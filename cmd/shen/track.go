package main

import . "github.com/pyrex41/shen-go/kl"

var TrackMain = MakeNative(func(__e *ControlFlow) {
tmp13535 := MakeNative(func(__e *ControlFlow) {
V5871 := __e.Get(1)
_ = V5871
tmp13536 := Call(__e, PrimFunc(symshen_4app), V5871, MakeString(";\n"), symshen_4a)


tmp13537 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("partial function "))
__typedS1, __typedOK1 := TypedString(tmp13536)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("partial function ")
__typedArg1 := tmp13536
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp13538 := Call(__e, PrimFunc(symstoutput))


tmp13539 := Call(__e, PrimFunc(sympr), tmp13537, tmp13538)


_ = tmp13539

tmp13548 := Call(__e, PrimFunc(symshen_4tracked_2), V5871)


tmp13549 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp13548)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp13548
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres13543 Obj

if True == tmp13549 {
tmp13545 := Call(__e, PrimFunc(symshen_4app), V5871, MakeString("? "), symshen_4a)


tmp13547 := Call(__e, PrimFunc(symy_1or_1n_2), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("track "))
__typedS1, __typedOK1 := TypedString(tmp13545)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("track ")
__typedArg1 := tmp13545
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())


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
tmp13541 := Call(__e, PrimFunc(symps), V5871)


tmp13542 := Call(__e, PrimFunc(symshen_4track_1function), tmp13541)


ifres13540 = tmp13542


} else {
ifres13540 = symshen_4ok


}

_ = ifres13540

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("aborted"))
}
__typedArg0 := MakeString("aborted")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}, 1)

tmp13550 := Call(__e, ns2_1set, symshen_4f_1error, tmp13535)


_ = tmp13550

tmp13551 := MakeNative(func(__e *ControlFlow) {
V5872 := __e.Get(1)
_ = V5872
tmp13552 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dtracking_d)
}
__typedArg0 := symshen_4_dtracking_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symelement_2), V5872, tmp13552)
return


}, 1)

tmp13553 := Call(__e, ns2_1set, symshen_4tracked_2, tmp13551)


_ = tmp13553

tmp13554 := MakeNative(func(__e *ControlFlow) {
V5873 := __e.Get(1)
_ = V5873
tmp13555 := Call(__e, PrimFunc(symps), V5873)


let__13524 := tmp13555
_ = let__13524

__e.TailApply(PrimFunc(symshen_4track_1function), let__13524)
return


}, 1)

tmp13556 := Call(__e, ns2_1set, symtrack, tmp13554)


_ = tmp13556

tmp13557 := MakeNative(func(__e *ControlFlow) {
V5877 := __e.Get(1)
_ = V5877
tmp13611 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5877)
}
__typedArg0 := V5877
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13585 Obj

if True == tmp13611 {
tmp13609 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5877)
}
__typedArg0 := V5877
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13610 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefun, tmp13609)
}
__typedArg0 := symdefun
__typedArg1 := tmp13609
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13587 Obj

if True == tmp13610 {
tmp13607 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5877)
}
__typedArg0 := V5877
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13608 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13607)
}
__typedArg0 := tmp13607
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13589 Obj

if True == tmp13608 {
tmp13604 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5877)
}
__typedArg0 := V5877
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13605 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13604)
}
__typedArg0 := tmp13604
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13606 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13605)
}
__typedArg0 := tmp13605
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13591 Obj

if True == tmp13606 {
tmp13600 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5877)
}
__typedArg0 := V5877
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13601 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13600)
}
__typedArg0 := tmp13600
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13602 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13601)
}
__typedArg0 := tmp13601
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13603 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13602)
}
__typedArg0 := tmp13602
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13593 Obj

if True == tmp13603 {
tmp13595 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5877)
}
__typedArg0 := V5877
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13596 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13595)
}
__typedArg0 := tmp13595
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13597 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13596)
}
__typedArg0 := tmp13596
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13598 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13597)
}
__typedArg0 := tmp13597
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13599 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp13598)
}
__typedArg0 := Nil
__typedArg1 := tmp13598
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

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
tmp13558 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5877)
}
__typedArg0 := V5877
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13559 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13558)
}
__typedArg0 := tmp13558
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13560 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5877)
}
__typedArg0 := V5877
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13561 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13560)
}
__typedArg0 := tmp13560
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13562 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13561)
}
__typedArg0 := tmp13561
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13563 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5877)
}
__typedArg0 := V5877
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13564 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13563)
}
__typedArg0 := tmp13563
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13565 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5877)
}
__typedArg0 := V5877
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13566 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13565)
}
__typedArg0 := tmp13565
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13567 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13566)
}
__typedArg0 := tmp13566
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13568 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5877)
}
__typedArg0 := V5877
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13569 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13568)
}
__typedArg0 := tmp13568
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13570 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13569)
}
__typedArg0 := tmp13569
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13571 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13570)
}
__typedArg0 := tmp13570
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13572 := Call(__e, PrimFunc(symshen_4insert_1tracking_1code), tmp13564, tmp13567, tmp13571)


