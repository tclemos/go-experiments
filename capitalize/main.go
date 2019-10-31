package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type NameObj struct {
	Name string `json:"name"`
}

func main() {
	http.HandleFunc("/", capitalize)
	go http.ListenAndServe(":8082", nil)
	o := NameObj{
		Name: "john doe",
	}
	b, _ := json.Marshal(o)
	if res, err := http.Post("http://localhost:8082/", "application/json", bytes.NewBuffer(b)); err != nil {
		fmt.Println(err)
	} else {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
		}
		defer res.Body.Close()
		fmt.Println("request body:", string(body))
	}

}

func capitalize(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error reading body"))
	}
	defer r.Body.Close()

	fmt.Println("request body:", string(body))

	o := new(NameObj)
	if err := json.Unmarshal(body, o); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid json body"))
	}

	pp := strings.Split(o.Name, " ")

	for i, p := range pp {
		pp[i] = fmt.Sprintf("%s%s", strings.ToUpper(p[0:1]), p[1:])
	}

	o.Name = strings.Join(pp, " ")

	res, _ := json.Marshal(o)
	w.Write(res)
}
