package model

func Create() {
	var schemaUserInfo = `
	CREATE TABLE IF NOT EXISTS user_info (
		uid bigint unsigned primary key auto_increment comment '用户ID',
		created_at timestamp default CURRENT_TIMESTAMP comment '创建时间',
		updated_at timestamp default CURRENT_TIMESTAMP comment '更新时间',
		is_deleted bool default false comment '是否被删除',
		username varchar(16) not null comment '用户名',
		password varchar(100) not null comment '密码',
		nickname varchar(16) not null comment '昵称',
		mail varchar(50) not null comment '邮箱',
		mail_verify bool not null default false comment '邮箱是否验证',
		status tinyint not null default 0 comment '用户状态',
		avatar varchar(255) not null default '' comment '用户头像',
		sign varchar(70) comment '个人签名',
		gender tinyint not null default 0 comment '性别',
		birthday varchar(10) comment '生日'
	);`

	var schemaUserPoint = `
	CREATE TABLE IF NOT EXISTS user_point (
	    uid bigint unsigned primary key auto_increment comment '用户ID',
	    created_at timestamp default CURRENT_TIMESTAMP comment '创建时间',
		updated_at timestamp default CURRENT_TIMESTAMP comment '更新时间',
		is_deleted bool default false comment '是否被删除',
	  	exp int not null default 0 comment '经验',
	  	coins bigint not null default 0 comment '硬币数',
	  	sign_time timestamp comment '签到时间'									    
	);`

	var schemaWords = `
	CREATE TABLE IF NOT EXISTS words (
  		wid bigint unsigned primary key auto_increment comment '文字ID',
  		created_at timestamp default CURRENT_TIMESTAMP comment '创建时间',
		updated_at timestamp default CURRENT_TIMESTAMP comment '更新时间',
		is_deleted bool default false comment '是否被删除',
		publisher bigint unsigned comment '发布者ID',
		status tinyint not null  comment '状态',
		content longtext comment '文字内容'
	);`

	DB.MustExec(schemaUserInfo)
	DB.MustExec(schemaUserPoint)
	DB.MustExec(schemaWords)
}
