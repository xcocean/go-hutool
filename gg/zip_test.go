package gg

import "testing"

func Test01(t *testing.T) {
	Zip_unzip("d:\\gg\\gg.zip", "d:\\gg\\ex1")
}

func Test02(t *testing.T) {
	println(Zip_zip("d:\\gg"))
}

func Test03(t *testing.T) {
	File_appendToZip("d:\\gg.zip", "d:\\gg2\\aaa.txt")
}
