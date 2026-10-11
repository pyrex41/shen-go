package main

import . "github.com/pyrex41/shen-go/kl"

var PrologMain = MakeNative(func(__e *ControlFlow) {
tmp10105 := MakeNative(func(__e *ControlFlow) {
V1599 := __e.Get(1)
_ = V1599
__e.TailApply(PrimFunc(symshen_4assert_d), V1599, symshen_4top)
return
}, 1)

tmp10106 := Call(__e, ns2_1set, symasserta, tmp10105)


_ = tmp10106

tmp10107 := MakeNative(func(__e *ControlFlow) {
V1600 := __e.Get(1)
_ = V1600
__e.TailApply(PrimFunc(symshen_4assert_d), V1600, symshen_4bottom)
return
}, 1)

tmp10108 := Call(__e, ns2_1set, symassertz, tmp10107)


_ = tmp10108

tmp10109 := MakeNative(func(__e *ControlFlow) {
V1601 := __e.Get(1)
_ = V1601
V1602 := __e.Get(2)
_ = V1602
tmp10136 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1601)
}
__typedArg0 := V1601
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10127 Obj

if True == tmp10136 {
tmp10134 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1601)
}
__typedArg0 := V1601
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10135 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10134)
}
__typedArg0 := tmp10134
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10129 Obj

if True == tmp10135 {
tmp10131 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1601)
}
__typedArg0 := V1601
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10132 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10131)
}
__typedArg0 := tmp10131
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10133 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_5_1_1, tmp10132)
}
__typedArg0 := sym_5_1_1
__typedArg1 := tmp10132
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10130 Obj

if True == tmp10133 {
ifres10130 = True


} else {
ifres10130 = False


}

ifres10129 = ifres10130


} else {
ifres10129 = False


}

var ifres10128 Obj

if True == ifres10129 {
ifres10128 = True


} else {
ifres10128 = False


}

ifres10127 = ifres10128


} else {
ifres10127 = False


}

if True == ifres10127 {
tmp10110 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1601)
}
__typedArg0 := V1601
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10111 := Call(__e, PrimFunc(symshen_4predicate), tmp10110)


let__9887 := tmp10111
_ = let__9887

tmp10112 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1601)
}
__typedArg0 := V1601
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10113 := Call(__e, PrimFunc(symshen_4terms), tmp10112)


let__9888 := tmp10113
_ = let__9888

tmp10114 := Call(__e, PrimFunc(symlength), let__9888)


let__9889 := tmp10114
_ = let__9889

tmp10115 := Call(__e, PrimFunc(symshen_4parameters), let__9889)


let__9890 := tmp10115
_ = let__9890

tmp10116 := Call(__e, PrimFunc(symarity), let__9887)


let__9891 := tmp10116
_ = let__9891

tmp10122 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__9891, MakeInteger(-1))
}
__typedArg0 := let__9891
__typedArg1 := MakeInteger(-1)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10117 Obj

if True == tmp10122 {
tmp10118 := Call(__e, PrimFunc(symshen_4create_1skeleton), let__9887, let__9890)


tmp10119 := Call(__e, PrimFunc(symeval), tmp10118)


_ = tmp10119

tmp10120 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp10121 := Call(__e, PrimFunc(symput), let__9887, symshen_4dynamic, Nil, tmp10120)


ifres10117 = tmp10121


} else {
ifres10117 = symshen_4skip


}

let__9892 := ifres10117
_ = let__9892

tmp10123 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1601)
}
__typedArg0 := V1601
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10124 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10123)
}
__typedArg0 := tmp10123
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10125 := Call(__e, PrimFunc(symshen_4insert_1info), let__9887, let__9888, tmp10124, V1601, V1602)


let__9893 := tmp10125
_ = let__9893

__e.Return(let__9887)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4assert_d)
return
}


}, 2)

tmp10137 := Call(__e, ns2_1set, symshen_4assert_d, tmp10109)


_ = tmp10137

tmp10138 := MakeNative(func(__e *ControlFlow) {
V1612 := __e.Get(1)
_ = V1612
tmp10140 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1612)
}
__typedArg0 := V1612
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp10140 {
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

tmp10141 := Call(__e, ns2_1set, symshen_4predicate, tmp10138)


_ = tmp10141

tmp10142 := MakeNative(func(__e *ControlFlow) {
V1617 := __e.Get(1)
_ = V1617
tmp10144 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1617)
}
__typedArg0 := V1617
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp10144 {
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

tmp10145 := Call(__e, ns2_1set, symshen_4terms, tmp10142)


_ = tmp10145

tmp10146 := MakeNative(func(__e *ControlFlow) {
V1618 := __e.Get(1)
_ = V1618
V1619 := __e.Get(2)
_ = V1619
tmp10147 := Call(__e, PrimFunc(symshen_4dynamic_1default), V1618, V1619)


tmp10148 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1618, tmp10147)
}
__typedArg0 := V1618
__typedArg1 := tmp10147
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefprolog, tmp10148)
}
__typedArg0 := symdefprolog
__typedArg1 := tmp10148
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 2)

tmp10149 := Call(__e, ns2_1set, symshen_4create_1skeleton, tmp10146)


_ = tmp10149

tmp10150 := MakeNative(func(__e *ControlFlow) {
V1620 := __e.Get(1)
_ = V1620
V1621 := __e.Get(2)
_ = V1621
tmp10151 := Call(__e, PrimFunc(symshen_4cons_1form), V1621)


tmp10152 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4dynamic, Nil)
}
__typedArg0 := symshen_4dynamic
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10153 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1620, tmp10152)
}
__typedArg0 := V1620
__typedArg1 := tmp10152
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10154 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symget, tmp10153)
}
__typedArg0 := symget
__typedArg1 := tmp10153
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10155 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp10154, Nil)
}
__typedArg0 := tmp10154
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10156 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp10151, tmp10155)
}
__typedArg0 := tmp10151
__typedArg1 := tmp10155
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10157 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4call_1dynamic, tmp10156)
}
__typedArg0 := symshen_4call_1dynamic
__typedArg1 := tmp10156
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10158 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(";"))
}
__typedArg0 := MakeString(";")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp10159 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp10158, Nil)
}
__typedArg0 := tmp10158
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10160 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp10157, tmp10159)
}
__typedArg0 := tmp10157
__typedArg1 := tmp10159
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10161 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1_1, tmp10160)
}
__typedArg0 := sym_5_1_1
__typedArg1 := tmp10160
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symappend), V1621, tmp10161)
return


}, 2)

tmp10162 := Call(__e, ns2_1set, symshen_4dynamic_1default, tmp10150)


_ = tmp10162

tmp10163 := MakeNative(func(__e *ControlFlow) {
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
tmp10164 := Call(__e, PrimFunc(symgensym), symshen_4g)


let__9894 := tmp10164
_ = let__9894

tmp10165 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__9894, Nil)
}
__typedArg0 := let__9894
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10166 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefprolog, tmp10165)
}
__typedArg0 := symdefprolog
__typedArg1 := tmp10165
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10167 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1_1, V1624)
}
__typedArg0 := sym_5_1_1
__typedArg1 := V1624
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10168 := Call(__e, PrimFunc(symappend), V1623, tmp10167)


tmp10169 := Call(__e, PrimFunc(symappend), tmp10166, tmp10168)


tmp10170 := Call(__e, PrimFunc(symeval), tmp10169)


let__9895 := tmp10170
_ = let__9895

tmp10171 := Call(__e, PrimFunc(symfn), let__9894)


tmp10172 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__9894, V1625)
}
__typedArg0 := let__9894
__typedArg1 := V1625
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10173 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp10171, tmp10172)
}
__typedArg0 := tmp10171
__typedArg1 := tmp10172
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__9896 := tmp10173
_ = let__9896

tmp10174 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp10175 := Call(__e, PrimFunc(symget), V1622, symshen_4dynamic, tmp10174)


let__9897 := tmp10175
_ = let__9897

tmp10180 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V1626, symshen_4top)
}
__typedArg0 := V1626
__typedArg1 := symshen_4top
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10176 Obj

if True == tmp10180 {
tmp10177 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__9896, let__9897)
}
__typedArg0 := let__9896
__typedArg1 := let__9897
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

ifres10176 = tmp10177


} else {
tmp10178 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__9896, Nil)
}
__typedArg0 := let__9896
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10179 := Call(__e, PrimFunc(symappend), let__9897, tmp10178)


ifres10176 = tmp10179


}

let__9898 := ifres10176
_ = let__9898

tmp10181 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symput), V1622, symshen_4dynamic, let__9898, tmp10181)
return


}, 5)

tmp10182 := Call(__e, ns2_1set, symshen_4insert_1info, tmp10163)


_ = tmp10182

tmp10183 := MakeNative(func(__e *ControlFlow) {
tmp10184 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dnames_d)
}
__typedArg0 := symshen_4_dnames_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

let__9899 := tmp10184
_ = let__9899

tmp10190 := Call(__e, PrimFunc(symempty_2), let__9899)


var ifres10185 Obj

if True == tmp10190 {
tmp10186 := Call(__e, PrimFunc(symgensym), symshen_4g)


ifres10185 = tmp10186


} else {
tmp10187 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__9899)
}
__typedArg0 := let__9899
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10188 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dnames_d, tmp10187)
}
__typedArg0 := symshen_4_dnames_d
__typedArg1 := tmp10187
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp10188

tmp10189 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__9899)
}
__typedArg0 := let__9899
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

ifres10185 = tmp10189


}

let__9900 := ifres10185
_ = let__9900

__e.Return(let__9900)
return


}, 0)

tmp10191 := Call(__e, ns2_1set, symshen_4newname, tmp10183)


_ = tmp10191

tmp10192 := MakeNative(func(__e *ControlFlow) {
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
tmp10204 := Call(__e, PrimFunc(symshen_4unlocked_2), V1637)


var ifres10193 Obj

if True == tmp10204 {
tmp10194 := Call(__e, PrimFunc(symshen_4lazyderef), V1635, V1636)


let__9902 := tmp10194
_ = let__9902

tmp10203 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__9902)
}
__typedArg0 := let__9902
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10195 Obj

if True == tmp10203 {
tmp10196 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__9902)
}
__typedArg0 := let__9902
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10197 := Call(__e, PrimFunc(symshen_4lazyderef), tmp10196, V1636)


let__9903 := tmp10197
_ = let__9903

tmp10202 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__9903)
}
__typedArg0 := let__9903
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10198 Obj

if True == tmp10202 {
tmp10199 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__9903)
}
__typedArg0 := let__9903
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__9904 := tmp10199
_ = let__9904

tmp10200 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp10200

tmp10201 := Call(__e, PrimFunc(symshen_4callrec), let__9904, V1634, V1636, V1637, V1638, V1639)


ifres10198 = tmp10201


} else {
ifres10198 = False


}

ifres10195 = ifres10198


} else {
ifres10195 = False


}

ifres10193 = ifres10195


} else {
ifres10193 = False


}

let__9901 := ifres10193
_ = let__9901

tmp10213 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__9901, False)
}
__typedArg0 := let__9901
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp10213 {
tmp10211 := Call(__e, PrimFunc(symshen_4unlocked_2), V1637)


if True == tmp10211 {
tmp10205 := Call(__e, PrimFunc(symshen_4lazyderef), V1635, V1636)


let__9905 := tmp10205
_ = let__9905

tmp10209 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__9905)
}
__typedArg0 := let__9905
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp10209 {
tmp10206 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__9905)
}
__typedArg0 := let__9905
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__9906 := tmp10206
_ = let__9906

tmp10207 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp10207

if PrimFunc(symshen_4call_1dynamic) == __self {
__self1, __self2, __self3, __self4, __self5, __self6 = V1634, let__9906, V1636, V1637, V1638, V1639
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4call_1dynamic), V1634, let__9906, V1636, V1637, V1638, V1639)
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
__e.Return(let__9901)
return
}


}, 6)

tmp10214 := Call(__e, ns2_1set, symshen_4call_1dynamic, tmp10192)


_ = tmp10214

tmp10215 := MakeNative(func(__e *ControlFlow) {
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
tmp10225 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V1647)
}
__typedArg0 := Nil
__typedArg1 := V1647
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp10225 {
tmp10216 := Call(__e, V1646, V1648)


tmp10217 := Call(__e, tmp10216, V1649)


tmp10218 := Call(__e, tmp10217, V1650)


__e.TailApply(tmp10218, V1651)
return


} else {
tmp10223 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1647)
}
__typedArg0 := V1647
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp10223 {
tmp10219 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1647)
}
__typedArg0 := V1647
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10220 := Call(__e, V1646, tmp10219)


tmp10221 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1647)
}
__typedArg0 := V1647
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4callrec) == __self {
__self1, __self2, __self3, __self4, __self5, __self6 = tmp10220, tmp10221, V1648, V1649, V1650, V1651
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4callrec), tmp10220, tmp10221, V1648, V1649, V1650, V1651)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4callrec)
return
}


}


}, 6)

tmp10226 := Call(__e, ns2_1set, symshen_4callrec, tmp10215)


_ = tmp10226

tmp10227 := MakeNative(func(__e *ControlFlow) {
V1652 := __e.Get(1)
_ = V1652
tmp10244 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1652)
}
__typedArg0 := V1652
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10235 Obj

if True == tmp10244 {
tmp10242 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1652)
}
__typedArg0 := V1652
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10243 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10242)
}
__typedArg0 := tmp10242
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10237 Obj

if True == tmp10243 {
tmp10239 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1652)
}
__typedArg0 := V1652
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10240 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10239)
}
__typedArg0 := tmp10239
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10241 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_5_1_1, tmp10240)
}
__typedArg0 := sym_5_1_1
__typedArg1 := tmp10240
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10238 Obj

if True == tmp10241 {
ifres10238 = True


} else {
ifres10238 = False


}

ifres10237 = ifres10238


} else {
ifres10237 = False


}

var ifres10236 Obj

if True == ifres10237 {
ifres10236 = True


} else {
ifres10236 = False


}

ifres10235 = ifres10236


} else {
ifres10235 = False


}

if True == ifres10235 {
tmp10228 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1652)
}
__typedArg0 := V1652
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10229 := Call(__e, PrimFunc(symshen_4predicate), tmp10228)


let__9907 := tmp10229
_ = let__9907

tmp10230 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp10231 := Call(__e, PrimFunc(symget), let__9907, symshen_4dynamic, tmp10230)


let__9908 := tmp10231
_ = let__9908

tmp10232 := Call(__e, PrimFunc(symshen_4retract_1clause), V1652, let__9908)


tmp10233 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symput), let__9907, symshen_4dynamic, tmp10232, tmp10233)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symretract)
return
}


}, 1)

tmp10245 := Call(__e, ns2_1set, symretract, tmp10227)


_ = tmp10245

tmp10246 := MakeNative(func(__e *ControlFlow) {
V1660 := __e.Get(1)
_ = V1660
V1661 := __e.Get(2)
_ = V1661
tmp10276 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V1661)
}
__typedArg0 := Nil
__typedArg1 := V1661
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp10276 {
__e.Return(Nil)
return
} else {
tmp10274 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1661)
}
__typedArg0 := V1661
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10259 Obj

if True == tmp10274 {
tmp10272 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1661)
}
__typedArg0 := V1661
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10273 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10272)
}
__typedArg0 := tmp10272
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10261 Obj

if True == tmp10273 {
tmp10269 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1661)
}
__typedArg0 := V1661
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10270 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10269)
}
__typedArg0 := tmp10269
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10271 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10270)
}
__typedArg0 := tmp10270
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10263 Obj

if True == tmp10271 {
tmp10265 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1661)
}
__typedArg0 := V1661
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10266 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10265)
}
__typedArg0 := tmp10265
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10267 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10266)
}
__typedArg0 := tmp10266
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10268 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V1660, tmp10267)
}
__typedArg0 := V1660
__typedArg1 := tmp10267
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10264 Obj

if True == tmp10268 {
ifres10264 = True


} else {
ifres10264 = False


}

ifres10263 = ifres10264


} else {
ifres10263 = False


}

var ifres10262 Obj

if True == ifres10263 {
ifres10262 = True


} else {
ifres10262 = False


}

ifres10261 = ifres10262


} else {
ifres10261 = False


}

var ifres10260 Obj

if True == ifres10261 {
ifres10260 = True


} else {
ifres10260 = False


}

ifres10259 = ifres10260


} else {
ifres10259 = False


}

if True == ifres10259 {
tmp10247 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1661)
}
__typedArg0 := V1661
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10248 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10247)
}
__typedArg0 := tmp10247
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10249 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10248)
}
__typedArg0 := tmp10248
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10250 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dnames_d)
}
__typedArg0 := symshen_4_dnames_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp10251 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp10249, tmp10250)
}
__typedArg0 := tmp10249
__typedArg1 := tmp10250
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10252 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dnames_d, tmp10251)
}
__typedArg0 := symshen_4_dnames_d
__typedArg1 := tmp10251
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp10252

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1661)
}
__typedArg0 := V1661
return Call(__e, PrimFunc(symtl), __typedArg0)
})())
return


} else {
tmp10257 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1661)
}
__typedArg0 := V1661
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp10257 {
tmp10253 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1661)
}
__typedArg0 := V1661
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10254 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1661)
}
__typedArg0 := V1661
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10255 := Call(__e, PrimFunc(symshen_4retract_1clause), V1660, tmp10254)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp10253, tmp10255)
}
__typedArg0 := tmp10253
__typedArg1 := tmp10255
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

tmp10277 := Call(__e, ns2_1set, symshen_4retract_1clause, tmp10246)


_ = tmp10277

tmp10278 := MakeNative(func(__e *ControlFlow) {
V1662 := __e.Get(1)
_ = V1662
V1663 := __e.Get(2)
_ = V1663
tmp10279 := MakeNative(func(__e *ControlFlow) {
Z1664 := __e.Get(1)
_ = Z1664
__e.TailApply(PrimFunc(symshen_4_5defprolog_6), Z1664)
return
}, 1)

tmp10280 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1662, V1663)
}
__typedArg0 := V1662
__typedArg1 := V1663
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symcompile), tmp10279, tmp10280)
return


}, 2)

tmp10281 := Call(__e, ns2_1set, symshen_4compile_1prolog, tmp10278)


_ = tmp10281

tmp10282 := MakeNative(func(__e *ControlFlow) {
V1665 := __e.Get(1)
_ = V1665
tmp10298 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1665)
}
__typedArg0 := V1665
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10283 Obj

if True == tmp10298 {
tmp10284 := Call(__e, PrimFunc(symhead), V1665)


let__9910 := tmp10284
_ = let__9910

tmp10285 := Call(__e, PrimFunc(symtail), V1665)


let__9911 := tmp10285
_ = let__9911

tmp10286 := Call(__e, PrimFunc(symshen_4_5clauses_6), let__9911)


let__9912 := tmp10286
_ = let__9912

tmp10296 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9912)


var ifres10287 Obj

if True == tmp10296 {
tmp10288 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10287 = tmp10288


} else {
tmp10289 := Call(__e, PrimFunc(symshen_4_5_1out), let__9912)


let__9913 := tmp10289
_ = let__9913

tmp10290 := Call(__e, PrimFunc(symshen_4in_1_6), let__9912)


let__9914 := tmp10290
_ = let__9914

tmp10291 := Call(__e, PrimFunc(symshen_4prolog_1arity_1check), let__9910, let__9913)


let__9915 := tmp10291
_ = let__9915

tmp10292 := MakeNative(func(__e *ControlFlow) {
Z1674 := __e.Get(1)
_ = Z1674
__e.TailApply(PrimFunc(symshen_4linearise_1clause), Z1674)
return
}, 1)

tmp10293 := Call(__e, PrimFunc(symmap), tmp10292, let__9913)


let__9916 := tmp10293
_ = let__9916

tmp10294 := Call(__e, PrimFunc(symshen_4horn_1clause_1procedure), let__9910, let__9916)


tmp10295 := Call(__e, PrimFunc(symshen_4comb), let__9914, tmp10294)


ifres10287 = tmp10295


}

ifres10283 = ifres10287


} else {
tmp10297 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10283 = tmp10297


}

let__9909 := ifres10283
_ = let__9909

tmp10300 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9909)


if True == tmp10300 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__9909)
return
}


}, 1)

tmp10301 := Call(__e, ns2_1set, symshen_4_5defprolog_6, tmp10282)


_ = tmp10301

tmp10302 := MakeNative(func(__e *ControlFlow) {
V1677 := __e.Get(1)
_ = V1677
V1678 := __e.Get(2)
_ = V1678
tmp10346 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1678)
}
__typedArg0 := V1678
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10327 Obj

if True == tmp10346 {
tmp10344 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1678)
}
__typedArg0 := V1678
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10345 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10344)
}
__typedArg0 := tmp10344
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10329 Obj

if True == tmp10345 {
tmp10341 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1678)
}
__typedArg0 := V1678
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10342 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10341)
}
__typedArg0 := tmp10341
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10343 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10342)
}
__typedArg0 := tmp10342
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10331 Obj

if True == tmp10343 {
tmp10337 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1678)
}
__typedArg0 := V1678
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10338 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10337)
}
__typedArg0 := tmp10337
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10339 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10338)
}
__typedArg0 := tmp10338
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10340 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp10339)
}
__typedArg0 := Nil
__typedArg1 := tmp10339
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10333 Obj

if True == tmp10340 {
tmp10335 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1678)
}
__typedArg0 := V1678
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10336 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp10335)
}
__typedArg0 := Nil
__typedArg1 := tmp10335
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10334 Obj

if True == tmp10336 {
ifres10334 = True


} else {
ifres10334 = False


}

ifres10333 = ifres10334


} else {
ifres10333 = False


}

var ifres10332 Obj

if True == ifres10333 {
ifres10332 = True


} else {
ifres10332 = False


}

ifres10331 = ifres10332


} else {
ifres10331 = False


}

var ifres10330 Obj

if True == ifres10331 {
ifres10330 = True


} else {
ifres10330 = False


}

ifres10329 = ifres10330


} else {
ifres10329 = False


}

var ifres10328 Obj

if True == ifres10329 {
ifres10328 = True


} else {
ifres10328 = False


}

ifres10327 = ifres10328


} else {
ifres10327 = False


}

if True == ifres10327 {
tmp10303 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1678)
}
__typedArg0 := V1678
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10304 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10303)
}
__typedArg0 := tmp10303
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(PrimFunc(symlength), tmp10304)
return


} else {
tmp10325 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1678)
}
__typedArg0 := V1678
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10310 Obj

if True == tmp10325 {
tmp10323 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1678)
}
__typedArg0 := V1678
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10324 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10323)
}
__typedArg0 := tmp10323
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10312 Obj

if True == tmp10324 {
tmp10320 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1678)
}
__typedArg0 := V1678
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10321 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10320)
}
__typedArg0 := tmp10320
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10322 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10321)
}
__typedArg0 := tmp10321
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10314 Obj

if True == tmp10322 {
tmp10316 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1678)
}
__typedArg0 := V1678
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10317 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10316)
}
__typedArg0 := tmp10316
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10318 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10317)
}
__typedArg0 := tmp10317
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10319 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp10318)
}
__typedArg0 := Nil
__typedArg1 := tmp10318
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10315 Obj

if True == tmp10319 {
ifres10315 = True


} else {
ifres10315 = False


}

ifres10314 = ifres10315


} else {
ifres10314 = False


}

var ifres10313 Obj

if True == ifres10314 {
ifres10313 = True


} else {
ifres10313 = False


}

ifres10312 = ifres10313


} else {
ifres10312 = False


}

