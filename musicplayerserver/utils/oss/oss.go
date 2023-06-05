package utils

import (
	"errors"
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var accessKey = "LTAI5tEZPqYCvUwvAcYvUy7W"
var accessKeySceret = "pFVyY9pLBITSsCmADv4EMidrjxuNcl"
var client *oss.Client

//
func CreateClient() {
	var err error
	client, err = oss.New("https://oss-cn-guangzhou.aliyuncs.com", accessKey, accessKeySceret)
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(-1)
    }
}

func UploadFile (ID int, fileHeader *multipart.FileHeader, filetype int) (string, error) {
	var url string
	fmt.Println("进入UploadFile成功！")
	bucket, err := client.Bucket("web-music-player")
    if err != nil {
		err = errors.New("获取OSS存储空间失败！")
		fmt.Println(err)
        return url, err
    }

	prefix := ""
	//根据文件类型设置前缀
	switch filetype{
	case 1://用户头像
		prefix = "image/userpic/"
	case 2:
	
	}
	objectKey := prefix + strconv.Itoa(ID) + filepath.Ext(fileHeader.Filename)//保存文件命名格式“用户/歌曲/专辑ID.文件扩展名”
	fmt.Println("objectKey:"+objectKey)
	// 打开文件内容并将其上传到 OSS
	fileContent, err := fileHeader.Open()
	if err != nil {
		err = errors.New("读取上传文件失败！")
		fmt.Println(err)
		return url, err
	}
	defer fileContent.Close()
	err = bucket.PutObject(objectKey, fileContent)
	if err != nil {
		fmt.Println(err.Error())
		err = errors.New("上传文件至OSS失败！")
		fmt.Println(err)
		return url, err
	}
	url = "https://web-music-player.oss-cn-guangzhou.aliyuncs.com/" + objectKey
	return url, err
}
