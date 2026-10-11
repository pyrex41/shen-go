package main

import . "github.com/pyrex41/shen-go/kl"

var TypesMain = MakeNative(func(__e *ControlFlow) {
tmp18621 := MakeNative(func(__e *ControlFlow) {
V6003 := __e.Get(1)
_ = V6003
V6004 := __e.Get(2)
_ = V6004
tmp18622 := Call(__e, PrimFunc(symshen_4rectify_1type), V6004)


let__18592 := tmp18622
_ = let__18592

tmp18623 := MakeNative(func(__e *ControlFlow) {
Z6007 := __e.Get(1)
_ = Z6007
__e.Return(MakeNative(func(__e *ControlFlow) {
Z6008 := __e.Get(1)
_ = Z6008
__e.Return(MakeNative(func(__e *ControlFlow) {
Z6009 := __e.Get(1)
_ = Z6009
__e.Return(MakeNative(func(__e *ControlFlow) {
Z6010 := __e.Get(1)
_ = Z6010
tmp18624 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp18624

tmp18625 := Call(__e, PrimFunc(symshen_4deref), V6003, Z6007)


tmp18626 := Call(__e, PrimFunc(symreceive), tmp18625)


tmp18627 := Call(__e, PrimFunc(symshen_4deref), let__18592, Z6007)


tmp18628 := Call(__e, PrimFunc(symreceive), tmp18627)


__e.TailApply(PrimFunc(symshen_4variancy), tmp18626, tmp18628, Z6007, Z6008, Z6009, Z6010)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18629 := Call(__e, PrimFunc(symshen_4prolog_1vector))


tmp18630 := Call(__e, tmp18623, tmp18629)


tmp18631 := Call(__e, PrimFunc(symvector), MakeInteger(0))


tmp18632 := Call(__e, PrimFunc(sym_8v), MakeInteger(0), tmp18631)


tmp18633 := Call(__e, PrimFunc(sym_8v), True, tmp18632)


tmp18634 := Call(__e, tmp18630, tmp18633)


tmp18635 := Call(__e, tmp18634, MakeInteger(0))


tmp18636 := MakeNative(func(__e *ControlFlow) {
__e.Return(True)
return
}, 0)

tmp18637 := Call(__e, tmp18635, tmp18636)


let__18593 := tmp18637
_ = let__18593

tmp18638 := Call(__e, PrimFunc(symshen_4prolog_1abstraction), V6004)


tmp18639 := Call(__e, PrimFunc(symeval_1kl), tmp18638)


let__18594 := tmp18639
_ = let__18594

tmp18640 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp18641 := Call(__e, PrimFunc(symshen_4assoc_1_6), V6003, let__18594, tmp18640)


tmp18642 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp18641)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp18641
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

let__18595 := tmp18642
_ = let__18595

__e.Return(V6003)
return


}, 2)

tmp18643 := Call(__e, ns2_1set, symdeclare, tmp18621)


_ = tmp18643

tmp18644 := MakeNative(func(__e *ControlFlow) {
V6013 := __e.Get(1)
_ = V6013
V6014 := __e.Get(2)
_ = V6014
V6015 := __e.Get(3)
_ = V6015
V6016 := __e.Get(4)
_ = V6016
V6017 := __e.Get(5)
_ = V6017
V6018 := __e.Get(6)
_ = V6018
tmp18655 := Call(__e, PrimFunc(symshen_4unlocked_2), V6016)


if True == tmp18655 {
tmp18645 := Call(__e, PrimFunc(symshen_4newpv), V6015)


let__18596 := tmp18645
_ = let__18596

tmp18646 := Call(__e, PrimFunc(symshen_4newpv), V6015)


let__18597 := tmp18646
_ = let__18597

tmp18647 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp18647

tmp18648 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V6013, Nil)
}
__typedArg0 := V6013
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18649 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfn, tmp18648)
}
__typedArg0 := symfn
__typedArg1 := tmp18648
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18650 := MakeNative(func(__e *ControlFlow) {
tmp18651 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4variants_2), V6013, let__18597, V6014, V6015, V6016, V6017, V6018)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4variancy_1signature), V6013, let__18596, let__18597, V6015, V6016, V6017, tmp18651)
return


}, 0)

tmp18652 := Call(__e, PrimFunc(symshen_4system_1S_1h), tmp18649, let__18596, Nil, V6015, V6016, V6017, tmp18650)


tmp18653 := Call(__e, PrimFunc(symshen_4gc), V6015, tmp18652)


__e.TailApply(PrimFunc(symshen_4gc), V6015, tmp18653)
return


} else {
__e.Return(False)
return
}


}, 6)

tmp18656 := Call(__e, ns2_1set, symshen_4variancy, tmp18644)


_ = tmp18656

tmp18657 := MakeNative(func(__e *ControlFlow) {
V6021 := __e.Get(1)
_ = V6021
V6022 := __e.Get(2)
_ = V6022
V6023 := __e.Get(3)
_ = V6023
V6024 := __e.Get(4)
_ = V6024
V6025 := __e.Get(5)
_ = V6025
V6026 := __e.Get(6)
_ = V6026
V6027 := __e.Get(7)
_ = V6027
let__18598 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V6026)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V6026
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
_ = let__18598

tmp18706 := Call(__e, PrimFunc(symshen_4unlocked_2), V6025)


var ifres18659 Obj

if True == tmp18706 {
tmp18660 := Call(__e, PrimFunc(symshen_4lazyderef), V6023, V6024)


let__18600 := tmp18660
_ = let__18600

tmp18661 := MakeNative(func(__e *ControlFlow) {
Z6032 := __e.Get(1)
_ = Z6032
tmp18662 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp18662

tmp18663 := MakeNative(func(__e *ControlFlow) {
tmp18664 := Call(__e, PrimFunc(symshen_4deref), V6021, V6024)


tmp18665 := Call(__e, PrimFunc(symarity), tmp18664)


tmp18666 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp18665, MakeInteger(0))
}
__typedArg0 := tmp18665
__typedArg1 := MakeInteger(0)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

tmp18667 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4cut), V6024, V6025, let__18598, V6027)
return
}, 0)

__e.TailApply(PrimFunc(symwhen), tmp18666, V6024, V6025, let__18598, tmp18667)
return


}, 0)

__e.TailApply(PrimFunc(symis_b), V6022, Z6032, V6024, V6025, let__18598, tmp18663)
return


}, 1)

