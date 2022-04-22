package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/golang/glog"
)

func main() {
	flag.Set("v", "4")
	flag.Parse()
	glog.V(2).Info("Starting http server...")

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/healthz", healthz)
	err := http.ListenAndServe(":8001", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func healthz(w http.ResponseWriter, r *http.Request) {
	fmt.Println("entering healthz handler")
	glog.Info("entering healthz handler")
	setHeaderVersion(w, r)
	io.WriteString(w, "200\n")
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("entering root handler")
	glog.Info("entering root handler")
	setHeaderVersion(w, r)
	io.WriteString(w, "===================Details of the http request header:============\n")
	for k, v := range r.Header {
		io.WriteString(w, fmt.Sprintf("%s=%s\n", k, v))
	}

}

func setHeaderVersion(w http.ResponseWriter, r *http.Request) {
	ver := os.Getenv("VERSION")
	glog.Info(fmt.Sprintf("VERSION is %s\n", ver))
	w.Header().Add("VERSION", ver)
	w.WriteHeader(http.StatusOK)
	glog.Info(fmt.Sprintf("Remote Addr is %s , http status is %d", r.RemoteAddr, http.StatusOK))
}
