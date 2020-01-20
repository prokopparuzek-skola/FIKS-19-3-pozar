package main

import "fmt"

func (v Vertex) String() string {
	return string(v.cross)
}

type Vertex struct {
	cross rune
	next  [5]int
}

func main() {
	var T int
	fmt.Scanf("%d", &T)
	for ; T > 0; T-- {
		var W, H, I, N, V int
		fmt.Scanf("%d%d%d%d%d", &W, &H, &I, &I, &N, &V)
		H -= 2 // prázdné řádky můžu přeskočit
		W -= 2 // prázdné sloupce můžu přeskočit
		println(H, W)
		// inicializace grafu
		var graph [][]Vertex
		graph = make([][]Vertex, H)
		for i := range graph {
			graph[i] = make([]Vertex, W)
		}
		// přeskočím 2 řádky, okraj a mezery
		fmt.Scanln()
		fmt.Scanln()
		for y := 0; y < H; y++ {
			fmt.Scanf("| ")
			for x := 0; x < W; x++ {
				var c rune
				fmt.Scan(&c)
				graph[y][x].cross = c
			}
			fmt.Scanf(" |")
		}
		for ; I > 0; I-- {
			fmt.Scanln()
		}
		// přeskočím 2 řádky, okraj a mezery
		fmt.Scanln()
		fmt.Scanln()
		for y := 0; y < H; y++ {
			for x := 0; x < W; x++ {
				fmt.Print(graph[y][x])
			}
			println()
		}
		println()
	}
}
