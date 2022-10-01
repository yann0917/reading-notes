# Rust 程序设计语言 

> [在线阅读地址-中文](https://rustwiki.org/zh-CN/book/title-page.html)
>
> [在线阅读地址-英文](https://doc.rust-lang.org/book/title-page.html)

[TOC]

## 入门指南

通过 rustup 安装 Rust 后，更新到最新版本 `rustup update`

要卸载 Rust 和 rustup `rustup self uninstall`

让浏览器打开本地文档 `rustup doc `

Rust 文件通常以 .rs 扩展名结尾。如果文件名中使用了多个单词，请使用**下划线**将它们隔开。

当看到一个 `!`，则意味着调用的是**宏**而不是普通的函数。

用分号（`;`，注意这是英文分号）结束该行，这表明该表达式已结束，下一个表达式已准备好开始。Rust 代码的大多数行都以一个 `;`结尾。

Cargo 是 Rust 的构建系统和包管理器.

* 使用 `cargo new xxx` 创建一个新项目
* 使用 `cargo build` 构建项目。
* 使用 `cargo run` 一步构建并运行项目。
* 使用 `cargo check` 构建项目而无需生成二进制文件来检查错误。
* 有别于将构建结果放在与源码相同的目录，Cargo 会将其放到 target/debug 目录。
* 使用 `cargo build --release` 来优化编译项目， 这会在 target/release 而不是 target/debug 下生成可执行文件。
* 使用 `cargo update` 升级 crate 


## 猜数字游戏

```rust
use rand::Rng;
use std::cmp::Ordering;
use std::io;

fn main() {
    println!("Guess the number!");

    let secret_number = rand::thread_rng().gen_range(1..101);

    loop {
        println!("Please input your guess.");

        let mut guess = String::new();

        io::stdin()
            .read_line(&mut guess)
            .expect("Failed to read line");

        let guess: u32 = match guess.trim().parse() {
            Ok(num) => num,
            Err(_) => continue,
        };

        println!("You guessed: {}", guess);

        match guess.cmp(&secret_number) {
            Ordering::Less => println!("Too small!"),
            Ordering::Greater => println!("Too big!"),
            Ordering::Equal => {
                println!("You win!");
                break;
            }
        }
    }
}


```
使用 `let` 声明来创建变量, 在变量名前使用 `mut` 来使一个变量可变。
```rust
let apples = 5; // 不可变
let mut bananas = 5; // 可变
```

`stdin()` 函数返回一个 `std::io::Stdin` 的实例，这代表终端标准输入句柄的类型。

就像变量一样，引用默认是不可变的。`&mut` 使其可变。

`{}` 是预留在特定位置的占位符。

Rust 允许用一个新值来**遮蔽** （shadow） 之前的值。

Rust 标准库中有很多叫做 `Result` 的类型：一个通用的 `Result` 以及在子模块中的特化版本，比如 `io::Result`。Result 类型是 枚举（enumerations），通常也写作 **enum**。枚举类型持有固定集合的值，这些值被称为枚举的**成员**（variant）。枚举往往与条件表达式 match 一起使用，可以方便地根据枚举值是哪个成员来执行不同的代码。

一个 `match` 表达式由分支（arm） 构成。一个分支包含一个用于匹配的模式（pattern），给到 match 的值与分支模式相匹配时，应该执行对应分支的代码。Rust 获取提供给 match 的值并逐个检查每个分支的模式。

字符串的  `parse()` 方法将字符串解析成数字。因为这个方法可以解析多种数字类型，因此需要告诉 Rust 具体的数字类型，这里通过 let guess: u32 指定。guess 后面的冒号（:）告诉 Rust 我们指定了变量的类型。Rust 有一些内建的数字类型；u32 是一个无符号的 32 位整型。
## 通用编程概念

### 变量和可变性

默认情况下**变量**是不可变的（immutable）, 当变量不可变时，这意味着一旦一个值绑定到一个变量名后，就不能更改该值了。

    1. 通过使用相同的变量名并重复使用 let 关键字来遮蔽（shadow）变量。
    2. 遮蔽和将变量标记为 mut 的方式不同，因为除非我们再次使用 let 关键字，否则若是我们不小心尝试重新赋值给这个变量，我们将得到一个编译错误。

**常量**（constant）是绑定到一个常量名且不允许更改的值，但是常量和变量之间存在一些差异。

    1. 常量不允许使用 mut。常量不仅仅默认不可变，而且自始至终不可变。常量使用 const 关键字而不是 let 关键字来声明，并且值的类型必须注明。
    2. 常量可以在任意作用域内声明，包括全局作用域。
    3. 常量只能设置为常量表达式，而不能是函数调用的结果或是只能在运行时计算得到的值。

Rust 常量的命名约定是**全部字母都使用大写，并使用下划线分隔单词**。

### 数据类型

Rust 的每个值都有确切的**数据类型**（data type），该类型告诉 Rust 数据是被指定成哪类数据，从而让 Rust 知道如何使用该数据。

标量（scalar）类型表示单个值。Rust 有 4 个基本的标量类型：**整型、浮点型、布尔型和字符。**

有符号整型以 i 开始，i 是英文单词 integer 的首字母，与之相反的是 u，代表无符号 unsigned 类型）。

每个有符号类型规定的数字范围是 `-(2n - 1) ~ 2n - 1 - 1`，其中 n 是该定义形式的位长度。所以 i8 可存储数字范围是 -(27) ~ 27 - 1，即 -128 ~ 127。无符号类型可以存储的数字范围是 `0 ~ 2n - 1`，所以 u8 能够存储的数字为 0 ~ 28 - 1，即 0 ~ 255。

可能属于多种数字类型的数字字面量允许使用类型后缀来指定类型，例如 `57u8`。数字字面量还可以使用 _ 作为可视分隔符以方便读数，如 1_000，此值和 1000 相同。

Rust 的浮点型是 `f32` 和 `f64`，它们的大小分别为 32 位和 64 位。默认浮点类型是 f64，因为在现代的 CPU 中它的速度与 f32 的几乎相同，但精度更高。所有浮点型都是有符号的。

Rust 中的布尔类型也有两个可能的值：`true` 和 `false`。布尔值的大小为 1 个字节。Rust 中的布尔类型使用 bool 声明。

Rust 的字符类型大小为 4 个字节，表示的是一个 Unicode 标量值。char 字面量采用单引号括起来。

复合类型（compound type）可以将多个值组合成一个类型。Rust 有两种基本的复合类型：**元组（tuple）和数组（array）**。

元组的长度是固定的：声明后，它们就无法增长或缩小。

我们通过在小括号内写入以逗号分隔的值列表来创建一个元组。元组中的每个位置都有一个类型，并且元组中不同值的类型不要求是相同的。

```rust
 let tup: (i32, f64, u8) = (500, 6.4, 1); // 声明后，它们就无法增长或缩小。
 let (x, y, z) = tup; // 解构元组
 println!("The value of y is: {}", y);
 let five_hundred = tup.0; // 或者使用 . 解构元组

```

没有任何值的元组 `()` 是一种特殊的类型，只有一个值，也写成 ()。该类型被称为**单元类型**（unit type），该值被称为单元值（unit value）。

Rust 中的数组具有固定长度，数组的每个元素必须具有相同的类型。

```rust
let a = [1, 2, 3, 4, 5]; // 方括号内以逗号分隔的列表形式将值写到数组中
let a: [i32; 5] = [1, 2, 3, 4, 5];

```
如果要为每个元素创建包含相同值的数组，可以指定初始值，后跟分号，然后在方括号中指定数组的长度。

```rust
let a = [3; 5]; //  这种写法与 let a = [3, 3, 3, 3, 3]; 效果相同，但更简洁。
```

当你尝试使用索引访问元素时，Rust 将检查你指定的索引是否小于数组长度。如果索引大于或等于数组长度，Rust 会出现 panic。

### 函数

```rust
fn main() {
    let x = plus_one(5);

    println!("The value of x is: {}", x);
    let y = {
        let x = 3;
        x + 1 // 行尾没有分号是表达式
    }; // 代码块 {} 是表达式

    println!("The value of y is: {}", y);
}

fn plus_one(x: i32) -> i32 { // -> 后声明它的类型 是返回值
    x + 1 // 行尾没有分号是表达式
}

```
Rust 代码中的函数和变量名使用下划线命名法（snake case，直译为蛇形命名法）规范风格。在下划线命名法中，所有字母都是小写并使用下划线分隔单词。

在函数签名中，必须声明每个参数的类型。

**语句**（statement）是执行一些操作但不返回值的指令。**表达式**（expression）计算并产生一个值。

函数调用是一个表达式。宏调用是一个表达式。我们用来创建新作用域的大括号（代码块） `{}` 也是一个表达式。

在 Rust 中，函数的返回值等同于函数体最后一个表达式的值。使用 return 关键字和指定值，可以从函数中提前返回；但大部分函数隐式返回最后一个表达式。

### 控制流

if 表达式允许根据条件执行不同的代码分支。Rust 并不会自动地将非布尔值转换为布尔值。你必须自始至终显式地使用布尔值作为 if 的条件。

Rust 有三种循环：loop、while 和 for。

如果存在嵌套循环，break 和 continue 应用于此时最内层的循环。你可以选择在一个循环上指定一个循环标签（loop label），然后将标签与 break 或 continue 一起使用，使这些关键字应用于已标记的循环而不是最内层的循环。

```rust
fn main() {
    let mut count = 0;
    'counting_up: loop {
        println!("count = {}", count);
        let mut remaining = 10;

        loop {
            println!("remaining = {}", remaining);
            if remaining == 9 {
                break;
            }
            if count == 2 {
                break 'counting_up;
            }
            remaining -= 1;
        }

        count += 1;
    }
    println!("End count = {}", count);
}

```

可以在用于停止循环的 break 表达式添加你想要返回的值。

```rust
fn main() {
    let mut counter = 0;
    let result = loop {
        counter += 1;

        if counter == 10 {
            break counter * 2;
        }
    };

    println!("The result is {}", result);
}
```
可以使用 for 循环来对一个集合的每个元素执行一些代码。
```rust
fn main() {
    let a = [10, 20, 30, 40, 50];

    for element in a {
        println!("the value is: {}", element);
    }

    // rev 用来反转区间（range）
    for number in (1..4).rev() {
        println!("{}!", number);
    }
    println!("LIFTOFF!!!");
}

```
## 认识所有权

## 使用结构体组织关联数据

## 枚举和模式匹配

## 使用包、Crate 和模块管理不断增长的项目

## 常见集合

## 错误处理

## 泛型、trait 与生命周期

## 编写自动化测试

## 一个 I/O 项目: 构建命令行程序

## 函数式语言功能：迭代器与闭包

## 更多关于 Cargo 和 Crates.io 的内容

## 智能指针

## 无畏并发

## 面对对象编程特性

## 模式和匹配

## 高级特征

## 最后的项目: 构建多线程 web 服务器

---
