package main

import . "github.com/pyrex41/shen-go/kl"

var TrackMain = MakeNative(func(__e *ControlFlow) {
tmp13731 := MakeNative(func(__e *ControlFlow) {
V5384 := __e.Get(1)
_ = V5384
tmp13732 := Call(__e, PrimFunc(symshen_4app), V5384, MakeString(";\n"), symshen_4a)


tmp13733 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("partial function "))
__typedS1, __typedOK1 := TypedString(tmp13732)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("partial function ")
__typedArg1 := tmp13732
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp13734 := Call(__e, PrimFunc(symstoutput))


tmp13735 := Call(__e, PrimFunc(sympr), tmp13733, tmp13734)


_ = tmp13735

tmp13744 := Call(__e, PrimFunc(symshen_4tracked_2), V5384)


tmp13745 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp13744)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp13744
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres13739 Obj

if True == tmp13745 {
tmp13741 := Call(__e, PrimFunc(symshen_4app), V5384, MakeString("? "), symshen_4a)


tmp13743 := Call(__e, PrimFunc(symy_1or_1n_2), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("track "))
__typedS1, __typedOK1 := TypedString(tmp13741)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("track ")
__typedArg1 := tmp13741
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())


var ifres13740 Obj

if True == tmp13743 {
ifres13740 = True


} else {
ifres13740 = False


}

ifres13739 = ifres13740


} else {
ifres13739 = False


}

var ifres13736 Obj

if True == ifres13739 {
tmp13737 := Call(__e, PrimFunc(symps), V5384)


tmp13738 := Call(__e, PrimFunc(symshen_4track_1function), tmp13737)


ifres13736 = tmp13738


} else {
ifres13736 = symshen_4ok


}

_ = ifres13736

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("aborted"))
}
__typedArg0 := MakeString("aborted")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}, 1)

tmp13746 := Call(__e, ns2_1set, symshen_4f_1error, tmp13731)


_ = tmp13746

tmp13747 := MakeNative(func(__e *ControlFlow) {
V5385 := __e.Get(1)
_ = V5385
tmp13748 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dtracking_d)
}
__typedArg0 := symshen_4_dtracking_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symelement_2), V5385, tmp13748)
return


}, 1)

tmp13749 := Call(__e, ns2_1set, symshen_4tracked_2, tmp13747)


_ = tmp13749

tmp13750 := MakeNative(func(__e *ControlFlow) {
V5386 := __e.Get(1)
_ = V5386
tmp13751 := MakeNative(func(__e *ControlFlow) {
W5387 := __e.Get(1)
_ = W5387
__e.TailApply(PrimFunc(symshen_4track_1function), W5387)
return
}, 1)

tmp13752 := Call(__e, PrimFunc(symps), V5386)


__e.TailApply(tmp13751, tmp13752)
return


}, 1)

tmp13753 := Call(__e, ns2_1set, symtrack, tmp13750)


_ = tmp13753

tmp13754 := MakeNative(func(__e *ControlFlow) {
V5390 := __e.Get(1)
_ = V5390
tmp13811 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5390)
}
__typedArg0 := V5390
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13785 Obj

if True == tmp13811 {
tmp13809 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5390)
}
__typedArg0 := V5390
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13810 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefun, tmp13809)
}
__typedArg0 := symdefun
__typedArg1 := tmp13809
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13787 Obj

if True == tmp13810 {
tmp13807 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5390)
}
__typedArg0 := V5390
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13808 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13807)
}
__typedArg0 := tmp13807
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13789 Obj

if True == tmp13808 {
tmp13804 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5390)
}
__typedArg0 := V5390
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13805 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13804)
}
__typedArg0 := tmp13804
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13806 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13805)
}
__typedArg0 := tmp13805
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13791 Obj

if True == tmp13806 {
tmp13800 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5390)
}
__typedArg0 := V5390
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13801 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13800)
}
__typedArg0 := tmp13800
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13802 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13801)
}
__typedArg0 := tmp13801
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13803 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13802)
}
__typedArg0 := tmp13802
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13793 Obj

if True == tmp13803 {
tmp13795 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5390)
}
__typedArg0 := V5390
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13796 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13795)
}
__typedArg0 := tmp13795
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13797 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13796)
}
__typedArg0 := tmp13796
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13798 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13797)
}
__typedArg0 := tmp13797
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13799 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp13798)
}
__typedArg0 := Nil
__typedArg1 := tmp13798
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13794 Obj

if True == tmp13799 {
ifres13794 = True


} else {
ifres13794 = False


}

ifres13793 = ifres13794


} else {
ifres13793 = False


}

var ifres13792 Obj

if True == ifres13793 {
ifres13792 = True


} else {
ifres13792 = False


}

ifres13791 = ifres13792


} else {
ifres13791 = False


}

var ifres13790 Obj

if True == ifres13791 {
ifres13790 = True


} else {
ifres13790 = False


}

ifres13789 = ifres13790


} else {
ifres13789 = False


}

var ifres13788 Obj

if True == ifres13789 {
ifres13788 = True


} else {
ifres13788 = False


}

ifres13787 = ifres13788


} else {
ifres13787 = False


}

var ifres13786 Obj

if True == ifres13787 {
ifres13786 = True


} else {
ifres13786 = False


}

ifres13785 = ifres13786


} else {
ifres13785 = False


}

