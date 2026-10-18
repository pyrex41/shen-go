package main

import . "github.com/pyrex41/shen-go/kl"

var PrologMain = MakeNative(func(__e *ControlFlow) {
tmp10112 := MakeNative(func(__e *ControlFlow) {
V1599 := __e.Get(1)
_ = V1599
__e.TailApply(PrimFunc(symshen_4assert_d), V1599, symshen_4top)
return
}, 1)

tmp10113 := Call(__e, ns2_1set, symasserta, tmp10112)


_ = tmp10113

tmp10114 := MakeNative(func(__e *ControlFlow) {
V1600 := __e.Get(1)
_ = V1600
__e.TailApply(PrimFunc(symshen_4assert_d), V1600, symshen_4bottom)
return
}, 1)

tmp10115 := Call(__e, ns2_1set, symassertz, tmp10114)


_ = tmp10115

tmp10116 := MakeNative(func(__e *ControlFlow) {
V1601 := __e.Get(1)
_ = V1601
V1602 := __e.Get(2)
_ = V1602
tmp10143 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1601)
}
__typedArg0 := V1601
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10134 Obj

if True == tmp10143 {
tmp10141 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1601)
}
__typedArg0 := V1601
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10142 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10141)
}
__typedArg0 := tmp10141
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10136 Obj

if True == tmp10142 {
tmp10138 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1601)
}
__typedArg0 := V1601
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10139 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10138)
}
__typedArg0 := tmp10138
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10140 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_5_1_1, tmp10139)
}
__typedArg0 := sym_5_1_1
__typedArg1 := tmp10139
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10137 Obj

if True == tmp10140 {
ifres10137 = True


} else {
ifres10137 = False


}

ifres10136 = ifres10137


} else {
ifres10136 = False


}

var ifres10135 Obj

if True == ifres10136 {
ifres10135 = True


} else {
ifres10135 = False


}

ifres10134 = ifres10135


} else {
ifres10134 = False


}

if True == ifres10134 {
tmp10117 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1601)
}
__typedArg0 := V1601
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10118 := Call(__e, PrimFunc(symshen_4predicate), tmp10117)


let__9894 := tmp10118
_ = let__9894

tmp10119 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1601)
}
__typedArg0 := V1601
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10120 := Call(__e, PrimFunc(symshen_4terms), tmp10119)


let__9895 := tmp10120
_ = let__9895

tmp10121 := Call(__e, PrimFunc(symlength), let__9895)


let__9896 := tmp10121
_ = let__9896

tmp10122 := Call(__e, PrimFunc(symshen_4parameters), let__9896)


let__9897 := tmp10122
_ = let__9897

tmp10123 := Call(__e, PrimFunc(symarity), let__9894)


let__9898 := tmp10123
_ = let__9898

tmp10129 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__9898, MakeInteger(-1))
}
__typedArg0 := let__9898
__typedArg1 := MakeInteger(-1)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10124 Obj

if True == tmp10129 {
tmp10125 := Call(__e, PrimFunc(symshen_4create_1skeleton), let__9894, let__9897)


tmp10126 := Call(__e, PrimFunc(symeval), tmp10125)


_ = tmp10126

tmp10127 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp10128 := Call(__e, PrimFunc(symput), let__9894, symshen_4dynamic, Nil, tmp10127)


ifres10124 = tmp10128


} else {
ifres10124 = symshen_4skip


}

let__9899 := ifres10124
_ = let__9899

tmp10130 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1601)
}
__typedArg0 := V1601
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10131 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10130)
}
__typedArg0 := tmp10130
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10132 := Call(__e, PrimFunc(symshen_4insert_1info), let__9894, let__9895, tmp10131, V1601, V1602)


let__9900 := tmp10132
_ = let__9900

__e.Return(let__9894)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4assert_d)
return
}


}, 2)

tmp10144 := Call(__e, ns2_1set, symshen_4assert_d, tmp10116)


_ = tmp10144

tmp10145 := MakeNative(func(__e *ControlFlow) {
V1612 := __e.Get(1)
_ = V1612
tmp10147 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1612)
}
__typedArg0 := V1612
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp10147 {
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

tmp10148 := Call(__e, ns2_1set, symshen_4predicate, tmp10145)


_ = tmp10148

tmp10149 := MakeNative(func(__e *ControlFlow) {
V1617 := __e.Get(1)
_ = V1617
tmp10151 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1617)
}
__typedArg0 := V1617
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp10151 {
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

tmp10152 := Call(__e, ns2_1set, symshen_4terms, tmp10149)


_ = tmp10152

tmp10153 := MakeNative(func(__e *ControlFlow) {
V1618 := __e.Get(1)
_ = V1618
V1619 := __e.Get(2)
_ = V1619
tmp10154 := Call(__e, PrimFunc(symshen_4dynamic_1default), V1618, V1619)


tmp10155 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1618, tmp10154)
}
__typedArg0 := V1618
__typedArg1 := tmp10154
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefprolog, tmp10155)
}
__typedArg0 := symdefprolog
__typedArg1 := tmp10155
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 2)

tmp10156 := Call(__e, ns2_1set, symshen_4create_1skeleton, tmp10153)


_ = tmp10156

tmp10157 := MakeNative(func(__e *ControlFlow) {
V1620 := __e.Get(1)
_ = V1620
V1621 := __e.Get(2)
_ = V1621
tmp10158 := Call(__e, PrimFunc(symshen_4cons_1form), V1621)


tmp10159 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4dynamic, Nil)
}
__typedArg0 := symshen_4dynamic
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10160 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1620, tmp10159)
}
__typedArg0 := V1620
__typedArg1 := tmp10159
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10161 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symget, tmp10160)
}
__typedArg0 := symget
__typedArg1 := tmp10160
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10162 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp10161, Nil)
}
__typedArg0 := tmp10161
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10163 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp10158, tmp10162)
}
__typedArg0 := tmp10158
__typedArg1 := tmp10162
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10164 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4call_1dynamic, tmp10163)
}
__typedArg0 := symshen_4call_1dynamic
__typedArg1 := tmp10163
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10165 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(";"))
}
__typedArg0 := MakeString(";")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp10166 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp10165, Nil)
}
__typedArg0 := tmp10165
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10167 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp10164, tmp10166)
}
__typedArg0 := tmp10164
__typedArg1 := tmp10166
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10168 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1_1, tmp10167)
}
__typedArg0 := sym_5_1_1
__typedArg1 := tmp10167
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symappend), V1621, tmp10168)
return


}, 2)

tmp10169 := Call(__e, ns2_1set, symshen_4dynamic_1default, tmp10157)


_ = tmp10169

tmp10170 := MakeNative(func(__e *ControlFlow) {
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
tmp10171 := Call(__e, PrimFunc(symgensym), symg)


let__9901 := tmp10171
_ = let__9901

tmp10172 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__9901, Nil)
}
__typedArg0 := let__9901
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10173 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefprolog, tmp10172)
}
__typedArg0 := symdefprolog
__typedArg1 := tmp10172
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10174 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1_1, V1624)
}
__typedArg0 := sym_5_1_1
__typedArg1 := V1624
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10175 := Call(__e, PrimFunc(symappend), V1623, tmp10174)


tmp10176 := Call(__e, PrimFunc(symappend), tmp10173, tmp10175)


tmp10177 := Call(__e, PrimFunc(symeval), tmp10176)


let__9902 := tmp10177
_ = let__9902

tmp10178 := Call(__e, PrimFunc(symfn), let__9901)


tmp10179 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__9901, V1625)
}
__typedArg0 := let__9901
__typedArg1 := V1625
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10180 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp10178, tmp10179)
}
__typedArg0 := tmp10178
__typedArg1 := tmp10179
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__9903 := tmp10180
_ = let__9903

tmp10181 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp10182 := Call(__e, PrimFunc(symget), V1622, symshen_4dynamic, tmp10181)


let__9904 := tmp10182
_ = let__9904

tmp10187 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V1626, symshen_4top)
}
__typedArg0 := V1626
__typedArg1 := symshen_4top
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10183 Obj

if True == tmp10187 {
tmp10184 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__9903, let__9904)
}
__typedArg0 := let__9903
__typedArg1 := let__9904
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

ifres10183 = tmp10184


} else {
tmp10185 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__9903, Nil)
}
__typedArg0 := let__9903
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10186 := Call(__e, PrimFunc(symappend), let__9904, tmp10185)


ifres10183 = tmp10186


}

let__9905 := ifres10183
_ = let__9905

tmp10188 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symput), V1622, symshen_4dynamic, let__9905, tmp10188)
return


}, 5)

tmp10189 := Call(__e, ns2_1set, symshen_4insert_1info, tmp10170)


_ = tmp10189

tmp10190 := MakeNative(func(__e *ControlFlow) {
tmp10191 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dnames_d)
}
__typedArg0 := symshen_4_dnames_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

let__9906 := tmp10191
_ = let__9906

tmp10197 := Call(__e, PrimFunc(symempty_2), let__9906)


var ifres10192 Obj

if True == tmp10197 {
tmp10193 := Call(__e, PrimFunc(symgensym), symg)


ifres10192 = tmp10193


} else {
tmp10194 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__9906)
}
__typedArg0 := let__9906
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10195 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dnames_d, tmp10194)
}
__typedArg0 := symshen_4_dnames_d
__typedArg1 := tmp10194
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp10195

tmp10196 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__9906)
}
__typedArg0 := let__9906
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

ifres10192 = tmp10196


}

let__9907 := ifres10192
_ = let__9907

__e.Return(let__9907)
return


}, 0)

tmp10198 := Call(__e, ns2_1set, symshen_4newname, tmp10190)


_ = tmp10198

tmp10199 := MakeNative(func(__e *ControlFlow) {
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
tmp10211 := Call(__e, PrimFunc(symshen_4unlocked_2), V1637)


var ifres10200 Obj

if True == tmp10211 {
tmp10201 := Call(__e, PrimFunc(symshen_4lazyderef), V1635, V1636)


let__9909 := tmp10201
_ = let__9909

tmp10210 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__9909)
}
__typedArg0 := let__9909
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10202 Obj

if True == tmp10210 {
tmp10203 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__9909)
}
__typedArg0 := let__9909
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10204 := Call(__e, PrimFunc(symshen_4lazyderef), tmp10203, V1636)


let__9910 := tmp10204
_ = let__9910

tmp10209 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__9910)
}
__typedArg0 := let__9910
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10205 Obj

if True == tmp10209 {
tmp10206 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__9910)
}
__typedArg0 := let__9910
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__9911 := tmp10206
_ = let__9911

tmp10207 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp10207

tmp10208 := Call(__e, PrimFunc(symshen_4callrec), let__9911, V1634, V1636, V1637, V1638, V1639)


ifres10205 = tmp10208


} else {
ifres10205 = False


}

ifres10202 = ifres10205


} else {
ifres10202 = False


}

ifres10200 = ifres10202


} else {
ifres10200 = False


}

let__9908 := ifres10200
_ = let__9908

tmp10220 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__9908, False)
}
__typedArg0 := let__9908
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp10220 {
tmp10218 := Call(__e, PrimFunc(symshen_4unlocked_2), V1637)


if True == tmp10218 {
tmp10212 := Call(__e, PrimFunc(symshen_4lazyderef), V1635, V1636)


let__9912 := tmp10212
_ = let__9912

tmp10216 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__9912)
}
__typedArg0 := let__9912
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp10216 {
tmp10213 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__9912)
}
__typedArg0 := let__9912
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__9913 := tmp10213
_ = let__9913

tmp10214 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp10214

if PrimFunc(symshen_4call_1dynamic) == __self {
__self1, __self2, __self3, __self4, __self5, __self6 = V1634, let__9913, V1636, V1637, V1638, V1639
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4call_1dynamic), V1634, let__9913, V1636, V1637, V1638, V1639)
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
__e.Return(let__9908)
return
}


}, 6)

tmp10221 := Call(__e, ns2_1set, symshen_4call_1dynamic, tmp10199)


_ = tmp10221

tmp10222 := MakeNative(func(__e *ControlFlow) {
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
tmp10232 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V1647)
}
__typedArg0 := Nil
__typedArg1 := V1647
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp10232 {
tmp10223 := Call(__e, V1646, V1648)


tmp10224 := Call(__e, tmp10223, V1649)


tmp10225 := Call(__e, tmp10224, V1650)


__e.TailApply(tmp10225, V1651)
return


} else {
tmp10230 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1647)
}
__typedArg0 := V1647
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp10230 {
tmp10226 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1647)
}
__typedArg0 := V1647
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10227 := Call(__e, V1646, tmp10226)


tmp10228 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1647)
}
__typedArg0 := V1647
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4callrec) == __self {
__self1, __self2, __self3, __self4, __self5, __self6 = tmp10227, tmp10228, V1648, V1649, V1650, V1651
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4callrec), tmp10227, tmp10228, V1648, V1649, V1650, V1651)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4callrec)
return
}


}


}, 6)

tmp10233 := Call(__e, ns2_1set, symshen_4callrec, tmp10222)


_ = tmp10233

tmp10234 := MakeNative(func(__e *ControlFlow) {
V1652 := __e.Get(1)
_ = V1652
tmp10251 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1652)
}
__typedArg0 := V1652
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10242 Obj

if True == tmp10251 {
tmp10249 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1652)
}
__typedArg0 := V1652
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10250 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10249)
}
__typedArg0 := tmp10249
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10244 Obj

if True == tmp10250 {
tmp10246 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1652)
}
__typedArg0 := V1652
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10247 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10246)
}
__typedArg0 := tmp10246
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10248 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_5_1_1, tmp10247)
}
__typedArg0 := sym_5_1_1
__typedArg1 := tmp10247
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10245 Obj

if True == tmp10248 {
ifres10245 = True


} else {
ifres10245 = False


}

ifres10244 = ifres10245


} else {
ifres10244 = False


}

var ifres10243 Obj

if True == ifres10244 {
ifres10243 = True


} else {
ifres10243 = False


}

ifres10242 = ifres10243


} else {
ifres10242 = False


}

if True == ifres10242 {
tmp10235 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1652)
}
__typedArg0 := V1652
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10236 := Call(__e, PrimFunc(symshen_4predicate), tmp10235)


let__9914 := tmp10236
_ = let__9914

tmp10237 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp10238 := Call(__e, PrimFunc(symget), let__9914, symshen_4dynamic, tmp10237)


let__9915 := tmp10238
_ = let__9915

tmp10239 := Call(__e, PrimFunc(symshen_4retract_1clause), V1652, let__9915)


tmp10240 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symput), let__9914, symshen_4dynamic, tmp10239, tmp10240)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symretract)
return
}


}, 1)

tmp10252 := Call(__e, ns2_1set, symretract, tmp10234)


_ = tmp10252

tmp10253 := MakeNative(func(__e *ControlFlow) {
V1660 := __e.Get(1)
_ = V1660
V1661 := __e.Get(2)
_ = V1661
tmp10283 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V1661)
}
__typedArg0 := Nil
__typedArg1 := V1661
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp10283 {
__e.Return(Nil)
return
} else {
tmp10281 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1661)
}
__typedArg0 := V1661
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10266 Obj

if True == tmp10281 {
tmp10279 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1661)
}
__typedArg0 := V1661
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10280 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10279)
}
__typedArg0 := tmp10279
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10268 Obj

if True == tmp10280 {
tmp10276 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1661)
}
__typedArg0 := V1661
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10277 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10276)
}
__typedArg0 := tmp10276
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10278 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10277)
}
__typedArg0 := tmp10277
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10270 Obj

if True == tmp10278 {
tmp10272 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1661)
}
__typedArg0 := V1661
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10273 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10272)
}
__typedArg0 := tmp10272
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10274 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10273)
}
__typedArg0 := tmp10273
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10275 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V1660, tmp10274)
}
__typedArg0 := V1660
__typedArg1 := tmp10274
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10271 Obj

if True == tmp10275 {
ifres10271 = True


} else {
ifres10271 = False


}

ifres10270 = ifres10271


} else {
ifres10270 = False


}

var ifres10269 Obj

if True == ifres10270 {
ifres10269 = True


} else {
ifres10269 = False


}

ifres10268 = ifres10269


} else {
ifres10268 = False


}

var ifres10267 Obj

if True == ifres10268 {
ifres10267 = True


} else {
ifres10267 = False


}

ifres10266 = ifres10267


} else {
ifres10266 = False


}

if True == ifres10266 {
tmp10254 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1661)
}
__typedArg0 := V1661
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10255 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10254)
}
__typedArg0 := tmp10254
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10256 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10255)
}
__typedArg0 := tmp10255
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10257 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dnames_d)
}
__typedArg0 := symshen_4_dnames_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp10258 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp10256, tmp10257)
}
__typedArg0 := tmp10256
__typedArg1 := tmp10257
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10259 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dnames_d, tmp10258)
}
__typedArg0 := symshen_4_dnames_d
__typedArg1 := tmp10258
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp10259

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1661)
}
__typedArg0 := V1661
return Call(__e, PrimFunc(symtl), __typedArg0)
})())
return


} else {
tmp10264 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1661)
}
__typedArg0 := V1661
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp10264 {
tmp10260 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1661)
}
__typedArg0 := V1661
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10261 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1661)
}
__typedArg0 := V1661
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10262 := Call(__e, PrimFunc(symshen_4retract_1clause), V1660, tmp10261)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp10260, tmp10262)
}
__typedArg0 := tmp10260
__typedArg1 := tmp10262
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

tmp10284 := Call(__e, ns2_1set, symshen_4retract_1clause, tmp10253)


_ = tmp10284

tmp10285 := MakeNative(func(__e *ControlFlow) {
V1662 := __e.Get(1)
_ = V1662
V1663 := __e.Get(2)
_ = V1663
tmp10286 := MakeNative(func(__e *ControlFlow) {
Z1664 := __e.Get(1)
_ = Z1664
__e.TailApply(PrimFunc(symshen_4_5defprolog_6), Z1664)
return
}, 1)

tmp10287 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1662, V1663)
}
__typedArg0 := V1662
__typedArg1 := V1663
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symcompile), tmp10286, tmp10287)
return


}, 2)

tmp10288 := Call(__e, ns2_1set, symshen_4compile_1prolog, tmp10285)


_ = tmp10288

tmp10289 := MakeNative(func(__e *ControlFlow) {
V1665 := __e.Get(1)
_ = V1665
tmp10305 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1665)
}
__typedArg0 := V1665
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10290 Obj

if True == tmp10305 {
tmp10291 := Call(__e, PrimFunc(symhead), V1665)


let__9917 := tmp10291
_ = let__9917

tmp10292 := Call(__e, PrimFunc(symtail), V1665)


let__9918 := tmp10292
_ = let__9918

tmp10293 := Call(__e, PrimFunc(symshen_4_5clauses_6), let__9918)


let__9919 := tmp10293
_ = let__9919

tmp10303 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9919)


var ifres10294 Obj

if True == tmp10303 {
tmp10295 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10294 = tmp10295


} else {
tmp10296 := Call(__e, PrimFunc(symshen_4_5_1out), let__9919)


let__9920 := tmp10296
_ = let__9920

tmp10297 := Call(__e, PrimFunc(symshen_4in_1_6), let__9919)


let__9921 := tmp10297
_ = let__9921

tmp10298 := Call(__e, PrimFunc(symshen_4prolog_1arity_1check), let__9917, let__9920)


let__9922 := tmp10298
_ = let__9922

tmp10299 := MakeNative(func(__e *ControlFlow) {
Z1674 := __e.Get(1)
_ = Z1674
__e.TailApply(PrimFunc(symshen_4linearise_1clause), Z1674)
return
}, 1)

tmp10300 := Call(__e, PrimFunc(symmap), tmp10299, let__9920)


let__9923 := tmp10300
_ = let__9923

tmp10301 := Call(__e, PrimFunc(symshen_4horn_1clause_1procedure), let__9917, let__9923)


tmp10302 := Call(__e, PrimFunc(symshen_4comb), let__9921, tmp10301)


ifres10294 = tmp10302


}

ifres10290 = ifres10294


} else {
tmp10304 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10290 = tmp10304


}

let__9916 := ifres10290
_ = let__9916

tmp10307 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9916)


if True == tmp10307 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__9916)
return
}


}, 1)

tmp10308 := Call(__e, ns2_1set, symshen_4_5defprolog_6, tmp10289)


_ = tmp10308

tmp10309 := MakeNative(func(__e *ControlFlow) {
V1677 := __e.Get(1)
_ = V1677
V1678 := __e.Get(2)
_ = V1678
tmp10353 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1678)
}
__typedArg0 := V1678
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10334 Obj

if True == tmp10353 {
tmp10351 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1678)
}
__typedArg0 := V1678
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10352 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10351)
}
__typedArg0 := tmp10351
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10336 Obj

if True == tmp10352 {
tmp10348 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1678)
}
__typedArg0 := V1678
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10349 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10348)
}
__typedArg0 := tmp10348
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10350 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10349)
}
__typedArg0 := tmp10349
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10338 Obj

if True == tmp10350 {
tmp10344 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1678)
}
__typedArg0 := V1678
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10345 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10344)
}
__typedArg0 := tmp10344
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10346 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10345)
}
__typedArg0 := tmp10345
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10347 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp10346)
}
__typedArg0 := Nil
__typedArg1 := tmp10346
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10340 Obj

if True == tmp10347 {
tmp10342 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1678)
}
__typedArg0 := V1678
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10343 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp10342)
}
__typedArg0 := Nil
__typedArg1 := tmp10342
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10341 Obj

if True == tmp10343 {
ifres10341 = True


} else {
ifres10341 = False


}

ifres10340 = ifres10341


} else {
ifres10340 = False


}

var ifres10339 Obj

if True == ifres10340 {
ifres10339 = True


} else {
ifres10339 = False


}

ifres10338 = ifres10339


} else {
ifres10338 = False


}

var ifres10337 Obj

if True == ifres10338 {
ifres10337 = True


} else {
ifres10337 = False


}

ifres10336 = ifres10337


} else {
ifres10336 = False


}

var ifres10335 Obj

if True == ifres10336 {
ifres10335 = True


} else {
ifres10335 = False


}

ifres10334 = ifres10335


} else {
ifres10334 = False


}

if True == ifres10334 {
tmp10310 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1678)
}
__typedArg0 := V1678
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10311 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10310)
}
__typedArg0 := tmp10310
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(PrimFunc(symlength), tmp10311)
return


} else {
tmp10332 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1678)
}
__typedArg0 := V1678
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10317 Obj

if True == tmp10332 {
tmp10330 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1678)
}
__typedArg0 := V1678
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10331 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10330)
}
__typedArg0 := tmp10330
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10319 Obj

if True == tmp10331 {
tmp10327 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1678)
}
__typedArg0 := V1678
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10328 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10327)
}
__typedArg0 := tmp10327
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10329 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10328)
}
__typedArg0 := tmp10328
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10321 Obj

if True == tmp10329 {
tmp10323 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1678)
}
__typedArg0 := V1678
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10324 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10323)
}
__typedArg0 := tmp10323
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10325 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10324)
}
__typedArg0 := tmp10324
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10326 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp10325)
}
__typedArg0 := Nil
__typedArg1 := tmp10325
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10322 Obj

if True == tmp10326 {
ifres10322 = True


} else {
ifres10322 = False


}

