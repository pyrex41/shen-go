package main

import . "github.com/pyrex41/shen-go/kl"

var LoadMain = MakeNative(func(__e *ControlFlow) {
tmp9761 := MakeNative(func(__e *ControlFlow) {
V907 := __e.Get(1)
_ = V907
tmp9762 := MakeNative(func(__e *ControlFlow) {
W908 := __e.Get(1)
_ = W908
tmp9763 := MakeNative(func(__e *ControlFlow) {
W909 := __e.Get(1)
_ = W909
tmp9764 := MakeNative(func(__e *ControlFlow) {
W915 := __e.Get(1)
_ = W915
__e.Return(symloaded)
return
}, 1)

var ifres9765 Obj

if True == W908 {
tmp9766 := Call(__e, PrimFunc(syminferences))


tmp9767 := Call(__e, PrimFunc(symshen_4app), tmp9766, MakeString(" inferences\n"), symshen_4a)


tmp9768 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("\ntypechecked in "))
__typedS1, __typedOK1 := TypedString(tmp9767)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("\ntypechecked in ")
__typedArg1 := tmp9767
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp9769 := Call(__e, PrimFunc(symstoutput))


tmp9770 := Call(__e, PrimFunc(sympr), tmp9768, tmp9769)


ifres9765 = tmp9770


} else {
ifres9765 = symshen_4skip


}

__e.TailApply(tmp9764, ifres9765)
return


}, 1)

tmp9771 := MakeNative(func(__e *ControlFlow) {
W910 := __e.Get(1)
_ = W910
tmp9772 := MakeNative(func(__e *ControlFlow) {
W911 := __e.Get(1)
_ = W911
tmp9773 := MakeNative(func(__e *ControlFlow) {
W912 := __e.Get(1)
_ = W912
tmp9774 := MakeNative(func(__e *ControlFlow) {
W913 := __e.Get(1)
_ = W913
tmp9775 := MakeNative(func(__e *ControlFlow) {
W914 := __e.Get(1)
_ = W914
__e.Return(W911)
return
}, 1)

tmp9776 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(W913)
}
__typedArg0 := W913
return Call(__e, PrimFunc(symstr), __typedArg0)
})()

tmp9778 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("\nrun time: "))
__typedS1, __typedOK1 := TypedString(tmp9776)
__typedS2, __typedOK2 := TypedString(MakeString(" secs\n"))
if __typedOK0 && __typedOK1 && __typedOK2 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + (__typedS1 + __typedS2)))
}}
__typedArg0 := MakeString("\nrun time: ")
__typedArg1 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(tmp9776)
__typedS1, __typedOK1 := TypedString(MakeString(" secs\n"))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := tmp9776
__typedArg1 := MakeString(" secs\n")
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp9779 := Call(__e, PrimFunc(symstoutput))


tmp9780 := Call(__e, PrimFunc(sympr), tmp9778, tmp9779)


__e.TailApply(tmp9775, tmp9780)
return


}, 1)

__e.TailApply(tmp9774, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(W912)
__typedN1, __typedOK1 := TypedFloat64(W910)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := W912
__typedArg1 := W910
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp9782 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symget_1time) {
return PrimGetTime(symrun)
}
__typedArg0 := symrun
return Call(__e, PrimFunc(symget_1time), __typedArg0)
})()

__e.TailApply(tmp9773, tmp9782)
return


}, 1)

tmp9783 := Call(__e, PrimFunc(symread_1file), V907)


tmp9784 := Call(__e, PrimFunc(symshen_4load_1help), W908, tmp9783)


__e.TailApply(tmp9772, tmp9784)
return


}, 1)

tmp9785 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symget_1time) {
return PrimGetTime(symrun)
}
__typedArg0 := symrun
return Call(__e, PrimFunc(symget_1time), __typedArg0)
})()

tmp9786 := Call(__e, tmp9771, tmp9785)


__e.TailApply(tmp9763, tmp9786)
return


}, 1)

tmp9787 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dtc_d)
}
__typedArg0 := symshen_4_dtc_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(tmp9762, tmp9787)
return


}, 1)

tmp9788 := Call(__e, ns2_1set, symload, tmp9761)


_ = tmp9788

tmp9789 := MakeNative(func(__e *ControlFlow) {
V918 := __e.Get(1)
_ = V918
V919 := __e.Get(2)
_ = V919
tmp9791 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(False, V918)
}
__typedArg0 := False
__typedArg1 := V918
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9791 {
__e.TailApply(PrimFunc(symshen_4eval_1and_1print), V919)
return
} else {
__e.TailApply(PrimFunc(symshen_4check_1eval_1and_1print), V919)
return
}


}, 2)

tmp9792 := Call(__e, ns2_1set, symshen_4load_1help, tmp9789)


_ = tmp9792

tmp9793 := MakeNative(func(__e *ControlFlow) {
V920 := __e.Get(1)
_ = V920
tmp9794 := MakeNative(func(__e *ControlFlow) {
Z921 := __e.Get(1)
_ = Z921
tmp9795 := Call(__e, PrimFunc(symshen_4shen_1_6kl), Z921)


tmp9796 := Call(__e, PrimFunc(symeval_1kl), tmp9795)


tmp9797 := Call(__e, PrimFunc(symshen_4app), tmp9796, MakeString("\n"), symshen_4s)


tmp9798 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp9797, tmp9798)
return


}, 1)

__e.TailApply(PrimFunc(symmap), tmp9794, V920)
return


}, 1)

