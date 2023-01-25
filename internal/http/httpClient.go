package httpClient

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func HttpCall() {
	client := &http.Client{}
	request, err := http.NewRequest("POST", "https://httpbin.org/post", nil)
	if err != nil {
		log.Fatal(err)
	}
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body)
}
