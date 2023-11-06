// 验证逻辑
package filehashchecking

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"os"
	"path/filepath"
)

// 计算Hash
func CountHash(data []byte) (hashString string) {
	hashByte := sha256.Sum256(data)
	hashString = hex.EncodeToString(hashByte[:])
	return hashString
}

// 检查Hash是否真确
func CheckingHash() (success []FileInfo, fail []FileInfo) {
	// 保存解析sha256.json的数据
	shaSaveData := []FileInfo{}

	// 读取Hash文件
	data, err := os.ReadFile(filepath.Join(ScanDirPath, "sha256.json"))
	ErrprDisplay(err)

	// 开始解析
	err = json.Unmarshal(data, &shaSaveData)
	ErrprDisplay(err)

	// 开始验证
	success, fail = shaVerify(shaSaveData)

	return success, fail
}

// 从sha256.json验证是否被篡改
func shaVerify(shaSaveData []FileInfo) (success []FileInfo, fail []FileInfo) {
	for _, v := range shaSaveData {
		data, err := os.ReadFile(v.Path)
		if err != nil {
			v.Path = err.Error()
			fail = append(fail, v)
		}

		if v.Hash == CountHash(data) {
			success = append(success, v)
		} else {
			// 无法同时读取和写入sha256.json文件Hash值，允许不同
			if v.Name != "sha256.json" {
				fail = append(fail, v)
			}
		}
	}

	return success, fail
}
