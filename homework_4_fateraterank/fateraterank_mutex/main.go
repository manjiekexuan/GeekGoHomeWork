package main

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"
)

type People struct {
	Name    string
	FatRate float64
}

var person_fatRate = map[string]float64{}
var lock sync.Mutex

func main() {
	people_num := 1000
	peoples := []People{}
	rand.Seed(time.Now().Unix())

	//生成1000个人
	for i := 0; i < people_num; i++ {
		peoples = append(peoples, People{
			Name:    fmt.Sprintf("Stu%d", i),
			FatRate: floats_rank(0.1, 0.4, 1),
		})
	}

	//将1000个人注册进  person_fatRate = map[string]float64{}
	for _, v := range peoples {
		v.register_person(v.Name, v.FatRate)
	}

	//完成1000人不停的插入，并返回对应的排名
	for {
		for i := 0; i < people_num; i++ {
			wg := sync.WaitGroup{}
			wg.Add(1)
			peoples[i].FatRate = floats_rank(0.1, 0.4, 1)
			go func(name string, fatRate float64, wg *sync.WaitGroup) {
				defer wg.Done()
				update_fatrate(peoples[i].Name, peoples[i].FatRate)
			}(peoples[i].Name, peoples[i].FatRate, &wg)
			wg.Wait()
			//获取排名
			rank, _ := createRand(peoples[i].Name)
			fmt.Printf("%s的排名是%d\n", peoples[i].Name, rank)
		}
	}

}

//更新插入姓名体脂率
func update_fatrate(name string, rate float64) {
	lock.Lock()
	person_fatRate[name] = rate
	lock.Unlock()
}

//注册姓名体脂率
func (p People) register_person(name string, fatRate float64) {
	lock.Lock()
	person_fatRate[name] = fatRate
	lock.Unlock()
}

//随机生成体脂率
func floats_rank(min, max float64, n int) float64 {
	res := make([]float64, n)
	for i := range res {
		res[i] = min + rand.Float64()*(max-min)
	}
	result := res[0]
	return result
}

//根据姓名获取排名
func createRand(name string) (rank int, fataRate float64) {
	fatRate2PersonMap := map[float64][]string{}
	rankArr := make([]float64, 0, len(person_fatRate))
	for nameItem, frItem := range person_fatRate {
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
