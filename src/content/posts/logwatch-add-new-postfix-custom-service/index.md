---
title: "Logwatch, add a new postfix custom service"
date: 2014-06-08T12:47:33Z
draft: false
tags: [ "sysadmin" ]
image: sysadmin.png
---

<p>Last days ago I had lot of problems with&nbsp;<em>MTA</em>&nbsp;servers,&nbsp;<a href="../../../../blog/postfix-deal-cracked-email-account.html">hacked accounts</a>, bad redirects policies and some other minor issues that kept me with the hands dirty in sysadmin keyboard. Needless to say how I love to put my&nbsp;<em>sysadm hat</em>&nbsp;and start to fix and/or optimize stuff, so from that bunch of problems was born this little script that makes my days easier than before.</p>
<p>We're talking about a&nbsp;<code>logwatch</code>&nbsp;custom service.&nbsp;<a href="http://sourceforge.net/projects/logwatch/files/">Logwatch</a>&nbsp;is a customizable log analysis system.&nbsp;<em>Logwatch</em>&nbsp;parses through your system's logs and creates a report analyzing areas that you specify.&nbsp;<em>Logwatch</em>&nbsp;is easy to use and will work right out of the package on most systems.&nbsp;I use&nbsp;<em>logwatch</em>&nbsp;to monitor common services running on servers. It sends me a daily report by mail telling me what happened last 24 hours. It's easy to add a new custom service, you have to put 3 files in the right place (<a href="https://www.debian.org/">Debian</a>&nbsp;like distribution):</p>
<ul>
<li><code>/etc/logwatch/conf/logfiles/my-postfix.conf</code>&nbsp;- Log configuration, the log files path and other minor options.</li>
<li><code>/etc/logwatch/conf/services/my-postfix.conf</code>&nbsp;- Service configuration, the title and the log file group we want to "inspect" (usually related to the above point).</li>
<li><code>/etc/logwatch/scripts/services/my-postfix</code>&nbsp;- The script that executes the command with the proper output you want to add in logwatch's report.</li>
</ul>
<h2>/etc/logwatch/conf/logfiles/my-postfix.conf</h2>

```
# /etc/logwatch/conf/logfiles/my-postfix.conf

# The LogFile path is relative to /var/log by default.
# You can change the default by setting LogDir.
LogFile = mail*.log

# This enables searching through zipped archives as well.
Archive = mail*.gz

# Expand the repeats (actually just removes them now).
*ExpandRepeats
```

<h2>/etc/logwatch/conf/services/my-postfix.conf</h2>

```
# /etc/logwatch/conf/services/my-postfix.conf

# The title shown in the report.
Title = "My Postfix"

# The name of the log file group (file name).
LogFile = my-postfix
```

<h2>/etc/logwatch/scripts/services/my-postfix</h2>

```
!/usr/bin/env bash
# /etc/logwatch/scripts/services/my-postfix

mailq | grep @ | awk '{ORS=(ORS==RS)?FS:RS; print $$NF}'
tot=`mailq | grep @ | awk '{ORS=(ORS==RS)?FS:RS; print $$NF}' | wc -l`

echo -e ""
echo "Total: ${tot}"
echo -e ""
echo -e ""
echo -e "Deferred emails from mail.log"
echo -e ""
grep "status" | grep -v "status=sent" | awk '{print $7" "$12}' | sort -rn | uniq -c | sort -rn
```

<h2>The output</h2>
<p>This is the report that the script sends me in the email, first part are the queued emails and the second part is a deferred list sorted by number of times it appears on mail.log:</p>

```
--------------------- My Postfix Begin ------------------------

 8A183B59       4347 Fri Jun  6 01:11:31  xxx@gmail.com -&gt; zzz@gmail.com
 88EE3B7C       2501 Thu Jun  5 16:28:42  xxx@domain.com -&gt; xxx@terra.es
 E16C1B3C        435 Thu Jun  5 13:34:28  xxx@xxx.kimsufi.com -&gt; root@xxx.kimsufi.com
 A4F3CB78       2501 Thu Jun  5 15:09:41  xxx@domain.com -&gt; zzz@terra.es
 AE0DBB8F       2501 Thu Jun  5 18:38:00  xxx@domain.com -&gt; zzz@terra.es
 AB746B6E       2501 Thu Jun  5 14:17:22  xxx@domain.com -&gt; zzz@terra.es
 AB5A289A        807 Wed Jun  4 06:26:47  xxx@xxx.kimsufi.com -&gt; root@xxx.kimsufi.com

 Total: 7

 Deferred emails from mail.log

     484 to=(root@xxxx.kimsufi.com), dsn=4.3.0,
     461 to=(zzz@terra.es), status=deferred
     170 to=(root@yyy.kimsufi.com), status=deferred
      56 to=(info@domain.com), status=deferred
      37 to=(zzz@gmail.com), dsn=4.7.0,
      31 to=(ooo@gmail.es), status=deferred
      31 to=(vvv@terra.es), status=deferred
      ...
       1 from=(soporte@domain.com),
       1 from=(root@vvv.kimsufi.com),

 ---------------------- My Postfix End -------------------------
```

<p>The main goal is to be able to take some decisions with a simple and quick glance.</p>
<h2>Recomendations</h2>
<p>As you can see in the last script, the bunch of files we selected to inspect (<code>mail*.log</code>) was the main input of the script, so we don't need to make a&nbsp;<code>cat</code>&nbsp;or something like that in the service script, they're treated as&nbsp;<em>STDIN</em>.</p>
<p>I must say too that you must activate&nbsp;<code>logrotate</code>&nbsp;on the logs to preserve&nbsp;<em>logwatch</em>&nbsp;eating cpu and harddisk for a long time. You can read more about how to add a service in logwatch&nbsp;<a href="http://alphahydrae.com/2012/08/logwatch-how-to-add-a-service/">here</a>.</p>
