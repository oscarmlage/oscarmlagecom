---
title: "Pyramid: set a cookie returning a HTTPFound"
date: 2015-03-30T10:27:13Z
draft: false
tags: [  ]
image: pyramid2015.png
---

<p>Probably it will not be the most common scenario, but if you want to save a cookie and make a redirection right away using <a href="http://docs.pylonsproject.org/docs/pyramid/en/latest/api/httpexceptions.html#pyramid.httpexceptions.HTTPFound">HTTPFound</a>, the cookie won't being saved. That's the fact.</p>
<p>It seems that <code>HTTPFound</code> takes a own headers parameter, probably overwritting the one you used to set the cookie, so theorically nothing happened with your <code>response.set_cookie()</code> unless you call <code>HTTPFound</code> with the proper argument:</p>

```
response.set_cookie('mycookie', value=myvalue)
return HTTPFound(location=request.environ['HTTP_REFERER']),
                 headers=response.headers)
```

<p>Thanks to <a href="http://stackoverflow.com/users/617711/rob-wouters">Rob Wouters</a> for the excellent reply in <a href="http://stackoverflow.com/questions/8746087/pyramid-how-to-set-cookie-without-renderer">this Stack Overflow thread</a>, it was hard to find it for me, that's the main reason of this post.</p>
