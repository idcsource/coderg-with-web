package controllor

import (
	"fmt"
	//"text/template"
	"os"
	"io/ioutil"
	
	//"coderg"
)


type WhyIpTemplData struct {
	Main MainOutData
	NowPageName string
	PageCss string
	WhyIpFile string
}

type WhyIpControllor struct {
	MainOutControllor
}

func (my *WhyIpControllor) Exec(){
	my.MainOutControllor.Exec();
	
	filename := my.Service.GlobalRelativePath + "static/file/why-ip.htm";
	file, err := os.Open(filename);
	if err != nil{
		fmt.Fprint(my.W, err);
		return;
	}
	filed, err := ioutil.ReadAll(file);
	if err != nil{
		fmt.Fprint(my.W, err);
		return;
	}
	filecontent := string(filed);
	
	my.Main.PageCss = append(my.Main.PageCss, "index.css");
	
	data := WhyIpTemplData{Main: my.Main, NowPageName : "为什么是一个Ip加port", WhyIpFile: filecontent};
	/*
	s1, err := template.ParseFiles(my.Service.GlobalRelativePath + "template/whyip.tmpl",my.Service.GlobalRelativePath + "template/mainout.tmpl");
	if err != nil {
		fmt.Println(err);
		return;
	}
	*/
	my.Service.Template["whyip"].ExecuteTemplate(my.W, "mainout", data);
}
