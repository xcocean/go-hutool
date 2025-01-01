package gg

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

// IsWindows 判断当前是否为Windows环境
func IsWindows() bool {
	return runtime.GOOS == "windows"
}

// 列出指定路径下的目录和文件
func File_ls(path string) []string {
	var dirs []string
	var files []string

	// 读取目录内容
	entries, err := os.ReadDir(path)
	if err != nil {
		if os.IsNotExist(err) {
			return []string{}
		}
		panic("无法读取目录: " + err.Error())
	}

	// 遍历目录内容
	for _, entry := range entries {
		fullPath := filepath.Join(path, entry.Name())
		if entry.IsDir() {
			dirs = append(dirs, fullPath)
		} else {
			files = append(files, fullPath)
		}
	}
	// 合并目录和文件列表
	return append(dirs, files...)
}

// 判断文件或目录是否为空
// 文件是否为空
// 目录：里面没有文件时为空 文件：文件大小为0时为空
func File_isEmpty(path string) bool {
	// 获取文件或目录信息
	fileInfo, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return true
		}
		panic("无法获取路径信息: " + err.Error())
	}

	// 如果是文件，检查大小是否为0
	if !fileInfo.IsDir() {
		return fileInfo.Size() == 0
	}

	// 如果是目录，检查是否包含文件或子目录
	entries, err := os.ReadDir(path)
	if err != nil {
		panic("无法读取目录内容: " + err.Error())
	}

	return len(entries) == 0
}

// 判断文件或目录是否不为空
func File_isNotEmpty(path string) bool {
	return !File_isEmpty(path)
}

// 判断目录是否为空，异常时 panic
func File_isDirEmpty(dirPath string) bool {
	// 读取目录内容
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		if os.IsNotExist(err) {
			return true
		}
		panic(fmt.Sprintf("无法读取目录内容: %v", err))
	}

	// 如果目录中没有条目，则为空
	return len(entries) == 0
}

// 递归遍历目录及其子目录中的所有文件
// 如果 path 是文件，直接返回过滤结果
func File_loopFiles(path string) []string {
	var files []string

	// 获取路径信息
	fileInfo, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return []string{}
		}
		panic(fmt.Sprintf("无法获取路径信息: %v", err))
	}

	// 如果是文件，直接返回
	if !fileInfo.IsDir() {
		return []string{path}
	}

	// 如果是目录，递归遍历
	err = filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			panic(fmt.Sprintf("访问路径失败: %v", err))
		}

		// 如果是文件，添加到结果中
		if !info.IsDir() {
			files = append(files, filePath)
		}

		return nil
	})

	if err != nil {
		panic(fmt.Sprintf("遍历目录失败: %v", err))
	}

	return files
}

// 创建多层目录并生成文件，元素名（多层目录名）
// 参数 names 是可变参数，表示目录和文件名
// file := File_file("d:\\a", "sub_dir", "example.txt")
func File_file(names ...string) *os.File {
	if len(names) == 0 {
		panic("至少提供一个文件名")
	}

	// 拼接完整路径
	path := filepath.Join(names...)

	// 创建目录（如果目录不存在）
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		panic(fmt.Sprintf("无法创建目录: %v", err))
	}

	// 创建文件
	file, err := os.Create(path)
	if err != nil {
		panic(fmt.Sprintf("无法创建文件: %v", err))
	}

	return file
}

// 创建文件，不创建目录
// 文件路径
func File_newFile(path string) *os.File {
	// 创建目录（如果目录不存在）
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		panic(fmt.Sprintf("无法创建目录: %v", err))
	}

	// 创建文件
	file, err := os.Create(path)
	if err != nil {
		panic(fmt.Sprintf("无法创建文件: %v", err))
	}

	return file
}

// 获取系统临时目录路径
// C:\Users\ADMINI~1\AppData\Local\Temp
func File_getTmpDirPath() string {
	return os.TempDir()
}

