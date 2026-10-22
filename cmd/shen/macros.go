package main

import . "github.com/pyrex41/shen-go/kl"

var MacrosMain = MakeNative(func(__e *ControlFlow) {
tmp8622 := MakeNative(func(__e *ControlFlow) {
V7129 := __e.Get(1)
_ = V7129
tmp8623 := MakeNative(func(__e *ControlFlow) {
Z7131 := __e.Get(1)
_ = Z7131
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Z7131)
}
__typedArg0 := Z7131
return Call(__e, PrimFunc(symtl), __typedArg0)
})())
return
}, 1)

tmp8624 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dmacros_d)
}
__typedArg0 := sym_dmacros_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp8625 := Call(__e, PrimFunc(symmap), tmp8623, tmp8624)


let__8531 := tmp8625
_ = let__8531

tmp8626 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsource_1form_1handlers_d)
}
__typedArg0 := symshen_4_dsource_1form_1handlers_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4macroexpand_1h), V7129, let__8531, let__8531, tmp8626)
return


}, 1)

tmp8627 := Call(__e, ns2_1set, symmacroexpand, tmp8622)


_ = tmp8627

tmp8628 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__self3 := __e.Get(3)
__self4 := __e.Get(4)
__selftop:
V7144 := __self1
_ = V7144
V7145 := __self2
_ = V7145
V7146 := __self3
_ = V7146
V7147 := __self4
_ = V7147
tmp8637 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V7145)
}
__typedArg0 := Nil
__typedArg1 := V7145
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp8637 {
__e.Return(V7144)
return
} else {
tmp8635 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7145)
}
__typedArg0 := V7145
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp8635 {
tmp8629 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7145)
}
__typedArg0 := V7145
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8630 := Call(__e, PrimFunc(symshen_4source_1macro_1walk), tmp8629, V7144, V7147)


let__8532 := tmp8630
_ = let__8532

tmp8633 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V7144, let__8532)
}
__typedArg0 := V7144
__typedArg1 := let__8532
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp8633 {
tmp8631 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7145)
}
__typedArg0 := V7145
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4macroexpand_1h) == __self {
__self1, __self2, __self3, __self4 = V7144, tmp8631, V7146, V7147
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4macroexpand_1h), V7144, tmp8631, V7146, V7147)
return


} else {
if PrimFunc(symshen_4macroexpand_1h) == __self {
__self1, __self2, __self3, __self4 = let__8532, V7146, V7146, V7147
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4macroexpand_1h), let__8532, V7146, V7146, V7147)
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

tmp8638 := Call(__e, ns2_1set, symshen_4macroexpand_1h, tmp8628)


_ = tmp8638

tmp8639 := MakeNative(func(__e *ControlFlow) {
V7152 := __e.Get(1)
_ = V7152
V7153 := __e.Get(2)
_ = V7153
V7154 := __e.Get(3)
_ = V7154
tmp8648 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V7154)
}
__typedArg0 := Nil
__typedArg1 := V7154
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp8648 {
__e.TailApply(PrimFunc(symshen_4walk), V7152, V7153)
return
} else {
tmp8645 := Call(__e, PrimFunc(symshen_4source_1form_1handler), V7153, V7154)


tmp8646 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp8645)
}
__typedArg0 := tmp8645
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp8646 {
__e.Return(V7153)
return
} else {
tmp8643 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7153)
}
__typedArg0 := V7153
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp8643 {
tmp8640 := MakeNative(func(__e *ControlFlow) {
Z7155 := __e.Get(1)
_ = Z7155
__e.TailApply(PrimFunc(symshen_4source_1macro_1walk), V7152, Z7155, V7154)
return
}, 1)

tmp8641 := Call(__e, PrimFunc(symmap), tmp8640, V7153)


__e.TailApply(V7152, tmp8641)
return


} else {
__e.TailApply(V7152, V7153)
return
}


}


}


}, 3)

tmp8649 := Call(__e, ns2_1set, symshen_4source_1macro_1walk, tmp8639)


_ = tmp8649

tmp8650 := MakeNative(func(__e *ControlFlow) {
V7156 := __e.Get(1)
_ = V7156
V7157 := __e.Get(2)
_ = V7157
tmp8654 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7157)
}
__typedArg0 := V7157
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp8654 {
tmp8651 := MakeNative(func(__e *ControlFlow) {
Z7158 := __e.Get(1)
_ = Z7158
__e.TailApply(PrimFunc(symshen_4walk), V7156, Z7158)
return
}, 1)

tmp8652 := Call(__e, PrimFunc(symmap), tmp8651, V7157)


__e.TailApply(V7156, tmp8652)
return


} else {
__e.TailApply(V7156, V7157)
return
}


}, 2)

tmp8655 := Call(__e, ns2_1set, symshen_4walk, tmp8650)


_ = tmp8655

tmp8656 := MakeNative(func(__e *ControlFlow) {
V7159 := __e.Get(1)
_ = V7159
tmp8657 := MakeNative(func(__e *ControlFlow) {
__e.Return(V7159)
return
}, 0)

let__8533 := tmp8657
_ = let__8533

tmp8943 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7159)
}
__typedArg0 := V7159
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp8943 {
tmp8658 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7159)
}
__typedArg0 := V7159
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__8534 := tmp8658
_ = let__8534

tmp8659 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7159)
}
__typedArg0 := V7159
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8535 := tmp8659
_ = let__8535

tmp8941 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefmacro, let__8534)
}
__typedArg0 := symdefmacro
__typedArg1 := let__8534
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8938 Obj

if True == tmp8941 {
tmp8940 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8939 Obj

if True == tmp8940 {
ifres8939 = True


} else {
ifres8939 = False


}

ifres8938 = ifres8939


} else {
ifres8938 = False


}

if True == ifres8938 {
tmp8660 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8661 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4process_1def), tmp8660, tmp8661)
return


} else {
tmp8936 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefcc, let__8534)
}
__typedArg0 := symdefcc
__typedArg1 := let__8534
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp8936 {
__e.TailApply(PrimFunc(symshen_4yacc_1_6shen), let__8535)
return
} else {
tmp8934 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symu_b, let__8534)
}
__typedArg0 := symu_b
__typedArg1 := let__8534
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8927 Obj

if True == tmp8934 {
tmp8933 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8929 Obj

if True == tmp8933 {
tmp8931 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8932 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp8931)
}
__typedArg0 := Nil
__typedArg1 := tmp8931
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8930 Obj

if True == tmp8932 {
ifres8930 = True


} else {
ifres8930 = False


}

ifres8929 = ifres8930


} else {
ifres8929 = False


}

var ifres8928 Obj

if True == ifres8929 {
ifres8928 = True


} else {
ifres8928 = False


}

ifres8927 = ifres8928


} else {
ifres8927 = False


}

if True == ifres8927 {
tmp8662 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8663 := Call(__e, PrimFunc(symshen_4make_1uppercase), tmp8662)


tmp8664 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8663, Nil)
}
__typedArg0 := tmp8663
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symprotect, tmp8664)
}
__typedArg0 := symprotect
__typedArg1 := tmp8664
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp8925 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symerror, let__8534)
}
__typedArg0 := symerror
__typedArg1 := let__8534
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8922 Obj

if True == tmp8925 {
tmp8924 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8923 Obj

if True == tmp8924 {
ifres8923 = True


} else {
ifres8923 = False


}

ifres8922 = ifres8923


} else {
ifres8922 = False


}

if True == ifres8922 {
tmp8665 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8666 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8667 := Call(__e, PrimFunc(symshen_4mkstr), tmp8665, tmp8666)


tmp8668 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8667, Nil)
}
__typedArg0 := tmp8667
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsimple_1error, tmp8668)
}
__typedArg0 := symsimple_1error
__typedArg1 := tmp8668
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp8920 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symoutput, let__8534)
}
__typedArg0 := symoutput
__typedArg1 := let__8534
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8917 Obj

if True == tmp8920 {
tmp8919 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8918 Obj

if True == tmp8919 {
ifres8918 = True


} else {
ifres8918 = False


}

ifres8917 = ifres8918


} else {
ifres8917 = False


}

if True == ifres8917 {
tmp8669 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8670 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8671 := Call(__e, PrimFunc(symshen_4mkstr), tmp8669, tmp8670)


tmp8672 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstoutput, Nil)
}
__typedArg0 := symstoutput
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8673 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8672, Nil)
}
__typedArg0 := tmp8672
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8674 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8671, tmp8673)
}
__typedArg0 := tmp8671
__typedArg1 := tmp8673
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympr, tmp8674)
}
__typedArg0 := sympr
__typedArg1 := tmp8674
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp8915 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sympr, let__8534)
}
__typedArg0 := sympr
__typedArg1 := let__8534
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8908 Obj

if True == tmp8915 {
tmp8914 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8910 Obj

if True == tmp8914 {
tmp8912 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8913 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp8912)
}
__typedArg0 := Nil
__typedArg1 := tmp8912
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8911 Obj

if True == tmp8913 {
ifres8911 = True


} else {
ifres8911 = False


}

ifres8910 = ifres8911


} else {
ifres8910 = False


}

var ifres8909 Obj

if True == ifres8910 {
ifres8909 = True


} else {
ifres8909 = False


}

ifres8908 = ifres8909


} else {
ifres8908 = False


}

if True == ifres8908 {
tmp8675 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8676 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstoutput, Nil)
}
__typedArg0 := symstoutput
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8677 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8676, Nil)
}
__typedArg0 := tmp8676
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8678 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8675, tmp8677)
}
__typedArg0 := tmp8675
__typedArg1 := tmp8677
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympr, tmp8678)
}
__typedArg0 := sympr
__typedArg1 := tmp8678
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp8906 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symmake_1string, let__8534)
}
__typedArg0 := symmake_1string
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

__e.TailApply(PrimFunc(symshen_4mkstr), tmp8679, tmp8680)
return


} else {
tmp8901 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlineread, let__8534)
}
__typedArg0 := symlineread
__typedArg1 := let__8534
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8898 Obj

if True == tmp8901 {
tmp8900 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, let__8535)
}
__typedArg0 := Nil
__typedArg1 := let__8535
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8899 Obj

if True == tmp8900 {
ifres8899 = True


} else {
ifres8899 = False


}

ifres8898 = ifres8899


} else {
ifres8898 = False


}

if True == ifres8898 {
tmp8681 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstinput, Nil)
}
__typedArg0 := symstinput
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8682 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8681, Nil)
}
__typedArg0 := tmp8681
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlineread, tmp8682)
}
__typedArg0 := symlineread
__typedArg1 := tmp8682
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp8896 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(syminput, let__8534)
}
__typedArg0 := syminput
__typedArg1 := let__8534
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8893 Obj

