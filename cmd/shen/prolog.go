package main

import . "github.com/pyrex41/shen-go/kl"

var PrologMain = MakeNative0(func(__e *ControlFlow) Obj {
tmp9554 := MakeNative1(func(__e *ControlFlow, V1599 Obj) Obj {
return TailCall2(__e, SymFn(symshen_4assert_d), V1599, symshen_4top)
panic("unreachable")
})

tmp9555 := Call2(__e, ns2_1set, symasserta, tmp9554)


_ = tmp9555

tmp9556 := MakeNative1(func(__e *ControlFlow, V1600 Obj) Obj {
return TailCall2(__e, SymFn(symshen_4assert_d), V1600, symshen_4bottom)
panic("unreachable")
})

tmp9557 := Call2(__e, ns2_1set, symassertz, tmp9556)


_ = tmp9557

tmp9558 := MakeNative2(func(__e *ControlFlow, V1601, V1602 Obj) Obj {
tmp9583 := GConsP(__e, V1601)

var ifres9576 Obj

if True == tmp9583 {
tmp9581 := GTl(__e, V1601)

tmp9582 := GConsP(__e, tmp9581)

var ifres9577 Obj

if True == tmp9582 {
tmp9578 := GTl(__e, V1601)

tmp9579 := GHd(__e, tmp9578)

tmp9580 := GEq(__e, sym_5_1_1, tmp9579)

ifres9577 = tmp9580


} else {
ifres9577 = False


}

ifres9576 = ifres9577


} else {
ifres9576 = False


}

if True == ifres9576 {
tmp9559 := GHd(__e, V1601)

tmp9560 := Call1(__e, SymFn(symshen_4predicate), tmp9559)


let__9336 := tmp9560
_ = let__9336

tmp9561 := GHd(__e, V1601)

tmp9562 := Call1(__e, SymFn(symshen_4terms), tmp9561)


let__9337 := tmp9562
_ = let__9337

tmp9563 := Call1(__e, SymFn(symlength), let__9337)


let__9338 := tmp9563
_ = let__9338

tmp9564 := Call1(__e, SymFn(symshen_4parameters), let__9338)


let__9339 := tmp9564
_ = let__9339

tmp9565 := Call1(__e, SymFn(symarity), let__9336)


let__9340 := tmp9565
_ = let__9340

tmp9571 := GEq(__e, let__9340, MakeInteger(-1))

var ifres9566 Obj

if True == tmp9571 {
tmp9567 := Call2(__e, SymFn(symshen_4create_1skeleton), let__9336, let__9339)


tmp9568 := Call1(__e, SymFn(symeval), tmp9567)


_ = tmp9568

tmp9569 := GValue(__e, sym_dproperty_1vector_d)

tmp9570 := Call4(__e, SymFn(symput), let__9336, symshen_4dynamic, Nil, tmp9569)


ifres9566 = tmp9570


} else {
ifres9566 = symshen_4skip


}

let__9341 := ifres9566
_ = let__9341

tmp9572 := GTl(__e, V1601)

tmp9573 := GTl(__e, tmp9572)

tmp9574 := Call5(__e, SymFn(symshen_4insert_1info), let__9336, let__9337, tmp9573, V1601, V1602)


let__9342 := tmp9574
_ = let__9342

return let__9336


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4assert_d)
}


panic("unreachable")
})

tmp9584 := Call2(__e, ns2_1set, symshen_4assert_d, tmp9558)


_ = tmp9584

tmp9585 := MakeNative1(func(__e *ControlFlow, V1612 Obj) Obj {
tmp9587 := GConsP(__e, V1612)

if True == tmp9587 {
return GHd(__e, V1612)
} else {
return V1612
}


panic("unreachable")
})

tmp9588 := Call2(__e, ns2_1set, symshen_4predicate, tmp9585)


_ = tmp9588

tmp9589 := MakeNative1(func(__e *ControlFlow, V1617 Obj) Obj {
tmp9591 := GConsP(__e, V1617)

if True == tmp9591 {
return GTl(__e, V1617)
} else {
return Nil
}


panic("unreachable")
})

tmp9592 := Call2(__e, ns2_1set, symshen_4terms, tmp9589)


_ = tmp9592

tmp9593 := MakeNative2(func(__e *ControlFlow, V1618, V1619 Obj) Obj {
tmp9594 := Call2(__e, SymFn(symshen_4dynamic_1default), V1618, V1619)


tmp9595 := GCons(__e, V1618, tmp9594)

return GCons(__e, symdefprolog, tmp9595)


panic("unreachable")
})

tmp9596 := Call2(__e, ns2_1set, symshen_4create_1skeleton, tmp9593)


_ = tmp9596

tmp9597 := MakeNative2(func(__e *ControlFlow, V1620, V1621 Obj) Obj {
tmp9598 := Call1(__e, SymFn(symshen_4cons_1form), V1621)


tmp9599 := GCons(__e, symshen_4dynamic, Nil)

tmp9600 := GCons(__e, V1620, tmp9599)

tmp9601 := GCons(__e, symget, tmp9600)

tmp9602 := GCons(__e, tmp9601, Nil)

tmp9603 := GCons(__e, tmp9598, tmp9602)

tmp9604 := GCons(__e, symshen_4call_1dynamic, tmp9603)

tmp9605 := PrimGuard1(__e, symintern, PrimIntern, MakeString(";"))

tmp9606 := GCons(__e, tmp9605, Nil)

tmp9607 := GCons(__e, tmp9604, tmp9606)

tmp9608 := GCons(__e, sym_5_1_1, tmp9607)

return TailCall2(__e, SymFn(symappend), V1621, tmp9608)


panic("unreachable")
})

tmp9609 := Call2(__e, ns2_1set, symshen_4dynamic_1default, tmp9597)


_ = tmp9609

tmp9610 := MakeNative5(func(__e *ControlFlow, V1622, V1623, V1624, V1625, V1626 Obj) Obj {
tmp9611 := Call1(__e, SymFn(symgensym), symg)


let__9343 := tmp9611
_ = let__9343

tmp9612 := GCons(__e, let__9343, Nil)

tmp9613 := GCons(__e, symdefprolog, tmp9612)

tmp9614 := GCons(__e, sym_5_1_1, V1624)

tmp9615 := Call2(__e, SymFn(symappend), V1623, tmp9614)


tmp9616 := Call2(__e, SymFn(symappend), tmp9613, tmp9615)


tmp9617 := Call1(__e, SymFn(symeval), tmp9616)


let__9344 := tmp9617
_ = let__9344

tmp9618 := Call1(__e, SymFn(symfn), let__9343)


tmp9619 := GCons(__e, let__9343, V1625)

tmp9620 := GCons(__e, tmp9618, tmp9619)

let__9345 := tmp9620
_ = let__9345

tmp9621 := GValue(__e, sym_dproperty_1vector_d)

tmp9622 := Call3(__e, SymFn(symget), V1622, symshen_4dynamic, tmp9621)


let__9346 := tmp9622
_ = let__9346

tmp9627 := GEq(__e, V1626, symshen_4top)

var ifres9623 Obj

if True == tmp9627 {
tmp9624 := GCons(__e, let__9345, let__9346)

ifres9623 = tmp9624


} else {
tmp9625 := GCons(__e, let__9345, Nil)

tmp9626 := Call2(__e, SymFn(symappend), let__9346, tmp9625)


ifres9623 = tmp9626


}

let__9347 := ifres9623
_ = let__9347

tmp9628 := GValue(__e, sym_dproperty_1vector_d)

return TailCall4(__e, SymFn(symput), V1622, symshen_4dynamic, let__9347, tmp9628)


panic("unreachable")
})

tmp9629 := Call2(__e, ns2_1set, symshen_4insert_1info, tmp9610)


_ = tmp9629

tmp9630 := MakeNative0(func(__e *ControlFlow) Obj {
tmp9631 := GValue(__e, symshen_4_dnames_d)

let__9348 := tmp9631
_ = let__9348

tmp9637 := Call1(__e, SymFn(symempty_2), let__9348)


var ifres9632 Obj

if True == tmp9637 {
tmp9633 := Call1(__e, SymFn(symgensym), symg)


ifres9632 = tmp9633


} else {
tmp9634 := GTl(__e, let__9348)

tmp9635 := PrimGuard2(__e, symset, PrimSet, symshen_4_dnames_d, tmp9634)

_ = tmp9635

tmp9636 := GHd(__e, let__9348)

ifres9632 = tmp9636


}

let__9349 := ifres9632
_ = let__9349

return let__9349


panic("unreachable")
})

tmp9638 := Call2(__e, ns2_1set, symshen_4newname, tmp9630)


_ = tmp9638

tmp9639 := (func() Obj {
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
tmp9651 := Call1(__e, SymFn(symshen_4unlocked_2), V1637)


var ifres9640 Obj

if True == tmp9651 {
tmp9641 := Call2(__e, SymFn(symshen_4lazyderef), V1635, V1636)


let__9351 := tmp9641
_ = let__9351

tmp9650 := GConsP(__e, let__9351)

var ifres9642 Obj

if True == tmp9650 {
tmp9643 := GHd(__e, let__9351)

tmp9644 := Call2(__e, SymFn(symshen_4lazyderef), tmp9643, V1636)


let__9352 := tmp9644
_ = let__9352

tmp9649 := GConsP(__e, let__9352)

var ifres9645 Obj

if True == tmp9649 {
tmp9646 := GHd(__e, let__9352)

let__9353 := tmp9646
_ = let__9353

tmp9647 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp9647

tmp9648 := Call6(__e, SymFn(symshen_4callrec), let__9353, V1634, V1636, V1637, V1638, V1639)


ifres9645 = tmp9648


} else {
ifres9645 = False


}

ifres9642 = ifres9645


} else {
ifres9642 = False


}

ifres9640 = ifres9642


} else {
ifres9640 = False


}

let__9350 := ifres9640
_ = let__9350

tmp9660 := GEq(__e, let__9350, False)

if True == tmp9660 {
tmp9658 := Call1(__e, SymFn(symshen_4unlocked_2), V1637)


if True == tmp9658 {
tmp9652 := Call2(__e, SymFn(symshen_4lazyderef), V1635, V1636)


let__9354 := tmp9652
_ = let__9354

tmp9656 := GConsP(__e, let__9354)

if True == tmp9656 {
tmp9653 := GTl(__e, let__9354)

let__9355 := tmp9653
_ = let__9355

tmp9654 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp9654

if SymFn(symshen_4call_1dynamic) == __self {
__self1, __self2, __self3, __self4, __self5, __self6 = V1634, let__9355, V1636, V1637, V1638, V1639
__e.Tick()
goto __selftop
}
return TailCall6(__e, SymFn(symshen_4call_1dynamic), V1634, let__9355, V1636, V1637, V1638, V1639)


} else {
return False
}


} else {
return False
}


} else {
return let__9350
}


panic("unreachable")
})
return __self
})()

tmp9661 := Call2(__e, ns2_1set, symshen_4call_1dynamic, tmp9639)


_ = tmp9661

tmp9662 := (func() Obj {
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
tmp9672 := GEq(__e, Nil, V1647)

if True == tmp9672 {
tmp9663 := Call1(__e, V1646, V1648)


tmp9664 := Call1(__e, tmp9663, V1649)


tmp9665 := Call1(__e, tmp9664, V1650)


return TailCall1(__e, tmp9665, V1651)


} else {
tmp9670 := GConsP(__e, V1647)

if True == tmp9670 {
tmp9666 := GHd(__e, V1647)

tmp9667 := Call1(__e, V1646, tmp9666)


tmp9668 := GTl(__e, V1647)

if SymFn(symshen_4callrec) == __self {
__self1, __self2, __self3, __self4, __self5, __self6 = tmp9667, tmp9668, V1648, V1649, V1650, V1651
__e.Tick()
goto __selftop
}
return TailCall6(__e, SymFn(symshen_4callrec), tmp9667, tmp9668, V1648, V1649, V1650, V1651)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4callrec)
}


}


panic("unreachable")
})
return __self
})()

tmp9673 := Call2(__e, ns2_1set, symshen_4callrec, tmp9662)


_ = tmp9673

tmp9674 := MakeNative1(func(__e *ControlFlow, V1652 Obj) Obj {
tmp9689 := GConsP(__e, V1652)

var ifres9682 Obj

if True == tmp9689 {
tmp9687 := GTl(__e, V1652)

tmp9688 := GConsP(__e, tmp9687)

var ifres9683 Obj

if True == tmp9688 {
tmp9684 := GTl(__e, V1652)

tmp9685 := GHd(__e, tmp9684)

tmp9686 := GEq(__e, sym_5_1_1, tmp9685)

ifres9683 = tmp9686


} else {
ifres9683 = False


}

ifres9682 = ifres9683


} else {
ifres9682 = False


}

if True == ifres9682 {
tmp9675 := GHd(__e, V1652)

tmp9676 := Call1(__e, SymFn(symshen_4predicate), tmp9675)


let__9356 := tmp9676
_ = let__9356

tmp9677 := GValue(__e, sym_dproperty_1vector_d)

tmp9678 := Call3(__e, SymFn(symget), let__9356, symshen_4dynamic, tmp9677)


let__9357 := tmp9678
_ = let__9357

tmp9679 := Call2(__e, SymFn(symshen_4retract_1clause), V1652, let__9357)


tmp9680 := GValue(__e, sym_dproperty_1vector_d)

return TailCall4(__e, SymFn(symput), let__9356, symshen_4dynamic, tmp9679, tmp9680)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symretract)
}


panic("unreachable")
})

tmp9690 := Call2(__e, ns2_1set, symretract, tmp9674)


_ = tmp9690

tmp9691 := MakeNative2(func(__e *ControlFlow, V1660, V1661 Obj) Obj {
tmp9718 := GEq(__e, Nil, V1661)

if True == tmp9718 {
return Nil
} else {
tmp9716 := GConsP(__e, V1661)

var ifres9704 Obj

if True == tmp9716 {
tmp9714 := GHd(__e, V1661)

tmp9715 := GConsP(__e, tmp9714)

var ifres9705 Obj

if True == tmp9715 {
tmp9711 := GHd(__e, V1661)

tmp9712 := GTl(__e, tmp9711)

tmp9713 := GConsP(__e, tmp9712)

var ifres9706 Obj

if True == tmp9713 {
tmp9707 := GHd(__e, V1661)

tmp9708 := GTl(__e, tmp9707)

tmp9709 := GTl(__e, tmp9708)

tmp9710 := GEq(__e, V1660, tmp9709)

ifres9706 = tmp9710


} else {
ifres9706 = False


}

ifres9705 = ifres9706


} else {
ifres9705 = False


}

ifres9704 = ifres9705


} else {
ifres9704 = False


}

if True == ifres9704 {
tmp9692 := GHd(__e, V1661)

tmp9693 := GTl(__e, tmp9692)

tmp9694 := GHd(__e, tmp9693)

tmp9695 := GValue(__e, symshen_4_dnames_d)

tmp9696 := GCons(__e, tmp9694, tmp9695)

tmp9697 := PrimGuard2(__e, symset, PrimSet, symshen_4_dnames_d, tmp9696)

_ = tmp9697

return GTl(__e, V1661)


} else {
tmp9702 := GConsP(__e, V1661)

if True == tmp9702 {
tmp9698 := GHd(__e, V1661)

tmp9699 := GTl(__e, V1661)

tmp9700 := Call2(__e, SymFn(symshen_4retract_1clause), V1660, tmp9699)


return GCons(__e, tmp9698, tmp9700)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4retract_1clause)
}


}


}


panic("unreachable")
})

tmp9719 := Call2(__e, ns2_1set, symshen_4retract_1clause, tmp9691)


_ = tmp9719

tmp9720 := MakeNative2(func(__e *ControlFlow, V1662, V1663 Obj) Obj {
tmp9721 := MakeNative1(func(__e *ControlFlow, Z1664 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4_5defprolog_6), Z1664)
panic("unreachable")
})

tmp9722 := GCons(__e, V1662, V1663)

return TailCall2(__e, SymFn(symcompile), tmp9721, tmp9722)


panic("unreachable")
})

tmp9723 := Call2(__e, ns2_1set, symshen_4compile_1prolog, tmp9720)


_ = tmp9723

tmp9724 := MakeNative1(func(__e *ControlFlow, V1665 Obj) Obj {
tmp9740 := GConsP(__e, V1665)

var ifres9725 Obj

if True == tmp9740 {
tmp9726 := Call1(__e, SymFn(symhead), V1665)


let__9359 := tmp9726
_ = let__9359

tmp9727 := Call1(__e, SymFn(symtail), V1665)


let__9360 := tmp9727
_ = let__9360

tmp9728 := Call1(__e, SymFn(symshen_4_5clauses_6), let__9360)


let__9361 := tmp9728
_ = let__9361

tmp9738 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9361)


var ifres9729 Obj

if True == tmp9738 {
tmp9730 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres9729 = tmp9730


} else {
tmp9731 := Call1(__e, SymFn(symshen_4_5_1out), let__9361)


let__9362 := tmp9731
_ = let__9362

tmp9732 := Call1(__e, SymFn(symshen_4in_1_6), let__9361)


let__9363 := tmp9732
_ = let__9363

tmp9733 := Call2(__e, SymFn(symshen_4prolog_1arity_1check), let__9359, let__9362)


let__9364 := tmp9733
_ = let__9364

tmp9734 := MakeNative1(func(__e *ControlFlow, Z1674 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4linearise_1clause), Z1674)
panic("unreachable")
})

tmp9735 := Call2(__e, SymFn(symmap), tmp9734, let__9362)


let__9365 := tmp9735
_ = let__9365

tmp9736 := Call2(__e, SymFn(symshen_4horn_1clause_1procedure), let__9359, let__9365)


tmp9737 := Call2(__e, SymFn(symshen_4comb), let__9363, tmp9736)


ifres9729 = tmp9737


}

ifres9725 = ifres9729


} else {
tmp9739 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres9725 = tmp9739


}

let__9358 := ifres9725
_ = let__9358

tmp9742 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9358)


if True == tmp9742 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__9358
}


panic("unreachable")
})

tmp9743 := Call2(__e, ns2_1set, symshen_4_5defprolog_6, tmp9724)


_ = tmp9743

tmp9744 := MakeNative2(func(__e *ControlFlow, V1677, V1678 Obj) Obj {
tmp9781 := GConsP(__e, V1678)

var ifres9766 Obj

if True == tmp9781 {
tmp9779 := GHd(__e, V1678)

tmp9780 := GConsP(__e, tmp9779)

var ifres9767 Obj

if True == tmp9780 {
tmp9776 := GHd(__e, V1678)

tmp9777 := GTl(__e, tmp9776)

tmp9778 := GConsP(__e, tmp9777)

var ifres9768 Obj

if True == tmp9778 {
tmp9772 := GHd(__e, V1678)

tmp9773 := GTl(__e, tmp9772)

tmp9774 := GTl(__e, tmp9773)

tmp9775 := GEq(__e, Nil, tmp9774)

var ifres9769 Obj

if True == tmp9775 {
tmp9770 := GTl(__e, V1678)

tmp9771 := GEq(__e, Nil, tmp9770)

ifres9769 = tmp9771


} else {
ifres9769 = False


}

ifres9768 = ifres9769


} else {
ifres9768 = False


}

ifres9767 = ifres9768


} else {
ifres9767 = False


}

ifres9766 = ifres9767


} else {
ifres9766 = False


}

if True == ifres9766 {
tmp9745 := GHd(__e, V1678)

tmp9746 := GHd(__e, tmp9745)

return TailCall1(__e, SymFn(symlength), tmp9746)


} else {
tmp9764 := GConsP(__e, V1678)

var ifres9752 Obj

if True == tmp9764 {
tmp9762 := GHd(__e, V1678)

tmp9763 := GConsP(__e, tmp9762)

var ifres9753 Obj

if True == tmp9763 {
tmp9759 := GHd(__e, V1678)

tmp9760 := GTl(__e, tmp9759)

tmp9761 := GConsP(__e, tmp9760)

var ifres9754 Obj

if True == tmp9761 {
tmp9755 := GHd(__e, V1678)

tmp9756 := GTl(__e, tmp9755)

tmp9757 := GTl(__e, tmp9756)

tmp9758 := GEq(__e, Nil, tmp9757)

ifres9754 = tmp9758


} else {
ifres9754 = False


}

ifres9753 = ifres9754


} else {
ifres9753 = False


}

ifres9752 = ifres9753


} else {
ifres9752 = False


}

if True == ifres9752 {
tmp9747 := GHd(__e, V1678)

tmp9748 := GHd(__e, tmp9747)

tmp9749 := Call1(__e, SymFn(symlength), tmp9748)


tmp9750 := GTl(__e, V1678)

return TailCall3(__e, SymFn(symshen_4pac_1h), V1677, tmp9749, tmp9750)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4prolog_1arity_1check)
}


}


panic("unreachable")
})

