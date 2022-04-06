---
title: "Emacs go to line y reload sin salir de Emacs"
date: 2012-08-13T19:39:29Z
draft: false
tags: [  ]
image: 
---

<p>Se me hace muy c&oacute;modo tener un atajo de teclado para la t&iacute;pica orden de "<em>ir a la linea X</em>", as&iacute; que en vez de escribir todo el tochaco de comando que hace falta en Emacs para ello (<em>M-x goto-line&lt;RET&gt;#linea&lt;RET&gt;</em>) voy a hacer un binding. Abro el archivo .emacs y agrego lo siguiente:</p>

```
;; Goto-line
short-cut key(global-set-key "C-l" 'goto-line)
```

<p>Una vez guardado y para hacer efectiva la configuraci&oacute;n sin tener que salir de Emacs ejecuto lo siguiente:&nbsp;</p>

```
M-x load-file 
~/.emacs 
```

<p>Y ya tengo lo que quer&iacute;a, mi shortcut de "<em>ir a la linea X</em>" en <em>Ctrl + L</em>. Tip b&aacute;sico de Emacs patrocinado por chocolate blanco Mercadona y agua de mineralizaci&oacute;n muy d&eacute;bil Bezoya.</p>
