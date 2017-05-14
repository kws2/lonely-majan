package main

import "testing"

func getTehai() Tehai {
	haipai := []Pai{
		M1,
		M2,
		M3,
		M4,
		M5,
		M6,
		M7,
		M8,
		M9,
		P1,
		P1,
		S2,
		S3}
	tehai := Tehai{tehai: haipai}
	return tehai
}

func TestIsAgariMentsu(t *testing.T) {
	tehai := getTehai()
	assertAgari(t, tehai, S1)
	assertNotAgari(t, tehai, S2)
	assertNotAgari(t, tehai, S3)
	assertAgari(t, tehai, S4)
}

func getTehai2() Tehai {
	haipai := []Pai{
		M2,
		M3,
		M3,
		M3,
		M3,
		M4,
		M5,
		M6,
		M7,
		M7,
		M7,
		M7,
		M8}
	tehai := Tehai{tehai: haipai}
	return tehai
}

func TestIsAgariMentsu2(t *testing.T) {
	tehai := getTehai2()
	assertAgari(t, tehai, M1)
	assertAgari(t, tehai, M2)
	// assertAgari(t, tehai, M3)
	assertAgari(t, tehai, M4)
	assertAgari(t, tehai, M5)
	assertAgari(t, tehai, M6)
	// assertAgari(t, tehai, M7)
	assertAgari(t, tehai, M8)
	assertAgari(t, tehai, M9)
}

func getTehai3() Tehai {
	haipai := []Pai{
		M2,
		M2,
		M3,
		M3,
		M4,
		M4,
		P5,
		P5,
		S3,
		S4,
		S4,
		S5,
		S5}
	tehai := Tehai{tehai: haipai}
	return tehai
}

func TestIsAgariMentsu3(t *testing.T) {
	tehai := getTehai3()
	// assertAgari(t, tehai, S3)
	assertAgari(t, tehai, S6)
}

func getTehaiTyuren() Tehai {
	haipai := []Pai{
		M1,
		M1,
		M1,
		M2,
		M3,
		M4,
		M5,
		M6,
		M7,
		M8,
		M9,
		M9,
		M9}
	tehai := Tehai{tehai: haipai}
	return tehai
}

func TestIsAgariTyuren(t *testing.T) {
	tehai := getTehaiTyuren()
	assertAgari(t, tehai, M1)
	assertAgari(t, tehai, M2)
	assertAgari(t, tehai, M3)
	assertAgari(t, tehai, M4)
	assertAgari(t, tehai, M5)
	assertAgari(t, tehai, M6)
	assertAgari(t, tehai, M7)
	assertAgari(t, tehai, M8)
	assertAgari(t, tehai, M9)
}

func getTehaiKokushimusou() Tehai {
	haipai := []Pai{
		M1,
		M9,
		P1,
		P9,
		S1,
		S9,
		TON,
		NAN,
		SHA,
		PEI,
		HAK,
		HAT,
		CHN}
	tehai := Tehai{tehai: haipai}
	return tehai
}

func TestIsAgariKokushimusou(t *testing.T) {
	tehai := getTehaiKokushimusou()
	assertAgari(t, tehai, M1)
	assertAgari(t, tehai, M9)
	assertAgari(t, tehai, P1)
	assertAgari(t, tehai, P9)
	assertAgari(t, tehai, S1)
	assertAgari(t, tehai, S9)
	assertAgari(t, tehai, TON)
	assertAgari(t, tehai, NAN)
	assertAgari(t, tehai, SHA)
	assertAgari(t, tehai, PEI)
	assertAgari(t, tehai, HAK)
	assertAgari(t, tehai, HAT)
	assertAgari(t, tehai, CHN)
	assertNotAgari(t, tehai, M2)
	assertNotAgari(t, tehai, P5)
	assertNotAgari(t, tehai, S8)
}

func getTehai7Toitsu() Tehai {
	haipai := []Pai{
		M1,
		M1,
		P1,
		P1,
		S1,
		S1,
		TON,
		TON,
		NAN,
		NAN,
		SHA,
		SHA,
		PEI}
	tehai := Tehai{tehai: haipai}
	return tehai
}
func TestIsAgari7Toitsu(t *testing.T) {
	tehai := getTehai7Toitsu()
	assertAgari(t, tehai, PEI)
	assertNotAgari(t, tehai, S2)
	assertNotAgari(t, tehai, S3)
}

func assertAgari(t *testing.T, tehai Tehai, tsumo Pai) {
	tehai.tsumo = tsumo
	if !tehai.isAgari() {
		t.Fail()
	}
}
func assertNotAgari(t *testing.T, tehai Tehai, tsumo Pai) {
	tehai.tsumo = tsumo
	if tehai.isAgari() {
		t.Fail()
	}
}