// getTmpDir 获取系统的临时文件目录
// C:\Users\ADMINI~1\AppData\Local\Temp
func File_getTmpDir() *os.File {
	// 获取临时文件目录
	dir := os.TempDir()
	if dir == "" {
		panic("无法获取临时文件目录")
	}
	open, _ := os.Open(dir)
	return open
}

// 获取当前用户的主目录
// C:\Users\Administrator
func File_getUserDir() string {
	// 获取当前用户
	currentUser, err := user.Current()
	if err != nil {
		panic(fmt.Sprintf("无法获取当前用户信息: %v", err))
	}

	// 获取用户主目录
	homeDir := currentUser.HomeDir
	if homeDir == "" {
		panic("无法获取用户主目录")
	}

	return homeDir
}

// 判断文件是否存在
// 如果 path 为空，返回 false
func File_exist(path string) bool {
	// 如果路径为空，返回 false
	if path == "" {
		return false
	}

	// 获取文件信息
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		// 其他错误（如权限问题）也返回 false
		return false
	}

	return true
}

// 判断文件是否存在
// 如果 file 为空，返回 false
func File_exist2(file *os.File) bool {
	if file == nil {
		return false
	}
	return File_exist(file.Name())
}

// 获取指定文件的最后修改时间
// println(File_lastModifiedTime("d:\\gg").String()) --> 2023-11-09 16:23:44 +0800 CST
func File_lastModifiedTime(path string) time.Time {
	// 获取文件信息
	fileInfo, err := os.Stat(path)
	if err != nil {
		panic(fmt.Sprintf("无法获取文件信息: %v", err))
	}

	// 返回最后修改时间
	return fileInfo.ModTime()
}

// 计算目录或文件的总大小
// 当给定对象为文件时，直接给出
// 当给定对象为目录时，遍历目录下的所有文件和目录，递归计算其大小，求和返回
func File_size(path string) int64 {
	// 获取文件或目录信息
	fileInfo, err := os.Stat(path)
	if err != nil {
		panic(fmt.Sprintf("无法获取路径信息: %v", err))
	}

	// 如果是文件，直接返回文件大小
	if !fileInfo.IsDir() {
		return fileInfo.Size()
	}

	// 如果是目录，递归计算总大小
	var totalSize int64
	err = filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			panic(fmt.Sprintf("访问路径失败: %v", err))
		}

		// 累加文件大小
		if !info.IsDir() {
			totalSize += info.Size()
		}

		return nil
	})

	if err != nil {
		panic(fmt.Sprintf("遍历目录失败: %v", err))
	}

	return totalSize
}

// 计算文件的总行数
func File_getTotalLines(filePath string) int {
	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		panic(fmt.Sprintf("无法打开文件: %v", err))
	}
	defer file.Close()

	// 使用 bufio.Scanner 逐行读取文件
	scanner := bufio.NewScanner(file)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}

	// 检查扫描过程中是否发生错误
	if err := scanner.Err(); err != nil {
		panic(fmt.Sprintf("读取文件失败: %v", err))
	}

	return lineCount
}

// 判断 filePath 的最后修改时间是否晚于 referencePath 的最后修改时间
func File_newerThan(filePath string, referencePath string) bool {
	// 获取 file 的最后修改时间
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		panic(fmt.Sprintf("无法获取文件信息: %v", err))
	}
	fileModTime := fileInfo.ModTime()

	// 获取 reference 的最后修改时间
	referenceInfo, err := os.Stat(referencePath)
	if err != nil {
		panic(fmt.Sprintf("无法获取参考文件信息: %v", err))
	}
	referenceModTime := referenceInfo.ModTime()

	// 判断 file 的最后修改时间是否晚于 reference 的最后修改时间
	return fileModTime.Before(referenceModTime)
}

// 判断 file 的最后修改时间是否晚于给定的时间戳（毫秒）
func File_newerThan2(filePath string, timeMillis int64) bool {
	// 获取 file 的最后修改时间
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		panic(fmt.Sprintf("无法获取文件信息: %v", err))
	}
	fileModTime := fileInfo.ModTime()

	// 将毫秒时间戳转换为 time.Time
	targetTime := time.UnixMilli(timeMillis)

	// 判断 file 的最后修改时间是否晚于目标时间
	return fileModTime.Before(targetTime)
}

