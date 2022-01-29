-- Удалить БД, если уже существует
DROP SCHEMA IF EXISTS `cashDB`;

-- Создает БД с чарсетом UTF-8
CREATE SCHEMA `cashDB` DEFAULT CHARACTER SET utf8;

-- Выбор созданной схемы как текущей активной
USE `cashDB`;

-- Создает таблицу
CREATE TABLE `cashDB`.`users` (
    `id` INT UNSIGNED NOT NULL,  -- нет автоинкремента т.к. изначально сущности создаются с заданным Id
    `value` DECIMAL(41,2) NOT NULL,
    PRIMARY KEY (`id`));
