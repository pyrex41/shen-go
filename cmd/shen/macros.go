package main

import . "github.com/pyrex41/shen-go/kl"

var MacrosMain = MakeNative(func(__e *ControlFlow) {
tmp8642 := MakeNative(func(__e *ControlFlow) {
V7150 := __e.Get(1)
_ = V7150
tmp8643 := MakeNative(func(__e *ControlFlow) {
Z7152 := __e.Get(1)
_ = Z7152
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Z7152)
}
__typedArg0 := Z7152
return Call(__e, PrimFunc(symtl), __typedArg0)
})())
return
}, 1)

tmp8644 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dmacros_d)
}
__typedArg0 := sym_dmacros_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp8645 := Call(__e, PrimFunc(symmap), tmp8643, tmp8644)


let__8551 := tmp8645
_ = let__8551

tmp8646 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsource_1form_1handlers_d)
}
__typedArg0 := symshen_4_dsource_1form_1handlers_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4macroexpand_1h), V7150, let__8551, let__8551, tmp8646)
return


}, 1)

tmp8647 := Call(__e, ns2_1set, symmacroexpand, tmp8642)


_ = tmp8647

tmp8648 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__self3 := __e.Get(3)
__self4 := __e.Get(4)
__selftop:
V7165 := __self1
_ = V7165
V7166 := __self2
_ = V7166
V7167 := __self3
_ = V7167
V7168 := __self4
_ = V7168
tmp8657 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V7166)
}
__typedArg0 := Nil
__typedArg1 := V7166
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp8657 {
__e.Return(V7165)
return
} else {
tmp8655 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7166)
}
__typedArg0 := V7166
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp8655 {
tmp8649 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7166)
}
__typedArg0 := V7166
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8650 := Call(__e, PrimFunc(symshen_4source_1macro_1walk), tmp8649, V7165, V7168)


let__8552 := tmp8650
_ = let__8552

tmp8653 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V7165, let__8552)
}
__typedArg0 := V7165
__typedArg1 := let__8552
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp8653 {
tmp8651 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7166)
}
__typedArg0 := V7166
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4macroexpand_1h) == __self {
__self1, __self2, __self3, __self4 = V7165, tmp8651, V7167, V7168
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4macroexpand_1h), V7165, tmp8651, V7167, V7168)
return


} else {
if PrimFunc(symshen_4macroexpand_1h) == __self {
__self1, __self2, __self3, __self4 = let__8552, V7167, V7167, V7168
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4macroexpand_1h), let__8552, V7167, V7167, V7168)
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

tmp8658 := Call(__e, ns2_1set, symshen_4macroexpand_1h, tmp8648)


_ = tmp8658

tmp8659 := MakeNative(func(__e *ControlFlow) {
V7173 := __e.Get(1)
_ = V7173
V7174 := __e.Get(2)
_ = V7174
V7175 := __e.Get(3)
_ = V7175
tmp8668 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V7175)
}
__typedArg0 := Nil
__typedArg1 := V7175
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp8668 {
__e.TailApply(PrimFunc(symshen_4walk), V7173, V7174)
return
} else {
tmp8665 := Call(__e, PrimFunc(symshen_4source_1form_1handler), V7174, V7175)


tmp8666 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp8665)
}
__typedArg0 := tmp8665
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp8666 {
__e.Return(V7174)
return
} else {
tmp8663 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7174)
}
__typedArg0 := V7174
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp8663 {
tmp8660 := MakeNative(func(__e *ControlFlow) {
Z7176 := __e.Get(1)
_ = Z7176
__e.TailApply(PrimFunc(symshen_4source_1macro_1walk), V7173, Z7176, V7175)
return
}, 1)

tmp8661 := Call(__e, PrimFunc(symmap), tmp8660, V7174)


__e.TailApply(V7173, tmp8661)
return


} else {
__e.TailApply(V7173, V7174)
return
}


}


}


}, 3)

tmp8669 := Call(__e, ns2_1set, symshen_4source_1macro_1walk, tmp8659)


_ = tmp8669

tmp8670 := MakeNative(func(__e *ControlFlow) {
V7177 := __e.Get(1)
_ = V7177
V7178 := __e.Get(2)
_ = V7178
tmp8674 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7178)
}
__typedArg0 := V7178
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp8674 {
tmp8671 := MakeNative(func(__e *ControlFlow) {
Z7179 := __e.Get(1)
_ = Z7179
__e.TailApply(PrimFunc(symshen_4walk), V7177, Z7179)
return
}, 1)

tmp8672 := Call(__e, PrimFunc(symmap), tmp8671, V7178)


__e.TailApply(V7177, tmp8672)
return


} else {
__e.TailApply(V7177, V7178)
return
}


}, 2)

tmp8675 := Call(__e, ns2_1set, symshen_4walk, tmp8670)


_ = tmp8675

tmp8676 := MakeNative(func(__e *ControlFlow) {
V7180 := __e.Get(1)
_ = V7180
tmp8677 := MakeNative(func(__e *ControlFlow) {
__e.Return(V7180)
return
}, 0)

let__8553 := tmp8677
_ = let__8553

tmp8963 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7180)
}
__typedArg0 := V7180
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp8963 {
tmp8678 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7180)
}
__typedArg0 := V7180
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__8554 := tmp8678
_ = let__8554

tmp8679 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7180)
}
__typedArg0 := V7180
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8555 := tmp8679
_ = let__8555

tmp8961 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefmacro, let__8554)
}
__typedArg0 := symdefmacro
__typedArg1 := let__8554
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8958 Obj

if True == tmp8961 {
tmp8960 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8555)
}
__typedArg0 := let__8555
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8959 Obj

if True == tmp8960 {
ifres8959 = True


} else {
ifres8959 = False


}

ifres8958 = ifres8959


} else {
ifres8958 = False


}

if True == ifres8958 {
tmp8680 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8555)
}
__typedArg0 := let__8555
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8681 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8555)
}
__typedArg0 := let__8555
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4process_1def), tmp8680, tmp8681)
return


} else {
tmp8956 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefcc, let__8554)
}
__typedArg0 := symdefcc
__typedArg1 := let__8554
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp8956 {
__e.TailApply(PrimFunc(symshen_4yacc_1_6shen), let__8555)
return
} else {
tmp8954 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symu_b, let__8554)
}
__typedArg0 := symu_b
__typedArg1 := let__8554
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8947 Obj

if True == tmp8954 {
tmp8953 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8555)
}
__typedArg0 := let__8555
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8949 Obj

if True == tmp8953 {
tmp8951 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8555)
}
__typedArg0 := let__8555
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8952 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp8951)
}
__typedArg0 := Nil
__typedArg1 := tmp8951
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8950 Obj

if True == tmp8952 {
ifres8950 = True


} else {
ifres8950 = False


}

ifres8949 = ifres8950


} else {
ifres8949 = False


}

var ifres8948 Obj

if True == ifres8949 {
ifres8948 = True


} else {
ifres8948 = False


}

ifres8947 = ifres8948


} else {
ifres8947 = False


}

if True == ifres8947 {
tmp8682 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8555)
}
__typedArg0 := let__8555
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8683 := Call(__e, PrimFunc(symshen_4make_1uppercase), tmp8682)


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
return PrimCons(symprotect, tmp8684)
}
__typedArg0 := symprotect
__typedArg1 := tmp8684
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp8945 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symerror, let__8554)
}
__typedArg0 := symerror
__typedArg1 := let__8554
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8942 Obj

if True == tmp8945 {
tmp8944 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8555)
}
__typedArg0 := let__8555
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8943 Obj

if True == tmp8944 {
ifres8943 = True


} else {
ifres8943 = False


}

ifres8942 = ifres8943


} else {
ifres8942 = False


}

if True == ifres8942 {
tmp8685 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8555)
}
__typedArg0 := let__8555
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8686 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8555)
}
__typedArg0 := let__8555
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8687 := Call(__e, PrimFunc(symshen_4mkstr), tmp8685, tmp8686)


tmp8688 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8687, Nil)
}
__typedArg0 := tmp8687
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsimple_1error, tmp8688)
}
__typedArg0 := symsimple_1error
__typedArg1 := tmp8688
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp8940 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symoutput, let__8554)
}
__typedArg0 := symoutput
__typedArg1 := let__8554
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8937 Obj

if True == tmp8940 {
tmp8939 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8555)
}
__typedArg0 := let__8555
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8938 Obj

if True == tmp8939 {
ifres8938 = True


} else {
ifres8938 = False


}

ifres8937 = ifres8938


} else {
ifres8937 = False


}

if True == ifres8937 {
tmp8689 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8555)
}
__typedArg0 := let__8555
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8690 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8555)
}
__typedArg0 := let__8555
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8691 := Call(__e, PrimFunc(symshen_4mkstr), tmp8689, tmp8690)


tmp8692 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstoutput, Nil)
}
__typedArg0 := symstoutput
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8693 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8692, Nil)
}
__typedArg0 := tmp8692
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8694 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8691, tmp8693)
}
__typedArg0 := tmp8691
__typedArg1 := tmp8693
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympr, tmp8694)
}
__typedArg0 := sympr
__typedArg1 := tmp8694
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp8935 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sympr, let__8554)
}
__typedArg0 := sympr
__typedArg1 := let__8554
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8928 Obj

if True == tmp8935 {
tmp8934 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8555)
}
__typedArg0 := let__8555
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8930 Obj

if True == tmp8934 {
tmp8932 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8555)
}
__typedArg0 := let__8555
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8933 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp8932)
}
__typedArg0 := Nil
__typedArg1 := tmp8932
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8931 Obj

if True == tmp8933 {
ifres8931 = True


} else {
ifres8931 = False


}

ifres8930 = ifres8931


} else {
ifres8930 = False


}

var ifres8929 Obj

if True == ifres8930 {
ifres8929 = True


} else {
ifres8929 = False


}

ifres8928 = ifres8929


} else {
ifres8928 = False


}

if True == ifres8928 {
tmp8695 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8555)
}
__typedArg0 := let__8555
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8696 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstoutput, Nil)
}
__typedArg0 := symstoutput
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8697 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8696, Nil)
}
__typedArg0 := tmp8696
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8698 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8695, tmp8697)
}
__typedArg0 := tmp8695
__typedArg1 := tmp8697
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympr, tmp8698)
}
__typedArg0 := sympr
__typedArg1 := tmp8698
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp8926 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symmake_1string, let__8554)
}
__typedArg0 := symmake_1string
__typedArg1 := let__8554
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8923 Obj

if True == tmp8926 {
tmp8925 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8555)
}
__typedArg0 := let__8555
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8924 Obj

if True == tmp8925 {
ifres8924 = True


} else {
ifres8924 = False


}

ifres8923 = ifres8924


} else {
ifres8923 = False


}