// 创建文件及其父目录，如果文件存在则直接返回
func File_touch(path string) *os.File {
	// 如果文件已经存在，直接打开并返回
	if _, err := os.Stat(path); err == nil {
		file, err := os.OpenFile(path, os.O_RDWR, 0644)
		if err != nil {
			panic(fmt.Sprintf("无法打开文件: %v", err))
		}
		return file
	}

	// 创建父目录
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		panic(fmt.Sprintf("无法创建目录: %v", err))
	}

	// 创建文件
	file, err := os.Create(path)
	if err != nil {
		panic(fmt.Sprintf("无法创建文件: %v", err))
	}

	return file
}

// 创建父文件夹，如果文件夹存在则直接返回
func File_mkParentDirs(path string) {
	// 获取父文件夹路径
	dir := filepath.Dir(path)

	// 如果文件夹已经存在，直接返回
	if _, err := os.Stat(dir); err == nil {
		return
	}

	// 创建父文件夹
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		panic(fmt.Sprintf("无法创建文件夹: %v", err))
	}
}

// 删除文件或文件夹
func File_delete(path string) {
	// 获取文件或文件夹信息
	fileInfo, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
		panic(fmt.Sprintf("无法获取路径信息: %v", err))
	}

	// 如果是文件，直接删除
	if !fileInfo.IsDir() {
		if err := os.Remove(path); err != nil {
			panic(fmt.Sprintf("无法删除文件: %v", err))
		}
		return
	}

	// 删除文件或文件夹
	if err := os.RemoveAll(path); err != nil {
		panic(fmt.Errorf("无法删除路径: %v", err))
	}
}

// 清空文件夹，删除文件夹中的所有文件和子文件夹
func File_clean(dirPath string) bool {
	// 获取文件夹信息
	fileInfo, err := os.Stat(dirPath)
	if err != nil {
		panic(fmt.Sprintf("无法获取路径信息: %v", err))
	}

	// 如果不是文件夹，直接返回 false
	if !fileInfo.IsDir() {
		return false
	}

	// 读取文件夹内容
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		panic(fmt.Sprintf("无法读取文件夹内容: %v", err))
	}

	// 删除文件夹中的所有文件和子文件夹
	for _, entry := range entries {
		fullPath := filepath.Join(dirPath, entry.Name())
		if err := os.RemoveAll(fullPath); err != nil {
			panic(fmt.Sprintf("无法删除路径: %v", err))
		}
	}

	return true
}

// 创建文件夹，如果文件夹存在则直接返回该文件夹的 File 对象
func File_mkdir(dirPath string) *os.File {
	// 如果文件夹已经存在，直接打开并返回
	if _, err := os.Stat(dirPath); err == nil {
		file, err := os.Open(dirPath)
		if err != nil {
			panic(fmt.Sprintf("无法打开文件夹: %v", err))
		}
		return file
	}

	// 创建文件夹
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		panic(fmt.Sprintf("无法创建文件夹: %v", err))
	}

	// 打开文件夹并返回
	file, err := os.Open(dirPath)
	if err != nil {
		panic(fmt.Sprintf("无法打开文件夹: %v", err))
	}

	return file
}

// 在默认临时文件目录下创建临时文件
//
//	file := File_createTempFile()
//	println(file.Name())
//	file.Close()
//	File_delete(file.Name())
func File_createTempFile() *os.File {
	// 在默认临时文件目录下创建临时文件
	file, err := os.CreateTemp("", "gg_temp_*.txt")
	if err != nil {
		panic(fmt.Sprintf("无法创建临时文件: %v", err))
	}

	return file
}

// 在默认临时文件目录下创建临时文件，并指定后缀
func File_createTempFile2(suffix string) *os.File {
	if suffix == "" {
		panic("后缀不能为空")
	}
	// 在默认临时文件目录下创建临时文件
	file, err := os.CreateTemp("", "gg_temp_*."+suffix)
	if err != nil {
		panic(fmt.Sprintf("无法创建临时文件: %v", err))
	}

	return file
}

