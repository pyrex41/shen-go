package main

import . "github.com/pyrex41/shen-go/kl"

var MacrosMain = MakeNative(func(__e *ControlFlow) {
tmp9135 := MakeNative(func(__e *ControlFlow) {
V5824 := __e.Get(1)
_ = V5824
tmp9136 := MakeNative(func(__e *ControlFlow) {
Z5826 := __e.Get(1)
_ = Z5826
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Z5826)
}
__typedArg0 := Z5826
return Call(__e, PrimFunc(symtl), __typedArg0)
})())
return
}, 1)

tmp9137 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dmacros_d)
}
__typedArg0 := sym_dmacros_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp9138 := Call(__e, PrimFunc(symmap), tmp9136, tmp9137)


W58259071 := tmp9138
_ = W58259071

__e.TailApply(PrimFunc(symshen_4macroexpand_1h), V5824, W58259071, W58259071)
return


}, 1)

tmp9139 := Call(__e, ns2_1set, symmacroexpand, tmp9135)


_ = tmp9139

tmp9140 := MakeNative(func(__e *ControlFlow) {
V5835 := __e.Get(1)
_ = V5835
V5836 := __e.Get(2)
_ = V5836
V5837 := __e.Get(3)
_ = V5837
tmp9149 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V5836)
}
__typedArg0 := Nil
__typedArg1 := V5836
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9149 {
__e.Return(V5835)
return
} else {
tmp9147 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5836)
}
__typedArg0 := V5836
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9147 {
tmp9141 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5836)
}
__typedArg0 := V5836
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9142 := Call(__e, PrimFunc(symshen_4walk), tmp9141, V5835)


W58389072 := tmp9142
_ = W58389072

tmp9145 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V5835, W58389072)
}
__typedArg0 := V5835
__typedArg1 := W58389072
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9145 {
tmp9143 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5836)
}
__typedArg0 := V5836
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4macroexpand_1h), V5835, tmp9143, V5837)
return


} else {
__e.TailApply(PrimFunc(symshen_4macroexpand_1h), W58389072, V5837, V5837)
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


}, 3)

tmp9150 := Call(__e, ns2_1set, symshen_4macroexpand_1h, tmp9140)


_ = tmp9150

tmp9151 := MakeNative(func(__e *ControlFlow) {
V5839 := __e.Get(1)
_ = V5839
V5840 := __e.Get(2)
_ = V5840
tmp9155 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5840)
}
__typedArg0 := V5840
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9155 {
tmp9152 := MakeNative(func(__e *ControlFlow) {
Z5841 := __e.Get(1)
_ = Z5841
__e.TailApply(PrimFunc(symshen_4walk), V5839, Z5841)
return
}, 1)

tmp9153 := Call(__e, PrimFunc(symmap), tmp9152, V5840)


__e.TailApply(V5839, tmp9153)
return


} else {
__e.TailApply(V5839, V5840)
return
}


}, 2)

tmp9156 := Call(__e, ns2_1set, symshen_4walk, tmp9151)


_ = tmp9156

tmp9157 := MakeNative(func(__e *ControlFlow) {
V5842 := __e.Get(1)
_ = V5842
tmp9158 := MakeNative(func(__e *ControlFlow) {
__e.Return(V5842)
return
}, 0)

GoTo58439073 := tmp9158
_ = GoTo58439073

tmp9456 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5842)
}
__typedArg0 := V5842
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9456 {
tmp9159 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5842)
}
__typedArg0 := V5842
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

Select58489074 := tmp9159
_ = Select58489074

tmp9160 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5842)
}
__typedArg0 := V5842
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

Select58499075 := tmp9160
_ = Select58499075

tmp9454 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefmacro, Select58489074)
}
__typedArg0 := symdefmacro
__typedArg1 := Select58489074
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9451 Obj

if True == tmp9454 {
tmp9453 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(Select58499075)
}
__typedArg0 := Select58499075
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9452 Obj

if True == tmp9453 {
ifres9452 = True


} else {
ifres9452 = False


}

ifres9451 = ifres9452


} else {
ifres9451 = False


}

if True == ifres9451 {
tmp9161 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(Select58499075)
}
__typedArg0 := Select58499075
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9162 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select58499075)
}
__typedArg0 := Select58499075
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4process_1def), tmp9161, tmp9162)
return


} else {
tmp9449 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefcc, Select58489074)
}
__typedArg0 := symdefcc
__typedArg1 := Select58489074
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9449 {
__e.TailApply(PrimFunc(symshen_4yacc_1_6shen), Select58499075)
return
} else {
tmp9447 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symu_b, Select58489074)
}
__typedArg0 := symu_b
__typedArg1 := Select58489074
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9440 Obj

if True == tmp9447 {
tmp9446 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(Select58499075)
}
__typedArg0 := Select58499075
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9442 Obj

if True == tmp9446 {
tmp9444 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select58499075)
}
__typedArg0 := Select58499075
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9445 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp9444)
}
__typedArg0 := Nil
__typedArg1 := tmp9444
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9443 Obj

if True == tmp9445 {
ifres9443 = True


} else {
ifres9443 = False


}

ifres9442 = ifres9443


} else {
ifres9442 = False


}

var ifres9441 Obj

if True == ifres9442 {
ifres9441 = True


} else {
ifres9441 = False


}

ifres9440 = ifres9441


} else {
ifres9440 = False


}

if True == ifres9440 {
tmp9163 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(Select58499075)
}
__typedArg0 := Select58499075
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9164 := Call(__e, PrimFunc(symshen_4make_1uppercase), tmp9163)


tmp9165 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9164, Nil)
}
__typedArg0 := tmp9164
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symprotect, tmp9165)
}
__typedArg0 := symprotect
__typedArg1 := tmp9165
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9438 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symerror, Select58489074)
}
__typedArg0 := symerror
__typedArg1 := Select58489074
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9435 Obj

if True == tmp9438 {
tmp9437 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(Select58499075)
}
__typedArg0 := Select58499075
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9436 Obj

if True == tmp9437 {
ifres9436 = True


} else {
ifres9436 = False


}

ifres9435 = ifres9436


} else {
ifres9435 = False


}

if True == ifres9435 {
tmp9166 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(Select58499075)
}
__typedArg0 := Select58499075
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9167 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select58499075)
}
__typedArg0 := Select58499075
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9168 := Call(__e, PrimFunc(symshen_4mkstr), tmp9166, tmp9167)


tmp9169 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9168, Nil)
}
__typedArg0 := tmp9168
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsimple_1error, tmp9169)
}
__typedArg0 := symsimple_1error
__typedArg1 := tmp9169
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9433 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symoutput, Select58489074)
}
__typedArg0 := symoutput
__typedArg1 := Select58489074
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9430 Obj

if True == tmp9433 {
tmp9432 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(Select58499075)
}
__typedArg0 := Select58499075
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9431 Obj

if True == tmp9432 {
ifres9431 = True


} else {
ifres9431 = False


}

ifres9430 = ifres9431


} else {
ifres9430 = False


}

if True == ifres9430 {
tmp9170 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(Select58499075)
}
__typedArg0 := Select58499075
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9171 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select58499075)
}
__typedArg0 := Select58499075
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9172 := Call(__e, PrimFunc(symshen_4mkstr), tmp9170, tmp9171)


tmp9173 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstoutput, Nil)
}
__typedArg0 := symstoutput
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9174 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9173, Nil)
}
__typedArg0 := tmp9173
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9175 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9172, tmp9174)
}
__typedArg0 := tmp9172
__typedArg1 := tmp9174
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympr, tmp9175)
}
__typedArg0 := sympr
__typedArg1 := tmp9175
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9428 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sympr, Select58489074)
}
__typedArg0 := sympr
__typedArg1 := Select58489074
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9421 Obj

if True == tmp9428 {
tmp9427 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(Select58499075)
}
__typedArg0 := Select58499075
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9423 Obj

if True == tmp9427 {
tmp9425 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select58499075)
}
__typedArg0 := Select58499075
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9426 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp9425)
}
__typedArg0 := Nil
__typedArg1 := tmp9425
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9424 Obj

if True == tmp9426 {
ifres9424 = True


} else {
ifres9424 = False


}

ifres9423 = ifres9424


} else {
ifres9423 = False


}

var ifres9422 Obj

if True == ifres9423 {
ifres9422 = True


} else {
ifres9422 = False


}

ifres9421 = ifres9422


} else {
ifres9421 = False


}

if True == ifres9421 {
tmp9176 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(Select58499075)
}
__typedArg0 := Select58499075
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9177 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstoutput, Nil)
}
__typedArg0 := symstoutput
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9178 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9177, Nil)
}
__typedArg0 := tmp9177
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9179 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9176, tmp9178)
}
__typedArg0 := tmp9176
__typedArg1 := tmp9178
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympr, tmp9179)
}
__typedArg0 := sympr
__typedArg1 := tmp9179
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9419 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symmake_1string, Select58489074)
}
__typedArg0 := symmake_1string
__typedArg1 := Select58489074
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9416 Obj

if True == tmp9419 {
tmp9418 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(Select58499075)
}
__typedArg0 := Select58499075
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9417 Obj

if True == tmp9418 {
ifres9417 = True


} else {
ifres9417 = False


}

ifres9416 = ifres9417


} else {
ifres9416 = False


}

if True == ifres9416 {
tmp9180 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(Select58499075)
}
__typedArg0 := Select58499075
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9181 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select58499075)
}
__typedArg0 := Select58499075
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4mkstr), tmp9180, tmp9181)
return


} else {
tmp9414 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlineread, Select58489074)
}
__typedArg0 := symlineread
__typedArg1 := Select58489074
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9411 Obj

if True == tmp9414 {
tmp9413 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, Select58499075)
}
__typedArg0 := Nil
__typedArg1 := Select58499075
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9412 Obj

if True == tmp9413 {
ifres9412 = True


} else {
ifres9412 = False


}

ifres9411 = ifres9412


} else {
ifres9411 = False


}

if True == ifres9411 {
tmp9182 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstinput, Nil)
}
__typedArg0 := symstinput
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9183 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9182, Nil)
}
__typedArg0 := tmp9182
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlineread, tmp9183)
}
__typedArg0 := symlineread
__typedArg1 := tmp9183
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9409 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(syminput, Select58489074)
}
__typedArg0 := syminput
__typedArg1 := Select58489074
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9406 Obj