ifres10321 = ifres10322


} else {
ifres10321 = False


}

var ifres10320 Obj

if True == ifres10321 {
ifres10320 = True


} else {
ifres10320 = False


}

ifres10319 = ifres10320


} else {
ifres10319 = False


}

var ifres10318 Obj

if True == ifres10319 {
ifres10318 = True


} else {
ifres10318 = False


}

ifres10317 = ifres10318


} else {
ifres10317 = False


}

if True == ifres10317 {
tmp10312 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1678)
}
__typedArg0 := V1678
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10313 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10312)
}
__typedArg0 := tmp10312
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10314 := Call(__e, PrimFunc(symlength), tmp10313)


tmp10315 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1678)
}
__typedArg0 := V1678
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4pac_1h), V1677, tmp10314, tmp10315)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4prolog_1arity_1check)
return
}


}


}, 2)

tmp10354 := Call(__e, ns2_1set, symshen_4prolog_1arity_1check, tmp10309)


_ = tmp10354

tmp10355 := MakeNative(func(__e *ControlFlow) {
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
tmp10371 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V1685)
}
__typedArg0 := Nil
__typedArg1 := V1685
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp10371 {
__e.Return(V1684)
return
} else {
tmp10369 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1685)
}
__typedArg0 := V1685
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10365 Obj

if True == tmp10369 {
tmp10367 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1685)
}
__typedArg0 := V1685
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10368 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10367)
}
__typedArg0 := tmp10367
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10366 Obj

if True == tmp10368 {
ifres10366 = True


} else {
ifres10366 = False


}

ifres10365 = ifres10366


} else {
ifres10365 = False


}

if True == ifres10365 {
tmp10360 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1685)
}
__typedArg0 := V1685
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10361 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10360)
}
__typedArg0 := tmp10360
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10362 := Call(__e, PrimFunc(symlength), tmp10361)


tmp10363 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V1684, tmp10362)
}
__typedArg0 := V1684
__typedArg1 := tmp10362
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp10363 {
tmp10356 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1685)
}
__typedArg0 := V1685
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4pac_1h) == __self {
__self1, __self2, __self3 = V1683, V1684, tmp10356
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4pac_1h), V1683, V1684, tmp10356)
return


} else {
tmp10357 := Call(__e, PrimFunc(symshen_4app), V1683, MakeString("\n"), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("arity error in prolog procedure "))
__typedS1, __typedOK1 := TypedString(tmp10357)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("arity error in prolog procedure ")
__typedArg1 := tmp10357
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("arity error in prolog procedure "))
__typedS1, __typedOK1 := TypedString(tmp10357)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("arity error in prolog procedure ")
__typedArg1 := tmp10357
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

tmp10372 := Call(__e, ns2_1set, symshen_4pac_1h, tmp10355)


_ = tmp10372

tmp10373 := MakeNative(func(__e *ControlFlow) {
V1686 := __e.Get(1)
_ = V1686
tmp10374 := Call(__e, PrimFunc(symshen_4_5clause_6), V1686)


let__9925 := tmp10374
_ = let__9925

tmp10387 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9925)


var ifres10375 Obj

if True == tmp10387 {
tmp10376 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10375 = tmp10376


} else {
tmp10377 := Call(__e, PrimFunc(symshen_4_5_1out), let__9925)


let__9926 := tmp10377
_ = let__9926

tmp10378 := Call(__e, PrimFunc(symshen_4in_1_6), let__9925)


let__9927 := tmp10378
_ = let__9927

tmp10379 := Call(__e, PrimFunc(symshen_4_5clauses_6), let__9927)


let__9928 := tmp10379
_ = let__9928

tmp10386 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9928)


var ifres10380 Obj

if True == tmp10386 {
tmp10381 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10380 = tmp10381


} else {
tmp10382 := Call(__e, PrimFunc(symshen_4_5_1out), let__9928)


let__9929 := tmp10382
_ = let__9929

tmp10383 := Call(__e, PrimFunc(symshen_4in_1_6), let__9928)


let__9930 := tmp10383
_ = let__9930

tmp10384 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__9926, let__9929)
}
__typedArg0 := let__9926
__typedArg1 := let__9929
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10385 := Call(__e, PrimFunc(symshen_4comb), let__9930, tmp10384)


ifres10380 = tmp10385


}

ifres10375 = ifres10380


}

let__9924 := ifres10375
_ = let__9924

tmp10403 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9924)


if True == tmp10403 {
tmp10388 := Call(__e, PrimFunc(sym_5_b_6), V1686)


let__9932 := tmp10388
_ = let__9932

tmp10399 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9932)


var ifres10389 Obj

if True == tmp10399 {
tmp10390 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10389 = tmp10390


} else {
tmp10391 := Call(__e, PrimFunc(symshen_4_5_1out), let__9932)


let__9933 := tmp10391
_ = let__9933

tmp10392 := Call(__e, PrimFunc(symshen_4in_1_6), let__9932)


let__9934 := tmp10392
_ = let__9934

tmp10397 := Call(__e, PrimFunc(symempty_2), let__9933)


var ifres10393 Obj

if True == tmp10397 {
ifres10393 = Nil


} else {
tmp10394 := Call(__e, PrimFunc(symshen_4app), let__9933, MakeString("\n ..."), symshen_4r)


tmp10396 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("Prolog syntax error here:\n "))
__typedS1, __typedOK1 := TypedString(tmp10394)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("Prolog syntax error here:\n ")
__typedArg1 := tmp10394
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("Prolog syntax error here:\n "))
__typedS1, __typedOK1 := TypedString(tmp10394)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("Prolog syntax error here:\n ")
__typedArg1 := tmp10394
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})()

ifres10393 = tmp10396


}

tmp10398 := Call(__e, PrimFunc(symshen_4comb), let__9934, ifres10393)


ifres10389 = tmp10398


}

let__9931 := ifres10389
_ = let__9931

tmp10401 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9931)


if True == tmp10401 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__9931)
return
}


} else {
__e.Return(let__9924)
return
}


}, 1)

tmp10404 := Call(__e, ns2_1set, symshen_4_5clauses_6, tmp10373)


_ = tmp10404

tmp10405 := MakeNative(func(__e *ControlFlow) {
V1698 := __e.Get(1)
_ = V1698
tmp10421 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1698)
}
__typedArg0 := V1698
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10412 Obj

if True == tmp10421 {
tmp10419 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1698)
}
__typedArg0 := V1698
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10420 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10419)
}
__typedArg0 := tmp10419
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10414 Obj

if True == tmp10420 {
tmp10416 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1698)
}
__typedArg0 := V1698
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10417 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10416)
}
__typedArg0 := tmp10416
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10418 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp10417)
}
__typedArg0 := Nil
__typedArg1 := tmp10417
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10415 Obj

if True == tmp10418 {
ifres10415 = True


} else {
ifres10415 = False


}

ifres10414 = ifres10415


} else {
ifres10414 = False


}

var ifres10413 Obj

if True == ifres10414 {
ifres10413 = True


} else {
ifres10413 = False


}

ifres10412 = ifres10413


} else {
ifres10412 = False


}

if True == ifres10412 {
tmp10406 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1698)
}
__typedArg0 := V1698
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10407 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1698)
}
__typedArg0 := V1698
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10408 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10407)
}
__typedArg0 := tmp10407
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10409 := Call(__e, PrimFunc(sym_8p), tmp10406, tmp10408)


tmp10410 := Call(__e, PrimFunc(symshen_4linearise), tmp10409)


__e.TailApply(PrimFunc(symshen_4lch), tmp10410)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4linearise_1clause)
return
}


}, 1)

tmp10422 := Call(__e, ns2_1set, symshen_4linearise_1clause, tmp10405)


_ = tmp10422

tmp10423 := MakeNative(func(__e *ControlFlow) {
V1699 := __e.Get(1)
_ = V1699
tmp10429 := Call(__e, PrimFunc(symtuple_2), V1699)


if True == tmp10429 {
tmp10424 := Call(__e, PrimFunc(symfst), V1699)


tmp10425 := Call(__e, PrimFunc(symsnd), V1699)


tmp10426 := Call(__e, PrimFunc(symshen_4lchh), tmp10425)


tmp10427 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp10426, Nil)
}
__typedArg0 := tmp10426
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp10424, tmp10427)
}
__typedArg0 := tmp10424
__typedArg1 := tmp10427
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4lch)
return
}


}, 1)

tmp10430 := Call(__e, ns2_1set, symshen_4lch, tmp10423)


_ = tmp10430

tmp10431 := MakeNative(func(__e *ControlFlow) {
V1700 := __e.Get(1)
_ = V1700
tmp10494 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1700)
}
__typedArg0 := V1700
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10443 Obj

if True == tmp10494 {
tmp10492 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1700)
}
__typedArg0 := V1700
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10493 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symwhere, tmp10492)
}
__typedArg0 := symwhere
__typedArg1 := tmp10492
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10445 Obj

if True == tmp10493 {
tmp10490 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1700)
}
__typedArg0 := V1700
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10491 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10490)
}
__typedArg0 := tmp10490
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10447 Obj

if True == tmp10491 {
tmp10487 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1700)
}
__typedArg0 := V1700
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10488 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10487)
}
__typedArg0 := tmp10487
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10489 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10488)
}
__typedArg0 := tmp10488
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10449 Obj

if True == tmp10489 {
tmp10483 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1700)
}
__typedArg0 := V1700
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10484 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10483)
}
__typedArg0 := tmp10483
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10485 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10484)
}
__typedArg0 := tmp10484
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10486 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_a, tmp10485)
}
__typedArg0 := sym_a
__typedArg1 := tmp10485
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10451 Obj

if True == tmp10486 {
tmp10479 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1700)
}
__typedArg0 := V1700
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10480 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10479)
}
__typedArg0 := tmp10479
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10481 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10480)
}
__typedArg0 := tmp10480
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10482 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10481)
}
__typedArg0 := tmp10481
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10453 Obj

if True == tmp10482 {
tmp10474 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1700)
}
__typedArg0 := V1700
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10475 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10474)
}
__typedArg0 := tmp10474
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10476 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10475)
}
__typedArg0 := tmp10475
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10477 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10476)
}
__typedArg0 := tmp10476
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10478 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10477)
}
__typedArg0 := tmp10477
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10455 Obj

if True == tmp10478 {
tmp10468 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1700)
}
__typedArg0 := V1700
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10469 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10468)
}
__typedArg0 := tmp10468
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10470 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10469)
}
__typedArg0 := tmp10469
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10471 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10470)
}
__typedArg0 := tmp10470
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10472 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10471)
}
__typedArg0 := tmp10471
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10473 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp10472)
}
__typedArg0 := Nil
__typedArg1 := tmp10472
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10457 Obj

if True == tmp10473 {
tmp10465 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1700)
}
__typedArg0 := V1700
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10466 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10465)
}
__typedArg0 := tmp10465
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10467 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10466)
}
__typedArg0 := tmp10466
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10459 Obj

if True == tmp10467 {
tmp10461 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1700)
}
__typedArg0 := V1700
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10462 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10461)
}
__typedArg0 := tmp10461
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10463 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10462)
}
__typedArg0 := tmp10462
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10464 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp10463)
}
__typedArg0 := Nil
__typedArg1 := tmp10463
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10460 Obj

if True == tmp10464 {
ifres10460 = True


} else {
ifres10460 = False


}

ifres10459 = ifres10460


} else {
ifres10459 = False


}

var ifres10458 Obj

if True == ifres10459 {
ifres10458 = True


} else {
ifres10458 = False


}

ifres10457 = ifres10458


} else {
ifres10457 = False


}

var ifres10456 Obj

if True == ifres10457 {
ifres10456 = True


} else {
ifres10456 = False


}

ifres10455 = ifres10456


} else {
ifres10455 = False


}

var ifres10454 Obj

if True == ifres10455 {
ifres10454 = True


} else {
ifres10454 = False


}

ifres10453 = ifres10454


} else {
ifres10453 = False


}

var ifres10452 Obj

if True == ifres10453 {
ifres10452 = True


} else {
ifres10452 = False


}

ifres10451 = ifres10452


} else {
ifres10451 = False


}

var ifres10450 Obj

if True == ifres10451 {
ifres10450 = True


} else {
ifres10450 = False


}

ifres10449 = ifres10450


} else {
ifres10449 = False


}

var ifres10448 Obj

if True == ifres10449 {
ifres10448 = True


} else {
ifres10448 = False


}

ifres10447 = ifres10448


} else {
ifres10447 = False


}

var ifres10446 Obj

if True == ifres10447 {
ifres10446 = True


} else {
ifres10446 = False


}

ifres10445 = ifres10446


} else {
ifres10445 = False


}

var ifres10444 Obj

if True == ifres10445 {
ifres10444 = True


} else {
ifres10444 = False


}

ifres10443 = ifres10444


} else {
ifres10443 = False


}

if True == ifres10443 {
tmp10433 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_doccurs_d)
}
__typedArg0 := symshen_4_doccurs_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

var ifres10432 Obj

if True == tmp10433 {
ifres10432 = symis_b


} else {
ifres10432 = symis


}

tmp10434 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1700)
}
__typedArg0 := V1700
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10435 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10434)
}
__typedArg0 := tmp10434
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10436 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10435)
}
__typedArg0 := tmp10435
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10437 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(ifres10432, tmp10436)
}
__typedArg0 := ifres10432
__typedArg1 := tmp10436
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10438 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1700)
}
__typedArg0 := V1700
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10439 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10438)
}
__typedArg0 := tmp10438
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10440 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10439)
}
__typedArg0 := tmp10439
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10441 := Call(__e, PrimFunc(symshen_4lchh), tmp10440)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp10437, tmp10441)
}
__typedArg0 := tmp10437
__typedArg1 := tmp10441
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return(V1700)
return
}


}, 1)

tmp10495 := Call(__e, ns2_1set, symshen_4lchh, tmp10431)


_ = tmp10495

tmp10496 := MakeNative(func(__e *ControlFlow) {
V1701 := __e.Get(1)
_ = V1701
tmp10497 := Call(__e, PrimFunc(symshen_4_5head_6), V1701)


let__9936 := tmp10497
_ = let__9936

tmp10520 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9936)


var ifres10498 Obj

if True == tmp10520 {
tmp10499 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10498 = tmp10499


} else {
tmp10500 := Call(__e, PrimFunc(symshen_4_5_1out), let__9936)


let__9937 := tmp10500
_ = let__9937

tmp10501 := Call(__e, PrimFunc(symshen_4in_1_6), let__9936)


let__9938 := tmp10501
_ = let__9938

tmp10519 := Call(__e, PrimFunc(symshen_4hds_a_2), let__9938, sym_5_1_1)


var ifres10502 Obj

if True == tmp10519 {
tmp10503 := Call(__e, PrimFunc(symtail), let__9938)


let__9939 := tmp10503
_ = let__9939

tmp10504 := Call(__e, PrimFunc(symshen_4_5body_6), let__9939)


let__9940 := tmp10504
_ = let__9940

tmp10517 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9940)


var ifres10505 Obj

if True == tmp10517 {
tmp10506 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10505 = tmp10506


} else {
tmp10507 := Call(__e, PrimFunc(symshen_4_5_1out), let__9940)


let__9941 := tmp10507
_ = let__9941

tmp10508 := Call(__e, PrimFunc(symshen_4in_1_6), let__9940)


let__9942 := tmp10508
_ = let__9942

tmp10509 := Call(__e, PrimFunc(symshen_4_5sc_6), let__9942)


let__9943 := tmp10509
_ = let__9943

tmp10516 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9943)


var ifres10510 Obj

if True == tmp10516 {
tmp10511 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10510 = tmp10511


} else {
tmp10512 := Call(__e, PrimFunc(symshen_4in_1_6), let__9943)


let__9944 := tmp10512
_ = let__9944

tmp10513 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__9941, Nil)
}
__typedArg0 := let__9941
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10514 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__9937, tmp10513)
}
__typedArg0 := let__9937
__typedArg1 := tmp10513
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10515 := Call(__e, PrimFunc(symshen_4comb), let__9944, tmp10514)


ifres10510 = tmp10515


}

ifres10505 = ifres10510


}

ifres10502 = ifres10505


} else {
tmp10518 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10502 = tmp10518


}

ifres10498 = ifres10502


}

let__9935 := ifres10498
_ = let__9935

tmp10522 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9935)


if True == tmp10522 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__9935)
return
}


}, 1)

tmp10523 := Call(__e, ns2_1set, symshen_4_5clause_6, tmp10496)


_ = tmp10523

tmp10524 := MakeNative(func(__e *ControlFlow) {
V1712 := __e.Get(1)
_ = V1712
tmp10525 := Call(__e, PrimFunc(symshen_4_5hterm_6), V1712)


let__9946 := tmp10525
_ = let__9946

tmp10538 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9946)


var ifres10526 Obj

if True == tmp10538 {
tmp10527 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10526 = tmp10527


} else {
tmp10528 := Call(__e, PrimFunc(symshen_4_5_1out), let__9946)


let__9947 := tmp10528
_ = let__9947

tmp10529 := Call(__e, PrimFunc(symshen_4in_1_6), let__9946)


let__9948 := tmp10529
_ = let__9948

tmp10530 := Call(__e, PrimFunc(symshen_4_5head_6), let__9948)


let__9949 := tmp10530
_ = let__9949

tmp10537 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9949)


var ifres10531 Obj

if True == tmp10537 {
tmp10532 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10531 = tmp10532


} else {
tmp10533 := Call(__e, PrimFunc(symshen_4_5_1out), let__9949)


let__9950 := tmp10533
_ = let__9950

tmp10534 := Call(__e, PrimFunc(symshen_4in_1_6), let__9949)


let__9951 := tmp10534
_ = let__9951

tmp10535 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__9947, let__9950)
}
__typedArg0 := let__9947
__typedArg1 := let__9950
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10536 := Call(__e, PrimFunc(symshen_4comb), let__9951, tmp10535)


ifres10531 = tmp10536


}

ifres10526 = ifres10531


}

let__9945 := ifres10526
_ = let__9945

tmp10548 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9945)


if True == tmp10548 {
tmp10539 := Call(__e, PrimFunc(sym_5e_6), V1712)


let__9953 := tmp10539
_ = let__9953

tmp10544 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9953)


var ifres10540 Obj

if True == tmp10544 {
tmp10541 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10540 = tmp10541


} else {
tmp10542 := Call(__e, PrimFunc(symshen_4in_1_6), let__9953)


let__9954 := tmp10542
_ = let__9954

tmp10543 := Call(__e, PrimFunc(symshen_4comb), let__9954, Nil)


ifres10540 = tmp10543


}

let__9952 := ifres10540
_ = let__9952

tmp10546 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9952)


if True == tmp10546 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__9952)
return
}


} else {
__e.Return(let__9945)
return
}


}, 1)

tmp10549 := Call(__e, ns2_1set, symshen_4_5head_6, tmp10524)


_ = tmp10549

tmp10550 := MakeNative(func(__e *ControlFlow) {
V1723 := __e.Get(1)
_ = V1723
tmp10563 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1723)
}
__typedArg0 := V1723
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10551 Obj

if True == tmp10563 {
tmp10552 := Call(__e, PrimFunc(symhead), V1723)


let__9956 := tmp10552
_ = let__9956

tmp10553 := Call(__e, PrimFunc(symtail), V1723)


let__9957 := tmp10553
_ = let__9957

tmp10561 := Call(__e, PrimFunc(symatom_2), let__9956)


var ifres10557 Obj

if True == tmp10561 {
tmp10559 := Call(__e, PrimFunc(symshen_4prolog_1keyword_2), let__9956)


tmp10560 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp10559)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp10559
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres10558 Obj

if True == tmp10560 {
ifres10558 = True


} else {
ifres10558 = False


}

ifres10557 = ifres10558


} else {
ifres10557 = False


}

var ifres10554 Obj

if True == ifres10557 {
tmp10555 := Call(__e, PrimFunc(symshen_4comb), let__9957, let__9956)


ifres10554 = tmp10555


} else {
tmp10556 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10554 = tmp10556


}

ifres10551 = ifres10554


} else {
tmp10562 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10551 = tmp10562


}

let__9955 := ifres10551
_ = let__9955

tmp10717 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9955)


if True == tmp10717 {
tmp10573 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1723)
}
__typedArg0 := V1723
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10564 Obj

if True == tmp10573 {
tmp10565 := Call(__e, PrimFunc(symhead), V1723)


let__9959 := tmp10565
_ = let__9959

tmp10566 := Call(__e, PrimFunc(symtail), V1723)


let__9960 := tmp10566
_ = let__9960

tmp10570 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp10571 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__9959, tmp10570)
}
__typedArg0 := let__9959
__typedArg1 := tmp10570
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10567 Obj

if True == tmp10571 {
tmp10568 := Call(__e, PrimFunc(symshen_4comb), let__9960, let__9959)


ifres10567 = tmp10568


} else {
tmp10569 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10567 = tmp10569


}

ifres10564 = ifres10567


} else {
tmp10572 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10564 = tmp10572


}

let__9958 := ifres10564
_ = let__9958

tmp10715 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9958)


if True == tmp10715 {
tmp10603 := Call(__e, PrimFunc(symshen_4ccons_2), V1723)


var ifres10574 Obj

if True == tmp10603 {
tmp10575 := Call(__e, PrimFunc(symhead), V1723)


let__9962 := tmp10575
_ = let__9962

tmp10576 := Call(__e, PrimFunc(symtail), V1723)


let__9963 := tmp10576
_ = let__9963

tmp10601 := Call(__e, PrimFunc(symshen_4hds_a_2), let__9962, symcons)


var ifres10577 Obj

if True == tmp10601 {
tmp10578 := Call(__e, PrimFunc(symtail), let__9962)


let__9964 := tmp10578
_ = let__9964

tmp10579 := Call(__e, PrimFunc(symshen_4_5hterm1_6), let__9964)


let__9965 := tmp10579
_ = let__9965

tmp10599 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9965)


var ifres10580 Obj

if True == tmp10599 {
tmp10581 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10580 = tmp10581


} else {
tmp10582 := Call(__e, PrimFunc(symshen_4_5_1out), let__9965)


let__9966 := tmp10582
_ = let__9966

tmp10583 := Call(__e, PrimFunc(symshen_4in_1_6), let__9965)


let__9967 := tmp10583
_ = let__9967

tmp10584 := Call(__e, PrimFunc(symshen_4_5hterm2_6), let__9967)


let__9968 := tmp10584
_ = let__9968

tmp10598 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9968)


var ifres10585 Obj

if True == tmp10598 {
tmp10586 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10585 = tmp10586


} else {
tmp10587 := Call(__e, PrimFunc(symshen_4_5_1out), let__9968)


let__9969 := tmp10587
_ = let__9969

tmp10588 := Call(__e, PrimFunc(symshen_4in_1_6), let__9968)


let__9970 := tmp10588
_ = let__9970

tmp10589 := Call(__e, PrimFunc(sym_5end_6), let__9970)


let__9971 := tmp10589
_ = let__9971

tmp10597 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9971)


var ifres10590 Obj

if True == tmp10597 {
tmp10591 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10590 = tmp10591


} else {
tmp10592 := Call(__e, PrimFunc(symshen_4in_1_6), let__9971)


let__9972 := tmp10592
_ = let__9972

tmp10593 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__9969, Nil)
}
__typedArg0 := let__9969
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10594 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__9966, tmp10593)
}
__typedArg0 := let__9966
__typedArg1 := tmp10593
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10595 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp10594)
}
__typedArg0 := symcons
__typedArg1 := tmp10594
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10596 := Call(__e, PrimFunc(symshen_4comb), let__9963, tmp10595)


