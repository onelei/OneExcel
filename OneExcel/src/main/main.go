package main

import (
	"OneExcel/src"
	"OneExcel/src/excelUtil"
	"OneExcel/src/fileUtil"
	"fmt"
	"os"
	"time"
)

func main() {
	if excelUtil.IsExcelClosed() != true {
		return
	}
	ExportAllExcel()
}

func ExportAllExcel() {
	fileUtil.DeleteDir(src.OUTPUT_DIR_LUA)

	timeNow := time.Now().UnixNano()
	fmt.Println("Exporting...")

	ExportAllExcelSheetData(excelUtil.GetAllData())

	timeNow = time.Now().UnixNano() - timeNow
	timeNow = timeNow / 1e6
	fmt.Printf("\nTotal time : %d ms. \n", timeNow)
	fmt.Println("\nExport Success...")

	fmt.Printf("Press any key to exit...")
	b := make([]byte, 1)
	os.Stdin.Read(b)
}

func ExportOneExcel(excelFilePath string) {
	timeNow := time.Now().UnixNano()
	fmt.Println("Exporting...")

	ExportAllExcelSheetData(excelUtil.GetData(excelFilePath))

	timeNow = time.Now().UnixNano() - timeNow
	timeNow = timeNow / 1e6
	fmt.Printf("\nTotal time : %d ms. \n", timeNow)
	fmt.Println("\nExport Success...")
}

func ExportAllExcelSheetData(sheetDatas []src.SheetData) {
	for i := range sheetDatas {
		sheetData := sheetDatas[i]
		ExportExcelSheetData(sheetData)
	}
}

func ExportExcelSheetData(sheetData src.SheetData) {
	excelUtil.WriteData(sheetData)
}
