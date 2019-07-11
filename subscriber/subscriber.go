package subscriber

import (
	"log"

	"context"

	balance "github.com/akililab/balance/proto/balance"
)

// Balance struct definition
type Balance struct{}

// Handle this handles subscriptions
func (e *Balance) Handle(ctx context.Context, msg *balance.Message) error {
	log.Print("Handler Received message: ", msg.Say)
	return nil
}

// Handler for subsctiber
func Handler(ctx context.Context, msg *balance.Message) error {
	log.Print("Function Received message: ", msg.Say)
	return nil
}
