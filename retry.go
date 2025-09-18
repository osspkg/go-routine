/*
 *  Copyright (c) 2024-2025 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package routine

import (
	"context"
	"time"
)

func Retry(ctx context.Context, count int, ttl time.Duration, call func(context.Context) error) error {
	var err error

	for i := 0; i < count; i++ {
		if err = call(ctx); err != nil {
			time.Sleep(ttl)
			ttl += ttl / time.Duration(count)

			continue
		}

		return nil
	}

	return err
}
