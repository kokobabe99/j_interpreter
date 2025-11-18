pkg main
imp "std/io"

type Quacker interface {
  quack(): string
}

type Duck struct {
  name: string
}

type Person struct {
  name: string
}

def (d: Duck) quack(): string {
  ret d.name + " says QUACK!";
}

def (p: Person) quack(): string {
  ret p.name + " says: I'm pretending to be a trump!";
}

def makeItQuack(x: Quacker) {
  var msg string = x.quack();
  io.Println(msg);
}

def main() {
  var d Duck;
  d.name = "Donald";

  var p Person;
  p.name = "Jackie";

  makeItQuack(d);
  makeItQuack(p);
}
