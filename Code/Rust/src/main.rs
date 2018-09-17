fn main() {
    var();
    datatype();
    compondtype();
}

// var and const
fn var() {
    // let y = 10;
    let mut x = 5;
    println!("The value of x is: {}", x);
    x = 6;
    // y = 122; // error: y is immutable var
    println!("The value of x is: {}", x);
    const MAX_POINTS: u32 = 100_000;
    println!("Const:{}", MAX_POINTS);

    //shadow
    let x = x * 2;
    println!("The value of x is: {}", x);

    let spaces = "   ";
    println!("The string len: {}", spaces.len());

    // let mut space = "   ";
    // space = space.len();
}

fn datatype(){
    let guess: u32 = "42".parse().expect("Not a number!"); // u32 is type annotation
    println!("The data type parse {}", guess);
    let _y: f32 = 3.0; // f32 type annotation
    let _heart_eyed_cat = '😻';
}

fn compondtype(){
    let _tup: (i32, f64, u8) = (500, 6.4, 1);
    let (_x, _y, _z) = _tup; // 模式匹配（pattern matching）
    println!("tuple {}",_tup.1);
}