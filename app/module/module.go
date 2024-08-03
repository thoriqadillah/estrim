package module

import (
	"estrim/app"
	"estrim/app/module/account"
	"estrim/app/module/compressor"
	"estrim/app/module/storage"
	"estrim/app/module/view"
)

func init() {
	app.RegisterService(account.NewService)
	app.RegisterService(compressor.NewService)
	app.RegisterService(storage.NewService)
	app.RegisterService(view.NewService)
}