var ifres10311 Obj

if True == ifres10312 {
ifres10311 = True


} else {
ifres10311 = False


}

ifres10310 = ifres10311


} else {
ifres10310 = False


}

if True == ifres10310 {
tmp10305 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1678)
}
__typedArg0 := V1678
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10306 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10305)
}
__typedArg0 := tmp10305
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10307 := Call(__e, PrimFunc(symlength), tmp10306)


tmp10308 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1678)
}
__typedArg0 := V1678
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4pac_1h), V1677, tmp10307, tmp10308)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4prolog_1arity_1check)
return
}


}


}, 2)

tmp10347 := Call(__e, ns2_1set, symshen_4prolog_1arity_1check, tmp10302)


_ = tmp10347

tmp10348 := MakeNative(func(__e *ControlFlow) {
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
tmp10364 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V1685)
}
__typedArg0 := Nil
__typedArg1 := V1685
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp10364 {
__e.Return(V1684)
return
} else {
tmp10362 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1685)
}
__typedArg0 := V1685
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10358 Obj

if True == tmp10362 {
tmp10360 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1685)
}
__typedArg0 := V1685
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10361 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10360)
}
__typedArg0 := tmp10360
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10359 Obj

if True == tmp10361 {
ifres10359 = True


} else {
ifres10359 = False


}

ifres10358 = ifres10359


} else {
ifres10358 = False


}

if True == ifres10358 {
tmp10353 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1685)
}
__typedArg0 := V1685
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10354 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10353)
}
__typedArg0 := tmp10353
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10355 := Call(__e, PrimFunc(symlength), tmp10354)


tmp10356 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V1684, tmp10355)
}
__typedArg0 := V1684
__typedArg1 := tmp10355
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp10356 {
tmp10349 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1685)
}
__typedArg0 := V1685
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4pac_1h) == __self {
__self1, __self2, __self3 = V1683, V1684, tmp10349
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4pac_1h), V1683, V1684, tmp10349)
return


} else {
tmp10350 := Call(__e, PrimFunc(symshen_4app), V1683, MakeString("\n"), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("arity error in prolog procedure "))
__typedS1, __typedOK1 := TypedString(tmp10350)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("arity error in prolog procedure ")
__typedArg1 := tmp10350
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("arity error in prolog procedure "))
__typedS1, __typedOK1 := TypedString(tmp10350)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("arity error in prolog procedure ")
__typedArg1 := tmp10350
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

tmp10365 := Call(__e, ns2_1set, symshen_4pac_1h, tmp10348)


_ = tmp10365

tmp10366 := MakeNative(func(__e *ControlFlow) {
V1686 := __e.Get(1)
_ = V1686
tmp10367 := Call(__e, PrimFunc(symshen_4_5clause_6), V1686)


let__9918 := tmp10367
_ = let__9918

tmp10380 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9918)


var ifres10368 Obj

if True == tmp10380 {
tmp10369 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10368 = tmp10369


} else {
tmp10370 := Call(__e, PrimFunc(symshen_4_5_1out), let__9918)


let__9919 := tmp10370
_ = let__9919

tmp10371 := Call(__e, PrimFunc(symshen_4in_1_6), let__9918)


let__9920 := tmp10371
_ = let__9920

tmp10372 := Call(__e, PrimFunc(symshen_4_5clauses_6), let__9920)


let__9921 := tmp10372
_ = let__9921

tmp10379 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9921)


var ifres10373 Obj

if True == tmp10379 {
tmp10374 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10373 = tmp10374


} else {
tmp10375 := Call(__e, PrimFunc(symshen_4_5_1out), let__9921)


let__9922 := tmp10375
_ = let__9922

tmp10376 := Call(__e, PrimFunc(symshen_4in_1_6), let__9921)


let__9923 := tmp10376
_ = let__9923

tmp10377 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__9919, let__9922)
}
__typedArg0 := let__9919
__typedArg1 := let__9922
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10378 := Call(__e, PrimFunc(symshen_4comb), let__9923, tmp10377)


ifres10373 = tmp10378


}

ifres10368 = ifres10373


}

let__9917 := ifres10368
_ = let__9917

tmp10396 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9917)


if True == tmp10396 {
tmp10381 := Call(__e, PrimFunc(sym_5_b_6), V1686)


let__9925 := tmp10381
_ = let__9925

tmp10392 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9925)


var ifres10382 Obj

if True == tmp10392 {
tmp10383 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10382 = tmp10383


} else {
tmp10384 := Call(__e, PrimFunc(symshen_4_5_1out), let__9925)


let__9926 := tmp10384
_ = let__9926

tmp10385 := Call(__e, PrimFunc(symshen_4in_1_6), let__9925)


let__9927 := tmp10385
_ = let__9927

tmp10390 := Call(__e, PrimFunc(symempty_2), let__9926)


var ifres10386 Obj

if True == tmp10390 {
ifres10386 = Nil


} else {
tmp10387 := Call(__e, PrimFunc(symshen_4app), let__9926, MakeString("\n ..."), symshen_4r)


tmp10389 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("Prolog syntax error here:\n "))
__typedS1, __typedOK1 := TypedString(tmp10387)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("Prolog syntax error here:\n ")
__typedArg1 := tmp10387
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("Prolog syntax error here:\n "))
__typedS1, __typedOK1 := TypedString(tmp10387)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("Prolog syntax error here:\n ")
__typedArg1 := tmp10387
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})()

ifres10386 = tmp10389


}

tmp10391 := Call(__e, PrimFunc(symshen_4comb), let__9927, ifres10386)


ifres10382 = tmp10391


}

let__9924 := ifres10382
_ = let__9924

tmp10394 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9924)


if True == tmp10394 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__9924)
return
}


} else {
__e.Return(let__9917)
return
}


}, 1)

tmp10397 := Call(__e, ns2_1set, symshen_4_5clauses_6, tmp10366)


_ = tmp10397

tmp10398 := MakeNative(func(__e *ControlFlow) {
V1698 := __e.Get(1)
_ = V1698
tmp10414 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1698)
}
__typedArg0 := V1698
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10405 Obj

if True == tmp10414 {
tmp10412 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1698)
}
__typedArg0 := V1698
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10413 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10412)
}
__typedArg0 := tmp10412
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10407 Obj

if True == tmp10413 {
tmp10409 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1698)
}
__typedArg0 := V1698
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10410 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10409)
}
__typedArg0 := tmp10409
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10411 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp10410)
}
__typedArg0 := Nil
__typedArg1 := tmp10410
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10408 Obj

if True == tmp10411 {
ifres10408 = True


} else {
ifres10408 = False


}

ifres10407 = ifres10408


} else {
ifres10407 = False


}

var ifres10406 Obj

if True == ifres10407 {
ifres10406 = True


} else {
ifres10406 = False


}

ifres10405 = ifres10406


} else {
ifres10405 = False


}

if True == ifres10405 {
tmp10399 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1698)
}
__typedArg0 := V1698
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10400 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1698)
}
__typedArg0 := V1698
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10401 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10400)
}
__typedArg0 := tmp10400
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10402 := Call(__e, PrimFunc(sym_8p), tmp10399, tmp10401)


tmp10403 := Call(__e, PrimFunc(symshen_4linearise), tmp10402)


__e.TailApply(PrimFunc(symshen_4lch), tmp10403)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4linearise_1clause)
return
}


}, 1)

tmp10415 := Call(__e, ns2_1set, symshen_4linearise_1clause, tmp10398)


_ = tmp10415

tmp10416 := MakeNative(func(__e *ControlFlow) {
V1699 := __e.Get(1)
_ = V1699
tmp10422 := Call(__e, PrimFunc(symtuple_2), V1699)


if True == tmp10422 {
tmp10417 := Call(__e, PrimFunc(symfst), V1699)


tmp10418 := Call(__e, PrimFunc(symsnd), V1699)


tmp10419 := Call(__e, PrimFunc(symshen_4lchh), tmp10418)


tmp10420 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp10419, Nil)
}
__typedArg0 := tmp10419
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp10417, tmp10420)
}
__typedArg0 := tmp10417
__typedArg1 := tmp10420
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4lch)
return
}


}, 1)

tmp10423 := Call(__e, ns2_1set, symshen_4lch, tmp10416)


_ = tmp10423

tmp10424 := MakeNative(func(__e *ControlFlow) {
V1700 := __e.Get(1)
_ = V1700
tmp10487 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1700)
}
__typedArg0 := V1700
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10436 Obj

if True == tmp10487 {
tmp10485 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1700)
}
__typedArg0 := V1700
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10486 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symwhere, tmp10485)
}
__typedArg0 := symwhere
__typedArg1 := tmp10485
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10438 Obj

if True == tmp10486 {
tmp10483 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1700)
}
__typedArg0 := V1700
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10484 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10483)
}
__typedArg0 := tmp10483
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10440 Obj

if True == tmp10484 {
tmp10480 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1700)
}
__typedArg0 := V1700
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10481 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10480)
}
__typedArg0 := tmp10480
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10482 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10481)
}
__typedArg0 := tmp10481
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10442 Obj

if True == tmp10482 {
tmp10476 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1700)
}
__typedArg0 := V1700
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10477 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10476)
}
__typedArg0 := tmp10476
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10478 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10477)
}
__typedArg0 := tmp10477
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10479 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_a, tmp10478)
}
__typedArg0 := sym_a
__typedArg1 := tmp10478
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10444 Obj

if True == tmp10479 {
tmp10472 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1700)
}
__typedArg0 := V1700
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10473 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10472)
}
__typedArg0 := tmp10472
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10474 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10473)
}
__typedArg0 := tmp10473
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10475 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10474)
}
__typedArg0 := tmp10474
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10446 Obj

if True == tmp10475 {
tmp10467 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1700)
}
__typedArg0 := V1700
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10468 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10467)
}
__typedArg0 := tmp10467
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10469 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10468)
}
__typedArg0 := tmp10468
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10470 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10469)
}
__typedArg0 := tmp10469
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10471 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10470)
}
__typedArg0 := tmp10470
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10448 Obj

if True == tmp10471 {
tmp10461 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1700)
}
__typedArg0 := V1700
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10462 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10461)
}
__typedArg0 := tmp10461
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10463 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10462)
}
__typedArg0 := tmp10462
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10464 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10463)
}
__typedArg0 := tmp10463
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10465 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10464)
}
__typedArg0 := tmp10464
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10466 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp10465)
}
__typedArg0 := Nil
__typedArg1 := tmp10465
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10450 Obj

if True == tmp10466 {
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10459)
}
__typedArg0 := tmp10459
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10452 Obj

if True == tmp10460 {
tmp10454 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1700)
}
__typedArg0 := V1700
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10455 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10454)
}
__typedArg0 := tmp10454
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10456 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10455)
}
__typedArg0 := tmp10455
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10457 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp10456)
}
__typedArg0 := Nil
__typedArg1 := tmp10456
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10453 Obj

if True == tmp10457 {
ifres10453 = True


} else {
ifres10453 = False


}

ifres10452 = ifres10453


} else {
ifres10452 = False


}

var ifres10451 Obj

if True == ifres10452 {
ifres10451 = True


} else {
ifres10451 = False


}

ifres10450 = ifres10451


} else {
ifres10450 = False


}

var ifres10449 Obj

if True == ifres10450 {
ifres10449 = True


} else {
ifres10449 = False


}

ifres10448 = ifres10449


} else {
ifres10448 = False


}

var ifres10447 Obj

if True == ifres10448 {
ifres10447 = True


} else {
ifres10447 = False


}

ifres10446 = ifres10447


} else {
ifres10446 = False


}

var ifres10445 Obj

if True == ifres10446 {
ifres10445 = True


} else {
ifres10445 = False


}

ifres10444 = ifres10445


} else {
ifres10444 = False


}

var ifres10443 Obj

if True == ifres10444 {
ifres10443 = True


} else {
ifres10443 = False


}

ifres10442 = ifres10443


} else {
ifres10442 = False


}

var ifres10441 Obj

if True == ifres10442 {
ifres10441 = True


} else {
ifres10441 = False


}

ifres10440 = ifres10441


} else {
ifres10440 = False


}

var ifres10439 Obj

if True == ifres10440 {
ifres10439 = True


} else {
ifres10439 = False


}

ifres10438 = ifres10439


} else {
ifres10438 = False


}

var ifres10437 Obj

if True == ifres10438 {
ifres10437 = True


} else {
ifres10437 = False


}

ifres10436 = ifres10437


} else {
ifres10436 = False


}

if True == ifres10436 {
tmp10426 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_doccurs_d)
}
__typedArg0 := symshen_4_doccurs_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

var ifres10425 Obj

if True == tmp10426 {
ifres10425 = symis_b


} else {
ifres10425 = symis


}

tmp10427 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1700)
}
__typedArg0 := V1700
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10428 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10427)
}
__typedArg0 := tmp10427
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10429 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10428)
}
__typedArg0 := tmp10428
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10430 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(ifres10425, tmp10429)
}
__typedArg0 := ifres10425
__typedArg1 := tmp10429
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10431 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1700)
}
__typedArg0 := V1700
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10432 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10431)
}
__typedArg0 := tmp10431
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10433 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10432)
}
__typedArg0 := tmp10432
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10434 := Call(__e, PrimFunc(symshen_4lchh), tmp10433)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp10430, tmp10434)
}
__typedArg0 := tmp10430
__typedArg1 := tmp10434
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return(V1700)
return
}


}, 1)

tmp10488 := Call(__e, ns2_1set, symshen_4lchh, tmp10424)


_ = tmp10488

tmp10489 := MakeNative(func(__e *ControlFlow) {
V1701 := __e.Get(1)
_ = V1701
tmp10490 := Call(__e, PrimFunc(symshen_4_5head_6), V1701)


let__9929 := tmp10490
_ = let__9929

tmp10513 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9929)


var ifres10491 Obj

if True == tmp10513 {
tmp10492 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10491 = tmp10492


} else {
tmp10493 := Call(__e, PrimFunc(symshen_4_5_1out), let__9929)


let__9930 := tmp10493
_ = let__9930

tmp10494 := Call(__e, PrimFunc(symshen_4in_1_6), let__9929)


let__9931 := tmp10494
_ = let__9931

tmp10512 := Call(__e, PrimFunc(symshen_4hds_a_2), let__9931, sym_5_1_1)


var ifres10495 Obj

if True == tmp10512 {
tmp10496 := Call(__e, PrimFunc(symtail), let__9931)


let__9932 := tmp10496
_ = let__9932

tmp10497 := Call(__e, PrimFunc(symshen_4_5body_6), let__9932)


let__9933 := tmp10497
_ = let__9933

tmp10510 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9933)


var ifres10498 Obj

if True == tmp10510 {
tmp10499 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10498 = tmp10499


} else {
tmp10500 := Call(__e, PrimFunc(symshen_4_5_1out), let__9933)


let__9934 := tmp10500
_ = let__9934

tmp10501 := Call(__e, PrimFunc(symshen_4in_1_6), let__9933)


let__9935 := tmp10501
_ = let__9935

tmp10502 := Call(__e, PrimFunc(symshen_4_5sc_6), let__9935)


let__9936 := tmp10502
_ = let__9936

tmp10509 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9936)


var ifres10503 Obj

if True == tmp10509 {
tmp10504 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10503 = tmp10504


} else {
tmp10505 := Call(__e, PrimFunc(symshen_4in_1_6), let__9936)


let__9937 := tmp10505
_ = let__9937

tmp10506 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__9934, Nil)
}
__typedArg0 := let__9934
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10507 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__9930, tmp10506)
}
__typedArg0 := let__9930
__typedArg1 := tmp10506
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10508 := Call(__e, PrimFunc(symshen_4comb), let__9937, tmp10507)


ifres10503 = tmp10508


}

ifres10498 = ifres10503


}

ifres10495 = ifres10498


} else {
tmp10511 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10495 = tmp10511


}

ifres10491 = ifres10495


}

let__9928 := ifres10491
_ = let__9928

tmp10515 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9928)


if True == tmp10515 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__9928)
return
}


}, 1)

tmp10516 := Call(__e, ns2_1set, symshen_4_5clause_6, tmp10489)


_ = tmp10516

tmp10517 := MakeNative(func(__e *ControlFlow) {
V1712 := __e.Get(1)
_ = V1712
tmp10518 := Call(__e, PrimFunc(symshen_4_5hterm_6), V1712)


let__9939 := tmp10518
_ = let__9939

tmp10531 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9939)


var ifres10519 Obj

if True == tmp10531 {
tmp10520 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10519 = tmp10520


} else {
tmp10521 := Call(__e, PrimFunc(symshen_4_5_1out), let__9939)


let__9940 := tmp10521
_ = let__9940

tmp10522 := Call(__e, PrimFunc(symshen_4in_1_6), let__9939)


let__9941 := tmp10522
_ = let__9941

tmp10523 := Call(__e, PrimFunc(symshen_4_5head_6), let__9941)


let__9942 := tmp10523
_ = let__9942

tmp10530 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9942)


var ifres10524 Obj

if True == tmp10530 {
tmp10525 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10524 = tmp10525


} else {
tmp10526 := Call(__e, PrimFunc(symshen_4_5_1out), let__9942)


let__9943 := tmp10526
_ = let__9943

tmp10527 := Call(__e, PrimFunc(symshen_4in_1_6), let__9942)


let__9944 := tmp10527
_ = let__9944

tmp10528 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__9940, let__9943)
}
__typedArg0 := let__9940
__typedArg1 := let__9943
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10529 := Call(__e, PrimFunc(symshen_4comb), let__9944, tmp10528)


ifres10524 = tmp10529


}

ifres10519 = ifres10524


}

let__9938 := ifres10519
_ = let__9938

tmp10541 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9938)


if True == tmp10541 {
tmp10532 := Call(__e, PrimFunc(sym_5e_6), V1712)


let__9946 := tmp10532
_ = let__9946

tmp10537 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9946)


var ifres10533 Obj

if True == tmp10537 {
tmp10534 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10533 = tmp10534


} else {
tmp10535 := Call(__e, PrimFunc(symshen_4in_1_6), let__9946)


let__9947 := tmp10535
_ = let__9947

tmp10536 := Call(__e, PrimFunc(symshen_4comb), let__9947, Nil)


ifres10533 = tmp10536


}

let__9945 := ifres10533
_ = let__9945

tmp10539 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9945)


if True == tmp10539 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__9945)
return
}


} else {
__e.Return(let__9938)
return
}


}, 1)

tmp10542 := Call(__e, ns2_1set, symshen_4_5head_6, tmp10517)


_ = tmp10542

tmp10543 := MakeNative(func(__e *ControlFlow) {
V1723 := __e.Get(1)
_ = V1723
tmp10556 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1723)
}
__typedArg0 := V1723
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10544 Obj

if True == tmp10556 {
tmp10545 := Call(__e, PrimFunc(symhead), V1723)


let__9949 := tmp10545
_ = let__9949

tmp10546 := Call(__e, PrimFunc(symtail), V1723)


let__9950 := tmp10546
_ = let__9950

tmp10554 := Call(__e, PrimFunc(symatom_2), let__9949)


var ifres10550 Obj

if True == tmp10554 {
tmp10552 := Call(__e, PrimFunc(symshen_4prolog_1keyword_2), let__9949)


tmp10553 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp10552)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp10552
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres10551 Obj

if True == tmp10553 {
ifres10551 = True


} else {
ifres10551 = False


}

ifres10550 = ifres10551


} else {
ifres10550 = False


}

var ifres10547 Obj

if True == ifres10550 {
tmp10548 := Call(__e, PrimFunc(symshen_4comb), let__9950, let__9949)


ifres10547 = tmp10548


} else {
tmp10549 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10547 = tmp10549


}

ifres10544 = ifres10547


} else {
tmp10555 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10544 = tmp10555


}

let__9948 := ifres10544
_ = let__9948

tmp10710 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9948)


if True == tmp10710 {
tmp10566 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1723)
}
__typedArg0 := V1723
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10557 Obj

if True == tmp10566 {
tmp10558 := Call(__e, PrimFunc(symhead), V1723)


let__9952 := tmp10558
_ = let__9952

tmp10559 := Call(__e, PrimFunc(symtail), V1723)


let__9953 := tmp10559
_ = let__9953

tmp10563 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp10564 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__9952, tmp10563)
}
__typedArg0 := let__9952
__typedArg1 := tmp10563
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10560 Obj

if True == tmp10564 {
tmp10561 := Call(__e, PrimFunc(symshen_4comb), let__9953, let__9952)


ifres10560 = tmp10561


} else {
tmp10562 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10560 = tmp10562


}

ifres10557 = ifres10560


} else {
tmp10565 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10557 = tmp10565


}

let__9951 := ifres10557
_ = let__9951

tmp10708 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9951)


if True == tmp10708 {
tmp10596 := Call(__e, PrimFunc(symshen_4ccons_2), V1723)


var ifres10567 Obj

if True == tmp10596 {
tmp10568 := Call(__e, PrimFunc(symhead), V1723)


let__9955 := tmp10568
_ = let__9955

tmp10569 := Call(__e, PrimFunc(symtail), V1723)


let__9956 := tmp10569
_ = let__9956

tmp10594 := Call(__e, PrimFunc(symshen_4hds_a_2), let__9955, symcons)


var ifres10570 Obj

if True == tmp10594 {
tmp10571 := Call(__e, PrimFunc(symtail), let__9955)


let__9957 := tmp10571
_ = let__9957

tmp10572 := Call(__e, PrimFunc(symshen_4_5hterm1_6), let__9957)


let__9958 := tmp10572
_ = let__9958

tmp10592 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9958)


var ifres10573 Obj

if True == tmp10592 {
tmp10574 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10573 = tmp10574


} else {
tmp10575 := Call(__e, PrimFunc(symshen_4_5_1out), let__9958)


let__9959 := tmp10575
_ = let__9959

tmp10576 := Call(__e, PrimFunc(symshen_4in_1_6), let__9958)


let__9960 := tmp10576
_ = let__9960

tmp10577 := Call(__e, PrimFunc(symshen_4_5hterm2_6), let__9960)


let__9961 := tmp10577
_ = let__9961

tmp10591 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9961)


var ifres10578 Obj

if True == tmp10591 {
tmp10579 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10578 = tmp10579


} else {
tmp10580 := Call(__e, PrimFunc(symshen_4_5_1out), let__9961)


let__9962 := tmp10580
_ = let__9962

tmp10581 := Call(__e, PrimFunc(symshen_4in_1_6), let__9961)


let__9963 := tmp10581
_ = let__9963

tmp10582 := Call(__e, PrimFunc(sym_5end_6), let__9963)


let__9964 := tmp10582
_ = let__9964

tmp10590 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9964)


var ifres10583 Obj

if True == tmp10590 {
tmp10584 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10583 = tmp10584


} else {
tmp10585 := Call(__e, PrimFunc(symshen_4in_1_6), let__9964)


let__9965 := tmp10585
_ = let__9965

tmp10586 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__9962, Nil)
}
__typedArg0 := let__9962
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10587 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__9959, tmp10586)
}
__typedArg0 := let__9959
__typedArg1 := tmp10586
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10588 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp10587)
}
__typedArg0 := symcons
__typedArg1 := tmp10587
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10589 := Call(__e, PrimFunc(symshen_4comb), let__9956, tmp10588)


ifres10583 = tmp10589


}

ifres10578 = ifres10583


}

ifres10573 = ifres10578


}

ifres10570 = ifres10573


} else {
tmp10593 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10570 = tmp10593


}

