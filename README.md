hkrpg-go

某动漫游戏的服务端再实现

随机首包密钥，随机生成seed，使数据更加安全

目前实现功能：

     1.走路模拟器
     2.队伍管理
     3.传送
     4.抽卡模拟器
     5.背包
     6.初级战斗
     7.养成
     8.初级模拟宇宙（1.6.0不支持）
     9.副本/周本/挑战
     10.忘却之庭(全关卡)

目前已实现功能中的bug:(因本人能力不足虽已尝试解决，但无法解决)

    1.角色无语音ui

目前实现的活动：

     1.角色试用
     2.巡星之礼
     3.忘却之庭-虚构叙事（不完善）

使用方法 ：

1. 安装mysql
2. 拉取项目文件
3. 编译
4. 运行后修改config.json
5. 运行

编译

   ```bash
   go mod tidy
   go build main.go
   ```
运行
```bash
  run main.exe / main
```

注：

* 如果你想帮助此项目，欢迎提交
