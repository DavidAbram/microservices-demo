package main

import (
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		url := fmt.Sprintf("http://micro-date:3000/")

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Fatal("NewRequest: ", err)
			return
		}

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal("Do: ", err)
			return
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		query := strings.Replace(html.EscapeString(r.URL.Path), "/", "", -1)

		if query == "" {
			fmt.Fprintf(w, "Hello, sir, today is %s", string(body))
		} else {
			fmt.Fprintf(w, "Hello, %s, today is %s", query, string(body))
		}

	})

	log.Fatal(http.ListenAndServe(":1337", nil))

}