tmp9782 := Call2(__e, ns2_1set, symshen_4prolog_1arity_1check, tmp9744)


_ = tmp9782

tmp9783 := (func() Obj {
var __self Obj
__self = MakeNative3(func(__e *ControlFlow, __self1, __self2, __self3 Obj) Obj {
__selftop:
V1683 := __self1
_ = V1683
V1684 := __self2
_ = V1684
V1685 := __self3
_ = V1685
tmp9798 := GEq(__e, Nil, V1685)

if True == tmp9798 {
return V1684
} else {
tmp9796 := GConsP(__e, V1685)

var ifres9793 Obj

if True == tmp9796 {
tmp9794 := GHd(__e, V1685)

tmp9795 := GConsP(__e, tmp9794)

ifres9793 = tmp9795


} else {
ifres9793 = False


}

if True == ifres9793 {
tmp9788 := GHd(__e, V1685)

tmp9789 := GHd(__e, tmp9788)

tmp9790 := Call1(__e, SymFn(symlength), tmp9789)


tmp9791 := GEq(__e, V1684, tmp9790)

if True == tmp9791 {
tmp9784 := GTl(__e, V1685)

if SymFn(symshen_4pac_1h) == __self {
__self1, __self2, __self3 = V1683, V1684, tmp9784
__e.Tick()
goto __selftop
}
return TailCall3(__e, SymFn(symshen_4pac_1h), V1683, V1684, tmp9784)


} else {
tmp9785 := Call3(__e, SymFn(symshen_4app), V1683, MakeString("\n"), symshen_4a)


return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("arity error in prolog procedure "), tmp9785))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("arity error in prolog procedure "), tmp9785)
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

tmp9799 := Call2(__e, ns2_1set, symshen_4pac_1h, tmp9783)


_ = tmp9799

tmp9800 := MakeNative1(func(__e *ControlFlow, V1686 Obj) Obj {
tmp9801 := Call1(__e, SymFn(symshen_4_5clause_6), V1686)


let__9367 := tmp9801
_ = let__9367

tmp9814 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9367)


var ifres9802 Obj

if True == tmp9814 {
tmp9803 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres9802 = tmp9803


} else {
tmp9804 := Call1(__e, SymFn(symshen_4_5_1out), let__9367)


let__9368 := tmp9804
_ = let__9368

tmp9805 := Call1(__e, SymFn(symshen_4in_1_6), let__9367)


let__9369 := tmp9805
_ = let__9369

tmp9806 := Call1(__e, SymFn(symshen_4_5clauses_6), let__9369)


let__9370 := tmp9806
_ = let__9370

tmp9813 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9370)


var ifres9807 Obj

if True == tmp9813 {
tmp9808 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres9807 = tmp9808


} else {
tmp9809 := Call1(__e, SymFn(symshen_4_5_1out), let__9370)


let__9371 := tmp9809
_ = let__9371

tmp9810 := Call1(__e, SymFn(symshen_4in_1_6), let__9370)


let__9372 := tmp9810
_ = let__9372

tmp9811 := GCons(__e, let__9368, let__9371)

tmp9812 := Call2(__e, SymFn(symshen_4comb), let__9372, tmp9811)


ifres9807 = tmp9812


}

ifres9802 = ifres9807


}

let__9366 := ifres9802
_ = let__9366

tmp9830 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9366)


if True == tmp9830 {
tmp9815 := Call1(__e, SymFn(sym_5_b_6), V1686)


let__9374 := tmp9815
_ = let__9374

tmp9826 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9374)


var ifres9816 Obj

if True == tmp9826 {
tmp9817 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres9816 = tmp9817


} else {
tmp9818 := Call1(__e, SymFn(symshen_4_5_1out), let__9374)


let__9375 := tmp9818
_ = let__9375

tmp9819 := Call1(__e, SymFn(symshen_4in_1_6), let__9374)


let__9376 := tmp9819
_ = let__9376

tmp9824 := Call1(__e, SymFn(symempty_2), let__9375)


var ifres9820 Obj

if True == tmp9824 {
ifres9820 = Nil


} else {
tmp9821 := Call3(__e, SymFn(symshen_4app), let__9375, MakeString("\n ..."), symshen_4r)


tmp9823 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("Prolog syntax error here:\n "), tmp9821))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("Prolog syntax error here:\n "), tmp9821)
return Call1(__e, SymFn(symsimple_1error), __typedArg0)
})()

ifres9820 = tmp9823


}

tmp9825 := Call2(__e, SymFn(symshen_4comb), let__9376, ifres9820)


ifres9816 = tmp9825


}

let__9373 := ifres9816
_ = let__9373

tmp9828 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9373)


if True == tmp9828 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__9373
}


} else {
return let__9366
}


panic("unreachable")
})

tmp9831 := Call2(__e, ns2_1set, symshen_4_5clauses_6, tmp9800)


_ = tmp9831

tmp9832 := MakeNative1(func(__e *ControlFlow, V1698 Obj) Obj {
tmp9846 := GConsP(__e, V1698)

var ifres9839 Obj

if True == tmp9846 {
tmp9844 := GTl(__e, V1698)

tmp9845 := GConsP(__e, tmp9844)

var ifres9840 Obj

if True == tmp9845 {
tmp9841 := GTl(__e, V1698)

tmp9842 := GTl(__e, tmp9841)

tmp9843 := GEq(__e, Nil, tmp9842)

ifres9840 = tmp9843


} else {
ifres9840 = False


}

ifres9839 = ifres9840


} else {
ifres9839 = False


}

if True == ifres9839 {
tmp9833 := GHd(__e, V1698)

tmp9834 := GTl(__e, V1698)

tmp9835 := GHd(__e, tmp9834)

tmp9836 := Call2(__e, SymFn(sym_8p), tmp9833, tmp9835)


tmp9837 := Call1(__e, SymFn(symshen_4linearise), tmp9836)


return TailCall1(__e, SymFn(symshen_4lch), tmp9837)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4linearise_1clause)
}


panic("unreachable")
})

tmp9847 := Call2(__e, ns2_1set, symshen_4linearise_1clause, tmp9832)


_ = tmp9847

tmp9848 := MakeNative1(func(__e *ControlFlow, V1699 Obj) Obj {
tmp9854 := Call1(__e, SymFn(symtuple_2), V1699)


if True == tmp9854 {
tmp9849 := Call1(__e, SymFn(symfst), V1699)


tmp9850 := Call1(__e, SymFn(symsnd), V1699)


tmp9851 := Call1(__e, SymFn(symshen_4lchh), tmp9850)


tmp9852 := GCons(__e, tmp9851, Nil)

return GCons(__e, tmp9849, tmp9852)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4lch)
}


panic("unreachable")
})

tmp9855 := Call2(__e, ns2_1set, symshen_4lch, tmp9848)


_ = tmp9855

tmp9856 := MakeNative1(func(__e *ControlFlow, V1700 Obj) Obj {
tmp9910 := GConsP(__e, V1700)

var ifres9868 Obj

if True == tmp9910 {
tmp9908 := GHd(__e, V1700)

tmp9909 := GEq(__e, symwhere, tmp9908)

var ifres9869 Obj

if True == tmp9909 {
tmp9906 := GTl(__e, V1700)

tmp9907 := GConsP(__e, tmp9906)

var ifres9870 Obj

if True == tmp9907 {
tmp9903 := GTl(__e, V1700)

tmp9904 := GHd(__e, tmp9903)

tmp9905 := GConsP(__e, tmp9904)

var ifres9871 Obj

if True == tmp9905 {
tmp9899 := GTl(__e, V1700)

tmp9900 := GHd(__e, tmp9899)

tmp9901 := GHd(__e, tmp9900)

tmp9902 := GEq(__e, sym_a, tmp9901)

var ifres9872 Obj

if True == tmp9902 {
tmp9895 := GTl(__e, V1700)

tmp9896 := GHd(__e, tmp9895)

tmp9897 := GTl(__e, tmp9896)

tmp9898 := GConsP(__e, tmp9897)

var ifres9873 Obj

if True == tmp9898 {
tmp9890 := GTl(__e, V1700)

tmp9891 := GHd(__e, tmp9890)

tmp9892 := GTl(__e, tmp9891)

tmp9893 := GTl(__e, tmp9892)

tmp9894 := GConsP(__e, tmp9893)

var ifres9874 Obj

if True == tmp9894 {
tmp9884 := GTl(__e, V1700)

tmp9885 := GHd(__e, tmp9884)

tmp9886 := GTl(__e, tmp9885)

tmp9887 := GTl(__e, tmp9886)

tmp9888 := GTl(__e, tmp9887)

tmp9889 := GEq(__e, Nil, tmp9888)

var ifres9875 Obj

if True == tmp9889 {
tmp9881 := GTl(__e, V1700)

tmp9882 := GTl(__e, tmp9881)

tmp9883 := GConsP(__e, tmp9882)

var ifres9876 Obj

if True == tmp9883 {
tmp9877 := GTl(__e, V1700)

tmp9878 := GTl(__e, tmp9877)

tmp9879 := GTl(__e, tmp9878)

tmp9880 := GEq(__e, Nil, tmp9879)

ifres9876 = tmp9880


} else {
ifres9876 = False


}

ifres9875 = ifres9876


} else {
ifres9875 = False


}

ifres9874 = ifres9875


} else {
ifres9874 = False


}

ifres9873 = ifres9874


} else {
ifres9873 = False


}

ifres9872 = ifres9873


} else {
ifres9872 = False


}

ifres9871 = ifres9872


} else {
ifres9871 = False


}

ifres9870 = ifres9871


} else {
ifres9870 = False


}

ifres9869 = ifres9870


} else {
ifres9869 = False


}

ifres9868 = ifres9869


} else {
ifres9868 = False


}

if True == ifres9868 {
tmp9858 := GValue(__e, symshen_4_doccurs_d)

var ifres9857 Obj

if True == tmp9858 {
ifres9857 = symis_b


} else {
ifres9857 = symis


}

tmp9859 := GTl(__e, V1700)

tmp9860 := GHd(__e, tmp9859)

tmp9861 := GTl(__e, tmp9860)

tmp9862 := GCons(__e, ifres9857, tmp9861)

tmp9863 := GTl(__e, V1700)

tmp9864 := GTl(__e, tmp9863)

tmp9865 := GHd(__e, tmp9864)

tmp9866 := Call1(__e, SymFn(symshen_4lchh), tmp9865)


return GCons(__e, tmp9862, tmp9866)


} else {
return V1700
}


panic("unreachable")
})

tmp9911 := Call2(__e, ns2_1set, symshen_4lchh, tmp9856)


_ = tmp9911

tmp9912 := MakeNative1(func(__e *ControlFlow, V1701 Obj) Obj {
tmp9913 := Call1(__e, SymFn(symshen_4_5head_6), V1701)


let__9378 := tmp9913
_ = let__9378

tmp9936 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9378)


var ifres9914 Obj

if True == tmp9936 {
tmp9915 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres9914 = tmp9915


} else {
tmp9916 := Call1(__e, SymFn(symshen_4_5_1out), let__9378)


let__9379 := tmp9916
_ = let__9379

tmp9917 := Call1(__e, SymFn(symshen_4in_1_6), let__9378)


let__9380 := tmp9917
_ = let__9380

tmp9935 := Call2(__e, SymFn(symshen_4hds_a_2), let__9380, sym_5_1_1)


var ifres9918 Obj

if True == tmp9935 {
tmp9919 := Call1(__e, SymFn(symtail), let__9380)


let__9381 := tmp9919
_ = let__9381

tmp9920 := Call1(__e, SymFn(symshen_4_5body_6), let__9381)


let__9382 := tmp9920
_ = let__9382

tmp9933 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9382)


var ifres9921 Obj

if True == tmp9933 {
tmp9922 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres9921 = tmp9922


} else {
tmp9923 := Call1(__e, SymFn(symshen_4_5_1out), let__9382)


let__9383 := tmp9923
_ = let__9383

tmp9924 := Call1(__e, SymFn(symshen_4in_1_6), let__9382)


let__9384 := tmp9924
_ = let__9384

tmp9925 := Call1(__e, SymFn(symshen_4_5sc_6), let__9384)


let__9385 := tmp9925
_ = let__9385

tmp9932 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9385)


var ifres9926 Obj

if True == tmp9932 {
tmp9927 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres9926 = tmp9927


} else {
tmp9928 := Call1(__e, SymFn(symshen_4in_1_6), let__9385)


let__9386 := tmp9928
_ = let__9386

tmp9929 := GCons(__e, let__9383, Nil)

tmp9930 := GCons(__e, let__9379, tmp9929)

tmp9931 := Call2(__e, SymFn(symshen_4comb), let__9386, tmp9930)


ifres9926 = tmp9931


}

ifres9921 = ifres9926


}

ifres9918 = ifres9921


} else {
tmp9934 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres9918 = tmp9934


}

ifres9914 = ifres9918


}

let__9377 := ifres9914
_ = let__9377

tmp9938 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9377)


if True == tmp9938 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__9377
}


panic("unreachable")
})

tmp9939 := Call2(__e, ns2_1set, symshen_4_5clause_6, tmp9912)


_ = tmp9939

tmp9940 := MakeNative1(func(__e *ControlFlow, V1712 Obj) Obj {
tmp9941 := Call1(__e, SymFn(symshen_4_5hterm_6), V1712)


let__9388 := tmp9941
_ = let__9388

tmp9954 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9388)


var ifres9942 Obj

if True == tmp9954 {
tmp9943 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres9942 = tmp9943


} else {
tmp9944 := Call1(__e, SymFn(symshen_4_5_1out), let__9388)


let__9389 := tmp9944
_ = let__9389

tmp9945 := Call1(__e, SymFn(symshen_4in_1_6), let__9388)


let__9390 := tmp9945
_ = let__9390

tmp9946 := Call1(__e, SymFn(symshen_4_5head_6), let__9390)


let__9391 := tmp9946
_ = let__9391

tmp9953 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9391)


var ifres9947 Obj

if True == tmp9953 {
tmp9948 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres9947 = tmp9948


} else {
tmp9949 := Call1(__e, SymFn(symshen_4_5_1out), let__9391)


let__9392 := tmp9949
_ = let__9392

tmp9950 := Call1(__e, SymFn(symshen_4in_1_6), let__9391)


let__9393 := tmp9950
_ = let__9393

tmp9951 := GCons(__e, let__9389, let__9392)

tmp9952 := Call2(__e, SymFn(symshen_4comb), let__9393, tmp9951)


ifres9947 = tmp9952


}

ifres9942 = ifres9947


}

let__9387 := ifres9942
_ = let__9387

tmp9964 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9387)


if True == tmp9964 {
tmp9955 := Call1(__e, SymFn(sym_5e_6), V1712)


let__9395 := tmp9955
_ = let__9395

tmp9960 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9395)


var ifres9956 Obj

if True == tmp9960 {
tmp9957 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres9956 = tmp9957


} else {
tmp9958 := Call1(__e, SymFn(symshen_4in_1_6), let__9395)


let__9396 := tmp9958
_ = let__9396

tmp9959 := Call2(__e, SymFn(symshen_4comb), let__9396, Nil)


ifres9956 = tmp9959


}

let__9394 := ifres9956
_ = let__9394

tmp9962 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9394)


if True == tmp9962 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__9394
}


} else {
return let__9387
}


panic("unreachable")
})

tmp9965 := Call2(__e, ns2_1set, symshen_4_5head_6, tmp9940)


_ = tmp9965

tmp9966 := MakeNative1(func(__e *ControlFlow, V1723 Obj) Obj {
tmp9978 := GConsP(__e, V1723)

var ifres9967 Obj

if True == tmp9978 {
tmp9968 := Call1(__e, SymFn(symhead), V1723)


let__9398 := tmp9968
_ = let__9398

tmp9969 := Call1(__e, SymFn(symtail), V1723)


let__9399 := tmp9969
_ = let__9399

tmp9976 := Call1(__e, SymFn(symatom_2), let__9398)


var ifres9973 Obj

if True == tmp9976 {
tmp9974 := Call1(__e, SymFn(symshen_4prolog_1keyword_2), let__9398)


ifres9973 = GNot(__e, tmp9974)


} else {
ifres9973 = False


}

var ifres9970 Obj

if True == ifres9973 {
tmp9971 := Call2(__e, SymFn(symshen_4comb), let__9399, let__9398)


ifres9970 = tmp9971


} else {
tmp9972 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres9970 = tmp9972


}

ifres9967 = ifres9970


} else {
tmp9977 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres9967 = tmp9977


}

let__9397 := ifres9967
_ = let__9397

tmp10132 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9397)


if True == tmp10132 {
tmp9988 := GConsP(__e, V1723)

var ifres9979 Obj

if True == tmp9988 {
tmp9980 := Call1(__e, SymFn(symhead), V1723)


let__9401 := tmp9980
_ = let__9401

tmp9981 := Call1(__e, SymFn(symtail), V1723)


let__9402 := tmp9981
_ = let__9402

tmp9985 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":"))

tmp9986 := GEq(__e, let__9401, tmp9985)

var ifres9982 Obj

if True == tmp9986 {
tmp9983 := Call2(__e, SymFn(symshen_4comb), let__9402, let__9401)


ifres9982 = tmp9983


} else {
tmp9984 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres9982 = tmp9984


}

ifres9979 = ifres9982


} else {
tmp9987 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres9979 = tmp9987


}

let__9400 := ifres9979
_ = let__9400

tmp10130 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9400)


if True == tmp10130 {
tmp10018 := Call1(__e, SymFn(symshen_4ccons_2), V1723)


var ifres9989 Obj

if True == tmp10018 {
tmp9990 := Call1(__e, SymFn(symhead), V1723)


let__9404 := tmp9990
_ = let__9404

tmp9991 := Call1(__e, SymFn(symtail), V1723)


let__9405 := tmp9991
_ = let__9405

tmp10016 := Call2(__e, SymFn(symshen_4hds_a_2), let__9404, symcons)


var ifres9992 Obj

if True == tmp10016 {
tmp9993 := Call1(__e, SymFn(symtail), let__9404)


let__9406 := tmp9993
_ = let__9406

tmp9994 := Call1(__e, SymFn(symshen_4_5hterm1_6), let__9406)


let__9407 := tmp9994
_ = let__9407

tmp10014 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9407)


var ifres9995 Obj

if True == tmp10014 {
tmp9996 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres9995 = tmp9996


} else {
tmp9997 := Call1(__e, SymFn(symshen_4_5_1out), let__9407)


let__9408 := tmp9997
_ = let__9408

tmp9998 := Call1(__e, SymFn(symshen_4in_1_6), let__9407)


let__9409 := tmp9998
_ = let__9409

tmp9999 := Call1(__e, SymFn(symshen_4_5hterm2_6), let__9409)


let__9410 := tmp9999
_ = let__9410

tmp10013 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9410)


var ifres10000 Obj

if True == tmp10013 {
tmp10001 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10000 = tmp10001


} else {
tmp10002 := Call1(__e, SymFn(symshen_4_5_1out), let__9410)


let__9411 := tmp10002
_ = let__9411

tmp10003 := Call1(__e, SymFn(symshen_4in_1_6), let__9410)


let__9412 := tmp10003
_ = let__9412

tmp10004 := Call1(__e, SymFn(sym_5end_6), let__9412)


let__9413 := tmp10004
_ = let__9413

tmp10012 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9413)


var ifres10005 Obj

if True == tmp10012 {
tmp10006 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10005 = tmp10006


} else {
tmp10007 := Call1(__e, SymFn(symshen_4in_1_6), let__9413)


let__9414 := tmp10007
_ = let__9414

tmp10008 := GCons(__e, let__9411, Nil)

tmp10009 := GCons(__e, let__9408, tmp10008)

tmp10010 := GCons(__e, symcons, tmp10009)

tmp10011 := Call2(__e, SymFn(symshen_4comb), let__9405, tmp10010)


ifres10005 = tmp10011


}

ifres10000 = ifres10005


}

ifres9995 = ifres10000


}

