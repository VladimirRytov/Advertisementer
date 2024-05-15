package exelexporter

import (
	"context"
	"errors"
	"io"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/VladimirRytov/advertisementer/internal/datatransferobjects"

	"github.com/tealeg/xlsx/v3"
)

type FileWriter interface {
	OpenForWrite(string) (io.WriteCloser, error)
}

type FileByNamer interface {
	FileByName(context.Context, string) (datatransferobjects.FileDTO, error)
}

type ExelReporter struct {
	fileWriter   FileWriter
	fileGetter   FileByNamer
	orderCounts  map[int]int
	reportParams datatransferobjects.ReportParams
	deployFolder string
}

func NewExelReporter(fileWriter FileWriter, fileGetter FileByNamer) *ExelReporter {
	return &ExelReporter{
		fileWriter:  fileWriter,
		fileGetter:  fileGetter,
		orderCounts: make(map[int]int)}
}

func (er *ExelReporter) NewReport(ctx context.Context, blocks []datatransferobjects.BlockAdvertisementReport, lines []datatransferobjects.LineAdvertisementReport, data datatransferobjects.ReportParams) error {
	er.reportParams = data
	table := xlsx.NewFile()

	err := er.checkFilePaths()
	if err != nil {
		return err
	}

	err = er.createDeployFolder()
	if err != nil {
		return err
	}

	err = os.MkdirAll(filepath.Join(er.reportParams.DeployPath, er.deployFolder, "Блочные объявления"), 0750)
	if err != nil {
		return err
	}

	lineSheet, err := er.newLineDoc(ctx, lines)
	if err != nil {
		return err
	}

	blockSheet, err := er.newBlockDoc(ctx, blocks)
	if err != nil {
		return err
	}

	_, err = table.AppendSheet(*lineSheet, "Строковые объявления")
	if err != nil {
		return err
	}

	_, err = table.AppendSheet(*blockSheet, "Блочные объявления")
	if err != nil {
		return err
	}
	return table.Save(filepath.Join(er.reportParams.DeployPath, er.deployFolder, "Объявления.xlsx"))
}

func (er *ExelReporter) newLineDoc(ctx context.Context, lines []datatransferobjects.LineAdvertisementReport) (*xlsx.Sheet, error) {
	var b strings.Builder

	select {
	case <-ctx.Done():
		return nil, ctx.Err()

	default:
		sh, err := xlsx.NewSheet("Стрококвые объявления")
		if err != nil {
			return nil, err
		}
		cols := xlsx.NewColForRange(1, 4)
		cols.SetWidth(25)
		sh.SetColParameters(cols)
		sh.SetColWidth(2, 2, 40)
		row := sh.AddRow()
		cel := row.GetCell(1)
		celDates := cel.GetStyle()
		celDates.Font.Size = 18
		celDates.Alignment.WrapText = true
		if er.reportParams.FromDate.Equal(er.reportParams.ToDate) {
			b.WriteString("за " + er.reportParams.FromDate.Format("02.01.2006"))
		} else {
			b.WriteString(er.reportParams.FromDate.Format("02.01.2006") + "-" + er.reportParams.ToDate.Format("02.01.2006"))
		}
		cel.SetString("Строковые объявления\n" + b.String())
		b.Reset()

		row = sh.AddRow()
		cell1 := row.AddCell()
		cell1st := cell1.GetStyle()
		cell1st.Font.Size = 14
		cell1.SetString("Номер объявления")

		cell2 := row.AddCell()
		cell2.SetString("Текст объявления")
		cell2st := cell2.GetStyle()
		cell2st.Font.Size = 14

		cell3 := row.AddCell()
		cell3.SetString("Метки")
		cell3st := cell3.GetStyle()
		cell3st.Font.Size = 14

		cell4 := row.AddCell()
		cell4.SetString("Замечания")
		cell4st := cell4.GetStyle()
		cell4st.Font.Size = 14

		for i := range lines {
			row := sh.AddRow()
			cell1 := row.AddCell()
			if er.orderCounts[lines[i].Line.OrderID] == 0 {
				cell1.SetString("№ " + strconv.Itoa(lines[i].Line.OrderID) + ".")
			} else {
				cell1.SetString("№ " + strconv.Itoa(lines[i].Line.OrderID) + "/" + strconv.Itoa(er.orderCounts[lines[i].Line.OrderID]) + ".")
			}
			er.orderCounts[lines[i].Line.OrderID]++

			cell2 := row.AddCell()
			cell2st := cell2.GetStyle()
			cell2st.Alignment.WrapText = true
			cell2.SetString(lines[i].Line.Text)

			cell3 := row.AddCell()
			cel3St := cell3.GetStyle()
			cel3St.Alignment.WrapText = true

			cell4 := row.AddCell()
			cell4St := cell4.GetStyle()
			cell4St.Alignment.WrapText = true

			b.WriteString(strings.Join(lines[i].Line.Tags, ", ") + "\n\n" + strings.Join(lines[i].Line.ExtraCharges, ", "))
			cell3.SetString(b.String())
			b.Reset()
			if !lines[i].Paid {
				cell4.SetString("Заказ № " + strconv.Itoa(lines[i].Line.OrderID) + " не оплачен. ")
				cell4St.Font.Color = "FF0000"
			}
		}
		return sh, nil
	}
}

