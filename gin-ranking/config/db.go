package config

const (
	Mysqldb = "root:123456@tcp(localhost:3306)/ranking?charset=utf8mb4&parseTime=true&loc=Local"
)

/*
创建数据库
CREATE DATABASE ranking;
创建用户表
CREATE TABLE users (
	id INT AUTO_INCREMENT PRIMARY KEY,
	username VARCHAR(255) NOT NULL,
	password VARCHAR(255) NOT NULL,
	add_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

创建活动表
CREATE TABLE activities (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    add_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
创建活动参加者表
CREATE TABLE players (
    id INT AUTO_INCREMENT PRIMARY KEY,
    activity_id INT NOT NULL,
    contestant_number VARCHAR(255) NOT NULL,
    nickname VARCHAR(255) NOT NULL,
    description TEXT,
    image VARCHAR(255),
    score DECIMAL(10, 2),
	add_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (activity_id) REFERENCES activities(id)
);
创建投票表
CREATE TABLE votes (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    player_id INT NOT NULL,
	add_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

插入样例数据：
INSERT INTO users (username, password, add_time, update_time) VALUES
('Alice', 'password123', NOW(), NOW()),
('Bob', 'password456', NOW(), NOW()),
('Charlie', 'password789', NOW(), NOW());


INSERT INTO activities (name, add_time) VALUES
('Coding Contest', NOW()),
('Art Competition', NOW()),
('Math Olympiad', NOW());


INSERT INTO players (activity_id, contestant_number, nickname, description, image, score, add_time, update_time) VALUES
(1, '001', 'CoderA', 'A passionate coder', 'image_url_1', 95.5, NOW(), NOW()),
(1, '002', 'CoderB', 'Loves challenges', 'image_url_2', 88.4, NOW(), NOW()),
(1, '003', 'CoderC', 'Never stops learning', 'image_url_3', 92.3, NOW(), NOW()),
(2, '001', 'ArtistA', 'Abstract art lover', 'image_url_4', 78.5, NOW(), NOW()),
(2, '002', 'ArtistB', 'Watercolor painter', 'image_url_5', 82.1, NOW(), NOW()),
(2, '003', 'ArtistC', 'Digital artist', 'image_url_6', 89.9, NOW(), NOW()),
(3, '001', 'MathWizA', 'Geometry guru', 'image_url_7', 97.0, NOW(), NOW()),
(3, '002', 'MathWizB', 'Algebra aficionado', 'image_url_8', 93.2, NOW(), NOW()),
(3, '003', 'MathWizC', 'Statistics specialist', 'image_url_9', 88.7, NOW(), NOW());


*/
