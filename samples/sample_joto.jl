pkg main
imp "std/io"

def main() {
  x := 0;

L1:
  io.Println("at L1, x =", x);
  x = x + 1;

  if x < 3 {
    joto L1;
  }

  io.Println("done");
}