package main

import . "github.com/pyrex41/shen-go/kl"

var TypesMain = MakeNative(func(__e *ControlFlow) {
tmp18628 := MakeNative(func(__e *ControlFlow) {
V5982 := __e.Get(1)
_ = V5982
V5983 := __e.Get(2)
_ = V5983
tmp18629 := Call(__e, PrimFunc(symshen_4rectify_1type), V5983)


let__18599 := tmp18629
_ = let__18599

tmp18630 := MakeNative(func(__e *ControlFlow) {
Z5986 := __e.Get(1)
_ = Z5986
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5987 := __e.Get(1)
_ = Z5987
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5988 := __e.Get(1)
_ = Z5988
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5989 := __e.Get(1)
_ = Z5989
tmp18631 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp18631

tmp18632 := Call(__e, PrimFunc(symshen_4deref), V5982, Z5986)


tmp18633 := Call(__e, PrimFunc(symreceive), tmp18632)


tmp18634 := Call(__e, PrimFunc(symshen_4deref), let__18599, Z5986)


tmp18635 := Call(__e, PrimFunc(symreceive), tmp18634)


__e.TailApply(PrimFunc(symshen_4variancy), tmp18633, tmp18635, Z5986, Z5987, Z5988, Z5989)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18636 := Call(__e, PrimFunc(symshen_4prolog_1vector))


tmp18637 := Call(__e, tmp18630, tmp18636)


tmp18638 := Call(__e, PrimFunc(symvector), MakeInteger(0))


tmp18639 := Call(__e, PrimFunc(sym_8v), MakeInteger(0), tmp18638)


tmp18640 := Call(__e, PrimFunc(sym_8v), True, tmp18639)


tmp18641 := Call(__e, tmp18637, tmp18640)


tmp18642 := Call(__e, tmp18641, MakeInteger(0))


tmp18643 := MakeNative(func(__e *ControlFlow) {
__e.Return(True)
return
}, 0)

tmp18644 := Call(__e, tmp18642, tmp18643)


let__18600 := tmp18644
_ = let__18600

tmp18645 := Call(__e, PrimFunc(symshen_4prolog_1abstraction), V5983)


tmp18646 := Call(__e, PrimFunc(symeval_1kl), tmp18645)


let__18601 := tmp18646
_ = let__18601

tmp18647 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp18648 := Call(__e, PrimFunc(symshen_4assoc_1_6), V5982, let__18601, tmp18647)


tmp18649 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp18648)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp18648
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

let__18602 := tmp18649
_ = let__18602

__e.Return(V5982)
return


}, 2)

tmp18650 := Call(__e, ns2_1set, symdeclare, tmp18628)


_ = tmp18650

tmp18651 := MakeNative(func(__e *ControlFlow) {
V5992 := __e.Get(1)
_ = V5992
V5993 := __e.Get(2)
_ = V5993
V5994 := __e.Get(3)
_ = V5994
V5995 := __e.Get(4)
_ = V5995
V5996 := __e.Get(5)
_ = V5996
V5997 := __e.Get(6)
_ = V5997
tmp18662 := Call(__e, PrimFunc(symshen_4unlocked_2), V5995)


if True == tmp18662 {
tmp18652 := Call(__e, PrimFunc(symshen_4newpv), V5994)


let__18603 := tmp18652
_ = let__18603

tmp18653 := Call(__e, PrimFunc(symshen_4newpv), V5994)


let__18604 := tmp18653
_ = let__18604

tmp18654 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp18654

tmp18655 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5992, Nil)
}
__typedArg0 := V5992
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18656 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfn, tmp18655)
}
__typedArg0 := symfn
__typedArg1 := tmp18655
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18657 := MakeNative(func(__e *ControlFlow) {
tmp18658 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4variants_2), V5992, let__18604, V5993, V5994, V5995, V5996, V5997)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4variancy_1signature), V5992, let__18603, let__18604, V5994, V5995, V5996, tmp18658)
return


}, 0)

tmp18659 := Call(__e, PrimFunc(symshen_4system_1S_1h), tmp18656, let__18603, Nil, V5994, V5995, V5996, tmp18657)


tmp18660 := Call(__e, PrimFunc(symshen_4gc), V5994, tmp18659)


__e.TailApply(PrimFunc(symshen_4gc), V5994, tmp18660)
return


} else {
__e.Return(False)
return
}


}, 6)