if True == ifres8923 {
tmp8699 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8555)
}
__typedArg0 := let__8555
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8700 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8555)
}
__typedArg0 := let__8555
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4mkstr), tmp8699, tmp8700)
return


} else {
tmp8921 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlineread, let__8554)
}
__typedArg0 := symlineread
__typedArg1 := let__8554
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8918 Obj

if True == tmp8921 {
tmp8920 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, let__8555)
}
__typedArg0 := Nil
__typedArg1 := let__8555
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8919 Obj

if True == tmp8920 {
ifres8919 = True


} else {
ifres8919 = False


}

ifres8918 = ifres8919


} else {
ifres8918 = False


}

if True == ifres8918 {
tmp8701 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstinput, Nil)
}
__typedArg0 := symstinput
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8702 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8701, Nil)
}
__typedArg0 := tmp8701
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlineread, tmp8702)
}
__typedArg0 := symlineread
__typedArg1 := tmp8702
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp8916 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(syminput, let__8554)
}
__typedArg0 := syminput
__typedArg1 := let__8554
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8913 Obj

if True == tmp8916 {
tmp8915 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, let__8555)
}
__typedArg0 := Nil
__typedArg1 := let__8555
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8914 Obj

if True == tmp8915 {
ifres8914 = True


} else {
ifres8914 = False


}

ifres8913 = ifres8914


} else {
ifres8913 = False


}

if True == ifres8913 {
tmp8703 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstinput, Nil)
}
__typedArg0 := symstinput
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8704 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8703, Nil)
}
__typedArg0 := tmp8703
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminput, tmp8704)
}
__typedArg0 := syminput
__typedArg1 := tmp8704
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp8911 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symread, let__8554)
}
__typedArg0 := symread
__typedArg1 := let__8554
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8908 Obj

if True == tmp8911 {
tmp8910 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, let__8555)
}
__typedArg0 := Nil
__typedArg1 := let__8555
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
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
tmp8705 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstinput, Nil)
}
__typedArg0 := symstinput
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8706 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8705, Nil)
}
__typedArg0 := tmp8705
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread, tmp8706)
}
__typedArg0 := symread
__typedArg1 := tmp8706
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp8906 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(syminput_7, let__8554)
}
__typedArg0 := syminput_7
__typedArg1 := let__8554
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8903 Obj

if True == tmp8906 {
tmp8905 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8555)
}
__typedArg0 := let__8555
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
__e.TailApply(PrimFunc(symshen_4process_1input_7), V7180)
return
} else {
tmp8901 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symread_1byte, let__8554)
}
__typedArg0 := symread_1byte
__typedArg1 := let__8554
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8898 Obj

if True == tmp8901 {
tmp8900 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, let__8555)
}
__typedArg0 := Nil
__typedArg1 := let__8555
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
__e.TailApply(PrimFunc(symshen_4process_1read_1byte))
return
} else {
tmp8896 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symprolog_2, let__8554)
}
__typedArg0 := symprolog_2
__typedArg1 := let__8554
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp8896 {
__e.TailApply(PrimFunc(symshen_4call_1prolog), let__8555)
return
} else {
tmp8894 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefprolog, let__8554)
}
__typedArg0 := symdefprolog
__typedArg1 := let__8554
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8891 Obj

if True == tmp8894 {
tmp8893 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8555)
}
__typedArg0 := let__8555
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8892 Obj

if True == tmp8893 {
ifres8892 = True


} else {
ifres8892 = False


}

ifres8891 = ifres8892


} else {
ifres8891 = False


}

if True == ifres8891 {
tmp8707 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8555)
}
__typedArg0 := let__8555
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8708 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8555)
}
__typedArg0 := let__8555
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4compile_1prolog), tmp8707, tmp8708)
return


} else {
tmp8889 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdatatype, let__8554)
}
__typedArg0 := symdatatype
__typedArg1 := let__8554
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8886 Obj

if True == tmp8889 {
tmp8888 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8555)
}
__typedArg0 := let__8555
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8887 Obj

if True == tmp8888 {
ifres8887 = True


} else {
ifres8887 = False


}

ifres8886 = ifres8887


} else {
ifres8886 = False


}

if True == ifres8886 {
tmp8709 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8555)
}
__typedArg0 := let__8555
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8710 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8555)
}
__typedArg0 := let__8555
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4process_1datatype), tmp8709, tmp8710)
return


} else {
tmp8884 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_8s, let__8554)
}
__typedArg0 := sym_8s
__typedArg1 := let__8554
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp8884 {
__e.TailApply(PrimFunc(symshen_4process_1_8s), V7180)
return
} else {
tmp8882 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symsynonyms, let__8554)
}
__typedArg0 := symsynonyms
__typedArg1 := let__8554
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp8882 {
__e.TailApply(PrimFunc(symshen_4process_1synonyms), let__8555)
return
} else {
tmp8880 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symnl, let__8554)
}
__typedArg0 := symnl
__typedArg1 := let__8554
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8877 Obj

if True == tmp8880 {
tmp8879 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, let__8555)
}
__typedArg0 := Nil
__typedArg1 := let__8555
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8878 Obj

if True == tmp8879 {
ifres8878 = True


} else {
ifres8878 = False


}

ifres8877 = ifres8878


} else {
ifres8877 = False


}

if True == ifres8877 {
tmp8711 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), Nil)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnl, tmp8711)
}
__typedArg0 := symnl
__typedArg1 := tmp8711
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp8875 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlet, let__8554)
}
__typedArg0 := symlet
__typedArg1 := let__8554
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp8875 {
__e.TailApply(PrimFunc(symshen_4process_1let), V7180)
return
} else {
tmp8873 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_c_4, let__8554)
}
__typedArg0 := sym_c_4
__typedArg1 := let__8554
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp8873 {
__e.TailApply(PrimFunc(symshen_4process_1lambda), V7180)
return
} else {
tmp8871 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcases, let__8554)
}
__typedArg0 := symcases
__typedArg1 := let__8554
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp8871 {
__e.TailApply(PrimFunc(symshen_4process_1cases), V7180)
return
} else {
tmp8869 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symtime, let__8554)
}
__typedArg0 := symtime
__typedArg1 := let__8554
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8862 Obj

if True == tmp8869 {
tmp8868 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8555)
}
__typedArg0 := let__8555
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8864 Obj

if True == tmp8868 {
tmp8866 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8555)
}
__typedArg0 := let__8555
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8867 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp8866)
}
__typedArg0 := Nil
__typedArg1 := tmp8866
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8865 Obj

if True == tmp8867 {
ifres8865 = True


} else {
ifres8865 = False


}

ifres8864 = ifres8865


} else {
ifres8864 = False


}

var ifres8863 Obj

if True == ifres8864 {
ifres8863 = True


} else {
ifres8863 = False


}

ifres8862 = ifres8863


} else {
ifres8862 = False


}

if True == ifres8862 {
tmp8712 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8555)
}
__typedArg0 := let__8555
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4process_1time), tmp8712)
return


} else {
tmp8860 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symput, let__8554)
}
__typedArg0 := symput
__typedArg1 := let__8554
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8845 Obj

if True == tmp8860 {
tmp8859 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8555)
}
__typedArg0 := let__8555
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8847 Obj

if True == tmp8859 {
tmp8849 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8555)
}
__typedArg0 := let__8555
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8556 := tmp8849
_ = let__8556

tmp8858 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8556)
}
__typedArg0 := let__8556
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8850 Obj

if True == tmp8858 {
tmp8852 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8556)
}
__typedArg0 := let__8556
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8557 := tmp8852
_ = let__8557

tmp8857 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8557)
}
__typedArg0 := let__8557
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8853 Obj

if True == tmp8857 {
tmp8855 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8557)
}
__typedArg0 := let__8557
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8856 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp8855)
}
__typedArg0 := Nil
__typedArg1 := tmp8855
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8854 Obj

if True == tmp8856 {
ifres8854 = True


} else {
ifres8854 = False


}

ifres8853 = ifres8854


} else {
ifres8853 = False


}

var ifres8851 Obj

if True == ifres8853 {
ifres8851 = True


} else {
ifres8851 = False


}

ifres8850 = ifres8851


} else {
ifres8850 = False


}

var ifres8848 Obj

if True == ifres8850 {
ifres8848 = True


} else {
ifres8848 = False


}

ifres8847 = ifres8848


} else {
ifres8847 = False


}

var ifres8846 Obj

if True == ifres8847 {
ifres8846 = True


} else {
ifres8846 = False


}

ifres8845 = ifres8846


} else {
ifres8845 = False


}

if True == ifres8845 {
tmp8713 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8555)
}
__typedArg0 := let__8555
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8558 := tmp8713
_ = let__8558

tmp8714 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8555)
}
__typedArg0 := let__8555
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8715 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8558)
}
__typedArg0 := let__8558
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8716 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8558)
}
__typedArg0 := let__8558
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8717 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp8716)
}
__typedArg0 := tmp8716
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8718 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dproperty_1vector_d, Nil)
}
__typedArg0 := sym_dproperty_1vector_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8719 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvalue, tmp8718)
}
__typedArg0 := symvalue
__typedArg1 := tmp8718
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8720 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8719, Nil)
}
__typedArg0 := tmp8719
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8721 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8717, tmp8720)
}
__typedArg0 := tmp8717
__typedArg1 := tmp8720
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8722 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8715, tmp8721)
}
__typedArg0 := tmp8715
__typedArg1 := tmp8721
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8723 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8714, tmp8722)
}
__typedArg0 := tmp8714
__typedArg1 := tmp8722
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symput, tmp8723)
}
__typedArg0 := symput
__typedArg1 := tmp8723
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp8843 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symget, let__8554)
}
__typedArg0 := symget
__typedArg1 := let__8554
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8832 Obj

if True == tmp8843 {
tmp8842 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8555)
}
__typedArg0 := let__8555
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8834 Obj

if True == tmp8842 {
tmp8836 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8555)
}
__typedArg0 := let__8555
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8559 := tmp8836
_ = let__8559

tmp8841 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8559)
}
__typedArg0 := let__8559
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8837 Obj

if True == tmp8841 {
tmp8839 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8559)
}
__typedArg0 := let__8559
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8840 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp8839)
}
__typedArg0 := Nil
__typedArg1 := tmp8839
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8838 Obj

if True == tmp8840 {
ifres8838 = True


} else {
ifres8838 = False


}

ifres8837 = ifres8838


} else {
ifres8837 = False


}

var ifres8835 Obj

if True == ifres8837 {
ifres8835 = True


} else {
ifres8835 = False


}

ifres8834 = ifres8835


} else {
ifres8834 = False


}

var ifres8833 Obj

if True == ifres8834 {
ifres8833 = True


} else {
ifres8833 = False


}

ifres8832 = ifres8833


} else {
ifres8832 = False


}

