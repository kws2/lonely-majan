package main

import "fmt"

const HAI_SHU = 34

type Pai int

const (
	//
	M1 Pai = iota
	M2
	M3
	M4
	M5
	M6
	M7
	M8
	M9
	//
	P1
	P2
	P3
	P4
	P5
	P6
	P7
	P8
	P9
	//
	S1
	S2
	S3
	S4
	S5
	S6
	S7
	S8
	S9
	TON
	NAN
	SHA
	PEI
	HAK
	HAT
	CHN
	N_PAI Pai = -1
)

type PaiType int

const (
	N_PAI_TYPE PaiType = iota
	MAN
	PIN
	SOU
	KAZ
	SAN
)

func (i Pai) Type() PaiType {
	if i >= P1 && i <= P9 {
		return MAN
	} else if i >= M1 && i <= M9 {
		return PIN
	} else if i >= S1 && i <= S9 {
		return SOU
	} else if i >= TON && i <= PEI {
		return KAZ
	} else if i >= HAK && i <= CHN {
		return SAN
	}
	return N_PAI_TYPE
}

func isYaotyu(pai Pai) bool {
	switch pai {
	case M1:
		return true
	case M9:
		return true
	case P1:
		return true
	case P9:
		return true
	case S1:
		return true
	case S9:
		return true
	case TON:
		return true
	case NAN:
		return true
	case SHA:
		return true
	case PEI:
		return true
	case HAK:
		return true
	case HAT:
		return true
	case CHN:
		return true
	default:
		return false
	}
}

func (p Pai) string() (string, error) {
	switch p {
	case M1:
		return "一", nil
	case M2:
		return "二", nil
	case M3:
		return "三", nil
	case M4:
		return "四", nil
	case M5:
		return "五", nil
	case M6:
		return "六", nil
	case M7:
		return "七", nil
	case M8:
		return "八", nil
	case M9:
		return "九", nil
	case P1:
		return "①", nil
	case P2:
		return "②", nil
	case P3:
		return "③", nil
	case P4:
		return "④", nil
	case P5:
		return "⑤", nil
	case P6:
		return "⑥", nil
	case P7:
		return "⑦", nil
	case P8:
		return "⑧", nil
	case P9:
		return "⑨", nil
	case S1:
		return "１", nil
	case S2:
		return "２", nil
	case S3:
		return "３", nil
	case S4:
		return "４", nil
	case S5:
		return "５", nil
	case S6:
		return "６", nil
	case S7:
		return "７", nil
	case S8:
		return "８", nil
	case S9:
		return "９", nil
	case TON:
		return "東", nil
	case NAN:
		return "南", nil
	case SHA:
		return "西", nil
	case PEI:
		return "北", nil
	case HAK:
		return "白", nil
	case HAT:
		return "發", nil
	case CHN:
		return "中", nil
	default:
		return "", fmt.Errorf("err %s", "想定外の牌")
	}
}
