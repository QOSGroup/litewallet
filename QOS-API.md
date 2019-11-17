# QOS Wallet SDK API

## 初始化钱包

* 方法:  InitWallet(name, storagePath string)
* 参数
    - name: 钱包名称 
    - storagePath: 存储路径
* 返回示例
    - 无

## 生成助记词

* 方法
    - ProduceMnemonic() string 
* 参数
    - 无
* 返回示例
```json
{
 "code":0,
 "data":"twist report state exchange army walnut thing that behave poem behind agree choice century side begin normal velvet flock just convince face staff mean"
}
```

## 创建账户(1)

使用默认名称创建账户, 默认名称格式为`Account${index}`

* 方法:  CreateAccount(password string) string
* 参数
    - password: 账户密码
* 返回示例

success:
    
```json

{
    "code": 0,
    "data": {
        "id": 1,
        "name": "Account3",
        "public_key": "qosaccpub1pmg7k5vqe8afm57grm5zf7m7fryz3ymh2a47rxlfandf0up62afsjm2d7c",
        "address": "qosacc1vrlzav25xmryvfchqtfuyguf3chemp94urzxf3",
        "private_key_enc": "/wuyL1bqmyvWDBl8UYUZ63oaMo0bP7EX1RgF1KbHK2qzVgvJRzfUaoX9DIt4v66zhjIIxAO61TEbqESo+/iT4KkD4QHuGvaLVhSbQUg+ztrycyjZRozD8A0F8MmQaBqRCXKwLqpCq/+JQ//2zoCSZhxC1h4gDZhNsUtNbs8808KBrV4A",
        "mnemonic": "hungry vault struggle width toward phrase flame artwork actor hover eyebrow kiss check smoke visa myself edit cabbage nose million light betray early injury"
    }
}

```

error: 

```json
{
    "code": 1,
    "message": "Key Account3 already exists"
}
```


## 创建账户(2)

* 方法:  CreateAccountWithName(name, password string) string 

使用指定名称创建账户

* 参数
    - name: 账户名称
    - password: 账户密码
* 返回示例
    - 同[创建账户(1)]()
    

## 创建账户(3)

* 方法:  CreateAccountWithMnemonic(name, password, mnemonic string) string

使用指定助记词创建账户, 若`name`未指定,则使用默认名称

* 参数
    - name: 账户名称
    - password: 账户密码
    - mnemonic: 助记词
* 返回示例
    - 同[创建账户(1)]()

## 获取账户(1)

根据账户地址获取账户信息

* 方法:  GetAccount(address string) string
* 参数
    - address: qos账户地址 
* 返回示例

success:

```json
{
    "code": 0,
    "data": {
        "id": 1,
        "name": "Account3",
        "public_key": "qosaccpub1wrr8lcy0f6q5efv9nl3e6fya6q7r2gkxpc5jz4fals5nfltfllvq4gkc5a",
        "address": "qosacc1zwcm3mxt6qj5zaus073823w57m5ug6qt4qsljl",
        "private_key_enc": "77m6AovYPAG8U54mOKEdxceiN5dwxXhZpAe2cGOhLfXsd5ztj9m3xPByZUpTLb88oxKwsvAscZw89kiyIY6KSoeHyv84NVQas4f0c7OdRLCgLKs1oYpCLI5s/rGaO+QXgjoxmyEOJIGaHCQR/xkRv1gQlHawW1z9vxe5nIQlJMWIrDu0"
    }
}
```

error:

```json
{
    "code": 1,
    "message": "Key Account5 not found"
}
```

```json
{
    "code": 1,
    "message": "Key qosacc1zwcm3mxt6qj5zaus073823w57m5ug6qt4qsljl not found"
}
```

## 获取账户(2)

根据账户名称获取账户信息

* 方法:  GetAccountByName(name string) string
* 参数
    - name: 账户名称
* 返回示例
    - 同[获取账户(1)]()

## 删除账户

* 方法:  DeleteAccount(address, password string) string
* 参数
    - address: qos账户地址 
    - password: 账户密码
* 返回示例

success:

```json
{"code":0}
```

error:

```json
{
    "code": 1,
    "message": "invalid account password"
}
```

## 导出账户

* 方法:  ExportAccount(address, password string) string
* 参数
    - address: qos账户地址 
    - password: 账户密码
* 返回示例

success

```json
{
    "code": 0,
    "data": {
        "id": 2,
        "name": "Account3",
        "public_key": "qosaccpub19m03rugvwfn4lrvmnnw4gsrdttza6g7dmxekx36y0rvq7cm6zfrq0wtnz6",
        "address": "qosacc1ge4c6x7rs7rmmehacdedyppknd3n3x6e8m0a8x",
        "private_key_enc": "K6hMQoDz9NN1RvpAjBm2Qao82VHucLrCAXgoUTa625tYI1VA8wSxBfHU1KlEEbkHNXfNpo/LwR5Bx/z+0+SUvpoBtFwtEDML8eQU6p+VTbOSCDBwTbMCzxTdRH5qx4eiSDX4Ng9xrJ0UMZMnPkM0jhF7/M6RSk8mzI5Pa67SZYkT7VCM",
        "private_key": "9069e0c30a9348b445ba69ef2ac737f94ac3619de7e0213a1f2e824281886d702edf11f10c72675f8d9b9cdd54406d5ac5dd23cdd9b363474478d80f637a1246"
    }
}
```

