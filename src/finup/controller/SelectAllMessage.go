package controller

import (
	db "finup/database"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

type Message struct {
	AppCustomerId    int
	AppRequestId     int
	LendRequestId    int
	AppCustomerName  string
	AppMobile        string
	AppIdNo          string
	AppLogin         int
	SalesNo          string
	AppStateType     string
	LendStatus       string
	LendCustomerId   string
	LendCustomerName string
	LendCustomerIdNo string
	LendMinStatus    string
}

func (m *Message) getAll(str string) (messages []Message, err error) {

	sqlStr := `
	SELECT
	lc.id AS app_customer_id ,
	lr.id AS app_request_id ,
	fl.id AS lend_id ,
	lc.customer_name as app_customer_name,
	lc.mobile AS app_moblie ,
	lc.id_no AS app_id_no ,
	lc.log_in_id,
	lc.sales_no,	
	lr.state_type ,	
	fl.status AS finup_lend_status ,
	fl.lend_customer_id,
	fc.name as lend_customer_name,
	fc.id_no AS finup_lend_id_no,
	lrs.sub_status as min_status

	FROM
	lend_app.app_customer lc
	LEFT JOIN lend_app.app_lend_request lr ON lc.id = lr.app_customer_id
	LEFT JOIN finup_lend.lend_request fl ON lr.id = fl.app_lend_request_id
	LEFT JOIN finup_lend.lend_customer fc ON fl.lend_customer_id = fc.id
	left join finup_lend.lend_request_substatus lrs on fl.id=lrs.lend_request_id
	WHERE lr.id=%s`

	sqlStr = fmt.Sprintf(sqlStr, str)
	rows, err := db.My.Query(sqlStr)
	if err != nil {
		log.Fatalln("query is error", err)
	}

	for rows.Next() {
		var message Message
		rows.Scan(&message.AppCustomerId, &message.AppRequestId, &message.LendRequestId, &message.AppCustomerName, &message.AppMobile,
			&message.AppIdNo, &message.AppLogin, &message.SalesNo, &message.AppStateType, &message.LendStatus, &message.LendCustomerId,
			&message.LendCustomerName, &message.LendCustomerIdNo, message.LendMinStatus)
		messages = append(messages, message)

	}

	return

}

func SelectAllMessage(c *gin.Context) {
	var m Message
	var messages, err = m.getAll("10002566")
	if err != nil {
		log.Fatalln("selectAllMessage is error")
	}
	c.JSON(http.StatusOK, gin.H{
		"code":     http.StatusOK,
		"messages": messages,
	})

}

func Posttest(c *gin.Context) {

	num := c.PostForm("num")
	data, _ := ioutil.ReadAll(c.Request.Body)
	fmt.Printf("ctx.Request.body: %v", string(data))

	fmt.Println(num)
	nick := c.DefaultPostForm("nick", "anonymous") // 此方法可以设置默认值
	c.JSON(http.StatusOK, gin.H{
		"num":  num + "哈哈哈",
		"nick": nick,
		"code": http.StatusOK,
	})

}

func HeadersAuth() gin.HandlerFunc {

	return func(context *gin.Context) {
       //权限验证 通过headers
		Auth := context.Request.Header.Get("Auth")

		if Auth=="YLS" {
			context.Next()
			return

		}
		context.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		context.Abort()

	}

}