ifres10590 = tmp10596


}

ifres10585 = ifres10590


}

ifres10580 = ifres10585


}

ifres10577 = ifres10580


} else {
tmp10600 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10577 = tmp10600


}

ifres10574 = ifres10577


} else {
tmp10602 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10574 = tmp10602


}

let__9961 := ifres10574
_ = let__9961

tmp10713 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9961)


if True == tmp10713 {
tmp10626 := Call(__e, PrimFunc(symshen_4ccons_2), V1723)


var ifres10604 Obj

if True == tmp10626 {
tmp10605 := Call(__e, PrimFunc(symhead), V1723)


let__9974 := tmp10605
_ = let__9974

tmp10606 := Call(__e, PrimFunc(symtail), V1723)


let__9975 := tmp10606
_ = let__9975

tmp10624 := Call(__e, PrimFunc(symshen_4hds_a_2), let__9974, sym_7)


var ifres10607 Obj

if True == tmp10624 {
tmp10608 := Call(__e, PrimFunc(symtail), let__9974)


let__9976 := tmp10608
_ = let__9976

tmp10609 := Call(__e, PrimFunc(symshen_4_5hterm_6), let__9976)


let__9977 := tmp10609
_ = let__9977

tmp10622 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9977)


var ifres10610 Obj

if True == tmp10622 {
tmp10611 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10610 = tmp10611


} else {
tmp10612 := Call(__e, PrimFunc(symshen_4_5_1out), let__9977)


let__9978 := tmp10612
_ = let__9978

tmp10613 := Call(__e, PrimFunc(symshen_4in_1_6), let__9977)


let__9979 := tmp10613
_ = let__9979

tmp10614 := Call(__e, PrimFunc(sym_5end_6), let__9979)


let__9980 := tmp10614
_ = let__9980

tmp10621 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9980)


var ifres10615 Obj

if True == tmp10621 {
tmp10616 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10615 = tmp10616


} else {
tmp10617 := Call(__e, PrimFunc(symshen_4in_1_6), let__9980)


let__9981 := tmp10617
_ = let__9981

tmp10618 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__9978, Nil)
}
__typedArg0 := let__9978
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10619 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4_7m, tmp10618)
}
__typedArg0 := symshen_4_7m
__typedArg1 := tmp10618
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10620 := Call(__e, PrimFunc(symshen_4comb), let__9975, tmp10619)


ifres10615 = tmp10620


}

ifres10610 = ifres10615


}

ifres10607 = ifres10610


} else {
tmp10623 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10607 = tmp10623


}

ifres10604 = ifres10607


} else {
tmp10625 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10604 = tmp10625


}

let__9973 := ifres10604
_ = let__9973

tmp10711 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9973)


if True == tmp10711 {
tmp10649 := Call(__e, PrimFunc(symshen_4ccons_2), V1723)


var ifres10627 Obj

if True == tmp10649 {
tmp10628 := Call(__e, PrimFunc(symhead), V1723)


let__9983 := tmp10628
_ = let__9983

tmp10629 := Call(__e, PrimFunc(symtail), V1723)


let__9984 := tmp10629
_ = let__9984

tmp10647 := Call(__e, PrimFunc(symshen_4hds_a_2), let__9983, sym_1)


var ifres10630 Obj

if True == tmp10647 {
tmp10631 := Call(__e, PrimFunc(symtail), let__9983)


let__9985 := tmp10631
_ = let__9985

tmp10632 := Call(__e, PrimFunc(symshen_4_5hterm_6), let__9985)


let__9986 := tmp10632
_ = let__9986

tmp10645 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9986)


var ifres10633 Obj

if True == tmp10645 {
tmp10634 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10633 = tmp10634


} else {
tmp10635 := Call(__e, PrimFunc(symshen_4_5_1out), let__9986)


let__9987 := tmp10635
_ = let__9987

tmp10636 := Call(__e, PrimFunc(symshen_4in_1_6), let__9986)


let__9988 := tmp10636
_ = let__9988

tmp10637 := Call(__e, PrimFunc(sym_5end_6), let__9988)


let__9989 := tmp10637
_ = let__9989

tmp10644 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9989)


var ifres10638 Obj

if True == tmp10644 {
tmp10639 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10638 = tmp10639


} else {
tmp10640 := Call(__e, PrimFunc(symshen_4in_1_6), let__9989)


let__9990 := tmp10640
_ = let__9990

tmp10641 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__9987, Nil)
}
__typedArg0 := let__9987
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10642 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4_1m, tmp10641)
}
__typedArg0 := symshen_4_1m
__typedArg1 := tmp10641
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10643 := Call(__e, PrimFunc(symshen_4comb), let__9984, tmp10642)


ifres10638 = tmp10643


}

ifres10633 = ifres10638


}

ifres10630 = ifres10633


} else {
tmp10646 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10630 = tmp10646


}

ifres10627 = ifres10630


} else {
tmp10648 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10627 = tmp10648


}

let__9982 := ifres10627
_ = let__9982

tmp10709 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9982)


if True == tmp10709 {
tmp10676 := Call(__e, PrimFunc(symshen_4ccons_2), V1723)


var ifres10650 Obj

if True == tmp10676 {
tmp10651 := Call(__e, PrimFunc(symhead), V1723)


let__9992 := tmp10651
_ = let__9992

tmp10652 := Call(__e, PrimFunc(symtail), V1723)


let__9993 := tmp10652
_ = let__9993

tmp10674 := Call(__e, PrimFunc(symshen_4hds_a_2), let__9992, symmode)


var ifres10653 Obj

if True == tmp10674 {
tmp10654 := Call(__e, PrimFunc(symtail), let__9992)


let__9994 := tmp10654
_ = let__9994

tmp10655 := Call(__e, PrimFunc(symshen_4_5hterm_6), let__9994)


let__9995 := tmp10655
_ = let__9995

tmp10672 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9995)


var ifres10656 Obj

if True == tmp10672 {
tmp10657 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10656 = tmp10657


} else {
tmp10658 := Call(__e, PrimFunc(symshen_4_5_1out), let__9995)


let__9996 := tmp10658
_ = let__9996

tmp10659 := Call(__e, PrimFunc(symshen_4in_1_6), let__9995)


let__9997 := tmp10659
_ = let__9997

tmp10671 := Call(__e, PrimFunc(symshen_4hds_a_2), let__9997, sym_7)


var ifres10660 Obj

if True == tmp10671 {
tmp10661 := Call(__e, PrimFunc(symtail), let__9997)


let__9998 := tmp10661
_ = let__9998

tmp10662 := Call(__e, PrimFunc(sym_5end_6), let__9998)


let__9999 := tmp10662
_ = let__9999

tmp10669 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9999)


var ifres10663 Obj

if True == tmp10669 {
tmp10664 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10663 = tmp10664


} else {
tmp10665 := Call(__e, PrimFunc(symshen_4in_1_6), let__9999)


let__10000 := tmp10665
_ = let__10000

tmp10666 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__9996, Nil)
}
__typedArg0 := let__9996
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10667 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4_7m, tmp10666)
}
__typedArg0 := symshen_4_7m
__typedArg1 := tmp10666
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10668 := Call(__e, PrimFunc(symshen_4comb), let__9993, tmp10667)


ifres10663 = tmp10668


}

ifres10660 = ifres10663


} else {
tmp10670 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10660 = tmp10670


}

ifres10656 = ifres10660


}

ifres10653 = ifres10656


} else {
tmp10673 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10653 = tmp10673


}

ifres10650 = ifres10653


} else {
tmp10675 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10650 = tmp10675


}

let__9991 := ifres10650
_ = let__9991

tmp10707 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__9991)


if True == tmp10707 {
tmp10703 := Call(__e, PrimFunc(symshen_4ccons_2), V1723)


var ifres10677 Obj

if True == tmp10703 {
tmp10678 := Call(__e, PrimFunc(symhead), V1723)


let__10002 := tmp10678
_ = let__10002

tmp10679 := Call(__e, PrimFunc(symtail), V1723)


let__10003 := tmp10679
_ = let__10003

tmp10701 := Call(__e, PrimFunc(symshen_4hds_a_2), let__10002, symmode)


var ifres10680 Obj

if True == tmp10701 {
tmp10681 := Call(__e, PrimFunc(symtail), let__10002)


let__10004 := tmp10681
_ = let__10004

tmp10682 := Call(__e, PrimFunc(symshen_4_5hterm_6), let__10004)


let__10005 := tmp10682
_ = let__10005

tmp10699 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10005)


var ifres10683 Obj

if True == tmp10699 {
tmp10684 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10683 = tmp10684


} else {
tmp10685 := Call(__e, PrimFunc(symshen_4_5_1out), let__10005)


let__10006 := tmp10685
_ = let__10006

tmp10686 := Call(__e, PrimFunc(symshen_4in_1_6), let__10005)


let__10007 := tmp10686
_ = let__10007

tmp10698 := Call(__e, PrimFunc(symshen_4hds_a_2), let__10007, sym_1)


var ifres10687 Obj

if True == tmp10698 {
tmp10688 := Call(__e, PrimFunc(symtail), let__10007)


let__10008 := tmp10688
_ = let__10008

tmp10689 := Call(__e, PrimFunc(sym_5end_6), let__10008)


let__10009 := tmp10689
_ = let__10009

tmp10696 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10009)


var ifres10690 Obj

if True == tmp10696 {
tmp10691 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10690 = tmp10691


} else {
tmp10692 := Call(__e, PrimFunc(symshen_4in_1_6), let__10009)


let__10010 := tmp10692
_ = let__10010

tmp10693 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10006, Nil)
}
__typedArg0 := let__10006
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10694 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4_1m, tmp10693)
}
__typedArg0 := symshen_4_1m
__typedArg1 := tmp10693
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10695 := Call(__e, PrimFunc(symshen_4comb), let__10003, tmp10694)


ifres10690 = tmp10695


}

ifres10687 = ifres10690


} else {
tmp10697 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10687 = tmp10697


}

ifres10683 = ifres10687


}

ifres10680 = ifres10683


} else {
tmp10700 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10680 = tmp10700


}

ifres10677 = ifres10680


} else {
tmp10702 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10677 = tmp10702


}

let__10001 := ifres10677
_ = let__10001

tmp10705 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10001)


if True == tmp10705 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__10001)
return
}


} else {
__e.Return(let__9991)
return
}


} else {
__e.Return(let__9982)
return
}


} else {
__e.Return(let__9973)
return
}


} else {
__e.Return(let__9961)
return
}


} else {
__e.Return(let__9958)
return
}


} else {
__e.Return(let__9955)
return
}


}, 1)

tmp10718 := Call(__e, ns2_1set, symshen_4_5hterm_6, tmp10550)


_ = tmp10718

tmp10719 := MakeNative(func(__e *ControlFlow) {
V1780 := __e.Get(1)
_ = V1780
tmp10720 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(";"))
}
__typedArg0 := MakeString(";")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp10721 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1_1, Nil)
}
__typedArg0 := sym_5_1_1
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10722 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp10720, tmp10721)
}
__typedArg0 := tmp10720
__typedArg1 := tmp10721
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symelement_2), V1780, tmp10722)
return


}, 1)

tmp10723 := Call(__e, ns2_1set, symshen_4prolog_1keyword_2, tmp10719)


_ = tmp10723

tmp10724 := MakeNative(func(__e *ControlFlow) {
V1781 := __e.Get(1)
_ = V1781
tmp10737 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsymbol_2) {
return PrimIsSymbol(V1781)
}
__typedArg0 := V1781
return Call(__e, PrimFunc(symsymbol_2), __typedArg0)
})()

if True == tmp10737 {
__e.Return(True)
return
} else {
tmp10735 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_2) {
return PrimIsString(V1781)
}
__typedArg0 := V1781
return Call(__e, PrimFunc(symstring_2), __typedArg0)
})()

var ifres10726 Obj

if True == tmp10735 {
ifres10726 = True


} else {
tmp10734 := Call(__e, PrimFunc(symboolean_2), V1781)


var ifres10728 Obj

if True == tmp10734 {
ifres10728 = True


} else {
tmp10733 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnumber_2) {
return PrimIsNumber(V1781)
}
__typedArg0 := V1781
return Call(__e, PrimFunc(symnumber_2), __typedArg0)
})()

var ifres10730 Obj

if True == tmp10733 {
ifres10730 = True


} else {
tmp10732 := Call(__e, PrimFunc(symempty_2), V1781)


var ifres10731 Obj

if True == tmp10732 {
ifres10731 = True


} else {
ifres10731 = False


}

ifres10730 = ifres10731


}

var ifres10729 Obj

if True == ifres10730 {
ifres10729 = True


} else {
ifres10729 = False


}

ifres10728 = ifres10729


}

var ifres10727 Obj

if True == ifres10728 {
ifres10727 = True


} else {
ifres10727 = False


}

ifres10726 = ifres10727


}

if True == ifres10726 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp10738 := Call(__e, ns2_1set, symatom_2, tmp10724)


_ = tmp10738

tmp10739 := MakeNative(func(__e *ControlFlow) {
V1782 := __e.Get(1)
_ = V1782
tmp10740 := Call(__e, PrimFunc(symshen_4_5hterm_6), V1782)


let__10012 := tmp10740
_ = let__10012

tmp10746 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10012)


var ifres10741 Obj

if True == tmp10746 {
tmp10742 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10741 = tmp10742


} else {
tmp10743 := Call(__e, PrimFunc(symshen_4_5_1out), let__10012)


let__10013 := tmp10743
_ = let__10013

tmp10744 := Call(__e, PrimFunc(symshen_4in_1_6), let__10012)


let__10014 := tmp10744
_ = let__10014

tmp10745 := Call(__e, PrimFunc(symshen_4comb), let__10014, let__10013)


ifres10741 = tmp10745


}

let__10011 := ifres10741
_ = let__10011

tmp10748 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10011)


if True == tmp10748 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__10011)
return
}


}, 1)

tmp10749 := Call(__e, ns2_1set, symshen_4_5hterm1_6, tmp10739)


_ = tmp10749

tmp10750 := MakeNative(func(__e *ControlFlow) {
V1787 := __e.Get(1)
_ = V1787
tmp10751 := Call(__e, PrimFunc(symshen_4_5hterm_6), V1787)


let__10016 := tmp10751
_ = let__10016

tmp10757 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10016)


var ifres10752 Obj

if True == tmp10757 {
tmp10753 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10752 = tmp10753


} else {
tmp10754 := Call(__e, PrimFunc(symshen_4_5_1out), let__10016)


let__10017 := tmp10754
_ = let__10017

tmp10755 := Call(__e, PrimFunc(symshen_4in_1_6), let__10016)


let__10018 := tmp10755
_ = let__10018

tmp10756 := Call(__e, PrimFunc(symshen_4comb), let__10018, let__10017)


ifres10752 = tmp10756


}

let__10015 := ifres10752
_ = let__10015

tmp10759 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10015)


if True == tmp10759 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__10015)
return
}


}, 1)

tmp10760 := Call(__e, ns2_1set, symshen_4_5hterm2_6, tmp10750)


_ = tmp10760

tmp10761 := MakeNative(func(__e *ControlFlow) {
V1792 := __e.Get(1)
_ = V1792
tmp10762 := Call(__e, PrimFunc(symshen_4_5literal_6), V1792)


let__10020 := tmp10762
_ = let__10020

tmp10775 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10020)


var ifres10763 Obj

if True == tmp10775 {
tmp10764 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10763 = tmp10764


} else {
tmp10765 := Call(__e, PrimFunc(symshen_4_5_1out), let__10020)


let__10021 := tmp10765
_ = let__10021

tmp10766 := Call(__e, PrimFunc(symshen_4in_1_6), let__10020)


let__10022 := tmp10766
_ = let__10022

tmp10767 := Call(__e, PrimFunc(symshen_4_5body_6), let__10022)


let__10023 := tmp10767
_ = let__10023

tmp10774 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10023)


var ifres10768 Obj

if True == tmp10774 {
tmp10769 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10768 = tmp10769


} else {
tmp10770 := Call(__e, PrimFunc(symshen_4_5_1out), let__10023)


let__10024 := tmp10770
_ = let__10024

tmp10771 := Call(__e, PrimFunc(symshen_4in_1_6), let__10023)


let__10025 := tmp10771
_ = let__10025

tmp10772 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10021, let__10024)
}
__typedArg0 := let__10021
__typedArg1 := let__10024
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10773 := Call(__e, PrimFunc(symshen_4comb), let__10025, tmp10772)


ifres10768 = tmp10773


}

ifres10763 = ifres10768


}

let__10019 := ifres10763
_ = let__10019

tmp10785 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10019)


if True == tmp10785 {
tmp10776 := Call(__e, PrimFunc(sym_5e_6), V1792)


let__10027 := tmp10776
_ = let__10027

tmp10781 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10027)


var ifres10777 Obj

if True == tmp10781 {
tmp10778 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10777 = tmp10778


} else {
tmp10779 := Call(__e, PrimFunc(symshen_4in_1_6), let__10027)


let__10028 := tmp10779
_ = let__10028

tmp10780 := Call(__e, PrimFunc(symshen_4comb), let__10028, Nil)


ifres10777 = tmp10780


}

let__10026 := ifres10777
_ = let__10026

tmp10783 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10026)


if True == tmp10783 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__10026)
return
}


} else {
__e.Return(let__10019)
return
}


}, 1)

tmp10786 := Call(__e, ns2_1set, symshen_4_5body_6, tmp10761)


_ = tmp10786

tmp10787 := MakeNative(func(__e *ControlFlow) {
V1803 := __e.Get(1)
_ = V1803
tmp10792 := Call(__e, PrimFunc(symshen_4hds_a_2), V1803, sym_b)


var ifres10788 Obj

if True == tmp10792 {
tmp10789 := Call(__e, PrimFunc(symtail), V1803)


let__10030 := tmp10789
_ = let__10030

tmp10790 := Call(__e, PrimFunc(symshen_4comb), let__10030, sym_b)


ifres10788 = tmp10790


} else {
tmp10791 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10788 = tmp10791


}

let__10029 := ifres10788
_ = let__10029

tmp10813 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10029)


if True == tmp10813 {
tmp10809 := Call(__e, PrimFunc(symshen_4ccons_2), V1803)


var ifres10793 Obj

if True == tmp10809 {
tmp10794 := Call(__e, PrimFunc(symhead), V1803)


let__10032 := tmp10794
_ = let__10032

tmp10795 := Call(__e, PrimFunc(symtail), V1803)


let__10033 := tmp10795
_ = let__10033

tmp10796 := Call(__e, PrimFunc(symshen_4_5bterms_6), let__10032)


let__10034 := tmp10796
_ = let__10034

tmp10807 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10034)


var ifres10797 Obj

if True == tmp10807 {
tmp10798 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10797 = tmp10798


} else {
tmp10799 := Call(__e, PrimFunc(symshen_4_5_1out), let__10034)


let__10035 := tmp10799
_ = let__10035

tmp10800 := Call(__e, PrimFunc(symshen_4in_1_6), let__10034)


let__10036 := tmp10800
_ = let__10036

tmp10801 := Call(__e, PrimFunc(sym_5end_6), let__10036)


let__10037 := tmp10801
_ = let__10037

tmp10806 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10037)


var ifres10802 Obj

if True == tmp10806 {
tmp10803 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10802 = tmp10803


} else {
tmp10804 := Call(__e, PrimFunc(symshen_4in_1_6), let__10037)


let__10038 := tmp10804
_ = let__10038

tmp10805 := Call(__e, PrimFunc(symshen_4comb), let__10033, let__10035)


ifres10802 = tmp10805


}

ifres10797 = ifres10802


}

ifres10793 = ifres10797


} else {
tmp10808 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10793 = tmp10808


}

let__10031 := ifres10793
_ = let__10031

tmp10811 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10031)


if True == tmp10811 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__10031)
return
}


} else {
__e.Return(let__10029)
return
}


}, 1)

tmp10814 := Call(__e, ns2_1set, symshen_4_5literal_6, tmp10787)


_ = tmp10814

tmp10815 := MakeNative(func(__e *ControlFlow) {
V1814 := __e.Get(1)
_ = V1814
tmp10816 := Call(__e, PrimFunc(symshen_4_5bterm_6), V1814)


let__10040 := tmp10816
_ = let__10040

tmp10829 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10040)


var ifres10817 Obj

if True == tmp10829 {
tmp10818 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10817 = tmp10818


} else {
tmp10819 := Call(__e, PrimFunc(symshen_4_5_1out), let__10040)


let__10041 := tmp10819
_ = let__10041

tmp10820 := Call(__e, PrimFunc(symshen_4in_1_6), let__10040)


let__10042 := tmp10820
_ = let__10042

tmp10821 := Call(__e, PrimFunc(symshen_4_5bterms_6), let__10042)


let__10043 := tmp10821
_ = let__10043

tmp10828 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10043)


var ifres10822 Obj

if True == tmp10828 {
tmp10823 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10822 = tmp10823


} else {
tmp10824 := Call(__e, PrimFunc(symshen_4_5_1out), let__10043)


let__10044 := tmp10824
_ = let__10044

tmp10825 := Call(__e, PrimFunc(symshen_4in_1_6), let__10043)


let__10045 := tmp10825
_ = let__10045

tmp10826 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10041, let__10044)
}
__typedArg0 := let__10041
__typedArg1 := let__10044
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10827 := Call(__e, PrimFunc(symshen_4comb), let__10045, tmp10826)


ifres10822 = tmp10827


}

ifres10817 = ifres10822


}

let__10039 := ifres10817
_ = let__10039

tmp10839 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10039)


if True == tmp10839 {
tmp10830 := Call(__e, PrimFunc(sym_5e_6), V1814)


let__10047 := tmp10830
_ = let__10047

tmp10835 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10047)


var ifres10831 Obj

if True == tmp10835 {
tmp10832 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10831 = tmp10832


} else {
tmp10833 := Call(__e, PrimFunc(symshen_4in_1_6), let__10047)


let__10048 := tmp10833
_ = let__10048

tmp10834 := Call(__e, PrimFunc(symshen_4comb), let__10048, Nil)


ifres10831 = tmp10834


}

let__10046 := ifres10831
_ = let__10046

tmp10837 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10046)


if True == tmp10837 {
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

tmp10840 := Call(__e, ns2_1set, symshen_4_5bterms_6, tmp10815)


_ = tmp10840

tmp10841 := MakeNative(func(__e *ControlFlow) {
V1825 := __e.Get(1)
_ = V1825
tmp10842 := Call(__e, PrimFunc(symshen_4_5wildcard_6), V1825)


let__10050 := tmp10842
_ = let__10050

tmp10848 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10050)


var ifres10843 Obj

if True == tmp10848 {
tmp10844 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10843 = tmp10844


} else {
tmp10845 := Call(__e, PrimFunc(symshen_4_5_1out), let__10050)


let__10051 := tmp10845
_ = let__10051

tmp10846 := Call(__e, PrimFunc(symshen_4in_1_6), let__10050)


let__10052 := tmp10846
_ = let__10052

tmp10847 := Call(__e, PrimFunc(symshen_4comb), let__10052, let__10051)


ifres10843 = tmp10847


}

let__10049 := ifres10843
_ = let__10049

tmp10880 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10049)


if True == tmp10880 {
tmp10857 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1825)
}
__typedArg0 := V1825
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10849 Obj