tmp9799 := Call(__e, ns2_1set, symshen_4eval_1and_1print, tmp9793)


_ = tmp9799

tmp9800 := MakeNative(func(__e *ControlFlow) {
V922 := __e.Get(1)
_ = V922
tmp9801 := MakeNative(func(__e *ControlFlow) {
W923 := __e.Get(1)
_ = W923
tmp9802 := MakeNative(func(__e *ControlFlow) {
W925 := __e.Get(1)
_ = W925
tmp9803 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4work_1through), V922)
return
}, 0)

tmp9804 := MakeNative(func(__e *ControlFlow) {
Z927 := __e.Get(1)
_ = Z927
__e.TailApply(PrimFunc(symshen_4unwind_1types), Z927, W923)
return
}, 1)

__e.TailApply(try_1catch, tmp9803, tmp9804)
return


}, 1)

tmp9805 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4assumetypes), W923)
return
}, 0)

tmp9806 := MakeNative(func(__e *ControlFlow) {
Z926 := __e.Get(1)
_ = Z926
__e.TailApply(PrimFunc(symshen_4unwind_1types), Z926, W923)
return
}, 1)

tmp9807 := Call(__e, try_1catch, tmp9805, tmp9806)


__e.TailApply(tmp9802, tmp9807)
return


}, 1)

tmp9808 := MakeNative(func(__e *ControlFlow) {
Z924 := __e.Get(1)
_ = Z924
__e.TailApply(PrimFunc(symshen_4typetable), Z924)
return
}, 1)

tmp9809 := Call(__e, PrimFunc(symmapcan), tmp9808, V922)


__e.TailApply(tmp9801, tmp9809)
return


}, 1)

tmp9810 := Call(__e, ns2_1set, symshen_4check_1eval_1and_1print, tmp9800)


_ = tmp9810

tmp9811 := MakeNative(func(__e *ControlFlow) {
V932 := __e.Get(1)
_ = V932
tmp9856 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V932)
}
__typedArg0 := V932
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9837 Obj

if True == tmp9856 {
tmp9854 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V932)
}
__typedArg0 := V932
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9855 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefine, tmp9854)
}
__typedArg0 := symdefine
__typedArg1 := tmp9854
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9839 Obj

if True == tmp9855 {
tmp9852 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V932)
}
__typedArg0 := V932
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9853 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9852)
}
__typedArg0 := tmp9852
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9841 Obj

if True == tmp9853 {
tmp9849 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V932)
}
__typedArg0 := V932
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9850 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9849)
}
__typedArg0 := tmp9849
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9851 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9850)
}
__typedArg0 := tmp9850
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9843 Obj

if True == tmp9851 {
tmp9845 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V932)
}
__typedArg0 := V932
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9846 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9845)
}
__typedArg0 := tmp9845
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9847 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9846)
}
__typedArg0 := tmp9846
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9848 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_i, tmp9847)
}
__typedArg0 := sym_i
__typedArg1 := tmp9847
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9844 Obj

if True == tmp9848 {
ifres9844 = True


} else {
ifres9844 = False


}

ifres9843 = ifres9844


} else {
ifres9843 = False


}

var ifres9842 Obj

if True == ifres9843 {
ifres9842 = True


} else {
ifres9842 = False


}

ifres9841 = ifres9842


} else {
ifres9841 = False


}

var ifres9840 Obj

if True == ifres9841 {
ifres9840 = True


} else {
ifres9840 = False


}

ifres9839 = ifres9840


} else {
ifres9839 = False


}

var ifres9838 Obj

if True == ifres9839 {
ifres9838 = True


} else {
ifres9838 = False


}

ifres9837 = ifres9838


} else {
ifres9837 = False


}

if True == ifres9837 {
tmp9812 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V932)
}
__typedArg0 := V932
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9813 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9812)
}
__typedArg0 := tmp9812
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9814 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V932)
}
__typedArg0 := V932
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9815 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9814)
}
__typedArg0 := tmp9814
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9816 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V932)
}
__typedArg0 := V932
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9817 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9816)
}
__typedArg0 := tmp9816
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9818 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9817)
}
__typedArg0 := tmp9817
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9819 := Call(__e, PrimFunc(symshen_4type_1F), tmp9815, tmp9818)


tmp9820 := Call(__e, PrimFunc(symshen_4rectify_1type), tmp9819)


tmp9821 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9820, Nil)
}
__typedArg0 := tmp9820
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9813, tmp9821)
}
__typedArg0 := tmp9813
__typedArg1 := tmp9821
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9835 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V932)
}
__typedArg0 := V932
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9827 Obj

if True == tmp9835 {
tmp9833 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V932)
}
__typedArg0 := V932
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9834 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefine, tmp9833)
}
__typedArg0 := symdefine
__typedArg1 := tmp9833
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9829 Obj

if True == tmp9834 {
tmp9831 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V932)
}
__typedArg0 := V932
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9832 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9831)
}
__typedArg0 := tmp9831
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9830 Obj

if True == tmp9832 {
ifres9830 = True


} else {
ifres9830 = False


}

ifres9829 = ifres9830


} else {
ifres9829 = False


}

var ifres9828 Obj

if True == ifres9829 {
ifres9828 = True


} else {
ifres9828 = False


}

ifres9827 = ifres9828


} else {
ifres9827 = False


}

