 /*
 provides the zLog programming interface to the Go language.
 Copyright (C) 2020 JA1ZLO.
*/
package main

import (
	_ "embed"
	"strings"
	"github.com/nextzlog/zylo"
)

var (
	hsmults int
)


//go:embed hs.dat
var hslist string

func zcities() string {
	return hslist
}


func zlaunch() {
	zylo.Notify("CQ!")
}

func zfinish() {
	zylo.Notify("Bye")
}

func zattach(test string, path string) {
	hsmults=0
}

func zdetach() {
	hsmults=0
}

func zinsert(qso *zylo.QSO) {
	if qso.GetMul2() == "HS" {
		hsmults = hsmults +1 
	}	
}

func zdelete(qso *zylo.QSO) {
	if qso.GetMul2() == "HS" {
		hsmults = hsmults -1 
	}
}

func zverify(qso *zylo.QSO) {
	//multi
	rcvd := strings.TrimSpace(qso.GetRcvd())
	if rcvd != "" {
		l_rcvd := len(rcvd)
	
		if rcvd[l_rcvd-1:l_rcvd]=="C"{
			qso.SetMul1(rcvd[0:l_rcvd-1])
			qso.SetMul2(rcvd[l_rcvd-1:l_rcvd])
		}
		if rcvd[l_rcvd-2:l_rcvd]=="HS"{
			qso.SetMul1(rcvd[0:l_rcvd-2])
			qso.SetMul2(rcvd[l_rcvd-2:l_rcvd])
		}
	}
	
	
	//score
	mode := int(qso.Mode)
	if qso.Dupe {
		qso.Score = 0
	} else {
		if mode == 0 {
			qso.Score = 3
		} else {
			qso.Score = 1
		}
	}
}

func zpoints(score, mults int) int {
	return score * (mults + hsmults)
}

func zeditor(key int, name string) bool {
	return	false
}

func zbutton(btn int, name string) bool {
	return false
}

func main() {}
