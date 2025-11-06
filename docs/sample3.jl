pkg main
imp "std/io"

def worker(c: channel string) {
  c <- "hello";
  c <- "world"
}

def main() {
  var msgs: channel string;
  j worker(msgs);
  fr range msgs {
    io.Println("tick")
  }
}