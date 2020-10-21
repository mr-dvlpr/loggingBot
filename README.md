##Logging bot for telegram

You can use this API for send your running app logs to certain Telegram Bot

***
####usage

**First**

add bellow parameters to .env file before using API

|KEY|Description|Mandatory|
|---|-----------|---------|
|LOGGING_BOT_TOKEN|Telegram bot token|✅|
|LOGGING_ADMIN_CHAT|telegram id of the users that your bot must send logs to them(you ccan add multiple users id (separated by ,). users must start the bot before adding their ids in this field|✅|
|LOGGING_API_HOST|your API server host(like : http://localhost|✅|
|LOGGING_API_PORT|your API server port(default value is 8080)|❌|


**Second**

use the sample code of [example.go](https://github.com/mr-dvlpr/loggingBot/example.go) for sending error message from your app to the telegram bot

**Third**

import bellow line in your code

```logging github.com/mr-dvlpr/loggingBot```

and then call `logging.Start()`