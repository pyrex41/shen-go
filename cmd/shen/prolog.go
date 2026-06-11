package main

import . "github.com/tiancaiamao/shen-go/kl"

var PrologMain = MakeNative(func(__e *ControlFlow) {
tmp8434 := MakeNative(func(__e *ControlFlow) {
V1577 := __e.Get(1)
_ = V1577
__e.TailApply(PrimFunc(symshen_4assert_d), V1577, symshen_4top)
return
}, 1)

tmp8435 := Call(__e, ns2_1set, symasserta, tmp8434)


_ = tmp8435

tmp8436 := MakeNative(func(__e *ControlFlow) {
V1578 := __e.Get(1)
_ = V1578
__e.TailApply(PrimFunc(symshen_4assert_d), V1578, symshen_4bottom)
return
}, 1)

tmp8437 := Call(__e, ns2_1set, symassertz, tmp8436)


_ = tmp8437

tmp8438 := MakeNative(func(__e *ControlFlow) {
V1579 := __e.Get(1)
_ = V1579
V1580 := __e.Get(2)
_ = V1580
tmp8472 := PrimIsPair(V1579)

var ifres8463 Obj

if True == tmp8472 {
tmp8470 := PrimTail(V1579)

tmp8471 := PrimIsPair(tmp8470)

var ifres8465 Obj

if True == tmp8471 {
tmp8467 := PrimTail(V1579)

tmp8468 := PrimHead(tmp8467)

tmp8469 := PrimEqual(sym_5_1_1, tmp8468)

var ifres8466 Obj

if True == tmp8469 {
ifres8466 = True


} else {
ifres8466 = False


}

ifres8465 = ifres8466


} else {
ifres8465 = False


}

var ifres8464 Obj

if True == ifres8465 {
ifres8464 = True


} else {
ifres8464 = False


}

ifres8463 = ifres8464


} else {
ifres8463 = False


}

if True == ifres8463 {
tmp8439 := MakeNative(func(__e *ControlFlow) {
W1581 := __e.Get(1)
_ = W1581
tmp8440 := MakeNative(func(__e *ControlFlow) {
W1582 := __e.Get(1)
_ = W1582
tmp8441 := MakeNative(func(__e *ControlFlow) {
W1583 := __e.Get(1)
_ = W1583
tmp8442 := MakeNative(func(__e *ControlFlow) {
W1584 := __e.Get(1)
_ = W1584
tmp8443 := MakeNative(func(__e *ControlFlow) {
W1585 := __e.Get(1)
_ = W1585
tmp8444 := MakeNative(func(__e *ControlFlow) {
W1586 := __e.Get(1)
_ = W1586
tmp8445 := MakeNative(func(__e *ControlFlow) {
W1587 := __e.Get(1)
_ = W1587
__e.Return(W1581)
return
}, 1)

tmp8446 := PrimTail(V1579)

tmp8447 := PrimTail(tmp8446)

tmp8448 := Call(__e, PrimFunc(symshen_4insert_1info), W1581, W1582, tmp8447, V1579, V1580)


__e.TailApply(tmp8445, tmp8448)
return


}, 1)

tmp8454 := PrimEqual(W1585, MakeNumber(-1))

var ifres8449 Obj

if True == tmp8454 {
tmp8450 := Call(__e, PrimFunc(symshen_4create_1skeleton), W1581, W1584)


tmp8451 := Call(__e, PrimFunc(symeval), tmp8450)


_ = tmp8451

tmp8452 := PrimValue(sym_dproperty_1vector_d)

tmp8453 := Call(__e, PrimFunc(symput), W1581, symshen_4dynamic, Nil, tmp8452)


ifres8449 = tmp8453


} else {
ifres8449 = symshen_4skip


}

__e.TailApply(tmp8444, ifres8449)
return


}, 1)

tmp8455 := Call(__e, PrimFunc(symarity), W1581)


__e.TailApply(tmp8443, tmp8455)
return


}, 1)

tmp8456 := Call(__e, PrimFunc(symshen_4parameters), W1583)


__e.TailApply(tmp8442, tmp8456)
return


}, 1)

tmp8457 := Call(__e, PrimFunc(symlength), W1582)


__e.TailApply(tmp8441, tmp8457)
return


}, 1)

tmp8458 := PrimHead(V1579)

tmp8459 := Call(__e, PrimFunc(symshen_4terms), tmp8458)


__e.TailApply(tmp8440, tmp8459)
return


}, 1)

tmp8460 := PrimHead(V1579)

tmp8461 := Call(__e, PrimFunc(symshen_4predicate), tmp8460)


__e.TailApply(tmp8439, tmp8461)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4assert_d)
return
}


}, 2)

tmp8473 := Call(__e, ns2_1set, symshen_4assert_d, tmp8438)


_ = tmp8473

tmp8474 := MakeNative(func(__e *ControlFlow) {
V1590 := __e.Get(1)
_ = V1590
tmp8476 := PrimIsPair(V1590)

if True == tmp8476 {
__e.Return(PrimHead(V1590))
return
} else {
__e.Return(V1590)
return
}


}, 1)

tmp8477 := Call(__e, ns2_1set, symshen_4predicate, tmp8474)


_ = tmp8477

tmp8478 := MakeNative(func(__e *ControlFlow) {
V1595 := __e.Get(1)
_ = V1595
tmp8480 := PrimIsPair(V1595)

if True == tmp8480 {
__e.Return(PrimTail(V1595))
return
} else {
__e.Return(Nil)
return
}


}, 1)

tmp8481 := Call(__e, ns2_1set, symshen_4terms, tmp8478)


_ = tmp8481

tmp8482 := MakeNative(func(__e *ControlFlow) {
V1596 := __e.Get(1)
_ = V1596
V1597 := __e.Get(2)
_ = V1597
tmp8483 := Call(__e, PrimFunc(symshen_4dynamic_1default), V1596, V1597)


tmp8484 := PrimCons(V1596, tmp8483)

__e.Return(PrimCons(symdefprolog, tmp8484))
return


}, 2)

tmp8485 := Call(__e, ns2_1set, symshen_4create_1skeleton, tmp8482)


_ = tmp8485

tmp8486 := MakeNative(func(__e *ControlFlow) {
V1598 := __e.Get(1)
_ = V1598
V1599 := __e.Get(2)
_ = V1599
tmp8487 := Call(__e, PrimFunc(symshen_4cons_1form), V1599)


tmp8488 := PrimCons(symshen_4dynamic, Nil)

tmp8489 := PrimCons(V1598, tmp8488)

tmp8490 := PrimCons(symget, tmp8489)

tmp8491 := PrimCons(tmp8490, Nil)

tmp8492 := PrimCons(tmp8487, tmp8491)

tmp8493 := PrimCons(symshen_4call_1dynamic, tmp8492)

tmp8494 := PrimIntern(MakeString(";"))

tmp8495 := PrimCons(tmp8494, Nil)

tmp8496 := PrimCons(tmp8493, tmp8495)

tmp8497 := PrimCons(sym_5_1_1, tmp8496)

__e.TailApply(PrimFunc(symappend), V1599, tmp8497)
return


}, 2)

tmp8498 := Call(__e, ns2_1set, symshen_4dynamic_1default, tmp8486)


_ = tmp8498

tmp8499 := MakeNative(func(__e *ControlFlow) {
V1600 := __e.Get(1)
_ = V1600
V1601 := __e.Get(2)
_ = V1601
V1602 := __e.Get(3)
_ = V1602
V1603 := __e.Get(4)
_ = V1603
V1604 := __e.Get(5)
_ = V1604
tmp8500 := MakeNative(func(__e *ControlFlow) {
W1605 := __e.Get(1)
_ = W1605
tmp8501 := MakeNative(func(__e *ControlFlow) {
W1606 := __e.Get(1)
_ = W1606
tmp8502 := MakeNative(func(__e *ControlFlow) {
W1607 := __e.Get(1)
_ = W1607
tmp8503 := MakeNative(func(__e *ControlFlow) {
W1608 := __e.Get(1)
_ = W1608
tmp8504 := MakeNative(func(__e *ControlFlow) {
W1609 := __e.Get(1)
_ = W1609
tmp8505 := PrimValue(sym_dproperty_1vector_d)

__e.TailApply(PrimFunc(symput), V1600, symshen_4dynamic, W1609, tmp8505)
return


}, 1)

tmp8510 := PrimEqual(V1604, symshen_4top)

var ifres8506 Obj

if True == tmp8510 {
tmp8507 := PrimCons(W1607, W1608)

ifres8506 = tmp8507


} else {
tmp8508 := PrimCons(W1607, Nil)

tmp8509 := Call(__e, PrimFunc(symappend), W1608, tmp8508)


ifres8506 = tmp8509


}

__e.TailApply(tmp8504, ifres8506)
return


}, 1)

tmp8511 := PrimValue(sym_dproperty_1vector_d)

tmp8512 := Call(__e, PrimFunc(symget), V1600, symshen_4dynamic, tmp8511)


__e.TailApply(tmp8503, tmp8512)
return


}, 1)

tmp8513 := Call(__e, PrimFunc(symfn), W1605)


tmp8514 := PrimCons(W1605, V1603)

tmp8515 := PrimCons(tmp8513, tmp8514)

__e.TailApply(tmp8502, tmp8515)
return


}, 1)

tmp8516 := PrimCons(W1605, Nil)

tmp8517 := PrimCons(symdefprolog, tmp8516)

tmp8518 := PrimCons(sym_5_1_1, V1602)

tmp8519 := Call(__e, PrimFunc(symappend), V1601, tmp8518)


tmp8520 := Call(__e, PrimFunc(symappend), tmp8517, tmp8519)


tmp8521 := Call(__e, PrimFunc(symeval), tmp8520)


__e.TailApply(tmp8501, tmp8521)
return


}, 1)

tmp8522 := Call(__e, PrimFunc(symgensym), symshen_4g)


__e.TailApply(tmp8500, tmp8522)
return


}, 5)

tmp8523 := Call(__e, ns2_1set, symshen_4insert_1info, tmp8499)


_ = tmp8523

tmp8524 := MakeNative(func(__e *ControlFlow) {
tmp8525 := MakeNative(func(__e *ControlFlow) {
W1610 := __e.Get(1)
_ = W1610
tmp8526 := MakeNative(func(__e *ControlFlow) {
W1611 := __e.Get(1)
_ = W1611
__e.Return(W1611)
return
}, 1)

tmp8532 := Call(__e, PrimFunc(symempty_2), W1610)


var ifres8527 Obj

if True == tmp8532 {
tmp8528 := Call(__e, PrimFunc(symgensym), symshen_4g)


ifres8527 = tmp8528


} else {
tmp8529 := PrimTail(W1610)

tmp8530 := PrimSet(symshen_4_dnames_d, tmp8529)

_ = tmp8530

tmp8531 := PrimHead(W1610)

ifres8527 = tmp8531


}

__e.TailApply(tmp8526, ifres8527)
return


}, 1)

tmp8533 := PrimValue(symshen_4_dnames_d)

__e.TailApply(tmp8525, tmp8533)
return


}, 0)

tmp8534 := Call(__e, ns2_1set, symshen_4newname, tmp8524)


_ = tmp8534

tmp8535 := MakeNative(func(__e *ControlFlow) {
V1612 := __e.Get(1)
_ = V1612
V1613 := __e.Get(2)
_ = V1613
V1614 := __e.Get(3)
_ = V1614
V1615 := __e.Get(4)
_ = V1615
V1616 := __e.Get(5)
_ = V1616
V1617 := __e.Get(6)
_ = V1617
tmp8536 := MakeNative(func(__e *ControlFlow) {
W1618 := __e.Get(1)
_ = W1618
tmp8547 := PrimEqual(W1618, False)

if True == tmp8547 {
tmp8545 := Call(__e, PrimFunc(symshen_4unlocked_2), V1615)


if True == tmp8545 {
tmp8537 := MakeNative(func(__e *ControlFlow) {
W1622 := __e.Get(1)
_ = W1622
tmp8542 := PrimIsPair(W1622)

if True == tmp8542 {
tmp8538 := MakeNative(func(__e *ControlFlow) {
W1623 := __e.Get(1)
_ = W1623
tmp8539 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp8539

__e.TailApply(PrimFunc(symshen_4call_1dynamic), V1612, W1623, V1614, V1615, V1616, V1617)
return


}, 1)

tmp8540 := PrimTail(W1622)

__e.TailApply(tmp8538, tmp8540)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp8543 := Call(__e, PrimFunc(symshen_4lazyderef), V1613, V1614)


__e.TailApply(tmp8537, tmp8543)
return


} else {
__e.Return(False)
return
}


} else {
__e.Return(W1618)
return
}


}, 1)

tmp8562 := Call(__e, PrimFunc(symshen_4unlocked_2), V1615)


var ifres8548 Obj

if True == tmp8562 {
tmp8549 := MakeNative(func(__e *ControlFlow) {
W1619 := __e.Get(1)
_ = W1619
tmp8559 := PrimIsPair(W1619)

if True == tmp8559 {
tmp8550 := MakeNative(func(__e *ControlFlow) {
W1620 := __e.Get(1)
_ = W1620
tmp8555 := PrimIsPair(W1620)

if True == tmp8555 {
tmp8551 := MakeNative(func(__e *ControlFlow) {
W1621 := __e.Get(1)
_ = W1621
tmp8552 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp8552

__e.TailApply(PrimFunc(symshen_4callrec), W1621, V1612, V1614, V1615, V1616, V1617)
return


}, 1)

tmp8553 := PrimHead(W1620)

__e.TailApply(tmp8551, tmp8553)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp8556 := PrimHead(W1619)

tmp8557 := Call(__e, PrimFunc(symshen_4lazyderef), tmp8556, V1614)


__e.TailApply(tmp8550, tmp8557)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp8560 := Call(__e, PrimFunc(symshen_4lazyderef), V1613, V1614)


tmp8561 := Call(__e, tmp8549, tmp8560)


ifres8548 = tmp8561


} else {
ifres8548 = False


}

__e.TailApply(tmp8536, ifres8548)
return


}, 6)

tmp8563 := Call(__e, ns2_1set, symshen_4call_1dynamic, tmp8535)


_ = tmp8563

tmp8564 := MakeNative(func(__e *ControlFlow) {
V1624 := __e.Get(1)
_ = V1624
V1625 := __e.Get(2)
_ = V1625
V1626 := __e.Get(3)
_ = V1626
V1627 := __e.Get(4)
_ = V1627
V1628 := __e.Get(5)
_ = V1628
V1629 := __e.Get(6)
_ = V1629
tmp8574 := PrimEqual(Nil, V1625)

if True == tmp8574 {
tmp8565 := Call(__e, V1624, V1626)


tmp8566 := Call(__e, tmp8565, V1627)


tmp8567 := Call(__e, tmp8566, V1628)


__e.TailApply(tmp8567, V1629)
return


} else {
tmp8572 := PrimIsPair(V1625)

if True == tmp8572 {
tmp8568 := PrimHead(V1625)

tmp8569 := Call(__e, V1624, tmp8568)


tmp8570 := PrimTail(V1625)

__e.TailApply(PrimFunc(symshen_4callrec), tmp8569, tmp8570, V1626, V1627, V1628, V1629)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4callrec)
return
}


}


}, 6)

tmp8575 := Call(__e, ns2_1set, symshen_4callrec, tmp8564)


_ = tmp8575

tmp8576 := MakeNative(func(__e *ControlFlow) {
V1630 := __e.Get(1)
_ = V1630
tmp8595 := PrimIsPair(V1630)

var ifres8586 Obj

if True == tmp8595 {
tmp8593 := PrimTail(V1630)

tmp8594 := PrimIsPair(tmp8593)

var ifres8588 Obj

if True == tmp8594 {
tmp8590 := PrimTail(V1630)

tmp8591 := PrimHead(tmp8590)

tmp8592 := PrimEqual(sym_5_1_1, tmp8591)

var ifres8589 Obj

if True == tmp8592 {
ifres8589 = True


} else {
ifres8589 = False


}

ifres8588 = ifres8589


} else {
ifres8588 = False


}

var ifres8587 Obj

if True == ifres8588 {
ifres8587 = True


} else {
ifres8587 = False


}

ifres8586 = ifres8587


} else {
ifres8586 = False


}

if True == ifres8586 {
tmp8577 := MakeNative(func(__e *ControlFlow) {
W1631 := __e.Get(1)
_ = W1631
tmp8578 := MakeNative(func(__e *ControlFlow) {
W1632 := __e.Get(1)
_ = W1632
tmp8579 := Call(__e, PrimFunc(symshen_4retract_1clause), V1630, W1632)


tmp8580 := PrimValue(sym_dproperty_1vector_d)

__e.TailApply(PrimFunc(symput), W1631, symshen_4dynamic, tmp8579, tmp8580)
return


}, 1)

tmp8581 := PrimValue(sym_dproperty_1vector_d)

tmp8582 := Call(__e, PrimFunc(symget), W1631, symshen_4dynamic, tmp8581)


__e.TailApply(tmp8578, tmp8582)
return


}, 1)

tmp8583 := PrimHead(V1630)

tmp8584 := Call(__e, PrimFunc(symshen_4predicate), tmp8583)


__e.TailApply(tmp8577, tmp8584)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symretract)
return
}


}, 1)

tmp8596 := Call(__e, ns2_1set, symretract, tmp8576)


_ = tmp8596

tmp8597 := MakeNative(func(__e *ControlFlow) {
V1638 := __e.Get(1)
_ = V1638
V1639 := __e.Get(2)
_ = V1639
tmp8627 := PrimEqual(Nil, V1639)

if True == tmp8627 {
__e.Return(Nil)
return
} else {
tmp8625 := PrimIsPair(V1639)

var ifres8610 Obj

if True == tmp8625 {
tmp8623 := PrimHead(V1639)

tmp8624 := PrimIsPair(tmp8623)

var ifres8612 Obj

if True == tmp8624 {
tmp8620 := PrimHead(V1639)

tmp8621 := PrimTail(tmp8620)

tmp8622 := PrimIsPair(tmp8621)

var ifres8614 Obj

if True == tmp8622 {
tmp8616 := PrimHead(V1639)

tmp8617 := PrimTail(tmp8616)

tmp8618 := PrimTail(tmp8617)

tmp8619 := PrimEqual(V1638, tmp8618)

var ifres8615 Obj

if True == tmp8619 {
ifres8615 = True


} else {
ifres8615 = False


}

ifres8614 = ifres8615


} else {
ifres8614 = False


}

var ifres8613 Obj

if True == ifres8614 {
ifres8613 = True


} else {
ifres8613 = False


}

ifres8612 = ifres8613


} else {
ifres8612 = False


}

var ifres8611 Obj

if True == ifres8612 {
ifres8611 = True


} else {
ifres8611 = False


}

ifres8610 = ifres8611


} else {
ifres8610 = False


}

if True == ifres8610 {
tmp8598 := PrimHead(V1639)

tmp8599 := PrimTail(tmp8598)

tmp8600 := PrimHead(tmp8599)

tmp8601 := PrimValue(symshen_4_dnames_d)

tmp8602 := PrimCons(tmp8600, tmp8601)

tmp8603 := PrimSet(symshen_4_dnames_d, tmp8602)

_ = tmp8603

__e.Return(PrimTail(V1639))
return


} else {
tmp8608 := PrimIsPair(V1639)

if True == tmp8608 {
tmp8604 := PrimHead(V1639)

tmp8605 := PrimTail(V1639)

tmp8606 := Call(__e, PrimFunc(symshen_4retract_1clause), V1638, tmp8605)


__e.Return(PrimCons(tmp8604, tmp8606))
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4retract_1clause)
return
}


}


}


}, 2)

tmp8628 := Call(__e, ns2_1set, symshen_4retract_1clause, tmp8597)


_ = tmp8628

tmp8629 := MakeNative(func(__e *ControlFlow) {
V1640 := __e.Get(1)
_ = V1640
V1641 := __e.Get(2)
_ = V1641
tmp8630 := MakeNative(func(__e *ControlFlow) {
Z1642 := __e.Get(1)
_ = Z1642
__e.TailApply(PrimFunc(symshen_4_5defprolog_6), Z1642)
return
}, 1)

tmp8631 := PrimCons(V1640, V1641)

__e.TailApply(PrimFunc(symcompile), tmp8630, tmp8631)
return


}, 2)

tmp8632 := Call(__e, ns2_1set, symshen_4compile_1prolog, tmp8629)


_ = tmp8632

tmp8633 := MakeNative(func(__e *ControlFlow) {
V1643 := __e.Get(1)
_ = V1643
tmp8634 := MakeNative(func(__e *ControlFlow) {
W1644 := __e.Get(1)
_ = W1644
tmp8636 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1644)


if True == tmp8636 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W1644)
return
}


}, 1)

tmp8658 := PrimIsPair(V1643)

var ifres8637 Obj

if True == tmp8658 {
tmp8638 := MakeNative(func(__e *ControlFlow) {
W1645 := __e.Get(1)
_ = W1645
tmp8639 := MakeNative(func(__e *ControlFlow) {
W1646 := __e.Get(1)
_ = W1646
tmp8640 := MakeNative(func(__e *ControlFlow) {
W1647 := __e.Get(1)
_ = W1647
tmp8652 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1647)


if True == tmp8652 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp8641 := MakeNative(func(__e *ControlFlow) {
W1648 := __e.Get(1)
_ = W1648
tmp8642 := MakeNative(func(__e *ControlFlow) {
W1649 := __e.Get(1)
_ = W1649
tmp8643 := MakeNative(func(__e *ControlFlow) {
W1650 := __e.Get(1)
_ = W1650
tmp8644 := MakeNative(func(__e *ControlFlow) {
W1651 := __e.Get(1)
_ = W1651
__e.TailApply(PrimFunc(symshen_4horn_1clause_1procedure), W1645, W1651)
return
}, 1)

tmp8645 := MakeNative(func(__e *ControlFlow) {
Z1652 := __e.Get(1)
_ = Z1652
__e.TailApply(PrimFunc(symshen_4linearise_1clause), Z1652)
return
}, 1)

tmp8646 := Call(__e, PrimFunc(symmap), tmp8645, W1648)


__e.TailApply(tmp8644, tmp8646)
return


}, 1)

tmp8647 := Call(__e, PrimFunc(symshen_4prolog_1arity_1check), W1645, W1648)


tmp8648 := Call(__e, tmp8643, tmp8647)


__e.TailApply(PrimFunc(symshen_4comb), W1649, tmp8648)
return


}, 1)

tmp8649 := Call(__e, PrimFunc(symshen_4in_1_6), W1647)


__e.TailApply(tmp8642, tmp8649)
return


}, 1)

tmp8650 := Call(__e, PrimFunc(symshen_4_5_1out), W1647)


__e.TailApply(tmp8641, tmp8650)
return


}


}, 1)

tmp8653 := Call(__e, PrimFunc(symshen_4_5clauses_6), W1646)


__e.TailApply(tmp8640, tmp8653)
return


}, 1)

tmp8654 := Call(__e, PrimFunc(symtail), V1643)


__e.TailApply(tmp8639, tmp8654)
return


}, 1)

tmp8655 := Call(__e, PrimFunc(symhead), V1643)


tmp8656 := Call(__e, tmp8638, tmp8655)


ifres8637 = tmp8656


} else {
tmp8657 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres8637 = tmp8657


}

__e.TailApply(tmp8634, ifres8637)
return


}, 1)

tmp8659 := Call(__e, ns2_1set, symshen_4_5defprolog_6, tmp8633)


_ = tmp8659

tmp8660 := MakeNative(func(__e *ControlFlow) {
V1655 := __e.Get(1)
_ = V1655
V1656 := __e.Get(2)
_ = V1656
tmp8704 := PrimIsPair(V1656)

var ifres8685 Obj

if True == tmp8704 {
tmp8702 := PrimHead(V1656)

tmp8703 := PrimIsPair(tmp8702)

var ifres8687 Obj

if True == tmp8703 {
tmp8699 := PrimHead(V1656)

tmp8700 := PrimTail(tmp8699)

tmp8701 := PrimIsPair(tmp8700)

var ifres8689 Obj

if True == tmp8701 {
tmp8695 := PrimHead(V1656)

tmp8696 := PrimTail(tmp8695)

tmp8697 := PrimTail(tmp8696)

tmp8698 := PrimEqual(Nil, tmp8697)

var ifres8691 Obj

if True == tmp8698 {
tmp8693 := PrimTail(V1656)

tmp8694 := PrimEqual(Nil, tmp8693)

var ifres8692 Obj

if True == tmp8694 {
ifres8692 = True


} else {
ifres8692 = False


}

ifres8691 = ifres8692


} else {
ifres8691 = False


}

var ifres8690 Obj

if True == ifres8691 {
ifres8690 = True


} else {
ifres8690 = False


}

ifres8689 = ifres8690


} else {
ifres8689 = False


}

var ifres8688 Obj

if True == ifres8689 {
ifres8688 = True


} else {
ifres8688 = False


}

ifres8687 = ifres8688


} else {
ifres8687 = False


}

var ifres8686 Obj

if True == ifres8687 {
ifres8686 = True


} else {
ifres8686 = False


}

ifres8685 = ifres8686


} else {
ifres8685 = False


}

if True == ifres8685 {
tmp8661 := PrimHead(V1656)

tmp8662 := PrimHead(tmp8661)

__e.TailApply(PrimFunc(symlength), tmp8662)
return


} else {
tmp8683 := PrimIsPair(V1656)

var ifres8668 Obj

if True == tmp8683 {
tmp8681 := PrimHead(V1656)

tmp8682 := PrimIsPair(tmp8681)

var ifres8670 Obj

if True == tmp8682 {
tmp8678 := PrimHead(V1656)

tmp8679 := PrimTail(tmp8678)

tmp8680 := PrimIsPair(tmp8679)

var ifres8672 Obj

if True == tmp8680 {
tmp8674 := PrimHead(V1656)

tmp8675 := PrimTail(tmp8674)

tmp8676 := PrimTail(tmp8675)

tmp8677 := PrimEqual(Nil, tmp8676)

var ifres8673 Obj

if True == tmp8677 {
ifres8673 = True


} else {
ifres8673 = False


}

ifres8672 = ifres8673


} else {
ifres8672 = False


}

var ifres8671 Obj

if True == ifres8672 {
ifres8671 = True


} else {
ifres8671 = False


}

ifres8670 = ifres8671


} else {
ifres8670 = False


}

var ifres8669 Obj

if True == ifres8670 {
ifres8669 = True


} else {
ifres8669 = False


}

ifres8668 = ifres8669


} else {
ifres8668 = False


}

if True == ifres8668 {
tmp8663 := PrimHead(V1656)

tmp8664 := PrimHead(tmp8663)

tmp8665 := Call(__e, PrimFunc(symlength), tmp8664)


tmp8666 := PrimTail(V1656)

__e.TailApply(PrimFunc(symshen_4pac_1h), V1655, tmp8665, tmp8666)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4prolog_1arity_1check)
return
}


}


}, 2)

tmp8705 := Call(__e, ns2_1set, symshen_4prolog_1arity_1check, tmp8660)


_ = tmp8705

tmp8706 := MakeNative(func(__e *ControlFlow) {
V1661 := __e.Get(1)
_ = V1661
V1662 := __e.Get(2)
_ = V1662
V1663 := __e.Get(3)
_ = V1663
tmp8722 := PrimEqual(Nil, V1663)

if True == tmp8722 {
__e.Return(V1662)
return
} else {
tmp8720 := PrimIsPair(V1663)

var ifres8716 Obj

if True == tmp8720 {
tmp8718 := PrimHead(V1663)

tmp8719 := PrimIsPair(tmp8718)

var ifres8717 Obj

if True == tmp8719 {
ifres8717 = True


} else {
ifres8717 = False


}

ifres8716 = ifres8717


} else {
ifres8716 = False


}

if True == ifres8716 {
tmp8711 := PrimHead(V1663)

tmp8712 := PrimHead(tmp8711)

tmp8713 := Call(__e, PrimFunc(symlength), tmp8712)


tmp8714 := PrimEqual(V1662, tmp8713)

if True == tmp8714 {
tmp8707 := PrimTail(V1663)

__e.TailApply(PrimFunc(symshen_4pac_1h), V1661, V1662, tmp8707)
return


} else {
tmp8708 := Call(__e, PrimFunc(symshen_4app), V1661, MakeString("\n"), symshen_4a)


tmp8709 := PrimStringConcat(MakeString("arity error in prolog procedure "), tmp8708)

__e.Return(PrimSimpleError(tmp8709))
return


}


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4pac_1h)
return
}


}


}, 3)

