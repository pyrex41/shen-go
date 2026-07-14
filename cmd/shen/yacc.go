package main

import . "github.com/tiancaiamao/shen-go/kl"

var YaccMain = MakeNative(func(__e *ControlFlow) {
tmp17364 := MakeNative(func(__e *ControlFlow) {
V112 := __e.Get(1)
_ = V112
V113 := __e.Get(2)
_ = V113
tmp17365 := MakeNative(func(__e *ControlFlow) {
W114 := __e.Get(1)
_ = W114
tmp17372 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W114)


if True == tmp17372 {
__e.Return(PrimSimpleError(MakeString("parse failure\n")))
return
} else {
tmp17370 := Call(__e, PrimFunc(symshen_4partial_1parse_1failure_2), W114)


if True == tmp17370 {
tmp17366 := Call(__e, PrimFunc(symshen_4in_1_6), W114)


tmp17367 := PrimSet(symshen_4_dresidue_d, tmp17366)

_ = tmp17367

tmp17368 := PrimValue(symshen_4_dresidue_d)

__e.TailApply(PrimFunc(symshen_4raise_1syntax_1error), tmp17368)
return


} else {
__e.TailApply(PrimFunc(symshen_4_5_1out), W114)
return
}


}


}, 1)

tmp17373 := Call(__e, V112, V113)


__e.TailApply(tmp17365, tmp17373)
return


}, 2)

tmp17374 := Call(__e, ns2_1set, symcompile, tmp17364)


_ = tmp17374

tmp17375 := MakeNative(func(__e *ControlFlow) {
V115 := __e.Get(1)
_ = V115
tmp17376 := PrimValue(sym_dmaximum_1print_1sequence_1size_d)

tmp17377 := Call(__e, PrimFunc(symshen_4syntax_1error_1message), tmp17376, MakeNumber(0), V115)


tmp17378 := PrimStringConcat(MakeString("syntax error here: "), tmp17377)

tmp17379 := Call(__e, PrimFunc(symshen_4proc_1nl), tmp17378)


__e.Return(PrimSimpleError(tmp17379))
return


}, 1)

tmp17380 := Call(__e, ns2_1set, symshen_4raise_1syntax_1error, tmp17375)


_ = tmp17380

tmp17381 := MakeNative(func(__e *ControlFlow) {
V123 := __e.Get(1)
_ = V123
V124 := __e.Get(2)
_ = V124
V125 := __e.Get(3)
_ = V125
tmp17392 := PrimEqual(Nil, V125)

if True == tmp17392 {
__e.Return(MakeString("\n"))
return
} else {
tmp17390 := PrimEqual(V123, V124)

if True == tmp17390 {
__e.Return(MakeString("...etc \n"))
return
} else {
tmp17388 := PrimIsPair(V125)

if True == tmp17388 {
tmp17382 := PrimHead(V125)

tmp17383 := Call(__e, PrimFunc(symshen_4app), tmp17382, MakeString(" "), symshen_4s)


tmp17384 := PrimNumberAdd(V124, MakeNumber(1))

tmp17385 := PrimTail(V125)

tmp17386 := Call(__e, PrimFunc(symshen_4syntax_1error_1message), V123, tmp17384, tmp17385)


__e.Return(PrimStringConcat(tmp17383, tmp17386))
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.syntax-error-message")))
return
}


}


}


}, 3)

tmp17393 := Call(__e, ns2_1set, symshen_4syntax_1error_1message, tmp17381)


_ = tmp17393

tmp17394 := MakeNative(func(__e *ControlFlow) {
V126 := __e.Get(1)
_ = V126
tmp17395 := Call(__e, PrimFunc(symfail))


__e.Return(PrimEqual(V126, tmp17395))
return


}, 1)

tmp17396 := Call(__e, ns2_1set, symshen_4parse_1failure_2, tmp17394)


_ = tmp17396

tmp17397 := MakeNative(func(__e *ControlFlow) {
V127 := __e.Get(1)
_ = V127
tmp17398 := Call(__e, PrimFunc(symshen_4in_1_6), V127)


__e.Return(PrimIsPair(tmp17398))
return


}, 1)

tmp17399 := Call(__e, ns2_1set, symshen_4partial_1parse_1failure_2, tmp17397)


_ = tmp17399

tmp17400 := MakeNative(func(__e *ControlFlow) {
V130 := __e.Get(1)
_ = V130
tmp17413 := PrimIsPair(V130)

var ifres17404 Obj

if True == tmp17413 {
tmp17411 := PrimTail(V130)

tmp17412 := PrimIsPair(tmp17411)

var ifres17406 Obj

if True == tmp17412 {
tmp17408 := PrimTail(V130)

tmp17409 := PrimTail(tmp17408)

tmp17410 := PrimEqual(Nil, tmp17409)

var ifres17407 Obj

if True == tmp17410 {
ifres17407 = True


} else {
ifres17407 = False


}

ifres17406 = ifres17407


} else {
ifres17406 = False


}

var ifres17405 Obj

if True == ifres17406 {
ifres17405 = True


} else {
ifres17405 = False


}

ifres17404 = ifres17405


} else {
ifres17404 = False


}

if True == ifres17404 {
tmp17401 := PrimTail(V130)

__e.Return(PrimHead(tmp17401))
return


} else {
tmp17402 := Call(__e, PrimFunc(symshen_4app), V130, MakeString(" is not a YACC stream\n"), symshen_4s)


__e.Return(PrimSimpleError(tmp17402))
return


}


}, 1)

tmp17414 := Call(__e, ns2_1set, symshen_4objectcode, tmp17400)


_ = tmp17414

tmp17415 := MakeNative(func(__e *ControlFlow) {
V131 := __e.Get(1)
_ = V131
tmp17416 := MakeNative(func(__e *ControlFlow) {
Z132 := __e.Get(1)
_ = Z132
__e.TailApply(PrimFunc(symshen_4_5yacc_6), Z132)
return
}, 1)

__e.TailApply(PrimFunc(symcompile), tmp17416, V131)
return


}, 1)

tmp17417 := Call(__e, ns2_1set, symshen_4yacc_1_6shen, tmp17415)


_ = tmp17417

tmp17418 := MakeNative(func(__e *ControlFlow) {
V133 := __e.Get(1)
_ = V133
tmp17419 := MakeNative(func(__e *ControlFlow) {
W134 := __e.Get(1)
_ = W134
tmp17421 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W134)


if True == tmp17421 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W134)
return
}


}, 1)

tmp17457 := PrimIsPair(V133)

var ifres17422 Obj

if True == tmp17457 {
tmp17423 := MakeNative(func(__e *ControlFlow) {
W135 := __e.Get(1)
_ = W135
tmp17424 := MakeNative(func(__e *ControlFlow) {
W136 := __e.Get(1)
_ = W136
tmp17425 := MakeNative(func(__e *ControlFlow) {
W137 := __e.Get(1)
_ = W137
tmp17451 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W137)


if True == tmp17451 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17426 := MakeNative(func(__e *ControlFlow) {
W138 := __e.Get(1)
_ = W138
tmp17427 := MakeNative(func(__e *ControlFlow) {
W139 := __e.Get(1)
_ = W139
tmp17428 := MakeNative(func(__e *ControlFlow) {
W140 := __e.Get(1)
_ = W140
tmp17446 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W140)


if True == tmp17446 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17429 := MakeNative(func(__e *ControlFlow) {
W141 := __e.Get(1)
_ = W141
tmp17430 := MakeNative(func(__e *ControlFlow) {
W142 := __e.Get(1)
_ = W142
tmp17431 := MakeNative(func(__e *ControlFlow) {
W143 := __e.Get(1)
_ = W143
tmp17432 := MakeNative(func(__e *ControlFlow) {
W144 := __e.Get(1)
_ = W144
__e.Return(W144)
return
}, 1)

tmp17433 := PrimCons(W135, Nil)

tmp17434 := PrimCons(symdefine, tmp17433)

tmp17435 := Call(__e, PrimFunc(symshen_4c_1rules_1_6shen), W138, W143, W141)


tmp17436 := PrimCons(tmp17435, Nil)

tmp17437 := PrimCons(sym_1_6, tmp17436)

tmp17438 := PrimCons(W143, tmp17437)

tmp17439 := Call(__e, PrimFunc(symappend), W138, tmp17438)


tmp17440 := Call(__e, PrimFunc(symappend), tmp17434, tmp17439)


__e.TailApply(tmp17432, tmp17440)
return


}, 1)

tmp17441 := Call(__e, PrimFunc(symgensym), symS)


tmp17442 := Call(__e, tmp17431, tmp17441)


__e.TailApply(PrimFunc(symshen_4comb), W142, tmp17442)
return


}, 1)

tmp17443 := Call(__e, PrimFunc(symshen_4in_1_6), W140)


__e.TailApply(tmp17430, tmp17443)
return


}, 1)

tmp17444 := Call(__e, PrimFunc(symshen_4_5_1out), W140)


__e.TailApply(tmp17429, tmp17444)
return


}


}, 1)

tmp17447 := Call(__e, PrimFunc(symshen_4_5c_1rules_6), W139)


__e.TailApply(tmp17428, tmp17447)
return


}, 1)

tmp17448 := Call(__e, PrimFunc(symshen_4in_1_6), W137)


__e.TailApply(tmp17427, tmp17448)
return


}, 1)

tmp17449 := Call(__e, PrimFunc(symshen_4_5_1out), W137)


__e.TailApply(tmp17426, tmp17449)
return


}


}, 1)

tmp17452 := Call(__e, PrimFunc(symshen_4_5yaccsig_6), W136)


__e.TailApply(tmp17425, tmp17452)
return


}, 1)

tmp17453 := Call(__e, PrimFunc(symtail), V133)


__e.TailApply(tmp17424, tmp17453)
return


}, 1)

tmp17454 := Call(__e, PrimFunc(symhead), V133)


tmp17455 := Call(__e, tmp17423, tmp17454)


ifres17422 = tmp17455


} else {
tmp17456 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17422 = tmp17456


}

__e.TailApply(tmp17419, ifres17422)
return


}, 1)

tmp17458 := Call(__e, ns2_1set, symshen_4_5yacc_6, tmp17418)


_ = tmp17458

tmp17459 := MakeNative(func(__e *ControlFlow) {
V145 := __e.Get(1)
_ = V145
tmp17460 := MakeNative(func(__e *ControlFlow) {
W146 := __e.Get(1)
_ = W146
tmp17472 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W146)


if True == tmp17472 {
tmp17461 := MakeNative(func(__e *ControlFlow) {
W161 := __e.Get(1)
_ = W161
tmp17463 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W161)


if True == tmp17463 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W161)
return
}


}, 1)

tmp17464 := MakeNative(func(__e *ControlFlow) {
W162 := __e.Get(1)
_ = W162
tmp17468 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W162)


if True == tmp17468 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17465 := MakeNative(func(__e *ControlFlow) {
W163 := __e.Get(1)
_ = W163
__e.TailApply(PrimFunc(symshen_4comb), W163, Nil)
return
}, 1)

tmp17466 := Call(__e, PrimFunc(symshen_4in_1_6), W162)


__e.TailApply(tmp17465, tmp17466)
return


}


}, 1)

tmp17469 := Call(__e, PrimFunc(sym_5e_6), V145)


tmp17470 := Call(__e, tmp17464, tmp17469)


__e.TailApply(tmp17461, tmp17470)
return


} else {
__e.Return(W146)
return
}


}, 1)

tmp17535 := PrimIsPair(V145)

var ifres17473 Obj

