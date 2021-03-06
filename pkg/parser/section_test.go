package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("sections - draft", func() {

	Context("valid sections", func() {

		It("header only", func() {
			source := "= a header"
			doctitle := []interface{}{
				types.StringElement{Content: "a header"},
			}
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.Section{
						Attributes: types.ElementAttributes{},
						Level:      0,
						Title:      doctitle,
						Elements:   []interface{}{},
					},
				},
			}
			Expect(source).To(BecomeDraftDocument(expected))
		})

		It("header with many spaces around content", func() {
			source := "= a header   "
			doctitle := []interface{}{
				types.StringElement{Content: "a header   "},
			}
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.Section{
						Attributes: types.ElementAttributes{},
						Level:      0,
						Title:      doctitle,
						Elements:   []interface{}{},
					},
				},
			}
			Expect(source).To(BecomeDraftDocument(expected))
		})

		It("header and paragraph", func() {
			source := `= a header

and a paragraph`

			doctitle := []interface{}{
				types.StringElement{Content: "a header"},
			}
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.Section{
						Attributes: types.ElementAttributes{},
						Level:      0,
						Title:      doctitle,
						Elements:   []interface{}{},
					},
					types.BlankLine{},
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "and a paragraph"},
							},
						},
					},
				},
			}
			Expect(source).To(BecomeDraftDocument(expected))
		})

		It("two sections with level 0", func() {
			source := `= a first header

= a second header`
			doctitle := []interface{}{
				types.StringElement{Content: "a first header"},
			}
			otherDoctitle := []interface{}{
				types.StringElement{Content: "a second header"},
			}

			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.Section{
						Attributes: types.ElementAttributes{},
						Level:      0,
						Title:      doctitle,
						Elements:   []interface{}{},
					},
					types.BlankLine{},
					types.Section{
						Attributes: types.ElementAttributes{},
						Level:      0,
						Title:      otherDoctitle,
						Elements:   []interface{}{},
					},
				},
			}
			Expect(source).To(BecomeDraftDocument(expected))
		})

		It("section level 1 alone", func() {
			source := `== section 1`
			section1Title := []interface{}{
				types.StringElement{Content: "section 1"},
			}
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.Section{
						Attributes: types.ElementAttributes{},
						Level:      1,
						Title:      section1Title,
						Elements:   []interface{}{},
					},
				},
			}
			Expect(source).To(BecomeDraftDocument(expected))
		})

		It("section level 1 with quoted text", func() {
			source := `==  *2 spaces and bold content*`
			sectionTitle := []interface{}{
				types.QuotedText{
					Kind: types.Bold,
					Elements: []interface{}{
						types.StringElement{Content: "2 spaces and bold content"},
					},
				},
			}
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.Section{
						Attributes: types.ElementAttributes{},
						Level:      1,
						Title:      sectionTitle,
						Elements:   []interface{}{},
					},
				},
			}
			Expect(source).To(BecomeDraftDocument(expected))
		})

		It("section level 0 with nested section level 1", func() {
			source := `= a header

== section 1`
			doctitle := []interface{}{
				types.StringElement{Content: "a header"},
			}
			section1Title := []interface{}{
				types.StringElement{Content: "section 1"},
			}
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.Section{
						Attributes: types.ElementAttributes{},
						Level:      0,
						Title:      doctitle,
						Elements:   []interface{}{},
					},
					types.BlankLine{},
					types.Section{
						Attributes: types.ElementAttributes{},
						Level:      1,
						Title:      section1Title,
						Elements:   []interface{}{},
					},
				},
			}
			Expect(source).To(BecomeDraftDocument(expected))
		})

		It("section level 0 with nested section level 2", func() {
			source := `= a header

=== section 2`
			doctitle := []interface{}{
				types.StringElement{Content: "a header"},
			}
			section2Title := []interface{}{
				types.StringElement{Content: "section 2"},
			}
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.Section{
						Attributes: types.ElementAttributes{},
						Level:      0,
						Title:      doctitle,
						Elements:   []interface{}{},
					},
					types.BlankLine{},
					types.Section{
						Attributes: types.ElementAttributes{},
						Level:      2,
						Title:      section2Title,
						Elements:   []interface{}{},
					},
				},
			}
			Expect(source).To(BecomeDraftDocument(expected))
		})

		It("section level 1 with immediate paragraph", func() {
			source := `== a title
and a paragraph`
			section1Title := []interface{}{
				types.StringElement{Content: "a title"},
			}
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.Section{
						Attributes: types.ElementAttributes{},
						Level:      1,
						Title:      section1Title,
						Elements:   []interface{}{},
					},
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "and a paragraph"},
							},
						},
					},
				},
			}
			Expect(source).To(BecomeDraftDocument(expected))
		})

		It("section level 1 with a paragraph separated by empty line", func() {
			source := `== a title
			
and a paragraph`
			section1Title := []interface{}{
				types.StringElement{Content: "a title"},
			}
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.Section{
						Attributes: types.ElementAttributes{},
						Level:      1,
						Title:      section1Title,
						Elements:   []interface{}{},
					},
					types.BlankLine{},
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "and a paragraph"},
							},
						},
					},
				},
			}
			Expect(source).To(BecomeDraftDocument(expected))
		})

		It("section level 1 with a paragraph separated by non-empty line", func() {
			source := "== a title\n    \nand a paragraph"
			section1Title := []interface{}{
				types.StringElement{Content: "a title"},
			}
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.Section{
						Attributes: types.ElementAttributes{},
						Level:      1,
						Title:      section1Title,
						Elements:   []interface{}{},
					},
					types.BlankLine{},
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "and a paragraph"},
							},
						},
					},
				},
			}
			Expect(source).To(BecomeDraftDocument(expected))
		})

		It("section levels 1, 2, 3, 2", func() {
			source := `= a header

== Section A
a paragraph

=== Section A.a
a paragraph

== Section B
a paragraph`
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.Section{
						Attributes: types.ElementAttributes{},
						Level:      0,
						Title: []interface{}{
							types.StringElement{Content: "a header"},
						},
						Elements: []interface{}{},
					},
					types.BlankLine{},
					types.Section{
						Attributes: types.ElementAttributes{},
						Level:      1,
						Title: []interface{}{
							types.StringElement{Content: "Section A"},
						},
						Elements: []interface{}{},
					},
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "a paragraph"},
							},
						},
					},
					types.BlankLine{},
					types.Section{
						Attributes: types.ElementAttributes{},
						Level:      2,
						Title: []interface{}{
							types.StringElement{Content: "Section A.a"},
						},
						Elements: []interface{}{},
					},
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "a paragraph"},
							},
						},
					},
					types.BlankLine{},
					types.Section{
						Attributes: types.ElementAttributes{},
						Level:      1,
						Title: []interface{}{
							types.StringElement{Content: "Section B"},
						},
						Elements: []interface{}{},
					},
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "a paragraph"},
							},
						},
					},
				},
			}
			Expect(source).To(BecomeDraftDocument(expected))
		})

		It("single section with custom IDs", func() {
			source := `[[custom_header]]
== a header`
			sectionTitle := []interface{}{
				types.StringElement{Content: "a header"},
			}
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.Section{
						Attributes: types.ElementAttributes{
							types.AttrID:       "custom_header",
							types.AttrCustomID: true,
						},
						Level:    1,
						Title:    sectionTitle,
						Elements: []interface{}{},
					},
				},
			}
			Expect(source).To(BecomeDraftDocument(expected))
		})

		It("multiple sections with custom IDs", func() {
			source := `[[custom_header]]
= a header

== Section F [[ignored]] [[foo]]

[[bar]]
== Section B
a paragraph`
			doctitle := []interface{}{
				types.StringElement{Content: "a header"},
			}
			fooTitle := []interface{}{
				types.StringElement{Content: "Section F "},
			}
			barTitle := []interface{}{
				types.StringElement{Content: "Section B"},
			}
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.Section{
						Attributes: types.ElementAttributes{
							types.AttrID:       "custom_header",
							types.AttrCustomID: true,
						},
						Level:    0,
						Title:    doctitle,
						Elements: []interface{}{},
					},
					types.BlankLine{},
					types.Section{
						Attributes: types.ElementAttributes{
							types.AttrID:       "foo",
							types.AttrCustomID: true,
						},
						Level:    1,
						Title:    fooTitle,
						Elements: []interface{}{},
					},
					types.BlankLine{},
					types.Section{
						Attributes: types.ElementAttributes{
							types.AttrID:       "bar",
							types.AttrCustomID: true,
						},
						Level:    1,
						Title:    barTitle,
						Elements: []interface{}{},
					},
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "a paragraph"},
							},
						},
					},
				},
			}
			Expect(source).To(BecomeDraftDocument(expected))
		})

		It("sections with same title", func() {
			source := `== section 1

== section 1`
			section1aTitle := []interface{}{
				types.StringElement{Content: "section 1"},
			}
			section1bTitle := []interface{}{
				types.StringElement{Content: "section 1"},
			}
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.Section{
						Attributes: types.ElementAttributes{},
						Level:      1,
						Title:      section1aTitle,
						Elements:   []interface{}{},
					},
					types.BlankLine{},
					types.Section{
						Attributes: types.ElementAttributes{},
						Level:      1,
						Title:      section1bTitle,
						Elements:   []interface{}{},
					},
				},
			}
			Expect(source).To(BecomeDraftDocument(expected))
		})

		It("section with link in title", func() {
			source := `== link to https://foo.bar
`
			section1aTitle := []interface{}{
				types.StringElement{Content: "link to "},
				types.InlineLink{
					Attributes: types.ElementAttributes{},
					Location: types.Location{
						Elements: []interface{}{
							types.StringElement{Content: "https://foo.bar"},
						},
					},
				},
			}
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.Section{
						Attributes: types.ElementAttributes{},
						Level:      1,
						Title:      section1aTitle,
						Elements:   []interface{}{},
					},
				},
			}
			Expect(source).To(BecomeDraftDocument(expected))
		})

		It("section 0, 1 and paragraph with bold quote", func() {

			source := `= a header
				
== section 1

a paragraph with *bold content*`

			title := []interface{}{
				types.StringElement{Content: "a header"},
			}
			section1Title := []interface{}{
				types.StringElement{Content: "section 1"},
			}
			expected := types.Document{
				Attributes: types.DocumentAttributes{},
				ElementReferences: types.ElementReferences{
					"_a_header":  title,
					"_section_1": section1Title,
				},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Section{
						Level: 0,
						Attributes: types.ElementAttributes{
							types.AttrID: "_a_header",
						},
						Title: title,
						Elements: []interface{}{
							types.Section{
								Level: 1,
								Title: section1Title,
								Attributes: types.ElementAttributes{
									types.AttrID: "_section_1",
								},
								Elements: []interface{}{
									types.Paragraph{
										Attributes: types.ElementAttributes{},
										Lines: [][]interface{}{
											{
												types.StringElement{Content: "a paragraph with "},
												types.QuotedText{
													Kind: types.Bold,
													Elements: []interface{}{
														types.StringElement{Content: "bold content"},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			}
			Expect(source).To(BecomeDocument(expected))
		})

	})

	Context("invalid sections", func() {
		It("header invalid - missing space", func() {
			source := "=a header"
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "=a header"},
							},
						},
					},
				}}
			Expect(source).To(BecomeDraftDocument(expected))
		})

		It("header invalid - header space", func() {
			source := " = a header with a prefix space"
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.LiteralBlock{
						Attributes: types.ElementAttributes{
							types.AttrKind:             types.Literal,
							types.AttrLiteralBlockType: types.LiteralBlockWithSpacesOnFirstLine,
						},
						Lines: []string{
							" = a header with a prefix space",
						},
					},
				},
			}
			Expect(source).To(BecomeDraftDocument(expected))
		})

		It("header with invalid section1", func() {
			source := `= a header

   == section with prefix space`
			title := []interface{}{
				types.StringElement{Content: "a header"},
			}
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.Section{
						Attributes: types.ElementAttributes{},
						Level:      0,
						Title:      title,
						Elements:   []interface{}{},
					},
					types.BlankLine{},
					types.LiteralBlock{
						Attributes: types.ElementAttributes{
							types.AttrKind:             types.Literal,
							types.AttrLiteralBlockType: types.LiteralBlockWithSpacesOnFirstLine,
						},
						Lines: []string{
							"   == section with prefix space",
						},
					},
				},
			}
			Expect(source).To(BecomeDraftDocument(expected))
		})

	})

	Context("unsupported section syntax", func() {

		It("should not fail with underlined title", func() {
			source := `Document Title
==============
Doc Writer <thedoc@asciidoctor.org>`
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: [][]interface{}{
							{
								types.StringElement{
									Content: "Document Title",
								},
							},
							{
								types.StringElement{
									Content: "==============",
								},
							},
							{
								types.StringElement{
									Content: "Doc Writer <thedoc@asciidoctor.org>",
								},
							},
						},
					},
				},
			}
			Expect(source).To(BecomeDraftDocument(expected))
		})
	})
})

