package main

import (
	"crypto/rand"
	"fmt"
	"html/template"
	"math/big"
	"net/http"
)

var tmplString = `    // content of index.html
    {{define "index"}}
    {{.var1}} is equal to {{.var2}}
    {{end}}`
var i int

func main() {
	i = 0
	http.HandleFunc("/", outputHandler)
	http.ListenAndServe(":8080", nil)
}

func outputHTML(w http.ResponseWriter, filename string, data interface{}) {
	t, err := template.ParseFiles(filename)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if err := t.Execute(w, data); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
func outputHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("refreshed : ", i)
	i++
	// do whatever you need to do
	wind := genRandNum(0, 100)
	water := genRandNum(0, 100)
	myvar := map[string]interface{}{"MyVar": wind, "MyVar2": water}
	outputHTML(w, "index.html", myvar)
}

func genRandNum(min, max int64) int64 {
	// calculate the max we will be using
	bg := big.NewInt(max - min)

	// get big.Int between 0 and bg
	// in this case 0 to 20
	n, err := rand.Int(rand.Reader, bg)
	if err != nil {
		panic(err)
	}

	// add n to min to support the passed in range
	return n.Int64() + min
}