// 在默认临时文件目录下创建临时文件，并指定前后缀
func File_createTempFile3(prefix, suffix string) *os.File {
	if suffix == "" {
		panic("后缀不能为空")
	}
	if prefix == "" {
		panic("前缀不能为空")
	}
	// 在默认临时文件目录下创建临时文件
	file, err := os.CreateTemp("", prefix+"_*."+suffix)
	if err != nil {
		panic(fmt.Sprintf("无法创建临时文件: %v", err))
	}

	return file
}

// 在指定目录下创建临时文件
func File_createTempFileByDir(dir string) *os.File {
	// 在默认临时文件目录下创建临时文件
	file, err := os.CreateTemp(dir, "gg_temp_*.txt")
	if err != nil {
		panic(fmt.Sprintf("无法创建临时文件: %v", err))
	}

	return file
}

// 在指定目录下创建临时文件，并指定后缀
func File_createTempFileByDir2(dir, suffix string) *os.File {
	if suffix == "" {
		panic("后缀不能为空")
	}
	// 在默认临时文件目录下创建临时文件
	file, err := os.CreateTemp(dir, "gg_temp_*."+suffix)
	if err != nil {
		panic(fmt.Sprintf("无法创建临时文件: %v", err))
	}

	return file
}

// 在指定目录下创建临时文件，并指定前后缀
func File_createTempFileByDir3(dir, prefix, suffix string) *os.File {
	if suffix == "" {
		panic("后缀不能为空")
	}
	if prefix == "" {
		panic("前缀不能为空")
	}
	// 在默认临时文件目录下创建临时文件
	file, err := os.CreateTemp(dir, prefix+"_*."+suffix)
	if err != nil {
		panic(fmt.Sprintf("无法创建临时文件: %v", err))
	}

	return file
}

// 拷贝文件
func File_copyFile(src, dst string) {
	// 打开源文件
	srcFile, err := os.Open(src)
	if err != nil {
		panic(fmt.Sprintf("无法打开源文件: %v", err))
	}
	defer srcFile.Close()

	// 创建目标文件
	dstFile, err := os.Create(dst)
	if err != nil {
		panic(fmt.Sprintf("无法创建目标文件: %v", err))
	}
	defer dstFile.Close()

	// 拷贝文件内容
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		panic(fmt.Sprintf("拷贝文件失败: %v", err))
	}

	// 确保目标文件的内容已刷新到磁盘
	if err := dstFile.Sync(); err != nil {
		panic(fmt.Sprintf("刷新文件失败: %v", err))
	}
}

// copyFile 复制文件
// isOverride 为 false 时，如果目标文件已存在且不允许覆盖，直接返回
func File_copyFile2(src, dst string, isOverride bool) {
	// 如果目标文件已存在且不允许覆盖，直接返回
	if !isOverride {
		if _, err := os.Stat(dst); err == nil {
			return
		}
	}

	// 打开源文件
	srcFile, err := os.Open(src)
	if err != nil {
		panic(fmt.Sprintf("无法打开源文件: %v", err))
	}
	defer srcFile.Close()

	// 创建目标文件
	dstFile, err := os.Create(dst)
	if err != nil {
		panic(fmt.Sprintf("无法创建目标文件: %v", err))
	}
	defer dstFile.Close()

	// 拷贝文件内容
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		panic(fmt.Sprintf("拷贝文件失败: %v", err))
	}

	// 确保目标文件的内容已刷新到磁盘
	if err := dstFile.Sync(); err != nil {
		panic(fmt.Sprintf("刷新文件失败: %v", err))
	}
}

// copyDir 递归复制目录
func File_copyDir(src, dst string, isOverride bool) {
	// 如果目标目录已存在且不允许覆盖，直接返回
	if !isOverride {
		if _, err := os.Stat(dst); err == nil {
			return
		}
	}

	// 创建目标目录
	if err := os.MkdirAll(dst, os.ModePerm); err != nil {
		panic(fmt.Sprintf("无法创建目标目录: %v", err))
	}

	// 读取源目录内容
	entries, err := os.ReadDir(src)
	if err != nil {
		panic(fmt.Sprintf("无法读取源目录内容: %v", err))
	}

	// 递归复制每个条目
	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			File_copyDir(srcPath, dstPath, isOverride)
		} else {
			File_copyFile2(srcPath, dstPath, isOverride)
		}
	}
}

