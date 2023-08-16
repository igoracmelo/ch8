package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := &chip8{}
	fmt.Println(c.render())
}

type chip8 struct {
	// program counter
	pc uint16

	// stack pointer
	sp uint8

	stack [16]uint16

	ram [4096]uint8

	display [32][64]bool

	// V0 - VF registers
	v [16]uint8

	// I register
	i uint16

	// delay timer, decremented at 60Hz
	dt uint8

	// sound timer, decremented at 60Hz and
	// plays buzzer while st != 0
	st uint8

	// period between CPU cycles
	period time.Duration

	// time of last CPU cycle
	lastCycle time.Time

	
}

func (c *chip8) render() string {
	s := ""
	for i := 0; i < len(c.display); i++ {
		for j := 0; j < len(c.display[0]); j++ {
			if c.display[i][j] {
				s += "#"
			} else {
				s += "."
			}
		}
		s += "\n"
	}

	s += "\n"
	for i := 0; i <= 0xF; i++ {
		s += fmt.Sprintf("V%1X: %02X  ", i, c.v[i])	
		if (i+1)%8 == 0 {
			s += "\n"
		}
	}
	s += fmt.Sprintf("PC: %04X  I: %04X  SP: %02X  DT: %02X  ST: %02X", c.pc, c.i, c.sp, c.dt, c.st)

	return s
}

func (c *chip8) step() {
	defer func() {
		c.lastCycle = time.Now()
	}()

	time.Sleep(c.period - time.Since(c.lastCycle))
}

func (c *chip8) process(opcode uint16) {
	switch opcode {
	case 0x00EE: // RET
		c.pc = c.stack[c.sp]
		c.sp--
		return
	case 0x00E0: // CLS
		// TODO
	}

	// unique prefix instructions
	switch opcode & 0xF000 {
	case 0x1000: // JP nnn
		c.pc = opcode & 0x0FFF
		return

	case 0x2000: // CALL nnn
		c.sp++
		c.stack[c.sp] = c.pc
		c.pc = opcode & 0x0FFF
		return

	case 0x3000: // SE Vx, b
		x := (opcode & 0x0F00) >> 8
		b := uint8(opcode & 0x00FF)
		if c.v[x] == b {
			c.pc++
		}
		return

	case 0x4000: // SNE Vx, b
		x := (opcode & 0x0F00) >> 8
		b := uint8(opcode & 0x00FF)
		if c.v[x] != b {
			c.pc++
		}
		return

	case 0x5000: // SE Vx, Vy
		x := (opcode & 0x0F00) >> 8
		y := (opcode & 0x00F0) >> 4
		if c.v[x] == c.v[y] {
			c.pc++
		}
		return

	case 0x9000: // SNE Vx, Vy
		x := (opcode & 0x0F00) >> 8
		y := (opcode & 0x00F0) >> 4
		if c.v[x] != c.v[y] {
			c.pc++
		}
		return

	case 0x6000: // LD Vx, b
		x := (opcode & 0x0F00) >> 8
		b := uint8(opcode & 0x00FF)
		c.v[x] = b
		return

	case 0x7000: // ADD Vx, b
		x := (opcode & 0x0F00) >> 8
		b := uint8(opcode & 0x00FF)
		c.v[x] += b // TODO: wrapping?
		return

	case 0xA000: // LD I, addr
		a := opcode & 0x0FFF
		c.i = a
		return

	case 0xB000: // JP V0, addr
		a := opcode & 0x0FFF + uint16(c.v[0])
		c.pc = a
		return

	case 0xC000: // RND b & Vx
		x := (opcode & 0x0F00) >> 8
		b := uint8(opcode & 0x00FF)
		c.v[x] = uint8(rand.Intn(255)) & b
		return

	case 0xD000: // DRW Vx, Vy, n
		// TODO
		x := (opcode & 0x0F00) >> 8
		y := (opcode & 0x00F0) >> 4
		n := uint8(opcode & 0x0F)
		_, _, _ = x, y, n
		return
	}
}