if True == ifres13785 {
tmp13755 := MakeNative(func(__e *ControlFlow) {
W5391 := __e.Get(1)
_ = W5391
tmp13756 := MakeNative(func(__e *ControlFlow) {
W5392 := __e.Get(1)
_ = W5392
tmp13757 := MakeNative(func(__e *ControlFlow) {
W5393 := __e.Get(1)
_ = W5393
tmp13758 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5390)
}
__typedArg0 := V5390
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13758)
}
__typedArg0 := tmp13758
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return


}, 1)

tmp13759 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5390)
}
__typedArg0 := V5390
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13760 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13759)
}
__typedArg0 := tmp13759
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13761 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dtracking_d)
}
__typedArg0 := symshen_4_dtracking_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp13762 := Call(__e, PrimFunc(symadjoin), tmp13760, tmp13761)


tmp13763 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dtracking_d, tmp13762)
}
__typedArg0 := symshen_4_dtracking_d
__typedArg1 := tmp13762
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

__e.TailApply(tmp13757, tmp13763)
return


}, 1)

tmp13764 := Call(__e, PrimFunc(symeval_1kl), W5391)


__e.TailApply(tmp13756, tmp13764)
return


}, 1)

tmp13765 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5390)
}
__typedArg0 := V5390
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13766 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13765)
}
__typedArg0 := tmp13765
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13767 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5390)
}
__typedArg0 := V5390
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13768 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13767)
}
__typedArg0 := tmp13767
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13769 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13768)
}
__typedArg0 := tmp13768
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13770 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5390)
}
__typedArg0 := V5390
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13771 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13770)
}
__typedArg0 := tmp13770
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13772 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5390)
}
__typedArg0 := V5390
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13773 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13772)
}
__typedArg0 := tmp13772
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13774 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13773)
}
__typedArg0 := tmp13773
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13775 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5390)
}
__typedArg0 := V5390
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13776 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13775)
}
__typedArg0 := tmp13775
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13777 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13776)
}
__typedArg0 := tmp13776
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13778 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13777)
}
__typedArg0 := tmp13777
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13779 := Call(__e, PrimFunc(symshen_4insert_1tracking_1code), tmp13771, tmp13774, tmp13778)


tmp13780 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13779, Nil)
}
__typedArg0 := tmp13779
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13781 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13769, tmp13780)
}
__typedArg0 := tmp13769
__typedArg1 := tmp13780
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13782 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13766, tmp13781)
}
__typedArg0 := tmp13766
__typedArg1 := tmp13781
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13783 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefun, tmp13782)
}
__typedArg0 := symdefun
__typedArg1 := tmp13782
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(tmp13755, tmp13783)
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

tmp13812 := Call(__e, ns2_1set, symshen_4track_1function, tmp13754)


_ = tmp13812

tmp13813 := MakeNative(func(__e *ControlFlow) {
V5394 := __e.Get(1)
_ = V5394
V5395 := __e.Get(2)
_ = V5395
V5396 := __e.Get(3)
_ = V5396
tmp13814 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4_dcall_d, Nil)
}
__typedArg0 := symshen_4_dcall_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13815 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvalue, tmp13814)
}
__typedArg0 := symvalue
__typedArg1 := tmp13814
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13816 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), Nil)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13817 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13815, tmp13816)
}
__typedArg0 := tmp13815
__typedArg1 := tmp13816
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13818 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_7, tmp13817)
}
__typedArg0 := sym_7
__typedArg1 := tmp13817
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13819 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13818, Nil)
}
__typedArg0 := tmp13818
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13820 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4_dcall_d, tmp13819)
}
__typedArg0 := symshen_4_dcall_d
__typedArg1 := tmp13819
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13821 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symset, tmp13820)
}
__typedArg0 := symset
__typedArg1 := tmp13820
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13822 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4_dcall_d, Nil)
}
__typedArg0 := symshen_4_dcall_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13823 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvalue, tmp13822)
}
__typedArg0 := symvalue
__typedArg1 := tmp13822
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13824 := Call(__e, PrimFunc(symshen_4prolog_1track), V5396, V5395)


tmp13825 := Call(__e, PrimFunc(symshen_4cons_1form), tmp13824)


