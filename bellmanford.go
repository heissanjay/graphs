package main 


import (
  "fmt"
  "math"
)

type Graph struct {
  vertices map[int]map[int]int
}

func NewGraph() *Graph {
  return &Graph{vertices: make(map[int]map[int]int)}
}

func (g *Graph) AddEdge(src, dest, weight int) {
    if g.vertices[src] == nil {
      g.vertices[src] = make(map[int]int)
    }
    g.vertices[src][dest] = weight
    // undirected graph
    if g.vertices[dest] == nil {
      g.vertices[dest] = make(map[int]int)
    }
    g.vertices[dest][src] = weight
}


func (g *Graph) bellmanford(source int) (map[int]int, bool) {
  distances := make(map[int]int)
  for vertex := range g.vertices {
    distances[vertex] = math.MaxInt64
  }
  
  distances[source] = 0
  // do the relaxation for len(vertices)-1 times
  for i := 1; i <= len(g.vertices)-1; i++ {
      for u, neighbors := range g.vertices {
          for neighbor, weight := range neighbors {
              if distances[u] != math.MaxInt64 && distances[u]+weight < distances[neighbor]{
                    distances[neighbor] = distances[u]+weight
              }
          }
      }
  }

  // negative cycle detection, try to do the relaxation one more time over len(vertices)-1 times
  for u, neighbors := range g.vertices {
    for neighbor, weight := range neighbors {
        if distances[u] != math.MaxInt64 && distances[u]+weight < distances[neighbor]{
          // cycle present
          return nil, true
        }
    }
  }
  return distances, false
}

func main(){
    g := Graph{
    vertices:map[int]map[int]int{
      0: {1: 5, 2: -2},
      1: {3: 3},
      2: {1: 2},
      3: {2: 1},
      },
    }

    dist, hasCycle := g.bellmanford(0)
    if hasCycle {
      fmt.Println("graph has cycle")
      return 
    }
    for u := range g.vertices {
      fmt.Println("distance from the vertex 0 to the vertex ->", u, dist[u])
    }
}
