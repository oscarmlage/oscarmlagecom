package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"html"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	Instance    string
	AccessToken string
	ReleaseDate time.Time
	BaseURL     string
	ContentDir  string
	StatePath   string
	Visibility  string
	DryRun      bool
	MaxChars    int
	Strict      bool
}

type Entry struct {
	Kind      string
	Title     string
	Date      time.Time
	Path      string
	Key       string
	URL       string
	Body      string
	Draft     bool
	ImagePath string
}

type StateEntry struct {
	MastodonID  string `json:"mastodon_id,omitempty"`
	MastodonURL string `json:"mastodon_url,omitempty"`
	PostedAt    string `json:"posted_at"`
	Kind        string `json:"kind"`
	Title       string `json:"title"`
}

type State map[string]StateEntry

type mastodonStatusResponse struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}

type mastodonMediaResponse struct {
	ID string `json:"id"`
}

var (
	frontMatterRe = regexp.MustCompile(`(?s)^---\s*\n(.*?)\n---\s*\n?`)
	shortcodeRe   = regexp.MustCompile(`(?s)\{\{<.*?>\}\}`)
	tagRe         = regexp.MustCompile(`(?s)<[^>]+>`)
	spaceRe       = regexp.MustCompile(`[ \t\r\f\v]+`)
	blankLinesRe  = regexp.MustCompile(`\n{3,}`)
)

func main() {
	cfg, ok, err := loadConfig()
	if errors.Is(err, flag.ErrHelp) {
		return
	}
	if err != nil {
		fail(cfg, err)
	}
	if !ok {
		fmt.Println("mastodon-sync: missing MASTODON_RELEASE_DATE or MASTODON_INSTANCE; skipping")
		return
	}

	state, err := loadState(cfg.StatePath)
	if err != nil {
		fail(cfg, err)
	}

	entries, err := findEntries(cfg)
	if err != nil {
		fail(cfg, err)
	}

	posted := 0
	for _, entry := range entries {
		if entry.Draft || entry.Date.Before(cfg.ReleaseDate) {
			continue
		}
		if _, exists := state[entry.Key]; exists {
			continue
		}

		status := formatStatus(entry, cfg.MaxChars)
		if cfg.DryRun {
			media := ""
			if entry.ImagePath != "" {
				media = "\n[media] " + entry.ImagePath
			}
			fmt.Printf("mastodon-sync: DRY-RUN would publish %s %s%s\n%s\n---\n", entry.Kind, entry.URL, media, status)
			continue
		}
		if cfg.AccessToken == "" {
			fail(cfg, errors.New("MASTODON_ACCESS_TOKEN/--token is required unless --dry-run is enabled"))
		}

		var mediaIDs []string
		if entry.ImagePath != "" {
			media, err := uploadMedia(cfg, entry.ImagePath)
			if err != nil {
				fail(cfg, fmt.Errorf("upload media %s: %w", entry.ImagePath, err))
			}
			mediaIDs = append(mediaIDs, media.ID)
		}

		resp, err := publishStatus(cfg, status, mediaIDs)
		if err != nil {
			fail(cfg, fmt.Errorf("publish %s: %w", entry.URL, err))
		}
		state[entry.Key] = StateEntry{
			MastodonID:  resp.ID,
			MastodonURL: resp.URL,
			PostedAt:    time.Now().Format(time.RFC3339),
			Kind:        entry.Kind,
			Title:       entry.Title,
		}
		if err := saveState(cfg.StatePath, state); err != nil {
			fail(cfg, err)
		}
		posted++
		fmt.Printf("mastodon-sync: published %s -> %s\n", entry.URL, resp.URL)
	}

	if posted == 0 && !cfg.DryRun {
		fmt.Println("mastodon-sync: nothing to publish")
	}
}

