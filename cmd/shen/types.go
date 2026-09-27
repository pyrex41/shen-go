package main

import . "github.com/pyrex41/shen-go/kl"

var TypesMain = MakeNative(func(__e *ControlFlow) {
tmp19031 := MakeNative(func(__e *ControlFlow) {
V5507 := __e.Get(1)
_ = V5507
V5508 := __e.Get(2)
_ = V5508
tmp19032 := Call(__e, PrimFunc(symshen_4rectify_1type), V5508)


let__19003 := tmp19032
_ = let__19003

tmp19033 := MakeNative(func(__e *ControlFlow) {
Z5511 := __e.Get(1)
_ = Z5511
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5512 := __e.Get(1)
_ = Z5512
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5513 := __e.Get(1)
_ = Z5513
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5514 := __e.Get(1)
_ = Z5514
tmp19034 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp19034

tmp19035 := Call(__e, PrimFunc(symshen_4deref), V5507, Z5511)


tmp19036 := Call(__e, PrimFunc(symreceive), tmp19035)


tmp19037 := Call(__e, PrimFunc(symshen_4deref), let__19003, Z5511)


tmp19038 := Call(__e, PrimFunc(symreceive), tmp19037)


__e.TailApply(PrimFunc(symshen_4variancy), tmp19036, tmp19038, Z5511, Z5512, Z5513, Z5514)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19039 := Call(__e, PrimFunc(symshen_4prolog_1vector))


tmp19040 := Call(__e, tmp19033, tmp19039)


tmp19041 := Call(__e, PrimFunc(symvector), MakeNumber(0))


tmp19042 := Call(__e, PrimFunc(sym_8v), MakeNumber(0), tmp19041)


tmp19043 := Call(__e, PrimFunc(sym_8v), True, tmp19042)


tmp19044 := Call(__e, tmp19040, tmp19043)


tmp19045 := Call(__e, tmp19044, MakeNumber(0))


tmp19046 := MakeNative(func(__e *ControlFlow) {
__e.Return(True)
return
}, 0)

tmp19047 := Call(__e, tmp19045, tmp19046)


let__19004 := tmp19047
_ = let__19004

tmp19048 := Call(__e, PrimFunc(symshen_4prolog_1abstraction), V5508)


tmp19049 := Call(__e, PrimFunc(symeval_1kl), tmp19048)


let__19005 := tmp19049
_ = let__19005

tmp19050 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp19051 := Call(__e, PrimFunc(symshen_4assoc_1_6), V5507, let__19005, tmp19050)


tmp19052 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp19051)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp19051
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

let__19006 := tmp19052
_ = let__19006

__e.Return(V5507)
return


}, 2)

tmp19053 := Call(__e, ns2_1set, symdeclare, tmp19031)


_ = tmp19053

tmp19054 := MakeNative(func(__e *ControlFlow) {
V5517 := __e.Get(1)
_ = V5517
V5518 := __e.Get(2)
_ = V5518
V5519 := __e.Get(3)
_ = V5519
V5520 := __e.Get(4)
_ = V5520
V5521 := __e.Get(5)
_ = V5521
V5522 := __e.Get(6)
_ = V5522
let__19007 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V5521)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V5521
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
_ = let__19007

tmp19108 := Call(__e, PrimFunc(symshen_4unlocked_2), V5520)


var ifres19056 Obj

if True == tmp19108 {
tmp19057 := Call(__e, PrimFunc(symshen_4lazyderef), V5518, V5519)


let__19009 := tmp19057
_ = let__19009

tmp19058 := MakeNative(func(__e *ControlFlow) {
Z5527 := __e.Get(1)
_ = Z5527
tmp19059 := Call(__e, PrimFunc(symshen_4newpv), V5519)


let__19011 := tmp19059
_ = let__19011

tmp19060 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp19060

tmp19061 := Call(__e, PrimFunc(symshen_4deref), V5517, V5519)


tmp19062 := Call(__e, PrimFunc(symarity), tmp19061)


tmp19063 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp19062, MakeNumber(0))
}
__typedArg0 := tmp19062
__typedArg1 := MakeNumber(0)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

tmp19064 := MakeNative(func(__e *ControlFlow) {
tmp19065 := MakeNative(func(__e *ControlFlow) {
tmp19066 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5517, Nil)
}
__typedArg0 := V5517
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19067 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfn, tmp19066)
}
__typedArg0 := symfn
__typedArg1 := tmp19066
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19068 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4variants_2), V5517, Z5527, let__19011, V5519, V5520, let__19007, V5522)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), tmp19067, let__19011, Nil, V5519, V5520, let__19007, tmp19068)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V5519, V5520, let__19007, tmp19065)
return


}, 0)

tmp19069 := Call(__e, PrimFunc(symwhen), tmp19063, V5519, V5520, let__19007, tmp19064)


__e.TailApply(PrimFunc(symshen_4gc), V5519, tmp19069)
return


}, 1)

let__19010 := tmp19058
_ = let__19010

tmp19107 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__19009)
}
__typedArg0 := let__19009
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres19070 Obj

if True == tmp19107 {
tmp19071 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__19009)
}
__typedArg0 := let__19009
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp19072 := Call(__e, PrimFunc(symshen_4lazyderef), tmp19071, V5519)


let__19012 := tmp19072
_ = let__19012

tmp19073 := MakeNative(func(__e *ControlFlow) {
tmp19074 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__19009)
}
__typedArg0 := let__19009
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp19075 := Call(__e, PrimFunc(symshen_4lazyderef), tmp19074, V5519)


let__19014 := tmp19075
_ = let__19014

tmp19076 := MakeNative(func(__e *ControlFlow) {
Z5533 := __e.Get(1)
_ = Z5533
__e.TailApply(let__19010, Z5533)
return
}, 1)

let__19015 := tmp19076
_ = let__19015

tmp19092 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__19014)
}
__typedArg0 := let__19014
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp19092 {
tmp19077 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__19014)
}
__typedArg0 := let__19014
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__19016 := tmp19077
_ = let__19016

tmp19078 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__19014)
}
__typedArg0 := let__19014
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp19079 := Call(__e, PrimFunc(symshen_4lazyderef), tmp19078, V5519)


let__19017 := tmp19079
_ = let__19017

tmp19080 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__19015, let__19016)
return
}, 0)

let__19018 := tmp19080
_ = let__19018

tmp19084 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__19017, Nil)
}
__typedArg0 := let__19017
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp19084 {
__e.TailApply(PrimFunc(symthaw), let__19018)
return
} else {
tmp19082 := Call(__e, PrimFunc(symshen_4pvar_2), let__19017)


if True == tmp19082 {
__e.TailApply(PrimFunc(symshen_4bind_b), let__19017, Nil, V5519, let__19018)
return
} else {
__e.Return(False)
return
}


}


} else {
tmp19090 := Call(__e, PrimFunc(symshen_4pvar_2), let__19014)


if True == tmp19090 {
tmp19085 := Call(__e, PrimFunc(symshen_4newpv), V5519)


let__19019 := tmp19085
_ = let__19019

tmp19086 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__19019, Nil)
}
__typedArg0 := let__19019
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19087 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__19015, let__19019)
return
}, 0)

tmp19088 := Call(__e, PrimFunc(symshen_4bind_b), let__19014, tmp19086, V5519, tmp19087)


__e.TailApply(PrimFunc(symshen_4gc), V5519, tmp19088)
return


} else {
__e.Return(False)
return
}


}


}, 0)

let__19013 := tmp19073
_ = let__19013

tmp19098 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__19012, sym_1_1_6)
}
__typedArg0 := let__19012
__typedArg1 := sym_1_1_6
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres19093 Obj

if True == tmp19098 {
tmp19094 := Call(__e, PrimFunc(symthaw), let__19013)


ifres19093 = tmp19094


} else {
tmp19097 := Call(__e, PrimFunc(symshen_4pvar_2), let__19012)


var ifres19095 Obj

if True == tmp19097 {
tmp19096 := Call(__e, PrimFunc(symshen_4bind_b), let__19012, sym_1_1_6, V5519, let__19013)


ifres19095 = tmp19096


} else {
ifres19095 = False


}

ifres19093 = ifres19095


}

ifres19070 = ifres19093


} else {
tmp19106 := Call(__e, PrimFunc(symshen_4pvar_2), let__19009)


var ifres19099 Obj

if True == tmp19106 {
tmp19100 := Call(__e, PrimFunc(symshen_4newpv), V5519)


let__19020 := tmp19100
_ = let__19020

tmp19101 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__19020, Nil)
}
__typedArg0 := let__19020
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19102 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19101)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19101
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19103 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__19010, let__19020)
return
}, 0)

tmp19104 := Call(__e, PrimFunc(symshen_4bind_b), let__19009, tmp19102, V5519, tmp19103)


tmp19105 := Call(__e, PrimFunc(symshen_4gc), V5519, tmp19104)


ifres19099 = tmp19105


} else {
ifres19099 = False


}

ifres19070 = ifres19099


}

ifres19056 = ifres19070


} else {
ifres19056 = False


}

let__19008 := ifres19056
_ = let__19008

tmp19121 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__19008, False)
}
__typedArg0 := let__19008
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp19121 {
tmp19117 := Call(__e, PrimFunc(symshen_4unlocked_2), V5520)


var ifres19109 Obj

if True == tmp19117 {
tmp19110 := Call(__e, PrimFunc(symshen_4newpv), V5519)


let__19022 := tmp19110
_ = let__19022

tmp19111 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp19111

tmp19112 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5517, Nil)
}
__typedArg0 := V5517
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19113 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfn, tmp19112)
}
__typedArg0 := symfn
__typedArg1 := tmp19112
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19114 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4variants_2), V5517, let__19022, V5518, V5519, V5520, let__19007, V5522)
return
}, 0)

tmp19115 := Call(__e, PrimFunc(symshen_4system_1S_1h), tmp19113, let__19022, Nil, V5519, V5520, let__19007, tmp19114)


tmp19116 := Call(__e, PrimFunc(symshen_4gc), V5519, tmp19115)


ifres19109 = tmp19116


} else {
ifres19109 = False


}

let__19021 := ifres19109
_ = let__19021

tmp19119 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__19021, False)
}
__typedArg0 := let__19021
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp19119 {
__e.TailApply(PrimFunc(symshen_4unlock), V5520, let__19007)
return
} else {
__e.Return(let__19021)
return
}


} else {
__e.Return(let__19008)
return
}


}, 6)

