-- Active: 1738724774868@@172.16.6.69@3306@icu
-- auto-generated definition
# 用户表
create table icu.users (
`id` int auto_increment comment 'Primary Key' primary key,
`name` varchar(255) null comment 'Name',
`avatar` varchar(255) null comment 'avatar',
`email` varchar(255) null comment 'Email',
`username` varchar(255) null comment 'Username',
`signature` varchar(255) null comment 'signature',
`password` varchar(255) null comment 'Password',
`create_at` datetime default CURRENT_TIMESTAMP null comment 'Create Time'
) comment 'user';
#文章
create table icu.articles (
    `id` int auto_increment comment 'Primary Key' primary key,
    `title` varchar(255) null comment 'title',
    `content` TEXT null comment 'content',
    `user_id` varchar(255) null comment 'userId',
    `create_at` datetime default CURRENT_TIMESTAMP null comment 'Create Time'
) comment 'article';

#标签表
create table icu.tags (
    `id` int auto_increment comment 'Primary Key' primary key,
    `tag` varchar(255) null comment 'tag',
    `user_id` varchar(255) null comment 'userId',
    `create_at` datetime default CURRENT_TIMESTAMP null comment 'Create Time'
) comment 'tags';

#标签文章多对多关联表
create table icu.article_tags (
    `id` int auto_increment comment 'Primary Key' primary key,
    `article_id` int null comment 'articleId',
    `tag_id` int null comment 'tagId',
    `create_at` datetime default CURRENT_TIMESTAMP null comment 'Create Time',
    `deleted_at` datetime DEFAULT NULL null comment 'deleted_at',
    `updated_at` datetime DEFAULT NULL null comment 'updated_at'
) comment 'article_tags';

ALTER TABLE icu.articles ADD COLUMN deleted_at DATETIME NULL;

ALTER TABLE icu.tags ADD COLUMN deleted_at DATETIME NULL;

ALTER TABLE icu.users ADD COLUMN deleted_at DATETIME NULL;

ALTER TABLE icu.articles ADD COLUMN updated_at DATETIME NULL;

ALTER TABLE icu.tags ADD COLUMN updated_at DATETIME NULL;

ALTER TABLE icu.users ADD COLUMN updated_at DATETIME NULL;