if True == tmp9409 {
tmp9408 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, Select58499075)
}
__typedArg0 := Nil
__typedArg1 := Select58499075
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9407 Obj

if True == tmp9408 {
ifres9407 = True


} else {
ifres9407 = False


}

ifres9406 = ifres9407


} else {
ifres9406 = False


}

if True == ifres9406 {
tmp9184 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstinput, Nil)
}
__typedArg0 := symstinput
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9185 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9184, Nil)
}
__typedArg0 := tmp9184
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminput, tmp9185)
}
__typedArg0 := syminput
__typedArg1 := tmp9185
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9404 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symread, Select58489074)
}
__typedArg0 := symread
__typedArg1 := Select58489074
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9401 Obj

if True == tmp9404 {
tmp9403 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, Select58499075)
}
__typedArg0 := Nil
__typedArg1 := Select58499075
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9402 Obj

if True == tmp9403 {
ifres9402 = True


} else {
ifres9402 = False


}

ifres9401 = ifres9402


} else {
ifres9401 = False


}

if True == ifres9401 {
tmp9186 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstinput, Nil)
}
__typedArg0 := symstinput
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9187 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9186, Nil)
}
__typedArg0 := tmp9186
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread, tmp9187)
}
__typedArg0 := symread
__typedArg1 := tmp9187
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9399 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(syminput_7, Select58489074)
}
__typedArg0 := syminput_7
__typedArg1 := Select58489074
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9396 Obj

if True == tmp9399 {
tmp9398 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(Select58499075)
}
__typedArg0 := Select58499075
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9397 Obj

if True == tmp9398 {
ifres9397 = True


} else {
ifres9397 = False


}

ifres9396 = ifres9397


} else {
ifres9396 = False


}

if True == ifres9396 {
__e.TailApply(PrimFunc(symshen_4process_1input_7), V5842)
return
} else {
tmp9394 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symread_1byte, Select58489074)
}
__typedArg0 := symread_1byte
__typedArg1 := Select58489074
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9391 Obj

if True == tmp9394 {
tmp9393 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, Select58499075)
}
__typedArg0 := Nil
__typedArg1 := Select58499075
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9392 Obj

if True == tmp9393 {
ifres9392 = True


} else {
ifres9392 = False


}

ifres9391 = ifres9392


} else {
ifres9391 = False


}

if True == ifres9391 {
__e.TailApply(PrimFunc(symshen_4process_1read_1byte))
return
} else {
tmp9389 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symprolog_2, Select58489074)
}
__typedArg0 := symprolog_2
__typedArg1 := Select58489074
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9389 {
__e.TailApply(PrimFunc(symshen_4call_1prolog), Select58499075)
return
} else {
tmp9387 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefprolog, Select58489074)
}
__typedArg0 := symdefprolog
__typedArg1 := Select58489074
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9384 Obj

if True == tmp9387 {
tmp9386 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(Select58499075)
}
__typedArg0 := Select58499075
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9385 Obj

if True == tmp9386 {
ifres9385 = True


} else {
ifres9385 = False


}

ifres9384 = ifres9385


} else {
ifres9384 = False


}

if True == ifres9384 {
tmp9188 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(Select58499075)
}
__typedArg0 := Select58499075
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9189 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select58499075)
}
__typedArg0 := Select58499075
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4compile_1prolog), tmp9188, tmp9189)
return


} else {
tmp9382 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdatatype, Select58489074)
}
__typedArg0 := symdatatype
__typedArg1 := Select58489074
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9379 Obj

if True == tmp9382 {
tmp9381 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(Select58499075)
}
__typedArg0 := Select58499075
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9380 Obj

if True == tmp9381 {
ifres9380 = True


} else {
ifres9380 = False


}

ifres9379 = ifres9380


} else {
ifres9379 = False


}

if True == ifres9379 {
tmp9190 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(Select58499075)
}
__typedArg0 := Select58499075
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9191 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select58499075)
}
__typedArg0 := Select58499075
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4process_1datatype), tmp9190, tmp9191)
return


} else {
tmp9377 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_8s, Select58489074)
}
__typedArg0 := sym_8s
__typedArg1 := Select58489074
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9377 {
__e.TailApply(PrimFunc(symshen_4process_1_8s), V5842)
return
} else {
tmp9375 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symsynonyms, Select58489074)
}
__typedArg0 := symsynonyms
__typedArg1 := Select58489074
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9375 {
__e.TailApply(PrimFunc(symshen_4process_1synonyms), Select58499075)
return
} else {
tmp9373 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symnl, Select58489074)
}
__typedArg0 := symnl
__typedArg1 := Select58489074
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9370 Obj

if True == tmp9373 {
tmp9372 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, Select58499075)
}
__typedArg0 := Nil
__typedArg1 := Select58499075
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9371 Obj

if True == tmp9372 {
ifres9371 = True


} else {
ifres9371 = False


}

ifres9370 = ifres9371


} else {
ifres9370 = False


}

if True == ifres9370 {
tmp9192 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), Nil)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnl, tmp9192)
}
__typedArg0 := symnl
__typedArg1 := tmp9192
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9368 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlet, Select58489074)
}
__typedArg0 := symlet
__typedArg1 := Select58489074
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9368 {
__e.TailApply(PrimFunc(symshen_4process_1let), V5842)
return
} else {
tmp9366 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_c_4, Select58489074)
}
__typedArg0 := sym_c_4
__typedArg1 := Select58489074
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9366 {
__e.TailApply(PrimFunc(symshen_4process_1lambda), V5842)
return
} else {
tmp9364 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcases, Select58489074)
}
__typedArg0 := symcases
__typedArg1 := Select58489074
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9364 {
__e.TailApply(PrimFunc(symshen_4process_1cases), V5842)
return
} else {
tmp9362 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symtime, Select58489074)
}
__typedArg0 := symtime
__typedArg1 := Select58489074
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9355 Obj

if True == tmp9362 {
tmp9361 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(Select58499075)
}
__typedArg0 := Select58499075
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9357 Obj

if True == tmp9361 {
tmp9359 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select58499075)
}
__typedArg0 := Select58499075
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9360 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp9359)
}
__typedArg0 := Nil
__typedArg1 := tmp9359
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9358 Obj

if True == tmp9360 {
ifres9358 = True


} else {
ifres9358 = False


}

ifres9357 = ifres9358


} else {
ifres9357 = False


}

var ifres9356 Obj

if True == ifres9357 {
ifres9356 = True


} else {
ifres9356 = False


}

ifres9355 = ifres9356


} else {
ifres9355 = False


}

if True == ifres9355 {
tmp9193 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(Select58499075)
}
__typedArg0 := Select58499075
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4process_1time), tmp9193)
return


} else {
tmp9353 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symput, Select58489074)
}
__typedArg0 := symput
__typedArg1 := Select58489074
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9335 Obj

if True == tmp9353 {
tmp9352 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(Select58499075)
}
__typedArg0 := Select58499075
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9337 Obj

if True == tmp9352 {
tmp9350 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select58499075)
}
__typedArg0 := Select58499075
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9351 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9350)
}
__typedArg0 := tmp9350
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9339 Obj

if True == tmp9351 {
tmp9347 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select58499075)
}
__typedArg0 := Select58499075
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9348 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9347)
}
__typedArg0 := tmp9347
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9349 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9348)
}
__typedArg0 := tmp9348
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9341 Obj

if True == tmp9349 {
tmp9343 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select58499075)
}
__typedArg0 := Select58499075
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9344 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9343)
}
__typedArg0 := tmp9343
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9345 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9344)
}
__typedArg0 := tmp9344
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9346 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp9345)
}
__typedArg0 := Nil
__typedArg1 := tmp9345
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9342 Obj

if True == tmp9346 {
ifres9342 = True


} else {
ifres9342 = False


}

ifres9341 = ifres9342


} else {
ifres9341 = False


}

var ifres9340 Obj

if True == ifres9341 {
ifres9340 = True


} else {
ifres9340 = False


}

ifres9339 = ifres9340


} else {
ifres9339 = False


}

var ifres9338 Obj

if True == ifres9339 {
ifres9338 = True


} else {
ifres9338 = False


}

ifres9337 = ifres9338


} else {
ifres9337 = False


}

var ifres9336 Obj

if True == ifres9337 {
ifres9336 = True


} else {
ifres9336 = False


}

ifres9335 = ifres9336


} else {
ifres9335 = False


}

if True == ifres9335 {
tmp9194 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(Select58499075)
}
__typedArg0 := Select58499075
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9195 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select58499075)
}
__typedArg0 := Select58499075
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9196 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9195)
}
__typedArg0 := tmp9195
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9197 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select58499075)
}
__typedArg0 := Select58499075
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9198 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9197)
}
__typedArg0 := tmp9197
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9199 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9198)
}
__typedArg0 := tmp9198
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9200 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dproperty_1vector_d, Nil)
}
__typedArg0 := sym_dproperty_1vector_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9201 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvalue, tmp9200)
}
__typedArg0 := symvalue
__typedArg1 := tmp9200
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9202 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9201, Nil)
}
__typedArg0 := tmp9201
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9203 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9199, tmp9202)
}
__typedArg0 := tmp9199
__typedArg1 := tmp9202
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9204 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9196, tmp9203)
}
__typedArg0 := tmp9196
__typedArg1 := tmp9203
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9205 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9194, tmp9204)
}
__typedArg0 := tmp9194
__typedArg1 := tmp9204
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symput, tmp9205)
}
__typedArg0 := symput
__typedArg1 := tmp9205
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9333 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symget, Select58489074)
}
__typedArg0 := symget
__typedArg1 := Select58489074
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9321 Obj

if True == tmp9333 {
tmp9332 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(Select58499075)
}
__typedArg0 := Select58499075
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9323 Obj

if True == tmp9332 {
tmp9330 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select58499075)
}
__typedArg0 := Select58499075
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9331 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9330)
}
__typedArg0 := tmp9330
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9325 Obj

if True == tmp9331 {
tmp9327 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select58499075)
}
__typedArg0 := Select58499075
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9328 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9327)
}
__typedArg0 := tmp9327
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9329 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp9328)
}
__typedArg0 := Nil
__typedArg1 := tmp9328
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9326 Obj

if True == tmp9329 {
ifres9326 = True


} else {
ifres9326 = False


}

ifres9325 = ifres9326


} else {
ifres9325 = False


}

var ifres9324 Obj

if True == ifres9325 {
ifres9324 = True


} else {
ifres9324 = False


}

