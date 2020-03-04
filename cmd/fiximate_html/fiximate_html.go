package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/daominah/quickfixdev"
)

func main() {
	resp, err := http.DefaultClient.Get(
		"https://fiximate.fixtrading.org/en/FIX.5.0SP2_EP252/fields_sorted_by_tagnum.html")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fields, err := quickfixdev.ParseFIXFieldsHtml(string(body))
	if err != nil {
		log.Fatal(err)
	}
	beauty, err := json.MarshalIndent(fields, "", "    ")
	fmt.Println(string(beauty))
}
