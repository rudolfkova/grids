package vectozavr

import (
	"math"
	"reflect"
	"testing"
)

func TestNewVec3(t *testing.T) {
	type args struct {
		x float64
		y float64
		z float64
	}
	tests := []struct {
		name string
		args args
		want Vec3
	}{
		{
			name: "testMaxFloat64",
			args: args{x: 823735, y: 8574249, z: 424242},
			want: Vec3{823735, 8574249, 424242},
		},
		{
			name: "testMaxFloat64",
			args: args{x: math.MaxFloat64, y: math.MaxFloat64, z: math.MaxFloat64},
			want: Vec3{math.MaxFloat64, math.MaxFloat64, math.MaxFloat64},
		},
		{
			name: "testMinFloat64",
			args: args{x: -math.MaxFloat64, y: -math.MaxFloat64, z: -math.MaxFloat64},
			want: Vec3{-math.MaxFloat64, -math.MaxFloat64, -math.MaxFloat64},
		},
		{
			name: "testZeroFloat64",
			args: args{x: 0, y: 0, z: 0},
			want: Vec3{0, 0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewVec3(tt.args.x, tt.args.y, tt.args.z); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewVec3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVec3_Add(t *testing.T) {
	type args struct {
		v2 Vec3
	}
	tests := []struct {
		name string
		v    Vec3
		args args
		want Vec3
	}{
		{
			name: "testAddVec3",
			v:    Vec3{1, 2, 3},
			args: args{v2: Vec3{3, 4, 5}},
			want: Vec3{4, 6, 8},
		},
		{
			name: "testMaxFloat64",
			v:    Vec3{math.MaxFloat64 / 2, math.MaxFloat64 / 2, math.MaxFloat64 / 2},
			args: args{v2: Vec3{math.MaxFloat64 / 2, math.MaxFloat64 / 2, math.MaxFloat64 / 2}},
			want: Vec3{math.MaxFloat64, math.MaxFloat64, math.MaxFloat64},
		},
		{
			name: "testMinFloat64",
			v:    Vec3{-math.MaxFloat64 / 2, -math.MaxFloat64 / 2, -math.MaxFloat64 / 2},
			args: args{v2: Vec3{-math.MaxFloat64 / 2, -math.MaxFloat64 / 2, -math.MaxFloat64 / 2}},
			want: Vec3{-math.MaxFloat64, -math.MaxFloat64, -math.MaxFloat64},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Add(tt.args.v2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vec3.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVec3_Sub(t *testing.T) {
	type args struct {
		v2 Vec3
	}
	tests := []struct {
		name string
		v    Vec3
		args args
		want Vec3
	}{
		{
			name: "testSubVec3",
			v:    Vec3{1, 2, 3},
			args: args{v2: Vec3{3, 4, 5}},
			want: Vec3{-2, -2, -2},
		},
		{
			name: "testMaxFloat64",
			v:    Vec3{math.MaxFloat64 / 2, math.MaxFloat64 / 2, math.MaxFloat64 / 2},
			args: args{v2: Vec3{-math.MaxFloat64 / 2, -math.MaxFloat64 / 2, -math.MaxFloat64 / 2}},
			want: Vec3{math.MaxFloat64, math.MaxFloat64, math.MaxFloat64},
		},
		{
			name: "testMinFloat64",
			v:    Vec3{-math.MaxFloat64 / 2, -math.MaxFloat64 / 2, -math.MaxFloat64 / 2},
			args: args{v2: Vec3{math.MaxFloat64 / 2, math.MaxFloat64 / 2, math.MaxFloat64 / 2}},
			want: Vec3{-math.MaxFloat64, -math.MaxFloat64, -math.MaxFloat64},
		},
		{
			name: "testZeroFloat64",
			v:    Vec3{math.MaxFloat64, math.MaxFloat64, math.MaxFloat64},
			args: args{v2: Vec3{math.MaxFloat64, math.MaxFloat64, math.MaxFloat64}},
			want: Vec3{0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Sub(tt.args.v2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vec3.Sub() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestVec3_Mul(t *testing.T) {
	type args struct {
		num float64
	}
	tests := []struct {
		name string
		v    Vec3
		args args
		want Vec3
	}{
		{
			name: "testMulFloat64",
			v:    Vec3{1, 2, 3},
			args: args{num: 2},
			want: Vec3{2, 4, 6},
		},
		{
			name: "testMaxFloat64",
			v:    Vec3{math.MaxFloat64 / 2, math.MaxFloat64 / 2, math.MaxFloat64 / 2},
			args: args{num: 2},
			want: Vec3{math.MaxFloat64, math.MaxFloat64, math.MaxFloat64},
		},
		{
			name: "testMinFloat64",
			v:    Vec3{-math.MaxFloat64 / 2, -math.MaxFloat64 / 2, -math.MaxFloat64 / 2},
			args: args{num: 2},
			want: Vec3{-math.MaxFloat64, -math.MaxFloat64, -math.MaxFloat64},
		},
		{
			name: "testZeroFloat64",
			v:    Vec3{math.MaxFloat64, math.MaxFloat64, math.MaxFloat64},
			args: args{num: 0},
			want: Vec3{0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Mul(tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vec3.Mul() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVec3_Div(t *testing.T) {
	type args struct {
		num float64
	}
	tests := []struct {
		name    string
		v       Vec3
		args    args
		want    Vec3
		wantErr bool
	}{
		{
			name:    "testDivFloat64",
			v:       Vec3{1, 2, 3},
			args:    args{num: 2},
			want:    Vec3{0.5, 1, 1.5},
			wantErr: false,
		},
		{
			name:    "testZeroFloat64",
			v:       Vec3{math.MaxFloat64, math.MaxFloat64, math.MaxFloat64},
			args:    args{num: 0},
			want:    Vec3{math.MaxFloat64, math.MaxFloat64, math.MaxFloat64},
			wantErr: true,
		},
		{
			name:    "testMaxFloat64",
			v:       Vec3{math.MaxFloat64 / 2, math.MaxFloat64 / 2, math.MaxFloat64 / 2},
			args:    args{num: 0.5},
			want:    Vec3{math.MaxFloat64, math.MaxFloat64, math.MaxFloat64},
			wantErr: false,
		},
		{
			name:    "testMinFloat64",
			v:       Vec3{-math.MaxFloat64 / 2, -math.MaxFloat64 / 2, -math.MaxFloat64 / 2},
			args:    args{num: 0.5},
			want:    Vec3{-math.MaxFloat64, -math.MaxFloat64, -math.MaxFloat64},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.v.Div(tt.args.num)
			if (err != nil) != tt.wantErr {
				t.Errorf("Vec3.Div() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vec3.Div() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVec3_Len(t *testing.T) {
	tests := []struct {
		name    string
		v       Vec3
		want    float64
		wantErr bool
	}{
		{
			name:    "testLen1",
			v:       Vec3{0, 2, 2},
			want:    math.Sqrt(8),
			wantErr: false,
		},
		{
			name:    "testLen2",
			v:       Vec3{3, 4, 0},
			want:    5,
			wantErr: false,
		},
		{
			name:    "testLen3",
			v:       Vec3{1, 2, 2},
			want:    3,
			wantErr: false,
		},
		{
			name:    "testLenZero",
			v:       Vec3{0, 0, 0},
			want:    0,
			wantErr: false,
		},
		{
			name:    "testLenMax",
			v:       Vec3{math.MaxFloat64, math.MaxFloat64, math.MaxFloat64},
			want:    0,
			wantErr: true,
		},
		{
			name:    "testLenMin",
			v:       Vec3{-math.MaxFloat64, -math.MaxFloat64, -math.MaxFloat64},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.v.Len()
			if (err != nil) != tt.wantErr {
				t.Errorf("Vec3.Len() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Vec3.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVec3_Normalize(t *testing.T) {
	tests := []struct {
		name    string
		v       Vec3
		want    Vec3
		wantErr bool
	}{
		{
			name:    "testNormalize1",
			v:       Vec3{0, 3, 4},
			want:    Vec3{0, 0.6, 0.8},
			wantErr: false,
		},
		{
			name:    "testNormalize2",
			v:       Vec3{3, 0, 4},
			want:    Vec3{0.6, 0, 0.8},
			wantErr: false,
		},
		{
			name:    "testNormalizeZero",
			v:       Vec3{0, 0, 0},
			want:    Vec3{0, 0, 0},
			wantErr: true,
		},
		{
			name:    "testNormalizeMax",
			v:       Vec3{math.MaxFloat64, math.MaxFloat64, math.MaxFloat64},
			want:    Vec3{math.MaxFloat64, math.MaxFloat64, math.MaxFloat64},
			wantErr: true,
		},
		{
			name:    "testNormalizeMin",
			v:       Vec3{-math.MaxFloat64, -math.MaxFloat64, -math.MaxFloat64},
			want:    Vec3{-math.MaxFloat64, -math.MaxFloat64, -math.MaxFloat64},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.v.Normalize()
			if (err != nil) != tt.wantErr {
				t.Errorf("Vec3.Normalize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vec3.Normalize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVec3_Dot(t *testing.T) {
	type args struct {
		v2 Vec3
	}
	tests := []struct {
		name string
		v    Vec3
		args args
		want float64
	}{
		{
			name: "testDot1",
			v:    Vec3{1, 2, 3},
			args: args{v2: Vec3{1, 2, 3}},
			want: 14,
		},
		{
			name: "testDot2",
			v:    Vec3{1, 2, 3},
			args: args{v2: Vec3{3, 2, 1}},
			want: 10,
		},
		{
			name: "testDotZero",
			v:    Vec3{0, 0, 0},
			args: args{v2: Vec3{0, 0, 0}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Dot(tt.args.v2); got != tt.want {
				t.Errorf("Vec3.Dot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVec3_Cross(t *testing.T) {
	type args struct {
		v2 Vec3
	}
	tests := []struct {
		name string
		v    Vec3
		args args
		want Vec3
	}{
		{
			name: "testCross1",
			v:    Vec3{1, 2, 3},
			args: args{v2: Vec3{3, 2, 1}},
			want: Vec3{-4, 8, -4},
		},
		{
			name: "testCross2",
			v:    Vec3{775, 657, 537},
			args: args{v2: Vec3{846, 235, 127}},
			want: Vec3{-42756, 355877, -373697},
		},
		{
			name: "testCross3",
			v:    Vec3{-775, -657, -537},
			args: args{v2: Vec3{846, 235, 127}},
			want: Vec3{42756, -355877, 373697},
		},
		{
			name: "testCross4",
			v:    Vec3{1, 2, 3},
			args: args{v2: Vec3{0, 0, 0}},
			want: Vec3{0, 0, 0},
		},
		{
			name: "testCrossZero",
			v:    Vec3{0, 0, 0},
			args: args{v2: Vec3{0, 0, 0}},
			want: Vec3{0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Cross(tt.args.v2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vec3.Cross() = %v, want %v", got, tt.want)
			}
		})
	}
}
