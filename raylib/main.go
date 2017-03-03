package main

import "github.com/gen2brain/raylib-go/raylib"

func main() {
	raylib.InitWindow(400, 225, "raylib [core] example - basic window")

	raylib.SetTargetFPS(60)

	for !raylib.WindowShouldClose() {
		raylib.BeginDrawing()

		raylib.ClearBackground(raylib.Black)

		raylib.DrawText("My first window!", 95, 100, 20, raylib.LightGray)

		raylib.EndDrawing()
	}

	raylib.CloseWindow()
}
