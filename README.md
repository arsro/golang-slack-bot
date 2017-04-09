# golang_slack_bot

# How To Use

```golang
import golang_slack_bot

// エラー処理は適宜入れる
client := createSlackClient()
client.setPostParamter(bot_name, icon_emoji, icon_url, channel)
response, err := sp.sendMessage( "hello world", "{your incomming webhook url}" )
```
