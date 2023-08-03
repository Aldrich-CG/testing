package servers

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"gupiao_project/model"
	"gupiao_project/utils"
	"sync"
)

func FlushPrice(data []string, prevData []model.StockData) ([]string, []model.StockData) {
	//fmt.Println(prevData)
	//fmt.Println(data)
	// 创建一个新的 Excel 文件
	f := excelize.NewFile()

	// 设置工作表名称
	sheetName := "Sheet1"

	//// 写入表头
	//f.SetCellValue(sheetName, "A1", "时间")

	// 定义上一次和当前获取的股票数据的切片
	var currentData []model.StockData
	// 设置并发的 goroutine 数量
	concurrency := 15

	// 创建一个 WaitGroup，用于等待所有 goroutine 完成
	var wg sync.WaitGroup

	// 创建一个有缓冲的 channel，用于控制并发的 goroutine 数量
	semaphore := make(chan struct{}, concurrency)

	// 启动 goroutine 获取实时股票价格
	for _, stockCode := range data {
		// 申请一个信号量
		semaphore <- struct{}{}

		wg.Add(1)
		go func(code string) {
			defer wg.Done()

			// 在 goroutine 中调用获取股票实时价格的函数
			stockQuote, err := utils.GetStockQuote(code)
			//fmt.Println(stockQuote)
			if err != nil {
				fmt.Println("获取股票实时价格失败：", err)
				return
			}

			// 将获取到的数据存入当前数据切片
			currentData = append(currentData, model.StockData{Code: code, Price: stockQuote["当前价格"]})

			// 释放信号量
			<-semaphore
		}(stockCode)
	}

	// 等待所有 goroutine 完成
	wg.Wait()
	//fmt.Println("---------", prevData)
	// 判断是否需要进行比较和处理
	if len(prevData) == 0 {
		// 第一次获取数据，直接保存到prevData切片中
		prevData = append(prevData, currentData...)
	} else {
		data, prevData = utils.ComparePrice(data, prevData, currentData)
		// 不是第一次获取数据，需要进行比较和处理
		// TODO: 在这里实现根据需求的比较和处理逻辑

		// 示例逻辑：将currentData中的数据与prevData中的数据进行比较
		// 如果当前价格低于上一次价格的0.5%，且低于次数小于5，则丢弃
		// 如果当前价格高于上一次价格的0.5%，则丢弃
		// 如果当前时间是9:23到9:26，并且当前价格小于等于上一次价格，则丢弃
		// 否则，将currentData中的数据添加到prevData中
		// 你可以根据实际需求来修改这里的比较和处理逻辑

		// 示例逻辑代码如下：
		//for _, current := range currentData {
		//	for _, prev := range prevData {
		//		if current.Code == prev.Code {
		//			if current.Price <= prev.Price*0.995 {
		//				// 低于上一次价格的0.5%，且低于次数小于5，丢弃
		//				// TODO: 在这里处理丢弃的逻辑
		//				data = utils.RemoveCodeFromData(data, current.Code)
		//			} else if current.Price > prev.Price*1.005 {
		//				// 高于上一次价格的0.5%，丢弃
		//				// TODO: 在这里处理丢弃的逻辑
		//				data = utils.RemoveCodeFromData(data, current.Code)
		//			} else if current.Price <= prev.Price && utils.IsBetweenTime923926() {
		//				// 9:23到9:26之间，且当前价格小于等于上一次价格，丢弃
		//				// TODO: 在这里处理丢弃的逻辑
		//				data = utils.RemoveCodeFromData(data, current.Code)
		//			} else {
		//				// 其他情况，将当前数据添加到prevData中
		//				prevData = append(prevData, current)
		//			}
		//		}
		//	}
		//}
	}
	//fmt.Println(prevData)
	// 处理获取到的结果
	//for _, stockData := range currentData {
	//	fmt.Printf("股票代码：%s，实时价格：%f\n", stockData.Code, stockData.Price)
	//}
	// 写入数据到 Excel 表格

	// 写入数据到工作表中
	for i, cell := range data {
		cellAddr, _ := excelize.CoordinatesToCellName(1, i+1) // 写入到第一列的第 i+1 行
		f.SetCellValue(sheetName, cellAddr, cell)
	}

	// 保存文件
	fileName := "output.xlsx"
	if err := f.SaveAs(fileName); err != nil {
		fmt.Println("保存 Excel 文件失败:", err)
	}

	fmt.Println("所有股票实时价格获取完成！")

	return data, prevData
}
