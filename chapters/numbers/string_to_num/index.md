---
title: Converting a string to number
question: How do I convert a string to number if possible in Go?
---

Go comes with an in-built library called strconv which makes this task very easy. You can convert a string to an integer and specify a base as well. You also have the flexibility to define the size of bits into which the number should fit into.

{% include example.html example="string_to_int" %}

The same library also contains functions such as ParseUint and ParseFloat which as the name suggest convert a string to unsigned integer and floating point number respectively.

{% include example.html example="string_to_float" %}
