---
title: "Django and memcache: clear cache keys"
date: 2014-02-05T13:35:54Z
draft: false
tags: [ "django", "code" ]
image: web-accelerators-memcached.jpg
---

<p>Let's play&nbsp;<a href="https://www.djangoproject.com/">Django</a> with <a href="http://memcached.org/">Memcached</a>. As the great framework Django is, it's so easy to activate <a href="https://docs.djangoproject.com/en/dev/topics/cache/">any kind of cache</a> in your project. <em>Memcached</em> is one of the options, but you can also work with <em>DatabaseCache</em>, <em>FileBasedCache</em>, <em>LocMemCache</em>, <em>MemcachedCache,</em>&nbsp;<em>DummyCache</em> (a kind of non-cache very useful for devel/test enviroments) or - of course - your own <em>CustomCache</em> if you want.</p>
<p><strong>Activating cache</strong></p>
<p>It's too easy to activate the cache feature, it's enough to set the preferences in settings, install <a href="https://pypi.python.org/pypi/python-memcached/">python-memcached</a>&nbsp;in your enviroment (in case you will use <em>MemcachedCache</em>), and not much more to do. A couple of examples:</p>
<p>1. Basic <em>FileBasedCache</em> settings:</p>

```
# settings.py
CACHES = {
    'default': {
        'BACKEND': 'django.core.cache.backends.filebased.FileBasedCache',
        'LOCATION': '/var/tmp/django_cache',
    }
}
```

<p>2. <em>MemcachedCache</em> settings:</p>

```
# settings.py
CACHES = {
    'default': {
        'BACKEND': 'django.core.cache.backends.memcached.MemcachedCache',
        'LOCATION': '127.0.0.1:11211',
    }
}
```

<p>3. Depending on the enviroment you can use <em>MemcachedCache</em> and <em>DummyCache</em>:</p>

```
# settings.devel.py
CACHES = {
    'default': {
        'BACKEND': 'django.core.cache.backends.dummy.DummyCache',
    }
}
```


```
# settings.production.py
CACHES = {
    'default': {
        'BACKEND': 'django.core.cache.backends.memcached.MemcachedCache',
        'LOCATION': '127.0.0.1:11211',
    }
}
```

<p><strong>Setting the places where cache will act</strong></p>
<p>Now that we have our project with a configured kind of cache, we must to say where and when to activate it. There are multiple ways to do it (on the <a href="https://docs.djangoproject.com/en/dev/topics/cache/#the-per-site-cache">whole site</a>, <a href="https://docs.djangoproject.com/en/dev/topics/cache/#the-per-view-cache">views</a>, <a href="https://docs.djangoproject.com/en/dev/topics/cache/#template-fragment-caching">templates</a>, <a href="https://docs.djangoproject.com/en/dev/topics/cache/#specifying-per-view-cache-in-the-urlconf">urls</a>...). I'm going to use... let's say urls. So, in our <em>urls.py</em> we have to set a time and activate the cache:</p>

```
# urls.py

from django.views.decorators.cache import cache_page

ctime = (60 * 24) # A day
urlpatterns = patterns('',
    url(r'^$',
        cache_page(ctime)(BlogIndexView.as_view()),
        {},
        'blog-index'
        ),
        ...
)
```

<p>A simple server reload will be enough to have cache running. We can see it on action with <a href="https://github.com/django-debug-toolbar/django-debug-toolbar">django-debug-toolbar</a>, <a href="https://github.com/bartTC/django-memcache-status">django-memcache-status</a> or something like that.</p>
<p><strong>Clean a specific key cache</strong></p>
<p>And now the funniest part. For example, talking about a blog tool, when you write a new post (or editing older one) the software should be able to remove some cache keys, i.e. the <em>blog-index</em> one (because you have a new post) and the <em>post-detail</em> other (because you must be able to inmediately see the changes in the post you're editting).</p>
<p>Following <a href="http://stackoverflow.com/questions/2268417/expire-a-view-cache-in-django">this link</a> I've created a cache.py with this content:</p>

```
# cache.py

# -*- coding: utf-8 -*-

def expire_view_cache(
    view_name,
    args=[],
    namespace=None,
    key_prefix=None,
    method="GET"):

    """
    This function allows you to invalidate any view-level cache.

    view_name: view function you wish to invalidate or it's named url pattern
    args: any arguments passed to the view function
    namepace: optioal, if an application namespace is needed
    key prefix: for the @cache_page decorator for the function (if any)

    http://stackoverflow.com/questions/2268417/expire-a-view-cache-in-django
    added: method to request to get the key generating properly
    """

    from django.core.urlresolvers import reverse
    from django.http import HttpRequest
    from django.utils.cache import get_cache_key
    from django.core.cache import cache
    from django.conf import settings
    # create a fake request object
    request = HttpRequest()
    request.method = method
    if settings.USE_I18N:
        request.LANGUAGE_CODE = settings.LANGUAGE_CODE
    # Loookup the request path:
    if namespace:
        view_name = namespace + ":" + view_name
    request.path = reverse(view_name, args=args)
    # get cache key, expire if the cached item exists:
    key = get_cache_key(request, key_prefix=key_prefix)
    if key:
        if cache.get(key):
            cache.set(key, None, 0)
        return True
    return False
```

<p>And last step is to call the&nbsp;expire_view_cache function on model form save hook (<em>admin.py</em> in this case):</p>

```
# admin.py

from cache import expire_view_cache

class PostAdminForm(admin.ModelAdmin):
    ...
    def save_model(self, request, obj, form, change):
        expire_view_cache("blog-index")
        expire_view_cache("post-detail", [obj.slug])
```

<p>And that's all, we are able to <em>clean/purge/remove</em> the cache when a new post is added or edited. As you can see in the code, cache is fun but you have to be careful to set it on the right way.</p>
