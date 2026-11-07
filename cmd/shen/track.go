package main

import . "github.com/pyrex41/shen-go/kl"

var TrackMain = MakeNative(func(__e *ControlFlow) {
tmp13555 := MakeNative(func(__e *ControlFlow) {
V5886 := __e.Get(1)
_ = V5886
tmp13556 := Call(__e, PrimFunc(symshen_4app), V5886, MakeString(";\n"), symshen_4a)


tmp13557 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("partial function "))
__typedS1, __typedOK1 := TypedString(tmp13556)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("partial function ")
__typedArg1 := tmp13556
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp13558 := Call(__e, PrimFunc(symstoutput))


tmp13559 := Call(__e, PrimFunc(sympr), tmp13557, tmp13558)


_ = tmp13559

tmp13568 := Call(__e, PrimFunc(symshen_4tracked_2), V5886)


tmp13569 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp13568)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp13568
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres13563 Obj

if True == tmp13569 {
tmp13565 := Call(__e, PrimFunc(symshen_4app), V5886, MakeString("? "), symshen_4a)


tmp13567 := Call(__e, PrimFunc(symy_1or_1n_2), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("track "))
__typedS1, __typedOK1 := TypedString(tmp13565)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("track ")
__typedArg1 := tmp13565
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())


var ifres13564 Obj

if True == tmp13567 {
ifres13564 = True


} else {
ifres13564 = False


}

ifres13563 = ifres13564


} else {
ifres13563 = False


}

var ifres13560 Obj

if True == ifres13563 {
tmp13561 := Call(__e, PrimFunc(symps), V5886)


tmp13562 := Call(__e, PrimFunc(symshen_4track_1function), tmp13561)


ifres13560 = tmp13562


} else {
ifres13560 = symshen_4ok


}

_ = ifres13560

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("aborted"))
}
__typedArg0 := MakeString("aborted")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}, 1)

tmp13570 := Call(__e, ns2_1set, symshen_4f_1error, tmp13555)


_ = tmp13570

tmp13571 := MakeNative(func(__e *ControlFlow) {
V5887 := __e.Get(1)
_ = V5887
tmp13572 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dtracking_d)
}
__typedArg0 := symshen_4_dtracking_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symelement_2), V5887, tmp13572)
return


}, 1)

tmp13573 := Call(__e, ns2_1set, symshen_4tracked_2, tmp13571)


_ = tmp13573

tmp13574 := MakeNative(func(__e *ControlFlow) {
V5888 := __e.Get(1)
_ = V5888
tmp13575 := Call(__e, PrimFunc(symps), V5888)


let__13544 := tmp13575
_ = let__13544

__e.TailApply(PrimFunc(symshen_4track_1function), let__13544)
return


}, 1)

tmp13576 := Call(__e, ns2_1set, symtrack, tmp13574)


_ = tmp13576

tmp13577 := MakeNative(func(__e *ControlFlow) {
V5892 := __e.Get(1)
_ = V5892
tmp13631 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5892)
}
__typedArg0 := V5892
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13605 Obj

if True == tmp13631 {
tmp13629 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5892)
}
__typedArg0 := V5892
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13630 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefun, tmp13629)
}
__typedArg0 := symdefun
__typedArg1 := tmp13629
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13607 Obj

if True == tmp13630 {
tmp13627 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5892)
}
__typedArg0 := V5892
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13628 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13627)
}
__typedArg0 := tmp13627
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13609 Obj

if True == tmp13628 {
tmp13624 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5892)
}
__typedArg0 := V5892
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13625 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13624)
}
__typedArg0 := tmp13624
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13626 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13625)
}
__typedArg0 := tmp13625
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13611 Obj

if True == tmp13626 {
tmp13620 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5892)
}
__typedArg0 := V5892
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13621 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13620)
}
__typedArg0 := tmp13620
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13622 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13621)
}
__typedArg0 := tmp13621
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13623 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13622)
}
__typedArg0 := tmp13622
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13613 Obj

if True == tmp13623 {
tmp13615 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5892)
}
__typedArg0 := V5892
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13616 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13615)
}
__typedArg0 := tmp13615
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13617 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13616)
}
__typedArg0 := tmp13616
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13618 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13617)
}
__typedArg0 := tmp13617
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13619 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp13618)
}
__typedArg0 := Nil
__typedArg1 := tmp13618
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13614 Obj

if True == tmp13619 {
ifres13614 = True


} else {
ifres13614 = False


}

ifres13613 = ifres13614


} else {
ifres13613 = False


}

var ifres13612 Obj

if True == ifres13613 {
ifres13612 = True


} else {
ifres13612 = False


}

ifres13611 = ifres13612


} else {
ifres13611 = False


}

var ifres13610 Obj

if True == ifres13611 {
ifres13610 = True


} else {
ifres13610 = False


}

ifres13609 = ifres13610


} else {
ifres13609 = False


}

var ifres13608 Obj

if True == ifres13609 {
ifres13608 = True


} else {
ifres13608 = False


}

ifres13607 = ifres13608


} else {
ifres13607 = False


}

var ifres13606 Obj

if True == ifres13607 {
ifres13606 = True


} else {
ifres13606 = False


}

ifres13605 = ifres13606


} else {
ifres13605 = False


}

