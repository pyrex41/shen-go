package main

import . "github.com/pyrex41/shen-go/kl"

var ReaderMain = MakeNative(func(__e *ControlFlow) {
tmp5335 := MakeNative(func(__e *ControlFlow) {
V2196 := __e.Get(1)
_ = V2196
tmp5336 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symread_1file_1as_1bytelist) {
return PrimReadFileAsByteList(V2196)
}
__typedArg0 := V2196
return Call(__e, PrimFunc(symread_1file_1as_1bytelist), __typedArg0)
})()

let__4910 := tmp5336
_ = let__4910

tmp5337 := MakeNative(func(__e *ControlFlow) {
tmp5338 := MakeNative(func(__e *ControlFlow) {
Z2199 := __e.Get(1)
_ = Z2199
__e.TailApply(PrimFunc(symshen_4_5s_1exprs_6), Z2199)
return
}, 1)

__e.TailApply(PrimFunc(symcompile), tmp5338, let__4910)
return


}, 0)

tmp5339 := MakeNative(func(__e *ControlFlow) {
Z2200 := __e.Get(1)
_ = Z2200
tmp5340 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dresidue_d)
}
__typedArg0 := symshen_4_dresidue_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4reader_1error), tmp5340)
return


}, 1)

tmp5341 := Call(__e, try_1catch, tmp5337, tmp5339)


let__4911 := tmp5341
_ = let__4911

tmp5342 := Call(__e, PrimFunc(symshen_4process_1sexprs), let__4911)


let__4912 := tmp5342
_ = let__4912

__e.Return(let__4912)
return


}, 1)

tmp5343 := Call(__e, ns2_1set, symread_1file, tmp5335)


_ = tmp5343

tmp5344 := MakeNative(func(__e *ControlFlow) {
V2202 := __e.Get(1)
_ = V2202
tmp5345 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dmaximum_1print_1sequence_1size_d)
}
__typedArg0 := sym_dmaximum_1print_1sequence_1size_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp5346 := Call(__e, PrimFunc(symshen_4reader_1error_1message), tmp5345, MakeInteger(0), V2202)


tmp5348 := Call(__e, PrimFunc(symshen_4proc_1nl), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("reader error near here: "))
__typedS1, __typedOK1 := TypedString(tmp5346)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("reader error near here: ")
__typedArg1 := tmp5346
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(tmp5348)
}
__typedArg0 := tmp5348
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}, 1)

tmp5349 := Call(__e, ns2_1set, symshen_4reader_1error, tmp5344)


_ = tmp5349

tmp5350 := MakeNative(func(__e *ControlFlow) {
V2210 := __e.Get(1)
_ = V2210
V2211 := __e.Get(2)
_ = V2211
V2212 := __e.Get(3)
_ = V2212
tmp5361 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V2212)
}
__typedArg0 := Nil
__typedArg1 := V2212
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp5361 {
__e.Return(MakeString(""))
return
} else {
tmp5359 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V2210, V2211)
}
__typedArg0 := V2210
__typedArg1 := V2211
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp5359 {
__e.Return(MakeString(""))
return
} else {
tmp5357 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2212)
}
__typedArg0 := V2212
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp5357 {
tmp5351 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2212)
}
__typedArg0 := V2212
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5352 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symn_1_6string) {
return PrimNumberToString(tmp5351)
}
__typedArg0 := tmp5351
return Call(__e, PrimFunc(symn_1_6string), __typedArg0)
})()

tmp5353 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V2211)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V2211
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()

tmp5354 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2212)
}
__typedArg0 := V2212
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5355 := Call(__e, PrimFunc(symshen_4reader_1error_1message), V2210, tmp5353, tmp5354)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(tmp5352)
__typedS1, __typedOK1 := TypedString(tmp5355)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := tmp5352
__typedArg1 := tmp5355
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

tmp5362 := Call(__e, ns2_1set, symshen_4reader_1error_1message, tmp5350)


_ = tmp5362

tmp5363 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dit_d)
}
__typedArg0 := symshen_4_dit_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp5364 := Call(__e, ns2_1set, symit, tmp5363)


_ = tmp5364

tmp5365 := MakeNative(func(__e *ControlFlow) {
V2213 := __e.Get(1)
_ = V2213
tmp5366 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symopen) {
return PrimOpenStream(V2213, symin)
}
__typedArg0 := V2213
__typedArg1 := symin
return Call(__e, PrimFunc(symopen), __typedArg0, __typedArg1)
})()

let__4913 := tmp5366
_ = let__4913

tmp5367 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symread_1byte) {
return PrimReadByte(let__4913)
}
__typedArg0 := let__4913
return Call(__e, PrimFunc(symread_1byte), __typedArg0)
})()

let__4914 := tmp5367
_ = let__4914

tmp5368 := Call(__e, PrimFunc(symshen_4read_1file_1as_1bytelist_1help), let__4913, let__4914, Nil)


let__4915 := tmp5368
_ = let__4915

tmp5369 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symclose) {
return PrimCloseStream(let__4913)
}
__typedArg0 := let__4913
return Call(__e, PrimFunc(symclose), __typedArg0)
})()

let__4916 := tmp5369
_ = let__4916

__e.TailApply(PrimFunc(symreverse), let__4915)
return


}, 1)

tmp5370 := Call(__e, ns2_1set, symread_1file_1as_1bytelist, tmp5365)


_ = tmp5370

tmp5371 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__self3 := __e.Get(3)
__selftop:
V2218 := __self1
_ = V2218
V2219 := __self2
_ = V2219
V2220 := __self3
_ = V2220
tmp5375 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(-1), V2219)
}
__typedArg0 := MakeInteger(-1)
__typedArg1 := V2219
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp5375 {
__e.Return(V2220)
return
} else {
tmp5372 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symread_1byte) {
return PrimReadByte(V2218)
}
__typedArg0 := V2218
return Call(__e, PrimFunc(symread_1byte), __typedArg0)
})()

tmp5373 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V2219, V2220)
}
__typedArg0 := V2219
__typedArg1 := V2220
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

if PrimFunc(symshen_4read_1file_1as_1bytelist_1help) == __self {
__self1, __self2, __self3 = V2218, tmp5372, tmp5373
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4read_1file_1as_1bytelist_1help), V2218, tmp5372, tmp5373)
return


}


}, 3)

tmp5376 := Call(__e, ns2_1set, symshen_4read_1file_1as_1bytelist_1help, tmp5371)


_ = tmp5376

tmp5377 := MakeNative(func(__e *ControlFlow) {
V2221 := __e.Get(1)
_ = V2221
tmp5378 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symopen) {
return PrimOpenStream(V2221, symin)
}
__typedArg0 := V2221
__typedArg1 := symin
return Call(__e, PrimFunc(symopen), __typedArg0, __typedArg1)
})()

let__4917 := tmp5378
_ = let__4917

tmp5379 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symread_1byte) {
return PrimReadByte(let__4917)
}
__typedArg0 := let__4917
return Call(__e, PrimFunc(symread_1byte), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4rfas_1h), let__4917, tmp5379, MakeString(""))
return


}, 1)

tmp5380 := Call(__e, ns2_1set, symread_1file_1as_1string, tmp5377)


_ = tmp5380

tmp5381 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__self3 := __e.Get(3)
__selftop:
V2223 := __self1
_ = V2223
V2224 := __self2
_ = V2224
V2225 := __self3
_ = V2225
tmp5387 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(-1), V2224)
}
__typedArg0 := MakeInteger(-1)
__typedArg1 := V2224
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp5387 {
tmp5382 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symclose) {
return PrimCloseStream(V2223)
}
__typedArg0 := V2223
return Call(__e, PrimFunc(symclose), __typedArg0)
})()

_ = tmp5382

__e.Return(V2225)
return


} else {
tmp5383 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symread_1byte) {
return PrimReadByte(V2223)
}
__typedArg0 := V2223
return Call(__e, PrimFunc(symread_1byte), __typedArg0)
})()

tmp5384 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symn_1_6string) {
return PrimNumberToString(V2224)
}
__typedArg0 := V2224
return Call(__e, PrimFunc(symn_1_6string), __typedArg0)
})()

if PrimFunc(symshen_4rfas_1h) == __self {
__self1, __self2, __self3 = V2223, tmp5383, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(V2225)
__typedS1, __typedOK1 := TypedString(tmp5384)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := V2225
__typedArg1 := tmp5384
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4rfas_1h), V2223, tmp5383, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(V2225)
__typedS1, __typedOK1 := TypedString(tmp5384)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := V2225
__typedArg1 := tmp5384
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
return


}


}, 3)

tmp5388 := Call(__e, ns2_1set, symshen_4rfas_1h, tmp5381)


_ = tmp5388

tmp5389 := MakeNative(func(__e *ControlFlow) {
V2226 := __e.Get(1)
_ = V2226
tmp5390 := Call(__e, PrimFunc(symread), V2226)


__e.TailApply(PrimFunc(symeval_1kl), tmp5390)
return


}, 1)

tmp5391 := Call(__e, ns2_1set, syminput, tmp5389)


_ = tmp5391

tmp5392 := MakeNative(func(__e *ControlFlow) {
V2227 := __e.Get(1)
_ = V2227
V2228 := __e.Get(2)
_ = V2228
tmp5393 := Call(__e, PrimFunc(symshen_4monotype), V2227)


let__4918 := tmp5393
_ = let__4918

tmp5394 := Call(__e, PrimFunc(symread), V2228)


let__4919 := tmp5394
_ = let__4919

tmp5400 := Call(__e, PrimFunc(symshen_4typecheck), let__4919, V2227)


tmp5401 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(False, tmp5400)
}
__typedArg0 := False
__typedArg1 := tmp5400
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp5401 {
tmp5395 := Call(__e, PrimFunc(symshen_4app), V2227, MakeString("\n"), symshen_4r)


tmp5397 := Call(__e, PrimFunc(symshen_4app), let__4919, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(" is not of type "))
__typedS1, __typedOK1 := TypedString(tmp5395)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(" is not of type ")
__typedArg1 := tmp5395
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4r)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("type error: "))
__typedS1, __typedOK1 := TypedString(tmp5397)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("type error: ")
__typedArg1 := tmp5397
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("type error: "))
__typedS1, __typedOK1 := TypedString(tmp5397)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("type error: ")
__typedArg1 := tmp5397
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


} else {
__e.TailApply(PrimFunc(symeval_1kl), let__4919)
return
}


}, 2)

tmp5402 := Call(__e, ns2_1set, symshen_4input_1h_7, tmp5392)


_ = tmp5402

tmp5403 := MakeNative(func(__e *ControlFlow) {
V2231 := __e.Get(1)
_ = V2231
tmp5410 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2231)
}
__typedArg0 := V2231
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp5410 {
tmp5404 := MakeNative(func(__e *ControlFlow) {
Z2232 := __e.Get(1)
_ = Z2232
__e.TailApply(PrimFunc(symshen_4monotype), Z2232)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp5404, V2231)
return


} else {
tmp5408 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(V2231)
}
__typedArg0 := V2231
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

if True == tmp5408 {
tmp5405 := Call(__e, PrimFunc(symshen_4app), V2231, MakeString("\n"), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("input+ expects a monotype: not "))
__typedS1, __typedOK1 := TypedString(tmp5405)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("input+ expects a monotype: not ")
__typedArg1 := tmp5405
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("input+ expects a monotype: not "))
__typedS1, __typedOK1 := TypedString(tmp5405)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("input+ expects a monotype: not ")
__typedArg1 := tmp5405
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

tmp5411 := Call(__e, ns2_1set, symshen_4monotype, tmp5403)


_ = tmp5411

tmp5412 := MakeNative(func(__e *ControlFlow) {
V2233 := __e.Get(1)
_ = V2233
tmp5413 := Call(__e, PrimFunc(symshen_4my_1read_1byte), V2233)


tmp5414 := MakeNative(func(__e *ControlFlow) {
Z2234 := __e.Get(1)
_ = Z2234
__e.TailApply(PrimFunc(symshen_4return_2), Z2234)
return
}, 1)

__e.TailApply(PrimFunc(symshen_4read_1loop), V2233, tmp5413, Nil, tmp5414)
return


}, 1)

tmp5415 := Call(__e, ns2_1set, symlineread, tmp5412)


_ = tmp5415

tmp5416 := MakeNative(func(__e *ControlFlow) {
V2235 := __e.Get(1)
_ = V2235
tmp5417 := Call(__e, PrimFunc(symshen_4str_1_6bytes), V2235)


let__4920 := tmp5417
_ = let__4920

tmp5418 := MakeNative(func(__e *ControlFlow) {
Z2238 := __e.Get(1)
_ = Z2238
__e.TailApply(PrimFunc(symshen_4_5s_1exprs_6), Z2238)
return
}, 1)

tmp5419 := Call(__e, PrimFunc(symcompile), tmp5418, let__4920)


let__4921 := tmp5419
_ = let__4921

tmp5420 := Call(__e, PrimFunc(symshen_4process_1sexprs), let__4921)


let__4922 := tmp5420
_ = let__4922

__e.Return(let__4922)
return


}, 1)

tmp5421 := Call(__e, ns2_1set, symread_1from_1string, tmp5416)


_ = tmp5421

tmp5422 := MakeNative(func(__e *ControlFlow) {
V2240 := __e.Get(1)
_ = V2240
tmp5423 := Call(__e, PrimFunc(symshen_4str_1_6bytes), V2240)


let__4923 := tmp5423
_ = let__4923

tmp5424 := MakeNative(func(__e *ControlFlow) {
Z2243 := __e.Get(1)
_ = Z2243
__e.TailApply(PrimFunc(symshen_4_5s_1exprs_6), Z2243)
return
}, 1)

tmp5425 := Call(__e, PrimFunc(symcompile), tmp5424, let__4923)


let__4924 := tmp5425
_ = let__4924

__e.Return(let__4924)
return


}, 1)

tmp5426 := Call(__e, ns2_1set, symread_1from_1string_1unprocessed, tmp5422)


_ = tmp5426

tmp5427 := MakeNative(func(__e *ControlFlow) {
V2244 := __e.Get(1)
_ = V2244
tmp5435 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString(""), V2244)
}
__typedArg0 := MakeString("")
__typedArg1 := V2244
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp5435 {
__e.Return(Nil)
return
} else {
tmp5433 := Call(__e, PrimFunc(symshen_4_7string_2), V2244)


if True == tmp5433 {
tmp5428 := Call(__e, PrimFunc(symhdstr), V2244)


tmp5429 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_1_6n) {
return PrimStringToNumber(tmp5428)
}
__typedArg0 := tmp5428
return Call(__e, PrimFunc(symstring_1_6n), __typedArg0)
})()

tmp5431 := Call(__e, PrimFunc(symshen_4str_1_6bytes), (func() Obj {
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
return PrimCons(tmp5429, tmp5431)
}
__typedArg0 := tmp5429
__typedArg1 := tmp5431
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

tmp5436 := Call(__e, ns2_1set, symshen_4str_1_6bytes, tmp5427)


_ = tmp5436

tmp5437 := MakeNative(func(__e *ControlFlow) {
V2245 := __e.Get(1)
_ = V2245
tmp5438 := Call(__e, PrimFunc(symshen_4my_1read_1byte), V2245)


tmp5439 := MakeNative(func(__e *ControlFlow) {
Z2246 := __e.Get(1)
_ = Z2246
__e.TailApply(PrimFunc(symshen_4whitespace_2), Z2246)
return
}, 1)

tmp5440 := Call(__e, PrimFunc(symshen_4read_1loop), V2245, tmp5438, Nil, tmp5439)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp5440)
}
__typedArg0 := tmp5440
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return


}, 1)

tmp5441 := Call(__e, ns2_1set, symread, tmp5437)


_ = tmp5441

tmp5442 := MakeNative(func(__e *ControlFlow) {
V2247 := __e.Get(1)
_ = V2247
tmp5445 := Call(__e, PrimFunc(symshen_4char_1stinput_2), V2247)


if True == tmp5445 {
tmp5443 := Call(__e, PrimFunc(symshen_4read_1unit_1string), V2247)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_1_6n) {
return PrimStringToNumber(tmp5443)
}
__typedArg0 := tmp5443
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

tmp5446 := Call(__e, ns2_1set, symshen_4my_1read_1byte, tmp5442)


_ = tmp5446

tmp5447 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__self3 := __e.Get(3)
__self4 := __e.Get(4)
__selftop:
V2252 := __self1
_ = V2252
V2253 := __self2
_ = V2253
V2254 := __self3
_ = V2254
V2255 := __self4
_ = V2255
tmp5469 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(94), V2253)
}
__typedArg0 := MakeInteger(94)
__typedArg1 := V2253
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp5469 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("read aborted"))
}
__typedArg0 := MakeString("read aborted")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
} else {
tmp5467 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(-1), V2253)
}
__typedArg0 := MakeInteger(-1)
__typedArg1 := V2253
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp5467 {
tmp5450 := Call(__e, PrimFunc(symempty_2), V2254)


if True == tmp5450 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("error: empty stream"))
}
__typedArg0 := MakeString("error: empty stream")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
} else {
tmp5448 := MakeNative(func(__e *ControlFlow) {
Z2256 := __e.Get(1)
_ = Z2256
__e.TailApply(PrimFunc(symshen_4_5s_1exprs_6), Z2256)
return
}, 1)

__e.TailApply(PrimFunc(symcompile), tmp5448, V2254)
return


}


} else {
tmp5465 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(0), V2253)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := V2253
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp5465 {
tmp5451 := Call(__e, PrimFunc(symshen_4my_1read_1byte), V2252)


if PrimFunc(symshen_4read_1loop) == __self {
__self1, __self2, __self3, __self4 = V2252, tmp5451, V2254, V2255
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4read_1loop), V2252, tmp5451, V2254, V2255)
return


} else {
tmp5463 := Call(__e, V2255, V2253)


if True == tmp5463 {
tmp5452 := Call(__e, PrimFunc(symshen_4try_1parse), V2254)


let__4925 := tmp5452
_ = let__4925

tmp5458 := Call(__e, PrimFunc(symshen_4nothing_1doing_2), let__4925)


if True == tmp5458 {
tmp5453 := Call(__e, PrimFunc(symshen_4my_1read_1byte), V2252)


tmp5454 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V2253, Nil)
}
__typedArg0 := V2253
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp5455 := Call(__e, PrimFunc(symappend), V2254, tmp5454)


if PrimFunc(symshen_4read_1loop) == __self {
__self1, __self2, __self3, __self4 = V2252, tmp5453, tmp5455, V2255
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4read_1loop), V2252, tmp5453, tmp5455, V2255)
return


} else {
tmp5456 := Call(__e, PrimFunc(symshen_4record_1it), V2254)


_ = tmp5456

__e.Return(let__4925)
return


}


} else {
tmp5459 := Call(__e, PrimFunc(symshen_4my_1read_1byte), V2252)


tmp5460 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V2253, Nil)
}
__typedArg0 := V2253
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp5461 := Call(__e, PrimFunc(symappend), V2254, tmp5460)


if PrimFunc(symshen_4read_1loop) == __self {
__self1, __self2, __self3, __self4 = V2252, tmp5459, tmp5461, V2255
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4read_1loop), V2252, tmp5459, tmp5461, V2255)
return


}


}


}


}


}, 4)

tmp5470 := Call(__e, ns2_1set, symshen_4read_1loop, tmp5447)


_ = tmp5470

tmp5471 := MakeNative(func(__e *ControlFlow) {
V2258 := __e.Get(1)
_ = V2258
tmp5472 := MakeNative(func(__e *ControlFlow) {
tmp5473 := MakeNative(func(__e *ControlFlow) {
Z2260 := __e.Get(1)
_ = Z2260
__e.TailApply(PrimFunc(symshen_4_5s_1exprs_6), Z2260)
return
}, 1)

__e.TailApply(PrimFunc(symcompile), tmp5473, V2258)
return


}, 0)

tmp5474 := MakeNative(func(__e *ControlFlow) {
Z2261 := __e.Get(1)
_ = Z2261
__e.Return(symshen_4i_1failed_b)
return
}, 1)

tmp5475 := Call(__e, try_1catch, tmp5472, tmp5474)


let__4926 := tmp5475
_ = let__4926

tmp5477 := Call(__e, PrimFunc(symshen_4nothing_1doing_2), let__4926)


if True == tmp5477 {
__e.Return(symshen_4i_1failed_b)
return
} else {
__e.TailApply(PrimFunc(symshen_4process_1sexprs), let__4926)
return
}


}, 1)

tmp5478 := Call(__e, ns2_1set, symshen_4try_1parse, tmp5471)


_ = tmp5478

tmp5479 := MakeNative(func(__e *ControlFlow) {
V2264 := __e.Get(1)
_ = V2264
tmp5483 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4i_1failed_b, V2264)
}
__typedArg0 := symshen_4i_1failed_b
__typedArg1 := V2264
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp5483 {
__e.Return(True)
return
} else {
tmp5481 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V2264)
}
__typedArg0 := Nil
__typedArg1 := V2264
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp5481 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp5484 := Call(__e, ns2_1set, symshen_4nothing_1doing_2, tmp5479)


_ = tmp5484

tmp5485 := MakeNative(func(__e *ControlFlow) {
V2265 := __e.Get(1)
_ = V2265
tmp5486 := Call(__e, PrimFunc(symshen_4bytes_1_6string), V2265)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dit_d, tmp5486)
}
__typedArg0 := symshen_4_dit_d
__typedArg1 := tmp5486
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp5487 := Call(__e, ns2_1set, symshen_4record_1it, tmp5485)


_ = tmp5487

tmp5488 := MakeNative(func(__e *ControlFlow) {
V2266 := __e.Get(1)
_ = V2266
tmp5496 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V2266)
}
__typedArg0 := Nil
__typedArg1 := V2266
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp5496 {
__e.Return(MakeString(""))
return
} else {
tmp5494 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2266)
}
__typedArg0 := V2266
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp5494 {
tmp5489 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2266)
}
__typedArg0 := V2266
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5490 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symn_1_6string) {
return PrimNumberToString(tmp5489)
}
__typedArg0 := tmp5489
return Call(__e, PrimFunc(symn_1_6string), __typedArg0)
})()

tmp5491 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2266)
}
__typedArg0 := V2266
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5492 := Call(__e, PrimFunc(symshen_4bytes_1_6string), tmp5491)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(tmp5490)
__typedS1, __typedOK1 := TypedString(tmp5492)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := tmp5490
__typedArg1 := tmp5492
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

tmp5497 := Call(__e, ns2_1set, symshen_4bytes_1_6string, tmp5488)


_ = tmp5497

tmp5498 := MakeNative(func(__e *ControlFlow) {
V2267 := __e.Get(1)
_ = V2267
tmp5499 := Call(__e, PrimFunc(symshen_4unpackage_emacroexpand), V2267)


let__4927 := tmp5499
_ = let__4927

tmp5500 := Call(__e, PrimFunc(symshen_4find_1arities), let__4927)


let__4928 := tmp5500
_ = let__4928

tmp5501 := Call(__e, PrimFunc(symshen_4find_1types), let__4927)


let__4929 := tmp5501
_ = let__4929

tmp5502 := MakeNative(func(__e *ControlFlow) {
Z2271 := __e.Get(1)
_ = Z2271
__e.TailApply(PrimFunc(symshen_4process_1applications), Z2271, let__4929)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp5502, let__4927)
return


}, 1)

tmp5503 := Call(__e, ns2_1set, symshen_4process_1sexprs, tmp5498)


_ = tmp5503

tmp5504 := MakeNative(func(__e *ControlFlow) {
V2272 := __e.Get(1)
_ = V2272
tmp5526 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2272)
}
__typedArg0 := V2272
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres5517 Obj

if True == tmp5526 {
tmp5524 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2272)
}
__typedArg0 := V2272
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5525 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp5524)
}
__typedArg0 := tmp5524
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres5519 Obj

if True == tmp5525 {
tmp5521 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2272)
}
__typedArg0 := V2272
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5522 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp5523 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp5521, tmp5522)
}
__typedArg0 := tmp5521
__typedArg1 := tmp5522
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres5520 Obj

if True == tmp5523 {
ifres5520 = True


} else {
ifres5520 = False


}

ifres5519 = ifres5520


} else {
ifres5519 = False


}

var ifres5518 Obj

if True == ifres5519 {
ifres5518 = True


} else {
ifres5518 = False


}

ifres5517 = ifres5518


} else {
ifres5517 = False


}

if True == ifres5517 {
tmp5505 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2272)
}
__typedArg0 := V2272
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5506 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp5505)
}
__typedArg0 := tmp5505
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5507 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2272)
}
__typedArg0 := V2272
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5508 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp5507)
}
__typedArg0 := tmp5507
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5509 := Call(__e, PrimFunc(symshen_4find_1types), tmp5508)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp5506, tmp5509)
}
__typedArg0 := tmp5506
__typedArg1 := tmp5509
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp5515 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2272)
}
__typedArg0 := V2272
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp5515 {
tmp5510 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2272)
}
__typedArg0 := V2272
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5511 := Call(__e, PrimFunc(symshen_4find_1types), tmp5510)


tmp5512 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2272)
}
__typedArg0 := V2272
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5513 := Call(__e, PrimFunc(symshen_4find_1types), tmp5512)


__e.TailApply(PrimFunc(symappend), tmp5511, tmp5513)
return


} else {
__e.Return(Nil)
return
}


}


}, 1)

tmp5527 := Call(__e, ns2_1set, symshen_4find_1types, tmp5504)


_ = tmp5527

tmp5528 := MakeNative(func(__e *ControlFlow) {
V2275 := __e.Get(1)
_ = V2275
tmp5577 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2275)
}
__typedArg0 := V2275
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres5558 Obj

if True == tmp5577 {
tmp5575 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2275)
}
__typedArg0 := V2275
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5576 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefine, tmp5575)
}
__typedArg0 := symdefine
__typedArg1 := tmp5575
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres5560 Obj

if True == tmp5576 {
tmp5573 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2275)
}
__typedArg0 := V2275
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5574 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp5573)
}
__typedArg0 := tmp5573
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres5562 Obj

if True == tmp5574 {
tmp5570 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2275)
}
__typedArg0 := V2275
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5571 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp5570)
}
__typedArg0 := tmp5570
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5572 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp5571)
}
__typedArg0 := tmp5571
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres5564 Obj

if True == tmp5572 {
tmp5566 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2275)
}
__typedArg0 := V2275
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5567 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp5566)
}
__typedArg0 := tmp5566
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5568 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp5567)
}
__typedArg0 := tmp5567
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5569 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_i, tmp5568)
}
__typedArg0 := sym_i
__typedArg1 := tmp5568
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres5565 Obj

if True == tmp5569 {
ifres5565 = True


} else {
ifres5565 = False


}

ifres5564 = ifres5565


} else {
ifres5564 = False


}

var ifres5563 Obj

if True == ifres5564 {
ifres5563 = True


} else {
ifres5563 = False


}

ifres5562 = ifres5563


} else {
ifres5562 = False


}

var ifres5561 Obj

if True == ifres5562 {
ifres5561 = True


} else {
ifres5561 = False


}

ifres5560 = ifres5561


} else {
ifres5560 = False


}

var ifres5559 Obj

if True == ifres5560 {
ifres5559 = True


} else {
ifres5559 = False


}

ifres5558 = ifres5559


} else {
ifres5558 = False


}

if True == ifres5558 {
tmp5529 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2275)
}
__typedArg0 := V2275
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5530 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp5529)
}
__typedArg0 := tmp5529
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5531 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2275)
}
__typedArg0 := V2275
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5532 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp5531)
}
__typedArg0 := tmp5531
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5533 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2275)
}
__typedArg0 := V2275
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5534 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp5533)
}
__typedArg0 := tmp5533
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5535 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp5534)
}
__typedArg0 := tmp5534
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5536 := Call(__e, PrimFunc(symshen_4find_1arity), tmp5532, MakeInteger(1), tmp5535)


__e.TailApply(PrimFunc(symshen_4store_1arity), tmp5530, tmp5536)
return


} else {
tmp5556 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2275)
}
__typedArg0 := V2275
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres5548 Obj

if True == tmp5556 {
tmp5554 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2275)
}
__typedArg0 := V2275
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5555 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefine, tmp5554)
}
__typedArg0 := symdefine
__typedArg1 := tmp5554
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres5550 Obj

if True == tmp5555 {
tmp5552 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2275)
}
__typedArg0 := V2275
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5553 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp5552)
}
__typedArg0 := tmp5552
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres5551 Obj

if True == tmp5553 {
ifres5551 = True


} else {
ifres5551 = False


}

ifres5550 = ifres5551


} else {
ifres5550 = False


}

var ifres5549 Obj

if True == ifres5550 {
ifres5549 = True


} else {
ifres5549 = False


}

ifres5548 = ifres5549


} else {
ifres5548 = False


}

if True == ifres5548 {
tmp5537 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2275)
}
__typedArg0 := V2275
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5538 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp5537)
}
__typedArg0 := tmp5537
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5539 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2275)
}
__typedArg0 := V2275
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5540 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp5539)
}
__typedArg0 := tmp5539
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5541 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2275)
}
__typedArg0 := V2275
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5542 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp5541)
}
__typedArg0 := tmp5541
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5543 := Call(__e, PrimFunc(symshen_4find_1arity), tmp5540, MakeInteger(0), tmp5542)


__e.TailApply(PrimFunc(symshen_4store_1arity), tmp5538, tmp5543)
return


} else {
tmp5546 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2275)
}
__typedArg0 := V2275
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp5546 {
tmp5544 := MakeNative(func(__e *ControlFlow) {
Z2276 := __e.Get(1)
_ = Z2276
__e.TailApply(PrimFunc(symshen_4find_1arities), Z2276)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp5544, V2275)
return


} else {
__e.Return(symshen_4skip)
return
}


}


}


}, 1)

tmp5578 := Call(__e, ns2_1set, symshen_4find_1arities, tmp5528)


_ = tmp5578

