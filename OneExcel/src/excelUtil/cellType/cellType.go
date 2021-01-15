package cellType

import (
	"bytes"
	"log"
	"strconv"
	"strings"
)

func GetLuaCellValue(cellValue string, cellType string) string {
	fixCellType := strings.ToLower(cellType)
	var buffer bytes.Buffer
	switch fixCellType {
	case "int", "float":
		return GetNumber(cellValue)
		break
	case "bool":
		return GetBool(cellValue)
		break
	case "string":
		return GetString(cellValue, buffer)
	case "vector2":
		return GetVector2(cellValue, buffer)
		break
	case "vector3":
		return GetVector3(cellValue, buffer)
		break
	case "vector4":
		return GetVector4(cellValue, buffer)
		break
	case "bool[]":
		return GetBoolArray(cellValue, buffer)
		break
	case "int[]", "float[]":
		return GetNumberArray(cellValue, buffer)
		break
	case "string[]":
		return GetStringArray(cellValue, buffer)
		break
	case "vector2[]":
		return GetVector2Array(cellValue, buffer)
		break
	case "vector3[]":
		return GetVector3Array(cellValue, buffer)
		break
	case "vector4[]":
		return GetVector4Array(cellValue, buffer)
		break
	default:
		log.Fatalf("Unknow type: %s ", cellType)
		return cellValue
		break
	}
	return cellValue
}

func GetNumber(cellValue string) string {
	return cellValue
}

func GetBool(cellValue string) string {
	value, err := strconv.Atoi(cellValue)
	if err != nil {
		return "false"
	}
	if value > 0 {
		return "true"
	}
	return "false"
}

func GetString(cellValue string, buffer bytes.Buffer) string {
	buffer.WriteString("'")
	buffer.WriteString(cellValue)
	buffer.WriteString("'")
	return buffer.String()
}

func GetVector2(cellValue string, buffer bytes.Buffer) string {
	//1=2
	//Vector2(1,2)
	cellArray := strings.Split(cellValue, "=")
	if len(cellArray) != 2 {
		log.Fatalf("GetVector2 error : %s ", cellValue)
		return cellValue
	}
	buffer.WriteString("Vector2(")
	buffer.WriteString(cellArray[0])
	buffer.WriteString(",")
	buffer.WriteString(cellArray[1])
	buffer.WriteString(")")
	return buffer.String()
}

func GetVector3(cellValue string, buffer bytes.Buffer) string {
	//1=2=3
	//Vector3(1,2,3)
	cellArray := strings.Split(cellValue, "=")
	if len(cellArray) != 3 {
		log.Fatalf("GetVector3 error : %s ", cellValue)
		return cellValue
	}
	buffer.WriteString("Vector3(")
	buffer.WriteString(cellArray[0])
	buffer.WriteString(",")
	buffer.WriteString(cellArray[1])
	buffer.WriteString(",")
	buffer.WriteString(cellArray[2])
	buffer.WriteString(")")
	return buffer.String()
}

func GetVector4(cellValue string, buffer bytes.Buffer) string {
	//1=2=3=4
	//Vector4(1,2,3,4)
	cellArray := strings.Split(cellValue, "=")
	if len(cellArray) != 4 {
		log.Fatalf("GetVector4 error : %s ", cellValue)
		return cellValue
	}
	buffer.WriteString("Vector4(")
	buffer.WriteString(cellArray[0])
	buffer.WriteString(",")
	buffer.WriteString(cellArray[1])
	buffer.WriteString(",")
	buffer.WriteString(cellArray[2])
	buffer.WriteString(",")
	buffer.WriteString(cellArray[3])
	buffer.WriteString(")")
	return buffer.String()
}

func GetNumberArray(cellValue string, buffer bytes.Buffer) string {
	//1|2|3|4|5
	//{1,2,3,4,5}
	cellArray := strings.Split(cellValue, "|")
	buffer.WriteString("{")
	length := len(cellArray)
	length--
	for i := range cellArray {
		num := cellArray[i]
		buffer.WriteString(GetNumber(num))
		if i != length {
			buffer.WriteString(",")
		}
	}
	buffer.WriteString("}")
	return buffer.String()
}

func GetBoolArray(cellValue string, buffer bytes.Buffer) string {
	//0|1|0|1|0
	//{false,true,false,true,false}
	cellArray := strings.Split(cellValue, "|")
	buffer.WriteString("{")
	length := len(cellArray)
	length--
	for i := range cellArray {
		num := cellArray[i]
		buffer.WriteString(GetBool(num))
		if i != length {
			buffer.WriteString(",")
		}
	}
	buffer.WriteString("}")
	return buffer.String()
}

func GetStringArray(cellValue string, buffer bytes.Buffer) string {
	//"a"|"b"|"c"|"
	//{"a","b","c"}
	cellArray := strings.Split(cellValue, "|")
	buffer.WriteString("{")
	length := len(cellArray)
	length--
	var buffer2 bytes.Buffer
	for i := range cellArray {
		num := cellArray[i]
		buffer2.Reset()
		buffer.WriteString(GetString(num, buffer2))
		if i != length {
			buffer.WriteString(",")
		}
	}
	buffer.WriteString("}")
	return buffer.String()
}

func GetVector2Array(cellValue string, buffer bytes.Buffer) string {
	cellArray := strings.Split(cellValue, "|")
	buffer.WriteString("{")
	length := len(cellArray)
	length--
	var buffer2 bytes.Buffer
	for i := range cellArray {
		num := cellArray[i]
		buffer2.Reset()
		buffer.WriteString(GetVector2(num, buffer2))
		if i != length {
			buffer.WriteString(",")
		}
	}
	buffer.WriteString("}")
	return buffer.String()
}

func GetVector3Array(cellValue string, buffer bytes.Buffer) string {
	cellArray := strings.Split(cellValue, "|")
	buffer.WriteString("{")
	length := len(cellArray)
	length--
	var buffer2 bytes.Buffer
	for i := range cellArray {
		num := cellArray[i]
		buffer2.Reset()
		buffer.WriteString(GetVector3(num, buffer2))
		if i != length {
			buffer.WriteString(",")
		}
	}
	buffer.WriteString("}")
	return buffer.String()
}

func GetVector4Array(cellValue string, buffer bytes.Buffer) string {
	cellArray := strings.Split(cellValue, "|")
	buffer.WriteString("{")
	length := len(cellArray)
	length--
	var buffer2 bytes.Buffer
	for i := range cellArray {
		num := cellArray[i]
		buffer2.Reset()
		buffer.WriteString(GetVector4(num, buffer2))
		if i != length {
			buffer.WriteString(",")
		}
	}
	buffer.WriteString("}")
	return buffer.String()
}