ifres10567 = ifres10570


} else {
tmp10595 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10567 = tmp10595


}

let__9954 := ifres10567
_ = let__9954

tmp10706 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9954)


if True == tmp10706 {
tmp10619 := Call(__e, PrimFunc(symshen_4ccons_2), V1723)


var ifres10597 Obj

if True == tmp10619 {
tmp10598 := Call(__e, PrimFunc(symhead), V1723)


let__9967 := tmp10598
_ = let__9967

tmp10599 := Call(__e, PrimFunc(symtail), V1723)


let__9968 := tmp10599
_ = let__9968

tmp10617 := Call(__e, PrimFunc(symshen_4hds_a_2), let__9967, sym_7)


var ifres10600 Obj

if True == tmp10617 {
tmp10601 := Call(__e, PrimFunc(symtail), let__9967)


let__9969 := tmp10601
_ = let__9969

tmp10602 := Call(__e, PrimFunc(symshen_4_5hterm_6), let__9969)


let__9970 := tmp10602
_ = let__9970

tmp10615 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9970)


var ifres10603 Obj

if True == tmp10615 {
tmp10604 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10603 = tmp10604


} else {
tmp10605 := Call(__e, PrimFunc(symshen_4_5_1out), let__9970)


let__9971 := tmp10605
_ = let__9971

tmp10606 := Call(__e, PrimFunc(symshen_4in_1_6), let__9970)


let__9972 := tmp10606
_ = let__9972

tmp10607 := Call(__e, PrimFunc(sym_5end_6), let__9972)


let__9973 := tmp10607
_ = let__9973

tmp10614 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9973)


var ifres10608 Obj

if True == tmp10614 {
tmp10609 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10608 = tmp10609


} else {
tmp10610 := Call(__e, PrimFunc(symshen_4in_1_6), let__9973)


let__9974 := tmp10610
_ = let__9974

tmp10611 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__9971, Nil)
}
__typedArg0 := let__9971
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10612 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4_7m, tmp10611)
}
__typedArg0 := symshen_4_7m
__typedArg1 := tmp10611
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10613 := Call(__e, PrimFunc(symshen_4comb), let__9968, tmp10612)


ifres10608 = tmp10613


}

ifres10603 = ifres10608


}

ifres10600 = ifres10603


} else {
tmp10616 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10600 = tmp10616


}

ifres10597 = ifres10600


} else {
tmp10618 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10597 = tmp10618


}

let__9966 := ifres10597
_ = let__9966

tmp10704 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9966)


if True == tmp10704 {
tmp10642 := Call(__e, PrimFunc(symshen_4ccons_2), V1723)


var ifres10620 Obj

if True == tmp10642 {
tmp10621 := Call(__e, PrimFunc(symhead), V1723)


let__9976 := tmp10621
_ = let__9976

tmp10622 := Call(__e, PrimFunc(symtail), V1723)


let__9977 := tmp10622
_ = let__9977

tmp10640 := Call(__e, PrimFunc(symshen_4hds_a_2), let__9976, sym_1)


var ifres10623 Obj

if True == tmp10640 {
tmp10624 := Call(__e, PrimFunc(symtail), let__9976)


let__9978 := tmp10624
_ = let__9978

tmp10625 := Call(__e, PrimFunc(symshen_4_5hterm_6), let__9978)


let__9979 := tmp10625
_ = let__9979

tmp10638 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9979)


var ifres10626 Obj

if True == tmp10638 {
tmp10627 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10626 = tmp10627


} else {
tmp10628 := Call(__e, PrimFunc(symshen_4_5_1out), let__9979)


let__9980 := tmp10628
_ = let__9980

tmp10629 := Call(__e, PrimFunc(symshen_4in_1_6), let__9979)


let__9981 := tmp10629
_ = let__9981

tmp10630 := Call(__e, PrimFunc(sym_5end_6), let__9981)


let__9982 := tmp10630
_ = let__9982

tmp10637 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9982)


var ifres10631 Obj

if True == tmp10637 {
tmp10632 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10631 = tmp10632


} else {
tmp10633 := Call(__e, PrimFunc(symshen_4in_1_6), let__9982)


let__9983 := tmp10633
_ = let__9983

tmp10634 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__9980, Nil)
}
__typedArg0 := let__9980
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10635 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4_1m, tmp10634)
}
__typedArg0 := symshen_4_1m
__typedArg1 := tmp10634
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10636 := Call(__e, PrimFunc(symshen_4comb), let__9977, tmp10635)


ifres10631 = tmp10636


}

ifres10626 = ifres10631


}

ifres10623 = ifres10626


} else {
tmp10639 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10623 = tmp10639


}

ifres10620 = ifres10623


} else {
tmp10641 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10620 = tmp10641


}

let__9975 := ifres10620
_ = let__9975

tmp10702 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9975)


if True == tmp10702 {
tmp10669 := Call(__e, PrimFunc(symshen_4ccons_2), V1723)


var ifres10643 Obj

if True == tmp10669 {
tmp10644 := Call(__e, PrimFunc(symhead), V1723)


let__9985 := tmp10644
_ = let__9985

tmp10645 := Call(__e, PrimFunc(symtail), V1723)


let__9986 := tmp10645
_ = let__9986

tmp10667 := Call(__e, PrimFunc(symshen_4hds_a_2), let__9985, symmode)


var ifres10646 Obj

if True == tmp10667 {
tmp10647 := Call(__e, PrimFunc(symtail), let__9985)


let__9987 := tmp10647
_ = let__9987

tmp10648 := Call(__e, PrimFunc(symshen_4_5hterm_6), let__9987)


let__9988 := tmp10648
_ = let__9988

tmp10665 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9988)


var ifres10649 Obj

if True == tmp10665 {
tmp10650 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10649 = tmp10650


} else {
tmp10651 := Call(__e, PrimFunc(symshen_4_5_1out), let__9988)


let__9989 := tmp10651
_ = let__9989

tmp10652 := Call(__e, PrimFunc(symshen_4in_1_6), let__9988)


let__9990 := tmp10652
_ = let__9990

tmp10664 := Call(__e, PrimFunc(symshen_4hds_a_2), let__9990, sym_7)


var ifres10653 Obj

if True == tmp10664 {
tmp10654 := Call(__e, PrimFunc(symtail), let__9990)


let__9991 := tmp10654
_ = let__9991

tmp10655 := Call(__e, PrimFunc(sym_5end_6), let__9991)


let__9992 := tmp10655
_ = let__9992

tmp10662 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9992)


var ifres10656 Obj

if True == tmp10662 {
tmp10657 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10656 = tmp10657


} else {
tmp10658 := Call(__e, PrimFunc(symshen_4in_1_6), let__9992)


let__9993 := tmp10658
_ = let__9993

tmp10659 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__9989, Nil)
}
__typedArg0 := let__9989
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10660 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4_7m, tmp10659)
}
__typedArg0 := symshen_4_7m
__typedArg1 := tmp10659
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10661 := Call(__e, PrimFunc(symshen_4comb), let__9986, tmp10660)


ifres10656 = tmp10661


}

ifres10653 = ifres10656


} else {
tmp10663 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10653 = tmp10663


}

ifres10649 = ifres10653


}

ifres10646 = ifres10649


} else {
tmp10666 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10646 = tmp10666


}

ifres10643 = ifres10646


} else {
tmp10668 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10643 = tmp10668


}

let__9984 := ifres10643
_ = let__9984

tmp10700 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9984)


if True == tmp10700 {
tmp10696 := Call(__e, PrimFunc(symshen_4ccons_2), V1723)


var ifres10670 Obj

if True == tmp10696 {
tmp10671 := Call(__e, PrimFunc(symhead), V1723)


let__9995 := tmp10671
_ = let__9995

tmp10672 := Call(__e, PrimFunc(symtail), V1723)


let__9996 := tmp10672
_ = let__9996

tmp10694 := Call(__e, PrimFunc(symshen_4hds_a_2), let__9995, symmode)


var ifres10673 Obj

if True == tmp10694 {
tmp10674 := Call(__e, PrimFunc(symtail), let__9995)


let__9997 := tmp10674
_ = let__9997

tmp10675 := Call(__e, PrimFunc(symshen_4_5hterm_6), let__9997)


let__9998 := tmp10675
_ = let__9998

tmp10692 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9998)


var ifres10676 Obj

if True == tmp10692 {
tmp10677 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10676 = tmp10677


} else {
tmp10678 := Call(__e, PrimFunc(symshen_4_5_1out), let__9998)


let__9999 := tmp10678
_ = let__9999

tmp10679 := Call(__e, PrimFunc(symshen_4in_1_6), let__9998)


let__10000 := tmp10679
_ = let__10000

tmp10691 := Call(__e, PrimFunc(symshen_4hds_a_2), let__10000, sym_1)


var ifres10680 Obj

if True == tmp10691 {
tmp10681 := Call(__e, PrimFunc(symtail), let__10000)


let__10001 := tmp10681
_ = let__10001

tmp10682 := Call(__e, PrimFunc(sym_5end_6), let__10001)


let__10002 := tmp10682
_ = let__10002

tmp10689 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10002)


var ifres10683 Obj

if True == tmp10689 {
tmp10684 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10683 = tmp10684


} else {
tmp10685 := Call(__e, PrimFunc(symshen_4in_1_6), let__10002)


let__10003 := tmp10685
_ = let__10003

tmp10686 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__9999, Nil)
}
__typedArg0 := let__9999
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10687 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4_1m, tmp10686)
}
__typedArg0 := symshen_4_1m
__typedArg1 := tmp10686
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10688 := Call(__e, PrimFunc(symshen_4comb), let__9996, tmp10687)


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

let__9994 := ifres10670
_ = let__9994

tmp10698 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9994)


if True == tmp10698 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__9994)
return
}


} else {
__e.Return(let__9984)
return
}


} else {
__e.Return(let__9975)
return
}


} else {
__e.Return(let__9966)
return
}


} else {
__e.Return(let__9954)
return
}


} else {
__e.Return(let__9951)
return
}


} else {
__e.Return(let__9948)
return
}


}, 1)

tmp10711 := Call(__e, ns2_1set, symshen_4_5hterm_6, tmp10543)


_ = tmp10711

tmp10712 := MakeNative(func(__e *ControlFlow) {
V1780 := __e.Get(1)
_ = V1780
tmp10713 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(";"))
}
__typedArg0 := MakeString(";")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp10714 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1_1, Nil)
}
__typedArg0 := sym_5_1_1
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10715 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp10713, tmp10714)
}
__typedArg0 := tmp10713
__typedArg1 := tmp10714
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symelement_2), V1780, tmp10715)
return


}, 1)

tmp10716 := Call(__e, ns2_1set, symshen_4prolog_1keyword_2, tmp10712)


_ = tmp10716

tmp10717 := MakeNative(func(__e *ControlFlow) {
V1781 := __e.Get(1)
_ = V1781
tmp10730 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsymbol_2) {
return PrimIsSymbol(V1781)
}
__typedArg0 := V1781
return Call(__e, PrimFunc(symsymbol_2), __typedArg0)
})()

if True == tmp10730 {
__e.Return(True)
return
} else {
tmp10728 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_2) {
return PrimIsString(V1781)
}
__typedArg0 := V1781
return Call(__e, PrimFunc(symstring_2), __typedArg0)
})()

var ifres10719 Obj

if True == tmp10728 {
ifres10719 = True


} else {
tmp10727 := Call(__e, PrimFunc(symboolean_2), V1781)


var ifres10721 Obj

if True == tmp10727 {
ifres10721 = True


} else {
tmp10726 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnumber_2) {
return PrimIsNumber(V1781)
}
__typedArg0 := V1781
return Call(__e, PrimFunc(symnumber_2), __typedArg0)
})()

var ifres10723 Obj

if True == tmp10726 {
ifres10723 = True


} else {
tmp10725 := Call(__e, PrimFunc(symempty_2), V1781)


var ifres10724 Obj

if True == tmp10725 {
ifres10724 = True


} else {
ifres10724 = False


}

ifres10723 = ifres10724


}

var ifres10722 Obj

if True == ifres10723 {
ifres10722 = True


} else {
ifres10722 = False


}

ifres10721 = ifres10722


}

var ifres10720 Obj

if True == ifres10721 {
ifres10720 = True


} else {
ifres10720 = False


}

ifres10719 = ifres10720


}

if True == ifres10719 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp10731 := Call(__e, ns2_1set, symatom_2, tmp10717)


_ = tmp10731

tmp10732 := MakeNative(func(__e *ControlFlow) {
V1782 := __e.Get(1)
_ = V1782
tmp10733 := Call(__e, PrimFunc(symshen_4_5hterm_6), V1782)


let__10005 := tmp10733
_ = let__10005

tmp10739 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10005)


var ifres10734 Obj

if True == tmp10739 {
tmp10735 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10734 = tmp10735


} else {
tmp10736 := Call(__e, PrimFunc(symshen_4_5_1out), let__10005)


let__10006 := tmp10736
_ = let__10006

tmp10737 := Call(__e, PrimFunc(symshen_4in_1_6), let__10005)


let__10007 := tmp10737
_ = let__10007

tmp10738 := Call(__e, PrimFunc(symshen_4comb), let__10007, let__10006)


ifres10734 = tmp10738


}

let__10004 := ifres10734
_ = let__10004

tmp10741 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10004)


if True == tmp10741 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__10004)
return
}


}, 1)

tmp10742 := Call(__e, ns2_1set, symshen_4_5hterm1_6, tmp10732)


_ = tmp10742

tmp10743 := MakeNative(func(__e *ControlFlow) {
V1787 := __e.Get(1)
_ = V1787
tmp10744 := Call(__e, PrimFunc(symshen_4_5hterm_6), V1787)


let__10009 := tmp10744
_ = let__10009

tmp10750 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10009)


var ifres10745 Obj

if True == tmp10750 {
tmp10746 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10745 = tmp10746


} else {
tmp10747 := Call(__e, PrimFunc(symshen_4_5_1out), let__10009)


let__10010 := tmp10747
_ = let__10010

tmp10748 := Call(__e, PrimFunc(symshen_4in_1_6), let__10009)


let__10011 := tmp10748
_ = let__10011

tmp10749 := Call(__e, PrimFunc(symshen_4comb), let__10011, let__10010)


ifres10745 = tmp10749


}

let__10008 := ifres10745
_ = let__10008

tmp10752 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10008)


if True == tmp10752 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__10008)
return
}


}, 1)

tmp10753 := Call(__e, ns2_1set, symshen_4_5hterm2_6, tmp10743)


_ = tmp10753

tmp10754 := MakeNative(func(__e *ControlFlow) {
V1792 := __e.Get(1)
_ = V1792
tmp10755 := Call(__e, PrimFunc(symshen_4_5literal_6), V1792)


let__10013 := tmp10755
_ = let__10013

tmp10768 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10013)


var ifres10756 Obj

if True == tmp10768 {
tmp10757 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10756 = tmp10757


} else {
tmp10758 := Call(__e, PrimFunc(symshen_4_5_1out), let__10013)


let__10014 := tmp10758
_ = let__10014

tmp10759 := Call(__e, PrimFunc(symshen_4in_1_6), let__10013)


let__10015 := tmp10759
_ = let__10015

tmp10760 := Call(__e, PrimFunc(symshen_4_5body_6), let__10015)


let__10016 := tmp10760
_ = let__10016

tmp10767 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10016)


var ifres10761 Obj

if True == tmp10767 {
tmp10762 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10761 = tmp10762


} else {
tmp10763 := Call(__e, PrimFunc(symshen_4_5_1out), let__10016)


let__10017 := tmp10763
_ = let__10017

tmp10764 := Call(__e, PrimFunc(symshen_4in_1_6), let__10016)


let__10018 := tmp10764
_ = let__10018

tmp10765 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10014, let__10017)
}
__typedArg0 := let__10014
__typedArg1 := let__10017
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10766 := Call(__e, PrimFunc(symshen_4comb), let__10018, tmp10765)


ifres10761 = tmp10766


}

ifres10756 = ifres10761


}

let__10012 := ifres10756
_ = let__10012

tmp10778 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10012)


if True == tmp10778 {
tmp10769 := Call(__e, PrimFunc(sym_5e_6), V1792)


let__10020 := tmp10769
_ = let__10020

tmp10774 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10020)


var ifres10770 Obj

if True == tmp10774 {
tmp10771 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10770 = tmp10771


} else {
tmp10772 := Call(__e, PrimFunc(symshen_4in_1_6), let__10020)


let__10021 := tmp10772
_ = let__10021

tmp10773 := Call(__e, PrimFunc(symshen_4comb), let__10021, Nil)


ifres10770 = tmp10773


}

let__10019 := ifres10770
_ = let__10019

tmp10776 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10019)


if True == tmp10776 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__10019)
return
}


} else {
__e.Return(let__10012)
return
}


}, 1)

tmp10779 := Call(__e, ns2_1set, symshen_4_5body_6, tmp10754)


_ = tmp10779

tmp10780 := MakeNative(func(__e *ControlFlow) {
V1803 := __e.Get(1)
_ = V1803
tmp10785 := Call(__e, PrimFunc(symshen_4hds_a_2), V1803, sym_b)


var ifres10781 Obj

if True == tmp10785 {
tmp10782 := Call(__e, PrimFunc(symtail), V1803)


let__10023 := tmp10782
_ = let__10023

tmp10783 := Call(__e, PrimFunc(symshen_4comb), let__10023, sym_b)


ifres10781 = tmp10783


} else {
tmp10784 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10781 = tmp10784


}

let__10022 := ifres10781
_ = let__10022

tmp10806 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10022)


if True == tmp10806 {
tmp10802 := Call(__e, PrimFunc(symshen_4ccons_2), V1803)


var ifres10786 Obj

if True == tmp10802 {
tmp10787 := Call(__e, PrimFunc(symhead), V1803)


let__10025 := tmp10787
_ = let__10025

tmp10788 := Call(__e, PrimFunc(symtail), V1803)


let__10026 := tmp10788
_ = let__10026

tmp10789 := Call(__e, PrimFunc(symshen_4_5bterms_6), let__10025)


let__10027 := tmp10789
_ = let__10027

tmp10800 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10027)


var ifres10790 Obj

if True == tmp10800 {
tmp10791 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10790 = tmp10791


} else {
tmp10792 := Call(__e, PrimFunc(symshen_4_5_1out), let__10027)


let__10028 := tmp10792
_ = let__10028

tmp10793 := Call(__e, PrimFunc(symshen_4in_1_6), let__10027)


let__10029 := tmp10793
_ = let__10029

tmp10794 := Call(__e, PrimFunc(sym_5end_6), let__10029)


let__10030 := tmp10794
_ = let__10030

tmp10799 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10030)


var ifres10795 Obj

if True == tmp10799 {
tmp10796 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10795 = tmp10796


} else {
tmp10797 := Call(__e, PrimFunc(symshen_4in_1_6), let__10030)


let__10031 := tmp10797
_ = let__10031

tmp10798 := Call(__e, PrimFunc(symshen_4comb), let__10026, let__10028)


ifres10795 = tmp10798


}

ifres10790 = ifres10795


}

ifres10786 = ifres10790


} else {
tmp10801 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10786 = tmp10801


}

let__10024 := ifres10786
_ = let__10024

tmp10804 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10024)


if True == tmp10804 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__10024)
return
}


} else {
__e.Return(let__10022)
return
}


}, 1)

tmp10807 := Call(__e, ns2_1set, symshen_4_5literal_6, tmp10780)


_ = tmp10807

tmp10808 := MakeNative(func(__e *ControlFlow) {
V1814 := __e.Get(1)
_ = V1814
tmp10809 := Call(__e, PrimFunc(symshen_4_5bterm_6), V1814)


let__10033 := tmp10809
_ = let__10033

tmp10822 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10033)


var ifres10810 Obj

if True == tmp10822 {
tmp10811 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10810 = tmp10811


} else {
tmp10812 := Call(__e, PrimFunc(symshen_4_5_1out), let__10033)


let__10034 := tmp10812
_ = let__10034

tmp10813 := Call(__e, PrimFunc(symshen_4in_1_6), let__10033)


let__10035 := tmp10813
_ = let__10035

tmp10814 := Call(__e, PrimFunc(symshen_4_5bterms_6), let__10035)


let__10036 := tmp10814
_ = let__10036

tmp10821 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10036)


var ifres10815 Obj

if True == tmp10821 {
tmp10816 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10815 = tmp10816


} else {
tmp10817 := Call(__e, PrimFunc(symshen_4_5_1out), let__10036)


let__10037 := tmp10817
_ = let__10037

tmp10818 := Call(__e, PrimFunc(symshen_4in_1_6), let__10036)


let__10038 := tmp10818
_ = let__10038

tmp10819 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10034, let__10037)
}
__typedArg0 := let__10034
__typedArg1 := let__10037
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10820 := Call(__e, PrimFunc(symshen_4comb), let__10038, tmp10819)


ifres10815 = tmp10820


}

ifres10810 = ifres10815


}

let__10032 := ifres10810
_ = let__10032

tmp10832 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10032)


if True == tmp10832 {
tmp10823 := Call(__e, PrimFunc(sym_5e_6), V1814)


let__10040 := tmp10823
_ = let__10040

tmp10828 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10040)


var ifres10824 Obj

if True == tmp10828 {
tmp10825 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10824 = tmp10825


} else {
tmp10826 := Call(__e, PrimFunc(symshen_4in_1_6), let__10040)


let__10041 := tmp10826
_ = let__10041

tmp10827 := Call(__e, PrimFunc(symshen_4comb), let__10041, Nil)


ifres10824 = tmp10827


}

let__10039 := ifres10824
_ = let__10039

tmp10830 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10039)


if True == tmp10830 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__10039)
return
}


} else {
__e.Return(let__10032)
return
}


}, 1)

tmp10833 := Call(__e, ns2_1set, symshen_4_5bterms_6, tmp10808)


_ = tmp10833

tmp10834 := MakeNative(func(__e *ControlFlow) {
V1825 := __e.Get(1)
_ = V1825
tmp10835 := Call(__e, PrimFunc(symshen_4_5wildcard_6), V1825)


let__10043 := tmp10835
_ = let__10043

tmp10841 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10043)


var ifres10836 Obj

if True == tmp10841 {
tmp10837 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10836 = tmp10837


} else {
tmp10838 := Call(__e, PrimFunc(symshen_4_5_1out), let__10043)


let__10044 := tmp10838
_ = let__10044

tmp10839 := Call(__e, PrimFunc(symshen_4in_1_6), let__10043)


let__10045 := tmp10839
_ = let__10045

tmp10840 := Call(__e, PrimFunc(symshen_4comb), let__10045, let__10044)


ifres10836 = tmp10840


}

let__10042 := ifres10836
_ = let__10042

tmp10873 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10042)


if True == tmp10873 {
tmp10850 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1825)
}
__typedArg0 := V1825
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10842 Obj

if True == tmp10850 {
tmp10843 := Call(__e, PrimFunc(symhead), V1825)


let__10047 := tmp10843
_ = let__10047

tmp10844 := Call(__e, PrimFunc(symtail), V1825)


let__10048 := tmp10844
_ = let__10048

tmp10848 := Call(__e, PrimFunc(symatom_2), let__10047)


var ifres10845 Obj

if True == tmp10848 {
tmp10846 := Call(__e, PrimFunc(symshen_4comb), let__10048, let__10047)


ifres10845 = tmp10846


} else {
tmp10847 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10845 = tmp10847


}

ifres10842 = ifres10845


} else {
tmp10849 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10842 = tmp10849


}

let__10046 := ifres10842
_ = let__10046

