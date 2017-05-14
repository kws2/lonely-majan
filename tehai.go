package main

import (
	"fmt"
	"sort"
)

type Tehai struct {
	tehai []Pai
	tsumo Pai
}

func (tehai Tehai) string() (string, error) {
	var ret = ""
	t, err := tehai.sort()
	if err != nil {
		return "", err
	}
	for _, pai := range t.tehai {
		s, err := pai.string()
		if err != nil {
			return "", err
		}
		ret += s
		ret += "　"
	}
	return ret, nil
}

func (tehai *Tehai) drop(i int) Pai {
	var ret Pai
	if i < len(tehai.tehai) {
		ret = tehai.tehai[i]
		tehai.tehai = append(tehai.tehai[:i], tehai.tehai[(i+1):]...)
		tehai.tehai = append(tehai.tehai, tehai.tsumo)
	} else {
		ret = tehai.tsumo
	}
	tehai.tsumo = N_PAI
	return ret
}

func (tehai *Tehai) dropTsumo() Pai {
	ret := tehai.tsumo
	tehai.tsumo = N_PAI
	return ret
}

func (tehai *Tehai) add(yama *Yama) Pai {

	// ツモ牌を切ってない
	if tehai.tsumo != N_PAI {
		fmt.Println("ERROR 本当は多牌")
		tehai.tehai = append(tehai.tehai, tehai.tsumo)
	}

	tehai.tsumo = yama.getPai()
	return tehai.tsumo
}

func (tehai Tehai) sort() (Tehai, error) {
	sort.Slice(tehai.tehai, func(i, j int) bool {
		return tehai.tehai[i] < tehai.tehai[j]
	})
	return tehai, nil
}

func (tehai Tehai) getAllTehai() []Pai {
	// ツモ牌も手配に含める
	t := append(tehai.tehai, tehai.tsumo)
	return t
}

func (tehai Tehai) isAgari() bool {

	// ツモってない
	if tehai.tsumo == N_PAI {
		// return [HAI_SHU]Pai{}, fmt.Errorf("err %s", "ERROR 少牌")
		return false
	}
	t := tehai.getAllTehai()

	c := countPaiFromTehai(t)

	// 国士チェック
	if isKokushimusou(c) {
		return true
	}

	// 面子手チェック
	for i := range c {
		tmp := countPaiFromTehai(t)
		if isHead(tmp, i) {
			// 雀頭を削除
			tmp[i] += -2
			for j := range tmp {
				if isAnko(tmp, j) {
					// 暗刻を削除
					tmp[j] += -3
				}
			}
			for k := range tmp {
				for tmp[k] != 0 {
					if isShuntsu(tmp, k) {
						tmp[k] += -1
						tmp[k+1] += -1
						tmp[k+2] += -1
					} else {
						break
					}
				}
			}
			sum := 0
			for _, x := range tmp {
				sum += x
			}
			if sum == 0 {
				return true
			}
		} else {
			continue
		}
	}

	// 七対子チェック
	if is7Toitsu(c) {
		return true
	}

	return false
}

// []Pai(手牌)から[34]int(34種の牌をいくつ持っているか)に変換する
// 例: 萬子の11112345678999の場合
// []Pai   : [0 0 0 0 1 2 3 4 5 6 7 8 8 8]
// [34]int : [4 1 1 1 1 1 1 1 3 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
func countPaiFromTehai(pais []Pai) [HAI_SHU]int {
	var c [HAI_SHU]int
	for _, pai := range pais {
		c[pai]++
	}
	return c
}

func isHead(pais [HAI_SHU]int, i int) bool {
	if pais[i] >= 2 {
		return true
	}
	return false
}

func isAnko(pais [HAI_SHU]int, i int) bool {
	if pais[i] >= 3 {
		return true
	}
	return false
}

func isShuntsu(pais [HAI_SHU]int, i int) bool {
	p := Pai(i)
	// 字牌は順子ではない
	if p.Type() != MAN && p.Type() != PIN && p.Type() != SOU {
		return false
	}
	// 数牌の8以上はチェック対象外
	if p == M8 || p == M9 || p == P8 || p == P9 || p == S8 || p == S9 {
		return false
	}
	if pais[i] >= 1 && pais[i+1] >= 1 && pais[i+2] >= 1 {
		return true
	}
	return false
}

func isKokushimusou(pais [HAI_SHU]int) bool {
	for i, count := range pais {
		if isYaotyu(Pai(i)) {
			if count == 0 {
				return false
			}
		} else {
			if count > 0 {
				return false
			}
		}
	}
	return true
}

func is7Toitsu(pais [HAI_SHU]int) bool {
	for _, count := range pais {
		if count == 0 || count == 2 {
			continue
		}
		return false
	}
	return true
}
