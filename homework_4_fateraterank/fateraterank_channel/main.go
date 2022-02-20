package main

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"
)

type People struct {
	name    string
	rateFat float64
	rank    int
}

func randFlots(min, max float64, n int) float64 {
	res := make([]float64, n)
	for i := range res {
		res[i] = min + rand.Float64()*(max-min)
	}
	result := res[0]
	return result
}

func createRand(name string, rank chan []People) {
	slicePeople := <-rank
	for _, people := range slicePeople {
		if people.name == name {
			fmt.Printf("%s的体脂排名为%d\n", name, people.rank)
		}
	}
}

func main() {
	for {
		ChaneelPepole := make(chan People, 1000)
		ChRank := make(chan []People, 1000)
		slicePeople := []People{}
		counterPeople := 1000
		wg := sync.WaitGroup{}
		wg.Add(counterPeople)
		rand.Seed(time.Now().Unix())

		for i := 0; i < counterPeople; i++ {
			go func(i int, wg *sync.WaitGroup) {
				defer wg.Done()
				var people = People{
					name:    fmt.Sprintf("Stu%d", i),
					rateFat: randFlots(0.1, 0.4, 1),
				}

				ChaneelPepole <- people

				createRand(people.name, ChRank)
			}(i, &wg)
		}

		finishedFileCount := 0
		for people := range ChaneelPepole {
			finishedFileCount++
			slicePeople = append(slicePeople, people)
			if finishedFileCount == counterPeople {
				close(ChaneelPepole)
			}
		}

		sort.Slice(slicePeople, func(i, j int) bool {
			return slicePeople[i].rateFat < slicePeople[j].rateFat
		})

		for i, _ := range slicePeople {
			slicePeople[i].rank = i + 1
		}

		for i := 0; i < counterPeople; i++ {
			ChRank <- slicePeople
		}

		wg.Wait()
	}
}
