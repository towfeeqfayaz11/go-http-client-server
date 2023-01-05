package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	resp1, err := c.Get("http://localhost:3333")
	if err != nil {
		fmt.Printf("Error %s", err)
		return
	}
	resp2, err := c.Get("http://localhost:3333/hello")
	if err != nil {
		fmt.Printf("Error %s", err)
		return
	}
	defer resp1.Body.Close()
	defer resp2.Body.Close()
	body1, err := ioutil.ReadAll(resp1.Body)
	body2, err := ioutil.ReadAll(resp2.Body)
	fmt.Printf("Body of from / endpoint : %s", body1)
	fmt.Printf("Body of from /hello endpoint : %s", body2)

}
