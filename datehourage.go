package main
import "fmt"
import "time"

func main() {
	now := time.Now()
	birthday := time.Date(1999, time.April, 18, 0, 0, 0, 0, time.UTC)
	date := now.Format("2006-01-02")
	hour := now.Format("15")
	age := now.Year() - birthday.Year()
	fmt.Printf("Date: %s, Hour: %s, Age: %d\n", date, hour, age)
}