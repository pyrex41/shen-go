package main

import . "github.com/pyrex41/shen-go/kl"

var MacrosMain = MakeNative(func(__e *ControlFlow) {
tmp8914 := MakeNative(func(__e *ControlFlow) {
V5824 := __e.Get(1)
_ = V5824
tmp8915 := MakeNative(func(__e *ControlFlow) {
W5825 := __e.Get(1)
_ = W5825
__e.TailApply(PrimFunc(symshen_4macroexpand_1h), V5824, W5825, W5825)
return
}, 1)

tmp8916 := MakeNative(func(__e *ControlFlow) {
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

tmp8917 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dmacros_d)
}
__typedArg0 := sym_dmacros_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp8918 := Call(__e, PrimFunc(symmap), tmp8916, tmp8917)


__e.TailApply(tmp8915, tmp8918)
return


}, 1)

tmp8919 := Call(__e, ns2_1set, symmacroexpand, tmp8914)


_ = tmp8919

tmp8920 := MakeNative(func(__e *ControlFlow) {
V5835 := __e.Get(1)
_ = V5835
V5836 := __e.Get(2)
_ = V5836
V5837 := __e.Get(3)
_ = V5837
tmp8930 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V5836)
}
__typedArg0 := Nil
__typedArg1 := V5836
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp8930 {
__e.Return(V5835)
return
} else {
tmp8928 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5836)
}
__typedArg0 := V5836
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp8928 {
tmp8921 := MakeNative(func(__e *ControlFlow) {
W5838 := __e.Get(1)
_ = W5838
tmp8924 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V5835, W5838)
}
__typedArg0 := V5835
__typedArg1 := W5838
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp8924 {
tmp8922 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5836)
}
__typedArg0 := V5836
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4macroexpand_1h), V5835, tmp8922, V5837)
return


} else {
__e.TailApply(PrimFunc(symshen_4macroexpand_1h), W5838, V5837, V5837)
return
}


}, 1)

tmp8925 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5836)
}
__typedArg0 := V5836
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8926 := Call(__e, PrimFunc(symshen_4walk), tmp8925, V5835)


__e.TailApply(tmp8921, tmp8926)
return


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

tmp8931 := Call(__e, ns2_1set, symshen_4macroexpand_1h, tmp8920)


_ = tmp8931

tmp8932 := MakeNative(func(__e *ControlFlow) {
V5839 := __e.Get(1)
_ = V5839
V5840 := __e.Get(2)
_ = V5840
tmp8936 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5840)
}
__typedArg0 := V5840
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp8936 {
tmp8933 := MakeNative(func(__e *ControlFlow) {
Z5841 := __e.Get(1)
_ = Z5841
__e.TailApply(PrimFunc(symshen_4walk), V5839, Z5841)
return
}, 1)

tmp8934 := Call(__e, PrimFunc(symmap), tmp8933, V5840)


__e.TailApply(V5839, tmp8934)
return


} else {
__e.TailApply(V5839, V5840)
return
}


}, 2)

tmp8937 := Call(__e, ns2_1set, symshen_4walk, tmp8932)


_ = tmp8937

tmp8938 := MakeNative(func(__e *ControlFlow) {
V5842 := __e.Get(1)
_ = V5842
tmp8939 := MakeNative(func(__e *ControlFlow) {
GoTo5843 := __e.Get(1)
_ = GoTo5843
tmp9242 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5842)
}
__typedArg0 := V5842
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9242 {
tmp8940 := MakeNative(func(__e *ControlFlow) {
Select5848 := __e.Get(1)
_ = Select5848
tmp8941 := MakeNative(func(__e *ControlFlow) {
Select5849 := __e.Get(1)
_ = Select5849
tmp9238 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefmacro, Select5848)
}
__typedArg0 := symdefmacro
__typedArg1 := Select5848
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9235 Obj

if True == tmp9238 {
tmp9237 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(Select5849)
}
__typedArg0 := Select5849
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9236 Obj

if True == tmp9237 {
ifres9236 = True


} else {
ifres9236 = False


}

ifres9235 = ifres9236


} else {
ifres9235 = False


}

if True == ifres9235 {
tmp8942 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(Select5849)
}
__typedArg0 := Select5849
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8943 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select5849)
}
__typedArg0 := Select5849
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4process_1def), tmp8942, tmp8943)
return


} else {
tmp9233 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefcc, Select5848)
}
__typedArg0 := symdefcc
__typedArg1 := Select5848
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9233 {
__e.TailApply(PrimFunc(symshen_4yacc_1_6shen), Select5849)
return
} else {
tmp9231 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symu_b, Select5848)
}
__typedArg0 := symu_b
__typedArg1 := Select5848
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9224 Obj

if True == tmp9231 {
tmp9230 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(Select5849)
}
__typedArg0 := Select5849
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9226 Obj

if True == tmp9230 {
tmp9228 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select5849)
}
__typedArg0 := Select5849
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9229 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp9228)
}
__typedArg0 := Nil
__typedArg1 := tmp9228
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9227 Obj

if True == tmp9229 {
ifres9227 = True


} else {
ifres9227 = False


}

ifres9226 = ifres9227


} else {
ifres9226 = False


}

var ifres9225 Obj

if True == ifres9226 {
ifres9225 = True


} else {
ifres9225 = False


}

ifres9224 = ifres9225


} else {
ifres9224 = False


}

if True == ifres9224 {
tmp8944 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(Select5849)
}
__typedArg0 := Select5849
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8945 := Call(__e, PrimFunc(symshen_4make_1uppercase), tmp8944)


tmp8946 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8945, Nil)
}
__typedArg0 := tmp8945
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symprotect, tmp8946)
}
__typedArg0 := symprotect
__typedArg1 := tmp8946
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9222 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symerror, Select5848)
}
__typedArg0 := symerror
__typedArg1 := Select5848
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9219 Obj

if True == tmp9222 {
tmp9221 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(Select5849)
}
__typedArg0 := Select5849
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9220 Obj

if True == tmp9221 {
ifres9220 = True


} else {
ifres9220 = False


}

ifres9219 = ifres9220


} else {
ifres9219 = False


}

if True == ifres9219 {
tmp8947 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(Select5849)
}
__typedArg0 := Select5849
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8948 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select5849)
}
__typedArg0 := Select5849
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8949 := Call(__e, PrimFunc(symshen_4mkstr), tmp8947, tmp8948)


tmp8950 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8949, Nil)
}
__typedArg0 := tmp8949
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsimple_1error, tmp8950)
}
__typedArg0 := symsimple_1error
__typedArg1 := tmp8950
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9217 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symoutput, Select5848)
}
__typedArg0 := symoutput
__typedArg1 := Select5848
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9214 Obj

if True == tmp9217 {
tmp9216 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(Select5849)
}
__typedArg0 := Select5849
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9215 Obj

if True == tmp9216 {
ifres9215 = True


} else {
ifres9215 = False


}

ifres9214 = ifres9215


} else {
ifres9214 = False


}

if True == ifres9214 {
tmp8951 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(Select5849)
}
__typedArg0 := Select5849
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8952 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select5849)
}
__typedArg0 := Select5849
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8953 := Call(__e, PrimFunc(symshen_4mkstr), tmp8951, tmp8952)


tmp8954 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstoutput, Nil)
}
__typedArg0 := symstoutput
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8955 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8954, Nil)
}
__typedArg0 := tmp8954
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8956 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8953, tmp8955)
}
__typedArg0 := tmp8953
__typedArg1 := tmp8955
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympr, tmp8956)
}
__typedArg0 := sympr
__typedArg1 := tmp8956
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9212 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sympr, Select5848)
}
__typedArg0 := sympr
__typedArg1 := Select5848
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9205 Obj

if True == tmp9212 {
tmp9211 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(Select5849)
}
__typedArg0 := Select5849
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9207 Obj

if True == tmp9211 {
tmp9209 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select5849)
}
__typedArg0 := Select5849
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9210 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp9209)
}
__typedArg0 := Nil
__typedArg1 := tmp9209
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9208 Obj

if True == tmp9210 {
ifres9208 = True


} else {
ifres9208 = False


}

ifres9207 = ifres9208


} else {
ifres9207 = False


}

var ifres9206 Obj

if True == ifres9207 {
ifres9206 = True


} else {
ifres9206 = False


}

ifres9205 = ifres9206


} else {
ifres9205 = False


}

if True == ifres9205 {
tmp8957 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(Select5849)
}
__typedArg0 := Select5849
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8958 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstoutput, Nil)
}
__typedArg0 := symstoutput
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8959 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8958, Nil)
}
__typedArg0 := tmp8958
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8960 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8957, tmp8959)
}
__typedArg0 := tmp8957
__typedArg1 := tmp8959
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympr, tmp8960)
}
__typedArg0 := sympr
__typedArg1 := tmp8960
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9203 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symmake_1string, Select5848)
}
__typedArg0 := symmake_1string
__typedArg1 := Select5848
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9200 Obj

if True == tmp9203 {
tmp9202 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(Select5849)
}
__typedArg0 := Select5849
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9201 Obj

if True == tmp9202 {
ifres9201 = True


} else {
ifres9201 = False


}

ifres9200 = ifres9201


} else {
ifres9200 = False


}

if True == ifres9200 {
tmp8961 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(Select5849)
}
__typedArg0 := Select5849
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8962 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select5849)
}
__typedArg0 := Select5849
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4mkstr), tmp8961, tmp8962)
return


} else {
tmp9198 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlineread, Select5848)
}
__typedArg0 := symlineread
__typedArg1 := Select5848
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9195 Obj

if True == tmp9198 {
tmp9197 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, Select5849)
}
__typedArg0 := Nil
__typedArg1 := Select5849
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9196 Obj

if True == tmp9197 {
ifres9196 = True


} else {
ifres9196 = False


}

ifres9195 = ifres9196


} else {
ifres9195 = False


}

if True == ifres9195 {
tmp8963 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstinput, Nil)
}
__typedArg0 := symstinput
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

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
return PrimCons(symlineread, tmp8964)
}
__typedArg0 := symlineread
__typedArg1 := tmp8964
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9193 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(syminput, Select5848)
}
__typedArg0 := syminput
__typedArg1 := Select5848
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9190 Obj

if True == tmp9193 {
tmp9192 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, Select5849)
}
__typedArg0 := Nil
__typedArg1 := Select5849
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9191 Obj

if True == tmp9192 {
ifres9191 = True


} else {
ifres9191 = False


}

ifres9190 = ifres9191


} else {
ifres9190 = False


}

if True == ifres9190 {
tmp8965 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstinput, Nil)
}
__typedArg0 := symstinput
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8966 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8965, Nil)
}
__typedArg0 := tmp8965
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminput, tmp8966)
}
__typedArg0 := syminput
__typedArg1 := tmp8966
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9188 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symread, Select5848)
}
__typedArg0 := symread
__typedArg1 := Select5848
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9185 Obj

