package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/zignig/viewer/util"
)

var conf util.Config

func main() {
	fmt.Println("Running Hub Server")
	conf := util.GetConfig("universe.toml")
	u := AndLetThereBeLight(conf)
	u.conf = conf
	fmt.Println(u)
	go h.run()
	go u.world.Run()

	r := gin.Default()
	r.LoadHTMLFiles("index.html")
	r.Static("static", "static")
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
	r.GET("/ws", func(c *gin.Context) {
		u.wshandler(c.Writer, c.Request)
	})
	r.GET("/asset/*path", u.asset)
	r.Run(":8090")
}

func (u *universe) asset(c *gin.Context) {
	// send to asset manager
	path := c.Params.ByName("path")
	data, err := u.cache.Cat(path)
	if err != nil {
		c.String(500, err.Error())
	}
	c.Data(200, "", data)
}

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (u *universe) wshandler(w http.ResponseWriter, r *http.Request) {
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
	}
	c := &connection{send: make(chan []byte, 256), ws: conn}
	h.register <- c
	//world.register <- c
	// todo  , move this to write pump and push a new player
	go c.writePump()
	c.readPump()
}
