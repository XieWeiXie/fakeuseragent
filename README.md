### Fake UserAgent


![](https://img.shields.io/badge/fake_useragent_go_ng-v1.0.0-519dd9.svg)
![](https://img.shields.io/badge/language-golang-orange.svg)
[![](https://img.shields.io/badge/weibo-@谢小小路-red.svg)](https://weibo.com/1948244870/profile?topnav=1&wvr=6)
[![](https://img.shields.io/badge/jianshu-@谢小路-F59581.svg)](https://www.jianshu.com/u/58f0817209aa)


### 1. Features

从下面几个网站中抓取到 User-Agent 数据。

- [http://useragentstring.com/](http://useragentstring.com/)
- [https://www.w3schools.com/browsers/default.asp](https://www.w3schools.com/browsers/default.asp)
- [http://d2g6u4gh6d9rq0.cloudfront.net/browsers/fake_useragent_0.1.9.json](http://d2g6u4gh6d9rq0.cloudfront.net/browsers/fake_useragent_0.1.9.json)


### 2. User-Agent 的知识

- User-Agent 用户代理
- 是一个特殊的字符串
- 作用是服务器识别客户使用的操作系统及版本、CPU 类型、浏览器及版本、浏览器渲染引擎、浏览器语言、浏览器插件等

```text

User-Agent: Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.181 Safari/537.36
```


### 3. Usage

```text
func main() {
	fakeUserAgent := application.NewFakeUserAgent(true, false, false)
	fmt.Println(fakeUserAgent.Random())
	fmt.Println(fakeUserAgent.Safari())
	fmt.Println(fakeUserAgent.Chrome())
	fmt.Println(fakeUserAgent.IE())
	fmt.Println(fakeUserAgent.Opera())
}


```

结果：

```go
Mozilla/4.0 (compatible; MSIE 6.0; Windows ME) Opera 7.53  [en]
Mozilla/5.0 (Macintosh; U; PPC Mac OS X 10_5_3; en-us) AppleWebKit/525.18 (KHTML, like Gecko) Version/3.1.1 Safari/525.20
Mozilla/5.0 (Windows; U; Windows NT 6.0; en-US) AppleWebKit/532.1 (KHTML, like Gecko) Chrome/4.0.220.1 Safari/532.1
Mozilla/5.0 (Windows NT 6.1; WOW64; Trident/7.0; AS; rv:11.0) like Gecko
Opera/9.64 (X11; Linux x86_64; U; en-GB) Presto/2.1.1

```

#### 3.1 方法

- Random 随机得到一个 UserAgent
- IE/Msie/InternetExplorer 返回IE 浏览器UserAgent
- FF/FireFox 返回 FireFox 浏览器UserAgent
- Google/Chrome 返回 Chrome 浏览器UserAgent
- Opera 返回 Opera 浏览器UserAgent

#### 3.3 参数

- UserAgentStringOk: http://useragentstring.com/ 网站的UserAgent
- CloudFrontNetOk: http://d2g6u4gh6d9rq0.cloudfront.net/browsers/fake_useragent_0.1.9.json 的UserAgent
- CacheOk : 代码内置的随机 50 个UserAgent


参考：[Python 版 fake-useragent](https://github.com/hellysmile/fake-useragent)

### License
MIT ©wuxiaoxiaoshen

