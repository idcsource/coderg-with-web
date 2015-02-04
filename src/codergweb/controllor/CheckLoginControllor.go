package controllor

import(
	//"database/sql"
	"fmt"
	
	"coderg"
)

type AdminLoginData struct {
	WebSiteName string
	AdminName string
	PageCss []string
	PageJs []string
}

type CheckLoginControllor struct {
	coderg.Controllor
	Main AdminLoginData
}

func (c *CheckLoginControllor) CheckLogin()(iflogin bool){
	theCookieName, _ := c.Service.WebConfig.GetString("admin","cookie");
	theNotLogin, _ := c.Service.WebConfig.GetString("admin","nologin");
	
	user_cookie, err := c.R.Cookie(theCookieName);
	if err != nil {
		fmt.Fprint(c.W,"<script language='javascript'>window.location.href='" + theNotLogin + "'</script>");
		iflogin = false;
		return;
	}
	user := user_cookie.Value;
	
	var sqlgetname string;
	checkuser_err := c.Service.Database.QueryRow("select name from admins where name = '"+user+"'").Scan(&sqlgetname);
	
	//if checkuser_err == sql.ErrNoRows {
	if checkuser_err != nil {
		fmt.Fprint(c.W,"<script language='javascript'>window.location.href='" + theNotLogin + "'</script>");
		iflogin = false;
		return;
	}
	iflogin = true;
	return;
}

func (c *CheckLoginControllor) Exec(){
	
	theCookieName, _ := c.Service.WebConfig.GetString("admin","cookie");
	
	user_cookie, _ := c.R.Cookie(theCookieName);
	user := user_cookie.Value;
	
	c.Main = AdminLoginData{};
	c.Main.AdminName = user;
	c.Main.WebSiteName , _ = c.Service.WebConfig.GetString("web","name");
	
	c.Main.PageCss = []string{"header.css"};
	c.Main.PageJs = make([]string,0);
}
