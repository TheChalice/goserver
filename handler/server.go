package handler

import (
	"code.cloudfoundry.org/lager"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"

	"io/ioutil"
	"net/http"
	"os"
)

const JSON = "application/json"

var log lager.Logger
var ssoHost, dfHost string

func init() {
	//dfHost = getENV("DFHOST")
	log = lager.NewLogger("Token_Proxy")
	log.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG))
}

type Postbody struct {
	N string `json:"n"`
	M string `json:"m"`
	H string `json:"H"`
	L string `json:"L"`
	ATTR []Point `json:"attr"`
	LOAD []Load `json:"load"`
}
type Point struct {
	SOURCE Pointtype `json:"source"`
	TARGET Pointtype `json:"target"`
	MATERIAL string `json:"material"`
	SECTION string `json:"section"`
	SIZE Sizetype `json:"size"`
}
type Pointtype struct {
	POINT string `json:"point"`
	COORDINATE Coordinate `json:"coordinate"`
}
type Coordinate struct {
	X int `json:"x"`
	Y int `json:"y"`
}
type Sizetype struct {
	T3 string `json:"t3"`
	Tw string `json:"tw"`
	T2 string `json:"t2,omitempty"`
	TF string `json:"tf,omitempty"`
}
type Load struct {
	SOURCE Pointtype `json:"source"`
	TARGET Pointtype `json:"target"`
	GERINO string `json:"region"`
	INTRODUCE string `json:"introduce"`
	DEADLOAD string `json:"deadload"`
	LIVELOAD string `json:"liveload"`
}


func Cluster(c *gin.Context) {

	//token := getToken(c)
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("CreateBC Read Request.Body error", err)
	}
	var postbody Postbody

	_ = json.Unmarshal(rBody, &postbody)
	fmt.Printf("%#v\n",postbody)

	//users :=c.Request.Header["Authorization"][0]

	//errorRsp := ErrorResponse{}

	//result := ""

	//dfRsp, err := trRequest("GET", "https://"+dfHost+"/oauth/authorize?client_id=openshift-challenging-client&response_type=token", user)

	//if err != nil {
	//	log.Error("GetToken request fail error", err)
	//	errorRsp.Description = "GetToken request fail error"
	//	errorRsp.Error = err.Error()
	//	c.JSON(500, errorRsp)
	//	return
	//}
	//
	//if value, ok := dfRsp.Header["Location"]; ok {
	//	result = getDFtoken(value[0])
	//} else {
	//	log.Debug("GetToken respose header fail error")
	//	errorRsp.Description = "GetToken respose header fail error"
	//	errorRsp.Error = err.Error()
	//	c.JSON(500, errorRsp)
	//	return
	//}

	//if len(result) == 0 {
	//	log.Debug("GetToken respose header[location]  fail error")
	//	errorRsp.Description = "GetToken respose header[location]  fail error"
	//	errorRsp.Error = err.Error()
	//	c.JSON(500, errorRsp)
	//	return
	//}

	//rsp := RstResponse{}
	//rsp.User = "admin"
	////rsp.Token = result
	//fmt.Printf("%#v\n",rsp)
	c.JSON(http.StatusOK, postbody)
	return

}
