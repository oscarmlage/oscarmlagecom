---
title: "Mastodon BOFHers update to 4.7.2"
date: 2026-09-18T18:10:00+02:00
draft: false
tags: ["mastodon", "sysadmin", "selfhosted"]
---

I manage a small [mastodon](https://joinmastodon.org/) instance for BOFHers. You know, old computer people armed with their whips, ready to lart anything that moves too close to a keyboard.

Doing admin work for that kind of crowd is a high risk sport. Doing it for free is even better. If something breaks, there is no mercy. If everything works, nobody says anything. So, the usual sysadmin deal.

Today I spent the afternoon updating the instance. It was not a single jump, but a few of them, one after another.

And I have to say, it felt nice to do some things by hand again in this AI driven world. Just me, a shell, some logs, and the usual fear of breaking production. No magic agent, no shiny button, no "we optimized your workflow". Only the old way: read, run, wait, swear a bit, continue.

- mastodon `4.3.4` → `4.4.25`
- mastodon `4.4.25` → `4.5.18`
- mastodon `4.5.18` → `4.6.8`
- mastodon `4.6.8` → `4.7.2`

The machine is the current BOFHers mastodon server: Debian 12, Ryzen 5 PRO 3600, 32GB RAM and a 1TB disk. Nothing too fancy, but enough for this. This post is mostly a note for future me. The last step was from `4.6.8` to `4.7.2`, and this is what I did.

## Before touching anything

First, clean a bit and check disk space.

```sh
$ make cleanmedia
$ df -h | grep md1
/dev/md1        933G  660G  227G  75% /
```

Then stop the instance and make a copy of the current directory. I know, not very elegant, but it is simple and it works.

```sh
$ make stop
$ time rsync -av --exclude='public/' --exclude='backup/' mastodon mastodon-copia-4.6.8

real    1m24.798s
```

## Get the new code

There were local changes, so I saved them first.

```sh
$ git stash
$ git stash list
stash@{0}: WIP on (no branch): 8f1847df31 Bump version to v4.6.8
```

Then move to `main`, pull, fetch tags and checkout the new version.

```sh
$ git switch main
$ git pull
$ git fetch --tags
$ git checkout v4.7.2
```

And then bring back the local changes.

```sh
$ git stash pop
Auto-merging docker-compose.yml
CONFLICT (content): Merge conflict in docker-compose.yml
```

Of course there was a conflict in `docker-compose.yml`, because there is always something. I fixed it and kept going.

## Build, migrations and restart

Build took around thirteen minutes.

```sh
$ time env DOCKER_BUILDKIT=1 COMPOSE_DOCKER_CLI_BUILD=1 docker-compose build

real    12m57.744s
```

Then the database migrations. First with post-deployment migrations skipped, as mastodon docs say for this kind of update.

```sh
$ time docker-compose run --rm -e SKIP_POST_DEPLOYMENT_MIGRATIONS=true web rails db:migrate
```

After that, stop again, clear cache, run the rest of migrations and start.

```sh
$ make stop
$ docker-compose run --rm web bin/tootctl cache clear
$ docker-compose run --rm web rails db:migrate
$ make start
```

## Clean docker stuff

After all the builds, Docker had a lot of old layers around. This recovered almost 100GB, which is not bad.

```sh
$ docker builder prune -f
Total reclaimed space: 95.41GB
```

And that was it. mastodon ended at `4.7.2`, the instance came back, and the BOFHers can keep doing BOFHers things.
