package main

import . "github.com/pyrex41/shen-go/kl"

var PrologMain = MakeNative0(func(__e *ControlFlow) Obj {
tmp9534 := MakeNative1(func(__e *ControlFlow, V1599 Obj) Obj {
return TailCall2(__e, SymFn(symshen_4assert_d), V1599, symshen_4top)
panic("unreachable")
})

tmp9535 := Call2(__e, ns2_1set, symasserta, tmp9534)


_ = tmp9535

tmp9536 := MakeNative1(func(__e *ControlFlow, V1600 Obj) Obj {
return TailCall2(__e, SymFn(symshen_4assert_d), V1600, symshen_4bottom)
panic("unreachable")
})

tmp9537 := Call2(__e, ns2_1set, symassertz, tmp9536)


_ = tmp9537

tmp9538 := MakeNative2(func(__e *ControlFlow, V1601, V1602 Obj) Obj {
tmp9563 := GConsP(__e, V1601)

var ifres9556 Obj

if True == tmp9563 {
tmp9561 := GTl(__e, V1601)

tmp9562 := GConsP(__e, tmp9561)

var ifres9557 Obj

if True == tmp9562 {
tmp9558 := GTl(__e, V1601)

tmp9559 := GHd(__e, tmp9558)

tmp9560 := GEq(__e, sym_5_1_1, tmp9559)

ifres9557 = tmp9560


} else {
ifres9557 = False


}

ifres9556 = ifres9557


} else {
ifres9556 = False


}

if True == ifres9556 {
tmp9539 := GHd(__e, V1601)

tmp9540 := Call1(__e, SymFn(symshen_4predicate), tmp9539)


let__9316 := tmp9540
_ = let__9316

tmp9541 := GHd(__e, V1601)

tmp9542 := Call1(__e, SymFn(symshen_4terms), tmp9541)


let__9317 := tmp9542
_ = let__9317

tmp9543 := Call1(__e, SymFn(symlength), let__9317)


let__9318 := tmp9543
_ = let__9318

tmp9544 := Call1(__e, SymFn(symshen_4parameters), let__9318)


let__9319 := tmp9544
_ = let__9319

tmp9545 := Call1(__e, SymFn(symarity), let__9316)


let__9320 := tmp9545
_ = let__9320

tmp9551 := GEq(__e, let__9320, MakeInteger(-1))

var ifres9546 Obj

if True == tmp9551 {
tmp9547 := Call2(__e, SymFn(symshen_4create_1skeleton), let__9316, let__9319)


tmp9548 := Call1(__e, SymFn(symeval), tmp9547)


_ = tmp9548

tmp9549 := GValue(__e, sym_dproperty_1vector_d)

tmp9550 := Call4(__e, SymFn(symput), let__9316, symshen_4dynamic, Nil, tmp9549)


ifres9546 = tmp9550


} else {
ifres9546 = symshen_4skip


}

let__9321 := ifres9546
_ = let__9321

tmp9552 := GTl(__e, V1601)

tmp9553 := GTl(__e, tmp9552)

tmp9554 := Call5(__e, SymFn(symshen_4insert_1info), let__9316, let__9317, tmp9553, V1601, V1602)


let__9322 := tmp9554
_ = let__9322

return let__9316


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4assert_d)
}


panic("unreachable")
})

tmp9564 := Call2(__e, ns2_1set, symshen_4assert_d, tmp9538)


_ = tmp9564

tmp9565 := MakeNative1(func(__e *ControlFlow, V1612 Obj) Obj {
tmp9567 := GConsP(__e, V1612)

if True == tmp9567 {
return GHd(__e, V1612)
} else {
return V1612
}


panic("unreachable")
})

tmp9568 := Call2(__e, ns2_1set, symshen_4predicate, tmp9565)


_ = tmp9568

tmp9569 := MakeNative1(func(__e *ControlFlow, V1617 Obj) Obj {
tmp9571 := GConsP(__e, V1617)

if True == tmp9571 {
return GTl(__e, V1617)
} else {
return Nil
}


panic("unreachable")
})

tmp9572 := Call2(__e, ns2_1set, symshen_4terms, tmp9569)


_ = tmp9572

tmp9573 := MakeNative2(func(__e *ControlFlow, V1618, V1619 Obj) Obj {
tmp9574 := Call2(__e, SymFn(symshen_4dynamic_1default), V1618, V1619)


tmp9575 := GCons(__e, V1618, tmp9574)

return GCons(__e, symdefprolog, tmp9575)


panic("unreachable")
})

tmp9576 := Call2(__e, ns2_1set, symshen_4create_1skeleton, tmp9573)


_ = tmp9576

tmp9577 := MakeNative2(func(__e *ControlFlow, V1620, V1621 Obj) Obj {
tmp9578 := Call1(__e, SymFn(symshen_4cons_1form), V1621)


tmp9579 := GCons(__e, symshen_4dynamic, Nil)

tmp9580 := GCons(__e, V1620, tmp9579)

tmp9581 := GCons(__e, symget, tmp9580)

tmp9582 := GCons(__e, tmp9581, Nil)

tmp9583 := GCons(__e, tmp9578, tmp9582)

tmp9584 := GCons(__e, symshen_4call_1dynamic, tmp9583)

tmp9585 := PrimGuard1(__e, symintern, PrimIntern, MakeString(";"))

tmp9586 := GCons(__e, tmp9585, Nil)

tmp9587 := GCons(__e, tmp9584, tmp9586)

tmp9588 := GCons(__e, sym_5_1_1, tmp9587)

return TailCall2(__e, SymFn(symappend), V1621, tmp9588)


panic("unreachable")
})

tmp9589 := Call2(__e, ns2_1set, symshen_4dynamic_1default, tmp9577)


_ = tmp9589

tmp9590 := MakeNative5(func(__e *ControlFlow, V1622, V1623, V1624, V1625, V1626 Obj) Obj {
tmp9591 := Call1(__e, SymFn(symgensym), symg)


let__9323 := tmp9591
_ = let__9323

tmp9592 := GCons(__e, let__9323, Nil)

tmp9593 := GCons(__e, symdefprolog, tmp9592)

tmp9594 := GCons(__e, sym_5_1_1, V1624)

tmp9595 := Call2(__e, SymFn(symappend), V1623, tmp9594)


tmp9596 := Call2(__e, SymFn(symappend), tmp9593, tmp9595)


tmp9597 := Call1(__e, SymFn(symeval), tmp9596)


let__9324 := tmp9597
_ = let__9324

tmp9598 := Call1(__e, SymFn(symfn), let__9323)


tmp9599 := GCons(__e, let__9323, V1625)

tmp9600 := GCons(__e, tmp9598, tmp9599)

let__9325 := tmp9600
_ = let__9325

tmp9601 := GValue(__e, sym_dproperty_1vector_d)

tmp9602 := Call3(__e, SymFn(symget), V1622, symshen_4dynamic, tmp9601)


let__9326 := tmp9602
_ = let__9326

tmp9607 := GEq(__e, V1626, symshen_4top)

var ifres9603 Obj

if True == tmp9607 {
tmp9604 := GCons(__e, let__9325, let__9326)

ifres9603 = tmp9604


} else {
tmp9605 := GCons(__e, let__9325, Nil)

tmp9606 := Call2(__e, SymFn(symappend), let__9326, tmp9605)


ifres9603 = tmp9606


}

let__9327 := ifres9603
_ = let__9327

tmp9608 := GValue(__e, sym_dproperty_1vector_d)

return TailCall4(__e, SymFn(symput), V1622, symshen_4dynamic, let__9327, tmp9608)


panic("unreachable")
})

tmp9609 := Call2(__e, ns2_1set, symshen_4insert_1info, tmp9590)


_ = tmp9609

tmp9610 := MakeNative0(func(__e *ControlFlow) Obj {
tmp9611 := GValue(__e, symshen_4_dnames_d)

let__9328 := tmp9611
_ = let__9328

tmp9617 := Call1(__e, SymFn(symempty_2), let__9328)


var ifres9612 Obj

if True == tmp9617 {
tmp9613 := Call1(__e, SymFn(symgensym), symg)


ifres9612 = tmp9613


} else {
tmp9614 := GTl(__e, let__9328)

tmp9615 := PrimGuard2(__e, symset, PrimSet, symshen_4_dnames_d, tmp9614)

_ = tmp9615

tmp9616 := GHd(__e, let__9328)

ifres9612 = tmp9616


}

let__9329 := ifres9612
_ = let__9329

return let__9329


panic("unreachable")
})

tmp9618 := Call2(__e, ns2_1set, symshen_4newname, tmp9610)


_ = tmp9618

tmp9619 := (func() Obj {
var __self Obj
__self = MakeNative6(func(__e *ControlFlow, __self1, __self2, __self3, __self4, __self5, __self6 Obj) Obj {
__selftop:
V1634 := __self1
_ = V1634
V1635 := __self2
_ = V1635
V1636 := __self3
_ = V1636
V1637 := __self4
_ = V1637
V1638 := __self5
_ = V1638
V1639 := __self6
_ = V1639
tmp9631 := Call1(__e, SymFn(symshen_4unlocked_2), V1637)


var ifres9620 Obj

if True == tmp9631 {
tmp9621 := Call2(__e, SymFn(symshen_4lazyderef), V1635, V1636)


let__9331 := tmp9621
_ = let__9331

tmp9630 := GConsP(__e, let__9331)

var ifres9622 Obj

if True == tmp9630 {
tmp9623 := GHd(__e, let__9331)

tmp9624 := Call2(__e, SymFn(symshen_4lazyderef), tmp9623, V1636)


let__9332 := tmp9624
_ = let__9332

tmp9629 := GConsP(__e, let__9332)

var ifres9625 Obj

if True == tmp9629 {
tmp9626 := GHd(__e, let__9332)

let__9333 := tmp9626
_ = let__9333

tmp9627 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp9627

tmp9628 := Call6(__e, SymFn(symshen_4callrec), let__9333, V1634, V1636, V1637, V1638, V1639)


ifres9625 = tmp9628


} else {
ifres9625 = False


}

ifres9622 = ifres9625


} else {
ifres9622 = False


}

ifres9620 = ifres9622


} else {
ifres9620 = False


}

let__9330 := ifres9620
_ = let__9330

tmp9640 := GEq(__e, let__9330, False)

if True == tmp9640 {
tmp9638 := Call1(__e, SymFn(symshen_4unlocked_2), V1637)


if True == tmp9638 {
tmp9632 := Call2(__e, SymFn(symshen_4lazyderef), V1635, V1636)


let__9334 := tmp9632
_ = let__9334

tmp9636 := GConsP(__e, let__9334)

if True == tmp9636 {
tmp9633 := GTl(__e, let__9334)

let__9335 := tmp9633
_ = let__9335

tmp9634 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp9634

if SymFn(symshen_4call_1dynamic) == __self {
__self1, __self2, __self3, __self4, __self5, __self6 = V1634, let__9335, V1636, V1637, V1638, V1639
__e.Tick()
goto __selftop
}
return TailCall6(__e, SymFn(symshen_4call_1dynamic), V1634, let__9335, V1636, V1637, V1638, V1639)


} else {
return False
}


} else {
return False
}


} else {
return let__9330
}


panic("unreachable")
})
return __self
})()

tmp9641 := Call2(__e, ns2_1set, symshen_4call_1dynamic, tmp9619)


_ = tmp9641

tmp9642 := (func() Obj {
var __self Obj
__self = MakeNative6(func(__e *ControlFlow, __self1, __self2, __self3, __self4, __self5, __self6 Obj) Obj {
__selftop:
V1646 := __self1
_ = V1646
V1647 := __self2
_ = V1647
V1648 := __self3
_ = V1648
V1649 := __self4
_ = V1649
V1650 := __self5
_ = V1650
V1651 := __self6
_ = V1651
tmp9652 := GEq(__e, Nil, V1647)

if True == tmp9652 {
tmp9643 := Call1(__e, V1646, V1648)


tmp9644 := Call1(__e, tmp9643, V1649)


tmp9645 := Call1(__e, tmp9644, V1650)


return TailCall1(__e, tmp9645, V1651)


} else {
tmp9650 := GConsP(__e, V1647)

if True == tmp9650 {
tmp9646 := GHd(__e, V1647)

tmp9647 := Call1(__e, V1646, tmp9646)


tmp9648 := GTl(__e, V1647)

if SymFn(symshen_4callrec) == __self {
__self1, __self2, __self3, __self4, __self5, __self6 = tmp9647, tmp9648, V1648, V1649, V1650, V1651
__e.Tick()
goto __selftop
}
return TailCall6(__e, SymFn(symshen_4callrec), tmp9647, tmp9648, V1648, V1649, V1650, V1651)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4callrec)
}


}


panic("unreachable")
})
return __self
})()

tmp9653 := Call2(__e, ns2_1set, symshen_4callrec, tmp9642)


_ = tmp9653

tmp9654 := MakeNative1(func(__e *ControlFlow, V1652 Obj) Obj {
tmp9669 := GConsP(__e, V1652)

var ifres9662 Obj

if True == tmp9669 {
tmp9667 := GTl(__e, V1652)

tmp9668 := GConsP(__e, tmp9667)

var ifres9663 Obj

if True == tmp9668 {
tmp9664 := GTl(__e, V1652)

tmp9665 := GHd(__e, tmp9664)

tmp9666 := GEq(__e, sym_5_1_1, tmp9665)

ifres9663 = tmp9666


} else {
ifres9663 = False


}

ifres9662 = ifres9663


} else {
ifres9662 = False


}

if True == ifres9662 {
tmp9655 := GHd(__e, V1652)

tmp9656 := Call1(__e, SymFn(symshen_4predicate), tmp9655)


let__9336 := tmp9656
_ = let__9336

tmp9657 := GValue(__e, sym_dproperty_1vector_d)

tmp9658 := Call3(__e, SymFn(symget), let__9336, symshen_4dynamic, tmp9657)


let__9337 := tmp9658
_ = let__9337

tmp9659 := Call2(__e, SymFn(symshen_4retract_1clause), V1652, let__9337)


tmp9660 := GValue(__e, sym_dproperty_1vector_d)

return TailCall4(__e, SymFn(symput), let__9336, symshen_4dynamic, tmp9659, tmp9660)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symretract)
}


panic("unreachable")
})

tmp9670 := Call2(__e, ns2_1set, symretract, tmp9654)


_ = tmp9670

tmp9671 := MakeNative2(func(__e *ControlFlow, V1660, V1661 Obj) Obj {
tmp9698 := GEq(__e, Nil, V1661)

if True == tmp9698 {
return Nil
} else {
tmp9696 := GConsP(__e, V1661)

var ifres9684 Obj

if True == tmp9696 {
tmp9694 := GHd(__e, V1661)

tmp9695 := GConsP(__e, tmp9694)

var ifres9685 Obj

if True == tmp9695 {
tmp9691 := GHd(__e, V1661)

tmp9692 := GTl(__e, tmp9691)

tmp9693 := GConsP(__e, tmp9692)

var ifres9686 Obj

if True == tmp9693 {
tmp9687 := GHd(__e, V1661)

tmp9688 := GTl(__e, tmp9687)

tmp9689 := GTl(__e, tmp9688)

tmp9690 := GEq(__e, V1660, tmp9689)

ifres9686 = tmp9690


} else {
ifres9686 = False


}

ifres9685 = ifres9686


} else {
ifres9685 = False


}

ifres9684 = ifres9685


} else {
ifres9684 = False


}

if True == ifres9684 {
tmp9672 := GHd(__e, V1661)

tmp9673 := GTl(__e, tmp9672)

tmp9674 := GHd(__e, tmp9673)

tmp9675 := GValue(__e, symshen_4_dnames_d)

tmp9676 := GCons(__e, tmp9674, tmp9675)

tmp9677 := PrimGuard2(__e, symset, PrimSet, symshen_4_dnames_d, tmp9676)

_ = tmp9677

return GTl(__e, V1661)


} else {
tmp9682 := GConsP(__e, V1661)

if True == tmp9682 {
tmp9678 := GHd(__e, V1661)

tmp9679 := GTl(__e, V1661)

tmp9680 := Call2(__e, SymFn(symshen_4retract_1clause), V1660, tmp9679)


return GCons(__e, tmp9678, tmp9680)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4retract_1clause)
}


}


}


panic("unreachable")
})

tmp9699 := Call2(__e, ns2_1set, symshen_4retract_1clause, tmp9671)


_ = tmp9699

tmp9700 := MakeNative2(func(__e *ControlFlow, V1662, V1663 Obj) Obj {
tmp9701 := MakeNative1(func(__e *ControlFlow, Z1664 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4_5defprolog_6), Z1664)
panic("unreachable")
})

tmp9702 := GCons(__e, V1662, V1663)

return TailCall2(__e, SymFn(symcompile), tmp9701, tmp9702)


panic("unreachable")
})

tmp9703 := Call2(__e, ns2_1set, symshen_4compile_1prolog, tmp9700)


_ = tmp9703

tmp9704 := MakeNative1(func(__e *ControlFlow, V1665 Obj) Obj {
tmp9720 := GConsP(__e, V1665)

var ifres9705 Obj

if True == tmp9720 {
tmp9706 := Call1(__e, SymFn(symhead), V1665)


let__9339 := tmp9706
_ = let__9339

tmp9707 := Call1(__e, SymFn(symtail), V1665)


let__9340 := tmp9707
_ = let__9340

tmp9708 := Call1(__e, SymFn(symshen_4_5clauses_6), let__9340)


let__9341 := tmp9708
_ = let__9341

tmp9718 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9341)


var ifres9709 Obj

if True == tmp9718 {
tmp9710 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres9709 = tmp9710


} else {
tmp9711 := Call1(__e, SymFn(symshen_4_5_1out), let__9341)


let__9342 := tmp9711
_ = let__9342

tmp9712 := Call1(__e, SymFn(symshen_4in_1_6), let__9341)


let__9343 := tmp9712
_ = let__9343

tmp9713 := Call2(__e, SymFn(symshen_4prolog_1arity_1check), let__9339, let__9342)


let__9344 := tmp9713
_ = let__9344

tmp9714 := MakeNative1(func(__e *ControlFlow, Z1674 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4linearise_1clause), Z1674)
panic("unreachable")
})

tmp9715 := Call2(__e, SymFn(symmap), tmp9714, let__9342)


let__9345 := tmp9715
_ = let__9345

tmp9716 := Call2(__e, SymFn(symshen_4horn_1clause_1procedure), let__9339, let__9345)


tmp9717 := Call2(__e, SymFn(symshen_4comb), let__9343, tmp9716)


ifres9709 = tmp9717


}

ifres9705 = ifres9709


} else {
tmp9719 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres9705 = tmp9719


}

let__9338 := ifres9705
_ = let__9338

tmp9722 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9338)


if True == tmp9722 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__9338
}


panic("unreachable")
})

tmp9723 := Call2(__e, ns2_1set, symshen_4_5defprolog_6, tmp9704)


_ = tmp9723

tmp9724 := MakeNative2(func(__e *ControlFlow, V1677, V1678 Obj) Obj {
tmp9761 := GConsP(__e, V1678)

var ifres9746 Obj

if True == tmp9761 {
tmp9759 := GHd(__e, V1678)

tmp9760 := GConsP(__e, tmp9759)

var ifres9747 Obj

if True == tmp9760 {
tmp9756 := GHd(__e, V1678)

tmp9757 := GTl(__e, tmp9756)

tmp9758 := GConsP(__e, tmp9757)

var ifres9748 Obj

if True == tmp9758 {
tmp9752 := GHd(__e, V1678)

tmp9753 := GTl(__e, tmp9752)

tmp9754 := GTl(__e, tmp9753)

tmp9755 := GEq(__e, Nil, tmp9754)

var ifres9749 Obj

if True == tmp9755 {
tmp9750 := GTl(__e, V1678)

tmp9751 := GEq(__e, Nil, tmp9750)

ifres9749 = tmp9751


} else {
ifres9749 = False


}

ifres9748 = ifres9749


} else {
ifres9748 = False


}

ifres9747 = ifres9748


} else {
ifres9747 = False


}

ifres9746 = ifres9747


} else {
ifres9746 = False


}

if True == ifres9746 {
tmp9725 := GHd(__e, V1678)

tmp9726 := GHd(__e, tmp9725)

return TailCall1(__e, SymFn(symlength), tmp9726)


} else {
tmp9744 := GConsP(__e, V1678)

var ifres9732 Obj

if True == tmp9744 {
tmp9742 := GHd(__e, V1678)

tmp9743 := GConsP(__e, tmp9742)

var ifres9733 Obj

if True == tmp9743 {
tmp9739 := GHd(__e, V1678)

tmp9740 := GTl(__e, tmp9739)

tmp9741 := GConsP(__e, tmp9740)

var ifres9734 Obj

if True == tmp9741 {
tmp9735 := GHd(__e, V1678)

tmp9736 := GTl(__e, tmp9735)

tmp9737 := GTl(__e, tmp9736)

tmp9738 := GEq(__e, Nil, tmp9737)

ifres9734 = tmp9738


} else {
ifres9734 = False


}

ifres9733 = ifres9734


} else {
ifres9733 = False


}

ifres9732 = ifres9733


} else {
ifres9732 = False


}

if True == ifres9732 {
tmp9727 := GHd(__e, V1678)

tmp9728 := GHd(__e, tmp9727)

tmp9729 := Call1(__e, SymFn(symlength), tmp9728)


tmp9730 := GTl(__e, V1678)

return TailCall3(__e, SymFn(symshen_4pac_1h), V1677, tmp9729, tmp9730)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4prolog_1arity_1check)
}


}


panic("unreachable")
})

