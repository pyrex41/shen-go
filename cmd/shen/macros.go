package main

import . "github.com/pyrex41/shen-go/kl"

var MacrosMain = MakeNative(func(__e *ControlFlow) {
tmp8596 := MakeNative(func(__e *ControlFlow) {
V7069 := __e.Get(1)
_ = V7069
tmp8597 := MakeNative(func(__e *ControlFlow) {
Z7071 := __e.Get(1)
_ = Z7071
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Z7071)
}
__typedArg0 := Z7071
return Call(__e, PrimFunc(symtl), __typedArg0)
})())
return
}, 1)

tmp8598 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dmacros_d)
}
__typedArg0 := sym_dmacros_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp8599 := Call(__e, PrimFunc(symmap), tmp8597, tmp8598)


let__8531 := tmp8599
_ = let__8531

tmp8600 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsource_1form_1handlers_d)
}
__typedArg0 := symshen_4_dsource_1form_1handlers_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4macroexpand_1h), V7069, let__8531, let__8531, tmp8600)
return


}, 1)

tmp8601 := Call(__e, ns2_1set, symmacroexpand, tmp8596)


_ = tmp8601

tmp8602 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__self3 := __e.Get(3)
__self4 := __e.Get(4)
__selftop:
V7084 := __self1
_ = V7084
V7085 := __self2
_ = V7085
V7086 := __self3
_ = V7086
V7087 := __self4
_ = V7087
tmp8611 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V7085)
}
__typedArg0 := Nil
__typedArg1 := V7085
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp8611 {
__e.Return(V7084)
return
} else {
tmp8609 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7085)
}
__typedArg0 := V7085
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp8609 {
tmp8603 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7085)
}
__typedArg0 := V7085
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8604 := Call(__e, PrimFunc(symshen_4source_1macro_1walk), tmp8603, V7084, V7087)


let__8532 := tmp8604
_ = let__8532

tmp8607 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V7084, let__8532)
}
__typedArg0 := V7084
__typedArg1 := let__8532
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp8607 {
tmp8605 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7085)
}
__typedArg0 := V7085
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4macroexpand_1h) == __self {
__self1, __self2, __self3, __self4 = V7084, tmp8605, V7086, V7087
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4macroexpand_1h), V7084, tmp8605, V7086, V7087)
return


} else {
if PrimFunc(symshen_4macroexpand_1h) == __self {
__self1, __self2, __self3, __self4 = let__8532, V7086, V7086, V7087
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4macroexpand_1h), let__8532, V7086, V7086, V7087)
return
}


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.macroexpand-h"))
}
__typedArg0 := MakeString("implementation error in shen.macroexpand-h")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 4)

tmp8612 := Call(__e, ns2_1set, symshen_4macroexpand_1h, tmp8602)


_ = tmp8612

tmp8613 := MakeNative(func(__e *ControlFlow) {
V7092 := __e.Get(1)
_ = V7092
V7093 := __e.Get(2)
_ = V7093
V7094 := __e.Get(3)
_ = V7094
tmp8622 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V7094)
}
__typedArg0 := Nil
__typedArg1 := V7094
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp8622 {
__e.TailApply(PrimFunc(symshen_4walk), V7092, V7093)
return
} else {
tmp8619 := Call(__e, PrimFunc(symshen_4source_1form_1handler), V7093, V7094)


tmp8620 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp8619)
}
__typedArg0 := tmp8619
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp8620 {
__e.Return(V7093)
return
} else {
tmp8617 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7093)
}
__typedArg0 := V7093
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp8617 {
tmp8614 := MakeNative(func(__e *ControlFlow) {
Z7095 := __e.Get(1)
_ = Z7095
__e.TailApply(PrimFunc(symshen_4source_1macro_1walk), V7092, Z7095, V7094)
return
}, 1)

tmp8615 := Call(__e, PrimFunc(symmap), tmp8614, V7093)


__e.TailApply(V7092, tmp8615)
return


} else {
__e.TailApply(V7092, V7093)
return
}


}


}


}, 3)

tmp8623 := Call(__e, ns2_1set, symshen_4source_1macro_1walk, tmp8613)


_ = tmp8623

tmp8624 := MakeNative(func(__e *ControlFlow) {
V7096 := __e.Get(1)
_ = V7096
V7097 := __e.Get(2)
_ = V7097
tmp8628 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7097)
}
__typedArg0 := V7097
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp8628 {
tmp8625 := MakeNative(func(__e *ControlFlow) {
Z7098 := __e.Get(1)
_ = Z7098
__e.TailApply(PrimFunc(symshen_4walk), V7096, Z7098)
return
}, 1)

tmp8626 := Call(__e, PrimFunc(symmap), tmp8625, V7097)


__e.TailApply(V7096, tmp8626)
return


} else {
__e.TailApply(V7096, V7097)
return
}


}, 2)

tmp8629 := Call(__e, ns2_1set, symshen_4walk, tmp8624)


_ = tmp8629

tmp8630 := MakeNative(func(__e *ControlFlow) {
V7099 := __e.Get(1)
_ = V7099
tmp8631 := MakeNative(func(__e *ControlFlow) {
__e.Return(V7099)
return
}, 0)

let__8533 := tmp8631
_ = let__8533

tmp8929 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7099)
}
__typedArg0 := V7099
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp8929 {
tmp8632 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7099)
}
__typedArg0 := V7099
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__8534 := tmp8632
_ = let__8534

tmp8633 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7099)
}
__typedArg0 := V7099
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8535 := tmp8633
_ = let__8535

tmp8927 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefmacro, let__8534)
}
__typedArg0 := symdefmacro
__typedArg1 := let__8534
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8924 Obj

if True == tmp8927 {
tmp8926 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8925 Obj

if True == tmp8926 {
ifres8925 = True


} else {
ifres8925 = False


}

ifres8924 = ifres8925


} else {
ifres8924 = False


}

if True == ifres8924 {
tmp8634 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8635 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4process_1def), tmp8634, tmp8635)
return


} else {
tmp8922 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefcc, let__8534)
}
__typedArg0 := symdefcc
__typedArg1 := let__8534
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp8922 {
__e.TailApply(PrimFunc(symshen_4yacc_1_6shen), let__8535)
return
} else {
tmp8920 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symu_b, let__8534)
}
__typedArg0 := symu_b
__typedArg1 := let__8534
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8913 Obj

if True == tmp8920 {
tmp8919 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8915 Obj

if True == tmp8919 {
tmp8917 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8918 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp8917)
}
__typedArg0 := Nil
__typedArg1 := tmp8917
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8916 Obj

if True == tmp8918 {
ifres8916 = True


} else {
ifres8916 = False


}

ifres8915 = ifres8916


} else {
ifres8915 = False


}

var ifres8914 Obj

if True == ifres8915 {
ifres8914 = True


} else {
ifres8914 = False


}

ifres8913 = ifres8914


} else {
ifres8913 = False


}

if True == ifres8913 {
tmp8636 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8637 := Call(__e, PrimFunc(symshen_4make_1uppercase), tmp8636)


tmp8638 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8637, Nil)
}
__typedArg0 := tmp8637
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symprotect, tmp8638)
}
__typedArg0 := symprotect
__typedArg1 := tmp8638
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp8911 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symerror, let__8534)
}
__typedArg0 := symerror
__typedArg1 := let__8534
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8908 Obj

if True == tmp8911 {
tmp8910 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8909 Obj

if True == tmp8910 {
ifres8909 = True


} else {
ifres8909 = False


}

ifres8908 = ifres8909


} else {
ifres8908 = False


}

if True == ifres8908 {
tmp8639 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8640 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8641 := Call(__e, PrimFunc(symshen_4mkstr), tmp8639, tmp8640)


tmp8642 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8641, Nil)
}
__typedArg0 := tmp8641
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsimple_1error, tmp8642)
}
__typedArg0 := symsimple_1error
__typedArg1 := tmp8642
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp8906 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symoutput, let__8534)
}
__typedArg0 := symoutput
__typedArg1 := let__8534
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8903 Obj

if True == tmp8906 {
tmp8905 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8904 Obj

if True == tmp8905 {
ifres8904 = True


} else {
ifres8904 = False


}

ifres8903 = ifres8904


} else {
ifres8903 = False


}

if True == ifres8903 {
tmp8643 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8644 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8645 := Call(__e, PrimFunc(symshen_4mkstr), tmp8643, tmp8644)


tmp8646 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstoutput, Nil)
}
__typedArg0 := symstoutput
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8647 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8646, Nil)
}
__typedArg0 := tmp8646
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8648 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8645, tmp8647)
}
__typedArg0 := tmp8645
__typedArg1 := tmp8647
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympr, tmp8648)
}
__typedArg0 := sympr
__typedArg1 := tmp8648
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp8901 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sympr, let__8534)
}
__typedArg0 := sympr
__typedArg1 := let__8534
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8894 Obj

if True == tmp8901 {
tmp8900 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8896 Obj

if True == tmp8900 {
tmp8898 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8899 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp8898)
}
__typedArg0 := Nil
__typedArg1 := tmp8898
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8897 Obj

if True == tmp8899 {
ifres8897 = True


} else {
ifres8897 = False


}

ifres8896 = ifres8897


} else {
ifres8896 = False


}

var ifres8895 Obj

if True == ifres8896 {
ifres8895 = True


} else {
ifres8895 = False


}

ifres8894 = ifres8895


} else {
ifres8894 = False


}

if True == ifres8894 {
tmp8649 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8650 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstoutput, Nil)
}
__typedArg0 := symstoutput
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8651 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8650, Nil)
}
__typedArg0 := tmp8650
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8652 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8649, tmp8651)
}
__typedArg0 := tmp8649
__typedArg1 := tmp8651
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympr, tmp8652)
}
__typedArg0 := sympr
__typedArg1 := tmp8652
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp8892 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symmake_1string, let__8534)
}
__typedArg0 := symmake_1string
__typedArg1 := let__8534
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8889 Obj

if True == tmp8892 {
tmp8891 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8890 Obj

if True == tmp8891 {
ifres8890 = True


} else {
ifres8890 = False


}

ifres8889 = ifres8890


} else {
ifres8889 = False


}

if True == ifres8889 {
tmp8653 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8654 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4mkstr), tmp8653, tmp8654)
return


} else {
tmp8887 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlineread, let__8534)
}
__typedArg0 := symlineread
__typedArg1 := let__8534
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8884 Obj

if True == tmp8887 {
tmp8886 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, let__8535)
}
__typedArg0 := Nil
__typedArg1 := let__8535
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8885 Obj

if True == tmp8886 {
ifres8885 = True


} else {
ifres8885 = False


}

