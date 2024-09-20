---
title: "The Medialog"
date: 2024-08-27T21:00:52+02:00
draft: false
tags: [blog]
---

Lately, I've been diving - a bit more - into Hugo *templates* and *archetypes*, and I decided to create a **[medialog](/medialog/)** to track all the media I'm currently consuming: movies, series, books, games, music... This space helps me keep everything organized, and I'm hoping it won't be forgotten—like some other parts of this site have been!

### Creating the Medialog Layout

I created a new layout called `medialog.html` to define how I want my media log to be displayed. This layout includes sections for each type of media I'm tracking. To display the content, I used Hugo's `range` function, iterating over my taxonomies to showcase what I’ve been watching, reading, and more.

Here’s a snippet of how the layout looks:

```html
<main class="main wrapper">
    <h1>Medialog</h1>

    {{ if .Site.Taxonomies.tags.watching }}
    <h2>Watching</h2>
    <div class="media-grid horizontal">
        {{ range (.Site.Taxonomies.tags.watching).Pages }}
            {{ partial "medialog/watching.html" . }}
        {{ end }}
    </div>
    {{ end }}

    {{ if .Site.Taxonomies.tags.reading }}
    <h2>Reading</h2>
    <div class="media-grid vertical">
        {{ range (.Site.Taxonomies.tags.reading).Pages }}
            {{ partial "medialog/reading.html" . }}
        {{ end }}
    </div>
    {{ end }}

</main>
```

### Adding Content with Partials

I used Hugo's `partials` to keep things modular and reusable. Here’s an example of a partial that renders a media item with its rating:

```html
<a href="{{ .Params.Link }}" title="{{ .Title }}">
  <div class="item-wrapper">
    <div class="meta-text">
        <div class="gridheader">{{ .Title }}</div>
        <div class="subheader">
          {{ $rate := .Params.Rate }}
          {{ range $i, $e := (slice 0 1 2 3 4) }}{{ if lt $i $rate }}⭐{{ end }}{{ end }}
        </div>
      </div>
      <img src="listening/_assets/{{ .Params.Image }}"
        alt="{{ .Title }}" loading="lazy" decoding="async" />
    </div>
</a>
```

### Content structure

Each content entry in the `medialog` follows a simple structure. For example, a movie entry might look like this:

```ini
---
title: "Movie title"
link: https://www.imdb.com/title/ttXXX/
subtitle: 2018
year: 2018
rate: 3
image: movie-image.png
date: 2024-07-11T10:44:54Z
draft: false
tags: watching
---
```

### Handling Layout Issues

I ran into an issue where Hugo was rendering my `medialog` content using `default/list.html` instead of my custom layout. The solution was to properly configure taxonomies in `config.yaml` and ensure the layout was correctly referenced in the frontmatter. I also had to move `layouts/medialog.html` to `layouts/medialog/list.html` to get everything working properly.

###  Conclusion

Building this `medialog` has been a fun and rewarding experience. It’s a simple yet effective way to track my media consumption, and I’m excited to keep it updated (hopefully this time, it won’t fall by the wayside!). If you're interested in doing something similar, I encourage you to dive into Hugo and explore its powerful template features.