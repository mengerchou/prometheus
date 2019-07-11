```
一：写在前面
现在有许多云厂商，我们的监控只能从平台中取，往往有许多细节万一对方不支持就无法监控，下面来一波实战，让我们能够轻松的监控云上的mysql


1、找一台和rds网络通并且加入了白名单的机器作为中转：机器上需要安装docker
```bash
cat new.sh 
#!/bin/bash
docker run -d \
  -p 9104:9104 \
  -e DATA_SOURCE_NAME="user:passwd@(rm-xxxxxx.mysql.rds.aliyuncs.com)/" \
  prom/mysqld-exporter
```

2、然后找prometheus上更改配置文件：
```yaml
  - job_name: 'mysql-discorvery'
    file_sd_configs: 
      - files:
        - /data/prometheus-2.3.2/conf.d/mysql-discovery.json
 ```
 
mysql-discovery.json内容如下：
```json
cat /data/prometheus-2.3.2/conf.d/mysql-discovery.json
[{"targets": ["yourip:9104"],"labels": {"instance": "xxxxx"}},{"targets": ["yourip:9105"],"labels": {"instance": "xxx"}}]
```

4、导入grafana的dashboard：在mysql目录文件夹下







