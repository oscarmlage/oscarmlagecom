---
title: "Pyramid: 2 - Entendiendo la estructura del proyecto"
date: 2013-10-31T19:00:00Z
draft: false
tags: [ "python", "code" ]
image: 
---

<p>En la <a href="../../../../pyramid-1-montando-entorno-y-proyecto-zodb.html">anterior entrada</a> hemos visto lo m&aacute;s b&aacute;sico de c&oacute;mo crear un entorno y un proyecto inicial con <strong>Pyramid + ZODB</strong>, en este post vamos a intentar entender el proyecto y adaptarlo un poco a nuestra forma de trabajar, para ello nada mejor que empezar ech&aacute;ndole un vistazo a lo que <em>pcreate</em> ha hecho por nosotros:</p>

```
$ cd src/
$ ls
CHANGES.txt     Data.fs.lock    README.txt      setup.cfg       src.egg-info/
Data.fs         Data.fs.tmp     development.ini setup.py
Data.fs.index   MANIFEST.in     production.ini  src/
```

<p>En este primer nivel, adem&aacute;s de los correspondientes ficheros de texto README y CHANGES, debemos destacar varios elementos:</p>
<ul>
<li>Archivos <strong>.ini</strong> (developent.ini, production.ini...): Son los archivos de configuraci&oacute;n de entorno. Pyramid sabe que habr&aacute; configuraciones distintas dependiendo del entorno as&iacute; que nos ofrece la posibilidad de utilizar archivos distintos para jugar a nuestro antojo.</li>
<li><strong>Data.*</strong>: Son los archivos que utiliza ZODB como base de datos, donde se guardar&aacute; toda la informaci&oacute;n persistente.</li>
<li><strong>setup.py</strong>, <strong>MANIFEST.in</strong>: Archivos de Python necesarios por si queremos empaquetar nuestra aplicaci&oacute;n como bundle.</li>
<li><strong>src/</strong>: Directorio donde realmente ir&aacute; la l&oacute;gica de la aplicaci&oacute;n, modelos, vistas y templates (MVT).</li>
</ul>
<p>Bajando otro nivel (<em>src/</em> nuevamente) nos encontramos con un proyecto en blanco listo para empezar a picar c&oacute;digo. Decir que Pyramid (como Django y otros frameworks), siguen el patr&oacute;n MVT (model-view-template), donde el modelo se encarga de enlazar con los datos persistentes, la vista es la l&oacute;gica de la aplicaci&oacute;n (controlador) y los templates son las plantillas que moldean la aplicaci&oacute;n en lenguaje de marcado.</p>

```
$ cd src/
$ ls
__init__.py   models.py    static/       tests.py     views.py
templates/
```

<p>Creo que no hay mucho que explicar, la nomenclatura es m&aacute;s que expl&iacute;cita. Por defecto Pyramid nos ofrece llevar toda la l&oacute;gica de la aplicaci&oacute;n en un fichero (<strong>views.py</strong>) y el/los modelo(s) en otro (<strong>models.py</strong>); sin embargo como es muy listo y sabe que la aplicaci&oacute;n tendr&aacute; m&aacute;s de una plantilla nos ofrece el directorio <strong>templates/</strong> para que vayamos guardando ah&iacute; dentro nuestras obras de arte, lo mismo que con los archivos est&aacute;ticos (<strong>static/</strong>).</p>
<p>&iquest;Qu&eacute; pasar&iacute;a si queremos, por organizaci&oacute;n o limpieza, separar los modelos o la l&oacute;gica de la aplicaci&oacute;n en ficheros distintos?. El siguiente paso ser&aacute; modificar m&iacute;nimamente la forma de trabajar con Pyramid para que lea los modelos de un directorio <strong>models/</strong> y la l&oacute;gica del directorio <strong>views/</strong>.</p>
<p>La sencillez de Python nos permite crear un m&oacute;dulo importable en cualquier directorio que exista un __init__.py vac&iacute;o, de forma que si creamos el directorio models/ y dentro ubicamos este m&aacute;gico recurso, ya podr&iacute;amos importar nuestros modelos llam&aacute;ndolos con un "import models.my_model". Ya hemos hecho la mitad del trabajo hecho.</p>
<p>Para la otra mitad debemos convertir nuestro views.py en el __init__.py del directorio views/ de forma que cargue ese archivo por defecto antes de leer ninguna otra vista. Podemos reescribirlo de la siguiente forma:</p>

```
from pyramid.view import view_config
from .models import my_model
BASE_TMPL = 'src:templates/'

@view_config(context=my_model, renderer=BASE_TMPL + 'index.pt')
def my_view(request):
    return {'project': 'src'}
```

<p>Importamos un my_model.py dentro de models/ y aprovechamos para configurar templates/ como el directorio por defecto de las plantillas. Ahora est&aacute; todo un poco m&aacute;s ordenado y sin apenas esfuerzo.</p>
<p>El siguiente paso es intentar entender el <a href="http://docs.pylonsproject.org/projects/pyramid/en/1.0-branch/narr/urldispatch.html">URL Dispatcher</a> de Pyramid y, si nos vemos con fuerza y ganas, el <em>Traversal</em> de la <em>ZODB</em> y sus recursos <em>Root</em> y <em>Folder</em>.</p>
