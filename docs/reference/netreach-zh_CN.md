# Nethttp

## 基本描述 

对于这种任务，每个kdoctor agent都会相互发送http请求，请求地址为每一个 agent 的 pod ip 、service ip、ingress ip 等等，并获得成功率和平均延迟。它可以指定成功条件来判断结果是否成功。并且，可以通过聚合API获取详细的报告。

## 配置说明

```shell
apiVersion: kdoctor.io/v1beta1
kind: NetReach
metadata:
  creationTimestamp: "2023-07-26T06:51:00Z"
  generation: 1
  name: netreach-test
  resourceVersion: "205202"
  uid: 2dcee736-3225-4025-9664-5dedf349e16d
spec:
  expect:
    meanAccessDelayInMs: 1500
    successRate: 1
  request:
    durationInSecond: 10
    perRequestTimeoutInMS: 1000
    qps: 10
  schedule:
    roundNumber: 1
    roundTimeoutMinute: 1
    schedule: 0 1
  target:
    clusterIP: true
    enableLatencyMetric: false
    endpoint: true
    ingress: true
    ipv4: true
    ipv6: true
    loadBalancer: true
    multusInterface: false
    nodePort: true
status:
  doneRound: 1
  expectedRound: 1
  finish: true
  history:
  - deadLineTimeStamp: "2023-07-26T06:52:00Z"
    duration: 15.553875143s
    endTimeStamp: "2023-07-26T06:51:15Z"
    expectedActorNumber: 2
    failedAgentNodeList: []
    notReportAgentNodeList: []
    roundNumber: 1
    startTimeStamp: "2023-07-26T06:51:00Z"
    status: succeed
    succeedAgentNodeList:
    - kdoctor-control-plane
    - kdoctor-worker
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

      >注意: 使用 agent 请求时，所有的 agent 都会向目标地址进行请求，因此实际的 qps 会大于设置的 qps。

* spec.target: 请求目标设置

      clusterIP: kdoctor agent 向的集群IPv4或IPv6的  kdoctor agent cluster IP 地址发送 http 请求。
      
      endpoint: kdoctor agent 向的集群IPv4或IPv6的  kdoctor agent pod IP 地址发送 http 请求。
      
      multusInterface: 若 kdoctor agent 中有多个网卡，是否向所有的网卡发送 http 请求。
      
      ipv4: 测试 ipv4 地址，注意：kdoctor 的配置 configmap 中 要开启 enableIPv4。
      
      ipv6: 测试 ipv6 地址，注意：kdoctor 的配置 configmap 中 要开启 enableIPv6。
      
      ingress: 向 kdoctor agent 的 ingress ipv4 或 ipv6 地址发送请求。
      
      nodePort: 根据ipv4和ipv6，向 kdoctor agnent的每个本地节点的 nodePort ipv4或ipv6地址发送http请求。
      
      enableLatencyMetric: 是否开启延时矩阵，若开启后，会统计所有请求的延时分布，同时会增加内存使用量。 


* spec.expect: 定义任务成功条件

      meanAccessDelayInMs: 平均延时，如果实际平均延时大于设置的平均延时，任务失败。

      successRate: 所有 http 请求的成功率，当 http 状态码 在 200-400 之间为成功请求，如果实际成功率小于设置的成功率，任务失败。

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

验证整个集群网络是否正常，每个 agent 都可以到达集群每一个角落

```shell

cat <<EOF > netreachhealthy-test.yaml
apiVersion: kdoctor.io/v1beta1
kind: NetReach
metadata:
  name: netreach-test
spec:
  expect:
    meanAccessDelayInMs: 1500
    successRate: 1
  request:
    durationInSecond: 10
    perRequestTimeoutInMS: 1000
    qps: 10
  schedule:
    roundNumber: 1
    roundTimeoutMinute: 1
    schedule: 0 1
  target:
    clusterIP: true
    enableLatencyMetric: false
    endpoint: true
    ingress: true
    ipv4: true
    ipv6: true
    loadBalancer: true
    multusInterface: false
    nodePort: true
EOF
kubectl apply -f netreachhealthy-test.yaml

