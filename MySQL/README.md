#此项目用于监控远端的rds并通过grafana展示相关监控项

1 需要有一台安装了docker的机器，并能够网络上和rds或者远端mysql网络互通，白名单账号授权做好
	
	执行 docker_start_node.sh 如果有多个可以继续按照同样的格式添加
	
2 prometheus的server端增加这台docker机器的ip端口配置文件

3 导入到grafana目录下的json文件即可查看