if True == ifres9827 {
tmp9822 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V932)
}
__typedArg0 := V932
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9823 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9822)
}
__typedArg0 := tmp9822
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9824 := Call(__e, PrimFunc(symshen_4app), tmp9823, MakeString("\n"), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("missing { in "))
__typedS1, __typedOK1 := TypedString(tmp9824)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("missing { in ")
__typedArg1 := tmp9824
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("missing { in "))
__typedS1, __typedOK1 := TypedString(tmp9824)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("missing { in ")
__typedArg1 := tmp9824
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


} else {
__e.Return(Nil)
return
}


}


}, 1)

tmp9857 := Call(__e, ns2_1set, symshen_4typetable, tmp9811)


_ = tmp9857

tmp9858 := MakeNative(func(__e *ControlFlow) {
V939 := __e.Get(1)
_ = V939
V940 := __e.Get(2)
_ = V940
tmp9871 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V940)
}
__typedArg0 := V940
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9867 Obj

if True == tmp9871 {
tmp9869 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V940)
}
__typedArg0 := V940
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9870 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_j, tmp9869)
}
__typedArg0 := sym_j
__typedArg1 := tmp9869
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9868 Obj

if True == tmp9870 {
ifres9868 = True


} else {
ifres9868 = False


}

ifres9867 = ifres9868


} else {
ifres9867 = False


}

if True == ifres9867 {
__e.Return(Nil)
return
} else {
tmp9865 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V940)
}
__typedArg0 := V940
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9865 {
tmp9859 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V940)
}
__typedArg0 := V940
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9860 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V940)
}
__typedArg0 := V940
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9861 := Call(__e, PrimFunc(symshen_4type_1F), V939, tmp9860)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9859, tmp9861)
}
__typedArg0 := tmp9859
__typedArg1 := tmp9861
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9862 := Call(__e, PrimFunc(symshen_4app), V939, MakeString("\n"), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("missing } in "))
__typedS1, __typedOK1 := TypedString(tmp9862)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("missing } in ")
__typedArg1 := tmp9862
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("missing } in "))
__typedS1, __typedOK1 := TypedString(tmp9862)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("missing } in ")
__typedArg1 := tmp9862
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}


}


}, 2)

tmp9872 := Call(__e, ns2_1set, symshen_4type_1F, tmp9858)


_ = tmp9872

tmp9873 := MakeNative(func(__e *ControlFlow) {
V943 := __e.Get(1)
_ = V943
tmp9887 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V943)
}
__typedArg0 := Nil
__typedArg1 := V943
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9887 {
__e.Return(Nil)
return
} else {
tmp9885 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V943)
}
__typedArg0 := V943
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9881 Obj

if True == tmp9885 {
tmp9883 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V943)
}
__typedArg0 := V943
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9884 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9883)
}
__typedArg0 := tmp9883
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9882 Obj

if True == tmp9884 {
ifres9882 = True


} else {
ifres9882 = False


}

ifres9881 = ifres9882


} else {
ifres9881 = False


}

if True == ifres9881 {
tmp9874 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V943)
}
__typedArg0 := V943
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9875 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V943)
}
__typedArg0 := V943
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9876 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9875)
}
__typedArg0 := tmp9875
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9877 := Call(__e, PrimFunc(symdeclare), tmp9874, tmp9876)


_ = tmp9877

tmp9878 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V943)
}
__typedArg0 := V943
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9879 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9878)
}
__typedArg0 := tmp9878
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4assumetypes), tmp9879)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.assumetype"))
}
__typedArg0 := MakeString("implementation error in shen.assumetype")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 1)

tmp9888 := Call(__e, ns2_1set, symshen_4assumetypes, tmp9873)


_ = tmp9888

tmp9889 := MakeNative(func(__e *ControlFlow) {
V948 := __e.Get(1)
_ = V948
V949 := __e.Get(2)
_ = V949
tmp9900 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V949)
}
__typedArg0 := V949
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9896 Obj

if True == tmp9900 {
tmp9898 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V949)
}
__typedArg0 := V949
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9899 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9898)
}
__typedArg0 := tmp9898
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9897 Obj

if True == tmp9899 {
ifres9897 = True


} else {
ifres9897 = False


}

ifres9896 = ifres9897


} else {
ifres9896 = False


}

if True == ifres9896 {
tmp9890 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V949)
}
__typedArg0 := V949
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9891 := Call(__e, PrimFunc(symdestroy), tmp9890)


_ = tmp9891

tmp9892 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V949)
}
__typedArg0 := V949
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9893 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9892)
}
__typedArg0 := tmp9892
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4unwind_1types), V948, tmp9893)
return


} else {
tmp9894 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symerror_1to_1string) {
return PrimErrorToString(V948)
}
__typedArg0 := V948
return Call(__e, PrimFunc(symerror_1to_1string), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(tmp9894)
}
__typedArg0 := tmp9894
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}


}, 2)

tmp9901 := Call(__e, ns2_1set, symshen_4unwind_1types, tmp9889)


_ = tmp9901

tmp9902 := MakeNative(func(__e *ControlFlow) {
V952 := __e.Get(1)
_ = V952
tmp9951 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V952)
}
__typedArg0 := Nil
__typedArg1 := V952
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9951 {
__e.Return(Nil)
return
} else {
tmp9949 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V952)
}
__typedArg0 := V952
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9934 Obj

if True == tmp9949 {
tmp9947 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V952)
}
__typedArg0 := V952
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9948 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9947)
}
__typedArg0 := tmp9947
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9936 Obj

