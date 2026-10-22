package main

import . "github.com/pyrex41/shen-go/kl"

var TypesMain = MakeNative(func(__e *ControlFlow) {
tmp18638 := MakeNative(func(__e *ControlFlow) {
V6010 := __e.Get(1)
_ = V6010
V6011 := __e.Get(2)
_ = V6011
tmp18639 := Call(__e, PrimFunc(symshen_4rectify_1type), V6011)


let__18609 := tmp18639
_ = let__18609

tmp18640 := MakeNative(func(__e *ControlFlow) {
Z6014 := __e.Get(1)
_ = Z6014
__e.Return(MakeNative(func(__e *ControlFlow) {
Z6015 := __e.Get(1)
_ = Z6015
__e.Return(MakeNative(func(__e *ControlFlow) {
Z6016 := __e.Get(1)
_ = Z6016
__e.Return(MakeNative(func(__e *ControlFlow) {
Z6017 := __e.Get(1)
_ = Z6017
tmp18641 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp18641

tmp18642 := Call(__e, PrimFunc(symshen_4deref), V6010, Z6014)


tmp18643 := Call(__e, PrimFunc(symreceive), tmp18642)


tmp18644 := Call(__e, PrimFunc(symshen_4deref), let__18609, Z6014)


tmp18645 := Call(__e, PrimFunc(symreceive), tmp18644)


__e.TailApply(PrimFunc(symshen_4variancy), tmp18643, tmp18645, Z6014, Z6015, Z6016, Z6017)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18646 := Call(__e, PrimFunc(symshen_4prolog_1vector))


tmp18647 := Call(__e, tmp18640, tmp18646)


tmp18648 := Call(__e, PrimFunc(symvector), MakeInteger(0))


tmp18649 := Call(__e, PrimFunc(sym_8v), MakeInteger(0), tmp18648)


tmp18650 := Call(__e, PrimFunc(sym_8v), True, tmp18649)


tmp18651 := Call(__e, tmp18647, tmp18650)


tmp18652 := Call(__e, tmp18651, MakeInteger(0))


tmp18653 := MakeNative(func(__e *ControlFlow) {
__e.Return(True)
return
}, 0)

tmp18654 := Call(__e, tmp18652, tmp18653)


let__18610 := tmp18654
_ = let__18610

tmp18655 := Call(__e, PrimFunc(symshen_4prolog_1abstraction), V6011)


tmp18656 := Call(__e, PrimFunc(symeval_1kl), tmp18655)


let__18611 := tmp18656
_ = let__18611

tmp18657 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp18658 := Call(__e, PrimFunc(symshen_4assoc_1_6), V6010, let__18611, tmp18657)


tmp18659 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp18658)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp18658
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

let__18612 := tmp18659
_ = let__18612

__e.Return(V6010)
return


}, 2)

tmp18660 := Call(__e, ns2_1set, symdeclare, tmp18638)


_ = tmp18660

tmp18661 := MakeNative(func(__e *ControlFlow) {
V6020 := __e.Get(1)
_ = V6020
V6021 := __e.Get(2)
_ = V6021
V6022 := __e.Get(3)
_ = V6022
V6023 := __e.Get(4)
_ = V6023
V6024 := __e.Get(5)
_ = V6024
V6025 := __e.Get(6)
_ = V6025
tmp18672 := Call(__e, PrimFunc(symshen_4unlocked_2), V6023)


if True == tmp18672 {
tmp18662 := Call(__e, PrimFunc(symshen_4newpv), V6022)


let__18613 := tmp18662
_ = let__18613

tmp18663 := Call(__e, PrimFunc(symshen_4newpv), V6022)


let__18614 := tmp18663
_ = let__18614

tmp18664 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp18664

tmp18665 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V6020, Nil)
}
__typedArg0 := V6020
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18666 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfn, tmp18665)
}
__typedArg0 := symfn
__typedArg1 := tmp18665
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18667 := MakeNative(func(__e *ControlFlow) {
tmp18668 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4variants_2), V6020, let__18614, V6021, V6022, V6023, V6024, V6025)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4variancy_1signature), V6020, let__18613, let__18614, V6022, V6023, V6024, tmp18668)
return


}, 0)

tmp18669 := Call(__e, PrimFunc(symshen_4system_1S_1h), tmp18666, let__18613, Nil, V6022, V6023, V6024, tmp18667)


tmp18670 := Call(__e, PrimFunc(symshen_4gc), V6022, tmp18669)


__e.TailApply(PrimFunc(symshen_4gc), V6022, tmp18670)
return


} else {
__e.Return(False)
return
}


}, 6)

tmp18673 := Call(__e, ns2_1set, symshen_4variancy, tmp18661)


_ = tmp18673

tmp18674 := MakeNative(func(__e *ControlFlow) {
V6028 := __e.Get(1)
_ = V6028
V6029 := __e.Get(2)
_ = V6029
V6030 := __e.Get(3)
_ = V6030
V6031 := __e.Get(4)
_ = V6031
V6032 := __e.Get(5)
_ = V6032
V6033 := __e.Get(6)
_ = V6033
V6034 := __e.Get(7)
_ = V6034
let__18615 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V6033)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V6033
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
_ = let__18615

tmp18723 := Call(__e, PrimFunc(symshen_4unlocked_2), V6032)


var ifres18676 Obj

if True == tmp18723 {
tmp18677 := Call(__e, PrimFunc(symshen_4lazyderef), V6030, V6031)


let__18617 := tmp18677
_ = let__18617

tmp18678 := MakeNative(func(__e *ControlFlow) {
Z6039 := __e.Get(1)
_ = Z6039
tmp18679 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp18679

tmp18680 := MakeNative(func(__e *ControlFlow) {
tmp18681 := Call(__e, PrimFunc(symshen_4deref), V6028, V6031)


tmp18682 := Call(__e, PrimFunc(symarity), tmp18681)


tmp18683 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp18682, MakeInteger(0))
}
__typedArg0 := tmp18682
__typedArg1 := MakeInteger(0)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

tmp18684 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4cut), V6031, V6032, let__18615, V6034)
return
}, 0)

__e.TailApply(PrimFunc(symwhen), tmp18683, V6031, V6032, let__18615, tmp18684)
return


}, 0)

__e.TailApply(PrimFunc(symis_b), V6029, Z6039, V6031, V6032, let__18615, tmp18680)
return


}, 1)

