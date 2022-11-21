SHOW DATABASES;

CREATE DATABASE `transactionscript` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
CREATE DATABASE `domainmodel` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
CREATE DATABASE `eventsourceddomainmodel` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;

USE transactionscript;

CREATE TABLE `pets`(
	id VARCHAR(36) PRIMARY KEY,
	name VARCHAR(30) NOT NULL,
	age TINYINT NOT NULL,
	added_at DATETIME NOT NULL,
	is_sold BOOLEAN NOT NULL DEFAULT FALSE,
	sold_at DATETIME
)


USE domainmodel;

CREATE TABLE `pets`(
	id VARCHAR(36) PRIMARY KEY,
	name VARCHAR(30) NOT NULL,
	age TINYINT NOT NULL,
	added_at DATETIME NOT NULL,
	is_sold BOOLEAN NOT NULL DEFAULT FALSE,
	sold_at DATETIME,
	version INT UNSIGNED NOT NULL DEFAULT 0
)


USE eventsourceddomainmodel;

CREATE TABLE `pets`(
	id VARCHAR(36) PRIMARY KEY,
	aggr_id VARCHAR(36) NOT NULL,
	event_id INT UNSIGNED NOT NULL,
	event_type VARCHAR(255) NOT NULL,
	created_at DATETIME NOT NULL,
	event_body JSON NOT NULL,
	INDEX `idx_aggr_id` (`aggr_id`)
)

SELECT * FROM `pets` ORDER BY aggr_id, event_id

