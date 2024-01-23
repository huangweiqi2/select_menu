
### 生成api

```
goctl api go -api api/main.api -dir .
```
### 生成文档

```api
goctl api plugin -plugin goctl-swagger="swagger -filename swagger.json" -api api/main.api -dir docs
```