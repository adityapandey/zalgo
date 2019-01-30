package zalgo

import (
	"io"
	"math/rand"
	"strings"
	"unicode/utf8"
)

var (
	upRunes = []rune{
		0x030d, 0x030e, 0x0304, 0x0305,
		0x033f, 0x0311, 0x0306, 0x0310,
		0x0352, 0x0357, 0x0351, 0x0307,
		0x0308, 0x030a, 0x0342, 0x0343,
		0x0344, 0x034a, 0x034b, 0x034c,
		0x0303, 0x0302, 0x030c, 0x0350,
		0x0300, 0x0301, 0x030b, 0x030f,
		0x0312, 0x0313, 0x0314, 0x033d,
		0x0309, 0x0363, 0x0364, 0x0365,
		0x0366, 0x0367, 0x0368, 0x0369,
		0x036a, 0x036b, 0x036c, 0x036d,
		0x036e, 0x036f, 0x033e, 0x035b,
		0x0346, 0x031a,
	}

	midRunes = []rune{
		0x0315, 0x031b, 0x0340, 0x0341,
		0x0358, 0x0321, 0x0322, 0x0327,
		0x0328, 0x0334, 0x0335, 0x0336,
		0x034f, 0x035c, 0x035d, 0x035e,
		0x035f, 0x0360, 0x0362, 0x0338,
		0x0337, 0x0361, 0x0489,
	}

	downRunes = []rune{
		0x0316, 0x0317, 0x0318, 0x0319,
		0x031c, 0x031d, 0x031e, 0x031f,
		0x0320, 0x0324, 0x0325, 0x0326,
		0x0329, 0x032a, 0x032b, 0x032c,
		0x032d, 0x032e, 0x032f, 0x0330,
		0x0331, 0x0332, 0x0333, 0x0339,
		0x033a, 0x033b, 0x033c, 0x0345,
		0x0347, 0x0348, 0x0349, 0x034d,
		0x034e, 0x0353, 0x0354, 0x0355,
		0x0356, 0x0359, 0x035a, 0x0323,
	}
)

var allRunes = make(map[rune]struct{})

func init() {
	for _, r := range upRunes {
		allRunes[r] = struct{}{}
	}
	for _, r := range midRunes {
		allRunes[r] = struct{}{}
	}
	for _, r := range downRunes {
		allRunes[r] = struct{}{}
	}
}

func rnd(n int) int {
	return rand.Int() % n
}

func randZalgo(z []rune) rune {
	return z[rnd(len(z))]
}

type level int

const (
	Min level = iota
	Normal
	Max
)

type Options struct {
	Corruption    level
	Up, Mid, Down bool
}

type Writer struct {
	w   io.Writer
	Opt Options
	b   []byte
}

func (z *Writer) Write(p []byte) (int, error) {
	z.b = append(z.b, p...)
	i := 0
	for len(z.b) > 0 {
		r, size := utf8.DecodeRune(z.b)
		if r == utf8.RuneError {
			return i, io.ErrShortWrite
		}
		i += size
		z.b = z.b[size:]
		if _, ok := allRunes[r]; ok {
			continue
		}

		var s strings.Builder

		if _, err := s.WriteRune(r); err != nil {
			return i, err
		}
		var nUp, nMid, nDown int
		switch z.Opt.Corruption {
		case Min:
			nUp, nMid, nDown = rnd(8), rnd(2), rnd(8)
		case Normal:
			nUp, nMid, nDown = rnd(16)/2+1, rnd(6)/2, rnd(16)/2+1
		case Max:
			nUp, nMid, nDown = rnd(64)/4+3, rnd(16)/4+1, rnd(64)/4+3
		}
		if z.Opt.Up {
			for j := 0; j < nUp; j++ {
				if _, err := s.WriteRune(randZalgo(upRunes)); err != nil {
					return i, err
				}
			}
		}
		if z.Opt.Mid {
			for j := 0; j < nMid; j++ {
				if _, err := s.WriteRune(randZalgo(midRunes)); err != nil {
					return i, err
				}
			}
		}
		if z.Opt.Down {
			for j := 0; j < nDown; j++ {
				if _, err := s.WriteRune(randZalgo(downRunes)); err != nil {
					return i, err
				}
			}
		}
		if _, err := z.w.Write([]byte(s.String())); err != nil {
			return i, err
		}
	}
	return len(p), nil
}

func NewWriter(w io.Writer) *Writer {
	return &Writer{w: w, Opt: Options{Corruption: Min}}
}
