---
title: "Postfix resolving /etc/hosts entries"
date: 2015-10-26T21:52:30Z
draft: false
tags: [ "sysadmin" ]
image: 1fec99ef1cdd13e30fdec537f42ae21d.jpg
---

<p>Most of times when you are setting up a mail server you need to test that all is going fine. Well, in fact tests should happen *<em>after</em>* the server is configured. Talking about an integral email solution (send + receive), some kind of tests come to my mind:</p>
<ul>
<li>Auth tests v&iacute;a smtp/pop/imap</li>
<li>Send an email to a local existent/inexistent user</li>
<li>Send an email to a remote existent/inexistent user</li>
<li>Carefully test main internet email providers (gmail, yahoo and hotmail)</li>
<li>Receive an email from a local account</li>
<li>Receive an email from a remote account</li>
<li>Check filters and other custom stuff</li>
</ul>
<p>Nothing so complicate... unless your server is not a <em>DNS MX</em> endpoint. I mean if the server is not in "<em>production</em>" state (because we're configuring it) there shouldn't be any entry pointing to it, so how could we test the incoming email?.</p>
<p>We need to tweak other mail server (the one from which we're going to send the test email) to say: "hey friend, the email you're going to send is not for the real <em>MX</em> entry in <em>DNS</em> for example.com, please deliver it to this other IP". Thanks <a href="http://blog.e-shell.org/">Wu</a> for the idea. May you think how? of course, <code>/etc/hosts</code>:</p>

```
12.12.12.12    mail.example.com
```

<p>But, if the email server you're tweaking is Postfix you'll have to change a bit the configuration, because it's not taking care of <code>/etc/hosts</code> by default, you need to change a couple of directives (<a href="http://www.postfix.org/postconf.5.html#lmtp_host_lookup">lmtp_host_lookup</a>, <a href="http://www.postfix.org/postconf.5.html#smtp_host_lookup">smtp_host_lookup</a>) to say: "hey friend, take in account <code>/etc/hosts</code> before to go to dns stuff". Here we go, in <code>main.cf</code>:</p>

```
lmtp_host_lookup = native
smtp_host_lookup = native
```

<p>Now restart the postfix and write an email (<code>echo "testing" | main account@example.com</code>) while tailing logs to ensure all is happening in a right way.</p>
