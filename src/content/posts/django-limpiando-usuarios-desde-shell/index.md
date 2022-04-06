---
title: "Django: limpiando usuarios desde shell"
date: 2013-09-12T23:03:29Z
draft: false
tags: [ "django", "code" ]
image: 
---

<p>Es tal el inter&eacute;s por el <a href="http://www.vamosavidas.com/foros/">foro de VamosaVidas</a> (si, ese que romp&iacute; en su d&iacute;a y <a href="http://www.vamosavidas.com/los-duendes-del-foro-nos-lo-han-devuelto.html">he vuelto a arreglar</a>) que todos los d&iacute;as tenemos cientos de registros. Realmente se trata de robots que se dedican a hacer spam, pero eso es lo de menos :P.</p>
<p>Mientras no actualizo y arreglo <a href="https://github.com/pennersr/django-allauth">django-authall</a>&nbsp;es un poco co&ntilde;azo andar mirando en el admin el &uacute;ltimo usuario bueno y borrar manualmente uno a uno cotejando que no me equivoque, as&iacute; que he pensado en un par de lineas que no deber&iacute;an hacer demasiado da&ntilde;o si se usan con precauci&oacute;n. Desde la shell de nuestro proyecto podemos hacer algo as&iacute;:</p>

```
$ python manage.py shell
&gt;&gt;&gt; from django.contrib.auth.models import User
&gt;&gt;&gt; list(User.objects.order_by('date_joined'))
&gt;&gt;&gt; [user.delete() for user in list(User.objects.order_by('date_joined'))[100:]]
```

<p>Primero estamos listando todos los usuarios por orden de fecha de registro, los convertimos en lista para que no salgan los datos "<em>truncados</em>". Y en segundo lugar estamos eliminando todos excepto los 100 primeros.</p>
<p>Esta es la sencillez y potencia de Python.</p>
