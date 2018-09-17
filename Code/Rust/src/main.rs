fn main() {
    var();
    datatype();
    fun();
    control_flow();
    ownership();
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
    println!("ownership:");
    println!("\tval:{}", s);
}