if True == tmp10857 {
tmp10850 := Call(__e, PrimFunc(symhead), V1825)


let__10054 := tmp10850
_ = let__10054

tmp10851 := Call(__e, PrimFunc(symtail), V1825)


let__10055 := tmp10851
_ = let__10055

tmp10855 := Call(__e, PrimFunc(symatom_2), let__10054)


var ifres10852 Obj

if True == tmp10855 {
tmp10853 := Call(__e, PrimFunc(symshen_4comb), let__10055, let__10054)


ifres10852 = tmp10853


} else {
tmp10854 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10852 = tmp10854


}

ifres10849 = ifres10852


} else {
tmp10856 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10849 = tmp10856


}

let__10053 := ifres10849
_ = let__10053

tmp10878 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10053)


if True == tmp10878 {
tmp10874 := Call(__e, PrimFunc(symshen_4ccons_2), V1825)


var ifres10858 Obj

if True == tmp10874 {
tmp10859 := Call(__e, PrimFunc(symhead), V1825)


let__10057 := tmp10859
_ = let__10057

tmp10860 := Call(__e, PrimFunc(symtail), V1825)


let__10058 := tmp10860
_ = let__10058

tmp10861 := Call(__e, PrimFunc(symshen_4_5bterms_6), let__10057)


let__10059 := tmp10861
_ = let__10059

tmp10872 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10059)


var ifres10862 Obj

if True == tmp10872 {
tmp10863 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10862 = tmp10863


} else {
tmp10864 := Call(__e, PrimFunc(symshen_4_5_1out), let__10059)


let__10060 := tmp10864
_ = let__10060

tmp10865 := Call(__e, PrimFunc(symshen_4in_1_6), let__10059)


let__10061 := tmp10865
_ = let__10061

tmp10866 := Call(__e, PrimFunc(sym_5end_6), let__10061)


let__10062 := tmp10866
_ = let__10062

tmp10871 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10062)


var ifres10867 Obj

if True == tmp10871 {
tmp10868 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10867 = tmp10868


} else {
tmp10869 := Call(__e, PrimFunc(symshen_4in_1_6), let__10062)


let__10063 := tmp10869
_ = let__10063

tmp10870 := Call(__e, PrimFunc(symshen_4comb), let__10058, let__10060)


ifres10867 = tmp10870


}

ifres10862 = ifres10867


}

ifres10858 = ifres10862


} else {
tmp10873 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10858 = tmp10873


}

let__10056 := ifres10858
_ = let__10056

tmp10876 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10056)


if True == tmp10876 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__10056)
return
}


} else {
__e.Return(let__10053)
return
}


} else {
__e.Return(let__10049)
return
}


}, 1)

tmp10881 := Call(__e, ns2_1set, symshen_4_5bterm_6, tmp10841)


_ = tmp10881

tmp10882 := MakeNative(func(__e *ControlFlow) {
V1841 := __e.Get(1)
_ = V1841
tmp10892 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1841)
}
__typedArg0 := V1841
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10883 Obj

if True == tmp10892 {
tmp10884 := Call(__e, PrimFunc(symhead), V1841)


let__10065 := tmp10884
_ = let__10065

tmp10885 := Call(__e, PrimFunc(symtail), V1841)


let__10066 := tmp10885
_ = let__10066

tmp10890 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__10065, sym__)
}
__typedArg0 := let__10065
__typedArg1 := sym__
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10886 Obj

if True == tmp10890 {
tmp10887 := Call(__e, PrimFunc(symgensym), symY)


tmp10888 := Call(__e, PrimFunc(symshen_4comb), let__10066, tmp10887)


ifres10886 = tmp10888


} else {
tmp10889 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10886 = tmp10889


}

ifres10883 = ifres10886


} else {
tmp10891 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10883 = tmp10891


}

let__10064 := ifres10883
_ = let__10064

tmp10894 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10064)


if True == tmp10894 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__10064)
return
}


}, 1)

tmp10895 := Call(__e, ns2_1set, symshen_4_5wildcard_6, tmp10882)


_ = tmp10895

tmp10896 := MakeNative(func(__e *ControlFlow) {
V1845 := __e.Get(1)
_ = V1845
tmp10905 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1845)
}
__typedArg0 := V1845
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10897 Obj

if True == tmp10905 {
tmp10898 := Call(__e, PrimFunc(symhead), V1845)


let__10068 := tmp10898
_ = let__10068

tmp10899 := Call(__e, PrimFunc(symtail), V1845)


let__10069 := tmp10899
_ = let__10069

tmp10903 := Call(__e, PrimFunc(symshen_4semicolon_2), let__10068)


var ifres10900 Obj

if True == tmp10903 {
tmp10901 := Call(__e, PrimFunc(symshen_4comb), let__10069, let__10068)


ifres10900 = tmp10901


} else {
tmp10902 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10900 = tmp10902


}

ifres10897 = ifres10900


} else {
tmp10904 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10897 = tmp10904


}

let__10067 := ifres10897
_ = let__10067

tmp10907 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10067)


if True == tmp10907 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__10067)
return
}


}, 1)

tmp10908 := Call(__e, ns2_1set, symshen_4_5sc_6, tmp10896)


_ = tmp10908

tmp10909 := MakeNative(func(__e *ControlFlow) {
V1849 := __e.Get(1)
_ = V1849
V1850 := __e.Get(2)
_ = V1850
tmp10910 := Call(__e, PrimFunc(symgensym), symB)


let__10070 := tmp10910
_ = let__10070

tmp10911 := Call(__e, PrimFunc(symgensym), symL)


let__10071 := tmp10911
_ = let__10071

tmp10912 := Call(__e, PrimFunc(symgensym), symK)


let__10072 := tmp10912
_ = let__10072

tmp10913 := Call(__e, PrimFunc(symgensym), symC)


let__10073 := tmp10913
_ = let__10073

tmp10914 := Call(__e, PrimFunc(symshen_4prolog_1parameters), V1850)


let__10074 := tmp10914
_ = let__10074

tmp10915 := Call(__e, PrimFunc(symshen_4hascut_2), V1850)


let__10075 := tmp10915
_ = let__10075

tmp10916 := Call(__e, PrimFunc(symshen_4prolog_1fbody), V1850, let__10074, let__10070, let__10071, let__10072, let__10073, let__10075)


let__10076 := tmp10916
_ = let__10076

var ifres10917 Obj

if True == let__10075 {
tmp10918 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), Nil)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10919 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10072, tmp10918)
}
__typedArg0 := let__10072
__typedArg1 := tmp10918
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10920 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_7, tmp10919)
}
__typedArg0 := sym_7
__typedArg1 := tmp10919
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10921 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10076, Nil)
}
__typedArg0 := let__10076
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10922 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp10920, tmp10921)
}
__typedArg0 := tmp10920
__typedArg1 := tmp10921
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10923 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10072, tmp10922)
}
__typedArg0 := let__10072
__typedArg1 := tmp10922
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10924 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp10923)
}
__typedArg0 := symlet
__typedArg1 := tmp10923
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

ifres10917 = tmp10924


} else {
ifres10917 = let__10076


}

let__10077 := ifres10917
_ = let__10077

tmp10925 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_6, Nil)
}
__typedArg0 := sym_1_6
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10926 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10073, tmp10925)
}
__typedArg0 := let__10073
__typedArg1 := tmp10925
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10927 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10072, tmp10926)
}
__typedArg0 := let__10072
__typedArg1 := tmp10926
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10928 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10071, tmp10927)
}
__typedArg0 := let__10071
__typedArg1 := tmp10927
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10929 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10070, tmp10928)
}
__typedArg0 := let__10070
__typedArg1 := tmp10928
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10930 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10077, Nil)
}
__typedArg0 := let__10077
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10931 := Call(__e, PrimFunc(symappend), tmp10929, tmp10930)


tmp10932 := Call(__e, PrimFunc(symappend), let__10074, tmp10931)


tmp10933 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1849, tmp10932)
}
__typedArg0 := V1849
__typedArg1 := tmp10932
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10934 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefine, tmp10933)
}
__typedArg0 := symdefine
__typedArg1 := tmp10933
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__10078 := tmp10934
_ = let__10078

__e.Return(let__10078)
return


}, 2)

tmp10935 := Call(__e, ns2_1set, symshen_4horn_1clause_1procedure, tmp10909)


_ = tmp10935

tmp10936 := MakeNative(func(__e *ControlFlow) {
V1862 := __e.Get(1)
_ = V1862
tmp10946 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_b, V1862)
}
__typedArg0 := sym_b
__typedArg1 := V1862
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp10946 {
__e.Return(True)
return
} else {
tmp10944 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1862)
}
__typedArg0 := V1862
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp10944 {
tmp10941 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1862)
}
__typedArg0 := V1862
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10942 := Call(__e, PrimFunc(symshen_4hascut_2), tmp10941)


if True == tmp10942 {
__e.Return(True)
return
} else {
tmp10938 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1862)
}
__typedArg0 := V1862
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10939 := Call(__e, PrimFunc(symshen_4hascut_2), tmp10938)


if True == tmp10939 {
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

tmp10947 := Call(__e, ns2_1set, symshen_4hascut_2, tmp10936)


_ = tmp10947

tmp10948 := MakeNative(func(__e *ControlFlow) {
V1867 := __e.Get(1)
_ = V1867
tmp10957 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1867)
}
__typedArg0 := V1867
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10953 Obj

if True == tmp10957 {
tmp10955 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1867)
}
__typedArg0 := V1867
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10956 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10955)
}
__typedArg0 := tmp10955
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10954 Obj

if True == tmp10956 {
ifres10954 = True


} else {
ifres10954 = False


}

ifres10953 = ifres10954


} else {
ifres10953 = False


}

if True == ifres10953 {
tmp10949 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1867)
}
__typedArg0 := V1867
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10950 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10949)
}
__typedArg0 := tmp10949
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10951 := Call(__e, PrimFunc(symlength), tmp10950)


__e.TailApply(PrimFunc(symshen_4parameters), tmp10951)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4prolog_1parameters)
return
}


}, 1)

tmp10958 := Call(__e, ns2_1set, symshen_4prolog_1parameters, tmp10948)


_ = tmp10958

tmp10959 := MakeNative(func(__e *ControlFlow) {
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
tmp11049 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V1888)
}
__typedArg0 := Nil
__typedArg1 := V1888
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11046 Obj

if True == tmp11049 {
tmp11048 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(True, V1894)
}
__typedArg0 := True
__typedArg1 := V1894
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

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
tmp10960 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1892, Nil)
}
__typedArg0 := V1892
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10961 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1891, tmp10960)
}
__typedArg0 := V1891
__typedArg1 := tmp10960
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4unlock, tmp10961)
}
__typedArg0 := symshen_4unlock
__typedArg1 := tmp10961
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp11044 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1888)
}
__typedArg0 := V1888
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11022 Obj

if True == tmp11044 {
tmp11042 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1888)
}
__typedArg0 := V1888
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11043 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11042)
}
__typedArg0 := tmp11042
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11024 Obj

if True == tmp11043 {
tmp11039 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1888)
}
__typedArg0 := V1888
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11040 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11039)
}
__typedArg0 := tmp11039
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11041 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11040)
}
__typedArg0 := tmp11040
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11026 Obj

if True == tmp11041 {
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11036)
}
__typedArg0 := tmp11036
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11038 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp11037)
}
__typedArg0 := Nil
__typedArg1 := tmp11037
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11028 Obj

if True == tmp11038 {
tmp11033 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1888)
}
__typedArg0 := V1888
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
tmp11032 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(False, V1894)
}
__typedArg0 := False
__typedArg1 := V1894
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11031 Obj

if True == tmp11032 {
ifres11031 = True


} else {
ifres11031 = False


}

ifres11030 = ifres11031


} else {
ifres11030 = False


}

var ifres11029 Obj

if True == ifres11030 {
ifres11029 = True


} else {
ifres11029 = False


}

ifres11028 = ifres11029


} else {
ifres11028 = False


}

var ifres11027 Obj

if True == ifres11028 {
ifres11027 = True


} else {
ifres11027 = False


}

ifres11026 = ifres11027


} else {
ifres11026 = False


}

var ifres11025 Obj

if True == ifres11026 {
ifres11025 = True


} else {
ifres11025 = False


}

ifres11024 = ifres11025


} else {
ifres11024 = False


}

var ifres11023 Obj

if True == ifres11024 {
ifres11023 = True


} else {
ifres11023 = False


}

ifres11022 = ifres11023


} else {
ifres11022 = False


}

if True == ifres11022 {
tmp10962 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1888)
}
__typedArg0 := V1888
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10963 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10962)
}
__typedArg0 := tmp10962
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10964 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1888)
}
__typedArg0 := V1888
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10965 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10964)
}
__typedArg0 := tmp10964
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10966 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10965)
}
__typedArg0 := tmp10965
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10967 := Call(__e, PrimFunc(symshen_4continue), tmp10963, tmp10966, V1890, V1891, V1892, V1893)


let__10079 := tmp10967
_ = let__10079

tmp10968 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1891, Nil)
}
__typedArg0 := V1891
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10969 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4unlocked_2, tmp10968)
}
__typedArg0 := symshen_4unlocked_2
__typedArg1 := tmp10968
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

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

tmp10972 := Call(__e, PrimFunc(symshen_4compile_1head), symshen_4_7m, tmp10971, V1889, V1890, let__10079)


tmp10973 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(False, Nil)
}
__typedArg0 := False
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10974 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp10972, tmp10973)
}
__typedArg0 := tmp10972
__typedArg1 := tmp10973
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10975 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp10969, tmp10974)
}
__typedArg0 := tmp10969
__typedArg1 := tmp10974
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp10975)
}
__typedArg0 := symif
__typedArg1 := tmp10975
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp11020 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1888)
}
__typedArg0 := V1888
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11005 Obj

if True == tmp11020 {
tmp11018 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1888)
}
__typedArg0 := V1888
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11019 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11018)
}
__typedArg0 := tmp11018
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11007 Obj

if True == tmp11019 {
tmp11015 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1888)
}
__typedArg0 := V1888
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11016 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11015)
}
__typedArg0 := tmp11015
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11017 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11016)
}
__typedArg0 := tmp11016
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11009 Obj

if True == tmp11017 {
tmp11011 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1888)
}
__typedArg0 := V1888
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11012 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11011)
}
__typedArg0 := tmp11011
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11013 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11012)
}
__typedArg0 := tmp11012
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11014 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp11013)
}
__typedArg0 := Nil
__typedArg1 := tmp11013
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11010 Obj

if True == tmp11014 {
ifres11010 = True


} else {
ifres11010 = False


}

ifres11009 = ifres11010


} else {
ifres11009 = False


}

var ifres11008 Obj

if True == ifres11009 {
ifres11008 = True


} else {
ifres11008 = False


}

ifres11007 = ifres11008


} else {
ifres11007 = False


}

var ifres11006 Obj

if True == ifres11007 {
ifres11006 = True


} else {
ifres11006 = False


}

ifres11005 = ifres11006


} else {
ifres11005 = False


}

if True == ifres11005 {
tmp10976 := Call(__e, PrimFunc(symgensym), symC)


let__10080 := tmp10976
_ = let__10080

tmp10977 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1888)
}
__typedArg0 := V1888
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10978 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10977)
}
__typedArg0 := tmp10977
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10979 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1888)
}
__typedArg0 := V1888
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10980 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10979)
}
__typedArg0 := tmp10979
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10981 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10980)
}
__typedArg0 := tmp10980
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10982 := Call(__e, PrimFunc(symshen_4continue), tmp10978, tmp10981, V1890, V1891, V1892, V1893)


let__10081 := tmp10982
_ = let__10081

tmp10983 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1891, Nil)
}
__typedArg0 := V1891
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10984 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4unlocked_2, tmp10983)
}
__typedArg0 := symshen_4unlocked_2
__typedArg1 := tmp10983
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10985 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1888)
}
__typedArg0 := V1888
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10986 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10985)
}
__typedArg0 := tmp10985
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10987 := Call(__e, PrimFunc(symshen_4compile_1head), symshen_4_7m, tmp10986, V1889, V1890, let__10081)


tmp10988 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(False, Nil)
}
__typedArg0 := False
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10989 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp10987, tmp10988)
}
__typedArg0 := tmp10987
__typedArg1 := tmp10988
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10990 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp10984, tmp10989)
}
__typedArg0 := tmp10984
__typedArg1 := tmp10989
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10991 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp10990)
}
__typedArg0 := symif
__typedArg1 := tmp10990
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10992 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(False, Nil)
}
__typedArg0 := False
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10993 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10080, tmp10992)
}
__typedArg0 := let__10080
__typedArg1 := tmp10992
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10994 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_a, tmp10993)
}
__typedArg0 := sym_a
__typedArg1 := tmp10993
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10995 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1888)
}
__typedArg0 := V1888
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10996 := Call(__e, PrimFunc(symshen_4prolog_1fbody), tmp10995, V1889, V1890, V1891, V1892, V1893, V1894)


tmp10997 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10080, Nil)
}
__typedArg0 := let__10080
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10998 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp10996, tmp10997)
}
__typedArg0 := tmp10996
__typedArg1 := tmp10997
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10999 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp10994, tmp10998)
}
__typedArg0 := tmp10994
__typedArg1 := tmp10998
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11000 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp10999)
}
__typedArg0 := symif
__typedArg1 := tmp10999
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11001 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11000, Nil)
}
__typedArg0 := tmp11000
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11002 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp10991, tmp11001)
}
__typedArg0 := tmp10991
__typedArg1 := tmp11001
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11003 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10080, tmp11002)
}
__typedArg0 := let__10080
__typedArg1 := tmp11002
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp11003)
}
__typedArg0 := symlet
__typedArg1 := tmp11003
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

tmp11050 := Call(__e, ns2_1set, symshen_4prolog_1fbody, tmp10959)


_ = tmp11050

tmp11051 := MakeNative(func(__e *ControlFlow) {
V1898 := __e.Get(1)
_ = V1898
V1899 := __e.Get(2)
_ = V1899
tmp11056 := Call(__e, PrimFunc(symshen_4locked_2), V1898)


var ifres11053 Obj

if True == tmp11056 {
tmp11055 := Call(__e, PrimFunc(symshen_4fits_2), V1899, V1898)


var ifres11054 Obj

if True == tmp11055 {
ifres11054 = True


} else {
ifres11054 = False


}

ifres11053 = ifres11054


} else {
ifres11053 = False


}

if True == ifres11053 {
__e.TailApply(PrimFunc(symshen_4openlock), V1898)
return
} else {
__e.Return(False)
return
}


}, 2)

tmp11057 := Call(__e, ns2_1set, symshen_4unlock, tmp11051)


_ = tmp11057

tmp11058 := MakeNative(func(__e *ControlFlow) {
V1900 := __e.Get(1)
_ = V1900
tmp11059 := Call(__e, PrimFunc(symshen_4unlocked_2), V1900)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp11059)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp11059
return Call(__e, PrimFunc(symnot), __typedArg0)
})())
return


}, 1)

tmp11060 := Call(__e, ns2_1set, symshen_4locked_2, tmp11058)


_ = tmp11060

tmp11061 := MakeNative(func(__e *ControlFlow) {
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

tmp11062 := Call(__e, ns2_1set, symshen_4unlocked_2, tmp11061)


_ = tmp11062

tmp11063 := MakeNative(func(__e *ControlFlow) {
V1902 := __e.Get(1)
_ = V1902
tmp11064 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(V1902, MakeInteger(1), True)
}
__typedArg0 := V1902
__typedArg1 := MakeInteger(1)
__typedArg2 := True
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

_ = tmp11064

__e.Return(False)
return


}, 1)

tmp11065 := Call(__e, ns2_1set, symshen_4openlock, tmp11063)


_ = tmp11065

tmp11066 := MakeNative(func(__e *ControlFlow) {
V1903 := __e.Get(1)
_ = V1903
V1904 := __e.Get(2)
_ = V1904
tmp11067 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V1904, MakeInteger(2))
}
__typedArg0 := V1904
__typedArg1 := MakeInteger(2)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V1903, tmp11067)
}
__typedArg0 := V1903
__typedArg1 := tmp11067
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})())
return


}, 2)

tmp11068 := Call(__e, ns2_1set, symshen_4fits_2, tmp11066)


_ = tmp11068

tmp11069 := MakeNative(func(__e *ControlFlow) {
V1907 := __e.Get(1)
_ = V1907
V1908 := __e.Get(2)
_ = V1908
V1909 := __e.Get(3)
_ = V1909
V1910 := __e.Get(4)
_ = V1910
tmp11070 := Call(__e, PrimFunc(symthaw), V1910)


let__10082 := tmp11070
_ = let__10082

tmp11075 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__10082, False)
}
__typedArg0 := let__10082
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11072 Obj

if True == tmp11075 {
tmp11074 := Call(__e, PrimFunc(symshen_4unlocked_2), V1908)


var ifres11073 Obj

if True == tmp11074 {
ifres11073 = True


} else {
ifres11073 = False


}

ifres11072 = ifres11073


} else {
ifres11072 = False


}

if True == ifres11072 {
__e.TailApply(PrimFunc(symshen_4lock), V1909, V1908)
return
} else {
__e.Return(let__10082)
return
}


}, 4)

tmp11076 := Call(__e, ns2_1set, symshen_4cut, tmp11069)


_ = tmp11076

tmp11077 := MakeNative(func(__e *ControlFlow) {
V1912 := __e.Get(1)
_ = V1912
V1913 := __e.Get(2)
_ = V1913
tmp11078 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(V1913, MakeInteger(1), False)
}
__typedArg0 := V1913
__typedArg1 := MakeInteger(1)
__typedArg2 := False
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

let__10083 := tmp11078
_ = let__10083

tmp11079 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(V1913, MakeInteger(2), V1912)
}
__typedArg0 := V1913
__typedArg1 := MakeInteger(2)
__typedArg2 := V1912
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

let__10084 := tmp11079
_ = let__10084

__e.Return(False)
return


}, 2)

tmp11080 := Call(__e, ns2_1set, symshen_4lock, tmp11077)


_ = tmp11080

tmp11081 := MakeNative(func(__e *ControlFlow) {
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
tmp11082 := Call(__e, PrimFunc(symshen_4extract_1vars), V1916)


let__10085 := tmp11082
_ = let__10085

tmp11083 := Call(__e, PrimFunc(symshen_4extract_1free_1vars), V1917)


let__10086 := tmp11083
_ = let__10086

tmp11084 := Call(__e, PrimFunc(symdifference), let__10086, let__10085)


let__10087 := tmp11084
_ = let__10087

tmp11085 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4incinfs, Nil)
}
__typedArg0 := symshen_4incinfs
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11086 := Call(__e, PrimFunc(symshen_4compile_1body), V1917, V1918, V1919, V1920, V1921)