ifres8884 = ifres8885


} else {
ifres8884 = False


}

if True == ifres8884 {
tmp8655 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstinput, Nil)
}
__typedArg0 := symstinput
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8656 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8655, Nil)
}
__typedArg0 := tmp8655
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlineread, tmp8656)
}
__typedArg0 := symlineread
__typedArg1 := tmp8656
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp8882 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(syminput, let__8534)
}
__typedArg0 := syminput
__typedArg1 := let__8534
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8879 Obj

if True == tmp8882 {
tmp8881 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, let__8535)
}
__typedArg0 := Nil
__typedArg1 := let__8535
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8880 Obj

if True == tmp8881 {
ifres8880 = True


} else {
ifres8880 = False


}

ifres8879 = ifres8880


} else {
ifres8879 = False


}

if True == ifres8879 {
tmp8657 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstinput, Nil)
}
__typedArg0 := symstinput
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8658 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8657, Nil)
}
__typedArg0 := tmp8657
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminput, tmp8658)
}
__typedArg0 := syminput
__typedArg1 := tmp8658
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp8877 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symread, let__8534)
}
__typedArg0 := symread
__typedArg1 := let__8534
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8874 Obj

if True == tmp8877 {
tmp8876 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, let__8535)
}
__typedArg0 := Nil
__typedArg1 := let__8535
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8875 Obj

if True == tmp8876 {
ifres8875 = True


} else {
ifres8875 = False


}

ifres8874 = ifres8875


} else {
ifres8874 = False


}

if True == ifres8874 {
tmp8659 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstinput, Nil)
}
__typedArg0 := symstinput
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8660 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8659, Nil)
}
__typedArg0 := tmp8659
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread, tmp8660)
}
__typedArg0 := symread
__typedArg1 := tmp8660
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp8872 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(syminput_7, let__8534)
}
__typedArg0 := syminput_7
__typedArg1 := let__8534
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8869 Obj

if True == tmp8872 {
tmp8871 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8870 Obj

if True == tmp8871 {
ifres8870 = True


} else {
ifres8870 = False


}

ifres8869 = ifres8870


} else {
ifres8869 = False


}

if True == ifres8869 {
__e.TailApply(PrimFunc(symshen_4process_1input_7), V7099)
return
} else {
tmp8867 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symread_1byte, let__8534)
}
__typedArg0 := symread_1byte
__typedArg1 := let__8534
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8864 Obj

if True == tmp8867 {
tmp8866 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, let__8535)
}
__typedArg0 := Nil
__typedArg1 := let__8535
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8865 Obj

if True == tmp8866 {
ifres8865 = True


} else {
ifres8865 = False


}

ifres8864 = ifres8865


} else {
ifres8864 = False


}

if True == ifres8864 {
__e.TailApply(PrimFunc(symshen_4process_1read_1byte))
return
} else {
tmp8862 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symprolog_2, let__8534)
}
__typedArg0 := symprolog_2
__typedArg1 := let__8534
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp8862 {
__e.TailApply(PrimFunc(symshen_4call_1prolog), let__8535)
return
} else {
tmp8860 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefprolog, let__8534)
}
__typedArg0 := symdefprolog
__typedArg1 := let__8534
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8857 Obj

if True == tmp8860 {
tmp8859 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8858 Obj

if True == tmp8859 {
ifres8858 = True


} else {
ifres8858 = False


}

ifres8857 = ifres8858


} else {
ifres8857 = False


}

if True == ifres8857 {
tmp8661 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8662 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4compile_1prolog), tmp8661, tmp8662)
return


} else {
tmp8855 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdatatype, let__8534)
}
__typedArg0 := symdatatype
__typedArg1 := let__8534
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8852 Obj

if True == tmp8855 {
tmp8854 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8853 Obj

if True == tmp8854 {
ifres8853 = True


} else {
ifres8853 = False


}

ifres8852 = ifres8853


} else {
ifres8852 = False


}

if True == ifres8852 {
tmp8663 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8664 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4process_1datatype), tmp8663, tmp8664)
return


} else {
tmp8850 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_8s, let__8534)
}
__typedArg0 := sym_8s
__typedArg1 := let__8534
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp8850 {
__e.TailApply(PrimFunc(symshen_4process_1_8s), V7099)
return
} else {
tmp8848 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symsynonyms, let__8534)
}
__typedArg0 := symsynonyms
__typedArg1 := let__8534
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp8848 {
__e.TailApply(PrimFunc(symshen_4process_1synonyms), let__8535)
return
} else {
tmp8846 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symnl, let__8534)
}
__typedArg0 := symnl
__typedArg1 := let__8534
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8843 Obj

if True == tmp8846 {
tmp8845 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, let__8535)
}
__typedArg0 := Nil
__typedArg1 := let__8535
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8844 Obj

if True == tmp8845 {
ifres8844 = True


} else {
ifres8844 = False


}

ifres8843 = ifres8844


} else {
ifres8843 = False


}

if True == ifres8843 {
tmp8665 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), Nil)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnl, tmp8665)
}
__typedArg0 := symnl
__typedArg1 := tmp8665
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp8841 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlet, let__8534)
}
__typedArg0 := symlet
__typedArg1 := let__8534
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp8841 {
__e.TailApply(PrimFunc(symshen_4process_1let), V7099)
return
} else {
tmp8839 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_c_4, let__8534)
}
__typedArg0 := sym_c_4
__typedArg1 := let__8534
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp8839 {
__e.TailApply(PrimFunc(symshen_4process_1lambda), V7099)
return
} else {
tmp8837 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcases, let__8534)
}
__typedArg0 := symcases
__typedArg1 := let__8534
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp8837 {
__e.TailApply(PrimFunc(symshen_4process_1cases), V7099)
return
} else {
tmp8835 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symtime, let__8534)
}
__typedArg0 := symtime
__typedArg1 := let__8534
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8828 Obj

if True == tmp8835 {
tmp8834 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8830 Obj

if True == tmp8834 {
tmp8832 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8833 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp8832)
}
__typedArg0 := Nil
__typedArg1 := tmp8832
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8831 Obj

if True == tmp8833 {
ifres8831 = True


} else {
ifres8831 = False


}

ifres8830 = ifres8831


} else {
ifres8830 = False


}

var ifres8829 Obj

if True == ifres8830 {
ifres8829 = True


} else {
ifres8829 = False


}

ifres8828 = ifres8829


} else {
ifres8828 = False


}

if True == ifres8828 {
tmp8666 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4process_1time), tmp8666)
return


} else {
tmp8826 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symput, let__8534)
}
__typedArg0 := symput
__typedArg1 := let__8534
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8808 Obj

if True == tmp8826 {
tmp8825 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8810 Obj

if True == tmp8825 {
tmp8823 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8824 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp8823)
}
__typedArg0 := tmp8823
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8812 Obj

if True == tmp8824 {
tmp8820 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8821 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp8820)
}
__typedArg0 := tmp8820
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8822 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp8821)
}
__typedArg0 := tmp8821
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8814 Obj

if True == tmp8822 {
tmp8816 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8817 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp8816)
}
__typedArg0 := tmp8816
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8818 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp8817)
}
__typedArg0 := tmp8817
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8819 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp8818)
}
__typedArg0 := Nil
__typedArg1 := tmp8818
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8815 Obj

if True == tmp8819 {
ifres8815 = True


} else {
ifres8815 = False


}

ifres8814 = ifres8815


} else {
ifres8814 = False


}

var ifres8813 Obj

if True == ifres8814 {
ifres8813 = True


} else {
ifres8813 = False


}

ifres8812 = ifres8813


} else {
ifres8812 = False


}

var ifres8811 Obj

if True == ifres8812 {
ifres8811 = True


} else {
ifres8811 = False


}

ifres8810 = ifres8811


} else {
ifres8810 = False


}

var ifres8809 Obj

if True == ifres8810 {
ifres8809 = True


} else {
ifres8809 = False


}

ifres8808 = ifres8809


} else {
ifres8808 = False


}

if True == ifres8808 {
tmp8667 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8668 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8669 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp8668)
}
__typedArg0 := tmp8668
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8670 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8671 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp8670)
}
__typedArg0 := tmp8670
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8672 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp8671)
}
__typedArg0 := tmp8671
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8673 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dproperty_1vector_d, Nil)
}
__typedArg0 := sym_dproperty_1vector_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8674 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvalue, tmp8673)
}
__typedArg0 := symvalue
__typedArg1 := tmp8673
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8675 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8674, Nil)
}
__typedArg0 := tmp8674
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8676 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8672, tmp8675)
}
__typedArg0 := tmp8672
__typedArg1 := tmp8675
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8677 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8669, tmp8676)
}
__typedArg0 := tmp8669
__typedArg1 := tmp8676
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8678 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8667, tmp8677)
}
__typedArg0 := tmp8667
__typedArg1 := tmp8677
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symput, tmp8678)
}
__typedArg0 := symput
__typedArg1 := tmp8678
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp8806 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symget, let__8534)
}
__typedArg0 := symget
__typedArg1 := let__8534
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8794 Obj

if True == tmp8806 {
tmp8805 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8796 Obj

if True == tmp8805 {
tmp8803 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8804 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp8803)
}
__typedArg0 := tmp8803
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8798 Obj

if True == tmp8804 {
tmp8800 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8801 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp8800)
}
__typedArg0 := tmp8800
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8802 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp8801)
}
__typedArg0 := Nil
__typedArg1 := tmp8801
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8799 Obj

if True == tmp8802 {
ifres8799 = True


} else {
ifres8799 = False


}

ifres8798 = ifres8799


} else {
ifres8798 = False


}

var ifres8797 Obj

if True == ifres8798 {
ifres8797 = True


} else {
ifres8797 = False


}

ifres8796 = ifres8797


} else {
ifres8796 = False


}

var ifres8795 Obj

if True == ifres8796 {
ifres8795 = True


} else {
ifres8795 = False


}

ifres8794 = ifres8795


} else {
ifres8794 = False


}

if True == ifres8794 {
tmp8679 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8680 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8681 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp8680)
}
__typedArg0 := tmp8680
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8682 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dproperty_1vector_d, Nil)
}
__typedArg0 := sym_dproperty_1vector_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8683 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvalue, tmp8682)
}
__typedArg0 := symvalue
__typedArg1 := tmp8682
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8684 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8683, Nil)
}
__typedArg0 := tmp8683
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8685 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8681, tmp8684)
}
__typedArg0 := tmp8681
__typedArg1 := tmp8684
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8686 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8679, tmp8685)
}
__typedArg0 := tmp8679
__typedArg1 := tmp8685
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symget, tmp8686)
}
__typedArg0 := symget
__typedArg1 := tmp8686
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp8792 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symunput, let__8534)
}
__typedArg0 := symunput
__typedArg1 := let__8534
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8780 Obj

