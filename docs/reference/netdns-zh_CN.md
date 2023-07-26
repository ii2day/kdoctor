# netdns

## 基本描述

对于这种任务，每个 kdoctor 代理都会向指定的目标发送 Dns 请求，并获得成功率和平均延迟。 它可以指定成功条件来告知结果成功或失败。

## 配置说明

### 集群 Dns 检查任务

```shell

apiVersion: kdoctor.io/v1beta1
kind: Netdns
metadata:
  creationTimestamp: "2023-07-26T06:51:44Z"
  generation: 1
  name: netdns-e2e
  resourceVersion: "205644"
  uid: 2ded6338-f305-4a87-897c-a680c5ca1139
spec:
  expect:
    meanAccessDelayInMs: 1500
    successRate: 1
  request:
    domain: netdns-e2e.kubernetes.default.svc.cluster.local
    durationInSecond: 10
    perRequestTimeoutInMS: 1000
    protocol: udp
    qps: 10
  schedule:
    roundNumber: 1
    roundTimeoutMinute: 1
    schedule: 0 1
  target:
    enableLatencyMetric: false
    targetDns:
      serviceName: kube-dns
      serviceNamespace: kube-system
      testIPv4: true
      testIPv6: true
status:
  doneRound: 1
  expectedRound: 1
  finish: true
  history:
  - deadLineTimeStamp: "2023-07-26T06:52:45Z"
    duration: 15.138758142s
    endTimeStamp: "2023-07-26T06:52:00Z"
    expectedActorNumber: 2
    failedAgentNodeList: []
    notReportAgentNodeList: []
    roundNumber: 1
    startTimeStamp: "2023-07-26T06:51:45Z"
    status: succeed
    succeedAgentNodeList:
    - kdoctor-worker
    - kdoctor-control-plane
  lastRoundStatus: succeed
```

### 指定 dns server 检查任务

```shell

apiVersion: kdoctor.io/v1beta1
kind: Netdns
metadata:
  creationTimestamp: "2023-07-26T06:51:45Z"
  generation: 1
  name: netdns-e2e
  resourceVersion: "205645"
  uid: 6dc206ee-cbcd-4fb8-8fdd-f060ecb7d8a1
spec:
  expect:
    meanAccessDelayInMs: 1500
    successRate: 1
  request:
    domain: netdns-e2e.kubernetes.default.svc.cluster.local
    durationInSecond: 10
    perRequestTimeoutInMS: 1000
    protocol: udp
    qps: 10
  schedule:
    roundNumber: 1
    roundTimeoutMinute: 1
    schedule: 0 1
  target:
    enableLatencyMetric: false
    targetUser:
      port: 53
      server: 172.41.54.83
status:
  doneRound: 1
  expectedRound: 1
  finish: true
  history:
  - deadLineTimeStamp: "2023-07-26T06:52:45Z"
    duration: 15.177016659s
    endTimeStamp: "2023-07-26T06:52:00Z"
    expectedActorNumber: 2
    failedAgentNodeList: []
    notReportAgentNodeList: []
    roundNumber: 1
    startTimeStamp: "2023-07-26T06:51:45Z"
    status: succeed
    succeedAgentNodeList:
    - kdoctor-worker
    - kdoctor-control-plane
  lastRoundStatus: succeed
  
```

* spec.schedule: 如果调度任务.

      roundNumber: 任务需要执行的轮数，如果为-1表示周期性执行

      schedule: 任务开始时间，支持 linux crontab 写法 与 sample 写法
                linux crontab ： * 1 * * *
                sample : 1 1 ，第一位表示多少分钟之后开启任务，第二位表示每一轮任务的间隔时间。
      roundTimeoutMinute: 每轮任务的超时时间，如果到达超时时间任务还没有结束，判定为任务失败。

* spec.request: 如何对目标地址进行请求

      durationInSecond: 每轮任务的请求时间

      perRequestTimeoutInMS: 每个请求的超时时间，达到超时时间，判定该请求失败

      qps: 每一个 agent 每秒请求数量

      domain： dns 请求解析的域名

      >注意: 使用 agent 请求时，所有的 agent 都会向目标地址进行请求，因此实际的 qps 会大于设置的 qps。