tmp5579 := MakeNative(func(__e *ControlFlow) {
V2277 := __e.Get(1)
_ = V2277
V2278 := __e.Get(2)
_ = V2278
tmp5580 := Call(__e, PrimFunc(symarity), V2277)


let__4930 := tmp5580
_ = let__4930

tmp5591 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__4930, MakeInteger(-1))
}
__typedArg0 := let__4930
__typedArg1 := MakeInteger(-1)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp5591 {
__e.TailApply(PrimFunc(symshen_4execute_1store_1arity), V2277, V2278)
return
} else {
tmp5589 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__4930, V2278)
}
__typedArg0 := let__4930
__typedArg1 := V2278
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp5589 {
__e.Return(symshen_4skip)
return
} else {
tmp5587 := Call(__e, PrimFunc(symshen_4sysfunc_2), V2277)


if True == tmp5587 {
tmp5581 := Call(__e, PrimFunc(symshen_4app), V2277, MakeString(" is a system function\n"), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(tmp5581)
}
__typedArg0 := tmp5581
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


} else {
tmp5582 := Call(__e, PrimFunc(symshen_4app), V2277, MakeString(" may cause errors\n"), symshen_4a)


tmp5583 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("changing the arity of "))
__typedS1, __typedOK1 := TypedString(tmp5582)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("changing the arity of ")
__typedArg1 := tmp5582
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp5584 := Call(__e, PrimFunc(symstoutput))


tmp5585 := Call(__e, PrimFunc(sympr), tmp5583, tmp5584)


_ = tmp5585

__e.TailApply(PrimFunc(symshen_4execute_1store_1arity), V2277, V2278)
return


}


}


}


}, 2)

tmp5592 := Call(__e, ns2_1set, symshen_4store_1arity, tmp5579)


_ = tmp5592

tmp5593 := MakeNative(func(__e *ControlFlow) {
V2280 := __e.Get(1)
_ = V2280
V2281 := __e.Get(2)
_ = V2281
tmp5598 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(0), V2281)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := V2281
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp5598 {
tmp5594 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symput), V2280, symarity, MakeInteger(0), tmp5594)
return


} else {
tmp5595 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp5596 := Call(__e, PrimFunc(symput), V2280, symarity, V2281, tmp5595)


_ = tmp5596

__e.TailApply(PrimFunc(symshen_4update_1lambdatable), V2280, V2281)
return


}


}, 2)

tmp5599 := Call(__e, ns2_1set, symshen_4execute_1store_1arity, tmp5593)


_ = tmp5599

tmp5600 := MakeNative(func(__e *ControlFlow) {
V2282 := __e.Get(1)
_ = V2282
V2283 := __e.Get(2)
_ = V2283
tmp5601 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dlambdatable_d)
}
__typedArg0 := symshen_4_dlambdatable_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

let__4931 := tmp5601
_ = let__4931

tmp5602 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V2282, Nil)
}
__typedArg0 := V2282
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp5603 := Call(__e, PrimFunc(symshen_4lambda_1function), tmp5602, V2283)


tmp5604 := Call(__e, PrimFunc(symeval_1kl), tmp5603)


let__4932 := tmp5604
_ = let__4932

tmp5605 := Call(__e, PrimFunc(symshen_4assoc_1_6), V2282, let__4932, let__4931)


let__4933 := tmp5605
_ = let__4933

tmp5606 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dlambdatable_d, let__4933)
}
__typedArg0 := symshen_4_dlambdatable_d
__typedArg1 := let__4933
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

let__4934 := tmp5606
_ = let__4934

__e.Return(let__4934)
return


}, 2)

tmp5607 := Call(__e, ns2_1set, symshen_4update_1lambdatable, tmp5600)


_ = tmp5607

tmp5608 := MakeNative(func(__e *ControlFlow) {
V2290 := __e.Get(1)
_ = V2290
V2291 := __e.Get(2)
_ = V2291
tmp5624 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(0), V2291)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := V2291
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp5624 {
__e.Return(symshen_4skip)
return
} else {
tmp5622 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(1), V2291)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := V2291
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp5622 {
tmp5609 := Call(__e, PrimFunc(symgensym), symY)


let__4935 := tmp5609
_ = let__4935

tmp5610 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__4935, Nil)
}
__typedArg0 := let__4935
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp5611 := Call(__e, PrimFunc(symappend), V2290, tmp5610)


tmp5612 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp5611, Nil)
}
__typedArg0 := tmp5611
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp5613 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__4935, tmp5612)
}
__typedArg0 := let__4935
__typedArg1 := tmp5612
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp5613)
}
__typedArg0 := symlambda
__typedArg1 := tmp5613
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp5614 := Call(__e, PrimFunc(symgensym), symY)


let__4936 := tmp5614
_ = let__4936

tmp5615 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__4936, Nil)
}
__typedArg0 := let__4936
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp5616 := Call(__e, PrimFunc(symappend), V2290, tmp5615)


tmp5618 := Call(__e, PrimFunc(symshen_4lambda_1function), tmp5616, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(V2291)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := V2291
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})())


tmp5619 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp5618, Nil)
}
__typedArg0 := tmp5618
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp5620 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__4936, tmp5619)
}
__typedArg0 := let__4936
__typedArg1 := tmp5619
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp5620)
}
__typedArg0 := symlambda
__typedArg1 := tmp5620
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


}


}, 2)

tmp5625 := Call(__e, ns2_1set, symshen_4lambda_1function, tmp5608)


_ = tmp5625

tmp5626 := MakeNative(func(__e *ControlFlow) {
V2303 := __e.Get(1)
_ = V2303
V2304 := __e.Get(2)
_ = V2304
V2305 := __e.Get(3)
_ = V2305
tmp5649 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V2305)
}
__typedArg0 := Nil
__typedArg1 := V2305
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp5649 {
tmp5627 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V2303, V2304)
}
__typedArg0 := V2303
__typedArg1 := V2304
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp5627, Nil)
}
__typedArg0 := tmp5627
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp5647 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2305)
}
__typedArg0 := V2305
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres5638 Obj

if True == tmp5647 {
tmp5645 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2305)
}
__typedArg0 := V2305
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5646 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp5645)
}
__typedArg0 := tmp5645
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres5640 Obj

if True == tmp5646 {
tmp5642 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2305)
}
__typedArg0 := V2305
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5643 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp5642)
}
__typedArg0 := tmp5642
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5644 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V2303, tmp5643)
}
__typedArg0 := V2303
__typedArg1 := tmp5643
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres5641 Obj

if True == tmp5644 {
ifres5641 = True


} else {
ifres5641 = False


}

ifres5640 = ifres5641


} else {
ifres5640 = False


}

var ifres5639 Obj

if True == ifres5640 {
ifres5639 = True


} else {
ifres5639 = False


}

ifres5638 = ifres5639


} else {
ifres5638 = False


}

if True == ifres5638 {
tmp5628 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2305)
}
__typedArg0 := V2305
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5629 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp5628)
}
__typedArg0 := tmp5628
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5630 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp5629, V2304)
}
__typedArg0 := tmp5629
__typedArg1 := V2304
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp5631 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2305)
}
__typedArg0 := V2305
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp5630, tmp5631)
}
__typedArg0 := tmp5630
__typedArg1 := tmp5631
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp5636 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2305)
}
__typedArg0 := V2305
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp5636 {
tmp5632 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2305)
}
__typedArg0 := V2305
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5633 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2305)
}
__typedArg0 := V2305
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5634 := Call(__e, PrimFunc(symshen_4assoc_1_6), V2303, V2304, tmp5633)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp5632, tmp5634)
}
__typedArg0 := tmp5632
__typedArg1 := tmp5634
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

tmp5650 := Call(__e, ns2_1set, symshen_4assoc_1_6, tmp5626)


_ = tmp5650

tmp5651 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__self3 := __e.Get(3)
__selftop:
V2320 := __self1
_ = V2320
V2321 := __self2
_ = V2321
V2322 := __self3
_ = V2322
tmp5698 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(0), V2321)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := V2321
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres5691 Obj

if True == tmp5698 {
tmp5697 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2322)
}
__typedArg0 := V2322
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres5693 Obj

if True == tmp5697 {
tmp5695 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2322)
}
__typedArg0 := V2322
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5696 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp5695, sym_1_6)
}
__typedArg0 := tmp5695
__typedArg1 := sym_1_6
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres5694 Obj

if True == tmp5696 {
ifres5694 = True


} else {
ifres5694 = False


}

ifres5693 = ifres5694


} else {
ifres5693 = False


}

var ifres5692 Obj

if True == ifres5693 {
ifres5692 = True


} else {
ifres5692 = False


}

ifres5691 = ifres5692


} else {
ifres5691 = False


}

if True == ifres5691 {
__e.Return(MakeInteger(0))
return
} else {
tmp5689 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(0), V2321)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := V2321
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres5682 Obj

if True == tmp5689 {
tmp5688 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2322)
}
__typedArg0 := V2322
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres5684 Obj

if True == tmp5688 {
tmp5686 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2322)
}
__typedArg0 := V2322
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5687 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp5686, sym_5_1)
}
__typedArg0 := tmp5686
__typedArg1 := sym_5_1
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres5685 Obj

if True == tmp5687 {
ifres5685 = True


} else {
ifres5685 = False


}

ifres5684 = ifres5685


} else {
ifres5684 = False


}

var ifres5683 Obj

if True == ifres5684 {
ifres5683 = True


} else {
ifres5683 = False


}

ifres5682 = ifres5683


} else {
ifres5682 = False


}

if True == ifres5682 {
__e.Return(MakeInteger(0))
return
} else {
tmp5680 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(0), V2321)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := V2321
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres5677 Obj

if True == tmp5680 {
tmp5679 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2322)
}
__typedArg0 := V2322
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres5678 Obj

if True == tmp5679 {
ifres5678 = True


} else {
ifres5678 = False


}

ifres5677 = ifres5678


} else {
ifres5677 = False


}

if True == ifres5677 {
tmp5652 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2322)
}
__typedArg0 := V2322
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5653 := Call(__e, PrimFunc(symshen_4find_1arity), V2320, MakeInteger(0), tmp5652)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(1))
__typedN1, __typedOK1 := TypedFloat64(tmp5653)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp5653
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())
return


} else {
tmp5675 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(1), V2321)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := V2321
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres5668 Obj

if True == tmp5675 {
tmp5674 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2322)
}
__typedArg0 := V2322
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres5670 Obj

if True == tmp5674 {
tmp5672 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2322)
}
__typedArg0 := V2322
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5673 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_j, tmp5672)
}
__typedArg0 := sym_j
__typedArg1 := tmp5672
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres5671 Obj

if True == tmp5673 {
ifres5671 = True


} else {
ifres5671 = False


}

ifres5670 = ifres5671


} else {
ifres5670 = False


}

var ifres5669 Obj

if True == ifres5670 {
ifres5669 = True


} else {
ifres5669 = False


}

ifres5668 = ifres5669


} else {
ifres5668 = False


}

if True == ifres5668 {
tmp5654 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2322)
}
__typedArg0 := V2322
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4find_1arity) == __self {
__self1, __self2, __self3 = V2320, MakeInteger(0), tmp5654
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4find_1arity), V2320, MakeInteger(0), tmp5654)
return


} else {
tmp5666 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(1), V2321)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := V2321
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres5663 Obj

if True == tmp5666 {
tmp5665 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2322)
}
__typedArg0 := V2322
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres5664 Obj

if True == tmp5665 {
ifres5664 = True


} else {
ifres5664 = False


}

ifres5663 = ifres5664


} else {
ifres5663 = False


}

if True == ifres5663 {
tmp5655 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2322)
}
__typedArg0 := V2322
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4find_1arity) == __self {
__self1, __self2, __self3 = V2320, MakeInteger(1), tmp5655
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4find_1arity), V2320, MakeInteger(1), tmp5655)
return


} else {
tmp5661 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(1), V2321)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := V2321
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp5661 {
tmp5656 := Call(__e, PrimFunc(symshen_4app), V2320, MakeString(" definition: missing }\n"), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("syntax error in "))
__typedS1, __typedOK1 := TypedString(tmp5656)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("syntax error in ")
__typedArg1 := tmp5656
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("syntax error in "))
__typedS1, __typedOK1 := TypedString(tmp5656)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("syntax error in ")
__typedArg1 := tmp5656
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


} else {
tmp5658 := Call(__e, PrimFunc(symshen_4app), V2320, MakeString(" definition: missing -> or <-\n"), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("syntax error in "))
__typedS1, __typedOK1 := TypedString(tmp5658)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("syntax error in ")
__typedArg1 := tmp5658
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("syntax error in "))
__typedS1, __typedOK1 := TypedString(tmp5658)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("syntax error in ")
__typedArg1 := tmp5658
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

tmp5699 := Call(__e, ns2_1set, symshen_4find_1arity, tmp5651)


_ = tmp5699

tmp5700 := MakeNative(func(__e *ControlFlow) {
V2323 := __e.Get(1)
_ = V2323
tmp5701 := Call(__e, PrimFunc(symshen_4_5lsb_6), V2323)


let__4938 := tmp5701
_ = let__4938

tmp5725 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__4938)


var ifres5702 Obj

if True == tmp5725 {
tmp5703 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5702 = tmp5703


} else {
tmp5704 := Call(__e, PrimFunc(symshen_4in_1_6), let__4938)


let__4939 := tmp5704
_ = let__4939

tmp5705 := Call(__e, PrimFunc(symshen_4_5s_1exprs1_6), let__4939)


let__4940 := tmp5705
_ = let__4940

tmp5724 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__4940)


var ifres5706 Obj

if True == tmp5724 {
tmp5707 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5706 = tmp5707


} else {
tmp5708 := Call(__e, PrimFunc(symshen_4_5_1out), let__4940)


let__4941 := tmp5708
_ = let__4941

tmp5709 := Call(__e, PrimFunc(symshen_4in_1_6), let__4940)


let__4942 := tmp5709
_ = let__4942

tmp5710 := Call(__e, PrimFunc(symshen_4_5rsb_6), let__4942)


let__4943 := tmp5710
_ = let__4943

tmp5723 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__4943)


var ifres5711 Obj

if True == tmp5723 {
tmp5712 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5711 = tmp5712


} else {
tmp5713 := Call(__e, PrimFunc(symshen_4in_1_6), let__4943)


let__4944 := tmp5713
_ = let__4944

tmp5714 := Call(__e, PrimFunc(symshen_4_5s_1exprs2_6), let__4944)


let__4945 := tmp5714
_ = let__4945

tmp5722 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__4945)


var ifres5715 Obj

if True == tmp5722 {
tmp5716 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5715 = tmp5716


} else {
tmp5717 := Call(__e, PrimFunc(symshen_4_5_1out), let__4945)


let__4946 := tmp5717
_ = let__4946

tmp5718 := Call(__e, PrimFunc(symshen_4in_1_6), let__4945)


let__4947 := tmp5718
_ = let__4947

tmp5719 := Call(__e, PrimFunc(symshen_4cons_1form), let__4941)


tmp5720 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp5719, let__4946)
}
__typedArg0 := tmp5719
__typedArg1 := let__4946
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp5721 := Call(__e, PrimFunc(symshen_4comb), let__4947, tmp5720)


ifres5715 = tmp5721


}

ifres5711 = ifres5715


}

ifres5706 = ifres5711


}

ifres5702 = ifres5706


}

let__4937 := ifres5702
_ = let__4937

tmp5919 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__4937)


if True == tmp5919 {
tmp5726 := Call(__e, PrimFunc(symshen_4_5lrb_6), V2323)


let__4949 := tmp5726
_ = let__4949

tmp5749 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__4949)


var ifres5727 Obj

if True == tmp5749 {
tmp5728 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5727 = tmp5728


} else {
tmp5729 := Call(__e, PrimFunc(symshen_4in_1_6), let__4949)


let__4950 := tmp5729
_ = let__4950

tmp5730 := Call(__e, PrimFunc(symshen_4_5s_1exprs1_6), let__4950)


let__4951 := tmp5730
_ = let__4951

tmp5748 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__4951)


var ifres5731 Obj

if True == tmp5748 {
tmp5732 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5731 = tmp5732


} else {
tmp5733 := Call(__e, PrimFunc(symshen_4_5_1out), let__4951)


let__4952 := tmp5733
_ = let__4952

tmp5734 := Call(__e, PrimFunc(symshen_4in_1_6), let__4951)


let__4953 := tmp5734
_ = let__4953

tmp5735 := Call(__e, PrimFunc(symshen_4_5rrb_6), let__4953)


let__4954 := tmp5735
_ = let__4954

tmp5747 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__4954)


var ifres5736 Obj

if True == tmp5747 {
tmp5737 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5736 = tmp5737


} else {
tmp5738 := Call(__e, PrimFunc(symshen_4in_1_6), let__4954)


let__4955 := tmp5738
_ = let__4955

tmp5739 := Call(__e, PrimFunc(symshen_4_5s_1exprs2_6), let__4955)


let__4956 := tmp5739
_ = let__4956

tmp5746 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__4956)


var ifres5740 Obj

if True == tmp5746 {
tmp5741 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5740 = tmp5741


} else {
tmp5742 := Call(__e, PrimFunc(symshen_4_5_1out), let__4956)


let__4957 := tmp5742
_ = let__4957

tmp5743 := Call(__e, PrimFunc(symshen_4in_1_6), let__4956)


let__4958 := tmp5743
_ = let__4958

tmp5744 := Call(__e, PrimFunc(symshen_4add_1sexpr), let__4952, let__4957)


tmp5745 := Call(__e, PrimFunc(symshen_4comb), let__4958, tmp5744)


ifres5740 = tmp5745


}

ifres5736 = ifres5740


}

ifres5731 = ifres5736


}

ifres5727 = ifres5731


}

let__4948 := ifres5727
_ = let__4948

tmp5917 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__4948)


if True == tmp5917 {
tmp5750 := Call(__e, PrimFunc(symshen_4_5lcurly_6), V2323)


let__4960 := tmp5750
_ = let__4960

tmp5762 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__4960)


var ifres5751 Obj

if True == tmp5762 {
tmp5752 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5751 = tmp5752


} else {
tmp5753 := Call(__e, PrimFunc(symshen_4in_1_6), let__4960)


let__4961 := tmp5753
_ = let__4961

tmp5754 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), let__4961)


let__4962 := tmp5754
_ = let__4962

tmp5761 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__4962)


var ifres5755 Obj

if True == tmp5761 {
tmp5756 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5755 = tmp5756


} else {
tmp5757 := Call(__e, PrimFunc(symshen_4_5_1out), let__4962)


let__4963 := tmp5757
_ = let__4963

tmp5758 := Call(__e, PrimFunc(symshen_4in_1_6), let__4962)


let__4964 := tmp5758
_ = let__4964

tmp5759 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_i, let__4963)
}
__typedArg0 := sym_i
__typedArg1 := let__4963
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp5760 := Call(__e, PrimFunc(symshen_4comb), let__4964, tmp5759)


ifres5755 = tmp5760


}

ifres5751 = ifres5755


}

let__4959 := ifres5751
_ = let__4959

tmp5915 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__4959)


if True == tmp5915 {
tmp5763 := Call(__e, PrimFunc(symshen_4_5rcurly_6), V2323)


let__4966 := tmp5763
_ = let__4966

tmp5775 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__4966)


var ifres5764 Obj

if True == tmp5775 {
tmp5765 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5764 = tmp5765


} else {
tmp5766 := Call(__e, PrimFunc(symshen_4in_1_6), let__4966)


let__4967 := tmp5766
_ = let__4967

tmp5767 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), let__4967)


let__4968 := tmp5767
_ = let__4968

tmp5774 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__4968)


var ifres5768 Obj

if True == tmp5774 {
tmp5769 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5768 = tmp5769


} else {
tmp5770 := Call(__e, PrimFunc(symshen_4_5_1out), let__4968)


let__4969 := tmp5770
_ = let__4969

tmp5771 := Call(__e, PrimFunc(symshen_4in_1_6), let__4968)


let__4970 := tmp5771
_ = let__4970

tmp5772 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_j, let__4969)
}
__typedArg0 := sym_j
__typedArg1 := let__4969
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp5773 := Call(__e, PrimFunc(symshen_4comb), let__4970, tmp5772)


ifres5768 = tmp5773


}

ifres5764 = ifres5768


}

let__4965 := ifres5764
_ = let__4965

tmp5913 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__4965)


if True == tmp5913 {
tmp5776 := Call(__e, PrimFunc(symshen_4_5bar_6), V2323)


let__4972 := tmp5776
_ = let__4972

tmp5788 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__4972)


var ifres5777 Obj

if True == tmp5788 {
tmp5778 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5777 = tmp5778


} else {
tmp5779 := Call(__e, PrimFunc(symshen_4in_1_6), let__4972)


let__4973 := tmp5779
_ = let__4973

tmp5780 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), let__4973)


let__4974 := tmp5780
_ = let__4974

tmp5787 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__4974)


var ifres5781 Obj

if True == tmp5787 {
tmp5782 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5781 = tmp5782


} else {
tmp5783 := Call(__e, PrimFunc(symshen_4_5_1out), let__4974)


let__4975 := tmp5783
_ = let__4975

tmp5784 := Call(__e, PrimFunc(symshen_4in_1_6), let__4974)


let__4976 := tmp5784
_ = let__4976

tmp5785 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symbar_b, let__4975)
}
__typedArg0 := symbar_b
__typedArg1 := let__4975
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp5786 := Call(__e, PrimFunc(symshen_4comb), let__4976, tmp5785)


ifres5781 = tmp5786


}

ifres5777 = ifres5781


}

let__4971 := ifres5777
_ = let__4971

tmp5911 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__4971)


if True == tmp5911 {
tmp5789 := Call(__e, PrimFunc(symshen_4_5semicolon_6), V2323)


let__4978 := tmp5789
_ = let__4978

tmp5802 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__4978)


var ifres5790 Obj

if True == tmp5802 {
tmp5791 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5790 = tmp5791


} else {
tmp5792 := Call(__e, PrimFunc(symshen_4in_1_6), let__4978)


let__4979 := tmp5792
_ = let__4979

tmp5793 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), let__4979)


let__4980 := tmp5793
_ = let__4980

tmp5801 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__4980)


var ifres5794 Obj

if True == tmp5801 {
tmp5795 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5794 = tmp5795


} else {
tmp5796 := Call(__e, PrimFunc(symshen_4_5_1out), let__4980)


let__4981 := tmp5796
_ = let__4981

tmp5797 := Call(__e, PrimFunc(symshen_4in_1_6), let__4980)


let__4982 := tmp5797
_ = let__4982

tmp5798 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(";"))
}
__typedArg0 := MakeString(";")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp5799 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp5798, let__4981)
}
__typedArg0 := tmp5798
__typedArg1 := let__4981
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp5800 := Call(__e, PrimFunc(symshen_4comb), let__4982, tmp5799)


ifres5794 = tmp5800


}

ifres5790 = ifres5794


}

let__4977 := ifres5790
_ = let__4977

tmp5909 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__4977)


if True == tmp5909 {
tmp5803 := Call(__e, PrimFunc(symshen_4_5colon_6), V2323)


let__4984 := tmp5803
_ = let__4984

tmp5821 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__4984)


var ifres5804 Obj

if True == tmp5821 {
tmp5805 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5804 = tmp5805


} else {
tmp5806 := Call(__e, PrimFunc(symshen_4in_1_6), let__4984)


let__4985 := tmp5806
_ = let__4985

tmp5807 := Call(__e, PrimFunc(symshen_4_5equal_6), let__4985)


let__4986 := tmp5807
_ = let__4986

tmp5820 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__4986)


var ifres5808 Obj

if True == tmp5820 {
tmp5809 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5808 = tmp5809


} else {
tmp5810 := Call(__e, PrimFunc(symshen_4in_1_6), let__4986)


let__4987 := tmp5810
_ = let__4987

tmp5811 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), let__4987)


let__4988 := tmp5811
_ = let__4988

tmp5819 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__4988)


var ifres5812 Obj

if True == tmp5819 {
tmp5813 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5812 = tmp5813


} else {
tmp5814 := Call(__e, PrimFunc(symshen_4_5_1out), let__4988)


let__4989 := tmp5814
_ = let__4989

tmp5815 := Call(__e, PrimFunc(symshen_4in_1_6), let__4988)


let__4990 := tmp5815
_ = let__4990

tmp5816 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":="))
}
__typedArg0 := MakeString(":=")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp5817 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp5816, let__4989)
}
__typedArg0 := tmp5816
__typedArg1 := let__4989
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp5818 := Call(__e, PrimFunc(symshen_4comb), let__4990, tmp5817)


ifres5812 = tmp5818


}

ifres5808 = ifres5812


}

ifres5804 = ifres5808


}

let__4983 := ifres5804
_ = let__4983

tmp5907 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__4983)


if True == tmp5907 {
tmp5822 := Call(__e, PrimFunc(symshen_4_5colon_6), V2323)


let__4992 := tmp5822
_ = let__4992

tmp5835 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__4992)


var ifres5823 Obj

if True == tmp5835 {
tmp5824 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5823 = tmp5824


} else {
tmp5825 := Call(__e, PrimFunc(symshen_4in_1_6), let__4992)


let__4993 := tmp5825
_ = let__4993

tmp5826 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), let__4993)


let__4994 := tmp5826
_ = let__4994

tmp5834 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__4994)


var ifres5827 Obj

if True == tmp5834 {
tmp5828 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5827 = tmp5828


} else {
tmp5829 := Call(__e, PrimFunc(symshen_4_5_1out), let__4994)


let__4995 := tmp5829
_ = let__4995

tmp5830 := Call(__e, PrimFunc(symshen_4in_1_6), let__4994)


let__4996 := tmp5830
_ = let__4996

tmp5831 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp5832 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp5831, let__4995)
}
__typedArg0 := tmp5831
__typedArg1 := let__4995
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp5833 := Call(__e, PrimFunc(symshen_4comb), let__4996, tmp5832)


ifres5827 = tmp5833


}

ifres5823 = ifres5827


}

let__4991 := ifres5823
_ = let__4991

tmp5905 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__4991)


if True == tmp5905 {
tmp5836 := Call(__e, PrimFunc(symshen_4_5comma_6), V2323)


let__4998 := tmp5836
_ = let__4998

tmp5849 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__4998)


var ifres5837 Obj

if True == tmp5849 {
tmp5838 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5837 = tmp5838


} else {
tmp5839 := Call(__e, PrimFunc(symshen_4in_1_6), let__4998)


let__4999 := tmp5839
_ = let__4999

tmp5840 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), let__4999)


let__5000 := tmp5840
_ = let__5000

tmp5848 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5000)


var ifres5841 Obj

if True == tmp5848 {
tmp5842 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5841 = tmp5842


} else {
tmp5843 := Call(__e, PrimFunc(symshen_4_5_1out), let__5000)


let__5001 := tmp5843
_ = let__5001

tmp5844 := Call(__e, PrimFunc(symshen_4in_1_6), let__5000)


let__5002 := tmp5844
_ = let__5002

tmp5845 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(","))
}
__typedArg0 := MakeString(",")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp5846 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp5845, let__5001)
}
__typedArg0 := tmp5845
__typedArg1 := let__5001
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp5847 := Call(__e, PrimFunc(symshen_4comb), let__5002, tmp5846)


ifres5841 = tmp5847


}

ifres5837 = ifres5841


}

let__4997 := ifres5837
_ = let__4997

tmp5903 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__4997)


if True == tmp5903 {
tmp5850 := Call(__e, PrimFunc(symshen_4_5comment_6), V2323)


let__5004 := tmp5850
_ = let__5004

tmp5861 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5004)


var ifres5851 Obj

if True == tmp5861 {
tmp5852 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5851 = tmp5852


} else {
tmp5853 := Call(__e, PrimFunc(symshen_4in_1_6), let__5004)


let__5005 := tmp5853
_ = let__5005

tmp5854 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), let__5005)


let__5006 := tmp5854
_ = let__5006

tmp5860 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5006)


var ifres5855 Obj

if True == tmp5860 {
tmp5856 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5855 = tmp5856


} else {
tmp5857 := Call(__e, PrimFunc(symshen_4_5_1out), let__5006)


let__5007 := tmp5857
_ = let__5007

tmp5858 := Call(__e, PrimFunc(symshen_4in_1_6), let__5006)


let__5008 := tmp5858
_ = let__5008

tmp5859 := Call(__e, PrimFunc(symshen_4comb), let__5008, let__5007)


ifres5855 = tmp5859


}

ifres5851 = ifres5855


}

let__5003 := ifres5851
_ = let__5003

tmp5901 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5003)


if True == tmp5901 {
tmp5862 := Call(__e, PrimFunc(symshen_4_5atom_6), V2323)


let__5010 := tmp5862
_ = let__5010

tmp5875 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5010)


var ifres5863 Obj

if True == tmp5875 {
tmp5864 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5863 = tmp5864


} else {
tmp5865 := Call(__e, PrimFunc(symshen_4_5_1out), let__5010)


let__5011 := tmp5865
_ = let__5011

tmp5866 := Call(__e, PrimFunc(symshen_4in_1_6), let__5010)


let__5012 := tmp5866
_ = let__5012

tmp5867 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), let__5012)


let__5013 := tmp5867
_ = let__5013

tmp5874 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5013)


var ifres5868 Obj

if True == tmp5874 {
tmp5869 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5868 = tmp5869


} else {
tmp5870 := Call(__e, PrimFunc(symshen_4_5_1out), let__5013)


let__5014 := tmp5870
_ = let__5014

tmp5871 := Call(__e, PrimFunc(symshen_4in_1_6), let__5013)


let__5015 := tmp5871
_ = let__5015

tmp5872 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__5011, let__5014)
}
__typedArg0 := let__5011
__typedArg1 := let__5014
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp5873 := Call(__e, PrimFunc(symshen_4comb), let__5015, tmp5872)


ifres5868 = tmp5873


}

ifres5863 = ifres5868


}

let__5009 := ifres5863
_ = let__5009

tmp5899 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5009)


if True == tmp5899 {
tmp5876 := Call(__e, PrimFunc(symshen_4_5whitespaces_6), V2323)


let__5017 := tmp5876
_ = let__5017

tmp5887 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5017)


