---
title: "Using mercurial a la darcs"
date: 2015-04-09T10:54:41Z
draft: false
tags: [ "code" ]
image: 5771027bc55731cae1a8fa6b44a43754.jpg
---

<p>Nowadays using a version control system is as basic as using an editor. I have to admit I'm a mercurial fanboy because it's clear, simple and written in python.</p>
<p>In some other projects we're using <a href="http://darcs.net/">darcs</a>, similar to <a href="http://mercurial.selenic.com/">mercurial</a> but with a different approach. As we've organized the darcs project, we must pull, create a <em>patch</em> with the new features, and send it to our QA department for review and bug hunting. If there are errors we must fix them and rewrite + resend the patch for a new review. Once the code is ok, the SYS department pushes the patch to <em>staging</em> (if there are more tests to be done) and finally to <em>production</em> enviroment.</p>
<p>The mercurial flow, ordinarily, is something like <em>pull</em>, write new features (or fix bugs), <em>commit</em> or <em>record</em> and <em>push</em>. We can have different repos like <em>repo-devel</em>, <em>repo-staging</em> and <em>repo-production</em> and play testing new features and <em>pushing</em> stuff between repos and enviroments. But what if we want to work with <em>patches</em> in a similar way that the one I've mentioned above?. Easy peasy.</p>
<p>I would have never noticed, but this is one of the advantages to work with smart people as <a href="https://twitter.com/WuShell">Borja</a>, the flow in this case should be... <em>pull</em>, write new features (or fix bugs), create a <em>patch</em> with <code>hg diff &gt; new_feature.patch</code> and send it to review. The reviewer will <em>apply</em> with <code>patch -p1 &lt; new_feature.patch</code> and he/she can edit, <em>commit</em>, <em>record</em> or <em>revert</em>. So easy too!.</p>
<p>Note that I'm not used to <a href="http://mercurial.selenic.com/wiki/MqExtension">Mercurial Queues Extension</a>, probably a better way to proceed in this case.</p>
