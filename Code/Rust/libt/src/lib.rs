use std::error::Error;
pub mod client;
mod t;

pub mod networkt;

pub struct Config {
    pub query: String,
    pub filename: String,
}

impl Config {
    pub fn new(args: &[String]) -> Result<Config, &'static str> {
        if args.len() < 3 {
            return Err("not enough arguments");
        }

        let query = args[1].clone();
        let filename = args[2].clone();

        Ok(Config { query, filename })
    }
}


pub fn run(config: Config) -> Result<(), Box<Error>> {
    var();
    datatype();
    fun();
    control_flow();
    ownership();
    reference();
    slicet();
    structf();
    enumss();
    matchh();
    t::t();
    collection();
    errorr();
    generics();
    traitt();
    Ok(())
}

/// Adds one to the number given.
///
/// # Examples
///
/// ```
/// let five = 5;
///
/// assert_eq!(6, libt::add_one(5)); // Doc-tests
/// ```
pub fn add_one(x: i32) -> i32 {
    x + 1
}

// var and const
fn var() {
    println!("var:");
    // let y = 10;
    let mut x = 5;
    println!("\tThe value of x is: {}", x);
    x = 6;
    // y = 122; // error: y is immutable var
    println!("\tThe value of x is: {}", x);
    const MAX_POINTS: u32 = 100_000;
    println!("\tConst:{}", MAX_POINTS);

    //shadow
    let x = x * 2;
    println!("\tThe value of x is: {}", x);

    let spaces = "   ";
    println!("\tThe string len: {}", spaces.len());

    // let mut space = "   ";
    // space = space.len();
}

fn datatype() {
    println!("data type:");
    let guess: u32 = "42".parse().expect("Not a number!"); // u32 is type annotation
    println!("\tThe data type parse {}", guess);
    let _y: f32 = 3.0; // f32 type annotation
    let _heart_eyed_cat = '😻';

    let _tup: (i32, f64, u8) = (500, 6.4, 1);
    let (_x, _y, _z) = _tup; // 模式匹配（pattern matching）
    println!("\ttuple {}", _tup.1);

    let months = [
        "January",
        "February",
        "March",
        "April",
        "May",
        "June",
        "July",
        "August",
        "September",
        "October",
        "November",
        "December",
    ];
    println!("\tone:{}", months[0]);
}

fn fun() {
    simple_fun(10, 20);
}

fn simple_fun(x: i32, y: i32) {
    println!("define func:");
    println!("\tThe value of x is: {}", x);
    println!("\tThe value of y is: {}", y);
    let x = {
        let y = 1;
        y + 1
    };
    println!("\tThe result of the expression: {}", x);

    println!("\tHad return {}", return_fun());
}

fn return_fun() -> i32 {
    20
}

fn control_flow() {
    control_flow_if();
    control_flow_loop();
}

fn control_flow_if() {
    let number = 3;
    println!("Control flow if");
    if number < 5 {
        println!("\tcondition was true");
    } else if number < 3 {
        println!("\tcondition was false");
    } else {
        println!("\t???");
    }
    let number = if true { 5 } else { 6 };
    println!("\tuse if in let expression ret: {}", number)
}

fn control_flow_loop() {
    loop {
        println!("\tloop again!and break");
        break;
    }

    let mut number = 3;
    while number != 0 {
        println!("\t{}!", number);
        number = number - 1;
    }
    println!("\tLIFTOFF!!!");

    let a = [10, 20, 30, 40, 50];
    println!("\tfor iter array");
    for element in a.iter() {
        println!("\tfor iter the value is: {}", element);
    }
    println!("\tfor i in like:");
    for number in (1..4).rev() {
        println!("\t{}!", number);
    }
}

