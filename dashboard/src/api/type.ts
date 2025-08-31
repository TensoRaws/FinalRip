export interface Error {
  message: string
  [property: string]: any
}

export type TaskStatus = 'pending' | 'running' | 'completed'

export interface PingResponse {
  error?: Error
  success: boolean
  [property: string]: any
}

export interface NewTaskRequest {
  video_key: string
  [property: string]: any
}

export interface NewTaskResponse {
  error?: Error
  success: boolean
  [property: string]: any
}

export interface StartTaskRequest {
  encode_param: string
  script: string
  video_key: string
  slice?: boolean
  timeout?: number
  queue?: string
  [property: string]: any
}

export interface StartTaskResponse {
  error?: Error
  success: boolean
  [property: string]: any
}

export interface GetTaskProgressRequest {
  video_key: string
  [property: string]: any
}

export interface GetTaskProgressResponse {
  data?: {
    create_at: number
    encode_key: string
    encode_param: string
    encode_size: string
    encode_url: string
    key: string
    progress: {
      clip_key: string
      clip_url: string
      completed: boolean
      encode_key: string
      encode_url: string
      index: number
      [property: string]: any
    }[]
    script: string
    size: string
    status: TaskStatus
    url: string
    [property: string]: any
  }
  error?: Error
  success: boolean
  [property: string]: any
}

export interface OSSPresignedURLRequest {
  video_key: string
  [property: string]: any
}

export interface OSSPresignedURLResponse {
  data?: {
    exist: boolean
    url: string
    [property: string]: any
  }
  error?: Error
  success: boolean
  [property: string]: any
}

export interface ClearTaskRequest {
  video_key: string
  [property: string]: any
}

export interface ClearTaskResponse {
  error?: Error
  success: boolean
  [property: string]: any
}

export interface RetryEncodeTaskRequest {
  video_key: string
  index: number
  timeout?: number
  queue?: string
  [property: string]: any
}

export interface RetryEncodeTaskResponse {
  error?: Error
  success: boolean
  [property: string]: any
}

export interface RetryMergeTaskRequest {
  video_key: string
  [property: string]: any
}

export interface RetryMergeTaskResponse {
  error?: Error
  success: boolean
  [property: string]: any
}

export interface TaskListResquest {
  completed: boolean
  pending: boolean
  running: boolean
  [property: string]: any
}

export interface TaskListResponse {
  data?: {
    create_at: number
    encode_key: string
    encode_param: string
    encode_url: string
    key: string
    script: string
    status: TaskStatus
    url: string
    [property: string]: any
  }[]
  error?: Error
  success: boolean
  [property: string]: any
}
