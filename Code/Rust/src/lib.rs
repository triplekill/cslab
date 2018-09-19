
#[cfg(test)]
mod tests {

    #[test]
    fn exploration() {
        assert_eq!(2 + 2, 4);
    }

    #[test]
    #[should_panic(expected="Make this test fail")]
    fn another() {
        panic!("Make this test fail");
    }
    #[test]
    #[ignore]
    fn assertt() {
        assert!(true);
    }
    #[test]
    fn assert_eqq() {
        assert_eq!(1, 1);
        assert_ne!(1, 2);
    }
    #[test]
    fn main(){
    }
}