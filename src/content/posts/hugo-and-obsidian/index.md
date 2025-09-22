---
title: Hugo and Obsidian
date: 2024-12-23T16:20:30+01:00
draft: false
tags: []
---

Some time ago, I decided to integrate [Hugo](https://gohugo.io/), the static site generator, with [Obsidian](https://obsidian.md/), my favorite note-taking tool, so that writing becomes the only obstacle to publishing new content. This post explains how I set it up and shares some tips to reduce the friction between writing, generating, and publishing.

## Hugo

My use of [Hugo](https://gohugo.io/) is limited to my development machine, where I have it set up to run in a [Docker](https://www.docker.com/) container in a relatively simple way. Once it's running, I access to `http://localhost:1313` to preview the post (or any changes) on a local development server. Here’s my `docker-compose.yml`:

```yaml
version: "3.1"

services:

  build:
    image: klakegg/hugo:0.88.0-ext-alpine
    command: build
    volumes:
      - "${PWD}/src:/src"

  serve:
    image: klakegg/hugo:0.88.0-ext-alpine
    command: server -vv
    volumes:
      - "${PWD}/src:/src"
    ports:
      - 1313:1313
    stdin_open: true
    tty: true

```

I could use `docker compose up -d`, but we *lazy folks* are always looking for shortcuts, so for me, it’s just `make start`.

Once I’m happy with the new content, I run a synchronization script against the production server. It could be something like `rsync -e ssh -avz . user@prod:/path/to/oscarmlage.com`, but again, I prefer shortcuts. For me, it’s just a simple `make deploy`.

The black magic that handles all of this is called a `Makefile`. It’s located at the root of the project and looks something like this:

```yaml
build:
	docker-compose -f docker-compose.yml up build

serve:
	docker-compose -f docker-compose.yml up serve

bash:
	docker-compose -f docker-compose.yml exec serve bash

shell:
	docker-compose -f docker-compose.yml run build shell

deploy: up
sync: up
up: build
	rsync -e 'ssh -p 22235' -avz src/public/ user@prod:/path/to/oscarmlage/
```

To give you an idea, here’s the project structure:

```ini
.
├── Makefile
├── docker-compose.yml
├── src
│   ├── config.yml
│   ├── archetypes
│   │   ├── default.md
│   │   └── posts
│   │       ├── gallery
│   │       │   └── delete.me
│   │       └── index.md
│   ├── content
│   │   ├── posts
│   │   │   ├── a-movie-for-a-change
│   │   │   │   └── index.md

```

There’s not much more to say about the *Hugo* configuration. The only notable detail is that the archetype for posts, as seen in the *skel*, ensures that each post has its own directory with an `index.md` inside, along with a `gallery/` directory if there’s a need to attach any photos or galleries.

## Obsidian

On the *Obsidian* side, the goal was to keep the workflow as smooth as possible without duplicating efforts. The first decision was whether to create a separate vault exclusively for the blog or to use the same vault I already had for my daily notes. Since I wasn’t keen on setting up another vault with the plugins, theme, and other configurations I’d need (I guess programmers are naturally like this :P), I decided to use the same one. The trick in this case was to add `src/content` to my vault (located in a completely different directory). Nothing a *symbolic link* can’t fix:

```sh
$ ls ~/vaults/Personal/Projects/oscarmlage.com/
total 0
0 drwxr-xr-x@ Oct 25 12:59 .
0 drwxr-xr-x@ Nov 14 18:55 ..
0 lrwxr-xr-x  Oct 25 12:59 medialog -> ~/oscarmlage.com/src/content/medialog
0 lrwxr-xr-x  Oct 25 12:59 microposts -> ~/oscarmlage.com/src/content/microposts
0 lrwxr-xr-x  Oct 25 12:59 posts -> ~/oscarmlage.com/src/content/posts
[...]
```

This way, from *Obsidian*, I can create new content directly in the correct location. However, I wasn’t completely satisfied because, as I mentioned earlier, creating a post requires setting up a directory with the slug, an `index.md`, and the frontmatter… and that’s something that would be fantastic to automate. This is where [Templater](https://github.com/SilentVoid13/Templater) comes into play — a plugin that allows us to do exactly that, among many other things: executing *custom actions* when creating new content within a directory.

{{< gallery match="gallery/*.png" sortOrder="asc" rowHeight="150" margins="5" thumbnailResizeOptions="600x600 q90 Lanczos" previewType="blur" embedPreview="true" >}}

As shown in the screenshot, Templater's configuration includes an option called *Folder templates*, which triggers a "script/template" every time a new file is created in that folder. We select the folder for our blog posts on one side, and the `post.md` template on the other, which would look something like this:

```js
<%*
function toSlug(title) {
    return title
        .toLowerCase()
        .normalize("NFD")
        .replace(/[\u0300-\u036f]/g, "")
        .replace(/[^a-z0-9\s-]/g, "")
        .trim()
        .replace(/\s+/g, "-");
}

let qcFileName = await tp.system.prompt("Post Title")
let titleName = toSlug(qcFileName);
await tp.file.rename (titleName);
await tp.file.move("/Projects/oscarmlage.com/posts/" + titleName + "/index");
-%>
---
title: <% qcFileName %>
date: <% tp.file.creation_date("YYYY-MM-DDTHH:mm:ssZ") %>
draft: false
tags: []
---

```

As we can see in the code above, every time we add a new post to the configured directory (the symbolic link we previously created connecting `src/content` from _Hugo_ with _Obsidian_), the flow of what _Obsidian_ does is as follows:

- A prompt asks us for the post title: `Hugo and Obsidian`.
- It converts that title into a slug.
- Renames and moves the file to its proper location: `posts/hugo-and-obsidian/index.md`.
- Adds the _frontmatter_ to the file, so all we need to focus on is writing the post content.

## Flow

My workflow for creating new posts with this integration is quite simple and efficient. First, I open _Obsidian_ and navigate to the directory where the blog posts are stored. There, I create a new note and assign the post title in the modal. Thanks to the integration with _Templater_, the trigger automatically runs and applies the predefined template with everything needed for _Hugo_ (archetype, date, tags...). Once the trigger finishes, I simply start writing the post content directly in _Obsidian_.

<video width=100% controls autoplay>
    <source src="gallery/hugo-and-obsidian-01.mp4" type="video/webm">
    Your browser does not support the video tag.  
</video>


Since [version 1.8](https://obsidian.md/changelog/2024-12-18-desktop-v1.8.0/), *Obsidian* includes a Core Plugin called **Web Viewer**, which allows us to have the web page and the `index.md` side by side. This way, as we write in one *Obsidian* pane, we can see the result in the other in real time. All you need is *Hugo* serving locally (`make serve`) and pointing the **Web Viewer** to the new post's URL: `http://localhost:1313/posts/hugo-and-obsidian/`.

<video width=100% controls autoplay>
    <source src="gallery/blog-edit-obsidian.mp4" type="video/webm">
    Your browser does not support the video tag.  
</video>

Once the creation process is complete and the content has been reviewed, we return to our beloved terminal and run `make deploy` to synchronize the content with the production server. The content is now published and available to everyone.

<video width=100% controls autoplay>
    <source src="gallery/make-deploy.mp4" type="video/webm">
    Your browser does not support the video tag.  
</video>

This integration between **Hugo** and **Obsidian** allows me to maintain a smooth and centralized workflow, reducing friction when creating and publishing content. It might also fit into your publishing process — perhaps not the entire process, but if you can make use of some of the steps, this post will have been worth it.