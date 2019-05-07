package model

import (
	"book/db"
	"book/util"
	"database/sql"
	"fmt"
)

type LxUser struct {
	Id         int
	Username   string
	Userpwd    string
	CreateTime int
	LoginIp    string
	LoginTime  int

	CreateTimeF string
	LoginTimeF  string
}

type LxToken struct {
	id     int
	Token  string
	Status int
	UserId int
}

type LxArticle struct {
	Id          int
	UserId      int
	Title       string
	Content     string
	CreateTime  int
	ReadCount   int
	CreateTimeF string
	User        LxUser
}

func GetUser(name string) LxUser {
	rows := db.Query("select * FROM lx_user where username ='" + name + "' order by id desc limit 1")

	var user LxUser;
	if rows.Next() {
		user = parserUser(rows)
	}
	rows.Close()
	return user
}

func GetUserById(id int) LxUser {
	rows := db.Query(fmt.Sprintf("select * FROM lx_user where id = %d", id))

	var user LxUser;
	if rows.Next() {
		user = parserUser(rows)
	}

	rows.Close()

	return user
}

func GetUsers(sql string) []LxUser {
	rows := db.Query(sql)

	var users []LxUser;
	for rows.Next() {
		user := parserUser(rows)
		users = append(users, user)
	}
	rows.Close()
	return users
}

func InsertUser(user LxUser) int64 {
	//INSERT INTO `lx_user` (`id`, `username`, `userpwd`, `CreateTime`, `loginIp`, `LoginTime`) VALUES (NULL, 'admin', MD5('admin'), UNIX_TIMESTAMP(), NULL, '0');
	str := "INSERT INTO `lx_user` (`id`, `username`, `userpwd`, `createTime`) VALUES (NULL, '%s', '%s', UNIX_TIMESTAMP())"
	return db.Insert(fmt.Sprintf(str, user.Username, util.Md5(user.Userpwd)))
}

func UpdateUser(user LxUser) int64 {
	//UPDATE `lx_user` SET  `loginIp` =  '127.0.0.1' WHERE  `lx_user`.`id` =1;
	str := "UPDATE `lx_user` SET  `loginIp` = '%s' , loginTime= '%d' WHERE `id` = %d"
	return db.Update(fmt.Sprintf(str, user.LoginIp, user.LoginTime, user.Id))
}

func parserUser(rows *sql.Rows) LxUser {
	var user LxUser;
	err := rows.Scan(&user.Id, &user.Username, &user.Userpwd, &user.CreateTime, &user.LoginIp, &user.LoginTime)
	util.ShowError(err)

	user.CreateTimeF = util.TimeFormat(user.CreateTime)
	user.LoginTimeF = util.TimeFormat(user.LoginTime)
	return user
}

func GetToken(t string) LxToken {
	rows := db.Query("select * FROM lx_token where token ='" + t + "' limit 1")

	var token LxToken;
	if rows.Next() {
		err := rows.Scan(&token.id, &token.Token, &token.Status, &token.UserId)
		util.ShowError(err)
	}
	rows.Close()

	return token
}

func InsertToken(token LxToken) int64 {
	UpdateToken(token)

	str := "INSERT INTO `lx_token` (`id`, `token`, `status`, `userId`) VALUES (NULL, '%s', '%d', '%d')"
	return db.Insert(fmt.Sprintf(str, token.Token, token.Status, token.UserId))
}

func UpdateToken(token LxToken) int64 {
	str := "UPDATE `lx_token` SET  `status` =  '0' WHERE  `userId` = %d"
	return db.Update(fmt.Sprintf(str, token.UserId))
}

func GetArticle(id int) LxArticle {
	rows := db.Query(fmt.Sprintf("select * FROM lx_article where id = %d", id))
	var article LxArticle;
	if rows.Next() {
		article = parserArticle(rows)
	}

	rows.Close()

	return article
}

func GetArticles(sql string) []LxArticle {
	rows := db.Query(sql)
	var articles []LxArticle;
	for rows.Next() {
		article := parserArticle(rows)
		articles = append(articles, article)
	}
	rows.Close()
	return articles
}

func InsertArticle(article LxArticle) int64 {
	str := "INSERT INTO `lx_article` (`id`, `userId`, `title`, `content`, `createTime`) VALUES (NULL, '%d', '%s','%s' ,UNIX_TIMESTAMP())"
	return db.Insert(fmt.Sprintf(str, article.UserId, article.Title, article.Content))
}

func UpdateArticle(id int) int64 {
	str := "UPDATE `lx_article` SET  `readCount` = `readCount` + 1 WHERE  `id` = %d"
	return db.Update(fmt.Sprintf(str, id))
}

func DeleteArticle(id int64) int64 {
	str := "DELETE FROM `lx_article` WHERE `id` = %d"
	return db.Update(fmt.Sprintf(str, id))
}

func parserArticle(rows *sql.Rows) LxArticle {
	var article LxArticle;
	err := rows.Scan(&article.Id, &article.UserId, &article.Title, &article.Content, &article.CreateTime, &article.ReadCount)
	util.ShowError(err)

	article.CreateTimeF = util.TimeFormat(article.CreateTime)
	article.User = GetUserById(article.UserId)
	article.User.Userpwd = ""
	return article;
}
