---
title: "IRSSI: Hilighting your nick"
date: 2020-02-11T20:59:09Z
draft: false
tags: [  ]
image: irssi-hilighting-your-nick.jpg
---

<p>I'm still - and proudly - using <a href="https://es.wikipedia.org/wiki/Internet_Relay_Chat">irc</a> for daily contact with my team mates and it's perfectly fine for our purposes. Well, I miss so much a couple of things (like edit a message with typos or the ability to easily share a photo/screenshot) but we're witty enough to manage among us. Other than that it's really fun.</p>
<p>Lately, due the amount of log, I've found something a bit weird with my <a href="https://irssi.org/">irssi</a> config when my is mentioned. I have a couple of perl scripts that sends me a notification to the phone when someone writes my nick and it seemed to me that it was working randomly but... it wasn't.</p>
<p>Just in case you're wondering about that scripts, I have two, one of them working with Android devices called <a href="https://irssinotifier.appspot.com/">irssinotifier.pl</a>&nbsp;and as I've lately switched to iOS I got <a href="https://www.prowlapp.com/">prowl.pl</a>&nbsp;(not free but it's so cheap) as the other one working in the Apple world.</p>
<p>What was the problem then? It seems that irssi was highlighting only when someone mentioned me at the beginning of the phrase like this&nbsp;<code>r0sk:&nbsp;</code>but it wasn't doing it in a normal nick mention in the middle of a sentence like <code>hey r0sk sup there!?</code> or similar.</p>
<p>I knew that something like this could be configured so reading a bit about <a href="https://irssi.org/documentation/help/hilight/">irssi hilight</a> found the way to be alerted even if they are shouting <code>r000ssskkk!!</code> to me, the solution can be applied in two ways, in the irssi command line:</p>

```
/hilight -regexp r0+s+k+
```

<p>And, of course, in the <code>.irssi/config</code> file, because if you have put it in the command line and the setting is not saved to the config file, it won't work in case of restart:</p>

```
hilights = (
  { text = "r0+s+k+"; nick = "yes"; word = "yes"; regexp = "yes"; }
)
```

<p>And that's the reason why my phone would beep even if you get so angry with me.</p>