ifres9992 = ifres9995


} else {
tmp10015 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres9992 = tmp10015


}

ifres9989 = ifres9992


} else {
tmp10017 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres9989 = tmp10017


}

let__9403 := ifres9989
_ = let__9403

tmp10128 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9403)


if True == tmp10128 {
tmp10041 := Call1(__e, SymFn(symshen_4ccons_2), V1723)


var ifres10019 Obj

if True == tmp10041 {
tmp10020 := Call1(__e, SymFn(symhead), V1723)


let__9416 := tmp10020
_ = let__9416

tmp10021 := Call1(__e, SymFn(symtail), V1723)


let__9417 := tmp10021
_ = let__9417

tmp10039 := Call2(__e, SymFn(symshen_4hds_a_2), let__9416, sym_7)


var ifres10022 Obj

if True == tmp10039 {
tmp10023 := Call1(__e, SymFn(symtail), let__9416)


let__9418 := tmp10023
_ = let__9418

tmp10024 := Call1(__e, SymFn(symshen_4_5hterm_6), let__9418)


let__9419 := tmp10024
_ = let__9419

tmp10037 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9419)


var ifres10025 Obj

if True == tmp10037 {
tmp10026 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10025 = tmp10026


} else {
tmp10027 := Call1(__e, SymFn(symshen_4_5_1out), let__9419)


let__9420 := tmp10027
_ = let__9420

tmp10028 := Call1(__e, SymFn(symshen_4in_1_6), let__9419)


let__9421 := tmp10028
_ = let__9421

tmp10029 := Call1(__e, SymFn(sym_5end_6), let__9421)


let__9422 := tmp10029
_ = let__9422

tmp10036 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9422)


var ifres10030 Obj

if True == tmp10036 {
tmp10031 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10030 = tmp10031


} else {
tmp10032 := Call1(__e, SymFn(symshen_4in_1_6), let__9422)


let__9423 := tmp10032
_ = let__9423

tmp10033 := GCons(__e, let__9420, Nil)

tmp10034 := GCons(__e, symshen_4_7m, tmp10033)

tmp10035 := Call2(__e, SymFn(symshen_4comb), let__9417, tmp10034)


ifres10030 = tmp10035


}

ifres10025 = ifres10030


}

ifres10022 = ifres10025


} else {
tmp10038 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10022 = tmp10038


}

ifres10019 = ifres10022


} else {
tmp10040 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10019 = tmp10040


}

let__9415 := ifres10019
_ = let__9415

tmp10126 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9415)


if True == tmp10126 {
tmp10064 := Call1(__e, SymFn(symshen_4ccons_2), V1723)


var ifres10042 Obj

if True == tmp10064 {
tmp10043 := Call1(__e, SymFn(symhead), V1723)


let__9425 := tmp10043
_ = let__9425

tmp10044 := Call1(__e, SymFn(symtail), V1723)


let__9426 := tmp10044
_ = let__9426

tmp10062 := Call2(__e, SymFn(symshen_4hds_a_2), let__9425, sym_1)


var ifres10045 Obj

if True == tmp10062 {
tmp10046 := Call1(__e, SymFn(symtail), let__9425)


let__9427 := tmp10046
_ = let__9427

tmp10047 := Call1(__e, SymFn(symshen_4_5hterm_6), let__9427)


let__9428 := tmp10047
_ = let__9428

tmp10060 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9428)


var ifres10048 Obj

if True == tmp10060 {
tmp10049 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10048 = tmp10049


} else {
tmp10050 := Call1(__e, SymFn(symshen_4_5_1out), let__9428)


let__9429 := tmp10050
_ = let__9429

tmp10051 := Call1(__e, SymFn(symshen_4in_1_6), let__9428)


let__9430 := tmp10051
_ = let__9430

tmp10052 := Call1(__e, SymFn(sym_5end_6), let__9430)


let__9431 := tmp10052
_ = let__9431

tmp10059 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9431)


var ifres10053 Obj

if True == tmp10059 {
tmp10054 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10053 = tmp10054


} else {
tmp10055 := Call1(__e, SymFn(symshen_4in_1_6), let__9431)


let__9432 := tmp10055
_ = let__9432

tmp10056 := GCons(__e, let__9429, Nil)

tmp10057 := GCons(__e, symshen_4_1m, tmp10056)

tmp10058 := Call2(__e, SymFn(symshen_4comb), let__9426, tmp10057)


ifres10053 = tmp10058


}

ifres10048 = ifres10053


}

ifres10045 = ifres10048


} else {
tmp10061 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10045 = tmp10061


}

ifres10042 = ifres10045


} else {
tmp10063 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10042 = tmp10063


}

let__9424 := ifres10042
_ = let__9424

tmp10124 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9424)


if True == tmp10124 {
tmp10091 := Call1(__e, SymFn(symshen_4ccons_2), V1723)


var ifres10065 Obj

if True == tmp10091 {
tmp10066 := Call1(__e, SymFn(symhead), V1723)


let__9434 := tmp10066
_ = let__9434

tmp10067 := Call1(__e, SymFn(symtail), V1723)


let__9435 := tmp10067
_ = let__9435

tmp10089 := Call2(__e, SymFn(symshen_4hds_a_2), let__9434, symmode)


var ifres10068 Obj

if True == tmp10089 {
tmp10069 := Call1(__e, SymFn(symtail), let__9434)


let__9436 := tmp10069
_ = let__9436

tmp10070 := Call1(__e, SymFn(symshen_4_5hterm_6), let__9436)


let__9437 := tmp10070
_ = let__9437

tmp10087 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9437)


var ifres10071 Obj

if True == tmp10087 {
tmp10072 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10071 = tmp10072


} else {
tmp10073 := Call1(__e, SymFn(symshen_4_5_1out), let__9437)


let__9438 := tmp10073
_ = let__9438

tmp10074 := Call1(__e, SymFn(symshen_4in_1_6), let__9437)


let__9439 := tmp10074
_ = let__9439

tmp10086 := Call2(__e, SymFn(symshen_4hds_a_2), let__9439, sym_7)


var ifres10075 Obj

if True == tmp10086 {
tmp10076 := Call1(__e, SymFn(symtail), let__9439)


let__9440 := tmp10076
_ = let__9440

tmp10077 := Call1(__e, SymFn(sym_5end_6), let__9440)


let__9441 := tmp10077
_ = let__9441

tmp10084 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9441)


var ifres10078 Obj

if True == tmp10084 {
tmp10079 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10078 = tmp10079


} else {
tmp10080 := Call1(__e, SymFn(symshen_4in_1_6), let__9441)


let__9442 := tmp10080
_ = let__9442

tmp10081 := GCons(__e, let__9438, Nil)

tmp10082 := GCons(__e, symshen_4_7m, tmp10081)

tmp10083 := Call2(__e, SymFn(symshen_4comb), let__9435, tmp10082)


ifres10078 = tmp10083


}

ifres10075 = ifres10078


} else {
tmp10085 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10075 = tmp10085


}

ifres10071 = ifres10075


}

ifres10068 = ifres10071


} else {
tmp10088 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10068 = tmp10088


}

ifres10065 = ifres10068


} else {
tmp10090 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10065 = tmp10090


}

let__9433 := ifres10065
_ = let__9433

tmp10122 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9433)


if True == tmp10122 {
tmp10118 := Call1(__e, SymFn(symshen_4ccons_2), V1723)


var ifres10092 Obj

if True == tmp10118 {
tmp10093 := Call1(__e, SymFn(symhead), V1723)


let__9444 := tmp10093
_ = let__9444

tmp10094 := Call1(__e, SymFn(symtail), V1723)


let__9445 := tmp10094
_ = let__9445

tmp10116 := Call2(__e, SymFn(symshen_4hds_a_2), let__9444, symmode)


var ifres10095 Obj

if True == tmp10116 {
tmp10096 := Call1(__e, SymFn(symtail), let__9444)


let__9446 := tmp10096
_ = let__9446

tmp10097 := Call1(__e, SymFn(symshen_4_5hterm_6), let__9446)


let__9447 := tmp10097
_ = let__9447

tmp10114 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9447)


var ifres10098 Obj

if True == tmp10114 {
tmp10099 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10098 = tmp10099


} else {
tmp10100 := Call1(__e, SymFn(symshen_4_5_1out), let__9447)


let__9448 := tmp10100
_ = let__9448

tmp10101 := Call1(__e, SymFn(symshen_4in_1_6), let__9447)


let__9449 := tmp10101
_ = let__9449

tmp10113 := Call2(__e, SymFn(symshen_4hds_a_2), let__9449, sym_1)


var ifres10102 Obj

if True == tmp10113 {
tmp10103 := Call1(__e, SymFn(symtail), let__9449)


let__9450 := tmp10103
_ = let__9450

tmp10104 := Call1(__e, SymFn(sym_5end_6), let__9450)


let__9451 := tmp10104
_ = let__9451

tmp10111 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9451)


var ifres10105 Obj

if True == tmp10111 {
tmp10106 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10105 = tmp10106


} else {
tmp10107 := Call1(__e, SymFn(symshen_4in_1_6), let__9451)


let__9452 := tmp10107
_ = let__9452

tmp10108 := GCons(__e, let__9448, Nil)

tmp10109 := GCons(__e, symshen_4_1m, tmp10108)

tmp10110 := Call2(__e, SymFn(symshen_4comb), let__9445, tmp10109)


ifres10105 = tmp10110


}

ifres10102 = ifres10105


} else {
tmp10112 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10102 = tmp10112


}

ifres10098 = ifres10102


}

ifres10095 = ifres10098


} else {
tmp10115 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10095 = tmp10115


}

ifres10092 = ifres10095


} else {
tmp10117 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10092 = tmp10117


}

let__9443 := ifres10092
_ = let__9443

tmp10120 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9443)


if True == tmp10120 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__9443
}


} else {
return let__9433
}


} else {
return let__9424
}


} else {
return let__9415
}


} else {
return let__9403
}


} else {
return let__9400
}


} else {
return let__9397
}


panic("unreachable")
})

tmp10133 := Call2(__e, ns2_1set, symshen_4_5hterm_6, tmp9966)


_ = tmp10133

tmp10134 := MakeNative1(func(__e *ControlFlow, V1780 Obj) Obj {
tmp10135 := PrimGuard1(__e, symintern, PrimIntern, MakeString(";"))

tmp10136 := GCons(__e, sym_5_1_1, Nil)

tmp10137 := GCons(__e, tmp10135, tmp10136)

return TailCall2(__e, SymFn(symelement_2), V1780, tmp10137)


panic("unreachable")
})

tmp10138 := Call2(__e, ns2_1set, symshen_4prolog_1keyword_2, tmp10134)


_ = tmp10138

tmp10139 := MakeNative1(func(__e *ControlFlow, V1781 Obj) Obj {
tmp10147 := PrimGuard1(__e, symsymbol_2, PrimIsSymbol, V1781)

if True == tmp10147 {
return True
} else {
tmp10145 := PrimGuard1(__e, symstring_2, PrimIsString, V1781)

if True == tmp10145 {
return True
} else {
tmp10143 := Call1(__e, SymFn(symboolean_2), V1781)


if True == tmp10143 {
return True
} else {
tmp10141 := PrimGuard1(__e, symnumber_2, PrimIsNumber, V1781)

if True == tmp10141 {
return True
} else {
return TailCall1(__e, SymFn(symempty_2), V1781)
}


}


}


}


panic("unreachable")
})

tmp10148 := Call2(__e, ns2_1set, symatom_2, tmp10139)


_ = tmp10148

tmp10149 := MakeNative1(func(__e *ControlFlow, V1782 Obj) Obj {
tmp10150 := Call1(__e, SymFn(symshen_4_5hterm_6), V1782)


let__9454 := tmp10150
_ = let__9454

tmp10156 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9454)


var ifres10151 Obj

if True == tmp10156 {
tmp10152 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10151 = tmp10152


} else {
tmp10153 := Call1(__e, SymFn(symshen_4_5_1out), let__9454)


let__9455 := tmp10153
_ = let__9455

tmp10154 := Call1(__e, SymFn(symshen_4in_1_6), let__9454)


let__9456 := tmp10154
_ = let__9456

tmp10155 := Call2(__e, SymFn(symshen_4comb), let__9456, let__9455)


ifres10151 = tmp10155


}

let__9453 := ifres10151
_ = let__9453

tmp10158 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9453)


if True == tmp10158 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__9453
}


panic("unreachable")
})

tmp10159 := Call2(__e, ns2_1set, symshen_4_5hterm1_6, tmp10149)


_ = tmp10159

tmp10160 := MakeNative1(func(__e *ControlFlow, V1787 Obj) Obj {
tmp10161 := Call1(__e, SymFn(symshen_4_5hterm_6), V1787)


let__9458 := tmp10161
_ = let__9458

tmp10167 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9458)


var ifres10162 Obj

if True == tmp10167 {
tmp10163 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10162 = tmp10163


} else {
tmp10164 := Call1(__e, SymFn(symshen_4_5_1out), let__9458)


let__9459 := tmp10164
_ = let__9459

tmp10165 := Call1(__e, SymFn(symshen_4in_1_6), let__9458)


let__9460 := tmp10165
_ = let__9460

tmp10166 := Call2(__e, SymFn(symshen_4comb), let__9460, let__9459)


ifres10162 = tmp10166


}

let__9457 := ifres10162
_ = let__9457

tmp10169 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9457)


if True == tmp10169 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__9457
}


panic("unreachable")
})

tmp10170 := Call2(__e, ns2_1set, symshen_4_5hterm2_6, tmp10160)


_ = tmp10170

tmp10171 := MakeNative1(func(__e *ControlFlow, V1792 Obj) Obj {
tmp10172 := Call1(__e, SymFn(symshen_4_5literal_6), V1792)


let__9462 := tmp10172
_ = let__9462

tmp10185 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9462)


var ifres10173 Obj

if True == tmp10185 {
tmp10174 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10173 = tmp10174


} else {
tmp10175 := Call1(__e, SymFn(symshen_4_5_1out), let__9462)


let__9463 := tmp10175
_ = let__9463

tmp10176 := Call1(__e, SymFn(symshen_4in_1_6), let__9462)


let__9464 := tmp10176
_ = let__9464

tmp10177 := Call1(__e, SymFn(symshen_4_5body_6), let__9464)


let__9465 := tmp10177
_ = let__9465

tmp10184 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9465)


var ifres10178 Obj

if True == tmp10184 {
tmp10179 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10178 = tmp10179


} else {
tmp10180 := Call1(__e, SymFn(symshen_4_5_1out), let__9465)


let__9466 := tmp10180
_ = let__9466

tmp10181 := Call1(__e, SymFn(symshen_4in_1_6), let__9465)


let__9467 := tmp10181
_ = let__9467

tmp10182 := GCons(__e, let__9463, let__9466)

tmp10183 := Call2(__e, SymFn(symshen_4comb), let__9467, tmp10182)


ifres10178 = tmp10183


}

ifres10173 = ifres10178


}

let__9461 := ifres10173
_ = let__9461

tmp10195 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9461)


if True == tmp10195 {
tmp10186 := Call1(__e, SymFn(sym_5e_6), V1792)


let__9469 := tmp10186
_ = let__9469

tmp10191 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9469)


var ifres10187 Obj

if True == tmp10191 {
tmp10188 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10187 = tmp10188


} else {
tmp10189 := Call1(__e, SymFn(symshen_4in_1_6), let__9469)


let__9470 := tmp10189
_ = let__9470

tmp10190 := Call2(__e, SymFn(symshen_4comb), let__9470, Nil)


ifres10187 = tmp10190


}

let__9468 := ifres10187
_ = let__9468

tmp10193 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9468)


if True == tmp10193 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__9468
}


} else {
return let__9461
}


panic("unreachable")
})

tmp10196 := Call2(__e, ns2_1set, symshen_4_5body_6, tmp10171)


_ = tmp10196

tmp10197 := MakeNative1(func(__e *ControlFlow, V1803 Obj) Obj {
tmp10202 := Call2(__e, SymFn(symshen_4hds_a_2), V1803, sym_b)


var ifres10198 Obj

if True == tmp10202 {
tmp10199 := Call1(__e, SymFn(symtail), V1803)


let__9472 := tmp10199
_ = let__9472

tmp10200 := Call2(__e, SymFn(symshen_4comb), let__9472, sym_b)


ifres10198 = tmp10200


} else {
tmp10201 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10198 = tmp10201


}

let__9471 := ifres10198
_ = let__9471

tmp10223 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9471)


if True == tmp10223 {
tmp10219 := Call1(__e, SymFn(symshen_4ccons_2), V1803)


var ifres10203 Obj

if True == tmp10219 {
tmp10204 := Call1(__e, SymFn(symhead), V1803)


let__9474 := tmp10204
_ = let__9474

tmp10205 := Call1(__e, SymFn(symtail), V1803)


let__9475 := tmp10205
_ = let__9475

tmp10206 := Call1(__e, SymFn(symshen_4_5bterms_6), let__9474)


let__9476 := tmp10206
_ = let__9476

tmp10217 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9476)


var ifres10207 Obj

if True == tmp10217 {
tmp10208 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10207 = tmp10208


} else {
tmp10209 := Call1(__e, SymFn(symshen_4_5_1out), let__9476)


let__9477 := tmp10209
_ = let__9477

tmp10210 := Call1(__e, SymFn(symshen_4in_1_6), let__9476)


let__9478 := tmp10210
_ = let__9478

tmp10211 := Call1(__e, SymFn(sym_5end_6), let__9478)


let__9479 := tmp10211
_ = let__9479

tmp10216 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9479)


var ifres10212 Obj

if True == tmp10216 {
tmp10213 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10212 = tmp10213


} else {
tmp10214 := Call1(__e, SymFn(symshen_4in_1_6), let__9479)


let__9480 := tmp10214
_ = let__9480

tmp10215 := Call2(__e, SymFn(symshen_4comb), let__9475, let__9477)


ifres10212 = tmp10215


}

ifres10207 = ifres10212


}

ifres10203 = ifres10207


} else {
tmp10218 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10203 = tmp10218


}

let__9473 := ifres10203
_ = let__9473

tmp10221 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9473)


if True == tmp10221 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__9473
}


} else {
return let__9471
}


panic("unreachable")
})

tmp10224 := Call2(__e, ns2_1set, symshen_4_5literal_6, tmp10197)


_ = tmp10224

tmp10225 := MakeNative1(func(__e *ControlFlow, V1814 Obj) Obj {
tmp10226 := Call1(__e, SymFn(symshen_4_5bterm_6), V1814)


let__9482 := tmp10226
_ = let__9482

tmp10239 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9482)


var ifres10227 Obj

if True == tmp10239 {
tmp10228 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10227 = tmp10228


} else {
tmp10229 := Call1(__e, SymFn(symshen_4_5_1out), let__9482)


let__9483 := tmp10229
_ = let__9483

tmp10230 := Call1(__e, SymFn(symshen_4in_1_6), let__9482)


let__9484 := tmp10230
_ = let__9484

tmp10231 := Call1(__e, SymFn(symshen_4_5bterms_6), let__9484)


let__9485 := tmp10231
_ = let__9485

tmp10238 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9485)


var ifres10232 Obj

if True == tmp10238 {
tmp10233 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10232 = tmp10233


} else {
tmp10234 := Call1(__e, SymFn(symshen_4_5_1out), let__9485)


let__9486 := tmp10234
_ = let__9486

tmp10235 := Call1(__e, SymFn(symshen_4in_1_6), let__9485)


let__9487 := tmp10235
_ = let__9487

tmp10236 := GCons(__e, let__9483, let__9486)

tmp10237 := Call2(__e, SymFn(symshen_4comb), let__9487, tmp10236)


ifres10232 = tmp10237


}

ifres10227 = ifres10232


}

