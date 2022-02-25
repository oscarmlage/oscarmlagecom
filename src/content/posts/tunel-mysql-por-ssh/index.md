---
title: "Túnel MySQL por SSH"
date: 2013-08-28T12:51:02Z
draft: false
tags: [  ]
image: 
---

<p>&iquest;Nunca os ha pasado aquello de que algo que conceptualmente no deber&iacute;a funcionar, lo pruebas m&aacute;s por desesperaci&oacute;n que por otra cosa y resulta que al final no resulta ser una mala idea?, pues este podr&iacute;a perfectamente ser uno de esos expedientes X.</p>
<p>Resulta que en el trabajo tenemos un problema de red, para cubrir toda la superficie tenemos que conectar dos routers con WDS (adem&aacute;s de la red cableada y el resto de elementos). El problema surge al acceder desde la parte inal&aacute;mbrica a un servidor MySQL (cableado), hay veces, momentos puntuales, en los que ese acceso se hace inpracticable por el tiempo de respuesta.</p>
<p>Tan s&oacute;lo pasa con ese servicio y en ese servidor. Incluso el resto de servicios de ese servidor responden razonablemente bien. Y desde la red cableada tampoco hemos notificado problemas. Conste que tampoco hemos hecho muchas m&aacute;s investigaciones puesto que se trata de un servidor dentro del entorno de desarrollo.</p>
<p>Y en un momento de lucidez se me ocurri&oacute; crear un t&uacute;nel por SSH pensando que no iba a funcionar, obviamente si el problema est&aacute; en la <em>capa 2</em>, por mucho que queramos pintar por encima...</p>

```
$ sudo ssh -fNg -L 6606:192.168.1.111:3306 root@192.168.1.111
```

<p>La sorpresa llega cuando llevo varios d&iacute;as conect&aacute;ndome "<em>en local</em>"&nbsp;y todo parece ir como la seda. No me ha vuelto a salir barba esperando una query. As&iacute; que siguiendo la principal regla inform&aacute;tica del "<em>si funciona no lo toques</em>", agrego un "<em>no quiero saber c&oacute;mo, pero que siga funcionando</em>".</p>
