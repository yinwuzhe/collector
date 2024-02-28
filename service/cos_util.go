package service

import (
	"net/http"
	"net/url"
	"os"

	"github.com/tencentyun/cos-go-sdk-v5"
)

var cosAddr = "https://weixin-1251754822.cos.ap-guangzhou.myqcloud.com"

var instance *cos.Client

func GetCosClient() *cos.Client {

	if instance == nil {
		u, _ := url.Parse(cosAddr)
		b := &cos.BaseURL{BucketURL: u}
		instance := cos.NewClient(b, &http.Client{

			Transport: &cos.AuthorizationTransport{
				SecretID:  os.Getenv("COS_SecretID"),
				SecretKey: os.Getenv("COS_SecretKey"),
			},
		})
		return instance
	}

	return instance
}
