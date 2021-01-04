package utils

import (
	"Alertgateway/common/share"
	"Alertgateway/dao"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/dyvmsapi"
	"github.com/gin-gonic/gin"
)

func NoResponse(c *gin.Context) {
	// c.JSON(http.StatusNotFound, gin.H{
	//     "status": 404,
	//     "error":  "404 ,page not exists!",
	// })
	c.HTML(http.StatusNotFound, "error_return_index.html", gin.H{"msg": "请求地址错误!"})
}

func Md5sum(strs string) string {
	h := md5.New()
	h.Write([]byte(strs)) // 需要加密的字符串
	cipherStr := h.Sum(nil)
	md5res := fmt.Sprintf("%s", hex.EncodeToString(cipherStr)) // 输出加密结果
	return md5res
}

//获取对象的方法
func TypeRef(params interface{}) {
	value := reflect.ValueOf(params)
	typ := value.Type()
	for i := 0; i < value.NumMethod(); i++ {
		fmt.Println(fmt.Sprintf("Method: [%d]%s ...and..   Type is: %v", i, typ.Method(i).Name, typ.Method(i).Type))
	}
}

//阿里云打电话
func AliyunCall(accessKeyId, accessSecret, TtsCode, CalledShowNumber string, cn *[]dao.CalledTelNum, tp dao.TtsParam) error {
	client, err := dyvmsapi.NewClientWithAccessKey("cn-hangzhou", accessKeyId, accessSecret)
	if err != nil {
		fmt.Println("aliyun api NewClientWithAccessKey is err:", err)
	}
	request := dyvmsapi.CreateSingleCallByTtsRequest()
	request.Scheme = "https"
	request.TtsCode = TtsCode
	request.CalledShowNumber = CalledShowNumber

	// jsontp, err := json.Marshal(tp)
	// if err != nil {
	// 	fmt.Println("参数转换失败!", err)
	// 	return errors.New("params turn falsed")
	// }

	// fmt.Println("jsontp---", "string(jsontp)", string(jsontp), fmt.Sprintf("{\"content\":\"%s\"}", tp.AlertInfo))
	request.TtsParam = fmt.Sprintf("{\"content\":\"%s\"}", tp.AlertInfo) //string(jsontp)

	if len(*cn) != 0 {

		for _, item := range *cn {
			fmt.Println("*********阿里云的电话:******", item.Cnumber, "===", item.IsCall)
			if item.IsCall {
				request.CalledNumber = item.Cnumber
				_, err := client.SingleCallByTts(request)
				if err != nil {

					fmt.Println("阿里云调用错误!", err)
					return errors.New("called err")
				}
			}

		}
	} else {
		return errors.New("called num is null")
	}
	return nil

}

//判断参数传递过来的是哪个模板,是否是在维护时间
func TlsLimitTime(tp *dao.TtsParam) bool {
	vt := *dao.GetVoicetem()
	for _, item := range vt {

		if item.Name == tp.AlertPf {
			//看下维护 状态
			if item.IsCanUsed {

				fmt.Println(item.LimitTime)
				splitres := strings.Split(item.LimitTime, " - ")

				startTime, _ := time.Parse("2006-01-02 15:04:05", splitres[0])
				endTime, _ := time.Parse("2006-01-02 15:04:05", splitres[1])

				share.Logger.Warn("开始时间:", startTime, startTime.Unix())
				share.Logger.Warn("结束时间:", endTime, endTime.Unix())

				NowTime := time.Now().Unix()

				if NowTime < startTime.Unix() || NowTime > endTime.Unix() {

					share.Logger.Warn("[不在维护间....]")
					return true
				}
				fmt.Println("在维护间")
				share.Logger.Warn("[在维护间....]")
				return false
			} else {

				share.Logger.Warn("[维护状态停止....]")
				return true
			}
		}

	}

	return false
}

func ReturnOK(c *gin.Context, msg interface{}) {
	if value, ok := msg.(string); ok {
		c.JSON(http.StatusOK, gin.H{
			"code":      "200",
			"message":   value,
			"timestamp": time.Now().Format("2006-01-02 15:04:05"),
		})
	} else {
		c.JSON(http.StatusOK, msg)
	}

}

func ReturnErr(c *gin.Context, msg interface{}) {
	if value, ok := msg.(string); ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":      "500",
			"message":   value,
			"timestamp": time.Now().Format("2006-01-02 15:04:05"),
		})
	} else {
		c.JSON(http.StatusBadRequest, msg)
	}

}
