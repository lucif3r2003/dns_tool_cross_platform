package main

import (
	"dns_tool_cross_platform/internal/dns"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("DNS Switcher")

	googleBtn := widget.NewButton("Set Google DNS", func() {
		err := dns.SetGoogle()
		if err != nil {
			dialog.ShowError(err, w)
		} else {
			dialog.ShowInformation("Thành công", "Đã đổi sang Google DNS", w)
		}
	})

	defaultBtn := widget.NewButton("Set Default DNS", func() {
		err := dns.SetDefault()
		if err != nil {
			dialog.ShowError(err, w)
		} else {
			dialog.ShowInformation("Thành công", "Đã đổi về DNS mặc định", w)
		}
	})

	w.SetContent(
		container.NewVBox(
			widget.NewLabel("Chọn DNS muốn đổi:"),
			googleBtn,
			defaultBtn,
		),
	)

	w.Resize(fyne.NewSize(300, 200))
	fmt.Println("App starting...")
	w.ShowAndRun()
}