tmp8723 := Call(__e, ns2_1set, symshen_4pac_1h, tmp8706)


_ = tmp8723

tmp8724 := MakeNative(func(__e *ControlFlow) {
V1664 := __e.Get(1)
_ = V1664
tmp8725 := MakeNative(func(__e *ControlFlow) {
W1665 := __e.Get(1)
_ = W1665
tmp8744 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1665)


if True == tmp8744 {
tmp8726 := MakeNative(func(__e *ControlFlow) {
W1672 := __e.Get(1)
_ = W1672
tmp8728 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1672)


if True == tmp8728 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W1672)
return
}


}, 1)

tmp8729 := MakeNative(func(__e *ControlFlow) {
W1673 := __e.Get(1)
_ = W1673
tmp8740 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1673)


if True == tmp8740 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp8730 := MakeNative(func(__e *ControlFlow) {
W1674 := __e.Get(1)
_ = W1674
tmp8731 := MakeNative(func(__e *ControlFlow) {
W1675 := __e.Get(1)
_ = W1675
tmp8736 := Call(__e, PrimFunc(symempty_2), W1674)


var ifres8732 Obj

if True == tmp8736 {
ifres8732 = Nil


} else {
tmp8733 := Call(__e, PrimFunc(symshen_4app), W1674, MakeString("\n ..."), symshen_4r)


tmp8734 := PrimStringConcat(MakeString("Prolog syntax error here:\n "), tmp8733)

tmp8735 := PrimSimpleError(tmp8734)

ifres8732 = tmp8735


}

__e.TailApply(PrimFunc(symshen_4comb), W1675, ifres8732)
return


}, 1)

tmp8737 := Call(__e, PrimFunc(symshen_4in_1_6), W1673)


__e.TailApply(tmp8731, tmp8737)
return


}, 1)

tmp8738 := Call(__e, PrimFunc(symshen_4_5_1out), W1673)


__e.TailApply(tmp8730, tmp8738)
return


}


}, 1)

tmp8741 := Call(__e, PrimFunc(sym_5_b_6), V1664)


tmp8742 := Call(__e, tmp8729, tmp8741)


__e.TailApply(tmp8726, tmp8742)
return


} else {
__e.Return(W1665)
return
}


}, 1)

tmp8745 := MakeNative(func(__e *ControlFlow) {
W1666 := __e.Get(1)
_ = W1666
tmp8760 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1666)


if True == tmp8760 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp8746 := MakeNative(func(__e *ControlFlow) {
W1667 := __e.Get(1)
_ = W1667
tmp8747 := MakeNative(func(__e *ControlFlow) {
W1668 := __e.Get(1)
_ = W1668
tmp8748 := MakeNative(func(__e *ControlFlow) {
W1669 := __e.Get(1)
_ = W1669
tmp8755 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1669)


if True == tmp8755 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp8749 := MakeNative(func(__e *ControlFlow) {
W1670 := __e.Get(1)
_ = W1670
tmp8750 := MakeNative(func(__e *ControlFlow) {
W1671 := __e.Get(1)
_ = W1671
tmp8751 := PrimCons(W1667, W1670)

__e.TailApply(PrimFunc(symshen_4comb), W1671, tmp8751)
return


}, 1)

tmp8752 := Call(__e, PrimFunc(symshen_4in_1_6), W1669)


__e.TailApply(tmp8750, tmp8752)
return


}, 1)

tmp8753 := Call(__e, PrimFunc(symshen_4_5_1out), W1669)


__e.TailApply(tmp8749, tmp8753)
return


}


}, 1)

tmp8756 := Call(__e, PrimFunc(symshen_4_5clauses_6), W1668)


__e.TailApply(tmp8748, tmp8756)
return


}, 1)

tmp8757 := Call(__e, PrimFunc(symshen_4in_1_6), W1666)


__e.TailApply(tmp8747, tmp8757)
return


}, 1)

tmp8758 := Call(__e, PrimFunc(symshen_4_5_1out), W1666)


__e.TailApply(tmp8746, tmp8758)
return


}


}, 1)

tmp8761 := Call(__e, PrimFunc(symshen_4_5clause_6), V1664)


tmp8762 := Call(__e, tmp8745, tmp8761)


__e.TailApply(tmp8725, tmp8762)
return


}, 1)

tmp8763 := Call(__e, ns2_1set, symshen_4_5clauses_6, tmp8724)


_ = tmp8763

tmp8764 := MakeNative(func(__e *ControlFlow) {
V1676 := __e.Get(1)
_ = V1676
tmp8780 := PrimIsPair(V1676)

var ifres8771 Obj

if True == tmp8780 {
tmp8778 := PrimTail(V1676)

tmp8779 := PrimIsPair(tmp8778)

var ifres8773 Obj

if True == tmp8779 {
tmp8775 := PrimTail(V1676)

tmp8776 := PrimTail(tmp8775)

tmp8777 := PrimEqual(Nil, tmp8776)

var ifres8774 Obj

if True == tmp8777 {
ifres8774 = True


} else {
ifres8774 = False


}

ifres8773 = ifres8774


} else {
ifres8773 = False


}

var ifres8772 Obj

if True == ifres8773 {
ifres8772 = True


} else {
ifres8772 = False


}

ifres8771 = ifres8772


} else {
ifres8771 = False


}

if True == ifres8771 {
tmp8765 := PrimHead(V1676)

tmp8766 := PrimTail(V1676)

tmp8767 := PrimHead(tmp8766)

tmp8768 := Call(__e, PrimFunc(sym_8p), tmp8765, tmp8767)


tmp8769 := Call(__e, PrimFunc(symshen_4linearise), tmp8768)


__e.TailApply(PrimFunc(symshen_4lch), tmp8769)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4linearise_1clause)
return
}


}, 1)

tmp8781 := Call(__e, ns2_1set, symshen_4linearise_1clause, tmp8764)


_ = tmp8781

tmp8782 := MakeNative(func(__e *ControlFlow) {
V1677 := __e.Get(1)
_ = V1677
tmp8788 := Call(__e, PrimFunc(symtuple_2), V1677)


if True == tmp8788 {
tmp8783 := Call(__e, PrimFunc(symfst), V1677)


tmp8784 := Call(__e, PrimFunc(symsnd), V1677)


tmp8785 := Call(__e, PrimFunc(symshen_4lchh), tmp8784)


tmp8786 := PrimCons(tmp8785, Nil)

__e.Return(PrimCons(tmp8783, tmp8786))
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4lch)
return
}


}, 1)

tmp8789 := Call(__e, ns2_1set, symshen_4lch, tmp8782)


_ = tmp8789

tmp8790 := MakeNative(func(__e *ControlFlow) {
V1678 := __e.Get(1)
_ = V1678
tmp8853 := PrimIsPair(V1678)

var ifres8802 Obj

if True == tmp8853 {
tmp8851 := PrimHead(V1678)

tmp8852 := PrimEqual(symwhere, tmp8851)

var ifres8804 Obj

if True == tmp8852 {
tmp8849 := PrimTail(V1678)

tmp8850 := PrimIsPair(tmp8849)

var ifres8806 Obj

if True == tmp8850 {
tmp8846 := PrimTail(V1678)

tmp8847 := PrimHead(tmp8846)

tmp8848 := PrimIsPair(tmp8847)

var ifres8808 Obj

if True == tmp8848 {
tmp8842 := PrimTail(V1678)

tmp8843 := PrimHead(tmp8842)

tmp8844 := PrimHead(tmp8843)

tmp8845 := PrimEqual(sym_a, tmp8844)

var ifres8810 Obj

if True == tmp8845 {
tmp8838 := PrimTail(V1678)

tmp8839 := PrimHead(tmp8838)

tmp8840 := PrimTail(tmp8839)

tmp8841 := PrimIsPair(tmp8840)

var ifres8812 Obj

if True == tmp8841 {
tmp8833 := PrimTail(V1678)

tmp8834 := PrimHead(tmp8833)

tmp8835 := PrimTail(tmp8834)

tmp8836 := PrimTail(tmp8835)

tmp8837 := PrimIsPair(tmp8836)

var ifres8814 Obj

if True == tmp8837 {
tmp8827 := PrimTail(V1678)

tmp8828 := PrimHead(tmp8827)

tmp8829 := PrimTail(tmp8828)

tmp8830 := PrimTail(tmp8829)

tmp8831 := PrimTail(tmp8830)

tmp8832 := PrimEqual(Nil, tmp8831)

var ifres8816 Obj

if True == tmp8832 {
tmp8824 := PrimTail(V1678)

tmp8825 := PrimTail(tmp8824)

tmp8826 := PrimIsPair(tmp8825)

var ifres8818 Obj

if True == tmp8826 {
tmp8820 := PrimTail(V1678)

tmp8821 := PrimTail(tmp8820)

tmp8822 := PrimTail(tmp8821)

tmp8823 := PrimEqual(Nil, tmp8822)

var ifres8819 Obj

if True == tmp8823 {
ifres8819 = True


} else {
ifres8819 = False


}

ifres8818 = ifres8819


} else {
ifres8818 = False


}

var ifres8817 Obj

if True == ifres8818 {
ifres8817 = True


} else {
ifres8817 = False


}

ifres8816 = ifres8817


} else {
ifres8816 = False


}

var ifres8815 Obj

if True == ifres8816 {
ifres8815 = True


} else {
ifres8815 = False


}

ifres8814 = ifres8815


} else {
ifres8814 = False


}

var ifres8813 Obj

if True == ifres8814 {
ifres8813 = True


} else {
ifres8813 = False


}

ifres8812 = ifres8813


} else {
ifres8812 = False


}

var ifres8811 Obj

if True == ifres8812 {
ifres8811 = True


} else {
ifres8811 = False


}

ifres8810 = ifres8811


} else {
ifres8810 = False


}

var ifres8809 Obj

if True == ifres8810 {
ifres8809 = True


} else {
ifres8809 = False


}

ifres8808 = ifres8809


} else {
ifres8808 = False


}

var ifres8807 Obj

if True == ifres8808 {
ifres8807 = True


} else {
ifres8807 = False


}

ifres8806 = ifres8807


} else {
ifres8806 = False


}

var ifres8805 Obj

if True == ifres8806 {
ifres8805 = True


} else {
ifres8805 = False


}

ifres8804 = ifres8805


} else {
ifres8804 = False


}

var ifres8803 Obj

if True == ifres8804 {
ifres8803 = True


} else {
ifres8803 = False


}

ifres8802 = ifres8803


} else {
ifres8802 = False


}

if True == ifres8802 {
tmp8792 := PrimValue(symshen_4_doccurs_d)

var ifres8791 Obj

if True == tmp8792 {
ifres8791 = symis_b


} else {
ifres8791 = symis


}

tmp8793 := PrimTail(V1678)

tmp8794 := PrimHead(tmp8793)

tmp8795 := PrimTail(tmp8794)

tmp8796 := PrimCons(ifres8791, tmp8795)

tmp8797 := PrimTail(V1678)

tmp8798 := PrimTail(tmp8797)

tmp8799 := PrimHead(tmp8798)

tmp8800 := Call(__e, PrimFunc(symshen_4lchh), tmp8799)


__e.Return(PrimCons(tmp8796, tmp8800))
return


} else {
__e.Return(V1678)
return
}


}, 1)

tmp8854 := Call(__e, ns2_1set, symshen_4lchh, tmp8790)


_ = tmp8854

tmp8855 := MakeNative(func(__e *ControlFlow) {
V1679 := __e.Get(1)
_ = V1679
tmp8856 := MakeNative(func(__e *ControlFlow) {
W1680 := __e.Get(1)
_ = W1680
tmp8858 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1680)


if True == tmp8858 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W1680)
return
}


}, 1)

tmp8859 := MakeNative(func(__e *ControlFlow) {
W1681 := __e.Get(1)
_ = W1681
tmp8885 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1681)


if True == tmp8885 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp8860 := MakeNative(func(__e *ControlFlow) {
W1682 := __e.Get(1)
_ = W1682
tmp8861 := MakeNative(func(__e *ControlFlow) {
W1683 := __e.Get(1)
_ = W1683
tmp8881 := Call(__e, PrimFunc(symshen_4hds_a_2), W1683, sym_5_1_1)


if True == tmp8881 {
tmp8862 := MakeNative(func(__e *ControlFlow) {
W1684 := __e.Get(1)
_ = W1684
tmp8863 := MakeNative(func(__e *ControlFlow) {
W1685 := __e.Get(1)
_ = W1685
tmp8877 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1685)


if True == tmp8877 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp8864 := MakeNative(func(__e *ControlFlow) {
W1686 := __e.Get(1)
_ = W1686
tmp8865 := MakeNative(func(__e *ControlFlow) {
W1687 := __e.Get(1)
_ = W1687
tmp8866 := MakeNative(func(__e *ControlFlow) {
W1688 := __e.Get(1)
_ = W1688
tmp8872 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1688)


if True == tmp8872 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp8867 := MakeNative(func(__e *ControlFlow) {
W1689 := __e.Get(1)
_ = W1689
tmp8868 := PrimCons(W1686, Nil)

tmp8869 := PrimCons(W1682, tmp8868)

__e.TailApply(PrimFunc(symshen_4comb), W1689, tmp8869)
return


}, 1)

tmp8870 := Call(__e, PrimFunc(symshen_4in_1_6), W1688)


__e.TailApply(tmp8867, tmp8870)
return


}


}, 1)

tmp8873 := Call(__e, PrimFunc(symshen_4_5sc_6), W1687)


__e.TailApply(tmp8866, tmp8873)
return


}, 1)

tmp8874 := Call(__e, PrimFunc(symshen_4in_1_6), W1685)


__e.TailApply(tmp8865, tmp8874)
return


}, 1)

tmp8875 := Call(__e, PrimFunc(symshen_4_5_1out), W1685)


__e.TailApply(tmp8864, tmp8875)
return


}


}, 1)

tmp8878 := Call(__e, PrimFunc(symshen_4_5body_6), W1684)


__e.TailApply(tmp8863, tmp8878)
return


}, 1)

tmp8879 := Call(__e, PrimFunc(symtail), W1683)


__e.TailApply(tmp8862, tmp8879)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp8882 := Call(__e, PrimFunc(symshen_4in_1_6), W1681)


__e.TailApply(tmp8861, tmp8882)
return


}, 1)

tmp8883 := Call(__e, PrimFunc(symshen_4_5_1out), W1681)


__e.TailApply(tmp8860, tmp8883)
return


}


}, 1)

tmp8886 := Call(__e, PrimFunc(symshen_4_5head_6), V1679)


tmp8887 := Call(__e, tmp8859, tmp8886)


__e.TailApply(tmp8856, tmp8887)
return


}, 1)

tmp8888 := Call(__e, ns2_1set, symshen_4_5clause_6, tmp8855)


_ = tmp8888

tmp8889 := MakeNative(func(__e *ControlFlow) {
V1690 := __e.Get(1)
_ = V1690
tmp8890 := MakeNative(func(__e *ControlFlow) {
W1691 := __e.Get(1)
_ = W1691
tmp8902 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1691)


if True == tmp8902 {
tmp8891 := MakeNative(func(__e *ControlFlow) {
W1698 := __e.Get(1)
_ = W1698
tmp8893 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1698)


if True == tmp8893 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W1698)
return
}


}, 1)

tmp8894 := MakeNative(func(__e *ControlFlow) {
W1699 := __e.Get(1)
_ = W1699
tmp8898 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1699)


if True == tmp8898 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp8895 := MakeNative(func(__e *ControlFlow) {
W1700 := __e.Get(1)
_ = W1700
__e.TailApply(PrimFunc(symshen_4comb), W1700, Nil)
return
}, 1)

tmp8896 := Call(__e, PrimFunc(symshen_4in_1_6), W1699)


__e.TailApply(tmp8895, tmp8896)
return


}


}, 1)

tmp8899 := Call(__e, PrimFunc(sym_5e_6), V1690)


tmp8900 := Call(__e, tmp8894, tmp8899)


__e.TailApply(tmp8891, tmp8900)
return


} else {
__e.Return(W1691)
return
}


}, 1)

tmp8903 := MakeNative(func(__e *ControlFlow) {
W1692 := __e.Get(1)
_ = W1692
tmp8918 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1692)


if True == tmp8918 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp8904 := MakeNative(func(__e *ControlFlow) {
W1693 := __e.Get(1)
_ = W1693
tmp8905 := MakeNative(func(__e *ControlFlow) {
W1694 := __e.Get(1)
_ = W1694
tmp8906 := MakeNative(func(__e *ControlFlow) {
W1695 := __e.Get(1)
_ = W1695
tmp8913 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1695)


if True == tmp8913 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp8907 := MakeNative(func(__e *ControlFlow) {
W1696 := __e.Get(1)
_ = W1696
tmp8908 := MakeNative(func(__e *ControlFlow) {
W1697 := __e.Get(1)
_ = W1697
tmp8909 := PrimCons(W1693, W1696)

__e.TailApply(PrimFunc(symshen_4comb), W1697, tmp8909)
return


}, 1)

tmp8910 := Call(__e, PrimFunc(symshen_4in_1_6), W1695)


__e.TailApply(tmp8908, tmp8910)
return


}, 1)

tmp8911 := Call(__e, PrimFunc(symshen_4_5_1out), W1695)


__e.TailApply(tmp8907, tmp8911)
return


}


}, 1)

tmp8914 := Call(__e, PrimFunc(symshen_4_5head_6), W1694)


__e.TailApply(tmp8906, tmp8914)
return


}, 1)

tmp8915 := Call(__e, PrimFunc(symshen_4in_1_6), W1692)


__e.TailApply(tmp8905, tmp8915)
return


}, 1)

tmp8916 := Call(__e, PrimFunc(symshen_4_5_1out), W1692)


__e.TailApply(tmp8904, tmp8916)
return


}


}, 1)

tmp8919 := Call(__e, PrimFunc(symshen_4_5hterm_6), V1690)


tmp8920 := Call(__e, tmp8903, tmp8919)


__e.TailApply(tmp8890, tmp8920)
return


}, 1)

tmp8921 := Call(__e, ns2_1set, symshen_4_5head_6, tmp8889)


_ = tmp8921

