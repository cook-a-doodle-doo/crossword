package main

import (
	"fmt"
	"reflect"
	"time"
)

type bord [9][9][]byte

func (b bord) Init() bord {
	for i, y := range b {
		for j, x := range y {
			if x[0] == 0 {
				b[i][j] = []byte{1, 2, 3, 4, 5, 6, 7, 8, 9}
			}
		}
	}
	return b
}

func (b bord) Show() {
	for i, y := range b {
		if i%3 == 0 {
			fmt.Println("-------------------------------")
		}
		for j, x := range y {
			if j%3 == 0 {
				fmt.Printf("|")
			}
			if len(x) == 1 {
				fmt.Printf(" %d ", x[0])
			} else {
				fmt.Printf("   ")
			}
		}
		fmt.Printf("|\n")
	}
	fmt.Println("-------------------------------")
}

func remove(bytes []byte, search byte) []byte {
	result := []byte{}
	for _, v := range bytes {
		if v != search {
			result = append(result, v)
		}
	}
	return result
}

func (b bord) haveNum(x, y int, num byte) bool {
	cx := x / 3
	cy := y / 3
	for i := cy * 3; i < (cy+1)*3; i++ {
		for j := cx * 3; j < (cx+1)*3; j++ {
			if x == j && y == i {
				continue
			}
			if len(b[i][j]) != 1 {
				continue
			}
			if b[i][j][0] == num {
				return true
			}
		}
	}
	//about row
	for i := 0; i < 9; i++ {
		if i == x {
			continue
		}
		if len(b[y][i]) != 1 {
			continue
		}
		if b[y][i][0] == num {
			return true
		}
	}
	//about col
	for i := 0; i < 9; i++ {
		if i == y {
			continue
		}
		if len(b[i][x]) != 1 {
			continue
		}
		if b[i][x][0] == num {
			return true
		}
	}
	return false
}

func (b bord) Update() (bord, bool) {
	end := true
	for i, y := range b {
		for j, x := range y {
			if len(x) == 1 {
				continue
			}
			result := []byte{}
			for _, v := range x {
				if b.haveNum(j, i, v) {
					continue
				}
				result = append(result, v)
			}
			if !reflect.DeepEqual(b[i][j], result) {
				end = false
			}
			b[i][j] = result
		}
	}
	return b, end
}

func main() {
	bord := bord{
		{{0}, {7}, {0}, {0}, {1}, {0}, {0}, {5}, {0}},
		{{2}, {0}, {0}, {0}, {7}, {0}, {0}, {0}, {4}},
		{{0}, {0}, {1}, {3}, {0}, {9}, {2}, {0}, {0}},
		{{0}, {0}, {3}, {9}, {0}, {2}, {6}, {0}, {0}},
		{{6}, {2}, {0}, {0}, {0}, {0}, {0}, {1}, {9}},
		{{0}, {0}, {4}, {1}, {0}, {5}, {8}, {0}, {0}},
		{{0}, {0}, {7}, {4}, {0}, {8}, {3}, {0}, {0}},
		{{3}, {0}, {0}, {0}, {2}, {0}, {0}, {0}, {6}},
		{{0}, {1}, {0}, {0}, {5}, {0}, {0}, {4}, {0}}}
	bord = bord.Init()
	for {
		fmt.Println(bord[0])
		fmt.Println(bord[1])
		fmt.Println(bord[2])
		fmt.Println(bord[3])
		fmt.Println(bord[4])
		fmt.Println(bord[5])
		fmt.Println(bord[6])
		fmt.Println(bord[7])
		fmt.Println(bord[8])
		bord.Show()
		b, end := bord.Update()
		if end {
			break
		}
		bord = b
		time.Sleep(time.Second)
	}
	fmt.Println("読みは未実装 やれば出来るけどめんどい Updateに追加")
}
