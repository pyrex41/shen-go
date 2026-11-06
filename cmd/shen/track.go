package main

import . "github.com/pyrex41/shen-go/kl"

var TrackMain = MakeNative0(func(__e *ControlFlow) Obj {
tmp13535 := MakeNative1(func(__e *ControlFlow, V5871 Obj) Obj {
tmp13536 := Call3(__e, SymFn(symshen_4app), V5871, MakeString(";\n"), symshen_4a)


tmp13537 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("partial function "), tmp13536)

tmp13538 := Call0(__e, SymFn(symstoutput))


tmp13539 := Call2(__e, SymFn(sympr), tmp13537, tmp13538)


_ = tmp13539

tmp13548 := Call1(__e, SymFn(symshen_4tracked_2), V5871)


tmp13549 := GNot(__e, tmp13548)

var ifres13543 Obj

if True == tmp13549 {
tmp13545 := Call3(__e, SymFn(symshen_4app), V5871, MakeString("? "), symshen_4a)


tmp13547 := Call1(__e, SymFn(symy_1or_1n_2), PrimGuard2(__e, symcn, PrimStringConcat, MakeString("track "), tmp13545))


var ifres13544 Obj

if True == tmp13547 {
ifres13544 = True


} else {
ifres13544 = False


}

ifres13543 = ifres13544


} else {
ifres13543 = False


}

var ifres13540 Obj

if True == ifres13543 {
tmp13541 := Call1(__e, SymFn(symps), V5871)


tmp13542 := Call1(__e, SymFn(symshen_4track_1function), tmp13541)


ifres13540 = tmp13542


} else {
ifres13540 = symshen_4ok


}

_ = ifres13540

return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("aborted"))


panic("unreachable")
})

tmp13550 := Call2(__e, ns2_1set, symshen_4f_1error, tmp13535)


_ = tmp13550

tmp13551 := MakeNative1(func(__e *ControlFlow, V5872 Obj) Obj {
tmp13552 := GValue(__e, symshen_4_dtracking_d)

return TailCall2(__e, SymFn(symelement_2), V5872, tmp13552)


panic("unreachable")
})

tmp13553 := Call2(__e, ns2_1set, symshen_4tracked_2, tmp13551)


_ = tmp13553

tmp13554 := MakeNative1(func(__e *ControlFlow, V5873 Obj) Obj {
tmp13555 := Call1(__e, SymFn(symps), V5873)


let__13524 := tmp13555
_ = let__13524

return TailCall1(__e, SymFn(symshen_4track_1function), let__13524)


panic("unreachable")
})

tmp13556 := Call2(__e, ns2_1set, symtrack, tmp13554)


_ = tmp13556

tmp13557 := MakeNative1(func(__e *ControlFlow, V5877 Obj) Obj {
tmp13611 := GConsP(__e, V5877)

var ifres13585 Obj

if True == tmp13611 {
tmp13609 := GHd(__e, V5877)

tmp13610 := GEq(__e, symdefun, tmp13609)

var ifres13587 Obj

if True == tmp13610 {
tmp13607 := GTl(__e, V5877)

tmp13608 := GConsP(__e, tmp13607)

var ifres13589 Obj

if True == tmp13608 {
tmp13604 := GTl(__e, V5877)

tmp13605 := GTl(__e, tmp13604)

tmp13606 := GConsP(__e, tmp13605)

var ifres13591 Obj

if True == tmp13606 {
tmp13600 := GTl(__e, V5877)

tmp13601 := GTl(__e, tmp13600)

tmp13602 := GTl(__e, tmp13601)

tmp13603 := GConsP(__e, tmp13602)

var ifres13593 Obj

if True == tmp13603 {
tmp13595 := GTl(__e, V5877)

tmp13596 := GTl(__e, tmp13595)

tmp13597 := GTl(__e, tmp13596)

tmp13598 := GTl(__e, tmp13597)

tmp13599 := GEq(__e, Nil, tmp13598)

var ifres13594 Obj

if True == tmp13599 {
ifres13594 = True


} else {
ifres13594 = False


}

ifres13593 = ifres13594


} else {
ifres13593 = False


}

var ifres13592 Obj

if True == ifres13593 {
ifres13592 = True


} else {
ifres13592 = False


}

ifres13591 = ifres13592


} else {
ifres13591 = False


}

var ifres13590 Obj

if True == ifres13591 {
ifres13590 = True


} else {
ifres13590 = False


}

ifres13589 = ifres13590


} else {
ifres13589 = False


}

var ifres13588 Obj

if True == ifres13589 {
ifres13588 = True


} else {
ifres13588 = False


}

ifres13587 = ifres13588


} else {
ifres13587 = False


}

var ifres13586 Obj

if True == ifres13587 {
ifres13586 = True


} else {
ifres13586 = False


}

ifres13585 = ifres13586


} else {
ifres13585 = False


}

if True == ifres13585 {
tmp13558 := GTl(__e, V5877)

tmp13559 := GHd(__e, tmp13558)

tmp13560 := GTl(__e, V5877)

tmp13561 := GTl(__e, tmp13560)

tmp13562 := GHd(__e, tmp13561)

tmp13563 := GTl(__e, V5877)

tmp13564 := GHd(__e, tmp13563)

tmp13565 := GTl(__e, V5877)

tmp13566 := GTl(__e, tmp13565)

tmp13567 := GHd(__e, tmp13566)

tmp13568 := GTl(__e, V5877)

tmp13569 := GTl(__e, tmp13568)

tmp13570 := GTl(__e, tmp13569)

tmp13571 := GHd(__e, tmp13570)

tmp13572 := Call3(__e, SymFn(symshen_4insert_1tracking_1code), tmp13564, tmp13567, tmp13571)


tmp13573 := GCons(__e, tmp13572, Nil)

tmp13574 := GCons(__e, tmp13562, tmp13573)

tmp13575 := GCons(__e, tmp13559, tmp13574)

tmp13576 := GCons(__e, symdefun, tmp13575)

let__13525 := tmp13576
_ = let__13525

tmp13577 := Call1(__e, SymFn(symeval_1kl), let__13525)


let__13526 := tmp13577
_ = let__13526

tmp13578 := GTl(__e, V5877)

tmp13579 := GHd(__e, tmp13578)

tmp13580 := GValue(__e, symshen_4_dtracking_d)

tmp13581 := Call2(__e, SymFn(symadjoin), tmp13579, tmp13580)


tmp13582 := PrimGuard2(__e, symset, PrimSet, symshen_4_dtracking_d, tmp13581)

let__13527 := tmp13582
_ = let__13527

tmp13583 := GTl(__e, V5877)

return GHd(__e, tmp13583)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.track-function"))
}


panic("unreachable")
})