tmp13573 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13572, Nil)
}
__typedArg0 := tmp13572
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13574 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13562, tmp13573)
}
__typedArg0 := tmp13562
__typedArg1 := tmp13573
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13575 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13559, tmp13574)
}
__typedArg0 := tmp13559
__typedArg1 := tmp13574
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13576 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefun, tmp13575)
}
__typedArg0 := symdefun
__typedArg1 := tmp13575
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__13525 := tmp13576
_ = let__13525

tmp13577 := Call(__e, PrimFunc(symeval_1kl), let__13525)


let__13526 := tmp13577
_ = let__13526

tmp13578 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5877)
}
__typedArg0 := V5877
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dtracking_d)
}
__typedArg0 := symshen_4_dtracking_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp13581 := Call(__e, PrimFunc(symadjoin), tmp13579, tmp13580)


tmp13582 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dtracking_d, tmp13581)
}
__typedArg0 := symshen_4_dtracking_d
__typedArg1 := tmp13581
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

let__13527 := tmp13582
_ = let__13527

tmp13583 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5877)
}
__typedArg0 := V5877
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13583)
}
__typedArg0 := tmp13583
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

tmp13612 := Call(__e, ns2_1set, symshen_4track_1function, tmp13557)


_ = tmp13612

tmp13613 := MakeNative(func(__e *ControlFlow) {
V5881 := __e.Get(1)
_ = V5881
V5882 := __e.Get(2)
_ = V5882
V5883 := __e.Get(3)
_ = V5883
tmp13614 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4_dcall_d, Nil)
}
__typedArg0 := symshen_4_dcall_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13615 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvalue, tmp13614)
}
__typedArg0 := symvalue
__typedArg1 := tmp13614
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13616 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), Nil)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13617 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13615, tmp13616)
}
__typedArg0 := tmp13615
__typedArg1 := tmp13616
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13618 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_7, tmp13617)
}
__typedArg0 := sym_7
__typedArg1 := tmp13617
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13619 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13618, Nil)
}
__typedArg0 := tmp13618
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13620 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4_dcall_d, tmp13619)
}
__typedArg0 := symshen_4_dcall_d
__typedArg1 := tmp13619
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13621 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symset, tmp13620)
}
__typedArg0 := symset
__typedArg1 := tmp13620
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13622 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4_dcall_d, Nil)
}
__typedArg0 := symshen_4_dcall_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13623 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvalue, tmp13622)
}
__typedArg0 := symvalue
__typedArg1 := tmp13622
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13624 := Call(__e, PrimFunc(symshen_4prolog_1track), V5883, V5882)


tmp13625 := Call(__e, PrimFunc(symshen_4cons_1form), tmp13624)


tmp13626 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13625, Nil)
}
__typedArg0 := tmp13625
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13627 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5881, tmp13626)
}
__typedArg0 := V5881
__typedArg1 := tmp13626
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13628 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13623, tmp13627)
}
__typedArg0 := tmp13623
__typedArg1 := tmp13627
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13629 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4input_1track, tmp13628)
}
__typedArg0 := symshen_4input_1track
__typedArg1 := tmp13628
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13630 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4terpri_1or_1read_1char, Nil)
}
__typedArg0 := symshen_4terpri_1or_1read_1char
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13631 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4_dcall_d, Nil)
}
__typedArg0 := symshen_4_dcall_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13632 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvalue, tmp13631)
}
__typedArg0 := symvalue
__typedArg1 := tmp13631
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13633 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symResult, Nil)
}
__typedArg0 := symResult
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13634 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5881, tmp13633)
}
__typedArg0 := V5881
__typedArg1 := tmp13633
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13635 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13632, tmp13634)
}
__typedArg0 := tmp13632
__typedArg1 := tmp13634
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13636 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4output_1track, tmp13635)
}
__typedArg0 := symshen_4output_1track
__typedArg1 := tmp13635
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13637 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4_dcall_d, Nil)
}
__typedArg0 := symshen_4_dcall_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13638 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvalue, tmp13637)
}
__typedArg0 := symvalue
__typedArg1 := tmp13637
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13639 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), Nil)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13640 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13638, tmp13639)
}
__typedArg0 := tmp13638
__typedArg1 := tmp13639
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13641 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1, tmp13640)
}
__typedArg0 := sym_1
__typedArg1 := tmp13640
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13642 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13641, Nil)
}
__typedArg0 := tmp13641
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13643 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4_dcall_d, tmp13642)
}
__typedArg0 := symshen_4_dcall_d
__typedArg1 := tmp13642
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13644 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symset, tmp13643)
}
__typedArg0 := symset
__typedArg1 := tmp13643
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13645 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4terpri_1or_1read_1char, Nil)
}
__typedArg0 := symshen_4terpri_1or_1read_1char
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13646 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symResult, Nil)
}
__typedArg0 := symResult
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13647 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13645, tmp13646)
}
__typedArg0 := tmp13645
__typedArg1 := tmp13646
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13648 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdo, tmp13647)
}
__typedArg0 := symdo
__typedArg1 := tmp13647
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13649 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13648, Nil)
}
__typedArg0 := tmp13648
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13650 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13644, tmp13649)
}
__typedArg0 := tmp13644
__typedArg1 := tmp13649
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13651 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdo, tmp13650)
}
__typedArg0 := symdo
__typedArg1 := tmp13650
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13652 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13651, Nil)
}
__typedArg0 := tmp13651
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13653 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13636, tmp13652)
}
__typedArg0 := tmp13636
__typedArg1 := tmp13652
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13654 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdo, tmp13653)
}
__typedArg0 := symdo
__typedArg1 := tmp13653
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13655 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13654, Nil)
}
__typedArg0 := tmp13654
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13656 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5883, tmp13655)
}
__typedArg0 := V5883
__typedArg1 := tmp13655
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13657 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symResult, tmp13656)
}
__typedArg0 := symResult
__typedArg1 := tmp13656
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13658 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp13657)
}
__typedArg0 := symlet
__typedArg1 := tmp13657
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13659 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13658, Nil)
}
__typedArg0 := tmp13658
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13660 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13630, tmp13659)
}
__typedArg0 := tmp13630
__typedArg1 := tmp13659
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13661 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdo, tmp13660)
}
__typedArg0 := symdo
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
return PrimCons(tmp13629, tmp13662)
}
__typedArg0 := tmp13629
__typedArg1 := tmp13662
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13664 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdo, tmp13663)
}
__typedArg0 := symdo
__typedArg1 := tmp13663
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13665 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13664, Nil)
}
__typedArg0 := tmp13664
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13666 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13621, tmp13665)
}
__typedArg0 := tmp13621
__typedArg1 := tmp13665
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdo, tmp13666)
}
__typedArg0 := symdo
__typedArg1 := tmp13666
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 3)

