package main

import . "github.com/pyrex41/shen-go/kl"

var PrologMain = MakeNative(func(__e *ControlFlow) {
tmp10132 := MakeNative(func(__e *ControlFlow) {
V1599 := __e.Get(1)
_ = V1599
__e.TailApply(PrimFunc(symshen_4assert_d), V1599, symshen_4top)
return
}, 1)

tmp10133 := Call(__e, ns2_1set, symasserta, tmp10132)


_ = tmp10133

tmp10134 := MakeNative(func(__e *ControlFlow) {
V1600 := __e.Get(1)
_ = V1600
__e.TailApply(PrimFunc(symshen_4assert_d), V1600, symshen_4bottom)
return
}, 1)

tmp10135 := Call(__e, ns2_1set, symassertz, tmp10134)


_ = tmp10135

tmp10136 := MakeNative(func(__e *ControlFlow) {
V1601 := __e.Get(1)
_ = V1601
V1602 := __e.Get(2)
_ = V1602
tmp10163 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1601)
}
__typedArg0 := V1601
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10154 Obj

if True == tmp10163 {
tmp10161 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1601)
}
__typedArg0 := V1601
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10162 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10161)
}
__typedArg0 := tmp10161
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10156 Obj

if True == tmp10162 {
tmp10158 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1601)
}
__typedArg0 := V1601
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10159 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10158)
}
__typedArg0 := tmp10158
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10160 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_5_1_1, tmp10159)
}
__typedArg0 := sym_5_1_1
__typedArg1 := tmp10159
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10157 Obj

if True == tmp10160 {
ifres10157 = True


} else {
ifres10157 = False


}

ifres10156 = ifres10157


} else {
ifres10156 = False


}

var ifres10155 Obj

if True == ifres10156 {
ifres10155 = True


} else {
ifres10155 = False


}

ifres10154 = ifres10155


} else {
ifres10154 = False


}

if True == ifres10154 {
tmp10137 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1601)
}
__typedArg0 := V1601
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10138 := Call(__e, PrimFunc(symshen_4predicate), tmp10137)


let__9914 := tmp10138
_ = let__9914

tmp10139 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1601)
}
__typedArg0 := V1601
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10140 := Call(__e, PrimFunc(symshen_4terms), tmp10139)


let__9915 := tmp10140
_ = let__9915

tmp10141 := Call(__e, PrimFunc(symlength), let__9915)


let__9916 := tmp10141
_ = let__9916

tmp10142 := Call(__e, PrimFunc(symshen_4parameters), let__9916)


let__9917 := tmp10142
_ = let__9917

tmp10143 := Call(__e, PrimFunc(symarity), let__9914)


let__9918 := tmp10143
_ = let__9918

tmp10149 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__9918, MakeInteger(-1))
}
__typedArg0 := let__9918
__typedArg1 := MakeInteger(-1)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10144 Obj

if True == tmp10149 {
tmp10145 := Call(__e, PrimFunc(symshen_4create_1skeleton), let__9914, let__9917)


tmp10146 := Call(__e, PrimFunc(symeval), tmp10145)


_ = tmp10146

tmp10147 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp10148 := Call(__e, PrimFunc(symput), let__9914, symshen_4dynamic, Nil, tmp10147)


ifres10144 = tmp10148


} else {
ifres10144 = symshen_4skip


}

let__9919 := ifres10144
_ = let__9919

tmp10150 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1601)
}
__typedArg0 := V1601
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10151 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10150)
}
__typedArg0 := tmp10150
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10152 := Call(__e, PrimFunc(symshen_4insert_1info), let__9914, let__9915, tmp10151, V1601, V1602)


let__9920 := tmp10152
_ = let__9920

__e.Return(let__9914)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4assert_d)
return
}


}, 2)

tmp10164 := Call(__e, ns2_1set, symshen_4assert_d, tmp10136)


_ = tmp10164

tmp10165 := MakeNative(func(__e *ControlFlow) {
V1612 := __e.Get(1)
_ = V1612
tmp10167 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1612)
}
__typedArg0 := V1612
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp10167 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1612)
}
__typedArg0 := V1612
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return
} else {
__e.Return(V1612)
return
}


}, 1)

tmp10168 := Call(__e, ns2_1set, symshen_4predicate, tmp10165)


_ = tmp10168

tmp10169 := MakeNative(func(__e *ControlFlow) {
V1617 := __e.Get(1)
_ = V1617
tmp10171 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1617)
}
__typedArg0 := V1617
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp10171 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1617)
}
__typedArg0 := V1617
return Call(__e, PrimFunc(symtl), __typedArg0)
})())
return
} else {
__e.Return(Nil)
return
}


}, 1)

tmp10172 := Call(__e, ns2_1set, symshen_4terms, tmp10169)


_ = tmp10172

tmp10173 := MakeNative(func(__e *ControlFlow) {
V1618 := __e.Get(1)
_ = V1618
V1619 := __e.Get(2)
_ = V1619
tmp10174 := Call(__e, PrimFunc(symshen_4dynamic_1default), V1618, V1619)


tmp10175 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1618, tmp10174)
}
__typedArg0 := V1618
__typedArg1 := tmp10174
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefprolog, tmp10175)
}
__typedArg0 := symdefprolog
__typedArg1 := tmp10175
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 2)

tmp10176 := Call(__e, ns2_1set, symshen_4create_1skeleton, tmp10173)


_ = tmp10176

tmp10177 := MakeNative(func(__e *ControlFlow) {
V1620 := __e.Get(1)
_ = V1620
V1621 := __e.Get(2)
_ = V1621
tmp10178 := Call(__e, PrimFunc(symshen_4cons_1form), V1621)


tmp10179 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4dynamic, Nil)
}
__typedArg0 := symshen_4dynamic
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10180 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1620, tmp10179)
}
__typedArg0 := V1620
__typedArg1 := tmp10179
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10181 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symget, tmp10180)
}
__typedArg0 := symget
__typedArg1 := tmp10180
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10182 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp10181, Nil)
}
__typedArg0 := tmp10181
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10183 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp10178, tmp10182)
}
__typedArg0 := tmp10178
__typedArg1 := tmp10182
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10184 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4call_1dynamic, tmp10183)
}
__typedArg0 := symshen_4call_1dynamic
__typedArg1 := tmp10183
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10185 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(";"))
}
__typedArg0 := MakeString(";")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp10186 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp10185, Nil)
}
__typedArg0 := tmp10185
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10187 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp10184, tmp10186)
}
__typedArg0 := tmp10184
__typedArg1 := tmp10186
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10188 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1_1, tmp10187)
}
__typedArg0 := sym_5_1_1
__typedArg1 := tmp10187
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symappend), V1621, tmp10188)
return


}, 2)

tmp10189 := Call(__e, ns2_1set, symshen_4dynamic_1default, tmp10177)


_ = tmp10189

tmp10190 := MakeNative(func(__e *ControlFlow) {
V1622 := __e.Get(1)
_ = V1622
V1623 := __e.Get(2)
_ = V1623
V1624 := __e.Get(3)
_ = V1624
V1625 := __e.Get(4)
_ = V1625
V1626 := __e.Get(5)
_ = V1626
tmp10191 := Call(__e, PrimFunc(symgensym), symg)


let__9921 := tmp10191
_ = let__9921

tmp10192 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__9921, Nil)
}
__typedArg0 := let__9921
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10193 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefprolog, tmp10192)
}
__typedArg0 := symdefprolog
__typedArg1 := tmp10192
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10194 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1_1, V1624)
}
__typedArg0 := sym_5_1_1
__typedArg1 := V1624
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10195 := Call(__e, PrimFunc(symappend), V1623, tmp10194)


tmp10196 := Call(__e, PrimFunc(symappend), tmp10193, tmp10195)


tmp10197 := Call(__e, PrimFunc(symeval), tmp10196)


let__9922 := tmp10197
_ = let__9922

tmp10198 := Call(__e, PrimFunc(symfn), let__9921)


tmp10199 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__9921, V1625)
}
__typedArg0 := let__9921
__typedArg1 := V1625
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10200 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp10198, tmp10199)
}
__typedArg0 := tmp10198
__typedArg1 := tmp10199
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__9923 := tmp10200
_ = let__9923

tmp10201 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp10202 := Call(__e, PrimFunc(symget), V1622, symshen_4dynamic, tmp10201)


let__9924 := tmp10202
_ = let__9924

tmp10207 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V1626, symshen_4top)
}
__typedArg0 := V1626
__typedArg1 := symshen_4top
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10203 Obj

if True == tmp10207 {
tmp10204 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__9923, let__9924)
}
__typedArg0 := let__9923
__typedArg1 := let__9924
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

ifres10203 = tmp10204


} else {
tmp10205 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__9923, Nil)
}
__typedArg0 := let__9923
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10206 := Call(__e, PrimFunc(symappend), let__9924, tmp10205)


ifres10203 = tmp10206


}

let__9925 := ifres10203
_ = let__9925

tmp10208 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symput), V1622, symshen_4dynamic, let__9925, tmp10208)
return


}, 5)

tmp10209 := Call(__e, ns2_1set, symshen_4insert_1info, tmp10190)


_ = tmp10209

tmp10210 := MakeNative(func(__e *ControlFlow) {
tmp10211 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dnames_d)
}
__typedArg0 := symshen_4_dnames_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

let__9926 := tmp10211
_ = let__9926

tmp10217 := Call(__e, PrimFunc(symempty_2), let__9926)


var ifres10212 Obj

if True == tmp10217 {
tmp10213 := Call(__e, PrimFunc(symgensym), symg)


ifres10212 = tmp10213


} else {
tmp10214 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__9926)
}
__typedArg0 := let__9926
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10215 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dnames_d, tmp10214)
}
__typedArg0 := symshen_4_dnames_d
__typedArg1 := tmp10214
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp10215

tmp10216 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__9926)
}
__typedArg0 := let__9926
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

ifres10212 = tmp10216


}

let__9927 := ifres10212
_ = let__9927

__e.Return(let__9927)
return


}, 0)

tmp10218 := Call(__e, ns2_1set, symshen_4newname, tmp10210)


_ = tmp10218

tmp10219 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__self3 := __e.Get(3)
__self4 := __e.Get(4)
__self5 := __e.Get(5)
__self6 := __e.Get(6)
__selftop:
V1634 := __self1
_ = V1634
V1635 := __self2
_ = V1635
V1636 := __self3
_ = V1636
V1637 := __self4
_ = V1637
V1638 := __self5
_ = V1638
V1639 := __self6
_ = V1639
tmp10231 := Call(__e, PrimFunc(symshen_4unlocked_2), V1637)


var ifres10220 Obj

if True == tmp10231 {
tmp10221 := Call(__e, PrimFunc(symshen_4lazyderef), V1635, V1636)


let__9929 := tmp10221
_ = let__9929

tmp10230 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__9929)
}
__typedArg0 := let__9929
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10222 Obj

if True == tmp10230 {
tmp10223 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__9929)
}
__typedArg0 := let__9929
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10224 := Call(__e, PrimFunc(symshen_4lazyderef), tmp10223, V1636)


let__9930 := tmp10224
_ = let__9930

tmp10229 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__9930)
}
__typedArg0 := let__9930
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10225 Obj

if True == tmp10229 {
tmp10226 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__9930)
}
__typedArg0 := let__9930
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__9931 := tmp10226
_ = let__9931

tmp10227 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp10227

tmp10228 := Call(__e, PrimFunc(symshen_4callrec), let__9931, V1634, V1636, V1637, V1638, V1639)


ifres10225 = tmp10228


} else {
ifres10225 = False


}

ifres10222 = ifres10225


} else {
ifres10222 = False


}

ifres10220 = ifres10222


} else {
ifres10220 = False


}

let__9928 := ifres10220
_ = let__9928

tmp10240 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__9928, False)
}
__typedArg0 := let__9928
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp10240 {
tmp10238 := Call(__e, PrimFunc(symshen_4unlocked_2), V1637)


if True == tmp10238 {
tmp10232 := Call(__e, PrimFunc(symshen_4lazyderef), V1635, V1636)


let__9932 := tmp10232
_ = let__9932

tmp10236 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__9932)
}
__typedArg0 := let__9932
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp10236 {
tmp10233 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__9932)
}
__typedArg0 := let__9932
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__9933 := tmp10233
_ = let__9933

tmp10234 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp10234

if PrimFunc(symshen_4call_1dynamic) == __self {
__self1, __self2, __self3, __self4, __self5, __self6 = V1634, let__9933, V1636, V1637, V1638, V1639
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4call_1dynamic), V1634, let__9933, V1636, V1637, V1638, V1639)
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
__e.Return(let__9928)
return
}


}, 6)

tmp10241 := Call(__e, ns2_1set, symshen_4call_1dynamic, tmp10219)


_ = tmp10241

tmp10242 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__self3 := __e.Get(3)
__self4 := __e.Get(4)
__self5 := __e.Get(5)
__self6 := __e.Get(6)
__selftop:
V1646 := __self1
_ = V1646
V1647 := __self2
_ = V1647
V1648 := __self3
_ = V1648
V1649 := __self4
_ = V1649
V1650 := __self5
_ = V1650
V1651 := __self6
_ = V1651
tmp10252 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V1647)
}
__typedArg0 := Nil
__typedArg1 := V1647
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp10252 {
tmp10243 := Call(__e, V1646, V1648)


tmp10244 := Call(__e, tmp10243, V1649)


tmp10245 := Call(__e, tmp10244, V1650)


__e.TailApply(tmp10245, V1651)
return


} else {
tmp10250 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1647)
}
__typedArg0 := V1647
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp10250 {
tmp10246 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1647)
}
__typedArg0 := V1647
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10247 := Call(__e, V1646, tmp10246)


tmp10248 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1647)
}
__typedArg0 := V1647
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4callrec) == __self {
__self1, __self2, __self3, __self4, __self5, __self6 = tmp10247, tmp10248, V1648, V1649, V1650, V1651
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4callrec), tmp10247, tmp10248, V1648, V1649, V1650, V1651)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4callrec)
return
}


}


}, 6)

tmp10253 := Call(__e, ns2_1set, symshen_4callrec, tmp10242)


_ = tmp10253

tmp10254 := MakeNative(func(__e *ControlFlow) {
V1652 := __e.Get(1)
_ = V1652
tmp10271 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1652)
}
__typedArg0 := V1652
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10262 Obj

if True == tmp10271 {
tmp10269 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1652)
}
__typedArg0 := V1652
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10270 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10269)
}
__typedArg0 := tmp10269
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10264 Obj

if True == tmp10270 {
tmp10266 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1652)
}
__typedArg0 := V1652
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10267 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10266)
}
__typedArg0 := tmp10266
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10268 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_5_1_1, tmp10267)
}
__typedArg0 := sym_5_1_1
__typedArg1 := tmp10267
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10265 Obj

if True == tmp10268 {
ifres10265 = True


} else {
ifres10265 = False


}

ifres10264 = ifres10265


} else {
ifres10264 = False


}

var ifres10263 Obj

if True == ifres10264 {
ifres10263 = True


} else {
ifres10263 = False


}

ifres10262 = ifres10263


} else {
ifres10262 = False


}

if True == ifres10262 {
tmp10255 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1652)
}
__typedArg0 := V1652
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10256 := Call(__e, PrimFunc(symshen_4predicate), tmp10255)


let__9934 := tmp10256
_ = let__9934

tmp10257 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp10258 := Call(__e, PrimFunc(symget), let__9934, symshen_4dynamic, tmp10257)


let__9935 := tmp10258
_ = let__9935

tmp10259 := Call(__e, PrimFunc(symshen_4retract_1clause), V1652, let__9935)


tmp10260 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symput), let__9934, symshen_4dynamic, tmp10259, tmp10260)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symretract)
return
}


}, 1)

tmp10272 := Call(__e, ns2_1set, symretract, tmp10254)


_ = tmp10272

tmp10273 := MakeNative(func(__e *ControlFlow) {
V1660 := __e.Get(1)
_ = V1660
V1661 := __e.Get(2)
_ = V1661
tmp10303 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V1661)
}
__typedArg0 := Nil
__typedArg1 := V1661
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp10303 {
__e.Return(Nil)
return
} else {
tmp10301 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1661)
}
__typedArg0 := V1661
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10286 Obj

if True == tmp10301 {
tmp10299 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1661)
}
__typedArg0 := V1661
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10300 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10299)
}
__typedArg0 := tmp10299
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10288 Obj

if True == tmp10300 {
tmp10296 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1661)
}
__typedArg0 := V1661
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10297 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10296)
}
__typedArg0 := tmp10296
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10298 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10297)
}
__typedArg0 := tmp10297
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10290 Obj

if True == tmp10298 {
tmp10292 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1661)
}
__typedArg0 := V1661
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10293 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10292)
}
__typedArg0 := tmp10292
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10294 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10293)
}
__typedArg0 := tmp10293
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10295 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V1660, tmp10294)
}
__typedArg0 := V1660
__typedArg1 := tmp10294
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10291 Obj

if True == tmp10295 {
ifres10291 = True


} else {
ifres10291 = False


}

ifres10290 = ifres10291


} else {
ifres10290 = False


}

var ifres10289 Obj

if True == ifres10290 {
ifres10289 = True


} else {
ifres10289 = False


}

ifres10288 = ifres10289


} else {
ifres10288 = False


}

var ifres10287 Obj

if True == ifres10288 {
ifres10287 = True


} else {
ifres10287 = False


}

ifres10286 = ifres10287


} else {
ifres10286 = False


}

if True == ifres10286 {
tmp10274 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1661)
}
__typedArg0 := V1661
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10275 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10274)
}
__typedArg0 := tmp10274
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10276 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10275)
}
__typedArg0 := tmp10275
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10277 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dnames_d)
}
__typedArg0 := symshen_4_dnames_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp10278 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp10276, tmp10277)
}
__typedArg0 := tmp10276
__typedArg1 := tmp10277
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10279 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dnames_d, tmp10278)
}
__typedArg0 := symshen_4_dnames_d
__typedArg1 := tmp10278
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp10279

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1661)
}
__typedArg0 := V1661
return Call(__e, PrimFunc(symtl), __typedArg0)
})())
return


} else {
tmp10284 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1661)
}
__typedArg0 := V1661
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp10284 {
tmp10280 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1661)
}
__typedArg0 := V1661
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10281 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1661)
}
__typedArg0 := V1661
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10282 := Call(__e, PrimFunc(symshen_4retract_1clause), V1660, tmp10281)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp10280, tmp10282)
}
__typedArg0 := tmp10280
__typedArg1 := tmp10282
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4retract_1clause)
return
}


}


}


}, 2)

tmp10304 := Call(__e, ns2_1set, symshen_4retract_1clause, tmp10273)


_ = tmp10304

tmp10305 := MakeNative(func(__e *ControlFlow) {
V1662 := __e.Get(1)
_ = V1662
V1663 := __e.Get(2)
_ = V1663
tmp10306 := MakeNative(func(__e *ControlFlow) {
Z1664 := __e.Get(1)
_ = Z1664
__e.TailApply(PrimFunc(symshen_4_5defprolog_6), Z1664)
return
}, 1)

tmp10307 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1662, V1663)
}
__typedArg0 := V1662
__typedArg1 := V1663
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symcompile), tmp10306, tmp10307)
return


}, 2)

tmp10308 := Call(__e, ns2_1set, symshen_4compile_1prolog, tmp10305)


_ = tmp10308

tmp10309 := MakeNative(func(__e *ControlFlow) {
V1665 := __e.Get(1)
_ = V1665
tmp10325 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1665)
}
__typedArg0 := V1665
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10310 Obj

if True == tmp10325 {
tmp10311 := Call(__e, PrimFunc(symhead), V1665)


let__9937 := tmp10311
_ = let__9937

tmp10312 := Call(__e, PrimFunc(symtail), V1665)


let__9938 := tmp10312
_ = let__9938

tmp10313 := Call(__e, PrimFunc(symshen_4_5clauses_6), let__9938)


let__9939 := tmp10313
_ = let__9939

tmp10323 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9939)


var ifres10314 Obj

if True == tmp10323 {
tmp10315 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10314 = tmp10315


} else {
tmp10316 := Call(__e, PrimFunc(symshen_4_5_1out), let__9939)


let__9940 := tmp10316
_ = let__9940

tmp10317 := Call(__e, PrimFunc(symshen_4in_1_6), let__9939)


let__9941 := tmp10317
_ = let__9941

tmp10318 := Call(__e, PrimFunc(symshen_4prolog_1arity_1check), let__9937, let__9940)


let__9942 := tmp10318
_ = let__9942

tmp10319 := MakeNative(func(__e *ControlFlow) {
Z1674 := __e.Get(1)
_ = Z1674
__e.TailApply(PrimFunc(symshen_4linearise_1clause), Z1674)
return
}, 1)

tmp10320 := Call(__e, PrimFunc(symmap), tmp10319, let__9940)


let__9943 := tmp10320
_ = let__9943

tmp10321 := Call(__e, PrimFunc(symshen_4horn_1clause_1procedure), let__9937, let__9943)


tmp10322 := Call(__e, PrimFunc(symshen_4comb), let__9941, tmp10321)


ifres10314 = tmp10322


}

ifres10310 = ifres10314


} else {
tmp10324 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10310 = tmp10324


}

let__9936 := ifres10310
_ = let__9936

tmp10327 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9936)


if True == tmp10327 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__9936)
return
}


}, 1)

tmp10328 := Call(__e, ns2_1set, symshen_4_5defprolog_6, tmp10309)


_ = tmp10328

tmp10329 := MakeNative(func(__e *ControlFlow) {
V1677 := __e.Get(1)
_ = V1677
V1678 := __e.Get(2)
_ = V1678
tmp10373 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1678)
}
__typedArg0 := V1678
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10354 Obj

if True == tmp10373 {
tmp10371 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1678)
}
__typedArg0 := V1678
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10372 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10371)
}
__typedArg0 := tmp10371
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10356 Obj

if True == tmp10372 {
tmp10368 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1678)
}
__typedArg0 := V1678
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10369 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10368)
}
__typedArg0 := tmp10368
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10370 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10369)
}
__typedArg0 := tmp10369
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10358 Obj

if True == tmp10370 {
tmp10364 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1678)
}
__typedArg0 := V1678
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10365 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10364)
}
__typedArg0 := tmp10364
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10366 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10365)
}
__typedArg0 := tmp10365
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10367 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp10366)
}
__typedArg0 := Nil
__typedArg1 := tmp10366
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10360 Obj

if True == tmp10367 {
tmp10362 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1678)
}
__typedArg0 := V1678
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10363 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp10362)
}
__typedArg0 := Nil
__typedArg1 := tmp10362
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10361 Obj

if True == tmp10363 {
ifres10361 = True


} else {
ifres10361 = False


}

ifres10360 = ifres10361


} else {
ifres10360 = False


}

var ifres10359 Obj

if True == ifres10360 {
ifres10359 = True


} else {
ifres10359 = False


}

ifres10358 = ifres10359


} else {
ifres10358 = False


}

var ifres10357 Obj

if True == ifres10358 {
ifres10357 = True


} else {
ifres10357 = False


}

ifres10356 = ifres10357


} else {
ifres10356 = False


}

var ifres10355 Obj

if True == ifres10356 {
ifres10355 = True


} else {
ifres10355 = False


}

ifres10354 = ifres10355


} else {
ifres10354 = False


}

if True == ifres10354 {
tmp10330 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1678)
}
__typedArg0 := V1678
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10331 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10330)
}
__typedArg0 := tmp10330
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(PrimFunc(symlength), tmp10331)
return


} else {
tmp10352 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1678)
}
__typedArg0 := V1678
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10337 Obj

if True == tmp10352 {
tmp10350 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1678)
}
__typedArg0 := V1678
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10351 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10350)
}
__typedArg0 := tmp10350
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10339 Obj

if True == tmp10351 {
tmp10347 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1678)
}
__typedArg0 := V1678
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10348 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10347)
}
__typedArg0 := tmp10347
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10349 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10348)
}
__typedArg0 := tmp10348
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10341 Obj

if True == tmp10349 {
tmp10343 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1678)
}
__typedArg0 := V1678
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10344 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10343)
}
__typedArg0 := tmp10343
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10345 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10344)
}
__typedArg0 := tmp10344
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10346 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp10345)
}
__typedArg0 := Nil
__typedArg1 := tmp10345
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10342 Obj

if True == tmp10346 {
ifres10342 = True


} else {
ifres10342 = False


}

ifres10341 = ifres10342


} else {
ifres10341 = False


}

var ifres10340 Obj

if True == ifres10341 {
ifres10340 = True


} else {
ifres10340 = False


}

ifres10339 = ifres10340


} else {
ifres10339 = False


}

var ifres10338 Obj

if True == ifres10339 {
ifres10338 = True


} else {
ifres10338 = False


}

ifres10337 = ifres10338


} else {
ifres10337 = False


}

if True == ifres10337 {
tmp10332 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1678)
}
__typedArg0 := V1678
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10333 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10332)
}
__typedArg0 := tmp10332
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10334 := Call(__e, PrimFunc(symlength), tmp10333)


