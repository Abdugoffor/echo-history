package history

import (
	"encoding/json"
	"fmt"
	"reflect"

	"gorm.io/gorm"
)

type Logger struct {
	DB *gorm.DB
}

func NewLogger(db *gorm.DB) *Logger {
	return &Logger{DB: db}
}

func (l *Logger) writeLog(tx *gorm.DB, action string, newModel interface{}, oldModel interface{}) error {
	table := tx.Statement.Table
	modelID := getPrimaryKey(newModel)

	var oldJSON, newJSON []byte
	var err error

	if oldModel != nil {
		oldJSON, err = json.Marshal(oldModel)
		if err != nil {
			return err
		}
	}

	if newModel != nil {
		newJSON, err = json.Marshal(newModel)
		if err != nil {
			return err
		}
	}

	act := action
	h := History{
		Table:    &table,
		ModelID:  modelID,
		Action:   &act,
		OldValue: oldJSON,
		NewValue: newJSON,
	}

	return tx.Create(&h).Error
}

// primary key olish
func getPrimaryKey(model interface{}) *int64 {
	if model == nil {
		return nil
	}

	v := reflect.ValueOf(model)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	field := v.FieldByName("ID")
	if !field.IsValid() {
		return nil
	}

	switch field.Kind() {
	case reflect.Int, reflect.Int64, reflect.Int32:
		id := field.Int()
		return &id
	default:
		return nil
	}
}
