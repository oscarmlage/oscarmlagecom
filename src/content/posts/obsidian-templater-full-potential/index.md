---
title: Obsidian Templater, full potential
date: 2024-11-13T09:37:59+01:00
draft: false
tags: []
---

I love automating anything that even hints at being repetitive. I don’t mind investing time in it because I know that sooner or later I’ll be grateful I did. Another one of my “sins” (from an excess) is documentation—I’m always thinking about my future self.

In almost every project I’m involved with, I feel the need to keep private notes that, in some way, link to the project’s public documentation. For example, if I’m working on an issue, I tend to write down a step-by-step log of what I’ve done, what I’m currently working on, and what’s next. I also make other notes that might help me solve it (snippets, related code, resources, links, etc.). Obviously, it wouldn’t be appropriate to share all of this information in the public issue (description, comments, etc...).

The solution to this dilemma has usually involved creating a note manually in the project’s `issues/` folder, matching the issue number (and title) of the one in the version control system (GitHub, Forgejo, Gitea, etc.). And let’s be honest—it’s a hassle.

So, since I’m managing my notes in [Obsidian](https://obsidian.md) and already using the [Templater](https://github.com/SilentVoid13/Templater) plugin for other tasks, I decided to invest time in automating this process.

First, we need to create a template that will serve as the base for all these new notes that will be created in the `issues/` directory mentioned earlier. This setup will allow *Templater* to take control of creating these notes and applying the template logic. Here’s the plan:

1. Ask me for the title of the issue and, optionally, a brief description.
2. Run a `curl` command against the version control server ([Forgejo](https://forgejo.org) or [Gitea](https://about.gitea.com) in this case) to create an issue with the given title.
3. Capture the issue number (e.g. `34`) from the `curl` response.
4. Create a new note with the template’s content (including a reference to the URL of the created issue for convenience).
5. Rename the note’s title to add the issue number at the beginning (e.g. `34 - Issue Title`).

These are basically the steps I used to do manually, so now let’s head to the `templates/` directory in *Obsidian* and create a `repo-issue` template with the following content:

```js
<%*
    const title = await tp.system.prompt("Issue title");
    const content = await tp.system.prompt("Issue content");
    const token = "your-forgejo-or-gitea-token";
    const url = "https://git.yourdomain.com/api/v1/repos/owner/repo/issues";

    const curlCommand = `curl -s -X POST -H "Content-Type: application/json" -H "Authorization: token ${token}" -d "{\\"title\\": \\"${title}\\", \\"body\\": \\"${content}\\"}" ${url}`;
    const exec = await tp.user.curl({ command: curlCommand });

    const response = JSON.parse(exec);
    const issueNumber = response.number.toString().padStart(3, '0');
    const issueUrl = response.url;

    const titleWithNumber = `${issueNumber} - ${title}`;
    await tp.file.rename(titleWithNumber);
%>

# <% titleWithNumber %>

**Issue URL:** [Issue #<% issueNumber %>](<% issueUrl %>)

<% content %>

## 📄 Log

- [[<% tp.file.creation_date("YYYY-MM-DD") %>]] Issue created

## 📔️ Notes

## 🔗 Resources

## ✅ Tasks

- [ ] 

```

We still need to add the function we’re calling `tp.user.curl()` to make the request for creating the issue. To do this, we need to create a file, let’s say `curl.js`, in our `scripts/` directory, and tell *Templater* in the configuration to load that file so that the method is available. The content of our `curl.js` would look something like this:

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

Now we have all the ingredients prepared and ready, we just need to cook them (the *Templater* configuration), and there are two simple steps:

1. Load the script to make `tp.user.curl()` available.
{{< gallery match="gallery/tp-user-curl*" sortOrder="asc" rowHeight="150" margins="5" thumbnailResizeOptions="600x600 q90 Lanczos" previewType="blur" embedPreview="true" >}}
1. Configure the `issues/` directory to run the `repo-issue` template we created earlier.
{{< gallery match="gallery/repo-issue*" sortOrder="asc" rowHeight="150" margins="5" thumbnailResizeOptions="600x600 q90 Lanczos" previewType="blur" embedPreview="true" >}}

And with that, everything is automated. Every time we create a new note in the `issues/` directory, not only will the note with the corresponding template content be created, but an issue will also be created in the version control server, and the note will be renamed with the issue number and link directly in the note.

{{< gallery match="gallery/issue*" sortOrder="asc" rowHeight="150" margins="5" thumbnailResizeOptions="600x600 q90 Lanczos" previewType="blur" embedPreview="true" >}}

For little details like this, *Obsidian* and *Templater* has earned a special place in my heart.