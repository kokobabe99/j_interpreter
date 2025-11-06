pkg main

def add(a: i32, b: i32): i32 {
  ret a + b * 2
}

def main(): i32 {
  var r: i32 = add(3, 4);
  ret r
}