package main

import . "github.com/pyrex41/shen-go/kl"

var ReaderMain = MakeNative(func(__e *ControlFlow) {
tmp4865 := MakeNative(func(__e *ControlFlow) {
V2196 := __e.Get(1)
_ = V2196
tmp4866 := MakeNative(func(__e *ControlFlow) {
W2197 := __e.Get(1)
_ = W2197
tmp4867 := MakeNative(func(__e *ControlFlow) {
W2198 := __e.Get(1)
_ = W2198
tmp4868 := MakeNative(func(__e *ControlFlow) {
W2201 := __e.Get(1)
_ = W2201
__e.Return(W2201)
return
}, 1)

tmp4869 := Call(__e, PrimFunc(symshen_4process_1sexprs), W2198)


__e.TailApply(tmp4868, tmp4869)
return


}, 1)

tmp4870 := MakeNative(func(__e *ControlFlow) {
tmp4871 := MakeNative(func(__e *ControlFlow) {
Z2199 := __e.Get(1)
_ = Z2199
__e.TailApply(PrimFunc(symshen_4_5s_1exprs_6), Z2199)
return
}, 1)

__e.TailApply(PrimFunc(symcompile), tmp4871, W2197)
return


}, 0)

tmp4872 := MakeNative(func(__e *ControlFlow) {
Z2200 := __e.Get(1)
_ = Z2200
tmp4873 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dresidue_d)
}
__typedArg0 := symshen_4_dresidue_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4reader_1error), tmp4873)
return


}, 1)

tmp4874 := Call(__e, try_1catch, tmp4870, tmp4872)


__e.TailApply(tmp4867, tmp4874)
return


}, 1)

tmp4875 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symread_1file_1as_1bytelist) {
return PrimReadFileAsByteList(V2196)
}
__typedArg0 := V2196
return Call(__e, PrimFunc(symread_1file_1as_1bytelist), __typedArg0)
})()

__e.TailApply(tmp4866, tmp4875)
return


}, 1)

tmp4876 := Call(__e, ns2_1set, symread_1file, tmp4865)


_ = tmp4876

tmp4877 := MakeNative(func(__e *ControlFlow) {
V2202 := __e.Get(1)
_ = V2202
tmp4878 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dmaximum_1print_1sequence_1size_d)
}
__typedArg0 := sym_dmaximum_1print_1sequence_1size_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp4879 := Call(__e, PrimFunc(symshen_4reader_1error_1message), tmp4878, MakeNumber(0), V2202)


tmp4881 := Call(__e, PrimFunc(symshen_4proc_1nl), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("reader error near here: "))
__typedS1, __typedOK1 := TypedString(tmp4879)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("reader error near here: ")
__typedArg1 := tmp4879
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(tmp4881)
}
__typedArg0 := tmp4881
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}, 1)

tmp4882 := Call(__e, ns2_1set, symshen_4reader_1error, tmp4877)


_ = tmp4882

tmp4883 := MakeNative(func(__e *ControlFlow) {
V2210 := __e.Get(1)
_ = V2210
V2211 := __e.Get(2)
_ = V2211
V2212 := __e.Get(3)
_ = V2212
tmp4894 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V2212)
}
__typedArg0 := Nil
__typedArg1 := V2212
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4894 {
__e.Return(MakeString(""))
return
} else {
tmp4892 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V2210, V2211)
}
__typedArg0 := V2210
__typedArg1 := V2211
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4892 {
__e.Return(MakeString(""))
return
} else {
tmp4890 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2212)
}
__typedArg0 := V2212
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp4890 {
tmp4884 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2212)
}
__typedArg0 := V2212
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4885 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symn_1_6string) {
return PrimNumberToString(tmp4884)
}
__typedArg0 := tmp4884
return Call(__e, PrimFunc(symn_1_6string), __typedArg0)
})()

tmp4886 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V2211)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V2211
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()

tmp4887 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2212)
}
__typedArg0 := V2212
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4888 := Call(__e, PrimFunc(symshen_4reader_1error_1message), V2210, tmp4886, tmp4887)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(tmp4885)
__typedS1, __typedOK1 := TypedString(tmp4888)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := tmp4885
__typedArg1 := tmp4888
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("partial function shen.reader-error-message"))
}
__typedArg0 := MakeString("partial function shen.reader-error-message")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}


}, 3)

tmp4895 := Call(__e, ns2_1set, symshen_4reader_1error_1message, tmp4883)


_ = tmp4895

tmp4896 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dit_d)
}
__typedArg0 := symshen_4_dit_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp4897 := Call(__e, ns2_1set, symit, tmp4896)


_ = tmp4897

tmp4898 := MakeNative(func(__e *ControlFlow) {
V2213 := __e.Get(1)
_ = V2213
tmp4899 := MakeNative(func(__e *ControlFlow) {
W2214 := __e.Get(1)
_ = W2214
tmp4900 := MakeNative(func(__e *ControlFlow) {
W2215 := __e.Get(1)
_ = W2215
tmp4901 := MakeNative(func(__e *ControlFlow) {
W2216 := __e.Get(1)
_ = W2216
tmp4902 := MakeNative(func(__e *ControlFlow) {
W2217 := __e.Get(1)
_ = W2217
__e.TailApply(PrimFunc(symreverse), W2216)
return
}, 1)

tmp4903 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symclose) {
return PrimCloseStream(W2214)
}
__typedArg0 := W2214
return Call(__e, PrimFunc(symclose), __typedArg0)
})()

__e.TailApply(tmp4902, tmp4903)
return


}, 1)

tmp4904 := Call(__e, PrimFunc(symshen_4read_1file_1as_1bytelist_1help), W2214, W2215, Nil)


__e.TailApply(tmp4901, tmp4904)
return


}, 1)

tmp4905 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symread_1byte) {
return PrimReadByte(W2214)
}
__typedArg0 := W2214
return Call(__e, PrimFunc(symread_1byte), __typedArg0)
})()

__e.TailApply(tmp4900, tmp4905)
return


}, 1)

tmp4906 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symopen) {
return PrimOpenStream(V2213, symin)
}
__typedArg0 := V2213
__typedArg1 := symin
return Call(__e, PrimFunc(symopen), __typedArg0, __typedArg1)
})()

__e.TailApply(tmp4899, tmp4906)
return


}, 1)

tmp4907 := Call(__e, ns2_1set, symread_1file_1as_1bytelist, tmp4898)


_ = tmp4907

tmp4908 := MakeNative(func(__e *ControlFlow) {
V2218 := __e.Get(1)
_ = V2218
V2219 := __e.Get(2)
_ = V2219
V2220 := __e.Get(3)
_ = V2220
tmp4912 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeNumber(-1), V2219)
}
__typedArg0 := MakeNumber(-1)
__typedArg1 := V2219
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4912 {
__e.Return(V2220)
return
} else {
tmp4909 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symread_1byte) {
return PrimReadByte(V2218)
}
__typedArg0 := V2218
return Call(__e, PrimFunc(symread_1byte), __typedArg0)
})()

tmp4910 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V2219, V2220)
}
__typedArg0 := V2219
__typedArg1 := V2220
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4read_1file_1as_1bytelist_1help), V2218, tmp4909, tmp4910)
return


}


}, 3)

tmp4913 := Call(__e, ns2_1set, symshen_4read_1file_1as_1bytelist_1help, tmp4908)


_ = tmp4913

tmp4914 := MakeNative(func(__e *ControlFlow) {
V2221 := __e.Get(1)
_ = V2221
tmp4915 := MakeNative(func(__e *ControlFlow) {
W2222 := __e.Get(1)
_ = W2222
tmp4916 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symread_1byte) {
return PrimReadByte(W2222)
}
__typedArg0 := W2222
return Call(__e, PrimFunc(symread_1byte), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4rfas_1h), W2222, tmp4916, MakeString(""))
return


}, 1)

tmp4917 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symopen) {
return PrimOpenStream(V2221, symin)
}
__typedArg0 := V2221
__typedArg1 := symin
return Call(__e, PrimFunc(symopen), __typedArg0, __typedArg1)
})()

__e.TailApply(tmp4915, tmp4917)
return


}, 1)

tmp4918 := Call(__e, ns2_1set, symread_1file_1as_1string, tmp4914)


_ = tmp4918

tmp4919 := MakeNative(func(__e *ControlFlow) {
V2223 := __e.Get(1)
_ = V2223
V2224 := __e.Get(2)
_ = V2224
V2225 := __e.Get(3)
_ = V2225
tmp4925 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeNumber(-1), V2224)
}
__typedArg0 := MakeNumber(-1)
__typedArg1 := V2224
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4925 {
tmp4920 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symclose) {
return PrimCloseStream(V2223)
}
__typedArg0 := V2223
return Call(__e, PrimFunc(symclose), __typedArg0)
})()

_ = tmp4920

__e.Return(V2225)
return


} else {
tmp4921 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symread_1byte) {
return PrimReadByte(V2223)
}
__typedArg0 := V2223
return Call(__e, PrimFunc(symread_1byte), __typedArg0)
})()

tmp4922 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symn_1_6string) {
return PrimNumberToString(V2224)
}
__typedArg0 := V2224
return Call(__e, PrimFunc(symn_1_6string), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4rfas_1h), V2223, tmp4921, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(V2225)
__typedS1, __typedOK1 := TypedString(tmp4922)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := V2225
__typedArg1 := tmp4922
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
return


}


}, 3)

tmp4926 := Call(__e, ns2_1set, symshen_4rfas_1h, tmp4919)


_ = tmp4926

tmp4927 := MakeNative(func(__e *ControlFlow) {
V2226 := __e.Get(1)
_ = V2226
tmp4928 := Call(__e, PrimFunc(symread), V2226)


__e.TailApply(PrimFunc(symeval_1kl), tmp4928)
return


}, 1)

tmp4929 := Call(__e, ns2_1set, syminput, tmp4927)


_ = tmp4929

tmp4930 := MakeNative(func(__e *ControlFlow) {
V2227 := __e.Get(1)
_ = V2227
V2228 := __e.Get(2)
_ = V2228
tmp4931 := MakeNative(func(__e *ControlFlow) {
W2229 := __e.Get(1)
_ = W2229
tmp4932 := MakeNative(func(__e *ControlFlow) {
W2230 := __e.Get(1)
_ = W2230
tmp4938 := Call(__e, PrimFunc(symshen_4typecheck), W2230, V2227)


tmp4939 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(False, tmp4938)
}
__typedArg0 := False
__typedArg1 := tmp4938
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4939 {
tmp4933 := Call(__e, PrimFunc(symshen_4app), V2227, MakeString("\n"), symshen_4r)


tmp4935 := Call(__e, PrimFunc(symshen_4app), W2230, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(" is not of type "))
__typedS1, __typedOK1 := TypedString(tmp4933)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(" is not of type ")
__typedArg1 := tmp4933
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4r)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("type error: "))
__typedS1, __typedOK1 := TypedString(tmp4935)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("type error: ")
__typedArg1 := tmp4935
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("type error: "))
__typedS1, __typedOK1 := TypedString(tmp4935)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("type error: ")
__typedArg1 := tmp4935
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


} else {
__e.TailApply(PrimFunc(symeval_1kl), W2230)
return
}


}, 1)

tmp4940 := Call(__e, PrimFunc(symread), V2228)


__e.TailApply(tmp4932, tmp4940)
return


}, 1)

tmp4941 := Call(__e, PrimFunc(symshen_4monotype), V2227)


__e.TailApply(tmp4931, tmp4941)
return


}, 2)

tmp4942 := Call(__e, ns2_1set, symshen_4input_1h_7, tmp4930)


_ = tmp4942

tmp4943 := MakeNative(func(__e *ControlFlow) {
V2231 := __e.Get(1)
_ = V2231
tmp4950 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2231)
}
__typedArg0 := V2231
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp4950 {
tmp4944 := MakeNative(func(__e *ControlFlow) {
Z2232 := __e.Get(1)
_ = Z2232
__e.TailApply(PrimFunc(symshen_4monotype), Z2232)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp4944, V2231)
return


} else {
tmp4948 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(V2231)
}
__typedArg0 := V2231
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

if True == tmp4948 {
tmp4945 := Call(__e, PrimFunc(symshen_4app), V2231, MakeString("\n"), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("input+ expects a monotype: not "))
__typedS1, __typedOK1 := TypedString(tmp4945)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("input+ expects a monotype: not ")
__typedArg1 := tmp4945
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("input+ expects a monotype: not "))
__typedS1, __typedOK1 := TypedString(tmp4945)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("input+ expects a monotype: not ")
__typedArg1 := tmp4945
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


} else {
__e.Return(V2231)
return
}


}


}, 1)

tmp4951 := Call(__e, ns2_1set, symshen_4monotype, tmp4943)


_ = tmp4951

tmp4952 := MakeNative(func(__e *ControlFlow) {
V2233 := __e.Get(1)
_ = V2233
tmp4953 := Call(__e, PrimFunc(symshen_4my_1read_1byte), V2233)


tmp4954 := MakeNative(func(__e *ControlFlow) {
Z2234 := __e.Get(1)
_ = Z2234
__e.TailApply(PrimFunc(symshen_4return_2), Z2234)
return
}, 1)

__e.TailApply(PrimFunc(symshen_4read_1loop), V2233, tmp4953, Nil, tmp4954)
return


}, 1)

tmp4955 := Call(__e, ns2_1set, symlineread, tmp4952)


_ = tmp4955

tmp4956 := MakeNative(func(__e *ControlFlow) {
V2235 := __e.Get(1)
_ = V2235
tmp4957 := MakeNative(func(__e *ControlFlow) {
W2236 := __e.Get(1)
_ = W2236
tmp4958 := MakeNative(func(__e *ControlFlow) {
W2237 := __e.Get(1)
_ = W2237
tmp4959 := MakeNative(func(__e *ControlFlow) {
W2239 := __e.Get(1)
_ = W2239
__e.Return(W2239)
return
}, 1)

tmp4960 := Call(__e, PrimFunc(symshen_4process_1sexprs), W2237)


__e.TailApply(tmp4959, tmp4960)
return


}, 1)

tmp4961 := MakeNative(func(__e *ControlFlow) {
Z2238 := __e.Get(1)
_ = Z2238
__e.TailApply(PrimFunc(symshen_4_5s_1exprs_6), Z2238)
return
}, 1)

tmp4962 := Call(__e, PrimFunc(symcompile), tmp4961, W2236)


__e.TailApply(tmp4958, tmp4962)
return


}, 1)

tmp4963 := Call(__e, PrimFunc(symshen_4str_1_6bytes), V2235)


__e.TailApply(tmp4957, tmp4963)
return


}, 1)

tmp4964 := Call(__e, ns2_1set, symread_1from_1string, tmp4956)


_ = tmp4964

tmp4965 := MakeNative(func(__e *ControlFlow) {
V2240 := __e.Get(1)
_ = V2240
tmp4966 := MakeNative(func(__e *ControlFlow) {
W2241 := __e.Get(1)
_ = W2241
tmp4967 := MakeNative(func(__e *ControlFlow) {
W2242 := __e.Get(1)
_ = W2242
__e.Return(W2242)
return
}, 1)

tmp4968 := MakeNative(func(__e *ControlFlow) {
Z2243 := __e.Get(1)
_ = Z2243
__e.TailApply(PrimFunc(symshen_4_5s_1exprs_6), Z2243)
return
}, 1)

tmp4969 := Call(__e, PrimFunc(symcompile), tmp4968, W2241)


__e.TailApply(tmp4967, tmp4969)
return


}, 1)

tmp4970 := Call(__e, PrimFunc(symshen_4str_1_6bytes), V2240)


__e.TailApply(tmp4966, tmp4970)
return


}, 1)

tmp4971 := Call(__e, ns2_1set, symread_1from_1string_1unprocessed, tmp4965)


_ = tmp4971

tmp4972 := MakeNative(func(__e *ControlFlow) {
V2244 := __e.Get(1)
_ = V2244
tmp4980 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString(""), V2244)
}
__typedArg0 := MakeString("")
__typedArg1 := V2244
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4980 {
__e.Return(Nil)
return
} else {
tmp4978 := Call(__e, PrimFunc(symshen_4_7string_2), V2244)


if True == tmp4978 {
tmp4973 := Call(__e, PrimFunc(symhdstr), V2244)


tmp4974 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_1_6n) {
return PrimStringToNumber(tmp4973)
}
__typedArg0 := tmp4973
return Call(__e, PrimFunc(symstring_1_6n), __typedArg0)
})()

tmp4976 := Call(__e, PrimFunc(symshen_4str_1_6bytes), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V2244)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V2244
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4974, tmp4976)
}
__typedArg0 := tmp4974
__typedArg1 := tmp4976
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("partial function shen.str->bytes"))
}
__typedArg0 := MakeString("partial function shen.str->bytes")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 1)

tmp4981 := Call(__e, ns2_1set, symshen_4str_1_6bytes, tmp4972)


_ = tmp4981

tmp4982 := MakeNative(func(__e *ControlFlow) {
V2245 := __e.Get(1)
_ = V2245
tmp4983 := Call(__e, PrimFunc(symshen_4my_1read_1byte), V2245)


tmp4984 := MakeNative(func(__e *ControlFlow) {
Z2246 := __e.Get(1)
_ = Z2246
__e.TailApply(PrimFunc(symshen_4whitespace_2), Z2246)
return
}, 1)

tmp4985 := Call(__e, PrimFunc(symshen_4read_1loop), V2245, tmp4983, Nil, tmp4984)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4985)
}
__typedArg0 := tmp4985
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return


}, 1)

tmp4986 := Call(__e, ns2_1set, symread, tmp4982)


_ = tmp4986

tmp4987 := MakeNative(func(__e *ControlFlow) {
V2247 := __e.Get(1)
_ = V2247
tmp4990 := Call(__e, PrimFunc(symshen_4char_1stinput_2), V2247)


if True == tmp4990 {
tmp4988 := Call(__e, PrimFunc(symshen_4read_1unit_1string), V2247)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_1_6n) {
return PrimStringToNumber(tmp4988)
}
__typedArg0 := tmp4988
return Call(__e, PrimFunc(symstring_1_6n), __typedArg0)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symread_1byte) {
return PrimReadByte(V2247)
}
__typedArg0 := V2247
return Call(__e, PrimFunc(symread_1byte), __typedArg0)
})())
return
}


}, 1)

tmp4991 := Call(__e, ns2_1set, symshen_4my_1read_1byte, tmp4987)


_ = tmp4991

tmp4992 := MakeNative(func(__e *ControlFlow) {
V2252 := __e.Get(1)
_ = V2252
V2253 := __e.Get(2)
_ = V2253
V2254 := __e.Get(3)
_ = V2254
V2255 := __e.Get(4)
_ = V2255
tmp5015 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeNumber(94), V2253)
}
__typedArg0 := MakeNumber(94)
__typedArg1 := V2253
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp5015 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("read aborted"))
}
__typedArg0 := MakeString("read aborted")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
} else {
tmp5013 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeNumber(-1), V2253)
}
__typedArg0 := MakeNumber(-1)
__typedArg1 := V2253
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp5013 {
tmp4995 := Call(__e, PrimFunc(symempty_2), V2254)


if True == tmp4995 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("error: empty stream"))
}
__typedArg0 := MakeString("error: empty stream")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
} else {
tmp4993 := MakeNative(func(__e *ControlFlow) {
Z2256 := __e.Get(1)
_ = Z2256
__e.TailApply(PrimFunc(symshen_4_5s_1exprs_6), Z2256)
return
}, 1)

__e.TailApply(PrimFunc(symcompile), tmp4993, V2254)
return


}


} else {
tmp5011 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeNumber(0), V2253)
}
__typedArg0 := MakeNumber(0)
__typedArg1 := V2253
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp5011 {
tmp4996 := Call(__e, PrimFunc(symshen_4my_1read_1byte), V2252)


__e.TailApply(PrimFunc(symshen_4read_1loop), V2252, tmp4996, V2254, V2255)
return


} else {
tmp5009 := Call(__e, V2255, V2253)


if True == tmp5009 {
tmp4997 := MakeNative(func(__e *ControlFlow) {
W2257 := __e.Get(1)
_ = W2257
tmp5003 := Call(__e, PrimFunc(symshen_4nothing_1doing_2), W2257)


if True == tmp5003 {
tmp4998 := Call(__e, PrimFunc(symshen_4my_1read_1byte), V2252)


tmp4999 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V2253, Nil)
}
__typedArg0 := V2253
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp5000 := Call(__e, PrimFunc(symappend), V2254, tmp4999)


__e.TailApply(PrimFunc(symshen_4read_1loop), V2252, tmp4998, tmp5000, V2255)
return


} else {
tmp5001 := Call(__e, PrimFunc(symshen_4record_1it), V2254)


_ = tmp5001

__e.Return(W2257)
return


}


}, 1)

tmp5004 := Call(__e, PrimFunc(symshen_4try_1parse), V2254)


__e.TailApply(tmp4997, tmp5004)
return


} else {
tmp5005 := Call(__e, PrimFunc(symshen_4my_1read_1byte), V2252)


tmp5006 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V2253, Nil)
}
__typedArg0 := V2253
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp5007 := Call(__e, PrimFunc(symappend), V2254, tmp5006)


__e.TailApply(PrimFunc(symshen_4read_1loop), V2252, tmp5005, tmp5007, V2255)
return


}


}


}


}


}, 4)

tmp5016 := Call(__e, ns2_1set, symshen_4read_1loop, tmp4992)


_ = tmp5016

tmp5017 := MakeNative(func(__e *ControlFlow) {
V2258 := __e.Get(1)
_ = V2258
tmp5018 := MakeNative(func(__e *ControlFlow) {
W2259 := __e.Get(1)
_ = W2259
tmp5020 := Call(__e, PrimFunc(symshen_4nothing_1doing_2), W2259)


if True == tmp5020 {
__e.Return(symshen_4i_1failed_b)
return
} else {
__e.TailApply(PrimFunc(symshen_4process_1sexprs), W2259)
return
}


}, 1)

tmp5021 := MakeNative(func(__e *ControlFlow) {
tmp5022 := MakeNative(func(__e *ControlFlow) {
Z2260 := __e.Get(1)
_ = Z2260
__e.TailApply(PrimFunc(symshen_4_5s_1exprs_6), Z2260)
return
}, 1)

__e.TailApply(PrimFunc(symcompile), tmp5022, V2258)
return


}, 0)

tmp5023 := MakeNative(func(__e *ControlFlow) {
Z2261 := __e.Get(1)
_ = Z2261
__e.Return(symshen_4i_1failed_b)
return
}, 1)

tmp5024 := Call(__e, try_1catch, tmp5021, tmp5023)


__e.TailApply(tmp5018, tmp5024)
return


}, 1)

tmp5025 := Call(__e, ns2_1set, symshen_4try_1parse, tmp5017)


_ = tmp5025

tmp5026 := MakeNative(func(__e *ControlFlow) {
V2264 := __e.Get(1)
_ = V2264
tmp5030 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4i_1failed_b, V2264)
}
__typedArg0 := symshen_4i_1failed_b
__typedArg1 := V2264
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp5030 {
__e.Return(True)
return
} else {
tmp5028 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V2264)
}
__typedArg0 := Nil
__typedArg1 := V2264
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp5028 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp5031 := Call(__e, ns2_1set, symshen_4nothing_1doing_2, tmp5026)


_ = tmp5031

tmp5032 := MakeNative(func(__e *ControlFlow) {
V2265 := __e.Get(1)
_ = V2265
tmp5033 := Call(__e, PrimFunc(symshen_4bytes_1_6string), V2265)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dit_d, tmp5033)
}
__typedArg0 := symshen_4_dit_d
__typedArg1 := tmp5033
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp5034 := Call(__e, ns2_1set, symshen_4record_1it, tmp5032)


_ = tmp5034

tmp5035 := MakeNative(func(__e *ControlFlow) {
V2266 := __e.Get(1)
_ = V2266
tmp5043 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V2266)
}
__typedArg0 := Nil
__typedArg1 := V2266
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp5043 {
__e.Return(MakeString(""))
return
} else {
tmp5041 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2266)
}
__typedArg0 := V2266
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp5041 {
tmp5036 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2266)
}
__typedArg0 := V2266
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5037 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symn_1_6string) {
return PrimNumberToString(tmp5036)
}
__typedArg0 := tmp5036
return Call(__e, PrimFunc(symn_1_6string), __typedArg0)
})()

tmp5038 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2266)
}
__typedArg0 := V2266
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5039 := Call(__e, PrimFunc(symshen_4bytes_1_6string), tmp5038)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(tmp5037)
__typedS1, __typedOK1 := TypedString(tmp5039)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := tmp5037
__typedArg1 := tmp5039
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("partial function shen.bytes->string"))
}
__typedArg0 := MakeString("partial function shen.bytes->string")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 1)

tmp5044 := Call(__e, ns2_1set, symshen_4bytes_1_6string, tmp5035)


_ = tmp5044

tmp5045 := MakeNative(func(__e *ControlFlow) {
V2267 := __e.Get(1)
_ = V2267
tmp5046 := MakeNative(func(__e *ControlFlow) {
W2268 := __e.Get(1)
_ = W2268
tmp5047 := MakeNative(func(__e *ControlFlow) {
W2269 := __e.Get(1)
_ = W2269
tmp5048 := MakeNative(func(__e *ControlFlow) {
W2270 := __e.Get(1)
_ = W2270
tmp5049 := MakeNative(func(__e *ControlFlow) {
Z2271 := __e.Get(1)
_ = Z2271
__e.TailApply(PrimFunc(symshen_4process_1applications), Z2271, W2270)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp5049, W2268)
return


}, 1)

tmp5050 := Call(__e, PrimFunc(symshen_4find_1types), W2268)


__e.TailApply(tmp5048, tmp5050)
return


}, 1)

tmp5051 := Call(__e, PrimFunc(symshen_4find_1arities), W2268)


__e.TailApply(tmp5047, tmp5051)
return


}, 1)

tmp5052 := Call(__e, PrimFunc(symshen_4unpackage_emacroexpand), V2267)


__e.TailApply(tmp5046, tmp5052)
return


}, 1)

tmp5053 := Call(__e, ns2_1set, symshen_4process_1sexprs, tmp5045)


_ = tmp5053

tmp5054 := MakeNative(func(__e *ControlFlow) {
V2272 := __e.Get(1)
_ = V2272
tmp5076 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2272)
}
__typedArg0 := V2272
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres5067 Obj

if True == tmp5076 {
tmp5074 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2272)
}
__typedArg0 := V2272
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5075 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp5074)
}
__typedArg0 := tmp5074
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres5069 Obj

if True == tmp5075 {
tmp5071 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2272)
}
__typedArg0 := V2272
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5072 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp5073 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp5071, tmp5072)
}
__typedArg0 := tmp5071
__typedArg1 := tmp5072
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres5070 Obj

if True == tmp5073 {
ifres5070 = True


} else {
ifres5070 = False


}

ifres5069 = ifres5070


} else {
ifres5069 = False


}

var ifres5068 Obj

if True == ifres5069 {
ifres5068 = True


} else {
ifres5068 = False


}

ifres5067 = ifres5068


} else {
ifres5067 = False


}

if True == ifres5067 {
tmp5055 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2272)
}
__typedArg0 := V2272
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5056 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp5055)
}
__typedArg0 := tmp5055
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5057 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2272)
}
__typedArg0 := V2272
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5058 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp5057)
}
__typedArg0 := tmp5057
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5059 := Call(__e, PrimFunc(symshen_4find_1types), tmp5058)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp5056, tmp5059)
}
__typedArg0 := tmp5056
__typedArg1 := tmp5059
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp5065 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2272)
}
__typedArg0 := V2272
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp5065 {
tmp5060 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2272)
}
__typedArg0 := V2272
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5061 := Call(__e, PrimFunc(symshen_4find_1types), tmp5060)


tmp5062 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2272)
}
__typedArg0 := V2272
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5063 := Call(__e, PrimFunc(symshen_4find_1types), tmp5062)


__e.TailApply(PrimFunc(symappend), tmp5061, tmp5063)
return


} else {
__e.Return(Nil)
return
}


}


}, 1)

tmp5077 := Call(__e, ns2_1set, symshen_4find_1types, tmp5054)


_ = tmp5077

tmp5078 := MakeNative(func(__e *ControlFlow) {
V2275 := __e.Get(1)
_ = V2275
tmp5127 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2275)
}
__typedArg0 := V2275
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres5108 Obj

if True == tmp5127 {
tmp5125 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2275)
}
__typedArg0 := V2275
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5126 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefine, tmp5125)
}
__typedArg0 := symdefine
__typedArg1 := tmp5125
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres5110 Obj

if True == tmp5126 {
tmp5123 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2275)
}
__typedArg0 := V2275
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5124 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp5123)
}
__typedArg0 := tmp5123
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres5112 Obj

if True == tmp5124 {
tmp5120 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2275)
}
__typedArg0 := V2275
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5121 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp5120)
}
__typedArg0 := tmp5120
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5122 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp5121)
}
__typedArg0 := tmp5121
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres5114 Obj

if True == tmp5122 {
tmp5116 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2275)
}
__typedArg0 := V2275
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5117 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp5116)
}
__typedArg0 := tmp5116
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5118 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp5117)
}
__typedArg0 := tmp5117
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5119 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_i, tmp5118)
}
__typedArg0 := sym_i
__typedArg1 := tmp5118
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres5115 Obj

if True == tmp5119 {
ifres5115 = True


} else {
ifres5115 = False


}

ifres5114 = ifres5115


} else {
ifres5114 = False


}

var ifres5113 Obj

if True == ifres5114 {
ifres5113 = True


} else {
ifres5113 = False


}

ifres5112 = ifres5113


} else {
ifres5112 = False


}

var ifres5111 Obj

if True == ifres5112 {
ifres5111 = True


} else {
ifres5111 = False


}

ifres5110 = ifres5111


} else {
ifres5110 = False


}

var ifres5109 Obj

if True == ifres5110 {
ifres5109 = True


} else {
ifres5109 = False


}

ifres5108 = ifres5109


} else {
ifres5108 = False


}

if True == ifres5108 {
tmp5079 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2275)
}
__typedArg0 := V2275
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5080 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp5079)
}
__typedArg0 := tmp5079
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5081 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2275)
}
__typedArg0 := V2275
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5082 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp5081)
}
__typedArg0 := tmp5081
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5083 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2275)
}
__typedArg0 := V2275
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5084 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp5083)
}
__typedArg0 := tmp5083
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5085 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp5084)
}
__typedArg0 := tmp5084
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5086 := Call(__e, PrimFunc(symshen_4find_1arity), tmp5082, MakeNumber(1), tmp5085)


__e.TailApply(PrimFunc(symshen_4store_1arity), tmp5080, tmp5086)
return


} else {
tmp5106 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2275)
}
__typedArg0 := V2275
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres5098 Obj

if True == tmp5106 {
tmp5104 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2275)
}
__typedArg0 := V2275
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5105 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefine, tmp5104)
}
__typedArg0 := symdefine
__typedArg1 := tmp5104
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres5100 Obj

if True == tmp5105 {
tmp5102 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2275)
}
__typedArg0 := V2275
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5103 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp5102)
}
__typedArg0 := tmp5102
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres5101 Obj

if True == tmp5103 {
ifres5101 = True


} else {
ifres5101 = False


}

ifres5100 = ifres5101


} else {
ifres5100 = False


}

var ifres5099 Obj

if True == ifres5100 {
ifres5099 = True


} else {
ifres5099 = False


}

ifres5098 = ifres5099


} else {
ifres5098 = False


}

if True == ifres5098 {
tmp5087 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2275)
}
__typedArg0 := V2275
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5088 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp5087)
}
__typedArg0 := tmp5087
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5089 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2275)
}
__typedArg0 := V2275
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5090 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp5089)
}
__typedArg0 := tmp5089
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5091 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2275)
}
__typedArg0 := V2275
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5092 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp5091)
}
__typedArg0 := tmp5091
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5093 := Call(__e, PrimFunc(symshen_4find_1arity), tmp5090, MakeNumber(0), tmp5092)


__e.TailApply(PrimFunc(symshen_4store_1arity), tmp5088, tmp5093)
return


} else {
tmp5096 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2275)
}
__typedArg0 := V2275
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp5096 {
tmp5094 := MakeNative(func(__e *ControlFlow) {
Z2276 := __e.Get(1)
_ = Z2276
__e.TailApply(PrimFunc(symshen_4find_1arities), Z2276)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp5094, V2275)
return


} else {
__e.Return(symshen_4skip)
return
}


}


}


}, 1)

tmp5128 := Call(__e, ns2_1set, symshen_4find_1arities, tmp5078)


_ = tmp5128

tmp5129 := MakeNative(func(__e *ControlFlow) {
V2277 := __e.Get(1)
_ = V2277
V2278 := __e.Get(2)
_ = V2278
tmp5130 := MakeNative(func(__e *ControlFlow) {
W2279 := __e.Get(1)
_ = W2279
tmp5141 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W2279, MakeNumber(-1))
}
__typedArg0 := W2279
__typedArg1 := MakeNumber(-1)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp5141 {
__e.TailApply(PrimFunc(symshen_4execute_1store_1arity), V2277, V2278)
return
} else {
tmp5139 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W2279, V2278)
}
__typedArg0 := W2279
__typedArg1 := V2278
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp5139 {
__e.Return(symshen_4skip)
return
} else {
tmp5137 := Call(__e, PrimFunc(symshen_4sysfunc_2), V2277)


if True == tmp5137 {
tmp5131 := Call(__e, PrimFunc(symshen_4app), V2277, MakeString(" is a system function\n"), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(tmp5131)
}
__typedArg0 := tmp5131
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


} else {
tmp5132 := Call(__e, PrimFunc(symshen_4app), V2277, MakeString(" may cause errors\n"), symshen_4a)


tmp5133 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("changing the arity of "))
__typedS1, __typedOK1 := TypedString(tmp5132)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("changing the arity of ")
__typedArg1 := tmp5132
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp5134 := Call(__e, PrimFunc(symstoutput))


tmp5135 := Call(__e, PrimFunc(sympr), tmp5133, tmp5134)


_ = tmp5135

__e.TailApply(PrimFunc(symshen_4execute_1store_1arity), V2277, V2278)
return


}


}


}


}, 1)

tmp5142 := Call(__e, PrimFunc(symarity), V2277)


__e.TailApply(tmp5130, tmp5142)
return


}, 2)

tmp5143 := Call(__e, ns2_1set, symshen_4store_1arity, tmp5129)


_ = tmp5143

tmp5144 := MakeNative(func(__e *ControlFlow) {
V2280 := __e.Get(1)
_ = V2280
V2281 := __e.Get(2)
_ = V2281
tmp5149 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeNumber(0), V2281)
}
__typedArg0 := MakeNumber(0)
__typedArg1 := V2281
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp5149 {
tmp5145 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symput), V2280, symarity, MakeNumber(0), tmp5145)
return


} else {
tmp5146 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp5147 := Call(__e, PrimFunc(symput), V2280, symarity, V2281, tmp5146)


_ = tmp5147

__e.TailApply(PrimFunc(symshen_4update_1lambdatable), V2280, V2281)
return


}


}, 2)

tmp5150 := Call(__e, ns2_1set, symshen_4execute_1store_1arity, tmp5144)


_ = tmp5150

tmp5151 := MakeNative(func(__e *ControlFlow) {
V2282 := __e.Get(1)
_ = V2282
V2283 := __e.Get(2)
_ = V2283
tmp5152 := MakeNative(func(__e *ControlFlow) {
W2284 := __e.Get(1)
_ = W2284
tmp5153 := MakeNative(func(__e *ControlFlow) {
W2285 := __e.Get(1)
_ = W2285
tmp5154 := MakeNative(func(__e *ControlFlow) {
W2286 := __e.Get(1)
_ = W2286
tmp5155 := MakeNative(func(__e *ControlFlow) {
W2287 := __e.Get(1)
_ = W2287
__e.Return(W2287)
return
}, 1)

tmp5156 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dlambdatable_d, W2286)
}
__typedArg0 := symshen_4_dlambdatable_d
__typedArg1 := W2286
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

__e.TailApply(tmp5155, tmp5156)
return


}, 1)

tmp5157 := Call(__e, PrimFunc(symshen_4assoc_1_6), V2282, W2285, W2284)


__e.TailApply(tmp5154, tmp5157)
return


}, 1)

tmp5158 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V2282, Nil)
}
__typedArg0 := V2282
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp5159 := Call(__e, PrimFunc(symshen_4lambda_1function), tmp5158, V2283)