if True == tmp17535 {
tmp17474 := MakeNative(func(__e *ControlFlow) {
W147 := __e.Get(1)
_ = W147
tmp17475 := MakeNative(func(__e *ControlFlow) {
W148 := __e.Get(1)
_ = W148
tmp17530 := Call(__e, PrimFunc(symshen_4ccons_2), W148)


if True == tmp17530 {
tmp17476 := MakeNative(func(__e *ControlFlow) {
W149 := __e.Get(1)
_ = W149
tmp17477 := MakeNative(func(__e *ControlFlow) {
W150 := __e.Get(1)
_ = W150
tmp17526 := Call(__e, PrimFunc(symshen_4hds_a_2), W149, symlist)


if True == tmp17526 {
tmp17478 := MakeNative(func(__e *ControlFlow) {
W151 := __e.Get(1)
_ = W151
tmp17523 := PrimIsPair(W151)

if True == tmp17523 {
tmp17479 := MakeNative(func(__e *ControlFlow) {
W152 := __e.Get(1)
_ = W152
tmp17480 := MakeNative(func(__e *ControlFlow) {
W153 := __e.Get(1)
_ = W153
tmp17481 := MakeNative(func(__e *ControlFlow) {
W154 := __e.Get(1)
_ = W154
tmp17518 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W154)


if True == tmp17518 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17482 := MakeNative(func(__e *ControlFlow) {
W155 := __e.Get(1)
_ = W155
tmp17515 := Call(__e, PrimFunc(symshen_4hds_a_2), W150, sym_a_a_6)


if True == tmp17515 {
tmp17483 := MakeNative(func(__e *ControlFlow) {
W156 := __e.Get(1)
_ = W156
tmp17512 := PrimIsPair(W156)

if True == tmp17512 {
tmp17484 := MakeNative(func(__e *ControlFlow) {
W157 := __e.Get(1)
_ = W157
tmp17485 := MakeNative(func(__e *ControlFlow) {
W158 := __e.Get(1)
_ = W158
tmp17508 := PrimIsPair(W158)

if True == tmp17508 {
tmp17486 := MakeNative(func(__e *ControlFlow) {
W159 := __e.Get(1)
_ = W159
tmp17487 := MakeNative(func(__e *ControlFlow) {
W160 := __e.Get(1)
_ = W160
tmp17504 := PrimEqual(sym_i, W147)

var ifres17501 Obj

if True == tmp17504 {
tmp17503 := PrimEqual(sym_j, W159)

var ifres17502 Obj

if True == tmp17503 {
ifres17502 = True


} else {
ifres17502 = False


}

ifres17501 = ifres17502


} else {
ifres17501 = False


}

if True == ifres17501 {
tmp17488 := PrimCons(W152, Nil)

tmp17489 := PrimCons(symlist, tmp17488)

tmp17490 := PrimCons(W152, Nil)

tmp17491 := PrimCons(symlist, tmp17490)

tmp17492 := PrimCons(W157, Nil)

tmp17493 := PrimCons(tmp17491, tmp17492)

tmp17494 := PrimCons(symstr, tmp17493)

tmp17495 := PrimCons(sym_j, Nil)

tmp17496 := PrimCons(tmp17494, tmp17495)

tmp17497 := PrimCons(sym_1_1_6, tmp17496)

tmp17498 := PrimCons(tmp17489, tmp17497)

tmp17499 := PrimCons(sym_i, tmp17498)

__e.TailApply(PrimFunc(symshen_4comb), W160, tmp17499)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp17505 := Call(__e, PrimFunc(symtail), W158)


__e.TailApply(tmp17487, tmp17505)
return


}, 1)

tmp17506 := Call(__e, PrimFunc(symhead), W158)


__e.TailApply(tmp17486, tmp17506)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp17509 := Call(__e, PrimFunc(symtail), W156)


__e.TailApply(tmp17485, tmp17509)
return


}, 1)

tmp17510 := Call(__e, PrimFunc(symhead), W156)


__e.TailApply(tmp17484, tmp17510)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp17513 := Call(__e, PrimFunc(symtail), W150)


__e.TailApply(tmp17483, tmp17513)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp17516 := Call(__e, PrimFunc(symshen_4in_1_6), W154)


__e.TailApply(tmp17482, tmp17516)
return


}


}, 1)

tmp17519 := Call(__e, PrimFunc(sym_5end_6), W153)


__e.TailApply(tmp17481, tmp17519)
return


}, 1)

tmp17520 := Call(__e, PrimFunc(symtail), W151)


__e.TailApply(tmp17480, tmp17520)
return


}, 1)

tmp17521 := Call(__e, PrimFunc(symhead), W151)


__e.TailApply(tmp17479, tmp17521)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp17524 := Call(__e, PrimFunc(symtail), W149)


__e.TailApply(tmp17478, tmp17524)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp17527 := Call(__e, PrimFunc(symtail), W148)


__e.TailApply(tmp17477, tmp17527)
return


}, 1)

tmp17528 := Call(__e, PrimFunc(symhead), W148)


__e.TailApply(tmp17476, tmp17528)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp17531 := Call(__e, PrimFunc(symtail), V145)


__e.TailApply(tmp17475, tmp17531)
return


}, 1)

tmp17532 := Call(__e, PrimFunc(symhead), V145)


tmp17533 := Call(__e, tmp17474, tmp17532)


ifres17473 = tmp17533


} else {
tmp17534 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17473 = tmp17534


}

__e.TailApply(tmp17460, ifres17473)
return


}, 1)

tmp17536 := Call(__e, ns2_1set, symshen_4_5yaccsig_6, tmp17459)


_ = tmp17536

tmp17537 := MakeNative(func(__e *ControlFlow) {
V164 := __e.Get(1)
_ = V164
tmp17538 := MakeNative(func(__e *ControlFlow) {
W165 := __e.Get(1)
_ = W165
tmp17557 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W165)


if True == tmp17557 {
tmp17539 := MakeNative(func(__e *ControlFlow) {
W172 := __e.Get(1)
_ = W172
tmp17541 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W172)


if True == tmp17541 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W172)
return
}


}, 1)

tmp17542 := MakeNative(func(__e *ControlFlow) {
W173 := __e.Get(1)
_ = W173
tmp17553 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W173)


if True == tmp17553 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17543 := MakeNative(func(__e *ControlFlow) {
W174 := __e.Get(1)
_ = W174
tmp17544 := MakeNative(func(__e *ControlFlow) {
W175 := __e.Get(1)
_ = W175
tmp17549 := Call(__e, PrimFunc(symempty_2), W174)


var ifres17545 Obj

if True == tmp17549 {
ifres17545 = Nil


} else {
tmp17546 := Call(__e, PrimFunc(symshen_4app), W174, MakeString("\n ..."), symshen_4r)


tmp17547 := PrimStringConcat(MakeString("YACC syntax error here:\n "), tmp17546)

tmp17548 := PrimSimpleError(tmp17547)

ifres17545 = tmp17548


}

__e.TailApply(PrimFunc(symshen_4comb), W175, ifres17545)
return


}, 1)

tmp17550 := Call(__e, PrimFunc(symshen_4in_1_6), W173)


__e.TailApply(tmp17544, tmp17550)
return


}, 1)

tmp17551 := Call(__e, PrimFunc(symshen_4_5_1out), W173)


__e.TailApply(tmp17543, tmp17551)
return


}


}, 1)

tmp17554 := Call(__e, PrimFunc(sym_5_b_6), V164)


tmp17555 := Call(__e, tmp17542, tmp17554)


__e.TailApply(tmp17539, tmp17555)
return


} else {
__e.Return(W165)
return
}


}, 1)

tmp17558 := MakeNative(func(__e *ControlFlow) {
W166 := __e.Get(1)
_ = W166
tmp17573 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W166)


if True == tmp17573 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17559 := MakeNative(func(__e *ControlFlow) {
W167 := __e.Get(1)
_ = W167
tmp17560 := MakeNative(func(__e *ControlFlow) {
W168 := __e.Get(1)
_ = W168
tmp17561 := MakeNative(func(__e *ControlFlow) {
W169 := __e.Get(1)
_ = W169
tmp17568 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W169)


if True == tmp17568 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17562 := MakeNative(func(__e *ControlFlow) {
W170 := __e.Get(1)
_ = W170
tmp17563 := MakeNative(func(__e *ControlFlow) {
W171 := __e.Get(1)
_ = W171
tmp17564 := PrimCons(W167, W170)

__e.TailApply(PrimFunc(symshen_4comb), W171, tmp17564)
return


}, 1)

tmp17565 := Call(__e, PrimFunc(symshen_4in_1_6), W169)


__e.TailApply(tmp17563, tmp17565)
return


}, 1)

tmp17566 := Call(__e, PrimFunc(symshen_4_5_1out), W169)


__e.TailApply(tmp17562, tmp17566)
return


}


}, 1)

tmp17569 := Call(__e, PrimFunc(symshen_4_5c_1rules_6), W168)


__e.TailApply(tmp17561, tmp17569)
return


}, 1)

tmp17570 := Call(__e, PrimFunc(symshen_4in_1_6), W166)


__e.TailApply(tmp17560, tmp17570)
return


}, 1)

tmp17571 := Call(__e, PrimFunc(symshen_4_5_1out), W166)


__e.TailApply(tmp17559, tmp17571)
return


}


}, 1)

tmp17574 := Call(__e, PrimFunc(symshen_4_5c_1rule_6), V164)


tmp17575 := Call(__e, tmp17558, tmp17574)


__e.TailApply(tmp17538, tmp17575)
return


}, 1)

tmp17576 := Call(__e, ns2_1set, symshen_4_5c_1rules_6, tmp17537)


_ = tmp17576

tmp17577 := MakeNative(func(__e *ControlFlow) {
V176 := __e.Get(1)
_ = V176
tmp17578 := MakeNative(func(__e *ControlFlow) {
W177 := __e.Get(1)
_ = W177
tmp17601 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W177)


if True == tmp17601 {
tmp17579 := MakeNative(func(__e *ControlFlow) {
W186 := __e.Get(1)
_ = W186
tmp17581 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W186)


if True == tmp17581 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W186)
return
}


}, 1)

tmp17582 := MakeNative(func(__e *ControlFlow) {
W187 := __e.Get(1)
_ = W187
tmp17597 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W187)


if True == tmp17597 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17583 := MakeNative(func(__e *ControlFlow) {
W188 := __e.Get(1)
_ = W188
tmp17584 := MakeNative(func(__e *ControlFlow) {
W189 := __e.Get(1)
_ = W189
tmp17585 := MakeNative(func(__e *ControlFlow) {
W190 := __e.Get(1)
_ = W190
tmp17592 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W190)


if True == tmp17592 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17586 := MakeNative(func(__e *ControlFlow) {
W191 := __e.Get(1)
_ = W191
tmp17587 := Call(__e, PrimFunc(symshen_4autocomplete), W188)


tmp17588 := PrimCons(tmp17587, Nil)

tmp17589 := PrimCons(W188, tmp17588)

__e.TailApply(PrimFunc(symshen_4comb), W191, tmp17589)
return


}, 1)

tmp17590 := Call(__e, PrimFunc(symshen_4in_1_6), W190)


__e.TailApply(tmp17586, tmp17590)
return


}


}, 1)

tmp17593 := Call(__e, PrimFunc(symshen_4_5sc_6), W189)


__e.TailApply(tmp17585, tmp17593)
return


}, 1)

tmp17594 := Call(__e, PrimFunc(symshen_4in_1_6), W187)


__e.TailApply(tmp17584, tmp17594)
return


}, 1)

tmp17595 := Call(__e, PrimFunc(symshen_4_5_1out), W187)


__e.TailApply(tmp17583, tmp17595)
return


}


}, 1)

tmp17598 := Call(__e, PrimFunc(symshen_4_5syntax_6), V176)


tmp17599 := Call(__e, tmp17582, tmp17598)


__e.TailApply(tmp17579, tmp17599)
return


} else {
__e.Return(W177)
return
}


}, 1)

tmp17602 := MakeNative(func(__e *ControlFlow) {
W178 := __e.Get(1)
_ = W178
tmp17624 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W178)


if True == tmp17624 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17603 := MakeNative(func(__e *ControlFlow) {
W179 := __e.Get(1)
_ = W179
tmp17604 := MakeNative(func(__e *ControlFlow) {
W180 := __e.Get(1)
_ = W180
tmp17605 := MakeNative(func(__e *ControlFlow) {
W181 := __e.Get(1)
_ = W181
tmp17619 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W181)


if True == tmp17619 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17606 := MakeNative(func(__e *ControlFlow) {
W182 := __e.Get(1)
_ = W182
tmp17607 := MakeNative(func(__e *ControlFlow) {
W183 := __e.Get(1)
_ = W183
tmp17608 := MakeNative(func(__e *ControlFlow) {
W184 := __e.Get(1)
_ = W184
tmp17614 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W184)


if True == tmp17614 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17609 := MakeNative(func(__e *ControlFlow) {
W185 := __e.Get(1)
_ = W185
tmp17610 := PrimCons(W182, Nil)

tmp17611 := PrimCons(W179, tmp17610)

__e.TailApply(PrimFunc(symshen_4comb), W185, tmp17611)
return


}, 1)

tmp17612 := Call(__e, PrimFunc(symshen_4in_1_6), W184)


__e.TailApply(tmp17609, tmp17612)
return


}


}, 1)

tmp17615 := Call(__e, PrimFunc(symshen_4_5sc_6), W183)


__e.TailApply(tmp17608, tmp17615)
return


}, 1)

tmp17616 := Call(__e, PrimFunc(symshen_4in_1_6), W181)


__e.TailApply(tmp17607, tmp17616)
return


}, 1)

tmp17617 := Call(__e, PrimFunc(symshen_4_5_1out), W181)


__e.TailApply(tmp17606, tmp17617)
return


}


}, 1)

tmp17620 := Call(__e, PrimFunc(symshen_4_5semantics_6), W180)


__e.TailApply(tmp17605, tmp17620)
return


}, 1)

tmp17621 := Call(__e, PrimFunc(symshen_4in_1_6), W178)


__e.TailApply(tmp17604, tmp17621)
return


}, 1)

tmp17622 := Call(__e, PrimFunc(symshen_4_5_1out), W178)


