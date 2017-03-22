package dt

import "github.com/gin-gonic/gin"

func init() {
	DtRegister("cdn.delServer", NewCdnDelServer)
}

type CdnDelServer struct {
	param Param
}

func NewCdnDelServer(param Param) (Action, error) {
	return &CdnDelServer{
		param: param,
	}, nil
}

func (d *CdnDelServer) Response() (data gin.H, err error) {
	servername := d.param.Get.Get("servername")

	sqlstr := "delete from cdn.server where servername=:servername"
	_, err = d.param.Db.NamedExec(sqlstr,
		map[string]interface{}{
			"servername": servername,
		})
	return
}
