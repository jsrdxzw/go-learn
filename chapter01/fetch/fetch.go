package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const HttpsProtocol = "https://"

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, HttpsProtocol) {
			url = HttpsProtocol + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch err %v", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		// to avoid allocating memory, we can use io.copy
		//io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("get body: %s", b)
	}
}
