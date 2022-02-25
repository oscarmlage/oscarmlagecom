---
title: "Cron: variables based on command output"
date: 2015-09-18T08:13:29Z
draft: false
tags: [ "sysadmin" ]
image: 41996742bd0a8110686906487cc1090a.jpg
---

<p>I had a problem with a cron job, a not-so-annoying but daily-repeating one. It's like a drill boring your mind slowly. Every time I got an email with the failing report the drill bit made more internal damage.</p>

```
/bin/sh: -c: unexpected EOF while looking for matching `''
/bin/sh: -c: line 1 systax error: unexpected end of file
```

<p>I was just trying to put a date into a file's name to backup a <a href="http://spamassassin.apache.org/">SpamAssassin</a> bayes database, something like <code>bayes-201508</code>, but it seemed that it didn't like to cron. I've tried it many ways:</p>

```
30 * * * * /usr/bin/sa-learn --backup &gt; /data/bayes-$(date +%Y%m).back
30 * * * * /usr/bin/sa-learn --backup &gt; /data/bayes-$(date '+%Y%m').back
30 * * * * /bin/bash -l -c '/usr/bin/sa-learn --backup &gt; /data/bayes-$(date +%Y%m).back'
```

<p>Thought it was caused to crontabs being running with <em>sh</em> instead of <em>bash</em> but at last nothing seemed to work in the proper way. I've experienced similar problems other times so I've decided to write the solution here: just escape the date call arguments:</p>

```
30 * * * * /usr/bin/sa-learn --backup &gt; /data/bayes-$(date +%Y%m).back
```

<p>And it will work like a charm. Hope this helps in the future.</p>
