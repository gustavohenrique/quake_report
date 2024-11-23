package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"quake_report/src/adapters/converters"
	"quake_report/src/adapters/dto"
	"quake_report/src/domain/services"
)

func main() {
	var logFile = os.Getenv("LOG_FILE")
	if logFile == "" {
		logFile = "qgames.log"
	}
	logContent, err := os.ReadFile(logFile)
	if err != nil {
		log.Fatalln("[ERROR]", err)
	}
	games := services.NewGameService().Parse(logContent)
	report := services.NewReportService().Generate(games)
	presentation := converters.NewReportConverter().FromModelToPresentation(report)
	printJSON(presentation)
}

func printJSON(presentation dto.ReportPresentation) {
	jsonReport, err := json.MarshalIndent(presentation, "", "  ")
	if err != nil {
		log.Fatalln("[ERROR] JSON marshal:", err)
		return
	}
	fmt.Println(string(jsonReport))
}
