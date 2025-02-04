CREATE TABLE icu.users (
    id INT NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT 'Primary Key',
    name VARCHAR(255) COMMENT 'Name',
    email VARCHAR(255) COMMENT 'Email',
    username VARCHAR(255) COMMENT 'Username',
    password VARCHAR(255) COMMENT 'Password',
    create_time DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT 'Create Time'
) COMMENT 'user';