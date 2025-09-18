/*
 *  Copyright (c) 2024-2025 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package routine

import (
	"context"
	"sync/atomic"
	"testing"
	"time"
)

func TestUnit_Ticker_Run(t1 *testing.T) {
	var counter uint64

	t := Ticker{
		Interval: 100 * time.Millisecond,
		OnStart:  true,
		Calls: []TickFunc{
			func(context.Context, time.Time) {
				atomic.AddUint64(&counter, 1)
			},
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	t.Run(ctx)

	if counter < 9 || counter > 11 {
		t1.Errorf("counter: %d", counter)
	}
}
