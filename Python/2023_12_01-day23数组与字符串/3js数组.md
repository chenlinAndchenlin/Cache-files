# 数组

## 一、数组的定义与概述

### 1、概述

数组的字面意思就是一组数据，一组(一般情况下相同类型)的数据（不一定都是数字）

数组的作用是：使用单独的变量名来存储一系列的值。

### 2、数组的定义

+ new Array（参数，参数,...）: 只有一个数字参数时是数组的长度（new可以省略,但一般尽量写上）

  ```javascript
  var arr = new Array();   //定义一个空数组 
  // 定义长度为10 的数组
  var arr = new Array(10);  //创建一个包含 10 个元素的数组，没有赋值
  var name=new Array(3)
  name[0]="迪丽热巴"
  name[1]="刘亦菲"
  name[2]="Lucky"
  ```

+ 创建带值的数组

  ```javascript
  var arr1 = new Array("芙蓉姐姐",30); //创建一个数组有两个元素
  ```

+ 字面量定义方式

  ```javascript
  var arr2 = [1,2,3,4,5];  //字面量定义方式
  ```

+ 快捷定义二维数组

  ```javascript
  var arr3 = [['a','b','c','d'],'b','c','d']; //快捷定义二维数组
  console.log(arr3[0][0])
  ```

  ​

  ​	

## 二、数组的访问

### 1、数组的下标

+ 下标就是索引，即元素的序号，从0开始，下标最大取值是：数组的长度(length)-1；

+ 下标可以是变量或表达式。

### 2、使用数组元素（访问） 

+ arr[0]:表示数组的第一个元素，0是下标，也叫索引

+ arr[1]:表示数组的第二个元素，1是下标

### 3、数组的长度（length属性）

+ 数组元素的个数    

  arr.length

+ length属性,不止是只读的，可以设置

  ```javascript
  var colors = new Array("red", "blue", "green");
  colors.length = 2;
  colors[2];
  ```

+ 注意

  + 如果设置的length的长度大于原有长度 则使用empty来占位
  + 如果设置的length的长度小于原有长度  则会将小于长度个数的元素删除掉




## 三、数组的使用与常用方法

### 1、数组的赋值

+ 概述

  给数组赋值，就是给数组的元素赋值，需要通过给数组的每个元素一一赋值，


+ 实例

  ```javascript
  arr[0] = 20;  //让数组的第一个元素的值为20；

  arr[1] = 12;  //让数组的第二个元素的值为12；
  ```

+ 以下为通过循环给数组的每个元素赋值，赋成下标的平方。

  ```javascript
   for(var i=0;i<10;i++){
     arr[i] = i*i
   }
  ```

+ 以下为通过循环给数组的每个元素赋值，随机数:Math.random()

  ```javascript
   for(var i=0;i<10;i++){
     arr[i] = Math.random();
   }
  ```


### 2、使用

+ 概述

  不能一次使用整个数组，使用数组就是在使用数组的每个元素，因为数组相当于若干个相同类型的变量。

+ 遍历数组：
  + 普通for循环

    ```javascript
    for(var i=0;i<5;i++){
       document.write(arr[i]);
    }
    ```

  + for...in语句用于遍历数组或者对象的属性（快速遍历）

    ```javascript
    for(var i in arr){
      	document.write(arr[i]);
    }
    ```

### 3、常用方法

+ **push()**         

  可以接收任意数量的参数，把它们逐个添加到数组的末尾，并返回修改后数组的长度。

  数组末尾添加一个元素，并且返回长度

  `alert(box.push('lucky')); `

+ **pop() **          

  从数组末尾移除最后一个元素，减少数组的 length 值，然后返回移除的元素。  

  移除数组末尾元素,并返回移除的元素

  `box.pop()`

+ **shift()**         

  从数组前端移除一个元素

  `alert(box.shift());	移除数组开头元素，并返回移除的元素`

+ **unshift()** 

   从数组前端添加一个或多个元素。

   `alert(box.unshift('广东','深圳'));		数组开头添加两个元素`

+ **reverse()**     

  逆向排序,原数组也被逆向排序了

  ```javascript
  var box = [1,2,3,4,5];
  alert(box.reverse()); //逆向排序方法，返回排序后的数组
  alert(box);    
  ```

+ **sort()**            

  从小到大排序 , 原数组也被升序排序了

  ```javascript
  var box = [4,1,7,3,9,2];
  alert(box.sort());   //从小到大排序，返回排序后的数组
  alert(box);  
  ```

+ **concat()** 

  追加数据, 创建一个新数组, 不改变原数组

  ```javascript
  var box = [2, 3, 4, '绿箭侠', '黑寡妇'];
  var box2 = box.concat('美队', '雷神');
  alert(box);
  alert(box2);  // 类型依然为数组
  var box2 = box+[5,6]  // 为字符串的连接 字符串类型
  ```

+ **slice()**

  不修改原数组, 将原数组中的指定区域数据提取出来

  ```javascript
  var box = [2, 3, 4, "绿巨人"];
  var box2 = box.slice(1, 3); //并没有修改原数组,将原数组中的元素提取出来,生成新数组,取的是下标在区域: [1,3)
  alert(box);
  alert(box2);
  ```

+ **box.join()**

   用数组的元素组成字符串

  ```javascript
  var box = [1, 2, 3, 4];
  console.log(box.join());    // 1,2,3,4 拼接成字符串用  默认用逗号作为分隔
  console.log(box.join('-')); // 1-2-3-4 拼接成字符串用-作为值得分隔
  ```

+ **Array.map()**  

   对于每个值进行处理并返回处理后的值作为新数组中的元素 

  ```javascript
   var numbers = [4, 9, 16, 25];
   new_numbers = numbers.map(function(v){
     return v + 10
   });
   // [14, 19, 26, 35]
  ```

+ 冒泡排序(了解)

  ```javascript
   //冒泡排序
   var arr = [1,4,2,7,3,5];
     for(var i=0;i<arr.length-1;i++){
       for(var j=0;j<arr.length-1-i;j++){
         if(arr[j]>arr[j+1]){
           var max = arr[j];
           arr[j] = arr[j+1];
           arr[j+1] = max;
         }
       }
     }
   console.log(arr);
  ```




