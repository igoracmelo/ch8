package main

import (
	"testing"
)

func Test_process(t *testing.T) {
	c := &chip8{}
	c.process(0x1111) // JP 111
	if c.pc != 0x0111 {
		t.Errorf("pc - want: 0x0111, got: %04X", c.pc)
	}

	c.process(0x2555) // CALL 555
	if c.pc != 0x0555 {
		t.Errorf("pc - want: 0x0555, got: %04X", c.pc)
	}

	c.process(0x00EE) // RET
	if c.pc != 0x0111 {
		t.Errorf("pc - want: 0x0111, got: %04X", c.pc)
	}

	c.pc = 0
	c.process(0x3F00) // SE VF, 00
	if c.pc != 1 {
		t.Errorf("pc - want: 0x0001 got: %04X", c.pc)
	}

	c.pc = 0
	c.process(0x3F05) // SE VF, 05
	if c.pc != 0 {
		t.Errorf("pc - want: 0x0000 got: %04X", c.pc)
	}

	c.pc = 0
	c.process(0x4F09) // SNE VF, 09
	if c.pc != 1 {
		t.Errorf("pc - want: 0x0001 got: %04X", c.pc)
	}

	c.pc = 0
	c.process(0x4F00) // SNE VF, 00
	if c.pc != 0 {
		t.Errorf("pc - want: 0x0000 got: %04X", c.pc)
	}

	c.pc = 0
	c.process(0x5AB0) // SE VA, VB
	if c.pc != 1 {
		t.Errorf("pc - want: 0x0001 got: %04X", c.pc)
	}

	c.pc = 0
	c.v[0xA] = 5
	c.v[0xB] = 2
	c.process(0x9AB0) // SNE VA, VB
	if c.pc != 1 {
		t.Errorf("pc - want: 0x0001 got: %04X", c.pc)
	}

	c.process(0x6533) // LD V5, 33
	if c.v[5] != 0x33 {
		t.Errorf("V5 - want: 0x33 got: %02X", c.v[5])
	}

	c.v[7] = 0
	c.process(0x7722) // ADD V7, 22
	if c.v[7] != 0x22 {
		t.Errorf("V7 - want: 0x22 got: %02X", c.v[7])
	}

	c.i = 10
	c.process(0xA333) // LD I, 333
	if c.i != 0x333 {
		t.Errorf("I - want: 0x0333 got: %04X", c.i)
	}

	c.v[0] = 5
	c.process(0xB444) // JP 444 + V0
	if c.pc != 0x449 {
		t.Errorf("I - want: 0x0449 got: %04X", c.pc)
	}

}