tmp10335 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1678)
}
__typedArg0 := V1678
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4pac_1h), V1677, tmp10334, tmp10335)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4prolog_1arity_1check)
return
}


}


}, 2)

tmp10374 := Call(__e, ns2_1set, symshen_4prolog_1arity_1check, tmp10329)


_ = tmp10374

tmp10375 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__self3 := __e.Get(3)
__selftop:
V1683 := __self1
_ = V1683
V1684 := __self2
_ = V1684
V1685 := __self3
_ = V1685
tmp10391 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V1685)
}
__typedArg0 := Nil
__typedArg1 := V1685
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp10391 {
__e.Return(V1684)
return
} else {
tmp10389 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1685)
}
__typedArg0 := V1685
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10385 Obj

if True == tmp10389 {
tmp10387 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1685)
}
__typedArg0 := V1685
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10388 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10387)
}
__typedArg0 := tmp10387
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10386 Obj

if True == tmp10388 {
ifres10386 = True


} else {
ifres10386 = False


}

ifres10385 = ifres10386


} else {
ifres10385 = False


}

if True == ifres10385 {
tmp10380 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1685)
}
__typedArg0 := V1685
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10381 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10380)
}
__typedArg0 := tmp10380
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10382 := Call(__e, PrimFunc(symlength), tmp10381)


tmp10383 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V1684, tmp10382)
}
__typedArg0 := V1684
__typedArg1 := tmp10382
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp10383 {
tmp10376 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1685)
}
__typedArg0 := V1685
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4pac_1h) == __self {
__self1, __self2, __self3 = V1683, V1684, tmp10376
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4pac_1h), V1683, V1684, tmp10376)
return


} else {
tmp10377 := Call(__e, PrimFunc(symshen_4app), V1683, MakeString("\n"), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("arity error in prolog procedure "))
__typedS1, __typedOK1 := TypedString(tmp10377)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("arity error in prolog procedure ")
__typedArg1 := tmp10377
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("arity error in prolog procedure "))
__typedS1, __typedOK1 := TypedString(tmp10377)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("arity error in prolog procedure ")
__typedArg1 := tmp10377
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4pac_1h)
return
}


}


}, 3)

tmp10392 := Call(__e, ns2_1set, symshen_4pac_1h, tmp10375)


_ = tmp10392

tmp10393 := MakeNative(func(__e *ControlFlow) {
V1686 := __e.Get(1)
_ = V1686
tmp10394 := Call(__e, PrimFunc(symshen_4_5clause_6), V1686)


let__9945 := tmp10394
_ = let__9945

tmp10407 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9945)


var ifres10395 Obj

if True == tmp10407 {
tmp10396 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10395 = tmp10396


} else {
tmp10397 := Call(__e, PrimFunc(symshen_4_5_1out), let__9945)


let__9946 := tmp10397
_ = let__9946

tmp10398 := Call(__e, PrimFunc(symshen_4in_1_6), let__9945)


let__9947 := tmp10398
_ = let__9947

tmp10399 := Call(__e, PrimFunc(symshen_4_5clauses_6), let__9947)


let__9948 := tmp10399
_ = let__9948

tmp10406 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9948)


var ifres10400 Obj

if True == tmp10406 {
tmp10401 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10400 = tmp10401


} else {
tmp10402 := Call(__e, PrimFunc(symshen_4_5_1out), let__9948)


let__9949 := tmp10402
_ = let__9949

tmp10403 := Call(__e, PrimFunc(symshen_4in_1_6), let__9948)


let__9950 := tmp10403
_ = let__9950

tmp10404 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__9946, let__9949)
}
__typedArg0 := let__9946
__typedArg1 := let__9949
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10405 := Call(__e, PrimFunc(symshen_4comb), let__9950, tmp10404)


ifres10400 = tmp10405


}

ifres10395 = ifres10400


}

let__9944 := ifres10395
_ = let__9944

tmp10423 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9944)


if True == tmp10423 {
tmp10408 := Call(__e, PrimFunc(sym_5_b_6), V1686)


let__9952 := tmp10408
_ = let__9952

tmp10419 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9952)


var ifres10409 Obj

if True == tmp10419 {
tmp10410 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10409 = tmp10410


} else {
tmp10411 := Call(__e, PrimFunc(symshen_4_5_1out), let__9952)


let__9953 := tmp10411
_ = let__9953

tmp10412 := Call(__e, PrimFunc(symshen_4in_1_6), let__9952)


let__9954 := tmp10412
_ = let__9954

tmp10417 := Call(__e, PrimFunc(symempty_2), let__9953)


var ifres10413 Obj

if True == tmp10417 {
ifres10413 = Nil


} else {
tmp10414 := Call(__e, PrimFunc(symshen_4app), let__9953, MakeString("\n ..."), symshen_4r)


tmp10416 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("Prolog syntax error here:\n "))
__typedS1, __typedOK1 := TypedString(tmp10414)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("Prolog syntax error here:\n ")
__typedArg1 := tmp10414
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("Prolog syntax error here:\n "))
__typedS1, __typedOK1 := TypedString(tmp10414)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("Prolog syntax error here:\n ")
__typedArg1 := tmp10414
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})()

ifres10413 = tmp10416


}

tmp10418 := Call(__e, PrimFunc(symshen_4comb), let__9954, ifres10413)


ifres10409 = tmp10418


}

let__9951 := ifres10409
_ = let__9951

tmp10421 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9951)


if True == tmp10421 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__9951)
return
}


} else {
__e.Return(let__9944)
return
}


}, 1)

tmp10424 := Call(__e, ns2_1set, symshen_4_5clauses_6, tmp10393)


_ = tmp10424

tmp10425 := MakeNative(func(__e *ControlFlow) {
V1698 := __e.Get(1)
_ = V1698
tmp10441 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1698)
}
__typedArg0 := V1698
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10432 Obj

if True == tmp10441 {
tmp10439 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1698)
}
__typedArg0 := V1698
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10440 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10439)
}
__typedArg0 := tmp10439
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10434 Obj

if True == tmp10440 {
tmp10436 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1698)
}
__typedArg0 := V1698
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10437 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10436)
}
__typedArg0 := tmp10436
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10438 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp10437)
}
__typedArg0 := Nil
__typedArg1 := tmp10437
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10435 Obj

if True == tmp10438 {
ifres10435 = True


} else {
ifres10435 = False


}

ifres10434 = ifres10435


} else {
ifres10434 = False


}

var ifres10433 Obj

if True == ifres10434 {
ifres10433 = True


} else {
ifres10433 = False


}

ifres10432 = ifres10433


} else {
ifres10432 = False


}

if True == ifres10432 {
tmp10426 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1698)
}
__typedArg0 := V1698
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10427 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1698)
}
__typedArg0 := V1698
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10428 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10427)
}
__typedArg0 := tmp10427
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10429 := Call(__e, PrimFunc(sym_8p), tmp10426, tmp10428)


tmp10430 := Call(__e, PrimFunc(symshen_4linearise), tmp10429)


__e.TailApply(PrimFunc(symshen_4lch), tmp10430)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4linearise_1clause)
return
}


}, 1)

tmp10442 := Call(__e, ns2_1set, symshen_4linearise_1clause, tmp10425)


_ = tmp10442

tmp10443 := MakeNative(func(__e *ControlFlow) {
V1699 := __e.Get(1)
_ = V1699
tmp10449 := Call(__e, PrimFunc(symtuple_2), V1699)


if True == tmp10449 {
tmp10444 := Call(__e, PrimFunc(symfst), V1699)


tmp10445 := Call(__e, PrimFunc(symsnd), V1699)


tmp10446 := Call(__e, PrimFunc(symshen_4lchh), tmp10445)


tmp10447 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp10446, Nil)
}
__typedArg0 := tmp10446
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp10444, tmp10447)
}
__typedArg0 := tmp10444
__typedArg1 := tmp10447
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4lch)
return
}


}, 1)

tmp10450 := Call(__e, ns2_1set, symshen_4lch, tmp10443)


_ = tmp10450

tmp10451 := MakeNative(func(__e *ControlFlow) {
V1700 := __e.Get(1)
_ = V1700
tmp10514 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1700)
}
__typedArg0 := V1700
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10463 Obj

if True == tmp10514 {
tmp10512 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1700)
}
__typedArg0 := V1700
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10513 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symwhere, tmp10512)
}
__typedArg0 := symwhere
__typedArg1 := tmp10512
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10465 Obj

if True == tmp10513 {
tmp10510 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1700)
}
__typedArg0 := V1700
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10511 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10510)
}
__typedArg0 := tmp10510
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10467 Obj

if True == tmp10511 {
tmp10507 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1700)
}
__typedArg0 := V1700
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10508 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10507)
}
__typedArg0 := tmp10507
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10509 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10508)
}
__typedArg0 := tmp10508
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10469 Obj

if True == tmp10509 {
tmp10503 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1700)
}
__typedArg0 := V1700
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10504 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10503)
}
__typedArg0 := tmp10503
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10505 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10504)
}
__typedArg0 := tmp10504
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10506 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_a, tmp10505)
}
__typedArg0 := sym_a
__typedArg1 := tmp10505
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10471 Obj

if True == tmp10506 {
tmp10499 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1700)
}
__typedArg0 := V1700
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10500 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10499)
}
__typedArg0 := tmp10499
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10501 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10500)
}
__typedArg0 := tmp10500
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10502 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10501)
}
__typedArg0 := tmp10501
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10473 Obj

if True == tmp10502 {
tmp10494 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1700)
}
__typedArg0 := V1700
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10495 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10494)
}
__typedArg0 := tmp10494
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10496 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10495)
}
__typedArg0 := tmp10495
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10497 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10496)
}
__typedArg0 := tmp10496
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10498 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10497)
}
__typedArg0 := tmp10497
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10475 Obj

if True == tmp10498 {
tmp10488 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1700)
}
__typedArg0 := V1700
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10489 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10488)
}
__typedArg0 := tmp10488
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10490 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10489)
}
__typedArg0 := tmp10489
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10491 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10490)
}
__typedArg0 := tmp10490
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10492 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10491)
}
__typedArg0 := tmp10491
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10493 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp10492)
}
__typedArg0 := Nil
__typedArg1 := tmp10492
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10477 Obj

if True == tmp10493 {
tmp10485 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1700)
}
__typedArg0 := V1700
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10486 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10485)
}
__typedArg0 := tmp10485
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10487 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10486)
}
__typedArg0 := tmp10486
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10479 Obj

if True == tmp10487 {
tmp10481 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1700)
}
__typedArg0 := V1700
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10482 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10481)
}
__typedArg0 := tmp10481
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10483 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10482)
}
__typedArg0 := tmp10482
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10484 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp10483)
}
__typedArg0 := Nil
__typedArg1 := tmp10483
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10480 Obj

if True == tmp10484 {
ifres10480 = True


} else {
ifres10480 = False


}

ifres10479 = ifres10480


} else {
ifres10479 = False


}

var ifres10478 Obj

if True == ifres10479 {
ifres10478 = True


} else {
ifres10478 = False


}

ifres10477 = ifres10478


} else {
ifres10477 = False


}

var ifres10476 Obj

if True == ifres10477 {
ifres10476 = True


} else {
ifres10476 = False


}

ifres10475 = ifres10476


} else {
ifres10475 = False


}

var ifres10474 Obj

if True == ifres10475 {
ifres10474 = True


} else {
ifres10474 = False


}

ifres10473 = ifres10474


} else {
ifres10473 = False


}

var ifres10472 Obj

if True == ifres10473 {
ifres10472 = True


} else {
ifres10472 = False


}

ifres10471 = ifres10472


} else {
ifres10471 = False


}

var ifres10470 Obj

if True == ifres10471 {
ifres10470 = True


} else {
ifres10470 = False


}

ifres10469 = ifres10470


} else {
ifres10469 = False


}

var ifres10468 Obj

if True == ifres10469 {
ifres10468 = True


} else {
ifres10468 = False


}

ifres10467 = ifres10468


} else {
ifres10467 = False


}

var ifres10466 Obj

if True == ifres10467 {
ifres10466 = True


} else {
ifres10466 = False


}

ifres10465 = ifres10466


} else {
ifres10465 = False


}

var ifres10464 Obj

if True == ifres10465 {
ifres10464 = True


} else {
ifres10464 = False


}

ifres10463 = ifres10464


} else {
ifres10463 = False


}

if True == ifres10463 {
tmp10453 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_doccurs_d)
}
__typedArg0 := symshen_4_doccurs_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

var ifres10452 Obj

if True == tmp10453 {
ifres10452 = symis_b


} else {
ifres10452 = symis


}

tmp10454 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1700)
}
__typedArg0 := V1700
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10455 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10454)
}
__typedArg0 := tmp10454
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10456 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10455)
}
__typedArg0 := tmp10455
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10457 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(ifres10452, tmp10456)
}
__typedArg0 := ifres10452
__typedArg1 := tmp10456
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10458 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1700)
}
__typedArg0 := V1700
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10459 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10458)
}
__typedArg0 := tmp10458
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10460 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10459)
}
__typedArg0 := tmp10459
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10461 := Call(__e, PrimFunc(symshen_4lchh), tmp10460)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp10457, tmp10461)
}
__typedArg0 := tmp10457
__typedArg1 := tmp10461
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return(V1700)
return
}


}, 1)

tmp10515 := Call(__e, ns2_1set, symshen_4lchh, tmp10451)


_ = tmp10515

tmp10516 := MakeNative(func(__e *ControlFlow) {
V1701 := __e.Get(1)
_ = V1701
tmp10517 := Call(__e, PrimFunc(symshen_4_5head_6), V1701)


let__9956 := tmp10517
_ = let__9956

tmp10540 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9956)


var ifres10518 Obj

if True == tmp10540 {
tmp10519 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10518 = tmp10519


} else {
tmp10520 := Call(__e, PrimFunc(symshen_4_5_1out), let__9956)


let__9957 := tmp10520
_ = let__9957

tmp10521 := Call(__e, PrimFunc(symshen_4in_1_6), let__9956)


let__9958 := tmp10521
_ = let__9958

tmp10539 := Call(__e, PrimFunc(symshen_4hds_a_2), let__9958, sym_5_1_1)


var ifres10522 Obj

if True == tmp10539 {
tmp10523 := Call(__e, PrimFunc(symtail), let__9958)


let__9959 := tmp10523
_ = let__9959

tmp10524 := Call(__e, PrimFunc(symshen_4_5body_6), let__9959)


let__9960 := tmp10524
_ = let__9960

tmp10537 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9960)


var ifres10525 Obj

if True == tmp10537 {
tmp10526 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10525 = tmp10526


} else {
tmp10527 := Call(__e, PrimFunc(symshen_4_5_1out), let__9960)


let__9961 := tmp10527
_ = let__9961

tmp10528 := Call(__e, PrimFunc(symshen_4in_1_6), let__9960)


let__9962 := tmp10528
_ = let__9962

tmp10529 := Call(__e, PrimFunc(symshen_4_5sc_6), let__9962)


let__9963 := tmp10529
_ = let__9963

tmp10536 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9963)


var ifres10530 Obj

if True == tmp10536 {
tmp10531 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10530 = tmp10531


} else {
tmp10532 := Call(__e, PrimFunc(symshen_4in_1_6), let__9963)


let__9964 := tmp10532
_ = let__9964

tmp10533 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__9961, Nil)
}
__typedArg0 := let__9961
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10534 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__9957, tmp10533)
}
__typedArg0 := let__9957
__typedArg1 := tmp10533
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10535 := Call(__e, PrimFunc(symshen_4comb), let__9964, tmp10534)


ifres10530 = tmp10535


}

ifres10525 = ifres10530


}

ifres10522 = ifres10525


} else {
tmp10538 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10522 = tmp10538


}

ifres10518 = ifres10522


}

let__9955 := ifres10518
_ = let__9955

tmp10542 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9955)


if True == tmp10542 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__9955)
return
}


}, 1)

tmp10543 := Call(__e, ns2_1set, symshen_4_5clause_6, tmp10516)


_ = tmp10543

tmp10544 := MakeNative(func(__e *ControlFlow) {
V1712 := __e.Get(1)
_ = V1712
tmp10545 := Call(__e, PrimFunc(symshen_4_5hterm_6), V1712)


let__9966 := tmp10545
_ = let__9966

tmp10558 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9966)


var ifres10546 Obj

if True == tmp10558 {
tmp10547 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10546 = tmp10547


} else {
tmp10548 := Call(__e, PrimFunc(symshen_4_5_1out), let__9966)


let__9967 := tmp10548
_ = let__9967

tmp10549 := Call(__e, PrimFunc(symshen_4in_1_6), let__9966)


let__9968 := tmp10549
_ = let__9968

tmp10550 := Call(__e, PrimFunc(symshen_4_5head_6), let__9968)


let__9969 := tmp10550
_ = let__9969

tmp10557 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9969)


var ifres10551 Obj

if True == tmp10557 {
tmp10552 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10551 = tmp10552


} else {
tmp10553 := Call(__e, PrimFunc(symshen_4_5_1out), let__9969)


let__9970 := tmp10553
_ = let__9970

tmp10554 := Call(__e, PrimFunc(symshen_4in_1_6), let__9969)


let__9971 := tmp10554
_ = let__9971

tmp10555 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__9967, let__9970)
}
__typedArg0 := let__9967
__typedArg1 := let__9970
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10556 := Call(__e, PrimFunc(symshen_4comb), let__9971, tmp10555)


ifres10551 = tmp10556


}

ifres10546 = ifres10551


}

let__9965 := ifres10546
_ = let__9965

tmp10568 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9965)


if True == tmp10568 {
tmp10559 := Call(__e, PrimFunc(sym_5e_6), V1712)


let__9973 := tmp10559
_ = let__9973

tmp10564 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9973)


var ifres10560 Obj

if True == tmp10564 {
tmp10561 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10560 = tmp10561


} else {
tmp10562 := Call(__e, PrimFunc(symshen_4in_1_6), let__9973)


let__9974 := tmp10562
_ = let__9974

tmp10563 := Call(__e, PrimFunc(symshen_4comb), let__9974, Nil)


ifres10560 = tmp10563


}

let__9972 := ifres10560
_ = let__9972

tmp10566 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9972)


if True == tmp10566 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__9972)
return
}


} else {
__e.Return(let__9965)
return
}


}, 1)

tmp10569 := Call(__e, ns2_1set, symshen_4_5head_6, tmp10544)


_ = tmp10569

tmp10570 := MakeNative(func(__e *ControlFlow) {
V1723 := __e.Get(1)
_ = V1723
tmp10583 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1723)
}
__typedArg0 := V1723
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10571 Obj

if True == tmp10583 {
tmp10572 := Call(__e, PrimFunc(symhead), V1723)


let__9976 := tmp10572
_ = let__9976

tmp10573 := Call(__e, PrimFunc(symtail), V1723)


let__9977 := tmp10573
_ = let__9977

tmp10581 := Call(__e, PrimFunc(symatom_2), let__9976)


var ifres10577 Obj

if True == tmp10581 {
tmp10579 := Call(__e, PrimFunc(symshen_4prolog_1keyword_2), let__9976)


tmp10580 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp10579)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp10579
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres10578 Obj

if True == tmp10580 {
ifres10578 = True


} else {
ifres10578 = False


}

ifres10577 = ifres10578


} else {
ifres10577 = False


}

var ifres10574 Obj

if True == ifres10577 {
tmp10575 := Call(__e, PrimFunc(symshen_4comb), let__9977, let__9976)


ifres10574 = tmp10575


} else {
tmp10576 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10574 = tmp10576


}

ifres10571 = ifres10574


} else {
tmp10582 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10571 = tmp10582


}

let__9975 := ifres10571
_ = let__9975

tmp10737 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9975)


if True == tmp10737 {
tmp10593 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1723)
}
__typedArg0 := V1723
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10584 Obj

if True == tmp10593 {
tmp10585 := Call(__e, PrimFunc(symhead), V1723)


let__9979 := tmp10585
_ = let__9979

tmp10586 := Call(__e, PrimFunc(symtail), V1723)


let__9980 := tmp10586
_ = let__9980

tmp10590 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp10591 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__9979, tmp10590)
}
__typedArg0 := let__9979
__typedArg1 := tmp10590
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10587 Obj

if True == tmp10591 {
tmp10588 := Call(__e, PrimFunc(symshen_4comb), let__9980, let__9979)


ifres10587 = tmp10588


} else {
tmp10589 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10587 = tmp10589


}

ifres10584 = ifres10587


} else {
tmp10592 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10584 = tmp10592


}

let__9978 := ifres10584
_ = let__9978

tmp10735 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9978)


if True == tmp10735 {
tmp10623 := Call(__e, PrimFunc(symshen_4ccons_2), V1723)


var ifres10594 Obj

if True == tmp10623 {
tmp10595 := Call(__e, PrimFunc(symhead), V1723)


let__9982 := tmp10595
_ = let__9982

tmp10596 := Call(__e, PrimFunc(symtail), V1723)


let__9983 := tmp10596
_ = let__9983

tmp10621 := Call(__e, PrimFunc(symshen_4hds_a_2), let__9982, symcons)


var ifres10597 Obj

if True == tmp10621 {
tmp10598 := Call(__e, PrimFunc(symtail), let__9982)


let__9984 := tmp10598
_ = let__9984

tmp10599 := Call(__e, PrimFunc(symshen_4_5hterm1_6), let__9984)


let__9985 := tmp10599
_ = let__9985

tmp10619 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9985)


var ifres10600 Obj

if True == tmp10619 {
tmp10601 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10600 = tmp10601


} else {
tmp10602 := Call(__e, PrimFunc(symshen_4_5_1out), let__9985)


let__9986 := tmp10602
_ = let__9986

tmp10603 := Call(__e, PrimFunc(symshen_4in_1_6), let__9985)


let__9987 := tmp10603
_ = let__9987

tmp10604 := Call(__e, PrimFunc(symshen_4_5hterm2_6), let__9987)


let__9988 := tmp10604
_ = let__9988

tmp10618 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9988)


var ifres10605 Obj

if True == tmp10618 {
tmp10606 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10605 = tmp10606


} else {
tmp10607 := Call(__e, PrimFunc(symshen_4_5_1out), let__9988)


let__9989 := tmp10607
_ = let__9989

tmp10608 := Call(__e, PrimFunc(symshen_4in_1_6), let__9988)


let__9990 := tmp10608
_ = let__9990

tmp10609 := Call(__e, PrimFunc(sym_5end_6), let__9990)


let__9991 := tmp10609
_ = let__9991

tmp10617 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9991)


var ifres10610 Obj

if True == tmp10617 {
tmp10611 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10610 = tmp10611


} else {
tmp10612 := Call(__e, PrimFunc(symshen_4in_1_6), let__9991)


let__9992 := tmp10612
_ = let__9992

tmp10613 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__9989, Nil)
}
__typedArg0 := let__9989
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10614 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__9986, tmp10613)
}
__typedArg0 := let__9986
__typedArg1 := tmp10613
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10615 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp10614)
}
__typedArg0 := symcons
__typedArg1 := tmp10614
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10616 := Call(__e, PrimFunc(symshen_4comb), let__9983, tmp10615)


ifres10610 = tmp10616


}

ifres10605 = ifres10610


}

ifres10600 = ifres10605


}

ifres10597 = ifres10600


} else {
tmp10620 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10597 = tmp10620


}

ifres10594 = ifres10597


} else {
tmp10622 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10594 = tmp10622


}

let__9981 := ifres10594
_ = let__9981

tmp10733 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9981)


if True == tmp10733 {
tmp10646 := Call(__e, PrimFunc(symshen_4ccons_2), V1723)


var ifres10624 Obj

if True == tmp10646 {
tmp10625 := Call(__e, PrimFunc(symhead), V1723)


let__9994 := tmp10625
_ = let__9994

tmp10626 := Call(__e, PrimFunc(symtail), V1723)


let__9995 := tmp10626
_ = let__9995

tmp10644 := Call(__e, PrimFunc(symshen_4hds_a_2), let__9994, sym_7)


var ifres10627 Obj

if True == tmp10644 {
tmp10628 := Call(__e, PrimFunc(symtail), let__9994)


let__9996 := tmp10628
_ = let__9996

tmp10629 := Call(__e, PrimFunc(symshen_4_5hterm_6), let__9996)


let__9997 := tmp10629
_ = let__9997

tmp10642 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9997)


var ifres10630 Obj

if True == tmp10642 {
tmp10631 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10630 = tmp10631


} else {
tmp10632 := Call(__e, PrimFunc(symshen_4_5_1out), let__9997)


let__9998 := tmp10632
_ = let__9998

tmp10633 := Call(__e, PrimFunc(symshen_4in_1_6), let__9997)


let__9999 := tmp10633
_ = let__9999

tmp10634 := Call(__e, PrimFunc(sym_5end_6), let__9999)


let__10000 := tmp10634
_ = let__10000

tmp10641 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10000)


var ifres10635 Obj

if True == tmp10641 {
tmp10636 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10635 = tmp10636


} else {
tmp10637 := Call(__e, PrimFunc(symshen_4in_1_6), let__10000)


let__10001 := tmp10637
_ = let__10001

tmp10638 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__9998, Nil)
}
__typedArg0 := let__9998
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10639 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4_7m, tmp10638)
}
__typedArg0 := symshen_4_7m
__typedArg1 := tmp10638
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10640 := Call(__e, PrimFunc(symshen_4comb), let__9995, tmp10639)


