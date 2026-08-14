# mastodon-sync

Publishes new Hugo `posts` and `microposts` to Mastodon after deploy.

For microposts, if an `image` front matter value or `gallery/` image exists, the first image is attached to the Mastodon status. Normal posts never attach media.

It is idempotent: already published source URLs are stored in `utils/mastodon-sync/state.json`.

## Configuration

Copy the example env file and fill in your token:

```sh
cp utils/mastodon-sync/.env.example utils/mastodon-sync/.env
```

Required to publish:

```sh
MASTODON_INSTANCE=https://mastodon.bofhers.es
MASTODON_ACCESS_TOKEN=...
```

You can pass the release date and one-shot options from the CLI:

```sh
make mastodon-sync ARGS="--from 2026-06-25"
make mastodon-sync ARGS="--from 2026-06-25 --dry-run"
make mastodon-sync ARGS="--from 2026-06-25 --live" # overrides MASTODON_DRY_RUN=1
```

Optional env fallback / defaults:

```sh
MASTODON_RELEASE_DATE=2026-06-25 # fallback for --from
MASTODON_VISIBILITY=public       # public, unlisted, private, direct
MASTODON_DRY_RUN=1               # fallback for --dry-run
MASTODON_STRICT=1                # fail make deploy on sync errors
SITE_BASE_URL=https://oscarmlage.com
```

The real `utils/mastodon-sync/.env` file is ignored by git.

If `MASTODON_RELEASE_DATE` or `MASTODON_INSTANCE` are missing the tool skips silently, so deploys do not break on machines without Mastodon configured.

## Run with Docker

No local Go installation is needed. Use the Compose service:

```sh
make mastodon-sync
```

Dry-run example:

```sh
make mastodon-sync ARGS="--from 2026-06-25 --dry-run"
```

The CLI supports these flags, with env vars kept as fallbacks: `--instance`, `--token`, `--from`, `--visibility`, `--dry-run`, `--live`, `--strict`, `--max-chars`, `--base-url`, `--content-dir`, `--state-path`.

`make deploy` / `make up` runs this automatically after the `rsync`.