tmp13826 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13825, Nil)
}
__typedArg0 := tmp13825
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13827 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5394, tmp13826)
}
__typedArg0 := V5394
__typedArg1 := tmp13826
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13828 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13823, tmp13827)
}
__typedArg0 := tmp13823
__typedArg1 := tmp13827
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13829 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4input_1track, tmp13828)
}
__typedArg0 := symshen_4input_1track
__typedArg1 := tmp13828
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13830 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4terpri_1or_1read_1char, Nil)
}
__typedArg0 := symshen_4terpri_1or_1read_1char
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13831 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4_dcall_d, Nil)
}
__typedArg0 := symshen_4_dcall_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13832 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvalue, tmp13831)
}
__typedArg0 := symvalue
__typedArg1 := tmp13831
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13833 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symResult, Nil)
}
__typedArg0 := symResult
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13834 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5394, tmp13833)
}
__typedArg0 := V5394
__typedArg1 := tmp13833
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13835 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13832, tmp13834)
}
__typedArg0 := tmp13832
__typedArg1 := tmp13834
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13836 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4output_1track, tmp13835)
}
__typedArg0 := symshen_4output_1track
__typedArg1 := tmp13835
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13837 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4_dcall_d, Nil)
}
__typedArg0 := symshen_4_dcall_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13838 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvalue, tmp13837)
}
__typedArg0 := symvalue
__typedArg1 := tmp13837
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13839 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), Nil)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13840 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13838, tmp13839)
}
__typedArg0 := tmp13838
__typedArg1 := tmp13839
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13841 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1, tmp13840)
}
__typedArg0 := sym_1
__typedArg1 := tmp13840
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13842 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13841, Nil)
}
__typedArg0 := tmp13841
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13843 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4_dcall_d, tmp13842)
}
__typedArg0 := symshen_4_dcall_d
__typedArg1 := tmp13842
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13844 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symset, tmp13843)
}
__typedArg0 := symset
__typedArg1 := tmp13843
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13845 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4terpri_1or_1read_1char, Nil)
}
__typedArg0 := symshen_4terpri_1or_1read_1char
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13846 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symResult, Nil)
}
__typedArg0 := symResult
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13847 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13845, tmp13846)
}
__typedArg0 := tmp13845
__typedArg1 := tmp13846
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13848 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdo, tmp13847)
}
__typedArg0 := symdo
__typedArg1 := tmp13847
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13849 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13848, Nil)
}
__typedArg0 := tmp13848
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13850 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13844, tmp13849)
}
__typedArg0 := tmp13844
__typedArg1 := tmp13849
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13851 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdo, tmp13850)
}
__typedArg0 := symdo
__typedArg1 := tmp13850
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13852 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13851, Nil)
}
__typedArg0 := tmp13851
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13853 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13836, tmp13852)
}
__typedArg0 := tmp13836
__typedArg1 := tmp13852
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13854 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdo, tmp13853)
}
__typedArg0 := symdo
__typedArg1 := tmp13853
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13855 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13854, Nil)
}
__typedArg0 := tmp13854
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13856 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5396, tmp13855)
}
__typedArg0 := V5396
__typedArg1 := tmp13855
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13857 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symResult, tmp13856)
}
__typedArg0 := symResult
__typedArg1 := tmp13856
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13858 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp13857)
}
__typedArg0 := symlet
__typedArg1 := tmp13857
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13859 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13858, Nil)
}
__typedArg0 := tmp13858
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13860 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13830, tmp13859)
}
__typedArg0 := tmp13830
__typedArg1 := tmp13859
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13861 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdo, tmp13860)
}
__typedArg0 := symdo
__typedArg1 := tmp13860
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13862 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13861, Nil)
}
__typedArg0 := tmp13861
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13863 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13829, tmp13862)
}
__typedArg0 := tmp13829
__typedArg1 := tmp13862
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13864 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdo, tmp13863)
}
__typedArg0 := symdo
__typedArg1 := tmp13863
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

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
return PrimCons(tmp13821, tmp13865)
}
__typedArg0 := tmp13821
__typedArg1 := tmp13865
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdo, tmp13866)
}
__typedArg0 := symdo
__typedArg1 := tmp13866
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 3)

tmp13867 := Call(__e, ns2_1set, symshen_4insert_1tracking_1code, tmp13813)


_ = tmp13867

tmp13868 := MakeNative(func(__e *ControlFlow) {
V5397 := __e.Get(1)
_ = V5397
V5398 := __e.Get(2)
_ = V5398
tmp13871 := Call(__e, PrimFunc(symoccurrences), symshen_4incinfs, V5397)


tmp13872 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp13871, MakeNumber(0))
}
__typedArg0 := tmp13871
__typedArg1 := MakeNumber(0)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13872 {
__e.Return(V5398)
return
} else {
tmp13869 := Call(__e, PrimFunc(symshen_4vector_1parameter), V5398)


__e.TailApply(PrimFunc(symshen_4vector_1dereference), V5398, tmp13869)
return


}


}, 2)

tmp13873 := Call(__e, ns2_1set, symshen_4prolog_1track, tmp13868)


_ = tmp13873

tmp13874 := MakeNative(func(__e *ControlFlow) {
V5401 := __e.Get(1)
_ = V5401
tmp13903 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V5401)
}
__typedArg0 := Nil
__typedArg1 := V5401
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13903 {
__e.Return(Nil)
return
} else {
tmp13901 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5401)
}
__typedArg0 := V5401
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13879 Obj

if True == tmp13901 {
tmp13899 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5401)
}
__typedArg0 := V5401
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13900 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13899)
}
__typedArg0 := tmp13899
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13881 Obj

if True == tmp13900 {
tmp13896 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5401)
}
__typedArg0 := V5401
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13897 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13896)
}
__typedArg0 := tmp13896
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13898 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13897)
}
__typedArg0 := tmp13897
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13883 Obj

if True == tmp13898 {
tmp13892 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5401)
}
__typedArg0 := V5401
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13893 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13892)
}
__typedArg0 := tmp13892
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13894 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13893)
}
__typedArg0 := tmp13893
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13895 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13894)
}
__typedArg0 := tmp13894
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13885 Obj

if True == tmp13895 {
tmp13887 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5401)
}
__typedArg0 := V5401
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13888 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13887)
}
__typedArg0 := tmp13887
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13889 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13888)
}
__typedArg0 := tmp13888
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp13890)
}
__typedArg0 := Nil
__typedArg1 := tmp13890
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13886 Obj

if True == tmp13891 {
ifres13886 = True


} else {
ifres13886 = False


}

ifres13885 = ifres13886


} else {
ifres13885 = False


}

var ifres13884 Obj

if True == ifres13885 {
ifres13884 = True


} else {
ifres13884 = False


}

ifres13883 = ifres13884


} else {
ifres13883 = False


}

var ifres13882 Obj

if True == ifres13883 {
ifres13882 = True


} else {
ifres13882 = False


}

ifres13881 = ifres13882


} else {
ifres13881 = False


}

var ifres13880 Obj

if True == ifres13881 {
ifres13880 = True


} else {
ifres13880 = False


}

ifres13879 = ifres13880


} else {
ifres13879 = False


}

