---
title: "Apache + Squid + Nginx"
date: 2011-10-28T23:50:34Z
draft: false
tags: [  ]
image: 
---

<p>
	&iexcl;Menuda combinaci&oacute;n!. A decir verdad empec&eacute; jugando un poco con el maldito <em>slowloris</em> y al final acab&eacute; montando este batiburrillo de servidores, primero para <em>paliar</em> el efecto del dichoso gusano y segundo para preparar el servidor para la inminente nueva versi&oacute;n del blog - <em>que me gustar&iacute;a estrenar con el d&eacute;cimo aniversario de este humilde rinconcito</em> -.</p>
<p>
	En un esquema inicial anal&oacute;gico de esos que tantos nos gustan podemos ver la pirula (pido perd&oacute;n de antemano por la calidad de la foto):<!--more--></p>
<p style="text-align: center;">
	<img alt="" src="gallery/apache-squid-nginx.jpg" style="width: 300px; height: 305px;" /></p>
<p>
	Probablemente no sea necesario tener 3 servidores para este entorno, seguro que puedo poner a escuchar Nginx en el puerto 80 y redirigir todo el tr&aacute;fico din&aacute;mico al 81 a la vez que responde al tr&aacute;fico est&aacute;tico de ciertos subdominios (static*) de forma autom&aacute;tica y en la misma instancia. Pero ten&iacute;a la tarde libre y me apetec&iacute;a probar una configuraci&oacute;n <em>rara</em> con Squid.</p>
<p>
	<strong>Apache</strong></p>
<p>
	La configuraci&oacute;n de Apache es de lo m&aacute;s sencilla, lo &uacute;nico que he hecho ha sido ponerlo a escuchar en el puerto 81 por defecto, y a todos sus <em>Virtualhosts</em> tambi&eacute;n. No hab&iacute;a mucho m&aacute;s que tocar puesto que ya ten&iacute;a el <em>mod_php</em> y todas las dependencias instaladas.</p>
<p>
	<strong>Nginx</strong></p>
<p>
	Nunca hab&iacute;a jugado con &eacute;l y a primera vista me gust&oacute; la sencillez de sus archivos de configuraci&oacute;n. Tampoco ten&iacute;a que hacer gran cosa, ponerlo a escuchar en el puerto 82 y poco m&aacute;s puesto que s&oacute;lo servir&iacute;a contenido est&aacute;tico. Cre&eacute; los <em>Virtualhosts</em> que atender&iacute;an las peticiones est&aacute;ticas, los activ&eacute; v&iacute;a enlace simb&oacute;lico a <em>sites-enabled</em> y poco m&aacute;s. La configuraci&oacute;n de un <em>Virtualhost</em> cualquiera:</p>

```
server {
        listen   82;
        server_name  static.dominio.com;
        root /home/www/dominio/sd/static/;
        autoindex on;
}
```

<p>
	<strong>Squid</strong></p>
<p>
	Aqu&iacute; vino la diversi&oacute;n, &iquest;c&oacute;mo decirle a Squid que balanceara el tr&aacute;fico din&aacute;mico al puerto 81 y el est&aacute;tico al puerto 82?. Despu&eacute;s de leer la documentaci&oacute;n y hacer varias pruebas con <em>cache_peer</em>,&nbsp; y <em>cache_peer_domain</em> he llegado a la conclusi&oacute;n que la configuraci&oacute;n &quot;<em>buena</em>&quot; es la siguiente:</p>

```
cache_peer ip parent 81 0 no-query originserver name=server_1
cache_peer_domain server_1 dominio.com www.dominio.com
cache_peer ip parent 82 0 no-query originserver name=server_2
cache_peer_domain server_2 static.dominio.com
```

<p>
	Como se puede ver, habr&iacute;a que cambiar <em>ip</em> por la ip p&uacute;blica correspondiente y los <em>fqdn</em> por los reales.</p>
<p>
	<strong>Conclusi&oacute;n</strong></p>
<p>
	He pasado una tarde agradable en compa&ntilde;&iacute;a de mis amigos los servidores web. No, en serio, al final he conseguido reproducir el escenario que me hab&iacute;a propuesto, (a&uacute;n sabiendo que se podr&iacute;a mejorar), he frenado los sockets incompletos de Apache y he preparado el servidor para servir est&aacute;ticos de forma independiente del contenido din&aacute;mico (que en breve cambiar&aacute; de PHP a Python + Django).</p>
<p>
	No ha estado mal :).</p>
