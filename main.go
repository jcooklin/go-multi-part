package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "go-multi-part"
	app.Usage = "make a multi part file"
	app.Flags = []cli.Flag{
		cli.StringSliceFlag{
			Name:  "file",
			Value: &cli.StringSlice{},
		},
	}
	app.Action = createMultipPart

	app.Run(os.Args)
}

func createMultipPart(c *cli.Context) {
	buffer := bytes.NewBuffer([]byte{})
	mpw := multipart.NewWriter(buffer)

	for _, file := range c.StringSlice("file") {
		w, err := mpw.CreateFormFile("pulse-plugin", file)
		if err != nil {
			panic(err)
		}

		fb, err := ioutil.ReadFile(file)
		if err != nil {
			panic(err)
		}
		w.Write(fb)
	}

	err := mpw.Close()
	if err != nil {
		panic(err)
	}

	//print out the multiart file
	fmt.Println(buffer.String())
}
