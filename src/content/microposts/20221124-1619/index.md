---
title: 20221124-1619
date: 2022-11-24 16:19:16 +00:00
draft: false
tags: [micropost]
image:
---

<p>Maybe your <a href="https://mastodon.bofhers.es/tags/mastodon" class="mention hashtag" rel="tag">#<span>mastodon</span></a> instance is playing tricks while uploading GIF files (dunno why it&#39;s not happening with other image formats):</p><p>413 Entity too large.</p><p>In our case this is a dockerized instance with `jwilder/nginx-proxy` in the middle, the fix is easy:</p><p>1. Use a nginx config volume: `./data/conf:/etc/nginx/conf.d`<br />2. create a `client_max_body_size.conf`file in the volume with &quot;client_max_body_size 100m;&quot; inside.<br />3. Restart the instance</p><p><a href="https://mastodon.bofhers.es/tags/docker" class="mention hashtag" rel="tag">#<span>docker</span></a> <a href="https://mastodon.bofhers.es/tags/nginx" class="mention hashtag" rel="tag">#<span>nginx</span></a> <a href="https://mastodon.bofhers.es/tags/jwilder" class="mention hashtag" rel="tag">#<span>jwilder</span></a> <a href="https://mastodon.bofhers.es/tags/nginxproxy" class="mention hashtag" rel="tag">#<span>nginxproxy</span></a> <a href="https://mastodon.bofhers.es/tags/micropost" class="mention hashtag" rel="tag">#<span>micropost</span></a> <a href="https://mastodon.bofhers.es/tags/repost" class="mention hashtag" rel="tag">#<span>repost</span></a></p>