tmp8922 := MakeNative(func(__e *ControlFlow) {
V1701 := __e.Get(1)
_ = V1701
tmp8923 := MakeNative(func(__e *ControlFlow) {
W1702 := __e.Get(1)
_ = W1702
tmp9111 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1702)


if True == tmp9111 {
tmp8924 := MakeNative(func(__e *ControlFlow) {
W1705 := __e.Get(1)
_ = W1705
tmp9098 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1705)


if True == tmp9098 {
tmp8925 := MakeNative(func(__e *ControlFlow) {
W1708 := __e.Get(1)
_ = W1708
tmp9059 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1708)


if True == tmp9059 {
tmp8926 := MakeNative(func(__e *ControlFlow) {
W1720 := __e.Get(1)
_ = W1720
tmp9029 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1720)


if True == tmp9029 {
tmp8927 := MakeNative(func(__e *ControlFlow) {
W1729 := __e.Get(1)
_ = W1729
tmp8999 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1729)


if True == tmp8999 {
tmp8928 := MakeNative(func(__e *ControlFlow) {
W1738 := __e.Get(1)
_ = W1738
tmp8965 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1738)


if True == tmp8965 {
tmp8929 := MakeNative(func(__e *ControlFlow) {
W1748 := __e.Get(1)
_ = W1748
tmp8931 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1748)


if True == tmp8931 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W1748)
return
}


}, 1)

tmp8963 := Call(__e, PrimFunc(symshen_4ccons_2), V1701)


var ifres8932 Obj

if True == tmp8963 {
tmp8933 := MakeNative(func(__e *ControlFlow) {
W1749 := __e.Get(1)
_ = W1749
tmp8934 := MakeNative(func(__e *ControlFlow) {
W1750 := __e.Get(1)
_ = W1750
tmp8958 := Call(__e, PrimFunc(symshen_4hds_a_2), W1749, symmode)


if True == tmp8958 {
tmp8935 := MakeNative(func(__e *ControlFlow) {
W1751 := __e.Get(1)
_ = W1751
tmp8936 := MakeNative(func(__e *ControlFlow) {
W1752 := __e.Get(1)
_ = W1752
tmp8954 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1752)


if True == tmp8954 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp8937 := MakeNative(func(__e *ControlFlow) {
W1753 := __e.Get(1)
_ = W1753
tmp8938 := MakeNative(func(__e *ControlFlow) {
W1754 := __e.Get(1)
_ = W1754
tmp8950 := Call(__e, PrimFunc(symshen_4hds_a_2), W1754, sym_1)


if True == tmp8950 {
tmp8939 := MakeNative(func(__e *ControlFlow) {
W1755 := __e.Get(1)
_ = W1755
tmp8940 := MakeNative(func(__e *ControlFlow) {
W1756 := __e.Get(1)
_ = W1756
tmp8946 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1756)


if True == tmp8946 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp8941 := MakeNative(func(__e *ControlFlow) {
W1757 := __e.Get(1)
_ = W1757
tmp8942 := PrimCons(W1753, Nil)

tmp8943 := PrimCons(symshen_4_1m, tmp8942)

__e.TailApply(PrimFunc(symshen_4comb), W1750, tmp8943)
return


}, 1)

tmp8944 := Call(__e, PrimFunc(symshen_4in_1_6), W1756)


__e.TailApply(tmp8941, tmp8944)
return


}


}, 1)

tmp8947 := Call(__e, PrimFunc(sym_5end_6), W1755)


__e.TailApply(tmp8940, tmp8947)
return


}, 1)

tmp8948 := Call(__e, PrimFunc(symtail), W1754)


__e.TailApply(tmp8939, tmp8948)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp8951 := Call(__e, PrimFunc(symshen_4in_1_6), W1752)


__e.TailApply(tmp8938, tmp8951)
return


}, 1)

tmp8952 := Call(__e, PrimFunc(symshen_4_5_1out), W1752)


__e.TailApply(tmp8937, tmp8952)
return


}


}, 1)

tmp8955 := Call(__e, PrimFunc(symshen_4_5hterm_6), W1751)


__e.TailApply(tmp8936, tmp8955)
return


}, 1)

tmp8956 := Call(__e, PrimFunc(symtail), W1749)


__e.TailApply(tmp8935, tmp8956)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp8959 := Call(__e, PrimFunc(symtail), V1701)


__e.TailApply(tmp8934, tmp8959)
return


}, 1)

tmp8960 := Call(__e, PrimFunc(symhead), V1701)


tmp8961 := Call(__e, tmp8933, tmp8960)


ifres8932 = tmp8961


} else {
tmp8962 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres8932 = tmp8962


}

__e.TailApply(tmp8929, ifres8932)
return


} else {
__e.Return(W1738)
return
}


}, 1)

tmp8997 := Call(__e, PrimFunc(symshen_4ccons_2), V1701)


var ifres8966 Obj

if True == tmp8997 {
tmp8967 := MakeNative(func(__e *ControlFlow) {
W1739 := __e.Get(1)
_ = W1739
tmp8968 := MakeNative(func(__e *ControlFlow) {
W1740 := __e.Get(1)
_ = W1740
tmp8992 := Call(__e, PrimFunc(symshen_4hds_a_2), W1739, symmode)


if True == tmp8992 {
tmp8969 := MakeNative(func(__e *ControlFlow) {
W1741 := __e.Get(1)
_ = W1741
tmp8970 := MakeNative(func(__e *ControlFlow) {
W1742 := __e.Get(1)
_ = W1742
tmp8988 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1742)


if True == tmp8988 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp8971 := MakeNative(func(__e *ControlFlow) {
W1743 := __e.Get(1)
_ = W1743
tmp8972 := MakeNative(func(__e *ControlFlow) {
W1744 := __e.Get(1)
_ = W1744
tmp8984 := Call(__e, PrimFunc(symshen_4hds_a_2), W1744, sym_7)


if True == tmp8984 {
tmp8973 := MakeNative(func(__e *ControlFlow) {
W1745 := __e.Get(1)
_ = W1745
tmp8974 := MakeNative(func(__e *ControlFlow) {
W1746 := __e.Get(1)
_ = W1746
tmp8980 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1746)


if True == tmp8980 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp8975 := MakeNative(func(__e *ControlFlow) {
W1747 := __e.Get(1)
_ = W1747
tmp8976 := PrimCons(W1743, Nil)

tmp8977 := PrimCons(symshen_4_7m, tmp8976)

__e.TailApply(PrimFunc(symshen_4comb), W1740, tmp8977)
return


}, 1)

tmp8978 := Call(__e, PrimFunc(symshen_4in_1_6), W1746)


__e.TailApply(tmp8975, tmp8978)
return


}


}, 1)

tmp8981 := Call(__e, PrimFunc(sym_5end_6), W1745)


__e.TailApply(tmp8974, tmp8981)
return


}, 1)

tmp8982 := Call(__e, PrimFunc(symtail), W1744)


__e.TailApply(tmp8973, tmp8982)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp8985 := Call(__e, PrimFunc(symshen_4in_1_6), W1742)


__e.TailApply(tmp8972, tmp8985)
return


}, 1)

tmp8986 := Call(__e, PrimFunc(symshen_4_5_1out), W1742)


__e.TailApply(tmp8971, tmp8986)
return


}


}, 1)

tmp8989 := Call(__e, PrimFunc(symshen_4_5hterm_6), W1741)


__e.TailApply(tmp8970, tmp8989)
return


}, 1)

tmp8990 := Call(__e, PrimFunc(symtail), W1739)


__e.TailApply(tmp8969, tmp8990)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp8993 := Call(__e, PrimFunc(symtail), V1701)


__e.TailApply(tmp8968, tmp8993)
return


}, 1)

tmp8994 := Call(__e, PrimFunc(symhead), V1701)


tmp8995 := Call(__e, tmp8967, tmp8994)


ifres8966 = tmp8995


} else {
tmp8996 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres8966 = tmp8996


}

__e.TailApply(tmp8928, ifres8966)
return


} else {
__e.Return(W1729)
return
}


}, 1)

tmp9027 := Call(__e, PrimFunc(symshen_4ccons_2), V1701)


var ifres9000 Obj

if True == tmp9027 {
tmp9001 := MakeNative(func(__e *ControlFlow) {
W1730 := __e.Get(1)
_ = W1730
tmp9002 := MakeNative(func(__e *ControlFlow) {
W1731 := __e.Get(1)
_ = W1731
tmp9022 := Call(__e, PrimFunc(symshen_4hds_a_2), W1730, sym_1)


if True == tmp9022 {
tmp9003 := MakeNative(func(__e *ControlFlow) {
W1732 := __e.Get(1)
_ = W1732
tmp9004 := MakeNative(func(__e *ControlFlow) {
W1733 := __e.Get(1)
_ = W1733
tmp9018 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1733)


if True == tmp9018 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp9005 := MakeNative(func(__e *ControlFlow) {
W1734 := __e.Get(1)
_ = W1734
tmp9006 := MakeNative(func(__e *ControlFlow) {
W1735 := __e.Get(1)
_ = W1735
tmp9007 := MakeNative(func(__e *ControlFlow) {
W1736 := __e.Get(1)
_ = W1736
tmp9013 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1736)


if True == tmp9013 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp9008 := MakeNative(func(__e *ControlFlow) {
W1737 := __e.Get(1)
_ = W1737
tmp9009 := PrimCons(W1734, Nil)

tmp9010 := PrimCons(symshen_4_1m, tmp9009)

__e.TailApply(PrimFunc(symshen_4comb), W1731, tmp9010)
return


}, 1)

tmp9011 := Call(__e, PrimFunc(symshen_4in_1_6), W1736)


__e.TailApply(tmp9008, tmp9011)
return


}


}, 1)

tmp9014 := Call(__e, PrimFunc(sym_5end_6), W1735)


__e.TailApply(tmp9007, tmp9014)
return


}, 1)

tmp9015 := Call(__e, PrimFunc(symshen_4in_1_6), W1733)


__e.TailApply(tmp9006, tmp9015)
return


}, 1)

tmp9016 := Call(__e, PrimFunc(symshen_4_5_1out), W1733)


__e.TailApply(tmp9005, tmp9016)
return


}


}, 1)

tmp9019 := Call(__e, PrimFunc(symshen_4_5hterm_6), W1732)


__e.TailApply(tmp9004, tmp9019)
return


}, 1)

tmp9020 := Call(__e, PrimFunc(symtail), W1730)


__e.TailApply(tmp9003, tmp9020)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp9023 := Call(__e, PrimFunc(symtail), V1701)


__e.TailApply(tmp9002, tmp9023)
return


}, 1)

tmp9024 := Call(__e, PrimFunc(symhead), V1701)


tmp9025 := Call(__e, tmp9001, tmp9024)


ifres9000 = tmp9025


} else {
tmp9026 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres9000 = tmp9026


}

__e.TailApply(tmp8927, ifres9000)
return


} else {
__e.Return(W1720)
return
}


}, 1)

tmp9057 := Call(__e, PrimFunc(symshen_4ccons_2), V1701)


var ifres9030 Obj

if True == tmp9057 {
tmp9031 := MakeNative(func(__e *ControlFlow) {
W1721 := __e.Get(1)
_ = W1721
tmp9032 := MakeNative(func(__e *ControlFlow) {
W1722 := __e.Get(1)
_ = W1722
tmp9052 := Call(__e, PrimFunc(symshen_4hds_a_2), W1721, sym_7)


if True == tmp9052 {
tmp9033 := MakeNative(func(__e *ControlFlow) {
W1723 := __e.Get(1)
_ = W1723
tmp9034 := MakeNative(func(__e *ControlFlow) {
W1724 := __e.Get(1)
_ = W1724
tmp9048 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1724)


if True == tmp9048 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp9035 := MakeNative(func(__e *ControlFlow) {
W1725 := __e.Get(1)
_ = W1725
tmp9036 := MakeNative(func(__e *ControlFlow) {
W1726 := __e.Get(1)
_ = W1726
tmp9037 := MakeNative(func(__e *ControlFlow) {
W1727 := __e.Get(1)
_ = W1727
tmp9043 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1727)


if True == tmp9043 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp9038 := MakeNative(func(__e *ControlFlow) {
W1728 := __e.Get(1)
_ = W1728
tmp9039 := PrimCons(W1725, Nil)

tmp9040 := PrimCons(symshen_4_7m, tmp9039)

__e.TailApply(PrimFunc(symshen_4comb), W1722, tmp9040)
return


}, 1)

tmp9041 := Call(__e, PrimFunc(symshen_4in_1_6), W1727)


__e.TailApply(tmp9038, tmp9041)
return


}


}, 1)

tmp9044 := Call(__e, PrimFunc(sym_5end_6), W1726)


__e.TailApply(tmp9037, tmp9044)
return


}, 1)

tmp9045 := Call(__e, PrimFunc(symshen_4in_1_6), W1724)


__e.TailApply(tmp9036, tmp9045)
return


}, 1)

tmp9046 := Call(__e, PrimFunc(symshen_4_5_1out), W1724)


__e.TailApply(tmp9035, tmp9046)
return


}


}, 1)

tmp9049 := Call(__e, PrimFunc(symshen_4_5hterm_6), W1723)


__e.TailApply(tmp9034, tmp9049)
return


}, 1)

tmp9050 := Call(__e, PrimFunc(symtail), W1721)


__e.TailApply(tmp9033, tmp9050)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp9053 := Call(__e, PrimFunc(symtail), V1701)


__e.TailApply(tmp9032, tmp9053)
return


}, 1)

tmp9054 := Call(__e, PrimFunc(symhead), V1701)


tmp9055 := Call(__e, tmp9031, tmp9054)


ifres9030 = tmp9055


} else {
tmp9056 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres9030 = tmp9056


}

__e.TailApply(tmp8926, ifres9030)
return


} else {
__e.Return(W1708)
return
}


}, 1)

tmp9096 := Call(__e, PrimFunc(symshen_4ccons_2), V1701)


var ifres9060 Obj

if True == tmp9096 {
tmp9061 := MakeNative(func(__e *ControlFlow) {
W1709 := __e.Get(1)
_ = W1709
tmp9062 := MakeNative(func(__e *ControlFlow) {
W1710 := __e.Get(1)
_ = W1710
tmp9091 := Call(__e, PrimFunc(symshen_4hds_a_2), W1709, symcons)


if True == tmp9091 {
tmp9063 := MakeNative(func(__e *ControlFlow) {
W1711 := __e.Get(1)
_ = W1711
tmp9064 := MakeNative(func(__e *ControlFlow) {
W1712 := __e.Get(1)
_ = W1712
tmp9087 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1712)


if True == tmp9087 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp9065 := MakeNative(func(__e *ControlFlow) {
W1713 := __e.Get(1)
_ = W1713
tmp9066 := MakeNative(func(__e *ControlFlow) {
W1714 := __e.Get(1)
_ = W1714
tmp9067 := MakeNative(func(__e *ControlFlow) {
W1715 := __e.Get(1)
_ = W1715
tmp9082 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1715)


if True == tmp9082 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp9068 := MakeNative(func(__e *ControlFlow) {
W1716 := __e.Get(1)
_ = W1716
tmp9069 := MakeNative(func(__e *ControlFlow) {
W1717 := __e.Get(1)
_ = W1717
tmp9070 := MakeNative(func(__e *ControlFlow) {
W1718 := __e.Get(1)
_ = W1718
tmp9077 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1718)


if True == tmp9077 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp9071 := MakeNative(func(__e *ControlFlow) {
W1719 := __e.Get(1)
_ = W1719
tmp9072 := PrimCons(W1716, Nil)

tmp9073 := PrimCons(W1713, tmp9072)

tmp9074 := PrimCons(symcons, tmp9073)

__e.TailApply(PrimFunc(symshen_4comb), W1710, tmp9074)
return


}, 1)

tmp9075 := Call(__e, PrimFunc(symshen_4in_1_6), W1718)


__e.TailApply(tmp9071, tmp9075)
return


}


}, 1)

tmp9078 := Call(__e, PrimFunc(sym_5end_6), W1717)


__e.TailApply(tmp9070, tmp9078)
return


}, 1)

tmp9079 := Call(__e, PrimFunc(symshen_4in_1_6), W1715)


__e.TailApply(tmp9069, tmp9079)
return


}, 1)

tmp9080 := Call(__e, PrimFunc(symshen_4_5_1out), W1715)


__e.TailApply(tmp9068, tmp9080)
return


}


}, 1)

tmp9083 := Call(__e, PrimFunc(symshen_4_5hterm2_6), W1714)


__e.TailApply(tmp9067, tmp9083)
return


}, 1)

tmp9084 := Call(__e, PrimFunc(symshen_4in_1_6), W1712)


__e.TailApply(tmp9066, tmp9084)
return


}, 1)

tmp9085 := Call(__e, PrimFunc(symshen_4_5_1out), W1712)


__e.TailApply(tmp9065, tmp9085)
return


}


}, 1)

tmp9088 := Call(__e, PrimFunc(symshen_4_5hterm1_6), W1711)


__e.TailApply(tmp9064, tmp9088)
return


}, 1)

tmp9089 := Call(__e, PrimFunc(symtail), W1709)


__e.TailApply(tmp9063, tmp9089)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp9092 := Call(__e, PrimFunc(symtail), V1701)


__e.TailApply(tmp9062, tmp9092)
return


}, 1)

tmp9093 := Call(__e, PrimFunc(symhead), V1701)


tmp9094 := Call(__e, tmp9061, tmp9093)


ifres9060 = tmp9094


} else {
tmp9095 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres9060 = tmp9095


}

__e.TailApply(tmp8925, ifres9060)
return


} else {
__e.Return(W1705)
return
}


}, 1)

tmp9109 := PrimIsPair(V1701)

var ifres9099 Obj

if True == tmp9109 {
tmp9100 := MakeNative(func(__e *ControlFlow) {
W1706 := __e.Get(1)
_ = W1706
tmp9101 := MakeNative(func(__e *ControlFlow) {
W1707 := __e.Get(1)
_ = W1707
tmp9103 := PrimIntern(MakeString(":"))

tmp9104 := PrimEqual(W1706, tmp9103)

if True == tmp9104 {
__e.TailApply(PrimFunc(symshen_4comb), W1707, W1706)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp9105 := Call(__e, PrimFunc(symtail), V1701)


__e.TailApply(tmp9101, tmp9105)
return


}, 1)

tmp9106 := Call(__e, PrimFunc(symhead), V1701)


tmp9107 := Call(__e, tmp9100, tmp9106)


ifres9099 = tmp9107


} else {
tmp9108 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres9099 = tmp9108


}

__e.TailApply(tmp8924, ifres9099)
return


} else {
__e.Return(W1702)
return
}


}, 1)

tmp9125 := PrimIsPair(V1701)

var ifres9112 Obj

if True == tmp9125 {
tmp9113 := MakeNative(func(__e *ControlFlow) {
W1703 := __e.Get(1)
_ = W1703
tmp9114 := MakeNative(func(__e *ControlFlow) {
W1704 := __e.Get(1)
_ = W1704
tmp9120 := Call(__e, PrimFunc(symatom_2), W1703)


var ifres9116 Obj

if True == tmp9120 {
tmp9118 := Call(__e, PrimFunc(symshen_4prolog_1keyword_2), W1703)


tmp9119 := PrimNot(tmp9118)

var ifres9117 Obj

if True == tmp9119 {
ifres9117 = True


} else {
ifres9117 = False


}

ifres9116 = ifres9117


} else {
ifres9116 = False


}

if True == ifres9116 {
__e.TailApply(PrimFunc(symshen_4comb), W1704, W1703)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp9121 := Call(__e, PrimFunc(symtail), V1701)


__e.TailApply(tmp9114, tmp9121)
return


}, 1)

tmp9122 := Call(__e, PrimFunc(symhead), V1701)


tmp9123 := Call(__e, tmp9113, tmp9122)


ifres9112 = tmp9123


} else {
tmp9124 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres9112 = tmp9124


}

__e.TailApply(tmp8923, ifres9112)
return


}, 1)

tmp9126 := Call(__e, ns2_1set, symshen_4_5hterm_6, tmp8922)


_ = tmp9126

tmp9127 := MakeNative(func(__e *ControlFlow) {
V1758 := __e.Get(1)
_ = V1758
tmp9128 := PrimIntern(MakeString(";"))

tmp9129 := PrimCons(sym_5_1_1, Nil)

tmp9130 := PrimCons(tmp9128, tmp9129)

__e.TailApply(PrimFunc(symelement_2), V1758, tmp9130)
return


}, 1)

tmp9131 := Call(__e, ns2_1set, symshen_4prolog_1keyword_2, tmp9127)


_ = tmp9131

tmp9132 := MakeNative(func(__e *ControlFlow) {
V1759 := __e.Get(1)
_ = V1759
tmp9145 := PrimIsSymbol(V1759)

if True == tmp9145 {
__e.Return(True)
return
} else {
tmp9143 := PrimIsString(V1759)

var ifres9134 Obj

if True == tmp9143 {
ifres9134 = True


} else {
tmp9142 := Call(__e, PrimFunc(symboolean_2), V1759)


var ifres9136 Obj

if True == tmp9142 {
ifres9136 = True


} else {
tmp9141 := PrimIsNumber(V1759)

var ifres9138 Obj

if True == tmp9141 {
ifres9138 = True


} else {
tmp9140 := Call(__e, PrimFunc(symempty_2), V1759)


var ifres9139 Obj

if True == tmp9140 {
ifres9139 = True


} else {
ifres9139 = False


}

ifres9138 = ifres9139


}

var ifres9137 Obj

if True == ifres9138 {
ifres9137 = True


} else {
ifres9137 = False


}

ifres9136 = ifres9137


}

var ifres9135 Obj

if True == ifres9136 {
ifres9135 = True


} else {
ifres9135 = False


}

ifres9134 = ifres9135


}

if True == ifres9134 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp9146 := Call(__e, ns2_1set, symatom_2, tmp9132)


_ = tmp9146

tmp9147 := MakeNative(func(__e *ControlFlow) {
V1760 := __e.Get(1)
_ = V1760
tmp9148 := MakeNative(func(__e *ControlFlow) {
W1761 := __e.Get(1)
_ = W1761
tmp9150 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1761)


if True == tmp9150 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W1761)
return
}


}, 1)

tmp9151 := MakeNative(func(__e *ControlFlow) {
W1762 := __e.Get(1)
_ = W1762
tmp9157 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1762)


if True == tmp9157 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp9152 := MakeNative(func(__e *ControlFlow) {
W1763 := __e.Get(1)
_ = W1763
tmp9153 := MakeNative(func(__e *ControlFlow) {
W1764 := __e.Get(1)
_ = W1764
__e.TailApply(PrimFunc(symshen_4comb), W1764, W1763)
return
}, 1)

tmp9154 := Call(__e, PrimFunc(symshen_4in_1_6), W1762)


__e.TailApply(tmp9153, tmp9154)
return


}, 1)

tmp9155 := Call(__e, PrimFunc(symshen_4_5_1out), W1762)


__e.TailApply(tmp9152, tmp9155)
return


}


}, 1)

tmp9158 := Call(__e, PrimFunc(symshen_4_5hterm_6), V1760)


tmp9159 := Call(__e, tmp9151, tmp9158)


__e.TailApply(tmp9148, tmp9159)
return


}, 1)

tmp9160 := Call(__e, ns2_1set, symshen_4_5hterm1_6, tmp9147)


_ = tmp9160

tmp9161 := MakeNative(func(__e *ControlFlow) {
V1765 := __e.Get(1)
_ = V1765
tmp9162 := MakeNative(func(__e *ControlFlow) {
W1766 := __e.Get(1)
_ = W1766
tmp9164 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1766)


if True == tmp9164 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W1766)
return
}


}, 1)

tmp9165 := MakeNative(func(__e *ControlFlow) {
W1767 := __e.Get(1)
_ = W1767
tmp9171 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1767)


if True == tmp9171 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp9166 := MakeNative(func(__e *ControlFlow) {
W1768 := __e.Get(1)
_ = W1768
tmp9167 := MakeNative(func(__e *ControlFlow) {
W1769 := __e.Get(1)
_ = W1769
__e.TailApply(PrimFunc(symshen_4comb), W1769, W1768)
return
}, 1)

tmp9168 := Call(__e, PrimFunc(symshen_4in_1_6), W1767)


__e.TailApply(tmp9167, tmp9168)
return


}, 1)

tmp9169 := Call(__e, PrimFunc(symshen_4_5_1out), W1767)


__e.TailApply(tmp9166, tmp9169)
return


}


}, 1)

tmp9172 := Call(__e, PrimFunc(symshen_4_5hterm_6), V1765)


tmp9173 := Call(__e, tmp9165, tmp9172)


__e.TailApply(tmp9162, tmp9173)
return


}, 1)

tmp9174 := Call(__e, ns2_1set, symshen_4_5hterm2_6, tmp9161)


_ = tmp9174

tmp9175 := MakeNative(func(__e *ControlFlow) {
V1770 := __e.Get(1)
_ = V1770
tmp9176 := MakeNative(func(__e *ControlFlow) {
W1771 := __e.Get(1)
_ = W1771
tmp9188 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1771)


if True == tmp9188 {
tmp9177 := MakeNative(func(__e *ControlFlow) {
W1778 := __e.Get(1)
_ = W1778
tmp9179 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1778)


if True == tmp9179 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W1778)
return
}


}, 1)

tmp9180 := MakeNative(func(__e *ControlFlow) {
W1779 := __e.Get(1)
_ = W1779
tmp9184 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1779)


if True == tmp9184 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp9181 := MakeNative(func(__e *ControlFlow) {
W1780 := __e.Get(1)
_ = W1780
__e.TailApply(PrimFunc(symshen_4comb), W1780, Nil)
return
}, 1)

tmp9182 := Call(__e, PrimFunc(symshen_4in_1_6), W1779)


__e.TailApply(tmp9181, tmp9182)
return


}


}, 1)

tmp9185 := Call(__e, PrimFunc(sym_5e_6), V1770)


tmp9186 := Call(__e, tmp9180, tmp9185)


__e.TailApply(tmp9177, tmp9186)
return


} else {
__e.Return(W1771)
return
}


}, 1)

tmp9189 := MakeNative(func(__e *ControlFlow) {
W1772 := __e.Get(1)
_ = W1772
tmp9204 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1772)


if True == tmp9204 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp9190 := MakeNative(func(__e *ControlFlow) {
W1773 := __e.Get(1)
_ = W1773
tmp9191 := MakeNative(func(__e *ControlFlow) {
W1774 := __e.Get(1)
_ = W1774
tmp9192 := MakeNative(func(__e *ControlFlow) {
W1775 := __e.Get(1)
_ = W1775
tmp9199 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1775)


if True == tmp9199 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp9193 := MakeNative(func(__e *ControlFlow) {
W1776 := __e.Get(1)
_ = W1776
tmp9194 := MakeNative(func(__e *ControlFlow) {
W1777 := __e.Get(1)
_ = W1777
tmp9195 := PrimCons(W1773, W1776)

__e.TailApply(PrimFunc(symshen_4comb), W1777, tmp9195)
return


}, 1)

tmp9196 := Call(__e, PrimFunc(symshen_4in_1_6), W1775)


__e.TailApply(tmp9194, tmp9196)
return


}, 1)

tmp9197 := Call(__e, PrimFunc(symshen_4_5_1out), W1775)


__e.TailApply(tmp9193, tmp9197)
return


}


}, 1)

tmp9200 := Call(__e, PrimFunc(symshen_4_5body_6), W1774)


__e.TailApply(tmp9192, tmp9200)
return


}, 1)

tmp9201 := Call(__e, PrimFunc(symshen_4in_1_6), W1772)


__e.TailApply(tmp9191, tmp9201)
return


}, 1)

tmp9202 := Call(__e, PrimFunc(symshen_4_5_1out), W1772)


__e.TailApply(tmp9190, tmp9202)
return


}


}, 1)

tmp9205 := Call(__e, PrimFunc(symshen_4_5literal_6), V1770)


tmp9206 := Call(__e, tmp9189, tmp9205)


__e.TailApply(tmp9176, tmp9206)
return


}, 1)

tmp9207 := Call(__e, ns2_1set, symshen_4_5body_6, tmp9175)


_ = tmp9207

