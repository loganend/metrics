SET NAMES utf8 COLLATE 'utf8_unicode_ci';
SET foreign_key_checks = 1;
SET time_zone = '+00:00';
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';
SET default_storage_engine = InnoDB;
SET CHARACTER SET utf8;

DROP DATABASE IF EXISTS metrics;

CREATE DATABASE metrics DEFAULT CHARSET = utf8 COLLATE = utf8_unicode_ci;
USE metrics;


CREATE TABLE user (
    id INT(10) UNSIGNED NOT NULL,
    
    age INT(4) NOT NULL,
    sex VARCHAR(10) NOT NULL,

    PRIMARY KEY (id)
);

CREATE INDEX user_index
    on user (id);


CREATE TABLE stat (
    id INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    
    user_id INT(10) UNSIGNED NOT NULL,

    action VARCHAR(25) NOT NULL,
    datetime TIMESTAMP NOT NULL,
    
    CONSTRAINT `f_note_user` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    
    PRIMARY KEY (id)
);

CREATE INDEX stat_index
    on stat (datetime, action);