let__18618 := tmp18678
_ = let__18618

tmp18722 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__18617)
}
__typedArg0 := let__18617
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18685 Obj

if True == tmp18722 {
tmp18686 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__18617)
}
__typedArg0 := let__18617
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18687 := Call(__e, PrimFunc(symshen_4lazyderef), tmp18686, V6031)


let__18619 := tmp18687
_ = let__18619

tmp18688 := MakeNative(func(__e *ControlFlow) {
tmp18689 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__18617)
}
__typedArg0 := let__18617
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18690 := Call(__e, PrimFunc(symshen_4lazyderef), tmp18689, V6031)


let__18621 := tmp18690
_ = let__18621

tmp18691 := MakeNative(func(__e *ControlFlow) {
Z6044 := __e.Get(1)
_ = Z6044
__e.TailApply(let__18618, Z6044)
return
}, 1)

let__18622 := tmp18691
_ = let__18622

tmp18707 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__18621)
}
__typedArg0 := let__18621
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp18707 {
tmp18692 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__18621)
}
__typedArg0 := let__18621
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__18623 := tmp18692
_ = let__18623

tmp18693 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__18621)
}
__typedArg0 := let__18621
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18694 := Call(__e, PrimFunc(symshen_4lazyderef), tmp18693, V6031)


let__18624 := tmp18694
_ = let__18624

tmp18695 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__18622, let__18623)
return
}, 0)

let__18625 := tmp18695
_ = let__18625

tmp18699 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__18624, Nil)
}
__typedArg0 := let__18624
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp18699 {
__e.TailApply(PrimFunc(symthaw), let__18625)
return
} else {
tmp18697 := Call(__e, PrimFunc(symshen_4pvar_2), let__18624)


if True == tmp18697 {
__e.TailApply(PrimFunc(symshen_4bind_b), let__18624, Nil, V6031, let__18625)
return
} else {
__e.Return(False)
return
}


}


} else {
tmp18705 := Call(__e, PrimFunc(symshen_4pvar_2), let__18621)


if True == tmp18705 {
tmp18700 := Call(__e, PrimFunc(symshen_4newpv), V6031)


let__18626 := tmp18700
_ = let__18626

tmp18701 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18626, Nil)
}
__typedArg0 := let__18626
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18702 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__18622, let__18626)
return
}, 0)

