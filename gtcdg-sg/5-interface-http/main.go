package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	// Method 1 -- Todd McLeod
	res, err := http.Get("http://google.com")

	if err != nil {
		log.Fatal(err)
	}

	page, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", page)

	// // METHOD 2
	// // var resp *http.Response
	// // var err error
	//g
	// resp, err := http.Get("http://google.com")
	// if err != nil {
	// 	os.Exit(1)
	// }

	// fmt.Println(resp)
}

// func (r *http.Response) pretty() {
// 	// pretty print json response
// 	// https://stackoverflow.com/questions/19038598/how-can-i-pretty-print-json-using-go/42426889#42426889
// 	var out io.Writer
// 	enc := json.NewEncoder(out)
// 	enc.SetIndent("", "    ")
// 	err := enc.Encode(r)
// 	if err != nil {
// 		panic(err)
// 	}
// }
