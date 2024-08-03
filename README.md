# Go Cookbook

Source for [the Go Cookbook](http://golangcookbook.com), a
community built collection of practical
recipes and contributions for real world Golang development, supported by
[SuperOrbital](https://superorbit.al).

### Contributing

The Go Cookbook is supported by
[SuperOrbital](https://superorbit.al), but built by the
community, so your contributions are very welcome.  Just send
a pull request for any changes or additions.

#### Adding a New Recipe

Recipes are generated from the [_data/chapters.yml
file](https://github.com/golangcookbook/golangcookbook.github.io/blob/master/_data/chapters.yml),
which is used to build [the index
page](https://github.com/golangcookbook/golangcookbook.github.io/blob/master/index.md).
The `chapters.yml` file format is:

```
- title: Strings
  recipes:
  - title: Concatenating Strings
    path: /chapters/strings/concatenation
    wip: true
  - title: Detecting a Substring
    path: /chapters/strings/detecting
  - title: Detecting All Substrings
    path: false
```

Specifying `wip: true` puts "[Work in progress]" in front of
the recipe in the index.  Setting `path: false` causes the
recipe to be listed without a link.

#### Recipe Format

Recipe files have a couple of required properties, specified
in the preamble like such:

``` yaml
---
title: Processing a String One Word or Character at a Time
question: Given a string, how do I break it into words or characters and process each one in turn?
---
```

Also, to show example code, recipes can make use of the `{%
include example.html example="filename" %}` helper. This will
include both `filename.go` as example source code, and
`filename.expected.escaped` as the output of that code.  If
you're curious, [here's the definition for that
helper](https://github.com/golangcookbook/golangcookbook.github.io/blob/master/_includes/example.html).
The `filename.expected.escaped` files are automatically
generated by the `./run_tests` script, described below.

### Run locally

This project uses Docker for local development.  Once you
have Docker installed:

``` console
$ ./run server
$ open http://localhost:4000
```

### Run the tests

``` console
$ ./run tests
```

The `run_tests` script not only runs the tests, but also
ensures that the `.expected`, `.go.escaped`, and
`.expected.escaped` files are in place.  These files are then
used in the recipe templates to ensure all examples are up to
date and correct.

### License

[![Creative Commons
License](https://i.creativecommons.org/l/by-nc-sa/4.0/88x31.png)](http://creativecommons.org/licenses/by-nc-sa/4.0/)

The Go Cookbook copyright
[SuperOrbital](https://superorbit.al), and is licensed under
a [Creative Commons Attribution-NonCommercial-ShareAlike 4.0
International
License](http://creativecommons.org/licenses/by-nc-sa/4.0/).