// copy 复制文件或目录到目标路径，支持覆盖选项
// File_copy("d:\\gg", "d:\\gg2", false)
func File_copy(srcPath, destPath string, isOverride bool) {
	// 获取源路径信息
	srcInfo, err := os.Stat(srcPath)
	if err != nil {
		panic(fmt.Sprintf("无法获取源路径信息: %v", err))
	}

	// 如果是文件，复制文件
	if !srcInfo.IsDir() {
		File_copyFile2(srcPath, destPath, isOverride)
		return
	}

	// 如果是目录，递归复制目录
	File_copyDir(srcPath, destPath, isOverride)
}

// 移动文件或目录到目标路径，支持覆盖选项
func File_move(src, target string, isOverride bool) {
	// 获取源路径信息
	srcInfo, err := os.Stat(src)
	if err != nil {
		panic(fmt.Sprintf("无法获取源路径信息: %v", err))
	}

	// 如果目标路径已存在
	if _, err := os.Stat(target); err == nil {
		if !isOverride {
			panic(fmt.Sprintf("目标路径已存在且不允许覆盖: %s", target))
		}

		// 如果目标路径是目录，递归删除
		if err := os.RemoveAll(target); err != nil {
			panic(fmt.Sprintf("无法删除目标路径: %v", err))
		}
	}

	// 移动文件或目录
	if err := os.Rename(src, target); err != nil {
		// 如果跨设备移动失败，使用复制+删除的方式
		if err := moveCrossDevice(src, target, srcInfo); err != nil {
			panic(fmt.Sprintf("无法移动文件或目录: %v", err))
		}
	}
}

// 跨设备移动文件或目录（复制+删除）
func moveCrossDevice(src, target string, srcInfo os.FileInfo) error {
	// 如果是文件，复制文件
	if !srcInfo.IsDir() {
		if err := copyFile(src, target); err != nil {
			return err
		}
		return os.Remove(src)
	}

	// 如果是目录，递归复制目录
	if err := copyDir(src, target); err != nil {
		return err
	}
	return os.RemoveAll(src)
}

// 复制文件
func copyFile(src, dst string) error {
	// 打开源文件
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// 创建目标文件
	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	// 拷贝文件内容
	if _, err := io.Copy(dstFile, srcFile); err != nil {
		return err
	}

	// 确保目标文件的内容已刷新到磁盘
	return dstFile.Sync()
}

// 递归复制目录
func copyDir(src, dst string) error {
	// 创建目标目录
	if err := os.MkdirAll(dst, os.ModePerm); err != nil {
		return err
	}

	// 读取源目录内容
	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	// 递归复制每个条目
	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			if err := copyDir(srcPath, dstPath); err != nil {
				return err
			}
		} else {
			if err := copyFile(srcPath, dstPath); err != nil {
				return err
			}
		}
	}

	return nil
}

// 修改文件或目录的文件名，支持覆盖选项
// File_rename("d:\\gg\\example2.txt", "example2-222.txt", false)
func File_rename(filePath, newName string, isOverride bool) string {
	// 获取文件或目录的父目录
	dir := filepath.Dir(filePath)

	// 构建新路径
	newPath := filepath.Join(dir, newName)

	// 如果目标路径已存在
	if _, err := os.Stat(newPath); err == nil {
		if !isOverride {
			panic(fmt.Sprintf("目标路径已存在且不允许覆盖: %s", newPath))
		}

		// 如果目标路径是目录，递归删除
		if err := os.RemoveAll(newPath); err != nil {
			panic(fmt.Sprintf("无法删除目标路径: %v", err))
		}
	}

	// 重命名文件或目录
	if err := os.Rename(filePath, newPath); err != nil {
		panic(fmt.Sprintf("无法重命名文件或目录: %v", err))
	}

	return newPath
}

