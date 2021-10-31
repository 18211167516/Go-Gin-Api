docker环境安装

## 1、docker更换阿里云镜像源

```
1.在windows命令行执行docker-machine ssh [machine-name]进入VM bash

2.sudo vi /var/lib/boot2docker/profile

3.在--label provider=virtualbox的下一行添加
--registry-mirror=https://k5paawyp.mirror.aliyuncs.com

4.重启docker服务：
sudo /etc/init.d/docker restart或者重启VM：exit退出VM bash，
在windows命令行中执行docker-machine restart
5.直接执行命令：ggdocker-machine ssh default "echo 'EXTRA_ARGS=\"--registry-mirror=https://k5paawyp.mirror.aliyuncs.com\"'  sudo tee -a /var/lib/boot2dock
er/profile" docker-machine restart default

```

## 2、基于docker

> [!NOTE]
> 目录结构

```
    ├─config 	       （配置文件目录）
    |  ├─conf.d      （nginx的conf.d）
    |  ├─mysql.cnf   （mysql配置）
    |  ├─redis.conf  （redis配置）
    ├─log 	         （日志目录）
    |  ├─mysql       （mysql日志）
    ├─mysql  	       （mysql数据目录）
    ├─docker-compose.yaml  （docker-composer） 
 
```

> [!NOTE]
> docker-composer.yaml

```
version: "3"
services:
  mysql:
    image: mysql:5.6
    ports:
      - "3306:3306"
    volumes:
      - ./conf/mysql.cnf:/etc/mysql/conf.d/mysql.cnf:ro
      - ./mysql:/var/lib/mysql/new/:rw
      - ./log/mysql/:/var/log/mysql/:rw
    restart: always
    networks:
      - default
    environment:
      MYSQL_ROOT_PASSWORD: "123456"
      TZ: "$TZ"

  redis:
    image: redis:5.0.3-alpine
    ports:
      - "6379:6379"
    volumes:
      - ./conf/redis.conf:/etc/redis.conf:ro
    restart: always
    entrypoint: ["redis-server", "/etc/redis.conf"]
    environment:
      TZ: "TZ=Asia/Shanghai"
    networks:
      - default
  memcached:
    image: memcached:latest
    ports:
      - "11211:11211"
```

## 3、docker-compose 

```
docker-compose up  编译并启动
docker-compose restart 重启容器
```