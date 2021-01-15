package src

import "bytes"

const EXCEL_ONE_PATH = "./Excels/UITextDataAuto.xlsx"
const EXCEL_DIR = "./Excels/"
const HEAD_SHEET_INGORE = "#"
const HEAD_CELL_KEY = "*"

type SheetData struct {
	SheetName  string
	HeadLength int

	DescList    []string
	KeyTypeList []string
	KeyNameList []string

	HasTwoKey bool
	//Two key
	KeyValueList [][]string
}

func GetExcelFilePathName(fileNameWithExt string) string {
	var buffer bytes.Buffer
	buffer.WriteString(EXCEL_DIR)
	buffer.WriteString(fileNameWithExt)
	return buffer.String()
}

// Lua
const OUTPUT_DIR_LUA = "./Outputs/Lua/"

func GetLuaFilePathName(fileName string) string {
	var buffer bytes.Buffer
	buffer.WriteString(OUTPUT_DIR_LUA)
	buffer.WriteString(fileName)
	buffer.WriteString(".lua")
	return buffer.String()
}