tmp18663 := Call(__e, ns2_1set, symshen_4variancy, tmp18651)


_ = tmp18663

tmp18664 := MakeNative(func(__e *ControlFlow) {
V6000 := __e.Get(1)
_ = V6000
V6001 := __e.Get(2)
_ = V6001
V6002 := __e.Get(3)
_ = V6002
V6003 := __e.Get(4)
_ = V6003
V6004 := __e.Get(5)
_ = V6004
V6005 := __e.Get(6)
_ = V6005
V6006 := __e.Get(7)
_ = V6006
let__18605 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V6005)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V6005
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
_ = let__18605

tmp18713 := Call(__e, PrimFunc(symshen_4unlocked_2), V6004)


var ifres18666 Obj

if True == tmp18713 {
tmp18667 := Call(__e, PrimFunc(symshen_4lazyderef), V6002, V6003)


let__18607 := tmp18667
_ = let__18607

tmp18668 := MakeNative(func(__e *ControlFlow) {
Z6011 := __e.Get(1)
_ = Z6011
tmp18669 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp18669

tmp18670 := MakeNative(func(__e *ControlFlow) {
tmp18671 := Call(__e, PrimFunc(symshen_4deref), V6000, V6003)


tmp18672 := Call(__e, PrimFunc(symarity), tmp18671)


tmp18673 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp18672, MakeInteger(0))
}
__typedArg0 := tmp18672
__typedArg1 := MakeInteger(0)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

tmp18674 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4cut), V6003, V6004, let__18605, V6006)
return
}, 0)

__e.TailApply(PrimFunc(symwhen), tmp18673, V6003, V6004, let__18605, tmp18674)
return


}, 0)

__e.TailApply(PrimFunc(symis_b), V6001, Z6011, V6003, V6004, let__18605, tmp18670)
return


}, 1)

let__18608 := tmp18668
_ = let__18608

tmp18712 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__18607)
}
__typedArg0 := let__18607
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18675 Obj

if True == tmp18712 {
tmp18676 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__18607)
}
__typedArg0 := let__18607
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18677 := Call(__e, PrimFunc(symshen_4lazyderef), tmp18676, V6003)


let__18609 := tmp18677
_ = let__18609

tmp18678 := MakeNative(func(__e *ControlFlow) {
tmp18679 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__18607)
}
__typedArg0 := let__18607
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18680 := Call(__e, PrimFunc(symshen_4lazyderef), tmp18679, V6003)


let__18611 := tmp18680
_ = let__18611

tmp18681 := MakeNative(func(__e *ControlFlow) {
Z6016 := __e.Get(1)
_ = Z6016
__e.TailApply(let__18608, Z6016)
return
}, 1)

let__18612 := tmp18681
_ = let__18612

tmp18697 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__18611)
}
__typedArg0 := let__18611
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp18697 {
tmp18682 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__18611)
}
__typedArg0 := let__18611
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__18613 := tmp18682
_ = let__18613

tmp18683 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__18611)
}
__typedArg0 := let__18611
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18684 := Call(__e, PrimFunc(symshen_4lazyderef), tmp18683, V6003)


let__18614 := tmp18684
_ = let__18614

tmp18685 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__18612, let__18613)
return
}, 0)

let__18615 := tmp18685
_ = let__18615

tmp18689 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__18614, Nil)
}
__typedArg0 := let__18614
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp18689 {
__e.TailApply(PrimFunc(symthaw), let__18615)
return
} else {
tmp18687 := Call(__e, PrimFunc(symshen_4pvar_2), let__18614)


if True == tmp18687 {
__e.TailApply(PrimFunc(symshen_4bind_b), let__18614, Nil, V6003, let__18615)
return
} else {
__e.Return(False)
return
}


}


} else {
tmp18695 := Call(__e, PrimFunc(symshen_4pvar_2), let__18611)


if True == tmp18695 {
tmp18690 := Call(__e, PrimFunc(symshen_4newpv), V6003)


let__18616 := tmp18690
_ = let__18616

tmp18691 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18616, Nil)
}
__typedArg0 := let__18616
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18692 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__18612, let__18616)
return
}, 0)

tmp18693 := Call(__e, PrimFunc(symshen_4bind_b), let__18611, tmp18691, V6003, tmp18692)


__e.TailApply(PrimFunc(symshen_4gc), V6003, tmp18693)
return


} else {
__e.Return(False)
return
}


}


}, 0)

