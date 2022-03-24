---
title: "Call date inside an alias"
date: 2015-07-02T21:03:20Z
draft: false
tags: [  ]
image: alias.jpg
---

<p>Stupidity of the day: calculate the <code>date</code> in an double quoted bash <code>alias</code>. It seems that, if you don't escape the <code>date</code> call, it's called at the time of the <code>alias</code> definition, so it's not the real behaviour I was looking for:</p>

```
$ date +%Y%m_%H%M
201507_2252
$ alias foo="echo `date +%Y%m_%H%M`"
$ foo
201507_2252
```

<p>Wait a minute and...</p>

```
$ date +%Y%m_%H%M
201507_2253
$ foo
201507_2252
```

<p>Not really usable when you want the date at the time you run the alias, not when it's defined, so you have to escape the date call this way:</p>

```
$ date +%Y%m_%H%M
201507_2253
$ alias foo="echo `date +%Y%m_%H%M`"
$ foo
201507_2253
```

<p>Wait a minute, or a couple of minutes and...</p>

```
$ foo
201507_2255
```

<p>That's all, stupid but so useful trick when you want to define something like:</p>

```
alias test_cov="py.test -s --tb=native --cov proj --cov-report term-missing | tee logs/proj-`date +%Y%m_%H%M`.log"
```

<p>Enjoy!</p>
