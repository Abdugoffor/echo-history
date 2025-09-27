package history

import (
	"gorm.io/gorm"
)

// RegisterHooks barcha callbacklarni ulamiz
func RegisterHooks(db *gorm.DB) {
	logger := NewLogger(db)

	// CREATE
	db.Callback().Create().After("gorm:create").Register("history:create", func(tx *gorm.DB) {
		if tx.Error == nil {
			_ = logger.writeLog(tx, "create", tx.Statement.Dest, nil)
		}
	})

	// UPDATE
	db.Callback().Update().After("gorm:update").Register("history:update", func(tx *gorm.DB) {
		if tx.Error == nil {
			_ = logger.writeLog(tx, "update", tx.Statement.Dest, tx.Statement.ReflectValue.Interface())
		}
	})

	// DELETE (soft delete)
	db.Callback().Delete().After("gorm:delete").Register("history:delete", func(tx *gorm.DB) {
		if tx.Error == nil {
			_ = logger.writeLog(tx, "delete", nil, tx.Statement.Dest)
		}
	})

	// RESTORE (gorm unscoped update qilib o‘zi chaqiriladi)
	db.Callback().Update().After("gorm:restore").Register("history:restore", func(tx *gorm.DB) {
		if tx.Error == nil {
			_ = logger.writeLog(tx, "restore", tx.Statement.Dest, nil)
		}
	})
}
