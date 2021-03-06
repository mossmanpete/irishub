# iriscli stake delegation

## 描述

基于委托者和验证者地址查询委托交易

## 用法

```
iriscli stake delegation [flags]
```
打印帮助信息
```
iriscli stake delegation --help
```
## 特有的flags

| 名称, 速记             | 默认值                      | 描述                                                                 | 必需     |
| --------------------- | -------------------------- | -------------------------------------------------------------------- | -------- |
| --address-delegator   |                            | [string] 委托者bech地址                                               | Yes      |
| --address-validator   |                            | [string] 验证者bech地址                                               | Yes      |

## 示例

### 查询验证者

```
iriscli stake delegation --address-validator=fva106nhdckyf996q69v3qdxwe6y7408pvyvfcwqmd --address-delegator=faa106nhdckyf996q69v3qdxwe6y7408pvyvufy0x2

```

运行成功以后，返回的结果如下：

```txt
Delegation
Delegator: faa13lcwnxpyn2ea3skzmek64vvnp97jsk8qmhl6vx
Validator: fva15grv3xg3ekxh9xrf79zd0w077krgv5xf6d6thd
Shares: 200.0000000
Height: 290
```
