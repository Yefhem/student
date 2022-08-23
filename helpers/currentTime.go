package helpers

import (
	"time"

	"github.com/Yefhem/student-syllabus/model"
)

func CurrentTime() model.CurrentTime {

	currentTime := model.CurrentTime{
		Day:   time.Now().Day(),
		Month: int(time.Now().Month()),
		Year:  time.Now().Year(),
		Hour:  time.Now().Hour(),
	}

	return currentTime

}
