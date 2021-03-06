# Golang-iDevice-CLI

[![license](https://img.shields.io/github/license/electricbubble/gidevice-cli)](https://github.com/electricbubble/gidevice-cli/blob/master/LICENSE)

## Installation

https://github.com/electricbubble/gidevice-cli/releases

#### Devices

```shell
$ gidevice list
$ gidevice listen
```

#### DeveloperDiskImage

```shell
$ gidevice mount -l
# gidevice mount -l -u=39xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx7
$ gidevice mount -d=/path/.../DeviceSupport/14.4/
$ gidevice mount /path/.../DeviceSupport/14.4/DeveloperDiskImage.dmg /path/.../DeviceSupport/14.4/DeveloperDiskImage.dmg.signature
```

#### App

```shell
$ gidevice applist
$ gidevice applist -t=all -u=39xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx7

$ gidevice launch com.apple.Preferences
$ gidevice kill com.apple.Preferences

$ gidevice install /path/.../WebDriverAgentRunner-Runner.ipa
$ gidevice uninstall com.leixipaopao.WebDriverAgentRunner.xctrunner

$ gidevice ps
```

#### Forward

```shell
# Default port local=8100 remote=8100
$ gidevice forward
$ gidevice forward -l=9100 -r=9100 -u=39xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx7
```

#### XCTest

```shell
$ gidevice xctest com.leixipaopao.WebDriverAgentRunner.xctrunner
# Only the logs contained in the text are displayed
$ gidevice xctest com.leixipaopao.WebDriverAgentRunner.xctrunner --contains="ServerURLHere->" --contains="Running tests..." --contains="Built at"
$ gidevice xctest com.leixipaopao.WebDriverAgentRunner.xctrunner --env=USE_PORT=8200 --env=MJPEG_SERVER_PORT=9200
```

#### Syslog

```shell
$ gidevice syslog
```

#### CrashReport

```shell
$ gidevice crashreport /path/.../local/dir/ -e -k
```

#### Screenshot

```shell
$ gidevice screenshot
$ gidevice screenshot -u=39xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx7 -o=/path/..../screenshot.png
```


#### DeviceInfo

```shell
$ gidevice info
$ gidevice info -u=39xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx7 --json=true   # for json format output
```


## Thanks

| |About|
|---|---|
|[libimobiledevice/libimobiledevice](https://github.com/libimobiledevice/libimobiledevice)|A cross-platform protocol library to communicate with iOS devices|
|[anonymous5l/iConsole](https://github.com/anonymous5l/iConsole)|iOS usbmuxd communication impl iTunes protocol|
|[alibaba/taobao-iphone-device](https://github.com/alibaba/taobao-iphone-device)|tidevice can be used to communicate with iPhone device|
|**[electricbubble/gidevice](https://github.com/electricbubble/gidevice)**|communicate with iOS devices implemented with Golang|

Thank you [JetBrains](https://www.jetbrains.com/?from=gwda) for providing free open source licenses
