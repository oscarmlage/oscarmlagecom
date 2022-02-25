---
title: "Pyramid: 5 - Views y Chameleon Templates"
date: 2013-11-03T19:00:00Z
draft: false
tags: [ "python", "code" ]
image: 
---

<p>Siguiendo con esta serie de art&iacute;culos sobre Pyramid y ZODB ya hemos escrito nuestros primeros modelos, as&iacute; que ya podemos interactuar con ellos de alguna forma: a trav&eacute;s de las <strong>views</strong>.</p>
<p>Las views o vistas en el patr&oacute;n MVT (no confundir con el mismo t&eacute;rmino "vista" del patr&oacute;n MVC), son las que albergan la l&oacute;gica de la aplicaci&oacute;n (lo equivalente al "controlador" en MVC). Por un lado conectan con el <em>model</em> y por otro lado con las <em>templates</em>&nbsp;para cerrar el proceso <em>request/response</em>.</p>
<p>Vamos a ver un ejemplo de view sencilla, que pasa un par de variables a template:</p>

```
# -*- coding: utf-8 -*-

from pyramid.view import view_config
from src.models.rootfolder import RootFolder

BASE_TMPL = 'src:templates/'

@view_config(context=RootFolder, renderer=BASE_TMPL + 'index.pt')
def homepage(context, request):

    return {'title': 'PyTwitter',
            'description': 'Our twitter clon, with Pyramid'}
```

<p>Definimos este m&eacute;todo en el contexto <em>RootFolder</em>, si recordamos, en la entrada anterior hemos configurado para que el punto de entrada a nuestra aplicaci&oacute;n sea <em>RootFolder</em>, de forma que esta view que hemos escrito ser&aacute; la principal de la aplicaci&oacute;n.</p>
<p>Adem&aacute;s estamos llamando al template <em>index.pt</em> y le hemos pasado dos variables, <em>title</em> y <em>description</em>. A continuaci&oacute;n creamos el fichero <em>index.pt</em> dentro de nuestro directorio <em>templates/</em> con el siguiente contenido:</p>

```
&lt;!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd"&gt;
&lt;html xmlns="http://www.w3.org/1999/xhtml" xml:lang="en" xmlns:tal="http://xml.zope.org/namespaces/tal"&gt;
&lt;head&gt;
  &lt;title&gt;${title}&lt;/title&gt;
  &lt;meta http-equiv="Content-Type" content="text/html;charset=UTF-8"/&gt;
  &lt;meta name="description" content="${description}" /&gt;&lt;/head&gt;
    &lt;body&gt;
      &lt;h1&gt;${title}&lt;/h1&gt;
      &lt;p&gt;${description}&lt;/p&gt;
    &lt;/body&gt;
&lt;/html&gt;
```

<p>Vemos como el uso de las variables que se pasan de view a template es muy sencillo. En Pyramid tenemos (por defecto) plantillas basadas en <a href="http://www.makotemplates.org/">Mako</a> o en <a href="https://chameleon.readthedocs.org/en/latest/">Chameleon</a>. A partir de la versi&oacute;n 1.5 debemos configurar en el m&eacute;todo <em>main()</em> cual de los dos vamos a usar.</p>

```
def main(global_config, **settings):
    config = Configurator(root_factory=root_factory, settings=settings)
    config.include('pyramid_chameleon')
    ...
```

<p>Personalmente prefiero <em>Mako</em>, que es el que usan sitios como <a href="http://www.python.org">Python.org</a> o <a href="http://www.reddit.com">reddit.com</a>, se parece m&aacute;s a <a href="http://jinja.pocoo.org/">Jiinja</a>, <a href="http://twig.sensiolabs.org">Twig</a> o a las plantillas de Django, que ya conocemos bastante. As&iacute; que por probar algo nuevo vamos a utilizar <strong>Chameleon</strong>.</p>
<p>Remito, para m&aacute;s informaci&oacute;n, a la <a href="https://chameleon.readthedocs.org/en/latest/">documentaci&oacute;n oficial de Chameleon</a>, pues escapa del objetivo de esta serie de art&iacute;culos hablar de un motor de plantillas u otro.</p>
<p>Hasta aqu&iacute; hemos creado ya lo que ser&aacute; la base de nuestra aplicaci&oacute;n: <em>model</em>, <em>view</em> y <em>template</em>. A partir de ahora intentaremos introducir m&aacute;s elementos como manejo de formularios, assets, tests y otros peque&ntilde;os trucos que har&aacute;n m&aacute;s sencillo tanto la programaci&oacute;n como el mantenimiento de nuestras aplicaciones.</p>
