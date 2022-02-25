---
title: "copy-paste mode in tmux 2.4"
date: 2017-05-25T10:47:12Z
draft: false
tags: [  ]
image: tmux-iterm.png
---

<p>Some changes has happened in the last version of <a href="https://tmux.github.io/">tmux</a>. Suddenly the <em>copy-paste </em>&nbsp;was not running but had no time to research the reason until minutes ago:</p>

```
bind-key -Tcopy-mode-vi Escape cancel
bind-key -Tcopy-mode-vi 'v' send -X begin-selection
bind-key -Tcopy-mode-vi 'V' send -X select-line
bind-key -Tcopy-mode-vi 'y' send -X copy-pipe-and-cancel "reattach-to-user-namespace pbcopy"
bind-key p paste-buffer
unbind -Tcopy-mode-vi Enter
bind-key -Tcopy-mode-vi Enter send -X copy-pipe-and-cancel "reattach-to-user-namespace pbcopy"
bind-key -Tcopy-mode-vi MouseDragEnd1Pane send -X copy-pipe-and-cancel "reattach-to-user-namespace pbcopy"
```

<p>There are a couple of issues (<a href="https://github.com/tmux/tmux/issues/592">#592</a>, <a href="https://github.com/gpakosz/.tmux/issues/42">#42</a>) related. Glad this amazing feature is back again.</p>
<p><iframe src="https://www.youtube.com/embed/ho4355YKf4Y" width="100%" height="420"></iframe></p>
