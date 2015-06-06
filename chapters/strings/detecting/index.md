---
title: Detecting a Substring
question: How do I determine if a string contains another string or pattern?
---

Go prides itself on being a low level language with a rich set of functionality included by default via the standard packages.  The `strings` package is a great example of this.

### Detecting Substrings

* `strings.Contains`
* `strings.ContainsAny`
* `strings.HasPrefix`
* `strings.HasSuffix`

{% include example.html example="substrings" %}

### Detecting Patterns

{% include example.html example="regexp" %}

### Finding Matched Strings

{% include example.html example="returned_regexp" %}

### Detecting Patterns that Include Special Characters

{% include example.html example="safe_regexp" %}

### Precompiling Regular Expressions

* `regexp.Compile`
* `regexp.MustCompile`

{% include example.html example="precompiled_regexp" %}

