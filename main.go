package main

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime"
	"time"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

func main() {
	url := "http://hello-world-demo.ir-e1.cloudhub.io/api/helloworld"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<p>Welcome to Demo Application Changed!</p>")
		fmt.Fprintf(w, "<p>Server time:"+time.Now().Format("2006-01-02 15:04:05")+"</p>")
		fmt.Fprintf(w, "<p>Server OS:"+runtime.GOOS+"<p>")
		fmt.Fprintf(w, "<p>Mulesoft Response:"+getMulesoft(url)+"</p>")
	})
	http.ListenAndServe(":80", nil)
}

func getMulesoft(url string) string {
	r, err := myClient.Get(url)
	if err != nil {
		fmt.Printf(err.Error())
	}
	defer r.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)

	return buf.String()
}