if True == tmp8792 {
tmp8791 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8782 Obj

if True == tmp8791 {
tmp8789 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8790 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp8789)
}
__typedArg0 := tmp8789
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8784 Obj

if True == tmp8790 {
tmp8786 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8787 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp8786)
}
__typedArg0 := tmp8786
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8788 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp8787)
}
__typedArg0 := Nil
__typedArg1 := tmp8787
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8785 Obj

if True == tmp8788 {
ifres8785 = True


} else {
ifres8785 = False


}

ifres8784 = ifres8785


} else {
ifres8784 = False


}

var ifres8783 Obj

if True == ifres8784 {
ifres8783 = True


} else {
ifres8783 = False


}

ifres8782 = ifres8783


} else {
ifres8782 = False


}

var ifres8781 Obj

if True == ifres8782 {
ifres8781 = True


} else {
ifres8781 = False


}

ifres8780 = ifres8781


} else {
ifres8780 = False


}

if True == ifres8780 {
tmp8687 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8688 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8689 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp8688)
}
__typedArg0 := tmp8688
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8690 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dproperty_1vector_d, Nil)
}
__typedArg0 := sym_dproperty_1vector_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8691 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvalue, tmp8690)
}
__typedArg0 := symvalue
__typedArg1 := tmp8690
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8692 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8691, Nil)
}
__typedArg0 := tmp8691
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8693 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8689, tmp8692)
}
__typedArg0 := tmp8689
__typedArg1 := tmp8692
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8694 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8687, tmp8693)
}
__typedArg0 := tmp8687
__typedArg1 := tmp8693
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunput, tmp8694)
}
__typedArg0 := symunput
__typedArg1 := tmp8694
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp8778 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4_8c, let__8534)
}
__typedArg0 := symshen_4_8c
__typedArg1 := let__8534
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8771 Obj

if True == tmp8778 {
tmp8777 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8773 Obj

if True == tmp8777 {
tmp8775 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8776 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp8775)
}
__typedArg0 := Nil
__typedArg1 := tmp8775
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8774 Obj

if True == tmp8776 {
ifres8774 = True


} else {
ifres8774 = False


}

ifres8773 = ifres8774


} else {
ifres8773 = False


}

var ifres8772 Obj

if True == ifres8773 {
ifres8772 = True


} else {
ifres8772 = False


}

ifres8771 = ifres8772


} else {
ifres8771 = False


}

if True == ifres8771 {
tmp8695 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4rcons__form), tmp8695)
return


} else {
tmp8696 := MakeNative(func(__e *ControlFlow) {
tmp8724 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8704 Obj

if True == tmp8724 {
tmp8722 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8723 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp8722)
}
__typedArg0 := tmp8722
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8706 Obj

if True == tmp8723 {
tmp8719 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8720 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp8719)
}
__typedArg0 := tmp8719
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8721 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp8720)
}
__typedArg0 := tmp8720
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8708 Obj

if True == tmp8721 {
tmp8710 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdo, Nil)
}
__typedArg0 := symdo
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8711 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_d, tmp8710)
}
__typedArg0 := sym_d
__typedArg1 := tmp8710
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8712 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_7, tmp8711)
}
__typedArg0 := sym_7
__typedArg1 := tmp8711
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8713 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symor, tmp8712)
}
__typedArg0 := symor
__typedArg1 := tmp8712
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8714 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symand, tmp8713)
}
__typedArg0 := symand
__typedArg1 := tmp8713
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8715 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symappend, tmp8714)
}
__typedArg0 := symappend
__typedArg1 := tmp8714
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8716 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8v, tmp8715)
}
__typedArg0 := sym_8v
__typedArg1 := tmp8715
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8717 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8p, tmp8716)
}
__typedArg0 := sym_8p
__typedArg1 := tmp8716
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8718 := Call(__e, PrimFunc(symelement_2), let__8534, tmp8717)


var ifres8709 Obj

if True == tmp8718 {
ifres8709 = True


} else {
ifres8709 = False


}

ifres8708 = ifres8709


} else {
ifres8708 = False


}

var ifres8707 Obj

if True == ifres8708 {
ifres8707 = True


} else {
ifres8707 = False


}

ifres8706 = ifres8707


} else {
ifres8706 = False


}

var ifres8705 Obj

if True == ifres8706 {
ifres8705 = True


} else {
ifres8705 = False


}

ifres8704 = ifres8705


} else {
ifres8704 = False


}

if True == ifres8704 {
tmp8697 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8698 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8699 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__8534, tmp8698)
}
__typedArg0 := let__8534
__typedArg1 := tmp8698
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8700 := Call(__e, PrimFunc(symshen_4process_1assoc), tmp8699)


tmp8701 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8700, Nil)
}
__typedArg0 := tmp8700
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8702 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8697, tmp8701)
}
__typedArg0 := tmp8697
__typedArg1 := tmp8701
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__8534, tmp8702)
}
__typedArg0 := let__8534
__typedArg1 := tmp8702
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.TailApply(PrimFunc(symthaw), let__8533)
return
}


}, 0)

let__8536 := tmp8696
_ = let__8536

tmp8769 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4_8ch, let__8534)
}
__typedArg0 := symshen_4_8ch
__typedArg1 := let__8534
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp8769 {
tmp8767 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp8767 {
tmp8725 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__8537 := tmp8725
_ = let__8537

tmp8726 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8538 := tmp8726
_ = let__8538

tmp8765 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8537)
}
__typedArg0 := let__8537
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8741 Obj

if True == tmp8765 {
tmp8763 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8537)
}
__typedArg0 := let__8537
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8764 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp8763)
}
__typedArg0 := tmp8763
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8743 Obj

if True == tmp8764 {
tmp8760 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8537)
}
__typedArg0 := let__8537
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8761 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp8760)
}
__typedArg0 := tmp8760
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8762 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp8761)
}
__typedArg0 := tmp8761
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8745 Obj

if True == tmp8762 {
tmp8756 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8537)
}
__typedArg0 := let__8537
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8757 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp8756)
}
__typedArg0 := tmp8756
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8758 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp8757)
}
__typedArg0 := tmp8757
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8759 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp8758)
}
__typedArg0 := Nil
__typedArg1 := tmp8758
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8747 Obj

if True == tmp8759 {
tmp8755 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, let__8538)
}
__typedArg0 := Nil
__typedArg1 := let__8538
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8749 Obj

if True == tmp8755 {
tmp8751 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8537)
}
__typedArg0 := let__8537
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8752 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp8751)
}
__typedArg0 := tmp8751
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8753 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp8754 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp8752, tmp8753)
}
__typedArg0 := tmp8752
__typedArg1 := tmp8753
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8750 Obj

if True == tmp8754 {
ifres8750 = True


} else {
ifres8750 = False


}

ifres8749 = ifres8750


} else {
ifres8749 = False


}

var ifres8748 Obj

if True == ifres8749 {
ifres8748 = True


} else {
ifres8748 = False


}

ifres8747 = ifres8748


} else {
ifres8747 = False


}

var ifres8746 Obj

if True == ifres8747 {
ifres8746 = True


} else {
ifres8746 = False


}

ifres8745 = ifres8746


} else {
ifres8745 = False


}

var ifres8744 Obj

if True == ifres8745 {
ifres8744 = True


} else {
ifres8744 = False


}

ifres8743 = ifres8744


} else {
ifres8743 = False


}

var ifres8742 Obj

if True == ifres8743 {
ifres8742 = True


} else {
ifres8742 = False


}

ifres8741 = ifres8742


} else {
ifres8741 = False


}

if True == ifres8741 {
tmp8727 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8537)
}
__typedArg0 := let__8537
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8728 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8537)
}
__typedArg0 := let__8537
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8729 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp8728)
}
__typedArg0 := tmp8728
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8730 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8537)
}
__typedArg0 := let__8537
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8731 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp8730)
}
__typedArg0 := tmp8730
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8732 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_7, tmp8731)
}
__typedArg0 := sym_7
__typedArg1 := tmp8731
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8733 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8732, Nil)
}
__typedArg0 := tmp8732
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8734 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8729, tmp8733)
}
__typedArg0 := tmp8729
__typedArg1 := tmp8733
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8735 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8727, tmp8734)
}
__typedArg0 := tmp8727
__typedArg1 := tmp8734
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8736 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8735, Nil)
}
__typedArg0 := tmp8735
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8737 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1, tmp8736)
}
__typedArg0 := sym_1
__typedArg1 := tmp8736
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4cons_1form_1respect_1modes), tmp8737)
return


} else {
tmp8739 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, let__8538)
}
__typedArg0 := Nil
__typedArg1 := let__8538
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp8739 {
__e.TailApply(PrimFunc(symshen_4cons_1form_1respect_1modes), let__8537)
return
} else {
__e.TailApply(PrimFunc(symthaw), let__8536)
return
}


}


} else {
__e.TailApply(PrimFunc(symthaw), let__8536)
return
}


} else {
__e.TailApply(PrimFunc(symthaw), let__8536)
return
}


}


}


}


}


}


}


}


}


}


}


}


}


}


}


}


}


}


}


}


}


}


}


}


}


}


}


} else {
__e.TailApply(PrimFunc(symthaw), let__8533)
return
}


}, 1)

tmp8930 := Call(__e, ns2_1set, symshen_4macros, tmp8630)


_ = tmp8930

tmp8931 := MakeNative(func(__e *ControlFlow) {
V7107 := __e.Get(1)
_ = V7107
tmp8932 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4process_1input_7)
return
}, 0)

let__8539 := tmp8932
_ = let__8539

tmp8956 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7107)
}
__typedArg0 := V7107
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp8956 {
tmp8933 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7107)
}
__typedArg0 := V7107
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8540 := tmp8933
_ = let__8540

tmp8953 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7107)
}
__typedArg0 := V7107
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8954 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(syminput_7, tmp8953)
}
__typedArg0 := syminput_7
__typedArg1 := tmp8953
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp8954 {
tmp8951 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8540)
}
__typedArg0 := let__8540
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp8951 {
tmp8934 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8540)
}
__typedArg0 := let__8540
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__8541 := tmp8934
_ = let__8541

tmp8935 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8540)
}
__typedArg0 := let__8540
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8542 := tmp8935
_ = let__8542