if True == ifres13605 {
tmp13578 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5892)
}
__typedArg0 := V5892
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13579 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13578)
}
__typedArg0 := tmp13578
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13580 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5892)
}
__typedArg0 := V5892
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13581 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13580)
}
__typedArg0 := tmp13580
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13582 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13581)
}
__typedArg0 := tmp13581
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13583 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5892)
}
__typedArg0 := V5892
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13584 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13583)
}
__typedArg0 := tmp13583
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13585 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5892)
}
__typedArg0 := V5892
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13586 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13585)
}
__typedArg0 := tmp13585
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13587 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13586)
}
__typedArg0 := tmp13586
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13588 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5892)
}
__typedArg0 := V5892
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13589 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13588)
}
__typedArg0 := tmp13588
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13590 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13589)
}
__typedArg0 := tmp13589
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13591 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13590)
}
__typedArg0 := tmp13590
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13592 := Call(__e, PrimFunc(symshen_4insert_1tracking_1code), tmp13584, tmp13587, tmp13591)


tmp13593 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13592, Nil)
}
__typedArg0 := tmp13592
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13594 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13582, tmp13593)
}
__typedArg0 := tmp13582
__typedArg1 := tmp13593
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13595 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13579, tmp13594)
}
__typedArg0 := tmp13579
__typedArg1 := tmp13594
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13596 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefun, tmp13595)
}
__typedArg0 := symdefun
__typedArg1 := tmp13595
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__13545 := tmp13596
_ = let__13545

tmp13597 := Call(__e, PrimFunc(symeval_1kl), let__13545)


let__13546 := tmp13597
_ = let__13546

tmp13598 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5892)
}
__typedArg0 := V5892
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13599 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13598)
}
__typedArg0 := tmp13598
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13600 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dtracking_d)
}
__typedArg0 := symshen_4_dtracking_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp13601 := Call(__e, PrimFunc(symadjoin), tmp13599, tmp13600)


tmp13602 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dtracking_d, tmp13601)
}
__typedArg0 := symshen_4_dtracking_d
__typedArg1 := tmp13601
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

let__13547 := tmp13602
_ = let__13547

tmp13603 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5892)
}
__typedArg0 := V5892
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13603)
}
__typedArg0 := tmp13603
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.track-function"))
}
__typedArg0 := MakeString("implementation error in shen.track-function")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}, 1)

tmp13632 := Call(__e, ns2_1set, symshen_4track_1function, tmp13577)


_ = tmp13632

tmp13633 := MakeNative(func(__e *ControlFlow) {
V5896 := __e.Get(1)
_ = V5896
V5897 := __e.Get(2)
_ = V5897
V5898 := __e.Get(3)
_ = V5898
tmp13634 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4_dcall_d, Nil)
}
__typedArg0 := symshen_4_dcall_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13635 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvalue, tmp13634)
}
__typedArg0 := symvalue
__typedArg1 := tmp13634
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13636 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), Nil)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13637 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13635, tmp13636)
}
__typedArg0 := tmp13635
__typedArg1 := tmp13636
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13638 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_7, tmp13637)
}
__typedArg0 := sym_7
__typedArg1 := tmp13637
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13639 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13638, Nil)
}
__typedArg0 := tmp13638
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13640 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4_dcall_d, tmp13639)
}
__typedArg0 := symshen_4_dcall_d
__typedArg1 := tmp13639
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13641 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symset, tmp13640)
}
__typedArg0 := symset
__typedArg1 := tmp13640
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13642 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4_dcall_d, Nil)
}
__typedArg0 := symshen_4_dcall_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13643 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvalue, tmp13642)
}
__typedArg0 := symvalue
__typedArg1 := tmp13642
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13644 := Call(__e, PrimFunc(symshen_4prolog_1track), V5898, V5897)


tmp13645 := Call(__e, PrimFunc(symshen_4cons_1form), tmp13644)


tmp13646 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13645, Nil)
}
__typedArg0 := tmp13645
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13647 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5896, tmp13646)
}
__typedArg0 := V5896
__typedArg1 := tmp13646
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13648 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13643, tmp13647)
}
__typedArg0 := tmp13643
__typedArg1 := tmp13647
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13649 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4input_1track, tmp13648)
}
__typedArg0 := symshen_4input_1track
__typedArg1 := tmp13648
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13650 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4terpri_1or_1read_1char, Nil)
}
__typedArg0 := symshen_4terpri_1or_1read_1char
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13651 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4_dcall_d, Nil)
}
__typedArg0 := symshen_4_dcall_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13652 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvalue, tmp13651)
}
__typedArg0 := symvalue
__typedArg1 := tmp13651
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13653 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symResult, Nil)
}
__typedArg0 := symResult
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13654 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5896, tmp13653)
}
__typedArg0 := V5896
__typedArg1 := tmp13653
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13655 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13652, tmp13654)
}
__typedArg0 := tmp13652
__typedArg1 := tmp13654
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13656 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4output_1track, tmp13655)
}
__typedArg0 := symshen_4output_1track
__typedArg1 := tmp13655
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13657 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4_dcall_d, Nil)
}
__typedArg0 := symshen_4_dcall_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13658 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvalue, tmp13657)
}
__typedArg0 := symvalue
__typedArg1 := tmp13657
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13659 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), Nil)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13660 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13658, tmp13659)
}
__typedArg0 := tmp13658
__typedArg1 := tmp13659
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13661 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1, tmp13660)
}
__typedArg0 := sym_1
__typedArg1 := tmp13660
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13662 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13661, Nil)
}
__typedArg0 := tmp13661
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13663 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4_dcall_d, tmp13662)
}
__typedArg0 := symshen_4_dcall_d
__typedArg1 := tmp13662
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13664 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symset, tmp13663)
}
__typedArg0 := symset
__typedArg1 := tmp13663
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13665 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4terpri_1or_1read_1char, Nil)
}
__typedArg0 := symshen_4terpri_1or_1read_1char
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13666 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symResult, Nil)
}
__typedArg0 := symResult
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13667 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13665, tmp13666)
}
__typedArg0 := tmp13665
__typedArg1 := tmp13666
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13668 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdo, tmp13667)
}
__typedArg0 := symdo
__typedArg1 := tmp13667
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13669 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13668, Nil)
}
__typedArg0 := tmp13668
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13670 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13664, tmp13669)
}
__typedArg0 := tmp13664
__typedArg1 := tmp13669
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13671 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdo, tmp13670)
}
__typedArg0 := symdo
__typedArg1 := tmp13670
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13672 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13671, Nil)
}
__typedArg0 := tmp13671
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13673 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13656, tmp13672)
}
__typedArg0 := tmp13656
__typedArg1 := tmp13672
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13674 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdo, tmp13673)
}
__typedArg0 := symdo
__typedArg1 := tmp13673
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13675 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13674, Nil)
}
__typedArg0 := tmp13674
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13676 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5898, tmp13675)
}
__typedArg0 := V5898
__typedArg1 := tmp13675
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13677 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symResult, tmp13676)
}
__typedArg0 := symResult
__typedArg1 := tmp13676
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13678 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp13677)
}
__typedArg0 := symlet
__typedArg1 := tmp13677
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13679 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13678, Nil)
}
__typedArg0 := tmp13678
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13680 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13650, tmp13679)
}
__typedArg0 := tmp13650
__typedArg1 := tmp13679
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13681 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdo, tmp13680)
}
__typedArg0 := symdo
__typedArg1 := tmp13680
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13682 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13681, Nil)
}
__typedArg0 := tmp13681
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13683 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13649, tmp13682)
}
__typedArg0 := tmp13649
__typedArg1 := tmp13682
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13684 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdo, tmp13683)
}
__typedArg0 := symdo
__typedArg1 := tmp13683
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13685 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13684, Nil)
}
__typedArg0 := tmp13684
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13686 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13641, tmp13685)
}
__typedArg0 := tmp13641
__typedArg1 := tmp13685
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdo, tmp13686)
}
__typedArg0 := symdo
__typedArg1 := tmp13686
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 3)