var ifres5877 Obj

if True == tmp5887 {
tmp5878 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5877 = tmp5878


} else {
tmp5879 := Call(__e, PrimFunc(symshen_4in_1_6), let__5017)


let__5018 := tmp5879
_ = let__5018

tmp5880 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), let__5018)


let__5019 := tmp5880
_ = let__5019

tmp5886 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5019)


var ifres5881 Obj

if True == tmp5886 {
tmp5882 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5881 = tmp5882


} else {
tmp5883 := Call(__e, PrimFunc(symshen_4_5_1out), let__5019)


let__5020 := tmp5883
_ = let__5020

tmp5884 := Call(__e, PrimFunc(symshen_4in_1_6), let__5019)


let__5021 := tmp5884
_ = let__5021

tmp5885 := Call(__e, PrimFunc(symshen_4comb), let__5021, let__5020)


ifres5881 = tmp5885


}

ifres5877 = ifres5881


}

let__5016 := ifres5877
_ = let__5016

tmp5897 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5016)


if True == tmp5897 {
tmp5888 := Call(__e, PrimFunc(sym_5e_6), V2323)


let__5023 := tmp5888
_ = let__5023

tmp5893 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5023)


var ifres5889 Obj

if True == tmp5893 {
tmp5890 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5889 = tmp5890


} else {
tmp5891 := Call(__e, PrimFunc(symshen_4in_1_6), let__5023)


let__5024 := tmp5891
_ = let__5024

tmp5892 := Call(__e, PrimFunc(symshen_4comb), let__5024, Nil)


ifres5889 = tmp5892


}

let__5022 := ifres5889
_ = let__5022

tmp5895 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5022)


if True == tmp5895 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5022)
return
}


} else {
__e.Return(let__5016)
return
}


} else {
__e.Return(let__5009)
return
}


} else {
__e.Return(let__5003)
return
}


} else {
__e.Return(let__4997)
return
}


} else {
__e.Return(let__4991)
return
}


} else {
__e.Return(let__4983)
return
}


} else {
__e.Return(let__4977)
return
}


} else {
__e.Return(let__4971)
return
}


} else {
__e.Return(let__4965)
return
}


} else {
__e.Return(let__4959)
return
}


} else {
__e.Return(let__4948)
return
}


} else {
__e.Return(let__4937)
return
}


}, 1)

tmp5920 := Call(__e, ns2_1set, symshen_4_5s_1exprs_6, tmp5700)


_ = tmp5920

tmp5921 := MakeNative(func(__e *ControlFlow) {
V2412 := __e.Get(1)
_ = V2412
V2413 := __e.Get(2)
_ = V2413
tmp5939 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2412)
}
__typedArg0 := V2412
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres5926 Obj

if True == tmp5939 {
tmp5937 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2412)
}
__typedArg0 := V2412
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5938 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_3, tmp5937)
}
__typedArg0 := sym_3
__typedArg1 := tmp5937
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres5928 Obj

if True == tmp5938 {
tmp5935 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2412)
}
__typedArg0 := V2412
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5936 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp5935)
}
__typedArg0 := tmp5935
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres5930 Obj

if True == tmp5936 {
tmp5932 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2412)
}
__typedArg0 := V2412
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5933 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp5932)
}
__typedArg0 := tmp5932
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5934 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp5933)
}
__typedArg0 := Nil
__typedArg1 := tmp5933
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres5931 Obj

if True == tmp5934 {
ifres5931 = True


} else {
ifres5931 = False


}

ifres5930 = ifres5931


} else {
ifres5930 = False


}

var ifres5929 Obj

if True == ifres5930 {
ifres5929 = True


} else {
ifres5929 = False


}

ifres5928 = ifres5929


} else {
ifres5928 = False


}

var ifres5927 Obj

if True == ifres5928 {
ifres5927 = True


} else {
ifres5927 = False


}

ifres5926 = ifres5927


} else {
ifres5926 = False


}

if True == ifres5926 {
tmp5922 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2412)
}
__typedArg0 := V2412
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5923 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp5922)
}
__typedArg0 := tmp5922
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5924 := Call(__e, PrimFunc(symexplode), tmp5923)


__e.TailApply(PrimFunc(symappend), tmp5924, V2413)
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

tmp5940 := Call(__e, ns2_1set, symshen_4add_1sexpr, tmp5921)


_ = tmp5940

tmp5941 := MakeNative(func(__e *ControlFlow) {
V2414 := __e.Get(1)
_ = V2414
tmp5946 := Call(__e, PrimFunc(symshen_4hds_a_2), V2414, MakeInteger(91))


var ifres5942 Obj

if True == tmp5946 {
tmp5943 := Call(__e, PrimFunc(symtail), V2414)


let__5026 := tmp5943
_ = let__5026

tmp5944 := Call(__e, PrimFunc(symshen_4comb), let__5026, symshen_4skip)


ifres5942 = tmp5944


} else {
tmp5945 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5942 = tmp5945


}

let__5025 := ifres5942
_ = let__5025

tmp5948 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5025)


if True == tmp5948 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5025)
return
}


}, 1)

tmp5949 := Call(__e, ns2_1set, symshen_4_5lsb_6, tmp5941)


_ = tmp5949

tmp5950 := MakeNative(func(__e *ControlFlow) {
V2417 := __e.Get(1)
_ = V2417
tmp5955 := Call(__e, PrimFunc(symshen_4hds_a_2), V2417, MakeInteger(93))


var ifres5951 Obj

if True == tmp5955 {
tmp5952 := Call(__e, PrimFunc(symtail), V2417)


let__5028 := tmp5952
_ = let__5028

tmp5953 := Call(__e, PrimFunc(symshen_4comb), let__5028, symshen_4skip)


ifres5951 = tmp5953


} else {
tmp5954 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5951 = tmp5954


}

let__5027 := ifres5951
_ = let__5027

tmp5957 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5027)


if True == tmp5957 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5027)
return
}


}, 1)

tmp5958 := Call(__e, ns2_1set, symshen_4_5rsb_6, tmp5950)


_ = tmp5958

tmp5959 := MakeNative(func(__e *ControlFlow) {
V2420 := __e.Get(1)
_ = V2420
tmp5960 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), V2420)


let__5030 := tmp5960
_ = let__5030

tmp5966 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5030)


var ifres5961 Obj

if True == tmp5966 {
tmp5962 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5961 = tmp5962


} else {
tmp5963 := Call(__e, PrimFunc(symshen_4_5_1out), let__5030)


let__5031 := tmp5963
_ = let__5031

tmp5964 := Call(__e, PrimFunc(symshen_4in_1_6), let__5030)


let__5032 := tmp5964
_ = let__5032

tmp5965 := Call(__e, PrimFunc(symshen_4comb), let__5032, let__5031)


ifres5961 = tmp5965


}

let__5029 := ifres5961
_ = let__5029

tmp5968 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5029)


if True == tmp5968 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5029)
return
}


}, 1)

tmp5969 := Call(__e, ns2_1set, symshen_4_5s_1exprs1_6, tmp5959)


_ = tmp5969

tmp5970 := MakeNative(func(__e *ControlFlow) {
V2425 := __e.Get(1)
_ = V2425
tmp5971 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), V2425)


let__5034 := tmp5971
_ = let__5034

tmp5977 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5034)


var ifres5972 Obj

if True == tmp5977 {
tmp5973 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5972 = tmp5973


} else {
tmp5974 := Call(__e, PrimFunc(symshen_4_5_1out), let__5034)


let__5035 := tmp5974
_ = let__5035

tmp5975 := Call(__e, PrimFunc(symshen_4in_1_6), let__5034)


let__5036 := tmp5975
_ = let__5036

tmp5976 := Call(__e, PrimFunc(symshen_4comb), let__5036, let__5035)


ifres5972 = tmp5976


}

let__5033 := ifres5972
_ = let__5033

tmp5979 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5033)


if True == tmp5979 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5033)
return
}


}, 1)

tmp5980 := Call(__e, ns2_1set, symshen_4_5s_1exprs2_6, tmp5970)


_ = tmp5980

tmp5981 := MakeNative(func(__e *ControlFlow) {
V2431 := __e.Get(1)
_ = V2431
tmp6038 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V2431)
}
__typedArg0 := Nil
__typedArg1 := V2431
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp6038 {
__e.Return(Nil)
return
} else {
tmp6036 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2431)
}
__typedArg0 := V2431
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres6016 Obj

if True == tmp6036 {
tmp6034 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2431)
}
__typedArg0 := V2431
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6035 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp6034)
}
__typedArg0 := tmp6034
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres6018 Obj

if True == tmp6035 {
tmp6031 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2431)
}
__typedArg0 := V2431
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6032 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp6031)
}
__typedArg0 := tmp6031
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6033 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp6032)
}
__typedArg0 := tmp6032
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres6020 Obj

if True == tmp6033 {
tmp6027 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2431)
}
__typedArg0 := V2431
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6028 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp6027)
}
__typedArg0 := tmp6027
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6029 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp6028)
}
__typedArg0 := tmp6028
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6030 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp6029)
}
__typedArg0 := Nil
__typedArg1 := tmp6029
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres6022 Obj

if True == tmp6030 {
tmp6024 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2431)
}
__typedArg0 := V2431
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6025 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp6024)
}
__typedArg0 := tmp6024
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp6026 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp6025, symbar_b)
}
__typedArg0 := tmp6025
__typedArg1 := symbar_b
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres6023 Obj

if True == tmp6026 {
ifres6023 = True


} else {
ifres6023 = False


}

ifres6022 = ifres6023


} else {
ifres6022 = False


}

var ifres6021 Obj

if True == ifres6022 {
ifres6021 = True


} else {
ifres6021 = False


}

ifres6020 = ifres6021


} else {
ifres6020 = False


}

var ifres6019 Obj

if True == ifres6020 {
ifres6019 = True


} else {
ifres6019 = False


}

ifres6018 = ifres6019


} else {
ifres6018 = False


}

var ifres6017 Obj

if True == ifres6018 {
ifres6017 = True


} else {
ifres6017 = False


}

ifres6016 = ifres6017


} else {
ifres6016 = False


}

if True == ifres6016 {
tmp5982 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2431)
}
__typedArg0 := V2431
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5983 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2431)
}
__typedArg0 := V2431
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5984 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp5983)
}
__typedArg0 := tmp5983
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5985 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp5982, tmp5984)
}
__typedArg0 := tmp5982
__typedArg1 := tmp5984
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp5985)
}
__typedArg0 := symcons
__typedArg1 := tmp5985
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp6014 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2431)
}
__typedArg0 := V2431
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres5994 Obj

if True == tmp6014 {
tmp6012 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2431)
}
__typedArg0 := V2431
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6013 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp6012)
}
__typedArg0 := tmp6012
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres5996 Obj

if True == tmp6013 {
tmp6009 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2431)
}
__typedArg0 := V2431
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6010 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp6009)
}
__typedArg0 := tmp6009
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6011 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp6010)
}
__typedArg0 := tmp6010
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres5998 Obj

if True == tmp6011 {
tmp6005 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2431)
}
__typedArg0 := V2431
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6006 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp6005)
}
__typedArg0 := tmp6005
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6007 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp6006)
}
__typedArg0 := tmp6006
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6008 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp6007)
}
__typedArg0 := tmp6007
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres6000 Obj

if True == tmp6008 {
tmp6002 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2431)
}
__typedArg0 := V2431
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6003 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp6002)
}
__typedArg0 := tmp6002
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp6004 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp6003, symbar_b)
}
__typedArg0 := tmp6003
__typedArg1 := symbar_b
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres6001 Obj

if True == tmp6004 {
ifres6001 = True


} else {
ifres6001 = False


}

ifres6000 = ifres6001


} else {
ifres6000 = False


}

var ifres5999 Obj

if True == ifres6000 {
ifres5999 = True


} else {
ifres5999 = False


}

ifres5998 = ifres5999


} else {
ifres5998 = False


}

var ifres5997 Obj

if True == ifres5998 {
ifres5997 = True


} else {
ifres5997 = False


}

ifres5996 = ifres5997


} else {
ifres5996 = False


}

var ifres5995 Obj

if True == ifres5996 {
ifres5995 = True


} else {
ifres5995 = False


}

ifres5994 = ifres5995


} else {
ifres5994 = False


}

if True == ifres5994 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("misapplication of |\n"))
}
__typedArg0 := MakeString("misapplication of |\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
} else {
tmp5992 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2431)
}
__typedArg0 := V2431
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp5992 {
tmp5986 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2431)
}
__typedArg0 := V2431
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5987 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2431)
}
__typedArg0 := V2431
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5988 := Call(__e, PrimFunc(symshen_4cons_1form), tmp5987)


tmp5989 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp5988, Nil)
}
__typedArg0 := tmp5988
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp5990 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp5986, tmp5989)
}
__typedArg0 := tmp5986
__typedArg1 := tmp5989
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp5990)
}
__typedArg0 := symcons
__typedArg1 := tmp5990
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

tmp6039 := Call(__e, ns2_1set, symshen_4cons_1form, tmp5981)


_ = tmp6039

tmp6040 := MakeNative(func(__e *ControlFlow) {
V2432 := __e.Get(1)
_ = V2432
tmp6045 := Call(__e, PrimFunc(symshen_4hds_a_2), V2432, MakeInteger(40))


var ifres6041 Obj

if True == tmp6045 {
tmp6042 := Call(__e, PrimFunc(symtail), V2432)


let__5038 := tmp6042
_ = let__5038

tmp6043 := Call(__e, PrimFunc(symshen_4comb), let__5038, symshen_4skip)


ifres6041 = tmp6043


} else {
tmp6044 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6041 = tmp6044


}

let__5037 := ifres6041
_ = let__5037

tmp6047 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5037)


if True == tmp6047 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5037)
return
}


}, 1)

tmp6048 := Call(__e, ns2_1set, symshen_4_5lrb_6, tmp6040)


_ = tmp6048

tmp6049 := MakeNative(func(__e *ControlFlow) {
V2435 := __e.Get(1)
_ = V2435
tmp6054 := Call(__e, PrimFunc(symshen_4hds_a_2), V2435, MakeInteger(41))


var ifres6050 Obj

if True == tmp6054 {
tmp6051 := Call(__e, PrimFunc(symtail), V2435)


let__5040 := tmp6051
_ = let__5040

tmp6052 := Call(__e, PrimFunc(symshen_4comb), let__5040, symshen_4skip)


ifres6050 = tmp6052


} else {
tmp6053 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6050 = tmp6053


}

let__5039 := ifres6050
_ = let__5039

tmp6056 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5039)


if True == tmp6056 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5039)
return
}


}, 1)

tmp6057 := Call(__e, ns2_1set, symshen_4_5rrb_6, tmp6049)


_ = tmp6057

tmp6058 := MakeNative(func(__e *ControlFlow) {
V2438 := __e.Get(1)
_ = V2438
tmp6063 := Call(__e, PrimFunc(symshen_4hds_a_2), V2438, MakeInteger(123))


var ifres6059 Obj

if True == tmp6063 {
tmp6060 := Call(__e, PrimFunc(symtail), V2438)


let__5042 := tmp6060
_ = let__5042

tmp6061 := Call(__e, PrimFunc(symshen_4comb), let__5042, symshen_4skip)


ifres6059 = tmp6061


} else {
tmp6062 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6059 = tmp6062


}

let__5041 := ifres6059
_ = let__5041

tmp6065 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5041)


if True == tmp6065 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5041)
return
}


}, 1)

tmp6066 := Call(__e, ns2_1set, symshen_4_5lcurly_6, tmp6058)


_ = tmp6066

tmp6067 := MakeNative(func(__e *ControlFlow) {
V2441 := __e.Get(1)
_ = V2441
tmp6072 := Call(__e, PrimFunc(symshen_4hds_a_2), V2441, MakeInteger(125))


var ifres6068 Obj

if True == tmp6072 {
tmp6069 := Call(__e, PrimFunc(symtail), V2441)


let__5044 := tmp6069
_ = let__5044

tmp6070 := Call(__e, PrimFunc(symshen_4comb), let__5044, symshen_4skip)


ifres6068 = tmp6070


} else {
tmp6071 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6068 = tmp6071


}

let__5043 := ifres6068
_ = let__5043

tmp6074 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5043)


if True == tmp6074 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5043)
return
}


}, 1)

tmp6075 := Call(__e, ns2_1set, symshen_4_5rcurly_6, tmp6067)


_ = tmp6075

tmp6076 := MakeNative(func(__e *ControlFlow) {
V2444 := __e.Get(1)
_ = V2444
tmp6081 := Call(__e, PrimFunc(symshen_4hds_a_2), V2444, MakeInteger(124))


var ifres6077 Obj

if True == tmp6081 {
tmp6078 := Call(__e, PrimFunc(symtail), V2444)


let__5046 := tmp6078
_ = let__5046

tmp6079 := Call(__e, PrimFunc(symshen_4comb), let__5046, symshen_4skip)


ifres6077 = tmp6079


} else {
tmp6080 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6077 = tmp6080


}

let__5045 := ifres6077
_ = let__5045

tmp6083 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5045)


if True == tmp6083 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5045)
return
}


}, 1)

tmp6084 := Call(__e, ns2_1set, symshen_4_5bar_6, tmp6076)


_ = tmp6084

tmp6085 := MakeNative(func(__e *ControlFlow) {
V2447 := __e.Get(1)
_ = V2447
tmp6090 := Call(__e, PrimFunc(symshen_4hds_a_2), V2447, MakeInteger(59))


var ifres6086 Obj

if True == tmp6090 {
tmp6087 := Call(__e, PrimFunc(symtail), V2447)


let__5048 := tmp6087
_ = let__5048

tmp6088 := Call(__e, PrimFunc(symshen_4comb), let__5048, symshen_4skip)


ifres6086 = tmp6088


} else {
tmp6089 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6086 = tmp6089


}

let__5047 := ifres6086
_ = let__5047

tmp6092 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5047)


if True == tmp6092 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5047)
return
}


}, 1)

tmp6093 := Call(__e, ns2_1set, symshen_4_5semicolon_6, tmp6085)


_ = tmp6093

tmp6094 := MakeNative(func(__e *ControlFlow) {
V2450 := __e.Get(1)
_ = V2450
tmp6099 := Call(__e, PrimFunc(symshen_4hds_a_2), V2450, MakeInteger(58))


var ifres6095 Obj

if True == tmp6099 {
tmp6096 := Call(__e, PrimFunc(symtail), V2450)


let__5050 := tmp6096
_ = let__5050

tmp6097 := Call(__e, PrimFunc(symshen_4comb), let__5050, symshen_4skip)


ifres6095 = tmp6097


} else {
tmp6098 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6095 = tmp6098


}

let__5049 := ifres6095
_ = let__5049

tmp6101 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5049)


if True == tmp6101 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5049)
return
}


}, 1)

tmp6102 := Call(__e, ns2_1set, symshen_4_5colon_6, tmp6094)


_ = tmp6102

tmp6103 := MakeNative(func(__e *ControlFlow) {
V2453 := __e.Get(1)
_ = V2453
tmp6108 := Call(__e, PrimFunc(symshen_4hds_a_2), V2453, MakeInteger(44))


var ifres6104 Obj

if True == tmp6108 {
tmp6105 := Call(__e, PrimFunc(symtail), V2453)


let__5052 := tmp6105
_ = let__5052

tmp6106 := Call(__e, PrimFunc(symshen_4comb), let__5052, symshen_4skip)


ifres6104 = tmp6106


} else {
tmp6107 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6104 = tmp6107


}

let__5051 := ifres6104
_ = let__5051

tmp6110 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5051)


if True == tmp6110 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5051)
return
}


}, 1)

tmp6111 := Call(__e, ns2_1set, symshen_4_5comma_6, tmp6103)


_ = tmp6111

tmp6112 := MakeNative(func(__e *ControlFlow) {
V2456 := __e.Get(1)
_ = V2456
tmp6117 := Call(__e, PrimFunc(symshen_4hds_a_2), V2456, MakeInteger(61))


var ifres6113 Obj

if True == tmp6117 {
tmp6114 := Call(__e, PrimFunc(symtail), V2456)


let__5054 := tmp6114
_ = let__5054

tmp6115 := Call(__e, PrimFunc(symshen_4comb), let__5054, symshen_4skip)


ifres6113 = tmp6115


} else {
tmp6116 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6113 = tmp6116


}

let__5053 := ifres6113
_ = let__5053

tmp6119 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5053)


if True == tmp6119 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5053)
return
}


}, 1)

tmp6120 := Call(__e, ns2_1set, symshen_4_5equal_6, tmp6112)


_ = tmp6120

tmp6121 := MakeNative(func(__e *ControlFlow) {
V2459 := __e.Get(1)
_ = V2459
tmp6122 := Call(__e, PrimFunc(symshen_4_5singleline_6), V2459)


let__5056 := tmp6122
_ = let__5056

tmp6127 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5056)


var ifres6123 Obj

if True == tmp6127 {
tmp6124 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6123 = tmp6124


} else {
tmp6125 := Call(__e, PrimFunc(symshen_4in_1_6), let__5056)


let__5057 := tmp6125
_ = let__5057

tmp6126 := Call(__e, PrimFunc(symshen_4comb), let__5057, symshen_4skip)


ifres6123 = tmp6126


}

let__5055 := ifres6123
_ = let__5055

tmp6137 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5055)


if True == tmp6137 {
tmp6128 := Call(__e, PrimFunc(symshen_4_5multiline_6), V2459)


let__5059 := tmp6128
_ = let__5059

tmp6133 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5059)


var ifres6129 Obj

if True == tmp6133 {
tmp6130 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6129 = tmp6130


} else {
tmp6131 := Call(__e, PrimFunc(symshen_4in_1_6), let__5059)


let__5060 := tmp6131
_ = let__5060

tmp6132 := Call(__e, PrimFunc(symshen_4comb), let__5060, symshen_4skip)


ifres6129 = tmp6132


}

let__5058 := ifres6129
_ = let__5058

tmp6135 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5058)


if True == tmp6135 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5058)
return
}


} else {
__e.Return(let__5055)
return
}


}, 1)

tmp6138 := Call(__e, ns2_1set, symshen_4_5comment_6, tmp6121)


_ = tmp6138

tmp6139 := MakeNative(func(__e *ControlFlow) {
V2466 := __e.Get(1)
_ = V2466
tmp6140 := Call(__e, PrimFunc(symshen_4_5backslash_6), V2466)


let__5062 := tmp6140
_ = let__5062

tmp6160 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5062)


var ifres6141 Obj

if True == tmp6160 {
tmp6142 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6141 = tmp6142


} else {
tmp6143 := Call(__e, PrimFunc(symshen_4in_1_6), let__5062)


let__5063 := tmp6143
_ = let__5063

tmp6144 := Call(__e, PrimFunc(symshen_4_5backslash_6), let__5063)


let__5064 := tmp6144
_ = let__5064

tmp6159 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5064)


var ifres6145 Obj

if True == tmp6159 {
tmp6146 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6145 = tmp6146


} else {
tmp6147 := Call(__e, PrimFunc(symshen_4in_1_6), let__5064)


let__5065 := tmp6147
_ = let__5065

tmp6148 := Call(__e, PrimFunc(symshen_4_5shortnatters_6), let__5065)


let__5066 := tmp6148
_ = let__5066

tmp6158 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5066)


var ifres6149 Obj

if True == tmp6158 {
tmp6150 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6149 = tmp6150


} else {
tmp6151 := Call(__e, PrimFunc(symshen_4in_1_6), let__5066)


let__5067 := tmp6151
_ = let__5067

tmp6152 := Call(__e, PrimFunc(symshen_4_5returns_6), let__5067)


let__5068 := tmp6152
_ = let__5068

tmp6157 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5068)


var ifres6153 Obj

if True == tmp6157 {
tmp6154 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6153 = tmp6154


} else {
tmp6155 := Call(__e, PrimFunc(symshen_4in_1_6), let__5068)


let__5069 := tmp6155
_ = let__5069

tmp6156 := Call(__e, PrimFunc(symshen_4comb), let__5069, symshen_4skip)


ifres6153 = tmp6156


}

ifres6149 = ifres6153


}

ifres6145 = ifres6149


}

ifres6141 = ifres6145


}

let__5061 := ifres6141
_ = let__5061

tmp6162 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5061)


if True == tmp6162 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5061)
return
}


}, 1)

tmp6163 := Call(__e, ns2_1set, symshen_4_5singleline_6, tmp6139)


_ = tmp6163

tmp6164 := MakeNative(func(__e *ControlFlow) {
V2476 := __e.Get(1)
_ = V2476
tmp6169 := Call(__e, PrimFunc(symshen_4hds_a_2), V2476, MakeInteger(92))


var ifres6165 Obj

if True == tmp6169 {
tmp6166 := Call(__e, PrimFunc(symtail), V2476)


let__5071 := tmp6166
_ = let__5071

tmp6167 := Call(__e, PrimFunc(symshen_4comb), let__5071, symshen_4skip)


ifres6165 = tmp6167


} else {
tmp6168 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6165 = tmp6168


}

let__5070 := ifres6165
_ = let__5070

tmp6171 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5070)


if True == tmp6171 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5070)
return
}


}, 1)

tmp6172 := Call(__e, ns2_1set, symshen_4_5backslash_6, tmp6164)


_ = tmp6172

tmp6173 := MakeNative(func(__e *ControlFlow) {
V2479 := __e.Get(1)
_ = V2479
tmp6174 := Call(__e, PrimFunc(symshen_4_5shortnatter_6), V2479)


let__5073 := tmp6174
_ = let__5073

tmp6184 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5073)


var ifres6175 Obj

if True == tmp6184 {
tmp6176 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6175 = tmp6176


} else {
tmp6177 := Call(__e, PrimFunc(symshen_4in_1_6), let__5073)


let__5074 := tmp6177
_ = let__5074

tmp6178 := Call(__e, PrimFunc(symshen_4_5shortnatters_6), let__5074)


let__5075 := tmp6178
_ = let__5075

tmp6183 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5075)


var ifres6179 Obj

if True == tmp6183 {
tmp6180 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6179 = tmp6180


} else {
tmp6181 := Call(__e, PrimFunc(symshen_4in_1_6), let__5075)


let__5076 := tmp6181
_ = let__5076

tmp6182 := Call(__e, PrimFunc(symshen_4comb), let__5076, symshen_4skip)


ifres6179 = tmp6182


}

ifres6175 = ifres6179


}

let__5072 := ifres6175
_ = let__5072

tmp6194 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5072)


if True == tmp6194 {
tmp6185 := Call(__e, PrimFunc(sym_5e_6), V2479)


let__5078 := tmp6185
_ = let__5078

tmp6190 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5078)


var ifres6186 Obj

if True == tmp6190 {
tmp6187 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6186 = tmp6187


} else {
tmp6188 := Call(__e, PrimFunc(symshen_4in_1_6), let__5078)


let__5079 := tmp6188
_ = let__5079

tmp6189 := Call(__e, PrimFunc(symshen_4comb), let__5079, symshen_4skip)


ifres6186 = tmp6189


}

let__5077 := ifres6186
_ = let__5077

tmp6192 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5077)


if True == tmp6192 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5077)
return
}


} else {
__e.Return(let__5072)
return
}


}, 1)

tmp6195 := Call(__e, ns2_1set, symshen_4_5shortnatters_6, tmp6173)


_ = tmp6195

tmp6196 := MakeNative(func(__e *ControlFlow) {
V2488 := __e.Get(1)
_ = V2488
tmp6206 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2488)
}
__typedArg0 := V2488
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres6197 Obj

if True == tmp6206 {
tmp6198 := Call(__e, PrimFunc(symhead), V2488)


let__5081 := tmp6198
_ = let__5081

tmp6199 := Call(__e, PrimFunc(symtail), V2488)


let__5082 := tmp6199
_ = let__5082

tmp6203 := Call(__e, PrimFunc(symshen_4return_2), let__5081)


tmp6204 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp6203)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp6203
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres6200 Obj

if True == tmp6204 {
tmp6201 := Call(__e, PrimFunc(symshen_4comb), let__5082, symshen_4skip)


ifres6200 = tmp6201


} else {
tmp6202 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6200 = tmp6202


}

ifres6197 = ifres6200


} else {
tmp6205 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6197 = tmp6205


}

let__5080 := ifres6197
_ = let__5080

tmp6208 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5080)


if True == tmp6208 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5080)
return
}


}, 1)

tmp6209 := Call(__e, ns2_1set, symshen_4_5shortnatter_6, tmp6196)


_ = tmp6209

tmp6210 := MakeNative(func(__e *ControlFlow) {
V2492 := __e.Get(1)
_ = V2492
tmp6211 := Call(__e, PrimFunc(symshen_4_5return_6), V2492)


let__5084 := tmp6211
_ = let__5084

tmp6221 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5084)


var ifres6212 Obj

if True == tmp6221 {
tmp6213 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6212 = tmp6213


} else {
tmp6214 := Call(__e, PrimFunc(symshen_4in_1_6), let__5084)


let__5085 := tmp6214
_ = let__5085

tmp6215 := Call(__e, PrimFunc(symshen_4_5returns_6), let__5085)


let__5086 := tmp6215
_ = let__5086

tmp6220 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5086)


var ifres6216 Obj

if True == tmp6220 {
tmp6217 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6216 = tmp6217


} else {
tmp6218 := Call(__e, PrimFunc(symshen_4in_1_6), let__5086)


let__5087 := tmp6218
_ = let__5087

tmp6219 := Call(__e, PrimFunc(symshen_4comb), let__5087, symshen_4skip)


ifres6216 = tmp6219


}

ifres6212 = ifres6216


}

let__5083 := ifres6212
_ = let__5083

tmp6231 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5083)


if True == tmp6231 {
tmp6222 := Call(__e, PrimFunc(symshen_4_5return_6), V2492)


let__5089 := tmp6222
_ = let__5089

tmp6227 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5089)


var ifres6223 Obj

