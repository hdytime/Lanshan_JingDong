# 蓝山寒假双人考核-goDong

> Go实现的婴儿级别京东网站后端

### 接口文档

------

[接口文档在这里](https://apifox.com/apidoc/shared-0ebfaf5e-4c45-4c2d-9c32-3951fc176f64)

### 前言💖

------

和前端同学一起合作完成了这个项目，在过程中发现自己很多问题，踩了一些坑，但同时也收获了不少。

### 技术栈💫

------

<img src="https://github.com/hdytime/Lanshan_JingDong/blob/master/images/color.png" style="zoom: 10%;" />

- [Gin](https://gin-gonic.com/zh-cn/)

> Gin 是一个用 Go (Golang) 编写的 HTTP Web 框架。 它具有类似 Martini 的 API，但性能比 Martini 快 40 倍。如果你需要极好的性能，使用 Gin 吧。

<img src="https://github.com/hdytime/Lanshan_JingDong/blob/master/images/R.png" style="zoom:7%;" />

- [MySQL](https://www.mysql.com/)

> 一个关系型数据库管理系统，由瑞典MySQL AB 公司开发，属于 Oracle 旗下产品。MySQL 是最流行的关系型数据库管理系统关系型数据库管理系统之一，在 WEB 应用方面，MySQL是最好的 RDBMS (Relational Database Management System，关系数据库管理系统) 应用软件之一

<img src="https://github.com/hdytime/Lanshan_JingDong/blob/master/images/redis-1536x864.jpg" style="zoom:20%;" />

- [Redis](https://redis.io/)

> Redis（Remote Dictionary Server）是一个使用[ANSI C](https://zh.wikipedia.org/wiki/ANSI_C)编写的[开源](https://zh.wikipedia.org/wiki/开源)、支持[网络](https://zh.wikipedia.org/wiki/电脑网络)、基于[内存](https://zh.wikipedia.org/wiki/内存)、[分布式](https://zh.wikipedia.org/wiki/分布式缓存)、可选[持久性](https://zh.wikipedia.org/wiki/持久性)的[键值对存储数据库](https://zh.wikipedia.org/wiki/键值-值数据库)。

<img src="https://github.com/hdytime/Lanshan_JingDong/blob/master/images/gorm.png" style="zoom:10%;" />

- [GORM](https://gorm.io/zh_CN/docs/index.html)

> The fantastic ORM library for Golang aims to be developer friendly.

<img src="https://github.com/hdytime/Lanshan_JingDong/blob/master/images/logo.png" style="zoom:25%;" />

- [Viper](https://link.zhihu.com/?target=https%3A//github.com/spf13/viper)

> Viper是适用于Go应用程序的完整配置解决方案。它被设计用于在应用程序中工作，并且可以处理所有类型的配置需求和格式。

<img src="https://github.com/hdytime/Lanshan_JingDong/blob/master/images/logo (1).png" style="zoom: 33%;" />

- [zap](https://github.com/uber-go/zap)

> zap是非常快的、结构化的，分日志级别的 Go 日志库。由Uber公司使用Go语言编写且开源。

### 功能💦

------

- [x] 用户注册
- [x] 用户登录
- [x] 用户个人主页
- [x] 用户信息修改
- [x] 商家发布商品信息
- [x] 用户查看商品
- [x] 用户将商品加入购物车
- [x] 用户购买商品
- [ ] [热榜](https://ranking.m.jd.com/rankingHome/rankingHome)
- [x] 用户密码加盐加密
- [ ] 用户登录有短信登录、邮箱登录、第三方登录多种形式
- [ ] 验证码（登录，注册，修改密码）
- [x] 用户状态保存使用 JWT 或 Session
- [x] 搜索功能（搜索商品）
- [x] 评论功能（售后评价）
- [ ] 用户浏览记录
- [x] 数据缓存（提高访问速度）
- [x] 考虑后端安全性（xxs，sql注入，cors，csrf 等）
- [ ] 秒杀活动
- [ ] 将项目部署上线（包括前端和后端的项目，也就是登录你的网站能够像正常的网站一样访问）
- [ ] 使用 https 加密

### 亮点😊

------

- **Redis中key-value设计**

![image-20240203122456867](https://github.com/hdytime/Lanshan_JingDong/blob/master/images/image-20240203122456867.png)

运用**用户名-商品id**作为**key**，**不同**的**value**来表示用户评论状态，0表示用户已购买商品未评论该商品，1表示购买也评论过该商品。

![image-20240203123203618](https://github.com/hdytime/Lanshan_JingDong/blob/master/images/image-20240203123203618.png)

- **商品序列号使用雪花算法生成**

> SnowFlake 中文意思为雪花，故称为雪花算法。最早是 Twitter 公司在其内部用于分布式环境下生成唯一 ID。在2014年开源 scala 语言版本。

雪花算法的优点是：

1. **有业务含义，并且可自定义。** 雪花算法的 ID 每一位都有特殊的含义，我们从 ID 的不同位数就可以推断出对应的含义。此外，我们还可根据自身需要，自行增删每个部分的位数，从而实现自定义的雪花算法。
2. **ID 单调增加，有利于提高写入性能。** 雪花算法的 ID 最后部分是递增的序列号，因此其生成的 ID 是递增的，将其作为[数据库](https://www.zhihu.com/search?q=数据库&search_source=Entity&hybrid_search_source=Entity&hybrid_search_extra={"sourceType"%3A"answer"%2C"sourceId"%3A2649891559})主键 ID 时可以实现顺序写入，从而提高写入性能。
3. **不依赖第三方系统。** 雪花算法的生成方式，不依赖第三方系统或[中间件](https://www.zhihu.com/search?q=中间件&search_source=Entity&hybrid_search_source=Entity&hybrid_search_extra={"sourceType"%3A"answer"%2C"sourceId"%3A2649891559})，因此其稳定性较高。
4. **解决了安全问题。** 雪花算法生成的 ID 是单调递增的，但其递增步长又不是确定的，因此无法从 ID 的差值推断出生成的数量，从而可以保护业务隐私。

![image-20240203123523170](https://github.com/hdytime/Lanshan_JingDong/blob/master/images/image-20240203123523170.png)

- **用户密码加盐加密**

> 使用"GoLang golang.org/x/crypto/bcrypt" 模块，通过该模块可以快速实现密码的存储处理。

BCrypt加密： 一种加盐的单向Hash，不可逆的加密算法，同一种明文（plaintext），每次加密后的密文都不一样，而且不可反向破解生成明文，破解难度很大。

![image-20240203165255430](https://github.com/hdytime/Lanshan_JingDong/blob/master/images/image-20240203165255430.png)

加盐加密的几个好处：

1. **密码唯一性**：通过使用随机生成的盐值，在不同的加密过程中产生的密码哈希值也会不同。这意味着即使两个用户的初始密码相同，最终加密后的密码哈希也会不同。这样可以确保每个用户的密码哈希值都是独一无二的，增加了密码的唯一性。
2. **防止彩虹表攻击**：彩虹表是一种预先计算并存储密码哈希值的数据表，攻击者可以使用它来快速破解大量哈希密码。通过引入盐值，即使密码相同，但每个用户使用的盐值都不同，导致生成的哈希密码也不同。这使得彩虹表攻击变得非常困难，因为攻击者无法预先计算和匹配所有可能的盐值。
3. **增加哈希强度**：盐值增加了哈希函数的强度，使得破解密码变得更为困难。通过在密码和盐的组合上执行哈希函数，攻击者需要显著增加破解密码的工作量。即使使用常见的密码也无法通过简单的散列算法直接破解。
4. **抵御碰撞攻击**：碰撞攻击是指通过找到两个不同的明文密码，然后经过哈希算法后得到相同的哈希值。使用盐值可以显著减少碰撞攻击的风险。即使明文密码相同，使用不同的盐值也会生成不同的哈希值，增加了攻击者找到碰撞的难度。

- **JWT认证**

> JWT（JSON Web Token）是一种用于认证和授权的开放标准。它是一种基于JSON的安全传输机制，通过在网络中传递签名的令牌来验证和授权用户访问资源。

<img src="https://github.com/hdytime/Lanshan_JingDong/blob/master/images/jwt.png" style="zoom:25%;" />

JWT的优点包括：

1. **基于标准化**：JWT是一个开放的标准，广泛被支持和使用，具有广泛的跨平台兼容性。
2. **去中心化**：JWT是无状态的，不需要在服务器端存储会话信息。服务器可以基于每个请求中的JWT完成认证和授权。
3. **安全性**：JWT使用数字签名验证令牌的完整性，确保令牌在传输过程中不被篡改或伪造。
4. **扩展性**：JWT的声明部分可以自由定制，可以包含任意的用户信息和其他相关信息。
