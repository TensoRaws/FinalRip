---
title: FinalRip
language_tabs:
  - shell: Shell
  - http: HTTP
  - javascript: JavaScript
  - ruby: Ruby
  - python: Python
  - php: PHP
  - java: Java
  - go: Go
toc_footers: []
includes: []
search: true
code_clipboard: true
highlight_theme: darkula
headingLevel: 2
generator: "@tarslib/widdershins v4.0.23"
---

# FinalRip

Base URLs:

# Authentication

- API Key (apikey-header-token)
  - Parameter Name: **token**, in: header.

# task

## POST Start

POST /api/v1/task/start

> Body 请求参数

```yaml
script: "import os\r

  import vapoursynth as vs\r

  from vapoursynth import core\r

  \r

  clip = core.bs.VideoSource(source=os.getenv('FINALRIP_SOURCE'))\r

  clip.set_output()"
encode_param: ffmpeg -i - -vcodec libx265 -crf 16
video_key: Roshidere-08.mkv
```

### 请求参数

| 名称           | 位置 | 类型   | 必选 | 说明               |
| -------------- | ---- | ------ | ---- | ------------------ |
| body           | body | object | 否   | none               |
| » script       | body | string | 是   | vapoursynth script |
| » encode_param | body | string | 是   | encoder param      |
| » video_key    | body | string | 是   | video oss key      |

> 返回示例

> 200 Response

```json
{
  "success": true,
  "error": {
    "message": "string"
  }
}
```

### 返回结果

| 状态码 | 状态码含义                                              | 说明 | 数据模型 |
| ------ | ------------------------------------------------------- | ---- | -------- |
| 200    | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | none | Inline   |

### 返回数据结构

状态码 **200**

| 名称       | 类型    | 必选  | 约束 | 中文名 | 说明 |
| ---------- | ------- | ----- | ---- | ------ | ---- |
| » success  | boolean | true  | none |        | none |
| » error    | object  | false | none |        | none |
| »» message | string  | true  | none |        | none |

## POST New

POST /api/v1/task/new

new a task after upload oss

> Body 请求参数

```yaml
video_key: Roshidere-08.mkv
```

### 请求参数

| 名称        | 位置 | 类型   | 必选 | 说明          |
| ----------- | ---- | ------ | ---- | ------------- |
| body        | body | object | 否   | none          |
| » video_key | body | string | 是   | video oss key |

> 返回示例

> 200 Response

```json
{
  "success": true,
  "error": {
    "message": "string"
  }
}
```

### 返回结果

| 状态码 | 状态码含义                                              | 说明 | 数据模型 |
| ------ | ------------------------------------------------------- | ---- | -------- |
| 200    | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | none | Inline   |

### 返回数据结构

状态码 **200**

| 名称       | 类型    | 必选  | 约束 | 中文名 | 说明 |
| ---------- | ------- | ----- | ---- | ------ | ---- |
| » success  | boolean | true  | none |        | none |
| » error    | object  | false | none |        | none |
| »» message | string  | true  | none |        | none |

## POST Clear

POST /api/v1/task/clear

new a task after upload oss

> Body 请求参数

```yaml
video_key: Roshidere-06.mkv
```

### 请求参数

| 名称        | 位置 | 类型   | 必选 | 说明          |
| ----------- | ---- | ------ | ---- | ------------- |
| body        | body | object | 否   | none          |
| » video_key | body | string | 是   | video oss key |

> 返回示例

> 200 Response

```json
{
  "success": true,
  "error": {
    "message": "string"
  }
}
```

### 返回结果

| 状态码 | 状态码含义                                              | 说明 | 数据模型 |
| ------ | ------------------------------------------------------- | ---- | -------- |
| 200    | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | none | Inline   |

### 返回数据结构

状态码 **200**

| 名称       | 类型    | 必选  | 约束 | 中文名 | 说明 |
| ---------- | ------- | ----- | ---- | ------ | ---- |
| » success  | boolean | true  | none |        | none |
| » error    | object  | false | none |        | none |
| »» message | string  | true  | none |        | none |

