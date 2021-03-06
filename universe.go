package main

import (
	"fmt"
	"html/template"

	"github.com/zignig/cohort/assets"
	"github.com/zignig/cohort/util"
	"github.com/zignig/cohort/world"
)

type universe struct {
	conf  *util.Config
	world *world.World
	h     *hub
	cache *assets.Cache
	templ *template.Template
	store *Store // local sqlite sotre
}

func AndLetThereBeLight(config *util.Config) *universe {
	fmt.Println("FATOOOOMPSH")
	u := &universe{}
	u.conf = config
	u.cache = assets.NewCache()
	u.store = NewStore("")
	u.world = world.NewWorld(config, u.cache)
	u.h = NewHub(u.world)
	return u
}

func (u *universe) run() {
	go u.world.Run()
	go u.h.run()
}

func (u *universe) String() (s string) {
	return "REALLY BIG"
}
