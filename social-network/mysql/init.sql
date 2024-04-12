# 1. 创建数据库
CREATE DATABASE IF NOT EXISTS socialDB;

# 2. 使用创建的数据库
USE socialDB;

# 3. 创建 users 表
CREATE TABLE IF NOT EXISTS `users` (
    `id` BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `username` VARCHAR(255) NOT NULL,
    `password` VARCHAR(255) NOT NULL,
    `add_time` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `update_time` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

# 4. 创建 moments 表
CREATE TABLE IF NOT EXISTS `moments` (
    `id` BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id` BIGINT(20) UNSIGNED NOT NULL,
    `content` TEXT NOT NULL,
    `image_paths` TEXT,  -- 存储图片路径，可以是JSON格式以支持多张图片
    `add_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

# 5. 创建 likes 表
CREATE TABLE IF NOT EXISTS `likes` (
    `id` BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
    `moment_id` BIGINT(20) UNSIGNED NOT NULL,
    `user_id` BIGINT(20) UNSIGNED,
    `add_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `unique_like` (`moment_id`,`user_id`), -- 确保一个用户对同一条朋友圈只能点赞一次
    KEY `idx_user_id` (`user_id`),
    CONSTRAINT `fk_likes_moments` FOREIGN KEY (`moment_id`) REFERENCES `moments` (`id`) ON DELETE CASCADE,
    CONSTRAINT `fk_likes_users` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE SET NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;