tmp13667 := Call(__e, ns2_1set, symshen_4insert_1tracking_1code, tmp13613)


_ = tmp13667

tmp13668 := MakeNative(func(__e *ControlFlow) {
V5884 := __e.Get(1)
_ = V5884
V5885 := __e.Get(2)
_ = V5885
tmp13671 := Call(__e, PrimFunc(symoccurrences), symshen_4incinfs, V5884)


tmp13672 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp13671, MakeInteger(0))
}
__typedArg0 := tmp13671
__typedArg1 := MakeInteger(0)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13672 {
__e.Return(V5885)
return
} else {
tmp13669 := Call(__e, PrimFunc(symshen_4vector_1parameter), V5885)


__e.TailApply(PrimFunc(symshen_4vector_1dereference), V5885, tmp13669)
return


}


}, 2)

tmp13673 := Call(__e, ns2_1set, symshen_4prolog_1track, tmp13668)


_ = tmp13673

tmp13674 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__selftop:
V5888 := __self1
_ = V5888
tmp13703 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V5888)
}
__typedArg0 := Nil
__typedArg1 := V5888
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13703 {
__e.Return(Nil)
return
} else {
tmp13701 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5888)
}
__typedArg0 := V5888
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13679 Obj

if True == tmp13701 {
tmp13699 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5888)
}
__typedArg0 := V5888
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13700 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13699)
}
__typedArg0 := tmp13699
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13681 Obj

if True == tmp13700 {
tmp13696 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5888)
}
__typedArg0 := V5888
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13697 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13696)
}
__typedArg0 := tmp13696
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13698 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13697)
}
__typedArg0 := tmp13697
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13683 Obj

if True == tmp13698 {
tmp13692 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5888)
}
__typedArg0 := V5888
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13693 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13692)
}
__typedArg0 := tmp13692
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13694 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13693)
}
__typedArg0 := tmp13693
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13695 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13694)
}
__typedArg0 := tmp13694
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13685 Obj

if True == tmp13695 {
tmp13687 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5888)
}
__typedArg0 := V5888
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13688 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13687)
}
__typedArg0 := tmp13687
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13689 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13688)
}
__typedArg0 := tmp13688
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13690 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13689)
}
__typedArg0 := tmp13689
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13691 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp13690)
}
__typedArg0 := Nil
__typedArg1 := tmp13690
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

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
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5888)
}
__typedArg0 := V5888
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return
} else {
tmp13677 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5888)
}
__typedArg0 := V5888
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp13677 {
tmp13675 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5888)
}
__typedArg0 := V5888
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4vector_1parameter) == __self {
__self1 = tmp13675
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4vector_1parameter), tmp13675)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4vector_1parameter)
return
}


}


}


}, 1)

tmp13704 := Call(__e, ns2_1set, symshen_4vector_1parameter, tmp13674)


_ = tmp13704

