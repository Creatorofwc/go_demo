package main

import (
  "fmt"
  "encoding/json"
  "net/http"
  "time"
  "runtime"
)

func healthCheck(w http.ResponseWriter, r *http.Request) {
    status := map[string]string{"status": "ok"}
    jsonResponse, err := json.Marshal(status)

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(jsonResponse)
}
func appHandler(w http.ResponseWriter, r *http.Request) {

  fmt.Println(time.Now(), "Hello from my new fresh server")

}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<p>Welcome to Demo Application Changed Again!</p>")
		fmt.Fprintf(w, "<p>Server time:"+time.Now().Format("2006-01-02 15:04:05")+"</p>")
		fmt.Fprintf(w, "<p>Server OS:"+runtime.GOOS+"<p>")
	}) 
    mux.HandleFunc("/healthcheck", healthCheck)

    err := http.ListenAndServe(":80", mux)
    if err != nil {
        panic(err)
    }
}