if True == tmp8896 {
tmp8895 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, let__8535)
}
__typedArg0 := Nil
__typedArg1 := let__8535
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8894 Obj

if True == tmp8895 {
ifres8894 = True


} else {
ifres8894 = False


}

ifres8893 = ifres8894


} else {
ifres8893 = False


}

if True == ifres8893 {
tmp8683 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstinput, Nil)
}
__typedArg0 := symstinput
__typedArg1 := Nil
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

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminput, tmp8684)
}
__typedArg0 := syminput
__typedArg1 := tmp8684
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp8891 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symread, let__8534)
}
__typedArg0 := symread
__typedArg1 := let__8534
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8888 Obj

if True == tmp8891 {
tmp8890 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, let__8535)
}
__typedArg0 := Nil
__typedArg1 := let__8535
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8889 Obj

if True == tmp8890 {
ifres8889 = True


} else {
ifres8889 = False


}

ifres8888 = ifres8889


} else {
ifres8888 = False


}

if True == ifres8888 {
tmp8685 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstinput, Nil)
}
__typedArg0 := symstinput
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8686 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8685, Nil)
}
__typedArg0 := tmp8685
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread, tmp8686)
}
__typedArg0 := symread
__typedArg1 := tmp8686
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp8886 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(syminput_7, let__8534)
}
__typedArg0 := syminput_7
__typedArg1 := let__8534
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8883 Obj

if True == tmp8886 {
tmp8885 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8884 Obj

if True == tmp8885 {
ifres8884 = True


} else {
ifres8884 = False


}

ifres8883 = ifres8884


} else {
ifres8883 = False


}

if True == ifres8883 {
__e.TailApply(PrimFunc(symshen_4process_1input_7), V7159)
return
} else {
tmp8881 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symread_1byte, let__8534)
}
__typedArg0 := symread_1byte
__typedArg1 := let__8534
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8878 Obj

if True == tmp8881 {
tmp8880 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, let__8535)
}
__typedArg0 := Nil
__typedArg1 := let__8535
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8879 Obj

if True == tmp8880 {
ifres8879 = True


} else {
ifres8879 = False


}

ifres8878 = ifres8879


} else {
ifres8878 = False


}

if True == ifres8878 {
__e.TailApply(PrimFunc(symshen_4process_1read_1byte))
return
} else {
tmp8876 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symprolog_2, let__8534)
}
__typedArg0 := symprolog_2
__typedArg1 := let__8534
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp8876 {
__e.TailApply(PrimFunc(symshen_4call_1prolog), let__8535)
return
} else {
tmp8874 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefprolog, let__8534)
}
__typedArg0 := symdefprolog
__typedArg1 := let__8534
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8871 Obj

if True == tmp8874 {
tmp8873 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8872 Obj

if True == tmp8873 {
ifres8872 = True


} else {
ifres8872 = False


}

ifres8871 = ifres8872


} else {
ifres8871 = False


}

if True == ifres8871 {
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

__e.TailApply(PrimFunc(symshen_4compile_1prolog), tmp8687, tmp8688)
return


} else {
tmp8869 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdatatype, let__8534)
}
__typedArg0 := symdatatype
__typedArg1 := let__8534
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8866 Obj

if True == tmp8869 {
tmp8868 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8867 Obj

if True == tmp8868 {
ifres8867 = True


} else {
ifres8867 = False


}

ifres8866 = ifres8867


} else {
ifres8866 = False


}

if True == ifres8866 {
tmp8689 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8690 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4process_1datatype), tmp8689, tmp8690)
return


} else {
tmp8864 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_8s, let__8534)
}
__typedArg0 := sym_8s
__typedArg1 := let__8534
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp8864 {
__e.TailApply(PrimFunc(symshen_4process_1_8s), V7159)
return
} else {
tmp8862 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symsynonyms, let__8534)
}
__typedArg0 := symsynonyms
__typedArg1 := let__8534
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp8862 {
__e.TailApply(PrimFunc(symshen_4process_1synonyms), let__8535)
return
} else {
tmp8860 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symnl, let__8534)
}
__typedArg0 := symnl
__typedArg1 := let__8534
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8857 Obj

if True == tmp8860 {
tmp8859 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, let__8535)
}
__typedArg0 := Nil
__typedArg1 := let__8535
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
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
tmp8691 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), Nil)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnl, tmp8691)
}
__typedArg0 := symnl
__typedArg1 := tmp8691
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp8855 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlet, let__8534)
}
__typedArg0 := symlet
__typedArg1 := let__8534
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp8855 {
__e.TailApply(PrimFunc(symshen_4process_1let), V7159)
return
} else {
tmp8853 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_c_4, let__8534)
}
__typedArg0 := sym_c_4
__typedArg1 := let__8534
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp8853 {
__e.TailApply(PrimFunc(symshen_4process_1lambda), V7159)
return
} else {
tmp8851 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcases, let__8534)
}
__typedArg0 := symcases
__typedArg1 := let__8534
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp8851 {
__e.TailApply(PrimFunc(symshen_4process_1cases), V7159)
return
} else {
tmp8849 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symtime, let__8534)
}
__typedArg0 := symtime
__typedArg1 := let__8534
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8842 Obj

if True == tmp8849 {
tmp8848 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8844 Obj

if True == tmp8848 {
tmp8846 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8847 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp8846)
}
__typedArg0 := Nil
__typedArg1 := tmp8846
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8845 Obj

if True == tmp8847 {
ifres8845 = True


} else {
ifres8845 = False


}

ifres8844 = ifres8845


} else {
ifres8844 = False


}

var ifres8843 Obj

if True == ifres8844 {
ifres8843 = True


} else {
ifres8843 = False


}

ifres8842 = ifres8843


} else {
ifres8842 = False


}

if True == ifres8842 {
tmp8692 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4process_1time), tmp8692)
return


} else {
tmp8840 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symput, let__8534)
}
__typedArg0 := symput
__typedArg1 := let__8534
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8825 Obj

if True == tmp8840 {
tmp8839 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8827 Obj

if True == tmp8839 {
tmp8829 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8536 := tmp8829
_ = let__8536

tmp8838 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8536)
}
__typedArg0 := let__8536
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8830 Obj

if True == tmp8838 {
tmp8832 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8536)
}
__typedArg0 := let__8536
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8537 := tmp8832
_ = let__8537

tmp8837 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8537)
}
__typedArg0 := let__8537
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8833 Obj

if True == tmp8837 {
tmp8835 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8537)
}
__typedArg0 := let__8537
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8836 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp8835)
}
__typedArg0 := Nil
__typedArg1 := tmp8835
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8834 Obj

if True == tmp8836 {
ifres8834 = True


} else {
ifres8834 = False


}

ifres8833 = ifres8834


} else {
ifres8833 = False


}

var ifres8831 Obj

if True == ifres8833 {
ifres8831 = True


} else {
ifres8831 = False


}

ifres8830 = ifres8831


} else {
ifres8830 = False


}

var ifres8828 Obj

if True == ifres8830 {
ifres8828 = True


} else {
ifres8828 = False


}

ifres8827 = ifres8828


} else {
ifres8827 = False


}

var ifres8826 Obj

if True == ifres8827 {
ifres8826 = True


} else {
ifres8826 = False


}

ifres8825 = ifres8826


} else {
ifres8825 = False


}

if True == ifres8825 {
tmp8693 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8538 := tmp8693
_ = let__8538

tmp8694 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8695 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8538)
}
__typedArg0 := let__8538
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8696 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8538)
}
__typedArg0 := let__8538
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8697 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp8696)
}
__typedArg0 := tmp8696
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8698 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dproperty_1vector_d, Nil)
}
__typedArg0 := sym_dproperty_1vector_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8699 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvalue, tmp8698)
}
__typedArg0 := symvalue
__typedArg1 := tmp8698
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8700 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8699, Nil)
}
__typedArg0 := tmp8699
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8701 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8697, tmp8700)
}
__typedArg0 := tmp8697
__typedArg1 := tmp8700
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8702 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8695, tmp8701)
}
__typedArg0 := tmp8695
__typedArg1 := tmp8701
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8703 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8694, tmp8702)
}
__typedArg0 := tmp8694
__typedArg1 := tmp8702
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symput, tmp8703)
}
__typedArg0 := symput
__typedArg1 := tmp8703
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp8823 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symget, let__8534)
}
__typedArg0 := symget
__typedArg1 := let__8534
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8812 Obj

if True == tmp8823 {
tmp8822 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8535)
}
__typedArg0 := let__8535
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

let__8539 := tmp8816
_ = let__8539

tmp8821 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8539)
}
__typedArg0 := let__8539
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8817 Obj

if True == tmp8821 {
tmp8819 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8539)
}
__typedArg0 := let__8539
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8820 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp8819)
}
__typedArg0 := Nil
__typedArg1 := tmp8819
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8818 Obj

if True == tmp8820 {
ifres8818 = True


} else {
ifres8818 = False


}

ifres8817 = ifres8818


} else {
ifres8817 = False


}

var ifres8815 Obj

if True == ifres8817 {
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

if True == ifres8812 {
tmp8704 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8705 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8706 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp8705)
}
__typedArg0 := tmp8705
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8707 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dproperty_1vector_d, Nil)
}
__typedArg0 := sym_dproperty_1vector_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8708 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvalue, tmp8707)
}
__typedArg0 := symvalue
__typedArg1 := tmp8707
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8709 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8708, Nil)
}
__typedArg0 := tmp8708
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8710 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8706, tmp8709)
}
__typedArg0 := tmp8706
__typedArg1 := tmp8709
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8711 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8704, tmp8710)
}
__typedArg0 := tmp8704
__typedArg1 := tmp8710
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symget, tmp8711)
}
__typedArg0 := symget
__typedArg1 := tmp8711
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp8810 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symunput, let__8534)
}
__typedArg0 := symunput
__typedArg1 := let__8534
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8799 Obj

if True == tmp8810 {
tmp8809 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8801 Obj

if True == tmp8809 {
tmp8803 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8540 := tmp8803
_ = let__8540

tmp8808 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8540)
}
__typedArg0 := let__8540
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8804 Obj

if True == tmp8808 {
tmp8806 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8540)
}
__typedArg0 := let__8540
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8807 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp8806)
}
__typedArg0 := Nil
__typedArg1 := tmp8806
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8805 Obj

if True == tmp8807 {
ifres8805 = True


} else {
ifres8805 = False


}