__e.TailApply(tmp17603, tmp17622)
return


}


}, 1)

tmp17625 := Call(__e, PrimFunc(symshen_4_5syntax_6), V176)


tmp17626 := Call(__e, tmp17602, tmp17625)


__e.TailApply(tmp17578, tmp17626)
return


}, 1)

tmp17627 := Call(__e, ns2_1set, symshen_4_5c_1rule_6, tmp17577)


_ = tmp17627

tmp17628 := MakeNative(func(__e *ControlFlow) {
V192 := __e.Get(1)
_ = V192
tmp17657 := PrimIsPair(V192)

var ifres17649 Obj

if True == tmp17657 {
tmp17655 := PrimTail(V192)

tmp17656 := PrimEqual(Nil, tmp17655)

var ifres17651 Obj

if True == tmp17656 {
tmp17653 := PrimHead(V192)

tmp17654 := Call(__e, PrimFunc(symshen_4non_1terminal_2), tmp17653)


var ifres17652 Obj

if True == tmp17654 {
ifres17652 = True


} else {
ifres17652 = False


}

ifres17651 = ifres17652


} else {
ifres17651 = False


}

var ifres17650 Obj

if True == ifres17651 {
ifres17650 = True


} else {
ifres17650 = False


}

ifres17649 = ifres17650


} else {
ifres17649 = False


}

if True == ifres17649 {
__e.Return(PrimHead(V192))
return
} else {
tmp17647 := PrimIsPair(V192)

var ifres17643 Obj

if True == tmp17647 {
tmp17645 := PrimHead(V192)

tmp17646 := Call(__e, PrimFunc(symshen_4non_1terminal_2), tmp17645)


var ifres17644 Obj

if True == tmp17646 {
ifres17644 = True


} else {
ifres17644 = False


}

ifres17643 = ifres17644


} else {
ifres17643 = False


}

if True == ifres17643 {
tmp17629 := PrimHead(V192)

tmp17630 := PrimTail(V192)

tmp17631 := Call(__e, PrimFunc(symshen_4autocomplete), tmp17630)


tmp17632 := PrimCons(tmp17631, Nil)

tmp17633 := PrimCons(tmp17629, tmp17632)

__e.Return(PrimCons(symappend, tmp17633))
return


} else {
tmp17641 := PrimIsPair(V192)

if True == tmp17641 {
tmp17634 := PrimHead(V192)

tmp17635 := Call(__e, PrimFunc(symshen_4autocomplete), tmp17634)


tmp17636 := PrimTail(V192)

tmp17637 := Call(__e, PrimFunc(symshen_4autocomplete), tmp17636)


tmp17638 := PrimCons(tmp17637, Nil)

tmp17639 := PrimCons(tmp17635, tmp17638)

__e.Return(PrimCons(symcons, tmp17639))
return


} else {
__e.Return(V192)
return
}


}


}


}, 1)

tmp17658 := Call(__e, ns2_1set, symshen_4autocomplete, tmp17628)


_ = tmp17658

tmp17659 := MakeNative(func(__e *ControlFlow) {
V193 := __e.Get(1)
_ = V193
tmp17666 := PrimIsSymbol(V193)

if True == tmp17666 {
tmp17661 := MakeNative(func(__e *ControlFlow) {
W194 := __e.Get(1)
_ = W194
tmp17662 := MakeNative(func(__e *ControlFlow) {
Z195 := __e.Get(1)
_ = Z195
__e.TailApply(PrimFunc(symshen_4_5non_1terminal_2_6), Z195)
return
}, 1)

__e.TailApply(PrimFunc(symcompile), tmp17662, W194)
return


}, 1)

tmp17663 := Call(__e, PrimFunc(symexplode), V193)


tmp17664 := Call(__e, tmp17661, tmp17663)


if True == tmp17664 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


} else {
__e.Return(False)
return
}


}, 1)

tmp17667 := Call(__e, ns2_1set, symshen_4non_1terminal_2, tmp17659)


_ = tmp17667

tmp17668 := MakeNative(func(__e *ControlFlow) {
V196 := __e.Get(1)
_ = V196
tmp17669 := MakeNative(func(__e *ControlFlow) {
W197 := __e.Get(1)
_ = W197
tmp17691 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W197)


if True == tmp17691 {
tmp17670 := MakeNative(func(__e *ControlFlow) {
W202 := __e.Get(1)
_ = W202
tmp17682 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W202)


if True == tmp17682 {
tmp17671 := MakeNative(func(__e *ControlFlow) {
W205 := __e.Get(1)
_ = W205
tmp17673 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W205)


if True == tmp17673 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W205)
return
}


}, 1)

tmp17674 := MakeNative(func(__e *ControlFlow) {
W206 := __e.Get(1)
_ = W206
tmp17678 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W206)


if True == tmp17678 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17675 := MakeNative(func(__e *ControlFlow) {
W207 := __e.Get(1)
_ = W207
__e.TailApply(PrimFunc(symshen_4comb), W207, False)
return
}, 1)

tmp17676 := Call(__e, PrimFunc(symshen_4in_1_6), W206)


__e.TailApply(tmp17675, tmp17676)
return


}


}, 1)

tmp17679 := Call(__e, PrimFunc(sym_5_b_6), V196)


tmp17680 := Call(__e, tmp17674, tmp17679)


__e.TailApply(tmp17671, tmp17680)
return


} else {
__e.Return(W202)
return
}


}, 1)

tmp17683 := MakeNative(func(__e *ControlFlow) {
W203 := __e.Get(1)
_ = W203
tmp17687 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W203)


if True == tmp17687 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17684 := MakeNative(func(__e *ControlFlow) {
W204 := __e.Get(1)
_ = W204
__e.TailApply(PrimFunc(symshen_4comb), W204, True)
return
}, 1)

tmp17685 := Call(__e, PrimFunc(symshen_4in_1_6), W203)


__e.TailApply(tmp17684, tmp17685)
return


}


}, 1)

tmp17688 := Call(__e, PrimFunc(symshen_4_5non_1terminal_1name_6), V196)


tmp17689 := Call(__e, tmp17683, tmp17688)


__e.TailApply(tmp17670, tmp17689)
return


} else {
__e.Return(W197)
return
}


}, 1)

tmp17692 := MakeNative(func(__e *ControlFlow) {
W198 := __e.Get(1)
_ = W198
tmp17702 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W198)


if True == tmp17702 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17693 := MakeNative(func(__e *ControlFlow) {
W199 := __e.Get(1)
_ = W199
tmp17694 := MakeNative(func(__e *ControlFlow) {
W200 := __e.Get(1)
_ = W200
tmp17698 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W200)


if True == tmp17698 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17695 := MakeNative(func(__e *ControlFlow) {
W201 := __e.Get(1)
_ = W201
__e.TailApply(PrimFunc(symshen_4comb), W201, True)
return
}, 1)

tmp17696 := Call(__e, PrimFunc(symshen_4in_1_6), W200)


__e.TailApply(tmp17695, tmp17696)
return


}


}, 1)

tmp17699 := Call(__e, PrimFunc(symshen_4_5non_1terminal_1name_6), W199)


__e.TailApply(tmp17694, tmp17699)
return


}, 1)

tmp17700 := Call(__e, PrimFunc(symshen_4in_1_6), W198)


__e.TailApply(tmp17693, tmp17700)
return


}


}, 1)

tmp17703 := Call(__e, PrimFunc(symshen_4_5packagenames_6), V196)


tmp17704 := Call(__e, tmp17692, tmp17703)


__e.TailApply(tmp17669, tmp17704)
return


}, 1)

tmp17705 := Call(__e, ns2_1set, symshen_4_5non_1terminal_2_6, tmp17668)


_ = tmp17705

tmp17706 := MakeNative(func(__e *ControlFlow) {
V208 := __e.Get(1)
_ = V208
tmp17707 := MakeNative(func(__e *ControlFlow) {
W209 := __e.Get(1)
_ = W209
tmp17723 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W209)


if True == tmp17723 {
tmp17708 := MakeNative(func(__e *ControlFlow) {
W215 := __e.Get(1)
_ = W215
tmp17710 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W215)


if True == tmp17710 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W215)
return
}


}, 1)

tmp17711 := MakeNative(func(__e *ControlFlow) {
W216 := __e.Get(1)
_ = W216
tmp17719 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W216)


if True == tmp17719 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17712 := MakeNative(func(__e *ControlFlow) {
W217 := __e.Get(1)
_ = W217
tmp17716 := Call(__e, PrimFunc(symshen_4hds_a_2), W217, MakeString("."))


if True == tmp17716 {
tmp17713 := MakeNative(func(__e *ControlFlow) {
W218 := __e.Get(1)
_ = W218
__e.TailApply(PrimFunc(symshen_4comb), W218, symshen_4skip)
return
}, 1)

tmp17714 := Call(__e, PrimFunc(symtail), W217)


__e.TailApply(tmp17713, tmp17714)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp17717 := Call(__e, PrimFunc(symshen_4in_1_6), W216)


__e.TailApply(tmp17712, tmp17717)
return


}


}, 1)

tmp17720 := Call(__e, PrimFunc(symshen_4_5packagename_6), V208)


tmp17721 := Call(__e, tmp17711, tmp17720)


__e.TailApply(tmp17708, tmp17721)
return


} else {
__e.Return(W209)
return
}


}, 1)

tmp17724 := MakeNative(func(__e *ControlFlow) {
W210 := __e.Get(1)
_ = W210
tmp17738 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W210)


if True == tmp17738 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17725 := MakeNative(func(__e *ControlFlow) {
W211 := __e.Get(1)
_ = W211
tmp17735 := Call(__e, PrimFunc(symshen_4hds_a_2), W211, MakeString("."))


if True == tmp17735 {
tmp17726 := MakeNative(func(__e *ControlFlow) {
W212 := __e.Get(1)
_ = W212
tmp17727 := MakeNative(func(__e *ControlFlow) {
W213 := __e.Get(1)
_ = W213
tmp17731 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W213)


if True == tmp17731 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17728 := MakeNative(func(__e *ControlFlow) {
W214 := __e.Get(1)
_ = W214
__e.TailApply(PrimFunc(symshen_4comb), W214, symshen_4skip)
return
}, 1)

tmp17729 := Call(__e, PrimFunc(symshen_4in_1_6), W213)


__e.TailApply(tmp17728, tmp17729)
return


}


}, 1)

tmp17732 := Call(__e, PrimFunc(symshen_4_5packagenames_6), W212)


__e.TailApply(tmp17727, tmp17732)
return


}, 1)

tmp17733 := Call(__e, PrimFunc(symtail), W211)


__e.TailApply(tmp17726, tmp17733)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp17736 := Call(__e, PrimFunc(symshen_4in_1_6), W210)


__e.TailApply(tmp17725, tmp17736)
return


}


}, 1)

tmp17739 := Call(__e, PrimFunc(symshen_4_5packagename_6), V208)


tmp17740 := Call(__e, tmp17724, tmp17739)


__e.TailApply(tmp17707, tmp17740)
return


}, 1)

tmp17741 := Call(__e, ns2_1set, symshen_4_5packagenames_6, tmp17706)


_ = tmp17741

tmp17742 := MakeNative(func(__e *ControlFlow) {
V219 := __e.Get(1)
_ = V219
tmp17743 := MakeNative(func(__e *ControlFlow) {
W220 := __e.Get(1)
_ = W220
tmp17755 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W220)


if True == tmp17755 {
tmp17744 := MakeNative(func(__e *ControlFlow) {
W225 := __e.Get(1)
_ = W225
tmp17746 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W225)


if True == tmp17746 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W225)
return
}


}, 1)

tmp17747 := MakeNative(func(__e *ControlFlow) {
W226 := __e.Get(1)
_ = W226
tmp17751 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W226)


if True == tmp17751 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17748 := MakeNative(func(__e *ControlFlow) {
W227 := __e.Get(1)
_ = W227
__e.TailApply(PrimFunc(symshen_4comb), W227, symshen_4skip)
return
}, 1)

tmp17749 := Call(__e, PrimFunc(symshen_4in_1_6), W226)


__e.TailApply(tmp17748, tmp17749)
return


}


}, 1)

tmp17752 := Call(__e, PrimFunc(sym_5e_6), V219)


tmp17753 := Call(__e, tmp17747, tmp17752)


__e.TailApply(tmp17744, tmp17753)
return


} else {
__e.Return(W220)
return
}


}, 1)

tmp17756 := MakeNative(func(__e *ControlFlow) {
W221 := __e.Get(1)
_ = W221
tmp17766 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W221)


if True == tmp17766 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17757 := MakeNative(func(__e *ControlFlow) {
W222 := __e.Get(1)
_ = W222
tmp17758 := MakeNative(func(__e *ControlFlow) {
W223 := __e.Get(1)
_ = W223
tmp17762 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W223)


if True == tmp17762 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17759 := MakeNative(func(__e *ControlFlow) {
W224 := __e.Get(1)
_ = W224
__e.TailApply(PrimFunc(symshen_4comb), W224, symshen_4skip)
return
}, 1)

tmp17760 := Call(__e, PrimFunc(symshen_4in_1_6), W223)


__e.TailApply(tmp17759, tmp17760)
return


}


}, 1)

tmp17763 := Call(__e, PrimFunc(symshen_4_5packagename_6), W222)


__e.TailApply(tmp17758, tmp17763)
return


}, 1)

