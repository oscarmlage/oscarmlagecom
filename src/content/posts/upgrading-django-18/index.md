---
title: "Upgrading Django to 1.8"
date: 2015-06-06T00:24:40Z
draft: false
tags: [  ]
image: django18.jpg
---

<p>As <a href="https://docs.djangoproject.com/en/1.8/howto/upgrade-version/">they said</a>, having last version has several benefits:</p>
<ul>
<li>New features and improvements are added.</li>
<li>Bugs are fixed.</li>
<li>Older version of Django will eventually no longer receive security updates.</li>
<li>Upgrading as each new Django release is available makes future upgrades less painful by keeping your code base up to date.</li>
</ul>
<p>Also it's even greater that <a href="https://www.djangoproject.com/weblog/2015/apr/01/release-18-final/">1.8</a> is a <code>LTS</code> (<em>Long-term support</em>), that means we will get bug fixes and security updates for a guaranteed period of time, typically 3+ years.</p>
<p>How the upgrade has affected to some <a href="https://docs.djangoproject.com/en/1.4/">Django 1.4</a> production code?, hard to say in few words, it depends on how the code is (+ 3rd party pieces), but main spots - <em>to take it carefully</em> - I've noted were:</p>
<ul>
<li>Changes in how settings are organized, take a look to your settings and try to split it in environment files (<a href="https://github.com/twoscoops/django-twoscoops-project/tree/develop/project_name/project_name/settings">tsd example</a>).</li>
<li><a href="https://docs.djangoproject.com/en/1.8/topics/migrations/">Migration tool</a> is now builtin (not depending on <a href="https://south.readthedocs.org/en/latest/">South</a> any more since <em>1.7</em>), take care while migrating between migrating tool <code>:P</code>.</li>
<li>I've never used URL's with namespaces before <code>({% url 'namespace:view' %})</code>. It's a nice feature available in <em>1.4</em> too (didn't know) if you pretend to reuse an app in more than a project.</li>
<li>In <em>urls.py</em> <code>django.conf.urls.default</code> is becoming in <code>django.conf.urls</code> (from <em>1.6</em>).</li>
<li>Cache key is now the <a href="https://docs.djangoproject.com/en/1.7/topics/cache/#using-vary-headers">ABSOLUTE URL</a> (full request required to set/get a key).</li>
<li><em>ManyToMany</em> field warning related to the stupid <code>null=True</code>.</li>
<li>Query filtering by <em>datetime</em> has changed a bit.</li>
<li>Changes in <a href="https://docs.djangoproject.com/en/1.8/howto/deployment/wsgi/">wsgi</a> and <a href="https://docs.djangoproject.com/en/1.8/howto/deployment/wsgi/gunicorn/">gunicorn</a> files (not using <em>gunicorn_django</em> any more).</li>
<li>Don't forget to enable <code><a href="https://docs.djangoproject.com/en/1.8/ref/settings/#allowed-hosts">ALLOWED_HOSTS</a></code> in your production settings (from <em>1.7</em>).</li>
</ul>
<p>It would be very tough if this above points pretend to be a guide or something like that. Not by a long shot. My recommendation is - of course - take a cup of your favourite drink, a bunch of patience... start reading the <a href="https://docs.djangoproject.com/en/1.8/releases/1.8/">release notes</a> and, above all, have as much fun as you can <code>:)</code>.</p>