if True == tmp9948 {
tmp9944 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V952)
}
__typedArg0 := V952
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9945 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9944)
}
__typedArg0 := tmp9944
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9946 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9945)
}
__typedArg0 := tmp9945
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9938 Obj

if True == tmp9946 {
tmp9940 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V952)
}
__typedArg0 := V952
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9941 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9940)
}
__typedArg0 := tmp9940
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9942 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp9943 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp9941, tmp9942)
}
__typedArg0 := tmp9941
__typedArg1 := tmp9942
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9939 Obj

if True == tmp9943 {
ifres9939 = True


} else {
ifres9939 = False


}

ifres9938 = ifres9939


} else {
ifres9938 = False


}

var ifres9937 Obj

if True == ifres9938 {
ifres9937 = True


} else {
ifres9937 = False


}

ifres9936 = ifres9937


} else {
ifres9936 = False


}

var ifres9935 Obj

if True == ifres9936 {
ifres9935 = True


} else {
ifres9935 = False


}

ifres9934 = ifres9935


} else {
ifres9934 = False


}

if True == ifres9934 {
tmp9903 := MakeNative(func(__e *ControlFlow) {
W953 := __e.Get(1)
_ = W953
tmp9919 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W953, False)
}
__typedArg0 := W953
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9919 {
__e.TailApply(PrimFunc(symshen_4type_1error))
return
} else {
tmp9904 := MakeNative(func(__e *ControlFlow) {
W954 := __e.Get(1)
_ = W954
tmp9905 := MakeNative(func(__e *ControlFlow) {
W955 := __e.Get(1)
_ = W955
tmp9906 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V952)
}
__typedArg0 := V952
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9907 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9906)
}
__typedArg0 := tmp9906
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9908 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9907)
}
__typedArg0 := tmp9907
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4work_1through), tmp9908)
return


}, 1)

tmp9909 := Call(__e, PrimFunc(symshen_4pretty_1type), W953)


tmp9910 := Call(__e, PrimFunc(symshen_4app), tmp9909, MakeString("\n"), symshen_4r)


tmp9912 := Call(__e, PrimFunc(symshen_4app), W954, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(" : "))
__typedS1, __typedOK1 := TypedString(tmp9910)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(" : ")
__typedArg1 := tmp9910
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4s)


tmp9913 := Call(__e, PrimFunc(symstoutput))


tmp9914 := Call(__e, PrimFunc(sympr), tmp9912, tmp9913)


__e.TailApply(tmp9905, tmp9914)
return


}, 1)

tmp9915 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V952)
}
__typedArg0 := V952
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9916 := Call(__e, PrimFunc(symshen_4shen_1_6kl), tmp9915)


tmp9917 := Call(__e, PrimFunc(symeval_1kl), tmp9916)


__e.TailApply(tmp9904, tmp9917)
return


}


}, 1)

tmp9920 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V952)
}
__typedArg0 := V952
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9921 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V952)
}
__typedArg0 := V952
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9922 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9921)
}
__typedArg0 := tmp9921
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9923 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9922)
}
__typedArg0 := tmp9922
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9924 := Call(__e, PrimFunc(symshen_4typecheck), tmp9920, tmp9923)


__e.TailApply(tmp9903, tmp9924)
return


} else {
tmp9932 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V952)
}
__typedArg0 := V952
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9932 {
tmp9925 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V952)
}
__typedArg0 := V952
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9926 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp9927 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V952)
}
__typedArg0 := V952
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9928 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp9927)
}
__typedArg0 := symA
__typedArg1 := tmp9927
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9929 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9926, tmp9928)
}
__typedArg0 := tmp9926
__typedArg1 := tmp9928
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9930 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9925, tmp9929)
}
__typedArg0 := tmp9925
__typedArg1 := tmp9929
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4work_1through), tmp9930)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.work-through"))
}
__typedArg0 := MakeString("implementation error in shen.work-through")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}


}, 1)

tmp9952 := Call(__e, ns2_1set, symshen_4work_1through, tmp9902)


_ = tmp9952

tmp9953 := MakeNative(func(__e *ControlFlow) {
V957 := __e.Get(1)
_ = V957
tmp10095 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V957)
}
__typedArg0 := V957
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9969 Obj

if True == tmp10095 {
tmp10093 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V957)
}
__typedArg0 := V957
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10094 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10093)
}
__typedArg0 := tmp10093
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9971 Obj

if True == tmp10094 {
tmp10090 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V957)
}
__typedArg0 := V957
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10091 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10090)
}
__typedArg0 := tmp10090
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10092 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlist, tmp10091)
}
__typedArg0 := symlist
__typedArg1 := tmp10091
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9973 Obj

if True == tmp10092 {
tmp10087 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V957)
}
__typedArg0 := V957
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10088 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10087)
}
__typedArg0 := tmp10087
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10089 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10088)
}
__typedArg0 := tmp10088
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9975 Obj

if True == tmp10089 {
tmp10083 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V957)
}
__typedArg0 := V957
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10084 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10083)
}
__typedArg0 := tmp10083
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10085 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10084)
}
__typedArg0 := tmp10084
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10086 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp10085)
}
__typedArg0 := Nil
__typedArg1 := tmp10085
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9977 Obj

if True == tmp10086 {
tmp10081 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V957)
}
__typedArg0 := V957
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10082 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10081)
}
__typedArg0 := tmp10081
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9979 Obj