// 判断给定路径是否为目录
func File_isDirectory(path string) bool {
	// 获取路径信息
	fileInfo, err := os.Stat(path)
	if err != nil {
		panic(fmt.Sprintf("无法获取路径信息: %v", err))
	}

	// 判断是否为目录
	return fileInfo.IsDir()
}

// 判断给定的 File 对象是否为目录
func File_isDirectory2(file *os.File) bool {
	// 获取文件信息
	fileInfo, err := file.Stat()
	if err != nil {
		panic(fmt.Sprintf("无法获取文件信息: %v", err))
	}

	// 判断是否为目录
	return fileInfo.IsDir()
}

// 判断给定路径是否为文件
func File_isFile(path string) bool {
	// 获取路径信息
	fileInfo, err := os.Stat(path)
	if err != nil {
		panic(fmt.Sprintf("无法获取路径信息: %v", err))
	}

	// 判断是否为文件
	return !fileInfo.IsDir()
}

// 判断给定的 File 对象是否为文件
func File_isFile2(file *os.File) bool {
	// 获取文件信息
	fileInfo, err := file.Stat()
	if err != nil {
		panic(fmt.Sprintf("无法获取文件信息: %v", err))
	}

	// 判断是否为文件
	return !fileInfo.IsDir()
}

// 比较两个文件内容是否相同，首先比较长度，长度一致再比较内容
// println(File_contentEquals("d:\\gg\\example99.txt", "d:\\gg2\\example99.txt"))
func File_contentEquals(file1, file2 string) bool {
	// 获取 file1 的文件信息
	open1, err := os.Open(file1)
	if err != nil {
		panic(fmt.Sprintf("无法获取 file1 的文件信息: %v", err))
	}
	defer open1.Close()

	open2, err := os.Open(file1)
	if err != nil {
		panic(fmt.Sprintf("无法获取 file2 的文件信息: %v", err))
	}
	defer open2.Close()
	return File_contentEquals2(open1, open2)
}

// 比较两个文件内容是否相同，首先比较长度，长度一致再比较内容
func File_contentEquals2(file1, file2 *os.File) bool {
	// 获取 file1 的文件信息
	file1Info, err := file1.Stat()
	if err != nil {
		panic(fmt.Sprintf("无法获取 file1 的文件信息: %v", err))
	}

	// 获取 file2 的文件信息
	file2Info, err := file2.Stat()
	if err != nil {
		panic(fmt.Sprintf("无法获取 file2 的文件信息: %v", err))
	}

	// 比较文件长度
	if file1Info.Size() != file2Info.Size() {
		return false
	}

	// 重置文件指针到开头
	_, err = file1.Seek(0, io.SeekStart)
	if err != nil {
		panic(fmt.Sprintf("无法重置 file1 的文件指针: %v", err))
	}
	_, err = file2.Seek(0, io.SeekStart)
	if err != nil {
		panic(fmt.Sprintf("无法重置 file2 的文件指针: %v", err))
	}

	// 逐字节比较文件内容
	buf1 := make([]byte, 1024)
	buf2 := make([]byte, 1024)
	for {
		n1, err1 := file1.Read(buf1)
		n2, err2 := file2.Read(buf2)

		// 检查读取错误
		if err1 != nil && err1 != io.EOF {
			panic(fmt.Sprintf("读取 file1 失败: %v", err1))
		}
		if err2 != nil && err2 != io.EOF {
			panic(fmt.Sprintf("读取 file2 失败: %v", err2))
		}

		// 比较读取的字节数
		if n1 != n2 {
			return false
		}

		// 比较读取的内容
		for i := 0; i < n1; i++ {
			if buf1[i] != buf2[i] {
				return false
			}
		}

		// 如果到达文件末尾，退出循环
		if err1 == io.EOF || err2 == io.EOF {
			break
		}
	}

	return true
}