let__9481 := ifres10227
_ = let__9481

tmp10249 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9481)


if True == tmp10249 {
tmp10240 := Call1(__e, SymFn(sym_5e_6), V1814)


let__9489 := tmp10240
_ = let__9489

tmp10245 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9489)


var ifres10241 Obj

if True == tmp10245 {
tmp10242 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10241 = tmp10242


} else {
tmp10243 := Call1(__e, SymFn(symshen_4in_1_6), let__9489)


let__9490 := tmp10243
_ = let__9490

tmp10244 := Call2(__e, SymFn(symshen_4comb), let__9490, Nil)


ifres10241 = tmp10244


}

let__9488 := ifres10241
_ = let__9488

tmp10247 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9488)


if True == tmp10247 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__9488
}


} else {
return let__9481
}


panic("unreachable")
})

tmp10250 := Call2(__e, ns2_1set, symshen_4_5bterms_6, tmp10225)


_ = tmp10250

tmp10251 := MakeNative1(func(__e *ControlFlow, V1825 Obj) Obj {
tmp10252 := Call1(__e, SymFn(symshen_4_5wildcard_6), V1825)


let__9492 := tmp10252
_ = let__9492

tmp10258 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9492)


var ifres10253 Obj

if True == tmp10258 {
tmp10254 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10253 = tmp10254


} else {
tmp10255 := Call1(__e, SymFn(symshen_4_5_1out), let__9492)


let__9493 := tmp10255
_ = let__9493

tmp10256 := Call1(__e, SymFn(symshen_4in_1_6), let__9492)


let__9494 := tmp10256
_ = let__9494

tmp10257 := Call2(__e, SymFn(symshen_4comb), let__9494, let__9493)


ifres10253 = tmp10257


}

let__9491 := ifres10253
_ = let__9491

tmp10290 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9491)


if True == tmp10290 {
tmp10267 := GConsP(__e, V1825)

var ifres10259 Obj

if True == tmp10267 {
tmp10260 := Call1(__e, SymFn(symhead), V1825)


let__9496 := tmp10260
_ = let__9496

tmp10261 := Call1(__e, SymFn(symtail), V1825)


let__9497 := tmp10261
_ = let__9497

tmp10265 := Call1(__e, SymFn(symatom_2), let__9496)


var ifres10262 Obj

if True == tmp10265 {
tmp10263 := Call2(__e, SymFn(symshen_4comb), let__9497, let__9496)


ifres10262 = tmp10263


} else {
tmp10264 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10262 = tmp10264


}

ifres10259 = ifres10262


} else {
tmp10266 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10259 = tmp10266


}

let__9495 := ifres10259
_ = let__9495

tmp10288 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9495)


if True == tmp10288 {
tmp10284 := Call1(__e, SymFn(symshen_4ccons_2), V1825)


var ifres10268 Obj

if True == tmp10284 {
tmp10269 := Call1(__e, SymFn(symhead), V1825)


let__9499 := tmp10269
_ = let__9499

tmp10270 := Call1(__e, SymFn(symtail), V1825)


let__9500 := tmp10270
_ = let__9500

tmp10271 := Call1(__e, SymFn(symshen_4_5bterms_6), let__9499)


let__9501 := tmp10271
_ = let__9501

tmp10282 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9501)


var ifres10272 Obj

if True == tmp10282 {
tmp10273 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10272 = tmp10273


} else {
tmp10274 := Call1(__e, SymFn(symshen_4_5_1out), let__9501)


let__9502 := tmp10274
_ = let__9502

tmp10275 := Call1(__e, SymFn(symshen_4in_1_6), let__9501)


let__9503 := tmp10275
_ = let__9503

tmp10276 := Call1(__e, SymFn(sym_5end_6), let__9503)


let__9504 := tmp10276
_ = let__9504

tmp10281 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9504)


var ifres10277 Obj

if True == tmp10281 {
tmp10278 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10277 = tmp10278


} else {
tmp10279 := Call1(__e, SymFn(symshen_4in_1_6), let__9504)


let__9505 := tmp10279
_ = let__9505

tmp10280 := Call2(__e, SymFn(symshen_4comb), let__9500, let__9502)


ifres10277 = tmp10280


}

ifres10272 = ifres10277


}

ifres10268 = ifres10272


} else {
tmp10283 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10268 = tmp10283


}

let__9498 := ifres10268
_ = let__9498

tmp10286 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9498)


if True == tmp10286 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__9498
}


} else {
return let__9495
}


} else {
return let__9491
}


panic("unreachable")
})

tmp10291 := Call2(__e, ns2_1set, symshen_4_5bterm_6, tmp10251)


_ = tmp10291

tmp10292 := MakeNative1(func(__e *ControlFlow, V1841 Obj) Obj {
tmp10302 := GConsP(__e, V1841)

var ifres10293 Obj

if True == tmp10302 {
tmp10294 := Call1(__e, SymFn(symhead), V1841)


let__9507 := tmp10294
_ = let__9507

tmp10295 := Call1(__e, SymFn(symtail), V1841)


let__9508 := tmp10295
_ = let__9508

tmp10300 := GEq(__e, let__9507, sym__)

var ifres10296 Obj

if True == tmp10300 {
tmp10297 := Call1(__e, SymFn(symgensym), symY)


tmp10298 := Call2(__e, SymFn(symshen_4comb), let__9508, tmp10297)


ifres10296 = tmp10298


} else {
tmp10299 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10296 = tmp10299


}

ifres10293 = ifres10296


} else {
tmp10301 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10293 = tmp10301


}

let__9506 := ifres10293
_ = let__9506

tmp10304 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9506)


if True == tmp10304 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__9506
}


panic("unreachable")
})

tmp10305 := Call2(__e, ns2_1set, symshen_4_5wildcard_6, tmp10292)


_ = tmp10305

tmp10306 := MakeNative1(func(__e *ControlFlow, V1845 Obj) Obj {
tmp10315 := GConsP(__e, V1845)

var ifres10307 Obj

if True == tmp10315 {
tmp10308 := Call1(__e, SymFn(symhead), V1845)


let__9510 := tmp10308
_ = let__9510

tmp10309 := Call1(__e, SymFn(symtail), V1845)


let__9511 := tmp10309
_ = let__9511

tmp10313 := Call1(__e, SymFn(symshen_4semicolon_2), let__9510)


var ifres10310 Obj

if True == tmp10313 {
tmp10311 := Call2(__e, SymFn(symshen_4comb), let__9511, let__9510)


ifres10310 = tmp10311


} else {
tmp10312 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10310 = tmp10312


}

ifres10307 = ifres10310


} else {
tmp10314 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10307 = tmp10314


}

let__9509 := ifres10307
_ = let__9509

tmp10317 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9509)


if True == tmp10317 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__9509
}


panic("unreachable")
})

tmp10318 := Call2(__e, ns2_1set, symshen_4_5sc_6, tmp10306)


_ = tmp10318

tmp10319 := MakeNative2(func(__e *ControlFlow, V1849, V1850 Obj) Obj {
tmp10320 := Call1(__e, SymFn(symgensym), symB)


let__9512 := tmp10320
_ = let__9512

tmp10321 := Call1(__e, SymFn(symgensym), symL)


let__9513 := tmp10321
_ = let__9513

tmp10322 := Call1(__e, SymFn(symgensym), symK)


let__9514 := tmp10322
_ = let__9514

tmp10323 := Call1(__e, SymFn(symgensym), symC)


let__9515 := tmp10323
_ = let__9515

tmp10324 := Call1(__e, SymFn(symshen_4prolog_1parameters), V1850)


let__9516 := tmp10324
_ = let__9516

tmp10325 := Call1(__e, SymFn(symshen_4hascut_2), V1850)


let__9517 := tmp10325
_ = let__9517

tmp10326 := Call(__e, SymFn(symshen_4prolog_1fbody), V1850, let__9516, let__9512, let__9513, let__9514, let__9515, let__9517)


let__9518 := tmp10326
_ = let__9518

var ifres10327 Obj

if True == let__9517 {
tmp10328 := GCons(__e, MakeInteger(1), Nil)

tmp10329 := GCons(__e, let__9514, tmp10328)

tmp10330 := GCons(__e, sym_7, tmp10329)

tmp10331 := GCons(__e, let__9518, Nil)

tmp10332 := GCons(__e, tmp10330, tmp10331)

tmp10333 := GCons(__e, let__9514, tmp10332)

tmp10334 := GCons(__e, symlet, tmp10333)

ifres10327 = tmp10334


} else {
ifres10327 = let__9518


}

let__9519 := ifres10327
_ = let__9519

tmp10335 := GCons(__e, sym_1_6, Nil)

tmp10336 := GCons(__e, let__9515, tmp10335)

tmp10337 := GCons(__e, let__9514, tmp10336)

tmp10338 := GCons(__e, let__9513, tmp10337)

tmp10339 := GCons(__e, let__9512, tmp10338)

tmp10340 := GCons(__e, let__9519, Nil)

tmp10341 := Call2(__e, SymFn(symappend), tmp10339, tmp10340)


tmp10342 := Call2(__e, SymFn(symappend), let__9516, tmp10341)


tmp10343 := GCons(__e, V1849, tmp10342)

tmp10344 := GCons(__e, symdefine, tmp10343)

let__9520 := tmp10344
_ = let__9520

return let__9520


panic("unreachable")
})

tmp10345 := Call2(__e, ns2_1set, symshen_4horn_1clause_1procedure, tmp10319)


_ = tmp10345

tmp10346 := (func() Obj {
var __self Obj
__self = MakeNative1(func(__e *ControlFlow, __self1 Obj) Obj {
__selftop:
V1862 := __self1
_ = V1862
tmp10354 := GEq(__e, sym_b, V1862)

if True == tmp10354 {
return True
} else {
tmp10352 := GConsP(__e, V1862)

if True == tmp10352 {
tmp10349 := GHd(__e, V1862)

tmp10350 := Call1(__e, SymFn(symshen_4hascut_2), tmp10349)


if True == tmp10350 {
return True
} else {
tmp10347 := GTl(__e, V1862)

if SymFn(symshen_4hascut_2) == __self {
__self1 = tmp10347
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4hascut_2), tmp10347)


}


} else {
return False
}


}


panic("unreachable")
})
return __self
})()

tmp10355 := Call2(__e, ns2_1set, symshen_4hascut_2, tmp10346)


_ = tmp10355

tmp10356 := MakeNative1(func(__e *ControlFlow, V1867 Obj) Obj {
tmp10364 := GConsP(__e, V1867)

var ifres10361 Obj

if True == tmp10364 {
tmp10362 := GHd(__e, V1867)

tmp10363 := GConsP(__e, tmp10362)

ifres10361 = tmp10363


} else {
ifres10361 = False


}

if True == ifres10361 {
tmp10357 := GHd(__e, V1867)

tmp10358 := GHd(__e, tmp10357)

tmp10359 := Call1(__e, SymFn(symlength), tmp10358)


return TailCall1(__e, SymFn(symshen_4parameters), tmp10359)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4prolog_1parameters)
}


panic("unreachable")
})

tmp10365 := Call2(__e, ns2_1set, symshen_4prolog_1parameters, tmp10356)


_ = tmp10365

tmp10366 := MakeNative(func(__e *ControlFlow) {
V1888 := __e.Get(1)
V1889 := __e.Get(2)
V1890 := __e.Get(3)
V1891 := __e.Get(4)
V1892 := __e.Get(5)
V1893 := __e.Get(6)
V1894 := __e.Get(7)
__e.ReturnDirect(func() Obj {
tmp10447 := GEq(__e, Nil, V1888)

var ifres10445 Obj

if True == tmp10447 {
tmp10446 := GEq(__e, True, V1894)

ifres10445 = tmp10446


} else {
ifres10445 = False


}

if True == ifres10445 {
tmp10367 := GCons(__e, V1892, Nil)

tmp10368 := GCons(__e, V1891, tmp10367)

return GCons(__e, symshen_4unlock, tmp10368)


} else {
tmp10443 := GConsP(__e, V1888)

var ifres10426 Obj

if True == tmp10443 {
tmp10441 := GHd(__e, V1888)

tmp10442 := GConsP(__e, tmp10441)

var ifres10427 Obj

if True == tmp10442 {
tmp10438 := GHd(__e, V1888)

tmp10439 := GTl(__e, tmp10438)

tmp10440 := GConsP(__e, tmp10439)

var ifres10428 Obj

if True == tmp10440 {
tmp10434 := GHd(__e, V1888)

tmp10435 := GTl(__e, tmp10434)

tmp10436 := GTl(__e, tmp10435)

tmp10437 := GEq(__e, Nil, tmp10436)

var ifres10429 Obj

if True == tmp10437 {
tmp10432 := GTl(__e, V1888)

tmp10433 := GEq(__e, Nil, tmp10432)

var ifres10430 Obj

if True == tmp10433 {
tmp10431 := GEq(__e, False, V1894)

ifres10430 = tmp10431


} else {
ifres10430 = False


}

ifres10429 = ifres10430


} else {
ifres10429 = False


}

ifres10428 = ifres10429


} else {
ifres10428 = False


}

ifres10427 = ifres10428


} else {
ifres10427 = False


}

ifres10426 = ifres10427


} else {
ifres10426 = False


}

if True == ifres10426 {
tmp10369 := GHd(__e, V1888)

tmp10370 := GHd(__e, tmp10369)

tmp10371 := GHd(__e, V1888)

tmp10372 := GTl(__e, tmp10371)

tmp10373 := GHd(__e, tmp10372)

tmp10374 := Call6(__e, SymFn(symshen_4continue), tmp10370, tmp10373, V1890, V1891, V1892, V1893)


let__9521 := tmp10374
_ = let__9521

tmp10375 := GCons(__e, V1891, Nil)

tmp10376 := GCons(__e, symshen_4unlocked_2, tmp10375)

tmp10377 := GHd(__e, V1888)

tmp10378 := GHd(__e, tmp10377)

tmp10379 := Call5(__e, SymFn(symshen_4compile_1head), symshen_4_7m, tmp10378, V1889, V1890, let__9521)


tmp10380 := GCons(__e, False, Nil)

tmp10381 := GCons(__e, tmp10379, tmp10380)

tmp10382 := GCons(__e, tmp10376, tmp10381)

return GCons(__e, symif, tmp10382)


} else {
tmp10424 := GConsP(__e, V1888)

var ifres10412 Obj

if True == tmp10424 {
tmp10422 := GHd(__e, V1888)

tmp10423 := GConsP(__e, tmp10422)

var ifres10413 Obj

if True == tmp10423 {
tmp10419 := GHd(__e, V1888)

tmp10420 := GTl(__e, tmp10419)

tmp10421 := GConsP(__e, tmp10420)

var ifres10414 Obj

if True == tmp10421 {
tmp10415 := GHd(__e, V1888)

tmp10416 := GTl(__e, tmp10415)

tmp10417 := GTl(__e, tmp10416)

tmp10418 := GEq(__e, Nil, tmp10417)

ifres10414 = tmp10418


} else {
ifres10414 = False


}

ifres10413 = ifres10414


} else {
ifres10413 = False


}

ifres10412 = ifres10413


} else {
ifres10412 = False


}

if True == ifres10412 {
tmp10383 := Call1(__e, SymFn(symgensym), symC)


let__9522 := tmp10383
_ = let__9522

tmp10384 := GHd(__e, V1888)

tmp10385 := GHd(__e, tmp10384)

tmp10386 := GHd(__e, V1888)

tmp10387 := GTl(__e, tmp10386)

tmp10388 := GHd(__e, tmp10387)

tmp10389 := Call6(__e, SymFn(symshen_4continue), tmp10385, tmp10388, V1890, V1891, V1892, V1893)


let__9523 := tmp10389
_ = let__9523

tmp10390 := GCons(__e, V1891, Nil)

tmp10391 := GCons(__e, symshen_4unlocked_2, tmp10390)

tmp10392 := GHd(__e, V1888)

tmp10393 := GHd(__e, tmp10392)

tmp10394 := Call5(__e, SymFn(symshen_4compile_1head), symshen_4_7m, tmp10393, V1889, V1890, let__9523)


tmp10395 := GCons(__e, False, Nil)

tmp10396 := GCons(__e, tmp10394, tmp10395)

tmp10397 := GCons(__e, tmp10391, tmp10396)

tmp10398 := GCons(__e, symif, tmp10397)

tmp10399 := GCons(__e, False, Nil)

tmp10400 := GCons(__e, let__9522, tmp10399)

tmp10401 := GCons(__e, sym_a, tmp10400)

tmp10402 := GTl(__e, V1888)

tmp10403 := Call(__e, SymFn(symshen_4prolog_1fbody), tmp10402, V1889, V1890, V1891, V1892, V1893, V1894)


tmp10404 := GCons(__e, let__9522, Nil)

tmp10405 := GCons(__e, tmp10403, tmp10404)

tmp10406 := GCons(__e, tmp10401, tmp10405)

tmp10407 := GCons(__e, symif, tmp10406)

tmp10408 := GCons(__e, tmp10407, Nil)

tmp10409 := GCons(__e, tmp10398, tmp10408)

tmp10410 := GCons(__e, let__9522, tmp10409)

return GCons(__e, symlet, tmp10410)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.prolog-fbody"))
}


}


}


panic("unreachable")
}())
}, 7)

tmp10448 := Call2(__e, ns2_1set, symshen_4prolog_1fbody, tmp10366)


_ = tmp10448

tmp10449 := MakeNative2(func(__e *ControlFlow, V1898, V1899 Obj) Obj {
tmp10453 := Call1(__e, SymFn(symshen_4locked_2), V1898)


var ifres10451 Obj

if True == tmp10453 {
tmp10452 := Call2(__e, SymFn(symshen_4fits_2), V1899, V1898)


ifres10451 = tmp10452


} else {
ifres10451 = False


}

if True == ifres10451 {
return TailCall1(__e, SymFn(symshen_4openlock), V1898)
} else {
return False
}


panic("unreachable")
})

tmp10454 := Call2(__e, ns2_1set, symshen_4unlock, tmp10449)


_ = tmp10454

tmp10455 := MakeNative1(func(__e *ControlFlow, V1900 Obj) Obj {
tmp10456 := Call1(__e, SymFn(symshen_4unlocked_2), V1900)


return GNot(__e, tmp10456)


panic("unreachable")
})

tmp10457 := Call2(__e, ns2_1set, symshen_4locked_2, tmp10455)


_ = tmp10457

tmp10458 := MakeNative1(func(__e *ControlFlow, V1901 Obj) Obj {
return GVecGet(__e, V1901, MakeInteger(1))
panic("unreachable")
})

tmp10459 := Call2(__e, ns2_1set, symshen_4unlocked_2, tmp10458)


_ = tmp10459

tmp10460 := MakeNative1(func(__e *ControlFlow, V1902 Obj) Obj {
tmp10461 := GVecSet(__e, V1902, MakeInteger(1), True)

_ = tmp10461

return False


panic("unreachable")
})

tmp10462 := Call2(__e, ns2_1set, symshen_4openlock, tmp10460)


_ = tmp10462

tmp10463 := MakeNative2(func(__e *ControlFlow, V1903, V1904 Obj) Obj {
tmp10464 := GVecGet(__e, V1904, MakeInteger(2))

return GEq(__e, V1903, tmp10464)


panic("unreachable")
})

tmp10465 := Call2(__e, ns2_1set, symshen_4fits_2, tmp10463)


_ = tmp10465

tmp10466 := MakeNative4(func(__e *ControlFlow, V1907, V1908, V1909, V1910 Obj) Obj {
tmp10467 := Call1(__e, SymFn(symthaw), V1910)


let__9524 := tmp10467
_ = let__9524

tmp10471 := GEq(__e, let__9524, False)

var ifres10469 Obj

if True == tmp10471 {
tmp10470 := Call1(__e, SymFn(symshen_4unlocked_2), V1908)


ifres10469 = tmp10470


} else {
ifres10469 = False


}

if True == ifres10469 {
return TailCall2(__e, SymFn(symshen_4lock), V1909, V1908)
} else {
return let__9524
}


panic("unreachable")
})

