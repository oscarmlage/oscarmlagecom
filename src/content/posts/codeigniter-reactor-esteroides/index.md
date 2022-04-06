---
title: "¿Codeigniter-Reactor + esteroides?"
date: 2011-02-04T20:09:38Z
draft: false
tags: [ "php", "code" ]
image: 
---

<p>
	Es algo que todav&iacute;a no llego a entender ni asimilar. He pasado la mayor parte de la tarde para configurar la nueva versi&oacute;n 2.0 de <a href="https://github.com/philsturgeon/codeigniter-reactor/">CodeIgniter-Reactor</a> con varias librer&iacute;as que, para mi forma de desarrollar, son indispensables en cualquier framework de programaci&oacute;n orientado a web:</p>
<ul>
	<li>
		<a href="https://bitbucket.org/wiredesignz/codeigniter-modular-extensions-hmvc">HMVC</a>: Gracias a la librer&iacute;a de <em>Wiredesignz</em> podemos ordenar nuestro c&oacute;digo en m&oacute;dulos y simplificar la l&oacute;gica de la aplicaci&oacute;n.</li>
	<li>
		<a href="https://github.com/pyrocms/pyrocms/tree/master/system/pyrocms/modules/modules">ModulesModule</a>: Ahora que todo ser&aacute; un m&oacute;dulo, no vendr&iacute;a mal un m&oacute;dulo encargado de ejecutar las tareas m&aacute;s comunes de los m&oacute;dulos (instalar, desinstalar, actualizar...). Un m&oacute;dulo de m&oacute;dulos.</li>
	<li>
		<a href="https://github.com/pyrocms/pyrocms/tree/master/system/pyrocms/modules/settings">SettingsModule</a>: No me gusta cargar la configuraci&oacute;n desde ficheros <em>config/*</em>, por comodidad y para que el usuario pueda cambiar cualquier par&aacute;metro ajustable de una aplicaci&oacute;n, prefiero hacerlo en base de datos y cachearlo a disco si es necesario. Me quedo con la librer&iacute;a de <a href="http://pyrocms.com">PyroCMS</a>.</li>
	<li>
		<a href="https://github.com/pyrocms/pyrocms/tree/master/system/pyrocms/modules/themes">ThemesModule</a>: Otra caracter&iacute;stica imprescindible ser&iacute;a tener una aplicaci&oacute;n <em>themeable</em> y que desde un interfaz de administraci&oacute;n se pueda cambiar tranquilamente el aspecto de la misma. Para ello podemos hacer uso de este m&oacute;dulo capaz de instalar y desinstalar themes.</li>
	<li>
		<a href="https://github.com/pyrocms/pyrocms/blob/master/system/pyrocms/libraries/Template.php">TemplateLibrary</a> y <a href="https://github.com/pyrocms/pyrocms/blob/master/system/pyrocms/libraries/Tags.php">TagsLibrary</a>: Una vez hemos decidido hacer la aplicaci&oacute;n <em>modular</em> y <em>themeable</em> siguiendo el patr&oacute;n <em>MVC</em>, un buen lenguaje de <em>template</em> para que los dise&ntilde;adores no se vuelvan locos con el lenguaje din&aacute;mico ser&iacute;a un punto m&aacute;s.</li>
	<li>
		<a href="https://github.com/philsturgeon/codeigniter-migrations/">MigrationsLibrary</a>: Una vez lo pruebas se convierte en <em>musthave</em>. Se trata de una librer&iacute;a para hacer migraciones de versiones en base de datos. Gestiona los cambios entre versiones de forma sencilla.</li>
</ul>
<p>
	<!--more-->Como pod&eacute;is ver siguiendo los enlaces, casi todo lo necesario para montar el esquema inicial de un framework PHP decente est&aacute; en los repositorios de <a href="https://www.philsturgeon.co.uk/">Phil Sturgeon</a> (<a href="https://bitbucket.org/philsturgeon/">bitbucket</a> y <a href="https://github.com/philsturgeon/">github</a>). Su <a href="http://www.pyrocms.org">PyroCMS</a> cuenta con todas estas caracter&iacute;sticas pero tiene muchas otras librer&iacute;as que no son necesarias -a mi modo de ver- para tomarlo como base para cualquier otro tipo de proyecto.</p>
<p>
	Demasiado complejo de ensamblar, una vez he conseguido hacerlo funcionar m&aacute;s o menos todo ya no tengo ganas de empezar a programar la aplicaci&oacute;n que ten&iacute;a en mente. &iquest;Por qu&eacute; no hay *nada* con todas esas funciones&nbsp; de serie en PHP?. Si, ya s&eacute; que con <a href="http://www.rubyonrails.org/">Ruby on Rails</a> o <a href="http://djangoproject.org">Django</a> hubiera tenido todas estas caracter&iacute;sticas -y m&aacute;s- en un par de comandos, pero ten&iacute;a que desenga&ntilde;arme.</p>
<p>
	En este sentido PHP sigue estando tan verde como hace a&ntilde;os, me hace perder demasiado tiempo a&uacute;n conociendo la estructura del lenguaje. Todos sabemos como es la curva de aprendizaje al principio, pero -aunque duela- ha llegado el momento de invertir en otra cosa. <em>Shame on me</em>?</p>
