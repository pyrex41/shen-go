package main

import . "github.com/pyrex41/shen-go/kl"

var PrologMain = MakeNative0(func(__e *ControlFlow) Obj {
tmp10112 := MakeNative1(func(__e *ControlFlow, V1599 Obj) Obj {
return TailCall2(__e, SymFn(symshen_4assert_d), V1599, symshen_4top)
panic("unreachable")
})

tmp10113 := Call2(__e, ns2_1set, symasserta, tmp10112)


_ = tmp10113

tmp10114 := MakeNative1(func(__e *ControlFlow, V1600 Obj) Obj {
return TailCall2(__e, SymFn(symshen_4assert_d), V1600, symshen_4bottom)
panic("unreachable")
})

tmp10115 := Call2(__e, ns2_1set, symassertz, tmp10114)


_ = tmp10115

tmp10116 := MakeNative2(func(__e *ControlFlow, V1601, V1602 Obj) Obj {
tmp10143 := GConsP(__e, V1601)

var ifres10134 Obj

if True == tmp10143 {
tmp10141 := GTl(__e, V1601)

tmp10142 := GConsP(__e, tmp10141)

var ifres10136 Obj

if True == tmp10142 {
tmp10138 := GTl(__e, V1601)

tmp10139 := GHd(__e, tmp10138)

tmp10140 := GEq(__e, sym_5_1_1, tmp10139)

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
tmp10117 := GHd(__e, V1601)

tmp10118 := Call1(__e, SymFn(symshen_4predicate), tmp10117)


let__9894 := tmp10118
_ = let__9894

tmp10119 := GHd(__e, V1601)

tmp10120 := Call1(__e, SymFn(symshen_4terms), tmp10119)


let__9895 := tmp10120
_ = let__9895

tmp10121 := Call1(__e, SymFn(symlength), let__9895)


let__9896 := tmp10121
_ = let__9896

tmp10122 := Call1(__e, SymFn(symshen_4parameters), let__9896)


let__9897 := tmp10122
_ = let__9897

tmp10123 := Call1(__e, SymFn(symarity), let__9894)


let__9898 := tmp10123
_ = let__9898

tmp10129 := GEq(__e, let__9898, MakeInteger(-1))

var ifres10124 Obj

if True == tmp10129 {
tmp10125 := Call2(__e, SymFn(symshen_4create_1skeleton), let__9894, let__9897)


tmp10126 := Call1(__e, SymFn(symeval), tmp10125)


_ = tmp10126

tmp10127 := GValue(__e, sym_dproperty_1vector_d)

tmp10128 := Call4(__e, SymFn(symput), let__9894, symshen_4dynamic, Nil, tmp10127)


ifres10124 = tmp10128


} else {
ifres10124 = symshen_4skip


}

let__9899 := ifres10124
_ = let__9899

tmp10130 := GTl(__e, V1601)

tmp10131 := GTl(__e, tmp10130)

tmp10132 := Call5(__e, SymFn(symshen_4insert_1info), let__9894, let__9895, tmp10131, V1601, V1602)


let__9900 := tmp10132
_ = let__9900

return let__9894


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4assert_d)
}


panic("unreachable")
})

tmp10144 := Call2(__e, ns2_1set, symshen_4assert_d, tmp10116)


_ = tmp10144

tmp10145 := MakeNative1(func(__e *ControlFlow, V1612 Obj) Obj {
tmp10147 := GConsP(__e, V1612)

if True == tmp10147 {
return GHd(__e, V1612)
} else {
return V1612
}


panic("unreachable")
})

tmp10148 := Call2(__e, ns2_1set, symshen_4predicate, tmp10145)


_ = tmp10148

tmp10149 := MakeNative1(func(__e *ControlFlow, V1617 Obj) Obj {
tmp10151 := GConsP(__e, V1617)

if True == tmp10151 {
return GTl(__e, V1617)
} else {
return Nil
}


panic("unreachable")
})

tmp10152 := Call2(__e, ns2_1set, symshen_4terms, tmp10149)


_ = tmp10152

tmp10153 := MakeNative2(func(__e *ControlFlow, V1618, V1619 Obj) Obj {
tmp10154 := Call2(__e, SymFn(symshen_4dynamic_1default), V1618, V1619)


tmp10155 := GCons(__e, V1618, tmp10154)

return GCons(__e, symdefprolog, tmp10155)


panic("unreachable")
})

tmp10156 := Call2(__e, ns2_1set, symshen_4create_1skeleton, tmp10153)


_ = tmp10156

tmp10157 := MakeNative2(func(__e *ControlFlow, V1620, V1621 Obj) Obj {
tmp10158 := Call1(__e, SymFn(symshen_4cons_1form), V1621)


tmp10159 := GCons(__e, symshen_4dynamic, Nil)

tmp10160 := GCons(__e, V1620, tmp10159)

tmp10161 := GCons(__e, symget, tmp10160)

tmp10162 := GCons(__e, tmp10161, Nil)

tmp10163 := GCons(__e, tmp10158, tmp10162)

tmp10164 := GCons(__e, symshen_4call_1dynamic, tmp10163)

tmp10165 := PrimGuard1(__e, symintern, PrimIntern, MakeString(";"))

tmp10166 := GCons(__e, tmp10165, Nil)

tmp10167 := GCons(__e, tmp10164, tmp10166)

tmp10168 := GCons(__e, sym_5_1_1, tmp10167)

return TailCall2(__e, SymFn(symappend), V1621, tmp10168)


panic("unreachable")
})

tmp10169 := Call2(__e, ns2_1set, symshen_4dynamic_1default, tmp10157)


_ = tmp10169

tmp10170 := MakeNative5(func(__e *ControlFlow, V1622, V1623, V1624, V1625, V1626 Obj) Obj {
tmp10171 := Call1(__e, SymFn(symgensym), symg)


let__9901 := tmp10171
_ = let__9901

tmp10172 := GCons(__e, let__9901, Nil)

tmp10173 := GCons(__e, symdefprolog, tmp10172)

tmp10174 := GCons(__e, sym_5_1_1, V1624)

tmp10175 := Call2(__e, SymFn(symappend), V1623, tmp10174)


tmp10176 := Call2(__e, SymFn(symappend), tmp10173, tmp10175)


tmp10177 := Call1(__e, SymFn(symeval), tmp10176)


let__9902 := tmp10177
_ = let__9902

tmp10178 := Call1(__e, SymFn(symfn), let__9901)


tmp10179 := GCons(__e, let__9901, V1625)

tmp10180 := GCons(__e, tmp10178, tmp10179)

let__9903 := tmp10180
_ = let__9903

tmp10181 := GValue(__e, sym_dproperty_1vector_d)

tmp10182 := Call3(__e, SymFn(symget), V1622, symshen_4dynamic, tmp10181)


let__9904 := tmp10182
_ = let__9904

tmp10187 := GEq(__e, V1626, symshen_4top)

var ifres10183 Obj

if True == tmp10187 {
tmp10184 := GCons(__e, let__9903, let__9904)

ifres10183 = tmp10184


} else {
tmp10185 := GCons(__e, let__9903, Nil)

tmp10186 := Call2(__e, SymFn(symappend), let__9904, tmp10185)


ifres10183 = tmp10186


}

let__9905 := ifres10183
_ = let__9905

tmp10188 := GValue(__e, sym_dproperty_1vector_d)

return TailCall4(__e, SymFn(symput), V1622, symshen_4dynamic, let__9905, tmp10188)


panic("unreachable")
})

tmp10189 := Call2(__e, ns2_1set, symshen_4insert_1info, tmp10170)


_ = tmp10189

tmp10190 := MakeNative0(func(__e *ControlFlow) Obj {
tmp10191 := GValue(__e, symshen_4_dnames_d)

let__9906 := tmp10191
_ = let__9906

tmp10197 := Call1(__e, SymFn(symempty_2), let__9906)


var ifres10192 Obj

if True == tmp10197 {
tmp10193 := Call1(__e, SymFn(symgensym), symg)


ifres10192 = tmp10193


} else {
tmp10194 := GTl(__e, let__9906)

tmp10195 := PrimGuard2(__e, symset, PrimSet, symshen_4_dnames_d, tmp10194)

_ = tmp10195

tmp10196 := GHd(__e, let__9906)

ifres10192 = tmp10196


}

let__9907 := ifres10192
_ = let__9907

return let__9907


panic("unreachable")
})

tmp10198 := Call2(__e, ns2_1set, symshen_4newname, tmp10190)


_ = tmp10198

tmp10199 := (func() Obj {
var __self Obj
__self = MakeNative6(func(__e *ControlFlow, __self1, __self2, __self3, __self4, __self5, __self6 Obj) Obj {
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
tmp10211 := Call1(__e, SymFn(symshen_4unlocked_2), V1637)


var ifres10200 Obj

if True == tmp10211 {
tmp10201 := Call2(__e, SymFn(symshen_4lazyderef), V1635, V1636)


let__9909 := tmp10201
_ = let__9909

tmp10210 := GConsP(__e, let__9909)

var ifres10202 Obj

if True == tmp10210 {
tmp10203 := GHd(__e, let__9909)

tmp10204 := Call2(__e, SymFn(symshen_4lazyderef), tmp10203, V1636)


let__9910 := tmp10204
_ = let__9910

tmp10209 := GConsP(__e, let__9910)

var ifres10205 Obj

if True == tmp10209 {
tmp10206 := GHd(__e, let__9910)

let__9911 := tmp10206
_ = let__9911

tmp10207 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp10207

tmp10208 := Call6(__e, SymFn(symshen_4callrec), let__9911, V1634, V1636, V1637, V1638, V1639)


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

tmp10220 := GEq(__e, let__9908, False)

if True == tmp10220 {
tmp10218 := Call1(__e, SymFn(symshen_4unlocked_2), V1637)


if True == tmp10218 {
tmp10212 := Call2(__e, SymFn(symshen_4lazyderef), V1635, V1636)


let__9912 := tmp10212
_ = let__9912

tmp10216 := GConsP(__e, let__9912)

if True == tmp10216 {
tmp10213 := GTl(__e, let__9912)

let__9913 := tmp10213
_ = let__9913

tmp10214 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp10214

if SymFn(symshen_4call_1dynamic) == __self {
__self1, __self2, __self3, __self4, __self5, __self6 = V1634, let__9913, V1636, V1637, V1638, V1639
__e.Tick()
goto __selftop
}
return TailCall6(__e, SymFn(symshen_4call_1dynamic), V1634, let__9913, V1636, V1637, V1638, V1639)


} else {
return False
}


} else {
return False
}


} else {
return let__9908
}


panic("unreachable")
})
return __self
})()

tmp10221 := Call2(__e, ns2_1set, symshen_4call_1dynamic, tmp10199)


_ = tmp10221

tmp10222 := (func() Obj {
var __self Obj
__self = MakeNative6(func(__e *ControlFlow, __self1, __self2, __self3, __self4, __self5, __self6 Obj) Obj {
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
tmp10232 := GEq(__e, Nil, V1647)

if True == tmp10232 {
tmp10223 := Call1(__e, V1646, V1648)


tmp10224 := Call1(__e, tmp10223, V1649)


tmp10225 := Call1(__e, tmp10224, V1650)


return TailCall1(__e, tmp10225, V1651)


} else {
tmp10230 := GConsP(__e, V1647)

if True == tmp10230 {
tmp10226 := GHd(__e, V1647)

tmp10227 := Call1(__e, V1646, tmp10226)


tmp10228 := GTl(__e, V1647)

if SymFn(symshen_4callrec) == __self {
__self1, __self2, __self3, __self4, __self5, __self6 = tmp10227, tmp10228, V1648, V1649, V1650, V1651
__e.Tick()
goto __selftop
}
return TailCall6(__e, SymFn(symshen_4callrec), tmp10227, tmp10228, V1648, V1649, V1650, V1651)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4callrec)
}


}


panic("unreachable")
})
return __self
})()

tmp10233 := Call2(__e, ns2_1set, symshen_4callrec, tmp10222)


_ = tmp10233

tmp10234 := MakeNative1(func(__e *ControlFlow, V1652 Obj) Obj {
tmp10251 := GConsP(__e, V1652)

var ifres10242 Obj

if True == tmp10251 {
tmp10249 := GTl(__e, V1652)

tmp10250 := GConsP(__e, tmp10249)

var ifres10244 Obj

if True == tmp10250 {
tmp10246 := GTl(__e, V1652)

tmp10247 := GHd(__e, tmp10246)

tmp10248 := GEq(__e, sym_5_1_1, tmp10247)

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
tmp10235 := GHd(__e, V1652)

tmp10236 := Call1(__e, SymFn(symshen_4predicate), tmp10235)


let__9914 := tmp10236
_ = let__9914

tmp10237 := GValue(__e, sym_dproperty_1vector_d)

tmp10238 := Call3(__e, SymFn(symget), let__9914, symshen_4dynamic, tmp10237)


let__9915 := tmp10238
_ = let__9915

tmp10239 := Call2(__e, SymFn(symshen_4retract_1clause), V1652, let__9915)


tmp10240 := GValue(__e, sym_dproperty_1vector_d)

return TailCall4(__e, SymFn(symput), let__9914, symshen_4dynamic, tmp10239, tmp10240)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symretract)
}


panic("unreachable")
})

tmp10252 := Call2(__e, ns2_1set, symretract, tmp10234)


_ = tmp10252

tmp10253 := MakeNative2(func(__e *ControlFlow, V1660, V1661 Obj) Obj {
tmp10283 := GEq(__e, Nil, V1661)

if True == tmp10283 {
return Nil
} else {
tmp10281 := GConsP(__e, V1661)

var ifres10266 Obj

if True == tmp10281 {
tmp10279 := GHd(__e, V1661)

tmp10280 := GConsP(__e, tmp10279)

var ifres10268 Obj

if True == tmp10280 {
tmp10276 := GHd(__e, V1661)

tmp10277 := GTl(__e, tmp10276)

tmp10278 := GConsP(__e, tmp10277)

var ifres10270 Obj

if True == tmp10278 {
tmp10272 := GHd(__e, V1661)

tmp10273 := GTl(__e, tmp10272)

tmp10274 := GTl(__e, tmp10273)

tmp10275 := GEq(__e, V1660, tmp10274)

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
tmp10254 := GHd(__e, V1661)

tmp10255 := GTl(__e, tmp10254)

tmp10256 := GHd(__e, tmp10255)

tmp10257 := GValue(__e, symshen_4_dnames_d)

tmp10258 := GCons(__e, tmp10256, tmp10257)

tmp10259 := PrimGuard2(__e, symset, PrimSet, symshen_4_dnames_d, tmp10258)

_ = tmp10259

return GTl(__e, V1661)


} else {
tmp10264 := GConsP(__e, V1661)

if True == tmp10264 {
tmp10260 := GHd(__e, V1661)

tmp10261 := GTl(__e, V1661)

tmp10262 := Call2(__e, SymFn(symshen_4retract_1clause), V1660, tmp10261)


return GCons(__e, tmp10260, tmp10262)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4retract_1clause)
}


}


}


panic("unreachable")
})

tmp10284 := Call2(__e, ns2_1set, symshen_4retract_1clause, tmp10253)


_ = tmp10284

tmp10285 := MakeNative2(func(__e *ControlFlow, V1662, V1663 Obj) Obj {
tmp10286 := MakeNative1(func(__e *ControlFlow, Z1664 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4_5defprolog_6), Z1664)
panic("unreachable")
})

tmp10287 := GCons(__e, V1662, V1663)

return TailCall2(__e, SymFn(symcompile), tmp10286, tmp10287)


panic("unreachable")
})

tmp10288 := Call2(__e, ns2_1set, symshen_4compile_1prolog, tmp10285)


_ = tmp10288

tmp10289 := MakeNative1(func(__e *ControlFlow, V1665 Obj) Obj {
tmp10305 := GConsP(__e, V1665)

var ifres10290 Obj

if True == tmp10305 {
tmp10291 := Call1(__e, SymFn(symhead), V1665)


let__9917 := tmp10291
_ = let__9917

tmp10292 := Call1(__e, SymFn(symtail), V1665)


let__9918 := tmp10292
_ = let__9918

tmp10293 := Call1(__e, SymFn(symshen_4_5clauses_6), let__9918)


let__9919 := tmp10293
_ = let__9919

tmp10303 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9919)


var ifres10294 Obj

if True == tmp10303 {
tmp10295 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10294 = tmp10295


} else {
tmp10296 := Call1(__e, SymFn(symshen_4_5_1out), let__9919)


let__9920 := tmp10296
_ = let__9920

tmp10297 := Call1(__e, SymFn(symshen_4in_1_6), let__9919)


let__9921 := tmp10297
_ = let__9921

tmp10298 := Call2(__e, SymFn(symshen_4prolog_1arity_1check), let__9917, let__9920)


let__9922 := tmp10298
_ = let__9922

tmp10299 := MakeNative1(func(__e *ControlFlow, Z1674 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4linearise_1clause), Z1674)
panic("unreachable")
})

tmp10300 := Call2(__e, SymFn(symmap), tmp10299, let__9920)


let__9923 := tmp10300
_ = let__9923

tmp10301 := Call2(__e, SymFn(symshen_4horn_1clause_1procedure), let__9917, let__9923)


tmp10302 := Call2(__e, SymFn(symshen_4comb), let__9921, tmp10301)


ifres10294 = tmp10302


}

ifres10290 = ifres10294


} else {
tmp10304 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10290 = tmp10304


}

let__9916 := ifres10290
_ = let__9916

tmp10307 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9916)


if True == tmp10307 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__9916
}


panic("unreachable")
})

tmp10308 := Call2(__e, ns2_1set, symshen_4_5defprolog_6, tmp10289)


_ = tmp10308

