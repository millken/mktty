package dt

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/millken/mktty/common"
)

func init() {
	DtRegister("cdn.getList", NewCdnGetList)
}

type CdnGetList struct {
	param Param
}

type Servers struct {
	Id                  int
	Servername, Setting string
}

func NewCdnGetList(param Param) (Action, error) {
	return &CdnGetList{
		param: param,
	}, nil
}

func (d *CdnGetList) Response() (data gin.H, err error) {
	maxID := common.StrToInt(d.param.Get.Get("maxid"))
	limit := common.StrToInt(d.param.Get.Get("limit"))
	offset := common.StrToInt(d.param.Get.Get("offset"))

	sqlstr := fmt.Sprintf("SELECT * FROM cdn.server where id<=%d order by id asc limit %d offset %d", maxID, limit, offset)
	records := []Servers{}
	d.param.Db.Select(&records, sqlstr)
	data = gin.H{"records": records}
	return
}
