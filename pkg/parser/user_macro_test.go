package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("user macros", func() {

	Context("inline macros", func() {

		It("inline macro empty", func() {
			source := "AAA hello:[]"
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: [][]interface{}{
					{
						types.StringElement{
							Content: "AAA ",
						},
						types.UserMacro{
							Kind:       types.InlineMacro,
							Name:       "hello",
							Value:      "",
							Attributes: types.ElementAttributes{},
							RawText:    "hello:[]",
						},
					},
				},
			}
			Expect(source).To(BecomeDocumentBlock(expected))
		})

		It("inline macro with attribute", func() {
			source := `AAA hello:[suffix="!!!!!"]`
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: [][]interface{}{
					{
						types.StringElement{
							Content: "AAA ",
						},
						types.UserMacro{
							Kind:  types.InlineMacro,
							Name:  "hello",
							Value: "",
							Attributes: types.ElementAttributes{
								"suffix": "!!!!!",
							},
							RawText: `hello:[suffix="!!!!!"]`,
						},
					},
				},
			}
			Expect(source).To(BecomeDocumentBlock(expected))
		})

		It("inline macro with value", func() {
			source := `AAA hello:JohnDoe[]`
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: [][]interface{}{
					{
						types.StringElement{
							Content: "AAA ",
						},
						types.UserMacro{
							Kind:       types.InlineMacro,
							Name:       "hello",
							Value:      "JohnDoe",
							Attributes: types.ElementAttributes{},
							RawText:    "hello:JohnDoe[]",
						},
					},
				},
			}
			Expect(source).To(BecomeDocumentBlock(expected))
		})

		It("inline user macro with value and attributes", func() {
			source := "repository: git:some/url.git[key1=value1,key2=value2]"
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: [][]interface{}{
					{
						types.StringElement{
							Content: "repository: ",
						},
						types.UserMacro{
							Kind:  types.InlineMacro,
							Name:  "git",
							Value: "some/url.git",
							Attributes: types.ElementAttributes{
								"key1": "value1",
								"key2": "value2",
							},
							RawText: "git:some/url.git[key1=value1,key2=value2]",
						},
					},
				},
			}
			Expect(source).To(BecomeDocumentBlock(expected))
		})
	})

	Context("user macros", func() {

		It("user macro block without value", func() {

			source := "git::[]"
			expected := types.UserMacro{
				Kind:       types.BlockMacro,
				Name:       "git",
				Value:      "",
				Attributes: types.ElementAttributes{},
				RawText:    "git::[]",
			}
			Expect(source).To(BecomeDocumentBlock(expected))
		})

		It("user block macro with value and attributes", func() {
			source := "git::some/url.git[key1=value1,key2=value2]"
			expected := types.UserMacro{
				Kind:  types.BlockMacro,
				Name:  "git",
				Value: "some/url.git",
				Attributes: types.ElementAttributes{
					"key1": "value1",
					"key2": "value2",
				},
				RawText: "git::some/url.git[key1=value1,key2=value2]",
			}
			Expect(source).To(BecomeDocumentBlock(expected))
		})

		It("user macro block with attribute", func() {
			source := `git::[key1="value1"]`
			expected := types.UserMacro{
				Kind:  types.BlockMacro,
				Name:  "git",
				Value: "",
				Attributes: types.ElementAttributes{
					"key1": "value1",
				},
				RawText: `git::[key1="value1"]`,
			}
			Expect(source).To(BecomeDocumentBlock(expected))
		})

		It("user macro block with value", func() {
			source := `git::some/url.git[]`
			expected := types.UserMacro{
				Kind:       types.BlockMacro,
				Name:       "git",
				Value:      "some/url.git",
				Attributes: types.ElementAttributes{},
				RawText:    "git::some/url.git[]",
			}
			Expect(source).To(BecomeDocumentBlock(expected))
		})

	})
})
