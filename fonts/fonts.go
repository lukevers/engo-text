package fonts

import (
	"github.com/golang/freetype/truetype"
	"io/ioutil"
	"log"
	"path/filepath"
)

var fonts map[string]*truetype.Font = make(map[string]*truetype.Font)

type Font string

const FONT_PRIMARY Font = "SourceSansPro-Regular.ttf"

func init() {
	files, err := filepath.Glob("resources/fonts/*.ttf")
	if err != nil {
		log.Fatal("Could not load font files", err)
	}

	for _, f := range files {
		_, name := filepath.Split(f)
		data, err := ioutil.ReadFile(f)
		if err != nil {
			log.Fatal("Could not read font file", err)
		}

		fonts[name], err = truetype.Parse(data)
		if err != nil {
			log.Fatal("Could not parse file as ttf", err)
		}
	}
}

func Get(font Font) *truetype.Font {
	return fonts[string(font)]
}
