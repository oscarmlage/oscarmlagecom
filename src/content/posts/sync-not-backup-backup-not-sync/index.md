---
title: "Sync is not backup, backup is not sync."
date: 2016-09-18T09:52:16Z
draft: false
tags: [ "sysadmin" ]
image: raspberry-pi-berries.jpg
---

<p>I had an issue - <em>nothing new</em> - with the backups - <em>ouch</em> - and this time it was my fault one hundred percent.</p>
<p>Days ago I've bought a <a href="https://www.raspberrypi.org/">Raspberry Pi</a>, just to have fun with the possibilities (media center, retro games, download center...) you can do almost everything you can imagine. One of the ideas I had was to create a usb softraid with 2 external usb drives and add the raspi as <a href="https://syncthing.net/">Syncthing</a> node, because I'm using syncthing across some nodes to sync essential data, as backup (<strong>*err*</strong>).</p>
<p>While playing with the usb drives I noticed that the raspi was not able to feed both disks because of the limited voltage. It was running perfectly fine with one external disk but when attaching the second one we got weird noises in drives. Crazy.</p>
<p>Even crazier when noticed myself that syncthing was not down while playing with this stuff. And the craziest was to check that I've lost data in this process. It seems that something really weird happened to the main disk when attaching second one and it was like some folders just dissapeared.</p>
<p>I was able to recover all the data - or I think so - because I unexpectedly&nbsp;had a syncthing node that was down. It was what saved me this time. Moral of the fable: sync does not mean backup and viceversa.</p>
<p>I had to review the backup policy and just tune up the nodes making some completes and incremental data backups. Tools I've been using: <a href="http://rdiff-backup.nongnu.org/">rdiff-backup</a>, <a href="http://duplicity.nongnu.org/">duplicity</a>, <a href="http://duply.net/">duply</a> and <a href="https://syncthing.net/">syncthing</a>.</p>
<p>I've not finished yet but I think I did an further improvement of what I had. Time will tell. You just remember: <strong>Sync is not backup and backup is not sync</strong>.</p>