func loadConfig() (Config, bool, error) {
	maxChars := 500
	if v := os.Getenv("MASTODON_MAX_CHARS"); v != "" {
		n, err := strconv.Atoi(v)
		if err != nil || n < 100 {
			return Config{}, false, fmt.Errorf("invalid MASTODON_MAX_CHARS: %q", v)
		}
		maxChars = n
	}

	cfg := Config{
		Instance:    strings.TrimRight(os.Getenv("MASTODON_INSTANCE"), "/"),
		AccessToken: os.Getenv("MASTODON_ACCESS_TOKEN"),
		BaseURL:     strings.TrimRight(getenv("SITE_BASE_URL", "https://oscarmlage.com"), "/"),
		ContentDir:  getenv("MASTODON_CONTENT_DIR", "src/content"),
		StatePath:   getenv("MASTODON_STATE_PATH", "utils/mastodon-sync/state.json"),
		Visibility:  getenv("MASTODON_VISIBILITY", "public"),
		DryRun:      truthy(os.Getenv("MASTODON_DRY_RUN")),
		MaxChars:    maxChars,
		Strict:      truthy(os.Getenv("MASTODON_STRICT")),
	}

	from := os.Getenv("MASTODON_RELEASE_DATE")
	fs := flag.NewFlagSet("mastodon-sync", flag.ContinueOnError)
	fs.StringVar(&cfg.Instance, "instance", cfg.Instance, "Mastodon instance URL")
	fs.StringVar(&cfg.AccessToken, "token", cfg.AccessToken, "Mastodon access token")
	fs.StringVar(&from, "from", from, "publish entries dated on/after this date (YYYY-MM-DD)")
	fs.StringVar(&cfg.BaseURL, "base-url", cfg.BaseURL, "site base URL")
	fs.StringVar(&cfg.ContentDir, "content-dir", cfg.ContentDir, "Hugo content directory")
	fs.StringVar(&cfg.StatePath, "state-path", cfg.StatePath, "state JSON path")
	fs.StringVar(&cfg.Visibility, "visibility", cfg.Visibility, "Mastodon visibility: public, unlisted, private, direct")
	fs.BoolVar(&cfg.DryRun, "dry-run", cfg.DryRun, "preview without publishing")
	live := false
	fs.BoolVar(&live, "live", false, "force real publishing, overriding MASTODON_DRY_RUN")
	fs.BoolVar(&cfg.Strict, "strict", cfg.Strict, "exit non-zero on errors")
	fs.IntVar(&cfg.MaxChars, "max-chars", cfg.MaxChars, "maximum Mastodon status length")
	if err := fs.Parse(os.Args[1:]); err != nil {
		return cfg, false, err
	}

	if live {
		cfg.DryRun = false
	}
	cfg.Instance = strings.TrimRight(cfg.Instance, "/")
	cfg.BaseURL = strings.TrimRight(cfg.BaseURL, "/")
	if cfg.MaxChars < 100 {
		return cfg, false, fmt.Errorf("invalid --max-chars: %d", cfg.MaxChars)
	}
	if from == "" || cfg.Instance == "" {
		return cfg, false, nil
	}
	d, err := time.Parse("2006-01-02", from)
	if err != nil {
		return cfg, false, fmt.Errorf("invalid --from/MASTODON_RELEASE_DATE, expected YYYY-MM-DD: %w", err)
	}
	cfg.ReleaseDate = d
	return cfg, true, nil
}

func findEntries(cfg Config) ([]Entry, error) {
	var entries []Entry
	for _, kind := range []string{"microposts", "posts"} {
		root := filepath.Join(cfg.ContentDir, kind)
		if err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if d.IsDir() || filepath.Base(path) != "index.md" {
				return nil
			}
			entry, err := parseEntry(path, kind, cfg.BaseURL)
			if err != nil {
				return fmt.Errorf("%s: %w", path, err)
			}
			entries = append(entries, entry)
			return nil
		}); err != nil {
			return nil, err
		}
	}
	sort.Slice(entries, func(i, j int) bool { return entries[i].Date.Before(entries[j].Date) })
	return entries, nil
}

