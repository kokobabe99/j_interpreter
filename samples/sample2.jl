pkg main
imp "std/io"

def add(a: i32, b: i32): i32 {
  ret a + b * 2;
}

def main() {
  var r i32 = add(3, 4);
  io.Println("add", r);
}