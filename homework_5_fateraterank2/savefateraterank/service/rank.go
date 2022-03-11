package main

import (
	"fmt"
	"math"
	"sort"
)

type RankItem struct {
	Name    string
	FatRate float64
}
type FatRateRank struct {
	Item []RankItem
}

var personFatRate = map[string]float64{}

func (f *FatRateRank) inputRecord(name string, fatRate ...float64) {
	minFatRate := math.MaxFloat64
	for _, value := range fatRate {
		if value <= minFatRate {
			minFatRate = value
		}
	}
	for i, item := range f.Item {
		if item.Name == name {
			item.FatRate = minFatRate
		}
		f.Item[i] = item //?  啥意思
	}
	personFatRate[name] = minFatRate
}

func (f *FatRateRank) getRand(name string) (rank int, fataRate float64) {
	fatRate2PersonMap := map[float64][]string{}
	rankArr := make([]float64, 0, len(personFatRate))
	for nameItem, frItem := range personFatRate {
		fatRate2PersonMap[frItem] = append(fatRate2PersonMap[frItem], nameItem)
		rankArr = append(rankArr, frItem)
	}
	sort.Float64s(rankArr)
	for i, frItem := range rankArr {
		_names := fatRate2PersonMap[frItem]
		for _, _name := range _names {
			if _name == name {
				rank = i + 1
				fataRate = frItem
				return
			}

		}

	}
	return 0, 0
}

//冒泡排序
func (f *FatRateRank) getRandByBubble(name string) (rank int, fataRate float64) {
	fatRate2PersonMap := map[float64][]string{}
	rankArr := make([]float64, 0, len(personFatRate))
	for nameItem, frItem := range personFatRate {
		fatRate2PersonMap[frItem] = append(fatRate2PersonMap[frItem], nameItem)
		rankArr = append(rankArr, frItem)
	}
	bubbleSort(rankArr)
	for i, frItem := range rankArr {
		_names := fatRate2PersonMap[frItem]
		for _, _name := range _names {
			if _name == name {
				rank = i + 1
				fataRate = frItem
				return
			}

		}

	}
	return 0, 0
}

func bubbleSort(arry []float64) {
	for i := 0; i < len(arry); i++ {
		for j := 0; j < len(arry)-i-1; j++ {
			if arry[j] > arry[j+1] {
				arry[j], arry[j+1] = arry[j+1], arry[j]
			}
		}
		//fmt.Println("中间排序:", arry)
	}

	fmt.Println("最终排序:", arry)
}

//快速排序
func (f *FatRateRank) getRandByQuick(name string) (rank int, fataRate float64) {
	fatRate2PersonMap := map[float64][]string{}
	rankArr := make([]float64, 0, len(personFatRate))
	for nameItem, frItem := range personFatRate {
		fatRate2PersonMap[frItem] = append(fatRate2PersonMap[frItem], nameItem)
		rankArr = append(rankArr, frItem)
	}
	quickSort(rankArr)
	for i, frItem := range rankArr {
		_names := fatRate2PersonMap[frItem]
		for _, _name := range _names {
			if _name == name {
				rank = i + 1
				fataRate = frItem
				return
			}

		}

	}
	return 0, 0
}

func quickSort(arry []float64) {
	for i := 0; i < len(arry); i++ {
		for j := 0; j < len(arry)-i-1; j++ {
			if arry[j] > arry[j+1] {
				arry[j], arry[j+1] = arry[j+1], arry[j]
			}
		}
		//fmt.Println("中间排序:", arry)
	}

	fmt.Println("最终排序:", arry)
}
