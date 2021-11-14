package main

import (
	"context"
	"golvl2/api/handlers"
	"golvl2/api/routers/gorilla"
	"golvl2/api/server"
	"golvl2/app"
	"golvl2/db/memory"
	"os"
	"os/signal"
	"sync"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)

	db := memory.NewStore()

	a := app.NewApp(db)

	h := handlers.NewHandlers(db)

	r := gorilla.NewRouter(h)

	s := server.NewServer(":8000", r)

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go a.Serve(ctx, wg, s)

	<-ctx.Done()
	cancel()
	wg.Wait()
}
