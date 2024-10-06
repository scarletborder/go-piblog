# 托管
博客托管和图床

## Example

```bash
go build
./host -f etc/host.yaml
```

## 功能

### 博客
- 上传结构化的Document
- 删除指定Document记录

### 图床
- 托管
  - 支持base64或url of picture
  - 根据图片的base64计算散列值作为key
- 删除指定图片