package main

import (
	"fmt"

	"github.com/trenchesdeveloper/toll-calculator/types"
)

type MemoryStore struct{}

func (s *MemoryStore) StoreDistance(distance types.Distance) error {
	fmt.Println("Storing distance")
	return nil
}


func NewMemoryStore() *MemoryStore {
	return &MemoryStore{}
}