ifres10635 = tmp10640


}

ifres10630 = ifres10635


}

ifres10627 = ifres10630


} else {
tmp10643 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10627 = tmp10643


}

ifres10624 = ifres10627


} else {
tmp10645 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10624 = tmp10645


}

let__9993 := ifres10624
_ = let__9993

tmp10731 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9993)


if True == tmp10731 {
tmp10669 := Call(__e, PrimFunc(symshen_4ccons_2), V1723)


var ifres10647 Obj

if True == tmp10669 {
tmp10648 := Call(__e, PrimFunc(symhead), V1723)


let__10003 := tmp10648
_ = let__10003

tmp10649 := Call(__e, PrimFunc(symtail), V1723)


let__10004 := tmp10649
_ = let__10004

tmp10667 := Call(__e, PrimFunc(symshen_4hds_a_2), let__10003, sym_1)


var ifres10650 Obj

if True == tmp10667 {
tmp10651 := Call(__e, PrimFunc(symtail), let__10003)


let__10005 := tmp10651
_ = let__10005

tmp10652 := Call(__e, PrimFunc(symshen_4_5hterm_6), let__10005)


let__10006 := tmp10652
_ = let__10006

tmp10665 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10006)


var ifres10653 Obj

if True == tmp10665 {
tmp10654 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10653 = tmp10654


} else {
tmp10655 := Call(__e, PrimFunc(symshen_4_5_1out), let__10006)


let__10007 := tmp10655
_ = let__10007

tmp10656 := Call(__e, PrimFunc(symshen_4in_1_6), let__10006)


let__10008 := tmp10656
_ = let__10008

tmp10657 := Call(__e, PrimFunc(sym_5end_6), let__10008)


let__10009 := tmp10657
_ = let__10009

tmp10664 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10009)


var ifres10658 Obj

if True == tmp10664 {
tmp10659 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10658 = tmp10659


} else {
tmp10660 := Call(__e, PrimFunc(symshen_4in_1_6), let__10009)


let__10010 := tmp10660
_ = let__10010

tmp10661 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10007, Nil)
}
__typedArg0 := let__10007
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10662 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4_1m, tmp10661)
}
__typedArg0 := symshen_4_1m
__typedArg1 := tmp10661
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10663 := Call(__e, PrimFunc(symshen_4comb), let__10004, tmp10662)


ifres10658 = tmp10663


}

ifres10653 = ifres10658


}

ifres10650 = ifres10653


} else {
tmp10666 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10650 = tmp10666


}

ifres10647 = ifres10650


} else {
tmp10668 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10647 = tmp10668


}

let__10002 := ifres10647
_ = let__10002

tmp10729 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10002)


if True == tmp10729 {
tmp10696 := Call(__e, PrimFunc(symshen_4ccons_2), V1723)


var ifres10670 Obj

if True == tmp10696 {
tmp10671 := Call(__e, PrimFunc(symhead), V1723)


let__10012 := tmp10671
_ = let__10012

tmp10672 := Call(__e, PrimFunc(symtail), V1723)


let__10013 := tmp10672
_ = let__10013

tmp10694 := Call(__e, PrimFunc(symshen_4hds_a_2), let__10012, symmode)


var ifres10673 Obj

if True == tmp10694 {
tmp10674 := Call(__e, PrimFunc(symtail), let__10012)


let__10014 := tmp10674
_ = let__10014

tmp10675 := Call(__e, PrimFunc(symshen_4_5hterm_6), let__10014)


let__10015 := tmp10675
_ = let__10015

tmp10692 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10015)


var ifres10676 Obj

if True == tmp10692 {
tmp10677 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10676 = tmp10677


} else {
tmp10678 := Call(__e, PrimFunc(symshen_4_5_1out), let__10015)


let__10016 := tmp10678
_ = let__10016

tmp10679 := Call(__e, PrimFunc(symshen_4in_1_6), let__10015)


let__10017 := tmp10679
_ = let__10017

tmp10691 := Call(__e, PrimFunc(symshen_4hds_a_2), let__10017, sym_7)


var ifres10680 Obj

if True == tmp10691 {
tmp10681 := Call(__e, PrimFunc(symtail), let__10017)


let__10018 := tmp10681
_ = let__10018

tmp10682 := Call(__e, PrimFunc(sym_5end_6), let__10018)


let__10019 := tmp10682
_ = let__10019

tmp10689 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10019)


var ifres10683 Obj

if True == tmp10689 {
tmp10684 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10683 = tmp10684


} else {
tmp10685 := Call(__e, PrimFunc(symshen_4in_1_6), let__10019)


let__10020 := tmp10685
_ = let__10020

tmp10686 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10016, Nil)
}
__typedArg0 := let__10016
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10687 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4_7m, tmp10686)
}
__typedArg0 := symshen_4_7m
__typedArg1 := tmp10686
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10688 := Call(__e, PrimFunc(symshen_4comb), let__10013, tmp10687)


ifres10683 = tmp10688


}

ifres10680 = ifres10683


} else {
tmp10690 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10680 = tmp10690


}

ifres10676 = ifres10680


}

ifres10673 = ifres10676


} else {
tmp10693 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10673 = tmp10693


}

ifres10670 = ifres10673


} else {
tmp10695 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10670 = tmp10695


}

let__10011 := ifres10670
_ = let__10011

tmp10727 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10011)


if True == tmp10727 {
tmp10723 := Call(__e, PrimFunc(symshen_4ccons_2), V1723)


var ifres10697 Obj

if True == tmp10723 {
tmp10698 := Call(__e, PrimFunc(symhead), V1723)


let__10022 := tmp10698
_ = let__10022

tmp10699 := Call(__e, PrimFunc(symtail), V1723)


let__10023 := tmp10699
_ = let__10023

tmp10721 := Call(__e, PrimFunc(symshen_4hds_a_2), let__10022, symmode)


var ifres10700 Obj

if True == tmp10721 {
tmp10701 := Call(__e, PrimFunc(symtail), let__10022)


let__10024 := tmp10701
_ = let__10024

tmp10702 := Call(__e, PrimFunc(symshen_4_5hterm_6), let__10024)


let__10025 := tmp10702
_ = let__10025

tmp10719 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10025)


var ifres10703 Obj

if True == tmp10719 {
tmp10704 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10703 = tmp10704


} else {
tmp10705 := Call(__e, PrimFunc(symshen_4_5_1out), let__10025)


let__10026 := tmp10705
_ = let__10026

tmp10706 := Call(__e, PrimFunc(symshen_4in_1_6), let__10025)


let__10027 := tmp10706
_ = let__10027

tmp10718 := Call(__e, PrimFunc(symshen_4hds_a_2), let__10027, sym_1)


var ifres10707 Obj

if True == tmp10718 {
tmp10708 := Call(__e, PrimFunc(symtail), let__10027)


let__10028 := tmp10708
_ = let__10028

tmp10709 := Call(__e, PrimFunc(sym_5end_6), let__10028)


let__10029 := tmp10709
_ = let__10029

tmp10716 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10029)


var ifres10710 Obj

if True == tmp10716 {
tmp10711 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10710 = tmp10711


} else {
tmp10712 := Call(__e, PrimFunc(symshen_4in_1_6), let__10029)


let__10030 := tmp10712
_ = let__10030

tmp10713 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10026, Nil)
}
__typedArg0 := let__10026
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10714 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4_1m, tmp10713)
}
__typedArg0 := symshen_4_1m
__typedArg1 := tmp10713
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10715 := Call(__e, PrimFunc(symshen_4comb), let__10023, tmp10714)


ifres10710 = tmp10715


}

ifres10707 = ifres10710


} else {
tmp10717 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10707 = tmp10717


}

ifres10703 = ifres10707


}

ifres10700 = ifres10703


} else {
tmp10720 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10700 = tmp10720


}

ifres10697 = ifres10700


} else {
tmp10722 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10697 = tmp10722


}

let__10021 := ifres10697
_ = let__10021

tmp10725 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10021)


if True == tmp10725 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__10021)
return
}


} else {
__e.Return(let__10011)
return
}


} else {
__e.Return(let__10002)
return
}


} else {
__e.Return(let__9993)
return
}


} else {
__e.Return(let__9981)
return
}


} else {
__e.Return(let__9978)
return
}


} else {
__e.Return(let__9975)
return
}


}, 1)

tmp10738 := Call(__e, ns2_1set, symshen_4_5hterm_6, tmp10570)


_ = tmp10738

tmp10739 := MakeNative(func(__e *ControlFlow) {
V1780 := __e.Get(1)
_ = V1780
tmp10740 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(";"))
}
__typedArg0 := MakeString(";")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp10741 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1_1, Nil)
}
__typedArg0 := sym_5_1_1
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10742 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp10740, tmp10741)
}
__typedArg0 := tmp10740
__typedArg1 := tmp10741
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symelement_2), V1780, tmp10742)
return


}, 1)

tmp10743 := Call(__e, ns2_1set, symshen_4prolog_1keyword_2, tmp10739)


_ = tmp10743

tmp10744 := MakeNative(func(__e *ControlFlow) {
V1781 := __e.Get(1)
_ = V1781
tmp10757 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsymbol_2) {
return PrimIsSymbol(V1781)
}
__typedArg0 := V1781
return Call(__e, PrimFunc(symsymbol_2), __typedArg0)
})()

if True == tmp10757 {
__e.Return(True)
return
} else {
tmp10755 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_2) {
return PrimIsString(V1781)
}
__typedArg0 := V1781
return Call(__e, PrimFunc(symstring_2), __typedArg0)
})()

var ifres10746 Obj

if True == tmp10755 {
ifres10746 = True


} else {
tmp10754 := Call(__e, PrimFunc(symboolean_2), V1781)


var ifres10748 Obj

if True == tmp10754 {
ifres10748 = True


} else {
tmp10753 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnumber_2) {
return PrimIsNumber(V1781)
}
__typedArg0 := V1781
return Call(__e, PrimFunc(symnumber_2), __typedArg0)
})()

var ifres10750 Obj

if True == tmp10753 {
ifres10750 = True


} else {
tmp10752 := Call(__e, PrimFunc(symempty_2), V1781)


var ifres10751 Obj

if True == tmp10752 {
ifres10751 = True


} else {
ifres10751 = False


}

ifres10750 = ifres10751


}

var ifres10749 Obj

if True == ifres10750 {
ifres10749 = True


} else {
ifres10749 = False


}

ifres10748 = ifres10749


}

var ifres10747 Obj

if True == ifres10748 {
ifres10747 = True


} else {
ifres10747 = False


}

ifres10746 = ifres10747


}

if True == ifres10746 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp10758 := Call(__e, ns2_1set, symatom_2, tmp10744)


_ = tmp10758

tmp10759 := MakeNative(func(__e *ControlFlow) {
V1782 := __e.Get(1)
_ = V1782
tmp10760 := Call(__e, PrimFunc(symshen_4_5hterm_6), V1782)


let__10032 := tmp10760
_ = let__10032

tmp10766 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10032)


var ifres10761 Obj

if True == tmp10766 {
tmp10762 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10761 = tmp10762


} else {
tmp10763 := Call(__e, PrimFunc(symshen_4_5_1out), let__10032)


let__10033 := tmp10763
_ = let__10033

tmp10764 := Call(__e, PrimFunc(symshen_4in_1_6), let__10032)


let__10034 := tmp10764
_ = let__10034

tmp10765 := Call(__e, PrimFunc(symshen_4comb), let__10034, let__10033)


ifres10761 = tmp10765


}

let__10031 := ifres10761
_ = let__10031

tmp10768 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10031)


if True == tmp10768 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__10031)
return
}


}, 1)

tmp10769 := Call(__e, ns2_1set, symshen_4_5hterm1_6, tmp10759)


_ = tmp10769

tmp10770 := MakeNative(func(__e *ControlFlow) {
V1787 := __e.Get(1)
_ = V1787
tmp10771 := Call(__e, PrimFunc(symshen_4_5hterm_6), V1787)


let__10036 := tmp10771
_ = let__10036

tmp10777 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10036)


var ifres10772 Obj

if True == tmp10777 {
tmp10773 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10772 = tmp10773


} else {
tmp10774 := Call(__e, PrimFunc(symshen_4_5_1out), let__10036)


let__10037 := tmp10774
_ = let__10037

tmp10775 := Call(__e, PrimFunc(symshen_4in_1_6), let__10036)


let__10038 := tmp10775
_ = let__10038

tmp10776 := Call(__e, PrimFunc(symshen_4comb), let__10038, let__10037)


ifres10772 = tmp10776


}

let__10035 := ifres10772
_ = let__10035

tmp10779 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10035)


if True == tmp10779 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__10035)
return
}


}, 1)

tmp10780 := Call(__e, ns2_1set, symshen_4_5hterm2_6, tmp10770)


_ = tmp10780

tmp10781 := MakeNative(func(__e *ControlFlow) {
V1792 := __e.Get(1)
_ = V1792
tmp10782 := Call(__e, PrimFunc(symshen_4_5literal_6), V1792)


let__10040 := tmp10782
_ = let__10040

tmp10795 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10040)


var ifres10783 Obj

if True == tmp10795 {
tmp10784 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10783 = tmp10784


} else {
tmp10785 := Call(__e, PrimFunc(symshen_4_5_1out), let__10040)


let__10041 := tmp10785
_ = let__10041

tmp10786 := Call(__e, PrimFunc(symshen_4in_1_6), let__10040)


let__10042 := tmp10786
_ = let__10042

tmp10787 := Call(__e, PrimFunc(symshen_4_5body_6), let__10042)


let__10043 := tmp10787
_ = let__10043

tmp10794 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10043)


var ifres10788 Obj

if True == tmp10794 {
tmp10789 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10788 = tmp10789


} else {
tmp10790 := Call(__e, PrimFunc(symshen_4_5_1out), let__10043)


let__10044 := tmp10790
_ = let__10044

tmp10791 := Call(__e, PrimFunc(symshen_4in_1_6), let__10043)


let__10045 := tmp10791
_ = let__10045

tmp10792 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10041, let__10044)
}
__typedArg0 := let__10041
__typedArg1 := let__10044
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10793 := Call(__e, PrimFunc(symshen_4comb), let__10045, tmp10792)


ifres10788 = tmp10793


}

ifres10783 = ifres10788


}

let__10039 := ifres10783
_ = let__10039

tmp10805 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10039)


if True == tmp10805 {
tmp10796 := Call(__e, PrimFunc(sym_5e_6), V1792)


let__10047 := tmp10796
_ = let__10047

tmp10801 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10047)


var ifres10797 Obj

if True == tmp10801 {
tmp10798 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10797 = tmp10798


} else {
tmp10799 := Call(__e, PrimFunc(symshen_4in_1_6), let__10047)


let__10048 := tmp10799
_ = let__10048

tmp10800 := Call(__e, PrimFunc(symshen_4comb), let__10048, Nil)


ifres10797 = tmp10800


}

let__10046 := ifres10797
_ = let__10046

tmp10803 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10046)


if True == tmp10803 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__10046)
return
}


} else {
__e.Return(let__10039)
return
}


}, 1)

tmp10806 := Call(__e, ns2_1set, symshen_4_5body_6, tmp10781)


_ = tmp10806

tmp10807 := MakeNative(func(__e *ControlFlow) {
V1803 := __e.Get(1)
_ = V1803
tmp10812 := Call(__e, PrimFunc(symshen_4hds_a_2), V1803, sym_b)


var ifres10808 Obj

if True == tmp10812 {
tmp10809 := Call(__e, PrimFunc(symtail), V1803)


let__10050 := tmp10809
_ = let__10050

tmp10810 := Call(__e, PrimFunc(symshen_4comb), let__10050, sym_b)


ifres10808 = tmp10810


} else {
tmp10811 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10808 = tmp10811


}

let__10049 := ifres10808
_ = let__10049

tmp10833 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10049)


if True == tmp10833 {
tmp10829 := Call(__e, PrimFunc(symshen_4ccons_2), V1803)


var ifres10813 Obj

if True == tmp10829 {
tmp10814 := Call(__e, PrimFunc(symhead), V1803)


let__10052 := tmp10814
_ = let__10052

tmp10815 := Call(__e, PrimFunc(symtail), V1803)


let__10053 := tmp10815
_ = let__10053

tmp10816 := Call(__e, PrimFunc(symshen_4_5bterms_6), let__10052)


let__10054 := tmp10816
_ = let__10054

tmp10827 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10054)


var ifres10817 Obj

if True == tmp10827 {
tmp10818 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10817 = tmp10818


} else {
tmp10819 := Call(__e, PrimFunc(symshen_4_5_1out), let__10054)


let__10055 := tmp10819
_ = let__10055

tmp10820 := Call(__e, PrimFunc(symshen_4in_1_6), let__10054)


let__10056 := tmp10820
_ = let__10056

tmp10821 := Call(__e, PrimFunc(sym_5end_6), let__10056)


let__10057 := tmp10821
_ = let__10057

tmp10826 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10057)


var ifres10822 Obj

if True == tmp10826 {
tmp10823 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10822 = tmp10823


} else {
tmp10824 := Call(__e, PrimFunc(symshen_4in_1_6), let__10057)


let__10058 := tmp10824
_ = let__10058

tmp10825 := Call(__e, PrimFunc(symshen_4comb), let__10053, let__10055)


ifres10822 = tmp10825


}

ifres10817 = ifres10822


}

ifres10813 = ifres10817


} else {
tmp10828 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10813 = tmp10828


}

let__10051 := ifres10813
_ = let__10051

tmp10831 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10051)


if True == tmp10831 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__10051)
return
}


} else {
__e.Return(let__10049)
return
}


}, 1)

tmp10834 := Call(__e, ns2_1set, symshen_4_5literal_6, tmp10807)


_ = tmp10834

tmp10835 := MakeNative(func(__e *ControlFlow) {
V1814 := __e.Get(1)
_ = V1814
tmp10836 := Call(__e, PrimFunc(symshen_4_5bterm_6), V1814)


let__10060 := tmp10836
_ = let__10060

tmp10849 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10060)


var ifres10837 Obj

if True == tmp10849 {
tmp10838 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10837 = tmp10838


} else {
tmp10839 := Call(__e, PrimFunc(symshen_4_5_1out), let__10060)


let__10061 := tmp10839
_ = let__10061

tmp10840 := Call(__e, PrimFunc(symshen_4in_1_6), let__10060)


let__10062 := tmp10840
_ = let__10062

tmp10841 := Call(__e, PrimFunc(symshen_4_5bterms_6), let__10062)


let__10063 := tmp10841
_ = let__10063

tmp10848 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10063)


var ifres10842 Obj

if True == tmp10848 {
tmp10843 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10842 = tmp10843


} else {
tmp10844 := Call(__e, PrimFunc(symshen_4_5_1out), let__10063)


let__10064 := tmp10844
_ = let__10064

tmp10845 := Call(__e, PrimFunc(symshen_4in_1_6), let__10063)


let__10065 := tmp10845
_ = let__10065

tmp10846 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10061, let__10064)
}
__typedArg0 := let__10061
__typedArg1 := let__10064
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10847 := Call(__e, PrimFunc(symshen_4comb), let__10065, tmp10846)


ifres10842 = tmp10847


}

ifres10837 = ifres10842


}

let__10059 := ifres10837
_ = let__10059

tmp10859 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10059)


if True == tmp10859 {
tmp10850 := Call(__e, PrimFunc(sym_5e_6), V1814)


let__10067 := tmp10850
_ = let__10067

tmp10855 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10067)


var ifres10851 Obj

if True == tmp10855 {
tmp10852 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10851 = tmp10852


} else {
tmp10853 := Call(__e, PrimFunc(symshen_4in_1_6), let__10067)


let__10068 := tmp10853
_ = let__10068

tmp10854 := Call(__e, PrimFunc(symshen_4comb), let__10068, Nil)


ifres10851 = tmp10854


}

let__10066 := ifres10851
_ = let__10066

tmp10857 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10066)


if True == tmp10857 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__10066)
return
}


} else {
__e.Return(let__10059)
return
}


}, 1)

tmp10860 := Call(__e, ns2_1set, symshen_4_5bterms_6, tmp10835)


_ = tmp10860

tmp10861 := MakeNative(func(__e *ControlFlow) {
V1825 := __e.Get(1)
_ = V1825
tmp10862 := Call(__e, PrimFunc(symshen_4_5wildcard_6), V1825)


let__10070 := tmp10862
_ = let__10070

tmp10868 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10070)


var ifres10863 Obj

if True == tmp10868 {
tmp10864 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10863 = tmp10864


} else {
tmp10865 := Call(__e, PrimFunc(symshen_4_5_1out), let__10070)


let__10071 := tmp10865
_ = let__10071

tmp10866 := Call(__e, PrimFunc(symshen_4in_1_6), let__10070)


let__10072 := tmp10866
_ = let__10072

tmp10867 := Call(__e, PrimFunc(symshen_4comb), let__10072, let__10071)


ifres10863 = tmp10867


}

let__10069 := ifres10863
_ = let__10069

tmp10900 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10069)


if True == tmp10900 {
tmp10877 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1825)
}
__typedArg0 := V1825
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10869 Obj

if True == tmp10877 {
tmp10870 := Call(__e, PrimFunc(symhead), V1825)


let__10074 := tmp10870
_ = let__10074

tmp10871 := Call(__e, PrimFunc(symtail), V1825)


let__10075 := tmp10871
_ = let__10075

tmp10875 := Call(__e, PrimFunc(symatom_2), let__10074)


var ifres10872 Obj

if True == tmp10875 {
tmp10873 := Call(__e, PrimFunc(symshen_4comb), let__10075, let__10074)


ifres10872 = tmp10873


} else {
tmp10874 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10872 = tmp10874


}

ifres10869 = ifres10872


} else {
tmp10876 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10869 = tmp10876


}

let__10073 := ifres10869
_ = let__10073

tmp10898 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10073)


if True == tmp10898 {
tmp10894 := Call(__e, PrimFunc(symshen_4ccons_2), V1825)


var ifres10878 Obj

if True == tmp10894 {
tmp10879 := Call(__e, PrimFunc(symhead), V1825)


let__10077 := tmp10879
_ = let__10077

tmp10880 := Call(__e, PrimFunc(symtail), V1825)


let__10078 := tmp10880
_ = let__10078

tmp10881 := Call(__e, PrimFunc(symshen_4_5bterms_6), let__10077)


let__10079 := tmp10881
_ = let__10079

tmp10892 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10079)


var ifres10882 Obj

if True == tmp10892 {
tmp10883 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10882 = tmp10883


} else {
tmp10884 := Call(__e, PrimFunc(symshen_4_5_1out), let__10079)


let__10080 := tmp10884
_ = let__10080

tmp10885 := Call(__e, PrimFunc(symshen_4in_1_6), let__10079)


let__10081 := tmp10885
_ = let__10081

tmp10886 := Call(__e, PrimFunc(sym_5end_6), let__10081)


let__10082 := tmp10886
_ = let__10082

tmp10891 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10082)


var ifres10887 Obj

if True == tmp10891 {
tmp10888 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10887 = tmp10888


} else {
tmp10889 := Call(__e, PrimFunc(symshen_4in_1_6), let__10082)


let__10083 := tmp10889
_ = let__10083

tmp10890 := Call(__e, PrimFunc(symshen_4comb), let__10078, let__10080)


ifres10887 = tmp10890


}

ifres10882 = ifres10887


}

ifres10878 = ifres10882


} else {
tmp10893 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10878 = tmp10893


}

let__10076 := ifres10878
_ = let__10076

tmp10896 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10076)


if True == tmp10896 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__10076)
return
}


} else {
__e.Return(let__10073)
return
}


} else {
__e.Return(let__10069)
return
}


}, 1)

tmp10901 := Call(__e, ns2_1set, symshen_4_5bterm_6, tmp10861)


_ = tmp10901

tmp10902 := MakeNative(func(__e *ControlFlow) {
V1841 := __e.Get(1)
_ = V1841
tmp10912 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1841)
}
__typedArg0 := V1841
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10903 Obj

if True == tmp10912 {
tmp10904 := Call(__e, PrimFunc(symhead), V1841)


let__10085 := tmp10904
_ = let__10085

tmp10905 := Call(__e, PrimFunc(symtail), V1841)


let__10086 := tmp10905
_ = let__10086

tmp10910 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__10085, sym__)
}
__typedArg0 := let__10085
__typedArg1 := sym__
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10906 Obj

