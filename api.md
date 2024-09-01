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

> Body Parameters

```yaml
script: "import os\r

  import vapoursynth as vs\r

  from vapoursynth import core\r

  from vsrealesrgan import realesrgan, RealESRGANModel\r

  \r

  clip = core.bs.VideoSource(source=os.getenv('FINALRIP_SOURCE'))\r

  clip = core.resize.Bicubic(clip=clip, format=vs.RGBH)\r

  clip = realesrgan(clip=clip,
  model=RealESRGANModel.AnimeJaNai_HD_V3_Compact_2x)\r

  clip = core.resize.Bicubic(clip=clip, matrix_s=\"709\",
  format=vs.YUV420P16)\r

  clip.set_output()"
encode_param: ffmpeg -i - -vcodec libx265 -crf 16
video_key: Roshidere-08.mkv
```

### Params

| Name           | Location | Type   | Required | Description        |
| -------------- | -------- | ------ | -------- | ------------------ |
| body           | body     | object | no       | none               |
| » script       | body     | string | yes      | vapoursynth script |
| » encode_param | body     | string | yes      | encoder param      |
| » video_key    | body     | string | yes      | video oss key      |

> Response Examples

> 200 Response

```json
{
  "success": true,
  "error": {
    "message": "string"
  }
}
```

### Responses

| HTTP Status Code | Meaning                                                 | Description | Data schema |
| ---------------- | ------------------------------------------------------- | ----------- | ----------- |
| 200              | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | none        | Inline      |

### Responses Data Schema

HTTP Status Code **200**

| Name       | Type    | Required | Restrictions | Title | description |
| ---------- | ------- | -------- | ------------ | ----- | ----------- |
| » success  | boolean | true     | none         |       | none        |
| » error    | object  | false    | none         |       | none        |
| »» message | string  | true     | none         |       | none        |

## POST New

POST /api/v1/task/new

new a task after upload oss

> Body Parameters

```yaml
video_key: Roshidere-08.mkv
```

### Params

| Name        | Location | Type   | Required | Description   |
| ----------- | -------- | ------ | -------- | ------------- |
| body        | body     | object | no       | none          |
| » video_key | body     | string | yes      | video oss key |

> Response Examples

> 200 Response

```json
{
  "success": true,
  "error": {
    "message": "string"
  }
}
```

### Responses

| HTTP Status Code | Meaning                                                 | Description | Data schema |
| ---------------- | ------------------------------------------------------- | ----------- | ----------- |
| 200              | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | none        | Inline      |

### Responses Data Schema

HTTP Status Code **200**

| Name       | Type    | Required | Restrictions | Title | description |
| ---------- | ------- | -------- | ------------ | ----- | ----------- |
| » success  | boolean | true     | none         |       | none        |
| » error    | object  | false    | none         |       | none        |
| »» message | string  | true     | none         |       | none        |

## POST Clear

POST /api/v1/task/clear

new a task after upload oss

> Body Parameters

```yaml
video_key: Roshidere-06.mkv
```

### Params

| Name        | Location | Type   | Required | Description   |
| ----------- | -------- | ------ | -------- | ------------- |
| body        | body     | object | no       | none          |
| » video_key | body     | string | yes      | video oss key |

> Response Examples

> 200 Response

```json
{
  "success": true,
  "error": {
    "message": "string"
  }
}
```

### Responses

| HTTP Status Code | Meaning                                                 | Description | Data schema |
| ---------------- | ------------------------------------------------------- | ----------- | ----------- |
| 200              | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | none        | Inline      |

### Responses Data Schema

HTTP Status Code **200**

| Name       | Type    | Required | Restrictions | Title | description |
| ---------- | ------- | -------- | ------------ | ----- | ----------- |
| » success  | boolean | true     | none         |       | none        |
| » error    | object  | false    | none         |       | none        |
| »» message | string  | true     | none         |       | none        |

## GET Progress

GET /api/v1/task/progress

### Params

| Name      | Location | Type   | Required | Description |
| --------- | -------- | ------ | -------- | ----------- |
| video_key | query    | string | yes      | none        |

> Response Examples

> 200 Response

```json
{
  "success": true,
  "error": {
    "message": "string"
  },
  "data": {
    "progress": [true],
    "encode_key": "string",
    "encode_url": "string",
    "encode_param": "string",
    "script": "string"
  }
}
```