tmp13705 := MakeNative(func(__e *ControlFlow) {
V5891 := __e.Get(1)
_ = V5891
V5892 := __e.Get(2)
_ = V5892
tmp13739 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V5892)
}
__typedArg0 := Nil
__typedArg1 := V5892
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13739 {
__e.Return(V5891)
return
} else {
tmp13737 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5891)
}
__typedArg0 := V5891
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13715 Obj

if True == tmp13737 {
tmp13735 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5891)
}
__typedArg0 := V5891
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13736 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13735)
}
__typedArg0 := tmp13735
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13717 Obj

if True == tmp13736 {
tmp13732 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5891)
}
__typedArg0 := V5891
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13733 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13732)
}
__typedArg0 := tmp13732
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13734 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13733)
}
__typedArg0 := tmp13733
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13719 Obj

if True == tmp13734 {
tmp13728 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5891)
}
__typedArg0 := V5891
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13729 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13728)
}
__typedArg0 := tmp13728
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13730 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13729)
}
__typedArg0 := tmp13729
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13731 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13730)
}
__typedArg0 := tmp13730
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13721 Obj

if True == tmp13731 {
tmp13723 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5891)
}
__typedArg0 := V5891
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13724 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13723)
}
__typedArg0 := tmp13723
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13725 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13724)
}
__typedArg0 := tmp13724
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13726 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13725)
}
__typedArg0 := tmp13725
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13727 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp13726)
}
__typedArg0 := Nil
__typedArg1 := tmp13726
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

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
__e.Return(V5891)
return
} else {
tmp13713 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5891)
}
__typedArg0 := V5891
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp13713 {
tmp13706 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5891)
}
__typedArg0 := V5891
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13707 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5892, Nil)
}
__typedArg0 := V5892
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13708 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13706, tmp13707)
}
__typedArg0 := tmp13706
__typedArg1 := tmp13707
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13709 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4deref, tmp13708)
}
__typedArg0 := symshen_4deref
__typedArg1 := tmp13708
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13710 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5891)
}
__typedArg0 := V5891
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13711 := Call(__e, PrimFunc(symshen_4vector_1dereference), tmp13710, V5892)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13709, tmp13711)
}
__typedArg0 := tmp13709
__typedArg1 := tmp13711
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

tmp13740 := Call(__e, ns2_1set, symshen_4vector_1dereference, tmp13705)


_ = tmp13740

tmp13741 := MakeNative(func(__e *ControlFlow) {
V5895 := __e.Get(1)
_ = V5895
tmp13745 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_7, V5895)
}
__typedArg0 := sym_7
__typedArg1 := V5895
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13745 {
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
tmp13743 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_1, V5895)
}
__typedArg0 := sym_1
__typedArg1 := V5895
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13743 {
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

tmp13746 := Call(__e, ns2_1set, symstep, tmp13741)


_ = tmp13746

tmp13747 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dstep_d)
}
__typedArg0 := symshen_4_dstep_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp13748 := Call(__e, ns2_1set, symstep_2, tmp13747)


_ = tmp13748

tmp13749 := MakeNative(func(__e *ControlFlow) {
V5898 := __e.Get(1)
_ = V5898
tmp13753 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_7, V5898)
}
__typedArg0 := sym_7
__typedArg1 := V5898
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13753 {
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
tmp13751 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_1, V5898)
}
__typedArg0 := sym_1
__typedArg1 := V5898
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13751 {
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

tmp13754 := Call(__e, ns2_1set, symspy, tmp13749)


_ = tmp13754

tmp13755 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dspy_d)
}
__typedArg0 := symshen_4_dspy_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp13756 := Call(__e, ns2_1set, symspy_2, tmp13755)


_ = tmp13756

tmp13757 := MakeNative(func(__e *ControlFlow) {
tmp13761 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dstep_d)
}
__typedArg0 := symshen_4_dstep_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

if True == tmp13761 {
tmp13758 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dstinput_d)
}
__typedArg0 := sym_dstinput_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp13759 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symread_1byte) {
return PrimReadByte(tmp13758)
}
__typedArg0 := tmp13758
return Call(__e, PrimFunc(symread_1byte), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4check_1byte), tmp13759)
return


} else {
__e.TailApply(PrimFunc(symnl), MakeInteger(1))
return
}


}, 0)

tmp13762 := Call(__e, ns2_1set, symshen_4terpri_1or_1read_1char, tmp13757)


_ = tmp13762

