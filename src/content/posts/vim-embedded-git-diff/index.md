---
title: "Vim: embedded git diff"
date: 2023-10-11T10:30:17Z
draft: false
tags: []
image: 
---

Sometimes, all you want is a simple and rapid method to review all the changes made across different branches, in a format you can easily edit. This is particularly useful for quickly assessing which tests need to be added when working on a new branch.

My initial thoughts led me to consider various plugins, like [fugitive](https://github.com/tpope/vim-fugitive) or [lazygit](https://github.com/kdheepak/lazygit.nvim), as well as manual approaches such as copying `git diff` output into the buffer. But I had a hunch that there had to be a quicker way, and I was right:

```vim
:r!git diff master..my-branch --no-color
:set syntax=git
```

The solution lies in vim's powerful `:r` command. This feature allows you to import the output of an external command directly into your vim buffer. Breaking it down, `git diff` generates a comparison between the `master` and `my-branch` with the `--no-color` flag to remove any color formatting. Then vim imports this output into the buffer.

To enhance the readability, we apply git syntax highlighting to the inserted text using `:set syntax=git`. This combination of vim commands offers a fast and efficient solution for evaluating changes, navigating code, and identifying the tests that need to be incorporated.

I love `vim`.
