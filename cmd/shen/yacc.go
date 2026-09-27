package main

import . "github.com/pyrex41/shen-go/kl"

var YaccMain = MakeNative(func(__e *ControlFlow) {
tmp17924 := MakeNative(func(__e *ControlFlow) {
V112 := __e.Get(1)
_ = V112
V113 := __e.Get(2)
_ = V113
tmp17925 := Call(__e, V112, V113)


let__17786 := tmp17925
_ = let__17786

tmp17932 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17786)


if True == tmp17932 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("parse failure\n"))
}
__typedArg0 := MakeString("parse failure\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
} else {
tmp17930 := Call(__e, PrimFunc(symshen_4partial_1parse_1failure_2), let__17786)


if True == tmp17930 {
tmp17926 := Call(__e, PrimFunc(symshen_4in_1_6), let__17786)


tmp17927 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dresidue_d, tmp17926)
}
__typedArg0 := symshen_4_dresidue_d
__typedArg1 := tmp17926
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp17927

tmp17928 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dresidue_d)
}
__typedArg0 := symshen_4_dresidue_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4raise_1syntax_1error), tmp17928)
return


} else {
__e.TailApply(PrimFunc(symshen_4_5_1out), let__17786)
return
}


}


}, 2)

tmp17933 := Call(__e, ns2_1set, symcompile, tmp17924)


_ = tmp17933

tmp17934 := MakeNative(func(__e *ControlFlow) {
V115 := __e.Get(1)
_ = V115
tmp17935 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dmaximum_1print_1sequence_1size_d)
}
__typedArg0 := sym_dmaximum_1print_1sequence_1size_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp17936 := Call(__e, PrimFunc(symshen_4syntax_1error_1message), tmp17935, MakeInteger(0), V115)


tmp17938 := Call(__e, PrimFunc(symshen_4proc_1nl), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("syntax error here: "))
__typedS1, __typedOK1 := TypedString(tmp17936)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("syntax error here: ")
__typedArg1 := tmp17936
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(tmp17938)
}
__typedArg0 := tmp17938
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}, 1)

tmp17939 := Call(__e, ns2_1set, symshen_4raise_1syntax_1error, tmp17934)


_ = tmp17939

tmp17940 := MakeNative(func(__e *ControlFlow) {
V123 := __e.Get(1)
_ = V123
V124 := __e.Get(2)
_ = V124
V125 := __e.Get(3)
_ = V125
tmp17951 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V125)
}
__typedArg0 := Nil
__typedArg1 := V125
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17951 {
__e.Return(MakeString("\n"))
return
} else {
tmp17949 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V123, V124)
}
__typedArg0 := V123
__typedArg1 := V124
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17949 {
__e.Return(MakeString("...etc \n"))
return
} else {
tmp17947 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V125)
}
__typedArg0 := V125
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17947 {
tmp17941 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V125)
}
__typedArg0 := V125
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17942 := Call(__e, PrimFunc(symshen_4app), tmp17941, MakeString(" "), symshen_4s)


tmp17943 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V124)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V124
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()

tmp17944 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V125)
}
__typedArg0 := V125
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17945 := Call(__e, PrimFunc(symshen_4syntax_1error_1message), V123, tmp17943, tmp17944)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(tmp17942)
__typedS1, __typedOK1 := TypedString(tmp17945)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := tmp17942
__typedArg1 := tmp17945
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("partial function shen.syntax-error-message"))
}
__typedArg0 := MakeString("partial function shen.syntax-error-message")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}


}, 3)

tmp17952 := Call(__e, ns2_1set, symshen_4syntax_1error_1message, tmp17940)


_ = tmp17952

tmp17953 := MakeNative(func(__e *ControlFlow) {
V126 := __e.Get(1)
_ = V126
tmp17954 := Call(__e, PrimFunc(symfail))


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V126, tmp17954)
}
__typedArg0 := V126
__typedArg1 := tmp17954
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp17955 := Call(__e, ns2_1set, symshen_4parse_1failure_2, tmp17953)


_ = tmp17955

tmp17956 := MakeNative(func(__e *ControlFlow) {
V127 := __e.Get(1)
_ = V127
tmp17957 := Call(__e, PrimFunc(symshen_4in_1_6), V127)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp17957)
}
__typedArg0 := tmp17957
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})())
return


}, 1)

tmp17958 := Call(__e, ns2_1set, symshen_4partial_1parse_1failure_2, tmp17956)


_ = tmp17958

tmp17959 := MakeNative(func(__e *ControlFlow) {
V130 := __e.Get(1)
_ = V130
tmp17972 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V130)
}
__typedArg0 := V130
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17963 Obj

if True == tmp17972 {
tmp17970 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V130)
}
__typedArg0 := V130
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17971 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp17970)
}
__typedArg0 := tmp17970
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17965 Obj

if True == tmp17971 {
tmp17967 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V130)
}
__typedArg0 := V130
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17968 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp17967)
}
__typedArg0 := tmp17967
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17969 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp17968)
}
__typedArg0 := Nil
__typedArg1 := tmp17968
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17966 Obj

if True == tmp17969 {
ifres17966 = True


} else {
ifres17966 = False


}

ifres17965 = ifres17966


} else {
ifres17965 = False


}

var ifres17964 Obj

if True == ifres17965 {
ifres17964 = True


} else {
ifres17964 = False


}

ifres17963 = ifres17964


} else {
ifres17963 = False


}

if True == ifres17963 {
tmp17960 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V130)
}
__typedArg0 := V130
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp17960)
}
__typedArg0 := tmp17960
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return


} else {
tmp17961 := Call(__e, PrimFunc(symshen_4app), V130, MakeString(" is not a YACC stream\n"), symshen_4s)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(tmp17961)
}
__typedArg0 := tmp17961
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}


}, 1)

tmp17973 := Call(__e, ns2_1set, symshen_4objectcode, tmp17959)


_ = tmp17973

tmp17974 := MakeNative(func(__e *ControlFlow) {
V131 := __e.Get(1)
_ = V131
tmp17975 := MakeNative(func(__e *ControlFlow) {
Z132 := __e.Get(1)
_ = Z132
__e.TailApply(PrimFunc(symshen_4_5yacc_6), Z132)
return
}, 1)

__e.TailApply(PrimFunc(symcompile), tmp17975, V131)
return


}, 1)

tmp17976 := Call(__e, ns2_1set, symshen_4yacc_1_6shen, tmp17974)


_ = tmp17976

tmp17977 := MakeNative(func(__e *ControlFlow) {
V133 := __e.Get(1)
_ = V133
tmp18004 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V133)
}
__typedArg0 := V133
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17978 Obj

if True == tmp18004 {
tmp17979 := Call(__e, PrimFunc(symhead), V133)


let__17788 := tmp17979
_ = let__17788

tmp17980 := Call(__e, PrimFunc(symtail), V133)


let__17789 := tmp17980
_ = let__17789

tmp17981 := Call(__e, PrimFunc(symshen_4_5yaccsig_6), let__17789)


let__17790 := tmp17981
_ = let__17790

tmp18002 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17790)


var ifres17982 Obj

if True == tmp18002 {
tmp17983 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17982 = tmp17983


} else {
tmp17984 := Call(__e, PrimFunc(symshen_4_5_1out), let__17790)


let__17791 := tmp17984
_ = let__17791

tmp17985 := Call(__e, PrimFunc(symshen_4in_1_6), let__17790)


let__17792 := tmp17985
_ = let__17792

tmp17986 := Call(__e, PrimFunc(symshen_4_5c_1rules_6), let__17792)


let__17793 := tmp17986
_ = let__17793

tmp18001 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17793)


var ifres17987 Obj

if True == tmp18001 {
tmp17988 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17987 = tmp17988


} else {
tmp17989 := Call(__e, PrimFunc(symshen_4_5_1out), let__17793)


let__17794 := tmp17989
_ = let__17794

tmp17990 := Call(__e, PrimFunc(symshen_4in_1_6), let__17793)


let__17795 := tmp17990
_ = let__17795

tmp17991 := Call(__e, PrimFunc(symgensym), symS)


let__17796 := tmp17991
_ = let__17796

tmp17992 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17788, Nil)
}
__typedArg0 := let__17788
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17993 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefine, tmp17992)
}
__typedArg0 := symdefine
__typedArg1 := tmp17992
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17994 := Call(__e, PrimFunc(symshen_4c_1rules_1_6shen), let__17791, let__17796, let__17794)


tmp17995 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17994, Nil)
}
__typedArg0 := tmp17994
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17996 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_6, tmp17995)
}
__typedArg0 := sym_1_6
__typedArg1 := tmp17995
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17997 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17796, tmp17996)
}
__typedArg0 := let__17796
__typedArg1 := tmp17996
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17998 := Call(__e, PrimFunc(symappend), let__17791, tmp17997)


tmp17999 := Call(__e, PrimFunc(symappend), tmp17993, tmp17998)


let__17797 := tmp17999
_ = let__17797

tmp18000 := Call(__e, PrimFunc(symshen_4comb), let__17795, let__17797)


ifres17987 = tmp18000


}

ifres17982 = ifres17987


}

ifres17978 = ifres17982


} else {
tmp18003 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres17978 = tmp18003


}

let__17787 := ifres17978
_ = let__17787

tmp18006 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17787)


if True == tmp18006 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__17787)
return
}


}, 1)

tmp18007 := Call(__e, ns2_1set, symshen_4_5yacc_6, tmp17977)


_ = tmp18007

tmp18008 := MakeNative(func(__e *ControlFlow) {
V145 := __e.Get(1)
_ = V145
tmp18065 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V145)
}
__typedArg0 := V145
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18009 Obj

if True == tmp18065 {
tmp18010 := Call(__e, PrimFunc(symhead), V145)


let__17799 := tmp18010
_ = let__17799

tmp18011 := Call(__e, PrimFunc(symtail), V145)


let__17800 := tmp18011
_ = let__17800

tmp18063 := Call(__e, PrimFunc(symshen_4ccons_2), let__17800)


var ifres18012 Obj

if True == tmp18063 {
tmp18013 := Call(__e, PrimFunc(symhead), let__17800)


let__17801 := tmp18013
_ = let__17801

tmp18014 := Call(__e, PrimFunc(symtail), let__17800)


let__17802 := tmp18014
_ = let__17802

tmp18061 := Call(__e, PrimFunc(symshen_4hds_a_2), let__17801, symlist)


var ifres18015 Obj

if True == tmp18061 {
tmp18016 := Call(__e, PrimFunc(symtail), let__17801)


let__17803 := tmp18016
_ = let__17803

tmp18059 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__17803)
}
__typedArg0 := let__17803
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18017 Obj

if True == tmp18059 {
tmp18018 := Call(__e, PrimFunc(symhead), let__17803)


let__17804 := tmp18018
_ = let__17804

tmp18019 := Call(__e, PrimFunc(symtail), let__17803)


let__17805 := tmp18019
_ = let__17805

tmp18020 := Call(__e, PrimFunc(sym_5end_6), let__17805)


let__17806 := tmp18020
_ = let__17806

tmp18057 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17806)


var ifres18021 Obj

if True == tmp18057 {
tmp18022 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres18021 = tmp18022


} else {
tmp18023 := Call(__e, PrimFunc(symshen_4in_1_6), let__17806)


let__17807 := tmp18023
_ = let__17807

tmp18056 := Call(__e, PrimFunc(symshen_4hds_a_2), let__17802, sym_a_a_6)


var ifres18024 Obj

if True == tmp18056 {
tmp18025 := Call(__e, PrimFunc(symtail), let__17802)


let__17808 := tmp18025
_ = let__17808

tmp18054 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__17808)
}
__typedArg0 := let__17808
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18026 Obj

if True == tmp18054 {
tmp18027 := Call(__e, PrimFunc(symhead), let__17808)


let__17809 := tmp18027
_ = let__17809

tmp18028 := Call(__e, PrimFunc(symtail), let__17808)


let__17810 := tmp18028
_ = let__17810

tmp18052 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__17810)
}
__typedArg0 := let__17810
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18029 Obj

if True == tmp18052 {
tmp18030 := Call(__e, PrimFunc(symhead), let__17810)


let__17811 := tmp18030
_ = let__17811

tmp18031 := Call(__e, PrimFunc(symtail), let__17810)


let__17812 := tmp18031
_ = let__17812

tmp18050 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_i, let__17799)
}
__typedArg0 := sym_i
__typedArg1 := let__17799
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18047 Obj