tmp13763 := MakeNative(func(__e *ControlFlow) {
V5901 := __e.Get(1)
_ = V5901
tmp13765 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(94), V5901)
}
__typedArg0 := MakeInteger(94)
__typedArg1 := V5901
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13765 {
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

tmp13766 := Call(__e, ns2_1set, symshen_4check_1byte, tmp13763)


_ = tmp13766

tmp13767 := MakeNative(func(__e *ControlFlow) {
V5902 := __e.Get(1)
_ = V5902
V5903 := __e.Get(2)
_ = V5903
V5904 := __e.Get(3)
_ = V5904
tmp13768 := Call(__e, PrimFunc(symshen_4spaces), V5902)


tmp13769 := Call(__e, PrimFunc(symshen_4spaces), V5902)


tmp13770 := Call(__e, PrimFunc(symshen_4app), tmp13769, MakeString(""), symshen_4a)


tmp13772 := Call(__e, PrimFunc(symshen_4app), V5903, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(" \n"))
__typedS1, __typedOK1 := TypedString(tmp13770)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(" \n")
__typedArg1 := tmp13770
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


tmp13774 := Call(__e, PrimFunc(symshen_4app), V5902, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("> Inputs to "))
__typedS1, __typedOK1 := TypedString(tmp13772)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("> Inputs to ")
__typedArg1 := tmp13772
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


tmp13776 := Call(__e, PrimFunc(symshen_4app), tmp13768, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("<"))
__typedS1, __typedOK1 := TypedString(tmp13774)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("<")
__typedArg1 := tmp13774
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


tmp13777 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("\n"))
__typedS1, __typedOK1 := TypedString(tmp13776)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("\n")
__typedArg1 := tmp13776
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp13778 := Call(__e, PrimFunc(symstoutput))


tmp13779 := Call(__e, PrimFunc(sympr), tmp13777, tmp13778)


_ = tmp13779

__e.TailApply(PrimFunc(symshen_4recursively_1print), V5904)
return


}, 3)

tmp13780 := Call(__e, ns2_1set, symshen_4input_1track, tmp13767)


_ = tmp13780

tmp13781 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__selftop:
V5907 := __self1
_ = V5907
tmp13791 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V5907)
}
__typedArg0 := Nil
__typedArg1 := V5907
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13791 {
tmp13782 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), MakeString(" ==>"), tmp13782)
return


} else {
tmp13789 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5907)
}
__typedArg0 := V5907
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp13789 {
tmp13783 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5907)
}
__typedArg0 := V5907
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13784 := Call(__e, PrimFunc(symprint), tmp13783)


_ = tmp13784

tmp13785 := Call(__e, PrimFunc(symstoutput))


tmp13786 := Call(__e, PrimFunc(sympr), MakeString(", "), tmp13785)


_ = tmp13786

tmp13787 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5907)
}
__typedArg0 := V5907
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4recursively_1print) == __self {
__self1 = tmp13787
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4recursively_1print), tmp13787)
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

tmp13792 := Call(__e, ns2_1set, symshen_4recursively_1print, tmp13781)


_ = tmp13792

tmp13793 := MakeNative(func(__e *ControlFlow) {
V5908 := __e.Get(1)
_ = V5908
tmp13797 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(0), V5908)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := V5908
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13797 {
__e.Return(MakeString(""))
return
} else {
tmp13795 := Call(__e, PrimFunc(symshen_4spaces), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(V5908)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := V5908
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})())


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(" "))
__typedS1, __typedOK1 := TypedString(tmp13795)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(" ")
__typedArg1 := tmp13795
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
return


}


}, 1)

tmp13798 := Call(__e, ns2_1set, symshen_4spaces, tmp13793)


_ = tmp13798

tmp13799 := MakeNative(func(__e *ControlFlow) {
V5909 := __e.Get(1)
_ = V5909
V5910 := __e.Get(2)
_ = V5910
V5911 := __e.Get(3)
_ = V5911
tmp13800 := Call(__e, PrimFunc(symshen_4spaces), V5909)


tmp13801 := Call(__e, PrimFunc(symshen_4spaces), V5909)


tmp13802 := Call(__e, PrimFunc(symshen_4app), V5911, MakeString(""), symshen_4s)


tmp13804 := Call(__e, PrimFunc(symshen_4app), tmp13801, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("==> "))
__typedS1, __typedOK1 := TypedString(tmp13802)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("==> ")
__typedArg1 := tmp13802
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


tmp13806 := Call(__e, PrimFunc(symshen_4app), V5910, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(" \n"))
__typedS1, __typedOK1 := TypedString(tmp13804)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(" \n")
__typedArg1 := tmp13804
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


tmp13808 := Call(__e, PrimFunc(symshen_4app), V5909, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("> Output from "))
__typedS1, __typedOK1 := TypedString(tmp13806)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("> Output from ")
__typedArg1 := tmp13806
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


tmp13810 := Call(__e, PrimFunc(symshen_4app), tmp13800, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("<"))
__typedS1, __typedOK1 := TypedString(tmp13808)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("<")
__typedArg1 := tmp13808
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


tmp13811 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("\n"))
__typedS1, __typedOK1 := TypedString(tmp13810)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("\n")
__typedArg1 := tmp13810
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp13812 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp13811, tmp13812)
return


}, 3)

tmp13813 := Call(__e, ns2_1set, symshen_4output_1track, tmp13799)


_ = tmp13813

tmp13814 := MakeNative(func(__e *ControlFlow) {
V5912 := __e.Get(1)
_ = V5912
tmp13815 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dtracking_d)
}
__typedArg0 := symshen_4_dtracking_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp13816 := Call(__e, PrimFunc(symremove), V5912, tmp13815)