tmp10309 := MakeNative2(func(__e *ControlFlow, V1677, V1678 Obj) Obj {
tmp10353 := GConsP(__e, V1678)

var ifres10334 Obj

if True == tmp10353 {
tmp10351 := GHd(__e, V1678)

tmp10352 := GConsP(__e, tmp10351)

var ifres10336 Obj

if True == tmp10352 {
tmp10348 := GHd(__e, V1678)

tmp10349 := GTl(__e, tmp10348)

tmp10350 := GConsP(__e, tmp10349)

var ifres10338 Obj

if True == tmp10350 {
tmp10344 := GHd(__e, V1678)

tmp10345 := GTl(__e, tmp10344)

tmp10346 := GTl(__e, tmp10345)

tmp10347 := GEq(__e, Nil, tmp10346)

var ifres10340 Obj

if True == tmp10347 {
tmp10342 := GTl(__e, V1678)

tmp10343 := GEq(__e, Nil, tmp10342)

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
tmp10310 := GHd(__e, V1678)

tmp10311 := GHd(__e, tmp10310)

return TailCall1(__e, SymFn(symlength), tmp10311)


} else {
tmp10332 := GConsP(__e, V1678)

var ifres10317 Obj

if True == tmp10332 {
tmp10330 := GHd(__e, V1678)

tmp10331 := GConsP(__e, tmp10330)

var ifres10319 Obj

if True == tmp10331 {
tmp10327 := GHd(__e, V1678)

tmp10328 := GTl(__e, tmp10327)

tmp10329 := GConsP(__e, tmp10328)

var ifres10321 Obj

if True == tmp10329 {
tmp10323 := GHd(__e, V1678)

tmp10324 := GTl(__e, tmp10323)

tmp10325 := GTl(__e, tmp10324)

tmp10326 := GEq(__e, Nil, tmp10325)

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
tmp10312 := GHd(__e, V1678)

tmp10313 := GHd(__e, tmp10312)

tmp10314 := Call1(__e, SymFn(symlength), tmp10313)


tmp10315 := GTl(__e, V1678)

return TailCall3(__e, SymFn(symshen_4pac_1h), V1677, tmp10314, tmp10315)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4prolog_1arity_1check)
}


}


panic("unreachable")
})

tmp10354 := Call2(__e, ns2_1set, symshen_4prolog_1arity_1check, tmp10309)


_ = tmp10354

tmp10355 := (func() Obj {
var __self Obj
__self = MakeNative3(func(__e *ControlFlow, __self1, __self2, __self3 Obj) Obj {
__selftop:
V1683 := __self1
_ = V1683
V1684 := __self2
_ = V1684
V1685 := __self3
_ = V1685
tmp10371 := GEq(__e, Nil, V1685)

if True == tmp10371 {
return V1684
} else {
tmp10369 := GConsP(__e, V1685)

var ifres10365 Obj

if True == tmp10369 {
tmp10367 := GHd(__e, V1685)

tmp10368 := GConsP(__e, tmp10367)

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
tmp10360 := GHd(__e, V1685)

tmp10361 := GHd(__e, tmp10360)

tmp10362 := Call1(__e, SymFn(symlength), tmp10361)


tmp10363 := GEq(__e, V1684, tmp10362)

if True == tmp10363 {
tmp10356 := GTl(__e, V1685)

if SymFn(symshen_4pac_1h) == __self {
__self1, __self2, __self3 = V1683, V1684, tmp10356
__e.Tick()
goto __selftop
}
return TailCall3(__e, SymFn(symshen_4pac_1h), V1683, V1684, tmp10356)


} else {
tmp10357 := Call3(__e, SymFn(symshen_4app), V1683, MakeString("\n"), symshen_4a)


return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("arity error in prolog procedure "), tmp10357))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("arity error in prolog procedure "), tmp10357)
return Call1(__e, SymFn(symsimple_1error), __typedArg0)
})()


}


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4pac_1h)
}


}


panic("unreachable")
})
return __self
})()

tmp10372 := Call2(__e, ns2_1set, symshen_4pac_1h, tmp10355)


_ = tmp10372

tmp10373 := MakeNative1(func(__e *ControlFlow, V1686 Obj) Obj {
tmp10374 := Call1(__e, SymFn(symshen_4_5clause_6), V1686)


let__9925 := tmp10374
_ = let__9925

tmp10387 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9925)


var ifres10375 Obj

if True == tmp10387 {
tmp10376 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10375 = tmp10376


} else {
tmp10377 := Call1(__e, SymFn(symshen_4_5_1out), let__9925)


let__9926 := tmp10377
_ = let__9926

tmp10378 := Call1(__e, SymFn(symshen_4in_1_6), let__9925)


let__9927 := tmp10378
_ = let__9927

tmp10379 := Call1(__e, SymFn(symshen_4_5clauses_6), let__9927)


let__9928 := tmp10379
_ = let__9928

tmp10386 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9928)


var ifres10380 Obj

if True == tmp10386 {
tmp10381 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10380 = tmp10381


} else {
tmp10382 := Call1(__e, SymFn(symshen_4_5_1out), let__9928)


let__9929 := tmp10382
_ = let__9929

tmp10383 := Call1(__e, SymFn(symshen_4in_1_6), let__9928)


let__9930 := tmp10383
_ = let__9930

tmp10384 := GCons(__e, let__9926, let__9929)

tmp10385 := Call2(__e, SymFn(symshen_4comb), let__9930, tmp10384)


ifres10380 = tmp10385


}

ifres10375 = ifres10380


}

let__9924 := ifres10375
_ = let__9924

tmp10403 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9924)


if True == tmp10403 {
tmp10388 := Call1(__e, SymFn(sym_5_b_6), V1686)


let__9932 := tmp10388
_ = let__9932

tmp10399 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9932)


var ifres10389 Obj

if True == tmp10399 {
tmp10390 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10389 = tmp10390


} else {
tmp10391 := Call1(__e, SymFn(symshen_4_5_1out), let__9932)


let__9933 := tmp10391
_ = let__9933

tmp10392 := Call1(__e, SymFn(symshen_4in_1_6), let__9932)


let__9934 := tmp10392
_ = let__9934

tmp10397 := Call1(__e, SymFn(symempty_2), let__9933)


var ifres10393 Obj

if True == tmp10397 {
ifres10393 = Nil


} else {
tmp10394 := Call3(__e, SymFn(symshen_4app), let__9933, MakeString("\n ..."), symshen_4r)


tmp10396 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("Prolog syntax error here:\n "), tmp10394))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("Prolog syntax error here:\n "), tmp10394)
return Call1(__e, SymFn(symsimple_1error), __typedArg0)
})()

ifres10393 = tmp10396


}

tmp10398 := Call2(__e, SymFn(symshen_4comb), let__9934, ifres10393)


ifres10389 = tmp10398


}

let__9931 := ifres10389
_ = let__9931

tmp10401 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9931)


if True == tmp10401 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__9931
}


} else {
return let__9924
}


panic("unreachable")
})

tmp10404 := Call2(__e, ns2_1set, symshen_4_5clauses_6, tmp10373)


_ = tmp10404

tmp10405 := MakeNative1(func(__e *ControlFlow, V1698 Obj) Obj {
tmp10421 := GConsP(__e, V1698)

var ifres10412 Obj

if True == tmp10421 {
tmp10419 := GTl(__e, V1698)

tmp10420 := GConsP(__e, tmp10419)

var ifres10414 Obj

if True == tmp10420 {
tmp10416 := GTl(__e, V1698)

tmp10417 := GTl(__e, tmp10416)

tmp10418 := GEq(__e, Nil, tmp10417)

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
tmp10406 := GHd(__e, V1698)

tmp10407 := GTl(__e, V1698)

tmp10408 := GHd(__e, tmp10407)

tmp10409 := Call2(__e, SymFn(sym_8p), tmp10406, tmp10408)


tmp10410 := Call1(__e, SymFn(symshen_4linearise), tmp10409)


return TailCall1(__e, SymFn(symshen_4lch), tmp10410)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4linearise_1clause)
}


panic("unreachable")
})

tmp10422 := Call2(__e, ns2_1set, symshen_4linearise_1clause, tmp10405)


_ = tmp10422

tmp10423 := MakeNative1(func(__e *ControlFlow, V1699 Obj) Obj {
tmp10429 := Call1(__e, SymFn(symtuple_2), V1699)


if True == tmp10429 {
tmp10424 := Call1(__e, SymFn(symfst), V1699)


tmp10425 := Call1(__e, SymFn(symsnd), V1699)


tmp10426 := Call1(__e, SymFn(symshen_4lchh), tmp10425)


tmp10427 := GCons(__e, tmp10426, Nil)

return GCons(__e, tmp10424, tmp10427)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4lch)
}


panic("unreachable")
})

tmp10430 := Call2(__e, ns2_1set, symshen_4lch, tmp10423)


_ = tmp10430

tmp10431 := MakeNative1(func(__e *ControlFlow, V1700 Obj) Obj {
tmp10494 := GConsP(__e, V1700)

var ifres10443 Obj

if True == tmp10494 {
tmp10492 := GHd(__e, V1700)

tmp10493 := GEq(__e, symwhere, tmp10492)

var ifres10445 Obj

if True == tmp10493 {
tmp10490 := GTl(__e, V1700)

tmp10491 := GConsP(__e, tmp10490)

var ifres10447 Obj

if True == tmp10491 {
tmp10487 := GTl(__e, V1700)

tmp10488 := GHd(__e, tmp10487)

tmp10489 := GConsP(__e, tmp10488)

var ifres10449 Obj

if True == tmp10489 {
tmp10483 := GTl(__e, V1700)

tmp10484 := GHd(__e, tmp10483)

tmp10485 := GHd(__e, tmp10484)

tmp10486 := GEq(__e, sym_a, tmp10485)

var ifres10451 Obj

if True == tmp10486 {
tmp10479 := GTl(__e, V1700)

tmp10480 := GHd(__e, tmp10479)

tmp10481 := GTl(__e, tmp10480)

tmp10482 := GConsP(__e, tmp10481)

var ifres10453 Obj

if True == tmp10482 {
tmp10474 := GTl(__e, V1700)

tmp10475 := GHd(__e, tmp10474)

tmp10476 := GTl(__e, tmp10475)

tmp10477 := GTl(__e, tmp10476)

tmp10478 := GConsP(__e, tmp10477)

var ifres10455 Obj

if True == tmp10478 {
tmp10468 := GTl(__e, V1700)

tmp10469 := GHd(__e, tmp10468)

tmp10470 := GTl(__e, tmp10469)

tmp10471 := GTl(__e, tmp10470)

tmp10472 := GTl(__e, tmp10471)

tmp10473 := GEq(__e, Nil, tmp10472)

var ifres10457 Obj

if True == tmp10473 {
tmp10465 := GTl(__e, V1700)

tmp10466 := GTl(__e, tmp10465)

tmp10467 := GConsP(__e, tmp10466)

var ifres10459 Obj

if True == tmp10467 {
tmp10461 := GTl(__e, V1700)

tmp10462 := GTl(__e, tmp10461)

tmp10463 := GTl(__e, tmp10462)

tmp10464 := GEq(__e, Nil, tmp10463)

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
tmp10433 := GValue(__e, symshen_4_doccurs_d)

var ifres10432 Obj

if True == tmp10433 {
ifres10432 = symis_b


} else {
ifres10432 = symis


}

tmp10434 := GTl(__e, V1700)

tmp10435 := GHd(__e, tmp10434)

tmp10436 := GTl(__e, tmp10435)

tmp10437 := GCons(__e, ifres10432, tmp10436)

tmp10438 := GTl(__e, V1700)

tmp10439 := GTl(__e, tmp10438)

tmp10440 := GHd(__e, tmp10439)

tmp10441 := Call1(__e, SymFn(symshen_4lchh), tmp10440)


return GCons(__e, tmp10437, tmp10441)


} else {
return V1700
}


panic("unreachable")
})

tmp10495 := Call2(__e, ns2_1set, symshen_4lchh, tmp10431)


_ = tmp10495

tmp10496 := MakeNative1(func(__e *ControlFlow, V1701 Obj) Obj {
tmp10497 := Call1(__e, SymFn(symshen_4_5head_6), V1701)


let__9936 := tmp10497
_ = let__9936

tmp10520 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9936)


var ifres10498 Obj

if True == tmp10520 {
tmp10499 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10498 = tmp10499


} else {
tmp10500 := Call1(__e, SymFn(symshen_4_5_1out), let__9936)


let__9937 := tmp10500
_ = let__9937

tmp10501 := Call1(__e, SymFn(symshen_4in_1_6), let__9936)


let__9938 := tmp10501
_ = let__9938

tmp10519 := Call2(__e, SymFn(symshen_4hds_a_2), let__9938, sym_5_1_1)


var ifres10502 Obj

if True == tmp10519 {
tmp10503 := Call1(__e, SymFn(symtail), let__9938)


let__9939 := tmp10503
_ = let__9939

tmp10504 := Call1(__e, SymFn(symshen_4_5body_6), let__9939)


let__9940 := tmp10504
_ = let__9940

tmp10517 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9940)


var ifres10505 Obj

if True == tmp10517 {
tmp10506 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10505 = tmp10506


} else {
tmp10507 := Call1(__e, SymFn(symshen_4_5_1out), let__9940)


let__9941 := tmp10507
_ = let__9941

tmp10508 := Call1(__e, SymFn(symshen_4in_1_6), let__9940)


let__9942 := tmp10508
_ = let__9942

tmp10509 := Call1(__e, SymFn(symshen_4_5sc_6), let__9942)


let__9943 := tmp10509
_ = let__9943

tmp10516 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9943)


var ifres10510 Obj

if True == tmp10516 {
tmp10511 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10510 = tmp10511


} else {
tmp10512 := Call1(__e, SymFn(symshen_4in_1_6), let__9943)


let__9944 := tmp10512
_ = let__9944

tmp10513 := GCons(__e, let__9941, Nil)

tmp10514 := GCons(__e, let__9937, tmp10513)

tmp10515 := Call2(__e, SymFn(symshen_4comb), let__9944, tmp10514)


ifres10510 = tmp10515


}

ifres10505 = ifres10510


}

ifres10502 = ifres10505


} else {
tmp10518 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10502 = tmp10518


}

ifres10498 = ifres10502


}

let__9935 := ifres10498
_ = let__9935

tmp10522 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9935)


if True == tmp10522 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__9935
}


panic("unreachable")
})

tmp10523 := Call2(__e, ns2_1set, symshen_4_5clause_6, tmp10496)


_ = tmp10523

tmp10524 := MakeNative1(func(__e *ControlFlow, V1712 Obj) Obj {
tmp10525 := Call1(__e, SymFn(symshen_4_5hterm_6), V1712)


let__9946 := tmp10525
_ = let__9946

tmp10538 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9946)


var ifres10526 Obj

if True == tmp10538 {
tmp10527 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10526 = tmp10527


} else {
tmp10528 := Call1(__e, SymFn(symshen_4_5_1out), let__9946)


let__9947 := tmp10528
_ = let__9947

tmp10529 := Call1(__e, SymFn(symshen_4in_1_6), let__9946)


let__9948 := tmp10529
_ = let__9948

tmp10530 := Call1(__e, SymFn(symshen_4_5head_6), let__9948)


let__9949 := tmp10530
_ = let__9949

tmp10537 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9949)


var ifres10531 Obj

if True == tmp10537 {
tmp10532 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10531 = tmp10532


} else {
tmp10533 := Call1(__e, SymFn(symshen_4_5_1out), let__9949)


let__9950 := tmp10533
_ = let__9950

tmp10534 := Call1(__e, SymFn(symshen_4in_1_6), let__9949)


let__9951 := tmp10534
_ = let__9951

tmp10535 := GCons(__e, let__9947, let__9950)

tmp10536 := Call2(__e, SymFn(symshen_4comb), let__9951, tmp10535)


ifres10531 = tmp10536


}

ifres10526 = ifres10531


}

let__9945 := ifres10526
_ = let__9945

tmp10548 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9945)


if True == tmp10548 {
tmp10539 := Call1(__e, SymFn(sym_5e_6), V1712)


let__9953 := tmp10539
_ = let__9953

tmp10544 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9953)


var ifres10540 Obj

if True == tmp10544 {
tmp10541 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10540 = tmp10541


} else {
tmp10542 := Call1(__e, SymFn(symshen_4in_1_6), let__9953)


let__9954 := tmp10542
_ = let__9954

tmp10543 := Call2(__e, SymFn(symshen_4comb), let__9954, Nil)


ifres10540 = tmp10543


}

let__9952 := ifres10540
_ = let__9952

tmp10546 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9952)


if True == tmp10546 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__9952
}


} else {
return let__9945
}


panic("unreachable")
})

tmp10549 := Call2(__e, ns2_1set, symshen_4_5head_6, tmp10524)


_ = tmp10549

tmp10550 := MakeNative1(func(__e *ControlFlow, V1723 Obj) Obj {
tmp10563 := GConsP(__e, V1723)

var ifres10551 Obj

if True == tmp10563 {
tmp10552 := Call1(__e, SymFn(symhead), V1723)


let__9956 := tmp10552
_ = let__9956

tmp10553 := Call1(__e, SymFn(symtail), V1723)


let__9957 := tmp10553
_ = let__9957

tmp10561 := Call1(__e, SymFn(symatom_2), let__9956)


var ifres10557 Obj

if True == tmp10561 {
tmp10559 := Call1(__e, SymFn(symshen_4prolog_1keyword_2), let__9956)


tmp10560 := GNot(__e, tmp10559)

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
tmp10555 := Call2(__e, SymFn(symshen_4comb), let__9957, let__9956)


ifres10554 = tmp10555


} else {
tmp10556 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10554 = tmp10556


}

ifres10551 = ifres10554


} else {
tmp10562 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10551 = tmp10562


}

let__9955 := ifres10551
_ = let__9955

tmp10717 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9955)


if True == tmp10717 {
tmp10573 := GConsP(__e, V1723)

var ifres10564 Obj

if True == tmp10573 {
tmp10565 := Call1(__e, SymFn(symhead), V1723)


let__9959 := tmp10565
_ = let__9959

tmp10566 := Call1(__e, SymFn(symtail), V1723)


let__9960 := tmp10566
_ = let__9960

tmp10570 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":"))

tmp10571 := GEq(__e, let__9959, tmp10570)

var ifres10567 Obj

if True == tmp10571 {
tmp10568 := Call2(__e, SymFn(symshen_4comb), let__9960, let__9959)


ifres10567 = tmp10568


} else {
tmp10569 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10567 = tmp10569


}

ifres10564 = ifres10567


} else {
tmp10572 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10564 = tmp10572


}

let__9958 := ifres10564
_ = let__9958

tmp10715 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9958)


if True == tmp10715 {
tmp10603 := Call1(__e, SymFn(symshen_4ccons_2), V1723)


var ifres10574 Obj

if True == tmp10603 {
tmp10575 := Call1(__e, SymFn(symhead), V1723)


let__9962 := tmp10575
_ = let__9962

tmp10576 := Call1(__e, SymFn(symtail), V1723)


let__9963 := tmp10576
_ = let__9963

tmp10601 := Call2(__e, SymFn(symshen_4hds_a_2), let__9962, symcons)


var ifres10577 Obj

if True == tmp10601 {
tmp10578 := Call1(__e, SymFn(symtail), let__9962)


let__9964 := tmp10578
_ = let__9964

tmp10579 := Call1(__e, SymFn(symshen_4_5hterm1_6), let__9964)


let__9965 := tmp10579
_ = let__9965

tmp10599 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9965)


var ifres10580 Obj

if True == tmp10599 {
tmp10581 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10580 = tmp10581


} else {
tmp10582 := Call1(__e, SymFn(symshen_4_5_1out), let__9965)


let__9966 := tmp10582
_ = let__9966

tmp10583 := Call1(__e, SymFn(symshen_4in_1_6), let__9965)


let__9967 := tmp10583
_ = let__9967

tmp10584 := Call1(__e, SymFn(symshen_4_5hterm2_6), let__9967)


let__9968 := tmp10584
_ = let__9968

tmp10598 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9968)


var ifres10585 Obj

if True == tmp10598 {
tmp10586 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10585 = tmp10586


} else {
tmp10587 := Call1(__e, SymFn(symshen_4_5_1out), let__9968)


let__9969 := tmp10587
_ = let__9969

tmp10588 := Call1(__e, SymFn(symshen_4in_1_6), let__9968)


let__9970 := tmp10588
_ = let__9970

tmp10589 := Call1(__e, SymFn(sym_5end_6), let__9970)


let__9971 := tmp10589
_ = let__9971

tmp10597 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9971)


var ifres10590 Obj

if True == tmp10597 {
tmp10591 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10590 = tmp10591


} else {
tmp10592 := Call1(__e, SymFn(symshen_4in_1_6), let__9971)


let__9972 := tmp10592
_ = let__9972

tmp10593 := GCons(__e, let__9969, Nil)

tmp10594 := GCons(__e, let__9966, tmp10593)

tmp10595 := GCons(__e, symcons, tmp10594)

tmp10596 := Call2(__e, SymFn(symshen_4comb), let__9963, tmp10595)


ifres10590 = tmp10596


}

ifres10585 = ifres10590


}

ifres10580 = ifres10585


}

