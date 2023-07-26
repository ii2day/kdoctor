# AppHttpHealthy

## 基本描述 

对于这种任务，每个kdoctor agent都会向指定的目标发送http请求，并获得成功率和平均延迟。它可以指定成功条件来判断结果是否成功。并且，可以通过聚合API获取详细的报告。

## 配置说明
```shell
apiVersion: kdoctor.io/v1beta1
kind: AppHttpHealthy
metadata:
  name: httphealthy
spec:
  request:
    durationInSecond: 10
    perRequestTimeoutInMS: 1000
    qps: 10
  schedule:
    roundNumber: 2
    roundTimeoutMinute: 1
    schedule: 1 1
  expect:
    meanAccessDelayInMs: 10000
    successRate: 1
    statusCode: 200
  target:
    bodyConfigmapName: http-body
    bodyConfigmapNamespace: kube-system
    header:
    - Accept:text/html
    host: https://10.6.172.20:9443
    http2: false
    method: PUT
    tlsSecretName: https-cert
    tlsSecretNamespace: kube-system
    enableLatencyMetric: false
status:
  doneRound: 2
  expectedRound: 2
  finish: true
  history:
  - deadLineTimeStamp: "2023-05-24T08:03:05Z"
    duration: 15.092667522s
    endTimeStamp: "2023-05-24T08:02:20Z"
    expectedActorNumber: 2
    failedAgentNodeList: []
    notReportAgentNodeList: []
    roundNumber: 2
    startTimeStamp: "2023-05-24T08:02:05Z"
    status: succeed
    succeedAgentNodeList:
    - kdoctor-worker
    - kdoctor-control-plane
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

      host:  http 请求地址，例如 service ip, pod ip, service domain ，自定义url

      method: http 请求方法, 只支持 GET POST PUT DELETE CONNECT OPTIONS PATCH HEAD
        
      bodyConfigmapName: http 请求 body 存放的 configmap 名称
 
      bodyConfigmapNamespace: http 请求 body 的 configmap 命名空间
        
      tlsSecretName:  https 请求证书存放的 secret 名称
        
      tlsSecretNamespace: https 请求证书存放的 secret 命名空间

      header:   http 请求头

      http2: 使用使用 http2 协议进行请求开关

      enableLatencyMetric: 是否开启延时矩阵，若开启后，会统计所有请求的延时分布，同时会增加内存使用量。

* spec.expect: 定义任务成功条件

      meanAccessDelayInMs: 平均延时，如果实际平均延时大于设置的平均延时，任务失败。

      successRate: 所有 http 请求的成功率，当 http 状态码 在 200-400 之间为成功请求，如果实际成功率小于设置的成功率，任务失败。
      
      statusCode（选填）: 期望返回的状态码，若请求返回期望以外的状态吗，任务失败。

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

验证整个网络是否正常，每个 agent 都可以到达特定的地址

```shell

cat <<EOF > test-apphttphealthy.yaml
apiVersion: kdoctor.io/v1beta1
kind: AppHttpHealthy
metadata:
  name: apphttphealth-test
spec:
  expect:
    meanAccessDelayInMs: 1500
    successRate: 1
  request:
    durationInSecond: 10
    perRequestTimeoutInMS: 2000
    qps: 10
  schedule:
    roundNumber: 1
    roundTimeoutMinute: 1
    schedule: 0 1
  target:
    enableLatencyMetric: false
    host: http://172.41.15.187:80?task=apphttphealth-test
    http2: false
    method: GET
EOF
kubectl apply -f test-apphttphealthy.yaml

```

## 任务报告

当 kdoctor 未启用聚合报告时，所有报告都将打印在 kdoctor agent 的标准输出中，使用以下命令获取其报告。

```shell
kubectl logs -n kdoctor  kdoctor-agent-v4vzx | jq 'select( .TaskName=="apphttphealthy.apphttphealth-test" )'
```

当 kdoctor 启用聚合报告时，所有报告将被  kdoctor-controller 收集，并通过 k8s 聚合 api 查看，使用以下命令获取报告。

```shell
kubectl get kdoctorreport apphttphealth-test -oyaml
```


### 报告详情
```shell
apiVersion: system.kdoctor.io/v1beta1
kind: KdoctorReport
metadata:
  creationTimestamp: null
  name: apphttphealth-test
