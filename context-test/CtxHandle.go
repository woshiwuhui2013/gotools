package CtxPkg;

import (
	"context"
	"fmt"
)

func SubRoute(ctx context.Context)  {



	fmt.Println("sub sub routine")
	select {
	case <-ctx.Done():
	   fmt.Println("sub cancel")

	}



}
func CtxHello()  {
	parent := context.Background()
	ctx, cacel := context.WithCancel(parent)


	go SubRoute(ctx)


	defer cacel()
}