tmp5160 := Call(__e, PrimFunc(symeval_1kl), tmp5159)


__e.TailApply(tmp5153, tmp5160)
return


}, 1)

tmp5161 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dlambdatable_d)
}
__typedArg0 := symshen_4_dlambdatable_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(tmp5152, tmp5161)
return


}, 2)

tmp5162 := Call(__e, ns2_1set, symshen_4update_1lambdatable, tmp5151)


_ = tmp5162

tmp5163 := MakeNative(func(__e *ControlFlow) {
V2290 := __e.Get(1)
_ = V2290
V2291 := __e.Get(2)
_ = V2291
tmp5181 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeNumber(0), V2291)
}
__typedArg0 := MakeNumber(0)
__typedArg1 := V2291
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp5181 {
__e.Return(symshen_4skip)
return
} else {
tmp5179 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeNumber(1), V2291)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := V2291
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp5179 {
tmp5164 := MakeNative(func(__e *ControlFlow) {
W2292 := __e.Get(1)
_ = W2292
tmp5165 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W2292, Nil)
}
__typedArg0 := W2292
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp5166 := Call(__e, PrimFunc(symappend), V2290, tmp5165)


tmp5167 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp5166, Nil)
}
__typedArg0 := tmp5166
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp5168 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W2292, tmp5167)
}
__typedArg0 := W2292
__typedArg1 := tmp5167
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp5168)
}
__typedArg0 := symlambda
__typedArg1 := tmp5168
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp5169 := Call(__e, PrimFunc(symgensym), symY)


__e.TailApply(tmp5164, tmp5169)
return


} else {
tmp5170 := MakeNative(func(__e *ControlFlow) {
W2293 := __e.Get(1)
_ = W2293
tmp5171 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W2293, Nil)
}
__typedArg0 := W2293
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp5172 := Call(__e, PrimFunc(symappend), V2290, tmp5171)


tmp5174 := Call(__e, PrimFunc(symshen_4lambda_1function), tmp5172, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(V2291)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := V2291
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})())


tmp5175 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp5174, Nil)
}
__typedArg0 := tmp5174
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp5176 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W2293, tmp5175)
}
__typedArg0 := W2293
__typedArg1 := tmp5175
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp5176)
}
__typedArg0 := symlambda
__typedArg1 := tmp5176
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp5177 := Call(__e, PrimFunc(symgensym), symY)


__e.TailApply(tmp5170, tmp5177)
return


}


}


}, 2)

tmp5182 := Call(__e, ns2_1set, symshen_4lambda_1function, tmp5163)


_ = tmp5182

tmp5183 := MakeNative(func(__e *ControlFlow) {
V2303 := __e.Get(1)
_ = V2303
V2304 := __e.Get(2)
_ = V2304
V2305 := __e.Get(3)
_ = V2305
tmp5206 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V2305)
}
__typedArg0 := Nil
__typedArg1 := V2305
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp5206 {
tmp5184 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V2303, V2304)
}
__typedArg0 := V2303
__typedArg1 := V2304
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp5184, Nil)
}
__typedArg0 := tmp5184
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp5204 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2305)
}
__typedArg0 := V2305
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres5195 Obj

if True == tmp5204 {
tmp5202 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2305)
}
__typedArg0 := V2305
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5203 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp5202)
}
__typedArg0 := tmp5202
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres5197 Obj

if True == tmp5203 {
tmp5199 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2305)
}
__typedArg0 := V2305
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5200 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp5199)
}
__typedArg0 := tmp5199
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5201 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V2303, tmp5200)
}
__typedArg0 := V2303
__typedArg1 := tmp5200
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres5198 Obj

if True == tmp5201 {
ifres5198 = True


} else {
ifres5198 = False


}

ifres5197 = ifres5198


} else {
ifres5197 = False


}

var ifres5196 Obj

if True == ifres5197 {
ifres5196 = True


} else {
ifres5196 = False


}

ifres5195 = ifres5196


} else {
ifres5195 = False


}

if True == ifres5195 {
tmp5185 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2305)
}
__typedArg0 := V2305
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5186 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp5185)
}
__typedArg0 := tmp5185
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5187 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp5186, V2304)
}
__typedArg0 := tmp5186
__typedArg1 := V2304
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp5188 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2305)
}
__typedArg0 := V2305
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp5187, tmp5188)
}
__typedArg0 := tmp5187
__typedArg1 := tmp5188
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp5193 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2305)
}
__typedArg0 := V2305
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp5193 {
tmp5189 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2305)
}
__typedArg0 := V2305
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5190 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2305)
}
__typedArg0 := V2305
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5191 := Call(__e, PrimFunc(symshen_4assoc_1_6), V2303, V2304, tmp5190)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp5189, tmp5191)
}
__typedArg0 := tmp5189
__typedArg1 := tmp5191
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.assoc->"))
}
__typedArg0 := MakeString("implementation error in shen.assoc->")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}


}, 3)

tmp5207 := Call(__e, ns2_1set, symshen_4assoc_1_6, tmp5183)


_ = tmp5207

tmp5208 := MakeNative(func(__e *ControlFlow) {
V2320 := __e.Get(1)
_ = V2320
V2321 := __e.Get(2)
_ = V2321
V2322 := __e.Get(3)
_ = V2322
tmp5255 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeNumber(0), V2321)
}
__typedArg0 := MakeNumber(0)
__typedArg1 := V2321
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres5248 Obj

if True == tmp5255 {
tmp5254 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2322)
}
__typedArg0 := V2322
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres5250 Obj

if True == tmp5254 {
tmp5252 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2322)
}
__typedArg0 := V2322
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5253 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp5252, sym_1_6)
}
__typedArg0 := tmp5252
__typedArg1 := sym_1_6
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres5251 Obj

if True == tmp5253 {
ifres5251 = True


} else {
ifres5251 = False


}

ifres5250 = ifres5251


} else {
ifres5250 = False


}

var ifres5249 Obj

if True == ifres5250 {
ifres5249 = True


} else {
ifres5249 = False


}

ifres5248 = ifres5249


} else {
ifres5248 = False


}

if True == ifres5248 {
__e.Return(MakeNumber(0))
return
} else {
tmp5246 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeNumber(0), V2321)
}
__typedArg0 := MakeNumber(0)
__typedArg1 := V2321
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres5239 Obj

if True == tmp5246 {
tmp5245 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2322)
}
__typedArg0 := V2322
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres5241 Obj

if True == tmp5245 {
tmp5243 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2322)
}
__typedArg0 := V2322
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5244 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp5243, sym_5_1)
}
__typedArg0 := tmp5243
__typedArg1 := sym_5_1
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres5242 Obj

if True == tmp5244 {
ifres5242 = True


} else {
ifres5242 = False


}

ifres5241 = ifres5242


} else {
ifres5241 = False


}

var ifres5240 Obj

if True == ifres5241 {
ifres5240 = True


} else {
ifres5240 = False


}

ifres5239 = ifres5240


} else {
ifres5239 = False


}

if True == ifres5239 {
__e.Return(MakeNumber(0))
return
} else {
tmp5237 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeNumber(0), V2321)
}
__typedArg0 := MakeNumber(0)
__typedArg1 := V2321
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres5234 Obj

if True == tmp5237 {
tmp5236 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2322)
}
__typedArg0 := V2322
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres5235 Obj

if True == tmp5236 {
ifres5235 = True


} else {
ifres5235 = False


}

ifres5234 = ifres5235


} else {
ifres5234 = False


}

if True == ifres5234 {
tmp5209 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2322)
}
__typedArg0 := V2322
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5210 := Call(__e, PrimFunc(symshen_4find_1arity), V2320, MakeNumber(0), tmp5209)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(1))
__typedN1, __typedOK1 := TypedFloat64(tmp5210)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp5210
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())
return


} else {
tmp5232 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeNumber(1), V2321)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := V2321
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres5225 Obj

if True == tmp5232 {
tmp5231 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2322)
}
__typedArg0 := V2322
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres5227 Obj

if True == tmp5231 {
tmp5229 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2322)
}
__typedArg0 := V2322
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5230 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_j, tmp5229)
}
__typedArg0 := sym_j
__typedArg1 := tmp5229
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres5228 Obj

if True == tmp5230 {
ifres5228 = True


} else {
ifres5228 = False


}

ifres5227 = ifres5228


} else {
ifres5227 = False


}

var ifres5226 Obj

if True == ifres5227 {
ifres5226 = True


} else {
ifres5226 = False


}

ifres5225 = ifres5226


} else {
ifres5225 = False


}

if True == ifres5225 {
tmp5211 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2322)
}
__typedArg0 := V2322
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4find_1arity), V2320, MakeNumber(0), tmp5211)
return


} else {
tmp5223 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeNumber(1), V2321)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := V2321
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres5220 Obj

if True == tmp5223 {
tmp5222 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2322)
}
__typedArg0 := V2322
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres5221 Obj

if True == tmp5222 {
ifres5221 = True


} else {
ifres5221 = False


}

ifres5220 = ifres5221


} else {
ifres5220 = False


}

if True == ifres5220 {
tmp5212 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2322)
}
__typedArg0 := V2322
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4find_1arity), V2320, MakeNumber(1), tmp5212)
return


} else {
tmp5218 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeNumber(1), V2321)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := V2321
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp5218 {
tmp5213 := Call(__e, PrimFunc(symshen_4app), V2320, MakeString(" definition: missing }\n"), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("syntax error in "))
__typedS1, __typedOK1 := TypedString(tmp5213)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("syntax error in ")
__typedArg1 := tmp5213
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("syntax error in "))
__typedS1, __typedOK1 := TypedString(tmp5213)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("syntax error in ")
__typedArg1 := tmp5213
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


} else {
tmp5215 := Call(__e, PrimFunc(symshen_4app), V2320, MakeString(" definition: missing -> or <-\n"), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("syntax error in "))
__typedS1, __typedOK1 := TypedString(tmp5215)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("syntax error in ")
__typedArg1 := tmp5215
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("syntax error in "))
__typedS1, __typedOK1 := TypedString(tmp5215)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("syntax error in ")
__typedArg1 := tmp5215
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}


}


}


}


}


}


}, 3)

tmp5256 := Call(__e, ns2_1set, symshen_4find_1arity, tmp5208)


_ = tmp5256

tmp5257 := MakeNative(func(__e *ControlFlow) {
V2323 := __e.Get(1)
_ = V2323
tmp5258 := MakeNative(func(__e *ControlFlow) {
W2324 := __e.Get(1)
_ = W2324
tmp5503 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2324)


if True == tmp5503 {
tmp5259 := MakeNative(func(__e *ControlFlow) {
W2335 := __e.Get(1)
_ = W2335
tmp5471 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2335)


if True == tmp5471 {
tmp5260 := MakeNative(func(__e *ControlFlow) {
W2346 := __e.Get(1)
_ = W2346
tmp5453 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2346)


if True == tmp5453 {
tmp5261 := MakeNative(func(__e *ControlFlow) {
W2352 := __e.Get(1)
_ = W2352
tmp5435 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2352)


if True == tmp5435 {
tmp5262 := MakeNative(func(__e *ControlFlow) {
W2358 := __e.Get(1)
_ = W2358
tmp5417 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2358)


if True == tmp5417 {
tmp5263 := MakeNative(func(__e *ControlFlow) {
W2364 := __e.Get(1)
_ = W2364
tmp5398 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2364)


if True == tmp5398 {
tmp5264 := MakeNative(func(__e *ControlFlow) {
W2370 := __e.Get(1)
_ = W2370
tmp5373 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2370)


if True == tmp5373 {
tmp5265 := MakeNative(func(__e *ControlFlow) {
W2378 := __e.Get(1)
_ = W2378
tmp5354 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2378)


if True == tmp5354 {
tmp5266 := MakeNative(func(__e *ControlFlow) {
W2384 := __e.Get(1)
_ = W2384
tmp5335 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2384)


if True == tmp5335 {
tmp5267 := MakeNative(func(__e *ControlFlow) {
W2390 := __e.Get(1)
_ = W2390
tmp5318 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2390)


if True == tmp5318 {
tmp5268 := MakeNative(func(__e *ControlFlow) {
W2396 := __e.Get(1)
_ = W2396
tmp5298 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2396)


if True == tmp5298 {
tmp5269 := MakeNative(func(__e *ControlFlow) {
W2403 := __e.Get(1)
_ = W2403
tmp5281 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2403)


if True == tmp5281 {
tmp5270 := MakeNative(func(__e *ControlFlow) {
W2409 := __e.Get(1)
_ = W2409
tmp5272 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2409)


if True == tmp5272 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2409)
return
}


}, 1)

tmp5273 := MakeNative(func(__e *ControlFlow) {
W2410 := __e.Get(1)
_ = W2410
tmp5277 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2410)


if True == tmp5277 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5274 := MakeNative(func(__e *ControlFlow) {
W2411 := __e.Get(1)
_ = W2411
__e.TailApply(PrimFunc(symshen_4comb), W2411, Nil)
return
}, 1)

tmp5275 := Call(__e, PrimFunc(symshen_4in_1_6), W2410)


__e.TailApply(tmp5274, tmp5275)
return


}


}, 1)

tmp5278 := Call(__e, PrimFunc(sym_5e_6), V2323)


tmp5279 := Call(__e, tmp5273, tmp5278)


__e.TailApply(tmp5270, tmp5279)
return


} else {
__e.Return(W2403)
return
}


}, 1)

tmp5282 := MakeNative(func(__e *ControlFlow) {
W2404 := __e.Get(1)
_ = W2404
tmp5294 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2404)


if True == tmp5294 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5283 := MakeNative(func(__e *ControlFlow) {
W2405 := __e.Get(1)
_ = W2405
tmp5284 := MakeNative(func(__e *ControlFlow) {
W2406 := __e.Get(1)
_ = W2406
tmp5290 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2406)


if True == tmp5290 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5285 := MakeNative(func(__e *ControlFlow) {
W2407 := __e.Get(1)
_ = W2407
tmp5286 := MakeNative(func(__e *ControlFlow) {
W2408 := __e.Get(1)
_ = W2408
__e.TailApply(PrimFunc(symshen_4comb), W2408, W2407)
return
}, 1)

tmp5287 := Call(__e, PrimFunc(symshen_4in_1_6), W2406)


__e.TailApply(tmp5286, tmp5287)
return


}, 1)

tmp5288 := Call(__e, PrimFunc(symshen_4_5_1out), W2406)


__e.TailApply(tmp5285, tmp5288)
return


}


}, 1)

tmp5291 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), W2405)


__e.TailApply(tmp5284, tmp5291)
return


}, 1)

tmp5292 := Call(__e, PrimFunc(symshen_4in_1_6), W2404)


__e.TailApply(tmp5283, tmp5292)
return


}


}, 1)

tmp5295 := Call(__e, PrimFunc(symshen_4_5whitespaces_6), V2323)


tmp5296 := Call(__e, tmp5282, tmp5295)


__e.TailApply(tmp5269, tmp5296)
return


} else {
__e.Return(W2396)
return
}


}, 1)

tmp5299 := MakeNative(func(__e *ControlFlow) {
W2397 := __e.Get(1)
_ = W2397
tmp5314 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2397)


if True == tmp5314 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5300 := MakeNative(func(__e *ControlFlow) {
W2398 := __e.Get(1)
_ = W2398
tmp5301 := MakeNative(func(__e *ControlFlow) {
W2399 := __e.Get(1)
_ = W2399
tmp5302 := MakeNative(func(__e *ControlFlow) {
W2400 := __e.Get(1)
_ = W2400
tmp5309 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2400)


if True == tmp5309 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5303 := MakeNative(func(__e *ControlFlow) {
W2401 := __e.Get(1)
_ = W2401
tmp5304 := MakeNative(func(__e *ControlFlow) {
W2402 := __e.Get(1)
_ = W2402
tmp5305 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W2398, W2401)
}
__typedArg0 := W2398
__typedArg1 := W2401
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4comb), W2402, tmp5305)
return


}, 1)

tmp5306 := Call(__e, PrimFunc(symshen_4in_1_6), W2400)


__e.TailApply(tmp5304, tmp5306)
return


}, 1)

tmp5307 := Call(__e, PrimFunc(symshen_4_5_1out), W2400)


__e.TailApply(tmp5303, tmp5307)
return


}


}, 1)

tmp5310 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), W2399)


__e.TailApply(tmp5302, tmp5310)
return


}, 1)

tmp5311 := Call(__e, PrimFunc(symshen_4in_1_6), W2397)


__e.TailApply(tmp5301, tmp5311)
return


}, 1)

tmp5312 := Call(__e, PrimFunc(symshen_4_5_1out), W2397)


__e.TailApply(tmp5300, tmp5312)
return


}


}, 1)

tmp5315 := Call(__e, PrimFunc(symshen_4_5atom_6), V2323)


tmp5316 := Call(__e, tmp5299, tmp5315)


__e.TailApply(tmp5268, tmp5316)
return


} else {
__e.Return(W2390)
return
}


}, 1)

tmp5319 := MakeNative(func(__e *ControlFlow) {
W2391 := __e.Get(1)
_ = W2391
tmp5331 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2391)


if True == tmp5331 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5320 := MakeNative(func(__e *ControlFlow) {
W2392 := __e.Get(1)
_ = W2392
tmp5321 := MakeNative(func(__e *ControlFlow) {
W2393 := __e.Get(1)
_ = W2393
tmp5327 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2393)


if True == tmp5327 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5322 := MakeNative(func(__e *ControlFlow) {
W2394 := __e.Get(1)
_ = W2394
tmp5323 := MakeNative(func(__e *ControlFlow) {
W2395 := __e.Get(1)
_ = W2395
__e.TailApply(PrimFunc(symshen_4comb), W2395, W2394)
return
}, 1)

tmp5324 := Call(__e, PrimFunc(symshen_4in_1_6), W2393)


__e.TailApply(tmp5323, tmp5324)
return


}, 1)

tmp5325 := Call(__e, PrimFunc(symshen_4_5_1out), W2393)


__e.TailApply(tmp5322, tmp5325)
return


}


}, 1)

tmp5328 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), W2392)


__e.TailApply(tmp5321, tmp5328)
return


}, 1)

tmp5329 := Call(__e, PrimFunc(symshen_4in_1_6), W2391)


__e.TailApply(tmp5320, tmp5329)
return


}


}, 1)

tmp5332 := Call(__e, PrimFunc(symshen_4_5comment_6), V2323)


tmp5333 := Call(__e, tmp5319, tmp5332)


__e.TailApply(tmp5267, tmp5333)
return


} else {
__e.Return(W2384)
return
}


}, 1)

tmp5336 := MakeNative(func(__e *ControlFlow) {
W2385 := __e.Get(1)
_ = W2385
tmp5350 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2385)


if True == tmp5350 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5337 := MakeNative(func(__e *ControlFlow) {
W2386 := __e.Get(1)
_ = W2386
tmp5338 := MakeNative(func(__e *ControlFlow) {
W2387 := __e.Get(1)
_ = W2387
tmp5346 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2387)


if True == tmp5346 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5339 := MakeNative(func(__e *ControlFlow) {
W2388 := __e.Get(1)
_ = W2388
tmp5340 := MakeNative(func(__e *ControlFlow) {
W2389 := __e.Get(1)
_ = W2389
tmp5341 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(","))
}
__typedArg0 := MakeString(",")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp5342 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp5341, W2388)
}
__typedArg0 := tmp5341
__typedArg1 := W2388
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4comb), W2389, tmp5342)
return


}, 1)

tmp5343 := Call(__e, PrimFunc(symshen_4in_1_6), W2387)


__e.TailApply(tmp5340, tmp5343)
return


}, 1)

tmp5344 := Call(__e, PrimFunc(symshen_4_5_1out), W2387)


__e.TailApply(tmp5339, tmp5344)
return


}


}, 1)

tmp5347 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), W2386)


__e.TailApply(tmp5338, tmp5347)
return


}, 1)

tmp5348 := Call(__e, PrimFunc(symshen_4in_1_6), W2385)


__e.TailApply(tmp5337, tmp5348)
return


}


}, 1)

tmp5351 := Call(__e, PrimFunc(symshen_4_5comma_6), V2323)


tmp5352 := Call(__e, tmp5336, tmp5351)


__e.TailApply(tmp5266, tmp5352)
return


} else {
__e.Return(W2378)
return
}


}, 1)

tmp5355 := MakeNative(func(__e *ControlFlow) {
W2379 := __e.Get(1)
_ = W2379
tmp5369 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2379)


if True == tmp5369 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5356 := MakeNative(func(__e *ControlFlow) {
W2380 := __e.Get(1)
_ = W2380
tmp5357 := MakeNative(func(__e *ControlFlow) {
W2381 := __e.Get(1)
_ = W2381
tmp5365 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2381)


if True == tmp5365 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5358 := MakeNative(func(__e *ControlFlow) {
W2382 := __e.Get(1)
_ = W2382
tmp5359 := MakeNative(func(__e *ControlFlow) {
W2383 := __e.Get(1)
_ = W2383
tmp5360 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp5361 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp5360, W2382)
}
__typedArg0 := tmp5360
__typedArg1 := W2382
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4comb), W2383, tmp5361)
return


}, 1)

tmp5362 := Call(__e, PrimFunc(symshen_4in_1_6), W2381)


__e.TailApply(tmp5359, tmp5362)
return


}, 1)

tmp5363 := Call(__e, PrimFunc(symshen_4_5_1out), W2381)


__e.TailApply(tmp5358, tmp5363)
return


}


}, 1)

tmp5366 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), W2380)


__e.TailApply(tmp5357, tmp5366)
return


}, 1)

tmp5367 := Call(__e, PrimFunc(symshen_4in_1_6), W2379)


__e.TailApply(tmp5356, tmp5367)
return


}


}, 1)

tmp5370 := Call(__e, PrimFunc(symshen_4_5colon_6), V2323)


tmp5371 := Call(__e, tmp5355, tmp5370)


__e.TailApply(tmp5265, tmp5371)
return


} else {
__e.Return(W2370)
return
}


}, 1)

tmp5374 := MakeNative(func(__e *ControlFlow) {
W2371 := __e.Get(1)
_ = W2371
tmp5394 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2371)


if True == tmp5394 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5375 := MakeNative(func(__e *ControlFlow) {
W2372 := __e.Get(1)
_ = W2372
tmp5376 := MakeNative(func(__e *ControlFlow) {
W2373 := __e.Get(1)
_ = W2373
tmp5390 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2373)


if True == tmp5390 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5377 := MakeNative(func(__e *ControlFlow) {
W2374 := __e.Get(1)
_ = W2374
tmp5378 := MakeNative(func(__e *ControlFlow) {
W2375 := __e.Get(1)
_ = W2375
tmp5386 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2375)


if True == tmp5386 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5379 := MakeNative(func(__e *ControlFlow) {
W2376 := __e.Get(1)
_ = W2376
tmp5380 := MakeNative(func(__e *ControlFlow) {
W2377 := __e.Get(1)
_ = W2377
tmp5381 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":="))
}
__typedArg0 := MakeString(":=")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp5382 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp5381, W2376)
}
__typedArg0 := tmp5381
__typedArg1 := W2376
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4comb), W2377, tmp5382)
return


}, 1)

tmp5383 := Call(__e, PrimFunc(symshen_4in_1_6), W2375)


__e.TailApply(tmp5380, tmp5383)
return


}, 1)

tmp5384 := Call(__e, PrimFunc(symshen_4_5_1out), W2375)


__e.TailApply(tmp5379, tmp5384)
return


}


}, 1)

tmp5387 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), W2374)


__e.TailApply(tmp5378, tmp5387)
return


}, 1)

tmp5388 := Call(__e, PrimFunc(symshen_4in_1_6), W2373)


__e.TailApply(tmp5377, tmp5388)
return


}


}, 1)

tmp5391 := Call(__e, PrimFunc(symshen_4_5equal_6), W2372)


__e.TailApply(tmp5376, tmp5391)
return


}, 1)

tmp5392 := Call(__e, PrimFunc(symshen_4in_1_6), W2371)


__e.TailApply(tmp5375, tmp5392)
return


}


}, 1)

tmp5395 := Call(__e, PrimFunc(symshen_4_5colon_6), V2323)


tmp5396 := Call(__e, tmp5374, tmp5395)


__e.TailApply(tmp5264, tmp5396)
return


} else {
__e.Return(W2364)
return
}


}, 1)

tmp5399 := MakeNative(func(__e *ControlFlow) {
W2365 := __e.Get(1)
_ = W2365
tmp5413 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2365)


if True == tmp5413 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5400 := MakeNative(func(__e *ControlFlow) {
W2366 := __e.Get(1)
_ = W2366
tmp5401 := MakeNative(func(__e *ControlFlow) {
W2367 := __e.Get(1)
_ = W2367
tmp5409 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2367)


if True == tmp5409 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5402 := MakeNative(func(__e *ControlFlow) {
W2368 := __e.Get(1)
_ = W2368
tmp5403 := MakeNative(func(__e *ControlFlow) {
W2369 := __e.Get(1)
_ = W2369
tmp5404 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(";"))
}
__typedArg0 := MakeString(";")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp5405 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp5404, W2368)
}
__typedArg0 := tmp5404
__typedArg1 := W2368
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4comb), W2369, tmp5405)
return


}, 1)

tmp5406 := Call(__e, PrimFunc(symshen_4in_1_6), W2367)


__e.TailApply(tmp5403, tmp5406)
return


}, 1)

tmp5407 := Call(__e, PrimFunc(symshen_4_5_1out), W2367)


__e.TailApply(tmp5402, tmp5407)
return


}


}, 1)

tmp5410 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), W2366)


__e.TailApply(tmp5401, tmp5410)
return


}, 1)

tmp5411 := Call(__e, PrimFunc(symshen_4in_1_6), W2365)


__e.TailApply(tmp5400, tmp5411)
return


}


}, 1)

tmp5414 := Call(__e, PrimFunc(symshen_4_5semicolon_6), V2323)


tmp5415 := Call(__e, tmp5399, tmp5414)


__e.TailApply(tmp5263, tmp5415)
return


} else {
__e.Return(W2358)
return
}


}, 1)

tmp5418 := MakeNative(func(__e *ControlFlow) {
W2359 := __e.Get(1)
_ = W2359
tmp5431 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2359)


if True == tmp5431 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5419 := MakeNative(func(__e *ControlFlow) {
W2360 := __e.Get(1)
_ = W2360
tmp5420 := MakeNative(func(__e *ControlFlow) {
W2361 := __e.Get(1)
_ = W2361
tmp5427 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2361)


if True == tmp5427 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5421 := MakeNative(func(__e *ControlFlow) {
W2362 := __e.Get(1)
_ = W2362
tmp5422 := MakeNative(func(__e *ControlFlow) {
W2363 := __e.Get(1)
_ = W2363
tmp5423 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symbar_b, W2362)
}
__typedArg0 := symbar_b
__typedArg1 := W2362
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4comb), W2363, tmp5423)
return


}, 1)

tmp5424 := Call(__e, PrimFunc(symshen_4in_1_6), W2361)


__e.TailApply(tmp5422, tmp5424)
return


}, 1)

tmp5425 := Call(__e, PrimFunc(symshen_4_5_1out), W2361)


__e.TailApply(tmp5421, tmp5425)
return


}


}, 1)

tmp5428 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), W2360)


__e.TailApply(tmp5420, tmp5428)
return


}, 1)

tmp5429 := Call(__e, PrimFunc(symshen_4in_1_6), W2359)


__e.TailApply(tmp5419, tmp5429)
return


}


}, 1)

tmp5432 := Call(__e, PrimFunc(symshen_4_5bar_6), V2323)


tmp5433 := Call(__e, tmp5418, tmp5432)


__e.TailApply(tmp5262, tmp5433)
return


} else {
__e.Return(W2352)
return
}


}, 1)

tmp5436 := MakeNative(func(__e *ControlFlow) {
W2353 := __e.Get(1)
_ = W2353
tmp5449 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2353)


if True == tmp5449 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5437 := MakeNative(func(__e *ControlFlow) {
W2354 := __e.Get(1)
_ = W2354
tmp5438 := MakeNative(func(__e *ControlFlow) {
W2355 := __e.Get(1)
_ = W2355
tmp5445 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2355)


if True == tmp5445 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5439 := MakeNative(func(__e *ControlFlow) {
W2356 := __e.Get(1)
_ = W2356
tmp5440 := MakeNative(func(__e *ControlFlow) {
W2357 := __e.Get(1)
_ = W2357
tmp5441 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_j, W2356)
}
__typedArg0 := sym_j
__typedArg1 := W2356
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4comb), W2357, tmp5441)
return


}, 1)

tmp5442 := Call(__e, PrimFunc(symshen_4in_1_6), W2355)


__e.TailApply(tmp5440, tmp5442)
return


}, 1)

tmp5443 := Call(__e, PrimFunc(symshen_4_5_1out), W2355)


__e.TailApply(tmp5439, tmp5443)
return


}


}, 1)

tmp5446 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), W2354)


__e.TailApply(tmp5438, tmp5446)
return


}, 1)

tmp5447 := Call(__e, PrimFunc(symshen_4in_1_6), W2353)


__e.TailApply(tmp5437, tmp5447)
return


}


}, 1)

tmp5450 := Call(__e, PrimFunc(symshen_4_5rcurly_6), V2323)


tmp5451 := Call(__e, tmp5436, tmp5450)


__e.TailApply(tmp5261, tmp5451)
return


} else {
__e.Return(W2346)
return
}


}, 1)

tmp5454 := MakeNative(func(__e *ControlFlow) {
W2347 := __e.Get(1)
_ = W2347
tmp5467 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2347)


if True == tmp5467 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5455 := MakeNative(func(__e *ControlFlow) {
W2348 := __e.Get(1)
_ = W2348
tmp5456 := MakeNative(func(__e *ControlFlow) {
W2349 := __e.Get(1)
_ = W2349
tmp5463 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2349)


if True == tmp5463 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5457 := MakeNative(func(__e *ControlFlow) {
W2350 := __e.Get(1)
_ = W2350
tmp5458 := MakeNative(func(__e *ControlFlow) {
W2351 := __e.Get(1)
_ = W2351
tmp5459 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_i, W2350)
}
__typedArg0 := sym_i
__typedArg1 := W2350
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4comb), W2351, tmp5459)
return


}, 1)

tmp5460 := Call(__e, PrimFunc(symshen_4in_1_6), W2349)


__e.TailApply(tmp5458, tmp5460)
return


}, 1)

tmp5461 := Call(__e, PrimFunc(symshen_4_5_1out), W2349)


__e.TailApply(tmp5457, tmp5461)
return


}


}, 1)

tmp5464 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), W2348)


__e.TailApply(tmp5456, tmp5464)
return


}, 1)

tmp5465 := Call(__e, PrimFunc(symshen_4in_1_6), W2347)


__e.TailApply(tmp5455, tmp5465)
return


}


}, 1)

tmp5468 := Call(__e, PrimFunc(symshen_4_5lcurly_6), V2323)


tmp5469 := Call(__e, tmp5454, tmp5468)


__e.TailApply(tmp5260, tmp5469)
return


} else {
__e.Return(W2335)
return
}


}, 1)

tmp5472 := MakeNative(func(__e *ControlFlow) {
W2336 := __e.Get(1)
_ = W2336
tmp5499 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2336)


if True == tmp5499 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5473 := MakeNative(func(__e *ControlFlow) {
W2337 := __e.Get(1)
_ = W2337
tmp5474 := MakeNative(func(__e *ControlFlow) {
W2338 := __e.Get(1)
_ = W2338
tmp5495 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2338)


if True == tmp5495 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5475 := MakeNative(func(__e *ControlFlow) {
W2339 := __e.Get(1)
_ = W2339
tmp5476 := MakeNative(func(__e *ControlFlow) {
W2340 := __e.Get(1)
_ = W2340
tmp5477 := MakeNative(func(__e *ControlFlow) {
W2341 := __e.Get(1)
_ = W2341
tmp5490 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2341)


if True == tmp5490 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5478 := MakeNative(func(__e *ControlFlow) {
W2342 := __e.Get(1)
_ = W2342
tmp5479 := MakeNative(func(__e *ControlFlow) {
W2343 := __e.Get(1)
_ = W2343
tmp5486 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2343)


if True == tmp5486 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5480 := MakeNative(func(__e *ControlFlow) {
W2344 := __e.Get(1)
_ = W2344
tmp5481 := MakeNative(func(__e *ControlFlow) {
W2345 := __e.Get(1)
_ = W2345
tmp5482 := Call(__e, PrimFunc(symshen_4add_1sexpr), W2339, W2344)


__e.TailApply(PrimFunc(symshen_4comb), W2345, tmp5482)
return


}, 1)

tmp5483 := Call(__e, PrimFunc(symshen_4in_1_6), W2343)


__e.TailApply(tmp5481, tmp5483)
return


}, 1)

tmp5484 := Call(__e, PrimFunc(symshen_4_5_1out), W2343)


__e.TailApply(tmp5480, tmp5484)
return


}


}, 1)

tmp5487 := Call(__e, PrimFunc(symshen_4_5s_1exprs2_6), W2342)


__e.TailApply(tmp5479, tmp5487)
return


}, 1)

tmp5488 := Call(__e, PrimFunc(symshen_4in_1_6), W2341)


__e.TailApply(tmp5478, tmp5488)
return


}


}, 1)

tmp5491 := Call(__e, PrimFunc(symshen_4_5rrb_6), W2340)


__e.TailApply(tmp5477, tmp5491)
return


}, 1)

tmp5492 := Call(__e, PrimFunc(symshen_4in_1_6), W2338)


__e.TailApply(tmp5476, tmp5492)
return


}, 1)

tmp5493 := Call(__e, PrimFunc(symshen_4_5_1out), W2338)


__e.TailApply(tmp5475, tmp5493)
return


}


}, 1)

tmp5496 := Call(__e, PrimFunc(symshen_4_5s_1exprs1_6), W2337)


__e.TailApply(tmp5474, tmp5496)
return


}, 1)

tmp5497 := Call(__e, PrimFunc(symshen_4in_1_6), W2336)


__e.TailApply(tmp5473, tmp5497)
return


}


}, 1)

tmp5500 := Call(__e, PrimFunc(symshen_4_5lrb_6), V2323)


tmp5501 := Call(__e, tmp5472, tmp5500)


__e.TailApply(tmp5259, tmp5501)
return


} else {
__e.Return(W2324)
return
}


}, 1)

tmp5504 := MakeNative(func(__e *ControlFlow) {
W2325 := __e.Get(1)
_ = W2325
tmp5532 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2325)


if True == tmp5532 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5505 := MakeNative(func(__e *ControlFlow) {
W2326 := __e.Get(1)
_ = W2326
tmp5506 := MakeNative(func(__e *ControlFlow) {
W2327 := __e.Get(1)
_ = W2327
tmp5528 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2327)


if True == tmp5528 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5507 := MakeNative(func(__e *ControlFlow) {
W2328 := __e.Get(1)
_ = W2328
tmp5508 := MakeNative(func(__e *ControlFlow) {
W2329 := __e.Get(1)
_ = W2329
tmp5509 := MakeNative(func(__e *ControlFlow) {
W2330 := __e.Get(1)
_ = W2330
tmp5523 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2330)


if True == tmp5523 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5510 := MakeNative(func(__e *ControlFlow) {
W2331 := __e.Get(1)
_ = W2331
tmp5511 := MakeNative(func(__e *ControlFlow) {
W2332 := __e.Get(1)
_ = W2332
tmp5519 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2332)


if True == tmp5519 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5512 := MakeNative(func(__e *ControlFlow) {
W2333 := __e.Get(1)
_ = W2333
tmp5513 := MakeNative(func(__e *ControlFlow) {
W2334 := __e.Get(1)
_ = W2334
tmp5514 := Call(__e, PrimFunc(symshen_4cons_1form), W2328)


tmp5515 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp5514, W2333)
}
__typedArg0 := tmp5514
__typedArg1 := W2333
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4comb), W2334, tmp5515)
return


}, 1)

tmp5516 := Call(__e, PrimFunc(symshen_4in_1_6), W2332)


__e.TailApply(tmp5513, tmp5516)
return


}, 1)

tmp5517 := Call(__e, PrimFunc(symshen_4_5_1out), W2332)


__e.TailApply(tmp5512, tmp5517)
return


}


}, 1)

tmp5520 := Call(__e, PrimFunc(symshen_4_5s_1exprs2_6), W2331)


__e.TailApply(tmp5511, tmp5520)
return


}, 1)

tmp5521 := Call(__e, PrimFunc(symshen_4in_1_6), W2330)


__e.TailApply(tmp5510, tmp5521)
return


}


}, 1)