if True == tmp10082 {
tmp10078 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V957)
}
__typedArg0 := V957
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10079 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10078)
}
__typedArg0 := tmp10078
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10080 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_1_1_6, tmp10079)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp10079
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9981 Obj

if True == tmp10080 {
tmp10075 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V957)
}
__typedArg0 := V957
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10076 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10075)
}
__typedArg0 := tmp10075
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10077 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10076)
}
__typedArg0 := tmp10076
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9983 Obj

if True == tmp10077 {
tmp10071 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V957)
}
__typedArg0 := V957
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10072 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10071)
}
__typedArg0 := tmp10071
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10073 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10072)
}
__typedArg0 := tmp10072
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10074 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10073)
}
__typedArg0 := tmp10073
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9985 Obj

if True == tmp10074 {
tmp10066 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V957)
}
__typedArg0 := V957
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10067 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10066)
}
__typedArg0 := tmp10066
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10068 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10067)
}
__typedArg0 := tmp10067
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10069 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10068)
}
__typedArg0 := tmp10068
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10070 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symstr, tmp10069)
}
__typedArg0 := symstr
__typedArg1 := tmp10069
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9987 Obj

if True == tmp10070 {
tmp10061 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V957)
}
__typedArg0 := V957
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10062 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10061)
}
__typedArg0 := tmp10061
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10063 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10062)
}
__typedArg0 := tmp10062
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10064 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10063)
}
__typedArg0 := tmp10063
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10065 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10064)
}
__typedArg0 := tmp10064
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9989 Obj

if True == tmp10065 {
tmp10055 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V957)
}
__typedArg0 := V957
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10056 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10055)
}
__typedArg0 := tmp10055
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10057 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10056)
}
__typedArg0 := tmp10056
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10058 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10057)
}
__typedArg0 := tmp10057
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10059 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10058)
}
__typedArg0 := tmp10058
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10060 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10059)
}
__typedArg0 := tmp10059
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9991 Obj

if True == tmp10060 {
tmp10048 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V957)
}
__typedArg0 := V957
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10049 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10048)
}
__typedArg0 := tmp10048
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10050 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10049)
}
__typedArg0 := tmp10049
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10051 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10050)
}
__typedArg0 := tmp10050
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10052 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10051)
}
__typedArg0 := tmp10051
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10053 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10052)
}
__typedArg0 := tmp10052
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10054 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlist, tmp10053)
}
__typedArg0 := symlist
__typedArg1 := tmp10053
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9993 Obj

if True == tmp10054 {
tmp10041 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V957)
}
__typedArg0 := V957
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10042 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10041)
}
__typedArg0 := tmp10041
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10043 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10042)
}
__typedArg0 := tmp10042
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10044 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10043)
}
__typedArg0 := tmp10043
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10045 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10044)
}
__typedArg0 := tmp10044
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10046 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10045)
}
__typedArg0 := tmp10045
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10047 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10046)
}
__typedArg0 := tmp10046
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9995 Obj

if True == tmp10047 {
tmp10033 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V957)
}
__typedArg0 := V957
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10034 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10033)
}
__typedArg0 := tmp10033
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10035 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10034)
}
__typedArg0 := tmp10034
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10036 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10035)
}
__typedArg0 := tmp10035
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10037 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10036)
}
__typedArg0 := tmp10036
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10038 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10037)
}
__typedArg0 := tmp10037
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10039 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10038)
}
__typedArg0 := tmp10038
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10040 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp10039)
}
__typedArg0 := Nil
__typedArg1 := tmp10039
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9997 Obj

if True == tmp10040 {
tmp10027 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V957)
}
__typedArg0 := V957
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10028 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10027)
}
__typedArg0 := tmp10027
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10029 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10028)
}
__typedArg0 := tmp10028
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10030 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10029)
}
__typedArg0 := tmp10029
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10031 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10030)
}
__typedArg0 := tmp10030
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10032 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10031)
}
__typedArg0 := tmp10031
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9999 Obj

if True == tmp10032 {
tmp10020 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V957)
}
__typedArg0 := V957
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10021 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10020)
}
__typedArg0 := tmp10020
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10022 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10021)
}
__typedArg0 := tmp10021
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10023 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10022)
}
__typedArg0 := tmp10022
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10024 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10023)
}
__typedArg0 := tmp10023
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10025 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10024)
}
__typedArg0 := tmp10024
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10026 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp10025)
}
__typedArg0 := Nil
__typedArg1 := tmp10025
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10001 Obj

if True == tmp10026 {
tmp10016 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V957)
}
__typedArg0 := V957
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10017 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10016)
}
__typedArg0 := tmp10016
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10018 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10017)
}
__typedArg0 := tmp10017
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10019 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp10018)
}
__typedArg0 := Nil
__typedArg1 := tmp10018
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10003 Obj

if True == tmp10019 {
tmp10005 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V957)
}
__typedArg0 := V957
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10006 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10005)
}
__typedArg0 := tmp10005
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10007 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10006)
}
__typedArg0 := tmp10006
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10008 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V957)
}
__typedArg0 := V957
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10009 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10008)
}
__typedArg0 := tmp10008
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10010 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10009)
}
__typedArg0 := tmp10009
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10011 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10010)
}
__typedArg0 := tmp10010
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10012 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10011)
}
__typedArg0 := tmp10011
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10013 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10012)
}
__typedArg0 := tmp10012
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10014 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10013)
}
__typedArg0 := tmp10013
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10015 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp10007, tmp10014)
}
__typedArg0 := tmp10007
__typedArg1 := tmp10014
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10004 Obj