let__18610 := tmp18678
_ = let__18610

tmp18703 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__18609, sym_1_1_6)
}
__typedArg0 := let__18609
__typedArg1 := sym_1_1_6
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18698 Obj

if True == tmp18703 {
tmp18699 := Call(__e, PrimFunc(symthaw), let__18610)


ifres18698 = tmp18699


} else {
tmp18702 := Call(__e, PrimFunc(symshen_4pvar_2), let__18609)


var ifres18700 Obj

if True == tmp18702 {
tmp18701 := Call(__e, PrimFunc(symshen_4bind_b), let__18609, sym_1_1_6, V6003, let__18610)


ifres18700 = tmp18701


} else {
ifres18700 = False


}

ifres18698 = ifres18700


}

ifres18675 = ifres18698


} else {
tmp18711 := Call(__e, PrimFunc(symshen_4pvar_2), let__18607)


var ifres18704 Obj

if True == tmp18711 {
tmp18705 := Call(__e, PrimFunc(symshen_4newpv), V6003)


let__18617 := tmp18705
_ = let__18617

tmp18706 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18617, Nil)
}
__typedArg0 := let__18617
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18707 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp18706)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp18706
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18708 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__18608, let__18617)
return
}, 0)

tmp18709 := Call(__e, PrimFunc(symshen_4bind_b), let__18607, tmp18707, V6003, tmp18708)


tmp18710 := Call(__e, PrimFunc(symshen_4gc), V6003, tmp18709)


ifres18704 = tmp18710


} else {
ifres18704 = False


}

ifres18675 = ifres18704


}

ifres18666 = ifres18675


} else {
ifres18666 = False


}

let__18606 := ifres18666
_ = let__18606

tmp18721 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__18606, False)
}
__typedArg0 := let__18606
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp18721 {
tmp18717 := Call(__e, PrimFunc(symshen_4unlocked_2), V6004)


var ifres18714 Obj

if True == tmp18717 {
tmp18715 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp18715

tmp18716 := Call(__e, PrimFunc(symis_b), V6001, V6002, V6003, V6004, let__18605, V6006)


ifres18714 = tmp18716


} else {
ifres18714 = False


}

let__18618 := ifres18714
_ = let__18618

tmp18719 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__18618, False)
}
__typedArg0 := let__18618
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp18719 {
__e.TailApply(PrimFunc(symshen_4unlock), V6004, let__18605)
return
} else {
__e.Return(let__18618)
return
}


} else {
__e.Return(let__18606)
return
}


}, 7)

tmp18722 := Call(__e, ns2_1set, symshen_4variancy_1signature, tmp18664)


_ = tmp18722

tmp18723 := MakeNative(func(__e *ControlFlow) {
V6023 := __e.Get(1)
_ = V6023
V6024 := __e.Get(2)
_ = V6024
V6025 := __e.Get(3)
_ = V6025
V6026 := __e.Get(4)
_ = V6026
V6027 := __e.Get(5)
_ = V6027
V6028 := __e.Get(6)
_ = V6028
V6029 := __e.Get(7)
_ = V6029
tmp18727 := Call(__e, PrimFunc(symshen_4unlocked_2), V6027)


var ifres18724 Obj

if True == tmp18727 {
tmp18725 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp18725

tmp18726 := Call(__e, PrimFunc(symis_b), V6024, V6025, V6026, V6027, V6028, V6029)


ifres18724 = tmp18726


} else {
ifres18724 = False


}

let__18619 := ifres18724
_ = let__18619

tmp18739 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__18619, False)
}
__typedArg0 := let__18619
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp18739 {
tmp18737 := Call(__e, PrimFunc(symshen_4unlocked_2), V6027)


if True == tmp18737 {
tmp18728 := Call(__e, PrimFunc(symshen_4newpv), V6026)


let__18620 := tmp18728
_ = let__18620

tmp18729 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp18729

tmp18730 := Call(__e, PrimFunc(symshen_4deref), V6023, V6026)


tmp18731 := Call(__e, PrimFunc(symshen_4app), tmp18730, MakeString(" may create errors\n"), symshen_4a)


tmp18732 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("warning: changing the type of "))
__typedS1, __typedOK1 := TypedString(tmp18731)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("warning: changing the type of ")
__typedArg1 := tmp18731
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp18733 := Call(__e, PrimFunc(symstoutput))


tmp18734 := Call(__e, PrimFunc(sympr), tmp18732, tmp18733)


tmp18735 := Call(__e, PrimFunc(symis), let__18620, tmp18734, V6026, V6027, V6028, V6029)


__e.TailApply(PrimFunc(symshen_4gc), V6026, tmp18735)
return


} else {
__e.Return(False)
return
}


} else {
__e.Return(let__18619)
return
}


}, 7)

