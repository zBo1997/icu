-- Active: 1738724774868@@172.16.6.69@3306@icu
-- auto-generated definition
create table icu.users (
    id int auto_increment comment 'Primary Key' primary key,
    name varchar(255) null comment 'Name',
    avatar varchar(255) null comment 'avatar',
    email varchar(255) null comment 'Email',
    username varchar(255) null comment 'Username',
    signature varchar(255) null comment 'signature',
    password varchar(255) null comment 'Password',
    create_time datetime default CURRENT_TIMESTAMP null comment 'Create Time'
) comment 'user';