/*
 *  Copyright (c) 2024-2025 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package tick_test

import (
	"context"
	"sync/atomic"
	"testing"
	"time"

	"go.osspkg.com/routine/tick"
)

func TestUnit_Ticker_Run(t *testing.T) {
	var counter uint64

	tik := tick.Ticker{
		Calls: []tick.Config{
			{
				Name:     "text",
				OnStart:  true,
				Interval: time.Millisecond * 100,
				Func: func(ctx context.Context, ct time.Time) error {
					atomic.AddUint64(&counter, 1)
					return nil
				},
			},
		},
		OnError: func(n string, e error) {
			t.Log(n, e)
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	tik.Run(ctx)

	cnt := atomic.LoadUint64(&counter)

	if cnt < 8 {
		t.Errorf("counter: %d", cnt)

	}
	t.Logf("counter: %d", cnt)
}
