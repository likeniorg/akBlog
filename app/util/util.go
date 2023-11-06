package util

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var errorMessage = "内部错误，请查看日志"

// 执行命令
func Shell(cmd string) {
	tmp := exec.Command("bash", "-c", cmd)
	tmp.Stdout = os.Stdout
	tmp.Stderr = os.Stderr
	tmp.Run()
}

// 错误显示函数
func ErrprDisplay(err interface{}) {
	if err != nil {
		fmt.Println(err)
		return
	}
}

// 替换文件内字符串
func ReplaceFileString(filePath, old, new string) error {

	data, err := os.ReadFile(filePath)
	ErrprDisplay(err)

	newString := strings.Replace(string(data), old, new, 1)

	err = os.WriteFile(filePath, []byte(newString), 0600)
	ErrprDisplay(err)

	return err
}
