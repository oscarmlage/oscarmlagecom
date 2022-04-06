---
title: "Mercurial en Fedora Core 4 y CentOS 5"
date: 2010-03-08T17:22:14Z
draft: false
tags: [  ]
image: 
---

Necesitaba instalar <a href="http://mercurial.selenic.com/">Mercurial</a> en varias máquinas totalmente desactualizadas, concretamente una <em>Fedora Core release 4 (Stentz)</em> y una <em>CentOS release 5 (final)</em>, pensé que iba a ser un lío de dependencias pero al final ha resultado inexplicablemente más sencillo de lo esperado:
<!--more-->
<h2>Fedora Core 4</h2>

```
# cat /etc/fedora-release 
Fedora Core release 4 (Stentz)
# wget http://packages.sw.be/mercurial/mercurial-1.0.2-1.fc4.rf.i386.rpm
# rpm -Uvh mercurial-1.0.2-1.fc4.rf.i386.rpm
# hg --version
Mercurial Distributed SCM (version 1.0.2)
```


<h2>CentOS 5</h2>

```
# cat /etc/redhat-release 
CentOS release 5 (Final)
# wget http://packages.sw.be/mercurial/mercurial-1.3.1-1.el5.rf.i386.rpm
# rpm -Uvh mercurial-1.3.1-1.el5.rf.i386.rpm
# hg --version
Mercurial Distributed SCM (version 1.3.1)
```


<h2>Binarios</h2>
Vale que no sean la última versión (1.5 <abbr title="en este momento">atm</abbr>) pero estos binarios me han ahorrado bastante tiempo y evitado un problema, mola :).