tmp8949 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, let__8542)
}
__typedArg0 := Nil
__typedArg1 := let__8542
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp8949 {
tmp8936 := Call(__e, PrimFunc(symshen_4rcons__form), let__8541)


tmp8937 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstinput, Nil)
}
__typedArg0 := symstinput
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8938 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8937, Nil)
}
__typedArg0 := tmp8937
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8939 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8936, tmp8938)
}
__typedArg0 := tmp8936
__typedArg1 := tmp8938
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4input_1h_7, tmp8939)
}
__typedArg0 := symshen_4input_1h_7
__typedArg1 := tmp8939
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp8947 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8542)
}
__typedArg0 := let__8542
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8943 Obj

if True == tmp8947 {
tmp8945 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8542)
}
__typedArg0 := let__8542
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8946 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp8945)
}
__typedArg0 := Nil
__typedArg1 := tmp8945
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8944 Obj

if True == tmp8946 {
ifres8944 = True


} else {
ifres8944 = False


}

ifres8943 = ifres8944


} else {
ifres8943 = False


}

if True == ifres8943 {
tmp8940 := Call(__e, PrimFunc(symshen_4rcons__form), let__8541)


tmp8941 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8940, let__8542)
}
__typedArg0 := tmp8940
__typedArg1 := let__8542
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4input_1h_7, tmp8941)
}
__typedArg0 := symshen_4input_1h_7
__typedArg1 := tmp8941
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.TailApply(PrimFunc(symthaw), let__8539)
return
}


}


} else {
__e.TailApply(PrimFunc(symthaw), let__8539)
return
}


} else {
__e.TailApply(PrimFunc(symthaw), let__8539)
return
}


} else {
__e.TailApply(PrimFunc(symthaw), let__8539)
return
}


}, 1)

tmp8957 := Call(__e, ns2_1set, symshen_4process_1input_7, tmp8931)


_ = tmp8957

tmp8958 := MakeNative(func(__e *ControlFlow) {
V7114 := __e.Get(1)
_ = V7114
tmp8959 := MakeNative(func(__e *ControlFlow) {
__e.Return(V7114)
return
}, 0)

let__8543 := tmp8959
_ = let__8543

tmp8991 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7114)
}
__typedArg0 := V7114
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp8991 {
tmp8960 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7114)
}
__typedArg0 := V7114
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__8544 := tmp8960
_ = let__8544

tmp8961 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7114)
}
__typedArg0 := V7114
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8545 := tmp8961
_ = let__8545

tmp8989 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_7, let__8544)
}
__typedArg0 := sym_7
__typedArg1 := let__8544
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8982 Obj

if True == tmp8989 {
tmp8988 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8545)
}
__typedArg0 := let__8545
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8984 Obj

if True == tmp8988 {
tmp8986 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8545)
}
__typedArg0 := let__8545
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8987 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp8986)
}
__typedArg0 := Nil
__typedArg1 := tmp8986
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8985 Obj

if True == tmp8987 {
ifres8985 = True


} else {
ifres8985 = False


}

ifres8984 = ifres8985


} else {
ifres8984 = False


}

var ifres8983 Obj

if True == ifres8984 {
ifres8983 = True


} else {
ifres8983 = False


}

ifres8982 = ifres8983


} else {
ifres8982 = False


}

if True == ifres8982 {
tmp8962 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8545)
}
__typedArg0 := let__8545
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8963 := Call(__e, PrimFunc(symshen_4cons_1form_1respect_1modes), tmp8962)


tmp8964 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8963, Nil)
}
__typedArg0 := tmp8963
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_7, tmp8964)
}
__typedArg0 := sym_7
__typedArg1 := tmp8964
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp8980 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_1, let__8544)
}
__typedArg0 := sym_1
__typedArg1 := let__8544
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8973 Obj

if True == tmp8980 {
tmp8979 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8545)
}
__typedArg0 := let__8545
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8975 Obj

if True == tmp8979 {
tmp8977 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8545)
}
__typedArg0 := let__8545
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8978 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp8977)
}
__typedArg0 := Nil
__typedArg1 := tmp8977
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8976 Obj

if True == tmp8978 {
ifres8976 = True


} else {
ifres8976 = False


}

ifres8975 = ifres8976


} else {
ifres8975 = False


}

var ifres8974 Obj

if True == ifres8975 {
ifres8974 = True


} else {
ifres8974 = False


}

ifres8973 = ifres8974


} else {
ifres8973 = False


}

if True == ifres8973 {
tmp8965 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8545)
}
__typedArg0 := let__8545
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8966 := Call(__e, PrimFunc(symshen_4cons_1form_1respect_1modes), tmp8965)


tmp8967 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8966, Nil)
}
__typedArg0 := tmp8966
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1, tmp8967)
}
__typedArg0 := sym_1
__typedArg1 := tmp8967
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp8968 := Call(__e, PrimFunc(symshen_4cons_1form_1respect_1modes), let__8544)


tmp8969 := Call(__e, PrimFunc(symshen_4cons_1form_1respect_1modes), let__8545)


tmp8970 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8969, Nil)
}
__typedArg0 := tmp8969
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8971 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8968, tmp8970)
}
__typedArg0 := tmp8968
__typedArg1 := tmp8970
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp8971)
}
__typedArg0 := symcons
__typedArg1 := tmp8971
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


}


} else {
__e.TailApply(PrimFunc(symthaw), let__8543)
return
}


}, 1)

tmp8992 := Call(__e, ns2_1set, symshen_4cons_1form_1respect_1modes, tmp8958)


_ = tmp8992

tmp8993 := MakeNative(func(__e *ControlFlow) {
V7118 := __e.Get(1)
_ = V7118
V7119 := __e.Get(2)
_ = V7119
tmp8994 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symX, Nil)
}
__typedArg0 := symX
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8995 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_6, tmp8994)
}
__typedArg0 := sym_1_6
__typedArg1 := tmp8994
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8996 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symX, tmp8995)
}
__typedArg0 := symX
__typedArg1 := tmp8995
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__8546 := tmp8996
_ = let__8546

tmp8997 := Call(__e, PrimFunc(symappend), V7119, let__8546)


tmp8998 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V7118, tmp8997)
}
__typedArg0 := V7118
__typedArg1 := tmp8997
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8999 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefine, tmp8998)
}
__typedArg0 := symdefine
__typedArg1 := tmp8998
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9000 := Call(__e, PrimFunc(symeval), tmp8999)


let__8547 := tmp9000
_ = let__8547

tmp9001 := Call(__e, PrimFunc(symfn), V7118)


tmp9002 := Call(__e, PrimFunc(symshen_4record_1macro), V7118, tmp9001)


let__8548 := tmp9002
_ = let__8548

__e.Return(V7118)
return


}, 2)

tmp9003 := Call(__e, ns2_1set, symshen_4process_1def, tmp8993)


_ = tmp9003

tmp9004 := MakeNative(func(__e *ControlFlow) {
V7123 := __e.Get(1)
_ = V7123
tmp9044 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7123)
}
__typedArg0 := V7123
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9018 Obj

if True == tmp9044 {
tmp9042 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7123)
}
__typedArg0 := V7123
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9043 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlet, tmp9042)
}
__typedArg0 := symlet
__typedArg1 := tmp9042
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9020 Obj

if True == tmp9043 {
tmp9040 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7123)
}
__typedArg0 := V7123
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9041 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9040)
}
__typedArg0 := tmp9040
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9022 Obj

if True == tmp9041 {
tmp9037 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7123)
}
__typedArg0 := V7123
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9038 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9037)
}
__typedArg0 := tmp9037
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9039 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9038)
}
__typedArg0 := tmp9038
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9024 Obj

if True == tmp9039 {
tmp9033 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7123)
}
__typedArg0 := V7123
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9034 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9033)
}
__typedArg0 := tmp9033
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9035 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9034)
}
__typedArg0 := tmp9034
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9036 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9035)
}
__typedArg0 := tmp9035
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9026 Obj

if True == tmp9036 {
tmp9028 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7123)
}
__typedArg0 := V7123
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9029 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9028)
}
__typedArg0 := tmp9028
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9030 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9029)
}
__typedArg0 := tmp9029
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9031 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9030)
}
__typedArg0 := tmp9030
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9032 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9031)
}
__typedArg0 := tmp9031
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9027 Obj

if True == tmp9032 {
ifres9027 = True


} else {
ifres9027 = False


}

ifres9026 = ifres9027


} else {
ifres9026 = False


}

var ifres9025 Obj

if True == ifres9026 {
ifres9025 = True


} else {
ifres9025 = False


}

ifres9024 = ifres9025


} else {
ifres9024 = False


}

var ifres9023 Obj

if True == ifres9024 {
ifres9023 = True


} else {
ifres9023 = False


}

ifres9022 = ifres9023


} else {
ifres9022 = False


}

var ifres9021 Obj

if True == ifres9022 {
ifres9021 = True


} else {
ifres9021 = False


}

ifres9020 = ifres9021


} else {
ifres9020 = False


}

var ifres9019 Obj

if True == ifres9020 {
ifres9019 = True


} else {
ifres9019 = False


}

ifres9018 = ifres9019


} else {
ifres9018 = False


}

if True == ifres9018 {
tmp9005 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7123)
}
__typedArg0 := V7123
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9006 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9005)
}
__typedArg0 := tmp9005
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9007 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7123)
}
__typedArg0 := V7123
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9008 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9007)
}
__typedArg0 := tmp9007
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9009 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9008)
}
__typedArg0 := tmp9008
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9010 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7123)
}
__typedArg0 := V7123
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9011 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9010)
}
__typedArg0 := tmp9010
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9012 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9011)
}
__typedArg0 := tmp9011
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9013 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp9012)
}
__typedArg0 := symlet
__typedArg1 := tmp9012
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9014 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9013, Nil)
}
__typedArg0 := tmp9013
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9015 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9009, tmp9014)
}
__typedArg0 := tmp9009
__typedArg1 := tmp9014
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9016 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9006, tmp9015)
}
__typedArg0 := tmp9006
__typedArg1 := tmp9015
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp9016)
}
__typedArg0 := symlet
__typedArg1 := tmp9016
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return(V7123)
return
}


}, 1)

tmp9045 := Call(__e, ns2_1set, symshen_4process_1let, tmp9004)


_ = tmp9045

tmp9046 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__selftop:
V7124 := __self1
_ = V7124
tmp9047 := MakeNative(func(__e *ControlFlow) {
__e.Return(V7124)
return
}, 0)

let__8549 := tmp9047
_ = let__8549

tmp9077 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7124)
}
__typedArg0 := V7124
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9077 {
tmp9048 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7124)
}
__typedArg0 := V7124
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8550 := tmp9048
_ = let__8550

tmp9074 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7124)
}
__typedArg0 := V7124
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9075 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_8s, tmp9074)
}
__typedArg0 := sym_8s
__typedArg1 := tmp9074
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9075 {
tmp9072 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8550)
}
__typedArg0 := let__8550
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9072 {
tmp9049 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8550)
}
__typedArg0 := let__8550
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__8551 := tmp9049
_ = let__8551