tmp11087 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11086, Nil)
}
__typedArg0 := tmp11086
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11088 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11085, tmp11087)
}
__typedArg0 := tmp11085
__typedArg1 := tmp11087
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11089 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdo, tmp11088)
}
__typedArg0 := symdo
__typedArg1 := tmp11088
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__10088 := tmp11089
_ = let__10088

__e.TailApply(PrimFunc(symshen_4stpart), let__10087, let__10088, V1918)
return


}, 6)

tmp11090 := Call(__e, ns2_1set, symshen_4continue, tmp11081)


_ = tmp11090

tmp11091 := MakeNative(func(__e *ControlFlow) {
V1928 := __e.Get(1)
_ = V1928
tmp11126 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1928)
}
__typedArg0 := V1928
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11107 Obj

if True == tmp11126 {
tmp11124 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1928)
}
__typedArg0 := V1928
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11125 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlambda, tmp11124)
}
__typedArg0 := symlambda
__typedArg1 := tmp11124
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11109 Obj

if True == tmp11125 {
tmp11122 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1928)
}
__typedArg0 := V1928
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11123 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11122)
}
__typedArg0 := tmp11122
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11111 Obj

if True == tmp11123 {
tmp11119 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1928)
}
__typedArg0 := V1928
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11120 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11119)
}
__typedArg0 := tmp11119
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11121 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11120)
}
__typedArg0 := tmp11120
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11113 Obj

if True == tmp11121 {
tmp11115 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1928)
}
__typedArg0 := V1928
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11116 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11115)
}
__typedArg0 := tmp11115
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11117 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11116)
}
__typedArg0 := tmp11116
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11118 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp11117)
}
__typedArg0 := Nil
__typedArg1 := tmp11117
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11114 Obj

if True == tmp11118 {
ifres11114 = True


} else {
ifres11114 = False


}

ifres11113 = ifres11114


} else {
ifres11113 = False


}

var ifres11112 Obj

if True == ifres11113 {
ifres11112 = True


} else {
ifres11112 = False


}

ifres11111 = ifres11112


} else {
ifres11111 = False


}

var ifres11110 Obj

if True == ifres11111 {
ifres11110 = True


} else {
ifres11110 = False


}

ifres11109 = ifres11110


} else {
ifres11109 = False


}

var ifres11108 Obj

if True == ifres11109 {
ifres11108 = True


} else {
ifres11108 = False


}

ifres11107 = ifres11108


} else {
ifres11107 = False


}

if True == ifres11107 {
tmp11092 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1928)
}
__typedArg0 := V1928
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11093 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11092)
}
__typedArg0 := tmp11092
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11094 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1928)
}
__typedArg0 := V1928
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11095 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11094)
}
__typedArg0 := tmp11094
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11096 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11095)
}
__typedArg0 := tmp11095
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11097 := Call(__e, PrimFunc(symshen_4extract_1free_1vars), tmp11096)


__e.TailApply(PrimFunc(symremove), tmp11093, tmp11097)
return


} else {
tmp11105 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1928)
}
__typedArg0 := V1928
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp11105 {
tmp11098 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1928)
}
__typedArg0 := V1928
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11099 := Call(__e, PrimFunc(symshen_4extract_1free_1vars), tmp11098)


tmp11100 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1928)
}
__typedArg0 := V1928
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11101 := Call(__e, PrimFunc(symshen_4extract_1free_1vars), tmp11100)


__e.TailApply(PrimFunc(symunion), tmp11099, tmp11101)
return


} else {
tmp11103 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(V1928)
}
__typedArg0 := V1928
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

if True == tmp11103 {
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

tmp11127 := Call(__e, ns2_1set, symshen_4extract_1free_1vars, tmp11091)


_ = tmp11127

tmp11128 := MakeNative(func(__e *ControlFlow) {
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
tmp11162 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V1945)
}
__typedArg0 := Nil
__typedArg1 := V1945
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp11162 {
tmp11129 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1949, Nil)
}
__typedArg0 := V1949
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symthaw, tmp11129)
}
__typedArg0 := symthaw
__typedArg1 := tmp11129
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp11160 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1945)
}
__typedArg0 := V1945
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11156 Obj

if True == tmp11160 {
tmp11158 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1945)
}
__typedArg0 := V1945
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11159 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_b, tmp11158)
}
__typedArg0 := sym_b
__typedArg1 := tmp11158
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11157 Obj

if True == tmp11159 {
ifres11157 = True


} else {
ifres11157 = False


}

ifres11156 = ifres11157


} else {
ifres11156 = False


}

if True == ifres11156 {
tmp11130 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4cut, Nil)
}
__typedArg0 := symshen_4cut
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11131 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1945)
}
__typedArg0 := V1945
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11132 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11130, tmp11131)
}
__typedArg0 := tmp11130
__typedArg1 := tmp11131
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

if PrimFunc(symshen_4compile_1body) == __self {
__self1, __self2, __self3, __self4, __self5 = tmp11132, V1946, V1947, V1948, V1949
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4compile_1body), tmp11132, V1946, V1947, V1948, V1949)
return


} else {
tmp11154 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1945)
}
__typedArg0 := V1945
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11150 Obj

if True == tmp11154 {
tmp11152 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1945)
}
__typedArg0 := V1945
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11153 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp11152)
}
__typedArg0 := Nil
__typedArg1 := tmp11152
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11151 Obj

if True == tmp11153 {
ifres11151 = True


} else {
ifres11151 = False


}

ifres11150 = ifres11151


} else {
ifres11150 = False


}

if True == ifres11150 {
tmp11133 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1945)
}
__typedArg0 := V1945
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11134 := Call(__e, PrimFunc(symshen_4deref_1calls), tmp11133, V1946)


tmp11135 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1949, Nil)
}
__typedArg0 := V1949
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11136 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1948, tmp11135)
}
__typedArg0 := V1948
__typedArg1 := tmp11135
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11137 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1947, tmp11136)
}
__typedArg0 := V1947
__typedArg1 := tmp11136
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11138 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1946, tmp11137)
}
__typedArg0 := V1946
__typedArg1 := tmp11137
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symappend), tmp11134, tmp11138)
return


} else {
tmp11148 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1945)
}
__typedArg0 := V1945
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp11148 {
tmp11139 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1945)
}
__typedArg0 := V1945
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11140 := Call(__e, PrimFunc(symshen_4deref_1calls), tmp11139, V1946)


let__10089 := tmp11140
_ = let__10089

tmp11141 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1945)
}
__typedArg0 := V1945
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11142 := Call(__e, PrimFunc(symshen_4freeze_1literals), tmp11141, V1946, V1947, V1948, V1949)


tmp11143 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11142, Nil)
}
__typedArg0 := tmp11142
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11144 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1948, tmp11143)
}
__typedArg0 := V1948
__typedArg1 := tmp11143
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11145 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1947, tmp11144)
}
__typedArg0 := V1947
__typedArg1 := tmp11144
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11146 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1946, tmp11145)
}
__typedArg0 := V1946
__typedArg1 := tmp11145
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symappend), let__10089, tmp11146)
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

tmp11163 := Call(__e, ns2_1set, symshen_4compile_1body, tmp11128)


_ = tmp11163

tmp11164 := MakeNative(func(__e *ControlFlow) {
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
tmp11187 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V1967)
}
__typedArg0 := Nil
__typedArg1 := V1967
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp11187 {
__e.Return(V1971)
return
} else {
tmp11185 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1967)
}
__typedArg0 := V1967
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11181 Obj

if True == tmp11185 {
tmp11183 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1967)
}
__typedArg0 := V1967
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11184 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_b, tmp11183)
}
__typedArg0 := sym_b
__typedArg1 := tmp11183
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11182 Obj

if True == tmp11184 {
ifres11182 = True


} else {
ifres11182 = False


}

ifres11181 = ifres11182


} else {
ifres11181 = False


}

if True == ifres11181 {
tmp11165 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4cut, Nil)
}
__typedArg0 := symshen_4cut
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11166 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1967)
}
__typedArg0 := V1967
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11167 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11165, tmp11166)
}
__typedArg0 := tmp11165
__typedArg1 := tmp11166
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

if PrimFunc(symshen_4freeze_1literals) == __self {
__self1, __self2, __self3, __self4, __self5 = tmp11167, V1968, V1969, V1970, V1971
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4freeze_1literals), tmp11167, V1968, V1969, V1970, V1971)
return


} else {
tmp11179 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1967)
}
__typedArg0 := V1967
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp11179 {
tmp11168 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1967)
}
__typedArg0 := V1967
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11169 := Call(__e, PrimFunc(symshen_4deref_1calls), tmp11168, V1968)


let__10090 := tmp11169
_ = let__10090

tmp11170 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1967)
}
__typedArg0 := V1967
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11171 := Call(__e, PrimFunc(symshen_4freeze_1literals), tmp11170, V1968, V1969, V1970, V1971)


tmp11172 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11171, Nil)
}
__typedArg0 := tmp11171
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11173 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1970, tmp11172)
}
__typedArg0 := V1970
__typedArg1 := tmp11172
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11174 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1969, tmp11173)
}
__typedArg0 := V1969
__typedArg1 := tmp11173
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11175 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1968, tmp11174)
}
__typedArg0 := V1968
__typedArg1 := tmp11174
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11176 := Call(__e, PrimFunc(symappend), let__10090, tmp11175)


tmp11177 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11176, Nil)
}
__typedArg0 := tmp11176
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfreeze, tmp11177)
}
__typedArg0 := symfreeze
__typedArg1 := tmp11177
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

tmp11188 := Call(__e, ns2_1set, symshen_4freeze_1literals, tmp11164)


_ = tmp11188

tmp11189 := MakeNative(func(__e *ControlFlow) {
V1977 := __e.Get(1)
_ = V1977
V1978 := __e.Get(2)
_ = V1978
tmp11204 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1977)
}
__typedArg0 := V1977
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11200 Obj

if True == tmp11204 {
tmp11202 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1977)
}
__typedArg0 := V1977
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11203 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symfork, tmp11202)
}
__typedArg0 := symfork
__typedArg1 := tmp11202
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11201 Obj

if True == tmp11203 {
ifres11201 = True


} else {
ifres11201 = False


}

ifres11200 = ifres11201


} else {
ifres11200 = False


}

if True == ifres11200 {
tmp11190 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1977)
}
__typedArg0 := V1977
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11191 := Call(__e, PrimFunc(symshen_4deref_1forked_1literals), tmp11190, V1978)


tmp11192 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11191, Nil)
}
__typedArg0 := tmp11191
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfork, tmp11192)
}
__typedArg0 := symfork
__typedArg1 := tmp11192
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp11198 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1977)
}
__typedArg0 := V1977
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp11198 {
tmp11193 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1977)
}
__typedArg0 := V1977
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11194 := MakeNative(func(__e *ControlFlow) {
Z1979 := __e.Get(1)
_ = Z1979
__e.TailApply(PrimFunc(symshen_4function_1calls), Z1979, V1978)
return
}, 1)

tmp11195 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1977)
}
__typedArg0 := V1977
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11196 := Call(__e, PrimFunc(symmap), tmp11194, tmp11195)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11193, tmp11196)
}
__typedArg0 := tmp11193
__typedArg1 := tmp11196
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

tmp11205 := Call(__e, ns2_1set, symshen_4deref_1calls, tmp11189)


_ = tmp11205

tmp11206 := MakeNative(func(__e *ControlFlow) {
V1986 := __e.Get(1)
_ = V1986
V1987 := __e.Get(2)
_ = V1987
tmp11216 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V1986)
}
__typedArg0 := Nil
__typedArg1 := V1986
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp11216 {
__e.Return(Nil)
return
} else {
tmp11214 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1986)
}
__typedArg0 := V1986
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp11214 {
tmp11207 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1986)
}
__typedArg0 := V1986
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11208 := Call(__e, PrimFunc(symshen_4deref_1calls), tmp11207, V1987)


tmp11209 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1986)
}
__typedArg0 := V1986
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11210 := Call(__e, PrimFunc(symshen_4deref_1forked_1literals), tmp11209, V1987)


tmp11211 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11210, Nil)
}
__typedArg0 := tmp11210
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11212 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11208, tmp11211)
}
__typedArg0 := tmp11208
__typedArg1 := tmp11211
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp11212)
}
__typedArg0 := symcons
__typedArg1 := tmp11212
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

tmp11217 := Call(__e, ns2_1set, symshen_4deref_1forked_1literals, tmp11206)


_ = tmp11217

tmp11218 := MakeNative(func(__e *ControlFlow) {
V1990 := __e.Get(1)
_ = V1990
V1991 := __e.Get(2)
_ = V1991
tmp11250 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1990)
}
__typedArg0 := V1990
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11231 Obj

if True == tmp11250 {
tmp11248 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1990)
}
__typedArg0 := V1990
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11249 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcons, tmp11248)
}
__typedArg0 := symcons
__typedArg1 := tmp11248
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11233 Obj

if True == tmp11249 {
tmp11246 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1990)
}
__typedArg0 := V1990
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11247 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11246)
}
__typedArg0 := tmp11246
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11235 Obj

if True == tmp11247 {
tmp11243 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1990)
}
__typedArg0 := V1990
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11244 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11243)
}
__typedArg0 := tmp11243
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11245 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11244)
}
__typedArg0 := tmp11244
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11237 Obj

if True == tmp11245 {
tmp11239 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1990)
}
__typedArg0 := V1990
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11240 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11239)
}
__typedArg0 := tmp11239
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11241 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11240)
}
__typedArg0 := tmp11240
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11242 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp11241)
}
__typedArg0 := Nil
__typedArg1 := tmp11241
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11238 Obj

if True == tmp11242 {
ifres11238 = True


} else {
ifres11238 = False


}

ifres11237 = ifres11238


} else {
ifres11237 = False


}

var ifres11236 Obj

if True == ifres11237 {
ifres11236 = True


} else {
ifres11236 = False


}

ifres11235 = ifres11236


} else {
ifres11235 = False


}

var ifres11234 Obj

if True == ifres11235 {
ifres11234 = True


} else {
ifres11234 = False


}

ifres11233 = ifres11234


} else {
ifres11233 = False


}

var ifres11232 Obj

if True == ifres11233 {
ifres11232 = True


} else {
ifres11232 = False


}

ifres11231 = ifres11232


} else {
ifres11231 = False


}

if True == ifres11231 {
tmp11219 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1990)
}
__typedArg0 := V1990
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11220 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11219)
}
__typedArg0 := tmp11219
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11221 := Call(__e, PrimFunc(symshen_4function_1calls), tmp11220, V1991)


tmp11222 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1990)
}
__typedArg0 := V1990
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11223 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11222)
}
__typedArg0 := tmp11222
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11224 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11223)
}
__typedArg0 := tmp11223
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11225 := Call(__e, PrimFunc(symshen_4function_1calls), tmp11224, V1991)


tmp11226 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11225, Nil)
}
__typedArg0 := tmp11225
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11227 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11221, tmp11226)
}
__typedArg0 := tmp11221
__typedArg1 := tmp11226
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp11227)
}
__typedArg0 := symcons
__typedArg1 := tmp11227
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp11229 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1990)
}
__typedArg0 := V1990
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp11229 {
__e.TailApply(PrimFunc(symshen_4deref_1terms), V1990, V1991, Nil)
return
} else {
__e.Return(V1990)
return
}


}


}, 2)

tmp11251 := Call(__e, ns2_1set, symshen_4function_1calls, tmp11218)


_ = tmp11251

tmp11252 := MakeNative(func(__e *ControlFlow) {
V2000 := __e.Get(1)
_ = V2000
V2001 := __e.Get(2)
_ = V2001
V2002 := __e.Get(3)
_ = V2002
tmp11346 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2000)
}
__typedArg0 := V2000
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11333 Obj

if True == tmp11346 {
tmp11344 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2000)
}
__typedArg0 := V2000
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11345 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(0), tmp11344)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp11344
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11335 Obj

if True == tmp11345 {
tmp11342 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2000)
}
__typedArg0 := V2000
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11343 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11342)
}
__typedArg0 := tmp11342
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11337 Obj

if True == tmp11343 {
tmp11339 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2000)
}
__typedArg0 := V2000
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11340 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11339)
}
__typedArg0 := tmp11339
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11341 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp11340)
}
__typedArg0 := Nil
__typedArg1 := tmp11340
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11338 Obj

if True == tmp11341 {
ifres11338 = True


} else {
ifres11338 = False


}

ifres11337 = ifres11338


} else {
ifres11337 = False


}

var ifres11336 Obj

if True == ifres11337 {
ifres11336 = True


} else {
ifres11336 = False


}

ifres11335 = ifres11336


} else {
ifres11335 = False


}

var ifres11334 Obj

if True == ifres11335 {
ifres11334 = True


} else {
ifres11334 = False


}

ifres11333 = ifres11334


} else {
ifres11333 = False


}

if True == ifres11333 {
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

tmp11261 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(tmp11260)
}
__typedArg0 := tmp11260
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

if True == tmp11261 {
tmp11253 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2000)
}
__typedArg0 := V2000
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11253)
}
__typedArg0 := tmp11253
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return


} else {
tmp11254 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2000)
}
__typedArg0 := V2000
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11255 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11254)
}
__typedArg0 := tmp11254
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11256 := Call(__e, PrimFunc(symshen_4app), tmp11255, MakeString("\n"), symshen_4s)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("attempt to optimise a non-variable "))
__typedS1, __typedOK1 := TypedString(tmp11256)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("attempt to optimise a non-variable ")
__typedArg1 := tmp11256
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("attempt to optimise a non-variable "))
__typedS1, __typedOK1 := TypedString(tmp11256)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("attempt to optimise a non-variable ")
__typedArg1 := tmp11256
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}


} else {
tmp11331 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2000)
}
__typedArg0 := V2000
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11318 Obj

if True == tmp11331 {
tmp11329 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2000)
}
__typedArg0 := V2000
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11330 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(1), tmp11329)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp11329
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11320 Obj

if True == tmp11330 {
tmp11327 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2000)
}
__typedArg0 := V2000
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11328 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11327)
}
__typedArg0 := tmp11327
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11322 Obj

if True == tmp11328 {
tmp11324 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2000)
}
__typedArg0 := V2000
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11325 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11324)
}
__typedArg0 := tmp11324
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11326 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp11325)
}
__typedArg0 := Nil
__typedArg1 := tmp11325
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11323 Obj

if True == tmp11326 {
ifres11323 = True


} else {
ifres11323 = False


}

ifres11322 = ifres11323


} else {
ifres11322 = False


}

var ifres11321 Obj

if True == ifres11322 {
ifres11321 = True


} else {
ifres11321 = False


}

ifres11320 = ifres11321


} else {
ifres11320 = False


}

var ifres11319 Obj

if True == ifres11320 {
ifres11319 = True


} else {
ifres11319 = False


}

ifres11318 = ifres11319


} else {
ifres11318 = False


}

if True == ifres11318 {
tmp11271 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2000)
}
__typedArg0 := V2000
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11272 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11271)
}
__typedArg0 := tmp11271
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11273 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(tmp11272)
}
__typedArg0 := tmp11272
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

if True == tmp11273 {
tmp11262 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2000)
}
__typedArg0 := V2000
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11263 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11262)
}
__typedArg0 := tmp11262
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11264 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V2001, Nil)
}
__typedArg0 := V2001
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11265 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11263, tmp11264)
}
__typedArg0 := tmp11263
__typedArg1 := tmp11264
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4lazyderef, tmp11265)
}
__typedArg0 := symshen_4lazyderef
__typedArg1 := tmp11265
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp11266 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2000)
}
__typedArg0 := V2000
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11267 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11266)
}
__typedArg0 := tmp11266
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11268 := Call(__e, PrimFunc(symshen_4app), tmp11267, MakeString("\n"), symshen_4s)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("attempt to optimise a non-variable "))
__typedS1, __typedOK1 := TypedString(tmp11268)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("attempt to optimise a non-variable ")
__typedArg1 := tmp11268
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("attempt to optimise a non-variable "))
__typedS1, __typedOK1 := TypedString(tmp11268)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("attempt to optimise a non-variable ")
__typedArg1 := tmp11268
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}


} else {
tmp11315 := Call(__e, PrimFunc(symelement_2), V2000, V2002)


tmp11316 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp11315)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp11315
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres11312 Obj

if True == tmp11316 {
tmp11314 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(V2000)
}
__typedArg0 := V2000
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

var ifres11313 Obj

if True == tmp11314 {
ifres11313 = True


} else {
ifres11313 = False


}

ifres11312 = ifres11313


} else {
ifres11312 = False


}

if True == ifres11312 {
tmp11274 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V2001, Nil)
}
__typedArg0 := V2001
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11275 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V2000, tmp11274)
}
__typedArg0 := V2000
__typedArg1 := tmp11274
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4deref, tmp11275)
}
__typedArg0 := symshen_4deref
__typedArg1 := tmp11275
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp11310 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2000)
}
__typedArg0 := V2000
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11291 Obj

if True == tmp11310 {
tmp11308 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2000)
}
__typedArg0 := V2000
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11309 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlambda, tmp11308)
}
__typedArg0 := symlambda
__typedArg1 := tmp11308
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11293 Obj

if True == tmp11309 {
tmp11306 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2000)
}
__typedArg0 := V2000
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11307 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11306)
}
__typedArg0 := tmp11306
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11295 Obj

if True == tmp11307 {
tmp11303 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2000)
}
__typedArg0 := V2000
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11304 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11303)
}
__typedArg0 := tmp11303
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11305 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11304)
}
__typedArg0 := tmp11304
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11297 Obj

if True == tmp11305 {
tmp11299 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2000)
}
__typedArg0 := V2000
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11300 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11299)
}
__typedArg0 := tmp11299
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11301 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11300)
}
__typedArg0 := tmp11300
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11302 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp11301)
}
__typedArg0 := Nil
__typedArg1 := tmp11301
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11298 Obj

if True == tmp11302 {
ifres11298 = True


} else {
ifres11298 = False


}

ifres11297 = ifres11298


} else {
ifres11297 = False


}

var ifres11296 Obj

if True == ifres11297 {
ifres11296 = True


} else {
ifres11296 = False


}

ifres11295 = ifres11296


} else {
ifres11295 = False


}

var ifres11294 Obj

if True == ifres11295 {
ifres11294 = True


} else {
ifres11294 = False


}

ifres11293 = ifres11294


} else {
ifres11293 = False


}

var ifres11292 Obj

if True == ifres11293 {
ifres11292 = True


} else {
ifres11292 = False


}

ifres11291 = ifres11292


} else {
ifres11291 = False


}

if True == ifres11291 {
tmp11276 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2000)
}
__typedArg0 := V2000
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11277 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11276)
}
__typedArg0 := tmp11276
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11278 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2000)
}
__typedArg0 := V2000
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11279 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11278)
}
__typedArg0 := tmp11278
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2000)
}
__typedArg0 := V2000
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11282 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11281)
}
__typedArg0 := tmp11281
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11283 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11282, V2002)
}
__typedArg0 := tmp11282
__typedArg1 := V2002
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11284 := Call(__e, PrimFunc(symshen_4deref_1terms), tmp11280, V2001, tmp11283)


