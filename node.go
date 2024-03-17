package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Node struct {
	SecLevel uint8
	Breached bool
}

func (n *Node) Draw(x float32, y float32) {
	pos := rl.Vector2{X: x, Y: y}
	radius := float32(30)
	var color rl.Color

	switch {
	case n.IsMouseOver(x, y, radius):
		color = rl.Yellow
	case n.Breached:
		color = rl.Green
	default:
		color = rl.Red
	}

	rl.DrawCircleV(pos, radius, color)
}

func (n *Node) IsMouseOver(x, y, radius float32) bool {
	mouseX := rl.GetMouseX()
	mouseY := rl.GetMouseY()
	distance := math.Sqrt(float64((x-float32(mouseX))*(x-float32(mouseX)) + (y-float32(mouseY))*(y-float32(mouseY))))
	return distance < float64(radius)
}

func (n *Node) HandleMouseClick(x, y, radius float32) {
	if n.IsMouseOver(x, y, radius) && rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		n.Breached = true
	}
}
