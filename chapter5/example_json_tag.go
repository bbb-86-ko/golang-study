package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type MyStruct struct {
	Title    string `json:"title"`
	Internal string `json:"-"`
	Value    int64  `json:",omitempty"`
	ID       int64  `json:"id,string"`
}

func main() {
	var mystruct MyStruct
	b, err := json.Marshal(mystruct)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(b))
}