tmp17764 := Call(__e, PrimFunc(symshen_4in_1_6), W221)


__e.TailApply(tmp17757, tmp17764)
return


}


}, 1)

tmp17767 := Call(__e, PrimFunc(symshen_4_5packagechar_6), V219)


tmp17768 := Call(__e, tmp17756, tmp17767)


__e.TailApply(tmp17743, tmp17768)
return


}, 1)

tmp17769 := Call(__e, ns2_1set, symshen_4_5packagename_6, tmp17742)


_ = tmp17769

tmp17770 := MakeNative(func(__e *ControlFlow) {
V228 := __e.Get(1)
_ = V228
tmp17771 := MakeNative(func(__e *ControlFlow) {
W229 := __e.Get(1)
_ = W229
tmp17773 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W229)


if True == tmp17773 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W229)
return
}


}, 1)

tmp17784 := PrimIsPair(V228)

var ifres17774 Obj

if True == tmp17784 {
tmp17775 := MakeNative(func(__e *ControlFlow) {
W230 := __e.Get(1)
_ = W230
tmp17776 := MakeNative(func(__e *ControlFlow) {
W231 := __e.Get(1)
_ = W231
tmp17778 := PrimEqual(W230, MakeString("."))

tmp17779 := PrimNot(tmp17778)

if True == tmp17779 {
__e.TailApply(PrimFunc(symshen_4comb), W231, symshen_4skip)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp17780 := Call(__e, PrimFunc(symtail), V228)


__e.TailApply(tmp17776, tmp17780)
return


}, 1)

tmp17781 := Call(__e, PrimFunc(symhead), V228)


tmp17782 := Call(__e, tmp17775, tmp17781)


ifres17774 = tmp17782


} else {
tmp17783 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17774 = tmp17783


}

__e.TailApply(tmp17771, ifres17774)
return


}, 1)

tmp17785 := Call(__e, ns2_1set, symshen_4_5packagechar_6, tmp17770)


_ = tmp17785

tmp17786 := MakeNative(func(__e *ControlFlow) {
V232 := __e.Get(1)
_ = V232
tmp17787 := MakeNative(func(__e *ControlFlow) {
W233 := __e.Get(1)
_ = W233
tmp17789 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W233)


if True == tmp17789 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W233)
return
}


}, 1)

tmp17812 := Call(__e, PrimFunc(symshen_4hds_a_2), V232, MakeString("<"))


var ifres17790 Obj

if True == tmp17812 {
tmp17791 := MakeNative(func(__e *ControlFlow) {
W234 := __e.Get(1)
_ = W234
tmp17792 := MakeNative(func(__e *ControlFlow) {
W235 := __e.Get(1)
_ = W235
tmp17807 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W235)


if True == tmp17807 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17793 := MakeNative(func(__e *ControlFlow) {
W236 := __e.Get(1)
_ = W236
tmp17794 := MakeNative(func(__e *ControlFlow) {
W237 := __e.Get(1)
_ = W237
tmp17796 := MakeNative(func(__e *ControlFlow) {
W238 := __e.Get(1)
_ = W238
tmp17801 := PrimIsPair(W238)

if True == tmp17801 {
tmp17798 := PrimHead(W238)

tmp17799 := PrimEqual(tmp17798, MakeString(">"))

if True == tmp17799 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


} else {
__e.Return(False)
return
}


}, 1)

tmp17802 := Call(__e, PrimFunc(symreverse), W236)


tmp17803 := Call(__e, tmp17796, tmp17802)


if True == tmp17803 {
__e.TailApply(PrimFunc(symshen_4comb), W237, symshen_4skip)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp17804 := Call(__e, PrimFunc(symshen_4in_1_6), W235)


__e.TailApply(tmp17794, tmp17804)
return


}, 1)

tmp17805 := Call(__e, PrimFunc(symshen_4_5_1out), W235)


__e.TailApply(tmp17793, tmp17805)
return


}


}, 1)

tmp17808 := Call(__e, PrimFunc(sym_5_b_6), W234)


__e.TailApply(tmp17792, tmp17808)
return


}, 1)

tmp17809 := Call(__e, PrimFunc(symtail), V232)


tmp17810 := Call(__e, tmp17791, tmp17809)


ifres17790 = tmp17810


} else {
tmp17811 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17790 = tmp17811


}

__e.TailApply(tmp17787, ifres17790)
return


}, 1)

tmp17813 := Call(__e, ns2_1set, symshen_4_5non_1terminal_1name_6, tmp17786)


_ = tmp17813

tmp17814 := MakeNative(func(__e *ControlFlow) {
V239 := __e.Get(1)
_ = V239
tmp17815 := PrimIntern(MakeString(";"))

__e.Return(PrimEqual(V239, tmp17815))
return


}, 1)

tmp17816 := Call(__e, ns2_1set, symshen_4semicolon_2, tmp17814)


_ = tmp17816

tmp17817 := MakeNative(func(__e *ControlFlow) {
V240 := __e.Get(1)
_ = V240
tmp17818 := MakeNative(func(__e *ControlFlow) {
W241 := __e.Get(1)
_ = W241
tmp17820 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W241)


if True == tmp17820 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W241)
return
}


}, 1)

tmp17830 := PrimIsPair(V240)

var ifres17821 Obj

if True == tmp17830 {
tmp17822 := MakeNative(func(__e *ControlFlow) {
W242 := __e.Get(1)
_ = W242
tmp17823 := MakeNative(func(__e *ControlFlow) {
W243 := __e.Get(1)
_ = W243
tmp17825 := Call(__e, PrimFunc(symshen_4colon_1equal_2), W242)


if True == tmp17825 {
__e.TailApply(PrimFunc(symshen_4comb), W243, symshen_4skip)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp17826 := Call(__e, PrimFunc(symtail), V240)


__e.TailApply(tmp17823, tmp17826)
return


}, 1)

tmp17827 := Call(__e, PrimFunc(symhead), V240)


tmp17828 := Call(__e, tmp17822, tmp17827)


ifres17821 = tmp17828


} else {
tmp17829 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17821 = tmp17829


}

__e.TailApply(tmp17818, ifres17821)
return


}, 1)

tmp17831 := Call(__e, ns2_1set, symshen_4_5colon_1equal_6, tmp17817)


_ = tmp17831

tmp17832 := MakeNative(func(__e *ControlFlow) {
V244 := __e.Get(1)
_ = V244
tmp17833 := PrimIntern(MakeString(":="))

__e.Return(PrimEqual(tmp17833, V244))
return


}, 1)

tmp17834 := Call(__e, ns2_1set, symshen_4colon_1equal_2, tmp17832)


_ = tmp17834

tmp17835 := MakeNative(func(__e *ControlFlow) {
V245 := __e.Get(1)
_ = V245
tmp17836 := MakeNative(func(__e *ControlFlow) {
W246 := __e.Get(1)
_ = W246
tmp17851 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W246)


if True == tmp17851 {
tmp17837 := MakeNative(func(__e *ControlFlow) {
W253 := __e.Get(1)
_ = W253
tmp17839 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W253)


if True == tmp17839 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W253)
return
}


}, 1)

tmp17840 := MakeNative(func(__e *ControlFlow) {
W254 := __e.Get(1)
_ = W254
tmp17847 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W254)


if True == tmp17847 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17841 := MakeNative(func(__e *ControlFlow) {
W255 := __e.Get(1)
_ = W255
tmp17842 := MakeNative(func(__e *ControlFlow) {
W256 := __e.Get(1)
_ = W256
tmp17843 := PrimCons(W255, Nil)

__e.TailApply(PrimFunc(symshen_4comb), W256, tmp17843)
return


}, 1)

tmp17844 := Call(__e, PrimFunc(symshen_4in_1_6), W254)


__e.TailApply(tmp17842, tmp17844)
return


}, 1)

tmp17845 := Call(__e, PrimFunc(symshen_4_5_1out), W254)


__e.TailApply(tmp17841, tmp17845)
return


}


}, 1)

tmp17848 := Call(__e, PrimFunc(symshen_4_5syntax_1item_6), V245)


tmp17849 := Call(__e, tmp17840, tmp17848)


__e.TailApply(tmp17837, tmp17849)
return


} else {
__e.Return(W246)
return
}


}, 1)

tmp17852 := MakeNative(func(__e *ControlFlow) {
W247 := __e.Get(1)
_ = W247
tmp17867 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W247)


if True == tmp17867 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17853 := MakeNative(func(__e *ControlFlow) {
W248 := __e.Get(1)
_ = W248
tmp17854 := MakeNative(func(__e *ControlFlow) {
W249 := __e.Get(1)
_ = W249
tmp17855 := MakeNative(func(__e *ControlFlow) {
W250 := __e.Get(1)
_ = W250
tmp17862 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W250)


if True == tmp17862 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17856 := MakeNative(func(__e *ControlFlow) {
W251 := __e.Get(1)
_ = W251
tmp17857 := MakeNative(func(__e *ControlFlow) {
W252 := __e.Get(1)
_ = W252
tmp17858 := PrimCons(W248, W251)

__e.TailApply(PrimFunc(symshen_4comb), W252, tmp17858)
return


}, 1)

tmp17859 := Call(__e, PrimFunc(symshen_4in_1_6), W250)


__e.TailApply(tmp17857, tmp17859)
return


}, 1)

tmp17860 := Call(__e, PrimFunc(symshen_4_5_1out), W250)


__e.TailApply(tmp17856, tmp17860)
return


}


}, 1)

tmp17863 := Call(__e, PrimFunc(symshen_4_5syntax_6), W249)


__e.TailApply(tmp17855, tmp17863)
return


}, 1)

tmp17864 := Call(__e, PrimFunc(symshen_4in_1_6), W247)


__e.TailApply(tmp17854, tmp17864)
return


}, 1)

tmp17865 := Call(__e, PrimFunc(symshen_4_5_1out), W247)


__e.TailApply(tmp17853, tmp17865)
return


}


}, 1)

tmp17868 := Call(__e, PrimFunc(symshen_4_5syntax_1item_6), V245)


tmp17869 := Call(__e, tmp17852, tmp17868)


__e.TailApply(tmp17836, tmp17869)
return


}, 1)

tmp17870 := Call(__e, ns2_1set, symshen_4_5syntax_6, tmp17835)


_ = tmp17870

tmp17871 := MakeNative(func(__e *ControlFlow) {
V257 := __e.Get(1)
_ = V257
tmp17872 := MakeNative(func(__e *ControlFlow) {
W258 := __e.Get(1)
_ = W258
tmp17874 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W258)


if True == tmp17874 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W258)
return
}


}, 1)

tmp17884 := PrimIsPair(V257)

var ifres17875 Obj