tmp13817 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dtracking_d, tmp13816)
}
__typedArg0 := symshen_4_dtracking_d
__typedArg1 := tmp13816
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp13817

tmp13818 := MakeNative(func(__e *ControlFlow) {
tmp13819 := Call(__e, PrimFunc(symps), V5912)


__e.TailApply(PrimFunc(symeval), tmp13819)
return


}, 0)

tmp13820 := MakeNative(func(__e *ControlFlow) {
Z5913 := __e.Get(1)
_ = Z5913
__e.Return(V5912)
return
}, 1)

tmp13821 := Call(__e, try_1catch, tmp13818, tmp13820)


_ = tmp13821

__e.Return(V5912)
return


}, 1)

tmp13822 := Call(__e, ns2_1set, symuntrack, tmp13814)


_ = tmp13822

tmp13823 := MakeNative(func(__e *ControlFlow) {
V5914 := __e.Get(1)
_ = V5914
V5915 := __e.Get(2)
_ = V5915
__e.TailApply(PrimFunc(symshen_4remove_1h), V5914, V5915, Nil)
return
}, 2)

tmp13824 := Call(__e, ns2_1set, symremove, tmp13823)


_ = tmp13824

tmp13825 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__self3 := __e.Get(3)
__selftop:
V5925 := __self1
_ = V5925
V5926 := __self2
_ = V5926
V5927 := __self3
_ = V5927
tmp13840 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V5926)
}
__typedArg0 := Nil
__typedArg1 := V5926
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13840 {
__e.TailApply(PrimFunc(symreverse), V5927)
return
} else {
tmp13838 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5926)
}
__typedArg0 := V5926
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13834 Obj

if True == tmp13838 {
tmp13836 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5926)
}
__typedArg0 := V5926
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13837 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V5925, tmp13836)
}
__typedArg0 := V5925
__typedArg1 := tmp13836
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

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
tmp13826 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5926)
}
__typedArg0 := V5926
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13827 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5926)
}
__typedArg0 := V5926
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4remove_1h) == __self {
__self1, __self2, __self3 = tmp13826, tmp13827, V5927
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4remove_1h), tmp13826, tmp13827, V5927)
return


} else {
tmp13832 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5926)
}
__typedArg0 := V5926
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp13832 {
tmp13828 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5926)
}
__typedArg0 := V5926
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13829 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5926)
}
__typedArg0 := V5926
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13830 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13829, V5927)
}
__typedArg0 := tmp13829
__typedArg1 := V5927
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

if PrimFunc(symshen_4remove_1h) == __self {
__self1, __self2, __self3 = V5925, tmp13828, tmp13830
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4remove_1h), V5925, tmp13828, tmp13830)
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

tmp13841 := Call(__e, ns2_1set, symshen_4remove_1h, tmp13825)


_ = tmp13841

tmp13842 := MakeNative(func(__e *ControlFlow) {
V5928 := __e.Get(1)
_ = V5928
tmp13843 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dprofiled_d)
}
__typedArg0 := symshen_4_dprofiled_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp13844 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5928, tmp13843)
}
__typedArg0 := V5928
__typedArg1 := tmp13843
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13845 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dprofiled_d, tmp13844)
}
__typedArg0 := symshen_4_dprofiled_d
__typedArg1 := tmp13844
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp13845

tmp13846 := Call(__e, PrimFunc(symps), V5928)


__e.TailApply(PrimFunc(symshen_4profile_1help), tmp13846)
return


}, 1)

tmp13847 := Call(__e, ns2_1set, symprofile, tmp13842)


_ = tmp13847

tmp13848 := MakeNative(func(__e *ControlFlow) {
V5931 := __e.Get(1)
_ = V5931
tmp13913 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5931)
}
__typedArg0 := V5931
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13887 Obj

if True == tmp13913 {
tmp13911 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5931)
}
__typedArg0 := V5931
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13912 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefun, tmp13911)
}
__typedArg0 := symdefun
__typedArg1 := tmp13911
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13889 Obj

if True == tmp13912 {
tmp13909 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5931)
}
__typedArg0 := V5931
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13910 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13909)
}
__typedArg0 := tmp13909
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13891 Obj

if True == tmp13910 {
tmp13906 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5931)
}
__typedArg0 := V5931
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13907 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13906)
}
__typedArg0 := tmp13906
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13908 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13907)
}
__typedArg0 := tmp13907
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13893 Obj

if True == tmp13908 {
tmp13902 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5931)
}
__typedArg0 := V5931
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13903 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13902)
}
__typedArg0 := tmp13902
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13904 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13903)
}
__typedArg0 := tmp13903
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13905 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13904)
}
__typedArg0 := tmp13904
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13895 Obj

