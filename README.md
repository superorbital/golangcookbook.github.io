# Go Cookbook

Source for [the Go Cookbook](http://golangcookbook.com).

### Run locally

```
$ bundle install
$ jekyll serve -B
$ open http://localhost:4000
```

### Run the tests

```
$ ./run_tests
```

...or...

```
$ ./run_tests chapters/strings/reverse/test_reverse_by_character.go
```

The `run_tests` script not only runs the tests, but also ensures that the `.expected`, `.go.escaped`, and `.expected.escaped` files are in place.  These files are then used in the chapter templates to ensure all examples are up to date and correct.

### License

[![Creative Commons License](https://i.creativecommons.org/l/by-nc-sa/4.0/88x31.png)](http://creativecommons.org/licenses/by-nc-sa/4.0/)

The Go Cookbook is licensed under a [Creative Commons Attribution-NonCommercial-ShareAlike 4.0 International License](http://creativecommons.org/licenses/by-nc-sa/4.0/).


