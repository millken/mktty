package dt

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/millken/mktty/common"
)

func init() {
	DtRegister("dns.getList", NewDnsGetList)
}

type DnsGetList struct {
	param Param
}

type Record struct {
	Id     int
	Domain string
	Value  []byte
}

func NewDnsGetList(param Param) (Action, error) {
	return &DnsGetList{
		param: param,
	}, nil
}

func (d *DnsGetList) Response() (data gin.H, err error) {
	maxID := common.StrToInt(d.param.Get.Get("maxid"))
	limit := common.StrToInt(d.param.Get.Get("limit"))
	offset := common.StrToInt(d.param.Get.Get("offset"))

	sqlstr := fmt.Sprintf("SELECT * FROM config.record where id<=%d order by id asc limit %d offset %d", maxID, limit, offset)
	records := []Record{}
	d.param.Db.Select(&records, sqlstr)
	data = gin.H{"records": records}
	return
}
