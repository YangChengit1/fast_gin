package config

import (
	"fmt"
	"github.com/glebarez/sqlite"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBMode string

const (
	DBMysqlMode  = "mysql"
	DBPgsqlMode  = "pgsql"
	DBSqliteMode = "sqlite"
)

type DB struct {
	Mode     DBMode `yaml:"mode"` // 模式 mysql pgsql sqlite ""(“”表示不连接)
	DBName   string `yaml:"db_name"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func (db DB) DSN() gorm.Dialector {
	switch db.Mode {
	case DBMysqlMode:
		dsn := fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			db.User,
			db.Password,
			db.Host,
			db.Port,
			db.DBName,
		)
		return mysql.Open(dsn)
	case DBPgsqlMode:
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
			db.Host,
			db.User,
			db.Password,
			db.DBName,
			db.Port,
		)
		return postgres.Open(dsn)
	case DBSqliteMode:
		return sqlite.Open(db.DBName)
	default:
		logrus.Warnf("未配置数据库连接")
		return nil
	}
}
