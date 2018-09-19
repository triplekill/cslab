use super::*;
#[test]
fn it_works() {
    assert_eq!(2 + 2, 4);
    ::client::connect(); // 从根模块开始并列出整个路径
                         // 用 super 在层级中上移到当前模块的上一级模块
    super::client::connect();
}
#[test]
fn tsupper() {
    use super::client; //使其不再相对于根模块而是相对于父模块。
    client::connect();
}

#[test]
fn runall() {
    let mut config = Config {
        query: String::from("query"),
        filename: String::from("filename"),
    };
    run(config);
}
