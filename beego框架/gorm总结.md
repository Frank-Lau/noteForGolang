# gorm总结

## 1.示例表

```go
		type User struct {
			Id int
			Name string `orm:"unique"`
			Passwd string `orm:"size(20)"`
			Articles []*Article `orm:"rel(m2m)"` //设置多对多关系
		}
		//文章结构体
		type Article struct {
			Id int `orm:"pk;auto"`
			ArtiName string `orm:"size(20)"`
			Atime time.Time `orm:"auto_now"`
			Acount int `orm:"default(0);null"`
			Acontent string `orm:"size(500)"`
			Aimg string  `orm:"size(100)"`
		
			ArticleType*ArticleType `orm:"rel(fk)"` //设置一对多关系
			Users []*User `orm:"reverse(many)"`  //设置多对多的反向关系
		}
		//类型表
		type ArticleType struct {
			Id int
			Tname string `orm:"size(20)"`
			Articles []*Article `orm:"reverse(many)"` //设置一对多的反向关系
		}
		
		
		func init(){
			//1.连接数据库
			orm.RegisterDataBase("default","mysql","root:123456@tcp(127.0.0.1:3306)/test?charset=utf8")
			//2.注册表
			orm.RegisterModel(new(User),new(Article),new(ArticleType))
			//3.生成表
			//1.数据库别名
			//2.是否强制更新
			//3.创建表过程是否可见
			orm.RunSyncdb("default",false,true)
		}
```



​		1.一对一:

​         		rel(one)

​        		 reverse(noe)



​		2.一对多:    本质是给多的那张表添加了外键

​        		rel(fk)

​        		reverse(many)



​		3.多对多    本质是创建了一张关系表

​         		rel(m2m)

​        		reverse(many)

## 2.限制条件

​		

|          限制条件           |                             作用                             |
| :-------------------------: | :----------------------------------------------------------: |
|             pk              |                       设置给字段为主键                       |
|            auto             |                自增,但是要求该字段必须为整型                 |
|         default(0)          |       设置该字段默认值,需要注意字段类型和默认类型一致        |
|          size(100)          |       设置该字段的长度为100字节,一般用来设置字符串类型       |
|            null             |            实体类默认不为空,设置该字段将允许为空             |
|           unique            |                      设置该字段全局唯一                      |
| digitsdigits(12);decimals(4 | 设置浮点数位数和精度。比如这个是说，浮点数总共12位，小数位为四位。 |
|          auto_now           |          针对时间类型字段，作用是保存数据的更新时间          |
|        auto_now_add         |          针对时间类型字段,作用是保存数据的添加时间           |

​		注意：**当模型定义里没有主键时，符合int, int32, int64, uint, uint32, uint64 类型且名称为 Id 的 Field 将被视为主键，能够自增.** "

​		Mysql中时间类型有date和datetime两种类型，但是我们go里面只有time.time一种类型，如果项目里面要求精确的话，就需要指定类型，指定类型用的是type(date)或者type(datetime)



## 3.单表操作

### 		单表插入

```go
					//插入数据库
					//获取orm对象
					o := orm.NewOrm()
					//获取要插入的对象
					var article models.Article
					//给对象赋值
					article.ArtiName = articleName
					article.Acontent = content
					//这一步需要注意，我们存储的图片是图片地址，没有 .
					article.Aimg = "/static/img/"+fileName+fileExt
					//插入
					o.Insert(&article)


```
## 4.多表操作

### 		1.一对多插入	

```go
					//处理数据
					//数据库的插入操作
					//获取orm对象
					o := orm.NewOrm()
					//获取插入对象
					var article models.Article
					//给插入对象赋值
					article.Title = artileName
					article.Content = content
					article.Image = "/static/image/"+fileName
				
					//根据类型名称获取类型对象
					var articleType models.ArticleType
					articleType.TypeName = typeName
					o.Read(&articleType,"TypeName")
					article.ArticleType = &articleType
				
					//插入
					_,err = o.Insert(&article)

			
```

### 		2.一对多查询

### 			高级查询

​					

| 函数名       | 作用                                   | 用法                                                         |
| :----------- | :------------------------------------- | :----------------------------------------------------------- |
| Limit()      | 获取部分数据                           | 有两个参数，第一个参数是指定获取几条数据，第二个参数指定从哪里获取qs.Limit(size,start)。返回值还是qs |
| OrderBy()    | 根据指定的字段排序                     | 只有一个参数，参数作用是指定按照哪个字段排序，返回值是qs     |
| Distinct()   | 去重                                   | 没有参数，返回值是qs                                         |
| Count()      | 查询符合条件的数据条目数               | 没有参数，返回值是查询到的条目数和错误信息                   |
| Filter()     | 过滤器，相当于SQL语句中的where         | 有两个参数，第一个参数是指定查询条件，第二个参数是值         |
| All()        | 把查询到的数据全部存储到指定的容器里面 | 只有一个参数，指定存储查询对象的存储容器                     |
| RelatedSel() | 多表查询的时候使用，指定关联的数据库表 | 参数长度不限，关联几个表，放几个参数                         |

​					

```go
					//分页查询
					pageSize := int64(2)
					pageCount := float64(count) / float64(pageSize)
					pageCount = math.Ceil(pageCount)
					start := pageSize * (int64(pageIndex)  -1 )
					qs.Limit(pageSize,start).RelatedSel("ArticleType").Filter("ArticleType__TypeName",typeName).All(&articles)					
```

​				**惰性查询:**	表中含有一对多关系,默认不查询 另外一张表, 解决办法:**添加RelatedSel()**

​            

### 		

### 		3.多对多插入

​					查看文章详情

​					1) 知道插入数据到哪里

