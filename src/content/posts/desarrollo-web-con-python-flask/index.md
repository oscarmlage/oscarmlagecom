---
title: "Desarrollo web con Python: Flask"
date: 2012-02-15T18:47:34Z
draft: false
tags: [  ]
image: 
---

<p>
	Hace alg&uacute;n tiempo empec&eacute; a utilizar <a href="http://flask.pocoo.org/">Flask</a> para un proyecto personal. Flask es un mini framework de desarrollo en python que me ha convencido desde el principio por su sencillez. Intentar&eacute; exponer un peque&ntilde;o ejemplo para que os hag&aacute;is una idea de c&oacute;mo funciona.</p>
<p style="text-align: center; ">
	<img alt="" src="gallery/flask.png" style="width: 200px; height: 78px; " /></p>
<p>
	Y como <a href="http://toporojo.es/blog/">&Aacute;lex</a> ha escrito un <a href="http://toporojo.es/blog/2012/02/11/desarrollo-web-con-python-pylons/">completo tutorial sobre Pylons/Pyramid</a> con un ejemplo de aplicaci&oacute;n, no quer&iacute;a ser menos y explicar mis aventuras y desventuras con esta otra peque&ntilde;a joya de Python as&iacute; que all&aacute; vamos.</p>
<p>
	<!--more--></p>
<p>
	<strong>Instalaci&oacute;n&nbsp;</strong></p>
<p>
	Flask est&aacute; basado en &nbsp;<a href="http://werkzeug.pocoo.org/">Werkzeug</a>, una librer&iacute;a WSGI para Python. Si usamos <a href="http://pypi.python.org/pypi/pip">pip</a> y <a href="http://pypi.python.org/pypi/virtualenv">virtualenv</a> para crear el entorno virtual y comenzar a trabajar (situaci&oacute;n que <a href="http://www.userlinux.net/django-virtualenv-pip.html">ya he explicado en su d&iacute;a</a>) vemos que las dependencias de Flask son m&iacute;nimas:</p>

```
$ sudo easy_install pip
$ pip install virtualenv
$ mkdir -p flaskblog/{src,env}
$ cd flaskblog/
$ virtualenv --distribute --no-site-packages env/
$ pip -E env/ install flask
Downloading Flask-0.8.tar.gz
Downloading/unpacking Werkzeug&gt;=0.6.1 (from flask)
Downloading/unpacking Jinja2&gt;=2.4 (from flask)
Installing collected packages: flask, Jinja2, Werkzeug
$ pip -E env/ install flask-wtf
Downloading Flask-WTF-0.5.2.tar.gz
Downloading/unpacking WTForms (from flask-wtf)
Installing collected packages: flask-wtf, WTForms
$ pip -E env/ install flask-sqlalchemy
```

<p>
	Una &nbsp;vez creado el entorno virtual instalamos <em>Flask</em>, que depende de <em>Werkzeug</em> y de <em>Jinja2</em> como sistema de templates. Tambi&eacute;n instalamos <a href="http://packages.python.org/Flask-WTF/">Flask-WTF</a>&nbsp;para tener integraci&oacute;n con <em>WTForms</em> y hacer m&aacute;s sencillo el uso de formularios con funciones avanzadas (csrf, etc...). Y por &uacute;ltimo <em>Flask-SQLAlchemy </em>como complemento ORM para la base de datos.</p>
<p>
	<strong>Requirements y configuraci&oacute;n</strong></p>
<p>
	A continuaci&oacute;n vamos a entrar por primera vez en el entorno y crear el fichero de requisitos - yo normalmente le llamo <em>requirements.txt</em> - para poder regenerar el entorno virtual tantas veces y en tantas m&aacute;quinas como queramos:</p>

```
$ source env/bin/activate
(env)$ pip freeze &gt; src/requirements.txt
```

<p>
	Si luego queremos regenerar el entorno a partir del fichero lo haremos ejecutando el siguiente comando:</p>

```
$ pip install -E nuevo_env/ -r src/requirements.txt 
```

<p>
	Empezamos a montar el esqueleto de la aplicaci&oacute;n dentro del directorio <em>src/&nbsp;</em>creando un directorio para los templates, como queremos que nuestra aplicaci&oacute;n pueda tener varios templates a elegir, creamos tambi&eacute;n el &quot;por defecto&quot;:</p>

```
(env)$ mkdir -p templates/default/
```

<p>
	Ya que estamos empezando una aplicaci&oacute;n desde cero nos gustar&iacute;a tener un m&iacute;nimo de configuraci&oacute;n para la misma as&iacute; que vamos a usar un fichero para que guarde ciertos valores variables:</p>