func (er *ExelReporter) newBlockDoc(ctx context.Context, blocks []datatransferobjects.BlockAdvertisementReport) (*xlsx.Sheet, error) {
	var b strings.Builder
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		sh, err := xlsx.NewSheet("Блочные объявления")
		if err != nil {
			return nil, err
		}
		cols := xlsx.NewColForRange(1, 6)
		cols.SetWidth(25)
		sh.SetColParameters(cols)
		sh.SetColWidth(2, 2, 40)
		row := sh.AddRow()
		cel := row.GetCell(1)
		celDates := cel.GetStyle()
		celDates.Font.Size = 18
		celDates.Alignment.WrapText = true
		if er.reportParams.FromDate.Equal(er.reportParams.ToDate) {
			b.WriteString("за " + er.reportParams.FromDate.Format("02.01.2006"))
		} else {
			b.WriteString(er.reportParams.FromDate.Format("02.01.2006") + "-" + er.reportParams.ToDate.Format("02.01.2006"))
		}
		cel.SetString("Блочные объявления\n" + b.String())
		b.Reset()

		row = sh.AddRow()
		cell1 := row.AddCell()
		cell1st := cell1.GetStyle()
		cell1st.Font.Size = 14
		cell1.SetString("Номер объявления")

		cell2 := row.AddCell()
		cell2.SetString("Площадь")
		cell2st := cell2.GetStyle()
		cell2st.Font.Size = 14

		cell3 := row.AddCell()
		cell3.SetString("Текст объявления")
		cell3st := cell3.GetStyle()
		cell3st.Font.Size = 14

		cell4 := row.AddCell()
		cell4.SetString("Метки")
		cell4st := cell4.GetStyle()
		cell4st.Font.Size = 14

		cell5 := row.AddCell()
		cell5.SetString("Имя файла")
		cell5st := cell5.GetStyle()
		cell5st.Font.Size = 14

		cell6 := row.AddCell()
		cell6.SetString("Замечания")
		cell6st := cell6.GetStyle()
		cell6st.Font.Size = 14

		for i := range blocks {
			row := sh.AddRow()
			cell1 := row.AddCell()
			if er.orderCounts[blocks[i].Block.OrderID] == 0 {
				cell1.SetString("№ " + strconv.Itoa(blocks[i].Block.OrderID) + ".")
			} else {
				cell1.SetString("№ " + strconv.Itoa(blocks[i].Block.OrderID) + "/" + strconv.Itoa(er.orderCounts[blocks[i].Block.OrderID]) + ".")
			}
			er.orderCounts[blocks[i].Block.OrderID]++

			cell2 := row.AddCell()
			cell2st := cell2.GetStyle()
			cell2st.Alignment.WrapText = true
			cell2.SetInt(int(blocks[i].Block.Size))

			cell3 := row.AddCell()
			cell3st := cell3.GetStyle()
			cell3st.Alignment.WrapText = true
			cell3.SetString(blocks[i].Block.Text)

			cell4 := row.AddCell()
			cel4St := cell4.GetStyle()
			cel4St.Alignment.WrapText = true

			b.WriteString(strings.Join(blocks[i].Block.Tags, ", ") + "\n\n" + strings.Join(blocks[i].Block.ExtraCharges, ", "))
			cell3.SetString(b.String())
			b.Reset()

			cell5 := row.AddCell()
			cel5St := cell5.GetStyle()
			cel5St.Alignment.WrapText = true

			cell6 := row.AddCell()
			cel6St := cell6.GetStyle()
			cel6St.Alignment.WrapText = true

			cell5.SetString(blocks[i].Block.FileName)
			if !blocks[i].Paid {
				cell6.SetString("Заказ № " + strconv.Itoa(blocks[i].Block.OrderID) + " не оплачен. ")
			}

			file, err := er.fileGetter.FileByName(ctx, path.Join(er.reportParams.BlocksFolderPath, blocks[i].Block.FileName))
			if err != nil {
				cell6.SetString(cell6.String() + "\n" + "Файл " + blocks[i].Block.FileName + " не найден. ")
				cel6St.Font.Color = "FF0000"
				continue
			}

			deployFile, err := er.fileWriter.OpenForWrite(path.Join(er.reportParams.DeployPath, er.deployFolder, "Блочные объявления", blocks[i].Block.FileName))
			if err != nil {
				return nil, err
			}
			defer deployFile.Close()

			writed, _ := deployFile.Write(file.Data)
			if writed != len(file.Data) {
				cell6.SetString(cell5.String() + "Произошла ошибка при копировании файла " + blocks[i].Block.FileName + ". ")
			}

		}

		return sh, nil
	}
}

func (er *ExelReporter) createDeployFolder() error {
	var folderName string
	if er.reportParams.FromDate.Equal(er.reportParams.ToDate) {
		folderName = "Сводка за " + er.reportParams.FromDate.Format("02.01.2006")
	} else {
		folderName = "Сводка за " + er.reportParams.FromDate.Format("02.01.2006") + "-" + er.reportParams.ToDate.Format("02.01.2006")

	}
	err := os.MkdirAll(filepath.Join(er.reportParams.DeployPath, folderName), 0750)
	if err != nil {
		return err
	}
	er.deployFolder = folderName
	return nil
}

func (er *ExelReporter) checkFilePaths() error {
	bf, err := os.Stat(er.reportParams.BlocksFolderPath)
	if err != nil {
		return err
	}
	if !bf.IsDir() {
		return errors.New(er.reportParams.BlocksFolderPath + " не является директорией")
	}
	dp, err := os.Stat(er.reportParams.DeployPath)
	if err != nil {
		return err
	}
	if !dp.IsDir() {
		return errors.New(er.reportParams.DeployPath + " не является директорией")
	}
	return nil
}