tmp10472 := Call2(__e, ns2_1set, symshen_4cut, tmp10466)


_ = tmp10472

tmp10473 := MakeNative2(func(__e *ControlFlow, V1912, V1913 Obj) Obj {
tmp10474 := GVecSet(__e, V1913, MakeInteger(1), False)

let__9525 := tmp10474
_ = let__9525

tmp10475 := GVecSet(__e, V1913, MakeInteger(2), V1912)

let__9526 := tmp10475
_ = let__9526

return False


panic("unreachable")
})

tmp10476 := Call2(__e, ns2_1set, symshen_4lock, tmp10473)


_ = tmp10476

tmp10477 := MakeNative6(func(__e *ControlFlow, V1916, V1917, V1918, V1919, V1920, V1921 Obj) Obj {
tmp10478 := Call1(__e, SymFn(symshen_4extract_1vars), V1916)


let__9527 := tmp10478
_ = let__9527

tmp10479 := Call1(__e, SymFn(symshen_4extract_1free_1vars), V1917)


let__9528 := tmp10479
_ = let__9528

tmp10480 := Call2(__e, SymFn(symdifference), let__9528, let__9527)


let__9529 := tmp10480
_ = let__9529

tmp10481 := GCons(__e, symshen_4incinfs, Nil)

tmp10482 := Call5(__e, SymFn(symshen_4compile_1body), V1917, V1918, V1919, V1920, V1921)


tmp10483 := GCons(__e, tmp10482, Nil)

tmp10484 := GCons(__e, tmp10481, tmp10483)

tmp10485 := GCons(__e, symdo, tmp10484)

let__9530 := tmp10485
_ = let__9530

return TailCall3(__e, SymFn(symshen_4stpart), let__9529, let__9530, V1918)


panic("unreachable")
})

tmp10486 := Call2(__e, ns2_1set, symshen_4continue, tmp10477)


_ = tmp10486

tmp10487 := MakeNative1(func(__e *ControlFlow, V1928 Obj) Obj {
tmp10518 := GConsP(__e, V1928)

var ifres10503 Obj

if True == tmp10518 {
tmp10516 := GHd(__e, V1928)

tmp10517 := GEq(__e, symlambda, tmp10516)

var ifres10504 Obj

if True == tmp10517 {
tmp10514 := GTl(__e, V1928)

tmp10515 := GConsP(__e, tmp10514)

var ifres10505 Obj

if True == tmp10515 {
tmp10511 := GTl(__e, V1928)

tmp10512 := GTl(__e, tmp10511)

tmp10513 := GConsP(__e, tmp10512)

var ifres10506 Obj

if True == tmp10513 {
tmp10507 := GTl(__e, V1928)

tmp10508 := GTl(__e, tmp10507)

tmp10509 := GTl(__e, tmp10508)

tmp10510 := GEq(__e, Nil, tmp10509)

ifres10506 = tmp10510


} else {
ifres10506 = False


}

ifres10505 = ifres10506


} else {
ifres10505 = False


}

ifres10504 = ifres10505


} else {
ifres10504 = False


}

ifres10503 = ifres10504


} else {
ifres10503 = False


}

if True == ifres10503 {
tmp10488 := GTl(__e, V1928)

tmp10489 := GHd(__e, tmp10488)

tmp10490 := GTl(__e, V1928)

tmp10491 := GTl(__e, tmp10490)

tmp10492 := GHd(__e, tmp10491)

tmp10493 := Call1(__e, SymFn(symshen_4extract_1free_1vars), tmp10492)


return TailCall2(__e, SymFn(symremove), tmp10489, tmp10493)


} else {
tmp10501 := GConsP(__e, V1928)

if True == tmp10501 {
tmp10494 := GHd(__e, V1928)

tmp10495 := Call1(__e, SymFn(symshen_4extract_1free_1vars), tmp10494)


tmp10496 := GTl(__e, V1928)

tmp10497 := Call1(__e, SymFn(symshen_4extract_1free_1vars), tmp10496)


return TailCall2(__e, SymFn(symunion), tmp10495, tmp10497)


} else {
tmp10499 := PrimGuard1(__e, symvariable_2, PrimIsVariable, V1928)

if True == tmp10499 {
return GCons(__e, V1928, Nil)
} else {
return Nil
}


}


}


panic("unreachable")
})

tmp10519 := Call2(__e, ns2_1set, symshen_4extract_1free_1vars, tmp10487)


_ = tmp10519

tmp10520 := (func() Obj {
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
tmp10552 := GEq(__e, Nil, V1945)

if True == tmp10552 {
tmp10521 := GCons(__e, V1949, Nil)

return GCons(__e, symthaw, tmp10521)


} else {
tmp10550 := GConsP(__e, V1945)

var ifres10547 Obj

if True == tmp10550 {
tmp10548 := GHd(__e, V1945)

tmp10549 := GEq(__e, sym_b, tmp10548)

ifres10547 = tmp10549


} else {
ifres10547 = False


}

if True == ifres10547 {
tmp10522 := GCons(__e, symshen_4cut, Nil)

tmp10523 := GTl(__e, V1945)

tmp10524 := GCons(__e, tmp10522, tmp10523)

if SymFn(symshen_4compile_1body) == __self {
__self1, __self2, __self3, __self4, __self5 = tmp10524, V1946, V1947, V1948, V1949
__e.Tick()
goto __selftop
}
return TailCall5(__e, SymFn(symshen_4compile_1body), tmp10524, V1946, V1947, V1948, V1949)


} else {
tmp10545 := GConsP(__e, V1945)

var ifres10542 Obj

if True == tmp10545 {
tmp10543 := GTl(__e, V1945)

tmp10544 := GEq(__e, Nil, tmp10543)

ifres10542 = tmp10544


} else {
ifres10542 = False


}

if True == ifres10542 {
tmp10525 := GHd(__e, V1945)

tmp10526 := Call2(__e, SymFn(symshen_4deref_1calls), tmp10525, V1946)


tmp10527 := GCons(__e, V1949, Nil)

tmp10528 := GCons(__e, V1948, tmp10527)

tmp10529 := GCons(__e, V1947, tmp10528)

tmp10530 := GCons(__e, V1946, tmp10529)

return TailCall2(__e, SymFn(symappend), tmp10526, tmp10530)


} else {
tmp10540 := GConsP(__e, V1945)

if True == tmp10540 {
tmp10531 := GHd(__e, V1945)

tmp10532 := Call2(__e, SymFn(symshen_4deref_1calls), tmp10531, V1946)


let__9531 := tmp10532
_ = let__9531

tmp10533 := GTl(__e, V1945)

tmp10534 := Call5(__e, SymFn(symshen_4freeze_1literals), tmp10533, V1946, V1947, V1948, V1949)


tmp10535 := GCons(__e, tmp10534, Nil)

tmp10536 := GCons(__e, V1948, tmp10535)

tmp10537 := GCons(__e, V1947, tmp10536)

tmp10538 := GCons(__e, V1946, tmp10537)

return TailCall2(__e, SymFn(symappend), let__9531, tmp10538)


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

tmp10553 := Call2(__e, ns2_1set, symshen_4compile_1body, tmp10520)


_ = tmp10553

tmp10554 := (func() Obj {
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
tmp10576 := GEq(__e, Nil, V1967)

if True == tmp10576 {
return V1971
} else {
tmp10574 := GConsP(__e, V1967)

var ifres10571 Obj

if True == tmp10574 {
tmp10572 := GHd(__e, V1967)

tmp10573 := GEq(__e, sym_b, tmp10572)

ifres10571 = tmp10573


} else {
ifres10571 = False


}

if True == ifres10571 {
tmp10555 := GCons(__e, symshen_4cut, Nil)

tmp10556 := GTl(__e, V1967)

tmp10557 := GCons(__e, tmp10555, tmp10556)

if SymFn(symshen_4freeze_1literals) == __self {
__self1, __self2, __self3, __self4, __self5 = tmp10557, V1968, V1969, V1970, V1971
__e.Tick()
goto __selftop
}
return TailCall5(__e, SymFn(symshen_4freeze_1literals), tmp10557, V1968, V1969, V1970, V1971)


} else {
tmp10569 := GConsP(__e, V1967)

if True == tmp10569 {
tmp10558 := GHd(__e, V1967)

tmp10559 := Call2(__e, SymFn(symshen_4deref_1calls), tmp10558, V1968)


let__9532 := tmp10559
_ = let__9532

tmp10560 := GTl(__e, V1967)

tmp10561 := Call5(__e, SymFn(symshen_4freeze_1literals), tmp10560, V1968, V1969, V1970, V1971)


tmp10562 := GCons(__e, tmp10561, Nil)

tmp10563 := GCons(__e, V1970, tmp10562)

tmp10564 := GCons(__e, V1969, tmp10563)

tmp10565 := GCons(__e, V1968, tmp10564)

tmp10566 := Call2(__e, SymFn(symappend), let__9532, tmp10565)


tmp10567 := GCons(__e, tmp10566, Nil)

return GCons(__e, symfreeze, tmp10567)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.freeze-literals"))
}


}


}


panic("unreachable")
})
return __self
})()

tmp10577 := Call2(__e, ns2_1set, symshen_4freeze_1literals, tmp10554)


_ = tmp10577

tmp10578 := MakeNative2(func(__e *ControlFlow, V1977, V1978 Obj) Obj {
tmp10592 := GConsP(__e, V1977)

var ifres10589 Obj

if True == tmp10592 {
tmp10590 := GHd(__e, V1977)

tmp10591 := GEq(__e, symfork, tmp10590)

ifres10589 = tmp10591


} else {
ifres10589 = False


}

if True == ifres10589 {
tmp10579 := GTl(__e, V1977)

tmp10580 := Call2(__e, SymFn(symshen_4deref_1forked_1literals), tmp10579, V1978)


tmp10581 := GCons(__e, tmp10580, Nil)

return GCons(__e, symfork, tmp10581)


} else {
tmp10587 := GConsP(__e, V1977)

if True == tmp10587 {
tmp10582 := GHd(__e, V1977)

tmp10583 := MakeNative1(func(__e *ControlFlow, Z1979 Obj) Obj {
return TailCall2(__e, SymFn(symshen_4function_1calls), Z1979, V1978)
panic("unreachable")
})

tmp10584 := GTl(__e, V1977)

tmp10585 := Call2(__e, SymFn(symmap), tmp10583, tmp10584)


return GCons(__e, tmp10582, tmp10585)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.deref-calls"))
}


}


panic("unreachable")
})

tmp10593 := Call2(__e, ns2_1set, symshen_4deref_1calls, tmp10578)


_ = tmp10593

tmp10594 := MakeNative2(func(__e *ControlFlow, V1986, V1987 Obj) Obj {
tmp10604 := GEq(__e, Nil, V1986)

if True == tmp10604 {
return Nil
} else {
tmp10602 := GConsP(__e, V1986)

if True == tmp10602 {
tmp10595 := GHd(__e, V1986)

tmp10596 := Call2(__e, SymFn(symshen_4deref_1calls), tmp10595, V1987)


tmp10597 := GTl(__e, V1986)

tmp10598 := Call2(__e, SymFn(symshen_4deref_1forked_1literals), tmp10597, V1987)


tmp10599 := GCons(__e, tmp10598, Nil)

tmp10600 := GCons(__e, tmp10596, tmp10599)

return GCons(__e, symcons, tmp10600)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("fork requires a list of literals\n"))
}


}


panic("unreachable")
})

tmp10605 := Call2(__e, ns2_1set, symshen_4deref_1forked_1literals, tmp10594)


_ = tmp10605

tmp10606 := MakeNative2(func(__e *ControlFlow, V1990, V1991 Obj) Obj {
tmp10634 := GConsP(__e, V1990)

var ifres10619 Obj

if True == tmp10634 {
tmp10632 := GHd(__e, V1990)

tmp10633 := GEq(__e, symcons, tmp10632)

var ifres10620 Obj

if True == tmp10633 {
tmp10630 := GTl(__e, V1990)

tmp10631 := GConsP(__e, tmp10630)

var ifres10621 Obj

if True == tmp10631 {
tmp10627 := GTl(__e, V1990)

tmp10628 := GTl(__e, tmp10627)

tmp10629 := GConsP(__e, tmp10628)

var ifres10622 Obj

if True == tmp10629 {
tmp10623 := GTl(__e, V1990)

tmp10624 := GTl(__e, tmp10623)

tmp10625 := GTl(__e, tmp10624)

tmp10626 := GEq(__e, Nil, tmp10625)

ifres10622 = tmp10626


} else {
ifres10622 = False


}

ifres10621 = ifres10622


} else {
ifres10621 = False


}

ifres10620 = ifres10621


} else {
ifres10620 = False


}

ifres10619 = ifres10620


} else {
ifres10619 = False


}

if True == ifres10619 {
tmp10607 := GTl(__e, V1990)

tmp10608 := GHd(__e, tmp10607)

tmp10609 := Call2(__e, SymFn(symshen_4function_1calls), tmp10608, V1991)


tmp10610 := GTl(__e, V1990)

tmp10611 := GTl(__e, tmp10610)

tmp10612 := GHd(__e, tmp10611)

tmp10613 := Call2(__e, SymFn(symshen_4function_1calls), tmp10612, V1991)


tmp10614 := GCons(__e, tmp10613, Nil)

tmp10615 := GCons(__e, tmp10609, tmp10614)

return GCons(__e, symcons, tmp10615)


} else {
tmp10617 := GConsP(__e, V1990)

if True == tmp10617 {
return TailCall3(__e, SymFn(symshen_4deref_1terms), V1990, V1991, Nil)
} else {
return V1990
}


}


panic("unreachable")
})

tmp10635 := Call2(__e, ns2_1set, symshen_4function_1calls, tmp10606)


_ = tmp10635

tmp10636 := MakeNative3(func(__e *ControlFlow, V2000, V2001, V2002 Obj) Obj {
tmp10719 := GConsP(__e, V2000)

var ifres10709 Obj

if True == tmp10719 {
tmp10717 := GHd(__e, V2000)

tmp10718 := GEq(__e, MakeInteger(0), tmp10717)

var ifres10710 Obj

if True == tmp10718 {
tmp10715 := GTl(__e, V2000)

tmp10716 := GConsP(__e, tmp10715)

var ifres10711 Obj

if True == tmp10716 {
tmp10712 := GTl(__e, V2000)

tmp10713 := GTl(__e, tmp10712)

tmp10714 := GEq(__e, Nil, tmp10713)

ifres10711 = tmp10714


} else {
ifres10711 = False


}

ifres10710 = ifres10711


} else {
ifres10710 = False


}

ifres10709 = ifres10710


} else {
ifres10709 = False


}

if True == ifres10709 {
tmp10643 := GTl(__e, V2000)

tmp10644 := GHd(__e, tmp10643)

tmp10645 := PrimGuard1(__e, symvariable_2, PrimIsVariable, tmp10644)

if True == tmp10645 {
tmp10637 := GTl(__e, V2000)

return GHd(__e, tmp10637)


} else {
tmp10638 := GTl(__e, V2000)

tmp10639 := GHd(__e, tmp10638)

tmp10640 := Call3(__e, SymFn(symshen_4app), tmp10639, MakeString("\n"), symshen_4s)


return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("attempt to optimise a non-variable "), tmp10640))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("attempt to optimise a non-variable "), tmp10640)
return Call1(__e, SymFn(symsimple_1error), __typedArg0)
})()


}


} else {
tmp10707 := GConsP(__e, V2000)

var ifres10697 Obj

if True == tmp10707 {
tmp10705 := GHd(__e, V2000)

tmp10706 := GEq(__e, MakeInteger(1), tmp10705)

var ifres10698 Obj

if True == tmp10706 {
tmp10703 := GTl(__e, V2000)

tmp10704 := GConsP(__e, tmp10703)

var ifres10699 Obj

if True == tmp10704 {
tmp10700 := GTl(__e, V2000)

tmp10701 := GTl(__e, tmp10700)

tmp10702 := GEq(__e, Nil, tmp10701)

ifres10699 = tmp10702


} else {
ifres10699 = False


}

ifres10698 = ifres10699


} else {
ifres10698 = False


}

ifres10697 = ifres10698


} else {
ifres10697 = False


}

if True == ifres10697 {
tmp10655 := GTl(__e, V2000)

tmp10656 := GHd(__e, tmp10655)

tmp10657 := PrimGuard1(__e, symvariable_2, PrimIsVariable, tmp10656)

if True == tmp10657 {
tmp10646 := GTl(__e, V2000)

tmp10647 := GHd(__e, tmp10646)

tmp10648 := GCons(__e, V2001, Nil)

tmp10649 := GCons(__e, tmp10647, tmp10648)

return GCons(__e, symshen_4lazyderef, tmp10649)


} else {
tmp10650 := GTl(__e, V2000)

tmp10651 := GHd(__e, tmp10650)

tmp10652 := Call3(__e, SymFn(symshen_4app), tmp10651, MakeString("\n"), symshen_4s)


return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("attempt to optimise a non-variable "), tmp10652))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("attempt to optimise a non-variable "), tmp10652)
return Call1(__e, SymFn(symsimple_1error), __typedArg0)
})()


}


} else {
tmp10694 := Call2(__e, SymFn(symelement_2), V2000, V2002)


tmp10695 := GNot(__e, tmp10694)

var ifres10692 Obj

if True == tmp10695 {
tmp10693 := PrimGuard1(__e, symvariable_2, PrimIsVariable, V2000)

ifres10692 = tmp10693


} else {
ifres10692 = False


}

if True == ifres10692 {
tmp10658 := GCons(__e, V2001, Nil)

tmp10659 := GCons(__e, V2000, tmp10658)

return GCons(__e, symshen_4deref, tmp10659)


} else {
tmp10690 := GConsP(__e, V2000)

var ifres10675 Obj

if True == tmp10690 {
tmp10688 := GHd(__e, V2000)

tmp10689 := GEq(__e, symlambda, tmp10688)

var ifres10676 Obj

if True == tmp10689 {
tmp10686 := GTl(__e, V2000)

tmp10687 := GConsP(__e, tmp10686)

var ifres10677 Obj

if True == tmp10687 {
tmp10683 := GTl(__e, V2000)

tmp10684 := GTl(__e, tmp10683)

tmp10685 := GConsP(__e, tmp10684)

var ifres10678 Obj

if True == tmp10685 {
tmp10679 := GTl(__e, V2000)

tmp10680 := GTl(__e, tmp10679)

tmp10681 := GTl(__e, tmp10680)

tmp10682 := GEq(__e, Nil, tmp10681)

ifres10678 = tmp10682


} else {
ifres10678 = False


}

ifres10677 = ifres10678


} else {
ifres10677 = False


}

ifres10676 = ifres10677


} else {
ifres10676 = False


}

ifres10675 = ifres10676


} else {
ifres10675 = False


}

if True == ifres10675 {
tmp10660 := GTl(__e, V2000)

tmp10661 := GHd(__e, tmp10660)

tmp10662 := GTl(__e, V2000)

tmp10663 := GTl(__e, tmp10662)

tmp10664 := GHd(__e, tmp10663)

tmp10665 := GTl(__e, V2000)

tmp10666 := GHd(__e, tmp10665)

tmp10667 := GCons(__e, tmp10666, V2002)

tmp10668 := Call3(__e, SymFn(symshen_4deref_1terms), tmp10664, V2001, tmp10667)


tmp10669 := GCons(__e, tmp10668, Nil)

tmp10670 := GCons(__e, tmp10661, tmp10669)

return GCons(__e, symlambda, tmp10670)


} else {
tmp10673 := GConsP(__e, V2000)

if True == tmp10673 {
tmp10671 := MakeNative1(func(__e *ControlFlow, Z2003 Obj) Obj {
return TailCall3(__e, SymFn(symshen_4deref_1terms), Z2003, V2001, V2002)
panic("unreachable")
})

return TailCall2(__e, SymFn(symmap), tmp10671, V2000)


} else {
return V2000
}


}


}


}


}


panic("unreachable")
})

