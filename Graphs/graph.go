package graphs

import "fmt"

type Graph struct {
	Vertexes map[string][]string
}

func (g *Graph) Init() {
	g.Vertexes = make(map[string][]string)
}

func (g *Graph) AddVertex(key string) bool {
	if g.Vertexes[key] == nil {
		g.Vertexes[key] = []string{}
		fmt.Println("Added vertex", g.Vertexes[key])
		return true
	}

	return false
}

func (g *Graph) AddEdge(vertex1, vertex2 string) bool {

	if g.Vertexes[vertex1] != nil && g.Vertexes[vertex2] != nil {
		g.Vertexes[vertex1] = append(g.Vertexes[vertex1], vertex2)
		g.Vertexes[vertex2] = append(g.Vertexes[vertex2], vertex1)

		return true
	}

	return false
}

func (g *Graph) RemoveEdge(vertex1, vertex2 string) bool {
	if g.Vertexes[vertex1] != nil && g.Vertexes[vertex2] != nil {
		g.Vertexes[vertex1] = remove(g.Vertexes[vertex1], vertex2)
		g.Vertexes[vertex2] = remove(g.Vertexes[vertex2], vertex1)
		return true
	}

	return false
}

func (g *Graph) RemoveVertex(vertex1 string) {
	// Find is vertex contain in the others
	for _, key := range g.Vertexes[vertex1] {
		g.RemoveEdge(key, vertex1)
	}

	delete(g.Vertexes, vertex1)
}

func remove(adjacency []string, text string) []string {
	var result []string
	for i, v := range adjacency {
		if text == v {
			result = append(result, adjacency[:i]...)
			result = append(result, adjacency[i+1:]...)
		}
	}

	return result
}

func (g *Graph) Print() {
	fmt.Println(g.Vertexes)
}

func Play() {
	g := Graph{}

	g.Init()
	g.AddVertex("1")
	g.AddVertex("2")
	g.AddVertex("3")
	g.AddVertex("4")
	g.AddVertex("5")

	g.AddEdge("1", "2")
	g.AddEdge("1", "3")
	g.AddEdge("1", "4")

	g.RemoveEdge("1", "3")

	g.RemoveVertex("4")

	g.Print()
}
