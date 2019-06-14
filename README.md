# go_nwlib

## Noah Wang(熊尼玛)的GO公共库

### cache

#### cache.go

| Name | Type | Desc |
| ----- | ----- | ----- |
| CacheMgr | struct | |

### conf

```yaml
cache
database
oss
setting
``` 

### database

#### database.go

| Name | Type | Desc |
| ----- | ----- | ----- |
| DBManager | struct | |

### error

| Name | Type | Desc |
| ----- | ----- | ----- |
| GeneralRaiseError | struct | |

### models

#### User

#### WechatMappUser

### token

| Name | Type | Desc |
| ----- | ----- | ----- |
| User | interface | |
| TokenMgr | struct | |
| Encode | func | |
| Decode | func | |
| GenerateJWT | func | |
| GetJWT | func | |
| GetClaims | func | |
| CheckJWT | func | |
| CleanJWT  | func | |