tmp9762 := Call2(__e, ns2_1set, symshen_4prolog_1arity_1check, tmp9724)


_ = tmp9762

tmp9763 := (func() Obj {
var __self Obj
__self = MakeNative3(func(__e *ControlFlow, __self1, __self2, __self3 Obj) Obj {
__selftop:
V1683 := __self1
_ = V1683
V1684 := __self2
_ = V1684
V1685 := __self3
_ = V1685
tmp9778 := GEq(__e, Nil, V1685)

if True == tmp9778 {
return V1684
} else {
tmp9776 := GConsP(__e, V1685)

var ifres9773 Obj

if True == tmp9776 {
tmp9774 := GHd(__e, V1685)

tmp9775 := GConsP(__e, tmp9774)

ifres9773 = tmp9775


} else {
ifres9773 = False


}

if True == ifres9773 {
tmp9768 := GHd(__e, V1685)

tmp9769 := GHd(__e, tmp9768)

tmp9770 := Call1(__e, SymFn(symlength), tmp9769)


tmp9771 := GEq(__e, V1684, tmp9770)

if True == tmp9771 {
tmp9764 := GTl(__e, V1685)

if SymFn(symshen_4pac_1h) == __self {
__self1, __self2, __self3 = V1683, V1684, tmp9764
__e.Tick()
goto __selftop
}
return TailCall3(__e, SymFn(symshen_4pac_1h), V1683, V1684, tmp9764)


} else {
tmp9765 := Call3(__e, SymFn(symshen_4app), V1683, MakeString("\n"), symshen_4a)


return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("arity error in prolog procedure "), tmp9765))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("arity error in prolog procedure "), tmp9765)
return Call1(__e, SymFn(symsimple_1error), __typedArg0)
})()


}


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4pac_1h)
}


}


panic("unreachable")
})
return __self
})()

tmp9779 := Call2(__e, ns2_1set, symshen_4pac_1h, tmp9763)


_ = tmp9779

tmp9780 := MakeNative1(func(__e *ControlFlow, V1686 Obj) Obj {
tmp9781 := Call1(__e, SymFn(symshen_4_5clause_6), V1686)


let__9347 := tmp9781
_ = let__9347

tmp9794 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9347)


var ifres9782 Obj

if True == tmp9794 {
tmp9783 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres9782 = tmp9783


} else {
tmp9784 := Call1(__e, SymFn(symshen_4_5_1out), let__9347)


let__9348 := tmp9784
_ = let__9348

tmp9785 := Call1(__e, SymFn(symshen_4in_1_6), let__9347)


let__9349 := tmp9785
_ = let__9349

tmp9786 := Call1(__e, SymFn(symshen_4_5clauses_6), let__9349)


let__9350 := tmp9786
_ = let__9350

tmp9793 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9350)


var ifres9787 Obj

if True == tmp9793 {
tmp9788 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres9787 = tmp9788


} else {
tmp9789 := Call1(__e, SymFn(symshen_4_5_1out), let__9350)


let__9351 := tmp9789
_ = let__9351

tmp9790 := Call1(__e, SymFn(symshen_4in_1_6), let__9350)


let__9352 := tmp9790
_ = let__9352

tmp9791 := GCons(__e, let__9348, let__9351)

tmp9792 := Call2(__e, SymFn(symshen_4comb), let__9352, tmp9791)


ifres9787 = tmp9792


}

ifres9782 = ifres9787


}

let__9346 := ifres9782
_ = let__9346

tmp9810 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9346)


if True == tmp9810 {
tmp9795 := Call1(__e, SymFn(sym_5_b_6), V1686)


let__9354 := tmp9795
_ = let__9354

tmp9806 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9354)


var ifres9796 Obj

if True == tmp9806 {
tmp9797 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres9796 = tmp9797


} else {
tmp9798 := Call1(__e, SymFn(symshen_4_5_1out), let__9354)


let__9355 := tmp9798
_ = let__9355

tmp9799 := Call1(__e, SymFn(symshen_4in_1_6), let__9354)


let__9356 := tmp9799
_ = let__9356

tmp9804 := Call1(__e, SymFn(symempty_2), let__9355)


var ifres9800 Obj

if True == tmp9804 {
ifres9800 = Nil


} else {
tmp9801 := Call3(__e, SymFn(symshen_4app), let__9355, MakeString("\n ..."), symshen_4r)


tmp9803 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("Prolog syntax error here:\n "), tmp9801))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("Prolog syntax error here:\n "), tmp9801)
return Call1(__e, SymFn(symsimple_1error), __typedArg0)
})()

ifres9800 = tmp9803


}

tmp9805 := Call2(__e, SymFn(symshen_4comb), let__9356, ifres9800)


ifres9796 = tmp9805


}

let__9353 := ifres9796
_ = let__9353

tmp9808 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9353)


if True == tmp9808 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__9353
}


} else {
return let__9346
}


panic("unreachable")
})

tmp9811 := Call2(__e, ns2_1set, symshen_4_5clauses_6, tmp9780)


_ = tmp9811

tmp9812 := MakeNative1(func(__e *ControlFlow, V1698 Obj) Obj {
tmp9826 := GConsP(__e, V1698)

var ifres9819 Obj

if True == tmp9826 {
tmp9824 := GTl(__e, V1698)

tmp9825 := GConsP(__e, tmp9824)

var ifres9820 Obj

if True == tmp9825 {
tmp9821 := GTl(__e, V1698)

tmp9822 := GTl(__e, tmp9821)

tmp9823 := GEq(__e, Nil, tmp9822)

ifres9820 = tmp9823


} else {
ifres9820 = False


}

ifres9819 = ifres9820


} else {
ifres9819 = False


}

if True == ifres9819 {
tmp9813 := GHd(__e, V1698)

tmp9814 := GTl(__e, V1698)

tmp9815 := GHd(__e, tmp9814)

tmp9816 := Call2(__e, SymFn(sym_8p), tmp9813, tmp9815)


tmp9817 := Call1(__e, SymFn(symshen_4linearise), tmp9816)


return TailCall1(__e, SymFn(symshen_4lch), tmp9817)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4linearise_1clause)
}


panic("unreachable")
})

tmp9827 := Call2(__e, ns2_1set, symshen_4linearise_1clause, tmp9812)


_ = tmp9827

tmp9828 := MakeNative1(func(__e *ControlFlow, V1699 Obj) Obj {
tmp9834 := Call1(__e, SymFn(symtuple_2), V1699)


if True == tmp9834 {
tmp9829 := Call1(__e, SymFn(symfst), V1699)


tmp9830 := Call1(__e, SymFn(symsnd), V1699)


tmp9831 := Call1(__e, SymFn(symshen_4lchh), tmp9830)


tmp9832 := GCons(__e, tmp9831, Nil)

return GCons(__e, tmp9829, tmp9832)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4lch)
}


panic("unreachable")
})

tmp9835 := Call2(__e, ns2_1set, symshen_4lch, tmp9828)


_ = tmp9835

tmp9836 := MakeNative1(func(__e *ControlFlow, V1700 Obj) Obj {
tmp9890 := GConsP(__e, V1700)

var ifres9848 Obj

if True == tmp9890 {
tmp9888 := GHd(__e, V1700)

tmp9889 := GEq(__e, symwhere, tmp9888)

var ifres9849 Obj

if True == tmp9889 {
tmp9886 := GTl(__e, V1700)

tmp9887 := GConsP(__e, tmp9886)

var ifres9850 Obj

if True == tmp9887 {
tmp9883 := GTl(__e, V1700)

tmp9884 := GHd(__e, tmp9883)

tmp9885 := GConsP(__e, tmp9884)

var ifres9851 Obj

if True == tmp9885 {
tmp9879 := GTl(__e, V1700)

tmp9880 := GHd(__e, tmp9879)

tmp9881 := GHd(__e, tmp9880)

tmp9882 := GEq(__e, sym_a, tmp9881)

var ifres9852 Obj

if True == tmp9882 {
tmp9875 := GTl(__e, V1700)

tmp9876 := GHd(__e, tmp9875)

tmp9877 := GTl(__e, tmp9876)

tmp9878 := GConsP(__e, tmp9877)

var ifres9853 Obj

if True == tmp9878 {
tmp9870 := GTl(__e, V1700)

tmp9871 := GHd(__e, tmp9870)

tmp9872 := GTl(__e, tmp9871)

tmp9873 := GTl(__e, tmp9872)

tmp9874 := GConsP(__e, tmp9873)

var ifres9854 Obj

if True == tmp9874 {
tmp9864 := GTl(__e, V1700)

tmp9865 := GHd(__e, tmp9864)

tmp9866 := GTl(__e, tmp9865)

tmp9867 := GTl(__e, tmp9866)

tmp9868 := GTl(__e, tmp9867)

tmp9869 := GEq(__e, Nil, tmp9868)

var ifres9855 Obj

if True == tmp9869 {
tmp9861 := GTl(__e, V1700)

tmp9862 := GTl(__e, tmp9861)

tmp9863 := GConsP(__e, tmp9862)

var ifres9856 Obj

if True == tmp9863 {
tmp9857 := GTl(__e, V1700)

tmp9858 := GTl(__e, tmp9857)

tmp9859 := GTl(__e, tmp9858)

tmp9860 := GEq(__e, Nil, tmp9859)

ifres9856 = tmp9860


} else {
ifres9856 = False


}

ifres9855 = ifres9856


} else {
ifres9855 = False


}

ifres9854 = ifres9855


} else {
ifres9854 = False


}

ifres9853 = ifres9854


} else {
ifres9853 = False


}

ifres9852 = ifres9853


} else {
ifres9852 = False


}

ifres9851 = ifres9852


} else {
ifres9851 = False


}

ifres9850 = ifres9851


} else {
ifres9850 = False


}

ifres9849 = ifres9850


} else {
ifres9849 = False


}

ifres9848 = ifres9849


} else {
ifres9848 = False


}

if True == ifres9848 {
tmp9838 := GValue(__e, symshen_4_doccurs_d)

var ifres9837 Obj

if True == tmp9838 {
ifres9837 = symis_b


} else {
ifres9837 = symis


}

tmp9839 := GTl(__e, V1700)

tmp9840 := GHd(__e, tmp9839)

tmp9841 := GTl(__e, tmp9840)

tmp9842 := GCons(__e, ifres9837, tmp9841)

tmp9843 := GTl(__e, V1700)

tmp9844 := GTl(__e, tmp9843)

tmp9845 := GHd(__e, tmp9844)

tmp9846 := Call1(__e, SymFn(symshen_4lchh), tmp9845)


return GCons(__e, tmp9842, tmp9846)


} else {
return V1700
}


panic("unreachable")
})

tmp9891 := Call2(__e, ns2_1set, symshen_4lchh, tmp9836)


_ = tmp9891

tmp9892 := MakeNative1(func(__e *ControlFlow, V1701 Obj) Obj {
tmp9893 := Call1(__e, SymFn(symshen_4_5head_6), V1701)


let__9358 := tmp9893
_ = let__9358

tmp9916 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9358)


var ifres9894 Obj

if True == tmp9916 {
tmp9895 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres9894 = tmp9895


} else {
tmp9896 := Call1(__e, SymFn(symshen_4_5_1out), let__9358)


let__9359 := tmp9896
_ = let__9359

tmp9897 := Call1(__e, SymFn(symshen_4in_1_6), let__9358)


let__9360 := tmp9897
_ = let__9360

tmp9915 := Call2(__e, SymFn(symshen_4hds_a_2), let__9360, sym_5_1_1)


var ifres9898 Obj

if True == tmp9915 {
tmp9899 := Call1(__e, SymFn(symtail), let__9360)


let__9361 := tmp9899
_ = let__9361

tmp9900 := Call1(__e, SymFn(symshen_4_5body_6), let__9361)


let__9362 := tmp9900
_ = let__9362

tmp9913 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9362)


var ifres9901 Obj

if True == tmp9913 {
tmp9902 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres9901 = tmp9902


} else {
tmp9903 := Call1(__e, SymFn(symshen_4_5_1out), let__9362)


let__9363 := tmp9903
_ = let__9363

tmp9904 := Call1(__e, SymFn(symshen_4in_1_6), let__9362)


let__9364 := tmp9904
_ = let__9364

tmp9905 := Call1(__e, SymFn(symshen_4_5sc_6), let__9364)


let__9365 := tmp9905
_ = let__9365

tmp9912 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9365)


var ifres9906 Obj

if True == tmp9912 {
tmp9907 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres9906 = tmp9907


} else {
tmp9908 := Call1(__e, SymFn(symshen_4in_1_6), let__9365)


let__9366 := tmp9908
_ = let__9366

tmp9909 := GCons(__e, let__9363, Nil)

tmp9910 := GCons(__e, let__9359, tmp9909)

tmp9911 := Call2(__e, SymFn(symshen_4comb), let__9366, tmp9910)


ifres9906 = tmp9911


}

ifres9901 = ifres9906


}

ifres9898 = ifres9901


} else {
tmp9914 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres9898 = tmp9914


}

ifres9894 = ifres9898


}

let__9357 := ifres9894
_ = let__9357

tmp9918 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9357)


if True == tmp9918 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__9357
}


panic("unreachable")
})

tmp9919 := Call2(__e, ns2_1set, symshen_4_5clause_6, tmp9892)


_ = tmp9919

tmp9920 := MakeNative1(func(__e *ControlFlow, V1712 Obj) Obj {
tmp9921 := Call1(__e, SymFn(symshen_4_5hterm_6), V1712)


let__9368 := tmp9921
_ = let__9368

tmp9934 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9368)


var ifres9922 Obj

if True == tmp9934 {
tmp9923 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres9922 = tmp9923


} else {
tmp9924 := Call1(__e, SymFn(symshen_4_5_1out), let__9368)


let__9369 := tmp9924
_ = let__9369

tmp9925 := Call1(__e, SymFn(symshen_4in_1_6), let__9368)


let__9370 := tmp9925
_ = let__9370

tmp9926 := Call1(__e, SymFn(symshen_4_5head_6), let__9370)


let__9371 := tmp9926
_ = let__9371

tmp9933 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9371)


var ifres9927 Obj

if True == tmp9933 {
tmp9928 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres9927 = tmp9928


} else {
tmp9929 := Call1(__e, SymFn(symshen_4_5_1out), let__9371)


let__9372 := tmp9929
_ = let__9372

tmp9930 := Call1(__e, SymFn(symshen_4in_1_6), let__9371)


let__9373 := tmp9930
_ = let__9373

tmp9931 := GCons(__e, let__9369, let__9372)

tmp9932 := Call2(__e, SymFn(symshen_4comb), let__9373, tmp9931)


ifres9927 = tmp9932


}

ifres9922 = ifres9927


}

let__9367 := ifres9922
_ = let__9367

tmp9944 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9367)


if True == tmp9944 {
tmp9935 := Call1(__e, SymFn(sym_5e_6), V1712)


let__9375 := tmp9935
_ = let__9375

tmp9940 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9375)


var ifres9936 Obj

if True == tmp9940 {
tmp9937 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres9936 = tmp9937


} else {
tmp9938 := Call1(__e, SymFn(symshen_4in_1_6), let__9375)


let__9376 := tmp9938
_ = let__9376

tmp9939 := Call2(__e, SymFn(symshen_4comb), let__9376, Nil)


ifres9936 = tmp9939


}

let__9374 := ifres9936
_ = let__9374

tmp9942 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9374)


if True == tmp9942 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__9374
}


} else {
return let__9367
}


panic("unreachable")
})

tmp9945 := Call2(__e, ns2_1set, symshen_4_5head_6, tmp9920)


_ = tmp9945

tmp9946 := MakeNative1(func(__e *ControlFlow, V1723 Obj) Obj {
tmp9958 := GConsP(__e, V1723)

var ifres9947 Obj

if True == tmp9958 {
tmp9948 := Call1(__e, SymFn(symhead), V1723)


let__9378 := tmp9948
_ = let__9378

tmp9949 := Call1(__e, SymFn(symtail), V1723)


let__9379 := tmp9949
_ = let__9379

tmp9956 := Call1(__e, SymFn(symatom_2), let__9378)


var ifres9953 Obj

if True == tmp9956 {
tmp9954 := Call1(__e, SymFn(symshen_4prolog_1keyword_2), let__9378)


ifres9953 = GNot(__e, tmp9954)


} else {
ifres9953 = False


}

var ifres9950 Obj

if True == ifres9953 {
tmp9951 := Call2(__e, SymFn(symshen_4comb), let__9379, let__9378)


ifres9950 = tmp9951


} else {
tmp9952 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres9950 = tmp9952


}

ifres9947 = ifres9950


} else {
tmp9957 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres9947 = tmp9957


}

let__9377 := ifres9947
_ = let__9377

tmp10112 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9377)


if True == tmp10112 {
tmp9968 := GConsP(__e, V1723)

var ifres9959 Obj

if True == tmp9968 {
tmp9960 := Call1(__e, SymFn(symhead), V1723)


let__9381 := tmp9960
_ = let__9381

tmp9961 := Call1(__e, SymFn(symtail), V1723)


let__9382 := tmp9961
_ = let__9382

tmp9965 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":"))

tmp9966 := GEq(__e, let__9381, tmp9965)

var ifres9962 Obj

if True == tmp9966 {
tmp9963 := Call2(__e, SymFn(symshen_4comb), let__9382, let__9381)


ifres9962 = tmp9963


} else {
tmp9964 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres9962 = tmp9964


}

ifres9959 = ifres9962


} else {
tmp9967 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres9959 = tmp9967


}

let__9380 := ifres9959
_ = let__9380

tmp10110 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9380)


if True == tmp10110 {
tmp9998 := Call1(__e, SymFn(symshen_4ccons_2), V1723)


var ifres9969 Obj

if True == tmp9998 {
tmp9970 := Call1(__e, SymFn(symhead), V1723)


let__9384 := tmp9970
_ = let__9384

tmp9971 := Call1(__e, SymFn(symtail), V1723)


let__9385 := tmp9971
_ = let__9385

tmp9996 := Call2(__e, SymFn(symshen_4hds_a_2), let__9384, symcons)


var ifres9972 Obj

if True == tmp9996 {
tmp9973 := Call1(__e, SymFn(symtail), let__9384)


let__9386 := tmp9973
_ = let__9386

tmp9974 := Call1(__e, SymFn(symshen_4_5hterm1_6), let__9386)


let__9387 := tmp9974
_ = let__9387

tmp9994 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9387)


var ifres9975 Obj

if True == tmp9994 {
tmp9976 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres9975 = tmp9976


} else {
tmp9977 := Call1(__e, SymFn(symshen_4_5_1out), let__9387)


let__9388 := tmp9977
_ = let__9388

tmp9978 := Call1(__e, SymFn(symshen_4in_1_6), let__9387)


let__9389 := tmp9978
_ = let__9389

tmp9979 := Call1(__e, SymFn(symshen_4_5hterm2_6), let__9389)


let__9390 := tmp9979
_ = let__9390

tmp9993 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9390)


var ifres9980 Obj

if True == tmp9993 {
tmp9981 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres9980 = tmp9981


} else {
tmp9982 := Call1(__e, SymFn(symshen_4_5_1out), let__9390)


let__9391 := tmp9982
_ = let__9391

tmp9983 := Call1(__e, SymFn(symshen_4in_1_6), let__9390)


let__9392 := tmp9983
_ = let__9392

tmp9984 := Call1(__e, SymFn(sym_5end_6), let__9392)


let__9393 := tmp9984
_ = let__9393

tmp9992 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9393)


var ifres9985 Obj

if True == tmp9992 {
tmp9986 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres9985 = tmp9986


} else {
tmp9987 := Call1(__e, SymFn(symshen_4in_1_6), let__9393)


let__9394 := tmp9987
_ = let__9394

tmp9988 := GCons(__e, let__9391, Nil)

tmp9989 := GCons(__e, let__9388, tmp9988)

tmp9990 := GCons(__e, symcons, tmp9989)

tmp9991 := Call2(__e, SymFn(symshen_4comb), let__9385, tmp9990)


ifres9985 = tmp9991


}

ifres9980 = ifres9985


}

ifres9975 = ifres9980


}