```

## 任务报告

当 kdoctor 未启用聚合报告时，所有报告都将打印在 kdoctor agent 的标准输出中，使用以下命令获取其报告。

```shell
kubectl logs -n kdoctor  kdoctor-agent-v4vzx | jq 'select( .TaskName=="apphttphealthy.netreach-test" )'
```

当 kdoctor 启用聚合报告时，所有报告将被  kdoctor-controller 收集，并通过 k8s 聚合 api 查看，使用以下命令获取报告。

```shell
kubectl get kdoctorreport netreach-test -oyaml
```

### 报告详情
```shell
apiVersion: system.kdoctor.io/v1beta1
kind: KdoctorReport
metadata:
  creationTimestamp: null
  name: netreach-test
spec:
  FailedRoundNumber: null
  FinishedRoundNumber: 1
  Report:
  - EndTimeStamp: "2023-07-26T06:51:11Z"
    NetReachTask:
      Detail:
      - MeanDelay: 9.68
        Metrics:
          Duration: 10.029926681s
          EndTime: "2023-07-26T06:51:10Z"
          Errors: {}
          Latencies:
            Max_inMx: 0
            Mean_inMs: 9.68
            Min_inMs: 0
            P50_inMs: 0
            P90_inMs: 0
            P95_inMs: 0
            P99_inMs: 0
          RequestCounts: 100
          StartTime: "2023-07-26T06:51:00Z"
          StatusCodes:
            "200": 100
          SuccessCounts: 100
          TPS: 9.970162612398065
          TotalDataSize: 34989 byte
        Succeed: true
        SucceedRate: 1
        TargetMethod: GET
        TargetName: AgentPodV4IP_kdoctor-agent-gp5mh_172.40.1.17
        TargetUrl: http://172.40.1.17:80
      - MeanDelay: 93.62
        Metrics:
          Duration: 10.143304754s
          EndTime: "2023-07-26T06:51:10Z"
          Errors: {}
          Latencies:
            Max_inMx: 0
            Mean_inMs: 93.62
            Min_inMs: 0
            P50_inMs: 0
            P90_inMs: 0
            P95_inMs: 0
            P99_inMs: 0
          RequestCounts: 100
          StartTime: "2023-07-26T06:51:00Z"
          StatusCodes:
            "200": 100
          SuccessCounts: 100
          TPS: 9.858719857605099
          TotalDataSize: 35557 byte
        Succeed: true
        SucceedRate: 1
        TargetMethod: GET
        TargetName: AgentLoadbalancerV6IP_fc00:f853:ccd:e793::50:80
        TargetUrl: http://fc00:f853:ccd:e793::50:80
      - MeanDelay: 99.84
        Metrics:
          Duration: 10.126165561s
          EndTime: "2023-07-26T06:51:10Z"
          Errors: {}
          Latencies:
            Max_inMx: 0
            Mean_inMs: 99.84
            Min_inMs: 0
            P50_inMs: 0
            P90_inMs: 0
            P95_inMs: 0
            P99_inMs: 0
          RequestCounts: 100
          StartTime: "2023-07-26T06:51:00Z"
          StatusCodes:
            "200": 100
          SuccessCounts: 100
          TPS: 9.875406381379033
          TotalDataSize: 35001 byte
        Succeed: true
        SucceedRate: 1
        TargetMethod: GET
        TargetName: AgentPodV4IP_kdoctor-agent-xkjn4_172.40.0.12
        TargetUrl: http://172.40.0.12:80
      - MeanDelay: 111.49
        Metrics:
          Duration: 10.12720135s
          EndTime: "2023-07-26T06:51:10Z"
          Errors: {}
          Latencies:
            Max_inMx: 0
            Mean_inMs: 111.49
            Min_inMs: 0
            P50_inMs: 0
            P90_inMs: 0
            P95_inMs: 0
            P99_inMs: 0
          RequestCounts: 100
          StartTime: "2023-07-26T06:51:00Z"
          StatusCodes:
            "200": 100
          SuccessCounts: 100
          TPS: 9.874396345442465
          TotalDataSize: 35491 byte
        Succeed: true
        SucceedRate: 1
        TargetMethod: GET
        TargetName: AgentNodePortV6IP_fc00:f853:ccd:e793::2_30893
        TargetUrl: http://fc00:f853:ccd:e793::2:30893
      - MeanDelay: 90.35
        Metrics:
          Duration: 10.128250564s
          EndTime: "2023-07-26T06:51:10Z"
          Errors: {}
          Latencies:
            Max_inMx: 0
            Mean_inMs: 90.35
            Min_inMs: 0
            P50_inMs: 0
            P90_inMs: 0
            P95_inMs: 0
            P99_inMs: 0
          RequestCounts: 100
          StartTime: "2023-07-26T06:51:00Z"
          StatusCodes:
            "200": 100
          SuccessCounts: 100
          TPS: 9.873373428915892
          TotalDataSize: 34885 byte
        Succeed: true
        SucceedRate: 1
        TargetMethod: GET
        TargetName: AgentNodePortV4IP_172.18.0.2_30431
        TargetUrl: http://172.18.0.2:30431
      - MeanDelay: 9.475728
        Metrics:
          Duration: 11.00557214s
          EndTime: "2023-07-26T06:51:11Z"
          Errors: {}
          Latencies:
            Max_inMx: 0
            Mean_inMs: 9.475728
            Min_inMs: 0
            P50_inMs: 0
            P90_inMs: 0
            P95_inMs: 0
            P99_inMs: 0
          RequestCounts: 103
          StartTime: "2023-07-26T06:51:00Z"
          StatusCodes:
            "200": 103
          SuccessCounts: 103
          TPS: 9.35889553852854
          TotalDataSize: 35923 byte
        Succeed: true
        SucceedRate: 1
        TargetMethod: GET
        TargetName: AgentLoadbalancerV4IP_172.18.0.51:80
        TargetUrl: http://172.18.0.51:80
      - MeanDelay: 87.84615
        Metrics:
          Duration: 11.01140748s
          EndTime: "2023-07-26T06:51:11Z"
          Errors: {}
          Latencies:
            Max_inMx: 0
            Mean_inMs: 87.84615
            Min_inMs: 0
            P50_inMs: 0
            P90_inMs: 0
            P95_inMs: 0
            P99_inMs: 0
          RequestCounts: 104
          StartTime: "2023-07-26T06:51:00Z"
          StatusCodes:
            "200": 104
          SuccessCounts: 104
          TPS: 9.444750835794153
          TotalDataSize: 36201 byte
        Succeed: true
        SucceedRate: 1
        TargetMethod: GET
        TargetName: AgentPodV6IP_kdoctor-agent-gp5mh_fd40:0:0:1::11
        TargetUrl: http://fd40:0:0:1::11:80
      - MeanDelay: 82.68519
        Metrics:
          Duration: 11.012395311s
          EndTime: "2023-07-26T06:51:11Z"
          Errors: {}
          Latencies:
            Max_inMx: 0
            Mean_inMs: 82.68519
            Min_inMs: 0
            P50_inMs: 0
            P90_inMs: 0
            P95_inMs: 0
            P99_inMs: 0
          RequestCounts: 108
          StartTime: "2023-07-26T06:51:00Z"
          StatusCodes:
            "200": 108
          SuccessCounts: 108
          TPS: 9.807130687737077
          TotalDataSize: 37583 byte
        Succeed: true
        SucceedRate: 1
        TargetMethod: GET
        TargetName: AgentClusterV6IP_fd41::7ba5:80
        TargetUrl: http://fd41::7ba5:80
      - MeanDelay: 80.76471
        Metrics:
          Duration: 11.030862964s
          EndTime: "2023-07-26T06:51:11Z"
          Errors: {}
          Latencies:
            Max_inMx: 0
            Mean_inMs: 80.76471
            Min_inMs: 0
            P50_inMs: 0
            P90_inMs: 0
            P95_inMs: 0
            P99_inMs: 0
          RequestCounts: 102
          StartTime: "2023-07-26T06:51:00Z"
          StatusCodes:
            "200": 102
          SuccessCounts: 102
          TPS: 9.246783350757251
          TotalDataSize: 35639 byte
        Succeed: true
        SucceedRate: 1
        TargetMethod: GET
        TargetName: AgentClusterV4IP_172.41.15.186:80
        TargetUrl: http://172.41.15.186:80
      - MeanDelay: 92.836365
        Metrics:
          Duration: 11.016150879s
          EndTime: "2023-07-26T06:51:11Z"
          Errors: {}
          Latencies:
            Max_inMx: 0
            Mean_inMs: 92.836365
            Min_inMs: 0
            P50_inMs: 0
            P90_inMs: 0
            P95_inMs: 0
            P99_inMs: 0
          RequestCounts: 110
          StartTime: "2023-07-26T06:51:00Z"
          StatusCodes:
            "200": 110
          SuccessCounts: 110
          TPS: 9.985338909046002
          TotalDataSize: 38292 byte
        Succeed: true
        SucceedRate: 1
        TargetMethod: GET
        TargetName: AgentPodV6IP_kdoctor-agent-xkjn4_fd40::c
        TargetUrl: http://fd40::c:80
      - MeanDelay: 99.915886
        Metrics:
          Duration: 11.023896937s
          EndTime: "2023-07-26T06:51:11Z"
          Errors: {}
          Latencies:
            Max_inMx: 0
            Mean_inMs: 99.915886
            Min_inMs: 0
            P50_inMs: 0
            P90_inMs: 0
            P95_inMs: 0
            P99_inMs: 0
          RequestCounts: 107
          StartTime: "2023-07-26T06:51:00Z"
          StatusCodes:
            "200": 107
          SuccessCounts: 107
          TPS: 9.706186533808303
          TotalDataSize: 0 byte
        Succeed: true
        SucceedRate: 1
        TargetMethod: GET
        TargetName: AgentIngress_http://172.18.0.50/kdoctoragent
        TargetUrl: http://172.18.0.50/kdoctoragent
      Succeed: true
      TargetNumber: 11
      TargetType: NetReach
    NetReachTaskSpec:
      expect:
        meanAccessDelayInMs: 1500
        successRate: 1
      request:
        durationInSecond: 10
        perRequestTimeoutInMS: 1000
        qps: 10
      schedule:
        roundNumber: 1
        roundTimeoutMinute: 1
        schedule: 0 1
      target:
        clusterIP: true
        endpoint: true
        ingress: true
        ipv4: true
        ipv6: true
        loadBalancer: true
        nodePort: true
    NodeName: kdoctor-control-plane
    PodName: kdoctor-agent-xkjn4
    ReportType: agent test report
    RoundDuration: 11.264679643s
    RoundNumber: 1
    RoundResult: succeed
    StartTimeStamp: "2023-07-26T06:51:00Z"
    TaskName: netreach.netreach-test
    TaskType: NetReach
  - EndTimeStamp: "2023-07-26T06:51:11Z"
    NetReachTask:
      Detail:
      - MeanDelay: 6.9
        Metrics:
          Duration: 10.0414793s
          EndTime: "2023-07-26T06:51:10Z"
          Errors: {}
          Latencies:
            Max_inMx: 0
            Mean_inMs: 6.9
            Min_inMs: 0
            P50_inMs: 0
            P90_inMs: 0
            P95_inMs: 0
            P99_inMs: 0
          RequestCounts: 100
          StartTime: "2023-07-26T06:51:00Z"
          StatusCodes:
            "200": 100
          SuccessCounts: 100
          TPS: 9.958692042516086
          TotalDataSize: 34994 byte
        Succeed: true
        SucceedRate: 1
        TargetMethod: GET
        TargetName: AgentPodV4IP_kdoctor-agent-xkjn4_172.40.0.12
        TargetUrl: http://172.40.0.12:80
      - MeanDelay: 9.22
        Metrics:
          Duration: 10.039099132s
          EndTime: "2023-07-26T06:51:10Z"
          Errors: {}
          Latencies:
            Max_inMx: 0
            Mean_inMs: 9.22
            Min_inMs: 0
            P50_inMs: 0
            P90_inMs: 0
            P95_inMs: 0
            P99_inMs: 0
          RequestCounts: 100
          StartTime: "2023-07-26T06:51:00Z"
          StatusCodes:
            "200": 100
          SuccessCounts: 100
          TPS: 9.961053146815365
          TotalDataSize: 34881 byte
        Succeed: true
        SucceedRate: 1
        TargetMethod: GET
        TargetName: AgentNodePortV4IP_172.18.0.3_30431
        TargetUrl: http://172.18.0.3:30431
      - MeanDelay: 53.51
        Metrics:
          Duration: 10.099769204s
          EndTime: "2023-07-26T06:51:10Z"
          Errors: {}
          Latencies:
            Max_inMx: 0
            Mean_inMs: 53.51
            Min_inMs: 0
            P50_inMs: 0
            P90_inMs: 0
            P95_inMs: 0
            P99_inMs: 0
          RequestCounts: 100
          StartTime: "2023-07-26T06:51:00Z"
          StatusCodes:
            "200": 100
          SuccessCounts: 100
          TPS: 9.901216352587062
          TotalDataSize: 35001 byte
        Succeed: true
        SucceedRate: 1
        TargetMethod: GET
        TargetName: AgentPodV4IP_kdoctor-agent-gp5mh_172.40.1.17
        TargetUrl: http://172.40.1.17:80
      - MeanDelay: 86.95
        Metrics:
          Duration: 10.110452892s
          EndTime: "2023-07-26T06:51:10Z"
          Errors: {}
          Latencies:
            Max_inMx: 0
            Mean_inMs: 86.95
            Min_inMs: 0
            P50_inMs: 0
            P90_inMs: 0
            P95_inMs: 0
            P99_inMs: 0
          RequestCounts: 100
          StartTime: "2023-07-26T06:51:00Z"
          StatusCodes:
            "200": 100
          SuccessCounts: 100
          TPS: 9.890753764267675
          TotalDataSize: 35448 byte
        Succeed: true
        SucceedRate: 1
        TargetMethod: GET
        TargetName: AgentClusterV6IP_fd41::7ba5:80
        TargetUrl: http://fd41::7ba5:80
      - MeanDelay: 94.9
        Metrics:
          Duration: 10.110119339s
          EndTime: "2023-07-26T06:51:10Z"
          Errors: {}
          Latencies:
            Max_inMx: 0
            Mean_inMs: 94.9
            Min_inMs: 0
            P50_inMs: 0
            P90_inMs: 0
            P95_inMs: 0
            P99_inMs: 0
          RequestCounts: 100
          StartTime: "2023-07-26T06:51:00Z"
          StatusCodes:
            "200": 100
          SuccessCounts: 100
          TPS: 9.891080079959874
          TotalDataSize: 34887 byte
        Succeed: true
        SucceedRate: 1
        TargetMethod: GET
        TargetName: AgentLoadbalancerV4IP_172.18.0.51:80
        TargetUrl: http://172.18.0.51:80
      - MeanDelay: 89.85
        Metrics:
          Duration: 10.098886424s
          EndTime: "2023-07-26T06:51:10Z"
          Errors: {}
          Latencies:
            Max_inMx: 0
            Mean_inMs: 89.85
            Min_inMs: 0
            P50_inMs: 0
            P90_inMs: 0
            P95_inMs: 0
            P99_inMs: 0
          RequestCounts: 100
          StartTime: "2023-07-26T06:51:00Z"
          StatusCodes:
            "200": 100
          SuccessCounts: 100
          TPS: 9.902081853534865
          TotalDataSize: 35839 byte
        Succeed: true
        SucceedRate: 1
        TargetMethod: GET
        TargetName: AgentNodePortV6IP_fc00:f853:ccd:e793::3_30893
        TargetUrl: http://fc00:f853:ccd:e793::3:30893
      - MeanDelay: 95.28
        Metrics:
          Duration: 10.128277339s
          EndTime: "2023-07-26T06:51:10Z"
          Errors: {}
          Latencies:
            Max_inMx: 0
            Mean_inMs: 95.28
            Min_inMs: 0
            P50_inMs: 0
            P90_inMs: 0
            P95_inMs: 0
            P99_inMs: 0
          RequestCounts: 100
          StartTime: "2023-07-26T06:51:00Z"
          StatusCodes:
            "200": 100
          SuccessCounts: 100
          TPS: 9.873347327777001
          TotalDataSize: 35501 byte
        Succeed: true
        SucceedRate: 1
        TargetMethod: GET
        TargetName: AgentPodV6IP_kdoctor-agent-xkjn4_fd40::c
        TargetUrl: http://fd40::c:80
      - MeanDelay: 9.431192
        Metrics:
          Duration: 11.011572265s
          EndTime: "2023-07-26T06:51:11Z"
          Errors: {}
          Latencies:
            Max_inMx: 0
            Mean_inMs: 9.431192
            Min_inMs: 0
            P50_inMs: 0
            P90_inMs: 0
            P95_inMs: 0
            P99_inMs: 0
          RequestCounts: 109
          StartTime: "2023-07-26T06:51:00Z"
          StatusCodes:
            "200": 109
          SuccessCounts: 109
          TPS: 9.898677262143
          TotalDataSize: 38091 byte
        Succeed: true
        SucceedRate: 1
        TargetMethod: GET
        TargetName: AgentClusterV4IP_172.41.15.186:80
        TargetUrl: http://172.41.15.186:80
      - MeanDelay: 28.686274
        Metrics:
          Duration: 11.023513429s
          EndTime: "2023-07-26T06:51:11Z"
          Errors: {}
          Latencies:
            Max_inMx: 0
            Mean_inMs: 28.686274
            Min_inMs: 0
            P50_inMs: 0
            P90_inMs: 0
            P95_inMs: 0
            P99_inMs: 0
          RequestCounts: 102
          StartTime: "2023-07-26T06:51:00Z"
          StatusCodes:
            "200": 102
          SuccessCounts: 102
          TPS: 9.252948314252016
          TotalDataSize: 0 byte
        Succeed: true
        SucceedRate: 1
        TargetMethod: GET
        TargetName: AgentIngress_http://172.18.0.50/kdoctoragent
        TargetUrl: http://172.18.0.50/kdoctoragent
      - MeanDelay: 65.37273
        Metrics:
          Duration: 11.007873525s
          EndTime: "2023-07-26T06:51:11Z"
          Errors: {}
          Latencies:
            Max_inMx: 0
            Mean_inMs: 65.37273
            Min_inMs: 0
            P50_inMs: 0
            P90_inMs: 0
            P95_inMs: 0
            P99_inMs: 0
          RequestCounts: 110
          StartTime: "2023-07-26T06:51:00Z"
          StatusCodes:
            "200": 110
          SuccessCounts: 110
          TPS: 9.99284736967397
          TotalDataSize: 39443 byte
        Succeed: true
        SucceedRate: 1
        TargetMethod: GET
        TargetName: AgentLoadbalancerV6IP_fc00:f853:ccd:e793::50:80
        TargetUrl: http://fc00:f853:ccd:e793::50:80
      - MeanDelay: 76.98039
        Metrics:
          Duration: 11.006378962s
          EndTime: "2023-07-26T06:51:11Z"
          Errors: {}
          Latencies:
            Max_inMx: 0
            Mean_inMs: 76.98039
            Min_inMs: 0
            P50_inMs: 0
            P90_inMs: 0
            P95_inMs: 0
            P99_inMs: 0
          RequestCounts: 102
          StartTime: "2023-07-26T06:51:00Z"
          StatusCodes:
            "200": 102
          SuccessCounts: 102
          TPS: 9.267353082440595
          TotalDataSize: 36215 byte
        Succeed: true
        SucceedRate: 1
        TargetMethod: GET
        TargetName: AgentPodV6IP_kdoctor-agent-gp5mh_fd40:0:0:1::11
        TargetUrl: http://fd40:0:0:1::11:80
      Succeed: true
      TargetNumber: 11
      TargetType: NetReach
    NetReachTaskSpec:
      expect:
        meanAccessDelayInMs: 1500
        successRate: 1
      request:
        durationInSecond: 10
        perRequestTimeoutInMS: 1000
        qps: 10
      schedule:
        roundNumber: 1
        roundTimeoutMinute: 1
        schedule: 0 1
      target:
        clusterIP: true
        endpoint: true
        ingress: true
        ipv4: true
        ipv6: true
        loadBalancer: true
        nodePort: true
    NodeName: kdoctor-worker
    PodName: kdoctor-agent-gp5mh
    ReportType: agent test report
    RoundDuration: 11.257636789s
    RoundNumber: 1
    RoundResult: succeed
    StartTimeStamp: "2023-07-26T06:51:00Z"
    TaskName: netreach.netreach-test
    TaskType: NetReach
  ReportRoundNumber: 1
  RoundNumber: 1
  Status: Finished
  TaskName: netreach-test
  TaskType: NetReach
```
