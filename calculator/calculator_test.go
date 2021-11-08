//go:build !integration
// +build !integration

package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isEven(t *testing.T) {
	t.Parallel()

	type args struct {
		n uint64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "true",
			args: args{
				n: 4,
			},
			want: true,
		},
		{
			name: "false",
			args: args{
				n: 5,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isEven(tt.args.n); got != tt.want {
				t.Errorf("isEven() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isOdd(t *testing.T) {
	t.Parallel()

	type args struct {
		n uint64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "true",
			args: args{
				n: 5,
			},
			want: true,
		},
		{
			name: "false",
			args: args{
				n: 4,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isOdd(tt.args.n); got != tt.want {
				t.Errorf("isOdd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculate(t *testing.T) {
	t.Parallel()

	type args struct {
		n     uint64
		steps int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "92",
			args: args{
				n:     92,
				steps: 0,
			},
			want: 19,
		},
		{
			name: "98",
			args: args{
				n:     98,
				steps: 0,
			},
			want: 27,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculate(tt.args.n, tt.args.steps); got != tt.want {
				t.Errorf("calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateCollatzConjectureSteps(t *testing.T) {
	t.Parallel()

	var n uint64 = 92
	ch := make(chan map[uint64]int, 1)

	CalculateCollatzConjectureSteps(n, ch)
	close(ch)

	assert.Equal(t, map[uint64]int{92: 19}, <-ch)
}
