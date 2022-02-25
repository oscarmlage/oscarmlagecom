---
title: "Music: new homekit"
date: 2017-08-30T08:26:03Z
draft: false
tags: [  ]
image: IMG_20170830_103325.jpg
---

<p>Not always but sometimes I need some <em>muzak</em> to focus myself in whatever I'm doing. To be honest I'm not the kind of people that wears (head|ear)phones all the time, I don't like, but from time to time it's ok with me to put them on to feel more isolated.</p>
<p>As I've said <a href="https://www.userlinux.net/banda-sonora-productiva.html">here</a> before, since I have a solid disk (yep, that ones that are fast as hell but size is not something to celebrate) I'm mostly using streaming music (<em>grooveshark</em>, <em>google play music</em>, <em>spotify</em>, <em>plex</em>) but... what if you're not at home with a reliable connection or no data plan?.</p>
<p>Shit happens, mostly when you're working remotely, so I've decided to <em>backup</em>&nbsp;my online lists - at least the most used ones - to an external hard drive following the next steps:</p>
<ul>
<li>As I've been using <code>youtube-dl</code> before, I've thought that best approach could be to reproduce all the lists in youtube and download them to different directories. So first step is to create manually the lists. In fact, as I was using youtube already, I had half of the work done.</li>
<li>Once all the lists are ready, just download them with <code>youtube-dl</code>. Easy peasy with this two commands:                      
<ul>
<li><code>$ youtube-dl --add-metadata --metadata-from-title "%(artist)s - %(title)s"</code> <br /> <code>--prefer-ffmpeg --no-post-overwrites -x -i --audio-format mp3 --audio-quality 320K</code><br /> <code>--embed-thumbnail -o "%(title)s.%(ext)s" https://youtube.com/url/to/list</code></li>
<li><code>$ youtube-dl --add-metadata -x --audio-format mp3 https://youtube.com/url/to/video</code></li>
</ul>
</li>
<li>Once I had the files "<em>physically</em>" on the hard drive I realized that some of them were in an unknown format for me: <a href="https://en.wikipedia.org/wiki/Opus_(audio_format)">opus</a>. Some players (like <code>VLC</code>) were able to play that format but <code>cmus</code> wasn't, so I needed to convert them to <em>mp3/ogg</em>...&nbsp;<code>ffmpeg</code> to the rescue:                     
<ul>
<li><code>for i in *.opus; do ffmpeg -i "$i" -f mp3 "${i%.*}.mp3"; done</code></li>
</ul>
</li>
</ul>
<p>Also, as extra step, I've bought <a href="https://www.amazon.es/gp/product/B0721JT49H">this adapter</a>, took an old microSD and <em>rsynced</em>&nbsp;the Music folder. And now I can enjoy it on the smartphone too without worrying about space or data plan limits.</p>
