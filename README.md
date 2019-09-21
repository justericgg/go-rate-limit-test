# Go: Rate Limiter 練習 

## 啟用方式

### Heroku 位置(直接訪問即可)

https://go-rate-limit-test.herokuapp.com/rate-limit

### Docker

> $ docker-compose up -d

可以修改 docker-compose 內的 environment 設定來控制連線數量以及時間區間

```
LIMIT=60 // 這是連線數量
WINDOW_TIME_SEC=60 // 這是時間區間
```

#### 測試方式

可以將 header 中的 X-Forwarded-For 更改 ip 來做不同 ip 的 token 取得的測試

```cassandraql
curl -X GET --header "X-Forwarded-For: 1.1.1.1" http://localhost:8080/rate-limit
curl -X GET --header "X-Forwarded-For: 1.1.1.2" http://localhost:8080/rate-limit
```

### 直接 build or run

> go run <Project>/cmd/http/main.go [PORT] [LIMIT] [WINDOW_TIME_SEC]
> Ex: go run <Project>/cmd/http/main.go 8888 60 60

## 目的

- 簡單實作一個 Rate Limiter 的 Gateway

- 每個 IP 每分鐘僅能接受 n個 requests

- 在首頁顯示目前的 ip, ip 在 1分鐘 內取到的 第n個 request, 超過限制的話則顯示 "Error"
request 則顯示 Error

## 選擇

因為實作練習, 為了方便 Demo 暫不考慮用任何儲存方案, 直接以 in-memory 方式實作

限流服務設計上, 個人目前了解的程度, 基本上考量兩種做法

1. Leaky Bucket
2. Token Bucket

前者較適合做流速的控制, 後者適合做時間區段內量的控制, 故這邊選用後者來實作練習

## 程式說明

實作 Token Bucket 當下想到的做法有:

1. 程式啟動一個 background job (go-routine) 定時填充 token 到 bucket 內

> 這裡帶出一個問題是, 有可能會有該分鐘最後一秒進來 n個 request 而下一分鐘第一秒又進來 n個 request 的問題, 這會造成 1分鐘內有 2n個 request 進來

> 而另一個問題是, 每個 ip 需要建立 1個 bucket 來計算, 是否會造成太多的 background job 而造成問題

2. 建立 ip 的 map 來存放各 token bucket, 每個 request 進來的時候會在 bucket 內記錄最後取到 token 的時間, 當目前進來的時間 - 上次取 token 的時間超過 1分鐘的時候, 填充 token

選擇第二種做法

## 程式結構

![image](https://github.com/justericgg/go-rate-limit-test/blob/master/assets/images/ratelimiter_diagram.png)

## 沒有考慮到的部分

以下為因為 Demo 沒有考慮到很細節的部分

- 單純用一個 endpoint 來做 demo, 沒有實作在 middleware 中
- 沒有實作取到 token 後將 request proxy 到後面的 server 這段
- ip 的 map 有可能會很久沒有相同的 ip request 再進來, 沒有特別做一個 background 做清理的動作
- ip 的空間沒有仔細計算過單台機器是否存的下的問題, 若不行可能要考慮前面一台 HA 做分流