if True == tmp13905 {
tmp13897 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5931)
}
__typedArg0 := V5931
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13898 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13897)
}
__typedArg0 := tmp13897
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13899 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13898)
}
__typedArg0 := tmp13898
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13900 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13899)
}
__typedArg0 := tmp13899
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13901 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp13900)
}
__typedArg0 := Nil
__typedArg1 := tmp13900
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

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
tmp13849 := Call(__e, PrimFunc(symgensym), symshen_4f)


let__13528 := tmp13849
_ = let__13528

tmp13850 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5931)
}
__typedArg0 := V5931
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13851 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13850)
}
__typedArg0 := tmp13850
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13852 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5931)
}
__typedArg0 := V5931
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13853 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13852)
}
__typedArg0 := tmp13852
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13854 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13853)
}
__typedArg0 := tmp13853
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13855 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5931)
}
__typedArg0 := V5931
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13856 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13855)
}
__typedArg0 := tmp13855
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13857 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5931)
}
__typedArg0 := V5931
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13858 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13857)
}
__typedArg0 := tmp13857
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13859 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13858)
}
__typedArg0 := tmp13858
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13860 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5931)
}
__typedArg0 := V5931
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13861 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13860)
}
__typedArg0 := tmp13860
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13862 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13861)
}
__typedArg0 := tmp13861
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13863 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__13528, tmp13862)
}
__typedArg0 := let__13528
__typedArg1 := tmp13862
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13864 := Call(__e, PrimFunc(symshen_4profile_1func), tmp13856, tmp13859, tmp13863)


tmp13865 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13864, Nil)
}
__typedArg0 := tmp13864
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13866 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13854, tmp13865)
}
__typedArg0 := tmp13854
__typedArg1 := tmp13865
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13867 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13851, tmp13866)
}
__typedArg0 := tmp13851
__typedArg1 := tmp13866
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13868 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefun, tmp13867)
}
__typedArg0 := symdefun
__typedArg1 := tmp13867
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__13529 := tmp13868
_ = let__13529

tmp13869 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5931)
}
__typedArg0 := V5931
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13870 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13869)
}
__typedArg0 := tmp13869
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
return PrimTail(V5931)
}
__typedArg0 := V5931
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13873 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13872)
}
__typedArg0 := tmp13872
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13874 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5931)
}
__typedArg0 := V5931
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13875 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13874)
}
__typedArg0 := tmp13874
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13876 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13875)
}
__typedArg0 := tmp13875
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13877 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13876)
}
__typedArg0 := tmp13876
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13878 := Call(__e, PrimFunc(symsubst), let__13528, tmp13873, tmp13877)


tmp13879 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13878, Nil)
}
__typedArg0 := tmp13878
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13880 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13871, tmp13879)
}
__typedArg0 := tmp13871
__typedArg1 := tmp13879
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13881 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__13528, tmp13880)
}
__typedArg0 := let__13528
__typedArg1 := tmp13880
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13882 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefun, tmp13881)
}
__typedArg0 := symdefun
__typedArg1 := tmp13881
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__13530 := tmp13882
_ = let__13530

tmp13883 := Call(__e, PrimFunc(symeval_1kl), let__13529)


let__13531 := tmp13883
_ = let__13531

tmp13884 := Call(__e, PrimFunc(symeval_1kl), let__13530)


let__13532 := tmp13884
_ = let__13532

tmp13885 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5931)
}
__typedArg0 := V5931
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13885)
}
__typedArg0 := tmp13885
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

tmp13914 := Call(__e, ns2_1set, symshen_4profile_1help, tmp13848)


_ = tmp13914

tmp13915 := MakeNative(func(__e *ControlFlow) {
V5937 := __e.Get(1)
_ = V5937
tmp13916 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dprofiled_d)
}
__typedArg0 := symshen_4_dprofiled_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp13917 := Call(__e, PrimFunc(symremove), V5937, tmp13916)


tmp13918 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dprofiled_d, tmp13917)
}
__typedArg0 := symshen_4_dprofiled_d
__typedArg1 := tmp13917
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp13918

tmp13919 := MakeNative(func(__e *ControlFlow) {
tmp13920 := Call(__e, PrimFunc(symps), V5937)


__e.TailApply(PrimFunc(symeval), tmp13920)
return


}, 0)

tmp13921 := MakeNative(func(__e *ControlFlow) {
Z5938 := __e.Get(1)
_ = Z5938
__e.Return(V5937)
return
}, 1)

__e.TailApply(try_1catch, tmp13919, tmp13921)
return


}, 1)

tmp13922 := Call(__e, ns2_1set, symunprofile, tmp13915)


_ = tmp13922

tmp13923 := MakeNative(func(__e *ControlFlow) {
V5939 := __e.Get(1)
_ = V5939
tmp13924 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dprofiled_d)
}
__typedArg0 := symshen_4_dprofiled_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symelement_2), V5939, tmp13924)
return


}, 1)

tmp13925 := Call(__e, ns2_1set, symshen_4profiled_2, tmp13923)


_ = tmp13925

