package graphs

import "fmt"

type Graph struct {
	AdjacencyList map[string][]string
}

func (g *Graph) Init() {
	g.AdjacencyList = make(map[string][]string)
}

func (g *Graph) AddVertex(vertex string) bool {
	if g.AdjacencyList[vertex] == nil {
		g.AdjacencyList[vertex] = []string{}
		fmt.Println("Added vertex", g.AdjacencyList[vertex])
		return true
	}

	return false
}

func (g *Graph) AddEdge(vertex1 string, vertex2 string) bool {
	if g.AdjacencyList[vertex1] != nil && g.AdjacencyList[vertex2] != nil {
		g.AdjacencyList[vertex1] = append(g.AdjacencyList[vertex1], vertex2)
		g.AdjacencyList[vertex2] = append(g.AdjacencyList[vertex2], vertex1)

		return true
	}

	return false
}

func (g *Graph) Print() {
	fmt.Println(g.AdjacencyList)
}

func Play() {
	g := Graph{}

	g.Init()
	fmt.Println(g.AddVertex("1"))
	fmt.Println(g.AddVertex("2"))

	fmt.Println(g.AddEdge("1", "2"))

	g.Print()
}
