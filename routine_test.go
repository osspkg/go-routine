/*
 *  Copyright (c) 2024 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package routine

import (
	"context"
	"fmt"
	"testing"
)

func TestUnit_Parallel(t *testing.T) {
	Parallel(
		func() {
			fmt.Println("a")
		}, func() {
			fmt.Println("b")
		}, func() {
			fmt.Println("c")
		},
	)
}

func TestUnit_Repeat(t *testing.T) {
	i := 0

	Repeat(context.TODO(), 3, func(ctx context.Context) error {
		i++
		return fmt.Errorf("err")
	})

	if i != 3 {
		t.Log(i)
		t.FailNow()
	}
}
