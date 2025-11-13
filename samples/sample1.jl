pkg main
imp "std/io"

def handler() {
  msg := recover();
  if msg != "" { io.Println("recovered:", msg) }
}

def main() {
  io.Println("ok");
  later handler();
  panic "boom";
}