if True == tmp6227 {
tmp6224 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6223 = tmp6224


} else {
tmp6225 := Call(__e, PrimFunc(symshen_4in_1_6), let__5089)


let__5090 := tmp6225
_ = let__5090

tmp6226 := Call(__e, PrimFunc(symshen_4comb), let__5090, symshen_4skip)


ifres6223 = tmp6226


}

let__5088 := ifres6223
_ = let__5088

tmp6229 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5088)


if True == tmp6229 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5088)
return
}


} else {
__e.Return(let__5083)
return
}


}, 1)

tmp6232 := Call(__e, ns2_1set, symshen_4_5returns_6, tmp6210)


_ = tmp6232

tmp6233 := MakeNative(func(__e *ControlFlow) {
V2501 := __e.Get(1)
_ = V2501
tmp6242 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2501)
}
__typedArg0 := V2501
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres6234 Obj

if True == tmp6242 {
tmp6235 := Call(__e, PrimFunc(symhead), V2501)


let__5092 := tmp6235
_ = let__5092

tmp6236 := Call(__e, PrimFunc(symtail), V2501)


let__5093 := tmp6236
_ = let__5093

tmp6240 := Call(__e, PrimFunc(symshen_4return_2), let__5092)


var ifres6237 Obj

if True == tmp6240 {
tmp6238 := Call(__e, PrimFunc(symshen_4comb), let__5093, symshen_4skip)


ifres6237 = tmp6238


} else {
tmp6239 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6237 = tmp6239


}

ifres6234 = ifres6237


} else {
tmp6241 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6234 = tmp6241


}

let__5091 := ifres6234
_ = let__5091

tmp6244 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5091)


if True == tmp6244 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5091)
return
}


}, 1)

tmp6245 := Call(__e, ns2_1set, symshen_4_5return_6, tmp6233)


_ = tmp6245

tmp6246 := MakeNative(func(__e *ControlFlow) {
V2505 := __e.Get(1)
_ = V2505
tmp6247 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(13), Nil)
}
__typedArg0 := MakeInteger(13)
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6248 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(10), tmp6247)
}
__typedArg0 := MakeInteger(10)
__typedArg1 := tmp6247
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6249 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(9), tmp6248)
}
__typedArg0 := MakeInteger(9)
__typedArg1 := tmp6248
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symelement_2), V2505, tmp6249)
return


}, 1)

tmp6250 := Call(__e, ns2_1set, symshen_4return_2, tmp6246)


_ = tmp6250

tmp6251 := MakeNative(func(__e *ControlFlow) {
V2506 := __e.Get(1)
_ = V2506
tmp6252 := Call(__e, PrimFunc(symshen_4_5backslash_6), V2506)


let__5095 := tmp6252
_ = let__5095

tmp6267 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5095)


var ifres6253 Obj

if True == tmp6267 {
tmp6254 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6253 = tmp6254


} else {
tmp6255 := Call(__e, PrimFunc(symshen_4in_1_6), let__5095)


let__5096 := tmp6255
_ = let__5096

tmp6256 := Call(__e, PrimFunc(symshen_4_5times_6), let__5096)


let__5097 := tmp6256
_ = let__5097

tmp6266 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5097)


var ifres6257 Obj

if True == tmp6266 {
tmp6258 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6257 = tmp6258


} else {
tmp6259 := Call(__e, PrimFunc(symshen_4in_1_6), let__5097)


let__5098 := tmp6259
_ = let__5098

tmp6260 := Call(__e, PrimFunc(symshen_4_5longnatter_6), let__5098)


let__5099 := tmp6260
_ = let__5099

tmp6265 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5099)


var ifres6261 Obj

if True == tmp6265 {
tmp6262 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6261 = tmp6262


} else {
tmp6263 := Call(__e, PrimFunc(symshen_4in_1_6), let__5099)


let__5100 := tmp6263
_ = let__5100

tmp6264 := Call(__e, PrimFunc(symshen_4comb), let__5100, symshen_4skip)


ifres6261 = tmp6264


}

ifres6257 = ifres6261


}

ifres6253 = ifres6257


}

let__5094 := ifres6253
_ = let__5094

tmp6269 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5094)


if True == tmp6269 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5094)
return
}


}, 1)

tmp6270 := Call(__e, ns2_1set, symshen_4_5multiline_6, tmp6251)


_ = tmp6270

tmp6271 := MakeNative(func(__e *ControlFlow) {
V2514 := __e.Get(1)
_ = V2514
tmp6276 := Call(__e, PrimFunc(symshen_4hds_a_2), V2514, MakeInteger(42))


var ifres6272 Obj

if True == tmp6276 {
tmp6273 := Call(__e, PrimFunc(symtail), V2514)


let__5102 := tmp6273
_ = let__5102

tmp6274 := Call(__e, PrimFunc(symshen_4comb), let__5102, symshen_4skip)


ifres6272 = tmp6274


} else {
tmp6275 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6272 = tmp6275


}

let__5101 := ifres6272
_ = let__5101

tmp6278 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5101)


if True == tmp6278 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5101)
return
}


}, 1)

tmp6279 := Call(__e, ns2_1set, symshen_4_5times_6, tmp6271)


_ = tmp6279

tmp6280 := MakeNative(func(__e *ControlFlow) {
V2517 := __e.Get(1)
_ = V2517
tmp6281 := Call(__e, PrimFunc(symshen_4_5comment_6), V2517)


let__5104 := tmp6281
_ = let__5104

tmp6291 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5104)


var ifres6282 Obj

if True == tmp6291 {
tmp6283 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6282 = tmp6283


} else {
tmp6284 := Call(__e, PrimFunc(symshen_4in_1_6), let__5104)


let__5105 := tmp6284
_ = let__5105

tmp6285 := Call(__e, PrimFunc(symshen_4_5longnatter_6), let__5105)


let__5106 := tmp6285
_ = let__5106

tmp6290 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5106)


var ifres6286 Obj

if True == tmp6290 {
tmp6287 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6286 = tmp6287


} else {
tmp6288 := Call(__e, PrimFunc(symshen_4in_1_6), let__5106)


let__5107 := tmp6288
_ = let__5107

tmp6289 := Call(__e, PrimFunc(symshen_4comb), let__5107, symshen_4skip)


ifres6286 = tmp6289


}

ifres6282 = ifres6286


}

let__5103 := ifres6282
_ = let__5103

tmp6318 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5103)


if True == tmp6318 {
tmp6292 := Call(__e, PrimFunc(symshen_4_5times_6), V2517)


let__5109 := tmp6292
_ = let__5109

tmp6302 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5109)


var ifres6293 Obj

if True == tmp6302 {
tmp6294 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6293 = tmp6294


} else {
tmp6295 := Call(__e, PrimFunc(symshen_4in_1_6), let__5109)


let__5110 := tmp6295
_ = let__5110

tmp6296 := Call(__e, PrimFunc(symshen_4_5backslash_6), let__5110)


let__5111 := tmp6296
_ = let__5111

tmp6301 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5111)


var ifres6297 Obj

if True == tmp6301 {
tmp6298 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6297 = tmp6298


} else {
tmp6299 := Call(__e, PrimFunc(symshen_4in_1_6), let__5111)


let__5112 := tmp6299
_ = let__5112

tmp6300 := Call(__e, PrimFunc(symshen_4comb), let__5112, symshen_4skip)


ifres6297 = tmp6300


}

ifres6293 = ifres6297


}

let__5108 := ifres6293
_ = let__5108

tmp6316 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5108)


if True == tmp6316 {
tmp6312 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2517)
}
__typedArg0 := V2517
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres6303 Obj

if True == tmp6312 {
tmp6304 := Call(__e, PrimFunc(symtail), V2517)


let__5114 := tmp6304
_ = let__5114

tmp6305 := Call(__e, PrimFunc(symshen_4_5longnatter_6), let__5114)


let__5115 := tmp6305
_ = let__5115

tmp6310 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5115)


var ifres6306 Obj

if True == tmp6310 {
tmp6307 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6306 = tmp6307


} else {
tmp6308 := Call(__e, PrimFunc(symshen_4in_1_6), let__5115)


let__5116 := tmp6308
_ = let__5116

tmp6309 := Call(__e, PrimFunc(symshen_4comb), let__5116, symshen_4skip)


ifres6306 = tmp6309


}

ifres6303 = ifres6306


} else {
tmp6311 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6303 = tmp6311


}

let__5113 := ifres6303
_ = let__5113

tmp6314 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5113)


if True == tmp6314 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5113)
return
}


} else {
__e.Return(let__5108)
return
}


} else {
__e.Return(let__5103)
return
}


}, 1)

tmp6319 := Call(__e, ns2_1set, symshen_4_5longnatter_6, tmp6280)


_ = tmp6319

tmp6320 := MakeNative(func(__e *ControlFlow) {
V2532 := __e.Get(1)
_ = V2532
tmp6321 := Call(__e, PrimFunc(symshen_4_5str_6), V2532)


let__5118 := tmp6321
_ = let__5118

tmp6327 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5118)


var ifres6322 Obj

if True == tmp6327 {
tmp6323 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6322 = tmp6323


} else {
tmp6324 := Call(__e, PrimFunc(symshen_4_5_1out), let__5118)


let__5119 := tmp6324
_ = let__5119

tmp6325 := Call(__e, PrimFunc(symshen_4in_1_6), let__5118)


let__5120 := tmp6325
_ = let__5120

tmp6326 := Call(__e, PrimFunc(symshen_4comb), let__5120, let__5119)


ifres6322 = tmp6326


}

let__5117 := ifres6322
_ = let__5117

tmp6352 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5117)


if True == tmp6352 {
tmp6328 := Call(__e, PrimFunc(symshen_4_5number_6), V2532)


let__5122 := tmp6328
_ = let__5122

tmp6334 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5122)


var ifres6329 Obj

if True == tmp6334 {
tmp6330 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6329 = tmp6330


} else {
tmp6331 := Call(__e, PrimFunc(symshen_4_5_1out), let__5122)


let__5123 := tmp6331
_ = let__5123

tmp6332 := Call(__e, PrimFunc(symshen_4in_1_6), let__5122)


let__5124 := tmp6332
_ = let__5124

tmp6333 := Call(__e, PrimFunc(symshen_4comb), let__5124, let__5123)


ifres6329 = tmp6333


}

let__5121 := ifres6329
_ = let__5121

tmp6350 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5121)


if True == tmp6350 {
tmp6335 := Call(__e, PrimFunc(symshen_4_5sym_6), V2532)


let__5126 := tmp6335
_ = let__5126

tmp6346 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5126)


var ifres6336 Obj

if True == tmp6346 {
tmp6337 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6336 = tmp6337


} else {
tmp6338 := Call(__e, PrimFunc(symshen_4_5_1out), let__5126)


let__5127 := tmp6338
_ = let__5127

tmp6339 := Call(__e, PrimFunc(symshen_4in_1_6), let__5126)


let__5128 := tmp6339
_ = let__5128

tmp6344 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__5127, MakeString("<>"))
}
__typedArg0 := let__5127
__typedArg1 := MakeString("<>")
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres6340 Obj

if True == tmp6344 {
tmp6341 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), Nil)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6342 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp6341)
}
__typedArg0 := symvector
__typedArg1 := tmp6341
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

ifres6340 = tmp6342


} else {
tmp6343 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(let__5127)
}
__typedArg0 := let__5127
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

ifres6340 = tmp6343


}

tmp6345 := Call(__e, PrimFunc(symshen_4comb), let__5128, ifres6340)


ifres6336 = tmp6345


}

let__5125 := ifres6336
_ = let__5125

tmp6348 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5125)


if True == tmp6348 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5125)
return
}


} else {
__e.Return(let__5121)
return
}


} else {
__e.Return(let__5117)
return
}


}, 1)

tmp6353 := Call(__e, ns2_1set, symshen_4_5atom_6, tmp6320)


_ = tmp6353

tmp6354 := MakeNative(func(__e *ControlFlow) {
V2545 := __e.Get(1)
_ = V2545
tmp6355 := Call(__e, PrimFunc(symshen_4_5alpha_6), V2545)


let__5130 := tmp6355
_ = let__5130

tmp6368 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5130)


var ifres6356 Obj

if True == tmp6368 {
tmp6357 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6356 = tmp6357


} else {
tmp6358 := Call(__e, PrimFunc(symshen_4_5_1out), let__5130)


let__5131 := tmp6358
_ = let__5131

tmp6359 := Call(__e, PrimFunc(symshen_4in_1_6), let__5130)


let__5132 := tmp6359
_ = let__5132

tmp6360 := Call(__e, PrimFunc(symshen_4_5alphanums_6), let__5132)


let__5133 := tmp6360
_ = let__5133

tmp6367 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5133)


var ifres6361 Obj

if True == tmp6367 {
tmp6362 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6361 = tmp6362


} else {
tmp6363 := Call(__e, PrimFunc(symshen_4_5_1out), let__5133)


let__5134 := tmp6363
_ = let__5134

tmp6364 := Call(__e, PrimFunc(symshen_4in_1_6), let__5133)


let__5135 := tmp6364
_ = let__5135

tmp6366 := Call(__e, PrimFunc(symshen_4comb), let__5135, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(let__5131)
__typedS1, __typedOK1 := TypedString(let__5134)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := let__5131
__typedArg1 := let__5134
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())


ifres6361 = tmp6366


}

ifres6356 = ifres6361


}

let__5129 := ifres6356
_ = let__5129

tmp6370 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5129)


if True == tmp6370 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5129)
return
}


}, 1)

tmp6371 := Call(__e, ns2_1set, symshen_4_5sym_6, tmp6354)


_ = tmp6371

tmp6372 := MakeNative(func(__e *ControlFlow) {
V2553 := __e.Get(1)
_ = V2553
tmp6382 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2553)
}
__typedArg0 := V2553
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres6373 Obj

if True == tmp6382 {
tmp6374 := Call(__e, PrimFunc(symhead), V2553)


let__5137 := tmp6374
_ = let__5137

tmp6375 := Call(__e, PrimFunc(symtail), V2553)


let__5138 := tmp6375
_ = let__5138

tmp6380 := Call(__e, PrimFunc(symshen_4alpha_2), let__5137)


var ifres6376 Obj

if True == tmp6380 {
tmp6377 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symn_1_6string) {
return PrimNumberToString(let__5137)
}
__typedArg0 := let__5137
return Call(__e, PrimFunc(symn_1_6string), __typedArg0)
})()

tmp6378 := Call(__e, PrimFunc(symshen_4comb), let__5138, tmp6377)


ifres6376 = tmp6378


} else {
tmp6379 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6376 = tmp6379


}

ifres6373 = ifres6376


} else {
tmp6381 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6373 = tmp6381


}

let__5136 := ifres6373
_ = let__5136

tmp6384 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5136)


if True == tmp6384 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5136)
return
}


}, 1)

tmp6385 := Call(__e, ns2_1set, symshen_4_5alpha_6, tmp6372)


_ = tmp6385

tmp6386 := MakeNative(func(__e *ControlFlow) {
V2557 := __e.Get(1)
_ = V2557
tmp6393 := Call(__e, PrimFunc(symshen_4lowercase_2), V2557)


if True == tmp6393 {
__e.Return(True)
return
} else {
tmp6391 := Call(__e, PrimFunc(symshen_4uppercase_2), V2557)


var ifres6388 Obj

if True == tmp6391 {
ifres6388 = True


} else {
tmp6390 := Call(__e, PrimFunc(symshen_4misc_2), V2557)


var ifres6389 Obj

if True == tmp6390 {
ifres6389 = True


} else {
ifres6389 = False


}

ifres6388 = ifres6389


}

if True == ifres6388 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp6394 := Call(__e, ns2_1set, symshen_4alpha_2, tmp6386)


_ = tmp6394

tmp6395 := MakeNative(func(__e *ControlFlow) {
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
__typedArg1 := MakeInteger(97)
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
__typedArg1 := MakeInteger(122)
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

tmp6400 := Call(__e, ns2_1set, symshen_4lowercase_2, tmp6395)


_ = tmp6400

tmp6401 := MakeNative(func(__e *ControlFlow) {
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
__typedArg1 := MakeInteger(65)
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
__typedArg1 := MakeInteger(90)
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

tmp6406 := Call(__e, ns2_1set, symshen_4uppercase_2, tmp6401)


_ = tmp6406

tmp6407 := MakeNative(func(__e *ControlFlow) {
V2560 := __e.Get(1)
_ = V2560
tmp6408 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(96), Nil)
}
__typedArg0 := MakeInteger(96)
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6409 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(35), tmp6408)
}
__typedArg0 := MakeInteger(35)
__typedArg1 := tmp6408
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6410 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(39), tmp6409)
}
__typedArg0 := MakeInteger(39)
__typedArg1 := tmp6409
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6411 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(37), tmp6410)
}
__typedArg0 := MakeInteger(37)
__typedArg1 := tmp6410
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6412 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(38), tmp6411)
}
__typedArg0 := MakeInteger(38)
__typedArg1 := tmp6411
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6413 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(60), tmp6412)
}
__typedArg0 := MakeInteger(60)
__typedArg1 := tmp6412
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6414 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(62), tmp6413)
}
__typedArg0 := MakeInteger(62)
__typedArg1 := tmp6413
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6415 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(46), tmp6414)
}
__typedArg0 := MakeInteger(46)
__typedArg1 := tmp6414
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6416 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(126), tmp6415)
}
__typedArg0 := MakeInteger(126)
__typedArg1 := tmp6415
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6417 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(64), tmp6416)
}
__typedArg0 := MakeInteger(64)
__typedArg1 := tmp6416
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6418 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(33), tmp6417)
}
__typedArg0 := MakeInteger(33)
__typedArg1 := tmp6417
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6419 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(36), tmp6418)
}
__typedArg0 := MakeInteger(36)
__typedArg1 := tmp6418
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6420 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(63), tmp6419)
}
__typedArg0 := MakeInteger(63)
__typedArg1 := tmp6419
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6421 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(95), tmp6420)
}
__typedArg0 := MakeInteger(95)
__typedArg1 := tmp6420
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6422 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(43), tmp6421)
}
__typedArg0 := MakeInteger(43)
__typedArg1 := tmp6421
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6423 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(47), tmp6422)
}
__typedArg0 := MakeInteger(47)
__typedArg1 := tmp6422
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6424 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(42), tmp6423)
}
__typedArg0 := MakeInteger(42)
__typedArg1 := tmp6423
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6425 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(45), tmp6424)
}
__typedArg0 := MakeInteger(45)
__typedArg1 := tmp6424
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6426 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(61), tmp6425)
}
__typedArg0 := MakeInteger(61)
__typedArg1 := tmp6425
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symelement_2), V2560, tmp6426)
return


}, 1)

tmp6427 := Call(__e, ns2_1set, symshen_4misc_2, tmp6407)


_ = tmp6427

tmp6428 := MakeNative(func(__e *ControlFlow) {
V2561 := __e.Get(1)
_ = V2561
tmp6429 := Call(__e, PrimFunc(symshen_4_5alphanum_6), V2561)


let__5140 := tmp6429
_ = let__5140

tmp6442 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5140)


var ifres6430 Obj

if True == tmp6442 {
tmp6431 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6430 = tmp6431


} else {
tmp6432 := Call(__e, PrimFunc(symshen_4_5_1out), let__5140)


let__5141 := tmp6432
_ = let__5141

tmp6433 := Call(__e, PrimFunc(symshen_4in_1_6), let__5140)


let__5142 := tmp6433
_ = let__5142

tmp6434 := Call(__e, PrimFunc(symshen_4_5alphanums_6), let__5142)


let__5143 := tmp6434
_ = let__5143

tmp6441 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5143)


var ifres6435 Obj

if True == tmp6441 {
tmp6436 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6435 = tmp6436


} else {
tmp6437 := Call(__e, PrimFunc(symshen_4_5_1out), let__5143)


let__5144 := tmp6437
_ = let__5144

tmp6438 := Call(__e, PrimFunc(symshen_4in_1_6), let__5143)


let__5145 := tmp6438
_ = let__5145

tmp6440 := Call(__e, PrimFunc(symshen_4comb), let__5145, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(let__5141)
__typedS1, __typedOK1 := TypedString(let__5144)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := let__5141
__typedArg1 := let__5144
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())


ifres6435 = tmp6440


}

ifres6430 = ifres6435


}

let__5139 := ifres6430
_ = let__5139

tmp6452 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5139)


if True == tmp6452 {
tmp6443 := Call(__e, PrimFunc(sym_5e_6), V2561)


let__5147 := tmp6443
_ = let__5147

tmp6448 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5147)


var ifres6444 Obj

if True == tmp6448 {
tmp6445 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6444 = tmp6445


} else {
tmp6446 := Call(__e, PrimFunc(symshen_4in_1_6), let__5147)


let__5148 := tmp6446
_ = let__5148

tmp6447 := Call(__e, PrimFunc(symshen_4comb), let__5148, MakeString(""))


ifres6444 = tmp6447


}

let__5146 := ifres6444
_ = let__5146

tmp6450 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5146)


if True == tmp6450 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5146)
return
}


} else {
__e.Return(let__5139)
return
}


}, 1)

tmp6453 := Call(__e, ns2_1set, symshen_4_5alphanums_6, tmp6428)


_ = tmp6453

tmp6454 := MakeNative(func(__e *ControlFlow) {
V2572 := __e.Get(1)
_ = V2572
tmp6455 := Call(__e, PrimFunc(symshen_4_5alpha_6), V2572)


let__5150 := tmp6455
_ = let__5150

tmp6461 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5150)


var ifres6456 Obj

if True == tmp6461 {
tmp6457 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6456 = tmp6457


} else {
tmp6458 := Call(__e, PrimFunc(symshen_4_5_1out), let__5150)


let__5151 := tmp6458
_ = let__5151

tmp6459 := Call(__e, PrimFunc(symshen_4in_1_6), let__5150)


let__5152 := tmp6459
_ = let__5152

tmp6460 := Call(__e, PrimFunc(symshen_4comb), let__5152, let__5151)


ifres6456 = tmp6460


}

let__5149 := ifres6456
_ = let__5149

tmp6472 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5149)


if True == tmp6472 {
tmp6462 := Call(__e, PrimFunc(symshen_4_5numeral_6), V2572)


let__5154 := tmp6462
_ = let__5154

tmp6468 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5154)


var ifres6463 Obj

if True == tmp6468 {
tmp6464 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6463 = tmp6464


} else {
tmp6465 := Call(__e, PrimFunc(symshen_4_5_1out), let__5154)


let__5155 := tmp6465
_ = let__5155

tmp6466 := Call(__e, PrimFunc(symshen_4in_1_6), let__5154)


let__5156 := tmp6466
_ = let__5156

tmp6467 := Call(__e, PrimFunc(symshen_4comb), let__5156, let__5155)


ifres6463 = tmp6467


}

let__5153 := ifres6463
_ = let__5153

tmp6470 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5153)


if True == tmp6470 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5153)
return
}


} else {
__e.Return(let__5149)
return
}


}, 1)

tmp6473 := Call(__e, ns2_1set, symshen_4_5alphanum_6, tmp6454)


_ = tmp6473

tmp6474 := MakeNative(func(__e *ControlFlow) {
V2581 := __e.Get(1)
_ = V2581
tmp6484 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2581)
}
__typedArg0 := V2581
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres6475 Obj

if True == tmp6484 {
tmp6476 := Call(__e, PrimFunc(symhead), V2581)


let__5158 := tmp6476
_ = let__5158

tmp6477 := Call(__e, PrimFunc(symtail), V2581)


let__5159 := tmp6477
_ = let__5159

tmp6482 := Call(__e, PrimFunc(symshen_4digit_2), let__5158)


var ifres6478 Obj

if True == tmp6482 {
tmp6479 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symn_1_6string) {
return PrimNumberToString(let__5158)
}
__typedArg0 := let__5158
return Call(__e, PrimFunc(symn_1_6string), __typedArg0)
})()

tmp6480 := Call(__e, PrimFunc(symshen_4comb), let__5159, tmp6479)


ifres6478 = tmp6480


} else {
tmp6481 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6478 = tmp6481


}

ifres6475 = ifres6478


} else {
tmp6483 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6475 = tmp6483


}

let__5157 := ifres6475
_ = let__5157

tmp6486 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5157)


if True == tmp6486 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5157)
return
}


}, 1)

tmp6487 := Call(__e, ns2_1set, symshen_4_5numeral_6, tmp6474)


_ = tmp6487

tmp6488 := MakeNative(func(__e *ControlFlow) {
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
__typedArg1 := MakeInteger(48)
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
__typedArg1 := MakeInteger(57)
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

tmp6493 := Call(__e, ns2_1set, symshen_4digit_2, tmp6488)


_ = tmp6493

tmp6494 := MakeNative(func(__e *ControlFlow) {
V2586 := __e.Get(1)
_ = V2586
tmp6495 := Call(__e, PrimFunc(symshen_4_5dbq_6), V2586)


let__5161 := tmp6495
_ = let__5161

tmp6511 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5161)


var ifres6496 Obj

if True == tmp6511 {
tmp6497 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6496 = tmp6497


} else {
tmp6498 := Call(__e, PrimFunc(symshen_4in_1_6), let__5161)


let__5162 := tmp6498
_ = let__5162

tmp6499 := Call(__e, PrimFunc(symshen_4_5strcontents_6), let__5162)


let__5163 := tmp6499
_ = let__5163

tmp6510 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5163)


var ifres6500 Obj

if True == tmp6510 {
tmp6501 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6500 = tmp6501


} else {
tmp6502 := Call(__e, PrimFunc(symshen_4_5_1out), let__5163)


let__5164 := tmp6502
_ = let__5164

tmp6503 := Call(__e, PrimFunc(symshen_4in_1_6), let__5163)


let__5165 := tmp6503
_ = let__5165

tmp6504 := Call(__e, PrimFunc(symshen_4_5dbq_6), let__5165)


let__5166 := tmp6504
_ = let__5166

tmp6509 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5166)


var ifres6505 Obj

if True == tmp6509 {
tmp6506 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6505 = tmp6506


} else {
tmp6507 := Call(__e, PrimFunc(symshen_4in_1_6), let__5166)


let__5167 := tmp6507
_ = let__5167

tmp6508 := Call(__e, PrimFunc(symshen_4comb), let__5167, let__5164)


ifres6505 = tmp6508


}

ifres6500 = ifres6505


}

ifres6496 = ifres6500


}

let__5160 := ifres6496
_ = let__5160

tmp6513 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5160)


if True == tmp6513 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5160)
return
}


}, 1)

tmp6514 := Call(__e, ns2_1set, symshen_4_5str_6, tmp6494)


_ = tmp6514

tmp6515 := MakeNative(func(__e *ControlFlow) {
V2595 := __e.Get(1)
_ = V2595
tmp6520 := Call(__e, PrimFunc(symshen_4hds_a_2), V2595, MakeInteger(34))


var ifres6516 Obj

if True == tmp6520 {
tmp6517 := Call(__e, PrimFunc(symtail), V2595)


let__5169 := tmp6517
_ = let__5169

tmp6518 := Call(__e, PrimFunc(symshen_4comb), let__5169, symshen_4skip)


ifres6516 = tmp6518


} else {
tmp6519 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6516 = tmp6519


}

let__5168 := ifres6516
_ = let__5168

tmp6522 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5168)


if True == tmp6522 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5168)
return
}


}, 1)

tmp6523 := Call(__e, ns2_1set, symshen_4_5dbq_6, tmp6515)


_ = tmp6523

tmp6524 := MakeNative(func(__e *ControlFlow) {
V2598 := __e.Get(1)
_ = V2598
tmp6525 := Call(__e, PrimFunc(symshen_4_5strc_6), V2598)


let__5171 := tmp6525
_ = let__5171

tmp6538 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5171)


var ifres6526 Obj

if True == tmp6538 {
tmp6527 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6526 = tmp6527


} else {
tmp6528 := Call(__e, PrimFunc(symshen_4_5_1out), let__5171)


let__5172 := tmp6528
_ = let__5172

tmp6529 := Call(__e, PrimFunc(symshen_4in_1_6), let__5171)


let__5173 := tmp6529
_ = let__5173

tmp6530 := Call(__e, PrimFunc(symshen_4_5strcontents_6), let__5173)


let__5174 := tmp6530
_ = let__5174

tmp6537 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5174)


var ifres6531 Obj

if True == tmp6537 {
tmp6532 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6531 = tmp6532


} else {
tmp6533 := Call(__e, PrimFunc(symshen_4_5_1out), let__5174)


let__5175 := tmp6533
_ = let__5175

tmp6534 := Call(__e, PrimFunc(symshen_4in_1_6), let__5174)


let__5176 := tmp6534
_ = let__5176

tmp6536 := Call(__e, PrimFunc(symshen_4comb), let__5176, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(let__5172)
__typedS1, __typedOK1 := TypedString(let__5175)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := let__5172
__typedArg1 := let__5175
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())


ifres6531 = tmp6536


}

ifres6526 = ifres6531


}

let__5170 := ifres6526
_ = let__5170

tmp6548 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5170)


if True == tmp6548 {
tmp6539 := Call(__e, PrimFunc(sym_5e_6), V2598)


let__5178 := tmp6539
_ = let__5178

tmp6544 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5178)


var ifres6540 Obj

if True == tmp6544 {
tmp6541 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6540 = tmp6541


} else {
tmp6542 := Call(__e, PrimFunc(symshen_4in_1_6), let__5178)


let__5179 := tmp6542
_ = let__5179

tmp6543 := Call(__e, PrimFunc(symshen_4comb), let__5179, MakeString(""))


ifres6540 = tmp6543


}

let__5177 := ifres6540
_ = let__5177

tmp6546 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5177)


if True == tmp6546 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5177)
return
}


} else {
__e.Return(let__5170)
return
}


}, 1)

tmp6549 := Call(__e, ns2_1set, symshen_4_5strcontents_6, tmp6524)


_ = tmp6549