ifres8804 = ifres8805


} else {
ifres8804 = False


}

var ifres8802 Obj

if True == ifres8804 {
ifres8802 = True


} else {
ifres8802 = False


}

ifres8801 = ifres8802


} else {
ifres8801 = False


}

var ifres8800 Obj

if True == ifres8801 {
ifres8800 = True


} else {
ifres8800 = False


}

ifres8799 = ifres8800


} else {
ifres8799 = False


}

if True == ifres8799 {
tmp8712 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8713 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8714 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp8713)
}
__typedArg0 := tmp8713
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8715 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dproperty_1vector_d, Nil)
}
__typedArg0 := sym_dproperty_1vector_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8716 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvalue, tmp8715)
}
__typedArg0 := symvalue
__typedArg1 := tmp8715
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8717 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8716, Nil)
}
__typedArg0 := tmp8716
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8718 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8714, tmp8717)
}
__typedArg0 := tmp8714
__typedArg1 := tmp8717
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8719 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8712, tmp8718)
}
__typedArg0 := tmp8712
__typedArg1 := tmp8718
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunput, tmp8719)
}
__typedArg0 := symunput
__typedArg1 := tmp8719
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp8797 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4_8c, let__8534)
}
__typedArg0 := symshen_4_8c
__typedArg1 := let__8534
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8790 Obj

if True == tmp8797 {
tmp8796 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8792 Obj

if True == tmp8796 {
tmp8794 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8795 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp8794)
}
__typedArg0 := Nil
__typedArg1 := tmp8794
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8793 Obj

if True == tmp8795 {
ifres8793 = True


} else {
ifres8793 = False


}

ifres8792 = ifres8793


} else {
ifres8792 = False


}

var ifres8791 Obj

if True == ifres8792 {
ifres8791 = True


} else {
ifres8791 = False


}

ifres8790 = ifres8791


} else {
ifres8790 = False


}

if True == ifres8790 {
tmp8720 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4rcons__form), tmp8720)
return


} else {
tmp8721 := MakeNative(func(__e *ControlFlow) {
tmp8748 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8729 Obj

if True == tmp8748 {
tmp8731 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8542 := tmp8731
_ = let__8542

tmp8747 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8542)
}
__typedArg0 := let__8542
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8732 Obj

if True == tmp8747 {
tmp8745 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8542)
}
__typedArg0 := let__8542
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8746 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp8745)
}
__typedArg0 := tmp8745
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8734 Obj

if True == tmp8746 {
tmp8736 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdo, Nil)
}
__typedArg0 := symdo
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8737 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_d, tmp8736)
}
__typedArg0 := sym_d
__typedArg1 := tmp8736
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8738 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_7, tmp8737)
}
__typedArg0 := sym_7
__typedArg1 := tmp8737
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8739 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symor, tmp8738)
}
__typedArg0 := symor
__typedArg1 := tmp8738
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8740 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symand, tmp8739)
}
__typedArg0 := symand
__typedArg1 := tmp8739
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8741 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symappend, tmp8740)
}
__typedArg0 := symappend
__typedArg1 := tmp8740
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8742 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8v, tmp8741)
}
__typedArg0 := sym_8v
__typedArg1 := tmp8741
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8743 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8p, tmp8742)
}
__typedArg0 := sym_8p
__typedArg1 := tmp8742
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8744 := Call(__e, PrimFunc(symelement_2), let__8534, tmp8743)


var ifres8735 Obj

if True == tmp8744 {
ifres8735 = True


} else {
ifres8735 = False


}

ifres8734 = ifres8735


} else {
ifres8734 = False


}

var ifres8733 Obj

if True == ifres8734 {
ifres8733 = True


} else {
ifres8733 = False


}

ifres8732 = ifres8733


} else {
ifres8732 = False


}

var ifres8730 Obj

if True == ifres8732 {
ifres8730 = True


} else {
ifres8730 = False


}

ifres8729 = ifres8730


} else {
ifres8729 = False


}

if True == ifres8729 {
tmp8722 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8723 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8724 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__8534, tmp8723)
}
__typedArg0 := let__8534
__typedArg1 := tmp8723
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8725 := Call(__e, PrimFunc(symshen_4process_1assoc), tmp8724)


tmp8726 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8725, Nil)
}
__typedArg0 := tmp8725
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8727 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8722, tmp8726)
}
__typedArg0 := tmp8722
__typedArg1 := tmp8726
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__8534, tmp8727)
}
__typedArg0 := let__8534
__typedArg1 := tmp8727
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.TailApply(PrimFunc(symthaw), let__8533)
return
}


}, 0)

let__8541 := tmp8721
_ = let__8541

tmp8788 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4_8ch, let__8534)
}
__typedArg0 := symshen_4_8ch
__typedArg1 := let__8534
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp8788 {
tmp8786 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp8786 {
tmp8749 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__8543 := tmp8749
_ = let__8543

tmp8750 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8535)
}
__typedArg0 := let__8535
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8544 := tmp8750
_ = let__8544

tmp8784 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8543)
}
__typedArg0 := let__8543
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8764 Obj

if True == tmp8784 {
tmp8766 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8543)
}
__typedArg0 := let__8543
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8545 := tmp8766
_ = let__8545

tmp8783 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8545)
}
__typedArg0 := let__8545
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8767 Obj

if True == tmp8783 {
tmp8769 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8545)
}
__typedArg0 := let__8545
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8546 := tmp8769
_ = let__8546

tmp8782 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8546)
}
__typedArg0 := let__8546
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8770 Obj

if True == tmp8782 {
tmp8780 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8546)
}
__typedArg0 := let__8546
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8781 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp8780)
}
__typedArg0 := Nil
__typedArg1 := tmp8780
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8772 Obj

if True == tmp8781 {
tmp8779 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, let__8544)
}
__typedArg0 := Nil
__typedArg1 := let__8544
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8774 Obj

if True == tmp8779 {
tmp8776 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8545)
}
__typedArg0 := let__8545
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8777 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp8778 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp8776, tmp8777)
}
__typedArg0 := tmp8776
__typedArg1 := tmp8777
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8775 Obj

if True == tmp8778 {
ifres8775 = True


} else {
ifres8775 = False


}

ifres8774 = ifres8775


} else {
ifres8774 = False


}

var ifres8773 Obj

if True == ifres8774 {
ifres8773 = True


} else {
ifres8773 = False


}

ifres8772 = ifres8773


} else {
ifres8772 = False


}

var ifres8771 Obj

if True == ifres8772 {
ifres8771 = True


} else {
ifres8771 = False


}

ifres8770 = ifres8771


} else {
ifres8770 = False


}

var ifres8768 Obj

if True == ifres8770 {
ifres8768 = True


} else {
ifres8768 = False


}

ifres8767 = ifres8768


} else {
ifres8767 = False


}

var ifres8765 Obj

if True == ifres8767 {
ifres8765 = True


} else {
ifres8765 = False


}

ifres8764 = ifres8765


} else {
ifres8764 = False


}

if True == ifres8764 {
tmp8751 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8543)
}
__typedArg0 := let__8543
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8547 := tmp8751
_ = let__8547

tmp8752 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8543)
}
__typedArg0 := let__8543
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8753 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8547)
}
__typedArg0 := let__8547
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8754 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8547)
}
__typedArg0 := let__8547
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8755 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_7, tmp8754)
}
__typedArg0 := sym_7
__typedArg1 := tmp8754
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8756 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8755, Nil)
}
__typedArg0 := tmp8755
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8757 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8753, tmp8756)
}
__typedArg0 := tmp8753
__typedArg1 := tmp8756
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8758 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8752, tmp8757)
}
__typedArg0 := tmp8752
__typedArg1 := tmp8757
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8759 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8758, Nil)
}
__typedArg0 := tmp8758
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8760 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1, tmp8759)
}
__typedArg0 := sym_1
__typedArg1 := tmp8759
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4cons_1form_1respect_1modes), tmp8760)
return


} else {
tmp8762 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, let__8544)
}
__typedArg0 := Nil
__typedArg1 := let__8544
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp8762 {
__e.TailApply(PrimFunc(symshen_4cons_1form_1respect_1modes), let__8543)
return
} else {
__e.TailApply(PrimFunc(symthaw), let__8541)
return
}


}


} else {
__e.TailApply(PrimFunc(symthaw), let__8541)
return
}


} else {
__e.TailApply(PrimFunc(symthaw), let__8541)
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

tmp8944 := Call(__e, ns2_1set, symshen_4macros, tmp8656)


_ = tmp8944

tmp8945 := MakeNative(func(__e *ControlFlow) {
V7176 := __e.Get(1)
_ = V7176
tmp8946 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4process_1input_7)
return
}, 0)

let__8548 := tmp8946
_ = let__8548

tmp8970 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7176)
}
__typedArg0 := V7176
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp8970 {
tmp8947 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7176)
}
__typedArg0 := V7176
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8549 := tmp8947
_ = let__8549

tmp8967 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7176)
}
__typedArg0 := V7176
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8968 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(syminput_7, tmp8967)
}
__typedArg0 := syminput_7
__typedArg1 := tmp8967
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp8968 {
tmp8965 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8549)
}
__typedArg0 := let__8549
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp8965 {
tmp8948 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8549)
}
__typedArg0 := let__8549
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__8550 := tmp8948
_ = let__8550

tmp8949 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8549)
}
__typedArg0 := let__8549
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8551 := tmp8949
_ = let__8551

tmp8963 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, let__8551)
}
__typedArg0 := Nil
__typedArg1 := let__8551
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp8963 {
tmp8950 := Call(__e, PrimFunc(symshen_4rcons__form), let__8550)


tmp8951 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstinput, Nil)
}
__typedArg0 := symstinput
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8952 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8951, Nil)
}
__typedArg0 := tmp8951
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8953 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8950, tmp8952)
}
__typedArg0 := tmp8950
__typedArg1 := tmp8952
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4input_1h_7, tmp8953)
}
__typedArg0 := symshen_4input_1h_7
__typedArg1 := tmp8953
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp8961 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8551)
}
__typedArg0 := let__8551
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8957 Obj

if True == tmp8961 {
tmp8959 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8551)
}
__typedArg0 := let__8551
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8960 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp8959)
}
__typedArg0 := Nil
__typedArg1 := tmp8959
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8958 Obj

if True == tmp8960 {
ifres8958 = True


} else {
ifres8958 = False


}

ifres8957 = ifres8958


} else {
ifres8957 = False


}

