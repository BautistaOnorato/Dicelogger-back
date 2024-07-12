package handler

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/proyecto-dnd/backend/internal/report"
)

type ReportHandler struct {
	service report.ReportGenerator
}

func NewReportHandler(service *report.ReportGenerator) *ReportHandler {
	return &ReportHandler{service: *service}
}

func (h *ReportHandler) HandlerGetSessionReport() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		report, err := h.service.GenerateEventSessionReport(id)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.Writer.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
		c.Writer.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=report-session-%d-%s.xlsx", id, time.Now().Format("2006-01-02 15:04:30")))
		c.Writer.Write(report.Bytes())
	}
}

func (h *ReportHandler) HandlerGetCharacterCampaignReport() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		report, err := h.service.GenerateCharactersByCampaignReport(id)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.Writer.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
		c.Writer.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=report-characters-campaign-%d-%s.xlsx", id, time.Now().Format("2006-01-02 15:04")))
		c.Writer.Write(report.Bytes())
	}
}