tmp6550 := MakeNative(func(__e *ControlFlow) {
V2609 := __e.Get(1)
_ = V2609
tmp6551 := Call(__e, PrimFunc(symshen_4_5control_6), V2609)


let__5181 := tmp6551
_ = let__5181

tmp6557 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5181)


var ifres6552 Obj

if True == tmp6557 {
tmp6553 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6552 = tmp6553


} else {
tmp6554 := Call(__e, PrimFunc(symshen_4_5_1out), let__5181)


let__5182 := tmp6554
_ = let__5182

tmp6555 := Call(__e, PrimFunc(symshen_4in_1_6), let__5181)


let__5183 := tmp6555
_ = let__5183

tmp6556 := Call(__e, PrimFunc(symshen_4comb), let__5183, let__5182)


ifres6552 = tmp6556


}

let__5180 := ifres6552
_ = let__5180

tmp6568 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5180)


if True == tmp6568 {
tmp6558 := Call(__e, PrimFunc(symshen_4_5notdbq_6), V2609)


let__5185 := tmp6558
_ = let__5185

tmp6564 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5185)


var ifres6559 Obj

if True == tmp6564 {
tmp6560 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6559 = tmp6560


} else {
tmp6561 := Call(__e, PrimFunc(symshen_4_5_1out), let__5185)


let__5186 := tmp6561
_ = let__5186

tmp6562 := Call(__e, PrimFunc(symshen_4in_1_6), let__5185)


let__5187 := tmp6562
_ = let__5187

tmp6563 := Call(__e, PrimFunc(symshen_4comb), let__5187, let__5186)


ifres6559 = tmp6563


}

let__5184 := ifres6559
_ = let__5184

tmp6566 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5184)


if True == tmp6566 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5184)
return
}


} else {
__e.Return(let__5180)
return
}


}, 1)

tmp6569 := Call(__e, ns2_1set, symshen_4_5strc_6, tmp6550)


_ = tmp6569

tmp6570 := MakeNative(func(__e *ControlFlow) {
V2618 := __e.Get(1)
_ = V2618
tmp6571 := Call(__e, PrimFunc(symshen_4_5lowC_6), V2618)


let__5189 := tmp6571
_ = let__5189

tmp6593 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5189)


var ifres6572 Obj

if True == tmp6593 {
tmp6573 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6572 = tmp6573


} else {
tmp6574 := Call(__e, PrimFunc(symshen_4in_1_6), let__5189)


let__5190 := tmp6574
_ = let__5190

tmp6575 := Call(__e, PrimFunc(symshen_4_5hash_6), let__5190)


let__5191 := tmp6575
_ = let__5191

tmp6592 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5191)


var ifres6576 Obj

if True == tmp6592 {
tmp6577 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6576 = tmp6577


} else {
tmp6578 := Call(__e, PrimFunc(symshen_4in_1_6), let__5191)


let__5192 := tmp6578
_ = let__5192

tmp6579 := Call(__e, PrimFunc(symshen_4_5integer_6), let__5192)


let__5193 := tmp6579
_ = let__5193

tmp6591 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5193)


var ifres6580 Obj

if True == tmp6591 {
tmp6581 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6580 = tmp6581


} else {
tmp6582 := Call(__e, PrimFunc(symshen_4_5_1out), let__5193)


let__5194 := tmp6582
_ = let__5194

tmp6583 := Call(__e, PrimFunc(symshen_4in_1_6), let__5193)


let__5195 := tmp6583
_ = let__5195

tmp6584 := Call(__e, PrimFunc(symshen_4_5semicolon_6), let__5195)


let__5196 := tmp6584
_ = let__5196

tmp6590 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5196)


var ifres6585 Obj

if True == tmp6590 {
tmp6586 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6585 = tmp6586


} else {
tmp6587 := Call(__e, PrimFunc(symshen_4in_1_6), let__5196)


let__5197 := tmp6587
_ = let__5197

tmp6588 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symn_1_6string) {
return PrimNumberToString(let__5194)
}
__typedArg0 := let__5194
return Call(__e, PrimFunc(symn_1_6string), __typedArg0)
})()

tmp6589 := Call(__e, PrimFunc(symshen_4comb), let__5197, tmp6588)


ifres6585 = tmp6589


}

ifres6580 = ifres6585


}

ifres6576 = ifres6580


}

ifres6572 = ifres6576


}

let__5188 := ifres6572
_ = let__5188

tmp6595 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5188)


if True == tmp6595 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5188)
return
}


}, 1)

tmp6596 := Call(__e, ns2_1set, symshen_4_5control_6, tmp6570)


_ = tmp6596

tmp6597 := MakeNative(func(__e *ControlFlow) {
V2629 := __e.Get(1)
_ = V2629
tmp6608 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2629)
}
__typedArg0 := V2629
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres6598 Obj

if True == tmp6608 {
tmp6599 := Call(__e, PrimFunc(symhead), V2629)


let__5199 := tmp6599
_ = let__5199

tmp6600 := Call(__e, PrimFunc(symtail), V2629)


let__5200 := tmp6600
_ = let__5200

tmp6605 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__5199, MakeInteger(34))
}
__typedArg0 := let__5199
__typedArg1 := MakeInteger(34)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

tmp6606 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp6605)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp6605
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres6601 Obj

if True == tmp6606 {
tmp6602 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symn_1_6string) {
return PrimNumberToString(let__5199)
}
__typedArg0 := let__5199
return Call(__e, PrimFunc(symn_1_6string), __typedArg0)
})()

tmp6603 := Call(__e, PrimFunc(symshen_4comb), let__5200, tmp6602)


ifres6601 = tmp6603


} else {
tmp6604 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6601 = tmp6604


}

ifres6598 = ifres6601


} else {
tmp6607 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6598 = tmp6607


}

let__5198 := ifres6598
_ = let__5198

tmp6610 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5198)


if True == tmp6610 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5198)
return
}


}, 1)

tmp6611 := Call(__e, ns2_1set, symshen_4_5notdbq_6, tmp6597)


_ = tmp6611

tmp6612 := MakeNative(func(__e *ControlFlow) {
V2633 := __e.Get(1)
_ = V2633
tmp6617 := Call(__e, PrimFunc(symshen_4hds_a_2), V2633, MakeInteger(99))


var ifres6613 Obj

if True == tmp6617 {
tmp6614 := Call(__e, PrimFunc(symtail), V2633)


let__5202 := tmp6614
_ = let__5202

tmp6615 := Call(__e, PrimFunc(symshen_4comb), let__5202, symshen_4skip)


ifres6613 = tmp6615


} else {
tmp6616 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6613 = tmp6616


}

let__5201 := ifres6613
_ = let__5201

tmp6619 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5201)


if True == tmp6619 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5201)
return
}


}, 1)

tmp6620 := Call(__e, ns2_1set, symshen_4_5lowC_6, tmp6612)


_ = tmp6620

tmp6621 := MakeNative(func(__e *ControlFlow) {
V2636 := __e.Get(1)
_ = V2636
tmp6626 := Call(__e, PrimFunc(symshen_4hds_a_2), V2636, MakeInteger(35))


var ifres6622 Obj

if True == tmp6626 {
tmp6623 := Call(__e, PrimFunc(symtail), V2636)


let__5204 := tmp6623
_ = let__5204

tmp6624 := Call(__e, PrimFunc(symshen_4comb), let__5204, symshen_4skip)


ifres6622 = tmp6624


} else {
tmp6625 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6622 = tmp6625


}

let__5203 := ifres6622
_ = let__5203

tmp6628 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5203)


if True == tmp6628 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5203)
return
}


}, 1)

tmp6629 := Call(__e, ns2_1set, symshen_4_5hash_6, tmp6621)


_ = tmp6629

tmp6630 := MakeNative(func(__e *ControlFlow) {
V2639 := __e.Get(1)
_ = V2639
tmp6631 := Call(__e, PrimFunc(symshen_4_5minus_6), V2639)


let__5206 := tmp6631
_ = let__5206

tmp6643 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5206)


var ifres6632 Obj

if True == tmp6643 {
tmp6633 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6632 = tmp6633


} else {
tmp6634 := Call(__e, PrimFunc(symshen_4in_1_6), let__5206)


let__5207 := tmp6634
_ = let__5207

tmp6635 := Call(__e, PrimFunc(symshen_4_5number_6), let__5207)


let__5208 := tmp6635
_ = let__5208

tmp6642 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5208)


var ifres6636 Obj

if True == tmp6642 {
tmp6637 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6636 = tmp6637


} else {
tmp6638 := Call(__e, PrimFunc(symshen_4_5_1out), let__5208)


let__5209 := tmp6638
_ = let__5209

tmp6639 := Call(__e, PrimFunc(symshen_4in_1_6), let__5208)


let__5210 := tmp6639
_ = let__5210

tmp6641 := Call(__e, PrimFunc(symshen_4comb), let__5210, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(0))
__typedN1, __typedOK1 := TypedFloat64(let__5209)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := MakeInteger(0)
__typedArg1 := let__5209
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})())


ifres6636 = tmp6641


}

ifres6632 = ifres6636


}

let__5205 := ifres6632
_ = let__5205

tmp6686 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5205)


if True == tmp6686 {
tmp6644 := Call(__e, PrimFunc(symshen_4_5plus_6), V2639)


let__5212 := tmp6644
_ = let__5212

tmp6655 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5212)


var ifres6645 Obj

if True == tmp6655 {
tmp6646 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6645 = tmp6646


} else {
tmp6647 := Call(__e, PrimFunc(symshen_4in_1_6), let__5212)


let__5213 := tmp6647
_ = let__5213

tmp6648 := Call(__e, PrimFunc(symshen_4_5number_6), let__5213)


let__5214 := tmp6648
_ = let__5214

tmp6654 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5214)


var ifres6649 Obj

if True == tmp6654 {
tmp6650 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6649 = tmp6650


} else {
tmp6651 := Call(__e, PrimFunc(symshen_4_5_1out), let__5214)


let__5215 := tmp6651
_ = let__5215

tmp6652 := Call(__e, PrimFunc(symshen_4in_1_6), let__5214)


let__5216 := tmp6652
_ = let__5216

tmp6653 := Call(__e, PrimFunc(symshen_4comb), let__5216, let__5215)


ifres6649 = tmp6653


}

ifres6645 = ifres6649


}

let__5211 := ifres6645
_ = let__5211

tmp6684 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5211)


if True == tmp6684 {
tmp6656 := Call(__e, PrimFunc(symshen_4_5e_1number_6), V2639)


let__5218 := tmp6656
_ = let__5218

tmp6662 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5218)


var ifres6657 Obj

if True == tmp6662 {
tmp6658 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6657 = tmp6658


} else {
tmp6659 := Call(__e, PrimFunc(symshen_4_5_1out), let__5218)


let__5219 := tmp6659
_ = let__5219

tmp6660 := Call(__e, PrimFunc(symshen_4in_1_6), let__5218)


let__5220 := tmp6660
_ = let__5220

tmp6661 := Call(__e, PrimFunc(symshen_4comb), let__5220, let__5219)


ifres6657 = tmp6661


}

let__5217 := ifres6657
_ = let__5217

tmp6682 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5217)


if True == tmp6682 {
tmp6663 := Call(__e, PrimFunc(symshen_4_5float_6), V2639)


let__5222 := tmp6663
_ = let__5222

tmp6669 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5222)


var ifres6664 Obj

if True == tmp6669 {
tmp6665 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6664 = tmp6665


} else {
tmp6666 := Call(__e, PrimFunc(symshen_4_5_1out), let__5222)


let__5223 := tmp6666
_ = let__5223

tmp6667 := Call(__e, PrimFunc(symshen_4in_1_6), let__5222)


let__5224 := tmp6667
_ = let__5224

tmp6668 := Call(__e, PrimFunc(symshen_4comb), let__5224, let__5223)


ifres6664 = tmp6668


}

let__5221 := ifres6664
_ = let__5221

tmp6680 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5221)


if True == tmp6680 {
tmp6670 := Call(__e, PrimFunc(symshen_4_5integer_6), V2639)


let__5226 := tmp6670
_ = let__5226

tmp6676 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5226)


var ifres6671 Obj

if True == tmp6676 {
tmp6672 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6671 = tmp6672


} else {
tmp6673 := Call(__e, PrimFunc(symshen_4_5_1out), let__5226)


let__5227 := tmp6673
_ = let__5227

tmp6674 := Call(__e, PrimFunc(symshen_4in_1_6), let__5226)


let__5228 := tmp6674
_ = let__5228

tmp6675 := Call(__e, PrimFunc(symshen_4comb), let__5228, let__5227)


ifres6671 = tmp6675


}

let__5225 := ifres6671
_ = let__5225

tmp6678 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5225)


if True == tmp6678 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5225)
return
}


} else {
__e.Return(let__5221)
return
}


} else {
__e.Return(let__5217)
return
}


} else {
__e.Return(let__5211)
return
}


} else {
__e.Return(let__5205)
return
}


}, 1)

tmp6687 := Call(__e, ns2_1set, symshen_4_5number_6, tmp6630)


_ = tmp6687

tmp6688 := MakeNative(func(__e *ControlFlow) {
V2664 := __e.Get(1)
_ = V2664
tmp6693 := Call(__e, PrimFunc(symshen_4hds_a_2), V2664, MakeInteger(45))


var ifres6689 Obj

if True == tmp6693 {
tmp6690 := Call(__e, PrimFunc(symtail), V2664)


let__5230 := tmp6690
_ = let__5230

tmp6691 := Call(__e, PrimFunc(symshen_4comb), let__5230, symshen_4skip)


ifres6689 = tmp6691


} else {
tmp6692 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6689 = tmp6692


}

let__5229 := ifres6689
_ = let__5229

tmp6695 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5229)


if True == tmp6695 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5229)
return
}


}, 1)

tmp6696 := Call(__e, ns2_1set, symshen_4_5minus_6, tmp6688)


_ = tmp6696

tmp6697 := MakeNative(func(__e *ControlFlow) {
V2667 := __e.Get(1)
_ = V2667
tmp6702 := Call(__e, PrimFunc(symshen_4hds_a_2), V2667, MakeInteger(43))


var ifres6698 Obj

if True == tmp6702 {
tmp6699 := Call(__e, PrimFunc(symtail), V2667)


let__5232 := tmp6699
_ = let__5232

tmp6700 := Call(__e, PrimFunc(symshen_4comb), let__5232, symshen_4skip)


ifres6698 = tmp6700


} else {
tmp6701 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6698 = tmp6701


}

let__5231 := ifres6698
_ = let__5231

tmp6704 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5231)


if True == tmp6704 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5231)
return
}


}, 1)

tmp6705 := Call(__e, ns2_1set, symshen_4_5plus_6, tmp6697)


_ = tmp6705

tmp6706 := MakeNative(func(__e *ControlFlow) {
V2670 := __e.Get(1)
_ = V2670
tmp6707 := Call(__e, PrimFunc(symshen_4_5digits_6), V2670)


let__5234 := tmp6707
_ = let__5234

tmp6714 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5234)


var ifres6708 Obj

if True == tmp6714 {
tmp6709 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6708 = tmp6709


} else {
tmp6710 := Call(__e, PrimFunc(symshen_4_5_1out), let__5234)


let__5235 := tmp6710
_ = let__5235

tmp6711 := Call(__e, PrimFunc(symshen_4in_1_6), let__5234)


let__5236 := tmp6711
_ = let__5236

tmp6712 := Call(__e, PrimFunc(symshen_4compute_1integer), let__5235)


tmp6713 := Call(__e, PrimFunc(symshen_4comb), let__5236, tmp6712)


ifres6708 = tmp6713


}

let__5233 := ifres6708
_ = let__5233

tmp6716 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5233)


if True == tmp6716 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5233)
return
}


}, 1)

tmp6717 := Call(__e, ns2_1set, symshen_4_5integer_6, tmp6706)


_ = tmp6717

tmp6718 := MakeNative(func(__e *ControlFlow) {
V2675 := __e.Get(1)
_ = V2675
tmp6719 := Call(__e, PrimFunc(symshen_4_5digit_6), V2675)


let__5238 := tmp6719
_ = let__5238

tmp6732 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5238)


var ifres6720 Obj

if True == tmp6732 {
tmp6721 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6720 = tmp6721


} else {
tmp6722 := Call(__e, PrimFunc(symshen_4_5_1out), let__5238)


let__5239 := tmp6722
_ = let__5239

tmp6723 := Call(__e, PrimFunc(symshen_4in_1_6), let__5238)


let__5240 := tmp6723
_ = let__5240

tmp6724 := Call(__e, PrimFunc(symshen_4_5digits_6), let__5240)


let__5241 := tmp6724
_ = let__5241

tmp6731 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5241)


var ifres6725 Obj

if True == tmp6731 {
tmp6726 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6725 = tmp6726


} else {
tmp6727 := Call(__e, PrimFunc(symshen_4_5_1out), let__5241)


let__5242 := tmp6727
_ = let__5242

tmp6728 := Call(__e, PrimFunc(symshen_4in_1_6), let__5241)


let__5243 := tmp6728
_ = let__5243

tmp6729 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__5239, let__5242)
}
__typedArg0 := let__5239
__typedArg1 := let__5242
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6730 := Call(__e, PrimFunc(symshen_4comb), let__5243, tmp6729)


ifres6725 = tmp6730


}

ifres6720 = ifres6725


}

let__5237 := ifres6720
_ = let__5237

tmp6744 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5237)


if True == tmp6744 {
tmp6733 := Call(__e, PrimFunc(symshen_4_5digit_6), V2675)


let__5245 := tmp6733
_ = let__5245

tmp6740 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5245)


var ifres6734 Obj

if True == tmp6740 {
tmp6735 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6734 = tmp6735


} else {
tmp6736 := Call(__e, PrimFunc(symshen_4_5_1out), let__5245)


let__5246 := tmp6736
_ = let__5246

tmp6737 := Call(__e, PrimFunc(symshen_4in_1_6), let__5245)


let__5247 := tmp6737
_ = let__5247

tmp6738 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__5246, Nil)
}
__typedArg0 := let__5246
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6739 := Call(__e, PrimFunc(symshen_4comb), let__5247, tmp6738)


ifres6734 = tmp6739


}

let__5244 := ifres6734
_ = let__5244

tmp6742 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5244)


if True == tmp6742 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5244)
return
}


} else {
__e.Return(let__5237)
return
}


}, 1)

tmp6745 := Call(__e, ns2_1set, symshen_4_5digits_6, tmp6718)


_ = tmp6745

tmp6746 := MakeNative(func(__e *ControlFlow) {
V2687 := __e.Get(1)
_ = V2687
tmp6756 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2687)
}
__typedArg0 := V2687
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres6747 Obj

if True == tmp6756 {
tmp6748 := Call(__e, PrimFunc(symhead), V2687)


let__5249 := tmp6748
_ = let__5249

tmp6749 := Call(__e, PrimFunc(symtail), V2687)


let__5250 := tmp6749
_ = let__5250

tmp6754 := Call(__e, PrimFunc(symshen_4digit_2), let__5249)


var ifres6750 Obj

if True == tmp6754 {
tmp6751 := Call(__e, PrimFunc(symshen_4byte_1_6digit), let__5249)


tmp6752 := Call(__e, PrimFunc(symshen_4comb), let__5250, tmp6751)


ifres6750 = tmp6752


} else {
tmp6753 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6750 = tmp6753


}

ifres6747 = ifres6750


} else {
tmp6755 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6747 = tmp6755


}

let__5248 := ifres6747
_ = let__5248

tmp6758 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5248)


if True == tmp6758 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5248)
return
}


}, 1)

tmp6759 := Call(__e, ns2_1set, symshen_4_5digit_6, tmp6746)


_ = tmp6759

tmp6760 := MakeNative(func(__e *ControlFlow) {
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
__typedArg1 := MakeInteger(48)
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})())
return
}, 1)

tmp6761 := Call(__e, ns2_1set, symshen_4byte_1_6digit, tmp6760)


_ = tmp6761

tmp6762 := MakeNative(func(__e *ControlFlow) {
V2692 := __e.Get(1)
_ = V2692
tmp6763 := Call(__e, PrimFunc(symreverse), V2692)


__e.TailApply(PrimFunc(symshen_4compute_1integer_1h), tmp6763, MakeInteger(0))
return


}, 1)

tmp6764 := Call(__e, ns2_1set, symshen_4compute_1integer, tmp6762)


_ = tmp6764

tmp6765 := MakeNative(func(__e *ControlFlow) {
V2695 := __e.Get(1)
_ = V2695
V2696 := __e.Get(2)
_ = V2696
tmp6775 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V2695)
}
__typedArg0 := Nil
__typedArg1 := V2695
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp6775 {
__e.Return(MakeInteger(0))
return
} else {
tmp6773 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2695)
}
__typedArg0 := V2695
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp6773 {
tmp6766 := Call(__e, PrimFunc(symshen_4expt), MakeInteger(10), V2696)


tmp6767 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2695)
}
__typedArg0 := V2695
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp6768 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_d) {
__typedN0, __typedOK0 := TypedFloat64(tmp6766)
__typedN1, __typedOK1 := TypedFloat64(tmp6767)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_d) {
return TypedMaterializeNumber((__typedN0 * __typedN1))
}}
__typedArg0 := tmp6766
__typedArg1 := tmp6767
return Call(__e, PrimFunc(sym_d), __typedArg0, __typedArg1)
})()

tmp6769 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2695)
}
__typedArg0 := V2695
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6771 := Call(__e, PrimFunc(symshen_4compute_1integer_1h), tmp6769, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V2696)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V2696
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(tmp6768)
__typedN1, __typedOK1 := TypedFloat64(tmp6771)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := tmp6768
__typedArg1 := tmp6771
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

tmp6776 := Call(__e, ns2_1set, symshen_4compute_1integer_1h, tmp6765)


_ = tmp6776

tmp6777 := MakeNative(func(__e *ControlFlow) {
V2699 := __e.Get(1)
_ = V2699
V2700 := __e.Get(2)
_ = V2700
tmp6785 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(0), V2700)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := V2700
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp6785 {
__e.Return(MakeInteger(1))
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
__typedArg1 := MakeInteger(0)
return Call(__e, PrimFunc(sym_6), __typedArg0, __typedArg1)
})() {
tmp6779 := Call(__e, PrimFunc(symshen_4expt), V2699, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(V2700)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := V2700
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})())


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_d) {
__typedN0, __typedOK0 := TypedFloat64(V2699)
__typedN1, __typedOK1 := TypedFloat64(tmp6779)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_d) {
return TypedMaterializeNumber((__typedN0 * __typedN1))
}}
__typedArg0 := V2699
__typedArg1 := tmp6779
return Call(__e, PrimFunc(sym_d), __typedArg0, __typedArg1)
})())
return


} else {
tmp6781 := Call(__e, PrimFunc(symshen_4expt), V2699, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V2700)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V2700
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_c) {
__typedN0, __typedOK0 := TypedFloat64(tmp6781)
__typedN1, __typedOK1 := TypedFloat64(V2699)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_c) {
return TypedMaterializeNumber(TypedDivideValue(__typedN0, __typedN1))
}}
__typedArg0 := tmp6781
__typedArg1 := V2699
return Call(__e, PrimFunc(sym_c), __typedArg0, __typedArg1)
})())
return


}


}


}, 2)

tmp6786 := Call(__e, ns2_1set, symshen_4expt, tmp6777)


_ = tmp6786

tmp6787 := MakeNative(func(__e *ControlFlow) {
V2701 := __e.Get(1)
_ = V2701
tmp6788 := Call(__e, PrimFunc(symshen_4_5integer_6), V2701)


let__5252 := tmp6788
_ = let__5252

tmp6806 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5252)


var ifres6789 Obj

if True == tmp6806 {
tmp6790 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6789 = tmp6790


} else {
tmp6791 := Call(__e, PrimFunc(symshen_4_5_1out), let__5252)


let__5253 := tmp6791
_ = let__5253

tmp6792 := Call(__e, PrimFunc(symshen_4in_1_6), let__5252)


let__5254 := tmp6792
_ = let__5254

tmp6793 := Call(__e, PrimFunc(symshen_4_5stop_6), let__5254)


let__5255 := tmp6793
_ = let__5255

tmp6805 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5255)


var ifres6794 Obj

if True == tmp6805 {
tmp6795 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6794 = tmp6795


} else {
tmp6796 := Call(__e, PrimFunc(symshen_4in_1_6), let__5255)


let__5256 := tmp6796
_ = let__5256

tmp6797 := Call(__e, PrimFunc(symshen_4_5fraction_6), let__5256)


let__5257 := tmp6797
_ = let__5257

tmp6804 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5257)


var ifres6798 Obj

if True == tmp6804 {
tmp6799 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6798 = tmp6799


} else {
tmp6800 := Call(__e, PrimFunc(symshen_4_5_1out), let__5257)


let__5258 := tmp6800
_ = let__5258

tmp6801 := Call(__e, PrimFunc(symshen_4in_1_6), let__5257)


let__5259 := tmp6801
_ = let__5259

tmp6803 := Call(__e, PrimFunc(symshen_4comb), let__5259, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(let__5253)
__typedN1, __typedOK1 := TypedFloat64(let__5258)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := let__5253
__typedArg1 := let__5258
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())


ifres6798 = tmp6803


}

ifres6794 = ifres6798


}

ifres6789 = ifres6794


}

let__5251 := ifres6789
_ = let__5251

tmp6822 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5251)


if True == tmp6822 {
tmp6807 := Call(__e, PrimFunc(symshen_4_5stop_6), V2701)


let__5261 := tmp6807
_ = let__5261

tmp6818 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5261)


var ifres6808 Obj

if True == tmp6818 {
tmp6809 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6808 = tmp6809


} else {
tmp6810 := Call(__e, PrimFunc(symshen_4in_1_6), let__5261)


let__5262 := tmp6810
_ = let__5262

tmp6811 := Call(__e, PrimFunc(symshen_4_5fraction_6), let__5262)


let__5263 := tmp6811
_ = let__5263

tmp6817 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5263)


var ifres6812 Obj

if True == tmp6817 {
tmp6813 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6812 = tmp6813


} else {
tmp6814 := Call(__e, PrimFunc(symshen_4_5_1out), let__5263)


let__5264 := tmp6814
_ = let__5264

tmp6815 := Call(__e, PrimFunc(symshen_4in_1_6), let__5263)


let__5265 := tmp6815
_ = let__5265

tmp6816 := Call(__e, PrimFunc(symshen_4comb), let__5265, let__5264)


ifres6812 = tmp6816


}

ifres6808 = ifres6812


}

let__5260 := ifres6808
_ = let__5260

tmp6820 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5260)


if True == tmp6820 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5260)
return
}


} else {
__e.Return(let__5251)
return
}


}, 1)

tmp6823 := Call(__e, ns2_1set, symshen_4_5float_6, tmp6787)


_ = tmp6823

tmp6824 := MakeNative(func(__e *ControlFlow) {
V2717 := __e.Get(1)
_ = V2717
tmp6829 := Call(__e, PrimFunc(symshen_4hds_a_2), V2717, MakeInteger(46))


var ifres6825 Obj

if True == tmp6829 {
tmp6826 := Call(__e, PrimFunc(symtail), V2717)


let__5267 := tmp6826
_ = let__5267

tmp6827 := Call(__e, PrimFunc(symshen_4comb), let__5267, symshen_4skip)


ifres6825 = tmp6827


} else {
tmp6828 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6825 = tmp6828


}

let__5266 := ifres6825
_ = let__5266

tmp6831 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5266)


if True == tmp6831 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5266)
return
}


}, 1)

tmp6832 := Call(__e, ns2_1set, symshen_4_5stop_6, tmp6824)


_ = tmp6832

tmp6833 := MakeNative(func(__e *ControlFlow) {
V2720 := __e.Get(1)
_ = V2720
tmp6834 := Call(__e, PrimFunc(symshen_4_5digits_6), V2720)


let__5269 := tmp6834
_ = let__5269

tmp6841 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5269)


var ifres6835 Obj

if True == tmp6841 {
tmp6836 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6835 = tmp6836


} else {
tmp6837 := Call(__e, PrimFunc(symshen_4_5_1out), let__5269)


let__5270 := tmp6837
_ = let__5270

tmp6838 := Call(__e, PrimFunc(symshen_4in_1_6), let__5269)


let__5271 := tmp6838
_ = let__5271

tmp6839 := Call(__e, PrimFunc(symshen_4compute_1fraction), let__5270)


tmp6840 := Call(__e, PrimFunc(symshen_4comb), let__5271, tmp6839)


ifres6835 = tmp6840


}

let__5268 := ifres6835
_ = let__5268

tmp6843 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5268)


if True == tmp6843 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5268)
return
}


}, 1)

tmp6844 := Call(__e, ns2_1set, symshen_4_5fraction_6, tmp6833)


_ = tmp6844

tmp6845 := MakeNative(func(__e *ControlFlow) {
V2725 := __e.Get(1)
_ = V2725
__e.TailApply(PrimFunc(symshen_4compute_1fraction_1h), V2725, MakeInteger(-1))
return
}, 1)

tmp6846 := Call(__e, ns2_1set, symshen_4compute_1fraction, tmp6845)


_ = tmp6846

tmp6847 := MakeNative(func(__e *ControlFlow) {
V2728 := __e.Get(1)
_ = V2728
V2729 := __e.Get(2)
_ = V2729
tmp6857 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V2728)
}
__typedArg0 := Nil
__typedArg1 := V2728
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp6857 {
__e.Return(MakeInteger(0))
return
} else {
tmp6855 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2728)
}
__typedArg0 := V2728
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp6855 {
tmp6848 := Call(__e, PrimFunc(symshen_4expt), MakeInteger(10), V2729)


tmp6849 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2728)
}
__typedArg0 := V2728
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp6850 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_d) {
__typedN0, __typedOK0 := TypedFloat64(tmp6848)
__typedN1, __typedOK1 := TypedFloat64(tmp6849)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_d) {
return TypedMaterializeNumber((__typedN0 * __typedN1))
}}
__typedArg0 := tmp6848
__typedArg1 := tmp6849
return Call(__e, PrimFunc(sym_d), __typedArg0, __typedArg1)
})()

tmp6851 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2728)
}
__typedArg0 := V2728
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6853 := Call(__e, PrimFunc(symshen_4compute_1fraction_1h), tmp6851, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(V2729)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := V2729
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})())


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(tmp6850)
__typedN1, __typedOK1 := TypedFloat64(tmp6853)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := tmp6850
__typedArg1 := tmp6853
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