tmp10871 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10046)


if True == tmp10871 {
tmp10867 := Call(__e, PrimFunc(symshen_4ccons_2), V1825)


var ifres10851 Obj

if True == tmp10867 {
tmp10852 := Call(__e, PrimFunc(symhead), V1825)


let__10050 := tmp10852
_ = let__10050

tmp10853 := Call(__e, PrimFunc(symtail), V1825)


let__10051 := tmp10853
_ = let__10051

tmp10854 := Call(__e, PrimFunc(symshen_4_5bterms_6), let__10050)


let__10052 := tmp10854
_ = let__10052

tmp10865 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10052)


var ifres10855 Obj

if True == tmp10865 {
tmp10856 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10855 = tmp10856


} else {
tmp10857 := Call(__e, PrimFunc(symshen_4_5_1out), let__10052)


let__10053 := tmp10857
_ = let__10053

tmp10858 := Call(__e, PrimFunc(symshen_4in_1_6), let__10052)


let__10054 := tmp10858
_ = let__10054

tmp10859 := Call(__e, PrimFunc(sym_5end_6), let__10054)


let__10055 := tmp10859
_ = let__10055

tmp10864 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10055)


var ifres10860 Obj

if True == tmp10864 {
tmp10861 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10860 = tmp10861


} else {
tmp10862 := Call(__e, PrimFunc(symshen_4in_1_6), let__10055)


let__10056 := tmp10862
_ = let__10056

tmp10863 := Call(__e, PrimFunc(symshen_4comb), let__10051, let__10053)


ifres10860 = tmp10863


}

ifres10855 = ifres10860


}

ifres10851 = ifres10855


} else {
tmp10866 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10851 = tmp10866


}

let__10049 := ifres10851
_ = let__10049

tmp10869 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10049)


if True == tmp10869 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__10049)
return
}


} else {
__e.Return(let__10046)
return
}


} else {
__e.Return(let__10042)
return
}


}, 1)

tmp10874 := Call(__e, ns2_1set, symshen_4_5bterm_6, tmp10834)


_ = tmp10874

tmp10875 := MakeNative(func(__e *ControlFlow) {
V1841 := __e.Get(1)
_ = V1841
tmp10885 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1841)
}
__typedArg0 := V1841
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10876 Obj

if True == tmp10885 {
tmp10877 := Call(__e, PrimFunc(symhead), V1841)


let__10058 := tmp10877
_ = let__10058

tmp10878 := Call(__e, PrimFunc(symtail), V1841)


let__10059 := tmp10878
_ = let__10059

tmp10883 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__10058, sym__)
}
__typedArg0 := let__10058
__typedArg1 := sym__
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10879 Obj

if True == tmp10883 {
tmp10880 := Call(__e, PrimFunc(symgensym), symY)


tmp10881 := Call(__e, PrimFunc(symshen_4comb), let__10059, tmp10880)


ifres10879 = tmp10881


} else {
tmp10882 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10879 = tmp10882


}

ifres10876 = ifres10879


} else {
tmp10884 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10876 = tmp10884


}

let__10057 := ifres10876
_ = let__10057

tmp10887 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10057)


if True == tmp10887 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__10057)
return
}


}, 1)

tmp10888 := Call(__e, ns2_1set, symshen_4_5wildcard_6, tmp10875)


_ = tmp10888

tmp10889 := MakeNative(func(__e *ControlFlow) {
V1845 := __e.Get(1)
_ = V1845
tmp10898 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1845)
}
__typedArg0 := V1845
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10890 Obj

if True == tmp10898 {
tmp10891 := Call(__e, PrimFunc(symhead), V1845)


let__10061 := tmp10891
_ = let__10061

tmp10892 := Call(__e, PrimFunc(symtail), V1845)


let__10062 := tmp10892
_ = let__10062

tmp10896 := Call(__e, PrimFunc(symshen_4semicolon_2), let__10061)


var ifres10893 Obj

if True == tmp10896 {
tmp10894 := Call(__e, PrimFunc(symshen_4comb), let__10062, let__10061)


ifres10893 = tmp10894


} else {
tmp10895 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10893 = tmp10895


}

ifres10890 = ifres10893


} else {
tmp10897 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10890 = tmp10897


}

let__10060 := ifres10890
_ = let__10060

tmp10900 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10060)


if True == tmp10900 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__10060)
return
}


}, 1)

tmp10901 := Call(__e, ns2_1set, symshen_4_5sc_6, tmp10889)


_ = tmp10901

tmp10902 := MakeNative(func(__e *ControlFlow) {
V1849 := __e.Get(1)
_ = V1849
V1850 := __e.Get(2)
_ = V1850
tmp10903 := Call(__e, PrimFunc(symgensym), symB)


let__10063 := tmp10903
_ = let__10063

tmp10904 := Call(__e, PrimFunc(symgensym), symL)


let__10064 := tmp10904
_ = let__10064

tmp10905 := Call(__e, PrimFunc(symgensym), symK)


let__10065 := tmp10905
_ = let__10065

tmp10906 := Call(__e, PrimFunc(symgensym), symC)


let__10066 := tmp10906
_ = let__10066

tmp10907 := Call(__e, PrimFunc(symshen_4prolog_1parameters), V1850)


let__10067 := tmp10907
_ = let__10067

tmp10908 := Call(__e, PrimFunc(symshen_4hascut_2), V1850)


let__10068 := tmp10908
_ = let__10068

tmp10909 := Call(__e, PrimFunc(symshen_4prolog_1fbody), V1850, let__10067, let__10063, let__10064, let__10065, let__10066, let__10068)


let__10069 := tmp10909
_ = let__10069

var ifres10910 Obj

if True == let__10068 {
tmp10911 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), Nil)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10912 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10065, tmp10911)
}
__typedArg0 := let__10065
__typedArg1 := tmp10911
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10913 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_7, tmp10912)
}
__typedArg0 := sym_7
__typedArg1 := tmp10912
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10914 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10069, Nil)
}
__typedArg0 := let__10069
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10915 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp10913, tmp10914)
}
__typedArg0 := tmp10913
__typedArg1 := tmp10914
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10916 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10065, tmp10915)
}
__typedArg0 := let__10065
__typedArg1 := tmp10915
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10917 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp10916)
}
__typedArg0 := symlet
__typedArg1 := tmp10916
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

ifres10910 = tmp10917


} else {
ifres10910 = let__10069


}

let__10070 := ifres10910
_ = let__10070

tmp10918 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_6, Nil)
}
__typedArg0 := sym_1_6
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10919 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10066, tmp10918)
}
__typedArg0 := let__10066
__typedArg1 := tmp10918
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10920 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10065, tmp10919)
}
__typedArg0 := let__10065
__typedArg1 := tmp10919
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10921 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10064, tmp10920)
}
__typedArg0 := let__10064
__typedArg1 := tmp10920
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10922 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10063, tmp10921)
}
__typedArg0 := let__10063
__typedArg1 := tmp10921
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10923 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10070, Nil)
}
__typedArg0 := let__10070
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10924 := Call(__e, PrimFunc(symappend), tmp10922, tmp10923)


tmp10925 := Call(__e, PrimFunc(symappend), let__10067, tmp10924)


tmp10926 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1849, tmp10925)
}
__typedArg0 := V1849
__typedArg1 := tmp10925
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10927 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefine, tmp10926)
}
__typedArg0 := symdefine
__typedArg1 := tmp10926
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__10071 := tmp10927
_ = let__10071

__e.Return(let__10071)
return


}, 2)

tmp10928 := Call(__e, ns2_1set, symshen_4horn_1clause_1procedure, tmp10902)


_ = tmp10928

tmp10929 := MakeNative(func(__e *ControlFlow) {
V1862 := __e.Get(1)
_ = V1862
tmp10939 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_b, V1862)
}
__typedArg0 := sym_b
__typedArg1 := V1862
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp10939 {
__e.Return(True)
return
} else {
tmp10937 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1862)
}
__typedArg0 := V1862
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp10937 {
tmp10934 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1862)
}
__typedArg0 := V1862
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10935 := Call(__e, PrimFunc(symshen_4hascut_2), tmp10934)


if True == tmp10935 {
__e.Return(True)
return
} else {
tmp10931 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1862)
}
__typedArg0 := V1862
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10932 := Call(__e, PrimFunc(symshen_4hascut_2), tmp10931)


if True == tmp10932 {
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

tmp10940 := Call(__e, ns2_1set, symshen_4hascut_2, tmp10929)


_ = tmp10940

tmp10941 := MakeNative(func(__e *ControlFlow) {
V1867 := __e.Get(1)
_ = V1867
tmp10950 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1867)
}
__typedArg0 := V1867
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10946 Obj

if True == tmp10950 {
tmp10948 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1867)
}
__typedArg0 := V1867
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10949 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10948)
}
__typedArg0 := tmp10948
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10947 Obj

if True == tmp10949 {
ifres10947 = True


} else {
ifres10947 = False


}

ifres10946 = ifres10947


} else {
ifres10946 = False


}

if True == ifres10946 {
tmp10942 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1867)
}
__typedArg0 := V1867
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10943 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10942)
}
__typedArg0 := tmp10942
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10944 := Call(__e, PrimFunc(symlength), tmp10943)


__e.TailApply(PrimFunc(symshen_4parameters), tmp10944)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4prolog_1parameters)
return
}


}, 1)

tmp10951 := Call(__e, ns2_1set, symshen_4prolog_1parameters, tmp10941)


_ = tmp10951

tmp10952 := MakeNative(func(__e *ControlFlow) {
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
tmp11042 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V1888)
}
__typedArg0 := Nil
__typedArg1 := V1888
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11039 Obj

if True == tmp11042 {
tmp11041 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(True, V1894)
}
__typedArg0 := True
__typedArg1 := V1894
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11040 Obj

if True == tmp11041 {
ifres11040 = True


} else {
ifres11040 = False


}

ifres11039 = ifres11040


} else {
ifres11039 = False


}

if True == ifres11039 {
tmp10953 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1892, Nil)
}
__typedArg0 := V1892
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10954 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1891, tmp10953)
}
__typedArg0 := V1891
__typedArg1 := tmp10953
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4unlock, tmp10954)
}
__typedArg0 := symshen_4unlock
__typedArg1 := tmp10954
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp11037 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1888)
}
__typedArg0 := V1888
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11015 Obj

if True == tmp11037 {
tmp11035 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1888)
}
__typedArg0 := V1888
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11036 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11035)
}
__typedArg0 := tmp11035
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11017 Obj

if True == tmp11036 {
tmp11032 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1888)
}
__typedArg0 := V1888
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11033 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11032)
}
__typedArg0 := tmp11032
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11034 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11033)
}
__typedArg0 := tmp11033
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11019 Obj

if True == tmp11034 {
tmp11028 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1888)
}
__typedArg0 := V1888
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11029 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11028)
}
__typedArg0 := tmp11028
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11030 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11029)
}
__typedArg0 := tmp11029
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11031 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp11030)
}
__typedArg0 := Nil
__typedArg1 := tmp11030
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11021 Obj

if True == tmp11031 {
tmp11026 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1888)
}
__typedArg0 := V1888
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11027 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp11026)
}
__typedArg0 := Nil
__typedArg1 := tmp11026
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11023 Obj

if True == tmp11027 {
tmp11025 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(False, V1894)
}
__typedArg0 := False
__typedArg1 := V1894
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11024 Obj

if True == tmp11025 {
ifres11024 = True


} else {
ifres11024 = False


}

ifres11023 = ifres11024


} else {
ifres11023 = False


}

var ifres11022 Obj

if True == ifres11023 {
ifres11022 = True


} else {
ifres11022 = False


}

ifres11021 = ifres11022


} else {
ifres11021 = False


}

var ifres11020 Obj

if True == ifres11021 {
ifres11020 = True


} else {
ifres11020 = False


}

ifres11019 = ifres11020


} else {
ifres11019 = False


}

var ifres11018 Obj

if True == ifres11019 {
ifres11018 = True


} else {
ifres11018 = False


}

ifres11017 = ifres11018


} else {
ifres11017 = False


}

var ifres11016 Obj

if True == ifres11017 {
ifres11016 = True


} else {
ifres11016 = False


}

ifres11015 = ifres11016


} else {
ifres11015 = False


}

if True == ifres11015 {
tmp10955 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1888)
}
__typedArg0 := V1888
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10956 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10955)
}
__typedArg0 := tmp10955
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10957 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1888)
}
__typedArg0 := V1888
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10958 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10957)
}
__typedArg0 := tmp10957
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10959 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10958)
}
__typedArg0 := tmp10958
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10960 := Call(__e, PrimFunc(symshen_4continue), tmp10956, tmp10959, V1890, V1891, V1892, V1893)


let__10072 := tmp10960
_ = let__10072

tmp10961 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1891, Nil)
}
__typedArg0 := V1891
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10962 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4unlocked_2, tmp10961)
}
__typedArg0 := symshen_4unlocked_2
__typedArg1 := tmp10961
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10963 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1888)
}
__typedArg0 := V1888
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10964 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10963)
}
__typedArg0 := tmp10963
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10965 := Call(__e, PrimFunc(symshen_4compile_1head), symshen_4_7m, tmp10964, V1889, V1890, let__10072)


tmp10966 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(False, Nil)
}
__typedArg0 := False
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10967 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp10965, tmp10966)
}
__typedArg0 := tmp10965
__typedArg1 := tmp10966
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10968 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp10962, tmp10967)
}
__typedArg0 := tmp10962
__typedArg1 := tmp10967
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp10968)
}
__typedArg0 := symif
__typedArg1 := tmp10968
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp11013 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1888)
}
__typedArg0 := V1888
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10998 Obj

if True == tmp11013 {
tmp11011 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1888)
}
__typedArg0 := V1888
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11012 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11011)
}
__typedArg0 := tmp11011
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11000 Obj

if True == tmp11012 {
tmp11008 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1888)
}
__typedArg0 := V1888
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11009 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11008)
}
__typedArg0 := tmp11008
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11010 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11009)
}
__typedArg0 := tmp11009
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11002 Obj

if True == tmp11010 {
tmp11004 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1888)
}
__typedArg0 := V1888
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11005 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11004)
}
__typedArg0 := tmp11004
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11006 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11005)
}
__typedArg0 := tmp11005
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11007 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp11006)
}
__typedArg0 := Nil
__typedArg1 := tmp11006
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11003 Obj

if True == tmp11007 {
ifres11003 = True


} else {
ifres11003 = False


}

ifres11002 = ifres11003


} else {
ifres11002 = False


}

var ifres11001 Obj

if True == ifres11002 {
ifres11001 = True


} else {
ifres11001 = False


}

ifres11000 = ifres11001


} else {
ifres11000 = False


}

var ifres10999 Obj

if True == ifres11000 {
ifres10999 = True


} else {
ifres10999 = False


}

ifres10998 = ifres10999


} else {
ifres10998 = False


}

if True == ifres10998 {
tmp10969 := Call(__e, PrimFunc(symgensym), symC)


let__10073 := tmp10969
_ = let__10073

tmp10970 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1888)
}
__typedArg0 := V1888
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10971 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10970)
}
__typedArg0 := tmp10970
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10972 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1888)
}
__typedArg0 := V1888
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10973 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10972)
}
__typedArg0 := tmp10972
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10974 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10973)
}
__typedArg0 := tmp10973
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10975 := Call(__e, PrimFunc(symshen_4continue), tmp10971, tmp10974, V1890, V1891, V1892, V1893)


let__10074 := tmp10975
_ = let__10074

tmp10976 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1891, Nil)
}
__typedArg0 := V1891
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10977 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4unlocked_2, tmp10976)
}
__typedArg0 := symshen_4unlocked_2
__typedArg1 := tmp10976
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10978 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1888)
}
__typedArg0 := V1888
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10979 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10978)
}
__typedArg0 := tmp10978
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10980 := Call(__e, PrimFunc(symshen_4compile_1head), symshen_4_7m, tmp10979, V1889, V1890, let__10074)


tmp10981 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(False, Nil)
}
__typedArg0 := False
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10982 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp10980, tmp10981)
}
__typedArg0 := tmp10980
__typedArg1 := tmp10981
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10983 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp10977, tmp10982)
}
__typedArg0 := tmp10977
__typedArg1 := tmp10982
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10984 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp10983)
}
__typedArg0 := symif
__typedArg1 := tmp10983
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10985 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(False, Nil)
}
__typedArg0 := False
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10986 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10073, tmp10985)
}
__typedArg0 := let__10073
__typedArg1 := tmp10985
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10987 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_a, tmp10986)
}
__typedArg0 := sym_a
__typedArg1 := tmp10986
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10988 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1888)
}
__typedArg0 := V1888
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10989 := Call(__e, PrimFunc(symshen_4prolog_1fbody), tmp10988, V1889, V1890, V1891, V1892, V1893, V1894)


tmp10990 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10073, Nil)
}
__typedArg0 := let__10073
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10991 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp10989, tmp10990)
}
__typedArg0 := tmp10989
__typedArg1 := tmp10990
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10992 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp10987, tmp10991)
}
__typedArg0 := tmp10987
__typedArg1 := tmp10991
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10993 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp10992)
}
__typedArg0 := symif
__typedArg1 := tmp10992
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10994 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp10993, Nil)
}
__typedArg0 := tmp10993
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10995 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp10984, tmp10994)
}
__typedArg0 := tmp10984
__typedArg1 := tmp10994
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10996 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10073, tmp10995)
}
__typedArg0 := let__10073
__typedArg1 := tmp10995
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp10996)
}
__typedArg0 := symlet
__typedArg1 := tmp10996
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

tmp11043 := Call(__e, ns2_1set, symshen_4prolog_1fbody, tmp10952)


_ = tmp11043

tmp11044 := MakeNative(func(__e *ControlFlow) {
V1898 := __e.Get(1)
_ = V1898
V1899 := __e.Get(2)
_ = V1899
tmp11049 := Call(__e, PrimFunc(symshen_4locked_2), V1898)


var ifres11046 Obj

if True == tmp11049 {
tmp11048 := Call(__e, PrimFunc(symshen_4fits_2), V1899, V1898)


var ifres11047 Obj

if True == tmp11048 {
ifres11047 = True


} else {
ifres11047 = False


}

ifres11046 = ifres11047


} else {
ifres11046 = False


}

if True == ifres11046 {
__e.TailApply(PrimFunc(symshen_4openlock), V1898)
return
} else {
__e.Return(False)
return
}


}, 2)

tmp11050 := Call(__e, ns2_1set, symshen_4unlock, tmp11044)


_ = tmp11050

tmp11051 := MakeNative(func(__e *ControlFlow) {
V1900 := __e.Get(1)
_ = V1900
tmp11052 := Call(__e, PrimFunc(symshen_4unlocked_2), V1900)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp11052)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp11052
return Call(__e, PrimFunc(symnot), __typedArg0)
})())
return


}, 1)

tmp11053 := Call(__e, ns2_1set, symshen_4locked_2, tmp11051)


_ = tmp11053

tmp11054 := MakeNative(func(__e *ControlFlow) {
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

tmp11055 := Call(__e, ns2_1set, symshen_4unlocked_2, tmp11054)


_ = tmp11055

tmp11056 := MakeNative(func(__e *ControlFlow) {
V1902 := __e.Get(1)
_ = V1902
tmp11057 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(V1902, MakeInteger(1), True)
}
__typedArg0 := V1902
__typedArg1 := MakeInteger(1)
__typedArg2 := True
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

_ = tmp11057

__e.Return(False)
return


}, 1)

tmp11058 := Call(__e, ns2_1set, symshen_4openlock, tmp11056)


_ = tmp11058

tmp11059 := MakeNative(func(__e *ControlFlow) {
V1903 := __e.Get(1)
_ = V1903
V1904 := __e.Get(2)
_ = V1904
tmp11060 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V1904, MakeInteger(2))
}
__typedArg0 := V1904
__typedArg1 := MakeInteger(2)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V1903, tmp11060)
}
__typedArg0 := V1903
__typedArg1 := tmp11060
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})())
return


}, 2)

tmp11061 := Call(__e, ns2_1set, symshen_4fits_2, tmp11059)


_ = tmp11061

tmp11062 := MakeNative(func(__e *ControlFlow) {
V1907 := __e.Get(1)
_ = V1907
V1908 := __e.Get(2)
_ = V1908
V1909 := __e.Get(3)
_ = V1909
V1910 := __e.Get(4)
_ = V1910
tmp11063 := Call(__e, PrimFunc(symthaw), V1910)


let__10075 := tmp11063
_ = let__10075

tmp11068 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__10075, False)
}
__typedArg0 := let__10075
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11065 Obj

if True == tmp11068 {
tmp11067 := Call(__e, PrimFunc(symshen_4unlocked_2), V1908)


var ifres11066 Obj

if True == tmp11067 {
ifres11066 = True


} else {
ifres11066 = False


}

ifres11065 = ifres11066


} else {
ifres11065 = False


}

if True == ifres11065 {
__e.TailApply(PrimFunc(symshen_4lock), V1909, V1908)
return
} else {
__e.Return(let__10075)
return
}


}, 4)

tmp11069 := Call(__e, ns2_1set, symshen_4cut, tmp11062)


_ = tmp11069

tmp11070 := MakeNative(func(__e *ControlFlow) {
V1912 := __e.Get(1)
_ = V1912
V1913 := __e.Get(2)
_ = V1913
tmp11071 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(V1913, MakeInteger(1), False)
}
__typedArg0 := V1913
__typedArg1 := MakeInteger(1)
__typedArg2 := False
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

let__10076 := tmp11071
_ = let__10076

tmp11072 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(V1913, MakeInteger(2), V1912)
}
__typedArg0 := V1913
__typedArg1 := MakeInteger(2)
__typedArg2 := V1912
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

let__10077 := tmp11072
_ = let__10077

__e.Return(False)
return


}, 2)

tmp11073 := Call(__e, ns2_1set, symshen_4lock, tmp11070)


_ = tmp11073

tmp11074 := MakeNative(func(__e *ControlFlow) {
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
tmp11075 := Call(__e, PrimFunc(symshen_4extract_1vars), V1916)


let__10078 := tmp11075
_ = let__10078

tmp11076 := Call(__e, PrimFunc(symshen_4extract_1free_1vars), V1917)


let__10079 := tmp11076
_ = let__10079

tmp11077 := Call(__e, PrimFunc(symdifference), let__10079, let__10078)


let__10080 := tmp11077
_ = let__10080

tmp11078 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4incinfs, Nil)
}
__typedArg0 := symshen_4incinfs
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11079 := Call(__e, PrimFunc(symshen_4compile_1body), V1917, V1918, V1919, V1920, V1921)


tmp11080 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11079, Nil)
}
__typedArg0 := tmp11079
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11081 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11078, tmp11080)
}
__typedArg0 := tmp11078
__typedArg1 := tmp11080
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11082 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdo, tmp11081)
}
__typedArg0 := symdo
__typedArg1 := tmp11081
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__10081 := tmp11082
_ = let__10081

__e.TailApply(PrimFunc(symshen_4stpart), let__10080, let__10081, V1918)
return


}, 6)

tmp11083 := Call(__e, ns2_1set, symshen_4continue, tmp11074)


_ = tmp11083

tmp11084 := MakeNative(func(__e *ControlFlow) {
V1928 := __e.Get(1)
_ = V1928
tmp11119 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1928)
}
__typedArg0 := V1928
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11100 Obj

if True == tmp11119 {
tmp11117 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1928)
}
__typedArg0 := V1928
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11118 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlambda, tmp11117)
}
__typedArg0 := symlambda
__typedArg1 := tmp11117
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11102 Obj

