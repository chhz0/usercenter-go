package main

import (
	"context"

	"github.com/chhz0/usercenter-go/internal/usercenter"
)

func main() {
	if err := usercenter.NewUserCenter().Run(context.Background()); err != nil {
		panic(err)
	}
}
