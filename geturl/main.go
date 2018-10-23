package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func getURL(url string, c chan []string) error {
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
		return err
	}

	c <- []string{url, string(body)}

	return nil
}

func main() {
	adrs := []string{
		"http://google.com",
		"http://seznam.cz",
		"http://root.cz",
		"http://abclinuxu.cz",
		"http://echo24.cz",
	}

	c := make(chan []string)

	for _, adr := range adrs {
		go getURL(adr, c)
	}

	for i := 0; i < len(adrs); i++ {
		select {
		case msg := <-c:
			fmt.Println(msg[0])
			fmt.Println("==============")
			fmt.Println(msg[1])
		}
	}

	close(c)
}
