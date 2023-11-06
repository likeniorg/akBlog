# 环境安装
## Go Install
安装Go语言环境，环境变量如果不生效（shell中显示go命令不存在），重启电脑或注销当前用户重新登录
```bash
wget https://dl.google.com/go/go1.21.3.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.21.3.linux-amd64.tar.gz
echo "export PATH=$PATH:/usr/local/go/bin" > ~/.profile
export PATH=$PATH:/usr/local/go/bin
```
## 常用Linux系统安装hugo 
* debian
    sudo apt install hugo
* centos
    sudo yum install hugo
* archlinux
    sudo pacman -Sy hugo

## 初始化使用
```bash
# hugo/public中内容移动到web目录
./tool/buildToWeb
# 编译akBlog
go build .
# 执行akBlog
./akBlog
```

# 设计理念
## 前端
前端需要一个优美页面来展示文章，否则无法吸引用户，但是从头写页面过于费时，所以采用hugo生成静态网页显示博客内容
## 后端
相较于UI界面，我更喜欢在shell中执行操作(绝对不是懒得写界面)，所以后台操作通过akClient程序来与服务器进行交互