if True == tmp18050 {
tmp18049 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_j, let__17811)
}
__typedArg0 := sym_j
__typedArg1 := let__17811
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18048 Obj

if True == tmp18049 {
ifres18048 = True


} else {
ifres18048 = False


}

ifres18047 = ifres18048


} else {
ifres18047 = False


}

var ifres18032 Obj

if True == ifres18047 {
tmp18033 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17804, Nil)
}
__typedArg0 := let__17804
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18034 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp18033)
}
__typedArg0 := symlist
__typedArg1 := tmp18033
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18035 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17804, Nil)
}
__typedArg0 := let__17804
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18036 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp18035)
}
__typedArg0 := symlist
__typedArg1 := tmp18035
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18037 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17809, Nil)
}
__typedArg0 := let__17809
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18038 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18036, tmp18037)
}
__typedArg0 := tmp18036
__typedArg1 := tmp18037
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18039 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp18038)
}
__typedArg0 := symstr
__typedArg1 := tmp18038
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18040 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_j, Nil)
}
__typedArg0 := sym_j
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18041 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18039, tmp18040)
}
__typedArg0 := tmp18039
__typedArg1 := tmp18040
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18042 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp18041)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp18041
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18043 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18034, tmp18042)
}
__typedArg0 := tmp18034
__typedArg1 := tmp18042
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18044 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_i, tmp18043)
}
__typedArg0 := sym_i
__typedArg1 := tmp18043
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18045 := Call(__e, PrimFunc(symshen_4comb), let__17812, tmp18044)


ifres18032 = tmp18045


} else {
tmp18046 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres18032 = tmp18046


}

ifres18029 = ifres18032


} else {
tmp18051 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres18029 = tmp18051


}

ifres18026 = ifres18029


} else {
tmp18053 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres18026 = tmp18053


}

ifres18024 = ifres18026


} else {
tmp18055 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres18024 = tmp18055


}

ifres18021 = ifres18024


}

ifres18017 = ifres18021


} else {
tmp18058 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres18017 = tmp18058


}

ifres18015 = ifres18017


} else {
tmp18060 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres18015 = tmp18060


}

ifres18012 = ifres18015


} else {
tmp18062 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres18012 = tmp18062


}

ifres18009 = ifres18012


} else {
tmp18064 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres18009 = tmp18064


}

let__17798 := ifres18009
_ = let__17798

tmp18075 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17798)


if True == tmp18075 {
tmp18066 := Call(__e, PrimFunc(sym_5e_6), V145)


let__17814 := tmp18066
_ = let__17814

tmp18071 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17814)


var ifres18067 Obj

if True == tmp18071 {
tmp18068 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres18067 = tmp18068


} else {
tmp18069 := Call(__e, PrimFunc(symshen_4in_1_6), let__17814)


let__17815 := tmp18069
_ = let__17815

tmp18070 := Call(__e, PrimFunc(symshen_4comb), let__17815, Nil)


ifres18067 = tmp18070


}

let__17813 := ifres18067
_ = let__17813

tmp18073 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17813)


if True == tmp18073 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__17813)
return
}


} else {
__e.Return(let__17798)
return
}


}, 1)

tmp18076 := Call(__e, ns2_1set, symshen_4_5yaccsig_6, tmp18008)


_ = tmp18076

tmp18077 := MakeNative(func(__e *ControlFlow) {
V164 := __e.Get(1)
_ = V164
tmp18078 := Call(__e, PrimFunc(symshen_4_5c_1rule_6), V164)


let__17817 := tmp18078
_ = let__17817

tmp18091 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17817)


var ifres18079 Obj

if True == tmp18091 {
tmp18080 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres18079 = tmp18080


} else {
tmp18081 := Call(__e, PrimFunc(symshen_4_5_1out), let__17817)


let__17818 := tmp18081
_ = let__17818

tmp18082 := Call(__e, PrimFunc(symshen_4in_1_6), let__17817)


let__17819 := tmp18082
_ = let__17819

tmp18083 := Call(__e, PrimFunc(symshen_4_5c_1rules_6), let__17819)


let__17820 := tmp18083
_ = let__17820

tmp18090 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17820)


var ifres18084 Obj

if True == tmp18090 {
tmp18085 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres18084 = tmp18085


} else {
tmp18086 := Call(__e, PrimFunc(symshen_4_5_1out), let__17820)


let__17821 := tmp18086
_ = let__17821

tmp18087 := Call(__e, PrimFunc(symshen_4in_1_6), let__17820)


let__17822 := tmp18087
_ = let__17822

tmp18088 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17818, let__17821)
}
__typedArg0 := let__17818
__typedArg1 := let__17821
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18089 := Call(__e, PrimFunc(symshen_4comb), let__17822, tmp18088)


ifres18084 = tmp18089


}

ifres18079 = ifres18084


}

let__17816 := ifres18079
_ = let__17816

tmp18107 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17816)


if True == tmp18107 {
tmp18092 := Call(__e, PrimFunc(sym_5_b_6), V164)


let__17824 := tmp18092
_ = let__17824

tmp18103 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17824)


var ifres18093 Obj

if True == tmp18103 {
tmp18094 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres18093 = tmp18094


} else {
tmp18095 := Call(__e, PrimFunc(symshen_4_5_1out), let__17824)


let__17825 := tmp18095
_ = let__17825

tmp18096 := Call(__e, PrimFunc(symshen_4in_1_6), let__17824)


let__17826 := tmp18096
_ = let__17826

tmp18101 := Call(__e, PrimFunc(symempty_2), let__17825)


var ifres18097 Obj

if True == tmp18101 {
ifres18097 = Nil


} else {
tmp18098 := Call(__e, PrimFunc(symshen_4app), let__17825, MakeString("\n ..."), symshen_4r)


tmp18100 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("YACC syntax error here:\n "))
__typedS1, __typedOK1 := TypedString(tmp18098)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("YACC syntax error here:\n ")
__typedArg1 := tmp18098
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("YACC syntax error here:\n "))
__typedS1, __typedOK1 := TypedString(tmp18098)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("YACC syntax error here:\n ")
__typedArg1 := tmp18098
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})()

ifres18097 = tmp18100


}

tmp18102 := Call(__e, PrimFunc(symshen_4comb), let__17826, ifres18097)


ifres18093 = tmp18102


}

let__17823 := ifres18093
_ = let__17823

tmp18105 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17823)


if True == tmp18105 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__17823)
return
}


} else {
__e.Return(let__17816)
return
}


}, 1)

tmp18108 := Call(__e, ns2_1set, symshen_4_5c_1rules_6, tmp18077)


_ = tmp18108

tmp18109 := MakeNative(func(__e *ControlFlow) {
V176 := __e.Get(1)
_ = V176
tmp18110 := Call(__e, PrimFunc(symshen_4_5syntax_6), V176)


let__17828 := tmp18110
_ = let__17828

tmp18129 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17828)


var ifres18111 Obj

if True == tmp18129 {
tmp18112 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres18111 = tmp18112


} else {
tmp18113 := Call(__e, PrimFunc(symshen_4_5_1out), let__17828)


let__17829 := tmp18113
_ = let__17829

tmp18114 := Call(__e, PrimFunc(symshen_4in_1_6), let__17828)


let__17830 := tmp18114
_ = let__17830

tmp18115 := Call(__e, PrimFunc(symshen_4_5semantics_6), let__17830)


let__17831 := tmp18115
_ = let__17831

tmp18128 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17831)


var ifres18116 Obj

if True == tmp18128 {
tmp18117 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres18116 = tmp18117


} else {
tmp18118 := Call(__e, PrimFunc(symshen_4_5_1out), let__17831)


let__17832 := tmp18118
_ = let__17832

tmp18119 := Call(__e, PrimFunc(symshen_4in_1_6), let__17831)


let__17833 := tmp18119
_ = let__17833

tmp18120 := Call(__e, PrimFunc(symshen_4_5sc_6), let__17833)


let__17834 := tmp18120
_ = let__17834

tmp18127 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17834)


var ifres18121 Obj

if True == tmp18127 {
tmp18122 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres18121 = tmp18122


} else {
tmp18123 := Call(__e, PrimFunc(symshen_4in_1_6), let__17834)


let__17835 := tmp18123
_ = let__17835

tmp18124 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17832, Nil)
}
__typedArg0 := let__17832
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18125 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17829, tmp18124)
}
__typedArg0 := let__17829
__typedArg1 := tmp18124
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18126 := Call(__e, PrimFunc(symshen_4comb), let__17835, tmp18125)


ifres18121 = tmp18126


}

ifres18116 = ifres18121


}

ifres18111 = ifres18116


}

let__17827 := ifres18111
_ = let__17827

tmp18148 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17827)


if True == tmp18148 {
tmp18130 := Call(__e, PrimFunc(symshen_4_5syntax_6), V176)


let__17837 := tmp18130
_ = let__17837

tmp18144 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17837)


var ifres18131 Obj

if True == tmp18144 {
tmp18132 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres18131 = tmp18132


} else {
tmp18133 := Call(__e, PrimFunc(symshen_4_5_1out), let__17837)


let__17838 := tmp18133
_ = let__17838

tmp18134 := Call(__e, PrimFunc(symshen_4in_1_6), let__17837)


let__17839 := tmp18134
_ = let__17839

tmp18135 := Call(__e, PrimFunc(symshen_4_5sc_6), let__17839)


let__17840 := tmp18135
_ = let__17840

tmp18143 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17840)


var ifres18136 Obj

if True == tmp18143 {
tmp18137 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres18136 = tmp18137


} else {
tmp18138 := Call(__e, PrimFunc(symshen_4in_1_6), let__17840)


let__17841 := tmp18138
_ = let__17841

tmp18139 := Call(__e, PrimFunc(symshen_4autocomplete), let__17838)


tmp18140 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18139, Nil)
}
__typedArg0 := tmp18139
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18141 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17838, tmp18140)
}
__typedArg0 := let__17838
__typedArg1 := tmp18140
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18142 := Call(__e, PrimFunc(symshen_4comb), let__17841, tmp18141)


ifres18136 = tmp18142


}

ifres18131 = ifres18136


}

let__17836 := ifres18131
_ = let__17836

tmp18146 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17836)


if True == tmp18146 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__17836)
return
}


} else {
__e.Return(let__17827)
return
}


}, 1)

tmp18149 := Call(__e, ns2_1set, symshen_4_5c_1rule_6, tmp18109)


_ = tmp18149

tmp18150 := MakeNative(func(__e *ControlFlow) {
V192 := __e.Get(1)
_ = V192
tmp18179 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V192)
}
__typedArg0 := V192
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18171 Obj

if True == tmp18179 {
tmp18177 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V192)
}
__typedArg0 := V192
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18178 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp18177)
}
__typedArg0 := Nil
__typedArg1 := tmp18177
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18173 Obj

if True == tmp18178 {
tmp18175 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V192)
}
__typedArg0 := V192
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18176 := Call(__e, PrimFunc(symshen_4non_1terminal_2), tmp18175)


var ifres18174 Obj

if True == tmp18176 {
ifres18174 = True


} else {
ifres18174 = False


}

ifres18173 = ifres18174


} else {
ifres18173 = False


}

var ifres18172 Obj

if True == ifres18173 {
ifres18172 = True


} else {
ifres18172 = False


}

ifres18171 = ifres18172


} else {
ifres18171 = False


}

if True == ifres18171 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V192)
}
__typedArg0 := V192
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return
} else {
tmp18169 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V192)
}
__typedArg0 := V192
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18165 Obj

if True == tmp18169 {
tmp18167 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V192)
}
__typedArg0 := V192
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18168 := Call(__e, PrimFunc(symshen_4non_1terminal_2), tmp18167)


var ifres18166 Obj

if True == tmp18168 {
ifres18166 = True


} else {
ifres18166 = False


}

ifres18165 = ifres18166


} else {
ifres18165 = False


}

if True == ifres18165 {
tmp18151 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V192)
}
__typedArg0 := V192
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18152 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V192)
}
__typedArg0 := V192
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18153 := Call(__e, PrimFunc(symshen_4autocomplete), tmp18152)


tmp18154 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18153, Nil)
}
__typedArg0 := tmp18153
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18155 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18151, tmp18154)
}
__typedArg0 := tmp18151
__typedArg1 := tmp18154
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symappend, tmp18155)
}
__typedArg0 := symappend
__typedArg1 := tmp18155
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp18163 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V192)
}
__typedArg0 := V192
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp18163 {
tmp18156 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V192)
}
__typedArg0 := V192
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18157 := Call(__e, PrimFunc(symshen_4autocomplete), tmp18156)