tmp19122 := Call(__e, ns2_1set, symshen_4variancy, tmp19054)


_ = tmp19122

tmp19123 := MakeNative(func(__e *ControlFlow) {
V5541 := __e.Get(1)
_ = V5541
V5542 := __e.Get(2)
_ = V5542
V5543 := __e.Get(3)
_ = V5543
V5544 := __e.Get(4)
_ = V5544
V5545 := __e.Get(5)
_ = V5545
V5546 := __e.Get(6)
_ = V5546
V5547 := __e.Get(7)
_ = V5547
tmp19127 := Call(__e, PrimFunc(symshen_4unlocked_2), V5545)


var ifres19124 Obj

if True == tmp19127 {
tmp19125 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp19125

tmp19126 := Call(__e, PrimFunc(symis_b), V5542, V5543, V5544, V5545, V5546, V5547)


ifres19124 = tmp19126


} else {
ifres19124 = False


}

let__19023 := ifres19124
_ = let__19023

tmp19139 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__19023, False)
}
__typedArg0 := let__19023
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp19139 {
tmp19137 := Call(__e, PrimFunc(symshen_4unlocked_2), V5545)


if True == tmp19137 {
tmp19128 := Call(__e, PrimFunc(symshen_4newpv), V5544)


let__19024 := tmp19128
_ = let__19024

tmp19129 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp19129

tmp19130 := Call(__e, PrimFunc(symshen_4deref), V5541, V5544)


tmp19131 := Call(__e, PrimFunc(symshen_4app), tmp19130, MakeString(" may create errors\n"), symshen_4a)


tmp19132 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("warning: changing the type of "))
__typedS1, __typedOK1 := TypedString(tmp19131)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("warning: changing the type of ")
__typedArg1 := tmp19131
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp19133 := Call(__e, PrimFunc(symstoutput))


tmp19134 := Call(__e, PrimFunc(sympr), tmp19132, tmp19133)


tmp19135 := Call(__e, PrimFunc(symis), let__19024, tmp19134, V5544, V5545, V5546, V5547)


__e.TailApply(PrimFunc(symshen_4gc), V5544, tmp19135)
return


} else {
__e.Return(False)
return
}


} else {
__e.Return(let__19023)
return
}


}, 7)

tmp19140 := Call(__e, ns2_1set, symshen_4variants_2, tmp19123)


_ = tmp19140

tmp19141 := MakeNative(func(__e *ControlFlow) {
V5550 := __e.Get(1)
_ = V5550
tmp19142 := Call(__e, PrimFunc(symgensym), symB)


let__19025 := tmp19142
_ = let__19025

tmp19143 := Call(__e, PrimFunc(symgensym), symL)


let__19026 := tmp19143
_ = let__19026

tmp19144 := Call(__e, PrimFunc(symgensym), symKey)


let__19027 := tmp19144
_ = let__19027

tmp19145 := Call(__e, PrimFunc(symgensym), symC)


let__19028 := tmp19145
_ = let__19028

tmp19146 := Call(__e, PrimFunc(symgensym), symV)


let__19029 := tmp19146
_ = let__19029

tmp19147 := Call(__e, PrimFunc(symshen_4extract_1vars), V5550)


let__19030 := tmp19147
_ = let__19030

tmp19148 := Call(__e, PrimFunc(symshen_4rcons__form), V5550)


tmp19149 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__19028, Nil)
}
__typedArg0 := let__19028
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19150 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__19027, tmp19149)
}
__typedArg0 := let__19027
__typedArg1 := tmp19149
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19151 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__19026, tmp19150)
}
__typedArg0 := let__19026
__typedArg1 := tmp19150
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19152 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__19025, tmp19151)
}
__typedArg0 := let__19025
__typedArg1 := tmp19151
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19153 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19148, tmp19152)
}
__typedArg0 := tmp19148
__typedArg1 := tmp19152
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19154 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__19029, tmp19153)
}
__typedArg0 := let__19029
__typedArg1 := tmp19153
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19155 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symis_b, tmp19154)
}
__typedArg0 := symis_b
__typedArg1 := tmp19154
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19156 := Call(__e, PrimFunc(symshen_4stpart), let__19030, tmp19155, let__19025)


tmp19157 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19156, Nil)
}
__typedArg0 := tmp19156
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19158 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__19028, tmp19157)
}
__typedArg0 := let__19028
__typedArg1 := tmp19157
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19159 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp19158)
}
__typedArg0 := symlambda
__typedArg1 := tmp19158
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19160 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19159, Nil)
}
__typedArg0 := tmp19159
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19161 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__19027, tmp19160)
}
__typedArg0 := let__19027
__typedArg1 := tmp19160
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19162 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp19161)
}
__typedArg0 := symlambda
__typedArg1 := tmp19161
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19163 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19162, Nil)
}
__typedArg0 := tmp19162
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19164 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__19026, tmp19163)
}
__typedArg0 := let__19026
__typedArg1 := tmp19163
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19165 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp19164)
}
__typedArg0 := symlambda
__typedArg1 := tmp19164
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19166 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19165, Nil)
}
__typedArg0 := tmp19165
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19167 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__19025, tmp19166)
}
__typedArg0 := let__19025
__typedArg1 := tmp19166
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19168 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp19167)
}
__typedArg0 := symlambda
__typedArg1 := tmp19167
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19169 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19168, Nil)
}
__typedArg0 := tmp19168
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19170 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__19029, tmp19169)
}
__typedArg0 := let__19029
__typedArg1 := tmp19169
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp19170)
}
__typedArg0 := symlambda
__typedArg1 := tmp19170
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp19171 := Call(__e, ns2_1set, symshen_4prolog_1abstraction, tmp19141)


_ = tmp19171

tmp19172 := MakeNative(func(__e *ControlFlow) {
V5557 := __e.Get(1)
_ = V5557
__e.Return(V5557)
return
}, 1)

tmp19173 := Call(__e, ns2_1set, symshen_4demod, tmp19172)


_ = tmp19173

tmp19174 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19175 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19174)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19174
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19176 := Call(__e, PrimFunc(symdeclare), symabort, tmp19175)


_ = tmp19176

tmp19177 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19178 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19177)
}
__typedArg0 := symlist
__typedArg1 := tmp19177
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19179 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19178, Nil)
}
__typedArg0 := tmp19178
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19180 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19179)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19179
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19181 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp19180)
}
__typedArg0 := symstring
__typedArg1 := tmp19180
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19182 := Call(__e, PrimFunc(symdeclare), symabsolute, tmp19181)


_ = tmp19182

tmp19183 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19184 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19183)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19183
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19185 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp19184)
}
__typedArg0 := symA
__typedArg1 := tmp19184
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19186 := Call(__e, PrimFunc(symdeclare), symabsvector_2, tmp19185)


_ = tmp19186

tmp19187 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19188 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19187)
}
__typedArg0 := symlist
__typedArg1 := tmp19187
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19189 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19190 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19189)
}
__typedArg0 := symlist
__typedArg1 := tmp19189
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19191 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19190, Nil)
}
__typedArg0 := tmp19190
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19192 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19191)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19191
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19193 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19188, tmp19192)
}
__typedArg0 := tmp19188
__typedArg1 := tmp19192
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19194 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19193, Nil)
}
__typedArg0 := tmp19193
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19195 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19194)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19194
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19196 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp19195)
}
__typedArg0 := symA
__typedArg1 := tmp19195
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19197 := Call(__e, PrimFunc(symdeclare), symadjoin, tmp19196)


_ = tmp19197

tmp19198 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19199 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19198)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19198
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19200 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, tmp19199)
}
__typedArg0 := symboolean
__typedArg1 := tmp19199
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19201 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19200, Nil)
}
__typedArg0 := tmp19200
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19202 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19201)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19201
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19203 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, tmp19202)
}
__typedArg0 := symboolean
__typedArg1 := tmp19202
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19204 := Call(__e, PrimFunc(symdeclare), symand, tmp19203)


_ = tmp19204

tmp19205 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19206 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19205)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19205
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19207 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp19206)
}
__typedArg0 := symsymbol
__typedArg1 := tmp19206
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19208 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19207, Nil)
}
__typedArg0 := tmp19207
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19209 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19208)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19208
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19210 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp19209)
}
__typedArg0 := symstring
__typedArg1 := tmp19209
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19211 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19210, Nil)
}
__typedArg0 := tmp19210
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19212 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19211)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19211
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19213 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp19212)
}
__typedArg0 := symA
__typedArg1 := tmp19212
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19214 := Call(__e, PrimFunc(symdeclare), symshen_4app, tmp19213)


_ = tmp19214

tmp19215 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19216 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19215)
}
__typedArg0 := symlist
__typedArg1 := tmp19215
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19217 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19218 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19217)
}
__typedArg0 := symlist
__typedArg1 := tmp19217
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19219 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19220 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19219)
}
__typedArg0 := symlist
__typedArg1 := tmp19219
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19221 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19220, Nil)
}
__typedArg0 := tmp19220
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19222 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19221)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19221
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19223 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19218, tmp19222)
}
__typedArg0 := tmp19218
__typedArg1 := tmp19222
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19224 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19223, Nil)
}
__typedArg0 := tmp19223
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19225 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19224)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19224
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19226 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19216, tmp19225)
}
__typedArg0 := tmp19216
__typedArg1 := tmp19225
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19227 := Call(__e, PrimFunc(symdeclare), symappend, tmp19226)


_ = tmp19227

tmp19228 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19229 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19228)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19228
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19230 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp19229)
}
__typedArg0 := symA
__typedArg1 := tmp19229
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19231 := Call(__e, PrimFunc(symdeclare), symarity, tmp19230)


_ = tmp19231

tmp19232 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19233 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19232)
}
__typedArg0 := symlist
__typedArg1 := tmp19232
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19234 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19233, Nil)
}
__typedArg0 := tmp19233
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19235 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19234)
}
__typedArg0 := symlist
__typedArg1 := tmp19234
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19236 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19237 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19236)
}
__typedArg0 := symlist
__typedArg1 := tmp19236
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19238 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19237, Nil)
}
__typedArg0 := tmp19237
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19239 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19238)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19238
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19240 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19235, tmp19239)
}
__typedArg0 := tmp19235
__typedArg1 := tmp19239
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19241 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19240, Nil)
}
__typedArg0 := tmp19240
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19242 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19241)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19241
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19243 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp19242)
}
__typedArg0 := symA
__typedArg1 := tmp19242
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19244 := Call(__e, PrimFunc(symdeclare), symassoc, tmp19243)


