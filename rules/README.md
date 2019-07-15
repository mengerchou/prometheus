用于添加prometheus的rule规则，格式统一为rule_*.yml

第一类为recoarding_rule

example如下：

```yaml
groups:
  - name: grpc_start_total_combine
    interval: 1m
    rules:
    - record: job:grpc_server_started_total:sum
      expr: sum(grpc_server_started_total) by (job)
      labels:
         grpc_start_total_combine: total
```

第二类为alert_rule：

example如下：

```yaml
groups:
- name: MySQLStatsAlert
  rules:
  - alert: MySQL is down
    expr: mysql_up == 0
    for: 1m
    labels:
      severity: critical
    annotations:
      summary: "Instance {{ $labels.instance }} MySQL is down"
      description: "MySQL database is down. This requires immediate action!"
```