package main

import (
	"container/heap"
	"fmt"
	"math"
)

// PriorityQueue implementation for Dijkstra
type Item struct {
	node     int
	priority int
	index    int
}

// PriorityQueue implements heap.Interface
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(item *Item, node, priority int) {
	item.node = node
	item.priority = priority
	heap.Fix(pq, item.index)
}

// Graph represents a weighted graph using adjacency list
type Graph struct {
	vertices map[int]map[int]int
}

// NewGraph initializes a new graph
func NewGraph() *Graph {
	return &Graph{vertices: make(map[int]map[int]int)}
}

// AddEdge adds an edge between two vertices with a weight
func (g *Graph) AddEdge(src, dest, weight int) {
	if g.vertices[src] == nil {
		g.vertices[src] = make(map[int]int)
	}
	g.vertices[src][dest] = weight
	if g.vertices[dest] == nil {
		g.vertices[dest] = make(map[int]int)
	}
	g.vertices[dest][src] = weight // For undirected graph
}

// Dijkstra finds the shortest path from a source node
func (g *Graph) Dijkstra(source int) map[int]int {
	// Initialize distances with infinity
	distances := make(map[int]int)
	for v := range g.vertices {
		distances[v] = math.MaxInt64
	}
	distances[source] = 0

	// Priority queue
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Item{node: source, priority: 0})

	visited := make(map[int]bool)

	for pq.Len() > 0 {
		// Get the node with the smallest distance
		item := heap.Pop(&pq).(*Item)
		u := item.node

		// If already visited, skip
		if visited[u] {
			continue
		}
		visited[u] = true

		// Update distances to neighboring vertices
		for neighbor, weight := range g.vertices[u] {
			if !visited[neighbor] {
				newDist := distances[u] + weight
				if newDist < distances[neighbor] {
					distances[neighbor] = newDist
					heap.Push(&pq, &Item{node: neighbor, priority: newDist})
				}
			}
		}
	}

	return distances
}

/*
Core Objective of Dijkstra's Algorithm is to find the shortest path between any two vertices, therby
it finds the distance between the source and the destination


- we initially assume the distance with the higher value such as infinity 
- and updating it with the minimum distance we get as we traverse
*/ 

/*
Bellman-Ford shortest path algorithm 

- it can handle negative weight edges unlike Dijkstra's algorithm 
- Works with both directed and undirected graph 
- Can handle negative weight edges 
- Detech the presence of negative weigth cycles ( a cycle whose total weight is negative)
- Time complexity is higher compared to Dijkstra, but its ability to handle negative weights makes it crucial in many cases
*/

func (g *Graph) bellmanford(source int) (map[int]int, bool){
    distance := make(map[int]int)

    for v := range g.vertices {
      distance[v] = math.MaxInt64
    }   
    distance[source] = 0
    for i := 1; i <= len(g.vertices)-1; i++ {
      for u, neighbor := range g.vertices {
          for neighbor, weight := range neighbor {
            if distance[u] != math.MaxInt64 && distance[u]+weight < distance[neighbor] {
              distance[neighbor] = distance[u]+weight
            }
          }
      }
    }
    
    //negative cycle detection
    for u, neighbors := range g.vertices {
      for neighbor, weight := range neighbors {
        if distance[u] != math.MaxInt64 && distance[u]+weight < distance[neighbor]{
          return nil, true
        }
      }
    }

    return distance, false
}

// main function for demonstration
func main() {
	graph := NewGraph()
	graph.AddEdge(0, 1, 5)
	graph.AddEdge(0, 2, -2)
	graph.AddEdge(1, 3, 3)
	graph.AddEdge(3, 2, 1)

  //:q	fmt.Println("Shortest path distances from node 0:")
	distances, hasNegativeCycle := graph.bellmanford(0)
	if hasNegativeCycle{
    fmt.Println("Graph has negative cycle")
    return 
  } 
  for node, distance := range distances {
		fmt.Printf("Distance to node %d: %d\n", node, distance)
	}
}

