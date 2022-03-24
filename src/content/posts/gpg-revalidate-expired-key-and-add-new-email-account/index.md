---
title: "GPG: Revalidate expired key and add new email account"
date: 2014-07-05T11:57:04Z
draft: false
tags: [ "sysadmin" ]
image: Email-with-lock.jpg
---

<p>When you are creating a new GPG key you have to decide how long the key will be valid before to renew. In my case when I got my own key, I was not really thinking on use it daily, it was years ago, just for test and have fun.</p>
<p>Nowadays I'm giving GPG a second - and more professional - chance, so when I tried to configure the key in client I realize that it was expired. How could I do to revalidate it?. Dunno if this is the oficial way, but it works for me:</p>

```
$ gpg --edit-key A0YOURK3Y
gpg&gt; list
gpg&gt; key 0
gpg&gt; expire

Changing expiration time for the primary key.
Please specify how long the key should be valid.

Key is valid for? (0) 2y
Key expires at 06/06/16 08:51:25
Is this correct? (y/N) y

gpg&gt; save
```

<p>At this time we have to send the update to public servers:</p>

```
$ gpg --keyserver pgp.mit.edu --send-keys A0YOURK3Y
```

<p>From the time I've created the key until now I've changed my main email account, so other of my main goals is to add the new email account to the key:</p>

```
$ gpg --edit-key A0YOURK3Y
gpg&gt; adduid
Name: &Oacute;scar
Email: new@emailaccount.com
Comment: This is my superb and new email account
gpg&gt; save
```

<p>We have to send the update to the public servers again:</p>

```
$ gpg --keyserver pgp.mit.edu --send-keys A0YOURK3Y
```

<p>And that's all, now it's time to see how to properly use it from an email client, new story.</p>
