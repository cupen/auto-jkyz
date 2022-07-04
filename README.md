# ...


# Usage

## Windows

* 首先，安装个 chrome 浏览器
* 创建个快捷方式，在目标栏里追加启动参数 `--remote-debugging-port=9222  --user-data-dir=d:/remote-profile2` 
* 然后双击这个快捷方式打开 chrome 
* 随便找个工作目录，比如 `d://workbench/jkyz/`
* 在该目录下新建个配置文件 `conf.toml`
    ```toml
    [account]
    idtype="4" # 大陆护照用 4 
    username="hello"
    password="world"
    ```
* 下载执行文件 https://github.com/cupen/auto-jkyz/releases/download/v0.1.0rc1/auto-jkyz.exe

* 打开控制台（在该目录下文件管理器的地址栏输入 `cmd`即可）
* 输入 `auto-jkyz.exe`
* 然后你就可以看到，chrome 自动打开`健康驿站`。