tmp6858 := Call(__e, ns2_1set, symshen_4compute_1fraction_1h, tmp6847)


_ = tmp6858

tmp6859 := MakeNative(func(__e *ControlFlow) {
V2730 := __e.Get(1)
_ = V2730
tmp6860 := Call(__e, PrimFunc(symshen_4_5float_6), V2730)


let__5273 := tmp6860
_ = let__5273

tmp6878 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5273)


var ifres6861 Obj

if True == tmp6878 {
tmp6862 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6861 = tmp6862


} else {
tmp6863 := Call(__e, PrimFunc(symshen_4_5_1out), let__5273)


let__5274 := tmp6863
_ = let__5274

tmp6864 := Call(__e, PrimFunc(symshen_4in_1_6), let__5273)


let__5275 := tmp6864
_ = let__5275

tmp6865 := Call(__e, PrimFunc(symshen_4_5lowE_6), let__5275)


let__5276 := tmp6865
_ = let__5276

tmp6877 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5276)


var ifres6866 Obj

if True == tmp6877 {
tmp6867 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6866 = tmp6867


} else {
tmp6868 := Call(__e, PrimFunc(symshen_4in_1_6), let__5276)


let__5277 := tmp6868
_ = let__5277

tmp6869 := Call(__e, PrimFunc(symshen_4_5log10_6), let__5277)


let__5278 := tmp6869
_ = let__5278

tmp6876 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5278)


var ifres6870 Obj

if True == tmp6876 {
tmp6871 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6870 = tmp6871


} else {
tmp6872 := Call(__e, PrimFunc(symshen_4_5_1out), let__5278)


let__5279 := tmp6872
_ = let__5279

tmp6873 := Call(__e, PrimFunc(symshen_4in_1_6), let__5278)


let__5280 := tmp6873
_ = let__5280

tmp6874 := Call(__e, PrimFunc(symshen_4compute_1E), let__5274, let__5279)


tmp6875 := Call(__e, PrimFunc(symshen_4comb), let__5280, tmp6874)


ifres6870 = tmp6875


}

ifres6866 = ifres6870


}

ifres6861 = ifres6866


}

let__5272 := ifres6861
_ = let__5272

tmp6901 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5272)


if True == tmp6901 {
tmp6879 := Call(__e, PrimFunc(symshen_4_5integer_6), V2730)


let__5282 := tmp6879
_ = let__5282

tmp6897 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5282)


var ifres6880 Obj

if True == tmp6897 {
tmp6881 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6880 = tmp6881


} else {
tmp6882 := Call(__e, PrimFunc(symshen_4_5_1out), let__5282)


let__5283 := tmp6882
_ = let__5283

tmp6883 := Call(__e, PrimFunc(symshen_4in_1_6), let__5282)


let__5284 := tmp6883
_ = let__5284

tmp6884 := Call(__e, PrimFunc(symshen_4_5lowE_6), let__5284)


let__5285 := tmp6884
_ = let__5285

tmp6896 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5285)


var ifres6885 Obj

if True == tmp6896 {
tmp6886 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6885 = tmp6886


} else {
tmp6887 := Call(__e, PrimFunc(symshen_4in_1_6), let__5285)


let__5286 := tmp6887
_ = let__5286

tmp6888 := Call(__e, PrimFunc(symshen_4_5log10_6), let__5286)


let__5287 := tmp6888
_ = let__5287

tmp6895 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5287)


var ifres6889 Obj

if True == tmp6895 {
tmp6890 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6889 = tmp6890


} else {
tmp6891 := Call(__e, PrimFunc(symshen_4_5_1out), let__5287)


let__5288 := tmp6891
_ = let__5288

tmp6892 := Call(__e, PrimFunc(symshen_4in_1_6), let__5287)


let__5289 := tmp6892
_ = let__5289

tmp6893 := Call(__e, PrimFunc(symshen_4compute_1E), let__5283, let__5288)


tmp6894 := Call(__e, PrimFunc(symshen_4comb), let__5289, tmp6893)


ifres6889 = tmp6894


}

ifres6885 = ifres6889


}

ifres6880 = ifres6885


}

let__5281 := ifres6880
_ = let__5281

tmp6899 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5281)


if True == tmp6899 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5281)
return
}


} else {
__e.Return(let__5272)
return
}


}, 1)

tmp6902 := Call(__e, ns2_1set, symshen_4_5e_1number_6, tmp6859)


_ = tmp6902

tmp6903 := MakeNative(func(__e *ControlFlow) {
V2749 := __e.Get(1)
_ = V2749
tmp6904 := Call(__e, PrimFunc(symshen_4_5plus_6), V2749)


let__5291 := tmp6904
_ = let__5291

tmp6915 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5291)


var ifres6905 Obj

if True == tmp6915 {
tmp6906 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6905 = tmp6906


} else {
tmp6907 := Call(__e, PrimFunc(symshen_4in_1_6), let__5291)


let__5292 := tmp6907
_ = let__5292

tmp6908 := Call(__e, PrimFunc(symshen_4_5log10_6), let__5292)


let__5293 := tmp6908
_ = let__5293

tmp6914 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5293)


var ifres6909 Obj

if True == tmp6914 {
tmp6910 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6909 = tmp6910


} else {
tmp6911 := Call(__e, PrimFunc(symshen_4_5_1out), let__5293)


let__5294 := tmp6911
_ = let__5294

tmp6912 := Call(__e, PrimFunc(symshen_4in_1_6), let__5293)


let__5295 := tmp6912
_ = let__5295

tmp6913 := Call(__e, PrimFunc(symshen_4comb), let__5295, let__5294)


ifres6909 = tmp6913


}

ifres6905 = ifres6909


}

let__5290 := ifres6905
_ = let__5290

tmp6941 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5290)


if True == tmp6941 {
tmp6916 := Call(__e, PrimFunc(symshen_4_5minus_6), V2749)


let__5297 := tmp6916
_ = let__5297

tmp6928 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5297)


var ifres6917 Obj

if True == tmp6928 {
tmp6918 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6917 = tmp6918


} else {
tmp6919 := Call(__e, PrimFunc(symshen_4in_1_6), let__5297)


let__5298 := tmp6919
_ = let__5298

tmp6920 := Call(__e, PrimFunc(symshen_4_5log10_6), let__5298)


let__5299 := tmp6920
_ = let__5299

tmp6927 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5299)


var ifres6921 Obj

if True == tmp6927 {
tmp6922 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6921 = tmp6922


} else {
tmp6923 := Call(__e, PrimFunc(symshen_4_5_1out), let__5299)


let__5300 := tmp6923
_ = let__5300

tmp6924 := Call(__e, PrimFunc(symshen_4in_1_6), let__5299)


let__5301 := tmp6924
_ = let__5301

tmp6926 := Call(__e, PrimFunc(symshen_4comb), let__5301, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(0))
__typedN1, __typedOK1 := TypedFloat64(let__5300)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := MakeInteger(0)
__typedArg1 := let__5300
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})())


ifres6921 = tmp6926


}

ifres6917 = ifres6921


}

let__5296 := ifres6917
_ = let__5296

tmp6939 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5296)


if True == tmp6939 {
tmp6929 := Call(__e, PrimFunc(symshen_4_5integer_6), V2749)


let__5303 := tmp6929
_ = let__5303

tmp6935 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5303)


var ifres6930 Obj

if True == tmp6935 {
tmp6931 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6930 = tmp6931


} else {
tmp6932 := Call(__e, PrimFunc(symshen_4_5_1out), let__5303)


let__5304 := tmp6932
_ = let__5304

tmp6933 := Call(__e, PrimFunc(symshen_4in_1_6), let__5303)


let__5305 := tmp6933
_ = let__5305

tmp6934 := Call(__e, PrimFunc(symshen_4comb), let__5305, let__5304)


ifres6930 = tmp6934


}

let__5302 := ifres6930
_ = let__5302

tmp6937 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5302)


if True == tmp6937 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5302)
return
}


} else {
__e.Return(let__5296)
return
}


} else {
__e.Return(let__5290)
return
}


}, 1)

tmp6942 := Call(__e, ns2_1set, symshen_4_5log10_6, tmp6903)


_ = tmp6942

tmp6943 := MakeNative(func(__e *ControlFlow) {
V2766 := __e.Get(1)
_ = V2766
tmp6948 := Call(__e, PrimFunc(symshen_4hds_a_2), V2766, MakeInteger(101))


var ifres6944 Obj

if True == tmp6948 {
tmp6945 := Call(__e, PrimFunc(symtail), V2766)


let__5307 := tmp6945
_ = let__5307

tmp6946 := Call(__e, PrimFunc(symshen_4comb), let__5307, symshen_4skip)


ifres6944 = tmp6946


} else {
tmp6947 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6944 = tmp6947


}

let__5306 := ifres6944
_ = let__5306

tmp6950 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5306)


if True == tmp6950 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5306)
return
}


}, 1)

tmp6951 := Call(__e, ns2_1set, symshen_4_5lowE_6, tmp6943)


_ = tmp6951

tmp6952 := MakeNative(func(__e *ControlFlow) {
V2769 := __e.Get(1)
_ = V2769
V2770 := __e.Get(2)
_ = V2770
tmp6953 := Call(__e, PrimFunc(symshen_4expt), MakeInteger(10), V2770)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_d) {
__typedN0, __typedOK0 := TypedFloat64(V2769)
__typedN1, __typedOK1 := TypedFloat64(tmp6953)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_d) {
return TypedMaterializeNumber((__typedN0 * __typedN1))
}}
__typedArg0 := V2769
__typedArg1 := tmp6953
return Call(__e, PrimFunc(sym_d), __typedArg0, __typedArg1)
})())
return


}, 2)

tmp6954 := Call(__e, ns2_1set, symshen_4compute_1E, tmp6952)


_ = tmp6954

tmp6955 := MakeNative(func(__e *ControlFlow) {
V2771 := __e.Get(1)
_ = V2771
tmp6956 := Call(__e, PrimFunc(symshen_4_5whitespace_6), V2771)


let__5309 := tmp6956
_ = let__5309

tmp6966 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5309)


var ifres6957 Obj

if True == tmp6966 {
tmp6958 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6957 = tmp6958


} else {
tmp6959 := Call(__e, PrimFunc(symshen_4in_1_6), let__5309)


let__5310 := tmp6959
_ = let__5310

tmp6960 := Call(__e, PrimFunc(symshen_4_5whitespaces_6), let__5310)


let__5311 := tmp6960
_ = let__5311

tmp6965 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5311)


var ifres6961 Obj

if True == tmp6965 {
tmp6962 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6961 = tmp6962


} else {
tmp6963 := Call(__e, PrimFunc(symshen_4in_1_6), let__5311)


let__5312 := tmp6963
_ = let__5312

tmp6964 := Call(__e, PrimFunc(symshen_4comb), let__5312, symshen_4skip)


ifres6961 = tmp6964


}

ifres6957 = ifres6961


}

let__5308 := ifres6957
_ = let__5308

tmp6976 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5308)


if True == tmp6976 {
tmp6967 := Call(__e, PrimFunc(symshen_4_5whitespace_6), V2771)


let__5314 := tmp6967
_ = let__5314

tmp6972 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5314)


var ifres6968 Obj

if True == tmp6972 {
tmp6969 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6968 = tmp6969


} else {
tmp6970 := Call(__e, PrimFunc(symshen_4in_1_6), let__5314)


let__5315 := tmp6970
_ = let__5315

tmp6971 := Call(__e, PrimFunc(symshen_4comb), let__5315, symshen_4skip)


ifres6968 = tmp6971


}

let__5313 := ifres6968
_ = let__5313

tmp6974 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5313)


if True == tmp6974 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5313)
return
}


} else {
__e.Return(let__5308)
return
}


}, 1)

tmp6977 := Call(__e, ns2_1set, symshen_4_5whitespaces_6, tmp6955)


_ = tmp6977

tmp6978 := MakeNative(func(__e *ControlFlow) {
V2780 := __e.Get(1)
_ = V2780
tmp6987 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2780)
}
__typedArg0 := V2780
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres6979 Obj

if True == tmp6987 {
tmp6980 := Call(__e, PrimFunc(symhead), V2780)


let__5317 := tmp6980
_ = let__5317

tmp6981 := Call(__e, PrimFunc(symtail), V2780)


let__5318 := tmp6981
_ = let__5318

tmp6985 := Call(__e, PrimFunc(symshen_4whitespace_2), let__5317)


var ifres6982 Obj

if True == tmp6985 {
tmp6983 := Call(__e, PrimFunc(symshen_4comb), let__5318, symshen_4skip)


ifres6982 = tmp6983


} else {
tmp6984 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6982 = tmp6984


}

ifres6979 = ifres6982


} else {
tmp6986 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6979 = tmp6986


}

let__5316 := ifres6979
_ = let__5316

tmp6989 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5316)


if True == tmp6989 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5316)
return
}


}, 1)

tmp6990 := Call(__e, ns2_1set, symshen_4_5whitespace_6, tmp6978)


_ = tmp6990

tmp6991 := MakeNative(func(__e *ControlFlow) {
V2786 := __e.Get(1)
_ = V2786
tmp6999 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(32), V2786)
}
__typedArg0 := MakeInteger(32)
__typedArg1 := V2786
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp6999 {
__e.Return(True)
return
} else {
tmp6997 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(13), V2786)
}
__typedArg0 := MakeInteger(13)
__typedArg1 := V2786
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp6997 {
__e.Return(True)
return
} else {
tmp6995 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(10), V2786)
}
__typedArg0 := MakeInteger(10)
__typedArg1 := V2786
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp6995 {
__e.Return(True)
return
} else {
tmp6993 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(9), V2786)
}
__typedArg0 := MakeInteger(9)
__typedArg1 := V2786
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp6993 {
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

tmp7000 := Call(__e, ns2_1set, symshen_4whitespace_2, tmp6991)


_ = tmp7000

tmp7001 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__selftop:
V2787 := __self1
_ = V2787
tmp7023 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V2787)
}
__typedArg0 := Nil
__typedArg1 := V2787
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp7023 {
__e.Return(Nil)
return
} else {
tmp7021 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2787)
}
__typedArg0 := V2787
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7017 Obj

if True == tmp7021 {
tmp7019 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2787)
}
__typedArg0 := V2787
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7020 := Call(__e, PrimFunc(symshen_4packaged_2), tmp7019)


var ifres7018 Obj

if True == tmp7020 {
ifres7018 = True


} else {
ifres7018 = False


}

ifres7017 = ifres7018


} else {
ifres7017 = False


}

if True == ifres7017 {
tmp7002 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2787)
}
__typedArg0 := V2787
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7003 := Call(__e, PrimFunc(symshen_4unpackage), tmp7002)


tmp7004 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2787)
}
__typedArg0 := V2787
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7005 := Call(__e, PrimFunc(symappend), tmp7003, tmp7004)


if PrimFunc(symshen_4unpackage_emacroexpand) == __self {
__self1 = tmp7005
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4unpackage_emacroexpand), tmp7005)
return


} else {
tmp7015 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2787)
}
__typedArg0 := V2787
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp7015 {
tmp7006 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2787)
}
__typedArg0 := V2787
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7007 := Call(__e, PrimFunc(symmacroexpand), tmp7006)


let__5319 := tmp7007
_ = let__5319

tmp7013 := Call(__e, PrimFunc(symshen_4packaged_2), let__5319)


if True == tmp7013 {
tmp7008 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2787)
}
__typedArg0 := V2787
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7009 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__5319, tmp7008)
}
__typedArg0 := let__5319
__typedArg1 := tmp7008
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

if PrimFunc(symshen_4unpackage_emacroexpand) == __self {
__self1 = tmp7009
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4unpackage_emacroexpand), tmp7009)
return


} else {
tmp7010 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2787)
}
__typedArg0 := V2787
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7011 := Call(__e, PrimFunc(symshen_4unpackage_emacroexpand), tmp7010)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__5319, tmp7011)
}
__typedArg0 := let__5319
__typedArg1 := tmp7011
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


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

tmp7024 := Call(__e, ns2_1set, symshen_4unpackage_emacroexpand, tmp7001)


_ = tmp7024

tmp7025 := MakeNative(func(__e *ControlFlow) {
V2791 := __e.Get(1)
_ = V2791
tmp7040 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2791)
}
__typedArg0 := V2791
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7027 Obj

if True == tmp7040 {
tmp7038 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2791)
}
__typedArg0 := V2791
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7039 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sympackage, tmp7038)
}
__typedArg0 := sympackage
__typedArg1 := tmp7038
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7029 Obj

if True == tmp7039 {
tmp7036 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2791)
}
__typedArg0 := V2791
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7037 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7036)
}
__typedArg0 := tmp7036
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7031 Obj

if True == tmp7037 {
tmp7033 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2791)
}
__typedArg0 := V2791
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7034 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7033)
}
__typedArg0 := tmp7033
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7035 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7034)
}
__typedArg0 := tmp7034
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7032 Obj

if True == tmp7035 {
ifres7032 = True


} else {
ifres7032 = False


}

ifres7031 = ifres7032


} else {
ifres7031 = False


}

var ifres7030 Obj

if True == ifres7031 {
ifres7030 = True


} else {
ifres7030 = False


}

ifres7029 = ifres7030


} else {
ifres7029 = False


}

var ifres7028 Obj

if True == ifres7029 {
ifres7028 = True


} else {
ifres7028 = False


}

ifres7027 = ifres7028


} else {
ifres7027 = False


}

if True == ifres7027 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 1)

tmp7041 := Call(__e, ns2_1set, symshen_4packaged_2, tmp7025)


_ = tmp7041

tmp7042 := MakeNative(func(__e *ControlFlow) {
V2794 := __e.Get(1)
_ = V2794
tmp7099 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2794)
}
__typedArg0 := V2794
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7081 Obj

if True == tmp7099 {
tmp7097 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2794)
}
__typedArg0 := V2794
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7098 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sympackage, tmp7097)
}
__typedArg0 := sympackage
__typedArg1 := tmp7097
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7083 Obj

if True == tmp7098 {
tmp7095 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2794)
}
__typedArg0 := V2794
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7096 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7095)
}
__typedArg0 := tmp7095
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7085 Obj

if True == tmp7096 {
tmp7092 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2794)
}
__typedArg0 := V2794
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7093 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7092)
}
__typedArg0 := tmp7092
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7094 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symnull, tmp7093)
}
__typedArg0 := symnull
__typedArg1 := tmp7093
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7087 Obj

if True == tmp7094 {
tmp7089 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2794)
}
__typedArg0 := V2794
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7090 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7089)
}
__typedArg0 := tmp7089
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7091 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7090)
}
__typedArg0 := tmp7090
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7088 Obj

if True == tmp7091 {
ifres7088 = True


} else {
ifres7088 = False


}

ifres7087 = ifres7088


} else {
ifres7087 = False


}

var ifres7086 Obj

if True == ifres7087 {
ifres7086 = True


} else {
ifres7086 = False


}

ifres7085 = ifres7086


} else {
ifres7085 = False


}

var ifres7084 Obj

if True == ifres7085 {
ifres7084 = True


} else {
ifres7084 = False


}

ifres7083 = ifres7084


} else {
ifres7083 = False


}

var ifres7082 Obj

if True == ifres7083 {
ifres7082 = True


} else {
ifres7082 = False


}

ifres7081 = ifres7082


} else {
ifres7081 = False


}

if True == ifres7081 {
tmp7043 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2794)
}
__typedArg0 := V2794
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7044 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7043)
}
__typedArg0 := tmp7043
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7044)
}
__typedArg0 := tmp7044
return Call(__e, PrimFunc(symtl), __typedArg0)
})())
return


} else {
tmp7079 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2794)
}
__typedArg0 := V2794
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7066 Obj

if True == tmp7079 {
tmp7077 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2794)
}
__typedArg0 := V2794
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7078 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sympackage, tmp7077)
}
__typedArg0 := sympackage
__typedArg1 := tmp7077
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7068 Obj

if True == tmp7078 {
tmp7075 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2794)
}
__typedArg0 := V2794
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7076 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7075)
}
__typedArg0 := tmp7075
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7070 Obj

if True == tmp7076 {
tmp7072 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2794)
}
__typedArg0 := V2794
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7073 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7072)
}
__typedArg0 := tmp7072
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7074 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7073)
}
__typedArg0 := tmp7073
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7071 Obj

if True == tmp7074 {
ifres7071 = True


} else {
ifres7071 = False


}

ifres7070 = ifres7071


} else {
ifres7070 = False


}

var ifres7069 Obj

if True == ifres7070 {
ifres7069 = True


} else {
ifres7069 = False


}

ifres7068 = ifres7069


} else {
ifres7068 = False


}

var ifres7067 Obj

if True == ifres7068 {
ifres7067 = True


} else {
ifres7067 = False


}

ifres7066 = ifres7067


} else {
ifres7066 = False


}

if True == ifres7066 {
tmp7045 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2794)
}
__typedArg0 := V2794
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7046 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7045)
}
__typedArg0 := tmp7045
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7047 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7046)
}
__typedArg0 := tmp7046
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7048 := Call(__e, PrimFunc(symeval), tmp7047)


let__5320 := tmp7048
_ = let__5320

tmp7049 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2794)
}
__typedArg0 := V2794
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7050 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7049)
}
__typedArg0 := tmp7049
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7051 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(tmp7050)
}
__typedArg0 := tmp7050
return Call(__e, PrimFunc(symstr), __typedArg0)
})()

tmp7052 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2794)
}
__typedArg0 := V2794
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7053 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7052)
}
__typedArg0 := tmp7052
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7054 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7053)
}
__typedArg0 := tmp7053
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7055 := Call(__e, PrimFunc(symshen_4package_1symbols), tmp7051, let__5320, tmp7054)


let__5321 := tmp7055
_ = let__5321

tmp7056 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2794)
}
__typedArg0 := V2794
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7057 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7056)
}
__typedArg0 := tmp7056
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7058 := Call(__e, PrimFunc(symshen_4record_1external), tmp7057, let__5320)


let__5322 := tmp7058
_ = let__5322

tmp7059 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2794)
}
__typedArg0 := V2794
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7060 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7059)
}
__typedArg0 := tmp7059
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7061 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2794)
}
__typedArg0 := V2794
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7062 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7061)
}
__typedArg0 := tmp7061
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7063 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7062)
}
__typedArg0 := tmp7062
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7064 := Call(__e, PrimFunc(symshen_4record_1internal), tmp7060, let__5320, tmp7063)


let__5323 := tmp7064
_ = let__5323

__e.Return(let__5321)
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

tmp7100 := Call(__e, ns2_1set, symshen_4unpackage, tmp7042)


_ = tmp7100

tmp7101 := MakeNative(func(__e *ControlFlow) {
V2799 := __e.Get(1)
_ = V2799
V2800 := __e.Get(2)
_ = V2800
V2801 := __e.Get(3)
_ = V2801
tmp7102 := MakeNative(func(__e *ControlFlow) {
tmp7103 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symget), V2799, symshen_4internal_1symbols, tmp7103)
return


}, 0)

tmp7104 := MakeNative(func(__e *ControlFlow) {
Z2803 := __e.Get(1)
_ = Z2803
__e.Return(Nil)
return
}, 1)

tmp7105 := Call(__e, try_1catch, tmp7102, tmp7104)


let__5324 := tmp7105
_ = let__5324

tmp7106 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(V2799)
}
__typedArg0 := V2799
return Call(__e, PrimFunc(symstr), __typedArg0)
})()

tmp7107 := Call(__e, PrimFunc(symshen_4internal_1symbols), tmp7106, V2800, V2801)


let__5325 := tmp7107
_ = let__5325

tmp7108 := Call(__e, PrimFunc(symunion), let__5325, let__5324)


tmp7109 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symput), V2799, symshen_4internal_1symbols, tmp7108, tmp7109)
return


}, 3)

tmp7110 := Call(__e, ns2_1set, symshen_4record_1internal, tmp7101)


_ = tmp7110

tmp7111 := MakeNative(func(__e *ControlFlow) {
V2811 := __e.Get(1)
_ = V2811
V2812 := __e.Get(2)
_ = V2812
V2813 := __e.Get(3)
_ = V2813
tmp7120 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2813)
}
__typedArg0 := V2813
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp7120 {
tmp7112 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2813)
}
__typedArg0 := V2813
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7113 := Call(__e, PrimFunc(symshen_4internal_1symbols), V2811, V2812, tmp7112)


tmp7114 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2813)
}
__typedArg0 := V2813
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7115 := Call(__e, PrimFunc(symshen_4internal_1symbols), V2811, V2812, tmp7114)


__e.TailApply(PrimFunc(symunion), tmp7113, tmp7115)
return


} else {
tmp7118 := Call(__e, PrimFunc(symshen_4internal_2), V2813, V2811, V2812)


if True == tmp7118 {
tmp7116 := Call(__e, PrimFunc(symshen_4intern_1in_1package), V2811, V2813)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7116, Nil)
}
__typedArg0 := tmp7116
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

tmp7121 := Call(__e, ns2_1set, symshen_4internal_1symbols, tmp7111)


_ = tmp7121

tmp7122 := MakeNative(func(__e *ControlFlow) {
V2814 := __e.Get(1)
_ = V2814
V2815 := __e.Get(2)
_ = V2815
tmp7123 := MakeNative(func(__e *ControlFlow) {
tmp7124 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symget), V2814, symshen_4external_1symbols, tmp7124)
return


}, 0)

tmp7125 := MakeNative(func(__e *ControlFlow) {
Z2817 := __e.Get(1)
_ = Z2817
__e.Return(Nil)
return
}, 1)

tmp7126 := Call(__e, try_1catch, tmp7123, tmp7125)


let__5326 := tmp7126
_ = let__5326

tmp7127 := Call(__e, PrimFunc(symunion), V2815, let__5326)


tmp7128 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symput), V2814, symshen_4external_1symbols, tmp7127, tmp7128)
return


}, 2)

tmp7129 := Call(__e, ns2_1set, symshen_4record_1external, tmp7122)


_ = tmp7129

tmp7130 := MakeNative(func(__e *ControlFlow) {
V2822 := __e.Get(1)
_ = V2822
V2823 := __e.Get(2)
_ = V2823
V2824 := __e.Get(3)
_ = V2824
tmp7135 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2824)
}
__typedArg0 := V2824
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp7135 {
tmp7131 := MakeNative(func(__e *ControlFlow) {
Z2825 := __e.Get(1)
_ = Z2825
__e.TailApply(PrimFunc(symshen_4package_1symbols), V2822, V2823, Z2825)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp7131, V2824)
return


} else {
tmp7133 := Call(__e, PrimFunc(symshen_4internal_2), V2824, V2822, V2823)


if True == tmp7133 {
__e.TailApply(PrimFunc(symshen_4intern_1in_1package), V2822, V2824)
return
} else {
__e.Return(V2824)
return
}


}


}, 3)

tmp7136 := Call(__e, ns2_1set, symshen_4package_1symbols, tmp7130)


_ = tmp7136

tmp7137 := MakeNative(func(__e *ControlFlow) {
V2826 := __e.Get(1)
_ = V2826
V2827 := __e.Get(2)
_ = V2827
tmp7138 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(V2827)
}
__typedArg0 := V2827
return Call(__e, PrimFunc(symstr), __typedArg0)
})()

tmp7139 := Call(__e, PrimFunc(sym_8s), MakeString("."), tmp7138)


tmp7140 := Call(__e, PrimFunc(sym_8s), V2826, tmp7139)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(tmp7140)
}
__typedArg0 := tmp7140
return Call(__e, PrimFunc(symintern), __typedArg0)
})())
return


}, 2)

tmp7141 := Call(__e, ns2_1set, symshen_4intern_1in_1package, tmp7137)


_ = tmp7141

tmp7142 := MakeNative(func(__e *ControlFlow) {
V2828 := __e.Get(1)
_ = V2828
V2829 := __e.Get(2)
_ = V2829
V2830 := __e.Get(3)
_ = V2830
tmp7172 := Call(__e, PrimFunc(symelement_2), V2828, V2830)


if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp7172)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp7172
return Call(__e, PrimFunc(symnot), __typedArg0)
})() {
tmp7169 := Call(__e, PrimFunc(symshen_4sng_2), V2828)


tmp7170 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp7169)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp7169
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres7144 Obj

if True == tmp7170 {
tmp7167 := Call(__e, PrimFunc(symshen_4dbl_2), V2828)


tmp7168 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp7167)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp7167
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres7146 Obj

if True == tmp7168 {
tmp7166 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsymbol_2) {
return PrimIsSymbol(V2828)
}
__typedArg0 := V2828
return Call(__e, PrimFunc(symsymbol_2), __typedArg0)
})()

var ifres7148 Obj

if True == tmp7166 {
tmp7164 := Call(__e, PrimFunc(symshen_4sysfunc_2), V2828)


tmp7165 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp7164)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp7164
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres7150 Obj

if True == tmp7165 {
tmp7162 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(V2828)
}
__typedArg0 := V2828
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

tmp7163 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp7162)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp7162
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres7152 Obj

if True == tmp7163 {
tmp7159 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(V2828)
}
__typedArg0 := V2828
return Call(__e, PrimFunc(symstr), __typedArg0)
})()

tmp7160 := Call(__e, PrimFunc(symshen_4internal_1to_1shen_2), tmp7159)


tmp7161 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp7160)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp7160
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres7154 Obj

if True == tmp7161 {
tmp7156 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(V2828)
}
__typedArg0 := V2828
return Call(__e, PrimFunc(symstr), __typedArg0)
})()

tmp7157 := Call(__e, PrimFunc(symshen_4internal_1to_1P_2), V2829, tmp7156)


tmp7158 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp7157)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp7157
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres7155 Obj

if True == tmp7158 {
ifres7155 = True


} else {
ifres7155 = False


}

ifres7154 = ifres7155


} else {
ifres7154 = False


}

var ifres7153 Obj

if True == ifres7154 {
ifres7153 = True


} else {
ifres7153 = False


}

