package model

import "github.com/jinzhu/gorm"

// BuildQuery is
func BuildQuery(db *gorm.DB, table string, wheres []WhereQuery, order string) *gorm.DB {
	db = db.Table(table)

	if wheres != nil && len(wheres) > 0 {
		for _, where := range wheres {
			if where.Operator == "IS NULL" || where.Operator == "IS NOT NULL" {
				db = db.Where(where.Column + " " + where.Operator)
			} else {
				db = db.Where(where.Column+" "+where.Operator+" ?", where.Value)
			}
		}
	}

	if order == "" {
		order = "ASC"
	}

	return db.Order("id " + order).Limit(1000)
}