if True == tmp11118 {
tmp11115 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1928)
}
__typedArg0 := V1928
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11116 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11115)
}
__typedArg0 := tmp11115
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11104 Obj

if True == tmp11116 {
tmp11112 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1928)
}
__typedArg0 := V1928
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11113 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11112)
}
__typedArg0 := tmp11112
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11114 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11113)
}
__typedArg0 := tmp11113
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11106 Obj

if True == tmp11114 {
tmp11108 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1928)
}
__typedArg0 := V1928
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11109 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11108)
}
__typedArg0 := tmp11108
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11110 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11109)
}
__typedArg0 := tmp11109
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11111 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp11110)
}
__typedArg0 := Nil
__typedArg1 := tmp11110
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11107 Obj

if True == tmp11111 {
ifres11107 = True


} else {
ifres11107 = False


}

ifres11106 = ifres11107


} else {
ifres11106 = False


}

var ifres11105 Obj

if True == ifres11106 {
ifres11105 = True


} else {
ifres11105 = False


}

ifres11104 = ifres11105


} else {
ifres11104 = False


}

var ifres11103 Obj

if True == ifres11104 {
ifres11103 = True


} else {
ifres11103 = False


}

ifres11102 = ifres11103


} else {
ifres11102 = False


}

var ifres11101 Obj

if True == ifres11102 {
ifres11101 = True


} else {
ifres11101 = False


}

ifres11100 = ifres11101


} else {
ifres11100 = False


}

if True == ifres11100 {
tmp11085 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1928)
}
__typedArg0 := V1928
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11086 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11085)
}
__typedArg0 := tmp11085
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11087 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1928)
}
__typedArg0 := V1928
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11088 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11087)
}
__typedArg0 := tmp11087
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11089 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11088)
}
__typedArg0 := tmp11088
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11090 := Call(__e, PrimFunc(symshen_4extract_1free_1vars), tmp11089)


__e.TailApply(PrimFunc(symremove), tmp11086, tmp11090)
return


} else {
tmp11098 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1928)
}
__typedArg0 := V1928
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp11098 {
tmp11091 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1928)
}
__typedArg0 := V1928
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11092 := Call(__e, PrimFunc(symshen_4extract_1free_1vars), tmp11091)


tmp11093 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1928)
}
__typedArg0 := V1928
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11094 := Call(__e, PrimFunc(symshen_4extract_1free_1vars), tmp11093)


__e.TailApply(PrimFunc(symunion), tmp11092, tmp11094)
return


} else {
tmp11096 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(V1928)
}
__typedArg0 := V1928
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

if True == tmp11096 {
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

tmp11120 := Call(__e, ns2_1set, symshen_4extract_1free_1vars, tmp11084)


_ = tmp11120

tmp11121 := MakeNative(func(__e *ControlFlow) {
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
tmp11155 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V1945)
}
__typedArg0 := Nil
__typedArg1 := V1945
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp11155 {
tmp11122 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1949, Nil)
}
__typedArg0 := V1949
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symthaw, tmp11122)
}
__typedArg0 := symthaw
__typedArg1 := tmp11122
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp11153 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1945)
}
__typedArg0 := V1945
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11149 Obj

if True == tmp11153 {
tmp11151 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1945)
}
__typedArg0 := V1945
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11152 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_b, tmp11151)
}
__typedArg0 := sym_b
__typedArg1 := tmp11151
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11150 Obj

if True == tmp11152 {
ifres11150 = True


} else {
ifres11150 = False


}

ifres11149 = ifres11150


} else {
ifres11149 = False


}

if True == ifres11149 {
tmp11123 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4cut, Nil)
}
__typedArg0 := symshen_4cut
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11124 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1945)
}
__typedArg0 := V1945
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11125 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11123, tmp11124)
}
__typedArg0 := tmp11123
__typedArg1 := tmp11124
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

if PrimFunc(symshen_4compile_1body) == __self {
__self1, __self2, __self3, __self4, __self5 = tmp11125, V1946, V1947, V1948, V1949
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4compile_1body), tmp11125, V1946, V1947, V1948, V1949)
return


} else {
tmp11147 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1945)
}
__typedArg0 := V1945
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11143 Obj

if True == tmp11147 {
tmp11145 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1945)
}
__typedArg0 := V1945
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11146 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp11145)
}
__typedArg0 := Nil
__typedArg1 := tmp11145
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11144 Obj

if True == tmp11146 {
ifres11144 = True


} else {
ifres11144 = False


}

ifres11143 = ifres11144


} else {
ifres11143 = False


}

if True == ifres11143 {
tmp11126 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1945)
}
__typedArg0 := V1945
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11127 := Call(__e, PrimFunc(symshen_4deref_1calls), tmp11126, V1946)


tmp11128 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1949, Nil)
}
__typedArg0 := V1949
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11129 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1948, tmp11128)
}
__typedArg0 := V1948
__typedArg1 := tmp11128
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11130 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1947, tmp11129)
}
__typedArg0 := V1947
__typedArg1 := tmp11129
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11131 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1946, tmp11130)
}
__typedArg0 := V1946
__typedArg1 := tmp11130
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symappend), tmp11127, tmp11131)
return


} else {
tmp11141 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1945)
}
__typedArg0 := V1945
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp11141 {
tmp11132 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1945)
}
__typedArg0 := V1945
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11133 := Call(__e, PrimFunc(symshen_4deref_1calls), tmp11132, V1946)


let__10082 := tmp11133
_ = let__10082

tmp11134 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1945)
}
__typedArg0 := V1945
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11135 := Call(__e, PrimFunc(symshen_4freeze_1literals), tmp11134, V1946, V1947, V1948, V1949)


tmp11136 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11135, Nil)
}
__typedArg0 := tmp11135
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11137 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1948, tmp11136)
}
__typedArg0 := V1948
__typedArg1 := tmp11136
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11138 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1947, tmp11137)
}
__typedArg0 := V1947
__typedArg1 := tmp11137
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11139 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1946, tmp11138)
}
__typedArg0 := V1946
__typedArg1 := tmp11138
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symappend), let__10082, tmp11139)
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

tmp11156 := Call(__e, ns2_1set, symshen_4compile_1body, tmp11121)


_ = tmp11156

tmp11157 := MakeNative(func(__e *ControlFlow) {
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
tmp11180 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V1967)
}
__typedArg0 := Nil
__typedArg1 := V1967
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp11180 {
__e.Return(V1971)
return
} else {
tmp11178 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1967)
}
__typedArg0 := V1967
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11174 Obj

if True == tmp11178 {
tmp11176 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1967)
}
__typedArg0 := V1967
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11177 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_b, tmp11176)
}
__typedArg0 := sym_b
__typedArg1 := tmp11176
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11175 Obj

if True == tmp11177 {
ifres11175 = True


} else {
ifres11175 = False


}

ifres11174 = ifres11175


} else {
ifres11174 = False


}

if True == ifres11174 {
tmp11158 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4cut, Nil)
}
__typedArg0 := symshen_4cut
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11159 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1967)
}
__typedArg0 := V1967
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11160 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11158, tmp11159)
}
__typedArg0 := tmp11158
__typedArg1 := tmp11159
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

if PrimFunc(symshen_4freeze_1literals) == __self {
__self1, __self2, __self3, __self4, __self5 = tmp11160, V1968, V1969, V1970, V1971
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4freeze_1literals), tmp11160, V1968, V1969, V1970, V1971)
return


} else {
tmp11172 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1967)
}
__typedArg0 := V1967
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp11172 {
tmp11161 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1967)
}
__typedArg0 := V1967
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11162 := Call(__e, PrimFunc(symshen_4deref_1calls), tmp11161, V1968)


let__10083 := tmp11162
_ = let__10083

tmp11163 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1967)
}
__typedArg0 := V1967
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11164 := Call(__e, PrimFunc(symshen_4freeze_1literals), tmp11163, V1968, V1969, V1970, V1971)


tmp11165 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11164, Nil)
}
__typedArg0 := tmp11164
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11166 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1970, tmp11165)
}
__typedArg0 := V1970
__typedArg1 := tmp11165
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11167 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1969, tmp11166)
}
__typedArg0 := V1969
__typedArg1 := tmp11166
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11168 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1968, tmp11167)
}
__typedArg0 := V1968
__typedArg1 := tmp11167
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11169 := Call(__e, PrimFunc(symappend), let__10083, tmp11168)


tmp11170 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11169, Nil)
}
__typedArg0 := tmp11169
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfreeze, tmp11170)
}
__typedArg0 := symfreeze
__typedArg1 := tmp11170
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

tmp11181 := Call(__e, ns2_1set, symshen_4freeze_1literals, tmp11157)


_ = tmp11181

tmp11182 := MakeNative(func(__e *ControlFlow) {
V1977 := __e.Get(1)
_ = V1977
V1978 := __e.Get(2)
_ = V1978
tmp11197 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1977)
}
__typedArg0 := V1977
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11193 Obj

if True == tmp11197 {
tmp11195 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1977)
}
__typedArg0 := V1977
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11196 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symfork, tmp11195)
}
__typedArg0 := symfork
__typedArg1 := tmp11195
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11194 Obj

if True == tmp11196 {
ifres11194 = True


} else {
ifres11194 = False


}

ifres11193 = ifres11194


} else {
ifres11193 = False


}

if True == ifres11193 {
tmp11183 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1977)
}
__typedArg0 := V1977
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11184 := Call(__e, PrimFunc(symshen_4deref_1forked_1literals), tmp11183, V1978)


tmp11185 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11184, Nil)
}
__typedArg0 := tmp11184
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfork, tmp11185)
}
__typedArg0 := symfork
__typedArg1 := tmp11185
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp11191 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1977)
}
__typedArg0 := V1977
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp11191 {
tmp11186 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1977)
}
__typedArg0 := V1977
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11187 := MakeNative(func(__e *ControlFlow) {
Z1979 := __e.Get(1)
_ = Z1979
__e.TailApply(PrimFunc(symshen_4function_1calls), Z1979, V1978)
return
}, 1)

tmp11188 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1977)
}
__typedArg0 := V1977
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11189 := Call(__e, PrimFunc(symmap), tmp11187, tmp11188)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11186, tmp11189)
}
__typedArg0 := tmp11186
__typedArg1 := tmp11189
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

tmp11198 := Call(__e, ns2_1set, symshen_4deref_1calls, tmp11182)


_ = tmp11198

tmp11199 := MakeNative(func(__e *ControlFlow) {
V1986 := __e.Get(1)
_ = V1986
V1987 := __e.Get(2)
_ = V1987
tmp11209 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V1986)
}
__typedArg0 := Nil
__typedArg1 := V1986
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp11209 {
__e.Return(Nil)
return
} else {
tmp11207 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1986)
}
__typedArg0 := V1986
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp11207 {
tmp11200 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1986)
}
__typedArg0 := V1986
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11201 := Call(__e, PrimFunc(symshen_4deref_1calls), tmp11200, V1987)


tmp11202 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1986)
}
__typedArg0 := V1986
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11203 := Call(__e, PrimFunc(symshen_4deref_1forked_1literals), tmp11202, V1987)


tmp11204 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11203, Nil)
}
__typedArg0 := tmp11203
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11205 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11201, tmp11204)
}
__typedArg0 := tmp11201
__typedArg1 := tmp11204
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp11205)
}
__typedArg0 := symcons
__typedArg1 := tmp11205
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

tmp11210 := Call(__e, ns2_1set, symshen_4deref_1forked_1literals, tmp11199)


_ = tmp11210

tmp11211 := MakeNative(func(__e *ControlFlow) {
V1990 := __e.Get(1)
_ = V1990
V1991 := __e.Get(2)
_ = V1991
tmp11243 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1990)
}
__typedArg0 := V1990
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11224 Obj

if True == tmp11243 {
tmp11241 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1990)
}
__typedArg0 := V1990
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11242 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcons, tmp11241)
}
__typedArg0 := symcons
__typedArg1 := tmp11241
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11226 Obj

if True == tmp11242 {
tmp11239 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1990)
}
__typedArg0 := V1990
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11240 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11239)
}
__typedArg0 := tmp11239
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11228 Obj

if True == tmp11240 {
tmp11236 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1990)
}
__typedArg0 := V1990
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11237 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11236)
}
__typedArg0 := tmp11236
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11238 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11237)
}
__typedArg0 := tmp11237
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11230 Obj

if True == tmp11238 {
tmp11232 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1990)
}
__typedArg0 := V1990
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11233 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11232)
}
__typedArg0 := tmp11232
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11234 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11233)
}
__typedArg0 := tmp11233
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11235 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp11234)
}
__typedArg0 := Nil
__typedArg1 := tmp11234
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11231 Obj

if True == tmp11235 {
ifres11231 = True


} else {
ifres11231 = False


}

ifres11230 = ifres11231


} else {
ifres11230 = False


}

var ifres11229 Obj

if True == ifres11230 {
ifres11229 = True


} else {
ifres11229 = False


}

ifres11228 = ifres11229


} else {
ifres11228 = False


}

var ifres11227 Obj

if True == ifres11228 {
ifres11227 = True


} else {
ifres11227 = False


}

ifres11226 = ifres11227


} else {
ifres11226 = False


}

var ifres11225 Obj

if True == ifres11226 {
ifres11225 = True


} else {
ifres11225 = False


}

ifres11224 = ifres11225


} else {
ifres11224 = False


}

if True == ifres11224 {
tmp11212 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1990)
}
__typedArg0 := V1990
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11213 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11212)
}
__typedArg0 := tmp11212
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11214 := Call(__e, PrimFunc(symshen_4function_1calls), tmp11213, V1991)


tmp11215 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1990)
}
__typedArg0 := V1990
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11216 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11215)
}
__typedArg0 := tmp11215
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11217 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11216)
}
__typedArg0 := tmp11216
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11218 := Call(__e, PrimFunc(symshen_4function_1calls), tmp11217, V1991)


tmp11219 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11218, Nil)
}
__typedArg0 := tmp11218
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11220 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11214, tmp11219)
}
__typedArg0 := tmp11214
__typedArg1 := tmp11219
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp11220)
}
__typedArg0 := symcons
__typedArg1 := tmp11220
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp11222 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1990)
}
__typedArg0 := V1990
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp11222 {
__e.TailApply(PrimFunc(symshen_4deref_1terms), V1990, V1991, Nil)
return
} else {
__e.Return(V1990)
return
}


}


}, 2)

tmp11244 := Call(__e, ns2_1set, symshen_4function_1calls, tmp11211)


_ = tmp11244

tmp11245 := MakeNative(func(__e *ControlFlow) {
V2000 := __e.Get(1)
_ = V2000
V2001 := __e.Get(2)
_ = V2001
V2002 := __e.Get(3)
_ = V2002
tmp11339 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2000)
}
__typedArg0 := V2000
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11326 Obj

if True == tmp11339 {
tmp11337 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2000)
}
__typedArg0 := V2000
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11338 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(0), tmp11337)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp11337
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11328 Obj

if True == tmp11338 {
tmp11335 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2000)
}
__typedArg0 := V2000
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11336 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11335)
}
__typedArg0 := tmp11335
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11330 Obj

if True == tmp11336 {
tmp11332 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2000)
}
__typedArg0 := V2000
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11333 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11332)
}
__typedArg0 := tmp11332
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11334 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp11333)
}
__typedArg0 := Nil
__typedArg1 := tmp11333
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11331 Obj

if True == tmp11334 {
ifres11331 = True


} else {
ifres11331 = False


}

ifres11330 = ifres11331


} else {
ifres11330 = False


}

var ifres11329 Obj

if True == ifres11330 {
ifres11329 = True


} else {
ifres11329 = False


}

ifres11328 = ifres11329


} else {
ifres11328 = False


}

var ifres11327 Obj

if True == ifres11328 {
ifres11327 = True


} else {
ifres11327 = False


}

ifres11326 = ifres11327


} else {
ifres11326 = False


}

if True == ifres11326 {
tmp11252 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2000)
}
__typedArg0 := V2000
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11253 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11252)
}
__typedArg0 := tmp11252
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11254 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(tmp11253)
}
__typedArg0 := tmp11253
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

if True == tmp11254 {
tmp11246 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2000)
}
__typedArg0 := V2000
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11246)
}
__typedArg0 := tmp11246
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return


} else {
tmp11247 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2000)
}
__typedArg0 := V2000
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11248 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11247)
}
__typedArg0 := tmp11247
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11249 := Call(__e, PrimFunc(symshen_4app), tmp11248, MakeString("\n"), symshen_4s)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("attempt to optimise a non-variable "))
__typedS1, __typedOK1 := TypedString(tmp11249)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("attempt to optimise a non-variable ")
__typedArg1 := tmp11249
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("attempt to optimise a non-variable "))
__typedS1, __typedOK1 := TypedString(tmp11249)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("attempt to optimise a non-variable ")
__typedArg1 := tmp11249
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}


} else {
tmp11324 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2000)
}
__typedArg0 := V2000
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11311 Obj

if True == tmp11324 {
tmp11322 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2000)
}
__typedArg0 := V2000
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11323 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(1), tmp11322)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp11322
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11313 Obj

if True == tmp11323 {
tmp11320 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2000)
}
__typedArg0 := V2000
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11321 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11320)
}
__typedArg0 := tmp11320
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11315 Obj

if True == tmp11321 {
tmp11317 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2000)
}
__typedArg0 := V2000
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11318 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11317)
}
__typedArg0 := tmp11317
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11319 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp11318)
}
__typedArg0 := Nil
__typedArg1 := tmp11318
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11316 Obj

if True == tmp11319 {
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
tmp11264 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2000)
}
__typedArg0 := V2000
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11265 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11264)
}
__typedArg0 := tmp11264
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11266 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(tmp11265)
}
__typedArg0 := tmp11265
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

if True == tmp11266 {
tmp11255 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2000)
}
__typedArg0 := V2000
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11256 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11255)
}
__typedArg0 := tmp11255
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11257 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V2001, Nil)
}
__typedArg0 := V2001
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11258 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11256, tmp11257)
}
__typedArg0 := tmp11256
__typedArg1 := tmp11257
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4lazyderef, tmp11258)
}
__typedArg0 := symshen_4lazyderef
__typedArg1 := tmp11258
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp11259 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2000)
}
__typedArg0 := V2000
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11260 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11259)
}
__typedArg0 := tmp11259
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11261 := Call(__e, PrimFunc(symshen_4app), tmp11260, MakeString("\n"), symshen_4s)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("attempt to optimise a non-variable "))
__typedS1, __typedOK1 := TypedString(tmp11261)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("attempt to optimise a non-variable ")
__typedArg1 := tmp11261
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("attempt to optimise a non-variable "))
__typedS1, __typedOK1 := TypedString(tmp11261)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("attempt to optimise a non-variable ")
__typedArg1 := tmp11261
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}


} else {
tmp11308 := Call(__e, PrimFunc(symelement_2), V2000, V2002)


tmp11309 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp11308)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp11308
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres11305 Obj

if True == tmp11309 {
tmp11307 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(V2000)
}
__typedArg0 := V2000
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

var ifres11306 Obj

if True == tmp11307 {
ifres11306 = True


} else {
ifres11306 = False


}

ifres11305 = ifres11306


} else {
ifres11305 = False


}

if True == ifres11305 {
tmp11267 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V2001, Nil)
}
__typedArg0 := V2001
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11268 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V2000, tmp11267)
}
__typedArg0 := V2000
__typedArg1 := tmp11267
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4deref, tmp11268)
}
__typedArg0 := symshen_4deref
__typedArg1 := tmp11268
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp11303 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2000)
}
__typedArg0 := V2000
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11284 Obj

if True == tmp11303 {
tmp11301 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2000)
}
__typedArg0 := V2000
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11302 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlambda, tmp11301)
}
__typedArg0 := symlambda
__typedArg1 := tmp11301
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11286 Obj

if True == tmp11302 {
tmp11299 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2000)
}
__typedArg0 := V2000
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11300 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11299)
}
__typedArg0 := tmp11299
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11288 Obj

if True == tmp11300 {
tmp11296 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2000)
}
__typedArg0 := V2000
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11297 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11296)
}
__typedArg0 := tmp11296
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11298 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11297)
}
__typedArg0 := tmp11297
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11290 Obj

if True == tmp11298 {
tmp11292 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2000)
}
__typedArg0 := V2000
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11293 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11292)
}
__typedArg0 := tmp11292
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11294 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11293)
}
__typedArg0 := tmp11293
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11295 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp11294)
}
__typedArg0 := Nil
__typedArg1 := tmp11294
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11291 Obj

if True == tmp11295 {
ifres11291 = True


} else {
ifres11291 = False


}

ifres11290 = ifres11291


} else {
ifres11290 = False


}

var ifres11289 Obj

if True == ifres11290 {
ifres11289 = True


} else {
ifres11289 = False


}

ifres11288 = ifres11289


} else {
ifres11288 = False


}

var ifres11287 Obj

if True == ifres11288 {
ifres11287 = True


} else {
ifres11287 = False


}

ifres11286 = ifres11287


} else {
ifres11286 = False


}

var ifres11285 Obj

if True == ifres11286 {
ifres11285 = True


} else {
ifres11285 = False


}

ifres11284 = ifres11285


} else {
ifres11284 = False


}

if True == ifres11284 {
tmp11269 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2000)
}
__typedArg0 := V2000
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11270 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11269)
}
__typedArg0 := tmp11269
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11271 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2000)
}
__typedArg0 := V2000
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11272 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11271)
}
__typedArg0 := tmp11271
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11273 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11272)
}
__typedArg0 := tmp11272
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

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

tmp11276 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11275, V2002)
}
__typedArg0 := tmp11275
__typedArg1 := V2002
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11277 := Call(__e, PrimFunc(symshen_4deref_1terms), tmp11273, V2001, tmp11276)


tmp11278 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11277, Nil)
}
__typedArg0 := tmp11277
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11279 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11270, tmp11278)
}
__typedArg0 := tmp11270
__typedArg1 := tmp11278
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp11279)
}
__typedArg0 := symlambda
__typedArg1 := tmp11279
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp11282 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2000)
}
__typedArg0 := V2000
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp11282 {
tmp11280 := MakeNative(func(__e *ControlFlow) {
Z2003 := __e.Get(1)
_ = Z2003
__e.TailApply(PrimFunc(symshen_4deref_1terms), Z2003, V2001, V2002)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp11280, V2000)
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

tmp11340 := Call(__e, ns2_1set, symshen_4deref_1terms, tmp11245)


_ = tmp11340

tmp11341 := MakeNative(func(__e *ControlFlow) {
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
tmp11517 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V2022)
}
__typedArg0 := Nil
__typedArg1 := V2022
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11514 Obj

if True == tmp11517 {
tmp11516 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V2023)
}
__typedArg0 := Nil
__typedArg1 := V2023
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11515 Obj

if True == tmp11516 {
ifres11515 = True


} else {
ifres11515 = False


}

ifres11514 = ifres11515


} else {
ifres11514 = False


}

if True == ifres11514 {
__e.Return(V2025)
return
} else {
tmp11512 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11492 Obj

if True == tmp11512 {
tmp11510 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11511 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11510)
}
__typedArg0 := tmp11510
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11494 Obj

if True == tmp11511 {
tmp11507 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11508 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11507)
}
__typedArg0 := tmp11507
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11509 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4_7m, tmp11508)
}
__typedArg0 := symshen_4_7m
__typedArg1 := tmp11508
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11496 Obj

