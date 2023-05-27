package main

import (
	"golang.org/x/sync/errgroup"
	"log"
	"services/cmd/service/order/init"
)

func main() {
	r := errgroup.Group{}
	// start order-service
	r.Go(func() error {
		return init.StartOrderService()
	})
	// start customer-service
	//r.Go(func() error {
	//	return nil
	//})
	if err := r.Wait(); err != nil {
		log.Fatal(err)
	}
}