ifres10577 = ifres10580


} else {
tmp10600 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10577 = tmp10600


}

ifres10574 = ifres10577


} else {
tmp10602 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10574 = tmp10602


}

let__9961 := ifres10574
_ = let__9961

tmp10713 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9961)


if True == tmp10713 {
tmp10626 := Call1(__e, SymFn(symshen_4ccons_2), V1723)


var ifres10604 Obj

if True == tmp10626 {
tmp10605 := Call1(__e, SymFn(symhead), V1723)


let__9974 := tmp10605
_ = let__9974

tmp10606 := Call1(__e, SymFn(symtail), V1723)


let__9975 := tmp10606
_ = let__9975

tmp10624 := Call2(__e, SymFn(symshen_4hds_a_2), let__9974, sym_7)


var ifres10607 Obj

if True == tmp10624 {
tmp10608 := Call1(__e, SymFn(symtail), let__9974)


let__9976 := tmp10608
_ = let__9976

tmp10609 := Call1(__e, SymFn(symshen_4_5hterm_6), let__9976)


let__9977 := tmp10609
_ = let__9977

tmp10622 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9977)


var ifres10610 Obj

if True == tmp10622 {
tmp10611 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10610 = tmp10611


} else {
tmp10612 := Call1(__e, SymFn(symshen_4_5_1out), let__9977)


let__9978 := tmp10612
_ = let__9978

tmp10613 := Call1(__e, SymFn(symshen_4in_1_6), let__9977)


let__9979 := tmp10613
_ = let__9979

tmp10614 := Call1(__e, SymFn(sym_5end_6), let__9979)


let__9980 := tmp10614
_ = let__9980

tmp10621 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9980)


var ifres10615 Obj

if True == tmp10621 {
tmp10616 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10615 = tmp10616


} else {
tmp10617 := Call1(__e, SymFn(symshen_4in_1_6), let__9980)


let__9981 := tmp10617
_ = let__9981

tmp10618 := GCons(__e, let__9978, Nil)

tmp10619 := GCons(__e, symshen_4_7m, tmp10618)

tmp10620 := Call2(__e, SymFn(symshen_4comb), let__9975, tmp10619)


ifres10615 = tmp10620


}

ifres10610 = ifres10615


}

ifres10607 = ifres10610


} else {
tmp10623 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10607 = tmp10623


}

ifres10604 = ifres10607


} else {
tmp10625 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10604 = tmp10625


}

let__9973 := ifres10604
_ = let__9973

tmp10711 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9973)


if True == tmp10711 {
tmp10649 := Call1(__e, SymFn(symshen_4ccons_2), V1723)


var ifres10627 Obj

if True == tmp10649 {
tmp10628 := Call1(__e, SymFn(symhead), V1723)


let__9983 := tmp10628
_ = let__9983

tmp10629 := Call1(__e, SymFn(symtail), V1723)


let__9984 := tmp10629
_ = let__9984

tmp10647 := Call2(__e, SymFn(symshen_4hds_a_2), let__9983, sym_1)


var ifres10630 Obj

if True == tmp10647 {
tmp10631 := Call1(__e, SymFn(symtail), let__9983)


let__9985 := tmp10631
_ = let__9985

tmp10632 := Call1(__e, SymFn(symshen_4_5hterm_6), let__9985)


let__9986 := tmp10632
_ = let__9986

tmp10645 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9986)


var ifres10633 Obj

if True == tmp10645 {
tmp10634 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10633 = tmp10634


} else {
tmp10635 := Call1(__e, SymFn(symshen_4_5_1out), let__9986)


let__9987 := tmp10635
_ = let__9987

tmp10636 := Call1(__e, SymFn(symshen_4in_1_6), let__9986)


let__9988 := tmp10636
_ = let__9988

tmp10637 := Call1(__e, SymFn(sym_5end_6), let__9988)


let__9989 := tmp10637
_ = let__9989

tmp10644 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9989)


var ifres10638 Obj

if True == tmp10644 {
tmp10639 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10638 = tmp10639


} else {
tmp10640 := Call1(__e, SymFn(symshen_4in_1_6), let__9989)


let__9990 := tmp10640
_ = let__9990

tmp10641 := GCons(__e, let__9987, Nil)

tmp10642 := GCons(__e, symshen_4_1m, tmp10641)

tmp10643 := Call2(__e, SymFn(symshen_4comb), let__9984, tmp10642)


ifres10638 = tmp10643


}

ifres10633 = ifres10638


}

ifres10630 = ifres10633


} else {
tmp10646 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10630 = tmp10646


}

ifres10627 = ifres10630


} else {
tmp10648 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10627 = tmp10648


}

let__9982 := ifres10627
_ = let__9982

tmp10709 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9982)


if True == tmp10709 {
tmp10676 := Call1(__e, SymFn(symshen_4ccons_2), V1723)


var ifres10650 Obj

if True == tmp10676 {
tmp10651 := Call1(__e, SymFn(symhead), V1723)


let__9992 := tmp10651
_ = let__9992

tmp10652 := Call1(__e, SymFn(symtail), V1723)


let__9993 := tmp10652
_ = let__9993

tmp10674 := Call2(__e, SymFn(symshen_4hds_a_2), let__9992, symmode)


var ifres10653 Obj

if True == tmp10674 {
tmp10654 := Call1(__e, SymFn(symtail), let__9992)


let__9994 := tmp10654
_ = let__9994

tmp10655 := Call1(__e, SymFn(symshen_4_5hterm_6), let__9994)


let__9995 := tmp10655
_ = let__9995

tmp10672 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9995)


var ifres10656 Obj

if True == tmp10672 {
tmp10657 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10656 = tmp10657


} else {
tmp10658 := Call1(__e, SymFn(symshen_4_5_1out), let__9995)


let__9996 := tmp10658
_ = let__9996

tmp10659 := Call1(__e, SymFn(symshen_4in_1_6), let__9995)


let__9997 := tmp10659
_ = let__9997

tmp10671 := Call2(__e, SymFn(symshen_4hds_a_2), let__9997, sym_7)


var ifres10660 Obj

if True == tmp10671 {
tmp10661 := Call1(__e, SymFn(symtail), let__9997)


let__9998 := tmp10661
_ = let__9998

tmp10662 := Call1(__e, SymFn(sym_5end_6), let__9998)


let__9999 := tmp10662
_ = let__9999

tmp10669 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9999)


var ifres10663 Obj

if True == tmp10669 {
tmp10664 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10663 = tmp10664


} else {
tmp10665 := Call1(__e, SymFn(symshen_4in_1_6), let__9999)


let__10000 := tmp10665
_ = let__10000

tmp10666 := GCons(__e, let__9996, Nil)

tmp10667 := GCons(__e, symshen_4_7m, tmp10666)

tmp10668 := Call2(__e, SymFn(symshen_4comb), let__9993, tmp10667)


ifres10663 = tmp10668


}

ifres10660 = ifres10663


} else {
tmp10670 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10660 = tmp10670


}

ifres10656 = ifres10660


}

ifres10653 = ifres10656


} else {
tmp10673 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10653 = tmp10673


}

ifres10650 = ifres10653


} else {
tmp10675 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10650 = tmp10675


}

let__9991 := ifres10650
_ = let__9991

tmp10707 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__9991)


if True == tmp10707 {
tmp10703 := Call1(__e, SymFn(symshen_4ccons_2), V1723)


var ifres10677 Obj

if True == tmp10703 {
tmp10678 := Call1(__e, SymFn(symhead), V1723)


let__10002 := tmp10678
_ = let__10002

tmp10679 := Call1(__e, SymFn(symtail), V1723)


let__10003 := tmp10679
_ = let__10003

tmp10701 := Call2(__e, SymFn(symshen_4hds_a_2), let__10002, symmode)


var ifres10680 Obj

if True == tmp10701 {
tmp10681 := Call1(__e, SymFn(symtail), let__10002)


let__10004 := tmp10681
_ = let__10004

tmp10682 := Call1(__e, SymFn(symshen_4_5hterm_6), let__10004)


let__10005 := tmp10682
_ = let__10005

tmp10699 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__10005)


var ifres10683 Obj

if True == tmp10699 {
tmp10684 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10683 = tmp10684


} else {
tmp10685 := Call1(__e, SymFn(symshen_4_5_1out), let__10005)


let__10006 := tmp10685
_ = let__10006

tmp10686 := Call1(__e, SymFn(symshen_4in_1_6), let__10005)


let__10007 := tmp10686
_ = let__10007

tmp10698 := Call2(__e, SymFn(symshen_4hds_a_2), let__10007, sym_1)


var ifres10687 Obj

if True == tmp10698 {
tmp10688 := Call1(__e, SymFn(symtail), let__10007)


let__10008 := tmp10688
_ = let__10008

tmp10689 := Call1(__e, SymFn(sym_5end_6), let__10008)


let__10009 := tmp10689
_ = let__10009

tmp10696 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__10009)


var ifres10690 Obj

if True == tmp10696 {
tmp10691 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10690 = tmp10691


} else {
tmp10692 := Call1(__e, SymFn(symshen_4in_1_6), let__10009)


let__10010 := tmp10692
_ = let__10010

tmp10693 := GCons(__e, let__10006, Nil)

tmp10694 := GCons(__e, symshen_4_1m, tmp10693)

tmp10695 := Call2(__e, SymFn(symshen_4comb), let__10003, tmp10694)


ifres10690 = tmp10695


}

ifres10687 = ifres10690


} else {
tmp10697 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10687 = tmp10697


}

ifres10683 = ifres10687


}

ifres10680 = ifres10683


} else {
tmp10700 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10680 = tmp10700


}

ifres10677 = ifres10680


} else {
tmp10702 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10677 = tmp10702


}

let__10001 := ifres10677
_ = let__10001

tmp10705 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__10001)


if True == tmp10705 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__10001
}


} else {
return let__9991
}


} else {
return let__9982
}


} else {
return let__9973
}


} else {
return let__9961
}


} else {
return let__9958
}


} else {
return let__9955
}


panic("unreachable")
})

tmp10718 := Call2(__e, ns2_1set, symshen_4_5hterm_6, tmp10550)


_ = tmp10718

tmp10719 := MakeNative1(func(__e *ControlFlow, V1780 Obj) Obj {
tmp10720 := PrimGuard1(__e, symintern, PrimIntern, MakeString(";"))

tmp10721 := GCons(__e, sym_5_1_1, Nil)

tmp10722 := GCons(__e, tmp10720, tmp10721)

return TailCall2(__e, SymFn(symelement_2), V1780, tmp10722)


panic("unreachable")
})

tmp10723 := Call2(__e, ns2_1set, symshen_4prolog_1keyword_2, tmp10719)


_ = tmp10723

tmp10724 := MakeNative1(func(__e *ControlFlow, V1781 Obj) Obj {
tmp10737 := PrimGuard1(__e, symsymbol_2, PrimIsSymbol, V1781)

if True == tmp10737 {
return True
} else {
tmp10735 := PrimGuard1(__e, symstring_2, PrimIsString, V1781)

var ifres10726 Obj

if True == tmp10735 {
ifres10726 = True


} else {
tmp10734 := Call1(__e, SymFn(symboolean_2), V1781)


var ifres10728 Obj

if True == tmp10734 {
ifres10728 = True


} else {
tmp10733 := PrimGuard1(__e, symnumber_2, PrimIsNumber, V1781)

var ifres10730 Obj

if True == tmp10733 {
ifres10730 = True


} else {
tmp10732 := Call1(__e, SymFn(symempty_2), V1781)


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
return True
} else {
return False
}


}


panic("unreachable")
})

tmp10738 := Call2(__e, ns2_1set, symatom_2, tmp10724)


_ = tmp10738

tmp10739 := MakeNative1(func(__e *ControlFlow, V1782 Obj) Obj {
tmp10740 := Call1(__e, SymFn(symshen_4_5hterm_6), V1782)


let__10012 := tmp10740
_ = let__10012

tmp10746 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__10012)


var ifres10741 Obj

if True == tmp10746 {
tmp10742 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10741 = tmp10742


} else {
tmp10743 := Call1(__e, SymFn(symshen_4_5_1out), let__10012)


let__10013 := tmp10743
_ = let__10013

tmp10744 := Call1(__e, SymFn(symshen_4in_1_6), let__10012)


let__10014 := tmp10744
_ = let__10014

tmp10745 := Call2(__e, SymFn(symshen_4comb), let__10014, let__10013)


ifres10741 = tmp10745


}

let__10011 := ifres10741
_ = let__10011

tmp10748 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__10011)


if True == tmp10748 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__10011
}


panic("unreachable")
})

tmp10749 := Call2(__e, ns2_1set, symshen_4_5hterm1_6, tmp10739)


_ = tmp10749

tmp10750 := MakeNative1(func(__e *ControlFlow, V1787 Obj) Obj {
tmp10751 := Call1(__e, SymFn(symshen_4_5hterm_6), V1787)


let__10016 := tmp10751
_ = let__10016

tmp10757 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__10016)


var ifres10752 Obj

if True == tmp10757 {
tmp10753 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10752 = tmp10753


} else {
tmp10754 := Call1(__e, SymFn(symshen_4_5_1out), let__10016)


let__10017 := tmp10754
_ = let__10017

tmp10755 := Call1(__e, SymFn(symshen_4in_1_6), let__10016)


let__10018 := tmp10755
_ = let__10018

tmp10756 := Call2(__e, SymFn(symshen_4comb), let__10018, let__10017)


ifres10752 = tmp10756


}

let__10015 := ifres10752
_ = let__10015

tmp10759 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__10015)


if True == tmp10759 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__10015
}


panic("unreachable")
})

tmp10760 := Call2(__e, ns2_1set, symshen_4_5hterm2_6, tmp10750)


_ = tmp10760

tmp10761 := MakeNative1(func(__e *ControlFlow, V1792 Obj) Obj {
tmp10762 := Call1(__e, SymFn(symshen_4_5literal_6), V1792)


let__10020 := tmp10762
_ = let__10020

tmp10775 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__10020)


var ifres10763 Obj

if True == tmp10775 {
tmp10764 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10763 = tmp10764


} else {
tmp10765 := Call1(__e, SymFn(symshen_4_5_1out), let__10020)


let__10021 := tmp10765
_ = let__10021

tmp10766 := Call1(__e, SymFn(symshen_4in_1_6), let__10020)


let__10022 := tmp10766
_ = let__10022

tmp10767 := Call1(__e, SymFn(symshen_4_5body_6), let__10022)


let__10023 := tmp10767
_ = let__10023

tmp10774 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__10023)


var ifres10768 Obj

if True == tmp10774 {
tmp10769 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10768 = tmp10769


} else {
tmp10770 := Call1(__e, SymFn(symshen_4_5_1out), let__10023)


let__10024 := tmp10770
_ = let__10024

tmp10771 := Call1(__e, SymFn(symshen_4in_1_6), let__10023)


let__10025 := tmp10771
_ = let__10025

tmp10772 := GCons(__e, let__10021, let__10024)

tmp10773 := Call2(__e, SymFn(symshen_4comb), let__10025, tmp10772)


ifres10768 = tmp10773


}

ifres10763 = ifres10768


}

let__10019 := ifres10763
_ = let__10019

tmp10785 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__10019)


if True == tmp10785 {
tmp10776 := Call1(__e, SymFn(sym_5e_6), V1792)


let__10027 := tmp10776
_ = let__10027

tmp10781 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__10027)


var ifres10777 Obj

if True == tmp10781 {
tmp10778 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10777 = tmp10778


} else {
tmp10779 := Call1(__e, SymFn(symshen_4in_1_6), let__10027)


let__10028 := tmp10779
_ = let__10028

tmp10780 := Call2(__e, SymFn(symshen_4comb), let__10028, Nil)


ifres10777 = tmp10780


}

let__10026 := ifres10777
_ = let__10026

tmp10783 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__10026)


if True == tmp10783 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__10026
}


} else {
return let__10019
}


panic("unreachable")
})

tmp10786 := Call2(__e, ns2_1set, symshen_4_5body_6, tmp10761)


_ = tmp10786

tmp10787 := MakeNative1(func(__e *ControlFlow, V1803 Obj) Obj {
tmp10792 := Call2(__e, SymFn(symshen_4hds_a_2), V1803, sym_b)


var ifres10788 Obj

if True == tmp10792 {
tmp10789 := Call1(__e, SymFn(symtail), V1803)


let__10030 := tmp10789
_ = let__10030

tmp10790 := Call2(__e, SymFn(symshen_4comb), let__10030, sym_b)


ifres10788 = tmp10790


} else {
tmp10791 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10788 = tmp10791


}

let__10029 := ifres10788
_ = let__10029

tmp10813 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__10029)


if True == tmp10813 {
tmp10809 := Call1(__e, SymFn(symshen_4ccons_2), V1803)


var ifres10793 Obj

if True == tmp10809 {
tmp10794 := Call1(__e, SymFn(symhead), V1803)


let__10032 := tmp10794
_ = let__10032

tmp10795 := Call1(__e, SymFn(symtail), V1803)


let__10033 := tmp10795
_ = let__10033

tmp10796 := Call1(__e, SymFn(symshen_4_5bterms_6), let__10032)


let__10034 := tmp10796
_ = let__10034

tmp10807 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__10034)


var ifres10797 Obj

if True == tmp10807 {
tmp10798 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10797 = tmp10798


} else {
tmp10799 := Call1(__e, SymFn(symshen_4_5_1out), let__10034)


let__10035 := tmp10799
_ = let__10035

tmp10800 := Call1(__e, SymFn(symshen_4in_1_6), let__10034)


let__10036 := tmp10800
_ = let__10036

tmp10801 := Call1(__e, SymFn(sym_5end_6), let__10036)


let__10037 := tmp10801
_ = let__10037

tmp10806 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__10037)


var ifres10802 Obj

if True == tmp10806 {
tmp10803 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10802 = tmp10803


} else {
tmp10804 := Call1(__e, SymFn(symshen_4in_1_6), let__10037)


let__10038 := tmp10804
_ = let__10038

tmp10805 := Call2(__e, SymFn(symshen_4comb), let__10033, let__10035)


ifres10802 = tmp10805


}

ifres10797 = ifres10802


}

ifres10793 = ifres10797


} else {
tmp10808 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10793 = tmp10808


}

let__10031 := ifres10793
_ = let__10031

tmp10811 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__10031)


if True == tmp10811 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__10031
}


} else {
return let__10029
}


panic("unreachable")
})

tmp10814 := Call2(__e, ns2_1set, symshen_4_5literal_6, tmp10787)


_ = tmp10814

