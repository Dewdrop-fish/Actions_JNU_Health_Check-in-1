# Health Check-in Actions for JNU

<!-- ![GitHub repo size](https://img.shields.io/github/repo-size/azxj/jnu-stu-health-report)
![GitHub last commit](https://img.shields.io/github/last-commit/azxj/jnu-stu-health-report) -->
![GitHub](https://img.shields.io/github/license/Sakaizd/JNU_Health_Actions)
![GitHub](https://github.com/Sakaizd/Actions_JNU_Health_Check-in/actions/workflows/auto_checkin.yml/badge.svg)

## Contents
- [JNU Health Checkin Action](#jnu-health-checkin-action)
  - [Contents](#contents)
  - [Introduction](#introduction)
  - [Usage](#usage)
  - [Tips](#tips)
  - [Stop/Delete](#stopdelete)
  - [Attention](#attention)
  - [Acknowledgments](#acknowledgments)


## Introduction 

<!-- 某JNU的基于Github Action完成自动化的每日健康打卡工作，时间设置为 01:00 UTC，即每天早上9点。推送打卡结果到微信（可选） -->

This is an automated daily health check-in based on Github Action for JNU, the chekc-in time was set to 01:00 UTC.

 The check-in result can be push to WeChat (optional) 




## Usage
<!-- * 点击`Use this template`
* 在你的仓库点击 `Settings` -> `Secrets` 中多次 `New repository secret` 添加以下配置：
  - `USERNAME`：账号
  - `PASSWORD`：密码
  - `SCKEY`：Server酱SCKEY，用于微信推送结果。(可选) [Server酱](https://sct.ftqq.com/)
  - （这里的“学号”和“密码”与在[暨南大学学生健康打卡](https://stuhealth.jnu.edu.cn)界面所输入的相同）：
* 通过 `Action` -> `Auto-checkin`  -> `Run workflow` 运行一次即可实现每天自动打卡，运行进度和结果可以在`Actions`页面查看
* 当输出“插入问卷数据成功”时，表示本次打卡成功；当输出“重复提交问卷”时，表示今日已经打过卡。
* 如果配置了SCKEY，打卡结果会推送到微信 -->

* Click the `Use this template` button to create a new repository.
* In your repository, click `Settings` -> `Secrets` multiple times in `New repository secret` to add the following configuration:
   - `USERNAME`: account
   - `PASSWORD`: Password
   - `SCKEY`: Server Chan SCKEY, used for WeChat push results. (Optional) [Server Chan](https://sct.ftqq.com/)
   - The "student ID" and "password" here should references to the websites [Jinan University Student Health Check-in](https://stuhealth.jnu.edu.cn)


## Tips
To prevent unnecessary workflow runs, scheduled workflows may be disabled automatically. When a public repository is forked, scheduled workflows are disabled by default. In a public repository, scheduled workflows are automatically disabled when no repository activity has occurred in 60 days. You may change the Readme.md file every 60 days to prevent this.

## Stop/Delete
* Go to `Actions` page，click `All workflows`->`Auto-checkin`->`•••`->`Disable workflow`
* Go to `Settings` page，click`Delete this repository` to delete this repo

## Attention
* If you have already finish the check-in, the actions results will inform you that the submittion was replicated.
* The program will automatically fill the contents read from the last successful check-in.
* The account and  password are encrypted and stored in Github. 

## Acknowledgments
 [Source code](https://github.com/azxj/jnu-stu-health-report) 
 
 ## License
[MIT](https://github.com/Sakaizd/Actions-LEDE/blob/main/LICENSE)