ifres9972 = ifres9975


} else {
tmp9995 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres9972 = tmp9995


}

ifres9969 = ifres9972


} else {
tmp9997 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres9969 = tmp9997


}

let__9383 := ifres9969
_ = let__9383

tmp10108 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9383)


if True == tmp10108 {
tmp10021 := Call1(__e, SymFn(symshen_4ccons_2), V1723)


var ifres9999 Obj

if True == tmp10021 {
tmp10000 := Call1(__e, SymFn(symhead), V1723)


let__9396 := tmp10000
_ = let__9396

tmp10001 := Call1(__e, SymFn(symtail), V1723)


let__9397 := tmp10001
_ = let__9397

tmp10019 := Call2(__e, SymFn(symshen_4hds_a_2), let__9396, sym_7)


var ifres10002 Obj

if True == tmp10019 {
tmp10003 := Call1(__e, SymFn(symtail), let__9396)


let__9398 := tmp10003
_ = let__9398

tmp10004 := Call1(__e, SymFn(symshen_4_5hterm_6), let__9398)


let__9399 := tmp10004
_ = let__9399

tmp10017 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9399)


var ifres10005 Obj

if True == tmp10017 {
tmp10006 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10005 = tmp10006


} else {
tmp10007 := Call1(__e, SymFn(symshen_4_5_1out), let__9399)


let__9400 := tmp10007
_ = let__9400

tmp10008 := Call1(__e, SymFn(symshen_4in_1_6), let__9399)


let__9401 := tmp10008
_ = let__9401

tmp10009 := Call1(__e, SymFn(sym_5end_6), let__9401)


let__9402 := tmp10009
_ = let__9402

tmp10016 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9402)


var ifres10010 Obj

if True == tmp10016 {
tmp10011 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10010 = tmp10011


} else {
tmp10012 := Call1(__e, SymFn(symshen_4in_1_6), let__9402)


let__9403 := tmp10012
_ = let__9403

tmp10013 := GCons(__e, let__9400, Nil)

tmp10014 := GCons(__e, symshen_4_7m, tmp10013)

tmp10015 := Call2(__e, SymFn(symshen_4comb), let__9397, tmp10014)


ifres10010 = tmp10015


}

ifres10005 = ifres10010


}

ifres10002 = ifres10005


} else {
tmp10018 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10002 = tmp10018


}

ifres9999 = ifres10002


} else {
tmp10020 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres9999 = tmp10020


}

let__9395 := ifres9999
_ = let__9395

tmp10106 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9395)


if True == tmp10106 {
tmp10044 := Call1(__e, SymFn(symshen_4ccons_2), V1723)


var ifres10022 Obj

if True == tmp10044 {
tmp10023 := Call1(__e, SymFn(symhead), V1723)


let__9405 := tmp10023
_ = let__9405

tmp10024 := Call1(__e, SymFn(symtail), V1723)


let__9406 := tmp10024
_ = let__9406

tmp10042 := Call2(__e, SymFn(symshen_4hds_a_2), let__9405, sym_1)


var ifres10025 Obj

if True == tmp10042 {
tmp10026 := Call1(__e, SymFn(symtail), let__9405)


let__9407 := tmp10026
_ = let__9407

tmp10027 := Call1(__e, SymFn(symshen_4_5hterm_6), let__9407)


let__9408 := tmp10027
_ = let__9408

tmp10040 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9408)


var ifres10028 Obj

if True == tmp10040 {
tmp10029 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10028 = tmp10029


} else {
tmp10030 := Call1(__e, SymFn(symshen_4_5_1out), let__9408)


let__9409 := tmp10030
_ = let__9409

tmp10031 := Call1(__e, SymFn(symshen_4in_1_6), let__9408)


let__9410 := tmp10031
_ = let__9410

tmp10032 := Call1(__e, SymFn(sym_5end_6), let__9410)


let__9411 := tmp10032
_ = let__9411

tmp10039 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9411)


var ifres10033 Obj

if True == tmp10039 {
tmp10034 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10033 = tmp10034


} else {
tmp10035 := Call1(__e, SymFn(symshen_4in_1_6), let__9411)


let__9412 := tmp10035
_ = let__9412

tmp10036 := GCons(__e, let__9409, Nil)

tmp10037 := GCons(__e, symshen_4_1m, tmp10036)

tmp10038 := Call2(__e, SymFn(symshen_4comb), let__9406, tmp10037)


ifres10033 = tmp10038


}

ifres10028 = ifres10033


}

ifres10025 = ifres10028


} else {
tmp10041 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10025 = tmp10041


}

ifres10022 = ifres10025


} else {
tmp10043 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10022 = tmp10043


}

let__9404 := ifres10022
_ = let__9404

tmp10104 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9404)


if True == tmp10104 {
tmp10071 := Call1(__e, SymFn(symshen_4ccons_2), V1723)


var ifres10045 Obj

if True == tmp10071 {
tmp10046 := Call1(__e, SymFn(symhead), V1723)


let__9414 := tmp10046
_ = let__9414

tmp10047 := Call1(__e, SymFn(symtail), V1723)


let__9415 := tmp10047
_ = let__9415

tmp10069 := Call2(__e, SymFn(symshen_4hds_a_2), let__9414, symmode)


var ifres10048 Obj

if True == tmp10069 {
tmp10049 := Call1(__e, SymFn(symtail), let__9414)


let__9416 := tmp10049
_ = let__9416

tmp10050 := Call1(__e, SymFn(symshen_4_5hterm_6), let__9416)


let__9417 := tmp10050
_ = let__9417

tmp10067 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9417)


var ifres10051 Obj

if True == tmp10067 {
tmp10052 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10051 = tmp10052


} else {
tmp10053 := Call1(__e, SymFn(symshen_4_5_1out), let__9417)


let__9418 := tmp10053
_ = let__9418

tmp10054 := Call1(__e, SymFn(symshen_4in_1_6), let__9417)


let__9419 := tmp10054
_ = let__9419

tmp10066 := Call2(__e, SymFn(symshen_4hds_a_2), let__9419, sym_7)


var ifres10055 Obj

if True == tmp10066 {
tmp10056 := Call1(__e, SymFn(symtail), let__9419)


let__9420 := tmp10056
_ = let__9420

tmp10057 := Call1(__e, SymFn(sym_5end_6), let__9420)


let__9421 := tmp10057
_ = let__9421

tmp10064 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9421)


var ifres10058 Obj

if True == tmp10064 {
tmp10059 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10058 = tmp10059


} else {
tmp10060 := Call1(__e, SymFn(symshen_4in_1_6), let__9421)


let__9422 := tmp10060
_ = let__9422

tmp10061 := GCons(__e, let__9418, Nil)

tmp10062 := GCons(__e, symshen_4_7m, tmp10061)

tmp10063 := Call2(__e, SymFn(symshen_4comb), let__9415, tmp10062)


ifres10058 = tmp10063


}

ifres10055 = ifres10058


} else {
tmp10065 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10055 = tmp10065


}

ifres10051 = ifres10055


}

ifres10048 = ifres10051


} else {
tmp10068 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10048 = tmp10068


}

ifres10045 = ifres10048


} else {
tmp10070 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10045 = tmp10070


}

let__9413 := ifres10045
_ = let__9413

tmp10102 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9413)


if True == tmp10102 {
tmp10098 := Call1(__e, SymFn(symshen_4ccons_2), V1723)


var ifres10072 Obj

if True == tmp10098 {
tmp10073 := Call1(__e, SymFn(symhead), V1723)


let__9424 := tmp10073
_ = let__9424

tmp10074 := Call1(__e, SymFn(symtail), V1723)


let__9425 := tmp10074
_ = let__9425

tmp10096 := Call2(__e, SymFn(symshen_4hds_a_2), let__9424, symmode)


var ifres10075 Obj

if True == tmp10096 {
tmp10076 := Call1(__e, SymFn(symtail), let__9424)


let__9426 := tmp10076
_ = let__9426

tmp10077 := Call1(__e, SymFn(symshen_4_5hterm_6), let__9426)


let__9427 := tmp10077
_ = let__9427

tmp10094 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9427)


var ifres10078 Obj

if True == tmp10094 {
tmp10079 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10078 = tmp10079


} else {
tmp10080 := Call1(__e, SymFn(symshen_4_5_1out), let__9427)


let__9428 := tmp10080
_ = let__9428

tmp10081 := Call1(__e, SymFn(symshen_4in_1_6), let__9427)


let__9429 := tmp10081
_ = let__9429

tmp10093 := Call2(__e, SymFn(symshen_4hds_a_2), let__9429, sym_1)


var ifres10082 Obj

if True == tmp10093 {
tmp10083 := Call1(__e, SymFn(symtail), let__9429)


let__9430 := tmp10083
_ = let__9430

tmp10084 := Call1(__e, SymFn(sym_5end_6), let__9430)


let__9431 := tmp10084
_ = let__9431

tmp10091 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9431)


var ifres10085 Obj

if True == tmp10091 {
tmp10086 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10085 = tmp10086


} else {
tmp10087 := Call1(__e, SymFn(symshen_4in_1_6), let__9431)


let__9432 := tmp10087
_ = let__9432

tmp10088 := GCons(__e, let__9428, Nil)

tmp10089 := GCons(__e, symshen_4_1m, tmp10088)

tmp10090 := Call2(__e, SymFn(symshen_4comb), let__9425, tmp10089)


ifres10085 = tmp10090


}

ifres10082 = ifres10085


} else {
tmp10092 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10082 = tmp10092


}

ifres10078 = ifres10082


}

ifres10075 = ifres10078


} else {
tmp10095 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10075 = tmp10095


}

ifres10072 = ifres10075


} else {
tmp10097 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10072 = tmp10097


}

let__9423 := ifres10072
_ = let__9423

tmp10100 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9423)


if True == tmp10100 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__9423
}


} else {
return let__9413
}


} else {
return let__9404
}


} else {
return let__9395
}


} else {
return let__9383
}


} else {
return let__9380
}


} else {
return let__9377
}


panic("unreachable")
})

tmp10113 := Call2(__e, ns2_1set, symshen_4_5hterm_6, tmp9946)


_ = tmp10113

tmp10114 := MakeNative1(func(__e *ControlFlow, V1780 Obj) Obj {
tmp10115 := PrimGuard1(__e, symintern, PrimIntern, MakeString(";"))

tmp10116 := GCons(__e, sym_5_1_1, Nil)

tmp10117 := GCons(__e, tmp10115, tmp10116)

return TailCall2(__e, SymFn(symelement_2), V1780, tmp10117)


panic("unreachable")
})

tmp10118 := Call2(__e, ns2_1set, symshen_4prolog_1keyword_2, tmp10114)


_ = tmp10118

tmp10119 := MakeNative1(func(__e *ControlFlow, V1781 Obj) Obj {
tmp10127 := PrimGuard1(__e, symsymbol_2, PrimIsSymbol, V1781)

if True == tmp10127 {
return True
} else {
tmp10125 := PrimGuard1(__e, symstring_2, PrimIsString, V1781)

if True == tmp10125 {
return True
} else {
tmp10123 := Call1(__e, SymFn(symboolean_2), V1781)


if True == tmp10123 {
return True
} else {
tmp10121 := PrimGuard1(__e, symnumber_2, PrimIsNumber, V1781)

if True == tmp10121 {
return True
} else {
return TailCall1(__e, SymFn(symempty_2), V1781)
}


}


}


}


panic("unreachable")
})

tmp10128 := Call2(__e, ns2_1set, symatom_2, tmp10119)


_ = tmp10128

tmp10129 := MakeNative1(func(__e *ControlFlow, V1782 Obj) Obj {
tmp10130 := Call1(__e, SymFn(symshen_4_5hterm_6), V1782)


let__9434 := tmp10130
_ = let__9434

tmp10136 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9434)


var ifres10131 Obj

if True == tmp10136 {
tmp10132 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10131 = tmp10132


} else {
tmp10133 := Call1(__e, SymFn(symshen_4_5_1out), let__9434)


let__9435 := tmp10133
_ = let__9435

tmp10134 := Call1(__e, SymFn(symshen_4in_1_6), let__9434)


let__9436 := tmp10134
_ = let__9436

tmp10135 := Call2(__e, SymFn(symshen_4comb), let__9436, let__9435)


ifres10131 = tmp10135


}

let__9433 := ifres10131
_ = let__9433

tmp10138 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9433)


if True == tmp10138 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__9433
}


panic("unreachable")
})

tmp10139 := Call2(__e, ns2_1set, symshen_4_5hterm1_6, tmp10129)


_ = tmp10139

tmp10140 := MakeNative1(func(__e *ControlFlow, V1787 Obj) Obj {
tmp10141 := Call1(__e, SymFn(symshen_4_5hterm_6), V1787)


let__9438 := tmp10141
_ = let__9438

tmp10147 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9438)


var ifres10142 Obj

if True == tmp10147 {
tmp10143 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10142 = tmp10143


} else {
tmp10144 := Call1(__e, SymFn(symshen_4_5_1out), let__9438)


let__9439 := tmp10144
_ = let__9439

tmp10145 := Call1(__e, SymFn(symshen_4in_1_6), let__9438)


let__9440 := tmp10145
_ = let__9440

tmp10146 := Call2(__e, SymFn(symshen_4comb), let__9440, let__9439)


ifres10142 = tmp10146


}

let__9437 := ifres10142
_ = let__9437

tmp10149 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9437)


if True == tmp10149 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__9437
}


panic("unreachable")
})

tmp10150 := Call2(__e, ns2_1set, symshen_4_5hterm2_6, tmp10140)


_ = tmp10150

tmp10151 := MakeNative1(func(__e *ControlFlow, V1792 Obj) Obj {
tmp10152 := Call1(__e, SymFn(symshen_4_5literal_6), V1792)


let__9442 := tmp10152
_ = let__9442

tmp10165 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9442)


var ifres10153 Obj

if True == tmp10165 {
tmp10154 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10153 = tmp10154


} else {
tmp10155 := Call1(__e, SymFn(symshen_4_5_1out), let__9442)


let__9443 := tmp10155
_ = let__9443

tmp10156 := Call1(__e, SymFn(symshen_4in_1_6), let__9442)


let__9444 := tmp10156
_ = let__9444

tmp10157 := Call1(__e, SymFn(symshen_4_5body_6), let__9444)


let__9445 := tmp10157
_ = let__9445

tmp10164 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9445)


var ifres10158 Obj

if True == tmp10164 {
tmp10159 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10158 = tmp10159


} else {
tmp10160 := Call1(__e, SymFn(symshen_4_5_1out), let__9445)


let__9446 := tmp10160
_ = let__9446

tmp10161 := Call1(__e, SymFn(symshen_4in_1_6), let__9445)


let__9447 := tmp10161
_ = let__9447

tmp10162 := GCons(__e, let__9443, let__9446)

tmp10163 := Call2(__e, SymFn(symshen_4comb), let__9447, tmp10162)


ifres10158 = tmp10163


}

ifres10153 = ifres10158


}

let__9441 := ifres10153
_ = let__9441

tmp10175 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9441)


if True == tmp10175 {
tmp10166 := Call1(__e, SymFn(sym_5e_6), V1792)


let__9449 := tmp10166
_ = let__9449

tmp10171 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9449)


var ifres10167 Obj

if True == tmp10171 {
tmp10168 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10167 = tmp10168


} else {
tmp10169 := Call1(__e, SymFn(symshen_4in_1_6), let__9449)


let__9450 := tmp10169
_ = let__9450

tmp10170 := Call2(__e, SymFn(symshen_4comb), let__9450, Nil)


ifres10167 = tmp10170


}

let__9448 := ifres10167
_ = let__9448

tmp10173 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9448)


if True == tmp10173 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__9448
}


} else {
return let__9441
}


panic("unreachable")
})

tmp10176 := Call2(__e, ns2_1set, symshen_4_5body_6, tmp10151)


_ = tmp10176

tmp10177 := MakeNative1(func(__e *ControlFlow, V1803 Obj) Obj {
tmp10182 := Call2(__e, SymFn(symshen_4hds_a_2), V1803, sym_b)


var ifres10178 Obj

if True == tmp10182 {
tmp10179 := Call1(__e, SymFn(symtail), V1803)


let__9452 := tmp10179
_ = let__9452

tmp10180 := Call2(__e, SymFn(symshen_4comb), let__9452, sym_b)


ifres10178 = tmp10180


} else {
tmp10181 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10178 = tmp10181


}

let__9451 := ifres10178
_ = let__9451

tmp10203 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9451)


if True == tmp10203 {
tmp10199 := Call1(__e, SymFn(symshen_4ccons_2), V1803)


var ifres10183 Obj

if True == tmp10199 {
tmp10184 := Call1(__e, SymFn(symhead), V1803)


let__9454 := tmp10184
_ = let__9454

tmp10185 := Call1(__e, SymFn(symtail), V1803)


let__9455 := tmp10185
_ = let__9455

tmp10186 := Call1(__e, SymFn(symshen_4_5bterms_6), let__9454)


let__9456 := tmp10186
_ = let__9456

tmp10197 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9456)


var ifres10187 Obj

if True == tmp10197 {
tmp10188 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10187 = tmp10188


} else {
tmp10189 := Call1(__e, SymFn(symshen_4_5_1out), let__9456)


let__9457 := tmp10189
_ = let__9457

tmp10190 := Call1(__e, SymFn(symshen_4in_1_6), let__9456)


let__9458 := tmp10190
_ = let__9458

tmp10191 := Call1(__e, SymFn(sym_5end_6), let__9458)


let__9459 := tmp10191
_ = let__9459

tmp10196 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9459)


var ifres10192 Obj

if True == tmp10196 {
tmp10193 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10192 = tmp10193


} else {
tmp10194 := Call1(__e, SymFn(symshen_4in_1_6), let__9459)


let__9460 := tmp10194
_ = let__9460

tmp10195 := Call2(__e, SymFn(symshen_4comb), let__9455, let__9457)


ifres10192 = tmp10195


}

ifres10187 = ifres10192


}

ifres10183 = ifres10187


} else {
tmp10198 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10183 = tmp10198


}

let__9453 := ifres10183
_ = let__9453

tmp10201 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9453)


if True == tmp10201 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__9453
}


} else {
return let__9451
}


panic("unreachable")
})

tmp10204 := Call2(__e, ns2_1set, symshen_4_5literal_6, tmp10177)


_ = tmp10204

tmp10205 := MakeNative1(func(__e *ControlFlow, V1814 Obj) Obj {
tmp10206 := Call1(__e, SymFn(symshen_4_5bterm_6), V1814)


let__9462 := tmp10206
_ = let__9462

tmp10219 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9462)


var ifres10207 Obj

if True == tmp10219 {
tmp10208 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10207 = tmp10208


} else {
tmp10209 := Call1(__e, SymFn(symshen_4_5_1out), let__9462)


let__9463 := tmp10209
_ = let__9463

tmp10210 := Call1(__e, SymFn(symshen_4in_1_6), let__9462)


let__9464 := tmp10210
_ = let__9464

tmp10211 := Call1(__e, SymFn(symshen_4_5bterms_6), let__9464)


let__9465 := tmp10211
_ = let__9465

tmp10218 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9465)


var ifres10212 Obj

if True == tmp10218 {
tmp10213 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10212 = tmp10213


} else {
tmp10214 := Call1(__e, SymFn(symshen_4_5_1out), let__9465)


let__9466 := tmp10214
_ = let__9466

tmp10215 := Call1(__e, SymFn(symshen_4in_1_6), let__9465)


let__9467 := tmp10215
_ = let__9467

tmp10216 := GCons(__e, let__9463, let__9466)

tmp10217 := Call2(__e, SymFn(symshen_4comb), let__9467, tmp10216)


ifres10212 = tmp10217


}

ifres10207 = ifres10212


}

let__9461 := ifres10207
_ = let__9461

tmp10229 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9461)


if True == tmp10229 {
tmp10220 := Call1(__e, SymFn(sym_5e_6), V1814)


let__9469 := tmp10220
_ = let__9469

tmp10225 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9469)


