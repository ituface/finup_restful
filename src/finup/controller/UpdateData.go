package controller

import (
	db "finup/database"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

const (
	_update_annex_status = `UPDATE app_lend_annex_status SET status  = 'OBTAIN_SUCCESS' WHERE app_lend_request_id = ? AND type in (
                            SELECT material_type FROM app_apply_material WHERE access_mode = 'CAPTURE' AND is_enable = 1)`
	_required_picture = `select apm.material_type from lend_app.app_customer ac LEFT join lend_app.app_lend_request alr on alr.app_customer_id=ac.id
						left join lend_app.app_product_define apd on FIND_IN_SET(ac.shop_id,test_shop_code)
						left join lend_app.app_supplement_data asd on  apd.id=asd.product_id 
						left join lend_app.app_apply_material apm on asd.material_type=apm.material_type
						where alr.id=? and apd.product_type='QUICK2.0' and apm.access_mode='PHOTO' and (asd.required=1 or asd.group_type is not null)`
)


func Sql_manage(id string) int64 {

	rs, err := db.My.Exec(_update_annex_status, id)
	if err != nil {
		fmt.Println("异常产生")
	}
	rows, err := rs.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	return rows

}

func SqlSelectRequired(id string) []string {
  rows,err:=db.My.Query(_required_picture,id)
  if err!=nil{
  	log.Fatal("异常产生")
  }
  var Required_list [] string

  for rows.Next()  {
	  var one_r string
	  rows.Scan(&one_r)
	  Required_list= append(Required_list, one_r)
	}

  return Required_list

}
func Upate_manage(c *gin.Context) {
	app_request_id := c.PostForm("id")
	fmt.Println("app id---------",app_request_id)
	rows := Sql_manage(app_request_id)

	Rs:=SqlSelectRequired(app_request_id)
	fmt.Println("rs-------",Rs)
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"num":  rows,
		"required":Rs,


	})

}
