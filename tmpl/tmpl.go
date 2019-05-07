package tmpl

import (
	"book/model"
	"book/util"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func init() {

}

func Index(r *http.Request) string {
	//return "Hello, World!"
	h := NewHtml();
	//h.body("<h1>Hello, World</h1>")
	//h.body(util.GetViewTpl("index.html"))

	list := model.GetArticles("select * from lx_article order by id desc")
	var str string
	tml := `
		<div class="row">
        <div class="col-left">
			<img src="/view/res/img/file_101.353.png" class="img128"/>
				<a href="/view?id=%d" target="_blank">%s</a>
		</div>
        <div class="col-right">%s</div>
        <div class="col-right1"><a href="#">%s</a></div>
    </div>`
	for _, s := range list {
		str += fmt.Sprintf(tml, s.Id, s.Title, s.CreateTimeF, s.User.Username)
	}
	h.body(strings.Replace(util.GetViewTpl("index.html"), "{{content}}", str, -1))
	return h.Show("首页")
}

func SignIn(r *http.Request) string {
	h := NewHtml();
	h.body(util.GetViewTpl("sign_in.html"))
	return h.Show("登陆")
}

func SignUp(r *http.Request) string {
	h := NewHtml();
	h.body(util.GetViewTpl("sign_up.html"))
	return h.Show("注册")
}

func SignInSave(r *http.Request) string {
	res := "{\"code\":%d,\"info\":\"%s\",\"token\":\"%s\"}"

	r.ParseForm()
	raw := model.LxUser{
		Username: r.PostForm.Get("username"),
		Userpwd:  r.PostForm.Get("userpwd"),
	}

	user := model.GetUser(raw.Username)
	if user.Userpwd == util.Md5(raw.Userpwd) {
		user.LoginIp = "127::0"
		user.LoginTime = int(util.TimeI())
		token := model.LxToken{
			Token:  util.Md5(raw.Userpwd + strconv.FormatInt(int64(user.LoginTime), 10)),
			Status: 1,
			UserId: user.Id,
		}
		model.UpdateUser(user)
		model.InsertToken(token)
		return fmt.Sprintf(res, 1, "登陆成功", token.Token)
	}

	return fmt.Sprintf(res, 0, "登陆失败", "")
}

func SignUpSave(r *http.Request) string {
	res := "{\"code\":%d,\"info\":\"%s\"}"

	r.ParseForm()
	raw := model.LxUser{
		Username: r.PostForm.Get("username"),
		Userpwd:  r.PostForm.Get("userpwd"),
	}
	if len(raw.Userpwd) > 0 {
		user := model.GetUser(raw.Username)
		if user.CreateTime == 0 {
			raw.Userpwd = util.Md5(raw.Userpwd)
			model.InsertUser(raw)
			return fmt.Sprintf(res, 1, "注册成功")
		}
	}

	return fmt.Sprintf(res, 0, "注册失败")
}

func AddArticel(r *http.Request) string {
	h := NewHtml();
	h.body(util.GetViewTpl("add.html"))
	return h.Show("发表")
}

func AddSaveArticel(r *http.Request) string {
	res := "{\"code\":%d,\"info\":\"%s\",\"token\":\"%s\"}"

	t, _ := r.Cookie("token")
	token := getToken(t.Value)
	fmt.Println(token)
	if token.Status == 1 {
		r.ParseForm()
		raw := model.LxArticle{
			UserId:  token.UserId,
			Title:   r.PostForm.Get("title"),
			Content: strings.ReplaceAll(r.PostForm.Get("content"),"'","&#39;"),
		}

		model.InsertArticle(raw)
		return fmt.Sprintf(res, 1, "发表成功", token.Token)

	}

	return fmt.Sprintf(res, 0, "发表失败", "")
}

func GetArticle(r *http.Request) string {
	h := NewHtml();
	r.ParseForm()
	id, _ := strconv.Atoi(r.Form.Get("id"))

	a := model.GetArticle(id)
	model.UpdateArticle(id)
	var str string
	tml := `<div class="title"><h2>%s</h2></div>
			<div class="row">
				<div class="fr">作者:%s</div>
				<div class="fr">发表时间:%s</div>
				<div class="fr">阅读:%d</div>
			</div>
			<div class="content">%s</div>`

	str = fmt.Sprintf(tml, a.Title, a.User.Username, a.CreateTimeF, a.ReadCount, a.Content)

	h.body(strings.Replace(util.GetViewTpl("article.html"), "{{content}}", str, -1))
	return h.Show(a.Title + "|文章详情")
}

func getToken(token string) model.LxToken {
	return model.GetToken(token)
}
