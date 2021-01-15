package excelUtil

import (
	"OneExcel/src"
	"OneExcel/src/excelUtil/luaUtil"
	"OneExcel/src/excelUtil/xlsx"
	"OneExcel/src/fileUtil"
)

func GetData(excelFilePath string) []src.SheetData {
	//sheetDataList := excelize.GetData(config.EXCEL_ONE_PATH)
	return xlsx.GetData(excelFilePath)
}

func WriteData(sheetData src.SheetData) {
	luaUtil.Write(sheetData)
}

func IsExcelClosed() bool {
	return fileUtil.IsExcelClosed()
}

func GetAllData() []src.SheetData {
	var results []src.SheetData
	fileNames := fileUtil.GetAllExcelName()
	for _, fileName := range fileNames {
		filePath := src.GetExcelFilePathName(fileName)
		excelSheetData := GetData(filePath)
		for _, excelData := range excelSheetData {
			results = append(results, excelData)
		}
	}
	return results
}
