package exceltools

import (
	"github.com/extrame/xls"
	"github.com/tealeg/xlsx"
	"strconv"
	"strings"
)

//列
type Cell struct {
	Val string
}

//行
type Row struct {
	CellList []*Cell
}

//表
type Sheet struct {
	SheetName string
	RowList   []*Row
}

//存储excel
func SaveExcel(filename string, data []*Sheet) (err error) {
	file := xlsx.NewFile()
	index := 1
	for _, oneSheet := range data {
		if strings.EqualFold(oneSheet.SheetName, "") {
			oneSheet.SheetName = strconv.Itoa(index)
			index++
		}
		sheet, e := file.AddSheet(oneSheet.SheetName)
		if nil != e {
			err = e
			return
		}
		for _, oneRow := range oneSheet.RowList {
			row := sheet.AddRow()
			for _, oneCell := range oneRow.CellList {
				cell := row.AddCell()
				cell.Value = oneCell.Val
			}
		}
	}
	err = file.Save(filename)
	if nil != err {
		return
	}
	return
}

//读取excel
func GetExcel(filename string) (data []*Sheet, err error) {
	file, err := xlsx.OpenFile(filename)
	if nil != err {
		return
	}
	for _, sheet := range file.Sheets {
		oneSheet := &Sheet{
			SheetName: sheet.Name,
		}
		for _, row := range sheet.Rows {
			oneRow := &Row{}
			for _, cell := range row.Cells {
				oneCel := &Cell{
					Val: cell.String(),
				}
				oneRow.CellList = append(oneRow.CellList, oneCel)
			}
			oneSheet.RowList = append(oneSheet.RowList, oneRow)
		}
		data = append(data, oneSheet)
	}
	return
}

//读取xls
func GetXls(filename string, charset string) (data []*Sheet, err error) {
	if strings.EqualFold("", charset) {
		charset = "utf-8"
	}
	file, err := xls.Open(filename, charset)
	if nil != err {
		return
	}
	for i := 0; i < file.NumSheets(); i++ {
		//sheet
		sheet := file.GetSheet(i)
		oneSheet := &Sheet{}
		oneSheet.SheetName = sheet.Name
		rowNum := sheet.MaxRow
		rowList := make([]*Row, 0, rowNum)

		for r := int(0); r <= int(rowNum); r++ {
			//row
			oneFileRow := sheet.Row(r)
			if nil == oneFileRow {
				continue
			}
			cellList := make([]*Cell, 0, oneFileRow.LastCol())
			for c := oneFileRow.FirstCol(); c <= oneFileRow.LastCol(); c++ {
				val := strings.Trim(oneFileRow.Col(c)," ")

				cellList = append(cellList, &Cell{Val:val})
			}
			rowList = append(rowList, &Row{CellList: cellList})
		}
		oneSheet.RowList = rowList
		data = append(data, oneSheet)
	}
	return
}

