# 奔跑吧(пробег)
[![Build Status](https://travis-ci.org/Shuliyey/benpaoba.svg?branch=master)](https://travis-ci.org/Shuliyey/benpaoba)
## обзор
benpaoba(пробег) - это программа на основе командной строки, которая запускает набор скриптов в указанном порядке,
это удобный инструмент для разработчиков Devops Engineers

## Применение
### командная строка
```
benpaoba -p 3 -o minimum
```

### yaml/хранилище состав
```yaml
- scripts/
  - task1.sh
  - task2.sh
- benpaoba.yaml
```
