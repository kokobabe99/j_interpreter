pkg main
imp "std/io"

def main(): i32 {
  later handler()
  var x: i32 = 1_024;
  if x > 0 { io.Println("ok") }
  panic "boom";
  ret 0
}

def handler() {
  msg := recover();
  if msg != "" { io.Println("recovered:", msg) }
}