tmp18158 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V192)
}
__typedArg0 := V192
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18159 := Call(__e, PrimFunc(symshen_4autocomplete), tmp18158)


tmp18160 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18159, Nil)
}
__typedArg0 := tmp18159
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18161 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18157, tmp18160)
}
__typedArg0 := tmp18157
__typedArg1 := tmp18160
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp18161)
}
__typedArg0 := symcons
__typedArg1 := tmp18161
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return(V192)
return
}


}


}


}, 1)

tmp18180 := Call(__e, ns2_1set, symshen_4autocomplete, tmp18150)


_ = tmp18180

tmp18181 := MakeNative(func(__e *ControlFlow) {
V193 := __e.Get(1)
_ = V193
tmp18187 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsymbol_2) {
return PrimIsSymbol(V193)
}
__typedArg0 := V193
return Call(__e, PrimFunc(symsymbol_2), __typedArg0)
})()

if True == tmp18187 {
tmp18183 := Call(__e, PrimFunc(symexplode), V193)


let__17842 := tmp18183
_ = let__17842

tmp18184 := MakeNative(func(__e *ControlFlow) {
Z195 := __e.Get(1)
_ = Z195
__e.TailApply(PrimFunc(symshen_4_5non_1terminal_2_6), Z195)
return
}, 1)

tmp18185 := Call(__e, PrimFunc(symcompile), tmp18184, let__17842)


if True == tmp18185 {
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

tmp18188 := Call(__e, ns2_1set, symshen_4non_1terminal_2, tmp18181)


_ = tmp18188

tmp18189 := MakeNative(func(__e *ControlFlow) {
V196 := __e.Get(1)
_ = V196
tmp18190 := Call(__e, PrimFunc(symshen_4_5packagenames_6), V196)


let__17844 := tmp18190
_ = let__17844

tmp18200 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17844)


var ifres18191 Obj

if True == tmp18200 {
tmp18192 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres18191 = tmp18192


} else {
tmp18193 := Call(__e, PrimFunc(symshen_4in_1_6), let__17844)


let__17845 := tmp18193
_ = let__17845

tmp18194 := Call(__e, PrimFunc(symshen_4_5non_1terminal_1name_6), let__17845)


let__17846 := tmp18194
_ = let__17846

tmp18199 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17846)


var ifres18195 Obj

if True == tmp18199 {
tmp18196 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres18195 = tmp18196


} else {
tmp18197 := Call(__e, PrimFunc(symshen_4in_1_6), let__17846)


let__17847 := tmp18197
_ = let__17847

tmp18198 := Call(__e, PrimFunc(symshen_4comb), let__17847, True)


ifres18195 = tmp18198


}

ifres18191 = ifres18195


}

let__17843 := ifres18191
_ = let__17843

tmp18218 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17843)


if True == tmp18218 {
tmp18201 := Call(__e, PrimFunc(symshen_4_5non_1terminal_1name_6), V196)


let__17849 := tmp18201
_ = let__17849

tmp18206 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17849)


var ifres18202 Obj

if True == tmp18206 {
tmp18203 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres18202 = tmp18203


} else {
tmp18204 := Call(__e, PrimFunc(symshen_4in_1_6), let__17849)


let__17850 := tmp18204
_ = let__17850

tmp18205 := Call(__e, PrimFunc(symshen_4comb), let__17850, True)


ifres18202 = tmp18205


}

let__17848 := ifres18202
_ = let__17848

tmp18216 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17848)


if True == tmp18216 {
tmp18207 := Call(__e, PrimFunc(sym_5_b_6), V196)


let__17852 := tmp18207
_ = let__17852

tmp18212 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17852)


var ifres18208 Obj

if True == tmp18212 {
tmp18209 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres18208 = tmp18209


} else {
tmp18210 := Call(__e, PrimFunc(symshen_4in_1_6), let__17852)


let__17853 := tmp18210
_ = let__17853

tmp18211 := Call(__e, PrimFunc(symshen_4comb), let__17853, False)


ifres18208 = tmp18211


}

let__17851 := ifres18208
_ = let__17851

tmp18214 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17851)


if True == tmp18214 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__17851)
return
}


} else {
__e.Return(let__17848)
return
}


} else {
__e.Return(let__17843)
return
}


}, 1)

tmp18219 := Call(__e, ns2_1set, symshen_4_5non_1terminal_2_6, tmp18189)


_ = tmp18219

tmp18220 := MakeNative(func(__e *ControlFlow) {
V208 := __e.Get(1)
_ = V208
tmp18221 := Call(__e, PrimFunc(symshen_4_5packagename_6), V208)


let__17855 := tmp18221
_ = let__17855

tmp18235 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17855)


var ifres18222 Obj

if True == tmp18235 {
tmp18223 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres18222 = tmp18223


} else {
tmp18224 := Call(__e, PrimFunc(symshen_4in_1_6), let__17855)


let__17856 := tmp18224
_ = let__17856

tmp18234 := Call(__e, PrimFunc(symshen_4hds_a_2), let__17856, MakeString("."))


var ifres18225 Obj

if True == tmp18234 {
tmp18226 := Call(__e, PrimFunc(symtail), let__17856)


let__17857 := tmp18226
_ = let__17857

tmp18227 := Call(__e, PrimFunc(symshen_4_5packagenames_6), let__17857)


let__17858 := tmp18227
_ = let__17858

tmp18232 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17858)


var ifres18228 Obj

if True == tmp18232 {
tmp18229 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres18228 = tmp18229


} else {
tmp18230 := Call(__e, PrimFunc(symshen_4in_1_6), let__17858)


let__17859 := tmp18230
_ = let__17859

tmp18231 := Call(__e, PrimFunc(symshen_4comb), let__17859, symshen_4skip)


ifres18228 = tmp18231


}

ifres18225 = ifres18228


} else {
tmp18233 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres18225 = tmp18233


}

ifres18222 = ifres18225


}

let__17854 := ifres18222
_ = let__17854

tmp18249 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17854)


if True == tmp18249 {
tmp18236 := Call(__e, PrimFunc(symshen_4_5packagename_6), V208)


let__17861 := tmp18236
_ = let__17861

tmp18245 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17861)


var ifres18237 Obj

if True == tmp18245 {
tmp18238 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres18237 = tmp18238


} else {
tmp18239 := Call(__e, PrimFunc(symshen_4in_1_6), let__17861)


let__17862 := tmp18239
_ = let__17862

tmp18244 := Call(__e, PrimFunc(symshen_4hds_a_2), let__17862, MakeString("."))


var ifres18240 Obj

if True == tmp18244 {
tmp18241 := Call(__e, PrimFunc(symtail), let__17862)


let__17863 := tmp18241
_ = let__17863

tmp18242 := Call(__e, PrimFunc(symshen_4comb), let__17863, symshen_4skip)


ifres18240 = tmp18242


} else {
tmp18243 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres18240 = tmp18243


}

ifres18237 = ifres18240


}

let__17860 := ifres18237
_ = let__17860

tmp18247 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17860)


if True == tmp18247 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__17860)
return
}


} else {
__e.Return(let__17854)
return
}


}, 1)

tmp18250 := Call(__e, ns2_1set, symshen_4_5packagenames_6, tmp18220)


_ = tmp18250

tmp18251 := MakeNative(func(__e *ControlFlow) {
V219 := __e.Get(1)
_ = V219
tmp18252 := Call(__e, PrimFunc(symshen_4_5packagechar_6), V219)


let__17865 := tmp18252
_ = let__17865

tmp18262 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17865)


var ifres18253 Obj

if True == tmp18262 {
tmp18254 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres18253 = tmp18254


} else {
tmp18255 := Call(__e, PrimFunc(symshen_4in_1_6), let__17865)


let__17866 := tmp18255
_ = let__17866

tmp18256 := Call(__e, PrimFunc(symshen_4_5packagename_6), let__17866)


let__17867 := tmp18256
_ = let__17867

tmp18261 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17867)


var ifres18257 Obj

if True == tmp18261 {
tmp18258 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres18257 = tmp18258


} else {
tmp18259 := Call(__e, PrimFunc(symshen_4in_1_6), let__17867)


let__17868 := tmp18259
_ = let__17868

tmp18260 := Call(__e, PrimFunc(symshen_4comb), let__17868, symshen_4skip)


ifres18257 = tmp18260


}

ifres18253 = ifres18257


}

let__17864 := ifres18253
_ = let__17864

tmp18272 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17864)


if True == tmp18272 {
tmp18263 := Call(__e, PrimFunc(sym_5e_6), V219)


let__17870 := tmp18263
_ = let__17870

tmp18268 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17870)


var ifres18264 Obj

if True == tmp18268 {
tmp18265 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres18264 = tmp18265


} else {
tmp18266 := Call(__e, PrimFunc(symshen_4in_1_6), let__17870)


let__17871 := tmp18266
_ = let__17871

tmp18267 := Call(__e, PrimFunc(symshen_4comb), let__17871, symshen_4skip)


ifres18264 = tmp18267


}

let__17869 := ifres18264
_ = let__17869

tmp18270 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17869)


if True == tmp18270 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__17869)
return
}


} else {
__e.Return(let__17864)
return
}


}, 1)

tmp18273 := Call(__e, ns2_1set, symshen_4_5packagename_6, tmp18251)


_ = tmp18273

tmp18274 := MakeNative(func(__e *ControlFlow) {
V228 := __e.Get(1)
_ = V228
tmp18284 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V228)
}
__typedArg0 := V228
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18275 Obj

if True == tmp18284 {
tmp18276 := Call(__e, PrimFunc(symhead), V228)


let__17873 := tmp18276
_ = let__17873

tmp18277 := Call(__e, PrimFunc(symtail), V228)


let__17874 := tmp18277
_ = let__17874

tmp18281 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__17873, MakeString("."))
}
__typedArg0 := let__17873
__typedArg1 := MakeString(".")
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

tmp18282 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp18281)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp18281
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres18278 Obj

if True == tmp18282 {
tmp18279 := Call(__e, PrimFunc(symshen_4comb), let__17874, symshen_4skip)


ifres18278 = tmp18279


} else {
tmp18280 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres18278 = tmp18280


}

ifres18275 = ifres18278


} else {
tmp18283 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres18275 = tmp18283


}

let__17872 := ifres18275
_ = let__17872

tmp18286 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17872)


if True == tmp18286 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__17872)
return
}


}, 1)

tmp18287 := Call(__e, ns2_1set, symshen_4_5packagechar_6, tmp18274)


_ = tmp18287

tmp18288 := MakeNative(func(__e *ControlFlow) {
V232 := __e.Get(1)
_ = V232
tmp18307 := Call(__e, PrimFunc(symshen_4hds_a_2), V232, MakeString("<"))


var ifres18289 Obj

if True == tmp18307 {
tmp18290 := Call(__e, PrimFunc(symtail), V232)


let__17876 := tmp18290
_ = let__17876

tmp18291 := Call(__e, PrimFunc(sym_5_b_6), let__17876)


let__17877 := tmp18291
_ = let__17877

tmp18305 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17877)


var ifres18292 Obj

if True == tmp18305 {
tmp18293 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres18292 = tmp18293


} else {
tmp18294 := Call(__e, PrimFunc(symshen_4_5_1out), let__17877)


let__17878 := tmp18294
_ = let__17878

tmp18295 := Call(__e, PrimFunc(symshen_4in_1_6), let__17877)


let__17879 := tmp18295
_ = let__17879

tmp18299 := Call(__e, PrimFunc(symreverse), let__17878)


let__17880 := tmp18299
_ = let__17880

tmp18304 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__17880)
}
__typedArg0 := let__17880
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18300 Obj

if True == tmp18304 {
tmp18302 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__17880)
}
__typedArg0 := let__17880
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18303 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp18302, MakeString(">"))
}
__typedArg0 := tmp18302
__typedArg1 := MakeString(">")
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18301 Obj

if True == tmp18303 {
ifres18301 = True


} else {
ifres18301 = False


}

ifres18300 = ifres18301


} else {
ifres18300 = False


}

var ifres18296 Obj

if True == ifres18300 {
tmp18297 := Call(__e, PrimFunc(symshen_4comb), let__17879, symshen_4skip)


ifres18296 = tmp18297


} else {
tmp18298 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres18296 = tmp18298


}

ifres18292 = ifres18296


}

