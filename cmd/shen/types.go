package main

import . "github.com/pyrex41/shen-go/kl"

var TypesMain = MakeNative(func(__e *ControlFlow) {
tmp18654 := MakeNative(func(__e *ControlFlow) {
V5507 := __e.Get(1)
_ = V5507
V5508 := __e.Get(2)
_ = V5508
tmp18655 := MakeNative(func(__e *ControlFlow) {
W5509 := __e.Get(1)
_ = W5509
tmp18656 := MakeNative(func(__e *ControlFlow) {
W5510 := __e.Get(1)
_ = W5510
tmp18657 := MakeNative(func(__e *ControlFlow) {
W5515 := __e.Get(1)
_ = W5515
tmp18658 := MakeNative(func(__e *ControlFlow) {
W5516 := __e.Get(1)
_ = W5516
__e.Return(V5507)
return
}, 1)

tmp18659 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp18660 := Call(__e, PrimFunc(symshen_4assoc_1_6), V5507, W5515, tmp18659)


tmp18661 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp18660)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp18660
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

__e.TailApply(tmp18658, tmp18661)
return


}, 1)

tmp18662 := Call(__e, PrimFunc(symshen_4prolog_1abstraction), V5508)


tmp18663 := Call(__e, PrimFunc(symeval_1kl), tmp18662)


__e.TailApply(tmp18657, tmp18663)
return


}, 1)

tmp18664 := MakeNative(func(__e *ControlFlow) {
Z5511 := __e.Get(1)
_ = Z5511
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5512 := __e.Get(1)
_ = Z5512
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5513 := __e.Get(1)
_ = Z5513
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5514 := __e.Get(1)
_ = Z5514
tmp18665 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp18665

tmp18666 := Call(__e, PrimFunc(symshen_4deref), V5507, Z5511)


tmp18667 := Call(__e, PrimFunc(symreceive), tmp18666)


tmp18668 := Call(__e, PrimFunc(symshen_4deref), W5509, Z5511)


tmp18669 := Call(__e, PrimFunc(symreceive), tmp18668)


__e.TailApply(PrimFunc(symshen_4variancy), tmp18667, tmp18669, Z5511, Z5512, Z5513, Z5514)
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


tmp18672 := Call(__e, PrimFunc(symvector), MakeNumber(0))


tmp18673 := Call(__e, PrimFunc(sym_8v), MakeNumber(0), tmp18672)


tmp18674 := Call(__e, PrimFunc(sym_8v), True, tmp18673)


tmp18675 := Call(__e, tmp18671, tmp18674)


tmp18676 := Call(__e, tmp18675, MakeNumber(0))


tmp18677 := MakeNative(func(__e *ControlFlow) {
__e.Return(True)
return
}, 0)

tmp18678 := Call(__e, tmp18676, tmp18677)


__e.TailApply(tmp18656, tmp18678)
return


}, 1)

tmp18679 := Call(__e, PrimFunc(symshen_4rectify_1type), V5508)


__e.TailApply(tmp18655, tmp18679)
return


}, 2)

tmp18680 := Call(__e, ns2_1set, symdeclare, tmp18654)


_ = tmp18680

tmp18681 := MakeNative(func(__e *ControlFlow) {
V5517 := __e.Get(1)
_ = V5517
V5518 := __e.Get(2)
_ = V5518
V5519 := __e.Get(3)
_ = V5519
V5520 := __e.Get(4)
_ = V5520
V5521 := __e.Get(5)
_ = V5521
V5522 := __e.Get(6)
_ = V5522
tmp18682 := MakeNative(func(__e *ControlFlow) {
W5523 := __e.Get(1)
_ = W5523
tmp18683 := MakeNative(func(__e *ControlFlow) {
W5524 := __e.Get(1)
_ = W5524
tmp18698 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W5524, False)
}
__typedArg0 := W5524
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp18698 {
tmp18684 := MakeNative(func(__e *ControlFlow) {
W5539 := __e.Get(1)
_ = W5539
tmp18686 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W5539, False)
}
__typedArg0 := W5539
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp18686 {
__e.TailApply(PrimFunc(symshen_4unlock), V5520, W5523)
return
} else {
__e.Return(W5539)
return
}


}, 1)

tmp18696 := Call(__e, PrimFunc(symshen_4unlocked_2), V5520)


var ifres18687 Obj

if True == tmp18696 {
tmp18688 := MakeNative(func(__e *ControlFlow) {
W5540 := __e.Get(1)
_ = W5540
tmp18689 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp18689

tmp18690 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5517, Nil)
}
__typedArg0 := V5517
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18691 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfn, tmp18690)
}
__typedArg0 := symfn
__typedArg1 := tmp18690
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18692 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4variants_2), V5517, W5540, V5518, V5519, V5520, W5523, V5522)
return
}, 0)

tmp18693 := Call(__e, PrimFunc(symshen_4system_1S_1h), tmp18691, W5540, Nil, V5519, V5520, W5523, tmp18692)


__e.TailApply(PrimFunc(symshen_4gc), V5519, tmp18693)
return


}, 1)

tmp18694 := Call(__e, PrimFunc(symshen_4newpv), V5519)


tmp18695 := Call(__e, tmp18688, tmp18694)


ifres18687 = tmp18695


} else {
ifres18687 = False


}

__e.TailApply(tmp18684, ifres18687)
return


} else {
__e.Return(W5524)
return
}


}, 1)

tmp18761 := Call(__e, PrimFunc(symshen_4unlocked_2), V5520)


var ifres18699 Obj

if True == tmp18761 {
tmp18700 := MakeNative(func(__e *ControlFlow) {
W5525 := __e.Get(1)
_ = W5525
tmp18701 := MakeNative(func(__e *ControlFlow) {
W5526 := __e.Get(1)
_ = W5526
tmp18745 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W5525)
}
__typedArg0 := W5525
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp18745 {
tmp18702 := MakeNative(func(__e *ControlFlow) {
W5529 := __e.Get(1)
_ = W5529
tmp18703 := MakeNative(func(__e *ControlFlow) {
W5530 := __e.Get(1)
_ = W5530
tmp18707 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W5529, sym_1_1_6)
}
__typedArg0 := W5529
__typedArg1 := sym_1_1_6
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp18707 {
__e.TailApply(PrimFunc(symthaw), W5530)
return
} else {
tmp18705 := Call(__e, PrimFunc(symshen_4pvar_2), W5529)


if True == tmp18705 {
__e.TailApply(PrimFunc(symshen_4bind_b), W5529, sym_1_1_6, V5519, W5530)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp18708 := MakeNative(func(__e *ControlFlow) {
tmp18709 := MakeNative(func(__e *ControlFlow) {
W5531 := __e.Get(1)
_ = W5531
tmp18710 := MakeNative(func(__e *ControlFlow) {
W5532 := __e.Get(1)
_ = W5532
tmp18730 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W5531)
}
__typedArg0 := W5531
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp18730 {
tmp18711 := MakeNative(func(__e *ControlFlow) {
W5534 := __e.Get(1)
_ = W5534
tmp18712 := MakeNative(func(__e *ControlFlow) {
W5535 := __e.Get(1)
_ = W5535
tmp18713 := MakeNative(func(__e *ControlFlow) {
W5536 := __e.Get(1)
_ = W5536
tmp18717 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W5535, Nil)
}
__typedArg0 := W5535
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp18717 {
__e.TailApply(PrimFunc(symthaw), W5536)
return
} else {
tmp18715 := Call(__e, PrimFunc(symshen_4pvar_2), W5535)


if True == tmp18715 {
__e.TailApply(PrimFunc(symshen_4bind_b), W5535, Nil, V5519, W5536)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp18718 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W5532, W5534)
return
}, 0)

__e.TailApply(tmp18713, tmp18718)
return


}, 1)

tmp18719 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W5531)
}
__typedArg0 := W5531
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18720 := Call(__e, PrimFunc(symshen_4lazyderef), tmp18719, V5519)


__e.TailApply(tmp18712, tmp18720)
return


}, 1)

tmp18721 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W5531)
}
__typedArg0 := W5531
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp18711, tmp18721)
return


} else {
tmp18728 := Call(__e, PrimFunc(symshen_4pvar_2), W5531)


if True == tmp18728 {
tmp18722 := MakeNative(func(__e *ControlFlow) {
W5537 := __e.Get(1)
_ = W5537
tmp18723 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W5537, Nil)
}
__typedArg0 := W5537
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18724 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W5532, W5537)
return
}, 0)

tmp18725 := Call(__e, PrimFunc(symshen_4bind_b), W5531, tmp18723, V5519, tmp18724)


__e.TailApply(PrimFunc(symshen_4gc), V5519, tmp18725)
return


}, 1)

tmp18726 := Call(__e, PrimFunc(symshen_4newpv), V5519)


__e.TailApply(tmp18722, tmp18726)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp18731 := MakeNative(func(__e *ControlFlow) {
Z5533 := __e.Get(1)
_ = Z5533
__e.TailApply(W5526, Z5533)
return
}, 1)

__e.TailApply(tmp18710, tmp18731)
return


}, 1)

tmp18732 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W5525)
}
__typedArg0 := W5525
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18733 := Call(__e, PrimFunc(symshen_4lazyderef), tmp18732, V5519)


__e.TailApply(tmp18709, tmp18733)
return


}, 0)

__e.TailApply(tmp18703, tmp18708)
return


}, 1)

tmp18734 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W5525)
}
__typedArg0 := W5525
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18735 := Call(__e, PrimFunc(symshen_4lazyderef), tmp18734, V5519)


__e.TailApply(tmp18702, tmp18735)
return


} else {
tmp18743 := Call(__e, PrimFunc(symshen_4pvar_2), W5525)


if True == tmp18743 {
tmp18736 := MakeNative(func(__e *ControlFlow) {
W5538 := __e.Get(1)
_ = W5538
tmp18737 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W5538, Nil)
}
__typedArg0 := W5538
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18738 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp18737)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp18737
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18739 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W5526, W5538)
return
}, 0)

tmp18740 := Call(__e, PrimFunc(symshen_4bind_b), W5525, tmp18738, V5519, tmp18739)


__e.TailApply(PrimFunc(symshen_4gc), V5519, tmp18740)
return


}, 1)

tmp18741 := Call(__e, PrimFunc(symshen_4newpv), V5519)


