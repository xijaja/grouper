package main

import (
	g "github.com/AllenDang/giu"
	"grouper/aui"
)

// 程序入口
func main() {
	// 启动
	// app.Uper()
	// g.MasterWindowFlagsNotResizable  // MasterWindowFlagsMaximized
	wnd := g.NewMasterWindow("Grouper 🐟", 730, 600, g.MasterWindowFlagsNotResizable)
	// wnd.SetDropCallback(onDrop)
	go aui.Prg()
	wnd.Run(aui.Loop)
}
