# bistoury-sync
一个简单的web api服务，旨在解决k8s环境下使用bistoury的问题。因为k8s环境下pod的ip是动态变化的，需要经常去修改bistoury-ui里面的agent配置信息。

此工具已在内部使用半年时间，稳定运行。


#### 解决办法
在pod启动过程中启动一个shell脚本,脚本上报pod的主机名和ip信息,然后将这些信息写入bistoury数据库中。
当pod销毁时使用钩子触发注销接口，从数据库中删除此注册信息。

此方法非最优解决方案，推荐方法是改造java项目来集成注册中心，参考 https://www.cnblogs.com/xiaoqi/p/Bistoury.html

示例脚本文件: `scripts/bistoury-start.sh`
- 启动agent并注册 `bistoury-start.sh start`
- 注销 `bistoury-start.sh del`

#### k8s停止钩子使用示例
```yaml
imagePullPolicy: IfNotPresent
    lifecycle:
      preStop:
        exec:
          command:
          - /bin/sh
          - -c
          - kill -15 $(ps aux|grep java|grep -Ev 'grep|bistoury' |awk '{print $1}')
            ; /bistoury-start.sh del
    livenessProbe:
```

#### 注册API调用
```shell
    curl --location --request POST "${BISTOURY_API}/api/app" \
    --header "Content-Type: application/json" \
    --data-raw "{
        \"code\": \"${jobname}\",
        \"Name\": \"${jobname}\",
        \"Group_code\": \"${jobname}\",
        \"ip\":\"$(hostname -i)\",
        \"port\":${port},
        \"logdir\":\"${logdir}\",
        \"hostname\": \"$(hostname)\"
    }"
```
`变量说明`
- jobname 服务名
- port java程序端口号
- logdir java程序日志目录
- hostname pod内主机名

>POST方法注册实例,DELETE方法注销
#### 程序编译
```shell
cd cmd/bistoury-sync
go build
```

#### 配置文件
```
[server]
host=0.0.0.0
port=8090

[mysql]
host=x.x.x.x:3306
user=bistoury
password=123456
db=bistoury
```

#### 启动
```
./bistoury-sync
```

自动注册应用
![image](https://github.com/typ431127/bistoury-tool/assets/20376675/67862267-3ead-43b3-ab00-4365ef18964f)
![image](https://github.com/typ431127/bistoury-tool/assets/20376675/75796729-1b7e-4a36-b34f-0f1899271579)


