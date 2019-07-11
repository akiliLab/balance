package handler

import (
	"context"

	"github.com/micro/go-log"

	balance "github.com/akililab/balance/proto/balance"
)

// Balance struct defining Bank Balance
type Balance struct{}

// Check is a single request handler called via client.Call or the generated client code
func (e *Balance) Check(ctx context.Context, req *balance.Request, rsp *balance.Response) error {
	log.Log("Received Balance.ManageBalance request")
	rsp.Completed = true
	return nil
}
