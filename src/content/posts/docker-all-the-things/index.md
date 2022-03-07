---
title: "Docker all the things"
date: 2019-04-02T10:00:50Z
draft: false
tags: [  ]
image: docker-all-the-things.jpg
---

<p>Si un proyecto me ha llamado poderosamente la atenci&oacute;n - llegando incluso a ponerme muy pesado con el tema - en estos &uacute;ltimos a&ntilde;os en el contexto del mundo <em>sysadmin</em>, ese ha sido <a href="https://www.docker.com/">Docker</a>.</p>
<p>Me encanta. Eso es as&iacute; y aunque parece que &uacute;ltimamente en el <em>mundillo tech</em>&nbsp;abunda el hater aburrido, no veo motivos por los que tenga que esconderlo. Prefiero ser constructivo porque a fin de cuentas - para m&iacute; - el <em>Software Libre</em> versa b&aacute;sicamente sobre eso: libertad.</p>
<p>Podr&iacute;a hacer una lista de todo lo que me ha asombrado de esta <em>nueva</em>&nbsp;tecnolog&iacute;a, pero no he venido a eso. He venido a contaros brevemente c&oacute;mo <code>Docker</code> ha cambiado la forma de afrontar el entorno de desarrollo de mis &uacute;ltimos proyectos.</p>
<p>Este 2019 ha venido con m&uacute;ltiples sorpresas a nivel laboral, variopintos proyectos, cada uno con su stack y sus requisitos distintos. Diferentes lenguajes de programaci&oacute;n, diferentes maneras de guardar los datos pero todos ellos con un denominador com&uacute;n: una carretilla llena de problemas a la hora de montar el entorno de desarrollo.</p>
<p>Tengo proyectos con <code>MySQL-5.5</code>, con <code>MySQL-5.7</code>, <code>MariaDB-10.1.32</code>, <code>AuroraDB</code>, <code>Python-2.7</code>, <code>Python-3.6.5</code>, <code>Django</code>, <code>Flask</code>, <code>FalconFramework</code>, <code>Zappa</code>, <code>PHP-5.6</code>, <code>PHP-7.*</code>, <code>CodeIgniter</code>, <code>Laravel</code>, <code>Apache2.2</code>, <code>Apache2.4</code>, <code>Nginx</code>, <code>Jasper Reports</code>... Imaginaos por un momento que no exista ning&uacute;n tipo de soluci&oacute;n de virtualizaci&oacute;n barra containers barra whatever. &iexcl;La que se podr&iacute;a haber liado para montar todo esto!. El ejemplo m&aacute;s llamativo de los &uacute;ltimos d&iacute;as ha sido la dependencia de un <code>Jasper Server</code>, con todo su <code>JAVA</code>,&nbsp;<code>Tomcat</code>&nbsp;y dem&aacute;s dolores de cabeza... Y por otro lado tendr&iacute;amos algo tal que as&iacute;:</p>

```sh
$ docker pull bitnami/bitnami-docker-jasperreports
$ vim docker-compose.yml
... 
$ docker-compose up -d
```

<p>Con <code>Docker</code> los problemas (haberlos hailos) se simplifican a base de contenedores y lo que podr&iacute;an haber sido d&iacute;as (incluso semanas) de setup para empezar a trabajar, se han convertido en an&eacute;cdota.</p>
<p>As&iacute; que si sab&eacute;is de alguna tienda que venda camisetas con el t&iacute;pico <code>Docker all the things</code>... p&oacute;ngame, por lo menos, dos.</p>
