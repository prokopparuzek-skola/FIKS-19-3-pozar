package main

import (
	"bufio"
	"fmt"
	"os"
)

func (v Vertex) String() string {
	return fmt.Sprintf("%c", v.cross)
}

type Vertex struct {
	cross rune
	next  [5]int
}

func main() {
	scanner := bufio.NewReader(os.Stdin)
	var T int
	fmt.Fscanf(scanner, "%d\n", &T)
	for ; T > 0; T-- {
		var W, H, I, N, V int
		fmt.Fscanf(scanner, "%d%d%d%d%d\n", &W, &H, &I, &N, &V)
		println(W, H, I, N, V)
		H -= 2 // prázdné řádky můžu přeskočit
		W -= 2 // prázdné sloupce můžu přeskočit
		// inicializace grafu
		var graph [][]Vertex
		graph = make([][]Vertex, H)
		for i := range graph {
			graph[i] = make([]Vertex, W)
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
				graph[y][x].cross = c
			}
			scanner.ReadBytes('\n')
		}
		// přeskočím 2 řádky, okraj a mezery
		scanner.ReadBytes('\n')
		scanner.ReadBytes('\n')
		// scan vstupu
		for ; I > 0; I-- {
			scanner.ReadBytes('\n')
		}
		for y := 0; y < H; y++ {
			for x := 0; x < W; x++ {
				fmt.Print(graph[y][x])
			}
			fmt.Println()
		}
	}
}