let__18601 := tmp18661
_ = let__18601

tmp18705 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__18600)
}
__typedArg0 := let__18600
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18668 Obj

if True == tmp18705 {
tmp18669 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__18600)
}
__typedArg0 := let__18600
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18670 := Call(__e, PrimFunc(symshen_4lazyderef), tmp18669, V6024)


let__18602 := tmp18670
_ = let__18602

tmp18671 := MakeNative(func(__e *ControlFlow) {
tmp18672 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__18600)
}
__typedArg0 := let__18600
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18673 := Call(__e, PrimFunc(symshen_4lazyderef), tmp18672, V6024)


let__18604 := tmp18673
_ = let__18604

tmp18674 := MakeNative(func(__e *ControlFlow) {
Z6037 := __e.Get(1)
_ = Z6037
__e.TailApply(let__18601, Z6037)
return
}, 1)

let__18605 := tmp18674
_ = let__18605

tmp18690 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__18604)
}
__typedArg0 := let__18604
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp18690 {
tmp18675 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__18604)
}
__typedArg0 := let__18604
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__18606 := tmp18675
_ = let__18606

tmp18676 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__18604)
}
__typedArg0 := let__18604
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18677 := Call(__e, PrimFunc(symshen_4lazyderef), tmp18676, V6024)


let__18607 := tmp18677
_ = let__18607

tmp18678 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__18605, let__18606)
return
}, 0)

let__18608 := tmp18678
_ = let__18608

tmp18682 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__18607, Nil)
}
__typedArg0 := let__18607
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp18682 {
__e.TailApply(PrimFunc(symthaw), let__18608)
return
} else {
tmp18680 := Call(__e, PrimFunc(symshen_4pvar_2), let__18607)


if True == tmp18680 {
__e.TailApply(PrimFunc(symshen_4bind_b), let__18607, Nil, V6024, let__18608)
return
} else {
__e.Return(False)
return
}


}


} else {
tmp18688 := Call(__e, PrimFunc(symshen_4pvar_2), let__18604)


if True == tmp18688 {
tmp18683 := Call(__e, PrimFunc(symshen_4newpv), V6024)


let__18609 := tmp18683
_ = let__18609

tmp18684 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18609, Nil)
}
__typedArg0 := let__18609
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18685 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__18605, let__18609)
return
}, 0)

tmp18686 := Call(__e, PrimFunc(symshen_4bind_b), let__18604, tmp18684, V6024, tmp18685)


