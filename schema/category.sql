DROP TABLE IF EXISTS `af_category`;
CREATE TABLE `af_category`
(
    `id`          INT(11)                                 NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `cate_id`     INT(10) UNSIGNED                        NOT NULL COMMENT '分区ID',
    `cate_name`   VARCHAR(20) COLLATE utf8mb4_general_ci  NOT NULL COMMENT '分区名称',
    `description` VARCHAR(256) COLLATE utf8mb4_general_ci NOT NULL COMMENT '分区描述',
    `created_at`  TIMESTAMP                               NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`  TIMESTAMP                               NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_category_id` (`cate_id`),
    UNIQUE KEY `idx_category_name` (`cate_name`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci COMMENT '分区表';