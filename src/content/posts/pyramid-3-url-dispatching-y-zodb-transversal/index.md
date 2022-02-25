---
title: "Pyramid: 3 - URL Dispatching y ZODB Traversal"
date: 2013-11-01T19:00:00Z
draft: false
tags: [ "python", "code" ]
image: 
---

<p>En cap&iacute;tulos anteriores hemos visto c&oacute;mo <a href="../../../../pyramid-1-montando-entorno-y-proyecto-zodb.html">montar el entorno</a> y c&oacute;mo <a href="../../../../pyramid-2-entendiendo-la-estructura-del-proyecto.html">entenderlo y adaptarlo m&iacute;nimamente</a> a nuestras necesidades organizativas. Hoy trataremos de abordar el tema del despachador de URL's... &iexcl;qu&eacute; mal suena en castellano!.</p>
<p>El "URL Dispatching" nos ofrece una manera lo m&aacute;s simple posible de mapear <em>URLs</em> a <em>views</em>. En otras palabras, dada una petici&oacute;n a nuestra aplicaci&oacute;n, debemos saber qu&eacute; parte de la l&oacute;gica tiene que aplicarse para devolver una respuesta.</p>
<p>Para configurar dicho mapeado, en Pyramid se usa el m&eacute;todo&nbsp;<a href="http://docs.pylonsproject.org/projects/pyramid/en/1.0-branch/api/config.html#pyramid.config.Configurator.add_route">pyramid.config.Configurator.add_route()</a>, por ejemplo:</p>

```
config.add_route('myroute', '/helloworld/',
                 view='myproject.views.helloworld')
```

<p>Le decimos al framework que si hay una petici&oacute;n del tipo <em>http://aplicacion/helloword/</em> vaya a ejecutar el m&eacute;todo <em>helloworld()</em> dentro de <em>myproject.views</em>.</p>
<p>En la documentaci&oacute;n oficial nos recomiendan configurar las rutas de forma centralizada en el __init__.py del proyecto, de forma que por un lado tengamos todas las rutas definidas con <em>add_route</em> y por otro lado, en las vistas, conectemos cada uno de nuestros m&eacute;todos con una de esas rutas:</p>

```
config.add_route('myroute', '/prefix/{one}/{two}')
config.add_view(myview, route_name='myroute')
```

<p>Sin embargo cuando trabajamos con <a href="http://docs.pylonsproject.org/projects/pyramid/en/latest/tutorials/wiki/basiclayout.html">ZODB y Traversal</a> la situaci&oacute;n da un giro bastante radical y de poco nos sirve todo lo visto. Si recordamos, a la hora de generar el proyecto le hemos dicho que &iacute;bamos a utilizar ZODB, por lo que el __ini__.py generado es ligeramente distinto a la de un proyecto Pyramid normal.</p>

```
def main(global_config, **settings):
    config = Configurator(root_factory=root_factory, settings=settings)
    config.include('pyramid_chameleon')
    config.add_static_view('static', 'static', cache_max_age=3600)
    config.scan()
    return config.make_wsgi_app()
```

<p><strong>La primera clave</strong> que debemos tener clara es que, por defecto, las rutas no se definen con un add_route, sino que van directamente en el decorator <em>@view_config</em> de cada vista. Y el encargado de recopilarlas y buscarlas por toda la aplicaci&oacute;n es el <strong>config.scan()</strong>. De forma que si definimos una view con el <em>name='fulanito'</em>, eso se convertir&aacute; en una url <em>/fulanito/</em> dentro de su contexto:</p>

```
@view_config(name='fulanito', context=myproject.models.MyModel, renderer='templates/fulanito.pt')
```

<p>Por lo que, poco a poco, seg&uacute;n vayamos construyendo la l&oacute;gica de nuestra aplicaci&oacute;n en view, iremos elaborando tambi&eacute;n el mapeado de rutas de la misma.</p>
<p><strong>La segunda clave</strong> para entender el sistema enrutado es el <strong>context y Traversal</strong>. Y si es lioso de entender, imaginaos de explicar, voy a intentarlo como ejercicio de que "yo tambi&eacute;n lo he entendido", aunque lo mejor es pasarse por la <a href="http://developer.plone.org/serving/traversing.html">documentaci&oacute;n oficial</a>.</p>
<p>En un sistema gestor de base de datos orientado a objetos como <em>ZODB</em>, cada objeto tiene un <em>path</em> dependiendo de su ubicaci&oacute;n, <a href="http://developer.plone.org/glossary.html#term-traversal">Traversal</a> es el m&eacute;todo encargado de manejar y llegar hasta un objeto guardado en la ZODB a partir de su <em>path</em>.</p>
<p>Para poder verlo y entenderlo adecuadamente vamos a compararlo con un almac&eacute;n robotizado. Todos los distintos tipos de productos est&aacute;n guardados en su estanter&iacute;a y en su ubicaci&oacute;n correspondiente dentro del gran almac&eacute;n (objetos en la ZODB). Cuando en el control del almac&eacute;n entra un nuevo pedido (request, GET, POST...) hay un brazo robotizado que, dependiendo del pedido, se encarga de ir a la estanter&iacute;a adecuada y recoger el paquete correspondiente para servirlo, pues justo eso es lo que hace Traversal.</p>
<p>&iquest;Y d&oacute;nde definimos los tipos de objetos que vamos a utilizar en nuestros modelos?, la propia pregunta se responde: en los modelos, bien usando por defecto los disponibles,&nbsp;extendiendo sus funcionalidades o creando nuestros propios tipos de objeto. Guardar objetos en la ZODB de forma persistente es sencillo, existen <a href="http://www.zodb.org/en/latest/documentation/guide/transactions.html">transacciones</a>, <a href="http://www.zodb.org/en/latest/documentation/guide/prog-zodb.html#writing-a-persistent-class">clases</a> propias para ello, etc... nosotros nos basaremos en objetos del tipo <a href="http://docs.repoze.org/folder/">Folder</a> para extender su funcionalidad (en alg&uacute;n caso) y trabajar con ellos sin tener que empezar de cero.</p>
<p>Lo bueno de usar una base de datos orientada a objetos es que se adapta perfectamente a los tipos de datos con los que tengas pensado trabajar en tu aplicaci&oacute;n. Aunque tambi&eacute;n tiene cosas menos buenas, como por ejemplo tener que escribir scripts de migraci&oacute;n cada vez que cambies o agregues propiedades a tus objetos.</p>
<p>En la siguiente entrada intentaremos escribir nuestro primer modelo de datos funcional.</p>
