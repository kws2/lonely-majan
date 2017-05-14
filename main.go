package main

import (
	"fmt"
	"os"

	runewidth "github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

type LMJ struct {
	yama    Yama
	tehai   Tehai
	sutehai []Pai
	fail    bool
	success bool
}

var lmj LMJ

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc)

	lmj.yama = Yama{}.Create()
	lmj.tehai = lmj.yama.haipai()
	lmj.tehai.add(&lmj.yama)
	redraw_all()

mainloop:
	for {

		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				break mainloop
			default:
				if ev.Ch != 0 {
					if !lmj.success && !lmj.fail {
						choice(ev.Ch)
					}
				}
			}
		case termbox.EventError:
			panic(ev.Err)
		}
		redraw_all()
	}
}

func write(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x += runewidth.RuneWidth(c)
	}
}

const coldef = termbox.ColorDefault

func redraw_all() {

	termbox.Clear(coldef, coldef)
	w, h := termbox.Size()

	midy := h / 2
	midx := (w - 20) / 2

	ss := showTehai(lmj.tehai)
	tsumo, _ := lmj.tehai.tsumo.string()

	sutehai := stringSutehai(lmj.sutehai)

	// write(midx, midy-1, coldef, coldef, "12345678901234567890123456789012345678901234567890")

	// 捨牌
	write(midx-25, midy-3, coldef, coldef, sutehai)

	// 手牌
	write(midx-12, midy, coldef, coldef, ss)
	write(midx+42, midy, coldef, coldef, tsumo)
	if lmj.success {
		// ESC
		write(midx-6, midy+3, coldef, coldef, "Game Clear!!! Press ESC to quit")
	} else if lmj.fail {
		// ESC
		write(midx-6, midy+3, coldef, coldef, "失敗!!! Press ESC to quit")
	} else {
		// key
		write(midx-12, midy+1, coldef, coldef, "A   B   C   D   E   F   G   H   I   J   K   L   M     N")
		// ESC
		write(midx-6, midy+3, coldef, coldef, "Press ESC to quit")
	}
	// ヤマのデバッグ
	// write(midx-80, midy+5, coldef, coldef, logyama(lmj.yama.Yama[0:80]))
	// if 80 <= len(lmj.yama.Yama) {
	// 	write(midx-80, midy+6, coldef, coldef, logyama(lmj.yama.Yama[80:]))
	// }

	// 残数
	write(midx, midy+7, coldef, coldef, fmt.Sprintf("残:%v", paiLimit()))

	termbox.Flush()
}

func paiLimit() int {
	var wanpaiSize = 12
	var hoge = 80
	return len(lmj.yama.Yama) - wanpaiSize - hoge
}

func logyama(cards []Pai) string {
	var str = ""
	for _, pai := range cards {
		s, err := pai.string()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		str += s
	}
	return str
}

func showTehai(tehai Tehai) string {
	s, err := tehai.string()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return s
}

func stringSutehai(pais []Pai) string {
	var str string
	for _, p := range pais {
		s, _ := p.string()
		str += s
	}
	return str
}

func choice(r rune) {

	var drop Pai
	s := string(r)
	switch s {
	case "a":
		drop = lmj.tehai.drop(0)
	case "b":
		drop = lmj.tehai.drop(1)
	case "c":
		drop = lmj.tehai.drop(2)
	case "d":
		drop = lmj.tehai.drop(3)
	case "e":
		drop = lmj.tehai.drop(4)
	case "f":
		drop = lmj.tehai.drop(5)
	case "g":
		drop = lmj.tehai.drop(6)
	case "h":
		drop = lmj.tehai.drop(7)
	case "i":
		drop = lmj.tehai.drop(8)
	case "j":
		drop = lmj.tehai.drop(9)
	case "k":
		drop = lmj.tehai.drop(10)
	case "l":
		drop = lmj.tehai.drop(11)
	case "m":
		drop = lmj.tehai.drop(12)
	case "n":
		drop = lmj.tehai.dropTsumo()
	default:
		drop = lmj.tehai.dropTsumo()
	}
	lmj.sutehai = append(lmj.sutehai, drop)
	lmj.tehai.add(&lmj.yama)
	if lmj.tehai.isAgari() {
		lmj.success = true
		redraw_all()
	}
	if paiLimit() <= 0 {
		lmj.fail = true
		redraw_all()
	}
}
