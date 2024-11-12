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
	roll  float64

	P      vectozavr.Matrix
	R      vectozavr.Matrix
	S      vectozavr.Matrix
	Scale  vectozavr.Matrix
	worldE vectozavr.Matrix
	SP     vectozavr.Matrix

	oX, oY, oZ          axes
	newOX, newOY, newOZ axes

	cube  Cube
	place Place

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
	pos := vectozavr.NewVec4(10, b-1, -10, 0)
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
	g.place = Place{
		A: vectozavr.NewVec3(10, -1, -10),
		B: vectozavr.NewVec3(10, -1, 10),
		C: vectozavr.NewVec3(-10, -1, 10),
		D: vectozavr.NewVec3(-10, -1, -10),

		color: color.RGBA{255, 0, 0, 255},
	}

	g.cam.A = float64(g.w) / float64(g.h)
	g.cam.Far = 100
	g.cam.Near = 0
	g.cam.Fov = 90
	g.cam.Left = vectozavr.NewVec3(0, 0, 1)
	g.cam.At = g.oZ.v
	g.cam.Up = g.oY.v

	g.P = vectozavr.Projection(g.cam.Fov, g.cam.A, g.cam.Near, g.cam.Far)
	g.S = vectozavr.ScreenSpace(float64(g.w), float64(g.h))

	return g
}

func (g *Game) Update() error {
	g.R = vectozavr.RotationY(g.angle).MatMul(vectozavr.RotationX(g.tilt)).MatMul(vectozavr.RotationZ(g.roll))
	g.Scale = vectozavr.Scale(vectozavr.NewVec3(g.scale, g.scale, g.scale))
	g.SP = g.P.MatMul(g.S)
	g.cam.ViewMat()
	g.cam.ViewMatrix = g.cam.ViewMatrix.MatMul(g.R)
	g.cam.Left, _ = g.oX.v.Normalize()
	g.cam.At, _ = g.oZ.v.Normalize()
	g.cam.Up, _ = g.oY.v.Normalize()
	g.newAxes()
	// g.cam.Left = g.newOX.v
	// g.cam.At = g.newOZ.v
	// g.cam.Up = g.newOY.v
	//-----------------------------------------------------------------
	if ebiten.IsKeyPressed(ebiten.KeyEnter) {
		g.angle = 0
		g.tilt = 0
		g.scale = 1
		g.roll = 0
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
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.roll += 0.01
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.roll -= 0.01
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
type Place struct {
	A, B, C, D vectozavr.Vec3
	color      color.Color
}

func cubePrint(screen *ebiten.Image, g *Game, cube Cube) {
	cubePoints := [8]vectozavr.Vec4{cube.A, cube.B, cube.C, cube.D, cube.E, cube.F, cube.G, cube.H}
	screenCubePoint := []vectozavr.Vec4{}

	for _, screenPoint := range cubePoints {
		screenPoint = g.worldE.Vec4Mul(screenPoint)
		// screenPoint = g.cam.ViewMatrix.Vec4Mul(screenPoint)
		// screenPoint = g.SP.Vec4Mul(screenPoint)
		// screenPoint = g.Scale.Vec4Mul(screenPoint)
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

func placePrint(screen *ebiten.Image, g *Game, p Place) {
	var s [4]vectozavr.Vec4 = [4]vectozavr.Vec4{
		p.A.ToVec4(), p.B.ToVec4(), p.C.ToVec4(), p.D.ToVec4(),
	}
	var s2 []vectozavr.Vec4
	for i := 0; i < 4; i++ {
		screenPoint := g.worldE.Vec4Mul(s[i])
		screenPoint = screenPoint.Add(vectozavr.NewVec4(float64(g.w/2), float64(g.h/2), 0, 0))
		s2 = append(s2, screenPoint)
	}
	vector.StrokeLine(screen, float32(s2[0].X), float32(s2[0].Y), float32(s2[1].X), float32(s2[1].Y), 5, color.RGBA{255, 0, 0, 255}, false)
	vector.StrokeLine(screen, float32(s2[1].X), float32(s2[1].Y), float32(s2[2].X), float32(s2[2].Y), 5, color.RGBA{255, 0, 0, 255}, false)
	vector.StrokeLine(screen, float32(s2[2].X), float32(s2[2].Y), float32(s2[3].X), float32(s2[3].Y), 5, color.RGBA{255, 0, 0, 255}, false)
	vector.StrokeLine(screen, float32(s2[3].X), float32(s2[3].Y), float32(s2[0].X), float32(s2[0].Y), 5, color.RGBA{255, 0, 0, 255}, false)
}

// func (g *Game) axesPrint(screen *ebiten.Image) {
// g.newAxes()
// vector.StrokeLine(screen, float32(g.newOX.v.Y), float32(g.newOX.v.Z), float32(g.newOX.v.X), float32(g.newOX.v.Z), 2, color.RGBA{255, 0, 0, 255}, false)
// }

// func gridPrint(screen *ebiten.Image, g *Game, step float64, val int) {
// }

func (g *Game) newAxes() {

	// eA := vectozavr.NewMatrix([4][4]float64{
	// 	{g.oX.v.X, g.oY.v.X, g.oZ.v.X, 0},
	// 	{g.oX.v.Y, g.oY.v.Y, g.oZ.v.Y, 0},
	// 	{g.oX.v.Z, g.oY.v.Z, g.oZ.v.Z, 0},
	// 	{0, 0, 0, 1},
	// })
	// eA = g.SP.MatMul(eA)
	// eA = g.cam.ViewMatrix.MatMul(eA)
	// eA = g.Scale.MatMul(eA)
	// dv := vectozavr.Translation(vectozavr.ZeroVec3().Sub(g.cam.E))
	// eA = dv.MatMul(eA)
	// g.worldE = eA
	// g.newOX.v = eA.X()
	// g.newOY.v = eA.Y()
	// g.newOZ.v = eA.Z()

	g.newOX = g.oX
	g.newOY = g.oY
	g.newOZ = g.oZ

	g.newOX.v = g.cam.ViewMatrix.Vec3Mul(g.newOX.v)
	g.newOY.v = g.cam.ViewMatrix.Vec3Mul(g.newOY.v)
	g.newOZ.v = g.cam.ViewMatrix.Vec3Mul(g.newOZ.v)

	g.newOX.v = g.SP.Vec3Mul(g.newOX.v)
	g.newOY.v = g.SP.Vec3Mul(g.newOY.v)
	g.newOZ.v = g.SP.Vec3Mul(g.newOZ.v)

	g.newOX.v = g.Scale.Vec3Mul(g.newOX.v)
	g.newOY.v = g.Scale.Vec3Mul(g.newOY.v)
	g.newOZ.v = g.Scale.Vec3Mul(g.newOZ.v)

	g.worldE = vectozavr.NewMatrixVec3(g.newOX.v, g.newOY.v, g.newOZ.v)

	fmt.Println(g.worldE)

}

// func screenToWorld(mousePosVec4 vectozavr.Vec4, g *Game) vectozavr.Vec3 {
// }

// func point(v vectozavr.Vec3, g *Game) vectozavr.Vec3 {
// }

func (g *Game) Draw(screen *ebiten.Image) {

	// g.axesPrint(screen)
	placePrint(screen, g, g.place)
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
