package main

import (
	"time"

	"golang.org/x/sys/windows"
)

// showNotification 使用 Windows API 显示一个消息框。
func showNotification(title, message string) {
	// 使用 Windows API 显示通知
	const (
		MB_OK = 0x0
	)
	hwnd := HWND(0) // 使用 0 表示没有父窗口
	uType := uint32(MB_OK)
	windows.MessageBox(hwnd, windows.StringToUTF16Ptr(message), windows.StringToUTF16Ptr(title), uType)
}

func main() {
	// 设置提醒时间间隔，例如每 15 分钟提醒一次
	interval := 1 * time.Minute

	for {
		// 显示通知
		showNotification("提醒", "严守买卖纪律，情绪晴朗、最主流、最核心、最强攻量！")

		// 等待指定的时间间隔
		time.Sleep(interval)
	}
}

// 定义 HWND 类型，它是 windows.HWND 的别名
type HWND = windows.HWND
