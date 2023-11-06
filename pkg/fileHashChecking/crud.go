// Create-Read-Update-Delete
package filehashchecking

import (
	"encoding/json"
	"fmt"
	"os"
)

// 定义文件信息三要素
type FileInfo struct {
	Name string
	Path string
	Hash string
}

// 默认检测路径
var ScanDirPath = "web/files/"

// Hash文件保存路径
var sha256Path = ScanDirPath + "sha256.json"

// 创建Hash表
func CreateHash() {
	// 检查是否存在sha256.json文件
	if _, err := os.ReadFile(sha256Path); err != nil {
		// 错误的话是因为没有创建files目录
		err := os.Mkdir(ScanDirPath, 0700)
		ErrprDisplay(err)

		// 递归查找本地文件，返回值为全局变量fileInfos
		fileInfos := recursionRerurnFiles(ScanDirPath)

		// 写入json文件
		data, _ := json.MarshalIndent(fileInfos, "", "	")
		os.WriteFile(sha256Path, data, 0600)

	}
}

// 新增Hash
func AddHash(fileName string) {
	// 临时存储变量
	tmpInfo := FileInfo{}

	// 开始赋值
	tmpInfo.Name = fileName
	tmpInfo.Path = ScanDirPath + fileName
	data, err := os.ReadFile(tmpInfo.Path)
	ErrprDisplay(err)

	//计算Hash
	tmpInfo.Hash = CountHash(data)

	// 将临时存储信息添加到全部信息中
	fileInfos := GetHash()
	fileInfos = append(fileInfos, tmpInfo)

	// 转换为json格式数据
	filedata, err := json.MarshalIndent(fileInfos, "", "	")
	ErrprDisplay(err)

	// 写入json格式
	err = os.WriteFile(sha256Path, filedata, 0600)
	ErrprDisplay(err)

}

// 获取Hash
func GetHash() []FileInfo {
	data, err := os.ReadFile(sha256Path)
	ErrprDisplay(err)
	var fileinfos = []FileInfo{}
	err = json.Unmarshal(data, &fileinfos)
	ErrprDisplay(err)
	fmt.Println("Gethash")

	fmt.Println(fileinfos)
	return fileinfos
}

func ErrprDisplay(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}

// 更新信息
func UpdateHash() {

}

// 删除信息
func DeleteHash() {

}

// 递归返回多个文件信息
// 使用完该函数需要初始化 fileInfos=nil
func recursionRerurnFiles(dirName string) (fileInfos []FileInfo) {
	// 临时存储文件信息
	tmpSave := FileInfo{}

	if dirName == "" {
		dirName = "./"
	}

	// 读取文件目录
	dir, err := os.ReadDir(dirName)
	ErrprDisplay(err)

	// 递归开始
	for _, v := range dir {
		// 判断是不是目录
		if v.IsDir() {
			// 是目录递归执行
			fileInfos = append(fileInfos, recursionRerurnFiles(dirName+v.Name()+"/")...)
		} else {
			// 获取文件数据
			data, _ := os.ReadFile(dirName + v.Name())
			ErrprDisplay(err)

			// 保存文件三要素
			tmpSave.Hash = CountHash(data)
			tmpSave.Name = v.Name()
			tmpSave.Path = dirName + v.Name()

			// 将临时存储的信息存入文件信息集合
			fileInfos = append(fileInfos, tmpSave)
		}
	}
	return fileInfos
}
