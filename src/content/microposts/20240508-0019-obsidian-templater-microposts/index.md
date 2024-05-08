---
title: obsidian-templater-microposts
date: 2024-05-08 00:19:30 +0200
draft: false
tags: micropost
---

Today I've built a new template for [Obsidian Templater](https://github.com/SilentVoid13/Templater) that allows the microposts creation inside [Hugo](https://gohugo.io) running a simple command:

```javascript
<%*
let qcFileName = await tp.system.prompt("Note Title")
titleName = tp.date.now("YYYYMMDD-HHmm-") + qcFileName
await tp.file.rename (titleName)
await tp.file.move("/microposts/" + titleName + "/index");
-%>
---
title: <% qcFileName %>
date: <% tp.file.creation_date("YYYY-MM-DD HH:mm:ss") %>
draft: false
tags: micropost
---
```

I'm really starting to love this kind of things.  [#obsidian](https://mastodon.bofhers.es/tags/obsidian) [#TIL](https://mastodon.bofhers.es/tags/til)
