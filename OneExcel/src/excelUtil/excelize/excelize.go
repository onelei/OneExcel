package excelize

import (
	"OneExcel/src"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"strings"
	"time"
)

func GetData(filepath string) []src.SheetData {
	timeNow := time.Now().UnixNano()
	var results []src.SheetData
	f, err := excelize.OpenFile(filepath)
	if err != nil {
		fmt.Println(err)
		return results
	}

	// Get sheet
	for _, sheetName := range f.GetSheetMap() {
		if strings.HasPrefix(sheetName, src.HEAD_SHEET_INGORE) {
			continue
		}

		var descList []string
		var typeList []string
		var nameList []string
		var valueList []string

		rows := f.GetRows(sheetName)
		for rowIndex, row := range rows {
			for _, cell := range row {
				//Desc
				if rowIndex == 0 {
					descList = append(descList, cell)
				} else if rowIndex == 1 {
					typeList = append(typeList, cell)

				} else if rowIndex == 2 {
					nameList = append(nameList, cell)

				} else {
					valueList = append(valueList, cell)
				}
			}
		}

	}
	timeNow = time.Now().UnixNano() - timeNow
	timeNow = timeNow / 1e6
	fmt.Printf("\nTotal time : %d ms. \n", timeNow)
	return results
}
