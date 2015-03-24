---
---

``` go
// Reverse elements
for i, j := 0, len(elements)-1; i < j; i, j = i+1, j-1 {
    elements[i], elements[j] = elements[j], elements[i]
}
```
