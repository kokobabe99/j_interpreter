pkg main
imp "std/io"

def sum(n: i32): i32 {
  io.Println("enter sum, n =", n);

  if n <= 0 {
    io.Println("base case, n <= 0, ret 0");
    ret 0;
  }

  var prev i32 = sum(n - 1);
  ret n + prev;
}

def main() {
  var x i32 = 5;
  var r i32 = sum(x);
  io.Println("sum(", x, ") =", r);
}