package main

import (
	gobmi "github.com/armstrongli/go-bmi"
	"learn.go/savefateraterank/pkg/apis"

	//"learn/pkg/apis"
	"log"
)

type Calc struct {
}

func (c *Calc) BMI(person *apis.PersonInformation) (float64, error) {
	bmi, err := gobmi.BMI(float64(person.Weight), float64(person.Tall))
	if err != nil {
		log.Println("error when calculating bmi:", err)
		return -1, err
	}
	return bmi, err
}
func (c *Calc) FatRate(person *apis.PersonInformation) (float64, error) {
	bmi, _ := c.BMI(person)
	return gobmi.CalcFatRate(int(person.Age), bmi, person.Sex), nil

}
