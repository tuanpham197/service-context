package main

import (
	"log"

	sctx "github.com/tuanpham197/service-context"
)

func main() {
	const compId = "foo"

	serviceCtx := sctx.NewServiceContext(
		sctx.WithName("simple-component"),
		sctx.WithComponent(NewSimpleComponent(compId)),
	)

	if err := serviceCtx.Load(); err != nil {
		log.Fatal(err)
	}

	type CanGetValue interface {
		GetValue() string
	}

	comp := serviceCtx.MustGet(compId).(CanGetValue)

	log.Println(comp.GetValue())

	_ = serviceCtx.Stop()
}
