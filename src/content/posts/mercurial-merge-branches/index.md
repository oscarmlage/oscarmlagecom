---
title: "Mercurial: Merge branches"
date: 2013-07-17T10:38:15Z
draft: false
tags: [  ]
image: 
---

<p>Tarde o temprano llegar&iacute;a la hora de trabajar con <em>branches</em>. Y ello lleva inequ&iacute;vocamente a tener que mezclarlas alg&uacute;n d&iacute;a. Suponiendo un escenario de repositorio con dos ramas (<em>default</em>, <em>new</em>), una vez el c&oacute;digo es estable y est&aacute; probado, haremos lo siguiente:</p>

```
$ hg branch
new
$ hg pull -u
$ hg commit -m "[ADD] New features"
$ hg push
$ hg update default
$ hg branch
default
$ hg merge new
$ hg commit -m "[MER] Merging default and new branches"
$ hg push
```

<p>El proceso es l&oacute;gico, estamos trabajando en la rama <em>new</em>, hacemos commit de todos los cambios, nos pasamos a la rama <em>default</em>, hacemos un <em>merge</em> de una con otra y <em>commiteamos</em> de nuevo la mezcla. Una vez hecho todo &eacute;sto ya volvemos a tener una sola rama <em>default</em> en la que seguir la linea de desarrollo.</p>