__e.TailApply(tmp18736, tmp18741)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp18746 := MakeNative(func(__e *ControlFlow) {
Z5527 := __e.Get(1)
_ = Z5527
tmp18747 := MakeNative(func(__e *ControlFlow) {
W5528 := __e.Get(1)
_ = W5528
tmp18748 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp18748

tmp18749 := Call(__e, PrimFunc(symshen_4deref), V5517, V5519)


tmp18750 := Call(__e, PrimFunc(symarity), tmp18749)


tmp18751 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp18750, MakeNumber(0))
}
__typedArg0 := tmp18750
__typedArg1 := MakeNumber(0)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

tmp18752 := MakeNative(func(__e *ControlFlow) {
tmp18753 := MakeNative(func(__e *ControlFlow) {
tmp18754 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5517, Nil)
}
__typedArg0 := V5517
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18755 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfn, tmp18754)
}
__typedArg0 := symfn
__typedArg1 := tmp18754
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18756 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4variants_2), V5517, Z5527, W5528, V5519, V5520, W5523, V5522)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), tmp18755, W5528, Nil, V5519, V5520, W5523, tmp18756)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V5519, V5520, W5523, tmp18753)
return


}, 0)

tmp18757 := Call(__e, PrimFunc(symwhen), tmp18751, V5519, V5520, W5523, tmp18752)


__e.TailApply(PrimFunc(symshen_4gc), V5519, tmp18757)
return


}, 1)

tmp18758 := Call(__e, PrimFunc(symshen_4newpv), V5519)


__e.TailApply(tmp18747, tmp18758)
return


}, 1)

__e.TailApply(tmp18701, tmp18746)
return


}, 1)

tmp18759 := Call(__e, PrimFunc(symshen_4lazyderef), V5518, V5519)


tmp18760 := Call(__e, tmp18700, tmp18759)


ifres18699 = tmp18760


} else {
ifres18699 = False


}

__e.TailApply(tmp18683, ifres18699)
return


}, 1)

__e.TailApply(tmp18682, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V5521)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V5521
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())
return


}, 6)

tmp18763 := Call(__e, ns2_1set, symshen_4variancy, tmp18681)


_ = tmp18763

tmp18764 := MakeNative(func(__e *ControlFlow) {
V5541 := __e.Get(1)
_ = V5541
V5542 := __e.Get(2)
_ = V5542
V5543 := __e.Get(3)
_ = V5543
V5544 := __e.Get(4)
_ = V5544
V5545 := __e.Get(5)
_ = V5545
V5546 := __e.Get(6)
_ = V5546
V5547 := __e.Get(7)
_ = V5547
tmp18765 := MakeNative(func(__e *ControlFlow) {
W5548 := __e.Get(1)
_ = W5548
tmp18778 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W5548, False)
}
__typedArg0 := W5548
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp18778 {
tmp18776 := Call(__e, PrimFunc(symshen_4unlocked_2), V5545)


if True == tmp18776 {
tmp18766 := MakeNative(func(__e *ControlFlow) {
W5549 := __e.Get(1)
_ = W5549
tmp18767 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp18767

tmp18768 := Call(__e, PrimFunc(symshen_4deref), V5541, V5544)


tmp18769 := Call(__e, PrimFunc(symshen_4app), tmp18768, MakeString(" may create errors\n"), symshen_4a)


tmp18770 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("warning: changing the type of "))
__typedS1, __typedOK1 := TypedString(tmp18769)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("warning: changing the type of ")
__typedArg1 := tmp18769
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp18771 := Call(__e, PrimFunc(symstoutput))


tmp18772 := Call(__e, PrimFunc(sympr), tmp18770, tmp18771)


tmp18773 := Call(__e, PrimFunc(symis), W5549, tmp18772, V5544, V5545, V5546, V5547)


__e.TailApply(PrimFunc(symshen_4gc), V5544, tmp18773)
return


}, 1)

tmp18774 := Call(__e, PrimFunc(symshen_4newpv), V5544)


__e.TailApply(tmp18766, tmp18774)
return


} else {
__e.Return(False)
return
}


} else {
__e.Return(W5548)
return
}


}, 1)

tmp18782 := Call(__e, PrimFunc(symshen_4unlocked_2), V5545)


var ifres18779 Obj

if True == tmp18782 {
tmp18780 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp18780

tmp18781 := Call(__e, PrimFunc(symis_b), V5542, V5543, V5544, V5545, V5546, V5547)


ifres18779 = tmp18781


} else {
ifres18779 = False


}

__e.TailApply(tmp18765, ifres18779)
return


}, 7)

tmp18783 := Call(__e, ns2_1set, symshen_4variants_2, tmp18764)


_ = tmp18783

tmp18784 := MakeNative(func(__e *ControlFlow) {
V5550 := __e.Get(1)
_ = V5550
tmp18785 := MakeNative(func(__e *ControlFlow) {
W5551 := __e.Get(1)
_ = W5551
tmp18786 := MakeNative(func(__e *ControlFlow) {
W5552 := __e.Get(1)
_ = W5552
tmp18787 := MakeNative(func(__e *ControlFlow) {
W5553 := __e.Get(1)
_ = W5553
tmp18788 := MakeNative(func(__e *ControlFlow) {
W5554 := __e.Get(1)
_ = W5554
tmp18789 := MakeNative(func(__e *ControlFlow) {
W5555 := __e.Get(1)
_ = W5555
tmp18790 := MakeNative(func(__e *ControlFlow) {
W5556 := __e.Get(1)
_ = W5556
tmp18791 := Call(__e, PrimFunc(symshen_4rcons__form), V5550)


tmp18792 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W5554, Nil)
}
__typedArg0 := W5554
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18793 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W5553, tmp18792)
}
__typedArg0 := W5553
__typedArg1 := tmp18792
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18794 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W5552, tmp18793)
}
__typedArg0 := W5552
__typedArg1 := tmp18793
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18795 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W5551, tmp18794)
}
__typedArg0 := W5551
__typedArg1 := tmp18794
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18796 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18791, tmp18795)
}
__typedArg0 := tmp18791
__typedArg1 := tmp18795
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18797 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W5555, tmp18796)
}
__typedArg0 := W5555
__typedArg1 := tmp18796
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18798 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symis_b, tmp18797)
}
__typedArg0 := symis_b
__typedArg1 := tmp18797
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18799 := Call(__e, PrimFunc(symshen_4stpart), W5556, tmp18798, W5551)


tmp18800 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18799, Nil)
}
__typedArg0 := tmp18799
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18801 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W5554, tmp18800)
}
__typedArg0 := W5554
__typedArg1 := tmp18800
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18802 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp18801)
}
__typedArg0 := symlambda
__typedArg1 := tmp18801
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18803 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18802, Nil)
}
__typedArg0 := tmp18802
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18804 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W5553, tmp18803)
}
__typedArg0 := W5553
__typedArg1 := tmp18803
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18805 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp18804)
}
__typedArg0 := symlambda
__typedArg1 := tmp18804
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18806 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18805, Nil)
}
__typedArg0 := tmp18805
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18807 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W5552, tmp18806)
}
__typedArg0 := W5552
__typedArg1 := tmp18806
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18808 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp18807)
}
__typedArg0 := symlambda
__typedArg1 := tmp18807
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18809 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18808, Nil)
}
__typedArg0 := tmp18808
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18810 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W5551, tmp18809)
}
__typedArg0 := W5551
__typedArg1 := tmp18809
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18811 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp18810)
}
__typedArg0 := symlambda
__typedArg1 := tmp18810
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18812 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18811, Nil)
}
__typedArg0 := tmp18811
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18813 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W5555, tmp18812)
}
__typedArg0 := W5555
__typedArg1 := tmp18812
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp18813)
}
__typedArg0 := symlambda
__typedArg1 := tmp18813
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp18814 := Call(__e, PrimFunc(symshen_4extract_1vars), V5550)


__e.TailApply(tmp18790, tmp18814)
return


}, 1)

tmp18815 := Call(__e, PrimFunc(symgensym), symV)


__e.TailApply(tmp18789, tmp18815)
return


}, 1)

tmp18816 := Call(__e, PrimFunc(symgensym), symC)


__e.TailApply(tmp18788, tmp18816)
return


}, 1)

tmp18817 := Call(__e, PrimFunc(symgensym), symKey)


__e.TailApply(tmp18787, tmp18817)
return


}, 1)

tmp18818 := Call(__e, PrimFunc(symgensym), symL)


__e.TailApply(tmp18786, tmp18818)
return


}, 1)

tmp18819 := Call(__e, PrimFunc(symgensym), symB)


__e.TailApply(tmp18785, tmp18819)
return


}, 1)

tmp18820 := Call(__e, ns2_1set, symshen_4prolog_1abstraction, tmp18784)


_ = tmp18820

tmp18821 := MakeNative(func(__e *ControlFlow) {
V5557 := __e.Get(1)
_ = V5557
__e.Return(V5557)
return
}, 1)

tmp18822 := Call(__e, ns2_1set, symshen_4demod, tmp18821)


_ = tmp18822

tmp18823 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18824 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp18823)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp18823
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18825 := Call(__e, PrimFunc(symdeclare), symabort, tmp18824)


_ = tmp18825

tmp18826 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18827 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp18826)
}
__typedArg0 := symlist
__typedArg1 := tmp18826
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18828 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18827, Nil)
}
__typedArg0 := tmp18827
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18829 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp18828)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp18828
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18830 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp18829)
}
__typedArg0 := symstring
__typedArg1 := tmp18829
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18831 := Call(__e, PrimFunc(symdeclare), symabsolute, tmp18830)


_ = tmp18831

tmp18832 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18833 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp18832)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp18832
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18834 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp18833)
}
__typedArg0 := symA
__typedArg1 := tmp18833
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18835 := Call(__e, PrimFunc(symdeclare), symabsvector_2, tmp18834)


_ = tmp18835

tmp18836 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18837 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp18836)
}
__typedArg0 := symlist
__typedArg1 := tmp18836
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18838 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18839 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp18838)
}
__typedArg0 := symlist
__typedArg1 := tmp18838
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18840 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18839, Nil)
}
__typedArg0 := tmp18839
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18841 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp18840)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp18840
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18842 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18837, tmp18841)
}
__typedArg0 := tmp18837
__typedArg1 := tmp18841
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18843 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18842, Nil)
}
__typedArg0 := tmp18842
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18844 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp18843)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp18843
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18845 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp18844)
}
__typedArg0 := symA
__typedArg1 := tmp18844
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18846 := Call(__e, PrimFunc(symdeclare), symadjoin, tmp18845)


_ = tmp18846

tmp18847 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18848 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp18847)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp18847
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18849 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, tmp18848)
}
__typedArg0 := symboolean
__typedArg1 := tmp18848
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18850 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18849, Nil)
}
__typedArg0 := tmp18849
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18851 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp18850)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp18850
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18852 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, tmp18851)
}
__typedArg0 := symboolean
__typedArg1 := tmp18851
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18853 := Call(__e, PrimFunc(symdeclare), symand, tmp18852)


_ = tmp18853

