package employ

import (
	"robot-server/config"
	"robot-server/internal/core/regist"
	"robot-server/internal/micro/employ/service"

	"robot-server/api/apiarea"
	"robot-server/api/apicompany"
	"robot-server/api/apiemployee"
	"robot-server/api/apiexport"

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
