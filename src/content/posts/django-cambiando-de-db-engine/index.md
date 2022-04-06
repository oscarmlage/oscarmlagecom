---
title: "Django: Cambiando de DB Engine"
date: 2012-04-28T18:42:31Z
draft: false
tags: [ "code", "django" ]
image: 
---

<p>
	Normalmente las especificaciones de un entorno de desarrollo y las de un entorno en producci&oacute;n suelen ser bastante diferentes. Quiz&aacute;s en el primero buscas la comodidad mientras que cuando se hace el <em>deploy</em> a producci&oacute;n priman otras cosas. Este suele ser el caso del motor de base de datos cuando trabajas con <a href="http://djangoproject.org">Django</a>.</p>
<p>
	Durante el proceso de desarrollo es muy c&oacute;modo utilizar <em>sqlite </em>porque no requiere de ning&uacute;n servidor adicional y est&aacute; soportado de base en Django (<em>builtin</em>). Es posible que una vez acabado el desarrollo queramos cambiar a otro <em>SGBD</em> &quot;de verdad&quot; como pueden ser <em>MySQL</em> o <em>PostgreSQL</em>.</p>
<p>
	Lo que podr&iacute;a suponer un quebradero de cabeza en otras arquitecturas, en nuestro caso utilizando <em>dumpdata</em> y <em>loaddata</em>, se reduce a las siguientes tres lineas:</p>

```
$ python manage.py dumpdata --indent=4 --format=json &gt; fixtures.json
$ scp fixtures.json user@remote:/path/project/
(remote)$ python manage.py loaddata fixtures.json
```

<p>
	En la primera usamos <a href="https://docs.djangoproject.com/en/dev/ref/django-admin/#dumpdata-appname-appname-appname-model">dumpdata</a>&nbsp;para hacer un dumpeado de todos los datos de la aplicaci&oacute;n (sqlite) en formato json, posteriormente pasamos ese archivo al entorno de producci&oacute;n (en este caso pongamos de ejemplo a otro servidor) y all&iacute; finalmente (con el conector apuntando al nuevo gestor) ejecutamos el <a href="https://docs.djangoproject.com/en/dev/ref/django-admin/#loaddata-fixture-fixture">loaddata</a> para insertar los datos.</p>
<p>
	Et voil&agrave;!</p>