if True == tmp9188 {
tmp9187 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, Select5849)
}
__typedArg0 := Nil
__typedArg1 := Select5849
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9186 Obj

if True == tmp9187 {
ifres9186 = True


} else {
ifres9186 = False


}

ifres9185 = ifres9186


} else {
ifres9185 = False


}

if True == ifres9185 {
tmp8967 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstinput, Nil)
}
__typedArg0 := symstinput
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8968 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8967, Nil)
}
__typedArg0 := tmp8967
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread, tmp8968)
}
__typedArg0 := symread
__typedArg1 := tmp8968
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9183 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(syminput_7, Select5848)
}
__typedArg0 := syminput_7
__typedArg1 := Select5848
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9180 Obj

if True == tmp9183 {
tmp9182 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(Select5849)
}
__typedArg0 := Select5849
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9181 Obj

if True == tmp9182 {
ifres9181 = True


} else {
ifres9181 = False


}

ifres9180 = ifres9181


} else {
ifres9180 = False


}

if True == ifres9180 {
__e.TailApply(PrimFunc(symshen_4process_1input_7), V5842)
return
} else {
tmp9178 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symread_1byte, Select5848)
}
__typedArg0 := symread_1byte
__typedArg1 := Select5848
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9175 Obj

if True == tmp9178 {
tmp9177 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, Select5849)
}
__typedArg0 := Nil
__typedArg1 := Select5849
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9176 Obj

if True == tmp9177 {
ifres9176 = True


} else {
ifres9176 = False


}

ifres9175 = ifres9176


} else {
ifres9175 = False


}

if True == ifres9175 {
__e.TailApply(PrimFunc(symshen_4process_1read_1byte))
return
} else {
tmp9173 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symprolog_2, Select5848)
}
__typedArg0 := symprolog_2
__typedArg1 := Select5848
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9173 {
__e.TailApply(PrimFunc(symshen_4call_1prolog), Select5849)
return
} else {
tmp9171 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefprolog, Select5848)
}
__typedArg0 := symdefprolog
__typedArg1 := Select5848
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9168 Obj

if True == tmp9171 {
tmp9170 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(Select5849)
}
__typedArg0 := Select5849
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9169 Obj

if True == tmp9170 {
ifres9169 = True


} else {
ifres9169 = False


}

ifres9168 = ifres9169


} else {
ifres9168 = False


}

if True == ifres9168 {
tmp8969 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(Select5849)
}
__typedArg0 := Select5849
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8970 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select5849)
}
__typedArg0 := Select5849
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4compile_1prolog), tmp8969, tmp8970)
return


} else {
tmp9166 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdatatype, Select5848)
}
__typedArg0 := symdatatype
__typedArg1 := Select5848
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9163 Obj

if True == tmp9166 {
tmp9165 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(Select5849)
}
__typedArg0 := Select5849
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9164 Obj

if True == tmp9165 {
ifres9164 = True


} else {
ifres9164 = False


}

ifres9163 = ifres9164


} else {
ifres9163 = False


}

if True == ifres9163 {
tmp8971 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(Select5849)
}
__typedArg0 := Select5849
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8972 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select5849)
}
__typedArg0 := Select5849
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4process_1datatype), tmp8971, tmp8972)
return


} else {
tmp9161 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_8s, Select5848)
}
__typedArg0 := sym_8s
__typedArg1 := Select5848
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9161 {
__e.TailApply(PrimFunc(symshen_4process_1_8s), V5842)
return
} else {
tmp9159 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symsynonyms, Select5848)
}
__typedArg0 := symsynonyms
__typedArg1 := Select5848
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9159 {
__e.TailApply(PrimFunc(symshen_4process_1synonyms), Select5849)
return
} else {
tmp9157 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symnl, Select5848)
}
__typedArg0 := symnl
__typedArg1 := Select5848
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9154 Obj

if True == tmp9157 {
tmp9156 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, Select5849)
}
__typedArg0 := Nil
__typedArg1 := Select5849
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9155 Obj

if True == tmp9156 {
ifres9155 = True


} else {
ifres9155 = False


}

ifres9154 = ifres9155


} else {
ifres9154 = False


}

if True == ifres9154 {
tmp8973 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), Nil)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnl, tmp8973)
}
__typedArg0 := symnl
__typedArg1 := tmp8973
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9152 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlet, Select5848)
}
__typedArg0 := symlet
__typedArg1 := Select5848
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9152 {
__e.TailApply(PrimFunc(symshen_4process_1let), V5842)
return
} else {
tmp9150 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_c_4, Select5848)
}
__typedArg0 := sym_c_4
__typedArg1 := Select5848
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9150 {
__e.TailApply(PrimFunc(symshen_4process_1lambda), V5842)
return
} else {
tmp9148 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcases, Select5848)
}
__typedArg0 := symcases
__typedArg1 := Select5848
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9148 {
__e.TailApply(PrimFunc(symshen_4process_1cases), V5842)
return
} else {
tmp9146 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symtime, Select5848)
}
__typedArg0 := symtime
__typedArg1 := Select5848
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9139 Obj

if True == tmp9146 {
tmp9145 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(Select5849)
}
__typedArg0 := Select5849
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9141 Obj

if True == tmp9145 {
tmp9143 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select5849)
}
__typedArg0 := Select5849
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9144 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp9143)
}
__typedArg0 := Nil
__typedArg1 := tmp9143
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9142 Obj

if True == tmp9144 {
ifres9142 = True


} else {
ifres9142 = False


}

ifres9141 = ifres9142


} else {
ifres9141 = False


}

var ifres9140 Obj

if True == ifres9141 {
ifres9140 = True


} else {
ifres9140 = False


}

ifres9139 = ifres9140


} else {
ifres9139 = False


}

if True == ifres9139 {
tmp8974 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(Select5849)
}
__typedArg0 := Select5849
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4process_1time), tmp8974)
return


} else {
tmp9137 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symput, Select5848)
}
__typedArg0 := symput
__typedArg1 := Select5848
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9119 Obj

if True == tmp9137 {
tmp9136 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(Select5849)
}
__typedArg0 := Select5849
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9121 Obj

if True == tmp9136 {
tmp9134 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select5849)
}
__typedArg0 := Select5849
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9135 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9134)
}
__typedArg0 := tmp9134
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9123 Obj

if True == tmp9135 {
tmp9131 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select5849)
}
__typedArg0 := Select5849
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9132 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9131)
}
__typedArg0 := tmp9131
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9133 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9132)
}
__typedArg0 := tmp9132
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9125 Obj

if True == tmp9133 {
tmp9127 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select5849)
}
__typedArg0 := Select5849
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9128 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9127)
}
__typedArg0 := tmp9127
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp9129)
}
__typedArg0 := Nil
__typedArg1 := tmp9129
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9126 Obj

if True == tmp9130 {
ifres9126 = True


} else {
ifres9126 = False


}

ifres9125 = ifres9126


} else {
ifres9125 = False


}

var ifres9124 Obj

if True == ifres9125 {
ifres9124 = True


} else {
ifres9124 = False


}

ifres9123 = ifres9124


} else {
ifres9123 = False


}

var ifres9122 Obj

if True == ifres9123 {
ifres9122 = True


} else {
ifres9122 = False


}

ifres9121 = ifres9122


} else {
ifres9121 = False


}

var ifres9120 Obj

if True == ifres9121 {
ifres9120 = True


} else {
ifres9120 = False


}

ifres9119 = ifres9120


} else {
ifres9119 = False


}

if True == ifres9119 {
tmp8975 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(Select5849)
}
__typedArg0 := Select5849
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8976 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select5849)
}
__typedArg0 := Select5849
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8977 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp8976)
}
__typedArg0 := tmp8976
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8978 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select5849)
}
__typedArg0 := Select5849
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8979 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp8978)
}
__typedArg0 := tmp8978
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8980 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp8979)
}
__typedArg0 := tmp8979
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8981 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dproperty_1vector_d, Nil)
}
__typedArg0 := sym_dproperty_1vector_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8982 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvalue, tmp8981)
}
__typedArg0 := symvalue
__typedArg1 := tmp8981
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8983 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8982, Nil)
}
__typedArg0 := tmp8982
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8984 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8980, tmp8983)
}
__typedArg0 := tmp8980
__typedArg1 := tmp8983
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8985 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8977, tmp8984)
}
__typedArg0 := tmp8977
__typedArg1 := tmp8984
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8986 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8975, tmp8985)
}
__typedArg0 := tmp8975
__typedArg1 := tmp8985
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symput, tmp8986)
}
__typedArg0 := symput
__typedArg1 := tmp8986
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9117 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symget, Select5848)
}
__typedArg0 := symget
__typedArg1 := Select5848
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9105 Obj

if True == tmp9117 {
tmp9116 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(Select5849)
}
__typedArg0 := Select5849
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9107 Obj

if True == tmp9116 {
tmp9114 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select5849)
}
__typedArg0 := Select5849
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9115 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9114)
}
__typedArg0 := tmp9114
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9109 Obj

if True == tmp9115 {
tmp9111 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select5849)
}
__typedArg0 := Select5849
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9112 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9111)
}
__typedArg0 := tmp9111
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9113 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp9112)
}
__typedArg0 := Nil
__typedArg1 := tmp9112
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9110 Obj

if True == tmp9113 {
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
tmp8987 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(Select5849)
}
__typedArg0 := Select5849
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8988 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select5849)
}
__typedArg0 := Select5849
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8989 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp8988)
}
__typedArg0 := tmp8988
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8990 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dproperty_1vector_d, Nil)
}
__typedArg0 := sym_dproperty_1vector_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8991 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvalue, tmp8990)
}
__typedArg0 := symvalue
__typedArg1 := tmp8990
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8992 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8991, Nil)
}
__typedArg0 := tmp8991
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8993 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8989, tmp8992)
}
__typedArg0 := tmp8989
__typedArg1 := tmp8992
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8994 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8987, tmp8993)
}
__typedArg0 := tmp8987
__typedArg1 := tmp8993
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symget, tmp8994)
}
__typedArg0 := symget
__typedArg1 := tmp8994
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9103 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symunput, Select5848)
}
__typedArg0 := symunput
__typedArg1 := Select5848
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9091 Obj

if True == tmp9103 {
tmp9102 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(Select5849)
}
__typedArg0 := Select5849
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9093 Obj

if True == tmp9102 {
tmp9100 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select5849)
}
__typedArg0 := Select5849
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9101 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9100)
}
__typedArg0 := tmp9100
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9095 Obj

if True == tmp9101 {
tmp9097 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select5849)
}
__typedArg0 := Select5849
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9098 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9097)
}
__typedArg0 := tmp9097
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9099 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp9098)
}
__typedArg0 := Nil
__typedArg1 := tmp9098
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9096 Obj

if True == tmp9099 {
ifres9096 = True


} else {
ifres9096 = False


}

ifres9095 = ifres9096


} else {
ifres9095 = False


}

