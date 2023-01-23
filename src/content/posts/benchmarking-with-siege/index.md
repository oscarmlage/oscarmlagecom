---
title: "Benchmarking with Siege"
date: 2022-12-29T18:28:43Z
draft: false
tags: []
image: 
---

If you ask me, I enjoy starting a new project, that fresh feeling during the first steps is usually becoming into a not-so-fresh thing with time. Life is so :).

In this concrete case, we're talking about an API, I needed to research and make some benchmarking work to take the best decision possible for the stack. It needs to be as fast as we can, so we can take advantage of some tools like `ab` or `siege` to measure some kinds of requests and scenarios.

With those tools in one hand and some languages in the other (`go-lang` and `php` mainly, in fact, the benchmark would be go-lang using different libraries but I was curious about a compiled language vs the fastest interpreted one I know so  I added some php code), let's start having fun!.

All the tests were done in the same machine (macbook pro intel i9-9980hk) + docker containers + a dockerized mysql. The API endpoint makes a couple of SQL queries to return about 50 rows of data from a table, the response is being `jsonized`. There is no auth and no more logic steps, just an endpoint that makes a sql query and returns the data as json.

- Go-Lang + [Gorm](https://gorm.io) + Custom Router: the custom router uses the golang stdlib `"net/http"` to serve GET petitions.
- Go-Lang + [Gorm](https://gorm.io) + [Echo Server](https://echo.labstack.com).
- Go-Lang + [Gorm](https://gorm.io) + [Gin](https://github.com/gin-gonic/gin).
- Laravel + DB raw: as Laravel was the slowest, decided to use raw sql instead Eloquent to balance a bit the results :P, queries were done with `DB::Select()`.
- PHP "vanilla": some memories came to my mind using `mysqli_connect()`, `mysqli_query()`, and `mysqli_fetch_array()`.

I've run `siege` 3 times in different scenarios and calculated the average:

- 100 requests, only 1 concurrent
- 100 requests, 2 concurrents (50x2)
- 100 requests, 5 concurrents (20x5)
- 100 requests, 10 concurrents (10x10)
- 1000 requests, 10 concurrents (100x10)
- 10000 requests, 100 concurrents (100x100)
- Max requests in 30 seconds, 10 concurrents
- Max requests in 30 seconds, 100 concurrents
- Max requests in 30 seconds, 150 concurrents: I got some errors here, so it doesn't matter for the graphics

And the result is quite interesting:

{{< gallery match="gallery/*" sortOrder="asc" rowHeight="150" margins="5" thumbnailResizeOptions="752x436 q90 Lanczos"  previewType="blur" embedPreview="true" >}}

In this benchmark **Go+Gorm+Echo** is the clear winner, closely followed by **Go+Gorm+Gin**. About PHP, dunno if I did something wrong but there is a lot of difference between vanilla PHP (the results are kinda acceptable) and Laravel, still dunno why.

I had so much fun with this "exercise".
