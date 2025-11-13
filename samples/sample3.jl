pkg main
imp "std/io"

def worker(c: channel string) {
  c <- "hello";
  c <- "world";
  close(c);           
}

def main() {
  var msgs channel string = make(channel string, 2);
  j worker(msgs);        // v0 同步调用
  fr range msgs {        // 顺序消费队列里的两个元素
    io.Println("tick");
  }
}