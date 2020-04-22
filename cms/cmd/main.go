package main

// each time change the GOPATH in .zchrc just redo 'source .zshrc'
// https://sourabhbajaj.com/mac-setup/Go/
import (
	"go-projects/cms"
	"os"
)

func main() {
	p := &cms.Page{
		Title:   "Hello!",
		Content: "Body of Page", // last item must have comma
	}

	cms.Tmpl.ExecuteTemplate(os.Stdout, "index", p)
}
