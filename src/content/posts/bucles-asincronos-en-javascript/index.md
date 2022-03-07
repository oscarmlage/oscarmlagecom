---
title: "Bucles asíncronos en javascript"
date: 2019-06-20T17:26:24Z
draft: false
tags: [  ]
image: vuedotjs.jpg
---

<p>Muchas veces hay ciertas situaciones que hacen que me plantee si esto de programar es para m&iacute;, sobre todo cuando, despues de un <em>debug</em> largo y tedioso, te das cuenta que lo que pasaba o era una chorrada o fallo m&iacute;o por no conocer los internals del lenguaje correspondiente.</p>
<p>Quiero pensar que en mayor o menor medida le pasa alguna vez a todo el mundo.</p>
<p>En este caso he perdido las yemas de los dedos de tanto <code>console.log</code> que creo que podr&iacute;a ir a por el r&eacute;cord Guinness. En la l&oacute;gica de un compoenente <a href="https://vuejs.org/">Vue</a> al que me ha tocado meterle mano hay una situaci&oacute;n un tanto graciosa. Tengo que hacer varias llamadas <em>AJAX</em> en el m&eacute;todo <code>mounted()</code> para recoger datos de un <em>API</em>.</p>
<p>Nada fuera de lo com&uacute;n. Digamos que tengo un id de competici&oacute;n y tengo que llamar a un endpoint en el que que, dado ese id, me devuelve el nombre de la competici&oacute;n. Una vez conozco el nombre, tengo que llamar a otro endpoint que me devuelva un listado de rondas que tiene esa competici&oacute;n. Y por &uacute;ltimo, dado ese listado de rondas, hacer una tercera llamada que, por cada ronda, me facilite el listado de deportistas que van a competir.</p>
<p>Tengo que definir el m&eacute;todo como as&iacute;ncrono para poder hacer que unas peticiones esperen por el resultado de la anterior, hasta donde he le&iacute;do ning&uacute;n problema por definir <code>async mounted()</code> y jugar con <code>await</code> y <code>Promises</code>, pero &iquest;qu&eacute; opin&aacute;is de esa tercera llamada en bucle?.</p>
<p>Entiendo que si dentro del bucle <code>forEach</code> de rounds defino el m&eacute;todo como as&iacute;ncrono podr&eacute; hacer las llamadas correspondientes esperando su ejecuci&oacute;n con <code>await</code>, tal que as&iacute;:</p>

```js
Array.from(this.rounds).forEach( async round =&gt; {
    await this.get_competitors_list(round.round_id);
    [...]
});
```

<p>Error, <code>forEach</code> no va a esperar a que hayan acabado los elementos marcados como tal, simplemente ejecuta y sigue. As&iacute; que, leyendo un poco de <a href="https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Iteration_protocols">Iteration protocols</a> veo que en su lugar podemos usar <code>for..of</code>, que es una versi&oacute;n moderna del "for loop" de forma que, simplemente cambiando forEach y adaptando la sintaxis al nuevo bucle haremos que funcione adecuadamente:</p>

```js
for(let round of this.rounds) {
    await this.get_competitors_list(round.round_id);
    [...]
}
```

<p>Y eso es todo, ni m&aacute;s ni menos; una tarde entera de debug y seguir trazas aqu&iacute; y all&aacute;. Todo muy gracioso y divertido.</p>