tmp10815 := MakeNative1(func(__e *ControlFlow, V1814 Obj) Obj {
tmp10816 := Call1(__e, SymFn(symshen_4_5bterm_6), V1814)


let__10040 := tmp10816
_ = let__10040

tmp10829 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__10040)


var ifres10817 Obj

if True == tmp10829 {
tmp10818 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10817 = tmp10818


} else {
tmp10819 := Call1(__e, SymFn(symshen_4_5_1out), let__10040)


let__10041 := tmp10819
_ = let__10041

tmp10820 := Call1(__e, SymFn(symshen_4in_1_6), let__10040)


let__10042 := tmp10820
_ = let__10042

tmp10821 := Call1(__e, SymFn(symshen_4_5bterms_6), let__10042)


let__10043 := tmp10821
_ = let__10043

tmp10828 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__10043)


var ifres10822 Obj

if True == tmp10828 {
tmp10823 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10822 = tmp10823


} else {
tmp10824 := Call1(__e, SymFn(symshen_4_5_1out), let__10043)


let__10044 := tmp10824
_ = let__10044

tmp10825 := Call1(__e, SymFn(symshen_4in_1_6), let__10043)


let__10045 := tmp10825
_ = let__10045

tmp10826 := GCons(__e, let__10041, let__10044)

tmp10827 := Call2(__e, SymFn(symshen_4comb), let__10045, tmp10826)


ifres10822 = tmp10827


}

ifres10817 = ifres10822


}

let__10039 := ifres10817
_ = let__10039

tmp10839 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__10039)


if True == tmp10839 {
tmp10830 := Call1(__e, SymFn(sym_5e_6), V1814)


let__10047 := tmp10830
_ = let__10047

tmp10835 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__10047)


var ifres10831 Obj

if True == tmp10835 {
tmp10832 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10831 = tmp10832


} else {
tmp10833 := Call1(__e, SymFn(symshen_4in_1_6), let__10047)


let__10048 := tmp10833
_ = let__10048

tmp10834 := Call2(__e, SymFn(symshen_4comb), let__10048, Nil)


ifres10831 = tmp10834


}

let__10046 := ifres10831
_ = let__10046

tmp10837 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__10046)


if True == tmp10837 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__10046
}


} else {
return let__10039
}


panic("unreachable")
})

tmp10840 := Call2(__e, ns2_1set, symshen_4_5bterms_6, tmp10815)


_ = tmp10840

tmp10841 := MakeNative1(func(__e *ControlFlow, V1825 Obj) Obj {
tmp10842 := Call1(__e, SymFn(symshen_4_5wildcard_6), V1825)


let__10050 := tmp10842
_ = let__10050

tmp10848 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__10050)


var ifres10843 Obj

if True == tmp10848 {
tmp10844 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10843 = tmp10844


} else {
tmp10845 := Call1(__e, SymFn(symshen_4_5_1out), let__10050)


let__10051 := tmp10845
_ = let__10051

tmp10846 := Call1(__e, SymFn(symshen_4in_1_6), let__10050)


let__10052 := tmp10846
_ = let__10052

tmp10847 := Call2(__e, SymFn(symshen_4comb), let__10052, let__10051)


ifres10843 = tmp10847


}

let__10049 := ifres10843
_ = let__10049

tmp10880 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__10049)


if True == tmp10880 {
tmp10857 := GConsP(__e, V1825)

var ifres10849 Obj

if True == tmp10857 {
tmp10850 := Call1(__e, SymFn(symhead), V1825)


let__10054 := tmp10850
_ = let__10054

tmp10851 := Call1(__e, SymFn(symtail), V1825)


let__10055 := tmp10851
_ = let__10055

tmp10855 := Call1(__e, SymFn(symatom_2), let__10054)


var ifres10852 Obj

if True == tmp10855 {
tmp10853 := Call2(__e, SymFn(symshen_4comb), let__10055, let__10054)


ifres10852 = tmp10853


} else {
tmp10854 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10852 = tmp10854


}

ifres10849 = ifres10852


} else {
tmp10856 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10849 = tmp10856


}

let__10053 := ifres10849
_ = let__10053

tmp10878 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__10053)


if True == tmp10878 {
tmp10874 := Call1(__e, SymFn(symshen_4ccons_2), V1825)


var ifres10858 Obj

if True == tmp10874 {
tmp10859 := Call1(__e, SymFn(symhead), V1825)


let__10057 := tmp10859
_ = let__10057

tmp10860 := Call1(__e, SymFn(symtail), V1825)


let__10058 := tmp10860
_ = let__10058

tmp10861 := Call1(__e, SymFn(symshen_4_5bterms_6), let__10057)


let__10059 := tmp10861
_ = let__10059

tmp10872 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__10059)


var ifres10862 Obj

if True == tmp10872 {
tmp10863 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10862 = tmp10863


} else {
tmp10864 := Call1(__e, SymFn(symshen_4_5_1out), let__10059)


let__10060 := tmp10864
_ = let__10060

tmp10865 := Call1(__e, SymFn(symshen_4in_1_6), let__10059)


let__10061 := tmp10865
_ = let__10061

tmp10866 := Call1(__e, SymFn(sym_5end_6), let__10061)


let__10062 := tmp10866
_ = let__10062

tmp10871 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__10062)


var ifres10867 Obj

if True == tmp10871 {
tmp10868 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10867 = tmp10868


} else {
tmp10869 := Call1(__e, SymFn(symshen_4in_1_6), let__10062)


let__10063 := tmp10869
_ = let__10063

tmp10870 := Call2(__e, SymFn(symshen_4comb), let__10058, let__10060)


ifres10867 = tmp10870


}

ifres10862 = ifres10867


}

ifres10858 = ifres10862


} else {
tmp10873 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10858 = tmp10873


}

let__10056 := ifres10858
_ = let__10056

tmp10876 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__10056)


if True == tmp10876 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__10056
}


} else {
return let__10053
}


} else {
return let__10049
}


panic("unreachable")
})

tmp10881 := Call2(__e, ns2_1set, symshen_4_5bterm_6, tmp10841)


_ = tmp10881

tmp10882 := MakeNative1(func(__e *ControlFlow, V1841 Obj) Obj {
tmp10892 := GConsP(__e, V1841)

var ifres10883 Obj

if True == tmp10892 {
tmp10884 := Call1(__e, SymFn(symhead), V1841)


let__10065 := tmp10884
_ = let__10065

tmp10885 := Call1(__e, SymFn(symtail), V1841)


let__10066 := tmp10885
_ = let__10066

tmp10890 := GEq(__e, let__10065, sym__)

var ifres10886 Obj

if True == tmp10890 {
tmp10887 := Call1(__e, SymFn(symgensym), symY)


tmp10888 := Call2(__e, SymFn(symshen_4comb), let__10066, tmp10887)


ifres10886 = tmp10888


} else {
tmp10889 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10886 = tmp10889


}

ifres10883 = ifres10886


} else {
tmp10891 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10883 = tmp10891


}

let__10064 := ifres10883
_ = let__10064

tmp10894 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__10064)


if True == tmp10894 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__10064
}


panic("unreachable")
})

tmp10895 := Call2(__e, ns2_1set, symshen_4_5wildcard_6, tmp10882)


_ = tmp10895

tmp10896 := MakeNative1(func(__e *ControlFlow, V1845 Obj) Obj {
tmp10905 := GConsP(__e, V1845)

var ifres10897 Obj

if True == tmp10905 {
tmp10898 := Call1(__e, SymFn(symhead), V1845)


let__10068 := tmp10898
_ = let__10068

tmp10899 := Call1(__e, SymFn(symtail), V1845)


let__10069 := tmp10899
_ = let__10069

tmp10903 := Call1(__e, SymFn(symshen_4semicolon_2), let__10068)


var ifres10900 Obj

if True == tmp10903 {
tmp10901 := Call2(__e, SymFn(symshen_4comb), let__10069, let__10068)


ifres10900 = tmp10901


} else {
tmp10902 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10900 = tmp10902


}

ifres10897 = ifres10900


} else {
tmp10904 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres10897 = tmp10904


}

let__10067 := ifres10897
_ = let__10067

tmp10907 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__10067)


if True == tmp10907 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__10067
}


panic("unreachable")
})

tmp10908 := Call2(__e, ns2_1set, symshen_4_5sc_6, tmp10896)


_ = tmp10908

tmp10909 := MakeNative2(func(__e *ControlFlow, V1849, V1850 Obj) Obj {
tmp10910 := Call1(__e, SymFn(symgensym), symB)


let__10070 := tmp10910
_ = let__10070

tmp10911 := Call1(__e, SymFn(symgensym), symL)


let__10071 := tmp10911
_ = let__10071

tmp10912 := Call1(__e, SymFn(symgensym), symK)


let__10072 := tmp10912
_ = let__10072

tmp10913 := Call1(__e, SymFn(symgensym), symC)


let__10073 := tmp10913
_ = let__10073

tmp10914 := Call1(__e, SymFn(symshen_4prolog_1parameters), V1850)


let__10074 := tmp10914
_ = let__10074

tmp10915 := Call1(__e, SymFn(symshen_4hascut_2), V1850)


let__10075 := tmp10915
_ = let__10075

tmp10916 := Call(__e, SymFn(symshen_4prolog_1fbody), V1850, let__10074, let__10070, let__10071, let__10072, let__10073, let__10075)


let__10076 := tmp10916
_ = let__10076

var ifres10917 Obj

if True == let__10075 {
tmp10918 := GCons(__e, MakeInteger(1), Nil)

tmp10919 := GCons(__e, let__10072, tmp10918)

tmp10920 := GCons(__e, sym_7, tmp10919)

tmp10921 := GCons(__e, let__10076, Nil)

tmp10922 := GCons(__e, tmp10920, tmp10921)

tmp10923 := GCons(__e, let__10072, tmp10922)

tmp10924 := GCons(__e, symlet, tmp10923)

ifres10917 = tmp10924


} else {
ifres10917 = let__10076


}

let__10077 := ifres10917
_ = let__10077

tmp10925 := GCons(__e, sym_1_6, Nil)

tmp10926 := GCons(__e, let__10073, tmp10925)

tmp10927 := GCons(__e, let__10072, tmp10926)

tmp10928 := GCons(__e, let__10071, tmp10927)

tmp10929 := GCons(__e, let__10070, tmp10928)

tmp10930 := GCons(__e, let__10077, Nil)

tmp10931 := Call2(__e, SymFn(symappend), tmp10929, tmp10930)


tmp10932 := Call2(__e, SymFn(symappend), let__10074, tmp10931)


tmp10933 := GCons(__e, V1849, tmp10932)

tmp10934 := GCons(__e, symdefine, tmp10933)

let__10078 := tmp10934
_ = let__10078

return let__10078


panic("unreachable")
})

tmp10935 := Call2(__e, ns2_1set, symshen_4horn_1clause_1procedure, tmp10909)


_ = tmp10935

tmp10936 := MakeNative1(func(__e *ControlFlow, V1862 Obj) Obj {
tmp10946 := GEq(__e, sym_b, V1862)

if True == tmp10946 {
return True
} else {
tmp10944 := GConsP(__e, V1862)

if True == tmp10944 {
tmp10941 := GHd(__e, V1862)

tmp10942 := Call1(__e, SymFn(symshen_4hascut_2), tmp10941)


if True == tmp10942 {
return True
} else {
tmp10938 := GTl(__e, V1862)

tmp10939 := Call1(__e, SymFn(symshen_4hascut_2), tmp10938)


if True == tmp10939 {
return True
} else {
return False
}


}


} else {
return False
}


}


panic("unreachable")
})

tmp10947 := Call2(__e, ns2_1set, symshen_4hascut_2, tmp10936)


_ = tmp10947

tmp10948 := MakeNative1(func(__e *ControlFlow, V1867 Obj) Obj {
tmp10957 := GConsP(__e, V1867)

var ifres10953 Obj

if True == tmp10957 {
tmp10955 := GHd(__e, V1867)

tmp10956 := GConsP(__e, tmp10955)

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
tmp10949 := GHd(__e, V1867)

tmp10950 := GHd(__e, tmp10949)

tmp10951 := Call1(__e, SymFn(symlength), tmp10950)


return TailCall1(__e, SymFn(symshen_4parameters), tmp10951)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4prolog_1parameters)
}


panic("unreachable")
})

tmp10958 := Call2(__e, ns2_1set, symshen_4prolog_1parameters, tmp10948)


_ = tmp10958

