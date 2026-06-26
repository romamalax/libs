package graceful

import (
	"context"
	"fmt"
	"statter/pkg/logger"
	"time"
)

type Service interface {
	Start(context.Context) error
	Stop(context.Context) error
	Name() string
}

type Graceful struct {
	services []Service
	log      *logger.Logger
}

func New(log *logger.Logger, services ...Service) *Graceful {
	return &Graceful{
		log:      log,
		services: services,
	}
}

func (g *Graceful) Start(ctx context.Context) error {
	var t time.Time
	for _, service := range g.services {
		t = time.Now()
		if err := service.Start(ctx); err != nil {
			return err
		}
		g.log.Info(fmt.Sprintf("%s started", service.Name()), "duration", fmt.Sprintf("%fs", time.Since(t).Seconds()))
	}

	return nil
}

func (g *Graceful) Stop(ctx context.Context) {
	var t time.Time
	for i := len(g.services) - 1; i >= 0; i-- {
		t = time.Now()
		if err := g.services[i].Stop(ctx); err != nil {
			g.log.Error(fmt.Sprintf("%s stop failed", g.services[i].Name()), "error", err)
		}
		g.log.Info(fmt.Sprintf("%s stopped", g.services[i].Name()), "duration", fmt.Sprintf("%fs", time.Since(t).Seconds()))
	}
}
