package foobarqix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompute(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "abc => abc",
			input: "abc",
			want:  "abc",
		},
		{
			name:  "1 => 1",
			input: "1",
			want:  "1",
		},
		{
			name:  "2 => 2",
			input: "2",
			want:  "2",
		},
		{
			name:  "3 => FooFoo",
			input: "3",
			want:  "FooFoo",
		},
		{
			name:  "4 => 4",
			input: "4",
			want:  "4",
		},
		{
			name:  "5 => BarBar",
			input: "5",
			want:  "BarBar",
		},
		{
			name:  "6 => Foo",
			input: "6",
			want:  "Foo",
		},
		{
			name:  "7 => QixQix",
			input: "7",
			want:  "QixQix",
		},
		{
			name:  "8 => 8",
			input: "8",
			want:  "8",
		},
		{
			name:  "9 => Foo",
			input: "9",
			want:  "Foo",
		},
		{
			name:  "10 => Bar",
			input: "10",
			want:  "Bar",
		},
		{
			name:  "13 => Foo",
			input: "13",
			want:  "Foo",
		},
		{
			name:  "15 => FooBarBar",
			input: "15",
			want:  "FooBarBar",
		},
		{
			name:  "21 => FooQix",
			input: "21",
			want:  "FooQix",
		},
		{
			name:  "33 => FooFooFoo",
			input: "33",
			want:  "FooFooFoo",
		},
		{
			name:  "51 => FooBar",
			input: "51",
			want:  "FooBar",
		},
		{
			name:  "53 => BarFoo",
			input: "53",
			want:  "BarFoo",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Compute(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}

}