ifres9323 = ifres9324


} else {
ifres9323 = False


}

var ifres9322 Obj

if True == ifres9323 {
ifres9322 = True


} else {
ifres9322 = False


}

ifres9321 = ifres9322


} else {
ifres9321 = False


}

if True == ifres9321 {
tmp9206 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(Select58499075)
}
__typedArg0 := Select58499075
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9207 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select58499075)
}
__typedArg0 := Select58499075
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9208 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9207)
}
__typedArg0 := tmp9207
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9209 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dproperty_1vector_d, Nil)
}
__typedArg0 := sym_dproperty_1vector_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9210 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvalue, tmp9209)
}
__typedArg0 := symvalue
__typedArg1 := tmp9209
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9211 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9210, Nil)
}
__typedArg0 := tmp9210
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9212 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9208, tmp9211)
}
__typedArg0 := tmp9208
__typedArg1 := tmp9211
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9213 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9206, tmp9212)
}
__typedArg0 := tmp9206
__typedArg1 := tmp9212
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symget, tmp9213)
}
__typedArg0 := symget
__typedArg1 := tmp9213
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9319 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symunput, Select58489074)
}
__typedArg0 := symunput
__typedArg1 := Select58489074
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9307 Obj

if True == tmp9319 {
tmp9318 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(Select58499075)
}
__typedArg0 := Select58499075
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9309 Obj

if True == tmp9318 {
tmp9316 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select58499075)
}
__typedArg0 := Select58499075
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9317 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9316)
}
__typedArg0 := tmp9316
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9311 Obj

if True == tmp9317 {
tmp9313 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select58499075)
}
__typedArg0 := Select58499075
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9314 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9313)
}
__typedArg0 := tmp9313
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9315 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp9314)
}
__typedArg0 := Nil
__typedArg1 := tmp9314
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9312 Obj

if True == tmp9315 {
ifres9312 = True


} else {
ifres9312 = False


}

ifres9311 = ifres9312


} else {
ifres9311 = False


}

var ifres9310 Obj

if True == ifres9311 {
ifres9310 = True


} else {
ifres9310 = False


}

ifres9309 = ifres9310


} else {
ifres9309 = False


}

var ifres9308 Obj

if True == ifres9309 {
ifres9308 = True


} else {
ifres9308 = False


}

ifres9307 = ifres9308


} else {
ifres9307 = False


}

if True == ifres9307 {
tmp9214 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(Select58499075)
}
__typedArg0 := Select58499075
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9215 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select58499075)
}
__typedArg0 := Select58499075
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9216 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9215)
}
__typedArg0 := tmp9215
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9217 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dproperty_1vector_d, Nil)
}
__typedArg0 := sym_dproperty_1vector_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9218 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvalue, tmp9217)
}
__typedArg0 := symvalue
__typedArg1 := tmp9217
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9219 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9218, Nil)
}
__typedArg0 := tmp9218
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9220 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9216, tmp9219)
}
__typedArg0 := tmp9216
__typedArg1 := tmp9219
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9221 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9214, tmp9220)
}
__typedArg0 := tmp9214
__typedArg1 := tmp9220
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunput, tmp9221)
}
__typedArg0 := symunput
__typedArg1 := tmp9221
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9305 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4_8c, Select58489074)
}
__typedArg0 := symshen_4_8c
__typedArg1 := Select58489074
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9298 Obj

if True == tmp9305 {
tmp9304 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(Select58499075)
}
__typedArg0 := Select58499075
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9300 Obj

if True == tmp9304 {
tmp9302 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select58499075)
}
__typedArg0 := Select58499075
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9303 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp9302)
}
__typedArg0 := Nil
__typedArg1 := tmp9302
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9301 Obj

if True == tmp9303 {
ifres9301 = True


} else {
ifres9301 = False


}

ifres9300 = ifres9301


} else {
ifres9300 = False


}

var ifres9299 Obj

if True == ifres9300 {
ifres9299 = True


} else {
ifres9299 = False


}

ifres9298 = ifres9299


} else {
ifres9298 = False


}

if True == ifres9298 {
tmp9222 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(Select58499075)
}
__typedArg0 := Select58499075
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4rcons__form), tmp9222)
return


} else {
tmp9223 := MakeNative(func(__e *ControlFlow) {
tmp9251 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(Select58499075)
}
__typedArg0 := Select58499075
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9231 Obj

if True == tmp9251 {
tmp9249 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select58499075)
}
__typedArg0 := Select58499075
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9250 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9249)
}
__typedArg0 := tmp9249
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9233 Obj

if True == tmp9250 {
tmp9246 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select58499075)
}
__typedArg0 := Select58499075
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9247 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9246)
}
__typedArg0 := tmp9246
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9248 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9247)
}
__typedArg0 := tmp9247
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9235 Obj

if True == tmp9248 {
tmp9237 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdo, Nil)
}
__typedArg0 := symdo
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9238 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_d, tmp9237)
}
__typedArg0 := sym_d
__typedArg1 := tmp9237
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9239 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_7, tmp9238)
}
__typedArg0 := sym_7
__typedArg1 := tmp9238
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9240 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symor, tmp9239)
}
__typedArg0 := symor
__typedArg1 := tmp9239
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9241 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symand, tmp9240)
}
__typedArg0 := symand
__typedArg1 := tmp9240
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9242 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symappend, tmp9241)
}
__typedArg0 := symappend
__typedArg1 := tmp9241
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9243 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8v, tmp9242)
}
__typedArg0 := sym_8v
__typedArg1 := tmp9242
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9244 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8p, tmp9243)
}
__typedArg0 := sym_8p
__typedArg1 := tmp9243
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9245 := Call(__e, PrimFunc(symelement_2), Select58489074, tmp9244)


var ifres9236 Obj

if True == tmp9245 {
ifres9236 = True


} else {
ifres9236 = False


}

ifres9235 = ifres9236


} else {
ifres9235 = False


}

var ifres9234 Obj

if True == ifres9235 {
ifres9234 = True


} else {
ifres9234 = False


}

ifres9233 = ifres9234


} else {
ifres9233 = False


}

var ifres9232 Obj

if True == ifres9233 {
ifres9232 = True


} else {
ifres9232 = False


}

ifres9231 = ifres9232


} else {
ifres9231 = False


}

if True == ifres9231 {
tmp9224 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(Select58499075)
}
__typedArg0 := Select58499075
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9225 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select58499075)
}
__typedArg0 := Select58499075
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9226 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Select58489074, tmp9225)
}
__typedArg0 := Select58489074
__typedArg1 := tmp9225
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9227 := Call(__e, PrimFunc(symshen_4process_1assoc), tmp9226)


tmp9228 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9227, Nil)
}
__typedArg0 := tmp9227
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9229 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9224, tmp9228)
}
__typedArg0 := tmp9224
__typedArg1 := tmp9228
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Select58489074, tmp9229)
}
__typedArg0 := Select58489074
__typedArg1 := tmp9229
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo58439073)
return
}


}, 0)

GoTo58449076 := tmp9223
_ = GoTo58449076

tmp9296 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4_8ch, Select58489074)
}
__typedArg0 := symshen_4_8ch
__typedArg1 := Select58489074
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9296 {
tmp9294 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(Select58499075)
}
__typedArg0 := Select58499075
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9294 {
tmp9252 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(Select58499075)
}
__typedArg0 := Select58499075
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

Select58469077 := tmp9252
_ = Select58469077

tmp9253 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select58499075)
}
__typedArg0 := Select58499075
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

Select58479078 := tmp9253
_ = Select58479078

tmp9292 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(Select58469077)
}
__typedArg0 := Select58469077
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9268 Obj

if True == tmp9292 {
tmp9290 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select58469077)
}
__typedArg0 := Select58469077
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9291 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9290)
}
__typedArg0 := tmp9290
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9270 Obj

if True == tmp9291 {
tmp9287 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select58469077)
}
__typedArg0 := Select58469077
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9288 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9287)
}
__typedArg0 := tmp9287
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9289 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9288)
}
__typedArg0 := tmp9288
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9272 Obj

if True == tmp9289 {
tmp9283 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select58469077)
}
__typedArg0 := Select58469077
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9284 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9283)
}
__typedArg0 := tmp9283
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9285 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9284)
}
__typedArg0 := tmp9284
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9286 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp9285)
}
__typedArg0 := Nil
__typedArg1 := tmp9285
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9274 Obj

if True == tmp9286 {
tmp9282 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, Select58479078)
}
__typedArg0 := Nil
__typedArg1 := Select58479078
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9276 Obj

if True == tmp9282 {
tmp9278 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select58469077)
}
__typedArg0 := Select58469077
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9279 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9278)
}
__typedArg0 := tmp9278
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9280 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp9281 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp9279, tmp9280)
}
__typedArg0 := tmp9279
__typedArg1 := tmp9280
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9277 Obj

if True == tmp9281 {
ifres9277 = True


} else {
ifres9277 = False


}

ifres9276 = ifres9277


} else {
ifres9276 = False


}

var ifres9275 Obj

if True == ifres9276 {
ifres9275 = True


} else {
ifres9275 = False


}

ifres9274 = ifres9275


} else {
ifres9274 = False


}

var ifres9273 Obj

if True == ifres9274 {
ifres9273 = True


} else {
ifres9273 = False


}

ifres9272 = ifres9273


} else {
ifres9272 = False


}

var ifres9271 Obj

if True == ifres9272 {
ifres9271 = True


} else {
ifres9271 = False


}

ifres9270 = ifres9271


} else {
ifres9270 = False


}

var ifres9269 Obj

if True == ifres9270 {
ifres9269 = True


} else {
ifres9269 = False


}

ifres9268 = ifres9269


} else {
ifres9268 = False


}

