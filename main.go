package main

import (
  "fmt"
  "log"
  "net/http"
  "runtime"
  "time"
)

func appHandler(w http.ResponseWriter, r *http.Request) {

  fmt.Println(time.Now(), "Hello from my new fresh server")

}

func main() {
//  http.HandleFunc("/", appHandler)
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<p>Welcome to Demo Application Changed Again!</p>")
		fmt.Fprintf(w, "<p>Server time:"+time.Now().Format("2006-01-02 15:04:05")+"</p>")
		fmt.Fprintf(w, "<p>Server OS:"+runtime.GOOS+"<p>")
	})

  log.Println("Started, serving on port 80")
  err := http.ListenAndServe(":80", nil)

  if err != nil {
    log.Fatal(err.Error())
  }
}
