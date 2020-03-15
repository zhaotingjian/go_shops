CREATE TABLE `test`.`employees` (
    `user_id` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户id',
    `user_name` VARCHAR(28) NOT NULL COMMENT '用户名',
    `pass_word` VARCHAR(28) NOT NULL COMMENT '密码',
    `true_name` VARCHAR(28) NOT NULL COMMENT '真实姓名',
    `sex` TINYINT NOT NULL COMMENT '性别：1：男；2：女',
    `mobile` VARCHAR(20) NOT NULL COMMENT '手机号',
    `card` VARCHAR(18) NOT NULL COMMENT '身份证号',
    `create_time` INT NOT NULL COMMENT '创建时间',
    `age` TINYINT UNSIGNED NOT NULL COMMENT '年龄',
    `login_time` INT NOT NULL COMMENT '最近一次登入时间',
    `nick_name` INT NOT NULL COMMENT '昵称',
    `edit_time` INT NOT NULL COMMENT '编辑时间',
    `status` TINYINT NOT NULL COMMENT '是否禁用：1-正常；2-禁用',
    `is_test` TINYINT NOT NULL COMMENT '是否是测试账号：1-是；2-否',
    PRIMARY KEY (`user_id`)
) ENGINE = InnoDB CHARSET = utf8 COLLATE utf8_general_ci COMMENT = '员工表';

CREATE TABLE `test`.`novel` (
    `novel_id` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '小说主键',
    `title` VARCHAR(58) NOT NULL COMMENT '标题',
    `author` VARCHAR(28) NOT NULL COMMENT '作者',
    `create_time` INT NOT NULL COMMENT '创建时间',
    `edit_time` INT NOT NULL COMMENT '编辑时间',
    `des` VARCHAR(258) NOT NULL COMMENT '简介',
    `status` TINYINT NOT NULL COMMENT '前端是否显示：1-显示；2不显示',
    `is_del` TINYINT NOT NULL COMMENT '是否删除：1-正常；2-删除',
    PRIMARY KEY (`novel_id`)
) ENGINE = InnoDB COMMENT = '小说';