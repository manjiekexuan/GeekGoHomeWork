package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	var inputStr string
	var yesOrNo  string
	var sexWeight int
	var i int = 0
	var allFateRate float64
	//循环输入次数


	for  {
		fmt.Println("姓名;性别(男,女);身高(米);体重(公斤);年龄 ------> 例如:小治;男;1.81;86.5;30")
		fmt.Scanln(&inputStr)

		input := strings.Split(inputStr,";")
		i++
		fmt.Println(input)
		//判断参数个数是否正确
		if len(input) < 5 || len(input) > 5 {
			fmt.Println("输入的参数个数不正确请重新输入")
		}
		name := input[0]
		sex  := input[1]
		if sex == "男"{
			sexWeight = 1
		}else if sex == "女"{
			sexWeight = 0
		}
		height, err := strconv.ParseFloat(input[2], 64)
		if err != nil {
			fmt.Println("输入的身高不正确请重新输入")
			continue
		}
		weight, err := strconv.ParseFloat(input[3], 64)
		if err != nil {
			fmt.Println("输入的体重不正确请重新输入")
			continue
		}
		age,err:= strconv.ParseInt(input[4],10,0)
		if err != nil {
			fmt.Println("输入的年龄不正确请重新输入")
			continue
		}
		BMI,FatRate,device:= getPersonReport(weight,height,age,sexWeight)
		allFateRate = allFateRate + FatRate
		fmt.Println("姓名:",name,"BMI:",BMI,"体脂率:",FatRate,"建议:",device)
		fmt.Println("是否继续输入是/否(Y/N)")
		fmt.Scanln(&yesOrNo)
		if yesOrNo == "N" || yesOrNo == "n"{
			break
		}
	}
	fmt.Println("已经输入结束,报表如下")
	fmt.Println("总人数为:",i,"平均体脂率为",allFateRate/float64(i))
}


func getPersonReport(weight float64,height float64,age int64, sexWeight int )(float64,float64,string){
	BMI:=getBMI(weight,height)
	fatRate := getFatRate(BMI,age,sexWeight)
	return BMI,fatRate,getDevice(fatRate,age,sexWeight)

}

//获取建议
func getDevice(fatRate float64,age int64,sexWeight int )string{
	var content string
	if sexWeight == 1{
		content =getDeviceContentMan(fatRate,age)
	}else if sexWeight == 2{
		content = getDeviceContentWoman(fatRate,age)
	}
	return content
}

func getDeviceContentMan(fatRate float64,age int64) string {
	var content string
	switch  {
		case fatRate >= 0 && fatRate <=10:
			content = "偏瘦"
			fallthrough
		case fatRate >=30 && fatRate<=45:
			content = "严重肥胖"
			fallthrough
		case 18 <= age && age<= 39:
			if fatRate <= 16{
				content = "标准"
			}else if fatRate <= 21{
				content = "偏重"
			}else if fatRate >= 22 && fatRate <= 26{
				content = "肥胖"
			}
		case 40 <= age && age<= 59:
			if fatRate >= 12 && fatRate <= 17{
				content = "标准"
			}else if fatRate >=18 && fatRate <= 22{
				content = "偏重"
			}else if fatRate >=23 && fatRate<=27{
				content = "肥胖"
			}

		case age >= 60:
			if fatRate >= 14 && fatRate<=19{
				content = "标准"
			}else if fatRate >= 20 && fatRate<= 24{
				content = "偏重"
			}else if fatRate >= 25 && fatRate<=29 {
				content = "肥胖"
			}
	}
	return content
}

func getDeviceContentWoman(fatRate float64,age int64) string{
	var content string
	switch  {
		case fatRate >= 0 && fatRate <=20:
			content = "偏瘦"
			fallthrough
		case 18 <= age && age<= 39:
			if fatRate <= 27{
				content = "标准"
			}else if fatRate <= 34{
				content = "偏重"
			}else if fatRate <= 39{
				content = "肥胖"
			}
		case 40 <= age && age<= 59:
			if fatRate <= 28{
				content = "标准"
			}else if fatRate <= 35{
				content = "偏重"
			}else if fatRate<=40{
				content = "肥胖"
			}
		case age >= 60:
			if fatRate <=22{
				content = "偏瘦"
			} else if fatRate <= 29{
				content = "标准"
			}else if fatRate <= 36{
				content = "偏重"
			}else if  fatRate<=41 {
				content = "肥胖"
			}
	}
	return content
}


func getBMI(weight float64,height float64) float64{
	//bmi := weight / (height * height)
	return weight / (height * height)
}

func getFatRate(BMI float64,age int64,sexWeight int) float64{
	//fatRate := 1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*float64(sexWeight)
	return 1.2*BMI + 0.23*float64(age) - 5.4 - 10.8*float64(sexWeight)
}