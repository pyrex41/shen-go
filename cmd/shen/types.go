package main

import . "github.com/pyrex41/shen-go/kl"

var TypesMain = MakeNative0(func(__e *ControlFlow) Obj {
tmp18638 := MakeNative2(func(__e *ControlFlow, V6010, V6011 Obj) Obj {
tmp18639 := Call1(__e, SymFn(symshen_4rectify_1type), V6011)


let__18609 := tmp18639
_ = let__18609

tmp18640 := MakeNative1(func(__e *ControlFlow, Z6014 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Z6015 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Z6016 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Z6017 Obj) Obj {
tmp18641 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp18641

tmp18642 := Call2(__e, SymFn(symshen_4deref), V6010, Z6014)


tmp18643 := Call1(__e, SymFn(symreceive), tmp18642)


tmp18644 := Call2(__e, SymFn(symshen_4deref), let__18609, Z6014)


tmp18645 := Call1(__e, SymFn(symreceive), tmp18644)


return TailCall6(__e, SymFn(symshen_4variancy), tmp18643, tmp18645, Z6014, Z6015, Z6016, Z6017)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp18646 := Call0(__e, SymFn(symshen_4prolog_1vector))


tmp18647 := Call1(__e, tmp18640, tmp18646)


tmp18648 := Call1(__e, SymFn(symvector), MakeInteger(0))


tmp18649 := Call2(__e, SymFn(sym_8v), MakeInteger(0), tmp18648)


tmp18650 := Call2(__e, SymFn(sym_8v), True, tmp18649)


tmp18651 := Call1(__e, tmp18647, tmp18650)


tmp18652 := Call1(__e, tmp18651, MakeInteger(0))


tmp18653 := MakeNative0(func(__e *ControlFlow) Obj {
return True
panic("unreachable")
})

tmp18654 := Call1(__e, tmp18652, tmp18653)


let__18610 := tmp18654
_ = let__18610

tmp18655 := Call1(__e, SymFn(symshen_4prolog_1abstraction), V6011)


tmp18656 := Call1(__e, SymFn(symeval_1kl), tmp18655)


let__18611 := tmp18656
_ = let__18611

tmp18657 := GValue(__e, symshen_4_dsigf_d)

tmp18658 := Call3(__e, SymFn(symshen_4assoc_1_6), V6010, let__18611, tmp18657)


tmp18659 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp18658)

let__18612 := tmp18659
_ = let__18612

return V6010


panic("unreachable")
})

tmp18660 := Call2(__e, ns2_1set, symdeclare, tmp18638)


_ = tmp18660

tmp18661 := MakeNative6(func(__e *ControlFlow, V6020, V6021, V6022, V6023, V6024, V6025 Obj) Obj {
tmp18672 := Call1(__e, SymFn(symshen_4unlocked_2), V6023)


if True == tmp18672 {
tmp18662 := Call1(__e, SymFn(symshen_4newpv), V6022)


let__18613 := tmp18662
_ = let__18613

tmp18663 := Call1(__e, SymFn(symshen_4newpv), V6022)


let__18614 := tmp18663
_ = let__18614

tmp18664 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp18664

tmp18665 := GCons(__e, V6020, Nil)

tmp18666 := GCons(__e, symfn, tmp18665)

tmp18667 := MakeNative0(func(__e *ControlFlow) Obj {
tmp18668 := MakeNative0(func(__e *ControlFlow) Obj {
__e.TailApply(SymFn(symshen_4variants_2), V6020, let__18614, V6021, V6022, V6023, V6024, V6025)
return nil
panic("unreachable")
})

__e.TailApply(SymFn(symshen_4variancy_1signature), V6020, let__18613, let__18614, V6022, V6023, V6024, tmp18668)
return nil


panic("unreachable")
})

tmp18669 := Call(__e, SymFn(symshen_4system_1S_1h), tmp18666, let__18613, Nil, V6022, V6023, V6024, tmp18667)


tmp18670 := Call2(__e, SymFn(symshen_4gc), V6022, tmp18669)


return TailCall2(__e, SymFn(symshen_4gc), V6022, tmp18670)


} else {
return False
}


panic("unreachable")
})

tmp18673 := Call2(__e, ns2_1set, symshen_4variancy, tmp18661)


_ = tmp18673

tmp18674 := MakeNative(func(__e *ControlFlow) {
V6028 := __e.Get(1)
V6029 := __e.Get(2)
V6030 := __e.Get(3)
V6031 := __e.Get(4)
V6032 := __e.Get(5)
V6033 := __e.Get(6)
V6034 := __e.Get(7)
__e.ReturnDirect(func() Obj {
let__18615 := GAdd(__e, V6033, MakeInteger(1))
_ = let__18615

tmp18723 := Call1(__e, SymFn(symshen_4unlocked_2), V6032)


var ifres18676 Obj

if True == tmp18723 {
tmp18677 := Call2(__e, SymFn(symshen_4lazyderef), V6030, V6031)


let__18617 := tmp18677
_ = let__18617

tmp18678 := MakeNative1(func(__e *ControlFlow, Z6039 Obj) Obj {
tmp18679 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp18679

tmp18680 := MakeNative0(func(__e *ControlFlow) Obj {
tmp18681 := Call2(__e, SymFn(symshen_4deref), V6028, V6031)


tmp18682 := Call1(__e, SymFn(symarity), tmp18681)


tmp18683 := GEq(__e, tmp18682, MakeInteger(0))

tmp18684 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall4(__e, SymFn(symshen_4cut), V6031, V6032, let__18615, V6034)
panic("unreachable")
})

return TailCall5(__e, SymFn(symwhen), tmp18683, V6031, V6032, let__18615, tmp18684)


panic("unreachable")
})

return TailCall6(__e, SymFn(symis_b), V6029, Z6039, V6031, V6032, let__18615, tmp18680)


panic("unreachable")
})

let__18618 := tmp18678
_ = let__18618

tmp18722 := GConsP(__e, let__18617)

var ifres18685 Obj

if True == tmp18722 {
tmp18686 := GHd(__e, let__18617)

tmp18687 := Call2(__e, SymFn(symshen_4lazyderef), tmp18686, V6031)


let__18619 := tmp18687
_ = let__18619

tmp18688 := MakeNative0(func(__e *ControlFlow) Obj {
tmp18689 := GTl(__e, let__18617)

tmp18690 := Call2(__e, SymFn(symshen_4lazyderef), tmp18689, V6031)


let__18621 := tmp18690
_ = let__18621

tmp18691 := MakeNative1(func(__e *ControlFlow, Z6044 Obj) Obj {
return TailCall1(__e, let__18618, Z6044)
panic("unreachable")
})

let__18622 := tmp18691
_ = let__18622

tmp18707 := GConsP(__e, let__18621)

if True == tmp18707 {
tmp18692 := GHd(__e, let__18621)

let__18623 := tmp18692
_ = let__18623

tmp18693 := GTl(__e, let__18621)

tmp18694 := Call2(__e, SymFn(symshen_4lazyderef), tmp18693, V6031)


let__18624 := tmp18694
_ = let__18624

tmp18695 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, let__18622, let__18623)
panic("unreachable")
})

let__18625 := tmp18695
_ = let__18625

tmp18699 := GEq(__e, let__18624, Nil)

if True == tmp18699 {
return TailCall1(__e, SymFn(symthaw), let__18625)
} else {
tmp18697 := Call1(__e, SymFn(symshen_4pvar_2), let__18624)


if True == tmp18697 {
return TailCall4(__e, SymFn(symshen_4bind_b), let__18624, Nil, V6031, let__18625)
} else {
return False
}


}


} else {
tmp18705 := Call1(__e, SymFn(symshen_4pvar_2), let__18621)


if True == tmp18705 {
tmp18700 := Call1(__e, SymFn(symshen_4newpv), V6031)


let__18626 := tmp18700
_ = let__18626

tmp18701 := GCons(__e, let__18626, Nil)

tmp18702 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, let__18622, let__18626)
panic("unreachable")
})

tmp18703 := Call4(__e, SymFn(symshen_4bind_b), let__18621, tmp18701, V6031, tmp18702)


return TailCall2(__e, SymFn(symshen_4gc), V6031, tmp18703)


} else {
return False
}


}


panic("unreachable")
})

let__18620 := tmp18688
_ = let__18620

tmp18713 := GEq(__e, let__18619, sym_1_1_6)

var ifres18708 Obj

if True == tmp18713 {
tmp18709 := Call1(__e, SymFn(symthaw), let__18620)


ifres18708 = tmp18709


} else {
tmp18712 := Call1(__e, SymFn(symshen_4pvar_2), let__18619)


var ifres18710 Obj

if True == tmp18712 {
tmp18711 := Call4(__e, SymFn(symshen_4bind_b), let__18619, sym_1_1_6, V6031, let__18620)


ifres18710 = tmp18711


} else {
ifres18710 = False


}

ifres18708 = ifres18710


}

ifres18685 = ifres18708


} else {
tmp18721 := Call1(__e, SymFn(symshen_4pvar_2), let__18617)


var ifres18714 Obj

if True == tmp18721 {
tmp18715 := Call1(__e, SymFn(symshen_4newpv), V6031)


let__18627 := tmp18715
_ = let__18627

tmp18716 := GCons(__e, let__18627, Nil)

tmp18717 := GCons(__e, sym_1_1_6, tmp18716)

tmp18718 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, let__18618, let__18627)
panic("unreachable")
})