var ifres10221 Obj

if True == tmp10225 {
tmp10222 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10221 = tmp10222


} else {
tmp10223 := Call1(__e, SymFn(symshen_4in_1_6), let__9469)


let__9470 := tmp10223
_ = let__9470

tmp10224 := Call2(__e, SymFn(symshen_4comb), let__9470, Nil)


ifres10221 = tmp10224


}

let__9468 := ifres10221
_ = let__9468

tmp10227 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9468)


if True == tmp10227 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__9468
}


} else {
return let__9461
}


panic("unreachable")
})

tmp10230 := Call2(__e, ns2_1set, symshen_4_5bterms_6, tmp10205)


_ = tmp10230

tmp10231 := MakeNative1(func(__e *ControlFlow, V1825 Obj) Obj {
tmp10232 := Call1(__e, SymFn(symshen_4_5wildcard_6), V1825)


let__9472 := tmp10232
_ = let__9472

tmp10238 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9472)


var ifres10233 Obj

if True == tmp10238 {
tmp10234 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10233 = tmp10234


} else {
tmp10235 := Call1(__e, SymFn(symshen_4_5_1out), let__9472)


let__9473 := tmp10235
_ = let__9473

tmp10236 := Call1(__e, SymFn(symshen_4in_1_6), let__9472)


let__9474 := tmp10236
_ = let__9474

tmp10237 := Call2(__e, SymFn(symshen_4comb), let__9474, let__9473)


ifres10233 = tmp10237


}

let__9471 := ifres10233
_ = let__9471

tmp10270 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9471)


if True == tmp10270 {
tmp10247 := GConsP(__e, V1825)

var ifres10239 Obj

if True == tmp10247 {
tmp10240 := Call1(__e, SymFn(symhead), V1825)


let__9476 := tmp10240
_ = let__9476

tmp10241 := Call1(__e, SymFn(symtail), V1825)


let__9477 := tmp10241
_ = let__9477

tmp10245 := Call1(__e, SymFn(symatom_2), let__9476)


var ifres10242 Obj

if True == tmp10245 {
tmp10243 := Call2(__e, SymFn(symshen_4comb), let__9477, let__9476)


ifres10242 = tmp10243


} else {
tmp10244 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10242 = tmp10244


}

ifres10239 = ifres10242


} else {
tmp10246 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10239 = tmp10246


}

let__9475 := ifres10239
_ = let__9475

tmp10268 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9475)


if True == tmp10268 {
tmp10264 := Call1(__e, SymFn(symshen_4ccons_2), V1825)


var ifres10248 Obj

if True == tmp10264 {
tmp10249 := Call1(__e, SymFn(symhead), V1825)


let__9479 := tmp10249
_ = let__9479

tmp10250 := Call1(__e, SymFn(symtail), V1825)


let__9480 := tmp10250
_ = let__9480

tmp10251 := Call1(__e, SymFn(symshen_4_5bterms_6), let__9479)


let__9481 := tmp10251
_ = let__9481

tmp10262 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9481)


var ifres10252 Obj

if True == tmp10262 {
tmp10253 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10252 = tmp10253


} else {
tmp10254 := Call1(__e, SymFn(symshen_4_5_1out), let__9481)


let__9482 := tmp10254
_ = let__9482

tmp10255 := Call1(__e, SymFn(symshen_4in_1_6), let__9481)


let__9483 := tmp10255
_ = let__9483

tmp10256 := Call1(__e, SymFn(sym_5end_6), let__9483)


let__9484 := tmp10256
_ = let__9484

tmp10261 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9484)


var ifres10257 Obj

if True == tmp10261 {
tmp10258 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10257 = tmp10258


} else {
tmp10259 := Call1(__e, SymFn(symshen_4in_1_6), let__9484)


let__9485 := tmp10259
_ = let__9485

tmp10260 := Call2(__e, SymFn(symshen_4comb), let__9480, let__9482)


ifres10257 = tmp10260


}

ifres10252 = ifres10257


}

ifres10248 = ifres10252


} else {
tmp10263 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10248 = tmp10263


}

let__9478 := ifres10248
_ = let__9478

tmp10266 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9478)


if True == tmp10266 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__9478
}


} else {
return let__9475
}


} else {
return let__9471
}


panic("unreachable")
})

tmp10271 := Call2(__e, ns2_1set, symshen_4_5bterm_6, tmp10231)


_ = tmp10271

tmp10272 := MakeNative1(func(__e *ControlFlow, V1841 Obj) Obj {
tmp10282 := GConsP(__e, V1841)

var ifres10273 Obj

if True == tmp10282 {
tmp10274 := Call1(__e, SymFn(symhead), V1841)


let__9487 := tmp10274
_ = let__9487

tmp10275 := Call1(__e, SymFn(symtail), V1841)


let__9488 := tmp10275
_ = let__9488

tmp10280 := GEq(__e, let__9487, sym__)

var ifres10276 Obj

if True == tmp10280 {
tmp10277 := Call1(__e, SymFn(symgensym), symY)


tmp10278 := Call2(__e, SymFn(symshen_4comb), let__9488, tmp10277)


ifres10276 = tmp10278


} else {
tmp10279 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10276 = tmp10279


}

ifres10273 = ifres10276


} else {
tmp10281 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10273 = tmp10281


}

let__9486 := ifres10273
_ = let__9486

tmp10284 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9486)


if True == tmp10284 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__9486
}


panic("unreachable")
})

tmp10285 := Call2(__e, ns2_1set, symshen_4_5wildcard_6, tmp10272)


_ = tmp10285

tmp10286 := MakeNative1(func(__e *ControlFlow, V1845 Obj) Obj {
tmp10295 := GConsP(__e, V1845)

var ifres10287 Obj

if True == tmp10295 {
tmp10288 := Call1(__e, SymFn(symhead), V1845)


let__9490 := tmp10288
_ = let__9490

tmp10289 := Call1(__e, SymFn(symtail), V1845)


let__9491 := tmp10289
_ = let__9491

tmp10293 := Call1(__e, SymFn(symshen_4semicolon_2), let__9490)


var ifres10290 Obj

if True == tmp10293 {
tmp10291 := Call2(__e, SymFn(symshen_4comb), let__9491, let__9490)


ifres10290 = tmp10291


} else {
tmp10292 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10290 = tmp10292


}

ifres10287 = ifres10290


} else {
tmp10294 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10287 = tmp10294


}

let__9489 := ifres10287
_ = let__9489

tmp10297 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9489)


if True == tmp10297 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__9489
}


panic("unreachable")
})

tmp10298 := Call2(__e, ns2_1set, symshen_4_5sc_6, tmp10286)


_ = tmp10298

tmp10299 := MakeNative2(func(__e *ControlFlow, V1849, V1850 Obj) Obj {
tmp10300 := Call1(__e, SymFn(symgensym), symB)


let__9492 := tmp10300
_ = let__9492

tmp10301 := Call1(__e, SymFn(symgensym), symL)


let__9493 := tmp10301
_ = let__9493

tmp10302 := Call1(__e, SymFn(symgensym), symK)


let__9494 := tmp10302
_ = let__9494

tmp10303 := Call1(__e, SymFn(symgensym), symC)


let__9495 := tmp10303
_ = let__9495

tmp10304 := Call1(__e, SymFn(symshen_4prolog_1parameters), V1850)


let__9496 := tmp10304
_ = let__9496

tmp10305 := Call1(__e, SymFn(symshen_4hascut_2), V1850)


let__9497 := tmp10305
_ = let__9497

tmp10306 := Call(__e, SymFn(symshen_4prolog_1fbody), V1850, let__9496, let__9492, let__9493, let__9494, let__9495, let__9497)


let__9498 := tmp10306
_ = let__9498

var ifres10307 Obj

if True == let__9497 {
tmp10308 := GCons(__e, MakeInteger(1), Nil)

tmp10309 := GCons(__e, let__9494, tmp10308)

tmp10310 := GCons(__e, sym_7, tmp10309)

tmp10311 := GCons(__e, let__9498, Nil)

tmp10312 := GCons(__e, tmp10310, tmp10311)

tmp10313 := GCons(__e, let__9494, tmp10312)

tmp10314 := GCons(__e, symlet, tmp10313)

ifres10307 = tmp10314


} else {
ifres10307 = let__9498


}

let__9499 := ifres10307
_ = let__9499

tmp10315 := GCons(__e, sym_1_6, Nil)

tmp10316 := GCons(__e, let__9495, tmp10315)

tmp10317 := GCons(__e, let__9494, tmp10316)

tmp10318 := GCons(__e, let__9493, tmp10317)

tmp10319 := GCons(__e, let__9492, tmp10318)

tmp10320 := GCons(__e, let__9499, Nil)

tmp10321 := Call2(__e, SymFn(symappend), tmp10319, tmp10320)


tmp10322 := Call2(__e, SymFn(symappend), let__9496, tmp10321)


tmp10323 := GCons(__e, V1849, tmp10322)

tmp10324 := GCons(__e, symdefine, tmp10323)

let__9500 := tmp10324
_ = let__9500

return let__9500


panic("unreachable")
})

tmp10325 := Call2(__e, ns2_1set, symshen_4horn_1clause_1procedure, tmp10299)


_ = tmp10325

tmp10326 := (func() Obj {
var __self Obj
__self = MakeNative1(func(__e *ControlFlow, __self1 Obj) Obj {
__selftop:
V1862 := __self1
_ = V1862
tmp10334 := GEq(__e, sym_b, V1862)

if True == tmp10334 {
return True
} else {
tmp10332 := GConsP(__e, V1862)

if True == tmp10332 {
tmp10329 := GHd(__e, V1862)

tmp10330 := Call1(__e, SymFn(symshen_4hascut_2), tmp10329)


if True == tmp10330 {
return True
} else {
tmp10327 := GTl(__e, V1862)

if SymFn(symshen_4hascut_2) == __self {
__self1 = tmp10327
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4hascut_2), tmp10327)


}


} else {
return False
}


}


panic("unreachable")
})
return __self
})()

tmp10335 := Call2(__e, ns2_1set, symshen_4hascut_2, tmp10326)


_ = tmp10335

tmp10336 := MakeNative1(func(__e *ControlFlow, V1867 Obj) Obj {
tmp10344 := GConsP(__e, V1867)

var ifres10341 Obj

if True == tmp10344 {
tmp10342 := GHd(__e, V1867)

tmp10343 := GConsP(__e, tmp10342)

ifres10341 = tmp10343


} else {
ifres10341 = False


}

if True == ifres10341 {
tmp10337 := GHd(__e, V1867)

tmp10338 := GHd(__e, tmp10337)

tmp10339 := Call1(__e, SymFn(symlength), tmp10338)


return TailCall1(__e, SymFn(symshen_4parameters), tmp10339)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4prolog_1parameters)
}


panic("unreachable")
})

tmp10345 := Call2(__e, ns2_1set, symshen_4prolog_1parameters, tmp10336)


_ = tmp10345

tmp10346 := MakeNative(func(__e *ControlFlow) {
V1888 := __e.Get(1)
V1889 := __e.Get(2)
V1890 := __e.Get(3)
V1891 := __e.Get(4)
V1892 := __e.Get(5)
V1893 := __e.Get(6)
V1894 := __e.Get(7)
__e.ReturnDirect(func() Obj {
tmp10427 := GEq(__e, Nil, V1888)

var ifres10425 Obj

if True == tmp10427 {
tmp10426 := GEq(__e, True, V1894)

ifres10425 = tmp10426


} else {
ifres10425 = False


}

if True == ifres10425 {
tmp10347 := GCons(__e, V1892, Nil)

tmp10348 := GCons(__e, V1891, tmp10347)

return GCons(__e, symshen_4unlock, tmp10348)


} else {
tmp10423 := GConsP(__e, V1888)

var ifres10406 Obj

if True == tmp10423 {
tmp10421 := GHd(__e, V1888)

tmp10422 := GConsP(__e, tmp10421)

var ifres10407 Obj

if True == tmp10422 {
tmp10418 := GHd(__e, V1888)

tmp10419 := GTl(__e, tmp10418)

tmp10420 := GConsP(__e, tmp10419)

var ifres10408 Obj

if True == tmp10420 {
tmp10414 := GHd(__e, V1888)

tmp10415 := GTl(__e, tmp10414)

tmp10416 := GTl(__e, tmp10415)

tmp10417 := GEq(__e, Nil, tmp10416)

var ifres10409 Obj

if True == tmp10417 {
tmp10412 := GTl(__e, V1888)

tmp10413 := GEq(__e, Nil, tmp10412)

var ifres10410 Obj

if True == tmp10413 {
tmp10411 := GEq(__e, False, V1894)

ifres10410 = tmp10411


} else {
ifres10410 = False


}

ifres10409 = ifres10410


} else {
ifres10409 = False


}

ifres10408 = ifres10409


} else {
ifres10408 = False


}

ifres10407 = ifres10408


} else {
ifres10407 = False


}

ifres10406 = ifres10407


} else {
ifres10406 = False


}

if True == ifres10406 {
tmp10349 := GHd(__e, V1888)

tmp10350 := GHd(__e, tmp10349)

tmp10351 := GHd(__e, V1888)

tmp10352 := GTl(__e, tmp10351)

tmp10353 := GHd(__e, tmp10352)

tmp10354 := Call6(__e, SymFn(symshen_4continue), tmp10350, tmp10353, V1890, V1891, V1892, V1893)


let__9501 := tmp10354
_ = let__9501

tmp10355 := GCons(__e, V1891, Nil)

tmp10356 := GCons(__e, symshen_4unlocked_2, tmp10355)

tmp10357 := GHd(__e, V1888)

tmp10358 := GHd(__e, tmp10357)

tmp10359 := Call5(__e, SymFn(symshen_4compile_1head), symshen_4_7m, tmp10358, V1889, V1890, let__9501)


tmp10360 := GCons(__e, False, Nil)

tmp10361 := GCons(__e, tmp10359, tmp10360)

tmp10362 := GCons(__e, tmp10356, tmp10361)

return GCons(__e, symif, tmp10362)


} else {
tmp10404 := GConsP(__e, V1888)

var ifres10392 Obj

if True == tmp10404 {
tmp10402 := GHd(__e, V1888)

tmp10403 := GConsP(__e, tmp10402)

var ifres10393 Obj

if True == tmp10403 {
tmp10399 := GHd(__e, V1888)

tmp10400 := GTl(__e, tmp10399)

tmp10401 := GConsP(__e, tmp10400)

var ifres10394 Obj

if True == tmp10401 {
tmp10395 := GHd(__e, V1888)

tmp10396 := GTl(__e, tmp10395)

tmp10397 := GTl(__e, tmp10396)

tmp10398 := GEq(__e, Nil, tmp10397)

ifres10394 = tmp10398


} else {
ifres10394 = False


}

ifres10393 = ifres10394


} else {
ifres10393 = False


}

ifres10392 = ifres10393


} else {
ifres10392 = False


}

if True == ifres10392 {
tmp10363 := Call1(__e, SymFn(symgensym), symC)


let__9502 := tmp10363
_ = let__9502

tmp10364 := GHd(__e, V1888)

tmp10365 := GHd(__e, tmp10364)

tmp10366 := GHd(__e, V1888)

tmp10367 := GTl(__e, tmp10366)

tmp10368 := GHd(__e, tmp10367)

tmp10369 := Call6(__e, SymFn(symshen_4continue), tmp10365, tmp10368, V1890, V1891, V1892, V1893)


let__9503 := tmp10369
_ = let__9503

tmp10370 := GCons(__e, V1891, Nil)

tmp10371 := GCons(__e, symshen_4unlocked_2, tmp10370)

tmp10372 := GHd(__e, V1888)

tmp10373 := GHd(__e, tmp10372)

tmp10374 := Call5(__e, SymFn(symshen_4compile_1head), symshen_4_7m, tmp10373, V1889, V1890, let__9503)


tmp10375 := GCons(__e, False, Nil)

tmp10376 := GCons(__e, tmp10374, tmp10375)

tmp10377 := GCons(__e, tmp10371, tmp10376)

tmp10378 := GCons(__e, symif, tmp10377)

tmp10379 := GCons(__e, False, Nil)

tmp10380 := GCons(__e, let__9502, tmp10379)

tmp10381 := GCons(__e, sym_a, tmp10380)

tmp10382 := GTl(__e, V1888)

tmp10383 := Call(__e, SymFn(symshen_4prolog_1fbody), tmp10382, V1889, V1890, V1891, V1892, V1893, V1894)


tmp10384 := GCons(__e, let__9502, Nil)

tmp10385 := GCons(__e, tmp10383, tmp10384)

tmp10386 := GCons(__e, tmp10381, tmp10385)

tmp10387 := GCons(__e, symif, tmp10386)

tmp10388 := GCons(__e, tmp10387, Nil)

tmp10389 := GCons(__e, tmp10378, tmp10388)

tmp10390 := GCons(__e, let__9502, tmp10389)

return GCons(__e, symlet, tmp10390)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.prolog-fbody"))
}


}


}


panic("unreachable")
}())
}, 7)

tmp10428 := Call2(__e, ns2_1set, symshen_4prolog_1fbody, tmp10346)


_ = tmp10428

tmp10429 := MakeNative2(func(__e *ControlFlow, V1898, V1899 Obj) Obj {
tmp10433 := Call1(__e, SymFn(symshen_4locked_2), V1898)


var ifres10431 Obj

if True == tmp10433 {
tmp10432 := Call2(__e, SymFn(symshen_4fits_2), V1899, V1898)


ifres10431 = tmp10432


} else {
ifres10431 = False


}

if True == ifres10431 {
return TailCall1(__e, SymFn(symshen_4openlock), V1898)
} else {
return False
}


panic("unreachable")
})

tmp10434 := Call2(__e, ns2_1set, symshen_4unlock, tmp10429)


_ = tmp10434

tmp10435 := MakeNative1(func(__e *ControlFlow, V1900 Obj) Obj {
tmp10436 := Call1(__e, SymFn(symshen_4unlocked_2), V1900)


return GNot(__e, tmp10436)


panic("unreachable")
})

tmp10437 := Call2(__e, ns2_1set, symshen_4locked_2, tmp10435)


_ = tmp10437

tmp10438 := MakeNative1(func(__e *ControlFlow, V1901 Obj) Obj {
return GVecGet(__e, V1901, MakeInteger(1))
panic("unreachable")
})

tmp10439 := Call2(__e, ns2_1set, symshen_4unlocked_2, tmp10438)


_ = tmp10439

tmp10440 := MakeNative1(func(__e *ControlFlow, V1902 Obj) Obj {
tmp10441 := GVecSet(__e, V1902, MakeInteger(1), True)

_ = tmp10441

return False


panic("unreachable")
})

tmp10442 := Call2(__e, ns2_1set, symshen_4openlock, tmp10440)


_ = tmp10442

tmp10443 := MakeNative2(func(__e *ControlFlow, V1903, V1904 Obj) Obj {
tmp10444 := GVecGet(__e, V1904, MakeInteger(2))

return GEq(__e, V1903, tmp10444)


panic("unreachable")
})

tmp10445 := Call2(__e, ns2_1set, symshen_4fits_2, tmp10443)


_ = tmp10445

tmp10446 := MakeNative4(func(__e *ControlFlow, V1907, V1908, V1909, V1910 Obj) Obj {
tmp10447 := Call1(__e, SymFn(symthaw), V1910)


let__9504 := tmp10447
_ = let__9504

tmp10451 := GEq(__e, let__9504, False)

var ifres10449 Obj

if True == tmp10451 {
tmp10450 := Call1(__e, SymFn(symshen_4unlocked_2), V1908)


ifres10449 = tmp10450


} else {
ifres10449 = False


}

if True == ifres10449 {
return TailCall2(__e, SymFn(symshen_4lock), V1909, V1908)
} else {
return let__9504
}


panic("unreachable")
})

tmp10452 := Call2(__e, ns2_1set, symshen_4cut, tmp10446)


_ = tmp10452

tmp10453 := MakeNative2(func(__e *ControlFlow, V1912, V1913 Obj) Obj {
tmp10454 := GVecSet(__e, V1913, MakeInteger(1), False)

let__9505 := tmp10454
_ = let__9505

tmp10455 := GVecSet(__e, V1913, MakeInteger(2), V1912)

let__9506 := tmp10455
_ = let__9506

return False


panic("unreachable")
})

tmp10456 := Call2(__e, ns2_1set, symshen_4lock, tmp10453)


_ = tmp10456