tmp18703 := Call(__e, PrimFunc(symshen_4bind_b), let__18621, tmp18701, V6031, tmp18702)


__e.TailApply(PrimFunc(symshen_4gc), V6031, tmp18703)
return


} else {
__e.Return(False)
return
}


}


}, 0)

let__18620 := tmp18688
_ = let__18620

tmp18713 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__18619, sym_1_1_6)
}
__typedArg0 := let__18619
__typedArg1 := sym_1_1_6
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18708 Obj

if True == tmp18713 {
tmp18709 := Call(__e, PrimFunc(symthaw), let__18620)


ifres18708 = tmp18709


} else {
tmp18712 := Call(__e, PrimFunc(symshen_4pvar_2), let__18619)


var ifres18710 Obj

if True == tmp18712 {
tmp18711 := Call(__e, PrimFunc(symshen_4bind_b), let__18619, sym_1_1_6, V6031, let__18620)


ifres18710 = tmp18711


} else {
ifres18710 = False


}

ifres18708 = ifres18710


}

ifres18685 = ifres18708


} else {
tmp18721 := Call(__e, PrimFunc(symshen_4pvar_2), let__18617)


var ifres18714 Obj

if True == tmp18721 {
tmp18715 := Call(__e, PrimFunc(symshen_4newpv), V6031)


let__18627 := tmp18715
_ = let__18627

tmp18716 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18627, Nil)
}
__typedArg0 := let__18627
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18717 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp18716)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp18716
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18718 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__18618, let__18627)
return
}, 0)

tmp18719 := Call(__e, PrimFunc(symshen_4bind_b), let__18617, tmp18717, V6031, tmp18718)


tmp18720 := Call(__e, PrimFunc(symshen_4gc), V6031, tmp18719)


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

tmp18731 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__18616, False)
}
__typedArg0 := let__18616
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp18731 {
tmp18727 := Call(__e, PrimFunc(symshen_4unlocked_2), V6032)


var ifres18724 Obj

if True == tmp18727 {
tmp18725 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp18725

tmp18726 := Call(__e, PrimFunc(symis_b), V6029, V6030, V6031, V6032, let__18615, V6034)


ifres18724 = tmp18726


} else {
ifres18724 = False


}

let__18628 := ifres18724
_ = let__18628

tmp18729 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__18628, False)
}
__typedArg0 := let__18628
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp18729 {
__e.TailApply(PrimFunc(symshen_4unlock), V6032, let__18615)
return
} else {
__e.Return(let__18628)
return
}


} else {
__e.Return(let__18616)
return
}


}, 7)

tmp18732 := Call(__e, ns2_1set, symshen_4variancy_1signature, tmp18674)


_ = tmp18732

tmp18733 := MakeNative(func(__e *ControlFlow) {
V6051 := __e.Get(1)
_ = V6051
V6052 := __e.Get(2)
_ = V6052
V6053 := __e.Get(3)
_ = V6053
V6054 := __e.Get(4)
_ = V6054
V6055 := __e.Get(5)
_ = V6055
V6056 := __e.Get(6)
_ = V6056
V6057 := __e.Get(7)
_ = V6057
tmp18737 := Call(__e, PrimFunc(symshen_4unlocked_2), V6055)


var ifres18734 Obj

if True == tmp18737 {
tmp18735 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp18735

tmp18736 := Call(__e, PrimFunc(symis_b), V6052, V6053, V6054, V6055, V6056, V6057)


ifres18734 = tmp18736


} else {
ifres18734 = False


}

let__18629 := ifres18734
_ = let__18629

tmp18749 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__18629, False)
}
__typedArg0 := let__18629
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp18749 {
tmp18747 := Call(__e, PrimFunc(symshen_4unlocked_2), V6055)


if True == tmp18747 {
tmp18738 := Call(__e, PrimFunc(symshen_4newpv), V6054)


let__18630 := tmp18738
_ = let__18630

tmp18739 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp18739

tmp18740 := Call(__e, PrimFunc(symshen_4deref), V6051, V6054)


tmp18741 := Call(__e, PrimFunc(symshen_4app), tmp18740, MakeString(" may create errors\n"), symshen_4a)


tmp18742 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("warning: changing the type of "))
__typedS1, __typedOK1 := TypedString(tmp18741)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("warning: changing the type of ")
__typedArg1 := tmp18741
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp18743 := Call(__e, PrimFunc(symstoutput))


tmp18744 := Call(__e, PrimFunc(sympr), tmp18742, tmp18743)


tmp18745 := Call(__e, PrimFunc(symis), let__18630, tmp18744, V6054, V6055, V6056, V6057)


__e.TailApply(PrimFunc(symshen_4gc), V6054, tmp18745)
return


} else {
__e.Return(False)
return
}


} else {
__e.Return(let__18629)
return
}


}, 7)

