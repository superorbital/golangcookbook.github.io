---
title: Properly Capitalizing a Title
question: How do I convert a string into a properly capitalized title?
---

### Welcome To The Dollhouse?

Golang comes packaged with a good number of very useful string manipulation helpers in the `strings` package.  Unfortunately, the `strings.Title()` method is clever but naive in implementation.   Let's look at the source for `strings.Title()` (this is from the 1.4.1 version of Go):

``` go
// Title returns a copy of the string s with all Unicode letters that 
// begin words mapped to their title case.
//
// BUG: The rule Title uses for word boundaries does not handle Unicode
// punctuation properly.
func Title(s string) string {
	// Use a closure here to remember state.
	// Hackish but effective. Depends on Map scanning in order and calling
	// the closure once per rune.
	prev := ' '
	return Map(
		func(r rune) rune {
			if isSeparator(prev) {
				prev = r
				return unicode.ToTitle(r)
			}
			prev = r
			return r
		},
		s)
}
```

`strings.Map()` iterates over each character of a string, replacing it with whatever character is returned from the `func(r rune) rune` method supplied by the caller.  `strings.Title()` makes use of this in a clever way, by recording the current character as `prev`, so that the next run can tell whether it's at a word boundary.

The flaw in this implementation is that it doesn't account for small words. For example:  

{% include example.html example="simple" %}

### Welcome _to the_ Dollhouse!

Our English professors wouldn't be happy with the results.  Small words, such as "a" and "the" shouldn't be capitalized in a title.  In order to _properly_ capitalize a string, we have to do a bit more work.  

{% include example.html example="proper" %}

In the implementation above, we first split the string into an array of words, loop through them, and only capitalize those who don't match our hard-coded list of small words.  Finally, we join the words up into a single string, and return the result.  

Go doesn't have anything like an `array.Includes(element)` predicate, so we get a bit clever when determining if the word is a small word by making use of spaces as word boundaries.

As you can see in the output above, the printed title is grammatically correct.