fn ownership() {
    let mut s = String::from("hello");
    s.push_str(", world!"); // push_str() 在字符串后追加字面值
    println!("Ownership:");
    println!("\tval:{}", s);

    // move
    let s1 = String::from("hello");
    let s2 = s1;
    // println!("{}, world!", s1); // E0382 s1 was moved :because `s1` has type `std::string::String`, which does not implement the `Copy` trait;

    // clone
    let s3 = s2.clone();
    println!("\ts12 = {}, s3 = {}", s2, s3);

    // ownership in func
    let s = String::from("hello");
    fn takes_ownership(some_string: String) {
        // some_string 进入作用域
        println!("\t{}", some_string);
    } // 这里，some_string 移出作用域并调用 `drop` 方法。占用的内存被释放
    takes_ownership(s.clone());
    takes_ownership(s);
    // println!("leave lifetime{}", s); // E0382 moved
    let s = 5;
    fn makes_copy(some_integer: i32) {
        // some_integer 进入作用域
        println!("\t{}", some_integer);
    } // 这里，some_integer 移出作用域。不会有特殊操作
    makes_copy(s);
    println!("\tleave lifetime {}", s);

    fn takes_and_gives_back(a_string: String) -> String {
        // a_string 进入作用域
        a_string // 返回 a_string 并移出给调用的函数
    }
    let _s3 = takes_and_gives_back(String::from("hello")); // s2 被移动到 _s3

    fn calculate_length(s: String) -> (String, usize) {
        let length = s.len(); // len() 返回字符串的长度
        (s, length)
    }
    let s1 = String::from("hello");
    let (s2, len) = calculate_length(s1);
    println!("\tThe length of '{}' is {}.", s2, len);
}

fn reference() {
    println!("Reference");
    fn calculate_length(s: &String) -> usize {
        s.len()
    }
    let s1 = String::from("hello");
    let len = calculate_length(&s1);
    println!("\tThe length of '{}' is {}.", s1, len);

    fn change(some_string: &mut String) {
        some_string.push_str(", world");
    }
    let mut s = String::from("hello");
    change(&mut s);
    println!("\tchange the word 'hello' to :{}", s);
}

fn slicet() {
    let a = [1, 2, 3, 4, 5];
    let slice = &a[1..3];
    println!("slice:\t[0]:{}", slice[0]);
}

fn structf() {
    println!("struct:");
    #[derive(Debug)]
    struct User {
        username: String,
        email: String,
        sign_in_count: u64,
        active: bool,
    }
    impl User {
        // associated functions
        fn new() -> User {
            User {
                email: String::from(""),
                username: String::from(""),
                active: false,
                sign_in_count: 0,
            }
        }
        //method
        fn get_username(&self) -> String {
            self.username.clone()
        }
        fn set_username(&mut self, username: String) {
            self.username = username;
        }
    }
    impl User {}

    let user1 = User {
        email: String::from("someone@example.com"),
        username: String::from("user1"),
        active: true,
        sign_in_count: 1,
    };
    let mut user2 = User {
        username: String::from("user2"),
        ..user1
    };
    println!("\tuser2: {:?}", user2);

    fn printstruct(user: &User) {
        println!("\tUser:{:#?}", user);
    }
    printstruct(&user2);
    println!("\t{}", user2.get_username());
    user2.set_username(String::from("user2changed"));
    println!("\t{}", user2.get_username());
    println!("\t{:?}", User::new());
}

fn enumss() {
    println!("enums:");
    #[derive(Debug)]
    enum IpAddrKind {
        V4,
        V6,
    }
    #[derive(Debug)]
    struct IpAddr {
        kind: IpAddrKind,
        address: String,
    }
    let home = IpAddr {
        kind: IpAddrKind::V4,
        address: String::from("127.0.0.1"),
    };
    let _loopback = IpAddr {
        kind: IpAddrKind::V6,
        address: String::from("::1"),
    };
    println!("\t{:?}", home);
    #[derive(Debug)]
    enum IpAddrv2 {
        V4(String),
        V6(String),
    }
    let home = IpAddrv2::V4(String::from("127.0.0.1"));
    println!("\t{:?}", home);
    #[derive(Debug)]
    enum Message {
        Quit,                       //Quit 没有关联任何数据。
        Move { x: i32, y: i32 },    //Move 包含一个匿名结构体
        Write(String),              //Write 包含单独一个 String。
        ChangeColor(i32, i32, i32), //ChangeColor 包含三个 i32
    }
    impl Message {
        fn call(&self) {
            println!("\t{:?}", self);
        }
    }
    let m = Message::Write(String::from("hello"));
    m.call();
}