if True == tmp10910 {
tmp10907 := Call(__e, PrimFunc(symgensym), symY)


tmp10908 := Call(__e, PrimFunc(symshen_4comb), let__10086, tmp10907)


ifres10906 = tmp10908


} else {
tmp10909 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10906 = tmp10909


}

ifres10903 = ifres10906


} else {
tmp10911 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10903 = tmp10911


}

let__10084 := ifres10903
_ = let__10084

tmp10914 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10084)


if True == tmp10914 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__10084)
return
}


}, 1)

tmp10915 := Call(__e, ns2_1set, symshen_4_5wildcard_6, tmp10902)


_ = tmp10915

tmp10916 := MakeNative(func(__e *ControlFlow) {
V1845 := __e.Get(1)
_ = V1845
tmp10925 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1845)
}
__typedArg0 := V1845
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10917 Obj

if True == tmp10925 {
tmp10918 := Call(__e, PrimFunc(symhead), V1845)


let__10088 := tmp10918
_ = let__10088

tmp10919 := Call(__e, PrimFunc(symtail), V1845)


let__10089 := tmp10919
_ = let__10089

tmp10923 := Call(__e, PrimFunc(symshen_4semicolon_2), let__10088)


var ifres10920 Obj

if True == tmp10923 {
tmp10921 := Call(__e, PrimFunc(symshen_4comb), let__10089, let__10088)


ifres10920 = tmp10921


} else {
tmp10922 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10920 = tmp10922


}

ifres10917 = ifres10920


} else {
tmp10924 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10917 = tmp10924


}

let__10087 := ifres10917
_ = let__10087

tmp10927 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10087)


if True == tmp10927 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__10087)
return
}


}, 1)

tmp10928 := Call(__e, ns2_1set, symshen_4_5sc_6, tmp10916)


_ = tmp10928

tmp10929 := MakeNative(func(__e *ControlFlow) {
V1849 := __e.Get(1)
_ = V1849
V1850 := __e.Get(2)
_ = V1850
tmp10930 := Call(__e, PrimFunc(symgensym), symB)


let__10090 := tmp10930
_ = let__10090

tmp10931 := Call(__e, PrimFunc(symgensym), symL)


let__10091 := tmp10931
_ = let__10091

tmp10932 := Call(__e, PrimFunc(symgensym), symK)


let__10092 := tmp10932
_ = let__10092

tmp10933 := Call(__e, PrimFunc(symgensym), symC)


let__10093 := tmp10933
_ = let__10093

tmp10934 := Call(__e, PrimFunc(symshen_4prolog_1parameters), V1850)


let__10094 := tmp10934
_ = let__10094

tmp10935 := Call(__e, PrimFunc(symshen_4hascut_2), V1850)


let__10095 := tmp10935
_ = let__10095

tmp10936 := Call(__e, PrimFunc(symshen_4prolog_1fbody), V1850, let__10094, let__10090, let__10091, let__10092, let__10093, let__10095)


let__10096 := tmp10936
_ = let__10096

var ifres10937 Obj

if True == let__10095 {
tmp10938 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), Nil)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10939 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10092, tmp10938)
}
__typedArg0 := let__10092
__typedArg1 := tmp10938
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10940 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_7, tmp10939)
}
__typedArg0 := sym_7
__typedArg1 := tmp10939
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10941 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10096, Nil)
}
__typedArg0 := let__10096
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10942 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp10940, tmp10941)
}
__typedArg0 := tmp10940
__typedArg1 := tmp10941
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10943 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10092, tmp10942)
}
__typedArg0 := let__10092
__typedArg1 := tmp10942
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10944 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp10943)
}
__typedArg0 := symlet
__typedArg1 := tmp10943
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

ifres10937 = tmp10944


} else {
ifres10937 = let__10096


}

let__10097 := ifres10937
_ = let__10097

tmp10945 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_6, Nil)
}
__typedArg0 := sym_1_6
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10946 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10093, tmp10945)
}
__typedArg0 := let__10093
__typedArg1 := tmp10945
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10947 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10092, tmp10946)
}
__typedArg0 := let__10092
__typedArg1 := tmp10946
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10948 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10091, tmp10947)
}
__typedArg0 := let__10091
__typedArg1 := tmp10947
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10949 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10090, tmp10948)
}
__typedArg0 := let__10090
__typedArg1 := tmp10948
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10950 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10097, Nil)
}
__typedArg0 := let__10097
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10951 := Call(__e, PrimFunc(symappend), tmp10949, tmp10950)


tmp10952 := Call(__e, PrimFunc(symappend), let__10094, tmp10951)


tmp10953 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1849, tmp10952)
}
__typedArg0 := V1849
__typedArg1 := tmp10952
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10954 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefine, tmp10953)
}
__typedArg0 := symdefine
__typedArg1 := tmp10953
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__10098 := tmp10954
_ = let__10098

__e.Return(let__10098)
return


}, 2)

tmp10955 := Call(__e, ns2_1set, symshen_4horn_1clause_1procedure, tmp10929)


_ = tmp10955

tmp10956 := MakeNative(func(__e *ControlFlow) {
V1862 := __e.Get(1)
_ = V1862
tmp10966 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_b, V1862)
}
__typedArg0 := sym_b
__typedArg1 := V1862
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp10966 {
__e.Return(True)
return
} else {
tmp10964 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1862)
}
__typedArg0 := V1862
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp10964 {
tmp10961 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1862)
}
__typedArg0 := V1862
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10962 := Call(__e, PrimFunc(symshen_4hascut_2), tmp10961)


if True == tmp10962 {
__e.Return(True)
return
} else {
tmp10958 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1862)
}
__typedArg0 := V1862
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10959 := Call(__e, PrimFunc(symshen_4hascut_2), tmp10958)


if True == tmp10959 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


} else {
__e.Return(False)
return
}


}


}, 1)

tmp10967 := Call(__e, ns2_1set, symshen_4hascut_2, tmp10956)


_ = tmp10967

tmp10968 := MakeNative(func(__e *ControlFlow) {
V1867 := __e.Get(1)
_ = V1867
tmp10977 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1867)
}
__typedArg0 := V1867
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10973 Obj

if True == tmp10977 {
tmp10975 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1867)
}
__typedArg0 := V1867
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10976 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10975)
}
__typedArg0 := tmp10975
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10974 Obj

if True == tmp10976 {
ifres10974 = True


} else {
ifres10974 = False


}

ifres10973 = ifres10974


} else {
ifres10973 = False


}

if True == ifres10973 {
tmp10969 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1867)
}
__typedArg0 := V1867
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10970 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10969)
}
__typedArg0 := tmp10969
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10971 := Call(__e, PrimFunc(symlength), tmp10970)


__e.TailApply(PrimFunc(symshen_4parameters), tmp10971)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4prolog_1parameters)
return
}


}, 1)

tmp10978 := Call(__e, ns2_1set, symshen_4prolog_1parameters, tmp10968)


_ = tmp10978

tmp10979 := MakeNative(func(__e *ControlFlow) {
V1888 := __e.Get(1)
_ = V1888
V1889 := __e.Get(2)
_ = V1889
V1890 := __e.Get(3)
_ = V1890
V1891 := __e.Get(4)
_ = V1891
V1892 := __e.Get(5)
_ = V1892
V1893 := __e.Get(6)
_ = V1893
V1894 := __e.Get(7)
_ = V1894
tmp11069 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V1888)
}
__typedArg0 := Nil
__typedArg1 := V1888
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11066 Obj

if True == tmp11069 {
tmp11068 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(True, V1894)
}
__typedArg0 := True
__typedArg1 := V1894
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11067 Obj

if True == tmp11068 {
ifres11067 = True


} else {
ifres11067 = False


}

ifres11066 = ifres11067


} else {
ifres11066 = False


}

if True == ifres11066 {
tmp10980 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1892, Nil)
}
__typedArg0 := V1892
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10981 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1891, tmp10980)
}
__typedArg0 := V1891
__typedArg1 := tmp10980
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4unlock, tmp10981)
}
__typedArg0 := symshen_4unlock
__typedArg1 := tmp10981
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp11064 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1888)
}
__typedArg0 := V1888
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11042 Obj

if True == tmp11064 {
tmp11062 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1888)
}
__typedArg0 := V1888
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11063 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11062)
}
__typedArg0 := tmp11062
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11044 Obj

if True == tmp11063 {
tmp11059 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1888)
}
__typedArg0 := V1888
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11060 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11059)
}
__typedArg0 := tmp11059
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11061 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11060)
}
__typedArg0 := tmp11060
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11046 Obj

if True == tmp11061 {
tmp11055 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1888)
}
__typedArg0 := V1888
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11056 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11055)
}
__typedArg0 := tmp11055
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11057 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11056)
}
__typedArg0 := tmp11056
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11058 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp11057)
}
__typedArg0 := Nil
__typedArg1 := tmp11057
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11048 Obj

if True == tmp11058 {
tmp11053 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1888)
}
__typedArg0 := V1888
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11054 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp11053)
}
__typedArg0 := Nil
__typedArg1 := tmp11053
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11050 Obj

if True == tmp11054 {
tmp11052 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(False, V1894)
}
__typedArg0 := False
__typedArg1 := V1894
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11051 Obj

if True == tmp11052 {
ifres11051 = True


} else {
ifres11051 = False


}

ifres11050 = ifres11051


} else {
ifres11050 = False


}

var ifres11049 Obj

if True == ifres11050 {
ifres11049 = True


} else {
ifres11049 = False


}

ifres11048 = ifres11049


} else {
ifres11048 = False


}

var ifres11047 Obj

if True == ifres11048 {
ifres11047 = True


} else {
ifres11047 = False


}

ifres11046 = ifres11047


} else {
ifres11046 = False


}

var ifres11045 Obj

if True == ifres11046 {
ifres11045 = True


} else {
ifres11045 = False


}

ifres11044 = ifres11045


} else {
ifres11044 = False


}

var ifres11043 Obj

if True == ifres11044 {
ifres11043 = True


} else {
ifres11043 = False


}

ifres11042 = ifres11043


} else {
ifres11042 = False


}

if True == ifres11042 {
tmp10982 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1888)
}
__typedArg0 := V1888
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10983 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10982)
}
__typedArg0 := tmp10982
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10984 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1888)
}
__typedArg0 := V1888
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10985 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10984)
}
__typedArg0 := tmp10984
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10986 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10985)
}
__typedArg0 := tmp10985
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10987 := Call(__e, PrimFunc(symshen_4continue), tmp10983, tmp10986, V1890, V1891, V1892, V1893)


let__10099 := tmp10987
_ = let__10099

tmp10988 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1891, Nil)
}
__typedArg0 := V1891
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10989 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4unlocked_2, tmp10988)
}
__typedArg0 := symshen_4unlocked_2
__typedArg1 := tmp10988
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10990 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1888)
}
__typedArg0 := V1888
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10991 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10990)
}
__typedArg0 := tmp10990
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10992 := Call(__e, PrimFunc(symshen_4compile_1head), symshen_4_7m, tmp10991, V1889, V1890, let__10099)


tmp10993 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(False, Nil)
}
__typedArg0 := False
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10994 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp10992, tmp10993)
}
__typedArg0 := tmp10992
__typedArg1 := tmp10993
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10995 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp10989, tmp10994)
}
__typedArg0 := tmp10989
__typedArg1 := tmp10994
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp10995)
}
__typedArg0 := symif
__typedArg1 := tmp10995
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp11040 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1888)
}
__typedArg0 := V1888
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11025 Obj

if True == tmp11040 {
tmp11038 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1888)
}
__typedArg0 := V1888
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11039 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11038)
}
__typedArg0 := tmp11038
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11027 Obj

if True == tmp11039 {
tmp11035 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1888)
}
__typedArg0 := V1888
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11036 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11035)
}
__typedArg0 := tmp11035
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11037 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11036)
}
__typedArg0 := tmp11036
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11029 Obj

if True == tmp11037 {
tmp11031 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1888)
}
__typedArg0 := V1888
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11032 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11031)
}
__typedArg0 := tmp11031
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11033 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11032)
}
__typedArg0 := tmp11032
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11034 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp11033)
}
__typedArg0 := Nil
__typedArg1 := tmp11033
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11030 Obj

if True == tmp11034 {
ifres11030 = True


} else {
ifres11030 = False


}

ifres11029 = ifres11030


} else {
ifres11029 = False


}

var ifres11028 Obj

if True == ifres11029 {
ifres11028 = True


} else {
ifres11028 = False


}

ifres11027 = ifres11028


} else {
ifres11027 = False


}

var ifres11026 Obj

if True == ifres11027 {
ifres11026 = True


} else {
ifres11026 = False


}

ifres11025 = ifres11026


} else {
ifres11025 = False


}

if True == ifres11025 {
tmp10996 := Call(__e, PrimFunc(symgensym), symC)


let__10100 := tmp10996
_ = let__10100

tmp10997 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1888)
}
__typedArg0 := V1888
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10998 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10997)
}
__typedArg0 := tmp10997
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10999 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1888)
}
__typedArg0 := V1888
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11000 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10999)
}
__typedArg0 := tmp10999
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11001 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11000)
}
__typedArg0 := tmp11000
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11002 := Call(__e, PrimFunc(symshen_4continue), tmp10998, tmp11001, V1890, V1891, V1892, V1893)


let__10101 := tmp11002
_ = let__10101

tmp11003 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1891, Nil)
}
__typedArg0 := V1891
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11004 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4unlocked_2, tmp11003)
}
__typedArg0 := symshen_4unlocked_2
__typedArg1 := tmp11003
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11005 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1888)
}
__typedArg0 := V1888
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11006 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11005)
}
__typedArg0 := tmp11005
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11007 := Call(__e, PrimFunc(symshen_4compile_1head), symshen_4_7m, tmp11006, V1889, V1890, let__10101)


tmp11008 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(False, Nil)
}
__typedArg0 := False
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11009 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11007, tmp11008)
}
__typedArg0 := tmp11007
__typedArg1 := tmp11008
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11010 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11004, tmp11009)
}
__typedArg0 := tmp11004
__typedArg1 := tmp11009
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11011 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp11010)
}
__typedArg0 := symif
__typedArg1 := tmp11010
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11012 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(False, Nil)
}
__typedArg0 := False
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11013 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10100, tmp11012)
}
__typedArg0 := let__10100
__typedArg1 := tmp11012
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11014 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_a, tmp11013)
}
__typedArg0 := sym_a
__typedArg1 := tmp11013
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11015 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1888)
}
__typedArg0 := V1888
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11016 := Call(__e, PrimFunc(symshen_4prolog_1fbody), tmp11015, V1889, V1890, V1891, V1892, V1893, V1894)


tmp11017 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10100, Nil)
}
__typedArg0 := let__10100
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11018 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11016, tmp11017)
}
__typedArg0 := tmp11016
__typedArg1 := tmp11017
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11019 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11014, tmp11018)
}
__typedArg0 := tmp11014
__typedArg1 := tmp11018
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11020 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp11019)
}
__typedArg0 := symif
__typedArg1 := tmp11019
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11021 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11020, Nil)
}
__typedArg0 := tmp11020
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11022 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11011, tmp11021)
}
__typedArg0 := tmp11011
__typedArg1 := tmp11021
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11023 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10100, tmp11022)
}
__typedArg0 := let__10100
__typedArg1 := tmp11022
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp11023)
}
__typedArg0 := symlet
__typedArg1 := tmp11023
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.prolog-fbody"))
}
__typedArg0 := MakeString("implementation error in shen.prolog-fbody")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}


}, 7)

tmp11070 := Call(__e, ns2_1set, symshen_4prolog_1fbody, tmp10979)


_ = tmp11070

tmp11071 := MakeNative(func(__e *ControlFlow) {
V1898 := __e.Get(1)
_ = V1898
V1899 := __e.Get(2)
_ = V1899
tmp11076 := Call(__e, PrimFunc(symshen_4locked_2), V1898)


var ifres11073 Obj

if True == tmp11076 {
tmp11075 := Call(__e, PrimFunc(symshen_4fits_2), V1899, V1898)


var ifres11074 Obj

if True == tmp11075 {
ifres11074 = True


} else {
ifres11074 = False


}

ifres11073 = ifres11074


} else {
ifres11073 = False


}

if True == ifres11073 {
__e.TailApply(PrimFunc(symshen_4openlock), V1898)
return
} else {
__e.Return(False)
return
}


}, 2)

tmp11077 := Call(__e, ns2_1set, symshen_4unlock, tmp11071)


_ = tmp11077

tmp11078 := MakeNative(func(__e *ControlFlow) {
V1900 := __e.Get(1)
_ = V1900
tmp11079 := Call(__e, PrimFunc(symshen_4unlocked_2), V1900)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp11079)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp11079
return Call(__e, PrimFunc(symnot), __typedArg0)
})())
return


}, 1)

tmp11080 := Call(__e, ns2_1set, symshen_4locked_2, tmp11078)


_ = tmp11080

tmp11081 := MakeNative(func(__e *ControlFlow) {
V1901 := __e.Get(1)
_ = V1901
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V1901, MakeInteger(1))
}
__typedArg0 := V1901
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})())
return
}, 1)

tmp11082 := Call(__e, ns2_1set, symshen_4unlocked_2, tmp11081)


_ = tmp11082

tmp11083 := MakeNative(func(__e *ControlFlow) {
V1902 := __e.Get(1)
_ = V1902
tmp11084 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(V1902, MakeInteger(1), True)
}
__typedArg0 := V1902
__typedArg1 := MakeInteger(1)
__typedArg2 := True
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

_ = tmp11084

__e.Return(False)
return


}, 1)

tmp11085 := Call(__e, ns2_1set, symshen_4openlock, tmp11083)


_ = tmp11085

tmp11086 := MakeNative(func(__e *ControlFlow) {
V1903 := __e.Get(1)
_ = V1903
V1904 := __e.Get(2)
_ = V1904
tmp11087 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V1904, MakeInteger(2))
}
__typedArg0 := V1904
__typedArg1 := MakeInteger(2)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V1903, tmp11087)
}
__typedArg0 := V1903
__typedArg1 := tmp11087
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})())
return


}, 2)

tmp11088 := Call(__e, ns2_1set, symshen_4fits_2, tmp11086)


_ = tmp11088

tmp11089 := MakeNative(func(__e *ControlFlow) {
V1907 := __e.Get(1)
_ = V1907
V1908 := __e.Get(2)
_ = V1908
V1909 := __e.Get(3)
_ = V1909
V1910 := __e.Get(4)
_ = V1910
tmp11090 := Call(__e, PrimFunc(symthaw), V1910)


let__10102 := tmp11090
_ = let__10102

tmp11095 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__10102, False)
}
__typedArg0 := let__10102
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11092 Obj

if True == tmp11095 {
tmp11094 := Call(__e, PrimFunc(symshen_4unlocked_2), V1908)


var ifres11093 Obj

if True == tmp11094 {
ifres11093 = True


} else {
ifres11093 = False


}

ifres11092 = ifres11093


} else {
ifres11092 = False


}

if True == ifres11092 {
__e.TailApply(PrimFunc(symshen_4lock), V1909, V1908)
return
} else {
__e.Return(let__10102)
return
}


}, 4)

tmp11096 := Call(__e, ns2_1set, symshen_4cut, tmp11089)


_ = tmp11096

tmp11097 := MakeNative(func(__e *ControlFlow) {
V1912 := __e.Get(1)
_ = V1912
V1913 := __e.Get(2)
_ = V1913
tmp11098 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(V1913, MakeInteger(1), False)
}
__typedArg0 := V1913
__typedArg1 := MakeInteger(1)
__typedArg2 := False
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

let__10103 := tmp11098
_ = let__10103

tmp11099 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(V1913, MakeInteger(2), V1912)
}
__typedArg0 := V1913
__typedArg1 := MakeInteger(2)
__typedArg2 := V1912
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

let__10104 := tmp11099
_ = let__10104

__e.Return(False)
return


}, 2)

tmp11100 := Call(__e, ns2_1set, symshen_4lock, tmp11097)


_ = tmp11100

tmp11101 := MakeNative(func(__e *ControlFlow) {
V1916 := __e.Get(1)
_ = V1916
V1917 := __e.Get(2)
_ = V1917
V1918 := __e.Get(3)
_ = V1918
V1919 := __e.Get(4)
_ = V1919
V1920 := __e.Get(5)
_ = V1920
V1921 := __e.Get(6)
_ = V1921
tmp11102 := Call(__e, PrimFunc(symshen_4extract_1vars), V1916)


let__10105 := tmp11102
_ = let__10105

tmp11103 := Call(__e, PrimFunc(symshen_4extract_1free_1vars), V1917)


let__10106 := tmp11103
_ = let__10106

tmp11104 := Call(__e, PrimFunc(symdifference), let__10106, let__10105)


let__10107 := tmp11104
_ = let__10107

tmp11105 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4incinfs, Nil)
}
__typedArg0 := symshen_4incinfs
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11106 := Call(__e, PrimFunc(symshen_4compile_1body), V1917, V1918, V1919, V1920, V1921)


tmp11107 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11106, Nil)
}
__typedArg0 := tmp11106
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11108 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11105, tmp11107)
}
__typedArg0 := tmp11105
__typedArg1 := tmp11107
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11109 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdo, tmp11108)
}
__typedArg0 := symdo
__typedArg1 := tmp11108
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__10108 := tmp11109
_ = let__10108

__e.TailApply(PrimFunc(symshen_4stpart), let__10107, let__10108, V1918)
return


}, 6)

tmp11110 := Call(__e, ns2_1set, symshen_4continue, tmp11101)


_ = tmp11110

tmp11111 := MakeNative(func(__e *ControlFlow) {
V1928 := __e.Get(1)
_ = V1928
tmp11146 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1928)
}
__typedArg0 := V1928
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11127 Obj

if True == tmp11146 {
tmp11144 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1928)
}
__typedArg0 := V1928
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11145 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlambda, tmp11144)
}
__typedArg0 := symlambda
__typedArg1 := tmp11144
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11129 Obj

if True == tmp11145 {
tmp11142 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1928)
}
__typedArg0 := V1928
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11143 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11142)
}
__typedArg0 := tmp11142
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11131 Obj

if True == tmp11143 {
tmp11139 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1928)
}
__typedArg0 := V1928
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11140 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11139)
}
__typedArg0 := tmp11139
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11141 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11140)
}
__typedArg0 := tmp11140
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11133 Obj

if True == tmp11141 {
tmp11135 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1928)
}
__typedArg0 := V1928
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11136 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11135)
}
__typedArg0 := tmp11135
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11137 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11136)
}
__typedArg0 := tmp11136
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11138 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp11137)
}
__typedArg0 := Nil
__typedArg1 := tmp11137
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11134 Obj

if True == tmp11138 {
ifres11134 = True


} else {
ifres11134 = False


}

ifres11133 = ifres11134


} else {
ifres11133 = False


}

var ifres11132 Obj

if True == ifres11133 {
ifres11132 = True


} else {
ifres11132 = False


}

ifres11131 = ifres11132


} else {
ifres11131 = False


}

var ifres11130 Obj

if True == ifres11131 {
ifres11130 = True


} else {
ifres11130 = False


}

ifres11129 = ifres11130


} else {
ifres11129 = False


}

var ifres11128 Obj

if True == ifres11129 {
ifres11128 = True


} else {
ifres11128 = False


}

ifres11127 = ifres11128


} else {
ifres11127 = False


}

if True == ifres11127 {
tmp11112 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1928)
}
__typedArg0 := V1928
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11113 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11112)
}
__typedArg0 := tmp11112
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11114 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1928)
}
__typedArg0 := V1928
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11115 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11114)
}
__typedArg0 := tmp11114
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11116 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11115)
}
__typedArg0 := tmp11115
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11117 := Call(__e, PrimFunc(symshen_4extract_1free_1vars), tmp11116)


__e.TailApply(PrimFunc(symremove), tmp11113, tmp11117)
return


} else {
tmp11125 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1928)
}
__typedArg0 := V1928
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp11125 {
tmp11118 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1928)
}
__typedArg0 := V1928
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11119 := Call(__e, PrimFunc(symshen_4extract_1free_1vars), tmp11118)


tmp11120 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1928)
}
__typedArg0 := V1928
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11121 := Call(__e, PrimFunc(symshen_4extract_1free_1vars), tmp11120)


__e.TailApply(PrimFunc(symunion), tmp11119, tmp11121)
return


} else {
tmp11123 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(V1928)
}
__typedArg0 := V1928
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

if True == tmp11123 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1928, Nil)
}
__typedArg0 := V1928
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return
} else {
__e.Return(Nil)
return
}


}


}


}, 1)

tmp11147 := Call(__e, ns2_1set, symshen_4extract_1free_1vars, tmp11111)


_ = tmp11147

tmp11148 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__self3 := __e.Get(3)
__self4 := __e.Get(4)
__self5 := __e.Get(5)
__selftop:
V1945 := __self1
_ = V1945
V1946 := __self2
_ = V1946
V1947 := __self3
_ = V1947
V1948 := __self4
_ = V1948
V1949 := __self5
_ = V1949
tmp11182 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V1945)
}
__typedArg0 := Nil
__typedArg1 := V1945
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp11182 {
tmp11149 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1949, Nil)
}
__typedArg0 := V1949
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symthaw, tmp11149)
}
__typedArg0 := symthaw
__typedArg1 := tmp11149
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp11180 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1945)
}
__typedArg0 := V1945
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11176 Obj