ifres7152 = ifres7153


} else {
ifres7152 = False


}

var ifres7151 Obj

if True == ifres7152 {
ifres7151 = True


} else {
ifres7151 = False


}

ifres7150 = ifres7151


} else {
ifres7150 = False


}

var ifres7149 Obj

if True == ifres7150 {
ifres7149 = True


} else {
ifres7149 = False


}

ifres7148 = ifres7149


} else {
ifres7148 = False


}

var ifres7147 Obj

if True == ifres7148 {
ifres7147 = True


} else {
ifres7147 = False


}

ifres7146 = ifres7147


} else {
ifres7146 = False


}

var ifres7145 Obj

if True == ifres7146 {
ifres7145 = True


} else {
ifres7145 = False


}

ifres7144 = ifres7145


} else {
ifres7144 = False


}

if True == ifres7144 {
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

tmp7174 := Call(__e, ns2_1set, symshen_4internal_2, tmp7142)


_ = tmp7174

tmp7175 := MakeNative(func(__e *ControlFlow) {
V2835 := __e.Get(1)
_ = V2835
tmp7229 := Call(__e, PrimFunc(symshen_4_7string_2), V2835)


var ifres7177 Obj

if True == tmp7229 {
tmp7227 := Call(__e, PrimFunc(symhdstr), V2835)


tmp7228 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("s"), tmp7227)
}
__typedArg0 := MakeString("s")
__typedArg1 := tmp7227
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7179 Obj

if True == tmp7228 {
tmp7226 := Call(__e, PrimFunc(symshen_4_7string_2), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V2835)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V2835
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


var ifres7181 Obj

if True == tmp7226 {
tmp7223 := Call(__e, PrimFunc(symhdstr), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V2835)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V2835
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


tmp7224 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("h"), tmp7223)
}
__typedArg0 := MakeString("h")
__typedArg1 := tmp7223
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7183 Obj

if True == tmp7224 {
tmp7221 := Call(__e, PrimFunc(symshen_4_7string_2), (func() Obj {
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


var ifres7185 Obj

if True == tmp7221 {
tmp7217 := Call(__e, PrimFunc(symhdstr), (func() Obj {
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


tmp7218 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("e"), tmp7217)
}
__typedArg0 := MakeString("e")
__typedArg1 := tmp7217
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7187 Obj

if True == tmp7218 {
tmp7214 := Call(__e, PrimFunc(symshen_4_7string_2), (func() Obj {
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


var ifres7189 Obj

if True == tmp7214 {
tmp7209 := Call(__e, PrimFunc(symhdstr), (func() Obj {
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


tmp7210 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("n"), tmp7209)
}
__typedArg0 := MakeString("n")
__typedArg1 := tmp7209
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7191 Obj

if True == tmp7210 {
tmp7205 := Call(__e, PrimFunc(symshen_4_7string_2), (func() Obj {
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


var ifres7193 Obj

if True == tmp7205 {
tmp7199 := Call(__e, PrimFunc(symhdstr), (func() Obj {
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


tmp7200 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("."), tmp7199)
}
__typedArg0 := MakeString(".")
__typedArg1 := tmp7199
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7194 Obj

if True == tmp7200 {
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

var ifres7190 Obj

if True == ifres7191 {
ifres7190 = True


} else {
ifres7190 = False


}

ifres7189 = ifres7190


} else {
ifres7189 = False


}

var ifres7188 Obj

if True == ifres7189 {
ifres7188 = True


} else {
ifres7188 = False


}

ifres7187 = ifres7188


} else {
ifres7187 = False


}

var ifres7186 Obj

if True == ifres7187 {
ifres7186 = True


} else {
ifres7186 = False


}

ifres7185 = ifres7186


} else {
ifres7185 = False


}

var ifres7184 Obj

if True == ifres7185 {
ifres7184 = True


} else {
ifres7184 = False


}

ifres7183 = ifres7184


} else {
ifres7183 = False


}

var ifres7182 Obj

if True == ifres7183 {
ifres7182 = True


} else {
ifres7182 = False


}

ifres7181 = ifres7182


} else {
ifres7181 = False


}

var ifres7180 Obj

if True == ifres7181 {
ifres7180 = True


} else {
ifres7180 = False


}

ifres7179 = ifres7180


} else {
ifres7179 = False


}

var ifres7178 Obj

if True == ifres7179 {
ifres7178 = True


} else {
ifres7178 = False


}

ifres7177 = ifres7178


} else {
ifres7177 = False


}

if True == ifres7177 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 1)

tmp7230 := Call(__e, ns2_1set, symshen_4internal_1to_1shen_2, tmp7175)


_ = tmp7230

tmp7231 := MakeNative(func(__e *ControlFlow) {
V2836 := __e.Get(1)
_ = V2836
tmp7232 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp7233 := Call(__e, PrimFunc(symget), symshen, symshen_4external_1symbols, tmp7232)


__e.TailApply(PrimFunc(symelement_2), V2836, tmp7233)
return


}, 1)

tmp7234 := Call(__e, ns2_1set, symshen_4sysfunc_2, tmp7231)


_ = tmp7234

tmp7235 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__selftop:
V2844 := __self1
_ = V2844
V2845 := __self2
_ = V2845
tmp7256 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString(""), V2844)
}
__typedArg0 := MakeString("")
__typedArg1 := V2844
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7249 Obj

if True == tmp7256 {
tmp7255 := Call(__e, PrimFunc(symshen_4_7string_2), V2845)


var ifres7251 Obj

if True == tmp7255 {
tmp7253 := Call(__e, PrimFunc(symhdstr), V2845)


tmp7254 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("."), tmp7253)
}
__typedArg0 := MakeString(".")
__typedArg1 := tmp7253
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7252 Obj

if True == tmp7254 {
ifres7252 = True


} else {
ifres7252 = False


}

ifres7251 = ifres7252


} else {
ifres7251 = False


}

var ifres7250 Obj

if True == ifres7251 {
ifres7250 = True


} else {
ifres7250 = False


}

ifres7249 = ifres7250


} else {
ifres7249 = False


}

if True == ifres7249 {
__e.Return(True)
return
} else {
tmp7247 := Call(__e, PrimFunc(symshen_4_7string_2), V2844)


var ifres7239 Obj

if True == tmp7247 {
tmp7246 := Call(__e, PrimFunc(symshen_4_7string_2), V2845)


var ifres7241 Obj

if True == tmp7246 {
tmp7243 := Call(__e, PrimFunc(symhdstr), V2844)


tmp7244 := Call(__e, PrimFunc(symhdstr), V2845)


tmp7245 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp7243, tmp7244)
}
__typedArg0 := tmp7243
__typedArg1 := tmp7244
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7242 Obj

if True == tmp7245 {
ifres7242 = True


} else {
ifres7242 = False


}

ifres7241 = ifres7242


} else {
ifres7241 = False


}

var ifres7240 Obj

if True == ifres7241 {
ifres7240 = True


} else {
ifres7240 = False


}

ifres7239 = ifres7240


} else {
ifres7239 = False


}

if True == ifres7239 {
tmp7236 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V2844)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V2844
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()

if PrimFunc(symshen_4internal_1to_1P_2) == __self {
__self1, __self2 = tmp7236, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V2845)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V2845
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4internal_1to_1P_2), tmp7236, (func() Obj {
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

tmp7257 := Call(__e, ns2_1set, symshen_4internal_1to_1P_2, tmp7235)


_ = tmp7257

tmp7258 := MakeNative(func(__e *ControlFlow) {
V2848 := __e.Get(1)
_ = V2848
V2849 := __e.Get(2)
_ = V2849
tmp7271 := Call(__e, PrimFunc(symelement_2), V2848, V2849)


if True == tmp7271 {
__e.Return(V2848)
return
} else {
tmp7269 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2848)
}
__typedArg0 := V2848
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7265 Obj

if True == tmp7269 {
tmp7267 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2848)
}
__typedArg0 := V2848
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7268 := Call(__e, PrimFunc(symshen_4non_1application_2), tmp7267)


var ifres7266 Obj

if True == tmp7268 {
ifres7266 = True


} else {
ifres7266 = False


}

ifres7265 = ifres7266


} else {
ifres7265 = False


}

if True == ifres7265 {
tmp7259 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2848)
}
__typedArg0 := V2848
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4special_1case), tmp7259, V2848, V2849)
return


} else {
tmp7263 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2848)
}
__typedArg0 := V2848
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp7263 {
tmp7260 := MakeNative(func(__e *ControlFlow) {
Z2850 := __e.Get(1)
_ = Z2850
__e.TailApply(PrimFunc(symshen_4process_1applications), Z2850, V2849)
return
}, 1)

tmp7261 := Call(__e, PrimFunc(symmap), tmp7260, V2848)


__e.TailApply(PrimFunc(symshen_4process_1application), tmp7261, V2849)
return


} else {
__e.Return(V2848)
return
}


}


}


}, 2)

tmp7272 := Call(__e, ns2_1set, symshen_4process_1applications, tmp7258)


_ = tmp7272

tmp7273 := MakeNative(func(__e *ControlFlow) {
V2853 := __e.Get(1)
_ = V2853
tmp7283 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefine, V2853)
}
__typedArg0 := symdefine
__typedArg1 := V2853
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp7283 {
__e.Return(True)
return
} else {
tmp7281 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefun, V2853)
}
__typedArg0 := symdefun
__typedArg1 := V2853
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp7281 {
__e.Return(True)
return
} else {
tmp7279 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symsynonyms, V2853)
}
__typedArg0 := symsynonyms
__typedArg1 := V2853
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp7279 {
__e.Return(True)
return
} else {
tmp7277 := Call(__e, PrimFunc(symshen_4special_2), V2853)


if True == tmp7277 {
__e.Return(True)
return
} else {
tmp7275 := Call(__e, PrimFunc(symshen_4extraspecial_2), V2853)


if True == tmp7275 {
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

tmp7284 := Call(__e, ns2_1set, symshen_4non_1application_2, tmp7273)


_ = tmp7284

tmp7285 := MakeNative(func(__e *ControlFlow) {
V2858 := __e.Get(1)
_ = V2858
V2859 := __e.Get(2)
_ = V2859
V2860 := __e.Get(3)
_ = V2860
tmp7527 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlambda, V2858)
}
__typedArg0 := symlambda
__typedArg1 := V2858
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7505 Obj

if True == tmp7527 {
tmp7526 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7507 Obj

if True == tmp7526 {
tmp7524 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7525 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlambda, tmp7524)
}
__typedArg0 := symlambda
__typedArg1 := tmp7524
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7509 Obj

if True == tmp7525 {
tmp7522 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7523 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7522)
}
__typedArg0 := tmp7522
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7511 Obj

if True == tmp7523 {
tmp7519 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2859)
}
__typedArg0 := V2859
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7520)
}
__typedArg0 := tmp7520
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7513 Obj

if True == tmp7521 {
tmp7515 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7516 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7515)
}
__typedArg0 := tmp7515
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7517 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7516)
}
__typedArg0 := tmp7516
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7518 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp7517)
}
__typedArg0 := Nil
__typedArg1 := tmp7517
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7514 Obj

if True == tmp7518 {
ifres7514 = True


} else {
ifres7514 = False


}

ifres7513 = ifres7514


} else {
ifres7513 = False


}

var ifres7512 Obj

if True == ifres7513 {
ifres7512 = True


} else {
ifres7512 = False


}

ifres7511 = ifres7512


} else {
ifres7511 = False


}

var ifres7510 Obj

if True == ifres7511 {
ifres7510 = True


} else {
ifres7510 = False


}

ifres7509 = ifres7510


} else {
ifres7509 = False


}

var ifres7508 Obj

if True == ifres7509 {
ifres7508 = True


} else {
ifres7508 = False


}

ifres7507 = ifres7508


} else {
ifres7507 = False


}

var ifres7506 Obj

if True == ifres7507 {
ifres7506 = True


} else {
ifres7506 = False


}

ifres7505 = ifres7506


} else {
ifres7505 = False


}

if True == ifres7505 {
tmp7286 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7287 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7286)
}
__typedArg0 := tmp7286
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7288 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7289 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7288)
}
__typedArg0 := tmp7288
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7290 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7289)
}
__typedArg0 := tmp7289
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7291 := Call(__e, PrimFunc(symshen_4process_1applications), tmp7290, V2860)


tmp7292 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7291, Nil)
}
__typedArg0 := tmp7291
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7293 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7287, tmp7292)
}
__typedArg0 := tmp7287
__typedArg1 := tmp7292
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp7293)
}
__typedArg0 := symlambda
__typedArg1 := tmp7293
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp7503 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlet, V2858)
}
__typedArg0 := symlet
__typedArg1 := V2858
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7474 Obj

if True == tmp7503 {
tmp7502 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7476 Obj

if True == tmp7502 {
tmp7500 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7501 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlet, tmp7500)
}
__typedArg0 := symlet
__typedArg1 := tmp7500
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7478 Obj

if True == tmp7501 {
tmp7498 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7499 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7498)
}
__typedArg0 := tmp7498
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7480 Obj

if True == tmp7499 {
tmp7495 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7496 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7495)
}
__typedArg0 := tmp7495
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7497 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7496)
}
__typedArg0 := tmp7496
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7482 Obj

if True == tmp7497 {
tmp7491 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7492 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7491)
}
__typedArg0 := tmp7491
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7493 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7492)
}
__typedArg0 := tmp7492
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7494 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7493)
}
__typedArg0 := tmp7493
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7484 Obj

if True == tmp7494 {
tmp7486 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7487 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7486)
}
__typedArg0 := tmp7486
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7488 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7487)
}
__typedArg0 := tmp7487
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7489 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7488)
}
__typedArg0 := tmp7488
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7490 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp7489)
}
__typedArg0 := Nil
__typedArg1 := tmp7489
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7485 Obj

if True == tmp7490 {
ifres7485 = True


} else {
ifres7485 = False


}

ifres7484 = ifres7485


} else {
ifres7484 = False


}

var ifres7483 Obj

if True == ifres7484 {
ifres7483 = True


} else {
ifres7483 = False


}

ifres7482 = ifres7483


} else {
ifres7482 = False


}

var ifres7481 Obj

if True == ifres7482 {
ifres7481 = True


} else {
ifres7481 = False


}

ifres7480 = ifres7481


} else {
ifres7480 = False


}

var ifres7479 Obj

if True == ifres7480 {
ifres7479 = True


} else {
ifres7479 = False


}

ifres7478 = ifres7479


} else {
ifres7478 = False


}

var ifres7477 Obj

if True == ifres7478 {
ifres7477 = True


} else {
ifres7477 = False


}

ifres7476 = ifres7477


} else {
ifres7476 = False


}

var ifres7475 Obj

if True == ifres7476 {
ifres7475 = True


} else {
ifres7475 = False


}

ifres7474 = ifres7475


} else {
ifres7474 = False


}

if True == ifres7474 {
tmp7294 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7295 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7294)
}
__typedArg0 := tmp7294
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7296 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2859)
}
__typedArg0 := V2859
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7297)
}
__typedArg0 := tmp7297
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7299 := Call(__e, PrimFunc(symshen_4process_1applications), tmp7298, V2860)


tmp7300 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7301 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7300)
}
__typedArg0 := tmp7300
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7302 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7301)
}
__typedArg0 := tmp7301
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7303 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7302)
}
__typedArg0 := tmp7302
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7304 := Call(__e, PrimFunc(symshen_4process_1applications), tmp7303, V2860)


tmp7305 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7304, Nil)
}
__typedArg0 := tmp7304
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7306 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7299, tmp7305)
}
__typedArg0 := tmp7299
__typedArg1 := tmp7305
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7307 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7295, tmp7306)
}
__typedArg0 := tmp7295
__typedArg1 := tmp7306
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp7307)
}
__typedArg0 := symlet
__typedArg1 := tmp7307
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp7472 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefun, V2858)
}
__typedArg0 := symdefun
__typedArg1 := V2858
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7443 Obj

if True == tmp7472 {
tmp7471 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7445 Obj

if True == tmp7471 {
tmp7469 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7470 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefun, tmp7469)
}
__typedArg0 := symdefun
__typedArg1 := tmp7469
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7447 Obj

if True == tmp7470 {
tmp7467 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7468 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7467)
}
__typedArg0 := tmp7467
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7449 Obj

if True == tmp7468 {
tmp7464 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7465 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7464)
}
__typedArg0 := tmp7464
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7466 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7465)
}
__typedArg0 := tmp7465
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7451 Obj

if True == tmp7466 {
tmp7460 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7461 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7460)
}
__typedArg0 := tmp7460
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7462 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7461)
}
__typedArg0 := tmp7461
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7463 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7462)
}
__typedArg0 := tmp7462
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7453 Obj

if True == tmp7463 {
tmp7455 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7456 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7455)
}
__typedArg0 := tmp7455
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7457 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7456)
}
__typedArg0 := tmp7456
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7458 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7457)
}
__typedArg0 := tmp7457
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7459 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp7458)
}
__typedArg0 := Nil
__typedArg1 := tmp7458
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7454 Obj

if True == tmp7459 {
ifres7454 = True


} else {
ifres7454 = False


}

ifres7453 = ifres7454


} else {
ifres7453 = False


}

var ifres7452 Obj

if True == ifres7453 {
ifres7452 = True


} else {
ifres7452 = False


}

ifres7451 = ifres7452


} else {
ifres7451 = False


}

var ifres7450 Obj

if True == ifres7451 {
ifres7450 = True


} else {
ifres7450 = False


}

ifres7449 = ifres7450


} else {
ifres7449 = False


}

var ifres7448 Obj

if True == ifres7449 {
ifres7448 = True


} else {
ifres7448 = False


}

ifres7447 = ifres7448


} else {
ifres7447 = False


}

var ifres7446 Obj

if True == ifres7447 {
ifres7446 = True


} else {
ifres7446 = False


}

ifres7445 = ifres7446


} else {
ifres7445 = False


}

var ifres7444 Obj

if True == ifres7445 {
ifres7444 = True


} else {
ifres7444 = False


}

ifres7443 = ifres7444


} else {
ifres7443 = False


}

if True == ifres7443 {
__e.Return(V2859)
return
} else {
tmp7441 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefine, V2858)
}
__typedArg0 := symdefine
__typedArg1 := V2858
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7419 Obj

if True == tmp7441 {
tmp7440 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7421 Obj

if True == tmp7440 {
tmp7438 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7439 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefine, tmp7438)
}
__typedArg0 := symdefine
__typedArg1 := tmp7438
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7423 Obj

if True == tmp7439 {
tmp7436 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7437 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7436)
}
__typedArg0 := tmp7436
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7425 Obj

if True == tmp7437 {
tmp7433 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7434 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7433)
}
__typedArg0 := tmp7433
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7435 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7434)
}
__typedArg0 := tmp7434
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7427 Obj

if True == tmp7435 {
tmp7429 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7430 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7429)
}
__typedArg0 := tmp7429
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7431 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7430)
}
__typedArg0 := tmp7430
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7432 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_i, tmp7431)
}
__typedArg0 := sym_i
__typedArg1 := tmp7431
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7428 Obj

if True == tmp7432 {
ifres7428 = True


} else {
ifres7428 = False


}

ifres7427 = ifres7428


} else {
ifres7427 = False


}

var ifres7426 Obj

if True == ifres7427 {
ifres7426 = True


} else {
ifres7426 = False


}

ifres7425 = ifres7426


} else {
ifres7425 = False


}

var ifres7424 Obj

if True == ifres7425 {
ifres7424 = True


} else {
ifres7424 = False


}

ifres7423 = ifres7424


} else {
ifres7423 = False


}

var ifres7422 Obj

if True == ifres7423 {
ifres7422 = True


} else {
ifres7422 = False


}

ifres7421 = ifres7422


} else {
ifres7421 = False


}

var ifres7420 Obj

if True == ifres7421 {
ifres7420 = True


} else {
ifres7420 = False


}

ifres7419 = ifres7420


} else {
ifres7419 = False


}

if True == ifres7419 {
tmp7308 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7309 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7308)
}
__typedArg0 := tmp7308
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7310 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7311 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7310)
}
__typedArg0 := tmp7310
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7312 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7313 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7312)
}
__typedArg0 := tmp7312
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7314 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7313)
}
__typedArg0 := tmp7313
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7315 := Call(__e, PrimFunc(symshen_4process_1after_1type), tmp7311, tmp7314, V2860)


tmp7316 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_i, tmp7315)
}
__typedArg0 := sym_i
__typedArg1 := tmp7315
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7317 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7309, tmp7316)
}
__typedArg0 := tmp7309
__typedArg1 := tmp7316
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefine, tmp7317)
}
__typedArg0 := symdefine
__typedArg1 := tmp7317
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp7417 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefine, V2858)
}
__typedArg0 := symdefine
__typedArg1 := V2858
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7406 Obj

if True == tmp7417 {
tmp7416 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7408 Obj

if True == tmp7416 {
tmp7414 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7415 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefine, tmp7414)
}
__typedArg0 := symdefine
__typedArg1 := tmp7414
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7410 Obj

if True == tmp7415 {
tmp7412 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7413 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7412)
}
__typedArg0 := tmp7412
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7411 Obj

if True == tmp7413 {
ifres7411 = True


} else {
ifres7411 = False


}

ifres7410 = ifres7411


} else {
ifres7410 = False


}

var ifres7409 Obj

if True == ifres7410 {
ifres7409 = True


} else {
ifres7409 = False


}

ifres7408 = ifres7409


} else {
ifres7408 = False


}

var ifres7407 Obj

if True == ifres7408 {
ifres7407 = True


} else {
ifres7407 = False


}

ifres7406 = ifres7407


} else {
ifres7406 = False


}

if True == ifres7406 {
tmp7318 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7319 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7318)
}
__typedArg0 := tmp7318
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7320 := MakeNative(func(__e *ControlFlow) {
Z2861 := __e.Get(1)
_ = Z2861
__e.TailApply(PrimFunc(symshen_4process_1applications), Z2861, V2860)
return
}, 1)

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

tmp7323 := Call(__e, PrimFunc(symmap), tmp7320, tmp7322)


tmp7324 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7319, tmp7323)
}
__typedArg0 := tmp7319
__typedArg1 := tmp7323
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefine, tmp7324)
}
__typedArg0 := symdefine
__typedArg1 := tmp7324
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp7404 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symsynonyms, V2858)
}
__typedArg0 := symsynonyms
__typedArg1 := V2858
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp7404 {
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
tmp7402 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symtype, V2858)
}
__typedArg0 := symtype
__typedArg1 := V2858
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7380 Obj

if True == tmp7402 {
tmp7401 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7382 Obj

if True == tmp7401 {
tmp7399 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7400 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symtype, tmp7399)
}
__typedArg0 := symtype
__typedArg1 := tmp7399
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7384 Obj

if True == tmp7400 {
tmp7397 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7398 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7397)
}
__typedArg0 := tmp7397
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7386 Obj

if True == tmp7398 {
tmp7394 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7395 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7394)
}
__typedArg0 := tmp7394
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7396 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7395)
}
__typedArg0 := tmp7395
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7388 Obj

if True == tmp7396 {
tmp7390 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7391 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7390)
}
__typedArg0 := tmp7390
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7392 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7391)
}
__typedArg0 := tmp7391
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7393 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp7392)
}
__typedArg0 := Nil
__typedArg1 := tmp7392
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7389 Obj

if True == tmp7393 {
ifres7389 = True


} else {
ifres7389 = False


}

ifres7388 = ifres7389


} else {
ifres7388 = False


}

var ifres7387 Obj

if True == ifres7388 {
ifres7387 = True


} else {
ifres7387 = False


}

ifres7386 = ifres7387


} else {
ifres7386 = False


}

var ifres7385 Obj

if True == ifres7386 {
ifres7385 = True


} else {
ifres7385 = False


}

ifres7384 = ifres7385


} else {
ifres7384 = False


}

var ifres7383 Obj

if True == ifres7384 {
ifres7383 = True


} else {
ifres7383 = False


}

ifres7382 = ifres7383


} else {
ifres7382 = False


}

var ifres7381 Obj

if True == ifres7382 {
ifres7381 = True


} else {
ifres7381 = False


}

ifres7380 = ifres7381


} else {
ifres7380 = False


}

if True == ifres7380 {
tmp7325 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7326 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7325)
}
__typedArg0 := tmp7325
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7327 := Call(__e, PrimFunc(symshen_4process_1applications), tmp7326, V2860)


tmp7328 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7329 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7328)
}
__typedArg0 := tmp7328
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7330 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7327, tmp7329)
}
__typedArg0 := tmp7327
__typedArg1 := tmp7329
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtype, tmp7330)
}
__typedArg0 := symtype
__typedArg1 := tmp7330
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp7378 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(syminput_7, V2858)
}
__typedArg0 := syminput_7
__typedArg1 := V2858
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7356 Obj

if True == tmp7378 {
tmp7377 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7358 Obj

if True == tmp7377 {
tmp7375 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7376 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(syminput_7, tmp7375)
}
__typedArg0 := syminput_7
__typedArg1 := tmp7375
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7360 Obj

if True == tmp7376 {
tmp7373 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7374 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7373)
}
__typedArg0 := tmp7373
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7362 Obj

if True == tmp7374 {
tmp7370 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7371 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7370)
}
__typedArg0 := tmp7370
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7372 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7371)
}
__typedArg0 := tmp7371
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7364 Obj

if True == tmp7372 {
tmp7366 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7367 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7366)
}
__typedArg0 := tmp7366
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7368 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7367)
}
__typedArg0 := tmp7367
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7369 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp7368)
}
__typedArg0 := Nil
__typedArg1 := tmp7368
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7365 Obj

if True == tmp7369 {
ifres7365 = True


} else {
ifres7365 = False


}

ifres7364 = ifres7365


} else {
ifres7364 = False


}

var ifres7363 Obj

if True == ifres7364 {
ifres7363 = True


} else {
ifres7363 = False


}

ifres7362 = ifres7363


} else {
ifres7362 = False


}

var ifres7361 Obj

if True == ifres7362 {
ifres7361 = True


} else {
ifres7361 = False


}

ifres7360 = ifres7361


} else {
ifres7360 = False


}

var ifres7359 Obj

if True == ifres7360 {
ifres7359 = True


} else {
ifres7359 = False


}

ifres7358 = ifres7359


} else {
ifres7358 = False


}

var ifres7357 Obj

if True == ifres7358 {
ifres7357 = True


} else {
ifres7357 = False


}

ifres7356 = ifres7357


} else {
ifres7356 = False


}

if True == ifres7356 {
tmp7331 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7332 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7331)
}
__typedArg0 := tmp7331
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7333 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7334 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7333)
}
__typedArg0 := tmp7333
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7335 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7334)
}
__typedArg0 := tmp7334
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7336 := Call(__e, PrimFunc(symshen_4process_1applications), tmp7335, V2860)


tmp7337 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7336, Nil)
}
__typedArg0 := tmp7336
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7338 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7332, tmp7337)
}
__typedArg0 := tmp7332
__typedArg1 := tmp7337
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminput_7, tmp7338)
}
__typedArg0 := syminput_7
__typedArg1 := tmp7338
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp7354 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7350 Obj

if True == tmp7354 {
tmp7352 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7353 := Call(__e, PrimFunc(symshen_4special_2), tmp7352)


var ifres7351 Obj

if True == tmp7353 {
ifres7351 = True


} else {
ifres7351 = False


}

ifres7350 = ifres7351


} else {
ifres7350 = False


}

if True == ifres7350 {
tmp7339 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7340 := MakeNative(func(__e *ControlFlow) {
Z2862 := __e.Get(1)
_ = Z2862
__e.TailApply(PrimFunc(symshen_4process_1applications), Z2862, V2860)
return
}, 1)

tmp7341 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7342 := Call(__e, PrimFunc(symmap), tmp7340, tmp7341)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7339, tmp7342)
}
__typedArg0 := tmp7339
__typedArg1 := tmp7342
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp7348 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7344 Obj

if True == tmp7348 {
tmp7346 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2859)
}
__typedArg0 := V2859
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7347 := Call(__e, PrimFunc(symshen_4extraspecial_2), tmp7346)


var ifres7345 Obj

if True == tmp7347 {
ifres7345 = True


} else {
ifres7345 = False


}

ifres7344 = ifres7345


} else {
ifres7344 = False


}

if True == ifres7344 {
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

tmp7528 := Call(__e, ns2_1set, symshen_4special_1case, tmp7285)


_ = tmp7528

tmp7529 := MakeNative(func(__e *ControlFlow) {
V2865 := __e.Get(1)
_ = V2865
V2866 := __e.Get(2)
_ = V2866
V2867 := __e.Get(3)
_ = V2867
tmp7545 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2866)
}
__typedArg0 := V2866
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7541 Obj

if True == tmp7545 {
tmp7543 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2866)
}
__typedArg0 := V2866
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7544 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_j, tmp7543)
}
__typedArg0 := sym_j
__typedArg1 := tmp7543
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7542 Obj

if True == tmp7544 {
ifres7542 = True


} else {
ifres7542 = False


}

ifres7541 = ifres7542


} else {
ifres7541 = False


}

if True == ifres7541 {
tmp7530 := MakeNative(func(__e *ControlFlow) {
Z2868 := __e.Get(1)
_ = Z2868
__e.TailApply(PrimFunc(symshen_4process_1applications), Z2868, V2867)
return
}, 1)

tmp7531 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2866)
}
__typedArg0 := V2866
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7532 := Call(__e, PrimFunc(symmap), tmp7530, tmp7531)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_j, tmp7532)
}
__typedArg0 := sym_j
__typedArg1 := tmp7532
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp7539 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2866)
}
__typedArg0 := V2866
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp7539 {
tmp7533 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2866)
}
__typedArg0 := V2866
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7534 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2866)
}
__typedArg0 := V2866
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7535 := Call(__e, PrimFunc(symshen_4process_1after_1type), V2865, tmp7534, V2867)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7533, tmp7535)
}
__typedArg0 := tmp7533
__typedArg1 := tmp7535
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp7536 := Call(__e, PrimFunc(symshen_4app), V2865, MakeString("\n"), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("missing } in "))
__typedS1, __typedOK1 := TypedString(tmp7536)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("missing } in ")
__typedArg1 := tmp7536
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("missing } in "))
__typedS1, __typedOK1 := TypedString(tmp7536)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("missing } in ")
__typedArg1 := tmp7536
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}


}


}, 3)

