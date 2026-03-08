package posts

import (
	"embed"
)

// go:embed articles/*.md projects/*.md
var markdowns embed.FS

func get_md_from_slug(slug string) {

}

// This code will receive a md and parse it using Goldmark.
func posts_handler(slug string, markdowns []byte) {
	// get markdown whose metadata matches slug

}
