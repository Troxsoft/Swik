package pkg

import (
	"errors"
	"os"

	"github.com/jchv/go-webview2"
)

type WindowConfig struct {
	Width  uint
	Height uint
	Title  string
	Focus  bool
	Debug  bool
	Center bool
	Resize bool
}
type Window struct {
	window webview2.WebView
	config WindowConfig
}

var (
	WindowOnCreatedErr = errors.New("Error on the create window")
)

type BindValue struct {
	Value any `json:""`
}

func ReturnValue(data any) BindValue {
	return BindValue{
		Value: data,
	}
}
func (w *Window) SetHtml(html string) {
	w.window.SetHtml(html)
}
func (w *Window) Bind(name string, function func(data string) BindValue) error {

	return w.window.Bind(name, function)
}
func (w *Window) SetHtmlFromFile(file string) error {

	//fmt.Printf("%s\n", wd+"\\"+file)
	fi, err := os.ReadFile(file)

	if err != nil {
		return err
	}
	src := string(fi)
	w.SetHtml(src)
	return nil

}
func (w *Window) SetTitle(title string) {
	w.window.SetTitle(title)
	w.config.Title = title
}
func (w *Window) Title() string {
	return w.config.Title
}
func (window *Window) SetSize(w, h uint) {

	window.config.Width = w
	window.config.Height = h
	if window.config.Resize {

		window.window.SetSize(int(window.config.Width), int(window.config.Height), webview2.HintNone)
	} else {
		window.window.SetSize(int(window.config.Width), int(window.config.Height), webview2.HintFixed)
	}
}

func (w *Window) Destroy() {
	w.window.Destroy()
}
func (w *Window) Run() {
	w.window.Run()
}
func NewWindow(config WindowConfig) (*Window, error) {
	w := webview2.NewWithOptions(webview2.WebViewOptions{
		Debug:     config.Debug,
		AutoFocus: config.Focus,
		WindowOptions: webview2.WindowOptions{
			Title:  config.Title,
			Width:  config.Width,
			Height: config.Height,
			Center: config.Center,
		},
	})
	if w == nil {
		return nil, WindowOnCreatedErr
	}
	if config.Resize {

		w.SetSize(int(config.Width), int(config.Height), webview2.HintNone)
	} else {
		w.SetSize(int(config.Width), int(config.Height), webview2.HintFixed)
	}
	return &Window{
		window: w,
		config: config,
	}, nil
}
func NewConfig(title string, w, h uint) WindowConfig {
	return WindowConfig{
		Width:  w,
		Height: h,
		Title:  title,
		Focus:  true,
		Debug:  true,
		Center: true,
		Resize: true,
	}
}
