package controllor

import (
	"fmt"
	//"text/template"
	"os"
	"io/ioutil"
	
	//"coderg"
)


type IndexTemplData struct {
	Main MainOutData
	NowPageName string
	PageCss string
	MainGoFile string
}

type IndexControllor struct {
	MainOutControllor
}

func (my *IndexControllor) Exec(){
	my.MainOutControllor.Exec();
	
	filename := my.Service.GlobalRelativePath + "static/file/index-main-file.htm";
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
	
	data := IndexTemplData{Main: my.Main, NowPageName : "首页", MainGoFile: filecontent};
	/*
	s1, err := template.ParseFiles(my.Service.GlobalRelativePath + "template/index.tmpl",my.Service.GlobalRelativePath + "template/mainout.tmpl");
	if err != nil {
		fmt.Println(err);
		return;
	}
	s1.ExecuteTemplate(my.W, "mainout", data);
	*/
	my.Service.Template["index"].ExecuteTemplate(my.W, "mainout", data);
	//num, _ := coderg.Service.TemplateConfig["index"].GetString("loop1","num");
	//fmt.Println(my.Service.GlobalRelativePath);
	//coderg.NodeConfigGet(my.Runtime, my.Service.NodeTree);
	//fmt.Fprint(my.W, "这里是IndexControllor");
	//fmt.Fprint(my.W,my.Service.TemplateConfig);
}