if True == ifres9268 {
tmp9254 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(Select58469077)
}
__typedArg0 := Select58469077
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9255 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select58469077)
}
__typedArg0 := Select58469077
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9256 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9255)
}
__typedArg0 := tmp9255
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9257 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select58469077)
}
__typedArg0 := Select58469077
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9258 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9257)
}
__typedArg0 := tmp9257
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9259 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_7, tmp9258)
}
__typedArg0 := sym_7
__typedArg1 := tmp9258
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
return PrimCons(tmp9256, tmp9260)
}
__typedArg0 := tmp9256
__typedArg1 := tmp9260
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9262 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9254, tmp9261)
}
__typedArg0 := tmp9254
__typedArg1 := tmp9261
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9263 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9262, Nil)
}
__typedArg0 := tmp9262
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9264 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1, tmp9263)
}
__typedArg0 := sym_1
__typedArg1 := tmp9263
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4cons_1form_1respect_1modes), tmp9264)
return


} else {
tmp9266 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, Select58479078)
}
__typedArg0 := Nil
__typedArg1 := Select58479078
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9266 {
__e.TailApply(PrimFunc(symshen_4cons_1form_1respect_1modes), Select58469077)
return
} else {
__e.TailApply(PrimFunc(symthaw), GoTo58449076)
return
}


}


} else {
__e.TailApply(PrimFunc(symthaw), GoTo58449076)
return
}


} else {
__e.TailApply(PrimFunc(symthaw), GoTo58449076)
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
__e.TailApply(PrimFunc(symthaw), GoTo58439073)
return
}


}, 1)

tmp9457 := Call(__e, ns2_1set, symshen_4macros, tmp9157)


_ = tmp9457

tmp9458 := MakeNative(func(__e *ControlFlow) {
V5850 := __e.Get(1)
_ = V5850
tmp9459 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("partial function shen.process-input+"))
}
__typedArg0 := MakeString("partial function shen.process-input+")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}, 0)

GoTo58519079 := tmp9459
_ = GoTo58519079

tmp9483 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5850)
}
__typedArg0 := V5850
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9483 {
tmp9460 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5850)
}
__typedArg0 := V5850
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

Select58569080 := tmp9460
_ = Select58569080

tmp9480 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5850)
}
__typedArg0 := V5850
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9481 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(syminput_7, tmp9480)
}
__typedArg0 := syminput_7
__typedArg1 := tmp9480
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9481 {
tmp9478 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(Select58569080)
}
__typedArg0 := Select58569080
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9478 {
tmp9461 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(Select58569080)
}
__typedArg0 := Select58569080
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

Select58549081 := tmp9461
_ = Select58549081

tmp9462 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select58569080)
}
__typedArg0 := Select58569080
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

Select58559082 := tmp9462
_ = Select58559082

tmp9476 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, Select58559082)
}
__typedArg0 := Nil
__typedArg1 := Select58559082
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9476 {
tmp9463 := Call(__e, PrimFunc(symshen_4rcons__form), Select58549081)


tmp9464 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstinput, Nil)
}
__typedArg0 := symstinput
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9465 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9464, Nil)
}
__typedArg0 := tmp9464
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9466 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9463, tmp9465)
}
__typedArg0 := tmp9463
__typedArg1 := tmp9465
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4input_1h_7, tmp9466)
}
__typedArg0 := symshen_4input_1h_7
__typedArg1 := tmp9466
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9474 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(Select58559082)
}
__typedArg0 := Select58559082
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9470 Obj

if True == tmp9474 {
tmp9472 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select58559082)
}
__typedArg0 := Select58559082
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9473 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp9472)
}
__typedArg0 := Nil
__typedArg1 := tmp9472
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9471 Obj

if True == tmp9473 {
ifres9471 = True


} else {
ifres9471 = False


}

ifres9470 = ifres9471


} else {
ifres9470 = False


}

if True == ifres9470 {
tmp9467 := Call(__e, PrimFunc(symshen_4rcons__form), Select58549081)


tmp9468 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9467, Select58559082)
}
__typedArg0 := tmp9467
__typedArg1 := Select58559082
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4input_1h_7, tmp9468)
}
__typedArg0 := symshen_4input_1h_7
__typedArg1 := tmp9468
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo58519079)
return
}


}


} else {
__e.TailApply(PrimFunc(symthaw), GoTo58519079)
return
}


} else {
__e.TailApply(PrimFunc(symthaw), GoTo58519079)
return
}


} else {
__e.TailApply(PrimFunc(symthaw), GoTo58519079)
return
}


}, 1)

tmp9484 := Call(__e, ns2_1set, symshen_4process_1input_7, tmp9458)


_ = tmp9484

tmp9485 := MakeNative(func(__e *ControlFlow) {
V5857 := __e.Get(1)
_ = V5857
tmp9486 := MakeNative(func(__e *ControlFlow) {
__e.Return(V5857)
return
}, 0)

GoTo58589083 := tmp9486
_ = GoTo58589083

tmp9518 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5857)
}
__typedArg0 := V5857
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9518 {
tmp9487 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5857)
}
__typedArg0 := V5857
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

Select58599084 := tmp9487
_ = Select58599084

tmp9488 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5857)
}
__typedArg0 := V5857
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

Select58609085 := tmp9488
_ = Select58609085

tmp9516 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_7, Select58599084)
}
__typedArg0 := sym_7
__typedArg1 := Select58599084
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9509 Obj

if True == tmp9516 {
tmp9515 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(Select58609085)
}
__typedArg0 := Select58609085
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9511 Obj

if True == tmp9515 {
tmp9513 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select58609085)
}
__typedArg0 := Select58609085
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9514 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp9513)
}
__typedArg0 := Nil
__typedArg1 := tmp9513
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9512 Obj

if True == tmp9514 {
ifres9512 = True


} else {
ifres9512 = False


}

ifres9511 = ifres9512


} else {
ifres9511 = False


}

var ifres9510 Obj

if True == ifres9511 {
ifres9510 = True


} else {
ifres9510 = False


}

ifres9509 = ifres9510


} else {
ifres9509 = False


}

if True == ifres9509 {
tmp9489 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(Select58609085)
}
__typedArg0 := Select58609085
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9490 := Call(__e, PrimFunc(symshen_4cons_1form_1respect_1modes), tmp9489)


tmp9491 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9490, Nil)
}
__typedArg0 := tmp9490
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_7, tmp9491)
}
__typedArg0 := sym_7
__typedArg1 := tmp9491
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9507 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_1, Select58599084)
}
__typedArg0 := sym_1
__typedArg1 := Select58599084
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9500 Obj

if True == tmp9507 {
tmp9506 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(Select58609085)
}
__typedArg0 := Select58609085
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9502 Obj

if True == tmp9506 {
tmp9504 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select58609085)
}
__typedArg0 := Select58609085
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9505 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp9504)
}
__typedArg0 := Nil
__typedArg1 := tmp9504
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9503 Obj

if True == tmp9505 {
ifres9503 = True


} else {
ifres9503 = False


}

ifres9502 = ifres9503


} else {
ifres9502 = False


}

var ifres9501 Obj

if True == ifres9502 {
ifres9501 = True


} else {
ifres9501 = False


}

ifres9500 = ifres9501


} else {
ifres9500 = False


}

if True == ifres9500 {
tmp9492 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(Select58609085)
}
__typedArg0 := Select58609085
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9493 := Call(__e, PrimFunc(symshen_4cons_1form_1respect_1modes), tmp9492)


tmp9494 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9493, Nil)
}
__typedArg0 := tmp9493
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1, tmp9494)
}
__typedArg0 := sym_1
__typedArg1 := tmp9494
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9495 := Call(__e, PrimFunc(symshen_4cons_1form_1respect_1modes), Select58599084)


tmp9496 := Call(__e, PrimFunc(symshen_4cons_1form_1respect_1modes), Select58609085)


tmp9497 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9496, Nil)
}
__typedArg0 := tmp9496
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9498 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9495, tmp9497)
}
__typedArg0 := tmp9495
__typedArg1 := tmp9497
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp9498)
}
__typedArg0 := symcons
__typedArg1 := tmp9498
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


}


} else {
__e.TailApply(PrimFunc(symthaw), GoTo58589083)
return
}


}, 1)

tmp9519 := Call(__e, ns2_1set, symshen_4cons_1form_1respect_1modes, tmp9485)


_ = tmp9519

tmp9520 := MakeNative(func(__e *ControlFlow) {
V5861 := __e.Get(1)
_ = V5861
V5862 := __e.Get(2)
_ = V5862
tmp9521 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symX, Nil)
}
__typedArg0 := symX
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9522 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_6, tmp9521)
}
__typedArg0 := sym_1_6
__typedArg1 := tmp9521
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9523 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symX, tmp9522)
}
__typedArg0 := symX
__typedArg1 := tmp9522
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

W58639086 := tmp9523
_ = W58639086

tmp9524 := Call(__e, PrimFunc(symappend), V5862, W58639086)


tmp9525 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5861, tmp9524)
}
__typedArg0 := V5861
__typedArg1 := tmp9524
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9526 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefine, tmp9525)
}
__typedArg0 := symdefine
__typedArg1 := tmp9525
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9527 := Call(__e, PrimFunc(symeval), tmp9526)


W58649087 := tmp9527
_ = W58649087

tmp9528 := Call(__e, PrimFunc(symfn), V5861)


tmp9529 := Call(__e, PrimFunc(symshen_4record_1macro), V5861, tmp9528)


W58659088 := tmp9529
_ = W58659088

__e.Return(V5861)
return


}, 2)

tmp9530 := Call(__e, ns2_1set, symshen_4process_1def, tmp9520)


_ = tmp9530

tmp9531 := MakeNative(func(__e *ControlFlow) {
V5866 := __e.Get(1)
_ = V5866
tmp9571 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5866)
}
__typedArg0 := V5866
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9545 Obj

if True == tmp9571 {
tmp9569 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5866)
}
__typedArg0 := V5866
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9570 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlet, tmp9569)
}
__typedArg0 := symlet
__typedArg1 := tmp9569
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9547 Obj

if True == tmp9570 {
tmp9567 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5866)
}
__typedArg0 := V5866
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9568 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9567)
}
__typedArg0 := tmp9567
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9549 Obj

if True == tmp9568 {
tmp9564 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5866)
}
__typedArg0 := V5866
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9565 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9564)
}
__typedArg0 := tmp9564
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9566 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9565)
}
__typedArg0 := tmp9565
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9551 Obj

if True == tmp9566 {
tmp9560 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5866)
}
__typedArg0 := V5866
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9561 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9560)
}
__typedArg0 := tmp9560
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9562 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9561)
}
__typedArg0 := tmp9561
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9563 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9562)
}
__typedArg0 := tmp9562
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9553 Obj

if True == tmp9563 {
tmp9555 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5866)
}
__typedArg0 := V5866
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9556 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9555)
}
__typedArg0 := tmp9555
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9557 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9556)
}
__typedArg0 := tmp9556
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9558 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9557)
}
__typedArg0 := tmp9557
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9559 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9558)
}
__typedArg0 := tmp9558
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9554 Obj