if True == ifres8957 {
tmp8954 := Call(__e, PrimFunc(symshen_4rcons__form), let__8550)


tmp8955 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8954, let__8551)
}
__typedArg0 := tmp8954
__typedArg1 := let__8551
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4input_1h_7, tmp8955)
}
__typedArg0 := symshen_4input_1h_7
__typedArg1 := tmp8955
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.TailApply(PrimFunc(symthaw), let__8548)
return
}


}


} else {
__e.TailApply(PrimFunc(symthaw), let__8548)
return
}


} else {
__e.TailApply(PrimFunc(symthaw), let__8548)
return
}


} else {
__e.TailApply(PrimFunc(symthaw), let__8548)
return
}


}, 1)

tmp8971 := Call(__e, ns2_1set, symshen_4process_1input_7, tmp8945)


_ = tmp8971

tmp8972 := MakeNative(func(__e *ControlFlow) {
V7183 := __e.Get(1)
_ = V7183
tmp8973 := MakeNative(func(__e *ControlFlow) {
__e.Return(V7183)
return
}, 0)

let__8552 := tmp8973
_ = let__8552

tmp9014 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7183)
}
__typedArg0 := V7183
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9014 {
tmp8974 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7183)
}
__typedArg0 := V7183
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__8553 := tmp8974
_ = let__8553

tmp8975 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7183)
}
__typedArg0 := V7183
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8554 := tmp8975
_ = let__8554

tmp9012 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_7, let__8553)
}
__typedArg0 := sym_7
__typedArg1 := let__8553
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9005 Obj

if True == tmp9012 {
tmp9011 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8554)
}
__typedArg0 := let__8554
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9007 Obj

if True == tmp9011 {
tmp9009 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8554)
}
__typedArg0 := let__8554
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9010 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp9009)
}
__typedArg0 := Nil
__typedArg1 := tmp9009
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9008 Obj

if True == tmp9010 {
ifres9008 = True


} else {
ifres9008 = False


}

ifres9007 = ifres9008


} else {
ifres9007 = False


}

var ifres9006 Obj

if True == ifres9007 {
ifres9006 = True


} else {
ifres9006 = False


}

ifres9005 = ifres9006


} else {
ifres9005 = False


}

if True == ifres9005 {
tmp8976 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8554)
}
__typedArg0 := let__8554
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8977 := Call(__e, PrimFunc(symshen_4cons_1form_1respect_1modes), tmp8976)


tmp8978 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8977, Nil)
}
__typedArg0 := tmp8977
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_7, tmp8978)
}
__typedArg0 := sym_7
__typedArg1 := tmp8978
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9003 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_1, let__8553)
}
__typedArg0 := sym_1
__typedArg1 := let__8553
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8996 Obj

if True == tmp9003 {
tmp9002 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8554)
}
__typedArg0 := let__8554
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8998 Obj

if True == tmp9002 {
tmp9000 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8554)
}
__typedArg0 := let__8554
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9001 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp9000)
}
__typedArg0 := Nil
__typedArg1 := tmp9000
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8999 Obj

if True == tmp9001 {
ifres8999 = True


} else {
ifres8999 = False


}

ifres8998 = ifres8999


} else {
ifres8998 = False


}

var ifres8997 Obj

if True == ifres8998 {
ifres8997 = True


} else {
ifres8997 = False


}

ifres8996 = ifres8997


} else {
ifres8996 = False


}

if True == ifres8996 {
tmp8979 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8554)
}
__typedArg0 := let__8554
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8980 := Call(__e, PrimFunc(symshen_4cons_1form_1respect_1modes), tmp8979)


tmp8981 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8980, Nil)
}
__typedArg0 := tmp8980
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1, tmp8981)
}
__typedArg0 := sym_1
__typedArg1 := tmp8981
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp8994 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symbar_b, let__8553)
}
__typedArg0 := symbar_b
__typedArg1 := let__8553
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8987 Obj

if True == tmp8994 {
tmp8993 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8554)
}
__typedArg0 := let__8554
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8989 Obj

if True == tmp8993 {
tmp8991 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8554)
}
__typedArg0 := let__8554
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8992 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp8991)
}
__typedArg0 := Nil
__typedArg1 := tmp8991
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8990 Obj

if True == tmp8992 {
ifres8990 = True


} else {
ifres8990 = False


}

ifres8989 = ifres8990


} else {
ifres8989 = False


}

var ifres8988 Obj

if True == ifres8989 {
ifres8988 = True


} else {
ifres8988 = False


}

ifres8987 = ifres8988


} else {
ifres8987 = False


}

if True == ifres8987 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8554)
}
__typedArg0 := let__8554
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return
} else {
tmp8982 := Call(__e, PrimFunc(symshen_4cons_1form_1respect_1modes), let__8553)


tmp8983 := Call(__e, PrimFunc(symshen_4cons_1form_1respect_1modes), let__8554)


tmp8984 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8983, Nil)
}
__typedArg0 := tmp8983
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8985 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8982, tmp8984)
}
__typedArg0 := tmp8982
__typedArg1 := tmp8984
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp8985)
}
__typedArg0 := symcons
__typedArg1 := tmp8985
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


}


}


} else {
__e.TailApply(PrimFunc(symthaw), let__8552)
return
}


}, 1)

tmp9015 := Call(__e, ns2_1set, symshen_4cons_1form_1respect_1modes, tmp8972)


_ = tmp9015

tmp9016 := MakeNative(func(__e *ControlFlow) {
V7187 := __e.Get(1)
_ = V7187
V7188 := __e.Get(2)
_ = V7188
tmp9017 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symX, Nil)
}
__typedArg0 := symX
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9018 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_6, tmp9017)
}
__typedArg0 := sym_1_6
__typedArg1 := tmp9017
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9019 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symX, tmp9018)
}
__typedArg0 := symX
__typedArg1 := tmp9018
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__8555 := tmp9019
_ = let__8555

tmp9020 := Call(__e, PrimFunc(symappend), V7188, let__8555)


tmp9021 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V7187, tmp9020)
}
__typedArg0 := V7187
__typedArg1 := tmp9020
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9022 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefine, tmp9021)
}
__typedArg0 := symdefine
__typedArg1 := tmp9021
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9023 := Call(__e, PrimFunc(symeval), tmp9022)


let__8556 := tmp9023
_ = let__8556

tmp9024 := Call(__e, PrimFunc(symfn), V7187)


tmp9025 := Call(__e, PrimFunc(symshen_4record_1macro), V7187, tmp9024)


let__8557 := tmp9025
_ = let__8557

__e.Return(V7187)
return


}, 2)

tmp9026 := Call(__e, ns2_1set, symshen_4process_1def, tmp9016)


_ = tmp9026

tmp9027 := MakeNative(func(__e *ControlFlow) {
V7192 := __e.Get(1)
_ = V7192
tmp9059 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7192)
}
__typedArg0 := V7192
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9039 Obj

if True == tmp9059 {
tmp9041 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7192)
}
__typedArg0 := V7192
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8558 := tmp9041
_ = let__8558

tmp9057 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7192)
}
__typedArg0 := V7192
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9058 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlet, tmp9057)
}
__typedArg0 := symlet
__typedArg1 := tmp9057
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9042 Obj

if True == tmp9058 {
tmp9056 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8558)
}
__typedArg0 := let__8558
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9044 Obj

if True == tmp9056 {
tmp9046 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8558)
}
__typedArg0 := let__8558
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8559 := tmp9046
_ = let__8559

tmp9055 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8559)
}
__typedArg0 := let__8559
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9047 Obj

if True == tmp9055 {
tmp9049 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8559)
}
__typedArg0 := let__8559
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8560 := tmp9049
_ = let__8560

tmp9054 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8560)
}
__typedArg0 := let__8560
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9050 Obj

if True == tmp9054 {
tmp9052 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8560)
}
__typedArg0 := let__8560
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9053 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9052)
}
__typedArg0 := tmp9052
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9051 Obj

if True == tmp9053 {
ifres9051 = True


} else {
ifres9051 = False


}

ifres9050 = ifres9051


} else {
ifres9050 = False


}

var ifres9048 Obj

if True == ifres9050 {
ifres9048 = True


} else {
ifres9048 = False


}

ifres9047 = ifres9048


} else {
ifres9047 = False


}

var ifres9045 Obj

if True == ifres9047 {
ifres9045 = True


} else {
ifres9045 = False


}

ifres9044 = ifres9045


} else {
ifres9044 = False


}

var ifres9043 Obj

if True == ifres9044 {
ifres9043 = True


} else {
ifres9043 = False


}

ifres9042 = ifres9043


} else {
ifres9042 = False


}

var ifres9040 Obj

if True == ifres9042 {
ifres9040 = True


} else {
ifres9040 = False


}

ifres9039 = ifres9040


} else {
ifres9039 = False


}

if True == ifres9039 {
tmp9028 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7192)
}
__typedArg0 := V7192
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8561 := tmp9028
_ = let__8561

tmp9029 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8561)
}
__typedArg0 := let__8561
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9030 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8561)
}
__typedArg0 := let__8561
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9031 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9030)
}
__typedArg0 := tmp9030
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9032 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8561)
}
__typedArg0 := let__8561
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9033 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9032)
}
__typedArg0 := tmp9032
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9034 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp9033)
}
__typedArg0 := symlet
__typedArg1 := tmp9033
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9035 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9034, Nil)
}
__typedArg0 := tmp9034
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9036 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9031, tmp9035)
}
__typedArg0 := tmp9031
__typedArg1 := tmp9035
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9037 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9029, tmp9036)
}
__typedArg0 := tmp9029
__typedArg1 := tmp9036
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp9037)
}
__typedArg0 := symlet
__typedArg1 := tmp9037
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return(V7192)
return
}


}, 1)

tmp9060 := Call(__e, ns2_1set, symshen_4process_1let, tmp9027)


_ = tmp9060

tmp9061 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__selftop:
V7197 := __self1
_ = V7197
tmp9062 := MakeNative(func(__e *ControlFlow) {
__e.Return(V7197)
return
}, 0)

let__8562 := tmp9062
_ = let__8562

tmp9092 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7197)
}
__typedArg0 := V7197
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9092 {
tmp9063 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7197)
}
__typedArg0 := V7197
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8563 := tmp9063
_ = let__8563

tmp9089 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7197)
}
__typedArg0 := V7197
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9090 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_8s, tmp9089)
}
__typedArg0 := sym_8s
__typedArg1 := tmp9089
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9090 {
tmp9087 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8563)
}
__typedArg0 := let__8563
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9087 {
tmp9064 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8563)
}
__typedArg0 := let__8563
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__8564 := tmp9064
_ = let__8564

