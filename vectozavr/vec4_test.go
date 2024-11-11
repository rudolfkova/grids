package vectozavr

import (
	"math"
	"reflect"
	"testing"
)

func TestNewVec4(t *testing.T) {
	type args struct {
		x float64
		y float64
		z float64
		w float64
	}
	tests := []struct {
		name string
		args args
		want Vec4
	}{
		{
			name: "testMaxFloat64",
			args: args{x: 823735, y: 8574249, z: 424242, w: 123456},
			want: Vec4{823735, 8574249, 424242, 123456},
		},
		{
			name: "testMaxFloat64",
			args: args{x: math.MaxFloat64, y: math.MaxFloat64, z: math.MaxFloat64, w: math.MaxFloat64},
			want: Vec4{math.MaxFloat64, math.MaxFloat64, math.MaxFloat64, math.MaxFloat64},
		},
		{
			name: "testMinFloat64",
			args: args{x: -math.MaxFloat64, y: -math.MaxFloat64, z: -math.MaxFloat64, w: -math.MaxFloat64},
			want: Vec4{-math.MaxFloat64, -math.MaxFloat64, -math.MaxFloat64, -math.MaxFloat64},
		},
		{
			name: "testZeroFloat64",
			args: args{x: 0, y: 0, z: 0, w: 0},
			want: Vec4{0, 0, 0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewVec4(tt.args.x, tt.args.y, tt.args.z, tt.args.w); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewVec4() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVec4_Add(t *testing.T) {
	type args struct {
		v2 Vec4
	}
	tests := []struct {
		name string
		v    Vec4
		args args
		want Vec4
	}{
		{
			name: "testAddVec4",
			v:    Vec4{1, 2, 3, 4},
			args: args{v2: Vec4{4, 5, 6, 7}},
			want: Vec4{5, 7, 9, 11},
		},
		{
			name: "testMaxFloat64",
			v:    Vec4{math.MaxFloat64 / 2, math.MaxFloat64 / 2, math.MaxFloat64 / 2, math.MaxFloat64 / 2},
			args: args{v2: Vec4{math.MaxFloat64 / 2, math.MaxFloat64 / 2, math.MaxFloat64 / 2, math.MaxFloat64 / 2}},
			want: Vec4{math.MaxFloat64, math.MaxFloat64, math.MaxFloat64, math.MaxFloat64},
		},
		{
			name: "testMinFloat64",
			v:    Vec4{-math.MaxFloat64 / 2, -math.MaxFloat64 / 2, -math.MaxFloat64 / 2, -math.MaxFloat64 / 2},
			args: args{v2: Vec4{-math.MaxFloat64 / 2, -math.MaxFloat64 / 2, -math.MaxFloat64 / 2, -math.MaxFloat64 / 2}},
			want: Vec4{-math.MaxFloat64, -math.MaxFloat64, -math.MaxFloat64, -math.MaxFloat64},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Add(tt.args.v2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vec4.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVec4_Sub(t *testing.T) {
	type args struct {
		v2 Vec4
	}
	tests := []struct {
		name string
		v    Vec4
		args args
		want Vec4
	}{
		{
			name: "testSubVec4",
			v:    Vec4{1, 2, 3, 4},
			args: args{v2: Vec4{3, 4, 5, 6}},
			want: Vec4{-2, -2, -2, -2},
		},
		{
			name: "testMaxFloat64",
			v:    Vec4{math.MaxFloat64 / 2, math.MaxFloat64 / 2, math.MaxFloat64 / 2, math.MaxFloat64 / 2},
			args: args{v2: Vec4{-math.MaxFloat64 / 2, -math.MaxFloat64 / 2, -math.MaxFloat64 / 2, -math.MaxFloat64 / 2}},
			want: Vec4{math.MaxFloat64, math.MaxFloat64, math.MaxFloat64, math.MaxFloat64},
		},
		{
			name: "testMinFloat64",
			v:    Vec4{-math.MaxFloat64 / 2, -math.MaxFloat64 / 2, -math.MaxFloat64 / 2, -math.MaxFloat64 / 2},
			args: args{v2: Vec4{math.MaxFloat64 / 2, math.MaxFloat64 / 2, math.MaxFloat64 / 2, math.MaxFloat64 / 2}},
			want: Vec4{-math.MaxFloat64, -math.MaxFloat64, -math.MaxFloat64, -math.MaxFloat64},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Sub(tt.args.v2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vec4.Sub() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVec4_Mul(t *testing.T) {
	type args struct {
		num float64
	}
	tests := []struct {
		name string
		v    Vec4
		args args
		want Vec4
	}{
		{
			name: "testMulFloat64",
			v:    Vec4{1, 2, 3, 4},
			args: args{num: 2},
			want: Vec4{2, 4, 6, 8},
		},
		{
			name: "testMaxFloat64",
			v:    Vec4{math.MaxFloat64 / 2, math.MaxFloat64 / 2, math.MaxFloat64 / 2, math.MaxFloat64 / 2},
			args: args{num: 2},
			want: Vec4{math.MaxFloat64, math.MaxFloat64, math.MaxFloat64, math.MaxFloat64},
		},
		{
			name: "testMinFloat64",
			v:    Vec4{-math.MaxFloat64 / 2, -math.MaxFloat64 / 2, -math.MaxFloat64 / 2, -math.MaxFloat64 / 2},
			args: args{num: 2},
			want: Vec4{-math.MaxFloat64, -math.MaxFloat64, -math.MaxFloat64, -math.MaxFloat64},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Mul(tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vec4.Mul() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVec4_Div(t *testing.T) {
	type args struct {
		num float64
	}
	tests := []struct {
		name    string
		v       Vec4
		args    args
		want    Vec4
		wantErr bool
	}{
		{
			name:    "testDivFloat64",
			v:       Vec4{1, 2, 3, 4},
			args:    args{num: 2},
			want:    Vec4{0.5, 1, 1.5, 2},
			wantErr: false,
		},
		{
			name:    "testZeroFloat64",
			v:       Vec4{math.MaxFloat64, math.MaxFloat64, math.MaxFloat64, math.MaxFloat64},
			args:    args{num: 0},
			want:    Vec4{math.MaxFloat64, math.MaxFloat64, math.MaxFloat64, math.MaxFloat64},
			wantErr: true,
		},
		{
			name:    "testMaxFloat64",
			v:       Vec4{math.MaxFloat64 / 2, math.MaxFloat64 / 2, math.MaxFloat64 / 2, math.MaxFloat64 / 2},
			args:    args{num: 0.5},
			want:    Vec4{math.MaxFloat64, math.MaxFloat64, math.MaxFloat64, math.MaxFloat64},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.v.Div(tt.args.num)
			if (err != nil) != tt.wantErr {
				t.Errorf("Vec4.Div() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vec4.Div() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVec4_Len(t *testing.T) {
	tests := []struct {
		name    string
		v       Vec4
		want    float64
		wantErr bool
	}{
		{
			name:    "testLen1",
			v:       Vec4{0, 2, 2, 2},
			want:    math.Sqrt(12),
			wantErr: false,
		},
		{
			name:    "testLen2",
			v:       Vec4{3, 4, 0, 1},
			want:    math.Sqrt(26),
			wantErr: false,
		},
		{
			name:    "testLenInf",
			v:       Vec4{math.MaxFloat64, math.MaxFloat64, math.MaxFloat64, math.MaxFloat64},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.v.Len()
			if (err != nil) != tt.wantErr {
				t.Errorf("Vec4.Len() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Vec4.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVec4_Normalize(t *testing.T) {
	tests := []struct {
		name    string
		v       Vec4
		want    Vec4
		wantErr bool
	}{
		{
			name:    "testNormalize1",
			v:       Vec4{0, 3, 4, 0},
			want:    Vec4{0, 0.6, 0.8, 0},
			wantErr: false,
		},
		{
			name:    "testNormalizeZero",
			v:       Vec4{0, 0, 0, 0},
			want:    Vec4{0, 0, 0, 0},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.v.Normalize()
			if (err != nil) != tt.wantErr {
				t.Errorf("Vec4.Normalize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vec4.Normalize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVec4_Dot(t *testing.T) {
	type args struct {
		v2 Vec4
	}
	tests := []struct {
		name string
		v    Vec4
		args args
		want float64
	}{
		{
			name: "testDot1",
			v:    Vec4{1, 2, 3, 4},
			args: args{v2: Vec4{1, 2, 3, 4}},
			want: 30,
		},
		{
			name: "testDotZero",
			v:    Vec4{0, 0, 0, 0},
			args: args{v2: Vec4{0, 0, 0, 0}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Dot(tt.args.v2); got != tt.want {
				t.Errorf("Vec4.Dot() = %v, want %v", got, tt.want)
			}
		})
	}
}