### Responses

| HTTP Status Code | Meaning                                                 | Description | Data schema |
| ---------------- | ------------------------------------------------------- | ----------- | ----------- |
| 200              | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | none        | Inline      |

### Responses Data Schema

HTTP Status Code **200**

| Name            | Type      | Required | Restrictions | Title | description       |
| --------------- | --------- | -------- | ------------ | ----- | ----------------- |
| » success       | boolean   | true     | none         |       | none              |
| » error         | object    | false    | none         |       | none              |
| »» message      | string    | true     | none         |       | none              |
| » data          | object    | false    | none         |       | none              |
| »» progress     | [boolean] | true     | none         |       | none              |
| »» encode_key   | string    | true     | none         |       | 压制结果，oss key |
| »» encode_url   | string    | true     | none         |       | 压制结果, url     |
| »» encode_param | string    | true     | none         |       | 压制参数          |
| »» script       | string    | true     | none         |       | 压制脚本          |

## GET OSSPresigned

GET /api/v1/task/oss/presigned

### Params

| Name      | Location | Type   | Required | Description |
| --------- | -------- | ------ | -------- | ----------- |
| video_key | query    | string | no       | none        |

> Response Examples

> 200 Response

```json
{
  "success": true,
  "error": {
    "message": "string"
  },
  "data": {
    "url": "string"
  }
}
```

### Responses

| HTTP Status Code | Meaning                                                 | Description | Data schema |
| ---------------- | ------------------------------------------------------- | ----------- | ----------- |
| 200              | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | none        | Inline      |

### Responses Data Schema

HTTP Status Code **200**

| Name       | Type    | Required | Restrictions | Title | description |
| ---------- | ------- | -------- | ------------ | ----- | ----------- |
| » success  | boolean | true     | none         |       | none        |
| » error    | object  | false    | none         |       | none        |
| »» message | string  | true     | none         |       | none        |
| » data     | object  | false    | none         |       | none        |
| »» url     | string  | true     | none         |       | upload url  |

## POST RetryEncode

POST /api/v1/task/retry/encode

> Body Parameters

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

### Params

| Name           | Location | Type    | Required | Description        |
| -------------- | -------- | ------- | -------- | ------------------ |
| body           | body     | object  | no       | none               |
| » script       | body     | string  | yes      | vapoursynth script |
| » encode_param | body     | string  | yes      | encoder param      |
| » video_key    | body     | string  | yes      | video oss key      |
| » index        | body     | integer | yes      | video clip index   |

> Response Examples

> 200 Response

```json
{
  "success": true,
  "error": {
    "message": "string"
  }
}
```

### Responses

| HTTP Status Code | Meaning                                                 | Description | Data schema |
| ---------------- | ------------------------------------------------------- | ----------- | ----------- |
| 200              | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | none        | Inline      |

### Responses Data Schema

HTTP Status Code **200**

| Name       | Type    | Required | Restrictions | Title | description |
| ---------- | ------- | -------- | ------------ | ----- | ----------- |
| » success  | boolean | true     | none         |       | none        |
| » error    | object  | false    | none         |       | none        |
| »» message | string  | true     | none         |       | none        |

## POST RetryMerge

POST /api/v1/task/retry/merge

> Body Parameters

```yaml
video_key: Roshidere-06.mkv
```

### Params

| Name        | Location | Type   | Required | Description   |
| ----------- | -------- | ------ | -------- | ------------- |
| body        | body     | object | no       | none          |
| » video_key | body     | string | yes      | video oss key |

> Response Examples

> 200 Response

```json
{
  "success": true,
  "error": {
    "message": "string"
  }
}
```

### Responses

| HTTP Status Code | Meaning                                                 | Description | Data schema |
| ---------------- | ------------------------------------------------------- | ----------- | ----------- |
| 200              | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | none        | Inline      |

### Responses Data Schema

HTTP Status Code **200**

| Name       | Type    | Required | Restrictions | Title | description |
| ---------- | ------- | -------- | ------------ | ----- | ----------- |
| » success  | boolean | true     | none         |       | none        |
| » error    | object  | false    | none         |       | none        |
| »» message | string  | true     | none         |       | none        |

# Data Schema
