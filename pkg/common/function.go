package common

import "github.com/rehylas/wx/pkg/utils"

// 01:00:00~08:56:00
// 11:31:00~13:30:00
// 15:00:00~20:56:00

func CheckBSTime() bool {
	dt := utils.TimeStr()
	if dt >= "01:00:00" && dt <= "08:56:00" {
		return false
	}
	if dt >= "11:31:00" && dt <= "13:30:00" {
		return false
	}
	if dt >= "15:00:00" && dt <= "20:56:00" {
		return false
	}
	return true
}

func CheckBSTimeSymbol(Symbol string) bool {
	return false
}