tmp13926 := MakeNative(func(__e *ControlFlow) {
V5940 := __e.Get(1)
_ = V5940
V5941 := __e.Get(2)
_ = V5941
V5942 := __e.Get(3)
_ = V5942
tmp13927 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symrun, Nil)
}
__typedArg0 := symrun
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13928 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symget_1time, tmp13927)
}
__typedArg0 := symget_1time
__typedArg1 := tmp13927
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13929 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symrun, Nil)
}
__typedArg0 := symrun
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13930 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symget_1time, tmp13929)
}
__typedArg0 := symget_1time
__typedArg1 := tmp13929
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13931 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symStart, Nil)
}
__typedArg0 := symStart
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13932 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13930, tmp13931)
}
__typedArg0 := tmp13930
__typedArg1 := tmp13931
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13933 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1, tmp13932)
}
__typedArg0 := sym_1
__typedArg1 := tmp13932
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13934 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5940, Nil)
}
__typedArg0 := V5940
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13935 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4get_1profile, tmp13934)
}
__typedArg0 := symshen_4get_1profile
__typedArg1 := tmp13934
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13936 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symFinish, Nil)
}
__typedArg0 := symFinish
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13937 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13935, tmp13936)
}
__typedArg0 := tmp13935
__typedArg1 := tmp13936
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13938 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_7, tmp13937)
}
__typedArg0 := sym_7
__typedArg1 := tmp13937
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13939 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13938, Nil)
}
__typedArg0 := tmp13938
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13940 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5940, tmp13939)
}
__typedArg0 := V5940
__typedArg1 := tmp13939
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13941 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4put_1profile, tmp13940)
}
__typedArg0 := symshen_4put_1profile
__typedArg1 := tmp13940
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13942 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symResult, Nil)
}
__typedArg0 := symResult
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13943 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13941, tmp13942)
}
__typedArg0 := tmp13941
__typedArg1 := tmp13942
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13944 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symRecord, tmp13943)
}
__typedArg0 := symRecord
__typedArg1 := tmp13943
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13945 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp13944)
}
__typedArg0 := symlet
__typedArg1 := tmp13944
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13946 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13945, Nil)
}
__typedArg0 := tmp13945
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13947 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13933, tmp13946)
}
__typedArg0 := tmp13933
__typedArg1 := tmp13946
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13948 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symFinish, tmp13947)
}
__typedArg0 := symFinish
__typedArg1 := tmp13947
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13949 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp13948)
}
__typedArg0 := symlet
__typedArg1 := tmp13948
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13950 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13949, Nil)
}
__typedArg0 := tmp13949
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13951 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5942, tmp13950)
}
__typedArg0 := V5942
__typedArg1 := tmp13950
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13952 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symResult, tmp13951)
}
__typedArg0 := symResult
__typedArg1 := tmp13951
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13953 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp13952)
}
__typedArg0 := symlet
__typedArg1 := tmp13952
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13954 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13953, Nil)
}
__typedArg0 := tmp13953
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13955 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13928, tmp13954)
}
__typedArg0 := tmp13928
__typedArg1 := tmp13954
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13956 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symStart, tmp13955)
}
__typedArg0 := symStart
__typedArg1 := tmp13955
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp13956)
}
__typedArg0 := symlet
__typedArg1 := tmp13956
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 3)

tmp13957 := Call(__e, ns2_1set, symshen_4profile_1func, tmp13926)


_ = tmp13957

tmp13958 := MakeNative(func(__e *ControlFlow) {
V5943 := __e.Get(1)
_ = V5943
tmp13959 := Call(__e, PrimFunc(symshen_4get_1profile), V5943)


let__13533 := tmp13959
_ = let__13533

tmp13960 := Call(__e, PrimFunc(symshen_4put_1profile), V5943, MakeInteger(0))


let__13534 := tmp13960
_ = let__13534

__e.TailApply(PrimFunc(sym_8p), V5943, let__13533)
return


}, 1)

tmp13961 := Call(__e, ns2_1set, symprofile_1results, tmp13958)


_ = tmp13961

tmp13962 := MakeNative(func(__e *ControlFlow) {
V5946 := __e.Get(1)
_ = V5946
tmp13963 := MakeNative(func(__e *ControlFlow) {
tmp13964 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symget), V5946, symprofile, tmp13964)
return


}, 0)

tmp13965 := MakeNative(func(__e *ControlFlow) {
Z5947 := __e.Get(1)
_ = Z5947
__e.Return(MakeInteger(0))
return
}, 1)

__e.TailApply(try_1catch, tmp13963, tmp13965)
return


}, 1)

tmp13966 := Call(__e, ns2_1set, symshen_4get_1profile, tmp13962)


_ = tmp13966

tmp13967 := MakeNative(func(__e *ControlFlow) {
V5948 := __e.Get(1)
_ = V5948
V5949 := __e.Get(2)
_ = V5949
tmp13968 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symput), V5948, symprofile, V5949, tmp13968)
return


}, 2)

__e.TailApply(ns2_1set, symshen_4put_1profile, tmp13967)
return




}, 0)

