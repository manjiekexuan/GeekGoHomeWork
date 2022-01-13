package main

import (
	"fmt"
	"sort"
	"time"
)

func main() {
	elevator := &elevator{}
	elevator.inputRecord(3, 5)
	fmt.Println("当前电梯停靠的楼层为:", elevator.nowFloor)
	storey2 := &storey{}
	storey1 := &storey{}
	storey4 := &storey{}
	storey5 := &storey{}
	storey6 := &storey{}
	storey7 := &storey{}
	storey8 := &storey{}
	storey4.press(4)
	storey5.press(5)
	storey2.press(2)
	storey1.press(1)
	storey6.press(6)
	storey7.press(7)
	storey8.press(8)
	//注意，这里第一次输入的楼层需要输入两次，第一个是用来确定电梯运行的方向，第二个放进切片中进行排序
	elevator.goRun(storey8, storey8, storey4, storey5, storey2, storey7, storey1, storey6)
	fmt.Println("当前电梯停靠的楼层为:", elevator.nowFloor)
}

type elevator struct {
	defaultFloor int
	allFlor      int
	nowFloor     int
}

type storey struct {
	number int
}

func (s *storey) press(num int) {
	s.number = num

}
func (e *elevator) goRun(s1 *storey, s2 ...*storey) {
	//e.nowFloor = s1.number
	var sub int
	var direction int //(定义电梯运行的方向 1为上 0为下)
	var up []int
	var down []int
	//判断电梯上还是下
	if s1.number > e.nowFloor {
		//sub = s1.number - e.nowFloor
		//time.Sleep(time.Duration(sub) * time.Second) //电梯运行一层一秒
		//e.nowFloor = s1.number
		direction = 1
	} else {
		//sub = e.nowFloor - s1.number
		//time.Sleep(time.Duration(sub) * time.Second)
		//e.nowFloor = s1.number
		direction = 0
	}

	//如果一开始为上楼
	if direction == 1 {
		for _, item := range s2 {
			if item.number > e.nowFloor {
				up = append(up, item.number)
				sort.Ints(up)
				fmt.Println("up中的元素:", up)
			} else {
				down = append(down, item.number)
				sort.Sort(sort.Reverse(sort.IntSlice(down)))
				fmt.Println("down中的元素:", down)
			}

		}
		for _, value := range up {
			sub = value - e.nowFloor
			time.Sleep(time.Duration(sub) * time.Second)
			e.nowFloor = value
			fmt.Println("Hey,此时电梯的楼层在:", e.nowFloor)
		}

		//下楼
		for _, value := range down {
			sub = e.nowFloor - value
			time.Sleep(time.Duration(sub) * time.Second)
			e.nowFloor = value
			fmt.Println("Hey,此时电梯的楼层在:", e.nowFloor)
		}
	}

}

func (e *elevator) inputRecord(defaultFloor int, allFloor int) {
	e.defaultFloor = defaultFloor
	e.allFlor = allFloor
	e.nowFloor = e.defaultFloor
}

func (e *elevator) getElevatorFloor() interface{} { //case1 默认情况下
	//e.nowFloor = e.defaultFloor

	return e.nowFloor
}
