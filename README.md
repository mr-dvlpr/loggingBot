##Logging bot for telegram

You can use this API for send your running app logs to certain Telegram Bot

***
####usage

**First**

import bellow line in your code

```logging github.com/mr-dvlpr/loggingBot```

then

create new object of `MandatoryInfo` struct like bellow:

```
MandatoryInfo{
		BotToken:  "", 
		TgUserIds: "",
		ApiHost:   "",
		ApiPort:   0, // not mandatory field. If not set use 8080 for default port
	}
```

mandatory info decscribe as bellow:

|KEY|Description|Mandatory|
|---|-----------|---------|
|BotToken|Telegram bot token|✅|
|TgUserIds|telegram id of the users that your bot must send logs to them(you ccan add multiple users id (separated by ,). users must start the bot before adding their ids in this field|✅|
|ApiHost|your API server host(like : http://localhost)|✅|
|ApiPort|your API server port(default value is 8080)|❌|


**Second**

use the sample code of [example.go](https://github.com/mr-dvlpr/loggingBot/blob/master/example.go) for sending error message from your app to the telegram bot

**Third**

call `logging.Start()`

for sending message from API to your bot call `logging.SendError(message)`