var ifres9094 Obj

if True == ifres9095 {
ifres9094 = True


} else {
ifres9094 = False


}

ifres9093 = ifres9094


} else {
ifres9093 = False


}

var ifres9092 Obj

if True == ifres9093 {
ifres9092 = True


} else {
ifres9092 = False


}

ifres9091 = ifres9092


} else {
ifres9091 = False


}

if True == ifres9091 {
tmp8995 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(Select5849)
}
__typedArg0 := Select5849
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8996 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select5849)
}
__typedArg0 := Select5849
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8997 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp8996)
}
__typedArg0 := tmp8996
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8998 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_dproperty_1vector_d, Nil)
}
__typedArg0 := sym_dproperty_1vector_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8999 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvalue, tmp8998)
}
__typedArg0 := symvalue
__typedArg1 := tmp8998
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9000 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8999, Nil)
}
__typedArg0 := tmp8999
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9001 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8997, tmp9000)
}
__typedArg0 := tmp8997
__typedArg1 := tmp9000
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9002 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8995, tmp9001)
}
__typedArg0 := tmp8995
__typedArg1 := tmp9001
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunput, tmp9002)
}
__typedArg0 := symunput
__typedArg1 := tmp9002
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9089 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4_8c, Select5848)
}
__typedArg0 := symshen_4_8c
__typedArg1 := Select5848
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9082 Obj

if True == tmp9089 {
tmp9088 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(Select5849)
}
__typedArg0 := Select5849
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9084 Obj

if True == tmp9088 {
tmp9086 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select5849)
}
__typedArg0 := Select5849
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9087 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp9086)
}
__typedArg0 := Nil
__typedArg1 := tmp9086
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9085 Obj

if True == tmp9087 {
ifres9085 = True


} else {
ifres9085 = False


}

ifres9084 = ifres9085


} else {
ifres9084 = False


}

var ifres9083 Obj

if True == ifres9084 {
ifres9083 = True


} else {
ifres9083 = False


}

ifres9082 = ifres9083


} else {
ifres9082 = False


}

if True == ifres9082 {
tmp9003 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(Select5849)
}
__typedArg0 := Select5849
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4rcons__form), tmp9003)
return


} else {
tmp9004 := MakeNative(func(__e *ControlFlow) {
GoTo5844 := __e.Get(1)
_ = GoTo5844
tmp9051 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4_8ch, Select5848)
}
__typedArg0 := symshen_4_8ch
__typedArg1 := Select5848
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9051 {
tmp9049 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(Select5849)
}
__typedArg0 := Select5849
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9049 {
tmp9005 := MakeNative(func(__e *ControlFlow) {
Select5846 := __e.Get(1)
_ = Select5846
tmp9006 := MakeNative(func(__e *ControlFlow) {
Select5847 := __e.Get(1)
_ = Select5847
tmp9045 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(Select5846)
}
__typedArg0 := Select5846
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9021 Obj

if True == tmp9045 {
tmp9043 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select5846)
}
__typedArg0 := Select5846
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9044 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9043)
}
__typedArg0 := tmp9043
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9023 Obj

if True == tmp9044 {
tmp9040 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select5846)
}
__typedArg0 := Select5846
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9041 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9040)
}
__typedArg0 := tmp9040
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9042 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9041)
}
__typedArg0 := tmp9041
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9025 Obj

if True == tmp9042 {
tmp9036 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select5846)
}
__typedArg0 := Select5846
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9037 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9036)
}
__typedArg0 := tmp9036
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp9038)
}
__typedArg0 := Nil
__typedArg1 := tmp9038
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9027 Obj

if True == tmp9039 {
tmp9035 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, Select5847)
}
__typedArg0 := Nil
__typedArg1 := Select5847
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9029 Obj

if True == tmp9035 {
tmp9031 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select5846)
}
__typedArg0 := Select5846
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9032 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9031)
}
__typedArg0 := tmp9031
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9033 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp9034 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp9032, tmp9033)
}
__typedArg0 := tmp9032
__typedArg1 := tmp9033
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9030 Obj

if True == tmp9034 {
ifres9030 = True


} else {
ifres9030 = False


}

ifres9029 = ifres9030


} else {
ifres9029 = False


}

var ifres9028 Obj

if True == ifres9029 {
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

var ifres9024 Obj

if True == ifres9025 {
ifres9024 = True


} else {
ifres9024 = False


}

ifres9023 = ifres9024


} else {
ifres9023 = False


}

var ifres9022 Obj

if True == ifres9023 {
ifres9022 = True


} else {
ifres9022 = False


}

ifres9021 = ifres9022


} else {
ifres9021 = False


}

if True == ifres9021 {
tmp9007 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(Select5846)
}
__typedArg0 := Select5846
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9008 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select5846)
}
__typedArg0 := Select5846
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
return PrimTail(Select5846)
}
__typedArg0 := Select5846
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_7, tmp9011)
}
__typedArg0 := sym_7
__typedArg1 := tmp9011
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9013 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9012, Nil)
}
__typedArg0 := tmp9012
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9014 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9009, tmp9013)
}
__typedArg0 := tmp9009
__typedArg1 := tmp9013
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9015 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9007, tmp9014)
}
__typedArg0 := tmp9007
__typedArg1 := tmp9014
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9016 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9015, Nil)
}
__typedArg0 := tmp9015
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9017 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1, tmp9016)
}
__typedArg0 := sym_1
__typedArg1 := tmp9016
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4cons_1form_1respect_1modes), tmp9017)
return


} else {
tmp9019 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, Select5847)
}
__typedArg0 := Nil
__typedArg1 := Select5847
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9019 {
__e.TailApply(PrimFunc(symshen_4cons_1form_1respect_1modes), Select5846)
return
} else {
__e.TailApply(PrimFunc(symthaw), GoTo5844)
return
}


}


}, 1)

tmp9046 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select5849)
}
__typedArg0 := Select5849
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(tmp9006, tmp9046)
return


}, 1)

tmp9047 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(Select5849)
}
__typedArg0 := Select5849
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp9005, tmp9047)
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5844)
return
}


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5844)
return
}


}, 1)

tmp9052 := MakeNative(func(__e *ControlFlow) {
tmp9080 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(Select5849)
}
__typedArg0 := Select5849
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9060 Obj

if True == tmp9080 {
tmp9078 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select5849)
}
__typedArg0 := Select5849
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9079 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9078)
}
__typedArg0 := tmp9078
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9062 Obj

if True == tmp9079 {
tmp9075 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select5849)
}
__typedArg0 := Select5849
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9076 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9075)
}
__typedArg0 := tmp9075
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9077 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9076)
}
__typedArg0 := tmp9076
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9064 Obj

if True == tmp9077 {
tmp9066 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdo, Nil)
}
__typedArg0 := symdo
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9067 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_d, tmp9066)
}
__typedArg0 := sym_d
__typedArg1 := tmp9066
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9068 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_7, tmp9067)
}
__typedArg0 := sym_7
__typedArg1 := tmp9067
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9069 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symor, tmp9068)
}
__typedArg0 := symor
__typedArg1 := tmp9068
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9070 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symand, tmp9069)
}
__typedArg0 := symand
__typedArg1 := tmp9069
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9071 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symappend, tmp9070)
}
__typedArg0 := symappend
__typedArg1 := tmp9070
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9072 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8v, tmp9071)
}
__typedArg0 := sym_8v
__typedArg1 := tmp9071
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9073 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8p, tmp9072)
}
__typedArg0 := sym_8p
__typedArg1 := tmp9072
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9074 := Call(__e, PrimFunc(symelement_2), Select5848, tmp9073)


var ifres9065 Obj

if True == tmp9074 {
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

var ifres9061 Obj

if True == ifres9062 {
ifres9061 = True


} else {
ifres9061 = False


}

ifres9060 = ifres9061


} else {
ifres9060 = False


}

if True == ifres9060 {
tmp9053 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(Select5849)
}
__typedArg0 := Select5849
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9054 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select5849)
}
__typedArg0 := Select5849
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9055 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Select5848, tmp9054)
}
__typedArg0 := Select5848
__typedArg1 := tmp9054
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9056 := Call(__e, PrimFunc(symshen_4process_1assoc), tmp9055)


tmp9057 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9056, Nil)
}
__typedArg0 := tmp9056
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9058 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9053, tmp9057)
}
__typedArg0 := tmp9053
__typedArg1 := tmp9057
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Select5848, tmp9058)
}
__typedArg0 := Select5848
__typedArg1 := tmp9058
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5843)
return
}


}, 0)

__e.TailApply(tmp9004, tmp9052)
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


}, 1)

tmp9239 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5842)
}
__typedArg0 := V5842
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(tmp8941, tmp9239)
return


}, 1)

tmp9240 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5842)
}
__typedArg0 := V5842
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp8940, tmp9240)
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5843)
return
}


}, 1)

tmp9243 := MakeNative(func(__e *ControlFlow) {
__e.Return(V5842)
return
}, 0)

__e.TailApply(tmp8939, tmp9243)
return


}, 1)

tmp9244 := Call(__e, ns2_1set, symshen_4macros, tmp8938)


_ = tmp9244

tmp9245 := MakeNative(func(__e *ControlFlow) {
V5850 := __e.Get(1)
_ = V5850
tmp9246 := MakeNative(func(__e *ControlFlow) {
GoTo5851 := __e.Get(1)
_ = GoTo5851
tmp9273 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5850)
}
__typedArg0 := V5850
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9273 {
tmp9247 := MakeNative(func(__e *ControlFlow) {
Select5856 := __e.Get(1)
_ = Select5856
tmp9269 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5850)
}
__typedArg0 := V5850
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9270 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(syminput_7, tmp9269)
}
__typedArg0 := syminput_7
__typedArg1 := tmp9269
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9270 {
tmp9267 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(Select5856)
}
__typedArg0 := Select5856
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9267 {
tmp9248 := MakeNative(func(__e *ControlFlow) {
Select5854 := __e.Get(1)
_ = Select5854
tmp9249 := MakeNative(func(__e *ControlFlow) {
Select5855 := __e.Get(1)
_ = Select5855
tmp9263 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, Select5855)
}
__typedArg0 := Nil
__typedArg1 := Select5855
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9263 {
tmp9250 := Call(__e, PrimFunc(symshen_4rcons__form), Select5854)


tmp9251 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstinput, Nil)
}
__typedArg0 := symstinput
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9252 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9251, Nil)
}
__typedArg0 := tmp9251
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9253 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9250, tmp9252)
}
__typedArg0 := tmp9250
__typedArg1 := tmp9252
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4input_1h_7, tmp9253)
}
__typedArg0 := symshen_4input_1h_7
__typedArg1 := tmp9253
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9261 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(Select5855)
}
__typedArg0 := Select5855
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9257 Obj

if True == tmp9261 {
tmp9259 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select5855)
}
__typedArg0 := Select5855
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9260 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp9259)
}
__typedArg0 := Nil
__typedArg1 := tmp9259
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9258 Obj

