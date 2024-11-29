package excel

import (
	"fmt"
	"sort"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/xuri/excelize/v2"
)

func Export(app *fiber.Ctx) error {
	var data []map[string]interface{}

	fmt.Println(data)

	app.BodyParser(&data)
	fmt.Println(data)

	f := excelize.NewFile()

	if len(data) == 0 {
		return app.Status(fiber.StatusBadRequest).SendString("Nenhum dado encontrado para exportar")
	}

	headers := make([]string, 0, len(data[0]))
	for key := range data[0] {
		headers = append(headers, key)
	}

	sort.Sort(sort.StringSlice(headers))

	for i, header := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue("Sheet1", cell, header)
	}

	for rowIndex, rowData := range data {
		for colIndex, header := range headers {
			cell, _ := excelize.CoordinatesToCellName(colIndex+1, rowIndex+2)
			value := rowData[header]
			f.SetCellValue("Sheet1", cell, value)
		}
	}

	filename := uuid.New().String()

	if err := f.SaveAs(fmt.Sprintf("./public/%s.xlsx", filename)); err != nil {
		return err
	}

	return app.Status(200).JSON(fmt.Sprintf("%s/%s.xlsx", app.BaseURL(), filename))
}
