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

	// Label to show current DNS
	dnsLabel := widget.NewLabel("Current DNS: (loading...)")

	// Load current DNS
	current, err := dns.GetCurrentDNS()
	if err != nil {
		dnsLabel.SetText("Current DNS: [Error: " + err.Error() + "]")
	} else {
		dnsLabel.SetText("Current DNS: " + current)
	}

	googleBtn := widget.NewButton("Set Google DNS", func() {
		err := dns.SetGoogle()
		if err != nil {
			dialog.ShowError(err, w)
		} else {
			dialog.ShowInformation("Success", "DNS has been changed to Google DNS", w)
			newDNS, _ := dns.GetCurrentDNS()
			dnsLabel.SetText("Current DNS: " + newDNS)
		}
	})

	defaultBtn := widget.NewButton("Set Default DNS", func() {
		err := dns.SetDefault()
		if err != nil {
			dialog.ShowError(err, w)
		} else {
			dialog.ShowInformation("Success", "DNS has been reset to default", w)
			newDNS, _ := dns.GetCurrentDNS()
			dnsLabel.SetText("Current DNS: " + newDNS)
		}
	})

	w.SetContent(
		container.NewVBox(
			widget.NewLabel("Choose the DNS you want to set:"),
			dnsLabel,
			googleBtn,
			defaultBtn,
		),
	)

	w.Resize(fyne.NewSize(350, 220))
	fmt.Println("App starting...")
	w.ShowAndRun()
}
