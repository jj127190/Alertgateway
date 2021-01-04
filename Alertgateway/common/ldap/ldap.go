package ldap

import (
	"errors"
	"fmt"

	"github.com/go-ldap/ldap/v3"
	// "Alertgateway/common/utils"
	// "gopkg.in/ldap.v3"
	// "log"
)

func LdapVerifica(username, password string) error {
	var L *ldap.Conn
	L, err := ldap.DialURL("ldap://1.2.3.4:389")
	if err != nil {
		return err
	}
	defer L.Close()
	_, err = L.SimpleBind(&ldap.SimpleBindRequest{
		Username: "CN=admin,DC=blingabc,DC=com",
		Password: "123456",
	})
	if err != nil {
		return err
	}
	searchRequest := ldap.NewSearchRequest(
		"dc=blingabc,dc=com", // The base dn to search
		ldap.ScopeWholeSubtree,

		ldap.DerefAlways,
		0,
		0,
		false,
		"(cn=ops)",
		[]string{"memberUid"},

		nil,
	)

	cur, err := L.Search(searchRequest)
	if err != nil {
		return err
	}

	memberUids := cur.Entries[0].GetAttributeValues("memberUid")
	if len(memberUids) == 0 {
		return errors.New("账号不存在!")
	}
	var tagnum int16
	for _, item := range memberUids {
		if item == username {
			tagnum = 1
		}
	}
	if tagnum == 1 {
		return nil
		// 有问题,所以重新启动一个函数查密码
		// 开始查询密码
		// fmt.Println("开始查询密码!")
		// searchRequestsig := ldap.NewSearchRequest(
		// 	"dc=blingabc,dc=com",
		// 	ldap.ScopeWholeSubtree,
		// 	ldap.DerefAlways,
		// 	0,
		// 	0,
		// 	false,
		// 	// filter,
		// 	"(uid=huapang)",
		// 	// fmt.Sprintf("(uid=%v)", username),
		// 	[]string{"uid", "userPassowrd"},
		// 	nil,
		// )

		// curs, err := L.Search(searchRequestsig)
		// if err != nil {
		// 	return err
		// }
		// fmt.Println("111")
		// fmt.Println(curs.Entries[0])
		// fmt.Println(curs.Entries[0].GetAttributeValue("userPassword"))
		// ldapres := curs.Entries[0]
		// userPassword := ldapres.GetAttributeValue("userPassword")
		// fmt.Println("密码是", userPassword)
		// 开始查询密码

	} else {
		return errors.New("账号不存在!")
	}
	// return errors.New("未知错误!")

}

func GetLdaPass(username string) (string, error) {

	l, err := ldap.DialURL("ldap://1.2.3.4:389")
	if err != nil {
		return "", err
	}
	_, err = l.SimpleBind(&ldap.SimpleBindRequest{
		Username: "CN=admin,DC=blingabc,DC=com",
		Password: "123456",
	})
	if err != nil {
		return "", err
	}
	searchRequest := ldap.NewSearchRequest(
		"dc=blingabc,dc=com", // The base dn to search
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		fmt.Sprintf("(uid=%s)", username),
		[]string{"userPassword"},
		nil,
	)
	cur, err := l.Search(searchRequest)
	if err != nil {
		return "", err
	}
	ldapres := cur.Entries[0]
	userPassword := ldapres.GetAttributeValue("userPassword")
	defer l.Close()
	return userPassword, nil
}

//获取所有的账号
func GetLdapUsers() (error, []string) {
	var L *ldap.Conn
	L, err := ldap.DialURL("ldap://1.2.3.4:389")
	if err != nil {
		return err, nil
	}
	defer L.Close()
	_, err = L.SimpleBind(&ldap.SimpleBindRequest{
		Username: "CN=admin,DC=blingabc,DC=com",
		Password: "1234567",
	})
	if err != nil {
		return err, nil
	}
	searchRequest := ldap.NewSearchRequest(
		"dc=blingabc,dc=com", // The base dn to search
		ldap.ScopeWholeSubtree,

		ldap.DerefAlways,
		0,
		0,
		false,
		"(cn=ops)",
		[]string{"memberUid"},

		nil,
	)

	cur, err := L.Search(searchRequest)
	if err != nil {
		return err, nil
	}

	memberUids := cur.Entries[0].GetAttributeValues("memberUid")
	if len(memberUids) == 0 {
		return errors.New("账号不存在!"), nil
	}
	var users []string

	for _, item := range memberUids {
		fmt.Println(item)
		users = append(users, item)
	}
	return nil, users

	// return errors.New("未知错误!"),nil

}