tmp18719 := Call4(__e, SymFn(symshen_4bind_b), let__18617, tmp18717, V6031, tmp18718)


tmp18720 := Call2(__e, SymFn(symshen_4gc), V6031, tmp18719)


ifres18714 = tmp18720


} else {
ifres18714 = False


}

ifres18685 = ifres18714


}

ifres18676 = ifres18685


} else {
ifres18676 = False


}

let__18616 := ifres18676
_ = let__18616

tmp18731 := GEq(__e, let__18616, False)

if True == tmp18731 {
tmp18727 := Call1(__e, SymFn(symshen_4unlocked_2), V6032)


var ifres18724 Obj

if True == tmp18727 {
tmp18725 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp18725

tmp18726 := Call6(__e, SymFn(symis_b), V6029, V6030, V6031, V6032, let__18615, V6034)


ifres18724 = tmp18726


} else {
ifres18724 = False


}

let__18628 := ifres18724
_ = let__18628

tmp18729 := GEq(__e, let__18628, False)

if True == tmp18729 {
return TailCall2(__e, SymFn(symshen_4unlock), V6032, let__18615)
} else {
return let__18628
}


} else {
return let__18616
}


panic("unreachable")
}())
}, 7)

tmp18732 := Call2(__e, ns2_1set, symshen_4variancy_1signature, tmp18674)


