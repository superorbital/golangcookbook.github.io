---
title: Converting a string to number
question: How do I convert a string to number if possible in Go?
---

Go comes with a core package called `strconv` which makes this task very
easy. You can convert a string to an integer and specify a base as well. You
also have the flexibility to define the size of bits into which the number
should fit into.

{% include example.html example="string_to_int" %}

The `strconv` package also contains functions such as `ParseUint` and
`ParseFloat` which convert a string to unsigned integer and floating point
numbers respectively.

{% include example.html example="string_to_float" %}
