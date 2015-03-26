---
title: Cross Compiling
layout: default
status: unfinished
---

## {{ page.title }}
{:.recipe}

How do I compile a binary designed to run on a different platform than my local host?
{:.question}

Since Go is a statically compiled language, it's well designed for producing tiny, pre-compiled tools.  This is great for Unix administrators or DevOps engineers, as they can send their tools with them when they work on remove systems.  However, many developers use OS X or Windows on their local machine, and another OS like Linux on others.  How do you use one system to compile binaries that will work on the other?  This is called cross-compiling, and it's pretty simple with Go.

### Pre 1.5

Before 1.5, cross compilation was an arduous process, requiring [massive scripts](http://dave.cheney.net/2013/07/09/an-introduction-to-cross-compilation-with-go-1-1) to build separate go compilers for each supported platform.

### Post 1.5

Things have been much better since v1.5 of Go was released.  Go now supports 

http://dave.cheney.net/2015/03/03/cross-compilation-just-got-a-whole-lot-better-in-go-1-5

Let's take the example program, which simply prints the OS and Architecture it's compiled for.  While this will always be the same as the platform it's being executed on, it's really a compile-time value - not something determined at runtime.

```
package main

import "fmt"
import "runtime"

func main() {
    fmt.Printf("OS: %s\nArchitecture: %s\n", runtime.GOOS, runtime.GOARCH)
}
```

Now, let's compile this program for an Apple MacBook.  To do so, we simply set two environment variables: `$GOOS`, which is the target operating system, and `$GOARCH`, which is the target processor.  Then we run `go build` as normal:

```
$ export GOOS=darwin 
$ export GOARCH=386 
$ go build test.go
```

Note that the `test` executable that comes out the other end only runs on OS X, and cannot be run on Linux or Windows.  If, on the other hand, we wanted to compile for Microsoft Windows, we'd set `GOOS=windows` and `GOARCH=386`.

When we run the resulting binary on the right platform, we see:

```
$ ./test
OS: windows
Architecture: 386
```

### Testing via Docker?

TODO?

This is another example of Go's strengths as a language for building easily distributed tools.  It's almost trivial to modify your build process to produce binaries for every major platform your users may run.  This is used, for example, to enable the Cloud Foundry team to distribute [Linux, Windows, and OS X CLI tools](https://github.com/cloudfoundry/cli).