func parseEntry(path, kind, baseURL string) (Entry, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return Entry{}, err
	}
	text := strings.TrimLeft(string(raw), "\ufeff\n\r\t ")
	fm := map[string]string{}
	body := text
	if m := frontMatterRe.FindStringSubmatch(text); len(m) == 2 {
		fm = parseFrontMatter(m[1])
		body = text[len(m[0]):]
	}
	dt, err := parseDate(fm["date"])
	if err != nil {
		return Entry{}, err
	}
	slug := filepath.Base(filepath.Dir(path))
	postURL := fmt.Sprintf("%s/%s/%s/", baseURL, kind, slug)
	linkURL := postURL
	if kind == "microposts" {
		linkURL = fmt.Sprintf("%s/microposts/", baseURL)
	}
	imagePath := ""
	if kind == "microposts" {
		imagePath = firstMicropostImage(path, fm["image"])
	}
	return Entry{
		Kind:      strings.TrimSuffix(kind, "s"),
		Title:     cleanTitle(firstNonEmpty(fm["title"], slug)),
		Date:      dt,
		Path:      path,
		Key:       postURL,
		URL:       linkURL,
		Body:      plainText(body),
		Draft:     strings.EqualFold(strings.TrimSpace(fm["draft"]), "true"),
		ImagePath: imagePath,
	}, nil
}

func firstMicropostImage(indexPath, frontMatterImage string) string {
	baseDir := filepath.Dir(indexPath)
	if img := strings.TrimSpace(frontMatterImage); img != "" {
		candidate := filepath.Join(baseDir, img)
		if isImageFile(candidate) && fileExists(candidate) {
			return candidate
		}
	}

	galleryDir := filepath.Join(baseDir, "gallery")
	items, err := os.ReadDir(galleryDir)
	if err != nil {
		return ""
	}
	var images []string
	for _, item := range items {
		if item.IsDir() {
			continue
		}
		candidate := filepath.Join(galleryDir, item.Name())
		if isImageFile(candidate) {
			images = append(images, candidate)
		}
	}
	sort.Strings(images)
	if len(images) == 0 {
		return ""
	}
	return images[0]
}

func isImageFile(path string) bool {
	switch strings.ToLower(filepath.Ext(path)) {
	case ".jpg", ".jpeg", ".png", ".gif", ".webp":
		return true
	default:
		return false
	}
}

func fileExists(path string) bool {
	info, err := os.Stat(path)
	return err == nil && !info.IsDir()
}

func parseFrontMatter(s string) map[string]string {
	out := map[string]string{}
	for _, line := range strings.Split(s, "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") || strings.HasPrefix(line, "-") {
			continue
		}
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		val := strings.Trim(strings.TrimSpace(parts[1]), `"'`)
		out[key] = val
	}
	return out
}

func parseDate(s string) (time.Time, error) {
	s = spaceRe.ReplaceAllString(strings.Trim(strings.TrimSpace(s), `"'`), " ")
	layouts := []string{
		time.RFC3339,
		time.RFC1123Z,
		time.RFC1123,
		"2006-01-02 15:04:05 -0700",
		"2006-01-02 15:04:05 -07:00",
		"2006-01-02 15:04:05 MST",
		"2006-01-02",
	}
	for _, layout := range layouts {
		if t, err := time.Parse(layout, s); err == nil {
			return t, nil
		}
	}
	return time.Time{}, fmt.Errorf("unsupported date %q", s)
}

func plainText(s string) string {
	s = shortcodeRe.ReplaceAllString(s, "")
	s = strings.ReplaceAll(s, "</p>", "\n\n")
	s = strings.ReplaceAll(s, "<br>", "\n")
	s = strings.ReplaceAll(s, "<br/>", "\n")
	s = strings.ReplaceAll(s, "<br />", "\n")
	s = tagRe.ReplaceAllString(s, "")
	s = html.UnescapeString(s)
	lines := strings.Split(s, "\n")
	for i, line := range lines {
		lines[i] = strings.TrimSpace(spaceRe.ReplaceAllString(line, " "))
	}
	s = strings.Join(lines, "\n")
	s = blankLinesRe.ReplaceAllString(s, "\n\n")
	return strings.TrimSpace(s)
}

func formatStatus(e Entry, max int) string {
	link := "🔗 " + e.URL
	var prefix string
	if e.Kind == "post" {
		prefix = "New post: " + e.Title
		if e.Body != "" {
			prefix += "\n\n" + e.Body
		}
	} else {
		prefix = e.Body
	}

	available := max - len([]rune(link)) - 2
	if available < 0 {
		available = 0
	}
	prefix = truncateRunes(strings.TrimSpace(prefix), available)
	if prefix == "" {
		return link
	}
	return prefix + "\n\n" + link
}

