package camera

import (
	"github.com/rudolfkova/vectozavr/vectozavr"
)

type Camera struct {
	E                 vectozavr.Vec3
	ViewMatrix        vectozavr.Matrix
	InverseViewMatrix vectozavr.Matrix

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

func InverseTransform(left, up, at, e vectozavr.Vec3) vectozavr.Matrix {
	// Транспонируем подматрицу вращения (поскольку это ортогональная матрица)
	var rotation [4][4]float64
	rotation[0][0], rotation[1][0], rotation[2][0] = left.X, left.Y, left.Z
	rotation[0][1], rotation[1][1], rotation[2][1] = up.X, up.Y, up.Z
	rotation[0][2], rotation[1][2], rotation[2][2] = at.X, at.Y, at.Z
	rotation[3][3] = 1

	// Смещение
	rotation[0][3] = e.Dot(left)
	rotation[1][3] = e.Dot(up)
	rotation[2][3] = e.Dot(at)

	return vectozavr.NewMatrix(rotation)
}

func (c *Camera) InitCamera() {
	c.left = c.Left
	c.up = c.Up
	c.at = c.At
}

func (c *Camera) ViewMat() {
	c.ViewMatrix = ViewMatrix(c.Left, c.Up, c.At, c.E)
	c.InverseViewMatrix, _ = c.ViewMatrix.Inverse()
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
