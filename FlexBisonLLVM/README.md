# 简介
## 综述
1. 文法分析用Flex:将数据分隔成一个个的标记token (标示符identifiers，关键字keywords，数字numbers, 中括号brackets, 大括号braces, 等等etc.)
2. 语法分析用Bison: 在分析标记的时候生成抽象语法树. Bison 将会做掉几乎所有的这些工作, 我们定义好我们的抽象语法树就OK了.
3. 组装用LLVM: 这里我们将遍历我们的抽象语法树，并未每一个节点生成字节/机器码。 这听起来似乎很疯狂，但是这几乎就是最简单的 一步了.

## 安装
```sh
cd /path/to/ComputSciLab
vagrant up
vagrant ssh
sudo apt-get insta ll flex bison clang
验证:i
    cd test
    lex valid.l
    gcc lex.yy.c -lfl

cd /vagrant/FlexBisonLLVM/my_toy_compiler
make test
```

# 资料
* https://github.com/lsegal/my_toy_compiler
* https://coolshell.cn/articles/1547.html
* https://www.zybuluo.com/clarazhang/note/190166
* https://github.com/shdxiang/xy-compiler
* http://lesliezhu.com/2014/05/29/%E6%9E%84%E5%BB%BAToy%E7%BC%96%E8%AF%91%E5%99%A8%E2%80%94%E2%80%94%E5%9F%BA%E4%BA%8EFlex-Bison%E5%92%8CLLVM/