if True == tmp9260 {
ifres9258 = True


} else {
ifres9258 = False


}

ifres9257 = ifres9258


} else {
ifres9257 = False


}

if True == ifres9257 {
tmp9254 := Call(__e, PrimFunc(symshen_4rcons__form), Select5854)


tmp9255 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9254, Select5855)
}
__typedArg0 := tmp9254
__typedArg1 := Select5855
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4input_1h_7, tmp9255)
}
__typedArg0 := symshen_4input_1h_7
__typedArg1 := tmp9255
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5851)
return
}


}


}, 1)

tmp9264 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select5856)
}
__typedArg0 := Select5856
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(tmp9249, tmp9264)
return


}, 1)

tmp9265 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(Select5856)
}
__typedArg0 := Select5856
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp9248, tmp9265)
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5851)
return
}


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5851)
return
}


}, 1)

tmp9271 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5850)
}
__typedArg0 := V5850
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(tmp9247, tmp9271)
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5851)
return
}


}, 1)

tmp9274 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("partial function shen.process-input+"))
}
__typedArg0 := MakeString("partial function shen.process-input+")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}, 0)

__e.TailApply(tmp9246, tmp9274)
return


}, 1)

tmp9275 := Call(__e, ns2_1set, symshen_4process_1input_7, tmp9245)


_ = tmp9275

tmp9276 := MakeNative(func(__e *ControlFlow) {
V5857 := __e.Get(1)
_ = V5857
tmp9277 := MakeNative(func(__e *ControlFlow) {
GoTo5858 := __e.Get(1)
_ = GoTo5858
tmp9311 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5857)
}
__typedArg0 := V5857
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9311 {
tmp9278 := MakeNative(func(__e *ControlFlow) {
Select5859 := __e.Get(1)
_ = Select5859
tmp9279 := MakeNative(func(__e *ControlFlow) {
Select5860 := __e.Get(1)
_ = Select5860
tmp9307 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_7, Select5859)
}
__typedArg0 := sym_7
__typedArg1 := Select5859
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9300 Obj

if True == tmp9307 {
tmp9306 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(Select5860)
}
__typedArg0 := Select5860
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9302 Obj

if True == tmp9306 {
tmp9304 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select5860)
}
__typedArg0 := Select5860
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9305 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp9304)
}
__typedArg0 := Nil
__typedArg1 := tmp9304
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9303 Obj

if True == tmp9305 {
ifres9303 = True


} else {
ifres9303 = False


}

ifres9302 = ifres9303


} else {
ifres9302 = False


}

var ifres9301 Obj

if True == ifres9302 {
ifres9301 = True


} else {
ifres9301 = False


}

ifres9300 = ifres9301


} else {
ifres9300 = False


}

if True == ifres9300 {
tmp9280 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(Select5860)
}
__typedArg0 := Select5860
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9281 := Call(__e, PrimFunc(symshen_4cons_1form_1respect_1modes), tmp9280)


tmp9282 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9281, Nil)
}
__typedArg0 := tmp9281
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_7, tmp9282)
}
__typedArg0 := sym_7
__typedArg1 := tmp9282
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9298 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_1, Select5859)
}
__typedArg0 := sym_1
__typedArg1 := Select5859
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9291 Obj

if True == tmp9298 {
tmp9297 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(Select5860)
}
__typedArg0 := Select5860
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9293 Obj

if True == tmp9297 {
tmp9295 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select5860)
}
__typedArg0 := Select5860
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9296 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp9295)
}
__typedArg0 := Nil
__typedArg1 := tmp9295
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9294 Obj

if True == tmp9296 {
ifres9294 = True


} else {
ifres9294 = False


}

ifres9293 = ifres9294


} else {
ifres9293 = False


}

var ifres9292 Obj

if True == ifres9293 {
ifres9292 = True


} else {
ifres9292 = False


}

ifres9291 = ifres9292


} else {
ifres9291 = False


}

if True == ifres9291 {
tmp9283 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(Select5860)
}
__typedArg0 := Select5860
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9284 := Call(__e, PrimFunc(symshen_4cons_1form_1respect_1modes), tmp9283)


tmp9285 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9284, Nil)
}
__typedArg0 := tmp9284
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1, tmp9285)
}
__typedArg0 := sym_1
__typedArg1 := tmp9285
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9286 := Call(__e, PrimFunc(symshen_4cons_1form_1respect_1modes), Select5859)


tmp9287 := Call(__e, PrimFunc(symshen_4cons_1form_1respect_1modes), Select5860)


tmp9288 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9287, Nil)
}
__typedArg0 := tmp9287
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9289 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9286, tmp9288)
}
__typedArg0 := tmp9286
__typedArg1 := tmp9288
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp9289)
}
__typedArg0 := symcons
__typedArg1 := tmp9289
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


}


}, 1)

tmp9308 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5857)
}
__typedArg0 := V5857
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(tmp9279, tmp9308)
return


}, 1)

tmp9309 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5857)
}
__typedArg0 := V5857
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp9278, tmp9309)
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5858)
return
}


}, 1)

tmp9312 := MakeNative(func(__e *ControlFlow) {
__e.Return(V5857)
return
}, 0)

__e.TailApply(tmp9277, tmp9312)
return


}, 1)

tmp9313 := Call(__e, ns2_1set, symshen_4cons_1form_1respect_1modes, tmp9276)


_ = tmp9313

tmp9314 := MakeNative(func(__e *ControlFlow) {
V5861 := __e.Get(1)
_ = V5861
V5862 := __e.Get(2)
_ = V5862
tmp9315 := MakeNative(func(__e *ControlFlow) {
W5863 := __e.Get(1)
_ = W5863
tmp9316 := MakeNative(func(__e *ControlFlow) {
W5864 := __e.Get(1)
_ = W5864
tmp9317 := MakeNative(func(__e *ControlFlow) {
W5865 := __e.Get(1)
_ = W5865
__e.Return(V5861)
return
}, 1)

tmp9318 := Call(__e, PrimFunc(symfn), V5861)


tmp9319 := Call(__e, PrimFunc(symshen_4record_1macro), V5861, tmp9318)


__e.TailApply(tmp9317, tmp9319)
return


}, 1)

tmp9320 := Call(__e, PrimFunc(symappend), V5862, W5863)


tmp9321 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5861, tmp9320)
}
__typedArg0 := V5861
__typedArg1 := tmp9320
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9322 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefine, tmp9321)
}
__typedArg0 := symdefine
__typedArg1 := tmp9321
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9323 := Call(__e, PrimFunc(symeval), tmp9322)


__e.TailApply(tmp9316, tmp9323)
return


}, 1)

tmp9324 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symX, Nil)
}
__typedArg0 := symX
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9325 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_6, tmp9324)
}
__typedArg0 := sym_1_6
__typedArg1 := tmp9324
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9326 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symX, tmp9325)
}
__typedArg0 := symX
__typedArg1 := tmp9325
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(tmp9315, tmp9326)
return


}, 2)

tmp9327 := Call(__e, ns2_1set, symshen_4process_1def, tmp9314)


_ = tmp9327

tmp9328 := MakeNative(func(__e *ControlFlow) {
V5866 := __e.Get(1)
_ = V5866
tmp9368 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5866)
}
__typedArg0 := V5866
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9342 Obj

if True == tmp9368 {
tmp9366 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5866)
}
__typedArg0 := V5866
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9367 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlet, tmp9366)
}
__typedArg0 := symlet
__typedArg1 := tmp9366
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9344 Obj

if True == tmp9367 {
tmp9364 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5866)
}
__typedArg0 := V5866
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9365 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9364)
}
__typedArg0 := tmp9364
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9346 Obj

if True == tmp9365 {
tmp9361 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5866)
}
__typedArg0 := V5866
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9362 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9361)
}
__typedArg0 := tmp9361
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9363 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9362)
}
__typedArg0 := tmp9362
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9348 Obj

if True == tmp9363 {
tmp9357 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5866)
}
__typedArg0 := V5866
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9358 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9357)
}
__typedArg0 := tmp9357
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9359 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9358)
}
__typedArg0 := tmp9358
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9360 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9359)
}
__typedArg0 := tmp9359
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9350 Obj

if True == tmp9360 {
tmp9352 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5866)
}
__typedArg0 := V5866
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9353 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9352)
}
__typedArg0 := tmp9352
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9354 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9353)
}
__typedArg0 := tmp9353
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9355 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9354)
}
__typedArg0 := tmp9354
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9356 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9355)
}
__typedArg0 := tmp9355
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9351 Obj

if True == tmp9356 {
ifres9351 = True


} else {
ifres9351 = False


}

ifres9350 = ifres9351


} else {
ifres9350 = False


}

var ifres9349 Obj

if True == ifres9350 {
ifres9349 = True


} else {
ifres9349 = False


}

ifres9348 = ifres9349


} else {
ifres9348 = False


}

var ifres9347 Obj

if True == ifres9348 {
ifres9347 = True


} else {
ifres9347 = False


}

ifres9346 = ifres9347


} else {
ifres9346 = False


}

var ifres9345 Obj

if True == ifres9346 {
ifres9345 = True


} else {
ifres9345 = False


}

ifres9344 = ifres9345


} else {
ifres9344 = False


}

var ifres9343 Obj

if True == ifres9344 {
ifres9343 = True


} else {
ifres9343 = False


}

ifres9342 = ifres9343


} else {
ifres9342 = False


}

if True == ifres9342 {
tmp9329 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5866)
}
__typedArg0 := V5866
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9330 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9329)
}
__typedArg0 := tmp9329
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9331 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5866)
}
__typedArg0 := V5866
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9332 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9331)
}
__typedArg0 := tmp9331
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9333 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9332)
}
__typedArg0 := tmp9332
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9334 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5866)
}
__typedArg0 := V5866
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9335 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9334)
}
__typedArg0 := tmp9334
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9336 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9335)
}
__typedArg0 := tmp9335
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9337 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp9336)
}
__typedArg0 := symlet
__typedArg1 := tmp9336
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9338 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9337, Nil)
}
__typedArg0 := tmp9337
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9339 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9333, tmp9338)
}
__typedArg0 := tmp9333
__typedArg1 := tmp9338
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9340 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9330, tmp9339)
}
__typedArg0 := tmp9330
__typedArg1 := tmp9339
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp9340)
}
__typedArg0 := symlet
__typedArg1 := tmp9340
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return(V5866)
return
}


}, 1)

tmp9369 := Call(__e, ns2_1set, symshen_4process_1let, tmp9328)


_ = tmp9369