tmp13612 := Call2(__e, ns2_1set, symshen_4track_1function, tmp13557)


_ = tmp13612

tmp13613 := MakeNative3(func(__e *ControlFlow, V5881, V5882, V5883 Obj) Obj {
tmp13614 := GCons(__e, symshen_4_dcall_d, Nil)

tmp13615 := GCons(__e, symvalue, tmp13614)

tmp13616 := GCons(__e, MakeInteger(1), Nil)

tmp13617 := GCons(__e, tmp13615, tmp13616)

tmp13618 := GCons(__e, sym_7, tmp13617)

tmp13619 := GCons(__e, tmp13618, Nil)

tmp13620 := GCons(__e, symshen_4_dcall_d, tmp13619)

tmp13621 := GCons(__e, symset, tmp13620)

tmp13622 := GCons(__e, symshen_4_dcall_d, Nil)

tmp13623 := GCons(__e, symvalue, tmp13622)

tmp13624 := Call2(__e, SymFn(symshen_4prolog_1track), V5883, V5882)


tmp13625 := Call1(__e, SymFn(symshen_4cons_1form), tmp13624)


tmp13626 := GCons(__e, tmp13625, Nil)

tmp13627 := GCons(__e, V5881, tmp13626)

tmp13628 := GCons(__e, tmp13623, tmp13627)

tmp13629 := GCons(__e, symshen_4input_1track, tmp13628)

tmp13630 := GCons(__e, symshen_4terpri_1or_1read_1char, Nil)

tmp13631 := GCons(__e, symshen_4_dcall_d, Nil)

tmp13632 := GCons(__e, symvalue, tmp13631)

tmp13633 := GCons(__e, symResult, Nil)

tmp13634 := GCons(__e, V5881, tmp13633)

tmp13635 := GCons(__e, tmp13632, tmp13634)

tmp13636 := GCons(__e, symshen_4output_1track, tmp13635)

tmp13637 := GCons(__e, symshen_4_dcall_d, Nil)

tmp13638 := GCons(__e, symvalue, tmp13637)

tmp13639 := GCons(__e, MakeInteger(1), Nil)

tmp13640 := GCons(__e, tmp13638, tmp13639)

tmp13641 := GCons(__e, sym_1, tmp13640)

tmp13642 := GCons(__e, tmp13641, Nil)

tmp13643 := GCons(__e, symshen_4_dcall_d, tmp13642)

tmp13644 := GCons(__e, symset, tmp13643)

tmp13645 := GCons(__e, symshen_4terpri_1or_1read_1char, Nil)

tmp13646 := GCons(__e, symResult, Nil)

tmp13647 := GCons(__e, tmp13645, tmp13646)

tmp13648 := GCons(__e, symdo, tmp13647)

tmp13649 := GCons(__e, tmp13648, Nil)

tmp13650 := GCons(__e, tmp13644, tmp13649)

tmp13651 := GCons(__e, symdo, tmp13650)

tmp13652 := GCons(__e, tmp13651, Nil)

tmp13653 := GCons(__e, tmp13636, tmp13652)

tmp13654 := GCons(__e, symdo, tmp13653)

tmp13655 := GCons(__e, tmp13654, Nil)

tmp13656 := GCons(__e, V5883, tmp13655)

tmp13657 := GCons(__e, symResult, tmp13656)

tmp13658 := GCons(__e, symlet, tmp13657)

tmp13659 := GCons(__e, tmp13658, Nil)

tmp13660 := GCons(__e, tmp13630, tmp13659)

tmp13661 := GCons(__e, symdo, tmp13660)

tmp13662 := GCons(__e, tmp13661, Nil)

tmp13663 := GCons(__e, tmp13629, tmp13662)

tmp13664 := GCons(__e, symdo, tmp13663)

tmp13665 := GCons(__e, tmp13664, Nil)

tmp13666 := GCons(__e, tmp13621, tmp13665)

return GCons(__e, symdo, tmp13666)


panic("unreachable")
})

