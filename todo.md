## 已了解
- 使用大小开头表示public函数
- Init()函数是一个特殊的函数，用于初始化程序或包
- - 调试使用GOLAND
- 使用h2来MOCK DB不好，不如使用本地的db,设置

```mac 
mysql.server start
mysql -u root -p new_password


```
设置环境变量
```
user := os.Getenv("MYSQL_USERNAME")
	pwd := os.Getenv("MYSQL_PASSWORD")
	addr := os.Getenv("MYSQL_ADDRESS")
	dataBase := os.Getenv("MYSQL_DATABASE")
```
如果使用了类似于db.Get()这样的函数来获取数据库连接，那么可能会出现nil的情况。这是因为在测试中，你需要手动初始化数据库连接，而不是依赖于应用程序中的初始化函数。


## ToDo
- debug的时候出现了unexpected fault address 0x1a78000，