if True == tmp11180 {
tmp11178 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1945)
}
__typedArg0 := V1945
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11179 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_b, tmp11178)
}
__typedArg0 := sym_b
__typedArg1 := tmp11178
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11177 Obj

if True == tmp11179 {
ifres11177 = True


} else {
ifres11177 = False


}

ifres11176 = ifres11177


} else {
ifres11176 = False


}

if True == ifres11176 {
tmp11150 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4cut, Nil)
}
__typedArg0 := symshen_4cut
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11151 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1945)
}
__typedArg0 := V1945
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11152 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11150, tmp11151)
}
__typedArg0 := tmp11150
__typedArg1 := tmp11151
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

if PrimFunc(symshen_4compile_1body) == __self {
__self1, __self2, __self3, __self4, __self5 = tmp11152, V1946, V1947, V1948, V1949
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4compile_1body), tmp11152, V1946, V1947, V1948, V1949)
return


} else {
tmp11174 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1945)
}
__typedArg0 := V1945
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11170 Obj

if True == tmp11174 {
tmp11172 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1945)
}
__typedArg0 := V1945
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11173 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp11172)
}
__typedArg0 := Nil
__typedArg1 := tmp11172
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11171 Obj

if True == tmp11173 {
ifres11171 = True


} else {
ifres11171 = False


}

ifres11170 = ifres11171


} else {
ifres11170 = False


}

if True == ifres11170 {
tmp11153 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1945)
}
__typedArg0 := V1945
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11154 := Call(__e, PrimFunc(symshen_4deref_1calls), tmp11153, V1946)


tmp11155 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1949, Nil)
}
__typedArg0 := V1949
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11156 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1948, tmp11155)
}
__typedArg0 := V1948
__typedArg1 := tmp11155
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11157 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1947, tmp11156)
}
__typedArg0 := V1947
__typedArg1 := tmp11156
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11158 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1946, tmp11157)
}
__typedArg0 := V1946
__typedArg1 := tmp11157
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symappend), tmp11154, tmp11158)
return


} else {
tmp11168 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1945)
}
__typedArg0 := V1945
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp11168 {
tmp11159 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1945)
}
__typedArg0 := V1945
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11160 := Call(__e, PrimFunc(symshen_4deref_1calls), tmp11159, V1946)


let__10109 := tmp11160
_ = let__10109

tmp11161 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1945)
}
__typedArg0 := V1945
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11162 := Call(__e, PrimFunc(symshen_4freeze_1literals), tmp11161, V1946, V1947, V1948, V1949)


tmp11163 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11162, Nil)
}
__typedArg0 := tmp11162
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11164 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1948, tmp11163)
}
__typedArg0 := V1948
__typedArg1 := tmp11163
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11165 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1947, tmp11164)
}
__typedArg0 := V1947
__typedArg1 := tmp11164
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11166 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1946, tmp11165)
}
__typedArg0 := V1946
__typedArg1 := tmp11165
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symappend), let__10109, tmp11166)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.compile-fbody"))
}
__typedArg0 := MakeString("implementation error in shen.compile-fbody")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}


}


}, 5)

tmp11183 := Call(__e, ns2_1set, symshen_4compile_1body, tmp11148)


_ = tmp11183

tmp11184 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__self3 := __e.Get(3)
__self4 := __e.Get(4)
__self5 := __e.Get(5)
__selftop:
V1967 := __self1
_ = V1967
V1968 := __self2
_ = V1968
V1969 := __self3
_ = V1969
V1970 := __self4
_ = V1970
V1971 := __self5
_ = V1971
tmp11207 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V1967)
}
__typedArg0 := Nil
__typedArg1 := V1967
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp11207 {
__e.Return(V1971)
return
} else {
tmp11205 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1967)
}
__typedArg0 := V1967
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11201 Obj

if True == tmp11205 {
tmp11203 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1967)
}
__typedArg0 := V1967
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11204 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_b, tmp11203)
}
__typedArg0 := sym_b
__typedArg1 := tmp11203
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11202 Obj

if True == tmp11204 {
ifres11202 = True


} else {
ifres11202 = False


}

ifres11201 = ifres11202


} else {
ifres11201 = False


}

if True == ifres11201 {
tmp11185 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4cut, Nil)
}
__typedArg0 := symshen_4cut
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11186 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1967)
}
__typedArg0 := V1967
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11187 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11185, tmp11186)
}
__typedArg0 := tmp11185
__typedArg1 := tmp11186
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

if PrimFunc(symshen_4freeze_1literals) == __self {
__self1, __self2, __self3, __self4, __self5 = tmp11187, V1968, V1969, V1970, V1971
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4freeze_1literals), tmp11187, V1968, V1969, V1970, V1971)
return


} else {
tmp11199 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1967)
}
__typedArg0 := V1967
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp11199 {
tmp11188 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1967)
}
__typedArg0 := V1967
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11189 := Call(__e, PrimFunc(symshen_4deref_1calls), tmp11188, V1968)


let__10110 := tmp11189
_ = let__10110

tmp11190 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1967)
}
__typedArg0 := V1967
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11191 := Call(__e, PrimFunc(symshen_4freeze_1literals), tmp11190, V1968, V1969, V1970, V1971)


tmp11192 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11191, Nil)
}
__typedArg0 := tmp11191
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11193 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1970, tmp11192)
}
__typedArg0 := V1970
__typedArg1 := tmp11192
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11194 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1969, tmp11193)
}
__typedArg0 := V1969
__typedArg1 := tmp11193
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11195 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1968, tmp11194)
}
__typedArg0 := V1968
__typedArg1 := tmp11194
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11196 := Call(__e, PrimFunc(symappend), let__10110, tmp11195)


tmp11197 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11196, Nil)
}
__typedArg0 := tmp11196
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfreeze, tmp11197)
}
__typedArg0 := symfreeze
__typedArg1 := tmp11197
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.freeze-literals"))
}
__typedArg0 := MakeString("implementation error in shen.freeze-literals")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}


}, 5)

tmp11208 := Call(__e, ns2_1set, symshen_4freeze_1literals, tmp11184)


_ = tmp11208

tmp11209 := MakeNative(func(__e *ControlFlow) {
V1977 := __e.Get(1)
_ = V1977
V1978 := __e.Get(2)
_ = V1978
tmp11224 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1977)
}
__typedArg0 := V1977
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11220 Obj

if True == tmp11224 {
tmp11222 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1977)
}
__typedArg0 := V1977
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11223 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symfork, tmp11222)
}
__typedArg0 := symfork
__typedArg1 := tmp11222
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11221 Obj

if True == tmp11223 {
ifres11221 = True


} else {
ifres11221 = False


}

ifres11220 = ifres11221


} else {
ifres11220 = False


}

if True == ifres11220 {
tmp11210 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1977)
}
__typedArg0 := V1977
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11211 := Call(__e, PrimFunc(symshen_4deref_1forked_1literals), tmp11210, V1978)


tmp11212 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11211, Nil)
}
__typedArg0 := tmp11211
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfork, tmp11212)
}
__typedArg0 := symfork
__typedArg1 := tmp11212
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp11218 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1977)
}
__typedArg0 := V1977
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp11218 {
tmp11213 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1977)
}
__typedArg0 := V1977
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11214 := MakeNative(func(__e *ControlFlow) {
Z1979 := __e.Get(1)
_ = Z1979
__e.TailApply(PrimFunc(symshen_4function_1calls), Z1979, V1978)
return
}, 1)

tmp11215 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1977)
}
__typedArg0 := V1977
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11216 := Call(__e, PrimFunc(symmap), tmp11214, tmp11215)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11213, tmp11216)
}
__typedArg0 := tmp11213
__typedArg1 := tmp11216
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.deref-calls"))
}
__typedArg0 := MakeString("implementation error in shen.deref-calls")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 2)

tmp11225 := Call(__e, ns2_1set, symshen_4deref_1calls, tmp11209)


_ = tmp11225

tmp11226 := MakeNative(func(__e *ControlFlow) {
V1986 := __e.Get(1)
_ = V1986
V1987 := __e.Get(2)
_ = V1987
tmp11236 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V1986)
}
__typedArg0 := Nil
__typedArg1 := V1986
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp11236 {
__e.Return(Nil)
return
} else {
tmp11234 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1986)
}
__typedArg0 := V1986
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp11234 {
tmp11227 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1986)
}
__typedArg0 := V1986
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11228 := Call(__e, PrimFunc(symshen_4deref_1calls), tmp11227, V1987)


tmp11229 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1986)
}
__typedArg0 := V1986
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11230 := Call(__e, PrimFunc(symshen_4deref_1forked_1literals), tmp11229, V1987)


tmp11231 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11230, Nil)
}
__typedArg0 := tmp11230
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11232 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11228, tmp11231)
}
__typedArg0 := tmp11228
__typedArg1 := tmp11231
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp11232)
}
__typedArg0 := symcons
__typedArg1 := tmp11232
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("fork requires a list of literals\n"))
}
__typedArg0 := MakeString("fork requires a list of literals\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 2)

tmp11237 := Call(__e, ns2_1set, symshen_4deref_1forked_1literals, tmp11226)


_ = tmp11237

tmp11238 := MakeNative(func(__e *ControlFlow) {
V1990 := __e.Get(1)
_ = V1990
V1991 := __e.Get(2)
_ = V1991
tmp11270 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1990)
}
__typedArg0 := V1990
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11251 Obj

if True == tmp11270 {
tmp11268 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1990)
}
__typedArg0 := V1990
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11269 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcons, tmp11268)
}
__typedArg0 := symcons
__typedArg1 := tmp11268
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11253 Obj

if True == tmp11269 {
tmp11266 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1990)
}
__typedArg0 := V1990
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11267 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11266)
}
__typedArg0 := tmp11266
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11255 Obj

if True == tmp11267 {
tmp11263 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1990)
}
__typedArg0 := V1990
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11264 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11263)
}
__typedArg0 := tmp11263
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11265 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11264)
}
__typedArg0 := tmp11264
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11257 Obj

if True == tmp11265 {
tmp11259 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1990)
}
__typedArg0 := V1990
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11260 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11259)
}
__typedArg0 := tmp11259
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11261 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11260)
}
__typedArg0 := tmp11260
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11262 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp11261)
}
__typedArg0 := Nil
__typedArg1 := tmp11261
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11258 Obj

if True == tmp11262 {
ifres11258 = True


} else {
ifres11258 = False


}

ifres11257 = ifres11258


} else {
ifres11257 = False


}

var ifres11256 Obj

if True == ifres11257 {
ifres11256 = True


} else {
ifres11256 = False


}

ifres11255 = ifres11256


} else {
ifres11255 = False


}

var ifres11254 Obj

if True == ifres11255 {
ifres11254 = True


} else {
ifres11254 = False


}

ifres11253 = ifres11254


} else {
ifres11253 = False


}

var ifres11252 Obj

if True == ifres11253 {
ifres11252 = True


} else {
ifres11252 = False


}

ifres11251 = ifres11252


} else {
ifres11251 = False


}

if True == ifres11251 {
tmp11239 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1990)
}
__typedArg0 := V1990
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11240 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11239)
}
__typedArg0 := tmp11239
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11241 := Call(__e, PrimFunc(symshen_4function_1calls), tmp11240, V1991)


tmp11242 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1990)
}
__typedArg0 := V1990
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11243 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11242)
}
__typedArg0 := tmp11242
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11244 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11243)
}
__typedArg0 := tmp11243
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11245 := Call(__e, PrimFunc(symshen_4function_1calls), tmp11244, V1991)


tmp11246 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11245, Nil)
}
__typedArg0 := tmp11245
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11247 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11241, tmp11246)
}
__typedArg0 := tmp11241
__typedArg1 := tmp11246
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp11247)
}
__typedArg0 := symcons
__typedArg1 := tmp11247
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp11249 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1990)
}
__typedArg0 := V1990
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp11249 {
__e.TailApply(PrimFunc(symshen_4deref_1terms), V1990, V1991, Nil)
return
} else {
__e.Return(V1990)
return
}


}


}, 2)

tmp11271 := Call(__e, ns2_1set, symshen_4function_1calls, tmp11238)


_ = tmp11271

tmp11272 := MakeNative(func(__e *ControlFlow) {
V2000 := __e.Get(1)
_ = V2000
V2001 := __e.Get(2)
_ = V2001
V2002 := __e.Get(3)
_ = V2002
tmp11366 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2000)
}
__typedArg0 := V2000
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11353 Obj

if True == tmp11366 {
tmp11364 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2000)
}
__typedArg0 := V2000
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11365 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(0), tmp11364)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp11364
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11355 Obj

if True == tmp11365 {
tmp11362 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2000)
}
__typedArg0 := V2000
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11363 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11362)
}
__typedArg0 := tmp11362
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11357 Obj

if True == tmp11363 {
tmp11359 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2000)
}
__typedArg0 := V2000
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11360 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11359)
}
__typedArg0 := tmp11359
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11361 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp11360)
}
__typedArg0 := Nil
__typedArg1 := tmp11360
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11358 Obj

if True == tmp11361 {
ifres11358 = True


} else {
ifres11358 = False


}

ifres11357 = ifres11358


} else {
ifres11357 = False


}

var ifres11356 Obj

if True == ifres11357 {
ifres11356 = True


} else {
ifres11356 = False


}

ifres11355 = ifres11356


} else {
ifres11355 = False


}

var ifres11354 Obj

if True == ifres11355 {
ifres11354 = True


} else {
ifres11354 = False


}

ifres11353 = ifres11354


} else {
ifres11353 = False


}

if True == ifres11353 {
tmp11279 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2000)
}
__typedArg0 := V2000
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11280 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11279)
}
__typedArg0 := tmp11279
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11281 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(tmp11280)
}
__typedArg0 := tmp11280
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

if True == tmp11281 {
tmp11273 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2000)
}
__typedArg0 := V2000
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11273)
}
__typedArg0 := tmp11273
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return


} else {
tmp11274 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2000)
}
__typedArg0 := V2000
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11275 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11274)
}
__typedArg0 := tmp11274
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11276 := Call(__e, PrimFunc(symshen_4app), tmp11275, MakeString("\n"), symshen_4s)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("attempt to optimise a non-variable "))
__typedS1, __typedOK1 := TypedString(tmp11276)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("attempt to optimise a non-variable ")
__typedArg1 := tmp11276
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("attempt to optimise a non-variable "))
__typedS1, __typedOK1 := TypedString(tmp11276)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("attempt to optimise a non-variable ")
__typedArg1 := tmp11276
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}


} else {
tmp11351 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2000)
}
__typedArg0 := V2000
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11338 Obj

if True == tmp11351 {
tmp11349 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2000)
}
__typedArg0 := V2000
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11350 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(1), tmp11349)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp11349
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11340 Obj

if True == tmp11350 {
tmp11347 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2000)
}
__typedArg0 := V2000
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11348 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11347)
}
__typedArg0 := tmp11347
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11342 Obj

if True == tmp11348 {
tmp11344 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2000)
}
__typedArg0 := V2000
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11345 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11344)
}
__typedArg0 := tmp11344
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11346 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp11345)
}
__typedArg0 := Nil
__typedArg1 := tmp11345
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11343 Obj

if True == tmp11346 {
ifres11343 = True


} else {
ifres11343 = False


}

ifres11342 = ifres11343


} else {
ifres11342 = False


}

var ifres11341 Obj

if True == ifres11342 {
ifres11341 = True


} else {
ifres11341 = False


}

ifres11340 = ifres11341


} else {
ifres11340 = False


}

var ifres11339 Obj

if True == ifres11340 {
ifres11339 = True


} else {
ifres11339 = False


}

ifres11338 = ifres11339


} else {
ifres11338 = False


}

if True == ifres11338 {
tmp11291 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2000)
}
__typedArg0 := V2000
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11292 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11291)
}
__typedArg0 := tmp11291
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11293 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(tmp11292)
}
__typedArg0 := tmp11292
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

if True == tmp11293 {
tmp11282 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2000)
}
__typedArg0 := V2000
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11283 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11282)
}
__typedArg0 := tmp11282
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11284 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V2001, Nil)
}
__typedArg0 := V2001
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11285 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11283, tmp11284)
}
__typedArg0 := tmp11283
__typedArg1 := tmp11284
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4lazyderef, tmp11285)
}
__typedArg0 := symshen_4lazyderef
__typedArg1 := tmp11285
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp11286 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2000)
}
__typedArg0 := V2000
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11287 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11286)
}
__typedArg0 := tmp11286
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11288 := Call(__e, PrimFunc(symshen_4app), tmp11287, MakeString("\n"), symshen_4s)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("attempt to optimise a non-variable "))
__typedS1, __typedOK1 := TypedString(tmp11288)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("attempt to optimise a non-variable ")
__typedArg1 := tmp11288
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("attempt to optimise a non-variable "))
__typedS1, __typedOK1 := TypedString(tmp11288)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("attempt to optimise a non-variable ")
__typedArg1 := tmp11288
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}


} else {
tmp11335 := Call(__e, PrimFunc(symelement_2), V2000, V2002)


tmp11336 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp11335)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp11335
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres11332 Obj

if True == tmp11336 {
tmp11334 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(V2000)
}
__typedArg0 := V2000
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

var ifres11333 Obj

if True == tmp11334 {
ifres11333 = True


} else {
ifres11333 = False


}

ifres11332 = ifres11333


} else {
ifres11332 = False


}

if True == ifres11332 {
tmp11294 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V2001, Nil)
}
__typedArg0 := V2001
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11295 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V2000, tmp11294)
}
__typedArg0 := V2000
__typedArg1 := tmp11294
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4deref, tmp11295)
}
__typedArg0 := symshen_4deref
__typedArg1 := tmp11295
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp11330 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2000)
}
__typedArg0 := V2000
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11311 Obj

if True == tmp11330 {
tmp11328 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2000)
}
__typedArg0 := V2000
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11329 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlambda, tmp11328)
}
__typedArg0 := symlambda
__typedArg1 := tmp11328
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11313 Obj

if True == tmp11329 {
tmp11326 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2000)
}
__typedArg0 := V2000
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11327 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11326)
}
__typedArg0 := tmp11326
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11315 Obj

if True == tmp11327 {
tmp11323 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2000)
}
__typedArg0 := V2000
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11324 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11323)
}
__typedArg0 := tmp11323
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11325 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11324)
}
__typedArg0 := tmp11324
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11317 Obj

if True == tmp11325 {
tmp11319 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2000)
}
__typedArg0 := V2000
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11320 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11319)
}
__typedArg0 := tmp11319
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11321 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11320)
}
__typedArg0 := tmp11320
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11322 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp11321)
}
__typedArg0 := Nil
__typedArg1 := tmp11321
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11318 Obj

if True == tmp11322 {
ifres11318 = True


} else {
ifres11318 = False


}

ifres11317 = ifres11318


} else {
ifres11317 = False


}

var ifres11316 Obj

if True == ifres11317 {
ifres11316 = True


} else {
ifres11316 = False


}

ifres11315 = ifres11316


} else {
ifres11315 = False


}

var ifres11314 Obj

if True == ifres11315 {
ifres11314 = True


} else {
ifres11314 = False


}

ifres11313 = ifres11314


} else {
ifres11313 = False


}

var ifres11312 Obj

if True == ifres11313 {
ifres11312 = True


} else {
ifres11312 = False


}

ifres11311 = ifres11312


} else {
ifres11311 = False


}

if True == ifres11311 {
tmp11296 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2000)
}
__typedArg0 := V2000
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11297 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11296)
}
__typedArg0 := tmp11296
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11298 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2000)
}
__typedArg0 := V2000
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11299 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11298)
}
__typedArg0 := tmp11298
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11300 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11299)
}
__typedArg0 := tmp11299
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11301 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2000)
}
__typedArg0 := V2000
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11302 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11301)
}
__typedArg0 := tmp11301
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11303 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11302, V2002)
}
__typedArg0 := tmp11302
__typedArg1 := V2002
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11304 := Call(__e, PrimFunc(symshen_4deref_1terms), tmp11300, V2001, tmp11303)


tmp11305 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11304, Nil)
}
__typedArg0 := tmp11304
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11306 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11297, tmp11305)
}
__typedArg0 := tmp11297
__typedArg1 := tmp11305
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp11306)
}
__typedArg0 := symlambda
__typedArg1 := tmp11306
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp11309 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2000)
}
__typedArg0 := V2000
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp11309 {
tmp11307 := MakeNative(func(__e *ControlFlow) {
Z2003 := __e.Get(1)
_ = Z2003
__e.TailApply(PrimFunc(symshen_4deref_1terms), Z2003, V2001, V2002)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp11307, V2000)
return


} else {
__e.Return(V2000)
return
}


}


}


}


}


}, 3)

tmp11367 := Call(__e, ns2_1set, symshen_4deref_1terms, tmp11272)


_ = tmp11367

tmp11368 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__self3 := __e.Get(3)
__self4 := __e.Get(4)
__self5 := __e.Get(5)
__selftop:
V2021 := __self1
_ = V2021
V2022 := __self2
_ = V2022
V2023 := __self3
_ = V2023
V2024 := __self4
_ = V2024
V2025 := __self5
_ = V2025
tmp11544 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V2022)
}
__typedArg0 := Nil
__typedArg1 := V2022
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11541 Obj

if True == tmp11544 {
tmp11543 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V2023)
}
__typedArg0 := Nil
__typedArg1 := V2023
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11542 Obj

if True == tmp11543 {
ifres11542 = True


} else {
ifres11542 = False


}

ifres11541 = ifres11542


} else {
ifres11541 = False


}

if True == ifres11541 {
__e.Return(V2025)
return
} else {
tmp11539 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11519 Obj

if True == tmp11539 {
tmp11537 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11538 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11537)
}
__typedArg0 := tmp11537
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11521 Obj

if True == tmp11538 {
tmp11534 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11535 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11534)
}
__typedArg0 := tmp11534
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11536 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4_7m, tmp11535)
}
__typedArg0 := symshen_4_7m
__typedArg1 := tmp11535
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11523 Obj

if True == tmp11536 {
tmp11531 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11532 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11531)
}
__typedArg0 := tmp11531
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11533 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11532)
}
__typedArg0 := tmp11532
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11525 Obj

if True == tmp11533 {
tmp11527 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11528 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11527)
}
__typedArg0 := tmp11527
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11529 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11528)
}
__typedArg0 := tmp11528
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11530 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp11529)
}
__typedArg0 := Nil
__typedArg1 := tmp11529
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11526 Obj

if True == tmp11530 {
ifres11526 = True


} else {
ifres11526 = False


}

ifres11525 = ifres11526


} else {
ifres11525 = False


}

var ifres11524 Obj

if True == ifres11525 {
ifres11524 = True


} else {
ifres11524 = False


}

ifres11523 = ifres11524


} else {
ifres11523 = False


}

var ifres11522 Obj

if True == ifres11523 {
ifres11522 = True


} else {
ifres11522 = False


}

ifres11521 = ifres11522


} else {
ifres11521 = False


}

var ifres11520 Obj

if True == ifres11521 {
ifres11520 = True


} else {
ifres11520 = False


}

ifres11519 = ifres11520


} else {
ifres11519 = False


}

if True == ifres11519 {
tmp11369 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11370 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11369)
}
__typedArg0 := tmp11369
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11371 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11370)
}
__typedArg0 := tmp11370
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11372 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11373 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V2021, tmp11372)
}
__typedArg0 := V2021
__typedArg1 := tmp11372
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11374 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11371, tmp11373)
}
__typedArg0 := tmp11371
__typedArg1 := tmp11373
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11375 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4_7m, tmp11374)
}
__typedArg0 := symshen_4_7m
__typedArg1 := tmp11374
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

if PrimFunc(symshen_4compile_1head) == __self {
__self1, __self2, __self3, __self4, __self5 = V2021, tmp11375, V2023, V2024, V2025
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4compile_1head), V2021, tmp11375, V2023, V2024, V2025)
return


} else {
tmp11517 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11497 Obj

if True == tmp11517 {
tmp11515 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11516 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11515)
}
__typedArg0 := tmp11515
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11499 Obj

if True == tmp11516 {
tmp11512 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11513 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11512)
}
__typedArg0 := tmp11512
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11514 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4_1m, tmp11513)
}
__typedArg0 := symshen_4_1m
__typedArg1 := tmp11513
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11501 Obj

if True == tmp11514 {
tmp11509 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11510 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11509)
}
__typedArg0 := tmp11509
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11511 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11510)
}
__typedArg0 := tmp11510
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11503 Obj

if True == tmp11511 {
tmp11505 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11506 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11505)
}
__typedArg0 := tmp11505
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11507 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11506)
}
__typedArg0 := tmp11506
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11508 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp11507)
}
__typedArg0 := Nil
__typedArg1 := tmp11507
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11504 Obj

