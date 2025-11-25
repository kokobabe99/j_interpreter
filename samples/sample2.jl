pkg main
imp "std/io"

def add(a: i32, b: i32): i32 {
  ret a + b * 2;
}

def main() {
  var r i32 = add(3, 2);
  var d i32 = div(10, 2);
  io.Println("add", r);
  io.Println("div", d);
}

def div(a: i32, b: i32): i32 {
  ret a/b;
}