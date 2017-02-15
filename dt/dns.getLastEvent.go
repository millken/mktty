package dt

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func init() {
	DtRegister("dns.getLastEvent", NewDnsGetLastEvent)
}

type DnsGetLastEvent struct {
	param Param
}

type Event struct {
	Domain string
	Act    string
	Utime  string
	Value  interface{}
}

func NewDnsGetLastEvent(param Param) (Action, error) {
	return &DnsGetLastEvent{
		param: param,
	}, nil
}

func (d *DnsGetLastEvent) Response() (data gin.H, err error) {
	utime := d.param.Get.Get("utime")

	sqlstr := fmt.Sprintf("select event.domain, utime, act, value from config.event left outer join config.record on event.domain = record.domain where utime>'%s' order by utime asc", utime)
	records := []Event{}
	err = d.param.Db.Select(&records, sqlstr)
	if err != nil {
		log.Printf("[ERROR] query last event: %s", err)
		return
	}
	data = gin.H{"records": records}
	return
}
