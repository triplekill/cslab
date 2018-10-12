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