ifres18289 = ifres18292


} else {
tmp18306 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres18289 = tmp18306


}

let__17875 := ifres18289
_ = let__17875

tmp18309 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17875)


if True == tmp18309 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__17875)
return
}


}, 1)

tmp18310 := Call(__e, ns2_1set, symshen_4_5non_1terminal_1name_6, tmp18288)


_ = tmp18310

tmp18311 := MakeNative(func(__e *ControlFlow) {
V239 := __e.Get(1)
_ = V239
tmp18312 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(";"))
}
__typedArg0 := MakeString(";")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V239, tmp18312)
}
__typedArg0 := V239
__typedArg1 := tmp18312
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp18313 := Call(__e, ns2_1set, symshen_4semicolon_2, tmp18311)


_ = tmp18313

tmp18314 := MakeNative(func(__e *ControlFlow) {
V240 := __e.Get(1)
_ = V240
tmp18323 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V240)
}
__typedArg0 := V240
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18315 Obj

if True == tmp18323 {
tmp18316 := Call(__e, PrimFunc(symhead), V240)


let__17882 := tmp18316
_ = let__17882

tmp18317 := Call(__e, PrimFunc(symtail), V240)


let__17883 := tmp18317
_ = let__17883

tmp18321 := Call(__e, PrimFunc(symshen_4colon_1equal_2), let__17882)


var ifres18318 Obj

if True == tmp18321 {
tmp18319 := Call(__e, PrimFunc(symshen_4comb), let__17883, symshen_4skip)


ifres18318 = tmp18319


} else {
tmp18320 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres18318 = tmp18320


}

ifres18315 = ifres18318


} else {
tmp18322 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres18315 = tmp18322


}

let__17881 := ifres18315
_ = let__17881

tmp18325 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17881)


if True == tmp18325 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__17881)
return
}


}, 1)

tmp18326 := Call(__e, ns2_1set, symshen_4_5colon_1equal_6, tmp18314)


_ = tmp18326

tmp18327 := MakeNative(func(__e *ControlFlow) {
V244 := __e.Get(1)
_ = V244
tmp18328 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":="))
}
__typedArg0 := MakeString(":=")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp18328, V244)
}
__typedArg0 := tmp18328
__typedArg1 := V244
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp18329 := Call(__e, ns2_1set, symshen_4colon_1equal_2, tmp18327)


_ = tmp18329

tmp18330 := MakeNative(func(__e *ControlFlow) {
V245 := __e.Get(1)
_ = V245
tmp18331 := Call(__e, PrimFunc(symshen_4_5syntax_1item_6), V245)


let__17885 := tmp18331
_ = let__17885

tmp18344 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17885)


var ifres18332 Obj

if True == tmp18344 {
tmp18333 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres18332 = tmp18333


} else {
tmp18334 := Call(__e, PrimFunc(symshen_4_5_1out), let__17885)


let__17886 := tmp18334
_ = let__17886

tmp18335 := Call(__e, PrimFunc(symshen_4in_1_6), let__17885)


let__17887 := tmp18335
_ = let__17887

tmp18336 := Call(__e, PrimFunc(symshen_4_5syntax_6), let__17887)


let__17888 := tmp18336
_ = let__17888

tmp18343 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17888)


var ifres18337 Obj

if True == tmp18343 {
tmp18338 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres18337 = tmp18338


} else {
tmp18339 := Call(__e, PrimFunc(symshen_4_5_1out), let__17888)


let__17889 := tmp18339
_ = let__17889

tmp18340 := Call(__e, PrimFunc(symshen_4in_1_6), let__17888)


let__17890 := tmp18340
_ = let__17890

tmp18341 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17886, let__17889)
}
__typedArg0 := let__17886
__typedArg1 := let__17889
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18342 := Call(__e, PrimFunc(symshen_4comb), let__17890, tmp18341)


ifres18337 = tmp18342


}

ifres18332 = ifres18337


}

let__17884 := ifres18332
_ = let__17884

tmp18356 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17884)


if True == tmp18356 {
tmp18345 := Call(__e, PrimFunc(symshen_4_5syntax_1item_6), V245)


let__17892 := tmp18345
_ = let__17892

tmp18352 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17892)


var ifres18346 Obj

if True == tmp18352 {
tmp18347 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres18346 = tmp18347


} else {
tmp18348 := Call(__e, PrimFunc(symshen_4_5_1out), let__17892)


let__17893 := tmp18348
_ = let__17893

tmp18349 := Call(__e, PrimFunc(symshen_4in_1_6), let__17892)


let__17894 := tmp18349
_ = let__17894

tmp18350 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17893, Nil)
}
__typedArg0 := let__17893
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18351 := Call(__e, PrimFunc(symshen_4comb), let__17894, tmp18350)


ifres18346 = tmp18351


}

let__17891 := ifres18346
_ = let__17891

tmp18354 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17891)


if True == tmp18354 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__17891)
return
}


} else {
__e.Return(let__17884)
return
}


}, 1)

tmp18357 := Call(__e, ns2_1set, symshen_4_5syntax_6, tmp18330)


_ = tmp18357

tmp18358 := MakeNative(func(__e *ControlFlow) {
V257 := __e.Get(1)
_ = V257
tmp18367 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V257)
}
__typedArg0 := V257
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18359 Obj

if True == tmp18367 {
tmp18360 := Call(__e, PrimFunc(symhead), V257)


let__17896 := tmp18360
_ = let__17896

tmp18361 := Call(__e, PrimFunc(symtail), V257)


let__17897 := tmp18361
_ = let__17897

tmp18365 := Call(__e, PrimFunc(symshen_4syntax_1item_2), let__17896)


var ifres18362 Obj

if True == tmp18365 {
tmp18363 := Call(__e, PrimFunc(symshen_4comb), let__17897, let__17896)


ifres18362 = tmp18363


} else {
tmp18364 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres18362 = tmp18364


}

ifres18359 = ifres18362


} else {
tmp18366 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres18359 = tmp18366


}

let__17895 := ifres18359
_ = let__17895

tmp18369 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17895)


if True == tmp18369 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__17895)
return
}


}, 1)

tmp18370 := Call(__e, ns2_1set, symshen_4_5syntax_1item_6, tmp18358)


_ = tmp18370

tmp18371 := MakeNative(func(__e *ControlFlow) {
V263 := __e.Get(1)
_ = V263
tmp18407 := Call(__e, PrimFunc(symshen_4colon_1equal_2), V263)


if True == tmp18407 {
__e.Return(False)
return
} else {
tmp18405 := Call(__e, PrimFunc(symshen_4semicolon_2), V263)


if True == tmp18405 {
__e.Return(False)
return
} else {
tmp18403 := Call(__e, PrimFunc(symatom_2), V263)


if True == tmp18403 {
__e.Return(True)
return
} else {
tmp18401 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V263)
}
__typedArg0 := V263
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18382 Obj

if True == tmp18401 {
tmp18399 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V263)
}
__typedArg0 := V263
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18400 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcons, tmp18399)
}
__typedArg0 := symcons
__typedArg1 := tmp18399
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18384 Obj

if True == tmp18400 {
tmp18397 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V263)
}
__typedArg0 := V263
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18398 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18397)
}
__typedArg0 := tmp18397
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18386 Obj

if True == tmp18398 {
tmp18394 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V263)
}
__typedArg0 := V263
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18395 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18394)
}
__typedArg0 := tmp18394
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18396 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18395)
}
__typedArg0 := tmp18395
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18388 Obj

if True == tmp18396 {
tmp18390 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V263)
}
__typedArg0 := V263
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18391 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18390)
}
__typedArg0 := tmp18390
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18392 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18391)
}
__typedArg0 := tmp18391
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18393 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp18392)
}
__typedArg0 := Nil
__typedArg1 := tmp18392
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18389 Obj

if True == tmp18393 {
ifres18389 = True


} else {
ifres18389 = False


}

ifres18388 = ifres18389


} else {
ifres18388 = False


}

var ifres18387 Obj

if True == ifres18388 {
ifres18387 = True


} else {
ifres18387 = False


}

ifres18386 = ifres18387


} else {
ifres18386 = False


}

var ifres18385 Obj

if True == ifres18386 {
ifres18385 = True


} else {
ifres18385 = False


}

ifres18384 = ifres18385


} else {
ifres18384 = False


}

var ifres18383 Obj

if True == ifres18384 {
ifres18383 = True


} else {
ifres18383 = False


}

ifres18382 = ifres18383


} else {
ifres18382 = False


}

if True == ifres18382 {
tmp18378 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V263)
}
__typedArg0 := V263
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18379 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18378)
}
__typedArg0 := tmp18378
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18380 := Call(__e, PrimFunc(symshen_4syntax_1item_2), tmp18379)


if True == tmp18380 {
tmp18373 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V263)
}
__typedArg0 := V263
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18374 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18373)
}
__typedArg0 := tmp18373
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18375 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18374)
}
__typedArg0 := tmp18374
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18376 := Call(__e, PrimFunc(symshen_4syntax_1item_2), tmp18375)


if True == tmp18376 {
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


}


}


}


}, 1)

tmp18408 := Call(__e, ns2_1set, symshen_4syntax_1item_2, tmp18371)


_ = tmp18408

tmp18409 := MakeNative(func(__e *ControlFlow) {
V264 := __e.Get(1)
_ = V264
tmp18410 := Call(__e, PrimFunc(symshen_4_5colon_1equal_6), V264)


let__17899 := tmp18410
_ = let__17899

tmp18436 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17899)


var ifres18411 Obj

if True == tmp18436 {
tmp18412 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres18411 = tmp18412


} else {
tmp18413 := Call(__e, PrimFunc(symshen_4in_1_6), let__17899)


let__17900 := tmp18413
_ = let__17900

tmp18435 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__17900)
}
__typedArg0 := let__17900
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18414 Obj

if True == tmp18435 {
tmp18415 := Call(__e, PrimFunc(symhead), let__17900)


let__17901 := tmp18415
_ = let__17901

tmp18416 := Call(__e, PrimFunc(symtail), let__17900)


let__17902 := tmp18416
_ = let__17902

tmp18433 := Call(__e, PrimFunc(symshen_4hds_a_2), let__17902, symwhere)


var ifres18417 Obj

if True == tmp18433 {
tmp18418 := Call(__e, PrimFunc(symtail), let__17902)


let__17903 := tmp18418
_ = let__17903

tmp18431 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__17903)
}
__typedArg0 := let__17903
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18419 Obj

if True == tmp18431 {
tmp18420 := Call(__e, PrimFunc(symhead), let__17903)


let__17904 := tmp18420
_ = let__17904

tmp18421 := Call(__e, PrimFunc(symtail), let__17903)


let__17905 := tmp18421
_ = let__17905

tmp18428 := Call(__e, PrimFunc(symshen_4semicolon_2), let__17901)


tmp18429 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp18428)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp18428
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres18422 Obj

if True == tmp18429 {
tmp18423 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17901, Nil)
}
__typedArg0 := let__17901
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18424 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17904, tmp18423)
}
__typedArg0 := let__17904
__typedArg1 := tmp18423
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18425 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhere, tmp18424)
}
__typedArg0 := symwhere
__typedArg1 := tmp18424
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18426 := Call(__e, PrimFunc(symshen_4comb), let__17905, tmp18425)


ifres18422 = tmp18426


} else {
tmp18427 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres18422 = tmp18427


}

ifres18419 = ifres18422


} else {
tmp18430 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres18419 = tmp18430


}

ifres18417 = ifres18419


} else {
tmp18432 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres18417 = tmp18432


}

ifres18414 = ifres18417


} else {
tmp18434 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres18414 = tmp18434


}

ifres18411 = ifres18414


}

let__17898 := ifres18411
_ = let__17898

tmp18455 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17898)


if True == tmp18455 {
tmp18437 := Call(__e, PrimFunc(symshen_4_5colon_1equal_6), V264)


let__17907 := tmp18437
_ = let__17907

tmp18451 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17907)


var ifres18438 Obj

if True == tmp18451 {
tmp18439 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres18438 = tmp18439


} else {
tmp18440 := Call(__e, PrimFunc(symshen_4in_1_6), let__17907)


let__17908 := tmp18440
_ = let__17908

tmp18450 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__17908)
}
__typedArg0 := let__17908
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18441 Obj

