package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	WEST = iota
	EAST
	NORD
	NWEST
	NEAST
)

func toOdd(x int) int {
	if x%2 == 0 {
		return x
	} else {
		return x + 1
	}
}

func (v Vertex) String() string {
	var s string
	s = fmt.Sprintf("%c:%v", v.cross, v.next)
	return s
}

type Vertex struct {
	cross rune
	next  [5]int
}

func contain(mapa map[int]int, v int) bool {
	for test := range mapa {
		if test == v {
			return true
		}
	}
	return false
}

func line(Px, Py, x, y int) (Fx, Fy int) {
	// řeší X
	if Px < x {
		Fx = x + 1
	} else if Px > x {
		Fx = x - 1
	} else {
		Fx = x
	}
	// řeší Y
	if Py < y {
		Fy = y + 1
	} else if Py > y {
		Fy = y - 1
	} else {
		Fy = y
	}
	return
}

func fire(graph [][]Vertex, in [][]bool, monuments int) (out [][]bool) {
	out = make([][]bool, monuments)
	for i := range out {
		out[i] = make([]bool, len(in[0]))
	}
	for Istep, input := range in {
		mapa := make([][]map[int]int, len(graph))
		for i := range mapa {
			mapa[i] = make([]map[int]int, len(graph[0]))
			for j := range mapa[i] {
				mapa[i][j] = make(map[int]int) // if exist return step; odkud
			}
		}
		var AQueue, FQueue []int
		AQueue = make([]int, 0)
		FQueue = make([]int, 0)
		step := 0
		for i, v := range graph[len(graph)-1] {
			if v.cross == '^' {
				if input[step] {
					AQueue = append(AQueue, (len(graph)-1)*len(graph[0])+i)
				}
				step++
			}
		}
		step = 0 // počítadlo kroků
		for len(AQueue) > 0 {
			//fmt.Println(AQueue)
			for _, v := range AQueue {
				y := v / len(graph)
				x := v % len(graph)
				s := graph[y][x]
				c := s.cross
				switch c {
				case '^':
					fallthrough
				case '.':
					for _, to := range s.next {
						Fy := to / len(graph)
						Fx := to % len(graph)
						if to == -1 {
							continue
						}
						_, err := mapa[Fy][Fx][v]
						if err || contain(mapa[y][x], to) {
							continue
						}
						//fmt.Print(to, ": ")
						//fmt.Println(Fy, Fx)
						FQueue = append(FQueue, to)
						mapa[Fy][Fx][v] = step
					}
				case '=':
					for Pv, Ps := range mapa[y][x] {
						if Ps == (step - 1) {
							Py := Pv / len(graph)
							Px := Pv % len(graph)
							for _, to := range s.next {
								Fy := to / len(graph)
								Fx := to % len(graph)
								PredictX, PredictY := line(Px, Py, x, y)
								if to == -1 {
									continue
								} else if PredictX != Fx || PredictY != Fy {
									continue
								}
								_, err := mapa[Fy][Fx][v]
								if err || contain(mapa[y][x], to) {
									continue
								}
								//fmt.Print(to, ": ")
								//fmt.Println(Fy, Fx)
								FQueue = append(FQueue, to)
								mapa[Fy][Fx][v] = step
							}
						}
					}
				default:
					continue
				}
			}
			step++
			AQueue = FQueue
			FQueue = make([]int, 0)
		}
		step = 0
		//fmt.Println(mapa[0])
		for i, o := range graph[0] {
			if o.cross == '?' {
				if len(mapa[0][i]) > 0 {
					out[Istep][step] = true
				}
				step++
			}
		}
	}
	return
}

func main() {
	scanner := bufio.NewReader(os.Stdin)
	var T int
	fmt.Fscanf(scanner, "%d\n", &T)
	for ; T > 0; T-- {
		var W, H, I, N, V int
		fmt.Fscanf(scanner, "%d%d%d%d%d\n", &W, &H, &I, &N, &V)
		H -= 2 // prázdné řádky můžu přeskočit
		W -= 2 // prázdné sloupce můžu přeskočit
		// inicializace grafu
		var graph [][]rune
		graph = make([][]rune, H)
		for i := range graph {
			graph[i] = make([]rune, W)
		}
		// přeskočím 2 řádky, okraj a mezery
		scanner.ReadBytes('\n')
		scanner.ReadBytes('\n')
		for y := 0; y < H; y++ {
			scanner.ReadBytes(' ')
			for x := 0; x < W; x++ {
				var c rune
				tmp, _ := scanner.ReadByte()
				c = rune(tmp)
				graph[y][x] = c
			}
			scanner.ReadBytes('\n')
		}
		// přeskočím 2 řádky, okraj a mezery
		scanner.ReadBytes('\n')
		scanner.ReadBytes('\n')
		// scan vstupu
		var tmp int
		var input [][]bool
		input = make([][]bool, I)
		for i := 0; i < I; i++ {
			input[i] = make([]bool, N)
			for j := 0; j < N; j++ {
				fmt.Fscan(scanner, &tmp)
				if tmp == 1 {
					input[i][j] = true
				} else {
					input[i][j] = false
				}
			}
		}
		var storeGraph [][]Vertex
		storeGraph = make([][]Vertex, toOdd(H)/2)
		for i := range storeGraph {
			storeGraph[i] = make([]Vertex, toOdd(W)/2)
			for j := range storeGraph[i] {
				storeGraph[i][j].next = [5]int{-1, -1, -1, -1, -1}
			}
		}
		for y := H - 1; y >= 0; y-- {
			for x := 0; x < W; x++ {
				if x%2 == 0 && y%2 == 0 {
					storeGraph[y/2][x/2].cross = graph[y][x]
				} else {
					switch graph[y][x] {
					case '-':
						storeGraph[y/2][(x+1)/2].next[WEST] = y*toOdd(H)/4 + (x-1)/2 // H též potřeba / 2
						storeGraph[y/2][(x-1)/2].next[EAST] = y*toOdd(H)/4 + (x+1)/2
					case '|':
						storeGraph[(y+1)/2][x/2].next[NORD] = (y-1)*toOdd(H)/4 + x/2
					case '/':
						storeGraph[(y+1)/2][(x-1)/2].next[NEAST] = (y-1)*toOdd(H)/4 + (x+1)/2
					case '\\':
						storeGraph[(y+1)/2][(x+1)/2].next[NWEST] = (y-1)*toOdd(H)/4 + (x-1)/2
					}
				}
			}
		}
		out := fire(storeGraph, input, I)
		for _, ca := range out {
			for _, f := range ca[:len(ca)-1] {
				if f {
					fmt.Print("1 ")
				} else {
					fmt.Print("0 ")
				}
			}
			if ca[len(ca)-1] {
				fmt.Println("1")
			} else {
				fmt.Println("0")
			}
		}
		/*
			for _, y := range storeGraph {
				for _, v := range y {
					fmt.Print(v)
				}
				fmt.Println()
			}
		*/
	}
}
