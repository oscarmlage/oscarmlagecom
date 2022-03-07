---
title: "Por mí y por todos mis compañeros"
date: 2019-01-13T10:58:11Z
draft: false
tags: [ "sysadmin" ]
image: IMG_20181109_090634_Bokeh.jpg
---

<p>Algo tan sencillo y aparentemente inocuo como un punto de montaje puede convertirse en una letal arma de destrucci&oacute;n de paciencia/tiempo.</p>
<p>Hay una especie de regla no escrita que sigo desde hace bastante tiempo en las instalaciones de los servidores pertenecientes a <em>OVH</em>: centralizar todos los datos de servicios en <code>/home</code>. Los <em>DocumentRoot</em> de apache/nginx van en <code>/home/www</code>, las bases de datos en <code>/home/mysql</code>, etc...</p>
<p>&iquest;Por qu&eacute;?, porque hace alg&uacute;n tiempo, en las instalaciones por defecto, la partici&oacute;n con mayor tama&ntilde;o era esa, as&iacute; que lo ped&iacute;a a gritos. Y hasta el momento no he cambiado porque todos los scripts de deploy, provisionamiento y dem&aacute;s est&aacute;n prepardos para ello. Vagancia.</p>
<p>Estos &uacute;ltimos d&iacute;as he tenido un problema que nunca me hab&iacute;a pasado antes con un servidor. En este caso, inicialmente no hab&iacute;a separaci&oacute;n de particiones y todo estaba montado en <code>/</code>, pero ante la falta de espacio se contrat&oacute; un disco a mayores con la idea de montarlo en <code>/home</code> de forma separada y mover el grueso de la informaci&oacute;n ah&iacute; dentro:</p>

```sh
$ mkfs.ext4 /dev/sdb1
$ mkdir /home2
$ mount /dev/sdb1 /home2
$ rsync -az /home/ /home2/
$ umount /home2
$ mount /dev/sdb1 /home
```

<p>Creamos la partici&oacute;n del nuevo dispositivo, la montamos como <code>/home2</code>, sincronizamos datos, desmontamos <code>/home2</code> y la montamos de nuevo como <code>/home</code>. Hasta aqu&iacute; todo correcto &iquest;no?. Pues no. Si el antiguo <code>/home</code> no se borra queda ah&iacute; consumiendo espacio y - lo m&aacute;s gracioso - como posteriormente montamos <code>/dev/sdb1</code> como <code>/home</code>, queda oculto.</p>
<p>As&iacute; que, cuando vayas a mirar por qu&eacute; sigue qued&aacute;ndose <code>/</code> sin espacio habiendo descargado todo "<em>lo gordo</em>" te podr&aacute; pasar como a m&iacute; y perder 3 horas de research con <code>du</code>, <code>df</code>, scripts en <code>bash</code>, <code>lsof</code> y dem&aacute;s. O quiz&aacute;s seas un poco m&aacute;s avispado y te des cuenta. Soluci&oacute;n:</p>

```sh
$ mv /home /home_old
$ mount /dev/sdb1 /home
$ echo "Comprueba que todo funciona, varias veces, y posteriormente:"
$ rm -rf /home_old
```

<p>Moraleja: nunca juegues al escondite con tus sistemas. Y menos si eres t&uacute; el que pone unas normas que luego no recuerdas.</p>
