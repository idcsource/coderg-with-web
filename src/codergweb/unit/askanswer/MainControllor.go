package askanswer

import(
	//"text/template"
	//"fmt"
	"strconv"
	"strings"

	"coderg"
	"codergweb/controllor"
)

type AskAnswerOne struct {
	Id uint64
	Title string
	Email string
	Date uint64
	Content string
	IfAnswer bool
	AnswerDate uint64
	Answer string
}

type AskAnswerMainData struct {
	Main controllor.MainOutData
	NowPageName string
	RealNode string
	IfNext bool
	NextPage uint64
	IfPrev bool
	PrevPage uint64
	AskAnswer []AskAnswerOne
}


type MainControllor struct {
	controllor.MainOutControllor
}

func (c *MainControllor) Exec(){
	c.MainOutControllor.Exec();
	
	c.Main.PageCss = append(c.Main.PageCss, "askanswer.css");
	
	c.Main.PageJs = append(c.Main.PageJs, "jquery-2.1.3.min.js");
	c.Main.PageJs = append(c.Main.PageJs, "php.js");
	c.Main.PageJs = append(c.Main.PageJs, "js_function.js");
	c.Main.PageJs = append(c.Main.PageJs, "RequestProcess.js");
	
	myconf := c.Runtime.MyConfig;
	nodename ,_ := myconf.GetString("default","nodename");
	onepagei, _ := myconf.GetInt64("default","onepage");
	onepage := uint64(onepagei);
	node_unique, _ := myconf.GetInt64("default","node_unique");
	
	
	var nowpage uint64
	nowpage = 1;
	if len(c.Runtime.UrlRequest) != 0 {
		nowpagei, e := strconv.Atoi(c.Runtime.UrlRequest[0]);
		if e == nil {
			nowpage = uint64(nowpagei);
		}
	}
	
	data := AskAnswerMainData{Main: c.Main, NowPageName : nodename, RealNode:  c.Runtime.RealNode};
	data.AskAnswer = make([]AskAnswerOne,0,onepage);
	
	rp := coderg.NewInputProcessor();
	
	var allAskNum uint64
	c.Service.Database.QueryRow("select COUNT(*) FROM ask_not_answer WHERE node = $1", node_unique).Scan(&allAskNum);
	if allAskNum != 0 {
		aa_rows, _ :=c.Service.Database.Query("select id, title, email, content, datetime, answertime, answer From ask_not_answer WHERE node = $1 ORDER BY id DESC LIMIT $2 OFFSET $3", node_unique, onepage, (nowpage - 1) * onepage);
		for aa_rows.Next(){
			one := AskAnswerOne{};
			aa_rows.Scan(&one.Id, &one.Title, &one.Email, &one.Content, &one.Date, &one.AnswerDate, &one.Answer);
			one.Title = strings.TrimSpace(one.Title);
			one.Email = strings.TrimSpace(one.Email);
			one.Content = rp.TextareaOut(one.Content, true);
			if len(one.Answer) != 0 {
				one.IfAnswer = true;
			}else{
				one.IfAnswer = false;
			}
			one.Answer = rp.TextareaOut(one.Answer, true);
			data.AskAnswer = append(data.AskAnswer, one);
		}
	}
	if nowpage == 1{
		data.IfPrev = false;
	}else{
		data.IfPrev = true;
		data.PrevPage = nowpage - 1;
	}
	if nowpage * onepage >= allAskNum {
		data.IfNext = false;
	}else{
		data.IfNext = true;
		data.NextPage = nowpage + 1;
	}
	
	/*
	s1, err := template.ParseFiles(c.Service.GlobalRelativePath + "template/askanswer.tmpl",c.Service.GlobalRelativePath + "template/mainout.tmpl");
	if err != nil {
		fmt.Println(err);
		return;
	}
	*/
	c.Service.Template["askanswer"].ExecuteTemplate(c.W, "mainout", data);
}