_ = tmp19244

tmp19245 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19246 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19245)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19245
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19247 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp19246)
}
__typedArg0 := symA
__typedArg1 := tmp19246
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19248 := Call(__e, PrimFunc(symdeclare), symatom_2, tmp19247)


_ = tmp19248

tmp19249 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19250 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19249)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19249
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19251 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp19250)
}
__typedArg0 := symA
__typedArg1 := tmp19250
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19252 := Call(__e, PrimFunc(symdeclare), symboolean_2, tmp19251)


_ = tmp19252

tmp19253 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19254 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19253)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19253
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19255 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp19254)
}
__typedArg0 := symstring
__typedArg1 := tmp19254
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19256 := Call(__e, PrimFunc(symdeclare), symbootstrap, tmp19255)


_ = tmp19256

tmp19257 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19258 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19257)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19257
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19259 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp19258)
}
__typedArg0 := symsymbol
__typedArg1 := tmp19258
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19260 := Call(__e, PrimFunc(symdeclare), symbound_2, tmp19259)


_ = tmp19260

tmp19261 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19262 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19261)
}
__typedArg0 := symlist
__typedArg1 := tmp19261
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19263 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19264 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19263)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19263
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19265 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19262, tmp19264)
}
__typedArg0 := tmp19262
__typedArg1 := tmp19264
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19266 := Call(__e, PrimFunc(symdeclare), symshen_4ccons_2, tmp19265)


_ = tmp19266

tmp19267 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19268 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19267)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19267
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19269 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp19268)
}
__typedArg0 := symstring
__typedArg1 := tmp19268
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19270 := Call(__e, PrimFunc(symdeclare), symcd, tmp19269)


_ = tmp19270

tmp19271 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19272 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp19271)
}
__typedArg0 := symstream
__typedArg1 := tmp19271
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19273 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symB, Nil)
}
__typedArg0 := symB
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19274 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19273)
}
__typedArg0 := symlist
__typedArg1 := tmp19273
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19275 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19274, Nil)
}
__typedArg0 := tmp19274
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19276 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19275)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19275
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19277 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19272, tmp19276)
}
__typedArg0 := tmp19272
__typedArg1 := tmp19276
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19278 := Call(__e, PrimFunc(symdeclare), symclose, tmp19277)


_ = tmp19278

tmp19279 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19280 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19279)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19279
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19281 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp19280)
}
__typedArg0 := symstring
__typedArg1 := tmp19280
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19282 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19281, Nil)
}
__typedArg0 := tmp19281
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19283 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19282)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19282
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19284 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp19283)
}
__typedArg0 := symstring
__typedArg1 := tmp19283
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19285 := Call(__e, PrimFunc(symdeclare), symcn, tmp19284)


_ = tmp19285

tmp19286 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19287 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19286)
}
__typedArg0 := symlist
__typedArg1 := tmp19286
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19288 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19289 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19288)
}
__typedArg0 := symlist
__typedArg1 := tmp19288
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19290 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symB, Nil)
}
__typedArg0 := symB
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19291 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19289, tmp19290)
}
__typedArg0 := tmp19289
__typedArg1 := tmp19290
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19292 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp19291)
}
__typedArg0 := symstr
__typedArg1 := tmp19291
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19293 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19292, Nil)
}
__typedArg0 := tmp19292
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19294 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19293)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19293
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19295 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19287, tmp19294)
}
__typedArg0 := tmp19287
__typedArg1 := tmp19294
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19296 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19297 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19296)
}
__typedArg0 := symlist
__typedArg1 := tmp19296
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19298 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symB, Nil)
}
__typedArg0 := symB
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19299 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19298)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19298
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19300 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19297, tmp19299)
}
__typedArg0 := tmp19297
__typedArg1 := tmp19299
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19301 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19300, Nil)
}
__typedArg0 := tmp19300
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19302 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19301)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19301
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19303 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19295, tmp19302)
}
__typedArg0 := tmp19295
__typedArg1 := tmp19302
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19304 := Call(__e, PrimFunc(symdeclare), symcompile, tmp19303)


_ = tmp19304

tmp19305 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19306 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19305)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19305
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19307 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp19306)
}
__typedArg0 := symA
__typedArg1 := tmp19306
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19308 := Call(__e, PrimFunc(symdeclare), symcons_2, tmp19307)


_ = tmp19308

tmp19309 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19310 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19309)
}
__typedArg0 := symlist
__typedArg1 := tmp19309
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19311 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19310, Nil)
}
__typedArg0 := tmp19310
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19312 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19311)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19311
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19313 := Call(__e, PrimFunc(symdeclare), symdatatypes, tmp19312)


_ = tmp19313

tmp19314 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19315 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19314)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19314
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19316 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp19315)
}
__typedArg0 := symsymbol
__typedArg1 := tmp19315
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19317 := Call(__e, PrimFunc(symdeclare), symdestroy, tmp19316)


_ = tmp19317

tmp19318 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19319 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19318)
}
__typedArg0 := symlist
__typedArg1 := tmp19318
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19320 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19321 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19320)
}
__typedArg0 := symlist
__typedArg1 := tmp19320
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19322 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19323 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19322)
}
__typedArg0 := symlist
__typedArg1 := tmp19322
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19324 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19323, Nil)
}
__typedArg0 := tmp19323
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19325 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19324)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19324
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19326 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19321, tmp19325)
}
__typedArg0 := tmp19321
__typedArg1 := tmp19325
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19327 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19326, Nil)
}
__typedArg0 := tmp19326
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19328 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19327)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19327
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19329 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19319, tmp19328)
}
__typedArg0 := tmp19319
__typedArg1 := tmp19328
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19330 := Call(__e, PrimFunc(symdeclare), symdifference, tmp19329)


_ = tmp19330

tmp19331 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symB, Nil)
}
__typedArg0 := symB
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19332 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19331)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19331
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19333 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symB, tmp19332)
}
__typedArg0 := symB
__typedArg1 := tmp19332
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19334 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19333, Nil)
}
__typedArg0 := tmp19333
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19335 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19334)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19334
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19336 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp19335)
}
__typedArg0 := symA
__typedArg1 := tmp19335
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19337 := Call(__e, PrimFunc(symdeclare), symdo, tmp19336)


_ = tmp19337

tmp19338 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19339 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19338)
}
__typedArg0 := symlist
__typedArg1 := tmp19338
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19340 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19341 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19340)
}
__typedArg0 := symlist
__typedArg1 := tmp19340
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19342 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symB, Nil)
}
__typedArg0 := symB
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19343 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19342)
}
__typedArg0 := symlist
__typedArg1 := tmp19342
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19344 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19343, Nil)
}
__typedArg0 := tmp19343
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19345 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19341, tmp19344)
}
__typedArg0 := tmp19341
__typedArg1 := tmp19344
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19346 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp19345)
}
__typedArg0 := symstr
__typedArg1 := tmp19345
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19347 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19346, Nil)
}
__typedArg0 := tmp19346
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19348 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19347)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19347
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19349 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19339, tmp19348)
}
__typedArg0 := tmp19339
__typedArg1 := tmp19348
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19350 := Call(__e, PrimFunc(symdeclare), sym_5e_6, tmp19349)


_ = tmp19350

tmp19351 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19352 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19351)
}
__typedArg0 := symlist
__typedArg1 := tmp19351
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19353 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symB, Nil)
}
__typedArg0 := symB
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19354 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19353)
}
__typedArg0 := symlist
__typedArg1 := tmp19353
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19355 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19356 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19355)
}
__typedArg0 := symlist
__typedArg1 := tmp19355
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19357 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19356, Nil)
}
__typedArg0 := tmp19356
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19358 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19354, tmp19357)
}
__typedArg0 := tmp19354
__typedArg1 := tmp19357
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19359 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp19358)
}
__typedArg0 := symstr
__typedArg1 := tmp19358
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19360 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19359, Nil)
}
__typedArg0 := tmp19359
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19361 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19360)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19360
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19362 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19352, tmp19361)
}
__typedArg0 := tmp19352
__typedArg1 := tmp19361
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19363 := Call(__e, PrimFunc(symdeclare), sym_5_b_6, tmp19362)


_ = tmp19363

tmp19364 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19365 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19364)
}
__typedArg0 := symlist
__typedArg1 := tmp19364
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19366 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19367 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19366)
}
__typedArg0 := symlist
__typedArg1 := tmp19366
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19368 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symB, Nil)
}
__typedArg0 := symB
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19369 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19368)
}
__typedArg0 := symlist
__typedArg1 := tmp19368
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19370 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19369, Nil)
}
__typedArg0 := tmp19369
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19371 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19367, tmp19370)
}
__typedArg0 := tmp19367
__typedArg1 := tmp19370
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19372 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp19371)
}
__typedArg0 := symstr
__typedArg1 := tmp19371
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19373 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19372, Nil)
}
__typedArg0 := tmp19372
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19374 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19373)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19373
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19375 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19365, tmp19374)
}
__typedArg0 := tmp19365
__typedArg1 := tmp19374
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19376 := Call(__e, PrimFunc(symdeclare), sym_5end_6, tmp19375)


_ = tmp19376

tmp19377 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19378 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19377)
}
__typedArg0 := symlist
__typedArg1 := tmp19377
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19379 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symB, Nil)
}
__typedArg0 := symB
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19380 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19378, tmp19379)
}
__typedArg0 := tmp19378
__typedArg1 := tmp19379
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19381 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp19380)
}
__typedArg0 := symstr
__typedArg1 := tmp19380
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19382 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19383 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19382)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19382
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19384 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19381, tmp19383)
}
__typedArg0 := tmp19381
__typedArg1 := tmp19383
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19385 := Call(__e, PrimFunc(symdeclare), symshen_4parse_1failure_2, tmp19384)


_ = tmp19385

tmp19386 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19387 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19386)
}
__typedArg0 := symlist
__typedArg1 := tmp19386
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19388 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symB, Nil)
}
__typedArg0 := symB
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19389 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19387, tmp19388)
}
__typedArg0 := tmp19387
__typedArg1 := tmp19388
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19390 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp19389)
}
__typedArg0 := symstr
__typedArg1 := tmp19389
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19391 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19390, Nil)
}
__typedArg0 := tmp19390
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19392 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19391)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19391
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19393 := Call(__e, PrimFunc(symdeclare), symshen_4parse_1failure, tmp19392)


