package dt

import "github.com/gin-gonic/gin"

func init() {
	DtRegister("cdn.addServer", NewCdnAddServer)
}

type CdnAddServer struct {
	param Param
}

func NewCdnAddServer(param Param) (Action, error) {
	return &CdnAddServer{
		param: param,
	}, nil
}

func (d *CdnAddServer) Response() (data gin.H, err error) {
	servername := d.param.Get.Get("servername")
	setting := d.param.Get.Get("setting")

	sqlstr := "insert into cdn.server (servername, setting) VALUES (:servername, :setting) ON CONFLICT (servername) DO UPDATE SET setting = :setting"
	_, err = d.param.Db.NamedExec(sqlstr,
		map[string]interface{}{
			"servername": servername,
			"setting":    setting,
		})
	return
}
