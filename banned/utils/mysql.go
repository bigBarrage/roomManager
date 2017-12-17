package utils

import "database/sql"
import (
	_ "github.com/go-sql-driver/mysql"
	"errors"
	"github.com/bigBarrage/roomManager/banned"
)

var mysqlConn *sql.DB

var ERROR_MYSQL_CONN_IS_NOT_WORK = errors.New("mysql conn is not work")

func SetMySQLConn(db *sql.DB){
	mysqlConn = db
}

func LoadWordList()error{
	if err := mysqlConn.Ping();err != nil{
		return err
	}

	rows,err := mysqlConn.Query("SELECT words FROM words_list WHERE status =  1")

	if err != nil{
		return err
	}
	wordsList := make([]string,0,8)
	for rows.Next(){
		word := ""
		rows.Scan(&word)
		wordsList = append(wordsList,word)
	}
	banned.SetWordList(wordsList)
}

func LoadUserList()error{
	if err := mysqlConn.Ping();err != nil{
		return err
	}

	rows,err := mysqlConn.Query("SELECT userID FROM user_id_list WHERE status =  1")

	if err != nil{
		return err
	}
	userIdList := make([]string,0,8)
	for rows.Next(){
		id := ""
		rows.Scan(&ip)
		userIdList = append(userIdList,id)
	}
	banned.SetUserIDs(userIdList)
}

func LoadIpList()error{
	if err := mysqlConn.Ping();err != nil{
		return err
	}

	rows,err := mysqlConn.Query("SELECT ip FROM ip_list WHERE status =  1")

	if err != nil{
		return err
	}
	ipList := make([]string,0,8)
	for rows.Next(){
		ip := ""
		rows.Scan(&ip)
		ipList = append(ipList,ip)
	}
	banned.SetIpList(ipList)
}

func IsBannedWords(msg string) bool {
	return banned.IsBannedWords(msg)
}

func IsBannedUserID(userID string) bool {
	return banned.IsBannedUserID(userID)
}

func IsBannedIP(ip string) bool {
	return banned.IsBannedIP(ip)
}