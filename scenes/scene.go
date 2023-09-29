package screen

import (
	"awesomeProject/models"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type ClickerGame struct {
	Window         fyne.Window
	AutoClickState bool
	gameRender     *models.Game
}

func NewClickerGame(w fyne.Window) *ClickerGame {
	gameRender := models.NewGame(w)
	return &ClickerGame{
		Window:         w,
		AutoClickState: false,
		gameRender:     gameRender,
	}
}

func (game *ClickerGame) Start() {
	imagePath, _ := fyne.LoadResourceFromPath("assets/shovelandshovel.png")
	imageShovel := canvas.NewImageFromResource(imagePath)
	resourceLabel := widget.NewLabel("Recursos: 0")
	timeInGame := widget.NewLabel("Time: 0")

	buttonAuto := widget.NewButton("Comprar 1 esclavo (Costo: 70)", func() {
		game.AutoClickState = true
		if game.AutoClickState {
			go game.gameRender.WaitForAutoclick()
		}
	})

	buttonUpgrade := widget.NewButton("", func() {
		game.gameRender.ClickUpgrade()
	})

	fastClickbutton := widget.NewButton("Latigazos (Costo: 100)", func() {
		game.gameRender.Latigazo()
	})

	endClick := widget.NewButton("Finalizar (Costo: 1000)", func() {
		game.gameRender.FinishGame()
	})

	buttonClick := widget.NewButton("Puchale aquí", func() {
		game.gameRender.Click()
	})

	clickContainer := container.NewGridWrap(fyne.NewSize(450, 300), buttonClick)
	leftContainer := container.NewGridWrap(fyne.NewSize(450, 300), container.NewVBox(timeInGame, resourceLabel, buttonAuto, buttonUpgrade, fastClickbutton, endClick))
	leftContainerFinal := container.NewVBox(leftContainer, clickContainer)
	centerContainer := container.NewGridWrap(fyne.NewSize(450, 500), imageShovel)
	content := container.NewHBox(leftContainerFinal, container.NewCenter(centerContainer))
	game.Window.SetContent(content)

	go game.gameRender.Timer(timeInGame)
	go game.gameRender.ReloadResources(resourceLabel)
	go game.gameRender.ReloadPrices(buttonUpgrade, buttonAuto, fastClickbutton)
}
