package answernumber5

import (
	"log"
	"strings"
	"time"

	"gitee.com/go-package/carbon"
	"github.com/dustin/go-humanize"
)

func FormatRupiah(amount float64) string {
	humanizeValue := humanize.CommafWithDigits(amount, 0)
	stringValue := strings.Replace(humanizeValue, ",", ".", -1)
	return stringValue
}

func GetDateDifferent(targetTime string) string {
	result := carbon.Parse(targetTime).DiffForHumans()
	return result
}

func main() {

	log.Println("library format rupiah : ", FormatRupiah(100000))

	log.Println("date from library carbon : ", GetDateDifferent(time.Now().AddDate(-10, 0, 0).Format("2006-01-02 15:04:05")))

}