tmp5524 := Call(__e, PrimFunc(symshen_4_5rsb_6), W2329)


__e.TailApply(tmp5509, tmp5524)
return


}, 1)

tmp5525 := Call(__e, PrimFunc(symshen_4in_1_6), W2327)


__e.TailApply(tmp5508, tmp5525)
return


}, 1)

tmp5526 := Call(__e, PrimFunc(symshen_4_5_1out), W2327)


__e.TailApply(tmp5507, tmp5526)
return


}


}, 1)

tmp5529 := Call(__e, PrimFunc(symshen_4_5s_1exprs1_6), W2326)


__e.TailApply(tmp5506, tmp5529)
return


}, 1)

tmp5530 := Call(__e, PrimFunc(symshen_4in_1_6), W2325)


__e.TailApply(tmp5505, tmp5530)
return


}


}, 1)

tmp5533 := Call(__e, PrimFunc(symshen_4_5lsb_6), V2323)


tmp5534 := Call(__e, tmp5504, tmp5533)


__e.TailApply(tmp5258, tmp5534)
return


}, 1)

tmp5535 := Call(__e, ns2_1set, symshen_4_5s_1exprs_6, tmp5257)


_ = tmp5535

tmp5536 := MakeNative(func(__e *ControlFlow) {
V2412 := __e.Get(1)
_ = V2412
V2413 := __e.Get(2)
_ = V2413
tmp5554 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2412)
}
__typedArg0 := V2412
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres5541 Obj

if True == tmp5554 {
tmp5552 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2412)
}
__typedArg0 := V2412
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5553 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_3, tmp5552)
}
__typedArg0 := sym_3
__typedArg1 := tmp5552
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres5543 Obj

if True == tmp5553 {
tmp5550 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2412)
}
__typedArg0 := V2412
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5551 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp5550)
}
__typedArg0 := tmp5550
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres5545 Obj

if True == tmp5551 {
tmp5547 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2412)
}
__typedArg0 := V2412
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5548 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp5547)
}
__typedArg0 := tmp5547
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5549 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp5548)
}
__typedArg0 := Nil
__typedArg1 := tmp5548
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres5546 Obj

if True == tmp5549 {
ifres5546 = True


} else {
ifres5546 = False


}

ifres5545 = ifres5546


} else {
ifres5545 = False


}

var ifres5544 Obj

if True == ifres5545 {
ifres5544 = True


} else {
ifres5544 = False


}

ifres5543 = ifres5544


} else {
ifres5543 = False


}

var ifres5542 Obj

if True == ifres5543 {
ifres5542 = True


} else {
ifres5542 = False


}

ifres5541 = ifres5542


} else {
ifres5541 = False


}

if True == ifres5541 {
tmp5537 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2412)
}
__typedArg0 := V2412
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5538 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp5537)
}
__typedArg0 := tmp5537
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5539 := Call(__e, PrimFunc(symexplode), tmp5538)


__e.TailApply(PrimFunc(symappend), tmp5539, V2413)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V2412, V2413)
}
__typedArg0 := V2412
__typedArg1 := V2413
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return
}


}, 2)

tmp5555 := Call(__e, ns2_1set, symshen_4add_1sexpr, tmp5536)


_ = tmp5555

tmp5556 := MakeNative(func(__e *ControlFlow) {
V2414 := __e.Get(1)
_ = V2414
tmp5557 := MakeNative(func(__e *ControlFlow) {
W2415 := __e.Get(1)
_ = W2415
tmp5559 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2415)


if True == tmp5559 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2415)
return
}


}, 1)

tmp5565 := Call(__e, PrimFunc(symshen_4hds_a_2), V2414, MakeNumber(91))


var ifres5560 Obj

if True == tmp5565 {
tmp5561 := MakeNative(func(__e *ControlFlow) {
W2416 := __e.Get(1)
_ = W2416
__e.TailApply(PrimFunc(symshen_4comb), W2416, symshen_4skip)
return
}, 1)

tmp5562 := Call(__e, PrimFunc(symtail), V2414)


tmp5563 := Call(__e, tmp5561, tmp5562)


ifres5560 = tmp5563


} else {
tmp5564 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5560 = tmp5564


}

__e.TailApply(tmp5557, ifres5560)
return


}, 1)

tmp5566 := Call(__e, ns2_1set, symshen_4_5lsb_6, tmp5556)


_ = tmp5566

tmp5567 := MakeNative(func(__e *ControlFlow) {
V2417 := __e.Get(1)
_ = V2417
tmp5568 := MakeNative(func(__e *ControlFlow) {
W2418 := __e.Get(1)
_ = W2418
tmp5570 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2418)


if True == tmp5570 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2418)
return
}


}, 1)

tmp5576 := Call(__e, PrimFunc(symshen_4hds_a_2), V2417, MakeNumber(93))


var ifres5571 Obj

if True == tmp5576 {
tmp5572 := MakeNative(func(__e *ControlFlow) {
W2419 := __e.Get(1)
_ = W2419
__e.TailApply(PrimFunc(symshen_4comb), W2419, symshen_4skip)
return
}, 1)

tmp5573 := Call(__e, PrimFunc(symtail), V2417)


tmp5574 := Call(__e, tmp5572, tmp5573)


ifres5571 = tmp5574


} else {
tmp5575 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5571 = tmp5575


}

__e.TailApply(tmp5568, ifres5571)
return


}, 1)

tmp5577 := Call(__e, ns2_1set, symshen_4_5rsb_6, tmp5567)


_ = tmp5577

tmp5578 := MakeNative(func(__e *ControlFlow) {
V2420 := __e.Get(1)
_ = V2420
tmp5579 := MakeNative(func(__e *ControlFlow) {
W2421 := __e.Get(1)
_ = W2421
tmp5581 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2421)


if True == tmp5581 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2421)
return
}


}, 1)

tmp5582 := MakeNative(func(__e *ControlFlow) {
W2422 := __e.Get(1)
_ = W2422
tmp5588 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2422)


if True == tmp5588 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5583 := MakeNative(func(__e *ControlFlow) {
W2423 := __e.Get(1)
_ = W2423
tmp5584 := MakeNative(func(__e *ControlFlow) {
W2424 := __e.Get(1)
_ = W2424
__e.TailApply(PrimFunc(symshen_4comb), W2424, W2423)
return
}, 1)

tmp5585 := Call(__e, PrimFunc(symshen_4in_1_6), W2422)


__e.TailApply(tmp5584, tmp5585)
return


}, 1)

tmp5586 := Call(__e, PrimFunc(symshen_4_5_1out), W2422)


__e.TailApply(tmp5583, tmp5586)
return


}


}, 1)

tmp5589 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), V2420)


tmp5590 := Call(__e, tmp5582, tmp5589)


__e.TailApply(tmp5579, tmp5590)
return


}, 1)

tmp5591 := Call(__e, ns2_1set, symshen_4_5s_1exprs1_6, tmp5578)


_ = tmp5591

tmp5592 := MakeNative(func(__e *ControlFlow) {
V2425 := __e.Get(1)
_ = V2425
tmp5593 := MakeNative(func(__e *ControlFlow) {
W2426 := __e.Get(1)
_ = W2426
tmp5595 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2426)


if True == tmp5595 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2426)
return
}


}, 1)

tmp5596 := MakeNative(func(__e *ControlFlow) {
W2427 := __e.Get(1)
_ = W2427
tmp5602 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2427)


if True == tmp5602 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5597 := MakeNative(func(__e *ControlFlow) {
W2428 := __e.Get(1)
_ = W2428
tmp5598 := MakeNative(func(__e *ControlFlow) {
W2429 := __e.Get(1)
_ = W2429
__e.TailApply(PrimFunc(symshen_4comb), W2429, W2428)
return
}, 1)

tmp5599 := Call(__e, PrimFunc(symshen_4in_1_6), W2427)


__e.TailApply(tmp5598, tmp5599)
return


}, 1)

tmp5600 := Call(__e, PrimFunc(symshen_4_5_1out), W2427)


__e.TailApply(tmp5597, tmp5600)
return


}


}, 1)

tmp5603 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), V2425)


tmp5604 := Call(__e, tmp5596, tmp5603)


__e.TailApply(tmp5593, tmp5604)
return


}, 1)

tmp5605 := Call(__e, ns2_1set, symshen_4_5s_1exprs2_6, tmp5592)


_ = tmp5605

tmp5606 := MakeNative(func(__e *ControlFlow) {
V2431 := __e.Get(1)
_ = V2431
tmp5663 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V2431)
}
__typedArg0 := Nil
__typedArg1 := V2431
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp5663 {
__e.Return(Nil)
return
} else {
tmp5661 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2431)
}
__typedArg0 := V2431
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres5641 Obj

if True == tmp5661 {
tmp5659 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2431)
}
__typedArg0 := V2431
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5660 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp5659)
}
__typedArg0 := tmp5659
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres5643 Obj

if True == tmp5660 {
tmp5656 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2431)
}
__typedArg0 := V2431
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5657 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp5656)
}
__typedArg0 := tmp5656
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5658 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp5657)
}
__typedArg0 := tmp5657
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres5645 Obj

if True == tmp5658 {
tmp5652 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2431)
}
__typedArg0 := V2431
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5653 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp5652)
}
__typedArg0 := tmp5652
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5654 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp5653)
}
__typedArg0 := tmp5653
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5655 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp5654)
}
__typedArg0 := Nil
__typedArg1 := tmp5654
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres5647 Obj

if True == tmp5655 {
tmp5649 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2431)
}
__typedArg0 := V2431
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5650 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp5649)
}
__typedArg0 := tmp5649
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5651 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp5650, symbar_b)
}
__typedArg0 := tmp5650
__typedArg1 := symbar_b
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres5648 Obj

if True == tmp5651 {
ifres5648 = True


} else {
ifres5648 = False


}

ifres5647 = ifres5648


} else {
ifres5647 = False


}

var ifres5646 Obj

if True == ifres5647 {
ifres5646 = True


} else {
ifres5646 = False


}

ifres5645 = ifres5646


} else {
ifres5645 = False


}

var ifres5644 Obj

if True == ifres5645 {
ifres5644 = True


} else {
ifres5644 = False


}

ifres5643 = ifres5644


} else {
ifres5643 = False


}

var ifres5642 Obj

if True == ifres5643 {
ifres5642 = True


} else {
ifres5642 = False


}

ifres5641 = ifres5642


} else {
ifres5641 = False


}

if True == ifres5641 {
tmp5607 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2431)
}
__typedArg0 := V2431
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5608 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2431)
}
__typedArg0 := V2431
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5609 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp5608)
}
__typedArg0 := tmp5608
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5610 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp5607, tmp5609)
}
__typedArg0 := tmp5607
__typedArg1 := tmp5609
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp5610)
}
__typedArg0 := symcons
__typedArg1 := tmp5610
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp5639 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2431)
}
__typedArg0 := V2431
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres5619 Obj

if True == tmp5639 {
tmp5637 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2431)
}
__typedArg0 := V2431
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5638 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp5637)
}
__typedArg0 := tmp5637
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres5621 Obj

if True == tmp5638 {
tmp5634 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2431)
}
__typedArg0 := V2431
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5635 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp5634)
}
__typedArg0 := tmp5634
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5636 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp5635)
}
__typedArg0 := tmp5635
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres5623 Obj

if True == tmp5636 {
tmp5630 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2431)
}
__typedArg0 := V2431
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5631 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp5630)
}
__typedArg0 := tmp5630
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5632 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp5631)
}
__typedArg0 := tmp5631
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5633 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp5632)
}
__typedArg0 := tmp5632
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres5625 Obj

if True == tmp5633 {
tmp5627 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2431)
}
__typedArg0 := V2431
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5628 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp5627)
}
__typedArg0 := tmp5627
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5629 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp5628, symbar_b)
}
__typedArg0 := tmp5628
__typedArg1 := symbar_b
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres5626 Obj

if True == tmp5629 {
ifres5626 = True


} else {
ifres5626 = False


}

ifres5625 = ifres5626


} else {
ifres5625 = False


}

var ifres5624 Obj

if True == ifres5625 {
ifres5624 = True


} else {
ifres5624 = False


}

ifres5623 = ifres5624


} else {
ifres5623 = False


}

var ifres5622 Obj

if True == ifres5623 {
ifres5622 = True


} else {
ifres5622 = False


}

ifres5621 = ifres5622


} else {
ifres5621 = False


}

var ifres5620 Obj

if True == ifres5621 {
ifres5620 = True


} else {
ifres5620 = False


}

ifres5619 = ifres5620


} else {
ifres5619 = False


}

if True == ifres5619 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("misapplication of |\n"))
}
__typedArg0 := MakeString("misapplication of |\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
} else {
tmp5617 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2431)
}
__typedArg0 := V2431
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp5617 {
tmp5611 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2431)
}
__typedArg0 := V2431
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5612 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2431)
}
__typedArg0 := V2431
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5613 := Call(__e, PrimFunc(symshen_4cons_1form), tmp5612)


tmp5614 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp5613, Nil)
}
__typedArg0 := tmp5613
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp5615 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp5611, tmp5614)
}
__typedArg0 := tmp5611
__typedArg1 := tmp5614
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp5615)
}
__typedArg0 := symcons
__typedArg1 := tmp5615
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("partial function shen.cons-form"))
}
__typedArg0 := MakeString("partial function shen.cons-form")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}


}


}, 1)

tmp5664 := Call(__e, ns2_1set, symshen_4cons_1form, tmp5606)


_ = tmp5664

tmp5665 := MakeNative(func(__e *ControlFlow) {
V2432 := __e.Get(1)
_ = V2432
tmp5666 := MakeNative(func(__e *ControlFlow) {
W2433 := __e.Get(1)
_ = W2433
tmp5668 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2433)


if True == tmp5668 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2433)
return
}


}, 1)

tmp5674 := Call(__e, PrimFunc(symshen_4hds_a_2), V2432, MakeNumber(40))


var ifres5669 Obj

if True == tmp5674 {
tmp5670 := MakeNative(func(__e *ControlFlow) {
W2434 := __e.Get(1)
_ = W2434
__e.TailApply(PrimFunc(symshen_4comb), W2434, symshen_4skip)
return
}, 1)

tmp5671 := Call(__e, PrimFunc(symtail), V2432)


tmp5672 := Call(__e, tmp5670, tmp5671)


ifres5669 = tmp5672


} else {
tmp5673 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5669 = tmp5673


}

__e.TailApply(tmp5666, ifres5669)
return


}, 1)

tmp5675 := Call(__e, ns2_1set, symshen_4_5lrb_6, tmp5665)


_ = tmp5675

tmp5676 := MakeNative(func(__e *ControlFlow) {
V2435 := __e.Get(1)
_ = V2435
tmp5677 := MakeNative(func(__e *ControlFlow) {
W2436 := __e.Get(1)
_ = W2436
tmp5679 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2436)


if True == tmp5679 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2436)
return
}


}, 1)

tmp5685 := Call(__e, PrimFunc(symshen_4hds_a_2), V2435, MakeNumber(41))


var ifres5680 Obj

if True == tmp5685 {
tmp5681 := MakeNative(func(__e *ControlFlow) {
W2437 := __e.Get(1)
_ = W2437
__e.TailApply(PrimFunc(symshen_4comb), W2437, symshen_4skip)
return
}, 1)

tmp5682 := Call(__e, PrimFunc(symtail), V2435)


tmp5683 := Call(__e, tmp5681, tmp5682)


ifres5680 = tmp5683


} else {
tmp5684 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5680 = tmp5684


}

__e.TailApply(tmp5677, ifres5680)
return


}, 1)

tmp5686 := Call(__e, ns2_1set, symshen_4_5rrb_6, tmp5676)


_ = tmp5686

tmp5687 := MakeNative(func(__e *ControlFlow) {
V2438 := __e.Get(1)
_ = V2438
tmp5688 := MakeNative(func(__e *ControlFlow) {
W2439 := __e.Get(1)
_ = W2439
tmp5690 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2439)


if True == tmp5690 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2439)
return
}


}, 1)

tmp5696 := Call(__e, PrimFunc(symshen_4hds_a_2), V2438, MakeNumber(123))


var ifres5691 Obj

if True == tmp5696 {
tmp5692 := MakeNative(func(__e *ControlFlow) {
W2440 := __e.Get(1)
_ = W2440
__e.TailApply(PrimFunc(symshen_4comb), W2440, symshen_4skip)
return
}, 1)

tmp5693 := Call(__e, PrimFunc(symtail), V2438)


tmp5694 := Call(__e, tmp5692, tmp5693)


ifres5691 = tmp5694


} else {
tmp5695 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5691 = tmp5695


}

__e.TailApply(tmp5688, ifres5691)
return


}, 1)

tmp5697 := Call(__e, ns2_1set, symshen_4_5lcurly_6, tmp5687)


_ = tmp5697

tmp5698 := MakeNative(func(__e *ControlFlow) {
V2441 := __e.Get(1)
_ = V2441
tmp5699 := MakeNative(func(__e *ControlFlow) {
W2442 := __e.Get(1)
_ = W2442
tmp5701 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2442)


if True == tmp5701 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2442)
return
}


}, 1)

tmp5707 := Call(__e, PrimFunc(symshen_4hds_a_2), V2441, MakeNumber(125))


var ifres5702 Obj

if True == tmp5707 {
tmp5703 := MakeNative(func(__e *ControlFlow) {
W2443 := __e.Get(1)
_ = W2443
__e.TailApply(PrimFunc(symshen_4comb), W2443, symshen_4skip)
return
}, 1)

tmp5704 := Call(__e, PrimFunc(symtail), V2441)


tmp5705 := Call(__e, tmp5703, tmp5704)


ifres5702 = tmp5705


} else {
tmp5706 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5702 = tmp5706


}

__e.TailApply(tmp5699, ifres5702)
return


}, 1)

tmp5708 := Call(__e, ns2_1set, symshen_4_5rcurly_6, tmp5698)


_ = tmp5708

tmp5709 := MakeNative(func(__e *ControlFlow) {
V2444 := __e.Get(1)
_ = V2444
tmp5710 := MakeNative(func(__e *ControlFlow) {
W2445 := __e.Get(1)
_ = W2445
tmp5712 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2445)


if True == tmp5712 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2445)
return
}


}, 1)

tmp5718 := Call(__e, PrimFunc(symshen_4hds_a_2), V2444, MakeNumber(124))


var ifres5713 Obj

if True == tmp5718 {
tmp5714 := MakeNative(func(__e *ControlFlow) {
W2446 := __e.Get(1)
_ = W2446
__e.TailApply(PrimFunc(symshen_4comb), W2446, symshen_4skip)
return
}, 1)

tmp5715 := Call(__e, PrimFunc(symtail), V2444)


tmp5716 := Call(__e, tmp5714, tmp5715)


ifres5713 = tmp5716


} else {
tmp5717 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5713 = tmp5717


}

__e.TailApply(tmp5710, ifres5713)
return


}, 1)

tmp5719 := Call(__e, ns2_1set, symshen_4_5bar_6, tmp5709)


_ = tmp5719

tmp5720 := MakeNative(func(__e *ControlFlow) {
V2447 := __e.Get(1)
_ = V2447
tmp5721 := MakeNative(func(__e *ControlFlow) {
W2448 := __e.Get(1)
_ = W2448
tmp5723 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2448)


if True == tmp5723 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2448)
return
}


}, 1)

tmp5729 := Call(__e, PrimFunc(symshen_4hds_a_2), V2447, MakeNumber(59))


var ifres5724 Obj

if True == tmp5729 {
tmp5725 := MakeNative(func(__e *ControlFlow) {
W2449 := __e.Get(1)
_ = W2449
__e.TailApply(PrimFunc(symshen_4comb), W2449, symshen_4skip)
return
}, 1)

tmp5726 := Call(__e, PrimFunc(symtail), V2447)


tmp5727 := Call(__e, tmp5725, tmp5726)


ifres5724 = tmp5727


} else {
tmp5728 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5724 = tmp5728


}

__e.TailApply(tmp5721, ifres5724)
return


}, 1)

tmp5730 := Call(__e, ns2_1set, symshen_4_5semicolon_6, tmp5720)


_ = tmp5730

tmp5731 := MakeNative(func(__e *ControlFlow) {
V2450 := __e.Get(1)
_ = V2450
tmp5732 := MakeNative(func(__e *ControlFlow) {
W2451 := __e.Get(1)
_ = W2451
tmp5734 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2451)


if True == tmp5734 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2451)
return
}


}, 1)

tmp5740 := Call(__e, PrimFunc(symshen_4hds_a_2), V2450, MakeNumber(58))


var ifres5735 Obj

if True == tmp5740 {
tmp5736 := MakeNative(func(__e *ControlFlow) {
W2452 := __e.Get(1)
_ = W2452
__e.TailApply(PrimFunc(symshen_4comb), W2452, symshen_4skip)
return
}, 1)

tmp5737 := Call(__e, PrimFunc(symtail), V2450)


tmp5738 := Call(__e, tmp5736, tmp5737)


ifres5735 = tmp5738


} else {
tmp5739 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5735 = tmp5739


}

__e.TailApply(tmp5732, ifres5735)
return


}, 1)

tmp5741 := Call(__e, ns2_1set, symshen_4_5colon_6, tmp5731)


_ = tmp5741

tmp5742 := MakeNative(func(__e *ControlFlow) {
V2453 := __e.Get(1)
_ = V2453
tmp5743 := MakeNative(func(__e *ControlFlow) {
W2454 := __e.Get(1)
_ = W2454
tmp5745 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2454)


if True == tmp5745 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2454)
return
}


}, 1)

tmp5751 := Call(__e, PrimFunc(symshen_4hds_a_2), V2453, MakeNumber(44))


var ifres5746 Obj

if True == tmp5751 {
tmp5747 := MakeNative(func(__e *ControlFlow) {
W2455 := __e.Get(1)
_ = W2455
__e.TailApply(PrimFunc(symshen_4comb), W2455, symshen_4skip)
return
}, 1)

tmp5748 := Call(__e, PrimFunc(symtail), V2453)


tmp5749 := Call(__e, tmp5747, tmp5748)


ifres5746 = tmp5749


} else {
tmp5750 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5746 = tmp5750


}

__e.TailApply(tmp5743, ifres5746)
return


}, 1)

tmp5752 := Call(__e, ns2_1set, symshen_4_5comma_6, tmp5742)


_ = tmp5752

tmp5753 := MakeNative(func(__e *ControlFlow) {
V2456 := __e.Get(1)
_ = V2456
tmp5754 := MakeNative(func(__e *ControlFlow) {
W2457 := __e.Get(1)
_ = W2457
tmp5756 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2457)


if True == tmp5756 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2457)
return
}


}, 1)

tmp5762 := Call(__e, PrimFunc(symshen_4hds_a_2), V2456, MakeNumber(61))


var ifres5757 Obj

if True == tmp5762 {
tmp5758 := MakeNative(func(__e *ControlFlow) {
W2458 := __e.Get(1)
_ = W2458
__e.TailApply(PrimFunc(symshen_4comb), W2458, symshen_4skip)
return
}, 1)

tmp5759 := Call(__e, PrimFunc(symtail), V2456)


tmp5760 := Call(__e, tmp5758, tmp5759)


ifres5757 = tmp5760


} else {
tmp5761 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5757 = tmp5761


}

__e.TailApply(tmp5754, ifres5757)
return


}, 1)

tmp5763 := Call(__e, ns2_1set, symshen_4_5equal_6, tmp5753)


_ = tmp5763

tmp5764 := MakeNative(func(__e *ControlFlow) {
V2459 := __e.Get(1)
_ = V2459
tmp5765 := MakeNative(func(__e *ControlFlow) {
W2460 := __e.Get(1)
_ = W2460
tmp5777 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2460)


if True == tmp5777 {
tmp5766 := MakeNative(func(__e *ControlFlow) {
W2463 := __e.Get(1)
_ = W2463
tmp5768 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2463)


if True == tmp5768 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2463)
return
}


}, 1)

tmp5769 := MakeNative(func(__e *ControlFlow) {
W2464 := __e.Get(1)
_ = W2464
tmp5773 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2464)


if True == tmp5773 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5770 := MakeNative(func(__e *ControlFlow) {
W2465 := __e.Get(1)
_ = W2465
__e.TailApply(PrimFunc(symshen_4comb), W2465, symshen_4skip)
return
}, 1)

tmp5771 := Call(__e, PrimFunc(symshen_4in_1_6), W2464)


__e.TailApply(tmp5770, tmp5771)
return


}


}, 1)

tmp5774 := Call(__e, PrimFunc(symshen_4_5multiline_6), V2459)


tmp5775 := Call(__e, tmp5769, tmp5774)


__e.TailApply(tmp5766, tmp5775)
return


} else {
__e.Return(W2460)
return
}


}, 1)

tmp5778 := MakeNative(func(__e *ControlFlow) {
W2461 := __e.Get(1)
_ = W2461
tmp5782 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2461)


if True == tmp5782 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5779 := MakeNative(func(__e *ControlFlow) {
W2462 := __e.Get(1)
_ = W2462
__e.TailApply(PrimFunc(symshen_4comb), W2462, symshen_4skip)
return
}, 1)

tmp5780 := Call(__e, PrimFunc(symshen_4in_1_6), W2461)


__e.TailApply(tmp5779, tmp5780)
return


}


}, 1)

tmp5783 := Call(__e, PrimFunc(symshen_4_5singleline_6), V2459)


tmp5784 := Call(__e, tmp5778, tmp5783)


__e.TailApply(tmp5765, tmp5784)
return


}, 1)

tmp5785 := Call(__e, ns2_1set, symshen_4_5comment_6, tmp5764)


_ = tmp5785

tmp5786 := MakeNative(func(__e *ControlFlow) {
V2466 := __e.Get(1)
_ = V2466
tmp5787 := MakeNative(func(__e *ControlFlow) {
W2467 := __e.Get(1)
_ = W2467
tmp5789 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2467)


if True == tmp5789 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2467)
return
}


}, 1)

tmp5790 := MakeNative(func(__e *ControlFlow) {
W2468 := __e.Get(1)
_ = W2468
tmp5812 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2468)


if True == tmp5812 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5791 := MakeNative(func(__e *ControlFlow) {
W2469 := __e.Get(1)
_ = W2469
tmp5792 := MakeNative(func(__e *ControlFlow) {
W2470 := __e.Get(1)
_ = W2470
tmp5808 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2470)


if True == tmp5808 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5793 := MakeNative(func(__e *ControlFlow) {
W2471 := __e.Get(1)
_ = W2471
tmp5794 := MakeNative(func(__e *ControlFlow) {
W2472 := __e.Get(1)
_ = W2472
tmp5804 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2472)


if True == tmp5804 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5795 := MakeNative(func(__e *ControlFlow) {
W2473 := __e.Get(1)
_ = W2473
tmp5796 := MakeNative(func(__e *ControlFlow) {
W2474 := __e.Get(1)
_ = W2474
tmp5800 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2474)


if True == tmp5800 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5797 := MakeNative(func(__e *ControlFlow) {
W2475 := __e.Get(1)
_ = W2475
__e.TailApply(PrimFunc(symshen_4comb), W2475, symshen_4skip)
return
}, 1)

tmp5798 := Call(__e, PrimFunc(symshen_4in_1_6), W2474)


__e.TailApply(tmp5797, tmp5798)
return


}


}, 1)

tmp5801 := Call(__e, PrimFunc(symshen_4_5returns_6), W2473)


__e.TailApply(tmp5796, tmp5801)
return


}, 1)

tmp5802 := Call(__e, PrimFunc(symshen_4in_1_6), W2472)


__e.TailApply(tmp5795, tmp5802)
return


}


}, 1)

tmp5805 := Call(__e, PrimFunc(symshen_4_5shortnatters_6), W2471)


__e.TailApply(tmp5794, tmp5805)
return


}, 1)

tmp5806 := Call(__e, PrimFunc(symshen_4in_1_6), W2470)


__e.TailApply(tmp5793, tmp5806)
return


}


}, 1)

tmp5809 := Call(__e, PrimFunc(symshen_4_5backslash_6), W2469)


__e.TailApply(tmp5792, tmp5809)
return


}, 1)

tmp5810 := Call(__e, PrimFunc(symshen_4in_1_6), W2468)


__e.TailApply(tmp5791, tmp5810)
return


}


}, 1)

tmp5813 := Call(__e, PrimFunc(symshen_4_5backslash_6), V2466)


tmp5814 := Call(__e, tmp5790, tmp5813)


__e.TailApply(tmp5787, tmp5814)
return


}, 1)

tmp5815 := Call(__e, ns2_1set, symshen_4_5singleline_6, tmp5786)


_ = tmp5815

tmp5816 := MakeNative(func(__e *ControlFlow) {
V2476 := __e.Get(1)
_ = V2476
tmp5817 := MakeNative(func(__e *ControlFlow) {
W2477 := __e.Get(1)
_ = W2477
tmp5819 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2477)


if True == tmp5819 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2477)
return
}


}, 1)

tmp5825 := Call(__e, PrimFunc(symshen_4hds_a_2), V2476, MakeNumber(92))


var ifres5820 Obj

if True == tmp5825 {
tmp5821 := MakeNative(func(__e *ControlFlow) {
W2478 := __e.Get(1)
_ = W2478
__e.TailApply(PrimFunc(symshen_4comb), W2478, symshen_4skip)
return
}, 1)

tmp5822 := Call(__e, PrimFunc(symtail), V2476)


tmp5823 := Call(__e, tmp5821, tmp5822)


ifres5820 = tmp5823


} else {
tmp5824 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5820 = tmp5824


}

__e.TailApply(tmp5817, ifres5820)
return


}, 1)

tmp5826 := Call(__e, ns2_1set, symshen_4_5backslash_6, tmp5816)


_ = tmp5826

tmp5827 := MakeNative(func(__e *ControlFlow) {
V2479 := __e.Get(1)
_ = V2479
tmp5828 := MakeNative(func(__e *ControlFlow) {
W2480 := __e.Get(1)
_ = W2480
tmp5840 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2480)


if True == tmp5840 {
tmp5829 := MakeNative(func(__e *ControlFlow) {
W2485 := __e.Get(1)
_ = W2485
tmp5831 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2485)


if True == tmp5831 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2485)
return
}


}, 1)

tmp5832 := MakeNative(func(__e *ControlFlow) {
W2486 := __e.Get(1)
_ = W2486
tmp5836 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2486)


if True == tmp5836 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5833 := MakeNative(func(__e *ControlFlow) {
W2487 := __e.Get(1)
_ = W2487
__e.TailApply(PrimFunc(symshen_4comb), W2487, symshen_4skip)
return
}, 1)

tmp5834 := Call(__e, PrimFunc(symshen_4in_1_6), W2486)


__e.TailApply(tmp5833, tmp5834)
return


}


}, 1)

tmp5837 := Call(__e, PrimFunc(sym_5e_6), V2479)


tmp5838 := Call(__e, tmp5832, tmp5837)


__e.TailApply(tmp5829, tmp5838)
return


} else {
__e.Return(W2480)
return
}


}, 1)

tmp5841 := MakeNative(func(__e *ControlFlow) {
W2481 := __e.Get(1)
_ = W2481
tmp5851 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2481)


if True == tmp5851 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5842 := MakeNative(func(__e *ControlFlow) {
W2482 := __e.Get(1)
_ = W2482
tmp5843 := MakeNative(func(__e *ControlFlow) {
W2483 := __e.Get(1)
_ = W2483
tmp5847 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2483)


if True == tmp5847 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5844 := MakeNative(func(__e *ControlFlow) {
W2484 := __e.Get(1)
_ = W2484
__e.TailApply(PrimFunc(symshen_4comb), W2484, symshen_4skip)
return
}, 1)

tmp5845 := Call(__e, PrimFunc(symshen_4in_1_6), W2483)


__e.TailApply(tmp5844, tmp5845)
return


}


}, 1)

tmp5848 := Call(__e, PrimFunc(symshen_4_5shortnatters_6), W2482)


__e.TailApply(tmp5843, tmp5848)
return


}, 1)

tmp5849 := Call(__e, PrimFunc(symshen_4in_1_6), W2481)


__e.TailApply(tmp5842, tmp5849)
return


}


}, 1)

tmp5852 := Call(__e, PrimFunc(symshen_4_5shortnatter_6), V2479)


tmp5853 := Call(__e, tmp5841, tmp5852)


__e.TailApply(tmp5828, tmp5853)
return


}, 1)

tmp5854 := Call(__e, ns2_1set, symshen_4_5shortnatters_6, tmp5827)


_ = tmp5854

tmp5855 := MakeNative(func(__e *ControlFlow) {
V2488 := __e.Get(1)
_ = V2488
tmp5856 := MakeNative(func(__e *ControlFlow) {
W2489 := __e.Get(1)
_ = W2489
tmp5858 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2489)


if True == tmp5858 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2489)
return
}


}, 1)

tmp5869 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2488)
}
__typedArg0 := V2488
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres5859 Obj

if True == tmp5869 {
tmp5860 := MakeNative(func(__e *ControlFlow) {
W2490 := __e.Get(1)
_ = W2490
tmp5861 := MakeNative(func(__e *ControlFlow) {
W2491 := __e.Get(1)
_ = W2491
tmp5863 := Call(__e, PrimFunc(symshen_4return_2), W2490)


if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp5863)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp5863
return Call(__e, PrimFunc(symnot), __typedArg0)
})() {
__e.TailApply(PrimFunc(symshen_4comb), W2491, symshen_4skip)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp5865 := Call(__e, PrimFunc(symtail), V2488)


__e.TailApply(tmp5861, tmp5865)
return


}, 1)

tmp5866 := Call(__e, PrimFunc(symhead), V2488)


tmp5867 := Call(__e, tmp5860, tmp5866)


ifres5859 = tmp5867


} else {
tmp5868 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5859 = tmp5868


}

__e.TailApply(tmp5856, ifres5859)
return


}, 1)

tmp5870 := Call(__e, ns2_1set, symshen_4_5shortnatter_6, tmp5855)


_ = tmp5870

tmp5871 := MakeNative(func(__e *ControlFlow) {
V2492 := __e.Get(1)
_ = V2492
tmp5872 := MakeNative(func(__e *ControlFlow) {
W2493 := __e.Get(1)
_ = W2493
tmp5884 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2493)


if True == tmp5884 {
tmp5873 := MakeNative(func(__e *ControlFlow) {
W2498 := __e.Get(1)
_ = W2498
tmp5875 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2498)


if True == tmp5875 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2498)
return
}


}, 1)

tmp5876 := MakeNative(func(__e *ControlFlow) {
W2499 := __e.Get(1)
_ = W2499
tmp5880 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2499)


if True == tmp5880 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5877 := MakeNative(func(__e *ControlFlow) {
W2500 := __e.Get(1)
_ = W2500
__e.TailApply(PrimFunc(symshen_4comb), W2500, symshen_4skip)
return
}, 1)

tmp5878 := Call(__e, PrimFunc(symshen_4in_1_6), W2499)


__e.TailApply(tmp5877, tmp5878)
return


}


}, 1)

tmp5881 := Call(__e, PrimFunc(symshen_4_5return_6), V2492)


tmp5882 := Call(__e, tmp5876, tmp5881)


__e.TailApply(tmp5873, tmp5882)
return


} else {
__e.Return(W2493)
return
}


}, 1)

