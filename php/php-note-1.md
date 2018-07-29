
# Php知识点
## 标签:
``` php标签
<?php code ?> 
```
输出语句echo
变量定义前面要加$ （例如$a = 10）
连接变量{$xx} 也可以用点.
在html里显示
```php输出语句
echo<?php code ?>
```

## 单双引号的区别
* 双引号串中的内容可以被解释而且替换，而单引号串中的内容总被认为是普通字符。
* 例如：
> $foo = 2; 

> echo "foo is $foo"; // 打印结果: foo is 2 

> echo 'foo is $foo'; // 打印结果: foo is $foo

> echo "foo is $foo\n"; // 打印结果: foo is 2 (同时换行) 

> echo 'foo is $foo\n'; // 打印结果: foo is $foo\n)

## php的变量类型 有八个 三大类型

* 一 标准类型

> 整形 10

> 浮点型 10.3

> 布尔类型 true|false

> 字符串类型 $a = 'hello wrold'

* 二 复合类型 

> 数组 $arr = array(1,2)

> 对象 new object()

* 三 特殊类型

> null null

> 资源 wysql_connect('localhost','root','123') // 连接数据库的方法



## 数组和js数组的区别

* 数组和json 合在一起
* 索引数组
```
array('name','name1')
```

* 混合数组
```
array('name','name1','arg'=>'20')
```

```
array('name'=>'user1','age'=>'20','age'=>'boy')
```

``` 例如
echo ‘<h1>my name is {$arr['name']},my age is {$ar['age']....}’
```

## 对象定义

* 注意要先有类在new对象

* 对象连接（->）连接

* 对象定义和输出

```php对象的定义
class Person{
 public $name ='user'
	  public $age ='20'
	  public $sex ='boy'	
	  public function say () {
		echo 'my name is {$this-> name}'
	}
}
$a = new Person()
echo $a->age = 20
```

## 变量输出

* 1、Echo 例如 echo $a = ‘123’  => 123

* 2、Var_dump() Var_dump($a = ‘123’)  =>string(2) 123 打印出类型字节length

* 3、Print  Print $a = ‘123’  => 123

## Php变量类型测试

* 测试变量是否存在

> Isset() 变量不存在

> Empty() 是否为空 没定义变量也算 对象资源非空

## 变量类型测试

* 获取类型结果

> gettype 返回字符串 例如string array integer double（浮点型）object resource（资源）

> 精确判断类型 返回bool布尔

> Is_int() 判断整形

>Is_string()判断是不是字符串

> Is_array()判断是不是数组

> Is_object()判断是不是对象

> Is_null()判断是不是null

> Is_resource()判断是不是资源类型

> Is_scalar()判断是不是标准类型

> Is_numeric()判断是不是数字

> Is_callable()判断是不是函数

> Echo 布尔类型结果为1或空

> Var_dump 输出布尔

> (bool)$a强制类型转换布尔

> (string)$a强制类型转字符串

> ((int)$a强制类型转换整形

> (float)$a强制类型转换浮点型

> (array)$a强制类型转换数组

> (object)$a强制类型转换对象


## Php变量类型转换
* 自动转换
> 1、整形=>字符串 $a = 10 echo $a.’!!’

> 2、字符串=>整形 $a = ‘10abc’ echo $a + 10

> 3、所有类型 => 布尔值 $a = 1 是true null 0 false 空字符串等都是false

## 面试题

* 面试题-1

> $a = ‘1+2’ 

> echo $a + 10   =>  11
## 强制转换

## Php变量

## Php运算符
