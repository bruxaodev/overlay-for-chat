package main

import (
	"chat/config"
	"embed"
	"log"

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
	app.NewWebviewWindowWithOptions(application.WebviewWindowOptions{
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
		KeyBindings: map[string]func(window *application.WebviewWindow){
			"CmdOrCtrl+Shift+X": func(window *application.WebviewWindow) {
				window.SetFrameless(!frameless)
				frameless = !frameless
				x, y := window.Position()
				w, h := window.Size()
				config.SaveConfig(configPath, &config.Config{
					Frameless: frameless,
					X:         x,
					Y:         y,
					Width:     w,
					Height:    h,
					Link:      conf.Link,
				})
			},
		},
	})

	// Create a goroutine that emits an event containing the current time every second.
	// The frontend can listen to this event and update the UI accordingly.
	// go func() {
	// 	for {
	// 		// now := time.Now().Format(time.RFC1123)
	// 		// app.Events.Emit(&application.WailsEvent{
	// 		// 	Name: "time",
	// 		// 	Data: now,
	// 		// })
	// 		time.Sleep(time.Second)

	// 	}
	// }()

	// Run the application. This blocks until the application has been exited.
	err := app.Run()

	// If an error occurred while running the application, log it and exit.
	if err != nil {
		log.Fatal(err)
	}
}
