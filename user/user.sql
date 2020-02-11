// mysql
CREATE TABLE `user`(
    `user_id` char(32) NOT NULL COMMENT '用户id',
    `user_name` varchar(15) NOT NULL unique COMMENT '用户账号',
    `password` varchar(128) NOT NULL COMMENT '加密后的密码',
    `nick_name` varchar(15) NOT NULL unique COMMENT '用户昵称',
    `email` varchar(30),
    `remark` varchar(30),
    `is_active` bool NOT NULL default true,
    `is_superuser` bool NOT NULL default false,
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT '用户信息表';


// postgresql
CREATE TABLE "user"(
    "user_id" char(32) primary key NOT NULL ,
    "user_name" varchar(15) NOT NULL unique,
    "password" varchar(128) NOT NULL,
    "nick_name" varchar(15) NOT NULL unique,
    "email" varchar(30),
    "remark" varchar(30),
    "is_active" boolean NOT NULL default true,
    "is_superuser" boolean NOT NULL default false,
    "create_time" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "update_time" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);