tmp10457 := MakeNative6(func(__e *ControlFlow, V1916, V1917, V1918, V1919, V1920, V1921 Obj) Obj {
tmp10458 := Call1(__e, SymFn(symshen_4extract_1vars), V1916)


let__9507 := tmp10458
_ = let__9507

tmp10459 := Call1(__e, SymFn(symshen_4extract_1free_1vars), V1917)


let__9508 := tmp10459
_ = let__9508

tmp10460 := Call2(__e, SymFn(symdifference), let__9508, let__9507)


let__9509 := tmp10460
_ = let__9509

tmp10461 := GCons(__e, symshen_4incinfs, Nil)

tmp10462 := Call5(__e, SymFn(symshen_4compile_1body), V1917, V1918, V1919, V1920, V1921)


tmp10463 := GCons(__e, tmp10462, Nil)

tmp10464 := GCons(__e, tmp10461, tmp10463)

tmp10465 := GCons(__e, symdo, tmp10464)

let__9510 := tmp10465
_ = let__9510

return TailCall3(__e, SymFn(symshen_4stpart), let__9509, let__9510, V1918)


panic("unreachable")
})

tmp10466 := Call2(__e, ns2_1set, symshen_4continue, tmp10457)


_ = tmp10466

tmp10467 := MakeNative1(func(__e *ControlFlow, V1928 Obj) Obj {
tmp10498 := GConsP(__e, V1928)

var ifres10483 Obj

if True == tmp10498 {
tmp10496 := GHd(__e, V1928)

tmp10497 := GEq(__e, symlambda, tmp10496)

var ifres10484 Obj

if True == tmp10497 {
tmp10494 := GTl(__e, V1928)

tmp10495 := GConsP(__e, tmp10494)

var ifres10485 Obj

if True == tmp10495 {
tmp10491 := GTl(__e, V1928)

tmp10492 := GTl(__e, tmp10491)

tmp10493 := GConsP(__e, tmp10492)

var ifres10486 Obj

if True == tmp10493 {
tmp10487 := GTl(__e, V1928)

tmp10488 := GTl(__e, tmp10487)

tmp10489 := GTl(__e, tmp10488)

tmp10490 := GEq(__e, Nil, tmp10489)

ifres10486 = tmp10490


} else {
ifres10486 = False


}

ifres10485 = ifres10486


} else {
ifres10485 = False


}

ifres10484 = ifres10485


} else {
ifres10484 = False


}

ifres10483 = ifres10484


} else {
ifres10483 = False


}

if True == ifres10483 {
tmp10468 := GTl(__e, V1928)

tmp10469 := GHd(__e, tmp10468)

tmp10470 := GTl(__e, V1928)

tmp10471 := GTl(__e, tmp10470)

tmp10472 := GHd(__e, tmp10471)

tmp10473 := Call1(__e, SymFn(symshen_4extract_1free_1vars), tmp10472)


return TailCall2(__e, SymFn(symremove), tmp10469, tmp10473)


} else {
tmp10481 := GConsP(__e, V1928)

if True == tmp10481 {
tmp10474 := GHd(__e, V1928)

tmp10475 := Call1(__e, SymFn(symshen_4extract_1free_1vars), tmp10474)


tmp10476 := GTl(__e, V1928)

tmp10477 := Call1(__e, SymFn(symshen_4extract_1free_1vars), tmp10476)


return TailCall2(__e, SymFn(symunion), tmp10475, tmp10477)


} else {
tmp10479 := PrimGuard1(__e, symvariable_2, PrimIsVariable, V1928)

if True == tmp10479 {
return GCons(__e, V1928, Nil)
} else {
return Nil
}


}


}


panic("unreachable")
})

tmp10499 := Call2(__e, ns2_1set, symshen_4extract_1free_1vars, tmp10467)


_ = tmp10499

tmp10500 := (func() Obj {
var __self Obj
__self = MakeNative5(func(__e *ControlFlow, __self1, __self2, __self3, __self4, __self5 Obj) Obj {
__selftop:
V1945 := __self1
_ = V1945
V1946 := __self2
_ = V1946
V1947 := __self3
_ = V1947
V1948 := __self4
_ = V1948
V1949 := __self5
_ = V1949
tmp10532 := GEq(__e, Nil, V1945)

if True == tmp10532 {
tmp10501 := GCons(__e, V1949, Nil)

return GCons(__e, symthaw, tmp10501)


} else {
tmp10530 := GConsP(__e, V1945)

var ifres10527 Obj

if True == tmp10530 {
tmp10528 := GHd(__e, V1945)

tmp10529 := GEq(__e, sym_b, tmp10528)

ifres10527 = tmp10529


} else {
ifres10527 = False


}

if True == ifres10527 {
tmp10502 := GCons(__e, symshen_4cut, Nil)

tmp10503 := GTl(__e, V1945)

tmp10504 := GCons(__e, tmp10502, tmp10503)

if SymFn(symshen_4compile_1body) == __self {
__self1, __self2, __self3, __self4, __self5 = tmp10504, V1946, V1947, V1948, V1949
__e.Tick()
goto __selftop
}
return TailCall5(__e, SymFn(symshen_4compile_1body), tmp10504, V1946, V1947, V1948, V1949)


} else {
tmp10525 := GConsP(__e, V1945)

var ifres10522 Obj

if True == tmp10525 {
tmp10523 := GTl(__e, V1945)

tmp10524 := GEq(__e, Nil, tmp10523)

ifres10522 = tmp10524


} else {
ifres10522 = False


}

if True == ifres10522 {
tmp10505 := GHd(__e, V1945)

tmp10506 := Call2(__e, SymFn(symshen_4deref_1calls), tmp10505, V1946)


tmp10507 := GCons(__e, V1949, Nil)

tmp10508 := GCons(__e, V1948, tmp10507)

tmp10509 := GCons(__e, V1947, tmp10508)

tmp10510 := GCons(__e, V1946, tmp10509)

return TailCall2(__e, SymFn(symappend), tmp10506, tmp10510)


} else {
tmp10520 := GConsP(__e, V1945)

if True == tmp10520 {
tmp10511 := GHd(__e, V1945)

tmp10512 := Call2(__e, SymFn(symshen_4deref_1calls), tmp10511, V1946)


let__9511 := tmp10512
_ = let__9511

tmp10513 := GTl(__e, V1945)

tmp10514 := Call5(__e, SymFn(symshen_4freeze_1literals), tmp10513, V1946, V1947, V1948, V1949)


tmp10515 := GCons(__e, tmp10514, Nil)

tmp10516 := GCons(__e, V1948, tmp10515)

tmp10517 := GCons(__e, V1947, tmp10516)

tmp10518 := GCons(__e, V1946, tmp10517)

return TailCall2(__e, SymFn(symappend), let__9511, tmp10518)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.compile-fbody"))
}


}


}


}


panic("unreachable")
})
return __self
})()

tmp10533 := Call2(__e, ns2_1set, symshen_4compile_1body, tmp10500)


_ = tmp10533

tmp10534 := (func() Obj {
var __self Obj
__self = MakeNative5(func(__e *ControlFlow, __self1, __self2, __self3, __self4, __self5 Obj) Obj {
__selftop:
V1967 := __self1
_ = V1967
V1968 := __self2
_ = V1968
V1969 := __self3
_ = V1969
V1970 := __self4
_ = V1970
V1971 := __self5
_ = V1971
tmp10556 := GEq(__e, Nil, V1967)

if True == tmp10556 {
return V1971
} else {
tmp10554 := GConsP(__e, V1967)

var ifres10551 Obj

if True == tmp10554 {
tmp10552 := GHd(__e, V1967)

tmp10553 := GEq(__e, sym_b, tmp10552)

ifres10551 = tmp10553


} else {
ifres10551 = False


}

if True == ifres10551 {
tmp10535 := GCons(__e, symshen_4cut, Nil)

tmp10536 := GTl(__e, V1967)

tmp10537 := GCons(__e, tmp10535, tmp10536)

if SymFn(symshen_4freeze_1literals) == __self {
__self1, __self2, __self3, __self4, __self5 = tmp10537, V1968, V1969, V1970, V1971
__e.Tick()
goto __selftop
}
return TailCall5(__e, SymFn(symshen_4freeze_1literals), tmp10537, V1968, V1969, V1970, V1971)


} else {
tmp10549 := GConsP(__e, V1967)

if True == tmp10549 {
tmp10538 := GHd(__e, V1967)

tmp10539 := Call2(__e, SymFn(symshen_4deref_1calls), tmp10538, V1968)


let__9512 := tmp10539
_ = let__9512

tmp10540 := GTl(__e, V1967)

tmp10541 := Call5(__e, SymFn(symshen_4freeze_1literals), tmp10540, V1968, V1969, V1970, V1971)


tmp10542 := GCons(__e, tmp10541, Nil)

tmp10543 := GCons(__e, V1970, tmp10542)

tmp10544 := GCons(__e, V1969, tmp10543)

tmp10545 := GCons(__e, V1968, tmp10544)

tmp10546 := Call2(__e, SymFn(symappend), let__9512, tmp10545)


tmp10547 := GCons(__e, tmp10546, Nil)

return GCons(__e, symfreeze, tmp10547)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.freeze-literals"))
}


}


}


panic("unreachable")
})
return __self
})()

tmp10557 := Call2(__e, ns2_1set, symshen_4freeze_1literals, tmp10534)


_ = tmp10557

tmp10558 := MakeNative2(func(__e *ControlFlow, V1977, V1978 Obj) Obj {
tmp10572 := GConsP(__e, V1977)

var ifres10569 Obj

if True == tmp10572 {
tmp10570 := GHd(__e, V1977)

tmp10571 := GEq(__e, symfork, tmp10570)

ifres10569 = tmp10571


} else {
ifres10569 = False


}

if True == ifres10569 {
tmp10559 := GTl(__e, V1977)

tmp10560 := Call2(__e, SymFn(symshen_4deref_1forked_1literals), tmp10559, V1978)


tmp10561 := GCons(__e, tmp10560, Nil)

return GCons(__e, symfork, tmp10561)


} else {
tmp10567 := GConsP(__e, V1977)

if True == tmp10567 {
tmp10562 := GHd(__e, V1977)

tmp10563 := MakeNative1(func(__e *ControlFlow, Z1979 Obj) Obj {
return TailCall2(__e, SymFn(symshen_4function_1calls), Z1979, V1978)
panic("unreachable")
})

tmp10564 := GTl(__e, V1977)

tmp10565 := Call2(__e, SymFn(symmap), tmp10563, tmp10564)


return GCons(__e, tmp10562, tmp10565)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.deref-calls"))
}


}


panic("unreachable")
})

tmp10573 := Call2(__e, ns2_1set, symshen_4deref_1calls, tmp10558)


_ = tmp10573

tmp10574 := MakeNative2(func(__e *ControlFlow, V1986, V1987 Obj) Obj {
tmp10584 := GEq(__e, Nil, V1986)

if True == tmp10584 {
return Nil
} else {
tmp10582 := GConsP(__e, V1986)

if True == tmp10582 {
tmp10575 := GHd(__e, V1986)

tmp10576 := Call2(__e, SymFn(symshen_4deref_1calls), tmp10575, V1987)


tmp10577 := GTl(__e, V1986)

tmp10578 := Call2(__e, SymFn(symshen_4deref_1forked_1literals), tmp10577, V1987)


tmp10579 := GCons(__e, tmp10578, Nil)

tmp10580 := GCons(__e, tmp10576, tmp10579)

return GCons(__e, symcons, tmp10580)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("fork requires a list of literals\n"))
}


}


panic("unreachable")
})

tmp10585 := Call2(__e, ns2_1set, symshen_4deref_1forked_1literals, tmp10574)


_ = tmp10585

tmp10586 := MakeNative2(func(__e *ControlFlow, V1990, V1991 Obj) Obj {
tmp10614 := GConsP(__e, V1990)

var ifres10599 Obj

if True == tmp10614 {
tmp10612 := GHd(__e, V1990)

tmp10613 := GEq(__e, symcons, tmp10612)

var ifres10600 Obj

if True == tmp10613 {
tmp10610 := GTl(__e, V1990)

tmp10611 := GConsP(__e, tmp10610)

var ifres10601 Obj

if True == tmp10611 {
tmp10607 := GTl(__e, V1990)

tmp10608 := GTl(__e, tmp10607)

tmp10609 := GConsP(__e, tmp10608)

var ifres10602 Obj

if True == tmp10609 {
tmp10603 := GTl(__e, V1990)

tmp10604 := GTl(__e, tmp10603)

tmp10605 := GTl(__e, tmp10604)

tmp10606 := GEq(__e, Nil, tmp10605)

ifres10602 = tmp10606


} else {
ifres10602 = False


}

ifres10601 = ifres10602


} else {
ifres10601 = False


}

ifres10600 = ifres10601


} else {
ifres10600 = False


}

ifres10599 = ifres10600


} else {
ifres10599 = False


}

if True == ifres10599 {
tmp10587 := GTl(__e, V1990)

tmp10588 := GHd(__e, tmp10587)

tmp10589 := Call2(__e, SymFn(symshen_4function_1calls), tmp10588, V1991)


tmp10590 := GTl(__e, V1990)

tmp10591 := GTl(__e, tmp10590)

tmp10592 := GHd(__e, tmp10591)

tmp10593 := Call2(__e, SymFn(symshen_4function_1calls), tmp10592, V1991)


tmp10594 := GCons(__e, tmp10593, Nil)

tmp10595 := GCons(__e, tmp10589, tmp10594)

return GCons(__e, symcons, tmp10595)


} else {
tmp10597 := GConsP(__e, V1990)

if True == tmp10597 {
return TailCall3(__e, SymFn(symshen_4deref_1terms), V1990, V1991, Nil)
} else {
return V1990
}


}


panic("unreachable")
})

tmp10615 := Call2(__e, ns2_1set, symshen_4function_1calls, tmp10586)


_ = tmp10615

tmp10616 := MakeNative3(func(__e *ControlFlow, V2000, V2001, V2002 Obj) Obj {
tmp10699 := GConsP(__e, V2000)

var ifres10689 Obj

if True == tmp10699 {
tmp10697 := GHd(__e, V2000)

tmp10698 := GEq(__e, MakeInteger(0), tmp10697)

var ifres10690 Obj

if True == tmp10698 {
tmp10695 := GTl(__e, V2000)

tmp10696 := GConsP(__e, tmp10695)

var ifres10691 Obj

if True == tmp10696 {
tmp10692 := GTl(__e, V2000)

tmp10693 := GTl(__e, tmp10692)

tmp10694 := GEq(__e, Nil, tmp10693)

ifres10691 = tmp10694


} else {
ifres10691 = False


}

ifres10690 = ifres10691


} else {
ifres10690 = False


}

ifres10689 = ifres10690


} else {
ifres10689 = False


}

if True == ifres10689 {
tmp10623 := GTl(__e, V2000)

tmp10624 := GHd(__e, tmp10623)

tmp10625 := PrimGuard1(__e, symvariable_2, PrimIsVariable, tmp10624)

if True == tmp10625 {
tmp10617 := GTl(__e, V2000)

return GHd(__e, tmp10617)


} else {
tmp10618 := GTl(__e, V2000)

tmp10619 := GHd(__e, tmp10618)

tmp10620 := Call3(__e, SymFn(symshen_4app), tmp10619, MakeString("\n"), symshen_4s)


return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("attempt to optimise a non-variable "), tmp10620))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("attempt to optimise a non-variable "), tmp10620)
return Call1(__e, SymFn(symsimple_1error), __typedArg0)
})()


}


} else {
tmp10687 := GConsP(__e, V2000)

var ifres10677 Obj

if True == tmp10687 {
tmp10685 := GHd(__e, V2000)

tmp10686 := GEq(__e, MakeInteger(1), tmp10685)

var ifres10678 Obj

if True == tmp10686 {
tmp10683 := GTl(__e, V2000)

tmp10684 := GConsP(__e, tmp10683)

var ifres10679 Obj

if True == tmp10684 {
tmp10680 := GTl(__e, V2000)

tmp10681 := GTl(__e, tmp10680)

tmp10682 := GEq(__e, Nil, tmp10681)

ifres10679 = tmp10682


} else {
ifres10679 = False


}

ifres10678 = ifres10679


} else {
ifres10678 = False


}

ifres10677 = ifres10678


} else {
ifres10677 = False


}

if True == ifres10677 {
tmp10635 := GTl(__e, V2000)

tmp10636 := GHd(__e, tmp10635)

tmp10637 := PrimGuard1(__e, symvariable_2, PrimIsVariable, tmp10636)

if True == tmp10637 {
tmp10626 := GTl(__e, V2000)

tmp10627 := GHd(__e, tmp10626)

tmp10628 := GCons(__e, V2001, Nil)

tmp10629 := GCons(__e, tmp10627, tmp10628)

return GCons(__e, symshen_4lazyderef, tmp10629)


} else {
tmp10630 := GTl(__e, V2000)

tmp10631 := GHd(__e, tmp10630)

tmp10632 := Call3(__e, SymFn(symshen_4app), tmp10631, MakeString("\n"), symshen_4s)


return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("attempt to optimise a non-variable "), tmp10632))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("attempt to optimise a non-variable "), tmp10632)
return Call1(__e, SymFn(symsimple_1error), __typedArg0)
})()


}


} else {
tmp10674 := Call2(__e, SymFn(symelement_2), V2000, V2002)


tmp10675 := GNot(__e, tmp10674)

var ifres10672 Obj

if True == tmp10675 {
tmp10673 := PrimGuard1(__e, symvariable_2, PrimIsVariable, V2000)

ifres10672 = tmp10673


} else {
ifres10672 = False


}

if True == ifres10672 {
tmp10638 := GCons(__e, V2001, Nil)

tmp10639 := GCons(__e, V2000, tmp10638)

return GCons(__e, symshen_4deref, tmp10639)


} else {
tmp10670 := GConsP(__e, V2000)

var ifres10655 Obj

if True == tmp10670 {
tmp10668 := GHd(__e, V2000)

tmp10669 := GEq(__e, symlambda, tmp10668)

var ifres10656 Obj

if True == tmp10669 {
tmp10666 := GTl(__e, V2000)

tmp10667 := GConsP(__e, tmp10666)

var ifres10657 Obj

if True == tmp10667 {
tmp10663 := GTl(__e, V2000)

tmp10664 := GTl(__e, tmp10663)

tmp10665 := GConsP(__e, tmp10664)

var ifres10658 Obj

if True == tmp10665 {
tmp10659 := GTl(__e, V2000)

tmp10660 := GTl(__e, tmp10659)

tmp10661 := GTl(__e, tmp10660)

tmp10662 := GEq(__e, Nil, tmp10661)

ifres10658 = tmp10662


} else {
ifres10658 = False


}

ifres10657 = ifres10658


} else {
ifres10657 = False


}

ifres10656 = ifres10657


} else {
ifres10656 = False


}

ifres10655 = ifres10656


} else {
ifres10655 = False


}

if True == ifres10655 {
tmp10640 := GTl(__e, V2000)

tmp10641 := GHd(__e, tmp10640)

tmp10642 := GTl(__e, V2000)

tmp10643 := GTl(__e, tmp10642)

tmp10644 := GHd(__e, tmp10643)

tmp10645 := GTl(__e, V2000)

tmp10646 := GHd(__e, tmp10645)

tmp10647 := GCons(__e, tmp10646, V2002)

tmp10648 := Call3(__e, SymFn(symshen_4deref_1terms), tmp10644, V2001, tmp10647)


tmp10649 := GCons(__e, tmp10648, Nil)

tmp10650 := GCons(__e, tmp10641, tmp10649)

return GCons(__e, symlambda, tmp10650)


} else {
tmp10653 := GConsP(__e, V2000)

if True == tmp10653 {
tmp10651 := MakeNative1(func(__e *ControlFlow, Z2003 Obj) Obj {
return TailCall3(__e, SymFn(symshen_4deref_1terms), Z2003, V2001, V2002)
panic("unreachable")
})

return TailCall2(__e, SymFn(symmap), tmp10651, V2000)


} else {
return V2000
}


}


}


}


}


panic("unreachable")
})

tmp10700 := Call2(__e, ns2_1set, symshen_4deref_1terms, tmp10616)


_ = tmp10700

