pkg main
imp "std/io"

def worker(c: channel string) {
  c <- "hello";
  c <- "world";
  close(c);           
}

def main() {
  var msgs channel string = make(channel string, 2);
  j worker(msgs);
  fr range msgs {
    io.Println("tick");
  }
}