## 导入账户(1)

使用助记词导入新账户

* 方法:  ImportMnemonic(mnemonic, password string) string 
* 参数
    - mnemonic: 助记词 
    - password: 账户密码
* 返回示例
success:

```json
{
    "code": 0,
    "data": {
        "id": 11,
        "name": "Account11",
        "public_key": "qosaccpub1qqs9ffyl804cv0zt7h4y3xss0f9w3j4t5dyfkxp6jwpfl29tnveq078kxj",
        "address": "qosacc16y98jnnaq5vdw94fq3svuzl8j776ng6xa38d7n",
        "private_key_enc": "emRsmbPLJ4IFEhkol5zT/nkK1l5sJGlDkV5iZ9uxl+4ZID+ZDD497xToq+oaYobs6C8dVMOrPBLyxmix0J/AqcdZb4D8EwD6GCVhPod2M1T19E3kUejqNhjbR9iqkpoaxkDZnxBvK+/QR4UVbaxd/48oZEU4YuCIGxeFo5HqDlP7qXhd"
    }
}
```


## 导入账户(2)

使用私钥导入新账户, 私钥为[导出账户]()中的`private_key`

* 方法:  ImportPrivateKey(hexPrivateKey, password string) string
* 参数
    - hexPrivateKey: 16进制编码的私钥
    - password: 账户密码
* 返回示例
    - 同[导入账户(1)]()


## 获取账户列表

* 方法:  ListAllAccounts() string
* 参数
    - 无
* 返回示例
success:

```json
{
    "code": 0,
    "data": [
        {
            "id": 1,
            "name": "Account3",
            "public_key": "qosaccpub1wrr8lcy0f6q5efv9nl3e6fya6q7r2gkxpc5jz4fals5nfltfllvq4gkc5a",
            "address": "qosacc1zwcm3mxt6qj5zaus073823w57m5ug6qt4qsljl",
            "private_key_enc": "77m6AovYPAG8U54mOKEdxceiN5dwxXhZpAe2cGOhLfXsd5ztj9m3xPByZUpTLb88oxKwsvAscZw89kiyIY6KSoeHyv84NVQas4f0c7OdRLCgLKs1oYpCLI5s/rGaO+QXgjoxmyEOJIGaHCQR/xkRv1gQlHawW1z9vxe5nIQlJMWIrDu0"
        },
        {
            "id": 2,
            "name": "account-2",
            "public_key": "qosaccpub19m03rugvwfn4lrvmnnw4gsrdttza6g7dmxekx36y0rvq7cm6zfrq0wtnz6",
            "address": "qosacc1ge4c6x7rs7rmmehacdedyppknd3n3x6e8m0a8x",
            "private_key_enc": "K6hMQoDz9NN1RvpAjBm2Qao82VHucLrCAXgoUTa625tYI1VA8wSxBfHU1KlEEbkHNXfNpo/LwR5Bx/z+0+SUvpoBtFwtEDML8eQU6p+VTbOSCDBwTbMCzxTdRH5qx4eiSDX4Ng9xrJ0UMZMnPkM0jhF7/M6RSk8mzI5Pa67SZYkT7VCM"
        },
        {
            "id": 4,
            "name": "Account4",
            "public_key": "qosaccpub1vgd3rz4panu225l9kpwj5v3vl4rd5kvpam9rq4anfkml8e8dfr3skr58g2",
            "address": "qosacc184hy5purntjszxrgwy97gqfevmp7kfyju552ux",
            "private_key_enc": "XDlcxDedqjHeGsatVUQqBnWmWxVirTvxG7vv7taBF/bCOYBpmK/1UQ9cq061vadcx19JWWsLGl/xPHGc6RWbyvNELIzVTCAXzrj3xUA/gi/h3mhHG4eXC7+yWb4OKNSCBd7JrrwYqQdFNlcKKrsV3erOxMDzqi0HvmarphWvYPaHngBt"
        }
    ]
}
```

## 签名(1)

* 方法:  Sign(address, password, signStr string) string
* 参数
    - address: qos账户地址 
    - password: 账户密码
    - signStr: 待签名串
* 返回示例

返回的签名字符串以Base64编码表示

success: 
```json
{
    "code": 0,
    "data": "pII9YOD58mROHGL//lRh9p9lGhvtbzOFNMUmwirPu6mVY1IEBQGGzos1S7b/hzNY6/YX6ySYOVrAyA/euBwhAA==" //base64 encode
}
```

## 签名(2)

* 方法:  SignBase64(address, password, base64Str string) string
* 参数
    - address: qos账户地址 
    - password: 账户密码
    - base64Str: base64编码的待签名串
* 返回示例
    - 同[签名(1)]()