tmp10959 := MakeNative(func(__e *ControlFlow) {
V1888 := __e.Get(1)
V1889 := __e.Get(2)
V1890 := __e.Get(3)
V1891 := __e.Get(4)
V1892 := __e.Get(5)
V1893 := __e.Get(6)
V1894 := __e.Get(7)
__e.ReturnDirect(func() Obj {
tmp11049 := GEq(__e, Nil, V1888)

var ifres11046 Obj

if True == tmp11049 {
tmp11048 := GEq(__e, True, V1894)

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
tmp10960 := GCons(__e, V1892, Nil)

tmp10961 := GCons(__e, V1891, tmp10960)

return GCons(__e, symshen_4unlock, tmp10961)


} else {
tmp11044 := GConsP(__e, V1888)

var ifres11022 Obj

if True == tmp11044 {
tmp11042 := GHd(__e, V1888)

tmp11043 := GConsP(__e, tmp11042)

var ifres11024 Obj

if True == tmp11043 {
tmp11039 := GHd(__e, V1888)

tmp11040 := GTl(__e, tmp11039)

tmp11041 := GConsP(__e, tmp11040)

var ifres11026 Obj

if True == tmp11041 {
tmp11035 := GHd(__e, V1888)

tmp11036 := GTl(__e, tmp11035)

tmp11037 := GTl(__e, tmp11036)

tmp11038 := GEq(__e, Nil, tmp11037)

var ifres11028 Obj

if True == tmp11038 {
tmp11033 := GTl(__e, V1888)

tmp11034 := GEq(__e, Nil, tmp11033)

var ifres11030 Obj

if True == tmp11034 {
tmp11032 := GEq(__e, False, V1894)

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
tmp10962 := GHd(__e, V1888)

tmp10963 := GHd(__e, tmp10962)

tmp10964 := GHd(__e, V1888)

tmp10965 := GTl(__e, tmp10964)

tmp10966 := GHd(__e, tmp10965)

tmp10967 := Call6(__e, SymFn(symshen_4continue), tmp10963, tmp10966, V1890, V1891, V1892, V1893)


let__10079 := tmp10967
_ = let__10079

tmp10968 := GCons(__e, V1891, Nil)

tmp10969 := GCons(__e, symshen_4unlocked_2, tmp10968)

tmp10970 := GHd(__e, V1888)

tmp10971 := GHd(__e, tmp10970)

tmp10972 := Call5(__e, SymFn(symshen_4compile_1head), symshen_4_7m, tmp10971, V1889, V1890, let__10079)


tmp10973 := GCons(__e, False, Nil)

tmp10974 := GCons(__e, tmp10972, tmp10973)

tmp10975 := GCons(__e, tmp10969, tmp10974)

return GCons(__e, symif, tmp10975)


} else {
tmp11020 := GConsP(__e, V1888)

var ifres11005 Obj

if True == tmp11020 {
tmp11018 := GHd(__e, V1888)

tmp11019 := GConsP(__e, tmp11018)

var ifres11007 Obj

if True == tmp11019 {
tmp11015 := GHd(__e, V1888)

tmp11016 := GTl(__e, tmp11015)

tmp11017 := GConsP(__e, tmp11016)

var ifres11009 Obj

if True == tmp11017 {
tmp11011 := GHd(__e, V1888)

tmp11012 := GTl(__e, tmp11011)

tmp11013 := GTl(__e, tmp11012)

tmp11014 := GEq(__e, Nil, tmp11013)

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
tmp10976 := Call1(__e, SymFn(symgensym), symC)


let__10080 := tmp10976
_ = let__10080

tmp10977 := GHd(__e, V1888)

tmp10978 := GHd(__e, tmp10977)

tmp10979 := GHd(__e, V1888)

tmp10980 := GTl(__e, tmp10979)

tmp10981 := GHd(__e, tmp10980)

tmp10982 := Call6(__e, SymFn(symshen_4continue), tmp10978, tmp10981, V1890, V1891, V1892, V1893)


let__10081 := tmp10982
_ = let__10081

tmp10983 := GCons(__e, V1891, Nil)

tmp10984 := GCons(__e, symshen_4unlocked_2, tmp10983)

tmp10985 := GHd(__e, V1888)

tmp10986 := GHd(__e, tmp10985)

tmp10987 := Call5(__e, SymFn(symshen_4compile_1head), symshen_4_7m, tmp10986, V1889, V1890, let__10081)


tmp10988 := GCons(__e, False, Nil)

tmp10989 := GCons(__e, tmp10987, tmp10988)

tmp10990 := GCons(__e, tmp10984, tmp10989)

tmp10991 := GCons(__e, symif, tmp10990)

tmp10992 := GCons(__e, False, Nil)

tmp10993 := GCons(__e, let__10080, tmp10992)

tmp10994 := GCons(__e, sym_a, tmp10993)

tmp10995 := GTl(__e, V1888)

tmp10996 := Call(__e, SymFn(symshen_4prolog_1fbody), tmp10995, V1889, V1890, V1891, V1892, V1893, V1894)


tmp10997 := GCons(__e, let__10080, Nil)

tmp10998 := GCons(__e, tmp10996, tmp10997)

tmp10999 := GCons(__e, tmp10994, tmp10998)

tmp11000 := GCons(__e, symif, tmp10999)

tmp11001 := GCons(__e, tmp11000, Nil)

tmp11002 := GCons(__e, tmp10991, tmp11001)

tmp11003 := GCons(__e, let__10080, tmp11002)

return GCons(__e, symlet, tmp11003)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.prolog-fbody"))
}


}


}


panic("unreachable")
}())
}, 7)

tmp11050 := Call2(__e, ns2_1set, symshen_4prolog_1fbody, tmp10959)


_ = tmp11050

tmp11051 := MakeNative2(func(__e *ControlFlow, V1898, V1899 Obj) Obj {
tmp11056 := Call1(__e, SymFn(symshen_4locked_2), V1898)


var ifres11053 Obj

if True == tmp11056 {
tmp11055 := Call2(__e, SymFn(symshen_4fits_2), V1899, V1898)


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
return TailCall1(__e, SymFn(symshen_4openlock), V1898)
} else {
return False
}


panic("unreachable")
})

tmp11057 := Call2(__e, ns2_1set, symshen_4unlock, tmp11051)


_ = tmp11057

tmp11058 := MakeNative1(func(__e *ControlFlow, V1900 Obj) Obj {
tmp11059 := Call1(__e, SymFn(symshen_4unlocked_2), V1900)


return GNot(__e, tmp11059)


panic("unreachable")
})

tmp11060 := Call2(__e, ns2_1set, symshen_4locked_2, tmp11058)


_ = tmp11060

tmp11061 := MakeNative1(func(__e *ControlFlow, V1901 Obj) Obj {
return GVecGet(__e, V1901, MakeInteger(1))
panic("unreachable")
})

tmp11062 := Call2(__e, ns2_1set, symshen_4unlocked_2, tmp11061)


_ = tmp11062

tmp11063 := MakeNative1(func(__e *ControlFlow, V1902 Obj) Obj {
tmp11064 := GVecSet(__e, V1902, MakeInteger(1), True)

_ = tmp11064

return False


panic("unreachable")
})

tmp11065 := Call2(__e, ns2_1set, symshen_4openlock, tmp11063)


_ = tmp11065

tmp11066 := MakeNative2(func(__e *ControlFlow, V1903, V1904 Obj) Obj {
tmp11067 := GVecGet(__e, V1904, MakeInteger(2))

return GEq(__e, V1903, tmp11067)


panic("unreachable")
})

tmp11068 := Call2(__e, ns2_1set, symshen_4fits_2, tmp11066)


_ = tmp11068

tmp11069 := MakeNative4(func(__e *ControlFlow, V1907, V1908, V1909, V1910 Obj) Obj {
tmp11070 := Call1(__e, SymFn(symthaw), V1910)


let__10082 := tmp11070
_ = let__10082

tmp11075 := GEq(__e, let__10082, False)

var ifres11072 Obj

if True == tmp11075 {
tmp11074 := Call1(__e, SymFn(symshen_4unlocked_2), V1908)


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
return TailCall2(__e, SymFn(symshen_4lock), V1909, V1908)
} else {
return let__10082
}


panic("unreachable")
})

tmp11076 := Call2(__e, ns2_1set, symshen_4cut, tmp11069)


_ = tmp11076

tmp11077 := MakeNative2(func(__e *ControlFlow, V1912, V1913 Obj) Obj {
tmp11078 := GVecSet(__e, V1913, MakeInteger(1), False)

let__10083 := tmp11078
_ = let__10083

tmp11079 := GVecSet(__e, V1913, MakeInteger(2), V1912)

let__10084 := tmp11079
_ = let__10084

return False


panic("unreachable")
})

tmp11080 := Call2(__e, ns2_1set, symshen_4lock, tmp11077)


_ = tmp11080

tmp11081 := MakeNative6(func(__e *ControlFlow, V1916, V1917, V1918, V1919, V1920, V1921 Obj) Obj {
tmp11082 := Call1(__e, SymFn(symshen_4extract_1vars), V1916)


let__10085 := tmp11082
_ = let__10085

tmp11083 := Call1(__e, SymFn(symshen_4extract_1free_1vars), V1917)


let__10086 := tmp11083
_ = let__10086

tmp11084 := Call2(__e, SymFn(symdifference), let__10086, let__10085)


let__10087 := tmp11084
_ = let__10087

tmp11085 := GCons(__e, symshen_4incinfs, Nil)

tmp11086 := Call5(__e, SymFn(symshen_4compile_1body), V1917, V1918, V1919, V1920, V1921)


tmp11087 := GCons(__e, tmp11086, Nil)

tmp11088 := GCons(__e, tmp11085, tmp11087)

tmp11089 := GCons(__e, symdo, tmp11088)

let__10088 := tmp11089
_ = let__10088

return TailCall3(__e, SymFn(symshen_4stpart), let__10087, let__10088, V1918)


panic("unreachable")
})

tmp11090 := Call2(__e, ns2_1set, symshen_4continue, tmp11081)


_ = tmp11090

tmp11091 := MakeNative1(func(__e *ControlFlow, V1928 Obj) Obj {
tmp11126 := GConsP(__e, V1928)

var ifres11107 Obj

if True == tmp11126 {
tmp11124 := GHd(__e, V1928)

tmp11125 := GEq(__e, symlambda, tmp11124)

var ifres11109 Obj

if True == tmp11125 {
tmp11122 := GTl(__e, V1928)

tmp11123 := GConsP(__e, tmp11122)

var ifres11111 Obj

if True == tmp11123 {
tmp11119 := GTl(__e, V1928)

tmp11120 := GTl(__e, tmp11119)

tmp11121 := GConsP(__e, tmp11120)

var ifres11113 Obj

if True == tmp11121 {
tmp11115 := GTl(__e, V1928)

tmp11116 := GTl(__e, tmp11115)

tmp11117 := GTl(__e, tmp11116)

tmp11118 := GEq(__e, Nil, tmp11117)

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
tmp11092 := GTl(__e, V1928)

tmp11093 := GHd(__e, tmp11092)

tmp11094 := GTl(__e, V1928)

tmp11095 := GTl(__e, tmp11094)

tmp11096 := GHd(__e, tmp11095)

tmp11097 := Call1(__e, SymFn(symshen_4extract_1free_1vars), tmp11096)


return TailCall2(__e, SymFn(symremove), tmp11093, tmp11097)


} else {
tmp11105 := GConsP(__e, V1928)

if True == tmp11105 {
tmp11098 := GHd(__e, V1928)

tmp11099 := Call1(__e, SymFn(symshen_4extract_1free_1vars), tmp11098)


tmp11100 := GTl(__e, V1928)

tmp11101 := Call1(__e, SymFn(symshen_4extract_1free_1vars), tmp11100)


return TailCall2(__e, SymFn(symunion), tmp11099, tmp11101)


} else {
tmp11103 := PrimGuard1(__e, symvariable_2, PrimIsVariable, V1928)

if True == tmp11103 {
return GCons(__e, V1928, Nil)
} else {
return Nil
}


}


}


panic("unreachable")
})

tmp11127 := Call2(__e, ns2_1set, symshen_4extract_1free_1vars, tmp11091)


_ = tmp11127

tmp11128 := (func() Obj {
var __self Obj
__self = MakeNative5(func(__e *ControlFlow, __self1, __self2, __self3, __self4, __self5 Obj) Obj {
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
tmp11162 := GEq(__e, Nil, V1945)

if True == tmp11162 {
tmp11129 := GCons(__e, V1949, Nil)

return GCons(__e, symthaw, tmp11129)


} else {
tmp11160 := GConsP(__e, V1945)

var ifres11156 Obj

if True == tmp11160 {
tmp11158 := GHd(__e, V1945)

tmp11159 := GEq(__e, sym_b, tmp11158)

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
tmp11130 := GCons(__e, symshen_4cut, Nil)

tmp11131 := GTl(__e, V1945)

tmp11132 := GCons(__e, tmp11130, tmp11131)

if SymFn(symshen_4compile_1body) == __self {
__self1, __self2, __self3, __self4, __self5 = tmp11132, V1946, V1947, V1948, V1949
__e.Tick()
goto __selftop
}
return TailCall5(__e, SymFn(symshen_4compile_1body), tmp11132, V1946, V1947, V1948, V1949)


} else {
tmp11154 := GConsP(__e, V1945)

var ifres11150 Obj

if True == tmp11154 {
tmp11152 := GTl(__e, V1945)

tmp11153 := GEq(__e, Nil, tmp11152)

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
tmp11133 := GHd(__e, V1945)

tmp11134 := Call2(__e, SymFn(symshen_4deref_1calls), tmp11133, V1946)


tmp11135 := GCons(__e, V1949, Nil)

tmp11136 := GCons(__e, V1948, tmp11135)

tmp11137 := GCons(__e, V1947, tmp11136)

tmp11138 := GCons(__e, V1946, tmp11137)

return TailCall2(__e, SymFn(symappend), tmp11134, tmp11138)


} else {
tmp11148 := GConsP(__e, V1945)

if True == tmp11148 {
tmp11139 := GHd(__e, V1945)

tmp11140 := Call2(__e, SymFn(symshen_4deref_1calls), tmp11139, V1946)


let__10089 := tmp11140
_ = let__10089

tmp11141 := GTl(__e, V1945)

tmp11142 := Call5(__e, SymFn(symshen_4freeze_1literals), tmp11141, V1946, V1947, V1948, V1949)


tmp11143 := GCons(__e, tmp11142, Nil)

tmp11144 := GCons(__e, V1948, tmp11143)

tmp11145 := GCons(__e, V1947, tmp11144)

tmp11146 := GCons(__e, V1946, tmp11145)

return TailCall2(__e, SymFn(symappend), let__10089, tmp11146)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.compile-fbody"))
}


}


}


}


panic("unreachable")
})
return __self
})()

tmp11163 := Call2(__e, ns2_1set, symshen_4compile_1body, tmp11128)


_ = tmp11163

tmp11164 := (func() Obj {
var __self Obj
__self = MakeNative5(func(__e *ControlFlow, __self1, __self2, __self3, __self4, __self5 Obj) Obj {
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
tmp11187 := GEq(__e, Nil, V1967)

if True == tmp11187 {
return V1971
} else {
tmp11185 := GConsP(__e, V1967)

var ifres11181 Obj

if True == tmp11185 {
tmp11183 := GHd(__e, V1967)

tmp11184 := GEq(__e, sym_b, tmp11183)

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
tmp11165 := GCons(__e, symshen_4cut, Nil)

tmp11166 := GTl(__e, V1967)

tmp11167 := GCons(__e, tmp11165, tmp11166)

if SymFn(symshen_4freeze_1literals) == __self {
__self1, __self2, __self3, __self4, __self5 = tmp11167, V1968, V1969, V1970, V1971
__e.Tick()
goto __selftop
}
return TailCall5(__e, SymFn(symshen_4freeze_1literals), tmp11167, V1968, V1969, V1970, V1971)


} else {
tmp11179 := GConsP(__e, V1967)

if True == tmp11179 {
tmp11168 := GHd(__e, V1967)

tmp11169 := Call2(__e, SymFn(symshen_4deref_1calls), tmp11168, V1968)


let__10090 := tmp11169
_ = let__10090

tmp11170 := GTl(__e, V1967)

tmp11171 := Call5(__e, SymFn(symshen_4freeze_1literals), tmp11170, V1968, V1969, V1970, V1971)


tmp11172 := GCons(__e, tmp11171, Nil)

tmp11173 := GCons(__e, V1970, tmp11172)

tmp11174 := GCons(__e, V1969, tmp11173)

tmp11175 := GCons(__e, V1968, tmp11174)

tmp11176 := Call2(__e, SymFn(symappend), let__10090, tmp11175)


tmp11177 := GCons(__e, tmp11176, Nil)

return GCons(__e, symfreeze, tmp11177)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.freeze-literals"))
}


}


}


panic("unreachable")
})
return __self
})()

tmp11188 := Call2(__e, ns2_1set, symshen_4freeze_1literals, tmp11164)


_ = tmp11188

tmp11189 := MakeNative2(func(__e *ControlFlow, V1977, V1978 Obj) Obj {
tmp11204 := GConsP(__e, V1977)

var ifres11200 Obj

if True == tmp11204 {
tmp11202 := GHd(__e, V1977)

tmp11203 := GEq(__e, symfork, tmp11202)

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
tmp11190 := GTl(__e, V1977)

tmp11191 := Call2(__e, SymFn(symshen_4deref_1forked_1literals), tmp11190, V1978)


tmp11192 := GCons(__e, tmp11191, Nil)

return GCons(__e, symfork, tmp11192)


} else {
tmp11198 := GConsP(__e, V1977)

if True == tmp11198 {
tmp11193 := GHd(__e, V1977)

tmp11194 := MakeNative1(func(__e *ControlFlow, Z1979 Obj) Obj {
return TailCall2(__e, SymFn(symshen_4function_1calls), Z1979, V1978)
panic("unreachable")
})

tmp11195 := GTl(__e, V1977)

tmp11196 := Call2(__e, SymFn(symmap), tmp11194, tmp11195)


return GCons(__e, tmp11193, tmp11196)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.deref-calls"))
}


}


panic("unreachable")
})

tmp11205 := Call2(__e, ns2_1set, symshen_4deref_1calls, tmp11189)


_ = tmp11205

tmp11206 := MakeNative2(func(__e *ControlFlow, V1986, V1987 Obj) Obj {
tmp11216 := GEq(__e, Nil, V1986)

if True == tmp11216 {
return Nil
} else {
tmp11214 := GConsP(__e, V1986)

if True == tmp11214 {
tmp11207 := GHd(__e, V1986)

tmp11208 := Call2(__e, SymFn(symshen_4deref_1calls), tmp11207, V1987)


tmp11209 := GTl(__e, V1986)

tmp11210 := Call2(__e, SymFn(symshen_4deref_1forked_1literals), tmp11209, V1987)


tmp11211 := GCons(__e, tmp11210, Nil)

tmp11212 := GCons(__e, tmp11208, tmp11211)

return GCons(__e, symcons, tmp11212)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("fork requires a list of literals\n"))
}


}


panic("unreachable")
})

tmp11217 := Call2(__e, ns2_1set, symshen_4deref_1forked_1literals, tmp11206)


_ = tmp11217

tmp11218 := MakeNative2(func(__e *ControlFlow, V1990, V1991 Obj) Obj {
tmp11250 := GConsP(__e, V1990)

var ifres11231 Obj

if True == tmp11250 {
tmp11248 := GHd(__e, V1990)

tmp11249 := GEq(__e, symcons, tmp11248)

var ifres11233 Obj

if True == tmp11249 {
tmp11246 := GTl(__e, V1990)

tmp11247 := GConsP(__e, tmp11246)

var ifres11235 Obj

if True == tmp11247 {
tmp11243 := GTl(__e, V1990)

tmp11244 := GTl(__e, tmp11243)

tmp11245 := GConsP(__e, tmp11244)

var ifres11237 Obj

if True == tmp11245 {
tmp11239 := GTl(__e, V1990)

tmp11240 := GTl(__e, tmp11239)

tmp11241 := GTl(__e, tmp11240)

tmp11242 := GEq(__e, Nil, tmp11241)

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
tmp11219 := GTl(__e, V1990)

tmp11220 := GHd(__e, tmp11219)

tmp11221 := Call2(__e, SymFn(symshen_4function_1calls), tmp11220, V1991)


tmp11222 := GTl(__e, V1990)

tmp11223 := GTl(__e, tmp11222)

tmp11224 := GHd(__e, tmp11223)

tmp11225 := Call2(__e, SymFn(symshen_4function_1calls), tmp11224, V1991)


tmp11226 := GCons(__e, tmp11225, Nil)

tmp11227 := GCons(__e, tmp11221, tmp11226)

return GCons(__e, symcons, tmp11227)


} else {
tmp11229 := GConsP(__e, V1990)

if True == tmp11229 {
return TailCall3(__e, SymFn(symshen_4deref_1terms), V1990, V1991, Nil)
} else {
return V1990
}


}


panic("unreachable")
})

tmp11251 := Call2(__e, ns2_1set, symshen_4function_1calls, tmp11218)


_ = tmp11251

tmp11252 := MakeNative3(func(__e *ControlFlow, V2000, V2001, V2002 Obj) Obj {
tmp11346 := GConsP(__e, V2000)

var ifres11333 Obj

if True == tmp11346 {
tmp11344 := GHd(__e, V2000)

tmp11345 := GEq(__e, MakeInteger(0), tmp11344)

var ifres11335 Obj

if True == tmp11345 {
tmp11342 := GTl(__e, V2000)

tmp11343 := GConsP(__e, tmp11342)

var ifres11337 Obj

if True == tmp11343 {
tmp11339 := GTl(__e, V2000)

tmp11340 := GTl(__e, tmp11339)

tmp11341 := GEq(__e, Nil, tmp11340)

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
tmp11259 := GTl(__e, V2000)

tmp11260 := GHd(__e, tmp11259)

tmp11261 := PrimGuard1(__e, symvariable_2, PrimIsVariable, tmp11260)

if True == tmp11261 {
tmp11253 := GTl(__e, V2000)

return GHd(__e, tmp11253)


} else {
tmp11254 := GTl(__e, V2000)

tmp11255 := GHd(__e, tmp11254)

tmp11256 := Call3(__e, SymFn(symshen_4app), tmp11255, MakeString("\n"), symshen_4s)


return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("attempt to optimise a non-variable "), tmp11256))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("attempt to optimise a non-variable "), tmp11256)
return Call1(__e, SymFn(symsimple_1error), __typedArg0)
})()


}


} else {
tmp11331 := GConsP(__e, V2000)

var ifres11318 Obj

if True == tmp11331 {
tmp11329 := GHd(__e, V2000)

tmp11330 := GEq(__e, MakeInteger(1), tmp11329)

var ifres11320 Obj

if True == tmp11330 {
tmp11327 := GTl(__e, V2000)

tmp11328 := GConsP(__e, tmp11327)

var ifres11322 Obj

if True == tmp11328 {
tmp11324 := GTl(__e, V2000)

tmp11325 := GTl(__e, tmp11324)

tmp11326 := GEq(__e, Nil, tmp11325)

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
tmp11271 := GTl(__e, V2000)

tmp11272 := GHd(__e, tmp11271)

tmp11273 := PrimGuard1(__e, symvariable_2, PrimIsVariable, tmp11272)

if True == tmp11273 {
tmp11262 := GTl(__e, V2000)

tmp11263 := GHd(__e, tmp11262)

tmp11264 := GCons(__e, V2001, Nil)

tmp11265 := GCons(__e, tmp11263, tmp11264)

return GCons(__e, symshen_4lazyderef, tmp11265)


} else {
tmp11266 := GTl(__e, V2000)

tmp11267 := GHd(__e, tmp11266)

tmp11268 := Call3(__e, SymFn(symshen_4app), tmp11267, MakeString("\n"), symshen_4s)


return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("attempt to optimise a non-variable "), tmp11268))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("attempt to optimise a non-variable "), tmp11268)
return Call1(__e, SymFn(symsimple_1error), __typedArg0)
})()


}


} else {
tmp11315 := Call2(__e, SymFn(symelement_2), V2000, V2002)


tmp11316 := GNot(__e, tmp11315)

var ifres11312 Obj

if True == tmp11316 {
tmp11314 := PrimGuard1(__e, symvariable_2, PrimIsVariable, V2000)

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
tmp11274 := GCons(__e, V2001, Nil)

tmp11275 := GCons(__e, V2000, tmp11274)

return GCons(__e, symshen_4deref, tmp11275)


} else {
tmp11310 := GConsP(__e, V2000)

var ifres11291 Obj

if True == tmp11310 {
tmp11308 := GHd(__e, V2000)

tmp11309 := GEq(__e, symlambda, tmp11308)

var ifres11293 Obj

if True == tmp11309 {
tmp11306 := GTl(__e, V2000)

tmp11307 := GConsP(__e, tmp11306)

var ifres11295 Obj

if True == tmp11307 {
tmp11303 := GTl(__e, V2000)

tmp11304 := GTl(__e, tmp11303)

tmp11305 := GConsP(__e, tmp11304)

var ifres11297 Obj

if True == tmp11305 {
tmp11299 := GTl(__e, V2000)

tmp11300 := GTl(__e, tmp11299)

tmp11301 := GTl(__e, tmp11300)

tmp11302 := GEq(__e, Nil, tmp11301)

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
tmp11276 := GTl(__e, V2000)

tmp11277 := GHd(__e, tmp11276)

tmp11278 := GTl(__e, V2000)

tmp11279 := GTl(__e, tmp11278)

tmp11280 := GHd(__e, tmp11279)

tmp11281 := GTl(__e, V2000)

tmp11282 := GHd(__e, tmp11281)

tmp11283 := GCons(__e, tmp11282, V2002)

tmp11284 := Call3(__e, SymFn(symshen_4deref_1terms), tmp11280, V2001, tmp11283)


tmp11285 := GCons(__e, tmp11284, Nil)

tmp11286 := GCons(__e, tmp11277, tmp11285)

return GCons(__e, symlambda, tmp11286)


} else {
tmp11289 := GConsP(__e, V2000)

if True == tmp11289 {
tmp11287 := MakeNative1(func(__e *ControlFlow, Z2003 Obj) Obj {
return TailCall3(__e, SymFn(symshen_4deref_1terms), Z2003, V2001, V2002)
panic("unreachable")
})

return TailCall2(__e, SymFn(symmap), tmp11287, V2000)


} else {
return V2000
}


}


}


}


}


panic("unreachable")
})

tmp11347 := Call2(__e, ns2_1set, symshen_4deref_1terms, tmp11252)


_ = tmp11347

tmp11348 := (func() Obj {
var __self Obj
__self = MakeNative5(func(__e *ControlFlow, __self1, __self2, __self3, __self4, __self5 Obj) Obj {
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
tmp11524 := GEq(__e, Nil, V2022)

var ifres11521 Obj

if True == tmp11524 {
tmp11523 := GEq(__e, Nil, V2023)

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
return V2025
} else {
tmp11519 := GConsP(__e, V2022)

var ifres11499 Obj

if True == tmp11519 {
tmp11517 := GHd(__e, V2022)

tmp11518 := GConsP(__e, tmp11517)

var ifres11501 Obj

if True == tmp11518 {
tmp11514 := GHd(__e, V2022)

tmp11515 := GHd(__e, tmp11514)

tmp11516 := GEq(__e, symshen_4_7m, tmp11515)

var ifres11503 Obj

if True == tmp11516 {
tmp11511 := GHd(__e, V2022)

tmp11512 := GTl(__e, tmp11511)

tmp11513 := GConsP(__e, tmp11512)

var ifres11505 Obj

if True == tmp11513 {
tmp11507 := GHd(__e, V2022)

tmp11508 := GTl(__e, tmp11507)

tmp11509 := GTl(__e, tmp11508)

tmp11510 := GEq(__e, Nil, tmp11509)

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
tmp11349 := GHd(__e, V2022)

tmp11350 := GTl(__e, tmp11349)

tmp11351 := GHd(__e, tmp11350)

tmp11352 := GTl(__e, V2022)

tmp11353 := GCons(__e, V2021, tmp11352)

tmp11354 := GCons(__e, tmp11351, tmp11353)

tmp11355 := GCons(__e, symshen_4_7m, tmp11354)

if SymFn(symshen_4compile_1head) == __self {
__self1, __self2, __self3, __self4, __self5 = V2021, tmp11355, V2023, V2024, V2025
__e.Tick()
goto __selftop
}
return TailCall5(__e, SymFn(symshen_4compile_1head), V2021, tmp11355, V2023, V2024, V2025)


} else {
tmp11497 := GConsP(__e, V2022)

var ifres11477 Obj

if True == tmp11497 {
tmp11495 := GHd(__e, V2022)

tmp11496 := GConsP(__e, tmp11495)

var ifres11479 Obj

if True == tmp11496 {
tmp11492 := GHd(__e, V2022)

tmp11493 := GHd(__e, tmp11492)

tmp11494 := GEq(__e, symshen_4_1m, tmp11493)

var ifres11481 Obj

if True == tmp11494 {
tmp11489 := GHd(__e, V2022)

tmp11490 := GTl(__e, tmp11489)

tmp11491 := GConsP(__e, tmp11490)

var ifres11483 Obj

if True == tmp11491 {
tmp11485 := GHd(__e, V2022)

tmp11486 := GTl(__e, tmp11485)

tmp11487 := GTl(__e, tmp11486)

tmp11488 := GEq(__e, Nil, tmp11487)

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
tmp11356 := GHd(__e, V2022)

tmp11357 := GTl(__e, tmp11356)

tmp11358 := GHd(__e, tmp11357)

tmp11359 := GTl(__e, V2022)

tmp11360 := GCons(__e, V2021, tmp11359)

tmp11361 := GCons(__e, tmp11358, tmp11360)

tmp11362 := GCons(__e, symshen_4_1m, tmp11361)

if SymFn(symshen_4compile_1head) == __self {
__self1, __self2, __self3, __self4, __self5 = V2021, tmp11362, V2023, V2024, V2025
__e.Tick()
goto __selftop
}
return TailCall5(__e, SymFn(symshen_4compile_1head), V2021, tmp11362, V2023, V2024, V2025)


} else {
tmp11475 := GConsP(__e, V2022)

var ifres11471 Obj

if True == tmp11475 {
tmp11473 := GHd(__e, V2022)

tmp11474 := GEq(__e, symshen_4_1m, tmp11473)

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
tmp11363 := GTl(__e, V2022)

if SymFn(symshen_4compile_1head) == __self {
__self1, __self2, __self3, __self4, __self5 = symshen_4_1m, tmp11363, V2023, V2024, V2025
__e.Tick()
goto __selftop
}
return TailCall5(__e, SymFn(symshen_4compile_1head), symshen_4_1m, tmp11363, V2023, V2024, V2025)


} else {
tmp11469 := GConsP(__e, V2022)

var ifres11465 Obj

if True == tmp11469 {
tmp11467 := GHd(__e, V2022)

tmp11468 := GEq(__e, symshen_4_7m, tmp11467)

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
tmp11364 := GTl(__e, V2022)

if SymFn(symshen_4compile_1head) == __self {
__self1, __self2, __self3, __self4, __self5 = symshen_4_7m, tmp11364, V2023, V2024, V2025
__e.Tick()
goto __selftop
}
return TailCall5(__e, SymFn(symshen_4compile_1head), symshen_4_7m, tmp11364, V2023, V2024, V2025)


} else {
tmp11463 := GConsP(__e, V2022)

var ifres11456 Obj

if True == tmp11463 {
tmp11462 := GConsP(__e, V2023)

var ifres11458 Obj

if True == tmp11462 {
tmp11460 := GHd(__e, V2022)

tmp11461 := Call1(__e, SymFn(symshen_4wildcard_2), tmp11460)


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
tmp11365 := GTl(__e, V2022)

tmp11366 := GTl(__e, V2023)

if SymFn(symshen_4compile_1head) == __self {
__self1, __self2, __self3, __self4, __self5 = V2021, tmp11365, tmp11366, V2024, V2025
__e.Tick()
goto __selftop
}
return TailCall5(__e, SymFn(symshen_4compile_1head), V2021, tmp11365, tmp11366, V2024, V2025)


} else {
tmp11454 := GConsP(__e, V2022)

var ifres11450 Obj

if True == tmp11454 {
tmp11452 := GHd(__e, V2022)

tmp11453 := PrimGuard1(__e, symvariable_2, PrimIsVariable, tmp11452)

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
return TailCall5(__e, SymFn(symshen_4variable_1case), V2021, V2022, V2023, V2024, V2025)
} else {
tmp11448 := GEq(__e, symshen_4_1m, V2021)

var ifres11441 Obj

if True == tmp11448 {
tmp11447 := GConsP(__e, V2022)

var ifres11443 Obj

if True == tmp11447 {
tmp11445 := GHd(__e, V2022)

tmp11446 := Call1(__e, SymFn(symatom_2), tmp11445)


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
return TailCall4(__e, SymFn(symshen_4atom_1case_1minus), V2022, V2023, V2024, V2025)
} else {
tmp11439 := GEq(__e, symshen_4_1m, V2021)

var ifres11409 Obj

if True == tmp11439 {
tmp11438 := GConsP(__e, V2022)

var ifres11411 Obj

if True == tmp11438 {
tmp11436 := GHd(__e, V2022)

tmp11437 := GConsP(__e, tmp11436)

var ifres11413 Obj

if True == tmp11437 {
tmp11433 := GHd(__e, V2022)

tmp11434 := GHd(__e, tmp11433)

tmp11435 := GEq(__e, symcons, tmp11434)

var ifres11415 Obj

if True == tmp11435 {
tmp11430 := GHd(__e, V2022)

tmp11431 := GTl(__e, tmp11430)

tmp11432 := GConsP(__e, tmp11431)

var ifres11417 Obj

if True == tmp11432 {
tmp11426 := GHd(__e, V2022)

tmp11427 := GTl(__e, tmp11426)

tmp11428 := GTl(__e, tmp11427)

tmp11429 := GConsP(__e, tmp11428)

var ifres11419 Obj

if True == tmp11429 {
tmp11421 := GHd(__e, V2022)

tmp11422 := GTl(__e, tmp11421)

tmp11423 := GTl(__e, tmp11422)

tmp11424 := GTl(__e, tmp11423)

tmp11425 := GEq(__e, Nil, tmp11424)

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
return TailCall4(__e, SymFn(symshen_4cons_1case_1minus), V2022, V2023, V2024, V2025)
} else {
tmp11407 := GEq(__e, symshen_4_7m, V2021)

var ifres11400 Obj

if True == tmp11407 {
tmp11406 := GConsP(__e, V2022)

var ifres11402 Obj

if True == tmp11406 {
tmp11404 := GHd(__e, V2022)

tmp11405 := Call1(__e, SymFn(symatom_2), tmp11404)


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
return TailCall4(__e, SymFn(symshen_4atom_1case_1plus), V2022, V2023, V2024, V2025)
} else {
tmp11398 := GEq(__e, symshen_4_7m, V2021)

var ifres11368 Obj

if True == tmp11398 {
tmp11397 := GConsP(__e, V2022)

var ifres11370 Obj

if True == tmp11397 {
tmp11395 := GHd(__e, V2022)

tmp11396 := GConsP(__e, tmp11395)

var ifres11372 Obj

if True == tmp11396 {
tmp11392 := GHd(__e, V2022)

tmp11393 := GHd(__e, tmp11392)

tmp11394 := GEq(__e, symcons, tmp11393)

var ifres11374 Obj

if True == tmp11394 {
tmp11389 := GHd(__e, V2022)

tmp11390 := GTl(__e, tmp11389)

tmp11391 := GConsP(__e, tmp11390)

var ifres11376 Obj

if True == tmp11391 {
tmp11385 := GHd(__e, V2022)

tmp11386 := GTl(__e, tmp11385)

tmp11387 := GTl(__e, tmp11386)

tmp11388 := GConsP(__e, tmp11387)

var ifres11378 Obj

if True == tmp11388 {
tmp11380 := GHd(__e, V2022)

tmp11381 := GTl(__e, tmp11380)

tmp11382 := GTl(__e, tmp11381)

tmp11383 := GTl(__e, tmp11382)

tmp11384 := GEq(__e, Nil, tmp11383)

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
return TailCall4(__e, SymFn(symshen_4cons_1case_1plus), V2022, V2023, V2024, V2025)
} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.compile-head"))
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


panic("unreachable")
})
return __self
})()

tmp11525 := Call2(__e, ns2_1set, symshen_4compile_1head, tmp11348)


_ = tmp11525

tmp11526 := MakeNative5(func(__e *ControlFlow, V2036, V2037, V2038, V2039, V2040 Obj) Obj {
tmp11547 := GConsP(__e, V2037)

var ifres11544 Obj

if True == tmp11547 {
tmp11546 := GConsP(__e, V2038)

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
tmp11541 := GHd(__e, V2038)

tmp11542 := PrimGuard1(__e, symvariable_2, PrimIsVariable, tmp11541)

if True == tmp11542 {
tmp11527 := GTl(__e, V2037)

tmp11528 := GTl(__e, V2038)

tmp11529 := GHd(__e, V2038)

tmp11530 := GHd(__e, V2037)

tmp11531 := Call3(__e, SymFn(symsubst), tmp11529, tmp11530, V2040)


return TailCall5(__e, SymFn(symshen_4compile_1head), V2036, tmp11527, tmp11528, V2039, tmp11531)


} else {
tmp11532 := GHd(__e, V2037)

tmp11533 := GHd(__e, V2038)

tmp11534 := GTl(__e, V2037)

tmp11535 := GTl(__e, V2038)

tmp11536 := Call5(__e, SymFn(symshen_4compile_1head), V2036, tmp11534, tmp11535, V2039, V2040)


tmp11537 := GCons(__e, tmp11536, Nil)

tmp11538 := GCons(__e, tmp11533, tmp11537)

tmp11539 := GCons(__e, tmp11532, tmp11538)

return GCons(__e, symlet, tmp11539)


}


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.variable-case"))
}


panic("unreachable")
})

tmp11548 := Call2(__e, ns2_1set, symshen_4variable_1case, tmp11526)


_ = tmp11548

tmp11549 := MakeNative4(func(__e *ControlFlow, V2049, V2050, V2051, V2052 Obj) Obj {
tmp11573 := GConsP(__e, V2049)

var ifres11570 Obj

if True == tmp11573 {
tmp11572 := GConsP(__e, V2050)

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
tmp11550 := Call1(__e, SymFn(symgensym), symTm)


let__10091 := tmp11550
_ = let__10091

tmp11551 := GHd(__e, V2050)

tmp11552 := GCons(__e, V2051, Nil)

tmp11553 := GCons(__e, tmp11551, tmp11552)

tmp11554 := GCons(__e, symshen_4lazyderef, tmp11553)

tmp11555 := GHd(__e, V2049)

tmp11556 := GCons(__e, tmp11555, Nil)

tmp11557 := GCons(__e, let__10091, tmp11556)

tmp11558 := GCons(__e, sym_a, tmp11557)

tmp11559 := GTl(__e, V2049)

tmp11560 := GTl(__e, V2050)

tmp11561 := Call5(__e, SymFn(symshen_4compile_1head), symshen_4_1m, tmp11559, tmp11560, V2051, V2052)


tmp11562 := GCons(__e, False, Nil)

tmp11563 := GCons(__e, tmp11561, tmp11562)

tmp11564 := GCons(__e, tmp11558, tmp11563)

tmp11565 := GCons(__e, symif, tmp11564)

tmp11566 := GCons(__e, tmp11565, Nil)

tmp11567 := GCons(__e, tmp11554, tmp11566)

tmp11568 := GCons(__e, let__10091, tmp11567)

return GCons(__e, symlet, tmp11568)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.atom-case-minus"))
}


panic("unreachable")
})

tmp11574 := Call2(__e, ns2_1set, symshen_4atom_1case_1minus, tmp11549)


_ = tmp11574

tmp11575 := MakeNative4(func(__e *ControlFlow, V2062, V2063, V2064, V2065 Obj) Obj {
tmp11639 := GConsP(__e, V2062)

var ifres11609 Obj

if True == tmp11639 {
tmp11637 := GHd(__e, V2062)

tmp11638 := GConsP(__e, tmp11637)

var ifres11611 Obj

if True == tmp11638 {
tmp11634 := GHd(__e, V2062)

tmp11635 := GHd(__e, tmp11634)

tmp11636 := GEq(__e, symcons, tmp11635)

var ifres11613 Obj

if True == tmp11636 {
tmp11631 := GHd(__e, V2062)

tmp11632 := GTl(__e, tmp11631)

tmp11633 := GConsP(__e, tmp11632)

var ifres11615 Obj

if True == tmp11633 {
tmp11627 := GHd(__e, V2062)

tmp11628 := GTl(__e, tmp11627)

tmp11629 := GTl(__e, tmp11628)

tmp11630 := GConsP(__e, tmp11629)

var ifres11617 Obj

if True == tmp11630 {
tmp11622 := GHd(__e, V2062)

tmp11623 := GTl(__e, tmp11622)

tmp11624 := GTl(__e, tmp11623)

tmp11625 := GTl(__e, tmp11624)

tmp11626 := GEq(__e, Nil, tmp11625)

var ifres11619 Obj

if True == tmp11626 {
tmp11621 := GConsP(__e, V2063)

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
tmp11576 := Call1(__e, SymFn(symgensym), symTm)


let__10092 := tmp11576
_ = let__10092

tmp11577 := GHd(__e, V2063)

tmp11578 := GCons(__e, V2064, Nil)

tmp11579 := GCons(__e, tmp11577, tmp11578)

tmp11580 := GCons(__e, symshen_4lazyderef, tmp11579)

tmp11581 := GCons(__e, let__10092, Nil)

tmp11582 := GCons(__e, symcons_2, tmp11581)

tmp11583 := GHd(__e, V2062)

tmp11584 := GTl(__e, tmp11583)

tmp11585 := GHd(__e, tmp11584)

tmp11586 := GHd(__e, V2062)

tmp11587 := GTl(__e, tmp11586)

tmp11588 := GTl(__e, tmp11587)

tmp11589 := GHd(__e, tmp11588)

tmp11590 := GTl(__e, V2062)

tmp11591 := GCons(__e, tmp11589, tmp11590)

tmp11592 := GCons(__e, tmp11585, tmp11591)

tmp11593 := GCons(__e, let__10092, Nil)

tmp11594 := GCons(__e, symhd, tmp11593)

tmp11595 := GCons(__e, let__10092, Nil)

tmp11596 := GCons(__e, symtl, tmp11595)

tmp11597 := GTl(__e, V2063)

tmp11598 := GCons(__e, tmp11596, tmp11597)

tmp11599 := GCons(__e, tmp11594, tmp11598)

tmp11600 := Call5(__e, SymFn(symshen_4compile_1head), symshen_4_1m, tmp11592, tmp11599, V2064, V2065)


tmp11601 := GCons(__e, False, Nil)

tmp11602 := GCons(__e, tmp11600, tmp11601)

tmp11603 := GCons(__e, tmp11582, tmp11602)

tmp11604 := GCons(__e, symif, tmp11603)

tmp11605 := GCons(__e, tmp11604, Nil)

tmp11606 := GCons(__e, tmp11580, tmp11605)

tmp11607 := GCons(__e, let__10092, tmp11606)

return GCons(__e, symlet, tmp11607)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.cons-case-minus"))
}


panic("unreachable")
})

tmp11640 := Call2(__e, ns2_1set, symshen_4cons_1case_1minus, tmp11575)


_ = tmp11640

tmp11641 := MakeNative4(func(__e *ControlFlow, V2075, V2076, V2077, V2078 Obj) Obj {
tmp11685 := GConsP(__e, V2075)

var ifres11682 Obj

if True == tmp11685 {
tmp11684 := GConsP(__e, V2076)

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
tmp11642 := Call1(__e, SymFn(symgensym), symTm)


let__10093 := tmp11642
_ = let__10093

tmp11643 := Call1(__e, SymFn(symgensym), symGoTo)


let__10094 := tmp11643
_ = let__10094

tmp11644 := GHd(__e, V2076)

tmp11645 := GCons(__e, V2077, Nil)

tmp11646 := GCons(__e, tmp11644, tmp11645)

tmp11647 := GCons(__e, symshen_4lazyderef, tmp11646)

tmp11648 := GTl(__e, V2075)

tmp11649 := GTl(__e, V2076)

tmp11650 := Call5(__e, SymFn(symshen_4compile_1head), symshen_4_7m, tmp11648, tmp11649, V2077, V2078)


tmp11651 := GCons(__e, tmp11650, Nil)

tmp11652 := GCons(__e, symfreeze, tmp11651)

tmp11653 := GHd(__e, V2075)

tmp11654 := GCons(__e, tmp11653, Nil)

tmp11655 := GCons(__e, let__10093, tmp11654)

tmp11656 := GCons(__e, sym_a, tmp11655)

tmp11657 := GCons(__e, let__10094, Nil)

tmp11658 := GCons(__e, symthaw, tmp11657)

tmp11659 := GCons(__e, let__10093, Nil)

tmp11660 := GCons(__e, symshen_4pvar_2, tmp11659)

tmp11661 := GHd(__e, V2075)

tmp11662 := Call1(__e, SymFn(symshen_4demode), tmp11661)


tmp11663 := GCons(__e, let__10094, Nil)

tmp11664 := GCons(__e, V2077, tmp11663)

tmp11665 := GCons(__e, tmp11662, tmp11664)

tmp11666 := GCons(__e, let__10093, tmp11665)

tmp11667 := GCons(__e, symshen_4bind_b, tmp11666)

tmp11668 := GCons(__e, False, Nil)

tmp11669 := GCons(__e, tmp11667, tmp11668)

tmp11670 := GCons(__e, tmp11660, tmp11669)

tmp11671 := GCons(__e, symif, tmp11670)

tmp11672 := GCons(__e, tmp11671, Nil)

tmp11673 := GCons(__e, tmp11658, tmp11672)

tmp11674 := GCons(__e, tmp11656, tmp11673)

tmp11675 := GCons(__e, symif, tmp11674)

tmp11676 := GCons(__e, tmp11675, Nil)

tmp11677 := GCons(__e, tmp11652, tmp11676)

tmp11678 := GCons(__e, let__10094, tmp11677)

tmp11679 := GCons(__e, tmp11647, tmp11678)

tmp11680 := GCons(__e, let__10093, tmp11679)

return GCons(__e, symlet, tmp11680)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.atom-case-plus"))
}


panic("unreachable")
})

tmp11686 := Call2(__e, ns2_1set, symshen_4atom_1case_1plus, tmp11641)


_ = tmp11686

tmp11687 := MakeNative4(func(__e *ControlFlow, V2089, V2090, V2091, V2092 Obj) Obj {
tmp11778 := GConsP(__e, V2089)

var ifres11748 Obj

if True == tmp11778 {
tmp11776 := GHd(__e, V2089)

tmp11777 := GConsP(__e, tmp11776)

var ifres11750 Obj

if True == tmp11777 {
tmp11773 := GHd(__e, V2089)

tmp11774 := GHd(__e, tmp11773)

tmp11775 := GEq(__e, symcons, tmp11774)

var ifres11752 Obj

if True == tmp11775 {
tmp11770 := GHd(__e, V2089)

tmp11771 := GTl(__e, tmp11770)

tmp11772 := GConsP(__e, tmp11771)

var ifres11754 Obj

if True == tmp11772 {
tmp11766 := GHd(__e, V2089)

tmp11767 := GTl(__e, tmp11766)

tmp11768 := GTl(__e, tmp11767)

tmp11769 := GConsP(__e, tmp11768)

var ifres11756 Obj

if True == tmp11769 {
tmp11761 := GHd(__e, V2089)

tmp11762 := GTl(__e, tmp11761)

tmp11763 := GTl(__e, tmp11762)

tmp11764 := GTl(__e, tmp11763)

tmp11765 := GEq(__e, Nil, tmp11764)

var ifres11758 Obj

if True == tmp11765 {
tmp11760 := GConsP(__e, V2090)

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
tmp11688 := Call1(__e, SymFn(symgensym), symTm)


let__10095 := tmp11688
_ = let__10095

tmp11689 := Call1(__e, SymFn(symgensym), symGoTo)


let__10096 := tmp11689
_ = let__10096

tmp11690 := GHd(__e, V2089)

tmp11691 := GTl(__e, tmp11690)

tmp11692 := GHd(__e, tmp11691)

tmp11693 := GHd(__e, V2089)

tmp11694 := GTl(__e, tmp11693)

tmp11695 := GTl(__e, tmp11694)

tmp11696 := GHd(__e, tmp11695)

tmp11697 := GCons(__e, tmp11692, tmp11696)

tmp11698 := Call1(__e, SymFn(symshen_4extract_1vars), tmp11697)


let__10097 := tmp11698
_ = let__10097

tmp11699 := GHd(__e, V2089)

tmp11700 := Call1(__e, SymFn(symshen_4tame), tmp11699)


let__10098 := tmp11700
_ = let__10098

tmp11701 := Call1(__e, SymFn(symshen_4extract_1vars), let__10098)


let__10099 := tmp11701
_ = let__10099

tmp11702 := GHd(__e, V2090)

tmp11703 := GCons(__e, V2091, Nil)

tmp11704 := GCons(__e, tmp11702, tmp11703)

tmp11705 := GCons(__e, symshen_4lazyderef, tmp11704)

tmp11706 := GTl(__e, V2089)

tmp11707 := GTl(__e, V2090)

tmp11708 := Call5(__e, SymFn(symshen_4compile_1head), symshen_4_7m, tmp11706, tmp11707, V2091, V2092)


tmp11709 := Call2(__e, SymFn(symshen_4goto), let__10097, tmp11708)


tmp11710 := GCons(__e, let__10095, Nil)

tmp11711 := GCons(__e, symcons_2, tmp11710)

tmp11712 := GHd(__e, V2089)

tmp11713 := GTl(__e, tmp11712)

tmp11714 := GCons(__e, let__10095, Nil)

tmp11715 := GCons(__e, symhd, tmp11714)

tmp11716 := GCons(__e, let__10095, Nil)

tmp11717 := GCons(__e, symtl, tmp11716)

tmp11718 := GCons(__e, tmp11717, Nil)

tmp11719 := GCons(__e, tmp11715, tmp11718)

tmp11720 := Call2(__e, SymFn(symshen_4invoke), let__10096, let__10097)


tmp11721 := Call5(__e, SymFn(symshen_4compile_1head), symshen_4_7m, tmp11713, tmp11719, V2091, tmp11720)


tmp11722 := GCons(__e, let__10095, Nil)

tmp11723 := GCons(__e, symshen_4pvar_2, tmp11722)

tmp11724 := Call1(__e, SymFn(symshen_4demode), let__10098)


tmp11725 := Call2(__e, SymFn(symshen_4invoke), let__10096, let__10097)


tmp11726 := GCons(__e, tmp11725, Nil)

tmp11727 := GCons(__e, symfreeze, tmp11726)

tmp11728 := GCons(__e, tmp11727, Nil)

tmp11729 := GCons(__e, V2091, tmp11728)

tmp11730 := GCons(__e, tmp11724, tmp11729)

tmp11731 := GCons(__e, let__10095, tmp11730)

tmp11732 := GCons(__e, symshen_4bind_b, tmp11731)

tmp11733 := Call3(__e, SymFn(symshen_4stpart), let__10099, tmp11732, V2091)


tmp11734 := GCons(__e, False, Nil)

tmp11735 := GCons(__e, tmp11733, tmp11734)

tmp11736 := GCons(__e, tmp11723, tmp11735)

tmp11737 := GCons(__e, symif, tmp11736)

tmp11738 := GCons(__e, tmp11737, Nil)

tmp11739 := GCons(__e, tmp11721, tmp11738)

tmp11740 := GCons(__e, tmp11711, tmp11739)

tmp11741 := GCons(__e, symif, tmp11740)

tmp11742 := GCons(__e, tmp11741, Nil)

tmp11743 := GCons(__e, tmp11709, tmp11742)

tmp11744 := GCons(__e, let__10096, tmp11743)

tmp11745 := GCons(__e, tmp11705, tmp11744)

tmp11746 := GCons(__e, let__10095, tmp11745)

return GCons(__e, symlet, tmp11746)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.cons-case-plus"))
}


panic("unreachable")
})

tmp11779 := Call2(__e, ns2_1set, symshen_4cons_1case_1plus, tmp11687)


_ = tmp11779

tmp11780 := (func() Obj {
var __self Obj
__self = MakeNative1(func(__e *ControlFlow, __self1 Obj) Obj {
__selftop:
V2098 := __self1
_ = V2098
tmp11817 := GConsP(__e, V2098)

var ifres11804 Obj

if True == tmp11817 {
tmp11815 := GHd(__e, V2098)

tmp11816 := GEq(__e, symshen_4_7m, tmp11815)

var ifres11806 Obj

if True == tmp11816 {
tmp11813 := GTl(__e, V2098)

tmp11814 := GConsP(__e, tmp11813)

var ifres11808 Obj

if True == tmp11814 {
tmp11810 := GTl(__e, V2098)

tmp11811 := GTl(__e, tmp11810)

tmp11812 := GEq(__e, Nil, tmp11811)

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
tmp11781 := GTl(__e, V2098)

tmp11782 := GHd(__e, tmp11781)

if SymFn(symshen_4demode) == __self {
__self1 = tmp11782
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4demode), tmp11782)


} else {
tmp11802 := GConsP(__e, V2098)

var ifres11789 Obj

if True == tmp11802 {
tmp11800 := GHd(__e, V2098)

tmp11801 := GEq(__e, symshen_4_1m, tmp11800)

var ifres11791 Obj

if True == tmp11801 {
tmp11798 := GTl(__e, V2098)

tmp11799 := GConsP(__e, tmp11798)

var ifres11793 Obj

if True == tmp11799 {
tmp11795 := GTl(__e, V2098)

tmp11796 := GTl(__e, tmp11795)

tmp11797 := GEq(__e, Nil, tmp11796)

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
tmp11783 := GTl(__e, V2098)

tmp11784 := GHd(__e, tmp11783)

if SymFn(symshen_4demode) == __self {
__self1 = tmp11784
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4demode), tmp11784)


} else {
tmp11787 := GConsP(__e, V2098)

if True == tmp11787 {
tmp11785 := MakeNative1(func(__e *ControlFlow, Z2099 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4demode), Z2099)
panic("unreachable")
})

return TailCall2(__e, SymFn(symmap), tmp11785, V2098)


} else {
return V2098
}


}


}


panic("unreachable")
})
return __self
})()

tmp11818 := Call2(__e, ns2_1set, symshen_4demode, tmp11780)


_ = tmp11818

tmp11819 := MakeNative1(func(__e *ControlFlow, V2100 Obj) Obj {
tmp11824 := Call1(__e, SymFn(symshen_4wildcard_2), V2100)


if True == tmp11824 {
return TailCall1(__e, SymFn(symgensym), symY)
} else {
tmp11822 := GConsP(__e, V2100)

if True == tmp11822 {
tmp11820 := MakeNative1(func(__e *ControlFlow, Z2101 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4tame), Z2101)
panic("unreachable")
})

return TailCall2(__e, SymFn(symmap), tmp11820, V2100)


} else {
return V2100
}


}


panic("unreachable")
})

tmp11825 := Call2(__e, ns2_1set, symshen_4tame, tmp11819)


_ = tmp11825

tmp11826 := MakeNative2(func(__e *ControlFlow, V2102, V2103 Obj) Obj {
tmp11829 := GEq(__e, Nil, V2102)

if True == tmp11829 {
tmp11827 := GCons(__e, V2103, Nil)

return GCons(__e, symfreeze, tmp11827)


} else {
return TailCall2(__e, SymFn(symshen_4goto_1h), V2102, V2103)
}


panic("unreachable")
})

tmp11830 := Call2(__e, ns2_1set, symshen_4goto, tmp11826)


_ = tmp11830

tmp11831 := MakeNative2(func(__e *ControlFlow, V2104, V2105 Obj) Obj {
tmp11840 := GEq(__e, Nil, V2104)

if True == tmp11840 {
return V2105
} else {
tmp11838 := GConsP(__e, V2104)

if True == tmp11838 {
tmp11832 := GHd(__e, V2104)

tmp11833 := GTl(__e, V2104)

tmp11834 := Call2(__e, SymFn(symshen_4goto_1h), tmp11833, V2105)


tmp11835 := GCons(__e, tmp11834, Nil)

tmp11836 := GCons(__e, tmp11832, tmp11835)

return GCons(__e, symlambda, tmp11836)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4goto_1h)
}


}


panic("unreachable")
})

tmp11841 := Call2(__e, ns2_1set, symshen_4goto_1h, tmp11831)


_ = tmp11841

tmp11842 := MakeNative2(func(__e *ControlFlow, V2106, V2107 Obj) Obj {
tmp11845 := GEq(__e, Nil, V2107)

if True == tmp11845 {
tmp11843 := GCons(__e, V2106, Nil)

return GCons(__e, symthaw, tmp11843)


} else {
return GCons(__e, V2106, V2107)
}


panic("unreachable")
})

tmp11846 := Call2(__e, ns2_1set, symshen_4invoke, tmp11842)


_ = tmp11846

tmp11847 := MakeNative1(func(__e *ControlFlow, V2108 Obj) Obj {
return GEq(__e, V2108, sym__)
panic("unreachable")
})

tmp11848 := Call2(__e, ns2_1set, symshen_4wildcard_2, tmp11847)


_ = tmp11848

tmp11849 := MakeNative1(func(__e *ControlFlow, V2109 Obj) Obj {
tmp11856 := GVectorP(__e, V2109)

if True == tmp11856 {
tmp11851 := MakeNative0(func(__e *ControlFlow) Obj {
return GVecGet(__e, V2109, MakeInteger(0))
panic("unreachable")
})

tmp11852 := MakeNative1(func(__e *ControlFlow, Z2110 Obj) Obj {
return symshen_4not_1pvar
panic("unreachable")
})

tmp11853 := Call2(__e, try_1catch, tmp11851, tmp11852)


tmp11854 := GEq(__e, tmp11853, symshen_4pvar)

if True == tmp11854 {
return True
} else {
return False
}


} else {
return False
}


panic("unreachable")
})

tmp11857 := Call2(__e, ns2_1set, symshen_4pvar_2, tmp11849)


_ = tmp11857

tmp11858 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V2111 := __self1
_ = V2111
V2112 := __self2
_ = V2112
tmp11864 := Call1(__e, SymFn(symshen_4pvar_2), V2111)


if True == tmp11864 {
tmp11859 := GVecGet(__e, V2111, MakeInteger(1))

tmp11860 := GVecGet(__e, V2112, tmp11859)

let__10100 := tmp11860
_ = let__10100

tmp11862 := GEq(__e, let__10100, symshen_4_1null_1)

if True == tmp11862 {
return V2111
} else {
if SymFn(symshen_4lazyderef) == __self {
__self1, __self2 = let__10100, V2112
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4lazyderef), let__10100, V2112)
}


} else {
return V2111
}


panic("unreachable")
})
return __self
})()

tmp11865 := Call2(__e, ns2_1set, symshen_4lazyderef, tmp11858)


_ = tmp11865

tmp11866 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V2114 := __self1
_ = V2114
V2115 := __self2
_ = V2115
tmp11878 := GConsP(__e, V2114)

if True == tmp11878 {
tmp11867 := GHd(__e, V2114)

tmp11868 := Call2(__e, SymFn(symshen_4deref), tmp11867, V2115)


tmp11869 := GTl(__e, V2114)

tmp11870 := Call2(__e, SymFn(symshen_4deref), tmp11869, V2115)


return GCons(__e, tmp11868, tmp11870)


} else {
tmp11876 := Call1(__e, SymFn(symshen_4pvar_2), V2114)


if True == tmp11876 {
tmp11871 := GVecGet(__e, V2114, MakeInteger(1))

tmp11872 := GVecGet(__e, V2115, tmp11871)

let__10101 := tmp11872
_ = let__10101

tmp11874 := GEq(__e, let__10101, symshen_4_1null_1)

if True == tmp11874 {
return V2114
} else {
if SymFn(symshen_4deref) == __self {
__self1, __self2 = let__10101, V2115
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4deref), let__10101, V2115)
}


} else {
return V2114
}


}


panic("unreachable")
})
return __self
})()

tmp11879 := Call2(__e, ns2_1set, symshen_4deref, tmp11866)


_ = tmp11879

tmp11880 := MakeNative4(func(__e *ControlFlow, V2117, V2118, V2119, V2120 Obj) Obj {
tmp11881 := Call3(__e, SymFn(symshen_4bindv), V2117, V2118, V2119)


let__10102 := tmp11881
_ = let__10102

tmp11882 := Call1(__e, SymFn(symthaw), V2120)


let__10103 := tmp11882
_ = let__10103

tmp11884 := GEq(__e, let__10103, False)

if True == tmp11884 {
return TailCall3(__e, SymFn(symshen_4unwind), V2117, V2119, let__10103)
} else {
return let__10103
}


panic("unreachable")
})

tmp11885 := Call2(__e, ns2_1set, symshen_4bind_b, tmp11880)


_ = tmp11885

tmp11886 := MakeNative3(func(__e *ControlFlow, V2123, V2124, V2125 Obj) Obj {
tmp11887 := GVecGet(__e, V2123, MakeInteger(1))

return GVecSet(__e, V2125, tmp11887, V2124)


panic("unreachable")
})

tmp11888 := Call2(__e, ns2_1set, symshen_4bindv, tmp11886)


_ = tmp11888

tmp11889 := MakeNative3(func(__e *ControlFlow, V2126, V2127, V2128 Obj) Obj {
tmp11890 := GVecGet(__e, V2126, MakeInteger(1))

tmp11891 := GVecSet(__e, V2127, tmp11890, symshen_4_1null_1)

_ = tmp11891

return V2128


panic("unreachable")
})

tmp11892 := Call2(__e, ns2_1set, symshen_4unwind, tmp11889)


_ = tmp11892

tmp11893 := MakeNative3(func(__e *ControlFlow, V2137, V2138, V2139 Obj) Obj {
tmp11908 := GEq(__e, Nil, V2137)

if True == tmp11908 {
return V2138
} else {
tmp11906 := GConsP(__e, V2137)

if True == tmp11906 {
tmp11894 := GHd(__e, V2137)

tmp11895 := GCons(__e, V2139, Nil)

tmp11896 := GCons(__e, symshen_4newpv, tmp11895)

tmp11897 := GTl(__e, V2137)

tmp11898 := Call3(__e, SymFn(symshen_4stpart), tmp11897, V2138, V2139)


tmp11899 := GCons(__e, tmp11898, Nil)

tmp11900 := GCons(__e, V2139, tmp11899)

tmp11901 := GCons(__e, symshen_4gc, tmp11900)

tmp11902 := GCons(__e, tmp11901, Nil)

tmp11903 := GCons(__e, tmp11896, tmp11902)

tmp11904 := GCons(__e, tmp11894, tmp11903)

return GCons(__e, symlet, tmp11904)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.stpart"))
}


}


panic("unreachable")
})

tmp11909 := Call2(__e, ns2_1set, symshen_4stpart, tmp11893)


_ = tmp11909

tmp11910 := MakeNative2(func(__e *ControlFlow, V2140, V2141 Obj) Obj {
tmp11914 := GEq(__e, V2141, False)

if True == tmp11914 {
tmp11911 := Call1(__e, SymFn(symshen_4ticket_1number), V2140)


let__10104 := tmp11911
_ = let__10104

tmp11912 := Call2(__e, SymFn(symshen_4decrement_1ticket), let__10104, V2140)


_ = tmp11912

return V2141


} else {
return V2141
}


panic("unreachable")
})

tmp11915 := Call2(__e, ns2_1set, symshen_4gc, tmp11910)


_ = tmp11915

tmp11916 := MakeNative2(func(__e *ControlFlow, V2143, V2144 Obj) Obj {
return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(V2144, MakeInteger(1), GSub(__e, V2143, MakeInteger(1)))
}
__typedArg0 := V2144
__typedArg1 := MakeInteger(1)
__typedArg2 := GSub(__e, V2143, MakeInteger(1))
return Call3(__e, SymFn(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()


panic("unreachable")
})

tmp11918 := Call2(__e, ns2_1set, symshen_4decrement_1ticket, tmp11916)


_ = tmp11918

tmp11919 := MakeNative1(func(__e *ControlFlow, V2145 Obj) Obj {
tmp11920 := Call1(__e, SymFn(symshen_4ticket_1number), V2145)


let__10105 := tmp11920
_ = let__10105

tmp11921 := Call1(__e, SymFn(symshen_4make_1prolog_1variable), let__10105)


let__10106 := tmp11921
_ = let__10106

tmp11922 := Call2(__e, SymFn(symshen_4nextticket), V2145, let__10105)


let__10107 := tmp11922
_ = let__10107

return let__10106


panic("unreachable")
})

tmp11923 := Call2(__e, ns2_1set, symshen_4newpv, tmp11919)


_ = tmp11923

tmp11924 := MakeNative1(func(__e *ControlFlow, V2149 Obj) Obj {
return GVecGet(__e, V2149, MakeInteger(1))
panic("unreachable")
})

tmp11925 := Call2(__e, ns2_1set, symshen_4ticket_1number, tmp11924)


_ = tmp11925

tmp11926 := MakeNative2(func(__e *ControlFlow, V2150, V2151 Obj) Obj {
tmp11927 := GVecSet(__e, V2150, V2151, symshen_4_1null_1)

let__10108 := tmp11927
_ = let__10108

return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(let__10108, MakeInteger(1), GAdd(__e, V2151, MakeInteger(1)))
}
__typedArg0 := let__10108
__typedArg1 := MakeInteger(1)
__typedArg2 := GAdd(__e, V2151, MakeInteger(1))
return Call3(__e, SymFn(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()


panic("unreachable")
})

tmp11929 := Call2(__e, ns2_1set, symshen_4nextticket, tmp11926)


_ = tmp11929

tmp11930 := MakeNative1(func(__e *ControlFlow, V2153 Obj) Obj {
tmp11931 := PrimGuard1(__e, symabsvector, PrimAbsvector, MakeInteger(2))

tmp11932 := GVecSet(__e, tmp11931, MakeInteger(0), symshen_4pvar)

return GVecSet(__e, tmp11932, MakeInteger(1), V2153)


panic("unreachable")
})

tmp11933 := Call2(__e, ns2_1set, symshen_4make_1prolog_1variable, tmp11930)


_ = tmp11933

tmp11934 := MakeNative1(func(__e *ControlFlow, V2154 Obj) Obj {
tmp11935 := GVecGet(__e, V2154, MakeInteger(1))

tmp11936 := Call3(__e, SymFn(symshen_4app), tmp11935, MakeString(""), symshen_4a)


return PrimGuard2(__e, symcn, PrimStringConcat, MakeString("Var"), tmp11936)


panic("unreachable")
})

tmp11937 := Call2(__e, ns2_1set, symshen_4pvar, tmp11934)


_ = tmp11937

tmp11938 := MakeNative0(func(__e *ControlFlow) Obj {
tmp11939 := GValue(__e, symshen_4_dinfs_d)

return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dinfs_d, GAdd(__e, MakeInteger(1), tmp11939))
}
__typedArg0 := symshen_4_dinfs_d
__typedArg1 := GAdd(__e, MakeInteger(1), tmp11939)
return Call2(__e, SymFn(symset), __typedArg0, __typedArg1)
})()


panic("unreachable")
})

tmp11941 := Call2(__e, ns2_1set, symshen_4incinfs, tmp11938)


_ = tmp11941

tmp11942 := MakeNative1(func(__e *ControlFlow, V2155 Obj) Obj {
tmp11949 := PrimGuard1(__e, syminteger_2, PrimIsInteger, V2155)

var ifres11946 Obj

if True == tmp11949 {
tmp11948 := GGt(__e, V2155, MakeInteger(0))

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
return PrimGuard2(__e, symset, PrimSet, symshen_4_dsize_1prolog_1vector_d, V2155)
} else {
tmp11943 := Call3(__e, SymFn(symshen_4app), V2155, MakeString(""), symshen_4a)


return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("prolog vector size: size should be a positive integer; not "), tmp11943))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("prolog vector size: size should be a positive integer; not "), tmp11943)
return Call1(__e, SymFn(symsimple_1error), __typedArg0)
})()


}


panic("unreachable")
})

tmp11950 := Call2(__e, ns2_1set, symshen_4prolog_1vector_1size, tmp11942)


_ = tmp11950

tmp11951 := (func() Obj {
var __self Obj
__self = MakeNative4(func(__e *ControlFlow, __self1, __self2, __self3, __self4 Obj) Obj {
__selftop:
V2167 := __self1
_ = V2167
V2168 := __self2
_ = V2168
V2169 := __self3
_ = V2169
V2170 := __self4
_ = V2170
tmp11981 := GEq(__e, V2167, V2168)

if True == tmp11981 {
return TailCall1(__e, SymFn(symthaw), V2170)
} else {
tmp11979 := Call1(__e, SymFn(symshen_4pvar_2), V2167)


var ifres11974 Obj

if True == tmp11979 {
tmp11976 := Call2(__e, SymFn(symshen_4deref), V2168, V2169)


tmp11977 := Call2(__e, SymFn(symshen_4occurs_1check_2), V2167, tmp11976)


tmp11978 := GNot(__e, tmp11977)

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
return TailCall4(__e, SymFn(symshen_4bind_b), V2167, V2168, V2169, V2170)
} else {
tmp11972 := Call1(__e, SymFn(symshen_4pvar_2), V2168)


var ifres11967 Obj

if True == tmp11972 {
tmp11969 := Call2(__e, SymFn(symshen_4deref), V2167, V2169)


tmp11970 := Call2(__e, SymFn(symshen_4occurs_1check_2), V2168, tmp11969)


tmp11971 := GNot(__e, tmp11970)

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
return TailCall4(__e, SymFn(symshen_4bind_b), V2168, V2167, V2169, V2170)
} else {
tmp11965 := GConsP(__e, V2167)

var ifres11962 Obj

if True == tmp11965 {
tmp11964 := GConsP(__e, V2168)

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
tmp11952 := GHd(__e, V2167)

tmp11953 := Call2(__e, SymFn(symshen_4lazyderef), tmp11952, V2169)


tmp11954 := GHd(__e, V2168)

tmp11955 := Call2(__e, SymFn(symshen_4lazyderef), tmp11954, V2169)


tmp11956 := MakeNative0(func(__e *ControlFlow) Obj {
tmp11957 := GTl(__e, V2167)

tmp11958 := Call2(__e, SymFn(symshen_4lazyderef), tmp11957, V2169)


tmp11959 := GTl(__e, V2168)

tmp11960 := Call2(__e, SymFn(symshen_4lazyderef), tmp11959, V2169)


return TailCall4(__e, SymFn(symshen_4lzy_a_b), tmp11958, tmp11960, V2169, V2170)


panic("unreachable")
})

if SymFn(symshen_4lzy_a_b) == __self {
__self1, __self2, __self3, __self4 = tmp11953, tmp11955, V2169, tmp11956
__e.Tick()
goto __selftop
}
return TailCall4(__e, SymFn(symshen_4lzy_a_b), tmp11953, tmp11955, V2169, tmp11956)


} else {
return False
}


}


}


}


panic("unreachable")
})
return __self
})()

tmp11982 := Call2(__e, ns2_1set, symshen_4lzy_a_b, tmp11951)


_ = tmp11982

tmp11983 := (func() Obj {
var __self Obj
__self = MakeNative4(func(__e *ControlFlow, __self1, __self2, __self3, __self4 Obj) Obj {
__selftop:
V2182 := __self1
_ = V2182
V2183 := __self2
_ = V2183
V2184 := __self3
_ = V2184
V2185 := __self4
_ = V2185
tmp12003 := GEq(__e, V2182, V2183)

if True == tmp12003 {
return TailCall1(__e, SymFn(symthaw), V2185)
} else {
tmp12001 := Call1(__e, SymFn(symshen_4pvar_2), V2182)


if True == tmp12001 {
return TailCall4(__e, SymFn(symshen_4bind_b), V2182, V2183, V2184, V2185)
} else {
tmp11999 := Call1(__e, SymFn(symshen_4pvar_2), V2183)


if True == tmp11999 {
return TailCall4(__e, SymFn(symshen_4bind_b), V2183, V2182, V2184, V2185)
} else {
tmp11997 := GConsP(__e, V2182)

var ifres11994 Obj

if True == tmp11997 {
tmp11996 := GConsP(__e, V2183)

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
tmp11984 := GHd(__e, V2182)

tmp11985 := Call2(__e, SymFn(symshen_4lazyderef), tmp11984, V2184)


tmp11986 := GHd(__e, V2183)

tmp11987 := Call2(__e, SymFn(symshen_4lazyderef), tmp11986, V2184)


tmp11988 := MakeNative0(func(__e *ControlFlow) Obj {
tmp11989 := GTl(__e, V2182)

tmp11990 := Call2(__e, SymFn(symshen_4lazyderef), tmp11989, V2184)


tmp11991 := GTl(__e, V2183)

tmp11992 := Call2(__e, SymFn(symshen_4lazyderef), tmp11991, V2184)


return TailCall4(__e, SymFn(symshen_4lzy_a), tmp11990, tmp11992, V2184, V2185)


panic("unreachable")
})

if SymFn(symshen_4lzy_a) == __self {
__self1, __self2, __self3, __self4 = tmp11985, tmp11987, V2184, tmp11988
__e.Tick()
goto __selftop
}
return TailCall4(__e, SymFn(symshen_4lzy_a), tmp11985, tmp11987, V2184, tmp11988)


} else {
return False
}


}


}


}


panic("unreachable")
})
return __self
})()

tmp12004 := Call2(__e, ns2_1set, symshen_4lzy_a, tmp11983)


_ = tmp12004

tmp12005 := MakeNative2(func(__e *ControlFlow, V2191, V2192 Obj) Obj {
tmp12015 := GEq(__e, V2191, V2192)

if True == tmp12015 {
return True
} else {
tmp12013 := GConsP(__e, V2192)

if True == tmp12013 {
tmp12010 := GHd(__e, V2192)

tmp12011 := Call2(__e, SymFn(symshen_4occurs_1check_2), V2191, tmp12010)


if True == tmp12011 {
return True
} else {
tmp12007 := GTl(__e, V2192)

tmp12008 := Call2(__e, SymFn(symshen_4occurs_1check_2), V2191, tmp12007)


if True == tmp12008 {
return True
} else {
return False
}


}


} else {
return False
}


}


panic("unreachable")
})

tmp12016 := Call2(__e, ns2_1set, symshen_4occurs_1check_2, tmp12005)


_ = tmp12016

tmp12017 := MakeNative5(func(__e *ControlFlow, V2193, V2194, V2195, V2196, V2197 Obj) Obj {
tmp12018 := Call1(__e, V2193, V2194)


tmp12019 := Call1(__e, tmp12018, V2195)


tmp12020 := Call1(__e, tmp12019, V2196)


return TailCall1(__e, tmp12020, V2197)


panic("unreachable")
})

tmp12021 := Call2(__e, ns2_1set, symcall, tmp12017)


_ = tmp12021

tmp12022 := MakeNative5(func(__e *ControlFlow, V2204, V2205, V2206, V2207, V2208 Obj) Obj {
return TailCall2(__e, SymFn(symshen_4deref), V2204, V2205)
panic("unreachable")
})

tmp12023 := Call2(__e, ns2_1set, symreturn, tmp12022)


_ = tmp12023

tmp12024 := MakeNative5(func(__e *ControlFlow, V2215, V2216, V2217, V2218, V2219 Obj) Obj {
if True == V2215 {
return TailCall1(__e, SymFn(symthaw), V2219)
} else {
return False
}
panic("unreachable")
})

tmp12026 := Call2(__e, ns2_1set, symwhen, tmp12024)


_ = tmp12026

tmp12027 := MakeNative6(func(__e *ControlFlow, V2220, V2221, V2222, V2223, V2224, V2225 Obj) Obj {
tmp12028 := Call2(__e, SymFn(symshen_4lazyderef), V2220, V2222)


tmp12029 := Call2(__e, SymFn(symshen_4lazyderef), V2221, V2222)


return TailCall4(__e, SymFn(symshen_4lzy_a), tmp12028, tmp12029, V2222, V2225)


panic("unreachable")
})

tmp12030 := Call2(__e, ns2_1set, symis, tmp12027)


_ = tmp12030

tmp12031 := MakeNative6(func(__e *ControlFlow, V2226, V2227, V2228, V2229, V2230, V2231 Obj) Obj {
tmp12032 := Call2(__e, SymFn(symshen_4lazyderef), V2226, V2228)


tmp12033 := Call2(__e, SymFn(symshen_4lazyderef), V2227, V2228)


return TailCall4(__e, SymFn(symshen_4lzy_a_b), tmp12032, tmp12033, V2228, V2231)


panic("unreachable")
})

tmp12034 := Call2(__e, ns2_1set, symis_b, tmp12031)


_ = tmp12034

tmp12035 := MakeNative6(func(__e *ControlFlow, V2236, V2237, V2238, V2239, V2240, V2241 Obj) Obj {
return TailCall4(__e, SymFn(symshen_4bind_b), V2236, V2237, V2238, V2241)
panic("unreachable")
})

tmp12036 := Call2(__e, ns2_1set, symbind, tmp12035)


_ = tmp12036

tmp12037 := MakeNative5(func(__e *ControlFlow, V2242, V2243, V2244, V2245, V2246 Obj) Obj {
tmp12039 := Call2(__e, SymFn(symshen_4lazyderef), V2242, V2243)


tmp12040 := Call1(__e, SymFn(symshen_4pvar_2), tmp12039)


if True == tmp12040 {
return TailCall1(__e, SymFn(symthaw), V2246)
} else {
return False
}


panic("unreachable")
})

tmp12041 := Call2(__e, ns2_1set, symvar_2, tmp12037)


_ = tmp12041

tmp12042 := MakeNative1(func(__e *ControlFlow, V2249 Obj) Obj {
return MakeString("|prolog vector|")
panic("unreachable")
})

tmp12043 := Call2(__e, ns2_1set, symshen_4print_1prolog_1vector, tmp12042)


_ = tmp12043

tmp12044 := (func() Obj {
var __self Obj
__self = MakeNative5(func(__e *ControlFlow, __self1, __self2, __self3, __self4, __self5 Obj) Obj {
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
tmp12056 := GEq(__e, Nil, V2268)

if True == tmp12056 {
return False
} else {
tmp12054 := GConsP(__e, V2268)

if True == tmp12054 {
tmp12045 := GHd(__e, V2268)

tmp12046 := Call1(__e, tmp12045, V2269)


tmp12047 := Call1(__e, tmp12046, V2270)


tmp12048 := Call1(__e, tmp12047, V2271)


tmp12049 := Call1(__e, tmp12048, V2272)


let__10109 := tmp12049
_ = let__10109

tmp12052 := GEq(__e, let__10109, False)

if True == tmp12052 {
tmp12050 := GTl(__e, V2268)

if SymFn(symfork) == __self {
__self1, __self2, __self3, __self4, __self5 = tmp12050, V2269, V2270, V2271, V2272
__e.Tick()
goto __selftop
}
return TailCall5(__e, SymFn(symfork), tmp12050, V2269, V2270, V2271, V2272)


} else {
return let__10109
}


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("fork expects a list of literals\n"))
}


}


panic("unreachable")
})
return __self
})()

tmp12057 := Call2(__e, ns2_1set, symfork, tmp12044)


_ = tmp12057

tmp12058 := MakeNative(func(__e *ControlFlow) {
V2274 := __e.Get(1)
V2275 := __e.Get(2)
V2276 := __e.Get(3)
V2277 := __e.Get(4)
V2278 := __e.Get(5)
V2279 := __e.Get(6)
V2280 := __e.Get(7)
__e.ReturnDirect(func() Obj {
tmp12064 := Call1(__e, SymFn(symshen_4unlocked_2), V2278)


if True == tmp12064 {
tmp12059 := Call1(__e, SymFn(symshen_4newpv), V2277)


let__10110 := tmp12059
_ = let__10110

tmp12060 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp12060

tmp12061 := MakeNative0(func(__e *ControlFlow) Obj {
__e.TailApply(SymFn(symshen_4findall_1h), V2274, V2275, V2276, let__10110, V2277, V2278, V2279, V2280)
return nil
panic("unreachable")
})

tmp12062 := Call6(__e, SymFn(symis), let__10110, Nil, V2277, V2278, V2279, tmp12061)


return TailCall2(__e, SymFn(symshen_4gc), V2277, tmp12062)


} else {
return False
}


panic("unreachable")
}())
}, 7)

tmp12065 := Call2(__e, ns2_1set, symfindall, tmp12058)


_ = tmp12065

tmp12066 := MakeNative(func(__e *ControlFlow) {
V2282 := __e.Get(1)
V2283 := __e.Get(2)
V2284 := __e.Get(3)
V2285 := __e.Get(4)
V2286 := __e.Get(5)
V2287 := __e.Get(6)
V2288 := __e.Get(7)
V2289 := __e.Get(8)
__e.ReturnDirect(func() Obj {
tmp12071 := Call1(__e, SymFn(symshen_4unlocked_2), V2287)


var ifres12067 Obj

if True == tmp12071 {
tmp12068 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp12068

tmp12069 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall6(__e, SymFn(symshen_4overbind), V2282, V2285, V2286, V2287, V2288, V2289)
panic("unreachable")
})

tmp12070 := Call5(__e, SymFn(symcall), V2283, V2286, V2287, V2288, tmp12069)


ifres12067 = tmp12070


} else {
ifres12067 = False


}

let__10111 := ifres12067
_ = let__10111

tmp12076 := GEq(__e, let__10111, False)

if True == tmp12076 {
tmp12074 := Call1(__e, SymFn(symshen_4unlocked_2), V2287)


if True == tmp12074 {
tmp12072 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp12072

return TailCall6(__e, SymFn(symis_b), V2284, V2285, V2286, V2287, V2288, V2289)


} else {
return False
}


} else {
return let__10111
}


panic("unreachable")
}())
}, 8)

tmp12077 := Call2(__e, ns2_1set, symshen_4findall_1h, tmp12066)


_ = tmp12077

tmp12078 := MakeNative6(func(__e *ControlFlow, V2297, V2298, V2299, V2300, V2301, V2302 Obj) Obj {
tmp12079 := Call2(__e, SymFn(symshen_4deref), V2297, V2299)


tmp12080 := Call2(__e, SymFn(symshen_4lazyderef), V2298, V2299)


tmp12081 := GCons(__e, tmp12079, tmp12080)

tmp12082 := Call3(__e, SymFn(symshen_4bindv), V2298, tmp12081, V2299)


_ = tmp12082

return False


panic("unreachable")
})

tmp12083 := Call2(__e, ns2_1set, symshen_4overbind, tmp12078)


_ = tmp12083

tmp12084 := MakeNative1(func(__e *ControlFlow, V2305 Obj) Obj {
tmp12088 := GEq(__e, sym_7, V2305)

if True == tmp12088 {
return PrimGuard2(__e, symset, PrimSet, symshen_4_doccurs_d, True)
} else {
tmp12086 := GEq(__e, sym_1, V2305)

if True == tmp12086 {
return PrimGuard2(__e, symset, PrimSet, symshen_4_doccurs_d, False)
} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("occurs-check expects a + or a -.\n"))
}


}


panic("unreachable")
})

return TailCall2(__e, ns2_1set, symoccurs_1check, tmp12084)




panic("unreachable")
})