tmp18854 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18855 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp18854)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp18854
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18856 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp18855)
}
__typedArg0 := symsymbol
__typedArg1 := tmp18855
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18857 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18856, Nil)
}
__typedArg0 := tmp18856
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18858 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp18857)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp18857
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18859 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp18858)
}
__typedArg0 := symstring
__typedArg1 := tmp18858
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18860 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18859, Nil)
}
__typedArg0 := tmp18859
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18861 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp18860)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp18860
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18862 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp18861)
}
__typedArg0 := symA
__typedArg1 := tmp18861
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18863 := Call(__e, PrimFunc(symdeclare), symshen_4app, tmp18862)


_ = tmp18863

tmp18864 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18865 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp18864)
}
__typedArg0 := symlist
__typedArg1 := tmp18864
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18866 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18867 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp18866)
}
__typedArg0 := symlist
__typedArg1 := tmp18866
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18868 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18869 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp18868)
}
__typedArg0 := symlist
__typedArg1 := tmp18868
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18870 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18869, Nil)
}
__typedArg0 := tmp18869
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18871 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp18870)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp18870
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18872 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18867, tmp18871)
}
__typedArg0 := tmp18867
__typedArg1 := tmp18871
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18873 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18872, Nil)
}
__typedArg0 := tmp18872
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18874 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp18873)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp18873
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18875 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18865, tmp18874)
}
__typedArg0 := tmp18865
__typedArg1 := tmp18874
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18876 := Call(__e, PrimFunc(symdeclare), symappend, tmp18875)


_ = tmp18876

tmp18877 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18878 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp18877)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp18877
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18879 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp18878)
}
__typedArg0 := symA
__typedArg1 := tmp18878
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18880 := Call(__e, PrimFunc(symdeclare), symarity, tmp18879)


_ = tmp18880

tmp18881 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18882 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp18881)
}
__typedArg0 := symlist
__typedArg1 := tmp18881
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18883 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18882, Nil)
}
__typedArg0 := tmp18882
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18884 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp18883)
}
__typedArg0 := symlist
__typedArg1 := tmp18883
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18885 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18886 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp18885)
}
__typedArg0 := symlist
__typedArg1 := tmp18885
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18887 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18886, Nil)
}
__typedArg0 := tmp18886
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18888 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp18887)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp18887
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18889 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18884, tmp18888)
}
__typedArg0 := tmp18884
__typedArg1 := tmp18888
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18890 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18889, Nil)
}
__typedArg0 := tmp18889
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18891 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp18890)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp18890
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18892 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp18891)
}
__typedArg0 := symA
__typedArg1 := tmp18891
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18893 := Call(__e, PrimFunc(symdeclare), symassoc, tmp18892)


_ = tmp18893

tmp18894 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18895 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp18894)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp18894
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18896 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp18895)
}
__typedArg0 := symA
__typedArg1 := tmp18895
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18897 := Call(__e, PrimFunc(symdeclare), symatom_2, tmp18896)


_ = tmp18897

tmp18898 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18899 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp18898)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp18898
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18900 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp18899)
}
__typedArg0 := symA
__typedArg1 := tmp18899
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18901 := Call(__e, PrimFunc(symdeclare), symboolean_2, tmp18900)


_ = tmp18901

tmp18902 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18903 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp18902)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp18902
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18904 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp18903)
}
__typedArg0 := symstring
__typedArg1 := tmp18903
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18905 := Call(__e, PrimFunc(symdeclare), symbootstrap, tmp18904)


_ = tmp18905

tmp18906 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18907 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp18906)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp18906
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18908 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp18907)
}
__typedArg0 := symsymbol
__typedArg1 := tmp18907
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18909 := Call(__e, PrimFunc(symdeclare), symbound_2, tmp18908)


_ = tmp18909

tmp18910 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18911 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp18910)
}
__typedArg0 := symlist
__typedArg1 := tmp18910
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18912 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18913 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp18912)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp18912
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18914 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18911, tmp18913)
}
__typedArg0 := tmp18911
__typedArg1 := tmp18913
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18915 := Call(__e, PrimFunc(symdeclare), symshen_4ccons_2, tmp18914)


_ = tmp18915

tmp18916 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18917 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp18916)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp18916
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18918 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp18917)
}
__typedArg0 := symstring
__typedArg1 := tmp18917
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18919 := Call(__e, PrimFunc(symdeclare), symcd, tmp18918)


_ = tmp18919

tmp18920 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18921 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp18920)
}
__typedArg0 := symstream
__typedArg1 := tmp18920
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18922 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symB, Nil)
}
__typedArg0 := symB
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18923 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp18922)
}
__typedArg0 := symlist
__typedArg1 := tmp18922
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18924 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18923, Nil)
}
__typedArg0 := tmp18923
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18925 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp18924)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp18924
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18926 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18921, tmp18925)
}
__typedArg0 := tmp18921
__typedArg1 := tmp18925
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18927 := Call(__e, PrimFunc(symdeclare), symclose, tmp18926)


_ = tmp18927

tmp18928 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18929 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp18928)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp18928
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18930 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp18929)
}
__typedArg0 := symstring
__typedArg1 := tmp18929
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18931 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18930, Nil)
}
__typedArg0 := tmp18930
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18932 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp18931)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp18931
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18933 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp18932)
}
__typedArg0 := symstring
__typedArg1 := tmp18932
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18934 := Call(__e, PrimFunc(symdeclare), symcn, tmp18933)


_ = tmp18934

tmp18935 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18936 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp18935)
}
__typedArg0 := symlist
__typedArg1 := tmp18935
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18937 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18938 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp18937)
}
__typedArg0 := symlist
__typedArg1 := tmp18937
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18939 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symB, Nil)
}
__typedArg0 := symB
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18940 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18938, tmp18939)
}
__typedArg0 := tmp18938
__typedArg1 := tmp18939
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18941 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp18940)
}
__typedArg0 := symstr
__typedArg1 := tmp18940
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18942 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18941, Nil)
}
__typedArg0 := tmp18941
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18943 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp18942)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp18942
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18944 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18936, tmp18943)
}
__typedArg0 := tmp18936
__typedArg1 := tmp18943
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18945 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18946 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp18945)
}
__typedArg0 := symlist
__typedArg1 := tmp18945
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18947 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symB, Nil)
}
__typedArg0 := symB
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18948 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp18947)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp18947
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18949 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18946, tmp18948)
}
__typedArg0 := tmp18946
__typedArg1 := tmp18948
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18950 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18949, Nil)
}
__typedArg0 := tmp18949
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18951 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp18950)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp18950
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18952 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18944, tmp18951)
}
__typedArg0 := tmp18944
__typedArg1 := tmp18951
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18953 := Call(__e, PrimFunc(symdeclare), symcompile, tmp18952)


_ = tmp18953

tmp18954 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18955 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp18954)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp18954
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18956 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp18955)
}
__typedArg0 := symA
__typedArg1 := tmp18955
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18957 := Call(__e, PrimFunc(symdeclare), symcons_2, tmp18956)


_ = tmp18957

tmp18958 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18959 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp18958)
}
__typedArg0 := symlist
__typedArg1 := tmp18958
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18960 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18959, Nil)
}
__typedArg0 := tmp18959
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18961 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp18960)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp18960
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18962 := Call(__e, PrimFunc(symdeclare), symdatatypes, tmp18961)


_ = tmp18962

tmp18963 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18964 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp18963)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp18963
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18965 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp18964)
}
__typedArg0 := symsymbol
__typedArg1 := tmp18964
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18966 := Call(__e, PrimFunc(symdeclare), symdestroy, tmp18965)


_ = tmp18966

tmp18967 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18968 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp18967)
}
__typedArg0 := symlist
__typedArg1 := tmp18967
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18969 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18970 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp18969)
}
__typedArg0 := symlist
__typedArg1 := tmp18969
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18971 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18972 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp18971)
}
__typedArg0 := symlist
__typedArg1 := tmp18971
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18973 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18972, Nil)
}
__typedArg0 := tmp18972
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18974 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp18973)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp18973
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18975 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18970, tmp18974)
}
__typedArg0 := tmp18970
__typedArg1 := tmp18974
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18976 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18975, Nil)
}
__typedArg0 := tmp18975
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18977 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp18976)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp18976
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18978 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18968, tmp18977)
}
__typedArg0 := tmp18968
__typedArg1 := tmp18977
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18979 := Call(__e, PrimFunc(symdeclare), symdifference, tmp18978)


_ = tmp18979

tmp18980 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symB, Nil)
}
__typedArg0 := symB
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18981 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp18980)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp18980
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18982 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symB, tmp18981)
}
__typedArg0 := symB
__typedArg1 := tmp18981
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18983 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18982, Nil)
}
__typedArg0 := tmp18982
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18984 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp18983)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp18983
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18985 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp18984)
}
__typedArg0 := symA
__typedArg1 := tmp18984
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18986 := Call(__e, PrimFunc(symdeclare), symdo, tmp18985)


_ = tmp18986

tmp18987 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18988 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp18987)
}
__typedArg0 := symlist
__typedArg1 := tmp18987
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18989 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18990 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp18989)
}
__typedArg0 := symlist
__typedArg1 := tmp18989
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18991 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symB, Nil)
}
__typedArg0 := symB
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18992 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp18991)
}
__typedArg0 := symlist
__typedArg1 := tmp18991
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18993 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18992, Nil)
}
__typedArg0 := tmp18992
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18994 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18990, tmp18993)
}
__typedArg0 := tmp18990
__typedArg1 := tmp18993
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18995 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp18994)
}
__typedArg0 := symstr
__typedArg1 := tmp18994
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18996 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18995, Nil)
}
__typedArg0 := tmp18995
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18997 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp18996)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp18996
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18998 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18988, tmp18997)
}
__typedArg0 := tmp18988
__typedArg1 := tmp18997
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18999 := Call(__e, PrimFunc(symdeclare), sym_5e_6, tmp18998)


_ = tmp18999

tmp19000 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19001 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19000)
}
__typedArg0 := symlist
__typedArg1 := tmp19000
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19002 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symB, Nil)
}
__typedArg0 := symB
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19003 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19002)
}
__typedArg0 := symlist
__typedArg1 := tmp19002
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19004 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19005 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19004)
}
__typedArg0 := symlist
__typedArg1 := tmp19004
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19006 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19005, Nil)
}
__typedArg0 := tmp19005
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19007 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19003, tmp19006)
}
__typedArg0 := tmp19003
__typedArg1 := tmp19006
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19008 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp19007)
}
__typedArg0 := symstr
__typedArg1 := tmp19007
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19009 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19008, Nil)
}
__typedArg0 := tmp19008
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19010 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19009)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19009
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19011 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19001, tmp19010)
}
__typedArg0 := tmp19001
__typedArg1 := tmp19010
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19012 := Call(__e, PrimFunc(symdeclare), sym_5_b_6, tmp19011)