## GET Progress

GET /api/v1/task/progress

### 请求参数

| 名称      | 位置  | 类型   | 必选 | 说明 |
| --------- | ----- | ------ | ---- | ---- |
| video_key | query | string | 是   | none |

> 返回示例

> 200 Response

```json
{
  "success": true,
  "error": {
    "message": "string"
  },
  "data": {
    "progress": [
      {
        "completed": true,
        "index": 0,
        "clip_key": "string",
        "clip_url": "string",
        "encode_key": "string",
        "encode_url": "string"
      }
    ],
    "key": "string",
    "url": "string",
    "size": "string",
    "encode_key": "string",
    "encode_url": "string",
    "encode_size": "string",
    "encode_param": "string",
    "script": "string",
    "status": "string",
    "create_at": 0
  }
}
```

### 返回结果

| 状态码 | 状态码含义                                              | 说明 | 数据模型 |
| ------ | ------------------------------------------------------- | ---- | -------- |
| 200    | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | none | Inline   |

### 返回数据结构

状态码 **200**

| 名称            | 类型     | 必选  | 约束 | 中文名 | 说明                        |
| --------------- | -------- | ----- | ---- | ------ | --------------------------- |
| » success       | boolean  | true  | none |        | none                        |
| » error         | object   | false | none |        | none                        |
| »» message      | string   | true  | none |        | none                        |
| » data          | object   | false | none |        | none                        |
| »» progress     | [object] | true  | none |        | none                        |
| »»» completed   | boolean  | true  | none |        | none                        |
| »»» index       | number   | true  | none |        | none                        |
| »»» clip_key    | string   | true  | none |        | none                        |
| »»» clip_url    | string   | true  | none |        | none                        |
| »»» encode_key  | string   | true  | none |        | none                        |
| »»» encode_url  | string   | true  | none |        | none                        |
| »» key          | string   | true  | none |        | none                        |
| »» url          | string   | true  | none |        | none                        |
| »» size         | string   | true  | none |        | none                        |
| »» encode_key   | string   | true  | none |        | none                        |
| »» encode_url   | string   | true  | none |        | none                        |
| »» encode_size  | string   | true  | none |        | none                        |
| »» encode_param | string   | true  | none |        | none                        |
| »» script       | string   | true  | none |        | none                        |
| »» status       | string   | true  | none |        | pending, running, completed |
| »» create_at    | integer  | true  | none |        | unix time, int64            |

## GET OSSPresigned

GET /api/v1/task/oss/presigned

### 请求参数

| 名称      | 位置  | 类型   | 必选 | 说明 |
| --------- | ----- | ------ | ---- | ---- |
| video_key | query | string | 是   | none |

> 返回示例

> 200 Response

```json
{
  "success": true,
  "error": {
    "message": "string"
  },
  "data": {
    "url": "string",
    "exist": true
  }
}
```

### 返回结果

| 状态码 | 状态码含义                                              | 说明 | 数据模型 |
| ------ | ------------------------------------------------------- | ---- | -------- |
| 200    | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | none | Inline   |

### 返回数据结构

状态码 **200**

| 名称       | 类型    | 必选  | 约束 | 中文名 | 说明       |
| ---------- | ------- | ----- | ---- | ------ | ---------- |
| » success  | boolean | true  | none |        | none       |
| » error    | object  | false | none |        | none       |
| »» message | string  | true  | none |        | none       |
| » data     | object  | false | none |        | none       |
| »» url     | string  | true  | none |        | upload url |
| »» exist   | boolean | true  | none |        | none       |

## POST RetryEncode

POST /api/v1/task/retry/encode

> Body 请求参数

```yaml
script: |
  import os
  import vapoursynth as vs
  from vapoursynth import core

  clip = core.bs.VideoSource(source=os.getenv('FINALRIP_SOURCE'))
  clip.set_output()
encode_param: ffmpeg -i - -vcodec libx265 -crf 16
video_key: Roshidere-06.mkv
index: 2
```

