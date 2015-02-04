package admin

import(
	"coderg"
	"net/http"
	"fmt"
	//"codergweb/controllor"
)


type LoginDoControllor struct {
	coderg.Controllor
}

func (a *LoginDoControllor) Exec(){
	a.R.ParseForm();
	username := a.R.PostForm["username"][0];
	password := a.R.PostForm["password"][0];
	check := a.R.PostForm["check"][0];
	
	ip := coderg.NewInputProcessor();
	
	username,erri := ip.Text(username, true, 2, 250);
	if erri != 0 {
		fmt.Fprint(a.W, "用户名密码错误");
		return;
	}
	var dbuserpassword string;
	err := a.Service.Database.QueryRow("select pass from admins where name = '" + username + "'").Scan(&dbuserpassword);
	if err != nil {
		fmt.Fprint(a.W,"用户名密码错误");
		return;
	}
	
	dbuserpassword = check + dbuserpassword;
	dbuserpassword = coderg.GetSha1(dbuserpassword);
	
	if password == dbuserpassword {
		theCookieName, _ := a.Service.WebConfig.GetString("admin","cookie");
		cookie := http.Cookie{Name: theCookieName, Value: username, Path: "/", MaxAge: 0};
		http.SetCookie(a.W, &cookie);
		fmt.Fprint(a.W, "ok");
	}else{
		fmt.Fprint(a.W,"用户名密码错误");
	}
	return;
}