if True == tmp18450 {
tmp18442 := Call(__e, PrimFunc(symhead), let__17908)


let__17909 := tmp18442
_ = let__17909

tmp18443 := Call(__e, PrimFunc(symtail), let__17908)


let__17910 := tmp18443
_ = let__17910

tmp18447 := Call(__e, PrimFunc(symshen_4semicolon_2), let__17909)


tmp18448 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp18447)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp18447
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres18444 Obj

if True == tmp18448 {
tmp18445 := Call(__e, PrimFunc(symshen_4comb), let__17910, let__17909)


ifres18444 = tmp18445


} else {
tmp18446 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres18444 = tmp18446


}

ifres18441 = ifres18444


} else {
tmp18449 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres18441 = tmp18449


}

ifres18438 = ifres18441


}

let__17906 := ifres18438
_ = let__17906

tmp18453 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__17906)


if True == tmp18453 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__17906)
return
}


} else {
__e.Return(let__17898)
return
}


}, 1)

tmp18456 := Call(__e, ns2_1set, symshen_4_5semantics_6, tmp18409)


_ = tmp18456

tmp18457 := MakeNative(func(__e *ControlFlow) {
V286 := __e.Get(1)
_ = V286
V287 := __e.Get(2)
_ = V287
V288 := __e.Get(3)
_ = V288
tmp18465 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V288)
}
__typedArg0 := Nil
__typedArg1 := V288
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp18465 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4parse_1failure, Nil)
}
__typedArg0 := symshen_4parse_1failure
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return
} else {
tmp18463 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V288)
}
__typedArg0 := V288
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp18463 {
tmp18458 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V288)
}
__typedArg0 := V288
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18459 := Call(__e, PrimFunc(symshen_4c_1rule_1_6shen), V286, tmp18458, V287)


tmp18460 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V288)
}
__typedArg0 := V288
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18461 := Call(__e, PrimFunc(symshen_4c_1rules_1_6shen), V286, V287, tmp18460)


__e.TailApply(PrimFunc(symshen_4combine_1c_1code), tmp18459, tmp18461)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.c-rules->shen\n"))
}
__typedArg0 := MakeString("implementation error in shen.c-rules->shen\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 3)

tmp18466 := Call(__e, ns2_1set, symshen_4c_1rules_1_6shen, tmp18457)


_ = tmp18466

tmp18467 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symfail))
return
}, 0)

tmp18468 := Call(__e, ns2_1set, symshen_4parse_1failure, tmp18467)


_ = tmp18468

tmp18469 := MakeNative(func(__e *ControlFlow) {
V289 := __e.Get(1)
_ = V289
V290 := __e.Get(2)
_ = V290
tmp18470 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symResult, Nil)
}
__typedArg0 := symResult
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18471 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4parse_1failure_2, tmp18470)
}
__typedArg0 := symshen_4parse_1failure_2
__typedArg1 := tmp18470
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18472 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symResult, Nil)
}
__typedArg0 := symResult
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18473 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V290, tmp18472)
}
__typedArg0 := V290
__typedArg1 := tmp18472
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18474 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18471, tmp18473)
}
__typedArg0 := tmp18471
__typedArg1 := tmp18473
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18475 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp18474)
}
__typedArg0 := symif
__typedArg1 := tmp18474
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18476 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18475, Nil)
}
__typedArg0 := tmp18475
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18477 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V289, tmp18476)
}
__typedArg0 := V289
__typedArg1 := tmp18476
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18478 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symResult, tmp18477)
}
__typedArg0 := symResult
__typedArg1 := tmp18477
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp18478)
}
__typedArg0 := symlet
__typedArg1 := tmp18478
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 2)

tmp18479 := Call(__e, ns2_1set, symshen_4combine_1c_1code, tmp18469)


_ = tmp18479

tmp18480 := MakeNative(func(__e *ControlFlow) {
V297 := __e.Get(1)
_ = V297
V298 := __e.Get(2)
_ = V298
V299 := __e.Get(3)
_ = V299
tmp18494 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V298)
}
__typedArg0 := V298
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18485 Obj

if True == tmp18494 {
tmp18492 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V298)
}
__typedArg0 := V298
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18493 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18492)
}
__typedArg0 := tmp18492
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18487 Obj

if True == tmp18493 {
tmp18489 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V298)
}
__typedArg0 := V298
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18490 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18489)
}
__typedArg0 := tmp18489
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18491 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp18490)
}
__typedArg0 := Nil
__typedArg1 := tmp18490
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18488 Obj

if True == tmp18491 {
ifres18488 = True


} else {
ifres18488 = False


}

ifres18487 = ifres18488


} else {
ifres18487 = False


}

var ifres18486 Obj

if True == ifres18487 {
ifres18486 = True


} else {
ifres18486 = False


}

ifres18485 = ifres18486


} else {
ifres18485 = False


}

if True == ifres18485 {
tmp18481 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V298)
}
__typedArg0 := V298
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18482 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V298)
}
__typedArg0 := V298
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18483 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18482)
}
__typedArg0 := tmp18482
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4yacc_1syntax), V297, V299, tmp18481, tmp18483)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.c-rule->shen\n"))
}
__typedArg0 := MakeString("implementation error in shen.c-rule->shen\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}, 3)

tmp18495 := Call(__e, ns2_1set, symshen_4c_1rule_1_6shen, tmp18480)


_ = tmp18495

tmp18496 := MakeNative(func(__e *ControlFlow) {
V308 := __e.Get(1)
_ = V308
V309 := __e.Get(2)
_ = V309
V310 := __e.Get(3)
_ = V310
V311 := __e.Get(4)
_ = V311
tmp18560 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V310)
}
__typedArg0 := Nil
__typedArg1 := V310
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18538 Obj

if True == tmp18560 {
tmp18559 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V311)
}
__typedArg0 := V311
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18540 Obj

if True == tmp18559 {
tmp18557 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V311)
}
__typedArg0 := V311
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18558 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symwhere, tmp18557)
}
__typedArg0 := symwhere
__typedArg1 := tmp18557
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18542 Obj

if True == tmp18558 {
tmp18555 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V311)
}
__typedArg0 := V311
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18556 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18555)
}
__typedArg0 := tmp18555
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18544 Obj

if True == tmp18556 {
tmp18552 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V311)
}
__typedArg0 := V311
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18553 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18552)
}
__typedArg0 := tmp18552
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18554 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18553)
}
__typedArg0 := tmp18553
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18546 Obj

if True == tmp18554 {
tmp18548 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V311)
}
__typedArg0 := V311
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18549 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18548)
}
__typedArg0 := tmp18548
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18550 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18549)
}
__typedArg0 := tmp18549
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18551 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp18550)
}
__typedArg0 := Nil
__typedArg1 := tmp18550
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18547 Obj

if True == tmp18551 {
ifres18547 = True


} else {
ifres18547 = False


}

ifres18546 = ifres18547


} else {
ifres18546 = False


}

var ifres18545 Obj

if True == ifres18546 {
ifres18545 = True


} else {
ifres18545 = False


}

ifres18544 = ifres18545


} else {
ifres18544 = False


}

var ifres18543 Obj

if True == ifres18544 {
ifres18543 = True


} else {
ifres18543 = False


}

ifres18542 = ifres18543


} else {
ifres18542 = False


}

var ifres18541 Obj

if True == ifres18542 {
ifres18541 = True


} else {
ifres18541 = False


}

ifres18540 = ifres18541


} else {
ifres18540 = False


}

var ifres18539 Obj

if True == ifres18540 {
ifres18539 = True


} else {
ifres18539 = False


}

ifres18538 = ifres18539


} else {
ifres18538 = False


}

if True == ifres18538 {
tmp18497 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V311)
}
__typedArg0 := V311
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18498 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18497)
}
__typedArg0 := tmp18497
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18499 := Call(__e, PrimFunc(symshen_4process_1yacc_1semantics), tmp18498)


tmp18500 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V311)
}
__typedArg0 := V311
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18501 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18500)
}
__typedArg0 := tmp18500
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18502 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18501)
}
__typedArg0 := tmp18501
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18503 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V308, V309, Nil, tmp18502)


tmp18504 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4parse_1failure, Nil)
}
__typedArg0 := symshen_4parse_1failure
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18505 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18504, Nil)
}
__typedArg0 := tmp18504
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18506 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18503, tmp18505)
}
__typedArg0 := tmp18503
__typedArg1 := tmp18505
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18507 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18499, tmp18506)
}
__typedArg0 := tmp18499
__typedArg1 := tmp18506
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp18507)
}
__typedArg0 := symif
__typedArg1 := tmp18507
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp18536 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V310)
}
__typedArg0 := Nil
__typedArg1 := V310
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp18536 {
__e.TailApply(PrimFunc(symshen_4yacc_1semantics), V308, V309, V311)
return
} else {
tmp18534 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp18534 {
tmp18531 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18532 := Call(__e, PrimFunc(symshen_4non_1terminal_2), tmp18531)


if True == tmp18532 {
tmp18508 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18509 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4non_1terminalcode), V308, V309, tmp18508, tmp18509, V311)
return


} else {
tmp18528 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18529 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(tmp18528)
}
__typedArg0 := tmp18528
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

if True == tmp18529 {
tmp18510 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18511 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4variablecode), V308, V309, tmp18510, tmp18511, V311)
return


} else {
tmp18525 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18526 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym__, tmp18525)
}
__typedArg0 := sym__
__typedArg1 := tmp18525
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp18526 {
tmp18512 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18513 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4wildcardcode), V308, V309, tmp18512, tmp18513, V311)
return


} else {
tmp18522 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18523 := Call(__e, PrimFunc(symatom_2), tmp18522)


if True == tmp18523 {
tmp18514 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18515 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4terminalcode), V308, V309, tmp18514, tmp18515, V311)
return


} else {
tmp18519 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18520 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18519)
}
__typedArg0 := tmp18519
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp18520 {
tmp18516 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18517 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V310)
}
__typedArg0 := V310
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4conscode), V308, V309, tmp18516, tmp18517, V311)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.yacc-syntax\n"))
}
__typedArg0 := MakeString("implementation error in shen.yacc-syntax\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}


}


}


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.yacc-syntax\n"))
}
__typedArg0 := MakeString("implementation error in shen.yacc-syntax\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}


}, 4)

tmp18561 := Call(__e, ns2_1set, symshen_4yacc_1syntax, tmp18496)


_ = tmp18561

tmp18562 := MakeNative(func(__e *ControlFlow) {
V312 := __e.Get(1)
_ = V312
V313 := __e.Get(2)
_ = V313
V314 := __e.Get(3)
_ = V314
V315 := __e.Get(4)
_ = V315
V316 := __e.Get(5)
_ = V316
tmp18563 := Call(__e, PrimFunc(symconcat), symParse, V314)


let__17911 := tmp18563
_ = let__17911

tmp18564 := Call(__e, PrimFunc(symconcat), symAction, V314)


let__17912 := tmp18564
_ = let__17912

tmp18565 := Call(__e, PrimFunc(symconcat), symRemainder, V314)


let__17913 := tmp18565
_ = let__17913

tmp18566 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V313, Nil)
}
__typedArg0 := V313
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18567 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V314, tmp18566)
}
__typedArg0 := V314
__typedArg1 := tmp18566
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18568 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17911, Nil)
}
__typedArg0 := let__17911
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18569 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4parse_1failure_2, tmp18568)
}
__typedArg0 := symshen_4parse_1failure_2
__typedArg1 := tmp18568
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18570 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4parse_1failure, Nil)
}
__typedArg0 := symshen_4parse_1failure
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18571 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17911, Nil)
}
__typedArg0 := let__17911
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18572 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4in_1_6, tmp18571)
}
__typedArg0 := symshen_4in_1_6
__typedArg1 := tmp18571
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18573 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V312, let__17913, V315, V316)


tmp18574 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18573, Nil)
}
__typedArg0 := tmp18573
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18575 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18572, tmp18574)
}
__typedArg0 := tmp18572
__typedArg1 := tmp18574
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18576 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17913, tmp18575)
}
__typedArg0 := let__17913
__typedArg1 := tmp18575
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18577 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp18576)
}
__typedArg0 := symlet
__typedArg1 := tmp18576
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__17914 := tmp18577
_ = let__17914

tmp18588 := Call(__e, PrimFunc(symshen_4occurs_1check_2), V314, V316)


var ifres18585 Obj

if True == tmp18588 {
ifres18585 = True


} else {
tmp18587 := Call(__e, PrimFunc(symshen_4occurs_1check_2), let__17912, V316)


var ifres18586 Obj

if True == tmp18587 {
ifres18586 = True


} else {
ifres18586 = False


}

ifres18585 = ifres18586


}

