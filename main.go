package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"golang.org/x/net/websocket"
)

func main() {
	e := echo.New()
	// entry point for SPA react app served by echo
	e.Static("/", "public")
	e.GET("/", func(c echo.Context) error {
		c.File("public/index.html")
		return c.File("public/index.html")
	})
	// reactive webpage nonsense goes here
	e.GET("/ws", func(c echo.Context) error {
		websocket.Handler(func(ws *websocket.Conn) {
			defer ws.Close()
			for {
				// Write
				err := websocket.Message.Send(ws, "document.getElementById('1').style.backgroundColor = 'red'")
				if err != nil {
					c.Logger().Error(err)
					ws.Close()
				}

				// Read
				msg := ""
				err = websocket.Message.Receive(ws, &msg)
				if err != nil {
					c.Logger().Error(err)
				}
				fmt.Printf("%s\n", msg)
			}
		}).ServeHTTP(c.Response(), c.Request())
		return nil
	})

	e.Logger.Fatal(e.Start(":1234"))
}
