 /*
 provides the zLog programming interface to the Go language.
 Copyright (C) 2020 JA1ZLO.
*/
package main

import (
	_ "embed"
	"strings"
	"os"
)

var (
	hsmults int
)


//go:embed hs.dat
var cityMultiList string

func init() {
	CityMultiList = cityMultiList
	OnLaunchEvent = onLaunchEvent
	OnFinishEvent = onFinishEvent
	OnAssignEvent = onAssignEvent
	OnAttachEvent = onAttachEvent
	OnInsertEvent = onInsertEvent
	OnDeleteEvent = onDeleteEvent
	OnVerifyEvent = onVerifyEvent
	OnPointsEvent = onPointsEvent
}


func onLaunchEvent() {
	DisplayToast("CQ!")
	hsmults = 0
}

func onFinishEvent() {
	DisplayToast("Bye")
	hsmults = 0
}

func onAttachEvent(test string, path string) {
	DisplayToast(test)
}

func onAssignEvent(contest, configs string) {
	binary_data, err := os.ReadFile(Query("{F}"))
	if err == nil {
		initial_qsos := LoadZLO(binary_data)
		for _,qso_data := range initial_qsos{
			if qso_data.GetMul2() == "HS"{
				hsmults = hsmults +1
			}
		}		
	}
}


func onInsertEvent(qso *QSO) {
	if qso.GetMul2() == "HS" {
		hsmults = hsmults +1 
	}	
}

func  onDeleteEvent(qso *QSO) {
	if qso.GetMul2() == "HS" {
		hsmults = hsmults -1 
	}
}

func onVerifyEvent(qso *QSO) {
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

func onPointsEvent(score, mults int) int {
	return score * (mults + hsmults)
}