### 请求参数

| 名称           | 位置 | 类型    | 必选 | 说明               |
| -------------- | ---- | ------- | ---- | ------------------ |
| body           | body | object  | 否   | none               |
| » script       | body | string  | 是   | vapoursynth script |
| » encode_param | body | string  | 是   | encoder param      |
| » video_key    | body | string  | 是   | video oss key      |
| » index        | body | integer | 是   | video clip index   |

> 返回示例

> 200 Response

```json
{
  "success": true,
  "error": {
    "message": "string"
  }
}
```

### 返回结果

| 状态码 | 状态码含义                                              | 说明 | 数据模型 |
| ------ | ------------------------------------------------------- | ---- | -------- |
| 200    | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | none | Inline   |

### 返回数据结构

状态码 **200**

| 名称       | 类型    | 必选  | 约束 | 中文名 | 说明 |
| ---------- | ------- | ----- | ---- | ------ | ---- |
| » success  | boolean | true  | none |        | none |
| » error    | object  | false | none |        | none |
| »» message | string  | true  | none |        | none |

## POST RetryMerge

POST /api/v1/task/retry/merge

> Body 请求参数

```yaml
video_key: Roshidere-06.mkv
```

### 请求参数

| 名称        | 位置 | 类型   | 必选 | 说明          |
| ----------- | ---- | ------ | ---- | ------------- |
| body        | body | object | 否   | none          |
| » video_key | body | string | 是   | video oss key |

> 返回示例

> 200 Response

```json
{
  "success": true,
  "error": {
    "message": "string"
  }
}
```

### 返回结果

| 状态码 | 状态码含义                                              | 说明 | 数据模型 |
| ------ | ------------------------------------------------------- | ---- | -------- |
| 200    | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | none | Inline   |

### 返回数据结构

状态码 **200**

| 名称       | 类型    | 必选  | 约束 | 中文名 | 说明 |
| ---------- | ------- | ----- | ---- | ------ | ---- |
| » success  | boolean | true  | none |        | none |
| » error    | object  | false | none |        | none |
| »» message | string  | true  | none |        | none |

## GET List

GET /api/v1/task/list

### 请求参数

| 名称      | 位置  | 类型    | 必选 | 说明 |
| --------- | ----- | ------- | ---- | ---- |
| pending   | query | boolean | 是   | none |
| running   | query | boolean | 是   | none |
| completed | query | boolean | 是   | none |

> 返回示例

> 200 Response

```json
{
  "success": true,
  "error": {
    "message": "string"
  },
  "data": [
    {
      "key": "string",
      "url": "string",
      "encode_key": "string",
      "encode_url": "string",
      "encode_param": "string",
      "script": "string",
      "status": "string",
      "create_at": 0
    }
  ]
}
```

### 返回结果

| 状态码 | 状态码含义                                              | 说明 | 数据模型 |
| ------ | ------------------------------------------------------- | ---- | -------- |
| 200    | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | none | Inline   |

### 返回数据结构

状态码 **200**

| 名称            | 类型     | 必选  | 约束 | 中文名 | 说明                        |
| --------------- | -------- | ----- | ---- | ------ | --------------------------- |
| » success       | boolean  | true  | none |        | none                        |
| » error         | object   | false | none |        | none                        |
| »» message      | string   | true  | none |        | none                        |
| » data          | [object] | false | none |        | none                        |
| »» key          | string   | true  | none |        | none                        |
| »» url          | string   | true  | none |        | none                        |
| »» encode_key   | string   | true  | none |        | none                        |
| »» encode_url   | string   | true  | none |        | none                        |
| »» encode_param | string   | true  | none |        | none                        |
| »» script       | string   | true  | none |        | none                        |
| »» status       | string   | true  | none |        | pending, running, completed |
| »» create_at    | integer  | true  | none |        | unix time, int64            |

# 数据模型
