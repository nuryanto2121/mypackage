package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	type args struct {
		a int
		b int
	}

	test := []struct {
		name      string
		args      args
		want      int
		isPositif bool
	}{
		//test case sukses
		{
			name:      "should return ok",
			args:      args{5, 5},
			want:      10,
			isPositif: true,
		},
		//test case sukses
		{
			name:      "should return failed",
			args:      args{4, 2},
			want:      6,
			isPositif: false,
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			got := Add(tt.args.a, tt.args.b)
			assert.Equal(t, tt.want, got)

		})
	}

}

func TestSubtract(t *testing.T) {
	type args struct {
		a int
		b int
	}

	test := []struct {
		name      string
		args      args
		want      int
		isPositif bool
	}{
		//test case sukses
		{
			name:      "should return ok",
			args:      args{5, 5},
			want:      0,
			isPositif: true,
		},
		//test case sukses
		{
			name:      "should return failed",
			args:      args{4, 2},
			want:      2,
			isPositif: false,
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			got := Subtract(tt.args.a, tt.args.b)
			if got == tt.want {
				assert.Equal(t, tt.want, got)
			} else {
				t.Errorf("GetLevel() = %v, want %v", got, tt.want)
			}

		})
	}

}
