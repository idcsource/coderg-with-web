package askanswer

import(
	"fmt"
	"time"
	"coderg"
)


type AddControllor struct {
	coderg.Controllor
}

func (c *AddControllor) Exec(){
	myconf := c.Runtime.MyConfig;
	node_unique, _ := myconf.GetInt64("default","node_unique");
	
	c.R.ParseForm();
	var question, email, content string;
	_, ok1 := c.R.PostForm["question"];
	if ok1 == true {
		question = c.R.PostForm["question"][0];
	}
	_, ok2 := c.R.PostForm["email"];
	if ok2 == true {
		email = c.R.PostForm["email"][0];
	}
	_, ok3 := c.R.PostForm["content"];
	if ok3 == true {
		content = c.R.PostForm["content"][0];
	}
	if ok1 != true || ok2 != true || ok3 != true {
		fmt.Fprint(c.W, "参数错误");
		return;
	}
	
	ip := coderg.NewInputProcessor();
	question ,err1 := ip.Text(question, true, 4, 55);
	email, err2 := ip.Email(email,false,4,200);
	content, err3 := ip.Text(content, true, 4, 2000);
	if err1 != 0 || err2 != 0 || err3 != 0 {
		fmt.Fprint(c.W, "输入不符合要求");
		return;
	}
	
	nowtime := time.Now().Unix();
	
	p, _ := c.Service.Database.Prepare("insert into ask_not_answer (node, title, email, content, datetime) values ($1, $2, $3, $4, $5)");
	_, err := p.Exec(node_unique, question, email, content, nowtime);
	if err != nil {
		fmt.Fprint(c.W, "数据添加有误");
		return;
	}
	fmt.Fprint(c.W, "ok");
	return;
}
