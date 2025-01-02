package gg

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestFile(t *testing.T) {
	println(IsWindows())
	println(Str_toString(File_ls("d:\\temp")))
	println(File_isEmpty("d:\\temp"))
	println(File_isDirEmpty("d:\\temp"))
}

func TestDir(t *testing.T) {
	// 测试路径
	paths := []string{
		"d:\\temp",         // 目录
		"d:\\temp\\11.exe", // 文件
		"d:\\temp11",       // 不存在的路径
	}

	for _, path := range paths {
		fmt.Printf("遍历路径: '%s'\n", path)

		files := File_loopFiles(path)
		fmt.Println("找到的文件:")
		for _, file := range files {
			fmt.Println(file)
		}
		fmt.Println()
	}
}

func TestDir2(t *testing.T) {
	println(File_getTmpDirPath())
}

func Test3(t *testing.T) {
	// 测试创建文件
	file := File_file("d:\\gg", "sub_dir", "example.txt")
	defer file.Close()
}

func Test4(t *testing.T) {
	// 测试创建文件
	file := File_newFile("d:/gg/example.txt")
	defer file.Close()

	fmt.Printf("文件创建成功: %s\n", file.Name())
}

func Test5(t *testing.T) {
	println(File_getTmpDir().Name())
}

func Test6(t *testing.T) {
	println(File_getUserDir())
}

func Test7(t *testing.T) {
	println(File_exist("d:\\gg"))
}

func Test8(t *testing.T) {
	open, _ := os.Open("d:\\gg")
	println(File_exist2(open))
}

func Test9(t *testing.T) {
	println(File_lastModifiedTime("d:\\gg").String())
}

func Test10(t *testing.T) {
	println(File_size("d:\\gg"))
}

func Test11(t *testing.T) {
	println(File_getTotalLines("d:\\gg\\example.txt"))
}

func Test12(t *testing.T) {
	println(File_newerThan("D:\\sqlite-mybatis-magic.db", "d:\\gg\\example.txt"))
	println(File_newerThan("d:\\gg\\example.txt", "D:\\sqlite-mybatis-magic.db"))
}

func Test13(t *testing.T) {
	timeMillis := time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC).UnixMilli() // 2023-10-01 00:00:00 的时间戳
	println(File_newerThan2("d:\\gg\\example.txt", timeMillis))
}

func Test14(t *testing.T) {
	touch := File_touch("d:\\gg\\example2.txt")
	println(touch.Name())
}

func Test15(t *testing.T) {
	touch := File_touch("d:\\gg\\aa")
	println(touch.Name())
}

func Test16(t *testing.T) {
	File_delete("d:\\gg\\aa")
}

func Test17(t *testing.T) {
	println(File_clean("d:\\gg"))
}

func Test18(t *testing.T) {
	println(File_mkdir("d:\\gg").Name())
}

func Test19(t *testing.T) {
	file := File_createTempFile()
	println(file.Name())
	file.Close()
	File_delete(file.Name())
}

func Test20(t *testing.T) {
	File_copyFile("d:\\gg\\example.txt", "d:\\gg\\example99.txt")
}

func Test21(t *testing.T) {
	File_copy("d:\\gg", "d:\\gg2", false)
}

func Test22(t *testing.T) {
	File_move("d:\\gg\\example.txt", "d:\\gg2\\aaa.txt", true)
}

func Test23(t *testing.T) {
	println(File_rename("d:\\gg\\example2-111.txt", "example99.txt", true))
}

func Test24(t *testing.T) {
	println(File_contentEquals("d:\\gg\\example99.txt", "d:\\gg2\\example99.txt"))
}

func Test25(t *testing.T) {
	println(File_isModified("d:\\gg\\example99.txt", 123455))
}

func Test26(t *testing.T) {
	println(File_getSuffix("d:\\gg\\example99.txt"))
}

func Test27(t *testing.T) {
	println(File_getPrefix("d:\\gg\\example99.txt"))
}

func Test28(t *testing.T) {
	println(File_getName("d:\\gg\\example99.txt"))
}

func Test29(t *testing.T) {
	File_writeUtf8String("bb", "d:\\gg\\example99.txt")
}

func Test30(t *testing.T) {
	File_appendUtf8String("1234563df", "d:\\gg\\example99.txt")
}

func Test31(t *testing.T) {
	File_writeBytes([]byte("66666666666"), "d:\\gg\\example99.txt")
}

func Test32(t *testing.T) {
	println(File_getAppPath())
}

func TestRead01(t *testing.T) {
	bytes := File_readBytes("d:\\gg\\example99.txt")
	println("len:", len(bytes))
	println("content: " + string(bytes))
}

func TestRead02(t *testing.T) {
	str := File_readString("d:\\gg\\example99.txt")
	println("len:", len(str))
	println("content: " + str)
}

func TestRead03(t *testing.T) {
	aar := File_readLines("d:\\gg\\example99.txt")
	println("len:", len(aar))
	println("content: " + Str_toString(aar))
}

func TestRead04(t *testing.T) {
	// 示例：读取文件内容并按行分割成字符串数组
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	file, err := os.Open("d:\\gg\\example99.txt")
	if err != nil {
		panic(fmt.Sprintf("Error opening file: %s", err))
	}
	defer file.Close()

	fmt.Println("Opened file:", File_readLines2(file))
}

func TestRead05(t *testing.T) {
	// 示例：读取文件所有数据并转换为字符串
	file, err := os.Open("D:\\gg\\example99.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	content, err := File_readString2(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("File content:")
	fmt.Println(content)
}

func TestWrite01(t *testing.T) {
	// 示例：将字符串以 UTF-8 编码写入文件（覆盖模式）
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	// 打开文件（如果文件不存在则创建）
	file, err := os.OpenFile("D:\\gg\\example99.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(fmt.Sprintf("Error opening file: %s", err))
	}
	defer file.Close()

	// 写入内容
	content := "Hello, World!\nThis is a UTF-8 encoded string."
	File_writeUtf8String2(content, file)

	fmt.Println("Content written to file successfully.")
}
