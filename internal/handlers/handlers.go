package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/sounishnath003/pdf-generator-go/internal/models"
	"github.com/sounishnath003/pdf-generator-go/internal/service"
)

func HandleIndexPage(ctx echo.Context) error {
	data := models.IndexPageModel{
		Name:   "index",
		Date:   time.Now().Format("January-2006"),
		BadAss: "Batch-Sounish-Nath",
	}
	return ctx.Render(http.StatusOK, data.Name, &data)
}

func HandleReportGeneration(ctx echo.Context) error {
	data := service.GenerateLatestSalesReport()
	log.Println("latest.data", data)
	return ctx.Render(http.StatusOK, data.Name, &data)
}
