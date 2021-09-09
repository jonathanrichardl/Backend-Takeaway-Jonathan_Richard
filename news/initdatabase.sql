USE testers;
DROP TABLE tags;
DROP TABLE news;
CREATE TABLE news(
id integer PRIMARY KEY AUTO_INCREMENT,
title varchar(45),
topic varchar(25),
status varchar(25)
);
CREATE TABLE tags(
id integer PRIMARY KEY AUTO_INCREMENT,
tags varchar(255),
news_id integer,
FOREIGN KEY(news_id) REFERENCES news(id)
);
