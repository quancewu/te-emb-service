package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"te-redis-service/controllers"
	"te-redis-service/initalizers"
	"te-redis-service/models"
)

func init() {
	initalizers.LoadEnvVariables()
	// initalizers.ConnectToSQLITE()
	initalizers.ConnectToSqliteTimeseries()
	// initalizers.SyncDatabase()
	initalizers.ConnectToRedis()
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	subscriber, err := controllers.NewSubscriber()
	if err != nil {
		log.Fatalf("Failed to create subscriber: %v", err)
	}

	subscriber.RegisterHandler("datas", func(m models.Message) error {
		if err := controllers.InsertAmsData(m); err != nil {
			log.Printf("%v", err)
			return nil
		}
		return nil
	}, models.MessageFilter{
		Types: []models.MessageType{models.TypeNotification},
	})

	var wg sync.WaitGroup
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	// Start subscriber
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := subscriber.Subscribe(ctx, "datas"); err != nil {
			log.Printf("Subscriber error: %v", err)
		}
	}()

	<-sigChan
	log.Println("Shutting down...")
	cancel()
	wg.Wait()
}