var ifres18578 Obj

if True == ifres18585 {
tmp18579 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17911, Nil)
}
__typedArg0 := let__17911
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18580 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4_5_1out, tmp18579)
}
__typedArg0 := symshen_4_5_1out
__typedArg1 := tmp18579
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18581 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17914, Nil)
}
__typedArg0 := let__17914
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18582 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18580, tmp18581)
}
__typedArg0 := tmp18580
__typedArg1 := tmp18581
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18583 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17912, tmp18582)
}
__typedArg0 := let__17912
__typedArg1 := tmp18582
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18584 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp18583)
}
__typedArg0 := symlet
__typedArg1 := tmp18583
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

ifres18578 = tmp18584


} else {
ifres18578 = let__17914


}

tmp18589 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(ifres18578, Nil)
}
__typedArg0 := ifres18578
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18590 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18570, tmp18589)
}
__typedArg0 := tmp18570
__typedArg1 := tmp18589
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18591 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18569, tmp18590)
}
__typedArg0 := tmp18569
__typedArg1 := tmp18590
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18592 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp18591)
}
__typedArg0 := symif
__typedArg1 := tmp18591
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18593 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18592, Nil)
}
__typedArg0 := tmp18592
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18594 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18567, tmp18593)
}
__typedArg0 := tmp18567
__typedArg1 := tmp18593
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18595 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17911, tmp18594)
}
__typedArg0 := let__17911
__typedArg1 := tmp18594
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp18595)
}
__typedArg0 := symlet
__typedArg1 := tmp18595
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 5)

tmp18596 := Call(__e, ns2_1set, symshen_4non_1terminalcode, tmp18562)


_ = tmp18596

tmp18597 := MakeNative(func(__e *ControlFlow) {
V321 := __e.Get(1)
_ = V321
V322 := __e.Get(2)
_ = V322
V323 := __e.Get(3)
_ = V323
V324 := __e.Get(4)
_ = V324
V325 := __e.Get(5)
_ = V325
tmp18598 := Call(__e, PrimFunc(symgensym), symRemainder)


let__17915 := tmp18598
_ = let__17915

tmp18599 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V322, Nil)
}
__typedArg0 := V322
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18600 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons_2, tmp18599)
}
__typedArg0 := symcons_2
__typedArg1 := tmp18599
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18601 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V322, Nil)
}
__typedArg0 := V322
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18602 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtail, tmp18601)
}
__typedArg0 := symtail
__typedArg1 := tmp18601
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18603 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V321, let__17915, V324, V325)


tmp18604 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18603, Nil)
}
__typedArg0 := tmp18603
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18605 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18602, tmp18604)
}
__typedArg0 := tmp18602
__typedArg1 := tmp18604
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18606 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17915, tmp18605)
}
__typedArg0 := let__17915
__typedArg1 := tmp18605
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18607 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp18606)
}
__typedArg0 := symlet
__typedArg1 := tmp18606
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__17916 := tmp18607
_ = let__17916

tmp18615 := Call(__e, PrimFunc(symshen_4occurs_1check_2), V323, V325)


var ifres18608 Obj

if True == tmp18615 {
tmp18609 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V322, Nil)
}
__typedArg0 := V322
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18610 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhead, tmp18609)
}
__typedArg0 := symhead
__typedArg1 := tmp18609
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18611 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17916, Nil)
}
__typedArg0 := let__17916
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18612 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18610, tmp18611)
}
__typedArg0 := tmp18610
__typedArg1 := tmp18611
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18613 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V323, tmp18612)
}
__typedArg0 := V323
__typedArg1 := tmp18612
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18614 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp18613)
}
__typedArg0 := symlet
__typedArg1 := tmp18613
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

ifres18608 = tmp18614


} else {
ifres18608 = let__17916


}

tmp18616 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4parse_1failure, Nil)
}
__typedArg0 := symshen_4parse_1failure
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18617 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18616, Nil)
}
__typedArg0 := tmp18616
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18618 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(ifres18608, tmp18617)
}
__typedArg0 := ifres18608
__typedArg1 := tmp18617
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18619 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18600, tmp18618)
}
__typedArg0 := tmp18600
__typedArg1 := tmp18618
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp18619)
}
__typedArg0 := symif
__typedArg1 := tmp18619
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 5)

tmp18620 := Call(__e, ns2_1set, symshen_4variablecode, tmp18597)


_ = tmp18620

tmp18621 := MakeNative(func(__e *ControlFlow) {
V328 := __e.Get(1)
_ = V328
V329 := __e.Get(2)
_ = V329
V330 := __e.Get(3)
_ = V330
V331 := __e.Get(4)
_ = V331
V332 := __e.Get(5)
_ = V332
tmp18622 := Call(__e, PrimFunc(symgensym), symRemainder)


let__17917 := tmp18622
_ = let__17917

tmp18623 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V329, Nil)
}
__typedArg0 := V329
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18624 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons_2, tmp18623)
}
__typedArg0 := symcons_2
__typedArg1 := tmp18623
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18625 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V329, Nil)
}
__typedArg0 := V329
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18626 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtail, tmp18625)
}
__typedArg0 := symtail
__typedArg1 := tmp18625
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18627 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V328, let__17917, V331, V332)


tmp18628 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18627, Nil)
}
__typedArg0 := tmp18627
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18629 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18626, tmp18628)
}
__typedArg0 := tmp18626
__typedArg1 := tmp18628
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18630 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17917, tmp18629)
}
__typedArg0 := let__17917
__typedArg1 := tmp18629
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18631 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp18630)
}
__typedArg0 := symlet
__typedArg1 := tmp18630
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18632 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4parse_1failure, Nil)
}
__typedArg0 := symshen_4parse_1failure
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18633 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18632, Nil)
}
__typedArg0 := tmp18632
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18634 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18631, tmp18633)
}
__typedArg0 := tmp18631
__typedArg1 := tmp18633
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18635 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18624, tmp18634)
}
__typedArg0 := tmp18624
__typedArg1 := tmp18634
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp18635)
}
__typedArg0 := symif
__typedArg1 := tmp18635
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 5)

tmp18636 := Call(__e, ns2_1set, symshen_4wildcardcode, tmp18621)


_ = tmp18636

tmp18637 := MakeNative(func(__e *ControlFlow) {
V334 := __e.Get(1)
_ = V334
V335 := __e.Get(2)
_ = V335
V336 := __e.Get(3)
_ = V336
V337 := __e.Get(4)
_ = V337
V338 := __e.Get(5)
_ = V338
tmp18638 := Call(__e, PrimFunc(symgensym), symRemainder)


let__17918 := tmp18638
_ = let__17918

tmp18639 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V336, Nil)
}
__typedArg0 := V336
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18640 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V335, tmp18639)
}
__typedArg0 := V335
__typedArg1 := tmp18639
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18641 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4hds_a_2, tmp18640)
}
__typedArg0 := symshen_4hds_a_2
__typedArg1 := tmp18640
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18642 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V335, Nil)
}
__typedArg0 := V335
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18643 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtail, tmp18642)
}
__typedArg0 := symtail
__typedArg1 := tmp18642
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18644 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V334, let__17918, V337, V338)


tmp18645 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18644, Nil)
}
__typedArg0 := tmp18644
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18646 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18643, tmp18645)
}
__typedArg0 := tmp18643
__typedArg1 := tmp18645
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18647 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17918, tmp18646)
}
__typedArg0 := let__17918
__typedArg1 := tmp18646
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18648 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp18647)
}
__typedArg0 := symlet
__typedArg1 := tmp18647
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18649 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4parse_1failure, Nil)
}
__typedArg0 := symshen_4parse_1failure
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18650 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18649, Nil)
}
__typedArg0 := tmp18649
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18651 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18648, tmp18650)
}
__typedArg0 := tmp18648
__typedArg1 := tmp18650
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18652 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18641, tmp18651)
}
__typedArg0 := tmp18641
__typedArg1 := tmp18651
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp18652)
}
__typedArg0 := symif
__typedArg1 := tmp18652
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 5)

tmp18653 := Call(__e, ns2_1set, symshen_4terminalcode, tmp18637)


_ = tmp18653

tmp18654 := MakeNative(func(__e *ControlFlow) {
V347 := __e.Get(1)
_ = V347
V348 := __e.Get(2)
_ = V348
tmp18660 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V347)
}
__typedArg0 := V347
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18656 Obj

if True == tmp18660 {
tmp18658 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V347)
}
__typedArg0 := V347
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18659 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp18658, V348)
}
__typedArg0 := tmp18658
__typedArg1 := V348
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18657 Obj

if True == tmp18659 {
ifres18657 = True


} else {
ifres18657 = False


}

ifres18656 = ifres18657


} else {
ifres18656 = False


}

if True == ifres18656 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 2)

tmp18661 := Call(__e, ns2_1set, symshen_4hds_a_2, tmp18654)


_ = tmp18661

tmp18662 := MakeNative(func(__e *ControlFlow) {
V349 := __e.Get(1)
_ = V349
V350 := __e.Get(2)
_ = V350
V351 := __e.Get(3)
_ = V351
V352 := __e.Get(4)
_ = V352
V353 := __e.Get(5)
_ = V353
tmp18663 := Call(__e, PrimFunc(symgensym), symRemainder)


let__17919 := tmp18663
_ = let__17919

tmp18664 := Call(__e, PrimFunc(symgensym), symHd)


let__17920 := tmp18664
_ = let__17920

tmp18665 := Call(__e, PrimFunc(symgensym), symTl)


let__17921 := tmp18665
_ = let__17921

tmp18666 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V350, Nil)
}
__typedArg0 := V350
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18667 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4ccons_2, tmp18666)
}
__typedArg0 := symshen_4ccons_2
__typedArg1 := tmp18666
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18668 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V350, Nil)
}
__typedArg0 := V350
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18669 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhead, tmp18668)
}
__typedArg0 := symhead
__typedArg1 := tmp18668
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18670 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V350, Nil)
}
__typedArg0 := V350
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18671 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtail, tmp18670)
}
__typedArg0 := symtail
__typedArg1 := tmp18670
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18672 := Call(__e, PrimFunc(symshen_4decons), V351)


tmp18673 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5end_6, Nil)
}
__typedArg0 := sym_5end_6
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18674 := Call(__e, PrimFunc(symappend), tmp18672, tmp18673)


tmp18675 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V349, let__17921, V352, V353)


tmp18676 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18675, Nil)
}
__typedArg0 := tmp18675
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18677 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4processed, tmp18676)
}
__typedArg0 := symshen_4processed
__typedArg1 := tmp18676
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18678 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V349, let__17920, tmp18674, tmp18677)


tmp18679 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18678, Nil)
}
__typedArg0 := tmp18678
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18680 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18671, tmp18679)
}
__typedArg0 := tmp18671
__typedArg1 := tmp18679
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18681 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17921, tmp18680)
}
__typedArg0 := let__17921
__typedArg1 := tmp18680
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18682 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18669, tmp18681)
}
__typedArg0 := tmp18669
__typedArg1 := tmp18681
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18683 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17920, tmp18682)
}
__typedArg0 := let__17920
__typedArg1 := tmp18682
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18684 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp18683)
}
__typedArg0 := symlet
__typedArg1 := tmp18683
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18685 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4parse_1failure, Nil)
}
__typedArg0 := symshen_4parse_1failure
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18686 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18685, Nil)
}
__typedArg0 := tmp18685
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18687 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18684, tmp18686)
}
__typedArg0 := tmp18684
__typedArg1 := tmp18686
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18688 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18667, tmp18687)
}
__typedArg0 := tmp18667
__typedArg1 := tmp18687
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp18688)
}
__typedArg0 := symif
__typedArg1 := tmp18688
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 5)

tmp18689 := Call(__e, ns2_1set, symshen_4conscode, tmp18662)


_ = tmp18689

tmp18690 := MakeNative(func(__e *ControlFlow) {
V367 := __e.Get(1)
_ = V367
tmp18702 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V367)
}
__typedArg0 := V367
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18698 Obj

if True == tmp18702 {
tmp18700 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V367)
}
__typedArg0 := V367
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18701 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18700)
}
__typedArg0 := tmp18700
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18699 Obj

if True == tmp18701 {
ifres18699 = True


} else {
ifres18699 = False


}

ifres18698 = ifres18699


} else {
ifres18698 = False


}

if True == ifres18698 {
__e.Return(True)
return
} else {
tmp18696 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V367)
}
__typedArg0 := V367
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18692 Obj

