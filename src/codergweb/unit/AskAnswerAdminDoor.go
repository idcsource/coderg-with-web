package unit

import (
	//"fmt"
	
	"coderg"
	
	"codergweb/unit/askanswer"
)



type AskAnswerAdminDoor struct {
	
}
func (s *AskAnswerAdminDoor) Router()(route *coderg.UnitRouter){
	route = coderg.NewUnitRouter();
	route.Add("admin-ad",&askanswer.AdminAnswerDelControllor{});
	route.Add("admin",&askanswer.AdminControllor{});
	//route.Add("add",&askanswer.AddControllor{});
	//route.Add("/",&askanswer.MainControllor{});
	return; 
}
