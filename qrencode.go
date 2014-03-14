package main

import (
	"github.com/docopt/docopt.go"
	"github.com/qpliu/qrencode-go/qrencode"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"strconv"
)

var arguments map[string]interface{}

func init() {
	var err error

	arguments, err = docopt.Parse(`qrencode v1.0
Copyright (C) 2014 Calin Culianu <calin.culianu@gmail.com> 1Ca1inQuedcKdyELCTmN8AtKTTehebY4mC
	
Usage:
  qrencode [-jp] [-b GRID_BLOCKSIZE_PX] [-m MARGIN_PX] <string> <outfile>

Parameters:
  <string>               The string to encode
  <outfile>              The output filename (PNG default, or JPEG if -j option)
 
Options:
  -j                     Output image in JPEG format
  -p                     Output image in PNG format (default)
  -m MARGIN_PX           The size, in pixels, of the white border [default 4]
  -b GRID_BLOCKSIZE_PX   The size, in pixels, of each qr code grid element [default 1]
`, nil, true, "qrencode", false)
	if err != nil {
		log.Fatal(err)
	}

}

func main() {
	var err error

	s := arguments["<string>"].(string)
	o := arguments["<outfile>"].(string)

	bs, m := 10, 4

	if arguments["-b"] != nil {
		bs, err = strconv.Atoi(arguments["-b"].(string))

		if err != nil {
			log.Fatal(err)
		}
	}
	if arguments["-m"] != nil {
		m, err = strconv.Atoi(arguments["-m"].(string))

		if err != nil {
			log.Fatal(err)
		}
	}

	var grid *qrencode.BitGrid

	grid, err = qrencode.Encode(s, qrencode.ECLevelM)

	if err != nil {
		log.Fatal(err)
	}

	img := grid.ImageWithMargin(bs, m)

	outfile, err := os.Create(o)

	if err != nil {
		log.Fatal(err)
	}

	if arguments["-j"] != nil {
		err = jpeg.Encode(outfile, img, nil)
	} else {
		err = png.Encode(outfile, img)
	}

	if err != nil {
		log.Fatal(err)
	}

}