_ = tmp19012

tmp19013 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19014 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19013)
}
__typedArg0 := symlist
__typedArg1 := tmp19013
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19015 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19016 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19015)
}
__typedArg0 := symlist
__typedArg1 := tmp19015
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19017 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symB, Nil)
}
__typedArg0 := symB
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19018 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19017)
}
__typedArg0 := symlist
__typedArg1 := tmp19017
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19019 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19018, Nil)
}
__typedArg0 := tmp19018
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19020 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19016, tmp19019)
}
__typedArg0 := tmp19016
__typedArg1 := tmp19019
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19021 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp19020)
}
__typedArg0 := symstr
__typedArg1 := tmp19020
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19022 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19021, Nil)
}
__typedArg0 := tmp19021
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19023 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19022)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19022
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19024 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19014, tmp19023)
}
__typedArg0 := tmp19014
__typedArg1 := tmp19023
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19025 := Call(__e, PrimFunc(symdeclare), sym_5end_6, tmp19024)


_ = tmp19025

tmp19026 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19027 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19026)
}
__typedArg0 := symlist
__typedArg1 := tmp19026
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19028 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symB, Nil)
}
__typedArg0 := symB
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19029 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19027, tmp19028)
}
__typedArg0 := tmp19027
__typedArg1 := tmp19028
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19030 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp19029)
}
__typedArg0 := symstr
__typedArg1 := tmp19029
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19031 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19032 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19031)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19031
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19033 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19030, tmp19032)
}
__typedArg0 := tmp19030
__typedArg1 := tmp19032
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19034 := Call(__e, PrimFunc(symdeclare), symshen_4parse_1failure_2, tmp19033)


_ = tmp19034

tmp19035 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19036 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19035)
}
__typedArg0 := symlist
__typedArg1 := tmp19035
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19037 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symB, Nil)
}
__typedArg0 := symB
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19038 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19036, tmp19037)
}
__typedArg0 := tmp19036
__typedArg1 := tmp19037
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19039 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp19038)
}
__typedArg0 := symstr
__typedArg1 := tmp19038
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19040 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19039, Nil)
}
__typedArg0 := tmp19039
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19041 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19040)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19040
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19042 := Call(__e, PrimFunc(symdeclare), symshen_4parse_1failure, tmp19041)


_ = tmp19042

tmp19043 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19044 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19043)
}
__typedArg0 := symlist
__typedArg1 := tmp19043
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19045 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symB, Nil)
}
__typedArg0 := symB
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19046 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19044, tmp19045)
}
__typedArg0 := tmp19044
__typedArg1 := tmp19045
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19047 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp19046)
}
__typedArg0 := symstr
__typedArg1 := tmp19046
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19048 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symB, Nil)
}
__typedArg0 := symB
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19049 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19048)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19048
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19050 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19047, tmp19049)
}
__typedArg0 := tmp19047
__typedArg1 := tmp19049
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19051 := Call(__e, PrimFunc(symdeclare), symshen_4_5_1out, tmp19050)


_ = tmp19051

tmp19052 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19053 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19052)
}
__typedArg0 := symlist
__typedArg1 := tmp19052
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19054 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symB, Nil)
}
__typedArg0 := symB
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19055 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19053, tmp19054)
}
__typedArg0 := tmp19053
__typedArg1 := tmp19054
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19056 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp19055)
}
__typedArg0 := symstr
__typedArg1 := tmp19055
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19057 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19058 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19057)
}
__typedArg0 := symlist
__typedArg1 := tmp19057
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19059 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19058, Nil)
}
__typedArg0 := tmp19058
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19060 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19059)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19059
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19061 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19056, tmp19060)
}
__typedArg0 := tmp19056
__typedArg1 := tmp19060
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19062 := Call(__e, PrimFunc(symdeclare), symshen_4in_1_6, tmp19061)


_ = tmp19062

tmp19063 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19064 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19063)
}
__typedArg0 := symlist
__typedArg1 := tmp19063
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19065 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19066 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19065)
}
__typedArg0 := symlist
__typedArg1 := tmp19065
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19067 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symB, Nil)
}
__typedArg0 := symB
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19068 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19066, tmp19067)
}
__typedArg0 := tmp19066
__typedArg1 := tmp19067
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19069 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp19068)
}
__typedArg0 := symstr
__typedArg1 := tmp19068
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19070 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19069, Nil)
}
__typedArg0 := tmp19069
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19071 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19070)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19070
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19072 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symB, tmp19071)
}
__typedArg0 := symB
__typedArg1 := tmp19071
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19073 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19072, Nil)
}
__typedArg0 := tmp19072
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19074 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19073)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19073
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19075 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19064, tmp19074)
}
__typedArg0 := tmp19064
__typedArg1 := tmp19074
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19076 := Call(__e, PrimFunc(symdeclare), symshen_4comb, tmp19075)


_ = tmp19076

tmp19077 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19078 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19077)
}
__typedArg0 := symlist
__typedArg1 := tmp19077
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19079 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19080 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19079)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19079
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19081 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19078, tmp19080)
}
__typedArg0 := tmp19078
__typedArg1 := tmp19080
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19082 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19081, Nil)
}
__typedArg0 := tmp19081
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19083 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19082)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19082
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19084 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp19083)
}
__typedArg0 := symA
__typedArg1 := tmp19083
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19085 := Call(__e, PrimFunc(symdeclare), symelement_2, tmp19084)


_ = tmp19085

tmp19086 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19087 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19086)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19086
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19088 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp19087)
}
__typedArg0 := symA
__typedArg1 := tmp19087
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19089 := Call(__e, PrimFunc(symdeclare), symempty_2, tmp19088)


_ = tmp19089

tmp19090 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19091 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19090)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19090
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19092 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp19091)
}
__typedArg0 := symsymbol
__typedArg1 := tmp19091
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19093 := Call(__e, PrimFunc(symdeclare), symenable_1type_1theory, tmp19092)


_ = tmp19093

tmp19094 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19095 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19094)
}
__typedArg0 := symlist
__typedArg1 := tmp19094
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19096 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19095, Nil)
}
__typedArg0 := tmp19095
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19097 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19096)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19096
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19098 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp19097)
}
__typedArg0 := symsymbol
__typedArg1 := tmp19097
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19099 := Call(__e, PrimFunc(symdeclare), symexternal, tmp19098)


_ = tmp19099

tmp19100 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19101 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19100)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19100
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19102 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symexception, tmp19101)
}
__typedArg0 := symexception
__typedArg1 := tmp19101
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19103 := Call(__e, PrimFunc(symdeclare), symerror_1to_1string, tmp19102)


_ = tmp19103

tmp19104 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19105 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19104)
}
__typedArg0 := symlist
__typedArg1 := tmp19104
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19106 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19105, Nil)
}
__typedArg0 := tmp19105
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19107 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19106)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19106
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19108 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp19107)
}
__typedArg0 := symA
__typedArg1 := tmp19107
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19109 := Call(__e, PrimFunc(symdeclare), symexplode, tmp19108)


_ = tmp19109

tmp19110 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19111 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19110)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19110
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19112 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp19111)
}
__typedArg0 := symsymbol
__typedArg1 := tmp19111
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19113 := Call(__e, PrimFunc(symdeclare), symfactorise, tmp19112)


_ = tmp19113

tmp19114 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19115 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19114)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19114
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19116 := Call(__e, PrimFunc(symdeclare), symfactorise_2, tmp19115)


_ = tmp19116

tmp19117 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19118 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19117)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19117
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19119 := Call(__e, PrimFunc(symdeclare), symfail, tmp19118)


_ = tmp19119

tmp19120 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19121 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19120)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19120
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19122 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp19121)
}
__typedArg0 := symA
__typedArg1 := tmp19121
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19123 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19124 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19123)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19123
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19125 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp19124)
}
__typedArg0 := symA
__typedArg1 := tmp19124
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19126 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19125, Nil)
}
__typedArg0 := tmp19125
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19127 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19126)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19126
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19128 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19122, tmp19127)
}
__typedArg0 := tmp19122
__typedArg1 := tmp19127
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19129 := Call(__e, PrimFunc(symdeclare), symfix, tmp19128)


_ = tmp19129

tmp19130 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19131 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlazy, tmp19130)
}
__typedArg0 := symlazy
__typedArg1 := tmp19130
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19132 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19131, Nil)
}
__typedArg0 := tmp19131
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19133 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19132)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19132
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19134 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp19133)
}
__typedArg0 := symA
__typedArg1 := tmp19133
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19135 := Call(__e, PrimFunc(symdeclare), symfreeze, tmp19134)


_ = tmp19135

tmp19136 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symB, Nil)
}
__typedArg0 := symB
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19137 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_d, tmp19136)
}
__typedArg0 := sym_d
__typedArg1 := tmp19136
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19138 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp19137)
}
__typedArg0 := symA
__typedArg1 := tmp19137
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19139 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19140 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19139)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19139
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19141 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19138, tmp19140)
}
__typedArg0 := tmp19138
__typedArg1 := tmp19140
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19142 := Call(__e, PrimFunc(symdeclare), symfst, tmp19141)


_ = tmp19142

tmp19143 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19144 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19143)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19143
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19145 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp19144)
}
__typedArg0 := symsymbol
__typedArg1 := tmp19144
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19146 := Call(__e, PrimFunc(symdeclare), symgensym, tmp19145)


_ = tmp19146

tmp19147 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19148 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19147)
}
__typedArg0 := symlist
__typedArg1 := tmp19147
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19149 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19150 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19149)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19149
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19151 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp19150)
}
__typedArg0 := symA
__typedArg1 := tmp19150
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19152 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19151, Nil)
}
__typedArg0 := tmp19151
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19153 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19152)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19152
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19154 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19148, tmp19153)
}
__typedArg0 := tmp19148
__typedArg1 := tmp19153
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19155 := Call(__e, PrimFunc(symdeclare), symshen_4hds_a_2, tmp19154)


_ = tmp19155

tmp19156 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19157 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19156)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19156
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19158 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp19157)
}
__typedArg0 := symsymbol
__typedArg1 := tmp19157
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19159 := Call(__e, PrimFunc(symdeclare), symhush, tmp19158)


_ = tmp19159

tmp19160 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19161 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19160)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19160
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19162 := Call(__e, PrimFunc(symdeclare), symhush_2, tmp19161)


_ = tmp19162

tmp19163 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19164 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp19163)
}
__typedArg0 := symvector
__typedArg1 := tmp19163
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19165 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19166 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19165)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19165
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19167 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp19166)
}
__typedArg0 := symnumber
__typedArg1 := tmp19166
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19168 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19167, Nil)
}
__typedArg0 := tmp19167
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19169 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19168)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19168
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19170 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19164, tmp19169)
}
__typedArg0 := tmp19164
__typedArg1 := tmp19169
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19171 := Call(__e, PrimFunc(symdeclare), sym_5_1vector, tmp19170)


