package loader

import "fmt"

func CalculateMetrics(success, failed, total int, duration float64) {
	fmt.Println("Success Rate: %.2f%%\n", (float64(success)/float64(total))*100)
	fmt.Println("Failure Rate: %.2f%%\n", (float64(failed)/float64(total))*100)
	fmt.Println("Requests per second: %.2f\n", float64(total)/duration)
}
