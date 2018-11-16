# Cloudgo-io



---

设计一个 web 小应用，展示静态文件服务、js 请求支持、模板输出、表单处理、Filter 中间件设计等方面的能力。（不需要数据库支持）


## 文件结构

> |-- cloudgo-io
>     |-- main.go
>     |-- service
>     |   |-- service.go
>     |-- static
>     |   |-- css
>     |   |   |-- style.css
>     |   |-- images
>     |   |   |-- favicon.jpg
>     |   |   |-- login.jpg
>     |-- templates
>         |-- table.html
>         |-- login.html

## 使用框架

 - render 
 - Negroni, 
 - mux

## 测试结果
1、运行程序，打开http://localhost:8080/ 进入登录界面
![此处输入图片的描述][1]


  
2、输入用户名密码点击登录进入详情页面

![此处输入图片的描述][2]

3、访问http://localhost:8080/json

![此处输入图片的描述][3]


4、对 /unknown 给出开发中的提示

![此处输入图片的描述][4]

5、访问静态文件

![此处输入图片的描述][5]

6、服务器返回信息
![此处输入图片的描述][6]


  [1]: https://i.loli.net/2018/11/15/5bed83be5e16e.png
  [2]: https://i.loli.net/2018/11/15/5bed83f9ca1c3.png
  [3]: https://i.loli.net/2018/11/15/5bed8434dd266.png
  [4]: https://i.loli.net/2018/11/15/5bed8492cc422.png
  [5]: https://i.loli.net/2018/11/15/5bed84d67bff4.png
  [6]: https://i.loli.net/2018/11/15/5bed84fb893ca.png