```
(env)$ cat config.py
import os
DEBUG = True
_basedir = os.path.abspath(os.path.dirname(__file__))
DATA_PATH = os.path.join(_basedir, &#39;data&#39;)
DEFAULT_TPL = &#39;default&#39;
USERNAME = &#39;admin&#39;
PASSWORD = &#39;default&#39;
SECRET_KEY = &#39;devel secret key&#39;
URL = &#39;http://localhost:5000/&#39;
TITLE = &#39;OurApplication&#39;
VERSION = &#39;0.1&#39;
LANG = &#39;es&#39;
LANG_DIRECTION = &#39;ltr&#39;
YEAR = &#39;2012&#39;
del os
```

<p>
	Con la configuraci&oacute;n y el entorno listos, lo siguiente ser&aacute; empezar la aplicaci&oacute;n propiamente dicha, as&iacute; que manos a la obra.</p>
<p>
	<strong>Primeros pasos con Flask</strong></p>
<p>
	Vamos a reducir el grueso de la aplicaci&oacute;n a un solo archivo, le llamaremos <em>blog.py </em>- ya s&eacute; que no suena muy original pero imagino que ser&aacute; entendible -. Nuestro <em>blog.py</em> ser&aacute; el controlador principal, en &eacute;l incluiremos de forma excepcional el modelo, los formularios y los m&eacute;todos que se encargar&aacute;n de la l&oacute;gica y de llamar a las plantillas. Para empezar por el principio definimos el archivo como una aplicaci&oacute;n Flask y cargamos la configuraci&oacute;n que antes hemos creado:</p>

```
# -*- coding: utf-8 -*-
&quot;&quot;&quot;
OurApplication
~~~~~~~~~~~~~~
:copyright: (c) 2011 by Oscar M. Lage.
:license: BSD, see LICENSE for more details.
&quot;&quot;&quot;
import os
from flask import Flask, render_template, request, redirect
from werkzeug.routing import Rule
# Flask application and config
app = Flask(__name__)
app.config.from_object(&#39;config&#39;)
# Middleware to serve the static files
from werkzeug import SharedDataMiddleware
import os
app.wsgi_app = SharedDataMiddleware(app.wsgi_app, {
  &#39;/&#39;: os.path.join(os.path.dirname(__file__), &#39;templates&#39;, app.config[&#39;DEFAULT_TPL&#39;])
})
# Index
@app.route(&#39;/&#39;)
def index():
        return render_template(app.config[&#39;DEFAULT_TPL&#39;]+&#39;/index.html&#39;,
                            conf = app.config)
if __name__ == &#39;__main__&#39;:
    app.run()
```

<p>
	* Lo m&aacute;s raro son esas 3 lineas que hemos incluido en el archivo, una especie de <em>Middleware</em>&nbsp;para poder servir archivos est&aacute;ticos directamente desde el template seleccionado.</p>
<p>
	Una vez tenemos la aplicaci&oacute;n &quot;preparada&quot; hemos de crear la plantilla para que se pueda ejecutar sin errores, para ello dentro del directorio templates/default/ creamos el archivo index.html al que referenciamos en el controlador, el contenido del archivo podr&iacute;a ser algo as&iacute;:</p>

```
&lt;!DOCTYPE html PUBLIC &quot;-//W3C//DTD XHTML 1.0 Transitional//EN&quot;
 &quot;http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd&quot;&gt;
&lt;html xmlns=&quot;http://www.w3.org/1999/xhtml&quot; 
      xml:lang=&quot;{{ conf[&#39;LANG&#39;] }}&quot; lang=&quot;{{ conf[&#39;LANG&#39;] }}&quot; 
      dir=&quot;{{ conf[&#39;LANG_DIRECTION&#39;] }}&quot;&gt;
&lt;head&gt;
    &lt;title&gt;{% block title %}{% endblock %} [{{ conf[&#39;TITLE&#39;] }}]&lt;/title&gt;
&lt;/head&gt;
&lt;body&gt;
  &lt;div id=&quot;header&quot;&gt;
    &lt;h1&gt;{{ conf[&#39;TITLE&#39;] }}&lt;/h1&gt;
  &lt;/div&gt;
  &lt;div id=&quot;content&quot;&gt;
    Contenido
  &lt;/div&gt;
    &lt;div id=&quot;footer&quot;&gt;
        &lt;p&gt;{{ conf[&#39;TITLE&#39;] }} - {{ conf[&#39;YEAR&#39;] }}&lt;/p&gt;
    &lt;/div&gt;
&lt;/body&gt;
```

<p>
	Y ya podr&iacute;amos ejecutar por primera vez nuestra aplicaci&oacute;n desde consola para luego comprobar que todo est&aacute; correcto en navegador, as&iacute; que vamos a ello:</p>

```
$ python blog.py 
 * Running on http://127.0.0.1:5000/
 * Restarting with reloader
```

