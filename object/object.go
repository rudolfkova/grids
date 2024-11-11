package object

import (
	"github.com/rudolfkova/vectozavr/vectozavr"
)

type Object struct {
	position          vectozavr.Vec3
	TransformMatrix   vectozavr.Matrix
	angle             vectozavr.Vec3
	angleLeftUpLookAt vectozavr.Vec3
	left              vectozavr.Vec3
	up                vectozavr.Vec3
	lookAt            vectozavr.Vec3
}

func NewObject(m vectozavr.Matrix) *Object {
	return &Object{
		TransformMatrix: m,
	}
}

func (o *Object) GetX() vectozavr.Vec3 {
	return o.TransformMatrix.X()
}
func (o *Object) GetY() vectozavr.Vec3 {
	return o.TransformMatrix.Y()
}
func (o *Object) GetZ() vectozavr.Vec3 {
	return o.TransformMatrix.Z()
}
func (o *Object) GetPos() vectozavr.Vec3 {
	return o.position
}

func (o *Object) Transform(t vectozavr.Matrix) {
	o.TransformMatrix = o.TransformMatrix.MatMul(t)
}

func (o *Object) Left() {
	o.left = o.TransformMatrix.X()
}

func (o *Object) Up() {
	o.up = o.TransformMatrix.Y()
}

func (o *Object) LookAt() {
	o.lookAt = o.TransformMatrix.Z()
}

func (o *Object) TransformRelativePoint(point vectozavr.Vec3, transform vectozavr.Matrix) {
	// translate object in new coordinate system (connected with point)
	o.TransformMatrix = vectozavr.Translation(o.position.Sub(point)).MatMul(o.TransformMatrix)
	// transform object in the new coordinate system
	o.TransformMatrix = transform.MatMul(o.TransformMatrix)
	// translate object back in self connected coordinate system
	o.position = o.TransformMatrix.W().Add(point)
	o.TransformMatrix = vectozavr.Translation(o.TransformMatrix.W()).MatMul(o.TransformMatrix)

}

func (o *Object) Translate(v vectozavr.Vec3) {
	o.position = o.position.Add(v)
}

func (o *Object) Scale(s vectozavr.Vec3) {
	o.Transform(vectozavr.Scale(s))
}

func (o *Object) Rotate(a vectozavr.Vec3) {
	o.angle = o.angle.Add(a)
	o.Transform(vectozavr.Rotation(a))
}

func (o *Object) VRotate(v vectozavr.Vec3, a float64) {
	o.Transform(vectozavr.RotationV(v, a))
}

func (o *Object) RotateRelativePoint(s vectozavr.Vec3, r vectozavr.Vec3) {
	o.angle = o.angle.Add(r)
	o.TransformRelativePoint(s, vectozavr.Rotation(r))
}

func (o *Object) RotateLeft(rl float64) {
	o.angleLeftUpLookAt.X += rl
	o.VRotate(o.left, rl)
}

func (o *Object) RotateUp(rl float64) {
	o.angleLeftUpLookAt.Y += rl
	o.VRotate(o.up, rl)
}

func (o *Object) RotateLookAt(rl float64) {
	o.angleLeftUpLookAt.Z += rl
	o.VRotate(o.lookAt, rl)
}

func (o *Object) TranslateToPoint(point vectozavr.Vec3) {
	o.Translate(point.Sub(o.position))
}
