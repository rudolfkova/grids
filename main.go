package main

import (
	"fmt"
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/rudolfkova/vectozavr/camera"
	"github.com/rudolfkova/vectozavr/object"
	"github.com/rudolfkova/vectozavr/vectozavr"
)

type Game struct {
	w, h int

	scale float64
	tilt  float64
	angle float64

	P      vectozavr.Matrix
	R      vectozavr.Matrix
	S      vectozavr.Matrix
	Scale  vectozavr.Matrix
	worldE vectozavr.Matrix
	SP     vectozavr.Matrix

	oX, oY, oZ          axes
	newOX, newOY, newOZ axes

	cube Cube

	cam camera.Camera
}

func NewGame() *Game {
	g := &Game{
		scale: 1.0,
		w:     1000,
		h:     700,
	}

	g.oX.v = vectozavr.NewVec3(1, 0, 0)
	g.oX.color = color.RGBA{255, 0, 0, 255}
	g.oY.v = vectozavr.NewVec3(0, 1, 0)
	g.oY.color = color.RGBA{0, 255, 0, 255}
	g.oZ.v = vectozavr.NewVec3(0, 0, 1)
	g.oZ.color = color.RGBA{0, 0, 255, 255}

	b := 0.1
	pos := vectozavr.NewVec4(2, 0, 2, 0)
	g.cube = Cube{
		A: vectozavr.NewVec4(b, b, b, 1).Add(pos),
		B: vectozavr.NewVec4(b, b, -b, 1).Add(pos),
		C: vectozavr.NewVec4(b, -b, b, 1).Add(pos),
		D: vectozavr.NewVec4(b, -b, -b, 1).Add(pos),
		E: vectozavr.NewVec4(-b, b, b, 1).Add(pos),
		F: vectozavr.NewVec4(-b, b, -b, 1).Add(pos),
		G: vectozavr.NewVec4(-b, -b, b, 1).Add(pos),
		H: vectozavr.NewVec4(-b, -b, -b, 1).Add(pos),
	}

	g.cam.A = float64(g.w) / float64(g.h)
	g.cam.Far = 100
	g.cam.Near = 0
	g.cam.Fov = 90
	g.cam.Left = g.oX.v
	g.cam.At = g.oZ.v
	g.cam.Up = g.oY.v

	g.P = vectozavr.Projection(g.cam.Fov, g.cam.A, g.cam.Near, g.cam.Far)
	g.S = vectozavr.ScreenSpace(float64(g.w), float64(g.h))

	return g
}

func (g *Game) Update() error {
	g.R = vectozavr.RotationY(g.angle).MatMul(vectozavr.RotationX(g.tilt))
	g.Scale = vectozavr.Scale(vectozavr.NewVec3(g.scale, g.scale, g.scale))
	g.SP = g.S.MatMul(g.P)
	g.cam.ViewMat()
	g.cam.ViewMatrix = g.cam.ViewMatrix.MatMul(g.R)
	//-----------------------------------------------------------------
	if ebiten.IsKeyPressed(ebiten.KeyEnter) {
		g.angle = 0
		g.tilt = 0
		g.scale = 1
		g.oX.v = vectozavr.NewVec3(1, 0, 0)
		g.oY.v = vectozavr.NewVec3(0, 1, 0)
		g.oZ.v = vectozavr.NewVec3(0, 0, 1)
		g.cam.E = vectozavr.NewVec3(0, 0, 0)
	}
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		g.cam.Move(vectozavr.NewVec3(0, 0.01, 0))
	}
	if ebiten.IsKeyPressed(ebiten.KeyShift) {
		g.cam.Move(vectozavr.NewVec3(0, -0.01, 0))
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.cam.Move(vectozavr.NewVec3(0.01, 0, 0))
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {

		g.cam.Move(vectozavr.NewVec3(-0.01, 0, 0))
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {

		g.cam.Move(vectozavr.NewVec3(0, 0, -0.01))
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {

		g.cam.Move(vectozavr.NewVec3(0, 0, 0.01))
	}
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		g.angle += 0.01
	}
	if ebiten.IsKeyPressed(ebiten.KeyE) {
		g.angle -= 0.01
	}
	if g.angle > 2*math.Pi {
		g.angle = 0.0
	}
	if g.angle < 0.0 {
		g.angle = 2 * math.Pi
	}

	_, delta := ebiten.Wheel()
	if delta > 0 {
		g.scale *= 1.1
	}
	if delta < 0 {
		g.scale *= 0.9
	}

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.tilt += 0.01
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.tilt -= 0.01
	}
	if g.tilt > 2*math.Pi {
		g.tilt = 0
	}
	if g.tilt < 0.0 {
		g.tilt = 2 * math.Pi
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
	}
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
	}
	if ebiten.IsKeyPressed(ebiten.Key1) {
		g.tilt = math.Pi / 2
	}
	if ebiten.IsKeyPressed(ebiten.Key2) {
		g.angle = math.Pi / 2
	}

	return nil
}

type axes struct {
	v     vectozavr.Vec3
	color color.Color
}

type Cube struct {
	A, B, C, D, E, F, G, H vectozavr.Vec4
}