_ = tmp19171

tmp19172 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19173 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp19172)
}
__typedArg0 := symvector
__typedArg1 := tmp19172
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19174 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19175 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp19174)
}
__typedArg0 := symvector
__typedArg1 := tmp19174
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19176 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19175, Nil)
}
__typedArg0 := tmp19175
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19177 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19176)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19176
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19178 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp19177)
}
__typedArg0 := symA
__typedArg1 := tmp19177
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19179 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19178, Nil)
}
__typedArg0 := tmp19178
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19180 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19179)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19179
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19181 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp19180)
}
__typedArg0 := symnumber
__typedArg1 := tmp19180
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19182 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19181, Nil)
}
__typedArg0 := tmp19181
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19183 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19182)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19182
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19184 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19173, tmp19183)
}
__typedArg0 := tmp19173
__typedArg1 := tmp19183
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19185 := Call(__e, PrimFunc(symdeclare), symvector_1_6, tmp19184)


_ = tmp19185

tmp19186 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19187 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp19186)
}
__typedArg0 := symvector
__typedArg1 := tmp19186
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19188 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19187, Nil)
}
__typedArg0 := tmp19187
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19189 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19188)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19188
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19190 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp19189)
}
__typedArg0 := symnumber
__typedArg1 := tmp19189
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19191 := Call(__e, PrimFunc(symdeclare), symvector, tmp19190)


_ = tmp19191

tmp19192 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19193 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19192)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19192
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19194 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp19193)
}
__typedArg0 := symsymbol
__typedArg1 := tmp19193
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19195 := Call(__e, PrimFunc(symdeclare), symget_1time, tmp19194)


_ = tmp19195

tmp19196 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19197 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19196)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19196
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19198 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp19197)
}
__typedArg0 := symnumber
__typedArg1 := tmp19197
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19199 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19198, Nil)
}
__typedArg0 := tmp19198
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19200 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19199)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19199
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19201 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp19200)
}
__typedArg0 := symA
__typedArg1 := tmp19200
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19202 := Call(__e, PrimFunc(symdeclare), symhash, tmp19201)


_ = tmp19202

tmp19203 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19204 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19203)
}
__typedArg0 := symlist
__typedArg1 := tmp19203
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19205 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19206 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19205)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19205
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19207 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19204, tmp19206)
}
__typedArg0 := tmp19204
__typedArg1 := tmp19206
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19208 := Call(__e, PrimFunc(symdeclare), symhead, tmp19207)


_ = tmp19208

tmp19209 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19210 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp19209)
}
__typedArg0 := symvector
__typedArg1 := tmp19209
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19211 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19212 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19211)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19211
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19213 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19210, tmp19212)
}
__typedArg0 := tmp19210
__typedArg1 := tmp19212
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19214 := Call(__e, PrimFunc(symdeclare), symhdv, tmp19213)


_ = tmp19214

tmp19215 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19216 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19215)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19215
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19217 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp19216)
}
__typedArg0 := symstring
__typedArg1 := tmp19216
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19218 := Call(__e, PrimFunc(symdeclare), symhdstr, tmp19217)


_ = tmp19218

tmp19219 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19220 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19219)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19219
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19221 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp19220)
}
__typedArg0 := symA
__typedArg1 := tmp19220
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19222 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19221, Nil)
}
__typedArg0 := tmp19221
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19223 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19222)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19222
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19224 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp19223)
}
__typedArg0 := symA
__typedArg1 := tmp19223
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19225 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19224, Nil)
}
__typedArg0 := tmp19224
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19226 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19225)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19225
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19227 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, tmp19226)
}
__typedArg0 := symboolean
__typedArg1 := tmp19226
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19228 := Call(__e, PrimFunc(symdeclare), symif, tmp19227)


_ = tmp19228

tmp19229 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19230 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19229)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19229
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19231 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp19230)
}
__typedArg0 := symsymbol
__typedArg1 := tmp19230
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19232 := Call(__e, PrimFunc(symdeclare), symin_1package, tmp19231)


_ = tmp19232

tmp19233 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19234 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19233)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19233
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19235 := Call(__e, PrimFunc(symdeclare), symit, tmp19234)


_ = tmp19235

tmp19236 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19237 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19236)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19236
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19238 := Call(__e, PrimFunc(symdeclare), symimplementation, tmp19237)


_ = tmp19238

tmp19239 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19240 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19239)
}
__typedArg0 := symlist
__typedArg1 := tmp19239
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19241 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19242 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19241)
}
__typedArg0 := symlist
__typedArg1 := tmp19241
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19243 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19242, Nil)
}
__typedArg0 := tmp19242
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19244 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19243)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19243
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19245 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19240, tmp19244)
}
__typedArg0 := tmp19240
__typedArg1 := tmp19244
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19246 := Call(__e, PrimFunc(symdeclare), syminclude, tmp19245)


_ = tmp19246

tmp19247 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19248 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19247)
}
__typedArg0 := symlist
__typedArg1 := tmp19247
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19249 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19250 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19249)
}
__typedArg0 := symlist
__typedArg1 := tmp19249
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19251 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19250, Nil)
}
__typedArg0 := tmp19250
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19252 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19251)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19251
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19253 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19248, tmp19252)
}
__typedArg0 := tmp19248
__typedArg1 := tmp19252
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19254 := Call(__e, PrimFunc(symdeclare), syminclude_1all_1but, tmp19253)


_ = tmp19254

tmp19255 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19256 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19255)
}
__typedArg0 := symlist
__typedArg1 := tmp19255
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19257 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19256, Nil)
}
__typedArg0 := tmp19256
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19258 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19257)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19257
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19259 := Call(__e, PrimFunc(symdeclare), symincluded, tmp19258)


_ = tmp19259

tmp19260 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19261 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19260)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19260
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19262 := Call(__e, PrimFunc(symdeclare), syminferences, tmp19261)


_ = tmp19262

tmp19263 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19264 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19263)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19263
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19265 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp19264)
}
__typedArg0 := symstring
__typedArg1 := tmp19264
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19266 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19265, Nil)
}
__typedArg0 := tmp19265
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19267 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19266)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19266
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19268 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp19267)
}
__typedArg0 := symA
__typedArg1 := tmp19267
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19269 := Call(__e, PrimFunc(symdeclare), symshen_4insert, tmp19268)


_ = tmp19269

tmp19270 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19271 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19270)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19270
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19272 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp19271)
}
__typedArg0 := symA
__typedArg1 := tmp19271
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19273 := Call(__e, PrimFunc(symdeclare), syminteger_2, tmp19272)


_ = tmp19273

tmp19274 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19275 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19274)
}
__typedArg0 := symlist
__typedArg1 := tmp19274
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19276 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19275, Nil)
}
__typedArg0 := tmp19275
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19277 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19276)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19276
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19278 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp19277)
}
__typedArg0 := symsymbol
__typedArg1 := tmp19277
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19279 := Call(__e, PrimFunc(symdeclare), syminternal, tmp19278)


_ = tmp19279

tmp19280 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19281 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19280)
}
__typedArg0 := symlist
__typedArg1 := tmp19280
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19282 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19283 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19282)
}
__typedArg0 := symlist
__typedArg1 := tmp19282
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19284 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19285 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19284)
}
__typedArg0 := symlist
__typedArg1 := tmp19284
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19286 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19285, Nil)
}
__typedArg0 := tmp19285
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19287 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19286)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19286
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19288 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19283, tmp19287)
}
__typedArg0 := tmp19283
__typedArg1 := tmp19287
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19289 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19288, Nil)
}
__typedArg0 := tmp19288
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19290 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19289)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19289
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19291 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19281, tmp19290)
}
__typedArg0 := tmp19281
__typedArg1 := tmp19290
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19292 := Call(__e, PrimFunc(symdeclare), symintersection, tmp19291)


_ = tmp19292

tmp19293 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19294 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19293)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19293
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19295 := Call(__e, PrimFunc(symdeclare), symlanguage, tmp19294)


_ = tmp19295

tmp19296 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19297 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19296)
}
__typedArg0 := symlist
__typedArg1 := tmp19296
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19298 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19299 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19298)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19298
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19300 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19297, tmp19299)
}
__typedArg0 := tmp19297
__typedArg1 := tmp19299
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19301 := Call(__e, PrimFunc(symdeclare), symlength, tmp19300)


_ = tmp19301

tmp19302 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19303 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp19302)
}
__typedArg0 := symvector
__typedArg1 := tmp19302
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19304 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19305 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19304)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19304
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19306 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19303, tmp19305)
}
__typedArg0 := tmp19303
__typedArg1 := tmp19305
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19307 := Call(__e, PrimFunc(symdeclare), symlimit, tmp19306)


_ = tmp19307

tmp19308 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symin, Nil)
}
__typedArg0 := symin
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19309 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp19308)
}
__typedArg0 := symstream
__typedArg1 := tmp19308
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19310 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunit, Nil)
}
__typedArg0 := symunit
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19311 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19310)
}
__typedArg0 := symlist
__typedArg1 := tmp19310
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19312 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19311, Nil)
}
__typedArg0 := tmp19311
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19313 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19312)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19312
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19314 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19309, tmp19313)
}
__typedArg0 := tmp19309
__typedArg1 := tmp19313
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19315 := Call(__e, PrimFunc(symdeclare), symlineread, tmp19314)


_ = tmp19315

tmp19316 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19317 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19316)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19316
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19318 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp19317)
}
__typedArg0 := symstring
__typedArg1 := tmp19317
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19319 := Call(__e, PrimFunc(symdeclare), symload, tmp19318)


_ = tmp19319

tmp19320 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symB, Nil)
}
__typedArg0 := symB
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19321 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19320)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19320
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19322 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp19321)
}
__typedArg0 := symA
__typedArg1 := tmp19321
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19323 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19324 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19323)
}
__typedArg0 := symlist
__typedArg1 := tmp19323
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19325 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symB, Nil)
}
__typedArg0 := symB
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19326 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19325)
}
__typedArg0 := symlist
__typedArg1 := tmp19325
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19327 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19326, Nil)
}
__typedArg0 := tmp19326
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19328 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19327)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19327
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19329 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19324, tmp19328)
}
__typedArg0 := tmp19324
__typedArg1 := tmp19328
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19330 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19329, Nil)
}
__typedArg0 := tmp19329
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19331 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19330)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19330
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19332 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19322, tmp19331)
}
__typedArg0 := tmp19322
__typedArg1 := tmp19331
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19333 := Call(__e, PrimFunc(symdeclare), symmap, tmp19332)


_ = tmp19333