tmp5885 := MakeNative(func(__e *ControlFlow) {
W2494 := __e.Get(1)
_ = W2494
tmp5895 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2494)


if True == tmp5895 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5886 := MakeNative(func(__e *ControlFlow) {
W2495 := __e.Get(1)
_ = W2495
tmp5887 := MakeNative(func(__e *ControlFlow) {
W2496 := __e.Get(1)
_ = W2496
tmp5891 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2496)


if True == tmp5891 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5888 := MakeNative(func(__e *ControlFlow) {
W2497 := __e.Get(1)
_ = W2497
__e.TailApply(PrimFunc(symshen_4comb), W2497, symshen_4skip)
return
}, 1)

tmp5889 := Call(__e, PrimFunc(symshen_4in_1_6), W2496)


__e.TailApply(tmp5888, tmp5889)
return


}


}, 1)

tmp5892 := Call(__e, PrimFunc(symshen_4_5returns_6), W2495)


__e.TailApply(tmp5887, tmp5892)
return


}, 1)

tmp5893 := Call(__e, PrimFunc(symshen_4in_1_6), W2494)


__e.TailApply(tmp5886, tmp5893)
return


}


}, 1)

tmp5896 := Call(__e, PrimFunc(symshen_4_5return_6), V2492)


tmp5897 := Call(__e, tmp5885, tmp5896)


__e.TailApply(tmp5872, tmp5897)
return


}, 1)

tmp5898 := Call(__e, ns2_1set, symshen_4_5returns_6, tmp5871)


_ = tmp5898

tmp5899 := MakeNative(func(__e *ControlFlow) {
V2501 := __e.Get(1)
_ = V2501
tmp5900 := MakeNative(func(__e *ControlFlow) {
W2502 := __e.Get(1)
_ = W2502
tmp5902 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2502)


if True == tmp5902 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2502)
return
}


}, 1)

tmp5912 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2501)
}
__typedArg0 := V2501
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres5903 Obj

if True == tmp5912 {
tmp5904 := MakeNative(func(__e *ControlFlow) {
W2503 := __e.Get(1)
_ = W2503
tmp5905 := MakeNative(func(__e *ControlFlow) {
W2504 := __e.Get(1)
_ = W2504
tmp5907 := Call(__e, PrimFunc(symshen_4return_2), W2503)


if True == tmp5907 {
__e.TailApply(PrimFunc(symshen_4comb), W2504, symshen_4skip)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp5908 := Call(__e, PrimFunc(symtail), V2501)


__e.TailApply(tmp5905, tmp5908)
return


}, 1)

tmp5909 := Call(__e, PrimFunc(symhead), V2501)


tmp5910 := Call(__e, tmp5904, tmp5909)


ifres5903 = tmp5910


} else {
tmp5911 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5903 = tmp5911


}

__e.TailApply(tmp5900, ifres5903)
return


}, 1)

tmp5913 := Call(__e, ns2_1set, symshen_4_5return_6, tmp5899)


_ = tmp5913

tmp5914 := MakeNative(func(__e *ControlFlow) {
V2505 := __e.Get(1)
_ = V2505
tmp5915 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(13), Nil)
}
__typedArg0 := MakeNumber(13)
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp5916 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(10), tmp5915)
}
__typedArg0 := MakeNumber(10)
__typedArg1 := tmp5915
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp5917 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(9), tmp5916)
}
__typedArg0 := MakeNumber(9)
__typedArg1 := tmp5916
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symelement_2), V2505, tmp5917)
return


}, 1)

tmp5918 := Call(__e, ns2_1set, symshen_4return_2, tmp5914)


_ = tmp5918

tmp5919 := MakeNative(func(__e *ControlFlow) {
V2506 := __e.Get(1)
_ = V2506
tmp5920 := MakeNative(func(__e *ControlFlow) {
W2507 := __e.Get(1)
_ = W2507
tmp5922 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2507)


if True == tmp5922 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2507)
return
}


}, 1)

tmp5923 := MakeNative(func(__e *ControlFlow) {
W2508 := __e.Get(1)
_ = W2508
tmp5939 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2508)


if True == tmp5939 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5924 := MakeNative(func(__e *ControlFlow) {
W2509 := __e.Get(1)
_ = W2509
tmp5925 := MakeNative(func(__e *ControlFlow) {
W2510 := __e.Get(1)
_ = W2510
tmp5935 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2510)


if True == tmp5935 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5926 := MakeNative(func(__e *ControlFlow) {
W2511 := __e.Get(1)
_ = W2511
tmp5927 := MakeNative(func(__e *ControlFlow) {
W2512 := __e.Get(1)
_ = W2512
tmp5931 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2512)


if True == tmp5931 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5928 := MakeNative(func(__e *ControlFlow) {
W2513 := __e.Get(1)
_ = W2513
__e.TailApply(PrimFunc(symshen_4comb), W2513, symshen_4skip)
return
}, 1)

tmp5929 := Call(__e, PrimFunc(symshen_4in_1_6), W2512)


__e.TailApply(tmp5928, tmp5929)
return


}


}, 1)

tmp5932 := Call(__e, PrimFunc(symshen_4_5longnatter_6), W2511)


__e.TailApply(tmp5927, tmp5932)
return


}, 1)

tmp5933 := Call(__e, PrimFunc(symshen_4in_1_6), W2510)


__e.TailApply(tmp5926, tmp5933)
return


}


}, 1)

tmp5936 := Call(__e, PrimFunc(symshen_4_5times_6), W2509)


__e.TailApply(tmp5925, tmp5936)
return


}, 1)

tmp5937 := Call(__e, PrimFunc(symshen_4in_1_6), W2508)


__e.TailApply(tmp5924, tmp5937)
return


}


}, 1)

tmp5940 := Call(__e, PrimFunc(symshen_4_5backslash_6), V2506)


tmp5941 := Call(__e, tmp5923, tmp5940)


__e.TailApply(tmp5920, tmp5941)
return


}, 1)

tmp5942 := Call(__e, ns2_1set, symshen_4_5multiline_6, tmp5919)


_ = tmp5942

tmp5943 := MakeNative(func(__e *ControlFlow) {
V2514 := __e.Get(1)
_ = V2514
tmp5944 := MakeNative(func(__e *ControlFlow) {
W2515 := __e.Get(1)
_ = W2515
tmp5946 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2515)


if True == tmp5946 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2515)
return
}


}, 1)

tmp5952 := Call(__e, PrimFunc(symshen_4hds_a_2), V2514, MakeNumber(42))


var ifres5947 Obj

if True == tmp5952 {
tmp5948 := MakeNative(func(__e *ControlFlow) {
W2516 := __e.Get(1)
_ = W2516
__e.TailApply(PrimFunc(symshen_4comb), W2516, symshen_4skip)
return
}, 1)

tmp5949 := Call(__e, PrimFunc(symtail), V2514)


tmp5950 := Call(__e, tmp5948, tmp5949)


ifres5947 = tmp5950


} else {
tmp5951 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5947 = tmp5951


}

__e.TailApply(tmp5944, ifres5947)
return


}, 1)

tmp5953 := Call(__e, ns2_1set, symshen_4_5times_6, tmp5943)


_ = tmp5953

tmp5954 := MakeNative(func(__e *ControlFlow) {
V2517 := __e.Get(1)
_ = V2517
tmp5955 := MakeNative(func(__e *ControlFlow) {
W2518 := __e.Get(1)
_ = W2518
tmp5988 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2518)


if True == tmp5988 {
tmp5956 := MakeNative(func(__e *ControlFlow) {
W2523 := __e.Get(1)
_ = W2523
tmp5973 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2523)


if True == tmp5973 {
tmp5957 := MakeNative(func(__e *ControlFlow) {
W2528 := __e.Get(1)
_ = W2528
tmp5959 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2528)


if True == tmp5959 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2528)
return
}


}, 1)

tmp5971 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2517)
}
__typedArg0 := V2517
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres5960 Obj

if True == tmp5971 {
tmp5961 := MakeNative(func(__e *ControlFlow) {
W2529 := __e.Get(1)
_ = W2529
tmp5962 := MakeNative(func(__e *ControlFlow) {
W2530 := __e.Get(1)
_ = W2530
tmp5966 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2530)


if True == tmp5966 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5963 := MakeNative(func(__e *ControlFlow) {
W2531 := __e.Get(1)
_ = W2531
__e.TailApply(PrimFunc(symshen_4comb), W2531, symshen_4skip)
return
}, 1)

tmp5964 := Call(__e, PrimFunc(symshen_4in_1_6), W2530)


__e.TailApply(tmp5963, tmp5964)
return


}


}, 1)

tmp5967 := Call(__e, PrimFunc(symshen_4_5longnatter_6), W2529)


__e.TailApply(tmp5962, tmp5967)
return


}, 1)

tmp5968 := Call(__e, PrimFunc(symtail), V2517)


tmp5969 := Call(__e, tmp5961, tmp5968)


ifres5960 = tmp5969


} else {
tmp5970 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5960 = tmp5970


}

__e.TailApply(tmp5957, ifres5960)
return


} else {
__e.Return(W2523)
return
}


}, 1)

tmp5974 := MakeNative(func(__e *ControlFlow) {
W2524 := __e.Get(1)
_ = W2524
tmp5984 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2524)


if True == tmp5984 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5975 := MakeNative(func(__e *ControlFlow) {
W2525 := __e.Get(1)
_ = W2525
tmp5976 := MakeNative(func(__e *ControlFlow) {
W2526 := __e.Get(1)
_ = W2526
tmp5980 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2526)


if True == tmp5980 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5977 := MakeNative(func(__e *ControlFlow) {
W2527 := __e.Get(1)
_ = W2527
__e.TailApply(PrimFunc(symshen_4comb), W2527, symshen_4skip)
return
}, 1)

tmp5978 := Call(__e, PrimFunc(symshen_4in_1_6), W2526)


__e.TailApply(tmp5977, tmp5978)
return


}


}, 1)

tmp5981 := Call(__e, PrimFunc(symshen_4_5backslash_6), W2525)


__e.TailApply(tmp5976, tmp5981)
return


}, 1)

tmp5982 := Call(__e, PrimFunc(symshen_4in_1_6), W2524)


__e.TailApply(tmp5975, tmp5982)
return


}


}, 1)

tmp5985 := Call(__e, PrimFunc(symshen_4_5times_6), V2517)


tmp5986 := Call(__e, tmp5974, tmp5985)


__e.TailApply(tmp5956, tmp5986)
return


} else {
__e.Return(W2518)
return
}


}, 1)

tmp5989 := MakeNative(func(__e *ControlFlow) {
W2519 := __e.Get(1)
_ = W2519
tmp5999 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2519)


if True == tmp5999 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5990 := MakeNative(func(__e *ControlFlow) {
W2520 := __e.Get(1)
_ = W2520
tmp5991 := MakeNative(func(__e *ControlFlow) {
W2521 := __e.Get(1)
_ = W2521
tmp5995 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2521)


if True == tmp5995 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5992 := MakeNative(func(__e *ControlFlow) {
W2522 := __e.Get(1)
_ = W2522
__e.TailApply(PrimFunc(symshen_4comb), W2522, symshen_4skip)
return
}, 1)

tmp5993 := Call(__e, PrimFunc(symshen_4in_1_6), W2521)


__e.TailApply(tmp5992, tmp5993)
return


}


}, 1)

tmp5996 := Call(__e, PrimFunc(symshen_4_5longnatter_6), W2520)


__e.TailApply(tmp5991, tmp5996)
return


}, 1)

tmp5997 := Call(__e, PrimFunc(symshen_4in_1_6), W2519)


__e.TailApply(tmp5990, tmp5997)
return


}


}, 1)

tmp6000 := Call(__e, PrimFunc(symshen_4_5comment_6), V2517)


tmp6001 := Call(__e, tmp5989, tmp6000)


__e.TailApply(tmp5955, tmp6001)
return


}, 1)

tmp6002 := Call(__e, ns2_1set, symshen_4_5longnatter_6, tmp5954)


_ = tmp6002

tmp6003 := MakeNative(func(__e *ControlFlow) {
V2532 := __e.Get(1)
_ = V2532
tmp6004 := MakeNative(func(__e *ControlFlow) {
W2533 := __e.Get(1)
_ = W2533
tmp6035 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2533)


if True == tmp6035 {
tmp6005 := MakeNative(func(__e *ControlFlow) {
W2537 := __e.Get(1)
_ = W2537
tmp6024 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2537)


if True == tmp6024 {
tmp6006 := MakeNative(func(__e *ControlFlow) {
W2541 := __e.Get(1)
_ = W2541
tmp6008 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2541)


if True == tmp6008 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2541)
return
}


}, 1)

tmp6009 := MakeNative(func(__e *ControlFlow) {
W2542 := __e.Get(1)
_ = W2542
tmp6020 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2542)


if True == tmp6020 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6010 := MakeNative(func(__e *ControlFlow) {
W2543 := __e.Get(1)
_ = W2543
tmp6011 := MakeNative(func(__e *ControlFlow) {
W2544 := __e.Get(1)
_ = W2544
tmp6016 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W2543, MakeString("<>"))
}
__typedArg0 := W2543
__typedArg1 := MakeString("<>")
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres6012 Obj

if True == tmp6016 {
tmp6013 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(0), Nil)
}
__typedArg0 := MakeNumber(0)
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6014 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp6013)
}
__typedArg0 := symvector
__typedArg1 := tmp6013
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

ifres6012 = tmp6014


} else {
tmp6015 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(W2543)
}
__typedArg0 := W2543
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

ifres6012 = tmp6015


}

__e.TailApply(PrimFunc(symshen_4comb), W2544, ifres6012)
return


}, 1)

tmp6017 := Call(__e, PrimFunc(symshen_4in_1_6), W2542)


__e.TailApply(tmp6011, tmp6017)
return


}, 1)

tmp6018 := Call(__e, PrimFunc(symshen_4_5_1out), W2542)


__e.TailApply(tmp6010, tmp6018)
return


}


}, 1)

tmp6021 := Call(__e, PrimFunc(symshen_4_5sym_6), V2532)


tmp6022 := Call(__e, tmp6009, tmp6021)


__e.TailApply(tmp6006, tmp6022)
return


} else {
__e.Return(W2537)
return
}


}, 1)

tmp6025 := MakeNative(func(__e *ControlFlow) {
W2538 := __e.Get(1)
_ = W2538
tmp6031 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2538)


if True == tmp6031 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6026 := MakeNative(func(__e *ControlFlow) {
W2539 := __e.Get(1)
_ = W2539
tmp6027 := MakeNative(func(__e *ControlFlow) {
W2540 := __e.Get(1)
_ = W2540
__e.TailApply(PrimFunc(symshen_4comb), W2540, W2539)
return
}, 1)

tmp6028 := Call(__e, PrimFunc(symshen_4in_1_6), W2538)


__e.TailApply(tmp6027, tmp6028)
return


}, 1)

tmp6029 := Call(__e, PrimFunc(symshen_4_5_1out), W2538)


__e.TailApply(tmp6026, tmp6029)
return


}


}, 1)

tmp6032 := Call(__e, PrimFunc(symshen_4_5number_6), V2532)


tmp6033 := Call(__e, tmp6025, tmp6032)


__e.TailApply(tmp6005, tmp6033)
return


} else {
__e.Return(W2533)
return
}


}, 1)

tmp6036 := MakeNative(func(__e *ControlFlow) {
W2534 := __e.Get(1)
_ = W2534
tmp6042 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2534)


if True == tmp6042 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6037 := MakeNative(func(__e *ControlFlow) {
W2535 := __e.Get(1)
_ = W2535
tmp6038 := MakeNative(func(__e *ControlFlow) {
W2536 := __e.Get(1)
_ = W2536
__e.TailApply(PrimFunc(symshen_4comb), W2536, W2535)
return
}, 1)

tmp6039 := Call(__e, PrimFunc(symshen_4in_1_6), W2534)


__e.TailApply(tmp6038, tmp6039)
return


}, 1)

tmp6040 := Call(__e, PrimFunc(symshen_4_5_1out), W2534)


__e.TailApply(tmp6037, tmp6040)
return


}


}, 1)

tmp6043 := Call(__e, PrimFunc(symshen_4_5str_6), V2532)


tmp6044 := Call(__e, tmp6036, tmp6043)


__e.TailApply(tmp6004, tmp6044)
return


}, 1)

tmp6045 := Call(__e, ns2_1set, symshen_4_5atom_6, tmp6003)


_ = tmp6045

tmp6046 := MakeNative(func(__e *ControlFlow) {
V2545 := __e.Get(1)
_ = V2545
tmp6047 := MakeNative(func(__e *ControlFlow) {
W2546 := __e.Get(1)
_ = W2546
tmp6049 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2546)


if True == tmp6049 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2546)
return
}


}, 1)

tmp6050 := MakeNative(func(__e *ControlFlow) {
W2547 := __e.Get(1)
_ = W2547
tmp6065 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2547)


if True == tmp6065 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6051 := MakeNative(func(__e *ControlFlow) {
W2548 := __e.Get(1)
_ = W2548
tmp6052 := MakeNative(func(__e *ControlFlow) {
W2549 := __e.Get(1)
_ = W2549
tmp6053 := MakeNative(func(__e *ControlFlow) {
W2550 := __e.Get(1)
_ = W2550
tmp6060 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2550)


if True == tmp6060 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6054 := MakeNative(func(__e *ControlFlow) {
W2551 := __e.Get(1)
_ = W2551
tmp6055 := MakeNative(func(__e *ControlFlow) {
W2552 := __e.Get(1)
_ = W2552
__e.TailApply(PrimFunc(symshen_4comb), W2552, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(W2548)
__typedS1, __typedOK1 := TypedString(W2551)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := W2548
__typedArg1 := W2551
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp6057 := Call(__e, PrimFunc(symshen_4in_1_6), W2550)


__e.TailApply(tmp6055, tmp6057)
return


}, 1)

tmp6058 := Call(__e, PrimFunc(symshen_4_5_1out), W2550)


__e.TailApply(tmp6054, tmp6058)
return


}


}, 1)

tmp6061 := Call(__e, PrimFunc(symshen_4_5alphanums_6), W2549)


__e.TailApply(tmp6053, tmp6061)
return


}, 1)

tmp6062 := Call(__e, PrimFunc(symshen_4in_1_6), W2547)


__e.TailApply(tmp6052, tmp6062)
return


}, 1)

tmp6063 := Call(__e, PrimFunc(symshen_4_5_1out), W2547)


__e.TailApply(tmp6051, tmp6063)
return


}


}, 1)

tmp6066 := Call(__e, PrimFunc(symshen_4_5alpha_6), V2545)


tmp6067 := Call(__e, tmp6050, tmp6066)


__e.TailApply(tmp6047, tmp6067)
return


}, 1)

tmp6068 := Call(__e, ns2_1set, symshen_4_5sym_6, tmp6046)


_ = tmp6068

tmp6069 := MakeNative(func(__e *ControlFlow) {
V2553 := __e.Get(1)
_ = V2553
tmp6070 := MakeNative(func(__e *ControlFlow) {
W2554 := __e.Get(1)
_ = W2554
tmp6072 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2554)


if True == tmp6072 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2554)
return
}


}, 1)

tmp6083 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2553)
}
__typedArg0 := V2553
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres6073 Obj

if True == tmp6083 {
tmp6074 := MakeNative(func(__e *ControlFlow) {
W2555 := __e.Get(1)
_ = W2555
tmp6075 := MakeNative(func(__e *ControlFlow) {
W2556 := __e.Get(1)
_ = W2556
tmp6078 := Call(__e, PrimFunc(symshen_4alpha_2), W2555)


if True == tmp6078 {
tmp6076 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symn_1_6string) {
return PrimNumberToString(W2555)
}
__typedArg0 := W2555
return Call(__e, PrimFunc(symn_1_6string), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4comb), W2556, tmp6076)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp6079 := Call(__e, PrimFunc(symtail), V2553)


__e.TailApply(tmp6075, tmp6079)
return


}, 1)

tmp6080 := Call(__e, PrimFunc(symhead), V2553)


tmp6081 := Call(__e, tmp6074, tmp6080)


ifres6073 = tmp6081


} else {
tmp6082 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6073 = tmp6082


}

__e.TailApply(tmp6070, ifres6073)
return


}, 1)

tmp6084 := Call(__e, ns2_1set, symshen_4_5alpha_6, tmp6069)


_ = tmp6084

tmp6085 := MakeNative(func(__e *ControlFlow) {
V2557 := __e.Get(1)
_ = V2557
tmp6092 := Call(__e, PrimFunc(symshen_4lowercase_2), V2557)


if True == tmp6092 {
__e.Return(True)
return
} else {
tmp6090 := Call(__e, PrimFunc(symshen_4uppercase_2), V2557)


var ifres6087 Obj

if True == tmp6090 {
ifres6087 = True


} else {
tmp6089 := Call(__e, PrimFunc(symshen_4misc_2), V2557)


var ifres6088 Obj

if True == tmp6089 {
ifres6088 = True


} else {
ifres6088 = False


}

ifres6087 = ifres6088


}

if True == ifres6087 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp6093 := Call(__e, ns2_1set, symshen_4alpha_2, tmp6085)


_ = tmp6093

tmp6094 := MakeNative(func(__e *ControlFlow) {
V2558 := __e.Get(1)
_ = V2558
if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_6_a) {
__typedN0, __typedOK0 := TypedFloat64(V2558)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(97))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_6_a) {
return TypedMaterializeBoolean((__typedN0 >= __typedN1))
}}
__typedArg0 := V2558
__typedArg1 := MakeNumber(97)
return Call(__e, PrimFunc(sym_6_a), __typedArg0, __typedArg1)
})() {
if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_a) {
__typedN0, __typedOK0 := TypedFloat64(V2558)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(122))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_5_a) {
return TypedMaterializeBoolean((__typedN0 <= __typedN1))
}}
__typedArg0 := V2558
__typedArg1 := MakeNumber(122)
return Call(__e, PrimFunc(sym_5_a), __typedArg0, __typedArg1)
})() {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


} else {
__e.Return(False)
return
}


}, 1)

tmp6099 := Call(__e, ns2_1set, symshen_4lowercase_2, tmp6094)


_ = tmp6099

tmp6100 := MakeNative(func(__e *ControlFlow) {
V2559 := __e.Get(1)
_ = V2559
if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_6_a) {
__typedN0, __typedOK0 := TypedFloat64(V2559)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(65))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_6_a) {
return TypedMaterializeBoolean((__typedN0 >= __typedN1))
}}
__typedArg0 := V2559
__typedArg1 := MakeNumber(65)
return Call(__e, PrimFunc(sym_6_a), __typedArg0, __typedArg1)
})() {
if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_a) {
__typedN0, __typedOK0 := TypedFloat64(V2559)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(90))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_5_a) {
return TypedMaterializeBoolean((__typedN0 <= __typedN1))
}}
__typedArg0 := V2559
__typedArg1 := MakeNumber(90)
return Call(__e, PrimFunc(sym_5_a), __typedArg0, __typedArg1)
})() {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


} else {
__e.Return(False)
return
}


}, 1)

tmp6105 := Call(__e, ns2_1set, symshen_4uppercase_2, tmp6100)


_ = tmp6105

tmp6106 := MakeNative(func(__e *ControlFlow) {
V2560 := __e.Get(1)
_ = V2560
tmp6107 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(96), Nil)
}
__typedArg0 := MakeNumber(96)
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6108 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(35), tmp6107)
}
__typedArg0 := MakeNumber(35)
__typedArg1 := tmp6107
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6109 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(39), tmp6108)
}
__typedArg0 := MakeNumber(39)
__typedArg1 := tmp6108
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6110 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(37), tmp6109)
}
__typedArg0 := MakeNumber(37)
__typedArg1 := tmp6109
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6111 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(38), tmp6110)
}
__typedArg0 := MakeNumber(38)
__typedArg1 := tmp6110
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6112 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(60), tmp6111)
}
__typedArg0 := MakeNumber(60)
__typedArg1 := tmp6111
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6113 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(62), tmp6112)
}
__typedArg0 := MakeNumber(62)
__typedArg1 := tmp6112
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6114 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(46), tmp6113)
}
__typedArg0 := MakeNumber(46)
__typedArg1 := tmp6113
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6115 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(126), tmp6114)
}
__typedArg0 := MakeNumber(126)
__typedArg1 := tmp6114
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6116 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(64), tmp6115)
}
__typedArg0 := MakeNumber(64)
__typedArg1 := tmp6115
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6117 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(33), tmp6116)
}
__typedArg0 := MakeNumber(33)
__typedArg1 := tmp6116
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6118 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(36), tmp6117)
}
__typedArg0 := MakeNumber(36)
__typedArg1 := tmp6117
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6119 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(63), tmp6118)
}
__typedArg0 := MakeNumber(63)
__typedArg1 := tmp6118
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6120 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(95), tmp6119)
}
__typedArg0 := MakeNumber(95)
__typedArg1 := tmp6119
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6121 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(43), tmp6120)
}
__typedArg0 := MakeNumber(43)
__typedArg1 := tmp6120
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6122 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(47), tmp6121)
}
__typedArg0 := MakeNumber(47)
__typedArg1 := tmp6121
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6123 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(42), tmp6122)
}
__typedArg0 := MakeNumber(42)
__typedArg1 := tmp6122
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6124 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(45), tmp6123)
}
__typedArg0 := MakeNumber(45)
__typedArg1 := tmp6123
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6125 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(61), tmp6124)
}
__typedArg0 := MakeNumber(61)
__typedArg1 := tmp6124
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symelement_2), V2560, tmp6125)
return


}, 1)

tmp6126 := Call(__e, ns2_1set, symshen_4misc_2, tmp6106)


_ = tmp6126

tmp6127 := MakeNative(func(__e *ControlFlow) {
V2561 := __e.Get(1)
_ = V2561
tmp6128 := MakeNative(func(__e *ControlFlow) {
W2562 := __e.Get(1)
_ = W2562
tmp6140 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2562)


if True == tmp6140 {
tmp6129 := MakeNative(func(__e *ControlFlow) {
W2569 := __e.Get(1)
_ = W2569
tmp6131 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2569)


if True == tmp6131 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2569)
return
}


}, 1)

tmp6132 := MakeNative(func(__e *ControlFlow) {
W2570 := __e.Get(1)
_ = W2570
tmp6136 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2570)


if True == tmp6136 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6133 := MakeNative(func(__e *ControlFlow) {
W2571 := __e.Get(1)
_ = W2571
__e.TailApply(PrimFunc(symshen_4comb), W2571, MakeString(""))
return
}, 1)

tmp6134 := Call(__e, PrimFunc(symshen_4in_1_6), W2570)


__e.TailApply(tmp6133, tmp6134)
return


}


}, 1)

tmp6137 := Call(__e, PrimFunc(sym_5e_6), V2561)


tmp6138 := Call(__e, tmp6132, tmp6137)


__e.TailApply(tmp6129, tmp6138)
return


} else {
__e.Return(W2562)
return
}


}, 1)

tmp6141 := MakeNative(func(__e *ControlFlow) {
W2563 := __e.Get(1)
_ = W2563
tmp6156 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2563)


if True == tmp6156 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6142 := MakeNative(func(__e *ControlFlow) {
W2564 := __e.Get(1)
_ = W2564
tmp6143 := MakeNative(func(__e *ControlFlow) {
W2565 := __e.Get(1)
_ = W2565
tmp6144 := MakeNative(func(__e *ControlFlow) {
W2566 := __e.Get(1)
_ = W2566
tmp6151 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2566)


if True == tmp6151 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6145 := MakeNative(func(__e *ControlFlow) {
W2567 := __e.Get(1)
_ = W2567
tmp6146 := MakeNative(func(__e *ControlFlow) {
W2568 := __e.Get(1)
_ = W2568
__e.TailApply(PrimFunc(symshen_4comb), W2568, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(W2564)
__typedS1, __typedOK1 := TypedString(W2567)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := W2564
__typedArg1 := W2567
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp6148 := Call(__e, PrimFunc(symshen_4in_1_6), W2566)


__e.TailApply(tmp6146, tmp6148)
return


}, 1)

tmp6149 := Call(__e, PrimFunc(symshen_4_5_1out), W2566)


__e.TailApply(tmp6145, tmp6149)
return


}


}, 1)

tmp6152 := Call(__e, PrimFunc(symshen_4_5alphanums_6), W2565)


__e.TailApply(tmp6144, tmp6152)
return


}, 1)

tmp6153 := Call(__e, PrimFunc(symshen_4in_1_6), W2563)


__e.TailApply(tmp6143, tmp6153)
return


}, 1)

tmp6154 := Call(__e, PrimFunc(symshen_4_5_1out), W2563)


__e.TailApply(tmp6142, tmp6154)
return


}


}, 1)

tmp6157 := Call(__e, PrimFunc(symshen_4_5alphanum_6), V2561)


tmp6158 := Call(__e, tmp6141, tmp6157)


__e.TailApply(tmp6128, tmp6158)
return


}, 1)

tmp6159 := Call(__e, ns2_1set, symshen_4_5alphanums_6, tmp6127)


_ = tmp6159

tmp6160 := MakeNative(func(__e *ControlFlow) {
V2572 := __e.Get(1)
_ = V2572
tmp6161 := MakeNative(func(__e *ControlFlow) {
W2573 := __e.Get(1)
_ = W2573
tmp6175 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2573)


if True == tmp6175 {
tmp6162 := MakeNative(func(__e *ControlFlow) {
W2577 := __e.Get(1)
_ = W2577
tmp6164 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2577)


if True == tmp6164 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2577)
return
}


}, 1)

tmp6165 := MakeNative(func(__e *ControlFlow) {
W2578 := __e.Get(1)
_ = W2578
tmp6171 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2578)


if True == tmp6171 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6166 := MakeNative(func(__e *ControlFlow) {
W2579 := __e.Get(1)
_ = W2579
tmp6167 := MakeNative(func(__e *ControlFlow) {
W2580 := __e.Get(1)
_ = W2580
__e.TailApply(PrimFunc(symshen_4comb), W2580, W2579)
return
}, 1)

tmp6168 := Call(__e, PrimFunc(symshen_4in_1_6), W2578)


__e.TailApply(tmp6167, tmp6168)
return


}, 1)

tmp6169 := Call(__e, PrimFunc(symshen_4_5_1out), W2578)


__e.TailApply(tmp6166, tmp6169)
return


}


}, 1)

tmp6172 := Call(__e, PrimFunc(symshen_4_5numeral_6), V2572)


tmp6173 := Call(__e, tmp6165, tmp6172)


__e.TailApply(tmp6162, tmp6173)
return


} else {
__e.Return(W2573)
return
}


}, 1)

tmp6176 := MakeNative(func(__e *ControlFlow) {
W2574 := __e.Get(1)
_ = W2574
tmp6182 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2574)


if True == tmp6182 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6177 := MakeNative(func(__e *ControlFlow) {
W2575 := __e.Get(1)
_ = W2575
tmp6178 := MakeNative(func(__e *ControlFlow) {
W2576 := __e.Get(1)
_ = W2576
__e.TailApply(PrimFunc(symshen_4comb), W2576, W2575)
return
}, 1)

tmp6179 := Call(__e, PrimFunc(symshen_4in_1_6), W2574)


__e.TailApply(tmp6178, tmp6179)
return


}, 1)

tmp6180 := Call(__e, PrimFunc(symshen_4_5_1out), W2574)


__e.TailApply(tmp6177, tmp6180)
return


}


}, 1)

tmp6183 := Call(__e, PrimFunc(symshen_4_5alpha_6), V2572)


tmp6184 := Call(__e, tmp6176, tmp6183)


__e.TailApply(tmp6161, tmp6184)
return


}, 1)

tmp6185 := Call(__e, ns2_1set, symshen_4_5alphanum_6, tmp6160)


_ = tmp6185

tmp6186 := MakeNative(func(__e *ControlFlow) {
V2581 := __e.Get(1)
_ = V2581
tmp6187 := MakeNative(func(__e *ControlFlow) {
W2582 := __e.Get(1)
_ = W2582
tmp6189 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2582)


if True == tmp6189 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2582)
return
}


}, 1)

tmp6200 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2581)
}
__typedArg0 := V2581
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres6190 Obj

if True == tmp6200 {
tmp6191 := MakeNative(func(__e *ControlFlow) {
W2583 := __e.Get(1)
_ = W2583
tmp6192 := MakeNative(func(__e *ControlFlow) {
W2584 := __e.Get(1)
_ = W2584
tmp6195 := Call(__e, PrimFunc(symshen_4digit_2), W2583)


if True == tmp6195 {
tmp6193 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symn_1_6string) {
return PrimNumberToString(W2583)
}
__typedArg0 := W2583
return Call(__e, PrimFunc(symn_1_6string), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4comb), W2584, tmp6193)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp6196 := Call(__e, PrimFunc(symtail), V2581)


__e.TailApply(tmp6192, tmp6196)
return


}, 1)

tmp6197 := Call(__e, PrimFunc(symhead), V2581)


tmp6198 := Call(__e, tmp6191, tmp6197)


ifres6190 = tmp6198


} else {
tmp6199 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6190 = tmp6199


}

__e.TailApply(tmp6187, ifres6190)
return


}, 1)

tmp6201 := Call(__e, ns2_1set, symshen_4_5numeral_6, tmp6186)


_ = tmp6201

tmp6202 := MakeNative(func(__e *ControlFlow) {
V2585 := __e.Get(1)
_ = V2585
if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_6_a) {
__typedN0, __typedOK0 := TypedFloat64(V2585)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(48))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_6_a) {
return TypedMaterializeBoolean((__typedN0 >= __typedN1))
}}
__typedArg0 := V2585
__typedArg1 := MakeNumber(48)
return Call(__e, PrimFunc(sym_6_a), __typedArg0, __typedArg1)
})() {
if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_a) {
__typedN0, __typedOK0 := TypedFloat64(V2585)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(57))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_5_a) {
return TypedMaterializeBoolean((__typedN0 <= __typedN1))
}}
__typedArg0 := V2585
__typedArg1 := MakeNumber(57)
return Call(__e, PrimFunc(sym_5_a), __typedArg0, __typedArg1)
})() {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


} else {
__e.Return(False)
return
}


}, 1)

tmp6207 := Call(__e, ns2_1set, symshen_4digit_2, tmp6202)


_ = tmp6207

tmp6208 := MakeNative(func(__e *ControlFlow) {
V2586 := __e.Get(1)
_ = V2586
tmp6209 := MakeNative(func(__e *ControlFlow) {
W2587 := __e.Get(1)
_ = W2587
tmp6211 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2587)


if True == tmp6211 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2587)
return
}


}, 1)

tmp6212 := MakeNative(func(__e *ControlFlow) {
W2588 := __e.Get(1)
_ = W2588
tmp6230 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2588)


if True == tmp6230 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6213 := MakeNative(func(__e *ControlFlow) {
W2589 := __e.Get(1)
_ = W2589
tmp6214 := MakeNative(func(__e *ControlFlow) {
W2590 := __e.Get(1)
_ = W2590
tmp6226 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2590)


if True == tmp6226 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6215 := MakeNative(func(__e *ControlFlow) {
W2591 := __e.Get(1)
_ = W2591
tmp6216 := MakeNative(func(__e *ControlFlow) {
W2592 := __e.Get(1)
_ = W2592
tmp6217 := MakeNative(func(__e *ControlFlow) {
W2593 := __e.Get(1)
_ = W2593
tmp6221 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2593)


if True == tmp6221 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6218 := MakeNative(func(__e *ControlFlow) {
W2594 := __e.Get(1)
_ = W2594
__e.TailApply(PrimFunc(symshen_4comb), W2594, W2591)
return
}, 1)

tmp6219 := Call(__e, PrimFunc(symshen_4in_1_6), W2593)


__e.TailApply(tmp6218, tmp6219)
return


}


}, 1)

tmp6222 := Call(__e, PrimFunc(symshen_4_5dbq_6), W2592)


__e.TailApply(tmp6217, tmp6222)
return


}, 1)

tmp6223 := Call(__e, PrimFunc(symshen_4in_1_6), W2590)


__e.TailApply(tmp6216, tmp6223)
return


}, 1)

tmp6224 := Call(__e, PrimFunc(symshen_4_5_1out), W2590)


__e.TailApply(tmp6215, tmp6224)
return


}


}, 1)