tmp13687 := Call(__e, ns2_1set, symshen_4insert_1tracking_1code, tmp13633)


_ = tmp13687

tmp13688 := MakeNative(func(__e *ControlFlow) {
V5899 := __e.Get(1)
_ = V5899
V5900 := __e.Get(2)
_ = V5900
tmp13691 := Call(__e, PrimFunc(symoccurrences), symshen_4incinfs, V5899)


tmp13692 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp13691, MakeInteger(0))
}
__typedArg0 := tmp13691
__typedArg1 := MakeInteger(0)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13692 {
__e.Return(V5900)
return
} else {
tmp13689 := Call(__e, PrimFunc(symshen_4vector_1parameter), V5900)


__e.TailApply(PrimFunc(symshen_4vector_1dereference), V5900, tmp13689)
return


}


}, 2)

tmp13693 := Call(__e, ns2_1set, symshen_4prolog_1track, tmp13688)


_ = tmp13693

tmp13694 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__selftop:
V5903 := __self1
_ = V5903
tmp13723 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V5903)
}
__typedArg0 := Nil
__typedArg1 := V5903
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13723 {
__e.Return(Nil)
return
} else {
tmp13721 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5903)
}
__typedArg0 := V5903
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13699 Obj

if True == tmp13721 {
tmp13719 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5903)
}
__typedArg0 := V5903
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13720 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13719)
}
__typedArg0 := tmp13719
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13701 Obj

if True == tmp13720 {
tmp13716 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5903)
}
__typedArg0 := V5903
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13717 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13716)
}
__typedArg0 := tmp13716
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13718 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13717)
}
__typedArg0 := tmp13717
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13703 Obj

if True == tmp13718 {
tmp13712 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5903)
}
__typedArg0 := V5903
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13713 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13712)
}
__typedArg0 := tmp13712
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13714 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13713)
}
__typedArg0 := tmp13713
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13715 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13714)
}
__typedArg0 := tmp13714
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13705 Obj

if True == tmp13715 {
tmp13707 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5903)
}
__typedArg0 := V5903
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13708 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13707)
}
__typedArg0 := tmp13707
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13709 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13708)
}
__typedArg0 := tmp13708
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13710 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13709)
}
__typedArg0 := tmp13709
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13711 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp13710)
}
__typedArg0 := Nil
__typedArg1 := tmp13710
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13706 Obj

if True == tmp13711 {
ifres13706 = True


} else {
ifres13706 = False


}

ifres13705 = ifres13706


} else {
ifres13705 = False


}

var ifres13704 Obj

if True == ifres13705 {
ifres13704 = True


} else {
ifres13704 = False


}

ifres13703 = ifres13704


} else {
ifres13703 = False


}

var ifres13702 Obj

if True == ifres13703 {
ifres13702 = True


} else {
ifres13702 = False


}

ifres13701 = ifres13702


} else {
ifres13701 = False


}

var ifres13700 Obj

if True == ifres13701 {
ifres13700 = True


} else {
ifres13700 = False


}

ifres13699 = ifres13700


} else {
ifres13699 = False


}

if True == ifres13699 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5903)
}
__typedArg0 := V5903
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return
} else {
tmp13697 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5903)
}
__typedArg0 := V5903
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp13697 {
tmp13695 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5903)
}
__typedArg0 := V5903
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4vector_1parameter) == __self {
__self1 = tmp13695
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4vector_1parameter), tmp13695)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4vector_1parameter)
return
}


}


}


}, 1)

tmp13724 := Call(__e, ns2_1set, symshen_4vector_1parameter, tmp13694)