tmp18750 := Call(__e, ns2_1set, symshen_4variants_2, tmp18733)


_ = tmp18750

tmp18751 := MakeNative(func(__e *ControlFlow) {
V6060 := __e.Get(1)
_ = V6060
tmp18752 := Call(__e, PrimFunc(symgensym), symB)


let__18631 := tmp18752
_ = let__18631

tmp18753 := Call(__e, PrimFunc(symgensym), symL)


let__18632 := tmp18753
_ = let__18632

tmp18754 := Call(__e, PrimFunc(symgensym), symKey)


let__18633 := tmp18754
_ = let__18633

tmp18755 := Call(__e, PrimFunc(symgensym), symC)


let__18634 := tmp18755
_ = let__18634

tmp18756 := Call(__e, PrimFunc(symgensym), symV)


let__18635 := tmp18756
_ = let__18635

tmp18757 := Call(__e, PrimFunc(symshen_4extract_1vars), V6060)


let__18636 := tmp18757
_ = let__18636

tmp18758 := Call(__e, PrimFunc(symshen_4rcons__form), V6060)


tmp18759 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18634, Nil)
}
__typedArg0 := let__18634
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18760 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18633, tmp18759)
}
__typedArg0 := let__18633
__typedArg1 := tmp18759
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18761 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18632, tmp18760)
}
__typedArg0 := let__18632
__typedArg1 := tmp18760
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18762 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18631, tmp18761)
}
__typedArg0 := let__18631
__typedArg1 := tmp18761
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18763 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18758, tmp18762)
}
__typedArg0 := tmp18758
__typedArg1 := tmp18762
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18764 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18635, tmp18763)
}
__typedArg0 := let__18635
__typedArg1 := tmp18763
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18765 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symis_b, tmp18764)
}
__typedArg0 := symis_b
__typedArg1 := tmp18764
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18766 := Call(__e, PrimFunc(symshen_4stpart), let__18636, tmp18765, let__18631)


tmp18767 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18766, Nil)
}
__typedArg0 := tmp18766
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18768 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18634, tmp18767)
}
__typedArg0 := let__18634
__typedArg1 := tmp18767
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18769 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp18768)
}
__typedArg0 := symlambda
__typedArg1 := tmp18768
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18770 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18769, Nil)
}
__typedArg0 := tmp18769
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18771 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18633, tmp18770)
}
__typedArg0 := let__18633
__typedArg1 := tmp18770
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18772 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp18771)
}
__typedArg0 := symlambda
__typedArg1 := tmp18771
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18773 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18772, Nil)
}
__typedArg0 := tmp18772
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18774 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18632, tmp18773)
}
__typedArg0 := let__18632
__typedArg1 := tmp18773
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18775 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp18774)
}
__typedArg0 := symlambda
__typedArg1 := tmp18774
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18776 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18775, Nil)
}
__typedArg0 := tmp18775
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18777 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18631, tmp18776)
}
__typedArg0 := let__18631
__typedArg1 := tmp18776
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18778 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp18777)
}
__typedArg0 := symlambda
__typedArg1 := tmp18777
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18779 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18778, Nil)
}
__typedArg0 := tmp18778
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18780 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18635, tmp18779)
}
__typedArg0 := let__18635
__typedArg1 := tmp18779
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp18780)
}
__typedArg0 := symlambda
__typedArg1 := tmp18780
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp18781 := Call(__e, ns2_1set, symshen_4prolog_1abstraction, tmp18751)


_ = tmp18781

tmp18782 := MakeNative(func(__e *ControlFlow) {
V6067 := __e.Get(1)
_ = V6067
tmp18783 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_ddemodulation_1function_d)
}
__typedArg0 := symshen_4_ddemodulation_1function_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

let__18637 := tmp18783
_ = let__18637

__e.TailApply(let__18637, V6067)
return


}, 1)

__e.TailApply(ns2_1set, symshen_4demod, tmp18782)
return




}, 0)