tmp10701 := (func() Obj {
var __self Obj
__self = MakeNative5(func(__e *ControlFlow, __self1, __self2, __self3, __self4, __self5 Obj) Obj {
__selftop:
V2021 := __self1
_ = V2021
V2022 := __self2
_ = V2022
V2023 := __self3
_ = V2023
V2024 := __self4
_ = V2024
V2025 := __self5
_ = V2025
tmp10847 := GEq(__e, Nil, V2022)

var ifres10845 Obj

if True == tmp10847 {
tmp10846 := GEq(__e, Nil, V2023)

ifres10845 = tmp10846


} else {
ifres10845 = False


}

if True == ifres10845 {
return V2025
} else {
tmp10843 := GConsP(__e, V2022)

var ifres10827 Obj

if True == tmp10843 {
tmp10841 := GHd(__e, V2022)

tmp10842 := GConsP(__e, tmp10841)

var ifres10828 Obj

if True == tmp10842 {
tmp10838 := GHd(__e, V2022)

tmp10839 := GHd(__e, tmp10838)

tmp10840 := GEq(__e, symshen_4_7m, tmp10839)

var ifres10829 Obj

if True == tmp10840 {
tmp10835 := GHd(__e, V2022)

tmp10836 := GTl(__e, tmp10835)

tmp10837 := GConsP(__e, tmp10836)

var ifres10830 Obj

if True == tmp10837 {
tmp10831 := GHd(__e, V2022)

tmp10832 := GTl(__e, tmp10831)

tmp10833 := GTl(__e, tmp10832)

tmp10834 := GEq(__e, Nil, tmp10833)

ifres10830 = tmp10834


} else {
ifres10830 = False


}

ifres10829 = ifres10830


} else {
ifres10829 = False


}

ifres10828 = ifres10829


} else {
ifres10828 = False


}

ifres10827 = ifres10828


} else {
ifres10827 = False


}

if True == ifres10827 {
tmp10702 := GHd(__e, V2022)

tmp10703 := GTl(__e, tmp10702)

tmp10704 := GHd(__e, tmp10703)

tmp10705 := GTl(__e, V2022)

tmp10706 := GCons(__e, V2021, tmp10705)

tmp10707 := GCons(__e, tmp10704, tmp10706)

tmp10708 := GCons(__e, symshen_4_7m, tmp10707)

if SymFn(symshen_4compile_1head) == __self {
__self1, __self2, __self3, __self4, __self5 = V2021, tmp10708, V2023, V2024, V2025
__e.Tick()
goto __selftop
}
return TailCall5(__e, SymFn(symshen_4compile_1head), V2021, tmp10708, V2023, V2024, V2025)


} else {
tmp10825 := GConsP(__e, V2022)

var ifres10809 Obj

if True == tmp10825 {
tmp10823 := GHd(__e, V2022)

tmp10824 := GConsP(__e, tmp10823)

var ifres10810 Obj

if True == tmp10824 {
tmp10820 := GHd(__e, V2022)

tmp10821 := GHd(__e, tmp10820)

tmp10822 := GEq(__e, symshen_4_1m, tmp10821)

var ifres10811 Obj

if True == tmp10822 {
tmp10817 := GHd(__e, V2022)

tmp10818 := GTl(__e, tmp10817)

tmp10819 := GConsP(__e, tmp10818)

var ifres10812 Obj

if True == tmp10819 {
tmp10813 := GHd(__e, V2022)

tmp10814 := GTl(__e, tmp10813)

tmp10815 := GTl(__e, tmp10814)

tmp10816 := GEq(__e, Nil, tmp10815)

ifres10812 = tmp10816


} else {
ifres10812 = False


}

ifres10811 = ifres10812


} else {
ifres10811 = False


}

ifres10810 = ifres10811


} else {
ifres10810 = False


}

ifres10809 = ifres10810


} else {
ifres10809 = False


}

if True == ifres10809 {
tmp10709 := GHd(__e, V2022)

tmp10710 := GTl(__e, tmp10709)

tmp10711 := GHd(__e, tmp10710)

tmp10712 := GTl(__e, V2022)

tmp10713 := GCons(__e, V2021, tmp10712)

tmp10714 := GCons(__e, tmp10711, tmp10713)

tmp10715 := GCons(__e, symshen_4_1m, tmp10714)

if SymFn(symshen_4compile_1head) == __self {
__self1, __self2, __self3, __self4, __self5 = V2021, tmp10715, V2023, V2024, V2025
__e.Tick()
goto __selftop
}
return TailCall5(__e, SymFn(symshen_4compile_1head), V2021, tmp10715, V2023, V2024, V2025)


} else {
tmp10807 := GConsP(__e, V2022)

var ifres10804 Obj

if True == tmp10807 {
tmp10805 := GHd(__e, V2022)

tmp10806 := GEq(__e, symshen_4_1m, tmp10805)

ifres10804 = tmp10806


} else {
ifres10804 = False


}

if True == ifres10804 {
tmp10716 := GTl(__e, V2022)

if SymFn(symshen_4compile_1head) == __self {
__self1, __self2, __self3, __self4, __self5 = symshen_4_1m, tmp10716, V2023, V2024, V2025
__e.Tick()
goto __selftop
}
return TailCall5(__e, SymFn(symshen_4compile_1head), symshen_4_1m, tmp10716, V2023, V2024, V2025)


} else {
tmp10802 := GConsP(__e, V2022)

var ifres10799 Obj

if True == tmp10802 {
tmp10800 := GHd(__e, V2022)

tmp10801 := GEq(__e, symshen_4_7m, tmp10800)

ifres10799 = tmp10801


} else {
ifres10799 = False


}

if True == ifres10799 {
tmp10717 := GTl(__e, V2022)

if SymFn(symshen_4compile_1head) == __self {
__self1, __self2, __self3, __self4, __self5 = symshen_4_7m, tmp10717, V2023, V2024, V2025
__e.Tick()
goto __selftop
}
return TailCall5(__e, SymFn(symshen_4compile_1head), symshen_4_7m, tmp10717, V2023, V2024, V2025)


} else {
tmp10797 := GConsP(__e, V2022)

var ifres10792 Obj

if True == tmp10797 {
tmp10796 := GConsP(__e, V2023)

var ifres10793 Obj

if True == tmp10796 {
tmp10794 := GHd(__e, V2022)

tmp10795 := Call1(__e, SymFn(symshen_4wildcard_2), tmp10794)


ifres10793 = tmp10795


} else {
ifres10793 = False


}

ifres10792 = ifres10793


} else {
ifres10792 = False


}

if True == ifres10792 {
tmp10718 := GTl(__e, V2022)

tmp10719 := GTl(__e, V2023)

if SymFn(symshen_4compile_1head) == __self {
__self1, __self2, __self3, __self4, __self5 = V2021, tmp10718, tmp10719, V2024, V2025
__e.Tick()
goto __selftop
}
return TailCall5(__e, SymFn(symshen_4compile_1head), V2021, tmp10718, tmp10719, V2024, V2025)


} else {
tmp10790 := GConsP(__e, V2022)

var ifres10787 Obj

if True == tmp10790 {
tmp10788 := GHd(__e, V2022)

tmp10789 := PrimGuard1(__e, symvariable_2, PrimIsVariable, tmp10788)

ifres10787 = tmp10789


} else {
ifres10787 = False


}

if True == ifres10787 {
return TailCall5(__e, SymFn(symshen_4variable_1case), V2021, V2022, V2023, V2024, V2025)
} else {
tmp10785 := GEq(__e, symshen_4_1m, V2021)

var ifres10780 Obj

if True == tmp10785 {
tmp10784 := GConsP(__e, V2022)

var ifres10781 Obj

if True == tmp10784 {
tmp10782 := GHd(__e, V2022)

tmp10783 := Call1(__e, SymFn(symatom_2), tmp10782)


ifres10781 = tmp10783


} else {
ifres10781 = False


}

ifres10780 = ifres10781


} else {
ifres10780 = False


}

if True == ifres10780 {
return TailCall4(__e, SymFn(symshen_4atom_1case_1minus), V2022, V2023, V2024, V2025)
} else {
tmp10778 := GEq(__e, symshen_4_1m, V2021)

var ifres10754 Obj

if True == tmp10778 {
tmp10777 := GConsP(__e, V2022)

var ifres10755 Obj

if True == tmp10777 {
tmp10775 := GHd(__e, V2022)

tmp10776 := GConsP(__e, tmp10775)

var ifres10756 Obj

if True == tmp10776 {
tmp10772 := GHd(__e, V2022)

tmp10773 := GHd(__e, tmp10772)

tmp10774 := GEq(__e, symcons, tmp10773)

var ifres10757 Obj

if True == tmp10774 {
tmp10769 := GHd(__e, V2022)

tmp10770 := GTl(__e, tmp10769)

tmp10771 := GConsP(__e, tmp10770)

var ifres10758 Obj

if True == tmp10771 {
tmp10765 := GHd(__e, V2022)

tmp10766 := GTl(__e, tmp10765)

tmp10767 := GTl(__e, tmp10766)

tmp10768 := GConsP(__e, tmp10767)

var ifres10759 Obj

if True == tmp10768 {
tmp10760 := GHd(__e, V2022)

tmp10761 := GTl(__e, tmp10760)

tmp10762 := GTl(__e, tmp10761)

tmp10763 := GTl(__e, tmp10762)

tmp10764 := GEq(__e, Nil, tmp10763)

ifres10759 = tmp10764


} else {
ifres10759 = False


}

ifres10758 = ifres10759


} else {
ifres10758 = False


}

ifres10757 = ifres10758


} else {
ifres10757 = False


}

ifres10756 = ifres10757


} else {
ifres10756 = False


}

ifres10755 = ifres10756


} else {
ifres10755 = False


}

ifres10754 = ifres10755


} else {
ifres10754 = False


}

if True == ifres10754 {
return TailCall4(__e, SymFn(symshen_4cons_1case_1minus), V2022, V2023, V2024, V2025)
} else {
tmp10752 := GEq(__e, symshen_4_7m, V2021)

var ifres10747 Obj

if True == tmp10752 {
tmp10751 := GConsP(__e, V2022)

var ifres10748 Obj

if True == tmp10751 {
tmp10749 := GHd(__e, V2022)

tmp10750 := Call1(__e, SymFn(symatom_2), tmp10749)


ifres10748 = tmp10750


} else {
ifres10748 = False


}

ifres10747 = ifres10748


} else {
ifres10747 = False


}

if True == ifres10747 {
return TailCall4(__e, SymFn(symshen_4atom_1case_1plus), V2022, V2023, V2024, V2025)
} else {
tmp10745 := GEq(__e, symshen_4_7m, V2021)

var ifres10721 Obj

if True == tmp10745 {
tmp10744 := GConsP(__e, V2022)

var ifres10722 Obj

if True == tmp10744 {
tmp10742 := GHd(__e, V2022)

tmp10743 := GConsP(__e, tmp10742)

var ifres10723 Obj

if True == tmp10743 {
tmp10739 := GHd(__e, V2022)

tmp10740 := GHd(__e, tmp10739)

tmp10741 := GEq(__e, symcons, tmp10740)

var ifres10724 Obj

if True == tmp10741 {
tmp10736 := GHd(__e, V2022)

tmp10737 := GTl(__e, tmp10736)

tmp10738 := GConsP(__e, tmp10737)

var ifres10725 Obj

if True == tmp10738 {
tmp10732 := GHd(__e, V2022)

tmp10733 := GTl(__e, tmp10732)

tmp10734 := GTl(__e, tmp10733)

tmp10735 := GConsP(__e, tmp10734)

var ifres10726 Obj

if True == tmp10735 {
tmp10727 := GHd(__e, V2022)

tmp10728 := GTl(__e, tmp10727)

tmp10729 := GTl(__e, tmp10728)

tmp10730 := GTl(__e, tmp10729)

tmp10731 := GEq(__e, Nil, tmp10730)

ifres10726 = tmp10731


} else {
ifres10726 = False


}

ifres10725 = ifres10726


} else {
ifres10725 = False


}

ifres10724 = ifres10725


} else {
ifres10724 = False


}

ifres10723 = ifres10724


} else {
ifres10723 = False


}

ifres10722 = ifres10723


} else {
ifres10722 = False


}

ifres10721 = ifres10722


} else {
ifres10721 = False


}

if True == ifres10721 {
return TailCall4(__e, SymFn(symshen_4cons_1case_1plus), V2022, V2023, V2024, V2025)
} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.compile-head"))
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


panic("unreachable")
})
return __self
})()

tmp10848 := Call2(__e, ns2_1set, symshen_4compile_1head, tmp10701)


_ = tmp10848

tmp10849 := MakeNative5(func(__e *ControlFlow, V2036, V2037, V2038, V2039, V2040 Obj) Obj {
tmp10869 := GConsP(__e, V2037)

var ifres10867 Obj

if True == tmp10869 {
tmp10868 := GConsP(__e, V2038)

ifres10867 = tmp10868


} else {
ifres10867 = False


}

if True == ifres10867 {
tmp10864 := GHd(__e, V2038)

tmp10865 := PrimGuard1(__e, symvariable_2, PrimIsVariable, tmp10864)

if True == tmp10865 {
tmp10850 := GTl(__e, V2037)

tmp10851 := GTl(__e, V2038)

tmp10852 := GHd(__e, V2038)

tmp10853 := GHd(__e, V2037)

tmp10854 := Call3(__e, SymFn(symsubst), tmp10852, tmp10853, V2040)


return TailCall5(__e, SymFn(symshen_4compile_1head), V2036, tmp10850, tmp10851, V2039, tmp10854)


} else {
tmp10855 := GHd(__e, V2037)

tmp10856 := GHd(__e, V2038)

tmp10857 := GTl(__e, V2037)

tmp10858 := GTl(__e, V2038)

tmp10859 := Call5(__e, SymFn(symshen_4compile_1head), V2036, tmp10857, tmp10858, V2039, V2040)


tmp10860 := GCons(__e, tmp10859, Nil)

tmp10861 := GCons(__e, tmp10856, tmp10860)

tmp10862 := GCons(__e, tmp10855, tmp10861)

return GCons(__e, symlet, tmp10862)


}


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.variable-case"))
}


panic("unreachable")
})

tmp10870 := Call2(__e, ns2_1set, symshen_4variable_1case, tmp10849)


_ = tmp10870

tmp10871 := MakeNative4(func(__e *ControlFlow, V2049, V2050, V2051, V2052 Obj) Obj {
tmp10894 := GConsP(__e, V2049)

var ifres10892 Obj

if True == tmp10894 {
tmp10893 := GConsP(__e, V2050)

ifres10892 = tmp10893


} else {
ifres10892 = False


}

if True == ifres10892 {
tmp10872 := Call1(__e, SymFn(symgensym), symTm)


let__9513 := tmp10872
_ = let__9513

tmp10873 := GHd(__e, V2050)

tmp10874 := GCons(__e, V2051, Nil)

tmp10875 := GCons(__e, tmp10873, tmp10874)

tmp10876 := GCons(__e, symshen_4lazyderef, tmp10875)

tmp10877 := GHd(__e, V2049)

tmp10878 := GCons(__e, tmp10877, Nil)

tmp10879 := GCons(__e, let__9513, tmp10878)

tmp10880 := GCons(__e, sym_a, tmp10879)

tmp10881 := GTl(__e, V2049)

tmp10882 := GTl(__e, V2050)

tmp10883 := Call5(__e, SymFn(symshen_4compile_1head), symshen_4_1m, tmp10881, tmp10882, V2051, V2052)


tmp10884 := GCons(__e, False, Nil)

tmp10885 := GCons(__e, tmp10883, tmp10884)

tmp10886 := GCons(__e, tmp10880, tmp10885)

tmp10887 := GCons(__e, symif, tmp10886)

tmp10888 := GCons(__e, tmp10887, Nil)

tmp10889 := GCons(__e, tmp10876, tmp10888)

tmp10890 := GCons(__e, let__9513, tmp10889)

return GCons(__e, symlet, tmp10890)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.atom-case-minus"))
}


panic("unreachable")
})

tmp10895 := Call2(__e, ns2_1set, symshen_4atom_1case_1minus, tmp10871)


_ = tmp10895

tmp10896 := MakeNative4(func(__e *ControlFlow, V2062, V2063, V2064, V2065 Obj) Obj {
tmp10954 := GConsP(__e, V2062)

var ifres10930 Obj

if True == tmp10954 {
tmp10952 := GHd(__e, V2062)

tmp10953 := GConsP(__e, tmp10952)

var ifres10931 Obj

if True == tmp10953 {
tmp10949 := GHd(__e, V2062)

tmp10950 := GHd(__e, tmp10949)

tmp10951 := GEq(__e, symcons, tmp10950)

var ifres10932 Obj

if True == tmp10951 {
tmp10946 := GHd(__e, V2062)

tmp10947 := GTl(__e, tmp10946)

tmp10948 := GConsP(__e, tmp10947)

var ifres10933 Obj

if True == tmp10948 {
tmp10942 := GHd(__e, V2062)

tmp10943 := GTl(__e, tmp10942)

tmp10944 := GTl(__e, tmp10943)

tmp10945 := GConsP(__e, tmp10944)

var ifres10934 Obj

if True == tmp10945 {
tmp10937 := GHd(__e, V2062)

tmp10938 := GTl(__e, tmp10937)

tmp10939 := GTl(__e, tmp10938)

tmp10940 := GTl(__e, tmp10939)

tmp10941 := GEq(__e, Nil, tmp10940)

var ifres10935 Obj

if True == tmp10941 {
tmp10936 := GConsP(__e, V2063)

ifres10935 = tmp10936


} else {
ifres10935 = False


}

ifres10934 = ifres10935


} else {
ifres10934 = False


}

ifres10933 = ifres10934


} else {
ifres10933 = False


}

ifres10932 = ifres10933


} else {
ifres10932 = False


}

ifres10931 = ifres10932


} else {
ifres10931 = False


}

ifres10930 = ifres10931


} else {
ifres10930 = False


}

if True == ifres10930 {
tmp10897 := Call1(__e, SymFn(symgensym), symTm)


let__9514 := tmp10897
_ = let__9514

tmp10898 := GHd(__e, V2063)

tmp10899 := GCons(__e, V2064, Nil)

tmp10900 := GCons(__e, tmp10898, tmp10899)

tmp10901 := GCons(__e, symshen_4lazyderef, tmp10900)

tmp10902 := GCons(__e, let__9514, Nil)

tmp10903 := GCons(__e, symcons_2, tmp10902)

tmp10904 := GHd(__e, V2062)

tmp10905 := GTl(__e, tmp10904)

tmp10906 := GHd(__e, tmp10905)

tmp10907 := GHd(__e, V2062)

tmp10908 := GTl(__e, tmp10907)

tmp10909 := GTl(__e, tmp10908)

tmp10910 := GHd(__e, tmp10909)

tmp10911 := GTl(__e, V2062)

tmp10912 := GCons(__e, tmp10910, tmp10911)

tmp10913 := GCons(__e, tmp10906, tmp10912)

tmp10914 := GCons(__e, let__9514, Nil)

tmp10915 := GCons(__e, symhd, tmp10914)

tmp10916 := GCons(__e, let__9514, Nil)

tmp10917 := GCons(__e, symtl, tmp10916)

tmp10918 := GTl(__e, V2063)

tmp10919 := GCons(__e, tmp10917, tmp10918)

tmp10920 := GCons(__e, tmp10915, tmp10919)

tmp10921 := Call5(__e, SymFn(symshen_4compile_1head), symshen_4_1m, tmp10913, tmp10920, V2064, V2065)


tmp10922 := GCons(__e, False, Nil)

tmp10923 := GCons(__e, tmp10921, tmp10922)

tmp10924 := GCons(__e, tmp10903, tmp10923)

tmp10925 := GCons(__e, symif, tmp10924)

tmp10926 := GCons(__e, tmp10925, Nil)

tmp10927 := GCons(__e, tmp10901, tmp10926)

tmp10928 := GCons(__e, let__9514, tmp10927)

return GCons(__e, symlet, tmp10928)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.cons-case-minus"))
}


panic("unreachable")
})

tmp10955 := Call2(__e, ns2_1set, symshen_4cons_1case_1minus, tmp10896)


_ = tmp10955

tmp10956 := MakeNative4(func(__e *ControlFlow, V2075, V2076, V2077, V2078 Obj) Obj {
tmp10999 := GConsP(__e, V2075)

var ifres10997 Obj

if True == tmp10999 {
tmp10998 := GConsP(__e, V2076)

ifres10997 = tmp10998


} else {
ifres10997 = False


}

if True == ifres10997 {
tmp10957 := Call1(__e, SymFn(symgensym), symTm)


let__9515 := tmp10957
_ = let__9515

tmp10958 := Call1(__e, SymFn(symgensym), symGoTo)


let__9516 := tmp10958
_ = let__9516

tmp10959 := GHd(__e, V2076)

tmp10960 := GCons(__e, V2077, Nil)

tmp10961 := GCons(__e, tmp10959, tmp10960)

tmp10962 := GCons(__e, symshen_4lazyderef, tmp10961)

tmp10963 := GTl(__e, V2075)

tmp10964 := GTl(__e, V2076)

tmp10965 := Call5(__e, SymFn(symshen_4compile_1head), symshen_4_7m, tmp10963, tmp10964, V2077, V2078)


tmp10966 := GCons(__e, tmp10965, Nil)

tmp10967 := GCons(__e, symfreeze, tmp10966)

tmp10968 := GHd(__e, V2075)

tmp10969 := GCons(__e, tmp10968, Nil)

tmp10970 := GCons(__e, let__9515, tmp10969)

tmp10971 := GCons(__e, sym_a, tmp10970)

tmp10972 := GCons(__e, let__9516, Nil)

tmp10973 := GCons(__e, symthaw, tmp10972)

tmp10974 := GCons(__e, let__9515, Nil)

tmp10975 := GCons(__e, symshen_4pvar_2, tmp10974)

tmp10976 := GHd(__e, V2075)

tmp10977 := Call1(__e, SymFn(symshen_4demode), tmp10976)


tmp10978 := GCons(__e, let__9516, Nil)

tmp10979 := GCons(__e, V2077, tmp10978)

tmp10980 := GCons(__e, tmp10977, tmp10979)

tmp10981 := GCons(__e, let__9515, tmp10980)

tmp10982 := GCons(__e, symshen_4bind_b, tmp10981)

tmp10983 := GCons(__e, False, Nil)

tmp10984 := GCons(__e, tmp10982, tmp10983)

tmp10985 := GCons(__e, tmp10975, tmp10984)

tmp10986 := GCons(__e, symif, tmp10985)

tmp10987 := GCons(__e, tmp10986, Nil)

tmp10988 := GCons(__e, tmp10973, tmp10987)

tmp10989 := GCons(__e, tmp10971, tmp10988)

tmp10990 := GCons(__e, symif, tmp10989)

tmp10991 := GCons(__e, tmp10990, Nil)

tmp10992 := GCons(__e, tmp10967, tmp10991)

tmp10993 := GCons(__e, let__9516, tmp10992)

tmp10994 := GCons(__e, tmp10962, tmp10993)

tmp10995 := GCons(__e, let__9515, tmp10994)

return GCons(__e, symlet, tmp10995)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.atom-case-plus"))
}


panic("unreachable")
})

tmp11000 := Call2(__e, ns2_1set, symshen_4atom_1case_1plus, tmp10956)


_ = tmp11000

tmp11001 := MakeNative4(func(__e *ControlFlow, V2089, V2090, V2091, V2092 Obj) Obj {
tmp11086 := GConsP(__e, V2089)

var ifres11062 Obj

if True == tmp11086 {
tmp11084 := GHd(__e, V2089)

tmp11085 := GConsP(__e, tmp11084)

var ifres11063 Obj

if True == tmp11085 {
tmp11081 := GHd(__e, V2089)

tmp11082 := GHd(__e, tmp11081)

tmp11083 := GEq(__e, symcons, tmp11082)

var ifres11064 Obj

if True == tmp11083 {
tmp11078 := GHd(__e, V2089)

tmp11079 := GTl(__e, tmp11078)

tmp11080 := GConsP(__e, tmp11079)

var ifres11065 Obj

if True == tmp11080 {
tmp11074 := GHd(__e, V2089)

tmp11075 := GTl(__e, tmp11074)

tmp11076 := GTl(__e, tmp11075)

tmp11077 := GConsP(__e, tmp11076)

var ifres11066 Obj

if True == tmp11077 {
tmp11069 := GHd(__e, V2089)

tmp11070 := GTl(__e, tmp11069)

tmp11071 := GTl(__e, tmp11070)

tmp11072 := GTl(__e, tmp11071)

tmp11073 := GEq(__e, Nil, tmp11072)

var ifres11067 Obj

if True == tmp11073 {
tmp11068 := GConsP(__e, V2090)

ifres11067 = tmp11068


} else {
ifres11067 = False


}

ifres11066 = ifres11067


} else {
ifres11066 = False


}

ifres11065 = ifres11066


} else {
ifres11065 = False


}

ifres11064 = ifres11065


} else {
ifres11064 = False


}

ifres11063 = ifres11064


} else {
ifres11063 = False


}

ifres11062 = ifres11063


} else {
ifres11062 = False


}

if True == ifres11062 {
tmp11002 := Call1(__e, SymFn(symgensym), symTm)


let__9517 := tmp11002
_ = let__9517

tmp11003 := Call1(__e, SymFn(symgensym), symGoTo)


let__9518 := tmp11003
_ = let__9518

tmp11004 := GHd(__e, V2089)

tmp11005 := GTl(__e, tmp11004)

tmp11006 := GHd(__e, tmp11005)

tmp11007 := GHd(__e, V2089)

tmp11008 := GTl(__e, tmp11007)

tmp11009 := GTl(__e, tmp11008)

tmp11010 := GHd(__e, tmp11009)

tmp11011 := GCons(__e, tmp11006, tmp11010)

tmp11012 := Call1(__e, SymFn(symshen_4extract_1vars), tmp11011)


let__9519 := tmp11012
_ = let__9519

tmp11013 := GHd(__e, V2089)

tmp11014 := Call1(__e, SymFn(symshen_4tame), tmp11013)


let__9520 := tmp11014
_ = let__9520

tmp11015 := Call1(__e, SymFn(symshen_4extract_1vars), let__9520)


let__9521 := tmp11015
_ = let__9521

tmp11016 := GHd(__e, V2090)

tmp11017 := GCons(__e, V2091, Nil)

tmp11018 := GCons(__e, tmp11016, tmp11017)

tmp11019 := GCons(__e, symshen_4lazyderef, tmp11018)

tmp11020 := GTl(__e, V2089)

tmp11021 := GTl(__e, V2090)

tmp11022 := Call5(__e, SymFn(symshen_4compile_1head), symshen_4_7m, tmp11020, tmp11021, V2091, V2092)


tmp11023 := Call2(__e, SymFn(symshen_4goto), let__9519, tmp11022)


tmp11024 := GCons(__e, let__9517, Nil)

tmp11025 := GCons(__e, symcons_2, tmp11024)

tmp11026 := GHd(__e, V2089)

tmp11027 := GTl(__e, tmp11026)

tmp11028 := GCons(__e, let__9517, Nil)

tmp11029 := GCons(__e, symhd, tmp11028)

tmp11030 := GCons(__e, let__9517, Nil)

tmp11031 := GCons(__e, symtl, tmp11030)

tmp11032 := GCons(__e, tmp11031, Nil)

tmp11033 := GCons(__e, tmp11029, tmp11032)

tmp11034 := Call2(__e, SymFn(symshen_4invoke), let__9518, let__9519)


tmp11035 := Call5(__e, SymFn(symshen_4compile_1head), symshen_4_7m, tmp11027, tmp11033, V2091, tmp11034)


tmp11036 := GCons(__e, let__9517, Nil)

tmp11037 := GCons(__e, symshen_4pvar_2, tmp11036)

tmp11038 := Call1(__e, SymFn(symshen_4demode), let__9520)


tmp11039 := Call2(__e, SymFn(symshen_4invoke), let__9518, let__9519)


tmp11040 := GCons(__e, tmp11039, Nil)

tmp11041 := GCons(__e, symfreeze, tmp11040)

tmp11042 := GCons(__e, tmp11041, Nil)

tmp11043 := GCons(__e, V2091, tmp11042)

tmp11044 := GCons(__e, tmp11038, tmp11043)

tmp11045 := GCons(__e, let__9517, tmp11044)

tmp11046 := GCons(__e, symshen_4bind_b, tmp11045)

tmp11047 := Call3(__e, SymFn(symshen_4stpart), let__9521, tmp11046, V2091)


tmp11048 := GCons(__e, False, Nil)

tmp11049 := GCons(__e, tmp11047, tmp11048)

tmp11050 := GCons(__e, tmp11037, tmp11049)

tmp11051 := GCons(__e, symif, tmp11050)

tmp11052 := GCons(__e, tmp11051, Nil)

tmp11053 := GCons(__e, tmp11035, tmp11052)

tmp11054 := GCons(__e, tmp11025, tmp11053)

tmp11055 := GCons(__e, symif, tmp11054)

tmp11056 := GCons(__e, tmp11055, Nil)

tmp11057 := GCons(__e, tmp11023, tmp11056)

tmp11058 := GCons(__e, let__9518, tmp11057)

tmp11059 := GCons(__e, tmp11019, tmp11058)

tmp11060 := GCons(__e, let__9517, tmp11059)

return GCons(__e, symlet, tmp11060)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.cons-case-plus"))
}


panic("unreachable")
})

tmp11087 := Call2(__e, ns2_1set, symshen_4cons_1case_1plus, tmp11001)


_ = tmp11087

tmp11088 := (func() Obj {
var __self Obj
__self = MakeNative1(func(__e *ControlFlow, __self1 Obj) Obj {
__selftop:
V2098 := __self1
_ = V2098
tmp11119 := GConsP(__e, V2098)

var ifres11109 Obj

if True == tmp11119 {
tmp11117 := GHd(__e, V2098)

tmp11118 := GEq(__e, symshen_4_7m, tmp11117)

var ifres11110 Obj

if True == tmp11118 {
tmp11115 := GTl(__e, V2098)

tmp11116 := GConsP(__e, tmp11115)

var ifres11111 Obj

if True == tmp11116 {
tmp11112 := GTl(__e, V2098)

tmp11113 := GTl(__e, tmp11112)

tmp11114 := GEq(__e, Nil, tmp11113)

ifres11111 = tmp11114


} else {
ifres11111 = False


}

ifres11110 = ifres11111


} else {
ifres11110 = False


}

ifres11109 = ifres11110


} else {
ifres11109 = False


}

if True == ifres11109 {
tmp11089 := GTl(__e, V2098)

tmp11090 := GHd(__e, tmp11089)

if SymFn(symshen_4demode) == __self {
__self1 = tmp11090
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4demode), tmp11090)


} else {
tmp11107 := GConsP(__e, V2098)

var ifres11097 Obj

if True == tmp11107 {
tmp11105 := GHd(__e, V2098)

tmp11106 := GEq(__e, symshen_4_1m, tmp11105)

var ifres11098 Obj

if True == tmp11106 {
tmp11103 := GTl(__e, V2098)

tmp11104 := GConsP(__e, tmp11103)

var ifres11099 Obj

if True == tmp11104 {
tmp11100 := GTl(__e, V2098)

tmp11101 := GTl(__e, tmp11100)

tmp11102 := GEq(__e, Nil, tmp11101)

ifres11099 = tmp11102


} else {
ifres11099 = False


}

ifres11098 = ifres11099


} else {
ifres11098 = False


}

ifres11097 = ifres11098


} else {
ifres11097 = False


}

if True == ifres11097 {
tmp11091 := GTl(__e, V2098)

tmp11092 := GHd(__e, tmp11091)

if SymFn(symshen_4demode) == __self {
__self1 = tmp11092
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4demode), tmp11092)


} else {
tmp11095 := GConsP(__e, V2098)

if True == tmp11095 {
tmp11093 := MakeNative1(func(__e *ControlFlow, Z2099 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4demode), Z2099)
panic("unreachable")
})

return TailCall2(__e, SymFn(symmap), tmp11093, V2098)


} else {
return V2098
}


}


}


panic("unreachable")
})
return __self
})()

tmp11120 := Call2(__e, ns2_1set, symshen_4demode, tmp11088)


_ = tmp11120

tmp11121 := MakeNative1(func(__e *ControlFlow, V2100 Obj) Obj {
tmp11126 := Call1(__e, SymFn(symshen_4wildcard_2), V2100)


if True == tmp11126 {
return TailCall1(__e, SymFn(symgensym), symY)
} else {
tmp11124 := GConsP(__e, V2100)

if True == tmp11124 {
tmp11122 := MakeNative1(func(__e *ControlFlow, Z2101 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4tame), Z2101)
panic("unreachable")
})

return TailCall2(__e, SymFn(symmap), tmp11122, V2100)


} else {
return V2100
}


}


panic("unreachable")
})

tmp11127 := Call2(__e, ns2_1set, symshen_4tame, tmp11121)


_ = tmp11127

tmp11128 := MakeNative2(func(__e *ControlFlow, V2102, V2103 Obj) Obj {
tmp11131 := GEq(__e, Nil, V2102)

if True == tmp11131 {
tmp11129 := GCons(__e, V2103, Nil)

return GCons(__e, symfreeze, tmp11129)


} else {
return TailCall2(__e, SymFn(symshen_4goto_1h), V2102, V2103)
}


panic("unreachable")
})

tmp11132 := Call2(__e, ns2_1set, symshen_4goto, tmp11128)


_ = tmp11132

tmp11133 := MakeNative2(func(__e *ControlFlow, V2104, V2105 Obj) Obj {
tmp11142 := GEq(__e, Nil, V2104)

if True == tmp11142 {
return V2105
} else {
tmp11140 := GConsP(__e, V2104)

if True == tmp11140 {
tmp11134 := GHd(__e, V2104)

tmp11135 := GTl(__e, V2104)

tmp11136 := Call2(__e, SymFn(symshen_4goto_1h), tmp11135, V2105)


tmp11137 := GCons(__e, tmp11136, Nil)

tmp11138 := GCons(__e, tmp11134, tmp11137)

return GCons(__e, symlambda, tmp11138)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4goto_1h)
}


}


panic("unreachable")
})

tmp11143 := Call2(__e, ns2_1set, symshen_4goto_1h, tmp11133)


_ = tmp11143

tmp11144 := MakeNative2(func(__e *ControlFlow, V2106, V2107 Obj) Obj {
tmp11147 := GEq(__e, Nil, V2107)

if True == tmp11147 {
tmp11145 := GCons(__e, V2106, Nil)

return GCons(__e, symthaw, tmp11145)


} else {
return GCons(__e, V2106, V2107)
}


panic("unreachable")
})

tmp11148 := Call2(__e, ns2_1set, symshen_4invoke, tmp11144)


_ = tmp11148

tmp11149 := MakeNative1(func(__e *ControlFlow, V2108 Obj) Obj {
return GEq(__e, V2108, sym__)
panic("unreachable")
})

tmp11150 := Call2(__e, ns2_1set, symshen_4wildcard_2, tmp11149)


_ = tmp11150

tmp11151 := MakeNative1(func(__e *ControlFlow, V2109 Obj) Obj {
tmp11156 := GVectorP(__e, V2109)

if True == tmp11156 {
tmp11152 := MakeNative0(func(__e *ControlFlow) Obj {
return GVecGet(__e, V2109, MakeInteger(0))
panic("unreachable")
})

tmp11153 := MakeNative1(func(__e *ControlFlow, Z2110 Obj) Obj {
return symshen_4not_1pvar
panic("unreachable")
})

tmp11154 := Call2(__e, try_1catch, tmp11152, tmp11153)


return GEq(__e, tmp11154, symshen_4pvar)


} else {
return False
}


panic("unreachable")
})

tmp11157 := Call2(__e, ns2_1set, symshen_4pvar_2, tmp11151)


_ = tmp11157

tmp11158 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V2111 := __self1
_ = V2111
V2112 := __self2
_ = V2112
tmp11164 := Call1(__e, SymFn(symshen_4pvar_2), V2111)


if True == tmp11164 {
tmp11159 := GVecGet(__e, V2111, MakeInteger(1))

tmp11160 := GVecGet(__e, V2112, tmp11159)

let__9522 := tmp11160
_ = let__9522

tmp11162 := GEq(__e, let__9522, symshen_4_1null_1)

if True == tmp11162 {
return V2111
} else {
if SymFn(symshen_4lazyderef) == __self {
__self1, __self2 = let__9522, V2112
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4lazyderef), let__9522, V2112)
}


} else {
return V2111
}


panic("unreachable")
})
return __self
})()