if True == ifres8832 {
tmp8724 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8555)
}
__typedArg0 := let__8555
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8725 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8555)
}
__typedArg0 := let__8555
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8726 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp8725)
}
__typedArg0 := tmp8725
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8727 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dproperty_1vector_d, Nil)
}
__typedArg0 := sym_dproperty_1vector_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8728 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvalue, tmp8727)
}
__typedArg0 := symvalue
__typedArg1 := tmp8727
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8729 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8728, Nil)
}
__typedArg0 := tmp8728
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8730 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8726, tmp8729)
}
__typedArg0 := tmp8726
__typedArg1 := tmp8729
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8731 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8724, tmp8730)
}
__typedArg0 := tmp8724
__typedArg1 := tmp8730
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symget, tmp8731)
}
__typedArg0 := symget
__typedArg1 := tmp8731
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp8830 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symunput, let__8554)
}
__typedArg0 := symunput
__typedArg1 := let__8554
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8819 Obj

if True == tmp8830 {
tmp8829 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8555)
}
__typedArg0 := let__8555
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8821 Obj

if True == tmp8829 {
tmp8823 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8555)
}
__typedArg0 := let__8555
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8560 := tmp8823
_ = let__8560

tmp8828 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8560)
}
__typedArg0 := let__8560
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8824 Obj

if True == tmp8828 {
tmp8826 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8560)
}
__typedArg0 := let__8560
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8827 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp8826)
}
__typedArg0 := Nil
__typedArg1 := tmp8826
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8825 Obj

if True == tmp8827 {
ifres8825 = True


} else {
ifres8825 = False


}

ifres8824 = ifres8825


} else {
ifres8824 = False


}

var ifres8822 Obj

if True == ifres8824 {
ifres8822 = True


} else {
ifres8822 = False


}

ifres8821 = ifres8822


} else {
ifres8821 = False


}

var ifres8820 Obj

if True == ifres8821 {
ifres8820 = True


} else {
ifres8820 = False


}

ifres8819 = ifres8820


} else {
ifres8819 = False


}

if True == ifres8819 {
tmp8732 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8555)
}
__typedArg0 := let__8555
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8733 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8555)
}
__typedArg0 := let__8555
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8734 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp8733)
}
__typedArg0 := tmp8733
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8735 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dproperty_1vector_d, Nil)
}
__typedArg0 := sym_dproperty_1vector_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8736 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvalue, tmp8735)
}
__typedArg0 := symvalue
__typedArg1 := tmp8735
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8737 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8736, Nil)
}
__typedArg0 := tmp8736
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8738 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8734, tmp8737)
}
__typedArg0 := tmp8734
__typedArg1 := tmp8737
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8739 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8732, tmp8738)
}
__typedArg0 := tmp8732
__typedArg1 := tmp8738
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunput, tmp8739)
}
__typedArg0 := symunput
__typedArg1 := tmp8739
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp8817 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4_8c, let__8554)
}
__typedArg0 := symshen_4_8c
__typedArg1 := let__8554
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8810 Obj

if True == tmp8817 {
tmp8816 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8555)
}
__typedArg0 := let__8555
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8812 Obj

if True == tmp8816 {
tmp8814 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8555)
}
__typedArg0 := let__8555
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8815 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp8814)
}
__typedArg0 := Nil
__typedArg1 := tmp8814
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8813 Obj

if True == tmp8815 {
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

if True == ifres8810 {
tmp8740 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8555)
}
__typedArg0 := let__8555
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4rcons__form), tmp8740)
return


} else {
tmp8741 := MakeNative(func(__e *ControlFlow) {
tmp8768 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8555)
}
__typedArg0 := let__8555
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8749 Obj

if True == tmp8768 {
tmp8751 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8555)
}
__typedArg0 := let__8555
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8562 := tmp8751
_ = let__8562

tmp8767 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8562)
}
__typedArg0 := let__8562
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8752 Obj

if True == tmp8767 {
tmp8765 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8562)
}
__typedArg0 := let__8562
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8766 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp8765)
}
__typedArg0 := tmp8765
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8754 Obj

if True == tmp8766 {
tmp8756 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdo, Nil)
}
__typedArg0 := symdo
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8757 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_d, tmp8756)
}
__typedArg0 := sym_d
__typedArg1 := tmp8756
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8758 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_7, tmp8757)
}
__typedArg0 := sym_7
__typedArg1 := tmp8757
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8759 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symor, tmp8758)
}
__typedArg0 := symor
__typedArg1 := tmp8758
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8760 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symand, tmp8759)
}
__typedArg0 := symand
__typedArg1 := tmp8759
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8761 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symappend, tmp8760)
}
__typedArg0 := symappend
__typedArg1 := tmp8760
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8762 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8v, tmp8761)
}
__typedArg0 := sym_8v
__typedArg1 := tmp8761
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8763 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8p, tmp8762)
}
__typedArg0 := sym_8p
__typedArg1 := tmp8762
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8764 := Call(__e, PrimFunc(symelement_2), let__8554, tmp8763)


var ifres8755 Obj

if True == tmp8764 {
ifres8755 = True


} else {
ifres8755 = False


}

ifres8754 = ifres8755


} else {
ifres8754 = False


}

var ifres8753 Obj

if True == ifres8754 {
ifres8753 = True


} else {
ifres8753 = False


}

ifres8752 = ifres8753


} else {
ifres8752 = False


}

var ifres8750 Obj

if True == ifres8752 {
ifres8750 = True


} else {
ifres8750 = False


}

ifres8749 = ifres8750


} else {
ifres8749 = False


}

if True == ifres8749 {
tmp8742 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8555)
}
__typedArg0 := let__8555
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8743 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8555)
}
__typedArg0 := let__8555
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8744 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__8554, tmp8743)
}
__typedArg0 := let__8554
__typedArg1 := tmp8743
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8745 := Call(__e, PrimFunc(symshen_4process_1assoc), tmp8744)


tmp8746 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8745, Nil)
}
__typedArg0 := tmp8745
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8747 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8742, tmp8746)
}
__typedArg0 := tmp8742
__typedArg1 := tmp8746
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__8554, tmp8747)
}
__typedArg0 := let__8554
__typedArg1 := tmp8747
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.TailApply(PrimFunc(symthaw), let__8553)
return
}


}, 0)

let__8561 := tmp8741
_ = let__8561

tmp8808 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4_8ch, let__8554)
}
__typedArg0 := symshen_4_8ch
__typedArg1 := let__8554
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp8808 {
tmp8806 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8555)
}
__typedArg0 := let__8555
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp8806 {
tmp8769 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8555)
}
__typedArg0 := let__8555
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__8563 := tmp8769
_ = let__8563

tmp8770 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8555)
}
__typedArg0 := let__8555
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8564 := tmp8770
_ = let__8564

tmp8804 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8563)
}
__typedArg0 := let__8563
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8784 Obj

if True == tmp8804 {
tmp8786 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8563)
}
__typedArg0 := let__8563
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8565 := tmp8786
_ = let__8565

tmp8803 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8565)
}
__typedArg0 := let__8565
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8787 Obj

if True == tmp8803 {
tmp8789 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8565)
}
__typedArg0 := let__8565
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8566 := tmp8789
_ = let__8566

tmp8802 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8566)
}
__typedArg0 := let__8566
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8790 Obj

if True == tmp8802 {
tmp8800 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8566)
}
__typedArg0 := let__8566
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8801 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp8800)
}
__typedArg0 := Nil
__typedArg1 := tmp8800
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8792 Obj

if True == tmp8801 {
tmp8799 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, let__8564)
}
__typedArg0 := Nil
__typedArg1 := let__8564
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8794 Obj

if True == tmp8799 {
tmp8796 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8565)
}
__typedArg0 := let__8565
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8797 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp8798 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp8796, tmp8797)
}
__typedArg0 := tmp8796
__typedArg1 := tmp8797
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8795 Obj

if True == tmp8798 {
ifres8795 = True


} else {
ifres8795 = False


}

ifres8794 = ifres8795


} else {
ifres8794 = False


}

var ifres8793 Obj

if True == ifres8794 {
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

var ifres8788 Obj

if True == ifres8790 {
ifres8788 = True


} else {
ifres8788 = False


}

ifres8787 = ifres8788


} else {
ifres8787 = False


}

var ifres8785 Obj

if True == ifres8787 {
ifres8785 = True


} else {
ifres8785 = False


}

ifres8784 = ifres8785


} else {
ifres8784 = False


}

if True == ifres8784 {
tmp8771 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8563)
}
__typedArg0 := let__8563
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8567 := tmp8771
_ = let__8567

tmp8772 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8563)
}
__typedArg0 := let__8563
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8773 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8567)
}
__typedArg0 := let__8567
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8774 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8567)
}
__typedArg0 := let__8567
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8775 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_7, tmp8774)
}
__typedArg0 := sym_7
__typedArg1 := tmp8774
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8776 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8775, Nil)
}
__typedArg0 := tmp8775
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8777 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8773, tmp8776)
}
__typedArg0 := tmp8773
__typedArg1 := tmp8776
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8778 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8772, tmp8777)
}
__typedArg0 := tmp8772
__typedArg1 := tmp8777
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8779 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8778, Nil)
}
__typedArg0 := tmp8778
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8780 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1, tmp8779)
}
__typedArg0 := sym_1
__typedArg1 := tmp8779
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4cons_1form_1respect_1modes), tmp8780)
return


} else {
tmp8782 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, let__8564)
}
__typedArg0 := Nil
__typedArg1 := let__8564
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp8782 {
__e.TailApply(PrimFunc(symshen_4cons_1form_1respect_1modes), let__8563)
return
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
__e.TailApply(PrimFunc(symthaw), let__8553)
return
}


}, 1)

tmp8964 := Call(__e, ns2_1set, symshen_4macros, tmp8676)


_ = tmp8964

tmp8965 := MakeNative(func(__e *ControlFlow) {
V7197 := __e.Get(1)
_ = V7197
tmp8966 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4process_1input_7)
return
}, 0)

let__8568 := tmp8966
_ = let__8568

tmp8990 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7197)
}
__typedArg0 := V7197
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp8990 {
tmp8967 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7197)
}
__typedArg0 := V7197
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8569 := tmp8967
_ = let__8569

tmp8987 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7197)
}
__typedArg0 := V7197
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8988 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(syminput_7, tmp8987)
}
__typedArg0 := syminput_7
__typedArg1 := tmp8987
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp8988 {
tmp8985 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8569)
}
__typedArg0 := let__8569
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp8985 {
tmp8968 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8569)
}
__typedArg0 := let__8569
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__8570 := tmp8968
_ = let__8570

tmp8969 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8569)
}
__typedArg0 := let__8569
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8571 := tmp8969
_ = let__8571

