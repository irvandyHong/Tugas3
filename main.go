package main

import (
	"crypto/rand"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"math/big"
	"net/http"
)

var PORT = ":8080"

type Data struct {
	Wind  int `json:"wind"`
	Water int `json:"water"`
}

var refreshed int

func main() {
	refreshed = 0
	http.HandleFunc("/", outputHandler)
	http.ListenAndServe(PORT, nil)
	//updateJson(10, 10)
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
	refreshed++
	wind := genRandNum(0, 100)
	water := genRandNum(0, 100)
	if refreshed%2 != 0 {
		updateJson(int(water), int(wind))
		myvar := map[string]interface{}{"MyVar": wind, "wind": windStatus(int(wind)), "MyVar2": water, "water": waterStatus(int(water))}
		outputHTML(w, "index.html", myvar)
	}

}
func updateJson(water, wind int) {
	tmp := Data{
		Wind:  wind,
		Water: water,
	}
	file, err := json.Marshal(tmp)
	if err != nil {
		panic(err)
	}
	_ = ioutil.WriteFile("data.json", file, 0644)
	// tmp := `{"wind":wind,"water":water}`
	// jsonData := []byte(tmp)
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
func windStatus(tmp int) string {
	if tmp <= 6 {
		return "aman"
	} else if tmp >= 7 && tmp <= 15 {
		return "siaga"
	} else {
		return "bahaya"
	}
}
func waterStatus(tmp int) string {
	if tmp <= 5 {
		return "aman"
	} else if tmp >= 6 && tmp <= 8 {
		return "siaga"
	} else {
		return "bahaya"
	}
}
