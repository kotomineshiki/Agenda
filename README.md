# Agenda

----------------------------

本项目是一个使用Cobra实现的命令行会议管理系统

## 项目需求

* 利用Go语言进行编程
* 使用面向对象的思想设计程序，使得程序具有良好的结构命令，并能方便修改、扩展新的命令不会影响其他命令的代码
* 项目部署在 Github 上，合适多人协作，特别是代码归并
* 支持日志（原则上不使用debug调试程序）

## 运行条件

1. 将项目文件，放入本地Go工作环境的src目录下
2. 在命令行中进入该文件中，运行`go build`命令
3. 之后就可以在命令行中利用`./Agenda [command_name] ......`来进行项目运行

## 相关命令

![命令显示](https://github.com/kotomineshiki/Agenda/blob/master/Picture/%E5%B1%8F%E5%B9%95%E5%BF%AB%E7%85%A7%202018-11-02%20%E4%B8%8A%E5%8D%8810.18.32.png?raw=true)

上图是本项目所包含的所有命令的简单介绍，可以用`./Agenda help`来查看，具体命令介绍如下：

### register

>* 注册新用户，要求用户提供一个唯一的用户名、密码、邮箱和电话。且对于个参数的输入格式都做了合法性检查。

* 使用格式：

```
./Agenda register -u username -p password -e email -t phone
```

* 可用参数列表

```
  -e, --email string       新用户邮箱
  -h, --help               help for register
  -p, --password string    新用户密码
  -t, --telephone string   新用户电话号码
  -u, --username string    新用户名
```

* 使用截图

![register](https://github.com/kotomineshiki/Agenda/blob/master/Picture/%E5%B1%8F%E5%B9%95%E5%BF%AB%E7%85%A7%202018-11-02%20%E4%B8%8A%E5%8D%8810.38.34.png?raw=true)

### login

>* 用户登录,用户输入用户名和密码登陆。且只有在没有登陆的状态下，用户名注册过且密码正确则登陆成功。

* 使用格式：

```
./Agenda login -u username -p password
```

* 可用参数列表

```
  -h, --help              help for login
  -p, --password string   用于登录的用户名
  -u, --username string   注册过的用户名
```
* 使用截图

![login](https://github.com/kotomineshiki/Agenda/blob/master/Picture/%E5%B1%8F%E5%B9%95%E5%BF%AB%E7%85%A7%202018-11-02%20%E4%B8%8A%E5%8D%8810.49.41.png?raw=true)

### logout

>* 用户登出,要求用户处于登入状态才能正确运行该命令

* 使用格式
 
```
./Agenda logout
```

* 可用参数列表

```
  -h, --help   help for logout
```

* 使用截图

![logout](https://github.com/kotomineshiki/Agenda/blob/master/Picture/%E5%B1%8F%E5%B9%95%E5%BF%AB%E7%85%A7%202018-11-02%20%E4%B8%8A%E5%8D%8810.50.24.png?raw=true)

### queryuser

>* 查询当前注册过的用户除密码以外的信息，要求当前处于登陆状态

* 使用格式

```
./Agenda queryuser             
```

* 可用参数列表
 
```
  -h, --help          help for queryuser
```

* 使用截图

![queryuser](https://github.com/kotomineshiki/Agenda/blob/master/Picture/%E5%B1%8F%E5%B9%95%E5%BF%AB%E7%85%A7%202018-11-02%20%E4%B8%8A%E5%8D%8810.51.32.png?raw=true)

### deleteuser

>* 删除用户，删除当前处于登陆状态的用户信息，该用户发起的所有回忆会被删除，该用户为参与者的会议会删除该参与者

* 使用格式

```
./Agenda deleteuser
```


* 可用参数列表

```
  -h, --help   help for deleteuser
```

* 使用截图

![deleteuser](https://github.com/kotomineshiki/Agenda/blob/master/Picture/%E5%B1%8F%E5%B9%95%E5%BF%AB%E7%85%A7%202018-11-02%20%E4%B8%8B%E5%8D%882.03.17.png?raw=true)


### createMeeting

>* 创建会议，已登录的用户创建一个新的会议。要求提供会议名称、会议的参与者、会议开始/结束时间。若会议名称重复，时间不合法，没有有效参与者则创建失败。

* 使用格式 

```
./Agenda createMeeting -t [Title] -p [\"name1, name2\"] -s [yyyy-mm-dd/hh:mm] -e [yyyy-mm-dd/hh:mm]
```

* 可用参数列表

```
  -e, --EndTime string         meeting's endTime
  -p, --Participator strings   meeting's participator
  -s, --StartTime string       meeting's startTime
  -t, --Title string           meeting title
  -h, --help                   help for createMeeting
```

* 使用截图

![createmeeting](https://github.com/kotomineshiki/Agenda/blob/master/Picture/%E5%B1%8F%E5%B9%95%E5%BF%AB%E7%85%A7%202018-11-02%20%E4%B8%8A%E5%8D%8811.11.54.png?raw=true)

### addparticipator

>* 增加参与者，已登录的用户可以增加自己发起会议的参与者。要求提供会议名称和添加的参与者

* 使用格式
 
```
./Agenda addparticipator -t [meetingtitle] -p ["name1, name2"])
```

* 可用参数列表

```
  -h, --help                   help for addparticipator
  -p, --participator strings   participator(s) you want to add, input like "name1, name2"
  -t, --title string           the title of meeting
```

* 使用截图

![addparticipator](https://github.com/kotomineshiki/Agenda/blob/master/Picture/%E5%B1%8F%E5%B9%95%E5%BF%AB%E7%85%A7%202018-11-02%20%E4%B8%8A%E5%8D%8811.13.05.png?raw=true)

### removeparticipator

>* 移除参与者，已登录的用户可以移除自己发起会议的参与者。要求提供会议名称和添加的参与者。

* 使用格式
 
```
./Agenda removeparticipator -t [title] -p [\"name1, name2\"]
```

* 可用参数列表

```
  -h, --help                   help for removeparticipator
  -p, --participator strings   the participator(s) of the meeting, input like "name1, name2"
  -t, --title string           the title of the meeting

```

* 使用截图

![removeparticipator](https://github.com/kotomineshiki/Agenda/blob/master/Picture/%E5%B1%8F%E5%B9%95%E5%BF%AB%E7%85%A7%202018-11-02%20%E4%B8%8B%E5%8D%881.57.05.png?raw=true)

### deleteMeeting 

>* 删除会议，已登录的用户可以删除自己作为发起者的某次会议，要求提供会议名称

* 使用格式

```
./Agenda deleteMeeting -t [title]
```

* 可用参数列表

```
  -t, --Title string   meeting title
  -h, --help           help for deleteMeeting

```

* 使用截图

![deleteMeeting](https://github.com/kotomineshiki/Agenda/blob/master/Picture/%E5%B1%8F%E5%B9%95%E5%BF%AB%E7%85%A7%202018-11-02%20%E4%B8%8B%E5%8D%882.01.35.png?raw=true)


### clearmeeting

>* 已登录的用户可以删除所有自己作为发起者的会议。

* 使用格式

```
./Agenda clearmeeting
```

* 可用参数列表
```
  -h, --help   help for clearmeeting
```

* 使用截图

![clearmeeting](https://github.com/kotomineshiki/Agenda/blob/master/Picture/%E5%B1%8F%E5%B9%95%E5%BF%AB%E7%85%A7%202018-11-02%20%E4%B8%8B%E5%8D%882.02.17.png?raw=true)


### quitmeeting

>* 已登陆用户可以推出自己作为参与者的会议
    
* 使用格式

```
  ./Agenda quitmeeting -t [title]
```

* 可用参数列表

```
  -h, --help           help for quitmeeting
  -t, --title string   the title of the meeting

```


### querymeeting

>* 查询会议，已登陆用户可以查询一段时间内自己作为参与者和发起者的所有会议信息
    
* 使用格式

```
  ./Agenda querymeeting -s [yyyy-mm-dd/hh:mm] -e [yyyy-mm-dd/hh:mm]
```

* 可用参数列表

```
  -e, --endtime string     endtime of time interval
  -h, --help               help for querymeeting
  -s, --starttime string   starttime of time interval
```

* 使用截图

![querymeeting](https://github.com/kotomineshiki/Agenda/blob/master/Picture/%E5%B1%8F%E5%B9%95%E5%BF%AB%E7%85%A7%202018-11-02%20%E4%B8%8B%E5%8D%881.52.37.png?raw=true)

