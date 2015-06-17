---
title: Reversing a String by Word or Character
question: How do I reverse a string by word or character?
---

Reversing a string by word is very similar to the process to reverse an array.  The difference is that strings are immutable in Go.  Therefor, we must first convert the string to a mutable array of runes (`[]rune`), perform the reverse operation on that, and then re-cast to a string.

{% include example.html example="reverse_by_character" %}

Reversing a string by word is a similar process.  First, we convert the string into an array of strings where each entry is a word.  Next we apply the normal reverse loop to that array. Finally, we smush the results back together into a string that we can return to the caller.

{% include example.html example="reverse_by_word" %}