_ = tmp19393

tmp19394 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19395 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19394)
}
__typedArg0 := symlist
__typedArg1 := tmp19394
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19396 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symB, Nil)
}
__typedArg0 := symB
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19397 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19395, tmp19396)
}
__typedArg0 := tmp19395
__typedArg1 := tmp19396
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19398 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp19397)
}
__typedArg0 := symstr
__typedArg1 := tmp19397
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19399 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symB, Nil)
}
__typedArg0 := symB
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19400 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19399)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19399
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19401 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19398, tmp19400)
}
__typedArg0 := tmp19398
__typedArg1 := tmp19400
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19402 := Call(__e, PrimFunc(symdeclare), symshen_4_5_1out, tmp19401)


_ = tmp19402

tmp19403 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19404 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19403)
}
__typedArg0 := symlist
__typedArg1 := tmp19403
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19405 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symB, Nil)
}
__typedArg0 := symB
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19406 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19404, tmp19405)
}
__typedArg0 := tmp19404
__typedArg1 := tmp19405
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19407 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp19406)
}
__typedArg0 := symstr
__typedArg1 := tmp19406
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19408 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19409 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19408)
}
__typedArg0 := symlist
__typedArg1 := tmp19408
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19410 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19409, Nil)
}
__typedArg0 := tmp19409
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19411 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19410)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19410
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19412 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19407, tmp19411)
}
__typedArg0 := tmp19407
__typedArg1 := tmp19411
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19413 := Call(__e, PrimFunc(symdeclare), symshen_4in_1_6, tmp19412)


_ = tmp19413

tmp19414 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19415 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19414)
}
__typedArg0 := symlist
__typedArg1 := tmp19414
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19416 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19417 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19416)
}
__typedArg0 := symlist
__typedArg1 := tmp19416
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19418 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symB, Nil)
}
__typedArg0 := symB
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19419 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19417, tmp19418)
}
__typedArg0 := tmp19417
__typedArg1 := tmp19418
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19420 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp19419)
}
__typedArg0 := symstr
__typedArg1 := tmp19419
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19421 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19420, Nil)
}
__typedArg0 := tmp19420
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19422 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19421)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19421
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19423 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symB, tmp19422)
}
__typedArg0 := symB
__typedArg1 := tmp19422
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19424 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19423, Nil)
}
__typedArg0 := tmp19423
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19425 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19424)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19424
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19426 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19415, tmp19425)
}
__typedArg0 := tmp19415
__typedArg1 := tmp19425
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19427 := Call(__e, PrimFunc(symdeclare), symshen_4comb, tmp19426)


_ = tmp19427

tmp19428 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19429 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19428)
}
__typedArg0 := symlist
__typedArg1 := tmp19428
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19430 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19431 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19430)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19430
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19432 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19429, tmp19431)
}
__typedArg0 := tmp19429
__typedArg1 := tmp19431
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19433 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19432, Nil)
}
__typedArg0 := tmp19432
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19434 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19433)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19433
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19435 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp19434)
}
__typedArg0 := symA
__typedArg1 := tmp19434
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19436 := Call(__e, PrimFunc(symdeclare), symelement_2, tmp19435)


_ = tmp19436

tmp19437 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19438 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19437)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19437
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19439 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp19438)
}
__typedArg0 := symA
__typedArg1 := tmp19438
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19440 := Call(__e, PrimFunc(symdeclare), symempty_2, tmp19439)


_ = tmp19440

tmp19441 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19442 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19441)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19441
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19443 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp19442)
}
__typedArg0 := symsymbol
__typedArg1 := tmp19442
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19444 := Call(__e, PrimFunc(symdeclare), symenable_1type_1theory, tmp19443)


_ = tmp19444

tmp19445 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19446 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19445)
}
__typedArg0 := symlist
__typedArg1 := tmp19445
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19447 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19446, Nil)
}
__typedArg0 := tmp19446
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19448 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19447)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19447
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19449 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp19448)
}
__typedArg0 := symsymbol
__typedArg1 := tmp19448
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19450 := Call(__e, PrimFunc(symdeclare), symexternal, tmp19449)


_ = tmp19450

tmp19451 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19452 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19451)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19451
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19453 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symexception, tmp19452)
}
__typedArg0 := symexception
__typedArg1 := tmp19452
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19454 := Call(__e, PrimFunc(symdeclare), symerror_1to_1string, tmp19453)


_ = tmp19454

tmp19455 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19456 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19455)
}
__typedArg0 := symlist
__typedArg1 := tmp19455
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19457 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19456, Nil)
}
__typedArg0 := tmp19456
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19458 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19457)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19457
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19459 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp19458)
}
__typedArg0 := symA
__typedArg1 := tmp19458
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19460 := Call(__e, PrimFunc(symdeclare), symexplode, tmp19459)


_ = tmp19460

tmp19461 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19462 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19461)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19461
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19463 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp19462)
}
__typedArg0 := symsymbol
__typedArg1 := tmp19462
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19464 := Call(__e, PrimFunc(symdeclare), symfactorise, tmp19463)


_ = tmp19464

tmp19465 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19466 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19465)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19465
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19467 := Call(__e, PrimFunc(symdeclare), symfactorise_2, tmp19466)


_ = tmp19467

tmp19468 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19469 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19468)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19468
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19470 := Call(__e, PrimFunc(symdeclare), symfail, tmp19469)


_ = tmp19470

tmp19471 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19472 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19471)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19471
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19473 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp19472)
}
__typedArg0 := symA
__typedArg1 := tmp19472
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19474 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19475 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19474)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19474
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19476 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp19475)
}
__typedArg0 := symA
__typedArg1 := tmp19475
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19477 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19476, Nil)
}
__typedArg0 := tmp19476
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19478 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19477)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19477
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19479 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19473, tmp19478)
}
__typedArg0 := tmp19473
__typedArg1 := tmp19478
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19480 := Call(__e, PrimFunc(symdeclare), symfix, tmp19479)


_ = tmp19480

tmp19481 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19482 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlazy, tmp19481)
}
__typedArg0 := symlazy
__typedArg1 := tmp19481
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19483 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19482, Nil)
}
__typedArg0 := tmp19482
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19484 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19483)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19483
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19485 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp19484)
}
__typedArg0 := symA
__typedArg1 := tmp19484
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19486 := Call(__e, PrimFunc(symdeclare), symfreeze, tmp19485)


_ = tmp19486

tmp19487 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symB, Nil)
}
__typedArg0 := symB
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19488 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_d, tmp19487)
}
__typedArg0 := sym_d
__typedArg1 := tmp19487
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19489 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp19488)
}
__typedArg0 := symA
__typedArg1 := tmp19488
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19490 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19491 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19490)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19490
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19492 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19489, tmp19491)
}
__typedArg0 := tmp19489
__typedArg1 := tmp19491
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19493 := Call(__e, PrimFunc(symdeclare), symfst, tmp19492)


_ = tmp19493

tmp19494 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19495 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19494)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19494
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19496 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp19495)
}
__typedArg0 := symsymbol
__typedArg1 := tmp19495
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19497 := Call(__e, PrimFunc(symdeclare), symgensym, tmp19496)


_ = tmp19497

tmp19498 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19499 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19498)
}
__typedArg0 := symlist
__typedArg1 := tmp19498
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19500 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19501 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19500)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19500
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19502 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp19501)
}
__typedArg0 := symA
__typedArg1 := tmp19501
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19503 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19502, Nil)
}
__typedArg0 := tmp19502
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19504 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19503)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19503
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19505 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19499, tmp19504)
}
__typedArg0 := tmp19499
__typedArg1 := tmp19504
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19506 := Call(__e, PrimFunc(symdeclare), symshen_4hds_a_2, tmp19505)


_ = tmp19506

tmp19507 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19508 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19507)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19507
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19509 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp19508)
}
__typedArg0 := symsymbol
__typedArg1 := tmp19508
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19510 := Call(__e, PrimFunc(symdeclare), symhush, tmp19509)


_ = tmp19510

tmp19511 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19512 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19511)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19511
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19513 := Call(__e, PrimFunc(symdeclare), symhush_2, tmp19512)


_ = tmp19513

tmp19514 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19515 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp19514)
}
__typedArg0 := symvector
__typedArg1 := tmp19514
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19516 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19517 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19516)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19516
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19518 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp19517)
}
__typedArg0 := symnumber
__typedArg1 := tmp19517
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19519 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19518, Nil)
}
__typedArg0 := tmp19518
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19520 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19519)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19519
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19521 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19515, tmp19520)
}
__typedArg0 := tmp19515
__typedArg1 := tmp19520
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19522 := Call(__e, PrimFunc(symdeclare), sym_5_1vector, tmp19521)


_ = tmp19522

tmp19523 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19524 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp19523)
}
__typedArg0 := symvector
__typedArg1 := tmp19523
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19525 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19526 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp19525)
}
__typedArg0 := symvector
__typedArg1 := tmp19525
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19527 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19526, Nil)
}
__typedArg0 := tmp19526
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19528 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19527)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19527
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19529 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp19528)
}
__typedArg0 := symA
__typedArg1 := tmp19528
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19530 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19529, Nil)
}
__typedArg0 := tmp19529
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19531 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19530)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19530
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19532 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp19531)
}
__typedArg0 := symnumber
__typedArg1 := tmp19531
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19533 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19532, Nil)
}
__typedArg0 := tmp19532
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19534 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19533)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19533
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19535 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19524, tmp19534)
}
__typedArg0 := tmp19524
__typedArg1 := tmp19534
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19536 := Call(__e, PrimFunc(symdeclare), symvector_1_6, tmp19535)


_ = tmp19536

tmp19537 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19538 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp19537)
}
__typedArg0 := symvector
__typedArg1 := tmp19537
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19539 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19538, Nil)
}
__typedArg0 := tmp19538
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19540 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19539)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19539
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19541 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp19540)
}
__typedArg0 := symnumber
__typedArg1 := tmp19540
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19542 := Call(__e, PrimFunc(symdeclare), symvector, tmp19541)


_ = tmp19542

tmp19543 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19544 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19543)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19543
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19545 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp19544)
}
__typedArg0 := symsymbol
__typedArg1 := tmp19544
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19546 := Call(__e, PrimFunc(symdeclare), symget_1time, tmp19545)


