# go_nwlib

## Noah Wang(熊尼玛)的GO公共库

### cache

| Name | Type | Desc |
| ----- | ----- | ----- |
| CacheMgr | struct | |

### conf

| Name | Type | Desc |
| ----- | ----- | ----- |
| RedisConfig | struct | |
| DBConfig | struct | |
| OSSConfig | struct | |
| SettingsConfig | struct | |
| LoadConfig | func | |
| GetConfig | func | |
| CheckSettingConfig | func | |

### database

| Name | Type | Desc |
| ----- | ----- | ----- |
| DBManager | struct | |

### error

| Name | Type | Desc |
| ----- | ----- | ----- |
| GeneralRaiseError | func | |
| GeneralError | struct | |

### models

| Name | Type | Desc |
| ----- | ----- | ----- |
| User | struct | |
| WechatMappUser | struct | |

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