tmp9208 := MakeNative(func(__e *ControlFlow) {
V1781 := __e.Get(1)
_ = V1781
tmp9209 := MakeNative(func(__e *ControlFlow) {
W1782 := __e.Get(1)
_ = W1782
tmp9236 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1782)


if True == tmp9236 {
tmp9210 := MakeNative(func(__e *ControlFlow) {
W1784 := __e.Get(1)
_ = W1784
tmp9212 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1784)


if True == tmp9212 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W1784)
return
}


}, 1)

tmp9234 := Call(__e, PrimFunc(symshen_4ccons_2), V1781)


var ifres9213 Obj

if True == tmp9234 {
tmp9214 := MakeNative(func(__e *ControlFlow) {
W1785 := __e.Get(1)
_ = W1785
tmp9215 := MakeNative(func(__e *ControlFlow) {
W1786 := __e.Get(1)
_ = W1786
tmp9216 := MakeNative(func(__e *ControlFlow) {
W1787 := __e.Get(1)
_ = W1787
tmp9228 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1787)


if True == tmp9228 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp9217 := MakeNative(func(__e *ControlFlow) {
W1788 := __e.Get(1)
_ = W1788
tmp9218 := MakeNative(func(__e *ControlFlow) {
W1789 := __e.Get(1)
_ = W1789
tmp9219 := MakeNative(func(__e *ControlFlow) {
W1790 := __e.Get(1)
_ = W1790
tmp9223 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1790)


if True == tmp9223 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp9220 := MakeNative(func(__e *ControlFlow) {
W1791 := __e.Get(1)
_ = W1791
__e.TailApply(PrimFunc(symshen_4comb), W1786, W1788)
return
}, 1)

tmp9221 := Call(__e, PrimFunc(symshen_4in_1_6), W1790)


__e.TailApply(tmp9220, tmp9221)
return


}


}, 1)

tmp9224 := Call(__e, PrimFunc(sym_5end_6), W1789)


__e.TailApply(tmp9219, tmp9224)
return


}, 1)

tmp9225 := Call(__e, PrimFunc(symshen_4in_1_6), W1787)


__e.TailApply(tmp9218, tmp9225)
return


}, 1)

tmp9226 := Call(__e, PrimFunc(symshen_4_5_1out), W1787)


__e.TailApply(tmp9217, tmp9226)
return


}


}, 1)

tmp9229 := Call(__e, PrimFunc(symshen_4_5bterms_6), W1785)


__e.TailApply(tmp9216, tmp9229)
return


}, 1)

tmp9230 := Call(__e, PrimFunc(symtail), V1781)


__e.TailApply(tmp9215, tmp9230)
return


}, 1)

tmp9231 := Call(__e, PrimFunc(symhead), V1781)


tmp9232 := Call(__e, tmp9214, tmp9231)


ifres9213 = tmp9232


} else {
tmp9233 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres9213 = tmp9233


}

__e.TailApply(tmp9210, ifres9213)
return


} else {
__e.Return(W1782)
return
}


}, 1)

tmp9242 := Call(__e, PrimFunc(symshen_4hds_a_2), V1781, sym_b)


var ifres9237 Obj

if True == tmp9242 {
tmp9238 := MakeNative(func(__e *ControlFlow) {
W1783 := __e.Get(1)
_ = W1783
__e.TailApply(PrimFunc(symshen_4comb), W1783, sym_b)
return
}, 1)

tmp9239 := Call(__e, PrimFunc(symtail), V1781)


tmp9240 := Call(__e, tmp9238, tmp9239)


ifres9237 = tmp9240


} else {
tmp9241 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres9237 = tmp9241


}

__e.TailApply(tmp9209, ifres9237)
return


}, 1)

tmp9243 := Call(__e, ns2_1set, symshen_4_5literal_6, tmp9208)


_ = tmp9243

tmp9244 := MakeNative(func(__e *ControlFlow) {
V1792 := __e.Get(1)
_ = V1792
tmp9245 := MakeNative(func(__e *ControlFlow) {
W1793 := __e.Get(1)
_ = W1793
tmp9257 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1793)


if True == tmp9257 {
tmp9246 := MakeNative(func(__e *ControlFlow) {
W1800 := __e.Get(1)
_ = W1800
tmp9248 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1800)


if True == tmp9248 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W1800)
return
}


}, 1)

tmp9249 := MakeNative(func(__e *ControlFlow) {
W1801 := __e.Get(1)
_ = W1801
tmp9253 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1801)


if True == tmp9253 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp9250 := MakeNative(func(__e *ControlFlow) {
W1802 := __e.Get(1)
_ = W1802
__e.TailApply(PrimFunc(symshen_4comb), W1802, Nil)
return
}, 1)

tmp9251 := Call(__e, PrimFunc(symshen_4in_1_6), W1801)


__e.TailApply(tmp9250, tmp9251)
return


}


}, 1)

tmp9254 := Call(__e, PrimFunc(sym_5e_6), V1792)


tmp9255 := Call(__e, tmp9249, tmp9254)


__e.TailApply(tmp9246, tmp9255)
return


} else {
__e.Return(W1793)
return
}


}, 1)

tmp9258 := MakeNative(func(__e *ControlFlow) {
W1794 := __e.Get(1)
_ = W1794
tmp9273 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1794)


if True == tmp9273 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp9259 := MakeNative(func(__e *ControlFlow) {
W1795 := __e.Get(1)
_ = W1795
tmp9260 := MakeNative(func(__e *ControlFlow) {
W1796 := __e.Get(1)
_ = W1796
tmp9261 := MakeNative(func(__e *ControlFlow) {
W1797 := __e.Get(1)
_ = W1797
tmp9268 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1797)


if True == tmp9268 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp9262 := MakeNative(func(__e *ControlFlow) {
W1798 := __e.Get(1)
_ = W1798
tmp9263 := MakeNative(func(__e *ControlFlow) {
W1799 := __e.Get(1)
_ = W1799
tmp9264 := PrimCons(W1795, W1798)

__e.TailApply(PrimFunc(symshen_4comb), W1799, tmp9264)
return


}, 1)

tmp9265 := Call(__e, PrimFunc(symshen_4in_1_6), W1797)


__e.TailApply(tmp9263, tmp9265)
return


}, 1)

tmp9266 := Call(__e, PrimFunc(symshen_4_5_1out), W1797)


__e.TailApply(tmp9262, tmp9266)
return


}


}, 1)

tmp9269 := Call(__e, PrimFunc(symshen_4_5bterms_6), W1796)


__e.TailApply(tmp9261, tmp9269)
return


}, 1)

tmp9270 := Call(__e, PrimFunc(symshen_4in_1_6), W1794)


__e.TailApply(tmp9260, tmp9270)
return


}, 1)

tmp9271 := Call(__e, PrimFunc(symshen_4_5_1out), W1794)


__e.TailApply(tmp9259, tmp9271)
return


}


}, 1)

tmp9274 := Call(__e, PrimFunc(symshen_4_5bterm_6), V1792)


tmp9275 := Call(__e, tmp9258, tmp9274)


__e.TailApply(tmp9245, tmp9275)
return


}, 1)

tmp9276 := Call(__e, ns2_1set, symshen_4_5bterms_6, tmp9244)


_ = tmp9276

tmp9277 := MakeNative(func(__e *ControlFlow) {
V1803 := __e.Get(1)
_ = V1803
tmp9278 := MakeNative(func(__e *ControlFlow) {
W1804 := __e.Get(1)
_ = W1804
tmp9318 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1804)


if True == tmp9318 {
tmp9279 := MakeNative(func(__e *ControlFlow) {
W1808 := __e.Get(1)
_ = W1808
tmp9306 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1808)


if True == tmp9306 {
tmp9280 := MakeNative(func(__e *ControlFlow) {
W1811 := __e.Get(1)
_ = W1811
tmp9282 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1811)


if True == tmp9282 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W1811)
return
}


}, 1)

tmp9304 := Call(__e, PrimFunc(symshen_4ccons_2), V1803)


var ifres9283 Obj

if True == tmp9304 {
tmp9284 := MakeNative(func(__e *ControlFlow) {
W1812 := __e.Get(1)
_ = W1812
tmp9285 := MakeNative(func(__e *ControlFlow) {
W1813 := __e.Get(1)
_ = W1813
tmp9286 := MakeNative(func(__e *ControlFlow) {
W1814 := __e.Get(1)
_ = W1814
tmp9298 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1814)


if True == tmp9298 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp9287 := MakeNative(func(__e *ControlFlow) {
W1815 := __e.Get(1)
_ = W1815
tmp9288 := MakeNative(func(__e *ControlFlow) {
W1816 := __e.Get(1)
_ = W1816
tmp9289 := MakeNative(func(__e *ControlFlow) {
W1817 := __e.Get(1)
_ = W1817
tmp9293 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1817)


if True == tmp9293 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp9290 := MakeNative(func(__e *ControlFlow) {
W1818 := __e.Get(1)
_ = W1818
__e.TailApply(PrimFunc(symshen_4comb), W1813, W1815)
return
}, 1)

tmp9291 := Call(__e, PrimFunc(symshen_4in_1_6), W1817)


__e.TailApply(tmp9290, tmp9291)
return


}


}, 1)

tmp9294 := Call(__e, PrimFunc(sym_5end_6), W1816)


__e.TailApply(tmp9289, tmp9294)
return


}, 1)

tmp9295 := Call(__e, PrimFunc(symshen_4in_1_6), W1814)


__e.TailApply(tmp9288, tmp9295)
return


}, 1)

tmp9296 := Call(__e, PrimFunc(symshen_4_5_1out), W1814)


__e.TailApply(tmp9287, tmp9296)
return


}


}, 1)

tmp9299 := Call(__e, PrimFunc(symshen_4_5bterms_6), W1812)


__e.TailApply(tmp9286, tmp9299)
return


}, 1)

tmp9300 := Call(__e, PrimFunc(symtail), V1803)


__e.TailApply(tmp9285, tmp9300)
return


}, 1)

tmp9301 := Call(__e, PrimFunc(symhead), V1803)


tmp9302 := Call(__e, tmp9284, tmp9301)


ifres9283 = tmp9302


} else {
tmp9303 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres9283 = tmp9303


}

__e.TailApply(tmp9280, ifres9283)
return


} else {
__e.Return(W1808)
return
}


}, 1)

tmp9316 := PrimIsPair(V1803)

var ifres9307 Obj

if True == tmp9316 {
tmp9308 := MakeNative(func(__e *ControlFlow) {
W1809 := __e.Get(1)
_ = W1809
tmp9309 := MakeNative(func(__e *ControlFlow) {
W1810 := __e.Get(1)
_ = W1810
tmp9311 := Call(__e, PrimFunc(symatom_2), W1809)


if True == tmp9311 {
__e.TailApply(PrimFunc(symshen_4comb), W1810, W1809)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp9312 := Call(__e, PrimFunc(symtail), V1803)


__e.TailApply(tmp9309, tmp9312)
return


}, 1)

tmp9313 := Call(__e, PrimFunc(symhead), V1803)


tmp9314 := Call(__e, tmp9308, tmp9313)


ifres9307 = tmp9314


} else {
tmp9315 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres9307 = tmp9315


}

__e.TailApply(tmp9279, ifres9307)
return


} else {
__e.Return(W1804)
return
}


}, 1)

tmp9319 := MakeNative(func(__e *ControlFlow) {
W1805 := __e.Get(1)
_ = W1805
tmp9325 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1805)


if True == tmp9325 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp9320 := MakeNative(func(__e *ControlFlow) {
W1806 := __e.Get(1)
_ = W1806
tmp9321 := MakeNative(func(__e *ControlFlow) {
W1807 := __e.Get(1)
_ = W1807
__e.TailApply(PrimFunc(symshen_4comb), W1807, W1806)
return
}, 1)

tmp9322 := Call(__e, PrimFunc(symshen_4in_1_6), W1805)


__e.TailApply(tmp9321, tmp9322)
return


}, 1)

tmp9323 := Call(__e, PrimFunc(symshen_4_5_1out), W1805)


__e.TailApply(tmp9320, tmp9323)
return


}


}, 1)

tmp9326 := Call(__e, PrimFunc(symshen_4_5wildcard_6), V1803)


tmp9327 := Call(__e, tmp9319, tmp9326)


__e.TailApply(tmp9278, tmp9327)
return


}, 1)

tmp9328 := Call(__e, ns2_1set, symshen_4_5bterm_6, tmp9277)


_ = tmp9328

tmp9329 := MakeNative(func(__e *ControlFlow) {
V1819 := __e.Get(1)
_ = V1819
tmp9330 := MakeNative(func(__e *ControlFlow) {
W1820 := __e.Get(1)
_ = W1820
tmp9332 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1820)


if True == tmp9332 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W1820)
return
}


}, 1)

tmp9343 := PrimIsPair(V1819)

var ifres9333 Obj

if True == tmp9343 {
tmp9334 := MakeNative(func(__e *ControlFlow) {
W1821 := __e.Get(1)
_ = W1821
tmp9335 := MakeNative(func(__e *ControlFlow) {
W1822 := __e.Get(1)
_ = W1822
tmp9338 := PrimEqual(W1821, sym__)

if True == tmp9338 {
tmp9336 := Call(__e, PrimFunc(symgensym), symY)


__e.TailApply(PrimFunc(symshen_4comb), W1822, tmp9336)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp9339 := Call(__e, PrimFunc(symtail), V1819)


__e.TailApply(tmp9335, tmp9339)
return


}, 1)

tmp9340 := Call(__e, PrimFunc(symhead), V1819)


tmp9341 := Call(__e, tmp9334, tmp9340)


ifres9333 = tmp9341


} else {
tmp9342 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres9333 = tmp9342


}

__e.TailApply(tmp9330, ifres9333)
return


}, 1)

tmp9344 := Call(__e, ns2_1set, symshen_4_5wildcard_6, tmp9329)


_ = tmp9344

tmp9345 := MakeNative(func(__e *ControlFlow) {
V1823 := __e.Get(1)
_ = V1823
tmp9346 := MakeNative(func(__e *ControlFlow) {
W1824 := __e.Get(1)
_ = W1824
tmp9348 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1824)


if True == tmp9348 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W1824)
return
}


}, 1)

tmp9358 := PrimIsPair(V1823)

var ifres9349 Obj

