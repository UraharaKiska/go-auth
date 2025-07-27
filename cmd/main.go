package main

import (
	"context"

	"log"

	"github.com/UraharaKiska/go-auth/internal/app"
)


func main() {

	ctx := context.Background()
	a, err := app.NewApp(ctx)
	// log.Printf("App :%v", a)
	if err != nil {
		log.Fatalf("failed to init app%v: ", err)
	}
	err = a.Run()
	if err != nil {
		log.Fatalf("failed to init app%v: ", err)
	}
}
