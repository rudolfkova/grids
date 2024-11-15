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

	invP vectozavr.Matrix
	invS vectozavr.Matrix

	angle, tilt, roll float64
	pos               vectozavr.Vec4

	cam    camera.Camera
	visual bool

	pointXY []vectozavr.Vec3
	pointXZ []vectozavr.Vec3
	pointYZ []vectozavr.Vec3
}

func NewGame() *Game {
	g := &Game{
		scale: 1.0,
		w:     1000,
		h:     700,
	}
	g.pos = vectozavr.NewVec4(0, 0, 4, 1)
	g.P = vectozavr.Projection(60, float64(g.w)/float64(g.h), 1, 10)
	g.invP, _ = g.P.Inverse()
	g.S = vectozavr.ScreenSpace(float64(g.w), float64(g.h))
	g.invS, _ = g.S.Inverse()
	g.cam.At = vectozavr.NewVec3(0, 0, 1)
	g.cam.Up = vectozavr.NewVec3(0, 1, 0)
	g.cam.Left = vectozavr.NewVec3(1, 0, 0)
	g.cam.InitCamera()

	return g
}

func (g *Game) ScreenToWorld(mousePos vectozavr.Vec2) (XY, XZ, YZ vectozavr.Vec3) {
	tmp := vectozavr.NewVec4(((2.0*mousePos.X)/float64(g.w))-1.0, ((2.0*mousePos.Y)/float64(g.h))-1.0, -1, 1)
	itmp := g.invP.Vec4Mul(tmp)
	tmp = vectozavr.NewVec4(itmp.X, itmp.Y, -1, 0)
	direction, _ := g.cam.InverseViewMatrix.Vec3Mul(tmp.ToVec3()).Normalize()
	camPos := g.cam.InverseViewMatrix.Vec4Mul(vectozavr.NewVec4(0, 0, 0, 1)).ToVec3()

	N := vectozavr.NewVec3(0, 0, 1)
	t := -(camPos.Dot(N)) / direction.Dot(N)
	resultXY := camPos.Add(direction.Mul(t))

	N = vectozavr.NewVec3(0, 1, 0)
	t = -(camPos.Dot(N)) / direction.Dot(N)
	resultXZ := camPos.Add(direction.Mul(t))

	N = vectozavr.NewVec3(1, 0, 0)
	t = -(camPos.Dot(N)) / direction.Dot(N)
	resultYZ := camPos.Add(direction.Mul(t))
	return resultXY, resultXZ, resultYZ
}

func (g *Game) DrawGrid(screen *ebiten.Image, step float64, num float64) {
	for i := -num; i <= num; i++ {
		g.ProjLine(screen, vectozavr.NewVec3(i*step, -5, 0), vectozavr.NewVec3(i*step, 5, 0), vectozavr.NewVec3(0, 0, 0), color.RGBA{255, 0, 0, 255})
		g.ProjLine(screen, vectozavr.NewVec3(-5, i*step, 0), vectozavr.NewVec3(5, i*step, 0), vectozavr.NewVec3(0, 0, 0), color.RGBA{255, 0, 0, 255})
	}
	for i := -num; i <= num; i++ {
		g.ProjLine(screen, vectozavr.NewVec3(0, -5, i*step), vectozavr.NewVec3(0, 5, i*step), vectozavr.NewVec3(0, 0, 0), color.RGBA{0, 0, 255, 255})
		g.ProjLine(screen, vectozavr.NewVec3(0, i*step, -5), vectozavr.NewVec3(0, i*step, 5), vectozavr.NewVec3(0, 0, 0), color.RGBA{0, 0, 255, 255})
	}
	for i := -num; i <= num; i++ {
		g.ProjLine(screen, vectozavr.NewVec3(-5, 0, i*step), vectozavr.NewVec3(5, 0, i*step), vectozavr.NewVec3(0, 0, 0), color.RGBA{0, 255, 0, 255})
		g.ProjLine(screen, vectozavr.NewVec3(i*step, 0, -5), vectozavr.NewVec3(i*step, 0, 5), vectozavr.NewVec3(0, 0, 0), color.RGBA{0, 255, 0, 255})
	}
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

func (g *Game) DrawProjPoint(screen *ebiten.Image, p vectozavr.Vec3, color color.Color) {
	pVec4 := g.ProjPoint(p)
	if !g.visual {
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf(
			"X:%.f\nY:%.f\nZ:%.f",
			p.X, p.Y, p.Z), int(pVec4.X), int(pVec4.Y),
		)
	}
	vector.DrawFilledCircle(screen, float32(pVec4.X), float32(pVec4.Y), 10, color, false)
}

