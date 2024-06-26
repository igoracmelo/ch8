package main

import (
	"reflect"
	"slices"
	"testing"
)

func Test_drwVxVyN(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		c8 := newChip8()
		c8.i = 0x500
		c8.ram[0x500] = 0b10111001
		c8.v[0] = 0
		c8.v[1] = 0

		c8.drwVxVyN(0, 1, 1)

		got := c8.screen[0][:8]
		want := []bool{true, false, true, true, true, false, false, true}
		if !slices.Equal(got, want) {
			t.Fatalf("\nwant: %v\ngot: %v\n", want, got)
		}
	})

	t.Run("multiple", func(t *testing.T) {
		c8 := newChip8()
		c8.i = 0x500
		c8.ram[0x500] = 0b11111111
		c8.ram[0x501] = 0b10000001
		c8.v[0] = 0
		c8.v[1] = 0

		c8.drwVxVyN(0, 1, 2)

		got := [][]bool{
			c8.screen[0][:8],
			c8.screen[1][:8],
		}
		want := [][]bool{
			{true, true, true, true, true, true, true, true},
			{true, false, false, false, false, false, false, true},
		}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("\nwant: %v\ngot: %v\n", want, got)
		}
	})

	t.Run("multiple", func(t *testing.T) {
		c8 := newChip8()
		c8.i = 0x500
		c8.ram[0x500] = 0b10001101
		c8.ram[0x501] = 0b10000001
		c8.v[0] = 8
		c8.v[1] = 0

		c8.drwVxVyN(0, 1, 2)

		got := c8.screen[0][8:16]
		want := []bool{true, false, false, false, true, true, false, true}

		if !reflect.DeepEqual(got, want) {
			t.Fatalf("\nwant: %v\ngot: %v\n", want, got)
		}
	})
}

func Test_ldBVx(t *testing.T) {
	t.Parallel()

	t.Run("ge 100", func(t *testing.T) {
		c8 := newChip8()
		c8.i = 0
		c8.v[5] = 123
		c8.ldBVx(5)

		want := []uint8{1, 2, 3}
		got := c8.ram[c8.i : c8.i+3]

		if !reflect.DeepEqual(want, got) {
			t.Fatalf("want: %v, got: %v", want, got)
		}
	})

	t.Run("ge 10", func(t *testing.T) {
		c8 := newChip8()
		c8.i = 0
		c8.v[5] = 54
		c8.ldBVx(5)

		want := []uint8{0, 5, 4}
		got := c8.ram[c8.i : c8.i+3]

		if !reflect.DeepEqual(want, got) {
			t.Fatalf("want: %v, got: %v", want, got)
		}
	})

	t.Run("lt 10", func(t *testing.T) {
		c8 := newChip8()
		c8.i = 0
		c8.v[5] = 7
		c8.ldBVx(5)

		want := []uint8{0, 0, 7}
		got := c8.ram[c8.i : c8.i+3]

		if !reflect.DeepEqual(want, got) {
			t.Fatalf("want: %v, got: %v", want, got)
		}
	})
}
