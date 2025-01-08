package reports

import (
	"encoding/json"
	"os"
)

type Report struct {
	Success int
	Failed  int
	Total   int
	Time    float64
}

func SaveJSONReport(report Report, filename string) {
	file, _ := os.Create(filename)
	defer file.Close()
	json.NewEncoder(file).Encode(report)
}
