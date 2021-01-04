package dao

import (
	"Alertgateway/common/share"
	"fmt"
	"strconv"
	"strings"

	"github.com/jinzhu/gorm"
)

//数据都从数据库里获取

//被呼叫电话
type CalledTelNum struct {
	gorm.Model
	Cnumber    string `gorm:"type:varchar(100);not null" json:"cnumber"` //被呼电话
	BelongName string `gorm:"type:varchar(100)" json:"belongname"`       //所属人
	BelongDep  string `gorm:"type:varchar(100)" json:"belongdep"`        //所属部门
	Oper       string `gorm:"type:varchar(100)" json:"oper"`             //添加人操作
	IsCall     bool   `json:"iscall"`
}

//呼叫号码
type CalledShowTelNum struct {
	gorm.Model
	Csnumber   string `gorm:"type:varchar(100);not null"` //呼叫电话
	Oper       string `gorm:"type:varchar(100)"`          //添加人操作
	Statslimit string `gorm:"type:varchar(100)"`          //状态
	LimitTime  string `gorm:"type:varchar(100)"`          //限呼时段
}

//报警基本信息
type AlertMsg struct {
	gorm.Model
	AccessKeyId  string `gorm:"type:varchar(100);not null"`
	AccessSecret string `gorm:"type:varchar(100);not null"`
	TtsCode      string `gorm:"unique;not null"`
	PTimes       int
}

//语音模板
type Voicetem struct {
	gorm.Model
	Name          string `gorm:"type:varchar(100);not null" json:"name"`     //模板名称
	Variable      string `gorm:"type:varchar(300);not null" json:"variable"` //模板变量
	LimitTime     string `gorm:"type:varchar(100)" json:"limitTime"`         //模板维护时间段
	LimitTimeobj  string `gorm:"type:varchar(100)" json:"limitTimeobj"`      //模板维护时间前端对象
	TypePro       string `gorm:"type:varchar(100)" json:"typePro"`           //模板所属项目(类别)
	IsCanUsed     bool   `json:"iscanused"`                                  //是否维护
	Count         uint32 `json:"count"`                                      //报警次数
	CalledTelNums string `gorm:"type:varchar(300)"`                          //添加外键费劲,直接进行查询操作
}

// TelAlertmsg...

// ${name} 超过阈值 ${yuzhi}
type TtsParam struct {
	AlertPf string `json:"name"` //报警名称 == 模板名称
	// AlertType string `json:"alerttype"` //报警类型
	AlertInfo string `json:"content"` //报警信息
	// Recodtime string `json:"recode"`    //报警时间
	//对应的Voicetem语音模板的主键id
	// AlertMaTian string `json:"AlertMaTian"` //维护时间time.Time
}

//多选下拉框
type Selectxmdata struct {
	Name     uint   `json:"name"`
	Value    string `json:"value"`
	Selected bool   `json:"selected"`
}

func NewAlertObj(alertname string) (*AlertMsg, *CalledShowTelNum, *[]CalledTelNum) { //cn *[]dao.CalledTelNum
	am := AlertMsg{}
	cs := CalledShowTelNum{}
	vt := Voicetem{}
	cn := []CalledTelNum{}
	// var ctnslicestr  []string
	// var CalledTelNums []string
	GDB.First(&am)
	GDB.First(&cs)
	GDB.Find(&vt, "Name=?", alertname)

	if vt.CalledTelNums != "" {
		for _, item := range strings.Split(vt.CalledTelNums, ",") {
			currcn := CalledTelNum{}
			id, _ := strconv.Atoi(item)
			GDB.Find(&currcn, "id=?", uint(id))
			cn = append(cn, currcn)
		}
		return &am, &cs, &cn
	} else {
		return nil, nil, nil
	}

}

func GetAlertTempDB() []Voicetem {
	vt := []Voicetem{}
	GDB.Find(&vt)

	share.Logger.Warn("获取所有报警模板!")
	return vt
}

func GetAlertPeopleDB() []CalledTelNum {
	ct := []CalledTelNum{}
	GDB.Find(&ct)
	share.Logger.Warn("获取所有报警人!")
	return ct
}

