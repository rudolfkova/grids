package vectozavr

import (
	"math"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		x float64
		y float64
	}
	tests := []struct {
		name string
		args args
		want Vec2
	}{
		{
			name: "testMaxFloat64",
			args: args{
				x: 823735,
				y: 8574249,
			},
			want: Vec2{823735, 8574249},
		},
		{
			name: "testMaxFloat64",
			args: args{
				x: math.MaxFloat64,
				y: math.MaxFloat64,
			},
			want: Vec2{math.MaxFloat64, math.MaxFloat64},
		},
		{
			name: "testMinFloat64",
			args: args{
				x: -math.MaxFloat64,
				y: -math.MaxFloat64,
			},
			want: Vec2{-math.MaxFloat64, -math.MaxFloat64},
		},
		{
			name: "testZeroFloat64",
			args: args{
				x: 0,
				y: 0,
			},
			want: Vec2{0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVec2_Add(t *testing.T) {
	type args struct {
		v2 Vec2
	}
	tests := []struct {
		name string
		v    Vec2
		args args
		want Vec2
	}{
		{
			name: "testAddVec2",
			v:    Vec2{1, 2},
			args: args{v2: Vec2{3, 4}},
			want: Vec2{4, 6},
		},
		{
			name: "testMaxFloat64",
			v:    Vec2{math.MaxFloat64 / 2, math.MaxFloat64 / 2},
			args: args{v2: Vec2{math.MaxFloat64 / 2, math.MaxFloat64 / 2}},
			want: Vec2{math.MaxFloat64, math.MaxFloat64},
		},
		{
			name: "testMinFloat64",
			v:    Vec2{-math.MaxFloat64 / 2, -math.MaxFloat64 / 2},
			args: args{v2: Vec2{-math.MaxFloat64 / 2, -math.MaxFloat64 / 2}},
			want: Vec2{-math.MaxFloat64, -math.MaxFloat64},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Add(tt.args.v2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vec2.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVec2_Sub(t *testing.T) {
	type args struct {
		v2 Vec2
	}
	tests := []struct {
		name string
		v    Vec2
		args args
		want Vec2
	}{
		{
			name: "testSubVec2",
			v:    Vec2{1, 2},
			args: args{v2: Vec2{3, 4}},
			want: Vec2{-2, -2},
		},
		{
			name: "testMaxFloat64",
			v:    Vec2{math.MaxFloat64 / 2, math.MaxFloat64 / 2},
			args: args{v2: Vec2{-math.MaxFloat64 / 2, -math.MaxFloat64 / 2}},
			want: Vec2{math.MaxFloat64, math.MaxFloat64},
		},
		{
			name: "testMinFloat64",
			v:    Vec2{-math.MaxFloat64 / 2, -math.MaxFloat64 / 2},
			args: args{v2: Vec2{math.MaxFloat64 / 2, math.MaxFloat64 / 2}},
			want: Vec2{-math.MaxFloat64, -math.MaxFloat64},
		},
		{
			name: "testZeroFloat64",
			v:    Vec2{math.MaxFloat64, math.MaxFloat64},
			args: args{v2: Vec2{math.MaxFloat64, math.MaxFloat64}},
			want: Vec2{0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Sub(tt.args.v2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vec2.Sub() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVec2_Mul(t *testing.T) {
	type args struct {
		num float64
	}
	tests := []struct {
		name string
		v    Vec2
		args args
		want Vec2
	}{
		{
			name: "testMulFloat64",
			v:    Vec2{1, 2},
			args: args{num: 2},
			want: Vec2{2, 4},
		},
		{
			name: "testMaxFloat64",
			v:    Vec2{math.MaxFloat64 / 2, math.MaxFloat64 / 2},
			args: args{num: 2},
			want: Vec2{math.MaxFloat64, math.MaxFloat64},
		},
		{
			name: "testMinFloat64",
			v:    Vec2{-math.MaxFloat64 / 2, -math.MaxFloat64 / 2},
			args: args{num: 2},
			want: Vec2{-math.MaxFloat64, -math.MaxFloat64},
		},
		{
			name: "testZeroFloat64",
			v:    Vec2{math.MaxFloat64, math.MaxFloat64},
			args: args{num: 0},
			want: Vec2{0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Mul(tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vec2.Mul() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVec2_Div(t *testing.T) {
	type args struct {
		num float64
	}
	tests := []struct {
		name    string
		v       Vec2
		args    args
		want    Vec2
		wantErr bool
	}{
		{
			name:    "testDivFloat64",
			v:       Vec2{1, 2},
			args:    args{num: 2},
			want:    Vec2{0.5, 1},
			wantErr: false,
		},
		{
			name:    "testZeroFloat64",
			v:       Vec2{math.MaxFloat64, math.MaxFloat64},
			args:    args{num: 0},
			want:    Vec2{math.MaxFloat64, math.MaxFloat64},
			wantErr: true,
		},
		{
			name:    "testMaxFloat64",
			v:       Vec2{math.MaxFloat64 / 2, math.MaxFloat64 / 2},
			args:    args{num: 0.5},
			want:    Vec2{math.MaxFloat64, math.MaxFloat64},
			wantErr: false,
		},
		{
			name:    "testMinFloat64",
			v:       Vec2{-math.MaxFloat64 / 2, -math.MaxFloat64 / 2},
			args:    args{num: 0.5},
			want:    Vec2{-math.MaxFloat64, -math.MaxFloat64},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.v.Div(tt.args.num)
			if (err != nil) != tt.wantErr {
				t.Errorf("Vec2.Div() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vec2.Div() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVec2_Len(t *testing.T) {
	tests := []struct {
		name    string
		v       Vec2
		want    float64
		wantErr bool
	}{
		{
			name:    "testLen1",
			v:       Vec2{0, 2},
			want:    2,
			wantErr: false,
		},
		{
			name:    "testLen2",
			v:       Vec2{2, 0},
			want:    2,
			wantErr: false,
		},
		{
			name:    "testLen3",
			v:       Vec2{3, 4},
			want:    5,
			wantErr: false,
		},
		{
			name:    "testLenZero",
			v:       Vec2{0, 0},
			want:    0,
			wantErr: false,
		},
		{
			name:    "testLenMax",
			v:       Vec2{math.MaxFloat64, math.MaxFloat64},
			want:    0,
			wantErr: true,
		},

		{
			name:    "testLenMin",
			v:       Vec2{-math.MaxFloat64, -math.MaxFloat64},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.v.Len()
			if (err != nil) != tt.wantErr {
				t.Errorf("Vec2.Len() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Vec2.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVec2_Normalize(t *testing.T) {
	tests := []struct {
		name    string
		v       Vec2
		want    Vec2
		wantErr bool
	}{
		{
			name:    "testNormalize1",
			v:       Vec2{0, 2},
			want:    Vec2{0, 1},
			wantErr: false,
		},
		{
			name:    "testNormalize2",
			v:       Vec2{2, 0},
			want:    Vec2{1, 0},
			wantErr: false,
		},
		{
			name:    "testNormalize3",
			v:       Vec2{3, 4},
			want:    Vec2{0.6, 0.8},
			wantErr: false,
		},
		{
			name:    "testNormalizeZero",
			v:       Vec2{0, 0},
			want:    Vec2{0, 0},
			wantErr: true,
		},
		{
			name:    "testNormalizeMax",
			v:       Vec2{math.MaxFloat64, math.MaxFloat64},
			want:    Vec2{math.MaxFloat64, math.MaxFloat64},
			wantErr: true,
		},
		{
			name:    "testNormalizeMin",
			v:       Vec2{-math.MaxFloat64, -math.MaxFloat64},
			want:    Vec2{-math.MaxFloat64, -math.MaxFloat64},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.v.Normalize()
			if (err != nil) != tt.wantErr {
				t.Errorf("Vec2.Normalize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vec2.Normalize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVec2_Dot(t *testing.T) {
	type args struct {
		v2 Vec2
	}
	tests := []struct {
		name string
		v    Vec2
		args args
		want float64
	}{
		{
			name: "testDot1",
			v:    Vec2{1, 2},
			args: args{v2: Vec2{1, 2}},
			want: 5,
		},
		{
			name: "testDot2",
			v:    Vec2{0, 2},
			args: args{v2: Vec2{2, 0}},
			want: 0,
		},
		{
			name: "testDot3",
			v:    Vec2{2, 0},
			args: args{v2: Vec2{0, 2}},
			want: 0,
		},
		{
			name: "testDot4",
			v:    Vec2{2, -2},
			args: args{v2: Vec2{1, 2}},
			want: -2,
		},
		{
			name: "testDot5",
			v:    Vec2{0, 0},
			args: args{v2: Vec2{0, 0}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Dot(tt.args.v2); got != tt.want {
				t.Errorf("Vec2.Dot() = %v, want %v", got, tt.want)
			}
		})
	}
}