tmp6227 := Call(__e, PrimFunc(symshen_4_5strcontents_6), W2589)


__e.TailApply(tmp6214, tmp6227)
return


}, 1)

tmp6228 := Call(__e, PrimFunc(symshen_4in_1_6), W2588)


__e.TailApply(tmp6213, tmp6228)
return


}


}, 1)

tmp6231 := Call(__e, PrimFunc(symshen_4_5dbq_6), V2586)


tmp6232 := Call(__e, tmp6212, tmp6231)


__e.TailApply(tmp6209, tmp6232)
return


}, 1)

tmp6233 := Call(__e, ns2_1set, symshen_4_5str_6, tmp6208)


_ = tmp6233

tmp6234 := MakeNative(func(__e *ControlFlow) {
V2595 := __e.Get(1)
_ = V2595
tmp6235 := MakeNative(func(__e *ControlFlow) {
W2596 := __e.Get(1)
_ = W2596
tmp6237 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2596)


if True == tmp6237 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2596)
return
}


}, 1)

tmp6243 := Call(__e, PrimFunc(symshen_4hds_a_2), V2595, MakeNumber(34))


var ifres6238 Obj

if True == tmp6243 {
tmp6239 := MakeNative(func(__e *ControlFlow) {
W2597 := __e.Get(1)
_ = W2597
__e.TailApply(PrimFunc(symshen_4comb), W2597, symshen_4skip)
return
}, 1)

tmp6240 := Call(__e, PrimFunc(symtail), V2595)


tmp6241 := Call(__e, tmp6239, tmp6240)


ifres6238 = tmp6241


} else {
tmp6242 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6238 = tmp6242


}

__e.TailApply(tmp6235, ifres6238)
return


}, 1)

tmp6244 := Call(__e, ns2_1set, symshen_4_5dbq_6, tmp6234)


_ = tmp6244

tmp6245 := MakeNative(func(__e *ControlFlow) {
V2598 := __e.Get(1)
_ = V2598
tmp6246 := MakeNative(func(__e *ControlFlow) {
W2599 := __e.Get(1)
_ = W2599
tmp6258 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2599)


if True == tmp6258 {
tmp6247 := MakeNative(func(__e *ControlFlow) {
W2606 := __e.Get(1)
_ = W2606
tmp6249 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2606)


if True == tmp6249 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2606)
return
}


}, 1)

tmp6250 := MakeNative(func(__e *ControlFlow) {
W2607 := __e.Get(1)
_ = W2607
tmp6254 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2607)


if True == tmp6254 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6251 := MakeNative(func(__e *ControlFlow) {
W2608 := __e.Get(1)
_ = W2608
__e.TailApply(PrimFunc(symshen_4comb), W2608, MakeString(""))
return
}, 1)

tmp6252 := Call(__e, PrimFunc(symshen_4in_1_6), W2607)


__e.TailApply(tmp6251, tmp6252)
return


}


}, 1)

tmp6255 := Call(__e, PrimFunc(sym_5e_6), V2598)


tmp6256 := Call(__e, tmp6250, tmp6255)


__e.TailApply(tmp6247, tmp6256)
return


} else {
__e.Return(W2599)
return
}


}, 1)

tmp6259 := MakeNative(func(__e *ControlFlow) {
W2600 := __e.Get(1)
_ = W2600
tmp6274 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2600)


if True == tmp6274 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6260 := MakeNative(func(__e *ControlFlow) {
W2601 := __e.Get(1)
_ = W2601
tmp6261 := MakeNative(func(__e *ControlFlow) {
W2602 := __e.Get(1)
_ = W2602
tmp6262 := MakeNative(func(__e *ControlFlow) {
W2603 := __e.Get(1)
_ = W2603
tmp6269 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2603)


if True == tmp6269 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6263 := MakeNative(func(__e *ControlFlow) {
W2604 := __e.Get(1)
_ = W2604
tmp6264 := MakeNative(func(__e *ControlFlow) {
W2605 := __e.Get(1)
_ = W2605
__e.TailApply(PrimFunc(symshen_4comb), W2605, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(W2601)
__typedS1, __typedOK1 := TypedString(W2604)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := W2601
__typedArg1 := W2604
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp6266 := Call(__e, PrimFunc(symshen_4in_1_6), W2603)


__e.TailApply(tmp6264, tmp6266)
return


}, 1)

tmp6267 := Call(__e, PrimFunc(symshen_4_5_1out), W2603)


__e.TailApply(tmp6263, tmp6267)
return


}


}, 1)

tmp6270 := Call(__e, PrimFunc(symshen_4_5strcontents_6), W2602)


__e.TailApply(tmp6262, tmp6270)
return


}, 1)

tmp6271 := Call(__e, PrimFunc(symshen_4in_1_6), W2600)


__e.TailApply(tmp6261, tmp6271)
return


}, 1)

tmp6272 := Call(__e, PrimFunc(symshen_4_5_1out), W2600)


__e.TailApply(tmp6260, tmp6272)
return


}


}, 1)

tmp6275 := Call(__e, PrimFunc(symshen_4_5strc_6), V2598)


tmp6276 := Call(__e, tmp6259, tmp6275)


__e.TailApply(tmp6246, tmp6276)
return


}, 1)

tmp6277 := Call(__e, ns2_1set, symshen_4_5strcontents_6, tmp6245)


_ = tmp6277

tmp6278 := MakeNative(func(__e *ControlFlow) {
V2609 := __e.Get(1)
_ = V2609
tmp6279 := MakeNative(func(__e *ControlFlow) {
W2610 := __e.Get(1)
_ = W2610
tmp6293 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2610)


if True == tmp6293 {
tmp6280 := MakeNative(func(__e *ControlFlow) {
W2614 := __e.Get(1)
_ = W2614
tmp6282 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2614)


if True == tmp6282 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2614)
return
}


}, 1)

tmp6283 := MakeNative(func(__e *ControlFlow) {
W2615 := __e.Get(1)
_ = W2615
tmp6289 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2615)


if True == tmp6289 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6284 := MakeNative(func(__e *ControlFlow) {
W2616 := __e.Get(1)
_ = W2616
tmp6285 := MakeNative(func(__e *ControlFlow) {
W2617 := __e.Get(1)
_ = W2617
__e.TailApply(PrimFunc(symshen_4comb), W2617, W2616)
return
}, 1)

tmp6286 := Call(__e, PrimFunc(symshen_4in_1_6), W2615)


__e.TailApply(tmp6285, tmp6286)
return


}, 1)

tmp6287 := Call(__e, PrimFunc(symshen_4_5_1out), W2615)


__e.TailApply(tmp6284, tmp6287)
return


}


}, 1)

tmp6290 := Call(__e, PrimFunc(symshen_4_5notdbq_6), V2609)


tmp6291 := Call(__e, tmp6283, tmp6290)


__e.TailApply(tmp6280, tmp6291)
return


} else {
__e.Return(W2610)
return
}


}, 1)

tmp6294 := MakeNative(func(__e *ControlFlow) {
W2611 := __e.Get(1)
_ = W2611
tmp6300 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2611)


if True == tmp6300 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6295 := MakeNative(func(__e *ControlFlow) {
W2612 := __e.Get(1)
_ = W2612
tmp6296 := MakeNative(func(__e *ControlFlow) {
W2613 := __e.Get(1)
_ = W2613
__e.TailApply(PrimFunc(symshen_4comb), W2613, W2612)
return
}, 1)

tmp6297 := Call(__e, PrimFunc(symshen_4in_1_6), W2611)


__e.TailApply(tmp6296, tmp6297)
return


}, 1)

tmp6298 := Call(__e, PrimFunc(symshen_4_5_1out), W2611)


__e.TailApply(tmp6295, tmp6298)
return


}


}, 1)

tmp6301 := Call(__e, PrimFunc(symshen_4_5control_6), V2609)


tmp6302 := Call(__e, tmp6294, tmp6301)


__e.TailApply(tmp6279, tmp6302)
return


}, 1)

tmp6303 := Call(__e, ns2_1set, symshen_4_5strc_6, tmp6278)


_ = tmp6303

tmp6304 := MakeNative(func(__e *ControlFlow) {
V2618 := __e.Get(1)
_ = V2618
tmp6305 := MakeNative(func(__e *ControlFlow) {
W2619 := __e.Get(1)
_ = W2619
tmp6307 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2619)


if True == tmp6307 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2619)
return
}


}, 1)

tmp6308 := MakeNative(func(__e *ControlFlow) {
W2620 := __e.Get(1)
_ = W2620
tmp6333 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2620)


if True == tmp6333 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6309 := MakeNative(func(__e *ControlFlow) {
W2621 := __e.Get(1)
_ = W2621
tmp6310 := MakeNative(func(__e *ControlFlow) {
W2622 := __e.Get(1)
_ = W2622
tmp6329 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2622)


if True == tmp6329 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6311 := MakeNative(func(__e *ControlFlow) {
W2623 := __e.Get(1)
_ = W2623
tmp6312 := MakeNative(func(__e *ControlFlow) {
W2624 := __e.Get(1)
_ = W2624
tmp6325 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2624)


if True == tmp6325 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6313 := MakeNative(func(__e *ControlFlow) {
W2625 := __e.Get(1)
_ = W2625
tmp6314 := MakeNative(func(__e *ControlFlow) {
W2626 := __e.Get(1)
_ = W2626
tmp6315 := MakeNative(func(__e *ControlFlow) {
W2627 := __e.Get(1)
_ = W2627
tmp6320 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2627)


if True == tmp6320 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6316 := MakeNative(func(__e *ControlFlow) {
W2628 := __e.Get(1)
_ = W2628
tmp6317 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symn_1_6string) {
return PrimNumberToString(W2625)
}
__typedArg0 := W2625
return Call(__e, PrimFunc(symn_1_6string), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4comb), W2628, tmp6317)
return


}, 1)

tmp6318 := Call(__e, PrimFunc(symshen_4in_1_6), W2627)


__e.TailApply(tmp6316, tmp6318)
return


}


}, 1)

tmp6321 := Call(__e, PrimFunc(symshen_4_5semicolon_6), W2626)


__e.TailApply(tmp6315, tmp6321)
return


}, 1)

tmp6322 := Call(__e, PrimFunc(symshen_4in_1_6), W2624)


__e.TailApply(tmp6314, tmp6322)
return


}, 1)

tmp6323 := Call(__e, PrimFunc(symshen_4_5_1out), W2624)


__e.TailApply(tmp6313, tmp6323)
return


}


}, 1)

tmp6326 := Call(__e, PrimFunc(symshen_4_5integer_6), W2623)


__e.TailApply(tmp6312, tmp6326)
return


}, 1)

tmp6327 := Call(__e, PrimFunc(symshen_4in_1_6), W2622)


__e.TailApply(tmp6311, tmp6327)
return


}


}, 1)

tmp6330 := Call(__e, PrimFunc(symshen_4_5hash_6), W2621)


__e.TailApply(tmp6310, tmp6330)
return


}, 1)

tmp6331 := Call(__e, PrimFunc(symshen_4in_1_6), W2620)


__e.TailApply(tmp6309, tmp6331)
return


}


}, 1)

tmp6334 := Call(__e, PrimFunc(symshen_4_5lowC_6), V2618)


tmp6335 := Call(__e, tmp6308, tmp6334)


__e.TailApply(tmp6305, tmp6335)
return


}, 1)

tmp6336 := Call(__e, ns2_1set, symshen_4_5control_6, tmp6304)


_ = tmp6336

tmp6337 := MakeNative(func(__e *ControlFlow) {
V2629 := __e.Get(1)
_ = V2629
tmp6338 := MakeNative(func(__e *ControlFlow) {
W2630 := __e.Get(1)
_ = W2630
tmp6340 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2630)


if True == tmp6340 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2630)
return
}


}, 1)

tmp6352 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2629)
}
__typedArg0 := V2629
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres6341 Obj

if True == tmp6352 {
tmp6342 := MakeNative(func(__e *ControlFlow) {
W2631 := __e.Get(1)
_ = W2631
tmp6343 := MakeNative(func(__e *ControlFlow) {
W2632 := __e.Get(1)
_ = W2632
tmp6346 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W2631, MakeNumber(34))
}
__typedArg0 := W2631
__typedArg1 := MakeNumber(34)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp6346)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp6346
return Call(__e, PrimFunc(symnot), __typedArg0)
})() {
tmp6344 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symn_1_6string) {
return PrimNumberToString(W2631)
}
__typedArg0 := W2631
return Call(__e, PrimFunc(symn_1_6string), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4comb), W2632, tmp6344)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp6348 := Call(__e, PrimFunc(symtail), V2629)


__e.TailApply(tmp6343, tmp6348)
return


}, 1)

tmp6349 := Call(__e, PrimFunc(symhead), V2629)


tmp6350 := Call(__e, tmp6342, tmp6349)


ifres6341 = tmp6350


} else {
tmp6351 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6341 = tmp6351


}

__e.TailApply(tmp6338, ifres6341)
return


}, 1)

tmp6353 := Call(__e, ns2_1set, symshen_4_5notdbq_6, tmp6337)


_ = tmp6353

tmp6354 := MakeNative(func(__e *ControlFlow) {
V2633 := __e.Get(1)
_ = V2633
tmp6355 := MakeNative(func(__e *ControlFlow) {
W2634 := __e.Get(1)
_ = W2634
tmp6357 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2634)


if True == tmp6357 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2634)
return
}


}, 1)

tmp6363 := Call(__e, PrimFunc(symshen_4hds_a_2), V2633, MakeNumber(99))


var ifres6358 Obj

if True == tmp6363 {
tmp6359 := MakeNative(func(__e *ControlFlow) {
W2635 := __e.Get(1)
_ = W2635
__e.TailApply(PrimFunc(symshen_4comb), W2635, symshen_4skip)
return
}, 1)

tmp6360 := Call(__e, PrimFunc(symtail), V2633)


tmp6361 := Call(__e, tmp6359, tmp6360)


ifres6358 = tmp6361


} else {
tmp6362 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6358 = tmp6362


}

__e.TailApply(tmp6355, ifres6358)
return


}, 1)

tmp6364 := Call(__e, ns2_1set, symshen_4_5lowC_6, tmp6354)


_ = tmp6364

tmp6365 := MakeNative(func(__e *ControlFlow) {
V2636 := __e.Get(1)
_ = V2636
tmp6366 := MakeNative(func(__e *ControlFlow) {
W2637 := __e.Get(1)
_ = W2637
tmp6368 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2637)


if True == tmp6368 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2637)
return
}


}, 1)

tmp6374 := Call(__e, PrimFunc(symshen_4hds_a_2), V2636, MakeNumber(35))


var ifres6369 Obj

if True == tmp6374 {
tmp6370 := MakeNative(func(__e *ControlFlow) {
W2638 := __e.Get(1)
_ = W2638
__e.TailApply(PrimFunc(symshen_4comb), W2638, symshen_4skip)
return
}, 1)

tmp6371 := Call(__e, PrimFunc(symtail), V2636)


tmp6372 := Call(__e, tmp6370, tmp6371)


ifres6369 = tmp6372


} else {
tmp6373 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6369 = tmp6373


}

__e.TailApply(tmp6366, ifres6369)
return


}, 1)

tmp6375 := Call(__e, ns2_1set, symshen_4_5hash_6, tmp6365)


_ = tmp6375

tmp6376 := MakeNative(func(__e *ControlFlow) {
V2639 := __e.Get(1)
_ = V2639
tmp6377 := MakeNative(func(__e *ControlFlow) {
W2640 := __e.Get(1)
_ = W2640
tmp6433 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2640)


if True == tmp6433 {
tmp6378 := MakeNative(func(__e *ControlFlow) {
W2646 := __e.Get(1)
_ = W2646
tmp6416 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2646)


if True == tmp6416 {
tmp6379 := MakeNative(func(__e *ControlFlow) {
W2652 := __e.Get(1)
_ = W2652
tmp6405 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2652)


if True == tmp6405 {
tmp6380 := MakeNative(func(__e *ControlFlow) {
W2656 := __e.Get(1)
_ = W2656
tmp6394 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2656)


if True == tmp6394 {
tmp6381 := MakeNative(func(__e *ControlFlow) {
W2660 := __e.Get(1)
_ = W2660
tmp6383 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2660)


if True == tmp6383 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2660)
return
}


}, 1)

tmp6384 := MakeNative(func(__e *ControlFlow) {
W2661 := __e.Get(1)
_ = W2661
tmp6390 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2661)


if True == tmp6390 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6385 := MakeNative(func(__e *ControlFlow) {
W2662 := __e.Get(1)
_ = W2662
tmp6386 := MakeNative(func(__e *ControlFlow) {
W2663 := __e.Get(1)
_ = W2663
__e.TailApply(PrimFunc(symshen_4comb), W2663, W2662)
return
}, 1)

tmp6387 := Call(__e, PrimFunc(symshen_4in_1_6), W2661)


__e.TailApply(tmp6386, tmp6387)
return


}, 1)

tmp6388 := Call(__e, PrimFunc(symshen_4_5_1out), W2661)


__e.TailApply(tmp6385, tmp6388)
return


}


}, 1)

tmp6391 := Call(__e, PrimFunc(symshen_4_5integer_6), V2639)


tmp6392 := Call(__e, tmp6384, tmp6391)


__e.TailApply(tmp6381, tmp6392)
return


} else {
__e.Return(W2656)
return
}


}, 1)

tmp6395 := MakeNative(func(__e *ControlFlow) {
W2657 := __e.Get(1)
_ = W2657
tmp6401 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2657)


if True == tmp6401 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6396 := MakeNative(func(__e *ControlFlow) {
W2658 := __e.Get(1)
_ = W2658
tmp6397 := MakeNative(func(__e *ControlFlow) {
W2659 := __e.Get(1)
_ = W2659
__e.TailApply(PrimFunc(symshen_4comb), W2659, W2658)
return
}, 1)

tmp6398 := Call(__e, PrimFunc(symshen_4in_1_6), W2657)


__e.TailApply(tmp6397, tmp6398)
return


}, 1)

tmp6399 := Call(__e, PrimFunc(symshen_4_5_1out), W2657)


__e.TailApply(tmp6396, tmp6399)
return


}


}, 1)

tmp6402 := Call(__e, PrimFunc(symshen_4_5float_6), V2639)


tmp6403 := Call(__e, tmp6395, tmp6402)


__e.TailApply(tmp6380, tmp6403)
return


} else {
__e.Return(W2652)
return
}


}, 1)

tmp6406 := MakeNative(func(__e *ControlFlow) {
W2653 := __e.Get(1)
_ = W2653
tmp6412 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2653)


if True == tmp6412 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6407 := MakeNative(func(__e *ControlFlow) {
W2654 := __e.Get(1)
_ = W2654
tmp6408 := MakeNative(func(__e *ControlFlow) {
W2655 := __e.Get(1)
_ = W2655
__e.TailApply(PrimFunc(symshen_4comb), W2655, W2654)
return
}, 1)

tmp6409 := Call(__e, PrimFunc(symshen_4in_1_6), W2653)


__e.TailApply(tmp6408, tmp6409)
return


}, 1)

tmp6410 := Call(__e, PrimFunc(symshen_4_5_1out), W2653)


__e.TailApply(tmp6407, tmp6410)
return


}


}, 1)

tmp6413 := Call(__e, PrimFunc(symshen_4_5e_1number_6), V2639)


tmp6414 := Call(__e, tmp6406, tmp6413)


__e.TailApply(tmp6379, tmp6414)
return


} else {
__e.Return(W2646)
return
}


}, 1)

tmp6417 := MakeNative(func(__e *ControlFlow) {
W2647 := __e.Get(1)
_ = W2647
tmp6429 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2647)


if True == tmp6429 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6418 := MakeNative(func(__e *ControlFlow) {
W2648 := __e.Get(1)
_ = W2648
tmp6419 := MakeNative(func(__e *ControlFlow) {
W2649 := __e.Get(1)
_ = W2649
tmp6425 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2649)


if True == tmp6425 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6420 := MakeNative(func(__e *ControlFlow) {
W2650 := __e.Get(1)
_ = W2650
tmp6421 := MakeNative(func(__e *ControlFlow) {
W2651 := __e.Get(1)
_ = W2651
__e.TailApply(PrimFunc(symshen_4comb), W2651, W2650)
return
}, 1)

tmp6422 := Call(__e, PrimFunc(symshen_4in_1_6), W2649)


__e.TailApply(tmp6421, tmp6422)
return


}, 1)

tmp6423 := Call(__e, PrimFunc(symshen_4_5_1out), W2649)


__e.TailApply(tmp6420, tmp6423)
return


}


}, 1)

tmp6426 := Call(__e, PrimFunc(symshen_4_5number_6), W2648)


__e.TailApply(tmp6419, tmp6426)
return


}, 1)

tmp6427 := Call(__e, PrimFunc(symshen_4in_1_6), W2647)


__e.TailApply(tmp6418, tmp6427)
return


}


}, 1)

tmp6430 := Call(__e, PrimFunc(symshen_4_5plus_6), V2639)


tmp6431 := Call(__e, tmp6417, tmp6430)


__e.TailApply(tmp6378, tmp6431)
return


} else {
__e.Return(W2640)
return
}


}, 1)

tmp6434 := MakeNative(func(__e *ControlFlow) {
W2641 := __e.Get(1)
_ = W2641
tmp6447 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2641)


if True == tmp6447 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6435 := MakeNative(func(__e *ControlFlow) {
W2642 := __e.Get(1)
_ = W2642
tmp6436 := MakeNative(func(__e *ControlFlow) {
W2643 := __e.Get(1)
_ = W2643
tmp6443 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2643)


if True == tmp6443 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6437 := MakeNative(func(__e *ControlFlow) {
W2644 := __e.Get(1)
_ = W2644
tmp6438 := MakeNative(func(__e *ControlFlow) {
W2645 := __e.Get(1)
_ = W2645
__e.TailApply(PrimFunc(symshen_4comb), W2645, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(0))
__typedN1, __typedOK1 := TypedFloat64(W2644)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := MakeNumber(0)
__typedArg1 := W2644
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp6440 := Call(__e, PrimFunc(symshen_4in_1_6), W2643)


__e.TailApply(tmp6438, tmp6440)
return


}, 1)

tmp6441 := Call(__e, PrimFunc(symshen_4_5_1out), W2643)


__e.TailApply(tmp6437, tmp6441)
return


}


}, 1)

tmp6444 := Call(__e, PrimFunc(symshen_4_5number_6), W2642)


__e.TailApply(tmp6436, tmp6444)
return


}, 1)

tmp6445 := Call(__e, PrimFunc(symshen_4in_1_6), W2641)


__e.TailApply(tmp6435, tmp6445)
return


}


}, 1)

tmp6448 := Call(__e, PrimFunc(symshen_4_5minus_6), V2639)


tmp6449 := Call(__e, tmp6434, tmp6448)


__e.TailApply(tmp6377, tmp6449)
return


}, 1)

tmp6450 := Call(__e, ns2_1set, symshen_4_5number_6, tmp6376)


_ = tmp6450

tmp6451 := MakeNative(func(__e *ControlFlow) {
V2664 := __e.Get(1)
_ = V2664
tmp6452 := MakeNative(func(__e *ControlFlow) {
W2665 := __e.Get(1)
_ = W2665
tmp6454 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2665)


if True == tmp6454 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2665)
return
}


}, 1)

tmp6460 := Call(__e, PrimFunc(symshen_4hds_a_2), V2664, MakeNumber(45))


var ifres6455 Obj

if True == tmp6460 {
tmp6456 := MakeNative(func(__e *ControlFlow) {
W2666 := __e.Get(1)
_ = W2666
__e.TailApply(PrimFunc(symshen_4comb), W2666, symshen_4skip)
return
}, 1)

tmp6457 := Call(__e, PrimFunc(symtail), V2664)


tmp6458 := Call(__e, tmp6456, tmp6457)


ifres6455 = tmp6458


} else {
tmp6459 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6455 = tmp6459


}

__e.TailApply(tmp6452, ifres6455)
return


}, 1)

tmp6461 := Call(__e, ns2_1set, symshen_4_5minus_6, tmp6451)


_ = tmp6461

tmp6462 := MakeNative(func(__e *ControlFlow) {
V2667 := __e.Get(1)
_ = V2667
tmp6463 := MakeNative(func(__e *ControlFlow) {
W2668 := __e.Get(1)
_ = W2668
tmp6465 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2668)


if True == tmp6465 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2668)
return
}


}, 1)

tmp6471 := Call(__e, PrimFunc(symshen_4hds_a_2), V2667, MakeNumber(43))


var ifres6466 Obj

if True == tmp6471 {
tmp6467 := MakeNative(func(__e *ControlFlow) {
W2669 := __e.Get(1)
_ = W2669
__e.TailApply(PrimFunc(symshen_4comb), W2669, symshen_4skip)
return
}, 1)

tmp6468 := Call(__e, PrimFunc(symtail), V2667)


tmp6469 := Call(__e, tmp6467, tmp6468)


ifres6466 = tmp6469


} else {
tmp6470 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6466 = tmp6470


}

__e.TailApply(tmp6463, ifres6466)
return


}, 1)

tmp6472 := Call(__e, ns2_1set, symshen_4_5plus_6, tmp6462)


_ = tmp6472

tmp6473 := MakeNative(func(__e *ControlFlow) {
V2670 := __e.Get(1)
_ = V2670
tmp6474 := MakeNative(func(__e *ControlFlow) {
W2671 := __e.Get(1)
_ = W2671
tmp6476 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2671)


if True == tmp6476 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2671)
return
}


}, 1)

tmp6477 := MakeNative(func(__e *ControlFlow) {
W2672 := __e.Get(1)
_ = W2672
tmp6484 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2672)


if True == tmp6484 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6478 := MakeNative(func(__e *ControlFlow) {
W2673 := __e.Get(1)
_ = W2673
tmp6479 := MakeNative(func(__e *ControlFlow) {
W2674 := __e.Get(1)
_ = W2674
tmp6480 := Call(__e, PrimFunc(symshen_4compute_1integer), W2673)


__e.TailApply(PrimFunc(symshen_4comb), W2674, tmp6480)
return


}, 1)

tmp6481 := Call(__e, PrimFunc(symshen_4in_1_6), W2672)


__e.TailApply(tmp6479, tmp6481)
return


}, 1)

tmp6482 := Call(__e, PrimFunc(symshen_4_5_1out), W2672)


__e.TailApply(tmp6478, tmp6482)
return


}


}, 1)

tmp6485 := Call(__e, PrimFunc(symshen_4_5digits_6), V2670)


tmp6486 := Call(__e, tmp6477, tmp6485)


__e.TailApply(tmp6474, tmp6486)
return


}, 1)

tmp6487 := Call(__e, ns2_1set, symshen_4_5integer_6, tmp6473)


_ = tmp6487

tmp6488 := MakeNative(func(__e *ControlFlow) {
V2675 := __e.Get(1)
_ = V2675
tmp6489 := MakeNative(func(__e *ControlFlow) {
W2676 := __e.Get(1)
_ = W2676
tmp6504 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2676)


if True == tmp6504 {
tmp6490 := MakeNative(func(__e *ControlFlow) {
W2683 := __e.Get(1)
_ = W2683
tmp6492 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2683)


if True == tmp6492 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2683)
return
}


}, 1)

tmp6493 := MakeNative(func(__e *ControlFlow) {
W2684 := __e.Get(1)
_ = W2684
tmp6500 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2684)


if True == tmp6500 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6494 := MakeNative(func(__e *ControlFlow) {
W2685 := __e.Get(1)
_ = W2685
tmp6495 := MakeNative(func(__e *ControlFlow) {
W2686 := __e.Get(1)
_ = W2686
tmp6496 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W2685, Nil)
}
__typedArg0 := W2685
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4comb), W2686, tmp6496)
return


}, 1)

tmp6497 := Call(__e, PrimFunc(symshen_4in_1_6), W2684)


__e.TailApply(tmp6495, tmp6497)
return


}, 1)

tmp6498 := Call(__e, PrimFunc(symshen_4_5_1out), W2684)


__e.TailApply(tmp6494, tmp6498)
return


}


}, 1)

tmp6501 := Call(__e, PrimFunc(symshen_4_5digit_6), V2675)


tmp6502 := Call(__e, tmp6493, tmp6501)


__e.TailApply(tmp6490, tmp6502)
return


} else {
__e.Return(W2676)
return
}


}, 1)

tmp6505 := MakeNative(func(__e *ControlFlow) {
W2677 := __e.Get(1)
_ = W2677
tmp6520 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2677)


if True == tmp6520 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6506 := MakeNative(func(__e *ControlFlow) {
W2678 := __e.Get(1)
_ = W2678
tmp6507 := MakeNative(func(__e *ControlFlow) {
W2679 := __e.Get(1)
_ = W2679
tmp6508 := MakeNative(func(__e *ControlFlow) {
W2680 := __e.Get(1)
_ = W2680
tmp6515 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2680)


if True == tmp6515 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6509 := MakeNative(func(__e *ControlFlow) {
W2681 := __e.Get(1)
_ = W2681
tmp6510 := MakeNative(func(__e *ControlFlow) {
W2682 := __e.Get(1)
_ = W2682
tmp6511 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W2678, W2681)
}
__typedArg0 := W2678
__typedArg1 := W2681
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4comb), W2682, tmp6511)
return


}, 1)

tmp6512 := Call(__e, PrimFunc(symshen_4in_1_6), W2680)


__e.TailApply(tmp6510, tmp6512)
return


}, 1)

tmp6513 := Call(__e, PrimFunc(symshen_4_5_1out), W2680)


__e.TailApply(tmp6509, tmp6513)
return


}


}, 1)

tmp6516 := Call(__e, PrimFunc(symshen_4_5digits_6), W2679)


__e.TailApply(tmp6508, tmp6516)
return


}, 1)

tmp6517 := Call(__e, PrimFunc(symshen_4in_1_6), W2677)


__e.TailApply(tmp6507, tmp6517)
return


}, 1)

tmp6518 := Call(__e, PrimFunc(symshen_4_5_1out), W2677)


__e.TailApply(tmp6506, tmp6518)
return


}


}, 1)

tmp6521 := Call(__e, PrimFunc(symshen_4_5digit_6), V2675)


tmp6522 := Call(__e, tmp6505, tmp6521)


__e.TailApply(tmp6489, tmp6522)
return


}, 1)

tmp6523 := Call(__e, ns2_1set, symshen_4_5digits_6, tmp6488)


_ = tmp6523

tmp6524 := MakeNative(func(__e *ControlFlow) {
V2687 := __e.Get(1)
_ = V2687
tmp6525 := MakeNative(func(__e *ControlFlow) {
W2688 := __e.Get(1)
_ = W2688
tmp6527 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2688)


if True == tmp6527 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2688)
return
}


}, 1)

tmp6538 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2687)
}
__typedArg0 := V2687
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres6528 Obj

if True == tmp6538 {
tmp6529 := MakeNative(func(__e *ControlFlow) {
W2689 := __e.Get(1)
_ = W2689
tmp6530 := MakeNative(func(__e *ControlFlow) {
W2690 := __e.Get(1)
_ = W2690
tmp6533 := Call(__e, PrimFunc(symshen_4digit_2), W2689)


if True == tmp6533 {
tmp6531 := Call(__e, PrimFunc(symshen_4byte_1_6digit), W2689)


__e.TailApply(PrimFunc(symshen_4comb), W2690, tmp6531)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp6534 := Call(__e, PrimFunc(symtail), V2687)


__e.TailApply(tmp6530, tmp6534)
return


}, 1)

tmp6535 := Call(__e, PrimFunc(symhead), V2687)


tmp6536 := Call(__e, tmp6529, tmp6535)


ifres6528 = tmp6536


} else {
tmp6537 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6528 = tmp6537


}

__e.TailApply(tmp6525, ifres6528)
return


}, 1)

tmp6539 := Call(__e, ns2_1set, symshen_4_5digit_6, tmp6524)


_ = tmp6539

tmp6540 := MakeNative(func(__e *ControlFlow) {
V2691 := __e.Get(1)
_ = V2691
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(V2691)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(48))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := V2691
__typedArg1 := MakeNumber(48)
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})())
return
}, 1)

tmp6541 := Call(__e, ns2_1set, symshen_4byte_1_6digit, tmp6540)


_ = tmp6541

tmp6542 := MakeNative(func(__e *ControlFlow) {
V2692 := __e.Get(1)
_ = V2692
tmp6543 := Call(__e, PrimFunc(symreverse), V2692)


__e.TailApply(PrimFunc(symshen_4compute_1integer_1h), tmp6543, MakeNumber(0))
return


}, 1)

tmp6544 := Call(__e, ns2_1set, symshen_4compute_1integer, tmp6542)


_ = tmp6544

tmp6545 := MakeNative(func(__e *ControlFlow) {
V2695 := __e.Get(1)
_ = V2695
V2696 := __e.Get(2)
_ = V2696
tmp6555 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V2695)
}
__typedArg0 := Nil
__typedArg1 := V2695
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp6555 {
__e.Return(MakeNumber(0))
return
} else {
tmp6553 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2695)
}
__typedArg0 := V2695
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp6553 {
tmp6546 := Call(__e, PrimFunc(symshen_4expt), MakeNumber(10), V2696)


tmp6547 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2695)
}
__typedArg0 := V2695
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp6548 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_d) {
__typedN0, __typedOK0 := TypedFloat64(tmp6546)
__typedN1, __typedOK1 := TypedFloat64(tmp6547)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_d) {
return TypedMaterializeNumber((__typedN0 * __typedN1))
}}
__typedArg0 := tmp6546
__typedArg1 := tmp6547
return Call(__e, PrimFunc(sym_d), __typedArg0, __typedArg1)
})()

tmp6549 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2695)
}
__typedArg0 := V2695
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6551 := Call(__e, PrimFunc(symshen_4compute_1integer_1h), tmp6549, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V2696)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V2696
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(tmp6548)
__typedN1, __typedOK1 := TypedFloat64(tmp6551)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := tmp6548
__typedArg1 := tmp6551
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("partial function shen.compute-integer-h"))
}
__typedArg0 := MakeString("partial function shen.compute-integer-h")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 2)

tmp6556 := Call(__e, ns2_1set, symshen_4compute_1integer_1h, tmp6545)


_ = tmp6556

tmp6557 := MakeNative(func(__e *ControlFlow) {
V2699 := __e.Get(1)
_ = V2699
V2700 := __e.Get(2)
_ = V2700
tmp6565 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeNumber(0), V2700)
}
__typedArg0 := MakeNumber(0)
__typedArg1 := V2700
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp6565 {
__e.Return(MakeNumber(1))
return
} else {
if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_6) {
__typedN0, __typedOK0 := TypedFloat64(V2700)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(0))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_6) {
return TypedMaterializeBoolean((__typedN0 > __typedN1))
}}
__typedArg0 := V2700
__typedArg1 := MakeNumber(0)
return Call(__e, PrimFunc(sym_6), __typedArg0, __typedArg1)
})() {
tmp6559 := Call(__e, PrimFunc(symshen_4expt), V2699, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(V2700)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := V2700
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})())


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_d) {
__typedN0, __typedOK0 := TypedFloat64(V2699)
__typedN1, __typedOK1 := TypedFloat64(tmp6559)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_d) {
return TypedMaterializeNumber((__typedN0 * __typedN1))
}}
__typedArg0 := V2699
__typedArg1 := tmp6559
return Call(__e, PrimFunc(sym_d), __typedArg0, __typedArg1)
})())
return


} else {
tmp6561 := Call(__e, PrimFunc(symshen_4expt), V2699, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V2700)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V2700
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_c) {
__typedN0, __typedOK0 := TypedFloat64(tmp6561)
__typedN1, __typedOK1 := TypedFloat64(V2699)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_c) {
return TypedMaterializeNumber(TypedDivideValue(__typedN0, __typedN1))
}}
__typedArg0 := tmp6561
__typedArg1 := V2699
return Call(__e, PrimFunc(sym_c), __typedArg0, __typedArg1)
})())
return


}


}


}, 2)

tmp6566 := Call(__e, ns2_1set, symshen_4expt, tmp6557)


_ = tmp6566

