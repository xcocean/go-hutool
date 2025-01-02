package gg

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

// zip 将指定路径下的文件或目录打包成 ZIP 文件，保存到当前目录
// println(Zip_zip("d:\\gg")) --> d:\gg.zip
func Zip_zip(srcPath string) string {
	// 获取当前工作目录
	/*currentDir, err := os.Getwd()
	if err != nil {
		panic(fmt.Sprintf("无法获取当前工作目录: %v", err))
	}*/

	parent := File_getParent(srcPath)

	// 创建 ZIP 文件
	zipFilePath := filepath.Join(parent, filepath.Base(srcPath)+".zip")
	zipFile, err := os.Create(zipFilePath)
	if err != nil {
		panic(fmt.Sprintf("无法创建 ZIP 文件: %v", err))
	}
	defer zipFile.Close()

	// 创建 ZIP 写入器
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// 遍历源路径下的所有文件
	err = filepath.Walk(srcPath, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 创建 ZIP 文件中的文件头
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		// 设置文件头中的文件名（相对路径）
		relPath, err := filepath.Rel(srcPath, filePath)
		if err != nil {
			return err
		}
		header.Name = relPath

		// 如果是目录，添加目录条目
		if info.IsDir() {
			header.Name += "/"
			_, err = zipWriter.CreateHeader(header)
			return err
		}

		// 如果是文件，添加文件条目并写入内容
		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}
		file, err := os.Open(filePath)
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = io.Copy(writer, file)
		return err
	})

	if err != nil {
		panic(fmt.Sprintf("无法打包文件: %v", err))
	}

	// fmt.Printf("文件已打包到: %s\n", zipFilePath)
	return zipFilePath
}

// 在现有的 ZIP 文件中添加新文件或目录
func File_appendToZip(zipPath, appendFilePath string) {
	if !File_exist(appendFilePath) {
		panic("需要追加的文件不存在：" + appendFilePath)
	}
	if !File_exist(zipPath) {
		panic("zip压缩文件不存在：" + zipPath)
	}
	temp := File_getParent(zipPath)
	prefix := File_getPrefix(zipPath)
	temp = temp + string(os.PathSeparator) + prefix + "_" + Str_toString(time.Now().UnixMilli())
	// 解压到临时目录
	Zip_unzip(zipPath, temp)
	// 复制
	File_copy(appendFilePath, temp, true)
	Zip_zip(temp)
	File_delete(temp)
	// 删除原文件
	File_delete(zipPath)
	// 重命名
	File_rename(temp+".zip", File_getName(zipPath), true)
}

// unzip 解压 ZIP 文件到指定目录
// Zip_unzip("d:\\gg\\gg.zip", "d:\\gg\\ex1")
func Zip_unzip(zipFile, outDir string) {
	// 打开 ZIP 文件
	reader, err := zip.OpenReader(zipFile)
	if err != nil {
		panic(fmt.Sprintf("无法打开 ZIP 文件: %v", err))
	}
	defer reader.Close()

	// 遍历 ZIP 文件中的每个文件
	for _, file := range reader.File {
		// 构建解压后的文件路径
		filePath := filepath.Join(outDir, file.Name)

		// 如果是目录，创建目录
		if file.FileInfo().IsDir() {
			if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
				panic(fmt.Sprintf("无法创建目录: %v", err))
			}
			continue
		}

		// 创建文件所在的目录
		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			panic(fmt.Sprintf("无法创建目录: %v", err))
		}

		// 打开 ZIP 文件中的文件
		zipFile, err := file.Open()
		if err != nil {
			panic(fmt.Sprintf("无法打开 ZIP 文件中的文件: %v", err))
		}
		defer zipFile.Close()

		// 创建目标文件
		outFile, err := os.Create(filePath)
		if err != nil {
			panic(fmt.Sprintf("无法创建文件: %v", err))
		}
		defer outFile.Close()

		// 将 ZIP 文件中的内容复制到目标文件
		if _, err := io.Copy(outFile, zipFile); err != nil {
			panic(fmt.Sprintf("无法解压文件: %v", err))
		}
	}
}
