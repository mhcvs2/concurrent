- 需求

软件开发人员开发完成一个可运行软件之后，需要对软件进行性能评测，包括：

1）运行速度  
2）高负载下，是否能保证正确性  
3）保证正确性的前提下，伸缩性怎样  
4）正常工作的情况下，负载与系统资源使用率的关系  

- 输入参数

|输入|描述|目的|
|---|---|---|
|每秒载荷量|每秒发送的载荷数量|控制载荷发生器向软件发送载荷的频率|
|负载持续时间|持续发送载荷的时间|评估这段时间内软件性能的具体状况以及对系统资源的使用情况|
|处理超时时间|从向软件发出请求到软件返回的响应的最大耗时|计算出在给定每秒载荷量和负载持续时间的情况下软件的正确率|

- 输出参数

|输出|描述|目的|
|---|---|---|
|请求|请求的内容|检查响应内容的正确性|
|响应|响应的内容|检查响应内容的正确性|
|请求处理耗时|从向软件发送请求到软件响应的精准耗时|评估性能|