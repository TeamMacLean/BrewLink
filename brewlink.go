package main

import (
	"os"
	"github.com/codegangsta/cli"
	"strings"
	"encoding/json"
	"io/ioutil"
	"path"
	"errors"
	"log"
)

var (
	config Config
)

type Config struct {
	CellarPath   string `json:"CellarPath"`
	SoftwarePath string `json:"SoftwarePath"`
}

func main() {
	//new cli app
	app := cli.NewApp()
	app.Name = "BrewLink"
	app.Usage = "Link software installed with brew to somewhere else"
	app.Action = func(c *cli.Context) {

		err := loadConfig()
		check(err)

		//check that there is just one argument
		if (len(c.Args()) == 1) {
			magic(c.Args()[0], *c)
			//println("Hello", c.Args()[0])
		} else {
			//show user the way to use the app
			cli.ShowAppHelp(c)
		}

	}

	//process the above
	app.Run(os.Args)
}

func loadConfig() error {

	//read config
	dat, err := ioutil.ReadFile("./.brewlink.json")
	check(err)

	//create empty Config struct
	config = Config{}

	//unmarshal config file to Config struct
	err = json.Unmarshal(dat, &config)
	check(err)

	//both paths should exists to begin with
	sExists, sError := exists(config.SoftwarePath)
	check(sError)
	cExists, cError := exists(config.CellarPath)
	check(cError)

	if (!sExists) {
		return errors.New("it seems that the SoftwarePath in your config does not exist, please check .brewlink.json")
	}
	if (!cExists) {
		return errors.New("it seems that the CellarPath in your config does not exist, please check .brewlink.json")

	}

	return nil

}

//function taken from http://stackoverflow.com/questions/10510691/how-to-check-whether-a-file-or-directory-denoted-by-a-path-exists-in-golang
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func magic(a string, c cli.Context) {

	//split into chunks at '-'
	splitString := strings.Split(a, "-")

	if (len(splitString) < 2) {
		//too short
		println("Error: There should only be one arugument")
		cli.ShowAppHelp(&c)
	} else if (len(splitString) > 2) {
		//too long
		println("Error: There should only be one arugument")
		cli.ShowAppHelp(&c)
	} else {
		//just right
		toolName := splitString[0]
		toolVersion := splitString[1]

		//build old path
		oldPath := path.Join(config.CellarPath, toolName, toolVersion)

		//folder above x86_64
		newPathPreX86 := path.Join(config.SoftwarePath, toolName, toolVersion)

		//make link parent folder (mkdir -p)
		err := os.MkdirAll(newPathPreX86, 0755)
		check(err)

		//path to target
		symLinkTarget := path.Join(config.SoftwarePath, toolName, toolVersion, "x86_64")

		//create sym link
		//err = os.Symlink(oldPath, symLinkTarget)
		//check(err)

		println("linking %v to %v", oldPath, symLinkTarget)
	}
}

func check(e error) {
	if (e != nil) {
		log.Fatal(e)
	}
}