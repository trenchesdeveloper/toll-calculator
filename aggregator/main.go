package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"

	"github.com/trenchesdeveloper/toll-calculator/types"
)

func main() {
	listenAddR := flag.String("listenAddress", ":3000", "The address to listen on for HTTP requests.")
	flag.Parse()
	store := NewMemoryStore()
	svc := NewInvoiceAggregator(store)
	makeHTTPTransport(*listenAddR, svc)

	svc.AggregateDistance(types.Distance{})
}

func makeHTTPTransport(listenAddr string, svc Aggregator) {
	fmt.Println("HTTP transport running on: ", listenAddr)
	http.HandleFunc("/aggregate", handleAggregation(svc))

	http.ListenAndServe(listenAddr, nil)
}

func handleAggregation(svc Aggregator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var distance types.Distance

		if err := json.NewDecoder(r.Body).Decode(&distance); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := svc.AggregateDistance(distance); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

}
