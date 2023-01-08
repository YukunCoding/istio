package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/pprof"
	"os"
	"strings"
)

func main() {
	sm := http.NewServeMux()
	sm.HandleFunc("/debug/pprof/", pprof.Index)
	sm.HandleFunc("/debug/pprof/profile", pprof.Profile)
	sm.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	sm.HandleFunc("/debug/pprof/trace", pprof.Trace)
	// 1 read request header and write to response
	sm.HandleFunc("/writeHeader", writeHeaderToResponse)
	// 2 fetch current system variable and write response
	sm.HandleFunc("/fetchVariable", fetchOSVariable)
	// 3 Print Client IP, HTTP RESPONSE CODE TO console
	sm.HandleFunc("/printClientIp", fetchClientIP)
	// 4 WHEN you visit at localhost/healthz, you shou be response 200
	sm.HandleFunc("/healthz", healthz)

	if err := http.ListenAndServe(":80", sm); err != nil {
		log.Fatalf("I am sorry, Start server Failed!: %s", err.Error())
	}
}

// 1 read request header and write to response
func writeHeaderToResponse(w http.ResponseWriter, r *http.Request) {
	for key, value := range r.Header {
		for _, v := range value {
			fmt.Printf("K: %s, V: %s\n", key, v)
			w.Header().Set(key, v)
		}
	}
}

// 2 fetch current system variable and write response
func fetchOSVariable(w http.ResponseWriter, r *http.Request) {
	os.Setenv("VERSION", "talk is cheap, show me the code.")
	version := os.Getenv("VERSION")
	w.Header().Set("VERSION", version)
}

// 3 Print Client IP, HTTP RESPONSE CODE TO console
func fetchClientIP(w http.ResponseWriter, r *http.Request) {
	ip := r.Header.Get("X-Real-IP")
	if ip == "" {
		ip = strings.Split(r.RemoteAddr, ":")[0]
	}
	println(r.RemoteAddr)
	log.Printf("client IP: %s", ip)
}

// 4 WHEN you visit at localhost/healthz, you shou be response 200
func healthz(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "200")
}
