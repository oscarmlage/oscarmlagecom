---
title: "Django + virtualenv + pip"
date: 2011-02-23T21:37:56Z
draft: false
tags: [ "django", "code" ]
image: 
---

<p>No lo ten&iacute;a claro, pero cuando entend&iacute; lo que supon&iacute;a y c&oacute;mo se trabajaba con <em>virtualenv</em> + <em>pip</em> me decid&iacute; a probarlo. Voy a <em>intentar</em> explicar como se utilizan estas herramientas de una forma gen&eacute;rica, para hacernos una idea de lo que significa y los casos en los que se pueden aplicar. A grandes rasgos:</p>
<ul>
<li class="level1">
<div class="li"><a href="http://djangoproject.org">Django</a>: Framework en <em>python</em>, creo que no necesita mucha m&aacute;s explicaci&oacute;n.</div>
</li>
<li class="level1">
<div class="li"><a href="http://pypi.python.org/pypi/virtualenv">Virtualenv</a>: Herramienta necesaria para crear un entorno virtual de <em>python</em>, con las versiones espec&iacute;ficas de los paquetes y/o dependencias que hagan falta para el proyecto.</div>
</li>
<li class="level1">
<div class="li"><a href="http://pypi.python.org/pypi/pip">Pip</a>: Gestor/Instalador de esos paquetes (similar a <em>easy_install</em>).</div>
</li>
</ul>
<p>Con estas herramientas intentaremos instalar un entorno virtual independiente para gestionar todas las dependencias de nuestro proyecto.<!--more--></p>
<p><strong>Instalaci&oacute;n del entorno virtualenv + pip</strong></p>
<p>Primero instalamos <em>pip</em>, que como hemos mencionado anteriormente, es el gestor de <em>paquetes</em> con el que vamos a instalar el resto de componentes. Tiramos de <em>easy_install</em> para instalarlo. Una vez instalado <em>pip</em>, instalamos tambi&eacute;n <em>virtualenv</em> y tendremos la base necesaria para seguir trabajando:</p>

```
$ sudo easy_install pip
$ pip install virtualenv
```

<p>Vamos a crear el entorno. Supongamos 2 directorios de trabajo, <em>src/</em> y <em>env/</em>. El directorio <em>src/</em> lo destinaremos al c&oacute;digo fuente (podemos tirar de repositorio o iniciar nuestro proyecto directamente ah&iacute;), en <em>env/</em> meteremos todos los paquetes y dependencias del proyecto, osea el entorno:</p>

```
$ mkdir -p foo/{env,src}
$ cd foo/
$ virtualenv --distribute --no-site-packages env/
New python executable in env/bin/python
Installing distribute........................................done.
```

<p>Ahora ya estamos preparados para instalar las aplicaciones/dependencias dentro del entorno de trabajo, lo haremos "desde fuera" con la opci&oacute;n <em>-E</em>:</p>

```
$ pip -E env/ install django
Downloading/unpacking django
  Downloading Django-1.2.4.tar.gz (6.4Mb): 6.4Mb downloaded
  Running setup.py egg_info for package django
[...]
Successfully installed django
Cleaning up...
```


```
$ sudo pip -E env/ install -e hg+https://bitbucket.org/ubernostrum/django-registration#egg=django-registration
$ sudo pip -E env/ install -e git+https://github.com/flashingpumpkin/django-socialregistration.git#egg=django-socialregistration
[instalamos el resto de dependencias]
```

<p>Vamos a comprobar que todo est&aacute; correcto, para ello utilizamos una herramienta que se llama <em>yolk</em>, capaz de listar todos los paquetes que tenemos instalados en el entorno, primero instalamos esa herramienta, luego entramos en el entorno y la ejecutamos (aunque podr&iacute;amos hacer lo mismo con la orden <em>pip freeze </em>como se muestra a continuaci&oacute;n):</p>

```
$ pip -E env/ install yolk
$ source env/bin/activate
(env)$ yolk -l
Django          - 1.2.4        - active 
Python          - 2.6.6        - active development (/usr/lib/python2.6/lib-dynload)
distribute      - 0.6.14       - active 
pip             - 0.8.1        - active 
wsgiref         - 0.1.2        - active development (/usr/lib/python2.6)
yolk            - 0.4.1        - active 
```

<p><strong>Nota</strong>: Como se puede ver en el ejemplo de arriba, para entrar en el entorno que hemos creado usamos el comando <em>source env/bin/activate</em>, para salir del mismo usamos el comando <em>deactivate</em>.</p>