* spec.target: 请求目标设置 targetUser 和 targetDns 不可以同时设置

      enableLatencyMetric: 是否开启延时矩阵，若开启后，会统计所有请求的延时分布，同时会增加内存使用量。

      targetUser [optional]: 对用户自定义的 dns server 进行 dns 请求

        server: dns server 地址

        port:  dns server 端口

      targetDns: [optional]: 对集群的 dns server（coredns）进行 dns 请求 

        testIPv4: 请求 dns server 的 ipv4 地址并且请求 A 记录 

        testIPv6: 请求 dns server 的 ipv6 地址并且请求 AAAA 记录 

        serviceName: 集群内 dns service 的名称

        serviceNamespace: 集群内 dns service 的命名空间
 
        protocol: 请求协议,支持协议 udp、tcp、tcp-tls,默认使用 udp.

* spec.expect: 定义任务成功条件

      meanAccessDelayInMs: 平均延时，如果实际平均延时大于设置的平均延时，任务失败。

      successRate: 所有 http 请求的成功率，如果实际成功率小于设置的成功率，任务失败。

* status: 任务状态

      doneRound: 完成的任务轮数

      expectedRound: 期望执行的任务轮数

      finish: 是否完成所有轮任务

      lastRoundStatus: 最近一轮的任务结果

  history: 任务执行历史只显示近十轮任务

        roundNumber: 任务轮数

        status: 任务状态

        startTimeStamp: 本轮任务开始时间

        endTimeStamp: 本轮任务结束时间

        duration: 本轮任执行时间

        deadLineTimeStamp: 本轮任务 deadline

        failedAgentNodeList: 任务失败的 agent

        notReportAgentNodeList: 没有任务报告的 agent

        succeedAgentNodeList: 任务成功的 agent


## 使用例子

### 集群内 dns server 检查

验证整个网络是否正常，每个 agent 都可以到达集群内 dns server

```shell

cat <<EOF > netdns1.yaml
apiVersion: kdoctor.io/v1beta1
kind: Netdns
metadata:
  name: netdns-e2e-5144-852295818
spec:
  expect:
    meanAccessDelayInMs: 1500
    successRate: 1
  request:
    domain: kubernetes.default.svc.cluster.local
    durationInSecond: 10
    perRequestTimeoutInMS: 1000
    protocol: udp
    qps: 10
  schedule:
    roundNumber: 1
    roundTimeoutMinute: 1
    schedule: 0 1
  target:
    enableLatencyMetric: false
    targetDns:
      serviceName: kube-dns
      serviceNamespace: kube-system
      testIPv4: true
      testIPv6: true
EOF

kubectl apply -f netdns1.yaml

```

### 指定 dns server 检查


```shell

cat <<EOF > netdns2.yaml
apiVersion: kdoctor.io/v1beta1
kind: Netdns
metadata:
  generation: 1
  name: netdns-e2e-5145-9594432
spec:
  expect:
    meanAccessDelayInMs: 1500
    successRate: 1
  request:
    domain: netdns-e2e-5145-9594432.kubernetes.default.svc.cluster.local
    durationInSecond: 10
    perRequestTimeoutInMS: 1000
    protocol: udp
    qps: 10
  schedule:
    roundNumber: 1
    roundTimeoutMinute: 1
    schedule: 0 1
  target:
    enableLatencyMetric: false
    targetUser:
      port: 53
      server: 172.41.54.83
EOF

kubectl apply -f netdns2.yaml

```

## 任务报告

当 kdoctor 未启用聚合报告时，所有报告都将打印在 kdoctor agent 的标准输出中，使用以下命令获取其报告。

```shell
kubectl logs -n kdoctor  kdoctor-agent-v4vzx | jq 'select( .TaskName=="netreach.netdns-e2e-5145-9594432" )'
```

当 kdoctor 启用聚合报告时，所有报告将被  kdoctor-controller 收集，并通过 k8s 聚合 api 查看，使用以下命令获取报告。

```shell
kubectl get kdoctorreport netdns-e2e-5145-9594432 -oyaml
```