tmp11165 := Call2(__e, ns2_1set, symshen_4lazyderef, tmp11158)


_ = tmp11165

tmp11166 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V2114 := __self1
_ = V2114
V2115 := __self2
_ = V2115
tmp11178 := GConsP(__e, V2114)

if True == tmp11178 {
tmp11167 := GHd(__e, V2114)

tmp11168 := Call2(__e, SymFn(symshen_4deref), tmp11167, V2115)


tmp11169 := GTl(__e, V2114)

tmp11170 := Call2(__e, SymFn(symshen_4deref), tmp11169, V2115)


return GCons(__e, tmp11168, tmp11170)


} else {
tmp11176 := Call1(__e, SymFn(symshen_4pvar_2), V2114)


if True == tmp11176 {
tmp11171 := GVecGet(__e, V2114, MakeInteger(1))

tmp11172 := GVecGet(__e, V2115, tmp11171)

let__9523 := tmp11172
_ = let__9523

tmp11174 := GEq(__e, let__9523, symshen_4_1null_1)

if True == tmp11174 {
return V2114
} else {
if SymFn(symshen_4deref) == __self {
__self1, __self2 = let__9523, V2115
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4deref), let__9523, V2115)
}


} else {
return V2114
}


}


panic("unreachable")
})
return __self
})()

tmp11179 := Call2(__e, ns2_1set, symshen_4deref, tmp11166)


_ = tmp11179

tmp11180 := MakeNative4(func(__e *ControlFlow, V2117, V2118, V2119, V2120 Obj) Obj {
tmp11181 := Call3(__e, SymFn(symshen_4bindv), V2117, V2118, V2119)


let__9524 := tmp11181
_ = let__9524

tmp11182 := Call1(__e, SymFn(symthaw), V2120)


let__9525 := tmp11182
_ = let__9525

tmp11184 := GEq(__e, let__9525, False)

if True == tmp11184 {
return TailCall3(__e, SymFn(symshen_4unwind), V2117, V2119, let__9525)
} else {
return let__9525
}


panic("unreachable")
})

tmp11185 := Call2(__e, ns2_1set, symshen_4bind_b, tmp11180)


_ = tmp11185

tmp11186 := MakeNative3(func(__e *ControlFlow, V2123, V2124, V2125 Obj) Obj {
tmp11187 := GVecGet(__e, V2123, MakeInteger(1))

return GVecSet(__e, V2125, tmp11187, V2124)


panic("unreachable")
})

tmp11188 := Call2(__e, ns2_1set, symshen_4bindv, tmp11186)


_ = tmp11188

tmp11189 := MakeNative3(func(__e *ControlFlow, V2126, V2127, V2128 Obj) Obj {
tmp11190 := GVecGet(__e, V2126, MakeInteger(1))

tmp11191 := GVecSet(__e, V2127, tmp11190, symshen_4_1null_1)

_ = tmp11191

return V2128


panic("unreachable")
})

tmp11192 := Call2(__e, ns2_1set, symshen_4unwind, tmp11189)


_ = tmp11192

tmp11193 := MakeNative3(func(__e *ControlFlow, V2137, V2138, V2139 Obj) Obj {
tmp11208 := GEq(__e, Nil, V2137)

if True == tmp11208 {
return V2138
} else {
tmp11206 := GConsP(__e, V2137)

if True == tmp11206 {
tmp11194 := GHd(__e, V2137)

tmp11195 := GCons(__e, V2139, Nil)

tmp11196 := GCons(__e, symshen_4newpv, tmp11195)

tmp11197 := GTl(__e, V2137)

tmp11198 := Call3(__e, SymFn(symshen_4stpart), tmp11197, V2138, V2139)


tmp11199 := GCons(__e, tmp11198, Nil)

tmp11200 := GCons(__e, V2139, tmp11199)

tmp11201 := GCons(__e, symshen_4gc, tmp11200)

tmp11202 := GCons(__e, tmp11201, Nil)

tmp11203 := GCons(__e, tmp11196, tmp11202)

tmp11204 := GCons(__e, tmp11194, tmp11203)

return GCons(__e, symlet, tmp11204)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.stpart"))
}


}


panic("unreachable")
})

tmp11209 := Call2(__e, ns2_1set, symshen_4stpart, tmp11193)


_ = tmp11209

tmp11210 := MakeNative2(func(__e *ControlFlow, V2140, V2141 Obj) Obj {
tmp11214 := GEq(__e, V2141, False)

if True == tmp11214 {
tmp11211 := Call1(__e, SymFn(symshen_4ticket_1number), V2140)


let__9526 := tmp11211
_ = let__9526

tmp11212 := Call2(__e, SymFn(symshen_4decrement_1ticket), let__9526, V2140)


_ = tmp11212

return V2141


} else {
return V2141
}


panic("unreachable")
})

tmp11215 := Call2(__e, ns2_1set, symshen_4gc, tmp11210)


