import { of } from "rxjs"
import { DeviceItem, ImageItem, LogItem, ScriptItem } from "src/types"

export const MatDialogRefStub = {
    close: () => { }
}

export const MAT_DIALOG_DATA_STUB = {
    callback: () => { }
}

export const ApiServiceResponsesStub = {
    getScriptsResponse: () => {
        const s1 = new ScriptItem()
        s1.ID = 5
        s1.Name = 'hello'
        s1.Path = 'testpath'
        const s2 = new ScriptItem()
        s2.ID = 57
        s2.Name = 'hellotwo'
        s2.Path = 'testpathtoo'
        return [s1, s2]
    },

    getDevicesResponse: () => {
        const m1 = new DeviceItem()
        m1.Mac = 'testmac'
        m1.ID = 3
        const m2 = new DeviceItem()
        m2.Mac = 'mactest'
        m2.ID = 42
        return [m1, m2]
    },

    addDeviceResponse: () => {
        const m1 = new DeviceItem()
        m1.Mac = 'testmac'
        m1.ID = 3
        return m1
    },

    getLogsResponse: () => {
        const log1 = new LogItem()
        log1.ID = 3
        log1.Summary = 'hello'
        const log2 = new LogItem()
        log2.ID = 35
        log2.Summary = 'hello 2'
        return [log1, log2]
    }
}

export const ApiServiceStub = {
    getImages: () => { },

    syncImages: () => { },

    addImage: (script: ImageItem) => { },

    uploadImage: (path: string, image: File) => { },

    editImage: (id: number, image: ImageItem) => { },

    deleteImage: (id: number) => { },

    getScripts: () => {
        return of(ApiServiceResponsesStub.getScriptsResponse())
    },

    syncScripts: () => { },

    addScript: (script: ScriptItem) => { },

    uploadScript: (path: string, script: File) => { },

    uploadScriptText: (path: string, script: string) => { },

    editScript: (id: number, script: ScriptItem) => { },

    deleteScript: (id: number) => { },

    getDevices: () => {
        return of(ApiServiceResponsesStub.getDevicesResponse())
    },

    addDevice: (device: DeviceItem) => of(device),

    editDevice: (id: number, device: DeviceItem) => of(device),

    deleteDevice: (id: number) => {
        const device = new DeviceItem()
        device.ID = id
        return of(device)
    },

    getFileContent: (path: string) => { },

    getLogs: () => {
        return of(ApiServiceResponsesStub.getLogsResponse())
    },
}

export const ConfirmServiceStub = {
    ask: (callback: Function, prompt?: string) => {
        callback()
    }
}