if True == tmp10015 {
ifres10004 = True


} else {
ifres10004 = False


}

ifres10003 = ifres10004


} else {
ifres10003 = False


}

var ifres10002 Obj

if True == ifres10003 {
ifres10002 = True


} else {
ifres10002 = False


}

ifres10001 = ifres10002


} else {
ifres10001 = False


}

var ifres10000 Obj

if True == ifres10001 {
ifres10000 = True


} else {
ifres10000 = False


}

ifres9999 = ifres10000


} else {
ifres9999 = False


}

var ifres9998 Obj

if True == ifres9999 {
ifres9998 = True


} else {
ifres9998 = False


}

ifres9997 = ifres9998


} else {
ifres9997 = False


}

var ifres9996 Obj

if True == ifres9997 {
ifres9996 = True


} else {
ifres9996 = False


}

ifres9995 = ifres9996


} else {
ifres9995 = False


}

var ifres9994 Obj

if True == ifres9995 {
ifres9994 = True


} else {
ifres9994 = False


}

ifres9993 = ifres9994


} else {
ifres9993 = False


}

var ifres9992 Obj

if True == ifres9993 {
ifres9992 = True


} else {
ifres9992 = False


}

ifres9991 = ifres9992


} else {
ifres9991 = False


}

var ifres9990 Obj

if True == ifres9991 {
ifres9990 = True


} else {
ifres9990 = False


}

ifres9989 = ifres9990


} else {
ifres9989 = False


}

var ifres9988 Obj

if True == ifres9989 {
ifres9988 = True


} else {
ifres9988 = False


}

ifres9987 = ifres9988


} else {
ifres9987 = False


}

var ifres9986 Obj

if True == ifres9987 {
ifres9986 = True


} else {
ifres9986 = False


}

ifres9985 = ifres9986


} else {
ifres9985 = False


}

var ifres9984 Obj

if True == ifres9985 {
ifres9984 = True


} else {
ifres9984 = False


}

ifres9983 = ifres9984


} else {
ifres9983 = False


}

var ifres9982 Obj

if True == ifres9983 {
ifres9982 = True


} else {
ifres9982 = False


}

ifres9981 = ifres9982


} else {
ifres9981 = False


}

var ifres9980 Obj

if True == ifres9981 {
ifres9980 = True


} else {
ifres9980 = False


}

ifres9979 = ifres9980


} else {
ifres9979 = False


}

var ifres9978 Obj

if True == ifres9979 {
ifres9978 = True


} else {
ifres9978 = False


}

ifres9977 = ifres9978


} else {
ifres9977 = False


}

var ifres9976 Obj

if True == ifres9977 {
ifres9976 = True


} else {
ifres9976 = False


}

ifres9975 = ifres9976


} else {
ifres9975 = False


}

var ifres9974 Obj

if True == ifres9975 {
ifres9974 = True


} else {
ifres9974 = False


}

ifres9973 = ifres9974


} else {
ifres9973 = False


}

var ifres9972 Obj

if True == ifres9973 {
ifres9972 = True


} else {
ifres9972 = False


}

ifres9971 = ifres9972


} else {
ifres9971 = False


}

var ifres9970 Obj

if True == ifres9971 {
ifres9970 = True


} else {
ifres9970 = False


}

ifres9969 = ifres9970


} else {
ifres9969 = False


}

if True == ifres9969 {
tmp9954 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V957)
}
__typedArg0 := V957
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9955 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9954)
}
__typedArg0 := tmp9954
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9956 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9955)
}
__typedArg0 := tmp9955
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9957 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9956)
}
__typedArg0 := tmp9956
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9958 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9957)
}
__typedArg0 := tmp9957
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9959 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V957)
}
__typedArg0 := V957
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9960 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9959)
}
__typedArg0 := tmp9959
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9961 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9960)
}
__typedArg0 := tmp9960
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9962 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9961)
}
__typedArg0 := tmp9961
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9963 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9962)
}
__typedArg0 := tmp9962
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9964 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_a_a_6, tmp9963)
}
__typedArg0 := sym_a_a_6
__typedArg1 := tmp9963
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9958, tmp9964)
}
__typedArg0 := tmp9958
__typedArg1 := tmp9964
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9967 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V957)
}
__typedArg0 := V957
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9967 {
tmp9965 := MakeNative(func(__e *ControlFlow) {
Z958 := __e.Get(1)
_ = Z958
__e.TailApply(PrimFunc(symshen_4pretty_1type), Z958)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp9965, V957)
return


} else {
__e.Return(V957)
return
}


}


}, 1)

tmp10096 := Call(__e, ns2_1set, symshen_4pretty_1type, tmp9953)


_ = tmp10096

