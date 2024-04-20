package main

import (
	rl "github.com/mohsengreen1388/raylib-go-custom/raylib"
)

func init() {
	rl.SetCallbackFunc(main)
}

func main() {

	rl.SetConfigFlags(rl.FlagVsyncHint)
	rl.SetConfigFlags(rl.FlagFullscreenMode)

	rl.InitWindow(int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()), "Conway's Game of Life")
	rl.SetTargetFPS(30)

	cam := rl.Camera3D{}
	cam.Fovy = 45
	cam.Position = rl.Vector3{2, 2, 6}
	cam.Projection = rl.CameraPerspective
	cam.Up = rl.Vector3{0, 1, 0}
	cam.Target = rl.Vector3{0, 0, 0}

	tx := rl.LoadTexture("assets/raylib_logo.png")

	for !rl.WindowShouldClose() {

		rl.BeginDrawing()
		rl.ClearBackground(rl.Gray)
		rl.DrawTextureEx(tx, rl.Vector2{float32(rl.GetScreenWidth()) / 2.5, float32(rl.GetScreenHeight()) / 2.5}, 0, 1, rl.White)
		rl.DrawFPS(10, 10)
		rl.EndDrawing()

	}

	rl.CloseWindow()
}
