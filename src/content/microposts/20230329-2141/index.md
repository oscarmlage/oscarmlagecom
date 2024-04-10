---
title: 20230329-2141
date: 2023-03-29 21:41:08 +00:00
draft: false
tags: [micropost]
image:
---

<p>TIL: How to SSH shim a dockerized Gitea by passing SSH through from the host to the container.</p><p>On host:</p><p>1. Share user (GID/UID) with the container env<br />2. Add new keys and mount .ssh as new volume<br />3. Mock the gitea binary on host to pass the SSH original command to the container</p><p>For the client:</p><p>1. Optionally add a new key to interact with the server<br />2. Add the pubkey to Gitea using the web interface<br />3. Be sure you&#39;re using the key and enjoy your clone :)</p><p><a href="https://mastodon.bofhers.es/tags/docker" class="mention hashtag" rel="tag">#<span>docker</span></a> <a href="https://mastodon.bofhers.es/tags/ssh" class="mention hashtag" rel="tag">#<span>ssh</span></a> <a href="https://mastodon.bofhers.es/tags/gitea" class="mention hashtag" rel="tag">#<span>gitea</span></a> <a href="https://mastodon.bofhers.es/tags/til" class="mention hashtag" rel="tag">#<span>til</span></a> <a href="https://mastodon.bofhers.es/tags/micropost" class="mention hashtag" rel="tag">#<span>micropost</span></a></p>


