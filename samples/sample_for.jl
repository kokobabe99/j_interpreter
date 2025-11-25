pkg main
imp "std/io"

def main() {
  fr i := 0; i < 5; i = i + 1 {
    io.Println("i =", i);
  }
}