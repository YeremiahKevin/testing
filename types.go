package main

import "time"

type InsertRequestParameter struct {
	RequestID int64        `json:"request_id"`
	Data      []InsertData `json:"data"`
}

type InsertData struct {
	ID           int64   `json:"id"`
	Customer     string  `json:"customer"`
	Quantity     int     `json:"quantity"`
	Price        float64 `json:"price"`
	Timestamp    string  `json:"timestamp"`
	NewTimestamp time.Time
}
