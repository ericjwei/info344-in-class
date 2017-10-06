package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"

	"github.com/ericjwei/info344-a17/info344-in-class/zipsvr/handlers"
	"github.com/ericjwei/info344-a17/info344-in-class/zipsvr/models"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	w.Header().Add("Content-Type", "text/plain")
	w.Header().Add("Access-Control-Allow-Origin", "*") // allows the access a different local host for security
	fmt.Fprintf(w, "Hello %s!", name)
	// http.Error(w, "you faild", http.StatusForbidden)
	return
}

func memoryHandler(w http.ResponseWriter, r *http.Request) {
	runtime.GC() //
	stats := &runtime.MemStats{}
	runtime.ReadMemStats(stats)
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(stats)
}

const zipsPath = "/zips/"

func main() {
	addr := os.Getenv("ADDR")
	if len(addr) == 0 {
		addr = ":80"
	}
	zips, err := models.LoadZips("zips.csv")
	if err != nil {
		log.Fatalf("error loading zips: %v", err)
	}
	log.Printf("loaded %d zips", len(zips))

	cityIndex := models.ZipIndex{}
	for _, z := range zips {
		cityLower := strings.ToLower(z.City)
		cityIndex[cityLower] = append(cityIndex[cityLower], z)
	}

	// fmt.Println("Hello World!")
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler) //add / after hello if you want to allow any following path
	mux.HandleFunc("/memory", memoryHandler)

	cityHandler := &handlers.CityHandler{
		Index:      cityIndex,
		PathPrefix: zipsPath,
	}
	mux.Handle(zipsPath, cityHandler)

	fmt.Printf("server is listening at http://%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
