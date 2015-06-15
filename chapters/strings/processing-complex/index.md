---
title: Processing a String by Complex Separators and Patterns
question: Given a string with complex and variable separators, how do I break it into words?
---

Splitting strings based on a single separator character is fairly simple in Go, but what about more tenacious strings?  If a string uses multiple separators that should be treated equally, then `regexp.Split` is the tool for you:

{% include example.html example="regexp_separator" %}

However, some strings are more complex, still -- involving tokens wrapped in complex patterns.  For these situations, `regexp.FindAllStringSubmatch` combined with regexp capture groups will do what you need.

`regexp.FindAllStringSubmatch` will search a string for the regexp, and will return tuples of the matched string and any submatches found via capture groups.

{% include example.html example="regexp_match" %}
