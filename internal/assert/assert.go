package assert


func NoError(err error) {
  if err != nil {
    panic(err)
  }
}

func IntEqual(i1 int, i2 int, msg string) {
  if i1 != i2 {
    panic(msg)
  }
}
