package main

import (
	"encoding/json"
	"fmt"
	gobmi "github.com/armstrongli/go-bmi"
	"io/ioutil"
	"learn.go/savefateraterank/pkg/apis"
	"log"
	"os"
)

func main() {
	input := InputFromStd{}
	calc := &Calc{}
	rk := &FatRateRank{}
	records := NewRecord("C:\\Users\\Jack\\Desktop\\Yijie.txt")

	for {
		pi := input.GetInput()
		if err := records.savePersonInformation(pi); err != nil {
			log.Fatal("保存失败", err)
		}
		fr, err := calc.FatRate(pi)
		if err != nil {
			log.Fatal("计算体脂率失败:", err)
		}
		rk.inputRecord(pi.Name, fr)
		fmt.Println(personFatRate)
		//冒泡排序
		//rankResult, _ := rk.getRandByBubble(pi.Name)
		//快速排序
		rankResult, _ := rk.getRandByQuick(pi.Name)

		fmt.Println("排名结果是:", rankResult)
	}

}

func writeFile(filePath string, data []byte) {
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.Write(data)
	personInformation := apis.PersonInformation{}
	json.Unmarshal(data, &personInformation)
	fmt.Println("反序列化:", personInformation)
	bmi, _ := gobmi.BMI(float64(personInformation.Weight), float64(personInformation.Tall))
	fmt.Printf("%s的BMI是%v\n", personInformation.Name, bmi)
	fatRate := gobmi.CalcFatRate(int(personInformation.Age), bmi, personInformation.Sex)
	fmt.Printf("%s的体脂率是%v\n", personInformation.Name, fatRate)
}

func readFile(filePath string) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("读取出来的内容是:", string(data))
	//fmt.Println("开始计算体脂信息:", infos)
}

func caseStudy() {
	filePath := "C:\\Users\\Jack\\Desktop\\Yijie.txt"
	personInformation := apis.PersonInformation{
		Name:   `"戴一杰"`,
		Sex:    "男",
		Tall:   1.75,
		Weight: 75,
		Age:    26,
	}

	fmt.Printf("%v\n", personInformation)

	data, err := json.Marshal(personInformation)
	if err != nil {
		log.Fatal(err)

	}
	fmt.Println("Marshal的结果是(原生):", data)
	fmt.Println("Marshal的结果是(string):", string(data))
	writeFile(filePath, data) //ctrl+alt+m 将方法提取出来
	//writeFileWithAppend(filePath)
	readFile(filePath)
}