_ = tmp13724

tmp13725 := MakeNative(func(__e *ControlFlow) {
V5906 := __e.Get(1)
_ = V5906
V5907 := __e.Get(2)
_ = V5907
tmp13759 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V5907)
}
__typedArg0 := Nil
__typedArg1 := V5907
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13759 {
__e.Return(V5906)
return
} else {
tmp13757 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5906)
}
__typedArg0 := V5906
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13735 Obj

if True == tmp13757 {
tmp13755 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5906)
}
__typedArg0 := V5906
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13756 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13755)
}
__typedArg0 := tmp13755
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13737 Obj

if True == tmp13756 {
tmp13752 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5906)
}
__typedArg0 := V5906
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13753 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13752)
}
__typedArg0 := tmp13752
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13754 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13753)
}
__typedArg0 := tmp13753
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13739 Obj

if True == tmp13754 {
tmp13748 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5906)
}
__typedArg0 := V5906
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13749 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13748)
}
__typedArg0 := tmp13748
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13750 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13749)
}
__typedArg0 := tmp13749
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13751 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13750)
}
__typedArg0 := tmp13750
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13741 Obj

if True == tmp13751 {
tmp13743 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5906)
}
__typedArg0 := V5906
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13744 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13743)
}
__typedArg0 := tmp13743
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13745 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13744)
}
__typedArg0 := tmp13744
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13746 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13745)
}
__typedArg0 := tmp13745
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13747 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp13746)
}
__typedArg0 := Nil
__typedArg1 := tmp13746
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13742 Obj

if True == tmp13747 {
ifres13742 = True


} else {
ifres13742 = False


}

ifres13741 = ifres13742


} else {
ifres13741 = False


}

var ifres13740 Obj

if True == ifres13741 {
ifres13740 = True


} else {
ifres13740 = False


}

ifres13739 = ifres13740


} else {
ifres13739 = False


}

var ifres13738 Obj

if True == ifres13739 {
ifres13738 = True


} else {
ifres13738 = False


}

ifres13737 = ifres13738


} else {
ifres13737 = False


}

var ifres13736 Obj

if True == ifres13737 {
ifres13736 = True


} else {
ifres13736 = False


}

ifres13735 = ifres13736


} else {
ifres13735 = False


}

if True == ifres13735 {
__e.Return(V5906)
return
} else {
tmp13733 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5906)
}
__typedArg0 := V5906
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp13733 {
tmp13726 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5906)
}
__typedArg0 := V5906
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13727 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5907, Nil)
}
__typedArg0 := V5907
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13728 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13726, tmp13727)
}
__typedArg0 := tmp13726
__typedArg1 := tmp13727
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13729 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4deref, tmp13728)
}
__typedArg0 := symshen_4deref
__typedArg1 := tmp13728
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13730 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5906)
}
__typedArg0 := V5906
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13731 := Call(__e, PrimFunc(symshen_4vector_1dereference), tmp13730, V5907)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13729, tmp13731)
}
__typedArg0 := tmp13729
__typedArg1 := tmp13731
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4vector_1dereference)
return
}


}


}


}, 2)

tmp13760 := Call(__e, ns2_1set, symshen_4vector_1dereference, tmp13725)


_ = tmp13760

tmp13761 := MakeNative(func(__e *ControlFlow) {
V5910 := __e.Get(1)
_ = V5910
tmp13765 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_7, V5910)
}
__typedArg0 := sym_7
__typedArg1 := V5910
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13765 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dstep_d, True)
}
__typedArg0 := symshen_4_dstep_d
__typedArg1 := True
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return
} else {
tmp13763 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_1, V5910)
}
__typedArg0 := sym_1
__typedArg1 := V5910
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13763 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dstep_d, False)
}
__typedArg0 := symshen_4_dstep_d
__typedArg1 := False
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return
} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("step expects a + or a -.\n"))
}
__typedArg0 := MakeString("step expects a + or a -.\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 1)

tmp13766 := Call(__e, ns2_1set, symstep, tmp13761)


_ = tmp13766

tmp13767 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dstep_d)
}
__typedArg0 := symshen_4_dstep_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp13768 := Call(__e, ns2_1set, symstep_2, tmp13767)


_ = tmp13768

tmp13769 := MakeNative(func(__e *ControlFlow) {
V5913 := __e.Get(1)
_ = V5913
tmp13773 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_7, V5913)
}
__typedArg0 := sym_7
__typedArg1 := V5913
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13773 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dspy_d, True)
}
__typedArg0 := symshen_4_dspy_d
__typedArg1 := True
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return
} else {
tmp13771 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_1, V5913)
}
__typedArg0 := sym_1
__typedArg1 := V5913
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13771 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dspy_d, False)
}
__typedArg0 := symshen_4_dspy_d
__typedArg1 := False
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return
} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("spy expects a + or a -.\n"))
}
__typedArg0 := MakeString("spy expects a + or a -.\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 1)

tmp13774 := Call(__e, ns2_1set, symspy, tmp13769)


_ = tmp13774

tmp13775 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dspy_d)
}
__typedArg0 := symshen_4_dspy_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp13776 := Call(__e, ns2_1set, symspy_2, tmp13775)


_ = tmp13776

tmp13777 := MakeNative(func(__e *ControlFlow) {
tmp13781 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dstep_d)
}
__typedArg0 := symshen_4_dstep_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

if True == tmp13781 {
tmp13778 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dstinput_d)
}
__typedArg0 := sym_dstinput_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp13779 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symread_1byte) {
return PrimReadByte(tmp13778)
}
__typedArg0 := tmp13778
return Call(__e, PrimFunc(symread_1byte), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4check_1byte), tmp13779)
return


} else {
__e.TailApply(PrimFunc(symnl), MakeInteger(1))
return
}


}, 0)