fn matchh() {
    println!("match:");
    enum Coin {
        Penny,
        Nickel,
        Dime,
        Quarter,
    }

    fn value_in_cents(coin: Coin) -> u32 {
        match coin {
            Coin::Penny => {
                // 如果想要在分支中运行多行代码，可以使用大括号
                println!("\tmatch penny:");
                1
            }
            Coin::Nickel => 5,
            Coin::Dime => 10,
            Coin::Quarter => 25,
        }
    }
    println!("\tvalue: {}", value_in_cents(Coin::Penny));

    #[derive(Debug)] // So we can inspect the state in a minute
    enum UsState {
        Alabama,
        Alaska,
        // ... etc
    }
    enum Coin2 {
        Penny,
        Nickel,
        Dime,
        Quarter(UsState),
    }
    // 绑定值的模式
    fn value_in_cents2(coin: Coin2) -> u32 {
        match coin {
            Coin2::Penny => 1,
            Coin2::Nickel => 5,
            Coin2::Dime => 10,
            Coin2::Quarter(state) => {
                println!("\tState quarter from {:?}!", state);
                25
            }
        }
    }
    value_in_cents2(Coin2::Quarter(UsState::Alaska));

    // 匹配Option<T>
    fn plus_one(x: Option<i32>) -> Option<i32> {
        match x {
            None => None,
            Some(i) => Some(i + 1),
        }
    }
    let five = Some(5);
    let six = plus_one(five);
    let none = plus_one(None);

    let some_u8_value = 0u8;
    match some_u8_value {
        1 => println!("\tone"),
        3 => println!("\tthree"),
        5 => println!("\tfive"),
        7 => println!("\tseven"),
        _ => println!("\tover~"),
    }

    let some_u8_value = Some(0u8);
    match some_u8_value {
        Some(3) => println!("\tthree"),
        _ => (),
    }
    if let Some(3) = some_u8_value {
        println!("\tthree");
    } else {
        println!("\tif let else");
    }
}

fn collection() {
    println!("Collection:");
    let mut v: Vec<i32> = Vec::new();
    v.push(1);
    v.push(2);
    println!("\t{:?}", v);
    let mut v = vec![1, 2, 3];
    v.push(4);
    println!("\t{:?}", v);
    let third: &i32 = &v[2];
    println!("\t{:?}", third);
    let third: Option<&i32> = v.get(2);
    println!("\t{:?}", third);
    let mut v = vec![100, 32, 57];
    for i in &v {
        // 不使用&，则所有全就moved，之后就不可以使用了
        println!("\t{}", i);
    }
    for i in &mut v {
        *i += 50;
    }
    // save mutil type use enum
    #[derive(Debug)]
    enum SpreadsheetCell {
        Int(i32),
        Float(f64),
        Text(String),
    }

    let row = vec![
        SpreadsheetCell::Int(3),
        SpreadsheetCell::Text(String::from("blue")),
        SpreadsheetCell::Float(10.12),
    ];
    println!("\t{:?}", row);

    // string
    let hello = String::from("السلام عليكم");
    let hello = String::from("Dobrý den");
    let hello = String::from("Hello");
    let hello = String::from("שָׁלוֹם");
    let hello = String::from("नमस्ते");
    let hello = String::from("こんにちは");
    let hello = String::from("안녕하세요");

    let hello = String::from("Olá");
    let hello = String::from("Hola");
    let hello = String::from("Здравствуйте");

    for h in hello.chars() {
        println!("\t{}", h);
    }
    println!("\t{}", hello);
    let hello = String::from("你好");
    for h in hello.chars() {
        println!("\t{}", h);
    }

    // map
    use std::collections::HashMap;
    let mut scores = HashMap::new();
    scores.insert(String::from("Blue"), 10);
    scores.insert(String::from("Yellow"), 50);
    println!("\tmap: {:?}", scores);

    let teams = vec![String::from("Blue"), String::from("Yellow")];
    let initial_scores = vec![10, 50];
    let scores: HashMap<_, _> = teams.iter().zip(initial_scores.iter()).collect();
    println!("\tvector zip{:?}", scores);
    let team_name = String::from("Blue");
    let score = scores.get(&team_name);
    println!("\t{:?}", score);
    let score = match score {
        Some(p) => **p,
        _ => 0,
    };
    println!("\toption to i32:{}", score);
    for (key, value) in &scores {
        println!("\t{}: {}", key, value);
    }
    // scores.entry(String::from("Blue")).or_insert(50); 只有键没有对应的值时插入
}

