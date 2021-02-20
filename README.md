Kit 工具包
===

gls
---
记录用户sessionId, actionId, spanID值，用于存储调用

filetools
---
文件操作相关方法

|函数|入参数|返回值|方法说明|
|---|---|---|---|
|CheckFileExist|filename string|bool|检测文件是否存在|
|MakeDir|dirPath string|error|递归创建文件夹|
|ReadFile|filename string|string|读取文件|
|ReadFileByte|filename string|[]byte|读取byte文件|
|WriteFile|filename string, value string|error|写文件（追加）|
|WriteFileCover|filename string, value string|error|写文件（覆盖）|
|WriteFileByte|filename string, value []byte|error|写文件,字符类型|

dlog
---
+ 设置日志信息
```go
	dlog.SetLog(dlog.SetLogConf{
		LogType: dlog.LOG_TYPE_LOCAL,
		LogPath: "/Users/caoshuyu/WorkSpace/GoWork/Csy/src/logs",
		Prefix:  "id_generator",
	})
```
|参数|说明|可选值|默认值|
|---|---|---|---|
|LogType|日志存储类型|LOG_TYPE_LOCAL 本地文件存储,LOG_TYPE_NET 网络存储（暂不支持）|LOG_TYPE_LOCAL|
|LogPath|日志存储位置|实际存储位置|/data/logs|
|Prefix|项目名称|-|-|

+ 调用日志存储

|日志级别|调用方法|
|---|---|
|DEBUG|dlog.DEBUG("funcName", "GetIdUseSnowflake", "info", info)|
|INFO|dlog.INFO("funcName", "GetIdUseSnowflake", "info", info)|
|WARN|dlog.WARN("funcName", "GetIdUseSnowflake", "warn", warn)|
|ERROR|dlog.ERROR("funcName", "GetIdUseSnowflake", "error", err)|

uuid
---

|函数|入参数|返回值|方法说明|
|---|---|---|---|
|NewV1|-|UUID|基于当前时间戳和MAC地址|
|NewV2|domain byte|UUID|返回基于POSIX UID/GID的DCE安全UUID|
|NewV3|ns UUID, name string|UUID|基于命名空间UUID和名称的MD5哈希|
|NewV4|-|UUID|随机生成的UUID|
|NewV5|ns UUID, name string|基于命名空间UUID和名称的SHA-1散列|

stringtools
---

|函数|入参数|返回值|方法说明|
|---|---|---|---|
|InitialUpdateStr|string|string|首字母大写，有_的去掉后第一个字母大写|
|InitialLowStr|string|string|首字母小写，有_的去掉后第一个字母大写|
|UpperToUnderlineToUpper|string|string|变成大写字母，用_分割|
|UpperToUnderline|string|string|变成小写字母，用_分割|
|MakeRoundabout|int|string|制作问号|
|Md5|string|string|MD5加密|
|Time33|string|int64|Time33加密算法|
|IsNum|string|bool|判断字符串是不是数字|
|RemoveSliceRepeatStr|[]string,[]string|[]string|去除数组中重复的字符串|

timetools
---

|函数|入参数|返回值|方法说明|
|---|---|---|---|
|GetConstellation|month int64, day int64|int64|查询星座，具体返回值星座见文件内|
|GetNowDateTime|-|time.Time|获取当前时间|
|GetBeforeTime|beforeSecond int|time.Time|获取当前时间之前时间|
|GetNextTime|nextSecond int|time.Time|获取当前时间之后时间|
|TimeStamp|layout string, value string|time.Time, error|时间类型转时间|
|GetDayStart|time.Time|time.Time|获取零晨时间|
|GetDayEnd|time.Time|time.Time|获取一天最后一刻时间|
|SecondToDateTime|int64|time.Time|时间戳转时间|
|TimeFormat|t time.Time, layout string|string|格式化时间|

