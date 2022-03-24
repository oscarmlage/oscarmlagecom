---
title: "Transferir VPS de un proyecto Public Cloud OVH a otro"
date: 2016-05-12T10:33:40Z
draft: false
tags: [ "sysadmin" ]
image: ovh-vps-public-cloud.jpg
---

<p>Tengo que decir que soy un enamorado de lo que yo llamo "<em>virtualizaci&oacute;n de &uacute;ltima generaci&oacute;n</em>". Como no soy ning&uacute;n doctor en terminolog&iacute;a - ni mucho menos - hablo de <a href="https://www.docker.com/">Docker</a>, <a href="https://openvz.org/Main_Page">OpenVZ</a>, <a href="https://www.proxmox.com/en/">Proxmox</a>, <a href="https://www.openstack.org/">OpenStack</a>... que habitualmente tienen herramientas e interfaces m&aacute;s que asequibles para resolver muchos de los problemas que, a nivel de sistemas, se presentan hoy d&iacute;a.</p>
<p>En este caso concreto ten&iacute;a un cliente al que administro una m&aacute;quina muy b&aacute;sica, una especie de <em>LAMP</em> para sus proyectos personales y de desarrollo web. Hasta hace nada ten&iacute;a esa m&aacute;quina en mi cuenta de <a href="https://www.ovh.es/">OVH</a> y, por varios motivos nos apetec&iacute;a transferirla a su cuenta de <em>OVH</em>, donde gestiona pagos, dominios y el resto de productos.</p>
<p>Para m&aacute;s datos la m&aacute;quina es un <em>VPS</em> de un proyecto <em>Public Cloud</em> en el que tengo otras instancias que nada tienen que ver con ella (<em>mea culpa</em>, deber&iacute;a haber tenido un proyecto por cliente, ojal&aacute; lo hubiera sabido cuando decid&iacute; probar el public cloud).</p>
<p>Despues de varias consultas telef&oacute;nicas al soporte t&eacute;cnico de <em>OVH Espa&ntilde;a</em> y a trav&eacute;s de la cuenta de Twitter (<a href="https://twitter.com/ovh_support_es">@ovh_support_es</a>) consegu&iacute; los enlaces a las correspondientes gu&iacute;as que juegan con este tipo de servicios, me hice un caf&eacute; y me li&eacute; la manta a la cabeza dispuesto a divertirme un rato. Y as&iacute; fue.</p>
<p><strong>Poner en orden los t&eacute;rminos</strong></p>
<p>Dentro de <em>OVH</em> tenemos la tecnolog&iacute;a <em>Cloud</em>, donde podemos crear proyectos (proyectos <em>Public Cloud</em>). Dentro de esos proyectos es donde albergamos los <em>VPS</em>. En mi caso concreto tengo un s&oacute;lo proyecto <em>Public Cloud</em> y como quiero transferir un <em>VPS</em> de ese proyecto a otro cliente tengo dos opciones:</p>
<ol>
<li>Crear otro proyecto public cloud distinto que ser&aacute; el destinatario.</li>
<li>Entrar en la cuenta del cliente y crear un proyecto public cloud que actuar&aacute; como destinatario.</li>
</ol>
<p>Cualquiera de las dos opciones es v&aacute;lida, a fin de cuentas vamos a trabajar "<em>de proyecto public cloud a proyecto public cloud</em>" independientemente de quien sea el propietario de los mismos.</p>
<p><em>OVH</em> trabaja con <a href="https://www.openstack.org">OpenStack</a> de fondo. <em>OpenStack</em> es un proyecto computaci&oacute;n en la nube para proporcionar una infraestructura como servicio (<em>IaaS</em>). Es <em>open source</em> y distribuido bajo la licencia <em>Apache</em>. Consiste en varios proyectos relacionados entre s&iacute; que controlan estanques de control de procesamiento, almacenamiento y recursos de red a trav&eacute;s de un centro de datos, todos administrados a trav&eacute;s de un panel de control que permite a los administradores controlar mientras potencia a sus usuarios proveyendo los recursos a trav&eacute;s de una interfaz web. (v&iacute;a <a href="https://es.wikipedia.org/wiki/OpenStack">Wikipedia</a>). Algunos de esos proyectos o componentes que nos interesa tener en mente:</p>
<ul>
<li><strong>Compute (Nova)</strong>: Es el controlador de la estructura cloud, escrito en <em>Python</em> (<em>yay!</em>), dise&ntilde;ado para escalar horizontalmente en hardware standard.</li>
<li><strong>Dashboard (Horizon)</strong>: Proporciona una interfaz gr&aacute;fica (<em>web</em>) para el acceso, provisi&oacute;n y automatizaci&oacute;n de los recursos basados en cloud. Tambi&eacute;n gestiona los recursos de seguridad y/o acceso a <em>API</em>.</li>
<li><strong>Servicio de Imagen (Glance)</strong>: &nbsp;Proporciona servicios de descubrimiento, de inscripci&oacute;n y de entrega de los discos y del servidor de im&aacute;genes.</li>
</ul>
<p>Creo que con eso es suficiente, usaremos <strong>Nova</strong> como cliente de <em>OpenStack</em>, <strong>Horizon</strong> para crear usuarios y acceso a <em>API</em> y <strong>Glance</strong> para trabajar con las im&aacute;genes.</p>
<p><strong>Crear acceso Horizon</strong></p>
<p>El primer paso es crear el acceso a <a href="https://horizon.cloud.ovh.net/">Horizon</a>, la interfaz web <em>dashboard</em> de <em>OpenStack</em>. Para ello vamos al manager de cada uno de nuestros proyectos cloud (recordad que estamos trabajando con dos proyectos, el que tiene el VPS y el que lo quiere), <code>Gesti&oacute;n y consumo del proyecto</code> &rarr; <code>OpenStack</code> y creamos un nuevo usuario. Insisto, haremos &eacute;sto por cada uno de los dos proyectos que tenemos.</p>
<p>Una vez creamos el usuario ya tenemos acceso a <em>Horizon</em> para obtener las claves <em>API</em>, que ser&aacute; el siguiente paso.</p>
<p><strong>Preparar el entorno para usar OpenStack</strong></p>
<p>Ahora debemos preparar una m&aacute;quina para que act&uacute;e como cliente de <em>OpenStack</em> e interactuar con los proyectos y las m&aacute;quinas virtuales que tenemos, como las gu&iacute;as referencia de <em>OVH</em> hablan de <em>Debian</em>, <em>CentOS</em> y <em>Windows</em> cogemos una m&aacute;quina <em>Debian</em>&nbsp;cualquiera que tengamos por ah&iacute; perdida - que tenga espacio suficiente y, si est&aacute; en ovh mejor que mejor por la transferencia de red - e instalamos los clientes de <code>nova</code> y <code>glance</code>:</p>

