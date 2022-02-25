---
title: "Python, Zombie, forms, tables and tests"
date: 2015-04-07T22:25:11Z
draft: false
tags: [  ]
image: python.jpg
---

<p>Sometimes you feel that the world is against you, sometimes it's a computer what makes your life miserable. This time was a cocktail of elements. Trying to test a simple form functionally from <a href="https://github.com/ryanpetrello/python-zombie/">python-zombie</a> was like a nightmare.</p>
<p>I'm working testing a project using <code>node v0.8.23</code>, <code>npm 1.4.4</code>, <code>zombie 1.4.1</code> and <code>python-zombie 0.1.0a5</code>. The template I'm trying to render while testing has a code like:</p>

```
&lt;form action="custom/url/delete" method="POST"&gt;
    &lt;button&gt;Delete&lt;/button&gt;
&lt;/form&gt;
```

<p>The test I'm trying to run is:</p>

```
from zombie import Browser
...
def test_remove_category(self):
    browser = self.browser
    ...
    # Go to the edit link
    browser.clickLink('a[href$="/edit"]')
    import pdb; pdb.set_trace()
    browser.pressButton('form[action$=delete] button[type=submit]')
    ...
```

<p>The problem is the <code>pressButton</code> line (ignore the pdb for a moment), it's trying to press the button inside the form but the test is failing raising an error like:</p>

```
NodeError: Error: No BUTTON 'form[action$=delete] button[type=submit]'
```

<p>Trying to debug the problem - now is where pdb gets the focus - I got that the form was not nicely rendered:</p>

```
(Pdb) print(browser.html())
....
&lt;form method="POST" action="custom/url/delete"&gt;&lt;/form&gt;
    &lt;button type="submit" name="submit"&gt;Delete&lt;/button&gt;
```

<p>After an intense testing + browsing + reading task I've realized wtf was happening, taking a look to the template code we have the form inside atable, as child of a &lt;tbody&gt; element, something like:</p>

```
&lt;table&gt;
 &lt;thead&gt;
  ...
 &lt;/thead&gt;
 &lt;tbody&gt;
  &lt;tal:r tal:repeat="q context.values()"&gt;
   &lt;tr&gt;...&lt;/tr&gt;
   &lt;div&gt;&lt;form&gt;...&lt;/form&gt;&lt;/div&gt;
  &lt;/tal:r&gt;
 &lt;/tbody&gt;
&lt;/table&gt;
```

<p>But a form is not allowed to be a child of <code>table</code>, neither <code>tr</code> or <code>tbody</code> elements so I think it was the reason why it was not nicely rendered by zombie.</p>
<p>After all it was no problem of zombie, the template was weird and, we've to be honest... it was funny and a great lucky to see this failing test.</p>
