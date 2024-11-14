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
	scale float64
	w     int
	h     int

	P vectozavr.Matrix
	S vectozavr.Matrix

	angle, tilt, roll float64
	pos               vectozavr.Vec4

	cam camera.Camera
}

func NewGame() *Game {
	g := &Game{
		scale: 1.0,
		w:     700,
		h:     700,
	}
	g.P = vectozavr.Projection(90, float64(g.w)/float64(g.h), 1, 3)
	g.S = vectozavr.ScreenSpace(float64(g.w), float64(g.h))
	g.cam.At = vectozavr.NewVec3(0, 0, 1)
	g.cam.Up = vectozavr.NewVec3(0, 1, 0)
	g.cam.Left = vectozavr.NewVec3(1, 0, 0)

	return g
}

func (g *Game) ProjPoint(p vectozavr.Vec3) vectozavr.Vec4 {
	var newPoint vectozavr.Vec4
	//  = g.S.Vec4Mul(g.P.Vec4Mul(p.ToVec4().Add(g.pos)))
	newPoint = p.ToVec4()
	newPoint = g.cam.ViewMatrix.Vec4Mul(newPoint)
	newPoint = g.P.Vec4Mul(newPoint)
	newPoint, _ = newPoint.Div(newPoint.W)
	newPoint = g.S.Vec4Mul(newPoint)

	return newPoint
}

func (g *Game) keys() {
	if ebiten.IsKeyPressed(ebiten.KeyEnter) {
		g.angle = 0
		g.tilt = 0
		g.scale = 1
		g.roll = 0
		g.cam.E = vectozavr.NewVec3(0, 0, 0)
	}
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		g.cam.Move(vectozavr.NewVec3(0, 0.1, 0))
	}
	if ebiten.IsKeyPressed(ebiten.KeyShift) {
		g.cam.Move(vectozavr.NewVec3(0, -0.1, 0))
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.cam.Move(vectozavr.NewVec3(0.1, 0, 0))
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {

		g.cam.Move(vectozavr.NewVec3(-0.1, 0, 0))
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {

		g.cam.Move(vectozavr.NewVec3(0, 0, 0.1))
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {

		g.cam.Move(vectozavr.NewVec3(0, 0, -0.1))
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
		//
	}
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		//
	}
	if ebiten.IsKeyPressed(ebiten.Key1) {
		g.tilt = math.Pi / 2
	}
	if ebiten.IsKeyPressed(ebiten.Key2) {
		g.angle = math.Pi / 2
	}
}

func (g *Game) Update() error {
	g.cam.ViewMat()
	//-----------------------------------------------------------------
	g.keys()

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

func (g *Game) Draw(screen *ebiten.Image) {

	p1 := g.ProjPoint(vectozavr.NewVec3(1, -1, 0))
	p2 := g.ProjPoint(vectozavr.NewVec3(0, 1, 0))
	p3 := g.ProjPoint(vectozavr.NewVec3(-1, -1, 0))

	vector.StrokeLine(screen, float32(p1.X), float32(p1.Y), float32(p2.X), float32(p2.Y), 2, color.RGBA{255, 0, 0, 255}, false)
	vector.StrokeLine(screen, float32(p2.X), float32(p2.Y), float32(p3.X), float32(p3.Y), 2, color.RGBA{255, 0, 0, 255}, false)
	vector.StrokeLine(screen, float32(p3.X), float32(p3.Y), float32(p1.X), float32(p1.Y), 2, color.RGBA{255, 0, 0, 255}, false)

	ebitenutil.DebugPrintAt(screen, fmt.Sprintf(
		"g.Scale: %.2f, Tilt: %.2f, Angle: %.2f",
		g.scale, g.tilt, g.angle), 0, 0,
	)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf(
		"g.cam.E(cam pos): %.2f \n g.cam.ViewMatrix: %.2f",
		g.cam.E, g.cam.ViewMatrix), 0, g.h-50,
	)

	ebitenutil.DebugPrintAt(screen, fmt.Sprintf(
		"A: %.f\nB:%.f",
		g.pos, g.P), 0, g.w/2,
	)
	// ebitenutil.DebugPrintAt(screen, fmt.Sprintf(
	// 	"g.worldE: %.2f,\n pos: %.2f",
	// 	g.worldE, g.pos), g.w/2, g.h/2,
	// )

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
