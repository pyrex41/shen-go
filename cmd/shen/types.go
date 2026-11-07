package main

import . "github.com/pyrex41/shen-go/kl"

var TypesMain = MakeNative(func(__e *ControlFlow) {
tmp18658 := MakeNative(func(__e *ControlFlow) {
V6028 := __e.Get(1)
_ = V6028
V6029 := __e.Get(2)
_ = V6029
tmp18659 := MakeNative(func(__e *ControlFlow) {
tmp18660 := Call(__e, PrimFunc(symshen_4prolog_1abstraction), V6029)


__e.TailApply(PrimFunc(symeval_1kl), tmp18660)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4declare_1with_1signature), V6028, V6029, tmp18659)
return


}, 2)

tmp18661 := Call(__e, ns2_1set, symdeclare, tmp18658)


_ = tmp18661

tmp18662 := MakeNative(func(__e *ControlFlow) {
V6030 := __e.Get(1)
_ = V6030
V6031 := __e.Get(2)
_ = V6031
V6032 := __e.Get(3)
_ = V6032
tmp18663 := Call(__e, PrimFunc(symshen_4rectify_1type), V6031)


let__18629 := tmp18663
_ = let__18629

tmp18664 := MakeNative(func(__e *ControlFlow) {
Z6035 := __e.Get(1)
_ = Z6035
__e.Return(MakeNative(func(__e *ControlFlow) {
Z6036 := __e.Get(1)
_ = Z6036
__e.Return(MakeNative(func(__e *ControlFlow) {
Z6037 := __e.Get(1)
_ = Z6037
__e.Return(MakeNative(func(__e *ControlFlow) {
Z6038 := __e.Get(1)
_ = Z6038
tmp18665 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp18665

tmp18666 := Call(__e, PrimFunc(symshen_4deref), V6030, Z6035)


tmp18667 := Call(__e, PrimFunc(symreceive), tmp18666)


tmp18668 := Call(__e, PrimFunc(symshen_4deref), let__18629, Z6035)


tmp18669 := Call(__e, PrimFunc(symreceive), tmp18668)


__e.TailApply(PrimFunc(symshen_4variancy), tmp18667, tmp18669, Z6035, Z6036, Z6037, Z6038)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18670 := Call(__e, PrimFunc(symshen_4prolog_1vector))


tmp18671 := Call(__e, tmp18664, tmp18670)


tmp18672 := Call(__e, PrimFunc(symvector), MakeInteger(0))


tmp18673 := Call(__e, PrimFunc(sym_8v), MakeInteger(0), tmp18672)


tmp18674 := Call(__e, PrimFunc(sym_8v), True, tmp18673)


tmp18675 := Call(__e, tmp18671, tmp18674)


tmp18676 := Call(__e, tmp18675, MakeInteger(0))


tmp18677 := MakeNative(func(__e *ControlFlow) {
__e.Return(True)
return
}, 0)

tmp18678 := Call(__e, tmp18676, tmp18677)


let__18630 := tmp18678
_ = let__18630

tmp18679 := Call(__e, PrimFunc(symthaw), V6032)


let__18631 := tmp18679
_ = let__18631

tmp18680 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp18681 := Call(__e, PrimFunc(symshen_4assoc_1_6), V6030, let__18631, tmp18680)


tmp18682 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp18681)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp18681
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

let__18632 := tmp18682
_ = let__18632

__e.Return(V6030)
return


}, 3)

tmp18683 := Call(__e, ns2_1set, symshen_4declare_1with_1signature, tmp18662)


_ = tmp18683

tmp18684 := MakeNative(func(__e *ControlFlow) {
V6041 := __e.Get(1)
_ = V6041
V6042 := __e.Get(2)
_ = V6042
V6043 := __e.Get(3)
_ = V6043
V6044 := __e.Get(4)
_ = V6044
V6045 := __e.Get(5)
_ = V6045
V6046 := __e.Get(6)
_ = V6046
tmp18695 := Call(__e, PrimFunc(symshen_4unlocked_2), V6044)


if True == tmp18695 {
tmp18685 := Call(__e, PrimFunc(symshen_4newpv), V6043)


let__18633 := tmp18685
_ = let__18633

tmp18686 := Call(__e, PrimFunc(symshen_4newpv), V6043)


let__18634 := tmp18686
_ = let__18634

tmp18687 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp18687

tmp18688 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V6041, Nil)
}
__typedArg0 := V6041
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18689 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfn, tmp18688)
}
__typedArg0 := symfn
__typedArg1 := tmp18688
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18690 := MakeNative(func(__e *ControlFlow) {
tmp18691 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4variants_2), V6041, let__18634, V6042, V6043, V6044, V6045, V6046)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4variancy_1signature), V6041, let__18633, let__18634, V6043, V6044, V6045, tmp18691)
return


}, 0)

tmp18692 := Call(__e, PrimFunc(symshen_4system_1S_1h), tmp18689, let__18633, Nil, V6043, V6044, V6045, tmp18690)


tmp18693 := Call(__e, PrimFunc(symshen_4gc), V6043, tmp18692)


__e.TailApply(PrimFunc(symshen_4gc), V6043, tmp18693)
return


} else {
__e.Return(False)
return
}


}, 6)

tmp18696 := Call(__e, ns2_1set, symshen_4variancy, tmp18684)


_ = tmp18696

tmp18697 := MakeNative(func(__e *ControlFlow) {
V6049 := __e.Get(1)
_ = V6049
V6050 := __e.Get(2)
_ = V6050
V6051 := __e.Get(3)
_ = V6051
V6052 := __e.Get(4)
_ = V6052
V6053 := __e.Get(5)
_ = V6053
V6054 := __e.Get(6)
_ = V6054
V6055 := __e.Get(7)
_ = V6055
let__18635 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V6054)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V6054
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
_ = let__18635

tmp18746 := Call(__e, PrimFunc(symshen_4unlocked_2), V6053)


var ifres18699 Obj

if True == tmp18746 {
tmp18700 := Call(__e, PrimFunc(symshen_4lazyderef), V6051, V6052)


let__18637 := tmp18700
_ = let__18637

tmp18701 := MakeNative(func(__e *ControlFlow) {
Z6060 := __e.Get(1)
_ = Z6060
tmp18702 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp18702

tmp18703 := MakeNative(func(__e *ControlFlow) {
tmp18704 := Call(__e, PrimFunc(symshen_4deref), V6049, V6052)


tmp18705 := Call(__e, PrimFunc(symarity), tmp18704)


tmp18706 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp18705, MakeInteger(0))
}
__typedArg0 := tmp18705
__typedArg1 := MakeInteger(0)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

tmp18707 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4cut), V6052, V6053, let__18635, V6055)
return
}, 0)

__e.TailApply(PrimFunc(symwhen), tmp18706, V6052, V6053, let__18635, tmp18707)
return


}, 0)

__e.TailApply(PrimFunc(symis_b), V6050, Z6060, V6052, V6053, let__18635, tmp18703)
return


}, 1)

