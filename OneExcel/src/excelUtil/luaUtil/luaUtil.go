package luaUtil

import (
	"OneExcel/src"
	"OneExcel/src/excelUtil/cellType"
	"OneExcel/src/fileUtil"
	"bufio"
	"log"
	"os"
	"strconv"
)

func Write(sheetData src.SheetData) {
	fileUtil.CreateDir(src.OUTPUT_DIR_LUA)
	fileUtil.CreateFile(src.GetLuaFilePathName(sheetData.SheetName))
	file, err := os.OpenFile(src.GetLuaFilePathName(sheetData.SheetName), os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	write := bufio.NewWriter(file)

	WriteBegin(sheetData, write)
	if sheetData.HasTwoKey {
		WriteTwoKey(sheetData, write)
	} else {
		WriteOneKey(sheetData, write)
	}
	WriteEnd(sheetData, write)

	write.Flush()
}

func WriteOneKey(sheetData src.SheetData, write *bufio.Writer) {
	//Write value
	write.WriteString("local ")
	write.WriteString(sheetData.SheetName)
	write.WriteString(" =\n{\n")

	rowLength := len(sheetData.KeyValueList)

	for i := range sheetData.KeyValueList {
		cellValues := sheetData.KeyValueList[i]

		for j := range cellValues {
			cellValue := cellValues[j]
			if j == 0 {
				write.WriteString("\t[")
				write.WriteString(cellValue)
				write.WriteString("] = {")
			}

			_cellType := sheetData.KeyTypeList[j]
			fixCellValue := cellType.GetLuaCellValue(cellValue, _cellType)
			write.WriteString(fixCellValue)

			if j != sheetData.HeadLength-1 {
				write.WriteString(",")
			}

		}
		if i != rowLength-1 {
			write.WriteString("},\n")
		} else {
			write.WriteString("}\n")
		}
	}
}

func WriteTwoKey(sheetData src.SheetData, write *bufio.Writer) {
	//Write value
	write.WriteString("local ")
	write.WriteString(sheetData.SheetName)
	write.WriteString(" =\n{")

	Key1 := "0"
	//Key2 := "0"
	tmpKey1 := ""
	tmpKey2 := ""

	rowLength := len(sheetData.KeyValueList)
	for i := range sheetData.KeyValueList {
		cellValues := sheetData.KeyValueList[i]
		tmpKey1 = cellValues[0]
		tmpKey2 = cellValues[1]

		bRepeat := false
		if tmpKey1 == Key1 {
			bRepeat = true
		}

		if i != 0 {
			if !bRepeat {
				write.WriteString("},")
			} else {
				write.WriteString(",")
			}
		}

		if !bRepeat {

			write.WriteString("\n\t[")
			write.WriteString(tmpKey1)
			write.WriteString("] = {")

			write.WriteString("\t[")
			write.WriteString(tmpKey2)
			write.WriteString("] = {")
		} else {
			write.WriteString("\t[")
			write.WriteString(tmpKey2)
			write.WriteString("] = {")
		}

		Key1 = tmpKey1
		for j := range cellValues {
			cellValue := cellValues[j]
			_cellType := sheetData.KeyTypeList[j]
			fixCellValue := cellType.GetLuaCellValue(cellValue, _cellType)
			write.WriteString(fixCellValue)
			if j != sheetData.HeadLength-1 {
				write.WriteString(",")

			} else {
				write.WriteString("}")
			}
		}

		if i == rowLength-1 {
			write.WriteString("}\n")
		}
	}

}

func WriteBegin(sheetData src.SheetData, write *bufio.Writer) {
	//Write key
	write.WriteString("local k = {")
	for i := range sheetData.KeyNameList {
		key := sheetData.KeyNameList[i]
		write.WriteString(key)
		write.WriteString("=")
		write.WriteString(strconv.Itoa(i + 1))
		if i != sheetData.HeadLength-1 {
			write.WriteString(",")
		}
	}
	write.WriteString("}")
	write.WriteString("\n\n")
}

func WriteEnd(sheetData src.SheetData, write *bufio.Writer) {
	write.WriteString("}\n")
	write.WriteString(sheetData.SheetName)
	write.WriteString(".__k = k\nreturn ")
	write.WriteString(sheetData.SheetName)
}
