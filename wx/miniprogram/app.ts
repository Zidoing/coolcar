// app.ts
import {getSetting, getUserInfo} from "./utils/util";


App<IAppOption>({
    globalData: {},
    onLaunch() {
        // 展示本地存储能力
        const logs = wx.getStorageSync('logs') || []
        logs.unshift(Date.now())
        wx.setStorageSync('logs', logs)

        // 登录
        wx.login({
            success: res => {
                console.log(res.code)
                // 发送 res.code 到后台换取 openId, sessionKey, unionId
            },
        })

        getSetting().then(res => {
            if (res.authSetting["scope.userInfo"]) {
                return getUserInfo()
            }
            return undefined
        }).then(res => {
            if (!res) {
                return
            }
            console.log(res.userInfo)
            this.globalData.userInfo = res.userInfo

            if (this.userInfoReadyCallback) {
                this.userInfoReadyCallback(res)
            }
        });
    },
})