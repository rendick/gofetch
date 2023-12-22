package weather

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

var (
	Reset = "\033[0m"
	Red   = "\033[31m"
)

func Weather() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://wttr.in/London?format=%t%20|%20%h%20|%20%P%20|%20%w", nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(Red+"Weather: "+Reset+"%s\n", bodyText)
}
