package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	// 指定要列出的目录
	root := "C:/Users/25837/Desktop/tmp/tiktok"

	// 读取目录下的所有文件和子目录
	files, err := ioutil.ReadDir(root)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 遍历文件列表并输出文件名或子目录名
	for _, file := range files {
		if file.IsDir() {
			// 如果是子目录，继续递归列出子目录中的文件和子目录
			subdir := filepath.Join(root, file.Name())
			fmt.Println("Directory:", subdir)
			listFiles(subdir)
		} else {
			// 如果是文件，输出文件名
			fmt.Println("File:", file.Name())
		}
	}
}

// 递归列出子目录中的文件和子目录
func listFiles(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, file := range files {
		if file.IsDir() {
			subdir := filepath.Join(path, file.Name())
			fmt.Println("Directory:", subdir)
			listFiles(subdir)
		} else {
			fmt.Println("File:", file.Name())
		}
	}
}