<p>
	Si abrimos el navegador en la url proporcionada (http://127.0.0.1:5000/) obtendremos una pantalla similar a la de la siguiente figura:</p>
<p style="text-align: center; ">
	<img alt="" src="gallery/flask001.png" style="width: 300px; height: 219px; " /></p>
<p>
	<strong>SQLAlchemy</strong></p>
<p>
	Una vez tenemos la base y el esqueleto de la aplicaci&oacute;n vamos a incorporar una base de datos para hacerla din&aacute;mica. Para ello usaremos <a href="http://packages.python.org/Flask-SQLAlchemy/">Flask-SQLAlchemy</a>. Para trabajar con SQLAlchemy agregamos una nueva variable de configuraci&oacute;n para la base de datos:</p>

```
SQLALCHEMY_DATABASE_URI = &#39;sqlite:///&#39;+ os.path.join(os.path.dirname(__file__), &#39;database.db&#39;)
```

<p>
	Y le decimos a <em>blog.py</em> el esquema que vamos a utilizar:</p>

```
from flaskext.sqlalchemy import SQLAlchemy
...
db = SQLAlchemy(app)
...
# Model
class Blog(db.Model):
    __tablename__ = &#39;Blog&#39;
    __mapper_args__ = dict(order_by=&quot;date desc&quot;)
    id = db.Column(db.Integer, primary_key=True)
    subject = db.Column(db.Unicode(255))
    author = db.Column(db.Unicode(255))
    date = db.Column(db.DateTime())
    content = db.Column(db.Text())
```

<p>
	Por &uacute;ltimo creamos la base de datos desde una consola python para poder usarla desde cualquier parte del controlador, tan simple como entrar a python desde el directorio src/ donde tenemos blog.py y ejecutar lo siguiente:</p>

```
&gt;&gt;&gt; from blog import db
&gt;&gt;&gt; db.create_all()
```

<p>
	Ya tenemos el archivo <em>database.db</em> preparado para cargarlo de datos mediante los m&eacute;todos correspondientes de <em>blog.py</em>, algunos ejemplos:</p>

```
@app.route(&#39;/add&#39;, methods=[&#39;GET&#39;,&#39;POST&#39;])
def add():
	if request.method == &#39;POST&#39;:
		post = Blog(request.form[&#39;subject&#39;], request.form[&#39;content&#39;])
		db.session.add(post)
		db.session.commit()
		return redirect(url_for(&#39;index&#39;))
	return render_template(app.config[&#39;DEFAULT_TPL&#39;]+&#39;/add.html&#39;,
			       conf = app.config)
```

<p>
	<strong>Flask-WTForms</strong></p>
<p>
	Una vez tenemos preparada la base de datos y el m&eacute;todo que agregar&aacute; nuevos posts tan solo falta crear el formulario que nos permitir&aacute; tal funcionalidad y crear el template, poco m&aacute;s de un par de lineas:</p>

```
# Create Form
class CreateForm(Form):
    subject = TextField(&#39;Subject&#39;, [validators.required()])
    content = TextAreaField(&#39;Content&#39;, [validators.required(), validators.Length(min=1)])
```

<p>
	Y el template con el formulario correspondiente ser&iacute;a el siguiente:</p>

```
&lt;form method=&quot;post&quot; action=&quot;&quot;&gt;
      &lt;dl&gt;
        {{ form.csrf }}
        {{ form.subject.label }} {{ form.subject(style=&quot;width:100%&quot;) }}
	{% for error in form.subject.errors %} {{ error }} {% endfor %}
		&lt;br /&gt;
	{{ form.content.label }} {{form.content(style=&quot;height:100px;width:100%&quot;) }}
	{% for error in form.content.errors %} {{ error }} {% endfor %}
      &lt;/dl&gt;
      &lt;p&gt;&lt;input type=&quot;submit&quot; value=&quot;submit&quot;&gt;
    &lt;/form&gt;
```

<p>
	Y ya tendr&iacute;amos un formulario para agregar posts totalmente funcional. Entre esta vista, la de un listado de posts y la del detalle de los mismos (ver repositorio m&aacute;s abajo), tendr&iacute;amos una peque&ntilde;a aplicaci&oacute;n con <em>Flask</em> + <em>SQLAlchemy</em> + <em>WTForms</em>.</p>
<p>
	<strong>Resumiendo</strong></p>
<p>
	Para finalizar, adem&aacute;s de las vistas que hemos dicho que faltaban, tendr&iacute;amos que dar un poco m&aacute;s de colorido a los templates (css, im&aacute;genes...), crear un <em>layout.html</em> que se pueda extender desde el resto de plantillas y poner un poco de orden en todo lo explicado, pero como no quiero extenderme mucho m&aacute;s en el art&iacute;culo, nada mejor que un peque&ntilde;o repositorio en el que se pueden ir viendo los avances:</p>
<ul>
	<li>
		<a href="https://bitbucket.org/r0sk/flaskblog">Flaskblog (bitbucket)</a></li>
</ul>
<p>
	Espero que con este peque&ntilde;o y humilde ejemplo haya quedado m&aacute;s o menos claro el uso de este framework. Siempre es divertido probar cosas nuevas y cuando se trata de piezas tan sencillas y bien documentadas como las tratadas el placer es doble.</p>
