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

func fire(graph [][]Vertex, in [][]bool, monuments int) (out [][]bool) {
	out = make([][]bool, monuments)
	for i := range out {
		out[i] = make([]bool, len(in[0]))
	}
	for _, input := range in {
		println(input)
		mapa := make([][][][2]int, len(graph))
		for i := range mapa {
			mapa[i] = make([][][2]int, len(graph[0]))
			for j := range mapa[i] {
				mapa[i][j] = make([][2]int, 0) // 0: kdy; 1: kam
			}
		}
		var AQueue, FQueue []int
		var canFire map[int]bool
		AQueue = make([]int, 0)
		FQueue = make([]int, 0)
		canFire = make(map[int]bool)
		for i, v := range graph[len(graph)-1] {
			if v.cross == '^' {
				AQueue = append(AQueue, len(graph)-1*len(graph[0])+i)
			}
		}
		for len(AQueue) > 0 {
			for _, v := range AQueue {
				y := v / len(graph)
				x := v % len(graph)
				s := graph[y][x]
				c := s.cross
				switch c {
				case '.':
					for _, to := range s.next {
						if to == -1 {
							continue
						}
					}
				case '=':
				default:
					continue
				}
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
		var input [][]bool
		input = make([][]bool, I)
		for i := 0; i < I; i++ {
			input[i] = make([]bool, N)
			for j := 0; j < N; j++ {
				var tmp int
				fmt.Fscanf(scanner, " %d ", &tmp)
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
						storeGraph[y/2][(x+1)/2].next[WEST] = y/4*toOdd(H) + (x-1)/2 // H též potřeba / 2
						storeGraph[y/2][(x-1)/2].next[EAST] = y/4*toOdd(H) + (x+1)/2
					case '|':
						storeGraph[(y+1)/2][x/2].next[NORD] = (y-1)*toOdd(H)/4 + x/2
					case '/':
						storeGraph[(y+1)/2][(x-1)/2].next[NEAST] = (y-1)/4*toOdd(H) + (x+1)/2
					case '\\':
						storeGraph[(y+1)/2][(x+1)/2].next[NWEST] = (y-1)/4*toOdd(H) + (x-1)/2
					}
				}
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
