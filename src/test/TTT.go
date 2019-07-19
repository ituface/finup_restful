package main

import "fmt"

func main() {

	strs := `SELECT
	lc.id AS app_customer_id ,
		lc.customer_name as app_customer_name,
		lc.mobile AS app_moblie ,
		lc.id_no AS app_id_no ,
		lc.log_in_id,
		lc.sales_no,
		lr.id AS app_request_id ,
		lr.state ,
		lr.audit_state ,
		lr.state_type ,
		fl.id AS lend_id ,
		fl.'status' AS finup_lend_status ,
		fl.lend_customer_id,
		fc.id_no AS finup_lend_id_no,
		fc.name as lend_customer_name,
		lrs.sub_status as min_status
 	    FROM
  	lend_app.app_customer lc
	LEFT JOIN lend_app.app_lend_request lr ON lc.id = lr.app_customer_id
	LEFT JOIN finup_lend.lend_request fl ON lr.id = fl.app_lend_request_id
	LEFT JOIN finup_lend.lend_customer fc ON fl.lend_customer_id = fc.id
	left join finup_lend.lend_request_substatus lrs on fl.id=lrs.lend_request_id
	WHERE
	#fl.status='PROCESSING'
#fl.id=786339
	#lc.customer_name='测北海三';
	#lc.id=5562
	#lc.mobile='xy43bee39e1bab6d6277fc7f25d0e2086520160926'
lc.id_no='%s'`

	strs="fad%sfafa"


	s:=fmt.Sprintf(strs,"12")


	fmt.Println(s)



}