|参数|类型|可选值|说明|
|---|---|---|---|
|layout|string|RFC3339/TT/TTDAY/NOSPLITT|时间格式|
|value|string|-|时间值|

#### 调用方法示例
```go
	var (
		t  time.Time
		tn int64
		ts string
	)

	//获取当前时间
	t = timetools.GetNowDateTime()
	fmt.Println("获取当前时间:", t)
	//获取当前时间戳
	tn = timetools.GetNowDateTime().Unix()
	fmt.Println("获取当前时间戳:", tn)
	//获取一小时前时间
	t = timetools.GetBeforeTime(timetools.ONE_HOURS_SECOND)
	fmt.Println("获取一小时前时间:", t)
	//获取一小时后时间
	t = timetools.GetNextTime(timetools.ONE_HOURS_SECOND)
	fmt.Println("获取一小时后时间:", t)
	//获取今天日期 2006-01-02
	ts = timetools.TimeFormat(timetools.TTDAY, timetools.GetNowDateTime())
	fmt.Println("获取今天日期 2006-01-02:", ts)
	//获取今天日期 20060102
	ts = timetools.TimeFormat(timetools.NOSPLITTDAY, timetools.GetNowDateTime())
	fmt.Println("获取今天日期 20060102:", ts)
	//2006-01-02T15:04:05Z07:00 转 时间
	t, err := timetools.TimeStamp(timetools.RFC3339, "2006-01-02T15:04:05Z07:00")
	if nil != err {
		fmt.Println(err)
		return
	}
	fmt.Println("2006-01-02T15:04:05Z07:00 转 时间:", t)

	//2006-01-02T15:04:05Z07:00 转 时间戳
	tn, err = timetools.TimeStampUnix(timetools.RFC3339, "2006-01-02T15:04:05Z07:00")
	if nil != err {
		fmt.Println(err)
		return
	}
	fmt.Println("2006-01-02T15:04:05Z07:00 转 时间戳:", tn)
	//2006-01-02 15:04:05 转 时间戳
	tn, err = timetools.TimeStampUnix(timetools.TT, "2006-01-02 15:04:05")
	if nil != err {
		fmt.Println(err)
		return
	}
	fmt.Println("2006-01-02 15:04:05 转 时间戳:", tn)

	//获取昨日零晨时间
	t = timetools.GetDayStart(timetools.GetBeforeTime(timetools.ONE_DAY_SECOND))
	fmt.Println("昨日零晨时间:", t)

	//获取昨日最后一刻时间
	t = timetools.GetDayEnd(timetools.GetBeforeTime(timetools.ONE_DAY_SECOND))
	fmt.Println("获取昨日最后一刻时间:", t)

	//时间戳转时间 1613802876
	t = timetools.SecondToDateTime(1613802876)
	fmt.Println("时间戳转时间 1613802876:", t)

	//格式化时间
	ts = timetools.TimeFormat(timetools.TT, timetools.GetNowDateTime())
	fmt.Println("格式化时间:", ts)

	_, _, _ = t, tn, ts
```

etcdclient
---
#### 初始化客户端
```go
	//添加一个节点
	//etcdclient.AddEndpoint("127.0.0.1:2379")
	//设置多个节点，默认值为"127.0.0.1:2379"
	etcdclient.SetEndpoints([]string{
		"127.0.0.1:2379",
	})
	//设置超时时间，秒 ，默认值为5
	etcdclient.SetTimeOutSecond(5)
	//设置etcd用户名，没有可忽略
	etcdclient.SetUsername("root")
	//设置etcd密码，没有可忽略
	etcdclient.SetPassword("root")
	//初始化客户端
	etcdclient.InitClient()
```
#### 使用客户端
```go
    //使用客户端读取数据
	etcdclient.GetValue("key")

	//使用客户端监听数据
	listeningChan := make(chan []byte, 100)
	etcdclient.WatchValue("key", listeningChan)
	select {
	case val := <-listeningChan:
		fmt.Println(val)
	}
```




