package lib

import (
	"github.com/google/uuid"
	"time"
)

const DbTLayout = "2006-01-02 15:04:05"
const TimeWhenDelete = 72 * time.Hour
const ShowCount = 3

func GetUUID() string {
	return uuid.New().String()
}