tmp8983 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, let__8571)
}
__typedArg0 := Nil
__typedArg1 := let__8571
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp8983 {
tmp8970 := Call(__e, PrimFunc(symshen_4rcons__form), let__8570)


tmp8971 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstinput, Nil)
}
__typedArg0 := symstinput
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8972 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8971, Nil)
}
__typedArg0 := tmp8971
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8973 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8970, tmp8972)
}
__typedArg0 := tmp8970
__typedArg1 := tmp8972
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4input_1h_7, tmp8973)
}
__typedArg0 := symshen_4input_1h_7
__typedArg1 := tmp8973
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp8981 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8571)
}
__typedArg0 := let__8571
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8977 Obj

if True == tmp8981 {
tmp8979 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8571)
}
__typedArg0 := let__8571
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8980 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp8979)
}
__typedArg0 := Nil
__typedArg1 := tmp8979
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8978 Obj

if True == tmp8980 {
ifres8978 = True


} else {
ifres8978 = False


}

ifres8977 = ifres8978


} else {
ifres8977 = False


}

if True == ifres8977 {
tmp8974 := Call(__e, PrimFunc(symshen_4rcons__form), let__8570)


tmp8975 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8974, let__8571)
}
__typedArg0 := tmp8974
__typedArg1 := let__8571
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4input_1h_7, tmp8975)
}
__typedArg0 := symshen_4input_1h_7
__typedArg1 := tmp8975
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.TailApply(PrimFunc(symthaw), let__8568)
return
}


}


} else {
__e.TailApply(PrimFunc(symthaw), let__8568)
return
}


} else {
__e.TailApply(PrimFunc(symthaw), let__8568)
return
}


} else {
__e.TailApply(PrimFunc(symthaw), let__8568)
return
}


}, 1)

tmp8991 := Call(__e, ns2_1set, symshen_4process_1input_7, tmp8965)


_ = tmp8991

tmp8992 := MakeNative(func(__e *ControlFlow) {
V7204 := __e.Get(1)
_ = V7204
tmp8993 := MakeNative(func(__e *ControlFlow) {
__e.Return(V7204)
return
}, 0)

let__8572 := tmp8993
_ = let__8572

tmp9034 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7204)
}
__typedArg0 := V7204
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9034 {
tmp8994 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7204)
}
__typedArg0 := V7204
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__8573 := tmp8994
_ = let__8573

tmp8995 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7204)
}
__typedArg0 := V7204
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8574 := tmp8995
_ = let__8574

tmp9032 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_7, let__8573)
}
__typedArg0 := sym_7
__typedArg1 := let__8573
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9025 Obj

if True == tmp9032 {
tmp9031 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8574)
}
__typedArg0 := let__8574
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9027 Obj

if True == tmp9031 {
tmp9029 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8574)
}
__typedArg0 := let__8574
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9030 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp9029)
}
__typedArg0 := Nil
__typedArg1 := tmp9029
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9028 Obj

if True == tmp9030 {
ifres9028 = True


} else {
ifres9028 = False


}

ifres9027 = ifres9028


} else {
ifres9027 = False


}

var ifres9026 Obj

if True == ifres9027 {
ifres9026 = True


} else {
ifres9026 = False


}

ifres9025 = ifres9026


} else {
ifres9025 = False


}

if True == ifres9025 {
tmp8996 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8574)
}
__typedArg0 := let__8574
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8997 := Call(__e, PrimFunc(symshen_4cons_1form_1respect_1modes), tmp8996)


tmp8998 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8997, Nil)
}
__typedArg0 := tmp8997
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_7, tmp8998)
}
__typedArg0 := sym_7
__typedArg1 := tmp8998
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9023 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_1, let__8573)
}
__typedArg0 := sym_1
__typedArg1 := let__8573
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9016 Obj

if True == tmp9023 {
tmp9022 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8574)
}
__typedArg0 := let__8574
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9018 Obj

if True == tmp9022 {
tmp9020 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8574)
}
__typedArg0 := let__8574
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9021 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp9020)
}
__typedArg0 := Nil
__typedArg1 := tmp9020
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9019 Obj

if True == tmp9021 {
ifres9019 = True


} else {
ifres9019 = False


}

ifres9018 = ifres9019


} else {
ifres9018 = False


}

var ifres9017 Obj

if True == ifres9018 {
ifres9017 = True


} else {
ifres9017 = False


}

ifres9016 = ifres9017


} else {
ifres9016 = False


}

if True == ifres9016 {
tmp8999 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8574)
}
__typedArg0 := let__8574
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9000 := Call(__e, PrimFunc(symshen_4cons_1form_1respect_1modes), tmp8999)


tmp9001 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9000, Nil)
}
__typedArg0 := tmp9000
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1, tmp9001)
}
__typedArg0 := sym_1
__typedArg1 := tmp9001
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9014 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symbar_b, let__8573)
}
__typedArg0 := symbar_b
__typedArg1 := let__8573
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9007 Obj

if True == tmp9014 {
tmp9013 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8574)
}
__typedArg0 := let__8574
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9009 Obj

if True == tmp9013 {
tmp9011 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8574)
}
__typedArg0 := let__8574
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9012 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp9011)
}
__typedArg0 := Nil
__typedArg1 := tmp9011
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9010 Obj

if True == tmp9012 {
ifres9010 = True


} else {
ifres9010 = False


}

ifres9009 = ifres9010


} else {
ifres9009 = False


}

var ifres9008 Obj

if True == ifres9009 {
ifres9008 = True


} else {
ifres9008 = False


}

ifres9007 = ifres9008


} else {
ifres9007 = False


}

if True == ifres9007 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8574)
}
__typedArg0 := let__8574
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return
} else {
tmp9002 := Call(__e, PrimFunc(symshen_4cons_1form_1respect_1modes), let__8573)


tmp9003 := Call(__e, PrimFunc(symshen_4cons_1form_1respect_1modes), let__8574)


tmp9004 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9003, Nil)
}
__typedArg0 := tmp9003
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9005 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9002, tmp9004)
}
__typedArg0 := tmp9002
__typedArg1 := tmp9004
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp9005)
}
__typedArg0 := symcons
__typedArg1 := tmp9005
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


}


}


} else {
__e.TailApply(PrimFunc(symthaw), let__8572)
return
}


}, 1)

tmp9035 := Call(__e, ns2_1set, symshen_4cons_1form_1respect_1modes, tmp8992)


_ = tmp9035

tmp9036 := MakeNative(func(__e *ControlFlow) {
V7208 := __e.Get(1)
_ = V7208
V7209 := __e.Get(2)
_ = V7209
tmp9037 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symX, Nil)
}
__typedArg0 := symX
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9038 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_6, tmp9037)
}
__typedArg0 := sym_1_6
__typedArg1 := tmp9037
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9039 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symX, tmp9038)
}
__typedArg0 := symX
__typedArg1 := tmp9038
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__8575 := tmp9039
_ = let__8575

tmp9040 := Call(__e, PrimFunc(symappend), V7209, let__8575)


tmp9041 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V7208, tmp9040)
}
__typedArg0 := V7208
__typedArg1 := tmp9040
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9042 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefine, tmp9041)
}
__typedArg0 := symdefine
__typedArg1 := tmp9041
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9043 := Call(__e, PrimFunc(symeval), tmp9042)


let__8576 := tmp9043
_ = let__8576

tmp9044 := Call(__e, PrimFunc(symfn), V7208)


tmp9045 := Call(__e, PrimFunc(symshen_4record_1macro), V7208, tmp9044)


let__8577 := tmp9045
_ = let__8577

__e.Return(V7208)
return


}, 2)

tmp9046 := Call(__e, ns2_1set, symshen_4process_1def, tmp9036)


_ = tmp9046

tmp9047 := MakeNative(func(__e *ControlFlow) {
V7213 := __e.Get(1)
_ = V7213
tmp9079 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7213)
}
__typedArg0 := V7213
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9059 Obj

if True == tmp9079 {
tmp9061 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7213)
}
__typedArg0 := V7213
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8578 := tmp9061
_ = let__8578

tmp9077 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7213)
}
__typedArg0 := V7213
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9078 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlet, tmp9077)
}
__typedArg0 := symlet
__typedArg1 := tmp9077
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9062 Obj

if True == tmp9078 {
tmp9076 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8578)
}
__typedArg0 := let__8578
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9064 Obj

if True == tmp9076 {
tmp9066 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8578)
}
__typedArg0 := let__8578
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8579 := tmp9066
_ = let__8579

tmp9075 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8579)
}
__typedArg0 := let__8579
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9067 Obj

if True == tmp9075 {
tmp9069 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8579)
}
__typedArg0 := let__8579
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8580 := tmp9069
_ = let__8580

tmp9074 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8580)
}
__typedArg0 := let__8580
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9070 Obj

if True == tmp9074 {
tmp9072 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8580)
}
__typedArg0 := let__8580
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9073 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9072)
}
__typedArg0 := tmp9072
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9071 Obj

if True == tmp9073 {
ifres9071 = True


} else {
ifres9071 = False


}

ifres9070 = ifres9071


} else {
ifres9070 = False


}

var ifres9068 Obj

if True == ifres9070 {
ifres9068 = True


} else {
ifres9068 = False


}

ifres9067 = ifres9068


} else {
ifres9067 = False


}

var ifres9065 Obj

if True == ifres9067 {
ifres9065 = True


} else {
ifres9065 = False


}

ifres9064 = ifres9065


} else {
ifres9064 = False


}

var ifres9063 Obj

if True == ifres9064 {
ifres9063 = True


} else {
ifres9063 = False


}

ifres9062 = ifres9063


} else {
ifres9062 = False


}

var ifres9060 Obj

if True == ifres9062 {
ifres9060 = True


} else {
ifres9060 = False


}

ifres9059 = ifres9060


} else {
ifres9059 = False


}

if True == ifres9059 {
tmp9048 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7213)
}
__typedArg0 := V7213
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8581 := tmp9048
_ = let__8581

tmp9049 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8581)
}
__typedArg0 := let__8581
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9050 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8581)
}
__typedArg0 := let__8581
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9051 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9050)
}
__typedArg0 := tmp9050
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9052 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8581)
}
__typedArg0 := let__8581
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9053 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9052)
}
__typedArg0 := tmp9052
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9054 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp9053)
}
__typedArg0 := symlet
__typedArg1 := tmp9053
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9055 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9054, Nil)
}
__typedArg0 := tmp9054
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9056 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9051, tmp9055)
}
__typedArg0 := tmp9051
__typedArg1 := tmp9055
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9057 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9049, tmp9056)
}
__typedArg0 := tmp9049
__typedArg1 := tmp9056
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp9057)
}
__typedArg0 := symlet
__typedArg1 := tmp9057
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return(V7213)
return
}


}, 1)

tmp9080 := Call(__e, ns2_1set, symshen_4process_1let, tmp9047)


_ = tmp9080

tmp9081 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__selftop:
V7218 := __self1
_ = V7218
tmp9082 := MakeNative(func(__e *ControlFlow) {
__e.Return(V7218)
return
}, 0)

