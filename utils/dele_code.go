package utils

import "time"

// 从data中去除对应代码
func RemoveCodeFromData(data []string, codeToRemove string) []string {
	for i, code := range data {
		if code == codeToRemove {
			// 从data中去除对应代码
			data = append(data[:i], data[i+1:]...)
			break
		}
	}
	return data
}

// 示例函数，模拟判断是否在9:23到9:26之间的函数
func IsBetweenTime923926() bool {
	now := time.Now()
	startTime := time.Date(now.Year(), now.Month(), now.Day(), 9, 23, 0, 0, now.Location())
	endTime := time.Date(now.Year(), now.Month(), now.Day(), 9, 26, 0, 0, now.Location())
	return now.After(startTime) && now.Before(endTime)
}

func IsBetweenTime09150923() bool {
	now := time.Now()
	startTime := time.Date(now.Year(), now.Month(), now.Day(), 9, 15, 0, 0, now.Location())
	endTime := time.Date(now.Year(), now.Month(), now.Day(), 9, 22, 59, 0, now.Location())
	return now.After(startTime) && now.Before(endTime)
}
