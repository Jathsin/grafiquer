package types

import "html/template"

type Post struct {
	Title       string
	Slug        string // endpoint
	Parent      string // category
	Content     template.HTML
	Description string
	Order       int      // to sort posts
	Headers     []string // section names/ table of contents
	Metadata    Metadata
}

type Metadata struct {
	// SEO fields
	Title                     string
	Meta_description          string // HTML tag
	Meta_property_title       string // Open Graph, social media
	Meta_property_description string // Social sharing
	Meta_Og_URL               string // Canonical URL string
}
