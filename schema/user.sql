DROP TABLE IF EXISTS `af_user`;
CREATE TABLE `af_user`
(
    `uid`        INT(20) UNSIGNED   NOT NULL AUTO_INCREMENT COMMENT '用户ID',
    `username`   VARCHAR(24) UNIQUE NOT NULL COMMENT '用户名',
    `password`   CHAR(60)           NOT NULL COMMENT '密码密文',
    `nickname`   VARCHAR(15)        NOT NULL COMMENT '昵称',
    `email`      VARCHAR(50) UNIQUE NOT NULL COMMENT '邮箱',
    `avatar`     VARCHAR(100)       NOT NULL DEFAULT '' COMMENT '头像存储地址',
    `state`      TINYINT            NOT NULL DEFAULT 0 COMMENT '账号状态',
    `created_at` TIMESTAMP          NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` TIMESTAMP          NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`uid`)
) ENGINE = InnoDB COMMENT '用户表';

DROP TABLE IF EXISTS `af_user_info`;
CREATE TABLE `af_user_info`
(
    `uid`        INT(20) UNSIGNED NOT NULL COMMENT '用户ID',
    `gender`     TINYINT COMMENT '社会性别',
    `sign`       VARCHAR(100) COMMENT '个人签名',
    `birthday`   DATE COMMENT '生日',
    `province`   VARCHAR(20) COMMENT '省份',
    `city`       VARCHAR(20) COMMENT '城市',
    `created_at` TIMESTAMP        NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` TIMESTAMP        NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`uid`)
) ENGINE = InnoDB COMMENT '用户信息表';