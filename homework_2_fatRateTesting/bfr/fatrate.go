package bfr

import (
	"fmt"
	"unicode/utf8"
)

func CalcFatRate(bmi float64, age int, sex string) (fatRate float64, err error) {
	sexWeight := 0
	sexInLen := utf8.RuneCountInString(sex)

	if age > 150 {
		err = fmt.Errorf("年龄:%v不能大于150。", age)
		return
	}

	if sexInLen == 0 || sexInLen > 1 {
		err = fmt.Errorf("性别:%v,非法，请输入男|女\n", sex)
		return
	} else {
		if sex == "女" {
			sexWeight = 0
		}
		if sex == "男" {
			sexWeight = 1
		}
	}

	if bmi < 0 {
		err = fmt.Errorf("BMI=%v,入不能为负数。", bmi)
		return
	}

	if bmi == float64(0) {
		err = fmt.Errorf("输入值BMI=%v,但期望输入不能为0。", bmi)
		return
	}

	if age < 18 {
		err = fmt.Errorf("输入年龄为:%v,不能小于18岁。", age)
		return
	}

	fatRate = (1.2*bmi + 0.23*float64(age) - 5.4 - (10.8 * float64(sexWeight))) / 100
	return
}
