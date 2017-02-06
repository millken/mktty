package dt

import "github.com/gin-gonic/gin"

func init() {
	DtRegister("/dt.conf", NewDtConf)
}

type DtConf struct {
	ap ActionParam
}

func NewDtConf(param Param) (Action, error) {
	return &DtConf{
		ap: param.Ap,
	}, nil
}

//{"status":200,"data":[{"requestId":"1","status":200,"action":"/dt.conf","data":{"payPolicy":[{"types":"sdk"}]}}],"cookies":[]}
func (d *DtConf) Response() (data gin.H, err error) {
	data = gin.H{
		"requestId": d.ap.RequestId,
		"status":    200,
		"action":    d.ap.Action,
		"data": gin.H{
			"payPolicy": []gin.H{
				gin.H{
					"types": "sdk",
				},
			},
		},
	}
	return
}
