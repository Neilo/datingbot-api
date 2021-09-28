package health

import (
	"github.com/brotherhood228/dating-bot-api/pkg/metric"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

//Check for service
//check connection to db
func Check(c echo.Context) error {
	//todo добавить проверку коннешинов к базе
	now := time.Now()
	upDuration := now.Sub(upTime).Round(time.Second).Seconds()
	h := health{
		Status: statusOK,
		Error:  nil,
		UpTime: &upDuration,
	}
	return c.JSON(http.StatusOK, h)
}

type health struct {
	Status string   `json:"status"`
	Error  []string `json:"error,omitempty"`
	UpTime *float64 `json:"up_time"`
}

const (
	statusOK  = "ok"
	statusBad = "bad"
)

var upTime time.Time

//InitUpTime count up time seconds
func InitUpTime() {
	upTime = time.Now()

	for {
		c.Inc()
		time.Sleep(5 * time.Minute)
	}
}

var c = metric.MustRegisterCounter("test", "test")