tmp13667 := Call2(__e, ns2_1set, symshen_4insert_1tracking_1code, tmp13613)


_ = tmp13667

tmp13668 := MakeNative2(func(__e *ControlFlow, V5884, V5885 Obj) Obj {
tmp13671 := Call2(__e, SymFn(symoccurrences), symshen_4incinfs, V5884)


tmp13672 := GEq(__e, tmp13671, MakeInteger(0))

if True == tmp13672 {
return V5885
} else {
tmp13669 := Call1(__e, SymFn(symshen_4vector_1parameter), V5885)


return TailCall2(__e, SymFn(symshen_4vector_1dereference), V5885, tmp13669)


}


panic("unreachable")
})

tmp13673 := Call2(__e, ns2_1set, symshen_4prolog_1track, tmp13668)


_ = tmp13673

tmp13674 := (func() Obj {
var __self Obj
__self = MakeNative1(func(__e *ControlFlow, __self1 Obj) Obj {
__selftop:
V5888 := __self1
_ = V5888
tmp13703 := GEq(__e, Nil, V5888)

if True == tmp13703 {
return Nil
} else {
tmp13701 := GConsP(__e, V5888)

var ifres13679 Obj

if True == tmp13701 {
tmp13699 := GTl(__e, V5888)

tmp13700 := GConsP(__e, tmp13699)

var ifres13681 Obj

if True == tmp13700 {
tmp13696 := GTl(__e, V5888)

tmp13697 := GTl(__e, tmp13696)

tmp13698 := GConsP(__e, tmp13697)

var ifres13683 Obj

if True == tmp13698 {
tmp13692 := GTl(__e, V5888)

tmp13693 := GTl(__e, tmp13692)

tmp13694 := GTl(__e, tmp13693)

tmp13695 := GConsP(__e, tmp13694)

var ifres13685 Obj

if True == tmp13695 {
tmp13687 := GTl(__e, V5888)

tmp13688 := GTl(__e, tmp13687)

tmp13689 := GTl(__e, tmp13688)

tmp13690 := GTl(__e, tmp13689)

tmp13691 := GEq(__e, Nil, tmp13690)

var ifres13686 Obj

if True == tmp13691 {
ifres13686 = True


} else {
ifres13686 = False


}

ifres13685 = ifres13686


} else {
ifres13685 = False


}

var ifres13684 Obj

if True == ifres13685 {
ifres13684 = True


} else {
ifres13684 = False


}

ifres13683 = ifres13684


} else {
ifres13683 = False


}

var ifres13682 Obj

if True == ifres13683 {
ifres13682 = True


} else {
ifres13682 = False


}

ifres13681 = ifres13682


} else {
ifres13681 = False


}

var ifres13680 Obj

if True == ifres13681 {
ifres13680 = True


} else {
ifres13680 = False


}

ifres13679 = ifres13680


} else {
ifres13679 = False


}

if True == ifres13679 {
return GHd(__e, V5888)
} else {
tmp13677 := GConsP(__e, V5888)

if True == tmp13677 {
tmp13675 := GTl(__e, V5888)

if SymFn(symshen_4vector_1parameter) == __self {
__self1 = tmp13675
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4vector_1parameter), tmp13675)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4vector_1parameter)
}


}


}


panic("unreachable")
})
return __self
})()

tmp13704 := Call2(__e, ns2_1set, symshen_4vector_1parameter, tmp13674)


_ = tmp13704

