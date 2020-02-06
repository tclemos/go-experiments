package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", capitalize)
	go http.ListenAndServe(":8082", nil)

	reqBody := `{
		"this": "is",
		"a":    "complex json",
		"object": {
			"with": "sub properties",
			"to":   "be capitalized",
			"even": {
				"with" : "multiple sub structures"
			}
		},
		"did": "you get it?"
	}`

	b := []byte(reqBody)
	if res, err := http.Post("http://localhost:8082/", "application/json", bytes.NewBuffer(b)); err != nil {
		fmt.Println(err)
	} else {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
		}
		defer res.Body.Close()
		fmt.Println("response:", string(body))
	}

}

func capitalize(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error reading body"))
	}
	defer r.Body.Close()

	fmt.Println("request:", string(body))

	o := make(map[string]interface{}, 0)
	if err := json.Unmarshal(body, &o); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid json body"))
	}

	capitalizeMap(o)

	res, _ := json.MarshalIndent(o, "", "\t")
	w.Write(res)
}

func capitalizeMap(m map[string]interface{}) {
	for k, i := range m {

		switch v := i.(type) {
		default:
			continue
		case string:
			pp := strings.Split(v, " ")

			for i, p := range pp {
				pp[i] = strings.Title(p)
			}

			m[k] = strings.Join(pp, " ")
		case map[string]interface{}:
			capitalizeMap(v)
		}
	}
}