```
# apt-get update
# apt-get install python-glanceclient python-novaclient -y
# nova help
# glance help
```

<p>Ahora nos vamos a nuestro <em>Horizon</em> (hemos creado el usuario en el paso anterior) y creamos un acceso para poder trabajar contra &eacute;l con el cliente v&iacute;a api, en la barra lateral vemos <code>Accesos y seguridad</code> &rarr; <code>Acceso a la API</code> &rarr; <code>Descargar fichero RC de OpenStack</code>. Hacemos lo mismo con los dos proyectos que tenemos y conseguimos estos dos archivos:</p>
<ul>
<li><code>3653162169886180-openrc.sh</code> (mis proyecto p&uacute;blico, con varias instancias).</li>
<li><code>8844589259929700-openrc.sh</code> (el de mi cliente, con ninguna).</li>
</ul>
<p>Una vez tenemos los archivos los transferimos a la m&aacute;quina que actuar&aacute; como cliente <em>OpenStack</em> para poder trabajar contra ambos proyectos (he decidido guardar todo lo que tenga que ver con <em>OpenStack</em> en la ruta <code>/home/openstack</code>, cada uno es libre de hacerlo donde quiera):</p>

```
# scp 3653162169886180-openrc user@host:/home/openstack/micloud.sh
# scp 8844589259929700-openrc.sh user@host:/home/openstack/clientecloud.sh
```

<p>Ya tenemos los archivos de entorno, cargamos con el que queramos trabajar primero (la contrase&ntilde;a la tenemos en el panel donde anteriormente hemos creado el usuario para horizon), en este caso voy a cargar el de <code>micloud.sh</code> porque empezar&eacute; trabajando en ese proyecto:</p>

```
# source micloud.sh
Please enter your OpenStack Password:
```