var _ = Describe("sections - document", func() {

	Context("valid sections", func() {

		It("header only", func() {
			source := "= a header"
			doctitle := []interface{}{
				types.StringElement{Content: "a header"},
			}
			expected := types.Document{
				Attributes: types.DocumentAttributes{},
				ElementReferences: types.ElementReferences{
					"_a_header": doctitle,
				},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Section{
						Attributes: types.ElementAttributes{
							types.AttrID: "_a_header",
						},
						Level:    0,
						Title:    doctitle,
						Elements: []interface{}{},
					},
				},
			}
			Expect(source).To(BecomeDocument(expected))
		})

		It("header with many spaces around content", func() {
			source := "= a header   "
			doctitle := []interface{}{
				types.StringElement{Content: "a header   "},
			}
			expected := types.Document{
				Attributes: types.DocumentAttributes{},
				ElementReferences: types.ElementReferences{
					"_a_header": doctitle,
				},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Section{
						Attributes: types.ElementAttributes{
							types.AttrID: "_a_header",
						},
						Level:    0,
						Title:    doctitle,
						Elements: []interface{}{},
					},
				},
			}
			Expect(source).To(BecomeDocument(expected))
		})

		It("header and paragraph", func() {
			source := `= a header

and a paragraph`

			doctitle := []interface{}{
				types.StringElement{Content: "a header"},
			}
			expected := types.Document{
				Attributes: types.DocumentAttributes{},
				ElementReferences: types.ElementReferences{
					"_a_header": doctitle,
				},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Section{
						Attributes: types.ElementAttributes{
							types.AttrID: "_a_header",
						},
						Level: 0,
						Title: doctitle,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: [][]interface{}{
									{
										types.StringElement{Content: "and a paragraph"},
									},
								},
							},
						},
					},
				},
			}
			Expect(source).To(BecomeDocument(expected))
		})

		It("two sections with level 0", func() {
			source := `= a first header

= a second header`
			doctitle := []interface{}{
				types.StringElement{Content: "a first header"},
			}
			otherDoctitle := []interface{}{
				types.StringElement{Content: "a second header"},
			}

			expected := types.Document{
				Attributes: types.DocumentAttributes{},
				ElementReferences: types.ElementReferences{
					"_a_first_header":  doctitle,
					"_a_second_header": otherDoctitle,
				},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Section{
						Attributes: types.ElementAttributes{
							types.AttrID: "_a_first_header",
						},
						Level:    0,
						Title:    doctitle,
						Elements: []interface{}{},
					},
					types.Section{
						Attributes: types.ElementAttributes{
							types.AttrID: "_a_second_header",
						},
						Level:    0,
						Title:    otherDoctitle,
						Elements: []interface{}{},
					},
				},
			}
			Expect(source).To(BecomeDocument(expected))
		})

		It("section level 1 alone", func() {
			source := `== section 1`
			section1Title := []interface{}{
				types.StringElement{Content: "section 1"},
			}
			expected := types.Document{
				Attributes: types.DocumentAttributes{},
				ElementReferences: types.ElementReferences{
					"_section_1": section1Title,
				},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Section{
						Attributes: types.ElementAttributes{
							types.AttrID: "_section_1",
						},
						Level:    1,
						Title:    section1Title,
						Elements: []interface{}{},
					},
				},
			}
			Expect(source).To(BecomeDocument(expected))
		})

		It("section level 1 with quoted text", func() {
			source := `==  *2 spaces and bold content*`
			sectionTitle := []interface{}{
				types.QuotedText{
					Kind: types.Bold,
					Elements: []interface{}{
						types.StringElement{Content: "2 spaces and bold content"},
					},
				},
			}
			expected := types.Document{
				Attributes: types.DocumentAttributes{},
				ElementReferences: types.ElementReferences{
					"_2_spaces_and_bold_content": sectionTitle,
				},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Section{
						Attributes: types.ElementAttributes{
							types.AttrID: "_2_spaces_and_bold_content",
						},
						Level:    1,
						Title:    sectionTitle,
						Elements: []interface{}{},
					},
				},
			}
			Expect(source).To(BecomeDocument(expected))
		})

		It("section level 0 with nested section level 1", func() {
			source := `= a header

== section 1`
			doctitle := []interface{}{
				types.StringElement{Content: "a header"},
			}
			section1Title := []interface{}{
				types.StringElement{Content: "section 1"},
			}
			expected := types.Document{
				Attributes: types.DocumentAttributes{},
				ElementReferences: types.ElementReferences{
					"_a_header":  doctitle,
					"_section_1": section1Title,
				},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Section{
						Attributes: types.ElementAttributes{
							types.AttrID: "_a_header",
						},
						Level: 0,
						Title: doctitle,
						Elements: []interface{}{
							types.Section{
								Attributes: types.ElementAttributes{
									types.AttrID: "_section_1",
								},
								Level:    1,
								Title:    section1Title,
								Elements: []interface{}{},
							},
						},
					},
				},
			}
			Expect(source).To(BecomeDocument(expected))
		})

		It("section level 0 with nested section level 2", func() {
			source := `= a header

=== section 2`
			doctitle := []interface{}{
				types.StringElement{Content: "a header"},
			}
			section2Title := []interface{}{
				types.StringElement{Content: "section 2"},
			}
			expected := types.Document{
				Attributes: types.DocumentAttributes{},
				ElementReferences: types.ElementReferences{
					"_a_header":  doctitle,
					"_section_2": section2Title,
				},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Section{
						Attributes: types.ElementAttributes{
							types.AttrID: "_a_header",
						},
						Level: 0,
						Title: doctitle,
						Elements: []interface{}{
							types.Section{
								Attributes: types.ElementAttributes{
									types.AttrID: "_section_2",
								},
								Level:    2,
								Title:    section2Title,
								Elements: []interface{}{},
							},
						},
					},
				},
			}
			Expect(source).To(BecomeDocument(expected))
		})

		It("section level 1 with immediate paragraph", func() {
			source := `== a title
and a paragraph`
			section1Title := []interface{}{
				types.StringElement{Content: "a title"},
			}
			expected := types.Document{
				Attributes: types.DocumentAttributes{},
				ElementReferences: types.ElementReferences{
					"_a_title": section1Title,
				},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Section{
						Attributes: types.ElementAttributes{
							types.AttrID: "_a_title",
						},
						Level: 1,
						Title: section1Title,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: [][]interface{}{
									{
										types.StringElement{Content: "and a paragraph"},
									},
								},
							},
						},
					},
				},
			}
			Expect(source).To(BecomeDocument(expected))
		})

		It("section level 1 with a paragraph separated by empty line", func() {
			source := `== a title
			
and a paragraph`
			section1Title := []interface{}{
				types.StringElement{Content: "a title"},
			}
			expected := types.Document{
				Attributes: types.DocumentAttributes{},
				ElementReferences: types.ElementReferences{
					"_a_title": section1Title,
				},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Section{
						Attributes: types.ElementAttributes{
							types.AttrID: "_a_title",
						},
						Level: 1,
						Title: section1Title,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: [][]interface{}{
									{
										types.StringElement{Content: "and a paragraph"},
									},
								},
							},
						},
					},
				},
			}
			Expect(source).To(BecomeDocument(expected))
		})

		It("section level 1 with a paragraph separated by non-empty line", func() {
			source := "== a title\n    \nand a paragraph"
			section1Title := []interface{}{
				types.StringElement{Content: "a title"},
			}
			expected := types.Document{
				Attributes: types.DocumentAttributes{},
				ElementReferences: types.ElementReferences{
					"_a_title": section1Title,
				},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Section{
						Attributes: types.ElementAttributes{
							types.AttrID: "_a_title",
						},
						Level: 1,
						Title: section1Title,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: [][]interface{}{
									{
										types.StringElement{Content: "and a paragraph"},
									},
								},
							},
						},
					},
				},
			}
			Expect(source).To(BecomeDocument(expected))
		})

		It("section levels 1, 2, 3, 2", func() {
			source := `= a header

== Section A
a paragraph

=== Section A.a
a paragraph

== Section B
a paragraph`
			doctitle := []interface{}{
				types.StringElement{Content: "a header"},
			}
			sectionATitle := []interface{}{
				types.StringElement{Content: "Section A"},
			}
			sectionAaTitle := []interface{}{
				types.StringElement{Content: "Section A.a"},
			}
			sectionBTitle := []interface{}{
				types.StringElement{Content: "Section B"},
			}
			expected := types.Document{
				Attributes: types.DocumentAttributes{},
				ElementReferences: types.ElementReferences{
					"_a_header":    doctitle,
					"_section_a":   sectionATitle,
					"_section_a_a": sectionAaTitle,
					"_section_b":   sectionBTitle,
				},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Section{
						Attributes: types.ElementAttributes{
							types.AttrID: "_a_header",
						},
						Level: 0,
						Title: doctitle,
						Elements: []interface{}{
							types.Section{
								Attributes: types.ElementAttributes{
									types.AttrID: "_section_a",
								},
								Level: 1,
								Title: sectionATitle,
								Elements: []interface{}{
									types.Paragraph{
										Attributes: types.ElementAttributes{},
										Lines: [][]interface{}{
											{
												types.StringElement{Content: "a paragraph"},
											},
										},
									},
									types.Section{
										Attributes: types.ElementAttributes{
											types.AttrID: "_section_a_a",
										},
										Level: 2,
										Title: sectionAaTitle,
										Elements: []interface{}{
											types.Paragraph{
												Attributes: types.ElementAttributes{},
												Lines: [][]interface{}{
													{
														types.StringElement{Content: "a paragraph"},
													},
												},
											},
										},
									},
								},
							},
							types.Section{
								Attributes: types.ElementAttributes{
									types.AttrID: "_section_b",
								},
								Level: 1,
								Title: sectionBTitle,
								Elements: []interface{}{
									types.Paragraph{
										Attributes: types.ElementAttributes{},
										Lines: [][]interface{}{
											{
												types.StringElement{Content: "a paragraph"},
											},
										},
									},
								},
							},
						},
					},
				},
			}
			Expect(source).To(BecomeDocument(expected))
		})

		It("section levels 1, 2, 3, 3", func() {
			source := `= a header

== Section A
a paragraph

=== Section A.a
a paragraph

=== Section A.b
a paragraph`
			doctitle := []interface{}{
				types.StringElement{Content: "a header"},
			}
			sectionATitle := []interface{}{
				types.StringElement{Content: "Section A"},
			}
			sectionAaTitle := []interface{}{
				types.StringElement{Content: "Section A.a"},
			}
			sectionBTitle := []interface{}{
				types.StringElement{Content: "Section A.b"},
			}
			expected := types.Document{
				Attributes: types.DocumentAttributes{},
				ElementReferences: types.ElementReferences{
					"_a_header":    doctitle,
					"_section_a":   sectionATitle,
					"_section_a_a": sectionAaTitle,
					"_section_a_b": sectionBTitle,
				},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Section{
						Attributes: types.ElementAttributes{
							types.AttrID: "_a_header",
						},
						Level: 0,
						Title: doctitle,
						Elements: []interface{}{
							types.Section{
								Attributes: types.ElementAttributes{
									types.AttrID: "_section_a",
								},
								Level: 1,
								Title: sectionATitle,
								Elements: []interface{}{
									types.Paragraph{
										Attributes: types.ElementAttributes{},
										Lines: [][]interface{}{
											{
												types.StringElement{Content: "a paragraph"},
											},
										},
									},
									types.Section{
										Attributes: types.ElementAttributes{
											types.AttrID: "_section_a_a",
										},
										Level: 2,
										Title: sectionAaTitle,
										Elements: []interface{}{
											types.Paragraph{
												Attributes: types.ElementAttributes{},
												Lines: [][]interface{}{
													{
														types.StringElement{Content: "a paragraph"},
													},
												},
											},
										},
									},
									types.Section{
										Attributes: types.ElementAttributes{
											types.AttrID: "_section_a_b",
										},
										Level: 2,
										Title: sectionBTitle,
										Elements: []interface{}{
											types.Paragraph{
												Attributes: types.ElementAttributes{},
												Lines: [][]interface{}{
													{
														types.StringElement{Content: "a paragraph"},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			}
			Expect(source).To(BecomeDocument(expected))
		})

		It("section levels 1, 3, 4, 4", func() {
			source := `= a header

=== Section A
a paragraph

==== Section A.a
a paragraph

==== Section A.b
a paragraph`
			doctitle := []interface{}{
				types.StringElement{Content: "a header"},
			}
			sectionATitle := []interface{}{
				types.StringElement{Content: "Section A"},
			}
			sectionAaTitle := []interface{}{
				types.StringElement{Content: "Section A.a"},
			}
			sectionBTitle := []interface{}{
				types.StringElement{Content: "Section A.b"},
			}
			expected := types.Document{
				Attributes: types.DocumentAttributes{},
				ElementReferences: types.ElementReferences{
					"_a_header":    doctitle,
					"_section_a":   sectionATitle,
					"_section_a_a": sectionAaTitle,
					"_section_a_b": sectionBTitle,
				},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Section{
						Attributes: types.ElementAttributes{
							types.AttrID: "_a_header",
						},
						Level: 0,
						Title: doctitle,
						Elements: []interface{}{
							types.Section{
								Attributes: types.ElementAttributes{
									types.AttrID: "_section_a",
								},
								Level: 2,
								Title: sectionATitle,
								Elements: []interface{}{
									types.Paragraph{
										Attributes: types.ElementAttributes{},
										Lines: [][]interface{}{
											{
												types.StringElement{Content: "a paragraph"},
											},
										},
									},
									types.Section{
										Attributes: types.ElementAttributes{
											types.AttrID: "_section_a_a",
										},
										Level: 3,
										Title: sectionAaTitle,
										Elements: []interface{}{
											types.Paragraph{
												Attributes: types.ElementAttributes{},
												Lines: [][]interface{}{
													{
														types.StringElement{Content: "a paragraph"},
													},
												},
											},
										},
									},
									types.Section{
										Attributes: types.ElementAttributes{
											types.AttrID: "_section_a_b",
										},
										Level: 3,
										Title: sectionBTitle,
										Elements: []interface{}{
											types.Paragraph{
												Attributes: types.ElementAttributes{},
												Lines: [][]interface{}{
													{
														types.StringElement{Content: "a paragraph"},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			}
			Expect(source).To(BecomeDocument(expected))
		})

		It("single section with custom IDs", func() {
			source := `[[custom_header]]
== a header`
			sectionTitle := []interface{}{
				types.StringElement{Content: "a header"},
			}
			expected := types.Document{
				Attributes: types.DocumentAttributes{},
				ElementReferences: types.ElementReferences{
					"custom_header": sectionTitle,
				},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Section{
						Attributes: types.ElementAttributes{
							types.AttrID:       "custom_header",
							types.AttrCustomID: true,
						},
						Level:    1,
						Title:    sectionTitle,
						Elements: []interface{}{},
					},
				},
			}
			Expect(source).To(BecomeDocument(expected))
		})

		It("multiple sections with custom IDs", func() {
			source := `[[custom_header]]
= a header

== Section F [[ignored]] [[foo]]

[[bar]]
== Section B
a paragraph`
			doctitle := []interface{}{
				types.StringElement{Content: "a header"},
			}
			fooTitle := []interface{}{
				types.StringElement{Content: "Section F "},
			}
			barTitle := []interface{}{
				types.StringElement{Content: "Section B"},
			}
			expected := types.Document{
				Attributes: types.DocumentAttributes{},
				ElementReferences: types.ElementReferences{
					"custom_header": doctitle,
					"foo":           fooTitle,
					"bar":           barTitle,
				},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Section{
						Attributes: types.ElementAttributes{
							types.AttrID:       "custom_header",
							types.AttrCustomID: true,
						},
						Level: 0,
						Title: doctitle,
						Elements: []interface{}{
							types.Section{
								Attributes: types.ElementAttributes{
									types.AttrID:       "foo",
									types.AttrCustomID: true,
								},
								Level:    1,
								Title:    fooTitle,
								Elements: []interface{}{},
							},
							types.Section{
								Attributes: types.ElementAttributes{
									types.AttrID:       "bar",
									types.AttrCustomID: true,
								},
								Level: 1,
								Title: barTitle,
								Elements: []interface{}{
									types.Paragraph{
										Attributes: types.ElementAttributes{},
										Lines: [][]interface{}{
											{
												types.StringElement{Content: "a paragraph"},
											},
										},
									},
								},
							},
						},
					},
				},
			}
			Expect(source).To(BecomeDocument(expected))
		})

		It("sections with same title", func() {
			source := `== section 1

== section 1`
			section1aTitle := []interface{}{
				types.StringElement{Content: "section 1"},
			}
			section1bTitle := []interface{}{
				types.StringElement{Content: "section 1"},
			}

			expected := types.Document{
				Attributes: types.DocumentAttributes{},
				ElementReferences: types.ElementReferences{
					"_section_1":   section1aTitle,
					"_section_1_2": section1bTitle,
				},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Section{
						Attributes: types.ElementAttributes{
							types.AttrID: "_section_1",
						},
						Level:    1,
						Title:    section1aTitle,
						Elements: []interface{}{},
					},
					types.Section{
						Attributes: types.ElementAttributes{
							types.AttrID: "_section_1_2",
						},
						Level:    1,
						Title:    section1bTitle,
						Elements: []interface{}{},
					},
				},
			}
			Expect(source).To(BecomeDocument(expected))
		})

		It("section level 0 with nested section level 1 and custom ID prefix", func() {
			source := `= a header
:idprefix: custom_

== section 1`
			doctitle := []interface{}{
				types.StringElement{Content: "a header"},
			}
			section1Title := []interface{}{
				types.StringElement{Content: "section 1"},
			}
			expected := types.Document{
				Attributes: types.DocumentAttributes{
					types.AttrIDPrefix: "custom_",
				},
				ElementReferences: types.ElementReferences{
					"custom_a_header":  doctitle,
					"custom_section_1": section1Title,
				},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Section{
						Attributes: types.ElementAttributes{
							types.AttrID: "custom_a_header",
						},
						Level: 0,
						Title: doctitle,
						Elements: []interface{}{
							types.Section{
								Attributes: types.ElementAttributes{
									types.AttrID: "custom_section_1",
								},
								Level:    1,
								Title:    section1Title,
								Elements: []interface{}{},
							},
						},
					},
				},
			}
			Expect(source).To(BecomeDocument(expected))
		})

		It("section level 0 with nested sections level 1 and custom ID prefixes - with idprefix as doc attribute", func() {
			source := `= a header
:idprefix: custom1a_

== section 1a

:idprefix: custom1b_

== section 1b`
			doctitle := []interface{}{
				types.StringElement{Content: "a header"},
			}
			section1aTitle := []interface{}{
				types.StringElement{Content: "section 1a"},
			}
			section1bTitle := []interface{}{
				types.StringElement{Content: "section 1b"},
			}
			expected := types.Document{
				Attributes: types.DocumentAttributes{
					types.AttrIDPrefix: "custom1a_",
				},
				ElementReferences: types.ElementReferences{
					"custom1a_a_header":   doctitle,
					"custom1a_section_1a": section1aTitle,
					"custom1b_section_1b": section1bTitle,
				},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Section{
						Attributes: types.ElementAttributes{
							types.AttrID: "custom1a_a_header",
						},
						Level: 0,
						Title: doctitle,
						Elements: []interface{}{
							types.Section{
								Attributes: types.ElementAttributes{
									types.AttrID: "custom1a_section_1a",
								},
								Level:    1,
								Title:    section1aTitle,
								Elements: []interface{}{},
							},
							types.Section{
								Attributes: types.ElementAttributes{
									types.AttrID: "custom1b_section_1b",
								},
								Level:    1,
								Title:    section1bTitle,
								Elements: []interface{}{},
							},
						},
					},
				},
			}
			Expect(source).To(BecomeDocument(expected))
		})

		It("section level 0 with nested sections level 1 and custom ID prefixes - without idprefix as doc attribute", func() {
			source := `= a header

:idprefix: custom1a_

== section 1a

:idprefix: custom1b_

== section 1b`
			doctitle := []interface{}{
				types.StringElement{Content: "a header"},
			}
			section1aTitle := []interface{}{
				types.StringElement{Content: "section 1a"},
			}
			section1bTitle := []interface{}{
				types.StringElement{Content: "section 1b"},
			}
			expected := types.Document{
				Attributes: types.DocumentAttributes{},
				ElementReferences: types.ElementReferences{
					"_a_header":           doctitle,
					"custom1a_section_1a": section1aTitle,
					"custom1b_section_1b": section1bTitle,
				},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Section{
						Attributes: types.ElementAttributes{
							types.AttrID: "_a_header",
						},
						Level: 0,
						Title: doctitle,
						Elements: []interface{}{
							types.Section{
								Attributes: types.ElementAttributes{
									types.AttrID: "custom1a_section_1a",
								},
								Level:    1,
								Title:    section1aTitle,
								Elements: []interface{}{},
							},
							types.Section{
								Attributes: types.ElementAttributes{
									types.AttrID: "custom1b_section_1b",
								},
								Level:    1,
								Title:    section1bTitle,
								Elements: []interface{}{},
							},
						},
					},
				},
			}
			Expect(source).To(BecomeDocument(expected))
		})
	})

	Context("invalid sections", func() {

		It("header invalid - too many spaces", func() {
			source := "======= a header"
			expected := types.Document{
				Attributes:         types.DocumentAttributes{},
				ElementReferences:  types.ElementReferences{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "======= a header"},
							},
						},
					},
				}}
			Expect(source).To(BecomeDocument(expected))
		})

		It("header invalid - missing space", func() {
			source := "=a header"
			expected := types.Document{
				Attributes:         types.DocumentAttributes{},
				ElementReferences:  types.ElementReferences{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "=a header"},
							},
						},
					},
				}}
			Expect(source).To(BecomeDocument(expected))
		})

		It("header invalid - header space", func() {
			source := " = a header with a prefix space"
			expected := types.Document{
				Attributes:         types.DocumentAttributes{},
				ElementReferences:  types.ElementReferences{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.LiteralBlock{
						Attributes: types.ElementAttributes{
							types.AttrKind:             types.Literal,
							types.AttrLiteralBlockType: types.LiteralBlockWithSpacesOnFirstLine,
						},
						Lines: []string{
							" = a header with a prefix space",
						},
					},
				},
			}
			Expect(source).To(BecomeDocument(expected))
		})

		It("header with invalid section1", func() {
			source := `= a header

 == section with prefix space`
			title := []interface{}{
				types.StringElement{Content: "a header"},
			}
			expected := types.Document{
				Attributes: types.DocumentAttributes{},
				ElementReferences: types.ElementReferences{
					"_a_header": title,
				}, Footnotes: types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Section{
						Attributes: types.ElementAttributes{
							types.AttrID: "_a_header",
						},
						Level: 0,
						Title: title,
						Elements: []interface{}{
							types.LiteralBlock{
								Attributes: types.ElementAttributes{
									types.AttrKind:             types.Literal,
									types.AttrLiteralBlockType: types.LiteralBlockWithSpacesOnFirstLine,
								},
								Lines: []string{
									" == section with prefix space",
								},
							},
						},
					},
				},
			}
			Expect(source).To(BecomeDocument(expected))
		})

	})

	Context("unsupported section syntax", func() {

		It("should not fail with underlined title", func() {
			source := `Document Title
==============
Doc Writer <thedoc@asciidoctor.org>`
			expected := types.Document{
				Attributes:         types.DocumentAttributes{},
				ElementReferences:  types.ElementReferences{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: [][]interface{}{
							{
								types.StringElement{
									Content: "Document Title",
								},
							},
							{
								types.StringElement{
									Content: "==============",
								},
							},
							{
								types.StringElement{
									Content: "Doc Writer <thedoc@asciidoctor.org>",
								},
							},
						},
					},
				},
			}
			Expect(source).To(BecomeDocument(expected))
		})
	})
})