tmp13782 := Call(__e, ns2_1set, symshen_4terpri_1or_1read_1char, tmp13777)


_ = tmp13782

tmp13783 := MakeNative(func(__e *ControlFlow) {
V5916 := __e.Get(1)
_ = V5916
tmp13785 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(94), V5916)
}
__typedArg0 := MakeInteger(94)
__typedArg1 := V5916
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13785 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("aborted"))
}
__typedArg0 := MakeString("aborted")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
} else {
__e.Return(True)
return
}


}, 1)

tmp13786 := Call(__e, ns2_1set, symshen_4check_1byte, tmp13783)


_ = tmp13786

tmp13787 := MakeNative(func(__e *ControlFlow) {
V5917 := __e.Get(1)
_ = V5917
V5918 := __e.Get(2)
_ = V5918
V5919 := __e.Get(3)
_ = V5919
tmp13788 := Call(__e, PrimFunc(symshen_4spaces), V5917)


tmp13789 := Call(__e, PrimFunc(symshen_4spaces), V5917)


tmp13790 := Call(__e, PrimFunc(symshen_4app), tmp13789, MakeString(""), symshen_4a)


tmp13792 := Call(__e, PrimFunc(symshen_4app), V5918, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(" \n"))
__typedS1, __typedOK1 := TypedString(tmp13790)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(" \n")
__typedArg1 := tmp13790
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


tmp13794 := Call(__e, PrimFunc(symshen_4app), V5917, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("> Inputs to "))
__typedS1, __typedOK1 := TypedString(tmp13792)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("> Inputs to ")
__typedArg1 := tmp13792
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


tmp13796 := Call(__e, PrimFunc(symshen_4app), tmp13788, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("<"))
__typedS1, __typedOK1 := TypedString(tmp13794)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("<")
__typedArg1 := tmp13794
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


tmp13797 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("\n"))
__typedS1, __typedOK1 := TypedString(tmp13796)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("\n")
__typedArg1 := tmp13796
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp13798 := Call(__e, PrimFunc(symstoutput))


tmp13799 := Call(__e, PrimFunc(sympr), tmp13797, tmp13798)


_ = tmp13799

__e.TailApply(PrimFunc(symshen_4recursively_1print), V5919)
return


}, 3)

tmp13800 := Call(__e, ns2_1set, symshen_4input_1track, tmp13787)


_ = tmp13800

tmp13801 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__selftop:
V5922 := __self1
_ = V5922
tmp13811 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V5922)
}
__typedArg0 := Nil
__typedArg1 := V5922
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13811 {
tmp13802 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), MakeString(" ==>"), tmp13802)
return


} else {
tmp13809 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5922)
}
__typedArg0 := V5922
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp13809 {
tmp13803 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5922)
}
__typedArg0 := V5922
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13804 := Call(__e, PrimFunc(symprint), tmp13803)


_ = tmp13804

tmp13805 := Call(__e, PrimFunc(symstoutput))


tmp13806 := Call(__e, PrimFunc(sympr), MakeString(", "), tmp13805)


_ = tmp13806

tmp13807 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5922)
}
__typedArg0 := V5922
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4recursively_1print) == __self {
__self1 = tmp13807
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4recursively_1print), tmp13807)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.recursively-print"))
}
__typedArg0 := MakeString("implementation error in shen.recursively-print")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 1)

tmp13812 := Call(__e, ns2_1set, symshen_4recursively_1print, tmp13801)


_ = tmp13812

tmp13813 := MakeNative(func(__e *ControlFlow) {
V5923 := __e.Get(1)
_ = V5923
tmp13817 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(0), V5923)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := V5923
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13817 {
__e.Return(MakeString(""))
return
} else {
tmp13815 := Call(__e, PrimFunc(symshen_4spaces), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(V5923)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := V5923
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})())


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(" "))
__typedS1, __typedOK1 := TypedString(tmp13815)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(" ")
__typedArg1 := tmp13815
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
return


}


}, 1)

tmp13818 := Call(__e, ns2_1set, symshen_4spaces, tmp13813)


_ = tmp13818

tmp13819 := MakeNative(func(__e *ControlFlow) {
V5924 := __e.Get(1)
_ = V5924
V5925 := __e.Get(2)
_ = V5925
V5926 := __e.Get(3)
_ = V5926
tmp13820 := Call(__e, PrimFunc(symshen_4spaces), V5924)


tmp13821 := Call(__e, PrimFunc(symshen_4spaces), V5924)


tmp13822 := Call(__e, PrimFunc(symshen_4app), V5926, MakeString(""), symshen_4s)


tmp13824 := Call(__e, PrimFunc(symshen_4app), tmp13821, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("==> "))
__typedS1, __typedOK1 := TypedString(tmp13822)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("==> ")
__typedArg1 := tmp13822
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


tmp13826 := Call(__e, PrimFunc(symshen_4app), V5925, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(" \n"))
__typedS1, __typedOK1 := TypedString(tmp13824)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(" \n")
__typedArg1 := tmp13824
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


tmp13828 := Call(__e, PrimFunc(symshen_4app), V5924, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("> Output from "))
__typedS1, __typedOK1 := TypedString(tmp13826)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("> Output from ")
__typedArg1 := tmp13826
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


tmp13830 := Call(__e, PrimFunc(symshen_4app), tmp13820, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("<"))
__typedS1, __typedOK1 := TypedString(tmp13828)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("<")
__typedArg1 := tmp13828
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


tmp13831 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("\n"))
__typedS1, __typedOK1 := TypedString(tmp13830)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("\n")
__typedArg1 := tmp13830
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp13832 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp13831, tmp13832)
return


}, 3)

