package main

import (
	"context"
	"crypto/tls"
	"dam4rus/nifi-tui/internal/pkg/nifiapi"
	"fmt"
	"net/http"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	configuration := nifiapi.NewConfiguration()
	configuration.HTTPClient = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	configuration.Servers = nifiapi.ServerConfigurations{
		nifiapi.ServerConfiguration{
			URL: "https://localhost:8443/nifi-api",
		},
	}
	apiClient := nifiapi.NewAPIClient(configuration)
	accessToken, _, err := apiClient.AccessAPI.CreateAccessToken(context.Background()).
		Username("aa518870-4cc8-481e-8f69-95aa12a9c4dd").
		Password("MUAQnbOWaJ0BhAIhIf+jadsllNPm4tPr").
		Execute()
	if err != nil {
		panic(err)
	}
	apiClient.GetConfig().AddDefaultHeader("Authorization", "Bearer "+accessToken)
	accessStatus, _, err := apiClient.AccessAPI.GetAccessStatus(context.Background()).Execute()
	if err != nil {
		panic(err)
	}
	fmt.Printf("logged in user: %v\n", *accessStatus.AccessStatus.Identity)

	processors, _, err := apiClient.ProcessGroupsAPI.GetProcessors(context.Background(), "root").Execute()
	if err != nil {
		panic(err)
	}
	app := tview.NewApplication()
	pages := tview.NewPages()
	list := tview.NewList()
	var cancelFunc context.CancelFunc
	for _, processor := range processors.Processors {
		processor := processor
		list.AddItem(*processor.Component.Name, *processor.Id, 0, func() {
			loadingPage := tview.NewTextView().
				SetLabel("Loading processor details")
			loadingPage.SetTitle(*processor.Component.Name).
				SetTitleAlign(tview.AlignLeft).
				SetBorder(true)
			pages.AddPage("Loading", loadingPage, true, true)

			ctx, cancel := context.WithCancel(context.Background())
			cancelFunc = cancel
			go func(ctx context.Context) {
				processorDetails, _, err := apiClient.ProcessorsAPI.GetProcessor(ctx, *processor.Id).Execute()
				app.QueueUpdateDraw(func() {
					frontPage, _ := pages.GetFrontPage()
					if frontPage != "Loading" {
						return
					}
					if err != nil {
						loadingPage.SetLabel("Failed to load processor details")
						return
					}
					pages.RemovePage("Loading")
					processorProperties := tview.NewList()
					for k, v := range *processorDetails.Component.Config.Properties {
						k, v := k, v
						processorProperties.AddItem(k, v, 0, func() {
							form := tview.NewForm().
								AddTextArea(k, v, 100, 5, 0, nil).
								AddButton("Save", func() {
									pages.RemovePage("Property")
								}).
								AddButton("Cancel", func() {
									pages.RemovePage("Property")
								})

							form.SetBorder(true).
								SetTitle("Property").
								SetTitleAlign(tview.AlignLeft)

							grid := tview.NewGrid().
								SetColumns(0, 100, 0).
								SetRows(0, 11, 0).
								AddItem(form, 1, 1, 1, 1, 0, 0, true)

							pages.AddPage("Property", grid, true, true)
						})
					}
					processorProperties.SetTitle(*processor.Component.Name)
					processorProperties.SetTitleAlign(tview.AlignLeft)
					processorProperties.SetBorder(true)
					pages.AddPage("Processor details", processorProperties, true, true)
				})
			}(ctx)
		})
	}
	list.SetTitle("Root")
	list.SetTitleAlign(tview.AlignLeft)
	list.SetBorder(true)
	pages.AddPage("Root", list, true, true)

	app.SetRoot(pages, true)
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		frontPage, _ := pages.GetFrontPage()
		if frontPage == "Root" {
			if event.Rune() == 'q' {
				app.Stop()
				return nil
			}
		} else if event.Key() == tcell.KeyEsc {
			if cancelFunc != nil {
				cancelFunc()
			}
			pages.RemovePage(frontPage)
		}
		return event
	})
	if err := app.Run(); err != nil {
		panic(err)
	}
}
