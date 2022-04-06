---
title: "Pair programming con tmux"
date: 2013-02-25T08:56:09Z
draft: false
tags: [  ]
image: 
---

<p>Con screen <a href="../../../../screen-el-nuevo-kibitz.html">era sencillo</a>&nbsp;crear un entorno para el pair programming, pero con <a href="http://tmux.sourceforge.net/">tmux</a> tambi&eacute;n es trivial. Me refiero a compartir una sesi&oacute;n de terminal para poder interactuar desde varias localizaciones distintas:</p>

```
tmux -S /tmp/pair
chmod 777 /tmp/pair
tmux -S /tmp/pair attach
```

<p>Se trata de crear una sesi&oacute;n especificando el destino, asignarle permisos (en este caso a todo el mundo) y conectarse a esa sesi&oacute;n con el otro usuario deseado. A partir de ah&iacute; ambos usuarios (el que ha creado la sesi&oacute;n y el que se ha conectado posteriormente) pueden interactuar simult&aacute;neamente.</p>