tmp9065 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8563)
}
__typedArg0 := let__8563
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8565 := tmp9065
_ = let__8565

tmp9085 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8565)
}
__typedArg0 := let__8565
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9085 {
tmp9066 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8565)
}
__typedArg0 := let__8565
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8566 := tmp9066
_ = let__8566

tmp9083 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8566)
}
__typedArg0 := let__8566
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9083 {
tmp9067 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8s, let__8565)
}
__typedArg0 := sym_8s
__typedArg1 := let__8565
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9068 := Call(__e, PrimFunc(symshen_4process_1_8s), tmp9067)


tmp9069 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9068, Nil)
}
__typedArg0 := tmp9068
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9070 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__8564, tmp9069)
}
__typedArg0 := let__8564
__typedArg1 := tmp9069
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8s, tmp9070)
}
__typedArg0 := sym_8s
__typedArg1 := tmp9070
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9081 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, let__8566)
}
__typedArg0 := Nil
__typedArg1 := let__8566
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9078 Obj

if True == tmp9081 {
tmp9080 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_2) {
return PrimIsString(let__8564)
}
__typedArg0 := let__8564
return Call(__e, PrimFunc(symstring_2), __typedArg0)
})()

var ifres9079 Obj

if True == tmp9080 {
ifres9079 = True


} else {
ifres9079 = False


}

ifres9078 = ifres9079


} else {
ifres9078 = False


}

if True == ifres9078 {
tmp9071 := Call(__e, PrimFunc(symexplode), let__8564)


let__8567 := tmp9071
_ = let__8567

tmp9075 := Call(__e, PrimFunc(symlength), let__8567)


if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_6) {
__typedN0, __typedOK0 := TypedFloat64(tmp9075)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_6) {
return TypedMaterializeBoolean((__typedN0 > __typedN1))
}}
__typedArg0 := tmp9075
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_6), __typedArg0, __typedArg1)
})() {
tmp9072 := Call(__e, PrimFunc(symappend), let__8567, let__8565)


tmp9073 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8s, tmp9072)
}
__typedArg0 := sym_8s
__typedArg1 := tmp9072
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

if PrimFunc(symshen_4process_1_8s) == __self {
__self1 = tmp9073
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4process_1_8s), tmp9073)
return


} else {
__e.Return(V7197)
return
}


} else {
__e.TailApply(PrimFunc(symthaw), let__8562)
return
}


}


} else {
__e.TailApply(PrimFunc(symthaw), let__8562)
return
}


} else {
__e.TailApply(PrimFunc(symthaw), let__8562)
return
}


} else {
__e.TailApply(PrimFunc(symthaw), let__8562)
return
}


} else {
__e.TailApply(PrimFunc(symthaw), let__8562)
return
}


}, 1)

tmp9093 := Call(__e, ns2_1set, symshen_4process_1_8s, tmp9061)


_ = tmp9093

tmp9094 := MakeNative(func(__e *ControlFlow) {
V7207 := __e.Get(1)
_ = V7207
V7208 := __e.Get(2)
_ = V7208
tmp9095 := Call(__e, PrimFunc(symshen_4intern_1type), V7207)


let__8568 := tmp9095
_ = let__8568

tmp9096 := MakeNative(func(__e *ControlFlow) {
Z7211 := __e.Get(1)
_ = Z7211
__e.TailApply(PrimFunc(symshen_4_5datatype_6), Z7211)
return
}, 1)

tmp9097 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__8568, V7208)
}
__typedArg0 := let__8568
__typedArg1 := V7208
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9098 := Call(__e, PrimFunc(symcompile), tmp9096, tmp9097)


let__8569 := tmp9098
_ = let__8569

__e.Return(let__8568)
return


}, 2)

tmp9099 := Call(__e, ns2_1set, symshen_4process_1datatype, tmp9094)


_ = tmp9099

tmp9100 := MakeNative(func(__e *ControlFlow) {
V7212 := __e.Get(1)
_ = V7212
tmp9101 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(V7212)
}
__typedArg0 := V7212
return Call(__e, PrimFunc(symstr), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(tmp9101)
__typedS1, __typedOK1 := TypedString(MakeString("#type"))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := tmp9101
__typedArg1 := MakeString("#type")
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(tmp9101)
__typedS1, __typedOK1 := TypedString(MakeString("#type"))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := tmp9101
__typedArg1 := MakeString("#type")
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symintern), __typedArg0)
})())
return


}, 1)

tmp9103 := Call(__e, ns2_1set, symshen_4intern_1type, tmp9100)


_ = tmp9103

tmp9104 := MakeNative(func(__e *ControlFlow) {
V7213 := __e.Get(1)
_ = V7213
tmp9105 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsynonyms_d)
}
__typedArg0 := symshen_4_dsynonyms_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp9106 := Call(__e, PrimFunc(symappend), V7213, tmp9105)


tmp9107 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsynonyms_d, tmp9106)
}
__typedArg0 := symshen_4_dsynonyms_d
__typedArg1 := tmp9106
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4synonyms_1h), tmp9107)
return


}, 1)

tmp9108 := Call(__e, ns2_1set, symshen_4process_1synonyms, tmp9104)


_ = tmp9108

tmp9109 := MakeNative(func(__e *ControlFlow) {
V7216 := __e.Get(1)
_ = V7216
tmp9147 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7216)
}
__typedArg0 := V7216
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9119 Obj

if True == tmp9147 {
tmp9121 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7216)
}
__typedArg0 := V7216
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8570 := tmp9121
_ = let__8570

tmp9145 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7216)
}
__typedArg0 := V7216
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9146 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefun, tmp9145)
}
__typedArg0 := symdefun
__typedArg1 := tmp9145
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9122 Obj

if True == tmp9146 {
tmp9144 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8570)
}
__typedArg0 := let__8570
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9124 Obj

if True == tmp9144 {
tmp9126 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8570)
}
__typedArg0 := let__8570
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8571 := tmp9126
_ = let__8571

tmp9143 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8571)
}
__typedArg0 := let__8571
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9127 Obj

if True == tmp9143 {
tmp9129 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8571)
}
__typedArg0 := let__8571
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__8572 := tmp9129
_ = let__8572

tmp9130 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8571)
}
__typedArg0 := let__8571
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8573 := tmp9130
_ = let__8573

tmp9142 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8572)
}
__typedArg0 := let__8572
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9131 Obj

if True == tmp9142 {
tmp9140 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8572)
}
__typedArg0 := let__8572
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9141 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp9140)
}
__typedArg0 := Nil
__typedArg1 := tmp9140
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9133 Obj

if True == tmp9141 {
tmp9139 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8573)
}
__typedArg0 := let__8573
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9135 Obj

if True == tmp9139 {
tmp9137 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8573)
}
__typedArg0 := let__8573
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9138 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp9137)
}
__typedArg0 := Nil
__typedArg1 := tmp9137
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9136 Obj

if True == tmp9138 {
ifres9136 = True


} else {
ifres9136 = False


}

ifres9135 = ifres9136


} else {
ifres9135 = False


}

var ifres9134 Obj

if True == ifres9135 {
ifres9134 = True


} else {
ifres9134 = False


}

ifres9133 = ifres9134


} else {
ifres9133 = False


}

var ifres9132 Obj

if True == ifres9133 {
ifres9132 = True


} else {
ifres9132 = False


}

ifres9131 = ifres9132


} else {
ifres9131 = False


}

var ifres9128 Obj

if True == ifres9131 {
ifres9128 = True


} else {
ifres9128 = False


}

ifres9127 = ifres9128


} else {
ifres9127 = False


}

var ifres9125 Obj

if True == ifres9127 {
ifres9125 = True


} else {
ifres9125 = False


}

ifres9124 = ifres9125


} else {
ifres9124 = False


}

var ifres9123 Obj

if True == ifres9124 {
ifres9123 = True


} else {
ifres9123 = False


}

ifres9122 = ifres9123


} else {
ifres9122 = False


}

var ifres9120 Obj

if True == ifres9122 {
ifres9120 = True


} else {
ifres9120 = False


}

ifres9119 = ifres9120


} else {
ifres9119 = False


}

if True == ifres9119 {
tmp9110 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7216)
}
__typedArg0 := V7216
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8574 := tmp9110
_ = let__8574

tmp9111 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8574)
}
__typedArg0 := let__8574
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9112 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9111)
}
__typedArg0 := tmp9111
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9113 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9112)
}
__typedArg0 := tmp9112
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9114 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8574)
}
__typedArg0 := let__8574
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9115 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9114)
}
__typedArg0 := tmp9114
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9116 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9113, tmp9115)
}
__typedArg0 := tmp9113
__typedArg1 := tmp9115
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9117 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_c_4, tmp9116)
}
__typedArg0 := sym_c_4
__typedArg1 := tmp9116
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symeval), tmp9117)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4lambda_1of_1defun)
return
}


}, 1)

tmp9148 := Call(__e, ns2_1set, symshen_4lambda_1of_1defun, tmp9109)


_ = tmp9148

tmp9149 := MakeNative(func(__e *ControlFlow) {
V7222 := __e.Get(1)
_ = V7222
tmp9150 := MakeNative(func(__e *ControlFlow) {
Z7224 := __e.Get(1)
_ = Z7224
__e.TailApply(PrimFunc(symshen_4curry_1type), Z7224)
return
}, 1)

tmp9151 := Call(__e, PrimFunc(symmap), tmp9150, V7222)


let__8575 := tmp9151
_ = let__8575

tmp9152 := Call(__e, PrimFunc(symshen_4compile_1synonyms), let__8575)


tmp9153 := Call(__e, PrimFunc(symshen_4shendef_1_6kldef), symshen_4demod, tmp9152)


tmp9154 := Call(__e, PrimFunc(symshen_4lambda_1of_1defun), tmp9153)


let__8576 := tmp9154
_ = let__8576

tmp9155 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_ddemodulation_1function_d, let__8576)
}
__typedArg0 := symshen_4_ddemodulation_1function_d
__typedArg1 := let__8576
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

let__8577 := tmp9155
_ = let__8577

__e.Return(symsynonyms)
return


}, 1)

tmp9156 := Call(__e, ns2_1set, symshen_4synonyms_1h, tmp9149)


_ = tmp9156

