package utils

import (
	"gupiao_project/model"
)

func ComparePrice(data []string, prevData, currentData []model.StockData) ([]string, []model.StockData) {
	//fmt.Println("CCCCCCCCCCCCCCCC")
	// 示例逻辑代码如下：
	for _, current := range currentData {
		for _, prev := range prevData {
			if current.Code == prev.Code {
				if current.Price < prev.Price*0.999 && IsBetweenTime09150923() {
					// 低于上一次价格的0.5%，且低于次数小于5，丢弃
					// TODO: 在这里处理丢弃的逻辑

					data = RemoveCodeFromData(data, current.Code)
				} else if current.Price > prev.Price*1.001 && IsBetweenTime09150923() {
					// 高于上一次价格的0.5%，丢弃
					// TODO: 在这里处理丢弃的逻辑
					data = RemoveCodeFromData(data, current.Code)
				} else if current.Price < prev.Price && IsBetweenTime923926() {
					//
					// 9:23到9:26之间，且当前价格小于上一次价格，丢弃
					// TODO: 在这里处理丢弃的逻辑
					data = RemoveCodeFromData(data, current.Code)
				} else {
					// 其他情况，将当前数据添加到prevData中
					prevData = append(prevData, current)
				}
			}
		}
		//fmt.Println("ddddddd", current.Code)
	}
	// 将当前数据保存为上一次数据，用于下一次比较
	prevData = append(prevData[:0], currentData...)

	return data, prevData
}