tmp13705 := MakeNative2(func(__e *ControlFlow, V5891, V5892 Obj) Obj {
tmp13739 := GEq(__e, Nil, V5892)

if True == tmp13739 {
return V5891
} else {
tmp13737 := GConsP(__e, V5891)

var ifres13715 Obj

if True == tmp13737 {
tmp13735 := GTl(__e, V5891)

tmp13736 := GConsP(__e, tmp13735)

var ifres13717 Obj

if True == tmp13736 {
tmp13732 := GTl(__e, V5891)

tmp13733 := GTl(__e, tmp13732)

tmp13734 := GConsP(__e, tmp13733)

var ifres13719 Obj

if True == tmp13734 {
tmp13728 := GTl(__e, V5891)

tmp13729 := GTl(__e, tmp13728)

tmp13730 := GTl(__e, tmp13729)

tmp13731 := GConsP(__e, tmp13730)

var ifres13721 Obj

if True == tmp13731 {
tmp13723 := GTl(__e, V5891)

tmp13724 := GTl(__e, tmp13723)

tmp13725 := GTl(__e, tmp13724)

tmp13726 := GTl(__e, tmp13725)

tmp13727 := GEq(__e, Nil, tmp13726)

var ifres13722 Obj

if True == tmp13727 {
ifres13722 = True


} else {
ifres13722 = False


}

ifres13721 = ifres13722


} else {
ifres13721 = False


}

var ifres13720 Obj

if True == ifres13721 {
ifres13720 = True


} else {
ifres13720 = False


}

ifres13719 = ifres13720


} else {
ifres13719 = False


}

var ifres13718 Obj

if True == ifres13719 {
ifres13718 = True


} else {
ifres13718 = False


}

ifres13717 = ifres13718


} else {
ifres13717 = False


}

var ifres13716 Obj

if True == ifres13717 {
ifres13716 = True


} else {
ifres13716 = False


}

ifres13715 = ifres13716


} else {
ifres13715 = False


}

if True == ifres13715 {
return V5891
} else {
tmp13713 := GConsP(__e, V5891)

if True == tmp13713 {
tmp13706 := GHd(__e, V5891)

tmp13707 := GCons(__e, V5892, Nil)

tmp13708 := GCons(__e, tmp13706, tmp13707)

tmp13709 := GCons(__e, symshen_4deref, tmp13708)

tmp13710 := GTl(__e, V5891)

tmp13711 := Call2(__e, SymFn(symshen_4vector_1dereference), tmp13710, V5892)


return GCons(__e, tmp13709, tmp13711)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4vector_1dereference)
}


}


}


panic("unreachable")
})

tmp13740 := Call2(__e, ns2_1set, symshen_4vector_1dereference, tmp13705)


_ = tmp13740

tmp13741 := MakeNative1(func(__e *ControlFlow, V5895 Obj) Obj {
tmp13745 := GEq(__e, sym_7, V5895)

if True == tmp13745 {
return PrimGuard2(__e, symset, PrimSet, symshen_4_dstep_d, True)
} else {
tmp13743 := GEq(__e, sym_1, V5895)

if True == tmp13743 {
return PrimGuard2(__e, symset, PrimSet, symshen_4_dstep_d, False)
} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("step expects a + or a -.\n"))
}


}


panic("unreachable")
})

tmp13746 := Call2(__e, ns2_1set, symstep, tmp13741)


_ = tmp13746

tmp13747 := MakeNative0(func(__e *ControlFlow) Obj {
return GValue(__e, symshen_4_dstep_d)
panic("unreachable")
})

tmp13748 := Call2(__e, ns2_1set, symstep_2, tmp13747)


_ = tmp13748

tmp13749 := MakeNative1(func(__e *ControlFlow, V5898 Obj) Obj {
tmp13753 := GEq(__e, sym_7, V5898)

if True == tmp13753 {
return PrimGuard2(__e, symset, PrimSet, symshen_4_dspy_d, True)
} else {
tmp13751 := GEq(__e, sym_1, V5898)

if True == tmp13751 {
return PrimGuard2(__e, symset, PrimSet, symshen_4_dspy_d, False)
} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("spy expects a + or a -.\n"))
}


}


panic("unreachable")
})

tmp13754 := Call2(__e, ns2_1set, symspy, tmp13749)


_ = tmp13754

tmp13755 := MakeNative0(func(__e *ControlFlow) Obj {
return GValue(__e, symshen_4_dspy_d)
panic("unreachable")
})

tmp13756 := Call2(__e, ns2_1set, symspy_2, tmp13755)


_ = tmp13756

tmp13757 := MakeNative0(func(__e *ControlFlow) Obj {
tmp13761 := GValue(__e, symshen_4_dstep_d)

if True == tmp13761 {
tmp13758 := GValue(__e, sym_dstinput_d)

tmp13759 := PrimGuard1(__e, symread_1byte, PrimReadByte, tmp13758)

return TailCall1(__e, SymFn(symshen_4check_1byte), tmp13759)


} else {
return TailCall1(__e, SymFn(symnl), MakeInteger(1))
}


panic("unreachable")
})

tmp13762 := Call2(__e, ns2_1set, symshen_4terpri_1or_1read_1char, tmp13757)


_ = tmp13762

