package main

import (
	"fmt"

	"github.com/trenchesdeveloper/toll-calculator/types"
)

type Aggregator interface {
	AggregateDistance(types.Distance) error
}

type Storer interface {
	StoreDistance(types.Distance) error
}

type InvoiceAggregator struct {
	store Storer
}

func NewInvoiceAggregator(store Storer) *InvoiceAggregator {
	return &InvoiceAggregator{
		store: store,
	}
}

func (ia *InvoiceAggregator) AggregateDistance(distance types.Distance) error {
	fmt.Println("Processing and aggregating distance")
	return ia.store.StoreDistance(distance)

}
