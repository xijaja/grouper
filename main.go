package main

import (
	g "github.com/AllenDang/giu"
	"grouper/aui"
)

// 程序入口
func main() {
	// 启动
	// app.Uper()
	// g.MasterWindowFlagsNotResizable
	wnd := g.NewMasterWindow("Grouper 🐟", 800, 600, g.MasterWindowFlagsNotResizable)
	// wnd.SetDropCallback(onDrop)
	wnd.Run(aui.Loop)
}
