---
title: Processing a String One Word or Character at a Time
question: Given a string, how do I break it into words or characters and process each one in turn?
---

### Each Character

Because of Go's built in support for Unicode "runes", processing a string one character at a time is quite straightforward.  Simply iterate over the `range` of that string:

{% include example.html example="each_char" %}

### Each Word

Processing a string one word at a time is a bit more involved, and depends on your specific needs.  If you're fine with the unsophisticated approach of cutting the string into words based on whitespace, then you're in luck - `strings.Fields` was built just for you:

{% include example.html example="words" %}

### Without Punctuation

However, most applications will need a more grammatically tolerant approach, where punctuation is taken into account.  Here we have two options.  We can either make use of a `strings.Replacer`, which we generate via the `strings.NewReplacer` function:

{% include example.html example="without_punctuation" %}

Or we can achieve a bit more clarity by making use of `strings.Map`:

{% include example.html example="without_punctuation_using_map" %}

### Special Separators

There are other situations where you'd want to split a string based on a separator other than whitespace.  The UNIX `/etc/passwd` file, for example, contains lines of tokens separated by colons.  Splitting each line into the relevant pieces is easy in Go, with the `strings.Split` function, which is a more generic form of `strings.Fields`:

{% include example.html example="separator" %}