tmp6567 := MakeNative(func(__e *ControlFlow) {
V2701 := __e.Get(1)
_ = V2701
tmp6568 := MakeNative(func(__e *ControlFlow) {
W2702 := __e.Get(1)
_ = W2702
tmp6588 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2702)


if True == tmp6588 {
tmp6569 := MakeNative(func(__e *ControlFlow) {
W2711 := __e.Get(1)
_ = W2711
tmp6571 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2711)


if True == tmp6571 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2711)
return
}


}, 1)

tmp6572 := MakeNative(func(__e *ControlFlow) {
W2712 := __e.Get(1)
_ = W2712
tmp6584 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2712)


if True == tmp6584 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6573 := MakeNative(func(__e *ControlFlow) {
W2713 := __e.Get(1)
_ = W2713
tmp6574 := MakeNative(func(__e *ControlFlow) {
W2714 := __e.Get(1)
_ = W2714
tmp6580 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2714)


if True == tmp6580 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6575 := MakeNative(func(__e *ControlFlow) {
W2715 := __e.Get(1)
_ = W2715
tmp6576 := MakeNative(func(__e *ControlFlow) {
W2716 := __e.Get(1)
_ = W2716
__e.TailApply(PrimFunc(symshen_4comb), W2716, W2715)
return
}, 1)

tmp6577 := Call(__e, PrimFunc(symshen_4in_1_6), W2714)


__e.TailApply(tmp6576, tmp6577)
return


}, 1)

tmp6578 := Call(__e, PrimFunc(symshen_4_5_1out), W2714)


__e.TailApply(tmp6575, tmp6578)
return


}


}, 1)

tmp6581 := Call(__e, PrimFunc(symshen_4_5fraction_6), W2713)


__e.TailApply(tmp6574, tmp6581)
return


}, 1)

tmp6582 := Call(__e, PrimFunc(symshen_4in_1_6), W2712)


__e.TailApply(tmp6573, tmp6582)
return


}


}, 1)

tmp6585 := Call(__e, PrimFunc(symshen_4_5stop_6), V2701)


tmp6586 := Call(__e, tmp6572, tmp6585)


__e.TailApply(tmp6569, tmp6586)
return


} else {
__e.Return(W2702)
return
}


}, 1)

tmp6589 := MakeNative(func(__e *ControlFlow) {
W2703 := __e.Get(1)
_ = W2703
tmp6610 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2703)


if True == tmp6610 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6590 := MakeNative(func(__e *ControlFlow) {
W2704 := __e.Get(1)
_ = W2704
tmp6591 := MakeNative(func(__e *ControlFlow) {
W2705 := __e.Get(1)
_ = W2705
tmp6592 := MakeNative(func(__e *ControlFlow) {
W2706 := __e.Get(1)
_ = W2706
tmp6605 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2706)


if True == tmp6605 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6593 := MakeNative(func(__e *ControlFlow) {
W2707 := __e.Get(1)
_ = W2707
tmp6594 := MakeNative(func(__e *ControlFlow) {
W2708 := __e.Get(1)
_ = W2708
tmp6601 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2708)


if True == tmp6601 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6595 := MakeNative(func(__e *ControlFlow) {
W2709 := __e.Get(1)
_ = W2709
tmp6596 := MakeNative(func(__e *ControlFlow) {
W2710 := __e.Get(1)
_ = W2710
__e.TailApply(PrimFunc(symshen_4comb), W2710, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(W2704)
__typedN1, __typedOK1 := TypedFloat64(W2709)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := W2704
__typedArg1 := W2709
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp6598 := Call(__e, PrimFunc(symshen_4in_1_6), W2708)


__e.TailApply(tmp6596, tmp6598)
return


}, 1)

tmp6599 := Call(__e, PrimFunc(symshen_4_5_1out), W2708)


__e.TailApply(tmp6595, tmp6599)
return


}


}, 1)

tmp6602 := Call(__e, PrimFunc(symshen_4_5fraction_6), W2707)


__e.TailApply(tmp6594, tmp6602)
return


}, 1)

tmp6603 := Call(__e, PrimFunc(symshen_4in_1_6), W2706)


__e.TailApply(tmp6593, tmp6603)
return


}


}, 1)

tmp6606 := Call(__e, PrimFunc(symshen_4_5stop_6), W2705)


__e.TailApply(tmp6592, tmp6606)
return


}, 1)

tmp6607 := Call(__e, PrimFunc(symshen_4in_1_6), W2703)


__e.TailApply(tmp6591, tmp6607)
return


}, 1)

tmp6608 := Call(__e, PrimFunc(symshen_4_5_1out), W2703)


__e.TailApply(tmp6590, tmp6608)
return


}


}, 1)

tmp6611 := Call(__e, PrimFunc(symshen_4_5integer_6), V2701)


tmp6612 := Call(__e, tmp6589, tmp6611)


__e.TailApply(tmp6568, tmp6612)
return


}, 1)

tmp6613 := Call(__e, ns2_1set, symshen_4_5float_6, tmp6567)


_ = tmp6613

tmp6614 := MakeNative(func(__e *ControlFlow) {
V2717 := __e.Get(1)
_ = V2717
tmp6615 := MakeNative(func(__e *ControlFlow) {
W2718 := __e.Get(1)
_ = W2718
tmp6617 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2718)


if True == tmp6617 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2718)
return
}


}, 1)

tmp6623 := Call(__e, PrimFunc(symshen_4hds_a_2), V2717, MakeNumber(46))


var ifres6618 Obj

if True == tmp6623 {
tmp6619 := MakeNative(func(__e *ControlFlow) {
W2719 := __e.Get(1)
_ = W2719
__e.TailApply(PrimFunc(symshen_4comb), W2719, symshen_4skip)
return
}, 1)

tmp6620 := Call(__e, PrimFunc(symtail), V2717)


tmp6621 := Call(__e, tmp6619, tmp6620)


ifres6618 = tmp6621


} else {
tmp6622 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6618 = tmp6622


}

__e.TailApply(tmp6615, ifres6618)
return


}, 1)

tmp6624 := Call(__e, ns2_1set, symshen_4_5stop_6, tmp6614)


_ = tmp6624

tmp6625 := MakeNative(func(__e *ControlFlow) {
V2720 := __e.Get(1)
_ = V2720
tmp6626 := MakeNative(func(__e *ControlFlow) {
W2721 := __e.Get(1)
_ = W2721
tmp6628 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2721)


if True == tmp6628 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2721)
return
}


}, 1)

tmp6629 := MakeNative(func(__e *ControlFlow) {
W2722 := __e.Get(1)
_ = W2722
tmp6636 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2722)


if True == tmp6636 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6630 := MakeNative(func(__e *ControlFlow) {
W2723 := __e.Get(1)
_ = W2723
tmp6631 := MakeNative(func(__e *ControlFlow) {
W2724 := __e.Get(1)
_ = W2724
tmp6632 := Call(__e, PrimFunc(symshen_4compute_1fraction), W2723)


__e.TailApply(PrimFunc(symshen_4comb), W2724, tmp6632)
return


}, 1)

tmp6633 := Call(__e, PrimFunc(symshen_4in_1_6), W2722)


__e.TailApply(tmp6631, tmp6633)
return


}, 1)

tmp6634 := Call(__e, PrimFunc(symshen_4_5_1out), W2722)


__e.TailApply(tmp6630, tmp6634)
return


}


}, 1)

tmp6637 := Call(__e, PrimFunc(symshen_4_5digits_6), V2720)


tmp6638 := Call(__e, tmp6629, tmp6637)


__e.TailApply(tmp6626, tmp6638)
return


}, 1)

tmp6639 := Call(__e, ns2_1set, symshen_4_5fraction_6, tmp6625)


_ = tmp6639

tmp6640 := MakeNative(func(__e *ControlFlow) {
V2725 := __e.Get(1)
_ = V2725
__e.TailApply(PrimFunc(symshen_4compute_1fraction_1h), V2725, MakeNumber(-1))
return
}, 1)

tmp6641 := Call(__e, ns2_1set, symshen_4compute_1fraction, tmp6640)


_ = tmp6641

tmp6642 := MakeNative(func(__e *ControlFlow) {
V2728 := __e.Get(1)
_ = V2728
V2729 := __e.Get(2)
_ = V2729
tmp6652 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V2728)
}
__typedArg0 := Nil
__typedArg1 := V2728
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp6652 {
__e.Return(MakeNumber(0))
return
} else {
tmp6650 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2728)
}
__typedArg0 := V2728
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp6650 {
tmp6643 := Call(__e, PrimFunc(symshen_4expt), MakeNumber(10), V2729)


tmp6644 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2728)
}
__typedArg0 := V2728
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp6645 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_d) {
__typedN0, __typedOK0 := TypedFloat64(tmp6643)
__typedN1, __typedOK1 := TypedFloat64(tmp6644)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_d) {
return TypedMaterializeNumber((__typedN0 * __typedN1))
}}
__typedArg0 := tmp6643
__typedArg1 := tmp6644
return Call(__e, PrimFunc(sym_d), __typedArg0, __typedArg1)
})()

tmp6646 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2728)
}
__typedArg0 := V2728
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6648 := Call(__e, PrimFunc(symshen_4compute_1fraction_1h), tmp6646, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(V2729)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := V2729
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})())


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(tmp6645)
__typedN1, __typedOK1 := TypedFloat64(tmp6648)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := tmp6645
__typedArg1 := tmp6648
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("partial function shen.compute-fraction-h"))
}
__typedArg0 := MakeString("partial function shen.compute-fraction-h")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 2)

tmp6653 := Call(__e, ns2_1set, symshen_4compute_1fraction_1h, tmp6642)


_ = tmp6653

tmp6654 := MakeNative(func(__e *ControlFlow) {
V2730 := __e.Get(1)
_ = V2730
tmp6655 := MakeNative(func(__e *ControlFlow) {
W2731 := __e.Get(1)
_ = W2731
tmp6684 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2731)


if True == tmp6684 {
tmp6656 := MakeNative(func(__e *ControlFlow) {
W2740 := __e.Get(1)
_ = W2740
tmp6658 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2740)


if True == tmp6658 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2740)
return
}


}, 1)

tmp6659 := MakeNative(func(__e *ControlFlow) {
W2741 := __e.Get(1)
_ = W2741
tmp6680 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2741)


if True == tmp6680 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6660 := MakeNative(func(__e *ControlFlow) {
W2742 := __e.Get(1)
_ = W2742
tmp6661 := MakeNative(func(__e *ControlFlow) {
W2743 := __e.Get(1)
_ = W2743
tmp6662 := MakeNative(func(__e *ControlFlow) {
W2744 := __e.Get(1)
_ = W2744
tmp6675 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2744)


if True == tmp6675 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6663 := MakeNative(func(__e *ControlFlow) {
W2745 := __e.Get(1)
_ = W2745
tmp6664 := MakeNative(func(__e *ControlFlow) {
W2746 := __e.Get(1)
_ = W2746
tmp6671 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2746)


if True == tmp6671 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6665 := MakeNative(func(__e *ControlFlow) {
W2747 := __e.Get(1)
_ = W2747
tmp6666 := MakeNative(func(__e *ControlFlow) {
W2748 := __e.Get(1)
_ = W2748
tmp6667 := Call(__e, PrimFunc(symshen_4compute_1E), W2742, W2747)


__e.TailApply(PrimFunc(symshen_4comb), W2748, tmp6667)
return


}, 1)

tmp6668 := Call(__e, PrimFunc(symshen_4in_1_6), W2746)


__e.TailApply(tmp6666, tmp6668)
return


}, 1)

tmp6669 := Call(__e, PrimFunc(symshen_4_5_1out), W2746)


__e.TailApply(tmp6665, tmp6669)
return


}


}, 1)

tmp6672 := Call(__e, PrimFunc(symshen_4_5log10_6), W2745)


__e.TailApply(tmp6664, tmp6672)
return


}, 1)

tmp6673 := Call(__e, PrimFunc(symshen_4in_1_6), W2744)


__e.TailApply(tmp6663, tmp6673)
return


}


}, 1)

tmp6676 := Call(__e, PrimFunc(symshen_4_5lowE_6), W2743)


__e.TailApply(tmp6662, tmp6676)
return


}, 1)

tmp6677 := Call(__e, PrimFunc(symshen_4in_1_6), W2741)


__e.TailApply(tmp6661, tmp6677)
return


}, 1)

tmp6678 := Call(__e, PrimFunc(symshen_4_5_1out), W2741)


__e.TailApply(tmp6660, tmp6678)
return


}


}, 1)

tmp6681 := Call(__e, PrimFunc(symshen_4_5integer_6), V2730)


tmp6682 := Call(__e, tmp6659, tmp6681)


__e.TailApply(tmp6656, tmp6682)
return


} else {
__e.Return(W2731)
return
}


}, 1)

tmp6685 := MakeNative(func(__e *ControlFlow) {
W2732 := __e.Get(1)
_ = W2732
tmp6706 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2732)


if True == tmp6706 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6686 := MakeNative(func(__e *ControlFlow) {
W2733 := __e.Get(1)
_ = W2733
tmp6687 := MakeNative(func(__e *ControlFlow) {
W2734 := __e.Get(1)
_ = W2734
tmp6688 := MakeNative(func(__e *ControlFlow) {
W2735 := __e.Get(1)
_ = W2735
tmp6701 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2735)


if True == tmp6701 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6689 := MakeNative(func(__e *ControlFlow) {
W2736 := __e.Get(1)
_ = W2736
tmp6690 := MakeNative(func(__e *ControlFlow) {
W2737 := __e.Get(1)
_ = W2737
tmp6697 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2737)


if True == tmp6697 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6691 := MakeNative(func(__e *ControlFlow) {
W2738 := __e.Get(1)
_ = W2738
tmp6692 := MakeNative(func(__e *ControlFlow) {
W2739 := __e.Get(1)
_ = W2739
tmp6693 := Call(__e, PrimFunc(symshen_4compute_1E), W2733, W2738)


__e.TailApply(PrimFunc(symshen_4comb), W2739, tmp6693)
return


}, 1)

tmp6694 := Call(__e, PrimFunc(symshen_4in_1_6), W2737)


__e.TailApply(tmp6692, tmp6694)
return


}, 1)

tmp6695 := Call(__e, PrimFunc(symshen_4_5_1out), W2737)


__e.TailApply(tmp6691, tmp6695)
return


}


}, 1)

tmp6698 := Call(__e, PrimFunc(symshen_4_5log10_6), W2736)


__e.TailApply(tmp6690, tmp6698)
return


}, 1)

tmp6699 := Call(__e, PrimFunc(symshen_4in_1_6), W2735)


__e.TailApply(tmp6689, tmp6699)
return


}


}, 1)

tmp6702 := Call(__e, PrimFunc(symshen_4_5lowE_6), W2734)


__e.TailApply(tmp6688, tmp6702)
return


}, 1)

tmp6703 := Call(__e, PrimFunc(symshen_4in_1_6), W2732)


__e.TailApply(tmp6687, tmp6703)
return


}, 1)

tmp6704 := Call(__e, PrimFunc(symshen_4_5_1out), W2732)


__e.TailApply(tmp6686, tmp6704)
return


}


}, 1)

tmp6707 := Call(__e, PrimFunc(symshen_4_5float_6), V2730)


tmp6708 := Call(__e, tmp6685, tmp6707)


__e.TailApply(tmp6655, tmp6708)
return


}, 1)

tmp6709 := Call(__e, ns2_1set, symshen_4_5e_1number_6, tmp6654)


_ = tmp6709

tmp6710 := MakeNative(func(__e *ControlFlow) {
V2749 := __e.Get(1)
_ = V2749
tmp6711 := MakeNative(func(__e *ControlFlow) {
W2750 := __e.Get(1)
_ = W2750
tmp6744 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2750)


if True == tmp6744 {
tmp6712 := MakeNative(func(__e *ControlFlow) {
W2756 := __e.Get(1)
_ = W2756
tmp6726 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2756)


if True == tmp6726 {
tmp6713 := MakeNative(func(__e *ControlFlow) {
W2762 := __e.Get(1)
_ = W2762
tmp6715 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2762)


if True == tmp6715 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2762)
return
}


}, 1)

tmp6716 := MakeNative(func(__e *ControlFlow) {
W2763 := __e.Get(1)
_ = W2763
tmp6722 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2763)


if True == tmp6722 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6717 := MakeNative(func(__e *ControlFlow) {
W2764 := __e.Get(1)
_ = W2764
tmp6718 := MakeNative(func(__e *ControlFlow) {
W2765 := __e.Get(1)
_ = W2765
__e.TailApply(PrimFunc(symshen_4comb), W2765, W2764)
return
}, 1)

tmp6719 := Call(__e, PrimFunc(symshen_4in_1_6), W2763)


__e.TailApply(tmp6718, tmp6719)
return


}, 1)

tmp6720 := Call(__e, PrimFunc(symshen_4_5_1out), W2763)


__e.TailApply(tmp6717, tmp6720)
return


}


}, 1)

tmp6723 := Call(__e, PrimFunc(symshen_4_5integer_6), V2749)


tmp6724 := Call(__e, tmp6716, tmp6723)


__e.TailApply(tmp6713, tmp6724)
return


} else {
__e.Return(W2756)
return
}


}, 1)

tmp6727 := MakeNative(func(__e *ControlFlow) {
W2757 := __e.Get(1)
_ = W2757
tmp6740 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2757)


if True == tmp6740 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6728 := MakeNative(func(__e *ControlFlow) {
W2758 := __e.Get(1)
_ = W2758
tmp6729 := MakeNative(func(__e *ControlFlow) {
W2759 := __e.Get(1)
_ = W2759
tmp6736 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2759)


if True == tmp6736 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6730 := MakeNative(func(__e *ControlFlow) {
W2760 := __e.Get(1)
_ = W2760
tmp6731 := MakeNative(func(__e *ControlFlow) {
W2761 := __e.Get(1)
_ = W2761
__e.TailApply(PrimFunc(symshen_4comb), W2761, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(0))
__typedN1, __typedOK1 := TypedFloat64(W2760)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := MakeNumber(0)
__typedArg1 := W2760
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp6733 := Call(__e, PrimFunc(symshen_4in_1_6), W2759)


__e.TailApply(tmp6731, tmp6733)
return


}, 1)

tmp6734 := Call(__e, PrimFunc(symshen_4_5_1out), W2759)


__e.TailApply(tmp6730, tmp6734)
return


}


}, 1)

tmp6737 := Call(__e, PrimFunc(symshen_4_5log10_6), W2758)


__e.TailApply(tmp6729, tmp6737)
return


}, 1)

tmp6738 := Call(__e, PrimFunc(symshen_4in_1_6), W2757)


__e.TailApply(tmp6728, tmp6738)
return


}


}, 1)

tmp6741 := Call(__e, PrimFunc(symshen_4_5minus_6), V2749)


tmp6742 := Call(__e, tmp6727, tmp6741)


__e.TailApply(tmp6712, tmp6742)
return


} else {
__e.Return(W2750)
return
}


}, 1)

tmp6745 := MakeNative(func(__e *ControlFlow) {
W2751 := __e.Get(1)
_ = W2751
tmp6757 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2751)


if True == tmp6757 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6746 := MakeNative(func(__e *ControlFlow) {
W2752 := __e.Get(1)
_ = W2752
tmp6747 := MakeNative(func(__e *ControlFlow) {
W2753 := __e.Get(1)
_ = W2753
tmp6753 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2753)


if True == tmp6753 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6748 := MakeNative(func(__e *ControlFlow) {
W2754 := __e.Get(1)
_ = W2754
tmp6749 := MakeNative(func(__e *ControlFlow) {
W2755 := __e.Get(1)
_ = W2755
__e.TailApply(PrimFunc(symshen_4comb), W2755, W2754)
return
}, 1)

tmp6750 := Call(__e, PrimFunc(symshen_4in_1_6), W2753)


__e.TailApply(tmp6749, tmp6750)
return


}, 1)

tmp6751 := Call(__e, PrimFunc(symshen_4_5_1out), W2753)


__e.TailApply(tmp6748, tmp6751)
return


}


}, 1)

tmp6754 := Call(__e, PrimFunc(symshen_4_5log10_6), W2752)


__e.TailApply(tmp6747, tmp6754)
return


}, 1)

tmp6755 := Call(__e, PrimFunc(symshen_4in_1_6), W2751)


__e.TailApply(tmp6746, tmp6755)
return


}


}, 1)

tmp6758 := Call(__e, PrimFunc(symshen_4_5plus_6), V2749)


tmp6759 := Call(__e, tmp6745, tmp6758)


__e.TailApply(tmp6711, tmp6759)
return


}, 1)

tmp6760 := Call(__e, ns2_1set, symshen_4_5log10_6, tmp6710)


_ = tmp6760

tmp6761 := MakeNative(func(__e *ControlFlow) {
V2766 := __e.Get(1)
_ = V2766
tmp6762 := MakeNative(func(__e *ControlFlow) {
W2767 := __e.Get(1)
_ = W2767
tmp6764 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2767)


if True == tmp6764 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2767)
return
}


}, 1)

tmp6770 := Call(__e, PrimFunc(symshen_4hds_a_2), V2766, MakeNumber(101))


var ifres6765 Obj

if True == tmp6770 {
tmp6766 := MakeNative(func(__e *ControlFlow) {
W2768 := __e.Get(1)
_ = W2768
__e.TailApply(PrimFunc(symshen_4comb), W2768, symshen_4skip)
return
}, 1)

tmp6767 := Call(__e, PrimFunc(symtail), V2766)


tmp6768 := Call(__e, tmp6766, tmp6767)


ifres6765 = tmp6768


} else {
tmp6769 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6765 = tmp6769


}

__e.TailApply(tmp6762, ifres6765)
return


}, 1)

tmp6771 := Call(__e, ns2_1set, symshen_4_5lowE_6, tmp6761)


_ = tmp6771

tmp6772 := MakeNative(func(__e *ControlFlow) {
V2769 := __e.Get(1)
_ = V2769
V2770 := __e.Get(2)
_ = V2770
tmp6773 := Call(__e, PrimFunc(symshen_4expt), MakeNumber(10), V2770)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_d) {
__typedN0, __typedOK0 := TypedFloat64(V2769)
__typedN1, __typedOK1 := TypedFloat64(tmp6773)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_d) {
return TypedMaterializeNumber((__typedN0 * __typedN1))
}}
__typedArg0 := V2769
__typedArg1 := tmp6773
return Call(__e, PrimFunc(sym_d), __typedArg0, __typedArg1)
})())
return


}, 2)

tmp6774 := Call(__e, ns2_1set, symshen_4compute_1E, tmp6772)


_ = tmp6774

tmp6775 := MakeNative(func(__e *ControlFlow) {
V2771 := __e.Get(1)
_ = V2771
tmp6776 := MakeNative(func(__e *ControlFlow) {
W2772 := __e.Get(1)
_ = W2772
tmp6788 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2772)


if True == tmp6788 {
tmp6777 := MakeNative(func(__e *ControlFlow) {
W2777 := __e.Get(1)
_ = W2777
tmp6779 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2777)


if True == tmp6779 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2777)
return
}


}, 1)

tmp6780 := MakeNative(func(__e *ControlFlow) {
W2778 := __e.Get(1)
_ = W2778
tmp6784 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2778)


if True == tmp6784 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6781 := MakeNative(func(__e *ControlFlow) {
W2779 := __e.Get(1)
_ = W2779
__e.TailApply(PrimFunc(symshen_4comb), W2779, symshen_4skip)
return
}, 1)

tmp6782 := Call(__e, PrimFunc(symshen_4in_1_6), W2778)


__e.TailApply(tmp6781, tmp6782)
return


}


}, 1)

tmp6785 := Call(__e, PrimFunc(symshen_4_5whitespace_6), V2771)


tmp6786 := Call(__e, tmp6780, tmp6785)


__e.TailApply(tmp6777, tmp6786)
return


} else {
__e.Return(W2772)
return
}


}, 1)

tmp6789 := MakeNative(func(__e *ControlFlow) {
W2773 := __e.Get(1)
_ = W2773
tmp6799 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2773)


if True == tmp6799 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6790 := MakeNative(func(__e *ControlFlow) {
W2774 := __e.Get(1)
_ = W2774
tmp6791 := MakeNative(func(__e *ControlFlow) {
W2775 := __e.Get(1)
_ = W2775
tmp6795 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2775)


if True == tmp6795 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6792 := MakeNative(func(__e *ControlFlow) {
W2776 := __e.Get(1)
_ = W2776
__e.TailApply(PrimFunc(symshen_4comb), W2776, symshen_4skip)
return
}, 1)

tmp6793 := Call(__e, PrimFunc(symshen_4in_1_6), W2775)


__e.TailApply(tmp6792, tmp6793)
return


}


}, 1)

tmp6796 := Call(__e, PrimFunc(symshen_4_5whitespaces_6), W2774)


__e.TailApply(tmp6791, tmp6796)
return


}, 1)

tmp6797 := Call(__e, PrimFunc(symshen_4in_1_6), W2773)


__e.TailApply(tmp6790, tmp6797)
return


}


}, 1)

tmp6800 := Call(__e, PrimFunc(symshen_4_5whitespace_6), V2771)


tmp6801 := Call(__e, tmp6789, tmp6800)


__e.TailApply(tmp6776, tmp6801)
return


}, 1)

tmp6802 := Call(__e, ns2_1set, symshen_4_5whitespaces_6, tmp6775)


_ = tmp6802

tmp6803 := MakeNative(func(__e *ControlFlow) {
V2780 := __e.Get(1)
_ = V2780
tmp6804 := MakeNative(func(__e *ControlFlow) {
W2781 := __e.Get(1)
_ = W2781
tmp6806 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2781)


if True == tmp6806 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2781)
return
}


}, 1)

tmp6816 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2780)
}
__typedArg0 := V2780
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres6807 Obj

if True == tmp6816 {
tmp6808 := MakeNative(func(__e *ControlFlow) {
W2782 := __e.Get(1)
_ = W2782
tmp6809 := MakeNative(func(__e *ControlFlow) {
W2783 := __e.Get(1)
_ = W2783
tmp6811 := Call(__e, PrimFunc(symshen_4whitespace_2), W2782)


if True == tmp6811 {
__e.TailApply(PrimFunc(symshen_4comb), W2783, symshen_4skip)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp6812 := Call(__e, PrimFunc(symtail), V2780)


__e.TailApply(tmp6809, tmp6812)
return


}, 1)

tmp6813 := Call(__e, PrimFunc(symhead), V2780)


tmp6814 := Call(__e, tmp6808, tmp6813)


ifres6807 = tmp6814


} else {
tmp6815 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6807 = tmp6815


}

__e.TailApply(tmp6804, ifres6807)
return


}, 1)

tmp6817 := Call(__e, ns2_1set, symshen_4_5whitespace_6, tmp6803)


_ = tmp6817

tmp6818 := MakeNative(func(__e *ControlFlow) {
V2786 := __e.Get(1)
_ = V2786
tmp6826 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeNumber(32), V2786)
}
__typedArg0 := MakeNumber(32)
__typedArg1 := V2786
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp6826 {
__e.Return(True)
return
} else {
tmp6824 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeNumber(13), V2786)
}
__typedArg0 := MakeNumber(13)
__typedArg1 := V2786
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp6824 {
__e.Return(True)
return
} else {
tmp6822 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeNumber(10), V2786)
}
__typedArg0 := MakeNumber(10)
__typedArg1 := V2786
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp6822 {
__e.Return(True)
return
} else {
tmp6820 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeNumber(9), V2786)
}
__typedArg0 := MakeNumber(9)
__typedArg1 := V2786
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp6820 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


}


}


}, 1)

tmp6827 := Call(__e, ns2_1set, symshen_4whitespace_2, tmp6818)


_ = tmp6827

tmp6828 := MakeNative(func(__e *ControlFlow) {
V2787 := __e.Get(1)
_ = V2787
tmp6851 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V2787)
}
__typedArg0 := Nil
__typedArg1 := V2787
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp6851 {
__e.Return(Nil)
return
} else {
tmp6849 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2787)
}
__typedArg0 := V2787
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres6845 Obj

if True == tmp6849 {
tmp6847 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2787)
}
__typedArg0 := V2787
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp6848 := Call(__e, PrimFunc(symshen_4packaged_2), tmp6847)


var ifres6846 Obj

if True == tmp6848 {
ifres6846 = True


} else {
ifres6846 = False


}

ifres6845 = ifres6846


} else {
ifres6845 = False


}

if True == ifres6845 {
tmp6829 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2787)
}
__typedArg0 := V2787
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp6830 := Call(__e, PrimFunc(symshen_4unpackage), tmp6829)


tmp6831 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2787)
}
__typedArg0 := V2787
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6832 := Call(__e, PrimFunc(symappend), tmp6830, tmp6831)


__e.TailApply(PrimFunc(symshen_4unpackage_emacroexpand), tmp6832)
return


} else {
tmp6843 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2787)
}
__typedArg0 := V2787
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp6843 {
tmp6833 := MakeNative(func(__e *ControlFlow) {
W2788 := __e.Get(1)
_ = W2788
tmp6839 := Call(__e, PrimFunc(symshen_4packaged_2), W2788)


if True == tmp6839 {
tmp6834 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2787)
}
__typedArg0 := V2787
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6835 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W2788, tmp6834)
}
__typedArg0 := W2788
__typedArg1 := tmp6834
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4unpackage_emacroexpand), tmp6835)
return


} else {
tmp6836 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2787)
}
__typedArg0 := V2787
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6837 := Call(__e, PrimFunc(symshen_4unpackage_emacroexpand), tmp6836)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W2788, tmp6837)
}
__typedArg0 := W2788
__typedArg1 := tmp6837
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


}, 1)

tmp6840 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2787)
}
__typedArg0 := V2787
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp6841 := Call(__e, PrimFunc(symmacroexpand), tmp6840)


__e.TailApply(tmp6833, tmp6841)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("partial function shen.unpackage&macroexpand"))
}
__typedArg0 := MakeString("partial function shen.unpackage&macroexpand")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}


}, 1)

tmp6852 := Call(__e, ns2_1set, symshen_4unpackage_emacroexpand, tmp6828)


_ = tmp6852

tmp6853 := MakeNative(func(__e *ControlFlow) {
V2791 := __e.Get(1)
_ = V2791
tmp6868 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2791)
}
__typedArg0 := V2791
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres6855 Obj

if True == tmp6868 {
tmp6866 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2791)
}
__typedArg0 := V2791
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp6867 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sympackage, tmp6866)
}
__typedArg0 := sympackage
__typedArg1 := tmp6866
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres6857 Obj

if True == tmp6867 {
tmp6864 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2791)
}
__typedArg0 := V2791
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6865 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp6864)
}
__typedArg0 := tmp6864
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres6859 Obj

if True == tmp6865 {
tmp6861 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2791)
}
__typedArg0 := V2791
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6862 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp6861)
}
__typedArg0 := tmp6861
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6863 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp6862)
}
__typedArg0 := tmp6862
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres6860 Obj

if True == tmp6863 {
ifres6860 = True


} else {
ifres6860 = False


}

ifres6859 = ifres6860


} else {
ifres6859 = False


}

var ifres6858 Obj

if True == ifres6859 {
ifres6858 = True


} else {
ifres6858 = False


}

ifres6857 = ifres6858


} else {
ifres6857 = False


}

var ifres6856 Obj

if True == ifres6857 {
ifres6856 = True


} else {
ifres6856 = False


}

ifres6855 = ifres6856


} else {
ifres6855 = False


}

if True == ifres6855 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 1)

tmp6869 := Call(__e, ns2_1set, symshen_4packaged_2, tmp6853)


_ = tmp6869

tmp6870 := MakeNative(func(__e *ControlFlow) {
V2794 := __e.Get(1)
_ = V2794
tmp6931 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2794)
}
__typedArg0 := V2794
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres6913 Obj

if True == tmp6931 {
tmp6929 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2794)
}
__typedArg0 := V2794
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp6930 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sympackage, tmp6929)
}
__typedArg0 := sympackage
__typedArg1 := tmp6929
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres6915 Obj

if True == tmp6930 {
tmp6927 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2794)
}
__typedArg0 := V2794
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6928 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp6927)
}
__typedArg0 := tmp6927
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres6917 Obj

if True == tmp6928 {
tmp6924 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2794)
}
__typedArg0 := V2794
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6925 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp6924)
}
__typedArg0 := tmp6924
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp6926 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symnull, tmp6925)
}
__typedArg0 := symnull
__typedArg1 := tmp6925
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres6919 Obj

if True == tmp6926 {
tmp6921 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2794)
}
__typedArg0 := V2794
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6922 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp6921)
}
__typedArg0 := tmp6921
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6923 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp6922)
}
__typedArg0 := tmp6922
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres6920 Obj

if True == tmp6923 {
ifres6920 = True


} else {
ifres6920 = False


}

ifres6919 = ifres6920


} else {
ifres6919 = False


}

var ifres6918 Obj

if True == ifres6919 {
ifres6918 = True


} else {
ifres6918 = False


}

ifres6917 = ifres6918


} else {
ifres6917 = False


}

var ifres6916 Obj

if True == ifres6917 {
ifres6916 = True


} else {
ifres6916 = False


}

ifres6915 = ifres6916


} else {
ifres6915 = False


}

var ifres6914 Obj

if True == ifres6915 {
ifres6914 = True


} else {
ifres6914 = False


}

ifres6913 = ifres6914


} else {
ifres6913 = False


}

if True == ifres6913 {
tmp6871 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2794)
}
__typedArg0 := V2794
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6872 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp6871)
}
__typedArg0 := tmp6871
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp6872)
}
__typedArg0 := tmp6872
return Call(__e, PrimFunc(symtl), __typedArg0)
})())
return


} else {
tmp6911 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2794)
}
__typedArg0 := V2794
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres6898 Obj

if True == tmp6911 {
tmp6909 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2794)
}
__typedArg0 := V2794
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp6910 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sympackage, tmp6909)
}
__typedArg0 := sympackage
__typedArg1 := tmp6909
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres6900 Obj

if True == tmp6910 {
tmp6907 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2794)
}
__typedArg0 := V2794
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6908 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp6907)
}
__typedArg0 := tmp6907
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres6902 Obj

if True == tmp6908 {
tmp6904 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2794)
}
__typedArg0 := V2794
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6905 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp6904)
}
__typedArg0 := tmp6904
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6906 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp6905)
}
__typedArg0 := tmp6905
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres6903 Obj

if True == tmp6906 {
ifres6903 = True


} else {
ifres6903 = False


}

ifres6902 = ifres6903


} else {
ifres6902 = False


}

var ifres6901 Obj

if True == ifres6902 {
ifres6901 = True


} else {
ifres6901 = False


}

ifres6900 = ifres6901


} else {
ifres6900 = False


}

var ifres6899 Obj

if True == ifres6900 {
ifres6899 = True


} else {
ifres6899 = False


}

ifres6898 = ifres6899


} else {
ifres6898 = False


}

if True == ifres6898 {
tmp6873 := MakeNative(func(__e *ControlFlow) {
W2795 := __e.Get(1)
_ = W2795
tmp6874 := MakeNative(func(__e *ControlFlow) {
W2796 := __e.Get(1)
_ = W2796
tmp6875 := MakeNative(func(__e *ControlFlow) {
W2797 := __e.Get(1)
_ = W2797
tmp6876 := MakeNative(func(__e *ControlFlow) {
W2798 := __e.Get(1)
_ = W2798
__e.Return(W2796)
return
}, 1)

tmp6877 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2794)
}
__typedArg0 := V2794
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6878 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp6877)
}
__typedArg0 := tmp6877
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp6879 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2794)
}
__typedArg0 := V2794
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6880 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp6879)
}
__typedArg0 := tmp6879
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6881 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp6880)
}
__typedArg0 := tmp6880
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6882 := Call(__e, PrimFunc(symshen_4record_1internal), tmp6878, W2795, tmp6881)


__e.TailApply(tmp6876, tmp6882)
return


}, 1)

tmp6883 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2794)
}
__typedArg0 := V2794
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6884 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp6883)
}
__typedArg0 := tmp6883
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp6885 := Call(__e, PrimFunc(symshen_4record_1external), tmp6884, W2795)


