### [gim](https://github.com/alberliu/gim)消息处理
> 单用户多设备支持，离线消息同步

每个用户都会维护一个自增的序列号，当用户A给用户B发送消息是，首先会获取A的最大序列号，设置为这条消息的seq，持久化到用户A的消息列表， 
再通过长连接下发到用户A账号登录的所有设备，再获取用户B的最大序列号，设置为这条消息的seq，持久化到用户B的消息列表，再通过长连接下发
到用户B账号登录的所有设备。假如用户的某个设备不在线，在设备长连接登录时，用本地收到消息的最大序列号，到服务器做消息同步，这样就可以
保证离线消息不丢失。

> 读扩散和写扩散

- 读扩散

简介：群组成员发送消息时，先建立一个会话，都将这个消息写入这个会话中，同步离线消息时，需要同步这个会话的未同步消息

优点：每个消息只需要写入数据库一次就行，减少数据库访问次数，节省数据库空间

缺点：一个用户有n个群组，客户端每次同步消息时，要上传n个序列号，服务器要对这n个群组分别做消息同步

- 写扩散

简介：在群组中，每个用户维持一个自己的消息列表，当群组中有人发送消息时，给群组的每个用户的消息列表插入一条消息即可

优点：每个用户只需要维护一个序列号和消息列表

缺点：一个群组有多少人，就要插入多少条消息，当群组成员很多时，DB的压力会增大

> 消息转发逻辑选型以及特点

- 普通群组：

采用写扩散，群组成员信息持久化到数据库保存。支持消息离线同步。

- 超大群组：

采用读扩散，群组成员信息保存到redis,不支持离线消息同步。

### 核心流程时序图

- 长连接登录

![eaf3a08af9c64bbd.png](http://www.wailian.work/images/2019/10/26/eaf3a08af9c64bbd.png)

- 离线消息同步

![ef9c9452e65be3ced63573164fec7ed5.png](http://s1.wailian.download/2019/12/25/ef9c9452e65be3ced63573164fec7ed5.png)

- 心跳

![6ea6acf2cd4b956e.png](http://www.wailian.work/images/2019/10/26/6ea6acf2cd4b956e.png)

- 消息单发

![e000fda2f18e86f3.png](http://www.wailian.work/images/2019/10/26/e000fda2f18e86f3.png)

- 小群消息群发

![749fc468746055a8ecf3fba913b66885.png](http://s1.wailian.download/2019/12/26/749fc468746055a8ecf3fba913b66885.png)

- 大群消息群发

![e3f92bdbb3eef199d185c28292307497.png](http://s1.wailian.download/2019/12/26/e3f92bdbb3eef199d185c28292307497.png)