spec:
  FailedRoundNumber: null
  FinishedRoundNumber: 1
  Report:
  - EndTimeStamp: "2023-07-26T06:50:41Z"
    HttpAppHealthyTask:
      Detail:
      - MeanDelay: 108.2
        Metrics:
          Duration: 10.088486394s
          EndTime: "2023-07-26T06:50:41Z"
          Errors: {}
          Latencies:
            Max_inMx: 0
            Mean_inMs: 108.2
            Min_inMs: 0
            P50_inMs: 0
            P90_inMs: 0
            P95_inMs: 0
            P99_inMs: 0
          RequestCounts: 100
          StartTime: "2023-07-26T06:50:31Z"
          StatusCodes:
            "200": 100
          SuccessCounts: 100
          TPS: 9.912289722616242
          TotalDataSize: 40143 byte
        Succeed: true
        SucceedRate: 1
        TargetMethod: GET
        TargetName: HttpAppHealthy target
        TargetUrl: https://172.40.1.21/?task=apphttphealth-test
      Succeed: true
      TargetNumber: 1
      TargetType: HttpAppHealthy
    HttpAppHealthyTaskSpec:
      expect:
        meanAccessDelayInMs: 1500
        successRate: 1
      request:
        durationInSecond: 10
        perRequestTimeoutInMS: 2000
        qps: 10
      schedule:
        roundNumber: 1
        roundTimeoutMinute: 1
        schedule: 0 1
      target:
        host: https://172.40.1.21/?task=apphttphealth-test
        http2: true
        method: GET
        tlsSecretName: https-client-cert
        tlsSecretNamespace: ns-4923-24894817
    NodeName: kdoctor-control-plane
    PodName: kdoctor-agent-xkjn4
    ReportType: agent test report
    RoundDuration: 10.113839941s
    RoundNumber: 1
    RoundResult: succeed
    StartTimeStamp: "2023-07-26T06:50:31Z"
    TaskName: apphttphealthy.apphttphealth-test
    TaskType: AppHttpHealthy
  - EndTimeStamp: "2023-07-26T06:50:41Z"
    HttpAppHealthyTask:
      Detail:
      - MeanDelay: 100.24
        Metrics:
          Duration: 10.133478133s
          EndTime: "2023-07-26T06:50:41Z"
          Errors: {}
          Latencies:
            Max_inMx: 0
            Mean_inMs: 100.24
            Min_inMs: 0
            P50_inMs: 0
            P90_inMs: 0
            P95_inMs: 0
            P99_inMs: 0
          RequestCounts: 100
          StartTime: "2023-07-26T06:50:31Z"
          StatusCodes:
            "200": 100
          SuccessCounts: 100
          TPS: 9.868280040428246
          TotalDataSize: 40149 byte
        Succeed: true
        SucceedRate: 1
        TargetMethod: GET
        TargetName: HttpAppHealthy target
        TargetUrl: https://172.40.1.21/?task=apphttphealth-test
      Succeed: true
      TargetNumber: 1
      TargetType: HttpAppHealthy
    HttpAppHealthyTaskSpec:
      expect:
        meanAccessDelayInMs: 1500
        successRate: 1
      request:
        durationInSecond: 10
        perRequestTimeoutInMS: 2000
        qps: 10
      schedule:
        roundNumber: 1
        roundTimeoutMinute: 1
        schedule: 0 1
      target:
        host: https://172.40.1.21/?task=apphttphealth-test
        http2: true
        method: GET
        tlsSecretName: https-client-cert
        tlsSecretNamespace: ns-4923-24894817
    NodeName: kdoctor-worker
    PodName: kdoctor-agent-gp5mh
    ReportType: agent test report
    RoundDuration: 10.174635658s
    RoundNumber: 1
    RoundResult: succeed
    StartTimeStamp: "2023-07-26T06:50:31Z"
    TaskName: apphttphealthy.apphttphealth-test
    TaskType: AppHttpHealthy
  ReportRoundNumber: 1
  RoundNumber: 1
  Status: Finished
  TaskName: apphttphealth-test
  TaskType: AppHttpHealthy
```
