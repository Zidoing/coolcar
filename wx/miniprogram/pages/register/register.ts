Page({
    data: {
        genderIndex: 0,
        licImgURL: "",
        genders: ["未知", "男", "女", "其他"],
        birthDate: "1990-01-01",
        licNo: "",
        licName: "",
        state: "UNSUBMITTED" as "UNSUBMITTED" | "PENDING" | "VERIFIED"
    },
    onUploadLic() {
        console.log("onUploadLic")
        wx.chooseImage({
            success: result => {
                console.log(result, "success")
                if (result.tempFilePaths.length > 0) {
                    this.setData({
                        licImgURL: result.tempFilePaths[0]
                    })
                    setTimeout(() => {
                        this.setData({
                            licNo: "2222",
                            licName: "小霸王",
                            genderIndex: 1,
                            birthDate: "1991-1-1"
                        })
                    }, 1000)
                }
            }
        })
    },
    onGenderChange(e: any) {
        this.setData({
            genderIndex: e.detail.value
        })
    },
    onBirthDateChange(e: any) {
        this.setData({
            birthDate: e.detail.value
        })
    },
    onSubmit() {
        this.setData({
            state: "PENDING"
        })

        setTimeout(() => {
            this.onLicVerified()
        }, 3000);
    },
    onReSubmit() {
        this.setData({
            state: "UNSUBMITTED",
            licImgURL: ""
        })
    },
    onLicVerified() {
        this.setData({
            state: "VERIFIED"
        })
        wx.redirectTo({
            url: "/pages/lock/lock"
        })
    }
})