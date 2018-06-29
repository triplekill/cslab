# loginweb.file
实现了简单的登陆功能。

用户信息存在config4login.properties。当用户登陆时候，会从配置文件中找到是否存在用户，验证登陆，想session中标记一个特征。注销时候消除session即可。

配置白名单。当所配置白名单没有在当前url时，拦截当前请求，返回消息。否则放行。

## 使用方法
写loginRest,配置filter，拦截一切没登陆请求。
