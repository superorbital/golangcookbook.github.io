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

		It("can replace substrings", func() {
			Expect(strings.Replace("Hello World", "l", "L", 2)).To(Equal("HeLLo World"))
			Expect(strings.Replace("Hello World", "l", "L", -1)).To(Equal("HeLLo WorLd"))

			r := strings.NewReplacer("<", "&lt;", ">", "&gt;")
			Expect(r.Replace("<Hello> <World>")).To(Equal("&lt;Hello&gt; &lt;World&gt;"))
		})
	})

})