if True == tmp11508 {
ifres11504 = True


} else {
ifres11504 = False


}

ifres11503 = ifres11504


} else {
ifres11503 = False


}

var ifres11502 Obj

if True == ifres11503 {
ifres11502 = True


} else {
ifres11502 = False


}

ifres11501 = ifres11502


} else {
ifres11501 = False


}

var ifres11500 Obj

if True == ifres11501 {
ifres11500 = True


} else {
ifres11500 = False


}

ifres11499 = ifres11500


} else {
ifres11499 = False


}

var ifres11498 Obj

if True == ifres11499 {
ifres11498 = True


} else {
ifres11498 = False


}

ifres11497 = ifres11498


} else {
ifres11497 = False


}

if True == ifres11497 {
tmp11376 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11377 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11376)
}
__typedArg0 := tmp11376
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11378 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11377)
}
__typedArg0 := tmp11377
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11379 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11380 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V2021, tmp11379)
}
__typedArg0 := V2021
__typedArg1 := tmp11379
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11381 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11378, tmp11380)
}
__typedArg0 := tmp11378
__typedArg1 := tmp11380
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11382 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4_1m, tmp11381)
}
__typedArg0 := symshen_4_1m
__typedArg1 := tmp11381
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

if PrimFunc(symshen_4compile_1head) == __self {
__self1, __self2, __self3, __self4, __self5 = V2021, tmp11382, V2023, V2024, V2025
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4compile_1head), V2021, tmp11382, V2023, V2024, V2025)
return


} else {
tmp11495 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11491 Obj

if True == tmp11495 {
tmp11493 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11494 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4_1m, tmp11493)
}
__typedArg0 := symshen_4_1m
__typedArg1 := tmp11493
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11492 Obj

if True == tmp11494 {
ifres11492 = True


} else {
ifres11492 = False


}

ifres11491 = ifres11492


} else {
ifres11491 = False


}

if True == ifres11491 {
tmp11383 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4compile_1head) == __self {
__self1, __self2, __self3, __self4, __self5 = symshen_4_1m, tmp11383, V2023, V2024, V2025
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4compile_1head), symshen_4_1m, tmp11383, V2023, V2024, V2025)
return


} else {
tmp11489 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11485 Obj

if True == tmp11489 {
tmp11487 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11488 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4_7m, tmp11487)
}
__typedArg0 := symshen_4_7m
__typedArg1 := tmp11487
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11486 Obj

if True == tmp11488 {
ifres11486 = True


} else {
ifres11486 = False


}

ifres11485 = ifres11486


} else {
ifres11485 = False


}

if True == ifres11485 {
tmp11384 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4compile_1head) == __self {
__self1, __self2, __self3, __self4, __self5 = symshen_4_7m, tmp11384, V2023, V2024, V2025
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4compile_1head), symshen_4_7m, tmp11384, V2023, V2024, V2025)
return


} else {
tmp11483 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11476 Obj

if True == tmp11483 {
tmp11482 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2023)
}
__typedArg0 := V2023
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11478 Obj

if True == tmp11482 {
tmp11480 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11481 := Call(__e, PrimFunc(symshen_4wildcard_2), tmp11480)


var ifres11479 Obj

if True == tmp11481 {
ifres11479 = True


} else {
ifres11479 = False


}

ifres11478 = ifres11479


} else {
ifres11478 = False


}

var ifres11477 Obj

if True == ifres11478 {
ifres11477 = True


} else {
ifres11477 = False


}

ifres11476 = ifres11477


} else {
ifres11476 = False


}

if True == ifres11476 {
tmp11385 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11386 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2023)
}
__typedArg0 := V2023
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4compile_1head) == __self {
__self1, __self2, __self3, __self4, __self5 = V2021, tmp11385, tmp11386, V2024, V2025
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4compile_1head), V2021, tmp11385, tmp11386, V2024, V2025)
return


} else {
tmp11474 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11470 Obj

if True == tmp11474 {
tmp11472 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11473 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(tmp11472)
}
__typedArg0 := tmp11472
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

var ifres11471 Obj

if True == tmp11473 {
ifres11471 = True


} else {
ifres11471 = False


}

ifres11470 = ifres11471


} else {
ifres11470 = False


}

if True == ifres11470 {
__e.TailApply(PrimFunc(symshen_4variable_1case), V2021, V2022, V2023, V2024, V2025)
return
} else {
tmp11468 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4_1m, V2021)
}
__typedArg0 := symshen_4_1m
__typedArg1 := V2021
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11461 Obj

if True == tmp11468 {
tmp11467 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11463 Obj

if True == tmp11467 {
tmp11465 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11466 := Call(__e, PrimFunc(symatom_2), tmp11465)


var ifres11464 Obj

if True == tmp11466 {
ifres11464 = True


} else {
ifres11464 = False


}

ifres11463 = ifres11464


} else {
ifres11463 = False


}

var ifres11462 Obj

if True == ifres11463 {
ifres11462 = True


} else {
ifres11462 = False


}

ifres11461 = ifres11462


} else {
ifres11461 = False


}

if True == ifres11461 {
__e.TailApply(PrimFunc(symshen_4atom_1case_1minus), V2022, V2023, V2024, V2025)
return
} else {
tmp11459 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4_1m, V2021)
}
__typedArg0 := symshen_4_1m
__typedArg1 := V2021
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11429 Obj

if True == tmp11459 {
tmp11458 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11431 Obj

if True == tmp11458 {
tmp11456 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11457 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11456)
}
__typedArg0 := tmp11456
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11433 Obj

if True == tmp11457 {
tmp11453 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11454 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11453)
}
__typedArg0 := tmp11453
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11455 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcons, tmp11454)
}
__typedArg0 := symcons
__typedArg1 := tmp11454
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11435 Obj

if True == tmp11455 {
tmp11450 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11451 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11450)
}
__typedArg0 := tmp11450
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11452 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11451)
}
__typedArg0 := tmp11451
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11437 Obj

if True == tmp11452 {
tmp11446 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11447 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11446)
}
__typedArg0 := tmp11446
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11448 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11447)
}
__typedArg0 := tmp11447
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11449 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11448)
}
__typedArg0 := tmp11448
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11439 Obj

if True == tmp11449 {
tmp11441 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11442 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11441)
}
__typedArg0 := tmp11441
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11443 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11442)
}
__typedArg0 := tmp11442
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11444 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11443)
}
__typedArg0 := tmp11443
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11445 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp11444)
}
__typedArg0 := Nil
__typedArg1 := tmp11444
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11440 Obj

if True == tmp11445 {
ifres11440 = True


} else {
ifres11440 = False


}

ifres11439 = ifres11440


} else {
ifres11439 = False


}

var ifres11438 Obj

if True == ifres11439 {
ifres11438 = True


} else {
ifres11438 = False


}

ifres11437 = ifres11438


} else {
ifres11437 = False


}

var ifres11436 Obj

if True == ifres11437 {
ifres11436 = True


} else {
ifres11436 = False


}

ifres11435 = ifres11436


} else {
ifres11435 = False


}

var ifres11434 Obj

if True == ifres11435 {
ifres11434 = True


} else {
ifres11434 = False


}

ifres11433 = ifres11434


} else {
ifres11433 = False


}

var ifres11432 Obj

if True == ifres11433 {
ifres11432 = True


} else {
ifres11432 = False


}

ifres11431 = ifres11432


} else {
ifres11431 = False


}

var ifres11430 Obj

if True == ifres11431 {
ifres11430 = True


} else {
ifres11430 = False


}

ifres11429 = ifres11430


} else {
ifres11429 = False


}

if True == ifres11429 {
__e.TailApply(PrimFunc(symshen_4cons_1case_1minus), V2022, V2023, V2024, V2025)
return
} else {
tmp11427 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4_7m, V2021)
}
__typedArg0 := symshen_4_7m
__typedArg1 := V2021
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11420 Obj

if True == tmp11427 {
tmp11426 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11422 Obj

if True == tmp11426 {
tmp11424 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11425 := Call(__e, PrimFunc(symatom_2), tmp11424)


var ifres11423 Obj

if True == tmp11425 {
ifres11423 = True


} else {
ifres11423 = False


}

ifres11422 = ifres11423


} else {
ifres11422 = False


}

var ifres11421 Obj

if True == ifres11422 {
ifres11421 = True


} else {
ifres11421 = False


}

ifres11420 = ifres11421


} else {
ifres11420 = False


}

if True == ifres11420 {
__e.TailApply(PrimFunc(symshen_4atom_1case_1plus), V2022, V2023, V2024, V2025)
return
} else {
tmp11418 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4_7m, V2021)
}
__typedArg0 := symshen_4_7m
__typedArg1 := V2021
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11388 Obj

if True == tmp11418 {
tmp11417 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11390 Obj

if True == tmp11417 {
tmp11415 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11416 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11415)
}
__typedArg0 := tmp11415
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11392 Obj

if True == tmp11416 {
tmp11412 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11413 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11412)
}
__typedArg0 := tmp11412
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11414 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcons, tmp11413)
}
__typedArg0 := symcons
__typedArg1 := tmp11413
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11394 Obj

if True == tmp11414 {
tmp11409 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11410 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11409)
}
__typedArg0 := tmp11409
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11411 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11410)
}
__typedArg0 := tmp11410
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11396 Obj

if True == tmp11411 {
tmp11405 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11406 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11405)
}
__typedArg0 := tmp11405
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11407 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11406)
}
__typedArg0 := tmp11406
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11408 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11407)
}
__typedArg0 := tmp11407
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11398 Obj

if True == tmp11408 {
tmp11400 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11401 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11400)
}
__typedArg0 := tmp11400
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11402 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11401)
}
__typedArg0 := tmp11401
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11403 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11402)
}
__typedArg0 := tmp11402
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11404 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp11403)
}
__typedArg0 := Nil
__typedArg1 := tmp11403
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11399 Obj

if True == tmp11404 {
ifres11399 = True


} else {
ifres11399 = False


}

ifres11398 = ifres11399


} else {
ifres11398 = False


}

var ifres11397 Obj

if True == ifres11398 {
ifres11397 = True


} else {
ifres11397 = False


}

ifres11396 = ifres11397


} else {
ifres11396 = False


}

var ifres11395 Obj

if True == ifres11396 {
ifres11395 = True


} else {
ifres11395 = False


}

ifres11394 = ifres11395


} else {
ifres11394 = False


}

var ifres11393 Obj

if True == ifres11394 {
ifres11393 = True


} else {
ifres11393 = False


}

ifres11392 = ifres11393


} else {
ifres11392 = False


}

var ifres11391 Obj

if True == ifres11392 {
ifres11391 = True


} else {
ifres11391 = False


}

ifres11390 = ifres11391


} else {
ifres11390 = False


}

var ifres11389 Obj

if True == ifres11390 {
ifres11389 = True


} else {
ifres11389 = False


}

ifres11388 = ifres11389


} else {
ifres11388 = False


}

if True == ifres11388 {
__e.TailApply(PrimFunc(symshen_4cons_1case_1plus), V2022, V2023, V2024, V2025)
return
} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.compile-head"))
}
__typedArg0 := MakeString("implementation error in shen.compile-head")
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


}


}, 5)

tmp11545 := Call(__e, ns2_1set, symshen_4compile_1head, tmp11368)


_ = tmp11545

tmp11546 := MakeNative(func(__e *ControlFlow) {
V2036 := __e.Get(1)
_ = V2036
V2037 := __e.Get(2)
_ = V2037
V2038 := __e.Get(3)
_ = V2038
V2039 := __e.Get(4)
_ = V2039
V2040 := __e.Get(5)
_ = V2040
tmp11567 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2037)
}
__typedArg0 := V2037
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11564 Obj

if True == tmp11567 {
tmp11566 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2038)
}
__typedArg0 := V2038
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11565 Obj

if True == tmp11566 {
ifres11565 = True


} else {
ifres11565 = False


}

ifres11564 = ifres11565


} else {
ifres11564 = False


}

if True == ifres11564 {
tmp11561 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2038)
}
__typedArg0 := V2038
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11562 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(tmp11561)
}
__typedArg0 := tmp11561
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

if True == tmp11562 {
tmp11547 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2037)
}
__typedArg0 := V2037
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11548 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2038)
}
__typedArg0 := V2038
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11549 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2038)
}
__typedArg0 := V2038
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11550 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2037)
}
__typedArg0 := V2037
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11551 := Call(__e, PrimFunc(symsubst), tmp11549, tmp11550, V2040)


__e.TailApply(PrimFunc(symshen_4compile_1head), V2036, tmp11547, tmp11548, V2039, tmp11551)
return


} else {
tmp11552 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2037)
}
__typedArg0 := V2037
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11553 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2038)
}
__typedArg0 := V2038
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11554 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2037)
}
__typedArg0 := V2037
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11555 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2038)
}
__typedArg0 := V2038
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11556 := Call(__e, PrimFunc(symshen_4compile_1head), V2036, tmp11554, tmp11555, V2039, V2040)


tmp11557 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11556, Nil)
}
__typedArg0 := tmp11556
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11558 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11553, tmp11557)
}
__typedArg0 := tmp11553
__typedArg1 := tmp11557
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11559 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11552, tmp11558)
}
__typedArg0 := tmp11552
__typedArg1 := tmp11558
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp11559)
}
__typedArg0 := symlet
__typedArg1 := tmp11559
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.variable-case"))
}
__typedArg0 := MakeString("implementation error in shen.variable-case")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}, 5)

tmp11568 := Call(__e, ns2_1set, symshen_4variable_1case, tmp11546)


_ = tmp11568

tmp11569 := MakeNative(func(__e *ControlFlow) {
V2049 := __e.Get(1)
_ = V2049
V2050 := __e.Get(2)
_ = V2050
V2051 := __e.Get(3)
_ = V2051
V2052 := __e.Get(4)
_ = V2052
tmp11593 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2049)
}
__typedArg0 := V2049
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11590 Obj

if True == tmp11593 {
tmp11592 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2050)
}
__typedArg0 := V2050
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11591 Obj

if True == tmp11592 {
ifres11591 = True


} else {
ifres11591 = False


}

ifres11590 = ifres11591


} else {
ifres11590 = False


}

if True == ifres11590 {
tmp11570 := Call(__e, PrimFunc(symgensym), symTm)


let__10111 := tmp11570
_ = let__10111

tmp11571 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2050)
}
__typedArg0 := V2050
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11572 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V2051, Nil)
}
__typedArg0 := V2051
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11573 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11571, tmp11572)
}
__typedArg0 := tmp11571
__typedArg1 := tmp11572
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11574 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4lazyderef, tmp11573)
}
__typedArg0 := symshen_4lazyderef
__typedArg1 := tmp11573
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11575 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2049)
}
__typedArg0 := V2049
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11576 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11575, Nil)
}
__typedArg0 := tmp11575
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11577 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10111, tmp11576)
}
__typedArg0 := let__10111
__typedArg1 := tmp11576
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11578 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_a, tmp11577)
}
__typedArg0 := sym_a
__typedArg1 := tmp11577
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11579 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2049)
}
__typedArg0 := V2049
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11580 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2050)
}
__typedArg0 := V2050
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11581 := Call(__e, PrimFunc(symshen_4compile_1head), symshen_4_1m, tmp11579, tmp11580, V2051, V2052)


tmp11582 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(False, Nil)
}
__typedArg0 := False
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11583 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11581, tmp11582)
}
__typedArg0 := tmp11581
__typedArg1 := tmp11582
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11584 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11578, tmp11583)
}
__typedArg0 := tmp11578
__typedArg1 := tmp11583
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11585 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp11584)
}
__typedArg0 := symif
__typedArg1 := tmp11584
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11586 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11585, Nil)
}
__typedArg0 := tmp11585
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11587 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11574, tmp11586)
}
__typedArg0 := tmp11574
__typedArg1 := tmp11586
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11588 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10111, tmp11587)
}
__typedArg0 := let__10111
__typedArg1 := tmp11587
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp11588)
}
__typedArg0 := symlet
__typedArg1 := tmp11588
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.atom-case-minus"))
}
__typedArg0 := MakeString("implementation error in shen.atom-case-minus")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}, 4)

tmp11594 := Call(__e, ns2_1set, symshen_4atom_1case_1minus, tmp11569)


_ = tmp11594

tmp11595 := MakeNative(func(__e *ControlFlow) {
V2062 := __e.Get(1)
_ = V2062
V2063 := __e.Get(2)
_ = V2063
V2064 := __e.Get(3)
_ = V2064
V2065 := __e.Get(4)
_ = V2065
tmp11659 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2062)
}
__typedArg0 := V2062
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11629 Obj

if True == tmp11659 {
tmp11657 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2062)
}
__typedArg0 := V2062
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11658 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11657)
}
__typedArg0 := tmp11657
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11631 Obj

if True == tmp11658 {
tmp11654 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2062)
}
__typedArg0 := V2062
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11655 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11654)
}
__typedArg0 := tmp11654
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11656 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcons, tmp11655)
}
__typedArg0 := symcons
__typedArg1 := tmp11655
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11633 Obj

if True == tmp11656 {
tmp11651 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2062)
}
__typedArg0 := V2062
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11652 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11651)
}
__typedArg0 := tmp11651
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11653 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11652)
}
__typedArg0 := tmp11652
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11635 Obj

if True == tmp11653 {
tmp11647 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2062)
}
__typedArg0 := V2062
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11648 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11647)
}
__typedArg0 := tmp11647
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11649 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11648)
}
__typedArg0 := tmp11648
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11650 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11649)
}
__typedArg0 := tmp11649
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11637 Obj

if True == tmp11650 {
tmp11642 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2062)
}
__typedArg0 := V2062
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11643 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11642)
}
__typedArg0 := tmp11642
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11644 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11643)
}
__typedArg0 := tmp11643
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11645 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11644)
}
__typedArg0 := tmp11644
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11646 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp11645)
}
__typedArg0 := Nil
__typedArg1 := tmp11645
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11639 Obj

if True == tmp11646 {
tmp11641 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2063)
}
__typedArg0 := V2063
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11640 Obj

if True == tmp11641 {
ifres11640 = True


} else {
ifres11640 = False


}

ifres11639 = ifres11640


} else {
ifres11639 = False


}

var ifres11638 Obj

if True == ifres11639 {
ifres11638 = True


} else {
ifres11638 = False


}

ifres11637 = ifres11638


} else {
ifres11637 = False


}

var ifres11636 Obj

if True == ifres11637 {
ifres11636 = True


} else {
ifres11636 = False


}

ifres11635 = ifres11636


} else {
ifres11635 = False


}

var ifres11634 Obj

if True == ifres11635 {
ifres11634 = True


} else {
ifres11634 = False


}

ifres11633 = ifres11634


} else {
ifres11633 = False


}

var ifres11632 Obj

if True == ifres11633 {
ifres11632 = True


} else {
ifres11632 = False


}

ifres11631 = ifres11632


} else {
ifres11631 = False


}

var ifres11630 Obj

if True == ifres11631 {
ifres11630 = True


} else {
ifres11630 = False


}

ifres11629 = ifres11630


} else {
ifres11629 = False


}

if True == ifres11629 {
tmp11596 := Call(__e, PrimFunc(symgensym), symTm)


let__10112 := tmp11596
_ = let__10112

tmp11597 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2063)
}
__typedArg0 := V2063
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11598 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V2064, Nil)
}
__typedArg0 := V2064
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11599 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11597, tmp11598)
}
__typedArg0 := tmp11597
__typedArg1 := tmp11598
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11600 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4lazyderef, tmp11599)
}
__typedArg0 := symshen_4lazyderef
__typedArg1 := tmp11599
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11601 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10112, Nil)
}
__typedArg0 := let__10112
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11602 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons_2, tmp11601)
}
__typedArg0 := symcons_2
__typedArg1 := tmp11601
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11603 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2062)
}
__typedArg0 := V2062
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11604 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11603)
}
__typedArg0 := tmp11603
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11605 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11604)
}
__typedArg0 := tmp11604
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11606 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2062)
}
__typedArg0 := V2062
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11607 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11606)
}
__typedArg0 := tmp11606
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11608 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11607)
}
__typedArg0 := tmp11607
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11609 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11608)
}
__typedArg0 := tmp11608
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11610 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2062)
}
__typedArg0 := V2062
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11611 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11609, tmp11610)
}
__typedArg0 := tmp11609
__typedArg1 := tmp11610
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11612 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11605, tmp11611)
}
__typedArg0 := tmp11605
__typedArg1 := tmp11611
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11613 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10112, Nil)
}
__typedArg0 := let__10112
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11614 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhd, tmp11613)
}
__typedArg0 := symhd
__typedArg1 := tmp11613
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11615 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10112, Nil)
}
__typedArg0 := let__10112
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11616 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtl, tmp11615)
}
__typedArg0 := symtl
__typedArg1 := tmp11615
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11617 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2063)
}
__typedArg0 := V2063
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11618 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11616, tmp11617)
}
__typedArg0 := tmp11616
__typedArg1 := tmp11617
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11619 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11614, tmp11618)
}
__typedArg0 := tmp11614
__typedArg1 := tmp11618
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11620 := Call(__e, PrimFunc(symshen_4compile_1head), symshen_4_1m, tmp11612, tmp11619, V2064, V2065)


tmp11621 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(False, Nil)
}
__typedArg0 := False
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11622 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11620, tmp11621)
}
__typedArg0 := tmp11620
__typedArg1 := tmp11621
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11623 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11602, tmp11622)
}
__typedArg0 := tmp11602
__typedArg1 := tmp11622
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11624 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp11623)
}
__typedArg0 := symif
__typedArg1 := tmp11623
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11625 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11624, Nil)
}
__typedArg0 := tmp11624
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11626 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11600, tmp11625)
}
__typedArg0 := tmp11600
__typedArg1 := tmp11625
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11627 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10112, tmp11626)
}
__typedArg0 := let__10112
__typedArg1 := tmp11626
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp11627)
}
__typedArg0 := symlet
__typedArg1 := tmp11627
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.cons-case-minus"))
}
__typedArg0 := MakeString("implementation error in shen.cons-case-minus")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}, 4)

tmp11660 := Call(__e, ns2_1set, symshen_4cons_1case_1minus, tmp11595)


_ = tmp11660

tmp11661 := MakeNative(func(__e *ControlFlow) {
V2075 := __e.Get(1)
_ = V2075
V2076 := __e.Get(2)
_ = V2076
V2077 := __e.Get(3)
_ = V2077
V2078 := __e.Get(4)
_ = V2078
tmp11705 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2075)
}
__typedArg0 := V2075
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11702 Obj

if True == tmp11705 {
tmp11704 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2076)
}
__typedArg0 := V2076
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11703 Obj

if True == tmp11704 {
ifres11703 = True


} else {
ifres11703 = False


}

ifres11702 = ifres11703


} else {
ifres11702 = False


}

if True == ifres11702 {
tmp11662 := Call(__e, PrimFunc(symgensym), symTm)


let__10113 := tmp11662
_ = let__10113

tmp11663 := Call(__e, PrimFunc(symgensym), symGoTo)


let__10114 := tmp11663
_ = let__10114

tmp11664 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2076)
}
__typedArg0 := V2076
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11665 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V2077, Nil)
}
__typedArg0 := V2077
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11666 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11664, tmp11665)
}
__typedArg0 := tmp11664
__typedArg1 := tmp11665
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11667 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4lazyderef, tmp11666)
}
__typedArg0 := symshen_4lazyderef
__typedArg1 := tmp11666
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11668 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2075)
}
__typedArg0 := V2075
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11669 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2076)
}
__typedArg0 := V2076
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11670 := Call(__e, PrimFunc(symshen_4compile_1head), symshen_4_7m, tmp11668, tmp11669, V2077, V2078)


tmp11671 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11670, Nil)
}
__typedArg0 := tmp11670
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11672 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfreeze, tmp11671)
}
__typedArg0 := symfreeze
__typedArg1 := tmp11671
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11673 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2075)
}
__typedArg0 := V2075
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11674 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11673, Nil)
}
__typedArg0 := tmp11673
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11675 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10113, tmp11674)
}
__typedArg0 := let__10113
__typedArg1 := tmp11674
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11676 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_a, tmp11675)
}
__typedArg0 := sym_a
__typedArg1 := tmp11675
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11677 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10114, Nil)
}
__typedArg0 := let__10114
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11678 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symthaw, tmp11677)
}
__typedArg0 := symthaw
__typedArg1 := tmp11677
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11679 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10113, Nil)
}
__typedArg0 := let__10113
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11680 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4pvar_2, tmp11679)
}
__typedArg0 := symshen_4pvar_2
__typedArg1 := tmp11679
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11681 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2075)
}
__typedArg0 := V2075
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11682 := Call(__e, PrimFunc(symshen_4demode), tmp11681)