let__18638 := tmp18701
_ = let__18638

tmp18745 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__18637)
}
__typedArg0 := let__18637
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18708 Obj

if True == tmp18745 {
tmp18709 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__18637)
}
__typedArg0 := let__18637
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18710 := Call(__e, PrimFunc(symshen_4lazyderef), tmp18709, V6052)


let__18639 := tmp18710
_ = let__18639

tmp18711 := MakeNative(func(__e *ControlFlow) {
tmp18712 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__18637)
}
__typedArg0 := let__18637
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18713 := Call(__e, PrimFunc(symshen_4lazyderef), tmp18712, V6052)


let__18641 := tmp18713
_ = let__18641

tmp18714 := MakeNative(func(__e *ControlFlow) {
Z6065 := __e.Get(1)
_ = Z6065
__e.TailApply(let__18638, Z6065)
return
}, 1)

let__18642 := tmp18714
_ = let__18642

tmp18730 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__18641)
}
__typedArg0 := let__18641
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp18730 {
tmp18715 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__18641)
}
__typedArg0 := let__18641
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__18643 := tmp18715
_ = let__18643

tmp18716 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__18641)
}
__typedArg0 := let__18641
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18717 := Call(__e, PrimFunc(symshen_4lazyderef), tmp18716, V6052)


let__18644 := tmp18717
_ = let__18644

tmp18718 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__18642, let__18643)
return
}, 0)

let__18645 := tmp18718
_ = let__18645

tmp18722 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__18644, Nil)
}
__typedArg0 := let__18644
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp18722 {
__e.TailApply(PrimFunc(symthaw), let__18645)
return
} else {
tmp18720 := Call(__e, PrimFunc(symshen_4pvar_2), let__18644)


if True == tmp18720 {
__e.TailApply(PrimFunc(symshen_4bind_b), let__18644, Nil, V6052, let__18645)
return
} else {
__e.Return(False)
return
}


}


} else {
tmp18728 := Call(__e, PrimFunc(symshen_4pvar_2), let__18641)


if True == tmp18728 {
tmp18723 := Call(__e, PrimFunc(symshen_4newpv), V6052)


let__18646 := tmp18723
_ = let__18646

tmp18724 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18646, Nil)
}
__typedArg0 := let__18646
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18725 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__18642, let__18646)
return
}, 0)