if True == tmp17884 {
tmp17876 := MakeNative(func(__e *ControlFlow) {
W259 := __e.Get(1)
_ = W259
tmp17877 := MakeNative(func(__e *ControlFlow) {
W260 := __e.Get(1)
_ = W260
tmp17879 := Call(__e, PrimFunc(symshen_4syntax_1item_2), W259)


if True == tmp17879 {
__e.TailApply(PrimFunc(symshen_4comb), W260, W259)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp17880 := Call(__e, PrimFunc(symtail), V257)


__e.TailApply(tmp17877, tmp17880)
return


}, 1)

tmp17881 := Call(__e, PrimFunc(symhead), V257)


tmp17882 := Call(__e, tmp17876, tmp17881)


ifres17875 = tmp17882


} else {
tmp17883 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17875 = tmp17883


}

__e.TailApply(tmp17872, ifres17875)
return


}, 1)

tmp17885 := Call(__e, ns2_1set, symshen_4_5syntax_1item_6, tmp17871)


_ = tmp17885

tmp17886 := MakeNative(func(__e *ControlFlow) {
V263 := __e.Get(1)
_ = V263
tmp17922 := Call(__e, PrimFunc(symshen_4colon_1equal_2), V263)


if True == tmp17922 {
__e.Return(False)
return
} else {
tmp17920 := Call(__e, PrimFunc(symshen_4semicolon_2), V263)


if True == tmp17920 {
__e.Return(False)
return
} else {
tmp17918 := Call(__e, PrimFunc(symatom_2), V263)


if True == tmp17918 {
__e.Return(True)
return
} else {
tmp17916 := PrimIsPair(V263)

var ifres17897 Obj

if True == tmp17916 {
tmp17914 := PrimHead(V263)

tmp17915 := PrimEqual(symcons, tmp17914)

var ifres17899 Obj

if True == tmp17915 {
tmp17912 := PrimTail(V263)

tmp17913 := PrimIsPair(tmp17912)

var ifres17901 Obj

if True == tmp17913 {
tmp17909 := PrimTail(V263)

tmp17910 := PrimTail(tmp17909)

tmp17911 := PrimIsPair(tmp17910)

var ifres17903 Obj

if True == tmp17911 {
tmp17905 := PrimTail(V263)

tmp17906 := PrimTail(tmp17905)

tmp17907 := PrimTail(tmp17906)

tmp17908 := PrimEqual(Nil, tmp17907)

var ifres17904 Obj

if True == tmp17908 {
ifres17904 = True


} else {
ifres17904 = False


}

ifres17903 = ifres17904


} else {
ifres17903 = False


}

var ifres17902 Obj

if True == ifres17903 {
ifres17902 = True


} else {
ifres17902 = False


}

ifres17901 = ifres17902


} else {
ifres17901 = False


}

var ifres17900 Obj

if True == ifres17901 {
ifres17900 = True


} else {
ifres17900 = False


}

ifres17899 = ifres17900


} else {
ifres17899 = False


}

var ifres17898 Obj

if True == ifres17899 {
ifres17898 = True


} else {
ifres17898 = False


}

ifres17897 = ifres17898


} else {
ifres17897 = False


}

if True == ifres17897 {
tmp17893 := PrimTail(V263)

tmp17894 := PrimHead(tmp17893)

tmp17895 := Call(__e, PrimFunc(symshen_4syntax_1item_2), tmp17894)


if True == tmp17895 {
tmp17888 := PrimTail(V263)

tmp17889 := PrimTail(tmp17888)

tmp17890 := PrimHead(tmp17889)

tmp17891 := Call(__e, PrimFunc(symshen_4syntax_1item_2), tmp17890)


if True == tmp17891 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


} else {
__e.Return(False)
return
}


} else {
__e.Return(False)
return
}


}


}


}


}, 1)

tmp17923 := Call(__e, ns2_1set, symshen_4syntax_1item_2, tmp17886)


_ = tmp17923

tmp17924 := MakeNative(func(__e *ControlFlow) {
V264 := __e.Get(1)
_ = V264
tmp17925 := MakeNative(func(__e *ControlFlow) {
W265 := __e.Get(1)
_ = W265
tmp17946 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W265)


if True == tmp17946 {
tmp17926 := MakeNative(func(__e *ControlFlow) {
W273 := __e.Get(1)
_ = W273
tmp17928 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W273)


if True == tmp17928 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W273)
return
}


}, 1)

tmp17929 := MakeNative(func(__e *ControlFlow) {
W274 := __e.Get(1)
_ = W274
tmp17942 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W274)


if True == tmp17942 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17930 := MakeNative(func(__e *ControlFlow) {
W275 := __e.Get(1)
_ = W275
tmp17939 := PrimIsPair(W275)

if True == tmp17939 {
tmp17931 := MakeNative(func(__e *ControlFlow) {
W276 := __e.Get(1)
_ = W276
tmp17932 := MakeNative(func(__e *ControlFlow) {
W277 := __e.Get(1)
_ = W277
tmp17934 := Call(__e, PrimFunc(symshen_4semicolon_2), W276)


tmp17935 := PrimNot(tmp17934)

if True == tmp17935 {
__e.TailApply(PrimFunc(symshen_4comb), W277, W276)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp17936 := Call(__e, PrimFunc(symtail), W275)


__e.TailApply(tmp17932, tmp17936)
return


}, 1)

tmp17937 := Call(__e, PrimFunc(symhead), W275)


__e.TailApply(tmp17931, tmp17937)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp17940 := Call(__e, PrimFunc(symshen_4in_1_6), W274)


__e.TailApply(tmp17930, tmp17940)
return


}


}, 1)

tmp17943 := Call(__e, PrimFunc(symshen_4_5colon_1equal_6), V264)


tmp17944 := Call(__e, tmp17929, tmp17943)


__e.TailApply(tmp17926, tmp17944)
return


} else {
__e.Return(W265)
return
}


}, 1)

tmp17947 := MakeNative(func(__e *ControlFlow) {
W266 := __e.Get(1)
_ = W266
tmp17973 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W266)


if True == tmp17973 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp17948 := MakeNative(func(__e *ControlFlow) {
W267 := __e.Get(1)
_ = W267
tmp17970 := PrimIsPair(W267)

if True == tmp17970 {
tmp17949 := MakeNative(func(__e *ControlFlow) {
W268 := __e.Get(1)
_ = W268
tmp17950 := MakeNative(func(__e *ControlFlow) {
W269 := __e.Get(1)
_ = W269
tmp17966 := Call(__e, PrimFunc(symshen_4hds_a_2), W269, symwhere)


if True == tmp17966 {
tmp17951 := MakeNative(func(__e *ControlFlow) {
W270 := __e.Get(1)
_ = W270
tmp17963 := PrimIsPair(W270)

if True == tmp17963 {
tmp17952 := MakeNative(func(__e *ControlFlow) {
W271 := __e.Get(1)
_ = W271
tmp17953 := MakeNative(func(__e *ControlFlow) {
W272 := __e.Get(1)
_ = W272
tmp17958 := Call(__e, PrimFunc(symshen_4semicolon_2), W268)


tmp17959 := PrimNot(tmp17958)

if True == tmp17959 {
tmp17954 := PrimCons(W268, Nil)

tmp17955 := PrimCons(W271, tmp17954)

tmp17956 := PrimCons(symwhere, tmp17955)

__e.TailApply(PrimFunc(symshen_4comb), W272, tmp17956)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp17960 := Call(__e, PrimFunc(symtail), W270)


__e.TailApply(tmp17953, tmp17960)
return


}, 1)

tmp17961 := Call(__e, PrimFunc(symhead), W270)


__e.TailApply(tmp17952, tmp17961)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp17964 := Call(__e, PrimFunc(symtail), W269)


__e.TailApply(tmp17951, tmp17964)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp17967 := Call(__e, PrimFunc(symtail), W267)


__e.TailApply(tmp17950, tmp17967)
return


}, 1)

tmp17968 := Call(__e, PrimFunc(symhead), W267)


__e.TailApply(tmp17949, tmp17968)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp17971 := Call(__e, PrimFunc(symshen_4in_1_6), W266)


__e.TailApply(tmp17948, tmp17971)
return


}


}, 1)

tmp17974 := Call(__e, PrimFunc(symshen_4_5colon_1equal_6), V264)


tmp17975 := Call(__e, tmp17947, tmp17974)


__e.TailApply(tmp17925, tmp17975)
return


}, 1)

tmp17976 := Call(__e, ns2_1set, symshen_4_5semantics_6, tmp17924)


_ = tmp17976

tmp17977 := MakeNative(func(__e *ControlFlow) {
V286 := __e.Get(1)
_ = V286
V287 := __e.Get(2)
_ = V287
V288 := __e.Get(3)
_ = V288
tmp17985 := PrimEqual(Nil, V288)

if True == tmp17985 {
__e.Return(PrimCons(symshen_4parse_1failure, Nil))
return
} else {
tmp17983 := PrimIsPair(V288)

if True == tmp17983 {
tmp17978 := PrimHead(V288)

tmp17979 := Call(__e, PrimFunc(symshen_4c_1rule_1_6shen), V286, tmp17978, V287)


tmp17980 := PrimTail(V288)

tmp17981 := Call(__e, PrimFunc(symshen_4c_1rules_1_6shen), V286, V287, tmp17980)


__e.TailApply(PrimFunc(symshen_4combine_1c_1code), tmp17979, tmp17981)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.c-rules->shen\n")))
return
}


}


}, 3)

tmp17986 := Call(__e, ns2_1set, symshen_4c_1rules_1_6shen, tmp17977)


_ = tmp17986

tmp17987 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symfail))
return
}, 0)

tmp17988 := Call(__e, ns2_1set, symshen_4parse_1failure, tmp17987)


_ = tmp17988

tmp17989 := MakeNative(func(__e *ControlFlow) {
V289 := __e.Get(1)
_ = V289
V290 := __e.Get(2)
_ = V290
tmp17990 := PrimCons(symResult, Nil)

tmp17991 := PrimCons(symshen_4parse_1failure_2, tmp17990)

tmp17992 := PrimCons(symResult, Nil)

tmp17993 := PrimCons(V290, tmp17992)

tmp17994 := PrimCons(tmp17991, tmp17993)

tmp17995 := PrimCons(symif, tmp17994)

tmp17996 := PrimCons(tmp17995, Nil)

tmp17997 := PrimCons(V289, tmp17996)

tmp17998 := PrimCons(symResult, tmp17997)

__e.Return(PrimCons(symlet, tmp17998))
return


}, 2)

tmp17999 := Call(__e, ns2_1set, symshen_4combine_1c_1code, tmp17989)


_ = tmp17999

tmp18000 := MakeNative(func(__e *ControlFlow) {
V297 := __e.Get(1)
_ = V297
V298 := __e.Get(2)
_ = V298
V299 := __e.Get(3)
_ = V299
tmp18014 := PrimIsPair(V298)

var ifres18005 Obj

if True == tmp18014 {
tmp18012 := PrimTail(V298)

tmp18013 := PrimIsPair(tmp18012)

var ifres18007 Obj

if True == tmp18013 {
tmp18009 := PrimTail(V298)

tmp18010 := PrimTail(tmp18009)

tmp18011 := PrimEqual(Nil, tmp18010)

var ifres18008 Obj

if True == tmp18011 {
ifres18008 = True


} else {
ifres18008 = False


}

ifres18007 = ifres18008


} else {
ifres18007 = False


}

var ifres18006 Obj

if True == ifres18007 {
ifres18006 = True


} else {
ifres18006 = False


}

ifres18005 = ifres18006


} else {
ifres18005 = False


}

if True == ifres18005 {
tmp18001 := PrimHead(V298)

tmp18002 := PrimTail(V298)

tmp18003 := PrimHead(tmp18002)

__e.TailApply(PrimFunc(symshen_4yacc_1syntax), V297, V299, tmp18001, tmp18003)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.c-rule->shen\n")))
return
}


}, 3)

tmp18015 := Call(__e, ns2_1set, symshen_4c_1rule_1_6shen, tmp18000)


_ = tmp18015

tmp18016 := MakeNative(func(__e *ControlFlow) {
V308 := __e.Get(1)
_ = V308
V309 := __e.Get(2)
_ = V309
V310 := __e.Get(3)
_ = V310
V311 := __e.Get(4)
_ = V311
tmp18080 := PrimEqual(Nil, V310)

var ifres18058 Obj

if True == tmp18080 {
tmp18079 := PrimIsPair(V311)

var ifres18060 Obj

if True == tmp18079 {
tmp18077 := PrimHead(V311)

tmp18078 := PrimEqual(symwhere, tmp18077)

var ifres18062 Obj

if True == tmp18078 {
tmp18075 := PrimTail(V311)

tmp18076 := PrimIsPair(tmp18075)

var ifres18064 Obj

if True == tmp18076 {
tmp18072 := PrimTail(V311)

tmp18073 := PrimTail(tmp18072)

tmp18074 := PrimIsPair(tmp18073)

var ifres18066 Obj

if True == tmp18074 {
tmp18068 := PrimTail(V311)

tmp18069 := PrimTail(tmp18068)

tmp18070 := PrimTail(tmp18069)

tmp18071 := PrimEqual(Nil, tmp18070)

var ifres18067 Obj

if True == tmp18071 {
ifres18067 = True


} else {
ifres18067 = False


}

ifres18066 = ifres18067


} else {
ifres18066 = False


}

var ifres18065 Obj

if True == ifres18066 {
ifres18065 = True


} else {
ifres18065 = False


}

ifres18064 = ifres18065


} else {
ifres18064 = False


}

var ifres18063 Obj

if True == ifres18064 {
ifres18063 = True


} else {
ifres18063 = False


}

ifres18062 = ifres18063


} else {
ifres18062 = False


}

var ifres18061 Obj

if True == ifres18062 {
ifres18061 = True


} else {
ifres18061 = False


}

ifres18060 = ifres18061


} else {
ifres18060 = False


}

var ifres18059 Obj

if True == ifres18060 {
ifres18059 = True


} else {
ifres18059 = False


}

ifres18058 = ifres18059


} else {
ifres18058 = False


}

if True == ifres18058 {
tmp18017 := PrimTail(V311)

tmp18018 := PrimHead(tmp18017)

tmp18019 := Call(__e, PrimFunc(symshen_4process_1yacc_1semantics), tmp18018)


tmp18020 := PrimTail(V311)

tmp18021 := PrimTail(tmp18020)

tmp18022 := PrimHead(tmp18021)

tmp18023 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V308, V309, Nil, tmp18022)


tmp18024 := PrimCons(symshen_4parse_1failure, Nil)

tmp18025 := PrimCons(tmp18024, Nil)

tmp18026 := PrimCons(tmp18023, tmp18025)

tmp18027 := PrimCons(tmp18019, tmp18026)

__e.Return(PrimCons(symif, tmp18027))
return


} else {
tmp18056 := PrimEqual(Nil, V310)

if True == tmp18056 {
__e.TailApply(PrimFunc(symshen_4yacc_1semantics), V308, V309, V311)
return
} else {
tmp18054 := PrimIsPair(V310)

if True == tmp18054 {
tmp18051 := PrimHead(V310)

tmp18052 := Call(__e, PrimFunc(symshen_4non_1terminal_2), tmp18051)


if True == tmp18052 {
tmp18028 := PrimHead(V310)

tmp18029 := PrimTail(V310)

__e.TailApply(PrimFunc(symshen_4non_1terminalcode), V308, V309, tmp18028, tmp18029, V311)
return


} else {
tmp18048 := PrimHead(V310)

tmp18049 := PrimIsVariable(tmp18048)

if True == tmp18049 {
tmp18030 := PrimHead(V310)

tmp18031 := PrimTail(V310)

__e.TailApply(PrimFunc(symshen_4variablecode), V308, V309, tmp18030, tmp18031, V311)
return


} else {
tmp18045 := PrimHead(V310)

tmp18046 := PrimEqual(sym__, tmp18045)

if True == tmp18046 {
tmp18032 := PrimHead(V310)

tmp18033 := PrimTail(V310)

__e.TailApply(PrimFunc(symshen_4wildcardcode), V308, V309, tmp18032, tmp18033, V311)
return


} else {
tmp18042 := PrimHead(V310)

tmp18043 := Call(__e, PrimFunc(symatom_2), tmp18042)


if True == tmp18043 {
tmp18034 := PrimHead(V310)

tmp18035 := PrimTail(V310)

__e.TailApply(PrimFunc(symshen_4terminalcode), V308, V309, tmp18034, tmp18035, V311)
return


} else {
tmp18039 := PrimHead(V310)

tmp18040 := PrimIsPair(tmp18039)

if True == tmp18040 {
tmp18036 := PrimHead(V310)

tmp18037 := PrimTail(V310)

__e.TailApply(PrimFunc(symshen_4conscode), V308, V309, tmp18036, tmp18037, V311)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.yacc-syntax\n")))
return
}


}


}


}


}


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.yacc-syntax\n")))
return
}


}


}


}, 4)

tmp18081 := Call(__e, ns2_1set, symshen_4yacc_1syntax, tmp18016)


_ = tmp18081

tmp18082 := MakeNative(func(__e *ControlFlow) {
V312 := __e.Get(1)
_ = V312
V313 := __e.Get(2)
_ = V313
V314 := __e.Get(3)
_ = V314
V315 := __e.Get(4)
_ = V315
V316 := __e.Get(5)
_ = V316
tmp18083 := MakeNative(func(__e *ControlFlow) {
W317 := __e.Get(1)
_ = W317
tmp18084 := MakeNative(func(__e *ControlFlow) {
W318 := __e.Get(1)
_ = W318
tmp18085 := MakeNative(func(__e *ControlFlow) {
W319 := __e.Get(1)
_ = W319
tmp18086 := PrimCons(V313, Nil)

tmp18087 := PrimCons(V314, tmp18086)

tmp18088 := PrimCons(W317, Nil)

tmp18089 := PrimCons(symshen_4parse_1failure_2, tmp18088)

tmp18090 := PrimCons(symshen_4parse_1failure, Nil)

tmp18091 := MakeNative(func(__e *ControlFlow) {
W320 := __e.Get(1)
_ = W320
tmp18101 := Call(__e, PrimFunc(symshen_4occurs_1check_2), V314, V316)


var ifres18098 Obj

if True == tmp18101 {
ifres18098 = True


} else {
tmp18100 := Call(__e, PrimFunc(symshen_4occurs_1check_2), W318, V316)


var ifres18099 Obj

if True == tmp18100 {
ifres18099 = True


} else {
ifres18099 = False


}

ifres18098 = ifres18099


}

if True == ifres18098 {
tmp18092 := PrimCons(W317, Nil)

tmp18093 := PrimCons(symshen_4_5_1out, tmp18092)

tmp18094 := PrimCons(W320, Nil)

tmp18095 := PrimCons(tmp18093, tmp18094)

tmp18096 := PrimCons(W318, tmp18095)

__e.Return(PrimCons(symlet, tmp18096))
return


} else {
__e.Return(W320)
return
}


}, 1)

tmp18102 := PrimCons(W317, Nil)

tmp18103 := PrimCons(symshen_4in_1_6, tmp18102)

tmp18104 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V312, W319, V315, V316)


tmp18105 := PrimCons(tmp18104, Nil)

tmp18106 := PrimCons(tmp18103, tmp18105)

tmp18107 := PrimCons(W319, tmp18106)

tmp18108 := PrimCons(symlet, tmp18107)

tmp18109 := Call(__e, tmp18091, tmp18108)


tmp18110 := PrimCons(tmp18109, Nil)

tmp18111 := PrimCons(tmp18090, tmp18110)

tmp18112 := PrimCons(tmp18089, tmp18111)

tmp18113 := PrimCons(symif, tmp18112)

tmp18114 := PrimCons(tmp18113, Nil)

tmp18115 := PrimCons(tmp18087, tmp18114)

tmp18116 := PrimCons(W317, tmp18115)

__e.Return(PrimCons(symlet, tmp18116))
return


}, 1)

tmp18117 := Call(__e, PrimFunc(symconcat), symRemainder, V314)


__e.TailApply(tmp18085, tmp18117)
return


}, 1)

tmp18118 := Call(__e, PrimFunc(symconcat), symAction, V314)


__e.TailApply(tmp18084, tmp18118)
return


}, 1)

tmp18119 := Call(__e, PrimFunc(symconcat), symParse, V314)


__e.TailApply(tmp18083, tmp18119)
return


}, 5)

tmp18120 := Call(__e, ns2_1set, symshen_4non_1terminalcode, tmp18082)


_ = tmp18120

tmp18121 := MakeNative(func(__e *ControlFlow) {
V321 := __e.Get(1)
_ = V321
V322 := __e.Get(2)
_ = V322
V323 := __e.Get(3)
_ = V323
V324 := __e.Get(4)
_ = V324
V325 := __e.Get(5)
_ = V325
tmp18122 := MakeNative(func(__e *ControlFlow) {
W326 := __e.Get(1)
_ = W326
tmp18123 := PrimCons(V322, Nil)

tmp18124 := PrimCons(symcons_2, tmp18123)

tmp18125 := MakeNative(func(__e *ControlFlow) {
W327 := __e.Get(1)
_ = W327
tmp18132 := Call(__e, PrimFunc(symshen_4occurs_1check_2), V323, V325)


if True == tmp18132 {
tmp18126 := PrimCons(V322, Nil)

tmp18127 := PrimCons(symhead, tmp18126)

tmp18128 := PrimCons(W327, Nil)

tmp18129 := PrimCons(tmp18127, tmp18128)

tmp18130 := PrimCons(V323, tmp18129)

__e.Return(PrimCons(symlet, tmp18130))
return


} else {
__e.Return(W327)
return
}


}, 1)

tmp18133 := PrimCons(V322, Nil)

tmp18134 := PrimCons(symtail, tmp18133)

tmp18135 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V321, W326, V324, V325)


tmp18136 := PrimCons(tmp18135, Nil)

tmp18137 := PrimCons(tmp18134, tmp18136)

tmp18138 := PrimCons(W326, tmp18137)

tmp18139 := PrimCons(symlet, tmp18138)

tmp18140 := Call(__e, tmp18125, tmp18139)


tmp18141 := PrimCons(symshen_4parse_1failure, Nil)

tmp18142 := PrimCons(tmp18141, Nil)

tmp18143 := PrimCons(tmp18140, tmp18142)

tmp18144 := PrimCons(tmp18124, tmp18143)

__e.Return(PrimCons(symif, tmp18144))
return


}, 1)

tmp18145 := Call(__e, PrimFunc(symgensym), symRemainder)


__e.TailApply(tmp18122, tmp18145)
return


}, 5)

tmp18146 := Call(__e, ns2_1set, symshen_4variablecode, tmp18121)


_ = tmp18146

tmp18147 := MakeNative(func(__e *ControlFlow) {
V328 := __e.Get(1)
_ = V328
V329 := __e.Get(2)
_ = V329
V330 := __e.Get(3)
_ = V330
V331 := __e.Get(4)
_ = V331
V332 := __e.Get(5)
_ = V332
tmp18148 := MakeNative(func(__e *ControlFlow) {
W333 := __e.Get(1)
_ = W333
tmp18149 := PrimCons(V329, Nil)

tmp18150 := PrimCons(symcons_2, tmp18149)

tmp18151 := PrimCons(V329, Nil)

tmp18152 := PrimCons(symtail, tmp18151)

tmp18153 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V328, W333, V331, V332)


tmp18154 := PrimCons(tmp18153, Nil)

tmp18155 := PrimCons(tmp18152, tmp18154)

tmp18156 := PrimCons(W333, tmp18155)

tmp18157 := PrimCons(symlet, tmp18156)

tmp18158 := PrimCons(symshen_4parse_1failure, Nil)

tmp18159 := PrimCons(tmp18158, Nil)

tmp18160 := PrimCons(tmp18157, tmp18159)

tmp18161 := PrimCons(tmp18150, tmp18160)

__e.Return(PrimCons(symif, tmp18161))
return


}, 1)

tmp18162 := Call(__e, PrimFunc(symgensym), symRemainder)


__e.TailApply(tmp18148, tmp18162)
return


}, 5)

tmp18163 := Call(__e, ns2_1set, symshen_4wildcardcode, tmp18147)


_ = tmp18163

tmp18164 := MakeNative(func(__e *ControlFlow) {
V334 := __e.Get(1)
_ = V334
V335 := __e.Get(2)
_ = V335
V336 := __e.Get(3)
_ = V336
V337 := __e.Get(4)
_ = V337
V338 := __e.Get(5)
_ = V338
tmp18165 := MakeNative(func(__e *ControlFlow) {
W339 := __e.Get(1)
_ = W339
tmp18166 := PrimCons(V336, Nil)

tmp18167 := PrimCons(V335, tmp18166)

tmp18168 := PrimCons(symshen_4hds_a_2, tmp18167)

tmp18169 := PrimCons(V335, Nil)

tmp18170 := PrimCons(symtail, tmp18169)

tmp18171 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V334, W339, V337, V338)


tmp18172 := PrimCons(tmp18171, Nil)

tmp18173 := PrimCons(tmp18170, tmp18172)

tmp18174 := PrimCons(W339, tmp18173)

tmp18175 := PrimCons(symlet, tmp18174)

tmp18176 := PrimCons(symshen_4parse_1failure, Nil)

tmp18177 := PrimCons(tmp18176, Nil)

tmp18178 := PrimCons(tmp18175, tmp18177)

tmp18179 := PrimCons(tmp18168, tmp18178)

__e.Return(PrimCons(symif, tmp18179))
return


}, 1)

tmp18180 := Call(__e, PrimFunc(symgensym), symRemainder)


__e.TailApply(tmp18165, tmp18180)
return


}, 5)

tmp18181 := Call(__e, ns2_1set, symshen_4terminalcode, tmp18164)


_ = tmp18181

tmp18182 := MakeNative(func(__e *ControlFlow) {
V347 := __e.Get(1)
_ = V347
V348 := __e.Get(2)
_ = V348
tmp18188 := PrimIsPair(V347)

var ifres18184 Obj

if True == tmp18188 {
tmp18186 := PrimHead(V347)

tmp18187 := PrimEqual(tmp18186, V348)

var ifres18185 Obj

if True == tmp18187 {
ifres18185 = True


} else {
ifres18185 = False


}

ifres18184 = ifres18185


} else {
ifres18184 = False


}

if True == ifres18184 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 2)

tmp18189 := Call(__e, ns2_1set, symshen_4hds_a_2, tmp18182)


_ = tmp18189

tmp18190 := MakeNative(func(__e *ControlFlow) {
V349 := __e.Get(1)
_ = V349
V350 := __e.Get(2)
_ = V350
V351 := __e.Get(3)
_ = V351
V352 := __e.Get(4)
_ = V352
V353 := __e.Get(5)
_ = V353
tmp18191 := MakeNative(func(__e *ControlFlow) {
W354 := __e.Get(1)
_ = W354
tmp18192 := MakeNative(func(__e *ControlFlow) {
W355 := __e.Get(1)
_ = W355
tmp18193 := MakeNative(func(__e *ControlFlow) {
W356 := __e.Get(1)
_ = W356
tmp18194 := PrimCons(V350, Nil)

tmp18195 := PrimCons(symshen_4ccons_2, tmp18194)

tmp18196 := PrimCons(V350, Nil)

tmp18197 := PrimCons(symhead, tmp18196)

tmp18198 := PrimCons(V350, Nil)

tmp18199 := PrimCons(symtail, tmp18198)

tmp18200 := Call(__e, PrimFunc(symshen_4decons), V351)


tmp18201 := PrimCons(sym_5end_6, Nil)

tmp18202 := Call(__e, PrimFunc(symappend), tmp18200, tmp18201)


tmp18203 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V349, W356, V352, V353)


tmp18204 := PrimCons(tmp18203, Nil)

tmp18205 := PrimCons(symshen_4processed, tmp18204)

tmp18206 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V349, W355, tmp18202, tmp18205)


tmp18207 := PrimCons(tmp18206, Nil)

tmp18208 := PrimCons(tmp18199, tmp18207)

tmp18209 := PrimCons(W356, tmp18208)

tmp18210 := PrimCons(tmp18197, tmp18209)

tmp18211 := PrimCons(W355, tmp18210)

tmp18212 := PrimCons(symlet, tmp18211)

tmp18213 := PrimCons(symshen_4parse_1failure, Nil)

tmp18214 := PrimCons(tmp18213, Nil)

tmp18215 := PrimCons(tmp18212, tmp18214)

tmp18216 := PrimCons(tmp18195, tmp18215)

__e.Return(PrimCons(symif, tmp18216))
return


}, 1)

tmp18217 := Call(__e, PrimFunc(symgensym), symTl)


__e.TailApply(tmp18193, tmp18217)
return


}, 1)

tmp18218 := Call(__e, PrimFunc(symgensym), symHd)


__e.TailApply(tmp18192, tmp18218)
return


}, 1)

tmp18219 := Call(__e, PrimFunc(symgensym), symRemainder)


__e.TailApply(tmp18191, tmp18219)
return


}, 5)

tmp18220 := Call(__e, ns2_1set, symshen_4conscode, tmp18190)


_ = tmp18220

tmp18221 := MakeNative(func(__e *ControlFlow) {
V367 := __e.Get(1)
_ = V367
tmp18233 := PrimIsPair(V367)

var ifres18229 Obj

if True == tmp18233 {
tmp18231 := PrimHead(V367)

tmp18232 := PrimIsPair(tmp18231)

var ifres18230 Obj

if True == tmp18232 {
ifres18230 = True


} else {
ifres18230 = False


}

ifres18229 = ifres18230


} else {
ifres18229 = False


}

if True == ifres18229 {
__e.Return(True)
return
} else {
tmp18227 := PrimIsPair(V367)

var ifres18223 Obj

if True == tmp18227 {
tmp18225 := PrimHead(V367)

tmp18226 := PrimEqual(Nil, tmp18225)

var ifres18224 Obj

if True == tmp18226 {
ifres18224 = True


} else {
ifres18224 = False


}

ifres18223 = ifres18224


} else {
ifres18223 = False


}

if True == ifres18223 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp18234 := Call(__e, ns2_1set, symshen_4ccons_2, tmp18221)


_ = tmp18234

tmp18235 := MakeNative(func(__e *ControlFlow) {
V368 := __e.Get(1)
_ = V368
tmp18262 := PrimIsPair(V368)

var ifres18243 Obj

if True == tmp18262 {
tmp18260 := PrimHead(V368)

tmp18261 := PrimEqual(symcons, tmp18260)

var ifres18245 Obj

if True == tmp18261 {
tmp18258 := PrimTail(V368)

tmp18259 := PrimIsPair(tmp18258)

var ifres18247 Obj

if True == tmp18259 {
tmp18255 := PrimTail(V368)

tmp18256 := PrimTail(tmp18255)

tmp18257 := PrimIsPair(tmp18256)

var ifres18249 Obj

if True == tmp18257 {
tmp18251 := PrimTail(V368)

tmp18252 := PrimTail(tmp18251)

tmp18253 := PrimTail(tmp18252)

tmp18254 := PrimEqual(Nil, tmp18253)

var ifres18250 Obj

if True == tmp18254 {
ifres18250 = True


} else {
ifres18250 = False


}

ifres18249 = ifres18250


} else {
ifres18249 = False


}

var ifres18248 Obj

if True == ifres18249 {
ifres18248 = True


} else {
ifres18248 = False


}

ifres18247 = ifres18248


} else {
ifres18247 = False


}

var ifres18246 Obj

if True == ifres18247 {
ifres18246 = True


} else {
ifres18246 = False


}

ifres18245 = ifres18246


} else {
ifres18245 = False


}

var ifres18244 Obj

if True == ifres18245 {
ifres18244 = True


} else {
ifres18244 = False


}

ifres18243 = ifres18244


} else {
ifres18243 = False


}

if True == ifres18243 {
tmp18236 := PrimTail(V368)

tmp18237 := PrimHead(tmp18236)

tmp18238 := PrimTail(V368)

tmp18239 := PrimTail(tmp18238)

tmp18240 := PrimHead(tmp18239)

tmp18241 := Call(__e, PrimFunc(symshen_4decons), tmp18240)


__e.Return(PrimCons(tmp18237, tmp18241))
return


} else {
__e.Return(V368)
return
}


}, 1)

tmp18263 := Call(__e, ns2_1set, symshen_4decons, tmp18235)


_ = tmp18263

tmp18264 := MakeNative(func(__e *ControlFlow) {
V369 := __e.Get(1)
_ = V369
V370 := __e.Get(2)
_ = V370
tmp18265 := PrimCons(V370, Nil)

__e.Return(PrimCons(V369, tmp18265))
return


}, 2)

tmp18266 := Call(__e, ns2_1set, symshen_4comb, tmp18264)


_ = tmp18266

tmp18267 := MakeNative(func(__e *ControlFlow) {
V375 := __e.Get(1)
_ = V375
V376 := __e.Get(2)
_ = V376
V377 := __e.Get(3)
_ = V377
tmp18289 := PrimIsPair(V377)

var ifres18276 Obj

if True == tmp18289 {
tmp18287 := PrimHead(V377)

tmp18288 := PrimEqual(symshen_4processed, tmp18287)

var ifres18278 Obj

if True == tmp18288 {
tmp18285 := PrimTail(V377)

tmp18286 := PrimIsPair(tmp18285)

var ifres18280 Obj

if True == tmp18286 {
tmp18282 := PrimTail(V377)

tmp18283 := PrimTail(tmp18282)

tmp18284 := PrimEqual(Nil, tmp18283)

var ifres18281 Obj

if True == tmp18284 {
ifres18281 = True


} else {
ifres18281 = False


}

ifres18280 = ifres18281


} else {
ifres18280 = False


}

var ifres18279 Obj

if True == ifres18280 {
ifres18279 = True


} else {
ifres18279 = False


}

ifres18278 = ifres18279


} else {
ifres18278 = False


}

var ifres18277 Obj

if True == ifres18278 {
ifres18277 = True


} else {
ifres18277 = False


}

ifres18276 = ifres18277


} else {
ifres18276 = False


}

if True == ifres18276 {
tmp18268 := PrimTail(V377)

__e.Return(PrimHead(tmp18268))
return


} else {
tmp18269 := MakeNative(func(__e *ControlFlow) {
W378 := __e.Get(1)
_ = W378
tmp18270 := MakeNative(func(__e *ControlFlow) {
W379 := __e.Get(1)
_ = W379
tmp18271 := PrimCons(W379, Nil)

tmp18272 := PrimCons(V376, tmp18271)

__e.Return(PrimCons(symshen_4comb, tmp18272))
return


}, 1)

tmp18273 := Call(__e, PrimFunc(symshen_4use_1type_1info), V375, W378)


__e.TailApply(tmp18270, tmp18273)
return


}, 1)

tmp18274 := Call(__e, PrimFunc(symshen_4process_1yacc_1semantics), V377)


__e.TailApply(tmp18269, tmp18274)
return


}


}, 3)

tmp18290 := Call(__e, ns2_1set, symshen_4yacc_1semantics, tmp18267)


_ = tmp18290

tmp18291 := MakeNative(func(__e *ControlFlow) {
V383 := __e.Get(1)
_ = V383
V384 := __e.Get(2)
_ = V384
tmp18479 := PrimIsPair(V383)

var ifres18300 Obj

if True == tmp18479 {
tmp18477 := PrimHead(V383)

tmp18478 := PrimEqual(sym_i, tmp18477)

var ifres18302 Obj

if True == tmp18478 {
tmp18475 := PrimTail(V383)

tmp18476 := PrimIsPair(tmp18475)

var ifres18304 Obj

if True == tmp18476 {
tmp18472 := PrimTail(V383)

tmp18473 := PrimHead(tmp18472)

tmp18474 := PrimIsPair(tmp18473)

var ifres18306 Obj

if True == tmp18474 {
tmp18468 := PrimTail(V383)

tmp18469 := PrimHead(tmp18468)

tmp18470 := PrimHead(tmp18469)

tmp18471 := PrimEqual(symlist, tmp18470)

var ifres18308 Obj

if True == tmp18471 {
tmp18464 := PrimTail(V383)

tmp18465 := PrimHead(tmp18464)

tmp18466 := PrimTail(tmp18465)

tmp18467 := PrimIsPair(tmp18466)

var ifres18310 Obj

if True == tmp18467 {
tmp18459 := PrimTail(V383)

tmp18460 := PrimHead(tmp18459)

tmp18461 := PrimTail(tmp18460)

tmp18462 := PrimTail(tmp18461)

tmp18463 := PrimEqual(Nil, tmp18462)

var ifres18312 Obj

if True == tmp18463 {
tmp18456 := PrimTail(V383)

tmp18457 := PrimTail(tmp18456)

tmp18458 := PrimIsPair(tmp18457)

var ifres18314 Obj

if True == tmp18458 {
tmp18452 := PrimTail(V383)

tmp18453 := PrimTail(tmp18452)

tmp18454 := PrimHead(tmp18453)

tmp18455 := PrimEqual(sym_1_1_6, tmp18454)

var ifres18316 Obj

if True == tmp18455 {
tmp18448 := PrimTail(V383)

tmp18449 := PrimTail(tmp18448)

tmp18450 := PrimTail(tmp18449)

tmp18451 := PrimIsPair(tmp18450)

var ifres18318 Obj

if True == tmp18451 {
tmp18443 := PrimTail(V383)

tmp18444 := PrimTail(tmp18443)

tmp18445 := PrimTail(tmp18444)

tmp18446 := PrimHead(tmp18445)

tmp18447 := PrimIsPair(tmp18446)

var ifres18320 Obj

if True == tmp18447 {
tmp18437 := PrimTail(V383)

tmp18438 := PrimTail(tmp18437)

tmp18439 := PrimTail(tmp18438)

tmp18440 := PrimHead(tmp18439)

tmp18441 := PrimHead(tmp18440)

tmp18442 := PrimEqual(symstr, tmp18441)

var ifres18322 Obj

if True == tmp18442 {
tmp18431 := PrimTail(V383)

tmp18432 := PrimTail(tmp18431)

tmp18433 := PrimTail(tmp18432)

tmp18434 := PrimHead(tmp18433)

tmp18435 := PrimTail(tmp18434)

tmp18436 := PrimIsPair(tmp18435)

var ifres18324 Obj

if True == tmp18436 {
tmp18424 := PrimTail(V383)

tmp18425 := PrimTail(tmp18424)

tmp18426 := PrimTail(tmp18425)

tmp18427 := PrimHead(tmp18426)

tmp18428 := PrimTail(tmp18427)

tmp18429 := PrimHead(tmp18428)

tmp18430 := PrimIsPair(tmp18429)

var ifres18326 Obj

if True == tmp18430 {
tmp18416 := PrimTail(V383)

tmp18417 := PrimTail(tmp18416)

tmp18418 := PrimTail(tmp18417)

tmp18419 := PrimHead(tmp18418)

tmp18420 := PrimTail(tmp18419)

tmp18421 := PrimHead(tmp18420)

tmp18422 := PrimHead(tmp18421)

tmp18423 := PrimEqual(symlist, tmp18422)

var ifres18328 Obj

if True == tmp18423 {
tmp18408 := PrimTail(V383)

tmp18409 := PrimTail(tmp18408)

tmp18410 := PrimTail(tmp18409)

tmp18411 := PrimHead(tmp18410)

tmp18412 := PrimTail(tmp18411)

tmp18413 := PrimHead(tmp18412)

tmp18414 := PrimTail(tmp18413)

tmp18415 := PrimIsPair(tmp18414)

var ifres18330 Obj

if True == tmp18415 {
tmp18399 := PrimTail(V383)

tmp18400 := PrimTail(tmp18399)

tmp18401 := PrimTail(tmp18400)

tmp18402 := PrimHead(tmp18401)

tmp18403 := PrimTail(tmp18402)

tmp18404 := PrimHead(tmp18403)

tmp18405 := PrimTail(tmp18404)

tmp18406 := PrimTail(tmp18405)

tmp18407 := PrimEqual(Nil, tmp18406)

var ifres18332 Obj

if True == tmp18407 {
tmp18392 := PrimTail(V383)

tmp18393 := PrimTail(tmp18392)

tmp18394 := PrimTail(tmp18393)

tmp18395 := PrimHead(tmp18394)

tmp18396 := PrimTail(tmp18395)

tmp18397 := PrimTail(tmp18396)

tmp18398 := PrimIsPair(tmp18397)

var ifres18334 Obj

if True == tmp18398 {
tmp18384 := PrimTail(V383)

tmp18385 := PrimTail(tmp18384)

tmp18386 := PrimTail(tmp18385)

tmp18387 := PrimHead(tmp18386)

tmp18388 := PrimTail(tmp18387)

tmp18389 := PrimTail(tmp18388)

tmp18390 := PrimTail(tmp18389)

tmp18391 := PrimEqual(Nil, tmp18390)

var ifres18336 Obj

if True == tmp18391 {
tmp18379 := PrimTail(V383)

tmp18380 := PrimTail(tmp18379)

tmp18381 := PrimTail(tmp18380)

tmp18382 := PrimTail(tmp18381)

tmp18383 := PrimIsPair(tmp18382)

var ifres18338 Obj

if True == tmp18383 {
tmp18373 := PrimTail(V383)

tmp18374 := PrimTail(tmp18373)

tmp18375 := PrimTail(tmp18374)

tmp18376 := PrimTail(tmp18375)

tmp18377 := PrimHead(tmp18376)

tmp18378 := PrimEqual(sym_j, tmp18377)

var ifres18340 Obj

if True == tmp18378 {
tmp18367 := PrimTail(V383)

tmp18368 := PrimTail(tmp18367)

tmp18369 := PrimTail(tmp18368)

tmp18370 := PrimTail(tmp18369)

tmp18371 := PrimTail(tmp18370)

tmp18372 := PrimEqual(Nil, tmp18371)

var ifres18342 Obj

if True == tmp18372 {
tmp18354 := PrimTail(V383)

tmp18355 := PrimHead(tmp18354)

tmp18356 := PrimTail(tmp18355)

tmp18357 := PrimHead(tmp18356)

tmp18358 := PrimTail(V383)

tmp18359 := PrimTail(tmp18358)

tmp18360 := PrimTail(tmp18359)

tmp18361 := PrimHead(tmp18360)

tmp18362 := PrimTail(tmp18361)

tmp18363 := PrimHead(tmp18362)

tmp18364 := PrimTail(tmp18363)

tmp18365 := PrimHead(tmp18364)

tmp18366 := PrimEqual(tmp18357, tmp18365)

var ifres18344 Obj

if True == tmp18366 {
tmp18346 := PrimTail(V383)

tmp18347 := PrimTail(tmp18346)

tmp18348 := PrimTail(tmp18347)

tmp18349 := PrimHead(tmp18348)

tmp18350 := PrimTail(tmp18349)

tmp18351 := PrimTail(tmp18350)

tmp18352 := PrimHead(tmp18351)

tmp18353 := Call(__e, PrimFunc(symshen_4monomorphic_2), tmp18352)


var ifres18345 Obj

if True == tmp18353 {
ifres18345 = True


} else {
ifres18345 = False


}

ifres18344 = ifres18345


} else {
ifres18344 = False


}

var ifres18343 Obj

if True == ifres18344 {
ifres18343 = True


} else {
ifres18343 = False


}

ifres18342 = ifres18343


} else {
ifres18342 = False


}

var ifres18341 Obj

if True == ifres18342 {
ifres18341 = True


} else {
ifres18341 = False


}

ifres18340 = ifres18341


} else {
ifres18340 = False


}

var ifres18339 Obj

if True == ifres18340 {
ifres18339 = True


} else {
ifres18339 = False


}

ifres18338 = ifres18339


} else {
ifres18338 = False


}

var ifres18337 Obj

if True == ifres18338 {
ifres18337 = True


} else {
ifres18337 = False


}

ifres18336 = ifres18337


} else {
ifres18336 = False


}

var ifres18335 Obj

if True == ifres18336 {
ifres18335 = True


} else {
ifres18335 = False


}

ifres18334 = ifres18335


} else {
ifres18334 = False


}

var ifres18333 Obj

if True == ifres18334 {
ifres18333 = True


} else {
ifres18333 = False


}

ifres18332 = ifres18333


} else {
ifres18332 = False


}

var ifres18331 Obj

if True == ifres18332 {
ifres18331 = True


} else {
ifres18331 = False


}

ifres18330 = ifres18331


} else {
ifres18330 = False


}

var ifres18329 Obj

if True == ifres18330 {
ifres18329 = True


} else {
ifres18329 = False


}

ifres18328 = ifres18329


} else {
ifres18328 = False


}

var ifres18327 Obj

if True == ifres18328 {
ifres18327 = True


} else {
ifres18327 = False


}

ifres18326 = ifres18327


} else {
ifres18326 = False


}

var ifres18325 Obj

if True == ifres18326 {
ifres18325 = True


} else {
ifres18325 = False


}

ifres18324 = ifres18325


} else {
ifres18324 = False


}

var ifres18323 Obj

if True == ifres18324 {
ifres18323 = True


} else {
ifres18323 = False


}

ifres18322 = ifres18323


} else {
ifres18322 = False


}

var ifres18321 Obj

if True == ifres18322 {
ifres18321 = True


} else {
ifres18321 = False


}

ifres18320 = ifres18321


} else {
ifres18320 = False


}

var ifres18319 Obj

if True == ifres18320 {
ifres18319 = True


} else {
ifres18319 = False


}

ifres18318 = ifres18319


} else {
ifres18318 = False


}

var ifres18317 Obj

if True == ifres18318 {
ifres18317 = True


} else {
ifres18317 = False


}

ifres18316 = ifres18317


} else {
ifres18316 = False


}

var ifres18315 Obj

if True == ifres18316 {
ifres18315 = True


} else {
ifres18315 = False


}

ifres18314 = ifres18315


} else {
ifres18314 = False


}

var ifres18313 Obj

if True == ifres18314 {
ifres18313 = True


} else {
ifres18313 = False


}

ifres18312 = ifres18313


} else {
ifres18312 = False


}

var ifres18311 Obj

if True == ifres18312 {
ifres18311 = True


} else {
ifres18311 = False


}

ifres18310 = ifres18311


} else {
ifres18310 = False


}

var ifres18309 Obj

if True == ifres18310 {
ifres18309 = True


} else {
ifres18309 = False


}

ifres18308 = ifres18309


} else {
ifres18308 = False


}

var ifres18307 Obj

if True == ifres18308 {
ifres18307 = True


} else {
ifres18307 = False


}

ifres18306 = ifres18307


} else {
ifres18306 = False


}

var ifres18305 Obj

if True == ifres18306 {
ifres18305 = True


} else {
ifres18305 = False


}

ifres18304 = ifres18305


} else {
ifres18304 = False


}

var ifres18303 Obj

if True == ifres18304 {
ifres18303 = True


} else {
ifres18303 = False


}

ifres18302 = ifres18303


} else {
ifres18302 = False


}

var ifres18301 Obj

if True == ifres18302 {
ifres18301 = True


} else {
ifres18301 = False


}

ifres18300 = ifres18301


} else {
ifres18300 = False


}

if True == ifres18300 {
tmp18292 := PrimTail(V383)

tmp18293 := PrimTail(tmp18292)

tmp18294 := PrimTail(tmp18293)

tmp18295 := PrimHead(tmp18294)

tmp18296 := PrimTail(tmp18295)

tmp18297 := PrimTail(tmp18296)

tmp18298 := PrimCons(V384, tmp18297)

__e.Return(PrimCons(symtype, tmp18298))
return


} else {
__e.Return(V384)
return
}


}, 2)

tmp18480 := Call(__e, ns2_1set, symshen_4use_1type_1info, tmp18291)


_ = tmp18480

tmp18481 := MakeNative(func(__e *ControlFlow) {
V387 := __e.Get(1)
_ = V387
tmp18491 := PrimIsVariable(V387)

if True == tmp18491 {
__e.Return(False)
return
} else {
tmp18489 := PrimIsPair(V387)

if True == tmp18489 {
tmp18486 := PrimHead(V387)

tmp18487 := Call(__e, PrimFunc(symshen_4monomorphic_2), tmp18486)


if True == tmp18487 {
tmp18483 := PrimTail(V387)

tmp18484 := Call(__e, PrimFunc(symshen_4monomorphic_2), tmp18483)


if True == tmp18484 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


} else {
__e.Return(False)
return
}


} else {
__e.Return(True)
return
}


}


}, 1)

tmp18492 := Call(__e, ns2_1set, symshen_4monomorphic_2, tmp18481)


_ = tmp18492

tmp18493 := MakeNative(func(__e *ControlFlow) {
V388 := __e.Get(1)
_ = V388
tmp18519 := PrimIsPair(V388)

var ifres18501 Obj

if True == tmp18519 {
tmp18517 := PrimHead(V388)

tmp18518 := PrimEqual(symprotect, tmp18517)

var ifres18503 Obj

if True == tmp18518 {
tmp18515 := PrimTail(V388)

tmp18516 := PrimIsPair(tmp18515)

var ifres18505 Obj

if True == tmp18516 {
tmp18512 := PrimTail(V388)

tmp18513 := PrimTail(tmp18512)

tmp18514 := PrimEqual(Nil, tmp18513)

var ifres18507 Obj

if True == tmp18514 {
tmp18509 := PrimTail(V388)

tmp18510 := PrimHead(tmp18509)

tmp18511 := Call(__e, PrimFunc(symshen_4non_1terminal_2), tmp18510)


var ifres18508 Obj

if True == tmp18511 {
ifres18508 = True


} else {
ifres18508 = False


}

ifres18507 = ifres18508


} else {
ifres18507 = False


}

var ifres18506 Obj

if True == ifres18507 {
ifres18506 = True


} else {
ifres18506 = False


}

ifres18505 = ifres18506


} else {
ifres18505 = False


}

var ifres18504 Obj

if True == ifres18505 {
ifres18504 = True


} else {
ifres18504 = False


}

ifres18503 = ifres18504


} else {
ifres18503 = False


}

var ifres18502 Obj

if True == ifres18503 {
ifres18502 = True


} else {
ifres18502 = False


}

ifres18501 = ifres18502


} else {
ifres18501 = False


}

if True == ifres18501 {
tmp18494 := PrimTail(V388)

__e.Return(PrimHead(tmp18494))
return


} else {
tmp18499 := PrimIsPair(V388)

if True == tmp18499 {
tmp18495 := MakeNative(func(__e *ControlFlow) {
Z389 := __e.Get(1)
_ = Z389
__e.TailApply(PrimFunc(symshen_4process_1yacc_1semantics), Z389)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp18495, V388)
return


} else {
tmp18497 := Call(__e, PrimFunc(symshen_4non_1terminal_2), V388)


if True == tmp18497 {
__e.TailApply(PrimFunc(symconcat), symAction, V388)
return
} else {
__e.Return(V388)
return
}


}


}


}, 1)

tmp18520 := Call(__e, ns2_1set, symshen_4process_1yacc_1semantics, tmp18493)


_ = tmp18520

tmp18521 := MakeNative(func(__e *ControlFlow) {
V390 := __e.Get(1)
_ = V390
tmp18522 := PrimTail(V390)

__e.Return(PrimHead(tmp18522))
return


}, 1)

tmp18523 := Call(__e, ns2_1set, symshen_4_5_1out, tmp18521)


_ = tmp18523

tmp18524 := MakeNative(func(__e *ControlFlow) {
V391 := __e.Get(1)
_ = V391
__e.Return(PrimHead(V391))
return
}, 1)

tmp18525 := Call(__e, ns2_1set, symshen_4in_1_6, tmp18524)


_ = tmp18525

tmp18526 := MakeNative(func(__e *ControlFlow) {
V392 := __e.Get(1)
_ = V392
tmp18527 := PrimCons(V392, Nil)

__e.Return(PrimCons(Nil, tmp18527))
return


}, 1)

tmp18528 := Call(__e, ns2_1set, sym_5_b_6, tmp18526)


_ = tmp18528

tmp18529 := MakeNative(func(__e *ControlFlow) {
V393 := __e.Get(1)
_ = V393
tmp18530 := PrimCons(Nil, Nil)

__e.Return(PrimCons(V393, tmp18530))
return


}, 1)

tmp18531 := Call(__e, ns2_1set, sym_5e_6, tmp18529)


_ = tmp18531

tmp18532 := MakeNative(func(__e *ControlFlow) {
V396 := __e.Get(1)
_ = V396
tmp18535 := PrimEqual(Nil, V396)

if True == tmp18535 {
tmp18533 := PrimCons(Nil, Nil)

__e.Return(PrimCons(Nil, tmp18533))
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

__e.TailApply(ns2_1set, sym_5end_6, tmp18532)
return




}, 0)

