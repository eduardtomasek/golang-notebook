package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func getURL(url string, c chan []string, e chan error) {
	resp, err := http.Get(url)

	if err != nil {
		e <- err
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		e <- err
		return
	}

	c <- []string{url, string(body)}

	return
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	adrs := []string{
		"http://google.com",
		"http://seznam.cz",
		"http://root.cz",
		"http://abclinuxu.cz",
		"http://echo24.cz",
		"http://nonexistingpage.xx",
	}

	c := make(chan []string)
	e := make(chan error)

	for _, adr := range adrs {
		go getURL(adr, c, e)
	}

	for i := 0; i < len(adrs); i++ {
		select {
		case msg := <-c:
			name := strings.Split(msg[0][7:], ".")[0]
			fmt.Println(name)

			f, err := os.Create(name + ".txt")
			check(err)

			_, err = f.WriteString(msg[1])
			check(err)
			f.Sync()
		case emsg := <-e:
			fmt.Println(emsg)
		}
	}

	close(c)
	close(e)
}
