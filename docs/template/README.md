
# 模板语法
* * *
#### if / else if / else
    {{ if foo == "asd" }}
        foo is 'asd'!
    {{ else if foo == 4711 }}
        foo is 4711!
    {{ else }}
        foo is something else!
    {{ end }}

#### range 循环
    {{ s := slice("foo", "bar", "asd") }}
    {{ range s }}
    {{.}}
    {{ end }}

#### include 引入文件
    {{ include "./header.html" }}

#### block 定义模块
    {{block copyright()}}
        <div>© ACME, Inc. 2020</div>
    {{end}}

#### yield 引入模块
    <footer>
        {{yield copyright()}}
    </footer>

#### extends 模板继承
    {{extends "./layout.html"}}
    {{block body()}}
        <main>
            This content can be yielded anywhere.
        </main>
    {{end}}

    <!-- file: "layout.html" -->
    <!DOCTYPE html>
    <html>
    <body>
        {{yield body()}}
    </body>
    </html>



> 更多模板语法，请参考golang jet: <br>
> https://github.com/CloudyKit/jet/blob/master/docs/syntax.md <br>
> https://github.com/CloudyKit/jet/blob/master/docs/builtins.md


# 数据模型

#### 文章

| 字段                       | 说明             |
|--------------------------|----------------|
| ID                       | ID             |
| Slug                     | 别名             |
| CategoryID               | 分类ID           |
| Title                    | 标题             |
| CreateTime               | 创建时间           |
| Thumbnail                | 缩略图            |
| Description              | 描述             |
| Keywords                 | 关键词  （查询列表时没有） |
| Content                  | 内容（查询列表时没有）    |
| Source                   | 来源（查询列表时没有）    |
| Extends                  | 扩展（查询列表时没有）    |
| URL()                    | url地址          |
| FullURL()                | 完整url地址        |
| CreateTimeFormat(layout) | 获取格式化的创建时间     |

#### 分类

| 字段                       | 说明            |
|--------------------------|---------------|
| ID                       | ID            |
| Slug                     | 别名            |
| Name                     | 名称            |
| ParentID                 | 父类ID          |
| CreateTime               | 创建时间          |
| Title                    | 标题            |
| Keywords                 | 关键词           |
| Description              | 描述            |
| URL()                    | url地址         |
| FullURL()                | 完整url地址       |
| PageURL(n)               | 分页url地址       |
| FullPageURL(n)           | 完整分页url地址     |
| CreateTimeFormat(layout) | 获取格式化的创建时间    |
| Children                 | 子分类           |

#### 标签
| 字段                       | 说明            |
|--------------------------|---------------|
| ID                       | ID            |
| Slug                     | 别名            |
| Name                     | 名称            |
| CreateTime               | 创建时间          |
| Title                    | 标题            |
| Keywords                 | 关键词           |
| Description              | 描述            |
| URL()                    | url地址         |
| FullURL()                | 完整url地址       |
| PageURL(n)               | 分页url地址       |
| FullPageURL(n)           | 完整分页url地址     |
| CreateTimeFormat(layout) | 获取格式化的创建时间    |

#### 链接
| 字段          | 说明     |
|-------------|--------|
| ID          | ID     |
| Name        | 名称     |
| URL         | 地址     |
| Logo        | logo   |
| CreateTime  | 创建时间   |
| ExpireTime  | 过期时间   |
| Note        | 备注     |
| Detect      | 是否自动检验 |
| DetectDelay | 自动检验延迟 |
| Status      | 链接状态   |




# 可调用标签
* * *

## 页面级标签

#### Page
    Page.Name // 页面名 article、category、tag、index、page
    Page.Title // 页面标题
    Page.Description // 页面描述
    Page.Keywords // 页面关键词
    Page.PageNumber // 当前页码
    Page.Path // 页面路径（template/page下文件专用，其他页面无效）

#### Data
    // 页面数据，不同的控制器对应不同数据
    // 文章页是文章数据，分类页是分类数据，标签页是标签数据...
    // 具体数据字段，请参考数据模型


## Config 配置
| 字段       | 说明   |
|----------|------|
| Site     | 站点   |
| Router   | 路由   |
| Upload   | 上传   |
| Cache    | 缓存   |
| Theme    | 主题   |
| Template | 模板   |
| Sitemap  | 站点地图 |

    // 举例
    Config.Site.Name // 站点名
    Config.Site.URL // 站点URL
    Config.Site.GetURL() // 格式化后的URL
    Config.Site.Title // 站点标题
    Config.Site.Keywords // 站点关键词
    Config.Site.Description // 站点描述
    Config.Upload.Domain // 上传域
    ... // 更多的可以打印出来自己看，基本和后台的配置对应，
    // 有一些模块的配置并没有开放出来,有的是觉得没必要，有的是为了安全考虑



## Widget 部件
> 部件可以通过后台配置（配置->模板）
#### head （用于html \<head>\</head>标签内）
    {{ Widget.Head() | raw }}
#### footer (用于body的页脚)
    {{ Widget.Footer() | raw }}
#### 轮播图
    {{ range i, v := Widget.Carousel()}}
        <div>
            <a href="{{v.Link}}">
                <img src="{{v.Image}}" /><p>{{v.Title}}</p>
            </a>
        </div>
    {{end}}
