package default_db

import (
	"github.com/izacgaldino23/products-api/config/database"
	"github.com/izacgaldino23/products-api/domain"
	"github.com/izacgaldino23/products-api/oops"
	"github.com/izacgaldino23/products-api/utils"
)

func Insert(object interface{}, tx *database.Transaction) (id int64, err error) {
	valueMap, err := utils.ParseStructToMap(object)
	if err != nil {
		return
	}

	if err = tx.Builder.
		Insert(domain.GetTableName(object)).
		SetMap(valueMap).
		Suffix("RETURNING id").
		Scan(&id); err != nil {
		return id, oops.Err(err)
	}

	return
}

func Update(object interface{}, id int64, tx *database.Transaction) (err error) {
	valueMap, err := utils.ParseStructToMap(object)
	if err != nil {
		return
	}

	if err = tx.Builder.
		Update(domain.GetTableName(object)).
		SetMap(valueMap).
		Where("id = ?", id).
		Suffix("RETURNING id").
		Scan(new(int64)); err != nil {
		return oops.Err(err)
	}

	return
}

func Delete(object interface{}, id int64, tx *database.Transaction) (err error) {
	if err = tx.Builder.
		Update(domain.GetTableName(object)).
		Set("removed_at", "NOW()").
		Where("id = ?", id).
		Suffix("RETURNING id").
		Scan(new(int64)); err != nil {
		return oops.Err(err)
	}

	return
}
