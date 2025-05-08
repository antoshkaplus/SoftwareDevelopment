To run as a script:
```bash
go run ${filename}.go
```

Golang does not usually use dynamic linking like C++, Python or .Net.
Usually the whole program is statically linked into one binary.
That simplifies dependency management.

People download Golang packages as source code to cache location 
on a local system and then compile on the current system.  
Golang packages are stored in `/home/<user>/go/pkg/mod`.  
Can also run the following command: `go env GOMODCACHE`.  

https://go.dev/doc/go-get-install-deprecation
https://www.reddit.com/r/golang/comments/174wsiw/where_packages_are_installed/

`go env` provides a bunch of set parameter variables.

GOPATH can be any directory on your system. In Unix examples, we will set it to \$HOME/go (the default since Go 1.8). 
Note that GOPATH must not be the same path as your Go installation. Another common setup is to set GOPATH=\$HOME.

```
gopls (pronounced "Go please") is the official Go language server developed by the Go team. It provides a wide variety of IDE features to any LSP-compatible editor.

You should not need to interact with gopls directly--it will be automatically integrated into your editor. The specific features and settings vary slightly by editor, so we recommend that you proceed to the documentation for your editor below. Also, the gopls documentation for each feature describes whether it is supported in each client editor.
```

```
The GOPATH environment variable specifies the location of your workspace. If no GOPATH is set, it is assumed to be $HOME/go on Unix systems and %USERPROFILE%\go on Windows. If you want to use a custom location as your workspace, you can set the GOPATH environment variable. This page explains how to set this variable on various platforms.
```

## Go Modules

Before `go mod` was introduced, people would keep projects under `$GOPATH/src` or `$GOPATH/src/github.com/user/package`.
There's no need for it now. Just place it wherever you like, and if you want to use other dependencies just do `go mod init project-name` (usually project-name will be your repo) then `go get` within that project will add the dependencies to your project.

When you call go mod init it will create go.mod and go.sum files. They are equivalent to package.json and package-lock.json in a node project. This two files hold the info about your project and your dependencies

TL;DR Place it wherever you like in your file system and run go mod init github.com/user/repo inside that folder

### Reference
https://www.reddit.com/r/golang/comments/q4lpr6/comment/hfzegv8/?utm_source=share&utm_medium=web3x&utm_name=web3xcss&utm_term=1&utm_content=share_button