tmp9370 := MakeNative(func(__e *ControlFlow) {
V5867 := __e.Get(1)
_ = V5867
tmp9371 := MakeNative(func(__e *ControlFlow) {
GoTo5869 := __e.Get(1)
_ = GoTo5869
tmp9406 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5867)
}
__typedArg0 := V5867
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9406 {
tmp9372 := MakeNative(func(__e *ControlFlow) {
Select5876 := __e.Get(1)
_ = Select5876
tmp9402 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5867)
}
__typedArg0 := V5867
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9403 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_8s, tmp9402)
}
__typedArg0 := sym_8s
__typedArg1 := tmp9402
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9403 {
tmp9400 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(Select5876)
}
__typedArg0 := Select5876
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9400 {
tmp9373 := MakeNative(func(__e *ControlFlow) {
Select5874 := __e.Get(1)
_ = Select5874
tmp9374 := MakeNative(func(__e *ControlFlow) {
Select5875 := __e.Get(1)
_ = Select5875
tmp9396 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(Select5875)
}
__typedArg0 := Select5875
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9396 {
tmp9375 := MakeNative(func(__e *ControlFlow) {
Select5873 := __e.Get(1)
_ = Select5873
tmp9393 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(Select5873)
}
__typedArg0 := Select5873
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9393 {
tmp9376 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8s, Select5875)
}
__typedArg0 := sym_8s
__typedArg1 := Select5875
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9377 := Call(__e, PrimFunc(symshen_4process_1_8s), tmp9376)


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
return PrimCons(Select5874, tmp9378)
}
__typedArg0 := Select5874
__typedArg1 := tmp9378
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8s, tmp9379)
}
__typedArg0 := sym_8s
__typedArg1 := tmp9379
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9391 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, Select5873)
}
__typedArg0 := Nil
__typedArg1 := Select5873
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9388 Obj

if True == tmp9391 {
tmp9390 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_2) {
return PrimIsString(Select5874)
}
__typedArg0 := Select5874
return Call(__e, PrimFunc(symstring_2), __typedArg0)
})()

var ifres9389 Obj

if True == tmp9390 {
ifres9389 = True


} else {
ifres9389 = False


}

ifres9388 = ifres9389


} else {
ifres9388 = False


}

if True == ifres9388 {
tmp9380 := MakeNative(func(__e *ControlFlow) {
W5868 := __e.Get(1)
_ = W5868
tmp9384 := Call(__e, PrimFunc(symlength), W5868)


if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_6) {
__typedN0, __typedOK0 := TypedFloat64(tmp9384)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_6) {
return TypedMaterializeBoolean((__typedN0 > __typedN1))
}}
__typedArg0 := tmp9384
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_6), __typedArg0, __typedArg1)
})() {
tmp9381 := Call(__e, PrimFunc(symappend), W5868, Select5875)


tmp9382 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8s, tmp9381)
}
__typedArg0 := sym_8s
__typedArg1 := tmp9381
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4process_1_8s), tmp9382)
return


} else {
__e.Return(V5867)
return
}


}, 1)

tmp9386 := Call(__e, PrimFunc(symexplode), Select5874)


__e.TailApply(tmp9380, tmp9386)
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5869)
return
}


}


}, 1)

tmp9394 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select5875)
}
__typedArg0 := Select5875
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(tmp9375, tmp9394)
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5869)
return
}


}, 1)

tmp9397 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select5876)
}
__typedArg0 := Select5876
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(tmp9374, tmp9397)
return


}, 1)

tmp9398 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(Select5876)
}
__typedArg0 := Select5876
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp9373, tmp9398)
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5869)
return
}


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5869)
return
}


}, 1)

tmp9404 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5867)
}
__typedArg0 := V5867
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(tmp9372, tmp9404)
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5869)
return
}


}, 1)

tmp9407 := MakeNative(func(__e *ControlFlow) {
__e.Return(V5867)
return
}, 0)

__e.TailApply(tmp9371, tmp9407)
return


}, 1)

tmp9408 := Call(__e, ns2_1set, symshen_4process_1_8s, tmp9370)


_ = tmp9408

tmp9409 := MakeNative(func(__e *ControlFlow) {
V5877 := __e.Get(1)
_ = V5877
V5878 := __e.Get(2)
_ = V5878
tmp9410 := MakeNative(func(__e *ControlFlow) {
W5879 := __e.Get(1)
_ = W5879
tmp9411 := MakeNative(func(__e *ControlFlow) {
W5880 := __e.Get(1)
_ = W5880
__e.Return(W5879)
return
}, 1)

tmp9412 := MakeNative(func(__e *ControlFlow) {
Z5881 := __e.Get(1)
_ = Z5881
__e.TailApply(PrimFunc(symshen_4_5datatype_6), Z5881)
return
}, 1)

tmp9413 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W5879, V5878)
}
__typedArg0 := W5879
__typedArg1 := V5878
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9414 := Call(__e, PrimFunc(symcompile), tmp9412, tmp9413)


__e.TailApply(tmp9411, tmp9414)
return


}, 1)

tmp9415 := Call(__e, PrimFunc(symshen_4intern_1type), V5877)


__e.TailApply(tmp9410, tmp9415)
return


}, 2)

tmp9416 := Call(__e, ns2_1set, symshen_4process_1datatype, tmp9409)


_ = tmp9416

tmp9417 := MakeNative(func(__e *ControlFlow) {
V5882 := __e.Get(1)
_ = V5882
tmp9418 := (func() Obj {
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
__typedS0, __typedOK0 := TypedString(tmp9418)
__typedS1, __typedOK1 := TypedString(MakeString("#type"))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := tmp9418
__typedArg1 := MakeString("#type")
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(tmp9418)
__typedS1, __typedOK1 := TypedString(MakeString("#type"))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := tmp9418
__typedArg1 := MakeString("#type")
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symintern), __typedArg0)
})())
return


}, 1)

tmp9420 := Call(__e, ns2_1set, symshen_4intern_1type, tmp9417)


_ = tmp9420

tmp9421 := MakeNative(func(__e *ControlFlow) {
V5883 := __e.Get(1)
_ = V5883
tmp9422 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsynonyms_d)
}
__typedArg0 := symshen_4_dsynonyms_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp9423 := Call(__e, PrimFunc(symappend), V5883, tmp9422)


tmp9424 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsynonyms_d, tmp9423)
}
__typedArg0 := symshen_4_dsynonyms_d
__typedArg1 := tmp9423
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4synonyms_1h), tmp9424)
return


}, 1)

tmp9425 := Call(__e, ns2_1set, symshen_4process_1synonyms, tmp9421)


_ = tmp9425

tmp9426 := MakeNative(func(__e *ControlFlow) {
V5884 := __e.Get(1)
_ = V5884
tmp9427 := MakeNative(func(__e *ControlFlow) {
W5885 := __e.Get(1)
_ = W5885
tmp9428 := MakeNative(func(__e *ControlFlow) {
W5887 := __e.Get(1)
_ = W5887
__e.Return(symsynonyms)
return
}, 1)

tmp9429 := Call(__e, PrimFunc(symshen_4compile_1synonyms), W5885)


tmp9430 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4demod, tmp9429)
}
__typedArg0 := symshen_4demod
__typedArg1 := tmp9429
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9431 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefine, tmp9430)
}
__typedArg0 := symdefine
__typedArg1 := tmp9430
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9432 := Call(__e, PrimFunc(symeval), tmp9431)


__e.TailApply(tmp9428, tmp9432)
return


}, 1)

tmp9433 := MakeNative(func(__e *ControlFlow) {
Z5886 := __e.Get(1)
_ = Z5886
__e.TailApply(PrimFunc(symshen_4curry_1type), Z5886)
return
}, 1)

tmp9434 := Call(__e, PrimFunc(symmap), tmp9433, V5884)


__e.TailApply(tmp9427, tmp9434)
return


}, 1)

tmp9435 := Call(__e, ns2_1set, symshen_4synonyms_1h, tmp9426)


_ = tmp9435

tmp9436 := MakeNative(func(__e *ControlFlow) {
V5890 := __e.Get(1)
_ = V5890
tmp9458 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V5890)
}
__typedArg0 := Nil
__typedArg1 := V5890
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9458 {
tmp9437 := MakeNative(func(__e *ControlFlow) {
W5891 := __e.Get(1)
_ = W5891
tmp9438 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W5891, Nil)
}
__typedArg0 := W5891
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9439 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_6, tmp9438)
}
__typedArg0 := sym_1_6
__typedArg1 := tmp9438
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W5891, tmp9439)
}
__typedArg0 := W5891
__typedArg1 := tmp9439
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp9440 := Call(__e, PrimFunc(symgensym), symX)


__e.TailApply(tmp9437, tmp9440)
return


} else {
tmp9456 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5890)
}
__typedArg0 := V5890
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9452 Obj

if True == tmp9456 {
tmp9454 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5890)
}
__typedArg0 := V5890
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9455 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9454)
}
__typedArg0 := tmp9454
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9453 Obj

if True == tmp9455 {
ifres9453 = True


} else {
ifres9453 = False


}

ifres9452 = ifres9453


} else {
ifres9452 = False


}

if True == ifres9452 {
tmp9441 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5890)
}
__typedArg0 := V5890
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9442 := Call(__e, PrimFunc(symshen_4rcons__form), tmp9441)


tmp9443 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5890)
}
__typedArg0 := V5890
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9444 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9443)
}
__typedArg0 := tmp9443
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9445 := Call(__e, PrimFunc(symshen_4rcons__form), tmp9444)


tmp9446 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5890)
}
__typedArg0 := V5890
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9447 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9446)
}
__typedArg0 := tmp9446
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9448 := Call(__e, PrimFunc(symshen_4compile_1synonyms), tmp9447)


tmp9449 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9445, tmp9448)
}
__typedArg0 := tmp9445
__typedArg1 := tmp9448
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9450 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_6, tmp9449)
}
__typedArg0 := sym_1_6
__typedArg1 := tmp9449
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9442, tmp9450)
}
__typedArg0 := tmp9442
__typedArg1 := tmp9450
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

tmp9459 := Call(__e, ns2_1set, symshen_4compile_1synonyms, tmp9436)


_ = tmp9459

tmp9460 := MakeNative(func(__e *ControlFlow) {
V5892 := __e.Get(1)
_ = V5892
tmp9461 := MakeNative(func(__e *ControlFlow) {
GoTo5893 := __e.Get(1)
_ = GoTo5893
tmp9489 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5892)
}
__typedArg0 := V5892
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9489 {
tmp9462 := MakeNative(func(__e *ControlFlow) {
Select5900 := __e.Get(1)
_ = Select5900
tmp9485 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5892)
}
__typedArg0 := V5892
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9486 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_c_4, tmp9485)
}
__typedArg0 := sym_c_4
__typedArg1 := tmp9485
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9486 {
tmp9483 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(Select5900)
}
__typedArg0 := Select5900
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9483 {
tmp9463 := MakeNative(func(__e *ControlFlow) {
Select5898 := __e.Get(1)
_ = Select5898
tmp9464 := MakeNative(func(__e *ControlFlow) {
Select5899 := __e.Get(1)
_ = Select5899
tmp9479 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(Select5899)
}
__typedArg0 := Select5899
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9479 {
tmp9465 := MakeNative(func(__e *ControlFlow) {
Select5897 := __e.Get(1)
_ = Select5897
tmp9476 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(Select5897)
}
__typedArg0 := Select5897
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9476 {
tmp9466 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_c_4, Select5899)
}
__typedArg0 := sym_c_4
__typedArg1 := Select5899
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9467 := Call(__e, PrimFunc(symshen_4process_1lambda), tmp9466)


tmp9468 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9467, Nil)
}
__typedArg0 := tmp9467
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9469 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Select5898, tmp9468)
}
__typedArg0 := Select5898
__typedArg1 := tmp9468
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp9469)
}
__typedArg0 := symlambda
__typedArg1 := tmp9469
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9474 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, Select5897)
}
__typedArg0 := Nil
__typedArg1 := Select5897
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9474 {
tmp9472 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(Select5898)
}
__typedArg0 := Select5898
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

