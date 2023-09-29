package main

import (
	"awesomeProject/scenes"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var scene *screen.ClickerGame

func main() {
	a := app.New()
	w := a.NewWindow("Minecraft chido")
	buttonStart := widget.NewButton("Start", func() {
		scene := screen.NewClickerGame(w)
		scene.Start()
	})
	containerInit := container.NewGridWrap(fyne.NewSize(100, 50), buttonStart)
	w.Resize(fyne.NewSize(900, 600))
	w.SetFixedSize(true)
	img, _ := fyne.LoadResourceFromPath("assets/minecraft.png")
	imageBg := canvas.NewImageFromResource(img)
	imageBg.Resize(fyne.NewSize(900, 600))
	containerInit.Move(fyne.NewPos(400, 300))
	containerObjects := container.NewWithoutLayout(imageBg, containerInit)
	w.SetContent(containerObjects)
	w.Show()
	a.Run()
}
