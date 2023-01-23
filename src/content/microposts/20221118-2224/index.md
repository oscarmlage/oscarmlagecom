---
title: 20221118-2224
date: 2022-11-18 22:24:02 +00:00
draft: false
tags: [micropost]
image:
---

<p>From time to time I get an error like this trying to execute a <a href="https://mastodon.bofhers.es/tags/make" class="mention hashtag" rel="tag">#<span>make</span></a> custom command:</p><p>$ make backup<br />make: &#39;backup&#39; is up to date</p><p>I thought that something was wrong with that backup command but it seems the error is because there is a `backup/` directory sibling to the `Makefile`.</p><p>If you change the directory name or the command name, the error is gone. Weird, indeed.</p><p><a href="https://mastodon.bofhers.es/tags/micropost" class="mention hashtag" rel="tag">#<span>micropost</span></a> <a href="https://mastodon.bofhers.es/tags/makefile" class="mention hashtag" rel="tag">#<span>makefile</span></a></p>