if True == ifres13879 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5401)
}
__typedArg0 := V5401
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return
} else {
tmp13877 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5401)
}
__typedArg0 := V5401
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp13877 {
tmp13875 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5401)
}
__typedArg0 := V5401
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4vector_1parameter), tmp13875)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("partial function shen.vector-parameter"))
}
__typedArg0 := MakeString("partial function shen.vector-parameter")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}


}, 1)

tmp13904 := Call(__e, ns2_1set, symshen_4vector_1parameter, tmp13874)


_ = tmp13904

tmp13905 := MakeNative(func(__e *ControlFlow) {
V5404 := __e.Get(1)
_ = V5404
V5405 := __e.Get(2)
_ = V5405
tmp13939 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V5405)
}
__typedArg0 := Nil
__typedArg1 := V5405
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13939 {
__e.Return(V5404)
return
} else {
tmp13937 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5404)
}
__typedArg0 := V5404
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13915 Obj

if True == tmp13937 {
tmp13935 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5404)
}
__typedArg0 := V5404
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13936 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13935)
}
__typedArg0 := tmp13935
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13917 Obj

if True == tmp13936 {
tmp13932 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5404)
}
__typedArg0 := V5404
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13933 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13932)
}
__typedArg0 := tmp13932
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13934 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13933)
}
__typedArg0 := tmp13933
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13919 Obj

if True == tmp13934 {
tmp13928 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5404)
}
__typedArg0 := V5404
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13929 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13928)
}
__typedArg0 := tmp13928
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13930 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13929)
}
__typedArg0 := tmp13929
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13931 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13930)
}
__typedArg0 := tmp13930
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13921 Obj

if True == tmp13931 {
tmp13923 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5404)
}
__typedArg0 := V5404
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13924)
}
__typedArg0 := tmp13924
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13926 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13925)
}
__typedArg0 := tmp13925
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13927 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp13926)
}
__typedArg0 := Nil
__typedArg1 := tmp13926
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13922 Obj

if True == tmp13927 {
ifres13922 = True


} else {
ifres13922 = False


}

ifres13921 = ifres13922


} else {
ifres13921 = False


}

var ifres13920 Obj

if True == ifres13921 {
ifres13920 = True


} else {
ifres13920 = False


}

ifres13919 = ifres13920


} else {
ifres13919 = False


}

var ifres13918 Obj

if True == ifres13919 {
ifres13918 = True


} else {
ifres13918 = False


}

ifres13917 = ifres13918


} else {
ifres13917 = False


}

var ifres13916 Obj

if True == ifres13917 {
ifres13916 = True


} else {
ifres13916 = False


}

ifres13915 = ifres13916


} else {
ifres13915 = False


}

if True == ifres13915 {
__e.Return(V5404)
return
} else {
tmp13913 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5404)
}
__typedArg0 := V5404
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp13913 {
tmp13906 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5404)
}
__typedArg0 := V5404
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13907 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5405, Nil)
}
__typedArg0 := V5405
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13908 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13906, tmp13907)
}
__typedArg0 := tmp13906
__typedArg1 := tmp13907
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13909 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4deref, tmp13908)
}
__typedArg0 := symshen_4deref
__typedArg1 := tmp13908
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13910 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5404)
}
__typedArg0 := V5404
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13911 := Call(__e, PrimFunc(symshen_4vector_1dereference), tmp13910, V5405)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13909, tmp13911)
}
__typedArg0 := tmp13909
__typedArg1 := tmp13911
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("partial function shen.vector-dereference"))
}
__typedArg0 := MakeString("partial function shen.vector-dereference")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}


}, 2)

tmp13940 := Call(__e, ns2_1set, symshen_4vector_1dereference, tmp13905)


_ = tmp13940

tmp13941 := MakeNative(func(__e *ControlFlow) {
V5408 := __e.Get(1)
_ = V5408
tmp13945 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_7, V5408)
}
__typedArg0 := sym_7
__typedArg1 := V5408
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13945 {
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
tmp13943 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_1, V5408)
}
__typedArg0 := sym_1
__typedArg1 := V5408
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13943 {
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

tmp13946 := Call(__e, ns2_1set, symstep, tmp13941)


_ = tmp13946

tmp13947 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dstep_d)
}
__typedArg0 := symshen_4_dstep_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp13948 := Call(__e, ns2_1set, symstep_2, tmp13947)


_ = tmp13948

tmp13949 := MakeNative(func(__e *ControlFlow) {
V5411 := __e.Get(1)
_ = V5411
tmp13953 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_7, V5411)
}
__typedArg0 := sym_7
__typedArg1 := V5411
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13953 {
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
tmp13951 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_1, V5411)
}
__typedArg0 := sym_1
__typedArg1 := V5411
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13951 {
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

tmp13954 := Call(__e, ns2_1set, symspy, tmp13949)


_ = tmp13954

tmp13955 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dspy_d)
}
__typedArg0 := symshen_4_dspy_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp13956 := Call(__e, ns2_1set, symspy_2, tmp13955)


_ = tmp13956

tmp13957 := MakeNative(func(__e *ControlFlow) {
tmp13961 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dstep_d)
}
__typedArg0 := symshen_4_dstep_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

if True == tmp13961 {
tmp13958 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dstinput_d)
}
__typedArg0 := sym_dstinput_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp13959 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symread_1byte) {
return PrimReadByte(tmp13958)
}
__typedArg0 := tmp13958
return Call(__e, PrimFunc(symread_1byte), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4check_1byte), tmp13959)
return


} else {
__e.TailApply(PrimFunc(symnl), MakeNumber(1))
return
}


}, 0)

tmp13962 := Call(__e, ns2_1set, symshen_4terpri_1or_1read_1char, tmp13957)


_ = tmp13962