tmp11285 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11284, Nil)
}
__typedArg0 := tmp11284
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11286 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11277, tmp11285)
}
__typedArg0 := tmp11277
__typedArg1 := tmp11285
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp11286)
}
__typedArg0 := symlambda
__typedArg1 := tmp11286
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp11289 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2000)
}
__typedArg0 := V2000
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp11289 {
tmp11287 := MakeNative(func(__e *ControlFlow) {
Z2003 := __e.Get(1)
_ = Z2003
__e.TailApply(PrimFunc(symshen_4deref_1terms), Z2003, V2001, V2002)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp11287, V2000)
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

tmp11347 := Call(__e, ns2_1set, symshen_4deref_1terms, tmp11252)


_ = tmp11347

tmp11348 := MakeNative(func(__e *ControlFlow) {
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
tmp11524 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V2022)
}
__typedArg0 := Nil
__typedArg1 := V2022
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11521 Obj

if True == tmp11524 {
tmp11523 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V2023)
}
__typedArg0 := Nil
__typedArg1 := V2023
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11522 Obj

if True == tmp11523 {
ifres11522 = True


} else {
ifres11522 = False


}

ifres11521 = ifres11522


} else {
ifres11521 = False


}

if True == ifres11521 {
__e.Return(V2025)
return
} else {
tmp11519 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11499 Obj

if True == tmp11519 {
tmp11517 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11518 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11517)
}
__typedArg0 := tmp11517
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11501 Obj

if True == tmp11518 {
tmp11514 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11515 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11514)
}
__typedArg0 := tmp11514
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11516 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4_7m, tmp11515)
}
__typedArg0 := symshen_4_7m
__typedArg1 := tmp11515
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11503 Obj

if True == tmp11516 {
tmp11511 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11512 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11511)
}
__typedArg0 := tmp11511
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11513 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11512)
}
__typedArg0 := tmp11512
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11505 Obj

if True == tmp11513 {
tmp11507 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11508 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11507)
}
__typedArg0 := tmp11507
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11509 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11508)
}
__typedArg0 := tmp11508
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11510 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp11509)
}
__typedArg0 := Nil
__typedArg1 := tmp11509
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11506 Obj

if True == tmp11510 {
ifres11506 = True


} else {
ifres11506 = False


}

ifres11505 = ifres11506


} else {
ifres11505 = False


}

var ifres11504 Obj

if True == ifres11505 {
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

if True == ifres11499 {
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
return PrimCons(symshen_4_7m, tmp11354)
}
__typedArg0 := symshen_4_7m
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
tmp11497 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11477 Obj

if True == tmp11497 {
tmp11495 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11496 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11495)
}
__typedArg0 := tmp11495
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11479 Obj

if True == tmp11496 {
tmp11492 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11493 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11492)
}
__typedArg0 := tmp11492
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

var ifres11481 Obj

if True == tmp11494 {
tmp11489 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11490 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11489)
}
__typedArg0 := tmp11489
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11491 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11490)
}
__typedArg0 := tmp11490
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11483 Obj

if True == tmp11491 {
tmp11485 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11486 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11485)
}
__typedArg0 := tmp11485
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11487 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11486)
}
__typedArg0 := tmp11486
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11488 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp11487)
}
__typedArg0 := Nil
__typedArg1 := tmp11487
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11484 Obj

if True == tmp11488 {
ifres11484 = True


} else {
ifres11484 = False


}

ifres11483 = ifres11484


} else {
ifres11483 = False


}

var ifres11482 Obj

if True == ifres11483 {
ifres11482 = True


} else {
ifres11482 = False


}

ifres11481 = ifres11482


} else {
ifres11481 = False


}

var ifres11480 Obj

if True == ifres11481 {
ifres11480 = True


} else {
ifres11480 = False


}

ifres11479 = ifres11480


} else {
ifres11479 = False


}

var ifres11478 Obj

if True == ifres11479 {
ifres11478 = True


} else {
ifres11478 = False


}

ifres11477 = ifres11478


} else {
ifres11477 = False


}

if True == ifres11477 {
tmp11356 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11357 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11356)
}
__typedArg0 := tmp11356
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11358 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11357)
}
__typedArg0 := tmp11357
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11359 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11360 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V2021, tmp11359)
}
__typedArg0 := V2021
__typedArg1 := tmp11359
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11361 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11358, tmp11360)
}
__typedArg0 := tmp11358
__typedArg1 := tmp11360
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11362 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4_1m, tmp11361)
}
__typedArg0 := symshen_4_1m
__typedArg1 := tmp11361
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

if PrimFunc(symshen_4compile_1head) == __self {
__self1, __self2, __self3, __self4, __self5 = V2021, tmp11362, V2023, V2024, V2025
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4compile_1head), V2021, tmp11362, V2023, V2024, V2025)
return


} else {
tmp11475 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11471 Obj

if True == tmp11475 {
tmp11473 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11474 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4_1m, tmp11473)
}
__typedArg0 := symshen_4_1m
__typedArg1 := tmp11473
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11472 Obj

if True == tmp11474 {
ifres11472 = True


} else {
ifres11472 = False


}

ifres11471 = ifres11472


} else {
ifres11471 = False


}

if True == ifres11471 {
tmp11363 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4compile_1head) == __self {
__self1, __self2, __self3, __self4, __self5 = symshen_4_1m, tmp11363, V2023, V2024, V2025
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4compile_1head), symshen_4_1m, tmp11363, V2023, V2024, V2025)
return


} else {
tmp11469 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11465 Obj

if True == tmp11469 {
tmp11467 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11468 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4_7m, tmp11467)
}
__typedArg0 := symshen_4_7m
__typedArg1 := tmp11467
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11466 Obj

if True == tmp11468 {
ifres11466 = True


} else {
ifres11466 = False


}

ifres11465 = ifres11466


} else {
ifres11465 = False


}

if True == ifres11465 {
tmp11364 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4compile_1head) == __self {
__self1, __self2, __self3, __self4, __self5 = symshen_4_7m, tmp11364, V2023, V2024, V2025
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4compile_1head), symshen_4_7m, tmp11364, V2023, V2024, V2025)
return


} else {
tmp11463 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11456 Obj

if True == tmp11463 {
tmp11462 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2023)
}
__typedArg0 := V2023
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

tmp11461 := Call(__e, PrimFunc(symshen_4wildcard_2), tmp11460)


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

var ifres11457 Obj

if True == ifres11458 {
ifres11457 = True


} else {
ifres11457 = False


}

ifres11456 = ifres11457


} else {
ifres11456 = False


}

if True == ifres11456 {
tmp11365 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11366 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2023)
}
__typedArg0 := V2023
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4compile_1head) == __self {
__self1, __self2, __self3, __self4, __self5 = V2021, tmp11365, tmp11366, V2024, V2025
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4compile_1head), V2021, tmp11365, tmp11366, V2024, V2025)
return


} else {
tmp11454 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11450 Obj

if True == tmp11454 {
tmp11452 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11453 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(tmp11452)
}
__typedArg0 := tmp11452
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

var ifres11451 Obj

if True == tmp11453 {
ifres11451 = True


} else {
ifres11451 = False


}

ifres11450 = ifres11451


} else {
ifres11450 = False


}

if True == ifres11450 {
__e.TailApply(PrimFunc(symshen_4variable_1case), V2021, V2022, V2023, V2024, V2025)
return
} else {
tmp11448 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4_1m, V2021)
}
__typedArg0 := symshen_4_1m
__typedArg1 := V2021
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11441 Obj

if True == tmp11448 {
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

tmp11446 := Call(__e, PrimFunc(symatom_2), tmp11445)


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

var ifres11442 Obj

if True == ifres11443 {
ifres11442 = True


} else {
ifres11442 = False


}

ifres11441 = ifres11442


} else {
ifres11441 = False


}

if True == ifres11441 {
__e.TailApply(PrimFunc(symshen_4atom_1case_1minus), V2022, V2023, V2024, V2025)
return
} else {
tmp11439 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4_1m, V2021)
}
__typedArg0 := symshen_4_1m
__typedArg1 := V2021
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11409 Obj

if True == tmp11439 {
tmp11438 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11411 Obj

if True == tmp11438 {
tmp11436 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11437 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11436)
}
__typedArg0 := tmp11436
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11413 Obj

if True == tmp11437 {
tmp11433 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11434 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11433)
}
__typedArg0 := tmp11433
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11435 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcons, tmp11434)
}
__typedArg0 := symcons
__typedArg1 := tmp11434
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11415 Obj

if True == tmp11435 {
tmp11430 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11431 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11430)
}
__typedArg0 := tmp11430
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11432 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11431)
}
__typedArg0 := tmp11431
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11417 Obj

if True == tmp11432 {
tmp11426 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11427 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11426)
}
__typedArg0 := tmp11426
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11428 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11427)
}
__typedArg0 := tmp11427
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11429 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11428)
}
__typedArg0 := tmp11428
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11419 Obj

if True == tmp11429 {
tmp11421 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11422 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11421)
}
__typedArg0 := tmp11421
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11423 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11422)
}
__typedArg0 := tmp11422
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11424 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11423)
}
__typedArg0 := tmp11423
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11425 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp11424)
}
__typedArg0 := Nil
__typedArg1 := tmp11424
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11420 Obj

if True == tmp11425 {
ifres11420 = True


} else {
ifres11420 = False


}

ifres11419 = ifres11420


} else {
ifres11419 = False


}

var ifres11418 Obj

if True == ifres11419 {
ifres11418 = True


} else {
ifres11418 = False


}

ifres11417 = ifres11418


} else {
ifres11417 = False


}

var ifres11416 Obj

if True == ifres11417 {
ifres11416 = True


} else {
ifres11416 = False


}

ifres11415 = ifres11416


} else {
ifres11415 = False


}

var ifres11414 Obj

if True == ifres11415 {
ifres11414 = True


} else {
ifres11414 = False


}

ifres11413 = ifres11414


} else {
ifres11413 = False


}

var ifres11412 Obj

if True == ifres11413 {
ifres11412 = True


} else {
ifres11412 = False


}

ifres11411 = ifres11412


} else {
ifres11411 = False


}

var ifres11410 Obj

if True == ifres11411 {
ifres11410 = True


} else {
ifres11410 = False


}

ifres11409 = ifres11410


} else {
ifres11409 = False


}

if True == ifres11409 {
__e.TailApply(PrimFunc(symshen_4cons_1case_1minus), V2022, V2023, V2024, V2025)
return
} else {
tmp11407 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4_7m, V2021)
}
__typedArg0 := symshen_4_7m
__typedArg1 := V2021
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11400 Obj

if True == tmp11407 {
tmp11406 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11402 Obj

if True == tmp11406 {
tmp11404 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11405 := Call(__e, PrimFunc(symatom_2), tmp11404)


var ifres11403 Obj

if True == tmp11405 {
ifres11403 = True


} else {
ifres11403 = False


}

ifres11402 = ifres11403


} else {
ifres11402 = False


}

var ifres11401 Obj

if True == ifres11402 {
ifres11401 = True


} else {
ifres11401 = False


}

ifres11400 = ifres11401


} else {
ifres11400 = False


}

if True == ifres11400 {
__e.TailApply(PrimFunc(symshen_4atom_1case_1plus), V2022, V2023, V2024, V2025)
return
} else {
tmp11398 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4_7m, V2021)
}
__typedArg0 := symshen_4_7m
__typedArg1 := V2021
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11368 Obj

if True == tmp11398 {
tmp11397 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11370 Obj

if True == tmp11397 {
tmp11395 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11396 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11395)
}
__typedArg0 := tmp11395
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11372 Obj

if True == tmp11396 {
tmp11392 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11393 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11392)
}
__typedArg0 := tmp11392
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11394 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcons, tmp11393)
}
__typedArg0 := symcons
__typedArg1 := tmp11393
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11374 Obj

if True == tmp11394 {
tmp11389 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11390 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11389)
}
__typedArg0 := tmp11389
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11391 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11390)
}
__typedArg0 := tmp11390
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11376 Obj

if True == tmp11391 {
tmp11385 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11386 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11385)
}
__typedArg0 := tmp11385
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11387 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11386)
}
__typedArg0 := tmp11386
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11388 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11387)
}
__typedArg0 := tmp11387
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11378 Obj

if True == tmp11388 {
tmp11380 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2022)
}
__typedArg0 := V2022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11381 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11380)
}
__typedArg0 := tmp11380
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11382 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11381)
}
__typedArg0 := tmp11381
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11383 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11382)
}
__typedArg0 := tmp11382
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11384 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp11383)
}
__typedArg0 := Nil
__typedArg1 := tmp11383
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11379 Obj

if True == tmp11384 {
ifres11379 = True


} else {
ifres11379 = False


}

ifres11378 = ifres11379


} else {
ifres11378 = False


}

var ifres11377 Obj

if True == ifres11378 {
ifres11377 = True


} else {
ifres11377 = False


}

ifres11376 = ifres11377


} else {
ifres11376 = False


}

var ifres11375 Obj

if True == ifres11376 {
ifres11375 = True


} else {
ifres11375 = False


}

ifres11374 = ifres11375


} else {
ifres11374 = False


}

var ifres11373 Obj

if True == ifres11374 {
ifres11373 = True


} else {
ifres11373 = False


}

ifres11372 = ifres11373


} else {
ifres11372 = False


}

var ifres11371 Obj

if True == ifres11372 {
ifres11371 = True


} else {
ifres11371 = False


}

ifres11370 = ifres11371


} else {
ifres11370 = False


}

var ifres11369 Obj

if True == ifres11370 {
ifres11369 = True


} else {
ifres11369 = False


}

ifres11368 = ifres11369


} else {
ifres11368 = False


}

if True == ifres11368 {
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

tmp11525 := Call(__e, ns2_1set, symshen_4compile_1head, tmp11348)


_ = tmp11525

tmp11526 := MakeNative(func(__e *ControlFlow) {
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
tmp11547 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2037)
}
__typedArg0 := V2037
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11544 Obj

if True == tmp11547 {
tmp11546 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2038)
}
__typedArg0 := V2038
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11545 Obj

if True == tmp11546 {
ifres11545 = True


} else {
ifres11545 = False


}

ifres11544 = ifres11545


} else {
ifres11544 = False


}

if True == ifres11544 {
tmp11541 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2038)
}
__typedArg0 := V2038
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11542 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(tmp11541)
}
__typedArg0 := tmp11541
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

if True == tmp11542 {
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

tmp11529 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2038)
}
__typedArg0 := V2038
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11530 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2037)
}
__typedArg0 := V2037
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11531 := Call(__e, PrimFunc(symsubst), tmp11529, tmp11530, V2040)


__e.TailApply(PrimFunc(symshen_4compile_1head), V2036, tmp11527, tmp11528, V2039, tmp11531)
return


} else {
tmp11532 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2037)
}
__typedArg0 := V2037
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11533 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2038)
}
__typedArg0 := V2038
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11534 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2037)
}
__typedArg0 := V2037
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11535 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2038)
}
__typedArg0 := V2038
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11536 := Call(__e, PrimFunc(symshen_4compile_1head), V2036, tmp11534, tmp11535, V2039, V2040)


tmp11537 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11536, Nil)
}
__typedArg0 := tmp11536
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11538 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11533, tmp11537)
}
__typedArg0 := tmp11533
__typedArg1 := tmp11537
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11539 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11532, tmp11538)
}
__typedArg0 := tmp11532
__typedArg1 := tmp11538
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp11539)
}
__typedArg0 := symlet
__typedArg1 := tmp11539
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

tmp11548 := Call(__e, ns2_1set, symshen_4variable_1case, tmp11526)


_ = tmp11548

tmp11549 := MakeNative(func(__e *ControlFlow) {
V2049 := __e.Get(1)
_ = V2049
V2050 := __e.Get(2)
_ = V2050
V2051 := __e.Get(3)
_ = V2051
V2052 := __e.Get(4)
_ = V2052
tmp11573 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2049)
}
__typedArg0 := V2049
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11570 Obj

if True == tmp11573 {
tmp11572 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2050)
}
__typedArg0 := V2050
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11571 Obj

if True == tmp11572 {
ifres11571 = True


} else {
ifres11571 = False


}

ifres11570 = ifres11571


} else {
ifres11570 = False


}

if True == ifres11570 {
tmp11550 := Call(__e, PrimFunc(symgensym), symTm)


let__10091 := tmp11550
_ = let__10091

tmp11551 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2050)
}
__typedArg0 := V2050
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11552 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V2051, Nil)
}
__typedArg0 := V2051
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11553 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11551, tmp11552)
}
__typedArg0 := tmp11551
__typedArg1 := tmp11552
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11554 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4lazyderef, tmp11553)
}
__typedArg0 := symshen_4lazyderef
__typedArg1 := tmp11553
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11555 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2049)
}
__typedArg0 := V2049
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11556 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11555, Nil)
}
__typedArg0 := tmp11555
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11557 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10091, tmp11556)
}
__typedArg0 := let__10091
__typedArg1 := tmp11556
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11558 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_a, tmp11557)
}
__typedArg0 := sym_a
__typedArg1 := tmp11557
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11559 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2049)
}
__typedArg0 := V2049
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11560 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2050)
}
__typedArg0 := V2050
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11561 := Call(__e, PrimFunc(symshen_4compile_1head), symshen_4_1m, tmp11559, tmp11560, V2051, V2052)


tmp11562 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(False, Nil)
}
__typedArg0 := False
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11563 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11561, tmp11562)
}
__typedArg0 := tmp11561
__typedArg1 := tmp11562
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11564 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11558, tmp11563)
}
__typedArg0 := tmp11558
__typedArg1 := tmp11563
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11565 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp11564)
}
__typedArg0 := symif
__typedArg1 := tmp11564
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11566 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11565, Nil)
}
__typedArg0 := tmp11565
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11567 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11554, tmp11566)
}
__typedArg0 := tmp11554
__typedArg1 := tmp11566
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11568 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10091, tmp11567)
}
__typedArg0 := let__10091
__typedArg1 := tmp11567
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp11568)
}
__typedArg0 := symlet
__typedArg1 := tmp11568
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

tmp11574 := Call(__e, ns2_1set, symshen_4atom_1case_1minus, tmp11549)


_ = tmp11574

tmp11575 := MakeNative(func(__e *ControlFlow) {
V2062 := __e.Get(1)
_ = V2062
V2063 := __e.Get(2)
_ = V2063
V2064 := __e.Get(3)
_ = V2064
V2065 := __e.Get(4)
_ = V2065
tmp11639 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2062)
}
__typedArg0 := V2062
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11609 Obj

if True == tmp11639 {
tmp11637 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2062)
}
__typedArg0 := V2062
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11638 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11637)
}
__typedArg0 := tmp11637
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11611 Obj

if True == tmp11638 {
tmp11634 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2062)
}
__typedArg0 := V2062
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11635 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11634)
}
__typedArg0 := tmp11634
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11636 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcons, tmp11635)
}
__typedArg0 := symcons
__typedArg1 := tmp11635
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11613 Obj

if True == tmp11636 {
tmp11631 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2062)
}
__typedArg0 := V2062
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11632 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11631)
}
__typedArg0 := tmp11631
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11633 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11632)
}
__typedArg0 := tmp11632
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11615 Obj

if True == tmp11633 {
tmp11627 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2062)
}
__typedArg0 := V2062
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11628 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11627)
}
__typedArg0 := tmp11627
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11629 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11628)
}
__typedArg0 := tmp11628
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11630 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11629)
}
__typedArg0 := tmp11629
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11617 Obj

if True == tmp11630 {
tmp11622 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2062)
}
__typedArg0 := V2062
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11623 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11622)
}
__typedArg0 := tmp11622
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11624 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11623)
}
__typedArg0 := tmp11623
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11625 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11624)
}
__typedArg0 := tmp11624
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11626 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp11625)
}
__typedArg0 := Nil
__typedArg1 := tmp11625
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11619 Obj

if True == tmp11626 {
tmp11621 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2063)
}
__typedArg0 := V2063
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11620 Obj

if True == tmp11621 {
ifres11620 = True


} else {
ifres11620 = False


}

ifres11619 = ifres11620


} else {
ifres11619 = False


}

var ifres11618 Obj

if True == ifres11619 {
ifres11618 = True


} else {
ifres11618 = False


}

ifres11617 = ifres11618


} else {
ifres11617 = False


}

var ifres11616 Obj

if True == ifres11617 {
ifres11616 = True


} else {
ifres11616 = False


}

ifres11615 = ifres11616


} else {
ifres11615 = False


}

var ifres11614 Obj

if True == ifres11615 {
ifres11614 = True


} else {
ifres11614 = False


}

ifres11613 = ifres11614


} else {
ifres11613 = False


}

var ifres11612 Obj

if True == ifres11613 {
ifres11612 = True


} else {
ifres11612 = False


}

ifres11611 = ifres11612


} else {
ifres11611 = False


}

var ifres11610 Obj

if True == ifres11611 {
ifres11610 = True


} else {
ifres11610 = False


}

ifres11609 = ifres11610


} else {
ifres11609 = False


}

if True == ifres11609 {
tmp11576 := Call(__e, PrimFunc(symgensym), symTm)


let__10092 := tmp11576
_ = let__10092

tmp11577 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2063)
}
__typedArg0 := V2063
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11578 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V2064, Nil)
}
__typedArg0 := V2064
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11579 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11577, tmp11578)
}
__typedArg0 := tmp11577
__typedArg1 := tmp11578
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11580 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4lazyderef, tmp11579)
}
__typedArg0 := symshen_4lazyderef
__typedArg1 := tmp11579
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11581 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10092, Nil)
}
__typedArg0 := let__10092
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11582 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons_2, tmp11581)
}
__typedArg0 := symcons_2
__typedArg1 := tmp11581
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11583 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2062)
}
__typedArg0 := V2062
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11584 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11583)
}
__typedArg0 := tmp11583
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11585 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11584)
}
__typedArg0 := tmp11584
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11586 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2062)
}
__typedArg0 := V2062
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11587 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11586)
}
__typedArg0 := tmp11586
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11588 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11587)
}
__typedArg0 := tmp11587
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11589 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11588)
}
__typedArg0 := tmp11588
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11590 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2062)
}
__typedArg0 := V2062
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
return PrimCons(tmp11585, tmp11591)
}
__typedArg0 := tmp11585
__typedArg1 := tmp11591
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11593 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10092, Nil)
}
__typedArg0 := let__10092
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11594 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhd, tmp11593)
}
__typedArg0 := symhd
__typedArg1 := tmp11593
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11595 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10092, Nil)
}
__typedArg0 := let__10092
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11596 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtl, tmp11595)
}
__typedArg0 := symtl
__typedArg1 := tmp11595
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11597 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2063)
}
__typedArg0 := V2063
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11598 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11596, tmp11597)
}
__typedArg0 := tmp11596
__typedArg1 := tmp11597
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11599 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11594, tmp11598)
}
__typedArg0 := tmp11594
__typedArg1 := tmp11598
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11600 := Call(__e, PrimFunc(symshen_4compile_1head), symshen_4_1m, tmp11592, tmp11599, V2064, V2065)


tmp11601 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(False, Nil)
}
__typedArg0 := False
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11602 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11600, tmp11601)
}
__typedArg0 := tmp11600
__typedArg1 := tmp11601
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11603 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11582, tmp11602)
}
__typedArg0 := tmp11582
__typedArg1 := tmp11602
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11604 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp11603)
}
__typedArg0 := symif
__typedArg1 := tmp11603
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11605 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11604, Nil)
}
__typedArg0 := tmp11604
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11606 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11580, tmp11605)
}
__typedArg0 := tmp11580
__typedArg1 := tmp11605
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11607 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10092, tmp11606)
}
__typedArg0 := let__10092
__typedArg1 := tmp11606
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp11607)
}
__typedArg0 := symlet
__typedArg1 := tmp11607
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