tmp9157 := MakeNative(func(__e *ControlFlow) {
V7229 := __e.Get(1)
_ = V7229
tmp9177 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V7229)
}
__typedArg0 := Nil
__typedArg1 := V7229
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9177 {
tmp9158 := Call(__e, PrimFunc(symgensym), symX)


let__8578 := tmp9158
_ = let__8578

tmp9159 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__8578, Nil)
}
__typedArg0 := let__8578
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9160 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_6, tmp9159)
}
__typedArg0 := sym_1_6
__typedArg1 := tmp9159
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__8578, tmp9160)
}
__typedArg0 := let__8578
__typedArg1 := tmp9160
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9175 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7229)
}
__typedArg0 := V7229
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9171 Obj

if True == tmp9175 {
tmp9173 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7229)
}
__typedArg0 := V7229
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9174 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9173)
}
__typedArg0 := tmp9173
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9172 Obj

if True == tmp9174 {
ifres9172 = True


} else {
ifres9172 = False


}

ifres9171 = ifres9172


} else {
ifres9171 = False


}

if True == ifres9171 {
tmp9161 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7229)
}
__typedArg0 := V7229
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8579 := tmp9161
_ = let__8579

tmp9162 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7229)
}
__typedArg0 := V7229
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9163 := Call(__e, PrimFunc(symshen_4rcons__form), tmp9162)


tmp9164 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8579)
}
__typedArg0 := let__8579
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9165 := Call(__e, PrimFunc(symshen_4rcons__form), tmp9164)


tmp9166 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8579)
}
__typedArg0 := let__8579
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9167 := Call(__e, PrimFunc(symshen_4compile_1synonyms), tmp9166)


tmp9168 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9165, tmp9167)
}
__typedArg0 := tmp9165
__typedArg1 := tmp9167
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9169 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_6, tmp9168)
}
__typedArg0 := sym_1_6
__typedArg1 := tmp9168
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9163, tmp9169)
}
__typedArg0 := tmp9163
__typedArg1 := tmp9169
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

tmp9178 := Call(__e, ns2_1set, symshen_4compile_1synonyms, tmp9157)


_ = tmp9178

tmp9179 := MakeNative(func(__e *ControlFlow) {
V7232 := __e.Get(1)
_ = V7232
tmp9180 := MakeNative(func(__e *ControlFlow) {
__e.Return(V7232)
return
}, 0)

let__8580 := tmp9180
_ = let__8580

tmp9204 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7232)
}
__typedArg0 := V7232
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9204 {
tmp9181 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7232)
}
__typedArg0 := V7232
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8581 := tmp9181
_ = let__8581

tmp9201 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7232)
}
__typedArg0 := V7232
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9202 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_c_4, tmp9201)
}
__typedArg0 := sym_c_4
__typedArg1 := tmp9201
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9202 {
tmp9199 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8581)
}
__typedArg0 := let__8581
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9199 {
tmp9182 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8581)
}
__typedArg0 := let__8581
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__8582 := tmp9182
_ = let__8582

tmp9183 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8581)
}
__typedArg0 := let__8581
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8583 := tmp9183
_ = let__8583

tmp9197 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8583)
}
__typedArg0 := let__8583
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9197 {
tmp9184 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8583)
}
__typedArg0 := let__8583
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8584 := tmp9184
_ = let__8584

tmp9195 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8584)
}
__typedArg0 := let__8584
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9195 {
tmp9185 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_c_4, let__8583)
}
__typedArg0 := sym_c_4
__typedArg1 := let__8583
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9186 := Call(__e, PrimFunc(symshen_4process_1lambda), tmp9185)


tmp9187 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9186, Nil)
}
__typedArg0 := tmp9186
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9188 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__8582, tmp9187)
}
__typedArg0 := let__8582
__typedArg1 := tmp9187
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp9188)
}
__typedArg0 := symlambda
__typedArg1 := tmp9188
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9193 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, let__8584)
}
__typedArg0 := Nil
__typedArg1 := let__8584
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9193 {
tmp9191 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(let__8582)
}
__typedArg0 := let__8582
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

if True == tmp9191 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, let__8581)
}
__typedArg0 := symlambda
__typedArg1 := let__8581
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return
} else {
tmp9189 := Call(__e, PrimFunc(symshen_4app), let__8582, MakeString(" is not a variable\n"), symshen_4s)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(tmp9189)
}
__typedArg0 := tmp9189
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}


} else {
__e.TailApply(PrimFunc(symthaw), let__8580)
return
}


}


} else {
__e.TailApply(PrimFunc(symthaw), let__8580)
return
}


} else {
__e.TailApply(PrimFunc(symthaw), let__8580)
return
}


} else {
__e.TailApply(PrimFunc(symthaw), let__8580)
return
}


} else {
__e.TailApply(PrimFunc(symthaw), let__8580)
return
}


}, 1)

tmp9205 := Call(__e, ns2_1set, symshen_4process_1lambda, tmp9179)


_ = tmp9205

tmp9206 := MakeNative(func(__e *ControlFlow) {
V7243 := __e.Get(1)
_ = V7243
tmp9207 := MakeNative(func(__e *ControlFlow) {
__e.Return(V7243)
return
}, 0)

let__8585 := tmp9207
_ = let__8585

tmp9241 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7243)
}
__typedArg0 := V7243
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9241 {
tmp9208 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7243)
}
__typedArg0 := V7243
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8586 := tmp9208
_ = let__8586

tmp9238 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7243)
}
__typedArg0 := V7243
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9239 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcases, tmp9238)
}
__typedArg0 := symcases
__typedArg1 := tmp9238
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9239 {
tmp9236 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8586)
}
__typedArg0 := let__8586
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9236 {
tmp9209 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8586)
}
__typedArg0 := let__8586
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__8587 := tmp9209
_ = let__8587

tmp9210 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8586)
}
__typedArg0 := let__8586
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8588 := tmp9210
_ = let__8588

tmp9234 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(True, let__8587)
}
__typedArg0 := True
__typedArg1 := let__8587
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9231 Obj

if True == tmp9234 {
tmp9233 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8588)
}
__typedArg0 := let__8588
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9232 Obj

if True == tmp9233 {
ifres9232 = True


} else {
ifres9232 = False


}

ifres9231 = ifres9232


} else {
ifres9231 = False


}

if True == ifres9231 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8588)
}
__typedArg0 := let__8588
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return
} else {
tmp9211 := MakeNative(func(__e *ControlFlow) {
tmp9213 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, let__8588)
}
__typedArg0 := Nil
__typedArg1 := let__8588
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9213 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("error: odd number of case elements\n"))
}
__typedArg0 := MakeString("error: odd number of case elements\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
} else {
__e.TailApply(PrimFunc(symthaw), let__8585)
return
}


}, 0)

let__8589 := tmp9211
_ = let__8589

tmp9229 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8588)
}
__typedArg0 := let__8588
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9229 {
tmp9214 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8588)
}
__typedArg0 := let__8588
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__8590 := tmp9214
_ = let__8590

tmp9215 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8588)
}
__typedArg0 := let__8588
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8591 := tmp9215
_ = let__8591

tmp9227 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, let__8591)
}
__typedArg0 := Nil
__typedArg1 := let__8591
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9227 {
tmp9216 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeString("error: cases exhausted"), Nil)
}
__typedArg0 := MakeString("error: cases exhausted")
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9217 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsimple_1error, tmp9216)
}
__typedArg0 := symsimple_1error
__typedArg1 := tmp9216
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9218 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9217, Nil)
}
__typedArg0 := tmp9217
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9219 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__8590, tmp9218)
}
__typedArg0 := let__8590
__typedArg1 := tmp9218
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9220 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__8587, tmp9219)
}
__typedArg0 := let__8587
__typedArg1 := tmp9219
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp9220)
}
__typedArg0 := symif
__typedArg1 := tmp9220
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9221 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcases, let__8591)
}
__typedArg0 := symcases
__typedArg1 := let__8591
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9222 := Call(__e, PrimFunc(symshen_4process_1cases), tmp9221)


tmp9223 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9222, Nil)
}
__typedArg0 := tmp9222
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9224 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__8590, tmp9223)
}
__typedArg0 := let__8590
__typedArg1 := tmp9223
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9225 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__8587, tmp9224)
}
__typedArg0 := let__8587
__typedArg1 := tmp9224
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp9225)
}
__typedArg0 := symif
__typedArg1 := tmp9225
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


} else {
__e.TailApply(PrimFunc(symthaw), let__8589)
return
}


}


} else {
__e.TailApply(PrimFunc(symthaw), let__8585)
return
}


} else {
__e.TailApply(PrimFunc(symthaw), let__8585)
return
}


} else {
__e.TailApply(PrimFunc(symthaw), let__8585)
return
}


}, 1)

tmp9242 := Call(__e, ns2_1set, symshen_4process_1cases, tmp9206)


_ = tmp9242

tmp9243 := MakeNative(func(__e *ControlFlow) {
V7253 := __e.Get(1)
_ = V7253
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
return PrimCons(symrun, Nil)
}
__typedArg0 := symrun
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9247 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symget_1time, tmp9246)
}
__typedArg0 := symget_1time
__typedArg1 := tmp9246
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9248 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symStart, Nil)
}
__typedArg0 := symStart
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9249 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symFinish, tmp9248)
}
__typedArg0 := symFinish
__typedArg1 := tmp9248
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9250 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1, tmp9249)
}
__typedArg0 := sym_1
__typedArg1 := tmp9249
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9251 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symTime, Nil)
}
__typedArg0 := symTime
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9252 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp9251)
}
__typedArg0 := symstr
__typedArg1 := tmp9251
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9253 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeString(" secs\n"), Nil)
}
__typedArg0 := MakeString(" secs\n")
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9254 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9252, tmp9253)
}
__typedArg0 := tmp9252
__typedArg1 := tmp9253
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9255 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcn, tmp9254)
}
__typedArg0 := symcn
__typedArg1 := tmp9254
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9256 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9255, Nil)
}
__typedArg0 := tmp9255
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9257 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeString("\nrun time: "), tmp9256)
}
__typedArg0 := MakeString("\nrun time: ")
__typedArg1 := tmp9256
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9258 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcn, tmp9257)
}
__typedArg0 := symcn
__typedArg1 := tmp9257
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9259 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstoutput, Nil)
}
__typedArg0 := symstoutput
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9260 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9259, Nil)
}
__typedArg0 := tmp9259
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9261 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9258, tmp9260)
}
__typedArg0 := tmp9258
__typedArg1 := tmp9260
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9262 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympr, tmp9261)
}
__typedArg0 := sympr
__typedArg1 := tmp9261
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9263 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symResult, Nil)
}
__typedArg0 := symResult
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9264 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9262, tmp9263)
}
__typedArg0 := tmp9262
__typedArg1 := tmp9263
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9265 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symMessage, tmp9264)
}
__typedArg0 := symMessage
__typedArg1 := tmp9264
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9266 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9250, tmp9265)
}
__typedArg0 := tmp9250
__typedArg1 := tmp9265
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9267 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symTime, tmp9266)
}
__typedArg0 := symTime
__typedArg1 := tmp9266
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9268 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9247, tmp9267)
}
__typedArg0 := tmp9247
__typedArg1 := tmp9267
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9269 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symFinish, tmp9268)
}
__typedArg0 := symFinish
__typedArg1 := tmp9268
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9270 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V7253, tmp9269)
}
__typedArg0 := V7253
__typedArg1 := tmp9269
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9271 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symResult, tmp9270)
}
__typedArg0 := symResult
__typedArg1 := tmp9270
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9272 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9245, tmp9271)
}
__typedArg0 := tmp9245
__typedArg1 := tmp9271
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9273 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symStart, tmp9272)
}
__typedArg0 := symStart
__typedArg1 := tmp9272
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp9273)
}
__typedArg0 := symlet
__typedArg1 := tmp9273
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp9274 := Call(__e, ns2_1set, symshen_4process_1time, tmp9243)