__e.TailApply(tmp6875, tmp6885)
return


}, 1)

tmp6886 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2794)
}
__typedArg0 := V2794
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6887 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp6886)
}
__typedArg0 := tmp6886
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp6888 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(tmp6887)
}
__typedArg0 := tmp6887
return Call(__e, PrimFunc(symstr), __typedArg0)
})()

tmp6889 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2794)
}
__typedArg0 := V2794
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6890 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp6889)
}
__typedArg0 := tmp6889
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6891 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp6890)
}
__typedArg0 := tmp6890
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6892 := Call(__e, PrimFunc(symshen_4package_1symbols), tmp6888, W2795, tmp6891)


__e.TailApply(tmp6874, tmp6892)
return


}, 1)

tmp6893 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2794)
}
__typedArg0 := V2794
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6894 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp6893)
}
__typedArg0 := tmp6893
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6895 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp6894)
}
__typedArg0 := tmp6894
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp6896 := Call(__e, PrimFunc(symeval), tmp6895)


__e.TailApply(tmp6873, tmp6896)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("partial function shen.unpackage"))
}
__typedArg0 := MakeString("partial function shen.unpackage")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 1)

tmp6932 := Call(__e, ns2_1set, symshen_4unpackage, tmp6870)


_ = tmp6932

tmp6933 := MakeNative(func(__e *ControlFlow) {
V2799 := __e.Get(1)
_ = V2799
V2800 := __e.Get(2)
_ = V2800
V2801 := __e.Get(3)
_ = V2801
tmp6934 := MakeNative(func(__e *ControlFlow) {
W2802 := __e.Get(1)
_ = W2802
tmp6935 := MakeNative(func(__e *ControlFlow) {
W2804 := __e.Get(1)
_ = W2804
tmp6936 := Call(__e, PrimFunc(symunion), W2804, W2802)


tmp6937 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symput), V2799, symshen_4internal_1symbols, tmp6936, tmp6937)
return


}, 1)

tmp6938 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(V2799)
}
__typedArg0 := V2799
return Call(__e, PrimFunc(symstr), __typedArg0)
})()

tmp6939 := Call(__e, PrimFunc(symshen_4internal_1symbols), tmp6938, V2800, V2801)


__e.TailApply(tmp6935, tmp6939)
return


}, 1)

tmp6940 := MakeNative(func(__e *ControlFlow) {
tmp6941 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symget), V2799, symshen_4internal_1symbols, tmp6941)
return


}, 0)

tmp6942 := MakeNative(func(__e *ControlFlow) {
Z2803 := __e.Get(1)
_ = Z2803
__e.Return(Nil)
return
}, 1)

tmp6943 := Call(__e, try_1catch, tmp6940, tmp6942)


__e.TailApply(tmp6934, tmp6943)
return


}, 3)

tmp6944 := Call(__e, ns2_1set, symshen_4record_1internal, tmp6933)


_ = tmp6944

tmp6945 := MakeNative(func(__e *ControlFlow) {
V2811 := __e.Get(1)
_ = V2811
V2812 := __e.Get(2)
_ = V2812
V2813 := __e.Get(3)
_ = V2813
tmp6954 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2813)
}
__typedArg0 := V2813
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp6954 {
tmp6946 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2813)
}
__typedArg0 := V2813
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp6947 := Call(__e, PrimFunc(symshen_4internal_1symbols), V2811, V2812, tmp6946)


tmp6948 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2813)
}
__typedArg0 := V2813
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6949 := Call(__e, PrimFunc(symshen_4internal_1symbols), V2811, V2812, tmp6948)


__e.TailApply(PrimFunc(symunion), tmp6947, tmp6949)
return


} else {
tmp6952 := Call(__e, PrimFunc(symshen_4internal_2), V2813, V2811, V2812)


if True == tmp6952 {
tmp6950 := Call(__e, PrimFunc(symshen_4intern_1in_1package), V2811, V2813)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp6950, Nil)
}
__typedArg0 := tmp6950
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return(Nil)
return
}


}


}, 3)

tmp6955 := Call(__e, ns2_1set, symshen_4internal_1symbols, tmp6945)


_ = tmp6955

tmp6956 := MakeNative(func(__e *ControlFlow) {
V2814 := __e.Get(1)
_ = V2814
V2815 := __e.Get(2)
_ = V2815
tmp6957 := MakeNative(func(__e *ControlFlow) {
W2816 := __e.Get(1)
_ = W2816
tmp6958 := Call(__e, PrimFunc(symunion), V2815, W2816)


tmp6959 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symput), V2814, symshen_4external_1symbols, tmp6958, tmp6959)
return


}, 1)

tmp6960 := MakeNative(func(__e *ControlFlow) {
tmp6961 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symget), V2814, symshen_4external_1symbols, tmp6961)
return


}, 0)

tmp6962 := MakeNative(func(__e *ControlFlow) {
Z2817 := __e.Get(1)
_ = Z2817
__e.Return(Nil)
return
}, 1)

tmp6963 := Call(__e, try_1catch, tmp6960, tmp6962)


__e.TailApply(tmp6957, tmp6963)
return


}, 2)

tmp6964 := Call(__e, ns2_1set, symshen_4record_1external, tmp6956)


_ = tmp6964

tmp6965 := MakeNative(func(__e *ControlFlow) {
V2822 := __e.Get(1)
_ = V2822
V2823 := __e.Get(2)
_ = V2823
V2824 := __e.Get(3)
_ = V2824
tmp6970 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2824)
}
__typedArg0 := V2824
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp6970 {
tmp6966 := MakeNative(func(__e *ControlFlow) {
Z2825 := __e.Get(1)
_ = Z2825
__e.TailApply(PrimFunc(symshen_4package_1symbols), V2822, V2823, Z2825)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp6966, V2824)
return


} else {
tmp6968 := Call(__e, PrimFunc(symshen_4internal_2), V2824, V2822, V2823)


if True == tmp6968 {
__e.TailApply(PrimFunc(symshen_4intern_1in_1package), V2822, V2824)
return
} else {
__e.Return(V2824)
return
}


}


}, 3)

tmp6971 := Call(__e, ns2_1set, symshen_4package_1symbols, tmp6965)


_ = tmp6971

tmp6972 := MakeNative(func(__e *ControlFlow) {
V2826 := __e.Get(1)
_ = V2826
V2827 := __e.Get(2)
_ = V2827
tmp6973 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(V2827)
}
__typedArg0 := V2827
return Call(__e, PrimFunc(symstr), __typedArg0)
})()

tmp6974 := Call(__e, PrimFunc(sym_8s), MakeString("."), tmp6973)


tmp6975 := Call(__e, PrimFunc(sym_8s), V2826, tmp6974)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(tmp6975)
}
__typedArg0 := tmp6975
return Call(__e, PrimFunc(symintern), __typedArg0)
})())
return


}, 2)

tmp6976 := Call(__e, ns2_1set, symshen_4intern_1in_1package, tmp6972)


_ = tmp6976

tmp6977 := MakeNative(func(__e *ControlFlow) {
V2828 := __e.Get(1)
_ = V2828
V2829 := __e.Get(2)
_ = V2829
V2830 := __e.Get(3)
_ = V2830
tmp7007 := Call(__e, PrimFunc(symelement_2), V2828, V2830)


if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp7007)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp7007
return Call(__e, PrimFunc(symnot), __typedArg0)
})() {
tmp7004 := Call(__e, PrimFunc(symshen_4sng_2), V2828)


tmp7005 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp7004)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp7004
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres6979 Obj

if True == tmp7005 {
tmp7002 := Call(__e, PrimFunc(symshen_4dbl_2), V2828)


tmp7003 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp7002)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp7002
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres6981 Obj

if True == tmp7003 {
tmp7001 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsymbol_2) {
return PrimIsSymbol(V2828)
}
__typedArg0 := V2828
return Call(__e, PrimFunc(symsymbol_2), __typedArg0)
})()

var ifres6983 Obj

if True == tmp7001 {
tmp6999 := Call(__e, PrimFunc(symshen_4sysfunc_2), V2828)


tmp7000 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp6999)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp6999
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres6985 Obj

if True == tmp7000 {
tmp6997 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(V2828)
}
__typedArg0 := V2828
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

tmp6998 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp6997)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp6997
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres6987 Obj

if True == tmp6998 {
tmp6994 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(V2828)
}
__typedArg0 := V2828
return Call(__e, PrimFunc(symstr), __typedArg0)
})()

tmp6995 := Call(__e, PrimFunc(symshen_4internal_1to_1shen_2), tmp6994)


tmp6996 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp6995)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp6995
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres6989 Obj

if True == tmp6996 {
tmp6991 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(V2828)
}
__typedArg0 := V2828
return Call(__e, PrimFunc(symstr), __typedArg0)
})()

tmp6992 := Call(__e, PrimFunc(symshen_4internal_1to_1P_2), V2829, tmp6991)


tmp6993 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp6992)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp6992
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres6990 Obj

if True == tmp6993 {
ifres6990 = True


} else {
ifres6990 = False


}

ifres6989 = ifres6990


} else {
ifres6989 = False


}

var ifres6988 Obj

if True == ifres6989 {
ifres6988 = True


} else {
ifres6988 = False


}

ifres6987 = ifres6988


} else {
ifres6987 = False


}

var ifres6986 Obj

if True == ifres6987 {
ifres6986 = True


} else {
ifres6986 = False


}

ifres6985 = ifres6986


} else {
ifres6985 = False


}

var ifres6984 Obj

if True == ifres6985 {
ifres6984 = True


} else {
ifres6984 = False


}

ifres6983 = ifres6984


} else {
ifres6983 = False


}

var ifres6982 Obj

if True == ifres6983 {
ifres6982 = True


} else {
ifres6982 = False


}

ifres6981 = ifres6982


} else {
ifres6981 = False


}

var ifres6980 Obj

if True == ifres6981 {
ifres6980 = True


} else {
ifres6980 = False


}

ifres6979 = ifres6980


} else {
ifres6979 = False


}

if True == ifres6979 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


} else {
__e.Return(False)
return
}


}, 3)

tmp7009 := Call(__e, ns2_1set, symshen_4internal_2, tmp6977)


_ = tmp7009

tmp7010 := MakeNative(func(__e *ControlFlow) {
V2835 := __e.Get(1)
_ = V2835
tmp7064 := Call(__e, PrimFunc(symshen_4_7string_2), V2835)


var ifres7012 Obj

if True == tmp7064 {
tmp7062 := Call(__e, PrimFunc(symhdstr), V2835)


tmp7063 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("s"), tmp7062)
}
__typedArg0 := MakeString("s")
__typedArg1 := tmp7062
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7014 Obj

if True == tmp7063 {
tmp7061 := Call(__e, PrimFunc(symshen_4_7string_2), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V2835)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V2835
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


var ifres7016 Obj

if True == tmp7061 {
tmp7058 := Call(__e, PrimFunc(symhdstr), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V2835)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V2835
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


tmp7059 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("h"), tmp7058)
}
__typedArg0 := MakeString("h")
__typedArg1 := tmp7058
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7018 Obj

if True == tmp7059 {
tmp7056 := Call(__e, PrimFunc(symshen_4_7string_2), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V2835)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(__typedS0)))
}}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V2835)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V2835
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


var ifres7020 Obj

if True == tmp7056 {
tmp7052 := Call(__e, PrimFunc(symhdstr), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V2835)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(__typedS0)))
}}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V2835)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V2835
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


tmp7053 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("e"), tmp7052)
}
__typedArg0 := MakeString("e")
__typedArg1 := tmp7052
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7022 Obj

if True == tmp7053 {
tmp7049 := Call(__e, PrimFunc(symshen_4_7string_2), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V2835)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(TypedStringTailValue(__typedS0))))
}}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V2835)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(__typedS0)))
}}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V2835)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V2835
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


var ifres7024 Obj

if True == tmp7049 {
tmp7044 := Call(__e, PrimFunc(symhdstr), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V2835)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(TypedStringTailValue(__typedS0))))
}}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V2835)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(__typedS0)))
}}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V2835)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V2835
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


tmp7045 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("n"), tmp7044)
}
__typedArg0 := MakeString("n")
__typedArg1 := tmp7044
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7026 Obj

if True == tmp7045 {
tmp7040 := Call(__e, PrimFunc(symshen_4_7string_2), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V2835)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(TypedStringTailValue(TypedStringTailValue(__typedS0)))))
}}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V2835)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(TypedStringTailValue(__typedS0))))
}}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V2835)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(__typedS0)))
}}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V2835)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V2835
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


var ifres7028 Obj

if True == tmp7040 {
tmp7034 := Call(__e, PrimFunc(symhdstr), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V2835)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(TypedStringTailValue(TypedStringTailValue(__typedS0)))))
}}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V2835)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(TypedStringTailValue(__typedS0))))
}}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V2835)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(__typedS0)))
}}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V2835)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V2835
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


tmp7035 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("."), tmp7034)
}
__typedArg0 := MakeString(".")
__typedArg1 := tmp7034
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7029 Obj

if True == tmp7035 {
ifres7029 = True


} else {
ifres7029 = False


}

ifres7028 = ifres7029


} else {
ifres7028 = False


}

var ifres7027 Obj

if True == ifres7028 {
ifres7027 = True


} else {
ifres7027 = False


}

ifres7026 = ifres7027


} else {
ifres7026 = False


}

var ifres7025 Obj

if True == ifres7026 {
ifres7025 = True


} else {
ifres7025 = False


}

ifres7024 = ifres7025


} else {
ifres7024 = False


}

var ifres7023 Obj

if True == ifres7024 {
ifres7023 = True


} else {
ifres7023 = False


}

ifres7022 = ifres7023


} else {
ifres7022 = False


}

var ifres7021 Obj

if True == ifres7022 {
ifres7021 = True


} else {
ifres7021 = False


}

ifres7020 = ifres7021


} else {
ifres7020 = False


}

var ifres7019 Obj

if True == ifres7020 {
ifres7019 = True


} else {
ifres7019 = False


}

ifres7018 = ifres7019


} else {
ifres7018 = False


}

var ifres7017 Obj

if True == ifres7018 {
ifres7017 = True


} else {
ifres7017 = False


}

ifres7016 = ifres7017


} else {
ifres7016 = False


}

var ifres7015 Obj

if True == ifres7016 {
ifres7015 = True


} else {
ifres7015 = False


}

ifres7014 = ifres7015


} else {
ifres7014 = False


}

var ifres7013 Obj

if True == ifres7014 {
ifres7013 = True


} else {
ifres7013 = False


}

ifres7012 = ifres7013


} else {
ifres7012 = False


}

if True == ifres7012 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 1)

tmp7065 := Call(__e, ns2_1set, symshen_4internal_1to_1shen_2, tmp7010)


_ = tmp7065

tmp7066 := MakeNative(func(__e *ControlFlow) {
V2836 := __e.Get(1)
_ = V2836
tmp7067 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp7068 := Call(__e, PrimFunc(symget), symshen, symshen_4external_1symbols, tmp7067)


__e.TailApply(PrimFunc(symelement_2), V2836, tmp7068)
return


}, 1)

tmp7069 := Call(__e, ns2_1set, symshen_4sysfunc_2, tmp7066)


_ = tmp7069

tmp7070 := MakeNative(func(__e *ControlFlow) {
V2844 := __e.Get(1)
_ = V2844
V2845 := __e.Get(2)
_ = V2845
tmp7091 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString(""), V2844)
}
__typedArg0 := MakeString("")
__typedArg1 := V2844
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7084 Obj

if True == tmp7091 {
tmp7090 := Call(__e, PrimFunc(symshen_4_7string_2), V2845)


var ifres7086 Obj

if True == tmp7090 {
tmp7088 := Call(__e, PrimFunc(symhdstr), V2845)


tmp7089 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("."), tmp7088)
}
__typedArg0 := MakeString(".")
__typedArg1 := tmp7088
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7087 Obj

if True == tmp7089 {
ifres7087 = True


} else {
ifres7087 = False


}

ifres7086 = ifres7087


} else {
ifres7086 = False


}

var ifres7085 Obj

if True == ifres7086 {
ifres7085 = True


} else {
ifres7085 = False


}

ifres7084 = ifres7085


} else {
ifres7084 = False


}

if True == ifres7084 {
__e.Return(True)
return
} else {
tmp7082 := Call(__e, PrimFunc(symshen_4_7string_2), V2844)


var ifres7074 Obj

if True == tmp7082 {
tmp7081 := Call(__e, PrimFunc(symshen_4_7string_2), V2845)


var ifres7076 Obj

if True == tmp7081 {
tmp7078 := Call(__e, PrimFunc(symhdstr), V2844)


tmp7079 := Call(__e, PrimFunc(symhdstr), V2845)


tmp7080 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp7078, tmp7079)
}
__typedArg0 := tmp7078
__typedArg1 := tmp7079
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7077 Obj

if True == tmp7080 {
ifres7077 = True


} else {
ifres7077 = False


}

ifres7076 = ifres7077


} else {
ifres7076 = False


}

var ifres7075 Obj

if True == ifres7076 {
ifres7075 = True


} else {
ifres7075 = False


}

ifres7074 = ifres7075


} else {
ifres7074 = False


}

if True == ifres7074 {
tmp7071 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V2844)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V2844
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4internal_1to_1P_2), tmp7071, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V2845)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V2845
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())
return


} else {
__e.Return(False)
return
}


}


}, 2)

tmp7092 := Call(__e, ns2_1set, symshen_4internal_1to_1P_2, tmp7070)


_ = tmp7092

tmp7093 := MakeNative(func(__e *ControlFlow) {
V2848 := __e.Get(1)
_ = V2848
V2849 := __e.Get(2)
_ = V2849
tmp7106 := Call(__e, PrimFunc(symelement_2), V2848, V2849)


if True == tmp7106 {
__e.Return(V2848)
return
} else {
tmp7104 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2848)
}
__typedArg0 := V2848
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7100 Obj

if True == tmp7104 {
tmp7102 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2848)
}
__typedArg0 := V2848
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7103 := Call(__e, PrimFunc(symshen_4non_1application_2), tmp7102)


var ifres7101 Obj

if True == tmp7103 {
ifres7101 = True


} else {
ifres7101 = False


}

ifres7100 = ifres7101


} else {
ifres7100 = False


}

if True == ifres7100 {
tmp7094 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2848)
}
__typedArg0 := V2848
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4special_1case), tmp7094, V2848, V2849)
return


} else {
tmp7098 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2848)
}
__typedArg0 := V2848
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp7098 {
tmp7095 := MakeNative(func(__e *ControlFlow) {
Z2850 := __e.Get(1)
_ = Z2850
__e.TailApply(PrimFunc(symshen_4process_1applications), Z2850, V2849)
return
}, 1)

tmp7096 := Call(__e, PrimFunc(symmap), tmp7095, V2848)


__e.TailApply(PrimFunc(symshen_4process_1application), tmp7096, V2849)
return


} else {
__e.Return(V2848)
return
}


}


}


}, 2)

tmp7107 := Call(__e, ns2_1set, symshen_4process_1applications, tmp7093)


_ = tmp7107

tmp7108 := MakeNative(func(__e *ControlFlow) {
V2853 := __e.Get(1)
_ = V2853
tmp7118 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefine, V2853)
}
__typedArg0 := symdefine
__typedArg1 := V2853
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp7118 {
__e.Return(True)
return
} else {
tmp7116 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefun, V2853)
}
__typedArg0 := symdefun
__typedArg1 := V2853
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp7116 {
__e.Return(True)
return
} else {
tmp7114 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symsynonyms, V2853)
}
__typedArg0 := symsynonyms
__typedArg1 := V2853
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp7114 {
__e.Return(True)
return
} else {
tmp7112 := Call(__e, PrimFunc(symshen_4special_2), V2853)


if True == tmp7112 {
__e.Return(True)
return
} else {
tmp7110 := Call(__e, PrimFunc(symshen_4extraspecial_2), V2853)


if True == tmp7110 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


}


}


}


}, 1)

tmp7119 := Call(__e, ns2_1set, symshen_4non_1application_2, tmp7108)


_ = tmp7119

tmp7120 := MakeNative(func(__e *ControlFlow) {
V2858 := __e.Get(1)
_ = V2858
V2859 := __e.Get(2)
_ = V2859
V2860 := __e.Get(3)
_ = V2860
tmp7362 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlambda, V2858)
}
__typedArg0 := symlambda
__typedArg1 := V2858
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7340 Obj

if True == tmp7362 {
tmp7361 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7342 Obj

if True == tmp7361 {
tmp7359 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7360 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlambda, tmp7359)
}
__typedArg0 := symlambda
__typedArg1 := tmp7359
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7344 Obj

if True == tmp7360 {
tmp7357 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7358 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7357)
}
__typedArg0 := tmp7357
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7346 Obj

if True == tmp7358 {
tmp7354 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7355 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7354)
}
__typedArg0 := tmp7354
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7356 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7355)
}
__typedArg0 := tmp7355
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7348 Obj

if True == tmp7356 {
tmp7350 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7351 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7350)
}
__typedArg0 := tmp7350
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7352 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7351)
}
__typedArg0 := tmp7351
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7353 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp7352)
}
__typedArg0 := Nil
__typedArg1 := tmp7352
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7349 Obj

if True == tmp7353 {
ifres7349 = True


} else {
ifres7349 = False


}

ifres7348 = ifres7349


} else {
ifres7348 = False


}

var ifres7347 Obj

if True == ifres7348 {
ifres7347 = True


} else {
ifres7347 = False


}

ifres7346 = ifres7347


} else {
ifres7346 = False


}

var ifres7345 Obj

if True == ifres7346 {
ifres7345 = True


} else {
ifres7345 = False


}

ifres7344 = ifres7345


} else {
ifres7344 = False


}

var ifres7343 Obj

if True == ifres7344 {
ifres7343 = True


} else {
ifres7343 = False


}

ifres7342 = ifres7343


} else {
ifres7342 = False


}

var ifres7341 Obj

if True == ifres7342 {
ifres7341 = True


} else {
ifres7341 = False


}

ifres7340 = ifres7341


} else {
ifres7340 = False


}

if True == ifres7340 {
tmp7121 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7122 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7121)
}
__typedArg0 := tmp7121
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7123 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7124 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7123)
}
__typedArg0 := tmp7123
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7125 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7124)
}
__typedArg0 := tmp7124
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7126 := Call(__e, PrimFunc(symshen_4process_1applications), tmp7125, V2860)


tmp7127 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7126, Nil)
}
__typedArg0 := tmp7126
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7128 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7122, tmp7127)
}
__typedArg0 := tmp7122
__typedArg1 := tmp7127
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp7128)
}
__typedArg0 := symlambda
__typedArg1 := tmp7128
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp7338 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlet, V2858)
}
__typedArg0 := symlet
__typedArg1 := V2858
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7309 Obj

if True == tmp7338 {
tmp7337 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7311 Obj

if True == tmp7337 {
tmp7335 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7336 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlet, tmp7335)
}
__typedArg0 := symlet
__typedArg1 := tmp7335
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7313 Obj

if True == tmp7336 {
tmp7333 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7334 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7333)
}
__typedArg0 := tmp7333
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7315 Obj

if True == tmp7334 {
tmp7330 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7331 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7330)
}
__typedArg0 := tmp7330
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7332 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7331)
}
__typedArg0 := tmp7331
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7317 Obj

if True == tmp7332 {
tmp7326 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7327 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7326)
}
__typedArg0 := tmp7326
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7328 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7327)
}
__typedArg0 := tmp7327
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7329 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7328)
}
__typedArg0 := tmp7328
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7319 Obj

if True == tmp7329 {
tmp7321 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7322 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7321)
}
__typedArg0 := tmp7321
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7323 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7322)
}
__typedArg0 := tmp7322
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7324 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7323)
}
__typedArg0 := tmp7323
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7325 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp7324)
}
__typedArg0 := Nil
__typedArg1 := tmp7324
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7320 Obj

if True == tmp7325 {
ifres7320 = True


} else {
ifres7320 = False


}

ifres7319 = ifres7320


} else {
ifres7319 = False


}

var ifres7318 Obj

if True == ifres7319 {
ifres7318 = True


} else {
ifres7318 = False


}

ifres7317 = ifres7318


} else {
ifres7317 = False


}

var ifres7316 Obj

if True == ifres7317 {
ifres7316 = True


} else {
ifres7316 = False


}

ifres7315 = ifres7316


} else {
ifres7315 = False


}

var ifres7314 Obj

if True == ifres7315 {
ifres7314 = True


} else {
ifres7314 = False


}

ifres7313 = ifres7314


} else {
ifres7313 = False


}

var ifres7312 Obj

if True == ifres7313 {
ifres7312 = True


} else {
ifres7312 = False


}

ifres7311 = ifres7312


} else {
ifres7311 = False


}

var ifres7310 Obj

if True == ifres7311 {
ifres7310 = True


} else {
ifres7310 = False


}

ifres7309 = ifres7310


} else {
ifres7309 = False


}

if True == ifres7309 {
tmp7129 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7130 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7129)
}
__typedArg0 := tmp7129
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7131 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7132 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7131)
}
__typedArg0 := tmp7131
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7133 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7132)
}
__typedArg0 := tmp7132
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7134 := Call(__e, PrimFunc(symshen_4process_1applications), tmp7133, V2860)


tmp7135 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7136 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7135)
}
__typedArg0 := tmp7135
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7137 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7136)
}
__typedArg0 := tmp7136
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7138 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7137)
}
__typedArg0 := tmp7137
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7139 := Call(__e, PrimFunc(symshen_4process_1applications), tmp7138, V2860)


tmp7140 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7139, Nil)
}
__typedArg0 := tmp7139
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7141 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7134, tmp7140)
}
__typedArg0 := tmp7134
__typedArg1 := tmp7140
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7142 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7130, tmp7141)
}
__typedArg0 := tmp7130
__typedArg1 := tmp7141
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp7142)
}
__typedArg0 := symlet
__typedArg1 := tmp7142
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp7307 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefun, V2858)
}
__typedArg0 := symdefun
__typedArg1 := V2858
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7278 Obj

if True == tmp7307 {
tmp7306 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7280 Obj

if True == tmp7306 {
tmp7304 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7305 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefun, tmp7304)
}
__typedArg0 := symdefun
__typedArg1 := tmp7304
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7282 Obj

if True == tmp7305 {
tmp7302 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7303 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7302)
}
__typedArg0 := tmp7302
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7284 Obj

if True == tmp7303 {
tmp7299 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7300 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7299)
}
__typedArg0 := tmp7299
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7301 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7300)
}
__typedArg0 := tmp7300
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7286 Obj

if True == tmp7301 {
tmp7295 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7296 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7295)
}
__typedArg0 := tmp7295
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7297 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7296)
}
__typedArg0 := tmp7296
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7298 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7297)
}
__typedArg0 := tmp7297
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7288 Obj

if True == tmp7298 {
tmp7290 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7291 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7290)
}
__typedArg0 := tmp7290
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7292 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7291)
}
__typedArg0 := tmp7291
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7293 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7292)
}
__typedArg0 := tmp7292
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7294 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp7293)
}
__typedArg0 := Nil
__typedArg1 := tmp7293
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7289 Obj

if True == tmp7294 {
ifres7289 = True


} else {
ifres7289 = False


}

ifres7288 = ifres7289


} else {
ifres7288 = False


}

var ifres7287 Obj

if True == ifres7288 {
ifres7287 = True


} else {
ifres7287 = False


}

ifres7286 = ifres7287


} else {
ifres7286 = False


}

var ifres7285 Obj

if True == ifres7286 {
ifres7285 = True


} else {
ifres7285 = False


}

ifres7284 = ifres7285


} else {
ifres7284 = False


}

var ifres7283 Obj

if True == ifres7284 {
ifres7283 = True


} else {
ifres7283 = False


}

ifres7282 = ifres7283


} else {
ifres7282 = False


}

var ifres7281 Obj

if True == ifres7282 {
ifres7281 = True


} else {
ifres7281 = False


}

ifres7280 = ifres7281


} else {
ifres7280 = False


}

var ifres7279 Obj

if True == ifres7280 {
ifres7279 = True


} else {
ifres7279 = False


}

ifres7278 = ifres7279


} else {
ifres7278 = False


}

if True == ifres7278 {
__e.Return(V2859)
return
} else {
tmp7276 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefine, V2858)
}
__typedArg0 := symdefine
__typedArg1 := V2858
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7254 Obj

if True == tmp7276 {
tmp7275 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7256 Obj

if True == tmp7275 {
tmp7273 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7274 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefine, tmp7273)
}
__typedArg0 := symdefine
__typedArg1 := tmp7273
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7258 Obj

if True == tmp7274 {
tmp7271 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7272 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7271)
}
__typedArg0 := tmp7271
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7260 Obj

if True == tmp7272 {
tmp7268 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7269 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7268)
}
__typedArg0 := tmp7268
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7270 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7269)
}
__typedArg0 := tmp7269
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7262 Obj

if True == tmp7270 {
tmp7264 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7265 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7264)
}
__typedArg0 := tmp7264
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7266 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7265)
}
__typedArg0 := tmp7265
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7267 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_i, tmp7266)
}
__typedArg0 := sym_i
__typedArg1 := tmp7266
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7263 Obj

if True == tmp7267 {
ifres7263 = True


} else {
ifres7263 = False


}

ifres7262 = ifres7263


} else {
ifres7262 = False


}

var ifres7261 Obj

if True == ifres7262 {
ifres7261 = True


} else {
ifres7261 = False


}

ifres7260 = ifres7261


} else {
ifres7260 = False


}

var ifres7259 Obj

if True == ifres7260 {
ifres7259 = True


} else {
ifres7259 = False


}

ifres7258 = ifres7259


} else {
ifres7258 = False


}

var ifres7257 Obj

if True == ifres7258 {
ifres7257 = True


} else {
ifres7257 = False


}

ifres7256 = ifres7257


} else {
ifres7256 = False


}

var ifres7255 Obj

if True == ifres7256 {
ifres7255 = True


} else {
ifres7255 = False


}

ifres7254 = ifres7255


} else {
ifres7254 = False


}

if True == ifres7254 {
tmp7143 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7144 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7143)
}
__typedArg0 := tmp7143
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7145 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7146 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7145)
}
__typedArg0 := tmp7145
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7147 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7148 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7147)
}
__typedArg0 := tmp7147
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7149 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7148)
}
__typedArg0 := tmp7148
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7150 := Call(__e, PrimFunc(symshen_4process_1after_1type), tmp7146, tmp7149, V2860)


tmp7151 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_i, tmp7150)
}
__typedArg0 := sym_i
__typedArg1 := tmp7150
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7152 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7144, tmp7151)
}
__typedArg0 := tmp7144
__typedArg1 := tmp7151
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefine, tmp7152)
}
__typedArg0 := symdefine
__typedArg1 := tmp7152
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp7252 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefine, V2858)
}
__typedArg0 := symdefine
__typedArg1 := V2858
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7241 Obj

if True == tmp7252 {
tmp7251 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7243 Obj

if True == tmp7251 {
tmp7249 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7250 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefine, tmp7249)
}
__typedArg0 := symdefine
__typedArg1 := tmp7249
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7245 Obj

if True == tmp7250 {
tmp7247 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7248 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7247)
}
__typedArg0 := tmp7247
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7246 Obj

if True == tmp7248 {
ifres7246 = True


} else {
ifres7246 = False


}

ifres7245 = ifres7246


} else {
ifres7245 = False


}

var ifres7244 Obj

if True == ifres7245 {
ifres7244 = True


} else {
ifres7244 = False


}

ifres7243 = ifres7244


} else {
ifres7243 = False


}

var ifres7242 Obj

if True == ifres7243 {
ifres7242 = True


} else {
ifres7242 = False


}

ifres7241 = ifres7242


} else {
ifres7241 = False


}

if True == ifres7241 {
tmp7153 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7154 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7153)
}
__typedArg0 := tmp7153
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7155 := MakeNative(func(__e *ControlFlow) {
Z2861 := __e.Get(1)
_ = Z2861
__e.TailApply(PrimFunc(symshen_4process_1applications), Z2861, V2860)
return
}, 1)

tmp7156 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7157 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7156)
}
__typedArg0 := tmp7156
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7158 := Call(__e, PrimFunc(symmap), tmp7155, tmp7157)


tmp7159 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7154, tmp7158)
}
__typedArg0 := tmp7154
__typedArg1 := tmp7158
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefine, tmp7159)
}
__typedArg0 := symdefine
__typedArg1 := tmp7159
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp7239 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symsynonyms, V2858)
}
__typedArg0 := symsynonyms
__typedArg1 := V2858
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp7239 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsynonyms, V2859)
}
__typedArg0 := symsynonyms
__typedArg1 := V2859
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return
} else {
tmp7237 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symtype, V2858)
}
__typedArg0 := symtype
__typedArg1 := V2858
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7215 Obj

if True == tmp7237 {
tmp7236 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7217 Obj

if True == tmp7236 {
tmp7234 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7235 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symtype, tmp7234)
}
__typedArg0 := symtype
__typedArg1 := tmp7234
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7219 Obj

if True == tmp7235 {
tmp7232 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7233 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7232)
}
__typedArg0 := tmp7232
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7221 Obj

if True == tmp7233 {
tmp7229 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7230 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7229)
}
__typedArg0 := tmp7229
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7231 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7230)
}
__typedArg0 := tmp7230
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7223 Obj

if True == tmp7231 {
tmp7225 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7226 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7225)
}
__typedArg0 := tmp7225
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7227 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7226)
}
__typedArg0 := tmp7226
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7228 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp7227)
}
__typedArg0 := Nil
__typedArg1 := tmp7227
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7224 Obj

if True == tmp7228 {
ifres7224 = True


} else {
ifres7224 = False


}

ifres7223 = ifres7224


} else {
ifres7223 = False


}

var ifres7222 Obj

if True == ifres7223 {
ifres7222 = True


} else {
ifres7222 = False


}

ifres7221 = ifres7222


} else {
ifres7221 = False


}

var ifres7220 Obj

if True == ifres7221 {
ifres7220 = True


} else {
ifres7220 = False


}

ifres7219 = ifres7220


} else {
ifres7219 = False


}

var ifres7218 Obj

if True == ifres7219 {
ifres7218 = True


} else {
ifres7218 = False


}

ifres7217 = ifres7218


} else {
ifres7217 = False


}

var ifres7216 Obj

if True == ifres7217 {
ifres7216 = True


} else {
ifres7216 = False


}

ifres7215 = ifres7216


} else {
ifres7215 = False


}

if True == ifres7215 {
tmp7160 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7161 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7160)
}
__typedArg0 := tmp7160
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7162 := Call(__e, PrimFunc(symshen_4process_1applications), tmp7161, V2860)


tmp7163 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7164 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7163)
}
__typedArg0 := tmp7163
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7165 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7162, tmp7164)
}
__typedArg0 := tmp7162
__typedArg1 := tmp7164
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtype, tmp7165)
}
__typedArg0 := symtype
__typedArg1 := tmp7165
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp7213 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(syminput_7, V2858)
}
__typedArg0 := syminput_7
__typedArg1 := V2858
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7191 Obj

if True == tmp7213 {
tmp7212 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7193 Obj

if True == tmp7212 {
tmp7210 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7211 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(syminput_7, tmp7210)
}
__typedArg0 := syminput_7
__typedArg1 := tmp7210
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7195 Obj

if True == tmp7211 {
tmp7208 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7209 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7208)
}
__typedArg0 := tmp7208
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7197 Obj

if True == tmp7209 {
tmp7205 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7206 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7205)
}
__typedArg0 := tmp7205
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7207 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7206)
}
__typedArg0 := tmp7206
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7199 Obj

if True == tmp7207 {
tmp7201 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7202 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7201)
}
__typedArg0 := tmp7201
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7203 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7202)
}
__typedArg0 := tmp7202
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7204 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp7203)
}
__typedArg0 := Nil
__typedArg1 := tmp7203
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7200 Obj

if True == tmp7204 {
ifres7200 = True


} else {
ifres7200 = False


}

ifres7199 = ifres7200


} else {
ifres7199 = False


}

var ifres7198 Obj

if True == ifres7199 {
ifres7198 = True


} else {
ifres7198 = False


}