<p><strong>Transferir snapshot de una cuenta a otra</strong></p>
<p>Te&oacute;ricamente ya tenemos todo listo para interactuar desde la m&aacute;quina que act&uacute;a de cliente con proyectos, instancias y las copias de seguridad. Recordamos que en el &uacute;ltimo comando del paso anterior hemos cargado las variables de entorno de <code>micloud.sh</code>, vamos a comprobar que todo es correcto, listamos las instancias:</p>

```
# nova list
+--------------------------------------+--------------------+--------+------------+-------------+
| ID                                   | Name               | Status | Task State | Power State |
+--------------------------------------+--------------------+--------+------------+-------------+
| ed677883-9c70-4b56-a21a-7dde34b4b149 | cliente.vps        | ACTIVE | -          | Running     |
| d12531a-d151-4d5a-9219-d9665094e4a64 | mi.vps             | ACTIVE | -          | Running     |
+--------------------------------------+--------------------+--------+------------+-------------+
```

<p>Tambi&eacute;n podemos usar <code>glance</code> para listar las im&aacute;genes que tenemos a nuestra disposici&oacute;n:</p>

```
# glance image-list
+--------------------------------------+----------------------------------+-------------+------------------+--------------+--------+
| ID                                   | Name                             | Disk Format | Container Format | Size         | Status |
+--------------------------------------+----------------------------------+-------------+------------------+--------------+--------+
| 92087692-94fd-4958-9784-61hy5a317590 | Centos 6                         | raw         | bare             | 8589934592   | active |
| c17f13b5-587f-4304-b550-eb939738289a | Centos 7                         | raw         | bare             | 2149580800   | active |
| 07dcb6a0-0a9a-47ed-bc36-546523ef65e3 | CoreOS stable 899.15.0           | raw         | bare             | 9116319744   | active |
| 73958794-ecf6-4e68-ab7f-1506eadab05b | Debian 7                         | raw         | bare             | 2149580800   | active |
| bdcb5042-3548-40d0-b06f-79551d3c4377 | Debian 8                         | raw         | bare             | 2149580800   | active |
| 7250cc02-ccc1-4a46-8361-a3d6d9111177 | Fedora 19                        | raw         | bare             | 2149580800   | active |
| 57b9722a-e6e8-4a55-8146-3e36a4777b78 | Fedora 20                        | raw         | bare             | 2149580800   | active |
| a5006914-ef04-41f3-8b90-6b99d0261a99 | FreeBSD 10.3 UFS                 | raw         | bare             | 5242880000   | active |
| 5a13b9a6-02f6-4f9f-bbb5-b131852848e8 | FreeBSD 10.3 ZFS                 | raw         | bare             | 5242880000   | active |
| 3bda2a66-5c24-4b1d-b850-83333b581674 | Ubuntu 12.04                     | raw         | bare             | 2149580800   | active |
| 9bfac38c-688f-4b63-bf3b-69155463d0e7 | Ubuntu 14.04                     | raw         | bare             | 10737418240  | active |
| 0c58e90e-168e-498a-a819-26792e4c769e | Ubuntu 15.10                     | qcow2       | bare             | 309854720    | active |
| cdc798d0-8688-45c1-bd7c-69bec0c2a0b2 | Ubuntu 16.04                     | raw         | bare             | 2361393152   | active |
| 7d983a54-d06b-488f-986c-ba0eaa980a51 | ubuntu-14.04-rescue              | raw         | bare             | 1073741824   | active |
| bb37762b-5a82-4c2b-b72b-91ea1016a941 | Windows-Server-2012-r2           | raw         | bare             | 107374182400 | active |
| 3e092033-906b-4ef1-bdfb-76155126f5ac | snapshot.mi.vps                  | qcow2       | bare             | 10324344832  | active |
+--------------------------------------+----------------------------------+-------------+------------------+--------------+--------+
```

<p>Bien, ahora que lo tenemos todo, nos disponemos a crear una snapshot del servidor que queremos transferir (cliente.vps), puesto que no tenemos ninguna creada:</p>

```
# nova image-create ed677883-9c70-4b56-a21a-7dde34b4b149 cliente_snapshot
# glance image-list | grep cliente
| 9c5d7d92-5b48-42be-8a49-f4f33b0235a4 | cliente_snapshot                   | raw         | bare             |              | queued |
```

