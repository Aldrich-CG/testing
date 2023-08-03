package main

import (
	"fmt"
	"gupiao_project/model"
	"gupiao_project/servers"
	"gupiao_project/utils"
	"time"
)

func main() {
	var prevData []model.StockData
	data, err := utils.GetAllStocksFromExcel()
	if err != nil {
		fmt.Println("获取所有股票代码失败：", err)
		return
	}

	// 设置程序运行的总时间为10秒
	runDuration := 13 * time.Minute

	// 启动一个定时器，在指定时间后调用 stopFunc 函数
	time.AfterFunc(runDuration, stopFunc)

	// 使用 time.Ticker 定时执行获取股票价格的操作
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		// 调用 FlushPrice 获取新的股票代码列表
		dataResult, prevDataResult := servers.FlushPrice(data, prevData)

		// 将获取到的数据保存为下一次调用的 data
		data = dataResult
		prevData = prevDataResult

		fmt.Println("此次执行后获取到的data为：")
		fmt.Println(data)
		fmt.Println("此次执行后获取到的data长度为：")
		fmt.Println(len(data))
		// 等待下一次执行
		<-ticker.C

	}

}

func stopFunc() {
	fmt.Println("程序运行时间已达到设定的限制，即将停止...")
	// 这里可以执行程序停止的操作，比如关闭连接、保存数据等
	// ...

	// 退出程序
	// 在实际应用中，你可能需要根据实际情况选择合适的退出方式
	// 这里使用了 panic() 来停止程序，只是示例用法，请根据实际情况选择合适的方式
	panic("程序已停止")
}
