package cms

import (
	"html/template"
)

// Capitalized first letter var (like Tmpl) is an exported variable
var Tmpl = template.Must(template.ParseGlob("../templates/*"))

type Page struct {
	Title   string
	Content string
	Posts   []*Post
}

type Post struct {
	Title string
}