tmp18726 := Call(__e, PrimFunc(symshen_4bind_b), let__18641, tmp18724, V6052, tmp18725)


__e.TailApply(PrimFunc(symshen_4gc), V6052, tmp18726)
return


} else {
__e.Return(False)
return
}


}


}, 0)

let__18640 := tmp18711
_ = let__18640

tmp18736 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__18639, sym_1_1_6)
}
__typedArg0 := let__18639
__typedArg1 := sym_1_1_6
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18731 Obj

if True == tmp18736 {
tmp18732 := Call(__e, PrimFunc(symthaw), let__18640)


ifres18731 = tmp18732


} else {
tmp18735 := Call(__e, PrimFunc(symshen_4pvar_2), let__18639)


var ifres18733 Obj

if True == tmp18735 {
tmp18734 := Call(__e, PrimFunc(symshen_4bind_b), let__18639, sym_1_1_6, V6052, let__18640)


ifres18733 = tmp18734


} else {
ifres18733 = False


}

ifres18731 = ifres18733


}

ifres18708 = ifres18731


} else {
tmp18744 := Call(__e, PrimFunc(symshen_4pvar_2), let__18637)


var ifres18737 Obj

if True == tmp18744 {
tmp18738 := Call(__e, PrimFunc(symshen_4newpv), V6052)


let__18647 := tmp18738
_ = let__18647

tmp18739 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18647, Nil)
}
__typedArg0 := let__18647
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18740 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp18739)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp18739
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18741 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__18638, let__18647)
return
}, 0)

tmp18742 := Call(__e, PrimFunc(symshen_4bind_b), let__18637, tmp18740, V6052, tmp18741)


tmp18743 := Call(__e, PrimFunc(symshen_4gc), V6052, tmp18742)


ifres18737 = tmp18743


} else {
ifres18737 = False


}

ifres18708 = ifres18737


}

ifres18699 = ifres18708


} else {
ifres18699 = False


}

let__18636 := ifres18699
_ = let__18636

tmp18754 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__18636, False)
}
__typedArg0 := let__18636
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp18754 {
tmp18750 := Call(__e, PrimFunc(symshen_4unlocked_2), V6053)


var ifres18747 Obj

if True == tmp18750 {
tmp18748 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp18748

tmp18749 := Call(__e, PrimFunc(symis_b), V6050, V6051, V6052, V6053, let__18635, V6055)


ifres18747 = tmp18749


} else {
ifres18747 = False


}

let__18648 := ifres18747
_ = let__18648

tmp18752 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__18648, False)
}
__typedArg0 := let__18648
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp18752 {
__e.TailApply(PrimFunc(symshen_4unlock), V6053, let__18635)
return
} else {
__e.Return(let__18648)
return
}


} else {
__e.Return(let__18636)
return
}


}, 7)

tmp18755 := Call(__e, ns2_1set, symshen_4variancy_1signature, tmp18697)


_ = tmp18755

tmp18756 := MakeNative(func(__e *ControlFlow) {
V6072 := __e.Get(1)
_ = V6072
V6073 := __e.Get(2)
_ = V6073
V6074 := __e.Get(3)
_ = V6074
V6075 := __e.Get(4)
_ = V6075
V6076 := __e.Get(5)
_ = V6076
V6077 := __e.Get(6)
_ = V6077
V6078 := __e.Get(7)
_ = V6078
tmp18760 := Call(__e, PrimFunc(symshen_4unlocked_2), V6076)


var ifres18757 Obj

if True == tmp18760 {
tmp18758 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp18758

tmp18759 := Call(__e, PrimFunc(symis_b), V6073, V6074, V6075, V6076, V6077, V6078)


ifres18757 = tmp18759


} else {
ifres18757 = False


}

let__18649 := ifres18757
_ = let__18649

tmp18772 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__18649, False)
}
__typedArg0 := let__18649
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp18772 {
tmp18770 := Call(__e, PrimFunc(symshen_4unlocked_2), V6076)


if True == tmp18770 {
tmp18761 := Call(__e, PrimFunc(symshen_4newpv), V6075)


let__18650 := tmp18761
_ = let__18650

tmp18762 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp18762

tmp18763 := Call(__e, PrimFunc(symshen_4deref), V6072, V6075)


tmp18764 := Call(__e, PrimFunc(symshen_4app), tmp18763, MakeString(" may create errors\n"), symshen_4a)


tmp18765 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("warning: changing the type of "))
__typedS1, __typedOK1 := TypedString(tmp18764)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("warning: changing the type of ")
__typedArg1 := tmp18764
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp18766 := Call(__e, PrimFunc(symstoutput))


tmp18767 := Call(__e, PrimFunc(sympr), tmp18765, tmp18766)


tmp18768 := Call(__e, PrimFunc(symis), let__18650, tmp18767, V6075, V6076, V6077, V6078)


__e.TailApply(PrimFunc(symshen_4gc), V6075, tmp18768)
return


} else {
__e.Return(False)
return
}


} else {
__e.Return(let__18649)
return
}


}, 7)

