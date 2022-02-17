
## How to install docker on ali cloud

### Install Docker latest
```bash
wget -O /etc/yum.repos.d/docker-ce.repo https://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo
yum install yum-plugin-releasever-adapter --disablerepo=* --enablerepo=plus
yum -y install docker-ce
```
### Start Docker
```bash
systemctl enable docker
systemctl start docker
systemctl status docker
```

