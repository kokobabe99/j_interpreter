pkg main
imp "std/io"

def main() {
  var x i32 = 10;

  if x > 12 {
    io.Println("x is greater than 5");
  } else {
    io.Println("x is NOT greater than 5");
  }
}