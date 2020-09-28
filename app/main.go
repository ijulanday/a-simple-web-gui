package main

import (
	"net/http"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

type hello struct {
    app.Compo
    name string
}


func (h *hello) Render() app.UI {
    return app.Div().Body(
        app.Main().Body(
            app.H1().Body(
                app.Text("Hello, "),
                app.If(h.name != "",
                    app.Text(h.name),
                ).Else(
                    app.Text("World"),
                ),
            ),
            app.Input().
                Value(h.name).
                Placeholder("What is your name?").
                AutoFocus(true).
                OnChange(h.OnInputChange),
        ),
    )
}

func (h *hello) OnInputChange(ctx app.Context, e app.Event) {
    h.name = ctx.JSSrc.Get("value").String()
    h.Update()
}

func main() {
	h := &app.Handler{
        Title:  "Hello Demo",
        Author: "Ian Ulanday",
    }

    if err := http.ListenAndServe(":80", h); err != nil {
        panic(err)
    }

    app.Route("/", &hello{})
    app.Route("/hello", &hello{})
    app.Run()
}
