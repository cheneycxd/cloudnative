package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/pprof"
	"os"

	"github.com/golang/glog"
)

func main() {
	flag.Set("v", "4")
	flag.Set("log_dir", "./log")
	flag.Parse()
	glog.V(2).Info("Starting http server...")

	// 	http.HandleFunc("/", rootHandler)
	// 	http.HandleFunc("/healthz", healthz)
	// 	err := http.ListenAndServe(":8001", nil)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 设置多路复用处理函数
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/healthz", healthz)
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	err := http.ListenAndServe(":8001", mux)
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
