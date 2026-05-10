# twitter_message_send

一个用于跟踪指定 X / Twitter 账号并把新内容推送到飞书的小工具。

它适合用来做个人信息流监控：把少量高价值账号放进配置里，定时拉取新帖子，去重后通过飞书发送提醒。

## What it does

- 定时拉取配置中的 X / Twitter 用户最新内容
- 对已推送 tweet 做短期去重，避免重复提醒
- 支持把消息发送到飞书
- 通过配置文件管理账号、鉴权信息和接收人

## Configuration

复制配置模板：

```bash
cp conf/config_template.yaml conf/my_config.yaml
```

然后填写：

- X / Twitter 请求所需的 headers、cookie、token
- 需要关注的 `user-ids`
- 飞书开放平台应用信息
- 接收人的 `lark_id`

配置文件里会包含个人 token / cookie，请不要提交真实配置。

## Run

```bash
go mod tidy
go run main.go
```

默认代码会读取：

```text
conf/my_config.yaml
```

## Notes

这个项目是个人信息监控工具，不是通用的 Twitter API SDK。它更关注“少量账号、及时提醒、个人使用”的场景。
