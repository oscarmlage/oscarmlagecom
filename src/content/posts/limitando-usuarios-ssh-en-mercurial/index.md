---
title: "Limitando usuarios ssh en Mercurial"
date: 2010-07-21T01:24:05Z
draft: false
tags: [  ]
image: 
---

<p>
	Si algo bueno tiene <a href="http://mercurial.selenic.com/">Mercurial</a> es que permite la autentificaci&oacute;n de usuarios a trav&eacute;s de SSH. Es muy sencillo agregar un nuevo usuario a un desarrollo/repositorio: <em>adduser</em> y con meterlo dentro del grupo correspondiente al desarrollo llegar&iacute;a. Pero &iquest;qu&eacute; ocurre si no queremos que ese usuario haga otra cosa que no sean comandos <em>hg</em>?.</p>
<p>
	Conociendo la existencia de <a href="http://www.selenic.com/repo/hg-stable/raw-file/tip/contrib/hg-ssh">hg-ssh</a> no ocurre demasiado, se trata de un script que hemos de referenciar en el <em>authorized_keys </em>del usuario que acabamos de crear de forma que todos los comandos entrantes pasen por este script. El script se encarga de parsear el comando que se pide en ejecuci&oacute;n: si es de la familia de <em>Mercurial</em> lo ejecuta, en cualquier otro caso mostrar&aacute; un error.</p>
<p>
	Ejemplo de <em>authorized_keys</em>:</p>

```
command=&quot;~/hg-ssh /home/repo1 /home/repo2&quot;,no-port-forwarding,no-X11-forwarding,no-agent-forwarding,no-pty ssh-dss AAAA...
```

<p>
	He optado por copiar el archivo hg-ssh en el directorio <em>home</em> del usuario, pero se podr&iacute;a referenciar directamente el que trae de ejemplo la instalaci&oacute;n de <em>Mercurial</em>.<!--more--></p>
<p>
	Nos encontramos con otro peque&ntilde;o inconveniente, si el script se tiene que referenciar en el archivo <em>authorized_keys</em> como he dicho antes, el usuario debe tener su clave p&uacute;blica DSA/RSA configurada en el servidor, as&iacute; que para estar seguros de que no se salta las restricciones impuestas por <em>hg-ssh</em>, deber&iacute;amos desactivar cualquier intento de acceso por contrase&ntilde;a:</p>

```
# nano /etc/ssh/sshd_config
Match Group desarrolladores
PasswordAuthentication no
```

<p>
	Reiniciamos <em>sshd</em> y todo deber&iacute;a funcionar de forma adecuada, el usuario podr&aacute; hacer <em>commits</em>, <em>push</em>, <em>pull</em>, <em>update</em>... desde la m&aacute;quina cuya clave <em>id_dsa.pub</em> hemos agregado en el servidor pero no podr&aacute; acceder a trav&eacute;s de ssh a ning&uacute;n otro comando, ni a shell; por lo tanto creo que se ha conseguido el principal objetivo.</p>