_ = tmp9274

tmp9275 := MakeNative(func(__e *ControlFlow) {
V7254 := __e.Get(1)
_ = V7254
tmp9296 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7254)
}
__typedArg0 := V7254
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9284 Obj

if True == tmp9296 {
tmp9286 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7254)
}
__typedArg0 := V7254
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8592 := tmp9286
_ = let__8592

tmp9295 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8592)
}
__typedArg0 := let__8592
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9287 Obj

if True == tmp9295 {
tmp9289 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8592)
}
__typedArg0 := let__8592
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8593 := tmp9289
_ = let__8593

tmp9294 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8593)
}
__typedArg0 := let__8593
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9290 Obj

if True == tmp9294 {
tmp9292 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8593)
}
__typedArg0 := let__8593
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9293 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9292)
}
__typedArg0 := tmp9292
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9291 Obj

if True == tmp9293 {
ifres9291 = True


} else {
ifres9291 = False


}

ifres9290 = ifres9291


} else {
ifres9290 = False


}

var ifres9288 Obj

if True == ifres9290 {
ifres9288 = True


} else {
ifres9288 = False


}

ifres9287 = ifres9288


} else {
ifres9287 = False


}

var ifres9285 Obj

if True == ifres9287 {
ifres9285 = True


} else {
ifres9285 = False


}

ifres9284 = ifres9285


} else {
ifres9284 = False


}

if True == ifres9284 {
tmp9276 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7254)
}
__typedArg0 := V7254
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__8594 := tmp9276
_ = let__8594

tmp9277 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7254)
}
__typedArg0 := V7254
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8595 := tmp9277
_ = let__8595

tmp9278 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8595)
}
__typedArg0 := let__8595
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9279 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8595)
}
__typedArg0 := let__8595
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9280 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__8594, tmp9279)
}
__typedArg0 := let__8594
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
return PrimCons(tmp9278, tmp9281)
}
__typedArg0 := tmp9278
__typedArg1 := tmp9281
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__8594, tmp9282)
}
__typedArg0 := let__8594
__typedArg1 := tmp9282
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return(V7254)
return
}


}, 1)

tmp9297 := Call(__e, ns2_1set, symshen_4process_1assoc, tmp9275)


_ = tmp9297

tmp9298 := MakeNative(func(__e *ControlFlow) {
V7259 := __e.Get(1)
_ = V7259
tmp9299 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(V7259)
}
__typedArg0 := V7259
return Call(__e, PrimFunc(symstr), __typedArg0)
})()

tmp9300 := Call(__e, PrimFunc(symshen_4mu_1h), tmp9299)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(tmp9300)
}
__typedArg0 := tmp9300
return Call(__e, PrimFunc(symintern), __typedArg0)
})())
return


}, 1)

tmp9301 := Call(__e, ns2_1set, symshen_4make_1uppercase, tmp9298)


_ = tmp9301

tmp9302 := MakeNative(func(__e *ControlFlow) {
V7260 := __e.Get(1)
_ = V7260
tmp9318 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString(""), V7260)
}
__typedArg0 := MakeString("")
__typedArg1 := V7260
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9318 {
__e.Return(MakeString(""))
return
} else {
tmp9316 := Call(__e, PrimFunc(symshen_4_7string_2), V7260)


if True == tmp9316 {
tmp9303 := Call(__e, PrimFunc(symhdstr), V7260)


tmp9304 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_1_6n) {
return PrimStringToNumber(tmp9303)
}
__typedArg0 := tmp9303
return Call(__e, PrimFunc(symstring_1_6n), __typedArg0)
})()

let__8596 := tmp9304
_ = let__8596

let__8597 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(let__8596)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(32))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := let__8596
__typedArg1 := MakeInteger(32)
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})()
_ = let__8597

tmp9312 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_6_a) {
__typedN0, __typedOK0 := TypedFloat64(let__8596)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(97))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_6_a) {
return TypedMaterializeBoolean((__typedN0 >= __typedN1))
}}
__typedArg0 := let__8596
__typedArg1 := MakeInteger(97)
return Call(__e, PrimFunc(sym_6_a), __typedArg0, __typedArg1)
})()

var ifres9309 Obj

if True == tmp9312 {
tmp9311 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_a) {
__typedN0, __typedOK0 := TypedFloat64(let__8596)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(122))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_5_a) {
return TypedMaterializeBoolean((__typedN0 <= __typedN1))
}}
__typedArg0 := let__8596
__typedArg1 := MakeInteger(122)
return Call(__e, PrimFunc(sym_5_a), __typedArg0, __typedArg1)
})()

var ifres9310 Obj

if True == tmp9311 {
ifres9310 = True


} else {
ifres9310 = False


}

ifres9309 = ifres9310


} else {
ifres9309 = False


}

var ifres9306 Obj

if True == ifres9309 {
tmp9307 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symn_1_6string) {
return PrimNumberToString(let__8597)
}
__typedArg0 := let__8597
return Call(__e, PrimFunc(symn_1_6string), __typedArg0)
})()

ifres9306 = tmp9307


} else {
tmp9308 := Call(__e, PrimFunc(symhdstr), V7260)


ifres9306 = tmp9308


}

let__8598 := ifres9306
_ = let__8598

tmp9314 := Call(__e, PrimFunc(symshen_4mu_1h), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V7260)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V7260
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


__e.TailApply(PrimFunc(sym_8s), let__8598, tmp9314)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4mu_1h)
return
}


}


}, 1)

tmp9319 := Call(__e, ns2_1set, symshen_4mu_1h, tmp9302)


_ = tmp9319

tmp9320 := MakeNative(func(__e *ControlFlow) {
V7264 := __e.Get(1)
_ = V7264
V7265 := __e.Get(2)
_ = V7265
tmp9321 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dmacros_d)
}
__typedArg0 := sym_dmacros_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp9322 := Call(__e, PrimFunc(symshen_4update_1assoc), V7264, V7265, tmp9321)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(sym_dmacros_d, tmp9322)
}
__typedArg0 := sym_dmacros_d
__typedArg1 := tmp9322
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return


}, 2)

tmp9323 := Call(__e, ns2_1set, symshen_4record_1macro, tmp9320)


_ = tmp9323

tmp9324 := MakeNative(func(__e *ControlFlow) {
V7275 := __e.Get(1)
_ = V7275
V7276 := __e.Get(2)
_ = V7276
V7277 := __e.Get(3)
_ = V7277
tmp9341 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V7277)
}
__typedArg0 := Nil
__typedArg1 := V7277
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9341 {
tmp9325 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V7275, V7276)
}
__typedArg0 := V7275
__typedArg1 := V7276
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9325, Nil)
}
__typedArg0 := tmp9325
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9326 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.update-assoc"))
}
__typedArg0 := MakeString("implementation error in shen.update-assoc")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}, 0)

let__8599 := tmp9326
_ = let__8599

tmp9339 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7277)
}
__typedArg0 := V7277
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9339 {
tmp9327 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7277)
}
__typedArg0 := V7277
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__8600 := tmp9327
_ = let__8600

tmp9328 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7277)
}
__typedArg0 := V7277
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8601 := tmp9328
_ = let__8601

tmp9337 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8600)
}
__typedArg0 := let__8600
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9333 Obj

if True == tmp9337 {
tmp9335 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8600)
}
__typedArg0 := let__8600
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9336 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V7275, tmp9335)
}
__typedArg0 := V7275
__typedArg1 := tmp9335
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9334 Obj

if True == tmp9336 {
ifres9334 = True


} else {
ifres9334 = False


}

ifres9333 = ifres9334


} else {
ifres9333 = False


}

if True == ifres9333 {
tmp9329 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8600)
}
__typedArg0 := let__8600
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9330 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9329, V7276)
}
__typedArg0 := tmp9329
__typedArg1 := V7276
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9330, let__8601)
}
__typedArg0 := tmp9330
__typedArg1 := let__8601
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9331 := Call(__e, PrimFunc(symshen_4update_1assoc), V7275, V7276, let__8601)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__8600, tmp9331)
}
__typedArg0 := let__8600
__typedArg1 := tmp9331
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


} else {
__e.TailApply(PrimFunc(symthaw), let__8599)
return
}


}


}, 3)

tmp9342 := Call(__e, ns2_1set, symshen_4update_1assoc, tmp9324)


_ = tmp9342

tmp9343 := MakeNative(func(__e *ControlFlow) {
tmp9351 := Call(__e, PrimFunc(symstinput))


tmp9352 := Call(__e, PrimFunc(symchar_1stream_2), tmp9351)


if True == tmp9352 {
tmp9344 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstinput, Nil)
}
__typedArg0 := symstinput
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9345 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9344, Nil)
}
__typedArg0 := tmp9344
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9346 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1unit_1string, tmp9345)
}
__typedArg0 := symread_1unit_1string
__typedArg1 := tmp9345
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9347 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9346, Nil)
}
__typedArg0 := tmp9346
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring_1_6n, tmp9347)
}
__typedArg0 := symstring_1_6n
__typedArg1 := tmp9347
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9348 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstinput, Nil)
}
__typedArg0 := symstinput
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9349 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9348, Nil)
}
__typedArg0 := tmp9348
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1byte, tmp9349)
}
__typedArg0 := symread_1byte
__typedArg1 := tmp9349
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


}, 0)

tmp9353 := Call(__e, ns2_1set, symshen_4process_1read_1byte, tmp9343)


