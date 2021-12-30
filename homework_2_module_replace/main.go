package main

import (
	"fmt"
	gobmi "github.com/armstrongli/go-bmi"
	"github.com/spf13/cobra"
)

var (
	Name   string
	Sex    string
	Height float64 //string
	Weight float64 //string
	Age    int
)

func InputFromCobra() {
	cmd := cobra.Command{
		Use:   "healthcheck",
		Short: "体脂计算器，体重、身高、性别、年龄。",
		Long:  "BMI根据体重身高性别和年龄进行科学运算得到最真实的数据",
		Run: func(cmd *cobra.Command, args []string) {

			fmt.Println("Name:", Name)
			fmt.Println("Sex:", Sex)
			fmt.Println("Height:", Height)
			fmt.Println("Weight:", Weight)
			fmt.Println("Age:", Age)

			// 计算
			bmi, err := gobmi.BMI(Weight, Height)

			if err != nil {
				fmt.Errorf("bmi异常错误信息:%v\n", err)
			}

			bfr := gobmi.CalcFatRate(bmi, Age, Sex)
			fmt.Printf("bmi:%v,bfr:%v\n", bmi, bfr)
		},
	}

	cmd.Flags().StringVar(&Name, "Name", "", "姓名")
	cmd.Flags().StringVar(&Sex, "Sex", "", "姓别")
	cmd.Flags().Float64Var(&Height, "Height", -1, "身高")
	cmd.Flags().Float64Var(&Weight, "Weight", -1, "体重")
	cmd.Flags().IntVar(&Age, "Age", -1, "年龄")

	// 运行命令行对象
	cmd.Execute()
}

func main() {
	//ArgsExample()
	InputFromCobra()
}
