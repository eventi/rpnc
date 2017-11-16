#!test/libs/bats/bin/bats

load 'libs/bats-support/load'
load 'libs/bats-assert/load'

@test "Should find the integer square root of 99000000" {
  result=$(./rpnc -e 'd2/[drr oo/+2/ro=~]s0*+.' 99000000 | head -1)
  assert_equal $result 9949
}