if True == tmp9358 {
tmp9350 := MakeNative(func(__e *ControlFlow) {
W1825 := __e.Get(1)
_ = W1825
tmp9351 := MakeNative(func(__e *ControlFlow) {
W1826 := __e.Get(1)
_ = W1826
tmp9353 := Call(__e, PrimFunc(symshen_4semicolon_2), W1825)


if True == tmp9353 {
__e.TailApply(PrimFunc(symshen_4comb), W1826, W1825)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp9354 := Call(__e, PrimFunc(symtail), V1823)


__e.TailApply(tmp9351, tmp9354)
return


}, 1)

tmp9355 := Call(__e, PrimFunc(symhead), V1823)


tmp9356 := Call(__e, tmp9350, tmp9355)


ifres9349 = tmp9356


} else {
tmp9357 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres9349 = tmp9357


}

__e.TailApply(tmp9346, ifres9349)
return


}, 1)

tmp9359 := Call(__e, ns2_1set, symshen_4_5sc_6, tmp9345)


_ = tmp9359

tmp9360 := MakeNative(func(__e *ControlFlow) {
V1827 := __e.Get(1)
_ = V1827
V1828 := __e.Get(2)
_ = V1828
tmp9361 := MakeNative(func(__e *ControlFlow) {
W1829 := __e.Get(1)
_ = W1829
tmp9362 := MakeNative(func(__e *ControlFlow) {
W1830 := __e.Get(1)
_ = W1830
tmp9363 := MakeNative(func(__e *ControlFlow) {
W1831 := __e.Get(1)
_ = W1831
tmp9364 := MakeNative(func(__e *ControlFlow) {
W1832 := __e.Get(1)
_ = W1832
tmp9365 := MakeNative(func(__e *ControlFlow) {
W1833 := __e.Get(1)
_ = W1833
tmp9366 := MakeNative(func(__e *ControlFlow) {
W1834 := __e.Get(1)
_ = W1834
tmp9367 := MakeNative(func(__e *ControlFlow) {
W1835 := __e.Get(1)
_ = W1835
tmp9368 := MakeNative(func(__e *ControlFlow) {
W1836 := __e.Get(1)
_ = W1836
tmp9369 := MakeNative(func(__e *ControlFlow) {
W1837 := __e.Get(1)
_ = W1837
__e.Return(W1837)
return
}, 1)

tmp9370 := PrimCons(sym_1_6, Nil)

tmp9371 := PrimCons(W1832, tmp9370)

tmp9372 := PrimCons(W1831, tmp9371)

tmp9373 := PrimCons(W1830, tmp9372)

tmp9374 := PrimCons(W1829, tmp9373)

tmp9375 := PrimCons(W1836, Nil)

tmp9376 := Call(__e, PrimFunc(symappend), tmp9374, tmp9375)


tmp9377 := Call(__e, PrimFunc(symappend), W1833, tmp9376)


tmp9378 := PrimCons(V1827, tmp9377)

tmp9379 := PrimCons(symdefine, tmp9378)

__e.TailApply(tmp9369, tmp9379)
return


}, 1)

var ifres9380 Obj

if True == W1834 {
tmp9381 := PrimCons(MakeNumber(1), Nil)

tmp9382 := PrimCons(W1831, tmp9381)

tmp9383 := PrimCons(sym_7, tmp9382)

tmp9384 := PrimCons(W1835, Nil)

tmp9385 := PrimCons(tmp9383, tmp9384)

tmp9386 := PrimCons(W1831, tmp9385)

tmp9387 := PrimCons(symlet, tmp9386)

ifres9380 = tmp9387


} else {
ifres9380 = W1835


}

__e.TailApply(tmp9368, ifres9380)
return


}, 1)

tmp9388 := Call(__e, PrimFunc(symshen_4prolog_1fbody), V1828, W1833, W1829, W1830, W1831, W1832, W1834)


__e.TailApply(tmp9367, tmp9388)
return


}, 1)

tmp9389 := Call(__e, PrimFunc(symshen_4hascut_2), V1828)


__e.TailApply(tmp9366, tmp9389)
return


}, 1)

tmp9390 := Call(__e, PrimFunc(symshen_4prolog_1parameters), V1828)


__e.TailApply(tmp9365, tmp9390)
return


}, 1)

tmp9391 := Call(__e, PrimFunc(symgensym), symC)


__e.TailApply(tmp9364, tmp9391)
return


}, 1)

tmp9392 := Call(__e, PrimFunc(symgensym), symK)


__e.TailApply(tmp9363, tmp9392)
return


}, 1)

tmp9393 := Call(__e, PrimFunc(symgensym), symL)


__e.TailApply(tmp9362, tmp9393)
return


}, 1)

tmp9394 := Call(__e, PrimFunc(symgensym), symB)


__e.TailApply(tmp9361, tmp9394)
return


}, 2)

tmp9395 := Call(__e, ns2_1set, symshen_4horn_1clause_1procedure, tmp9360)


_ = tmp9395

tmp9396 := MakeNative(func(__e *ControlFlow) {
V1840 := __e.Get(1)
_ = V1840
tmp9406 := PrimEqual(sym_b, V1840)

if True == tmp9406 {
__e.Return(True)
return
} else {
tmp9404 := PrimIsPair(V1840)

if True == tmp9404 {
tmp9401 := PrimHead(V1840)

tmp9402 := Call(__e, PrimFunc(symshen_4hascut_2), tmp9401)


if True == tmp9402 {
__e.Return(True)
return
} else {
tmp9398 := PrimTail(V1840)

tmp9399 := Call(__e, PrimFunc(symshen_4hascut_2), tmp9398)


if True == tmp9399 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


} else {
__e.Return(False)
return
}


}


}, 1)

tmp9407 := Call(__e, ns2_1set, symshen_4hascut_2, tmp9396)


_ = tmp9407

tmp9408 := MakeNative(func(__e *ControlFlow) {
V1845 := __e.Get(1)
_ = V1845
tmp9417 := PrimIsPair(V1845)

var ifres9413 Obj

if True == tmp9417 {
tmp9415 := PrimHead(V1845)

tmp9416 := PrimIsPair(tmp9415)

var ifres9414 Obj

if True == tmp9416 {
ifres9414 = True


} else {
ifres9414 = False


}

ifres9413 = ifres9414


} else {
ifres9413 = False


}

if True == ifres9413 {
tmp9409 := PrimHead(V1845)

tmp9410 := PrimHead(tmp9409)

tmp9411 := Call(__e, PrimFunc(symlength), tmp9410)


__e.TailApply(PrimFunc(symshen_4parameters), tmp9411)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4prolog_1parameters)
return
}


}, 1)

tmp9418 := Call(__e, ns2_1set, symshen_4prolog_1parameters, tmp9408)


_ = tmp9418

tmp9419 := MakeNative(func(__e *ControlFlow) {
V1866 := __e.Get(1)
_ = V1866
V1867 := __e.Get(2)
_ = V1867
V1868 := __e.Get(3)
_ = V1868
V1869 := __e.Get(4)
_ = V1869
V1870 := __e.Get(5)
_ = V1870
V1871 := __e.Get(6)
_ = V1871
V1872 := __e.Get(7)
_ = V1872
tmp9512 := PrimEqual(Nil, V1866)

var ifres9509 Obj

if True == tmp9512 {
tmp9511 := PrimEqual(True, V1872)

var ifres9510 Obj

if True == tmp9511 {
ifres9510 = True


} else {
ifres9510 = False


}

ifres9509 = ifres9510


} else {
ifres9509 = False


}

if True == ifres9509 {
tmp9420 := PrimCons(V1870, Nil)

tmp9421 := PrimCons(V1869, tmp9420)

__e.Return(PrimCons(symshen_4unlock, tmp9421))
return


} else {
tmp9507 := PrimIsPair(V1866)

var ifres9485 Obj

if True == tmp9507 {
tmp9505 := PrimHead(V1866)

tmp9506 := PrimIsPair(tmp9505)

var ifres9487 Obj

if True == tmp9506 {
tmp9502 := PrimHead(V1866)

tmp9503 := PrimTail(tmp9502)

tmp9504 := PrimIsPair(tmp9503)

var ifres9489 Obj

if True == tmp9504 {
tmp9498 := PrimHead(V1866)

tmp9499 := PrimTail(tmp9498)

tmp9500 := PrimTail(tmp9499)

tmp9501 := PrimEqual(Nil, tmp9500)

var ifres9491 Obj

if True == tmp9501 {
tmp9496 := PrimTail(V1866)

tmp9497 := PrimEqual(Nil, tmp9496)

var ifres9493 Obj

if True == tmp9497 {
tmp9495 := PrimEqual(False, V1872)

var ifres9494 Obj

if True == tmp9495 {
ifres9494 = True


} else {
ifres9494 = False


}

ifres9493 = ifres9494


} else {
ifres9493 = False


}

var ifres9492 Obj

if True == ifres9493 {
ifres9492 = True


} else {
ifres9492 = False


}

ifres9491 = ifres9492


} else {
ifres9491 = False


}

var ifres9490 Obj

if True == ifres9491 {
ifres9490 = True


} else {
ifres9490 = False


}

ifres9489 = ifres9490


} else {
ifres9489 = False


}

var ifres9488 Obj

if True == ifres9489 {
ifres9488 = True


} else {
ifres9488 = False


}

ifres9487 = ifres9488


} else {
ifres9487 = False


}

var ifres9486 Obj

if True == ifres9487 {
ifres9486 = True


} else {
ifres9486 = False


}

ifres9485 = ifres9486


} else {
ifres9485 = False


}

if True == ifres9485 {
tmp9422 := MakeNative(func(__e *ControlFlow) {
W1873 := __e.Get(1)
_ = W1873
tmp9423 := PrimCons(V1869, Nil)

tmp9424 := PrimCons(symshen_4unlocked_2, tmp9423)

tmp9425 := PrimHead(V1866)

tmp9426 := PrimHead(tmp9425)

tmp9427 := Call(__e, PrimFunc(symshen_4compile_1head), symshen_4_7m, tmp9426, V1867, V1868, W1873)


tmp9428 := PrimCons(False, Nil)

tmp9429 := PrimCons(tmp9427, tmp9428)

tmp9430 := PrimCons(tmp9424, tmp9429)

__e.Return(PrimCons(symif, tmp9430))
return


}, 1)

tmp9431 := PrimHead(V1866)

tmp9432 := PrimHead(tmp9431)

tmp9433 := PrimHead(V1866)

tmp9434 := PrimTail(tmp9433)

tmp9435 := PrimHead(tmp9434)

tmp9436 := Call(__e, PrimFunc(symshen_4continue), tmp9432, tmp9435, V1868, V1869, V1870, V1871)


__e.TailApply(tmp9422, tmp9436)
return


} else {
tmp9483 := PrimIsPair(V1866)

var ifres9468 Obj

if True == tmp9483 {
tmp9481 := PrimHead(V1866)

tmp9482 := PrimIsPair(tmp9481)

var ifres9470 Obj

if True == tmp9482 {
tmp9478 := PrimHead(V1866)

tmp9479 := PrimTail(tmp9478)

tmp9480 := PrimIsPair(tmp9479)

var ifres9472 Obj

if True == tmp9480 {
tmp9474 := PrimHead(V1866)

tmp9475 := PrimTail(tmp9474)

tmp9476 := PrimTail(tmp9475)

tmp9477 := PrimEqual(Nil, tmp9476)

var ifres9473 Obj

if True == tmp9477 {
ifres9473 = True


} else {
ifres9473 = False


}

ifres9472 = ifres9473


} else {
ifres9472 = False


}

var ifres9471 Obj

if True == ifres9472 {
ifres9471 = True


} else {
ifres9471 = False


}

ifres9470 = ifres9471


} else {
ifres9470 = False


}

var ifres9469 Obj

if True == ifres9470 {
ifres9469 = True


} else {
ifres9469 = False


}

ifres9468 = ifres9469


} else {
ifres9468 = False


}

if True == ifres9468 {
tmp9437 := MakeNative(func(__e *ControlFlow) {
W1874 := __e.Get(1)
_ = W1874
tmp9438 := MakeNative(func(__e *ControlFlow) {
W1875 := __e.Get(1)
_ = W1875
tmp9439 := PrimCons(V1869, Nil)

tmp9440 := PrimCons(symshen_4unlocked_2, tmp9439)

tmp9441 := PrimHead(V1866)

tmp9442 := PrimHead(tmp9441)

tmp9443 := Call(__e, PrimFunc(symshen_4compile_1head), symshen_4_7m, tmp9442, V1867, V1868, W1875)


tmp9444 := PrimCons(False, Nil)

tmp9445 := PrimCons(tmp9443, tmp9444)

tmp9446 := PrimCons(tmp9440, tmp9445)

tmp9447 := PrimCons(symif, tmp9446)

tmp9448 := PrimCons(False, Nil)

tmp9449 := PrimCons(W1874, tmp9448)

tmp9450 := PrimCons(sym_a, tmp9449)

tmp9451 := PrimTail(V1866)

tmp9452 := Call(__e, PrimFunc(symshen_4prolog_1fbody), tmp9451, V1867, V1868, V1869, V1870, V1871, V1872)


tmp9453 := PrimCons(W1874, Nil)

tmp9454 := PrimCons(tmp9452, tmp9453)

tmp9455 := PrimCons(tmp9450, tmp9454)

tmp9456 := PrimCons(symif, tmp9455)

tmp9457 := PrimCons(tmp9456, Nil)

tmp9458 := PrimCons(tmp9447, tmp9457)

tmp9459 := PrimCons(W1874, tmp9458)

__e.Return(PrimCons(symlet, tmp9459))
return


}, 1)

tmp9460 := PrimHead(V1866)

tmp9461 := PrimHead(tmp9460)

tmp9462 := PrimHead(V1866)

tmp9463 := PrimTail(tmp9462)

tmp9464 := PrimHead(tmp9463)

tmp9465 := Call(__e, PrimFunc(symshen_4continue), tmp9461, tmp9464, V1868, V1869, V1870, V1871)


__e.TailApply(tmp9438, tmp9465)
return


}, 1)

tmp9466 := Call(__e, PrimFunc(symgensym), symC)


__e.TailApply(tmp9437, tmp9466)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.prolog-fbody")))
return
}


}


}


}, 7)

tmp9513 := Call(__e, ns2_1set, symshen_4prolog_1fbody, tmp9419)


_ = tmp9513

tmp9514 := MakeNative(func(__e *ControlFlow) {
V1876 := __e.Get(1)
_ = V1876
V1877 := __e.Get(2)
_ = V1877
tmp9519 := Call(__e, PrimFunc(symshen_4locked_2), V1876)


var ifres9516 Obj

if True == tmp9519 {
tmp9518 := Call(__e, PrimFunc(symshen_4fits_2), V1877, V1876)


var ifres9517 Obj

if True == tmp9518 {
ifres9517 = True


} else {
ifres9517 = False


}

ifres9516 = ifres9517


} else {
ifres9516 = False


}

if True == ifres9516 {
__e.TailApply(PrimFunc(symshen_4openlock), V1876)
return
} else {
__e.Return(False)
return
}


}, 2)

tmp9520 := Call(__e, ns2_1set, symshen_4unlock, tmp9514)


_ = tmp9520

tmp9521 := MakeNative(func(__e *ControlFlow) {
V1878 := __e.Get(1)
_ = V1878
tmp9522 := Call(__e, PrimFunc(symshen_4unlocked_2), V1878)


__e.Return(PrimNot(tmp9522))
return


}, 1)

tmp9523 := Call(__e, ns2_1set, symshen_4locked_2, tmp9521)


_ = tmp9523

tmp9524 := MakeNative(func(__e *ControlFlow) {
V1879 := __e.Get(1)
_ = V1879
__e.Return(PrimVectorGet(V1879, MakeNumber(1)))
return
}, 1)

tmp9525 := Call(__e, ns2_1set, symshen_4unlocked_2, tmp9524)


_ = tmp9525

tmp9526 := MakeNative(func(__e *ControlFlow) {
V1880 := __e.Get(1)
_ = V1880
tmp9527 := PrimVectorSet(V1880, MakeNumber(1), True)

_ = tmp9527

__e.Return(False)
return


}, 1)

tmp9528 := Call(__e, ns2_1set, symshen_4openlock, tmp9526)


_ = tmp9528

tmp9529 := MakeNative(func(__e *ControlFlow) {
V1881 := __e.Get(1)
_ = V1881
V1882 := __e.Get(2)
_ = V1882
tmp9530 := PrimVectorGet(V1882, MakeNumber(2))

__e.Return(PrimEqual(V1881, tmp9530))
return


}, 2)

tmp9531 := Call(__e, ns2_1set, symshen_4fits_2, tmp9529)


_ = tmp9531

tmp9532 := MakeNative(func(__e *ControlFlow) {
V1885 := __e.Get(1)
_ = V1885
V1886 := __e.Get(2)
_ = V1886
V1887 := __e.Get(3)
_ = V1887
V1888 := __e.Get(4)
_ = V1888
tmp9533 := MakeNative(func(__e *ControlFlow) {
W1889 := __e.Get(1)
_ = W1889
tmp9538 := PrimEqual(W1889, False)

var ifres9535 Obj

if True == tmp9538 {
tmp9537 := Call(__e, PrimFunc(symshen_4unlocked_2), V1886)


var ifres9536 Obj

if True == tmp9537 {
ifres9536 = True


} else {
ifres9536 = False


}

ifres9535 = ifres9536


} else {
ifres9535 = False


}

if True == ifres9535 {
__e.TailApply(PrimFunc(symshen_4lock), V1887, V1886)
return
} else {
__e.Return(W1889)
return
}


}, 1)

tmp9539 := Call(__e, PrimFunc(symthaw), V1888)


__e.TailApply(tmp9533, tmp9539)
return


}, 4)

tmp9540 := Call(__e, ns2_1set, symshen_4cut, tmp9532)


_ = tmp9540

tmp9541 := MakeNative(func(__e *ControlFlow) {
V1890 := __e.Get(1)
_ = V1890
V1891 := __e.Get(2)
_ = V1891
tmp9542 := MakeNative(func(__e *ControlFlow) {
W1892 := __e.Get(1)
_ = W1892
tmp9543 := MakeNative(func(__e *ControlFlow) {
W1893 := __e.Get(1)
_ = W1893
__e.Return(False)
return
}, 1)

tmp9544 := PrimVectorSet(V1891, MakeNumber(2), V1890)

__e.TailApply(tmp9543, tmp9544)
return


}, 1)

tmp9545 := PrimVectorSet(V1891, MakeNumber(1), False)

__e.TailApply(tmp9542, tmp9545)
return


}, 2)

tmp9546 := Call(__e, ns2_1set, symshen_4lock, tmp9541)


_ = tmp9546

tmp9547 := MakeNative(func(__e *ControlFlow) {
V1894 := __e.Get(1)
_ = V1894
V1895 := __e.Get(2)
_ = V1895
V1896 := __e.Get(3)
_ = V1896
V1897 := __e.Get(4)
_ = V1897
V1898 := __e.Get(5)
_ = V1898
V1899 := __e.Get(6)
_ = V1899
tmp9548 := MakeNative(func(__e *ControlFlow) {
W1900 := __e.Get(1)
_ = W1900
tmp9549 := MakeNative(func(__e *ControlFlow) {
W1901 := __e.Get(1)
_ = W1901
tmp9550 := MakeNative(func(__e *ControlFlow) {
W1902 := __e.Get(1)
_ = W1902
tmp9551 := MakeNative(func(__e *ControlFlow) {
W1903 := __e.Get(1)
_ = W1903
__e.TailApply(PrimFunc(symshen_4stpart), W1902, W1903, V1896)
return
}, 1)

tmp9552 := PrimCons(symshen_4incinfs, Nil)

tmp9553 := Call(__e, PrimFunc(symshen_4compile_1body), V1895, V1896, V1897, V1898, V1899)


tmp9554 := PrimCons(tmp9553, Nil)

tmp9555 := PrimCons(tmp9552, tmp9554)

tmp9556 := PrimCons(symdo, tmp9555)

__e.TailApply(tmp9551, tmp9556)
return


}, 1)

tmp9557 := Call(__e, PrimFunc(symdifference), W1901, W1900)


__e.TailApply(tmp9550, tmp9557)
return


}, 1)

tmp9558 := Call(__e, PrimFunc(symshen_4extract_1free_1vars), V1895)


__e.TailApply(tmp9549, tmp9558)
return


}, 1)

tmp9559 := Call(__e, PrimFunc(symshen_4extract_1vars), V1894)


__e.TailApply(tmp9548, tmp9559)
return


}, 6)

tmp9560 := Call(__e, ns2_1set, symshen_4continue, tmp9547)


_ = tmp9560

tmp9561 := MakeNative(func(__e *ControlFlow) {
V1906 := __e.Get(1)
_ = V1906
tmp9596 := PrimIsPair(V1906)

var ifres9577 Obj

if True == tmp9596 {
tmp9594 := PrimHead(V1906)

tmp9595 := PrimEqual(symlambda, tmp9594)

var ifres9579 Obj

if True == tmp9595 {
tmp9592 := PrimTail(V1906)

tmp9593 := PrimIsPair(tmp9592)

var ifres9581 Obj

if True == tmp9593 {
tmp9589 := PrimTail(V1906)

tmp9590 := PrimTail(tmp9589)

tmp9591 := PrimIsPair(tmp9590)

var ifres9583 Obj

if True == tmp9591 {
tmp9585 := PrimTail(V1906)

tmp9586 := PrimTail(tmp9585)

tmp9587 := PrimTail(tmp9586)

tmp9588 := PrimEqual(Nil, tmp9587)

var ifres9584 Obj

if True == tmp9588 {
ifres9584 = True


} else {
ifres9584 = False


}

ifres9583 = ifres9584


} else {
ifres9583 = False


}

var ifres9582 Obj

if True == ifres9583 {
ifres9582 = True


} else {
ifres9582 = False


}

ifres9581 = ifres9582


} else {
ifres9581 = False


}

var ifres9580 Obj

if True == ifres9581 {
ifres9580 = True


} else {
ifres9580 = False


}

ifres9579 = ifres9580


} else {
ifres9579 = False


}

var ifres9578 Obj

if True == ifres9579 {
ifres9578 = True


} else {
ifres9578 = False


}

ifres9577 = ifres9578


} else {
ifres9577 = False


}

if True == ifres9577 {
tmp9562 := PrimTail(V1906)

tmp9563 := PrimHead(tmp9562)

tmp9564 := PrimTail(V1906)

tmp9565 := PrimTail(tmp9564)

tmp9566 := PrimHead(tmp9565)

tmp9567 := Call(__e, PrimFunc(symshen_4extract_1free_1vars), tmp9566)


__e.TailApply(PrimFunc(symremove), tmp9563, tmp9567)
return


} else {
tmp9575 := PrimIsPair(V1906)

if True == tmp9575 {
tmp9568 := PrimHead(V1906)

tmp9569 := Call(__e, PrimFunc(symshen_4extract_1free_1vars), tmp9568)


tmp9570 := PrimTail(V1906)

tmp9571 := Call(__e, PrimFunc(symshen_4extract_1free_1vars), tmp9570)


__e.TailApply(PrimFunc(symunion), tmp9569, tmp9571)
return


} else {
tmp9573 := PrimIsVariable(V1906)

if True == tmp9573 {
__e.Return(PrimCons(V1906, Nil))
return
} else {
__e.Return(Nil)
return
}


}


}


}, 1)

tmp9597 := Call(__e, ns2_1set, symshen_4extract_1free_1vars, tmp9561)


_ = tmp9597

tmp9598 := MakeNative(func(__e *ControlFlow) {
V1923 := __e.Get(1)
_ = V1923
V1924 := __e.Get(2)
_ = V1924
V1925 := __e.Get(3)
_ = V1925
V1926 := __e.Get(4)
_ = V1926
V1927 := __e.Get(5)
_ = V1927
tmp9633 := PrimEqual(Nil, V1923)

if True == tmp9633 {
tmp9599 := PrimCons(V1927, Nil)

__e.Return(PrimCons(symthaw, tmp9599))
return


} else {
tmp9631 := PrimIsPair(V1923)

var ifres9627 Obj

if True == tmp9631 {
tmp9629 := PrimHead(V1923)

tmp9630 := PrimEqual(sym_b, tmp9629)

var ifres9628 Obj

if True == tmp9630 {
ifres9628 = True


} else {
ifres9628 = False


}

ifres9627 = ifres9628


} else {
ifres9627 = False


}

if True == ifres9627 {
tmp9600 := PrimCons(symshen_4cut, Nil)

tmp9601 := PrimTail(V1923)

tmp9602 := PrimCons(tmp9600, tmp9601)

__e.TailApply(PrimFunc(symshen_4compile_1body), tmp9602, V1924, V1925, V1926, V1927)
return


} else {
tmp9625 := PrimIsPair(V1923)

var ifres9621 Obj

if True == tmp9625 {
tmp9623 := PrimTail(V1923)

tmp9624 := PrimEqual(Nil, tmp9623)

var ifres9622 Obj

if True == tmp9624 {
ifres9622 = True


} else {
ifres9622 = False


}

ifres9621 = ifres9622


} else {
ifres9621 = False


}

if True == ifres9621 {
tmp9603 := PrimHead(V1923)

tmp9604 := Call(__e, PrimFunc(symshen_4deref_1calls), tmp9603, V1924)


tmp9605 := PrimCons(V1927, Nil)

tmp9606 := PrimCons(V1926, tmp9605)

tmp9607 := PrimCons(V1925, tmp9606)

tmp9608 := PrimCons(V1924, tmp9607)

__e.TailApply(PrimFunc(symappend), tmp9604, tmp9608)
return


} else {
tmp9619 := PrimIsPair(V1923)

if True == tmp9619 {
tmp9609 := MakeNative(func(__e *ControlFlow) {
W1928 := __e.Get(1)
_ = W1928
tmp9610 := PrimTail(V1923)

tmp9611 := Call(__e, PrimFunc(symshen_4freeze_1literals), tmp9610, V1924, V1925, V1926, V1927)


tmp9612 := PrimCons(tmp9611, Nil)

tmp9613 := PrimCons(V1926, tmp9612)

tmp9614 := PrimCons(V1925, tmp9613)

tmp9615 := PrimCons(V1924, tmp9614)

__e.TailApply(PrimFunc(symappend), W1928, tmp9615)
return


}, 1)

tmp9616 := PrimHead(V1923)

tmp9617 := Call(__e, PrimFunc(symshen_4deref_1calls), tmp9616, V1924)


__e.TailApply(tmp9609, tmp9617)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.compile-fbody")))
return
}


}


}


}


}, 5)

tmp9634 := Call(__e, ns2_1set, symshen_4compile_1body, tmp9598)


_ = tmp9634

tmp9635 := MakeNative(func(__e *ControlFlow) {
V1945 := __e.Get(1)
_ = V1945
V1946 := __e.Get(2)
_ = V1946
V1947 := __e.Get(3)
_ = V1947
V1948 := __e.Get(4)
_ = V1948
V1949 := __e.Get(5)
_ = V1949
tmp9659 := PrimEqual(Nil, V1945)

if True == tmp9659 {
__e.Return(V1949)
return
} else {
tmp9657 := PrimIsPair(V1945)

var ifres9653 Obj

if True == tmp9657 {
tmp9655 := PrimHead(V1945)

tmp9656 := PrimEqual(sym_b, tmp9655)

var ifres9654 Obj

if True == tmp9656 {
ifres9654 = True


} else {
ifres9654 = False


}

ifres9653 = ifres9654


} else {
ifres9653 = False


}

if True == ifres9653 {
tmp9636 := PrimCons(symshen_4cut, Nil)

tmp9637 := PrimTail(V1945)

tmp9638 := PrimCons(tmp9636, tmp9637)

__e.TailApply(PrimFunc(symshen_4freeze_1literals), tmp9638, V1946, V1947, V1948, V1949)
return


} else {
tmp9651 := PrimIsPair(V1945)

if True == tmp9651 {
tmp9639 := MakeNative(func(__e *ControlFlow) {
W1950 := __e.Get(1)
_ = W1950
tmp9640 := PrimTail(V1945)

tmp9641 := Call(__e, PrimFunc(symshen_4freeze_1literals), tmp9640, V1946, V1947, V1948, V1949)


tmp9642 := PrimCons(tmp9641, Nil)

tmp9643 := PrimCons(V1948, tmp9642)

tmp9644 := PrimCons(V1947, tmp9643)

tmp9645 := PrimCons(V1946, tmp9644)

tmp9646 := Call(__e, PrimFunc(symappend), W1950, tmp9645)


tmp9647 := PrimCons(tmp9646, Nil)

__e.Return(PrimCons(symfreeze, tmp9647))
return


}, 1)

tmp9648 := PrimHead(V1945)

tmp9649 := Call(__e, PrimFunc(symshen_4deref_1calls), tmp9648, V1946)


__e.TailApply(tmp9639, tmp9649)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.freeze-literals")))
return
}


}


}


}, 5)

tmp9660 := Call(__e, ns2_1set, symshen_4freeze_1literals, tmp9635)


_ = tmp9660

tmp9661 := MakeNative(func(__e *ControlFlow) {
V1955 := __e.Get(1)
_ = V1955
V1956 := __e.Get(2)
_ = V1956
tmp9676 := PrimIsPair(V1955)

var ifres9672 Obj

if True == tmp9676 {
tmp9674 := PrimHead(V1955)

tmp9675 := PrimEqual(symfork, tmp9674)

var ifres9673 Obj

if True == tmp9675 {
ifres9673 = True


} else {
ifres9673 = False


}

ifres9672 = ifres9673


} else {
ifres9672 = False


}

if True == ifres9672 {
tmp9662 := PrimTail(V1955)

tmp9663 := Call(__e, PrimFunc(symshen_4deref_1forked_1literals), tmp9662, V1956)


tmp9664 := PrimCons(tmp9663, Nil)

__e.Return(PrimCons(symfork, tmp9664))
return


} else {
tmp9670 := PrimIsPair(V1955)

if True == tmp9670 {
tmp9665 := PrimHead(V1955)

tmp9666 := MakeNative(func(__e *ControlFlow) {
Z1957 := __e.Get(1)
_ = Z1957
__e.TailApply(PrimFunc(symshen_4function_1calls), Z1957, V1956)
return
}, 1)

tmp9667 := PrimTail(V1955)

tmp9668 := Call(__e, PrimFunc(symmap), tmp9666, tmp9667)


__e.Return(PrimCons(tmp9665, tmp9668))
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.deref-calls")))
return
}


}


}, 2)

tmp9677 := Call(__e, ns2_1set, symshen_4deref_1calls, tmp9661)


_ = tmp9677

tmp9678 := MakeNative(func(__e *ControlFlow) {
V1964 := __e.Get(1)
_ = V1964
V1965 := __e.Get(2)
_ = V1965
tmp9688 := PrimEqual(Nil, V1964)

if True == tmp9688 {
__e.Return(Nil)
return
} else {
tmp9686 := PrimIsPair(V1964)

if True == tmp9686 {
tmp9679 := PrimHead(V1964)

tmp9680 := Call(__e, PrimFunc(symshen_4deref_1calls), tmp9679, V1965)


tmp9681 := PrimTail(V1964)

tmp9682 := Call(__e, PrimFunc(symshen_4deref_1forked_1literals), tmp9681, V1965)


tmp9683 := PrimCons(tmp9682, Nil)

tmp9684 := PrimCons(tmp9680, tmp9683)

__e.Return(PrimCons(symcons, tmp9684))
return


} else {
__e.Return(PrimSimpleError(MakeString("fork requires a list of literals\n")))
return
}


}


}, 2)

tmp9689 := Call(__e, ns2_1set, symshen_4deref_1forked_1literals, tmp9678)


_ = tmp9689

tmp9690 := MakeNative(func(__e *ControlFlow) {
V1968 := __e.Get(1)
_ = V1968
V1969 := __e.Get(2)
_ = V1969
tmp9722 := PrimIsPair(V1968)

var ifres9703 Obj

if True == tmp9722 {
tmp9720 := PrimHead(V1968)

tmp9721 := PrimEqual(symcons, tmp9720)

var ifres9705 Obj

if True == tmp9721 {
tmp9718 := PrimTail(V1968)

tmp9719 := PrimIsPair(tmp9718)

var ifres9707 Obj

if True == tmp9719 {
tmp9715 := PrimTail(V1968)

tmp9716 := PrimTail(tmp9715)

tmp9717 := PrimIsPair(tmp9716)

var ifres9709 Obj

if True == tmp9717 {
tmp9711 := PrimTail(V1968)

tmp9712 := PrimTail(tmp9711)

tmp9713 := PrimTail(tmp9712)

tmp9714 := PrimEqual(Nil, tmp9713)

var ifres9710 Obj

if True == tmp9714 {
ifres9710 = True


} else {
ifres9710 = False


}

ifres9709 = ifres9710


} else {
ifres9709 = False


}

var ifres9708 Obj

if True == ifres9709 {
ifres9708 = True


} else {
ifres9708 = False


}

ifres9707 = ifres9708


} else {
ifres9707 = False


}

var ifres9706 Obj

if True == ifres9707 {
ifres9706 = True


} else {
ifres9706 = False


}

ifres9705 = ifres9706


} else {
ifres9705 = False


}

var ifres9704 Obj

if True == ifres9705 {
ifres9704 = True


} else {
ifres9704 = False


}

ifres9703 = ifres9704


} else {
ifres9703 = False


}

if True == ifres9703 {
tmp9691 := PrimTail(V1968)

tmp9692 := PrimHead(tmp9691)

tmp9693 := Call(__e, PrimFunc(symshen_4function_1calls), tmp9692, V1969)


tmp9694 := PrimTail(V1968)

tmp9695 := PrimTail(tmp9694)

tmp9696 := PrimHead(tmp9695)

tmp9697 := Call(__e, PrimFunc(symshen_4function_1calls), tmp9696, V1969)


tmp9698 := PrimCons(tmp9697, Nil)

tmp9699 := PrimCons(tmp9693, tmp9698)

__e.Return(PrimCons(symcons, tmp9699))
return


} else {
tmp9701 := PrimIsPair(V1968)

if True == tmp9701 {
__e.TailApply(PrimFunc(symshen_4deref_1terms), V1968, V1969, Nil)
return
} else {
__e.Return(V1968)
return
}


}


}, 2)

tmp9723 := Call(__e, ns2_1set, symshen_4function_1calls, tmp9690)


_ = tmp9723

tmp9724 := MakeNative(func(__e *ControlFlow) {
V1978 := __e.Get(1)
_ = V1978
V1979 := __e.Get(2)
_ = V1979
V1980 := __e.Get(3)
_ = V1980
tmp9818 := PrimIsPair(V1978)

var ifres9805 Obj

if True == tmp9818 {
tmp9816 := PrimHead(V1978)

tmp9817 := PrimEqual(MakeNumber(0), tmp9816)

var ifres9807 Obj

if True == tmp9817 {
tmp9814 := PrimTail(V1978)

tmp9815 := PrimIsPair(tmp9814)

var ifres9809 Obj

if True == tmp9815 {
tmp9811 := PrimTail(V1978)

tmp9812 := PrimTail(tmp9811)

tmp9813 := PrimEqual(Nil, tmp9812)

var ifres9810 Obj

if True == tmp9813 {
ifres9810 = True


} else {
ifres9810 = False


}

ifres9809 = ifres9810


} else {
ifres9809 = False


}

var ifres9808 Obj

if True == ifres9809 {
ifres9808 = True


} else {
ifres9808 = False


}

ifres9807 = ifres9808


} else {
ifres9807 = False


}

var ifres9806 Obj

if True == ifres9807 {
ifres9806 = True


} else {
ifres9806 = False


}

ifres9805 = ifres9806


} else {
ifres9805 = False


}

if True == ifres9805 {
tmp9731 := PrimTail(V1978)

tmp9732 := PrimHead(tmp9731)

tmp9733 := PrimIsVariable(tmp9732)

if True == tmp9733 {
tmp9725 := PrimTail(V1978)

__e.Return(PrimHead(tmp9725))
return


} else {
tmp9726 := PrimTail(V1978)

tmp9727 := PrimHead(tmp9726)

tmp9728 := Call(__e, PrimFunc(symshen_4app), tmp9727, MakeString("\n"), symshen_4s)


tmp9729 := PrimStringConcat(MakeString("attempt to optimise a non-variable "), tmp9728)

__e.Return(PrimSimpleError(tmp9729))
return


}


} else {
tmp9803 := PrimIsPair(V1978)

var ifres9790 Obj

if True == tmp9803 {
tmp9801 := PrimHead(V1978)

tmp9802 := PrimEqual(MakeNumber(1), tmp9801)

var ifres9792 Obj

if True == tmp9802 {
tmp9799 := PrimTail(V1978)

tmp9800 := PrimIsPair(tmp9799)

var ifres9794 Obj

if True == tmp9800 {
tmp9796 := PrimTail(V1978)

tmp9797 := PrimTail(tmp9796)

tmp9798 := PrimEqual(Nil, tmp9797)

var ifres9795 Obj

if True == tmp9798 {
ifres9795 = True


} else {
ifres9795 = False


}

ifres9794 = ifres9795


} else {
ifres9794 = False


}

var ifres9793 Obj

if True == ifres9794 {
ifres9793 = True


} else {
ifres9793 = False


}

ifres9792 = ifres9793


} else {
ifres9792 = False


}

var ifres9791 Obj

if True == ifres9792 {
ifres9791 = True


} else {
ifres9791 = False


}

ifres9790 = ifres9791


} else {
ifres9790 = False


}

if True == ifres9790 {
tmp9743 := PrimTail(V1978)

tmp9744 := PrimHead(tmp9743)

tmp9745 := PrimIsVariable(tmp9744)

if True == tmp9745 {
tmp9734 := PrimTail(V1978)

tmp9735 := PrimHead(tmp9734)

tmp9736 := PrimCons(V1979, Nil)

tmp9737 := PrimCons(tmp9735, tmp9736)

__e.Return(PrimCons(symshen_4lazyderef, tmp9737))
return


} else {
tmp9738 := PrimTail(V1978)

tmp9739 := PrimHead(tmp9738)

tmp9740 := Call(__e, PrimFunc(symshen_4app), tmp9739, MakeString("\n"), symshen_4s)


tmp9741 := PrimStringConcat(MakeString("attempt to optimise a non-variable "), tmp9740)

__e.Return(PrimSimpleError(tmp9741))
return


}


} else {
tmp9787 := Call(__e, PrimFunc(symelement_2), V1978, V1980)


tmp9788 := PrimNot(tmp9787)

var ifres9784 Obj

if True == tmp9788 {
tmp9786 := PrimIsVariable(V1978)

var ifres9785 Obj

if True == tmp9786 {
ifres9785 = True


} else {
ifres9785 = False


}

ifres9784 = ifres9785


} else {
ifres9784 = False


}

if True == ifres9784 {
tmp9746 := PrimCons(V1979, Nil)

tmp9747 := PrimCons(V1978, tmp9746)

__e.Return(PrimCons(symshen_4deref, tmp9747))
return


} else {
tmp9782 := PrimIsPair(V1978)

var ifres9763 Obj

if True == tmp9782 {
tmp9780 := PrimHead(V1978)

tmp9781 := PrimEqual(symlambda, tmp9780)

var ifres9765 Obj

if True == tmp9781 {
tmp9778 := PrimTail(V1978)

tmp9779 := PrimIsPair(tmp9778)

var ifres9767 Obj

if True == tmp9779 {
tmp9775 := PrimTail(V1978)

tmp9776 := PrimTail(tmp9775)

tmp9777 := PrimIsPair(tmp9776)

var ifres9769 Obj

if True == tmp9777 {
tmp9771 := PrimTail(V1978)

tmp9772 := PrimTail(tmp9771)

tmp9773 := PrimTail(tmp9772)

tmp9774 := PrimEqual(Nil, tmp9773)

var ifres9770 Obj

if True == tmp9774 {
ifres9770 = True


} else {
ifres9770 = False


}

ifres9769 = ifres9770


} else {
ifres9769 = False


}

var ifres9768 Obj

if True == ifres9769 {
ifres9768 = True


} else {
ifres9768 = False


}

ifres9767 = ifres9768


} else {
ifres9767 = False


}

var ifres9766 Obj

if True == ifres9767 {
ifres9766 = True


} else {
ifres9766 = False


}

ifres9765 = ifres9766


} else {
ifres9765 = False


}

var ifres9764 Obj

if True == ifres9765 {
ifres9764 = True


} else {
ifres9764 = False


}

ifres9763 = ifres9764


} else {
ifres9763 = False


}

if True == ifres9763 {
tmp9748 := PrimTail(V1978)

tmp9749 := PrimHead(tmp9748)

tmp9750 := PrimTail(V1978)

tmp9751 := PrimTail(tmp9750)

tmp9752 := PrimHead(tmp9751)

tmp9753 := PrimTail(V1978)

tmp9754 := PrimHead(tmp9753)

tmp9755 := PrimCons(tmp9754, V1980)

tmp9756 := Call(__e, PrimFunc(symshen_4deref_1terms), tmp9752, V1979, tmp9755)


tmp9757 := PrimCons(tmp9756, Nil)

tmp9758 := PrimCons(tmp9749, tmp9757)

__e.Return(PrimCons(symlambda, tmp9758))
return


} else {
tmp9761 := PrimIsPair(V1978)

if True == tmp9761 {
tmp9759 := MakeNative(func(__e *ControlFlow) {
Z1981 := __e.Get(1)
_ = Z1981
__e.TailApply(PrimFunc(symshen_4deref_1terms), Z1981, V1979, V1980)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp9759, V1978)
return


} else {
__e.Return(V1978)
return
}


}


}


}


}


}, 3)

tmp9819 := Call(__e, ns2_1set, symshen_4deref_1terms, tmp9724)


_ = tmp9819

tmp9820 := MakeNative(func(__e *ControlFlow) {
V1999 := __e.Get(1)
_ = V1999
V2000 := __e.Get(2)
_ = V2000
V2001 := __e.Get(3)
_ = V2001
V2002 := __e.Get(4)
_ = V2002
V2003 := __e.Get(5)
_ = V2003
tmp9996 := PrimEqual(Nil, V2000)

var ifres9993 Obj

if True == tmp9996 {
tmp9995 := PrimEqual(Nil, V2001)

var ifres9994 Obj

if True == tmp9995 {
ifres9994 = True


} else {
ifres9994 = False


}

ifres9993 = ifres9994


} else {
ifres9993 = False


}

if True == ifres9993 {
__e.Return(V2003)
return
} else {
tmp9991 := PrimIsPair(V2000)

var ifres9971 Obj

if True == tmp9991 {
tmp9989 := PrimHead(V2000)

tmp9990 := PrimIsPair(tmp9989)

var ifres9973 Obj

if True == tmp9990 {
tmp9986 := PrimHead(V2000)

tmp9987 := PrimHead(tmp9986)

tmp9988 := PrimEqual(symshen_4_7m, tmp9987)

var ifres9975 Obj

if True == tmp9988 {
tmp9983 := PrimHead(V2000)

tmp9984 := PrimTail(tmp9983)

tmp9985 := PrimIsPair(tmp9984)

var ifres9977 Obj

if True == tmp9985 {
tmp9979 := PrimHead(V2000)

tmp9980 := PrimTail(tmp9979)

tmp9981 := PrimTail(tmp9980)

tmp9982 := PrimEqual(Nil, tmp9981)

var ifres9978 Obj

if True == tmp9982 {
ifres9978 = True


} else {
ifres9978 = False


}

ifres9977 = ifres9978


} else {
ifres9977 = False


}

var ifres9976 Obj

if True == ifres9977 {
ifres9976 = True


} else {
ifres9976 = False


}

ifres9975 = ifres9976


} else {
ifres9975 = False


}

var ifres9974 Obj

if True == ifres9975 {
ifres9974 = True


} else {
ifres9974 = False


}

ifres9973 = ifres9974


} else {
ifres9973 = False


}

var ifres9972 Obj

if True == ifres9973 {
ifres9972 = True


} else {
ifres9972 = False


}

ifres9971 = ifres9972


} else {
ifres9971 = False


}

if True == ifres9971 {
tmp9821 := PrimHead(V2000)

tmp9822 := PrimTail(tmp9821)

tmp9823 := PrimHead(tmp9822)

tmp9824 := PrimTail(V2000)

tmp9825 := PrimCons(V1999, tmp9824)

tmp9826 := PrimCons(tmp9823, tmp9825)

tmp9827 := PrimCons(symshen_4_7m, tmp9826)

__e.TailApply(PrimFunc(symshen_4compile_1head), V1999, tmp9827, V2001, V2002, V2003)
return


} else {
tmp9969 := PrimIsPair(V2000)

var ifres9949 Obj

if True == tmp9969 {
tmp9967 := PrimHead(V2000)

tmp9968 := PrimIsPair(tmp9967)

var ifres9951 Obj

if True == tmp9968 {
tmp9964 := PrimHead(V2000)

tmp9965 := PrimHead(tmp9964)

tmp9966 := PrimEqual(symshen_4_1m, tmp9965)

var ifres9953 Obj

if True == tmp9966 {
tmp9961 := PrimHead(V2000)

tmp9962 := PrimTail(tmp9961)

tmp9963 := PrimIsPair(tmp9962)

var ifres9955 Obj

if True == tmp9963 {
tmp9957 := PrimHead(V2000)

tmp9958 := PrimTail(tmp9957)

tmp9959 := PrimTail(tmp9958)

tmp9960 := PrimEqual(Nil, tmp9959)

var ifres9956 Obj

if True == tmp9960 {
ifres9956 = True


} else {
ifres9956 = False


}

ifres9955 = ifres9956


} else {
ifres9955 = False


}

var ifres9954 Obj

if True == ifres9955 {
ifres9954 = True


} else {
ifres9954 = False


}

ifres9953 = ifres9954


} else {
ifres9953 = False


}

var ifres9952 Obj

if True == ifres9953 {
ifres9952 = True


} else {
ifres9952 = False


}

ifres9951 = ifres9952


} else {
ifres9951 = False


}

var ifres9950 Obj

if True == ifres9951 {
ifres9950 = True


} else {
ifres9950 = False


}

ifres9949 = ifres9950


} else {
ifres9949 = False


}

if True == ifres9949 {
tmp9828 := PrimHead(V2000)

tmp9829 := PrimTail(tmp9828)

tmp9830 := PrimHead(tmp9829)

tmp9831 := PrimTail(V2000)

tmp9832 := PrimCons(V1999, tmp9831)

tmp9833 := PrimCons(tmp9830, tmp9832)

tmp9834 := PrimCons(symshen_4_1m, tmp9833)

__e.TailApply(PrimFunc(symshen_4compile_1head), V1999, tmp9834, V2001, V2002, V2003)
return


} else {
tmp9947 := PrimIsPair(V2000)

var ifres9943 Obj

if True == tmp9947 {
tmp9945 := PrimHead(V2000)

tmp9946 := PrimEqual(symshen_4_1m, tmp9945)

var ifres9944 Obj

if True == tmp9946 {
ifres9944 = True


} else {
ifres9944 = False


}

ifres9943 = ifres9944


} else {
ifres9943 = False


}

if True == ifres9943 {
tmp9835 := PrimTail(V2000)

__e.TailApply(PrimFunc(symshen_4compile_1head), symshen_4_1m, tmp9835, V2001, V2002, V2003)
return


} else {
tmp9941 := PrimIsPair(V2000)

var ifres9937 Obj

if True == tmp9941 {
tmp9939 := PrimHead(V2000)

tmp9940 := PrimEqual(symshen_4_7m, tmp9939)

var ifres9938 Obj

if True == tmp9940 {
ifres9938 = True


} else {
ifres9938 = False


}

ifres9937 = ifres9938


} else {
ifres9937 = False


}

if True == ifres9937 {
tmp9836 := PrimTail(V2000)

__e.TailApply(PrimFunc(symshen_4compile_1head), symshen_4_7m, tmp9836, V2001, V2002, V2003)
return


} else {
tmp9935 := PrimIsPair(V2000)

var ifres9928 Obj

if True == tmp9935 {
tmp9934 := PrimIsPair(V2001)

var ifres9930 Obj

if True == tmp9934 {
tmp9932 := PrimHead(V2000)

tmp9933 := Call(__e, PrimFunc(symshen_4wildcard_2), tmp9932)


var ifres9931 Obj

if True == tmp9933 {
ifres9931 = True


} else {
ifres9931 = False


}

ifres9930 = ifres9931


} else {
ifres9930 = False


}

var ifres9929 Obj

if True == ifres9930 {
ifres9929 = True


} else {
ifres9929 = False


}

ifres9928 = ifres9929


} else {
ifres9928 = False


}

if True == ifres9928 {
tmp9837 := PrimTail(V2000)

tmp9838 := PrimTail(V2001)

__e.TailApply(PrimFunc(symshen_4compile_1head), V1999, tmp9837, tmp9838, V2002, V2003)
return


} else {
tmp9926 := PrimIsPair(V2000)

var ifres9922 Obj

if True == tmp9926 {
tmp9924 := PrimHead(V2000)

tmp9925 := PrimIsVariable(tmp9924)

var ifres9923 Obj

if True == tmp9925 {
ifres9923 = True


} else {
ifres9923 = False


}

ifres9922 = ifres9923


} else {
ifres9922 = False


}

if True == ifres9922 {
__e.TailApply(PrimFunc(symshen_4variable_1case), V1999, V2000, V2001, V2002, V2003)
return
} else {
tmp9920 := PrimEqual(symshen_4_1m, V1999)

var ifres9913 Obj

if True == tmp9920 {
tmp9919 := PrimIsPair(V2000)

var ifres9915 Obj

if True == tmp9919 {
tmp9917 := PrimHead(V2000)

tmp9918 := Call(__e, PrimFunc(symatom_2), tmp9917)


var ifres9916 Obj

if True == tmp9918 {
ifres9916 = True


} else {
ifres9916 = False


}

ifres9915 = ifres9916


} else {
ifres9915 = False


}

var ifres9914 Obj

if True == ifres9915 {
ifres9914 = True


} else {
ifres9914 = False


}

ifres9913 = ifres9914


} else {
ifres9913 = False


}

if True == ifres9913 {
__e.TailApply(PrimFunc(symshen_4atom_1case_1minus), V2000, V2001, V2002, V2003)
return
} else {
tmp9911 := PrimEqual(symshen_4_1m, V1999)

var ifres9881 Obj

if True == tmp9911 {
tmp9910 := PrimIsPair(V2000)

var ifres9883 Obj

if True == tmp9910 {
tmp9908 := PrimHead(V2000)

tmp9909 := PrimIsPair(tmp9908)

var ifres9885 Obj

if True == tmp9909 {
tmp9905 := PrimHead(V2000)

tmp9906 := PrimHead(tmp9905)

tmp9907 := PrimEqual(symcons, tmp9906)

var ifres9887 Obj

if True == tmp9907 {
tmp9902 := PrimHead(V2000)

tmp9903 := PrimTail(tmp9902)

tmp9904 := PrimIsPair(tmp9903)

var ifres9889 Obj

if True == tmp9904 {
tmp9898 := PrimHead(V2000)

tmp9899 := PrimTail(tmp9898)

tmp9900 := PrimTail(tmp9899)

tmp9901 := PrimIsPair(tmp9900)

var ifres9891 Obj

if True == tmp9901 {
tmp9893 := PrimHead(V2000)

tmp9894 := PrimTail(tmp9893)

tmp9895 := PrimTail(tmp9894)

tmp9896 := PrimTail(tmp9895)

tmp9897 := PrimEqual(Nil, tmp9896)

var ifres9892 Obj

if True == tmp9897 {
ifres9892 = True


} else {
ifres9892 = False


}

ifres9891 = ifres9892


} else {
ifres9891 = False


}

var ifres9890 Obj

if True == ifres9891 {
ifres9890 = True


} else {
ifres9890 = False


}

ifres9889 = ifres9890


} else {
ifres9889 = False


}

var ifres9888 Obj

if True == ifres9889 {
ifres9888 = True


} else {
ifres9888 = False


}

ifres9887 = ifres9888


} else {
ifres9887 = False


}

var ifres9886 Obj

if True == ifres9887 {
ifres9886 = True


} else {
ifres9886 = False


}

ifres9885 = ifres9886


} else {
ifres9885 = False


}

var ifres9884 Obj

if True == ifres9885 {
ifres9884 = True


} else {
ifres9884 = False


}

ifres9883 = ifres9884


} else {
ifres9883 = False


}

var ifres9882 Obj

if True == ifres9883 {
ifres9882 = True


} else {
ifres9882 = False


}

ifres9881 = ifres9882


} else {
ifres9881 = False


}

if True == ifres9881 {
__e.TailApply(PrimFunc(symshen_4cons_1case_1minus), V2000, V2001, V2002, V2003)
return
} else {
tmp9879 := PrimEqual(symshen_4_7m, V1999)

var ifres9872 Obj

if True == tmp9879 {
tmp9878 := PrimIsPair(V2000)

var ifres9874 Obj

if True == tmp9878 {
tmp9876 := PrimHead(V2000)

tmp9877 := Call(__e, PrimFunc(symatom_2), tmp9876)


var ifres9875 Obj

if True == tmp9877 {
ifres9875 = True


} else {
ifres9875 = False


}

ifres9874 = ifres9875


} else {
ifres9874 = False


}

var ifres9873 Obj

if True == ifres9874 {
ifres9873 = True


} else {
ifres9873 = False


}

ifres9872 = ifres9873


} else {
ifres9872 = False


}

if True == ifres9872 {
__e.TailApply(PrimFunc(symshen_4atom_1case_1plus), V2000, V2001, V2002, V2003)
return
} else {
tmp9870 := PrimEqual(symshen_4_7m, V1999)

var ifres9840 Obj

if True == tmp9870 {
tmp9869 := PrimIsPair(V2000)

var ifres9842 Obj

if True == tmp9869 {
tmp9867 := PrimHead(V2000)

tmp9868 := PrimIsPair(tmp9867)

var ifres9844 Obj

if True == tmp9868 {
tmp9864 := PrimHead(V2000)

tmp9865 := PrimHead(tmp9864)

tmp9866 := PrimEqual(symcons, tmp9865)

var ifres9846 Obj

if True == tmp9866 {
tmp9861 := PrimHead(V2000)

tmp9862 := PrimTail(tmp9861)

tmp9863 := PrimIsPair(tmp9862)

var ifres9848 Obj

if True == tmp9863 {
tmp9857 := PrimHead(V2000)

tmp9858 := PrimTail(tmp9857)

tmp9859 := PrimTail(tmp9858)

tmp9860 := PrimIsPair(tmp9859)

var ifres9850 Obj

if True == tmp9860 {
tmp9852 := PrimHead(V2000)

tmp9853 := PrimTail(tmp9852)

tmp9854 := PrimTail(tmp9853)

tmp9855 := PrimTail(tmp9854)

tmp9856 := PrimEqual(Nil, tmp9855)

var ifres9851 Obj

if True == tmp9856 {
ifres9851 = True


} else {
ifres9851 = False


}

ifres9850 = ifres9851


} else {
ifres9850 = False


}

var ifres9849 Obj

if True == ifres9850 {
ifres9849 = True


} else {
ifres9849 = False


}

ifres9848 = ifres9849


} else {
ifres9848 = False


}

var ifres9847 Obj

if True == ifres9848 {
ifres9847 = True


} else {
ifres9847 = False


}

ifres9846 = ifres9847


} else {
ifres9846 = False


}

var ifres9845 Obj

if True == ifres9846 {
ifres9845 = True


} else {
ifres9845 = False


}

ifres9844 = ifres9845


} else {
ifres9844 = False


}

var ifres9843 Obj

if True == ifres9844 {
ifres9843 = True


} else {
ifres9843 = False


}

ifres9842 = ifres9843


} else {
ifres9842 = False


}

var ifres9841 Obj

if True == ifres9842 {
ifres9841 = True


} else {
ifres9841 = False


}

ifres9840 = ifres9841


} else {
ifres9840 = False


}

if True == ifres9840 {
__e.TailApply(PrimFunc(symshen_4cons_1case_1plus), V2000, V2001, V2002, V2003)
return
} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.compile-head")))
return
}


}


}


}


}


}


}


}


}


}


}


}, 5)

tmp9997 := Call(__e, ns2_1set, symshen_4compile_1head, tmp9820)


_ = tmp9997

tmp9998 := MakeNative(func(__e *ControlFlow) {
V2014 := __e.Get(1)
_ = V2014
V2015 := __e.Get(2)
_ = V2015
V2016 := __e.Get(3)
_ = V2016
V2017 := __e.Get(4)
_ = V2017
V2018 := __e.Get(5)
_ = V2018
tmp10019 := PrimIsPair(V2015)

var ifres10016 Obj

if True == tmp10019 {
tmp10018 := PrimIsPair(V2016)

var ifres10017 Obj

if True == tmp10018 {
ifres10017 = True


} else {
ifres10017 = False


}

ifres10016 = ifres10017


} else {
ifres10016 = False


}

if True == ifres10016 {
tmp10013 := PrimHead(V2016)

tmp10014 := PrimIsVariable(tmp10013)

if True == tmp10014 {
tmp9999 := PrimTail(V2015)

tmp10000 := PrimTail(V2016)

tmp10001 := PrimHead(V2016)

tmp10002 := PrimHead(V2015)

tmp10003 := Call(__e, PrimFunc(symsubst), tmp10001, tmp10002, V2018)


__e.TailApply(PrimFunc(symshen_4compile_1head), V2014, tmp9999, tmp10000, V2017, tmp10003)
return


} else {
tmp10004 := PrimHead(V2015)

tmp10005 := PrimHead(V2016)

tmp10006 := PrimTail(V2015)

tmp10007 := PrimTail(V2016)

tmp10008 := Call(__e, PrimFunc(symshen_4compile_1head), V2014, tmp10006, tmp10007, V2017, V2018)


tmp10009 := PrimCons(tmp10008, Nil)

tmp10010 := PrimCons(tmp10005, tmp10009)

tmp10011 := PrimCons(tmp10004, tmp10010)

__e.Return(PrimCons(symlet, tmp10011))
return


}


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.variable-case")))
return
}


}, 5)

tmp10020 := Call(__e, ns2_1set, symshen_4variable_1case, tmp9998)


_ = tmp10020

tmp10021 := MakeNative(func(__e *ControlFlow) {
V2027 := __e.Get(1)
_ = V2027
V2028 := __e.Get(2)
_ = V2028
V2029 := __e.Get(3)
_ = V2029
V2030 := __e.Get(4)
_ = V2030
tmp10046 := PrimIsPair(V2027)

var ifres10043 Obj

if True == tmp10046 {
tmp10045 := PrimIsPair(V2028)

var ifres10044 Obj

if True == tmp10045 {
ifres10044 = True


} else {
ifres10044 = False


}

ifres10043 = ifres10044


} else {
ifres10043 = False


}

if True == ifres10043 {
tmp10022 := MakeNative(func(__e *ControlFlow) {
W2031 := __e.Get(1)
_ = W2031
tmp10023 := PrimHead(V2028)

tmp10024 := PrimCons(V2029, Nil)

tmp10025 := PrimCons(tmp10023, tmp10024)

tmp10026 := PrimCons(symshen_4lazyderef, tmp10025)

tmp10027 := PrimHead(V2027)

tmp10028 := PrimCons(tmp10027, Nil)

tmp10029 := PrimCons(W2031, tmp10028)

tmp10030 := PrimCons(sym_a, tmp10029)

tmp10031 := PrimTail(V2027)

tmp10032 := PrimTail(V2028)

tmp10033 := Call(__e, PrimFunc(symshen_4compile_1head), symshen_4_1m, tmp10031, tmp10032, V2029, V2030)


tmp10034 := PrimCons(False, Nil)

tmp10035 := PrimCons(tmp10033, tmp10034)

tmp10036 := PrimCons(tmp10030, tmp10035)

tmp10037 := PrimCons(symif, tmp10036)

tmp10038 := PrimCons(tmp10037, Nil)

tmp10039 := PrimCons(tmp10026, tmp10038)

tmp10040 := PrimCons(W2031, tmp10039)

__e.Return(PrimCons(symlet, tmp10040))
return


}, 1)

tmp10041 := Call(__e, PrimFunc(symgensym), symTm)


__e.TailApply(tmp10022, tmp10041)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.atom-case-minus")))
return
}


}, 4)

tmp10047 := Call(__e, ns2_1set, symshen_4atom_1case_1minus, tmp10021)


_ = tmp10047

tmp10048 := MakeNative(func(__e *ControlFlow) {
V2040 := __e.Get(1)
_ = V2040
V2041 := __e.Get(2)
_ = V2041
V2042 := __e.Get(3)
_ = V2042
V2043 := __e.Get(4)
_ = V2043
tmp10113 := PrimIsPair(V2040)

var ifres10083 Obj

if True == tmp10113 {
tmp10111 := PrimHead(V2040)

tmp10112 := PrimIsPair(tmp10111)

var ifres10085 Obj

if True == tmp10112 {
tmp10108 := PrimHead(V2040)

tmp10109 := PrimHead(tmp10108)

tmp10110 := PrimEqual(symcons, tmp10109)

var ifres10087 Obj

if True == tmp10110 {
tmp10105 := PrimHead(V2040)

tmp10106 := PrimTail(tmp10105)

tmp10107 := PrimIsPair(tmp10106)

var ifres10089 Obj

if True == tmp10107 {
tmp10101 := PrimHead(V2040)

tmp10102 := PrimTail(tmp10101)

tmp10103 := PrimTail(tmp10102)

tmp10104 := PrimIsPair(tmp10103)

var ifres10091 Obj

if True == tmp10104 {
tmp10096 := PrimHead(V2040)

tmp10097 := PrimTail(tmp10096)

tmp10098 := PrimTail(tmp10097)

tmp10099 := PrimTail(tmp10098)

tmp10100 := PrimEqual(Nil, tmp10099)

var ifres10093 Obj

if True == tmp10100 {
tmp10095 := PrimIsPair(V2041)

var ifres10094 Obj

if True == tmp10095 {
ifres10094 = True


} else {
ifres10094 = False


}

ifres10093 = ifres10094


} else {
ifres10093 = False


}

var ifres10092 Obj

if True == ifres10093 {
ifres10092 = True


} else {
ifres10092 = False


}

ifres10091 = ifres10092


} else {
ifres10091 = False


}

var ifres10090 Obj

if True == ifres10091 {
ifres10090 = True


} else {
ifres10090 = False


}

ifres10089 = ifres10090


} else {
ifres10089 = False


}

var ifres10088 Obj

if True == ifres10089 {
ifres10088 = True


} else {
ifres10088 = False


}

ifres10087 = ifres10088


} else {
ifres10087 = False


}

var ifres10086 Obj

if True == ifres10087 {
ifres10086 = True


} else {
ifres10086 = False


}

ifres10085 = ifres10086


} else {
ifres10085 = False


}

var ifres10084 Obj

if True == ifres10085 {
ifres10084 = True


} else {
ifres10084 = False


}

ifres10083 = ifres10084


} else {
ifres10083 = False


}

if True == ifres10083 {
tmp10049 := MakeNative(func(__e *ControlFlow) {
W2044 := __e.Get(1)
_ = W2044
tmp10050 := PrimHead(V2041)

tmp10051 := PrimCons(V2042, Nil)

tmp10052 := PrimCons(tmp10050, tmp10051)

tmp10053 := PrimCons(symshen_4lazyderef, tmp10052)

tmp10054 := PrimCons(W2044, Nil)

tmp10055 := PrimCons(symcons_2, tmp10054)

tmp10056 := PrimHead(V2040)

tmp10057 := PrimTail(tmp10056)

tmp10058 := PrimHead(tmp10057)

tmp10059 := PrimHead(V2040)

tmp10060 := PrimTail(tmp10059)

tmp10061 := PrimTail(tmp10060)

tmp10062 := PrimHead(tmp10061)

tmp10063 := PrimTail(V2040)

tmp10064 := PrimCons(tmp10062, tmp10063)

tmp10065 := PrimCons(tmp10058, tmp10064)

tmp10066 := PrimCons(W2044, Nil)

tmp10067 := PrimCons(symhd, tmp10066)

tmp10068 := PrimCons(W2044, Nil)

tmp10069 := PrimCons(symtl, tmp10068)

tmp10070 := PrimTail(V2041)

tmp10071 := PrimCons(tmp10069, tmp10070)

tmp10072 := PrimCons(tmp10067, tmp10071)

tmp10073 := Call(__e, PrimFunc(symshen_4compile_1head), symshen_4_1m, tmp10065, tmp10072, V2042, V2043)


tmp10074 := PrimCons(False, Nil)

tmp10075 := PrimCons(tmp10073, tmp10074)

tmp10076 := PrimCons(tmp10055, tmp10075)

tmp10077 := PrimCons(symif, tmp10076)

tmp10078 := PrimCons(tmp10077, Nil)

tmp10079 := PrimCons(tmp10053, tmp10078)

tmp10080 := PrimCons(W2044, tmp10079)

__e.Return(PrimCons(symlet, tmp10080))
return


}, 1)

tmp10081 := Call(__e, PrimFunc(symgensym), symTm)


__e.TailApply(tmp10049, tmp10081)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.cons-case-minus")))
return
}


}, 4)

tmp10114 := Call(__e, ns2_1set, symshen_4cons_1case_1minus, tmp10048)


_ = tmp10114

tmp10115 := MakeNative(func(__e *ControlFlow) {
V2053 := __e.Get(1)
_ = V2053
V2054 := __e.Get(2)
_ = V2054
V2055 := __e.Get(3)
_ = V2055
V2056 := __e.Get(4)
_ = V2056
tmp10161 := PrimIsPair(V2053)

var ifres10158 Obj

if True == tmp10161 {
tmp10160 := PrimIsPair(V2054)

var ifres10159 Obj

if True == tmp10160 {
ifres10159 = True


} else {
ifres10159 = False


}

ifres10158 = ifres10159


} else {
ifres10158 = False


}

if True == ifres10158 {
tmp10116 := MakeNative(func(__e *ControlFlow) {
W2057 := __e.Get(1)
_ = W2057
tmp10117 := MakeNative(func(__e *ControlFlow) {
W2058 := __e.Get(1)
_ = W2058
tmp10118 := PrimHead(V2054)

tmp10119 := PrimCons(V2055, Nil)

tmp10120 := PrimCons(tmp10118, tmp10119)

tmp10121 := PrimCons(symshen_4lazyderef, tmp10120)

tmp10122 := PrimTail(V2053)

tmp10123 := PrimTail(V2054)

tmp10124 := Call(__e, PrimFunc(symshen_4compile_1head), symshen_4_7m, tmp10122, tmp10123, V2055, V2056)


tmp10125 := PrimCons(tmp10124, Nil)

tmp10126 := PrimCons(symfreeze, tmp10125)

tmp10127 := PrimHead(V2053)

tmp10128 := PrimCons(tmp10127, Nil)

tmp10129 := PrimCons(W2057, tmp10128)

tmp10130 := PrimCons(sym_a, tmp10129)

tmp10131 := PrimCons(W2058, Nil)

tmp10132 := PrimCons(symthaw, tmp10131)

tmp10133 := PrimCons(W2057, Nil)

tmp10134 := PrimCons(symshen_4pvar_2, tmp10133)

tmp10135 := PrimHead(V2053)

tmp10136 := Call(__e, PrimFunc(symshen_4demode), tmp10135)


tmp10137 := PrimCons(W2058, Nil)

tmp10138 := PrimCons(V2055, tmp10137)

tmp10139 := PrimCons(tmp10136, tmp10138)

tmp10140 := PrimCons(W2057, tmp10139)

tmp10141 := PrimCons(symshen_4bind_b, tmp10140)

tmp10142 := PrimCons(False, Nil)

tmp10143 := PrimCons(tmp10141, tmp10142)

tmp10144 := PrimCons(tmp10134, tmp10143)

tmp10145 := PrimCons(symif, tmp10144)

tmp10146 := PrimCons(tmp10145, Nil)

tmp10147 := PrimCons(tmp10132, tmp10146)

tmp10148 := PrimCons(tmp10130, tmp10147)

tmp10149 := PrimCons(symif, tmp10148)

tmp10150 := PrimCons(tmp10149, Nil)

tmp10151 := PrimCons(tmp10126, tmp10150)

tmp10152 := PrimCons(W2058, tmp10151)

tmp10153 := PrimCons(tmp10121, tmp10152)

tmp10154 := PrimCons(W2057, tmp10153)

__e.Return(PrimCons(symlet, tmp10154))
return


}, 1)

tmp10155 := Call(__e, PrimFunc(symgensym), symGoTo)


__e.TailApply(tmp10117, tmp10155)
return


}, 1)

tmp10156 := Call(__e, PrimFunc(symgensym), symTm)


__e.TailApply(tmp10116, tmp10156)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.atom-case-plus")))
return
}


}, 4)

tmp10162 := Call(__e, ns2_1set, symshen_4atom_1case_1plus, tmp10115)


_ = tmp10162

tmp10163 := MakeNative(func(__e *ControlFlow) {
V2067 := __e.Get(1)
_ = V2067
V2068 := __e.Get(2)
_ = V2068
V2069 := __e.Get(3)
_ = V2069
V2070 := __e.Get(4)
_ = V2070
tmp10259 := PrimIsPair(V2067)

var ifres10229 Obj

if True == tmp10259 {
tmp10257 := PrimHead(V2067)

tmp10258 := PrimIsPair(tmp10257)

var ifres10231 Obj

if True == tmp10258 {
tmp10254 := PrimHead(V2067)

tmp10255 := PrimHead(tmp10254)

tmp10256 := PrimEqual(symcons, tmp10255)

var ifres10233 Obj

if True == tmp10256 {
tmp10251 := PrimHead(V2067)

tmp10252 := PrimTail(tmp10251)

tmp10253 := PrimIsPair(tmp10252)

var ifres10235 Obj

if True == tmp10253 {
tmp10247 := PrimHead(V2067)

tmp10248 := PrimTail(tmp10247)

tmp10249 := PrimTail(tmp10248)

tmp10250 := PrimIsPair(tmp10249)

var ifres10237 Obj

if True == tmp10250 {
tmp10242 := PrimHead(V2067)

tmp10243 := PrimTail(tmp10242)

tmp10244 := PrimTail(tmp10243)

tmp10245 := PrimTail(tmp10244)

tmp10246 := PrimEqual(Nil, tmp10245)

var ifres10239 Obj

if True == tmp10246 {
tmp10241 := PrimIsPair(V2068)

var ifres10240 Obj

if True == tmp10241 {
ifres10240 = True


} else {
ifres10240 = False


}

ifres10239 = ifres10240


} else {
ifres10239 = False


}

var ifres10238 Obj

if True == ifres10239 {
ifres10238 = True


} else {
ifres10238 = False


}

ifres10237 = ifres10238


} else {
ifres10237 = False


}

var ifres10236 Obj

if True == ifres10237 {
ifres10236 = True


} else {
ifres10236 = False


}

ifres10235 = ifres10236


} else {
ifres10235 = False


}

var ifres10234 Obj

if True == ifres10235 {
ifres10234 = True


} else {
ifres10234 = False


}

ifres10233 = ifres10234


} else {
ifres10233 = False


}

var ifres10232 Obj

if True == ifres10233 {
ifres10232 = True


} else {
ifres10232 = False


}

ifres10231 = ifres10232


} else {
ifres10231 = False


}

var ifres10230 Obj

if True == ifres10231 {
ifres10230 = True


} else {
ifres10230 = False


}

ifres10229 = ifres10230


} else {
ifres10229 = False


}

if True == ifres10229 {
tmp10164 := MakeNative(func(__e *ControlFlow) {
W2071 := __e.Get(1)
_ = W2071
tmp10165 := MakeNative(func(__e *ControlFlow) {
W2072 := __e.Get(1)
_ = W2072
tmp10166 := MakeNative(func(__e *ControlFlow) {
W2073 := __e.Get(1)
_ = W2073
tmp10167 := MakeNative(func(__e *ControlFlow) {
W2074 := __e.Get(1)
_ = W2074
tmp10168 := MakeNative(func(__e *ControlFlow) {
W2075 := __e.Get(1)
_ = W2075
tmp10169 := PrimHead(V2068)

tmp10170 := PrimCons(V2069, Nil)

tmp10171 := PrimCons(tmp10169, tmp10170)

tmp10172 := PrimCons(symshen_4lazyderef, tmp10171)

tmp10173 := PrimTail(V2067)

tmp10174 := PrimTail(V2068)

tmp10175 := Call(__e, PrimFunc(symshen_4compile_1head), symshen_4_7m, tmp10173, tmp10174, V2069, V2070)


tmp10176 := Call(__e, PrimFunc(symshen_4goto), W2073, tmp10175)


tmp10177 := PrimCons(W2071, Nil)

tmp10178 := PrimCons(symcons_2, tmp10177)

tmp10179 := PrimHead(V2067)

tmp10180 := PrimTail(tmp10179)

tmp10181 := PrimCons(W2071, Nil)

tmp10182 := PrimCons(symhd, tmp10181)

tmp10183 := PrimCons(W2071, Nil)

tmp10184 := PrimCons(symtl, tmp10183)

tmp10185 := PrimCons(tmp10184, Nil)

tmp10186 := PrimCons(tmp10182, tmp10185)

tmp10187 := Call(__e, PrimFunc(symshen_4invoke), W2072, W2073)


tmp10188 := Call(__e, PrimFunc(symshen_4compile_1head), symshen_4_7m, tmp10180, tmp10186, V2069, tmp10187)


tmp10189 := PrimCons(W2071, Nil)

tmp10190 := PrimCons(symshen_4pvar_2, tmp10189)

tmp10191 := Call(__e, PrimFunc(symshen_4demode), W2074)


tmp10192 := Call(__e, PrimFunc(symshen_4invoke), W2072, W2073)


tmp10193 := PrimCons(tmp10192, Nil)

tmp10194 := PrimCons(symfreeze, tmp10193)

tmp10195 := PrimCons(tmp10194, Nil)

tmp10196 := PrimCons(V2069, tmp10195)

tmp10197 := PrimCons(tmp10191, tmp10196)

tmp10198 := PrimCons(W2071, tmp10197)

tmp10199 := PrimCons(symshen_4bind_b, tmp10198)

tmp10200 := Call(__e, PrimFunc(symshen_4stpart), W2075, tmp10199, V2069)


tmp10201 := PrimCons(False, Nil)

tmp10202 := PrimCons(tmp10200, tmp10201)

tmp10203 := PrimCons(tmp10190, tmp10202)

tmp10204 := PrimCons(symif, tmp10203)

tmp10205 := PrimCons(tmp10204, Nil)

tmp10206 := PrimCons(tmp10188, tmp10205)

tmp10207 := PrimCons(tmp10178, tmp10206)

tmp10208 := PrimCons(symif, tmp10207)

tmp10209 := PrimCons(tmp10208, Nil)

tmp10210 := PrimCons(tmp10176, tmp10209)

tmp10211 := PrimCons(W2072, tmp10210)

tmp10212 := PrimCons(tmp10172, tmp10211)

tmp10213 := PrimCons(W2071, tmp10212)

__e.Return(PrimCons(symlet, tmp10213))
return


}, 1)

tmp10214 := Call(__e, PrimFunc(symshen_4extract_1vars), W2074)


__e.TailApply(tmp10168, tmp10214)
return


}, 1)

tmp10215 := PrimHead(V2067)

tmp10216 := Call(__e, PrimFunc(symshen_4tame), tmp10215)


__e.TailApply(tmp10167, tmp10216)
return


}, 1)

tmp10217 := PrimHead(V2067)

tmp10218 := PrimTail(tmp10217)

tmp10219 := PrimHead(tmp10218)

tmp10220 := PrimHead(V2067)

tmp10221 := PrimTail(tmp10220)

tmp10222 := PrimTail(tmp10221)

tmp10223 := PrimHead(tmp10222)

tmp10224 := PrimCons(tmp10219, tmp10223)

tmp10225 := Call(__e, PrimFunc(symshen_4extract_1vars), tmp10224)


__e.TailApply(tmp10166, tmp10225)
return


}, 1)

tmp10226 := Call(__e, PrimFunc(symgensym), symGoTo)


__e.TailApply(tmp10165, tmp10226)
return


}, 1)

tmp10227 := Call(__e, PrimFunc(symgensym), symTm)


__e.TailApply(tmp10164, tmp10227)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.cons-case-plus")))
return
}


}, 4)

tmp10260 := Call(__e, ns2_1set, symshen_4cons_1case_1plus, tmp10163)


_ = tmp10260

tmp10261 := MakeNative(func(__e *ControlFlow) {
V2076 := __e.Get(1)
_ = V2076
tmp10298 := PrimIsPair(V2076)

var ifres10285 Obj

if True == tmp10298 {
tmp10296 := PrimHead(V2076)

tmp10297 := PrimEqual(symshen_4_7m, tmp10296)

var ifres10287 Obj

if True == tmp10297 {
tmp10294 := PrimTail(V2076)

tmp10295 := PrimIsPair(tmp10294)

var ifres10289 Obj

if True == tmp10295 {
tmp10291 := PrimTail(V2076)

tmp10292 := PrimTail(tmp10291)

tmp10293 := PrimEqual(Nil, tmp10292)

var ifres10290 Obj

if True == tmp10293 {
ifres10290 = True


} else {
ifres10290 = False


}

ifres10289 = ifres10290


} else {
ifres10289 = False


}

var ifres10288 Obj

if True == ifres10289 {
ifres10288 = True


} else {
ifres10288 = False


}

ifres10287 = ifres10288


} else {
ifres10287 = False


}

var ifres10286 Obj

if True == ifres10287 {
ifres10286 = True


} else {
ifres10286 = False


}

ifres10285 = ifres10286


} else {
ifres10285 = False


}

if True == ifres10285 {
tmp10262 := PrimTail(V2076)

tmp10263 := PrimHead(tmp10262)

__e.TailApply(PrimFunc(symshen_4demode), tmp10263)
return


} else {
tmp10283 := PrimIsPair(V2076)

var ifres10270 Obj

if True == tmp10283 {
tmp10281 := PrimHead(V2076)

tmp10282 := PrimEqual(symshen_4_1m, tmp10281)

var ifres10272 Obj

if True == tmp10282 {
tmp10279 := PrimTail(V2076)

tmp10280 := PrimIsPair(tmp10279)

var ifres10274 Obj

if True == tmp10280 {
tmp10276 := PrimTail(V2076)

tmp10277 := PrimTail(tmp10276)

tmp10278 := PrimEqual(Nil, tmp10277)

var ifres10275 Obj

if True == tmp10278 {
ifres10275 = True


} else {
ifres10275 = False


}

ifres10274 = ifres10275


} else {
ifres10274 = False


}

var ifres10273 Obj

if True == ifres10274 {
ifres10273 = True


} else {
ifres10273 = False


}

ifres10272 = ifres10273


} else {
ifres10272 = False


}

var ifres10271 Obj

if True == ifres10272 {
ifres10271 = True


} else {
ifres10271 = False


}

ifres10270 = ifres10271


} else {
ifres10270 = False


}

if True == ifres10270 {
tmp10264 := PrimTail(V2076)

tmp10265 := PrimHead(tmp10264)

__e.TailApply(PrimFunc(symshen_4demode), tmp10265)
return


} else {
tmp10268 := PrimIsPair(V2076)

if True == tmp10268 {
tmp10266 := MakeNative(func(__e *ControlFlow) {
Z2077 := __e.Get(1)
_ = Z2077
__e.TailApply(PrimFunc(symshen_4demode), Z2077)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp10266, V2076)
return


} else {
__e.Return(V2076)
return
}


}


}


}, 1)

tmp10299 := Call(__e, ns2_1set, symshen_4demode, tmp10261)


_ = tmp10299

tmp10300 := MakeNative(func(__e *ControlFlow) {
V2078 := __e.Get(1)
_ = V2078
tmp10305 := Call(__e, PrimFunc(symshen_4wildcard_2), V2078)


if True == tmp10305 {
__e.TailApply(PrimFunc(symgensym), symY)
return
} else {
tmp10303 := PrimIsPair(V2078)

if True == tmp10303 {
tmp10301 := MakeNative(func(__e *ControlFlow) {
Z2079 := __e.Get(1)
_ = Z2079
__e.TailApply(PrimFunc(symshen_4tame), Z2079)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp10301, V2078)
return


} else {
__e.Return(V2078)
return
}


}


}, 1)

tmp10306 := Call(__e, ns2_1set, symshen_4tame, tmp10300)


_ = tmp10306

tmp10307 := MakeNative(func(__e *ControlFlow) {
V2080 := __e.Get(1)
_ = V2080
V2081 := __e.Get(2)
_ = V2081
tmp10310 := PrimEqual(Nil, V2080)

if True == tmp10310 {
tmp10308 := PrimCons(V2081, Nil)

__e.Return(PrimCons(symfreeze, tmp10308))
return


} else {
__e.TailApply(PrimFunc(symshen_4goto_1h), V2080, V2081)
return
}


}, 2)

tmp10311 := Call(__e, ns2_1set, symshen_4goto, tmp10307)


_ = tmp10311

tmp10312 := MakeNative(func(__e *ControlFlow) {
V2082 := __e.Get(1)
_ = V2082
V2083 := __e.Get(2)
_ = V2083
tmp10321 := PrimEqual(Nil, V2082)

if True == tmp10321 {
__e.Return(V2083)
return
} else {
tmp10319 := PrimIsPair(V2082)

if True == tmp10319 {
tmp10313 := PrimHead(V2082)

tmp10314 := PrimTail(V2082)

tmp10315 := Call(__e, PrimFunc(symshen_4goto_1h), tmp10314, V2083)


tmp10316 := PrimCons(tmp10315, Nil)

tmp10317 := PrimCons(tmp10313, tmp10316)

__e.Return(PrimCons(symlambda, tmp10317))
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4goto_1h)
return
}


}


}, 2)

tmp10322 := Call(__e, ns2_1set, symshen_4goto_1h, tmp10312)


_ = tmp10322

tmp10323 := MakeNative(func(__e *ControlFlow) {
V2084 := __e.Get(1)
_ = V2084
V2085 := __e.Get(2)
_ = V2085
tmp10326 := PrimEqual(Nil, V2085)

if True == tmp10326 {
tmp10324 := PrimCons(V2084, Nil)

__e.Return(PrimCons(symthaw, tmp10324))
return


} else {
__e.Return(PrimCons(V2084, V2085))
return
}


}, 2)

tmp10327 := Call(__e, ns2_1set, symshen_4invoke, tmp10323)


_ = tmp10327

tmp10328 := MakeNative(func(__e *ControlFlow) {
V2086 := __e.Get(1)
_ = V2086
__e.Return(PrimEqual(V2086, sym__))
return
}, 1)

tmp10329 := Call(__e, ns2_1set, symshen_4wildcard_2, tmp10328)


_ = tmp10329

tmp10330 := MakeNative(func(__e *ControlFlow) {
V2087 := __e.Get(1)
_ = V2087
tmp10337 := PrimIsVector(V2087)

if True == tmp10337 {
tmp10332 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimVectorGet(V2087, MakeNumber(0)))
return
}, 0)

tmp10333 := MakeNative(func(__e *ControlFlow) {
Z2088 := __e.Get(1)
_ = Z2088
__e.Return(symshen_4not_1pvar)
return
}, 1)

tmp10334 := Call(__e, try_1catch, tmp10332, tmp10333)


tmp10335 := PrimEqual(tmp10334, symshen_4pvar)

if True == tmp10335 {
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

tmp10338 := Call(__e, ns2_1set, symshen_4pvar_2, tmp10330)


_ = tmp10338

tmp10339 := MakeNative(func(__e *ControlFlow) {
V2089 := __e.Get(1)
_ = V2089
V2090 := __e.Get(2)
_ = V2090
tmp10346 := Call(__e, PrimFunc(symshen_4pvar_2), V2089)


if True == tmp10346 {
tmp10340 := MakeNative(func(__e *ControlFlow) {
W2091 := __e.Get(1)
_ = W2091
tmp10342 := PrimEqual(W2091, symshen_4_1null_1)

if True == tmp10342 {
__e.Return(V2089)
return
} else {
__e.TailApply(PrimFunc(symshen_4lazyderef), W2091, V2090)
return
}


}, 1)

tmp10343 := PrimVectorGet(V2089, MakeNumber(1))

tmp10344 := PrimVectorGet(V2090, tmp10343)

__e.TailApply(tmp10340, tmp10344)
return


} else {
__e.Return(V2089)
return
}


}, 2)

tmp10347 := Call(__e, ns2_1set, symshen_4lazyderef, tmp10339)


_ = tmp10347

tmp10348 := MakeNative(func(__e *ControlFlow) {
V2092 := __e.Get(1)
_ = V2092
V2093 := __e.Get(2)
_ = V2093
tmp10361 := PrimIsPair(V2092)

if True == tmp10361 {
tmp10349 := PrimHead(V2092)

tmp10350 := Call(__e, PrimFunc(symshen_4deref), tmp10349, V2093)


tmp10351 := PrimTail(V2092)

tmp10352 := Call(__e, PrimFunc(symshen_4deref), tmp10351, V2093)


__e.Return(PrimCons(tmp10350, tmp10352))
return


} else {
tmp10359 := Call(__e, PrimFunc(symshen_4pvar_2), V2092)


if True == tmp10359 {
tmp10353 := MakeNative(func(__e *ControlFlow) {
W2094 := __e.Get(1)
_ = W2094
tmp10355 := PrimEqual(W2094, symshen_4_1null_1)

if True == tmp10355 {
__e.Return(V2092)
return
} else {
__e.TailApply(PrimFunc(symshen_4deref), W2094, V2093)
return
}


}, 1)

tmp10356 := PrimVectorGet(V2092, MakeNumber(1))

tmp10357 := PrimVectorGet(V2093, tmp10356)

__e.TailApply(tmp10353, tmp10357)
return


} else {
__e.Return(V2092)
return
}


}


}, 2)

tmp10362 := Call(__e, ns2_1set, symshen_4deref, tmp10348)


_ = tmp10362

tmp10363 := MakeNative(func(__e *ControlFlow) {
V2095 := __e.Get(1)
_ = V2095
V2096 := __e.Get(2)
_ = V2096
V2097 := __e.Get(3)
_ = V2097
V2098 := __e.Get(4)
_ = V2098
tmp10364 := MakeNative(func(__e *ControlFlow) {
W2099 := __e.Get(1)
_ = W2099
tmp10365 := MakeNative(func(__e *ControlFlow) {
W2100 := __e.Get(1)
_ = W2100
tmp10367 := PrimEqual(W2100, False)

if True == tmp10367 {
__e.TailApply(PrimFunc(symshen_4unwind), V2095, V2097, W2100)
return
} else {
__e.Return(W2100)
return
}


}, 1)

tmp10368 := Call(__e, PrimFunc(symthaw), V2098)


__e.TailApply(tmp10365, tmp10368)
return


}, 1)

tmp10369 := Call(__e, PrimFunc(symshen_4bindv), V2095, V2096, V2097)


__e.TailApply(tmp10364, tmp10369)
return


}, 4)

tmp10370 := Call(__e, ns2_1set, symshen_4bind_b, tmp10363)


_ = tmp10370

tmp10371 := MakeNative(func(__e *ControlFlow) {
V2101 := __e.Get(1)
_ = V2101
V2102 := __e.Get(2)
_ = V2102
V2103 := __e.Get(3)
_ = V2103
tmp10372 := PrimVectorGet(V2101, MakeNumber(1))

__e.Return(PrimVectorSet(V2103, tmp10372, V2102))
return


}, 3)

tmp10373 := Call(__e, ns2_1set, symshen_4bindv, tmp10371)


_ = tmp10373

tmp10374 := MakeNative(func(__e *ControlFlow) {
V2104 := __e.Get(1)
_ = V2104
V2105 := __e.Get(2)
_ = V2105
V2106 := __e.Get(3)
_ = V2106
tmp10375 := PrimVectorGet(V2104, MakeNumber(1))

tmp10376 := PrimVectorSet(V2105, tmp10375, symshen_4_1null_1)

_ = tmp10376

__e.Return(V2106)
return


}, 3)

tmp10377 := Call(__e, ns2_1set, symshen_4unwind, tmp10374)


_ = tmp10377

tmp10378 := MakeNative(func(__e *ControlFlow) {
V2115 := __e.Get(1)
_ = V2115
V2116 := __e.Get(2)
_ = V2116
V2117 := __e.Get(3)
_ = V2117
tmp10393 := PrimEqual(Nil, V2115)

if True == tmp10393 {
__e.Return(V2116)
return
} else {
tmp10391 := PrimIsPair(V2115)

if True == tmp10391 {
tmp10379 := PrimHead(V2115)

tmp10380 := PrimCons(V2117, Nil)

tmp10381 := PrimCons(symshen_4newpv, tmp10380)

tmp10382 := PrimTail(V2115)

tmp10383 := Call(__e, PrimFunc(symshen_4stpart), tmp10382, V2116, V2117)


tmp10384 := PrimCons(tmp10383, Nil)

tmp10385 := PrimCons(V2117, tmp10384)

tmp10386 := PrimCons(symshen_4gc, tmp10385)

tmp10387 := PrimCons(tmp10386, Nil)

tmp10388 := PrimCons(tmp10381, tmp10387)

tmp10389 := PrimCons(tmp10379, tmp10388)

__e.Return(PrimCons(symlet, tmp10389))
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.stpart")))
return
}


}


}, 3)

tmp10394 := Call(__e, ns2_1set, symshen_4stpart, tmp10378)


_ = tmp10394

tmp10395 := MakeNative(func(__e *ControlFlow) {
V2118 := __e.Get(1)
_ = V2118
V2119 := __e.Get(2)
_ = V2119
tmp10400 := PrimEqual(V2119, False)

if True == tmp10400 {
tmp10396 := MakeNative(func(__e *ControlFlow) {
W2120 := __e.Get(1)
_ = W2120
tmp10397 := Call(__e, PrimFunc(symshen_4decrement_1ticket), W2120, V2118)


_ = tmp10397

__e.Return(V2119)
return


}, 1)

tmp10398 := Call(__e, PrimFunc(symshen_4ticket_1number), V2118)


__e.TailApply(tmp10396, tmp10398)
return


} else {
__e.Return(V2119)
return
}


}, 2)

tmp10401 := Call(__e, ns2_1set, symshen_4gc, tmp10395)


_ = tmp10401

tmp10402 := MakeNative(func(__e *ControlFlow) {
V2121 := __e.Get(1)
_ = V2121
V2122 := __e.Get(2)
_ = V2122
tmp10403 := PrimNumberSubtract(V2121, MakeNumber(1))

__e.Return(PrimVectorSet(V2122, MakeNumber(1), tmp10403))
return


}, 2)

tmp10404 := Call(__e, ns2_1set, symshen_4decrement_1ticket, tmp10402)


_ = tmp10404

tmp10405 := MakeNative(func(__e *ControlFlow) {
V2123 := __e.Get(1)
_ = V2123
tmp10406 := MakeNative(func(__e *ControlFlow) {
W2124 := __e.Get(1)
_ = W2124
tmp10407 := MakeNative(func(__e *ControlFlow) {
W2125 := __e.Get(1)
_ = W2125
tmp10408 := MakeNative(func(__e *ControlFlow) {
W2126 := __e.Get(1)
_ = W2126
__e.Return(W2125)
return
}, 1)

tmp10409 := Call(__e, PrimFunc(symshen_4nextticket), V2123, W2124)


__e.TailApply(tmp10408, tmp10409)
return


}, 1)

tmp10410 := Call(__e, PrimFunc(symshen_4make_1prolog_1variable), W2124)


__e.TailApply(tmp10407, tmp10410)
return


}, 1)

tmp10411 := Call(__e, PrimFunc(symshen_4ticket_1number), V2123)


__e.TailApply(tmp10406, tmp10411)
return


}, 1)

tmp10412 := Call(__e, ns2_1set, symshen_4newpv, tmp10405)


_ = tmp10412

tmp10413 := MakeNative(func(__e *ControlFlow) {
V2127 := __e.Get(1)
_ = V2127
__e.Return(PrimVectorGet(V2127, MakeNumber(1)))
return
}, 1)

tmp10414 := Call(__e, ns2_1set, symshen_4ticket_1number, tmp10413)


_ = tmp10414

tmp10415 := MakeNative(func(__e *ControlFlow) {
V2128 := __e.Get(1)
_ = V2128
V2129 := __e.Get(2)
_ = V2129
tmp10416 := MakeNative(func(__e *ControlFlow) {
W2130 := __e.Get(1)
_ = W2130
tmp10417 := PrimNumberAdd(V2129, MakeNumber(1))

__e.Return(PrimVectorSet(W2130, MakeNumber(1), tmp10417))
return


}, 1)

tmp10418 := PrimVectorSet(V2128, V2129, symshen_4_1null_1)

__e.TailApply(tmp10416, tmp10418)
return


}, 2)

tmp10419 := Call(__e, ns2_1set, symshen_4nextticket, tmp10415)


_ = tmp10419

tmp10420 := MakeNative(func(__e *ControlFlow) {
V2131 := __e.Get(1)
_ = V2131
tmp10421 := PrimAbsvector(MakeNumber(2))

tmp10422 := PrimVectorSet(tmp10421, MakeNumber(0), symshen_4pvar)

__e.Return(PrimVectorSet(tmp10422, MakeNumber(1), V2131))
return


}, 1)

tmp10423 := Call(__e, ns2_1set, symshen_4make_1prolog_1variable, tmp10420)


_ = tmp10423

tmp10424 := MakeNative(func(__e *ControlFlow) {
V2132 := __e.Get(1)
_ = V2132
tmp10425 := PrimVectorGet(V2132, MakeNumber(1))

tmp10426 := Call(__e, PrimFunc(symshen_4app), tmp10425, MakeString(""), symshen_4a)


__e.Return(PrimStringConcat(MakeString("Var"), tmp10426))
return


}, 1)

tmp10427 := Call(__e, ns2_1set, symshen_4pvar, tmp10424)


_ = tmp10427

tmp10428 := MakeNative(func(__e *ControlFlow) {
tmp10429 := PrimValue(symshen_4_dinfs_d)

tmp10430 := PrimNumberAdd(MakeNumber(1), tmp10429)

__e.Return(PrimSet(symshen_4_dinfs_d, tmp10430))
return


}, 0)

tmp10431 := Call(__e, ns2_1set, symshen_4incinfs, tmp10428)


_ = tmp10431

tmp10432 := MakeNative(func(__e *ControlFlow) {
V2133 := __e.Get(1)
_ = V2133
tmp10439 := PrimIsInteger(V2133)

var ifres10436 Obj

if True == tmp10439 {
tmp10438 := PrimGreatThan(V2133, MakeNumber(0))

var ifres10437 Obj

if True == tmp10438 {
ifres10437 = True


} else {
ifres10437 = False


}

ifres10436 = ifres10437


} else {
ifres10436 = False


}

if True == ifres10436 {
__e.Return(PrimSet(symshen_4_dsize_1prolog_1vector_d, V2133))
return
} else {
tmp10433 := Call(__e, PrimFunc(symshen_4app), V2133, MakeString(""), symshen_4a)


tmp10434 := PrimStringConcat(MakeString("prolog vector size: size should be a positive integer; not "), tmp10433)

__e.Return(PrimSimpleError(tmp10434))
return


}


}, 1)

tmp10440 := Call(__e, ns2_1set, symshen_4prolog_1vector_1size, tmp10432)


_ = tmp10440

tmp10441 := MakeNative(func(__e *ControlFlow) {
V2145 := __e.Get(1)
_ = V2145
V2146 := __e.Get(2)
_ = V2146
V2147 := __e.Get(3)
_ = V2147
V2148 := __e.Get(4)
_ = V2148
tmp10471 := PrimEqual(V2145, V2146)

if True == tmp10471 {
__e.TailApply(PrimFunc(symthaw), V2148)
return
} else {
tmp10469 := Call(__e, PrimFunc(symshen_4pvar_2), V2145)


var ifres10464 Obj

if True == tmp10469 {
tmp10466 := Call(__e, PrimFunc(symshen_4deref), V2146, V2147)


tmp10467 := Call(__e, PrimFunc(symshen_4occurs_1check_2), V2145, tmp10466)


tmp10468 := PrimNot(tmp10467)

var ifres10465 Obj

if True == tmp10468 {
ifres10465 = True


} else {
ifres10465 = False


}

ifres10464 = ifres10465


} else {
ifres10464 = False


}

if True == ifres10464 {
__e.TailApply(PrimFunc(symshen_4bind_b), V2145, V2146, V2147, V2148)
return
} else {
tmp10462 := Call(__e, PrimFunc(symshen_4pvar_2), V2146)


var ifres10457 Obj

if True == tmp10462 {
tmp10459 := Call(__e, PrimFunc(symshen_4deref), V2145, V2147)


tmp10460 := Call(__e, PrimFunc(symshen_4occurs_1check_2), V2146, tmp10459)


tmp10461 := PrimNot(tmp10460)

var ifres10458 Obj

if True == tmp10461 {
ifres10458 = True


} else {
ifres10458 = False


}

ifres10457 = ifres10458


} else {
ifres10457 = False


}

if True == ifres10457 {
__e.TailApply(PrimFunc(symshen_4bind_b), V2146, V2145, V2147, V2148)
return
} else {
tmp10455 := PrimIsPair(V2145)

var ifres10452 Obj

if True == tmp10455 {
tmp10454 := PrimIsPair(V2146)

var ifres10453 Obj

if True == tmp10454 {
ifres10453 = True


} else {
ifres10453 = False


}

ifres10452 = ifres10453


} else {
ifres10452 = False


}

if True == ifres10452 {
tmp10442 := PrimHead(V2145)

tmp10443 := Call(__e, PrimFunc(symshen_4lazyderef), tmp10442, V2147)


tmp10444 := PrimHead(V2146)

tmp10445 := Call(__e, PrimFunc(symshen_4lazyderef), tmp10444, V2147)


tmp10446 := MakeNative(func(__e *ControlFlow) {
tmp10447 := PrimTail(V2145)

tmp10448 := Call(__e, PrimFunc(symshen_4lazyderef), tmp10447, V2147)


tmp10449 := PrimTail(V2146)

tmp10450 := Call(__e, PrimFunc(symshen_4lazyderef), tmp10449, V2147)


__e.TailApply(PrimFunc(symshen_4lzy_a_b), tmp10448, tmp10450, V2147, V2148)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4lzy_a_b), tmp10443, tmp10445, V2147, tmp10446)
return


} else {
__e.Return(False)
return
}


}


}


}


}, 4)

tmp10472 := Call(__e, ns2_1set, symshen_4lzy_a_b, tmp10441)


_ = tmp10472

tmp10473 := MakeNative(func(__e *ControlFlow) {
V2160 := __e.Get(1)
_ = V2160
V2161 := __e.Get(2)
_ = V2161
V2162 := __e.Get(3)
_ = V2162
V2163 := __e.Get(4)
_ = V2163
tmp10493 := PrimEqual(V2160, V2161)

if True == tmp10493 {
__e.TailApply(PrimFunc(symthaw), V2163)
return
} else {
tmp10491 := Call(__e, PrimFunc(symshen_4pvar_2), V2160)


if True == tmp10491 {
__e.TailApply(PrimFunc(symshen_4bind_b), V2160, V2161, V2162, V2163)
return
} else {
tmp10489 := Call(__e, PrimFunc(symshen_4pvar_2), V2161)


if True == tmp10489 {
__e.TailApply(PrimFunc(symshen_4bind_b), V2161, V2160, V2162, V2163)
return
} else {
tmp10487 := PrimIsPair(V2160)

var ifres10484 Obj

if True == tmp10487 {
tmp10486 := PrimIsPair(V2161)

var ifres10485 Obj

if True == tmp10486 {
ifres10485 = True


} else {
ifres10485 = False


}

ifres10484 = ifres10485


} else {
ifres10484 = False


}

if True == ifres10484 {
tmp10474 := PrimHead(V2160)

tmp10475 := Call(__e, PrimFunc(symshen_4lazyderef), tmp10474, V2162)


tmp10476 := PrimHead(V2161)

tmp10477 := Call(__e, PrimFunc(symshen_4lazyderef), tmp10476, V2162)


tmp10478 := MakeNative(func(__e *ControlFlow) {
tmp10479 := PrimTail(V2160)

tmp10480 := Call(__e, PrimFunc(symshen_4lazyderef), tmp10479, V2162)


tmp10481 := PrimTail(V2161)

tmp10482 := Call(__e, PrimFunc(symshen_4lazyderef), tmp10481, V2162)


__e.TailApply(PrimFunc(symshen_4lzy_a), tmp10480, tmp10482, V2162, V2163)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4lzy_a), tmp10475, tmp10477, V2162, tmp10478)
return


} else {
__e.Return(False)
return
}


}


}


}


}, 4)

tmp10494 := Call(__e, ns2_1set, symshen_4lzy_a, tmp10473)


_ = tmp10494

tmp10495 := MakeNative(func(__e *ControlFlow) {
V2169 := __e.Get(1)
_ = V2169
V2170 := __e.Get(2)
_ = V2170
tmp10505 := PrimEqual(V2169, V2170)

if True == tmp10505 {
__e.Return(True)
return
} else {
tmp10503 := PrimIsPair(V2170)

if True == tmp10503 {
tmp10500 := PrimHead(V2170)

tmp10501 := Call(__e, PrimFunc(symshen_4occurs_1check_2), V2169, tmp10500)


if True == tmp10501 {
__e.Return(True)
return
} else {
tmp10497 := PrimTail(V2170)

tmp10498 := Call(__e, PrimFunc(symshen_4occurs_1check_2), V2169, tmp10497)


if True == tmp10498 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


} else {
__e.Return(False)
return
}


}


}, 2)

tmp10506 := Call(__e, ns2_1set, symshen_4occurs_1check_2, tmp10495)


_ = tmp10506

tmp10507 := MakeNative(func(__e *ControlFlow) {
V2171 := __e.Get(1)
_ = V2171
V2172 := __e.Get(2)
_ = V2172
V2173 := __e.Get(3)
_ = V2173
V2174 := __e.Get(4)
_ = V2174
V2175 := __e.Get(5)
_ = V2175
tmp10508 := Call(__e, V2171, V2172)


tmp10509 := Call(__e, tmp10508, V2173)


tmp10510 := Call(__e, tmp10509, V2174)


__e.TailApply(tmp10510, V2175)
return


}, 5)

tmp10511 := Call(__e, ns2_1set, symcall, tmp10507)


_ = tmp10511

tmp10512 := MakeNative(func(__e *ControlFlow) {
V2182 := __e.Get(1)
_ = V2182
V2183 := __e.Get(2)
_ = V2183
V2184 := __e.Get(3)
_ = V2184
V2185 := __e.Get(4)
_ = V2185
V2186 := __e.Get(5)
_ = V2186
__e.TailApply(PrimFunc(symshen_4deref), V2182, V2183)
return
}, 5)

tmp10513 := Call(__e, ns2_1set, symreturn, tmp10512)


_ = tmp10513

tmp10514 := MakeNative(func(__e *ControlFlow) {
V2193 := __e.Get(1)
_ = V2193
V2194 := __e.Get(2)
_ = V2194
V2195 := __e.Get(3)
_ = V2195
V2196 := __e.Get(4)
_ = V2196
V2197 := __e.Get(5)
_ = V2197
if True == V2193 {
__e.TailApply(PrimFunc(symthaw), V2197)
return
} else {
__e.Return(False)
return
}
}, 5)

tmp10516 := Call(__e, ns2_1set, symwhen, tmp10514)


_ = tmp10516

tmp10517 := MakeNative(func(__e *ControlFlow) {
V2198 := __e.Get(1)
_ = V2198
V2199 := __e.Get(2)
_ = V2199
V2200 := __e.Get(3)
_ = V2200
V2201 := __e.Get(4)
_ = V2201
V2202 := __e.Get(5)
_ = V2202
V2203 := __e.Get(6)
_ = V2203
tmp10518 := Call(__e, PrimFunc(symshen_4lazyderef), V2198, V2200)


tmp10519 := Call(__e, PrimFunc(symshen_4lazyderef), V2199, V2200)


__e.TailApply(PrimFunc(symshen_4lzy_a), tmp10518, tmp10519, V2200, V2203)
return


}, 6)

tmp10520 := Call(__e, ns2_1set, symis, tmp10517)


_ = tmp10520

tmp10521 := MakeNative(func(__e *ControlFlow) {
V2204 := __e.Get(1)
_ = V2204
V2205 := __e.Get(2)
_ = V2205
V2206 := __e.Get(3)
_ = V2206
V2207 := __e.Get(4)
_ = V2207
V2208 := __e.Get(5)
_ = V2208
V2209 := __e.Get(6)
_ = V2209
tmp10522 := Call(__e, PrimFunc(symshen_4lazyderef), V2204, V2206)


tmp10523 := Call(__e, PrimFunc(symshen_4lazyderef), V2205, V2206)


__e.TailApply(PrimFunc(symshen_4lzy_a_b), tmp10522, tmp10523, V2206, V2209)
return


}, 6)

tmp10524 := Call(__e, ns2_1set, symis_b, tmp10521)


_ = tmp10524

tmp10525 := MakeNative(func(__e *ControlFlow) {
V2214 := __e.Get(1)
_ = V2214
V2215 := __e.Get(2)
_ = V2215
V2216 := __e.Get(3)
_ = V2216
V2217 := __e.Get(4)
_ = V2217
V2218 := __e.Get(5)
_ = V2218
V2219 := __e.Get(6)
_ = V2219
__e.TailApply(PrimFunc(symshen_4bind_b), V2214, V2215, V2216, V2219)
return
}, 6)

tmp10526 := Call(__e, ns2_1set, symbind, tmp10525)


_ = tmp10526

tmp10527 := MakeNative(func(__e *ControlFlow) {
V2220 := __e.Get(1)
_ = V2220
V2221 := __e.Get(2)
_ = V2221
V2222 := __e.Get(3)
_ = V2222
V2223 := __e.Get(4)
_ = V2223
V2224 := __e.Get(5)
_ = V2224
tmp10529 := Call(__e, PrimFunc(symshen_4lazyderef), V2220, V2221)


tmp10530 := Call(__e, PrimFunc(symshen_4pvar_2), tmp10529)


if True == tmp10530 {
__e.TailApply(PrimFunc(symthaw), V2224)
return
} else {
__e.Return(False)
return
}


}, 5)

tmp10531 := Call(__e, ns2_1set, symvar_2, tmp10527)


_ = tmp10531

tmp10532 := MakeNative(func(__e *ControlFlow) {
V2227 := __e.Get(1)
_ = V2227
__e.Return(MakeString("|prolog vector|"))
return
}, 1)

tmp10533 := Call(__e, ns2_1set, symshen_4print_1prolog_1vector, tmp10532)


_ = tmp10533

tmp10534 := MakeNative(func(__e *ControlFlow) {
V2246 := __e.Get(1)
_ = V2246
V2247 := __e.Get(2)
_ = V2247
V2248 := __e.Get(3)
_ = V2248
V2249 := __e.Get(4)
_ = V2249
V2250 := __e.Get(5)
_ = V2250
tmp10547 := PrimEqual(Nil, V2246)

if True == tmp10547 {
__e.Return(False)
return
} else {
tmp10545 := PrimIsPair(V2246)

if True == tmp10545 {
tmp10535 := MakeNative(func(__e *ControlFlow) {
W2251 := __e.Get(1)
_ = W2251
tmp10538 := PrimEqual(W2251, False)

if True == tmp10538 {
tmp10536 := PrimTail(V2246)

__e.TailApply(PrimFunc(symfork), tmp10536, V2247, V2248, V2249, V2250)
return


} else {
__e.Return(W2251)
return
}


}, 1)

tmp10539 := PrimHead(V2246)

tmp10540 := Call(__e, tmp10539, V2247)


tmp10541 := Call(__e, tmp10540, V2248)


tmp10542 := Call(__e, tmp10541, V2249)


tmp10543 := Call(__e, tmp10542, V2250)


__e.TailApply(tmp10535, tmp10543)
return


} else {
__e.Return(PrimSimpleError(MakeString("fork expects a list of literals\n")))
return
}


}


}, 5)

tmp10548 := Call(__e, ns2_1set, symfork, tmp10534)


_ = tmp10548

tmp10549 := MakeNative(func(__e *ControlFlow) {
V2252 := __e.Get(1)
_ = V2252
V2253 := __e.Get(2)
_ = V2253
V2254 := __e.Get(3)
_ = V2254
V2255 := __e.Get(4)
_ = V2255
V2256 := __e.Get(5)
_ = V2256
V2257 := __e.Get(6)
_ = V2257
V2258 := __e.Get(7)
_ = V2258
tmp10556 := Call(__e, PrimFunc(symshen_4unlocked_2), V2256)


if True == tmp10556 {
tmp10550 := MakeNative(func(__e *ControlFlow) {
W2259 := __e.Get(1)
_ = W2259
tmp10551 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp10551

tmp10552 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4findall_1h), V2252, V2253, V2254, W2259, V2255, V2256, V2257, V2258)
return
}, 0)

tmp10553 := Call(__e, PrimFunc(symis), W2259, Nil, V2255, V2256, V2257, tmp10552)


__e.TailApply(PrimFunc(symshen_4gc), V2255, tmp10553)
return


}, 1)

tmp10554 := Call(__e, PrimFunc(symshen_4newpv), V2255)


__e.TailApply(tmp10550, tmp10554)
return


} else {
__e.Return(False)
return
}


}, 7)