tmp19334 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symB, Nil)
}
__typedArg0 := symB
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19335 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19334)
}
__typedArg0 := symlist
__typedArg1 := tmp19334
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19336 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19335, Nil)
}
__typedArg0 := tmp19335
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19337 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19336)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19336
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19338 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp19337)
}
__typedArg0 := symA
__typedArg1 := tmp19337
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19339 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19340 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19339)
}
__typedArg0 := symlist
__typedArg1 := tmp19339
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19341 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symB, Nil)
}
__typedArg0 := symB
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19342 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19341)
}
__typedArg0 := symlist
__typedArg1 := tmp19341
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19343 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19342, Nil)
}
__typedArg0 := tmp19342
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19344 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19343)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19343
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19345 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19340, tmp19344)
}
__typedArg0 := tmp19340
__typedArg1 := tmp19344
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19346 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19345, Nil)
}
__typedArg0 := tmp19345
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19347 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19346)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19346
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19348 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19338, tmp19347)
}
__typedArg0 := tmp19338
__typedArg1 := tmp19347
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19349 := Call(__e, PrimFunc(symdeclare), symmapcan, tmp19348)


_ = tmp19349

tmp19350 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19351 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19350)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19350
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19352 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp19351)
}
__typedArg0 := symnumber
__typedArg1 := tmp19351
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19353 := Call(__e, PrimFunc(symdeclare), symmaxinferences, tmp19352)


_ = tmp19353

tmp19354 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19355 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19354)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19354
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19356 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp19355)
}
__typedArg0 := symnumber
__typedArg1 := tmp19355
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19357 := Call(__e, PrimFunc(symdeclare), symn_1_6string, tmp19356)


_ = tmp19357

tmp19358 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19359 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19358)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19358
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19360 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp19359)
}
__typedArg0 := symnumber
__typedArg1 := tmp19359
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19361 := Call(__e, PrimFunc(symdeclare), symnl, tmp19360)


_ = tmp19361

tmp19362 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19363 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19362)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19362
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19364 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, tmp19363)
}
__typedArg0 := symboolean
__typedArg1 := tmp19363
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19365 := Call(__e, PrimFunc(symdeclare), symnot, tmp19364)


_ = tmp19365

tmp19366 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19367 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19366)
}
__typedArg0 := symlist
__typedArg1 := tmp19366
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19368 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19369 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19368)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19368
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19370 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19367, tmp19369)
}
__typedArg0 := tmp19367
__typedArg1 := tmp19369
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19371 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19370, Nil)
}
__typedArg0 := tmp19370
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19372 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19371)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19371
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19373 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp19372)
}
__typedArg0 := symnumber
__typedArg1 := tmp19372
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19374 := Call(__e, PrimFunc(symdeclare), symnth, tmp19373)


_ = tmp19374

tmp19375 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19376 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19375)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19375
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19377 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp19376)
}
__typedArg0 := symA
__typedArg1 := tmp19376
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19378 := Call(__e, PrimFunc(symdeclare), symnumber_2, tmp19377)


_ = tmp19378

tmp19379 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19380 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19379)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19379
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19381 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symB, tmp19380)
}
__typedArg0 := symB
__typedArg1 := tmp19380
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19382 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19381, Nil)
}
__typedArg0 := tmp19381
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19383 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19382)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19382
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19384 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp19383)
}
__typedArg0 := symA
__typedArg1 := tmp19383
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19385 := Call(__e, PrimFunc(symdeclare), symoccurrences, tmp19384)


_ = tmp19385

tmp19386 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19387 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19386)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19386
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19388 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp19387)
}
__typedArg0 := symsymbol
__typedArg1 := tmp19387
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19389 := Call(__e, PrimFunc(symdeclare), symoccurs_1check, tmp19388)


_ = tmp19389

tmp19390 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19391 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19390)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19390
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19392 := Call(__e, PrimFunc(symdeclare), symoccurs_2, tmp19391)


_ = tmp19392

tmp19393 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19394 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19393)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19393
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19395 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp19394)
}
__typedArg0 := symsymbol
__typedArg1 := tmp19394
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19396 := Call(__e, PrimFunc(symdeclare), symoptimise, tmp19395)


_ = tmp19396

tmp19397 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19398 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19397)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19397
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19399 := Call(__e, PrimFunc(symdeclare), symoptimise_2, tmp19398)


_ = tmp19399

tmp19400 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19401 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19400)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19400
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19402 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, tmp19401)
}
__typedArg0 := symboolean
__typedArg1 := tmp19401
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19403 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19402, Nil)
}
__typedArg0 := tmp19402
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19404 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19403)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19403
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19405 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, tmp19404)
}
__typedArg0 := symboolean
__typedArg1 := tmp19404
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19406 := Call(__e, PrimFunc(symdeclare), symor, tmp19405)


_ = tmp19406

tmp19407 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19408 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19407)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19407
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19409 := Call(__e, PrimFunc(symdeclare), symos, tmp19408)


_ = tmp19409

tmp19410 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19411 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19410)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19410
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19412 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp19411)
}
__typedArg0 := symsymbol
__typedArg1 := tmp19411
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19413 := Call(__e, PrimFunc(symdeclare), sympackage_2, tmp19412)


_ = tmp19413

tmp19414 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19415 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19414)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19414
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19416 := Call(__e, PrimFunc(symdeclare), symport, tmp19415)


_ = tmp19416

tmp19417 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19418 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19417)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19417
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19419 := Call(__e, PrimFunc(symdeclare), symporters, tmp19418)


_ = tmp19419

tmp19420 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19421 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19420)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19420
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19422 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp19421)
}
__typedArg0 := symnumber
__typedArg1 := tmp19421
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19423 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19422, Nil)
}
__typedArg0 := tmp19422
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19424 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19423)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19423
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19425 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp19424)
}
__typedArg0 := symstring
__typedArg1 := tmp19424
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19426 := Call(__e, PrimFunc(symdeclare), sympos, tmp19425)


_ = tmp19426

tmp19427 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symout, Nil)
}
__typedArg0 := symout
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19428 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp19427)
}
__typedArg0 := symstream
__typedArg1 := tmp19427
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19429 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19430 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19429)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19429
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19431 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19428, tmp19430)
}
__typedArg0 := tmp19428
__typedArg1 := tmp19430
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19432 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19431, Nil)
}
__typedArg0 := tmp19431
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19433 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19432)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19432
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19434 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp19433)
}
__typedArg0 := symstring
__typedArg1 := tmp19433
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19435 := Call(__e, PrimFunc(symdeclare), sympr, tmp19434)


_ = tmp19435

tmp19436 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19437 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19436)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19436
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19438 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp19437)
}
__typedArg0 := symA
__typedArg1 := tmp19437
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19439 := Call(__e, PrimFunc(symdeclare), symprint, tmp19438)


_ = tmp19439

tmp19440 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19441 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19440)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19440
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19442 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp19441)
}
__typedArg0 := symsymbol
__typedArg1 := tmp19441
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19443 := Call(__e, PrimFunc(symdeclare), symprofile, tmp19442)


_ = tmp19443

tmp19444 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19445 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19444)
}
__typedArg0 := symlist
__typedArg1 := tmp19444
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19446 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19447 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19446)
}
__typedArg0 := symlist
__typedArg1 := tmp19446
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19448 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19447, Nil)
}
__typedArg0 := tmp19447
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19449 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19448)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19448
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19450 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19445, tmp19449)
}
__typedArg0 := tmp19445
__typedArg1 := tmp19449
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19451 := Call(__e, PrimFunc(symdeclare), sympreclude, tmp19450)


_ = tmp19451

tmp19452 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19453 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19452)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19452
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19454 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp19453)
}
__typedArg0 := symstring
__typedArg1 := tmp19453
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19455 := Call(__e, PrimFunc(symdeclare), symshen_4proc_1nl, tmp19454)


_ = tmp19455

tmp19456 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19457 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_d, tmp19456)
}
__typedArg0 := sym_d
__typedArg1 := tmp19456
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19458 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp19457)
}
__typedArg0 := symsymbol
__typedArg1 := tmp19457
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19459 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19458, Nil)
}
__typedArg0 := tmp19458
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19460 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19459)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19459
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19461 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp19460)
}
__typedArg0 := symsymbol
__typedArg1 := tmp19460
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19462 := Call(__e, PrimFunc(symdeclare), symprofile_1results, tmp19461)


_ = tmp19462

tmp19463 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19464 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19463)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19463
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19465 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp19464)
}
__typedArg0 := symA
__typedArg1 := tmp19464
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19466 := Call(__e, PrimFunc(symdeclare), symprotect, tmp19465)


_ = tmp19466

tmp19467 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19468 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19467)
}
__typedArg0 := symlist
__typedArg1 := tmp19467
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19469 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19470 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19469)
}
__typedArg0 := symlist
__typedArg1 := tmp19469
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19471 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19470, Nil)
}
__typedArg0 := tmp19470
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19472 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19471)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19471
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19473 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19468, tmp19472)
}
__typedArg0 := tmp19468
__typedArg1 := tmp19472
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19474 := Call(__e, PrimFunc(symdeclare), sympreclude_1all_1but, tmp19473)


_ = tmp19474

tmp19475 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symout, Nil)
}
__typedArg0 := symout
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19476 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp19475)
}
__typedArg0 := symstream
__typedArg1 := tmp19475
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19477 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19478 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19477)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19477
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19479 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19476, tmp19478)
}
__typedArg0 := tmp19476
__typedArg1 := tmp19478
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19480 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19479, Nil)
}
__typedArg0 := tmp19479
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19481 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19480)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19480
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19482 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp19481)
}
__typedArg0 := symstring
__typedArg1 := tmp19481
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19483 := Call(__e, PrimFunc(symdeclare), symshen_4prhush, tmp19482)


_ = tmp19483

tmp19484 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19485 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19484)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19484
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19486 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp19485)
}
__typedArg0 := symnumber
__typedArg1 := tmp19485
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19487 := Call(__e, PrimFunc(symdeclare), symprolog_1memory, tmp19486)


_ = tmp19487

tmp19488 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunit, Nil)
}
__typedArg0 := symunit
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19489 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19488)
}
__typedArg0 := symlist
__typedArg1 := tmp19488
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19490 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19489, Nil)
}
__typedArg0 := tmp19489
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19491 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19490)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19490
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19492 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp19491)
}
__typedArg0 := symsymbol
__typedArg1 := tmp19491
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19493 := Call(__e, PrimFunc(symdeclare), symps, tmp19492)


_ = tmp19493

tmp19494 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symin, Nil)
}
__typedArg0 := symin
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19495 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp19494)
}
__typedArg0 := symstream
__typedArg1 := tmp19494
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19496 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunit, Nil)
}
__typedArg0 := symunit
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19497 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19496)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19496
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19498 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19495, tmp19497)
}
__typedArg0 := tmp19495
__typedArg1 := tmp19497
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19499 := Call(__e, PrimFunc(symdeclare), symread, tmp19498)


