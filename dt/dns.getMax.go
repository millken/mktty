package dt

import (
	"log"

	"github.com/gin-gonic/gin"
)

func init() {
	DtRegister("dns.getMax", NewDnsGetMax)
}

type DnsGetMax struct {
	param Param
}

func NewDnsGetMax(param Param) (Action, error) {
	return &DnsGetMax{
		param: param,
	}, nil
}

func (d *DnsGetMax) Response() (data gin.H, err error) {
	utime, err := d.getMaxTime()
	if err != nil {
		log.Printf("[ERROR] getMaxTime() : %s", err)
		return
	}
	maxID, countID, err := d.getMaxCount()
	if err != nil {
		log.Printf("[ERROR] getMaxTime() : %s", err)
		return
	}
	log.Printf("[DEBUG] postgres utime=%s, maxID=%d, countID=%d", utime, maxID, countID)
	data = gin.H{
		"utime": utime,
		"maxid": maxID,
		"total": countID,
	}
	return
}

func (d *DnsGetMax) getMaxCount() (maxID, countID int, err error) {
	err = d.param.Db.QueryRow("select max(id) max_id,count(id) count_id from dns.record").Scan(&maxID, &countID)
	return
}

func (d *DnsGetMax) getMaxTime() (utime string, err error) {
	err = d.param.Db.QueryRow("select max(utime) utime from dns.event").Scan(&utime)
	return
}
