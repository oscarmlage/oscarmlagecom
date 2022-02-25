---
title: "Moviendo entornos virtuales virtualenv"
date: 2016-02-11T21:05:00Z
draft: false
tags: [ "python", "code" ]
image: 
---

<p>Una de las ventajas de los entornos virtuales es que los podemos recrear en cualquier m&aacute;quina con muy poco coste de tiempo, siempre que tengamos los requisitos m&aacute;s o menos documentados es tan sencillo como lanzar un comando para volver a construirlo &iquest;o no?. La teor&iacute;a es fant&aacute;stica pero en la pr&aacute;ctica siempre surgen - como no - problemas <em>murphianos</em>.</p>
<p>En base a una necesidad muy concreta me ha tocado mover <a href="https://github.com/pypa/virtualenv">virtualenvs</a> de ubicaci&oacute;n para reorganizar tanto el c&oacute;digo como los entornos en el maltrecho disco duro de mi port&aacute;til, y por eso he creado esta peque&ntilde;a muestra de script "<em>quick and dirty</em>" (<code>movenv</code>):</p>

```
#!/bin/bash

VIRTUALENV=`which virtualenv`
OLD=$1 # /absolute/path/to/old/env
NEW=$2 # /absolute/path/to/new/env

${VIRTUALENV} --relocatable ${OLD}
mv ${OLD} ${NEW}
${SED} -i -e "s#${OLD}#${NEW}#g" ${NEW}/bin/activate
```

<p>Ahora s&oacute;lo tenemos que ejecutarlo <code>movenv /path/old/env /path/to/new/env</code> y comprobar que todo sigue en funcionando en la nueva ubicaci&oacute;n. Ni que decir tiene que el script es una peque&ntilde;a prueba de concepto, habr&iacute;a que a&ntilde;adir validaci&oacute;n de existencia de directorios y algunas otras cosas para que fuera plenamente funcional.</p>
