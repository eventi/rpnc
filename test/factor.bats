#!test/libs/bats/bin/bats

load 'libs/bats-support/load'
load 'libs/bats-assert/load'

@test "Should factor 99" {
  result=$( go run rpnc.go -e '[d2 0[+oo/o*rso=(s1 0:s2+d4=(1-)ood*>0s)](s)0*+oo=(0*:d./1)].' 99 | xargs )
  assert_equal $result "3 3 11"
}