let__8582 := tmp9082
_ = let__8582

tmp9112 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7218)
}
__typedArg0 := V7218
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9112 {
tmp9083 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7218)
}
__typedArg0 := V7218
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8583 := tmp9083
_ = let__8583

tmp9109 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7218)
}
__typedArg0 := V7218
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9110 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_8s, tmp9109)
}
__typedArg0 := sym_8s
__typedArg1 := tmp9109
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9110 {
tmp9107 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8583)
}
__typedArg0 := let__8583
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9107 {
tmp9084 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8583)
}
__typedArg0 := let__8583
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__8584 := tmp9084
_ = let__8584

tmp9085 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8583)
}
__typedArg0 := let__8583
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8585 := tmp9085
_ = let__8585

tmp9105 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8585)
}
__typedArg0 := let__8585
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9105 {
tmp9086 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8585)
}
__typedArg0 := let__8585
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8586 := tmp9086
_ = let__8586

tmp9103 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8586)
}
__typedArg0 := let__8586
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9103 {
tmp9087 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8s, let__8585)
}
__typedArg0 := sym_8s
__typedArg1 := let__8585
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9088 := Call(__e, PrimFunc(symshen_4process_1_8s), tmp9087)


tmp9089 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9088, Nil)
}
__typedArg0 := tmp9088
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9090 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__8584, tmp9089)
}
__typedArg0 := let__8584
__typedArg1 := tmp9089
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8s, tmp9090)
}
__typedArg0 := sym_8s
__typedArg1 := tmp9090
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9101 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, let__8586)
}
__typedArg0 := Nil
__typedArg1 := let__8586
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9098 Obj

if True == tmp9101 {
tmp9100 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_2) {
return PrimIsString(let__8584)
}
__typedArg0 := let__8584
return Call(__e, PrimFunc(symstring_2), __typedArg0)
})()

var ifres9099 Obj

if True == tmp9100 {
ifres9099 = True


} else {
ifres9099 = False


}

ifres9098 = ifres9099


} else {
ifres9098 = False


}

if True == ifres9098 {
tmp9091 := Call(__e, PrimFunc(symexplode), let__8584)


let__8587 := tmp9091
_ = let__8587

tmp9095 := Call(__e, PrimFunc(symlength), let__8587)


if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_6) {
__typedN0, __typedOK0 := TypedFloat64(tmp9095)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_6) {
return TypedMaterializeBoolean((__typedN0 > __typedN1))
}}
__typedArg0 := tmp9095
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_6), __typedArg0, __typedArg1)
})() {
tmp9092 := Call(__e, PrimFunc(symappend), let__8587, let__8585)


tmp9093 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8s, tmp9092)
}
__typedArg0 := sym_8s
__typedArg1 := tmp9092
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

if PrimFunc(symshen_4process_1_8s) == __self {
__self1 = tmp9093
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4process_1_8s), tmp9093)
return


} else {
__e.Return(V7218)
return
}


} else {
__e.TailApply(PrimFunc(symthaw), let__8582)
return
}


}


} else {
__e.TailApply(PrimFunc(symthaw), let__8582)
return
}


} else {
__e.TailApply(PrimFunc(symthaw), let__8582)
return
}


} else {
__e.TailApply(PrimFunc(symthaw), let__8582)
return
}


} else {
__e.TailApply(PrimFunc(symthaw), let__8582)
return
}


}, 1)

tmp9113 := Call(__e, ns2_1set, symshen_4process_1_8s, tmp9081)


_ = tmp9113

tmp9114 := MakeNative(func(__e *ControlFlow) {
V7228 := __e.Get(1)
_ = V7228
V7229 := __e.Get(2)
_ = V7229
tmp9115 := Call(__e, PrimFunc(symshen_4intern_1type), V7228)


let__8588 := tmp9115
_ = let__8588

tmp9116 := MakeNative(func(__e *ControlFlow) {
Z7232 := __e.Get(1)
_ = Z7232
__e.TailApply(PrimFunc(symshen_4_5datatype_6), Z7232)
return
}, 1)

tmp9117 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__8588, V7229)
}
__typedArg0 := let__8588
__typedArg1 := V7229
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9118 := Call(__e, PrimFunc(symcompile), tmp9116, tmp9117)


let__8589 := tmp9118
_ = let__8589

__e.Return(let__8588)
return


}, 2)

tmp9119 := Call(__e, ns2_1set, symshen_4process_1datatype, tmp9114)


_ = tmp9119

tmp9120 := MakeNative(func(__e *ControlFlow) {
V7233 := __e.Get(1)
_ = V7233
tmp9121 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(V7233)
}
__typedArg0 := V7233
return Call(__e, PrimFunc(symstr), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(tmp9121)
__typedS1, __typedOK1 := TypedString(MakeString("#type"))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := tmp9121
__typedArg1 := MakeString("#type")
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(tmp9121)
__typedS1, __typedOK1 := TypedString(MakeString("#type"))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := tmp9121
__typedArg1 := MakeString("#type")
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symintern), __typedArg0)
})())
return


}, 1)

tmp9123 := Call(__e, ns2_1set, symshen_4intern_1type, tmp9120)


_ = tmp9123

tmp9124 := MakeNative(func(__e *ControlFlow) {
V7234 := __e.Get(1)
_ = V7234
tmp9125 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsynonyms_d)
}
__typedArg0 := symshen_4_dsynonyms_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp9126 := Call(__e, PrimFunc(symappend), V7234, tmp9125)


tmp9127 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsynonyms_d, tmp9126)
}
__typedArg0 := symshen_4_dsynonyms_d
__typedArg1 := tmp9126
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4synonyms_1h), tmp9127)
return


}, 1)

tmp9128 := Call(__e, ns2_1set, symshen_4process_1synonyms, tmp9124)


_ = tmp9128

tmp9129 := MakeNative(func(__e *ControlFlow) {
V7237 := __e.Get(1)
_ = V7237
tmp9167 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7237)
}
__typedArg0 := V7237
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9139 Obj

if True == tmp9167 {
tmp9141 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7237)
}
__typedArg0 := V7237
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8590 := tmp9141
_ = let__8590

tmp9165 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7237)
}
__typedArg0 := V7237
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9166 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefun, tmp9165)
}
__typedArg0 := symdefun
__typedArg1 := tmp9165
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9142 Obj

if True == tmp9166 {
tmp9164 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8590)
}
__typedArg0 := let__8590
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9144 Obj

if True == tmp9164 {
tmp9146 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8590)
}
__typedArg0 := let__8590
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8591 := tmp9146
_ = let__8591

tmp9163 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8591)
}
__typedArg0 := let__8591
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9147 Obj

if True == tmp9163 {
tmp9149 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8591)
}
__typedArg0 := let__8591
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__8592 := tmp9149
_ = let__8592

tmp9150 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8591)
}
__typedArg0 := let__8591
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8593 := tmp9150
_ = let__8593

tmp9162 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8592)
}
__typedArg0 := let__8592
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9151 Obj

if True == tmp9162 {
tmp9160 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8592)
}
__typedArg0 := let__8592
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9161 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp9160)
}
__typedArg0 := Nil
__typedArg1 := tmp9160
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9153 Obj

if True == tmp9161 {
tmp9159 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8593)
}
__typedArg0 := let__8593
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9155 Obj

if True == tmp9159 {
tmp9157 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8593)
}
__typedArg0 := let__8593
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9158 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp9157)
}
__typedArg0 := Nil
__typedArg1 := tmp9157
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9156 Obj

if True == tmp9158 {
ifres9156 = True


} else {
ifres9156 = False


}

ifres9155 = ifres9156


} else {
ifres9155 = False


}

var ifres9154 Obj

if True == ifres9155 {
ifres9154 = True


} else {
ifres9154 = False


}

ifres9153 = ifres9154


} else {
ifres9153 = False


}

var ifres9152 Obj

if True == ifres9153 {
ifres9152 = True


} else {
ifres9152 = False


}

ifres9151 = ifres9152


} else {
ifres9151 = False


}

var ifres9148 Obj

if True == ifres9151 {
ifres9148 = True


} else {
ifres9148 = False


}

ifres9147 = ifres9148


} else {
ifres9147 = False


}

var ifres9145 Obj

if True == ifres9147 {
ifres9145 = True


} else {
ifres9145 = False


}

ifres9144 = ifres9145


} else {
ifres9144 = False


}

var ifres9143 Obj

if True == ifres9144 {
ifres9143 = True


} else {
ifres9143 = False


}

ifres9142 = ifres9143


} else {
ifres9142 = False


}

var ifres9140 Obj

if True == ifres9142 {
ifres9140 = True


} else {
ifres9140 = False


}

ifres9139 = ifres9140


} else {
ifres9139 = False


}

if True == ifres9139 {
tmp9130 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7237)
}
__typedArg0 := V7237
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8594 := tmp9130
_ = let__8594

tmp9131 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8594)
}
__typedArg0 := let__8594
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9132 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9131)
}
__typedArg0 := tmp9131
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9133 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9132)
}
__typedArg0 := tmp9132
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9134 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8594)
}
__typedArg0 := let__8594
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9135 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9134)
}
__typedArg0 := tmp9134
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9136 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9133, tmp9135)
}
__typedArg0 := tmp9133
__typedArg1 := tmp9135
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9137 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_c_4, tmp9136)
}
__typedArg0 := sym_c_4
__typedArg1 := tmp9136
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symeval), tmp9137)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4lambda_1of_1defun)
return
}


}, 1)

tmp9168 := Call(__e, ns2_1set, symshen_4lambda_1of_1defun, tmp9129)


_ = tmp9168

tmp9169 := MakeNative(func(__e *ControlFlow) {
V7243 := __e.Get(1)
_ = V7243
tmp9170 := MakeNative(func(__e *ControlFlow) {
Z7245 := __e.Get(1)
_ = Z7245
__e.TailApply(PrimFunc(symshen_4curry_1type), Z7245)
return
}, 1)

tmp9171 := Call(__e, PrimFunc(symmap), tmp9170, V7243)


let__8595 := tmp9171
_ = let__8595

tmp9172 := Call(__e, PrimFunc(symshen_4compile_1synonyms), let__8595)


tmp9173 := Call(__e, PrimFunc(symshen_4shendef_1_6kldef), symshen_4demod, tmp9172)


tmp9174 := Call(__e, PrimFunc(symshen_4lambda_1of_1defun), tmp9173)


let__8596 := tmp9174
_ = let__8596

tmp9175 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_ddemodulation_1function_d, let__8596)
}
__typedArg0 := symshen_4_ddemodulation_1function_d
__typedArg1 := let__8596
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

let__8597 := tmp9175
_ = let__8597

__e.Return(symsynonyms)
return


}, 1)

tmp9176 := Call(__e, ns2_1set, symshen_4synonyms_1h, tmp9169)


_ = tmp9176

