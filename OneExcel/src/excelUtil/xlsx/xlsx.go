package xlsx

import (
	"OneExcel/src"
	"fmt"
	"github.com/tealeg/xlsx"
	"github.com/wxnacy/wgo/arrays"
	"strings"
)

func GetData(filepath string) []src.SheetData {
	//timeNow := time.Now().UnixNano()
	var results []src.SheetData

	f, err := xlsx.OpenFile(filepath)
	if err != nil {
		fmt.Println(err)
		return results
	}

	// Get sheet
	for _, sheet := range f.Sheets {
		sheetName := sheet.Name
		if strings.HasPrefix(sheetName, src.HEAD_SHEET_INGORE) {
			continue
		}
		headLength := 0
		var descList []string
		var keyTypeList []string
		var keyNameList []string

		bHasTwoKey := false

		var keyValueList [][]string

		var rowIndexIgnore []int
		var cellIndexIgnore []int

		for rowIndex, row := range sheet.Rows {
			sheet.Row(rowIndex)
			var cellValueList []string
			for cellIndex, cell := range row.Cells {
				cellValue := cell.String()

				//Desc
				if rowIndex == 0 {
					//Ignored
					if strings.HasPrefix(cellValue, src.HEAD_SHEET_INGORE) {
						cellIndexIgnore = append(cellIndexIgnore, cellIndex)
						continue
					}
					//Two Key
					if cellIndex == 1 && strings.HasPrefix(cellValue, src.HEAD_CELL_KEY) {
						bHasTwoKey = true
					}
					headLength++
					descList = append(descList, cellValue)
				} else {
					//All cell ignored.
					if cellIndex == 0 && strings.HasPrefix(cellValue, src.HEAD_SHEET_INGORE) {
						rowIndexIgnore = append(rowIndexIgnore, rowIndex)
						break
					}
					//One cell ignored.
					if arrays.Contains(cellIndexIgnore, cellIndex) >= 0 {
						continue
					}
					if rowIndex == 1 {
						keyTypeList = append(keyTypeList, cellValue)
					} else if rowIndex == 2 {
						keyNameList = append(keyNameList, cellValue)
					} else {
						//All cell empty, Ignored
						if cellIndex == 0 && cellValue == "" {
							break
						}
						cellValueList = append(cellValueList, cellValue)
					}
				}
			}
			if len(cellValueList) <= 0 {
				continue
			}
			keyValueList = append(keyValueList, cellValueList)
		}
		sheetData := src.SheetData{
			SheetName:  sheetName,
			HeadLength: headLength,

			DescList:    descList,
			KeyTypeList: keyTypeList,
			KeyNameList: keyNameList,

			HasTwoKey: bHasTwoKey,

			KeyValueList: keyValueList,
		}
		results = append(results, sheetData)
	}
	//timeNow = time.Now().UnixNano() - timeNow
	//timeNow = timeNow / 1e6
	//fmt.Printf("\nTotal time : %d ms. \n", timeNow)
	return results
}
