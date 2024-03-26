package utils

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(Config)

type Config struct {
	Port         int `mapstructure:"port"`
	*Datasource  `mapstructure:"datasource"`
	*RedisConfig `mapstructure:"redis"`
	*AwsConfig   `mapstructure:"aws"`
	*Mq          `mapstructure:"mq"`
}
type AwsConfig struct {
	AccessKeyID        string    `mapstructure:"access_key_id"`
	SecretAccessKey    string    `mapstructure:"secret_access_key"`
	Region             string    `mapstructure:"region"`
	FileS3Config       *S3Config `mapstructure:"file_s3"`
	FrontS3Config      *S3Config `mapstructure:"front_s3"`
	FrontAdminS3Config *S3Config `mapstructure:"front_admin_s3"`
}

type RedisConfig struct {
	Host         string `mapstructure:"host"`
	Password     string `mapstructure:"password"`
	Port         int    `mapstructure:"port"`
	DB           int    `mapstructure:"db"`
	PoolSize     int    `mapstructure:"pool_size"`
	MinIdleConns int    `mapstructure:"min_idle_conns"`
}

type S3Config struct {
	BucketName string `mapstructure:"bucket_name"`
	DomainName string `mapstructure:"domain_name"`
	DistId     string `mapstructure:"dist_id"`
}

type Datasource struct {
	*Master `mapstructure:"master"`
	*Slave  `mapstructure:"slave"`
}
type Master struct {
	DriverName   string `mapstructure:"driver_name"`
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DB           string `mapstructure:"dbname"`
	Port         int    `mapstructure:"port"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

type Slave struct {
	Count        int      `mapstructure:"count"`
	DriverName   string   `mapstructure:"driver_name"`
	Hosts        []string `mapstructure:"hosts"`
	Users        []string `mapstructure:"users"`
	Passwords    []string `mapstructure:"passwords"`
	DBs          []string `mapstructure:"dbnames"`
	Ports        []int    `mapstructure:"ports"`
	MaxOpenConns int      `mapstructure:"max_open_conns"`
	MaxIdleConns int      `mapstructure:"max_idle_conns"`
}

type Amqp struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

type Mq struct {
	*Amqp `mapstructure:"amqp"`
}

func Init(filePath string) {
	// 方式1：直接指定配置文件路径（相对路径或者绝对路径）
	// 相对路径：相对执行的可执行文件的相对路径
	//viper.SetConfigFile("./etc/attachment-api.yaml")
	// 绝对路径：系统中实际的文件路径
	//viper.SetConfigFile("/Users/xxx/Desktop/bluebell/etc/attachment-api.yaml")

	// 方式2：指定配置文件名和配置文件的位置，viper自行查找可用的配置文件
	// 配置文件名不需要带后缀
	// 配置文件位置可配置多个
	//viper.SetConfigName("config") // 指定配置文件名（不带后缀）
	//viper.AddConfigPath(".") // 指定查找配置文件的路径（这里使用相对路径）
	//viper.AddConfigPath("./etc")      // 指定查找配置文件的路径（这里使用相对路径）

	// 基本上是配合远程配置中心使用的，告诉viper当前的数据使用什么格式去解析
	//viper.SetConfigType("json")

	viper.SetConfigFile(filePath)

	err := viper.ReadInConfig() // 读取配置信息
	if err != nil {
		// 读取配置信息失败
		fmt.Printf("viper.ReadInConfig failed, err:%v\n", err)
		panic(err)
		return
	}

	// 把读取到的配置信息反序列化到 Conf 变量中
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
		panic(err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了...")
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
			panic(err)
		}
	})
	return
}
