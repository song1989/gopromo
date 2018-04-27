package routers

import (
	"gopromo/app/http/controllers/Common/PsController"
	"gopromo/app/http/controllers/Common/TestController"

	"net/http"
)

func common() routeTypes {
	commonRouter := commonRouter{}
	return routeTypes{
		prefix: "common",
		handlers: []routeType{
			commonRouter.imageHandle(),
			commonRouter.testFunc(),
		},
	}
}

type commonRouter struct {
}

func (this *commonRouter) imageHandle() routeType {
	return routeType{
		prefix: "image",
		handler: map[string]http.HandlerFunc{
			"handle": midd.Chain(
				PsController.Handler,
				midd.Method("GET"),
				midd.PromoId(),
				//midd.Logging(),
				midd.Validate(map[string]string{
					//"name":      "required",
					"image_url": "required",
				}),
			),
		},
	}
}

func (this *commonRouter) testFunc() routeType {
	return routeType{
		prefix: "test",
		handler: map[string]http.HandlerFunc{
			"banyan": TestController.Banyan,
			"usertable": midd.Chain(
				TestController.Usertable,
				midd.Method("POST"),
				midd.Logging(),
				//midd.IsLogin(),
				midd.Validate(map[string]string{
					//"id": "between:1,10",
					//"mobile": "int",
				})),
			"testone":    TestController.TestOne,
			"testthrift": TestController.Testthrift,
		},
	}
}