__e.TailApply(PrimFunc(symshen_4gc), V6024, tmp18686)
return


} else {
__e.Return(False)
return
}


}


}, 0)

let__18603 := tmp18671
_ = let__18603

tmp18696 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__18602, sym_1_1_6)
}
__typedArg0 := let__18602
__typedArg1 := sym_1_1_6
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18691 Obj

if True == tmp18696 {
tmp18692 := Call(__e, PrimFunc(symthaw), let__18603)


ifres18691 = tmp18692


} else {
tmp18695 := Call(__e, PrimFunc(symshen_4pvar_2), let__18602)


var ifres18693 Obj

if True == tmp18695 {
tmp18694 := Call(__e, PrimFunc(symshen_4bind_b), let__18602, sym_1_1_6, V6024, let__18603)


ifres18693 = tmp18694


} else {
ifres18693 = False


}

ifres18691 = ifres18693


}

ifres18668 = ifres18691


} else {
tmp18704 := Call(__e, PrimFunc(symshen_4pvar_2), let__18600)


var ifres18697 Obj

if True == tmp18704 {
tmp18698 := Call(__e, PrimFunc(symshen_4newpv), V6024)


let__18610 := tmp18698
_ = let__18610

tmp18699 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18610, Nil)
}
__typedArg0 := let__18610
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18700 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp18699)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp18699
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18701 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__18601, let__18610)
return
}, 0)

tmp18702 := Call(__e, PrimFunc(symshen_4bind_b), let__18600, tmp18700, V6024, tmp18701)


tmp18703 := Call(__e, PrimFunc(symshen_4gc), V6024, tmp18702)


ifres18697 = tmp18703


} else {
ifres18697 = False


}

ifres18668 = ifres18697


}

ifres18659 = ifres18668


} else {
ifres18659 = False


}

let__18599 := ifres18659
_ = let__18599

tmp18714 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__18599, False)
}
__typedArg0 := let__18599
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp18714 {
tmp18710 := Call(__e, PrimFunc(symshen_4unlocked_2), V6025)


var ifres18707 Obj

if True == tmp18710 {
tmp18708 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp18708

tmp18709 := Call(__e, PrimFunc(symis_b), V6022, V6023, V6024, V6025, let__18598, V6027)


ifres18707 = tmp18709


} else {
ifres18707 = False


}

let__18611 := ifres18707
_ = let__18611

tmp18712 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__18611, False)
}
__typedArg0 := let__18611
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp18712 {
__e.TailApply(PrimFunc(symshen_4unlock), V6025, let__18598)
return
} else {
__e.Return(let__18611)
return
}


} else {
__e.Return(let__18599)
return
}


}, 7)

tmp18715 := Call(__e, ns2_1set, symshen_4variancy_1signature, tmp18657)


_ = tmp18715

tmp18716 := MakeNative(func(__e *ControlFlow) {
V6044 := __e.Get(1)
_ = V6044
V6045 := __e.Get(2)
_ = V6045
V6046 := __e.Get(3)
_ = V6046
V6047 := __e.Get(4)
_ = V6047
V6048 := __e.Get(5)
_ = V6048
V6049 := __e.Get(6)
_ = V6049
V6050 := __e.Get(7)
_ = V6050
tmp18720 := Call(__e, PrimFunc(symshen_4unlocked_2), V6048)


var ifres18717 Obj

if True == tmp18720 {
tmp18718 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp18718

tmp18719 := Call(__e, PrimFunc(symis_b), V6045, V6046, V6047, V6048, V6049, V6050)


ifres18717 = tmp18719


} else {
ifres18717 = False


}

let__18612 := ifres18717
_ = let__18612

tmp18732 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__18612, False)
}
__typedArg0 := let__18612
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp18732 {
tmp18730 := Call(__e, PrimFunc(symshen_4unlocked_2), V6048)


if True == tmp18730 {
tmp18721 := Call(__e, PrimFunc(symshen_4newpv), V6047)


let__18613 := tmp18721
_ = let__18613

tmp18722 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp18722

tmp18723 := Call(__e, PrimFunc(symshen_4deref), V6044, V6047)


tmp18724 := Call(__e, PrimFunc(symshen_4app), tmp18723, MakeString(" may create errors\n"), symshen_4a)


tmp18725 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("warning: changing the type of "))
__typedS1, __typedOK1 := TypedString(tmp18724)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("warning: changing the type of ")
__typedArg1 := tmp18724
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp18726 := Call(__e, PrimFunc(symstoutput))


tmp18727 := Call(__e, PrimFunc(sympr), tmp18725, tmp18726)


tmp18728 := Call(__e, PrimFunc(symis), let__18613, tmp18727, V6047, V6048, V6049, V6050)


__e.TailApply(PrimFunc(symshen_4gc), V6047, tmp18728)
return


} else {
__e.Return(False)
return
}


} else {
__e.Return(let__18612)
return
}


}, 7)

