---
title: "Apache: Alta carga de CPU"
date: 2011-08-09T10:27:52Z
draft: false
tags: [ "sysadmin" ]
image: 
---

<p>
	Llevo desde el fin de semana con la mosca detr&aacute;s de la oreja. Uno de los servidores que administro ha visto incrementada de forma inexperada su carga media de CPU sin motivo aparente. Donde el <em>load average</em> normal de 1 minuto variaba entre <em>0.40</em> y <em>0.80</em> de repente supon&iacute;a cargas tan elevadas como <em>60</em> o <em>100</em> unidades.</p>
<p>
	En esos momentos puntuales que llegaban a dejar la m&aacute;quina <em>zombie</em> el proceso que abarcaba un consumo de entre el <em>60%</em> y el <em>90%</em> de CPU era <em>apache2</em>. Intrigante que ni <em>error.log</em> ni <em>slow-queries.log</em> de MySQL (lo que normalmente suele ser cuello de botella) dieran ninguna pista.<!--more--></p>
<p>
	<em>Analytics</em> tampoco dec&iacute;a nada de un aumento considerable de visitas -m&aacute;s bien al contrario- y el resto de herramientas de monitorizaci&oacute;n parec&iacute;an c&oacute;mplices del problema (&iexcl;tener <em>tools</em> para &eacute;sto!).</p>
<p>
	En un alarde de desesperaci&oacute;n y viendo m&aacute;s o menos por d&oacute;nde pod&iacute;an venir los tiros -a trav&eacute;s del m&eacute;todo de <em>prueba y error</em>- localic&eacute; el <em>VirtualHost</em> que estaba dando problemas, un c&oacute;digo no auditado y afamado por su ausente optimizaci&oacute;n. Vamos que ya ten&iacute;a precedentes, aunque ninguno de esta &iacute;ndole.</p>
<p>
	Despu&eacute;s de desactivar distintas partes del dominio en m&aacute;s pruebas esperando focalizar el error en alg&uacute;n script concreto, la conclusi&oacute;n es que ten&iacute;a que ser algo que se inclu&iacute;a en todos los archivos, algo com&uacute;n a toda la web. As&iacute; que miramos los <em>includes</em> comunes (si, PHP) y llegamos a la conclusi&oacute;n de que era <strong>problema de sesiones</strong>, sin exculpar al programador.</p>
<p>
	Una de las primeras acciones de todos los scripts es definir el tiempo de sesi&oacute;n a <strong>3 d&iacute;as</strong>, definir el directorio donde se guardar&aacute;n los archivos de sesi&oacute;n y arrancar la sesi&oacute;n. Probablemente el programador no esperaba tener <em>60k</em> visitas y m&aacute;s de <em>200k</em> p&aacute;ginas vistas en los 3 d&iacute;as que dura cada sesi&oacute;n, pero est&aacute; claro que para <em>apache2</em> supon&iacute;a un problema el acceso de lectura/escritura a un mismo directorio con m&aacute;s de <em>90k</em> archivos.</p>
<p>
	La soluci&oacute;n inicial -a falta de m&aacute;s tiempo para cambiar el sistema de sesiones a <em>memcache</em>, <em>Redis</em> o cualquier otra soluci&oacute;n basada en RAM- era sencilla, reducir el tiempo de sesi&oacute;n, vaciar el directorio donde se guardan las sesiones, comprobar de nuevo el <em>load average</em> durante un intervalo representativo y una tarea programada que monitorice el contenido de ese directorio. Despu&eacute;s de todo la carga se ha vuelto a estabilizar entre <em>0.30</em> y <em>0.40</em>.</p>
<p>
	No s&eacute; si estoy para moralejas porque la soluci&oacute;n es temporal, pero como <a href="http://twitter.com/#!/r0sk/status/100688297642311681">lo prometido es deuda</a> me gustar&iacute;a terminar esta anotaci&oacute;n advirtiendo a todo el mundo sobre la fiabilidad del c&oacute;digo no auditado, el uso moderado de las sesiones y las endorfinas que libera uno cuando consigue resolver algo <em>&quot;as&iacute;&quot;</em>.</p>