_ = tmp18732

tmp18733 := MakeNative(func(__e *ControlFlow) {
V6051 := __e.Get(1)
V6052 := __e.Get(2)
V6053 := __e.Get(3)
V6054 := __e.Get(4)
V6055 := __e.Get(5)
V6056 := __e.Get(6)
V6057 := __e.Get(7)
__e.ReturnDirect(func() Obj {
tmp18737 := Call1(__e, SymFn(symshen_4unlocked_2), V6055)


var ifres18734 Obj

if True == tmp18737 {
tmp18735 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp18735

tmp18736 := Call6(__e, SymFn(symis_b), V6052, V6053, V6054, V6055, V6056, V6057)


ifres18734 = tmp18736


} else {
ifres18734 = False


}

let__18629 := ifres18734
_ = let__18629

tmp18749 := GEq(__e, let__18629, False)

if True == tmp18749 {
tmp18747 := Call1(__e, SymFn(symshen_4unlocked_2), V6055)


if True == tmp18747 {
tmp18738 := Call1(__e, SymFn(symshen_4newpv), V6054)


let__18630 := tmp18738
_ = let__18630

tmp18739 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp18739

tmp18740 := Call2(__e, SymFn(symshen_4deref), V6051, V6054)


tmp18741 := Call3(__e, SymFn(symshen_4app), tmp18740, MakeString(" may create errors\n"), symshen_4a)


tmp18742 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("warning: changing the type of "), tmp18741)

tmp18743 := Call0(__e, SymFn(symstoutput))


tmp18744 := Call2(__e, SymFn(sympr), tmp18742, tmp18743)


tmp18745 := Call6(__e, SymFn(symis), let__18630, tmp18744, V6054, V6055, V6056, V6057)


return TailCall2(__e, SymFn(symshen_4gc), V6054, tmp18745)


} else {
return False
}


} else {
return let__18629
}


panic("unreachable")
}())
}, 7)