tmp9050 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8550)
}
__typedArg0 := let__8550
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8552 := tmp9050
_ = let__8552

tmp9070 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8552)
}
__typedArg0 := let__8552
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9070 {
tmp9051 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8552)
}
__typedArg0 := let__8552
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8553 := tmp9051
_ = let__8553

tmp9068 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8553)
}
__typedArg0 := let__8553
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9068 {
tmp9052 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8s, let__8552)
}
__typedArg0 := sym_8s
__typedArg1 := let__8552
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9053 := Call(__e, PrimFunc(symshen_4process_1_8s), tmp9052)


tmp9054 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9053, Nil)
}
__typedArg0 := tmp9053
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9055 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__8551, tmp9054)
}
__typedArg0 := let__8551
__typedArg1 := tmp9054
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8s, tmp9055)
}
__typedArg0 := sym_8s
__typedArg1 := tmp9055
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9066 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, let__8553)
}
__typedArg0 := Nil
__typedArg1 := let__8553
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9063 Obj

if True == tmp9066 {
tmp9065 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_2) {
return PrimIsString(let__8551)
}
__typedArg0 := let__8551
return Call(__e, PrimFunc(symstring_2), __typedArg0)
})()

var ifres9064 Obj

if True == tmp9065 {
ifres9064 = True


} else {
ifres9064 = False


}

ifres9063 = ifres9064


} else {
ifres9063 = False


}

if True == ifres9063 {
tmp9056 := Call(__e, PrimFunc(symexplode), let__8551)


let__8554 := tmp9056
_ = let__8554

tmp9060 := Call(__e, PrimFunc(symlength), let__8554)


if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_6) {
__typedN0, __typedOK0 := TypedFloat64(tmp9060)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_6) {
return TypedMaterializeBoolean((__typedN0 > __typedN1))
}}
__typedArg0 := tmp9060
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_6), __typedArg0, __typedArg1)
})() {
tmp9057 := Call(__e, PrimFunc(symappend), let__8554, let__8552)


tmp9058 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8s, tmp9057)
}
__typedArg0 := sym_8s
__typedArg1 := tmp9057
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

if PrimFunc(symshen_4process_1_8s) == __self {
__self1 = tmp9058
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4process_1_8s), tmp9058)
return


} else {
__e.Return(V7124)
return
}


} else {
__e.TailApply(PrimFunc(symthaw), let__8549)
return
}


}


} else {
__e.TailApply(PrimFunc(symthaw), let__8549)
return
}


} else {
__e.TailApply(PrimFunc(symthaw), let__8549)
return
}


} else {
__e.TailApply(PrimFunc(symthaw), let__8549)
return
}


} else {
__e.TailApply(PrimFunc(symthaw), let__8549)
return
}


}, 1)

tmp9078 := Call(__e, ns2_1set, symshen_4process_1_8s, tmp9046)


_ = tmp9078

tmp9079 := MakeNative(func(__e *ControlFlow) {
V7134 := __e.Get(1)
_ = V7134
V7135 := __e.Get(2)
_ = V7135
tmp9080 := Call(__e, PrimFunc(symshen_4intern_1type), V7134)


let__8555 := tmp9080
_ = let__8555

tmp9081 := MakeNative(func(__e *ControlFlow) {
Z7138 := __e.Get(1)
_ = Z7138
__e.TailApply(PrimFunc(symshen_4_5datatype_6), Z7138)
return
}, 1)

tmp9082 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__8555, V7135)
}
__typedArg0 := let__8555
__typedArg1 := V7135
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9083 := Call(__e, PrimFunc(symcompile), tmp9081, tmp9082)


let__8556 := tmp9083
_ = let__8556

__e.Return(let__8555)
return


}, 2)

tmp9084 := Call(__e, ns2_1set, symshen_4process_1datatype, tmp9079)


_ = tmp9084

tmp9085 := MakeNative(func(__e *ControlFlow) {
V7139 := __e.Get(1)
_ = V7139
tmp9086 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(V7139)
}
__typedArg0 := V7139
return Call(__e, PrimFunc(symstr), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(tmp9086)
__typedS1, __typedOK1 := TypedString(MakeString("#type"))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := tmp9086
__typedArg1 := MakeString("#type")
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(tmp9086)
__typedS1, __typedOK1 := TypedString(MakeString("#type"))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := tmp9086
__typedArg1 := MakeString("#type")
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symintern), __typedArg0)
})())
return


}, 1)

tmp9088 := Call(__e, ns2_1set, symshen_4intern_1type, tmp9085)


_ = tmp9088

tmp9089 := MakeNative(func(__e *ControlFlow) {
V7140 := __e.Get(1)
_ = V7140
tmp9090 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsynonyms_d)
}
__typedArg0 := symshen_4_dsynonyms_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp9091 := Call(__e, PrimFunc(symappend), V7140, tmp9090)


tmp9092 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsynonyms_d, tmp9091)
}
__typedArg0 := symshen_4_dsynonyms_d
__typedArg1 := tmp9091
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4synonyms_1h), tmp9092)
return


}, 1)

tmp9093 := Call(__e, ns2_1set, symshen_4process_1synonyms, tmp9089)


_ = tmp9093

tmp9094 := MakeNative(func(__e *ControlFlow) {
V7143 := __e.Get(1)
_ = V7143
tmp9144 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7143)
}
__typedArg0 := V7143
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9105 Obj

if True == tmp9144 {
tmp9142 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7143)
}
__typedArg0 := V7143
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9143 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefun, tmp9142)
}
__typedArg0 := symdefun
__typedArg1 := tmp9142
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9107 Obj

if True == tmp9143 {
tmp9140 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7143)
}
__typedArg0 := V7143
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9141 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9140)
}
__typedArg0 := tmp9140
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9109 Obj

if True == tmp9141 {
tmp9137 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7143)
}
__typedArg0 := V7143
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9138 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9137)
}
__typedArg0 := tmp9137
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9139 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9138)
}
__typedArg0 := tmp9138
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9111 Obj

if True == tmp9139 {
tmp9133 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7143)
}
__typedArg0 := V7143
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9134 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9133)
}
__typedArg0 := tmp9133
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9135 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9134)
}
__typedArg0 := tmp9134
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9136 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9135)
}
__typedArg0 := tmp9135
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9113 Obj

if True == tmp9136 {
tmp9128 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7143)
}
__typedArg0 := V7143
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9129 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9128)
}
__typedArg0 := tmp9128
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9130 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9129)
}
__typedArg0 := tmp9129
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9131 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9130)
}
__typedArg0 := tmp9130
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9132 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp9131)
}
__typedArg0 := Nil
__typedArg1 := tmp9131
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9115 Obj

if True == tmp9132 {
tmp9124 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7143)
}
__typedArg0 := V7143
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9125 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9124)
}
__typedArg0 := tmp9124
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9126 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9125)
}
__typedArg0 := tmp9125
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9127 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9126)
}
__typedArg0 := tmp9126
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9117 Obj

if True == tmp9127 {
tmp9119 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7143)
}
__typedArg0 := V7143
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9120 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9119)
}
__typedArg0 := tmp9119
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9121 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9120)
}
__typedArg0 := tmp9120
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9122 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9121)
}
__typedArg0 := tmp9121
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9123 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp9122)
}
__typedArg0 := Nil
__typedArg1 := tmp9122
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9118 Obj

if True == tmp9123 {
ifres9118 = True


} else {
ifres9118 = False


}

ifres9117 = ifres9118


} else {
ifres9117 = False


}

var ifres9116 Obj

if True == ifres9117 {
ifres9116 = True


} else {
ifres9116 = False


}

ifres9115 = ifres9116


} else {
ifres9115 = False


}

var ifres9114 Obj

if True == ifres9115 {
ifres9114 = True


} else {
ifres9114 = False


}

ifres9113 = ifres9114


} else {
ifres9113 = False


}

var ifres9112 Obj

if True == ifres9113 {
ifres9112 = True


} else {
ifres9112 = False


}

ifres9111 = ifres9112


} else {
ifres9111 = False


}

var ifres9110 Obj

if True == ifres9111 {
ifres9110 = True


} else {
ifres9110 = False


}

ifres9109 = ifres9110


} else {
ifres9109 = False


}

var ifres9108 Obj

if True == ifres9109 {
ifres9108 = True


} else {
ifres9108 = False


}

ifres9107 = ifres9108


} else {
ifres9107 = False


}

var ifres9106 Obj

if True == ifres9107 {
ifres9106 = True


} else {
ifres9106 = False


}

ifres9105 = ifres9106


} else {
ifres9105 = False


}

if True == ifres9105 {
tmp9095 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7143)
}
__typedArg0 := V7143
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9096 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9095)
}
__typedArg0 := tmp9095
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9097 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9096)
}
__typedArg0 := tmp9096
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9098 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9097)
}
__typedArg0 := tmp9097
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9099 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7143)
}
__typedArg0 := V7143
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9100 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9099)
}
__typedArg0 := tmp9099
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9101 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9100)
}
__typedArg0 := tmp9100
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9102 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9098, tmp9101)
}
__typedArg0 := tmp9098
__typedArg1 := tmp9101
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9103 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_c_4, tmp9102)
}
__typedArg0 := sym_c_4
__typedArg1 := tmp9102
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symeval), tmp9103)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4lambda_1of_1defun)
return
}


}, 1)

tmp9145 := Call(__e, ns2_1set, symshen_4lambda_1of_1defun, tmp9094)


_ = tmp9145

tmp9146 := MakeNative(func(__e *ControlFlow) {
V7144 := __e.Get(1)
_ = V7144
tmp9147 := MakeNative(func(__e *ControlFlow) {
Z7146 := __e.Get(1)
_ = Z7146
__e.TailApply(PrimFunc(symshen_4curry_1type), Z7146)
return
}, 1)

tmp9148 := Call(__e, PrimFunc(symmap), tmp9147, V7144)


let__8557 := tmp9148
_ = let__8557

tmp9149 := Call(__e, PrimFunc(symshen_4compile_1synonyms), let__8557)


tmp9150 := Call(__e, PrimFunc(symshen_4shendef_1_6kldef), symshen_4demod, tmp9149)


tmp9151 := Call(__e, PrimFunc(symshen_4lambda_1of_1defun), tmp9150)


let__8558 := tmp9151
_ = let__8558

tmp9152 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_ddemodulation_1function_d, let__8558)
}
__typedArg0 := symshen_4_ddemodulation_1function_d
__typedArg1 := let__8558
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

