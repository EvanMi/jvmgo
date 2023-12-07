package main

import (
	"fmt"
	"jvmgo/ch11/clazzpath"
	"jvmgo/ch11/instructions/base"
	"jvmgo/ch11/rtda"
	"jvmgo/ch11/rtda/heap"
	"strings"
)

type JVM struct {
	cmd         *Cmd
	classLoader *heap.ClassLoader
	mainThread  *rtda.Thread
}

func newJVM(cmd *Cmd) *JVM {
	cp := clazzpath.Parse(cmd.XjreOption, cmd.cpOption)
	classLoader := heap.NewClassLoader(cp, cmd.verboseClassFlag)
	return &JVM{
		cmd:         cmd,
		classLoader: classLoader,
		mainThread:  rtda.NewThread(),
	}
}

func (jvm *JVM) start() {
	jvm.initVM()
	jvm.execMain()
}

func (jvm *JVM) initVM() {
	vmClass := jvm.classLoader.LoadClass("sun/misc/VM")
	base.InitClass(jvm.mainThread, vmClass)
	interpret(jvm.mainThread, jvm.cmd.verboseInstFlag)
}

func (jvm *JVM) execMain() {
	className := strings.Replace(jvm.cmd.class, ".", "/", -1)
	mainClass := jvm.classLoader.LoadClass(className)
	mainMethod := mainClass.GetMainMethod()
	if mainMethod == nil {
		fmt.Printf("Main method not found in class %s\n", jvm.cmd.class)
		return
	}
	argsArr := jvm.createArgsArray()
	frame := jvm.mainThread.NewFrame(mainMethod)
	frame.LocalVars().SetRef(0, argsArr)
	jvm.mainThread.PushFrame(frame)
	interpret(jvm.mainThread, jvm.cmd.verboseInstFlag)
}

func (jvm *JVM) createArgsArray() *heap.Object {
	stringClass := jvm.classLoader.LoadClass("java/lang/String")
	argsLen := uint(len(jvm.cmd.args))
	argsArr := stringClass.ArrayClass().NewArray(argsLen)
	jArgs := argsArr.Refs()
	for i, arg := range jvm.cmd.args {
		jArgs[i] = heap.JString(jvm.classLoader, arg)
	}
	return argsArr
}
