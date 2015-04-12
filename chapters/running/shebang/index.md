---
title: Running a file via a shebang line
question: How can I add a shebang line to the top of my Go file in order to run it directly, as though it were an interpreted script?
---

### "Shebang??"

In order to make it possible to execute scripts as though they were first class executables, UNIX systems will looks for what we refer to as a shebang line at the top of the file.  The origin of the name is murky.  Some think it came from sharp-bang or hash-bang -- contractions of `#` ("sharp") and `!` ("bang").  Others think the "SH" is in reference to the first UNIX shell, named "sh".

In any case, if a UNIX system sees that the first line of an executable begins with `#!`, then it will execute the file using whatever command is specified in the rest of the line.  For example, if there's a file named `/path/to/bar` that looks like:

```
#!/bin/foo
...
```

Then, when run, the kernel will execute `/bin/foo /path/to/bar`.  

This is great for interpreted languages like Ruby and Python which allow `#` as an innocuous comment.  A clever gopher might attempt to create a file named `/path/to/example` that looks like such:

``` golang
#!/usr/local/bin/go run
package main

import "fmt"

func main() {
    fmt.Println("Hello!")
}
```

And expect to run it and see "hello" in their terminal.  However, Go does not allow `#` as a comment, which causes this to fail with a syntax error:

```
$ ./fail.go
fail.go:1:1: illegal character U+0023 '#'
```

In order for this to work, the Go spec would have to be modified to accept `#` as a comment on the first line of the file.  However, the Go authors are highly opinionated.  When (repeatedly) asked for this feature, [Rob Pike has said](https://groups.google.com/d/msg/golang-nuts/iGHWoUQFHjg/_dbLKomrPmUJ):

> I have never said it cannot be done. I have always said it should not be done, and I have explained why.
> 
> "Useful" is not an argument for a feature. All features are useful; otherwise they would not be features. The issue is whether the feature justifies its cost. That is a judgement call, and my judgement is, no. Running compilers and linkers, doing megabytes of I/O, and creating temporary binary files is not a justifiable expense for running a small program. For large programs, the amortization is even more in favor of not doing this.
> 
> I am firmly against adding a feature to Go that encourages abuse of resources.
> 
> If you want the feature, use gorun or an equivalent wrapper program. That's what it's for.  I think believe [sic] gorun is a mistake, but I'm not stopping you from installing it and using it as you see fit.

### gorun

What's this gorun that has Rob up in arms?  The Ubuntu Linux distribution produced the [gorun utility](https://wiki.ubuntu.com/gorun) to solve the shebang issue.  As listed in the homepage, gorun will:

* write files under a safe directory in `$TMPDIR` (or `/tmp`), so that the actual script location isn't touched (may be read-only)
* avoid races between parallel executions of the same file
* automatically clean up old compiled files that remain unused for some time (without races)
* replace the process rather than using a child
* pass arguments to the compiled application properly
* handle well `GOROOT`, `GOROOT_FINAL` and the location of the toolchain

Importantly, gorun will also strip the shebang line from the executed file before running it via `go run`.

Because gorun is a separate tool, and not included in the Go distribution, it must first be installed via `go get launchpad.net/gorun`.  Once it's installed, you can make use of it quite simply:

```
#!/usr/bin/gorun

package main

func main() {
    println("Hello world!")
}
```

This file is now executable on the command line (after setting the execution bit via `chmod +x`, of course).  However, it's also technically invalid Go. This means you can't compile the same code via `go build` that you would run with gorun.

As we said earlier, this solution also requires that your users install gorun, making it less portable than ideal.

As of March 1st, 2015, [gorun does not compile on OS X](https://bugs.launchpad.net/gorun/+bug/1176275), but it looks like this will be resolved, soon.
{:.warning}

### Mimic the shebang via Bash

In UNIX systems, the shebang line is handled by the family of `exec` functions inside the fundamental libc.  The `exec` functions use what's referred to as magic numbers in the first bits of a file to determine what kind of executable it is.  The characters `#!` are simply hard coded as a magic number that corresponds to being run as a script.

But we can make clever use of the fact that `//` is both a comment in Go, *and* interpreted as a valid part of a file path in Bash, in order to trick Bash into running the file for us instead of `exec`.  This isn't technically a shebang line, but serves the same purpose quite nicely.  The resulting file looks like such:

```
//usr/bin/go run $0 $@; exit $?
package main

func main() {
    println("Hello world!")
}
```

This file is *both* a valid Bash script and valid Go source code.  When Bash runs this file, it executes `/usr/bin/go run $0 $@`, and then immediately exits.  If the file were named "/path/to/foo.go", and we ran it like `./foo.go bar baz`, then Bash would execute `/usr/bin/go run ./foo.go "bar" "baz"`.

At this point, the file is parsed by `go run`, which ignores the first line as a comment, and dutifully compiles and executes the rest of the file.

This trick doesn't require a separately installed command, and is relatively portable.  However, all is not roses:  

* The file must end in `.go`, since `go run` will happily ignore it otherwise.
* This only works when executed via a shell.  It won't work if executed directly by other programs.

### Trade-offs

Sadly, unless Rob changes his mind and modifies the Go spec, there is no one-size-fits-all solution in sight.  Both of the recipes above have their drawbacks, and choosing which one to use is entirely about your specific needs and tastes.

