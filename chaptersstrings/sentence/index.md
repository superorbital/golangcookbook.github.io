---
title: Turning an Array into a Sentence
question: How can I convert an array of words into a readable sentence of options?
---

If you're content with joining each of the strings in a given array using a consistent separator, then `strings.Join()` is the tool for you.  `Join()` takes an array of strings, and a final string that acts as the separator.  If we pass in `", "` as our separator, then we come pretty close to producing a sentence:

{% include example.html example="join" %}

Sometimes this is good enough, but we often want the resulting string to read more like and English sentence, offering all or some of the options.  To produce this, we can create a `ToSentence` method like such:

{% include example.html example="sentence" %}

...which is much more friendly when presented to our users.
