package main

import (
	"github.com/thomassilvi/Meuh"
	"github.com/thomassilvi/Meuh/controllers"
	"log"
)

//-------------------------------------------------------------------------------------------------

//#################################################################################################

func main() {
        log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

        meuh.Init("default.conf")
        controllers.InitMainTemplates()
        controllers.InitUserTemplates()
        controllers.InitUserRoutes()
        meuh.InitDefault()

        meuh.Run()
}

//#################################################################################################