_ = tmp9353

tmp9354 := MakeNative(func(__e *ControlFlow) {
V7281 := __e.Get(1)
_ = V7281
tmp9355 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4prolog_1vector, Nil)
}
__typedArg0 := symshen_4prolog_1vector
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__8602 := tmp9355
_ = let__8602

tmp9356 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), Nil)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9357 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp9356)
}
__typedArg0 := symvector
__typedArg1 := tmp9356
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9358 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9357, Nil)
}
__typedArg0 := tmp9357
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9359 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), tmp9358)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp9358
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9360 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(True, tmp9359)
}
__typedArg0 := True
__typedArg1 := tmp9359
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9361 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8v, tmp9360)
}
__typedArg0 := sym_8v
__typedArg1 := tmp9360
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__8603 := tmp9361
_ = let__8603

let__8604 := MakeInteger(0)
_ = let__8604

tmp9362 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(True, Nil)
}
__typedArg0 := True
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9363 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfreeze, tmp9362)
}
__typedArg0 := symfreeze
__typedArg1 := tmp9362
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__8605 := tmp9363
_ = let__8605

tmp9364 := MakeNative(func(__e *ControlFlow) {
Z7287 := __e.Get(1)
_ = Z7287
__e.TailApply(PrimFunc(symshen_4_5body_6), Z7287)
return
}, 1)

tmp9365 := Call(__e, PrimFunc(symcompile), tmp9364, V7281)


let__8606 := tmp9365
_ = let__8606

tmp9366 := Call(__e, PrimFunc(symshen_4received), V7281)


let__8607 := tmp9366
_ = let__8607

tmp9367 := Call(__e, PrimFunc(symgensym), symV)


let__8608 := tmp9367
_ = let__8608

tmp9368 := Call(__e, PrimFunc(symgensym), symL)


let__8609 := tmp9368
_ = let__8609

tmp9369 := Call(__e, PrimFunc(symgensym), symK)


let__8610 := tmp9369
_ = let__8610

tmp9370 := Call(__e, PrimFunc(symgensym), symC)


let__8611 := tmp9370
_ = let__8611

tmp9371 := Call(__e, PrimFunc(symshen_4continue), let__8607, let__8606, let__8608, let__8609, let__8610, let__8611)


tmp9372 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9371, Nil)
}
__typedArg0 := tmp9371
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9373 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__8611, tmp9372)
}
__typedArg0 := let__8611
__typedArg1 := tmp9372
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9374 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp9373)
}
__typedArg0 := symlambda
__typedArg1 := tmp9373
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

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
return PrimCons(let__8610, tmp9375)
}
__typedArg0 := let__8610
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
return PrimCons(let__8609, tmp9378)
}
__typedArg0 := let__8609
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
return PrimCons(let__8608, tmp9381)
}
__typedArg0 := let__8608
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

let__8612 := tmp9383
_ = let__8612

tmp9384 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__8605, Nil)
}
__typedArg0 := let__8605
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9385 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__8604, tmp9384)
}
__typedArg0 := let__8604
__typedArg1 := tmp9384
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9386 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__8603, tmp9385)
}
__typedArg0 := let__8603
__typedArg1 := tmp9385
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9387 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__8602, tmp9386)
}
__typedArg0 := let__8602
__typedArg1 := tmp9386
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__8612, tmp9387)
}
__typedArg0 := let__8612
__typedArg1 := tmp9387
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp9388 := Call(__e, ns2_1set, symshen_4call_1prolog, tmp9354)


_ = tmp9388

tmp9389 := MakeNative(func(__e *ControlFlow) {
V7296 := __e.Get(1)
_ = V7296
tmp9390 := MakeNative(func(__e *ControlFlow) {
__e.Return(Nil)
return
}, 0)

let__8613 := tmp9390
_ = let__8613

tmp9405 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7296)
}
__typedArg0 := V7296
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9405 {
tmp9391 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7296)
}
__typedArg0 := V7296
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__8614 := tmp9391
_ = let__8614

tmp9392 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7296)
}
__typedArg0 := V7296
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8615 := tmp9392
_ = let__8615

tmp9403 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symreceive, let__8614)
}
__typedArg0 := symreceive
__typedArg1 := let__8614
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9396 Obj

if True == tmp9403 {
tmp9402 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8615)
}
__typedArg0 := let__8615
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9398 Obj

if True == tmp9402 {
tmp9400 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8615)
}
__typedArg0 := let__8615
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9401 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp9400)
}
__typedArg0 := Nil
__typedArg1 := tmp9400
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9399 Obj

if True == tmp9401 {
ifres9399 = True


} else {
ifres9399 = False


}

ifres9398 = ifres9399


} else {
ifres9398 = False


}

var ifres9397 Obj

if True == ifres9398 {
ifres9397 = True


} else {
ifres9397 = False


}

ifres9396 = ifres9397


} else {
ifres9396 = False


}

if True == ifres9396 {
__e.Return(let__8615)
return
} else {
tmp9393 := Call(__e, PrimFunc(symshen_4received), let__8614)


tmp9394 := Call(__e, PrimFunc(symshen_4received), let__8615)


__e.TailApply(PrimFunc(symunion), tmp9393, tmp9394)
return


}


} else {
__e.TailApply(PrimFunc(symthaw), let__8613)
return
}


}, 1)

tmp9406 := Call(__e, ns2_1set, symshen_4received, tmp9389)


_ = tmp9406

tmp9407 := MakeNative(func(__e *ControlFlow) {
tmp9408 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dprolog_1memory_d)
}
__typedArg0 := symshen_4_dprolog_1memory_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp9409 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symabsvector) {
return PrimAbsvector(tmp9408)
}
__typedArg0 := tmp9408
return Call(__e, PrimFunc(symabsvector), __typedArg0)
})()

let__8616 := tmp9409
_ = let__8616

tmp9410 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(let__8616, MakeInteger(0), symshen_4print_1prolog_1vector)
}
__typedArg0 := let__8616
__typedArg1 := MakeInteger(0)
__typedArg2 := symshen_4print_1prolog_1vector
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

let__8617 := tmp9410
_ = let__8617

tmp9411 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(let__8616, MakeInteger(1), MakeInteger(2))
}
__typedArg0 := let__8616
__typedArg1 := MakeInteger(1)
__typedArg2 := MakeInteger(2)
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

let__8618 := tmp9411
_ = let__8618

__e.Return(let__8618)
return


}, 0)

tmp9412 := Call(__e, ns2_1set, symshen_4prolog_1vector, tmp9407)


_ = tmp9412

tmp9413 := MakeNative(func(__e *ControlFlow) {
V7303 := __e.Get(1)
_ = V7303
__e.Return(V7303)
return
}, 1)

tmp9414 := Call(__e, ns2_1set, symreceive, tmp9413)


_ = tmp9414

tmp9415 := MakeNative(func(__e *ControlFlow) {
V7304 := __e.Get(1)
_ = V7304
tmp9416 := MakeNative(func(__e *ControlFlow) {
__e.Return(V7304)
return
}, 0)

let__8619 := tmp9416
_ = let__8619

tmp9433 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7304)
}
__typedArg0 := V7304
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9433 {
tmp9417 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7304)
}
__typedArg0 := V7304
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__8620 := tmp9417
_ = let__8620

tmp9418 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7304)
}
__typedArg0 := V7304
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8621 := tmp9418
_ = let__8621

tmp9431 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symbar_b, let__8620)
}
__typedArg0 := symbar_b
__typedArg1 := let__8620
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9424 Obj

if True == tmp9431 {
tmp9430 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8621)
}
__typedArg0 := let__8621
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9426 Obj

if True == tmp9430 {
tmp9428 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8621)
}
__typedArg0 := let__8621
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9429 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp9428)
}
__typedArg0 := Nil
__typedArg1 := tmp9428
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9427 Obj

if True == tmp9429 {
ifres9427 = True


} else {
ifres9427 = False


}

ifres9426 = ifres9427


} else {
ifres9426 = False


}

var ifres9425 Obj

if True == ifres9426 {
ifres9425 = True


} else {
ifres9425 = False


}

ifres9424 = ifres9425


} else {
ifres9424 = False


}

if True == ifres9424 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8621)
}
__typedArg0 := let__8621
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return
} else {
tmp9419 := Call(__e, PrimFunc(symshen_4rcons__form), let__8620)


tmp9420 := Call(__e, PrimFunc(symshen_4rcons__form), let__8621)


tmp9421 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9420, Nil)
}
__typedArg0 := tmp9420
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9422 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9419, tmp9421)
}
__typedArg0 := tmp9419
__typedArg1 := tmp9421
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp9422)
}
__typedArg0 := symcons
__typedArg1 := tmp9422
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


} else {
__e.TailApply(PrimFunc(symthaw), let__8619)
return
}


}, 1)

tmp9434 := Call(__e, ns2_1set, symshen_4rcons__form, tmp9415)


_ = tmp9434

tmp9435 := MakeNative(func(__e *ControlFlow) {
V7308 := __e.Get(1)
_ = V7308
tmp9442 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7308)
}
__typedArg0 := V7308
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9442 {
tmp9436 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7308)
}
__typedArg0 := V7308
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9437 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7308)
}
__typedArg0 := V7308
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9438 := Call(__e, PrimFunc(symshen_4tuple_1up), tmp9437)


tmp9439 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9438, Nil)
}
__typedArg0 := tmp9438
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9440 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9436, tmp9439)
}
__typedArg0 := tmp9436
__typedArg1 := tmp9439
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8p, tmp9440)
}
__typedArg0 := sym_8p
__typedArg1 := tmp9440
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return(V7308)
return
}


}, 1)

tmp9443 := Call(__e, ns2_1set, symshen_4tuple_1up, tmp9435)


_ = tmp9443

tmp9444 := MakeNative(func(__e *ControlFlow) {
V7309 := __e.Get(1)
_ = V7309
tmp9445 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dmacros_d)
}
__typedArg0 := sym_dmacros_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp9446 := Call(__e, PrimFunc(symassoc), V7309, tmp9445)


tmp9447 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dmacros_d)
}
__typedArg0 := sym_dmacros_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp9448 := Call(__e, PrimFunc(symremove), tmp9446, tmp9447)


tmp9449 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(sym_dmacros_d, tmp9448)
}
__typedArg0 := sym_dmacros_d
__typedArg1 := tmp9448
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp9449

__e.Return(V7309)
return


}, 1)

__e.TailApply(ns2_1set, symundefmacro, tmp9444)
return




}, 0)

