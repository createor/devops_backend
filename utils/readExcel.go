package utils

import (
	"fmt"
	"path"

	"github.com/tealeg/xlsx"
)

/*
读取表格
*/

// file 文件路径
// col 列的数量
func ReadTable(file string, col int) ([][]string, error) {
	// 判断文件是否是xlsx格式
	fileExt := path.Ext(file)
	if fileExt != ".xlsx" {
		return nil, fmt.Errorf("格式错误:%v", fileExt)
	}
	result := make([][]string, 0)
	xlFile, err := xlsx.OpenFile(file)
	if err != nil {
		return nil, err
	}
	sheet := xlFile.Sheets[0] // 读取第一个sheet
	for _, row := range sheet.Rows {
		if col != 0 && len(row.Cells) != col {
			return nil, fmt.Errorf("数据缺失")
		}
		data := make([]string, 0)
		for _, cell := range row.Cells {
			value := cell.String() // 获取单元格数据
			data = append(data, value)
		}
		result = append(result, data)
	}
	return result, nil
}
