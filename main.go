package main

import (
	"retroeditor/parser"
)

func main() {
	p := parser.NewGameListParser("/home/sen/tmp/gamelist.xml")
	err := p.Parse()
	p.Remove("Ball Breakers")
	p.Save()
	if err != nil {
		panic(err)
	}
	// widgets.NewQApplication(len(os.Args), os.Args)
	// window := ui.NewRetroUI()
	// window.Show()
	// widgets.QApplication_Exec()
}
