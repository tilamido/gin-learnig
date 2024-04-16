// config/mysql.go

package config

const (
	MySQLDSN_DEFAULT = "root:123456@tcp(localhost:39001)/socialDB?charset=utf8mb4&parseTime=true&loc=Local"
)

var (
	MySQLDSN string
)

func init() {
	MySQLDSN = getEnv("DB_CONNECTION_STRING", MySQLDSN_DEFAULT)
}

/*
创建数据库
CREATE DATABASE socialDB;

使用数据库
use socialDB;

创建用户表
CREATE TABLE users (
	id INT AUTO_INCREMENT PRIMARY KEY,
	username VARCHAR(255) NOT NULL,
	password VARCHAR(255) NOT NULL,
	add_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

插入样例数据：
INSERT INTO users (username, password, add_time, update_time) VALUES
('Alice', 'password123', NOW(), NOW()),
('Bob', 'password456', NOW(), NOW()),
('Charlie', 'password789', NOW(), NOW());

*/
