CREATE TABLE quetions (
    id int auto_increment not null primary key, 
    title varchar(32), 
    content varchar(256),
    created_at TIMESTAMP
);

insert into book_db.posts (title, content) values ('hoge', 'hogehoge');