#### 分类导航
    {{ range i, v := Widget.Menu() }}
        <a href="{{v.URL()}}" target="_blank">{{v.Name}}</a>
    {{ end }}
#### 链接
    {{ range i, v := Widget.Link() }}
        <a href="{{v.URL}}" target="_blank">{{v.Name}}</a>
    {{ end }}

#### 首页文章列表
    {{ range i, v := Widget.IndexList() }}
        <a href="{{ v.URL() }}">{{v.Title}}</a>
        <img src="{{ v.Data.Thumbnail }}" />
        ......
    {{ end }}
#### 全局列表
    {{ range i, v := Widget.GlobalList() }}
        ......
    {{ end }}
#### 面包屑导航
    {{ range i, v := Widget.Breadcrumb(分类ID) }}
        <a href="{{ v.URL() }}">{{v.Name}}</a>
        ......
    {{ end }}
#### 标签云
    {{ range i, v := Widget.TagCloud() }}
        <a href="{{ v.URL() }}">{{v.Name}}</a>
        ......
    {{ end }}

#### 分类页列表（用在分类页模板中）
    // Data.ID 是当前分类ID
    // Page.PageNumber 是当前页码
    {{ res:= Widget.CategoryPageList(Data.ID, Page.PageNumber)}}

    // res 返回
    // res.List 文章列表
    // res.Count 文章统计
    // res.PageTotal 总页数
    // res.ExistNextPage 是否存在下一页
    // res.DisableCount 是否禁用count

    // 列表
    {{ range i, v := res.List }}
        ......
    {{ end }}

#### 标签页列表（用在标签页模板中）
    // Data.ID 是当前标签ID
    // Page.PageNumber 是当前页码
    {{ res:= Widget.TagPageList(Data.ID, Page.PageNumber)}}
    // res返回值同上


## 查询器
>可以复杂的查询一些数据<br>
>可用的链式命令：.Limit(10).Order("id desc").Comment("某某页查询文章")
### 文章

#### 根据ID获取文章
    Article().Get(1)
#### 获取文章列表
    Article().Limit(10).Order("id desc").List()
#### 文章伪随机列表
    Article().Limit(10).PseudorandomList()
#### 根据ID获取文章列表
    Article().ListByID(id1,id2,id3...)
#### 根据分类ID调用文章列表
    Article().ListByCategoryID(id1,id2,id3...)
#### 根据标签ID调用文章列表
    Article().ListByTagID(id1,id2,id3...)
#### 根据标签实体列表调用文章列表
    // 这个标签直接通过标签实体列表调用，可以用在文章页，直接通过文章的相关标签调取使用
    Article().ListByTags([]model.Tag)

### 分类
#### 根据ID获取分类
    Category().Get(1)
#### 获取分类列表
    Category().Limit(10).Order("id desc").List()
#### 分类伪随机列表
    Category().Limit(10).PseudorandomList()
#### 根据ID获取列表
    Category().ListByID(id1,id2,id3...)
#### 调用分类所有子分类列表
    Category().Children(parentID) // parentID为父分类ID
#### 根据ID获取分类以及其父分类
    Category().GetWithParent(id)

> 以下一些标签为了性能考虑只查询了一次分类列表，可以导致分类太多不显示的情况。<br>
> 可以调大更多配置里面的‘Categories limit’选项（默认调用前200个）

#### 调用分类所有后代列表
    Category().Descendants(rootID) // rootID为根分类ID
#### 获取分类以及其所有祖先（返回分类列表）
    Category().GetWithAncestors(id)
#### 调用分类列表以及其子分类
    Category().ListWithChildren(id1，id2, id3...)


### 标签
#### 根据ID获取标签
    Tag().Get(1)
#### 获取标签列表
    Tag().Limit(10).List()
#### 根据ID获取列表
    Tag().ListByID(id1,id2,id3...)
#### 标签伪随机列表
    Tag().Limit(10).PseudorandomList()
#### 通过文章id调用标签列表
    Tag().ListByArticleID(id1,id2,id3...) // id为文章ID

### 链接
#### 根据ID获取链接
    Link().Get(1)
#### 获取链接列表
    Link().List()
#### 根据ID获取列表
    Link().ListByID(id1,id2,id3...)
#### 获取公开的链接列表
    Link().ListPublic()

## Utils 工具集

#### 分页计算器
    res:= Utils.Pagination(page,pageTotal,limit) // 参数:当前页码,总页数,显示限制数量
    // 返回值：
    // res.Begin 开始页码
	// res.End  结束页码

#### 格式化时间戳
    Utils.FormatTimestamp(ts,layout) // 参数: 时间戳, 格式模板
    // 例子
    // 时间格式模板遵循golang格式: 2006代表年 01代表月 02代表日 15代表时 04代表分 05代表秒
    Utils.FormatTimestamp(1668052818, "2006-01-02 15:04:05")

#### 随机字符串
    Utils.RandString(n) // 参数n为生成的字符串长度
#### 随机数字
    Utils.RandInt(min,max) // 参数:[最小值,最大值)
#### 生成UUID
    Utils.UUID()

-----

>具体实践可以参考默认模板主题中的写法。