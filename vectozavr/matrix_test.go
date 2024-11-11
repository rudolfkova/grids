package vectozavr

import (
	"reflect"
	"testing"
)

func TestMatrix_MatMul(t *testing.T) {
	type args struct {
		n Matrix
	}
	tests := []struct {
		name string
		m    Matrix
		args args
		want Matrix
	}{
		{
			name: "test1",
			m:    NewMatrix([4][4]float64{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 1, 2, 3}, {4, 5, 6, 7}}),
			args: args{n: NewMatrix([4][4]float64{{9, 8, 7, 6}, {5, 4, 3, 2}, {1, 9, 8, 7}, {6, 5, 4, 3}})},
			want: NewMatrix([4][4]float64{{46, 63, 53, 43}, {130, 167, 141, 115}, {106, 109, 94, 79}, {109, 141, 119, 97}}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.MatMul(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Matrix.Mul() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix_Vec4Mul(t *testing.T) {
	type args struct {
		v Vec4
	}
	tests := []struct {
		name string
		m    Matrix
		args args
		want Vec4
	}{
		{
			name: "test1",
			m:    NewMatrix([4][4]float64{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 1, 2, 3}, {4, 5, 6, 7}}),
			args: args{v: Vec4{9, 5, 1, 6}},
			want: Vec4{46, 130, 106, 109},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Vec4Mul(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Matrix.Vec4Mul() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix_Vec3Mul(t *testing.T) {
	type args struct {
		v Vec3
	}
	tests := []struct {
		name string
		m    Matrix
		args args
		want Vec3
	}{
		{
			name: "test1",
			m:    NewMatrix([4][4]float64{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 1, 2, 3}, {4, 5, 6, 7}}),
			args: args{v: Vec3{9, 5, 1}},
			want: Vec3{22, 82, 88},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Vec3Mul(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Matrix.Vec3Mul() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIdentity(t *testing.T) {
	tests := []struct {
		name string
		want Matrix
	}{
		{
			name: "test1",
			want: NewMatrix([4][4]float64{{1, 0, 0, 0}, {0, 1, 0, 0}, {0, 0, 1, 0}, {0, 0, 0, 1}}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Identity(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Identity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConstant(t *testing.T) {
	type args struct {
		value float64
	}
	tests := []struct {
		name string
		args args
		want Matrix
	}{
		{
			name: "test1",
			args: args{value: 1},
			want: NewMatrix([4][4]float64{{1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constant(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constant() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestZeroMatrix(t *testing.T) {
	tests := []struct {
		name string
		want Matrix
	}{
		{
			name: "test1",
			want: NewMatrix([4][4]float64{{0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ZeroMatrix(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ZeroMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScale(t *testing.T) {
	type args struct {
		v Vec3
	}
	tests := []struct {
		name string
		args args
		want Matrix
	}{
		{
			name: "test1",
			args: args{v: Vec3{1, 2, 3}},
			want: NewMatrix([4][4]float64{{1, 0, 0, 0}, {0, 2, 0, 0}, {0, 0, 3, 0}, {0, 0, 0, 1}}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Scale(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Scale() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTranslate(t *testing.T) {
	type args struct {
		v Vec3
	}
	tests := []struct {
		name string
		args args
		want Matrix
	}{
		{
			name: "test1",
			args: args{v: Vec3{1, 2, 3}},
			want: NewMatrix([4][4]float64{{1, 0, 0, 1}, {0, 1, 0, 2}, {0, 0, 1, 3}, {0, 0, 0, 1}}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Translation(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Translation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRotationX(t *testing.T) {
	type args struct {
		angle float64
	}
	tests := []struct {
		name string
		args args
		want Matrix
	}{
		{
			name: "test1",
			args: args{angle: 1},
			want: NewMatrix([4][4]float64{{1, 0, 0, 0}, {0, 0.5403023058681398, -0.8414709848078965, 0}, {0, 0.8414709848078965, 0.5403023058681398, 0}, {0, 0, 0, 1}}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RotationX(tt.args.angle); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RotationX() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRotationY(t *testing.T) {
	type args struct {
		angle float64
	}
	tests := []struct {
		name string
		args args
		want Matrix
	}{
		{
			name: "test1",
			args: args{angle: 1},
			want: NewMatrix([4][4]float64{{0.5403023058681398, 0, 0.8414709848078965, 0}, {0, 1, 0, 0}, {-0.8414709848078965, 0, 0.5403023058681398, 0}, {0, 0, 0, 1}}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RotationY(tt.args.angle); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RotationY() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix_X(t *testing.T) {
	tests := []struct {
		name string
		m    Matrix
		want Vec3
	}{
		{
			name: "test1",
			m:    NewMatrix([4][4]float64{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}),
			want: Vec3{1, 5, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.X(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Matrix.X() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix_Y(t *testing.T) {
	tests := []struct {
		name string
		m    Matrix
		want Vec3
	}{
		{
			name: "test1",
			m:    NewMatrix([4][4]float64{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}),
			want: Vec3{2, 6, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Y(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Matrix.Y() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix_Z(t *testing.T) {
	tests := []struct {
		name string
		m    Matrix
		want Vec3
	}{
		{
			name: "test1",
			m:    NewMatrix([4][4]float64{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}),
			want: Vec3{3, 7, 11},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Z(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Matrix.Z() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix_W(t *testing.T) {
	tests := []struct {
		name string
		m    Matrix
		want Vec3
	}{
		{
			name: "test1",
			m:    NewMatrix([4][4]float64{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}),
			want: Vec3{4, 8, 12},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.W(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Matrix.W() = %v, want %v", got, tt.want)
			}
		})
	}
}