let__8559 := tmp9152
_ = let__8559

__e.Return(symsynonyms)
return


}, 1)

tmp9153 := Call(__e, ns2_1set, symshen_4synonyms_1h, tmp9146)


_ = tmp9153

tmp9154 := MakeNative(func(__e *ControlFlow) {
V7151 := __e.Get(1)
_ = V7151
tmp9175 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V7151)
}
__typedArg0 := Nil
__typedArg1 := V7151
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9175 {
tmp9155 := Call(__e, PrimFunc(symgensym), symX)


let__8560 := tmp9155
_ = let__8560

tmp9156 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__8560, Nil)
}
__typedArg0 := let__8560
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9157 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_6, tmp9156)
}
__typedArg0 := sym_1_6
__typedArg1 := tmp9156
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__8560, tmp9157)
}
__typedArg0 := let__8560
__typedArg1 := tmp9157
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9173 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7151)
}
__typedArg0 := V7151
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9169 Obj

if True == tmp9173 {
tmp9171 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7151)
}
__typedArg0 := V7151
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9172 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9171)
}
__typedArg0 := tmp9171
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9170 Obj

if True == tmp9172 {
ifres9170 = True


} else {
ifres9170 = False


}

ifres9169 = ifres9170


} else {
ifres9169 = False


}

if True == ifres9169 {
tmp9158 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7151)
}
__typedArg0 := V7151
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9159 := Call(__e, PrimFunc(symshen_4rcons__form), tmp9158)


tmp9160 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7151)
}
__typedArg0 := V7151
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9161 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9160)
}
__typedArg0 := tmp9160
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9162 := Call(__e, PrimFunc(symshen_4rcons__form), tmp9161)


tmp9163 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7151)
}
__typedArg0 := V7151
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9164 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9163)
}
__typedArg0 := tmp9163
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9165 := Call(__e, PrimFunc(symshen_4compile_1synonyms), tmp9164)


tmp9166 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9162, tmp9165)
}
__typedArg0 := tmp9162
__typedArg1 := tmp9165
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9167 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_6, tmp9166)
}
__typedArg0 := sym_1_6
__typedArg1 := tmp9166
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9159, tmp9167)
}
__typedArg0 := tmp9159
__typedArg1 := tmp9167
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("synonyms requires an even number of arguments\n"))
}
__typedArg0 := MakeString("synonyms requires an even number of arguments\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 1)

tmp9176 := Call(__e, ns2_1set, symshen_4compile_1synonyms, tmp9154)


_ = tmp9176

tmp9177 := MakeNative(func(__e *ControlFlow) {
V7153 := __e.Get(1)
_ = V7153
tmp9178 := MakeNative(func(__e *ControlFlow) {
__e.Return(V7153)
return
}, 0)

let__8561 := tmp9178
_ = let__8561

tmp9202 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7153)
}
__typedArg0 := V7153
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9202 {
tmp9179 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7153)
}
__typedArg0 := V7153
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8562 := tmp9179
_ = let__8562

tmp9199 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7153)
}
__typedArg0 := V7153
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9200 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_c_4, tmp9199)
}
__typedArg0 := sym_c_4
__typedArg1 := tmp9199
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9200 {
tmp9197 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8562)
}
__typedArg0 := let__8562
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9197 {
tmp9180 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8562)
}
__typedArg0 := let__8562
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__8563 := tmp9180
_ = let__8563

tmp9181 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8562)
}
__typedArg0 := let__8562
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8564 := tmp9181
_ = let__8564

tmp9195 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8564)
}
__typedArg0 := let__8564
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9195 {
tmp9182 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8564)
}
__typedArg0 := let__8564
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8565 := tmp9182
_ = let__8565

tmp9193 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8565)
}
__typedArg0 := let__8565
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9193 {
tmp9183 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_c_4, let__8564)
}
__typedArg0 := sym_c_4
__typedArg1 := let__8564
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9184 := Call(__e, PrimFunc(symshen_4process_1lambda), tmp9183)


tmp9185 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9184, Nil)
}
__typedArg0 := tmp9184
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9186 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__8563, tmp9185)
}
__typedArg0 := let__8563
__typedArg1 := tmp9185
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp9186)
}
__typedArg0 := symlambda
__typedArg1 := tmp9186
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9191 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, let__8565)
}
__typedArg0 := Nil
__typedArg1 := let__8565
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9191 {
tmp9189 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(let__8563)
}
__typedArg0 := let__8563
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

if True == tmp9189 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, let__8562)
}
__typedArg0 := symlambda
__typedArg1 := let__8562
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return
} else {
tmp9187 := Call(__e, PrimFunc(symshen_4app), let__8563, MakeString(" is not a variable\n"), symshen_4s)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(tmp9187)
}
__typedArg0 := tmp9187
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}


} else {
__e.TailApply(PrimFunc(symthaw), let__8561)
return
}


}


} else {
__e.TailApply(PrimFunc(symthaw), let__8561)
return
}


} else {
__e.TailApply(PrimFunc(symthaw), let__8561)
return
}


} else {
__e.TailApply(PrimFunc(symthaw), let__8561)
return
}


} else {
__e.TailApply(PrimFunc(symthaw), let__8561)
return
}


}, 1)

tmp9203 := Call(__e, ns2_1set, symshen_4process_1lambda, tmp9177)


_ = tmp9203

tmp9204 := MakeNative(func(__e *ControlFlow) {
V7164 := __e.Get(1)
_ = V7164
tmp9205 := MakeNative(func(__e *ControlFlow) {
__e.Return(V7164)
return
}, 0)

let__8566 := tmp9205
_ = let__8566

tmp9239 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7164)
}
__typedArg0 := V7164
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9239 {
tmp9206 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7164)
}
__typedArg0 := V7164
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8567 := tmp9206
_ = let__8567

tmp9236 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7164)
}
__typedArg0 := V7164
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9237 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcases, tmp9236)
}
__typedArg0 := symcases
__typedArg1 := tmp9236
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9237 {
tmp9234 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8567)
}
__typedArg0 := let__8567
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9234 {
tmp9207 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8567)
}
__typedArg0 := let__8567
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__8568 := tmp9207
_ = let__8568

tmp9208 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8567)
}
__typedArg0 := let__8567
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8569 := tmp9208
_ = let__8569

tmp9232 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(True, let__8568)
}
__typedArg0 := True
__typedArg1 := let__8568
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9229 Obj

if True == tmp9232 {
tmp9231 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8569)
}
__typedArg0 := let__8569
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9230 Obj

if True == tmp9231 {
ifres9230 = True


} else {
ifres9230 = False


}

ifres9229 = ifres9230


} else {
ifres9229 = False


}

if True == ifres9229 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8569)
}
__typedArg0 := let__8569
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return
} else {
tmp9209 := MakeNative(func(__e *ControlFlow) {
tmp9211 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, let__8569)
}
__typedArg0 := Nil
__typedArg1 := let__8569
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9211 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("error: odd number of case elements\n"))
}
__typedArg0 := MakeString("error: odd number of case elements\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
} else {
__e.TailApply(PrimFunc(symthaw), let__8566)
return
}


}, 0)

let__8570 := tmp9209
_ = let__8570

tmp9227 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8569)
}
__typedArg0 := let__8569
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9227 {
tmp9212 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8569)
}
__typedArg0 := let__8569
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__8571 := tmp9212
_ = let__8571

tmp9213 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8569)
}
__typedArg0 := let__8569
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8572 := tmp9213
_ = let__8572

tmp9225 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, let__8572)
}
__typedArg0 := Nil
__typedArg1 := let__8572
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9225 {
tmp9214 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeString("error: cases exhausted"), Nil)
}
__typedArg0 := MakeString("error: cases exhausted")
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9215 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsimple_1error, tmp9214)
}
__typedArg0 := symsimple_1error
__typedArg1 := tmp9214
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9216 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9215, Nil)
}
__typedArg0 := tmp9215
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9217 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__8571, tmp9216)
}
__typedArg0 := let__8571
__typedArg1 := tmp9216
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9218 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__8568, tmp9217)
}
__typedArg0 := let__8568
__typedArg1 := tmp9217
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp9218)
}
__typedArg0 := symif
__typedArg1 := tmp9218
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9219 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcases, let__8572)
}
__typedArg0 := symcases
__typedArg1 := let__8572
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9220 := Call(__e, PrimFunc(symshen_4process_1cases), tmp9219)


tmp9221 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9220, Nil)
}
__typedArg0 := tmp9220
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9222 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__8571, tmp9221)
}
__typedArg0 := let__8571
__typedArg1 := tmp9221
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9223 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__8568, tmp9222)
}
__typedArg0 := let__8568
__typedArg1 := tmp9222
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp9223)
}
__typedArg0 := symif
__typedArg1 := tmp9223
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


} else {
__e.TailApply(PrimFunc(symthaw), let__8570)
return
}


}


} else {
__e.TailApply(PrimFunc(symthaw), let__8566)
return
}


} else {
__e.TailApply(PrimFunc(symthaw), let__8566)
return
}


} else {
__e.TailApply(PrimFunc(symthaw), let__8566)
return
}


}, 1)

tmp9240 := Call(__e, ns2_1set, symshen_4process_1cases, tmp9204)


_ = tmp9240

