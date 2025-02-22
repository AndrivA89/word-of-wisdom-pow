// Package domain entity for citation
package domain

// Citation - domain entity for work with citation.
type Citation struct {
	text   string
	author string
}

// NewCitation - create new citation entity.
func NewCitation(text, author string) *Citation {
	return &Citation{
		text:   text,
		author: author,
	}
}

// GetText - return text of citation.
func (c *Citation) GetText() string {
	return c.text
}

// GetAuthor - return author of citation.
func (c *Citation) GetAuthor() string {
	return c.author
}