func (g *Game) ProjLine(screen *ebiten.Image, p1, p2 vectozavr.Vec3, pos vectozavr.Vec3, color color.Color) {
	//  = g.S.Vec4Mul(g.P.Vec4Mul(p.ToVec4().Add(g.pos)))
	p1Vec4 := g.ProjPoint(p1.Add(pos))
	p2Vec4 := g.ProjPoint(p2.Add(pos))

	if !g.visual {
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf(
			"X:%.f\nY:%.f",
			p1.X, p1.Y), int(p1Vec4.X), int(p1Vec4.Y),
		)
	}
	if !g.visual {
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf(
			"X:%.f\nY:%.f",
			p2.X, p2.Y), int(p2Vec4.X), int(p2Vec4.Y),
		)
	}

	vector.StrokeLine(screen, float32(p1Vec4.X), float32(p1Vec4.Y), float32(p2Vec4.X), float32(p2Vec4.Y), 1, color, false)
}

func (g *Game) keys() {
	if ebiten.IsKeyPressed(ebiten.KeyEnter) {
		g.cam.E = vectozavr.NewVec3(0, 0, 0)

	}
	if inpututil.IsKeyJustPressed(ebiten.Key5) {
		g.visual = !g.visual
	}
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		g.cam.Move(vectozavr.NewVec3(0, 0.05, 0))
	}
	if ebiten.IsKeyPressed(ebiten.KeyShift) {
		g.cam.Move(vectozavr.NewVec3(0, -0.05, 0))
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		dv, _ := g.cam.Left.Div(25.0)
		dv.Y = 0
		g.cam.Move(dv)
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {

		dv, _ := g.cam.Left.Div(-25.0)
		dv.Y = 0
		g.cam.Move(dv)
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		dv, _ := g.cam.At.Div(25.0)
		dv.Y = 0
		g.cam.Move(dv)
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {

		dv, _ := g.cam.At.Div(-25.0)
		dv.Y = 0
		g.cam.Move(dv)
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.cam.Rotate(0, 0.03)
		g.angle += 0.03
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.cam.Rotate(0, -0.03)
		g.angle -= 0.03
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
		g.cam.Rotate(0.03, 0)
		g.tilt += 0.03
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.cam.Rotate(-0.03, 0)
		g.tilt -= 0.03
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
		x, y := ebiten.CursorPosition()
		mousePos := vectozavr.NewVec2(float64(x), float64(y))
		XY, XZ, YZ := g.ScreenToWorld(mousePos)
		g.pointXY = append(g.pointXY, XY)
		g.pointXZ = append(g.pointXZ, XZ)
		g.pointYZ = append(g.pointYZ, YZ)

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
	g.cam.Tilt = g.tilt
	g.cam.Angle = g.angle
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

	// g.ProjLine(screen, vectozavr.NewVec3(1, -1, 0), vectozavr.NewVec3(0, 1, 0), vectozavr.NewVec3(0, 0, 4), color.RGBA{255, 0, 0, 255})
	// g.ProjLine(screen, vectozavr.NewVec3(0, 1, 0), vectozavr.NewVec3(-1, -1, 0), vectozavr.NewVec3(0, 0, 4), color.RGBA{255, 0, 0, 255})
	// g.ProjLine(screen, vectozavr.NewVec3(-1, -1, 0), vectozavr.NewVec3(1, -1, 0), vectozavr.NewVec3(0, 0, 4), color.RGBA{255, 0, 0, 255})

	// g.DrawProjPoint(screen, vectozavr.NewVec3(0, 0, 0), color.White)

	// g.ProjLine(screen, vectozavr.NewVec3(1, -1, 1), vectozavr.NewVec3(0, 1, 1), vectozavr.NewVec3(0, 0, 4), color.RGBA{255, 0, 0, 255})
	// g.ProjLine(screen, vectozavr.NewVec3(0, 1, 1), vectozavr.NewVec3(-1, -1, 1), vectozavr.NewVec3(0, 0, 4), color.RGBA{255, 0, 0, 255})
	// g.ProjLine(screen, vectozavr.NewVec3(-1, -1, 1), vectozavr.NewVec3(1, -1, 1), vectozavr.NewVec3(0, 0, 4), color.RGBA{255, 0, 0, 255})

	// g.ProjLine(screen, vectozavr.NewVec3(-1, -1, 0), vectozavr.NewVec3(-1, -1, 1), vectozavr.NewVec3(0, 0, 4), color.RGBA{255, 0, 0, 255})
	// g.ProjLine(screen, vectozavr.NewVec3(0, 1, 0), vectozavr.NewVec3(0, 1, 1), vectozavr.NewVec3(0, 0, 4), color.RGBA{255, 0, 0, 255})
	// g.ProjLine(screen, vectozavr.NewVec3(1, -1, 0), vectozavr.NewVec3(1, -1, 1), vectozavr.NewVec3(0, 0, 4), color.RGBA{255, 0, 0, 255})

	g.DrawGrid(screen, 0.5, 10)

	for _, p := range g.pointXY {
		g.DrawProjPoint(screen, p, color.RGBA{255, 0, 0, 255})
	}
	for _, p := range g.pointXZ {
		g.DrawProjPoint(screen, p, color.RGBA{0, 255, 0, 255})
	}
	for _, p := range g.pointYZ {
		g.DrawProjPoint(screen, p, color.RGBA{0, 0, 255, 255})
	}

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
