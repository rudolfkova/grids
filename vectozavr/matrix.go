package vectozavr

import (
	"math"
)

type Matrix struct {
	m [4][4]float64
}

// Создание новой матрицы
func NewMatrix(m [4][4]float64) Matrix {
	return Matrix{
		m: m,
	}

}

// Умножение матрицы на матрицу
func (m Matrix) MatMul(n Matrix) Matrix {
	var result [4][4]float64
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			for k := 0; k < 4; k++ {
				result[i][j] += m.m[i][k] * n.m[k][j]
			}

		}

	}
	return NewMatrix(result)
}

// Умножение матрицы на вектор размером 4
func (m Matrix) Vec4Mul(v Vec4) Vec4 {
	var v1 Vec4
	v1.X = m.m[0][0]*v.X + m.m[0][1]*v.Y + m.m[0][2]*v.Z + m.m[0][3]*v.W
	v1.Y = m.m[1][0]*v.X + m.m[1][1]*v.Y + m.m[1][2]*v.Z + m.m[1][3]*v.W
	v1.Z = m.m[2][0]*v.X + m.m[2][1]*v.Y + m.m[2][2]*v.Z + m.m[2][3]*v.W
	v1.W = m.m[3][0]*v.X + m.m[3][1]*v.Y + m.m[3][2]*v.Z + m.m[3][3]*v.W

	return v1

}

// Умножение матрицы на вектор размером 3
func (m Matrix) Vec3Mul(v Vec3) Vec3 {
	var v1 Vec3
	v1.X = m.m[0][0]*v.X + m.m[0][1]*v.Y + m.m[0][2]*v.Z
	v1.Y = m.m[1][0]*v.X + m.m[1][1]*v.Y + m.m[1][2]*v.Z
	v1.Z = m.m[2][0]*v.X + m.m[2][1]*v.Y + m.m[2][2]*v.Z

	return v1
}

// Единичная матрица
func Identity() Matrix {
	return NewMatrix([4][4]float64{
		{1, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	})
}

// Матрица, все элементы который заданные числа
func Constant(value float64) Matrix {
	var m Matrix
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			m.m[i][j] = value
		}
	}
	return m
}

// Пустая матрица (все значения нулевые)
func ZeroMatrix() Matrix {
	return NewMatrix([4][4]float64{})
}

// Матрица изменения масштаба
func Scale(v Vec3) Matrix {
	return NewMatrix([4][4]float64{
		{v.X, 0, 0, 0},
		{0, v.Y, 0, 0},
		{0, 0, v.Z, 0},
		{0, 0, 0, 1},
	})

}

// Матрица перемещения
func Translation(v Vec3) Matrix {
	return NewMatrix([4][4]float64{
		{1, 0, 0, v.X},
		{0, 1, 0, v.Y},
		{0, 0, 1, v.Z},
		{0, 0, 0, 1},
	})

}

// Матрица поворота вокруг оси X
func RotationX(angle float64) Matrix {
	c := math.Cos(angle)
	s := math.Sin(angle)

	return NewMatrix([4][4]float64{
		{1, 0, 0, 0},
		{0, c, -s, 0},
		{0, s, c, 0},
		{0, 0, 0, 1},
	})

}

// Матрица поворота вокруг оси Y
func RotationY(angle float64) Matrix {
	c := math.Cos(angle)
	s := math.Sin(angle)

	return NewMatrix([4][4]float64{
		{c, 0, s, 0},
		{0, 1, 0, 0},
		{-s, 0, c, 0},
		{0, 0, 0, 1},
	})

}

func InverseRotationY(angle float64) Matrix {
	return RotationY(-angle) // транспонирование для ортогональной матрицы
}

func InverseRotationX(angle float64) Matrix {
	return RotationX(-angle) // транспонирование для ортогональной матрицы
}

