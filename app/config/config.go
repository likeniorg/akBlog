// 生成全局配置和admin SSl证书
package config

import (
	"akBlog/app/util"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

// 配置文件路径
const cfgFilePath = "config/main.conf"

// 检测是否存在配置文件，如果不存在选择是否新建配置文件
func init() {
	if _, err := os.ReadFile(cfgFilePath); err != nil {
		fmt.Fprintf(os.Stderr, "错误原因：%s\n", err.Error())
		fmt.Println("未检测到配置文件，是否创建(y/n)")
		var is string
		fmt.Scanln(&is)
		if is != "y" {
			fmt.Println("退出成功")
			os.Exit(1)
		}
		SelectMode()
	}
}

// 配置信息
type ConfigInfo struct {
	// 服务器IP
	ServerIP string
	// 主域名
	Domain string
	// 公开web服务端口
	Port string
	// 管理员web服务端口
	AdminPort string
	// 公开web服务是否使用HTTPS
	PubHTTPS string
	// 管理员端口是否使用HTTPS协议
	AdminHTTPS string
	// 公开服务端口证书路径
	PubCaPath string
	// 管理员服务端口证书路径
	AdminCaPath string
}

// 快捷获取配置文件信息
func Get(cfgName string) string {
	if cfgInfo, err := get(cfgName); err == nil {
		return cfgInfo
	} else {
		fmt.Fprintf(os.Stderr, "%s", err.Error())
		return ""
	}

}

// 获取配置文件信息
func get(cfgName string) (string, error) {
	data, err := os.ReadFile(cfgFilePath)
	util.ErrprDisplay(err)

	c := ConfigInfo{}
	err = json.Unmarshal(data, &c)
	util.ErrprDisplay(err)

	switch cfgName {
	case "port":
		return c.Port, nil

	case "serverIP":
		return c.ServerIP, nil

	case "domain":
		return c.Domain, nil

	case "adminPort":
		return c.AdminPort, nil

	case "PubHTTPS":
		return c.PubHTTPS, nil

	case "adminHTTPS":
		return c.AdminHTTPS, nil

	case "pubCaPath":
		return c.PubCaPath, nil

	case "adminCaPath":
		return c.AdminCaPath, nil
	default:
		return "", errors.New("不存在配置：" + cfgName)
	}
}

// 选择部署模式
func SelectMode() {

	//配置文件夹创建
	os.Mkdir("config", 0700)
	os.Mkdir("config/cert/", 0700)
	os.Mkdir("config/cert/Ca", 0700)
	os.Mkdir("config/cert/adminCa", 0700)

	fmt.Println(`选择部署模式(默认本地部署)
	(0) 本地部署
		公开访问地址		 localhost:8080
		管理员访问地址		  localhost:59812
		均不使用HTTPS协议

	(1) 自定义
	`)

	var selectVar string
	fmt.Scanln(&selectVar)
	switch selectVar {
	case "0":
		// 默认配置信息
		configInfo := ConfigInfo{"127.0.0.1", "localhost", ":8080", ":59812", "", "", "config/cert/Ca/", "config/cert/adminCa/"}

		// 将数据写入配置文件
		data, err := json.MarshalIndent(configInfo, "", "	")
		util.ErrprDisplay(err)
		err = os.WriteFile(cfgFilePath, data, 0600)
		util.ErrprDisplay(err)

	case "1":
		createConfig()

	default:
		fmt.Println("输入不合法")
	}
	// 设置配置文件夹为只读
	util.Shell("chmod 400 config/cert/adminCa/*")
	util.Shell("chmod 400 " + cfgFilePath)
}

// akBlog配置创建
func createConfig() {

	// 配置信息
	configInfo := ConfigInfo{}

	//写入配置文件
	fmt.Println("输入你的服务器IP(默认:127.0.0.1)")
	fmt.Scanln(&configInfo.ServerIP)
	if configInfo.ServerIP == "" {
		configInfo.ServerIP = "127.0.0.1"
	}

	fmt.Println("输入你的域名(默认:localhost)")
	fmt.Scanln(&configInfo.Domain)
	if configInfo.Domain == "" {
		configInfo.Domain = "localhost"
	}

	fmt.Println("设置你的web端口(默认:8080)")
	fmt.Scanln(&configInfo.Port)
	if configInfo.Port == "" {
		configInfo.Port = ":8080"
	} else {
		configInfo.Port = ":" + configInfo.Port
	}

	fmt.Println("设置你的网站管理员端口(默认:59812)")
	fmt.Scanln(&configInfo.AdminPort)
	if configInfo.AdminPort == "" {
		configInfo.AdminPort = ":59812"
	} else {
		configInfo.AdminPort = ":" + configInfo.AdminPort
	}

	// 公开端口是否使用HTTPS协议
	fmt.Println(configInfo.Port + "端口是否使用HTTPS协议(y/n)")
	var PubHTTPS string
	fmt.Scanln(&PubHTTPS)
	if PubHTTPS == "y" {
		configInfo.PubHTTPS = "y"
		fmt.Println(configInfo.Port + "端口将使用https协议")
		fmt.Println("!!!")
		fmt.Println("公开服务端口证书路径：config/cert/Ca/")
		fmt.Println("证书文件名格式：\n 域名.crt\n域名.key\n不规范命名将无法正确导入证书")
		fmt.Println("!!!")
	}

	// 管理员端口是否使用HTTPS协议
	fmt.Println(configInfo.AdminPort + "端口是否使用HTTPS协议(y/n)")
	var adminHTTPS string
	fmt.Scanln(&adminHTTPS)
	if PubHTTPS == "y" {
		configInfo.PubHTTPS = "y"
		fmt.Println(configInfo.AdminHTTPS + "端口将使用https协议")
		fmt.Println("!!!")
		fmt.Println("管理员服务端口证书路径：config/cert/adminCa/")
		fmt.Println("!!!")
	}

	// 将数据写入配置文件
	data, err := json.MarshalIndent(configInfo, "", "	")
	util.ErrprDisplay(err)
	err = os.WriteFile(cfgFilePath, data, 0600)
	util.ErrprDisplay(err)

	// 控制台输出配置信息
	fmt.Println("网站管理员端口设置为" + Get("adminPort"))
	fmt.Println("web端口为" + Get("port"))
	fmt.Println("服务器域名是" + Get("domain"))
	fmt.Println("服务器IP是" + Get("serverIP"))

	// 创建管理员证书
	createCA(configInfo.Domain)

}

// 创建管理员https证书
func createCA(doMain string) {
	domain := Get("domain")
	util.Shell("openssl genrsa -out ./config/cert/adminCa/ca.key 4096")
	fmt.Println("生成证书进度：35%")

	util.Shell(`openssl req -x509 -new -nodes -sha512 -days 3650 \
	-subj "/C=CN/ST=Beijing/L=Beijing/O=example/OU=Personal/CN=` + domain + `" \
	-key ./config/cert/adminCa/ca.key \
	-out ./config/cert/adminCa/ca.crt`)
	fmt.Println("生成证书进度：50%")

	util.Shell(`openssl genrsa -out ./config/cert/adminCa/` + domain + `.key 4096`)
	fmt.Println("生成证书进度：65%")

	util.Shell(`openssl req -sha512 -new \
	-subj "/C=CN/ST=Beijing/L=Beijing/O=example/OU=Personal/CN=` + domain + `" \
	-key ./config/cert/adminCa/` + domain + `.key \
	-out ./config/cert/adminCa/` + domain + `.csr`)
	fmt.Println("生成证书进度：80%")

	util.Shell(`cat > ./config/cert/adminCa/v3.ext <<-EOF
	authorityKeyIdentifier=keyid,issuer
	basicConstraints=CA:FALSE
	keyUsage=digitalSignature, nonRepudiation, keyEncipherment, dataEncipherment
	extendedKeyUsage=serverAuth
	subjectAltName=@alt_names
	
	[alt_names]
	DNS.1=` + domain + `
	EOF`)
	fmt.Println("生成证书进度：95%")

	util.Shell(`openssl x509 -req -sha512 -days 3650 \
	-extfile ./config/cert/adminCa/v3.ext \
	-CA ./config/cert/adminCa/ca.crt -CAkey ./config/cert/adminCa/ca.key -CAcreateserial \
	-in ./config/cert/adminCa/` + domain + `.csr \
	-out ./config/cert/adminCa/` + domain + `.crt`)
	fmt.Println("生成证书进度：100%")

	// 删除不需要的证书
	util.Shell("rm -rf ./config/cert/adminCa/*.csr")
	util.Shell("rm -rf ./config/cert/adminCa/v3.ext")
	util.Shell("rm -rf ./config/cert/adminCa/ca.key")
	util.Shell("rm -rf ./config/cert/adminCa/ca.srl")
}
