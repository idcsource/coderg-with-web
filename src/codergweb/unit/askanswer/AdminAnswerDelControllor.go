package askanswer

import(
	"fmt"
	"strconv"
	"time"
	"coderg"
	"codergweb/controllor"
)

type AdminAnswerDelControllor struct {
	controllor.CheckLoginControllor
}

func (a *AdminAnswerDelControllor) Exec(){
	iflogin := a.CheckLogin();
	if iflogin == false {
		return;
	}
	a.CheckLoginControllor.Exec();
	
	a.R.ParseForm();
	
	var types int;
	var id uint64;
	var answer string;
	
	_, ok1 := a.R.PostForm["types"];
	if ok1 == true {
		var err error;
		types, err = strconv.Atoi(a.R.PostForm["types"][0]);
		if err != nil {
			fmt.Fprint(a.W, "参数错误");
			return;
		}
	}
	
	ip := coderg.NewInputProcessor();
	
	if types == 1{
		//回答问题的
		_, ok2 := a.R.PostForm["id"];
		if ok2 != true {
			fmt.Fprint(a.W, "参数错误");
			return;
		}
		ids, err := strconv.ParseInt(a.R.PostForm["id"][0],10,64);
		if err != nil{
			fmt.Fprint(a.W, "参数错误");
			return;
		}
		id = uint64(ids);
		answer = a.R.PostForm["answer"][0];
		answer, _ := ip.Text(answer, false, 4, 2000);
		
		nowtime := time.Now().Unix();
		
		p, _ := a.Service.Database.Prepare("update ask_not_answer set answertime = $1, answer = $2 where id = $3");
		_, err2 := p.Exec(nowtime, answer, id);
		if err2 != nil {
			fmt.Fprint(a.W, "数据添加有误");
			return;
		}else{
			fmt.Fprint(a.W, "ok");
			return;
		}
	}else if types == 2 {
		//删除问题
		_, ok2 := a.R.PostForm["id"];
		if ok2 != true {
			fmt.Fprint(a.W, "参数错误");
			return;
		}
		ids, err := strconv.ParseInt(a.R.PostForm["id"][0],10,64);
		if err != nil{
			fmt.Fprint(a.W, "参数错误");
			return;
		}
		id = uint64(ids);
		
		p, _ := a.Service.Database.Prepare("delete from ask_not_answer where id = $1");
		_, err2 := p.Exec(id);
		if err2 != nil {
			fmt.Fprint(a.W, "数据添加有误");
			return;
		}else{
			fmt.Fprint(a.W, "ok");
			return;
		}
	}else{
		fmt.Fprint(a.W, "参数错误");
		return;
	}
}
