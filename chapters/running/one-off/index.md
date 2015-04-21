---
title: Running a one-off File
question: What's the easiest way to run a single Go file?
---

First steps are often both fundamental and overlooked.

Go's unusual packaging and compilation process can be daunting and confusing to a new developer.  Unlike most compiled languages, Go actually comes with a very easy way of running a single file.  Simply execute the program with `go run filename.go`:

{% include example.html example="hello_world" %}

This allows Go to act as though it were a more dynamic language, and is useful for development and exploration.  However, Go's compiler is so incredibly fast, that it's possible to use this technique in a production setting:

```
$ time go run empty.go

real	0m0.103s
user	0m0.058s
sys	0m0.033s
```

Throughout this book, you'll see many Go programs.  Each one is available as a separate file in this book, and each example includes the `go run` command required to run the file along side the full output.
