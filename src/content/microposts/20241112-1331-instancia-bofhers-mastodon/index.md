---
title: instancia-bofhers-mastodon
date: 2024-11-12 11:31:51  +0200
draft: false
tags: micropost
---

En nuestra instancia [#BOFHers](https://mastodon.bofhers.es/tags/BOFHers) de [#Mastodon](https://mastodon.bofhers.es/tags/Mastodon) ([https://mastodon.bofhers.es](https://mastodon.bofhers.es/ "https://mastodon.bofhers.es/")) buscamos soluciones para mantener la actividad. Con 80 usuarios y 300+ GB de disco, en una VM de 8 GB RAM y 8 vCPUs, las actualizaciones son cada vez más complicadas, requiriendo casi el doble de espacio en disco, a pesar de nuestras tareas programadas para reducir cachés, assets y previews.

Agradecemos boost y/o cualquier consejo, sugerencia o aportación para que [#BOFHers](https://mastodon.bofhers.es/tags/BOFHers) siga siendo un espacio activo para todos.

Desde que agregamos el clean diario va algo mejor:

```sh
Filesystem Size Used Avail Use% Mounted on
/dev/xxx 492G 146G 323G 32% /xxx/mastodon
```  
Habría que quitarle 30Gb de los dumps diarios, que también están ahí  
```sh
load average: 2.61, 2.07, 2.32
```  
El cleanmedia diario:  
```sh
# tail -n 100 cleanmedia.log
3155/3155 |=========| Time: 00:02:08
Removed 3155 media attachments (approx. 3.71 GB)
1807/1807 |=========| Time: 00:01:09
Removed 1807 preview cards (approx. 343 MB)
```

No quiero entrar en comparaciones ni tampoco digo que no sea necesario todo esto, pero me parecen unos requisitos - de entrada - demasiado exagerados para interactuar con [#activitypub](https://mastodon.bofhers.es/tags/activitypub) .