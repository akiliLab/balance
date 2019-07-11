package handler

import (
	"context"

	"github.com/micro/go-log"

	balance "github.com/akililab/balance/proto/balance"
)

// Balance struct defining Bank Balance
type Balance struct{}

// ManageBalance is a single request handler called via client.Call or the generated client code
func (e *Balance) ManageBalance(ctx context.Context, req *balance.Request, rsp *balance.Response) error {
	log.Log("Received Balance.ManageBalance request")
	rsp.Completed = true
	return nil
}
