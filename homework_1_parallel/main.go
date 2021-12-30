package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var AB,CD string
	var AX,AY,BX,BY,CX,CY,DX,DY,A,B,C,D string

	fmt.Println("请输入AB线的两点坐标,例52,23;99,77")
	fmt.Scanln(&AB)
	A = strings.Split(AB,";")[0]
	B = strings.Split(AB,";")[1]
	AX = strings.Split(A,",")[0]
	AY = strings.Split(A,",")[1]
	BX = strings.Split(B,",")[0]
	BY = strings.Split(B,",")[1]


	fmt.Println("请输入CD线的两点坐标")
	fmt.Scanln(&CD)
	C = strings.Split(CD,";")[0]
	D = strings.Split(CD,";")[1]
	CX = strings.Split(C,",")[0]
	CY = strings.Split(C,",")[1]
	DX = strings.Split(D,",")[0]
	DY = strings.Split(D,",")[1]


	ax,_:= strconv.ParseFloat(AX,64)
	ay,_:= strconv.ParseFloat(AY,64)
	bx,_:= strconv.ParseFloat(BX,64)
	by,_:= strconv.ParseFloat(BY,64)
	cx,_:= strconv.ParseFloat(CX,64)
	cy,_:= strconv.ParseFloat(CY,64)
	dx,_:= strconv.ParseFloat(DX,64)
	dy,_:= strconv.ParseFloat(DY,64)

	k1:=(ay-by)/(ax-bx)
	k2:=(cy-dy)/(cx-dx)

	if k1 == k2{
		fmt.Println("平行")
	}else {
		fmt.Println("不平行")
	}

}