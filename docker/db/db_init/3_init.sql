CREATE TABLE quetions (
    id int auto_increment not null primary key, 
    title VARCHAR(32), 
    content VARCHAR(256)
);

CREATE TABLE quetionIds (
    id int auto_increment not null primary key, 
    titleid int
);