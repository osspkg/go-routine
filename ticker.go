/*
 *  Copyright (c) 2024-2025 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package routine

import (
	"context"
	"sync"
	"time"
)

type (
	Ticker struct {
		Interval time.Duration
		OnStart  bool
		Calls    []TickFunc

		wg sync.WaitGroup
	}

	TickFunc func(context.Context, time.Time)
)

func (t *Ticker) callAll(ctx context.Context, now time.Time) {
	t.wg.Add(len(t.Calls))
	for _, call := range t.Calls {
		call := call
		go func() {
			defer func() {
				recover() //nolint:errcheck
				t.wg.Done()
			}()

			call(ctx, now)
		}()
	}
}

func (t *Ticker) Run(ctx context.Context) {
	defer func() {
		t.wg.Wait()
	}()

	if t.OnStart {
		t.callAll(ctx, time.Now())
	}

	if t.Interval == 0 {
		t.Interval = time.Second
	}

	tick := time.NewTicker(t.Interval)
	defer tick.Stop()

	for {
		select {
		case <-ctx.Done():
			return

		case now := <-tick.C:
			t.callAll(ctx, now)
		}
	}
}
