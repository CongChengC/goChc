package model

import (
	"fmt"
	"testing"
)

// Go 的单元测试需要用到 testing 包
// 以及 go test 命令，而且对测试文件也有以下要求
// 1) 被测试的源文件和测试文件必须位于同一个包下
// 2) 测试文件必须要以_test.go 结尾
// 虽然 Go 对测试文件_test.go 的前缀没有强制要求，不过一般我们都设置
// 为被测试的文件的文件名，如：要对 user.go 进行测试，那么测试文件的
// 名字我们通常设置为 user_test.go
// 3) 测试文件中的测试函数为 TestXxx(t *testing.T)
// 其中 Xxx 的首字母必须是大写的英文字母
// 函数参数必须是 test.T 的指针类型（如果是 Benchmark 测试则参数是
// testing.B 的指针类型）

// TestMain函数可以在测试函数执行之前做一些其他操作
func TestMain(m *testing.M) {
	fmt.Println("测试开始：")
	//通过m.Run()来执行测试函数
	m.Run()
}

// 测试方式二之1
func TestUser(t *testing.T) {
	fmt.Println("开始测试User中的相关方法")
	//通过t.Run()来执行子测试函数
	t.Run("测试添加用户:", testAddUser)
}

// 测试方式二之2
// 如果函数名不是以Test开头，那么该函数默认不执行，我们可以将它设置成为一个子测试函数
func testAddUser(t *testing.T) {
	fmt.Println("子测试函数执行：")
	user := &User{}
	//调用添加用户的方法
	user.AddUser()
	user.AddUser2()
}

// 测试方式一
// // 往test数据库添加两个账户
// func TestAddUser(t *testing.T) {
// 	fmt.Println("子测试函数执行：")
// 	user := &User{}
// 	//调用添加用户的方法
// 	user.AddUser()
// 	user.AddUser2()
// }
