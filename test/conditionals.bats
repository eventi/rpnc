#!test/libs/bats/bin/bats

load 'libs/bats-support/load'
load 'libs/bats-assert/load'

@test "Should branch on ( : ) for non zero on stack" {
  program='d(99:66).'
  result=$(go run rpnc.go -e "$program" 1| head -1) # 1 is true
  assert_equal $result 99
  result=$(go run rpnc.go -e "$program" 9| head -1) # 9 is true
  assert_equal $result 99
  result=$(go run rpnc.go -e "$program" 0| head -1) # 0 is false
  assert_equal $result 66
}

fizzbuzz() { go run rpnc.go -e 'd3%(5%(0:5):5%(3:35))' $1 ; }

@test "Should nest conditionals" {
  program='d3%(5%(0:5):5%(3:35))'
  assert_equal $(fizzbuzz 99) 3
}

