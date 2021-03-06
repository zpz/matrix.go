// Copyright ©2013 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dense

import (
	check "launchpad.net/gocheck"
)

func (s *S) TestCholesky(c *check.C) {
	data := []struct {
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
	}

	for _, t := range data {
		cl, ok := Chol(t.a)
		c.Check(ok, check.Equals, t.spd)

		c.Check(Approx(
			Mult(cl.L(), T(cl.L(), nil), nil),
			t.a,
			1e-12),
			check.Equals, true)

		c.Check(Approx(
			Mult(t.a, cl.Solve(eye(3)), nil),
			eye(3),
			1e-12),
			check.Equals, true)

		ok = cl.Chol(t.a)
		c.Check(ok, check.Equals, t.spd)

		c.Check(Approx(
			Mult(t.a, cl.Inv(nil), nil),
			eye(3),
			1e-12),
			check.Equals, true)

		b := make_dense(3, 4, []float64{
			1, 2, 3, 4,
			4, 5, 2, 3,
			6, 7, 8, 9})
		c.Check(Approx(
			b,
			Mult(t.a, cl.Solve(Clone(b)), nil),
			1e-12),
			check.Equals, true)

		bt := T(b, nil)
		c.Check(Approx(
			bt,
			Mult(cl.SolveR(Clone(bt)), t.a, nil),
			1e-12),
			check.Equals, true)

		c.Check(Approx(
			Inv(Clone(t.a), nil),
			cl.Inv(nil),
			1e-12),
			check.Equals, true)

		c.Check(approx(
			t.a.Det(),
			cl.Det(),
			1e-12),
			check.Equals, true)
	}
}
