package nifitui

import (
	"context"
	"crypto/tls"
	"dam4rus/nifi-tui/internal/pkg/nifiapi"
	"fmt"
	"io"
	"net/http"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type App struct {
	app            *tview.Application
	pages          *tview.Pages
	apiClient      *nifiapi.APIClient
	processorTypes []nifiapi.DocumentedTypeDTO
	cancelFunc     context.CancelFunc
}

func NewApp(url, username, password string) (*App, error) {
	configuration := nifiapi.NewConfiguration()
	configuration.HTTPClient = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	configuration.Servers = nifiapi.ServerConfigurations{
		nifiapi.ServerConfiguration{
			URL: url,
		},
	}
	apiClient := nifiapi.NewAPIClient(configuration)
	accessToken, _, err := apiClient.AccessAPI.CreateAccessToken(context.Background()).
		Username(username).
		Password(password).
		Execute()
	if err != nil {
		return nil, err
	}
	apiClient.GetConfig().AddDefaultHeader("Authorization", "Bearer "+accessToken)
	accessStatus, _, err := apiClient.AccessAPI.GetAccessStatus(context.Background()).Execute()
	if err != nil {
		return nil, err
	}
	fmt.Printf("logged in user: %v\n", *accessStatus.AccessStatus.Identity)

	processorTypes, _, err := apiClient.FlowAPI.GetProcessorTypes(context.Background()).
		Execute()

	if err != nil {
		return nil, err
	}

	app := tview.NewApplication()
	pages := tview.NewPages()
	application := &App{
		app:            app,
		pages:          pages,
		apiClient:      apiClient,
		processorTypes: processorTypes.ProcessorTypes,
	}

	app.SetRoot(pages, true)
	app.SetInputCapture(application.inputCapture)
	return application, nil
}

func (a *App) Run() error {
	return a.app.Run()
}

func (a *App) EnterRootProcessGroupScreen() {
	rootProcessGroupScreen := newProcessGroupScreen(a, "root")
	rootProcessGroupScreen.enter()
}

func (a *App) inputCapture(event *tcell.EventKey) *tcell.EventKey {
	frontPage, _ := a.pages.GetFrontPage()
	if frontPage == "root" {
		if event.Rune() == 'q' {
			a.app.Stop()
			return nil
		}
	}
	return event
}

func (a *App) enterLoadingScreen(title, label string) *tview.TextView {
	loadingPage := tview.NewTextView().
		SetLabel(label)
	loadingPage.SetTitle(title).
		SetTitleAlign(tview.AlignLeft).
		SetBorder(true)
	a.pages.AddAndSwitchToPage("Loading", loadingPage, true)
	return loadingPage
}

func (a *App) showErrorDialog(response *http.Response, err error) {
	if err == nil {
		panic("`err` must not be nil")
	}
	body, readErr := io.ReadAll(response.Body)
	modal := tview.NewModal().
		AddButtons([]string{"Ok"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			a.pages.RemovePage("Error")
		})
	modal.SetTitle("Error").
		SetBorder(true)
	if readErr == nil {
		modal.SetText(string(body))
	} else {
		modal.SetText(err.Error())
	}
	a.pages.AddPage("Error", modal, true, true)
}