tmp9241 := MakeNative(func(__e *ControlFlow) {
V7174 := __e.Get(1)
_ = V7174
tmp9242 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symrun, Nil)
}
__typedArg0 := symrun
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9243 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symget_1time, tmp9242)
}
__typedArg0 := symget_1time
__typedArg1 := tmp9242
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9244 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symrun, Nil)
}
__typedArg0 := symrun
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9245 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symget_1time, tmp9244)
}
__typedArg0 := symget_1time
__typedArg1 := tmp9244
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9246 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symStart, Nil)
}
__typedArg0 := symStart
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9247 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symFinish, tmp9246)
}
__typedArg0 := symFinish
__typedArg1 := tmp9246
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9248 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1, tmp9247)
}
__typedArg0 := sym_1
__typedArg1 := tmp9247
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9249 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symTime, Nil)
}
__typedArg0 := symTime
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9250 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp9249)
}
__typedArg0 := symstr
__typedArg1 := tmp9249
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9251 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeString(" secs\n"), Nil)
}
__typedArg0 := MakeString(" secs\n")
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9252 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9250, tmp9251)
}
__typedArg0 := tmp9250
__typedArg1 := tmp9251
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9253 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcn, tmp9252)
}
__typedArg0 := symcn
__typedArg1 := tmp9252
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9254 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9253, Nil)
}
__typedArg0 := tmp9253
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9255 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeString("\nrun time: "), tmp9254)
}
__typedArg0 := MakeString("\nrun time: ")
__typedArg1 := tmp9254
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9256 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcn, tmp9255)
}
__typedArg0 := symcn
__typedArg1 := tmp9255
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9257 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstoutput, Nil)
}
__typedArg0 := symstoutput
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9258 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9257, Nil)
}
__typedArg0 := tmp9257
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9259 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9256, tmp9258)
}
__typedArg0 := tmp9256
__typedArg1 := tmp9258
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9260 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympr, tmp9259)
}
__typedArg0 := sympr
__typedArg1 := tmp9259
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9261 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symResult, Nil)
}
__typedArg0 := symResult
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9262 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9260, tmp9261)
}
__typedArg0 := tmp9260
__typedArg1 := tmp9261
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9263 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symMessage, tmp9262)
}
__typedArg0 := symMessage
__typedArg1 := tmp9262
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9264 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9248, tmp9263)
}
__typedArg0 := tmp9248
__typedArg1 := tmp9263
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9265 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symTime, tmp9264)
}
__typedArg0 := symTime
__typedArg1 := tmp9264
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9266 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9245, tmp9265)
}
__typedArg0 := tmp9245
__typedArg1 := tmp9265
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9267 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symFinish, tmp9266)
}
__typedArg0 := symFinish
__typedArg1 := tmp9266
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9268 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V7174, tmp9267)
}
__typedArg0 := V7174
__typedArg1 := tmp9267
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9269 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symResult, tmp9268)
}
__typedArg0 := symResult
__typedArg1 := tmp9268
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9270 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9243, tmp9269)
}
__typedArg0 := tmp9243
__typedArg1 := tmp9269
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9271 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symStart, tmp9270)
}
__typedArg0 := symStart
__typedArg1 := tmp9270
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp9271)
}
__typedArg0 := symlet
__typedArg1 := tmp9271
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp9272 := Call(__e, ns2_1set, symshen_4process_1time, tmp9241)


_ = tmp9272

tmp9273 := MakeNative(func(__e *ControlFlow) {
V7175 := __e.Get(1)
_ = V7175
tmp9299 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7175)
}
__typedArg0 := V7175
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9284 Obj

if True == tmp9299 {
tmp9297 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7175)
}
__typedArg0 := V7175
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9298 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9297)
}
__typedArg0 := tmp9297
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9286 Obj

if True == tmp9298 {
tmp9294 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7175)
}
__typedArg0 := V7175
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9295 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9294)
}
__typedArg0 := tmp9294
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9296 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9295)
}
__typedArg0 := tmp9295
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9288 Obj

if True == tmp9296 {
tmp9290 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7175)
}
__typedArg0 := V7175
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9291 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9290)
}
__typedArg0 := tmp9290
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9292 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9291)
}
__typedArg0 := tmp9291
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9293 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9292)
}
__typedArg0 := tmp9292
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9289 Obj

if True == tmp9293 {
ifres9289 = True


} else {
ifres9289 = False


}

ifres9288 = ifres9289


} else {
ifres9288 = False


}

var ifres9287 Obj

if True == ifres9288 {
ifres9287 = True


} else {
ifres9287 = False


}

ifres9286 = ifres9287


} else {
ifres9286 = False


}

var ifres9285 Obj

if True == ifres9286 {
ifres9285 = True


} else {
ifres9285 = False


}

ifres9284 = ifres9285


} else {
ifres9284 = False


}

if True == ifres9284 {
tmp9274 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7175)
}
__typedArg0 := V7175
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9275 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7175)
}
__typedArg0 := V7175
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9276 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9275)
}
__typedArg0 := tmp9275
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9277 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7175)
}
__typedArg0 := V7175
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9278 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7175)
}
__typedArg0 := V7175
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9279 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9278)
}
__typedArg0 := tmp9278
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9280 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9277, tmp9279)
}
__typedArg0 := tmp9277
__typedArg1 := tmp9279
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9281 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9280, Nil)
}
__typedArg0 := tmp9280
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9282 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9276, tmp9281)
}
__typedArg0 := tmp9276
__typedArg1 := tmp9281
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9274, tmp9282)
}
__typedArg0 := tmp9274
__typedArg1 := tmp9282
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return(V7175)
return
}


}, 1)

tmp9300 := Call(__e, ns2_1set, symshen_4process_1assoc, tmp9273)


_ = tmp9300

tmp9301 := MakeNative(func(__e *ControlFlow) {
V7176 := __e.Get(1)
_ = V7176
tmp9302 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(V7176)
}
__typedArg0 := V7176
return Call(__e, PrimFunc(symstr), __typedArg0)
})()

tmp9303 := Call(__e, PrimFunc(symshen_4mu_1h), tmp9302)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(tmp9303)
}
__typedArg0 := tmp9303
return Call(__e, PrimFunc(symintern), __typedArg0)
})())
return


}, 1)

tmp9304 := Call(__e, ns2_1set, symshen_4make_1uppercase, tmp9301)


_ = tmp9304

tmp9305 := MakeNative(func(__e *ControlFlow) {
V7177 := __e.Get(1)
_ = V7177
tmp9321 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString(""), V7177)
}
__typedArg0 := MakeString("")
__typedArg1 := V7177
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9321 {
__e.Return(MakeString(""))
return
} else {
tmp9319 := Call(__e, PrimFunc(symshen_4_7string_2), V7177)


if True == tmp9319 {
tmp9306 := Call(__e, PrimFunc(symhdstr), V7177)


tmp9307 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_1_6n) {
return PrimStringToNumber(tmp9306)
}
__typedArg0 := tmp9306
return Call(__e, PrimFunc(symstring_1_6n), __typedArg0)
})()

let__8573 := tmp9307
_ = let__8573

let__8574 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(let__8573)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(32))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := let__8573
__typedArg1 := MakeInteger(32)
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})()
_ = let__8574

tmp9315 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_6_a) {
__typedN0, __typedOK0 := TypedFloat64(let__8573)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(97))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_6_a) {
return TypedMaterializeBoolean((__typedN0 >= __typedN1))
}}
__typedArg0 := let__8573
__typedArg1 := MakeInteger(97)
return Call(__e, PrimFunc(sym_6_a), __typedArg0, __typedArg1)
})()

var ifres9312 Obj

if True == tmp9315 {
tmp9314 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_a) {
__typedN0, __typedOK0 := TypedFloat64(let__8573)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(122))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_5_a) {
return TypedMaterializeBoolean((__typedN0 <= __typedN1))
}}
__typedArg0 := let__8573
__typedArg1 := MakeInteger(122)
return Call(__e, PrimFunc(sym_5_a), __typedArg0, __typedArg1)
})()

var ifres9313 Obj

if True == tmp9314 {
ifres9313 = True


} else {
ifres9313 = False


}

ifres9312 = ifres9313


} else {
ifres9312 = False


}

var ifres9309 Obj

if True == ifres9312 {
tmp9310 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symn_1_6string) {
return PrimNumberToString(let__8574)
}
__typedArg0 := let__8574
return Call(__e, PrimFunc(symn_1_6string), __typedArg0)
})()

ifres9309 = tmp9310


} else {
tmp9311 := Call(__e, PrimFunc(symhdstr), V7177)


ifres9309 = tmp9311


}

let__8575 := ifres9309
_ = let__8575

tmp9317 := Call(__e, PrimFunc(symshen_4mu_1h), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V7177)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V7177
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


__e.TailApply(PrimFunc(sym_8s), let__8575, tmp9317)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4mu_1h)
return
}


}


}, 1)

tmp9322 := Call(__e, ns2_1set, symshen_4mu_1h, tmp9305)


_ = tmp9322

tmp9323 := MakeNative(func(__e *ControlFlow) {
V7181 := __e.Get(1)
_ = V7181
V7182 := __e.Get(2)
_ = V7182
tmp9324 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dmacros_d)
}
__typedArg0 := sym_dmacros_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp9325 := Call(__e, PrimFunc(symshen_4update_1assoc), V7181, V7182, tmp9324)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(sym_dmacros_d, tmp9325)
}
__typedArg0 := sym_dmacros_d
__typedArg1 := tmp9325
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return


}, 2)

tmp9326 := Call(__e, ns2_1set, symshen_4record_1macro, tmp9323)


_ = tmp9326

tmp9327 := MakeNative(func(__e *ControlFlow) {
V7192 := __e.Get(1)
_ = V7192
V7193 := __e.Get(2)
_ = V7193
V7194 := __e.Get(3)
_ = V7194
tmp9344 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V7194)
}
__typedArg0 := Nil
__typedArg1 := V7194
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9344 {
tmp9328 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V7192, V7193)
}
__typedArg0 := V7192
__typedArg1 := V7193
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9328, Nil)
}
__typedArg0 := tmp9328
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9329 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.update-assoc"))
}
__typedArg0 := MakeString("implementation error in shen.update-assoc")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}, 0)

let__8576 := tmp9329
_ = let__8576

tmp9342 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7194)
}
__typedArg0 := V7194
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9342 {
tmp9330 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7194)
}
__typedArg0 := V7194
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__8577 := tmp9330
_ = let__8577

tmp9331 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7194)
}
__typedArg0 := V7194
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8578 := tmp9331
_ = let__8578

tmp9340 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8577)
}
__typedArg0 := let__8577
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9336 Obj

if True == tmp9340 {
tmp9338 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8577)
}
__typedArg0 := let__8577
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9339 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V7192, tmp9338)
}
__typedArg0 := V7192
__typedArg1 := tmp9338
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9337 Obj

if True == tmp9339 {
ifres9337 = True


} else {
ifres9337 = False


}

ifres9336 = ifres9337


} else {
ifres9336 = False


}

if True == ifres9336 {
tmp9332 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8577)
}
__typedArg0 := let__8577
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9333 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9332, V7193)
}
__typedArg0 := tmp9332
__typedArg1 := V7193
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9333, let__8578)
}
__typedArg0 := tmp9333
__typedArg1 := let__8578
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9334 := Call(__e, PrimFunc(symshen_4update_1assoc), V7192, V7193, let__8578)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__8577, tmp9334)
}
__typedArg0 := let__8577
__typedArg1 := tmp9334
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


} else {
__e.TailApply(PrimFunc(symthaw), let__8576)
return
}


}


}, 3)

tmp9345 := Call(__e, ns2_1set, symshen_4update_1assoc, tmp9327)


_ = tmp9345

