package oss

type OSSConfig struct {
	ENDPOINT        string `yaml:"ENDPOINT"`
	AccessKeyId     string `yaml:"ACCESS_KEY_ID"`
	AccessKeySecret string `yaml:"ACCESS_KEY_SECRET"`
	BucketName      string `yaml:"BUCKET_NAME"`
}
