package strings_test

import (
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Strings", func() {
	Context("Substrings", func() {
		It("can detect substrings", func() {
			Expect(strings.Contains("Assume", "Ass")).To(Equal(true))

			Expect(strings.ContainsAny("team", "i")).To(Equal(false))
			Expect(strings.ContainsAny("Hello", "aeiou&y")).To(Equal(true))

			Expect(strings.HasPrefix("Hello World", "Hello")).To(Equal(true))
			Expect(strings.HasSuffix("Hello World", "World")).To(Equal(true))
		})

		It("can access by index", func() {
			s := "Hello World"
			Expect(s[:5]).To(Equal("Hello"))
			Expect(s[:len(s)-6]).To(Equal("Hello"))
			Expect(s[len(s)-5 : len(s)]).To(Equal("World"))
			Expect(s[1 : len(s)-1]).To(Equal("ello Worl"))
		})

		It("can split a string into words", func() {
			s := "Hello World"
			Expect(strings.Fields(s)).To(Equal([]string{"Hello", "World"}))
			Expect(strings.Split(s, "o")).To(Equal([]string{"Hell", " W", "rld"}))
		})

		It("can remove unwanted bits", func() {
			hello := " Hello "
			Expect(strings.Trim(hello, " ")).To(Equal("Hello"))
			Expect(strings.TrimSpace(hello)).To(Equal("Hello"))

			Expect(strings.TrimLeft(hello, " ")).To(Equal("Hello "))
			Expect(strings.TrimRight(hello, " ")).To(Equal(" Hello"))

			Expect(strings.TrimPrefix("FOO: bar", "FOO: ")).To(Equal("bar"))
			Expect(strings.TrimPrefix("bar", "FOO: ")).To(Equal("bar"))
		})

		It("can replace substrings", func() {
			Expect(strings.Replace("Hello World", "l", "L", 2)).To(Equal("HeLLo World"))
			Expect(strings.Replace("Hello World", "l", "L", -1)).To(Equal("HeLLo WorLd"))

			r := strings.NewReplacer("<", "&lt;", ">", "&gt;")
			Expect(r.Replace("<Hello> <World>")).To(Equal("&lt;Hello&gt; &lt;World&gt;"))
		})
	})

	It("can make a sentence", func() {
		a := []string{"this", "that", "the other"}
		Expect(strings.Join(a, ", ")).To(Equal("this, that, the other"))

		a[len(a)-1] = "and " + a[len(a)-1]
		Expect(strings.Join(a, ", ")).To(Equal("this, that, and the other"))
	})

	It("can titlize a string", func() {
		Expect(strings.Title("hello to the world!")).To(Equal("Hello To The World!"))
	})

	It("can properly titlize a string", func() {
		properTitle := func(input string) string {
			words := strings.Fields(input)
			output := make([]string, len(words))
			smallwords := "a an on the to"

			for index, word := range words {
				if strings.Contains(smallwords, word) {
					output[index] = word
				} else {
					output[index] = strings.Title(word)
				}
			}
			return strings.Join(output, " ")
		}

		s := "hello to the world!"
		Expect(properTitle(s)).To(Equal("Hello to the World!"))
	})
})
