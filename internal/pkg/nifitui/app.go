package nifitui

import (
	"context"
	"crypto/tls"
	"dam4rus/nifi-tui/internal/pkg/nifiapi"
	"fmt"
	"net/http"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type App struct {
	app        *tview.Application
	pages      *tview.Pages
	apiClient  *nifiapi.APIClient
	cancelFunc context.CancelFunc
}

func NewApp(url, username, password string) *App {
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
		panic(err)
	}
	apiClient.GetConfig().AddDefaultHeader("Authorization", "Bearer "+accessToken)
	accessStatus, _, err := apiClient.AccessAPI.GetAccessStatus(context.Background()).Execute()
	if err != nil {
		panic(err)
	}
	fmt.Printf("logged in user: %v\n", *accessStatus.AccessStatus.Identity)

	app := tview.NewApplication()
	pages := tview.NewPages()
	application := &App{
		app:       app,
		pages:     pages,
		apiClient: apiClient,
	}

	app.SetRoot(pages, true)
	app.SetInputCapture(application.inputCapture)
	return application
}

func (a *App) Run() error {
	return a.app.Run()
}

func (a *App) EnterRootProcessGroupScreen() {
	rootProcessGroupScreen := processGroupScreen{
		app:            a,
		processGroupId: "root",
	}
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
