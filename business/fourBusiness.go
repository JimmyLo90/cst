package business

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/zhyq132/cst/config"
)

//SellAskCount 购车询价统计
func SellAskCount(areaId int) int {
	var i int
	dbDsn := config.GetDbDSN()

	dbConn, err := sql.Open("mysql", dbDsn)
	if err != nil {
		return i
	}
	defer dbConn.Close()

	sql := " SELECT COUNT(*) FROM `4s_sell_ask` WHERE `area_id`=? AND `isdel`=0 AND `status`=0 "

	err = dbConn.QueryRow(sql, areaId).Scan(&i)
	if err != nil {
		return i
	}

	return i
}

//SellPromiseCount 购车询价统计
func SellPromiseCount(areaId int) int {
	var i int
	dbDsn := config.GetDbDSN()

	dbConn, err := sql.Open("mysql", dbDsn)
	if err != nil {
		return i
	}
	defer dbConn.Close()

	sql := " SELECT COUNT(*) FROM `4s_sell_promise` WHERE `area_id`=? AND `isdel`=0 AND `status`=0 "

	err = dbConn.QueryRow(sql, areaId).Scan(&i)
	if err != nil {
		return i
	}

	return i
}

//YangxiuCount 购车询价统计
func YangxiuCount(areaId int) int {
	var i int
	dbDsn := config.GetDbDSN()

	dbConn, err := sql.Open("mysql", dbDsn)
	if err != nil {
		return i
	}
	defer dbConn.Close()

	sql := " SELECT COUNT(*) FROM `4s_yangxiu` WHERE `a_areaId`=? AND `isDel`=0 AND `is_read`=0 "

	err = dbConn.QueryRow(sql, areaId).Scan(&i)
	if err != nil {
		return i
	}

	return i
}

//XubaoCount 购车询价统计
func XubaoCount(areaId int) int {
	var i int
	dbDsn := config.GetDbDSN()

	dbConn, err := sql.Open("mysql", dbDsn)
	if err != nil {
		return i
	}
	defer dbConn.Close()

	sql := " SELECT COUNT(*) FROM `4s_xubao` WHERE `a_areaId`=? AND `4s_xubao`.`isDel`=0 AND `is_read`=0 "

	err = dbConn.QueryRow(sql, areaId).Scan(&i)
	if err != nil {
		return i
	}

	return i
}

//SupportCount 购车询价统计
func SupportCount(areaId int) int {
	var i int
	dbDsn := config.GetDbDSN()

	dbConn, err := sql.Open("mysql", dbDsn)
	if err != nil {
		return i
	}
	defer dbConn.Close()

	sql := " SELECT COUNT(*) FROM `4s_support_record` WHERE `area_id`=? AND `4s_support_record`.`isDel`=0 AND `status`=0 "

	err = dbConn.QueryRow(sql, areaId).Scan(&i)
	if err != nil {
		return i
	}

	return i
}
