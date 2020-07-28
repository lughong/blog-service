drop database if exists `blog`;
create database `blog` default charset utf8mb4 collate utf8mb4_general_ci;

use blog;

drop table if exists `tag`;
create table `tag`(
	`id` int(10) unsigned not null auto_increment,
	`name` varchar(100) not null default '' comment'标签名称',
	`created_on` timestamp not null comment '创建时间',
	`created_by` varchar(100) not null default '' comment '创建人',
	`modified_on` timestamp not null comment '修改时间',
	`modified_by` varchar(100) not null default '' comment '修改人',
	`is_del` tinyint(1) unsigned not null default 0 comment '是否已删除：0.未删除，1.已删除',
    `deleted_on` timestamp not null comment '删除时间',
	`state` tinyint(1) unsigned not null default 0 comment '状态：0.禁止，1.启用',
	PRIMARY KEY(`id`)
) engine=Innodb default charset utf8mb4 collate utf8mb4_general_ci comment '标签表';

drop table if exists `article`;
create table `article`(
	`id` int(10) unsigned not null auto_increment,
	`title` varchar(100) not null default '' comment '文章标题',
	`desc` varchar(100) not null default '' comment '文章描述',
	`content` text not null comment '文章内容',
	`created_on` timestamp not null comment '创建时间',
    `created_by` varchar(100) not null default '' comment '创建人',
    `modified_on` timestamp not null comment '修改时间',
    `modified_by` varchar(100) not null default '' comment '修改人',
    `is_del` tinyint(1) not null default 0 comment '是否已删除：0.未删除，1.已删除',
    `deleted_on` timestamp not null comment '删除时间',
    `state` tinyint(1) unsigned not null default 0 comment '状态：0.禁止，1.启用',
	PRIMARY KEY(`id`)
) engine=Innodb default charset utf8mb4 collate utf8mb4_general_ci comment '文章表';

drop table if exists `article_tag`;
create table `article_tag`(
	`id` int(10) unsigned not null auto_increment,
	`tag_id` int(10) unsigned not null default 0,
	`article_id` int(10) unsigned not null default 0,
	`created_on` timestamp not null comment '创建时间',
    `created_by` varchar(100) not null default '' comment '创建人',
    `modified_on` timestamp not null comment '修改时间',
    `modified_by` varchar(100) not null default '' comment '修改人',
    `is_del` tinyint(1) not null default 0 comment '是否已删除：0.未删除，1.已删除',
    `deleted_on` timestamp not null comment '删除时间',
	PRIMARY KEY(`id`)
) engine=Innodb default charset utf8mb4 collate utf8mb4_general_ci comment '文章标签关联表';