if True == tmp11509 {
tmp11504 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11505 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11504)
}
__typedArg0 := tmp11504
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11506 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11505)
}
__typedArg0 := tmp11505
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11498 Obj

if True == tmp11506 {
tmp11500 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11501 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11500)
}
__typedArg0 := tmp11500
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11502 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11501)
}
__typedArg0 := tmp11501
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11503 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp11502)
}
__typedArg0 := Nil
__typedArg1 := tmp11502
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11499 Obj

if True == tmp11503 {
ifres11499 = True


} else {
ifres11499 = False


}

ifres11498 = ifres11499


} else {
ifres11498 = False


}

var ifres11497 Obj

if True == ifres11498 {
ifres11497 = True


} else {
ifres11497 = False


}

ifres11496 = ifres11497


} else {
ifres11496 = False


}

var ifres11495 Obj

if True == ifres11496 {
ifres11495 = True


} else {
ifres11495 = False


}

ifres11494 = ifres11495


} else {
ifres11494 = False


}

var ifres11493 Obj

if True == ifres11494 {
ifres11493 = True


} else {
ifres11493 = False


}

ifres11492 = ifres11493


} else {
ifres11492 = False


}

if True == ifres11492 {
tmp11342 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11343 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11342)
}
__typedArg0 := tmp11342
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11344 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11343)
}
__typedArg0 := tmp11343
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11345 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11346 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V2021, tmp11345)
}
__typedArg0 := V2021
__typedArg1 := tmp11345
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11347 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11344, tmp11346)
}
__typedArg0 := tmp11344
__typedArg1 := tmp11346
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11348 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4_7m, tmp11347)
}
__typedArg0 := symshen_4_7m
__typedArg1 := tmp11347
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

if PrimFunc(symshen_4compile_1head) == __self {
__self1, __self2, __self3, __self4, __self5 = V2021, tmp11348, V2023, V2024, V2025
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4compile_1head), V2021, tmp11348, V2023, V2024, V2025)
return


} else {
tmp11490 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11470 Obj

if True == tmp11490 {
tmp11488 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11489 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11488)
}
__typedArg0 := tmp11488
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11472 Obj

if True == tmp11489 {
tmp11485 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11486 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11485)
}
__typedArg0 := tmp11485
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11487 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4_1m, tmp11486)
}
__typedArg0 := symshen_4_1m
__typedArg1 := tmp11486
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11474 Obj

if True == tmp11487 {
tmp11482 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11483 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11482)
}
__typedArg0 := tmp11482
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11484 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11483)
}
__typedArg0 := tmp11483
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11476 Obj

if True == tmp11484 {
tmp11478 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11479 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11478)
}
__typedArg0 := tmp11478
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11480 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11479)
}
__typedArg0 := tmp11479
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11481 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp11480)
}
__typedArg0 := Nil
__typedArg1 := tmp11480
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11477 Obj

if True == tmp11481 {
ifres11477 = True


} else {
ifres11477 = False


}

ifres11476 = ifres11477


} else {
ifres11476 = False


}

var ifres11475 Obj

if True == ifres11476 {
ifres11475 = True


} else {
ifres11475 = False


}

ifres11474 = ifres11475


} else {
ifres11474 = False


}

var ifres11473 Obj

if True == ifres11474 {
ifres11473 = True


} else {
ifres11473 = False


}

ifres11472 = ifres11473


} else {
ifres11472 = False


}

var ifres11471 Obj

if True == ifres11472 {
ifres11471 = True


} else {
ifres11471 = False


}

ifres11470 = ifres11471


} else {
ifres11470 = False


}

if True == ifres11470 {
tmp11349 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11350 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11349)
}
__typedArg0 := tmp11349
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11351 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11350)
}
__typedArg0 := tmp11350
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11352 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11353 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V2021, tmp11352)
}
__typedArg0 := V2021
__typedArg1 := tmp11352
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11354 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11351, tmp11353)
}
__typedArg0 := tmp11351
__typedArg1 := tmp11353
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11355 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4_1m, tmp11354)
}
__typedArg0 := symshen_4_1m
__typedArg1 := tmp11354
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

if PrimFunc(symshen_4compile_1head) == __self {
__self1, __self2, __self3, __self4, __self5 = V2021, tmp11355, V2023, V2024, V2025
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4compile_1head), V2021, tmp11355, V2023, V2024, V2025)
return


} else {
tmp11468 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11464 Obj

if True == tmp11468 {
tmp11466 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11467 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4_1m, tmp11466)
}
__typedArg0 := symshen_4_1m
__typedArg1 := tmp11466
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11465 Obj

if True == tmp11467 {
ifres11465 = True


} else {
ifres11465 = False


}

ifres11464 = ifres11465


} else {
ifres11464 = False


}

if True == ifres11464 {
tmp11356 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4compile_1head) == __self {
__self1, __self2, __self3, __self4, __self5 = symshen_4_1m, tmp11356, V2023, V2024, V2025
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4compile_1head), symshen_4_1m, tmp11356, V2023, V2024, V2025)
return


} else {
tmp11462 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11458 Obj

if True == tmp11462 {
tmp11460 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11461 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4_7m, tmp11460)
}
__typedArg0 := symshen_4_7m
__typedArg1 := tmp11460
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11459 Obj

if True == tmp11461 {
ifres11459 = True


} else {
ifres11459 = False


}

ifres11458 = ifres11459


} else {
ifres11458 = False


}

if True == ifres11458 {
tmp11357 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4compile_1head) == __self {
__self1, __self2, __self3, __self4, __self5 = symshen_4_7m, tmp11357, V2023, V2024, V2025
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4compile_1head), symshen_4_7m, tmp11357, V2023, V2024, V2025)
return


} else {
tmp11456 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11449 Obj

if True == tmp11456 {
tmp11455 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2023)
}
__typedArg0 := V2023
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11451 Obj

if True == tmp11455 {
tmp11453 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11454 := Call(__e, PrimFunc(symshen_4wildcard_2), tmp11453)


var ifres11452 Obj

if True == tmp11454 {
ifres11452 = True


} else {
ifres11452 = False


}

ifres11451 = ifres11452


} else {
ifres11451 = False


}

var ifres11450 Obj

if True == ifres11451 {
ifres11450 = True


} else {
ifres11450 = False


}

ifres11449 = ifres11450


} else {
ifres11449 = False


}

if True == ifres11449 {
tmp11358 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11359 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2023)
}
__typedArg0 := V2023
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4compile_1head) == __self {
__self1, __self2, __self3, __self4, __self5 = V2021, tmp11358, tmp11359, V2024, V2025
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4compile_1head), V2021, tmp11358, tmp11359, V2024, V2025)
return


} else {
tmp11447 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11443 Obj

if True == tmp11447 {
tmp11445 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11446 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(tmp11445)
}
__typedArg0 := tmp11445
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

var ifres11444 Obj

if True == tmp11446 {
ifres11444 = True


} else {
ifres11444 = False


}

ifres11443 = ifres11444


} else {
ifres11443 = False


}

if True == ifres11443 {
__e.TailApply(PrimFunc(symshen_4variable_1case), V2021, V2022, V2023, V2024, V2025)
return
} else {
tmp11441 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4_1m, V2021)
}
__typedArg0 := symshen_4_1m
__typedArg1 := V2021
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11434 Obj

if True == tmp11441 {
tmp11440 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11436 Obj

if True == tmp11440 {
tmp11438 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11439 := Call(__e, PrimFunc(symatom_2), tmp11438)


var ifres11437 Obj

if True == tmp11439 {
ifres11437 = True


} else {
ifres11437 = False


}

ifres11436 = ifres11437


} else {
ifres11436 = False


}

var ifres11435 Obj

if True == ifres11436 {
ifres11435 = True


} else {
ifres11435 = False


}

ifres11434 = ifres11435


} else {
ifres11434 = False


}

if True == ifres11434 {
__e.TailApply(PrimFunc(symshen_4atom_1case_1minus), V2022, V2023, V2024, V2025)
return
} else {
tmp11432 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4_1m, V2021)
}
__typedArg0 := symshen_4_1m
__typedArg1 := V2021
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11402 Obj

if True == tmp11432 {
tmp11431 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11404 Obj

if True == tmp11431 {
tmp11429 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11430 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11429)
}
__typedArg0 := tmp11429
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11406 Obj

if True == tmp11430 {
tmp11426 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11427 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11426)
}
__typedArg0 := tmp11426
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11428 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcons, tmp11427)
}
__typedArg0 := symcons
__typedArg1 := tmp11427
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11408 Obj

if True == tmp11428 {
tmp11423 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11424 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11423)
}
__typedArg0 := tmp11423
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11425 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11424)
}
__typedArg0 := tmp11424
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11410 Obj

if True == tmp11425 {
tmp11419 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11420 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11419)
}
__typedArg0 := tmp11419
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11421 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11420)
}
__typedArg0 := tmp11420
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11422 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11421)
}
__typedArg0 := tmp11421
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11412 Obj

if True == tmp11422 {
tmp11414 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11415 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11414)
}
__typedArg0 := tmp11414
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11416 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11415)
}
__typedArg0 := tmp11415
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11417 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11416)
}
__typedArg0 := tmp11416
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11418 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp11417)
}
__typedArg0 := Nil
__typedArg1 := tmp11417
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11413 Obj

if True == tmp11418 {
ifres11413 = True


} else {
ifres11413 = False


}

ifres11412 = ifres11413


} else {
ifres11412 = False


}

var ifres11411 Obj

if True == ifres11412 {
ifres11411 = True


} else {
ifres11411 = False


}

ifres11410 = ifres11411


} else {
ifres11410 = False


}

var ifres11409 Obj

if True == ifres11410 {
ifres11409 = True


} else {
ifres11409 = False


}

ifres11408 = ifres11409


} else {
ifres11408 = False


}

var ifres11407 Obj

if True == ifres11408 {
ifres11407 = True


} else {
ifres11407 = False


}

ifres11406 = ifres11407


} else {
ifres11406 = False


}

var ifres11405 Obj

if True == ifres11406 {
ifres11405 = True


} else {
ifres11405 = False


}

ifres11404 = ifres11405


} else {
ifres11404 = False


}

var ifres11403 Obj

if True == ifres11404 {
ifres11403 = True


} else {
ifres11403 = False


}

ifres11402 = ifres11403


} else {
ifres11402 = False


}

if True == ifres11402 {
__e.TailApply(PrimFunc(symshen_4cons_1case_1minus), V2022, V2023, V2024, V2025)
return
} else {
tmp11400 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4_7m, V2021)
}
__typedArg0 := symshen_4_7m
__typedArg1 := V2021
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11393 Obj

if True == tmp11400 {
tmp11399 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11395 Obj

if True == tmp11399 {
tmp11397 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11398 := Call(__e, PrimFunc(symatom_2), tmp11397)


var ifres11396 Obj

if True == tmp11398 {
ifres11396 = True


} else {
ifres11396 = False


}

ifres11395 = ifres11396


} else {
ifres11395 = False


}

var ifres11394 Obj

if True == ifres11395 {
ifres11394 = True


} else {
ifres11394 = False


}

ifres11393 = ifres11394


} else {
ifres11393 = False


}

if True == ifres11393 {
__e.TailApply(PrimFunc(symshen_4atom_1case_1plus), V2022, V2023, V2024, V2025)
return
} else {
tmp11391 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4_7m, V2021)
}
__typedArg0 := symshen_4_7m
__typedArg1 := V2021
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11361 Obj

if True == tmp11391 {
tmp11390 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11363 Obj

if True == tmp11390 {
tmp11388 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11389 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11388)
}
__typedArg0 := tmp11388
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11365 Obj

if True == tmp11389 {
tmp11385 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11386 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11385)
}
__typedArg0 := tmp11385
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11387 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcons, tmp11386)
}
__typedArg0 := symcons
__typedArg1 := tmp11386
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11367 Obj

if True == tmp11387 {
tmp11382 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11383 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11382)
}
__typedArg0 := tmp11382
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11384 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11383)
}
__typedArg0 := tmp11383
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11369 Obj

if True == tmp11384 {
tmp11378 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11379 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11378)
}
__typedArg0 := tmp11378
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11380 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11379)
}
__typedArg0 := tmp11379
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11381 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11380)
}
__typedArg0 := tmp11380
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11371 Obj

if True == tmp11381 {
tmp11373 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11374 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11373)
}
__typedArg0 := tmp11373
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11375 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11374)
}
__typedArg0 := tmp11374
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11376 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11375)
}
__typedArg0 := tmp11375
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11377 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp11376)
}
__typedArg0 := Nil
__typedArg1 := tmp11376
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11372 Obj

if True == tmp11377 {
ifres11372 = True


} else {
ifres11372 = False


}

ifres11371 = ifres11372


} else {
ifres11371 = False


}

var ifres11370 Obj

if True == ifres11371 {
ifres11370 = True


} else {
ifres11370 = False


}

ifres11369 = ifres11370


} else {
ifres11369 = False


}

var ifres11368 Obj

if True == ifres11369 {
ifres11368 = True


} else {
ifres11368 = False


}

ifres11367 = ifres11368


} else {
ifres11367 = False


}

var ifres11366 Obj

if True == ifres11367 {
ifres11366 = True


} else {
ifres11366 = False


}

ifres11365 = ifres11366


} else {
ifres11365 = False


}

var ifres11364 Obj

if True == ifres11365 {
ifres11364 = True


} else {
ifres11364 = False


}

ifres11363 = ifres11364


} else {
ifres11363 = False


}

var ifres11362 Obj

if True == ifres11363 {
ifres11362 = True


} else {
ifres11362 = False


}

ifres11361 = ifres11362


} else {
ifres11361 = False


}

if True == ifres11361 {
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

tmp11518 := Call(__e, ns2_1set, symshen_4compile_1head, tmp11341)


_ = tmp11518

tmp11519 := MakeNative(func(__e *ControlFlow) {
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
tmp11540 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2037)
}
__typedArg0 := V2037
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11537 Obj

if True == tmp11540 {
tmp11539 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2038)
}
__typedArg0 := V2038
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11538 Obj

if True == tmp11539 {
ifres11538 = True


} else {
ifres11538 = False


}

ifres11537 = ifres11538


} else {
ifres11537 = False


}

if True == ifres11537 {
tmp11534 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2038)
}
__typedArg0 := V2038
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11535 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(tmp11534)
}
__typedArg0 := tmp11534
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

if True == tmp11535 {
tmp11520 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2037)
}
__typedArg0 := V2037
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11521 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2038)
}
__typedArg0 := V2038
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11522 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2038)
}
__typedArg0 := V2038
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11523 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2037)
}
__typedArg0 := V2037
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11524 := Call(__e, PrimFunc(symsubst), tmp11522, tmp11523, V2040)


__e.TailApply(PrimFunc(symshen_4compile_1head), V2036, tmp11520, tmp11521, V2039, tmp11524)
return


} else {
tmp11525 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2037)
}
__typedArg0 := V2037
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11526 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2038)
}
__typedArg0 := V2038
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11527 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2037)
}
__typedArg0 := V2037
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11528 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2038)
}
__typedArg0 := V2038
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11529 := Call(__e, PrimFunc(symshen_4compile_1head), V2036, tmp11527, tmp11528, V2039, V2040)


tmp11530 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11529, Nil)
}
__typedArg0 := tmp11529
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11531 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11526, tmp11530)
}
__typedArg0 := tmp11526
__typedArg1 := tmp11530
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11532 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11525, tmp11531)
}
__typedArg0 := tmp11525
__typedArg1 := tmp11531
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp11532)
}
__typedArg0 := symlet
__typedArg1 := tmp11532
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

tmp11541 := Call(__e, ns2_1set, symshen_4variable_1case, tmp11519)


_ = tmp11541

tmp11542 := MakeNative(func(__e *ControlFlow) {
V2049 := __e.Get(1)
_ = V2049
V2050 := __e.Get(2)
_ = V2050
V2051 := __e.Get(3)
_ = V2051
V2052 := __e.Get(4)
_ = V2052
tmp11566 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2049)
}
__typedArg0 := V2049
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11563 Obj

if True == tmp11566 {
tmp11565 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2050)
}
__typedArg0 := V2050
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11564 Obj

if True == tmp11565 {
ifres11564 = True


} else {
ifres11564 = False


}

ifres11563 = ifres11564


} else {
ifres11563 = False


}

if True == ifres11563 {
tmp11543 := Call(__e, PrimFunc(symgensym), symTm)


let__10084 := tmp11543
_ = let__10084

tmp11544 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2050)
}
__typedArg0 := V2050
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11545 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V2051, Nil)
}
__typedArg0 := V2051
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11546 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11544, tmp11545)
}
__typedArg0 := tmp11544
__typedArg1 := tmp11545
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11547 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4lazyderef, tmp11546)
}
__typedArg0 := symshen_4lazyderef
__typedArg1 := tmp11546
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11548 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2049)
}
__typedArg0 := V2049
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11549 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11548, Nil)
}
__typedArg0 := tmp11548
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11550 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10084, tmp11549)
}
__typedArg0 := let__10084
__typedArg1 := tmp11549
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11551 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_a, tmp11550)
}
__typedArg0 := sym_a
__typedArg1 := tmp11550
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11552 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2049)
}
__typedArg0 := V2049
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11553 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2050)
}
__typedArg0 := V2050
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11554 := Call(__e, PrimFunc(symshen_4compile_1head), symshen_4_1m, tmp11552, tmp11553, V2051, V2052)


tmp11555 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(False, Nil)
}
__typedArg0 := False
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11556 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11554, tmp11555)
}
__typedArg0 := tmp11554
__typedArg1 := tmp11555
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11557 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11551, tmp11556)
}
__typedArg0 := tmp11551
__typedArg1 := tmp11556
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11558 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp11557)
}
__typedArg0 := symif
__typedArg1 := tmp11557
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11559 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11558, Nil)
}
__typedArg0 := tmp11558
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11560 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11547, tmp11559)
}
__typedArg0 := tmp11547
__typedArg1 := tmp11559
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11561 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10084, tmp11560)
}
__typedArg0 := let__10084
__typedArg1 := tmp11560
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp11561)
}
__typedArg0 := symlet
__typedArg1 := tmp11561
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

tmp11567 := Call(__e, ns2_1set, symshen_4atom_1case_1minus, tmp11542)


_ = tmp11567

tmp11568 := MakeNative(func(__e *ControlFlow) {
V2062 := __e.Get(1)
_ = V2062
V2063 := __e.Get(2)
_ = V2063
V2064 := __e.Get(3)
_ = V2064
V2065 := __e.Get(4)
_ = V2065
tmp11632 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2062)
}
__typedArg0 := V2062
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11602 Obj

if True == tmp11632 {
tmp11630 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2062)
}
__typedArg0 := V2062
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11631 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11630)
}
__typedArg0 := tmp11630
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11604 Obj

if True == tmp11631 {
tmp11627 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2062)
}
__typedArg0 := V2062
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11628 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11627)
}
__typedArg0 := tmp11627
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11629 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcons, tmp11628)
}
__typedArg0 := symcons
__typedArg1 := tmp11628
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11606 Obj

if True == tmp11629 {
tmp11624 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2062)
}
__typedArg0 := V2062
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11625 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11624)
}
__typedArg0 := tmp11624
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11626 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11625)
}
__typedArg0 := tmp11625
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11608 Obj

if True == tmp11626 {
tmp11620 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2062)
}
__typedArg0 := V2062
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11621 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11620)
}
__typedArg0 := tmp11620
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11622 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11621)
}
__typedArg0 := tmp11621
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11623 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11622)
}
__typedArg0 := tmp11622
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11610 Obj

if True == tmp11623 {
tmp11615 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2062)
}
__typedArg0 := V2062
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11616 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11615)
}
__typedArg0 := tmp11615
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11617 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11616)
}
__typedArg0 := tmp11616
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11618 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11617)
}
__typedArg0 := tmp11617
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11619 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp11618)
}
__typedArg0 := Nil
__typedArg1 := tmp11618
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11612 Obj

if True == tmp11619 {
tmp11614 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2063)
}
__typedArg0 := V2063
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11613 Obj

if True == tmp11614 {
ifres11613 = True


} else {
ifres11613 = False


}

ifres11612 = ifres11613


} else {
ifres11612 = False


}

var ifres11611 Obj

if True == ifres11612 {
ifres11611 = True


} else {
ifres11611 = False


}

ifres11610 = ifres11611


} else {
ifres11610 = False


}

var ifres11609 Obj

if True == ifres11610 {
ifres11609 = True


} else {
ifres11609 = False


}

ifres11608 = ifres11609


} else {
ifres11608 = False


}

var ifres11607 Obj

if True == ifres11608 {
ifres11607 = True


} else {
ifres11607 = False


}

ifres11606 = ifres11607


} else {
ifres11606 = False


}

var ifres11605 Obj

if True == ifres11606 {
ifres11605 = True


} else {
ifres11605 = False


}

ifres11604 = ifres11605


} else {
ifres11604 = False


}

var ifres11603 Obj

if True == ifres11604 {
ifres11603 = True


} else {
ifres11603 = False


}

ifres11602 = ifres11603


} else {
ifres11602 = False


}

if True == ifres11602 {
tmp11569 := Call(__e, PrimFunc(symgensym), symTm)


let__10085 := tmp11569
_ = let__10085

tmp11570 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2063)
}
__typedArg0 := V2063
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11571 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V2064, Nil)
}
__typedArg0 := V2064
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11572 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11570, tmp11571)
}
__typedArg0 := tmp11570
__typedArg1 := tmp11571
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11573 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4lazyderef, tmp11572)
}
__typedArg0 := symshen_4lazyderef
__typedArg1 := tmp11572
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11574 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10085, Nil)
}
__typedArg0 := let__10085
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11575 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons_2, tmp11574)
}
__typedArg0 := symcons_2
__typedArg1 := tmp11574
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11576 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2062)
}
__typedArg0 := V2062
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11577 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11576)
}
__typedArg0 := tmp11576
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11578 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11577)
}
__typedArg0 := tmp11577
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11579 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2062)
}
__typedArg0 := V2062
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11580 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11579)
}
__typedArg0 := tmp11579
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11581 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11580)
}
__typedArg0 := tmp11580
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11582 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11581)
}
__typedArg0 := tmp11581
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11583 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2062)
}
__typedArg0 := V2062
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11584 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11582, tmp11583)
}
__typedArg0 := tmp11582
__typedArg1 := tmp11583
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11585 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11578, tmp11584)
}
__typedArg0 := tmp11578
__typedArg1 := tmp11584
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11586 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10085, Nil)
}
__typedArg0 := let__10085
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11587 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhd, tmp11586)
}
__typedArg0 := symhd
__typedArg1 := tmp11586
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11588 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10085, Nil)
}
__typedArg0 := let__10085
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11589 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtl, tmp11588)
}
__typedArg0 := symtl
__typedArg1 := tmp11588
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11590 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2063)
}
__typedArg0 := V2063
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11591 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11589, tmp11590)
}
__typedArg0 := tmp11589
__typedArg1 := tmp11590
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11592 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11587, tmp11591)
}
__typedArg0 := tmp11587
__typedArg1 := tmp11591
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11593 := Call(__e, PrimFunc(symshen_4compile_1head), symshen_4_1m, tmp11585, tmp11592, V2064, V2065)


