package main

import (
	"log"
	"log-service/data"
	"time"
)

type RPCServer struct {
	logRepo data.LogRepository
}

type RPCPayload struct {
	Name string
	Data string
}

func (r *RPCServer) LogInfo(payload RPCPayload, resp *string) error {
	logEntry := data.LogEntry{
		Name:      payload.Name,
		Data:      payload.Data,
		CreatedAt: time.Now(),
	}

	err := r.logRepo.Insert(logEntry)
	if err != nil {
		log.Println("error writing to Mongo via RPC", err)
		return err
	}

	*resp = "Processed payload via RPC: " + payload.Name
	return nil
}
