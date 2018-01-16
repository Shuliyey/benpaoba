# 奔跑吧(run) [![Build Status](https://travis-ci.org/Shuliyey/benpaoba.svg?branch=master)](https://travis-ci.org/Shuliyey/benpaoba)
## Overview
benpaoba(run) is a cli based program that runs a collection of scripts in specified order, it's a handy tool for Devops Engineers

## Usage
### command line
```
benpaoba -p 3 -o minimum
```

### yaml/repo structure
```yaml
- scripts/
  - task1.sh
  - task2.sh
- benpaoba.yaml
```
