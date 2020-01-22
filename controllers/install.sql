CREATE TABLE `blog`.`employee` (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID',
    `username` VARCHAR(30) NOT NULL COMMENT '用户名',
    `password` VARCHAR(30) NOT NULL COMMENT '密码',
    PRIMARY KEY (`id`)
) ENGINE = MyISAM;