tmp11683 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10114, Nil)
}
__typedArg0 := let__10114
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11684 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V2077, tmp11683)
}
__typedArg0 := V2077
__typedArg1 := tmp11683
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11685 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11682, tmp11684)
}
__typedArg0 := tmp11682
__typedArg1 := tmp11684
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11686 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10113, tmp11685)
}
__typedArg0 := let__10113
__typedArg1 := tmp11685
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11687 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4bind_b, tmp11686)
}
__typedArg0 := symshen_4bind_b
__typedArg1 := tmp11686
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11688 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(False, Nil)
}
__typedArg0 := False
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11689 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11687, tmp11688)
}
__typedArg0 := tmp11687
__typedArg1 := tmp11688
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11690 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11680, tmp11689)
}
__typedArg0 := tmp11680
__typedArg1 := tmp11689
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11691 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp11690)
}
__typedArg0 := symif
__typedArg1 := tmp11690
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11692 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11691, Nil)
}
__typedArg0 := tmp11691
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11693 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11678, tmp11692)
}
__typedArg0 := tmp11678
__typedArg1 := tmp11692
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11694 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11676, tmp11693)
}
__typedArg0 := tmp11676
__typedArg1 := tmp11693
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11695 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp11694)
}
__typedArg0 := symif
__typedArg1 := tmp11694
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11696 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11695, Nil)
}
__typedArg0 := tmp11695
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11697 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11672, tmp11696)
}
__typedArg0 := tmp11672
__typedArg1 := tmp11696
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11698 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10114, tmp11697)
}
__typedArg0 := let__10114
__typedArg1 := tmp11697
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11699 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11667, tmp11698)
}
__typedArg0 := tmp11667
__typedArg1 := tmp11698
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11700 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10113, tmp11699)
}
__typedArg0 := let__10113
__typedArg1 := tmp11699
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp11700)
}
__typedArg0 := symlet
__typedArg1 := tmp11700
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.atom-case-plus"))
}
__typedArg0 := MakeString("implementation error in shen.atom-case-plus")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}, 4)

tmp11706 := Call(__e, ns2_1set, symshen_4atom_1case_1plus, tmp11661)


_ = tmp11706

tmp11707 := MakeNative(func(__e *ControlFlow) {
V2089 := __e.Get(1)
_ = V2089
V2090 := __e.Get(2)
_ = V2090
V2091 := __e.Get(3)
_ = V2091
V2092 := __e.Get(4)
_ = V2092
tmp11798 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2089)
}
__typedArg0 := V2089
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11768 Obj

if True == tmp11798 {
tmp11796 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2089)
}
__typedArg0 := V2089
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11797 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11796)
}
__typedArg0 := tmp11796
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11770 Obj

if True == tmp11797 {
tmp11793 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2089)
}
__typedArg0 := V2089
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11794 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11793)
}
__typedArg0 := tmp11793
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11795 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcons, tmp11794)
}
__typedArg0 := symcons
__typedArg1 := tmp11794
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11772 Obj

if True == tmp11795 {
tmp11790 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2089)
}
__typedArg0 := V2089
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11791 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11790)
}
__typedArg0 := tmp11790
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11792 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11791)
}
__typedArg0 := tmp11791
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11774 Obj

if True == tmp11792 {
tmp11786 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2089)
}
__typedArg0 := V2089
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11787 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11786)
}
__typedArg0 := tmp11786
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11788 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11787)
}
__typedArg0 := tmp11787
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11789 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11788)
}
__typedArg0 := tmp11788
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11776 Obj

if True == tmp11789 {
tmp11781 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2089)
}
__typedArg0 := V2089
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11782 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11781)
}
__typedArg0 := tmp11781
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11783 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11782)
}
__typedArg0 := tmp11782
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11784 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11783)
}
__typedArg0 := tmp11783
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11785 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp11784)
}
__typedArg0 := Nil
__typedArg1 := tmp11784
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11778 Obj

if True == tmp11785 {
tmp11780 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2090)
}
__typedArg0 := V2090
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11779 Obj

if True == tmp11780 {
ifres11779 = True


} else {
ifres11779 = False


}

ifres11778 = ifres11779


} else {
ifres11778 = False


}

var ifres11777 Obj

if True == ifres11778 {
ifres11777 = True


} else {
ifres11777 = False


}

ifres11776 = ifres11777


} else {
ifres11776 = False


}

var ifres11775 Obj

if True == ifres11776 {
ifres11775 = True


} else {
ifres11775 = False


}

ifres11774 = ifres11775


} else {
ifres11774 = False


}

var ifres11773 Obj

if True == ifres11774 {
ifres11773 = True


} else {
ifres11773 = False


}

ifres11772 = ifres11773


} else {
ifres11772 = False


}

var ifres11771 Obj

if True == ifres11772 {
ifres11771 = True


} else {
ifres11771 = False


}

ifres11770 = ifres11771


} else {
ifres11770 = False


}

var ifres11769 Obj

if True == ifres11770 {
ifres11769 = True


} else {
ifres11769 = False


}

ifres11768 = ifres11769


} else {
ifres11768 = False


}

if True == ifres11768 {
tmp11708 := Call(__e, PrimFunc(symgensym), symTm)


let__10115 := tmp11708
_ = let__10115

tmp11709 := Call(__e, PrimFunc(symgensym), symGoTo)


let__10116 := tmp11709
_ = let__10116

tmp11710 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2089)
}
__typedArg0 := V2089
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11711 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11710)
}
__typedArg0 := tmp11710
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11712 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11711)
}
__typedArg0 := tmp11711
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11713 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2089)
}
__typedArg0 := V2089
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11714 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11713)
}
__typedArg0 := tmp11713
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11715 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11714)
}
__typedArg0 := tmp11714
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11716 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11715)
}
__typedArg0 := tmp11715
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11717 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11712, tmp11716)
}
__typedArg0 := tmp11712
__typedArg1 := tmp11716
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11718 := Call(__e, PrimFunc(symshen_4extract_1vars), tmp11717)


let__10117 := tmp11718
_ = let__10117

tmp11719 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2089)
}
__typedArg0 := V2089
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11720 := Call(__e, PrimFunc(symshen_4tame), tmp11719)


let__10118 := tmp11720
_ = let__10118

tmp11721 := Call(__e, PrimFunc(symshen_4extract_1vars), let__10118)


let__10119 := tmp11721
_ = let__10119

tmp11722 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2090)
}
__typedArg0 := V2090
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11723 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V2091, Nil)
}
__typedArg0 := V2091
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11724 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11722, tmp11723)
}
__typedArg0 := tmp11722
__typedArg1 := tmp11723
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11725 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4lazyderef, tmp11724)
}
__typedArg0 := symshen_4lazyderef
__typedArg1 := tmp11724
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11726 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2089)
}
__typedArg0 := V2089
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11727 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2090)
}
__typedArg0 := V2090
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11728 := Call(__e, PrimFunc(symshen_4compile_1head), symshen_4_7m, tmp11726, tmp11727, V2091, V2092)


tmp11729 := Call(__e, PrimFunc(symshen_4goto), let__10117, tmp11728)


tmp11730 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10115, Nil)
}
__typedArg0 := let__10115
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11731 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons_2, tmp11730)
}
__typedArg0 := symcons_2
__typedArg1 := tmp11730
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11732 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2089)
}
__typedArg0 := V2089
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11733 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11732)
}
__typedArg0 := tmp11732
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11734 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10115, Nil)
}
__typedArg0 := let__10115
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11735 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhd, tmp11734)
}
__typedArg0 := symhd
__typedArg1 := tmp11734
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11736 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10115, Nil)
}
__typedArg0 := let__10115
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11737 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtl, tmp11736)
}
__typedArg0 := symtl
__typedArg1 := tmp11736
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11738 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11737, Nil)
}
__typedArg0 := tmp11737
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11739 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11735, tmp11738)
}
__typedArg0 := tmp11735
__typedArg1 := tmp11738
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11740 := Call(__e, PrimFunc(symshen_4invoke), let__10116, let__10117)


tmp11741 := Call(__e, PrimFunc(symshen_4compile_1head), symshen_4_7m, tmp11733, tmp11739, V2091, tmp11740)


tmp11742 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10115, Nil)
}
__typedArg0 := let__10115
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11743 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4pvar_2, tmp11742)
}
__typedArg0 := symshen_4pvar_2
__typedArg1 := tmp11742
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11744 := Call(__e, PrimFunc(symshen_4demode), let__10118)


tmp11745 := Call(__e, PrimFunc(symshen_4invoke), let__10116, let__10117)


tmp11746 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11745, Nil)
}
__typedArg0 := tmp11745
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11747 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfreeze, tmp11746)
}
__typedArg0 := symfreeze
__typedArg1 := tmp11746
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11748 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11747, Nil)
}
__typedArg0 := tmp11747
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11749 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V2091, tmp11748)
}
__typedArg0 := V2091
__typedArg1 := tmp11748
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11750 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11744, tmp11749)
}
__typedArg0 := tmp11744
__typedArg1 := tmp11749
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11751 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10115, tmp11750)
}
__typedArg0 := let__10115
__typedArg1 := tmp11750
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11752 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4bind_b, tmp11751)
}
__typedArg0 := symshen_4bind_b
__typedArg1 := tmp11751
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11753 := Call(__e, PrimFunc(symshen_4stpart), let__10119, tmp11752, V2091)


tmp11754 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(False, Nil)
}
__typedArg0 := False
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11755 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11753, tmp11754)
}
__typedArg0 := tmp11753
__typedArg1 := tmp11754
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11756 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11743, tmp11755)
}
__typedArg0 := tmp11743
__typedArg1 := tmp11755
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11757 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp11756)
}
__typedArg0 := symif
__typedArg1 := tmp11756
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11758 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11757, Nil)
}
__typedArg0 := tmp11757
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11759 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11741, tmp11758)
}
__typedArg0 := tmp11741
__typedArg1 := tmp11758
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11760 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11731, tmp11759)
}
__typedArg0 := tmp11731
__typedArg1 := tmp11759
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11761 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp11760)
}
__typedArg0 := symif
__typedArg1 := tmp11760
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11762 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11761, Nil)
}
__typedArg0 := tmp11761
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11763 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11729, tmp11762)
}
__typedArg0 := tmp11729
__typedArg1 := tmp11762
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11764 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10116, tmp11763)
}
__typedArg0 := let__10116
__typedArg1 := tmp11763
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11765 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11725, tmp11764)
}
__typedArg0 := tmp11725
__typedArg1 := tmp11764
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11766 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10115, tmp11765)
}
__typedArg0 := let__10115
__typedArg1 := tmp11765
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp11766)
}
__typedArg0 := symlet
__typedArg1 := tmp11766
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.cons-case-plus"))
}
__typedArg0 := MakeString("implementation error in shen.cons-case-plus")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}, 4)

tmp11799 := Call(__e, ns2_1set, symshen_4cons_1case_1plus, tmp11707)


_ = tmp11799

tmp11800 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__selftop:
V2098 := __self1
_ = V2098
tmp11837 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2098)
}
__typedArg0 := V2098
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11824 Obj

if True == tmp11837 {
tmp11835 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2098)
}
__typedArg0 := V2098
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11836 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4_7m, tmp11835)
}
__typedArg0 := symshen_4_7m
__typedArg1 := tmp11835
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11826 Obj

if True == tmp11836 {
tmp11833 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2098)
}
__typedArg0 := V2098
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11834 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11833)
}
__typedArg0 := tmp11833
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11828 Obj

if True == tmp11834 {
tmp11830 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2098)
}
__typedArg0 := V2098
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11831 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11830)
}
__typedArg0 := tmp11830
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11832 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp11831)
}
__typedArg0 := Nil
__typedArg1 := tmp11831
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11829 Obj

if True == tmp11832 {
ifres11829 = True


} else {
ifres11829 = False


}

ifres11828 = ifres11829


} else {
ifres11828 = False


}

var ifres11827 Obj

if True == ifres11828 {
ifres11827 = True


} else {
ifres11827 = False


}

ifres11826 = ifres11827


} else {
ifres11826 = False


}

var ifres11825 Obj

if True == ifres11826 {
ifres11825 = True


} else {
ifres11825 = False


}

ifres11824 = ifres11825


} else {
ifres11824 = False


}

if True == ifres11824 {
tmp11801 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2098)
}
__typedArg0 := V2098
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11802 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11801)
}
__typedArg0 := tmp11801
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

if PrimFunc(symshen_4demode) == __self {
__self1 = tmp11802
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4demode), tmp11802)
return


} else {
tmp11822 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2098)
}
__typedArg0 := V2098
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11809 Obj

if True == tmp11822 {
tmp11820 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2098)
}
__typedArg0 := V2098
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11821 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4_1m, tmp11820)
}
__typedArg0 := symshen_4_1m
__typedArg1 := tmp11820
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11811 Obj

if True == tmp11821 {
tmp11818 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2098)
}
__typedArg0 := V2098
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11819 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11818)
}
__typedArg0 := tmp11818
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11813 Obj

if True == tmp11819 {
tmp11815 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2098)
}
__typedArg0 := V2098
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11816 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11815)
}
__typedArg0 := tmp11815
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11817 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp11816)
}
__typedArg0 := Nil
__typedArg1 := tmp11816
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11814 Obj

if True == tmp11817 {
ifres11814 = True


} else {
ifres11814 = False


}

ifres11813 = ifres11814


} else {
ifres11813 = False


}

var ifres11812 Obj

if True == ifres11813 {
ifres11812 = True


} else {
ifres11812 = False


}

ifres11811 = ifres11812


} else {
ifres11811 = False


}

var ifres11810 Obj

if True == ifres11811 {
ifres11810 = True


} else {
ifres11810 = False


}

ifres11809 = ifres11810


} else {
ifres11809 = False


}

if True == ifres11809 {
tmp11803 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2098)
}
__typedArg0 := V2098
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11804 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11803)
}
__typedArg0 := tmp11803
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

if PrimFunc(symshen_4demode) == __self {
__self1 = tmp11804
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4demode), tmp11804)
return


} else {
tmp11807 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2098)
}
__typedArg0 := V2098
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp11807 {
tmp11805 := MakeNative(func(__e *ControlFlow) {
Z2099 := __e.Get(1)
_ = Z2099
__e.TailApply(PrimFunc(symshen_4demode), Z2099)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp11805, V2098)
return


} else {
__e.Return(V2098)
return
}


}


}


}, 1)

tmp11838 := Call(__e, ns2_1set, symshen_4demode, tmp11800)


_ = tmp11838

tmp11839 := MakeNative(func(__e *ControlFlow) {
V2100 := __e.Get(1)
_ = V2100
tmp11844 := Call(__e, PrimFunc(symshen_4wildcard_2), V2100)


if True == tmp11844 {
__e.TailApply(PrimFunc(symgensym), symY)
return
} else {
tmp11842 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2100)
}
__typedArg0 := V2100
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp11842 {
tmp11840 := MakeNative(func(__e *ControlFlow) {
Z2101 := __e.Get(1)
_ = Z2101
__e.TailApply(PrimFunc(symshen_4tame), Z2101)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp11840, V2100)
return


} else {
__e.Return(V2100)
return
}


}


}, 1)

tmp11845 := Call(__e, ns2_1set, symshen_4tame, tmp11839)


_ = tmp11845

tmp11846 := MakeNative(func(__e *ControlFlow) {
V2102 := __e.Get(1)
_ = V2102
V2103 := __e.Get(2)
_ = V2103
tmp11849 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V2102)
}
__typedArg0 := Nil
__typedArg1 := V2102
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp11849 {
tmp11847 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V2103, Nil)
}
__typedArg0 := V2103
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfreeze, tmp11847)
}
__typedArg0 := symfreeze
__typedArg1 := tmp11847
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.TailApply(PrimFunc(symshen_4goto_1h), V2102, V2103)
return
}


}, 2)

tmp11850 := Call(__e, ns2_1set, symshen_4goto, tmp11846)


_ = tmp11850

tmp11851 := MakeNative(func(__e *ControlFlow) {
V2104 := __e.Get(1)
_ = V2104
V2105 := __e.Get(2)
_ = V2105
tmp11860 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V2104)
}
__typedArg0 := Nil
__typedArg1 := V2104
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp11860 {
__e.Return(V2105)
return
} else {
tmp11858 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2104)
}
__typedArg0 := V2104
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp11858 {
tmp11852 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2104)
}
__typedArg0 := V2104
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11853 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2104)
}
__typedArg0 := V2104
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11854 := Call(__e, PrimFunc(symshen_4goto_1h), tmp11853, V2105)


tmp11855 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11854, Nil)
}
__typedArg0 := tmp11854
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11856 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11852, tmp11855)
}
__typedArg0 := tmp11852
__typedArg1 := tmp11855
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp11856)
}
__typedArg0 := symlambda
__typedArg1 := tmp11856
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4goto_1h)
return
}


}


}, 2)

tmp11861 := Call(__e, ns2_1set, symshen_4goto_1h, tmp11851)


_ = tmp11861

tmp11862 := MakeNative(func(__e *ControlFlow) {
V2106 := __e.Get(1)
_ = V2106
V2107 := __e.Get(2)
_ = V2107
tmp11865 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V2107)
}
__typedArg0 := Nil
__typedArg1 := V2107
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp11865 {
tmp11863 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V2106, Nil)
}
__typedArg0 := V2106
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symthaw, tmp11863)
}
__typedArg0 := symthaw
__typedArg1 := tmp11863
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V2106, V2107)
}
__typedArg0 := V2106
__typedArg1 := V2107
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return
}


}, 2)

tmp11866 := Call(__e, ns2_1set, symshen_4invoke, tmp11862)


_ = tmp11866

tmp11867 := MakeNative(func(__e *ControlFlow) {
V2108 := __e.Get(1)
_ = V2108
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V2108, sym__)
}
__typedArg0 := V2108
__typedArg1 := sym__
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})())
return
}, 1)

tmp11868 := Call(__e, ns2_1set, symshen_4wildcard_2, tmp11867)


_ = tmp11868

tmp11869 := MakeNative(func(__e *ControlFlow) {
V2109 := __e.Get(1)
_ = V2109
tmp11876 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symabsvector_2) {
return PrimIsVector(V2109)
}
__typedArg0 := V2109
return Call(__e, PrimFunc(symabsvector_2), __typedArg0)
})()

if True == tmp11876 {
tmp11871 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V2109, MakeInteger(0))
}
__typedArg0 := V2109
__typedArg1 := MakeInteger(0)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})())
return
}, 0)

tmp11872 := MakeNative(func(__e *ControlFlow) {
Z2110 := __e.Get(1)
_ = Z2110
__e.Return(symshen_4not_1pvar)
return
}, 1)

tmp11873 := Call(__e, try_1catch, tmp11871, tmp11872)


tmp11874 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp11873, symshen_4pvar)
}
__typedArg0 := tmp11873
__typedArg1 := symshen_4pvar
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp11874 {
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

tmp11877 := Call(__e, ns2_1set, symshen_4pvar_2, tmp11869)


_ = tmp11877

tmp11878 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__selftop:
V2111 := __self1
_ = V2111
V2112 := __self2
_ = V2112
tmp11884 := Call(__e, PrimFunc(symshen_4pvar_2), V2111)


if True == tmp11884 {
tmp11879 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V2111, MakeInteger(1))
}
__typedArg0 := V2111
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})()

tmp11880 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V2112, tmp11879)
}
__typedArg0 := V2112
__typedArg1 := tmp11879
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})()

let__10120 := tmp11880
_ = let__10120

tmp11882 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__10120, symshen_4_1null_1)
}
__typedArg0 := let__10120
__typedArg1 := symshen_4_1null_1
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp11882 {
__e.Return(V2111)
return
} else {
if PrimFunc(symshen_4lazyderef) == __self {
__self1, __self2 = let__10120, V2112
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4lazyderef), let__10120, V2112)
return
}


} else {
__e.Return(V2111)
return
}


}, 2)

tmp11885 := Call(__e, ns2_1set, symshen_4lazyderef, tmp11878)


_ = tmp11885

tmp11886 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__selftop:
V2114 := __self1
_ = V2114
V2115 := __self2
_ = V2115
tmp11898 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2114)
}
__typedArg0 := V2114
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp11898 {
tmp11887 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2114)
}
__typedArg0 := V2114
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11888 := Call(__e, PrimFunc(symshen_4deref), tmp11887, V2115)


tmp11889 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2114)
}
__typedArg0 := V2114
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11890 := Call(__e, PrimFunc(symshen_4deref), tmp11889, V2115)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11888, tmp11890)
}
__typedArg0 := tmp11888
__typedArg1 := tmp11890
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp11896 := Call(__e, PrimFunc(symshen_4pvar_2), V2114)


if True == tmp11896 {
tmp11891 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V2114, MakeInteger(1))
}
__typedArg0 := V2114
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})()

tmp11892 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V2115, tmp11891)
}
__typedArg0 := V2115
__typedArg1 := tmp11891
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})()

let__10121 := tmp11892
_ = let__10121

tmp11894 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__10121, symshen_4_1null_1)
}
__typedArg0 := let__10121
__typedArg1 := symshen_4_1null_1
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp11894 {
__e.Return(V2114)
return
} else {
if PrimFunc(symshen_4deref) == __self {
__self1, __self2 = let__10121, V2115
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4deref), let__10121, V2115)
return
}


} else {
__e.Return(V2114)
return
}


}


}, 2)

tmp11899 := Call(__e, ns2_1set, symshen_4deref, tmp11886)


_ = tmp11899

tmp11900 := MakeNative(func(__e *ControlFlow) {
V2117 := __e.Get(1)
_ = V2117
V2118 := __e.Get(2)
_ = V2118
V2119 := __e.Get(3)
_ = V2119
V2120 := __e.Get(4)
_ = V2120
tmp11901 := Call(__e, PrimFunc(symshen_4bindv), V2117, V2118, V2119)


let__10122 := tmp11901
_ = let__10122

tmp11902 := Call(__e, PrimFunc(symthaw), V2120)


let__10123 := tmp11902
_ = let__10123

tmp11904 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__10123, False)
}
__typedArg0 := let__10123
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp11904 {
__e.TailApply(PrimFunc(symshen_4unwind), V2117, V2119, let__10123)
return
} else {
__e.Return(let__10123)
return
}


}, 4)

tmp11905 := Call(__e, ns2_1set, symshen_4bind_b, tmp11900)


_ = tmp11905

tmp11906 := MakeNative(func(__e *ControlFlow) {
V2123 := __e.Get(1)
_ = V2123
V2124 := __e.Get(2)
_ = V2124
V2125 := __e.Get(3)
_ = V2125
tmp11907 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V2123, MakeInteger(1))
}
__typedArg0 := V2123
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(V2125, tmp11907, V2124)
}
__typedArg0 := V2125
__typedArg1 := tmp11907
__typedArg2 := V2124
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})())
return


}, 3)

tmp11908 := Call(__e, ns2_1set, symshen_4bindv, tmp11906)


_ = tmp11908

tmp11909 := MakeNative(func(__e *ControlFlow) {
V2126 := __e.Get(1)
_ = V2126
V2127 := __e.Get(2)
_ = V2127
V2128 := __e.Get(3)
_ = V2128
tmp11910 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V2126, MakeInteger(1))
}
__typedArg0 := V2126
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})()

tmp11911 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(V2127, tmp11910, symshen_4_1null_1)
}
__typedArg0 := V2127
__typedArg1 := tmp11910
__typedArg2 := symshen_4_1null_1
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

_ = tmp11911

__e.Return(V2128)
return


}, 3)

tmp11912 := Call(__e, ns2_1set, symshen_4unwind, tmp11909)


_ = tmp11912

tmp11913 := MakeNative(func(__e *ControlFlow) {
V2137 := __e.Get(1)
_ = V2137
V2138 := __e.Get(2)
_ = V2138
V2139 := __e.Get(3)
_ = V2139
tmp11928 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V2137)
}
__typedArg0 := Nil
__typedArg1 := V2137
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp11928 {
__e.Return(V2138)
return
} else {
tmp11926 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2137)
}
__typedArg0 := V2137
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp11926 {
tmp11914 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2137)
}
__typedArg0 := V2137
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11915 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V2139, Nil)
}
__typedArg0 := V2139
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11916 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4newpv, tmp11915)
}
__typedArg0 := symshen_4newpv
__typedArg1 := tmp11915
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11917 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2137)
}
__typedArg0 := V2137
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11918 := Call(__e, PrimFunc(symshen_4stpart), tmp11917, V2138, V2139)


tmp11919 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11918, Nil)
}
__typedArg0 := tmp11918
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11920 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V2139, tmp11919)
}
__typedArg0 := V2139
__typedArg1 := tmp11919
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11921 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4gc, tmp11920)
}
__typedArg0 := symshen_4gc
__typedArg1 := tmp11920
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11922 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11921, Nil)
}
__typedArg0 := tmp11921
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11923 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11916, tmp11922)
}
__typedArg0 := tmp11916
__typedArg1 := tmp11922
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11924 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11914, tmp11923)
}
__typedArg0 := tmp11914
__typedArg1 := tmp11923
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp11924)
}
__typedArg0 := symlet
__typedArg1 := tmp11924
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.stpart"))
}
__typedArg0 := MakeString("implementation error in shen.stpart")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 3)

