---
title: "Casual fun with tests"
date: 2015-05-13T18:17:15Z
draft: false
tags: [  ]
image: PhantomJS.png
---

<p>Fun is not the word because I went to bed really annoyed last two days, but it's the only way I can handle this, having as much "fun" as I can.</p>
<p>In one of my projects we are using <a href="http://pytest.org/latest/">py.test</a> + <a href="http://splinter.readthedocs.org/en/latest/">splinter</a> + <a href="http://www.seleniumhq.org/">selenium</a> + (<a href="http://phantomjs.org/">phantomjs</a>|<a href="https://sites.google.com/a/chromium.org/chromedriver/">chromedriver</a>) for headless testing. It's a good combination when it works, but the question is that rarely does in my enviroment.</p>
<p>Don't know whose fault, first times I've pointed to <em>chromedriver</em> (tried many versions), then I've changed to <em>phantomjs</em> and crashes were still there. Tried <em>phantomjs</em> downloaded binary from oficial web, other one from <em>npm</em>, even tried building my own binary!. Then I've changed the <em>selenium</em> version (<em>2.45.0</em>, <em>2.44.0</em>, <em>2.43.0</em>...), <em>splinter</em> (from <em>0.6.0</em> to <em>0.7.2</em>), I've tried rebuilding the enviroment from scratch, increasing OS limits... but nothing seemed to work.</p>

```
PhantomJS:
...
project/tests/test_integration/headless_1.py ....
project/tests/test_integration/headless_2.py ...............
project/tests/test_integration/headless_3.py .
project/tests/test_integration/headless_4.py ..........
project/tests/test_integration/headless_5.py .....................................
project/tests/test_integration/headless_6.py ...............
project/tests/test_integration/headless_7.py ........
project/tests/test_integration/headless_8.py .......
project/tests/test_integration/headless_9.py ....
project/tests/test_integration/headless_10.py ..
project/tests/test_integration/headless_11.py ....
project/tests/test_integration/headless_12.py .........
project/tests/test_integration/headless_13.py ...........
project/tests/test_integration/headless_14.py ............FFF
project/tests/test_integration/headless_15.py FFFFFFFFFFFFFFFFFFFFFFFF
 
File ".../lib/python2.7/site-packages/selenium/webdriver/phantomjs/service.py", 
line 75, in start raise WebDriverException("Unable to start phantomjs with ghostdriver.", e)
WebDriverException: Message: Unable to start phantomjs with ghostdriver.
Screenshot: available via screen
 
================================ 28 failed, 1809 passed, 2 skipped in 1825.82 seconds ================================
 
 
ChromeDriver:
...
project/tests/test_integration/headless_1.py ....
project/tests/test_integration/headless_2.py ...............
project/tests/test_integration/headless_3.py .
project/tests/test_integration/headless_4.py ................................
project/tests/test_integration/headless_5.py ...............
project/tests/test_integration/headless_6.py ........
project/tests/test_integration/headless_7.py .......
project/tests/test_integration/headless_8.py ..FF
project/tests/test_integration/headless_9.py FF
project/tests/test_integration/headless_10.py FFFF
project/tests/test_integration/headless_11.py FFFFFFFFF
project/tests/test_integration/headless_12.py FFFFFFFFFFF
project/tests/test_integration/headless_13.py FFFFFFFFFFFFFFF
project/tests/test_integration/headless_14.py FFFFFFFFFFFFFFFFFFFFFFFF
 
File ".../lib/python2.7/site-packages/selenium/webdriver/chrome/service.py", 
line 70, in start http://code.google.com/p/selenium/wiki/ChromeDriver")
WebDriverException: Message: 'chromedriver' executable needs to be available in the path.
Please look at http://docs.seleniumhq.org/download/#thirdPartyDrivers and read up at
http://code.google.com/p/selenium/wiki/ChromeDriver
 
================================ 67 failed, 1770 passed, 2 skipped in 1182.77 seconds ================================
 
$ phantomjs -v
1.9.8
(tested with 1.9.8 downloaded binary, from npm and built one)
 
$ chromedriver -v
ChromeDriver 2.15.322455 (ae8db840dac8d0c453355d3d922c91adfb61df8f)
 
$ pip freeze
selenium==2.45.0
splinter==0.7.2
 
Python 2.7.6 (default, Nov 12 2013, 13:26:39)
 
$ ulimit
unlimited
 
$ sysctl -a | grep maxfil
kern.maxfiles = 12288
kern.maxfilesperproc = 10240
kern.maxfiles: 12288
kern.maxfilesperproc: 10240
 
$ sudo sysctl -w kern.maxfiles=65536
$ sudo sysctl -w kern.maxfilesperproc=65536
 
$ sysctl -a | grep maxfil
kern.maxfiles = 65536
kern.maxfilesperproc = 65536
kern.maxfiles: 65536
kern.maxfilesperproc: 65536
 
 
$ launchctl limit maxfiles
maxfiles     256            unlimited
 
$ launchctl limit maxproc
maxproc     709            1064
```

<p>Integration tests starts working great, but suddenly "it" starts to crash and once it crashes first time, I only get <code>F</code> til the end. I've opened a <a href="https://github.com/cobrateam/splinter/issues/400">issue in splinter github project</a>, I've asked in <em>#phantomjs</em>&nbsp;freenode irc channel but there are no clues about that strange behaviour. It exceeds my patience.</p>
<p>Casually one of my fellows said something about running tests slightly faster with multiprocessing option (<code>-n</code>) and once we tested it - just curiosity - we realized that tests were not failing in same way, moreover, they were finally passing!. Speed is not the key, because the improvement is not that much, but who cares... test are passing!!.</p>

```
$ py.test -s --tb=native --cov project --cov-report term-missing -n 2
```

<p>In the end we don't know what's happening with our enviroments, don't know who's fault, but we find a temporary solution in a stupid an casual way. "Fun".</p>