func AddalertpeopleDB(ctn CalledTelNum) {
	// ctn := dao.CalledTelNum{}
	// ctn := CalledTelNum{
	// 	Cnumber:    Cnumber,
	// 	BelongName: BelongName,
	// 	BelongDep:  BelongDep,
	// 	Oper:       Oper,
	// }
	// 创建记录
	GDB.Create(&ctn)
	share.Logger.Warn("创建报警人成功!")

}

func ChangeCnstat(iscall, id string) {

	i, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Atoi is err:", err)
		share.Logger.Warn("Atoi is err:", err)
	}
	// ctn := &CalledTelNum{ID: uint(i)}
	ctn := &CalledTelNum{}
	if iscall == "1" {

		GDB.Model(ctn).Where("id=?", uint(i)).Update("IsCall", true)

	} else {

		GDB.Model(ctn).Where("id=?", uint(i)).Update("IsCall", false)
	}

}

func Delctnwithid(id string) error {

	//物理删除
	GDB.Debug().Unscoped().Where("id = ?", id).Delete(&CalledTelNum{})
	// GDB.Debug().Where("id = ?", id).Delete(&PermissionDisgroup{})
	return nil
}

func Updatetel(user, tel string) error {
	ctn := CalledTelNum{}
	GDB.Debug().Where("belong_name = ?", user).First(&ctn)

	ctn.Cnumber = tel

	GDB.Save(&ctn)
	return nil
}

//报警模板
func GetVoicetem() *[]Voicetem {
	ctn := []Voicetem{}
	GDB.Find(&ctn)
	return &ctn
}

//更改数据库中报警模板中的参数
func SaveVtParms(p string, id string) {
	i, err := strconv.Atoi(id)
	if err != nil {

		share.Logger.Warn("Atoi is err:", err)

	}

	vt := &Voicetem{}

	GDB.Model(vt).Where("id=?", uint(i)).Update("Variable", p)

}

func ChangeTemTimebutton(id, iscanused string) {
	i, err := strconv.Atoi(id)
	if err != nil {

		share.Logger.Warn("Atoi is err:", err)
	}
	vt := &Voicetem{}
	if iscanused == "1" {
		GDB.Model(vt).Where("id=?", uint(i)).Update("IsCanUsed", true)
	} else {
		GDB.Model(vt).Where("id=?", uint(i)).Update("IsCanUsed", false)
	}
}

func AlertTemplaydatechange(value, vdate string) {
	vt := Voicetem{}

	GDB.Model(vt).Where("limit_timeobj=?", vdate).Update("LimitTime", value)
}

func GetCurrAlertPeople(id string) []Selectxmdata {
	var alerpeosobj []Selectxmdata
	for _, item := range GetAlertPeopleDB() {
		sx := Selectxmdata{}
		sx.Name = item.ID
		sx.Value = item.BelongName
		alerpeosobj = append(alerpeosobj, sx)
	}

	i, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Atoi is err:", err)
	}

	vt := Voicetem{}
	GDB.First(&vt, i)

	if vt.CalledTelNums != "" {

		for index, sx := range alerpeosobj {
			for _, slicei := range strings.Split(vt.CalledTelNums, ",") {
				if strconv.FormatUint(uint64(sx.Name), 10) == slicei {

					sx.Selected = true

					alerpeosobj[index] = sx
				}
			}

		}
	}
	return alerpeosobj
}

func SaveTemaddAlerPeo(res, id string) error {
	i, err := strconv.Atoi(id)
	if err != nil {

		share.Logger.Warn("Atoi is err:", err)
		return err
	}
	vt := &Voicetem{}

	sliceres := strings.TrimRight(res, ",")

	GDB.Model(vt).Where("id=?", uint(i)).Update("CalledTelNums", sliceres)
	return nil
}

func TempAddCon(t_name, t_content string) {

	vtinter := Voicetem{}
	GDB.Last(&vtinter)
	crrliobj := vtinter.LimitTimeobj
	vtltobj := strings.Split(crrliobj, "datetimestr")[1]
	cdtsint, err := strconv.Atoi(vtltobj)
	if err != nil {
		fmt.Println("转换失败!")
	}
	vt := Voicetem{
		Name:         t_name,
		Variable:     t_content,
		LimitTimeobj: "#datetimestr" + fmt.Sprintf("%d", cdtsint+1),
	}
	GDB.Create(&vt)
}

func DelTempMsg(id string) {
	//物理删除
	GDB.Debug().Unscoped().Where("id = ?", id).Delete(&Voicetem{})
}
