package history

import (
	"gorm.io/gorm"
)

// RegisterHooks barcha CRUD eventlarga hook qo‘shadi
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
			_ = logger.writeLog(tx, "update", tx.Statement.Dest, nil)
		}
	})

	// DELETE
	db.Callback().Delete().After("gorm:delete").Register("history:delete", func(tx *gorm.DB) {
		if tx.Error == nil {
			_ = logger.writeLog(tx, "delete", nil, tx.Statement.Dest)
		}
	})

	// RESTORE (gorm’da SoftDelete cancel qilinganda ishlaydi)
	db.Callback().Update().After("gorm:restore").Register("history:restore", func(tx *gorm.DB) {
		if tx.Error == nil {
			_ = logger.writeLog(tx, "restore", tx.Statement.Dest, nil)
		}
	})
}
