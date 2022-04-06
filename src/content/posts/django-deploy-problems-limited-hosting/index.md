---
title: "Django deploy: problems with limited hosting"
date: 2014-07-21T21:51:18Z
draft: false
tags: [ "django", "code" ]
image: django.png
---

<p>Some months ago I had to deal with a&nbsp;<a href="../../../../blog/symfony2-en-un-hosting-compartido.html">Symfony2 project in a shared hosting</a>&nbsp;(Spanish article) and now the big next deal is a similar task with&nbsp;<a href="https://www.djangoproject.com/">Django</a>.</p>
<p>The project is almost done and I have the hosting credentials, once I'm in I noticed that there is no chance to configure anything (Apache, WSGI or whatever) so I was a bit lost. Thanks to my&nbsp;<a href="http://ailalelo.com/">Ailalelo's mates</a>&nbsp;(they had lot of experience with this kind of situations) I found the proper configuration.</p>
<p>Hosting is&nbsp;<em>django-ready</em>, but the version they're running (<em>1.4.2</em>) is not the best choice, I want to install&nbsp;<em>1.6.x</em>, the one I have used to develop the project. The other big requirement is&nbsp;<em>virtualenv+pip</em>&nbsp;to retrieve the packages I'm using.</p>
<p>Mainly I've solved it with two files,&nbsp;<code>project.cgi</code>&nbsp;and&nbsp;<code>.htaccess</code>.</p>
<h2>project.cgi</h2>
<p>The hosting structure is like many others, I have access to a&nbsp;<em>homedir</em>&nbsp;with the following directories:</p>

```
myhome/
  logs/
  tmp/
  www/
    cgi-bin/
    index.html
```

<p>Before to say Apache what to do with our project, let's install&nbsp;<code>virtualenv</code>&nbsp;and all the requirements, my choice is to put the environment out of the&nbsp;<code>www/</code>&nbsp;directory:</p>

```
myhome/
  env/
  logs/
  tmp/
  www/
    cgi-bin/
    project/
    index.html
```


```
$ virtualenv env
$ . env/bin/activate
$ pip install -r www/project/requirements/production.txt
```

<p>Seems to be that apache's&nbsp;<em>mod_cgi</em>&nbsp;will process all the files you put in the&nbsp;<code>cgi-bin</code>&nbsp;directory, so we already know where to save our&nbsp;<code>project.cgi</code>. I have to tell apache to use my own&nbsp;<em>virtualenv</em>&nbsp;<code>python</code>, where the environment and the project are. And finally set some environment variables:</p>

```
#!/home/myhome/env/bin/python

import os
from os.path import abspath, dirname
from sys import path

actual = os.path.dirname(__file__)
path.insert(0, os.path.join(actual, "..", "project/project"))
path.insert(0, os.path.join(actual, "..", "env"))

# Set the DJANGO_SETTINGS_MODULE environment variable.
os.environ['PATH'] = os.environ['PATH'] + ':/home/myhome/www/project/node_modules/less/bin'
os.environ['SECRET_KEY'] = 'SECRETKEY'
os.environ['DJANGO_SETTINGS_MODULE'] = "project.settings.production"

from django.core.servers.fastcgi import runfastcgi
runfastcgi(method="threaded", daemonize="false")
```

<p>Note that I modified the PATH because I have to be able to use&nbsp;<code>less</code>&nbsp;binary, required for&nbsp;<a href="http://django-compressor.readthedocs.org/en/latest/">django-compressor</a>&nbsp;package. In this particular case my user was not allowed to install&nbsp;<em>node/less</em>&nbsp;in the system, so I had to install it locally, referencing the particular&nbsp;<code>node_modules</code>&nbsp;folder.</p>
<h2>.htaccess</h2>
<p>Now that Apache knows what to do, we should redirect almost all the incoming traffic to the&nbsp;<em>cgi</em>, so let's write some<code>.htaccess</code>&nbsp;rules:</p>

```
AddHandler fcgid-script .fcgi
RewriteEngine On

RewriteRule ^static/(.*)$ project/project/static/$1 [L]
RewriteRule ^media/(.*)$ project/project/media/$1 [L]

RewriteCond %{REQUEST_FILENAME} !-f
RewriteRule ^(.*)$ cgi-bin/project.cgi/$1 [QSA,L]

AddDefaultCharset UTF-8
```

<p>The&nbsp;<code>media/</code>&nbsp;and&nbsp;<code>static/</code>&nbsp;dirs are redirected to the proper location, because they didn't work as is. Not much more to say with this file, it's easy to understand I think.</p>
<h2>Remember</h2>
<ul>
<li>To properly install the environment.</li>
<li>Set up this two files (cgi and .htaccess).</li>
<li>Take care with node tools (bower, npm, less, node_modules...).</li>
<li>Most of times django-compressor is a PITA in shared hostings.</li>
<li>Take double care with static/media settings.</li>
<li>Set the proper permissions in the .cgi file.</li>
<li>Having a bit of patience and access to&nbsp;<a href="http://stackoverflow.com/">stackoverflow</a>&nbsp;is to have it a half fixed ;).</li>
</ul>
<p>&nbsp;</p>
<ul>
</ul>
