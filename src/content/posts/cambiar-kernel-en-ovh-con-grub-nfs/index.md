---
title: "Cambiar kernel en ovh con grub (nfs)"
date: 2013-03-11T07:59:44Z
draft: false
tags: [  ]
image: 
---

<p>El principal problema que existe para montar un servidor NFS en&nbsp;<span class="search_hit">OVH</span>&nbsp;es que los kernels que habilitan por defecto (<span class="search_hit">OVH</span>&nbsp;+ grsec), adem&aacute;s de venir con grsec como su propia descripci&oacute;n indica, no tienen soporte para instalar m&oacute;dulos de&nbsp;<span class="search_hit">kernel</span>. De forma que al intentar instalar el m&oacute;dulo para NFS nos dar&aacute; un fant&aacute;stico y maravilloso error. Por defecto tenemos el siguiente escenario:</p>

```
# uname -a
Linux ns302399.ovh.net 3.2.13-grsec-xxxx-grs-ipv6-64 #1 SMP Thu Mar 29 09:48:59 UTC 2012 x86_64 GNU/Linux
```

<p>As&iacute; que no queda otra que reemplazar el&nbsp;<span class="search_hit">kernel</span>&nbsp;por uno que ya lleve el m&oacute;dulo de nfs&nbsp;<em>builtin. </em>Llegados a este punto tenemos varias opciones: instalar un&nbsp;<span class="search_hit">kernel</span>&nbsp;desde fuentes, modificar el ya instalado o bajarnos uno de los kernels &rdquo;<em>-filer</em>&rdquo; de&nbsp;<span class="search_hit">OVH</span>. Los kernels "<em>-filer</em>" est&aacute;n pensados para este tipo de servidores de ficheros. Por sencillez optaremos por esta tercera opci&oacute;n.</p>
<p>Primeramente descargamos las im&aacute;genes del nuevo kernel que vamos a instalar desde el ftp oficial de&nbsp;<span class="search_hit">OVH</span>&nbsp;(<a class="urlextern" title="ftp://ftp.ovh.net/made-in-ovh/bzImage/" rel="nofollow" href="ftp://ftp.ovh.net/made-in-ovh/bzImage/">ftp://ftp.<span class="search_hit">ovh</span>.net/made-in-<span class="search_hit">ovh</span>/bzImage/</a>):</p>

```
# cd /boot/
# wget ftp://ftp.ovh.net/made-in-ovh/bzImage/3.2.13/bzImage-3.2.13-xxxx-grs-ipv6-64-filer
# wget ftp://ftp.ovh.net/made-in-ovh/bzImage/3.2.13/System.map-3.2.13-xxxx-grs-ipv6-64-filer
# ls
bzImage-3.2.13-xxxx-grs-ipv6-64        System.map-3.2.13-xxxx-grs-ipv6-64
bzImage-3.2.13-xxxx-grs-ipv6-64-filer  System.map-3.2.13-xxxx-grs-ipv6-64-filer
grub
```

<p>Ahora tendremos que modificar <em>grub</em> para que al arrancar la m&aacute;quina lo haga con el nuevo&nbsp;<span class="search_hit">kernel</span>. Lo hacemos editando el fichero&nbsp;<em>/etc/grub.d/06_OVHkernel</em>, en la l&iacute;nea que antes hac&iacute;a referencia al&nbsp;<span class="search_hit">kernel&nbsp;</span><em>bzImage*64</em>&nbsp;ahora le agregamos un&nbsp;<em>-filer</em>&nbsp;y guardamos los cambios:</p>

```
OS="Debian GNU/Linux"
LINUX_ROOT_DEVICE=${GRUB_DEVICE}
linux=`ls -1 -t /boot/bzImage*64-filer 2&gt;/dev/null | head -n1`
```

<p>Una vez tenemos el nuevo&nbsp;<span class="search_hit">kernel</span>&nbsp;referenciado en el archivo de plantilla&nbsp;<em>06_OVHkernel</em>, ejecutamos el&nbsp;<em>grub-mkconfig</em>&nbsp;para hacer efectiva la configuraci&oacute;n, es recomendable hacer antes una copia de seguridad de la configuraci&oacute;n anterior (sabe m&aacute;s el diablo por viejo que por diablo):</p>

```
# cp /boot/grub/grub.cfg /boot/grub/grub.cfg.old
# grub-mkconfig &gt; /boot/grub/grub.cfg
```

<p>Finalmente toca reiniciar, cruzar los dedos y, si todo ha ido correctamente, comprobar que ya tenemos el nuevo&nbsp;<span class="search_hit">kernel</span>&nbsp;listo para hacer funcionar nuestro servidor NFS:</p>

```
# uname -a
Linux ns302399.ovh.net 3.2.13-grsec-xxxx-grs-ipv6-64-filer #1 SMP Thu Mar 29 11:26:19 UTC 2012 x86_64 GNU/Linux
```

<p>No es complicado, pero es enga&ntilde;oso pensar en montar un NFS sobre OVH como una trivial tarea de poco m&aacute;s de 5 minutos.&nbsp;</p>