_ = tmp19546

tmp19547 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19548 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19547)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19547
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19549 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp19548)
}
__typedArg0 := symnumber
__typedArg1 := tmp19548
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19550 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19549, Nil)
}
__typedArg0 := tmp19549
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19551 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19550)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19550
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19552 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp19551)
}
__typedArg0 := symA
__typedArg1 := tmp19551
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19553 := Call(__e, PrimFunc(symdeclare), symhash, tmp19552)


_ = tmp19553

tmp19554 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19555 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19554)
}
__typedArg0 := symlist
__typedArg1 := tmp19554
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19556 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19557 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19556)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19556
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19558 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19555, tmp19557)
}
__typedArg0 := tmp19555
__typedArg1 := tmp19557
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19559 := Call(__e, PrimFunc(symdeclare), symhead, tmp19558)


_ = tmp19559

tmp19560 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19561 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp19560)
}
__typedArg0 := symvector
__typedArg1 := tmp19560
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19562 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19563 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19562)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19562
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19564 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19561, tmp19563)
}
__typedArg0 := tmp19561
__typedArg1 := tmp19563
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19565 := Call(__e, PrimFunc(symdeclare), symhdv, tmp19564)


_ = tmp19565

tmp19566 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19567 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19566)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19566
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19568 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp19567)
}
__typedArg0 := symstring
__typedArg1 := tmp19567
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19569 := Call(__e, PrimFunc(symdeclare), symhdstr, tmp19568)


_ = tmp19569

tmp19570 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19571 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19570)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19570
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19572 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp19571)
}
__typedArg0 := symA
__typedArg1 := tmp19571
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19573 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19572, Nil)
}
__typedArg0 := tmp19572
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19574 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19573)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19573
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19575 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp19574)
}
__typedArg0 := symA
__typedArg1 := tmp19574
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19576 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19575, Nil)
}
__typedArg0 := tmp19575
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19577 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19576)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19576
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19578 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, tmp19577)
}
__typedArg0 := symboolean
__typedArg1 := tmp19577
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19579 := Call(__e, PrimFunc(symdeclare), symif, tmp19578)


_ = tmp19579

tmp19580 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19581 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19580)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19580
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19582 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp19581)
}
__typedArg0 := symsymbol
__typedArg1 := tmp19581
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19583 := Call(__e, PrimFunc(symdeclare), symin_1package, tmp19582)


_ = tmp19583

tmp19584 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19585 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19584)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19584
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19586 := Call(__e, PrimFunc(symdeclare), symit, tmp19585)


_ = tmp19586

tmp19587 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19588 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19587)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19587
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19589 := Call(__e, PrimFunc(symdeclare), symimplementation, tmp19588)


_ = tmp19589

tmp19590 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19591 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19590)
}
__typedArg0 := symlist
__typedArg1 := tmp19590
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19592 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19593 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19592)
}
__typedArg0 := symlist
__typedArg1 := tmp19592
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19594 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19593, Nil)
}
__typedArg0 := tmp19593
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19595 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19594)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19594
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19596 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19591, tmp19595)
}
__typedArg0 := tmp19591
__typedArg1 := tmp19595
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19597 := Call(__e, PrimFunc(symdeclare), syminclude, tmp19596)


_ = tmp19597

tmp19598 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19599 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19598)
}
__typedArg0 := symlist
__typedArg1 := tmp19598
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19600 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19601 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19600)
}
__typedArg0 := symlist
__typedArg1 := tmp19600
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19602 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19601, Nil)
}
__typedArg0 := tmp19601
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19603 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19602)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19602
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19604 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19599, tmp19603)
}
__typedArg0 := tmp19599
__typedArg1 := tmp19603
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19605 := Call(__e, PrimFunc(symdeclare), syminclude_1all_1but, tmp19604)


_ = tmp19605

tmp19606 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19607 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19606)
}
__typedArg0 := symlist
__typedArg1 := tmp19606
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19608 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19607, Nil)
}
__typedArg0 := tmp19607
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19609 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19608)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19608
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19610 := Call(__e, PrimFunc(symdeclare), symincluded, tmp19609)


_ = tmp19610

tmp19611 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19612 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19611)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19611
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19613 := Call(__e, PrimFunc(symdeclare), syminferences, tmp19612)


_ = tmp19613

tmp19614 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19615 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19614)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19614
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19616 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp19615)
}
__typedArg0 := symstring
__typedArg1 := tmp19615
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19617 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19616, Nil)
}
__typedArg0 := tmp19616
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19618 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19617)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19617
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19619 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp19618)
}
__typedArg0 := symA
__typedArg1 := tmp19618
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19620 := Call(__e, PrimFunc(symdeclare), symshen_4insert, tmp19619)


_ = tmp19620

tmp19621 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19622 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19621)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19621
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19623 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp19622)
}
__typedArg0 := symA
__typedArg1 := tmp19622
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19624 := Call(__e, PrimFunc(symdeclare), syminteger_2, tmp19623)


_ = tmp19624

tmp19625 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19626 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19625)
}
__typedArg0 := symlist
__typedArg1 := tmp19625
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19627 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19626, Nil)
}
__typedArg0 := tmp19626
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19628 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19627)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19627
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19629 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp19628)
}
__typedArg0 := symsymbol
__typedArg1 := tmp19628
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19630 := Call(__e, PrimFunc(symdeclare), syminternal, tmp19629)


_ = tmp19630

tmp19631 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19632 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19631)
}
__typedArg0 := symlist
__typedArg1 := tmp19631
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19633 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19634 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19633)
}
__typedArg0 := symlist
__typedArg1 := tmp19633
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19635 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19636 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19635)
}
__typedArg0 := symlist
__typedArg1 := tmp19635
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19637 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19636, Nil)
}
__typedArg0 := tmp19636
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19638 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19637)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19637
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19639 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19634, tmp19638)
}
__typedArg0 := tmp19634
__typedArg1 := tmp19638
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19640 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19639, Nil)
}
__typedArg0 := tmp19639
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19641 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19640)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19640
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19642 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19632, tmp19641)
}
__typedArg0 := tmp19632
__typedArg1 := tmp19641
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19643 := Call(__e, PrimFunc(symdeclare), symintersection, tmp19642)


_ = tmp19643

tmp19644 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19645 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19644)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19644
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19646 := Call(__e, PrimFunc(symdeclare), symlanguage, tmp19645)


_ = tmp19646

tmp19647 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19648 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19647)
}
__typedArg0 := symlist
__typedArg1 := tmp19647
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19649 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19650 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19649)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19649
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19651 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19648, tmp19650)
}
__typedArg0 := tmp19648
__typedArg1 := tmp19650
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19652 := Call(__e, PrimFunc(symdeclare), symlength, tmp19651)


_ = tmp19652

tmp19653 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19654 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp19653)
}
__typedArg0 := symvector
__typedArg1 := tmp19653
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19655 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19656 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19655)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19655
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19657 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19654, tmp19656)
}
__typedArg0 := tmp19654
__typedArg1 := tmp19656
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19658 := Call(__e, PrimFunc(symdeclare), symlimit, tmp19657)


_ = tmp19658

tmp19659 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symin, Nil)
}
__typedArg0 := symin
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19660 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp19659)
}
__typedArg0 := symstream
__typedArg1 := tmp19659
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19661 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunit, Nil)
}
__typedArg0 := symunit
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19662 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19661)
}
__typedArg0 := symlist
__typedArg1 := tmp19661
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19663 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19662, Nil)
}
__typedArg0 := tmp19662
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19664 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19663)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19663
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19665 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19660, tmp19664)
}
__typedArg0 := tmp19660
__typedArg1 := tmp19664
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19666 := Call(__e, PrimFunc(symdeclare), symlineread, tmp19665)


_ = tmp19666

tmp19667 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19668 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19667)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19667
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19669 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp19668)
}
__typedArg0 := symstring
__typedArg1 := tmp19668
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19670 := Call(__e, PrimFunc(symdeclare), symload, tmp19669)


_ = tmp19670

tmp19671 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symB, Nil)
}
__typedArg0 := symB
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19672 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19671)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19671
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19673 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp19672)
}
__typedArg0 := symA
__typedArg1 := tmp19672
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19674 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19675 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19674)
}
__typedArg0 := symlist
__typedArg1 := tmp19674
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19676 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symB, Nil)
}
__typedArg0 := symB
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19677 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19676)
}
__typedArg0 := symlist
__typedArg1 := tmp19676
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19678 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19677, Nil)
}
__typedArg0 := tmp19677
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19679 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19678)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19678
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19680 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19675, tmp19679)
}
__typedArg0 := tmp19675
__typedArg1 := tmp19679
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19681 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19680, Nil)
}
__typedArg0 := tmp19680
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19682 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19681)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19681
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19683 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19673, tmp19682)
}
__typedArg0 := tmp19673
__typedArg1 := tmp19682
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19684 := Call(__e, PrimFunc(symdeclare), symmap, tmp19683)


_ = tmp19684

tmp19685 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symB, Nil)
}
__typedArg0 := symB
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19686 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19685)
}
__typedArg0 := symlist
__typedArg1 := tmp19685
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19687 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19686, Nil)
}
__typedArg0 := tmp19686
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19688 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19687)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19687
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19689 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp19688)
}
__typedArg0 := symA
__typedArg1 := tmp19688
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19690 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19691 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19690)
}
__typedArg0 := symlist
__typedArg1 := tmp19690
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19692 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symB, Nil)
}
__typedArg0 := symB
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19693 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19692)
}
__typedArg0 := symlist
__typedArg1 := tmp19692
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19694 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19693, Nil)
}
__typedArg0 := tmp19693
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19695 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19694)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19694
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19696 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19691, tmp19695)
}
__typedArg0 := tmp19691
__typedArg1 := tmp19695
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19697 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19696, Nil)
}
__typedArg0 := tmp19696
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19698 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19697)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19697
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19699 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19689, tmp19698)
}
__typedArg0 := tmp19689
__typedArg1 := tmp19698
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19700 := Call(__e, PrimFunc(symdeclare), symmapcan, tmp19699)


_ = tmp19700

tmp19701 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19702 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19701)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19701
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19703 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp19702)
}
__typedArg0 := symnumber
__typedArg1 := tmp19702
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19704 := Call(__e, PrimFunc(symdeclare), symmaxinferences, tmp19703)


_ = tmp19704