tmp10720 := Call2(__e, ns2_1set, symshen_4deref_1terms, tmp10636)


_ = tmp10720

tmp10721 := (func() Obj {
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
tmp10867 := GEq(__e, Nil, V2022)

var ifres10865 Obj

if True == tmp10867 {
tmp10866 := GEq(__e, Nil, V2023)

ifres10865 = tmp10866


} else {
ifres10865 = False


}

if True == ifres10865 {
return V2025
} else {
tmp10863 := GConsP(__e, V2022)

var ifres10847 Obj

if True == tmp10863 {
tmp10861 := GHd(__e, V2022)

tmp10862 := GConsP(__e, tmp10861)

var ifres10848 Obj

if True == tmp10862 {
tmp10858 := GHd(__e, V2022)

tmp10859 := GHd(__e, tmp10858)

tmp10860 := GEq(__e, symshen_4_7m, tmp10859)

var ifres10849 Obj

if True == tmp10860 {
tmp10855 := GHd(__e, V2022)

tmp10856 := GTl(__e, tmp10855)

tmp10857 := GConsP(__e, tmp10856)

var ifres10850 Obj

if True == tmp10857 {
tmp10851 := GHd(__e, V2022)

tmp10852 := GTl(__e, tmp10851)

tmp10853 := GTl(__e, tmp10852)

tmp10854 := GEq(__e, Nil, tmp10853)

ifres10850 = tmp10854


} else {
ifres10850 = False


}

ifres10849 = ifres10850


} else {
ifres10849 = False


}

ifres10848 = ifres10849


} else {
ifres10848 = False


}

ifres10847 = ifres10848


} else {
ifres10847 = False


}

if True == ifres10847 {
tmp10722 := GHd(__e, V2022)

tmp10723 := GTl(__e, tmp10722)

tmp10724 := GHd(__e, tmp10723)

tmp10725 := GTl(__e, V2022)

tmp10726 := GCons(__e, V2021, tmp10725)

tmp10727 := GCons(__e, tmp10724, tmp10726)

tmp10728 := GCons(__e, symshen_4_7m, tmp10727)

if SymFn(symshen_4compile_1head) == __self {
__self1, __self2, __self3, __self4, __self5 = V2021, tmp10728, V2023, V2024, V2025
__e.Tick()
goto __selftop
}
return TailCall5(__e, SymFn(symshen_4compile_1head), V2021, tmp10728, V2023, V2024, V2025)


} else {
tmp10845 := GConsP(__e, V2022)

var ifres10829 Obj

if True == tmp10845 {
tmp10843 := GHd(__e, V2022)

tmp10844 := GConsP(__e, tmp10843)

var ifres10830 Obj

if True == tmp10844 {
tmp10840 := GHd(__e, V2022)

tmp10841 := GHd(__e, tmp10840)

tmp10842 := GEq(__e, symshen_4_1m, tmp10841)

var ifres10831 Obj

if True == tmp10842 {
tmp10837 := GHd(__e, V2022)

tmp10838 := GTl(__e, tmp10837)

tmp10839 := GConsP(__e, tmp10838)

var ifres10832 Obj

if True == tmp10839 {
tmp10833 := GHd(__e, V2022)

tmp10834 := GTl(__e, tmp10833)

tmp10835 := GTl(__e, tmp10834)

tmp10836 := GEq(__e, Nil, tmp10835)

ifres10832 = tmp10836


} else {
ifres10832 = False


}

ifres10831 = ifres10832


} else {
ifres10831 = False


}

ifres10830 = ifres10831


} else {
ifres10830 = False


}

ifres10829 = ifres10830


} else {
ifres10829 = False


}

if True == ifres10829 {
tmp10729 := GHd(__e, V2022)

tmp10730 := GTl(__e, tmp10729)

tmp10731 := GHd(__e, tmp10730)

tmp10732 := GTl(__e, V2022)

tmp10733 := GCons(__e, V2021, tmp10732)

tmp10734 := GCons(__e, tmp10731, tmp10733)

tmp10735 := GCons(__e, symshen_4_1m, tmp10734)

if SymFn(symshen_4compile_1head) == __self {
__self1, __self2, __self3, __self4, __self5 = V2021, tmp10735, V2023, V2024, V2025
__e.Tick()
goto __selftop
}
return TailCall5(__e, SymFn(symshen_4compile_1head), V2021, tmp10735, V2023, V2024, V2025)


} else {
tmp10827 := GConsP(__e, V2022)

var ifres10824 Obj

if True == tmp10827 {
tmp10825 := GHd(__e, V2022)

tmp10826 := GEq(__e, symshen_4_1m, tmp10825)

ifres10824 = tmp10826


} else {
ifres10824 = False


}

if True == ifres10824 {
tmp10736 := GTl(__e, V2022)

if SymFn(symshen_4compile_1head) == __self {
__self1, __self2, __self3, __self4, __self5 = symshen_4_1m, tmp10736, V2023, V2024, V2025
__e.Tick()
goto __selftop
}
return TailCall5(__e, SymFn(symshen_4compile_1head), symshen_4_1m, tmp10736, V2023, V2024, V2025)


} else {
tmp10822 := GConsP(__e, V2022)

var ifres10819 Obj

if True == tmp10822 {
tmp10820 := GHd(__e, V2022)

tmp10821 := GEq(__e, symshen_4_7m, tmp10820)

ifres10819 = tmp10821


} else {
ifres10819 = False


}

if True == ifres10819 {
tmp10737 := GTl(__e, V2022)

if SymFn(symshen_4compile_1head) == __self {
__self1, __self2, __self3, __self4, __self5 = symshen_4_7m, tmp10737, V2023, V2024, V2025
__e.Tick()
goto __selftop
}
return TailCall5(__e, SymFn(symshen_4compile_1head), symshen_4_7m, tmp10737, V2023, V2024, V2025)


} else {
tmp10817 := GConsP(__e, V2022)

var ifres10812 Obj

if True == tmp10817 {
tmp10816 := GConsP(__e, V2023)

var ifres10813 Obj

if True == tmp10816 {
tmp10814 := GHd(__e, V2022)

tmp10815 := Call1(__e, SymFn(symshen_4wildcard_2), tmp10814)


ifres10813 = tmp10815


} else {
ifres10813 = False


}

ifres10812 = ifres10813


} else {
ifres10812 = False


}

if True == ifres10812 {
tmp10738 := GTl(__e, V2022)

tmp10739 := GTl(__e, V2023)

if SymFn(symshen_4compile_1head) == __self {
__self1, __self2, __self3, __self4, __self5 = V2021, tmp10738, tmp10739, V2024, V2025
__e.Tick()
goto __selftop
}
return TailCall5(__e, SymFn(symshen_4compile_1head), V2021, tmp10738, tmp10739, V2024, V2025)


} else {
tmp10810 := GConsP(__e, V2022)

var ifres10807 Obj

if True == tmp10810 {
tmp10808 := GHd(__e, V2022)

tmp10809 := PrimGuard1(__e, symvariable_2, PrimIsVariable, tmp10808)

ifres10807 = tmp10809


} else {
ifres10807 = False


}

if True == ifres10807 {
return TailCall5(__e, SymFn(symshen_4variable_1case), V2021, V2022, V2023, V2024, V2025)
} else {
tmp10805 := GEq(__e, symshen_4_1m, V2021)

var ifres10800 Obj

if True == tmp10805 {
tmp10804 := GConsP(__e, V2022)

var ifres10801 Obj

if True == tmp10804 {
tmp10802 := GHd(__e, V2022)

tmp10803 := Call1(__e, SymFn(symatom_2), tmp10802)


ifres10801 = tmp10803


} else {
ifres10801 = False


}

ifres10800 = ifres10801


} else {
ifres10800 = False


}

if True == ifres10800 {
return TailCall4(__e, SymFn(symshen_4atom_1case_1minus), V2022, V2023, V2024, V2025)
} else {
tmp10798 := GEq(__e, symshen_4_1m, V2021)

var ifres10774 Obj

if True == tmp10798 {
tmp10797 := GConsP(__e, V2022)

var ifres10775 Obj

if True == tmp10797 {
tmp10795 := GHd(__e, V2022)

tmp10796 := GConsP(__e, tmp10795)

var ifres10776 Obj

if True == tmp10796 {
tmp10792 := GHd(__e, V2022)

tmp10793 := GHd(__e, tmp10792)

tmp10794 := GEq(__e, symcons, tmp10793)

var ifres10777 Obj

if True == tmp10794 {
tmp10789 := GHd(__e, V2022)

tmp10790 := GTl(__e, tmp10789)

tmp10791 := GConsP(__e, tmp10790)

var ifres10778 Obj

if True == tmp10791 {
tmp10785 := GHd(__e, V2022)

tmp10786 := GTl(__e, tmp10785)

tmp10787 := GTl(__e, tmp10786)

tmp10788 := GConsP(__e, tmp10787)

var ifres10779 Obj

if True == tmp10788 {
tmp10780 := GHd(__e, V2022)

tmp10781 := GTl(__e, tmp10780)

tmp10782 := GTl(__e, tmp10781)

tmp10783 := GTl(__e, tmp10782)

tmp10784 := GEq(__e, Nil, tmp10783)

ifres10779 = tmp10784


} else {
ifres10779 = False


}

ifres10778 = ifres10779


} else {
ifres10778 = False


}

ifres10777 = ifres10778


} else {
ifres10777 = False


}

ifres10776 = ifres10777


} else {
ifres10776 = False


}

ifres10775 = ifres10776


} else {
ifres10775 = False


}

ifres10774 = ifres10775


} else {
ifres10774 = False


}

if True == ifres10774 {
return TailCall4(__e, SymFn(symshen_4cons_1case_1minus), V2022, V2023, V2024, V2025)
} else {
tmp10772 := GEq(__e, symshen_4_7m, V2021)

var ifres10767 Obj

if True == tmp10772 {
tmp10771 := GConsP(__e, V2022)

var ifres10768 Obj

if True == tmp10771 {
tmp10769 := GHd(__e, V2022)

tmp10770 := Call1(__e, SymFn(symatom_2), tmp10769)


ifres10768 = tmp10770


} else {
ifres10768 = False


}

ifres10767 = ifres10768


} else {
ifres10767 = False


}

if True == ifres10767 {
return TailCall4(__e, SymFn(symshen_4atom_1case_1plus), V2022, V2023, V2024, V2025)
} else {
tmp10765 := GEq(__e, symshen_4_7m, V2021)

var ifres10741 Obj

if True == tmp10765 {
tmp10764 := GConsP(__e, V2022)

var ifres10742 Obj

if True == tmp10764 {
tmp10762 := GHd(__e, V2022)

tmp10763 := GConsP(__e, tmp10762)

var ifres10743 Obj

if True == tmp10763 {
tmp10759 := GHd(__e, V2022)

tmp10760 := GHd(__e, tmp10759)

tmp10761 := GEq(__e, symcons, tmp10760)

var ifres10744 Obj

if True == tmp10761 {
tmp10756 := GHd(__e, V2022)

tmp10757 := GTl(__e, tmp10756)

tmp10758 := GConsP(__e, tmp10757)

var ifres10745 Obj

if True == tmp10758 {
tmp10752 := GHd(__e, V2022)

tmp10753 := GTl(__e, tmp10752)

tmp10754 := GTl(__e, tmp10753)

tmp10755 := GConsP(__e, tmp10754)

var ifres10746 Obj

if True == tmp10755 {
tmp10747 := GHd(__e, V2022)

tmp10748 := GTl(__e, tmp10747)

tmp10749 := GTl(__e, tmp10748)

tmp10750 := GTl(__e, tmp10749)

tmp10751 := GEq(__e, Nil, tmp10750)

ifres10746 = tmp10751


} else {
ifres10746 = False


}

ifres10745 = ifres10746


} else {
ifres10745 = False


}

ifres10744 = ifres10745


} else {
ifres10744 = False


}

ifres10743 = ifres10744


} else {
ifres10743 = False


}

ifres10742 = ifres10743


} else {
ifres10742 = False


}

ifres10741 = ifres10742


} else {
ifres10741 = False


}

if True == ifres10741 {
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

tmp10868 := Call2(__e, ns2_1set, symshen_4compile_1head, tmp10721)


_ = tmp10868

tmp10869 := MakeNative5(func(__e *ControlFlow, V2036, V2037, V2038, V2039, V2040 Obj) Obj {
tmp10889 := GConsP(__e, V2037)

var ifres10887 Obj

if True == tmp10889 {
tmp10888 := GConsP(__e, V2038)

ifres10887 = tmp10888


} else {
ifres10887 = False


}

if True == ifres10887 {
tmp10884 := GHd(__e, V2038)

tmp10885 := PrimGuard1(__e, symvariable_2, PrimIsVariable, tmp10884)

if True == tmp10885 {
tmp10870 := GTl(__e, V2037)

tmp10871 := GTl(__e, V2038)

tmp10872 := GHd(__e, V2038)

tmp10873 := GHd(__e, V2037)

tmp10874 := Call3(__e, SymFn(symsubst), tmp10872, tmp10873, V2040)


return TailCall5(__e, SymFn(symshen_4compile_1head), V2036, tmp10870, tmp10871, V2039, tmp10874)


} else {
tmp10875 := GHd(__e, V2037)

tmp10876 := GHd(__e, V2038)

tmp10877 := GTl(__e, V2037)

tmp10878 := GTl(__e, V2038)

tmp10879 := Call5(__e, SymFn(symshen_4compile_1head), V2036, tmp10877, tmp10878, V2039, V2040)


tmp10880 := GCons(__e, tmp10879, Nil)

tmp10881 := GCons(__e, tmp10876, tmp10880)

tmp10882 := GCons(__e, tmp10875, tmp10881)

return GCons(__e, symlet, tmp10882)


}


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.variable-case"))
}


panic("unreachable")
})

tmp10890 := Call2(__e, ns2_1set, symshen_4variable_1case, tmp10869)


_ = tmp10890

tmp10891 := MakeNative4(func(__e *ControlFlow, V2049, V2050, V2051, V2052 Obj) Obj {
tmp10914 := GConsP(__e, V2049)

var ifres10912 Obj

if True == tmp10914 {
tmp10913 := GConsP(__e, V2050)

ifres10912 = tmp10913


} else {
ifres10912 = False


}

if True == ifres10912 {
tmp10892 := Call1(__e, SymFn(symgensym), symTm)


let__9533 := tmp10892
_ = let__9533

tmp10893 := GHd(__e, V2050)

tmp10894 := GCons(__e, V2051, Nil)

tmp10895 := GCons(__e, tmp10893, tmp10894)

tmp10896 := GCons(__e, symshen_4lazyderef, tmp10895)

tmp10897 := GHd(__e, V2049)

tmp10898 := GCons(__e, tmp10897, Nil)

tmp10899 := GCons(__e, let__9533, tmp10898)

tmp10900 := GCons(__e, sym_a, tmp10899)

tmp10901 := GTl(__e, V2049)

tmp10902 := GTl(__e, V2050)

tmp10903 := Call5(__e, SymFn(symshen_4compile_1head), symshen_4_1m, tmp10901, tmp10902, V2051, V2052)


tmp10904 := GCons(__e, False, Nil)

tmp10905 := GCons(__e, tmp10903, tmp10904)

tmp10906 := GCons(__e, tmp10900, tmp10905)

tmp10907 := GCons(__e, symif, tmp10906)

tmp10908 := GCons(__e, tmp10907, Nil)

tmp10909 := GCons(__e, tmp10896, tmp10908)

tmp10910 := GCons(__e, let__9533, tmp10909)

return GCons(__e, symlet, tmp10910)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.atom-case-minus"))
}


panic("unreachable")
})

tmp10915 := Call2(__e, ns2_1set, symshen_4atom_1case_1minus, tmp10891)


_ = tmp10915

tmp10916 := MakeNative4(func(__e *ControlFlow, V2062, V2063, V2064, V2065 Obj) Obj {
tmp10974 := GConsP(__e, V2062)

var ifres10950 Obj

if True == tmp10974 {
tmp10972 := GHd(__e, V2062)

tmp10973 := GConsP(__e, tmp10972)

var ifres10951 Obj

if True == tmp10973 {
tmp10969 := GHd(__e, V2062)

tmp10970 := GHd(__e, tmp10969)

tmp10971 := GEq(__e, symcons, tmp10970)

var ifres10952 Obj

if True == tmp10971 {
tmp10966 := GHd(__e, V2062)

tmp10967 := GTl(__e, tmp10966)

tmp10968 := GConsP(__e, tmp10967)

var ifres10953 Obj

if True == tmp10968 {
tmp10962 := GHd(__e, V2062)

tmp10963 := GTl(__e, tmp10962)

tmp10964 := GTl(__e, tmp10963)

tmp10965 := GConsP(__e, tmp10964)

var ifres10954 Obj

if True == tmp10965 {
tmp10957 := GHd(__e, V2062)

tmp10958 := GTl(__e, tmp10957)

tmp10959 := GTl(__e, tmp10958)

tmp10960 := GTl(__e, tmp10959)

tmp10961 := GEq(__e, Nil, tmp10960)

var ifres10955 Obj

if True == tmp10961 {
tmp10956 := GConsP(__e, V2063)

ifres10955 = tmp10956


} else {
ifres10955 = False


}

ifres10954 = ifres10955


} else {
ifres10954 = False


}

ifres10953 = ifres10954


} else {
ifres10953 = False


}

ifres10952 = ifres10953


} else {
ifres10952 = False


}

ifres10951 = ifres10952


} else {
ifres10951 = False


}

ifres10950 = ifres10951


} else {
ifres10950 = False


}

if True == ifres10950 {
tmp10917 := Call1(__e, SymFn(symgensym), symTm)


let__9534 := tmp10917
_ = let__9534

tmp10918 := GHd(__e, V2063)

tmp10919 := GCons(__e, V2064, Nil)

tmp10920 := GCons(__e, tmp10918, tmp10919)

tmp10921 := GCons(__e, symshen_4lazyderef, tmp10920)

tmp10922 := GCons(__e, let__9534, Nil)

tmp10923 := GCons(__e, symcons_2, tmp10922)

tmp10924 := GHd(__e, V2062)

tmp10925 := GTl(__e, tmp10924)

tmp10926 := GHd(__e, tmp10925)

tmp10927 := GHd(__e, V2062)

tmp10928 := GTl(__e, tmp10927)

tmp10929 := GTl(__e, tmp10928)

tmp10930 := GHd(__e, tmp10929)

tmp10931 := GTl(__e, V2062)

tmp10932 := GCons(__e, tmp10930, tmp10931)

tmp10933 := GCons(__e, tmp10926, tmp10932)

tmp10934 := GCons(__e, let__9534, Nil)

tmp10935 := GCons(__e, symhd, tmp10934)

tmp10936 := GCons(__e, let__9534, Nil)

tmp10937 := GCons(__e, symtl, tmp10936)

tmp10938 := GTl(__e, V2063)

tmp10939 := GCons(__e, tmp10937, tmp10938)

tmp10940 := GCons(__e, tmp10935, tmp10939)

tmp10941 := Call5(__e, SymFn(symshen_4compile_1head), symshen_4_1m, tmp10933, tmp10940, V2064, V2065)


tmp10942 := GCons(__e, False, Nil)

tmp10943 := GCons(__e, tmp10941, tmp10942)

tmp10944 := GCons(__e, tmp10923, tmp10943)

tmp10945 := GCons(__e, symif, tmp10944)

tmp10946 := GCons(__e, tmp10945, Nil)

tmp10947 := GCons(__e, tmp10921, tmp10946)

tmp10948 := GCons(__e, let__9534, tmp10947)

return GCons(__e, symlet, tmp10948)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.cons-case-minus"))
}


panic("unreachable")
})

tmp10975 := Call2(__e, ns2_1set, symshen_4cons_1case_1minus, tmp10916)


_ = tmp10975

tmp10976 := MakeNative4(func(__e *ControlFlow, V2075, V2076, V2077, V2078 Obj) Obj {
tmp11019 := GConsP(__e, V2075)

var ifres11017 Obj

if True == tmp11019 {
tmp11018 := GConsP(__e, V2076)

ifres11017 = tmp11018


} else {
ifres11017 = False


}

if True == ifres11017 {
tmp10977 := Call1(__e, SymFn(symgensym), symTm)


let__9535 := tmp10977
_ = let__9535

tmp10978 := Call1(__e, SymFn(symgensym), symGoTo)


let__9536 := tmp10978
_ = let__9536

tmp10979 := GHd(__e, V2076)

tmp10980 := GCons(__e, V2077, Nil)

tmp10981 := GCons(__e, tmp10979, tmp10980)

tmp10982 := GCons(__e, symshen_4lazyderef, tmp10981)

tmp10983 := GTl(__e, V2075)

tmp10984 := GTl(__e, V2076)

tmp10985 := Call5(__e, SymFn(symshen_4compile_1head), symshen_4_7m, tmp10983, tmp10984, V2077, V2078)


tmp10986 := GCons(__e, tmp10985, Nil)

tmp10987 := GCons(__e, symfreeze, tmp10986)

tmp10988 := GHd(__e, V2075)

tmp10989 := GCons(__e, tmp10988, Nil)

tmp10990 := GCons(__e, let__9535, tmp10989)

tmp10991 := GCons(__e, sym_a, tmp10990)

tmp10992 := GCons(__e, let__9536, Nil)

tmp10993 := GCons(__e, symthaw, tmp10992)

tmp10994 := GCons(__e, let__9535, Nil)

tmp10995 := GCons(__e, symshen_4pvar_2, tmp10994)

tmp10996 := GHd(__e, V2075)

tmp10997 := Call1(__e, SymFn(symshen_4demode), tmp10996)


tmp10998 := GCons(__e, let__9536, Nil)

tmp10999 := GCons(__e, V2077, tmp10998)

tmp11000 := GCons(__e, tmp10997, tmp10999)

tmp11001 := GCons(__e, let__9535, tmp11000)

tmp11002 := GCons(__e, symshen_4bind_b, tmp11001)

tmp11003 := GCons(__e, False, Nil)

tmp11004 := GCons(__e, tmp11002, tmp11003)

tmp11005 := GCons(__e, tmp10995, tmp11004)

tmp11006 := GCons(__e, symif, tmp11005)

tmp11007 := GCons(__e, tmp11006, Nil)

tmp11008 := GCons(__e, tmp10993, tmp11007)

tmp11009 := GCons(__e, tmp10991, tmp11008)

tmp11010 := GCons(__e, symif, tmp11009)

tmp11011 := GCons(__e, tmp11010, Nil)

tmp11012 := GCons(__e, tmp10987, tmp11011)

tmp11013 := GCons(__e, let__9536, tmp11012)

tmp11014 := GCons(__e, tmp10982, tmp11013)

tmp11015 := GCons(__e, let__9535, tmp11014)

return GCons(__e, symlet, tmp11015)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.atom-case-plus"))
}


panic("unreachable")
})

tmp11020 := Call2(__e, ns2_1set, symshen_4atom_1case_1plus, tmp10976)


_ = tmp11020

tmp11021 := MakeNative4(func(__e *ControlFlow, V2089, V2090, V2091, V2092 Obj) Obj {
tmp11106 := GConsP(__e, V2089)

var ifres11082 Obj

if True == tmp11106 {
tmp11104 := GHd(__e, V2089)

tmp11105 := GConsP(__e, tmp11104)

var ifres11083 Obj

if True == tmp11105 {
tmp11101 := GHd(__e, V2089)

tmp11102 := GHd(__e, tmp11101)

tmp11103 := GEq(__e, symcons, tmp11102)

var ifres11084 Obj

if True == tmp11103 {
tmp11098 := GHd(__e, V2089)

tmp11099 := GTl(__e, tmp11098)

tmp11100 := GConsP(__e, tmp11099)

var ifres11085 Obj

if True == tmp11100 {
tmp11094 := GHd(__e, V2089)

tmp11095 := GTl(__e, tmp11094)

tmp11096 := GTl(__e, tmp11095)

tmp11097 := GConsP(__e, tmp11096)

var ifres11086 Obj

if True == tmp11097 {
tmp11089 := GHd(__e, V2089)

tmp11090 := GTl(__e, tmp11089)

tmp11091 := GTl(__e, tmp11090)

tmp11092 := GTl(__e, tmp11091)

tmp11093 := GEq(__e, Nil, tmp11092)

var ifres11087 Obj

if True == tmp11093 {
tmp11088 := GConsP(__e, V2090)

ifres11087 = tmp11088


} else {
ifres11087 = False


}

ifres11086 = ifres11087


} else {
ifres11086 = False


}

ifres11085 = ifres11086


} else {
ifres11085 = False


}

ifres11084 = ifres11085


} else {
ifres11084 = False


}

ifres11083 = ifres11084


} else {
ifres11083 = False


}

ifres11082 = ifres11083


} else {
ifres11082 = False


}

if True == ifres11082 {
tmp11022 := Call1(__e, SymFn(symgensym), symTm)


let__9537 := tmp11022
_ = let__9537

tmp11023 := Call1(__e, SymFn(symgensym), symGoTo)


let__9538 := tmp11023
_ = let__9538

tmp11024 := GHd(__e, V2089)

tmp11025 := GTl(__e, tmp11024)

tmp11026 := GHd(__e, tmp11025)

tmp11027 := GHd(__e, V2089)

tmp11028 := GTl(__e, tmp11027)

tmp11029 := GTl(__e, tmp11028)

tmp11030 := GHd(__e, tmp11029)

tmp11031 := GCons(__e, tmp11026, tmp11030)

tmp11032 := Call1(__e, SymFn(symshen_4extract_1vars), tmp11031)


let__9539 := tmp11032
_ = let__9539

tmp11033 := GHd(__e, V2089)

tmp11034 := Call1(__e, SymFn(symshen_4tame), tmp11033)


let__9540 := tmp11034
_ = let__9540

tmp11035 := Call1(__e, SymFn(symshen_4extract_1vars), let__9540)


let__9541 := tmp11035
_ = let__9541

tmp11036 := GHd(__e, V2090)

tmp11037 := GCons(__e, V2091, Nil)

tmp11038 := GCons(__e, tmp11036, tmp11037)

tmp11039 := GCons(__e, symshen_4lazyderef, tmp11038)

tmp11040 := GTl(__e, V2089)

tmp11041 := GTl(__e, V2090)

tmp11042 := Call5(__e, SymFn(symshen_4compile_1head), symshen_4_7m, tmp11040, tmp11041, V2091, V2092)


tmp11043 := Call2(__e, SymFn(symshen_4goto), let__9539, tmp11042)


tmp11044 := GCons(__e, let__9537, Nil)

tmp11045 := GCons(__e, symcons_2, tmp11044)

tmp11046 := GHd(__e, V2089)

tmp11047 := GTl(__e, tmp11046)

tmp11048 := GCons(__e, let__9537, Nil)

tmp11049 := GCons(__e, symhd, tmp11048)

tmp11050 := GCons(__e, let__9537, Nil)

tmp11051 := GCons(__e, symtl, tmp11050)

tmp11052 := GCons(__e, tmp11051, Nil)

tmp11053 := GCons(__e, tmp11049, tmp11052)

tmp11054 := Call2(__e, SymFn(symshen_4invoke), let__9538, let__9539)


tmp11055 := Call5(__e, SymFn(symshen_4compile_1head), symshen_4_7m, tmp11047, tmp11053, V2091, tmp11054)


tmp11056 := GCons(__e, let__9537, Nil)

tmp11057 := GCons(__e, symshen_4pvar_2, tmp11056)

tmp11058 := Call1(__e, SymFn(symshen_4demode), let__9540)


tmp11059 := Call2(__e, SymFn(symshen_4invoke), let__9538, let__9539)


tmp11060 := GCons(__e, tmp11059, Nil)

tmp11061 := GCons(__e, symfreeze, tmp11060)

tmp11062 := GCons(__e, tmp11061, Nil)

tmp11063 := GCons(__e, V2091, tmp11062)

tmp11064 := GCons(__e, tmp11058, tmp11063)

tmp11065 := GCons(__e, let__9537, tmp11064)

tmp11066 := GCons(__e, symshen_4bind_b, tmp11065)

tmp11067 := Call3(__e, SymFn(symshen_4stpart), let__9541, tmp11066, V2091)


tmp11068 := GCons(__e, False, Nil)

tmp11069 := GCons(__e, tmp11067, tmp11068)

tmp11070 := GCons(__e, tmp11057, tmp11069)

tmp11071 := GCons(__e, symif, tmp11070)

tmp11072 := GCons(__e, tmp11071, Nil)

tmp11073 := GCons(__e, tmp11055, tmp11072)

tmp11074 := GCons(__e, tmp11045, tmp11073)

tmp11075 := GCons(__e, symif, tmp11074)

tmp11076 := GCons(__e, tmp11075, Nil)

tmp11077 := GCons(__e, tmp11043, tmp11076)

tmp11078 := GCons(__e, let__9538, tmp11077)

tmp11079 := GCons(__e, tmp11039, tmp11078)

tmp11080 := GCons(__e, let__9537, tmp11079)

return GCons(__e, symlet, tmp11080)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.cons-case-plus"))
}


panic("unreachable")
})

tmp11107 := Call2(__e, ns2_1set, symshen_4cons_1case_1plus, tmp11021)


_ = tmp11107

tmp11108 := (func() Obj {
var __self Obj
__self = MakeNative1(func(__e *ControlFlow, __self1 Obj) Obj {
__selftop:
V2098 := __self1
_ = V2098
tmp11139 := GConsP(__e, V2098)

var ifres11129 Obj

if True == tmp11139 {
tmp11137 := GHd(__e, V2098)

tmp11138 := GEq(__e, symshen_4_7m, tmp11137)

var ifres11130 Obj

if True == tmp11138 {
tmp11135 := GTl(__e, V2098)

tmp11136 := GConsP(__e, tmp11135)

var ifres11131 Obj

if True == tmp11136 {
tmp11132 := GTl(__e, V2098)

tmp11133 := GTl(__e, tmp11132)

tmp11134 := GEq(__e, Nil, tmp11133)

ifres11131 = tmp11134


} else {
ifres11131 = False


}

ifres11130 = ifres11131


} else {
ifres11130 = False


}

ifres11129 = ifres11130


} else {
ifres11129 = False


}

if True == ifres11129 {
tmp11109 := GTl(__e, V2098)

tmp11110 := GHd(__e, tmp11109)

if SymFn(symshen_4demode) == __self {
__self1 = tmp11110
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4demode), tmp11110)


} else {
tmp11127 := GConsP(__e, V2098)

var ifres11117 Obj

if True == tmp11127 {
tmp11125 := GHd(__e, V2098)

tmp11126 := GEq(__e, symshen_4_1m, tmp11125)

var ifres11118 Obj

if True == tmp11126 {
tmp11123 := GTl(__e, V2098)

tmp11124 := GConsP(__e, tmp11123)

var ifres11119 Obj

if True == tmp11124 {
tmp11120 := GTl(__e, V2098)

tmp11121 := GTl(__e, tmp11120)

tmp11122 := GEq(__e, Nil, tmp11121)

ifres11119 = tmp11122


} else {
ifres11119 = False


}

ifres11118 = ifres11119


} else {
ifres11118 = False


}

ifres11117 = ifres11118


} else {
ifres11117 = False


}

if True == ifres11117 {
tmp11111 := GTl(__e, V2098)

tmp11112 := GHd(__e, tmp11111)

if SymFn(symshen_4demode) == __self {
__self1 = tmp11112
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4demode), tmp11112)


} else {
tmp11115 := GConsP(__e, V2098)

if True == tmp11115 {
tmp11113 := MakeNative1(func(__e *ControlFlow, Z2099 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4demode), Z2099)
panic("unreachable")
})

return TailCall2(__e, SymFn(symmap), tmp11113, V2098)


} else {
return V2098
}


}


}


panic("unreachable")
})
return __self
})()

tmp11140 := Call2(__e, ns2_1set, symshen_4demode, tmp11108)


_ = tmp11140

tmp11141 := MakeNative1(func(__e *ControlFlow, V2100 Obj) Obj {
tmp11146 := Call1(__e, SymFn(symshen_4wildcard_2), V2100)


if True == tmp11146 {
return TailCall1(__e, SymFn(symgensym), symY)
} else {
tmp11144 := GConsP(__e, V2100)

if True == tmp11144 {
tmp11142 := MakeNative1(func(__e *ControlFlow, Z2101 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4tame), Z2101)
panic("unreachable")
})

return TailCall2(__e, SymFn(symmap), tmp11142, V2100)


} else {
return V2100
}


}


panic("unreachable")
})

tmp11147 := Call2(__e, ns2_1set, symshen_4tame, tmp11141)


_ = tmp11147

tmp11148 := MakeNative2(func(__e *ControlFlow, V2102, V2103 Obj) Obj {
tmp11151 := GEq(__e, Nil, V2102)

if True == tmp11151 {
tmp11149 := GCons(__e, V2103, Nil)

return GCons(__e, symfreeze, tmp11149)


} else {
return TailCall2(__e, SymFn(symshen_4goto_1h), V2102, V2103)
}


panic("unreachable")
})

tmp11152 := Call2(__e, ns2_1set, symshen_4goto, tmp11148)


_ = tmp11152

tmp11153 := MakeNative2(func(__e *ControlFlow, V2104, V2105 Obj) Obj {
tmp11162 := GEq(__e, Nil, V2104)

if True == tmp11162 {
return V2105
} else {
tmp11160 := GConsP(__e, V2104)

if True == tmp11160 {
tmp11154 := GHd(__e, V2104)

tmp11155 := GTl(__e, V2104)

tmp11156 := Call2(__e, SymFn(symshen_4goto_1h), tmp11155, V2105)


tmp11157 := GCons(__e, tmp11156, Nil)

tmp11158 := GCons(__e, tmp11154, tmp11157)

return GCons(__e, symlambda, tmp11158)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4goto_1h)
}


}


panic("unreachable")
})

tmp11163 := Call2(__e, ns2_1set, symshen_4goto_1h, tmp11153)


_ = tmp11163

tmp11164 := MakeNative2(func(__e *ControlFlow, V2106, V2107 Obj) Obj {
tmp11167 := GEq(__e, Nil, V2107)

if True == tmp11167 {
tmp11165 := GCons(__e, V2106, Nil)

return GCons(__e, symthaw, tmp11165)


} else {
return GCons(__e, V2106, V2107)
}


panic("unreachable")
})

tmp11168 := Call2(__e, ns2_1set, symshen_4invoke, tmp11164)


_ = tmp11168

tmp11169 := MakeNative1(func(__e *ControlFlow, V2108 Obj) Obj {
return GEq(__e, V2108, sym__)
panic("unreachable")
})

tmp11170 := Call2(__e, ns2_1set, symshen_4wildcard_2, tmp11169)


_ = tmp11170

tmp11171 := MakeNative1(func(__e *ControlFlow, V2109 Obj) Obj {
tmp11176 := GVectorP(__e, V2109)

if True == tmp11176 {
tmp11172 := MakeNative0(func(__e *ControlFlow) Obj {
return GVecGet(__e, V2109, MakeInteger(0))
panic("unreachable")
})

tmp11173 := MakeNative1(func(__e *ControlFlow, Z2110 Obj) Obj {
return symshen_4not_1pvar
panic("unreachable")
})

tmp11174 := Call2(__e, try_1catch, tmp11172, tmp11173)


return GEq(__e, tmp11174, symshen_4pvar)


} else {
return False
}


panic("unreachable")
})

tmp11177 := Call2(__e, ns2_1set, symshen_4pvar_2, tmp11171)


_ = tmp11177

tmp11178 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V2111 := __self1
_ = V2111
V2112 := __self2
_ = V2112
tmp11184 := Call1(__e, SymFn(symshen_4pvar_2), V2111)


if True == tmp11184 {
tmp11179 := GVecGet(__e, V2111, MakeInteger(1))

tmp11180 := GVecGet(__e, V2112, tmp11179)

let__9542 := tmp11180
_ = let__9542

tmp11182 := GEq(__e, let__9542, symshen_4_1null_1)

if True == tmp11182 {
return V2111
} else {
if SymFn(symshen_4lazyderef) == __self {
__self1, __self2 = let__9542, V2112
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4lazyderef), let__9542, V2112)
}


} else {
return V2111
}


panic("unreachable")
})
return __self
})()

tmp11185 := Call2(__e, ns2_1set, symshen_4lazyderef, tmp11178)


_ = tmp11185

tmp11186 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V2114 := __self1
_ = V2114
V2115 := __self2
_ = V2115
tmp11198 := GConsP(__e, V2114)

if True == tmp11198 {
tmp11187 := GHd(__e, V2114)

tmp11188 := Call2(__e, SymFn(symshen_4deref), tmp11187, V2115)


tmp11189 := GTl(__e, V2114)

tmp11190 := Call2(__e, SymFn(symshen_4deref), tmp11189, V2115)


return GCons(__e, tmp11188, tmp11190)


} else {
tmp11196 := Call1(__e, SymFn(symshen_4pvar_2), V2114)


if True == tmp11196 {
tmp11191 := GVecGet(__e, V2114, MakeInteger(1))

tmp11192 := GVecGet(__e, V2115, tmp11191)

let__9543 := tmp11192
_ = let__9543

tmp11194 := GEq(__e, let__9543, symshen_4_1null_1)

if True == tmp11194 {
return V2114
} else {
if SymFn(symshen_4deref) == __self {
__self1, __self2 = let__9543, V2115
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4deref), let__9543, V2115)
}


} else {
return V2114
}


}


