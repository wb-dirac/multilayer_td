package rpc

import (
	"context"

	"go.uber.org/zap"
)

// NotifyAcks notifies engine about received acknowledgements.
func (e *Engine) NotifyAcks(ids []int64) {
	for _, id := range ids {
		e.mux.Lock()
		cb, ok := e.ack[id]
		e.mux.Unlock()

		if !ok {
			e.log.Warn("ack callback not set", zap.Int64("msg_id", id))
			continue
		}

		cb()
	}
}

// waitAck blocks until acknowledgement on message id is received.
func (e *Engine) waitAck(ctx context.Context, id int64) error {
	got := make(chan struct{})

	e.mux.Lock()
	e.ack[id] = func() { close(got) }
	e.mux.Unlock()

	defer func() {
		e.mux.Lock()
		delete(e.ack, id)
		e.mux.Unlock()
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-got:
		return nil
	}
}
