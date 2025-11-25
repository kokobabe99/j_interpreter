pkg main
imp "std/io"

def fib(n: i32): i32 {
  io.Println("enter fib, n =", n);

  if n < 0 {
    io.Println("guard: n < 0, force 0");
    ret 0;
  };

  if n <= 1 {
    io.Println("base case, n =", n);
    ret n;
  } else {
    var a i32 = fib(n - 1);
    var b i32 = fib(n - 2);
    ret a + b;
  };
}
def main() {
  var x i32 = 5;
  var r i32 = fib(x);
  io.Println("fib(", x, ") =", r);
}