tmp11640 := Call(__e, ns2_1set, symshen_4cons_1case_1minus, tmp11575)


_ = tmp11640

tmp11641 := MakeNative(func(__e *ControlFlow) {
V2075 := __e.Get(1)
_ = V2075
V2076 := __e.Get(2)
_ = V2076
V2077 := __e.Get(3)
_ = V2077
V2078 := __e.Get(4)
_ = V2078
tmp11685 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2075)
}
__typedArg0 := V2075
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11682 Obj

if True == tmp11685 {
tmp11684 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2076)
}
__typedArg0 := V2076
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11683 Obj

if True == tmp11684 {
ifres11683 = True


} else {
ifres11683 = False


}

ifres11682 = ifres11683


} else {
ifres11682 = False


}

if True == ifres11682 {
tmp11642 := Call(__e, PrimFunc(symgensym), symTm)


let__10093 := tmp11642
_ = let__10093

tmp11643 := Call(__e, PrimFunc(symgensym), symGoTo)


let__10094 := tmp11643
_ = let__10094

tmp11644 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2076)
}
__typedArg0 := V2076
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11645 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V2077, Nil)
}
__typedArg0 := V2077
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11646 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11644, tmp11645)
}
__typedArg0 := tmp11644
__typedArg1 := tmp11645
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11647 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4lazyderef, tmp11646)
}
__typedArg0 := symshen_4lazyderef
__typedArg1 := tmp11646
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11648 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2075)
}
__typedArg0 := V2075
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11649 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2076)
}
__typedArg0 := V2076
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11650 := Call(__e, PrimFunc(symshen_4compile_1head), symshen_4_7m, tmp11648, tmp11649, V2077, V2078)


tmp11651 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11650, Nil)
}
__typedArg0 := tmp11650
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11652 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfreeze, tmp11651)
}
__typedArg0 := symfreeze
__typedArg1 := tmp11651
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11653 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2075)
}
__typedArg0 := V2075
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11654 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11653, Nil)
}
__typedArg0 := tmp11653
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11655 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10093, tmp11654)
}
__typedArg0 := let__10093
__typedArg1 := tmp11654
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11656 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_a, tmp11655)
}
__typedArg0 := sym_a
__typedArg1 := tmp11655
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11657 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10094, Nil)
}
__typedArg0 := let__10094
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11658 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symthaw, tmp11657)
}
__typedArg0 := symthaw
__typedArg1 := tmp11657
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11659 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10093, Nil)
}
__typedArg0 := let__10093
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11660 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4pvar_2, tmp11659)
}
__typedArg0 := symshen_4pvar_2
__typedArg1 := tmp11659
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11661 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2075)
}
__typedArg0 := V2075
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11662 := Call(__e, PrimFunc(symshen_4demode), tmp11661)


tmp11663 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10094, Nil)
}
__typedArg0 := let__10094
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11664 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V2077, tmp11663)
}
__typedArg0 := V2077
__typedArg1 := tmp11663
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11665 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11662, tmp11664)
}
__typedArg0 := tmp11662
__typedArg1 := tmp11664
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11666 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10093, tmp11665)
}
__typedArg0 := let__10093
__typedArg1 := tmp11665
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11667 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4bind_b, tmp11666)
}
__typedArg0 := symshen_4bind_b
__typedArg1 := tmp11666
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11668 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(False, Nil)
}
__typedArg0 := False
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11669 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11667, tmp11668)
}
__typedArg0 := tmp11667
__typedArg1 := tmp11668
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11670 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11660, tmp11669)
}
__typedArg0 := tmp11660
__typedArg1 := tmp11669
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11671 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp11670)
}
__typedArg0 := symif
__typedArg1 := tmp11670
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11672 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11671, Nil)
}
__typedArg0 := tmp11671
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11673 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11658, tmp11672)
}
__typedArg0 := tmp11658
__typedArg1 := tmp11672
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11674 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11656, tmp11673)
}
__typedArg0 := tmp11656
__typedArg1 := tmp11673
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11675 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp11674)
}
__typedArg0 := symif
__typedArg1 := tmp11674
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11676 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11675, Nil)
}
__typedArg0 := tmp11675
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11677 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11652, tmp11676)
}
__typedArg0 := tmp11652
__typedArg1 := tmp11676
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11678 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10094, tmp11677)
}
__typedArg0 := let__10094
__typedArg1 := tmp11677
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11679 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11647, tmp11678)
}
__typedArg0 := tmp11647
__typedArg1 := tmp11678
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11680 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10093, tmp11679)
}
__typedArg0 := let__10093
__typedArg1 := tmp11679
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp11680)
}
__typedArg0 := symlet
__typedArg1 := tmp11680
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

tmp11686 := Call(__e, ns2_1set, symshen_4atom_1case_1plus, tmp11641)


_ = tmp11686

tmp11687 := MakeNative(func(__e *ControlFlow) {
V2089 := __e.Get(1)
_ = V2089
V2090 := __e.Get(2)
_ = V2090
V2091 := __e.Get(3)
_ = V2091
V2092 := __e.Get(4)
_ = V2092
tmp11778 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2089)
}
__typedArg0 := V2089
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11748 Obj

if True == tmp11778 {
tmp11776 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2089)
}
__typedArg0 := V2089
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11777 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11776)
}
__typedArg0 := tmp11776
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11750 Obj

if True == tmp11777 {
tmp11773 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2089)
}
__typedArg0 := V2089
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11774 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11773)
}
__typedArg0 := tmp11773
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11775 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcons, tmp11774)
}
__typedArg0 := symcons
__typedArg1 := tmp11774
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11752 Obj

if True == tmp11775 {
tmp11770 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2089)
}
__typedArg0 := V2089
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11771 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11770)
}
__typedArg0 := tmp11770
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11772 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11771)
}
__typedArg0 := tmp11771
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11754 Obj

if True == tmp11772 {
tmp11766 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2089)
}
__typedArg0 := V2089
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11767 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11766)
}
__typedArg0 := tmp11766
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11768 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11767)
}
__typedArg0 := tmp11767
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11769 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11768)
}
__typedArg0 := tmp11768
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11756 Obj

if True == tmp11769 {
tmp11761 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2089)
}
__typedArg0 := V2089
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11762 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11761)
}
__typedArg0 := tmp11761
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11763 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11762)
}
__typedArg0 := tmp11762
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11764 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11763)
}
__typedArg0 := tmp11763
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11765 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp11764)
}
__typedArg0 := Nil
__typedArg1 := tmp11764
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11758 Obj

if True == tmp11765 {
tmp11760 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2090)
}
__typedArg0 := V2090
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11759 Obj

if True == tmp11760 {
ifres11759 = True


} else {
ifres11759 = False


}

ifres11758 = ifres11759


} else {
ifres11758 = False


}

var ifres11757 Obj

if True == ifres11758 {
ifres11757 = True


} else {
ifres11757 = False


}

ifres11756 = ifres11757


} else {
ifres11756 = False


}

var ifres11755 Obj

if True == ifres11756 {
ifres11755 = True


} else {
ifres11755 = False


}

ifres11754 = ifres11755


} else {
ifres11754 = False


}

var ifres11753 Obj

if True == ifres11754 {
ifres11753 = True


} else {
ifres11753 = False


}

ifres11752 = ifres11753


} else {
ifres11752 = False


}

var ifres11751 Obj

if True == ifres11752 {
ifres11751 = True


} else {
ifres11751 = False


}

ifres11750 = ifres11751


} else {
ifres11750 = False


}

var ifres11749 Obj

if True == ifres11750 {
ifres11749 = True


} else {
ifres11749 = False


}

ifres11748 = ifres11749


} else {
ifres11748 = False


}

if True == ifres11748 {
tmp11688 := Call(__e, PrimFunc(symgensym), symTm)


let__10095 := tmp11688
_ = let__10095

tmp11689 := Call(__e, PrimFunc(symgensym), symGoTo)


let__10096 := tmp11689
_ = let__10096

tmp11690 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2089)
}
__typedArg0 := V2089
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11691 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11690)
}
__typedArg0 := tmp11690
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11692 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11691)
}
__typedArg0 := tmp11691
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11693 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2089)
}
__typedArg0 := V2089
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11694 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11693)
}
__typedArg0 := tmp11693
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11695 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11694)
}
__typedArg0 := tmp11694
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11696 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11695)
}
__typedArg0 := tmp11695
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11697 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11692, tmp11696)
}
__typedArg0 := tmp11692
__typedArg1 := tmp11696
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11698 := Call(__e, PrimFunc(symshen_4extract_1vars), tmp11697)


let__10097 := tmp11698
_ = let__10097

tmp11699 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2089)
}
__typedArg0 := V2089
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11700 := Call(__e, PrimFunc(symshen_4tame), tmp11699)


let__10098 := tmp11700
_ = let__10098

tmp11701 := Call(__e, PrimFunc(symshen_4extract_1vars), let__10098)


let__10099 := tmp11701
_ = let__10099

tmp11702 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2090)
}
__typedArg0 := V2090
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11703 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V2091, Nil)
}
__typedArg0 := V2091
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11704 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11702, tmp11703)
}
__typedArg0 := tmp11702
__typedArg1 := tmp11703
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11705 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4lazyderef, tmp11704)
}
__typedArg0 := symshen_4lazyderef
__typedArg1 := tmp11704
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11706 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2089)
}
__typedArg0 := V2089
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11707 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2090)
}
__typedArg0 := V2090
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11708 := Call(__e, PrimFunc(symshen_4compile_1head), symshen_4_7m, tmp11706, tmp11707, V2091, V2092)


tmp11709 := Call(__e, PrimFunc(symshen_4goto), let__10097, tmp11708)


tmp11710 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10095, Nil)
}
__typedArg0 := let__10095
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11711 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons_2, tmp11710)
}
__typedArg0 := symcons_2
__typedArg1 := tmp11710
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11712 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2089)
}
__typedArg0 := V2089
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11713 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11712)
}
__typedArg0 := tmp11712
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11714 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10095, Nil)
}
__typedArg0 := let__10095
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11715 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhd, tmp11714)
}
__typedArg0 := symhd
__typedArg1 := tmp11714
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11716 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10095, Nil)
}
__typedArg0 := let__10095
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11717 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtl, tmp11716)
}
__typedArg0 := symtl
__typedArg1 := tmp11716
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11718 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11717, Nil)
}
__typedArg0 := tmp11717
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11719 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11715, tmp11718)
}
__typedArg0 := tmp11715
__typedArg1 := tmp11718
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11720 := Call(__e, PrimFunc(symshen_4invoke), let__10096, let__10097)


tmp11721 := Call(__e, PrimFunc(symshen_4compile_1head), symshen_4_7m, tmp11713, tmp11719, V2091, tmp11720)


tmp11722 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10095, Nil)
}
__typedArg0 := let__10095
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11723 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4pvar_2, tmp11722)
}
__typedArg0 := symshen_4pvar_2
__typedArg1 := tmp11722
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11724 := Call(__e, PrimFunc(symshen_4demode), let__10098)


tmp11725 := Call(__e, PrimFunc(symshen_4invoke), let__10096, let__10097)


tmp11726 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11725, Nil)
}
__typedArg0 := tmp11725
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11727 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfreeze, tmp11726)
}
__typedArg0 := symfreeze
__typedArg1 := tmp11726
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11728 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11727, Nil)
}
__typedArg0 := tmp11727
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11729 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V2091, tmp11728)
}
__typedArg0 := V2091
__typedArg1 := tmp11728
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11730 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11724, tmp11729)
}
__typedArg0 := tmp11724
__typedArg1 := tmp11729
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11731 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10095, tmp11730)
}
__typedArg0 := let__10095
__typedArg1 := tmp11730
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11732 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4bind_b, tmp11731)
}
__typedArg0 := symshen_4bind_b
__typedArg1 := tmp11731
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11733 := Call(__e, PrimFunc(symshen_4stpart), let__10099, tmp11732, V2091)


tmp11734 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(False, Nil)
}
__typedArg0 := False
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11735 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11733, tmp11734)
}
__typedArg0 := tmp11733
__typedArg1 := tmp11734
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11736 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11723, tmp11735)
}
__typedArg0 := tmp11723
__typedArg1 := tmp11735
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11737 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp11736)
}
__typedArg0 := symif
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
return PrimCons(tmp11721, tmp11738)
}
__typedArg0 := tmp11721
__typedArg1 := tmp11738
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11740 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11711, tmp11739)
}
__typedArg0 := tmp11711
__typedArg1 := tmp11739
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11741 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp11740)
}
__typedArg0 := symif
__typedArg1 := tmp11740
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11742 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11741, Nil)
}
__typedArg0 := tmp11741
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11743 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11709, tmp11742)
}
__typedArg0 := tmp11709
__typedArg1 := tmp11742
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11744 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10096, tmp11743)
}
__typedArg0 := let__10096
__typedArg1 := tmp11743
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11745 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11705, tmp11744)
}
__typedArg0 := tmp11705
__typedArg1 := tmp11744
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11746 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10095, tmp11745)
}
__typedArg0 := let__10095
__typedArg1 := tmp11745
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp11746)
}
__typedArg0 := symlet
__typedArg1 := tmp11746
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

tmp11779 := Call(__e, ns2_1set, symshen_4cons_1case_1plus, tmp11687)


_ = tmp11779

tmp11780 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__selftop:
V2098 := __self1
_ = V2098
tmp11817 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2098)
}
__typedArg0 := V2098
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11804 Obj

if True == tmp11817 {
tmp11815 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2098)
}
__typedArg0 := V2098
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11816 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4_7m, tmp11815)
}
__typedArg0 := symshen_4_7m
__typedArg1 := tmp11815
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11806 Obj

if True == tmp11816 {
tmp11813 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2098)
}
__typedArg0 := V2098
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11814 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11813)
}
__typedArg0 := tmp11813
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11808 Obj

if True == tmp11814 {
tmp11810 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2098)
}
__typedArg0 := V2098
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11811 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11810)
}
__typedArg0 := tmp11810
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11812 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp11811)
}
__typedArg0 := Nil
__typedArg1 := tmp11811
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11809 Obj

if True == tmp11812 {
ifres11809 = True


} else {
ifres11809 = False


}

ifres11808 = ifres11809


} else {
ifres11808 = False


}

var ifres11807 Obj

if True == ifres11808 {
ifres11807 = True


} else {
ifres11807 = False


}

ifres11806 = ifres11807


} else {
ifres11806 = False


}

var ifres11805 Obj

if True == ifres11806 {
ifres11805 = True


} else {
ifres11805 = False


}

ifres11804 = ifres11805


} else {
ifres11804 = False


}

if True == ifres11804 {
tmp11781 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2098)
}
__typedArg0 := V2098
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11782 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11781)
}
__typedArg0 := tmp11781
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

if PrimFunc(symshen_4demode) == __self {
__self1 = tmp11782
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4demode), tmp11782)
return


} else {
tmp11802 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2098)
}
__typedArg0 := V2098
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11789 Obj

if True == tmp11802 {
tmp11800 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2098)
}
__typedArg0 := V2098
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11801 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4_1m, tmp11800)
}
__typedArg0 := symshen_4_1m
__typedArg1 := tmp11800
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11791 Obj

if True == tmp11801 {
tmp11798 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2098)
}
__typedArg0 := V2098
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11799 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11798)
}
__typedArg0 := tmp11798
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11793 Obj

if True == tmp11799 {
tmp11795 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2098)
}
__typedArg0 := V2098
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11796 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11795)
}
__typedArg0 := tmp11795
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11797 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp11796)
}
__typedArg0 := Nil
__typedArg1 := tmp11796
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11794 Obj

if True == tmp11797 {
ifres11794 = True


} else {
ifres11794 = False


}

ifres11793 = ifres11794


} else {
ifres11793 = False


}

var ifres11792 Obj

if True == ifres11793 {
ifres11792 = True


} else {
ifres11792 = False


}

ifres11791 = ifres11792


} else {
ifres11791 = False


}

var ifres11790 Obj

if True == ifres11791 {
ifres11790 = True


} else {
ifres11790 = False


}

ifres11789 = ifres11790


} else {
ifres11789 = False


}

if True == ifres11789 {
tmp11783 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2098)
}
__typedArg0 := V2098
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11784 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11783)
}
__typedArg0 := tmp11783
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

if PrimFunc(symshen_4demode) == __self {
__self1 = tmp11784
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4demode), tmp11784)
return


} else {
tmp11787 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2098)
}
__typedArg0 := V2098
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp11787 {
tmp11785 := MakeNative(func(__e *ControlFlow) {
Z2099 := __e.Get(1)
_ = Z2099
__e.TailApply(PrimFunc(symshen_4demode), Z2099)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp11785, V2098)
return


} else {
__e.Return(V2098)
return
}


}


}


}, 1)

tmp11818 := Call(__e, ns2_1set, symshen_4demode, tmp11780)


_ = tmp11818

tmp11819 := MakeNative(func(__e *ControlFlow) {
V2100 := __e.Get(1)
_ = V2100
tmp11824 := Call(__e, PrimFunc(symshen_4wildcard_2), V2100)


if True == tmp11824 {
__e.TailApply(PrimFunc(symgensym), symY)
return
} else {
tmp11822 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2100)
}
__typedArg0 := V2100
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp11822 {
tmp11820 := MakeNative(func(__e *ControlFlow) {
Z2101 := __e.Get(1)
_ = Z2101
__e.TailApply(PrimFunc(symshen_4tame), Z2101)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp11820, V2100)
return


} else {
__e.Return(V2100)
return
}


}


}, 1)

tmp11825 := Call(__e, ns2_1set, symshen_4tame, tmp11819)


_ = tmp11825

tmp11826 := MakeNative(func(__e *ControlFlow) {
V2102 := __e.Get(1)
_ = V2102
V2103 := __e.Get(2)
_ = V2103
tmp11829 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V2102)
}
__typedArg0 := Nil
__typedArg1 := V2102
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp11829 {
tmp11827 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V2103, Nil)
}
__typedArg0 := V2103
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfreeze, tmp11827)
}
__typedArg0 := symfreeze
__typedArg1 := tmp11827
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.TailApply(PrimFunc(symshen_4goto_1h), V2102, V2103)
return
}


}, 2)

tmp11830 := Call(__e, ns2_1set, symshen_4goto, tmp11826)


_ = tmp11830

tmp11831 := MakeNative(func(__e *ControlFlow) {
V2104 := __e.Get(1)
_ = V2104
V2105 := __e.Get(2)
_ = V2105
tmp11840 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V2104)
}
__typedArg0 := Nil
__typedArg1 := V2104
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp11840 {
__e.Return(V2105)
return
} else {
tmp11838 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2104)
}
__typedArg0 := V2104
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp11838 {
tmp11832 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2104)
}
__typedArg0 := V2104
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11833 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2104)
}
__typedArg0 := V2104
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11834 := Call(__e, PrimFunc(symshen_4goto_1h), tmp11833, V2105)


tmp11835 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11834, Nil)
}
__typedArg0 := tmp11834
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11836 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11832, tmp11835)
}
__typedArg0 := tmp11832
__typedArg1 := tmp11835
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp11836)
}
__typedArg0 := symlambda
__typedArg1 := tmp11836
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4goto_1h)
return
}


}


}, 2)

tmp11841 := Call(__e, ns2_1set, symshen_4goto_1h, tmp11831)


_ = tmp11841

tmp11842 := MakeNative(func(__e *ControlFlow) {
V2106 := __e.Get(1)
_ = V2106
V2107 := __e.Get(2)
_ = V2107
tmp11845 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V2107)
}
__typedArg0 := Nil
__typedArg1 := V2107
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp11845 {
tmp11843 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V2106, Nil)
}
__typedArg0 := V2106
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symthaw, tmp11843)
}
__typedArg0 := symthaw
__typedArg1 := tmp11843
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

tmp11846 := Call(__e, ns2_1set, symshen_4invoke, tmp11842)


_ = tmp11846

tmp11847 := MakeNative(func(__e *ControlFlow) {
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

tmp11848 := Call(__e, ns2_1set, symshen_4wildcard_2, tmp11847)


_ = tmp11848

tmp11849 := MakeNative(func(__e *ControlFlow) {
V2109 := __e.Get(1)
_ = V2109
tmp11856 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symabsvector_2) {
return PrimIsVector(V2109)
}
__typedArg0 := V2109
return Call(__e, PrimFunc(symabsvector_2), __typedArg0)
})()

if True == tmp11856 {
tmp11851 := MakeNative(func(__e *ControlFlow) {
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

tmp11852 := MakeNative(func(__e *ControlFlow) {
Z2110 := __e.Get(1)
_ = Z2110
__e.Return(symshen_4not_1pvar)
return
}, 1)

tmp11853 := Call(__e, try_1catch, tmp11851, tmp11852)


tmp11854 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp11853, symshen_4pvar)
}
__typedArg0 := tmp11853
__typedArg1 := symshen_4pvar
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp11854 {
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

tmp11857 := Call(__e, ns2_1set, symshen_4pvar_2, tmp11849)


_ = tmp11857

tmp11858 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__selftop:
V2111 := __self1
_ = V2111
V2112 := __self2
_ = V2112
tmp11864 := Call(__e, PrimFunc(symshen_4pvar_2), V2111)


if True == tmp11864 {
tmp11859 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V2111, MakeInteger(1))
}
__typedArg0 := V2111
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})()

tmp11860 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V2112, tmp11859)
}
__typedArg0 := V2112
__typedArg1 := tmp11859
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})()

let__10100 := tmp11860
_ = let__10100

tmp11862 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__10100, symshen_4_1null_1)
}
__typedArg0 := let__10100
__typedArg1 := symshen_4_1null_1
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp11862 {
__e.Return(V2111)
return
} else {
if PrimFunc(symshen_4lazyderef) == __self {
__self1, __self2 = let__10100, V2112
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4lazyderef), let__10100, V2112)
return
}


} else {
__e.Return(V2111)
return
}


}, 2)

tmp11865 := Call(__e, ns2_1set, symshen_4lazyderef, tmp11858)


_ = tmp11865

tmp11866 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__selftop:
V2114 := __self1
_ = V2114
V2115 := __self2
_ = V2115
tmp11878 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2114)
}
__typedArg0 := V2114
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp11878 {
tmp11867 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2114)
}
__typedArg0 := V2114
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11868 := Call(__e, PrimFunc(symshen_4deref), tmp11867, V2115)


tmp11869 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2114)
}
__typedArg0 := V2114
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11870 := Call(__e, PrimFunc(symshen_4deref), tmp11869, V2115)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11868, tmp11870)
}
__typedArg0 := tmp11868
__typedArg1 := tmp11870
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp11876 := Call(__e, PrimFunc(symshen_4pvar_2), V2114)


if True == tmp11876 {
tmp11871 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V2114, MakeInteger(1))
}
__typedArg0 := V2114
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})()

tmp11872 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V2115, tmp11871)
}
__typedArg0 := V2115
__typedArg1 := tmp11871
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})()

