package tool

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strings"
)

// ---------------------------------------------
// 文件遍历
// ---------------------------------------------

// GetFileList 获取文件列表
func GetFileList(path string, up func(newPath string)) {
	// 获取路径
	fs, _ := ioutil.ReadDir(path)
	// 固定的path
	if path[len(path)-1:] != "/" {
		path += "/"
	}
	// 执行上传
	for _, file := range fs {
		if file.IsDir() {
			// 遇到文件夹时就开启一个并发递归
			go GetFileList(path+file.Name()+"/", up)
		} else {
			newPath := path + file.Name()
			up(newPath) // 调用函数参数
		}
	}
}

// GetFilesAndDirs 获取指定目录下的所有文件和目录
// func GetFilesAndDirs(dirPth string) (files []string, dirs []string, err error) {
// 	err = filepath.Walk(dirPth, func(path string, info os.FileInfo, err error) error {
// 		files = append(files, path)
// 		return nil
// 	})
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	for _, file := range files {
// 		dirs = append(dirs, file[len(dirPth):])
// 	}
// 	return files, dirs, err
// }

// ---------------------------------------------
// 退出程序
// ---------------------------------------------

func GoodBye() {
	fmt.Println("⚠ 按回车或回复任意，退出程序。")
	reader := bufio.NewReader(os.Stdin) // 读取命令行
	osWin := isOsWindows()              // 当前系统是否为windows
	_ = readInput(reader, osWin)
	// 按 CTRL+C 或输入 exit 以退出程序
	// t := strings.Split(text, " ")
	// if len(t) == 1 && strings.Compare("exit", text) == 0 {
	// 	fmt.Sprintln("Bye~ Bye~")
	// 	os.Exit(1)
	// }
	if true {
		fmt.Sprintln("Bye~ Bye~")
		os.Exit(1)
	}
}

// 获取当前计算机系统类型，是否为Windows
func isOsWindows() bool {
	// runtime.GOARCH 返回当前的系统架构；runtime.GOOS 返回当前的操作系统。
	sysType := runtime.GOOS
	// fmt.Println(fmt.Sprintf("您的系统是%v，采用%v架构", runtime.GOOS, runtime.GOARCH))
	if sysType == "linux" {
		// LINUX系统
		// fmt.Println("Linux system")
		return false
	}
	if sysType == "windows" {
		// windows系统
		// fmt.Println("Windows system")
		return true
	}
	return false
}

// 读取用户输入
func readInput(reader *bufio.Reader, osWin bool) (text string) {
	if osWin {
		text, _ = reader.ReadString('\n')
		text = strings.Replace(text, "\r\n", "", -1)
	} else {
		text, _ = reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
	}
	return text
}