if True == tmp9472 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, Select5900)
}
__typedArg0 := symlambda
__typedArg1 := Select5900
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return
} else {
tmp9470 := Call(__e, PrimFunc(symshen_4app), Select5898, MakeString(" is not a variable\n"), symshen_4s)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(tmp9470)
}
__typedArg0 := tmp9470
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5893)
return
}


}


}, 1)

tmp9477 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select5899)
}
__typedArg0 := Select5899
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(tmp9465, tmp9477)
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5893)
return
}


}, 1)

tmp9480 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select5900)
}
__typedArg0 := Select5900
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(tmp9464, tmp9480)
return


}, 1)

tmp9481 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(Select5900)
}
__typedArg0 := Select5900
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp9463, tmp9481)
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5893)
return
}


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5893)
return
}


}, 1)

tmp9487 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5892)
}
__typedArg0 := V5892
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(tmp9462, tmp9487)
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5893)
return
}


}, 1)

tmp9490 := MakeNative(func(__e *ControlFlow) {
__e.Return(V5892)
return
}, 0)

__e.TailApply(tmp9461, tmp9490)
return


}, 1)

tmp9491 := Call(__e, ns2_1set, symshen_4process_1lambda, tmp9460)


_ = tmp9491

tmp9492 := MakeNative(func(__e *ControlFlow) {
V5903 := __e.Get(1)
_ = V5903
tmp9493 := MakeNative(func(__e *ControlFlow) {
GoTo5904 := __e.Get(1)
_ = GoTo5904
tmp9533 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5903)
}
__typedArg0 := V5903
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9533 {
tmp9494 := MakeNative(func(__e *ControlFlow) {
Select5912 := __e.Get(1)
_ = Select5912
tmp9529 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5903)
}
__typedArg0 := V5903
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9530 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcases, tmp9529)
}
__typedArg0 := symcases
__typedArg1 := tmp9529
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9530 {
tmp9527 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(Select5912)
}
__typedArg0 := Select5912
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9527 {
tmp9495 := MakeNative(func(__e *ControlFlow) {
Select5910 := __e.Get(1)
_ = Select5910
tmp9496 := MakeNative(func(__e *ControlFlow) {
Select5911 := __e.Get(1)
_ = Select5911
tmp9523 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(True, Select5910)
}
__typedArg0 := True
__typedArg1 := Select5910
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9520 Obj

if True == tmp9523 {
tmp9522 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(Select5911)
}
__typedArg0 := Select5911
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9521 Obj

if True == tmp9522 {
ifres9521 = True


} else {
ifres9521 = False


}

ifres9520 = ifres9521


} else {
ifres9520 = False


}

if True == ifres9520 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(Select5911)
}
__typedArg0 := Select5911
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return
} else {
tmp9497 := MakeNative(func(__e *ControlFlow) {
GoTo5907 := __e.Get(1)
_ = GoTo5907
tmp9515 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(Select5911)
}
__typedArg0 := Select5911
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9515 {
tmp9498 := MakeNative(func(__e *ControlFlow) {
Select5908 := __e.Get(1)
_ = Select5908
tmp9499 := MakeNative(func(__e *ControlFlow) {
Select5909 := __e.Get(1)
_ = Select5909
tmp9511 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, Select5909)
}
__typedArg0 := Nil
__typedArg1 := Select5909
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9511 {
tmp9500 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeString("error: cases exhausted"), Nil)
}
__typedArg0 := MakeString("error: cases exhausted")
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9501 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsimple_1error, tmp9500)
}
__typedArg0 := symsimple_1error
__typedArg1 := tmp9500
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9502 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9501, Nil)
}
__typedArg0 := tmp9501
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9503 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Select5908, tmp9502)
}
__typedArg0 := Select5908
__typedArg1 := tmp9502
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9504 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Select5910, tmp9503)
}
__typedArg0 := Select5910
__typedArg1 := tmp9503
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp9504)
}
__typedArg0 := symif
__typedArg1 := tmp9504
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9505 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcases, Select5909)
}
__typedArg0 := symcases
__typedArg1 := Select5909
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9506 := Call(__e, PrimFunc(symshen_4process_1cases), tmp9505)


tmp9507 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9506, Nil)
}
__typedArg0 := tmp9506
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9508 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Select5908, tmp9507)
}
__typedArg0 := Select5908
__typedArg1 := tmp9507
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9509 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Select5910, tmp9508)
}
__typedArg0 := Select5910
__typedArg1 := tmp9508
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp9509)
}
__typedArg0 := symif
__typedArg1 := tmp9509
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


}, 1)

tmp9512 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select5911)
}
__typedArg0 := Select5911
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(tmp9499, tmp9512)
return


}, 1)

tmp9513 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(Select5911)
}
__typedArg0 := Select5911
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp9498, tmp9513)
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5907)
return
}


}, 1)

tmp9516 := MakeNative(func(__e *ControlFlow) {
tmp9518 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, Select5911)
}
__typedArg0 := Nil
__typedArg1 := Select5911
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9518 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("error: odd number of case elements\n"))
}
__typedArg0 := MakeString("error: odd number of case elements\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
} else {
__e.TailApply(PrimFunc(symthaw), GoTo5904)
return
}


}, 0)

__e.TailApply(tmp9497, tmp9516)
return


}


}, 1)

tmp9524 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select5912)
}
__typedArg0 := Select5912
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(tmp9496, tmp9524)
return


}, 1)

tmp9525 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(Select5912)
}
__typedArg0 := Select5912
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp9495, tmp9525)
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5904)
return
}


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5904)
return
}


}, 1)

tmp9531 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5903)
}
__typedArg0 := V5903
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(tmp9494, tmp9531)
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5904)
return
}


}, 1)

tmp9534 := MakeNative(func(__e *ControlFlow) {
__e.Return(V5903)
return
}, 0)

__e.TailApply(tmp9493, tmp9534)
return


}, 1)

tmp9535 := Call(__e, ns2_1set, symshen_4process_1cases, tmp9492)


_ = tmp9535

tmp9536 := MakeNative(func(__e *ControlFlow) {
V5913 := __e.Get(1)
_ = V5913
tmp9537 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symrun, Nil)
}
__typedArg0 := symrun
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9538 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symget_1time, tmp9537)
}
__typedArg0 := symget_1time
__typedArg1 := tmp9537
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9539 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symrun, Nil)
}
__typedArg0 := symrun
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9540 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symget_1time, tmp9539)
}
__typedArg0 := symget_1time
__typedArg1 := tmp9539
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9541 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symStart, Nil)
}
__typedArg0 := symStart
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9542 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symFinish, tmp9541)
}
__typedArg0 := symFinish
__typedArg1 := tmp9541
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9543 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1, tmp9542)
}
__typedArg0 := sym_1
__typedArg1 := tmp9542
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9544 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symTime, Nil)
}
__typedArg0 := symTime
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9545 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp9544)
}
__typedArg0 := symstr
__typedArg1 := tmp9544
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9546 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeString(" secs\n"), Nil)
}
__typedArg0 := MakeString(" secs\n")
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9547 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9545, tmp9546)
}
__typedArg0 := tmp9545
__typedArg1 := tmp9546
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9548 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcn, tmp9547)
}
__typedArg0 := symcn
__typedArg1 := tmp9547
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9549 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9548, Nil)
}
__typedArg0 := tmp9548
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9550 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeString("\nrun time: "), tmp9549)
}
__typedArg0 := MakeString("\nrun time: ")
__typedArg1 := tmp9549
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9551 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcn, tmp9550)
}
__typedArg0 := symcn
__typedArg1 := tmp9550
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9552 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstoutput, Nil)
}
__typedArg0 := symstoutput
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9553 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9552, Nil)
}
__typedArg0 := tmp9552
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9554 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9551, tmp9553)
}
__typedArg0 := tmp9551
__typedArg1 := tmp9553
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9555 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sympr, tmp9554)
}
__typedArg0 := sympr
__typedArg1 := tmp9554
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9556 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symResult, Nil)
}
__typedArg0 := symResult
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9557 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9555, tmp9556)
}
__typedArg0 := tmp9555
__typedArg1 := tmp9556
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9558 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symMessage, tmp9557)
}
__typedArg0 := symMessage
__typedArg1 := tmp9557
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9559 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9543, tmp9558)
}
__typedArg0 := tmp9543
__typedArg1 := tmp9558
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9560 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symTime, tmp9559)
}
__typedArg0 := symTime
__typedArg1 := tmp9559
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9561 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9540, tmp9560)
}
__typedArg0 := tmp9540
__typedArg1 := tmp9560
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9562 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symFinish, tmp9561)
}
__typedArg0 := symFinish
__typedArg1 := tmp9561
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9563 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5913, tmp9562)
}
__typedArg0 := V5913
__typedArg1 := tmp9562
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9564 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symResult, tmp9563)
}
__typedArg0 := symResult
__typedArg1 := tmp9563
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9565 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9538, tmp9564)
}
__typedArg0 := tmp9538
__typedArg1 := tmp9564
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9566 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symStart, tmp9565)
}
__typedArg0 := symStart
__typedArg1 := tmp9565
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp9566)
}
__typedArg0 := symlet
__typedArg1 := tmp9566
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp9567 := Call(__e, ns2_1set, symshen_4process_1time, tmp9536)


_ = tmp9567

tmp9568 := MakeNative(func(__e *ControlFlow) {
V5914 := __e.Get(1)
_ = V5914
tmp9594 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5914)
}
__typedArg0 := V5914
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9579 Obj

if True == tmp9594 {
tmp9592 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5914)
}
__typedArg0 := V5914
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9593 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9592)
}
__typedArg0 := tmp9592
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9581 Obj

if True == tmp9593 {
tmp9589 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5914)
}
__typedArg0 := V5914
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9590 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9589)
}
__typedArg0 := tmp9589
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9591 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9590)
}
__typedArg0 := tmp9590
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9583 Obj

