export const FILE_TYPES = [ 'Image', 'Config', 'Script', 'Binary', 'Plain' ]

export class BaseModel {
  id: number | undefined
  createdAt: string | undefined
  updatedAt: string | undefined
  deletedAt: string | undefined
  type: string | undefined
}

export class FileItem extends BaseModel {
  name = ''
  path = ''
  fileType = ''
  lastAccessedAt = ''
}

export class DeviceItem extends BaseModel {
  name = ''
  mac = ''
  groupName = ''
  lastBootedAt = ''
  scriptID = ''
  script = new FileItem()
}

export class LogItem extends BaseModel {
  summary = ''
  detail = ''
}
