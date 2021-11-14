DROP TABLE IF EXISTS `af_word`;
CREATE TABLE `af_word`
(
    `id`         BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `wid`        BIGINT(20) UNSIGNED NOT NULL COMMENT '文字ID',
    `aid`        INT(20) UNSIGNED    NOT NULL COMMENT '发布者ID',
    `cat_id`     INT(10) UNSIGNED    NOT NULL COMMENT '分区ID',
    `status`     TINYINT(4)          NOT NULL DEFAULT '1' COMMENT '文字状态',
    `title`      VARCHAR(48)         NOT NULL COLLATE utf8mb4_general_ci COMMENT '用户名',
    `content`    VARCHAR(1024)       NOT NULL COLLATE utf8mb4_general_ci COMMENT '文字内容',
    `likes`      INT(11) UNSIGNED    NOT NULL DEFAULT 0 COMMENT '点赞数量',
    `created_at` TIMESTAMP           NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` TIMESTAMP           NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_word_id` (`wid`),
    KEY `idx_author_id` (`aid`),
    KEY `idx_category_id` (`cat_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci COMMENT '文字表';