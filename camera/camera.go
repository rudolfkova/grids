package camera

import (
	"github.com/rudolfkova/vectozavr/vectozavr"
)

type Camera struct {
	E          vectozavr.Vec3
	ViewMatrix vectozavr.Matrix

	Left vectozavr.Vec3
	Up   vectozavr.Vec3
	At   vectozavr.Vec3

	left, up, at vectozavr.Vec3

	Tilt, Angle float64
	V           vectozavr.Vec3

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

func (c *Camera) InitCamera() {
	c.left = c.Left
	c.up = c.Up
	c.at = c.At
}

func (c *Camera) ViewMat() {
	c.ViewMatrix = ViewMatrix(c.Left, c.Up, c.At, c.E)
}

func (c *Camera) Move(dv vectozavr.Vec3) {
	c.E = c.E.Add(dv)
}

func (c *Camera) Rotate(tilt, angle float64) {
	c.Vert()
	c.Up = vectozavr.RotationV(c.Left, tilt).Vec4Mul(c.Up.ToVec4()).ToVec3()
	c.At = vectozavr.RotationV(c.Left, tilt).Vec4Mul(c.At.ToVec4()).ToVec3()
	c.Left = vectozavr.RotationV(c.Left, tilt).Vec4Mul(c.Left.ToVec4()).ToVec3()

	c.Up = vectozavr.RotationV(c.V, angle).Vec4Mul(c.Up.ToVec4()).ToVec3()
	c.At = vectozavr.RotationV(c.V, angle).Vec4Mul(c.At.ToVec4()).ToVec3()
	c.Left = vectozavr.RotationV(c.V, angle).Vec4Mul(c.Left.ToVec4()).ToVec3()
}

func (c *Camera) Vert() {
	c.V = vectozavr.RotationV(c.Left, -c.Tilt).Vec4Mul(c.Up.ToVec4()).ToVec3()
}
