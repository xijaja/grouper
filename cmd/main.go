package main

import (
	g "github.com/AllenDang/giu"
	"grouper/common/aui"
)

// 程序入口
func main() {
	// GUI  // g.MasterWindowFlagsNotResizable  // MasterWindowFlagsMaximized
	wnd := g.NewMasterWindow("Grouper 🐟", 730, 600, g.MasterWindowFlagsNotResizable)
	// wnd.SetDropCallback(onDrop)
	wnd.Run(aui.Loop)
}