tmp9346 := MakeNative(func(__e *ControlFlow) {
tmp9354 := Call(__e, PrimFunc(symstinput))


tmp9355 := Call(__e, PrimFunc(symshen_4char_1stinput_2), tmp9354)


if True == tmp9355 {
tmp9347 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstinput, Nil)
}
__typedArg0 := symstinput
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9348 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9347, Nil)
}
__typedArg0 := tmp9347
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9349 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4read_1unit_1string, tmp9348)
}
__typedArg0 := symshen_4read_1unit_1string
__typedArg1 := tmp9348
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9350 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9349, Nil)
}
__typedArg0 := tmp9349
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring_1_6n, tmp9350)
}
__typedArg0 := symstring_1_6n
__typedArg1 := tmp9350
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9351 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstinput, Nil)
}
__typedArg0 := symstinput
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9352 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9351, Nil)
}
__typedArg0 := tmp9351
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1byte, tmp9352)
}
__typedArg0 := symread_1byte
__typedArg1 := tmp9352
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


}, 0)

tmp9356 := Call(__e, ns2_1set, symshen_4process_1read_1byte, tmp9346)


_ = tmp9356

tmp9357 := MakeNative(func(__e *ControlFlow) {
V7198 := __e.Get(1)
_ = V7198
tmp9358 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4prolog_1vector, Nil)
}
__typedArg0 := symshen_4prolog_1vector
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__8579 := tmp9358
_ = let__8579

tmp9359 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), Nil)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9360 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp9359)
}
__typedArg0 := symvector
__typedArg1 := tmp9359
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9361 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9360, Nil)
}
__typedArg0 := tmp9360
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9362 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp9361)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp9361
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9363 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(True, tmp9362)
}
__typedArg0 := True
__typedArg1 := tmp9362
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9364 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8v, tmp9363)
}
__typedArg0 := sym_8v
__typedArg1 := tmp9363
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__8580 := tmp9364
_ = let__8580

let__8581 := MakeInteger(0)
_ = let__8581

tmp9365 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(True, Nil)
}
__typedArg0 := True
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9366 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfreeze, tmp9365)
}
__typedArg0 := symfreeze
__typedArg1 := tmp9365
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__8582 := tmp9366
_ = let__8582

tmp9367 := MakeNative(func(__e *ControlFlow) {
Z7204 := __e.Get(1)
_ = Z7204
__e.TailApply(PrimFunc(symshen_4_5body_6), Z7204)
return
}, 1)

tmp9368 := Call(__e, PrimFunc(symcompile), tmp9367, V7198)


let__8583 := tmp9368
_ = let__8583

tmp9369 := Call(__e, PrimFunc(symshen_4received), V7198)


let__8584 := tmp9369
_ = let__8584

tmp9370 := Call(__e, PrimFunc(symgensym), symV)


let__8585 := tmp9370
_ = let__8585

tmp9371 := Call(__e, PrimFunc(symgensym), symL)


let__8586 := tmp9371
_ = let__8586

tmp9372 := Call(__e, PrimFunc(symgensym), symK)


let__8587 := tmp9372
_ = let__8587

tmp9373 := Call(__e, PrimFunc(symgensym), symC)


let__8588 := tmp9373
_ = let__8588

tmp9374 := Call(__e, PrimFunc(symshen_4continue), let__8584, let__8583, let__8585, let__8586, let__8587, let__8588)


tmp9375 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9374, Nil)
}
__typedArg0 := tmp9374
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9376 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__8588, tmp9375)
}
__typedArg0 := let__8588
__typedArg1 := tmp9375
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9377 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp9376)
}
__typedArg0 := symlambda
__typedArg1 := tmp9376
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9378 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9377, Nil)
}
__typedArg0 := tmp9377
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9379 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__8587, tmp9378)
}
__typedArg0 := let__8587
__typedArg1 := tmp9378
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9380 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp9379)
}
__typedArg0 := symlambda
__typedArg1 := tmp9379
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9381 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9380, Nil)
}
__typedArg0 := tmp9380
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9382 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__8586, tmp9381)
}
__typedArg0 := let__8586
__typedArg1 := tmp9381
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9383 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp9382)
}
__typedArg0 := symlambda
__typedArg1 := tmp9382
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9384 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9383, Nil)
}
__typedArg0 := tmp9383
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9385 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__8585, tmp9384)
}
__typedArg0 := let__8585
__typedArg1 := tmp9384
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9386 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp9385)
}
__typedArg0 := symlambda
__typedArg1 := tmp9385
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__8589 := tmp9386
_ = let__8589

tmp9387 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__8582, Nil)
}
__typedArg0 := let__8582
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9388 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__8581, tmp9387)
}
__typedArg0 := let__8581
__typedArg1 := tmp9387
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9389 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__8580, tmp9388)
}
__typedArg0 := let__8580
__typedArg1 := tmp9388
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9390 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__8579, tmp9389)
}
__typedArg0 := let__8579
__typedArg1 := tmp9389
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__8589, tmp9390)
}
__typedArg0 := let__8589
__typedArg1 := tmp9390
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp9391 := Call(__e, ns2_1set, symshen_4call_1prolog, tmp9357)


_ = tmp9391

tmp9392 := MakeNative(func(__e *ControlFlow) {
V7213 := __e.Get(1)
_ = V7213
tmp9393 := MakeNative(func(__e *ControlFlow) {
__e.Return(Nil)
return
}, 0)

let__8590 := tmp9393
_ = let__8590

tmp9408 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7213)
}
__typedArg0 := V7213
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9408 {
tmp9394 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7213)
}
__typedArg0 := V7213
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__8591 := tmp9394
_ = let__8591

tmp9395 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7213)
}
__typedArg0 := V7213
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8592 := tmp9395
_ = let__8592

tmp9406 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symreceive, let__8591)
}
__typedArg0 := symreceive
__typedArg1 := let__8591
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9399 Obj

if True == tmp9406 {
tmp9405 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8592)
}
__typedArg0 := let__8592
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9401 Obj

if True == tmp9405 {
tmp9403 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8592)
}
__typedArg0 := let__8592
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9404 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp9403)
}
__typedArg0 := Nil
__typedArg1 := tmp9403
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9402 Obj

if True == tmp9404 {
ifres9402 = True


} else {
ifres9402 = False


}

ifres9401 = ifres9402


} else {
ifres9401 = False


}

var ifres9400 Obj

if True == ifres9401 {
ifres9400 = True


} else {
ifres9400 = False


}

ifres9399 = ifres9400


} else {
ifres9399 = False


}

if True == ifres9399 {
__e.Return(let__8592)
return
} else {
tmp9396 := Call(__e, PrimFunc(symshen_4received), let__8591)


tmp9397 := Call(__e, PrimFunc(symshen_4received), let__8592)


__e.TailApply(PrimFunc(symunion), tmp9396, tmp9397)
return


}


} else {
__e.TailApply(PrimFunc(symthaw), let__8590)
return
}


}, 1)

tmp9409 := Call(__e, ns2_1set, symshen_4received, tmp9392)


_ = tmp9409

tmp9410 := MakeNative(func(__e *ControlFlow) {
tmp9411 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dprolog_1memory_d)
}
__typedArg0 := symshen_4_dprolog_1memory_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp9412 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symabsvector) {
return PrimAbsvector(tmp9411)
}
__typedArg0 := tmp9411
return Call(__e, PrimFunc(symabsvector), __typedArg0)
})()

let__8593 := tmp9412
_ = let__8593

tmp9413 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(let__8593, MakeInteger(0), symshen_4print_1prolog_1vector)
}
__typedArg0 := let__8593
__typedArg1 := MakeInteger(0)
__typedArg2 := symshen_4print_1prolog_1vector
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

let__8594 := tmp9413
_ = let__8594

tmp9414 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(let__8593, MakeInteger(1), MakeInteger(2))
}
__typedArg0 := let__8593
__typedArg1 := MakeInteger(1)
__typedArg2 := MakeInteger(2)
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

let__8595 := tmp9414
_ = let__8595

__e.Return(let__8595)
return


}, 0)

tmp9415 := Call(__e, ns2_1set, symshen_4prolog_1vector, tmp9410)


_ = tmp9415

tmp9416 := MakeNative(func(__e *ControlFlow) {
V7220 := __e.Get(1)
_ = V7220
__e.Return(V7220)
return
}, 1)

tmp9417 := Call(__e, ns2_1set, symreceive, tmp9416)


_ = tmp9417

tmp9418 := MakeNative(func(__e *ControlFlow) {
V7221 := __e.Get(1)
_ = V7221
tmp9426 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7221)
}
__typedArg0 := V7221
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9426 {
tmp9419 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7221)
}
__typedArg0 := V7221
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9420 := Call(__e, PrimFunc(symshen_4rcons__form), tmp9419)


tmp9421 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7221)
}
__typedArg0 := V7221
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9422 := Call(__e, PrimFunc(symshen_4rcons__form), tmp9421)


tmp9423 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9422, Nil)
}
__typedArg0 := tmp9422
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9424 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9420, tmp9423)
}
__typedArg0 := tmp9420
__typedArg1 := tmp9423
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp9424)
}
__typedArg0 := symcons
__typedArg1 := tmp9424
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return(V7221)
return
}


}, 1)

tmp9427 := Call(__e, ns2_1set, symshen_4rcons__form, tmp9418)


_ = tmp9427

tmp9428 := MakeNative(func(__e *ControlFlow) {
V7222 := __e.Get(1)
_ = V7222
tmp9435 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7222)
}
__typedArg0 := V7222
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9435 {
tmp9429 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7222)
}
__typedArg0 := V7222
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9430 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7222)
}
__typedArg0 := V7222
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9431 := Call(__e, PrimFunc(symshen_4tuple_1up), tmp9430)


tmp9432 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9431, Nil)
}
__typedArg0 := tmp9431
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9433 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9429, tmp9432)
}
__typedArg0 := tmp9429
__typedArg1 := tmp9432
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8p, tmp9433)
}
__typedArg0 := sym_8p
__typedArg1 := tmp9433
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return(V7222)
return
}


}, 1)

tmp9436 := Call(__e, ns2_1set, symshen_4tuple_1up, tmp9428)


_ = tmp9436

tmp9437 := MakeNative(func(__e *ControlFlow) {
V7223 := __e.Get(1)
_ = V7223
tmp9438 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dmacros_d)
}
__typedArg0 := sym_dmacros_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp9439 := Call(__e, PrimFunc(symassoc), V7223, tmp9438)


tmp9440 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dmacros_d)
}
__typedArg0 := sym_dmacros_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp9441 := Call(__e, PrimFunc(symremove), tmp9439, tmp9440)


tmp9442 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(sym_dmacros_d, tmp9441)
}
__typedArg0 := sym_dmacros_d
__typedArg1 := tmp9441
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp9442

__e.Return(V7223)
return


}, 1)

__e.TailApply(ns2_1set, symundefmacro, tmp9437)
return




}, 0)