tmp13833 := Call(__e, ns2_1set, symshen_4output_1track, tmp13819)


_ = tmp13833

tmp13834 := MakeNative(func(__e *ControlFlow) {
V5927 := __e.Get(1)
_ = V5927
tmp13835 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dtracking_d)
}
__typedArg0 := symshen_4_dtracking_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp13836 := Call(__e, PrimFunc(symremove), V5927, tmp13835)


tmp13837 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dtracking_d, tmp13836)
}
__typedArg0 := symshen_4_dtracking_d
__typedArg1 := tmp13836
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp13837

tmp13838 := MakeNative(func(__e *ControlFlow) {
tmp13839 := Call(__e, PrimFunc(symps), V5927)


__e.TailApply(PrimFunc(symeval), tmp13839)
return


}, 0)

tmp13840 := MakeNative(func(__e *ControlFlow) {
Z5928 := __e.Get(1)
_ = Z5928
__e.Return(V5927)
return
}, 1)

tmp13841 := Call(__e, try_1catch, tmp13838, tmp13840)


_ = tmp13841

__e.Return(V5927)
return


}, 1)

tmp13842 := Call(__e, ns2_1set, symuntrack, tmp13834)


_ = tmp13842

tmp13843 := MakeNative(func(__e *ControlFlow) {
V5929 := __e.Get(1)
_ = V5929
V5930 := __e.Get(2)
_ = V5930
__e.TailApply(PrimFunc(symshen_4remove_1h), V5929, V5930, Nil)
return
}, 2)

tmp13844 := Call(__e, ns2_1set, symremove, tmp13843)


_ = tmp13844

tmp13845 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__self3 := __e.Get(3)
__selftop:
V5940 := __self1
_ = V5940
V5941 := __self2
_ = V5941
V5942 := __self3
_ = V5942
tmp13860 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V5941)
}
__typedArg0 := Nil
__typedArg1 := V5941
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13860 {
__e.TailApply(PrimFunc(symreverse), V5942)
return
} else {
tmp13858 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5941)
}
__typedArg0 := V5941
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13854 Obj

if True == tmp13858 {
tmp13856 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5941)
}
__typedArg0 := V5941
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13857 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V5940, tmp13856)
}
__typedArg0 := V5940
__typedArg1 := tmp13856
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13855 Obj

if True == tmp13857 {
ifres13855 = True


} else {
ifres13855 = False


}

ifres13854 = ifres13855


} else {
ifres13854 = False


}

if True == ifres13854 {
tmp13846 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5941)
}
__typedArg0 := V5941
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13847 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5941)
}
__typedArg0 := V5941
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4remove_1h) == __self {
__self1, __self2, __self3 = tmp13846, tmp13847, V5942
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4remove_1h), tmp13846, tmp13847, V5942)
return


} else {
tmp13852 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5941)
}
__typedArg0 := V5941
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp13852 {
tmp13848 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5941)
}
__typedArg0 := V5941
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13849 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5941)
}
__typedArg0 := V5941
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13850 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13849, V5942)
}
__typedArg0 := tmp13849
__typedArg1 := V5942
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

if PrimFunc(symshen_4remove_1h) == __self {
__self1, __self2, __self3 = V5940, tmp13848, tmp13850
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4remove_1h), V5940, tmp13848, tmp13850)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.remove-h"))
}
__typedArg0 := MakeString("implementation error in shen.remove-h")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}


}, 3)

tmp13861 := Call(__e, ns2_1set, symshen_4remove_1h, tmp13845)


_ = tmp13861

tmp13862 := MakeNative(func(__e *ControlFlow) {
V5943 := __e.Get(1)
_ = V5943
tmp13863 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dprofiled_d)
}
__typedArg0 := symshen_4_dprofiled_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp13864 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5943, tmp13863)
}
__typedArg0 := V5943
__typedArg1 := tmp13863
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13865 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dprofiled_d, tmp13864)
}
__typedArg0 := symshen_4_dprofiled_d
__typedArg1 := tmp13864
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp13865

tmp13866 := Call(__e, PrimFunc(symps), V5943)


__e.TailApply(PrimFunc(symshen_4profile_1help), tmp13866)
return


}, 1)

tmp13867 := Call(__e, ns2_1set, symprofile, tmp13862)


_ = tmp13867

tmp13868 := MakeNative(func(__e *ControlFlow) {
V5946 := __e.Get(1)
_ = V5946
tmp13933 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5946)
}
__typedArg0 := V5946
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13907 Obj

if True == tmp13933 {
tmp13931 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5946)
}
__typedArg0 := V5946
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13932 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefun, tmp13931)
}
__typedArg0 := symdefun
__typedArg1 := tmp13931
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13909 Obj

if True == tmp13932 {
tmp13929 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5946)
}
__typedArg0 := V5946
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13930 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13929)
}
__typedArg0 := tmp13929
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13911 Obj

if True == tmp13930 {
tmp13926 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5946)
}
__typedArg0 := V5946
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13927 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13926)
}
__typedArg0 := tmp13926
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13928 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13927)
}
__typedArg0 := tmp13927
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13913 Obj

if True == tmp13928 {
tmp13922 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5946)
}
__typedArg0 := V5946
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13923 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13922)
}
__typedArg0 := tmp13922
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13924 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13923)
}
__typedArg0 := tmp13923
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13925 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13924)
}
__typedArg0 := tmp13924
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13915 Obj