tmp11594 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(False, Nil)
}
__typedArg0 := False
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11595 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11593, tmp11594)
}
__typedArg0 := tmp11593
__typedArg1 := tmp11594
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11596 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11575, tmp11595)
}
__typedArg0 := tmp11575
__typedArg1 := tmp11595
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11597 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp11596)
}
__typedArg0 := symif
__typedArg1 := tmp11596
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11598 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11597, Nil)
}
__typedArg0 := tmp11597
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11599 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11573, tmp11598)
}
__typedArg0 := tmp11573
__typedArg1 := tmp11598
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11600 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10085, tmp11599)
}
__typedArg0 := let__10085
__typedArg1 := tmp11599
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp11600)
}
__typedArg0 := symlet
__typedArg1 := tmp11600
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

tmp11633 := Call(__e, ns2_1set, symshen_4cons_1case_1minus, tmp11568)


_ = tmp11633

tmp11634 := MakeNative(func(__e *ControlFlow) {
V2075 := __e.Get(1)
_ = V2075
V2076 := __e.Get(2)
_ = V2076
V2077 := __e.Get(3)
_ = V2077
V2078 := __e.Get(4)
_ = V2078
tmp11678 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2075)
}
__typedArg0 := V2075
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11675 Obj

if True == tmp11678 {
tmp11677 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2076)
}
__typedArg0 := V2076
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11676 Obj

if True == tmp11677 {
ifres11676 = True


} else {
ifres11676 = False


}

ifres11675 = ifres11676


} else {
ifres11675 = False


}

if True == ifres11675 {
tmp11635 := Call(__e, PrimFunc(symgensym), symTm)


let__10086 := tmp11635
_ = let__10086

tmp11636 := Call(__e, PrimFunc(symgensym), symGoTo)


let__10087 := tmp11636
_ = let__10087

tmp11637 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2076)
}
__typedArg0 := V2076
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11638 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V2077, Nil)
}
__typedArg0 := V2077
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11639 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11637, tmp11638)
}
__typedArg0 := tmp11637
__typedArg1 := tmp11638
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11640 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4lazyderef, tmp11639)
}
__typedArg0 := symshen_4lazyderef
__typedArg1 := tmp11639
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11641 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2075)
}
__typedArg0 := V2075
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11642 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2076)
}
__typedArg0 := V2076
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11643 := Call(__e, PrimFunc(symshen_4compile_1head), symshen_4_7m, tmp11641, tmp11642, V2077, V2078)


tmp11644 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11643, Nil)
}
__typedArg0 := tmp11643
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11645 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfreeze, tmp11644)
}
__typedArg0 := symfreeze
__typedArg1 := tmp11644
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11646 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2075)
}
__typedArg0 := V2075
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11647 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11646, Nil)
}
__typedArg0 := tmp11646
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11648 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10086, tmp11647)
}
__typedArg0 := let__10086
__typedArg1 := tmp11647
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11649 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_a, tmp11648)
}
__typedArg0 := sym_a
__typedArg1 := tmp11648
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11650 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10087, Nil)
}
__typedArg0 := let__10087
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11651 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symthaw, tmp11650)
}
__typedArg0 := symthaw
__typedArg1 := tmp11650
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11652 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10086, Nil)
}
__typedArg0 := let__10086
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11653 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4pvar_2, tmp11652)
}
__typedArg0 := symshen_4pvar_2
__typedArg1 := tmp11652
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11654 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2075)
}
__typedArg0 := V2075
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11655 := Call(__e, PrimFunc(symshen_4demode), tmp11654)


tmp11656 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10087, Nil)
}
__typedArg0 := let__10087
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11657 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V2077, tmp11656)
}
__typedArg0 := V2077
__typedArg1 := tmp11656
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11658 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11655, tmp11657)
}
__typedArg0 := tmp11655
__typedArg1 := tmp11657
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11659 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10086, tmp11658)
}
__typedArg0 := let__10086
__typedArg1 := tmp11658
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11660 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4bind_b, tmp11659)
}
__typedArg0 := symshen_4bind_b
__typedArg1 := tmp11659
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11661 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(False, Nil)
}
__typedArg0 := False
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11662 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11660, tmp11661)
}
__typedArg0 := tmp11660
__typedArg1 := tmp11661
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11663 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11653, tmp11662)
}
__typedArg0 := tmp11653
__typedArg1 := tmp11662
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11664 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp11663)
}
__typedArg0 := symif
__typedArg1 := tmp11663
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11665 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11664, Nil)
}
__typedArg0 := tmp11664
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11666 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11651, tmp11665)
}
__typedArg0 := tmp11651
__typedArg1 := tmp11665
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11667 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11649, tmp11666)
}
__typedArg0 := tmp11649
__typedArg1 := tmp11666
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11668 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp11667)
}
__typedArg0 := symif
__typedArg1 := tmp11667
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11669 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11668, Nil)
}
__typedArg0 := tmp11668
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11670 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11645, tmp11669)
}
__typedArg0 := tmp11645
__typedArg1 := tmp11669
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11671 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10087, tmp11670)
}
__typedArg0 := let__10087
__typedArg1 := tmp11670
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11672 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11640, tmp11671)
}
__typedArg0 := tmp11640
__typedArg1 := tmp11671
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11673 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10086, tmp11672)
}
__typedArg0 := let__10086
__typedArg1 := tmp11672
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp11673)
}
__typedArg0 := symlet
__typedArg1 := tmp11673
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

tmp11679 := Call(__e, ns2_1set, symshen_4atom_1case_1plus, tmp11634)


_ = tmp11679

tmp11680 := MakeNative(func(__e *ControlFlow) {
V2089 := __e.Get(1)
_ = V2089
V2090 := __e.Get(2)
_ = V2090
V2091 := __e.Get(3)
_ = V2091
V2092 := __e.Get(4)
_ = V2092
tmp11771 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2089)
}
__typedArg0 := V2089
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11741 Obj

if True == tmp11771 {
tmp11769 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2089)
}
__typedArg0 := V2089
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11770 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11769)
}
__typedArg0 := tmp11769
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11743 Obj

if True == tmp11770 {
tmp11766 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2089)
}
__typedArg0 := V2089
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11767 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11766)
}
__typedArg0 := tmp11766
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11768 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcons, tmp11767)
}
__typedArg0 := symcons
__typedArg1 := tmp11767
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11745 Obj

if True == tmp11768 {
tmp11763 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2089)
}
__typedArg0 := V2089
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11764 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11763)
}
__typedArg0 := tmp11763
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11765 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11764)
}
__typedArg0 := tmp11764
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11747 Obj

if True == tmp11765 {
tmp11759 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2089)
}
__typedArg0 := V2089
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11760 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11759)
}
__typedArg0 := tmp11759
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11761 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11760)
}
__typedArg0 := tmp11760
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11762 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11761)
}
__typedArg0 := tmp11761
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11749 Obj

if True == tmp11762 {
tmp11754 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2089)
}
__typedArg0 := V2089
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11755 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11754)
}
__typedArg0 := tmp11754
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11756 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11755)
}
__typedArg0 := tmp11755
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11757 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11756)
}
__typedArg0 := tmp11756
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11758 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp11757)
}
__typedArg0 := Nil
__typedArg1 := tmp11757
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11751 Obj

if True == tmp11758 {
tmp11753 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2090)
}
__typedArg0 := V2090
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11752 Obj

if True == tmp11753 {
ifres11752 = True


} else {
ifres11752 = False


}

ifres11751 = ifres11752


} else {
ifres11751 = False


}

var ifres11750 Obj

if True == ifres11751 {
ifres11750 = True


} else {
ifres11750 = False


}

ifres11749 = ifres11750


} else {
ifres11749 = False


}

var ifres11748 Obj

if True == ifres11749 {
ifres11748 = True


} else {
ifres11748 = False


}

ifres11747 = ifres11748


} else {
ifres11747 = False


}

var ifres11746 Obj

if True == ifres11747 {
ifres11746 = True


} else {
ifres11746 = False


}

ifres11745 = ifres11746


} else {
ifres11745 = False


}

var ifres11744 Obj

if True == ifres11745 {
ifres11744 = True


} else {
ifres11744 = False


}

ifres11743 = ifres11744


} else {
ifres11743 = False


}

var ifres11742 Obj

if True == ifres11743 {
ifres11742 = True


} else {
ifres11742 = False


}

ifres11741 = ifres11742


} else {
ifres11741 = False


}

if True == ifres11741 {
tmp11681 := Call(__e, PrimFunc(symgensym), symTm)


let__10088 := tmp11681
_ = let__10088

tmp11682 := Call(__e, PrimFunc(symgensym), symGoTo)


let__10089 := tmp11682
_ = let__10089

tmp11683 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2089)
}
__typedArg0 := V2089
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11684 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11683)
}
__typedArg0 := tmp11683
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11685 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11684)
}
__typedArg0 := tmp11684
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11686 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2089)
}
__typedArg0 := V2089
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11687 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11686)
}
__typedArg0 := tmp11686
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11688 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11687)
}
__typedArg0 := tmp11687
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11689 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11688)
}
__typedArg0 := tmp11688
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11690 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11685, tmp11689)
}
__typedArg0 := tmp11685
__typedArg1 := tmp11689
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11691 := Call(__e, PrimFunc(symshen_4extract_1vars), tmp11690)


let__10090 := tmp11691
_ = let__10090

tmp11692 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2089)
}
__typedArg0 := V2089
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11693 := Call(__e, PrimFunc(symshen_4tame), tmp11692)


let__10091 := tmp11693
_ = let__10091

tmp11694 := Call(__e, PrimFunc(symshen_4extract_1vars), let__10091)


let__10092 := tmp11694
_ = let__10092

tmp11695 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2090)
}
__typedArg0 := V2090
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11696 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V2091, Nil)
}
__typedArg0 := V2091
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11697 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11695, tmp11696)
}
__typedArg0 := tmp11695
__typedArg1 := tmp11696
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11698 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4lazyderef, tmp11697)
}
__typedArg0 := symshen_4lazyderef
__typedArg1 := tmp11697
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11699 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2089)
}
__typedArg0 := V2089
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11700 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2090)
}
__typedArg0 := V2090
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11701 := Call(__e, PrimFunc(symshen_4compile_1head), symshen_4_7m, tmp11699, tmp11700, V2091, V2092)


tmp11702 := Call(__e, PrimFunc(symshen_4goto), let__10090, tmp11701)


tmp11703 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10088, Nil)
}
__typedArg0 := let__10088
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11704 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons_2, tmp11703)
}
__typedArg0 := symcons_2
__typedArg1 := tmp11703
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11705 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2089)
}
__typedArg0 := V2089
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11706 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11705)
}
__typedArg0 := tmp11705
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11707 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10088, Nil)
}
__typedArg0 := let__10088
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11708 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhd, tmp11707)
}
__typedArg0 := symhd
__typedArg1 := tmp11707
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11709 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10088, Nil)
}
__typedArg0 := let__10088
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11710 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtl, tmp11709)
}
__typedArg0 := symtl
__typedArg1 := tmp11709
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11711 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11710, Nil)
}
__typedArg0 := tmp11710
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11712 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11708, tmp11711)
}
__typedArg0 := tmp11708
__typedArg1 := tmp11711
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11713 := Call(__e, PrimFunc(symshen_4invoke), let__10089, let__10090)


tmp11714 := Call(__e, PrimFunc(symshen_4compile_1head), symshen_4_7m, tmp11706, tmp11712, V2091, tmp11713)


tmp11715 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10088, Nil)
}
__typedArg0 := let__10088
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11716 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4pvar_2, tmp11715)
}
__typedArg0 := symshen_4pvar_2
__typedArg1 := tmp11715
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11717 := Call(__e, PrimFunc(symshen_4demode), let__10091)


tmp11718 := Call(__e, PrimFunc(symshen_4invoke), let__10089, let__10090)


tmp11719 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11718, Nil)
}
__typedArg0 := tmp11718
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11720 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfreeze, tmp11719)
}
__typedArg0 := symfreeze
__typedArg1 := tmp11719
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11721 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11720, Nil)
}
__typedArg0 := tmp11720
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11722 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V2091, tmp11721)
}
__typedArg0 := V2091
__typedArg1 := tmp11721
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11723 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11717, tmp11722)
}
__typedArg0 := tmp11717
__typedArg1 := tmp11722
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11724 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10088, tmp11723)
}
__typedArg0 := let__10088
__typedArg1 := tmp11723
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11725 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4bind_b, tmp11724)
}
__typedArg0 := symshen_4bind_b
__typedArg1 := tmp11724
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11726 := Call(__e, PrimFunc(symshen_4stpart), let__10092, tmp11725, V2091)


tmp11727 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(False, Nil)
}
__typedArg0 := False
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11728 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11726, tmp11727)
}
__typedArg0 := tmp11726
__typedArg1 := tmp11727
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11729 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11716, tmp11728)
}
__typedArg0 := tmp11716
__typedArg1 := tmp11728
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11730 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp11729)
}
__typedArg0 := symif
__typedArg1 := tmp11729
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11731 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11730, Nil)
}
__typedArg0 := tmp11730
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11732 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11714, tmp11731)
}
__typedArg0 := tmp11714
__typedArg1 := tmp11731
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11733 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11704, tmp11732)
}
__typedArg0 := tmp11704
__typedArg1 := tmp11732
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11734 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp11733)
}
__typedArg0 := symif
__typedArg1 := tmp11733
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11735 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11734, Nil)
}
__typedArg0 := tmp11734
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11736 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11702, tmp11735)
}
__typedArg0 := tmp11702
__typedArg1 := tmp11735
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11737 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10089, tmp11736)
}
__typedArg0 := let__10089
__typedArg1 := tmp11736
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11738 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11698, tmp11737)
}
__typedArg0 := tmp11698
__typedArg1 := tmp11737
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11739 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10088, tmp11738)
}
__typedArg0 := let__10088
__typedArg1 := tmp11738
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp11739)
}
__typedArg0 := symlet
__typedArg1 := tmp11739
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

tmp11772 := Call(__e, ns2_1set, symshen_4cons_1case_1plus, tmp11680)


_ = tmp11772

tmp11773 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__selftop:
V2098 := __self1
_ = V2098
tmp11810 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2098)
}
__typedArg0 := V2098
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11797 Obj

if True == tmp11810 {
tmp11808 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2098)
}
__typedArg0 := V2098
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11809 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4_7m, tmp11808)
}
__typedArg0 := symshen_4_7m
__typedArg1 := tmp11808
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11799 Obj

if True == tmp11809 {
tmp11806 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2098)
}
__typedArg0 := V2098
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11807 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11806)
}
__typedArg0 := tmp11806
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11801 Obj

if True == tmp11807 {
tmp11803 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2098)
}
__typedArg0 := V2098
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11804 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11803)
}
__typedArg0 := tmp11803
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11805 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp11804)
}
__typedArg0 := Nil
__typedArg1 := tmp11804
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11802 Obj

if True == tmp11805 {
ifres11802 = True


} else {
ifres11802 = False


}

ifres11801 = ifres11802


} else {
ifres11801 = False


}

var ifres11800 Obj

if True == ifres11801 {
ifres11800 = True


} else {
ifres11800 = False


}

ifres11799 = ifres11800


} else {
ifres11799 = False


}

var ifres11798 Obj

if True == ifres11799 {
ifres11798 = True


} else {
ifres11798 = False


}

ifres11797 = ifres11798


} else {
ifres11797 = False


}

if True == ifres11797 {
tmp11774 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2098)
}
__typedArg0 := V2098
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11775 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11774)
}
__typedArg0 := tmp11774
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

if PrimFunc(symshen_4demode) == __self {
__self1 = tmp11775
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4demode), tmp11775)
return


} else {
tmp11795 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2098)
}
__typedArg0 := V2098
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11782 Obj

if True == tmp11795 {
tmp11793 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2098)
}
__typedArg0 := V2098
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11794 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4_1m, tmp11793)
}
__typedArg0 := symshen_4_1m
__typedArg1 := tmp11793
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11784 Obj

if True == tmp11794 {
tmp11791 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2098)
}
__typedArg0 := V2098
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11792 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11791)
}
__typedArg0 := tmp11791
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11786 Obj

if True == tmp11792 {
tmp11788 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2098)
}
__typedArg0 := V2098
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11789 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11788)
}
__typedArg0 := tmp11788
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11790 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp11789)
}
__typedArg0 := Nil
__typedArg1 := tmp11789
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11787 Obj

if True == tmp11790 {
ifres11787 = True


} else {
ifres11787 = False


}

ifres11786 = ifres11787


} else {
ifres11786 = False


}

var ifres11785 Obj

if True == ifres11786 {
ifres11785 = True


} else {
ifres11785 = False


}

ifres11784 = ifres11785


} else {
ifres11784 = False


}

var ifres11783 Obj

if True == ifres11784 {
ifres11783 = True


} else {
ifres11783 = False


}

ifres11782 = ifres11783


} else {
ifres11782 = False


}

if True == ifres11782 {
tmp11776 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2098)
}
__typedArg0 := V2098
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11777 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11776)
}
__typedArg0 := tmp11776
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

if PrimFunc(symshen_4demode) == __self {
__self1 = tmp11777
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4demode), tmp11777)
return


} else {
tmp11780 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2098)
}
__typedArg0 := V2098
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp11780 {
tmp11778 := MakeNative(func(__e *ControlFlow) {
Z2099 := __e.Get(1)
_ = Z2099
__e.TailApply(PrimFunc(symshen_4demode), Z2099)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp11778, V2098)
return


} else {
__e.Return(V2098)
return
}


}


}


}, 1)

tmp11811 := Call(__e, ns2_1set, symshen_4demode, tmp11773)


_ = tmp11811

tmp11812 := MakeNative(func(__e *ControlFlow) {
V2100 := __e.Get(1)
_ = V2100
tmp11817 := Call(__e, PrimFunc(symshen_4wildcard_2), V2100)


if True == tmp11817 {
__e.TailApply(PrimFunc(symgensym), symY)
return
} else {
tmp11815 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2100)
}
__typedArg0 := V2100
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp11815 {
tmp11813 := MakeNative(func(__e *ControlFlow) {
Z2101 := __e.Get(1)
_ = Z2101
__e.TailApply(PrimFunc(symshen_4tame), Z2101)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp11813, V2100)
return


} else {
__e.Return(V2100)
return
}


}


}, 1)

tmp11818 := Call(__e, ns2_1set, symshen_4tame, tmp11812)


_ = tmp11818

tmp11819 := MakeNative(func(__e *ControlFlow) {
V2102 := __e.Get(1)
_ = V2102
V2103 := __e.Get(2)
_ = V2103
tmp11822 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V2102)
}
__typedArg0 := Nil
__typedArg1 := V2102
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp11822 {
tmp11820 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V2103, Nil)
}
__typedArg0 := V2103
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfreeze, tmp11820)
}
__typedArg0 := symfreeze
__typedArg1 := tmp11820
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.TailApply(PrimFunc(symshen_4goto_1h), V2102, V2103)
return
}


}, 2)

tmp11823 := Call(__e, ns2_1set, symshen_4goto, tmp11819)


_ = tmp11823

tmp11824 := MakeNative(func(__e *ControlFlow) {
V2104 := __e.Get(1)
_ = V2104
V2105 := __e.Get(2)
_ = V2105
tmp11833 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V2104)
}
__typedArg0 := Nil
__typedArg1 := V2104
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp11833 {
__e.Return(V2105)
return
} else {
tmp11831 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2104)
}
__typedArg0 := V2104
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp11831 {
tmp11825 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2104)
}
__typedArg0 := V2104
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11826 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2104)
}
__typedArg0 := V2104
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11827 := Call(__e, PrimFunc(symshen_4goto_1h), tmp11826, V2105)


tmp11828 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11827, Nil)
}
__typedArg0 := tmp11827
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11829 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11825, tmp11828)
}
__typedArg0 := tmp11825
__typedArg1 := tmp11828
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp11829)
}
__typedArg0 := symlambda
__typedArg1 := tmp11829
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4goto_1h)
return
}


}


}, 2)

tmp11834 := Call(__e, ns2_1set, symshen_4goto_1h, tmp11824)


_ = tmp11834

tmp11835 := MakeNative(func(__e *ControlFlow) {
V2106 := __e.Get(1)
_ = V2106
V2107 := __e.Get(2)
_ = V2107
tmp11838 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V2107)
}
__typedArg0 := Nil
__typedArg1 := V2107
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp11838 {
tmp11836 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V2106, Nil)
}
__typedArg0 := V2106
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symthaw, tmp11836)
}
__typedArg0 := symthaw
__typedArg1 := tmp11836
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

tmp11839 := Call(__e, ns2_1set, symshen_4invoke, tmp11835)


_ = tmp11839

tmp11840 := MakeNative(func(__e *ControlFlow) {
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

tmp11841 := Call(__e, ns2_1set, symshen_4wildcard_2, tmp11840)


_ = tmp11841

tmp11842 := MakeNative(func(__e *ControlFlow) {
V2109 := __e.Get(1)
_ = V2109
tmp11849 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symabsvector_2) {
return PrimIsVector(V2109)
}
__typedArg0 := V2109
return Call(__e, PrimFunc(symabsvector_2), __typedArg0)
})()

if True == tmp11849 {
tmp11844 := MakeNative(func(__e *ControlFlow) {
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

tmp11845 := MakeNative(func(__e *ControlFlow) {
Z2110 := __e.Get(1)
_ = Z2110
__e.Return(symshen_4not_1pvar)
return
}, 1)

tmp11846 := Call(__e, try_1catch, tmp11844, tmp11845)


tmp11847 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp11846, symshen_4pvar)
}
__typedArg0 := tmp11846
__typedArg1 := symshen_4pvar
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp11847 {
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

tmp11850 := Call(__e, ns2_1set, symshen_4pvar_2, tmp11842)


_ = tmp11850

tmp11851 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__selftop:
V2111 := __self1
_ = V2111
V2112 := __self2
_ = V2112
tmp11857 := Call(__e, PrimFunc(symshen_4pvar_2), V2111)


if True == tmp11857 {
tmp11852 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V2111, MakeInteger(1))
}
__typedArg0 := V2111
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})()

tmp11853 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V2112, tmp11852)
}
__typedArg0 := V2112
__typedArg1 := tmp11852
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})()

let__10093 := tmp11853
_ = let__10093

tmp11855 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__10093, symshen_4_1null_1)
}
__typedArg0 := let__10093
__typedArg1 := symshen_4_1null_1
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp11855 {
__e.Return(V2111)
return
} else {
if PrimFunc(symshen_4lazyderef) == __self {
__self1, __self2 = let__10093, V2112
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4lazyderef), let__10093, V2112)
return
}


} else {
__e.Return(V2111)
return
}


}, 2)

tmp11858 := Call(__e, ns2_1set, symshen_4lazyderef, tmp11851)


_ = tmp11858

tmp11859 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__selftop:
V2114 := __self1
_ = V2114
V2115 := __self2
_ = V2115
tmp11871 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2114)
}
__typedArg0 := V2114
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp11871 {
tmp11860 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2114)
}
__typedArg0 := V2114
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11861 := Call(__e, PrimFunc(symshen_4deref), tmp11860, V2115)


tmp11862 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2114)
}
__typedArg0 := V2114
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11863 := Call(__e, PrimFunc(symshen_4deref), tmp11862, V2115)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11861, tmp11863)
}
__typedArg0 := tmp11861
__typedArg1 := tmp11863
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp11869 := Call(__e, PrimFunc(symshen_4pvar_2), V2114)


if True == tmp11869 {
tmp11864 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V2114, MakeInteger(1))
}
__typedArg0 := V2114
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})()

tmp11865 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V2115, tmp11864)
}
__typedArg0 := V2115
__typedArg1 := tmp11864
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})()