tmp13963 := MakeNative(func(__e *ControlFlow) {
V5414 := __e.Get(1)
_ = V5414
tmp13965 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeNumber(94), V5414)
}
__typedArg0 := MakeNumber(94)
__typedArg1 := V5414
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13965 {
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

tmp13966 := Call(__e, ns2_1set, symshen_4check_1byte, tmp13963)


_ = tmp13966

tmp13967 := MakeNative(func(__e *ControlFlow) {
V5415 := __e.Get(1)
_ = V5415
V5416 := __e.Get(2)
_ = V5416
V5417 := __e.Get(3)
_ = V5417
tmp13968 := Call(__e, PrimFunc(symshen_4spaces), V5415)


tmp13969 := Call(__e, PrimFunc(symshen_4spaces), V5415)


tmp13970 := Call(__e, PrimFunc(symshen_4app), tmp13969, MakeString(""), symshen_4a)


tmp13972 := Call(__e, PrimFunc(symshen_4app), V5416, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(" \n"))
__typedS1, __typedOK1 := TypedString(tmp13970)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(" \n")
__typedArg1 := tmp13970
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


tmp13974 := Call(__e, PrimFunc(symshen_4app), V5415, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("> Inputs to "))
__typedS1, __typedOK1 := TypedString(tmp13972)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("> Inputs to ")
__typedArg1 := tmp13972
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


tmp13976 := Call(__e, PrimFunc(symshen_4app), tmp13968, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("<"))
__typedS1, __typedOK1 := TypedString(tmp13974)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("<")
__typedArg1 := tmp13974
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


tmp13977 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("\n"))
__typedS1, __typedOK1 := TypedString(tmp13976)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("\n")
__typedArg1 := tmp13976
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp13978 := Call(__e, PrimFunc(symstoutput))


tmp13979 := Call(__e, PrimFunc(sympr), tmp13977, tmp13978)


_ = tmp13979

__e.TailApply(PrimFunc(symshen_4recursively_1print), V5417)
return


}, 3)

tmp13980 := Call(__e, ns2_1set, symshen_4input_1track, tmp13967)


_ = tmp13980

tmp13981 := MakeNative(func(__e *ControlFlow) {
V5420 := __e.Get(1)
_ = V5420
tmp13991 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V5420)
}
__typedArg0 := Nil
__typedArg1 := V5420
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13991 {
tmp13982 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), MakeString(" ==>"), tmp13982)
return


} else {
tmp13989 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5420)
}
__typedArg0 := V5420
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp13989 {
tmp13983 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5420)
}
__typedArg0 := V5420
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13984 := Call(__e, PrimFunc(symprint), tmp13983)


_ = tmp13984

tmp13985 := Call(__e, PrimFunc(symstoutput))


tmp13986 := Call(__e, PrimFunc(sympr), MakeString(", "), tmp13985)


_ = tmp13986

tmp13987 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5420)
}
__typedArg0 := V5420
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4recursively_1print), tmp13987)
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

tmp13992 := Call(__e, ns2_1set, symshen_4recursively_1print, tmp13981)


_ = tmp13992

tmp13993 := MakeNative(func(__e *ControlFlow) {
V5421 := __e.Get(1)
_ = V5421
tmp13997 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeNumber(0), V5421)
}
__typedArg0 := MakeNumber(0)
__typedArg1 := V5421
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13997 {
__e.Return(MakeString(""))
return
} else {
tmp13995 := Call(__e, PrimFunc(symshen_4spaces), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(V5421)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := V5421
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})())


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(" "))
__typedS1, __typedOK1 := TypedString(tmp13995)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(" ")
__typedArg1 := tmp13995
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
return


}


}, 1)

tmp13998 := Call(__e, ns2_1set, symshen_4spaces, tmp13993)


_ = tmp13998

tmp13999 := MakeNative(func(__e *ControlFlow) {
V5422 := __e.Get(1)
_ = V5422
V5423 := __e.Get(2)
_ = V5423
V5424 := __e.Get(3)
_ = V5424
tmp14000 := Call(__e, PrimFunc(symshen_4spaces), V5422)


tmp14001 := Call(__e, PrimFunc(symshen_4spaces), V5422)


tmp14002 := Call(__e, PrimFunc(symshen_4app), V5424, MakeString(""), symshen_4s)


tmp14004 := Call(__e, PrimFunc(symshen_4app), tmp14001, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("==> "))
__typedS1, __typedOK1 := TypedString(tmp14002)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("==> ")
__typedArg1 := tmp14002
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


tmp14006 := Call(__e, PrimFunc(symshen_4app), V5423, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(" \n"))
__typedS1, __typedOK1 := TypedString(tmp14004)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(" \n")
__typedArg1 := tmp14004
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


tmp14008 := Call(__e, PrimFunc(symshen_4app), V5422, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("> Output from "))
__typedS1, __typedOK1 := TypedString(tmp14006)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("> Output from ")
__typedArg1 := tmp14006
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


tmp14010 := Call(__e, PrimFunc(symshen_4app), tmp14000, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("<"))
__typedS1, __typedOK1 := TypedString(tmp14008)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("<")
__typedArg1 := tmp14008
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


tmp14011 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("\n"))
__typedS1, __typedOK1 := TypedString(tmp14010)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("\n")
__typedArg1 := tmp14010
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp14012 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp14011, tmp14012)
return


}, 3)

tmp14013 := Call(__e, ns2_1set, symshen_4output_1track, tmp13999)


_ = tmp14013

tmp14014 := MakeNative(func(__e *ControlFlow) {
V5425 := __e.Get(1)
_ = V5425
tmp14015 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dtracking_d)
}
__typedArg0 := symshen_4_dtracking_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp14016 := Call(__e, PrimFunc(symremove), V5425, tmp14015)