​					2)知道插入数据到哪个字段

​					3)知道插入什么数据

​					4)插入操作

```go
					//处理数据
					//查询数据
					//获取orm对象
					o := orm.NewOrm()
					//获取查询对象
					var article models.Article
					//给查询条件赋值
					article.Id = articleId
				
					//获取article对象   知道向哪里插入数据
				
					//获取多对多操作对象   知道插入到对象的哪个字段
					m2m := o.QueryM2M(&article,"Users")
					//第三步,获取要插入的数据   知道插入什么数据
					var user models.User
					userName := this.GetSession("userName")
					user.UserName = userName.(string)
					o.Read(&user,"UserName")
				
					//插入多对多关系
					m2m.Add(user)
```

### 		4多对多查询

​					展示有哪些人看了文章

```go
					//第一种多对多查询,无法使用高级查询
					o.LoadRelated(&article,"Users")
				
					////第二种多对多关系查询   正向插入，反向查询
					////filter  过滤器  指定查询条件，进行过滤查找
					var users []models.User
				
					//select * from user                     where article.Id == articleId
					o.QueryTable("User").Filter("Articles__Article__Id",articleId).Distinct().All(&users)
```

## 5.注意问题

​		orm中一对多关系删除的时候,如果删除的是外键所对应的数据,做的是级联删除,删除类型时,该类型对应的文章也会被删除

​		

|   cascade   | 级联删除(默认)                  |
| :---------: | ------------------------------- |
|  set_null   | 设置为null,需要设置 null = true |
| set_default | 设置为默认值,需要设置default值  |
| do_nothing  | 什么也不做,忽略                 |

​		

​	

```go
			//文章结构体
			type Article struct {
				Id int `orm:"pk;auto"`
				ArtiName string `orm:"size(20)"`
				Atime time.Time `orm:"auto_now"`
				Acount int `orm:"default(0);null"`
				Acontent string `orm:"size(500)"`
				Aimg string  `orm:"size(100)"`
			
				ArticleType*ArticleType `orm:"rel(fk);null;on_delete(set_null)"`
				Users []*User `orm:"reverse(many)"`
			}
```

