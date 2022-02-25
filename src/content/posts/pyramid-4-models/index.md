---
title: "Pyramid: 4 - Models"
date: 2013-11-02T19:00:00Z
draft: false
tags: [ "python", "code" ]
image: 
---

<p>Si hacemos memoria, hasta el momento ya sabemos <a href="../../../../pyramid-1-montando-entorno-y-proyecto-zodb.html">montar un proyecto y un entorno</a> m&iacute;nimo de trabajo, <a href="../../../../pyramid-2-entendiendo-la-estructura-del-proyecto.html">entendemos la estructura</a> del mismo, incluso nos hemos arriesgado a modificarla ligeramente seg&uacute;n nuestras necesidades; y sabemos a pies juntillas c&oacute;mo funciona el <a href="../../../../pyramid-3-url-dispatching-y-zodb-transversal.html">despachador de urls</a> de Traversal contra la ZODB.</p>
<p>Es el momento de crear nuestro primer modelo de datos. Los modelos indican la forma en la que vamos a guardar los datos de nuestra aplicaci&oacute;n en la <strong>ZODB</strong> as&iacute; que son una parte crucial en la que debemos tener especial cuidado.</p>
<p>La idea en la que vamos a trabajar para entender el modo de funcionamiento es algo sencillo, un trivial clon de Twitter en el que podamos actualizar nuestro estado y ver todos los estados anteriores que hemos escrito. Luego, si hay tiempo y ganas, lo iremos complicando poco a poco.</p>
<p>Pensando un poco en las propiedades de este modelo vamos a definirlo de forma que tenga <strong>id</strong>, <strong>twitt</strong> y <strong>published</strong>&nbsp;(fecha de publicaci&oacute;n). Para crearlo nos vamos al directorio <em>models/</em> y dentro de un nuevo fichero <em>twitt.py</em> agregamos el siguiente contenido:</p>

```
# -*- coding: utf-8 -*-

from uuid import uuid1
from persistent import Persistent

class Twitt(Persistent):

    def __init__(self, **kwargs):
        self.tid = uuid1()
        self.twitt = kwargs.get('twitt', None)
        self.published = kwargs.get('published', None)
        super(Twitt, self).__init__()
```

<p>Como vemos no se trata m&aacute;s que de una clase normal en Python que hereda de <strong>Persistent</strong> y tiene las propiedades que hemos descrito anteriormente. &iquest;Sencillo?, pues ya tenemos nuestro primer modelo. Adem&aacute;s y como trabajo extra vamos a intentar montar un modelo a mayores que sea el "root" de la aplicaci&oacute;n. Esto significa que cuando entremos a la aplicaci&oacute;n, por defecto, nos responda ese modelo.</p>
<p>Un modelo de ese tipo tiene que tener ciertas propiedades especiales, por eso, habitualmente para la petici&oacute;n principal de la aplicaci&oacute;n se utiliza un modelo de tipo <em>Folder</em> (<em>repoze.folder</em>). Como nosotros vamos a darle ciertas propiedades o m&eacute;todos especiales, en vez de usar Folder directamente, vamos a extender su funcionamiento. Creamos en <em>models/</em> nuestro <em>rootfolder.py</em> que actuar&aacute; como punto de entrada a la aplicaci&oacute;n:</p>

```
# -*- coding: utf-8 -*-

from repoze.folder import Folder
from src.models.twitts import Twitt

class RootFolder(Folder):

    def __init__(self):
        super(RootFolder, self).__init__()
```

<p><strong>Ojo</strong>: Para usar repoze.folder es posible que tengamos que instalarlo en nuestro entorno: <strong>pip install repoze.folder</strong></p>
<p>Una vez tenemos este modelo "especial" preparado debemos configurar el objeto <em>app_root</em> que tendremos dentro del m&eacute;todo <strong>app_maker()</strong>&nbsp;especific&aacute;ndole que RootFolder() ser&aacute; lo primero que vea la aplicaci&oacute;n. Yo tengo este m&eacute;todo en views/__init__.py, pero hay gente que trabajando con ZODB acostumbra a ponerlo en el models.py:</p>

```
def appmaker(zodb_root):
    if not 'app_root' in zodb_root:
        app_root = RootFolder()
        zodb_root['app_root'] = app_root
        import transaction
        transaction.commit()
    return zodb_root['app_root']
```

<p><strong>Tip</strong>: Una forma de comprobar que todo est&aacute; funcionando (no siempre pasa a la primera) es entrar v&iacute;a pshell y comprobar que realmente est&aacute; cargando el root que le hemos configurado. Si no lo hace puede ser que Data.fs est&eacute; corrupto porque se ha creado anteriormente y debamos eliminarlo/migrarlo. Para comprobarlo por <em>pshell</em> podemos hacer lo siguiente:</p>

```
$ pshell development.ini
&gt;&gt;&gt; root
&lt;src.models.rootfolder.RootFolder object None at 0x11037b758&gt;
```

<p>Hasta aqu&iacute; ya hemos creado nuestro primer modelo, un segundo modelo "especial" basado en repoze.folder para el root de nuestra aplicaci&oacute;n, y hemos empezado a usar la shell interactiva. En la siguiente entrada veremos c&oacute;mo escribir una view y una template con <a href="https://chameleon.readthedocs.org/en/latest/">Chameleon</a>.</p>
