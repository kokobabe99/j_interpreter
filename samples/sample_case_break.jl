pkg main
imp "std/io"

def main() {
  x := 2;

  switch x {
    case 1:
      io.Println("one");
      break;

    case 2:
      io.Println("two");
      break;

    dft:
      io.Println("other");
  }
}