package parser

import (
	"strconv"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// rearrangeSections moves elements into section to obtain a hierarchical document instead of a flat thing
func rearrangeSections(blocks []interface{}) (types.Document, error) {

	// use same logic as with list items:
	// only append a child section to her parent section when
	// a sibling or higher level section is processed.

	log.Debugf("rearranging sections in %d blocks...", len(blocks))
	tle := make([]interface{}, 0, len(blocks)) // top-level elements
	sections := make([]types.Section, 0, 6)    // the path to the current section (eg: []{section-level0, section-level1, etc.})
	elementRefs := types.ElementReferences{}
	footnotes := types.Footnotes{}
	footnoteRefs := types.FootnoteReferences{}
	var previous *types.Section // the current "parent" section
	for _, element := range blocks {
		if e, ok := element.(types.Section); ok {
			// avoid duplicate IDs in sections
			referenceSection(e, elementRefs)
			if previous == nil { // set first parent
				log.Debugf("setting section with title %v as a top-level element", e.Title)
				sections = append(sections, e)
			} else if e.Level > previous.Level { // add new level
				log.Debugf("adding section with title %v as the first section at level %d", e.Title, e.Level)
				sections = append(sections, e)
			} else { // replace at the deepest level
				sections = pruneSections(sections, e.Level)
				if len(sections) > 0 && sections[0].Level == e.Level {
					log.Debugf("moving section with title %v as a new top-level element", e.Title)
					tle = append(tle, sections[0])
					sections = make([]types.Section, 0, 6)
				}
				log.Debugf("adding section with title %v as another section at level %d", e.Title, e.Level)
				sections = append(sections, e)
				// if len(sections) == 1 { // we have new top-level element
				// 	log.Debugf("setting section with title %v as secondary top-level", e.Title)
				// 	tle = append(tle, &e)
				// } else {
				// 	log.Debugf("adding section with title %v as child of section at level %d", e.Title, (len(sections) - 2))
				// 	sections[len(sections)-2].AddElement(e) // attach to parent
				// }
			}
			previous = &e // pointer to new current parent
		} else {
			if previous == nil {
				// log.Debugf("adding element of type %T as a top-level element", element)
				tle = append(tle, element)
			} else {
				parentSection := &(sections[len(sections)-1])
				// log.Debugf("adding element of type %T as a child of section with level %d", element, parentSection.Level)
				(*parentSection).AddElement(element)
			}
		}
		// also collect footnotes
		if e, ok := element.(types.FootnotesContainer); ok {
			// log.Debugf("collecting footnotes on element of type %T", element)
			f, fr, err := e.Footnotes()
			if err != nil {
				return types.Document{}, errors.Wrap(err, "unable to collect footnotes in document")
			}
			footnotes = append(footnotes, f...)
			for k, v := range fr {
				footnoteRefs[k] = v
			}
		}
	}
	// process the remaining sections
	sections = pruneSections(sections, 1)
	if len(sections) > 0 {
		tle = append(tle, sections[0])
	}

	return types.Document{
		Attributes:         types.DocumentAttributes{},
		Elements:           tle,
		ElementReferences:  elementRefs,
		Footnotes:          footnotes,
		FootnoteReferences: footnoteRefs,
	}, nil
}

func referenceSection(e types.Section, elementRefs types.ElementReferences) {
	id := e.Attributes.GetAsString(types.AttrID)
	for i := 1; ; i++ {
		var key string
		if i == 1 {
			key = id
		} else {
			key = id + "_" + strconv.Itoa(i)
		}
		if _, found := elementRefs[key]; !found {
			elementRefs[key] = e.Title
			// override the element id
			e.Attributes[types.AttrID] = key
			break
		}
	}
	elementRefs[e.Attributes.GetAsString(types.AttrID)] = e.Title
}

func pruneSections(sections []types.Section, level int) []types.Section {
	if len(sections) > 0 && level > 0 { // && level < len(sections) {
		log.Debugf("pruning the section path with %d level(s) of deep", len(sections))
		// add the last list(s) as children of their parent, in reverse order,
		// because we copy the value, not the pointers
		cut := len(sections)
		for i := len(sections) - 1; i > 0 && sections[i].Level >= level; i-- {
			parentSection := &(sections[i-1])
			log.Debugf("appending section at depth %d (%v) to the last element of the parent section (%v)", i, sections[i].Title, parentSection.Title)
			(*parentSection).AddElement(sections[i])
			cut = i
		}
		// also, prune the pointers to the remaining sublists
		sections := sections[0:cut]
		log.Debugf("sections list has now %d top-level elements", len(sections))
		return sections
	}
	return sections
}
