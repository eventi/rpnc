#!test/libs/bats/bin/bats

load 'libs/bats-support/load'
load 'libs/bats-assert/load'

@test "Should loop on [ ] for non zero on stack" {
  result=$(./rpnc -e "[1-d.d]" 4| xargs) # 1 is true
  assert_equal "$result" "3 2 1 0"
}
@test "Should nest loops" {
  result=$(./rpnc -e "[1-[1-d3%]d.d]" 10| xargs) # 1 is true
  assert_equal "$result" "6 3 0"
}
