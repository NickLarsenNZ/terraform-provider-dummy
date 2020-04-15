provider "dummy" {
  some_block = {
    some_key = "blah"
  }
}

resource "dummy_file" "d" {
  path       = "./thefile.txt"
  firstline  = "alpha"
  secondline = "bravo"
}
