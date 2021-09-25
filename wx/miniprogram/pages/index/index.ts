// index.ts
// 获取应用实例

Page({
    isPageShowing: false,
    data: {
        setting: {
            skew: 0,
            rotate: 0,
            showLocation: true,
            showScale: true,
            subKey: "",
            layerStyle: -1,
            enableZoom: true,
            enableScroll: true,
            enableRotate: false,
            showCompass: false,
            enable3D: false,
            enableOverLooking: false,
            enableSatellite: false,
            enableTraffic: false
        },
        location: {
            latitude: 31,
            longitude: 120
        },
        scale: 10,
        markers: [{
            iconPath: "/resources/car.png",
            id: 0,
            latitude: 23.099994,
            longitude: 113.324520,
            width: 50,
            height: 50
        }, {
            iconPath: "/resources/car.png",
            id: 1,
            latitude: 23.099994,
            longitude: 113.324520,
            width: 50,
            height: 50
        }]
    },
    onMyLocationTap() {
        wx.getLocation({
            type: "gcj02",
            success: result => {
                this.setData({
                    location: {
                        latitude: result.latitude,
                        longitude: result.longitude
                    }
                })
            }
        })
    },
    moveCars() {
        const map = wx.createMapContext("map");
        const dest = {
            latitude: 23.099994,
            longitude: 113.324520,
        }

        const moveCar = () => {
            dest.latitude += 0.1
            dest.longitude += 0.1
            map.translateMarker({
                destination: {
                    latitude: dest.latitude,
                    longitude: dest.longitude
                },
                markerId: 0,
                autoRotate: false,
                rotate: 0,
                duration: 5000,
                animationEnd: () => {
                    if (this.isPageShowing) {
                        moveCar()
                    }
                }
            })
        }
        moveCar()
    },

    onShow(): void | Promise<void> {
        this.isPageShowing = true
    },

    onHide(): void | Promise<void> {
        this.isPageShowing = false
    },

    onScanClick() {
        wx.scanCode({
            success: result => {
                console.log(result)
                wx.navigateTo({
                    url: "/pages/register/register"
                })
            }
        })
    }

})