func File_isModified(file string, lastModifyTime int64) bool {
	// 获取文件的最后修改时间
	open, err := os.Open(file)
	if err != nil {
		panic(fmt.Sprintf("无法获取文件信息: %v", err))
	}
	defer open.Close()
	return File_isModified2(open, lastModifyTime)

}

// 判断文件是否被修改
func File_isModified2(file *os.File, lastModifyTime int64) bool {
	// 获取文件的最后修改时间
	fileInfo, err := file.Stat()
	if err != nil {
		panic(fmt.Sprintf("无法获取文件信息: %v", err))
	}
	fileModTime := fileInfo.ModTime().UnixMilli()
	return fileModTime != lastModifyTime
}

// 获取文件的后缀名（不带点号）
// println(File_getSuffix("d:\\gg\\example99.txt")) --> txt
func File_getSuffix(file string) string {
	open, err := os.Open(file)
	if err != nil {
		panic(fmt.Sprintf("无法获取文件信息: %v", err))
	}
	return File_getSuffix2(open)
}

// 获取文件的后缀名（不带点号）
// println(File_getSuffix("d:\\gg\\example99.txt")) --> txt
func File_getSuffix2(file *os.File) string {
	// 获取文件名
	fileName := file.Name()

	// 获取文件后缀名
	suffix := filepath.Ext(fileName)

	// 去掉点号
	if suffix != "" {
		suffix = strings.TrimPrefix(suffix, ".")
	}

	return suffix
}

// 从文件路径中提取文件名
// println(File_getName("d:\\gg\\example99.txt")) --> example99.txt
func File_getName(filePath string) string {
	// 获取文件名
	fileName := filepath.Base(filePath)

	// 如果路径为空，触发 panic
	if fileName == "." || fileName == "" {
		panic("文件路径无效")
	}

	return fileName
}

// 从文件名中提取主文件名（去掉后缀名的部分）
// println(File_getPrefix("d:\\gg\\example99.txt")) --> example99
func File_getPrefix(fileName string) string {
	// 去除文件名两端的空格
	fileName = strings.TrimSpace(fileName)

	// 如果文件名为空，触发 panic
	if fileName == "" {
		panic("文件名无效")
	}
	fileName = File_getName(fileName)
	// 获取主文件名（去掉后缀名）
	prefix := strings.TrimSuffix(fileName, filepath.Ext(fileName))
	return prefix
}

// 将字符串以 UTF-8 编码写入文件（覆盖模式），原内容将被覆盖
func File_writeUtf8String(content, path string) {
	// 打开文件（覆盖模式）
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(fmt.Sprintf("无法打开文件: %v", err))
	}
	defer file.Close()

	// 将字符串写入文件
	_, err = file.WriteString(content)
	if err != nil {
		panic(fmt.Sprintf("无法写入文件: %v", err))
	}

	// 确保内容刷新到磁盘
	err = file.Sync()
	if err != nil {
		panic(fmt.Sprintf("无法刷新文件内容: %v", err))
	}
}

// 将字符串以 UTF-8 编码追加写入文件
func File_appendUtf8String(content, path string) {
	// 打开文件（追加模式）
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(fmt.Sprintf("无法打开文件: %v", err))
	}
	defer file.Close()

	// 将字符串写入文件
	_, err = file.WriteString(content)
	if err != nil {
		panic(fmt.Sprintf("无法写入文件: %v", err))
	}

	// 确保内容刷新到磁盘
	err = file.Sync()
	if err != nil {
		panic(fmt.Sprintf("无法刷新文件内容: %v", err))
	}
}

// 将字节数组写入文件（覆盖模式）
func File_writeBytes(data []byte, path string) {
	// 打开文件（覆盖模式）
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(fmt.Sprintf("无法打开文件: %v", err))
	}
	defer file.Close()

	// 将字节数组写入文件
	_, err = file.Write(data)
	if err != nil {
		panic(fmt.Sprintf("无法写入文件: %v", err))
	}

	// 确保内容刷新到磁盘
	err = file.Sync()
	if err != nil {
		panic(fmt.Sprintf("无法刷新文件内容: %v", err))
	}
}
