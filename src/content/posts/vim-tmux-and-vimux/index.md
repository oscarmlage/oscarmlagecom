---
title: "Vim, tmux and vimux"
date: 2018-06-13T18:13:27Z
draft: false
tags: [  ]
image: tmux-vim-vimux.png
---

<p>Hey, vim + tmux calling again!</p>
<p>If you know me by now, you know I feel really happy if I'm able to work without touching the mouse at all. In this particular case I wanted to run the tests while working in the project.</p>
<p>I could just open a new tmux pane, run the tests command there and go back to vim, but if there is a plugin that can save me a couple of keystrokes... I need to test it!.</p>
<p style="text-align: center;"><img style="width: 100%; height: 100%;" src="gallery/tmux-vim-vimux.gif" alt="tmux, vim, vimux" /></p>
<p>You can <a href="gallery/vimux">Vimux</a> and it just opens a tmux pane (or window), execute the command you want, and go back to vim. The configuration I wrote (<code>~/.vimrc</code>):</p>

```vim
"---------------------------------------------------------
"------------- Vimux
"---------------------------------------------------------
let g:VimuxHeight = "20"
let g:VimuxOrientation = "v"
" let g:VimuxRunnerType = "pane"
map rr :call VimuxPromptCommand()
map rl :call VimuxRunLastCommand()
map rc :call VimuxCloseRunner()
```

<p>Easy peasy!.</p>
