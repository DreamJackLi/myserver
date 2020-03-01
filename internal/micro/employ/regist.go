package employ

import (
	"myserver/config"
	"myserver/internal/core/regist"
	"myserver/internal/micro/employ/service"

	"myserver/api/apiarea"
	"myserver/api/apicompany"
	"myserver/api/apiemployee"
	"myserver/api/apiexport"

	"google.golang.org/grpc"
)

func init() {
	// 注册 service
	config.LoadService(regist.ServerName[regist.ALL])
	regist.InitServer(regist.ServerName[regist.ALL], func(s *grpc.Server) {
		apiemployee.RegisterEmployeeServer(s, &service.EmployServer{})
		apicompany.RegisterCompanyServer(s, &service.CompanyServer{})
		apiarea.RegisterAreaServer(s, &service.AreaServer{})
		apiexport.RegisterExportServer(s, &service.ExportServer{})
	})
}