// Матрица поворота вокруг оси Z
func RotationZ(angle float64) Matrix {
	c := math.Cos(angle)
	s := math.Sin(angle)

	return NewMatrix([4][4]float64{
		{c, -s, 0, 0},
		{s, c, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	})
}

// Матрица поворота (по всем осям)
func Rotation(v Vec3) Matrix {
	r := RotationX(v.X).MatMul(RotationY(v.Y))
	return r.MatMul(RotationZ(v.Z))
}

func RotationV(v Vec3, a float64) Matrix {
	var r Matrix
	nv, err1 := v.Normalize()
	if err1 != nil {
		return ZeroMatrix()
	}
	c := math.Cos(a)
	s := math.Sin(a)
	r.m[0][0] = c + (1.0-c)*nv.X*nv.X
	r.m[0][1] = (1.0-c)*nv.X*nv.Y - s*nv.Z
	r.m[0][2] = (1.0-c)*nv.X*nv.Z + s*nv.Y

	r.m[1][0] = (1.0-c)*nv.X*nv.Y + s*nv.Z
	r.m[1][1] = c + (1.0-c)*nv.Y*nv.Y
	r.m[1][2] = (1.0-c)*nv.Y*nv.Z - s*nv.X

	r.m[2][0] = (1.0-c)*nv.X*nv.Z - s*nv.Y
	r.m[2][1] = (1.0-c)*nv.Y*nv.Z + s*nv.X
	r.m[2][2] = c + (1.0-c)*nv.Z*nv.Z

	r.m[3][3] = 1

	return r
}

// Получить X координаты из матрицы
func (m Matrix) X() Vec3 {
	return Vec3{m.m[0][0], m.m[1][0], m.m[2][0]}
}

// Получить Y координаты из матрицы
func (m Matrix) Y() Vec3 {
	return Vec3{m.m[0][1], m.m[1][1], m.m[2][1]}
}

//Получить Z координаты из матрицы

func (m Matrix) Z() Vec3 {
	return Vec3{m.m[0][2], m.m[1][2], m.m[2][2]}
}

//Получить W вектор из матрицы

func (m Matrix) W() Vec3 {
	return Vec3{m.m[0][3], m.m[1][3], m.m[2][3]}
}

// Создаёт патрицу проекции
func Projection(fov float64, aspect, ZNear, ZFar float64) Matrix {
	return NewMatrix([4][4]float64{
		{1.0 / (math.Tan(math.Pi*fov*0.5/180) * aspect), 0, 0, 0},
		{0, 1.0 / math.Tan(math.Pi*fov*0.5/180), 0, 0},
		{0, 0, ZFar / (ZFar - ZNear), -ZFar * ZNear / (ZFar - ZNear)},
		{0, 0, 1, 0},
	})
}

func InverseProjection(fov float64, aspect, ZNear, ZFar float64) Matrix {
	return NewMatrix([4][4]float64{
		{math.Tan(math.Pi*fov*0.5/180) * aspect, 0, 0, 0},
		{0, math.Tan(math.Pi * fov * 0.5 / 180), 0, 0},
		{0, 0, 0, 1},
		{0, 0, (ZFar - ZNear) / ZFar, (ZFar - ZNear) / -ZFar * ZNear},
	})
}

// Создаёт матрицу экранного пространства
func ScreenSpace(width, height float64) Matrix {
	return NewMatrix([4][4]float64{
		{-0.5 * width, 0, 0, 0.5 * width},
		{0, -0.5 * height, 0, 0.5 * height},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	})
}

func InverseScreenSpace(width, height float64) Matrix {
	return NewMatrix([4][4]float64{
		{-2.0 / width, 0, 0, 2.0 / width},
		{0, -2.0 / height, 0, 2.0 / height},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	})
}

func (m1 Matrix) Determinant() float64 {
	m := m1.m
	return m[0][0]*m[1][1]*m[2][2]*m[3][3] +
		m[0][1]*m[1][2]*m[2][3]*m[3][0] +
		m[0][2]*m[1][3]*m[2][0]*m[3][1] +
		m[0][3]*m[1][0]*m[2][1]*m[3][2] -
		m[0][3]*m[1][2]*m[2][1]*m[3][0] -
		m[0][0]*m[1][3]*m[2][2]*m[3][1] -
		m[0][1]*m[1][0]*m[2][3]*m[3][2] -
		m[0][2]*m[1][1]*m[2][0]*m[3][3]
}

func (m1 Matrix) Inverse() Matrix {
	matrix := m1.m
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if matrix[i][j] != 0 {
				matrix[i][j] = 1 / matrix[i][j]
			}
		}
	}
	return NewMatrix(matrix)
}

func NewMatrixVec3(X, Y, Z Vec3) Matrix {
	return NewMatrix([4][4]float64{
		{X.X, Y.X, Z.X, 0},
		{X.Y, Y.Y, Z.Y, 0},
		{X.Z, Y.Z, Z.Z, 0},
		{0, 0, 0, 1},
	})

}

func Bazis(X, Y, Z Vec3) Matrix {
	return NewMatrix([4][4]float64{
		{X.X, X.Y, X.Z, 0},
		{Y.X, Y.Y, Y.Z, 0},
		{Z.X, Z.Y, Z.Z, 0},
		{0, 0, 0, 1},
	})
}
