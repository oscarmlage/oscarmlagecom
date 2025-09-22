---
title: obsidian-templater-gitea
date: 2024-11-11 21:21:50  +0200
draft: false
tags: micropost
---

[#TIL](https://mastodon.bofhers.es/tags/TIL) that with [#Obsidian](https://mastodon.bofhers.es/tags/Obsidian) and the [#Templater](https://mastodon.bofhers.es/tags/Templater) plugin, it’s possible to automatically create a new note and a related issue in [#Gitea](https://mastodon.bofhers.es/tags/Gitea) or [#Forgejo](https://mastodon.bofhers.es/tags/Forgejo) at the same time. Made my day.

You can define a custom function that makes a curl request  as `tp.user.curl()`:

```js
const { exec } = require('child_process');

module.exports = async (args) => {
    return new Promise((resolve, reject) => {
        exec(args.command, (error, stdout, stderr) => {
            if (error) {
                console.error(`Error: ${error.message}`);
                reject(error);
            } else if (stderr) {
                console.error(`Stderr: ${stderr}`);
                reject(stderr);
            } else {
                console.log(`Result: ${stdout}`);
                resolve(stdout);
            }
        });
    });
};
```

Then the function call:

```js
<%*
const curlCommand = `curl -s -X POST -H "Content-Type: application/json" -H "Authorization: token ${token}" -d "{\\"title\\": \\"${title}\\", \\"body\\": \\"${content}\\"}" ${url}`;
const exec = await tp.user.curl({ command: curlCommand });
%>
```

When you execute this template, the curl request will be triggered, creating the issue as specified.