if True == tmp13925 {
tmp13917 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5946)
}
__typedArg0 := V5946
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13918 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13917)
}
__typedArg0 := tmp13917
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13919 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13918)
}
__typedArg0 := tmp13918
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13920 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13919)
}
__typedArg0 := tmp13919
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13921 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp13920)
}
__typedArg0 := Nil
__typedArg1 := tmp13920
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13916 Obj

if True == tmp13921 {
ifres13916 = True


} else {
ifres13916 = False


}

ifres13915 = ifres13916


} else {
ifres13915 = False


}

var ifres13914 Obj

if True == ifres13915 {
ifres13914 = True


} else {
ifres13914 = False


}

ifres13913 = ifres13914


} else {
ifres13913 = False


}

var ifres13912 Obj

if True == ifres13913 {
ifres13912 = True


} else {
ifres13912 = False


}

ifres13911 = ifres13912


} else {
ifres13911 = False


}

var ifres13910 Obj

if True == ifres13911 {
ifres13910 = True


} else {
ifres13910 = False


}

ifres13909 = ifres13910


} else {
ifres13909 = False


}

var ifres13908 Obj

if True == ifres13909 {
ifres13908 = True


} else {
ifres13908 = False


}

ifres13907 = ifres13908


} else {
ifres13907 = False


}

if True == ifres13907 {
tmp13869 := Call(__e, PrimFunc(symgensym), symshen_4f)


let__13548 := tmp13869
_ = let__13548

tmp13870 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5946)
}
__typedArg0 := V5946
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13871 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13870)
}
__typedArg0 := tmp13870
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13872 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5946)
}
__typedArg0 := V5946
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13873 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13872)
}
__typedArg0 := tmp13872
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13874 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13873)
}
__typedArg0 := tmp13873
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13875 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5946)
}
__typedArg0 := V5946
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13876 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13875)
}
__typedArg0 := tmp13875
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13877 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5946)
}
__typedArg0 := V5946
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13878 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13877)
}
__typedArg0 := tmp13877
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13879 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13878)
}
__typedArg0 := tmp13878
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13880 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5946)
}
__typedArg0 := V5946
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13881 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13880)
}
__typedArg0 := tmp13880
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13882 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13881)
}
__typedArg0 := tmp13881
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13883 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__13548, tmp13882)
}
__typedArg0 := let__13548
__typedArg1 := tmp13882
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13884 := Call(__e, PrimFunc(symshen_4profile_1func), tmp13876, tmp13879, tmp13883)


tmp13885 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13884, Nil)
}
__typedArg0 := tmp13884
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13886 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13874, tmp13885)
}
__typedArg0 := tmp13874
__typedArg1 := tmp13885
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13887 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13871, tmp13886)
}
__typedArg0 := tmp13871
__typedArg1 := tmp13886
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13888 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefun, tmp13887)
}
__typedArg0 := symdefun
__typedArg1 := tmp13887
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__13549 := tmp13888
_ = let__13549

tmp13889 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5946)
}
__typedArg0 := V5946
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13890 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13889)
}
__typedArg0 := tmp13889
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13891 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13890)
}
__typedArg0 := tmp13890
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13892 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5946)
}
__typedArg0 := V5946
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13893 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13892)
}
__typedArg0 := tmp13892
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13894 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5946)
}
__typedArg0 := V5946
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13895 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13894)
}
__typedArg0 := tmp13894
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13896 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13895)
}
__typedArg0 := tmp13895
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13897 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13896)
}
__typedArg0 := tmp13896
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13898 := Call(__e, PrimFunc(symsubst), let__13548, tmp13893, tmp13897)


tmp13899 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13898, Nil)
}
__typedArg0 := tmp13898
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13900 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13891, tmp13899)
}
__typedArg0 := tmp13891
__typedArg1 := tmp13899
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13901 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__13548, tmp13900)
}
__typedArg0 := let__13548
__typedArg1 := tmp13900
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13902 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefun, tmp13901)
}
__typedArg0 := symdefun
__typedArg1 := tmp13901
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__13550 := tmp13902
_ = let__13550

tmp13903 := Call(__e, PrimFunc(symeval_1kl), let__13549)


let__13551 := tmp13903
_ = let__13551

tmp13904 := Call(__e, PrimFunc(symeval_1kl), let__13550)


let__13552 := tmp13904
_ = let__13552

tmp13905 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5946)
}
__typedArg0 := V5946
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13905)
}
__typedArg0 := tmp13905
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("Cannot profile.\n"))
}
__typedArg0 := MakeString("Cannot profile.\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}, 1)

tmp13934 := Call(__e, ns2_1set, symshen_4profile_1help, tmp13868)


_ = tmp13934

tmp13935 := MakeNative(func(__e *ControlFlow) {
V5952 := __e.Get(1)
_ = V5952
tmp13936 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dprofiled_d)
}
__typedArg0 := symshen_4_dprofiled_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp13937 := Call(__e, PrimFunc(symremove), V5952, tmp13936)


tmp13938 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dprofiled_d, tmp13937)
}
__typedArg0 := symshen_4_dprofiled_d
__typedArg1 := tmp13937
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp13938

tmp13939 := MakeNative(func(__e *ControlFlow) {
tmp13940 := Call(__e, PrimFunc(symps), V5952)


__e.TailApply(PrimFunc(symeval), tmp13940)
return


}, 0)

tmp13941 := MakeNative(func(__e *ControlFlow) {
Z5953 := __e.Get(1)
_ = Z5953
__e.Return(V5952)
return
}, 1)

__e.TailApply(try_1catch, tmp13939, tmp13941)
return


}, 1)