tmp14017 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dtracking_d, tmp14016)
}
__typedArg0 := symshen_4_dtracking_d
__typedArg1 := tmp14016
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp14017

tmp14018 := MakeNative(func(__e *ControlFlow) {
tmp14019 := Call(__e, PrimFunc(symps), V5425)


__e.TailApply(PrimFunc(symeval), tmp14019)
return


}, 0)

tmp14020 := MakeNative(func(__e *ControlFlow) {
Z5426 := __e.Get(1)
_ = Z5426
__e.Return(V5425)
return
}, 1)

tmp14021 := Call(__e, try_1catch, tmp14018, tmp14020)


_ = tmp14021

__e.Return(V5425)
return


}, 1)

tmp14022 := Call(__e, ns2_1set, symuntrack, tmp14014)


_ = tmp14022

tmp14023 := MakeNative(func(__e *ControlFlow) {
V5427 := __e.Get(1)
_ = V5427
V5428 := __e.Get(2)
_ = V5428
__e.TailApply(PrimFunc(symshen_4remove_1h), V5427, V5428, Nil)
return
}, 2)

tmp14024 := Call(__e, ns2_1set, symremove, tmp14023)


_ = tmp14024

tmp14025 := MakeNative(func(__e *ControlFlow) {
V5438 := __e.Get(1)
_ = V5438
V5439 := __e.Get(2)
_ = V5439
V5440 := __e.Get(3)
_ = V5440
tmp14040 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V5439)
}
__typedArg0 := Nil
__typedArg1 := V5439
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp14040 {
__e.TailApply(PrimFunc(symreverse), V5440)
return
} else {
tmp14038 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5439)
}
__typedArg0 := V5439
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14034 Obj

if True == tmp14038 {
tmp14036 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5439)
}
__typedArg0 := V5439
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14037 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V5438, tmp14036)
}
__typedArg0 := V5438
__typedArg1 := tmp14036
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14035 Obj

if True == tmp14037 {
ifres14035 = True


} else {
ifres14035 = False


}

ifres14034 = ifres14035


} else {
ifres14034 = False


}

if True == ifres14034 {
tmp14026 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5439)
}
__typedArg0 := V5439
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14027 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5439)
}
__typedArg0 := V5439
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4remove_1h), tmp14026, tmp14027, V5440)
return


} else {
tmp14032 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5439)
}
__typedArg0 := V5439
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp14032 {
tmp14028 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5439)
}
__typedArg0 := V5439
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14029 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5439)
}
__typedArg0 := V5439
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14030 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14029, V5440)
}
__typedArg0 := tmp14029
__typedArg1 := V5440
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4remove_1h), V5438, tmp14028, tmp14030)
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

tmp14041 := Call(__e, ns2_1set, symshen_4remove_1h, tmp14025)


_ = tmp14041

tmp14042 := MakeNative(func(__e *ControlFlow) {
V5441 := __e.Get(1)
_ = V5441
tmp14043 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dprofiled_d)
}
__typedArg0 := symshen_4_dprofiled_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp14044 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5441, tmp14043)
}
__typedArg0 := V5441
__typedArg1 := tmp14043
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14045 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dprofiled_d, tmp14044)
}
__typedArg0 := symshen_4_dprofiled_d
__typedArg1 := tmp14044
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp14045

tmp14046 := Call(__e, PrimFunc(symps), V5441)


__e.TailApply(PrimFunc(symshen_4profile_1help), tmp14046)
return


}, 1)

tmp14047 := Call(__e, ns2_1set, symprofile, tmp14042)


_ = tmp14047

tmp14048 := MakeNative(func(__e *ControlFlow) {
V5444 := __e.Get(1)
_ = V5444
tmp14118 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5444)
}
__typedArg0 := V5444
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14092 Obj

if True == tmp14118 {
tmp14116 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5444)
}
__typedArg0 := V5444
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14117 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefun, tmp14116)
}
__typedArg0 := symdefun
__typedArg1 := tmp14116
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14094 Obj

if True == tmp14117 {
tmp14114 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5444)
}
__typedArg0 := V5444
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14115 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14114)
}
__typedArg0 := tmp14114
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14096 Obj

if True == tmp14115 {
tmp14111 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5444)
}
__typedArg0 := V5444
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14112 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14111)
}
__typedArg0 := tmp14111
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14113 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14112)
}
__typedArg0 := tmp14112
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14098 Obj

if True == tmp14113 {
tmp14107 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5444)
}
__typedArg0 := V5444
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14108 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14107)
}
__typedArg0 := tmp14107
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14109 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14108)
}
__typedArg0 := tmp14108
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14110 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14109)
}
__typedArg0 := tmp14109
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14100 Obj

if True == tmp14110 {
tmp14102 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5444)
}
__typedArg0 := V5444
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14103 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14102)
}
__typedArg0 := tmp14102
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14104 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14103)
}
__typedArg0 := tmp14103
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14105 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14104)
}
__typedArg0 := tmp14104
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14106 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp14105)
}
__typedArg0 := Nil
__typedArg1 := tmp14105
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14101 Obj

if True == tmp14106 {
ifres14101 = True


} else {
ifres14101 = False


}

ifres14100 = ifres14101


} else {
ifres14100 = False


}

var ifres14099 Obj

if True == ifres14100 {
ifres14099 = True


} else {
ifres14099 = False


}

ifres14098 = ifres14099


} else {
ifres14098 = False


}

var ifres14097 Obj

if True == ifres14098 {
ifres14097 = True


} else {
ifres14097 = False


}