```
(env)$ pip freeze
Django==1.2.4
distribute==0.6.14
wsgiref==0.1.2
yolk==0.4.1
```

<p>Ahora que tenemos m&aacute;s o menos clara la forma de actuar, vamos a hacer un <em>freeze</em> y guardar el archivo de requisitos dentro del repositorio para poder volver a reproducirlo en cualquier otra m&aacute;quina. Insisto, todos estos comandos pueden ejecutarse dentro del entorno (env) o fuera del mismo con la opci&oacute;n <em>-E entorno</em> de <em>pip</em>:</p>

```
(env)$ cd foo/src/
(env)$ pip freeze &gt; requirements.txt
(env)$ cat requirements.txt 
Django==1.2.4
South==0.7.3
distribute==0.6.14
wsgiref==0.1.2
yolk==0.4.1
(env)$ deactivate
$
```

<p>Y ya podemos empezar a programar el proyecto en dentro de <em>src/</em>.</p>
<p><strong>Replicando el entorno</strong></p>
<p>Suponemos ahora que estamos en otra m&aacute;quina -aunque en el ejemplo seguimos en la misma- donde queremos montar de nuevo todo el entorno. Primero creamos de nuevo la estructura anterior y luego hacemos un <em>pull</em> del repositorio la carpeta <em>src/</em> (como estamos en la misma m&aacute;quina lo simulando haciendo un <em>cp</em> :P):</p>

```
$ mkdir -p bar/{env,src}
$ cp -r foo/src/* bar/src
```

<p>Y ahora procuramos recrear el entorno, primero hacemos un entorno nuevo y vac&iacute;o en <em>bar/env/</em> y luego instalamos todo lo que hemos <em>logueado</em> en el <em>requirements.txt</em>:</p>

```
$ virtualenv --distribute --no-site-packages bar/env/
$ pip install -E bar/env/ -r bar/src/requirements.txt 
...
Successfully installed Django South ... yolk
Cleaning up...
```

<p>Comprobamos que todo est&eacute; correcto:</p>

```
$ source bar/env/bin/activate
(env)$ pip freeze
Django==1.2.4
distribute==0.6.14
wsgiref==0.1.2
yolk==0.4.1
(env)$ deactivate
$
```

<p><strong>Conclusi&oacute;n</strong></p>
<p><em>Virtualenv</em> y <em>pip</em> forman una combinaci&oacute;n excelente de elementos para hacer de un proyecto un <em>ente</em> sin dependencias frustradas y con todos los requisitos y versiones espec&iacute;ficas para garantizar el correcto funcionamiento del mismo. Ahora, por lo que me han contado, deber&iacute;a interesarme en <a href="http://fabfile.org/">Fabric</a> para automatizar los <em>deploys</em> y seguir haciendo magia.</p>
<p><strong>Referencias</strong></p>
<div class="level2 section_highlight">
<ul>
<li class="level1">
<div class="li"><a class="urlextern" title="http://www.saltycrane.com/blog/2009/05/notes-using-pip-and-virtualenv-django/" rel="nofollow" href="http://www.saltycrane.com/blog/2009/05/notes-using-pip-and-virtualenv-django/">Saltycrane</a></div>
</li>
<li class="level1">
<div class="li"><a class="urlextern" title="http://ryanwilliams.org/2009/Jun/09/deploying-django-sites-fabric-pip-and-virtualenv" rel="nofollow" href="http://ryanwilliams.org/2009/Jun/09/deploying-django-sites-fabric-pip-and-virtualenv">Ryanwilliams</a></div>
</li>
<li class="level1">
<div class="li"><a class="urlextern" title="http://techylinguist.com/how-to/python-and-django-dev-environment-virtualenv-and-pip" rel="nofollow" href="http://techylinguist.com/how-to/python-and-django-dev-environment-virtualenv-and-pip">Techylinguist.com</a></div>
</li>
<li class="level1">
<div class="li"><a class="urlextern" title="http://www.sebasmagri.com/2010/aug/12/pip-fabric-y-virtualenv-herramientas-para-el-desar/" rel="nofollow" href="http://www.sebasmagri.com/2010/aug/12/pip-fabric-y-virtualenv-herramientas-para-el-desar/">Sebasmagri 1</a></div>
</li>
<li class="level1">
<div class="li"><a class="urlextern" title="http://www.sebasmagri.com/2010/aug/14/esqueleto-para-proyectos-django-porque-la-forma-im/" rel="nofollow" href="http://www.sebasmagri.com/2010/aug/14/esqueleto-para-proyectos-django-porque-la-forma-im/">Sebasmagri 2</a></div>
</li>
</ul>
</div>
