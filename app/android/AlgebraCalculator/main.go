// +build darwin linux

package main

import (
	"golang.org/x/mobile/app"
	"golang.org/x/mobile/event/key"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
	"golang.org/x/mobile/event/touch"
	"golang.org/x/mobile/gl"
)

func main() {
	app.Main(run)
}

func run(a app.App) {
	var glctx gl.Context
	var sz size.Event
	for e := range a.Events() {
		switch e := a.Filter(e).(type) {
		case lifecycle.Event:
			switch e.Crosses(lifecycle.StageVisible) {
			case lifecycle.CrossOn:
				glctx, _ = e.DrawContext.(gl.Context)
				onStart(glctx)
				a.Send(paint.Event{})
			case lifecycle.CrossOff:
				onStop()
				glctx = nil
			}
		case size.Event:
			sz = e
		case paint.Event:
			if glctx == nil || e.External {
				continue
			}
			onPaint(glctx, sz)
			a.Publish()
			a.Send(paint.Event{})
		case touch.Event:
			if down := e.Type == touch.TypeBegin; down || e.Type == touch.TypeEnd {

			}
		case key.Event:
			if e.Code != key.CodeSpacebar {
				break
			}
			if down := e.Direction == key.DirPress; down || e.Direction == key.DirRelease {

			}
		}

	}
}

func onStart(glctx gl.Context) {

}

func onStop() {

}

func onPaint(glctx gl.Context, sz size.Event) {
	glctx.ClearColor(1, 1, 1, 1)
	glctx.Clear(gl.COLOR_BUFFER_BIT)
}
