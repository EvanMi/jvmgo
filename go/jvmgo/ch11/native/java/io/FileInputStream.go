package io

import (
	"jvmgo/ch11/native"
	"jvmgo/ch11/rtda"
)

const fis = "java/io/FileInputStream"

func init() {
	native.Register(fis, "initIDs", "()V", initIDs)
}

func initIDs(frame *rtda.Frame) {

}
