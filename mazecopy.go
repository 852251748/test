package main

import (
	"fmt"
	"os"
)

var dirs = []pos{
	{1, 0, nil},
	{-1, 0, nil},
	{0, 1, nil},
	{0, -1, nil},
}

func readMazeC(filename string) ([][]int, error) {
	var maze [][]int
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Open file faile", filename, err)
		return maze, err
	}
	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)

	maze = make([][]int, row)
	for v, _ := range maze {
		maze[v] = make([]int, col)
	}
	for i := 0; i < len(maze); i++ {
		for v, _ := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][v])
		}
	}
	fmt.Println(maze)
	return maze, nil
}

type pos struct {
	x, y int
	next *pos
}

func (p pos) add(mov pos) pos {
	var temp pos
	temp.x = p.x + mov.x
	temp.y = p.y + mov.y
	temp.next = &p
	return temp
}

func (p pos) check(mapp [][]int) (int, bool) {
	if p.x < 0 || p.x >= len(mapp) {
		return 0, false
	}
	if p.y < 0 || p.y >= len(mapp[p.x]) {
		return 0, false
	}
	return mapp[p.x][p.y], true
}

func calsteps(maze [][]int, sta pos, end pos) {

	//初始化路径图
	route := make([][]int, len(maze))
	for v, _ := range route {
		route[v] = make([]int, len(maze[v]))
	}
	//起点设置1
	route[sta.x][sta.y] = 1

	queue := make([]pos, 1)
	queue[0] = sta
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if cur.x == end.x && cur.y == end.y {
			for i := 0; i < len(route); i++ {
				for _, v := range route[i] {
					fmt.Print(fmt.Sprintf("%02d ", v))
				}
				fmt.Println()
			}

		}
		for _, v := range dirs {
			temp := cur.add(v)
			val, ok := temp.check(maze)
			if !ok || val == 1 {
				continue
			}
			val, ok = temp.check(route)
			if !ok || val != 0 {
				continue
			}
			queue = append(queue, temp)
			route[temp.x][temp.y] = route[cur.x][cur.y] + 1
		}
	}
}

func main() {
	maze, err := readMazeC("Maze.txt")
	if err != nil {
		fmt.Println("Open file faile", err)
		return
	}
	sta := pos{0, 0, nil}
	end := pos{4, 0, nil}
	calsteps(maze, sta, end)
}
