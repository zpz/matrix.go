// Copyright ©2013 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dense

import (
	check "launchpad.net/gocheck"
)

func (s *S) TestCholesky(c *check.C) {
	for _, t := range []struct {
		a   *Dense
		spd bool
	}{
		{
			a: make_dense(3, 3, []float64{
				4, 1, 1,
				1, 2, 3,
				1, 3, 6,
			}),

			spd: true,
		},
	} {
		cf := Cholesky(t.a)
		c.Check(cf.SPD, check.Equals, t.spd)

		lt := T(cf.L, nil)

		lc := Mult(cf.L, lt, nil)
		c.Check(EqualApprox(lc, t.a, 1e-12), check.Equals, true)

		x := cf.Solve(eye(3))

		t.a = Mult(t.a, x, nil)
		c.Check(EqualApprox(t.a, eye(3), 1e-12), check.Equals, true)
	}
}
