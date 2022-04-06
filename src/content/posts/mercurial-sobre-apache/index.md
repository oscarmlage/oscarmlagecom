---
title: "Mercurial sobre Apache"
date: 2010-05-07T14:52:47Z
draft: false
tags: [  ]
image: 
---

<p>
	Mi predilecci&oacute;n por <a href="http://mercurial.selenic.com/">Mercurial</a> ha quedado patente en <a href="http://www.userlinux.net/mercurial-automatizando-al-maximo.html">alg&uacute;n</a> que <a href="http://www.userlinux.net/primeros-pasos-con-mercurial.html">otro</a> <a href="http://www.userlinux.net/mercurial-hook-on-push.html">post</a>, as&iacute; que una vez estamos conforme con nuestro servidor de versiones llega el momento de dar un paso m&aacute;s. Intentaremos configurar un interfaz web para mostrar el c&oacute;digo a todo el mundo (una especie de <a href="http://trac.edgewall.org/">Trac</a> solo para c&oacute;digo y adaptado a <em>Mercurial</em>).</p>
<p>
	El proceso es tan sencillo como crear otro <em>VirtualHost</em> en tu Apache con unas caracter&iacute;sticas un poco especiales porque en vez de tirar de archivos din&aacute;micos (.php, .asp&hellip;) vamos a tirar de un cgi en Python, as&iacute; que la configuraci&oacute;n ser&iacute;a algo as&iacute;:</p>
<p>
	<!--more--></p>

```
&lt;VirtualHost *:80&gt;
        ServerAdmin webmaster@localhost
        DocumentRoot /home/www/userlinux/sd/mercurial/cgi-bin/
        ServerName hg.userlinux.net
        ScriptAlias / /home/www/userlinux/sd/mercurial/cgi-bin/hgwebdir.cgi/
        &lt;Directory &quot;/home/www/userlinux/sd/mercurial/cgi-bin/&quot;&gt;
                SetHandler cgi-script
                AllowOverride None
                Options +ExecCGI -MultiViews +SymLinksIfOwnerMatch
                Order allow,deny
                Allow from all
        &lt;/Directory&gt;
        ErrorLog  /home/www/userlinux/sd/mercurial/logs/error.log
        CustomLog /home/www/userlinux/sd/mercurial/logs/access.log combined
&lt;/VirtualHost&gt;
```

<p>
	Fijaos que el archivo <em>hgwebdir.cgi</em> es el que lleva todo el trabajo, y &iquest;de d&oacute;nde lo sacamos?. Si tenemos instalado Mercurial (cosa que se da por supuesta), ya lo tenemos en el sistema, tan solo hemos de copiarlo a su nueva ubicaci&oacute;n y tan contentos:</p>
<p>
	&nbsp;</p>

```
# cp /usr/share/doc/mercurial/examples/hgwebdir.cgi /home/www/userlinux/sd/mercurial/cgi-bin/ # chmod a+x /srv/hg/cgi-bin/hgwebdir.cgi
```

<p>Ahora vamos a configurar alguna que otra preferencia del interfaz web, para ello creamos al mismo nivel del <em>cgi</em> un archivo de configuraci&oacute;n llamado <em>hgweb.config</em>:</p>

```
[paths]
rcms-plugin-prueba = /home/mercurial/rcms-plugin-prueba/
[web]
style = gitweb
allow_archive = bz2 gz zip
maxchanges = 200
```

<p>
	En &eacute;l -como veis- definimos los repositorios que vamos a mostrar (nombre = ruta) y varias opciones de la web como su theme, si permitimos o no descargar el contenido del repositorio&hellip;</p>
<p>
	Ojo con los permisos, tanto del archivo de configuraci&oacute;n como de los repositorios, a los que ha de acceder el usuario propietario de Apache para leer. Un reinicio de Apache bastar&iacute;a para tener el repositorio listo:</p>
<p style="text-align: center;">
	<img alt="" src="gallery/mercurial-apache.png" style="width: 400px; height: 385px;" /></p>
