package history

import (
	"encoding/json"

	"gorm.io/gorm"
)

type Logger struct {
	DB *gorm.DB
}

func NewLogger(db *gorm.DB) *Logger {
	return &Logger{DB: db}
}

func (l *Logger) writeLog(tx *gorm.DB, action string, model interface{}, old interface{}) error {
	table := tx.Statement.Table
	var modelID *int64

	// primary key olish
	if pk := tx.Statement.Schema.PrioritizedPrimaryField; pk != nil {
		if v, ok := pk.ValueOf(tx.Statement.Context, model); ok {
			if id, ok := v.(int64); ok {
				modelID = &id
			}
		}
	}

	var oldJSON, newJSON []byte
	var err error

	if old != nil {
		oldJSON, err = json.Marshal(old)
		if err != nil {
			return err
		}
	}

	if model != nil {
		newJSON, err = json.Marshal(model)
		if err != nil {
			return err
		}
	}

	h := History{
		Table:   &table,
		ModelID: modelID,
		Action:  &action,
		OldValue: oldJSON,
		NewValue: newJSON,
	}

	return tx.Create(&h).Error
}