tmp18750 := Call2(__e, ns2_1set, symshen_4variants_2, tmp18733)


_ = tmp18750

tmp18751 := MakeNative1(func(__e *ControlFlow, V6060 Obj) Obj {
tmp18752 := Call1(__e, SymFn(symgensym), symB)


let__18631 := tmp18752
_ = let__18631

tmp18753 := Call1(__e, SymFn(symgensym), symL)


let__18632 := tmp18753
_ = let__18632

tmp18754 := Call1(__e, SymFn(symgensym), symKey)


let__18633 := tmp18754
_ = let__18633

tmp18755 := Call1(__e, SymFn(symgensym), symC)


let__18634 := tmp18755
_ = let__18634

tmp18756 := Call1(__e, SymFn(symgensym), symV)


let__18635 := tmp18756
_ = let__18635

tmp18757 := Call1(__e, SymFn(symshen_4extract_1vars), V6060)


let__18636 := tmp18757
_ = let__18636

tmp18758 := Call1(__e, SymFn(symshen_4rcons__form), V6060)


tmp18759 := GCons(__e, let__18634, Nil)

tmp18760 := GCons(__e, let__18633, tmp18759)

tmp18761 := GCons(__e, let__18632, tmp18760)

tmp18762 := GCons(__e, let__18631, tmp18761)

tmp18763 := GCons(__e, tmp18758, tmp18762)

tmp18764 := GCons(__e, let__18635, tmp18763)

tmp18765 := GCons(__e, symis_b, tmp18764)

tmp18766 := Call3(__e, SymFn(symshen_4stpart), let__18636, tmp18765, let__18631)


tmp18767 := GCons(__e, tmp18766, Nil)

tmp18768 := GCons(__e, let__18634, tmp18767)

tmp18769 := GCons(__e, symlambda, tmp18768)

tmp18770 := GCons(__e, tmp18769, Nil)

tmp18771 := GCons(__e, let__18633, tmp18770)

tmp18772 := GCons(__e, symlambda, tmp18771)

tmp18773 := GCons(__e, tmp18772, Nil)

tmp18774 := GCons(__e, let__18632, tmp18773)

tmp18775 := GCons(__e, symlambda, tmp18774)

tmp18776 := GCons(__e, tmp18775, Nil)

tmp18777 := GCons(__e, let__18631, tmp18776)

tmp18778 := GCons(__e, symlambda, tmp18777)

tmp18779 := GCons(__e, tmp18778, Nil)

tmp18780 := GCons(__e, let__18635, tmp18779)

return GCons(__e, symlambda, tmp18780)


panic("unreachable")
})

tmp18781 := Call2(__e, ns2_1set, symshen_4prolog_1abstraction, tmp18751)


_ = tmp18781

tmp18782 := MakeNative1(func(__e *ControlFlow, V6067 Obj) Obj {
tmp18783 := GValue(__e, symshen_4_ddemodulation_1function_d)

let__18637 := tmp18783
_ = let__18637

return TailCall1(__e, let__18637, V6067)


panic("unreachable")
})

return TailCall2(__e, ns2_1set, symshen_4demod, tmp18782)




panic("unreachable")
})

