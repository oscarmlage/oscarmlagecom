---
title: shotabline
date: 2025-10-30 12:14:19  +0200
draft: false
tags: micropost
---

I didn’t want to show the Neovim tabline. I had removed the `barbar` plugin, but the tabline was still there. I thought another plugin (like `winbar` or `lualine`) might be responsible, but in the end it turned out to be much simpler than that:

```lua
vim.opt.showtabline = 0
```