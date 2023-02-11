## 主题目录结构
| 目录或文件          | 说明                                                  |
|----------------|-----------------------------------------------------|
| page目录         | 可以在page目录下创建自定义页面，可以直接url访问。支持模板语法，可以调用template目录文件 |
| public目录       | 可以在此目录下存放静态文件，可以直接url访问                             |
| template目录     | 目录下必须包含固定的页面模板                                      |
| theme.json     | 主题信息                                                |
| screenshot.png | 主题预览图                                               |

### template目录 模板文件
| 文件名           | 说明   |
|---------------|------|
| index.html    | 首页   |
| article.html  | 文章页  |
| category.html | 分类页  |
| tag.html      | 标签页  |
| notFound.html | 404页 |



### theme.json

| 字段          | 说明   |
|-------------|------|
| name        | 主题名称 |
| version     | 版本   |
| author      | 作者   |
| description | 描述   |
| homepage    | 主页   |
| license     | 协议   |

> 主题文件夹需放到themes目录下<br>
> 不明白的可以查看默认的主题的结构和写法<br>
> 如果默认主题被删除，重新运行程序即可生成默认主题