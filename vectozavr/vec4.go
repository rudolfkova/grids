package vectozavr

import (
	"fmt"
	"math"
)

type Vec4 struct {
	X, Y, Z, W float64
}

func NewVec4(x, y, z, w float64) Vec4 {
	return Vec4{
		X: x, Y: y, Z: z, W: w,
	}
}

func (v Vec4) Add(v2 Vec4) Vec4 {
	return Vec4{X: v.X + v2.X, Y: v.Y + v2.Y, Z: v.Z + v2.Z, W: v.W + v2.W}
}

func (v Vec4) Sub(v2 Vec4) Vec4 {
	return Vec4{X: v.X - v2.X, Y: v.Y - v2.Y, Z: v.Z - v2.Z, W: v.W - v2.W}
}

func (v Vec4) Mul(num float64) Vec4 {
	return Vec4{X: v.X * num, Y: v.Y * num, Z: v.Z * num, W: v.W * num}
}

func (v Vec4) Div(num float64) (Vec4, error) {
	if num <= Zero && num >= -Zero {
		return v, fmt.Errorf("cannot divide by Zero")
	}
	return Vec4{X: v.X / num, Y: v.Y / num, Z: v.Z / num, W: v.W / num}, nil

}

func (v Vec4) Len() (float64, error) {
	l := math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z + v.W*v.W)
	if l == math.Inf(1) {
		return 0, fmt.Errorf("cannot calculate length of vector: length is infinity")
	}
	return l, nil
}

func (v Vec4) Normalize() (Vec4, error) {
	l, err1 := v.Len()
	if err1 != nil {
		return v, fmt.Errorf("cannot normalize vector: %v", err1)
	}
	v2, err2 := v.Div(l)
	if err2 != nil {
		return v, fmt.Errorf("cannot normalize vector: %v", err2)
	}
	return v2, nil
}

func (v Vec4) Dot(v2 Vec4) float64 {
	return v.X*v2.X + v.Y*v2.Y + v.Z*v2.Z + v.W*v2.W
}

func (v Vec4) ToVec3() Vec3 {

	return Vec3{X: v.X, Y: v.Y, Z: v.Z}
}