let__10094 := tmp11865
_ = let__10094

tmp11867 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__10094, symshen_4_1null_1)
}
__typedArg0 := let__10094
__typedArg1 := symshen_4_1null_1
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp11867 {
__e.Return(V2114)
return
} else {
if PrimFunc(symshen_4deref) == __self {
__self1, __self2 = let__10094, V2115
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4deref), let__10094, V2115)
return
}


} else {
__e.Return(V2114)
return
}


}


}, 2)

tmp11872 := Call(__e, ns2_1set, symshen_4deref, tmp11859)


_ = tmp11872

tmp11873 := MakeNative(func(__e *ControlFlow) {
V2117 := __e.Get(1)
_ = V2117
V2118 := __e.Get(2)
_ = V2118
V2119 := __e.Get(3)
_ = V2119
V2120 := __e.Get(4)
_ = V2120
tmp11874 := Call(__e, PrimFunc(symshen_4bindv), V2117, V2118, V2119)


let__10095 := tmp11874
_ = let__10095

tmp11875 := Call(__e, PrimFunc(symthaw), V2120)


let__10096 := tmp11875
_ = let__10096

tmp11877 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__10096, False)
}
__typedArg0 := let__10096
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp11877 {
__e.TailApply(PrimFunc(symshen_4unwind), V2117, V2119, let__10096)
return
} else {
__e.Return(let__10096)
return
}


}, 4)

tmp11878 := Call(__e, ns2_1set, symshen_4bind_b, tmp11873)


_ = tmp11878

tmp11879 := MakeNative(func(__e *ControlFlow) {
V2123 := __e.Get(1)
_ = V2123
V2124 := __e.Get(2)
_ = V2124
V2125 := __e.Get(3)
_ = V2125
tmp11880 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V2123, MakeInteger(1))
}
__typedArg0 := V2123
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(V2125, tmp11880, V2124)
}
__typedArg0 := V2125
__typedArg1 := tmp11880
__typedArg2 := V2124
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})())
return


}, 3)

tmp11881 := Call(__e, ns2_1set, symshen_4bindv, tmp11879)


_ = tmp11881

tmp11882 := MakeNative(func(__e *ControlFlow) {
V2126 := __e.Get(1)
_ = V2126
V2127 := __e.Get(2)
_ = V2127
V2128 := __e.Get(3)
_ = V2128
tmp11883 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V2126, MakeInteger(1))
}
__typedArg0 := V2126
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})()

tmp11884 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(V2127, tmp11883, symshen_4_1null_1)
}
__typedArg0 := V2127
__typedArg1 := tmp11883
__typedArg2 := symshen_4_1null_1
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

_ = tmp11884

__e.Return(V2128)
return


}, 3)

tmp11885 := Call(__e, ns2_1set, symshen_4unwind, tmp11882)


_ = tmp11885

tmp11886 := MakeNative(func(__e *ControlFlow) {
V2137 := __e.Get(1)
_ = V2137
V2138 := __e.Get(2)
_ = V2138
V2139 := __e.Get(3)
_ = V2139
tmp11901 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V2137)
}
__typedArg0 := Nil
__typedArg1 := V2137
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp11901 {
__e.Return(V2138)
return
} else {
tmp11899 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2137)
}
__typedArg0 := V2137
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp11899 {
tmp11887 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2137)
}
__typedArg0 := V2137
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11888 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V2139, Nil)
}
__typedArg0 := V2139
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11889 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4newpv, tmp11888)
}
__typedArg0 := symshen_4newpv
__typedArg1 := tmp11888
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11890 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2137)
}
__typedArg0 := V2137
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11891 := Call(__e, PrimFunc(symshen_4stpart), tmp11890, V2138, V2139)


tmp11892 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11891, Nil)
}
__typedArg0 := tmp11891
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11893 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V2139, tmp11892)
}
__typedArg0 := V2139
__typedArg1 := tmp11892
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11894 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4gc, tmp11893)
}
__typedArg0 := symshen_4gc
__typedArg1 := tmp11893
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11895 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11894, Nil)
}
__typedArg0 := tmp11894
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11896 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11889, tmp11895)
}
__typedArg0 := tmp11889
__typedArg1 := tmp11895
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11897 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11887, tmp11896)
}
__typedArg0 := tmp11887
__typedArg1 := tmp11896
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp11897)
}
__typedArg0 := symlet
__typedArg1 := tmp11897
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

tmp11902 := Call(__e, ns2_1set, symshen_4stpart, tmp11886)


_ = tmp11902

tmp11903 := MakeNative(func(__e *ControlFlow) {
V2140 := __e.Get(1)
_ = V2140
V2141 := __e.Get(2)
_ = V2141
tmp11907 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V2141, False)
}
__typedArg0 := V2141
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp11907 {
tmp11904 := Call(__e, PrimFunc(symshen_4ticket_1number), V2140)


let__10097 := tmp11904
_ = let__10097

tmp11905 := Call(__e, PrimFunc(symshen_4decrement_1ticket), let__10097, V2140)


_ = tmp11905

__e.Return(V2141)
return


} else {
__e.Return(V2141)
return
}


}, 2)

tmp11908 := Call(__e, ns2_1set, symshen_4gc, tmp11903)


_ = tmp11908

tmp11909 := MakeNative(func(__e *ControlFlow) {
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

tmp11911 := Call(__e, ns2_1set, symshen_4decrement_1ticket, tmp11909)


_ = tmp11911

tmp11912 := MakeNative(func(__e *ControlFlow) {
V2145 := __e.Get(1)
_ = V2145
tmp11913 := Call(__e, PrimFunc(symshen_4ticket_1number), V2145)


let__10098 := tmp11913
_ = let__10098

tmp11914 := Call(__e, PrimFunc(symshen_4make_1prolog_1variable), let__10098)


let__10099 := tmp11914
_ = let__10099

tmp11915 := Call(__e, PrimFunc(symshen_4nextticket), V2145, let__10098)


let__10100 := tmp11915
_ = let__10100

__e.Return(let__10099)
return


}, 1)

tmp11916 := Call(__e, ns2_1set, symshen_4newpv, tmp11912)


_ = tmp11916

tmp11917 := MakeNative(func(__e *ControlFlow) {
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

tmp11918 := Call(__e, ns2_1set, symshen_4ticket_1number, tmp11917)


_ = tmp11918

tmp11919 := MakeNative(func(__e *ControlFlow) {
V2150 := __e.Get(1)
_ = V2150
V2151 := __e.Get(2)
_ = V2151
tmp11920 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(V2150, V2151, symshen_4_1null_1)
}
__typedArg0 := V2150
__typedArg1 := V2151
__typedArg2 := symshen_4_1null_1
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

let__10101 := tmp11920
_ = let__10101

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(let__10101, MakeInteger(1), (func() Obj {
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
__typedArg0 := let__10101
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

tmp11922 := Call(__e, ns2_1set, symshen_4nextticket, tmp11919)


_ = tmp11922

tmp11923 := MakeNative(func(__e *ControlFlow) {
V2153 := __e.Get(1)
_ = V2153
tmp11924 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symabsvector) {
return PrimAbsvector(MakeInteger(2))
}
__typedArg0 := MakeInteger(2)
return Call(__e, PrimFunc(symabsvector), __typedArg0)
})()

tmp11925 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(tmp11924, MakeInteger(0), symshen_4pvar)
}
__typedArg0 := tmp11924
__typedArg1 := MakeInteger(0)
__typedArg2 := symshen_4pvar
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(tmp11925, MakeInteger(1), V2153)
}
__typedArg0 := tmp11925
__typedArg1 := MakeInteger(1)
__typedArg2 := V2153
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})())
return


}, 1)

tmp11926 := Call(__e, ns2_1set, symshen_4make_1prolog_1variable, tmp11923)


_ = tmp11926

tmp11927 := MakeNative(func(__e *ControlFlow) {
V2154 := __e.Get(1)
_ = V2154
tmp11928 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V2154, MakeInteger(1))
}
__typedArg0 := V2154
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})()

tmp11929 := Call(__e, PrimFunc(symshen_4app), tmp11928, MakeString(""), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("Var"))
__typedS1, __typedOK1 := TypedString(tmp11929)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("Var")
__typedArg1 := tmp11929
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp11930 := Call(__e, ns2_1set, symshen_4pvar, tmp11927)


_ = tmp11930

tmp11931 := MakeNative(func(__e *ControlFlow) {
tmp11932 := (func() Obj {
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
__typedN1, __typedOK1 := TypedFloat64(tmp11932)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp11932
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())
}
__typedArg0 := symshen_4_dinfs_d
__typedArg1 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(1))
__typedN1, __typedOK1 := TypedFloat64(tmp11932)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp11932
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return


}, 0)

tmp11934 := Call(__e, ns2_1set, symshen_4incinfs, tmp11931)


_ = tmp11934

tmp11935 := MakeNative(func(__e *ControlFlow) {
V2155 := __e.Get(1)
_ = V2155
tmp11942 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(syminteger_2) {
return PrimIsInteger(V2155)
}
__typedArg0 := V2155
return Call(__e, PrimFunc(syminteger_2), __typedArg0)
})()

var ifres11939 Obj

if True == tmp11942 {
tmp11941 := (func() Obj {
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

var ifres11940 Obj

if True == tmp11941 {
ifres11940 = True


} else {
ifres11940 = False


}

ifres11939 = ifres11940


} else {
ifres11939 = False


}

if True == ifres11939 {
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
tmp11936 := Call(__e, PrimFunc(symshen_4app), V2155, MakeString(""), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("prolog vector size: size should be a positive integer; not "))
__typedS1, __typedOK1 := TypedString(tmp11936)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("prolog vector size: size should be a positive integer; not ")
__typedArg1 := tmp11936
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("prolog vector size: size should be a positive integer; not "))
__typedS1, __typedOK1 := TypedString(tmp11936)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("prolog vector size: size should be a positive integer; not ")
__typedArg1 := tmp11936
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}


}, 1)

tmp11943 := Call(__e, ns2_1set, symshen_4prolog_1vector_1size, tmp11935)


_ = tmp11943

tmp11944 := MakeNative(func(__e *ControlFlow) {
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
tmp11974 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V2167, V2168)
}
__typedArg0 := V2167
__typedArg1 := V2168
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp11974 {
__e.TailApply(PrimFunc(symthaw), V2170)
return
} else {
tmp11972 := Call(__e, PrimFunc(symshen_4pvar_2), V2167)


var ifres11967 Obj

if True == tmp11972 {
tmp11969 := Call(__e, PrimFunc(symshen_4deref), V2168, V2169)


tmp11970 := Call(__e, PrimFunc(symshen_4occurs_1check_2), V2167, tmp11969)


tmp11971 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp11970)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp11970
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres11968 Obj

if True == tmp11971 {
ifres11968 = True


} else {
ifres11968 = False


}

ifres11967 = ifres11968


} else {
ifres11967 = False


}

if True == ifres11967 {
__e.TailApply(PrimFunc(symshen_4bind_b), V2167, V2168, V2169, V2170)
return
} else {
tmp11965 := Call(__e, PrimFunc(symshen_4pvar_2), V2168)


var ifres11960 Obj

if True == tmp11965 {
tmp11962 := Call(__e, PrimFunc(symshen_4deref), V2167, V2169)


tmp11963 := Call(__e, PrimFunc(symshen_4occurs_1check_2), V2168, tmp11962)


tmp11964 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp11963)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp11963
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres11961 Obj

if True == tmp11964 {
ifres11961 = True


} else {
ifres11961 = False


}

ifres11960 = ifres11961


} else {
ifres11960 = False


}

if True == ifres11960 {
__e.TailApply(PrimFunc(symshen_4bind_b), V2168, V2167, V2169, V2170)
return
} else {
tmp11958 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2167)
}
__typedArg0 := V2167
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11955 Obj

if True == tmp11958 {
tmp11957 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2168)
}
__typedArg0 := V2168
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11956 Obj

if True == tmp11957 {
ifres11956 = True


} else {
ifres11956 = False


}

ifres11955 = ifres11956


} else {
ifres11955 = False


}

if True == ifres11955 {
tmp11945 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2167)
}
__typedArg0 := V2167
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11946 := Call(__e, PrimFunc(symshen_4lazyderef), tmp11945, V2169)


tmp11947 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2168)
}
__typedArg0 := V2168
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11948 := Call(__e, PrimFunc(symshen_4lazyderef), tmp11947, V2169)


tmp11949 := MakeNative(func(__e *ControlFlow) {
tmp11950 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2167)
}
__typedArg0 := V2167
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11951 := Call(__e, PrimFunc(symshen_4lazyderef), tmp11950, V2169)


tmp11952 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2168)
}
__typedArg0 := V2168
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11953 := Call(__e, PrimFunc(symshen_4lazyderef), tmp11952, V2169)


__e.TailApply(PrimFunc(symshen_4lzy_a_b), tmp11951, tmp11953, V2169, V2170)
return


}, 0)

if PrimFunc(symshen_4lzy_a_b) == __self {
__self1, __self2, __self3, __self4 = tmp11946, tmp11948, V2169, tmp11949
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4lzy_a_b), tmp11946, tmp11948, V2169, tmp11949)
return


} else {
__e.Return(False)
return
}


}


}


}


}, 4)

tmp11975 := Call(__e, ns2_1set, symshen_4lzy_a_b, tmp11944)


_ = tmp11975

tmp11976 := MakeNative(func(__e *ControlFlow) {
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
tmp11996 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V2182, V2183)
}
__typedArg0 := V2182
__typedArg1 := V2183
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp11996 {
__e.TailApply(PrimFunc(symthaw), V2185)
return
} else {
tmp11994 := Call(__e, PrimFunc(symshen_4pvar_2), V2182)


if True == tmp11994 {
__e.TailApply(PrimFunc(symshen_4bind_b), V2182, V2183, V2184, V2185)
return
} else {
tmp11992 := Call(__e, PrimFunc(symshen_4pvar_2), V2183)


if True == tmp11992 {
__e.TailApply(PrimFunc(symshen_4bind_b), V2183, V2182, V2184, V2185)
return
} else {
tmp11990 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2182)
}
__typedArg0 := V2182
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11987 Obj

if True == tmp11990 {
tmp11989 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2183)
}
__typedArg0 := V2183
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11988 Obj

if True == tmp11989 {
ifres11988 = True


} else {
ifres11988 = False


}

ifres11987 = ifres11988


} else {
ifres11987 = False


}

if True == ifres11987 {
tmp11977 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2182)
}
__typedArg0 := V2182
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11978 := Call(__e, PrimFunc(symshen_4lazyderef), tmp11977, V2184)


tmp11979 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2183)
}
__typedArg0 := V2183
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11980 := Call(__e, PrimFunc(symshen_4lazyderef), tmp11979, V2184)


tmp11981 := MakeNative(func(__e *ControlFlow) {
tmp11982 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2182)
}
__typedArg0 := V2182
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11983 := Call(__e, PrimFunc(symshen_4lazyderef), tmp11982, V2184)


tmp11984 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2183)
}
__typedArg0 := V2183
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11985 := Call(__e, PrimFunc(symshen_4lazyderef), tmp11984, V2184)


__e.TailApply(PrimFunc(symshen_4lzy_a), tmp11983, tmp11985, V2184, V2185)
return


}, 0)

if PrimFunc(symshen_4lzy_a) == __self {
__self1, __self2, __self3, __self4 = tmp11978, tmp11980, V2184, tmp11981
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4lzy_a), tmp11978, tmp11980, V2184, tmp11981)
return


} else {
__e.Return(False)
return
}


}


}


}


}, 4)

tmp11997 := Call(__e, ns2_1set, symshen_4lzy_a, tmp11976)


_ = tmp11997

tmp11998 := MakeNative(func(__e *ControlFlow) {
V2191 := __e.Get(1)
_ = V2191
V2192 := __e.Get(2)
_ = V2192
tmp12008 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V2191, V2192)
}
__typedArg0 := V2191
__typedArg1 := V2192
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp12008 {
__e.Return(True)
return
} else {
tmp12006 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2192)
}
__typedArg0 := V2192
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp12006 {
tmp12003 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2192)
}
__typedArg0 := V2192
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12004 := Call(__e, PrimFunc(symshen_4occurs_1check_2), V2191, tmp12003)


if True == tmp12004 {
__e.Return(True)
return
} else {
tmp12000 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2192)
}
__typedArg0 := V2192
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12001 := Call(__e, PrimFunc(symshen_4occurs_1check_2), V2191, tmp12000)


if True == tmp12001 {
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

tmp12009 := Call(__e, ns2_1set, symshen_4occurs_1check_2, tmp11998)


_ = tmp12009

tmp12010 := MakeNative(func(__e *ControlFlow) {
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
tmp12011 := Call(__e, V2193, V2194)


tmp12012 := Call(__e, tmp12011, V2195)


tmp12013 := Call(__e, tmp12012, V2196)


__e.TailApply(tmp12013, V2197)
return


}, 5)

tmp12014 := Call(__e, ns2_1set, symcall, tmp12010)


_ = tmp12014

tmp12015 := MakeNative(func(__e *ControlFlow) {
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

tmp12016 := Call(__e, ns2_1set, symreturn, tmp12015)


_ = tmp12016

tmp12017 := MakeNative(func(__e *ControlFlow) {
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

tmp12019 := Call(__e, ns2_1set, symwhen, tmp12017)


_ = tmp12019

tmp12020 := MakeNative(func(__e *ControlFlow) {
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
tmp12021 := Call(__e, PrimFunc(symshen_4lazyderef), V2220, V2222)


tmp12022 := Call(__e, PrimFunc(symshen_4lazyderef), V2221, V2222)


__e.TailApply(PrimFunc(symshen_4lzy_a), tmp12021, tmp12022, V2222, V2225)
return


}, 6)

tmp12023 := Call(__e, ns2_1set, symis, tmp12020)


_ = tmp12023

tmp12024 := MakeNative(func(__e *ControlFlow) {
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
tmp12025 := Call(__e, PrimFunc(symshen_4lazyderef), V2226, V2228)


tmp12026 := Call(__e, PrimFunc(symshen_4lazyderef), V2227, V2228)


__e.TailApply(PrimFunc(symshen_4lzy_a_b), tmp12025, tmp12026, V2228, V2231)
return


}, 6)

tmp12027 := Call(__e, ns2_1set, symis_b, tmp12024)


_ = tmp12027

tmp12028 := MakeNative(func(__e *ControlFlow) {
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

tmp12029 := Call(__e, ns2_1set, symbind, tmp12028)


_ = tmp12029

tmp12030 := MakeNative(func(__e *ControlFlow) {
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
tmp12032 := Call(__e, PrimFunc(symshen_4lazyderef), V2242, V2243)


tmp12033 := Call(__e, PrimFunc(symshen_4pvar_2), tmp12032)


if True == tmp12033 {
__e.TailApply(PrimFunc(symthaw), V2246)
return
} else {
__e.Return(False)
return
}


}, 5)

tmp12034 := Call(__e, ns2_1set, symvar_2, tmp12030)


_ = tmp12034

tmp12035 := MakeNative(func(__e *ControlFlow) {
V2249 := __e.Get(1)
_ = V2249
__e.Return(MakeString("|prolog vector|"))
return
}, 1)

tmp12036 := Call(__e, ns2_1set, symshen_4print_1prolog_1vector, tmp12035)


_ = tmp12036

tmp12037 := MakeNative(func(__e *ControlFlow) {
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
tmp12049 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V2268)
}
__typedArg0 := Nil
__typedArg1 := V2268
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp12049 {
__e.Return(False)
return
} else {
tmp12047 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2268)
}
__typedArg0 := V2268
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp12047 {
tmp12038 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2268)
}
__typedArg0 := V2268
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12039 := Call(__e, tmp12038, V2269)


tmp12040 := Call(__e, tmp12039, V2270)


tmp12041 := Call(__e, tmp12040, V2271)


tmp12042 := Call(__e, tmp12041, V2272)


let__10102 := tmp12042
_ = let__10102

tmp12045 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__10102, False)
}
__typedArg0 := let__10102
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp12045 {
tmp12043 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2268)
}
__typedArg0 := V2268
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symfork) == __self {
__self1, __self2, __self3, __self4, __self5 = tmp12043, V2269, V2270, V2271, V2272
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symfork), tmp12043, V2269, V2270, V2271, V2272)
return


} else {
__e.Return(let__10102)
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

tmp12050 := Call(__e, ns2_1set, symfork, tmp12037)


_ = tmp12050

tmp12051 := MakeNative(func(__e *ControlFlow) {
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
tmp12057 := Call(__e, PrimFunc(symshen_4unlocked_2), V2278)


if True == tmp12057 {
tmp12052 := Call(__e, PrimFunc(symshen_4newpv), V2277)


let__10103 := tmp12052
_ = let__10103

tmp12053 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp12053

tmp12054 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4findall_1h), V2274, V2275, V2276, let__10103, V2277, V2278, V2279, V2280)
return
}, 0)

tmp12055 := Call(__e, PrimFunc(symis), let__10103, Nil, V2277, V2278, V2279, tmp12054)


__e.TailApply(PrimFunc(symshen_4gc), V2277, tmp12055)
return


} else {
__e.Return(False)
return
}


}, 7)

tmp12058 := Call(__e, ns2_1set, symfindall, tmp12051)


_ = tmp12058

tmp12059 := MakeNative(func(__e *ControlFlow) {
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
tmp12064 := Call(__e, PrimFunc(symshen_4unlocked_2), V2287)


var ifres12060 Obj

if True == tmp12064 {
tmp12061 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp12061

tmp12062 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4overbind), V2282, V2285, V2286, V2287, V2288, V2289)
return
}, 0)

tmp12063 := Call(__e, PrimFunc(symcall), V2283, V2286, V2287, V2288, tmp12062)


ifres12060 = tmp12063


} else {
ifres12060 = False


}

let__10104 := ifres12060
_ = let__10104

tmp12069 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__10104, False)
}
__typedArg0 := let__10104
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp12069 {
tmp12067 := Call(__e, PrimFunc(symshen_4unlocked_2), V2287)


if True == tmp12067 {
tmp12065 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp12065

__e.TailApply(PrimFunc(symis_b), V2284, V2285, V2286, V2287, V2288, V2289)
return


} else {
__e.Return(False)
return
}


} else {
__e.Return(let__10104)
return
}


}, 8)

tmp12070 := Call(__e, ns2_1set, symshen_4findall_1h, tmp12059)


_ = tmp12070

tmp12071 := MakeNative(func(__e *ControlFlow) {
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
tmp12072 := Call(__e, PrimFunc(symshen_4deref), V2297, V2299)


tmp12073 := Call(__e, PrimFunc(symshen_4lazyderef), V2298, V2299)


tmp12074 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12072, tmp12073)
}
__typedArg0 := tmp12072
__typedArg1 := tmp12073
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12075 := Call(__e, PrimFunc(symshen_4bindv), V2298, tmp12074, V2299)


_ = tmp12075

__e.Return(False)
return


}, 6)

tmp12076 := Call(__e, ns2_1set, symshen_4overbind, tmp12071)


_ = tmp12076

tmp12077 := MakeNative(func(__e *ControlFlow) {
V2305 := __e.Get(1)
_ = V2305
tmp12081 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_7, V2305)
}
__typedArg0 := sym_7
__typedArg1 := V2305
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp12081 {
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
tmp12079 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_1, V2305)
}
__typedArg0 := sym_1
__typedArg1 := V2305
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp12079 {
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

__e.TailApply(ns2_1set, symoccurs_1check, tmp12077)
return




}, 0)

