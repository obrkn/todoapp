CREATE DATABASE IF NOT EXISTS todoapp character SET utf8mb4 collate utf8mb4_bin;

USE todoapp;

SET FOREIGN_KEY_CHECKS=0;

DROP TABLE IF EXISTS tasks;
CREATE TABLE IF NOT EXISTS tasks (
  id INT UNSIGNED NOT NULL PRIMARY KEY auto_increment,
  content VARCHAR(255) NOT NULL,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) character set utf8mb4 collate utf8mb4_bin;