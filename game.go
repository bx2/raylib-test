package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	Grid *Grid
}

func NewGame(gridSize int) *Game {
	grid := NewGrid(gridSize)
	return &Game{
		Grid: grid,
	}
}

func (g *Game) Update() {
}

func (g *Game) Draw() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)
	rl.DrawFPS(5, 5)
	g.Grid.Draw()
	rl.EndDrawing()
}