tmp13763 := MakeNative1(func(__e *ControlFlow, V5901 Obj) Obj {
tmp13765 := GEq(__e, MakeInteger(94), V5901)

if True == tmp13765 {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("aborted"))
} else {
return True
}


panic("unreachable")
})

tmp13766 := Call2(__e, ns2_1set, symshen_4check_1byte, tmp13763)


_ = tmp13766

tmp13767 := MakeNative3(func(__e *ControlFlow, V5902, V5903, V5904 Obj) Obj {
tmp13768 := Call1(__e, SymFn(symshen_4spaces), V5902)


tmp13769 := Call1(__e, SymFn(symshen_4spaces), V5902)


tmp13770 := Call3(__e, SymFn(symshen_4app), tmp13769, MakeString(""), symshen_4a)


tmp13772 := Call3(__e, SymFn(symshen_4app), V5903, PrimGuard2(__e, symcn, PrimStringConcat, MakeString(" \n"), tmp13770), symshen_4a)


tmp13774 := Call3(__e, SymFn(symshen_4app), V5902, PrimGuard2(__e, symcn, PrimStringConcat, MakeString("> Inputs to "), tmp13772), symshen_4a)


tmp13776 := Call3(__e, SymFn(symshen_4app), tmp13768, PrimGuard2(__e, symcn, PrimStringConcat, MakeString("<"), tmp13774), symshen_4a)


tmp13777 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("\n"), tmp13776)

tmp13778 := Call0(__e, SymFn(symstoutput))


tmp13779 := Call2(__e, SymFn(sympr), tmp13777, tmp13778)


_ = tmp13779

return TailCall1(__e, SymFn(symshen_4recursively_1print), V5904)


panic("unreachable")
})

tmp13780 := Call2(__e, ns2_1set, symshen_4input_1track, tmp13767)


_ = tmp13780

tmp13781 := (func() Obj {
var __self Obj
__self = MakeNative1(func(__e *ControlFlow, __self1 Obj) Obj {
__selftop:
V5907 := __self1
_ = V5907
tmp13791 := GEq(__e, Nil, V5907)

if True == tmp13791 {
tmp13782 := Call0(__e, SymFn(symstoutput))


return TailCall2(__e, SymFn(sympr), MakeString(" ==>"), tmp13782)


} else {
tmp13789 := GConsP(__e, V5907)

if True == tmp13789 {
tmp13783 := GHd(__e, V5907)

tmp13784 := Call1(__e, SymFn(symprint), tmp13783)


_ = tmp13784

tmp13785 := Call0(__e, SymFn(symstoutput))


tmp13786 := Call2(__e, SymFn(sympr), MakeString(", "), tmp13785)


_ = tmp13786

tmp13787 := GTl(__e, V5907)

if SymFn(symshen_4recursively_1print) == __self {
__self1 = tmp13787
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4recursively_1print), tmp13787)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.recursively-print"))
}


}


panic("unreachable")
})
return __self
})()

tmp13792 := Call2(__e, ns2_1set, symshen_4recursively_1print, tmp13781)


_ = tmp13792

tmp13793 := MakeNative1(func(__e *ControlFlow, V5908 Obj) Obj {
tmp13797 := GEq(__e, MakeInteger(0), V5908)

if True == tmp13797 {
return MakeString("")
} else {
tmp13795 := Call1(__e, SymFn(symshen_4spaces), GSub(__e, V5908, MakeInteger(1)))


return PrimGuard2(__e, symcn, PrimStringConcat, MakeString(" "), tmp13795)


}


panic("unreachable")
})

tmp13798 := Call2(__e, ns2_1set, symshen_4spaces, tmp13793)


_ = tmp13798

tmp13799 := MakeNative3(func(__e *ControlFlow, V5909, V5910, V5911 Obj) Obj {
tmp13800 := Call1(__e, SymFn(symshen_4spaces), V5909)


tmp13801 := Call1(__e, SymFn(symshen_4spaces), V5909)


tmp13802 := Call3(__e, SymFn(symshen_4app), V5911, MakeString(""), symshen_4s)


tmp13804 := Call3(__e, SymFn(symshen_4app), tmp13801, PrimGuard2(__e, symcn, PrimStringConcat, MakeString("==> "), tmp13802), symshen_4a)


tmp13806 := Call3(__e, SymFn(symshen_4app), V5910, PrimGuard2(__e, symcn, PrimStringConcat, MakeString(" \n"), tmp13804), symshen_4a)


tmp13808 := Call3(__e, SymFn(symshen_4app), V5909, PrimGuard2(__e, symcn, PrimStringConcat, MakeString("> Output from "), tmp13806), symshen_4a)


tmp13810 := Call3(__e, SymFn(symshen_4app), tmp13800, PrimGuard2(__e, symcn, PrimStringConcat, MakeString("<"), tmp13808), symshen_4a)


tmp13811 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("\n"), tmp13810)

tmp13812 := Call0(__e, SymFn(symstoutput))


return TailCall2(__e, SymFn(sympr), tmp13811, tmp13812)


panic("unreachable")
})

