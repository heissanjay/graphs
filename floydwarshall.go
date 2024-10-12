package main

import (
	"fmt"
	"math"
)

func floydwarshall(vertices int, edges [][]int) [][]int {
	distance := make([][]int, vertices)
	for i := 0; i < vertices; i++ {
		distance[i] = make([]int, vertices)
		for j := 0; j < vertices; j++ {
			if i == j {
				distance[i][j] = 0 // Distance to self is 0
			} else {
				distance[i][j] = math.MaxInt64 // Initialize to infinity
			}
		}
	}

	for _, edge := range edges {
		u := edge[0]
		v := edge[1]
		w := edge[2]
		distance[u][v] = w // Set edge weight
	}

	for k := 0; k < vertices; k++ {
		for i := 0; i < vertices; i++ {
			for j := 0; j < vertices; j++ {
				if distance[i][k] == math.MaxInt64 || distance[k][j] == math.MaxInt64 {
					continue // Skip if either distance is infinity
				}
				if distance[i][j] > distance[i][k]+distance[k][j] {
					distance[i][j] = distance[i][k] + distance[k][j]
				}
			}
		}
	}

	return distance
}

func main() {
	edges := [][]int{
		{0, 1, 2}, // A -> B with weight 2
		{0, 2, 1}, // A -> C with weight 1
		{1, 2, 3}, // B -> C with weight 3
	}

	vertices := 3
	shortestPaths := floydwarshall(vertices, edges)

	fmt.Println("Distance Matrix:")
	for i := 0; i < vertices; i++ {
		for j := 0; j < vertices; j++ {
			if shortestPaths[i][j] == math.MaxInt64 {
				fmt.Print("inf ")
			} else {
				fmt.Printf("%d ", shortestPaths[i][j])
			}
		}
		fmt.Println()
	}
}