_ = tmp11215

tmp11216 := MakeNative2(func(__e *ControlFlow, V2143, V2144 Obj) Obj {
return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(V2144, MakeInteger(1), GSub(__e, V2143, MakeInteger(1)))
}
__typedArg0 := V2144
__typedArg1 := MakeInteger(1)
__typedArg2 := GSub(__e, V2143, MakeInteger(1))
return Call3(__e, SymFn(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()


panic("unreachable")
})

tmp11218 := Call2(__e, ns2_1set, symshen_4decrement_1ticket, tmp11216)


_ = tmp11218

tmp11219 := MakeNative1(func(__e *ControlFlow, V2145 Obj) Obj {
tmp11220 := Call1(__e, SymFn(symshen_4ticket_1number), V2145)


let__9527 := tmp11220
_ = let__9527

tmp11221 := Call1(__e, SymFn(symshen_4make_1prolog_1variable), let__9527)


let__9528 := tmp11221
_ = let__9528

tmp11222 := Call2(__e, SymFn(symshen_4nextticket), V2145, let__9527)


let__9529 := tmp11222
_ = let__9529

return let__9528


panic("unreachable")
})

tmp11223 := Call2(__e, ns2_1set, symshen_4newpv, tmp11219)


_ = tmp11223

tmp11224 := MakeNative1(func(__e *ControlFlow, V2149 Obj) Obj {
return GVecGet(__e, V2149, MakeInteger(1))
panic("unreachable")
})

tmp11225 := Call2(__e, ns2_1set, symshen_4ticket_1number, tmp11224)


_ = tmp11225

tmp11226 := MakeNative2(func(__e *ControlFlow, V2150, V2151 Obj) Obj {
tmp11227 := GVecSet(__e, V2150, V2151, symshen_4_1null_1)

let__9530 := tmp11227
_ = let__9530

return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(let__9530, MakeInteger(1), GAdd(__e, V2151, MakeInteger(1)))
}
__typedArg0 := let__9530
__typedArg1 := MakeInteger(1)
__typedArg2 := GAdd(__e, V2151, MakeInteger(1))
return Call3(__e, SymFn(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()


panic("unreachable")
})

tmp11229 := Call2(__e, ns2_1set, symshen_4nextticket, tmp11226)


_ = tmp11229

tmp11230 := MakeNative1(func(__e *ControlFlow, V2153 Obj) Obj {
tmp11231 := PrimGuard1(__e, symabsvector, PrimAbsvector, MakeInteger(2))

tmp11232 := GVecSet(__e, tmp11231, MakeInteger(0), symshen_4pvar)

return GVecSet(__e, tmp11232, MakeInteger(1), V2153)


panic("unreachable")
})

tmp11233 := Call2(__e, ns2_1set, symshen_4make_1prolog_1variable, tmp11230)


_ = tmp11233

tmp11234 := MakeNative1(func(__e *ControlFlow, V2154 Obj) Obj {
tmp11235 := GVecGet(__e, V2154, MakeInteger(1))

tmp11236 := Call3(__e, SymFn(symshen_4app), tmp11235, MakeString(""), symshen_4a)


return PrimGuard2(__e, symcn, PrimStringConcat, MakeString("Var"), tmp11236)


panic("unreachable")
})

tmp11237 := Call2(__e, ns2_1set, symshen_4pvar, tmp11234)


_ = tmp11237

tmp11238 := MakeNative0(func(__e *ControlFlow) Obj {
tmp11239 := GValue(__e, symshen_4_dinfs_d)

return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dinfs_d, GAdd(__e, MakeInteger(1), tmp11239))
}
__typedArg0 := symshen_4_dinfs_d
__typedArg1 := GAdd(__e, MakeInteger(1), tmp11239)
return Call2(__e, SymFn(symset), __typedArg0, __typedArg1)
})()


panic("unreachable")
})

tmp11241 := Call2(__e, ns2_1set, symshen_4incinfs, tmp11238)


_ = tmp11241

tmp11242 := MakeNative1(func(__e *ControlFlow, V2155 Obj) Obj {
tmp11248 := PrimGuard1(__e, syminteger_2, PrimIsInteger, V2155)

var ifres11246 Obj

if True == tmp11248 {
ifres11246 = GGt(__e, V2155, MakeInteger(0))


} else {
ifres11246 = False


}

if True == ifres11246 {
return PrimGuard2(__e, symset, PrimSet, symshen_4_dsize_1prolog_1vector_d, V2155)
} else {
tmp11243 := Call3(__e, SymFn(symshen_4app), V2155, MakeString(""), symshen_4a)


return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("prolog vector size: size should be a positive integer; not "), tmp11243))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("prolog vector size: size should be a positive integer; not "), tmp11243)
return Call1(__e, SymFn(symsimple_1error), __typedArg0)
})()


}


panic("unreachable")
})

tmp11249 := Call2(__e, ns2_1set, symshen_4prolog_1vector_1size, tmp11242)


_ = tmp11249

tmp11250 := (func() Obj {
var __self Obj
__self = MakeNative4(func(__e *ControlFlow, __self1, __self2, __self3, __self4 Obj) Obj {
__selftop:
V2167 := __self1
_ = V2167
V2168 := __self2
_ = V2168
V2169 := __self3
_ = V2169
V2170 := __self4
_ = V2170
tmp11277 := GEq(__e, V2167, V2168)

if True == tmp11277 {
return TailCall1(__e, SymFn(symthaw), V2170)
} else {
tmp11275 := Call1(__e, SymFn(symshen_4pvar_2), V2167)


var ifres11271 Obj

if True == tmp11275 {
tmp11272 := Call2(__e, SymFn(symshen_4deref), V2168, V2169)


tmp11273 := Call2(__e, SymFn(symshen_4occurs_1check_2), V2167, tmp11272)


ifres11271 = GNot(__e, tmp11273)


} else {
ifres11271 = False


}

if True == ifres11271 {
return TailCall4(__e, SymFn(symshen_4bind_b), V2167, V2168, V2169, V2170)
} else {
tmp11269 := Call1(__e, SymFn(symshen_4pvar_2), V2168)


var ifres11265 Obj

if True == tmp11269 {
tmp11266 := Call2(__e, SymFn(symshen_4deref), V2167, V2169)


tmp11267 := Call2(__e, SymFn(symshen_4occurs_1check_2), V2168, tmp11266)


ifres11265 = GNot(__e, tmp11267)


} else {
ifres11265 = False


}

if True == ifres11265 {
return TailCall4(__e, SymFn(symshen_4bind_b), V2168, V2167, V2169, V2170)
} else {
tmp11263 := GConsP(__e, V2167)

var ifres11261 Obj

if True == tmp11263 {
tmp11262 := GConsP(__e, V2168)

ifres11261 = tmp11262


} else {
ifres11261 = False


}

if True == ifres11261 {
tmp11251 := GHd(__e, V2167)

tmp11252 := Call2(__e, SymFn(symshen_4lazyderef), tmp11251, V2169)


tmp11253 := GHd(__e, V2168)

tmp11254 := Call2(__e, SymFn(symshen_4lazyderef), tmp11253, V2169)


tmp11255 := MakeNative0(func(__e *ControlFlow) Obj {
tmp11256 := GTl(__e, V2167)

tmp11257 := Call2(__e, SymFn(symshen_4lazyderef), tmp11256, V2169)


tmp11258 := GTl(__e, V2168)

tmp11259 := Call2(__e, SymFn(symshen_4lazyderef), tmp11258, V2169)


return TailCall4(__e, SymFn(symshen_4lzy_a_b), tmp11257, tmp11259, V2169, V2170)


panic("unreachable")
})

if SymFn(symshen_4lzy_a_b) == __self {
__self1, __self2, __self3, __self4 = tmp11252, tmp11254, V2169, tmp11255
__e.Tick()
goto __selftop
}
return TailCall4(__e, SymFn(symshen_4lzy_a_b), tmp11252, tmp11254, V2169, tmp11255)


} else {
return False
}


}


}


}


panic("unreachable")
})
return __self
})()

tmp11278 := Call2(__e, ns2_1set, symshen_4lzy_a_b, tmp11250)


_ = tmp11278

tmp11279 := (func() Obj {
var __self Obj
__self = MakeNative4(func(__e *ControlFlow, __self1, __self2, __self3, __self4 Obj) Obj {
__selftop:
V2182 := __self1
_ = V2182
V2183 := __self2
_ = V2183
V2184 := __self3
_ = V2184
V2185 := __self4
_ = V2185
tmp11298 := GEq(__e, V2182, V2183)

if True == tmp11298 {
return TailCall1(__e, SymFn(symthaw), V2185)
} else {
tmp11296 := Call1(__e, SymFn(symshen_4pvar_2), V2182)


if True == tmp11296 {
return TailCall4(__e, SymFn(symshen_4bind_b), V2182, V2183, V2184, V2185)
} else {
tmp11294 := Call1(__e, SymFn(symshen_4pvar_2), V2183)


if True == tmp11294 {
return TailCall4(__e, SymFn(symshen_4bind_b), V2183, V2182, V2184, V2185)
} else {
tmp11292 := GConsP(__e, V2182)

var ifres11290 Obj

if True == tmp11292 {
tmp11291 := GConsP(__e, V2183)

ifres11290 = tmp11291


} else {
ifres11290 = False


}

if True == ifres11290 {
tmp11280 := GHd(__e, V2182)

tmp11281 := Call2(__e, SymFn(symshen_4lazyderef), tmp11280, V2184)


tmp11282 := GHd(__e, V2183)

tmp11283 := Call2(__e, SymFn(symshen_4lazyderef), tmp11282, V2184)


tmp11284 := MakeNative0(func(__e *ControlFlow) Obj {
tmp11285 := GTl(__e, V2182)

tmp11286 := Call2(__e, SymFn(symshen_4lazyderef), tmp11285, V2184)


tmp11287 := GTl(__e, V2183)

tmp11288 := Call2(__e, SymFn(symshen_4lazyderef), tmp11287, V2184)


return TailCall4(__e, SymFn(symshen_4lzy_a), tmp11286, tmp11288, V2184, V2185)


panic("unreachable")
})

if SymFn(symshen_4lzy_a) == __self {
__self1, __self2, __self3, __self4 = tmp11281, tmp11283, V2184, tmp11284
__e.Tick()
goto __selftop
}
return TailCall4(__e, SymFn(symshen_4lzy_a), tmp11281, tmp11283, V2184, tmp11284)


} else {
return False
}


}


}


}


panic("unreachable")
})
return __self
})()

tmp11299 := Call2(__e, ns2_1set, symshen_4lzy_a, tmp11279)


_ = tmp11299

tmp11300 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V2191 := __self1
_ = V2191
V2192 := __self2
_ = V2192
tmp11308 := GEq(__e, V2191, V2192)

if True == tmp11308 {
return True
} else {
tmp11306 := GConsP(__e, V2192)

if True == tmp11306 {
tmp11303 := GHd(__e, V2192)

tmp11304 := Call2(__e, SymFn(symshen_4occurs_1check_2), V2191, tmp11303)


if True == tmp11304 {
return True
} else {
tmp11301 := GTl(__e, V2192)

if SymFn(symshen_4occurs_1check_2) == __self {
__self1, __self2 = V2191, tmp11301
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4occurs_1check_2), V2191, tmp11301)


}


} else {
return False
}


}


panic("unreachable")
})
return __self
})()

tmp11309 := Call2(__e, ns2_1set, symshen_4occurs_1check_2, tmp11300)


_ = tmp11309

tmp11310 := MakeNative5(func(__e *ControlFlow, V2193, V2194, V2195, V2196, V2197 Obj) Obj {
tmp11311 := Call1(__e, V2193, V2194)


tmp11312 := Call1(__e, tmp11311, V2195)


tmp11313 := Call1(__e, tmp11312, V2196)


return TailCall1(__e, tmp11313, V2197)


panic("unreachable")
})

tmp11314 := Call2(__e, ns2_1set, symcall, tmp11310)


_ = tmp11314

tmp11315 := MakeNative5(func(__e *ControlFlow, V2204, V2205, V2206, V2207, V2208 Obj) Obj {
return TailCall2(__e, SymFn(symshen_4deref), V2204, V2205)
panic("unreachable")
})

tmp11316 := Call2(__e, ns2_1set, symreturn, tmp11315)


_ = tmp11316

tmp11317 := MakeNative5(func(__e *ControlFlow, V2215, V2216, V2217, V2218, V2219 Obj) Obj {
if True == V2215 {
return TailCall1(__e, SymFn(symthaw), V2219)
} else {
return False
}
panic("unreachable")
})

tmp11319 := Call2(__e, ns2_1set, symwhen, tmp11317)


_ = tmp11319

tmp11320 := MakeNative6(func(__e *ControlFlow, V2220, V2221, V2222, V2223, V2224, V2225 Obj) Obj {
tmp11321 := Call2(__e, SymFn(symshen_4lazyderef), V2220, V2222)


tmp11322 := Call2(__e, SymFn(symshen_4lazyderef), V2221, V2222)


return TailCall4(__e, SymFn(symshen_4lzy_a), tmp11321, tmp11322, V2222, V2225)


panic("unreachable")
})

tmp11323 := Call2(__e, ns2_1set, symis, tmp11320)


_ = tmp11323

tmp11324 := MakeNative6(func(__e *ControlFlow, V2226, V2227, V2228, V2229, V2230, V2231 Obj) Obj {
tmp11325 := Call2(__e, SymFn(symshen_4lazyderef), V2226, V2228)


tmp11326 := Call2(__e, SymFn(symshen_4lazyderef), V2227, V2228)


return TailCall4(__e, SymFn(symshen_4lzy_a_b), tmp11325, tmp11326, V2228, V2231)


panic("unreachable")
})

tmp11327 := Call2(__e, ns2_1set, symis_b, tmp11324)


_ = tmp11327

tmp11328 := MakeNative6(func(__e *ControlFlow, V2236, V2237, V2238, V2239, V2240, V2241 Obj) Obj {
return TailCall4(__e, SymFn(symshen_4bind_b), V2236, V2237, V2238, V2241)
panic("unreachable")
})

tmp11329 := Call2(__e, ns2_1set, symbind, tmp11328)


_ = tmp11329

tmp11330 := MakeNative5(func(__e *ControlFlow, V2242, V2243, V2244, V2245, V2246 Obj) Obj {
tmp11332 := Call2(__e, SymFn(symshen_4lazyderef), V2242, V2243)


tmp11333 := Call1(__e, SymFn(symshen_4pvar_2), tmp11332)


if True == tmp11333 {
return TailCall1(__e, SymFn(symthaw), V2246)
} else {
return False
}


panic("unreachable")
})

tmp11334 := Call2(__e, ns2_1set, symvar_2, tmp11330)


_ = tmp11334

tmp11335 := MakeNative1(func(__e *ControlFlow, V2249 Obj) Obj {
return MakeString("|prolog vector|")
panic("unreachable")
})

tmp11336 := Call2(__e, ns2_1set, symshen_4print_1prolog_1vector, tmp11335)


_ = tmp11336

tmp11337 := (func() Obj {
var __self Obj
__self = MakeNative5(func(__e *ControlFlow, __self1, __self2, __self3, __self4, __self5 Obj) Obj {
__selftop:
V2268 := __self1
_ = V2268
V2269 := __self2
_ = V2269
V2270 := __self3
_ = V2270
V2271 := __self4
_ = V2271
V2272 := __self5
_ = V2272
tmp11349 := GEq(__e, Nil, V2268)

if True == tmp11349 {
return False
} else {
tmp11347 := GConsP(__e, V2268)

if True == tmp11347 {
tmp11338 := GHd(__e, V2268)

tmp11339 := Call1(__e, tmp11338, V2269)


tmp11340 := Call1(__e, tmp11339, V2270)


tmp11341 := Call1(__e, tmp11340, V2271)


tmp11342 := Call1(__e, tmp11341, V2272)


let__9531 := tmp11342
_ = let__9531

tmp11345 := GEq(__e, let__9531, False)

if True == tmp11345 {
tmp11343 := GTl(__e, V2268)

if SymFn(symfork) == __self {
__self1, __self2, __self3, __self4, __self5 = tmp11343, V2269, V2270, V2271, V2272
__e.Tick()
goto __selftop
}
return TailCall5(__e, SymFn(symfork), tmp11343, V2269, V2270, V2271, V2272)


} else {
return let__9531
}


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("fork expects a list of literals\n"))
}


}


panic("unreachable")
})
return __self
})()

tmp11350 := Call2(__e, ns2_1set, symfork, tmp11337)


_ = tmp11350

tmp11351 := MakeNative(func(__e *ControlFlow) {
V2274 := __e.Get(1)
V2275 := __e.Get(2)
V2276 := __e.Get(3)
V2277 := __e.Get(4)
V2278 := __e.Get(5)
V2279 := __e.Get(6)
V2280 := __e.Get(7)
__e.ReturnDirect(func() Obj {
tmp11357 := Call1(__e, SymFn(symshen_4unlocked_2), V2278)


if True == tmp11357 {
tmp11352 := Call1(__e, SymFn(symshen_4newpv), V2277)


let__9532 := tmp11352
_ = let__9532

tmp11353 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp11353

tmp11354 := MakeNative0(func(__e *ControlFlow) Obj {
__e.TailApply(SymFn(symshen_4findall_1h), V2274, V2275, V2276, let__9532, V2277, V2278, V2279, V2280)
return nil
panic("unreachable")
})

tmp11355 := Call6(__e, SymFn(symis), let__9532, Nil, V2277, V2278, V2279, tmp11354)


return TailCall2(__e, SymFn(symshen_4gc), V2277, tmp11355)


} else {
return False
}


panic("unreachable")
}())
}, 7)

tmp11358 := Call2(__e, ns2_1set, symfindall, tmp11351)


_ = tmp11358

tmp11359 := MakeNative(func(__e *ControlFlow) {
V2282 := __e.Get(1)
V2283 := __e.Get(2)
V2284 := __e.Get(3)
V2285 := __e.Get(4)
V2286 := __e.Get(5)
V2287 := __e.Get(6)
V2288 := __e.Get(7)
V2289 := __e.Get(8)
__e.ReturnDirect(func() Obj {
tmp11364 := Call1(__e, SymFn(symshen_4unlocked_2), V2287)


var ifres11360 Obj

if True == tmp11364 {
tmp11361 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp11361

tmp11362 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall6(__e, SymFn(symshen_4overbind), V2282, V2285, V2286, V2287, V2288, V2289)
panic("unreachable")
})

tmp11363 := Call5(__e, SymFn(symcall), V2283, V2286, V2287, V2288, tmp11362)


ifres11360 = tmp11363


} else {
ifres11360 = False


}

let__9533 := ifres11360
_ = let__9533

tmp11369 := GEq(__e, let__9533, False)

if True == tmp11369 {
tmp11367 := Call1(__e, SymFn(symshen_4unlocked_2), V2287)


if True == tmp11367 {
tmp11365 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp11365

return TailCall6(__e, SymFn(symis_b), V2284, V2285, V2286, V2287, V2288, V2289)


} else {
return False
}


} else {
return let__9533
}


panic("unreachable")
}())
}, 8)

tmp11370 := Call2(__e, ns2_1set, symshen_4findall_1h, tmp11359)


_ = tmp11370

tmp11371 := MakeNative6(func(__e *ControlFlow, V2297, V2298, V2299, V2300, V2301, V2302 Obj) Obj {
tmp11372 := Call2(__e, SymFn(symshen_4deref), V2297, V2299)


tmp11373 := Call2(__e, SymFn(symshen_4lazyderef), V2298, V2299)


tmp11374 := GCons(__e, tmp11372, tmp11373)

tmp11375 := Call3(__e, SymFn(symshen_4bindv), V2298, tmp11374, V2299)


_ = tmp11375

return False


panic("unreachable")
})

tmp11376 := Call2(__e, ns2_1set, symshen_4overbind, tmp11371)


_ = tmp11376

tmp11377 := MakeNative1(func(__e *ControlFlow, V2305 Obj) Obj {
tmp11381 := GEq(__e, sym_7, V2305)

if True == tmp11381 {
return PrimGuard2(__e, symset, PrimSet, symshen_4_doccurs_d, True)
} else {
tmp11379 := GEq(__e, sym_1, V2305)

if True == tmp11379 {
return PrimGuard2(__e, symset, PrimSet, symshen_4_doccurs_d, False)
} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("occurs-check expects a + or a -.\n"))
}


}


panic("unreachable")
})

return TailCall2(__e, ns2_1set, symoccurs_1check, tmp11377)




panic("unreachable")
})