ifres14096 = ifres14097


} else {
ifres14096 = False


}

var ifres14095 Obj

if True == ifres14096 {
ifres14095 = True


} else {
ifres14095 = False


}

ifres14094 = ifres14095


} else {
ifres14094 = False


}

var ifres14093 Obj

if True == ifres14094 {
ifres14093 = True


} else {
ifres14093 = False


}

ifres14092 = ifres14093


} else {
ifres14092 = False


}

if True == ifres14092 {
tmp14049 := MakeNative(func(__e *ControlFlow) {
W5445 := __e.Get(1)
_ = W5445
tmp14050 := MakeNative(func(__e *ControlFlow) {
W5446 := __e.Get(1)
_ = W5446
tmp14051 := MakeNative(func(__e *ControlFlow) {
W5447 := __e.Get(1)
_ = W5447
tmp14052 := MakeNative(func(__e *ControlFlow) {
W5448 := __e.Get(1)
_ = W5448
tmp14053 := MakeNative(func(__e *ControlFlow) {
W5449 := __e.Get(1)
_ = W5449
tmp14054 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5444)
}
__typedArg0 := V5444
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp14054)
}
__typedArg0 := tmp14054
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return


}, 1)

tmp14055 := Call(__e, PrimFunc(symeval_1kl), W5447)


__e.TailApply(tmp14053, tmp14055)
return


}, 1)

tmp14056 := Call(__e, PrimFunc(symeval_1kl), W5446)


__e.TailApply(tmp14052, tmp14056)
return


}, 1)

tmp14057 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5444)
}
__typedArg0 := V5444
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14058 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14057)
}
__typedArg0 := tmp14057
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14059 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp14058)
}
__typedArg0 := tmp14058
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14060 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5444)
}
__typedArg0 := V5444
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14061 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp14060)
}
__typedArg0 := tmp14060
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14062 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5444)
}
__typedArg0 := V5444
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14063 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14062)
}
__typedArg0 := tmp14062
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14064 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14063)
}
__typedArg0 := tmp14063
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14065 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp14064)
}
__typedArg0 := tmp14064
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14066 := Call(__e, PrimFunc(symsubst), W5445, tmp14061, tmp14065)


tmp14067 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14066, Nil)
}
__typedArg0 := tmp14066
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14068 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14059, tmp14067)
}
__typedArg0 := tmp14059
__typedArg1 := tmp14067
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14069 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W5445, tmp14068)
}
__typedArg0 := W5445
__typedArg1 := tmp14068
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14070 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefun, tmp14069)
}
__typedArg0 := symdefun
__typedArg1 := tmp14069
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(tmp14051, tmp14070)
return


}, 1)

tmp14071 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5444)
}
__typedArg0 := V5444
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14072 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp14071)
}
__typedArg0 := tmp14071
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14073 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5444)
}
__typedArg0 := V5444
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14074 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14073)
}
__typedArg0 := tmp14073
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14075 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp14074)
}
__typedArg0 := tmp14074
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14076 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5444)
}
__typedArg0 := V5444
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14077 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp14076)
}
__typedArg0 := tmp14076
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14078 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5444)
}
__typedArg0 := V5444
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14079 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14078)
}
__typedArg0 := tmp14078
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14080 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp14079)
}
__typedArg0 := tmp14079
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14081 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5444)
}
__typedArg0 := V5444
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14082 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14081)
}
__typedArg0 := tmp14081
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14083 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp14082)
}
__typedArg0 := tmp14082
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14084 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W5445, tmp14083)
}
__typedArg0 := W5445
__typedArg1 := tmp14083
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14085 := Call(__e, PrimFunc(symshen_4profile_1func), tmp14077, tmp14080, tmp14084)


tmp14086 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14085, Nil)
}
__typedArg0 := tmp14085
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14087 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14075, tmp14086)
}
__typedArg0 := tmp14075
__typedArg1 := tmp14086
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14088 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14072, tmp14087)
}
__typedArg0 := tmp14072
__typedArg1 := tmp14087
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14089 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefun, tmp14088)
}
__typedArg0 := symdefun
__typedArg1 := tmp14088
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(tmp14050, tmp14089)
return


}, 1)

tmp14090 := Call(__e, PrimFunc(symgensym), symshen_4f)


__e.TailApply(tmp14049, tmp14090)
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

tmp14119 := Call(__e, ns2_1set, symshen_4profile_1help, tmp14048)


_ = tmp14119

tmp14120 := MakeNative(func(__e *ControlFlow) {
V5450 := __e.Get(1)
_ = V5450
tmp14121 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dprofiled_d)
}
__typedArg0 := symshen_4_dprofiled_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp14122 := Call(__e, PrimFunc(symremove), V5450, tmp14121)


tmp14123 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dprofiled_d, tmp14122)
}
__typedArg0 := symshen_4_dprofiled_d
__typedArg1 := tmp14122
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp14123

tmp14124 := MakeNative(func(__e *ControlFlow) {
tmp14125 := Call(__e, PrimFunc(symps), V5450)


__e.TailApply(PrimFunc(symeval), tmp14125)
return


}, 0)

tmp14126 := MakeNative(func(__e *ControlFlow) {
Z5451 := __e.Get(1)
_ = Z5451
__e.Return(V5450)
return
}, 1)

__e.TailApply(try_1catch, tmp14124, tmp14126)
return


}, 1)

tmp14127 := Call(__e, ns2_1set, symunprofile, tmp14120)


_ = tmp14127

