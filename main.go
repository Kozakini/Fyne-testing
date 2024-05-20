package main
import(
    "fmt"
    "image/color"

    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/canvas"
    "fyne.io/fyne/v2/theme"
    "fyne.io/fyne/v2/layout"
    "fyne.io/fyne/v2/widget"
    "fyne.io/fyne/v2/container"
)



func main(){
    var x uint8 = 0
    a := app.New()
    w := a.NewWindow("Traffic Light")
    w.Resize(fyne.NewSize(1000,1000))
    fyne.CurrentApp().Settings().SetTheme(theme.DarkTheme())

    kolko1 := canvas.NewCircle(color.NRGBA{R:255,G:255,B:255,A:255})
    kolko1.Resize(fyne.NewSize(50,50))
    kolko2 := canvas.NewCircle(color.NRGBA{R:255,G:255,B:255,A:255})
    kolko2.Resize(fyne.NewSize(50,50))
    kolko3 := canvas.NewCircle(color.NRGBA{R:255,G:255,B:255,A:255})
    kolko3.Resize(fyne.NewSize(50,50))

    buttonR := widget.NewButton("Red",func(){
        fmt.Printf("%t",true)
        reset(kolko1,kolko2,kolko3)
        kolko1.FillColor=color.NRGBA{R:incrementMove(x),G:0,B:0,A:255}
        kolko1.Refresh()
    })

    buttonG := widget.NewButton("Green",func(){
        fmt.Printf("%t",true)
        reset(kolko1,kolko2,kolko3)
        kolko2.FillColor=color.NRGBA{R:0,G:incrementMove(x),B:0,A:255}
        kolko2.Refresh()
    })

    buttonB := widget.NewButton("Green",func(){
        fmt.Printf("%t",true)
        reset(kolko1,kolko2,kolko3)
        kolko3.FillColor=color.NRGBA{R:0,G:0,B:incrementMove(x),A:255}
        kolko3.Refresh()
    })
    buttonReset:=widget.NewButton("Reset",func(){
        kolko1.FillColor=color.NRGBA{R:255,G:255,B:255,A:255}
        kolko2.FillColor=color.NRGBA{R:255,G:255,B:255,A:255}
        kolko3.FillColor=color.NRGBA{R:255,G:255,B:255,A:255}
        kolko1.Refresh()
        kolko2.Refresh()
        kolko3.Refresh()
    })

    w.SetContent(
        container.NewHSplit(
            container.NewGridWithRows(
                7,
                layout.NewSpacer(),
                kolko1,
                layout.NewSpacer(),
                kolko2,
                layout.NewSpacer(),
                kolko3,
                layout.NewSpacer(),
            ),
            container.NewGridWithRows(
                9,
                layout.NewSpacer(),
                buttonR,
                //layout.NewSpacer(),
                buttonG,
                //layout.NewSpacer(),
                buttonB,
                buttonReset,
                layout.NewSpacer(),
            ),
        ),

    )
    w.ShowAndRun()
}

func incrementMove(x uint8) (a uint8){
    if x == 255{
        x=0
    } else if x== 0{
        x=255
    }
    return x
}

func reset(kolko1,kolko2,kolko3*canvas.Circle){
    kolko1.FillColor=color.NRGBA{R:255,G:255,B:255,A:255}
    kolko2.FillColor=color.NRGBA{R:255,G:255,B:255,A:255}
    kolko3.FillColor=color.NRGBA{R:255,G:255,B:255,A:255}
    kolko1.Refresh()
    kolko2.Refresh()
    kolko3.Refresh()
}







