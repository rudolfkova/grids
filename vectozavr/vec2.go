package vectozavr

import (
	"fmt"
	"math"
)

const (
	Zero = 0.001
)

// A 2D vector
type Vec2 struct {
	X, Y float64
}

// Creates a new Vec2 with the given coordinates
func NewVec2(x, y float64) Vec2 {
	return Vec2{X: x, Y: y}
}

// Adding two vectors
func (v Vec2) Add(v2 Vec2) Vec2 {
	return Vec2{X: v.X + v2.X, Y: v.Y + v2.Y}
}

// Subtracting two vectors
func (v Vec2) Sub(v2 Vec2) Vec2 {
	return Vec2{X: v.X - v2.X, Y: v.Y - v2.Y}
}

// Multiplying a vector by a number
func (v Vec2) Mul(num float64) Vec2 {
	return Vec2{X: v.X * num, Y: v.Y * num}
}

// Dividing a vector by a number
func (v Vec2) Div(num float64) (Vec2, error) {
	if num <= Zero && num >= -Zero {
		return v, fmt.Errorf("cannot divide by Zero")
	}
	return Vec2{X: v.X / num, Y: v.Y / num}, nil

}

// Returns the length of the vector
func (v Vec2) Len() (float64, error) {
	l := math.Sqrt(v.X*v.X + v.Y*v.Y)
	if l == math.Inf(1) {
		return 0, fmt.Errorf("cannot calculate length of vector: length is infinite")
	}
	return l, nil
}

// Normalizing a vector
func (v Vec2) Normalize() (Vec2, error) {
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
func (v Vec2) Dot(v2 Vec2) float64 {
	return v.X*v2.X + v.Y*v2.Y
}

func (v Vec2) ToVec4() Vec4 {
	return Vec4{X: v.X, Y: v.Y, Z: 0, W: 1}
}
