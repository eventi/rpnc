#!test/libs/bats/bin/bats

load 'libs/bats-support/load'
load 'libs/bats-assert/load'

@test "Should be verbose with -v" {
  result=$(./rpnc -v -e "99 ." 1 | xargs)
  assert_equal "$result" "Input: Program: 99 . 99 Machine: ( 1 )"
}

@test "Should be verbose with --verbose" {
  result=$(./rpnc --verbose -e "99 ." 1 | xargs)
  assert_equal "$result" "Input: Program: 99 . 99 Machine: ( 1 )"
}

@test "Should take input with -i" {
  result=$(./rpnc -v -i "9" -e "#." | xargs)
  assert_equal "$result" "Input: 9 Program: #. 57 Machine: ( )"
}

@test "Should take input with --input" {
  result=$(./rpnc -v --input "9" -e "#." | xargs)
  assert_equal "$result" "Input: 9 Program: #. 57 Machine: ( )"
}