tmp14128 := MakeNative(func(__e *ControlFlow) {
V5452 := __e.Get(1)
_ = V5452
tmp14129 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dprofiled_d)
}
__typedArg0 := symshen_4_dprofiled_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symelement_2), V5452, tmp14129)
return


}, 1)

tmp14130 := Call(__e, ns2_1set, symshen_4profiled_2, tmp14128)


_ = tmp14130

tmp14131 := MakeNative(func(__e *ControlFlow) {
V5453 := __e.Get(1)
_ = V5453
V5454 := __e.Get(2)
_ = V5454
V5455 := __e.Get(3)
_ = V5455
tmp14132 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symrun, Nil)
}
__typedArg0 := symrun
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14133 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symget_1time, tmp14132)
}
__typedArg0 := symget_1time
__typedArg1 := tmp14132
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14134 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symrun, Nil)
}
__typedArg0 := symrun
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14135 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symget_1time, tmp14134)
}
__typedArg0 := symget_1time
__typedArg1 := tmp14134
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14136 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symStart, Nil)
}
__typedArg0 := symStart
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14137 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14135, tmp14136)
}
__typedArg0 := tmp14135
__typedArg1 := tmp14136
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14138 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1, tmp14137)
}
__typedArg0 := sym_1
__typedArg1 := tmp14137
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14139 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5453, Nil)
}
__typedArg0 := V5453
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14140 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4get_1profile, tmp14139)
}
__typedArg0 := symshen_4get_1profile
__typedArg1 := tmp14139
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14141 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symFinish, Nil)
}
__typedArg0 := symFinish
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14142 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14140, tmp14141)
}
__typedArg0 := tmp14140
__typedArg1 := tmp14141
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14143 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_7, tmp14142)
}
__typedArg0 := sym_7
__typedArg1 := tmp14142
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14144 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14143, Nil)
}
__typedArg0 := tmp14143
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14145 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5453, tmp14144)
}
__typedArg0 := V5453
__typedArg1 := tmp14144
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14146 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4put_1profile, tmp14145)
}
__typedArg0 := symshen_4put_1profile
__typedArg1 := tmp14145
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14147 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symResult, Nil)
}
__typedArg0 := symResult
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14148 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14146, tmp14147)
}
__typedArg0 := tmp14146
__typedArg1 := tmp14147
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14149 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symRecord, tmp14148)
}
__typedArg0 := symRecord
__typedArg1 := tmp14148
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14150 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp14149)
}
__typedArg0 := symlet
__typedArg1 := tmp14149
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14151 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14150, Nil)
}
__typedArg0 := tmp14150
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14152 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14138, tmp14151)
}
__typedArg0 := tmp14138
__typedArg1 := tmp14151
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14153 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symFinish, tmp14152)
}
__typedArg0 := symFinish
__typedArg1 := tmp14152
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14154 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp14153)
}
__typedArg0 := symlet
__typedArg1 := tmp14153
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14155 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14154, Nil)
}
__typedArg0 := tmp14154
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14156 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5455, tmp14155)
}
__typedArg0 := V5455
__typedArg1 := tmp14155
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14157 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symResult, tmp14156)
}
__typedArg0 := symResult
__typedArg1 := tmp14156
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14158 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp14157)
}
__typedArg0 := symlet
__typedArg1 := tmp14157
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14159 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14158, Nil)
}
__typedArg0 := tmp14158
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14160 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14133, tmp14159)
}
__typedArg0 := tmp14133
__typedArg1 := tmp14159
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14161 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symStart, tmp14160)
}
__typedArg0 := symStart
__typedArg1 := tmp14160
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp14161)
}
__typedArg0 := symlet
__typedArg1 := tmp14161
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 3)

tmp14162 := Call(__e, ns2_1set, symshen_4profile_1func, tmp14131)


_ = tmp14162

tmp14163 := MakeNative(func(__e *ControlFlow) {
V5456 := __e.Get(1)
_ = V5456
tmp14164 := MakeNative(func(__e *ControlFlow) {
W5457 := __e.Get(1)
_ = W5457
tmp14165 := MakeNative(func(__e *ControlFlow) {
W5458 := __e.Get(1)
_ = W5458
__e.TailApply(PrimFunc(sym_8p), V5456, W5457)
return
}, 1)

tmp14166 := Call(__e, PrimFunc(symshen_4put_1profile), V5456, MakeNumber(0))


__e.TailApply(tmp14165, tmp14166)
return


}, 1)

tmp14167 := Call(__e, PrimFunc(symshen_4get_1profile), V5456)


__e.TailApply(tmp14164, tmp14167)
return


}, 1)

tmp14168 := Call(__e, ns2_1set, symprofile_1results, tmp14163)


_ = tmp14168

tmp14169 := MakeNative(func(__e *ControlFlow) {
V5459 := __e.Get(1)
_ = V5459
tmp14170 := MakeNative(func(__e *ControlFlow) {
tmp14171 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symget), V5459, symprofile, tmp14171)
return


}, 0)

tmp14172 := MakeNative(func(__e *ControlFlow) {
Z5460 := __e.Get(1)
_ = Z5460
__e.Return(MakeNumber(0))
return
}, 1)

__e.TailApply(try_1catch, tmp14170, tmp14172)
return


}, 1)

tmp14173 := Call(__e, ns2_1set, symshen_4get_1profile, tmp14169)


_ = tmp14173

tmp14174 := MakeNative(func(__e *ControlFlow) {
V5461 := __e.Get(1)
_ = V5461
V5462 := __e.Get(2)
_ = V5462
tmp14175 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symput), V5461, symprofile, V5462, tmp14175)
return


}, 2)

__e.TailApply(ns2_1set, symshen_4put_1profile, tmp14174)
return




}, 0)

