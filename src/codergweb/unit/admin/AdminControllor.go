package admin

import(
	"text/template"
	"fmt"
	"codergweb/controllor"
)

type AdminData struct{
	Main controllor.AdminLoginData
	NowPageName string
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
	
	a.Main.PageCss = append(a.Main.PageCss, "index.css");
	
	data := AdminData{Main: a.Main, NowPageName : "管理首页"};
	
	s1, err := template.ParseFiles(a.Service.GlobalRelativePath + "template/admin.tmpl",a.Service.GlobalRelativePath + "template/adminout.tmpl");
	if err != nil {
		fmt.Println(err);
		return;
	}
	s1.ExecuteTemplate(a.W, "adminout", data);
}
