package pubsub

import (
	"context"
	"errors"
	"log"
	"sync"
	"sync/atomic"
	"time"
)

type bus struct {
	bus     Bus
	closed  atomic.Bool
	closeCh chan struct{}
	wg      sync.WaitGroup
}

func New() PubSub {
	return &bus{
		closeCh: make(chan struct{}),
	}
}

func (a *bus) Init(_ context.Context) error {
	a.bus = NewEventbus(true)

	return nil
}

func (a *bus) Publish(_ context.Context, req *PublishRequest) error {
	if a.closed.Load() {
		return errors.New("component is closed")
	}

	a.bus.Publish(req.Topic, req.Data)

	return nil
}

func (a *bus) Subscribe(ctx context.Context, req SubscribeRequest, handler Handler) error {
	if a.closed.Load() {
		return errors.New("component is closed")
	}

	// For this component we allow built-in retries because it is backed by memory
	retryHandler := func(data []byte) {
		for range 10 {
			handleErr := handler(ctx, &NewMessage{Data: data, Topic: req.Topic, Metadata: req.Metadata})
			if handleErr == nil {
				break
			}
			log.Println(handleErr)
			select {
			case <-time.After(100 * time.Millisecond):
				// Nop
			case <-ctx.Done():
				return
			}
		}
	}
	err := a.bus.SubscribeAsync(req.Topic, retryHandler, true)
	if err != nil {
		return err
	}

	// Unsubscribe when context is done
	a.wg.Add(1)
	go func() {
		defer a.wg.Done()
		select {
		case <-ctx.Done():
		case <-a.closeCh:
		}
		err := a.bus.Unsubscribe(req.Topic, retryHandler)
		if err != nil {
			log.Printf("error while unsubscribing from topic %s: %v", req.Topic, err)
		}
	}()

	return nil
}

func (a *bus) Close() error {
	if a.closed.CompareAndSwap(false, true) {
		close(a.closeCh)
	}
	a.wg.Wait()
	return nil
}
