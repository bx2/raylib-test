package main

type Grid struct {
	Nodes []Node
}

func NewGrid(gridSize int) *Grid {
	return &Grid{
		Nodes: make([]Node, gridSize*gridSize),
	}
}

func (g *Grid) Draw() {
	const startX, startY = 100, 100
	const spacing = 100
	const nodesPerRow = 5

	for i, node := range g.Nodes {
		x := startX + float32(i%nodesPerRow)*spacing
		y := startY + float32(i/nodesPerRow)*spacing

		g.Nodes[i].HandleMouseClick(x, y, 30)
		node.Draw(x, y)
	}
}