fn errorr() {
    println!("Error");
    use std::fs::File;
    use std::io::ErrorKind;
    let f = File::open("hello.txt");

    let f = match f {
        Ok(file) => file,
        Err(ref error) if error.kind() == ErrorKind::NotFound => match File::create("hello.txt") {
            Ok(fc) => fc,
            Err(e) => panic!("Tried to create file but there was a problem: {:?}", e),
        },
        Err(error) => panic!("There was a problem opening the file: {:?}", error),
    };
    let f = File::open("hello.txt").unwrap();

    use std::io;
    use std::io::Read;

    fn read_username_from_file() -> Result<String, io::Error> {
        let f = File::open("hello.txt");
        let mut f = match f {
            Ok(file) => file,
            Err(e) => return Err(e),
        };
        let mut s = String::new();
        match f.read_to_string(&mut s) {
            Ok(_) => Ok(s),
            Err(e) => Err(e),
        }
    }
    fn read_username_from_file2() -> Result<String, io::Error> {
        let mut f = File::open("hello.txt")?;
        let mut s = String::new();
        f.read_to_string(&mut s)?;
        Ok(s)
    }
    let s = read_username_from_file();
    let s = match s {
        Ok(string) => string,
        Err(e) => String::new(),
    };
    println!("\t{:?}", s);
}

fn generics() {
    // TODO:
    fn largest<T: PartialOrd + Copy>(list: &[T]) -> T {
        let mut largest = list[0];

        for &item in list.iter() {
            if item > largest {
                largest = item;
            }
        }
        largest
    }

    #[derive(Debug)]
    struct Point<T> {
        x: T,
        y: T,
    }
    impl<T> Point<T> {
        fn x(&self) -> &T {
            &self.x
        }
    }
    //可以选择只为f32类型的point进行实现
    impl Point<f32> {
        fn distance_from_origin(&self) -> f32 {
            (self.x.powi(2) + self.y.powi(2)).sqrt()
        }
    }
    let integer = Point { x: 5, y: 10 };
    println!("\tstruct generics:{:?} x:{}", integer, integer.x());
    let flo = Point { x: 3.0, y: 4.0 };
    println!("\t{}", flo.distance_from_origin());
}

fn traitt() {
    pub trait Summarizable {
        fn author_summary(&self) -> String;

        fn summary(&self) -> String {
            format!("(Read more from {}...)", self.author_summary())
        }
    }

    pub struct Tweet {
        pub username: String,
        pub content: String,
        pub reply: bool,
        pub retweet: bool,
    }
    impl Summarizable for Tweet {
        fn author_summary(&self) -> String {
            format!("@{}", self.username)
        }
    }

    let tweet = Tweet {
        username: String::from("horse_ebooks"),
        content: String::from("of course, as you probably already know, people"),
        reply: false,
        retweet: false,
    };
    println!("\t1 new tweet: {}", tweet.summary());
    // >1 new tweet: (Read more from @horse_ebooks...)
}

#[cfg(test)]
mod tests;