func cubePrint(screen *ebiten.Image, g *Game, cube Cube) {
	cubePoints := [8]vectozavr.Vec4{cube.A, cube.B, cube.C, cube.D, cube.E, cube.F, cube.G, cube.H}
	screenCubePoint := []vectozavr.Vec4{}

	for _, screenPoint := range cubePoints {
		screenPoint = g.cam.ViewMatrix.Vec4Mul(screenPoint)
		screenPoint = g.SP.Vec4Mul(screenPoint)
		screenPoint = g.Scale.Vec4Mul(screenPoint)
		screenPoint = screenPoint.Add(vectozavr.NewVec4(float64(g.w/2), float64(g.h/2), 0, 0))
		screenCubePoint = append(screenCubePoint, screenPoint)
	}

	vector.StrokeLine(screen, float32(screenCubePoint[0].X), float32(screenCubePoint[0].Y), float32(screenCubePoint[1].X), float32(screenCubePoint[1].Y), 5, color.RGBA{255, 0, 0, 255}, false)
	vector.StrokeLine(screen, float32(screenCubePoint[0].X), float32(screenCubePoint[0].Y), float32(screenCubePoint[2].X), float32(screenCubePoint[2].Y), 5, color.RGBA{255, 0, 0, 255}, false)
	vector.StrokeLine(screen, float32(screenCubePoint[0].X), float32(screenCubePoint[0].Y), float32(screenCubePoint[4].X), float32(screenCubePoint[4].Y), 5, color.RGBA{255, 0, 0, 255}, false)

	vector.StrokeLine(screen, float32(screenCubePoint[1].X), float32(screenCubePoint[1].Y), float32(screenCubePoint[3].X), float32(screenCubePoint[3].Y), 5, color.RGBA{0, 255, 0, 255}, false)
	vector.StrokeLine(screen, float32(screenCubePoint[1].X), float32(screenCubePoint[1].Y), float32(screenCubePoint[5].X), float32(screenCubePoint[5].Y), 5, color.RGBA{0, 255, 0, 255}, false)

	vector.StrokeLine(screen, float32(screenCubePoint[2].X), float32(screenCubePoint[2].Y), float32(screenCubePoint[3].X), float32(screenCubePoint[3].Y), 5, color.RGBA{0, 0, 255, 255}, false)
	vector.StrokeLine(screen, float32(screenCubePoint[2].X), float32(screenCubePoint[2].Y), float32(screenCubePoint[6].X), float32(screenCubePoint[6].Y), 5, color.RGBA{0, 0, 255, 255}, false)

	vector.StrokeLine(screen, float32(screenCubePoint[3].X), float32(screenCubePoint[3].Y), float32(screenCubePoint[7].X), float32(screenCubePoint[7].Y), 5, color.RGBA{0, 255, 255, 255}, false)

	vector.StrokeLine(screen, float32(screenCubePoint[4].X), float32(screenCubePoint[4].Y), float32(screenCubePoint[5].X), float32(screenCubePoint[5].Y), 5, color.RGBA{255, 0, 255, 255}, false)
	vector.StrokeLine(screen, float32(screenCubePoint[4].X), float32(screenCubePoint[4].Y), float32(screenCubePoint[6].X), float32(screenCubePoint[6].Y), 5, color.RGBA{255, 0, 255, 255}, false)

	vector.StrokeLine(screen, float32(screenCubePoint[5].X), float32(screenCubePoint[5].Y), float32(screenCubePoint[7].X), float32(screenCubePoint[7].Y), 5, color.RGBA{255, 255, 0, 255}, false)

	vector.StrokeLine(screen, float32(screenCubePoint[6].X), float32(screenCubePoint[6].Y), float32(screenCubePoint[7].X), float32(screenCubePoint[7].Y), 5, color.RGBA{255, 255, 255, 255}, false)
}

// func axesPrint(screen *ebiten.Image, g *Game) {
// }

// func gridPrint(screen *ebiten.Image, g *Game, step float64, val int) {
// }

// func newPoint(v vectozavr.Vec3, g *Game) vectozavr.Vec3 {
// }

// func screenToWorld(mousePosVec4 vectozavr.Vec4, g *Game) vectozavr.Vec3 {
// }

// func point(v vectozavr.Vec3, g *Game) vectozavr.Vec3 {
// }

func (g *Game) Draw(screen *ebiten.Image) {

	cubePrint(screen, g, g.cube)

	ebitenutil.DebugPrintAt(screen, fmt.Sprintf(
		"g.Scale: %.2f, Tilt: %.2f, Angle: %.2f",
		g.scale, g.tilt, g.angle), 0, 0,
	)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf(
		"g.cam.E(cam pos): %.2f \n g.cam.ViewMatrix: %.2f",
		g.cam.E, g.cam.ViewMatrix), 0, g.h-50,
	)

	ebitenutil.DebugPrintAt(screen, fmt.Sprintf(
		"g.SP: %.f\ng.P:%.f",
		g.SP, g.P), 0, g.w/2,
	)

}

func (g *Game) Layout(w, h int) (int, int) {
	return w, h
}

func main() {
	var _ vectozavr.Vec3
	var _ object.Object
	g := NewGame()
	ebiten.SetWindowSize(g.w, g.h)
	ebiten.SetWindowTitle("Coords")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
