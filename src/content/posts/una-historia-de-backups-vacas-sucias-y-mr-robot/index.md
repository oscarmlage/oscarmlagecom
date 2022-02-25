---
title: "Una historia de backups, vacas sucias y Mr. Robot"
date: 2016-10-27T10:52:34Z
draft: false
tags: [ "sysadmin" ]
image: fsociety-mask-elliot.jpg
---

<p>Hace un rato, cuando he ido a abrir una de las aplicaciones de mi servidor me he dado cuenta que no funcionaba. La verdad es que no la uso demasiado a menudo pero ese <em>502</em> sonaba a muy raro porque nunca antes hab&iacute;a fallado.</p>
<p>Estas cosas pasan, me dije, as&iacute; que sin pensar mucho m&aacute;s he abierto el log de errores y lo he encontrado lleno de <em>proyecto/src/gunicorn.sock failed (2: No such file or directory) while connecting to upstream</em>, hmm, pero sin haber tocado nada entiendo yo que las rutas no se cambian solas &iquest;no?, bueno, comprobemos... <em>wt*</em>, &iexcl;&iexcl;<em>proyecto/src/</em>&nbsp;no existe!!, no hay nada, directorio en blanco, vac&iacute;o... &iquest;c&oacute;mo puede ser posible?.</p>
<p>Tranquilo, piensa en fr&iacute;o, t&uacute; haces copias de seguridad as&iacute; que no hay problema. Ponte la capucha Mr Robot e intenta pasar un buen rato con todo &eacute;sto.</p>
<p>Tienes raz&oacute;n, va, abriendo logs, el error ha aparecido el viernes 21 a eso de las 17:30 de la tarde. Haz memoria pero creo que ese viernes por la tarde... justo ese viernes y a esas horas... &iquest;no estabas probando los <em>PoC</em> de <em>dirtyc0w</em> en esa m&aacute;quina?. Ni con la capucha, pulsaciones subiendo... a ver, relaja, reinstalaci&oacute;n, restore de la &uacute;ltima copia y lecci&oacute;n aprendida &iquest;no?.</p>
<p>Ya pero... &iquest;uno de los <em>PoC</em> pudo haber borrado datos?, &iquest;c&oacute;mo es posible que nadie haya dicho nada?, demasiado raro. Espera, recuerdo que instantes despu&eacute;s de haberlo ejecutado la m&aacute;quina se qued&oacute; frita y tuve que reiniciar, &iquest;qu&eacute; dice el log del sistema?...</p>

```
Oct 21 17:31:08 me kernel: EXT4-fs (sda1): INFO: recovery required on readonly filesystem
Oct 21 17:31:08 me kernel: EXT4-fs (sda1): write access will be enabled during recovery
Oct 21 17:31:08 me kernel: EXT4-fs (sda1): orphan cleanup on readonly fs
Oct 21 17:31:08 me kernel: EXT4-fs (sda1): ext4_orphan_cleanup: deleting unreferenced inode 938198
Oct 21 17:31:08 me kernel: EXT4-fs (sda1): ext4_orphan_cleanup: deleting unreferenced inode 938126
Oct 21 17:31:08 me kernel: EXT4-fs (sda1): ext4_orphan_cleanup: deleting unreferenced inode 921486
Oct 21 17:31:08 me kernel: EXT4-fs (sda1): ext4_orphan_cleanup: deleting unreferenced inode 921485
```

<p>Va Mr Robot, deja de hacer el est&uacute;pido, qu&iacute;tate la capucha y p&aacute;sale un chequeo al disco m&aacute;s de vez en cuando, &iexcl;y ojo con los enlaces simb&oacute;licos!... No veas la ma&ntilde;anita que me est&aacute;s dando con estas tonter&iacute;as. Restaura el &uacute;ltimo backup y deja las gilipolleces para otro d&iacute;a.</p>
<p>Bueno, qu&eacute; digo &uacute;ltimo... &iexcl;el del 21!. Dime que guardas varias incrementales anda. &iexcl;Espera!, justo estos d&iacute;as has estado jugando con las copias que lo he visto... madre m&iacute;a... &iquest;por qu&eacute; no te quitas la capucha?, &iexcl;d&iacute; algo!, &iquest;&iexcl;por qu&eacute; te vas!?.</p>

```
# duply backup fetch proyecto/src/ /ruta/absoluta/al/proyecto/src/ 10D
```

<p>&iquest;Qu&eacute; ritmo te marca la pulserita pija esa que te has comprado?, &iquest;a cu&aacute;nto te bombea el coraz&oacute;n mientras <a href="http://duply.net/">duply</a> no acaba?. Menos mal que te hab&iacute;a dado por jugar con &eacute;l en su d&iacute;a. Esta vez te ha salvado el culo, adm&iacute;telo. Nunca est&aacute; de m&aacute;s tener otra copia a mayores, te lo dije.</p>
<p>Y as&iacute; es como, en una fr&iacute;a ma&ntilde;ana de octubre, he aprendido varias lecciones:</p>
<ul>
<li>No borrar incrementales a no ser por falta de espacio.</li>
<li>El tiempo perdido en mejorar la pol&iacute;tica de copias de seguridad es la mejor inversi&oacute;n, ever.</li>
<li><em>Duply</em> funciona cojonudamente, no quiero lanzar piedras contra el tejado de <em>rdiff</em>&nbsp;ni mucho menos -<em>de otras peores me ha sacado</em>-, pero <em>duply</em> hace magia blanca sin pedir permiso (comprimido, control de solapamiento, ejecuci&oacute;n secuencial s&oacute;lo en caso de no-fallo...).</li>
<li>La capacidad de almacenamiento est&aacute; demasiado barata como para racanear con ella.</li>
<li>Me encanta el modo capucha.</li>
</ul>