tmp19705 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19706 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19705)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19705
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19707 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp19706)
}
__typedArg0 := symnumber
__typedArg1 := tmp19706
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19708 := Call(__e, PrimFunc(symdeclare), symn_1_6string, tmp19707)


_ = tmp19708

tmp19709 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19710 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19709)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19709
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19711 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp19710)
}
__typedArg0 := symnumber
__typedArg1 := tmp19710
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19712 := Call(__e, PrimFunc(symdeclare), symnl, tmp19711)


_ = tmp19712

tmp19713 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19714 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19713)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19713
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19715 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, tmp19714)
}
__typedArg0 := symboolean
__typedArg1 := tmp19714
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19716 := Call(__e, PrimFunc(symdeclare), symnot, tmp19715)


_ = tmp19716

tmp19717 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19718 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19717)
}
__typedArg0 := symlist
__typedArg1 := tmp19717
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19719 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19720 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19719)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19719
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19721 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19718, tmp19720)
}
__typedArg0 := tmp19718
__typedArg1 := tmp19720
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19722 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19721, Nil)
}
__typedArg0 := tmp19721
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19723 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19722)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19722
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19724 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp19723)
}
__typedArg0 := symnumber
__typedArg1 := tmp19723
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19725 := Call(__e, PrimFunc(symdeclare), symnth, tmp19724)


_ = tmp19725

tmp19726 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19727 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19726)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19726
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19728 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp19727)
}
__typedArg0 := symA
__typedArg1 := tmp19727
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19729 := Call(__e, PrimFunc(symdeclare), symnumber_2, tmp19728)


_ = tmp19729

tmp19730 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19731 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19730)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19730
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19732 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symB, tmp19731)
}
__typedArg0 := symB
__typedArg1 := tmp19731
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19733 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19732, Nil)
}
__typedArg0 := tmp19732
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19734 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19733)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19733
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19735 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp19734)
}
__typedArg0 := symA
__typedArg1 := tmp19734
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19736 := Call(__e, PrimFunc(symdeclare), symoccurrences, tmp19735)


_ = tmp19736

tmp19737 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19738 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19737)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19737
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19739 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp19738)
}
__typedArg0 := symsymbol
__typedArg1 := tmp19738
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19740 := Call(__e, PrimFunc(symdeclare), symoccurs_1check, tmp19739)


_ = tmp19740

tmp19741 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19742 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19741)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19741
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19743 := Call(__e, PrimFunc(symdeclare), symoccurs_2, tmp19742)


_ = tmp19743

tmp19744 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19745 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19744)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19744
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19746 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp19745)
}
__typedArg0 := symsymbol
__typedArg1 := tmp19745
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19747 := Call(__e, PrimFunc(symdeclare), symoptimise, tmp19746)


_ = tmp19747

tmp19748 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19749 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19748)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19748
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19750 := Call(__e, PrimFunc(symdeclare), symoptimise_2, tmp19749)


_ = tmp19750

tmp19751 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19752 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19751)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19751
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19753 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, tmp19752)
}
__typedArg0 := symboolean
__typedArg1 := tmp19752
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19754 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19753, Nil)
}
__typedArg0 := tmp19753
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19755 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19754)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19754
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19756 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, tmp19755)
}
__typedArg0 := symboolean
__typedArg1 := tmp19755
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19757 := Call(__e, PrimFunc(symdeclare), symor, tmp19756)


_ = tmp19757

tmp19758 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19759 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19758)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19758
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19760 := Call(__e, PrimFunc(symdeclare), symos, tmp19759)


_ = tmp19760

tmp19761 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19762 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19761)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19761
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19763 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp19762)
}
__typedArg0 := symsymbol
__typedArg1 := tmp19762
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19764 := Call(__e, PrimFunc(symdeclare), sympackage_2, tmp19763)


_ = tmp19764

tmp19765 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19766 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19765)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19765
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19767 := Call(__e, PrimFunc(symdeclare), symport, tmp19766)


_ = tmp19767

tmp19768 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19769 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19768)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19768
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19770 := Call(__e, PrimFunc(symdeclare), symporters, tmp19769)


_ = tmp19770

tmp19771 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19772 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19771)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19771
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19773 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp19772)
}
__typedArg0 := symnumber
__typedArg1 := tmp19772
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19774 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19773, Nil)
}
__typedArg0 := tmp19773
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19775 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19774)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19774
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19776 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp19775)
}
__typedArg0 := symstring
__typedArg1 := tmp19775
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19777 := Call(__e, PrimFunc(symdeclare), sympos, tmp19776)


_ = tmp19777

tmp19778 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symout, Nil)
}
__typedArg0 := symout
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19779 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp19778)
}
__typedArg0 := symstream
__typedArg1 := tmp19778
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19780 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19781 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19780)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19780
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19782 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19779, tmp19781)
}
__typedArg0 := tmp19779
__typedArg1 := tmp19781
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19783 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19782, Nil)
}
__typedArg0 := tmp19782
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19784 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19783)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19783
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19785 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp19784)
}
__typedArg0 := symstring
__typedArg1 := tmp19784
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19786 := Call(__e, PrimFunc(symdeclare), sympr, tmp19785)


_ = tmp19786

tmp19787 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19788 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19787)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19787
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19789 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp19788)
}
__typedArg0 := symA
__typedArg1 := tmp19788
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19790 := Call(__e, PrimFunc(symdeclare), symprint, tmp19789)


_ = tmp19790

tmp19791 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19792 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19791)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19791
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19793 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp19792)
}
__typedArg0 := symsymbol
__typedArg1 := tmp19792
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19794 := Call(__e, PrimFunc(symdeclare), symprofile, tmp19793)


_ = tmp19794

tmp19795 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19796 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19795)
}
__typedArg0 := symlist
__typedArg1 := tmp19795
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19797 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19798 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19797)
}
__typedArg0 := symlist
__typedArg1 := tmp19797
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19799 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19798, Nil)
}
__typedArg0 := tmp19798
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19800 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19799)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19799
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19801 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19796, tmp19800)
}
__typedArg0 := tmp19796
__typedArg1 := tmp19800
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19802 := Call(__e, PrimFunc(symdeclare), sympreclude, tmp19801)


_ = tmp19802

tmp19803 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19804 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19803)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19803
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19805 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp19804)
}
__typedArg0 := symstring
__typedArg1 := tmp19804
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19806 := Call(__e, PrimFunc(symdeclare), symshen_4proc_1nl, tmp19805)


_ = tmp19806

tmp19807 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19808 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_d, tmp19807)
}
__typedArg0 := sym_d
__typedArg1 := tmp19807
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19809 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp19808)
}
__typedArg0 := symsymbol
__typedArg1 := tmp19808
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19810 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19809, Nil)
}
__typedArg0 := tmp19809
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19811 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19810)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19810
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19812 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp19811)
}
__typedArg0 := symsymbol
__typedArg1 := tmp19811
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19813 := Call(__e, PrimFunc(symdeclare), symprofile_1results, tmp19812)


_ = tmp19813

tmp19814 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19815 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19814)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19814
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19816 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp19815)
}
__typedArg0 := symA
__typedArg1 := tmp19815
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19817 := Call(__e, PrimFunc(symdeclare), symprotect, tmp19816)


_ = tmp19817

tmp19818 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19819 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19818)
}
__typedArg0 := symlist
__typedArg1 := tmp19818
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19820 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19821 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19820)
}
__typedArg0 := symlist
__typedArg1 := tmp19820
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19822 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19821, Nil)
}
__typedArg0 := tmp19821
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19823 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19822)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19822
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19824 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19819, tmp19823)
}
__typedArg0 := tmp19819
__typedArg1 := tmp19823
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19825 := Call(__e, PrimFunc(symdeclare), sympreclude_1all_1but, tmp19824)


_ = tmp19825

tmp19826 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symout, Nil)
}
__typedArg0 := symout
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19827 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp19826)
}
__typedArg0 := symstream
__typedArg1 := tmp19826
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19828 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19829 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19828)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19828
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19830 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19827, tmp19829)
}
__typedArg0 := tmp19827
__typedArg1 := tmp19829
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19831 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19830, Nil)
}
__typedArg0 := tmp19830
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19832 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19831)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19831
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19833 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp19832)
}
__typedArg0 := symstring
__typedArg1 := tmp19832
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19834 := Call(__e, PrimFunc(symdeclare), symshen_4prhush, tmp19833)


_ = tmp19834

tmp19835 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19836 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19835)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19835
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19837 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp19836)
}
__typedArg0 := symnumber
__typedArg1 := tmp19836
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19838 := Call(__e, PrimFunc(symdeclare), symprolog_1memory, tmp19837)


_ = tmp19838

tmp19839 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunit, Nil)
}
__typedArg0 := symunit
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19840 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19839)
}
__typedArg0 := symlist
__typedArg1 := tmp19839
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19841 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19840, Nil)
}
__typedArg0 := tmp19840
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19842 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19841)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19841
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19843 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp19842)
}
__typedArg0 := symsymbol
__typedArg1 := tmp19842
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19844 := Call(__e, PrimFunc(symdeclare), symps, tmp19843)


_ = tmp19844

tmp19845 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symin, Nil)
}
__typedArg0 := symin
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19846 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp19845)
}
__typedArg0 := symstream
__typedArg1 := tmp19845
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19847 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunit, Nil)
}
__typedArg0 := symunit
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19848 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19847)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19847
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19849 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19846, tmp19848)
}
__typedArg0 := tmp19846
__typedArg1 := tmp19848
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19850 := Call(__e, PrimFunc(symdeclare), symread, tmp19849)


_ = tmp19850

tmp19851 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symin, Nil)
}
__typedArg0 := symin
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19852 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp19851)
}
__typedArg0 := symstream
__typedArg1 := tmp19851
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19853 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19854 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19853)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19853
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19855 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19852, tmp19854)
}
__typedArg0 := tmp19852
__typedArg1 := tmp19854
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19856 := Call(__e, PrimFunc(symdeclare), symread_1byte, tmp19855)


_ = tmp19856

tmp19857 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19858 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19857)
}
__typedArg0 := symlist
__typedArg1 := tmp19857
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19859 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19858, Nil)
}
__typedArg0 := tmp19858
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19860 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19859)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19859
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19861 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp19860)
}
__typedArg0 := symstring
__typedArg1 := tmp19860
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19862 := Call(__e, PrimFunc(symdeclare), symread_1file_1as_1bytelist, tmp19861)