<p>Esperamos que pase de <code>queued</code> a <code>active</code> para tenerla disponible antes de continuar, una vez disponible nos fijamos en el id y la descargamos a disco:</p>

```
# glance image-download --file cliente_snapshot.qcow 9c5d7d92-5b48-42be-8a49-f4f33b0235a4
```

<p>Una vez tenemos la imagen descargada cambiamos las variables de entorno para trabajar contra el otro proyecto. Nota: Podemos trabajar directamente contra otro proyecto de otra cuenta de cliente, no hace falta hacerlo en la misma cuenta de OVH. Vemos que se trata del otro proyecto porque, en este caso, no tenemos ninguna instancia ejecut&aacute;ndose:</p>

```
# source clientecloud.sh
Please enter your OpenStack Password:
```


```
# nova list
+----+------+--------+------------+-------------+----------+
| ID | Name | Status | Task State | Power State | Networks |
+----+------+--------+------------+-------------+----------+
+----+------+--------+------------+-------------+----------+
```

<p>El siguiente paso ser&aacute; enviar el snapshot que acabamos de crear a este proyecto y - finalmente - crear la tan esperada instancia vps partiendo de esa imagen:</p>

```
# glance image-create --name cliente_snapshot --disk-format qcow2 --container-format bare --file cliente_snapshot.qcow
+------------------+--------------------------------------+
| Property         | Value                                |
+------------------+--------------------------------------+
| checksum         | 411b684734ae0a07bda42bb0da5f9166     |
| container_format | bare                                 |
| created_at       | 2016-05-12T09:23:47                  |
| deleted          | False                                |
| deleted_at       | None                                 |
| disk_format      | qcow2                                |
| id               | ab64a222-433a-4d89-8285-ed4a726b3b24 |
| is_public        | False                                |
| min_disk         | 0                                    |
| min_ram          | 0                                    |
| name             | cliente_snapshot                     |
| owner            | 4g4456hs21s121d12e6ff8d8r8qa8928     |
| protected        | False                                |
| size             | 4524015616                           |
| status           | active                               |
| updated_at       | 2016-05-12T09:29:55                  |
| virtual_size     | None                                 |
+------------------+--------------------------------------+
```

<p>Para la creaci&oacute;n de la nueva instancia recomiendo hacerlo directamente a trav&eacute;s de la interfaz web de ovh (o en horizon), porque ser&aacute; m&aacute;s sencillo seleccionar las caracter&iacute;sticas adecuadas al VPS, de todas formas tambi&eacute;n se puede hacer a trav&eacute;s del cliente, el comando ser&iacute;a el siguiente:</p>

```
# nova boot --key_name SSHKEY --flavor 98c1e679-5f2c-4069-b4da-4a4f7179b758 --image ab64a222-433a-4d89-8285-ed4a726b3b24 clientserver_from_snap
```

<p>Y con &eacute;sto ya estar&iacute;a el VPS transferido de un proyecto public cloud de una cuenta de ovh a otro proyecto public cloud de otra cuenta distinta. Ha sido un poco lioso pero el esfuerzo y la diversi&oacute;n han merecido la pena.</p>
<p><strong>Bola Extra</strong></p>
<p>&iexcl;Las gu&iacute;as!, sin ellas no habr&iacute;a sido capaz de hacer asolutamente nada:</p>
<ul>
<li><a href="https://www.ovh.es/publiccloud/guides/g1773.crear_un_acceso_a_horizon">Gu&iacute;a para crear un acceso a Horizon</a></li>
<li><a href="https://www.ovh.es/publiccloud/guides/g1851.preparar_el_entorno_para_utilizar_la_api_de_openstack">Gu&iacute;a para preparar una m&aacute;quina como cliente OpenStack</a></li>
<li><a href="https://www.ovh.es/publiccloud/guides/g1774.acceso_y_seguridad_en_horizon">Gu&iacute;a para conseguir el fichero de variables en Horizon</a></li>
<li><a href="https://www.ovh.es/publiccloud/guides/g1852.cargar_las_variables_del_entorno_de_openstack">Gu&iacute;a para cargar el entorno de proyecto desde el cliente OpenStack</a></li>
<li><a href="https://www.ovh.es/publiccloud/guides/g1853.transferir_la_copia_de_seguridad_de_una_instancia_de_un_centro_de_datos_a_otro">Gu&iacute;a para transferir una copia de seguridad entre cuentas</a></li>
</ul>
