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
var Enable = true

func main() {
	conf, _err := config.LoadConfig(configPath)
	if _err != nil {
		log.Fatal(_err)
	}
	frameless = conf.Frameless

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

	overlay := app.NewWebviewWindowWithOptions(application.WebviewWindowOptions{
		Title: "Chat",
		Mac: application.MacWindow{
			InvisibleTitleBarHeight: 50,
			Backdrop:                application.MacBackdropTranslucent,
			TitleBar:                application.MacTitleBarHiddenInset,
		},
		BackgroundType:    application.BackgroundTypeTransparent,
		AlwaysOnTop:       true,
		DisableResize:     false,
		Width:             conf.Width,
		Height:            conf.Height,
		Frameless:         conf.Frameless,
		URL:               conf.Link,
		X:                 conf.X,
		Y:                 conf.Y,
		IgnoreMouseEvents: conf.Frameless,
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
				overlay.SetFrameless(frameless)
				overlay.SetIgnoreMouseEvents(frameless)
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

		}
	}()

	// Run the application. This blocks until the application has been exited.
	err := app.Run()

	// If an error occurred while running the application, log it and exit.
	if err != nil {
		log.Fatal(err)
	}
}
