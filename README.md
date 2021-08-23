# JNU Health Checkin Action

<!-- ![GitHub repo size](https://img.shields.io/github/repo-size/azxj/jnu-stu-health-report)
![GitHub last commit](https://img.shields.io/github/last-commit/azxj/jnu-stu-health-report) -->
![GitHub](https://img.shields.io/github/license/Sakaizd/JNU_Health_Actions?style=for-the-badge)

## 目录
- [JNU Health Checkin Action](#jnu-health-checkin-action)
  - [目录](#目录)
  - [介绍 <a name="usages"></a>](#介绍-)
  - [使用方式](#使用方式)
  - [注意](#注意)
  - [停止/删除自动打卡方式](#停止删除自动打卡方式)
  - [其他说明](#其他说明)


## 介绍 <a name="usages"></a> 

某学校的基于 Github Action 完成自动化的每日健康打卡工作，时间设置为 01:00 UTC，即每天早上9点。推送打卡结果到微信（可选）

感谢 [原作者](https://github.com/azxj/jnu-stu-health-report) 提供的源代码。



## 使用方式
* 点击`Use this template`
* 在你的仓库点 `Settings` -> `Secrets` 中多次 `New repository secret` 添加以下配置：
  - `USERNAME`：账号
  - `PASSWORD`：密码
  - `SCKEY`：Server酱SCKEY，用于微信推送结果。(可选) [Server酱](https://sct.ftqq.com/)
  - （这里的“学号”和“密码”与在[暨南大学学生健康打卡](https://stuhealth.jnu.edu.cn)界面所输入的相同）：
* 通过 `Action` -> `Auto-checkin`  -> `Run workflow` 运行一次即可实现每天自动打卡，运行进度和结果可以在`Actions`页面查看
* 当输出“插入问卷数据成功”时，表示本次打卡成功；当输出“重复提交问卷”时，表示今日已经打过卡。
* 如果配置了 SCKEY，打卡结果会推送到微信

## 注意
* 由于 Github 的限制，仓库60天没有发生操作就会自动停止自动任务，建议60天内随便编辑下此文档续上。
* 如已经被停止，按使用方式所述重新启动


## 停止/删除自动打卡方式
* 进入`Actions`页，点击`All workflows`选项，切换成`Auto-checkin`，旁边`•••`菜单，选择`Disable workflow`停止自动打卡
* 进入`Settings`页，拉到底部，点击`Delete this repository`删除整个项目

## 其他说明
* 若当日已经打卡，程序会自动退出，不会重复打卡。
* 程序会自动读取上一次成功打卡的内容，作为这次打卡填写的内容。
* 帐号密码由 Github 加密存储，任务由 Github 自动执行，其中的所有隐私信息不会对外泄露。