tmp18733 := Call(__e, ns2_1set, symshen_4variants_2, tmp18716)


_ = tmp18733

tmp18734 := MakeNative(func(__e *ControlFlow) {
V6053 := __e.Get(1)
_ = V6053
tmp18735 := Call(__e, PrimFunc(symgensym), symB)


let__18614 := tmp18735
_ = let__18614

tmp18736 := Call(__e, PrimFunc(symgensym), symL)


let__18615 := tmp18736
_ = let__18615

tmp18737 := Call(__e, PrimFunc(symgensym), symKey)


let__18616 := tmp18737
_ = let__18616

tmp18738 := Call(__e, PrimFunc(symgensym), symC)


let__18617 := tmp18738
_ = let__18617

tmp18739 := Call(__e, PrimFunc(symgensym), symV)


let__18618 := tmp18739
_ = let__18618

tmp18740 := Call(__e, PrimFunc(symshen_4extract_1vars), V6053)


let__18619 := tmp18740
_ = let__18619

tmp18741 := Call(__e, PrimFunc(symshen_4rcons__form), V6053)


tmp18742 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18617, Nil)
}
__typedArg0 := let__18617
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18743 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18616, tmp18742)
}
__typedArg0 := let__18616
__typedArg1 := tmp18742
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18744 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18615, tmp18743)
}
__typedArg0 := let__18615
__typedArg1 := tmp18743
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18745 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18614, tmp18744)
}
__typedArg0 := let__18614
__typedArg1 := tmp18744
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18746 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18741, tmp18745)
}
__typedArg0 := tmp18741
__typedArg1 := tmp18745
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18747 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18618, tmp18746)
}
__typedArg0 := let__18618
__typedArg1 := tmp18746
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18748 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symis_b, tmp18747)
}
__typedArg0 := symis_b
__typedArg1 := tmp18747
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18749 := Call(__e, PrimFunc(symshen_4stpart), let__18619, tmp18748, let__18614)


tmp18750 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18749, Nil)
}
__typedArg0 := tmp18749
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18751 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18617, tmp18750)
}
__typedArg0 := let__18617
__typedArg1 := tmp18750
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18752 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp18751)
}
__typedArg0 := symlambda
__typedArg1 := tmp18751
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18753 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18752, Nil)
}
__typedArg0 := tmp18752
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18754 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18616, tmp18753)
}
__typedArg0 := let__18616
__typedArg1 := tmp18753
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18755 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp18754)
}
__typedArg0 := symlambda
__typedArg1 := tmp18754
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18756 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18755, Nil)
}
__typedArg0 := tmp18755
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18757 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18615, tmp18756)
}
__typedArg0 := let__18615
__typedArg1 := tmp18756
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18758 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp18757)
}
__typedArg0 := symlambda
__typedArg1 := tmp18757
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18759 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18758, Nil)
}
__typedArg0 := tmp18758
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18760 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18614, tmp18759)
}
__typedArg0 := let__18614
__typedArg1 := tmp18759
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18761 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp18760)
}
__typedArg0 := symlambda
__typedArg1 := tmp18760
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18762 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18761, Nil)
}
__typedArg0 := tmp18761
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18763 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__18618, tmp18762)
}
__typedArg0 := let__18618
__typedArg1 := tmp18762
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp18763)
}
__typedArg0 := symlambda
__typedArg1 := tmp18763
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp18764 := Call(__e, ns2_1set, symshen_4prolog_1abstraction, tmp18734)


_ = tmp18764

tmp18765 := MakeNative(func(__e *ControlFlow) {
V6060 := __e.Get(1)
_ = V6060
tmp18766 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_ddemodulation_1function_d)
}
__typedArg0 := symshen_4_ddemodulation_1function_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

let__18620 := tmp18766
_ = let__18620

__e.TailApply(let__18620, V6060)
return


}, 1)

__e.TailApply(ns2_1set, symshen_4demod, tmp18765)
return




}, 0)

