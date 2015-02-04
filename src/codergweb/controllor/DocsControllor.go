package controllor

import(
	"strings"
	//"text/template"
	"os"
	"io/ioutil"
	"fmt"
	
)

type DocsMenuStruct struct {
	Name string
	Mark string
}

type DocsTemplData struct {
	Main MainOutData
	NowPageName string
	IfNext bool
	NextName string
	NextMark string
	IfPrev bool
	PrevName string
	PrevMark string
	NowDoc string
	IfMain bool
	Menu []DocsMenuStruct
}

type DocsControllor struct {
	MainOutControllor
}

func (d *DocsControllor) Exec(){
	d.MainOutControllor.Exec();
	
	d.Main.PageCss = append(d.Main.PageCss, "markdown.css");
	d.Main.PageCss = append(d.Main.PageCss, "index.css");
	
	d.Main.PageJs = append(d.Main.PageJs, "jquery-2.1.3.min.js");
	d.Main.PageJs = append(d.Main.PageJs, "Markdown.Converter.js");
	
	/*
	s1, err := template.ParseFiles(d.Service.GlobalRelativePath + "template/docs.tmpl",d.Service.GlobalRelativePath + "template/mainout.tmpl");
	if err != nil {
		fmt.Println(err);
		return;
	}
	*/
	myconf := d.Runtime.MyConfig;
	data := DocsTemplData{Main: d.Main};
	data.IfNext = true;
	data.IfPrev = true;
	data.NowPageName , _ = myconf.GetString("default","nodename");  //节点名称
	filepath, _ := myconf.GetString("default","files");  //Doc文件存放位置
	filepath = d.Service.GlobalRelativePath + filepath;
	allmarks, _ := myconf.GetString("default","menu");  //所有文件的mark
	allmark := strings.Split(allmarks, ",");
	var nowmark string; //当前的mark
	if len(d.Runtime.UrlRequest) != 0 {
		nowmark = d.Runtime.UrlRequest[0];
	}else{
		nowmark, _ = myconf.GetString("default","main"); 
	}
	nowmarknum := 0;
	allmarknum := len(allmark);
	for i, one := range allmark {
		if one == nowmark {
			nowmarknum = i;
		}
	}
	filename := filepath + nowmark + ".md";
	if nowmarknum == 0 { 
		data.IfPrev = false;
		data.IfMain = true;
		data.Menu = make([]DocsMenuStruct,0,allmarknum);
		for i, one := range allmark {
			if i == 0 { continue; }
			name, _ := myconf.GetString(one, "name");
			oneMenu := DocsMenuStruct{Name : name, Mark : one};
			data.Menu = append(data.Menu,oneMenu);
		}
	}else{
		data.PrevMark = allmark[nowmarknum - 1];
		data.PrevName, _ = myconf.GetString(data.PrevMark, "name");
		data.IfMain = false;
	}
	if nowmarknum == allmarknum - 1 {
		data.IfNext = false;
	}else{
		data.NextMark = allmark[nowmarknum + 1];
		data.NextName, _ = myconf.GetString(data.NextMark, "name");
	}
	file, err := os.Open(filename);
	if err != nil{
		fmt.Fprint(d.W, err);
		return;
	}
	filed, err := ioutil.ReadAll(file);
	if err != nil{
		fmt.Fprint(d.W, err);
		return;
	}
	data.NowDoc = string(filed);
	d.Service.Template["docs"].ExecuteTemplate(d.W, "mainout", data);
}
