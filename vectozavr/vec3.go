package vectozavr

import (
	"fmt"
	"math"
)

// A 3D vector

type Vec3 struct {
	X, Y, Z float64
}

// Creates a new Vec3 with the given values
func NewVec3(x, y, z float64) Vec3 {
	return Vec3{X: x, Y: y, Z: z}
}

// Adding two vectors
func (v Vec3) Add(v2 Vec3) Vec3 {
	return Vec3{X: v.X + v2.X, Y: v.Y + v2.Y, Z: v.Z + v2.Z}
}

// Subtracting two vectors
func (v Vec3) Sub(v2 Vec3) Vec3 {
	return Vec3{X: v.X - v2.X, Y: v.Y - v2.Y, Z: v.Z - v2.Z}
}

// Multiplying a vector by a number
func (v Vec3) Mul(num float64) Vec3 {
	return Vec3{X: v.X * num, Y: v.Y * num, Z: v.Z * num}
}

// Dividing a vector by a number
func (v Vec3) Div(num float64) (Vec3, error) {
	if num <= Zero && num >= -Zero {
		return v, fmt.Errorf("cannot divide by Zero")
	}
	return Vec3{X: v.X / num, Y: v.Y / num, Z: v.Z / num}, nil

}

// Returns the length of the vector
func (v Vec3) Len() (float64, error) {
	l := math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
	if l == math.Inf(1) {
		return 0, fmt.Errorf("cannot calculate length of vector: length is infinity")
	}
	return l, nil
}

// Normalizing a vector
func (v Vec3) Normalize() (Vec3, error) {
	l, err1 := v.Len()
	if err1 != nil {
		return v, fmt.Errorf("cannot normalize: %v", err1)
	}
	v2, err2 := v.Div(l)
	if err2 != nil {
		return v, fmt.Errorf("cannot normalize: %v", err2)
	}
	return v2, nil
}

// The scalar product
func (v Vec3) Dot(v2 Vec3) float64 {
	return v.X*v2.X + v.Y*v2.Y + v.Z*v2.Z
}

// The vector product
func (v Vec3) Cross(v2 Vec3) Vec3 {
	return NewVec3(v.Y*v2.Z-v.Z*v2.Y, v.Z*v2.X-v.X*v2.Z, v.X*v2.Y-v.Y*v2.X)
}
