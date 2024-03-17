package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	const (
		screenWidth  = 800
		screenHeight = 720
	)
	rl.InitWindow(screenWidth, screenHeight, "[gridwars]")
	rl.SetTargetFPS(60)

	game := NewGame(5)

	for !rl.WindowShouldClose() {
		game.Update()
		game.Draw()
	}

	rl.CloseWindow()
}