ifres7197 = ifres7198


} else {
ifres7197 = False


}

var ifres7196 Obj

if True == ifres7197 {
ifres7196 = True


} else {
ifres7196 = False


}

ifres7195 = ifres7196


} else {
ifres7195 = False


}

var ifres7194 Obj

if True == ifres7195 {
ifres7194 = True


} else {
ifres7194 = False


}

ifres7193 = ifres7194


} else {
ifres7193 = False


}

var ifres7192 Obj

if True == ifres7193 {
ifres7192 = True


} else {
ifres7192 = False


}

ifres7191 = ifres7192


} else {
ifres7191 = False


}

if True == ifres7191 {
tmp7166 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7167 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7166)
}
__typedArg0 := tmp7166
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7168 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7169 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7168)
}
__typedArg0 := tmp7168
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7170 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7169)
}
__typedArg0 := tmp7169
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7171 := Call(__e, PrimFunc(symshen_4process_1applications), tmp7170, V2860)


tmp7172 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7171, Nil)
}
__typedArg0 := tmp7171
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7173 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7167, tmp7172)
}
__typedArg0 := tmp7167
__typedArg1 := tmp7172
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminput_7, tmp7173)
}
__typedArg0 := syminput_7
__typedArg1 := tmp7173
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp7189 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7185 Obj

if True == tmp7189 {
tmp7187 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7188 := Call(__e, PrimFunc(symshen_4special_2), tmp7187)


var ifres7186 Obj

if True == tmp7188 {
ifres7186 = True


} else {
ifres7186 = False


}

ifres7185 = ifres7186


} else {
ifres7185 = False


}

if True == ifres7185 {
tmp7174 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7175 := MakeNative(func(__e *ControlFlow) {
Z2862 := __e.Get(1)
_ = Z2862
__e.TailApply(PrimFunc(symshen_4process_1applications), Z2862, V2860)
return
}, 1)

tmp7176 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7177 := Call(__e, PrimFunc(symmap), tmp7175, tmp7176)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7174, tmp7177)
}
__typedArg0 := tmp7174
__typedArg1 := tmp7177
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp7183 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7179 Obj

if True == tmp7183 {
tmp7181 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7182 := Call(__e, PrimFunc(symshen_4extraspecial_2), tmp7181)


var ifres7180 Obj

if True == tmp7182 {
ifres7180 = True


} else {
ifres7180 = False


}

ifres7179 = ifres7180


} else {
ifres7179 = False


}

if True == ifres7179 {
__e.Return(V2859)
return
} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("partial function shen.special-case"))
}
__typedArg0 := MakeString("partial function shen.special-case")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
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


}, 3)

tmp7363 := Call(__e, ns2_1set, symshen_4special_1case, tmp7120)


_ = tmp7363

tmp7364 := MakeNative(func(__e *ControlFlow) {
V2865 := __e.Get(1)
_ = V2865
V2866 := __e.Get(2)
_ = V2866
V2867 := __e.Get(3)
_ = V2867
tmp7380 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2866)
}
__typedArg0 := V2866
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7376 Obj

if True == tmp7380 {
tmp7378 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2866)
}
__typedArg0 := V2866
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7379 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_j, tmp7378)
}
__typedArg0 := sym_j
__typedArg1 := tmp7378
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7377 Obj

if True == tmp7379 {
ifres7377 = True


} else {
ifres7377 = False


}

ifres7376 = ifres7377


} else {
ifres7376 = False


}

if True == ifres7376 {
tmp7365 := MakeNative(func(__e *ControlFlow) {
Z2868 := __e.Get(1)
_ = Z2868
__e.TailApply(PrimFunc(symshen_4process_1applications), Z2868, V2867)
return
}, 1)

tmp7366 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2866)
}
__typedArg0 := V2866
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7367 := Call(__e, PrimFunc(symmap), tmp7365, tmp7366)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_j, tmp7367)
}
__typedArg0 := sym_j
__typedArg1 := tmp7367
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp7374 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2866)
}
__typedArg0 := V2866
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp7374 {
tmp7368 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2866)
}
__typedArg0 := V2866
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7369 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2866)
}
__typedArg0 := V2866
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7370 := Call(__e, PrimFunc(symshen_4process_1after_1type), V2865, tmp7369, V2867)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7368, tmp7370)
}
__typedArg0 := tmp7368
__typedArg1 := tmp7370
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp7371 := Call(__e, PrimFunc(symshen_4app), V2865, MakeString("\n"), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("missing } in "))
__typedS1, __typedOK1 := TypedString(tmp7371)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("missing } in ")
__typedArg1 := tmp7371
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("missing } in "))
__typedS1, __typedOK1 := TypedString(tmp7371)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("missing } in ")
__typedArg1 := tmp7371
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}


}


}, 3)

tmp7381 := Call(__e, ns2_1set, symshen_4process_1after_1type, tmp7364)


_ = tmp7381

tmp7382 := MakeNative(func(__e *ControlFlow) {
V2869 := __e.Get(1)
_ = V2869
V2870 := __e.Get(2)
_ = V2870
tmp7427 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2869)
}
__typedArg0 := V2869
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp7427 {
tmp7383 := MakeNative(func(__e *ControlFlow) {
W2871 := __e.Get(1)
_ = W2871
tmp7384 := MakeNative(func(__e *ControlFlow) {
W2872 := __e.Get(1)
_ = W2872
tmp7421 := Call(__e, PrimFunc(symelement_2), V2869, V2870)


if True == tmp7421 {
__e.Return(V2869)
return
} else {
tmp7418 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2869)
}
__typedArg0 := V2869
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7419 := Call(__e, PrimFunc(symshen_4shen_1call_2), tmp7418)


if True == tmp7419 {
__e.Return(V2869)
return
} else {
tmp7416 := Call(__e, PrimFunc(symshen_4foreign_2), V2869)


if True == tmp7416 {
__e.TailApply(PrimFunc(symshen_4unpack_1foreign), V2869)
return
} else {
tmp7414 := Call(__e, PrimFunc(symshen_4fn_1call_2), V2869)


if True == tmp7414 {
__e.TailApply(PrimFunc(symshen_4fn_1call), V2869)
return
} else {
tmp7412 := Call(__e, PrimFunc(symshen_4zero_1place_2), V2869)


if True == tmp7412 {
__e.Return(V2869)
return
} else {
tmp7409 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2869)
}
__typedArg0 := V2869
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7410 := Call(__e, PrimFunc(symshen_4undefined_1f_2), tmp7409, W2871)


if True == tmp7410 {
tmp7385 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2869)
}
__typedArg0 := V2869
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7386 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7385, Nil)
}
__typedArg0 := tmp7385
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7387 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfn, tmp7386)
}
__typedArg0 := symfn
__typedArg1 := tmp7386
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7388 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2869)
}
__typedArg0 := V2869
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7389 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7387, tmp7388)
}
__typedArg0 := tmp7387
__typedArg1 := tmp7388
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4simple_1curry), tmp7389)
return


} else {
tmp7406 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2869)
}
__typedArg0 := V2869
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7407 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(tmp7406)
}
__typedArg0 := tmp7406
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

if True == tmp7407 {
__e.TailApply(PrimFunc(symshen_4simple_1curry), V2869)
return
} else {
tmp7403 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2869)
}
__typedArg0 := V2869
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7404 := Call(__e, PrimFunc(symshen_4application_2), tmp7403)


if True == tmp7404 {
__e.TailApply(PrimFunc(symshen_4simple_1curry), V2869)
return
} else {
tmp7400 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2869)
}
__typedArg0 := V2869
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7401 := Call(__e, PrimFunc(symshen_4partial_1application_d_2), tmp7400, W2871, W2872)


if True == tmp7401 {
__e.TailApply(PrimFunc(symshen_4lambda_1function), V2869, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(W2871)
__typedN1, __typedOK1 := TypedFloat64(W2872)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := W2871
__typedArg1 := W2872
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})())
return


} else {
tmp7397 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2869)
}
__typedArg0 := V2869
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7398 := Call(__e, PrimFunc(symshen_4overapplication_2), tmp7397, W2871, W2872)


if True == tmp7398 {
tmp7391 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2869)
}
__typedArg0 := V2869
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7392 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7391, Nil)
}
__typedArg0 := tmp7391
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7393 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfn, tmp7392)
}
__typedArg0 := symfn
__typedArg1 := tmp7392
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7394 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2869)
}
__typedArg0 := V2869
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7395 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7393, tmp7394)
}
__typedArg0 := tmp7393
__typedArg1 := tmp7394
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4simple_1curry), tmp7395)
return


} else {
__e.Return(V2869)
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


}, 1)

tmp7422 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2869)
}
__typedArg0 := V2869
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7423 := Call(__e, PrimFunc(symlength), tmp7422)


__e.TailApply(tmp7384, tmp7423)
return


}, 1)

tmp7424 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2869)
}
__typedArg0 := V2869
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7425 := Call(__e, PrimFunc(symarity), tmp7424)


__e.TailApply(tmp7383, tmp7425)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("partial function shen.process-application"))
}
__typedArg0 := MakeString("partial function shen.process-application")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}, 2)

tmp7428 := Call(__e, ns2_1set, symshen_4process_1application, tmp7382)


_ = tmp7428

tmp7429 := MakeNative(func(__e *ControlFlow) {
V2873 := __e.Get(1)
_ = V2873
tmp7455 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2873)
}
__typedArg0 := V2873
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7435 Obj

if True == tmp7455 {
tmp7453 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2873)
}
__typedArg0 := V2873
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7454 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7453)
}
__typedArg0 := tmp7453
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7437 Obj

if True == tmp7454 {
tmp7450 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2873)
}
__typedArg0 := V2873
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7451 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7450)
}
__typedArg0 := tmp7450
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7452 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symforeign, tmp7451)
}
__typedArg0 := symforeign
__typedArg1 := tmp7451
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7439 Obj

if True == tmp7452 {
tmp7447 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2873)
}
__typedArg0 := V2873
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7448 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7447)
}
__typedArg0 := tmp7447
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7449 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7448)
}
__typedArg0 := tmp7448
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7441 Obj

if True == tmp7449 {
tmp7443 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2873)
}
__typedArg0 := V2873
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7444 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7443)
}
__typedArg0 := tmp7443
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7445 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7444)
}
__typedArg0 := tmp7444
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7446 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp7445)
}
__typedArg0 := Nil
__typedArg1 := tmp7445
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7442 Obj

if True == tmp7446 {
ifres7442 = True


} else {
ifres7442 = False


}

ifres7441 = ifres7442


} else {
ifres7441 = False


}

var ifres7440 Obj

if True == ifres7441 {
ifres7440 = True


} else {
ifres7440 = False


}

ifres7439 = ifres7440


} else {
ifres7439 = False


}

var ifres7438 Obj

if True == ifres7439 {
ifres7438 = True


} else {
ifres7438 = False


}

ifres7437 = ifres7438


} else {
ifres7437 = False


}

var ifres7436 Obj

if True == ifres7437 {
ifres7436 = True


} else {
ifres7436 = False


}

ifres7435 = ifres7436


} else {
ifres7435 = False


}

if True == ifres7435 {
tmp7430 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2873)
}
__typedArg0 := V2873
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7431 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7430)
}
__typedArg0 := tmp7430
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7432 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7431)
}
__typedArg0 := tmp7431
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7433 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2873)
}
__typedArg0 := V2873
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7432, tmp7433)
}
__typedArg0 := tmp7432
__typedArg1 := tmp7433
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("partial function shen.unpack-foreign"))
}
__typedArg0 := MakeString("partial function shen.unpack-foreign")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}, 1)

tmp7456 := Call(__e, ns2_1set, symshen_4unpack_1foreign, tmp7429)


_ = tmp7456

tmp7457 := MakeNative(func(__e *ControlFlow) {
V2876 := __e.Get(1)
_ = V2876
tmp7479 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2876)
}
__typedArg0 := V2876
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7459 Obj

if True == tmp7479 {
tmp7477 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2876)
}
__typedArg0 := V2876
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7478 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7477)
}
__typedArg0 := tmp7477
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7461 Obj

if True == tmp7478 {
tmp7474 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2876)
}
__typedArg0 := V2876
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7475 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7474)
}
__typedArg0 := tmp7474
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7476 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symforeign, tmp7475)
}
__typedArg0 := symforeign
__typedArg1 := tmp7475
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7463 Obj

if True == tmp7476 {
tmp7471 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2876)
}
__typedArg0 := V2876
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7472 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7471)
}
__typedArg0 := tmp7471
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7473 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7472)
}
__typedArg0 := tmp7472
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7465 Obj

if True == tmp7473 {
tmp7467 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2876)
}
__typedArg0 := V2876
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7468 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7467)
}
__typedArg0 := tmp7467
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7469 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7468)
}
__typedArg0 := tmp7468
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7470 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp7469)
}
__typedArg0 := Nil
__typedArg1 := tmp7469
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7466 Obj

if True == tmp7470 {
ifres7466 = True


} else {
ifres7466 = False


}

ifres7465 = ifres7466


} else {
ifres7465 = False


}

var ifres7464 Obj

if True == ifres7465 {
ifres7464 = True


} else {
ifres7464 = False


}

ifres7463 = ifres7464


} else {
ifres7463 = False


}

var ifres7462 Obj

if True == ifres7463 {
ifres7462 = True


} else {
ifres7462 = False


}

ifres7461 = ifres7462


} else {
ifres7461 = False


}

var ifres7460 Obj

if True == ifres7461 {
ifres7460 = True


} else {
ifres7460 = False


}

ifres7459 = ifres7460


} else {
ifres7459 = False


}

if True == ifres7459 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 1)

tmp7480 := Call(__e, ns2_1set, symshen_4foreign_2, tmp7457)


_ = tmp7480

tmp7481 := MakeNative(func(__e *ControlFlow) {
V2879 := __e.Get(1)
_ = V2879
tmp7487 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2879)
}
__typedArg0 := V2879
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7483 Obj

if True == tmp7487 {
tmp7485 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2879)
}
__typedArg0 := V2879
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7486 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp7485)
}
__typedArg0 := Nil
__typedArg1 := tmp7485
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7484 Obj

if True == tmp7486 {
ifres7484 = True


} else {
ifres7484 = False


}

ifres7483 = ifres7484


} else {
ifres7483 = False


}

if True == ifres7483 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 1)

tmp7488 := Call(__e, ns2_1set, symshen_4zero_1place_2, tmp7481)


_ = tmp7488

tmp7489 := MakeNative(func(__e *ControlFlow) {
V2880 := __e.Get(1)
_ = V2880
tmp7494 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsymbol_2) {
return PrimIsSymbol(V2880)
}
__typedArg0 := V2880
return Call(__e, PrimFunc(symsymbol_2), __typedArg0)
})()

if True == tmp7494 {
tmp7491 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(V2880)
}
__typedArg0 := V2880
return Call(__e, PrimFunc(symstr), __typedArg0)
})()

tmp7492 := Call(__e, PrimFunc(symshen_4internal_1to_1shen_2), tmp7491)


if True == tmp7492 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


} else {
__e.Return(False)
return
}


}, 1)

tmp7495 := Call(__e, ns2_1set, symshen_4shen_1call_2, tmp7489)


_ = tmp7495

tmp7496 := MakeNative(func(__e *ControlFlow) {
V2885 := __e.Get(1)
_ = V2885
tmp7526 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2885)
}
__typedArg0 := V2885
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7513 Obj

if True == tmp7526 {
tmp7524 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2885)
}
__typedArg0 := V2885
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7525 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symprotect, tmp7524)
}
__typedArg0 := symprotect
__typedArg1 := tmp7524
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7515 Obj

if True == tmp7525 {
tmp7522 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2885)
}
__typedArg0 := V2885
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7523 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7522)
}
__typedArg0 := tmp7522
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7517 Obj

if True == tmp7523 {
tmp7519 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2885)
}
__typedArg0 := V2885
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7520 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7519)
}
__typedArg0 := tmp7519
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7521 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp7520)
}
__typedArg0 := Nil
__typedArg1 := tmp7520
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7518 Obj

if True == tmp7521 {
ifres7518 = True


} else {
ifres7518 = False


}

ifres7517 = ifres7518


} else {
ifres7517 = False


}

var ifres7516 Obj

if True == ifres7517 {
ifres7516 = True


} else {
ifres7516 = False


}

ifres7515 = ifres7516


} else {
ifres7515 = False


}

var ifres7514 Obj

if True == ifres7515 {
ifres7514 = True


} else {
ifres7514 = False


}

ifres7513 = ifres7514


} else {
ifres7513 = False


}

if True == ifres7513 {
__e.Return(False)
return
} else {
tmp7511 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2885)
}
__typedArg0 := V2885
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7498 Obj

if True == tmp7511 {
tmp7509 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2885)
}
__typedArg0 := V2885
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7510 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symforeign, tmp7509)
}
__typedArg0 := symforeign
__typedArg1 := tmp7509
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7500 Obj

if True == tmp7510 {
tmp7507 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2885)
}
__typedArg0 := V2885
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7508 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7507)
}
__typedArg0 := tmp7507
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7502 Obj

if True == tmp7508 {
tmp7504 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2885)
}
__typedArg0 := V2885
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7505 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7504)
}
__typedArg0 := tmp7504
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7506 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp7505)
}
__typedArg0 := Nil
__typedArg1 := tmp7505
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7503 Obj

if True == tmp7506 {
ifres7503 = True


} else {
ifres7503 = False


}

ifres7502 = ifres7503


} else {
ifres7502 = False


}

var ifres7501 Obj

if True == ifres7502 {
ifres7501 = True


} else {
ifres7501 = False


}

ifres7500 = ifres7501


} else {
ifres7500 = False


}

var ifres7499 Obj

if True == ifres7500 {
ifres7499 = True


} else {
ifres7499 = False


}

ifres7498 = ifres7499


} else {
ifres7498 = False


}

if True == ifres7498 {
__e.Return(False)
return
} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2885)
}
__typedArg0 := V2885
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})())
return
}


}


}, 1)

tmp7527 := Call(__e, ns2_1set, symshen_4application_2, tmp7496)


_ = tmp7527

tmp7528 := MakeNative(func(__e *ControlFlow) {
V2890 := __e.Get(1)
_ = V2890
V2891 := __e.Get(2)
_ = V2891
tmp7536 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeNumber(-1), V2891)
}
__typedArg0 := MakeNumber(-1)
__typedArg1 := V2891
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp7536 {
tmp7534 := Call(__e, PrimFunc(symshen_4lowercase_1symbol_2), V2890)


if True == tmp7534 {
tmp7530 := Call(__e, PrimFunc(symexternal), symshen)


tmp7531 := Call(__e, PrimFunc(symelement_2), V2890, tmp7530)


if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp7531)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp7531
return Call(__e, PrimFunc(symnot), __typedArg0)
})() {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


} else {
__e.Return(False)
return
}


} else {
__e.Return(False)
return
}


}, 2)

tmp7537 := Call(__e, ns2_1set, symshen_4undefined_1f_2, tmp7528)


_ = tmp7537

tmp7538 := MakeNative(func(__e *ControlFlow) {
V2892 := __e.Get(1)
_ = V2892
tmp7543 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsymbol_2) {
return PrimIsSymbol(V2892)
}
__typedArg0 := V2892
return Call(__e, PrimFunc(symsymbol_2), __typedArg0)
})()

if True == tmp7543 {
tmp7540 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(V2892)
}
__typedArg0 := V2892
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp7540)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp7540
return Call(__e, PrimFunc(symnot), __typedArg0)
})() {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


} else {
__e.Return(False)
return
}


}, 1)

tmp7544 := Call(__e, ns2_1set, symshen_4lowercase_1symbol_2, tmp7538)


_ = tmp7544

tmp7545 := MakeNative(func(__e *ControlFlow) {
V2893 := __e.Get(1)
_ = V2893
tmp7575 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2893)
}
__typedArg0 := V2893
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7566 Obj

if True == tmp7575 {
tmp7573 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2893)
}
__typedArg0 := V2893
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7574 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7573)
}
__typedArg0 := tmp7573
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7568 Obj

if True == tmp7574 {
tmp7570 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2893)
}
__typedArg0 := V2893
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7571 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7570)
}
__typedArg0 := tmp7570
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7572 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp7571)
}
__typedArg0 := Nil
__typedArg1 := tmp7571
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7569 Obj

if True == tmp7572 {
ifres7569 = True


} else {
ifres7569 = False


}

ifres7568 = ifres7569


} else {
ifres7568 = False


}

var ifres7567 Obj

if True == ifres7568 {
ifres7567 = True


} else {
ifres7567 = False


}

ifres7566 = ifres7567


} else {
ifres7566 = False


}

if True == ifres7566 {
__e.Return(V2893)
return
} else {
tmp7564 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2893)
}
__typedArg0 := V2893
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7555 Obj

if True == tmp7564 {
tmp7562 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2893)
}
__typedArg0 := V2893
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7563 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7562)
}
__typedArg0 := tmp7562
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7557 Obj

if True == tmp7563 {
tmp7559 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2893)
}
__typedArg0 := V2893
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7560 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7559)
}
__typedArg0 := tmp7559
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7561 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7560)
}
__typedArg0 := tmp7560
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7558 Obj

if True == tmp7561 {
ifres7558 = True


} else {
ifres7558 = False


}

ifres7557 = ifres7558


} else {
ifres7557 = False


}

var ifres7556 Obj

if True == ifres7557 {
ifres7556 = True


} else {
ifres7556 = False


}

ifres7555 = ifres7556


} else {
ifres7555 = False


}

if True == ifres7555 {
tmp7546 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2893)
}
__typedArg0 := V2893
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7547 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2893)
}
__typedArg0 := V2893
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7548 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7547)
}
__typedArg0 := tmp7547
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7549 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7548, Nil)
}
__typedArg0 := tmp7548
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7550 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7546, tmp7549)
}
__typedArg0 := tmp7546
__typedArg1 := tmp7549
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7551 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2893)
}
__typedArg0 := V2893
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7552 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7551)
}
__typedArg0 := tmp7551
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7553 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7550, tmp7552)
}
__typedArg0 := tmp7550
__typedArg1 := tmp7552
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4simple_1curry), tmp7553)
return


} else {
__e.Return(V2893)
return
}


}


}, 1)

tmp7576 := Call(__e, ns2_1set, symshen_4simple_1curry, tmp7545)


_ = tmp7576

tmp7577 := MakeNative(func(__e *ControlFlow) {
V2894 := __e.Get(1)
_ = V2894
__e.TailApply(PrimFunc(symfn), V2894)
return
}, 1)

tmp7578 := Call(__e, ns2_1set, symfunction, tmp7577)


_ = tmp7578

tmp7579 := MakeNative(func(__e *ControlFlow) {
V2895 := __e.Get(1)
_ = V2895
tmp7588 := Call(__e, PrimFunc(symarity), V2895)


tmp7589 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp7588, MakeNumber(0))
}
__typedArg0 := tmp7588
__typedArg1 := MakeNumber(0)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp7589 {
__e.TailApply(V2895)
return
} else {
tmp7580 := MakeNative(func(__e *ControlFlow) {
W2896 := __e.Get(1)
_ = W2896
tmp7584 := Call(__e, PrimFunc(symempty_2), W2896)


if True == tmp7584 {
tmp7581 := Call(__e, PrimFunc(symshen_4app), V2895, MakeString(" is undefined\n"), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("fn: "))
__typedS1, __typedOK1 := TypedString(tmp7581)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("fn: ")
__typedArg1 := tmp7581
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("fn: "))
__typedS1, __typedOK1 := TypedString(tmp7581)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("fn: ")
__typedArg1 := tmp7581
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W2896)
}
__typedArg0 := W2896
return Call(__e, PrimFunc(symtl), __typedArg0)
})())
return
}


}, 1)

tmp7585 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dlambdatable_d)
}
__typedArg0 := symshen_4_dlambdatable_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp7586 := Call(__e, PrimFunc(symassoc), V2895, tmp7585)


__e.TailApply(tmp7580, tmp7586)
return


}


}, 1)

tmp7590 := Call(__e, ns2_1set, symfn, tmp7579)


_ = tmp7590

tmp7591 := MakeNative(func(__e *ControlFlow) {
V2899 := __e.Get(1)
_ = V2899
tmp7621 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2899)
}
__typedArg0 := V2899
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7608 Obj

if True == tmp7621 {
tmp7619 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2899)
}
__typedArg0 := V2899
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7620 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symfn, tmp7619)
}
__typedArg0 := symfn
__typedArg1 := tmp7619
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7610 Obj

if True == tmp7620 {
tmp7617 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2899)
}
__typedArg0 := V2899
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7618 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7617)
}
__typedArg0 := tmp7617
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7612 Obj

if True == tmp7618 {
tmp7614 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2899)
}
__typedArg0 := V2899
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7615 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7614)
}
__typedArg0 := tmp7614
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7616 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp7615)
}
__typedArg0 := Nil
__typedArg1 := tmp7615
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7613 Obj

if True == tmp7616 {
ifres7613 = True


} else {
ifres7613 = False


}

ifres7612 = ifres7613


} else {
ifres7612 = False


}

var ifres7611 Obj

if True == ifres7612 {
ifres7611 = True


} else {
ifres7611 = False


}

ifres7610 = ifres7611


} else {
ifres7610 = False


}

var ifres7609 Obj

if True == ifres7610 {
ifres7609 = True


} else {
ifres7609 = False


}

ifres7608 = ifres7609


} else {
ifres7608 = False


}

if True == ifres7608 {
__e.Return(True)
return
} else {
tmp7606 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2899)
}
__typedArg0 := V2899
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7593 Obj

if True == tmp7606 {
tmp7604 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2899)
}
__typedArg0 := V2899
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7605 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symfunction, tmp7604)
}
__typedArg0 := symfunction
__typedArg1 := tmp7604
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7595 Obj

if True == tmp7605 {
tmp7602 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2899)
}
__typedArg0 := V2899
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7603 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7602)
}
__typedArg0 := tmp7602
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7597 Obj

if True == tmp7603 {
tmp7599 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2899)
}
__typedArg0 := V2899
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7600 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7599)
}
__typedArg0 := tmp7599
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7601 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp7600)
}
__typedArg0 := Nil
__typedArg1 := tmp7600
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7598 Obj

if True == tmp7601 {
ifres7598 = True


} else {
ifres7598 = False


}

ifres7597 = ifres7598


} else {
ifres7597 = False


}

var ifres7596 Obj

if True == ifres7597 {
ifres7596 = True


} else {
ifres7596 = False


}

ifres7595 = ifres7596


} else {
ifres7595 = False


}

var ifres7594 Obj

if True == ifres7595 {
ifres7594 = True


} else {
ifres7594 = False


}

ifres7593 = ifres7594


} else {
ifres7593 = False


}

if True == ifres7593 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp7622 := Call(__e, ns2_1set, symshen_4fn_1call_2, tmp7591)


_ = tmp7622

tmp7623 := MakeNative(func(__e *ControlFlow) {
V2900 := __e.Get(1)
_ = V2900
tmp7664 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2900)
}
__typedArg0 := V2900
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7651 Obj

if True == tmp7664 {
tmp7662 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2900)
}
__typedArg0 := V2900
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7663 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symfunction, tmp7662)
}
__typedArg0 := symfunction
__typedArg1 := tmp7662
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7653 Obj

if True == tmp7663 {
tmp7660 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2900)
}
__typedArg0 := V2900
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7661 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7660)
}
__typedArg0 := tmp7660
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7655 Obj

if True == tmp7661 {
tmp7657 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2900)
}
__typedArg0 := V2900
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7658 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7657)
}
__typedArg0 := tmp7657
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7659 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp7658)
}
__typedArg0 := Nil
__typedArg1 := tmp7658
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7656 Obj

if True == tmp7659 {
ifres7656 = True


} else {
ifres7656 = False


}

ifres7655 = ifres7656


} else {
ifres7655 = False


}

var ifres7654 Obj

if True == ifres7655 {
ifres7654 = True


} else {
ifres7654 = False


}

ifres7653 = ifres7654


} else {
ifres7653 = False


}

var ifres7652 Obj

if True == ifres7653 {
ifres7652 = True


} else {
ifres7652 = False


}

ifres7651 = ifres7652


} else {
ifres7651 = False


}

if True == ifres7651 {
tmp7624 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2900)
}
__typedArg0 := V2900
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7625 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfn, tmp7624)
}
__typedArg0 := symfn
__typedArg1 := tmp7624
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4fn_1call), tmp7625)
return


} else {
tmp7649 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2900)
}
__typedArg0 := V2900
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7636 Obj

if True == tmp7649 {
tmp7647 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2900)
}
__typedArg0 := V2900
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7648 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symfn, tmp7647)
}
__typedArg0 := symfn
__typedArg1 := tmp7647
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7638 Obj

if True == tmp7648 {
tmp7645 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2900)
}
__typedArg0 := V2900
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7646 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7645)
}
__typedArg0 := tmp7645
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7640 Obj

if True == tmp7646 {
tmp7642 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2900)
}
__typedArg0 := V2900
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7643 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7642)
}
__typedArg0 := tmp7642
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7644 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp7643)
}
__typedArg0 := Nil
__typedArg1 := tmp7643
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7641 Obj

if True == tmp7644 {
ifres7641 = True


} else {
ifres7641 = False


}

ifres7640 = ifres7641


} else {
ifres7640 = False


}

var ifres7639 Obj

if True == ifres7640 {
ifres7639 = True


} else {
ifres7639 = False


}

ifres7638 = ifres7639


} else {
ifres7638 = False


}

var ifres7637 Obj

if True == ifres7638 {
ifres7637 = True


} else {
ifres7637 = False


}

ifres7636 = ifres7637


} else {
ifres7636 = False


}

if True == ifres7636 {
tmp7626 := MakeNative(func(__e *ControlFlow) {
W2901 := __e.Get(1)
_ = W2901
tmp7631 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W2901, MakeNumber(-1))
}
__typedArg0 := W2901
__typedArg1 := MakeNumber(-1)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp7631 {
__e.Return(V2900)
return
} else {
tmp7629 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W2901, MakeNumber(0))
}
__typedArg0 := W2901
__typedArg1 := MakeNumber(0)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp7629 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2900)
}
__typedArg0 := V2900
return Call(__e, PrimFunc(symtl), __typedArg0)
})())
return
} else {
tmp7627 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2900)
}
__typedArg0 := V2900
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4lambda_1function), tmp7627, W2901)
return


}


}


}, 1)

tmp7632 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2900)
}
__typedArg0 := V2900
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7633 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7632)
}
__typedArg0 := tmp7632
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7634 := Call(__e, PrimFunc(symarity), tmp7633)


__e.TailApply(tmp7626, tmp7634)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("partial function shen.fn-call"))
}
__typedArg0 := MakeString("partial function shen.fn-call")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 1)

tmp7665 := Call(__e, ns2_1set, symshen_4fn_1call, tmp7623)


_ = tmp7665

tmp7666 := MakeNative(func(__e *ControlFlow) {
V2902 := __e.Get(1)
_ = V2902
V2903 := __e.Get(2)
_ = V2903
V2904 := __e.Get(3)
_ = V2904
tmp7667 := MakeNative(func(__e *ControlFlow) {
W2905 := __e.Get(1)
_ = W2905
tmp7668 := MakeNative(func(__e *ControlFlow) {
W2906 := __e.Get(1)
_ = W2906
__e.Return(W2905)
return
}, 1)

var ifres7674 Obj

if True == W2905 {
tmp7682 := Call(__e, PrimFunc(symshen_4loading_2))


var ifres7676 Obj

if True == tmp7682 {
tmp7678 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1, Nil)
}
__typedArg0 := sym_1
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7679 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_7, tmp7678)
}
__typedArg0 := sym_7
__typedArg1 := tmp7678
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7680 := Call(__e, PrimFunc(symelement_2), V2902, tmp7679)


tmp7681 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp7680)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp7680
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres7677 Obj

if True == tmp7681 {
ifres7677 = True


} else {
ifres7677 = False


}

ifres7676 = ifres7677


} else {
ifres7676 = False


}

var ifres7675 Obj

if True == ifres7676 {
ifres7675 = True


} else {
ifres7675 = False


}

ifres7674 = ifres7675


} else {
ifres7674 = False


}

var ifres7669 Obj

if True == ifres7674 {
tmp7670 := Call(__e, PrimFunc(symshen_4app), V2902, MakeString("\n"), symshen_4a)


tmp7671 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("partial application of "))
__typedS1, __typedOK1 := TypedString(tmp7670)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("partial application of ")
__typedArg1 := tmp7670
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp7672 := Call(__e, PrimFunc(symstoutput))


tmp7673 := Call(__e, PrimFunc(sympr), tmp7671, tmp7672)


ifres7669 = tmp7673


} else {
ifres7669 = symshen_4skip


}

__e.TailApply(tmp7668, ifres7669)
return


}, 1)

__e.TailApply(tmp7667, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_6) {
__typedN0, __typedOK0 := TypedFloat64(V2903)
__typedN1, __typedOK1 := TypedFloat64(V2904)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_6) {
return TypedMaterializeBoolean((__typedN0 > __typedN1))
}}
__typedArg0 := V2903
__typedArg1 := V2904
return Call(__e, PrimFunc(sym_6), __typedArg0, __typedArg1)
})())
return


}, 3)

tmp7684 := Call(__e, ns2_1set, symshen_4partial_1application_d_2, tmp7666)


_ = tmp7684

tmp7685 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dloading_2_d)
}
__typedArg0 := symshen_4_dloading_2_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp7686 := Call(__e, ns2_1set, symshen_4loading_2, tmp7685)


_ = tmp7686

tmp7687 := MakeNative(func(__e *ControlFlow) {
V2911 := __e.Get(1)
_ = V2911
V2912 := __e.Get(2)
_ = V2912
V2913 := __e.Get(3)
_ = V2913
tmp7705 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeNumber(-1), V2912)
}
__typedArg0 := MakeNumber(-1)
__typedArg1 := V2912
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp7705 {
__e.Return(False)
return
} else {
tmp7688 := MakeNative(func(__e *ControlFlow) {
W2914 := __e.Get(1)
_ = W2914
tmp7689 := MakeNative(func(__e *ControlFlow) {
W2915 := __e.Get(1)
_ = W2915
__e.Return(W2914)
return
}, 1)

var ifres7700 Obj

if True == W2914 {
tmp7702 := Call(__e, PrimFunc(symshen_4loading_2))


var ifres7701 Obj

if True == tmp7702 {
ifres7701 = True


} else {
ifres7701 = False


}

ifres7700 = ifres7701


} else {
ifres7700 = False


}

var ifres7690 Obj

if True == ifres7700 {
tmp7692 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V2913, MakeNumber(1))
}
__typedArg0 := V2913
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7691 Obj

if True == tmp7692 {
ifres7691 = MakeString("")


} else {
ifres7691 = MakeString("s")


}

tmp7693 := Call(__e, PrimFunc(symshen_4app), ifres7691, MakeString("\n"), symshen_4a)


tmp7695 := Call(__e, PrimFunc(symshen_4app), V2913, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(" argument"))
__typedS1, __typedOK1 := TypedString(tmp7693)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(" argument")
__typedArg1 := tmp7693
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


tmp7697 := Call(__e, PrimFunc(symshen_4app), V2911, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(" might not like "))
__typedS1, __typedOK1 := TypedString(tmp7695)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(" might not like ")
__typedArg1 := tmp7695
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


tmp7698 := Call(__e, PrimFunc(symstoutput))


tmp7699 := Call(__e, PrimFunc(sympr), tmp7697, tmp7698)


ifres7690 = tmp7699


} else {
ifres7690 = symshen_4skip


}

__e.TailApply(tmp7689, ifres7690)
return


}, 1)

__e.TailApply(tmp7688, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5) {
__typedN0, __typedOK0 := TypedFloat64(V2912)
__typedN1, __typedOK1 := TypedFloat64(V2913)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_5) {
return TypedMaterializeBoolean((__typedN0 < __typedN1))
}}
__typedArg0 := V2912
__typedArg1 := V2913
return Call(__e, PrimFunc(sym_5), __typedArg0, __typedArg1)
})())
return


}


}, 3)

__e.TailApply(ns2_1set, symshen_4overapplication_2, tmp7687)
return




}, 0)

