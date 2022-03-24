---
title: "Darcs: otra forma de trabajar"
date: 2013-01-11T17:18:13Z
draft: false
tags: [ "code" ]
image: 
---

<p>Si tengo que ser sincero, desde un principio he sido bastante reacio a la moda esta de los sistemas de control de versiones. Estaba tan agusto y acomodado con mis copias de seguridad y mis carpetas <em>fechadas</em> que me daba una pereza enorme ponerme a aprender un nuevo sistema que, sin duda, me obligar&iacute;a a cambiar la forma de trabajar.</p>
<p>La verdad, uno no sabe lo que se est&aacute; perdiendo hasta que lo prueba y aprende. Primero y muy vagamente <a href="http://cvs.nongnu.org/">CVS</a>, luego <a href="http://subversion.tigris.org/">Subversion</a> y posteriormente el gran salto de calidad pasando a los sistemas de control de versiones distribuidos. Me he quedado con el m&aacute;s vers&aacute;til dada mi forma de trabajar: <a href="http://mercurial.selenic.com/">Mercurial</a>.</p>
<p>Pero si algo he aprendido desde esa acomodada posici&oacute;n inicial es a no cerrar posibilidades, as&iacute; que gracias a un proyecto en el que estoy trabajando, he tenido la oportunidad de probar <a href="http://darcs.net/">Darcs</a>, eso si,&nbsp;a otro nivel.</p>
<p><img style="display: block; margin-left: auto; margin-right: auto;" src="gallery/>Y digo "a otro nivel" puesto que la forma de trabajar es m&aacute;s parecida a un proyecto <em>open source</em>&nbsp;en el que hay un <em>core team </em>que verifica y valida el trabajo del resto de la comunidad, pasando el c&oacute;digo un control de calidad previo y una fant&aacute;stica capa de&nbsp;testeo antes de su puesta en producci&oacute;n.&nbsp;El flujo ser&iacute;a el siguiente:</p>
<ul>
<li>Pull de las &uacute;ltimas novedades del proyecto</li>
<li>Se trabaja en las nueva caracter&iacute;sticas</li>
<li>Se "graban" los cambios peri&oacute;dicamente</li>
<li>Cuando varias de esas grabaciones completan una caracter&iacute;stica se crea un parche con todas esas grabaciones y se env&iacute;a al <em>maintainer</em>.</li>
<li>El <em>maintainer</em> verifica la validez - o no - del c&oacute;digo y tiene la capacidad para hacer push.</li>
<li>Dado el caso, se dispara un despliegue a preproducci&oacute;n o a producci&oacute;n directamente.</li>
</ul>

```
$ darcs pull
$ darcs whatsnew -l # muestra status
$ darcs record # hacer commit
$ darcs send -O # crear patch
$ darcs amend
$ darcs push
$ darcs apply parche.patch
$ darcs unpull
```

<p>Como v&eacute;is las &oacute;rdenes son id&eacute;nticas a cualquier otro sistema de control de versiones, de hecho podr&iacute;amos trabajar de forma id&eacute;ntica a como lo hacemos habitualmente con <em>Mercurial</em> o <a href="http://git-scm.com/">Git</a>, (<em>pull</em>, <em>commit</em>, <em>push</em>, etc...) sin embargo est&aacute; resultando una experiencia muy enriquecedora el tener otro flujo distinto al habitual.</p>
<p>En concreto hay una caracter&iacute;stica que me encanta de <em>Darcs</em>, y es el poder <em>commitear</em> partes de un archivo y no el archivo completo. Por ejemplo si hacemos dos cambios distintos en el mismo fichero, uno de ellos corrigiendo un bug y otro porque vamos a empezar una nueva <em>feature</em>, podemos hacer <em>record</em> (<em>commit)</em> del cambio que nos interese, no tiene por qu&eacute; reflejarse el archivo por completo. Algo que en mercurial por ejemplo no se puede hacer a no ser que se active la <a href="http://mercurial.selenic.com/wiki/RecordExtension">extensi&oacute;n record</a>.</p>
<p>Reconozco que todav&iacute;a me falta entrar en detalle con <em>Git</em>, pero al final todos los sistemas de control de versiones que he probado han acabado por gustarme de una u otra forma. As&iacute; que la moraleja ser&iacute;a sencilla, no comentas el mismo error que en un principio he cometido yo, prueba los que puedas y qu&eacute;date con el que m&aacute;s te guste, pero si desarrollas, dcvs es un <em>musthave</em>.</p>