tmp10097 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("type error\n"))
}
__typedArg0 := MakeString("type error\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}, 0)

tmp10098 := Call(__e, ns2_1set, symshen_4type_1error, tmp10097)


_ = tmp10098

tmp10099 := MakeNative(func(__e *ControlFlow) {
V959 := __e.Get(1)
_ = V959
tmp10100 := MakeNative(func(__e *ControlFlow) {
W960 := __e.Get(1)
_ = W960
tmp10101 := MakeNative(func(__e *ControlFlow) {
W961 := __e.Get(1)
_ = W961
tmp10102 := MakeNative(func(__e *ControlFlow) {
W962 := __e.Get(1)
_ = W962
tmp10103 := MakeNative(func(__e *ControlFlow) {
W963 := __e.Get(1)
_ = W963
tmp10104 := MakeNative(func(__e *ControlFlow) {
W965 := __e.Get(1)
_ = W965
__e.Return(W960)
return
}, 1)

tmp10105 := Call(__e, PrimFunc(symshen_4write_1kl), W963, W962)


__e.TailApply(tmp10104, tmp10105)
return


}, 1)

tmp10106 := MakeNative(func(__e *ControlFlow) {
Z964 := __e.Get(1)
_ = Z964
tmp10107 := Call(__e, PrimFunc(symshen_4shen_1_6kl_1h), Z964)


__e.TailApply(PrimFunc(symshen_4partial), tmp10107)
return


}, 1)

tmp10108 := Call(__e, PrimFunc(symmap), tmp10106, W961)


__e.TailApply(tmp10103, tmp10108)
return


}, 1)

tmp10109 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symopen) {
return PrimOpenStream(W960, symout)
}
__typedArg0 := W960
__typedArg1 := symout
return Call(__e, PrimFunc(symopen), __typedArg0, __typedArg1)
})()

__e.TailApply(tmp10102, tmp10109)
return


}, 1)

tmp10110 := Call(__e, PrimFunc(symread_1file), V959)


__e.TailApply(tmp10101, tmp10110)
return


}, 1)

tmp10111 := Call(__e, PrimFunc(symshen_4klfile), V959)


__e.TailApply(tmp10100, tmp10111)
return


}, 1)

tmp10112 := Call(__e, ns2_1set, symbootstrap, tmp10099)


_ = tmp10112

tmp10113 := MakeNative(func(__e *ControlFlow) {
V966 := __e.Get(1)
_ = V966
tmp10136 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V966)
}
__typedArg0 := V966
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10123 Obj

if True == tmp10136 {
tmp10134 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V966)
}
__typedArg0 := V966
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10135 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4f_1error, tmp10134)
}
__typedArg0 := symshen_4f_1error
__typedArg1 := tmp10134
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10125 Obj

if True == tmp10135 {
tmp10132 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V966)
}
__typedArg0 := V966
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10133 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10132)
}
__typedArg0 := tmp10132
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10127 Obj

if True == tmp10133 {
tmp10129 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V966)
}
__typedArg0 := V966
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10130 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10129)
}
__typedArg0 := tmp10129
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10131 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp10130)
}
__typedArg0 := Nil
__typedArg1 := tmp10130
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10128 Obj

if True == tmp10131 {
ifres10128 = True


} else {
ifres10128 = False


}

ifres10127 = ifres10128


} else {
ifres10127 = False


}

var ifres10126 Obj

if True == ifres10127 {
ifres10126 = True


} else {
ifres10126 = False


}

ifres10125 = ifres10126


} else {
ifres10125 = False


}

var ifres10124 Obj

if True == ifres10125 {
ifres10124 = True


} else {
ifres10124 = False


}

ifres10123 = ifres10124


} else {
ifres10123 = False


}

if True == ifres10123 {
tmp10114 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V966)
}
__typedArg0 := V966
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10115 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10114)
}
__typedArg0 := tmp10114
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10116 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(tmp10115)
}
__typedArg0 := tmp10115
return Call(__e, PrimFunc(symstr), __typedArg0)
})()

tmp10118 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("partial function "))
__typedS1, __typedOK1 := TypedString(tmp10116)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("partial function ")
__typedArg1 := tmp10116
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), Nil)
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("partial function "))
__typedS1, __typedOK1 := TypedString(tmp10116)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("partial function ")
__typedArg1 := tmp10116
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsimple_1error, tmp10118)
}
__typedArg0 := symsimple_1error
__typedArg1 := tmp10118
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp10121 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V966)
}
__typedArg0 := V966
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp10121 {
tmp10119 := MakeNative(func(__e *ControlFlow) {
Z967 := __e.Get(1)
_ = Z967
__e.TailApply(PrimFunc(symshen_4partial), Z967)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp10119, V966)
return


} else {
__e.Return(V966)
return
}


}


}, 1)

tmp10137 := Call(__e, ns2_1set, symshen_4partial, tmp10113)


_ = tmp10137

tmp10138 := MakeNative(func(__e *ControlFlow) {
V970 := __e.Get(1)
_ = V970
V971 := __e.Get(2)
_ = V971
tmp10152 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V970)
}
__typedArg0 := Nil
__typedArg1 := V970
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp10152 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symclose) {
return PrimCloseStream(V971)
}
__typedArg0 := V971
return Call(__e, PrimFunc(symclose), __typedArg0)
})())
return
} else {
tmp10150 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V970)
}
__typedArg0 := V970
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10146 Obj

if True == tmp10150 {
tmp10148 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V970)
}
__typedArg0 := V970
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10149 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10148)
}
__typedArg0 := tmp10148
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10147 Obj

if True == tmp10149 {
ifres10147 = True


} else {
ifres10147 = False


}

ifres10146 = ifres10147


} else {
ifres10146 = False


}

if True == ifres10146 {
tmp10139 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V970)
}
__typedArg0 := V970
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10140 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V970)
}
__typedArg0 := V970
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10141 := Call(__e, PrimFunc(symshen_4write_1kl_1h), tmp10140, V971)


_ = tmp10141

