---
title: "Restaurar InnoDB partiendo solamente de archivos"
date: 2013-12-05T14:42:23Z
draft: false
tags: [  ]
image: 
---

<p>El t&iacute;tulo suena muy a "<em>la he liado parda por favor &eacute;chame una mano</em>". En efecto. B&aacute;sicamente (y creo que lo repetir&eacute; varias veces a lo largo del art&iacute;culo para que queda claro), la mejor soluci&oacute;n para no llegar nunca a este punto es hacer dumps de vuestras bases de datos.</p>
<p>Olvidaos de <em>rsync</em>, <em>rdiff-backup</em> o cualquier soluci&oacute;n de backups que teng&aacute;is implementada contra ficheros, <em>InnoDB</em> no es <em>MyISAM</em>. Pero bueno, como he le&iacute;do en alguna de las referencias que pongo al final del art&iacute;culo, dejemos de llorar, nuestros datos todav&iacute;a deber&iacute;an estar ah&iacute;.</p>
<p><strong>Requisitos</strong></p>
<p>Para restaurar una base de datos <em>InnoDB</em> desde archivos (frm y binary logs), sin tener dumpeado completo, debemos tener lo siguiente:</p>
<ul>
<li>Backup del ibdata</li>
<li>Backup de los ib_logfiles</li>
<li>Backup del directorio de la base de datos (*.frm)</li>
</ul>
<p>Si tienes lo necesario podemos seguir con la intentona.</p>
<p><strong>Restaurando</strong></p>
<p>El primer paso es restaurar los ficheros del d&iacute;a adecuado, pongamos que queremos hacerlo de hace 10 d&iacute;as a trav&eacute;s de nuestro <em>rdiff-backup</em>:</p>

```
# mkdir ~/restaura/
# rdiff-backup -r 10D /backup/mysql/ibdata1 ~/restaura/ibdata1
# rdiff-backup -r 10D /backup/mysql/ib_logfile0 ~/restaura/ib_logfile0
# rdiff-backup -r 10D /backup/mysql/ib_logfile1 ~/restaura/ib_logfile1
# rdiff-backup -r 10D /backup/mysql/database ~/restaura/database/
```

<p>Los pasamos a la m&aacute;quina en la que vayamos a ejecutar el aislado de mysql:</p>

```
# scp -r ~/restaura/ oscar@Air.local:~/restaura-innodb/
```

<p><strong>Aislando MySQL</strong></p>
<p>Lo primero que hemos de tener claro es que para restaurar desde estos archivos en este escenario es que tenemos que aislar un <em>MySQL</em> solamente para &eacute;sto. Es decir, coger un <em>MySQL</em> que tengamos por ah&iacute; parado o sin mucho funcionamiento para ponerlo a funcionar s&oacute;lo con la base de datos que vamos a resturar.</p>
<p>&iquest;Por qu&eacute;?, porque los logs binarios (<em>ib_log</em> e <em>ibdata</em>) guardan informaci&oacute;n de todas las bases de datos del servidor, as&iacute; que no podemos sobreescribir los ya existentes, para ello podemos hacer una copia del directorio que actualmente est&aacute; funcionando en ese servidor que vamos a usar (para posteriormente dejarlo todo tal y como estaba).</p>
<p>En este caso concreto voy a utilizar el de mi Mac. Lo primero de todo es parar el servidor, Luego hago copia de seguridad y preparo el nuevo directorio <em>mysql5</em> para que el servidor lo detecte adecuadamente:</p>

```
$ sudo kill -9 `ps aux | grep mysql | awk '{print $2}'` ; ps aux | grep mysql
# cd /opt/local/var/db/
# mv mysql5 mysql5_backup
# mkdir mysql5
# cp mysql_backup/mysql-bin* mysql5
# cp -r mysql_backup/mysql/ mysql5
```

<p><strong>Agregando los datos restaurados</strong></p>
<p>Ahora que tenemos un entorno mysql m&iacute;nimo funcional (todav&iacute;a no hemos arrancado de nuevo el servidor), vamos a agregar la base de datos que hemos restaurado para poder acceder a los datos:</p>

```
# cp -r ~/restaura-innodb/ib_logfile* /opt/local/var/db/mysql5/
# cp -r ~/restaura-innodb/ibdata1 /opt/local/var/db/mysql5/
# cp -r ~/restaura-innodb/database /opt/local/var/db/mysql5/
```

<p>El nuevo directorio mysql5 tiene que quedar tal que as&iacute;:</p>

```
$ ls
Air.local.err       mysql-bin.000002    mysql-bin.000010    mysql-bin.000018
Air.local.pid       mysql-bin.000003    mysql-bin.000011    mysql-bin.000019
ib_logfile0     mysql-bin.000004    mysql-bin.000012    mysql-bin.000020
ib_logfile1     mysql-bin.000005    mysql-bin.000013    mysql-bin.000021
ibdata1         mysql-bin.000006    mysql-bin.000014    mysql-bin.index
database/        mysql-bin.000007    mysql-bin.000015
mysql/           mysql-bin.000008    mysql-bin.000016
mysql-bin.000001    mysql-bin.000009    mysql-bin.000017
```

<p>Comprobamos que todos los ficheros tienen permisos para el usuario <em>mysql</em> (<em>_mysql:_mysql</em> en el caso concreto del <em>MySQL</em> de <em>OSX</em> instalado v&iacute;a <a href="http://www.macports.org/">Macports</a>), y arrancamos de nuevo el servidor.&nbsp;Si todo ha ido bien y tenemos algo de suerte, deber&iacute;amos tener acceso a los datos de esa base de datos que se nos resist&iacute;a.</p>
<p><strong>Volverlo a dejar todo como estaba</strong></p>
<p>Este punto no es recomendable del todo. Como lecci&oacute;n aprendida no deber&iacute;amos dejar todo como estaba: deber&iacute;amos mejorar las copias de seguridad para que no vuelva a pasar esta pirula del <em>InnoDB</em>.</p>
<p>Sin embargo si nos referimos al servidor <em>MySQL</em> que aislamos, como hemos sido previsores, ser&aacute; sencillo dejarlo como estaba:</p>

```
# cd /opt/local/var/db/
# mv mysql5 mysql5_restauraciones
# mv mysql5_backup mysql5
```

<p>Y reiniciamos el servidor.</p>
<p><strong>Insistiendo en los dumps</strong></p>
<p>La mejor forma de hacer una copia de seguridad de una base de datos (en general, pero sobre todo si hablamos de <em>InnoDB</em>) es haciendo dumps completos, as&iacute; que es importante repetir hasta la saciedad que la pr&aacute;ctica aqu&iacute; descrita no es sana, ni recomendable, ni segura en absoluto. Ni para los <em>sysadmins</em>.</p>
<p><strong>Referencias</strong></p>
<ul>
<li><a href="http://egil.biz/how-to-recover-mysql-data-from-innodb/">http://getasysadmin.com/2012/08/recover-mysql-innodb-database-from-ibdata1-and-frm/</a></li>
<li><a href="http://egil.biz/how-to-recover-mysql-data-from-innodb/">http://egil.biz/how-to-recover-mysql-data-from-innodb/</a></li>
</ul>
