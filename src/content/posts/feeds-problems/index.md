---
title: "Feeds problems"
date: 2024-10-28T13:11:14+01:00
draft: false
tags: []
---

The other day a blog reader - I assume - let me know that my feed content was broken. The content syndication was failing because, in every post that loaded an image with the theme's gallery shortcode, the code appeared instead of the images:

```html
<description>Can you almost smell it?
    if (!jQuery) { alert(&#34;jquery is not loaded&#34;); } $( document ).ready(() = { const gallery = $(&#34;
</description>
```

I imagine [the shortcode I’m using](https://github.com/mfg92/hugo-shortcode-gallery/) (which is part of a theme) doesn’t take in account the feeds, so I had to spend some time looking for a possible solution. I also realized that my RSS feeds weren't including the full content of each post, only an "excerpt" version. Big mistake on my part, as I've always ranted about that.

Getting down to it, the biggest challenge I faced was finding something that could be executed **before** the shortcode, as it would be more complicated to fix once the `{{ <gallery ...> }}` had been replaced. That’s where `RawContent` came to the rescue. By creating a new `home.xml` template for the feeds, I was able to edit the `description` tag to, initially, disable the galleries:

```html
<description>
    {{ .RawContent | replaceRE `<\s*gallery[^>]*>` "`[image not loaded in rss]`" | markdownify | htmlEscape | safeHTML }}
</description>
```

Finally, I also separated the feeds by category to keep things more organized. So, as of now, the following feeds are available:

- [Posts](https://oscarmlage.com/index.xml): Blog posts.
- [Microposts or Log](https://oscarmlage.com/microposts/index.xml): Microposts from the blog (most relevant updates from Mastodon / Bluesky).
- [Watching](https://oscarmlage.com/medialog/watching/index.xml): Series and movies I’m tracking.
- [Reading](https://oscarmlage.com/medialog/reading/index.xml): Books I’m reading.
- [Listening](https://oscarmlage.com/medialog/listening/index.xml): Albums currently playing on my Spotify.
- [Playing](https://oscarmlage.com/medialog/playing/index.xml): Video games I’m currently playing.

Thanks [@benalb](https://mastodon.bofhers.es/@benalb) for the heads-up.

—

**2024-10-28 Update**: Found [an issue](https://github.com/mfg92/hugo-shortcode-gallery/issues/36) related to this problem.