panic("unreachable")
})
return __self
})()

tmp11199 := Call2(__e, ns2_1set, symshen_4deref, tmp11186)


_ = tmp11199

tmp11200 := MakeNative4(func(__e *ControlFlow, V2117, V2118, V2119, V2120 Obj) Obj {
tmp11201 := Call3(__e, SymFn(symshen_4bindv), V2117, V2118, V2119)


let__9544 := tmp11201
_ = let__9544

tmp11202 := Call1(__e, SymFn(symthaw), V2120)


let__9545 := tmp11202
_ = let__9545

tmp11204 := GEq(__e, let__9545, False)

if True == tmp11204 {
return TailCall3(__e, SymFn(symshen_4unwind), V2117, V2119, let__9545)
} else {
return let__9545
}


panic("unreachable")
})

tmp11205 := Call2(__e, ns2_1set, symshen_4bind_b, tmp11200)


_ = tmp11205

tmp11206 := MakeNative3(func(__e *ControlFlow, V2123, V2124, V2125 Obj) Obj {
tmp11207 := GVecGet(__e, V2123, MakeInteger(1))

return GVecSet(__e, V2125, tmp11207, V2124)


panic("unreachable")
})

tmp11208 := Call2(__e, ns2_1set, symshen_4bindv, tmp11206)


_ = tmp11208

tmp11209 := MakeNative3(func(__e *ControlFlow, V2126, V2127, V2128 Obj) Obj {
tmp11210 := GVecGet(__e, V2126, MakeInteger(1))

tmp11211 := GVecSet(__e, V2127, tmp11210, symshen_4_1null_1)

_ = tmp11211

return V2128


panic("unreachable")
})

tmp11212 := Call2(__e, ns2_1set, symshen_4unwind, tmp11209)


_ = tmp11212

tmp11213 := MakeNative3(func(__e *ControlFlow, V2137, V2138, V2139 Obj) Obj {
tmp11228 := GEq(__e, Nil, V2137)

if True == tmp11228 {
return V2138
} else {
tmp11226 := GConsP(__e, V2137)

if True == tmp11226 {
tmp11214 := GHd(__e, V2137)

tmp11215 := GCons(__e, V2139, Nil)

tmp11216 := GCons(__e, symshen_4newpv, tmp11215)

tmp11217 := GTl(__e, V2137)

tmp11218 := Call3(__e, SymFn(symshen_4stpart), tmp11217, V2138, V2139)


tmp11219 := GCons(__e, tmp11218, Nil)

tmp11220 := GCons(__e, V2139, tmp11219)

tmp11221 := GCons(__e, symshen_4gc, tmp11220)

tmp11222 := GCons(__e, tmp11221, Nil)

tmp11223 := GCons(__e, tmp11216, tmp11222)

tmp11224 := GCons(__e, tmp11214, tmp11223)

return GCons(__e, symlet, tmp11224)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.stpart"))
}


}


panic("unreachable")
})