tmp10557 := Call(__e, ns2_1set, symfindall, tmp10549)


_ = tmp10557

tmp10558 := MakeNative(func(__e *ControlFlow) {
V2260 := __e.Get(1)
_ = V2260
V2261 := __e.Get(2)
_ = V2261
V2262 := __e.Get(3)
_ = V2262
V2263 := __e.Get(4)
_ = V2263
V2264 := __e.Get(5)
_ = V2264
V2265 := __e.Get(6)
_ = V2265
V2266 := __e.Get(7)
_ = V2266
V2267 := __e.Get(8)
_ = V2267
tmp10559 := MakeNative(func(__e *ControlFlow) {
W2268 := __e.Get(1)
_ = W2268
tmp10564 := PrimEqual(W2268, False)

if True == tmp10564 {
tmp10562 := Call(__e, PrimFunc(symshen_4unlocked_2), V2265)


if True == tmp10562 {
tmp10560 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp10560

__e.TailApply(PrimFunc(symis_b), V2262, V2263, V2264, V2265, V2266, V2267)
return


} else {
__e.Return(False)
return
}


} else {
__e.Return(W2268)
return
}


}, 1)

tmp10569 := Call(__e, PrimFunc(symshen_4unlocked_2), V2265)


var ifres10565 Obj

if True == tmp10569 {
tmp10566 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp10566

tmp10567 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4overbind), V2260, V2263, V2264, V2265, V2266, V2267)
return
}, 0)

