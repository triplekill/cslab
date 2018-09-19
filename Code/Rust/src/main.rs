extern crate libt;
use libt::Config;
fn main() {
    let mut config = Config{query:String::from("query"),filename:String::from("filename")};
    if let Err(e) = libt::run(config) {
    }
}
