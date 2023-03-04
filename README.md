# Telegram Monitor

Telegram bot, which helps monitor websites.
It has got basic functions like add, delete, check websites. In future i am planning to add inline keyboard, automatic notification when the site come down and more features.
#

## Basics 
    /addsite - adding site to your list

    /deletesite - deleting site from yout list

    /status - sends status of your websites

    /help - shows available commands
#


## Techologies
![Go](https://img.shields.io/badge/-Go-090909?style=for-the-badge&logo=go)
![Postgresql](https://img.shields.io/badge/-Postgresql-090909?style=for-the-badge&logo=postgresql)
![Sql](https://img.shields.io/badge/-Sql-090909?style=for-the-badge&logo=mysql)

#
## Getting Started

```
$ https://github.com/fishkaoff/tg-monitor.git
```

```
$ go mod tidy
```

### Create file *preferences.env* in parent directory and paste: 

> TGTOKEN = [your telegram token from bot father] <br>
> DBURL = database url(postgresql)

### Create database and run this sql query:

> CREATE TABLE user_sites (
>   chatid INT NOT NULL UNIQUE,
>   site VARCHAR(255) NOT NULL,
>);

### Start the bot:
```
$ make start
```