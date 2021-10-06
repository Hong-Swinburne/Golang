package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"


	"github.com/golang/glog"
)

func main() {
	flag.Set("v", "4")
	glog.V(2).Info("Starting http server...")

	// create mux to access socket exclusively
	mux := http.NewServeMux()
	// register handlers for mux
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/healthz", healthz)
	fmt.Println("successfully registered all handlers")

	err := http.ListenAndServe(":80", mux)
	if err != nil {
		log.Fatal(err)
	}


}

func healthz(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "200\n")
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("entering root handler")

	io.WriteString(w, "===================Write details of the http request header to response header:============\n")
	for k, v := range r.Header {
		io.WriteString(w, fmt.Sprintf("%s=%s\n", k, v))
	}

	// write VERSION of environment variable into the response header
	out, err := exec.Command("go", "version").Output()
	if err != nil {
		log.Fatal(err)
	}
	io.WriteString(w, fmt.Sprintf("GO-Version=%s\n", out))

	//retrieve request ip
	ip := r.URL.Query().Get("RemoteAddr")
	if ip != "" {
		fmt.Sprint("Clinet IP is [%s]\n", ip)
	} else {
		fmt.Sprint("Can't retrieve client IP\n")
	}

	//retrieve Http status code
	// status_code = w.Header.Get("StatusCode")
	// if status_code != "" {
	// 	fmt.Sprint("Http status code is [%s]\n", status_code)
	// } else {
	// 	fmt.Sprint("Can't retrieve Http status code\n")
	}

}