tmp10568 := Call(__e, PrimFunc(symcall), V2261, V2264, V2265, V2266, tmp10567)


ifres10565 = tmp10568


} else {
ifres10565 = False


}

__e.TailApply(tmp10559, ifres10565)
return


}, 8)

tmp10570 := Call(__e, ns2_1set, symshen_4findall_1h, tmp10558)


_ = tmp10570

tmp10571 := MakeNative(func(__e *ControlFlow) {
V2275 := __e.Get(1)
_ = V2275
V2276 := __e.Get(2)
_ = V2276
V2277 := __e.Get(3)
_ = V2277
V2278 := __e.Get(4)
_ = V2278
V2279 := __e.Get(5)
_ = V2279
V2280 := __e.Get(6)
_ = V2280
tmp10572 := Call(__e, PrimFunc(symshen_4deref), V2275, V2277)


tmp10573 := Call(__e, PrimFunc(symshen_4lazyderef), V2276, V2277)


tmp10574 := PrimCons(tmp10572, tmp10573)

tmp10575 := Call(__e, PrimFunc(symshen_4bindv), V2276, tmp10574, V2277)


_ = tmp10575

__e.Return(False)
return


}, 6)

tmp10576 := Call(__e, ns2_1set, symshen_4overbind, tmp10571)


_ = tmp10576

tmp10577 := MakeNative(func(__e *ControlFlow) {
V2283 := __e.Get(1)
_ = V2283
tmp10581 := PrimEqual(sym_7, V2283)

if True == tmp10581 {
__e.Return(PrimSet(symshen_4_doccurs_d, True))
return
} else {
tmp10579 := PrimEqual(sym_1, V2283)

if True == tmp10579 {
__e.Return(PrimSet(symshen_4_doccurs_d, False))
return
} else {
__e.Return(PrimSimpleError(MakeString("occurs-check expects a + or a -.\n")))
return
}


}


}, 1)

__e.TailApply(ns2_1set, symoccurs_1check, tmp10577)
return




}, 0)

