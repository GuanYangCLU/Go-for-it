package main

// no comma for each line of import, use ()
import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// here is package level scope

func main() {
	// := means declare and assign in same line
	res, err := http.Get("http://api.theysaidso.com/qod.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	// defer like finally as exception handler
	defer res.Body.Close()
	contents, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		// fmt.Printf("%T", err) print its type
		return
	}
	fmt.Println(string(contents))
}

// after run 'go build' will generate an executable file, use ./filename to execute it in any system
