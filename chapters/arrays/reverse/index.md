---
title: Reversing an Array
question: How do I reverse an array?
---

Reversing an array is fairly straightforward in Go, due to multiple return
values.  Simply loop through the first half of the array, swapping each element
in turn with its mirror counterpart:

{%include example.html example="simple"%}

There's an even more terse implementation which makes use of the `,` operator
inside the Go for loop:

{%include example.html example="terse"%}

Note that both of these `reverse` functions take slices, which are passed by
value.  This is good for performance, but means we're actually modifying the
passed in array, leading to unexpected results.  In other words, the `reverse`
method could be written like such:

{%include example.html example="alternate"%}

Now `reverse` doesn't return another reference to the slice, which makes it
much more clear that it's modifying the slice in-place.  If we want to return a
modified copy of the slice, we must do so manually:

{%include example.html example="copy"%}

As you can see in the output above, the original array is untouched.

Since we're not modifying the array in-place, we can also simplify our
algorithm buy doing a more straight-forward element copy:

{%include example.html example="copy_simple"%}

