package main

import (
	crand "crypto/rand"
	"encoding/binary"
	"math/rand"
	"time"
)

type Yama struct {
	Yama []Pai
}

func (yama Yama) Create() Yama {
	for i := 0; i < HAI_SHU; i++ {
		for j := 0; j < 4; j++ {
			yama.Yama = append(yama.Yama, Pai(i))
		}
	}
	return yama
}

func (yama *Yama) getPai() Pai {
	n := random(len(yama.Yama))
	pai := yama.Yama[n]

	yama.Yama = append(yama.Yama[:n], yama.Yama[(n+1):]...)

	c := make([]Pai, len(yama.Yama))
	copy(c, yama.Yama)
	return pai
}

func random(max int) int {
	var s int64
	if err := binary.Read(crand.Reader, binary.LittleEndian, &s); err != nil {
		// crypto/rand からReadできなかった場合の代替手段
		s = time.Now().UnixNano()
	}
	rand.Seed(s)
	n := rand.Intn(max)
	return n
}

func (yama *Yama) haipai() Tehai {
	haipai := []Pai{
		yama.getPai(),
		yama.getPai(),
		yama.getPai(),
		yama.getPai(),
		yama.getPai(),
		yama.getPai(),
		yama.getPai(),
		yama.getPai(),
		yama.getPai(),
		yama.getPai(),
		yama.getPai(),
		yama.getPai(),
		yama.getPai()}
	t := Tehai{tehai: haipai, tsumo: N_PAI}
	return t
}

func (yama Yama) string() (string, error) {
	var ret = ""
	for _, pai := range yama.Yama {
		s, err := pai.string()
		if err != nil {
			return "", err
		}
		ret += s
		ret += "　"
	}
	return ret, nil
}