tmp18773 := Call(__e, ns2_1set, symshen_4variants_2, tmp18756)


_ = tmp18773

tmp18774 := MakeNative(func(__e *ControlFlow) {
V6081 := __e.Get(1)
_ = V6081
tmp18775 := Call(__e, PrimFunc(symgensym), symB)


let__18651 := tmp18775
_ = let__18651

tmp18776 := Call(__e, PrimFunc(symgensym), symL)


let__18652 := tmp18776
_ = let__18652

tmp18777 := Call(__e, PrimFunc(symgensym), symKey)


let__18653 := tmp18777
_ = let__18653

tmp18778 := Call(__e, PrimFunc(symgensym), symC)


let__18654 := tmp18778
_ = let__18654

tmp18779 := Call(__e, PrimFunc(symgensym), symV)


let__18655 := tmp18779
_ = let__18655

tmp18780 := Call(__e, PrimFunc(symshen_4extract_1vars), V6081)


let__18656 := tmp18780
_ = let__18656

tmp18781 := Call(__e, PrimFunc(symshen_4rcons__form), V6081)


tmp18782 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18654, Nil)
}
__typedArg0 := let__18654
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18783 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18653, tmp18782)
}
__typedArg0 := let__18653
__typedArg1 := tmp18782
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18784 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18652, tmp18783)
}
__typedArg0 := let__18652
__typedArg1 := tmp18783
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18785 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18651, tmp18784)
}
__typedArg0 := let__18651
__typedArg1 := tmp18784
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18786 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18781, tmp18785)
}
__typedArg0 := tmp18781
__typedArg1 := tmp18785
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18787 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18655, tmp18786)
}
__typedArg0 := let__18655
__typedArg1 := tmp18786
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18788 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symis_b, tmp18787)
}
__typedArg0 := symis_b
__typedArg1 := tmp18787
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18789 := Call(__e, PrimFunc(symshen_4stpart), let__18656, tmp18788, let__18651)


tmp18790 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18789, Nil)
}
__typedArg0 := tmp18789
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18791 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18654, tmp18790)
}
__typedArg0 := let__18654
__typedArg1 := tmp18790
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18792 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp18791)
}
__typedArg0 := symlambda
__typedArg1 := tmp18791
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18793 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18792, Nil)
}
__typedArg0 := tmp18792
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18794 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18653, tmp18793)
}
__typedArg0 := let__18653
__typedArg1 := tmp18793
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18795 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp18794)
}
__typedArg0 := symlambda
__typedArg1 := tmp18794
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18796 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18795, Nil)
}
__typedArg0 := tmp18795
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18797 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18652, tmp18796)
}
__typedArg0 := let__18652
__typedArg1 := tmp18796
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18798 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp18797)
}
__typedArg0 := symlambda
__typedArg1 := tmp18797
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18799 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18798, Nil)
}
__typedArg0 := tmp18798
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18800 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18651, tmp18799)
}
__typedArg0 := let__18651
__typedArg1 := tmp18799
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18801 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp18800)
}
__typedArg0 := symlambda
__typedArg1 := tmp18800
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18802 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18801, Nil)
}
__typedArg0 := tmp18801
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18803 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18655, tmp18802)
}
__typedArg0 := let__18655
__typedArg1 := tmp18802
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp18803)
}
__typedArg0 := symlambda
__typedArg1 := tmp18803
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp18804 := Call(__e, ns2_1set, symshen_4prolog_1abstraction, tmp18774)


_ = tmp18804

tmp18805 := MakeNative(func(__e *ControlFlow) {
V6088 := __e.Get(1)
_ = V6088
tmp18806 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_ddemodulation_1function_d)
}
__typedArg0 := symshen_4_ddemodulation_1function_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

let__18657 := tmp18806
_ = let__18657

__e.TailApply(let__18657, V6088)
return


}, 1)

__e.TailApply(ns2_1set, symshen_4demod, tmp18805)
return




}, 0)

