package models

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"time"
)

type Game struct {
	Resources     int
	PriceUpgrade1 int
	Upgrade       int
	CPS           int
	Slave         int
	Autoclick     int
	FastClick     int
	GameWindow    fyne.Window
	ResourceLabel widget.Label
	Speed         int
}

func NewGame(GameWindow fyne.Window) *Game {
	return &Game{
		Resources:     0,
		PriceUpgrade1: 100,
		Upgrade:       2,
		CPS:           1,
		Slave:         1,
		Autoclick:     70,
		FastClick:     200,
		Speed:         1000,
		GameWindow:    GameWindow,
	}
}

func (g *Game) Latigazo() {
	if g.Speed > 200 {
		g.Resources -= 100
		g.Speed -= 210
		g.FastClick *= 2
	}
}

func (g *Game) ClickUpgrade() {
	if g.Resources >= g.PriceUpgrade1 {
		g.Resources -= g.PriceUpgrade1
		g.CPS *= 2
		g.PriceUpgrade1 *= 4
		g.Upgrade++
	}
}

func (g *Game) Click() {
	g.Resources += g.CPS
	fmt.Print(g.Resources)
}

func (g *Game) WaitForAutoclick() {
	g.Resources -= g.Autoclick
	g.Autoclick *= 2
	g.Slave++
	for {
		time.Sleep(time.Duration(g.Speed) * time.Millisecond)
		g.Click()
	}
}

func (g *Game) ReloadResources(ResourceLabel *widget.Label) {
	for {
		time.Sleep(60 * time.Millisecond)
		if g.Resources < 0 {
			goodbye := canvas.NewText("Has perdido", color.White)
			goodbye.TextSize = 20
			g.GameWindow.SetContent(container.NewCenter(goodbye))
			time.Sleep(1 * time.Second)
			g.GameWindow.Close()
		}
		ResourceLabel.SetText("Recursos: " + fmt.Sprint(g.Resources))
	}
}

func (g *Game) ReloadPrices(ButtonUpgrade, ButtonAuto, FastClickButton *widget.Button) {
	for {
		ButtonUpgrade.SetText("Mejora de clic x" + fmt.Sprint(g.Upgrade) + " (Costo: " + fmt.Sprint(g.PriceUpgrade1) + ")")
		ButtonAuto.SetText("Comprar " + fmt.Sprint(g.Slave) + " esclavos (Costo: " + fmt.Sprint(g.Autoclick) + ")")
		FastClickButton.SetText("Latigazo " + "(Costo: " + fmt.Sprint(g.FastClick) + ")")
	}
}

func (g *Game) Timer(TimerInGame *widget.Label) {
	sec := 0
	for sec >= 0 {
		TimerInGame.SetText(fmt.Sprintf("Time: %.2d:%.2d\r", sec/60, sec%60))
		time.Sleep(1 * time.Second)
		sec++
	}
}

func (g *Game) FinishGame() {
	g.GameWindow.Close()
}