### 报告详情
```shell
apiVersion: system.kdoctor.io/v1beta1
kind: KdoctorReport
metadata:
  creationTimestamp: null
  name: netdns-e2e-5145-9594432
spec:
  FailedRoundNumber: null
  FinishedRoundNumber: 1
  Report:
  - EndTimeStamp: "2023-07-26T06:51:56Z"
    NodeName: kdoctor-control-plane
    PodName: kdoctor-agent-xkjn4
    ReportType: agent test report
    RoundDuration: 11.009567031s
    RoundNumber: 1
    RoundResult: succeed
    StartTimeStamp: "2023-07-26T06:51:45Z"
    TaskName: netdns.netdns-e2e-5145-9594432
    TaskType: Netdns
    netDNSTask:
      detail:
      - FailureReason: null
        MeanDelay: 1.3009709
        Metrics:
          DNSMethod: udp
          DNSServer: 172.41.54.83:53
          Duration: 11.006400674s
          EndTime: "2023-07-26T06:51:56Z"
          Errors: {}
          FailedCounts: 0
          Latencies:
            Max_inMx: 0
            Mean_inMs: 1.3009709
            Min_inMs: 0
            P50_inMs: 0
            P90_inMs: 0
            P95_inMs: 0
            P99_inMs: 0
          ReplyCode:
            NOERROR: 103
          RequestCounts: 103
          StartTime: "2023-07-26T06:51:45Z"
          SuccessCounts: 103
          TPS: 9.358191024547468
          TargetDomain: netdns-e2e-5145-9594432.kubernetes.default.svc.cluster.local.
        Succeed: true
        SucceedRate: 1
        TargetName: typeA_172.41.54.83:53_netdns-e2e-5145-9594432.kubernetes.default.svc.cluster.local
        TargetProtocol: udp
        TargetServer: 172.41.54.83:53
      succeed: true
      targetNumber: 1
      targetType: kdoctor agent
    netDNSTaskSpec:
      expect:
        meanAccessDelayInMs: 1500
        successRate: 1
      request:
        domain: netdns-e2e-5145-9594432.kubernetes.default.svc.cluster.local
        durationInSecond: 10
        perRequestTimeoutInMS: 1000
        protocol: udp
        qps: 10
      schedule:
        roundNumber: 1
        roundTimeoutMinute: 1
        schedule: 0 1
      target:
        targetUser:
          port: 53
          server: 172.41.54.83
  - EndTimeStamp: "2023-07-26T06:51:56Z"
    NodeName: kdoctor-worker
    PodName: kdoctor-agent-gp5mh
    ReportType: agent test report
    RoundDuration: 11.006289056s
    RoundNumber: 1
    RoundResult: succeed
    StartTimeStamp: "2023-07-26T06:51:45Z"
    TaskName: netdns.netdns-e2e-5145-9594432
    TaskType: Netdns
    netDNSTask:
      detail:
      - FailureReason: null
        MeanDelay: 1.1181818
        Metrics:
          DNSMethod: udp
          DNSServer: 172.41.54.83:53
          Duration: 11.003017817s
          EndTime: "2023-07-26T06:51:56Z"
          Errors: {}
          FailedCounts: 0
          Latencies:
            Max_inMx: 0
            Mean_inMs: 1.1181818
            Min_inMs: 0
            P50_inMs: 0
            P90_inMs: 0
            P95_inMs: 0
            P99_inMs: 0
          ReplyCode:
            NOERROR: 110
          RequestCounts: 110
          StartTime: "2023-07-26T06:51:45Z"
          SuccessCounts: 110
          TPS: 9.99725728245633
          TargetDomain: netdns-e2e-5145-9594432.kubernetes.default.svc.cluster.local.
        Succeed: true
        SucceedRate: 1
        TargetName: typeA_172.41.54.83:53_netdns-e2e-5145-9594432.kubernetes.default.svc.cluster.local
        TargetProtocol: udp
        TargetServer: 172.41.54.83:53
      succeed: true
      targetNumber: 1
      targetType: kdoctor agent
    netDNSTaskSpec:
      expect:
        meanAccessDelayInMs: 1500
        successRate: 1
      request:
        domain: netdns-e2e-5145-9594432.kubernetes.default.svc.cluster.local
        durationInSecond: 10
        perRequestTimeoutInMS: 1000
        protocol: udp
        qps: 10
      schedule:
        roundNumber: 1
        roundTimeoutMinute: 1
        schedule: 0 1
      target:
        targetUser:
          port: 53
          server: 172.41.54.83
  ReportRoundNumber: 1
  RoundNumber: 1
  Status: Finished
  TaskName: netdns-e2e-5145-9594432
  TaskType: Netdns
```