func uploadMedia(cfg Config, imagePath string) (mastodonMediaResponse, error) {
	file, err := os.Open(imagePath)
	if err != nil {
		return mastodonMediaResponse{}, err
	}
	defer file.Close()

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)
	part, err := writer.CreateFormFile("file", filepath.Base(imagePath))
	if err != nil {
		return mastodonMediaResponse{}, err
	}
	if _, err := io.Copy(part, file); err != nil {
		return mastodonMediaResponse{}, err
	}
	if err := writer.Close(); err != nil {
		return mastodonMediaResponse{}, err
	}

	req, err := http.NewRequest(http.MethodPost, cfg.Instance+"/api/v2/media", &body)
	if err != nil {
		return mastodonMediaResponse{}, err
	}
	req.Header.Set("Authorization", "Bearer "+cfg.AccessToken)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("User-Agent", "oscarmlage-hugo mastodon-sync")

	client := http.Client{Timeout: 60 * time.Second}
	res, err := client.Do(req)
	if err != nil {
		return mastodonMediaResponse{}, err
	}
	defer res.Body.Close()
	respBody, _ := io.ReadAll(io.LimitReader(res.Body, 4096))
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return mastodonMediaResponse{}, fmt.Errorf("mastodon media API returned %s: %s", res.Status, string(respBody))
	}
	var parsed mastodonMediaResponse
	if err := json.Unmarshal(respBody, &parsed); err != nil {
		return mastodonMediaResponse{}, err
	}
	return parsed, nil
}

func publishStatus(cfg Config, status string, mediaIDs []string) (mastodonStatusResponse, error) {
	endpoint := cfg.Instance + "/api/v1/statuses"
	form := url.Values{}
	form.Set("status", status)
	form.Set("visibility", cfg.Visibility)
	for _, id := range mediaIDs {
		form.Add("media_ids[]", id)
	}

	req, err := http.NewRequest(http.MethodPost, endpoint, strings.NewReader(form.Encode()))
	if err != nil {
		return mastodonStatusResponse{}, err
	}
	req.Header.Set("Authorization", "Bearer "+cfg.AccessToken)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", "oscarmlage-hugo mastodon-sync")

	client := http.Client{Timeout: 20 * time.Second}
	res, err := client.Do(req)
	if err != nil {
		return mastodonStatusResponse{}, err
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(io.LimitReader(res.Body, 4096))
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return mastodonStatusResponse{}, fmt.Errorf("mastodon API returned %s: %s", res.Status, string(body))
	}
	var parsed mastodonStatusResponse
	if err := json.Unmarshal(body, &parsed); err != nil {
		return mastodonStatusResponse{}, err
	}
	return parsed, nil
}

func loadState(path string) (State, error) {
	state := State{}
	b, err := os.ReadFile(path)
	if errors.Is(err, os.ErrNotExist) {
		return state, nil
	}
	if err != nil {
		return nil, err
	}
	if len(bytes.TrimSpace(b)) == 0 {
		return state, nil
	}
	return state, json.Unmarshal(b, &state)
}

func saveState(path string, state State) error {
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}
	b, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, append(b, '\n'), 0644)
}

func fail(cfg Config, err error) {
	if cfg.Strict {
		fmt.Fprintln(os.Stderr, "mastodon-sync:", err)
		os.Exit(1)
	}
	fmt.Fprintln(os.Stderr, "mastodon-sync:", err)
	os.Exit(0)
}

func getenv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func truthy(v string) bool {
	switch strings.ToLower(strings.TrimSpace(v)) {
	case "1", "true", "yes", "y", "on":
		return true
	default:
		return false
	}
}

func firstNonEmpty(vals ...string) string {
	for _, v := range vals {
		if strings.TrimSpace(v) != "" {
			return v
		}
	}
	return ""
}

func cleanTitle(s string) string {
	return strings.Trim(strings.TrimSpace(html.UnescapeString(s)), `"'`)
}

func truncateRunes(s string, max int) string {
	r := []rune(s)
	if len(r) <= max {
		return s
	}
	if max <= 1 {
		return string(r[:max])
	}
	return strings.TrimSpace(string(r[:max-1])) + "…"
}