tmp11929 := Call(__e, ns2_1set, symshen_4stpart, tmp11913)


_ = tmp11929

tmp11930 := MakeNative(func(__e *ControlFlow) {
V2140 := __e.Get(1)
_ = V2140
V2141 := __e.Get(2)
_ = V2141
tmp11934 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V2141, False)
}
__typedArg0 := V2141
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp11934 {
tmp11931 := Call(__e, PrimFunc(symshen_4ticket_1number), V2140)


let__10124 := tmp11931
_ = let__10124

tmp11932 := Call(__e, PrimFunc(symshen_4decrement_1ticket), let__10124, V2140)


_ = tmp11932

__e.Return(V2141)
return


} else {
__e.Return(V2141)
return
}


}, 2)

tmp11935 := Call(__e, ns2_1set, symshen_4gc, tmp11930)


_ = tmp11935

tmp11936 := MakeNative(func(__e *ControlFlow) {
V2143 := __e.Get(1)
_ = V2143
V2144 := __e.Get(2)
_ = V2144
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(V2144, MakeInteger(1), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(V2143)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := V2143
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})())
}
__typedArg0 := V2144
__typedArg1 := MakeInteger(1)
__typedArg2 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(V2143)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := V2143
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})())
return


}, 2)

tmp11938 := Call(__e, ns2_1set, symshen_4decrement_1ticket, tmp11936)


_ = tmp11938

tmp11939 := MakeNative(func(__e *ControlFlow) {
V2145 := __e.Get(1)
_ = V2145
tmp11940 := Call(__e, PrimFunc(symshen_4ticket_1number), V2145)


let__10125 := tmp11940
_ = let__10125

tmp11941 := Call(__e, PrimFunc(symshen_4make_1prolog_1variable), let__10125)


let__10126 := tmp11941
_ = let__10126

tmp11942 := Call(__e, PrimFunc(symshen_4nextticket), V2145, let__10125)


let__10127 := tmp11942
_ = let__10127

__e.Return(let__10126)
return


}, 1)

tmp11943 := Call(__e, ns2_1set, symshen_4newpv, tmp11939)


_ = tmp11943

tmp11944 := MakeNative(func(__e *ControlFlow) {
V2149 := __e.Get(1)
_ = V2149
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V2149, MakeInteger(1))
}
__typedArg0 := V2149
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})())
return
}, 1)

tmp11945 := Call(__e, ns2_1set, symshen_4ticket_1number, tmp11944)


_ = tmp11945

tmp11946 := MakeNative(func(__e *ControlFlow) {
V2150 := __e.Get(1)
_ = V2150
V2151 := __e.Get(2)
_ = V2151
tmp11947 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(V2150, V2151, symshen_4_1null_1)
}
__typedArg0 := V2150
__typedArg1 := V2151
__typedArg2 := symshen_4_1null_1
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

let__10128 := tmp11947
_ = let__10128

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(let__10128, MakeInteger(1), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V2151)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V2151
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())
}
__typedArg0 := let__10128
__typedArg1 := MakeInteger(1)
__typedArg2 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V2151)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V2151
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})())
return


}, 2)

tmp11949 := Call(__e, ns2_1set, symshen_4nextticket, tmp11946)


_ = tmp11949

tmp11950 := MakeNative(func(__e *ControlFlow) {
V2153 := __e.Get(1)
_ = V2153
tmp11951 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symabsvector) {
return PrimAbsvector(MakeInteger(2))
}
__typedArg0 := MakeInteger(2)
return Call(__e, PrimFunc(symabsvector), __typedArg0)
})()

tmp11952 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(tmp11951, MakeInteger(0), symshen_4pvar)
}
__typedArg0 := tmp11951
__typedArg1 := MakeInteger(0)
__typedArg2 := symshen_4pvar
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(tmp11952, MakeInteger(1), V2153)
}
__typedArg0 := tmp11952
__typedArg1 := MakeInteger(1)
__typedArg2 := V2153
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})())
return


}, 1)

tmp11953 := Call(__e, ns2_1set, symshen_4make_1prolog_1variable, tmp11950)


_ = tmp11953

tmp11954 := MakeNative(func(__e *ControlFlow) {
V2154 := __e.Get(1)
_ = V2154
tmp11955 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V2154, MakeInteger(1))
}
__typedArg0 := V2154
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})()

tmp11956 := Call(__e, PrimFunc(symshen_4app), tmp11955, MakeString(""), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("Var"))
__typedS1, __typedOK1 := TypedString(tmp11956)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("Var")
__typedArg1 := tmp11956
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp11957 := Call(__e, ns2_1set, symshen_4pvar, tmp11954)


_ = tmp11957

tmp11958 := MakeNative(func(__e *ControlFlow) {
tmp11959 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dinfs_d)
}
__typedArg0 := symshen_4_dinfs_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dinfs_d, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(1))
__typedN1, __typedOK1 := TypedFloat64(tmp11959)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp11959
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())
}
__typedArg0 := symshen_4_dinfs_d
__typedArg1 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(1))
__typedN1, __typedOK1 := TypedFloat64(tmp11959)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp11959
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return


}, 0)

tmp11961 := Call(__e, ns2_1set, symshen_4incinfs, tmp11958)


_ = tmp11961

tmp11962 := MakeNative(func(__e *ControlFlow) {
V2155 := __e.Get(1)
_ = V2155
tmp11969 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(syminteger_2) {
return PrimIsInteger(V2155)
}
__typedArg0 := V2155
return Call(__e, PrimFunc(syminteger_2), __typedArg0)
})()

var ifres11966 Obj

if True == tmp11969 {
tmp11968 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_6) {
__typedN0, __typedOK0 := TypedFloat64(V2155)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(0))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_6) {
return TypedMaterializeBoolean((__typedN0 > __typedN1))
}}
__typedArg0 := V2155
__typedArg1 := MakeInteger(0)
return Call(__e, PrimFunc(sym_6), __typedArg0, __typedArg1)
})()

var ifres11967 Obj

if True == tmp11968 {
ifres11967 = True


} else {
ifres11967 = False


}

ifres11966 = ifres11967


} else {
ifres11966 = False


}

if True == ifres11966 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsize_1prolog_1vector_d, V2155)
}
__typedArg0 := symshen_4_dsize_1prolog_1vector_d
__typedArg1 := V2155
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return
} else {
tmp11963 := Call(__e, PrimFunc(symshen_4app), V2155, MakeString(""), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("prolog vector size: size should be a positive integer; not "))
__typedS1, __typedOK1 := TypedString(tmp11963)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("prolog vector size: size should be a positive integer; not ")
__typedArg1 := tmp11963
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("prolog vector size: size should be a positive integer; not "))
__typedS1, __typedOK1 := TypedString(tmp11963)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("prolog vector size: size should be a positive integer; not ")
__typedArg1 := tmp11963
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}


}, 1)

tmp11970 := Call(__e, ns2_1set, symshen_4prolog_1vector_1size, tmp11962)


_ = tmp11970

tmp11971 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__self3 := __e.Get(3)
__self4 := __e.Get(4)
__selftop:
V2167 := __self1
_ = V2167
V2168 := __self2
_ = V2168
V2169 := __self3
_ = V2169
V2170 := __self4
_ = V2170
tmp12001 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V2167, V2168)
}
__typedArg0 := V2167
__typedArg1 := V2168
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp12001 {
__e.TailApply(PrimFunc(symthaw), V2170)
return
} else {
tmp11999 := Call(__e, PrimFunc(symshen_4pvar_2), V2167)


var ifres11994 Obj

if True == tmp11999 {
tmp11996 := Call(__e, PrimFunc(symshen_4deref), V2168, V2169)


tmp11997 := Call(__e, PrimFunc(symshen_4occurs_1check_2), V2167, tmp11996)


tmp11998 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp11997)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp11997
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres11995 Obj

if True == tmp11998 {
ifres11995 = True


} else {
ifres11995 = False


}

ifres11994 = ifres11995


} else {
ifres11994 = False


}

if True == ifres11994 {
__e.TailApply(PrimFunc(symshen_4bind_b), V2167, V2168, V2169, V2170)
return
} else {
tmp11992 := Call(__e, PrimFunc(symshen_4pvar_2), V2168)


var ifres11987 Obj

if True == tmp11992 {
tmp11989 := Call(__e, PrimFunc(symshen_4deref), V2167, V2169)


tmp11990 := Call(__e, PrimFunc(symshen_4occurs_1check_2), V2168, tmp11989)


tmp11991 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp11990)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp11990
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres11988 Obj

if True == tmp11991 {
ifres11988 = True


} else {
ifres11988 = False


}

ifres11987 = ifres11988


} else {
ifres11987 = False


}

if True == ifres11987 {
__e.TailApply(PrimFunc(symshen_4bind_b), V2168, V2167, V2169, V2170)
return
} else {
tmp11985 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2167)
}
__typedArg0 := V2167
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11982 Obj

if True == tmp11985 {
tmp11984 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2168)
}
__typedArg0 := V2168
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11983 Obj

if True == tmp11984 {
ifres11983 = True


} else {
ifres11983 = False


}

ifres11982 = ifres11983


} else {
ifres11982 = False


}

if True == ifres11982 {
tmp11972 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2167)
}
__typedArg0 := V2167
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11973 := Call(__e, PrimFunc(symshen_4lazyderef), tmp11972, V2169)


tmp11974 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2168)
}
__typedArg0 := V2168
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11975 := Call(__e, PrimFunc(symshen_4lazyderef), tmp11974, V2169)


tmp11976 := MakeNative(func(__e *ControlFlow) {
tmp11977 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2167)
}
__typedArg0 := V2167
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11978 := Call(__e, PrimFunc(symshen_4lazyderef), tmp11977, V2169)


tmp11979 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2168)
}
__typedArg0 := V2168
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11980 := Call(__e, PrimFunc(symshen_4lazyderef), tmp11979, V2169)


__e.TailApply(PrimFunc(symshen_4lzy_a_b), tmp11978, tmp11980, V2169, V2170)
return


}, 0)

if PrimFunc(symshen_4lzy_a_b) == __self {
__self1, __self2, __self3, __self4 = tmp11973, tmp11975, V2169, tmp11976
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4lzy_a_b), tmp11973, tmp11975, V2169, tmp11976)
return


} else {
__e.Return(False)
return
}


}


}


}


}, 4)

tmp12002 := Call(__e, ns2_1set, symshen_4lzy_a_b, tmp11971)


_ = tmp12002

tmp12003 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__self3 := __e.Get(3)
__self4 := __e.Get(4)
__selftop:
V2182 := __self1
_ = V2182
V2183 := __self2
_ = V2183
V2184 := __self3
_ = V2184
V2185 := __self4
_ = V2185
tmp12023 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V2182, V2183)
}
__typedArg0 := V2182
__typedArg1 := V2183
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp12023 {
__e.TailApply(PrimFunc(symthaw), V2185)
return
} else {
tmp12021 := Call(__e, PrimFunc(symshen_4pvar_2), V2182)


if True == tmp12021 {
__e.TailApply(PrimFunc(symshen_4bind_b), V2182, V2183, V2184, V2185)
return
} else {
tmp12019 := Call(__e, PrimFunc(symshen_4pvar_2), V2183)


if True == tmp12019 {
__e.TailApply(PrimFunc(symshen_4bind_b), V2183, V2182, V2184, V2185)
return
} else {
tmp12017 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2182)
}
__typedArg0 := V2182
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12014 Obj

if True == tmp12017 {
tmp12016 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2183)
}
__typedArg0 := V2183
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12015 Obj

if True == tmp12016 {
ifres12015 = True


} else {
ifres12015 = False


}

ifres12014 = ifres12015


} else {
ifres12014 = False


}

if True == ifres12014 {
tmp12004 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2182)
}
__typedArg0 := V2182
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12005 := Call(__e, PrimFunc(symshen_4lazyderef), tmp12004, V2184)


tmp12006 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2183)
}
__typedArg0 := V2183
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12007 := Call(__e, PrimFunc(symshen_4lazyderef), tmp12006, V2184)


tmp12008 := MakeNative(func(__e *ControlFlow) {
tmp12009 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2182)
}
__typedArg0 := V2182
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12010 := Call(__e, PrimFunc(symshen_4lazyderef), tmp12009, V2184)


tmp12011 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2183)
}
__typedArg0 := V2183
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12012 := Call(__e, PrimFunc(symshen_4lazyderef), tmp12011, V2184)


__e.TailApply(PrimFunc(symshen_4lzy_a), tmp12010, tmp12012, V2184, V2185)
return


}, 0)

if PrimFunc(symshen_4lzy_a) == __self {
__self1, __self2, __self3, __self4 = tmp12005, tmp12007, V2184, tmp12008
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4lzy_a), tmp12005, tmp12007, V2184, tmp12008)
return


} else {
__e.Return(False)
return
}


}


}


}


}, 4)

tmp12024 := Call(__e, ns2_1set, symshen_4lzy_a, tmp12003)


_ = tmp12024

tmp12025 := MakeNative(func(__e *ControlFlow) {
V2191 := __e.Get(1)
_ = V2191
V2192 := __e.Get(2)
_ = V2192
tmp12035 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V2191, V2192)
}
__typedArg0 := V2191
__typedArg1 := V2192
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp12035 {
__e.Return(True)
return
} else {
tmp12033 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2192)
}
__typedArg0 := V2192
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp12033 {
tmp12030 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2192)
}
__typedArg0 := V2192
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12031 := Call(__e, PrimFunc(symshen_4occurs_1check_2), V2191, tmp12030)


if True == tmp12031 {
__e.Return(True)
return
} else {
tmp12027 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2192)
}
__typedArg0 := V2192
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12028 := Call(__e, PrimFunc(symshen_4occurs_1check_2), V2191, tmp12027)


if True == tmp12028 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


} else {
__e.Return(False)
return
}


}


}, 2)

tmp12036 := Call(__e, ns2_1set, symshen_4occurs_1check_2, tmp12025)


_ = tmp12036

tmp12037 := MakeNative(func(__e *ControlFlow) {
V2193 := __e.Get(1)
_ = V2193
V2194 := __e.Get(2)
_ = V2194
V2195 := __e.Get(3)
_ = V2195
V2196 := __e.Get(4)
_ = V2196
V2197 := __e.Get(5)
_ = V2197
tmp12038 := Call(__e, V2193, V2194)


tmp12039 := Call(__e, tmp12038, V2195)


tmp12040 := Call(__e, tmp12039, V2196)


__e.TailApply(tmp12040, V2197)
return


}, 5)

tmp12041 := Call(__e, ns2_1set, symcall, tmp12037)


_ = tmp12041

tmp12042 := MakeNative(func(__e *ControlFlow) {
V2204 := __e.Get(1)
_ = V2204
V2205 := __e.Get(2)
_ = V2205
V2206 := __e.Get(3)
_ = V2206
V2207 := __e.Get(4)
_ = V2207
V2208 := __e.Get(5)
_ = V2208
__e.TailApply(PrimFunc(symshen_4deref), V2204, V2205)
return
}, 5)

tmp12043 := Call(__e, ns2_1set, symreturn, tmp12042)


_ = tmp12043

tmp12044 := MakeNative(func(__e *ControlFlow) {
V2215 := __e.Get(1)
_ = V2215
V2216 := __e.Get(2)
_ = V2216
V2217 := __e.Get(3)
_ = V2217
V2218 := __e.Get(4)
_ = V2218
V2219 := __e.Get(5)
_ = V2219
if True == V2215 {
__e.TailApply(PrimFunc(symthaw), V2219)
return
} else {
__e.Return(False)
return
}
}, 5)

tmp12046 := Call(__e, ns2_1set, symwhen, tmp12044)


_ = tmp12046

tmp12047 := MakeNative(func(__e *ControlFlow) {
V2220 := __e.Get(1)
_ = V2220
V2221 := __e.Get(2)
_ = V2221
V2222 := __e.Get(3)
_ = V2222
V2223 := __e.Get(4)
_ = V2223
V2224 := __e.Get(5)
_ = V2224
V2225 := __e.Get(6)
_ = V2225
tmp12048 := Call(__e, PrimFunc(symshen_4lazyderef), V2220, V2222)


tmp12049 := Call(__e, PrimFunc(symshen_4lazyderef), V2221, V2222)


__e.TailApply(PrimFunc(symshen_4lzy_a), tmp12048, tmp12049, V2222, V2225)
return


}, 6)

tmp12050 := Call(__e, ns2_1set, symis, tmp12047)


_ = tmp12050

tmp12051 := MakeNative(func(__e *ControlFlow) {
V2226 := __e.Get(1)
_ = V2226
V2227 := __e.Get(2)
_ = V2227
V2228 := __e.Get(3)
_ = V2228
V2229 := __e.Get(4)
_ = V2229
V2230 := __e.Get(5)
_ = V2230
V2231 := __e.Get(6)
_ = V2231
tmp12052 := Call(__e, PrimFunc(symshen_4lazyderef), V2226, V2228)


tmp12053 := Call(__e, PrimFunc(symshen_4lazyderef), V2227, V2228)


__e.TailApply(PrimFunc(symshen_4lzy_a_b), tmp12052, tmp12053, V2228, V2231)
return


}, 6)

tmp12054 := Call(__e, ns2_1set, symis_b, tmp12051)


_ = tmp12054

tmp12055 := MakeNative(func(__e *ControlFlow) {
V2236 := __e.Get(1)
_ = V2236
V2237 := __e.Get(2)
_ = V2237
V2238 := __e.Get(3)
_ = V2238
V2239 := __e.Get(4)
_ = V2239
V2240 := __e.Get(5)
_ = V2240
V2241 := __e.Get(6)
_ = V2241
__e.TailApply(PrimFunc(symshen_4bind_b), V2236, V2237, V2238, V2241)
return
}, 6)

tmp12056 := Call(__e, ns2_1set, symbind, tmp12055)


_ = tmp12056

tmp12057 := MakeNative(func(__e *ControlFlow) {
V2242 := __e.Get(1)
_ = V2242
V2243 := __e.Get(2)
_ = V2243
V2244 := __e.Get(3)
_ = V2244
V2245 := __e.Get(4)
_ = V2245
V2246 := __e.Get(5)
_ = V2246
tmp12059 := Call(__e, PrimFunc(symshen_4lazyderef), V2242, V2243)


tmp12060 := Call(__e, PrimFunc(symshen_4pvar_2), tmp12059)


if True == tmp12060 {
__e.TailApply(PrimFunc(symthaw), V2246)
return
} else {
__e.Return(False)
return
}


}, 5)

tmp12061 := Call(__e, ns2_1set, symvar_2, tmp12057)


_ = tmp12061

tmp12062 := MakeNative(func(__e *ControlFlow) {
V2249 := __e.Get(1)
_ = V2249
__e.Return(MakeString("|prolog vector|"))
return
}, 1)

tmp12063 := Call(__e, ns2_1set, symshen_4print_1prolog_1vector, tmp12062)


_ = tmp12063

tmp12064 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__self3 := __e.Get(3)
__self4 := __e.Get(4)
__self5 := __e.Get(5)
__selftop:
V2268 := __self1
_ = V2268
V2269 := __self2
_ = V2269
V2270 := __self3
_ = V2270
V2271 := __self4
_ = V2271
V2272 := __self5
_ = V2272
tmp12076 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V2268)
}
__typedArg0 := Nil
__typedArg1 := V2268
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp12076 {
__e.Return(False)
return
} else {
tmp12074 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2268)
}
__typedArg0 := V2268
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp12074 {
tmp12065 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2268)
}
__typedArg0 := V2268
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12066 := Call(__e, tmp12065, V2269)


tmp12067 := Call(__e, tmp12066, V2270)


tmp12068 := Call(__e, tmp12067, V2271)


tmp12069 := Call(__e, tmp12068, V2272)


let__10129 := tmp12069
_ = let__10129

tmp12072 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__10129, False)
}
__typedArg0 := let__10129
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp12072 {
tmp12070 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2268)
}
__typedArg0 := V2268
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symfork) == __self {
__self1, __self2, __self3, __self4, __self5 = tmp12070, V2269, V2270, V2271, V2272
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symfork), tmp12070, V2269, V2270, V2271, V2272)
return


} else {
__e.Return(let__10129)
return
}


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("fork expects a list of literals\n"))
}
__typedArg0 := MakeString("fork expects a list of literals\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 5)

tmp12077 := Call(__e, ns2_1set, symfork, tmp12064)


_ = tmp12077

tmp12078 := MakeNative(func(__e *ControlFlow) {
V2274 := __e.Get(1)
_ = V2274
V2275 := __e.Get(2)
_ = V2275
V2276 := __e.Get(3)
_ = V2276
V2277 := __e.Get(4)
_ = V2277
V2278 := __e.Get(5)
_ = V2278
V2279 := __e.Get(6)
_ = V2279
V2280 := __e.Get(7)
_ = V2280
tmp12084 := Call(__e, PrimFunc(symshen_4unlocked_2), V2278)


if True == tmp12084 {
tmp12079 := Call(__e, PrimFunc(symshen_4newpv), V2277)


let__10130 := tmp12079
_ = let__10130

tmp12080 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp12080

tmp12081 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4findall_1h), V2274, V2275, V2276, let__10130, V2277, V2278, V2279, V2280)
return
}, 0)

tmp12082 := Call(__e, PrimFunc(symis), let__10130, Nil, V2277, V2278, V2279, tmp12081)


__e.TailApply(PrimFunc(symshen_4gc), V2277, tmp12082)
return


} else {
__e.Return(False)
return
}


}, 7)

tmp12085 := Call(__e, ns2_1set, symfindall, tmp12078)


_ = tmp12085

tmp12086 := MakeNative(func(__e *ControlFlow) {
V2282 := __e.Get(1)
_ = V2282
V2283 := __e.Get(2)
_ = V2283
V2284 := __e.Get(3)
_ = V2284
V2285 := __e.Get(4)
_ = V2285
V2286 := __e.Get(5)
_ = V2286
V2287 := __e.Get(6)
_ = V2287
V2288 := __e.Get(7)
_ = V2288
V2289 := __e.Get(8)
_ = V2289
tmp12091 := Call(__e, PrimFunc(symshen_4unlocked_2), V2287)


var ifres12087 Obj

if True == tmp12091 {
tmp12088 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp12088

tmp12089 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4overbind), V2282, V2285, V2286, V2287, V2288, V2289)
return
}, 0)

tmp12090 := Call(__e, PrimFunc(symcall), V2283, V2286, V2287, V2288, tmp12089)


ifres12087 = tmp12090


} else {
ifres12087 = False


}

let__10131 := ifres12087
_ = let__10131

tmp12096 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__10131, False)
}
__typedArg0 := let__10131
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp12096 {
tmp12094 := Call(__e, PrimFunc(symshen_4unlocked_2), V2287)


if True == tmp12094 {
tmp12092 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp12092

__e.TailApply(PrimFunc(symis_b), V2284, V2285, V2286, V2287, V2288, V2289)
return


} else {
__e.Return(False)
return
}


} else {
__e.Return(let__10131)
return
}


}, 8)

tmp12097 := Call(__e, ns2_1set, symshen_4findall_1h, tmp12086)


_ = tmp12097

tmp12098 := MakeNative(func(__e *ControlFlow) {
V2297 := __e.Get(1)
_ = V2297
V2298 := __e.Get(2)
_ = V2298
V2299 := __e.Get(3)
_ = V2299
V2300 := __e.Get(4)
_ = V2300
V2301 := __e.Get(5)
_ = V2301
V2302 := __e.Get(6)
_ = V2302
tmp12099 := Call(__e, PrimFunc(symshen_4deref), V2297, V2299)


tmp12100 := Call(__e, PrimFunc(symshen_4lazyderef), V2298, V2299)


tmp12101 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12099, tmp12100)
}
__typedArg0 := tmp12099
__typedArg1 := tmp12100
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12102 := Call(__e, PrimFunc(symshen_4bindv), V2298, tmp12101, V2299)


_ = tmp12102

__e.Return(False)
return


}, 6)

tmp12103 := Call(__e, ns2_1set, symshen_4overbind, tmp12098)


_ = tmp12103

tmp12104 := MakeNative(func(__e *ControlFlow) {
V2305 := __e.Get(1)
_ = V2305
tmp12108 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_7, V2305)
}
__typedArg0 := sym_7
__typedArg1 := V2305
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp12108 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_doccurs_d, True)
}
__typedArg0 := symshen_4_doccurs_d
__typedArg1 := True
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return
} else {
tmp12106 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_1, V2305)
}
__typedArg0 := sym_1
__typedArg1 := V2305
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp12106 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_doccurs_d, False)
}
__typedArg0 := symshen_4_doccurs_d
__typedArg1 := False
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return
} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("occurs-check expects a + or a -.\n"))
}
__typedArg0 := MakeString("occurs-check expects a + or a -.\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 1)

__e.TailApply(ns2_1set, symoccurs_1check, tmp12104)
return




}, 0)

