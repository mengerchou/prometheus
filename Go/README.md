#此文主要讲解了如何用prometheus自动发现k8s里面的go项目并监控展示


1:首先在k8s里面一次应用四个yaml文件

```
2:其次程序中植入相关go-prometheus相关的插件
	这部分参考：https://github.com/prometheus/client_golang
		    https://github.com/grpc-ecosystem/go-grpc-prometheus
		    目录下的/go-sql-prometheus
```
```yaml  
3: 修改业务deployment的yaml在template下添加
  template:
    metadata:
      annotations:
        prometheus.io/port: "2112"
        prometheus.io/scrape: "true"
```