_ = tmp19499

tmp19500 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symin, Nil)
}
__typedArg0 := symin
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19501 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp19500)
}
__typedArg0 := symstream
__typedArg1 := tmp19500
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19502 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19503 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19502)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19502
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19504 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19501, tmp19503)
}
__typedArg0 := tmp19501
__typedArg1 := tmp19503
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19505 := Call(__e, PrimFunc(symdeclare), symread_1byte, tmp19504)


_ = tmp19505

tmp19506 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19507 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19506)
}
__typedArg0 := symlist
__typedArg1 := tmp19506
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19508 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19507, Nil)
}
__typedArg0 := tmp19507
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19509 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19508)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19508
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19510 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp19509)
}
__typedArg0 := symstring
__typedArg1 := tmp19509
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19511 := Call(__e, PrimFunc(symdeclare), symread_1file_1as_1bytelist, tmp19510)


_ = tmp19511

tmp19512 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19513 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19512)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19512
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19514 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp19513)
}
__typedArg0 := symstring
__typedArg1 := tmp19513
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19515 := Call(__e, PrimFunc(symdeclare), symread_1file_1as_1string, tmp19514)


_ = tmp19515

tmp19516 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunit, Nil)
}
__typedArg0 := symunit
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19517 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19516)
}
__typedArg0 := symlist
__typedArg1 := tmp19516
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19518 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19517, Nil)
}
__typedArg0 := tmp19517
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19519 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19518)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19518
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19520 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp19519)
}
__typedArg0 := symstring
__typedArg1 := tmp19519
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19521 := Call(__e, PrimFunc(symdeclare), symread_1file, tmp19520)


_ = tmp19521

tmp19522 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunit, Nil)
}
__typedArg0 := symunit
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19523 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19522)
}
__typedArg0 := symlist
__typedArg1 := tmp19522
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19524 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19523, Nil)
}
__typedArg0 := tmp19523
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19525 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19524)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19524
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19526 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp19525)
}
__typedArg0 := symstring
__typedArg1 := tmp19525
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19527 := Call(__e, PrimFunc(symdeclare), symread_1from_1string, tmp19526)


_ = tmp19527

tmp19528 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunit, Nil)
}
__typedArg0 := symunit
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19529 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19528)
}
__typedArg0 := symlist
__typedArg1 := tmp19528
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19530 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19529, Nil)
}
__typedArg0 := tmp19529
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19531 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19530)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19530
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19532 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp19531)
}
__typedArg0 := symstring
__typedArg1 := tmp19531
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19533 := Call(__e, PrimFunc(symdeclare), symread_1from_1string_1unprocessed, tmp19532)


_ = tmp19533

tmp19534 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19535 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19534)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19534
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19536 := Call(__e, PrimFunc(symdeclare), symrelease, tmp19535)


_ = tmp19536

tmp19537 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19538 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19537)
}
__typedArg0 := symlist
__typedArg1 := tmp19537
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19539 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19540 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19539)
}
__typedArg0 := symlist
__typedArg1 := tmp19539
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19541 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19540, Nil)
}
__typedArg0 := tmp19540
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19542 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19541)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19541
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19543 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19538, tmp19542)
}
__typedArg0 := tmp19538
__typedArg1 := tmp19542
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19544 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19543, Nil)
}
__typedArg0 := tmp19543
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19545 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19544)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19544
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19546 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp19545)
}
__typedArg0 := symA
__typedArg1 := tmp19545
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19547 := Call(__e, PrimFunc(symdeclare), symremove, tmp19546)


_ = tmp19547

tmp19548 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19549 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19548)
}
__typedArg0 := symlist
__typedArg1 := tmp19548
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19550 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19551 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19550)
}
__typedArg0 := symlist
__typedArg1 := tmp19550
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19552 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19551, Nil)
}
__typedArg0 := tmp19551
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19553 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19552)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19552
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19554 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19549, tmp19553)
}
__typedArg0 := tmp19549
__typedArg1 := tmp19553
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19555 := Call(__e, PrimFunc(symdeclare), symreverse, tmp19554)


_ = tmp19555

tmp19556 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19557 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19556)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19556
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19558 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp19557)
}
__typedArg0 := symstring
__typedArg1 := tmp19557
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19559 := Call(__e, PrimFunc(symdeclare), symsimple_1error, tmp19558)


_ = tmp19559

tmp19560 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symB, Nil)
}
__typedArg0 := symB
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19561 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_d, tmp19560)
}
__typedArg0 := sym_d
__typedArg1 := tmp19560
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19562 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp19561)
}
__typedArg0 := symA
__typedArg1 := tmp19561
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19563 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symB, Nil)
}
__typedArg0 := symB
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19564 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19563)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19563
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19565 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19562, tmp19564)
}
__typedArg0 := tmp19562
__typedArg1 := tmp19564
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19566 := Call(__e, PrimFunc(symdeclare), symsnd, tmp19565)


_ = tmp19566

tmp19567 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19568 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19567)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19567
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19569 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp19568)
}
__typedArg0 := symnumber
__typedArg1 := tmp19568
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19570 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19569, Nil)
}
__typedArg0 := tmp19569
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19571 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19570)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19570
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19572 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp19571)
}
__typedArg0 := symsymbol
__typedArg1 := tmp19571
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19573 := Call(__e, PrimFunc(symdeclare), symspecialise, tmp19572)


_ = tmp19573

tmp19574 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19575 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19574)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19574
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19576 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp19575)
}
__typedArg0 := symsymbol
__typedArg1 := tmp19575
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19577 := Call(__e, PrimFunc(symdeclare), symspy, tmp19576)


_ = tmp19577

tmp19578 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19579 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19578)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19578
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19580 := Call(__e, PrimFunc(symdeclare), symspy_2, tmp19579)


_ = tmp19580

tmp19581 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19582 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19581)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19581
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19583 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp19582)
}
__typedArg0 := symsymbol
__typedArg1 := tmp19582
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19584 := Call(__e, PrimFunc(symdeclare), symstep, tmp19583)


_ = tmp19584

tmp19585 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19586 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19585)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19585
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19587 := Call(__e, PrimFunc(symdeclare), symstep_2, tmp19586)


_ = tmp19587

tmp19588 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symin, Nil)
}
__typedArg0 := symin
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19589 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp19588)
}
__typedArg0 := symstream
__typedArg1 := tmp19588
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19590 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19589, Nil)
}
__typedArg0 := tmp19589
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19591 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19590)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19590
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19592 := Call(__e, PrimFunc(symdeclare), symstinput, tmp19591)


_ = tmp19592

tmp19593 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symout, Nil)
}
__typedArg0 := symout
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19594 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp19593)
}
__typedArg0 := symstream
__typedArg1 := tmp19593
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19595 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19594, Nil)
}
__typedArg0 := tmp19594
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19596 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19595)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19595
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19597 := Call(__e, PrimFunc(symdeclare), symstoutput, tmp19596)


_ = tmp19597

tmp19598 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19599 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19598)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19598
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19600 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp19599)
}
__typedArg0 := symA
__typedArg1 := tmp19599
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19601 := Call(__e, PrimFunc(symdeclare), symstring_2, tmp19600)


_ = tmp19601

tmp19602 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19603 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19602)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19602
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19604 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp19603)
}
__typedArg0 := symA
__typedArg1 := tmp19603
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19605 := Call(__e, PrimFunc(symdeclare), symstr, tmp19604)


_ = tmp19605

tmp19606 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19607 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19606)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19606
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19608 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp19607)
}
__typedArg0 := symstring
__typedArg1 := tmp19607
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19609 := Call(__e, PrimFunc(symdeclare), symstring_1_6n, tmp19608)


_ = tmp19609

tmp19610 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19611 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19610)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19610
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19612 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp19611)
}
__typedArg0 := symstring
__typedArg1 := tmp19611
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19613 := Call(__e, PrimFunc(symdeclare), symstring_1_6symbol, tmp19612)


_ = tmp19613

tmp19614 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19615 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19614)
}
__typedArg0 := symlist
__typedArg1 := tmp19614
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19616 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19617 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19616)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19616
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19618 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19615, tmp19617)
}
__typedArg0 := tmp19615
__typedArg1 := tmp19617
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19619 := Call(__e, PrimFunc(symdeclare), symsum, tmp19618)


_ = tmp19619

tmp19620 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19621 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19620)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19620
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19622 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp19621)
}
__typedArg0 := symA
__typedArg1 := tmp19621
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19623 := Call(__e, PrimFunc(symdeclare), symsymbol_2, tmp19622)


_ = tmp19623

tmp19624 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19625 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19624)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19624
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19626 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp19625)
}
__typedArg0 := symsymbol
__typedArg1 := tmp19625
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19627 := Call(__e, PrimFunc(symdeclare), symsystemf, tmp19626)


_ = tmp19627

tmp19628 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19629 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19628)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19628
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19630 := Call(__e, PrimFunc(symdeclare), symsystem_1S_2, tmp19629)


_ = tmp19630

tmp19631 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19632 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19631)
}
__typedArg0 := symlist
__typedArg1 := tmp19631
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19633 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19634 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19633)
}
__typedArg0 := symlist
__typedArg1 := tmp19633
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19635 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19634, Nil)
}
__typedArg0 := tmp19634
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19636 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19635)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19635
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19637 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19632, tmp19636)
}
__typedArg0 := tmp19632
__typedArg1 := tmp19636
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19638 := Call(__e, PrimFunc(symdeclare), symtail, tmp19637)


_ = tmp19638

tmp19639 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19640 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19639)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19639
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19641 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp19640)
}
__typedArg0 := symstring
__typedArg1 := tmp19640
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19642 := Call(__e, PrimFunc(symdeclare), symtlstr, tmp19641)


_ = tmp19642

tmp19643 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19644 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp19643)
}
__typedArg0 := symvector
__typedArg1 := tmp19643
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19645 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19646 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp19645)
}
__typedArg0 := symvector
__typedArg1 := tmp19645
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19647 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19646, Nil)
}
__typedArg0 := tmp19646
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19648 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19647)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19647
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19649 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19644, tmp19648)
}
__typedArg0 := tmp19644
__typedArg1 := tmp19648
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19650 := Call(__e, PrimFunc(symdeclare), symtlv, tmp19649)


_ = tmp19650

tmp19651 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19652 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19651)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19651
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19653 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp19652)
}
__typedArg0 := symsymbol
__typedArg1 := tmp19652
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19654 := Call(__e, PrimFunc(symdeclare), symtc, tmp19653)


_ = tmp19654

