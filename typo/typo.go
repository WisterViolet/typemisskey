// Package typo : Typoを生成する何か
package typo

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type keyPoint struct {
	ShiftFlag bool
	Row       int
	Column    int
}

type vec struct {
	X int
	Y int
}

func newKeyPoint(r, c int, s bool) *keyPoint {
	k := new(keyPoint)
	k.Row = r
	k.Column = c
	k.ShiftFlag = s
	return k
}

func newVec(x, y int) *vec {
	v := new(vec)
	v.X = x
	v.Y = y
	return v
}

// Generate : StringからTypoを生成する関数
func Generate(base string) (string, error) {
	/*
		base: 元の文字列
	*/
	if len(strings.Fields(base)) == 0 {
		err := errors.New("error! whitespace only")
		return "", err
	}
	keymap, keymapShift := make([]string, 4), make([]string, 4)
	keymap[0] = `1234567890-^\`
	keymap[1] = `qwertyuiop@[`
	keymap[2] = `asdfghjkl;:]`
	keymap[3] = `zxcvbnm,./\`
	keymapShift[0] = `!"#$%&'() ~=~|`
	keymapShift[1] = "QWERTYUIOP`{"
	keymapShift[2] = `ASDFGHJKL+*}`
	keymapShift[3] = `ZXCVBNM<>?_ `
	space := " "
	rand.Seed(time.Now().UnixNano())
	typoIndex := rand.Intn(len(base))
	for string(base[typoIndex]) == space {
		fmt.Println(string(base[typoIndex]))
		typoIndex = rand.Intn(len(base))
	}
	point := newKeyPoint(0, 0, false)
	for i := 0; i < 4; i++ {
		if p := strings.Index(keymap[i], string(base[typoIndex])); p != -1 {
			point.ShiftFlag = false
			point.Row = i
			point.Column = p
		} else if p := strings.Index(keymapShift[i], string(base[typoIndex])); p != -1 {
			point.ShiftFlag = true
			point.Row = i
			point.Column = p
		}
	}
	v := newVec(0, 0)
	switch point.Row {
	case 0:
		v.Y = rand.Intn(2)
	case 3:
		v.Y = rand.Intn(2) - 1
	default:
		v.Y = rand.Intn(3) - 1
	}
	switch l, c := len(keymap[point.Row]), point.Column; true {
	case c == 0:
		v.X = rand.Intn(2)
	case c+1 == l:
		v.X = rand.Intn(2) - 1
	default:
		v.X = rand.Intn(3) - 1
	}
	point.Row += v.Y
	if point.ShiftFlag == true {
		return base[0:typoIndex+1] + string(keymapShift[point.Row][point.Column]) + base[typoIndex+1:len(base)], nil
	}
	return base[0:typoIndex+1] + string(keymap[point.Row][point.Column]) + base[typoIndex+1:len(base)], nil
}
