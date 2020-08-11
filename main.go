// factorz
package main

import (
	"bytes"
	"fmt"

	"math/big"
	"math/rand"
	"time"
)

var letters = []rune("0123456789")

const (
	MOD1 int = 1
	DIV1 int = 0
)

func randSeq(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	if b[0] == 48 {
		return randSeq(n)
	}
	return string(b)
}
func genstr(s string, len int) string {
	var buffer bytes.Buffer
	for i := 0; i < len; i++ {
		buffer.WriteString(s)
	}
	return buffer.String()
}
func muldiv0(s1 string, s2 string, st int) string {
	n1 := new(big.Int)
	n1.SetString(s1, 10)
	n2 := new(big.Int)
	n2.SetString(s2, 10)
	result := new(big.Int)
	if st == MOD1 {
		result = result.Mod(n1, n2)
	} else if st == DIV1 {
		result = result.Div(n1, n2)
	}
	return result.String()
}
func rnd(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
func fact(s1 string, s2 string) {
	_ = time.Now()
	var n0, n1 string
	var str1, str2, result string
	str1 = s1
	str2 = s2
	for {
		n0 = muldiv0(str1, str2, MOD1)
		if n0 == "1" {
			str2 = randSeq(rnd(len(s1)/4, len(s1)-1))
			continue
		}
		if n0 != "0" {
			str2 = n0
			continue
		} else {
			result = muldiv0(str1, str2, DIV1)
			if result != "1" {
				result1 := muldiv0(s1, result, DIV1)
				fmt.Println("множитель: ", result1)
			}
			if result == "1" {
				str2 = randSeq(rnd(len(s1)/4, len(s1)-1))

				continue
			}
			str1 = result
			str2 = genstr("1", len(str1)-1)
			continue
		}
		_ = n1

	}
}

func main() {
	//	st := "25628048988421719547097505269140301621245447773978588029756396560670223499264"
	//	st := "58207747338624383418022883327000372892939975349820165606278788294480903138594178479250428119336894263454265872902945313285760578431825384789241881237034625"
	//	st := "1571910308760291888984793843057738587262697570697437503743890553816212418123888297554027376163831614110711571968215997115861578098656795678879116492531711577984647461349406872473821577628291279146501792431284752859886076824444143521883476794010257385713389786108553298224709233626535494072055950711316337145958969194230104931060033797622789904057700258800967374017699173059895299330311955038468909660420665504234747838828312926352360431525203393346458725181950081080010425160853488193728053695936990810617343017555208966832145813099976949389366251549035473187255502278087216486961845995957928374753570432887685874768818349206158301313998559081509487044323046777727386244"
	st := "12345678901234567890"
	fact(st, genstr("7", len(st)-2))
	//-----
	len11 := len(st) - 2
	var buffer bytes.Buffer
	for i := 1; i < len11+1; i++ {
		buffer.WriteString(fmt.Sprintf("%d", i))
	}
	fmt.Println(buffer.String())
	//
	fmt.Println("stop")
}
