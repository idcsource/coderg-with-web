package controllor

import (
	//"fmt"
	//"html/template"
	
	"coderg"
)

type MainOutData struct{
	WebSiteName string
	PageTitleKey string
	PageCss []string
	PageJs  []string
}

type MainOutControllor struct {
	coderg.Controllor
	Main MainOutData
}

func (my *MainOutControllor) Exec(){
	webname, _ := my.Service.WebConfig.GetString("web","name");
	titlekey, _ := my.Service.WebConfig.GetString("web","titleKey");
	my.Main = MainOutData{} ;
	my.Main.WebSiteName = webname;
	my.Main.PageTitleKey = titlekey;
	my.Main.PageCss = []string{"header.css"};
	my.Main.PageJs = make([]string,0);
	//num, _ := coderg.Service.TemplateConfig["index"].GetString("loop1","num");
	//fmt.Println(my.Service.GlobalRelativePath);
	//coderg.NodeConfigGet(my.Runtime, my.Service.NodeTree);
	//fmt.Fprint(my.W, "这里是IndexControllor");
	//fmt.Fprint(my.W,my.Service.TemplateConfig);
}
