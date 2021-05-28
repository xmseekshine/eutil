package local

import (
	"context"
	"fmt"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/tal-tech/go-zero/core/logx"
)

//新建存储桶
//上传文件对象
//下载文件对象
//删除文件对象
//批量上传文件对象
//批量下载文件对象
//批量删除文件对象
//根据配置获取对象存储Client
//不存在创建
//https://github.com/minio/minio-go

type (
	LocalCosClient interface {
		BucketExists(bucketName string, loc string) error
		CreateBucket(bucketName string, loc string) error
		UploadFile(bucketName, objectName, filePath, contentType string) (minio.UploadInfo, error)
		DownloadFile(bucketName, objectName, filePath string) error
		DeleteFile(bucketName, objectName string) error
		DeleteFiles(bucketName string, objectInfos []minio.ObjectInfo) error
	}

	defaultLocalCosClient struct {
		Client *minio.Client
	}

	MinIOConfig struct {
		EndPoint        string //对象存储服务的URL
		AccessKeyID     string //Access key是唯一标识你的账户的用户ID。
		SecretAccessKey string //Secret key是你账户的密码。
		UseSSL          bool   //true代表使用HTTPS
	}
)

func NewLocalCosClient(c MinIOConfig) LocalCosClient {
	opt := &minio.Options{
		Creds:  credentials.NewStaticV4(c.AccessKeyID, c.SecretAccessKey, ""),
		Secure: c.UseSSL,
	}
	// 初使化 minio client对象。
	client, err := minio.New(c.EndPoint, opt)
	if err != nil {
		log.Fatalln(err)
	}
	return &defaultLocalCosClient{
		Client: client,
	}
}

func (m *defaultLocalCosClient) BucketExists(bucketName string, loc string) error {
	ctx := context.Background()
	found, err := m.Client.BucketExists(ctx, bucketName)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if found {
		fmt.Println("Bucket found")
	}
	return nil
}

func (m *defaultLocalCosClient) CreateBucket(bucketName string, loc string) error {
	ctx := context.Background()
	err := m.Client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: loc})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := m.Client.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
	}
	return nil
}

func (m *defaultLocalCosClient) UploadFile(bucketName, objectName, filePath, contentType string) (minio.UploadInfo, error) {
	// objectName := "golden-oldies.zip"
	// filePath := "/tmp/golden-oldies.zip"
	// contentType := "application/zip"

	// Upload the zip file with FPutObject
	ctx := context.Background()
	info, err := m.Client.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalln(err)
	}

	return info, nil
}

func (m *defaultLocalCosClient) DownloadFile(bucketName, objectName, filePath string) error {
	err := m.Client.FGetObject(context.Background(), bucketName, objectName, filePath, minio.GetObjectOptions{})
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (m *defaultLocalCosClient) DeleteFile(bucketName, objectName string) error {
	opts := minio.RemoveObjectOptions{
		GovernanceBypass: true,
		VersionID:        "myversionid",
	}
	err := m.Client.RemoveObject(context.Background(), bucketName, objectName, opts)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

//TODO 返回错误的文件列表
func (m *defaultLocalCosClient) DeleteFiles(bucketName string, objectInfos []minio.ObjectInfo) error {
	objectsCh := make(chan minio.ObjectInfo)
	opts := minio.RemoveObjectsOptions{
		GovernanceBypass: true,
	}
	for _, v := range objectInfos {
		objectsCh <- v
	}
	for rErr := range m.Client.RemoveObjects(context.Background(), bucketName, objectsCh, opts) {
		logx.Errorf("rm object %s id %s err %s", rErr.ObjectName, rErr.VersionID, rErr.Err)
		fmt.Println(rErr)
	}
	return nil
}
