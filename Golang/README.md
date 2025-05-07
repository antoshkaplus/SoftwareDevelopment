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
Can also run the following command: `go env GOMODCACHE`

https://go.dev/doc/go-get-install-deprecation
https://www.reddit.com/r/golang/comments/174wsiw/where_packages_are_installed/

