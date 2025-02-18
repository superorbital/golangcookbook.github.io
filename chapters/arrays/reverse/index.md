---
title: Reversing an Array
question: How do I reverse an array?
---

Reversing an array is fairly straightforward in Go, due to multiple return
values.  Simply loop through the first half of the array, swapping each element
in turn with its mirror counterpart:

{%include example.html example="simple"%}

There's an even more concise implementation that uses the `,` operator
inside the Go for loop:

{%include example.html example="terse"%}

Note that both of these `reverse` functions take slices, which are passed by reference.
This is good for performance, but means we're actually modifying the
passed in array, this might not be the intended behavior.  In other words, the `reverse`
method could be written like such:

{%include example.html example="alternate"%}

Now `reverse` doesn't return another reference to the slice, which makes it
much more clear that it's modifying the slice in-place.  If you want to return a
modified copy of the slice without altering the original, you need to create
and return a new slice manually:

{%include example.html example="copy"%}

This ensures the original array remains untouched.
You can also simplify the algorithm by directly copying elements into a new slice:

{%include example.html example="copy_simple"%}
