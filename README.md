# gofetch

System information tool written in Golang 

# Available OS

GNU/Linux & Android

# Dependencies

```
lsb-release, xrandr
```

# Install

```bash
git clone https://github.com/rendick/gofetch.git
cd gofetch
go build
./gofetch
```

# Functions

### GNU/Linux

```
System: *system name* (System: GNU/Linux)
User: *user* (User: adolf)
Kernel: *kernel* (Kernel: linux-zen)
Shell: *shell* (Shell: /bin/zsh)
Uptime: *uptime* (10 hours, 10 min)
Resolution: *resoulution* (Resolution: 1920x1080)
Server: *server* (Server: x11) echo $XDG_SESSION_TYPE
Terminal: *terminal* (Terminal: kitty)
Battery: *battery* (Battery: 50%, 4 hours)
CPU: *name* | *temp* (CPU: AMD Ryzen 7 2700 | 90)
GPU: *name* | *temp* (GPU: NVIDIA RTX 3060 | 100 )
Theme: *theme* (Theme: Dracula)
WM & DE: *wm or de* (WM & DE: bspwm)~~
Memory: *freememory / allmemory* (Memory: 4223/15456)
Packages: *packages* (Packages: 2343)
```

# TODO

Add PKGBUILD for Arch Linux-like distros

# License

**GNU GPL v3**

