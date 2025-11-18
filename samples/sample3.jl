pkg main
imp "std/io"

def worker(c: channel string) {
  c <- "hello";
  c <- "world"; //error
  close(c);           
}

def main() {
  var msgs channel string = make(channel string, 2); // make type 
  j worker(msgs);
  fr range msgs {
    io.Println("tick");
  }
}