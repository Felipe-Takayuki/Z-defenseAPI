package main

import (
	"github.com/Felipe-Takayuki/sistema-de-defesa-dsinCC/config"
	"github.com/Felipe-Takayuki/sistema-de-defesa-dsinCC/routes"
)

func main() {
	 config.Init()
	
	routes.Init()
}
