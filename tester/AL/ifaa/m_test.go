package main

import(
	"fmt"
	"testing"
	

	"com.es.gateway.sdk/sdk"
	"com.es.gateway.sdk/sdk/app/ifaa"

)
/**
 * IFAA认证，服务开通地址：https://market.aliyun.com/products/57000002/cmapi00038615.html?spm=5176.shop.result.97.553928d8UHqi88&innerSource=search
 * 参考文档：https://esandinfo.yuque.com/books/share/324b13b7-7e24-40e0-a445-99d52c57e9db?#
 * 管理控制台地址：https://openali.esandcloud.com/
 */
 func Test1(t *testing.T){
	// TODO 替换成你自己的appcode, 获取APPCODE  可参考 https://esandinfo.yuque.com/yv6e1k/ulp2ub/fs2mm48opwox3xc4?singleDoc#
	appCode :=  "TODO"; 
	gateway := sdk.ESAlyunGateway{}
	gateway.Init(appCode)
	// 构造请求报文
	sAuthRequest := ifaa.SAuthRequest{}
    sAuthRequest.IfaaMsg = "oTRva8ai5LeONn5rUIfQU7ZvWgj6AdsDfhB3ijksLEeqLGJvECJkEOftNyimG8R7mgMfFtUFBc0gvGmgkrYt2bSuW5ixGDtelRBI7U1gqxGtm5xDFM+TdzUCZaclKP0lfSEhs2l0NRN3BLeAaUJh+lmGU5OSIpAIMACK6r5nwabvNp/v6XsAs+0/47XR5QMMJBVpFXf1kzG51xqjbVYfPoHbCjZnImwvC76qPjDiK8k0q2qBCx7QXTH04MLr3bdBsCJ0cy8mvu+AZC2lAbOebl7FHILPRa8+MBe4Arut7BXKvqbvYYVb7K7VK41EPLN/LQm1qjYO7VTiRuesmj16bRctjo77qEg1QSM4un30adm2Hbq4rp3vclBb8RAkcFykzqKa7JXSAy3U5vW1eDnpVlfb3q0XXTvTiZBR0Ah2hV4gPEWa9eoGQP/a22XWyhr+HY3bIkx2COGYnPge09QXxlKI5NPZq/qYEh68GTlYssfeAyvNx4oR4el7bcx59PQlT9o6JEKqIUqubHzuxs/bN/O+IuDkNrRZl4bBDM50nbC8nqVUc/tF8RWAeGRZixntcS4udq6lLXzagR2Hc7pcJK/Bzm2arvMikZb2v/uz9oX63XuySnHioWfzgqnxQYcOFJKbToD0rV4S5J2YmWryrjjk0G5bNvXehKD1mXvlUUtvthWzXj0OU/5/P93WAqwHx/F95AEFi5+iZSQ//HzAHdPbfN3EUnlNiCB9Ni/fs9XiljqgRn0Iou5OBWrR7wRkZ49DubderLYneujZZOyEaWZioSN9ZqNVMkGTY6X52hJUrCajV85XpX/4XxgApHGt3r7OjbBTf1m8TUqva5TY+9ckYyrz24p9gYtoiW0qjbRCb6N0FLRys9VILRwqaLjaHqD8VoFvCh4FAgJSJUhJdy4ePwmzn6jXffj51nrT16kx3hIjkLkPPukjGFkKDL4ier7Re+pOinPfiH87xBkid6nBQhG9WUZoH9j1OLbLMvBWbSxVSXq1rTYu+3Jt3bl7Y3qOsqURA2JiaPrnL4h4bwPcIq/Woy6kQvkFEqMPvAZF5mkDqd3c8OHvPAOee08lNQFBBEEsZUB5l7Fj/OdYwUeIy1thD5cprKRFRUoALi2m+hQeo6SRs0GF3YyAOwZTq9BMEjDnodnryEHDKf213VnSUhcTVZJhpJc6s/S57gQufaUnL4ktGygB9RR+8pIPNxALTdq2ZkNBV2brEZ1fYnGsZ+33EFgzFEdkec16/VuYBYah+K2fRqoOpITEuCrMsoqygSD9w3Fy1g2GAamL6/fT+Bpn8Ox8BWqms12DnndaOm2I6QEw4+sEoSwGsiuDRNdH07YQ5dXXOtDaQLnL64MZlk80xlcVm3PfTdl6VXjiUWhSUQlFtWweZKbk9ZOoJ8lpH/uNNnFRRUumrIVTU0EAiSo5oekCDZF9k+UwdGAbQDyRwnKRkfnOdInpli+bF+XmMUzQxUEJNrObPIqZLVi/OFHHKQG1GxYi766RqWfuERX7fGKyxnow2kAEJV0NM4RNJo5fVHYL+jRUHqJKc6zlhetyn//7VLYzqHFf17WG9BFFqWD2iR5T+RSsgwgT"

	bizContent := sAuthRequest.ToJsonStr()
	rspMsg := gateway.SendToGateWay(ifaa.SAuthRequestALIYUN_URL, bizContent)
	if rspMsg ==""{
		t.Error("服务器返回内容为空")
		return
	}
	//解析返回的报文内容
	fmt.Println("服务器返回内容为: ",rspMsg)
	sAuthResponse := ifaa.SAuthResponse{}
	fmt.Println("服务器返回业务数据 : ", sAuthResponse.ToJsonStr())

}