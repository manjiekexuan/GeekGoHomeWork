package main

import "testing"

type Rank interface {
	UpdateFR(name string, fr float64)
	GetRank(name string) int
}

type Client interface {
	UpdateFR(name string, fr float64)
	GetRank(name string) int
}

func TestHomework(t *testing.T) {
	//var rank Rank
	var clients []Client
	for i := 0; i < len(clients); i++ {
		go func(idx int) {
			//TODO add context to control exit

			go func() {
				for {
					go clients[idx].UpdateFR("戴一杰", 0.23)
				}
			}()

			go func() {
				for {
					go clients[idx].GetRank("戴一杰")
				}
			}()

		}(i)

	}
}
