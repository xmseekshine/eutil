package tencent

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/minio/minio-go/v7"
	"github.com/tencentyun/cos-go-sdk-v5"
)

//先实现简单文件上下传读写 了解后再增加更过参数
//新建存储桶
//上传文件对象
//下载文件对象
//删除文件对象
//批量上传文件对象
//批量下载文件对象
//批量删除文件对象
//根据配置获取对象存储Client
//不存在创建
// 对不同的存储桶需要配置不同的域名 ex:https://ft-site-1256195644.cos.ap-shanghai.myqcloud.com
//TODO 腾讯云存储支持大文件分块操作
// 地域：地域（Region）是腾讯云托管机房的分布地区，对象存储 COS 的数据存放在这些地域的存储桶中。您可以通过 COS，
//      将数据进行多地域存储。通常情况下，COS 建议您选择在与您业务最近的地域上创建存储桶，以满足低延迟、低成本以及合规性要求。
// 域名：默认域名指 COS 的默认存储桶域名，用户在创建存储桶时，由系统根据存储桶名称和地域自动生成。不同地域的存储桶有不同的默认域名。
type (
	TenCosClient interface {
		HeadBucket(bucketName string, loc string) error
		PutBucket(bucketName string, loc string) error
		GetBucket(bucketName string, loc string) error
		GetObjectList()
		UploadFile(bucketName, objectName, filePath, contentType string) (*cos.Response, error)
		DownloadFile(bucketName, objectName, filePath string) error
		DeleteFile(bucketName, objectName string) error
		DeleteFiles(bucketName string, objectInfos []minio.ObjectInfo) error
	}

	defaultTenCosClient struct {
		Client *cos.Client
	}

	TenCosConfig struct {
		BucketURL  string //访问 bucket, object 相关 API 的基础 URL（不包含 path 部分）
		ServiceURL string //访问 service API 的基础 URL（不包含 path 部分）
		BatchURL   string //访问 Batch API 的基础 URL （不包含 path 部分）
		CIURL      string //访问 CI 的基础 URL （不包含 path 部分）
		SecretID   string //Access key是唯一标识你的账户的用户ID。
		SecretKey  string //Secret key是你账户的密码。
	}
)

func NewTenCosClient(c *TenCosConfig) TenCosClient {
	u, _ := url.Parse(c.BucketURL)
	su, _ := url.Parse(c.ServiceURL)
	b := &cos.BaseURL{BucketURL: u, ServiceURL: su}
	// 1.永久密钥
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  c.SecretID,
			SecretKey: c.SecretKey,
		},
	})
	return &defaultTenCosClient{
		Client: client,
	}
}

//client 绑定了存储桶
func (m *defaultTenCosClient) HeadBucket(bucketName string, loc string) error {
	_, err := m.Client.Bucket.Head(context.Background())
	if err != nil {
		panic(err)
	}
	return nil
}

/*PUT Bucket 接口请求可以在指定账号下创建一个存储桶。该 API 接口不支持匿名请求，
您需要使用带 Authorization 签名认证的请求才能创建新的 Bucket 。创建存储桶的用户默认成为存储桶的持有者。
创建存储桶时，如果没有指定访问权限，则默认使用私有读写（private）权限。
如需创建多 AZ 存储桶，那么应当通过请求体指示存储桶配置，否则无需传入请求体。
*/
func (m *defaultTenCosClient) PutBucket(bucketName string, loc string) error {
	_, err := m.Client.Bucket.Put(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	return nil
}

//TODO 返回何种数据？
func (m *defaultTenCosClient) GetBucket(bucketName string, loc string) error {
	s, _, err := m.Client.Service.Get(context.Background())
	if err != nil {
		panic(err)
	}

	for _, b := range s.Buckets {
		fmt.Printf("%#v\n", b)
	}
	return nil
}

//TODO 返回对象列表
func (m *defaultTenCosClient) GetObjectList() {
	opt := &cos.BucketGetOptions{
		Prefix:  "test",
		MaxKeys: 3,
	}
	v, _, err := m.Client.Bucket.Get(context.Background(), opt)
	if err != nil {
		panic(err)
	}

	for _, c := range v.Contents {
		fmt.Printf("%s, %d\n", c.Key, c.Size)
	}
}

func (m *defaultTenCosClient) UploadFile(bucketName, objectName, filePath, contentType string) (*cos.Response, error) {
	// 2.通过本地文件上传对象
	rsp, err := m.Client.Object.PutFromFile(context.Background(), objectName, filePath, nil)
	if err != nil {
		panic(err)
	}
	return rsp, nil
}
func (m *defaultTenCosClient) DownloadFile(bucketName, objectName, filePath string) error {
	//获取对象到本地文件
	_, err := m.Client.Object.GetToFile(context.Background(), objectName, filePath, nil)
	if err != nil {
		panic(err)
	}
	return nil
}

//删除文件？
func (m *defaultTenCosClient) DeleteFile(bucketName, objectName string) error {
	//name := "test/objectPut.go"
	_, err := m.Client.Object.Delete(context.Background(), objectName)
	if err != nil {
		panic(err)
	}

	return nil
}
func (m *defaultTenCosClient) DeleteFiles(bucketName string, objectInfos []minio.ObjectInfo) error {
	return nil
}