if True == tmp9559 {
ifres9554 = True


} else {
ifres9554 = False


}

ifres9553 = ifres9554


} else {
ifres9553 = False


}

var ifres9552 Obj

if True == ifres9553 {
ifres9552 = True


} else {
ifres9552 = False


}

ifres9551 = ifres9552


} else {
ifres9551 = False


}

var ifres9550 Obj

if True == ifres9551 {
ifres9550 = True


} else {
ifres9550 = False


}

ifres9549 = ifres9550


} else {
ifres9549 = False


}

var ifres9548 Obj

if True == ifres9549 {
ifres9548 = True


} else {
ifres9548 = False


}

ifres9547 = ifres9548


} else {
ifres9547 = False


}

var ifres9546 Obj

if True == ifres9547 {
ifres9546 = True


} else {
ifres9546 = False


}

ifres9545 = ifres9546


} else {
ifres9545 = False


}

if True == ifres9545 {
tmp9532 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5866)
}
__typedArg0 := V5866
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9533 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9532)
}
__typedArg0 := tmp9532
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9534 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5866)
}
__typedArg0 := V5866
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9535 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9534)
}
__typedArg0 := tmp9534
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9536 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9535)
}
__typedArg0 := tmp9535
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9537 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5866)
}
__typedArg0 := V5866
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9538 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9537)
}
__typedArg0 := tmp9537
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9539 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9538)
}
__typedArg0 := tmp9538
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9540 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp9539)
}
__typedArg0 := symlet
__typedArg1 := tmp9539
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9541 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9540, Nil)
}
__typedArg0 := tmp9540
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9542 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9536, tmp9541)
}
__typedArg0 := tmp9536
__typedArg1 := tmp9541
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9543 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9533, tmp9542)
}
__typedArg0 := tmp9533
__typedArg1 := tmp9542
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp9543)
}
__typedArg0 := symlet
__typedArg1 := tmp9543
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return(V5866)
return
}


}, 1)

tmp9572 := Call(__e, ns2_1set, symshen_4process_1let, tmp9531)


_ = tmp9572

tmp9573 := MakeNative(func(__e *ControlFlow) {
V5867 := __e.Get(1)
_ = V5867
tmp9574 := MakeNative(func(__e *ControlFlow) {
__e.Return(V5867)
return
}, 0)

GoTo58699089 := tmp9574
_ = GoTo58699089

tmp9604 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5867)
}
__typedArg0 := V5867
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9604 {
tmp9575 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5867)
}
__typedArg0 := V5867
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

Select58769090 := tmp9575
_ = Select58769090

tmp9601 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5867)
}
__typedArg0 := V5867
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9602 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_8s, tmp9601)
}
__typedArg0 := sym_8s
__typedArg1 := tmp9601
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9602 {
tmp9599 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(Select58769090)
}
__typedArg0 := Select58769090
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9599 {
tmp9576 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(Select58769090)
}
__typedArg0 := Select58769090
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

Select58749091 := tmp9576
_ = Select58749091

tmp9577 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select58769090)
}
__typedArg0 := Select58769090
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

Select58759092 := tmp9577
_ = Select58759092

tmp9597 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(Select58759092)
}
__typedArg0 := Select58759092
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9597 {
tmp9578 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select58759092)
}
__typedArg0 := Select58759092
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

Select58739093 := tmp9578
_ = Select58739093

tmp9595 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(Select58739093)
}
__typedArg0 := Select58739093
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9595 {
tmp9579 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8s, Select58759092)
}
__typedArg0 := sym_8s
__typedArg1 := Select58759092
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9580 := Call(__e, PrimFunc(symshen_4process_1_8s), tmp9579)


tmp9581 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9580, Nil)
}
__typedArg0 := tmp9580
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9582 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Select58749091, tmp9581)
}
__typedArg0 := Select58749091
__typedArg1 := tmp9581
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8s, tmp9582)
}
__typedArg0 := sym_8s
__typedArg1 := tmp9582
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9593 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, Select58739093)
}
__typedArg0 := Nil
__typedArg1 := Select58739093
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9590 Obj

if True == tmp9593 {
tmp9592 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_2) {
return PrimIsString(Select58749091)
}
__typedArg0 := Select58749091
return Call(__e, PrimFunc(symstring_2), __typedArg0)
})()

var ifres9591 Obj

if True == tmp9592 {
ifres9591 = True


} else {
ifres9591 = False


}

ifres9590 = ifres9591


} else {
ifres9590 = False


}

if True == ifres9590 {
tmp9583 := Call(__e, PrimFunc(symexplode), Select58749091)


W58689094 := tmp9583
_ = W58689094

tmp9587 := Call(__e, PrimFunc(symlength), W58689094)


if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_6) {
__typedN0, __typedOK0 := TypedFloat64(tmp9587)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_6) {
return TypedMaterializeBoolean((__typedN0 > __typedN1))
}}
__typedArg0 := tmp9587
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_6), __typedArg0, __typedArg1)
})() {
tmp9584 := Call(__e, PrimFunc(symappend), W58689094, Select58759092)


tmp9585 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8s, tmp9584)
}
__typedArg0 := sym_8s
__typedArg1 := tmp9584
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4process_1_8s), tmp9585)
return


} else {
__e.Return(V5867)
return
}


} else {
__e.TailApply(PrimFunc(symthaw), GoTo58699089)
return
}


}


} else {
__e.TailApply(PrimFunc(symthaw), GoTo58699089)
return
}


} else {
__e.TailApply(PrimFunc(symthaw), GoTo58699089)
return
}


} else {
__e.TailApply(PrimFunc(symthaw), GoTo58699089)
return
}


} else {
__e.TailApply(PrimFunc(symthaw), GoTo58699089)
return
}


}, 1)

tmp9605 := Call(__e, ns2_1set, symshen_4process_1_8s, tmp9573)


_ = tmp9605

tmp9606 := MakeNative(func(__e *ControlFlow) {
V5877 := __e.Get(1)
_ = V5877
V5878 := __e.Get(2)
_ = V5878
tmp9607 := Call(__e, PrimFunc(symshen_4intern_1type), V5877)


W58799095 := tmp9607
_ = W58799095

tmp9608 := MakeNative(func(__e *ControlFlow) {
Z5881 := __e.Get(1)
_ = Z5881
__e.TailApply(PrimFunc(symshen_4_5datatype_6), Z5881)
return
}, 1)

tmp9609 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W58799095, V5878)
}
__typedArg0 := W58799095
__typedArg1 := V5878
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9610 := Call(__e, PrimFunc(symcompile), tmp9608, tmp9609)


W58809096 := tmp9610
_ = W58809096

__e.Return(W58799095)
return


}, 2)

tmp9611 := Call(__e, ns2_1set, symshen_4process_1datatype, tmp9606)


_ = tmp9611

tmp9612 := MakeNative(func(__e *ControlFlow) {
V5882 := __e.Get(1)
_ = V5882
tmp9613 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(V5882)
}
__typedArg0 := V5882
return Call(__e, PrimFunc(symstr), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(tmp9613)
__typedS1, __typedOK1 := TypedString(MakeString("#type"))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := tmp9613
__typedArg1 := MakeString("#type")
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(tmp9613)
__typedS1, __typedOK1 := TypedString(MakeString("#type"))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := tmp9613
__typedArg1 := MakeString("#type")
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symintern), __typedArg0)
})())
return


}, 1)

tmp9615 := Call(__e, ns2_1set, symshen_4intern_1type, tmp9612)


_ = tmp9615

tmp9616 := MakeNative(func(__e *ControlFlow) {
V5883 := __e.Get(1)
_ = V5883
tmp9617 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsynonyms_d)
}
__typedArg0 := symshen_4_dsynonyms_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp9618 := Call(__e, PrimFunc(symappend), V5883, tmp9617)


tmp9619 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsynonyms_d, tmp9618)
}
__typedArg0 := symshen_4_dsynonyms_d
__typedArg1 := tmp9618
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4synonyms_1h), tmp9619)
return


}, 1)

tmp9620 := Call(__e, ns2_1set, symshen_4process_1synonyms, tmp9616)


_ = tmp9620

tmp9621 := MakeNative(func(__e *ControlFlow) {
V5884 := __e.Get(1)
_ = V5884
tmp9622 := MakeNative(func(__e *ControlFlow) {
Z5886 := __e.Get(1)
_ = Z5886
__e.TailApply(PrimFunc(symshen_4curry_1type), Z5886)
return
}, 1)

tmp9623 := Call(__e, PrimFunc(symmap), tmp9622, V5884)


W58859097 := tmp9623
_ = W58859097

tmp9624 := Call(__e, PrimFunc(symshen_4compile_1synonyms), W58859097)


tmp9625 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4demod, tmp9624)
}
__typedArg0 := symshen_4demod
__typedArg1 := tmp9624
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9626 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefine, tmp9625)
}
__typedArg0 := symdefine
__typedArg1 := tmp9625
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9627 := Call(__e, PrimFunc(symeval), tmp9626)


W58879098 := tmp9627
_ = W58879098

__e.Return(symsynonyms)
return


}, 1)

tmp9628 := Call(__e, ns2_1set, symshen_4synonyms_1h, tmp9621)


_ = tmp9628

tmp9629 := MakeNative(func(__e *ControlFlow) {
V5890 := __e.Get(1)
_ = V5890
tmp9650 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V5890)
}
__typedArg0 := Nil
__typedArg1 := V5890
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9650 {
tmp9630 := Call(__e, PrimFunc(symgensym), symX)


W58919099 := tmp9630
_ = W58919099

tmp9631 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W58919099, Nil)
}
__typedArg0 := W58919099
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9632 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_6, tmp9631)
}
__typedArg0 := sym_1_6
__typedArg1 := tmp9631
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W58919099, tmp9632)
}
__typedArg0 := W58919099
__typedArg1 := tmp9632
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9648 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5890)
}
__typedArg0 := V5890
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9644 Obj

if True == tmp9648 {
tmp9646 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5890)
}
__typedArg0 := V5890
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9647 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9646)
}
__typedArg0 := tmp9646
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9645 Obj

if True == tmp9647 {
ifres9645 = True


} else {
ifres9645 = False


}

ifres9644 = ifres9645


} else {
ifres9644 = False


}

if True == ifres9644 {
tmp9633 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5890)
}
__typedArg0 := V5890
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9634 := Call(__e, PrimFunc(symshen_4rcons__form), tmp9633)


