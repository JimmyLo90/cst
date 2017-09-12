package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"

	"github.com/xuri/excelize"
)

var sheet2 string = "Sheet2"
var sheet3 string = "Sheet3"
var sheet4 string = "Sheet4"

func main() {
	xlsx := excelize.NewFile()

	brand(xlsx)

	jiyou(xlsx)

	//保存文件
	xlsx.SetActiveSheet(2)
	err := xlsx.SaveAs("D:\\test.xlsx")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
func brand(xlsx *excelize.File) {
	xlsx.NewSheet(2, sheet2)

	//品牌导图
	xlsx.SetCellValue(sheet2, "A1", "车型库来源")
	xlsx.SetCellValue(sheet2, "B1", "1")
	xlsx.SetCellValue(sheet2, "C1", "1=车商通；2=乐车邦车型")

	xlsx.SetCellValue(sheet2, "A2", "品牌")
	xlsx.SetCellValue(sheet2, "B2", "厂商")
	xlsx.SetCellValue(sheet2, "C2", "车型")
	xlsx.SetCellValue(sheet2, "D2", "年款")
	xlsx.SetCellValue(sheet2, "E2", "型号")
	xlsx.SetCellValue(sheet2, "F2", "a/b/c/d")
	xlsx.SetCellValue(sheet2, "G2", "机油用量")

	setNum := 3
	var buf bytes.Buffer
	for i := 0; i <= 10; i++ {
		buf.Reset()
		buf.WriteString("A")
		buf.WriteString(strconv.Itoa(setNum))
		xlsx.SetCellValue(sheet2, buf.String(), "本田")

		buf.Reset()
		buf.WriteString("B")
		buf.WriteString(strconv.Itoa(setNum))
		xlsx.SetCellValue(sheet2, buf.String(), "广汽本田")

		buf.Reset()
		buf.WriteString("C")
		buf.WriteString(strconv.Itoa(setNum))
		xlsx.SetCellValue(sheet2, buf.String(), "凌派")

		buf.Reset()
		buf.WriteString("D")
		buf.WriteString(strconv.Itoa(setNum))
		xlsx.SetCellValue(sheet2, buf.String(), "2015款")

		buf.Reset()
		buf.WriteString("E")
		buf.WriteString(strconv.Itoa(setNum))
		xlsx.SetCellValue(sheet2, buf.String(), "1.8L 手动豪华版")

		buf.Reset()
		buf.WriteString("F")
		buf.WriteString(strconv.Itoa(setNum))
		xlsx.SetCellValue(sheet2, buf.String(), "a")

		buf.Reset()
		buf.WriteString("G")
		buf.WriteString(strconv.Itoa(setNum))
		xlsx.SetCellValue(sheet2, buf.String(), "4.0")

		setNum++
	}

}

func jiyou(xlsx *excelize.File) {
	//机油数据规格
	xlsx.NewSheet(3, sheet3)
	xlsx.SetCellValue(sheet3, "A1", "填写顺序按照同规格的机油价格从低到高（例如：最便宜的机油填写在机油A中）！！！此表格只需填写常用机油（含原厂和非原厂）！！！")
	xlsx.MergeCell(sheet3, "A1", "J1")

	xlsx.SetCellValue(sheet3, "A2", "机油代码")
	xlsx.SetCellValue(sheet3, "B2", "规格（L）")
	xlsx.SetCellValue(sheet3, "C2", "原厂")
	xlsx.SetCellValue(sheet3, "D2", "品牌")
	xlsx.SetCellValue(sheet3, "E2", "系列")
	xlsx.SetCellValue(sheet3, "F2", "粘度")
	xlsx.SetCellValue(sheet3, "G2", "级别")
	xlsx.SetCellValue(sheet3, "H2", "类型（矿物/半合成/全合成）")
	xlsx.SetCellValue(sheet3, "I2", "零件编号")
	xlsx.SetCellValue(sheet3, "J2", "机油单价")

	xlsx.SetCellValue(sheet3, "A3", "机油A")

	jiyouSetNum := 3
	var jiyouBuf bytes.Buffer

	jiyouANum := 4
	for i := 1; i <= jiyouANum; i++ {
		jiyouBuf.Reset()
		jiyouBuf.WriteString("B")
		jiyouBuf.WriteString(strconv.Itoa(jiyouSetNum + i))
		xlsx.SetCellValue(sheet3, jiyouBuf.String(), i)

		jiyouBuf.Reset()
		jiyouBuf.WriteString("C")
		jiyouBuf.WriteString(strconv.Itoa(jiyouSetNum + i))
		xlsx.SetCellValue(sheet3, jiyouBuf.String(), i)

		jiyouBuf.Reset()
		jiyouBuf.WriteString("D")
		jiyouBuf.WriteString(strconv.Itoa(jiyouSetNum + i))
		xlsx.SetCellValue(sheet3, jiyouBuf.String(), i)

		jiyouBuf.Reset()
		jiyouBuf.WriteString("E")
		jiyouBuf.WriteString(strconv.Itoa(jiyouSetNum + i))
		xlsx.SetCellValue(sheet3, jiyouBuf.String(), i)

		jiyouBuf.Reset()
		jiyouBuf.WriteString("F")
		jiyouBuf.WriteString(strconv.Itoa(jiyouSetNum + i))
		xlsx.SetCellValue(sheet3, jiyouBuf.String(), i)

		jiyouBuf.Reset()
		jiyouBuf.WriteString("G")
		jiyouBuf.WriteString(strconv.Itoa(jiyouSetNum + i))
		xlsx.SetCellValue(sheet3, jiyouBuf.String(), i)

		jiyouBuf.Reset()
		jiyouBuf.WriteString("H")
		jiyouBuf.WriteString(strconv.Itoa(jiyouSetNum + i))
		xlsx.SetCellValue(sheet3, jiyouBuf.String(), i)

		jiyouBuf.Reset()
		jiyouBuf.WriteString("I")
		jiyouBuf.WriteString(strconv.Itoa(jiyouSetNum + i))
		xlsx.SetCellValue(sheet3, jiyouBuf.String(), i)

		jiyouBuf.Reset()
		jiyouBuf.WriteString("J")
		jiyouBuf.WriteString(strconv.Itoa(jiyouSetNum + i))
		xlsx.SetCellValue(sheet3, jiyouBuf.String(), i)

		if i == jiyouANum {
			xlsx.MergeCell(sheet3, "A"+strconv.Itoa(jiyouSetNum), "A"+strconv.Itoa(jiyouSetNum+i))
			jiyouSetNum += i
		}
	}
}
