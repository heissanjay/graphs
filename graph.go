package main 


import (
  "fmt"
  "container/list"
)

/*
First we will define the Graph type
*/

type Graph struct {
  vertices map[int][]int
}

func NewGraph() *Graph{
  return &Graph{vertices:make(map[int][]int)}
}

func (g *Graph) addEdge(src, dest int){
  g.vertices[src] = append(g.vertices[src], dest)
  g.vertices[dest] = append(g.vertices[dest], src)
}

// Depth first search

func (g *Graph) DFS(start int, visited map[int]bool){
  visited[start] = true
  fmt.Println(start)

  for _, neighbor := range g.vertices[start]{
    if !visited[neighbor] {
      g.DFS(neighbor, visited)
    }
  }
}

func (g *Graph) BFS(start int){
  visited := make(map[int]bool)
  queue := list.New()
  
  visited[start] = true
  queue.PushBack(start)


  for queue.Len() > 0 {
    vertex := queue.Front()
    queue.Remove(vertex)

    fmt.Println(vertex.Value)
    
    for _, neighbor := range g.vertices[vertex.Value.(int)] {
      if !visited[neighbor]{
        visited[neighbor] = true
        queue.PushBack(neighbor)
      }
    }

  }

}


func main(){
   graph := NewGraph()
   graph.addEdge(0, 1)
   graph.addEdge(0, 2)
   graph.addEdge(1, 3)
   graph.addEdge(1, 4)
   graph.addEdge(2, 5)
   graph.addEdge(5, 4)
   graph.addEdge(4, 3)


  visited := make(map[int]bool)
  graph.DFS(0, visited)
  fmt.Println("---")
  graph.BFS(0)

}