tmp9635 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5890)
}
__typedArg0 := V5890
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9636 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9635)
}
__typedArg0 := tmp9635
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9637 := Call(__e, PrimFunc(symshen_4rcons__form), tmp9636)


tmp9638 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5890)
}
__typedArg0 := V5890
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9639 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9638)
}
__typedArg0 := tmp9638
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9640 := Call(__e, PrimFunc(symshen_4compile_1synonyms), tmp9639)


tmp9641 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9637, tmp9640)
}
__typedArg0 := tmp9637
__typedArg1 := tmp9640
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9642 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_6, tmp9641)
}
__typedArg0 := sym_1_6
__typedArg1 := tmp9641
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9634, tmp9642)
}
__typedArg0 := tmp9634
__typedArg1 := tmp9642
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

tmp9651 := Call(__e, ns2_1set, symshen_4compile_1synonyms, tmp9629)


_ = tmp9651

tmp9652 := MakeNative(func(__e *ControlFlow) {
V5892 := __e.Get(1)
_ = V5892
tmp9653 := MakeNative(func(__e *ControlFlow) {
__e.Return(V5892)
return
}, 0)

GoTo58939100 := tmp9653
_ = GoTo58939100

tmp9677 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5892)
}
__typedArg0 := V5892
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9677 {
tmp9654 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5892)
}
__typedArg0 := V5892
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

Select59009101 := tmp9654
_ = Select59009101

tmp9674 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5892)
}
__typedArg0 := V5892
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9675 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_c_4, tmp9674)
}
__typedArg0 := sym_c_4
__typedArg1 := tmp9674
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9675 {
tmp9672 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(Select59009101)
}
__typedArg0 := Select59009101
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9672 {
tmp9655 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(Select59009101)
}
__typedArg0 := Select59009101
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

Select58989102 := tmp9655
_ = Select58989102

tmp9656 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select59009101)
}
__typedArg0 := Select59009101
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

Select58999103 := tmp9656
_ = Select58999103

tmp9670 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(Select58999103)
}
__typedArg0 := Select58999103
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9670 {
tmp9657 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select58999103)
}
__typedArg0 := Select58999103
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

Select58979104 := tmp9657
_ = Select58979104

tmp9668 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(Select58979104)
}
__typedArg0 := Select58979104
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9668 {
tmp9658 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_c_4, Select58999103)
}
__typedArg0 := sym_c_4
__typedArg1 := Select58999103
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9659 := Call(__e, PrimFunc(symshen_4process_1lambda), tmp9658)


tmp9660 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9659, Nil)
}
__typedArg0 := tmp9659
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9661 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Select58989102, tmp9660)
}
__typedArg0 := Select58989102
__typedArg1 := tmp9660
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp9661)
}
__typedArg0 := symlambda
__typedArg1 := tmp9661
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9666 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, Select58979104)
}
__typedArg0 := Nil
__typedArg1 := Select58979104
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9666 {
tmp9664 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(Select58989102)
}
__typedArg0 := Select58989102
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

if True == tmp9664 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, Select59009101)
}
__typedArg0 := symlambda
__typedArg1 := Select59009101
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return
} else {
tmp9662 := Call(__e, PrimFunc(symshen_4app), Select58989102, MakeString(" is not a variable\n"), symshen_4s)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(tmp9662)
}
__typedArg0 := tmp9662
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}


} else {
__e.TailApply(PrimFunc(symthaw), GoTo58939100)
return
}


}


} else {
__e.TailApply(PrimFunc(symthaw), GoTo58939100)
return
}


} else {
__e.TailApply(PrimFunc(symthaw), GoTo58939100)
return
}


} else {
__e.TailApply(PrimFunc(symthaw), GoTo58939100)
return
}


} else {
__e.TailApply(PrimFunc(symthaw), GoTo58939100)
return
}


}, 1)

tmp9678 := Call(__e, ns2_1set, symshen_4process_1lambda, tmp9652)


_ = tmp9678

tmp9679 := MakeNative(func(__e *ControlFlow) {
V5903 := __e.Get(1)
_ = V5903
tmp9680 := MakeNative(func(__e *ControlFlow) {
__e.Return(V5903)
return
}, 0)

GoTo59049105 := tmp9680
_ = GoTo59049105

tmp9714 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5903)
}
__typedArg0 := V5903
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9714 {
tmp9681 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5903)
}
__typedArg0 := V5903
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

Select59129106 := tmp9681
_ = Select59129106

tmp9711 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5903)
}
__typedArg0 := V5903
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9712 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcases, tmp9711)
}
__typedArg0 := symcases
__typedArg1 := tmp9711
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9712 {
tmp9709 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(Select59129106)
}
__typedArg0 := Select59129106
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9709 {
tmp9682 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(Select59129106)
}
__typedArg0 := Select59129106
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

Select59109107 := tmp9682
_ = Select59109107

tmp9683 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select59129106)
}
__typedArg0 := Select59129106
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

Select59119108 := tmp9683
_ = Select59119108

tmp9707 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(True, Select59109107)
}
__typedArg0 := True
__typedArg1 := Select59109107
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9704 Obj

if True == tmp9707 {
tmp9706 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(Select59119108)
}
__typedArg0 := Select59119108
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9705 Obj

if True == tmp9706 {
ifres9705 = True


} else {
ifres9705 = False


}

ifres9704 = ifres9705


} else {
ifres9704 = False


}

if True == ifres9704 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(Select59119108)
}
__typedArg0 := Select59119108
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return
} else {
tmp9684 := MakeNative(func(__e *ControlFlow) {
tmp9686 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, Select59119108)
}
__typedArg0 := Nil
__typedArg1 := Select59119108
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9686 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("error: odd number of case elements\n"))
}
__typedArg0 := MakeString("error: odd number of case elements\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
} else {
__e.TailApply(PrimFunc(symthaw), GoTo59049105)
return
}


}, 0)

GoTo59079109 := tmp9684
_ = GoTo59079109

tmp9702 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(Select59119108)
}
__typedArg0 := Select59119108
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9702 {
tmp9687 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(Select59119108)
}
__typedArg0 := Select59119108
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

Select59089110 := tmp9687
_ = Select59089110

tmp9688 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select59119108)
}
__typedArg0 := Select59119108
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

Select59099111 := tmp9688
_ = Select59099111

tmp9700 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, Select59099111)
}
__typedArg0 := Nil
__typedArg1 := Select59099111
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9700 {
tmp9689 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeString("error: cases exhausted"), Nil)
}
__typedArg0 := MakeString("error: cases exhausted")
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9690 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsimple_1error, tmp9689)
}
__typedArg0 := symsimple_1error
__typedArg1 := tmp9689
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9691 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9690, Nil)
}
__typedArg0 := tmp9690
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9692 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Select59089110, tmp9691)
}
__typedArg0 := Select59089110
__typedArg1 := tmp9691
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9693 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Select59109107, tmp9692)
}
__typedArg0 := Select59109107
__typedArg1 := tmp9692
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp9693)
}
__typedArg0 := symif
__typedArg1 := tmp9693
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9694 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcases, Select59099111)
}
__typedArg0 := symcases
__typedArg1 := Select59099111
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9695 := Call(__e, PrimFunc(symshen_4process_1cases), tmp9694)


tmp9696 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9695, Nil)
}
__typedArg0 := tmp9695
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9697 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Select59089110, tmp9696)
}
__typedArg0 := Select59089110
__typedArg1 := tmp9696
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9698 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Select59109107, tmp9697)
}
__typedArg0 := Select59109107
__typedArg1 := tmp9697
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp9698)
}
__typedArg0 := symif
__typedArg1 := tmp9698
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


} else {
__e.TailApply(PrimFunc(symthaw), GoTo59079109)
return
}


}


} else {
__e.TailApply(PrimFunc(symthaw), GoTo59049105)
return
}


} else {
__e.TailApply(PrimFunc(symthaw), GoTo59049105)
return
}


} else {
__e.TailApply(PrimFunc(symthaw), GoTo59049105)
return
}


}, 1)

tmp9715 := Call(__e, ns2_1set, symshen_4process_1cases, tmp9679)


_ = tmp9715

tmp9716 := MakeNative(func(__e *ControlFlow) {
V5913 := __e.Get(1)
_ = V5913
tmp9717 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symrun, Nil)
}
__typedArg0 := symrun
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9718 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symget_1time, tmp9717)
}
__typedArg0 := symget_1time
__typedArg1 := tmp9717
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9719 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symrun, Nil)
}
__typedArg0 := symrun
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9720 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symget_1time, tmp9719)
}
__typedArg0 := symget_1time
__typedArg1 := tmp9719
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9721 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symStart, Nil)
}
__typedArg0 := symStart
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9722 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symFinish, tmp9721)
}
__typedArg0 := symFinish
__typedArg1 := tmp9721
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9723 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1, tmp9722)
}
__typedArg0 := sym_1
__typedArg1 := tmp9722
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9724 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symTime, Nil)
}
__typedArg0 := symTime
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9725 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp9724)
}
__typedArg0 := symstr
__typedArg1 := tmp9724
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9726 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeString(" secs\n"), Nil)
}
__typedArg0 := MakeString(" secs\n")
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9727 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9725, tmp9726)
}
__typedArg0 := tmp9725
__typedArg1 := tmp9726
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9728 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcn, tmp9727)
}
__typedArg0 := symcn
__typedArg1 := tmp9727
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9729 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9728, Nil)
}
__typedArg0 := tmp9728
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9730 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeString("\nrun time: "), tmp9729)
}
__typedArg0 := MakeString("\nrun time: ")
__typedArg1 := tmp9729
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9731 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcn, tmp9730)
}
__typedArg0 := symcn
__typedArg1 := tmp9730
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9732 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstoutput, Nil)
}
__typedArg0 := symstoutput
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9733 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9732, Nil)
}
__typedArg0 := tmp9732
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9734 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9731, tmp9733)
}
__typedArg0 := tmp9731
__typedArg1 := tmp9733
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9735 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympr, tmp9734)
}
__typedArg0 := sympr
__typedArg1 := tmp9734
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9736 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symResult, Nil)
}
__typedArg0 := symResult
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9737 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9735, tmp9736)
}
__typedArg0 := tmp9735
__typedArg1 := tmp9736
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9738 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symMessage, tmp9737)
}
__typedArg0 := symMessage
__typedArg1 := tmp9737
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9739 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9723, tmp9738)
}
__typedArg0 := tmp9723
__typedArg1 := tmp9738
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9740 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symTime, tmp9739)
}
__typedArg0 := symTime
__typedArg1 := tmp9739
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9741 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9720, tmp9740)
}
__typedArg0 := tmp9720
__typedArg1 := tmp9740
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9742 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symFinish, tmp9741)
}
__typedArg0 := symFinish
__typedArg1 := tmp9741
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9743 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5913, tmp9742)
}
__typedArg0 := V5913
__typedArg1 := tmp9742
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9744 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symResult, tmp9743)
}
__typedArg0 := symResult
__typedArg1 := tmp9743
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9745 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9718, tmp9744)
}
__typedArg0 := tmp9718
__typedArg1 := tmp9744
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9746 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symStart, tmp9745)
}
__typedArg0 := symStart
__typedArg1 := tmp9745
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp9746)
}
__typedArg0 := symlet
__typedArg1 := tmp9746
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp9747 := Call(__e, ns2_1set, symshen_4process_1time, tmp9716)