tmp9177 := MakeNative(func(__e *ControlFlow) {
V7250 := __e.Get(1)
_ = V7250
tmp9197 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V7250)
}
__typedArg0 := Nil
__typedArg1 := V7250
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9197 {
tmp9178 := Call(__e, PrimFunc(symgensym), symX)


let__8598 := tmp9178
_ = let__8598

tmp9179 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__8598, Nil)
}
__typedArg0 := let__8598
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9180 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_6, tmp9179)
}
__typedArg0 := sym_1_6
__typedArg1 := tmp9179
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__8598, tmp9180)
}
__typedArg0 := let__8598
__typedArg1 := tmp9180
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9195 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7250)
}
__typedArg0 := V7250
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9191 Obj

if True == tmp9195 {
tmp9193 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7250)
}
__typedArg0 := V7250
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9194 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9193)
}
__typedArg0 := tmp9193
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9192 Obj

if True == tmp9194 {
ifres9192 = True


} else {
ifres9192 = False


}

ifres9191 = ifres9192


} else {
ifres9191 = False


}

if True == ifres9191 {
tmp9181 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7250)
}
__typedArg0 := V7250
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8599 := tmp9181
_ = let__8599

tmp9182 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7250)
}
__typedArg0 := V7250
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9183 := Call(__e, PrimFunc(symshen_4rcons__form), tmp9182)


tmp9184 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8599)
}
__typedArg0 := let__8599
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9185 := Call(__e, PrimFunc(symshen_4rcons__form), tmp9184)


tmp9186 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8599)
}
__typedArg0 := let__8599
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9187 := Call(__e, PrimFunc(symshen_4compile_1synonyms), tmp9186)


tmp9188 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9185, tmp9187)
}
__typedArg0 := tmp9185
__typedArg1 := tmp9187
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9189 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_6, tmp9188)
}
__typedArg0 := sym_1_6
__typedArg1 := tmp9188
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9183, tmp9189)
}
__typedArg0 := tmp9183
__typedArg1 := tmp9189
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

tmp9198 := Call(__e, ns2_1set, symshen_4compile_1synonyms, tmp9177)


_ = tmp9198

tmp9199 := MakeNative(func(__e *ControlFlow) {
V7253 := __e.Get(1)
_ = V7253
tmp9200 := MakeNative(func(__e *ControlFlow) {
__e.Return(V7253)
return
}, 0)

let__8600 := tmp9200
_ = let__8600

tmp9224 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7253)
}
__typedArg0 := V7253
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9224 {
tmp9201 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7253)
}
__typedArg0 := V7253
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8601 := tmp9201
_ = let__8601

tmp9221 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7253)
}
__typedArg0 := V7253
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9222 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_c_4, tmp9221)
}
__typedArg0 := sym_c_4
__typedArg1 := tmp9221
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9222 {
tmp9219 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8601)
}
__typedArg0 := let__8601
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9219 {
tmp9202 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8601)
}
__typedArg0 := let__8601
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__8602 := tmp9202
_ = let__8602

tmp9203 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8601)
}
__typedArg0 := let__8601
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8603 := tmp9203
_ = let__8603

tmp9217 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8603)
}
__typedArg0 := let__8603
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9217 {
tmp9204 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8603)
}
__typedArg0 := let__8603
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8604 := tmp9204
_ = let__8604

tmp9215 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8604)
}
__typedArg0 := let__8604
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9215 {
tmp9205 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_c_4, let__8603)
}
__typedArg0 := sym_c_4
__typedArg1 := let__8603
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9206 := Call(__e, PrimFunc(symshen_4process_1lambda), tmp9205)


tmp9207 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9206, Nil)
}
__typedArg0 := tmp9206
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9208 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__8602, tmp9207)
}
__typedArg0 := let__8602
__typedArg1 := tmp9207
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp9208)
}
__typedArg0 := symlambda
__typedArg1 := tmp9208
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9213 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, let__8604)
}
__typedArg0 := Nil
__typedArg1 := let__8604
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9213 {
tmp9211 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(let__8602)
}
__typedArg0 := let__8602
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

if True == tmp9211 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, let__8601)
}
__typedArg0 := symlambda
__typedArg1 := let__8601
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return
} else {
tmp9209 := Call(__e, PrimFunc(symshen_4app), let__8602, MakeString(" is not a variable\n"), symshen_4s)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(tmp9209)
}
__typedArg0 := tmp9209
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}


} else {
__e.TailApply(PrimFunc(symthaw), let__8600)
return
}


}


} else {
__e.TailApply(PrimFunc(symthaw), let__8600)
return
}


} else {
__e.TailApply(PrimFunc(symthaw), let__8600)
return
}


} else {
__e.TailApply(PrimFunc(symthaw), let__8600)
return
}


} else {
__e.TailApply(PrimFunc(symthaw), let__8600)
return
}


}, 1)

tmp9225 := Call(__e, ns2_1set, symshen_4process_1lambda, tmp9199)


_ = tmp9225

tmp9226 := MakeNative(func(__e *ControlFlow) {
V7264 := __e.Get(1)
_ = V7264
tmp9227 := MakeNative(func(__e *ControlFlow) {
__e.Return(V7264)
return
}, 0)

let__8605 := tmp9227
_ = let__8605

tmp9261 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7264)
}
__typedArg0 := V7264
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9261 {
tmp9228 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7264)
}
__typedArg0 := V7264
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8606 := tmp9228
_ = let__8606

tmp9258 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7264)
}
__typedArg0 := V7264
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9259 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcases, tmp9258)
}
__typedArg0 := symcases
__typedArg1 := tmp9258
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9259 {
tmp9256 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8606)
}
__typedArg0 := let__8606
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9256 {
tmp9229 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8606)
}
__typedArg0 := let__8606
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__8607 := tmp9229
_ = let__8607

tmp9230 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8606)
}
__typedArg0 := let__8606
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8608 := tmp9230
_ = let__8608

tmp9254 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(True, let__8607)
}
__typedArg0 := True
__typedArg1 := let__8607
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9251 Obj

if True == tmp9254 {
tmp9253 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8608)
}
__typedArg0 := let__8608
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9252 Obj

if True == tmp9253 {
ifres9252 = True


} else {
ifres9252 = False


}

ifres9251 = ifres9252


} else {
ifres9251 = False


}

if True == ifres9251 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8608)
}
__typedArg0 := let__8608
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return
} else {
tmp9231 := MakeNative(func(__e *ControlFlow) {
tmp9233 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, let__8608)
}
__typedArg0 := Nil
__typedArg1 := let__8608
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9233 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("error: odd number of case elements\n"))
}
__typedArg0 := MakeString("error: odd number of case elements\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
} else {
__e.TailApply(PrimFunc(symthaw), let__8605)
return
}


}, 0)

let__8609 := tmp9231
_ = let__8609

tmp9249 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8608)
}
__typedArg0 := let__8608
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9249 {
tmp9234 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8608)
}
__typedArg0 := let__8608
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__8610 := tmp9234
_ = let__8610

tmp9235 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8608)
}
__typedArg0 := let__8608
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8611 := tmp9235
_ = let__8611

tmp9247 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, let__8611)
}
__typedArg0 := Nil
__typedArg1 := let__8611
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9247 {
tmp9236 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeString("error: cases exhausted"), Nil)
}
__typedArg0 := MakeString("error: cases exhausted")
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9237 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsimple_1error, tmp9236)
}
__typedArg0 := symsimple_1error
__typedArg1 := tmp9236
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9238 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9237, Nil)
}
__typedArg0 := tmp9237
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9239 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__8610, tmp9238)
}
__typedArg0 := let__8610
__typedArg1 := tmp9238
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9240 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__8607, tmp9239)
}
__typedArg0 := let__8607
__typedArg1 := tmp9239
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp9240)
}
__typedArg0 := symif
__typedArg1 := tmp9240
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9241 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcases, let__8611)
}
__typedArg0 := symcases
__typedArg1 := let__8611
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9242 := Call(__e, PrimFunc(symshen_4process_1cases), tmp9241)


tmp9243 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9242, Nil)
}
__typedArg0 := tmp9242
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9244 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__8610, tmp9243)
}
__typedArg0 := let__8610
__typedArg1 := tmp9243
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9245 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__8607, tmp9244)
}
__typedArg0 := let__8607
__typedArg1 := tmp9244
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp9245)
}
__typedArg0 := symif
__typedArg1 := tmp9245
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


} else {
__e.TailApply(PrimFunc(symthaw), let__8609)
return
}


}


} else {
__e.TailApply(PrimFunc(symthaw), let__8605)
return
}


} else {
__e.TailApply(PrimFunc(symthaw), let__8605)
return
}


} else {
__e.TailApply(PrimFunc(symthaw), let__8605)
return
}


}, 1)

tmp9262 := Call(__e, ns2_1set, symshen_4process_1cases, tmp9226)


_ = tmp9262

tmp9263 := MakeNative(func(__e *ControlFlow) {
V7274 := __e.Get(1)
_ = V7274
tmp9264 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symrun, Nil)
}
__typedArg0 := symrun
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9265 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symget_1time, tmp9264)
}
__typedArg0 := symget_1time
__typedArg1 := tmp9264
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9266 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symrun, Nil)
}
__typedArg0 := symrun
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9267 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symget_1time, tmp9266)
}
__typedArg0 := symget_1time
__typedArg1 := tmp9266
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9268 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symStart, Nil)
}
__typedArg0 := symStart
__typedArg1 := Nil
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
return PrimCons(sym_1, tmp9269)
}
__typedArg0 := sym_1
__typedArg1 := tmp9269
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9271 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symTime, Nil)
}
__typedArg0 := symTime
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9272 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp9271)
}
__typedArg0 := symstr
__typedArg1 := tmp9271
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9273 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeString(" secs\n"), Nil)
}
__typedArg0 := MakeString(" secs\n")
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9274 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9272, tmp9273)
}
__typedArg0 := tmp9272
__typedArg1 := tmp9273
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9275 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcn, tmp9274)
}
__typedArg0 := symcn
__typedArg1 := tmp9274
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9276 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9275, Nil)
}
__typedArg0 := tmp9275
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9277 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeString("\nrun time: "), tmp9276)
}
__typedArg0 := MakeString("\nrun time: ")
__typedArg1 := tmp9276
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9278 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcn, tmp9277)
}
__typedArg0 := symcn
__typedArg1 := tmp9277
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9279 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstoutput, Nil)
}
__typedArg0 := symstoutput
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9280 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9279, Nil)
}
__typedArg0 := tmp9279
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9281 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9278, tmp9280)
}
__typedArg0 := tmp9278
__typedArg1 := tmp9280
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9282 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympr, tmp9281)
}
__typedArg0 := sympr
__typedArg1 := tmp9281
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9283 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symResult, Nil)
}
__typedArg0 := symResult
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9284 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9282, tmp9283)
}
__typedArg0 := tmp9282
__typedArg1 := tmp9283
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9285 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symMessage, tmp9284)
}
__typedArg0 := symMessage
__typedArg1 := tmp9284
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9286 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9270, tmp9285)
}
__typedArg0 := tmp9270
__typedArg1 := tmp9285
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9287 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symTime, tmp9286)
}
__typedArg0 := symTime
__typedArg1 := tmp9286
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9288 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9267, tmp9287)
}
__typedArg0 := tmp9267
__typedArg1 := tmp9287
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9289 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symFinish, tmp9288)
}
__typedArg0 := symFinish
__typedArg1 := tmp9288
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9290 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V7274, tmp9289)
}
__typedArg0 := V7274
__typedArg1 := tmp9289
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9291 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symResult, tmp9290)
}
__typedArg0 := symResult
__typedArg1 := tmp9290
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9292 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9265, tmp9291)
}
__typedArg0 := tmp9265
__typedArg1 := tmp9291
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9293 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symStart, tmp9292)
}
__typedArg0 := symStart
__typedArg1 := tmp9292
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp9293)
}
__typedArg0 := symlet
__typedArg1 := tmp9293
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp9294 := Call(__e, ns2_1set, symshen_4process_1time, tmp9263)


