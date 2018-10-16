var a = 1;
console.log("var: " + a);

if (false) {
  console.log("if true");
} else if (true) {
  console.log("if false");
}

var x = 3 % 2 === 0 ? 1 : 2;
var msg = x + "是" + (x % 2 === 0 ? "偶数" : "奇数");
switch (x) {
  case 1:
    console.log(msg);
    break;
  case 2:
    console.log(msg);
  default:
    console.log("others msg");
}

var i = 0;

while (i < 2) {
  console.log("i 当前为：" + i);
  i = i + 1;
}

for (var i = 0; i < 2; i++) {
  console.log("for i:" + i);
}

var obj = {
  foo: "foov",
  bar: "bar",
  getFoo: function() {
    return this.foo;
  }
};
for (var i in obj) {
  console.log("键名：", i);
  console.log("键值：", obj[i]);
}

do {
  console.log(i);
  i++;
} while (i < x);

function b64Encode(str) {
  return btoa(encodeURIComponent(str));
}

function b64Decode(str) {
  return decodeURIComponent(atob(str));
}

console.log(b64Encode("你好")); // "JUU0JUJEJUEwJUU1JUE1JUJE"
console.log(b64Decode("JUU0JUJEJUEwJUU1JUE1JUJE")); // "你好"

with (document.links[0]) {
  console.log(href);
  console.log(title);
  console.log(style);
}

function f(x) {
  if (x instanceof Function) {
    return x.name;
  }
  console.log(arguments);
  return "";
}
console.log(f(f));

function throwit() {
  throw new Error("");
}
function catchit() {
  try {
    throwit();
  } catch (e) {
    console.log(e.message); // print stack trace
  } finally {
    console.log("完成清理工作");
  }
}
catchit();

function UserError(message) {
  this.message = message || "默认信息";
  this.name = "UserError";
}

UserError.prototype = new Error();
UserError.prototype.constructor = UserError;

function Animal(name) {
  this.name = name;
}
Animal.prototype.color = "white";

var cat1 = new Animal("大毛");
var cat2 = new Animal("二毛");

cat1.color; // 'white'
cat2.color; // 'white'
cat1.color = "black";

cat1.color; // 'black'
cat2.color; // 'yellow'
Animal.prototype.color; // 'yellow';

function PubSub() {
  let yourMsg = {};
  yourMsg.peopleList = [];
  yourMsg.listen = function(fn) {
    this.peopleList.push(fn);
  };
  yourMsg.triger = function() {
    for (var i = 0, fn; (fn = this.peopleList[i++]); ) {
      fn.apply(this, arguments);
    }
  };

  yourMsg.listen(function(name) {
    console.log(`${name}收到了你的消息`);
  });
  yourMsg.listen(function(name) {
    console.log("哈哈");
  });

  yourMsg.triger("张三");
  yourMsg.triger("李四");
}
PubSub();