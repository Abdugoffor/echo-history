package history

import (
	"time"

	"gorm.io/datatypes"
)

type History struct {
	ID        int64          `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID    *int64         `json:"user_id" gorm:"index"`
	Table     *string        `json:"table" gorm:"size:255;index"`
	ModelID   *int64         `json:"model_id" gorm:"index"`
	IPAddress *string        `json:"ip_address" gorm:"size:50"`
	API       *string        `json:"api" gorm:"size:255"`
	OldValue  datatypes.JSON `json:"old_value" gorm:"type:jsonb"`
	NewValue  datatypes.JSON `json:"new_value" gorm:"type:jsonb"`
	Action    *string        `json:"action" gorm:"size:50"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
}

func (History) TableName() string {
	return "history"
}