__e.TailApply(PrimFunc(symshen_4write_1kl), tmp10139, V971)
return


} else {
tmp10144 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V970)
}
__typedArg0 := V970
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp10144 {
tmp10142 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V970)
}
__typedArg0 := V970
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4write_1kl), tmp10142, V971)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("partial function shen.write-kl"))
}
__typedArg0 := MakeString("partial function shen.write-kl")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}


}, 2)

tmp10153 := Call(__e, ns2_1set, symshen_4write_1kl, tmp10138)


_ = tmp10153

tmp10154 := MakeNative(func(__e *ControlFlow) {
V974 := __e.Get(1)
_ = V974
V975 := __e.Get(2)
_ = V975
tmp10194 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V974)
}
__typedArg0 := V974
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10157 Obj

if True == tmp10194 {
tmp10192 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V974)
}
__typedArg0 := V974
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10193 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefun, tmp10192)
}
__typedArg0 := symdefun
__typedArg1 := tmp10192
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10159 Obj

if True == tmp10193 {
tmp10190 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V974)
}
__typedArg0 := V974
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10191 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10190)
}
__typedArg0 := tmp10190
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10161 Obj

if True == tmp10191 {
tmp10187 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V974)
}
__typedArg0 := V974
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10188 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10187)
}
__typedArg0 := tmp10187
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10189 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symfail, tmp10188)
}
__typedArg0 := symfail
__typedArg1 := tmp10188
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10163 Obj

if True == tmp10189 {
tmp10184 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V974)
}
__typedArg0 := V974
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10185 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10184)
}
__typedArg0 := tmp10184
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10186 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10185)
}
__typedArg0 := tmp10185
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10165 Obj

if True == tmp10186 {
tmp10180 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V974)
}
__typedArg0 := V974
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10181 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10180)
}
__typedArg0 := tmp10180
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10182 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10181)
}
__typedArg0 := tmp10181
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10183 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp10182)
}
__typedArg0 := Nil
__typedArg1 := tmp10182
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10167 Obj

if True == tmp10183 {
tmp10176 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V974)
}
__typedArg0 := V974
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10177 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10176)
}
__typedArg0 := tmp10176
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10178 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10177)
}
__typedArg0 := tmp10177
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10179 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10178)
}
__typedArg0 := tmp10178
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10169 Obj

if True == tmp10179 {
tmp10171 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V974)
}
__typedArg0 := V974
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10172 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10171)
}
__typedArg0 := tmp10171
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10173 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10172)
}
__typedArg0 := tmp10172
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10174 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10173)
}
__typedArg0 := tmp10173
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10175 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp10174)
}
__typedArg0 := Nil
__typedArg1 := tmp10174
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10170 Obj

if True == tmp10175 {
ifres10170 = True


} else {
ifres10170 = False


}

ifres10169 = ifres10170


} else {
ifres10169 = False


}

var ifres10168 Obj

if True == ifres10169 {
ifres10168 = True


} else {
ifres10168 = False


}

ifres10167 = ifres10168


} else {
ifres10167 = False


}

var ifres10166 Obj

if True == ifres10167 {
ifres10166 = True


} else {
ifres10166 = False


}

ifres10165 = ifres10166


} else {
ifres10165 = False


}

var ifres10164 Obj

if True == ifres10165 {
ifres10164 = True


} else {
ifres10164 = False


}

ifres10163 = ifres10164


} else {
ifres10163 = False


}

var ifres10162 Obj

if True == ifres10163 {
ifres10162 = True


} else {
ifres10162 = False


}

ifres10161 = ifres10162


} else {
ifres10161 = False


}

var ifres10160 Obj

if True == ifres10161 {
ifres10160 = True


} else {
ifres10160 = False


}

ifres10159 = ifres10160


} else {
ifres10159 = False


}

var ifres10158 Obj

if True == ifres10159 {
ifres10158 = True


} else {
ifres10158 = False


}

ifres10157 = ifres10158


} else {
ifres10157 = False


}

if True == ifres10157 {
__e.TailApply(PrimFunc(sympr), MakeString("(defun fail () shen.fail!)"), V975)
return
} else {
tmp10155 := Call(__e, PrimFunc(symshen_4app), V974, MakeString("\n\n"), symshen_4r)


__e.TailApply(PrimFunc(sympr), tmp10155, V975)
return


}


}, 2)

tmp10195 := Call(__e, ns2_1set, symshen_4write_1kl_1h, tmp10154)


_ = tmp10195

tmp10196 := MakeNative(func(__e *ControlFlow) {
V976 := __e.Get(1)
_ = V976
tmp10205 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString(""), V976)
}
__typedArg0 := MakeString("")
__typedArg1 := V976
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp10205 {
__e.Return(MakeString(".kl"))
return
} else {
tmp10203 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString(".shen"), V976)
}
__typedArg0 := MakeString(".shen")
__typedArg1 := V976
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp10203 {
__e.Return(MakeString(".kl"))
return
} else {
tmp10201 := Call(__e, PrimFunc(symshen_4_7string_2), V976)


if True == tmp10201 {
tmp10197 := Call(__e, PrimFunc(symhdstr), V976)


tmp10199 := Call(__e, PrimFunc(symshen_4klfile), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V976)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V976
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


__e.TailApply(PrimFunc(sym_8s), tmp10197, tmp10199)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("partial function shen.klfile"))
}
__typedArg0 := MakeString("partial function shen.klfile")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}


}, 1)

__e.TailApply(ns2_1set, symshen_4klfile, tmp10196)
return




}, 0)