_ = tmp9294

tmp9295 := MakeNative(func(__e *ControlFlow) {
V7275 := __e.Get(1)
_ = V7275
tmp9316 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7275)
}
__typedArg0 := V7275
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9304 Obj

if True == tmp9316 {
tmp9306 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7275)
}
__typedArg0 := V7275
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8612 := tmp9306
_ = let__8612

tmp9315 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8612)
}
__typedArg0 := let__8612
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9307 Obj

if True == tmp9315 {
tmp9309 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8612)
}
__typedArg0 := let__8612
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8613 := tmp9309
_ = let__8613

tmp9314 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8613)
}
__typedArg0 := let__8613
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9310 Obj

if True == tmp9314 {
tmp9312 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8613)
}
__typedArg0 := let__8613
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9313 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9312)
}
__typedArg0 := tmp9312
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9311 Obj

if True == tmp9313 {
ifres9311 = True


} else {
ifres9311 = False


}

ifres9310 = ifres9311


} else {
ifres9310 = False


}

var ifres9308 Obj

if True == ifres9310 {
ifres9308 = True


} else {
ifres9308 = False


}

ifres9307 = ifres9308


} else {
ifres9307 = False


}

var ifres9305 Obj

if True == ifres9307 {
ifres9305 = True


} else {
ifres9305 = False


}

ifres9304 = ifres9305


} else {
ifres9304 = False


}

if True == ifres9304 {
tmp9296 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7275)
}
__typedArg0 := V7275
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__8614 := tmp9296
_ = let__8614

tmp9297 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7275)
}
__typedArg0 := V7275
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8615 := tmp9297
_ = let__8615

tmp9298 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8615)
}
__typedArg0 := let__8615
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9299 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8615)
}
__typedArg0 := let__8615
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9300 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__8614, tmp9299)
}
__typedArg0 := let__8614
__typedArg1 := tmp9299
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9301 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9300, Nil)
}
__typedArg0 := tmp9300
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9302 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9298, tmp9301)
}
__typedArg0 := tmp9298
__typedArg1 := tmp9301
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__8614, tmp9302)
}
__typedArg0 := let__8614
__typedArg1 := tmp9302
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return(V7275)
return
}


}, 1)

tmp9317 := Call(__e, ns2_1set, symshen_4process_1assoc, tmp9295)


_ = tmp9317

tmp9318 := MakeNative(func(__e *ControlFlow) {
V7280 := __e.Get(1)
_ = V7280
tmp9319 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(V7280)
}
__typedArg0 := V7280
return Call(__e, PrimFunc(symstr), __typedArg0)
})()

tmp9320 := Call(__e, PrimFunc(symshen_4mu_1h), tmp9319)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(tmp9320)
}
__typedArg0 := tmp9320
return Call(__e, PrimFunc(symintern), __typedArg0)
})())
return


}, 1)

tmp9321 := Call(__e, ns2_1set, symshen_4make_1uppercase, tmp9318)


_ = tmp9321

tmp9322 := MakeNative(func(__e *ControlFlow) {
V7281 := __e.Get(1)
_ = V7281
tmp9338 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString(""), V7281)
}
__typedArg0 := MakeString("")
__typedArg1 := V7281
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9338 {
__e.Return(MakeString(""))
return
} else {
tmp9336 := Call(__e, PrimFunc(symshen_4_7string_2), V7281)


if True == tmp9336 {
tmp9323 := Call(__e, PrimFunc(symhdstr), V7281)


tmp9324 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_1_6n) {
return PrimStringToNumber(tmp9323)
}
__typedArg0 := tmp9323
return Call(__e, PrimFunc(symstring_1_6n), __typedArg0)
})()

let__8616 := tmp9324
_ = let__8616

let__8617 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(let__8616)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(32))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := let__8616
__typedArg1 := MakeInteger(32)
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})()
_ = let__8617

tmp9332 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_6_a) {
__typedN0, __typedOK0 := TypedFloat64(let__8616)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(97))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_6_a) {
return TypedMaterializeBoolean((__typedN0 >= __typedN1))
}}
__typedArg0 := let__8616
__typedArg1 := MakeInteger(97)
return Call(__e, PrimFunc(sym_6_a), __typedArg0, __typedArg1)
})()

var ifres9329 Obj

if True == tmp9332 {
tmp9331 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_a) {
__typedN0, __typedOK0 := TypedFloat64(let__8616)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(122))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_5_a) {
return TypedMaterializeBoolean((__typedN0 <= __typedN1))
}}
__typedArg0 := let__8616
__typedArg1 := MakeInteger(122)
return Call(__e, PrimFunc(sym_5_a), __typedArg0, __typedArg1)
})()

var ifres9330 Obj

if True == tmp9331 {
ifres9330 = True


} else {
ifres9330 = False


}

ifres9329 = ifres9330


} else {
ifres9329 = False


}

var ifres9326 Obj

if True == ifres9329 {
tmp9327 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symn_1_6string) {
return PrimNumberToString(let__8617)
}
__typedArg0 := let__8617
return Call(__e, PrimFunc(symn_1_6string), __typedArg0)
})()

ifres9326 = tmp9327


} else {
tmp9328 := Call(__e, PrimFunc(symhdstr), V7281)


ifres9326 = tmp9328


}

let__8618 := ifres9326
_ = let__8618

tmp9334 := Call(__e, PrimFunc(symshen_4mu_1h), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V7281)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V7281
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


__e.TailApply(PrimFunc(sym_8s), let__8618, tmp9334)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4mu_1h)
return
}


}


}, 1)

tmp9339 := Call(__e, ns2_1set, symshen_4mu_1h, tmp9322)


_ = tmp9339

tmp9340 := MakeNative(func(__e *ControlFlow) {
V7285 := __e.Get(1)
_ = V7285
V7286 := __e.Get(2)
_ = V7286
tmp9341 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dmacros_d)
}
__typedArg0 := sym_dmacros_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp9342 := Call(__e, PrimFunc(symshen_4update_1assoc), V7285, V7286, tmp9341)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(sym_dmacros_d, tmp9342)
}
__typedArg0 := sym_dmacros_d
__typedArg1 := tmp9342
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return


}, 2)

tmp9343 := Call(__e, ns2_1set, symshen_4record_1macro, tmp9340)


_ = tmp9343

tmp9344 := MakeNative(func(__e *ControlFlow) {
V7296 := __e.Get(1)
_ = V7296
V7297 := __e.Get(2)
_ = V7297
V7298 := __e.Get(3)
_ = V7298
tmp9361 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V7298)
}
__typedArg0 := Nil
__typedArg1 := V7298
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9361 {
tmp9345 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V7296, V7297)
}
__typedArg0 := V7296
__typedArg1 := V7297
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9345, Nil)
}
__typedArg0 := tmp9345
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9346 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.update-assoc"))
}
__typedArg0 := MakeString("implementation error in shen.update-assoc")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}, 0)

let__8619 := tmp9346
_ = let__8619

tmp9359 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7298)
}
__typedArg0 := V7298
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9359 {
tmp9347 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7298)
}
__typedArg0 := V7298
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__8620 := tmp9347
_ = let__8620

tmp9348 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7298)
}
__typedArg0 := V7298
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8621 := tmp9348
_ = let__8621

tmp9357 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8620)
}
__typedArg0 := let__8620
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9353 Obj

if True == tmp9357 {
tmp9355 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8620)
}
__typedArg0 := let__8620
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9356 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V7296, tmp9355)
}
__typedArg0 := V7296
__typedArg1 := tmp9355
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9354 Obj

if True == tmp9356 {
ifres9354 = True


} else {
ifres9354 = False


}

ifres9353 = ifres9354


} else {
ifres9353 = False


}

if True == ifres9353 {
tmp9349 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8620)
}
__typedArg0 := let__8620
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9350 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9349, V7297)
}
__typedArg0 := tmp9349
__typedArg1 := V7297
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9350, let__8621)
}
__typedArg0 := tmp9350
__typedArg1 := let__8621
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9351 := Call(__e, PrimFunc(symshen_4update_1assoc), V7296, V7297, let__8621)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__8620, tmp9351)
}
__typedArg0 := let__8620
__typedArg1 := tmp9351
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


} else {
__e.TailApply(PrimFunc(symthaw), let__8619)
return
}


}


}, 3)

tmp9362 := Call(__e, ns2_1set, symshen_4update_1assoc, tmp9344)


_ = tmp9362

tmp9363 := MakeNative(func(__e *ControlFlow) {
tmp9371 := Call(__e, PrimFunc(symstinput))


tmp9372 := Call(__e, PrimFunc(symchar_1stream_2), tmp9371)


if True == tmp9372 {
tmp9364 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstinput, Nil)
}
__typedArg0 := symstinput
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9365 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9364, Nil)
}
__typedArg0 := tmp9364
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9366 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1unit_1string, tmp9365)
}
__typedArg0 := symread_1unit_1string
__typedArg1 := tmp9365
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9367 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9366, Nil)
}
__typedArg0 := tmp9366
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring_1_6n, tmp9367)
}
__typedArg0 := symstring_1_6n
__typedArg1 := tmp9367
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9368 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstinput, Nil)
}
__typedArg0 := symstinput
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9369 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9368, Nil)
}
__typedArg0 := tmp9368
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1byte, tmp9369)
}
__typedArg0 := symread_1byte
__typedArg1 := tmp9369
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


}, 0)

tmp9373 := Call(__e, ns2_1set, symshen_4process_1read_1byte, tmp9363)


_ = tmp9373

