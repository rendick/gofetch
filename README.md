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
  __ _  ___   
 / _  |/ _ \   User: %s 
| (_| | (_) |  Hostname: %s 
 \__, |\___/   Distribution: %s 
 |___/    _    Kernel: %s 
 / _| ___| |_  Shell: %s
| |_ / _ \ __| Display: %s
|  _|  __/ |_  Uptime: %s
|_|  \___|\__| Wm: %s
  ___| |__     Terminal: %s 
 / __|  _ \    CPU: %s
| (__| | | |   Server: %s 
 \___|_| |_|   Memory: %s 
```

### Android 

```
  __ _  ___   
 / _  |/ _ \   User: %s 
| (_| | (_) |  Hostname: %s 
 \__, |\___/   Distribution: %s 
 |___/    _    Kernel: %s 
 / _| ___| |_  
| |_ / _ \ __| Shell: %s
|  _|  __/ |_  Uptime: %s
|_|  \___|\__| 
  ___| |__     Terminal: %s 
 / __|  _ \    CPU: %s
| (__| | | |   Server: %s 
 \___|_| |_|   Memory: %s 
```

# TODO

Add PKGBUILD for Arch Linux-like distros

# License

**GNU GPL v3**

