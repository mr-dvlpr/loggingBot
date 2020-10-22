# Logging bot for telegram

You can use this API for send your running app logs to certain Telegram Bot

## usage

import bellow line in your code

```logging github.com/mr-dvlpr/loggingBot```

then

create a new object of `LoggingInfo` struct:

```
LoggingInfo{
    BotToken:  "", 
    TgUserIds: "",
    ApiHost:   "",
    AppName:   "",
    ApiPort:   0, // not mandatory field. If not set use 8080 for default port
}
```


call `logging.Start()`

for sending message from API to your bot call `logging.SendError(message)`

---
logging info describe bellow:

|KEY|Description|Mandatory|
|---|-----------|---------|
|BotToken|Telegram bot token|✅|
|TgUserIds|telegram id of the users that your bot must send logs to them(you ccan add multiple users id (separated by ,). users must start the bot before adding their ids in this field|✅|
|ApiHost|your API server host(like : http://localhost)|✅|
|AppName|The name of your app(showing in the top of bot messages)|❌|
|ApiPort|your API server port(default value is 8080)|❌|