tmp9374 := MakeNative(func(__e *ControlFlow) {
V7302 := __e.Get(1)
_ = V7302
tmp9375 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4prolog_1vector, Nil)
}
__typedArg0 := symshen_4prolog_1vector
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__8622 := tmp9375
_ = let__8622

tmp9376 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), Nil)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9377 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp9376)
}
__typedArg0 := symvector
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
return PrimCons(MakeInteger(0), tmp9378)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp9378
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9380 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(True, tmp9379)
}
__typedArg0 := True
__typedArg1 := tmp9379
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9381 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8v, tmp9380)
}
__typedArg0 := sym_8v
__typedArg1 := tmp9380
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__8623 := tmp9381
_ = let__8623

let__8624 := MakeInteger(0)
_ = let__8624

tmp9382 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(True, Nil)
}
__typedArg0 := True
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9383 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfreeze, tmp9382)
}
__typedArg0 := symfreeze
__typedArg1 := tmp9382
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__8625 := tmp9383
_ = let__8625

tmp9384 := MakeNative(func(__e *ControlFlow) {
Z7308 := __e.Get(1)
_ = Z7308
__e.TailApply(PrimFunc(symshen_4_5body_6), Z7308)
return
}, 1)

tmp9385 := Call(__e, PrimFunc(symcompile), tmp9384, V7302)


let__8626 := tmp9385
_ = let__8626

tmp9386 := Call(__e, PrimFunc(symshen_4received), V7302)


let__8627 := tmp9386
_ = let__8627

tmp9387 := Call(__e, PrimFunc(symgensym), symV)


let__8628 := tmp9387
_ = let__8628

tmp9388 := Call(__e, PrimFunc(symgensym), symL)


let__8629 := tmp9388
_ = let__8629

tmp9389 := Call(__e, PrimFunc(symgensym), symK)


let__8630 := tmp9389
_ = let__8630

tmp9390 := Call(__e, PrimFunc(symgensym), symC)


let__8631 := tmp9390
_ = let__8631

tmp9391 := Call(__e, PrimFunc(symshen_4continue), let__8627, let__8626, let__8628, let__8629, let__8630, let__8631)


tmp9392 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9391, Nil)
}
__typedArg0 := tmp9391
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9393 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__8631, tmp9392)
}
__typedArg0 := let__8631
__typedArg1 := tmp9392
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9394 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp9393)
}
__typedArg0 := symlambda
__typedArg1 := tmp9393
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9395 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9394, Nil)
}
__typedArg0 := tmp9394
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9396 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__8630, tmp9395)
}
__typedArg0 := let__8630
__typedArg1 := tmp9395
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9397 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp9396)
}
__typedArg0 := symlambda
__typedArg1 := tmp9396
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9398 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9397, Nil)
}
__typedArg0 := tmp9397
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9399 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__8629, tmp9398)
}
__typedArg0 := let__8629
__typedArg1 := tmp9398
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9400 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp9399)
}
__typedArg0 := symlambda
__typedArg1 := tmp9399
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9401 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9400, Nil)
}
__typedArg0 := tmp9400
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9402 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__8628, tmp9401)
}
__typedArg0 := let__8628
__typedArg1 := tmp9401
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9403 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp9402)
}
__typedArg0 := symlambda
__typedArg1 := tmp9402
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__8632 := tmp9403
_ = let__8632

tmp9404 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__8625, Nil)
}
__typedArg0 := let__8625
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9405 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__8624, tmp9404)
}
__typedArg0 := let__8624
__typedArg1 := tmp9404
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9406 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__8623, tmp9405)
}
__typedArg0 := let__8623
__typedArg1 := tmp9405
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9407 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__8622, tmp9406)
}
__typedArg0 := let__8622
__typedArg1 := tmp9406
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__8632, tmp9407)
}
__typedArg0 := let__8632
__typedArg1 := tmp9407
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp9408 := Call(__e, ns2_1set, symshen_4call_1prolog, tmp9374)


_ = tmp9408

tmp9409 := MakeNative(func(__e *ControlFlow) {
V7317 := __e.Get(1)
_ = V7317
tmp9410 := MakeNative(func(__e *ControlFlow) {
__e.Return(Nil)
return
}, 0)

let__8633 := tmp9410
_ = let__8633

tmp9425 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7317)
}
__typedArg0 := V7317
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9425 {
tmp9411 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7317)
}
__typedArg0 := V7317
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__8634 := tmp9411
_ = let__8634

tmp9412 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7317)
}
__typedArg0 := V7317
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8635 := tmp9412
_ = let__8635

tmp9423 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symreceive, let__8634)
}
__typedArg0 := symreceive
__typedArg1 := let__8634
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9416 Obj

if True == tmp9423 {
tmp9422 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8635)
}
__typedArg0 := let__8635
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9418 Obj

if True == tmp9422 {
tmp9420 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8635)
}
__typedArg0 := let__8635
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9421 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp9420)
}
__typedArg0 := Nil
__typedArg1 := tmp9420
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9419 Obj

if True == tmp9421 {
ifres9419 = True


} else {
ifres9419 = False


}

ifres9418 = ifres9419


} else {
ifres9418 = False


}

var ifres9417 Obj

if True == ifres9418 {
ifres9417 = True


} else {
ifres9417 = False


}

ifres9416 = ifres9417


} else {
ifres9416 = False


}

if True == ifres9416 {
__e.Return(let__8635)
return
} else {
tmp9413 := Call(__e, PrimFunc(symshen_4received), let__8634)


tmp9414 := Call(__e, PrimFunc(symshen_4received), let__8635)


__e.TailApply(PrimFunc(symunion), tmp9413, tmp9414)
return


}


} else {
__e.TailApply(PrimFunc(symthaw), let__8633)
return
}


}, 1)

tmp9426 := Call(__e, ns2_1set, symshen_4received, tmp9409)


_ = tmp9426

tmp9427 := MakeNative(func(__e *ControlFlow) {
tmp9428 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dprolog_1memory_d)
}
__typedArg0 := symshen_4_dprolog_1memory_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp9429 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symabsvector) {
return PrimAbsvector(tmp9428)
}
__typedArg0 := tmp9428
return Call(__e, PrimFunc(symabsvector), __typedArg0)
})()

let__8636 := tmp9429
_ = let__8636

tmp9430 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(let__8636, MakeInteger(0), symshen_4print_1prolog_1vector)
}
__typedArg0 := let__8636
__typedArg1 := MakeInteger(0)
__typedArg2 := symshen_4print_1prolog_1vector
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

let__8637 := tmp9430
_ = let__8637

tmp9431 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(let__8636, MakeInteger(1), MakeInteger(2))
}
__typedArg0 := let__8636
__typedArg1 := MakeInteger(1)
__typedArg2 := MakeInteger(2)
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

let__8638 := tmp9431
_ = let__8638

__e.Return(let__8638)
return


}, 0)

tmp9432 := Call(__e, ns2_1set, symshen_4prolog_1vector, tmp9427)


_ = tmp9432

tmp9433 := MakeNative(func(__e *ControlFlow) {
V7324 := __e.Get(1)
_ = V7324
__e.Return(V7324)
return
}, 1)

tmp9434 := Call(__e, ns2_1set, symreceive, tmp9433)


_ = tmp9434

tmp9435 := MakeNative(func(__e *ControlFlow) {
V7325 := __e.Get(1)
_ = V7325
tmp9436 := MakeNative(func(__e *ControlFlow) {
__e.Return(V7325)
return
}, 0)

let__8639 := tmp9436
_ = let__8639

tmp9453 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7325)
}
__typedArg0 := V7325
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9453 {
tmp9437 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7325)
}
__typedArg0 := V7325
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__8640 := tmp9437
_ = let__8640

tmp9438 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7325)
}
__typedArg0 := V7325
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__8641 := tmp9438
_ = let__8641

tmp9451 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symbar_b, let__8640)
}
__typedArg0 := symbar_b
__typedArg1 := let__8640
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9444 Obj

if True == tmp9451 {
tmp9450 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__8641)
}
__typedArg0 := let__8641
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9446 Obj

if True == tmp9450 {
tmp9448 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__8641)
}
__typedArg0 := let__8641
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9449 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp9448)
}
__typedArg0 := Nil
__typedArg1 := tmp9448
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9447 Obj

if True == tmp9449 {
ifres9447 = True


} else {
ifres9447 = False


}

ifres9446 = ifres9447


} else {
ifres9446 = False


}

var ifres9445 Obj

if True == ifres9446 {
ifres9445 = True


} else {
ifres9445 = False


}

ifres9444 = ifres9445


} else {
ifres9444 = False


}

if True == ifres9444 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__8641)
}
__typedArg0 := let__8641
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return
} else {
tmp9439 := Call(__e, PrimFunc(symshen_4rcons__form), let__8640)


tmp9440 := Call(__e, PrimFunc(symshen_4rcons__form), let__8641)


tmp9441 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9440, Nil)
}
__typedArg0 := tmp9440
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9442 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9439, tmp9441)
}
__typedArg0 := tmp9439
__typedArg1 := tmp9441
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp9442)
}
__typedArg0 := symcons
__typedArg1 := tmp9442
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


} else {
__e.TailApply(PrimFunc(symthaw), let__8639)
return
}


}, 1)

tmp9454 := Call(__e, ns2_1set, symshen_4rcons__form, tmp9435)


_ = tmp9454

tmp9455 := MakeNative(func(__e *ControlFlow) {
V7329 := __e.Get(1)
_ = V7329
tmp9462 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7329)
}
__typedArg0 := V7329
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9462 {
tmp9456 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7329)
}
__typedArg0 := V7329
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9457 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7329)
}
__typedArg0 := V7329
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9458 := Call(__e, PrimFunc(symshen_4tuple_1up), tmp9457)


tmp9459 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9458, Nil)
}
__typedArg0 := tmp9458
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9460 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9456, tmp9459)
}
__typedArg0 := tmp9456
__typedArg1 := tmp9459
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8p, tmp9460)
}
__typedArg0 := sym_8p
__typedArg1 := tmp9460
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return(V7329)
return
}


}, 1)

tmp9463 := Call(__e, ns2_1set, symshen_4tuple_1up, tmp9455)


_ = tmp9463

tmp9464 := MakeNative(func(__e *ControlFlow) {
V7330 := __e.Get(1)
_ = V7330
tmp9465 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dmacros_d)
}
__typedArg0 := sym_dmacros_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp9466 := Call(__e, PrimFunc(symassoc), V7330, tmp9465)


tmp9467 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dmacros_d)
}
__typedArg0 := sym_dmacros_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp9468 := Call(__e, PrimFunc(symremove), tmp9466, tmp9467)


tmp9469 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(sym_dmacros_d, tmp9468)
}
__typedArg0 := sym_dmacros_d
__typedArg1 := tmp9468
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp9469

__e.Return(V7330)
return


}, 1)

__e.TailApply(ns2_1set, symundefmacro, tmp9464)
return




}, 0)

