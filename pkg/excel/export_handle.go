package excel

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/xuri/excelize/v2"
)

type Request struct {
	Data    []map[string]interface{} `json:"data"`
	Headers []string                 `json:"headers"`
}

func ExportHandler(app *fiber.Ctx) error {
	var request Request

	app.BodyParser(&request)

	f := excelize.NewFile()

	if len(request.Data) == 0 {
		return app.Status(fiber.StatusBadRequest).SendString("Nenhum dado encontrado para exportar")
	}

	if len(request.Headers) == 0 {
		return app.Status(fiber.StatusBadRequest).SendString("Nenhum header encontrado")
	}

	for i, header := range request.Headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue("Sheet1", cell, header)
	}

	for rowIndex, rowData := range request.Data {
		for colIndex, header := range request.Headers {
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
