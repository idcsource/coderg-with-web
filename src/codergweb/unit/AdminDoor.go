package unit

import (
	//"fmt"
	
	"coderg"
	
	"codergweb/unit/admin"
)



type AdminDoor struct {
	
}
func (s *AdminDoor) Router()(route *coderg.UnitRouter){
	route = coderg.NewUnitRouter();
	route.Add("logout",&admin.LogoutControllor{});
	route.Add("logindo",&admin.LoginDoControllor{});
	route.Add("login",&admin.LoginControllor{});
	route.Add("/",&admin.AdminControllor{});
	return; 
}