tmp13813 := Call2(__e, ns2_1set, symshen_4output_1track, tmp13799)


_ = tmp13813

tmp13814 := MakeNative1(func(__e *ControlFlow, V5912 Obj) Obj {
tmp13815 := GValue(__e, symshen_4_dtracking_d)

tmp13816 := Call2(__e, SymFn(symremove), V5912, tmp13815)


tmp13817 := PrimGuard2(__e, symset, PrimSet, symshen_4_dtracking_d, tmp13816)

_ = tmp13817

tmp13818 := MakeNative0(func(__e *ControlFlow) Obj {
tmp13819 := Call1(__e, SymFn(symps), V5912)


return TailCall1(__e, SymFn(symeval), tmp13819)


panic("unreachable")
})

tmp13820 := MakeNative1(func(__e *ControlFlow, Z5913 Obj) Obj {
return V5912
panic("unreachable")
})

tmp13821 := Call2(__e, try_1catch, tmp13818, tmp13820)


_ = tmp13821

return V5912


panic("unreachable")
})

tmp13822 := Call2(__e, ns2_1set, symuntrack, tmp13814)


_ = tmp13822

tmp13823 := MakeNative2(func(__e *ControlFlow, V5914, V5915 Obj) Obj {
return TailCall3(__e, SymFn(symshen_4remove_1h), V5914, V5915, Nil)
panic("unreachable")
})

tmp13824 := Call2(__e, ns2_1set, symremove, tmp13823)


_ = tmp13824

tmp13825 := (func() Obj {
var __self Obj
__self = MakeNative3(func(__e *ControlFlow, __self1, __self2, __self3 Obj) Obj {
__selftop:
V5925 := __self1
_ = V5925
V5926 := __self2
_ = V5926
V5927 := __self3
_ = V5927
tmp13840 := GEq(__e, Nil, V5926)

if True == tmp13840 {
return TailCall1(__e, SymFn(symreverse), V5927)
} else {
tmp13838 := GConsP(__e, V5926)

var ifres13834 Obj

if True == tmp13838 {
tmp13836 := GHd(__e, V5926)

tmp13837 := GEq(__e, V5925, tmp13836)

var ifres13835 Obj

if True == tmp13837 {
ifres13835 = True


} else {
ifres13835 = False


}

ifres13834 = ifres13835


} else {
ifres13834 = False


}

if True == ifres13834 {
tmp13826 := GHd(__e, V5926)

tmp13827 := GTl(__e, V5926)

if SymFn(symshen_4remove_1h) == __self {
__self1, __self2, __self3 = tmp13826, tmp13827, V5927
__e.Tick()
goto __selftop
}
return TailCall3(__e, SymFn(symshen_4remove_1h), tmp13826, tmp13827, V5927)


} else {
tmp13832 := GConsP(__e, V5926)

if True == tmp13832 {
tmp13828 := GTl(__e, V5926)

tmp13829 := GHd(__e, V5926)

tmp13830 := GCons(__e, tmp13829, V5927)

if SymFn(symshen_4remove_1h) == __self {
__self1, __self2, __self3 = V5925, tmp13828, tmp13830
__e.Tick()
goto __selftop
}
return TailCall3(__e, SymFn(symshen_4remove_1h), V5925, tmp13828, tmp13830)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.remove-h"))
}


}


}


panic("unreachable")
})
return __self
})()

tmp13841 := Call2(__e, ns2_1set, symshen_4remove_1h, tmp13825)


_ = tmp13841

tmp13842 := MakeNative1(func(__e *ControlFlow, V5928 Obj) Obj {
tmp13843 := GValue(__e, symshen_4_dprofiled_d)

tmp13844 := GCons(__e, V5928, tmp13843)

tmp13845 := PrimGuard2(__e, symset, PrimSet, symshen_4_dprofiled_d, tmp13844)

_ = tmp13845

tmp13846 := Call1(__e, SymFn(symps), V5928)


return TailCall1(__e, SymFn(symshen_4profile_1help), tmp13846)


panic("unreachable")
})

tmp13847 := Call2(__e, ns2_1set, symprofile, tmp13842)


_ = tmp13847

