package main

type Point struct {
	x int
	y int
}

//NO idea how this works

func main() {

	maze := []string{
		"xxxxxxxxxx x",
		"x        x x",
		"x        x x",
		"x xxxxxxxx x",
		"x          x",
		"x xxxxxxxxxx",
	}

	start := Point{x: 10, y: 0}
	end := Point{x: 10, y: 0}

	mazeSolver(maze, "x", start, end)
}

func mazeSolver(maze []string, wall string, start Point, end Point) Point {
	var seen [][]bool
	var path []Point
	for i := 0; i < len(maze); i++ {

	}
}

var dir = [][2]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

func walk(maze []string, wall string, curr Point, end Point, seen [][]bool, path []Point) bool {
	//1. Base case
	//Off the map
	if curr.x < 0 || curr.x >= len(maze[0]) || curr.y < 0 || curr.y >= len(maze) {
		return false
	}
	// on a wall
	if string(maze[curr.y][curr.x]) == wall {
		return false
	}

	if curr.x == end.x && curr.y == end.y {
		path = append(path, end)
		return true

	}

	if seen[curr.y][curr.x] {
		return false
	}

	// Step in recursion - 3
	// pre
	// recurse
	// post

	seen[curr.y][curr.x] = true
	path = append(path, curr) //pre

	for i := 0; i < len(dir); i++ {

		x, y := dir[i][0], dir[i][1]
		if walk(maze, wall, Point{x: curr.x + x, y: curr.y + y}, end, seen, path) {
			return true
		}
	}

	_, path = pop(path) //post

	return false

}

// pop function to remove the last element from a slice of Point structs
func pop(slice []Point) (Point, []Point) {
	if len(slice) == 0 {
		// Return a zero value Point and the original slice if it's empty
		return Point{}, slice
	}
	// Get the last element
	lastElement := slice[len(slice)-1]
	// Return the last element and a new slice without the last element
	return lastElement, slice[:len(slice)-1]
}