if True == tmp9591 {
tmp9585 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5914)
}
__typedArg0 := V5914
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9586 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9585)
}
__typedArg0 := tmp9585
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9587 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9586)
}
__typedArg0 := tmp9586
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9588 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9587)
}
__typedArg0 := tmp9587
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9584 Obj

if True == tmp9588 {
ifres9584 = True


} else {
ifres9584 = False


}

ifres9583 = ifres9584


} else {
ifres9583 = False


}

var ifres9582 Obj

if True == ifres9583 {
ifres9582 = True


} else {
ifres9582 = False


}

ifres9581 = ifres9582


} else {
ifres9581 = False


}

var ifres9580 Obj

if True == ifres9581 {
ifres9580 = True


} else {
ifres9580 = False


}

ifres9579 = ifres9580


} else {
ifres9579 = False


}

if True == ifres9579 {
tmp9569 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5914)
}
__typedArg0 := V5914
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9570 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5914)
}
__typedArg0 := V5914
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9571 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9570)
}
__typedArg0 := tmp9570
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9572 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5914)
}
__typedArg0 := V5914
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9573 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5914)
}
__typedArg0 := V5914
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9574 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9573)
}
__typedArg0 := tmp9573
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9575 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9572, tmp9574)
}
__typedArg0 := tmp9572
__typedArg1 := tmp9574
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9576 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9575, Nil)
}
__typedArg0 := tmp9575
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9577 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9571, tmp9576)
}
__typedArg0 := tmp9571
__typedArg1 := tmp9576
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9569, tmp9577)
}
__typedArg0 := tmp9569
__typedArg1 := tmp9577
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return(V5914)
return
}


}, 1)

tmp9595 := Call(__e, ns2_1set, symshen_4process_1assoc, tmp9568)


_ = tmp9595

tmp9596 := MakeNative(func(__e *ControlFlow) {
V5915 := __e.Get(1)
_ = V5915
tmp9597 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(V5915)
}
__typedArg0 := V5915
return Call(__e, PrimFunc(symstr), __typedArg0)
})()

tmp9598 := Call(__e, PrimFunc(symshen_4mu_1h), tmp9597)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(tmp9598)
}
__typedArg0 := tmp9598
return Call(__e, PrimFunc(symintern), __typedArg0)
})())
return


}, 1)

tmp9599 := Call(__e, ns2_1set, symshen_4make_1uppercase, tmp9596)


_ = tmp9599

tmp9600 := MakeNative(func(__e *ControlFlow) {
V5916 := __e.Get(1)
_ = V5916
tmp9619 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString(""), V5916)
}
__typedArg0 := MakeString("")
__typedArg1 := V5916
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9619 {
__e.Return(MakeString(""))
return
} else {
tmp9617 := Call(__e, PrimFunc(symshen_4_7string_2), V5916)


if True == tmp9617 {
tmp9601 := MakeNative(func(__e *ControlFlow) {
W5917 := __e.Get(1)
_ = W5917
tmp9602 := MakeNative(func(__e *ControlFlow) {
W5918 := __e.Get(1)
_ = W5918
tmp9603 := MakeNative(func(__e *ControlFlow) {
W5919 := __e.Get(1)
_ = W5919
tmp9605 := Call(__e, PrimFunc(symshen_4mu_1h), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V5916)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V5916
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


__e.TailApply(PrimFunc(sym_8s), W5919, tmp9605)
return


}, 1)

tmp9612 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_6_a) {
__typedN0, __typedOK0 := TypedFloat64(W5917)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(97))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_6_a) {
return TypedMaterializeBoolean((__typedN0 >= __typedN1))
}}
__typedArg0 := W5917
__typedArg1 := MakeNumber(97)
return Call(__e, PrimFunc(sym_6_a), __typedArg0, __typedArg1)
})()

var ifres9609 Obj

if True == tmp9612 {
tmp9611 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_a) {
__typedN0, __typedOK0 := TypedFloat64(W5917)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(122))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_5_a) {
return TypedMaterializeBoolean((__typedN0 <= __typedN1))
}}
__typedArg0 := W5917
__typedArg1 := MakeNumber(122)
return Call(__e, PrimFunc(sym_5_a), __typedArg0, __typedArg1)
})()

var ifres9610 Obj

if True == tmp9611 {
ifres9610 = True


} else {
ifres9610 = False


}

ifres9609 = ifres9610


} else {
ifres9609 = False


}

var ifres9606 Obj

if True == ifres9609 {
tmp9607 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symn_1_6string) {
return PrimNumberToString(W5918)
}
__typedArg0 := W5918
return Call(__e, PrimFunc(symn_1_6string), __typedArg0)
})()

ifres9606 = tmp9607


} else {
tmp9608 := Call(__e, PrimFunc(symhdstr), V5916)


ifres9606 = tmp9608


}

__e.TailApply(tmp9603, ifres9606)
return


}, 1)

__e.TailApply(tmp9602, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(W5917)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(32))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := W5917
__typedArg1 := MakeNumber(32)
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp9614 := Call(__e, PrimFunc(symhdstr), V5916)


tmp9615 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_1_6n) {
return PrimStringToNumber(tmp9614)
}
__typedArg0 := tmp9614
return Call(__e, PrimFunc(symstring_1_6n), __typedArg0)
})()

__e.TailApply(tmp9601, tmp9615)
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

tmp9620 := Call(__e, ns2_1set, symshen_4mu_1h, tmp9600)


_ = tmp9620

tmp9621 := MakeNative(func(__e *ControlFlow) {
V5920 := __e.Get(1)
_ = V5920
V5921 := __e.Get(2)
_ = V5921
tmp9622 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dmacros_d)
}
__typedArg0 := sym_dmacros_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp9623 := Call(__e, PrimFunc(symshen_4update_1assoc), V5920, V5921, tmp9622)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(sym_dmacros_d, tmp9623)
}
__typedArg0 := sym_dmacros_d
__typedArg1 := tmp9623
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return


}, 2)

tmp9624 := Call(__e, ns2_1set, symshen_4record_1macro, tmp9621)


_ = tmp9624

tmp9625 := MakeNative(func(__e *ControlFlow) {
V5931 := __e.Get(1)
_ = V5931
V5932 := __e.Get(2)
_ = V5932
V5933 := __e.Get(3)
_ = V5933
tmp9645 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V5933)
}
__typedArg0 := Nil
__typedArg1 := V5933
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9645 {
tmp9626 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5931, V5932)
}
__typedArg0 := V5931
__typedArg1 := V5932
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9626, Nil)
}
__typedArg0 := tmp9626
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9627 := MakeNative(func(__e *ControlFlow) {
GoTo5934 := __e.Get(1)
_ = GoTo5934
tmp9642 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5933)
}
__typedArg0 := V5933
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9642 {
tmp9628 := MakeNative(func(__e *ControlFlow) {
Select5935 := __e.Get(1)
_ = Select5935
tmp9629 := MakeNative(func(__e *ControlFlow) {
Select5936 := __e.Get(1)
_ = Select5936
tmp9638 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(Select5935)
}
__typedArg0 := Select5935
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9634 Obj

if True == tmp9638 {
tmp9636 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(Select5935)
}
__typedArg0 := Select5935
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9637 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V5931, tmp9636)
}
__typedArg0 := V5931
__typedArg1 := tmp9636
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9635 Obj

if True == tmp9637 {
ifres9635 = True


} else {
ifres9635 = False


}

ifres9634 = ifres9635


} else {
ifres9634 = False


}

if True == ifres9634 {
tmp9630 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(Select5935)
}
__typedArg0 := Select5935
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9631 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9630, V5932)
}
__typedArg0 := tmp9630
__typedArg1 := V5932
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9631, Select5936)
}
__typedArg0 := tmp9631
__typedArg1 := Select5936
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9632 := Call(__e, PrimFunc(symshen_4update_1assoc), V5931, V5932, Select5936)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Select5935, tmp9632)
}
__typedArg0 := Select5935
__typedArg1 := tmp9632
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


}, 1)

tmp9639 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5933)
}
__typedArg0 := V5933
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(tmp9629, tmp9639)
return


}, 1)

tmp9640 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5933)
}
__typedArg0 := V5933
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp9628, tmp9640)
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5934)
return
}


}, 1)

tmp9643 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.update-assoc"))
}
__typedArg0 := MakeString("implementation error in shen.update-assoc")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}, 0)

__e.TailApply(tmp9627, tmp9643)
return


}


}, 3)

tmp9646 := Call(__e, ns2_1set, symshen_4update_1assoc, tmp9625)


_ = tmp9646

tmp9647 := MakeNative(func(__e *ControlFlow) {
tmp9655 := Call(__e, PrimFunc(symstinput))


tmp9656 := Call(__e, PrimFunc(symshen_4char_1stinput_2), tmp9655)


if True == tmp9656 {
tmp9648 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstinput, Nil)
}
__typedArg0 := symstinput
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9649 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9648, Nil)
}
__typedArg0 := tmp9648
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9650 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4read_1unit_1string, tmp9649)
}
__typedArg0 := symshen_4read_1unit_1string
__typedArg1 := tmp9649
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9651 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9650, Nil)
}
__typedArg0 := tmp9650
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring_1_6n, tmp9651)
}
__typedArg0 := symstring_1_6n
__typedArg1 := tmp9651
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9652 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstinput, Nil)
}
__typedArg0 := symstinput
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9653 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9652, Nil)
}
__typedArg0 := tmp9652
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symread_1byte, tmp9653)
}
__typedArg0 := symread_1byte
__typedArg1 := tmp9653
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


}, 0)

tmp9657 := Call(__e, ns2_1set, symshen_4process_1read_1byte, tmp9647)


_ = tmp9657

