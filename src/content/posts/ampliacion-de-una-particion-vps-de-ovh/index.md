---
title: "Ampliación de una partición, VPS de OVH"
date: 2016-04-20T08:21:14Z
draft: false
tags: [ "sysadmin" ]
image: vpsovh.jpg
---

<p>Gran invento los VPS de OVH, excepcional rendimiento a precio inigualable y ahora que han activado el servicio de <em>snapshots</em> no se puede pedir m&aacute;s por <em>3&euro;/mes</em>. Pero como no me llevo comisi&oacute;n, vamos al l&iacute;o t&eacute;cnico. Lo has empezado a usar, te ha molado y ahora el disco duro se te queda peque&ntilde;o (<em>10Gb</em> por defecto) as&iacute; que decides cambiar el tipo de servidor a otro con m&aacute;s disco... pero los cambios no se aplican autom&aacute;gicamente como en otras veces... as&iacute; que vamos a intentar arreglarlo.</p>
<p>Lo primero es entrar en la m&aacute;quina, ver el esquema de particionado y anotar alg&uacute;n que otro valor que nos va a ayudar m&aacute;s adelante:</p>

```
# mount
/dev/vda1 on / type ext4 (rw)
proc on /proc type proc (rw,noexec,nosuid,nodev)
sysfs on /sys type sysfs (rw,noexec,nosuid,nodev)
none on /sys/fs/cgroup type tmpfs (rw)
none on /sys/fs/fuse/connections type fusectl (rw)
none on /sys/kernel/debug type debugfs (rw)
none on /sys/kernel/security type securityfs (rw)
udev on /dev type devtmpfs (rw,mode=0755)
devpts on /dev/pts type devpts (rw,noexec,nosuid,gid=5,mode=0620)
tmpfs on /run type tmpfs (rw,noexec,nosuid,size=10%,mode=0755)
none on /run/lock type tmpfs (rw,noexec,nosuid,nodev,size=5242880)
none on /run/shm type tmpfs (rw,nosuid,nodev)
none on /run/user type tmpfs (rw,noexec,nosuid,nodev,size=104857600,mode=0755)
none on /sys/fs/pstore type pstore (rw)
systemd on /sys/fs/cgroup/systemd type cgroup (rw,noexec,nosuid,nodev,none,name=systemd)
```


```
# cat /etc/fstab
LABEL=cloudimg-rootfs  /   ext4  defaults  0 0
```


```
# df -h
Filesystem      Size  Used Avail Use% Mounted on
udev            1.9G  4.0K  1.9G   1% /dev
tmpfs           386M  892K  385M   1% /run
/dev/vda1       9.9G  7.5G  2.0G  80% /
none            4.0K     0  4.0K   0% /sys/fs/cgroup
none            5.0M     0  5.0M   0% /run/lock
none            1.9G   12K  1.9G   1% /run/shm
none            100M     0  100M   0% /run/user
```


```
# fdisk /dev/vda
Command (m for help): p

Disk /dev/vda: 21.5 GB, 21474836480 bytes
4 heads, 32 sectors/track, 327680 cylinders, total 41943040 sectors
Units = sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 512 bytes
I/O size (minimum/optimal): 512 bytes / 512 bytes
Disk identifier: 0x0005bac1

   Device Boot      Start         End      Blocks   Id  System
/dev/vda1   *        2048    20971519    10484736   83  Linux
```

<p>Anotamos los siguientes valores:</p>
<ul>
<li>Sector comienzo de la partici&oacute;n:&nbsp;<strong>2048</strong></li>
<li>Tama&ntilde;o usado de la partici&oacute;n:&nbsp;<strong>7.5Gb</strong></li>
<li>Tama&ntilde;o total del disco:&nbsp;<strong>21.5G</strong></li>
</ul>
<p>Como se puede ver, ya tenemos el nuevo disco (<em>20Gb</em>) en funcionamiento, pero las particiones todav&iacute;a est&aacute;n ancladas al anterior.&nbsp;Una vez anotados todos estos valores y teniendo a buen recaudo el resto de la informaci&oacute;n, tenemos que eliminar el sistema de particionado actual (<strong>d</strong>), comprobar que se ha eliminado -en este caso la &uacute;nica partici&oacute;n que hab&iacute;a- (<strong>p</strong>), recreamos el sistema de particiones (<strong>n</strong>) con el nuevo tama&ntilde;o y reescribimos la tabla de particiones (<strong>w</strong>). Todo esto usando <code>fdisk</code> en modo <code>live</code>, nada de <code>rescue</code> que por algo somos gallegos:</p>

```
# fdisk -u /dev/vda
Command (m for help): d
Selected partition 1

Command (m for help): p
   Device Boot      Start         End      Blocks   Id  System

Command (m for help): n
Partition type:
   p   primary (0 primary, 0 extended, 4 free)
   e   extended
Select (default p): p
Partition number (1-4, default 1):
Using default value 1
First sector (2048-41943039, default 2048):
Using default value 2048
Last sector, +sectors or +size{K,M,G} (2048-41943039, default 41943039):
Using default value 41943039

Command (m for help): p
   Device Boot      Start         End      Blocks   Id  System
/dev/vda1            2048    41943039    20970496   83  Linux

Command (m for help): w
The partition table has been altered!

Calling ioctl() to re-read partition table.

WARNING: Re-reading the partition table failed with error 16: Device or resource busy.
The kernel still uses the old table. The new table will be used at
the next reboot or after you run partprobe(8) or kpartx(8)
Syncing disks.
```

<p>Por defecto ya nos ha cogido los valores buenos al crear la nueva partici&oacute;n (<strong>2048</strong> como sector de inicio y <strong>41943039</strong> como sector de fin), si nos fijamos el sector final es pr&aacute;cticamente el doble del que ten&iacute;amos anteriormente (<em>20971519</em>), porque el nuevo disco es de <em>20Gb</em> en lugar de <em>10Gb</em>. Si no cogiera los valores por defecto tendr&iacute;amos que introducirlos manualmente.</p>
<p>Reiniciamos el vps y una vez haya reiniciado entramos de nuevo para redimensionar la nueva partici&oacute;n al tama&ntilde;o correspondiente:</p>

```
# df -h | grep vda1
/dev/vda1       9.9G  7.5G  2.0G  80% /
```


```
# resize2fs /dev/vda1
```


```
# df -h | grep vda1
/dev/vda1        20G  7.5G   12G  40% /
```

<p>Como se puede comprobar, ya tenemos la partici&oacute;n lista con su nuevo tama&ntilde;o. Al final ha sido m&aacute;s f&aacute;cil y menos traum&aacute;tico de lo esperado.</p>