_ = tmp9747

tmp9748 := MakeNative(func(__e *ControlFlow) {
V5914 := __e.Get(1)
_ = V5914
tmp9774 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5914)
}
__typedArg0 := V5914
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9759 Obj

if True == tmp9774 {
tmp9772 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5914)
}
__typedArg0 := V5914
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9773 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9772)
}
__typedArg0 := tmp9772
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9761 Obj

if True == tmp9773 {
tmp9769 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5914)
}
__typedArg0 := V5914
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9770 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9769)
}
__typedArg0 := tmp9769
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9771 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9770)
}
__typedArg0 := tmp9770
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9763 Obj

if True == tmp9771 {
tmp9765 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5914)
}
__typedArg0 := V5914
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9766 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9765)
}
__typedArg0 := tmp9765
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9767 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9766)
}
__typedArg0 := tmp9766
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9768 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9767)
}
__typedArg0 := tmp9767
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9764 Obj

if True == tmp9768 {
ifres9764 = True


} else {
ifres9764 = False


}

ifres9763 = ifres9764


} else {
ifres9763 = False


}

var ifres9762 Obj

if True == ifres9763 {
ifres9762 = True


} else {
ifres9762 = False


}

ifres9761 = ifres9762


} else {
ifres9761 = False


}

var ifres9760 Obj

if True == ifres9761 {
ifres9760 = True


} else {
ifres9760 = False


}

ifres9759 = ifres9760


} else {
ifres9759 = False


}

if True == ifres9759 {
tmp9749 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5914)
}
__typedArg0 := V5914
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9750 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5914)
}
__typedArg0 := V5914
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9751 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9750)
}
__typedArg0 := tmp9750
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9752 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5914)
}
__typedArg0 := V5914
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9753 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5914)
}
__typedArg0 := V5914
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9754 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9753)
}
__typedArg0 := tmp9753
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9755 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9752, tmp9754)
}
__typedArg0 := tmp9752
__typedArg1 := tmp9754
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9756 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9755, Nil)
}
__typedArg0 := tmp9755
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9757 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9751, tmp9756)
}
__typedArg0 := tmp9751
__typedArg1 := tmp9756
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9749, tmp9757)
}
__typedArg0 := tmp9749
__typedArg1 := tmp9757
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return(V5914)
return
}


}, 1)

tmp9775 := Call(__e, ns2_1set, symshen_4process_1assoc, tmp9748)


_ = tmp9775

tmp9776 := MakeNative(func(__e *ControlFlow) {
V5915 := __e.Get(1)
_ = V5915
tmp9777 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(V5915)
}
__typedArg0 := V5915
return Call(__e, PrimFunc(symstr), __typedArg0)
})()

tmp9778 := Call(__e, PrimFunc(symshen_4mu_1h), tmp9777)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(tmp9778)
}
__typedArg0 := tmp9778
return Call(__e, PrimFunc(symintern), __typedArg0)
})())
return


}, 1)

tmp9779 := Call(__e, ns2_1set, symshen_4make_1uppercase, tmp9776)


_ = tmp9779

tmp9780 := MakeNative(func(__e *ControlFlow) {
V5916 := __e.Get(1)
_ = V5916
tmp9796 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString(""), V5916)
}
__typedArg0 := MakeString("")
__typedArg1 := V5916
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9796 {
__e.Return(MakeString(""))
return
} else {
tmp9794 := Call(__e, PrimFunc(symshen_4_7string_2), V5916)


if True == tmp9794 {
tmp9781 := Call(__e, PrimFunc(symhdstr), V5916)


tmp9782 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_1_6n) {
return PrimStringToNumber(tmp9781)
}
__typedArg0 := tmp9781
return Call(__e, PrimFunc(symstring_1_6n), __typedArg0)
})()

W59179112 := tmp9782
_ = W59179112

W59189113 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(W59179112)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(32))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := W59179112
__typedArg1 := MakeNumber(32)
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})()
_ = W59189113

tmp9790 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_6_a) {
__typedN0, __typedOK0 := TypedFloat64(W59179112)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(97))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_6_a) {
return TypedMaterializeBoolean((__typedN0 >= __typedN1))
}}
__typedArg0 := W59179112
__typedArg1 := MakeNumber(97)
return Call(__e, PrimFunc(sym_6_a), __typedArg0, __typedArg1)
})()

var ifres9787 Obj

if True == tmp9790 {
tmp9789 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_a) {
__typedN0, __typedOK0 := TypedFloat64(W59179112)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(122))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_5_a) {
return TypedMaterializeBoolean((__typedN0 <= __typedN1))
}}
__typedArg0 := W59179112
__typedArg1 := MakeNumber(122)
return Call(__e, PrimFunc(sym_5_a), __typedArg0, __typedArg1)
})()

var ifres9788 Obj

if True == tmp9789 {
ifres9788 = True


} else {
ifres9788 = False


}

ifres9787 = ifres9788


} else {
ifres9787 = False


}

var ifres9784 Obj

if True == ifres9787 {
tmp9785 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symn_1_6string) {
return PrimNumberToString(W59189113)
}
__typedArg0 := W59189113
return Call(__e, PrimFunc(symn_1_6string), __typedArg0)
})()

ifres9784 = tmp9785


} else {
tmp9786 := Call(__e, PrimFunc(symhdstr), V5916)


ifres9784 = tmp9786


}

W59199114 := ifres9784
_ = W59199114

tmp9792 := Call(__e, PrimFunc(symshen_4mu_1h), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V5916)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V5916
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


__e.TailApply(PrimFunc(sym_8s), W59199114, tmp9792)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("partial function shen.mu-h"))
}
__typedArg0 := MakeString("partial function shen.mu-h")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 1)

tmp9797 := Call(__e, ns2_1set, symshen_4mu_1h, tmp9780)


_ = tmp9797

tmp9798 := MakeNative(func(__e *ControlFlow) {
V5920 := __e.Get(1)
_ = V5920
V5921 := __e.Get(2)
_ = V5921
tmp9799 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dmacros_d)
}
__typedArg0 := sym_dmacros_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp9800 := Call(__e, PrimFunc(symshen_4update_1assoc), V5920, V5921, tmp9799)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(sym_dmacros_d, tmp9800)
}
__typedArg0 := sym_dmacros_d
__typedArg1 := tmp9800
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return


}, 2)

tmp9801 := Call(__e, ns2_1set, symshen_4record_1macro, tmp9798)


_ = tmp9801

tmp9802 := MakeNative(func(__e *ControlFlow) {
V5931 := __e.Get(1)
_ = V5931
V5932 := __e.Get(2)
_ = V5932
V5933 := __e.Get(3)
_ = V5933
tmp9819 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V5933)
}
__typedArg0 := Nil
__typedArg1 := V5933
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9819 {
tmp9803 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5931, V5932)
}
__typedArg0 := V5931
__typedArg1 := V5932
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9803, Nil)
}
__typedArg0 := tmp9803
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9804 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.update-assoc"))
}
__typedArg0 := MakeString("implementation error in shen.update-assoc")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}, 0)

GoTo59349115 := tmp9804
_ = GoTo59349115

tmp9817 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5933)
}
__typedArg0 := V5933
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9817 {
tmp9805 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5933)
}
__typedArg0 := V5933
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

Select59359116 := tmp9805
_ = Select59359116

tmp9806 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5933)
}
__typedArg0 := V5933
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

Select59369117 := tmp9806
_ = Select59369117

tmp9815 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(Select59359116)
}
__typedArg0 := Select59359116
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9811 Obj

if True == tmp9815 {
tmp9813 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(Select59359116)
}
__typedArg0 := Select59359116
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9814 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V5931, tmp9813)
}
__typedArg0 := V5931
__typedArg1 := tmp9813
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9812 Obj

if True == tmp9814 {
ifres9812 = True


} else {
ifres9812 = False


}

ifres9811 = ifres9812


} else {
ifres9811 = False


}

if True == ifres9811 {
tmp9807 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(Select59359116)
}
__typedArg0 := Select59359116
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9808 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9807, V5932)
}
__typedArg0 := tmp9807
__typedArg1 := V5932
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9808, Select59369117)
}
__typedArg0 := tmp9808
__typedArg1 := Select59369117
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9809 := Call(__e, PrimFunc(symshen_4update_1assoc), V5931, V5932, Select59369117)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Select59359116, tmp9809)
}
__typedArg0 := Select59359116
__typedArg1 := tmp9809
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


} else {
__e.TailApply(PrimFunc(symthaw), GoTo59349115)
return
}


}


}, 3)

tmp9820 := Call(__e, ns2_1set, symshen_4update_1assoc, tmp9802)


_ = tmp9820

tmp9821 := MakeNative(func(__e *ControlFlow) {
tmp9829 := Call(__e, PrimFunc(symstinput))


tmp9830 := Call(__e, PrimFunc(symshen_4char_1stinput_2), tmp9829)


if True == tmp9830 {
tmp9822 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstinput, Nil)
}
__typedArg0 := symstinput
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9823 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9822, Nil)
}
__typedArg0 := tmp9822
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9824 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4read_1unit_1string, tmp9823)
}
__typedArg0 := symshen_4read_1unit_1string
__typedArg1 := tmp9823
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9825 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9824, Nil)
}
__typedArg0 := tmp9824
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring_1_6n, tmp9825)
}
__typedArg0 := symstring_1_6n
__typedArg1 := tmp9825
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9826 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstinput, Nil)
}
__typedArg0 := symstinput
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9827 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9826, Nil)
}
__typedArg0 := tmp9826
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1byte, tmp9827)
}
__typedArg0 := symread_1byte
__typedArg1 := tmp9827
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


}, 0)

