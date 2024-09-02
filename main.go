package main

import (
	"fmt"
	"os"
	"time"

	cmd "api-product/cmd"
)

const banner = `
    .___                            __     __   .__.__          __   
  __| _/____   _____ ______   _____/  |_  |  | _|__|  | _____ _/  |_ 
 / __ |/  _ \ /     \\____ \_/ __ \   __\ |  |/ |  |  | \__  \\   __\
/ /_/ (  <_> |  Y Y  |  |_> \  ___/|  |   |    <|  |  |__/ __ \|  |  
\____ |\____/|__|_|  |   __/ \___  |__|   |__|_ |__|____(____  |__|  
     \/            \/|__|        \/            \/            \/      

`

func main() {
	if tz := os.Getenv("TZ"); tz != "" {
		var err error
		time.Local, err = time.LoadLocation(tz)
		if err != nil {
			fmt.Printf("error loading location '%s': %v\n", tz, err)
		} else {
			fmt.Printf("location loaded '%s'\n", tz)
		}
	}

	fmt.Print(banner)
	cmd.Run()
}