tmp19655 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19656 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19655)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19655
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19657 := Call(__e, PrimFunc(symdeclare), symtc_2, tmp19656)


_ = tmp19657

tmp19658 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19659 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlazy, tmp19658)
}
__typedArg0 := symlazy
__typedArg1 := tmp19658
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19660 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19661 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19660)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19660
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19662 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19659, tmp19661)
}
__typedArg0 := tmp19659
__typedArg1 := tmp19661
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19663 := Call(__e, PrimFunc(symdeclare), symthaw, tmp19662)


_ = tmp19663

tmp19664 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19665 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19664)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19664
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19666 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp19665)
}
__typedArg0 := symsymbol
__typedArg1 := tmp19665
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19667 := Call(__e, PrimFunc(symdeclare), symtrack, tmp19666)


_ = tmp19667

tmp19668 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19669 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19668)
}
__typedArg0 := symlist
__typedArg1 := tmp19668
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19670 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19669, Nil)
}
__typedArg0 := tmp19669
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19671 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19670)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19670
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19672 := Call(__e, PrimFunc(symdeclare), symtracked, tmp19671)


_ = tmp19672

tmp19673 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19674 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19673)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19673
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19675 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symexception, tmp19674)
}
__typedArg0 := symexception
__typedArg1 := tmp19674
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19676 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19677 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19676)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19676
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19678 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19675, tmp19677)
}
__typedArg0 := tmp19675
__typedArg1 := tmp19677
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19679 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19678, Nil)
}
__typedArg0 := tmp19678
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19680 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19679)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19679
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19681 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp19680)
}
__typedArg0 := symA
__typedArg1 := tmp19680
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19682 := Call(__e, PrimFunc(symdeclare), symtrap_1error, tmp19681)


_ = tmp19682

tmp19683 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19684 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19683)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19683
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19685 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp19684)
}
__typedArg0 := symA
__typedArg1 := tmp19684
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19686 := Call(__e, PrimFunc(symdeclare), symtuple_2, tmp19685)


_ = tmp19686

tmp19687 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19688 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19687)
}
__typedArg0 := symlist
__typedArg1 := tmp19687
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19689 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19688, Nil)
}
__typedArg0 := tmp19688
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19690 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19689)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19689
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19691 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp19690)
}
__typedArg0 := symstring
__typedArg1 := tmp19690
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19692 := Call(__e, PrimFunc(symdeclare), symunabsolute, tmp19691)


_ = tmp19692

tmp19693 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19694 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19693)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19693
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19695 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp19694)
}
__typedArg0 := symsymbol
__typedArg1 := tmp19694
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19696 := Call(__e, PrimFunc(symdeclare), symundefmacro, tmp19695)


_ = tmp19696

tmp19697 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19698 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19697)
}
__typedArg0 := symlist
__typedArg1 := tmp19697
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19699 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19700 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19699)
}
__typedArg0 := symlist
__typedArg1 := tmp19699
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19701 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19702 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19701)
}
__typedArg0 := symlist
__typedArg1 := tmp19701
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19703 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19702, Nil)
}
__typedArg0 := tmp19702
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19704 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19703)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19703
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19705 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19700, tmp19704)
}
__typedArg0 := tmp19700
__typedArg1 := tmp19704
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19706 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19705, Nil)
}
__typedArg0 := tmp19705
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19707 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19706)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19706
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19708 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19698, tmp19707)
}
__typedArg0 := tmp19698
__typedArg1 := tmp19707
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19709 := Call(__e, PrimFunc(symdeclare), symunion, tmp19708)


_ = tmp19709

tmp19710 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19711 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19710)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19710
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19712 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp19711)
}
__typedArg0 := symsymbol
__typedArg1 := tmp19711
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19713 := Call(__e, PrimFunc(symdeclare), symunprofile, tmp19712)


_ = tmp19713

tmp19714 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19715 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19714)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19714
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19716 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp19715)
}
__typedArg0 := symsymbol
__typedArg1 := tmp19715
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19717 := Call(__e, PrimFunc(symdeclare), symuntrack, tmp19716)


_ = tmp19717

tmp19718 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19719 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19718)
}
__typedArg0 := symlist
__typedArg1 := tmp19718
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19720 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19719, Nil)
}
__typedArg0 := tmp19719
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19721 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19720)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19720
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19722 := Call(__e, PrimFunc(symdeclare), symuserdefs, tmp19721)


_ = tmp19722

tmp19723 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19724 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19723)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19723
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19725 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp19724)
}
__typedArg0 := symA
__typedArg1 := tmp19724
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19726 := Call(__e, PrimFunc(symdeclare), symvariable_2, tmp19725)


_ = tmp19726

tmp19727 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19728 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19727)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19727
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19729 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp19728)
}
__typedArg0 := symA
__typedArg1 := tmp19728
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19730 := Call(__e, PrimFunc(symdeclare), symvector_2, tmp19729)


_ = tmp19730

tmp19731 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19732 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19731)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19731
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19733 := Call(__e, PrimFunc(symdeclare), symversion, tmp19732)


_ = tmp19733

tmp19734 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19735 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19734)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19734
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19736 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp19735)
}
__typedArg0 := symA
__typedArg1 := tmp19735
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19737 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19736, Nil)
}
__typedArg0 := tmp19736
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19738 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19737)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19737
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19739 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp19738)
}
__typedArg0 := symstring
__typedArg1 := tmp19738
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19740 := Call(__e, PrimFunc(symdeclare), symwrite_1to_1file, tmp19739)


_ = tmp19740

tmp19741 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symout, Nil)
}
__typedArg0 := symout
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19742 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp19741)
}
__typedArg0 := symstream
__typedArg1 := tmp19741
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19743 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19744 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19743)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19743
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19745 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19742, tmp19744)
}
__typedArg0 := tmp19742
__typedArg1 := tmp19744
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19746 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19745, Nil)
}
__typedArg0 := tmp19745
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19747 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19746)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19746
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19748 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp19747)
}
__typedArg0 := symnumber
__typedArg1 := tmp19747
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19749 := Call(__e, PrimFunc(symdeclare), symwrite_1byte, tmp19748)


_ = tmp19749

tmp19750 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19751 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19750)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19750
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19752 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp19751)
}
__typedArg0 := symstring
__typedArg1 := tmp19751
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19753 := Call(__e, PrimFunc(symdeclare), symy_1or_1n_2, tmp19752)


_ = tmp19753

tmp19754 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19755 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19754)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19754
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19756 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp19755)
}
__typedArg0 := symnumber
__typedArg1 := tmp19755
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19757 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19756, Nil)
}
__typedArg0 := tmp19756
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19758 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19757)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19757
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19759 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp19758)
}
__typedArg0 := symnumber
__typedArg1 := tmp19758
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19760 := Call(__e, PrimFunc(symdeclare), sym_6, tmp19759)


_ = tmp19760

tmp19761 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19762 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19761)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19761
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19763 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp19762)
}
__typedArg0 := symnumber
__typedArg1 := tmp19762
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19764 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19763, Nil)
}
__typedArg0 := tmp19763
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19765 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19764)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19764
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19766 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp19765)
}
__typedArg0 := symnumber
__typedArg1 := tmp19765
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19767 := Call(__e, PrimFunc(symdeclare), sym_5, tmp19766)


_ = tmp19767

tmp19768 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19769 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19768)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19768
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19770 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp19769)
}
__typedArg0 := symnumber
__typedArg1 := tmp19769
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19771 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19770, Nil)
}
__typedArg0 := tmp19770
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19772 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19771)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19771
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19773 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp19772)
}
__typedArg0 := symnumber
__typedArg1 := tmp19772
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19774 := Call(__e, PrimFunc(symdeclare), sym_6_a, tmp19773)


_ = tmp19774

tmp19775 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19776 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19775)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19775
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19777 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp19776)
}
__typedArg0 := symnumber
__typedArg1 := tmp19776
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19778 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19777, Nil)
}
__typedArg0 := tmp19777
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19779 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19778)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19778
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19780 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp19779)
}
__typedArg0 := symnumber
__typedArg1 := tmp19779
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19781 := Call(__e, PrimFunc(symdeclare), sym_5_a, tmp19780)


_ = tmp19781

tmp19782 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19783 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19782)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19782
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19784 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp19783)
}
__typedArg0 := symA
__typedArg1 := tmp19783
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19785 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19784, Nil)
}
__typedArg0 := tmp19784
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19786 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19785)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19785
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19787 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp19786)
}
__typedArg0 := symA
__typedArg1 := tmp19786
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19788 := Call(__e, PrimFunc(symdeclare), sym_a, tmp19787)


_ = tmp19788

tmp19789 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19790 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19789)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19789
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19791 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp19790)
}
__typedArg0 := symnumber
__typedArg1 := tmp19790
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19792 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19791, Nil)
}
__typedArg0 := tmp19791
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19793 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19792)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19792
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19794 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp19793)
}
__typedArg0 := symnumber
__typedArg1 := tmp19793
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19795 := Call(__e, PrimFunc(symdeclare), sym_7, tmp19794)


_ = tmp19795

tmp19796 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19797 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19796)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19796
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19798 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp19797)
}
__typedArg0 := symnumber
__typedArg1 := tmp19797
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19799 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19798, Nil)
}
__typedArg0 := tmp19798
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19800 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19799)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19799
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19801 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp19800)
}
__typedArg0 := symnumber
__typedArg1 := tmp19800
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19802 := Call(__e, PrimFunc(symdeclare), sym_c, tmp19801)


_ = tmp19802

tmp19803 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19804 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19803)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19803
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19805 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp19804)
}
__typedArg0 := symnumber
__typedArg1 := tmp19804
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19806 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19805, Nil)
}
__typedArg0 := tmp19805
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19807 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19806)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19806
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19808 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp19807)
}
__typedArg0 := symnumber
__typedArg1 := tmp19807
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19809 := Call(__e, PrimFunc(symdeclare), sym_1, tmp19808)


_ = tmp19809

tmp19810 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19811 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19810)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19810
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19812 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp19811)
}
__typedArg0 := symnumber
__typedArg1 := tmp19811
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19813 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19812, Nil)
}
__typedArg0 := tmp19812
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19814 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19813)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19813
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19815 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp19814)
}
__typedArg0 := symnumber
__typedArg1 := tmp19814
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19816 := Call(__e, PrimFunc(symdeclare), sym_d, tmp19815)


_ = tmp19816

tmp19817 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19818 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19817)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19817
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19819 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symB, tmp19818)
}
__typedArg0 := symB
__typedArg1 := tmp19818
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19820 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19819, Nil)
}
__typedArg0 := tmp19819
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19821 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19820)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19820
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19822 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp19821)
}
__typedArg0 := symA
__typedArg1 := tmp19821
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symdeclare), sym_a_a, tmp19822)
return




}, 0)

