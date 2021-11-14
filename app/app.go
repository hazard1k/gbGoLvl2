package app

import (
	"context"
	"golvl2/app/domain"
	"sync"
)

type App struct {
	repos domain.Repositories
}

type HTTPServer interface {
	Start(r domain.Repositories)
	Stop()
}

func (a *App) Serve(ctx context.Context, wg *sync.WaitGroup, hs HTTPServer) {
	defer wg.Done()
	hs.Start(a.repos)
	<-ctx.Done()
	hs.Stop()
}

func NewApp(r domain.Repositories) *App {

	return &App{
		repos: r,
	}
}