_ = tmp19862

tmp19863 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19864 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19863)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19863
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19865 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp19864)
}
__typedArg0 := symstring
__typedArg1 := tmp19864
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19866 := Call(__e, PrimFunc(symdeclare), symread_1file_1as_1string, tmp19865)


_ = tmp19866

tmp19867 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunit, Nil)
}
__typedArg0 := symunit
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19868 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19867)
}
__typedArg0 := symlist
__typedArg1 := tmp19867
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19869 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19868, Nil)
}
__typedArg0 := tmp19868
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19870 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19869)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19869
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19871 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp19870)
}
__typedArg0 := symstring
__typedArg1 := tmp19870
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19872 := Call(__e, PrimFunc(symdeclare), symread_1file, tmp19871)


_ = tmp19872

tmp19873 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunit, Nil)
}
__typedArg0 := symunit
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19874 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19873)
}
__typedArg0 := symlist
__typedArg1 := tmp19873
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19875 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19874, Nil)
}
__typedArg0 := tmp19874
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19876 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19875)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19875
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19877 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp19876)
}
__typedArg0 := symstring
__typedArg1 := tmp19876
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19878 := Call(__e, PrimFunc(symdeclare), symread_1from_1string, tmp19877)


_ = tmp19878

tmp19879 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunit, Nil)
}
__typedArg0 := symunit
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19880 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19879)
}
__typedArg0 := symlist
__typedArg1 := tmp19879
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19881 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19880, Nil)
}
__typedArg0 := tmp19880
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19882 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19881)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19881
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19883 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp19882)
}
__typedArg0 := symstring
__typedArg1 := tmp19882
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19884 := Call(__e, PrimFunc(symdeclare), symread_1from_1string_1unprocessed, tmp19883)


_ = tmp19884

tmp19885 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19886 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19885)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19885
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19887 := Call(__e, PrimFunc(symdeclare), symrelease, tmp19886)


_ = tmp19887

tmp19888 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19889 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19888)
}
__typedArg0 := symlist
__typedArg1 := tmp19888
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19890 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19891 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19890)
}
__typedArg0 := symlist
__typedArg1 := tmp19890
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19892 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19891, Nil)
}
__typedArg0 := tmp19891
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19893 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19892)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19892
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19894 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19889, tmp19893)
}
__typedArg0 := tmp19889
__typedArg1 := tmp19893
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19895 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19894, Nil)
}
__typedArg0 := tmp19894
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19896 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19895)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19895
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19897 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp19896)
}
__typedArg0 := symA
__typedArg1 := tmp19896
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19898 := Call(__e, PrimFunc(symdeclare), symremove, tmp19897)


_ = tmp19898

tmp19899 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19900 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19899)
}
__typedArg0 := symlist
__typedArg1 := tmp19899
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19901 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19902 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19901)
}
__typedArg0 := symlist
__typedArg1 := tmp19901
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19903 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19902, Nil)
}
__typedArg0 := tmp19902
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19904 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19903)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19903
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19905 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19900, tmp19904)
}
__typedArg0 := tmp19900
__typedArg1 := tmp19904
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19906 := Call(__e, PrimFunc(symdeclare), symreverse, tmp19905)


_ = tmp19906

tmp19907 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19908 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19907)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19907
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19909 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp19908)
}
__typedArg0 := symstring
__typedArg1 := tmp19908
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19910 := Call(__e, PrimFunc(symdeclare), symsimple_1error, tmp19909)


_ = tmp19910

tmp19911 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symB, Nil)
}
__typedArg0 := symB
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19912 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_d, tmp19911)
}
__typedArg0 := sym_d
__typedArg1 := tmp19911
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19913 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp19912)
}
__typedArg0 := symA
__typedArg1 := tmp19912
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19914 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symB, Nil)
}
__typedArg0 := symB
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19915 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19914)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19914
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19916 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19913, tmp19915)
}
__typedArg0 := tmp19913
__typedArg1 := tmp19915
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19917 := Call(__e, PrimFunc(symdeclare), symsnd, tmp19916)


_ = tmp19917

tmp19918 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19919 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19918)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19918
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19920 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp19919)
}
__typedArg0 := symnumber
__typedArg1 := tmp19919
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19921 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19920, Nil)
}
__typedArg0 := tmp19920
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19922 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19921)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19921
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19923 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp19922)
}
__typedArg0 := symsymbol
__typedArg1 := tmp19922
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19924 := Call(__e, PrimFunc(symdeclare), symspecialise, tmp19923)


_ = tmp19924

tmp19925 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19926 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19925)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19925
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19927 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp19926)
}
__typedArg0 := symsymbol
__typedArg1 := tmp19926
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19928 := Call(__e, PrimFunc(symdeclare), symspy, tmp19927)


_ = tmp19928

tmp19929 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19930 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19929)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19929
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19931 := Call(__e, PrimFunc(symdeclare), symspy_2, tmp19930)


_ = tmp19931

tmp19932 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19933 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19932)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19932
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19934 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp19933)
}
__typedArg0 := symsymbol
__typedArg1 := tmp19933
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19935 := Call(__e, PrimFunc(symdeclare), symstep, tmp19934)


_ = tmp19935

tmp19936 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19937 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19936)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19936
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19938 := Call(__e, PrimFunc(symdeclare), symstep_2, tmp19937)


_ = tmp19938

tmp19939 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symin, Nil)
}
__typedArg0 := symin
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19940 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp19939)
}
__typedArg0 := symstream
__typedArg1 := tmp19939
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19941 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19940, Nil)
}
__typedArg0 := tmp19940
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19942 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19941)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19941
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19943 := Call(__e, PrimFunc(symdeclare), symstinput, tmp19942)


_ = tmp19943

tmp19944 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symout, Nil)
}
__typedArg0 := symout
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19945 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp19944)
}
__typedArg0 := symstream
__typedArg1 := tmp19944
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19946 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19945, Nil)
}
__typedArg0 := tmp19945
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19947 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19946)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19946
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19948 := Call(__e, PrimFunc(symdeclare), symstoutput, tmp19947)


_ = tmp19948

tmp19949 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19950 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19949)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19949
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19951 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp19950)
}
__typedArg0 := symA
__typedArg1 := tmp19950
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19952 := Call(__e, PrimFunc(symdeclare), symstring_2, tmp19951)


_ = tmp19952

tmp19953 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19954 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19953)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19953
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19955 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp19954)
}
__typedArg0 := symA
__typedArg1 := tmp19954
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19956 := Call(__e, PrimFunc(symdeclare), symstr, tmp19955)


_ = tmp19956

tmp19957 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19958 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19957)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19957
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19959 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp19958)
}
__typedArg0 := symstring
__typedArg1 := tmp19958
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19960 := Call(__e, PrimFunc(symdeclare), symstring_1_6n, tmp19959)


_ = tmp19960

tmp19961 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19962 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19961)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19961
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19963 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp19962)
}
__typedArg0 := symstring
__typedArg1 := tmp19962
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19964 := Call(__e, PrimFunc(symdeclare), symstring_1_6symbol, tmp19963)


_ = tmp19964

tmp19965 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19966 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19965)
}
__typedArg0 := symlist
__typedArg1 := tmp19965
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19967 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19968 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19967)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19967
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19969 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19966, tmp19968)
}
__typedArg0 := tmp19966
__typedArg1 := tmp19968
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19970 := Call(__e, PrimFunc(symdeclare), symsum, tmp19969)


_ = tmp19970

tmp19971 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19972 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19971)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19971
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19973 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp19972)
}
__typedArg0 := symA
__typedArg1 := tmp19972
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19974 := Call(__e, PrimFunc(symdeclare), symsymbol_2, tmp19973)


_ = tmp19974

tmp19975 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19976 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19975)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19975
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19977 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp19976)
}
__typedArg0 := symsymbol
__typedArg1 := tmp19976
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19978 := Call(__e, PrimFunc(symdeclare), symsystemf, tmp19977)


_ = tmp19978

tmp19979 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19980 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19979)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19979
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19981 := Call(__e, PrimFunc(symdeclare), symsystem_1S_2, tmp19980)


_ = tmp19981

tmp19982 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19983 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19982)
}
__typedArg0 := symlist
__typedArg1 := tmp19982
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19984 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19985 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp19984)
}
__typedArg0 := symlist
__typedArg1 := tmp19984
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19986 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19985, Nil)
}
__typedArg0 := tmp19985
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19987 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19986)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19986
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19988 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19983, tmp19987)
}
__typedArg0 := tmp19983
__typedArg1 := tmp19987
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19989 := Call(__e, PrimFunc(symdeclare), symtail, tmp19988)


_ = tmp19989

tmp19990 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19991 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19990)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19990
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19992 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp19991)
}
__typedArg0 := symstring
__typedArg1 := tmp19991
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19993 := Call(__e, PrimFunc(symdeclare), symtlstr, tmp19992)


_ = tmp19993

tmp19994 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19995 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp19994)
}
__typedArg0 := symvector
__typedArg1 := tmp19994
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19996 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19997 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp19996)
}
__typedArg0 := symvector
__typedArg1 := tmp19996
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19998 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19997, Nil)
}
__typedArg0 := tmp19997
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19999 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp19998)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp19998
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20000 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19995, tmp19999)
}
__typedArg0 := tmp19995
__typedArg1 := tmp19999
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20001 := Call(__e, PrimFunc(symdeclare), symtlv, tmp20000)


_ = tmp20001

tmp20002 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20003 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20002)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20002
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20004 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp20003)
}
__typedArg0 := symsymbol
__typedArg1 := tmp20003
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20005 := Call(__e, PrimFunc(symdeclare), symtc, tmp20004)


_ = tmp20005

tmp20006 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20007 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20006)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20006
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20008 := Call(__e, PrimFunc(symdeclare), symtc_2, tmp20007)


_ = tmp20008

tmp20009 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20010 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlazy, tmp20009)
}
__typedArg0 := symlazy
__typedArg1 := tmp20009
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20011 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20012 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20011)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20011
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20013 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20010, tmp20012)
}
__typedArg0 := tmp20010
__typedArg1 := tmp20012
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20014 := Call(__e, PrimFunc(symdeclare), symthaw, tmp20013)


_ = tmp20014

tmp20015 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20016 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20015)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20015
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20017 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp20016)
}
__typedArg0 := symsymbol
__typedArg1 := tmp20016
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20018 := Call(__e, PrimFunc(symdeclare), symtrack, tmp20017)