tmp9658 := MakeNative(func(__e *ControlFlow) {
V5937 := __e.Get(1)
_ = V5937
tmp9659 := MakeNative(func(__e *ControlFlow) {
W5938 := __e.Get(1)
_ = W5938
tmp9660 := MakeNative(func(__e *ControlFlow) {
W5939 := __e.Get(1)
_ = W5939
tmp9661 := MakeNative(func(__e *ControlFlow) {
W5940 := __e.Get(1)
_ = W5940
tmp9662 := MakeNative(func(__e *ControlFlow) {
W5941 := __e.Get(1)
_ = W5941
tmp9663 := MakeNative(func(__e *ControlFlow) {
W5942 := __e.Get(1)
_ = W5942
tmp9664 := MakeNative(func(__e *ControlFlow) {
W5944 := __e.Get(1)
_ = W5944
tmp9665 := MakeNative(func(__e *ControlFlow) {
W5945 := __e.Get(1)
_ = W5945
tmp9666 := MakeNative(func(__e *ControlFlow) {
W5946 := __e.Get(1)
_ = W5946
tmp9667 := MakeNative(func(__e *ControlFlow) {
W5947 := __e.Get(1)
_ = W5947
tmp9668 := MakeNative(func(__e *ControlFlow) {
W5948 := __e.Get(1)
_ = W5948
tmp9669 := MakeNative(func(__e *ControlFlow) {
W5949 := __e.Get(1)
_ = W5949
tmp9670 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W5941, Nil)
}
__typedArg0 := W5941
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9671 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W5940, tmp9670)
}
__typedArg0 := W5940
__typedArg1 := tmp9670
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9672 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W5939, tmp9671)
}
__typedArg0 := W5939
__typedArg1 := tmp9671
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9673 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W5938, tmp9672)
}
__typedArg0 := W5938
__typedArg1 := tmp9672
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W5949, tmp9673)
}
__typedArg0 := W5949
__typedArg1 := tmp9673
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp9674 := Call(__e, PrimFunc(symshen_4continue), W5944, W5942, W5945, W5946, W5947, W5948)


tmp9675 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9674, Nil)
}
__typedArg0 := tmp9674
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9676 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W5948, tmp9675)
}
__typedArg0 := W5948
__typedArg1 := tmp9675
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9677 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp9676)
}
__typedArg0 := symlambda
__typedArg1 := tmp9676
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9678 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9677, Nil)
}
__typedArg0 := tmp9677
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9679 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W5947, tmp9678)
}
__typedArg0 := W5947
__typedArg1 := tmp9678
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9680 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp9679)
}
__typedArg0 := symlambda
__typedArg1 := tmp9679
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9681 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9680, Nil)
}
__typedArg0 := tmp9680
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9682 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W5946, tmp9681)
}
__typedArg0 := W5946
__typedArg1 := tmp9681
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9683 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp9682)
}
__typedArg0 := symlambda
__typedArg1 := tmp9682
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9684 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9683, Nil)
}
__typedArg0 := tmp9683
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9685 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W5945, tmp9684)
}
__typedArg0 := W5945
__typedArg1 := tmp9684
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9686 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp9685)
}
__typedArg0 := symlambda
__typedArg1 := tmp9685
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(tmp9669, tmp9686)
return


}, 1)

tmp9687 := Call(__e, PrimFunc(symgensym), symC)


__e.TailApply(tmp9668, tmp9687)
return


}, 1)

tmp9688 := Call(__e, PrimFunc(symgensym), symK)


__e.TailApply(tmp9667, tmp9688)
return


}, 1)

tmp9689 := Call(__e, PrimFunc(symgensym), symL)


__e.TailApply(tmp9666, tmp9689)
return


}, 1)

tmp9690 := Call(__e, PrimFunc(symgensym), symV)


__e.TailApply(tmp9665, tmp9690)
return


}, 1)

tmp9691 := Call(__e, PrimFunc(symshen_4received), V5937)


__e.TailApply(tmp9664, tmp9691)
return


}, 1)

tmp9692 := MakeNative(func(__e *ControlFlow) {
Z5943 := __e.Get(1)
_ = Z5943
__e.TailApply(PrimFunc(symshen_4_5body_6), Z5943)
return
}, 1)

tmp9693 := Call(__e, PrimFunc(symcompile), tmp9692, V5937)


__e.TailApply(tmp9663, tmp9693)
return


}, 1)

tmp9694 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(True, Nil)
}
__typedArg0 := True
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9695 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfreeze, tmp9694)
}
__typedArg0 := symfreeze
__typedArg1 := tmp9694
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(tmp9662, tmp9695)
return


}, 1)

__e.TailApply(tmp9661, MakeNumber(0))
return


}, 1)

tmp9696 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(0), Nil)
}
__typedArg0 := MakeNumber(0)
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9697 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp9696)
}
__typedArg0 := symvector
__typedArg1 := tmp9696
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9698 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9697, Nil)
}
__typedArg0 := tmp9697
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9699 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(0), tmp9698)
}
__typedArg0 := MakeNumber(0)
__typedArg1 := tmp9698
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9700 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(True, tmp9699)
}
__typedArg0 := True
__typedArg1 := tmp9699
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9701 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8v, tmp9700)
}
__typedArg0 := sym_8v
__typedArg1 := tmp9700
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(tmp9660, tmp9701)
return


}, 1)

tmp9702 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4prolog_1vector, Nil)
}
__typedArg0 := symshen_4prolog_1vector
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(tmp9659, tmp9702)
return


}, 1)

tmp9703 := Call(__e, ns2_1set, symshen_4call_1prolog, tmp9658)


_ = tmp9703

tmp9704 := MakeNative(func(__e *ControlFlow) {
V5952 := __e.Get(1)
_ = V5952
tmp9705 := MakeNative(func(__e *ControlFlow) {
GoTo5953 := __e.Get(1)
_ = GoTo5953
tmp9722 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5952)
}
__typedArg0 := V5952
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9722 {
tmp9706 := MakeNative(func(__e *ControlFlow) {
Select5954 := __e.Get(1)
_ = Select5954
tmp9707 := MakeNative(func(__e *ControlFlow) {
Select5955 := __e.Get(1)
_ = Select5955
tmp9718 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symreceive, Select5954)
}
__typedArg0 := symreceive
__typedArg1 := Select5954
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9711 Obj

if True == tmp9718 {
tmp9717 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(Select5955)
}
__typedArg0 := Select5955
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9713 Obj

if True == tmp9717 {
tmp9715 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(Select5955)
}
__typedArg0 := Select5955
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9716 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp9715)
}
__typedArg0 := Nil
__typedArg1 := tmp9715
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9714 Obj

if True == tmp9716 {
ifres9714 = True


} else {
ifres9714 = False


}

ifres9713 = ifres9714


} else {
ifres9713 = False


}

var ifres9712 Obj

if True == ifres9713 {
ifres9712 = True


} else {
ifres9712 = False


}

ifres9711 = ifres9712


} else {
ifres9711 = False


}

if True == ifres9711 {
__e.Return(Select5955)
return
} else {
tmp9708 := Call(__e, PrimFunc(symshen_4received), Select5954)


tmp9709 := Call(__e, PrimFunc(symshen_4received), Select5955)


__e.TailApply(PrimFunc(symunion), tmp9708, tmp9709)
return


}


}, 1)

tmp9719 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5952)
}
__typedArg0 := V5952
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(tmp9707, tmp9719)
return


}, 1)

tmp9720 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5952)
}
__typedArg0 := V5952
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp9706, tmp9720)
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5953)
return
}


}, 1)

tmp9723 := MakeNative(func(__e *ControlFlow) {
__e.Return(Nil)
return
}, 0)

__e.TailApply(tmp9705, tmp9723)
return


}, 1)

tmp9724 := Call(__e, ns2_1set, symshen_4received, tmp9704)


_ = tmp9724

tmp9725 := MakeNative(func(__e *ControlFlow) {
tmp9726 := MakeNative(func(__e *ControlFlow) {
W5956 := __e.Get(1)
_ = W5956
tmp9727 := MakeNative(func(__e *ControlFlow) {
W5957 := __e.Get(1)
_ = W5957
tmp9728 := MakeNative(func(__e *ControlFlow) {
W5958 := __e.Get(1)
_ = W5958
__e.Return(W5958)
return
}, 1)

tmp9729 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(W5956, MakeNumber(1), MakeNumber(2))
}
__typedArg0 := W5956
__typedArg1 := MakeNumber(1)
__typedArg2 := MakeNumber(2)
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

__e.TailApply(tmp9728, tmp9729)
return


}, 1)

tmp9730 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(W5956, MakeNumber(0), symshen_4print_1prolog_1vector)
}
__typedArg0 := W5956
__typedArg1 := MakeNumber(0)
__typedArg2 := symshen_4print_1prolog_1vector
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

__e.TailApply(tmp9727, tmp9730)
return


}, 1)

tmp9731 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dprolog_1memory_d)
}
__typedArg0 := symshen_4_dprolog_1memory_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp9732 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symabsvector) {
return PrimAbsvector(tmp9731)
}
__typedArg0 := tmp9731
return Call(__e, PrimFunc(symabsvector), __typedArg0)
})()

__e.TailApply(tmp9726, tmp9732)
return


}, 0)

tmp9733 := Call(__e, ns2_1set, symshen_4prolog_1vector, tmp9725)


_ = tmp9733

tmp9734 := MakeNative(func(__e *ControlFlow) {
V5959 := __e.Get(1)
_ = V5959
__e.Return(V5959)
return
}, 1)

tmp9735 := Call(__e, ns2_1set, symreceive, tmp9734)


_ = tmp9735

tmp9736 := MakeNative(func(__e *ControlFlow) {
V5960 := __e.Get(1)
_ = V5960
tmp9744 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5960)
}
__typedArg0 := V5960
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9744 {
tmp9737 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5960)
}
__typedArg0 := V5960
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9738 := Call(__e, PrimFunc(symshen_4rcons__form), tmp9737)


tmp9739 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5960)
}
__typedArg0 := V5960
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9740 := Call(__e, PrimFunc(symshen_4rcons__form), tmp9739)


tmp9741 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9740, Nil)
}
__typedArg0 := tmp9740
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9742 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9738, tmp9741)
}
__typedArg0 := tmp9738
__typedArg1 := tmp9741
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp9742)
}
__typedArg0 := symcons
__typedArg1 := tmp9742
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return(V5960)
return
}


}, 1)

tmp9745 := Call(__e, ns2_1set, symshen_4rcons__form, tmp9736)


_ = tmp9745

tmp9746 := MakeNative(func(__e *ControlFlow) {
V5961 := __e.Get(1)
_ = V5961
tmp9753 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5961)
}
__typedArg0 := V5961
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9753 {
tmp9747 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5961)
}
__typedArg0 := V5961
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9748 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5961)
}
__typedArg0 := V5961
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9749 := Call(__e, PrimFunc(symshen_4tuple_1up), tmp9748)


tmp9750 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9749, Nil)
}
__typedArg0 := tmp9749
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9751 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9747, tmp9750)
}
__typedArg0 := tmp9747
__typedArg1 := tmp9750
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8p, tmp9751)
}
__typedArg0 := sym_8p
__typedArg1 := tmp9751
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return(V5961)
return
}


}, 1)

tmp9754 := Call(__e, ns2_1set, symshen_4tuple_1up, tmp9746)


_ = tmp9754

tmp9755 := MakeNative(func(__e *ControlFlow) {
V5962 := __e.Get(1)
_ = V5962
tmp9756 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dmacros_d)
}
__typedArg0 := sym_dmacros_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp9757 := Call(__e, PrimFunc(symassoc), V5962, tmp9756)


tmp9758 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dmacros_d)
}
__typedArg0 := sym_dmacros_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp9759 := Call(__e, PrimFunc(symremove), tmp9757, tmp9758)


tmp9760 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(sym_dmacros_d, tmp9759)
}
__typedArg0 := sym_dmacros_d
__typedArg1 := tmp9759
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp9760

__e.Return(V5962)
return


}, 1)

__e.TailApply(ns2_1set, symundefmacro, tmp9755)
return




}, 0)

