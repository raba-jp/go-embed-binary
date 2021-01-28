package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	_ "github.com/raba-jp/go-embed-binary/statik"
	"github.com/rakyll/statik/fs"
)

func main() {
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	fs, err := statikFS.Open("/ffmpeg")
	if err != nil {
		log.Fatal(err)
	}
	r, err := ioutil.ReadAll(fs)
	if err != nil {
		log.Fatal(err)
	}
	wd, _ := os.Getwd()
	binpath := filepath.Join(wd, "ffmpeg")
	if err := ioutil.WriteFile(binpath, r, 0o755); err != nil {
		log.Fatal(err)
	}

	out, err := exec.Command(binpath).CombinedOutput()
	fmt.Printf("%s\n", string(out))
	if err != nil {
		log.Fatal(err)
	}
}
