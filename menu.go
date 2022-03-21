package main

import "fmt"
 
func main() {
	var cmd string
	for {
		fmt.Println("============================")
		fmt.Println("h：hello world")
		fmt.Println("go：go")
		fmt.Println("vs：vscode")
		fmt.Println("q：exit")
		fmt.Print("请输入命令:")
		fmt.Scanln(&cmd)
		fmt.Println("============================")
		switch cmd {
		case "h":
			fmt.Println("hello world!")
		case "go":
			fmt.Println("go!")
		case "vs":
			fmt.Println("vscode!")
		case "q":
			return
		default:
			fmt.Println("请输入正确的命令！")
		}
 
	}
}