tmp18740 := Call(__e, ns2_1set, symshen_4variants_2, tmp18723)


_ = tmp18740

tmp18741 := MakeNative(func(__e *ControlFlow) {
V6032 := __e.Get(1)
_ = V6032
tmp18742 := Call(__e, PrimFunc(symgensym), symB)


let__18621 := tmp18742
_ = let__18621

tmp18743 := Call(__e, PrimFunc(symgensym), symL)


let__18622 := tmp18743
_ = let__18622

tmp18744 := Call(__e, PrimFunc(symgensym), symKey)


let__18623 := tmp18744
_ = let__18623

tmp18745 := Call(__e, PrimFunc(symgensym), symC)


let__18624 := tmp18745
_ = let__18624

tmp18746 := Call(__e, PrimFunc(symgensym), symV)


let__18625 := tmp18746
_ = let__18625

tmp18747 := Call(__e, PrimFunc(symshen_4extract_1vars), V6032)


let__18626 := tmp18747
_ = let__18626

tmp18748 := Call(__e, PrimFunc(symshen_4rcons__form), V6032)


tmp18749 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18624, Nil)
}
__typedArg0 := let__18624
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18750 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18623, tmp18749)
}
__typedArg0 := let__18623
__typedArg1 := tmp18749
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18751 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18622, tmp18750)
}
__typedArg0 := let__18622
__typedArg1 := tmp18750
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18752 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18621, tmp18751)
}
__typedArg0 := let__18621
__typedArg1 := tmp18751
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18753 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18748, tmp18752)
}
__typedArg0 := tmp18748
__typedArg1 := tmp18752
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18754 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18625, tmp18753)
}
__typedArg0 := let__18625
__typedArg1 := tmp18753
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18755 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symis_b, tmp18754)
}
__typedArg0 := symis_b
__typedArg1 := tmp18754
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18756 := Call(__e, PrimFunc(symshen_4stpart), let__18626, tmp18755, let__18621)


tmp18757 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18756, Nil)
}
__typedArg0 := tmp18756
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18758 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18624, tmp18757)
}
__typedArg0 := let__18624
__typedArg1 := tmp18757
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18759 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp18758)
}
__typedArg0 := symlambda
__typedArg1 := tmp18758
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18760 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18759, Nil)
}
__typedArg0 := tmp18759
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18761 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18623, tmp18760)
}
__typedArg0 := let__18623
__typedArg1 := tmp18760
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18762 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp18761)
}
__typedArg0 := symlambda
__typedArg1 := tmp18761
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18763 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18762, Nil)
}
__typedArg0 := tmp18762
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18764 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18622, tmp18763)
}
__typedArg0 := let__18622
__typedArg1 := tmp18763
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18765 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp18764)
}
__typedArg0 := symlambda
__typedArg1 := tmp18764
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18766 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18765, Nil)
}
__typedArg0 := tmp18765
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18767 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18621, tmp18766)
}
__typedArg0 := let__18621
__typedArg1 := tmp18766
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18768 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp18767)
}
__typedArg0 := symlambda
__typedArg1 := tmp18767
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18769 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18768, Nil)
}
__typedArg0 := tmp18768
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18770 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18625, tmp18769)
}
__typedArg0 := let__18625
__typedArg1 := tmp18769
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp18770)
}
__typedArg0 := symlambda
__typedArg1 := tmp18770
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp18771 := Call(__e, ns2_1set, symshen_4prolog_1abstraction, tmp18741)


_ = tmp18771

tmp18772 := MakeNative(func(__e *ControlFlow) {
V6039 := __e.Get(1)
_ = V6039
tmp18773 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_ddemodulation_1function_d)
}
__typedArg0 := symshen_4_ddemodulation_1function_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

let__18627 := tmp18773
_ = let__18627

__e.TailApply(let__18627, V6039)
return


}, 1)

__e.TailApply(ns2_1set, symshen_4demod, tmp18772)
return




}, 0)

