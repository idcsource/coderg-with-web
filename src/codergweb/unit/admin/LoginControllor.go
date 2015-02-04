package admin

import(
	"text/template"
	"fmt"
	"codergweb/controllor"
)

type LoginControllorData struct{
	Main controllor.MainOutData
	NowPageName string
	RealNode string
}


type LoginControllor struct {
	controllor.MainOutControllor
	//controllor.CheckLoginControllor
}

func (a *LoginControllor) Exec(){
	a.MainOutControllor.Exec();
	
	a.Main.PageCss = append(a.Main.PageCss, "index.css");
	
	a.Main.PageJs = append(a.Main.PageJs, "jquery-2.1.3.min.js");
	a.Main.PageJs = append(a.Main.PageJs, "php.js");
	
	s1, err := template.ParseFiles(a.Service.GlobalRelativePath + "template/login.tmpl",a.Service.GlobalRelativePath + "template/mainout.tmpl");
	if err != nil {
		fmt.Println(err);
		return;
	}
	
	data := LoginControllorData{ Main: a.Main, NowPageName : "Login", RealNode: a.Runtime.RealNode};
	s1.ExecuteTemplate(a.W, "mainout", data);
}
