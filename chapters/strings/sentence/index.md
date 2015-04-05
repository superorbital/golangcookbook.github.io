---
title: Turning an Array into a Sentence
layout: default
---

## {{ page.title }}
{:.recipe}

How can I convert an array of words into a readable sentence of options?
{:.question}

If you're content with joining each of the strings in a given array using a consistent separator, then `strings.Join()` is the tool for you.  `Join()` takes an array of strings, and a final string that acts as the separator.  If we pass in `", "` as our separator, then we come pretty close to producing a sentence:

```
package main

import (
  "fmt"
  "strings"
)

func main() {
  a := []string{"this", "that", "the other"}
  fmt.Println(strings.Join(a, ", "))
}
```

When run, outputs:

```
$ go run test.go
this, that, the other
```

Sometimes this is good enough, but we often want the resulting string to read more like and English sentence, offering all or some of the options.  To produce this, we can create a `ToSentence` method like such:

```
package main

import (
  "fmt"
  "strings"
)

func ToSentence(words []string, andOrOr string) string {
  l := len(words)
  wordsForSentence := make([]string, l)
  copy(wordsForSentence, words)
  wordsForSentence[l-1] = andOrOr + " " + wordsForSentence[l-1]
  return strings.Join(wordsForSentence, ", ")
}

func main() {
  a := []string{"this", "that", "the other"}
  fmt.Println(ToSentence(a, "or"))
  fmt.Println(ToSentence(a, "and"))
}
```

The above will print out:

```
$ go run test.go
this, that, or the other
this, that, and the other
```

...which is much more friendly when presented to our users.