if True == tmp18696 {
tmp18694 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V367)
}
__typedArg0 := V367
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18695 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp18694)
}
__typedArg0 := Nil
__typedArg1 := tmp18694
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18693 Obj

if True == tmp18695 {
ifres18693 = True


} else {
ifres18693 = False


}

ifres18692 = ifres18693


} else {
ifres18692 = False


}

if True == ifres18692 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp18703 := Call(__e, ns2_1set, symshen_4ccons_2, tmp18690)


_ = tmp18703

tmp18704 := MakeNative(func(__e *ControlFlow) {
V368 := __e.Get(1)
_ = V368
tmp18731 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V368)
}
__typedArg0 := V368
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18712 Obj

if True == tmp18731 {
tmp18729 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V368)
}
__typedArg0 := V368
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18730 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcons, tmp18729)
}
__typedArg0 := symcons
__typedArg1 := tmp18729
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18714 Obj

if True == tmp18730 {
tmp18727 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V368)
}
__typedArg0 := V368
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18728 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18727)
}
__typedArg0 := tmp18727
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18716 Obj

if True == tmp18728 {
tmp18724 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V368)
}
__typedArg0 := V368
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18725 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18724)
}
__typedArg0 := tmp18724
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18726 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18725)
}
__typedArg0 := tmp18725
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18718 Obj

if True == tmp18726 {
tmp18720 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V368)
}
__typedArg0 := V368
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18721 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18720)
}
__typedArg0 := tmp18720
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18722 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18721)
}
__typedArg0 := tmp18721
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18723 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp18722)
}
__typedArg0 := Nil
__typedArg1 := tmp18722
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18719 Obj

if True == tmp18723 {
ifres18719 = True


} else {
ifres18719 = False


}

ifres18718 = ifres18719


} else {
ifres18718 = False


}

var ifres18717 Obj

if True == ifres18718 {
ifres18717 = True


} else {
ifres18717 = False


}

ifres18716 = ifres18717


} else {
ifres18716 = False


}

var ifres18715 Obj

if True == ifres18716 {
ifres18715 = True


} else {
ifres18715 = False


}

ifres18714 = ifres18715


} else {
ifres18714 = False


}

var ifres18713 Obj

if True == ifres18714 {
ifres18713 = True


} else {
ifres18713 = False


}

ifres18712 = ifres18713


} else {
ifres18712 = False


}

if True == ifres18712 {
tmp18705 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V368)
}
__typedArg0 := V368
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18706 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18705)
}
__typedArg0 := tmp18705
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18707 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V368)
}
__typedArg0 := V368
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18708 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18707)
}
__typedArg0 := tmp18707
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18709 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18708)
}
__typedArg0 := tmp18708
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18710 := Call(__e, PrimFunc(symshen_4decons), tmp18709)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp18706, tmp18710)
}
__typedArg0 := tmp18706
__typedArg1 := tmp18710
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return(V368)
return
}


}, 1)

tmp18732 := Call(__e, ns2_1set, symshen_4decons, tmp18704)


_ = tmp18732

tmp18733 := MakeNative(func(__e *ControlFlow) {
V369 := __e.Get(1)
_ = V369
V370 := __e.Get(2)
_ = V370
tmp18734 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V370, Nil)
}
__typedArg0 := V370
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V369, tmp18734)
}
__typedArg0 := V369
__typedArg1 := tmp18734
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 2)

tmp18735 := Call(__e, ns2_1set, symshen_4comb, tmp18733)


_ = tmp18735

tmp18736 := MakeNative(func(__e *ControlFlow) {
V375 := __e.Get(1)
_ = V375
V376 := __e.Get(2)
_ = V376
V377 := __e.Get(3)
_ = V377
tmp18756 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V377)
}
__typedArg0 := V377
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18743 Obj

if True == tmp18756 {
tmp18754 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V377)
}
__typedArg0 := V377
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18755 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4processed, tmp18754)
}
__typedArg0 := symshen_4processed
__typedArg1 := tmp18754
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18745 Obj

if True == tmp18755 {
tmp18752 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V377)
}
__typedArg0 := V377
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18753 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18752)
}
__typedArg0 := tmp18752
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18747 Obj

if True == tmp18753 {
tmp18749 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V377)
}
__typedArg0 := V377
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18750 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18749)
}
__typedArg0 := tmp18749
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18751 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp18750)
}
__typedArg0 := Nil
__typedArg1 := tmp18750
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18748 Obj

if True == tmp18751 {
ifres18748 = True


} else {
ifres18748 = False


}

ifres18747 = ifres18748


} else {
ifres18747 = False


}

var ifres18746 Obj

if True == ifres18747 {
ifres18746 = True


} else {
ifres18746 = False


}

ifres18745 = ifres18746


} else {
ifres18745 = False


}

var ifres18744 Obj

if True == ifres18745 {
ifres18744 = True


} else {
ifres18744 = False


}

ifres18743 = ifres18744


} else {
ifres18743 = False


}

if True == ifres18743 {
tmp18737 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V377)
}
__typedArg0 := V377
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18737)
}
__typedArg0 := tmp18737
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return


} else {
tmp18738 := Call(__e, PrimFunc(symshen_4process_1yacc_1semantics), V377)


let__17922 := tmp18738
_ = let__17922

tmp18739 := Call(__e, PrimFunc(symshen_4use_1type_1info), V375, let__17922)


let__17923 := tmp18739
_ = let__17923

tmp18740 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__17923, Nil)
}
__typedArg0 := let__17923
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp18741 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V376, tmp18740)
}
__typedArg0 := V376
__typedArg1 := tmp18740
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4comb, tmp18741)
}
__typedArg0 := symshen_4comb
__typedArg1 := tmp18741
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


}, 3)

tmp18757 := Call(__e, ns2_1set, symshen_4yacc_1semantics, tmp18736)


_ = tmp18757

tmp18758 := MakeNative(func(__e *ControlFlow) {
V383 := __e.Get(1)
_ = V383
V384 := __e.Get(2)
_ = V384
tmp18946 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18767 Obj

if True == tmp18946 {
tmp18944 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18945 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_i, tmp18944)
}
__typedArg0 := sym_i
__typedArg1 := tmp18944
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18769 Obj

if True == tmp18945 {
tmp18942 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18943 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18942)
}
__typedArg0 := tmp18942
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18771 Obj

if True == tmp18943 {
tmp18939 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18940 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18939)
}
__typedArg0 := tmp18939
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18941 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18940)
}
__typedArg0 := tmp18940
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18773 Obj

if True == tmp18941 {
tmp18935 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18936 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18935)
}
__typedArg0 := tmp18935
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18937 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18936)
}
__typedArg0 := tmp18936
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18938 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlist, tmp18937)
}
__typedArg0 := symlist
__typedArg1 := tmp18937
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18775 Obj

if True == tmp18938 {
tmp18931 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18932 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18931)
}
__typedArg0 := tmp18931
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18933 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18932)
}
__typedArg0 := tmp18932
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18934 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18933)
}
__typedArg0 := tmp18933
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18777 Obj

if True == tmp18934 {
tmp18926 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18927 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18926)
}
__typedArg0 := tmp18926
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18928 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18927)
}
__typedArg0 := tmp18927
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18929 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18928)
}
__typedArg0 := tmp18928
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18930 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp18929)
}
__typedArg0 := Nil
__typedArg1 := tmp18929
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18779 Obj

if True == tmp18930 {
tmp18923 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18924 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18923)
}
__typedArg0 := tmp18923
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18925 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18924)
}
__typedArg0 := tmp18924
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18781 Obj

if True == tmp18925 {
tmp18919 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18920 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18919)
}
__typedArg0 := tmp18919
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18921 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18920)
}
__typedArg0 := tmp18920
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18922 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_1_1_6, tmp18921)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp18921
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18783 Obj

if True == tmp18922 {
tmp18915 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18916 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18915)
}
__typedArg0 := tmp18915
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18917 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18916)
}
__typedArg0 := tmp18916
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18918 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18917)
}
__typedArg0 := tmp18917
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18785 Obj

if True == tmp18918 {
tmp18910 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18911 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18910)
}
__typedArg0 := tmp18910
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18912 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18911)
}
__typedArg0 := tmp18911
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18913 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18912)
}
__typedArg0 := tmp18912
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18914 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18913)
}
__typedArg0 := tmp18913
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18787 Obj

if True == tmp18914 {
tmp18904 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18905 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18904)
}
__typedArg0 := tmp18904
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18906 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18905)
}
__typedArg0 := tmp18905
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18907 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18906)
}
__typedArg0 := tmp18906
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18908 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18907)
}
__typedArg0 := tmp18907
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18909 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symstr, tmp18908)
}
__typedArg0 := symstr
__typedArg1 := tmp18908
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18789 Obj

if True == tmp18909 {
tmp18898 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18899 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18898)
}
__typedArg0 := tmp18898
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18900 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18899)
}
__typedArg0 := tmp18899
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18901 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18900)
}
__typedArg0 := tmp18900
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18902 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18901)
}
__typedArg0 := tmp18901
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18903 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18902)
}
__typedArg0 := tmp18902
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18791 Obj

if True == tmp18903 {
tmp18891 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18892 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18891)
}
__typedArg0 := tmp18891
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18893 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18892)
}
__typedArg0 := tmp18892
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18894 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18893)
}
__typedArg0 := tmp18893
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18895 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18894)
}
__typedArg0 := tmp18894
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18896 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18895)
}
__typedArg0 := tmp18895
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18897 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18896)
}
__typedArg0 := tmp18896
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18793 Obj

if True == tmp18897 {
tmp18883 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18884 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18883)
}
__typedArg0 := tmp18883
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18885 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18884)
}
__typedArg0 := tmp18884
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18886 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18885)
}
__typedArg0 := tmp18885
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18887 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18886)
}
__typedArg0 := tmp18886
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18888 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18887)
}
__typedArg0 := tmp18887
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18889 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18888)
}
__typedArg0 := tmp18888
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18890 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlist, tmp18889)
}
__typedArg0 := symlist
__typedArg1 := tmp18889
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18795 Obj

if True == tmp18890 {
tmp18875 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18876 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18875)
}
__typedArg0 := tmp18875
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18877 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18876)
}
__typedArg0 := tmp18876
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18878 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18877)
}
__typedArg0 := tmp18877
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18879 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18878)
}
__typedArg0 := tmp18878
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18880 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18879)
}
__typedArg0 := tmp18879
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18881 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18880)
}
__typedArg0 := tmp18880
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18882 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18881)
}
__typedArg0 := tmp18881
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18797 Obj

if True == tmp18882 {
tmp18866 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18867 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18866)
}
__typedArg0 := tmp18866
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18868 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18867)
}
__typedArg0 := tmp18867
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18869 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18868)
}
__typedArg0 := tmp18868
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18870 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18869)
}
__typedArg0 := tmp18869
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18871 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18870)
}
__typedArg0 := tmp18870
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18872 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18871)
}
__typedArg0 := tmp18871
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18873 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18872)
}
__typedArg0 := tmp18872
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18874 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp18873)
}
__typedArg0 := Nil
__typedArg1 := tmp18873
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18799 Obj

if True == tmp18874 {
tmp18859 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18860 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18859)
}
__typedArg0 := tmp18859
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18861 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18860)
}
__typedArg0 := tmp18860
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18862 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18861)
}
__typedArg0 := tmp18861
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18863 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18862)
}
__typedArg0 := tmp18862
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18864 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18863)
}
__typedArg0 := tmp18863
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18865 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18864)
}
__typedArg0 := tmp18864
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18801 Obj

if True == tmp18865 {
tmp18851 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18852 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18851)
}
__typedArg0 := tmp18851
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18853 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18852)
}
__typedArg0 := tmp18852
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18854 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18853)
}
__typedArg0 := tmp18853
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18855 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18854)
}
__typedArg0 := tmp18854
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18856 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18855)
}
__typedArg0 := tmp18855
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18857 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18856)
}
__typedArg0 := tmp18856
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18858 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp18857)
}
__typedArg0 := Nil
__typedArg1 := tmp18857
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18803 Obj

if True == tmp18858 {
tmp18846 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18847 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18846)
}
__typedArg0 := tmp18846
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18848 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18847)
}
__typedArg0 := tmp18847
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18849 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18848)
}
__typedArg0 := tmp18848
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18850 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18849)
}
__typedArg0 := tmp18849
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18805 Obj

