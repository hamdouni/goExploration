package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {

	height := int32(rl.GetScreenHeight())
	width := int32(rl.GetScreenWidth())
	rl.InitWindow(width, height, "raylib [core] example - basic window")

	rl.SetTargetFPS(60)

	var posY int32 = 200

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.DrawText("Congrats! You created your first window!", 190, posY, 20, rl.Black)

		rl.EndDrawing()

		posY--
		if posY < 0 {
			posY = height - 100
		}
	}

	rl.CloseWindow()
}