tmp13942 := Call(__e, ns2_1set, symunprofile, tmp13935)


_ = tmp13942

tmp13943 := MakeNative(func(__e *ControlFlow) {
V5954 := __e.Get(1)
_ = V5954
tmp13944 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dprofiled_d)
}
__typedArg0 := symshen_4_dprofiled_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symelement_2), V5954, tmp13944)
return


}, 1)

tmp13945 := Call(__e, ns2_1set, symshen_4profiled_2, tmp13943)


_ = tmp13945

tmp13946 := MakeNative(func(__e *ControlFlow) {
V5955 := __e.Get(1)
_ = V5955
V5956 := __e.Get(2)
_ = V5956
V5957 := __e.Get(3)
_ = V5957
tmp13947 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symrun, Nil)
}
__typedArg0 := symrun
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13948 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symget_1time, tmp13947)
}
__typedArg0 := symget_1time
__typedArg1 := tmp13947
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13949 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symrun, Nil)
}
__typedArg0 := symrun
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13950 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symget_1time, tmp13949)
}
__typedArg0 := symget_1time
__typedArg1 := tmp13949
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13951 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symStart, Nil)
}
__typedArg0 := symStart
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13952 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13950, tmp13951)
}
__typedArg0 := tmp13950
__typedArg1 := tmp13951
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13953 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1, tmp13952)
}
__typedArg0 := sym_1
__typedArg1 := tmp13952
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13954 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5955, Nil)
}
__typedArg0 := V5955
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13955 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4get_1profile, tmp13954)
}
__typedArg0 := symshen_4get_1profile
__typedArg1 := tmp13954
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13956 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symFinish, Nil)
}
__typedArg0 := symFinish
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13957 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13955, tmp13956)
}
__typedArg0 := tmp13955
__typedArg1 := tmp13956
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13958 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_7, tmp13957)
}
__typedArg0 := sym_7
__typedArg1 := tmp13957
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13959 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13958, Nil)
}
__typedArg0 := tmp13958
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13960 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5955, tmp13959)
}
__typedArg0 := V5955
__typedArg1 := tmp13959
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13961 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4put_1profile, tmp13960)
}
__typedArg0 := symshen_4put_1profile
__typedArg1 := tmp13960
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13962 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symResult, Nil)
}
__typedArg0 := symResult
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13963 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13961, tmp13962)
}
__typedArg0 := tmp13961
__typedArg1 := tmp13962
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13964 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symRecord, tmp13963)
}
__typedArg0 := symRecord
__typedArg1 := tmp13963
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13965 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp13964)
}
__typedArg0 := symlet
__typedArg1 := tmp13964
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13966 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13965, Nil)
}
__typedArg0 := tmp13965
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13967 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13953, tmp13966)
}
__typedArg0 := tmp13953
__typedArg1 := tmp13966
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13968 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symFinish, tmp13967)
}
__typedArg0 := symFinish
__typedArg1 := tmp13967
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13969 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp13968)
}
__typedArg0 := symlet
__typedArg1 := tmp13968
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13970 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13969, Nil)
}
__typedArg0 := tmp13969
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13971 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5957, tmp13970)
}
__typedArg0 := V5957
__typedArg1 := tmp13970
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13972 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symResult, tmp13971)
}
__typedArg0 := symResult
__typedArg1 := tmp13971
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13973 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp13972)
}
__typedArg0 := symlet
__typedArg1 := tmp13972
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13974 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13973, Nil)
}
__typedArg0 := tmp13973
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13975 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13948, tmp13974)
}
__typedArg0 := tmp13948
__typedArg1 := tmp13974
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13976 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symStart, tmp13975)
}
__typedArg0 := symStart
__typedArg1 := tmp13975
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp13976)
}
__typedArg0 := symlet
__typedArg1 := tmp13976
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 3)

tmp13977 := Call(__e, ns2_1set, symshen_4profile_1func, tmp13946)


_ = tmp13977

tmp13978 := MakeNative(func(__e *ControlFlow) {
V5958 := __e.Get(1)
_ = V5958
tmp13979 := Call(__e, PrimFunc(symshen_4get_1profile), V5958)


let__13553 := tmp13979
_ = let__13553

tmp13980 := Call(__e, PrimFunc(symshen_4put_1profile), V5958, MakeInteger(0))


let__13554 := tmp13980
_ = let__13554

__e.TailApply(PrimFunc(sym_8p), V5958, let__13553)
return


}, 1)

tmp13981 := Call(__e, ns2_1set, symprofile_1results, tmp13978)


_ = tmp13981

tmp13982 := MakeNative(func(__e *ControlFlow) {
V5961 := __e.Get(1)
_ = V5961
tmp13983 := MakeNative(func(__e *ControlFlow) {
tmp13984 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symget), V5961, symprofile, tmp13984)
return


}, 0)

tmp13985 := MakeNative(func(__e *ControlFlow) {
Z5962 := __e.Get(1)
_ = Z5962
__e.Return(MakeInteger(0))
return
}, 1)

__e.TailApply(try_1catch, tmp13983, tmp13985)
return


}, 1)

tmp13986 := Call(__e, ns2_1set, symshen_4get_1profile, tmp13982)


_ = tmp13986

tmp13987 := MakeNative(func(__e *ControlFlow) {
V5963 := __e.Get(1)
_ = V5963
V5964 := __e.Get(2)
_ = V5964
tmp13988 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symput), V5963, symprofile, V5964, tmp13988)
return


}, 2)

__e.TailApply(ns2_1set, symshen_4put_1profile, tmp13987)
return




}, 0)

