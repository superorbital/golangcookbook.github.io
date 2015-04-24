---
title: Accessing Substrings
question: How do I access parts of a given string?
---

Go allows you to access parts of a string via the normal indexing primitives you're accustomed to using with arrays and slices.  In particular, it allows you to access `mystring[n:m]`.  Either `n` or `m` can be omitted, which implies from the beginning or to the end, respectively.  There's also the primitive `len(s)`, which gives you the number of characters in the string.

Here are some examples:

{% include example.html example="index" %}

Accessing strings by character, word or token are covered in the [Processing a String One Word or Character at a Time](/chapters/strings/processing) recipe.