tmp7546 := Call(__e, ns2_1set, symshen_4process_1after_1type, tmp7529)


_ = tmp7546

tmp7547 := MakeNative(func(__e *ControlFlow) {
V2869 := __e.Get(1)
_ = V2869
V2870 := __e.Get(2)
_ = V2870
tmp7590 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2869)
}
__typedArg0 := V2869
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp7590 {
tmp7548 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2869)
}
__typedArg0 := V2869
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7549 := Call(__e, PrimFunc(symarity), tmp7548)


let__5327 := tmp7549
_ = let__5327

tmp7550 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2869)
}
__typedArg0 := V2869
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7551 := Call(__e, PrimFunc(symlength), tmp7550)


let__5328 := tmp7551
_ = let__5328

tmp7588 := Call(__e, PrimFunc(symelement_2), V2869, V2870)


if True == tmp7588 {
__e.Return(V2869)
return
} else {
tmp7585 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2869)
}
__typedArg0 := V2869
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7586 := Call(__e, PrimFunc(symshen_4shen_1call_2), tmp7585)


if True == tmp7586 {
__e.Return(V2869)
return
} else {
tmp7583 := Call(__e, PrimFunc(symshen_4foreign_2), V2869)


if True == tmp7583 {
__e.TailApply(PrimFunc(symshen_4unpack_1foreign), V2869)
return
} else {
tmp7581 := Call(__e, PrimFunc(symshen_4fn_1call_2), V2869)


if True == tmp7581 {
__e.TailApply(PrimFunc(symshen_4fn_1call), V2869)
return
} else {
tmp7579 := Call(__e, PrimFunc(symshen_4zero_1place_2), V2869)


if True == tmp7579 {
__e.Return(V2869)
return
} else {
tmp7576 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2869)
}
__typedArg0 := V2869
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7577 := Call(__e, PrimFunc(symshen_4undefined_1f_2), tmp7576, let__5327)


if True == tmp7577 {
tmp7552 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2869)
}
__typedArg0 := V2869
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7553 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7552, Nil)
}
__typedArg0 := tmp7552
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7554 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfn, tmp7553)
}
__typedArg0 := symfn
__typedArg1 := tmp7553
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7555 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2869)
}
__typedArg0 := V2869
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7556 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7554, tmp7555)
}
__typedArg0 := tmp7554
__typedArg1 := tmp7555
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4simple_1curry), tmp7556)
return


} else {
tmp7573 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2869)
}
__typedArg0 := V2869
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7574 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(tmp7573)
}
__typedArg0 := tmp7573
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

if True == tmp7574 {
__e.TailApply(PrimFunc(symshen_4simple_1curry), V2869)
return
} else {
tmp7570 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2869)
}
__typedArg0 := V2869
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7571 := Call(__e, PrimFunc(symshen_4application_2), tmp7570)


if True == tmp7571 {
__e.TailApply(PrimFunc(symshen_4simple_1curry), V2869)
return
} else {
tmp7567 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2869)
}
__typedArg0 := V2869
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7568 := Call(__e, PrimFunc(symshen_4partial_1application_d_2), tmp7567, let__5327, let__5328)


if True == tmp7568 {
__e.TailApply(PrimFunc(symshen_4lambda_1function), V2869, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(let__5327)
__typedN1, __typedOK1 := TypedFloat64(let__5328)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := let__5327
__typedArg1 := let__5328
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})())
return


} else {
tmp7564 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2869)
}
__typedArg0 := V2869
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7565 := Call(__e, PrimFunc(symshen_4overapplication_2), tmp7564, let__5327, let__5328)


if True == tmp7565 {
tmp7558 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2869)
}
__typedArg0 := V2869
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7559 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7558, Nil)
}
__typedArg0 := tmp7558
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7560 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfn, tmp7559)
}
__typedArg0 := symfn
__typedArg1 := tmp7559
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7561 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2869)
}
__typedArg0 := V2869
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7562 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7560, tmp7561)
}
__typedArg0 := tmp7560
__typedArg1 := tmp7561
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4simple_1curry), tmp7562)
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

tmp7591 := Call(__e, ns2_1set, symshen_4process_1application, tmp7547)


_ = tmp7591

tmp7592 := MakeNative(func(__e *ControlFlow) {
V2873 := __e.Get(1)
_ = V2873
tmp7618 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2873)
}
__typedArg0 := V2873
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7598 Obj

if True == tmp7618 {
tmp7616 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2873)
}
__typedArg0 := V2873
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7617 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7616)
}
__typedArg0 := tmp7616
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7600 Obj

if True == tmp7617 {
tmp7613 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2873)
}
__typedArg0 := V2873
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7614 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7613)
}
__typedArg0 := tmp7613
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7615 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symforeign, tmp7614)
}
__typedArg0 := symforeign
__typedArg1 := tmp7614
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7602 Obj

if True == tmp7615 {
tmp7610 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2873)
}
__typedArg0 := V2873
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7611 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7610)
}
__typedArg0 := tmp7610
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7612 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7611)
}
__typedArg0 := tmp7611
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7604 Obj

if True == tmp7612 {
tmp7606 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2873)
}
__typedArg0 := V2873
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7607 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7606)
}
__typedArg0 := tmp7606
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7608 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7607)
}
__typedArg0 := tmp7607
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7609 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp7608)
}
__typedArg0 := Nil
__typedArg1 := tmp7608
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7605 Obj

if True == tmp7609 {
ifres7605 = True


} else {
ifres7605 = False


}

ifres7604 = ifres7605


} else {
ifres7604 = False


}

var ifres7603 Obj

if True == ifres7604 {
ifres7603 = True


} else {
ifres7603 = False


}

ifres7602 = ifres7603


} else {
ifres7602 = False


}

var ifres7601 Obj

if True == ifres7602 {
ifres7601 = True


} else {
ifres7601 = False


}

ifres7600 = ifres7601


} else {
ifres7600 = False


}

var ifres7599 Obj

if True == ifres7600 {
ifres7599 = True


} else {
ifres7599 = False


}

ifres7598 = ifres7599


} else {
ifres7598 = False


}

if True == ifres7598 {
tmp7593 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2873)
}
__typedArg0 := V2873
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7594 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7593)
}
__typedArg0 := tmp7593
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7595 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7594)
}
__typedArg0 := tmp7594
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7596 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2873)
}
__typedArg0 := V2873
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7595, tmp7596)
}
__typedArg0 := tmp7595
__typedArg1 := tmp7596
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

tmp7619 := Call(__e, ns2_1set, symshen_4unpack_1foreign, tmp7592)


_ = tmp7619

tmp7620 := MakeNative(func(__e *ControlFlow) {
V2876 := __e.Get(1)
_ = V2876
tmp7642 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2876)
}
__typedArg0 := V2876
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7622 Obj

if True == tmp7642 {
tmp7640 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2876)
}
__typedArg0 := V2876
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7641 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7640)
}
__typedArg0 := tmp7640
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7624 Obj

if True == tmp7641 {
tmp7637 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2876)
}
__typedArg0 := V2876
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7638 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7637)
}
__typedArg0 := tmp7637
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7639 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symforeign, tmp7638)
}
__typedArg0 := symforeign
__typedArg1 := tmp7638
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7626 Obj

if True == tmp7639 {
tmp7634 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2876)
}
__typedArg0 := V2876
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7635 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7634)
}
__typedArg0 := tmp7634
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7636 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7635)
}
__typedArg0 := tmp7635
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7628 Obj

if True == tmp7636 {
tmp7630 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2876)
}
__typedArg0 := V2876
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7631 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7630)
}
__typedArg0 := tmp7630
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7632 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7631)
}
__typedArg0 := tmp7631
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7633 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp7632)
}
__typedArg0 := Nil
__typedArg1 := tmp7632
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7629 Obj

if True == tmp7633 {
ifres7629 = True


} else {
ifres7629 = False


}

ifres7628 = ifres7629


} else {
ifres7628 = False


}

var ifres7627 Obj

if True == ifres7628 {
ifres7627 = True


} else {
ifres7627 = False


}

ifres7626 = ifres7627


} else {
ifres7626 = False


}

var ifres7625 Obj

if True == ifres7626 {
ifres7625 = True


} else {
ifres7625 = False


}

ifres7624 = ifres7625


} else {
ifres7624 = False


}

var ifres7623 Obj

if True == ifres7624 {
ifres7623 = True


} else {
ifres7623 = False


}

ifres7622 = ifres7623


} else {
ifres7622 = False


}

if True == ifres7622 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 1)

tmp7643 := Call(__e, ns2_1set, symshen_4foreign_2, tmp7620)


_ = tmp7643

tmp7644 := MakeNative(func(__e *ControlFlow) {
V2879 := __e.Get(1)
_ = V2879
tmp7650 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2879)
}
__typedArg0 := V2879
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7646 Obj

if True == tmp7650 {
tmp7648 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2879)
}
__typedArg0 := V2879
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7649 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp7648)
}
__typedArg0 := Nil
__typedArg1 := tmp7648
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7647 Obj

if True == tmp7649 {
ifres7647 = True


} else {
ifres7647 = False


}

ifres7646 = ifres7647


} else {
ifres7646 = False


}

if True == ifres7646 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 1)

tmp7651 := Call(__e, ns2_1set, symshen_4zero_1place_2, tmp7644)


_ = tmp7651

tmp7652 := MakeNative(func(__e *ControlFlow) {
V2880 := __e.Get(1)
_ = V2880
tmp7657 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsymbol_2) {
return PrimIsSymbol(V2880)
}
__typedArg0 := V2880
return Call(__e, PrimFunc(symsymbol_2), __typedArg0)
})()

if True == tmp7657 {
tmp7654 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(V2880)
}
__typedArg0 := V2880
return Call(__e, PrimFunc(symstr), __typedArg0)
})()

tmp7655 := Call(__e, PrimFunc(symshen_4internal_1to_1shen_2), tmp7654)


if True == tmp7655 {
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

tmp7658 := Call(__e, ns2_1set, symshen_4shen_1call_2, tmp7652)


_ = tmp7658

tmp7659 := MakeNative(func(__e *ControlFlow) {
V2885 := __e.Get(1)
_ = V2885
tmp7689 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2885)
}
__typedArg0 := V2885
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7676 Obj

if True == tmp7689 {
tmp7687 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2885)
}
__typedArg0 := V2885
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7688 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symprotect, tmp7687)
}
__typedArg0 := symprotect
__typedArg1 := tmp7687
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7678 Obj

if True == tmp7688 {
tmp7685 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2885)
}
__typedArg0 := V2885
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7686 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7685)
}
__typedArg0 := tmp7685
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7680 Obj

if True == tmp7686 {
tmp7682 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2885)
}
__typedArg0 := V2885
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7683 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7682)
}
__typedArg0 := tmp7682
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7684 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp7683)
}
__typedArg0 := Nil
__typedArg1 := tmp7683
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7681 Obj

if True == tmp7684 {
ifres7681 = True


} else {
ifres7681 = False


}

ifres7680 = ifres7681


} else {
ifres7680 = False


}

var ifres7679 Obj

if True == ifres7680 {
ifres7679 = True


} else {
ifres7679 = False


}

ifres7678 = ifres7679


} else {
ifres7678 = False


}

var ifres7677 Obj

if True == ifres7678 {
ifres7677 = True


} else {
ifres7677 = False


}

ifres7676 = ifres7677


} else {
ifres7676 = False


}

if True == ifres7676 {
__e.Return(False)
return
} else {
tmp7674 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2885)
}
__typedArg0 := V2885
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7661 Obj

if True == tmp7674 {
tmp7672 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2885)
}
__typedArg0 := V2885
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7673 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symforeign, tmp7672)
}
__typedArg0 := symforeign
__typedArg1 := tmp7672
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7663 Obj

if True == tmp7673 {
tmp7670 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2885)
}
__typedArg0 := V2885
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7671 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7670)
}
__typedArg0 := tmp7670
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7665 Obj

if True == tmp7671 {
tmp7667 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2885)
}
__typedArg0 := V2885
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7668 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7667)
}
__typedArg0 := tmp7667
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7669 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp7668)
}
__typedArg0 := Nil
__typedArg1 := tmp7668
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7666 Obj

if True == tmp7669 {
ifres7666 = True


} else {
ifres7666 = False


}

ifres7665 = ifres7666


} else {
ifres7665 = False


}

var ifres7664 Obj

if True == ifres7665 {
ifres7664 = True


} else {
ifres7664 = False


}

ifres7663 = ifres7664


} else {
ifres7663 = False


}

var ifres7662 Obj

if True == ifres7663 {
ifres7662 = True


} else {
ifres7662 = False


}

ifres7661 = ifres7662


} else {
ifres7661 = False


}

if True == ifres7661 {
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

tmp7690 := Call(__e, ns2_1set, symshen_4application_2, tmp7659)


_ = tmp7690

tmp7691 := MakeNative(func(__e *ControlFlow) {
V2890 := __e.Get(1)
_ = V2890
V2891 := __e.Get(2)
_ = V2891
tmp7699 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(-1), V2891)
}
__typedArg0 := MakeInteger(-1)
__typedArg1 := V2891
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp7699 {
tmp7697 := Call(__e, PrimFunc(symshen_4lowercase_1symbol_2), V2890)


if True == tmp7697 {
tmp7693 := Call(__e, PrimFunc(symexternal), symshen)


tmp7694 := Call(__e, PrimFunc(symelement_2), V2890, tmp7693)


if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp7694)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp7694
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

tmp7700 := Call(__e, ns2_1set, symshen_4undefined_1f_2, tmp7691)


_ = tmp7700

tmp7701 := MakeNative(func(__e *ControlFlow) {
V2892 := __e.Get(1)
_ = V2892
tmp7706 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsymbol_2) {
return PrimIsSymbol(V2892)
}
__typedArg0 := V2892
return Call(__e, PrimFunc(symsymbol_2), __typedArg0)
})()

if True == tmp7706 {
tmp7703 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(V2892)
}
__typedArg0 := V2892
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp7703)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp7703
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

tmp7707 := Call(__e, ns2_1set, symshen_4lowercase_1symbol_2, tmp7701)


_ = tmp7707

tmp7708 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__selftop:
V2893 := __self1
_ = V2893
tmp7738 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2893)
}
__typedArg0 := V2893
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7729 Obj

if True == tmp7738 {
tmp7736 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2893)
}
__typedArg0 := V2893
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7737 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7736)
}
__typedArg0 := tmp7736
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7731 Obj

if True == tmp7737 {
tmp7733 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2893)
}
__typedArg0 := V2893
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7734 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7733)
}
__typedArg0 := tmp7733
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7735 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp7734)
}
__typedArg0 := Nil
__typedArg1 := tmp7734
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7732 Obj

if True == tmp7735 {
ifres7732 = True


} else {
ifres7732 = False


}

ifres7731 = ifres7732


} else {
ifres7731 = False


}

var ifres7730 Obj

if True == ifres7731 {
ifres7730 = True


} else {
ifres7730 = False


}

ifres7729 = ifres7730


} else {
ifres7729 = False


}

if True == ifres7729 {
__e.Return(V2893)
return
} else {
tmp7727 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2893)
}
__typedArg0 := V2893
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7718 Obj

if True == tmp7727 {
tmp7725 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2893)
}
__typedArg0 := V2893
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7726 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7725)
}
__typedArg0 := tmp7725
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7720 Obj

if True == tmp7726 {
tmp7722 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2893)
}
__typedArg0 := V2893
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7723 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7722)
}
__typedArg0 := tmp7722
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7724 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7723)
}
__typedArg0 := tmp7723
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7721 Obj

if True == tmp7724 {
ifres7721 = True


} else {
ifres7721 = False


}

ifres7720 = ifres7721


} else {
ifres7720 = False


}

var ifres7719 Obj

if True == ifres7720 {
ifres7719 = True


} else {
ifres7719 = False


}

ifres7718 = ifres7719


} else {
ifres7718 = False


}

if True == ifres7718 {
tmp7709 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2893)
}
__typedArg0 := V2893
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7710 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2893)
}
__typedArg0 := V2893
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7711 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7710)
}
__typedArg0 := tmp7710
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7712 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7711, Nil)
}
__typedArg0 := tmp7711
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7713 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7709, tmp7712)
}
__typedArg0 := tmp7709
__typedArg1 := tmp7712
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7714 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2893)
}
__typedArg0 := V2893
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7715 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7714)
}
__typedArg0 := tmp7714
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7716 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7713, tmp7715)
}
__typedArg0 := tmp7713
__typedArg1 := tmp7715
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

if PrimFunc(symshen_4simple_1curry) == __self {
__self1 = tmp7716
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4simple_1curry), tmp7716)
return


} else {
__e.Return(V2893)
return
}


}


}, 1)

tmp7739 := Call(__e, ns2_1set, symshen_4simple_1curry, tmp7708)


_ = tmp7739

tmp7740 := MakeNative(func(__e *ControlFlow) {
V2894 := __e.Get(1)
_ = V2894
__e.TailApply(PrimFunc(symfn), V2894)
return
}, 1)

tmp7741 := Call(__e, ns2_1set, symfunction, tmp7740)


_ = tmp7741

tmp7742 := MakeNative(func(__e *ControlFlow) {
V2895 := __e.Get(1)
_ = V2895
tmp7750 := Call(__e, PrimFunc(symarity), V2895)


tmp7751 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp7750, MakeInteger(0))
}
__typedArg0 := tmp7750
__typedArg1 := MakeInteger(0)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp7751 {
__e.TailApply(V2895)
return
} else {
tmp7743 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dlambdatable_d)
}
__typedArg0 := symshen_4_dlambdatable_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp7744 := Call(__e, PrimFunc(symassoc), V2895, tmp7743)


let__5329 := tmp7744
_ = let__5329

tmp7748 := Call(__e, PrimFunc(symempty_2), let__5329)


if True == tmp7748 {
tmp7745 := Call(__e, PrimFunc(symshen_4app), V2895, MakeString(" is undefined\n"), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("fn: "))
__typedS1, __typedOK1 := TypedString(tmp7745)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("fn: ")
__typedArg1 := tmp7745
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("fn: "))
__typedS1, __typedOK1 := TypedString(tmp7745)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("fn: ")
__typedArg1 := tmp7745
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__5329)
}
__typedArg0 := let__5329
return Call(__e, PrimFunc(symtl), __typedArg0)
})())
return
}


}


}, 1)

tmp7752 := Call(__e, ns2_1set, symfn, tmp7742)


_ = tmp7752

tmp7753 := MakeNative(func(__e *ControlFlow) {
V2899 := __e.Get(1)
_ = V2899
tmp7783 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2899)
}
__typedArg0 := V2899
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7770 Obj

if True == tmp7783 {
tmp7781 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2899)
}
__typedArg0 := V2899
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7782 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symfn, tmp7781)
}
__typedArg0 := symfn
__typedArg1 := tmp7781
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7772 Obj

if True == tmp7782 {
tmp7779 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2899)
}
__typedArg0 := V2899
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7780 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7779)
}
__typedArg0 := tmp7779
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7774 Obj

if True == tmp7780 {
tmp7776 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2899)
}
__typedArg0 := V2899
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7777 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7776)
}
__typedArg0 := tmp7776
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7778 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp7777)
}
__typedArg0 := Nil
__typedArg1 := tmp7777
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7775 Obj

if True == tmp7778 {
ifres7775 = True


} else {
ifres7775 = False


}

ifres7774 = ifres7775


} else {
ifres7774 = False


}

var ifres7773 Obj

if True == ifres7774 {
ifres7773 = True


} else {
ifres7773 = False


}

ifres7772 = ifres7773


} else {
ifres7772 = False


}

var ifres7771 Obj

if True == ifres7772 {
ifres7771 = True


} else {
ifres7771 = False


}

ifres7770 = ifres7771


} else {
ifres7770 = False


}

if True == ifres7770 {
__e.Return(True)
return
} else {
tmp7768 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2899)
}
__typedArg0 := V2899
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7755 Obj

if True == tmp7768 {
tmp7766 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2899)
}
__typedArg0 := V2899
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7767 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symfunction, tmp7766)
}
__typedArg0 := symfunction
__typedArg1 := tmp7766
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7757 Obj

if True == tmp7767 {
tmp7764 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2899)
}
__typedArg0 := V2899
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7765 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7764)
}
__typedArg0 := tmp7764
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7759 Obj

if True == tmp7765 {
tmp7761 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2899)
}
__typedArg0 := V2899
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7762 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7761)
}
__typedArg0 := tmp7761
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7763 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp7762)
}
__typedArg0 := Nil
__typedArg1 := tmp7762
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7760 Obj

if True == tmp7763 {
ifres7760 = True


} else {
ifres7760 = False


}

ifres7759 = ifres7760


} else {
ifres7759 = False


}

var ifres7758 Obj

if True == ifres7759 {
ifres7758 = True


} else {
ifres7758 = False


}

ifres7757 = ifres7758


} else {
ifres7757 = False


}

var ifres7756 Obj

if True == ifres7757 {
ifres7756 = True


} else {
ifres7756 = False


}

ifres7755 = ifres7756


} else {
ifres7755 = False


}

if True == ifres7755 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp7784 := Call(__e, ns2_1set, symshen_4fn_1call_2, tmp7753)


_ = tmp7784

tmp7785 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__selftop:
V2900 := __self1
_ = V2900
tmp7825 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2900)
}
__typedArg0 := V2900
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7812 Obj

if True == tmp7825 {
tmp7823 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2900)
}
__typedArg0 := V2900
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7824 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symfunction, tmp7823)
}
__typedArg0 := symfunction
__typedArg1 := tmp7823
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7814 Obj

if True == tmp7824 {
tmp7821 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2900)
}
__typedArg0 := V2900
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7822 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7821)
}
__typedArg0 := tmp7821
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7816 Obj

if True == tmp7822 {
tmp7818 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2900)
}
__typedArg0 := V2900
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7819 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7818)
}
__typedArg0 := tmp7818
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7820 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp7819)
}
__typedArg0 := Nil
__typedArg1 := tmp7819
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7817 Obj

if True == tmp7820 {
ifres7817 = True


} else {
ifres7817 = False


}

ifres7816 = ifres7817


} else {
ifres7816 = False


}

var ifres7815 Obj

if True == ifres7816 {
ifres7815 = True


} else {
ifres7815 = False


}

ifres7814 = ifres7815


} else {
ifres7814 = False


}

var ifres7813 Obj

if True == ifres7814 {
ifres7813 = True


} else {
ifres7813 = False


}

ifres7812 = ifres7813


} else {
ifres7812 = False


}

if True == ifres7812 {
tmp7786 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2900)
}
__typedArg0 := V2900
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7787 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfn, tmp7786)
}
__typedArg0 := symfn
__typedArg1 := tmp7786
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

if PrimFunc(symshen_4fn_1call) == __self {
__self1 = tmp7787
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4fn_1call), tmp7787)
return


} else {
tmp7810 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2900)
}
__typedArg0 := V2900
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7797 Obj

if True == tmp7810 {
tmp7808 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2900)
}
__typedArg0 := V2900
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7809 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symfn, tmp7808)
}
__typedArg0 := symfn
__typedArg1 := tmp7808
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7799 Obj

if True == tmp7809 {
tmp7806 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2900)
}
__typedArg0 := V2900
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7807 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7806)
}
__typedArg0 := tmp7806
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7801 Obj

if True == tmp7807 {
tmp7803 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2900)
}
__typedArg0 := V2900
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7804 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7803)
}
__typedArg0 := tmp7803
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7805 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp7804)
}
__typedArg0 := Nil
__typedArg1 := tmp7804
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7802 Obj

if True == tmp7805 {
ifres7802 = True


} else {
ifres7802 = False


}

ifres7801 = ifres7802


} else {
ifres7801 = False


}

var ifres7800 Obj

if True == ifres7801 {
ifres7800 = True


} else {
ifres7800 = False


}

ifres7799 = ifres7800


} else {
ifres7799 = False


}

var ifres7798 Obj

if True == ifres7799 {
ifres7798 = True


} else {
ifres7798 = False


}

ifres7797 = ifres7798


} else {
ifres7797 = False


}

if True == ifres7797 {
tmp7788 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2900)
}
__typedArg0 := V2900
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7789 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7788)
}
__typedArg0 := tmp7788
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7790 := Call(__e, PrimFunc(symarity), tmp7789)


let__5330 := tmp7790
_ = let__5330

tmp7795 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__5330, MakeInteger(-1))
}
__typedArg0 := let__5330
__typedArg1 := MakeInteger(-1)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp7795 {
__e.Return(V2900)
return
} else {
tmp7793 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__5330, MakeInteger(0))
}
__typedArg0 := let__5330
__typedArg1 := MakeInteger(0)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp7793 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2900)
}
__typedArg0 := V2900
return Call(__e, PrimFunc(symtl), __typedArg0)
})())
return
} else {
tmp7791 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2900)
}
__typedArg0 := V2900
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4lambda_1function), tmp7791, let__5330)
return


}


}


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

tmp7826 := Call(__e, ns2_1set, symshen_4fn_1call, tmp7785)


_ = tmp7826

tmp7827 := MakeNative(func(__e *ControlFlow) {
V2902 := __e.Get(1)
_ = V2902
V2903 := __e.Get(2)
_ = V2903
V2904 := __e.Get(3)
_ = V2904
let__5331 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_6) {
__typedN0, __typedOK0 := TypedFloat64(V2903)
__typedN1, __typedOK1 := TypedFloat64(V2904)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_6) {
return TypedMaterializeBoolean((__typedN0 > __typedN1))
}}
__typedArg0 := V2903
__typedArg1 := V2904
return Call(__e, PrimFunc(sym_6), __typedArg0, __typedArg1)
})()
_ = let__5331

var ifres7834 Obj

if True == let__5331 {
tmp7842 := Call(__e, PrimFunc(symshen_4loading_2))


var ifres7836 Obj

if True == tmp7842 {
tmp7838 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1, Nil)
}
__typedArg0 := sym_1
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7839 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_7, tmp7838)
}
__typedArg0 := sym_7
__typedArg1 := tmp7838
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7840 := Call(__e, PrimFunc(symelement_2), V2902, tmp7839)


tmp7841 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp7840)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp7840
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres7837 Obj

if True == tmp7841 {
ifres7837 = True


} else {
ifres7837 = False


}

ifres7836 = ifres7837


} else {
ifres7836 = False


}

var ifres7835 Obj

if True == ifres7836 {
ifres7835 = True


} else {
ifres7835 = False


}

ifres7834 = ifres7835


} else {
ifres7834 = False


}

var ifres7829 Obj

if True == ifres7834 {
tmp7830 := Call(__e, PrimFunc(symshen_4app), V2902, MakeString("\n"), symshen_4a)


tmp7831 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("partial application of "))
__typedS1, __typedOK1 := TypedString(tmp7830)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("partial application of ")
__typedArg1 := tmp7830
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp7832 := Call(__e, PrimFunc(symstoutput))


tmp7833 := Call(__e, PrimFunc(sympr), tmp7831, tmp7832)


ifres7829 = tmp7833


} else {
ifres7829 = symshen_4skip


}

let__5332 := ifres7829
_ = let__5332

__e.Return(let__5331)
return


}, 3)

tmp7843 := Call(__e, ns2_1set, symshen_4partial_1application_d_2, tmp7827)


_ = tmp7843

tmp7844 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dloading_2_d)
}
__typedArg0 := symshen_4_dloading_2_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp7845 := Call(__e, ns2_1set, symshen_4loading_2, tmp7844)


_ = tmp7845

tmp7846 := MakeNative(func(__e *ControlFlow) {
V2911 := __e.Get(1)
_ = V2911
V2912 := __e.Get(2)
_ = V2912
V2913 := __e.Get(3)
_ = V2913
tmp7862 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(-1), V2912)
}
__typedArg0 := MakeInteger(-1)
__typedArg1 := V2912
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp7862 {
__e.Return(False)
return
} else {
let__5333 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5) {
__typedN0, __typedOK0 := TypedFloat64(V2912)
__typedN1, __typedOK1 := TypedFloat64(V2913)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_5) {
return TypedMaterializeBoolean((__typedN0 < __typedN1))
}}
__typedArg0 := V2912
__typedArg1 := V2913
return Call(__e, PrimFunc(sym_5), __typedArg0, __typedArg1)
})()
_ = let__5333

var ifres7858 Obj

if True == let__5333 {
tmp7860 := Call(__e, PrimFunc(symshen_4loading_2))


var ifres7859 Obj

if True == tmp7860 {
ifres7859 = True


} else {
ifres7859 = False


}

ifres7858 = ifres7859


} else {
ifres7858 = False


}

var ifres7848 Obj

if True == ifres7858 {
tmp7850 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V2913, MakeInteger(1))
}
__typedArg0 := V2913
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7849 Obj

if True == tmp7850 {
ifres7849 = MakeString("")


} else {
ifres7849 = MakeString("s")


}

tmp7851 := Call(__e, PrimFunc(symshen_4app), ifres7849, MakeString("\n"), symshen_4a)


tmp7853 := Call(__e, PrimFunc(symshen_4app), V2913, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(" argument"))
__typedS1, __typedOK1 := TypedString(tmp7851)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(" argument")
__typedArg1 := tmp7851
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


tmp7855 := Call(__e, PrimFunc(symshen_4app), V2911, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(" might not like "))
__typedS1, __typedOK1 := TypedString(tmp7853)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(" might not like ")
__typedArg1 := tmp7853
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


tmp7856 := Call(__e, PrimFunc(symstoutput))


tmp7857 := Call(__e, PrimFunc(sympr), tmp7855, tmp7856)


ifres7848 = tmp7857


} else {
ifres7848 = symshen_4skip


}

let__5334 := ifres7848
_ = let__5334

__e.Return(let__5333)
return


}


}, 3)

__e.TailApply(ns2_1set, symshen_4overapplication_2, tmp7846)
return




}, 0)

