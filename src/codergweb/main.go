package main

import(
	//"fmt"
	"text/template"
	
	"coderg"
	
	"codergweb/controllor"
	"codergweb/unit"
)



func main(){
	
	service :=coderg.ServicePrepare("codergweb.cfg");
	
	service.MainRouter.AddStatic("/img","static/img");
	service.MainRouter.AddStatic("/css","static/css");
	service.MainRouter.AddStatic("/js","static/js");
	
	/*
	 * 节点树结构
	 * /
	 *  - WhyIp
	 *  - AskAnswer
	 *      - add
	 *  - Docs
	 *  - Admin
	 *      - AskAnswer
	 */
	service.NodeTree = coderg.NewNodeTree(&controllor.IndexControllor{});
	service.NodeTree.AddNode("WhyIp","",&controllor.WhyIpControllor{});
	service.NodeTree.AddNode("Docs","nodetree/docs.cfg",&controllor.DocsControllor{});
	service.NodeTree.AddUnit("AskAnswer","nodetree/askanswer.cfg",&unit.AskAnswerDoor{});
	
	admin, _ := service.NodeTree.AddUnit("Admin","nodetree/admin.cfg",&unit.AdminDoor{});
	admin.AddUnit("AskAnswer","nodetree/askanswer.cfg",&unit.AskAnswerAdminDoor{});
	
	service.NodeTree.Router(service);
	
	
	//缓存模板
	local := service.GlobalRelativePath;
	service.Template["index"], _ = template.ParseFiles(local + "template/index.tmpl",local + "template/mainout.tmpl");
	service.Template["docs"], _ = template.ParseFiles(local + "template/docs.tmpl",local + "template/mainout.tmpl");
	service.Template["whyip"], _ = template.ParseFiles(local + "template/whyip.tmpl",local + "template/mainout.tmpl");
	service.Template["askanswer"], _ = template.ParseFiles(local + "template/askanswer.tmpl",local + "template/mainout.tmpl");
	
	service.Start();
	
	service.Wait();
}
