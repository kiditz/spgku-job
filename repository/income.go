package repository

import (
	"github.com/kiditz/spgku-api/db"
	e "github.com/kiditz/spgku-api/entity"
	"github.com/kiditz/spgku-api/utils"
	"github.com/labstack/echo/v4"
)

// GetIncomes used to insert campaign into briefs database
func GetIncomes(c echo.Context, filter *e.IncomeFilter) []e.Income {
	if filter.Limit == 0 {
		filter.Limit = 10
	}
	incomes := []e.Income{}
	user := utils.GetUser(c)
	userID := uint(user["id"].(float64))
	query := db.DB.Where("user_id = ?", userID)
	if filter.StartDate != "" && filter.EndDate != "" {
		query = query.Where("incomes.created_at between to_date(?, 'YYYY-MM-DD') and to_date(?, 'YYYY-MM-DD')", filter.StartDate, filter.EndDate)
	}
	if filter.CanWithdrawal {
		query = query.Where("can_withdrawal = ?", filter.CanWithdrawal)
	}
	if filter.HasWithdraw {
		query = query.Where("has_withdraw = ?", filter.HasWithdraw)
	}
	if filter.Query != "" {
		query = query.Joins("JOIN briefs b ON b.id = brief_id")
		query = query.Where("order_id ilike ? OR b.title ilike ?", "%"+filter.Query+"%", "%"+filter.Query+"%")
	}
	query = query.Preload("Brief")
	query = query.Offset(filter.Offset).Limit(filter.Limit).Order("id desc")
	query = query.Find(&incomes)
	return incomes
}

//FindIncomeInfo godoc
func FindIncomeInfo(c echo.Context) e.IncomeInfo {
	user := utils.GetUser(c)
	userID := uint(user["id"].(float64))
	var info e.IncomeInfo
	query := db.DB.Table("incomes i")
	query = query.Where("i.user_id = ?", userID)
	query = query.Select(`
	(
		SELECT cast(sum((amount - (10.0 / 100.0) * amount)) as integer)
			FROM incomes 
		WHERE can_withdrawal = true and has_withdraw = false and user_id = i.user_id
	) as total_withdrawal,
	(
		SELECT cast(sum((amount - (10.0 / 100.0) * amount)) as integer)
			FROM incomes 
		WHERE user_id = i.user_id
	) as total
	
	`).Scan(&info)
	return info
}