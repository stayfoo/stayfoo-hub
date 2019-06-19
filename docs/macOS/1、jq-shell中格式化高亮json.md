### 1、 jq

- 官方：https://stedolan.github.io/jq/

> `shell` 中格式化高亮 `json` 。

- Mac 安装：
 
```
brew install jq
```

- 使用：

```
curl 'https://api.github.com/repos/stedolan/jq/commits?per_page=5' | jq '.'


curl --request POST --url http://10.4.9.101:9088/v1/app/info --data '{"packageId":"1284"}' |jq
```

