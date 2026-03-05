package utils

import (
	"OperationAndMonitoring/mysql"
	"OperationAndMonitoring/mysql/db"
	"errors"
	"fmt"
	"path/filepath"
	"runtime"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

// Create
func Create(value interface{}) error {
	err := db.DB.Create(value).Error
	logDBError("Create", err)
	return err
}

// Save
func Save(value interface{}) error {
	err := db.DB.Save(value).Error
	logDBError("Save", err)
	return err
}

// Updates
func Updates(where interface{}, value interface{}) error {
	err := db.DB.Model(where).Updates(value).Error
	logDBError("Updates", err)
	return err
}

// Delete
func DeleteByModel(model interface{}) (count int64, err error) {
	tx := db.DB.Delete(model)
	err = tx.Error
	if err != nil {
		logDBError("DeleteByModel", err)
		return
	}
	count = tx.RowsAffected
	return
}

// Delete
func DeleteByWhere(model, where interface{}) (count int64, err error) {
	tx := db.DB.Where(where).Delete(model)
	err = tx.Error
	if err != nil {
		logDBError("DeleteByWhere", err)
		return
	}
	count = tx.RowsAffected
	return
}

// Delete
func DeleteByID(model interface{}, id uint64) (count int64, err error) {
	tx := db.DB.Where("id=?", id).Delete(model)
	err = tx.Error
	if err != nil {
		logDBError("DeleteByID", err)
		return
	}
	count = tx.RowsAffected
	return
}

// Delete
func DeleteByIDS(model interface{}, ids []uint64) (count int64, err error) {
	tx := db.DB.Where("id in (?)", ids).Delete(model)
	err = tx.Error
	if err != nil {
		logDBError("DeleteByIDS", err)
		return
	}
	count = tx.RowsAffected
	return
}

// First
func FirstByID(out interface{}, id int) (notFound bool, err error) {
	err = db.DB.First(out, id).Error
	notFound = isRecordNotFound(err)
	logDBError("FirstByID", err)
	return
}

// First
func First(where interface{}, out interface{}) (notFound bool, err error) {
	err = db.DB.Where(where).First(out).Error
	notFound = isRecordNotFound(err)
	logDBError("First", err)
	return
}

// Find
func Find(where interface{}, out interface{}, whereOrder ...mysql.PageWhereOrder) error {
	tx := applyPageWhereOrder(db.DB.Where(where), whereOrder...)
	err := tx.Find(out).Error
	logDBError("Find", err)
	return err
}

// Scan
func Scan(model, where interface{}, out interface{}) (notFound bool, err error) {
	err = db.DB.Model(model).Where(where).Scan(out).Error
	notFound = isRecordNotFound(err)
	logDBError("Scan", err)
	return
}

// ScanList
func ScanList(model, where interface{}, out interface{}, orders ...string) error {
	tx := db.DB.Model(model).Where(where)
	if len(orders) > 0 {
		for _, order := range orders {
			tx = tx.Order(order)
		}
	}
	err := tx.Scan(out).Error
	logDBError("ScanList", err)
	return err
}

// GetPage
func GetPage(model, where interface{}, out interface{}, pageIndex, pageSize int, totalCount *int64, whereOrder ...mysql.PageWhereOrder) error {
	tx := applyPageWhereOrder(db.DB.Model(model).Where(where), whereOrder...)
	err := tx.Count(totalCount).Error
	if err != nil {
		logDBError("GetPage.Count", err)
		return err
	}
	if *totalCount == 0 {
		return nil
	}
	err = tx.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(out).Error
	logDBError("GetPage.Find", err)
	return err
}

// PluckList
func PluckList(model, where interface{}, out interface{}, fieldName string) error {
	err := db.DB.Model(model).Where(where).Pluck(fieldName, out).Error
	logDBError("PluckList", err)
	return err
}

// PluckList
func Test(model, out interface{}, preload, association string) error {
	err := db.DB.Model(model).Association(association).Find(out)
	logDBError("Test", err)
	return err
}

func applyPageWhereOrder(tx *gorm.DB, whereOrder ...mysql.PageWhereOrder) *gorm.DB {
	if len(whereOrder) == 0 {
		return tx
	}
	for _, wo := range whereOrder {
		if wo.Order != "" {
			tx = tx.Order(wo.Order)
		}
		if wo.Where != "" {
			tx = tx.Where(wo.Where, wo.Value...)
		}
	}
	return tx
}

func isRecordNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}

func logDBError(op string, err error) {
	if err == nil || isRecordNotFound(err) {
		return
	}
	caller := caller(3)
	zap.L().Error("mysql operation failed",
		zap.String("operation", op),
		zap.String("caller", caller),
		zap.Error(err),
	)
}

func caller(skip int) string {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		return "unknown"
	}
	funcName := "unknown"
	if fn := runtime.FuncForPC(pc); fn != nil {
		funcName = fn.Name()
	}
	return fmt.Sprintf("%s:%d %s", filepath.Base(file), line, funcName)
}
