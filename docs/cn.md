# 奔跑吧(bēn pǎo ba) [![Build Status](https://travis-ci.org/Shuliyey/benpaoba.svg?branch=master)](https://travis-ci.org/Shuliyey/benpaoba)
## 简介
奔跑吧(bēn pǎo ba)是一个终端程序，用来执行有按先后顺序规定的终端脚本，是运维工程师的神器

## 用法
### 终端
```
benpaoba -p 3 -o minimum
```

### yaml/文件 格式
```yaml
- scripts/
  - task1.sh
  - task2.sh
- benpaoba.yaml
```
