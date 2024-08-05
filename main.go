package main

import (
	"embed"
	"log"

	"github.com/bruxaodev/overlay-for-chat/config"
	hook "github.com/robotn/gohook"
	"github.com/wailsapp/wails/v3/pkg/application"
)

// Wails uses Go's `embed` package to embed the frontend files into the binary.
// Any files in the frontend/dist folder will be embedded into the binary and
// made available to the frontend.
// See https://pkg.go.dev/embed for more information.

//go:embed frontend/dist
var assets embed.FS
var frameless = false
var configPath = "config.json"
var conf *config.Config
var Visible = true

// main function serves as the application's entry point. It initializes the application, creates a window,
// and starts a goroutine that emits a time-based event every second. It subsequently runs the application and
// logs any error that might occur.
func main() {
	conf, _err := config.LoadConfig(configPath)
	if _err != nil {
		log.Fatal(_err)
	}
	// Create a new Wails application by providing the necessary options.
	// Variables 'Name' and 'Description' are for application metadata.
	// 'Assets' configures the asset server with the 'FS' variable pointing to the frontend files.
	// 'Bind' is a list of Go struct instances. The frontend has access to the methods of these instances.
	// 'Mac' options tailor the application when running an macOS.
	app := application.New(application.Options{
		Name:        "Chat",
		Description: "chat",
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assets),
		},
		Mac: application.MacOptions{
			ApplicationShouldTerminateAfterLastWindowClosed: true,
		},
	})

	// Create a new window with the necessary options.
	// 'Title' is the title of the window.
	// 'Mac' options tailor the window when running on macOS.
	// 'BackgroundColour' is the background colour of the window.
	// 'URL' is the URL that will be loaded into the webview.
	overlay := app.NewWebviewWindowWithOptions(application.WebviewWindowOptions{
		Title: "Chat",
		Mac: application.MacWindow{
			InvisibleTitleBarHeight: 50,
			Backdrop:                application.MacBackdropTranslucent,
			TitleBar:                application.MacTitleBarHiddenInset,
		},
		BackgroundType: application.BackgroundTypeTransparent,
		Width:          conf.Width,
		Height:         conf.Height,
		AlwaysOnTop:    true,
		Frameless:      conf.Frameless,
		DisableResize:  false,
		URL:            conf.Link,
		X:              conf.X,
		Y:              conf.Y,
	})

	go func() {
		for {
			hook.Register(hook.KeyDown, conf.ShowHideWindowKey, func(e hook.Event) {
				Visible = !Visible
				if Visible {
					overlay.Show()
				} else {
					overlay.Hide()
				}
			})
			hook.Register(hook.KeyDown, conf.HideBarAndSaveKey, func(e hook.Event) {
				overlay.Show()
				frameless = !frameless
				overlay.SetFrameless(!frameless)
				x, y := overlay.Position()
				w, h := overlay.Size()
				config.SaveConfig(configPath, &config.Config{
					Frameless:         frameless,
					X:                 x,
					Y:                 y,
					Width:             w,
					Height:            h,
					Link:              conf.Link,
					HideBarAndSaveKey: conf.HideBarAndSaveKey,
					ShowHideWindowKey: conf.ShowHideWindowKey,
				})
			})

			s := hook.Start()
			<-hook.Process(s)

			// Sleep for 20 milliseconds before checking the event again.
			// time.Sleep(20 * time.Millisecond)
		}
	}()

	// Run the application. This blocks until the application has been exited.
	err := app.Run()

	// If an error occurred while running the application, log it and exit.
	if err != nil {
		log.Fatal(err)
	}
}
