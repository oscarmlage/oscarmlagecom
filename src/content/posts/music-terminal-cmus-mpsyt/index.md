---
title: "Music from terminal: cmus & mpsyt"
date: 2017-09-20T08:36:55Z
draft: false
tags: [  ]
image: music-from-terminal-cmus-mpsyt-01.png
---

<p>In a try to go back to my roots, lately, I'm using the terminal as much as I can and it's being a real pleasure. Even to listen music.</p>
<p>My consumption habits - talking about music - mostly go through <em>youtube</em> and local music, it depends on the situation. I was used to open whatever browser for <em>youtube</em> and whatever music player for local files (<em>vlc</em>, <em>itunes</em>...). Both (browsers and music players) are RAM drainers, so looked for a lighter solution.</p>
<p>I've found <a href="https://cmus.github.io/">cmus</a> as a really small and fast console music player, and it was so beautiful for my eyes so install and enjoy. If you're using <em>OSX</em> I recommend you to install <a href="https://github.com/PhilipTrauner/cmus-osx">cmus-osx</a> too, it gives you the control of cmus using the media keys on the keyboard, it's a must.</p>

```
$ brew install cmus
$ git clone https://github.com/azadkuh/cmus-osx.git
$ cd cmus-osx
$ pip3 install -r requirements.txt
$ ./setup.py install
$ cmus
```

<p>On the <em>youtube</em> hand the best I've found was <a href="https://github.com/mps-youtube/mps-youtube">mps-youtube</a> (<code>mpsyt</code>), a terminal based youtube player and downloader. It works with <code>youtube-dl</code> on background and it's able to search and play audio/video from youtube, create and import/export playlists... Enough for me.</p>

```
$ brew install mpv youtube-dl mps-youtube
$ mpsyt
```

<p>Enjoy the music from the dark side.</p>

{{< gallery match="gallery/*" sortOrder="asc" rowHeight="150" margins="5" thumbnailResizeOptions="600x600 q90 Lanczos"  previewType="blur" embedPreview="true" >}}

