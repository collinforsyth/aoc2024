package main

import (
	"collinforsyth/aoc2024/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay23(t *testing.T) {
	input := []byte(`kh-tc
qp-kh
de-cg
ka-co
yn-aq
qp-ub
cg-tb
vc-aq
tb-ka
wh-tc
yn-cg
kh-ub
ta-co
de-co
tc-td
tb-wq
wh-td
ta-ka
td-qp
aq-cg
wq-ub
ub-vc
de-ta
wq-aq
wq-vc
wh-yn
ka-de
kh-ta
co-tc
wh-qp
tb-vc
td-yn`)

	i := util.FromBytes(input)
	p := parseInput(i)
	assert.Equal(t, 7, partOne(p))

	assert.Equal(t, "co,de,ka,ta", partTwo(p))
}