tmp9831 := Call(__e, ns2_1set, symshen_4process_1read_1byte, tmp9821)


_ = tmp9831

tmp9832 := MakeNative(func(__e *ControlFlow) {
V5937 := __e.Get(1)
_ = V5937
tmp9833 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4prolog_1vector, Nil)
}
__typedArg0 := symshen_4prolog_1vector
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

W59389118 := tmp9833
_ = W59389118

tmp9834 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(0), Nil)
}
__typedArg0 := MakeNumber(0)
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9835 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp9834)
}
__typedArg0 := symvector
__typedArg1 := tmp9834
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9836 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9835, Nil)
}
__typedArg0 := tmp9835
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9837 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(0), tmp9836)
}
__typedArg0 := MakeNumber(0)
__typedArg1 := tmp9836
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9838 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(True, tmp9837)
}
__typedArg0 := True
__typedArg1 := tmp9837
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9839 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8v, tmp9838)
}
__typedArg0 := sym_8v
__typedArg1 := tmp9838
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

W59399119 := tmp9839
_ = W59399119

W59409120 := MakeNumber(0)
_ = W59409120

tmp9840 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(True, Nil)
}
__typedArg0 := True
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9841 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfreeze, tmp9840)
}
__typedArg0 := symfreeze
__typedArg1 := tmp9840
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

W59419121 := tmp9841
_ = W59419121

tmp9842 := MakeNative(func(__e *ControlFlow) {
Z5943 := __e.Get(1)
_ = Z5943
__e.TailApply(PrimFunc(symshen_4_5body_6), Z5943)
return
}, 1)

tmp9843 := Call(__e, PrimFunc(symcompile), tmp9842, V5937)


W59429122 := tmp9843
_ = W59429122

tmp9844 := Call(__e, PrimFunc(symshen_4received), V5937)


W59449123 := tmp9844
_ = W59449123

tmp9845 := Call(__e, PrimFunc(symgensym), symV)


W59459124 := tmp9845
_ = W59459124

tmp9846 := Call(__e, PrimFunc(symgensym), symL)


W59469125 := tmp9846
_ = W59469125

tmp9847 := Call(__e, PrimFunc(symgensym), symK)


W59479126 := tmp9847
_ = W59479126

tmp9848 := Call(__e, PrimFunc(symgensym), symC)


W59489127 := tmp9848
_ = W59489127

tmp9849 := Call(__e, PrimFunc(symshen_4continue), W59449123, W59429122, W59459124, W59469125, W59479126, W59489127)


tmp9850 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9849, Nil)
}
__typedArg0 := tmp9849
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9851 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W59489127, tmp9850)
}
__typedArg0 := W59489127
__typedArg1 := tmp9850
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9852 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp9851)
}
__typedArg0 := symlambda
__typedArg1 := tmp9851
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9853 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9852, Nil)
}
__typedArg0 := tmp9852
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9854 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W59479126, tmp9853)
}
__typedArg0 := W59479126
__typedArg1 := tmp9853
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9855 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp9854)
}
__typedArg0 := symlambda
__typedArg1 := tmp9854
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9856 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9855, Nil)
}
__typedArg0 := tmp9855
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9857 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W59469125, tmp9856)
}
__typedArg0 := W59469125
__typedArg1 := tmp9856
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9858 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp9857)
}
__typedArg0 := symlambda
__typedArg1 := tmp9857
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9859 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9858, Nil)
}
__typedArg0 := tmp9858
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9860 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W59459124, tmp9859)
}
__typedArg0 := W59459124
__typedArg1 := tmp9859
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9861 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp9860)
}
__typedArg0 := symlambda
__typedArg1 := tmp9860
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

W59499128 := tmp9861
_ = W59499128

tmp9862 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W59419121, Nil)
}
__typedArg0 := W59419121
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9863 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W59409120, tmp9862)
}
__typedArg0 := W59409120
__typedArg1 := tmp9862
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9864 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W59399119, tmp9863)
}
__typedArg0 := W59399119
__typedArg1 := tmp9863
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9865 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W59389118, tmp9864)
}
__typedArg0 := W59389118
__typedArg1 := tmp9864
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W59499128, tmp9865)
}
__typedArg0 := W59499128
__typedArg1 := tmp9865
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp9866 := Call(__e, ns2_1set, symshen_4call_1prolog, tmp9832)


_ = tmp9866

tmp9867 := MakeNative(func(__e *ControlFlow) {
V5952 := __e.Get(1)
_ = V5952
tmp9868 := MakeNative(func(__e *ControlFlow) {
__e.Return(Nil)
return
}, 0)

GoTo59539129 := tmp9868
_ = GoTo59539129

tmp9883 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5952)
}
__typedArg0 := V5952
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9883 {
tmp9869 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5952)
}
__typedArg0 := V5952
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

Select59549130 := tmp9869
_ = Select59549130

tmp9870 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5952)
}
__typedArg0 := V5952
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

Select59559131 := tmp9870
_ = Select59559131

tmp9881 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symreceive, Select59549130)
}
__typedArg0 := symreceive
__typedArg1 := Select59549130
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9874 Obj

if True == tmp9881 {
tmp9880 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(Select59559131)
}
__typedArg0 := Select59559131
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9876 Obj

if True == tmp9880 {
tmp9878 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select59559131)
}
__typedArg0 := Select59559131
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9879 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp9878)
}
__typedArg0 := Nil
__typedArg1 := tmp9878
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9877 Obj

if True == tmp9879 {
ifres9877 = True


} else {
ifres9877 = False


}

ifres9876 = ifres9877


} else {
ifres9876 = False


}

var ifres9875 Obj

if True == ifres9876 {
ifres9875 = True


} else {
ifres9875 = False


}

ifres9874 = ifres9875


} else {
ifres9874 = False


}

if True == ifres9874 {
__e.Return(Select59559131)
return
} else {
tmp9871 := Call(__e, PrimFunc(symshen_4received), Select59549130)


tmp9872 := Call(__e, PrimFunc(symshen_4received), Select59559131)


__e.TailApply(PrimFunc(symunion), tmp9871, tmp9872)
return


}


} else {
__e.TailApply(PrimFunc(symthaw), GoTo59539129)
return
}


}, 1)

tmp9884 := Call(__e, ns2_1set, symshen_4received, tmp9867)


_ = tmp9884

tmp9885 := MakeNative(func(__e *ControlFlow) {
tmp9886 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dprolog_1memory_d)
}
__typedArg0 := symshen_4_dprolog_1memory_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp9887 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symabsvector) {
return PrimAbsvector(tmp9886)
}
__typedArg0 := tmp9886
return Call(__e, PrimFunc(symabsvector), __typedArg0)
})()

W59569132 := tmp9887
_ = W59569132

tmp9888 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(W59569132, MakeNumber(0), symshen_4print_1prolog_1vector)
}
__typedArg0 := W59569132
__typedArg1 := MakeNumber(0)
__typedArg2 := symshen_4print_1prolog_1vector
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

W59579133 := tmp9888
_ = W59579133

tmp9889 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(W59569132, MakeNumber(1), MakeNumber(2))
}
__typedArg0 := W59569132
__typedArg1 := MakeNumber(1)
__typedArg2 := MakeNumber(2)
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

W59589134 := tmp9889
_ = W59589134

__e.Return(W59589134)
return


}, 0)

tmp9890 := Call(__e, ns2_1set, symshen_4prolog_1vector, tmp9885)


_ = tmp9890

tmp9891 := MakeNative(func(__e *ControlFlow) {
V5959 := __e.Get(1)
_ = V5959
__e.Return(V5959)
return
}, 1)

tmp9892 := Call(__e, ns2_1set, symreceive, tmp9891)


_ = tmp9892

tmp9893 := MakeNative(func(__e *ControlFlow) {
V5960 := __e.Get(1)
_ = V5960
tmp9901 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5960)
}
__typedArg0 := V5960
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9901 {
tmp9894 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5960)
}
__typedArg0 := V5960
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9895 := Call(__e, PrimFunc(symshen_4rcons__form), tmp9894)


tmp9896 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5960)
}
__typedArg0 := V5960
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9897 := Call(__e, PrimFunc(symshen_4rcons__form), tmp9896)


tmp9898 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9897, Nil)
}
__typedArg0 := tmp9897
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9899 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9895, tmp9898)
}
__typedArg0 := tmp9895
__typedArg1 := tmp9898
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp9899)
}
__typedArg0 := symcons
__typedArg1 := tmp9899
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return(V5960)
return
}


}, 1)

tmp9902 := Call(__e, ns2_1set, symshen_4rcons__form, tmp9893)


_ = tmp9902

tmp9903 := MakeNative(func(__e *ControlFlow) {
V5961 := __e.Get(1)
_ = V5961
tmp9910 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5961)
}
__typedArg0 := V5961
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9910 {
tmp9904 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5961)
}
__typedArg0 := V5961
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9905 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5961)
}
__typedArg0 := V5961
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9906 := Call(__e, PrimFunc(symshen_4tuple_1up), tmp9905)


tmp9907 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9906, Nil)
}
__typedArg0 := tmp9906
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9908 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9904, tmp9907)
}
__typedArg0 := tmp9904
__typedArg1 := tmp9907
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8p, tmp9908)
}
__typedArg0 := sym_8p
__typedArg1 := tmp9908
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return(V5961)
return
}


}, 1)

tmp9911 := Call(__e, ns2_1set, symshen_4tuple_1up, tmp9903)


_ = tmp9911

tmp9912 := MakeNative(func(__e *ControlFlow) {
V5962 := __e.Get(1)
_ = V5962
tmp9913 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dmacros_d)
}
__typedArg0 := sym_dmacros_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp9914 := Call(__e, PrimFunc(symassoc), V5962, tmp9913)


tmp9915 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dmacros_d)
}
__typedArg0 := sym_dmacros_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp9916 := Call(__e, PrimFunc(symremove), tmp9914, tmp9915)


tmp9917 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(sym_dmacros_d, tmp9916)
}
__typedArg0 := sym_dmacros_d
__typedArg1 := tmp9916
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp9917

__e.Return(V5962)
return


}, 1)

__e.TailApply(ns2_1set, symundefmacro, tmp9912)
return




}, 0)

