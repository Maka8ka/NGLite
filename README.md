## NGLite
* 基于区块链网络的匿名跨平台远控程序
* 实现原理:[链接](https://maka8ka.github.io/post/一个基于区块链网络的匿名远控/)

## 优势&劣势

理论上完全的匿名性，当然要是有人监测并分析了所有中间节点除外，目前节点约8W个

无需任何公网资源，只需要通信主机能上网即可

无需实名购买IP/域名/服务器/CDN等等资源

目前免杀性能优

连接稍多，体积较大，大家可自行通过upx等进行压缩

## 目前支持参数

```powershell
控制端
-n new 生成新的频道/群组/seed
-g 9e8124591f55d27b48ba907f2ad39e790ec589b3942dec2a19e7c2a96b751922 指定9e8124591f55d27b48ba907f2ad39e790ec589b3942dec2a19e7c2a96b751922频道
$mac$ip shell 对执行主机发送shell命令

被控端
-g 9e8124591f55d27b48ba907f2ad39e790ec589b3942dec2a19e7c2a96b751922 指定9e8124591f55d27b48ba907f2ad39e790ec589b3942dec2a19e7c2a96b751922频道
```

## Example
![example](https://raw.githubusercontent.com/Maka8ka/NGLite/main/example.png)

## 后续开发

介于P2P的特性以及分布的7w多个网络节点，这个网络天生具有大文件传输的优势及网络传输速度的优势。

后续考虑增加文件传输、内网穿透代理等功能。

代码写的比较乱，稍后整理一下，将功能分离后会更新源代码。

## Detection
![detection](https://raw.githubusercontent.com/Maka8ka/NGLite/main/detection.png)


## Warning!!!
该远控程序是方便运维人员使用的运维工具，请大家合理使用，产生一切非法行为及后果由使用者自负。

## My Sapce
[Maka8ka's Garden](https://maka8ka.github.io)
