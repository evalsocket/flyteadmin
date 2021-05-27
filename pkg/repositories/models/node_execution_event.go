package models

import (
	"time"
)

type NodeExecutionEvent struct {
	BaseModel
	NodeExecutionKey
	RequestID  string
	OccurredAt time.Time
	Phase      string `gorm:"primary_key" valid:"length(1|50)"`
}
