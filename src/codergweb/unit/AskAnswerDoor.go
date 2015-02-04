package unit

import (
	//"fmt"
	
	"coderg"
	
	"codergweb/unit/askanswer"
)



type AskAnswerDoor struct {
	
}
func (s *AskAnswerDoor) Router()(route *coderg.UnitRouter){
	route = coderg.NewUnitRouter();
	//route.Add("admin-ad",&askanswer.AdminAnswerDelControllor{});
	//route.Add("admin",&askanswer.AdminControllor{});
	route.Add("add",&askanswer.AddControllor{});
	route.Add("/",&askanswer.MainControllor{});
	return; 
}