if True == tmp18850 {
tmp18840 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18841 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18840)
}
__typedArg0 := tmp18840
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18842 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18841)
}
__typedArg0 := tmp18841
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18843 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18842)
}
__typedArg0 := tmp18842
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18844 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18843)
}
__typedArg0 := tmp18843
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18845 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_j, tmp18844)
}
__typedArg0 := sym_j
__typedArg1 := tmp18844
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18807 Obj

if True == tmp18845 {
tmp18834 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18835 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18834)
}
__typedArg0 := tmp18834
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18836 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18835)
}
__typedArg0 := tmp18835
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18837 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18836)
}
__typedArg0 := tmp18836
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18838 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18837)
}
__typedArg0 := tmp18837
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18839 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp18838)
}
__typedArg0 := Nil
__typedArg1 := tmp18838
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18809 Obj

if True == tmp18839 {
tmp18821 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18822 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18821)
}
__typedArg0 := tmp18821
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18823 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18822)
}
__typedArg0 := tmp18822
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18824 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18823)
}
__typedArg0 := tmp18823
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18825 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18826 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18825)
}
__typedArg0 := tmp18825
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18827 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18826)
}
__typedArg0 := tmp18826
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18828 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18827)
}
__typedArg0 := tmp18827
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18829 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18828)
}
__typedArg0 := tmp18828
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18830 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18829)
}
__typedArg0 := tmp18829
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18831 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18830)
}
__typedArg0 := tmp18830
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18832 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18831)
}
__typedArg0 := tmp18831
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18833 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp18824, tmp18832)
}
__typedArg0 := tmp18824
__typedArg1 := tmp18832
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18811 Obj

if True == tmp18833 {
tmp18813 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18814 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18813)
}
__typedArg0 := tmp18813
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18815 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18814)
}
__typedArg0 := tmp18814
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18816 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18815)
}
__typedArg0 := tmp18815
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18817 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18816)
}
__typedArg0 := tmp18816
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18818 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18817)
}
__typedArg0 := tmp18817
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18819 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18818)
}
__typedArg0 := tmp18818
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18820 := Call(__e, PrimFunc(symshen_4monomorphic_2), tmp18819)


var ifres18812 Obj

if True == tmp18820 {
ifres18812 = True


} else {
ifres18812 = False


}

ifres18811 = ifres18812


} else {
ifres18811 = False


}

var ifres18810 Obj

if True == ifres18811 {
ifres18810 = True


} else {
ifres18810 = False


}

ifres18809 = ifres18810


} else {
ifres18809 = False


}

var ifres18808 Obj

if True == ifres18809 {
ifres18808 = True


} else {
ifres18808 = False


}

ifres18807 = ifres18808


} else {
ifres18807 = False


}

var ifres18806 Obj

if True == ifres18807 {
ifres18806 = True


} else {
ifres18806 = False


}

ifres18805 = ifres18806


} else {
ifres18805 = False


}

var ifres18804 Obj

if True == ifres18805 {
ifres18804 = True


} else {
ifres18804 = False


}

ifres18803 = ifres18804


} else {
ifres18803 = False


}

var ifres18802 Obj

if True == ifres18803 {
ifres18802 = True


} else {
ifres18802 = False


}

ifres18801 = ifres18802


} else {
ifres18801 = False


}

var ifres18800 Obj

if True == ifres18801 {
ifres18800 = True


} else {
ifres18800 = False


}

ifres18799 = ifres18800


} else {
ifres18799 = False


}

var ifres18798 Obj

if True == ifres18799 {
ifres18798 = True


} else {
ifres18798 = False


}

ifres18797 = ifres18798


} else {
ifres18797 = False


}

var ifres18796 Obj

if True == ifres18797 {
ifres18796 = True


} else {
ifres18796 = False


}

ifres18795 = ifres18796


} else {
ifres18795 = False


}

var ifres18794 Obj

if True == ifres18795 {
ifres18794 = True


} else {
ifres18794 = False


}

ifres18793 = ifres18794


} else {
ifres18793 = False


}

var ifres18792 Obj

if True == ifres18793 {
ifres18792 = True


} else {
ifres18792 = False


}

ifres18791 = ifres18792


} else {
ifres18791 = False


}

var ifres18790 Obj

if True == ifres18791 {
ifres18790 = True


} else {
ifres18790 = False


}

ifres18789 = ifres18790


} else {
ifres18789 = False


}

var ifres18788 Obj

if True == ifres18789 {
ifres18788 = True


} else {
ifres18788 = False


}

ifres18787 = ifres18788


} else {
ifres18787 = False


}

var ifres18786 Obj

if True == ifres18787 {
ifres18786 = True


} else {
ifres18786 = False


}

ifres18785 = ifres18786


} else {
ifres18785 = False


}

var ifres18784 Obj

if True == ifres18785 {
ifres18784 = True


} else {
ifres18784 = False


}

ifres18783 = ifres18784


} else {
ifres18783 = False


}

var ifres18782 Obj

if True == ifres18783 {
ifres18782 = True


} else {
ifres18782 = False


}

ifres18781 = ifres18782


} else {
ifres18781 = False


}

var ifres18780 Obj

if True == ifres18781 {
ifres18780 = True


} else {
ifres18780 = False


}

ifres18779 = ifres18780


} else {
ifres18779 = False


}

var ifres18778 Obj

if True == ifres18779 {
ifres18778 = True


} else {
ifres18778 = False


}

ifres18777 = ifres18778


} else {
ifres18777 = False


}

var ifres18776 Obj

if True == ifres18777 {
ifres18776 = True


} else {
ifres18776 = False


}

ifres18775 = ifres18776


} else {
ifres18775 = False


}

var ifres18774 Obj

if True == ifres18775 {
ifres18774 = True


} else {
ifres18774 = False


}

ifres18773 = ifres18774


} else {
ifres18773 = False


}

var ifres18772 Obj

if True == ifres18773 {
ifres18772 = True


} else {
ifres18772 = False


}

ifres18771 = ifres18772


} else {
ifres18771 = False


}

var ifres18770 Obj

if True == ifres18771 {
ifres18770 = True


} else {
ifres18770 = False


}

ifres18769 = ifres18770


} else {
ifres18769 = False


}

var ifres18768 Obj

if True == ifres18769 {
ifres18768 = True


} else {
ifres18768 = False


}

ifres18767 = ifres18768


} else {
ifres18767 = False


}

if True == ifres18767 {
tmp18759 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V383)
}
__typedArg0 := V383
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18760 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18759)
}
__typedArg0 := tmp18759
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18761 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18760)
}
__typedArg0 := tmp18760
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18762 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18761)
}
__typedArg0 := tmp18761
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18763 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18762)
}
__typedArg0 := tmp18762
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18764 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18763)
}
__typedArg0 := tmp18763
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18765 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V384, tmp18764)
}
__typedArg0 := V384
__typedArg1 := tmp18764
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtype, tmp18765)
}
__typedArg0 := symtype
__typedArg1 := tmp18765
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return(V384)
return
}


}, 2)

tmp18947 := Call(__e, ns2_1set, symshen_4use_1type_1info, tmp18758)


_ = tmp18947

tmp18948 := MakeNative(func(__e *ControlFlow) {
V387 := __e.Get(1)
_ = V387
tmp18958 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(V387)
}
__typedArg0 := V387
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

if True == tmp18958 {
__e.Return(False)
return
} else {
tmp18956 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V387)
}
__typedArg0 := V387
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp18956 {
tmp18953 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V387)
}
__typedArg0 := V387
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18954 := Call(__e, PrimFunc(symshen_4monomorphic_2), tmp18953)


if True == tmp18954 {
tmp18950 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V387)
}
__typedArg0 := V387
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18951 := Call(__e, PrimFunc(symshen_4monomorphic_2), tmp18950)


if True == tmp18951 {
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
__e.Return(True)
return
}


}


}, 1)

tmp18959 := Call(__e, ns2_1set, symshen_4monomorphic_2, tmp18948)


_ = tmp18959

tmp18960 := MakeNative(func(__e *ControlFlow) {
V388 := __e.Get(1)
_ = V388
tmp18986 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V388)
}
__typedArg0 := V388
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18968 Obj

if True == tmp18986 {
tmp18984 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V388)
}
__typedArg0 := V388
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18985 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symprotect, tmp18984)
}
__typedArg0 := symprotect
__typedArg1 := tmp18984
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18970 Obj

if True == tmp18985 {
tmp18982 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V388)
}
__typedArg0 := V388
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18983 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18982)
}
__typedArg0 := tmp18982
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18972 Obj

if True == tmp18983 {
tmp18979 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V388)
}
__typedArg0 := V388
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18980 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18979)
}
__typedArg0 := tmp18979
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18981 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp18980)
}
__typedArg0 := Nil
__typedArg1 := tmp18980
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres18974 Obj

if True == tmp18981 {
tmp18976 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V388)
}
__typedArg0 := V388
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18977 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18976)
}
__typedArg0 := tmp18976
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18978 := Call(__e, PrimFunc(symshen_4non_1terminal_2), tmp18977)


var ifres18975 Obj

if True == tmp18978 {
ifres18975 = True


} else {
ifres18975 = False


}

ifres18974 = ifres18975


} else {
ifres18974 = False


}

var ifres18973 Obj

if True == ifres18974 {
ifres18973 = True


} else {
ifres18973 = False


}

ifres18972 = ifres18973


} else {
ifres18972 = False


}

var ifres18971 Obj

if True == ifres18972 {
ifres18971 = True


} else {
ifres18971 = False


}

ifres18970 = ifres18971


} else {
ifres18970 = False


}

var ifres18969 Obj

if True == ifres18970 {
ifres18969 = True


} else {
ifres18969 = False


}

ifres18968 = ifres18969


} else {
ifres18968 = False


}

if True == ifres18968 {
tmp18961 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V388)
}
__typedArg0 := V388
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18961)
}
__typedArg0 := tmp18961
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return


} else {
tmp18966 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V388)
}
__typedArg0 := V388
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp18966 {
tmp18962 := MakeNative(func(__e *ControlFlow) {
Z389 := __e.Get(1)
_ = Z389
__e.TailApply(PrimFunc(symshen_4process_1yacc_1semantics), Z389)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp18962, V388)
return


} else {
tmp18964 := Call(__e, PrimFunc(symshen_4non_1terminal_2), V388)


if True == tmp18964 {
__e.TailApply(PrimFunc(symconcat), symAction, V388)
return
} else {
__e.Return(V388)
return
}


}


}


}, 1)

tmp18987 := Call(__e, ns2_1set, symshen_4process_1yacc_1semantics, tmp18960)


_ = tmp18987

tmp18988 := MakeNative(func(__e *ControlFlow) {
V390 := __e.Get(1)
_ = V390
tmp18989 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V390)
}
__typedArg0 := V390
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18989)
}
__typedArg0 := tmp18989
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return


}, 1)

tmp18990 := Call(__e, ns2_1set, symshen_4_5_1out, tmp18988)


_ = tmp18990

tmp18991 := MakeNative(func(__e *ControlFlow) {
V391 := __e.Get(1)
_ = V391
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V391)
}
__typedArg0 := V391
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return
}, 1)

tmp18992 := Call(__e, ns2_1set, symshen_4in_1_6, tmp18991)


_ = tmp18992

tmp18993 := MakeNative(func(__e *ControlFlow) {
V392 := __e.Get(1)
_ = V392
tmp18994 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V392, Nil)
}
__typedArg0 := V392
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Nil, tmp18994)
}
__typedArg0 := Nil
__typedArg1 := tmp18994
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp18995 := Call(__e, ns2_1set, sym_5_b_6, tmp18993)


_ = tmp18995

tmp18996 := MakeNative(func(__e *ControlFlow) {
V393 := __e.Get(1)
_ = V393
tmp18997 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Nil, Nil)
}
__typedArg0 := Nil
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V393, tmp18997)
}
__typedArg0 := V393
__typedArg1 := tmp18997
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp18998 := Call(__e, ns2_1set, sym_5e_6, tmp18996)


_ = tmp18998

tmp18999 := MakeNative(func(__e *ControlFlow) {
V396 := __e.Get(1)
_ = V396
tmp19002 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V396)
}
__typedArg0 := Nil
__typedArg1 := V396
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp19002 {
tmp19000 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Nil, Nil)
}
__typedArg0 := Nil
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Nil, tmp19000)
}
__typedArg0 := Nil
__typedArg1 := tmp19000
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

__e.TailApply(ns2_1set, sym_5end_6, tmp18999)
return




}, 0)

