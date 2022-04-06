---
title: "Mercurial: Hook on push"
date: 2010-02-21T13:21:26Z
draft: false
tags: [  ]
image: 
---

Lo tenía pendiente desde que cambié de Subversion a Mercurial, sabía que se podía y que era algo trivial pero lo vas dejando y bueno, <em>just happens</em>. El caso es que cuando haces un <em>push</em> al servidor lo normal es hacer un <em>update</em> de su lado, así que el <em>hook</em> que lo automatiza es el correspondiente:

```
[server]$ cat /path/del/repo/.hg/hgrc
[hooks]
changegroup = hg update
```

Y fuera preocupaciones. La obligada referencia a <a href="http://the.taoofmac.com/space/blog/2007/09/14/1130">Tao of Mac</a>, desde donde me vino la inspiración mientras intentaba poner el <em>Macbook</em> a punto.