tmp11229 := Call2(__e, ns2_1set, symshen_4stpart, tmp11213)


_ = tmp11229

tmp11230 := MakeNative2(func(__e *ControlFlow, V2140, V2141 Obj) Obj {
tmp11234 := GEq(__e, V2141, False)

if True == tmp11234 {
tmp11231 := Call1(__e, SymFn(symshen_4ticket_1number), V2140)


let__9546 := tmp11231
_ = let__9546

tmp11232 := Call2(__e, SymFn(symshen_4decrement_1ticket), let__9546, V2140)


_ = tmp11232

return V2141


} else {
return V2141
}


panic("unreachable")
})

tmp11235 := Call2(__e, ns2_1set, symshen_4gc, tmp11230)


_ = tmp11235

tmp11236 := MakeNative2(func(__e *ControlFlow, V2143, V2144 Obj) Obj {
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

tmp11238 := Call2(__e, ns2_1set, symshen_4decrement_1ticket, tmp11236)


_ = tmp11238

tmp11239 := MakeNative1(func(__e *ControlFlow, V2145 Obj) Obj {
tmp11240 := Call1(__e, SymFn(symshen_4ticket_1number), V2145)


let__9547 := tmp11240
_ = let__9547

tmp11241 := Call1(__e, SymFn(symshen_4make_1prolog_1variable), let__9547)


let__9548 := tmp11241
_ = let__9548

tmp11242 := Call2(__e, SymFn(symshen_4nextticket), V2145, let__9547)


let__9549 := tmp11242
_ = let__9549

return let__9548


panic("unreachable")
})

tmp11243 := Call2(__e, ns2_1set, symshen_4newpv, tmp11239)


_ = tmp11243

tmp11244 := MakeNative1(func(__e *ControlFlow, V2149 Obj) Obj {
return GVecGet(__e, V2149, MakeInteger(1))
panic("unreachable")
})

tmp11245 := Call2(__e, ns2_1set, symshen_4ticket_1number, tmp11244)


_ = tmp11245

tmp11246 := MakeNative2(func(__e *ControlFlow, V2150, V2151 Obj) Obj {
tmp11247 := GVecSet(__e, V2150, V2151, symshen_4_1null_1)

let__9550 := tmp11247
_ = let__9550

return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(let__9550, MakeInteger(1), GAdd(__e, V2151, MakeInteger(1)))
}
__typedArg0 := let__9550
__typedArg1 := MakeInteger(1)
__typedArg2 := GAdd(__e, V2151, MakeInteger(1))
return Call3(__e, SymFn(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()


panic("unreachable")
})

tmp11249 := Call2(__e, ns2_1set, symshen_4nextticket, tmp11246)


_ = tmp11249

tmp11250 := MakeNative1(func(__e *ControlFlow, V2153 Obj) Obj {
tmp11251 := PrimGuard1(__e, symabsvector, PrimAbsvector, MakeInteger(2))

tmp11252 := GVecSet(__e, tmp11251, MakeInteger(0), symshen_4pvar)

return GVecSet(__e, tmp11252, MakeInteger(1), V2153)


panic("unreachable")
})

tmp11253 := Call2(__e, ns2_1set, symshen_4make_1prolog_1variable, tmp11250)


_ = tmp11253

tmp11254 := MakeNative1(func(__e *ControlFlow, V2154 Obj) Obj {
tmp11255 := GVecGet(__e, V2154, MakeInteger(1))

tmp11256 := Call3(__e, SymFn(symshen_4app), tmp11255, MakeString(""), symshen_4a)


return PrimGuard2(__e, symcn, PrimStringConcat, MakeString("Var"), tmp11256)


panic("unreachable")
})

tmp11257 := Call2(__e, ns2_1set, symshen_4pvar, tmp11254)


_ = tmp11257

tmp11258 := MakeNative0(func(__e *ControlFlow) Obj {
tmp11259 := GValue(__e, symshen_4_dinfs_d)

return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dinfs_d, GAdd(__e, MakeInteger(1), tmp11259))
}
__typedArg0 := symshen_4_dinfs_d
__typedArg1 := GAdd(__e, MakeInteger(1), tmp11259)
return Call2(__e, SymFn(symset), __typedArg0, __typedArg1)
})()


panic("unreachable")
})

tmp11261 := Call2(__e, ns2_1set, symshen_4incinfs, tmp11258)


_ = tmp11261

tmp11262 := MakeNative1(func(__e *ControlFlow, V2155 Obj) Obj {
tmp11268 := PrimGuard1(__e, syminteger_2, PrimIsInteger, V2155)

var ifres11266 Obj

if True == tmp11268 {
ifres11266 = GGt(__e, V2155, MakeInteger(0))


} else {
ifres11266 = False


}

if True == ifres11266 {
return PrimGuard2(__e, symset, PrimSet, symshen_4_dsize_1prolog_1vector_d, V2155)
} else {
tmp11263 := Call3(__e, SymFn(symshen_4app), V2155, MakeString(""), symshen_4a)


return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("prolog vector size: size should be a positive integer; not "), tmp11263))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("prolog vector size: size should be a positive integer; not "), tmp11263)
return Call1(__e, SymFn(symsimple_1error), __typedArg0)
})()


}


panic("unreachable")
})

tmp11269 := Call2(__e, ns2_1set, symshen_4prolog_1vector_1size, tmp11262)


_ = tmp11269

tmp11270 := (func() Obj {
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
tmp11297 := GEq(__e, V2167, V2168)

if True == tmp11297 {
return TailCall1(__e, SymFn(symthaw), V2170)
} else {
tmp11295 := Call1(__e, SymFn(symshen_4pvar_2), V2167)


var ifres11291 Obj

if True == tmp11295 {
tmp11292 := Call2(__e, SymFn(symshen_4deref), V2168, V2169)


tmp11293 := Call2(__e, SymFn(symshen_4occurs_1check_2), V2167, tmp11292)


ifres11291 = GNot(__e, tmp11293)


} else {
ifres11291 = False


}

if True == ifres11291 {
return TailCall4(__e, SymFn(symshen_4bind_b), V2167, V2168, V2169, V2170)
} else {
tmp11289 := Call1(__e, SymFn(symshen_4pvar_2), V2168)


var ifres11285 Obj

if True == tmp11289 {
tmp11286 := Call2(__e, SymFn(symshen_4deref), V2167, V2169)


tmp11287 := Call2(__e, SymFn(symshen_4occurs_1check_2), V2168, tmp11286)


ifres11285 = GNot(__e, tmp11287)


} else {
ifres11285 = False


}

if True == ifres11285 {
return TailCall4(__e, SymFn(symshen_4bind_b), V2168, V2167, V2169, V2170)
} else {
tmp11283 := GConsP(__e, V2167)

var ifres11281 Obj

if True == tmp11283 {
tmp11282 := GConsP(__e, V2168)

ifres11281 = tmp11282


} else {
ifres11281 = False


}

if True == ifres11281 {
tmp11271 := GHd(__e, V2167)

tmp11272 := Call2(__e, SymFn(symshen_4lazyderef), tmp11271, V2169)


tmp11273 := GHd(__e, V2168)

tmp11274 := Call2(__e, SymFn(symshen_4lazyderef), tmp11273, V2169)


tmp11275 := MakeNative0(func(__e *ControlFlow) Obj {
tmp11276 := GTl(__e, V2167)

tmp11277 := Call2(__e, SymFn(symshen_4lazyderef), tmp11276, V2169)


tmp11278 := GTl(__e, V2168)

tmp11279 := Call2(__e, SymFn(symshen_4lazyderef), tmp11278, V2169)


return TailCall4(__e, SymFn(symshen_4lzy_a_b), tmp11277, tmp11279, V2169, V2170)


panic("unreachable")
})

if SymFn(symshen_4lzy_a_b) == __self {
__self1, __self2, __self3, __self4 = tmp11272, tmp11274, V2169, tmp11275
__e.Tick()
goto __selftop
}
return TailCall4(__e, SymFn(symshen_4lzy_a_b), tmp11272, tmp11274, V2169, tmp11275)


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

tmp11298 := Call2(__e, ns2_1set, symshen_4lzy_a_b, tmp11270)


_ = tmp11298

tmp11299 := (func() Obj {
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
tmp11318 := GEq(__e, V2182, V2183)

if True == tmp11318 {
return TailCall1(__e, SymFn(symthaw), V2185)
} else {
tmp11316 := Call1(__e, SymFn(symshen_4pvar_2), V2182)


if True == tmp11316 {
return TailCall4(__e, SymFn(symshen_4bind_b), V2182, V2183, V2184, V2185)
} else {
tmp11314 := Call1(__e, SymFn(symshen_4pvar_2), V2183)


if True == tmp11314 {
return TailCall4(__e, SymFn(symshen_4bind_b), V2183, V2182, V2184, V2185)
} else {
tmp11312 := GConsP(__e, V2182)

var ifres11310 Obj

if True == tmp11312 {
tmp11311 := GConsP(__e, V2183)

ifres11310 = tmp11311


} else {
ifres11310 = False


}

if True == ifres11310 {
tmp11300 := GHd(__e, V2182)

tmp11301 := Call2(__e, SymFn(symshen_4lazyderef), tmp11300, V2184)


tmp11302 := GHd(__e, V2183)

tmp11303 := Call2(__e, SymFn(symshen_4lazyderef), tmp11302, V2184)


tmp11304 := MakeNative0(func(__e *ControlFlow) Obj {
tmp11305 := GTl(__e, V2182)

tmp11306 := Call2(__e, SymFn(symshen_4lazyderef), tmp11305, V2184)


tmp11307 := GTl(__e, V2183)

tmp11308 := Call2(__e, SymFn(symshen_4lazyderef), tmp11307, V2184)


return TailCall4(__e, SymFn(symshen_4lzy_a), tmp11306, tmp11308, V2184, V2185)


panic("unreachable")
})

if SymFn(symshen_4lzy_a) == __self {
__self1, __self2, __self3, __self4 = tmp11301, tmp11303, V2184, tmp11304
__e.Tick()
goto __selftop
}
return TailCall4(__e, SymFn(symshen_4lzy_a), tmp11301, tmp11303, V2184, tmp11304)


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

tmp11319 := Call2(__e, ns2_1set, symshen_4lzy_a, tmp11299)


_ = tmp11319

tmp11320 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V2191 := __self1
_ = V2191
V2192 := __self2
_ = V2192
tmp11328 := GEq(__e, V2191, V2192)

if True == tmp11328 {
return True
} else {
tmp11326 := GConsP(__e, V2192)

if True == tmp11326 {
tmp11323 := GHd(__e, V2192)

tmp11324 := Call2(__e, SymFn(symshen_4occurs_1check_2), V2191, tmp11323)


if True == tmp11324 {
return True
} else {
tmp11321 := GTl(__e, V2192)

if SymFn(symshen_4occurs_1check_2) == __self {
__self1, __self2 = V2191, tmp11321
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4occurs_1check_2), V2191, tmp11321)


}


} else {
return False
}


}


panic("unreachable")
})
return __self
})()

tmp11329 := Call2(__e, ns2_1set, symshen_4occurs_1check_2, tmp11320)


_ = tmp11329

tmp11330 := MakeNative5(func(__e *ControlFlow, V2193, V2194, V2195, V2196, V2197 Obj) Obj {
tmp11331 := Call1(__e, V2193, V2194)


tmp11332 := Call1(__e, tmp11331, V2195)


tmp11333 := Call1(__e, tmp11332, V2196)


return TailCall1(__e, tmp11333, V2197)


panic("unreachable")
})

tmp11334 := Call2(__e, ns2_1set, symcall, tmp11330)


_ = tmp11334

tmp11335 := MakeNative5(func(__e *ControlFlow, V2204, V2205, V2206, V2207, V2208 Obj) Obj {
return TailCall2(__e, SymFn(symshen_4deref), V2204, V2205)
panic("unreachable")
})

tmp11336 := Call2(__e, ns2_1set, symreturn, tmp11335)


_ = tmp11336

tmp11337 := MakeNative5(func(__e *ControlFlow, V2215, V2216, V2217, V2218, V2219 Obj) Obj {
if True == V2215 {
return TailCall1(__e, SymFn(symthaw), V2219)
} else {
return False
}
panic("unreachable")
})

tmp11339 := Call2(__e, ns2_1set, symwhen, tmp11337)


_ = tmp11339

tmp11340 := MakeNative6(func(__e *ControlFlow, V2220, V2221, V2222, V2223, V2224, V2225 Obj) Obj {
tmp11341 := Call2(__e, SymFn(symshen_4lazyderef), V2220, V2222)


tmp11342 := Call2(__e, SymFn(symshen_4lazyderef), V2221, V2222)


return TailCall4(__e, SymFn(symshen_4lzy_a), tmp11341, tmp11342, V2222, V2225)


panic("unreachable")
})

tmp11343 := Call2(__e, ns2_1set, symis, tmp11340)


_ = tmp11343

tmp11344 := MakeNative6(func(__e *ControlFlow, V2226, V2227, V2228, V2229, V2230, V2231 Obj) Obj {
tmp11345 := Call2(__e, SymFn(symshen_4lazyderef), V2226, V2228)


tmp11346 := Call2(__e, SymFn(symshen_4lazyderef), V2227, V2228)


return TailCall4(__e, SymFn(symshen_4lzy_a_b), tmp11345, tmp11346, V2228, V2231)


panic("unreachable")
})

tmp11347 := Call2(__e, ns2_1set, symis_b, tmp11344)


_ = tmp11347

tmp11348 := MakeNative6(func(__e *ControlFlow, V2236, V2237, V2238, V2239, V2240, V2241 Obj) Obj {
return TailCall4(__e, SymFn(symshen_4bind_b), V2236, V2237, V2238, V2241)
panic("unreachable")
})

tmp11349 := Call2(__e, ns2_1set, symbind, tmp11348)


_ = tmp11349

tmp11350 := MakeNative5(func(__e *ControlFlow, V2242, V2243, V2244, V2245, V2246 Obj) Obj {
tmp11352 := Call2(__e, SymFn(symshen_4lazyderef), V2242, V2243)


tmp11353 := Call1(__e, SymFn(symshen_4pvar_2), tmp11352)


if True == tmp11353 {
return TailCall1(__e, SymFn(symthaw), V2246)
} else {
return False
}


panic("unreachable")
})

tmp11354 := Call2(__e, ns2_1set, symvar_2, tmp11350)


_ = tmp11354

tmp11355 := MakeNative1(func(__e *ControlFlow, V2249 Obj) Obj {
return MakeString("|prolog vector|")
panic("unreachable")
})

tmp11356 := Call2(__e, ns2_1set, symshen_4print_1prolog_1vector, tmp11355)


_ = tmp11356

tmp11357 := (func() Obj {
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
tmp11369 := GEq(__e, Nil, V2268)

if True == tmp11369 {
return False
} else {
tmp11367 := GConsP(__e, V2268)

if True == tmp11367 {
tmp11358 := GHd(__e, V2268)

tmp11359 := Call1(__e, tmp11358, V2269)


tmp11360 := Call1(__e, tmp11359, V2270)


tmp11361 := Call1(__e, tmp11360, V2271)


tmp11362 := Call1(__e, tmp11361, V2272)


let__9551 := tmp11362
_ = let__9551

tmp11365 := GEq(__e, let__9551, False)

if True == tmp11365 {
tmp11363 := GTl(__e, V2268)

if SymFn(symfork) == __self {
__self1, __self2, __self3, __self4, __self5 = tmp11363, V2269, V2270, V2271, V2272
__e.Tick()
goto __selftop
}
return TailCall5(__e, SymFn(symfork), tmp11363, V2269, V2270, V2271, V2272)


} else {
return let__9551
}


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("fork expects a list of literals\n"))
}


}


panic("unreachable")
})
return __self
})()

tmp11370 := Call2(__e, ns2_1set, symfork, tmp11357)


_ = tmp11370

tmp11371 := MakeNative(func(__e *ControlFlow) {
V2274 := __e.Get(1)
V2275 := __e.Get(2)
V2276 := __e.Get(3)
V2277 := __e.Get(4)
V2278 := __e.Get(5)
V2279 := __e.Get(6)
V2280 := __e.Get(7)
__e.ReturnDirect(func() Obj {
tmp11377 := Call1(__e, SymFn(symshen_4unlocked_2), V2278)


if True == tmp11377 {
tmp11372 := Call1(__e, SymFn(symshen_4newpv), V2277)


let__9552 := tmp11372
_ = let__9552

tmp11373 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp11373

tmp11374 := MakeNative0(func(__e *ControlFlow) Obj {
__e.TailApply(SymFn(symshen_4findall_1h), V2274, V2275, V2276, let__9552, V2277, V2278, V2279, V2280)
return nil
panic("unreachable")
})

tmp11375 := Call6(__e, SymFn(symis), let__9552, Nil, V2277, V2278, V2279, tmp11374)


return TailCall2(__e, SymFn(symshen_4gc), V2277, tmp11375)


} else {
return False
}


panic("unreachable")
}())
}, 7)

tmp11378 := Call2(__e, ns2_1set, symfindall, tmp11371)


_ = tmp11378

tmp11379 := MakeNative(func(__e *ControlFlow) {
V2282 := __e.Get(1)
V2283 := __e.Get(2)
V2284 := __e.Get(3)
V2285 := __e.Get(4)
V2286 := __e.Get(5)
V2287 := __e.Get(6)
V2288 := __e.Get(7)
V2289 := __e.Get(8)
__e.ReturnDirect(func() Obj {
tmp11384 := Call1(__e, SymFn(symshen_4unlocked_2), V2287)


var ifres11380 Obj

if True == tmp11384 {
tmp11381 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp11381

tmp11382 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall6(__e, SymFn(symshen_4overbind), V2282, V2285, V2286, V2287, V2288, V2289)
panic("unreachable")
})

tmp11383 := Call5(__e, SymFn(symcall), V2283, V2286, V2287, V2288, tmp11382)


ifres11380 = tmp11383


} else {
ifres11380 = False


}

let__9553 := ifres11380
_ = let__9553

tmp11389 := GEq(__e, let__9553, False)

if True == tmp11389 {
tmp11387 := Call1(__e, SymFn(symshen_4unlocked_2), V2287)


if True == tmp11387 {
tmp11385 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp11385

return TailCall6(__e, SymFn(symis_b), V2284, V2285, V2286, V2287, V2288, V2289)


} else {
return False
}


} else {
return let__9553
}


panic("unreachable")
}())
}, 8)

tmp11390 := Call2(__e, ns2_1set, symshen_4findall_1h, tmp11379)


_ = tmp11390

tmp11391 := MakeNative6(func(__e *ControlFlow, V2297, V2298, V2299, V2300, V2301, V2302 Obj) Obj {
tmp11392 := Call2(__e, SymFn(symshen_4deref), V2297, V2299)


tmp11393 := Call2(__e, SymFn(symshen_4lazyderef), V2298, V2299)


tmp11394 := GCons(__e, tmp11392, tmp11393)

tmp11395 := Call3(__e, SymFn(symshen_4bindv), V2298, tmp11394, V2299)


_ = tmp11395

return False


panic("unreachable")
})

tmp11396 := Call2(__e, ns2_1set, symshen_4overbind, tmp11391)


_ = tmp11396

tmp11397 := MakeNative1(func(__e *ControlFlow, V2305 Obj) Obj {
tmp11401 := GEq(__e, sym_7, V2305)

if True == tmp11401 {
return PrimGuard2(__e, symset, PrimSet, symshen_4_doccurs_d, True)
} else {
tmp11399 := GEq(__e, sym_1, V2305)

if True == tmp11399 {
return PrimGuard2(__e, symset, PrimSet, symshen_4_doccurs_d, False)
} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("occurs-check expects a + or a -.\n"))
}


}


panic("unreachable")
})

return TailCall2(__e, ns2_1set, symoccurs_1check, tmp11397)




panic("unreachable")
})

