export const formatTime = (date: Date) => {
    const year = date.getFullYear()
    const month = date.getMonth() + 1
    const day = date.getDate()
    const hour = date.getHours()
    const minute = date.getMinutes()
    const second = date.getSeconds()

    return (
        [year, month, day].map(formatNumber).join('/') +
        ' ' +
        [hour, minute, second].map(formatNumber).join(':')
    )
}

const formatNumber = (n: number) => {
    const s = n.toString()
    return s[1] ? s : '0' + s
}


export function getSetting(): Promise<WechatMiniprogram.GetSettingSuccessCallbackResult> {
    return new Promise<WechatMiniprogram.GetSettingSuccessCallbackResult>((resolve, reject) => {
        wx.getSetting({
            success: resolve,
            fail: reject
        })
    })

}


export function getUserInfo(): Promise<WechatMiniprogram.GetUserInfoSuccessCallbackResult> {
    return new Promise<WechatMiniprogram.GetUserInfoSuccessCallbackResult>((resolve, reject) => {
        wx.getUserInfo({
            success: resolve,
            fail: reject
        })
    });
}