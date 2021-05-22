package model

func Create() {
	var schemaUserInfo = `
	CREATE TABLE IF NOT EXISTS user_base_info (
		uid bigint(20) UNSIGNED AUTO_INCREMENT COMMENT '用户ID',
		username varchar(16) NOT NULL COMMENT '用户名',
		password varchar(100) NOT NULL COMMENT '密码',
		nickname varchar(16) NOT NULL COMMENT '昵称',
		mail varchar(50) NOT NULL COMMENT '邮箱',
		mail_verified bool NOT NULL DEFAULT false COMMENT '邮箱是否验证',
		join_time timestamp NOT NULL COMMENT '注册时间',
		account_state tinyint(4) NOT NULL DEFAULT 0 COMMENT '账号状态，0:正常 1:封禁',
		sign varchar(70) NOT NULL DEFAULT '' COMMENT '个人签名',
		gender tinyint(4) NOT NULL DEFAULT 0 COMMENT '性别，0:保密 1:男 2:女 3:其他',
		credit int(11) NOT NULL DEFAULT 1000 COMMENT '信用值',
		birthday date NOT NULL DEFAULT '1970-01-01' COMMENT '生日',
		province varchar(32) NOT NULL DEFAULT '' COMMENT '省份',
		city varchar(32) NOT NULL DEFAULT '' COMMENT '城市',
		created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
		updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
		PRIMARY KEY (uid),
		UNIQUE KEY (username),
		UNIQUE KEY (mail)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户基本信息表';`

	// var schemaUserPoint = `
	// CREATE TABLE IF NOT EXISTS user_point (
	//     uid bigint unsigned primary key auto_increment COMMENT '用户ID',
	//     created_at timestamp DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	// 	updated_at timestamp DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
	// 	is_deleted bool DEFAULT false COMMENT '是否被删除',
	//   	exp int NOT NULL DEFAULT 0 COMMENT '经验',
	//   	coins bigint NOT NULL DEFAULT 0 COMMENT '硬币数',
	//   	sign_time timestamp COMMENT '签到时间'
	// );`
	DB.MustExec(schemaUserInfo)
	// DB.MustExec(schemaUserPoint)
}
