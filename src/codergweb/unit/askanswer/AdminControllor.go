package askanswer

import(
	"text/template"
	"fmt"
	"strings"
	"strconv"
	"coderg"
	"codergweb/controllor"
)

type AskAnswerAdminData struct {
	Main controllor.AdminLoginData
	NowPageName string
	RealNode string
	IfNext bool
	NextPage uint64
	IfPrev bool
	PrevPage uint64
	AskAnswer []AskAnswerOne
}


type AdminControllor struct {
	controllor.CheckLoginControllor
}

func (a *AdminControllor) Exec(){
	iflogin := a.CheckLogin();
	if iflogin == false {
		return;
	}
	a.CheckLoginControllor.Exec();
	
	a.Main.PageCss = append(a.Main.PageCss, "askanswer.css");
	
	a.Main.PageJs = append(a.Main.PageJs, "jquery-2.1.3.min.js");
	a.Main.PageJs = append(a.Main.PageJs, "php.js");
	a.Main.PageJs = append(a.Main.PageJs, "js_function.js");
	a.Main.PageJs = append(a.Main.PageJs, "RequestProcess.js");
	
	myconf := a.Runtime.MyConfig;
	nodename ,_ := myconf.GetString("default","nodename");
	onepagei, _ := myconf.GetInt64("default","onepage");
	onepage := uint64(onepagei);
	node_unique, _ := myconf.GetInt64("default","node_unique");
	
	var nowpage uint64
	nowpage = 1;
	if len(a.Runtime.UrlRequest) != 0 {
		nowpagei, e := strconv.Atoi(a.Runtime.UrlRequest[0]);
		if e == nil {
			nowpage = uint64(nowpagei);
		}
	}
	
	data := AskAnswerAdminData{Main: a.Main, NowPageName : nodename, RealNode:  a.Runtime.RealNode};
	data.AskAnswer = make([]AskAnswerOne,0,onepage);
	
	rp := coderg.NewInputProcessor();
	
	var allAskNum uint64
	a.Service.Database.QueryRow("select COUNT(*) FROM ask_not_answer WHERE node = $1", node_unique).Scan(&allAskNum);
	if allAskNum != 0 {
		aa_rows, _ :=a.Service.Database.Query("select id, title, email, content, datetime, answertime, answer From ask_not_answer WHERE node = $1 ORDER BY id DESC LIMIT $2 OFFSET $3", node_unique, onepage, (nowpage - 1) * onepage);
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
	
	s1, err := template.ParseFiles(a.Service.GlobalRelativePath + "template/askansweradmin.tmpl",a.Service.GlobalRelativePath + "template/adminout.tmpl");
	if err != nil {
		fmt.Println(err);
		return;
	}
	s1.ExecuteTemplate(a.W, "adminout", data);
}