tmp13848 := MakeNative1(func(__e *ControlFlow, V5931 Obj) Obj {
tmp13913 := GConsP(__e, V5931)

var ifres13887 Obj

if True == tmp13913 {
tmp13911 := GHd(__e, V5931)

tmp13912 := GEq(__e, symdefun, tmp13911)

var ifres13889 Obj

if True == tmp13912 {
tmp13909 := GTl(__e, V5931)

tmp13910 := GConsP(__e, tmp13909)

var ifres13891 Obj

if True == tmp13910 {
tmp13906 := GTl(__e, V5931)

tmp13907 := GTl(__e, tmp13906)

tmp13908 := GConsP(__e, tmp13907)

var ifres13893 Obj

if True == tmp13908 {
tmp13902 := GTl(__e, V5931)

tmp13903 := GTl(__e, tmp13902)

tmp13904 := GTl(__e, tmp13903)

tmp13905 := GConsP(__e, tmp13904)

var ifres13895 Obj

if True == tmp13905 {
tmp13897 := GTl(__e, V5931)

tmp13898 := GTl(__e, tmp13897)

tmp13899 := GTl(__e, tmp13898)

tmp13900 := GTl(__e, tmp13899)

tmp13901 := GEq(__e, Nil, tmp13900)

var ifres13896 Obj

if True == tmp13901 {
ifres13896 = True


} else {
ifres13896 = False


}

ifres13895 = ifres13896


} else {
ifres13895 = False


}

var ifres13894 Obj

if True == ifres13895 {
ifres13894 = True


} else {
ifres13894 = False


}

ifres13893 = ifres13894


} else {
ifres13893 = False


}

var ifres13892 Obj

if True == ifres13893 {
ifres13892 = True


} else {
ifres13892 = False


}

ifres13891 = ifres13892


} else {
ifres13891 = False


}

var ifres13890 Obj

if True == ifres13891 {
ifres13890 = True


} else {
ifres13890 = False


}

ifres13889 = ifres13890


} else {
ifres13889 = False


}

var ifres13888 Obj

if True == ifres13889 {
ifres13888 = True


} else {
ifres13888 = False


}

ifres13887 = ifres13888


} else {
ifres13887 = False


}

if True == ifres13887 {
tmp13849 := Call1(__e, SymFn(symgensym), symshen_4f)


let__13528 := tmp13849
_ = let__13528

tmp13850 := GTl(__e, V5931)

tmp13851 := GHd(__e, tmp13850)

tmp13852 := GTl(__e, V5931)

tmp13853 := GTl(__e, tmp13852)

tmp13854 := GHd(__e, tmp13853)

tmp13855 := GTl(__e, V5931)

tmp13856 := GHd(__e, tmp13855)

tmp13857 := GTl(__e, V5931)

tmp13858 := GTl(__e, tmp13857)

tmp13859 := GHd(__e, tmp13858)

tmp13860 := GTl(__e, V5931)

tmp13861 := GTl(__e, tmp13860)

tmp13862 := GHd(__e, tmp13861)

tmp13863 := GCons(__e, let__13528, tmp13862)

tmp13864 := Call3(__e, SymFn(symshen_4profile_1func), tmp13856, tmp13859, tmp13863)


tmp13865 := GCons(__e, tmp13864, Nil)

tmp13866 := GCons(__e, tmp13854, tmp13865)

tmp13867 := GCons(__e, tmp13851, tmp13866)

tmp13868 := GCons(__e, symdefun, tmp13867)

let__13529 := tmp13868
_ = let__13529

tmp13869 := GTl(__e, V5931)

tmp13870 := GTl(__e, tmp13869)

tmp13871 := GHd(__e, tmp13870)

tmp13872 := GTl(__e, V5931)

tmp13873 := GHd(__e, tmp13872)

tmp13874 := GTl(__e, V5931)

tmp13875 := GTl(__e, tmp13874)

tmp13876 := GTl(__e, tmp13875)

tmp13877 := GHd(__e, tmp13876)

tmp13878 := Call3(__e, SymFn(symsubst), let__13528, tmp13873, tmp13877)


tmp13879 := GCons(__e, tmp13878, Nil)

tmp13880 := GCons(__e, tmp13871, tmp13879)

tmp13881 := GCons(__e, let__13528, tmp13880)

tmp13882 := GCons(__e, symdefun, tmp13881)

let__13530 := tmp13882
_ = let__13530

tmp13883 := Call1(__e, SymFn(symeval_1kl), let__13529)


let__13531 := tmp13883
_ = let__13531

tmp13884 := Call1(__e, SymFn(symeval_1kl), let__13530)


let__13532 := tmp13884
_ = let__13532

tmp13885 := GTl(__e, V5931)

return GHd(__e, tmp13885)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("Cannot profile.\n"))
}


panic("unreachable")
})

tmp13914 := Call2(__e, ns2_1set, symshen_4profile_1help, tmp13848)


_ = tmp13914

tmp13915 := MakeNative1(func(__e *ControlFlow, V5937 Obj) Obj {
tmp13916 := GValue(__e, symshen_4_dprofiled_d)

tmp13917 := Call2(__e, SymFn(symremove), V5937, tmp13916)


tmp13918 := PrimGuard2(__e, symset, PrimSet, symshen_4_dprofiled_d, tmp13917)

_ = tmp13918

tmp13919 := MakeNative0(func(__e *ControlFlow) Obj {
tmp13920 := Call1(__e, SymFn(symps), V5937)


return TailCall1(__e, SymFn(symeval), tmp13920)


panic("unreachable")
})

tmp13921 := MakeNative1(func(__e *ControlFlow, Z5938 Obj) Obj {
return V5937
panic("unreachable")
})

return TailCall2(__e, try_1catch, tmp13919, tmp13921)


panic("unreachable")
})

