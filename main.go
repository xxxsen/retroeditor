package main

import (
	"flag"
	"os"
	"retroeditor/ui"

	"github.com/therecipe/qt/widgets"
)

var root = flag.String("root", "/home/sen/tmp/roms", "roms root path")

func main() {
	flag.Parse()
	widgets.NewQApplication(len(os.Args), os.Args)
	window := ui.NewRetroUI(*root)
	window.Show()
	widgets.QApplication_Exec()
}
