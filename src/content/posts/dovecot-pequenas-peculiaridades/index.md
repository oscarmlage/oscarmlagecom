---
title: "Dovecot, pequeñas peculiaridades"
date: 2010-09-07T09:20:13Z
draft: false
tags: [ "dovecot" ]
image: 
---

<p>
	Desde hace alg&uacute;n tiempo -y despu&eacute;s de haber lidiado con <a href="http://www.cyrusimap.org/">Cyrus</a> y <a href="http://www.courier-mta.org/">Courier</a>- he optado por <a href="http://dovecot.org/">Dovecot</a> como servidor <em>POP3</em> e <em>IMAP</em> para m&aacute;quinas en producci&oacute;n. Por varios motivos: la sencillez de configuraci&oacute;n, sigue los est&aacute;ndares, soporta <em>mbox</em> y <em>Maildir</em> y algo muy importante, tiene un backend de autentificaci&oacute;n <em>SMTP</em> compatible con <a href="http://www.postfix.org/">Postfix</a> (entre otros).</p>
<p>
	Sin duda el servicio de correo electr&oacute;nico es el menos agradecido y probablemente el m&aacute;s doloroso para el <em>sysadmin</em> pero el haber dado con esta combinaci&oacute;n de elementos me ha ahorrado un mont&oacute;n de problemas.</p>
<p>
	De todos modos en la &uacute;ltima instalaci&oacute;n que me ha tocado he encontrado un par de peculiaridades que me gustar&iacute;a documentar por si alguien se encuentra en la misma situaci&oacute;n.<!--more--></p>
<p>
	Habiendo instalado el mismo O.S., las mismas versiones de software y exactamente los mismos ficheros de configuraci&oacute;n a la hora de despachar correos me encuentro con un error inexperado en <em>mail.log</em>:</p>

```
... status=bounced (local configuration error)
```

<p>
	As&iacute; sin m&aacute;s descripci&oacute;n no puedo adivinar mucho as&iacute; que decido activar errores en <em>dovecot.conf</em> con la directiva de configuraci&oacute;n <em>log_path = /var/log/<span class="search_hit">dovecot</span>.log.</em> Ahora s&iacute; podemos sacar m&aacute;s informaci&oacute;n del <em>dovecot.log</em>:</p>

```
Fatal: postmaster_address  setting not given
```

<p>
	Esto ya es otra cosa, despu&eacute;s de un poco de <em>googling</em> corrijo el fallo agregando al fichero de configuraci&oacute;n una direcci&oacute;n de postmaster, sigue pareci&eacute;ndome raro porque <em>/etc/aliases</em> es el mismo que otras m&aacute;quinas y nunca hab&iacute;a notificado este problema antes pero bueno, es cuesti&oacute;n de agregar a <em>dovecot.conf</em>&nbsp; lo siguiente:</p>

```
protocol lda {
  postmaster_address = tu-postmaster@tu-dominio.com
}
```

<p>
	Reinicio el servicio y a funcionar... pero no por mucho tiempo puesto que al d&iacute;a siguiente me encuentro con el servicio parado, vuelvo a reiniciar y el proceso se vuelve a parar cada d&iacute;a, repitiendo la jugada. Volviendo a los logs -ese gran invento- veo que todos los d&iacute;as a eso de las <em>6:00am</em> suelta el siguiente mensaje:</p>

```
dovecot: 2010-09-05 05:59:53 Fatal: Time just moved backwards by 9 seconds. This might cause a lot of problems, so I&#39;ll just kill myself now. http://wiki.dovecot.org/TimeMovedBackwards
```

<p>
	Justo a esa hora tengo una tarea programada que sincroniza la hora del servidor con <em>rdate</em>. Al parecer Dovecot detecta que la hora ha cambiado y para no entrar en conflictos se hace el <em>harakiri.</em> Interesante, es algo que tienen documentado en <a href="http://wiki.dovecot.org/TimeMovedBackwards">su wiki</a> y te animan a que cambies <em>rdate/ntpdate</em> por <em>ntpd</em>, <em>clockspeed</em> o <em>chrony.</em></p>
<p>
	Nada del otro mundo pero s&iacute; me ha supuesto algo de tiempo saber el origen de los errores para poder subsanarlos as&iacute; que bueno, si al menos esta entrada ayuda a alguien o le ahorra alg&uacute;n dolor de cabeza me dar&eacute; por contento.</p>