tmp13922 := Call2(__e, ns2_1set, symunprofile, tmp13915)


_ = tmp13922

tmp13923 := MakeNative1(func(__e *ControlFlow, V5939 Obj) Obj {
tmp13924 := GValue(__e, symshen_4_dprofiled_d)

return TailCall2(__e, SymFn(symelement_2), V5939, tmp13924)


panic("unreachable")
})

tmp13925 := Call2(__e, ns2_1set, symshen_4profiled_2, tmp13923)


_ = tmp13925

tmp13926 := MakeNative3(func(__e *ControlFlow, V5940, V5941, V5942 Obj) Obj {
tmp13927 := GCons(__e, symrun, Nil)

tmp13928 := GCons(__e, symget_1time, tmp13927)

tmp13929 := GCons(__e, symrun, Nil)

tmp13930 := GCons(__e, symget_1time, tmp13929)

tmp13931 := GCons(__e, symStart, Nil)

tmp13932 := GCons(__e, tmp13930, tmp13931)

tmp13933 := GCons(__e, sym_1, tmp13932)

tmp13934 := GCons(__e, V5940, Nil)

tmp13935 := GCons(__e, symshen_4get_1profile, tmp13934)

tmp13936 := GCons(__e, symFinish, Nil)

tmp13937 := GCons(__e, tmp13935, tmp13936)

tmp13938 := GCons(__e, sym_7, tmp13937)

tmp13939 := GCons(__e, tmp13938, Nil)

tmp13940 := GCons(__e, V5940, tmp13939)

tmp13941 := GCons(__e, symshen_4put_1profile, tmp13940)

tmp13942 := GCons(__e, symResult, Nil)

tmp13943 := GCons(__e, tmp13941, tmp13942)

tmp13944 := GCons(__e, symRecord, tmp13943)

tmp13945 := GCons(__e, symlet, tmp13944)

tmp13946 := GCons(__e, tmp13945, Nil)

tmp13947 := GCons(__e, tmp13933, tmp13946)

tmp13948 := GCons(__e, symFinish, tmp13947)

tmp13949 := GCons(__e, symlet, tmp13948)

tmp13950 := GCons(__e, tmp13949, Nil)

tmp13951 := GCons(__e, V5942, tmp13950)

tmp13952 := GCons(__e, symResult, tmp13951)

tmp13953 := GCons(__e, symlet, tmp13952)

tmp13954 := GCons(__e, tmp13953, Nil)

tmp13955 := GCons(__e, tmp13928, tmp13954)

tmp13956 := GCons(__e, symStart, tmp13955)

return GCons(__e, symlet, tmp13956)


panic("unreachable")
})

tmp13957 := Call2(__e, ns2_1set, symshen_4profile_1func, tmp13926)


_ = tmp13957

tmp13958 := MakeNative1(func(__e *ControlFlow, V5943 Obj) Obj {
tmp13959 := Call1(__e, SymFn(symshen_4get_1profile), V5943)


let__13533 := tmp13959
_ = let__13533

tmp13960 := Call2(__e, SymFn(symshen_4put_1profile), V5943, MakeInteger(0))


let__13534 := tmp13960
_ = let__13534

return TailCall2(__e, SymFn(sym_8p), V5943, let__13533)


panic("unreachable")
})

tmp13961 := Call2(__e, ns2_1set, symprofile_1results, tmp13958)


_ = tmp13961

tmp13962 := MakeNative1(func(__e *ControlFlow, V5946 Obj) Obj {
tmp13963 := MakeNative0(func(__e *ControlFlow) Obj {
tmp13964 := GValue(__e, sym_dproperty_1vector_d)

return TailCall3(__e, SymFn(symget), V5946, symprofile, tmp13964)


panic("unreachable")
})

tmp13965 := MakeNative1(func(__e *ControlFlow, Z5947 Obj) Obj {
return MakeInteger(0)
panic("unreachable")
})

return TailCall2(__e, try_1catch, tmp13963, tmp13965)


panic("unreachable")
})

tmp13966 := Call2(__e, ns2_1set, symshen_4get_1profile, tmp13962)


_ = tmp13966

tmp13967 := MakeNative2(func(__e *ControlFlow, V5948, V5949 Obj) Obj {
tmp13968 := GValue(__e, sym_dproperty_1vector_d)

return TailCall4(__e, SymFn(symput), V5948, symprofile, V5949, tmp13968)


panic("unreachable")
})

return TailCall2(__e, ns2_1set, symshen_4put_1profile, tmp13967)




panic("unreachable")
})