_ = tmp20018

tmp20019 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20020 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20019)
}
__typedArg0 := symlist
__typedArg1 := tmp20019
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20021 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20020, Nil)
}
__typedArg0 := tmp20020
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20022 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20021)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20021
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20023 := Call(__e, PrimFunc(symdeclare), symtracked, tmp20022)


_ = tmp20023

tmp20024 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20025 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20024)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20024
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20026 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symexception, tmp20025)
}
__typedArg0 := symexception
__typedArg1 := tmp20025
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20027 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20028 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20027)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20027
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20029 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20026, tmp20028)
}
__typedArg0 := tmp20026
__typedArg1 := tmp20028
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20030 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20029, Nil)
}
__typedArg0 := tmp20029
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20031 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20030)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20030
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20032 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp20031)
}
__typedArg0 := symA
__typedArg1 := tmp20031
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20033 := Call(__e, PrimFunc(symdeclare), symtrap_1error, tmp20032)


_ = tmp20033

tmp20034 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20035 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20034)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20034
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20036 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp20035)
}
__typedArg0 := symA
__typedArg1 := tmp20035
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20037 := Call(__e, PrimFunc(symdeclare), symtuple_2, tmp20036)


_ = tmp20037

tmp20038 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20039 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20038)
}
__typedArg0 := symlist
__typedArg1 := tmp20038
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20040 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20039, Nil)
}
__typedArg0 := tmp20039
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20041 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20040)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20040
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20042 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp20041)
}
__typedArg0 := symstring
__typedArg1 := tmp20041
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20043 := Call(__e, PrimFunc(symdeclare), symunabsolute, tmp20042)


_ = tmp20043

tmp20044 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20045 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20044)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20044
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20046 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp20045)
}
__typedArg0 := symsymbol
__typedArg1 := tmp20045
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20047 := Call(__e, PrimFunc(symdeclare), symundefmacro, tmp20046)


_ = tmp20047

tmp20048 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20049 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20048)
}
__typedArg0 := symlist
__typedArg1 := tmp20048
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20050 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20051 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20050)
}
__typedArg0 := symlist
__typedArg1 := tmp20050
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20052 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20053 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20052)
}
__typedArg0 := symlist
__typedArg1 := tmp20052
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20054 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20053, Nil)
}
__typedArg0 := tmp20053
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20055 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20054)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20054
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20056 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20051, tmp20055)
}
__typedArg0 := tmp20051
__typedArg1 := tmp20055
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20057 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20056, Nil)
}
__typedArg0 := tmp20056
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20058 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20057)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20057
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20059 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20049, tmp20058)
}
__typedArg0 := tmp20049
__typedArg1 := tmp20058
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20060 := Call(__e, PrimFunc(symdeclare), symunion, tmp20059)


_ = tmp20060

tmp20061 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20062 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20061)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20061
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20063 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp20062)
}
__typedArg0 := symsymbol
__typedArg1 := tmp20062
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20064 := Call(__e, PrimFunc(symdeclare), symunprofile, tmp20063)


_ = tmp20064

tmp20065 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20066 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20065)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20065
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20067 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, tmp20066)
}
__typedArg0 := symsymbol
__typedArg1 := tmp20066
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20068 := Call(__e, PrimFunc(symdeclare), symuntrack, tmp20067)


_ = tmp20068

tmp20069 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsymbol, Nil)
}
__typedArg0 := symsymbol
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20070 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp20069)
}
__typedArg0 := symlist
__typedArg1 := tmp20069
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20071 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20070, Nil)
}
__typedArg0 := tmp20070
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20072 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20071)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20071
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20073 := Call(__e, PrimFunc(symdeclare), symuserdefs, tmp20072)


_ = tmp20073

tmp20074 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20075 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20074)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20074
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20076 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp20075)
}
__typedArg0 := symA
__typedArg1 := tmp20075
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20077 := Call(__e, PrimFunc(symdeclare), symvariable_2, tmp20076)


_ = tmp20077

tmp20078 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20079 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20078)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20078
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20080 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp20079)
}
__typedArg0 := symA
__typedArg1 := tmp20079
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20081 := Call(__e, PrimFunc(symdeclare), symvector_2, tmp20080)


_ = tmp20081

tmp20082 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20083 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20082)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20082
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20084 := Call(__e, PrimFunc(symdeclare), symversion, tmp20083)


_ = tmp20084

tmp20085 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, Nil)
}
__typedArg0 := symA
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20086 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20085)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20085
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20087 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp20086)
}
__typedArg0 := symA
__typedArg1 := tmp20086
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20088 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20087, Nil)
}
__typedArg0 := tmp20087
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20089 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20088)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20088
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20090 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp20089)
}
__typedArg0 := symstring
__typedArg1 := tmp20089
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20091 := Call(__e, PrimFunc(symdeclare), symwrite_1to_1file, tmp20090)


_ = tmp20091

tmp20092 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symout, Nil)
}
__typedArg0 := symout
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20093 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp20092)
}
__typedArg0 := symstream
__typedArg1 := tmp20092
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20094 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20095 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20094)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20094
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20096 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20093, tmp20095)
}
__typedArg0 := tmp20093
__typedArg1 := tmp20095
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20097 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20096, Nil)
}
__typedArg0 := tmp20096
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20098 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20097)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20097
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20099 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp20098)
}
__typedArg0 := symnumber
__typedArg1 := tmp20098
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20100 := Call(__e, PrimFunc(symdeclare), symwrite_1byte, tmp20099)


_ = tmp20100

tmp20101 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20102 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20101)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20101
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20103 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, tmp20102)
}
__typedArg0 := symstring
__typedArg1 := tmp20102
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20104 := Call(__e, PrimFunc(symdeclare), symy_1or_1n_2, tmp20103)


_ = tmp20104

tmp20105 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20106 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20105)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20105
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20107 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp20106)
}
__typedArg0 := symnumber
__typedArg1 := tmp20106
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20108 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20107, Nil)
}
__typedArg0 := tmp20107
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20109 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20108)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20108
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20110 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp20109)
}
__typedArg0 := symnumber
__typedArg1 := tmp20109
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20111 := Call(__e, PrimFunc(symdeclare), sym_6, tmp20110)


_ = tmp20111

tmp20112 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20113 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20112)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20112
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20114 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp20113)
}
__typedArg0 := symnumber
__typedArg1 := tmp20113
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20115 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20114, Nil)
}
__typedArg0 := tmp20114
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20116 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20115)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20115
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20117 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp20116)
}
__typedArg0 := symnumber
__typedArg1 := tmp20116
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20118 := Call(__e, PrimFunc(symdeclare), sym_5, tmp20117)


_ = tmp20118

tmp20119 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20120 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20119)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20119
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20121 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp20120)
}
__typedArg0 := symnumber
__typedArg1 := tmp20120
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20122 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20121, Nil)
}
__typedArg0 := tmp20121
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20123 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20122)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20122
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20124 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp20123)
}
__typedArg0 := symnumber
__typedArg1 := tmp20123
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20125 := Call(__e, PrimFunc(symdeclare), sym_6_a, tmp20124)


_ = tmp20125

tmp20126 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20127 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20126)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20126
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20128 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp20127)
}
__typedArg0 := symnumber
__typedArg1 := tmp20127
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20129 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20128, Nil)
}
__typedArg0 := tmp20128
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20130 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20129)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20129
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20131 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp20130)
}
__typedArg0 := symnumber
__typedArg1 := tmp20130
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20132 := Call(__e, PrimFunc(symdeclare), sym_5_a, tmp20131)


_ = tmp20132

tmp20133 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20134 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20133)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20133
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20135 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp20134)
}
__typedArg0 := symA
__typedArg1 := tmp20134
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20136 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20135, Nil)
}
__typedArg0 := tmp20135
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20137 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20136)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20136
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20138 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp20137)
}
__typedArg0 := symA
__typedArg1 := tmp20137
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20139 := Call(__e, PrimFunc(symdeclare), sym_a, tmp20138)


_ = tmp20139

tmp20140 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20141 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20140)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20140
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20142 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp20141)
}
__typedArg0 := symnumber
__typedArg1 := tmp20141
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20143 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20142, Nil)
}
__typedArg0 := tmp20142
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20144 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20143)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20143
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20145 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp20144)
}
__typedArg0 := symnumber
__typedArg1 := tmp20144
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20146 := Call(__e, PrimFunc(symdeclare), sym_7, tmp20145)


_ = tmp20146

tmp20147 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20148 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20147)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20147
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20149 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp20148)
}
__typedArg0 := symnumber
__typedArg1 := tmp20148
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20150 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20149, Nil)
}
__typedArg0 := tmp20149
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20151 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20150)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20150
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20152 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp20151)
}
__typedArg0 := symnumber
__typedArg1 := tmp20151
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20153 := Call(__e, PrimFunc(symdeclare), sym_c, tmp20152)


_ = tmp20153

tmp20154 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20155 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20154)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20154
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20156 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp20155)
}
__typedArg0 := symnumber
__typedArg1 := tmp20155
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20157 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20156, Nil)
}
__typedArg0 := tmp20156
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20158 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20157)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20157
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20159 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp20158)
}
__typedArg0 := symnumber
__typedArg1 := tmp20158
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20160 := Call(__e, PrimFunc(symdeclare), sym_1, tmp20159)


_ = tmp20160

tmp20161 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, Nil)
}
__typedArg0 := symnumber
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20162 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20161)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20161
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20163 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp20162)
}
__typedArg0 := symnumber
__typedArg1 := tmp20162
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20164 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20163, Nil)
}
__typedArg0 := tmp20163
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20165 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20164)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20164
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20166 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnumber, tmp20165)
}
__typedArg0 := symnumber
__typedArg1 := tmp20165
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20167 := Call(__e, PrimFunc(symdeclare), sym_d, tmp20166)


_ = tmp20167

tmp20168 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symboolean, Nil)
}
__typedArg0 := symboolean
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20169 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20168)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20168
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20170 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symB, tmp20169)
}
__typedArg0 := symB
__typedArg1 := tmp20169
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20171 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20170, Nil)
}
__typedArg0 := tmp20170
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20172 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp20171)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp20171
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20173 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp20172)
}
__typedArg0 := symA
__typedArg1 := tmp20172
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symdeclare), sym_a_a, tmp20173)
return




}, 0)

