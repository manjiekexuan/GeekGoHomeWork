package suggest

func Suggest(Sex string, Age float64, bfrCalcResult float64) string {
	pianshou := "偏瘦"
	biaozhun := "标准"
	pianzhong := "偏重"
	fat := "肥胖"
	yanzhongfeipang := "严重肥胖"
	superfat := "超级肥胖"
	Errorinput := "输入错误"
	age18 := "小于18岁"

	switch Sex {
	case "男":
		if Age >= 18 && Age <= 39 {
			if bfrCalcResult >= 0 && bfrCalcResult <= 0.1 {
				return pianshou
			} else if bfrCalcResult >= 0.1 && bfrCalcResult <= 0.16 {
				return biaozhun
			} else if bfrCalcResult >= 0.17 && bfrCalcResult <= 0.21 {
				return pianzhong
			} else if bfrCalcResult >= 0.22 && bfrCalcResult <= 0.26 {
				return fat
			} else if bfrCalcResult >= 0.27 && bfrCalcResult <= 0.45 {
				return yanzhongfeipang
			} else if bfrCalcResult >= 0.46 {
				return superfat
			} else {
				return Errorinput
			}
		} else if Age >= 40 && Age <= 59 {
			if bfrCalcResult >= 0 && bfrCalcResult <= 0.11 {
				return pianshou
			} else if bfrCalcResult >= 0.12 && bfrCalcResult <= 0.17 {
				return biaozhun
			} else if bfrCalcResult >= 0.18 && bfrCalcResult <= 0.22 {
				return pianzhong
			} else if bfrCalcResult >= 0.23 && bfrCalcResult <= 0.27 {
				return fat
			} else if bfrCalcResult >= 0.28 && bfrCalcResult <= 0.45 {
				return yanzhongfeipang
			} else if bfrCalcResult >= 0.46 {
				return superfat
			} else {
				return Errorinput
			}
		} else if Age >= 60 {
			if bfrCalcResult >= 0 && bfrCalcResult <= 0.13 {
				return pianshou
			} else if bfrCalcResult >= 0.14 && bfrCalcResult <= 0.19 {
				return biaozhun
			} else if bfrCalcResult >= 0.20 && bfrCalcResult <= 0.24 {
				return pianzhong
			} else if bfrCalcResult >= 0.25 && bfrCalcResult <= 0.29 {
				return fat
			} else if bfrCalcResult >= 0.30 && bfrCalcResult <= 0.45 {
				return yanzhongfeipang
			} else if bfrCalcResult >= 0.46 {
				return superfat
			} else {
				return Errorinput
			}
		} else {
			return age18
		}
	case "女":
		if Age >= 18 && Age <= 39 {
			if bfrCalcResult >= 0 && bfrCalcResult <= 0.20 {
				return pianshou
			} else if bfrCalcResult >= 0.21 && bfrCalcResult <= 0.27 {
				return biaozhun
			} else if bfrCalcResult >= 0.28 && bfrCalcResult <= 0.34 {
				return pianzhong
			} else if bfrCalcResult >= 0.35 && bfrCalcResult <= 0.39 {
				return fat
			} else if bfrCalcResult >= 0.40 && bfrCalcResult <= 0.45 {
				return yanzhongfeipang
			} else if bfrCalcResult >= 0.46 {
				return superfat
			} else {
				return Errorinput
			}
		} else if Age >= 40 && Age <= 59 {
			if bfrCalcResult >= 0 && bfrCalcResult <= 0.21 {
				return pianshou
			} else if bfrCalcResult >= 0.22 && bfrCalcResult <= 0.28 {
				return biaozhun
			} else if bfrCalcResult >= 0.29 && bfrCalcResult <= 0.35 {
				return pianzhong
			} else if bfrCalcResult >= 0.36 && bfrCalcResult <= 0.40 {
				return fat
			} else if bfrCalcResult >= 0.40 && bfrCalcResult <= 0.45 {
				return yanzhongfeipang
			} else if bfrCalcResult >= 0.46 {
				return superfat
			} else {
				return Errorinput
			}
		} else if Age >= 60 {
			if bfrCalcResult >= 0 && bfrCalcResult <= 0.22 {
				return pianshou
			} else if bfrCalcResult >= 0.23 && bfrCalcResult <= 0.29 {
				return biaozhun
			} else if bfrCalcResult >= 0.30 && bfrCalcResult <= 0.36 {
				return pianzhong
			} else if bfrCalcResult >= 0.37 && bfrCalcResult <= 0.41 {
				return fat
			} else if bfrCalcResult >= 0.42 && bfrCalcResult <= 0.45 {
				return yanzhongfeipang
			} else if bfrCalcResult >= 0.46 {
				return superfat
			} else {
				return Errorinput
			}
		} else {
			return age18
		}
	default:
		return Errorinput
	}
}
