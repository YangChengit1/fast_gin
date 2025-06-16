package md5

import (
	"crypto/md5"
	"encoding/hex"
	"io"
)

// MD5WithFile 计算hash值
func MD5WithFile(file io.Reader) string {
	// 创建一个新的 MD5 哈希计算器
	m := md5.New()
	// 将文件内容（io.Reader）的数据流写入哈希计算器
	io.Copy(m, file)
	// 计算最终的哈希值（字节切片形式）
	sum := m.Sum(nil)
	// 将字节切片转换为十六进制字符串并返回
	return hex.EncodeToString(sum)
}
