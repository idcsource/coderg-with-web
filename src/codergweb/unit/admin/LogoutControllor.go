package admin

import(
	"coderg"
	"fmt"
	"net/http"
	//"codergweb/controllor"
)


type LogoutControllor struct {
	coderg.Controllor
}

func (a *LogoutControllor) Exec(){
	theCookieName, _ := a.Service.WebConfig.GetString("admin","cookie");
	theNotLogin, _ := a.Service.WebConfig.GetString("admin","nologin");
	
	cookie := http.Cookie{Name: theCookieName, Value: "", Path: "/", MaxAge: -1};
	http.SetCookie(a.W, &cookie);
	fmt.Fprint(a.W,"<script language='javascript'>window.location.href='" + theNotLogin + "'</script>");
}
