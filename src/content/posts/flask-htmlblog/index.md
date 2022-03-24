---
title: "Flask-HTMLBlog"
date: 2012-06-07T00:32:24Z
draft: false
tags: [  ]
image: 
---

<p>
	As&iacute; de extra&ntilde;o suena el nombre del &uacute;ltimo proyecto que lleva dando vueltas entre las perdidas neuronas de mi extremidad superior. &iquest;De qu&eacute; se trata?, de un c&uacute;mulo de cosas...</p>
<p>
	Os habr&eacute;is dado cuenta que hace un tiempo me entretiene bastante todo lo relacionado con <a href="http://flask.pocoo.org/">Flask</a> y <a href="https://www.djangoproject.com/">Django</a>, pero como buen <em>newbie</em> que soy, no me sent&iacute;a c&oacute;modo sin un proyecto en el poder practicar el <em>learn by doing.<!--more--></em></p>
<p>
	Por otro lado, viendo lo abandonado que tengo <a href="http://userlinux.net">&eacute;sto</a> del blog y la de recursos que consume (esto &uacute;ltimo no es del todo cierto, pero me hab&iacute;a autoimpuesto mejorar el rendimiento), me he propuesto cambiar el gestor de contenidos por algo que sirva simplemente archivos est&aacute;ticos (&iquest;qu&eacute; es un post sino?).</p>
<p>
	La idea es sencilla: escribir archivos marcados en <em>html</em>, <em>rst</em>, <em>markdown</em>... y que un proceso los parsee y genere archivos est&aacute;ticos <em>.html</em> que autom&aacute;gicamente se subir&aacute;n a la carpeta <em>httpdocs</em> del servidor. Nada del otro mundo. Los principales mentores de esta idea han sido, entre otros&nbsp;<a href="http://tinkerer.bitbucket.org/">Tinkerer</a>, <a href="http://sphinx.pocoo.org">Sphinx</a>, <a href="http://disqus.com">Disqus</a>&nbsp;y&nbsp;<a href="https://github.com/mitsuhiko/lucumr">Lucumr</a>.</p>
<p>
	Hab&iacute;a hecho alguna que otra prueba con <a href="http://packages.python.org/Frozen-Flask/">Frozen-Flask</a> y me parec&iacute;a que pod&iacute;a encajar muy bien as&iacute; que en un alarde de productividad he creado la primera versi&oacute;n inestable 0.0.1. Si, es inestable, tiene mensajes <em>hardcodeados</em> y sin traducir, hay muchas cosas que todav&iacute;a no son configurables, el CLI de momento s&oacute;lo permite agregar y eliminar posts, el interfaz de ADMIN lo mismo (no tiene estad&iacute;sticas...), no soporta idiomas. Podr&iacute;a seguir enumerando <em>issues</em> pero me centrar&eacute; en que lo principal est&aacute; funcionando:</p>
<ul>
	<li>
		Agregar un post nuevo (desde CLI y desde ADMIN)</li>
	<li>
		Eliminar un post (desde CLI y desde ADMIN)</li>
	<li>
		Editar un post (desde CLI con cualquier editor de texto y desde ADMIN v&iacute;a web)</li>
	<li>
		Generar estructura de est&aacute;ticos .html (con paginaci&oacute;n, tags, b&uacute;squeda...)</li>
</ul>
<p>
	Cuando hablo de CLI y de ADMIN son dos interfaces distintas que he creado para interactuar con el sistema. Con CLI me refiero a una interacci&oacute;n en consola, se ve mejor con un par de ejemplos:</p>

```
$ ./manage.py newpost -t &quot;T&iacute;tulo del post&quot; -p &quot;articles/mi-primer-post&quot;
$ ./manage.py delpost articles/mi-primer-post
```

<p>
	Y cuando me refiero al ADMIN es un interfaz web mucho m&aacute;s sencillo en el que podemos, a trav&eacute;s de formularios, realizar las mismas acciones que por linea de comandos:</p>
<p style="text-align: center; ">
	<img alt="" src="gallery/flask-htmlblog.png" style="width: 450px; height: 248px; " /></p>
<p>
	Tambi&eacute;n he estado jugando un poco con <a href="http://readthedocs.org/">ReadTheDocs</a>, una fant&aacute;stica plataforma que nos permite leer online la documentaci&oacute;n del proyecto, as&iacute; que - aunque todav&iacute;a muy verde - aqu&iacute; hay algunos p&aacute;rrafos explicando qu&eacute; es Flask-HTMLBlog:</p>
<ul>
	<li>
		<a href="http://flask-htmlblog.readthedocs.org">Flask-HTMLBlog en ReadTheDocs</a></li>
</ul>
<p>
	El c&oacute;digo lo he subido a un repositorio en <a href="http://bitbucket.org">Bitbucket</a>, la idea es organizar y refactorizar todo lo que est&aacute; desordenado e ir implementando nuevas funcionalidades antes de hacer la migraci&oacute;n de Userlinux, puedes seguir las novedades a trav&eacute;s de:</p>
<ul>
	<li>
		<a href="https://bitbucket.org/r0sk/flask-htmlblog">Flask-HTMLBlog en Bitbucket</a></li>
</ul>
<p>
	Y, por supuesto, me encantar&iacute;a tener tu colaboraci&oacute;n en el proyecto, tanto con ideas como con lineas de c&oacute;digo que puedas aportar, as&iacute; que si&eacute;ntete libre de usarlo, seguirlo, clonarlo, <em>forkearlo</em> y mejorar este peque&ntilde;o generador de archivos html est&aacute;ticos.</p>
