# END OF this README
# This readme is old 
# Ethernet Over IP (EoIP)

### 为什么要写这个项目
有以下三点原因:
1. 我需要一个在Infiniband网络的IPOIB上面传输以太网帧, 并且进行桥接.
2. 我需要使用VPS NAT VPS自己的公网IP, 并且进行NAT, 方便进行科学上网(全局),和在VPS上的端口转发.
3. 尝试直接使用VPS进行上网.

### 这个项目的原理
这个项目使用TUN/TAP技术, 支持Linux和Windows(部分支持), 并且可以在100Mbit/s下跑满带宽(实测), 超高速不太行, 考虑过段时间放暑假后用C++或者Rust重写这个项目.这只是一个实验性质的项目, 我估计Go这种带GC的语言本身就不适合拿来做这种特别底层的项目.

你们可以先拿来实验, 有了Rust版本, Go和Rust会同时维护, 不用担心这个项目不被维护了.

### 如何使用
1. 如果局域网内使用, 可以直接直连
2. 公网使用需要用Zerotier进行中转, 因为这个需要双端连接, 双端运行Agent, 如果用来进行内网穿透和科学上网等等, 必须使用N2N,Zerotier等P2P VPN进行转发.

### 注意事项
1. 这个项目使用明文传输, 如果需要传输敏感数据, 需要外面套一层VPN, Zerotier的情况下, 一般可以不套, 因为Zerotier是一个P2P的VPN

