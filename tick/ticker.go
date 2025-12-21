/*
 *  Copyright (c) 2024-2025 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package tick

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type (
	Ticker struct {
		Calls   []Config
		OnError func(name string, err error)
		wg      sync.WaitGroup
	}

	Config struct {
		Name     string
		OnStart  bool
		Interval time.Duration
		Func     Func
	}

	Func func(context.Context, time.Time) error
)

func (t *Ticker) runConfig(ctx context.Context, c Config) {
	go func() {
		defer t.wg.Done()

		c.Interval = max(c.Interval, time.Millisecond*100)

		tik := time.NewTicker(c.Interval)
		defer tik.Stop()

		if c.OnStart {
			if err := c.Func(ctx, time.Now()); err != nil {
				t.OnError(c.Name, err)
			}
		}

		for {
			select {
			case <-ctx.Done():
				return
			case ct := <-tik.C:
				if err := c.Func(ctx, ct); err != nil {
					t.OnError(c.Name, err)
				}
			}
		}
	}()
}

func (t *Ticker) Run(ctx context.Context) {
	if t.OnError == nil {
		t.OnError = func(string, error) {}
	}

	for _, call := range t.Calls {
		if call.Func == nil {
			t.OnError(call.Name, fmt.Errorf("no function"))
			continue
		}

		t.wg.Add(1)
		call := call
		t.runConfig(ctx, call)
	}

	t.wg.Wait()
}