let__10101 := tmp11872
_ = let__10101

tmp11874 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__10101, symshen_4_1null_1)
}
__typedArg0 := let__10101
__typedArg1 := symshen_4_1null_1
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp11874 {
__e.Return(V2114)
return
} else {
if PrimFunc(symshen_4deref) == __self {
__self1, __self2 = let__10101, V2115
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4deref), let__10101, V2115)
return
}


} else {
__e.Return(V2114)
return
}


}


}, 2)

tmp11879 := Call(__e, ns2_1set, symshen_4deref, tmp11866)


_ = tmp11879

tmp11880 := MakeNative(func(__e *ControlFlow) {
V2117 := __e.Get(1)
_ = V2117
V2118 := __e.Get(2)
_ = V2118
V2119 := __e.Get(3)
_ = V2119
V2120 := __e.Get(4)
_ = V2120
tmp11881 := Call(__e, PrimFunc(symshen_4bindv), V2117, V2118, V2119)


let__10102 := tmp11881
_ = let__10102

tmp11882 := Call(__e, PrimFunc(symthaw), V2120)


let__10103 := tmp11882
_ = let__10103

tmp11884 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__10103, False)
}
__typedArg0 := let__10103
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp11884 {
__e.TailApply(PrimFunc(symshen_4unwind), V2117, V2119, let__10103)
return
} else {
__e.Return(let__10103)
return
}


}, 4)

tmp11885 := Call(__e, ns2_1set, symshen_4bind_b, tmp11880)


_ = tmp11885

tmp11886 := MakeNative(func(__e *ControlFlow) {
V2123 := __e.Get(1)
_ = V2123
V2124 := __e.Get(2)
_ = V2124
V2125 := __e.Get(3)
_ = V2125
tmp11887 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V2123, MakeInteger(1))
}
__typedArg0 := V2123
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(V2125, tmp11887, V2124)
}
__typedArg0 := V2125
__typedArg1 := tmp11887
__typedArg2 := V2124
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})())
return


}, 3)

tmp11888 := Call(__e, ns2_1set, symshen_4bindv, tmp11886)


_ = tmp11888

tmp11889 := MakeNative(func(__e *ControlFlow) {
V2126 := __e.Get(1)
_ = V2126
V2127 := __e.Get(2)
_ = V2127
V2128 := __e.Get(3)
_ = V2128
tmp11890 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V2126, MakeInteger(1))
}
__typedArg0 := V2126
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})()

tmp11891 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(V2127, tmp11890, symshen_4_1null_1)
}
__typedArg0 := V2127
__typedArg1 := tmp11890
__typedArg2 := symshen_4_1null_1
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

_ = tmp11891

__e.Return(V2128)
return


}, 3)

tmp11892 := Call(__e, ns2_1set, symshen_4unwind, tmp11889)


_ = tmp11892

tmp11893 := MakeNative(func(__e *ControlFlow) {
V2137 := __e.Get(1)
_ = V2137
V2138 := __e.Get(2)
_ = V2138
V2139 := __e.Get(3)
_ = V2139
tmp11908 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V2137)
}
__typedArg0 := Nil
__typedArg1 := V2137
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp11908 {
__e.Return(V2138)
return
} else {
tmp11906 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2137)
}
__typedArg0 := V2137
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp11906 {
tmp11894 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2137)
}
__typedArg0 := V2137
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11895 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V2139, Nil)
}
__typedArg0 := V2139
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11896 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4newpv, tmp11895)
}
__typedArg0 := symshen_4newpv
__typedArg1 := tmp11895
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11897 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2137)
}
__typedArg0 := V2137
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11898 := Call(__e, PrimFunc(symshen_4stpart), tmp11897, V2138, V2139)


tmp11899 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11898, Nil)
}
__typedArg0 := tmp11898
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11900 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V2139, tmp11899)
}
__typedArg0 := V2139
__typedArg1 := tmp11899
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11901 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4gc, tmp11900)
}
__typedArg0 := symshen_4gc
__typedArg1 := tmp11900
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11902 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11901, Nil)
}
__typedArg0 := tmp11901
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11903 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11896, tmp11902)
}
__typedArg0 := tmp11896
__typedArg1 := tmp11902
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11904 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11894, tmp11903)
}
__typedArg0 := tmp11894
__typedArg1 := tmp11903
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp11904)
}
__typedArg0 := symlet
__typedArg1 := tmp11904
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

tmp11909 := Call(__e, ns2_1set, symshen_4stpart, tmp11893)


_ = tmp11909

tmp11910 := MakeNative(func(__e *ControlFlow) {
V2140 := __e.Get(1)
_ = V2140
V2141 := __e.Get(2)
_ = V2141
tmp11914 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V2141, False)
}
__typedArg0 := V2141
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp11914 {
tmp11911 := Call(__e, PrimFunc(symshen_4ticket_1number), V2140)


let__10104 := tmp11911
_ = let__10104

tmp11912 := Call(__e, PrimFunc(symshen_4decrement_1ticket), let__10104, V2140)


_ = tmp11912

__e.Return(V2141)
return


} else {
__e.Return(V2141)
return
}


}, 2)

tmp11915 := Call(__e, ns2_1set, symshen_4gc, tmp11910)


_ = tmp11915

tmp11916 := MakeNative(func(__e *ControlFlow) {
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

tmp11918 := Call(__e, ns2_1set, symshen_4decrement_1ticket, tmp11916)


_ = tmp11918

tmp11919 := MakeNative(func(__e *ControlFlow) {
V2145 := __e.Get(1)
_ = V2145
tmp11920 := Call(__e, PrimFunc(symshen_4ticket_1number), V2145)


let__10105 := tmp11920
_ = let__10105

tmp11921 := Call(__e, PrimFunc(symshen_4make_1prolog_1variable), let__10105)


let__10106 := tmp11921
_ = let__10106

tmp11922 := Call(__e, PrimFunc(symshen_4nextticket), V2145, let__10105)


let__10107 := tmp11922
_ = let__10107

__e.Return(let__10106)
return


}, 1)

tmp11923 := Call(__e, ns2_1set, symshen_4newpv, tmp11919)


_ = tmp11923

tmp11924 := MakeNative(func(__e *ControlFlow) {
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

tmp11925 := Call(__e, ns2_1set, symshen_4ticket_1number, tmp11924)


_ = tmp11925

tmp11926 := MakeNative(func(__e *ControlFlow) {
V2150 := __e.Get(1)
_ = V2150
V2151 := __e.Get(2)
_ = V2151
tmp11927 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(V2150, V2151, symshen_4_1null_1)
}
__typedArg0 := V2150
__typedArg1 := V2151
__typedArg2 := symshen_4_1null_1
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

let__10108 := tmp11927
_ = let__10108

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(let__10108, MakeInteger(1), (func() Obj {
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
__typedArg0 := let__10108
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

tmp11929 := Call(__e, ns2_1set, symshen_4nextticket, tmp11926)


_ = tmp11929

tmp11930 := MakeNative(func(__e *ControlFlow) {
V2153 := __e.Get(1)
_ = V2153
tmp11931 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symabsvector) {
return PrimAbsvector(MakeInteger(2))
}
__typedArg0 := MakeInteger(2)
return Call(__e, PrimFunc(symabsvector), __typedArg0)
})()

tmp11932 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(tmp11931, MakeInteger(0), symshen_4pvar)
}
__typedArg0 := tmp11931
__typedArg1 := MakeInteger(0)
__typedArg2 := symshen_4pvar
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(tmp11932, MakeInteger(1), V2153)
}
__typedArg0 := tmp11932
__typedArg1 := MakeInteger(1)
__typedArg2 := V2153
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})())
return


}, 1)

tmp11933 := Call(__e, ns2_1set, symshen_4make_1prolog_1variable, tmp11930)


_ = tmp11933

tmp11934 := MakeNative(func(__e *ControlFlow) {
V2154 := __e.Get(1)
_ = V2154
tmp11935 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V2154, MakeInteger(1))
}
__typedArg0 := V2154
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})()

tmp11936 := Call(__e, PrimFunc(symshen_4app), tmp11935, MakeString(""), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("Var"))
__typedS1, __typedOK1 := TypedString(tmp11936)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("Var")
__typedArg1 := tmp11936
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp11937 := Call(__e, ns2_1set, symshen_4pvar, tmp11934)


_ = tmp11937

tmp11938 := MakeNative(func(__e *ControlFlow) {
tmp11939 := (func() Obj {
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
__typedN1, __typedOK1 := TypedFloat64(tmp11939)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp11939
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())
}
__typedArg0 := symshen_4_dinfs_d
__typedArg1 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(1))
__typedN1, __typedOK1 := TypedFloat64(tmp11939)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp11939
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return


}, 0)

tmp11941 := Call(__e, ns2_1set, symshen_4incinfs, tmp11938)


_ = tmp11941

tmp11942 := MakeNative(func(__e *ControlFlow) {
V2155 := __e.Get(1)
_ = V2155
tmp11949 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(syminteger_2) {
return PrimIsInteger(V2155)
}
__typedArg0 := V2155
return Call(__e, PrimFunc(syminteger_2), __typedArg0)
})()

var ifres11946 Obj

if True == tmp11949 {
tmp11948 := (func() Obj {
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

var ifres11947 Obj

if True == tmp11948 {
ifres11947 = True


} else {
ifres11947 = False


}

ifres11946 = ifres11947


} else {
ifres11946 = False


}

if True == ifres11946 {
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
tmp11943 := Call(__e, PrimFunc(symshen_4app), V2155, MakeString(""), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("prolog vector size: size should be a positive integer; not "))
__typedS1, __typedOK1 := TypedString(tmp11943)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("prolog vector size: size should be a positive integer; not ")
__typedArg1 := tmp11943
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("prolog vector size: size should be a positive integer; not "))
__typedS1, __typedOK1 := TypedString(tmp11943)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("prolog vector size: size should be a positive integer; not ")
__typedArg1 := tmp11943
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}


}, 1)

tmp11950 := Call(__e, ns2_1set, symshen_4prolog_1vector_1size, tmp11942)


_ = tmp11950

tmp11951 := MakeNative(func(__e *ControlFlow) {
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
tmp11981 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V2167, V2168)
}
__typedArg0 := V2167
__typedArg1 := V2168
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp11981 {
__e.TailApply(PrimFunc(symthaw), V2170)
return
} else {
tmp11979 := Call(__e, PrimFunc(symshen_4pvar_2), V2167)


var ifres11974 Obj

if True == tmp11979 {
tmp11976 := Call(__e, PrimFunc(symshen_4deref), V2168, V2169)


tmp11977 := Call(__e, PrimFunc(symshen_4occurs_1check_2), V2167, tmp11976)


tmp11978 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp11977)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp11977
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres11975 Obj

if True == tmp11978 {
ifres11975 = True


} else {
ifres11975 = False


}

ifres11974 = ifres11975


} else {
ifres11974 = False


}

if True == ifres11974 {
__e.TailApply(PrimFunc(symshen_4bind_b), V2167, V2168, V2169, V2170)
return
} else {
tmp11972 := Call(__e, PrimFunc(symshen_4pvar_2), V2168)


var ifres11967 Obj

if True == tmp11972 {
tmp11969 := Call(__e, PrimFunc(symshen_4deref), V2167, V2169)


tmp11970 := Call(__e, PrimFunc(symshen_4occurs_1check_2), V2168, tmp11969)


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
__e.TailApply(PrimFunc(symshen_4bind_b), V2168, V2167, V2169, V2170)
return
} else {
tmp11965 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2167)
}
__typedArg0 := V2167
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11962 Obj

if True == tmp11965 {
tmp11964 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2168)
}
__typedArg0 := V2168
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11963 Obj

if True == tmp11964 {
ifres11963 = True


} else {
ifres11963 = False


}

ifres11962 = ifres11963


} else {
ifres11962 = False


}

if True == ifres11962 {
tmp11952 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2167)
}
__typedArg0 := V2167
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11953 := Call(__e, PrimFunc(symshen_4lazyderef), tmp11952, V2169)


tmp11954 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2168)
}
__typedArg0 := V2168
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11955 := Call(__e, PrimFunc(symshen_4lazyderef), tmp11954, V2169)


tmp11956 := MakeNative(func(__e *ControlFlow) {
tmp11957 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2167)
}
__typedArg0 := V2167
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11958 := Call(__e, PrimFunc(symshen_4lazyderef), tmp11957, V2169)


tmp11959 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2168)
}
__typedArg0 := V2168
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11960 := Call(__e, PrimFunc(symshen_4lazyderef), tmp11959, V2169)


__e.TailApply(PrimFunc(symshen_4lzy_a_b), tmp11958, tmp11960, V2169, V2170)
return


}, 0)

if PrimFunc(symshen_4lzy_a_b) == __self {
__self1, __self2, __self3, __self4 = tmp11953, tmp11955, V2169, tmp11956
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4lzy_a_b), tmp11953, tmp11955, V2169, tmp11956)
return


} else {
__e.Return(False)
return
}


}


}


}


}, 4)

tmp11982 := Call(__e, ns2_1set, symshen_4lzy_a_b, tmp11951)


_ = tmp11982

tmp11983 := MakeNative(func(__e *ControlFlow) {
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
tmp12003 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V2182, V2183)
}
__typedArg0 := V2182
__typedArg1 := V2183
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp12003 {
__e.TailApply(PrimFunc(symthaw), V2185)
return
} else {
tmp12001 := Call(__e, PrimFunc(symshen_4pvar_2), V2182)


if True == tmp12001 {
__e.TailApply(PrimFunc(symshen_4bind_b), V2182, V2183, V2184, V2185)
return
} else {
tmp11999 := Call(__e, PrimFunc(symshen_4pvar_2), V2183)


if True == tmp11999 {
__e.TailApply(PrimFunc(symshen_4bind_b), V2183, V2182, V2184, V2185)
return
} else {
tmp11997 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2182)
}
__typedArg0 := V2182
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11994 Obj

if True == tmp11997 {
tmp11996 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2183)
}
__typedArg0 := V2183
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11995 Obj

if True == tmp11996 {
ifres11995 = True


} else {
ifres11995 = False


}

ifres11994 = ifres11995


} else {
ifres11994 = False


}

if True == ifres11994 {
tmp11984 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2182)
}
__typedArg0 := V2182
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11985 := Call(__e, PrimFunc(symshen_4lazyderef), tmp11984, V2184)


tmp11986 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2183)
}
__typedArg0 := V2183
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11987 := Call(__e, PrimFunc(symshen_4lazyderef), tmp11986, V2184)


tmp11988 := MakeNative(func(__e *ControlFlow) {
tmp11989 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2182)
}
__typedArg0 := V2182
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11990 := Call(__e, PrimFunc(symshen_4lazyderef), tmp11989, V2184)


tmp11991 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2183)
}
__typedArg0 := V2183
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11992 := Call(__e, PrimFunc(symshen_4lazyderef), tmp11991, V2184)


__e.TailApply(PrimFunc(symshen_4lzy_a), tmp11990, tmp11992, V2184, V2185)
return


}, 0)

if PrimFunc(symshen_4lzy_a) == __self {
__self1, __self2, __self3, __self4 = tmp11985, tmp11987, V2184, tmp11988
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4lzy_a), tmp11985, tmp11987, V2184, tmp11988)
return


} else {
__e.Return(False)
return
}


}


}


}


}, 4)

tmp12004 := Call(__e, ns2_1set, symshen_4lzy_a, tmp11983)


_ = tmp12004

tmp12005 := MakeNative(func(__e *ControlFlow) {
V2191 := __e.Get(1)
_ = V2191
V2192 := __e.Get(2)
_ = V2192
tmp12015 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V2191, V2192)
}
__typedArg0 := V2191
__typedArg1 := V2192
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp12015 {
__e.Return(True)
return
} else {
tmp12013 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2192)
}
__typedArg0 := V2192
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp12013 {
tmp12010 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2192)
}
__typedArg0 := V2192
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12011 := Call(__e, PrimFunc(symshen_4occurs_1check_2), V2191, tmp12010)


if True == tmp12011 {
__e.Return(True)
return
} else {
tmp12007 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2192)
}
__typedArg0 := V2192
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12008 := Call(__e, PrimFunc(symshen_4occurs_1check_2), V2191, tmp12007)


if True == tmp12008 {
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

tmp12016 := Call(__e, ns2_1set, symshen_4occurs_1check_2, tmp12005)


_ = tmp12016

tmp12017 := MakeNative(func(__e *ControlFlow) {
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
tmp12018 := Call(__e, V2193, V2194)


tmp12019 := Call(__e, tmp12018, V2195)


tmp12020 := Call(__e, tmp12019, V2196)


__e.TailApply(tmp12020, V2197)
return


}, 5)

tmp12021 := Call(__e, ns2_1set, symcall, tmp12017)


_ = tmp12021

tmp12022 := MakeNative(func(__e *ControlFlow) {
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

tmp12023 := Call(__e, ns2_1set, symreturn, tmp12022)


_ = tmp12023

tmp12024 := MakeNative(func(__e *ControlFlow) {
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

tmp12026 := Call(__e, ns2_1set, symwhen, tmp12024)


_ = tmp12026

tmp12027 := MakeNative(func(__e *ControlFlow) {
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
tmp12028 := Call(__e, PrimFunc(symshen_4lazyderef), V2220, V2222)


tmp12029 := Call(__e, PrimFunc(symshen_4lazyderef), V2221, V2222)


__e.TailApply(PrimFunc(symshen_4lzy_a), tmp12028, tmp12029, V2222, V2225)
return


}, 6)

tmp12030 := Call(__e, ns2_1set, symis, tmp12027)


_ = tmp12030

tmp12031 := MakeNative(func(__e *ControlFlow) {
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
tmp12032 := Call(__e, PrimFunc(symshen_4lazyderef), V2226, V2228)


tmp12033 := Call(__e, PrimFunc(symshen_4lazyderef), V2227, V2228)


__e.TailApply(PrimFunc(symshen_4lzy_a_b), tmp12032, tmp12033, V2228, V2231)
return


}, 6)

tmp12034 := Call(__e, ns2_1set, symis_b, tmp12031)


_ = tmp12034

tmp12035 := MakeNative(func(__e *ControlFlow) {
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

tmp12036 := Call(__e, ns2_1set, symbind, tmp12035)


_ = tmp12036

tmp12037 := MakeNative(func(__e *ControlFlow) {
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
tmp12039 := Call(__e, PrimFunc(symshen_4lazyderef), V2242, V2243)


tmp12040 := Call(__e, PrimFunc(symshen_4pvar_2), tmp12039)


if True == tmp12040 {
__e.TailApply(PrimFunc(symthaw), V2246)
return
} else {
__e.Return(False)
return
}


}, 5)

tmp12041 := Call(__e, ns2_1set, symvar_2, tmp12037)


_ = tmp12041

tmp12042 := MakeNative(func(__e *ControlFlow) {
V2249 := __e.Get(1)
_ = V2249
__e.Return(MakeString("|prolog vector|"))
return
}, 1)

tmp12043 := Call(__e, ns2_1set, symshen_4print_1prolog_1vector, tmp12042)


_ = tmp12043

tmp12044 := MakeNative(func(__e *ControlFlow) {
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
tmp12056 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V2268)
}
__typedArg0 := Nil
__typedArg1 := V2268
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp12056 {
__e.Return(False)
return
} else {
tmp12054 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2268)
}
__typedArg0 := V2268
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp12054 {
tmp12045 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2268)
}
__typedArg0 := V2268
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12046 := Call(__e, tmp12045, V2269)


tmp12047 := Call(__e, tmp12046, V2270)


tmp12048 := Call(__e, tmp12047, V2271)


tmp12049 := Call(__e, tmp12048, V2272)


let__10109 := tmp12049
_ = let__10109

tmp12052 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__10109, False)
}
__typedArg0 := let__10109
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp12052 {
tmp12050 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2268)
}
__typedArg0 := V2268
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symfork) == __self {
__self1, __self2, __self3, __self4, __self5 = tmp12050, V2269, V2270, V2271, V2272
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symfork), tmp12050, V2269, V2270, V2271, V2272)
return


} else {
__e.Return(let__10109)
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

tmp12057 := Call(__e, ns2_1set, symfork, tmp12044)


_ = tmp12057

tmp12058 := MakeNative(func(__e *ControlFlow) {
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
tmp12064 := Call(__e, PrimFunc(symshen_4unlocked_2), V2278)


if True == tmp12064 {
tmp12059 := Call(__e, PrimFunc(symshen_4newpv), V2277)


let__10110 := tmp12059
_ = let__10110

tmp12060 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp12060

tmp12061 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4findall_1h), V2274, V2275, V2276, let__10110, V2277, V2278, V2279, V2280)
return
}, 0)

tmp12062 := Call(__e, PrimFunc(symis), let__10110, Nil, V2277, V2278, V2279, tmp12061)


__e.TailApply(PrimFunc(symshen_4gc), V2277, tmp12062)
return


} else {
__e.Return(False)
return
}


}, 7)

tmp12065 := Call(__e, ns2_1set, symfindall, tmp12058)


_ = tmp12065

tmp12066 := MakeNative(func(__e *ControlFlow) {
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
tmp12071 := Call(__e, PrimFunc(symshen_4unlocked_2), V2287)


var ifres12067 Obj

if True == tmp12071 {
tmp12068 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp12068

tmp12069 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4overbind), V2282, V2285, V2286, V2287, V2288, V2289)
return
}, 0)

tmp12070 := Call(__e, PrimFunc(symcall), V2283, V2286, V2287, V2288, tmp12069)


ifres12067 = tmp12070


} else {
ifres12067 = False


}

let__10111 := ifres12067
_ = let__10111

tmp12076 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__10111, False)
}
__typedArg0 := let__10111
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp12076 {
tmp12074 := Call(__e, PrimFunc(symshen_4unlocked_2), V2287)


if True == tmp12074 {
tmp12072 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp12072

__e.TailApply(PrimFunc(symis_b), V2284, V2285, V2286, V2287, V2288, V2289)
return


} else {
__e.Return(False)
return
}


} else {
__e.Return(let__10111)
return
}


}, 8)

tmp12077 := Call(__e, ns2_1set, symshen_4findall_1h, tmp12066)


_ = tmp12077

tmp12078 := MakeNative(func(__e *ControlFlow) {
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
tmp12079 := Call(__e, PrimFunc(symshen_4deref), V2297, V2299)


tmp12080 := Call(__e, PrimFunc(symshen_4lazyderef), V2298, V2299)


tmp12081 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12079, tmp12080)
}
__typedArg0 := tmp12079
__typedArg1 := tmp12080
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12082 := Call(__e, PrimFunc(symshen_4bindv), V2298, tmp12081, V2299)


_ = tmp12082

__e.Return(False)
return


}, 6)

tmp12083 := Call(__e, ns2_1set, symshen_4overbind, tmp12078)


_ = tmp12083

tmp12084 := MakeNative(func(__e *ControlFlow) {
V2305 := __e.Get(1)
_ = V2305
tmp12088 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_7, V2305)
}
__typedArg0 := sym_7
__typedArg1 := V2305
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp12088 {
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
tmp12086 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_1, V2305)
}
__typedArg0 := sym_1
__typedArg1 := V2305
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp12086 {
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

__e.TailApply(ns2_1set, symoccurs_1check, tmp12084)
return




}, 0)

