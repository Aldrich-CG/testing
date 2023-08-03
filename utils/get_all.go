package utils

import (
	"github.com/xuri/excelize/v2"
)

//type StockInfo struct {
//	Code string
//}

type Code string

func GetAllStocksFromExcel() ([]string, error) {
	var stocks []string

	f, err := excelize.OpenFile("all-1.xlsx")
	if err != nil {
		return nil, err
	}

	rows, err := f.GetRows("Sheet1")
	if err != nil {
		return nil, err
	}

	for _, row := range rows {
		if len(row) >= 2 {
			code := row[0]

			stocks = append(stocks, code)
		}
	}

	return stocks, nil
}
