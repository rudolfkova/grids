package camera

import "github.com/rudolfkova/vectozavr/vectozavr"

type Camera struct {
	E          vectozavr.Vec3
	ViewMatrix vectozavr.Matrix

	Left vectozavr.Vec3
	Up   vectozavr.Vec3
	At   vectozavr.Vec3

	Fov  float64
	Near float64
	Far  float64
	A    float64
}

func ViewMatrix(left vectozavr.Vec3, up vectozavr.Vec3, at vectozavr.Vec3, e vectozavr.Vec3) vectozavr.Matrix {
	return vectozavr.NewMatrix(
		[4][4]float64{
			{left.X, left.Y, left.Z, -e.Dot(left)},
			{up.X, up.Y, up.Z, -e.Dot(up)},
			{at.X, at.Y, at.Z, -e.Dot(at)},
			{0, 0, 0, 1},
		},
	)
}

func (c *Camera) ViewMat() {
	c.ViewMatrix = ViewMatrix(c.Left, c.Up, c.At, c.E)
}

func (c *Camera) Move(dv vectozavr.Vec3) {
	c.E = c.E.Add(dv)
}
