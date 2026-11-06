package main

import . "github.com/pyrex41/shen-go/kl"

var InitMain = MakeNative0(func(__e *ControlFlow) Obj {
tmp18993 := MakeNative0(func(__e *ControlFlow) Obj {
tmp18994 := PrimGuard2(__e, symset, PrimSet, symshen_4_dhistory_d, Nil)

_ = tmp18994

tmp18995 := PrimGuard2(__e, symset, PrimSet, symshen_4_dtc_d, False)

_ = tmp18995

tmp18996 := Call1(__e, SymFn(symshen_4dict), MakeInteger(20000))


tmp18997 := PrimGuard2(__e, symset, PrimSet, sym_dproperty_1vector_d, tmp18996)

_ = tmp18997

tmp18998 := MakeNative1(func(__e *ControlFlow, X Obj) Obj {
return TailCall1(__e, SymFn(symshen_4macros), X)
panic("unreachable")
})

tmp18999 := GCons(__e, symshen_4macros, tmp18998)

tmp19000 := GCons(__e, tmp18999, Nil)

tmp19001 := PrimGuard2(__e, symset, PrimSet, sym_dmacros_d, tmp19000)

_ = tmp19001

tmp19002 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsource_1form_1handlers_d, Nil)

_ = tmp19002

tmp19003 := PrimGuard2(__e, symset, PrimSet, symshen_4_dgensym_d, MakeInteger(0))

_ = tmp19003

tmp19004 := PrimGuard2(__e, symset, PrimSet, symshen_4_dtracking_d, Nil)

_ = tmp19004

tmp19005 := PrimGuard2(__e, symset, PrimSet, symshen_4_dprofiled_d, Nil)

_ = tmp19005

tmp19006 := GCons(__e, symtype, Nil)

tmp19007 := GCons(__e, symshen_4input_1h_7, tmp19006)

tmp19008 := GCons(__e, symopen, tmp19007)

tmp19009 := GCons(__e, symset, tmp19008)

tmp19010 := GCons(__e, symwhere, tmp19009)

tmp19011 := GCons(__e, symlet, tmp19010)

tmp19012 := GCons(__e, symlambda, tmp19011)

tmp19013 := GCons(__e, symcons, tmp19012)

tmp19014 := GCons(__e, sym_8v, tmp19013)

tmp19015 := GCons(__e, sym_8s, tmp19014)

tmp19016 := GCons(__e, sym_8p, tmp19015)

tmp19017 := PrimGuard2(__e, symset, PrimSet, symshen_4_dspecial_d, tmp19016)

_ = tmp19017

tmp19018 := PrimGuard2(__e, symset, PrimSet, symshen_4_dextraspecial_d, Nil)

_ = tmp19018

tmp19019 := PrimGuard2(__e, symset, PrimSet, symshen_4_dspy_d, False)

_ = tmp19019

tmp19020 := PrimGuard2(__e, symset, PrimSet, symshen_4_ddatatypes_d, Nil)

_ = tmp19020

tmp19021 := PrimGuard2(__e, symset, PrimSet, symshen_4_dalldatatypes_d, Nil)

_ = tmp19021

tmp19022 := PrimGuard2(__e, symset, PrimSet, symshen_4_dshen_1type_1theory_1enabled_2_d, True)

_ = tmp19022

tmp19023 := PrimGuard2(__e, symset, PrimSet, symshen_4_dpackage_d, symnull)

_ = tmp19023

tmp19024 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsynonyms_d, Nil)

_ = tmp19024

tmp19025 := PrimGuard2(__e, symset, PrimSet, symshen_4_doccurs_d, True)

_ = tmp19025

tmp19026 := PrimGuard2(__e, symset, PrimSet, symshen_4_dfactorise_2_d, False)

_ = tmp19026

tmp19027 := PrimGuard2(__e, symset, PrimSet, symshen_4_dmaxinferences_d, MakeInteger(1000000))

_ = tmp19027

tmp19028 := PrimGuard2(__e, symset, PrimSet, sym_dmaximum_1print_1sequence_1size_d, MakeInteger(20))

_ = tmp19028

tmp19029 := PrimGuard2(__e, symset, PrimSet, symshen_4_dcall_d, MakeInteger(0))

_ = tmp19029

tmp19030 := PrimGuard2(__e, symset, PrimSet, symshen_4_dinfs_d, MakeInteger(0))

_ = tmp19030

tmp19031 := PrimGuard2(__e, symset, PrimSet, sym_dhush_d, False)

_ = tmp19031

tmp19032 := PrimGuard2(__e, symset, PrimSet, symshen_4_doptimise_d, False)

_ = tmp19032

tmp19033 := PrimGuard2(__e, symset, PrimSet, sym_dversion_d, MakeString("43.2"))

_ = tmp19033

tmp19034 := PrimGuard2(__e, symset, PrimSet, symshen_4_dnames_d, Nil)

_ = tmp19034

tmp19035 := PrimGuard2(__e, symset, PrimSet, symshen_4_dstep_d, False)

_ = tmp19035

tmp19036 := PrimGuard2(__e, symset, PrimSet, symshen_4_dit_d, MakeString(""))

_ = tmp19036

tmp19037 := PrimGuard2(__e, symset, PrimSet, symshen_4_dresidue_d, Nil)

_ = tmp19037

tmp19038 := PrimGuard2(__e, symset, PrimSet, sym_dabsolute_d, Nil)

_ = tmp19038

tmp19039 := PrimGuard2(__e, symset, PrimSet, symshen_4_dprolog_1memory_d, MakeInteger(1000))

_ = tmp19039

tmp19040 := PrimGuard2(__e, symset, PrimSet, symshen_4_dloading_2_d, False)

_ = tmp19040

tmp19041 := PrimGuard2(__e, symset, PrimSet, symshen_4_duserdefs_d, Nil)

_ = tmp19041

tmp19042 := MakeNative1(func(__e *ControlFlow, X Obj) Obj {
return X
panic("unreachable")
})

tmp19043 := PrimGuard2(__e, symset, PrimSet, symshen_4_ddemodulation_1function_d, tmp19042)

_ = tmp19043

tmp19044 := PrimGuard2(__e, symset, PrimSet, symshen_4_dcustom_1pattern_1compiler_d, False)

_ = tmp19044

tmp19045 := PrimGuard2(__e, symset, PrimSet, symshen_4_dcustom_1pattern_1reducer_d, False)

_ = tmp19045

tmp19048 := Call1(__e, SymFn(symbound_2), sym_dhome_1directory_d)


tmp19049 := GNot(__e, tmp19048)

var ifres19046 Obj

if True == tmp19049 {
tmp19047 := PrimGuard2(__e, symset, PrimSet, sym_dhome_1directory_d, MakeString(""))

ifres19046 = tmp19047


} else {
ifres19046 = symshen_4skip


}

_ = ifres19046

tmp19053 := Call1(__e, SymFn(symbound_2), sym_dsterror_d)


tmp19054 := GNot(__e, tmp19053)

var ifres19050 Obj

if True == tmp19054 {
tmp19051 := GValue(__e, sym_dstoutput_d)

tmp19052 := PrimGuard2(__e, symset, PrimSet, sym_dsterror_d, tmp19051)

ifres19050 = tmp19052


} else {
ifres19050 = symshen_4skip


}

_ = ifres19050

tmp19055 := Call1(__e, SymFn(symprolog_1memory), MakeInteger(10000))


_ = tmp19055

tmp19056 := PrimGuard2(__e, symset, PrimSet, symshen_4_dloading_2_d, False)

_ = tmp19056

tmp19057 := GCons(__e, MakeInteger(2), Nil)

tmp19058 := GCons(__e, sym_8s, tmp19057)

tmp19059 := GCons(__e, MakeInteger(2), tmp19058)

tmp19060 := GCons(__e, sym_8v, tmp19059)

tmp19061 := GCons(__e, MakeInteger(2), tmp19060)

tmp19062 := GCons(__e, sym_8p, tmp19061)

tmp19063 := GCons(__e, MakeInteger(1), tmp19062)

tmp19064 := GCons(__e, sym_5_b_6, tmp19063)

tmp19065 := GCons(__e, MakeInteger(1), tmp19064)

tmp19066 := GCons(__e, sym_5end_6, tmp19065)

tmp19067 := GCons(__e, MakeInteger(1), tmp19066)

tmp19068 := GCons(__e, sym_5e_6, tmp19067)

tmp19069 := GCons(__e, MakeInteger(2), tmp19068)

tmp19070 := GCons(__e, sym_a_a, tmp19069)

tmp19071 := GCons(__e, MakeInteger(2), tmp19070)

tmp19072 := GCons(__e, sym_1, tmp19071)

tmp19073 := GCons(__e, MakeInteger(2), tmp19072)

tmp19074 := GCons(__e, sym_c, tmp19073)

tmp19075 := GCons(__e, MakeInteger(2), tmp19074)

tmp19076 := GCons(__e, sym_d, tmp19075)

tmp19077 := GCons(__e, MakeInteger(2), tmp19076)

tmp19078 := GCons(__e, sym_7, tmp19077)

tmp19079 := GCons(__e, MakeInteger(1), tmp19078)

tmp19080 := GCons(__e, symy_1or_1n_2, tmp19079)

tmp19081 := GCons(__e, MakeInteger(2), tmp19080)

tmp19082 := GCons(__e, symwrite_1to_1file, tmp19081)

tmp19083 := GCons(__e, MakeInteger(2), tmp19082)

tmp19084 := GCons(__e, symwrite_1string, tmp19083)

tmp19085 := GCons(__e, MakeInteger(2), tmp19084)

tmp19086 := GCons(__e, symwrite_1byte, tmp19085)

tmp19087 := GCons(__e, MakeInteger(5), tmp19086)

tmp19088 := GCons(__e, symwhen, tmp19087)

tmp19089 := GCons(__e, MakeInteger(0), tmp19088)

tmp19090 := GCons(__e, symversion, tmp19089)

tmp19091 := GCons(__e, MakeInteger(5), tmp19090)

tmp19092 := GCons(__e, symvar_2, tmp19091)

tmp19093 := GCons(__e, MakeInteger(1), tmp19092)

tmp19094 := GCons(__e, symvariable_2, tmp19093)

tmp19095 := GCons(__e, MakeInteger(1), tmp19094)

tmp19096 := GCons(__e, symvalue, tmp19095)

tmp19097 := GCons(__e, MakeInteger(3), tmp19096)

tmp19098 := GCons(__e, symvector_1_6, tmp19097)

tmp19099 := GCons(__e, MakeInteger(1), tmp19098)

tmp19100 := GCons(__e, symvector_2, tmp19099)

tmp19101 := GCons(__e, MakeInteger(1), tmp19100)

tmp19102 := GCons(__e, symvector, tmp19101)

tmp19103 := GCons(__e, MakeInteger(0), tmp19102)

tmp19104 := GCons(__e, symuserdefs, tmp19103)

tmp19105 := GCons(__e, MakeInteger(2), tmp19104)

tmp19106 := GCons(__e, symupdate_1lambda_1table, tmp19105)

tmp19107 := GCons(__e, MakeInteger(1), tmp19106)

tmp19108 := GCons(__e, symundefmacro, tmp19107)

tmp19109 := GCons(__e, MakeInteger(1), tmp19108)

tmp19110 := GCons(__e, symuntrack, tmp19109)

tmp19111 := GCons(__e, MakeInteger(2), tmp19110)

tmp19112 := GCons(__e, symunion, tmp19111)

tmp19113 := GCons(__e, MakeInteger(1), tmp19112)

tmp19114 := GCons(__e, symunprofile, tmp19113)

tmp19115 := GCons(__e, MakeInteger(3), tmp19114)

tmp19116 := GCons(__e, symunput, tmp19115)

tmp19117 := GCons(__e, MakeInteger(1), tmp19116)

tmp19118 := GCons(__e, symundefmacro, tmp19117)

tmp19119 := GCons(__e, MakeInteger(1), tmp19118)

tmp19120 := GCons(__e, symunabsolute, tmp19119)

tmp19121 := GCons(__e, MakeInteger(5), tmp19120)

tmp19122 := GCons(__e, symreturn, tmp19121)

tmp19123 := GCons(__e, MakeInteger(2), tmp19122)

tmp19124 := GCons(__e, symtype, tmp19123)

tmp19125 := GCons(__e, MakeInteger(1), tmp19124)

tmp19126 := GCons(__e, symtuple_2, tmp19125)

tmp19127 := GCons(__e, MakeInteger(2), tmp19126)

tmp19128 := GCons(__e, symtrap_1error, tmp19127)

tmp19129 := GCons(__e, MakeInteger(0), tmp19128)

tmp19130 := GCons(__e, symtracked, tmp19129)

tmp19131 := GCons(__e, MakeInteger(1), tmp19130)

tmp19132 := GCons(__e, symtrack, tmp19131)

tmp19133 := GCons(__e, MakeInteger(1), tmp19132)

tmp19134 := GCons(__e, symtlstr, tmp19133)

tmp19135 := GCons(__e, MakeInteger(1), tmp19134)

tmp19136 := GCons(__e, symthaw, tmp19135)

tmp19137 := GCons(__e, MakeInteger(0), tmp19136)

tmp19138 := GCons(__e, symtc_2, tmp19137)

tmp19139 := GCons(__e, MakeInteger(1), tmp19138)

tmp19140 := GCons(__e, symtc, tmp19139)

tmp19141 := GCons(__e, MakeInteger(1), tmp19140)

tmp19142 := GCons(__e, symtl, tmp19141)

tmp19143 := GCons(__e, MakeInteger(1), tmp19142)

tmp19144 := GCons(__e, symtail, tmp19143)

tmp19145 := GCons(__e, MakeInteger(1), tmp19144)

tmp19146 := GCons(__e, symsystemf, tmp19145)

tmp19147 := GCons(__e, MakeInteger(1), tmp19146)

tmp19148 := GCons(__e, symsymbol_2, tmp19147)

tmp19149 := GCons(__e, MakeInteger(1), tmp19148)

tmp19150 := GCons(__e, symsum, tmp19149)

tmp19151 := GCons(__e, MakeInteger(3), tmp19150)

tmp19152 := GCons(__e, symsubst, tmp19151)

tmp19153 := GCons(__e, MakeInteger(1), tmp19152)

tmp19154 := GCons(__e, symstring_2, tmp19153)

tmp19155 := GCons(__e, MakeInteger(1), tmp19154)

tmp19156 := GCons(__e, symstring_1_6symbol, tmp19155)

tmp19157 := GCons(__e, MakeInteger(1), tmp19156)

tmp19158 := GCons(__e, symstring_1_6n, tmp19157)

tmp19159 := GCons(__e, MakeInteger(1), tmp19158)

tmp19160 := GCons(__e, symstr, tmp19159)

tmp19161 := GCons(__e, MakeInteger(0), tmp19160)

tmp19162 := GCons(__e, symstoutput, tmp19161)

tmp19163 := GCons(__e, MakeInteger(0), tmp19162)

tmp19164 := GCons(__e, symstinput, tmp19163)

tmp19165 := GCons(__e, MakeInteger(0), tmp19164)

tmp19166 := GCons(__e, symstep_2, tmp19165)

tmp19167 := GCons(__e, MakeInteger(1), tmp19166)

tmp19168 := GCons(__e, symstep, tmp19167)

tmp19169 := GCons(__e, MakeInteger(0), tmp19168)

tmp19170 := GCons(__e, symspy_2, tmp19169)

tmp19171 := GCons(__e, MakeInteger(1), tmp19170)

tmp19172 := GCons(__e, symspy, tmp19171)

tmp19173 := GCons(__e, MakeInteger(2), tmp19172)

tmp19174 := GCons(__e, symspecialise, tmp19173)

tmp19175 := GCons(__e, MakeInteger(1), tmp19174)

tmp19176 := GCons(__e, symsnd, tmp19175)

tmp19177 := GCons(__e, MakeInteger(1), tmp19176)

tmp19178 := GCons(__e, symsimple_1error, tmp19177)

tmp19179 := GCons(__e, MakeInteger(2), tmp19178)

tmp19180 := GCons(__e, symset, tmp19179)

tmp19181 := GCons(__e, MakeInteger(1), tmp19180)

tmp19182 := GCons(__e, symreverse, tmp19181)

tmp19183 := GCons(__e, MakeInteger(2), tmp19182)

tmp19184 := GCons(__e, symremove, tmp19183)

tmp19185 := GCons(__e, MakeInteger(1), tmp19184)

tmp19186 := GCons(__e, symshen_4unregister_1source_1form, tmp19185)

tmp19187 := GCons(__e, MakeInteger(2), tmp19186)

tmp19188 := GCons(__e, symshen_4register_1source_1form, tmp19187)

tmp19189 := GCons(__e, MakeInteger(0), tmp19188)

tmp19190 := GCons(__e, symrelease, tmp19189)

tmp19191 := GCons(__e, MakeInteger(1), tmp19190)

tmp19192 := GCons(__e, symreceive, tmp19191)

tmp19193 := GCons(__e, MakeInteger(1), tmp19192)

tmp19194 := GCons(__e, symread_1unit_1string, tmp19193)

tmp19195 := GCons(__e, MakeInteger(1), tmp19194)

tmp19196 := GCons(__e, symread_1from_1string_1unprocessed, tmp19195)

tmp19197 := GCons(__e, MakeInteger(1), tmp19196)

tmp19198 := GCons(__e, symread_1from_1string, tmp19197)

tmp19199 := GCons(__e, MakeInteger(1), tmp19198)

tmp19200 := GCons(__e, symread_1byte, tmp19199)

tmp19201 := GCons(__e, MakeInteger(1), tmp19200)

tmp19202 := GCons(__e, symread, tmp19201)

tmp19203 := GCons(__e, MakeInteger(1), tmp19202)

tmp19204 := GCons(__e, symread_1file, tmp19203)

tmp19205 := GCons(__e, MakeInteger(1), tmp19204)

tmp19206 := GCons(__e, symread_1file_1as_1bytelist, tmp19205)

tmp19207 := GCons(__e, MakeInteger(1), tmp19206)

tmp19208 := GCons(__e, symread_1file_1as_1string, tmp19207)

tmp19209 := GCons(__e, MakeInteger(4), tmp19208)

tmp19210 := GCons(__e, symput, tmp19209)

tmp19211 := GCons(__e, MakeInteger(1), tmp19210)

tmp19212 := GCons(__e, symprotect, tmp19211)

tmp19213 := GCons(__e, MakeInteger(1), tmp19212)

tmp19214 := GCons(__e, sympreclude_1all_1but, tmp19213)

tmp19215 := GCons(__e, MakeInteger(1), tmp19214)

tmp19216 := GCons(__e, sympreclude, tmp19215)

tmp19217 := GCons(__e, MakeInteger(1), tmp19216)

tmp19218 := GCons(__e, symps, tmp19217)

tmp19219 := GCons(__e, MakeInteger(2), tmp19218)

tmp19220 := GCons(__e, sympr, tmp19219)

tmp19221 := GCons(__e, MakeInteger(1), tmp19220)

tmp19222 := GCons(__e, symprofile_1results, tmp19221)

tmp19223 := GCons(__e, MakeInteger(1), tmp19222)

tmp19224 := GCons(__e, symprolog_1memory, tmp19223)

tmp19225 := GCons(__e, MakeInteger(1), tmp19224)

tmp19226 := GCons(__e, symshen_4printF, tmp19225)

tmp19227 := GCons(__e, MakeInteger(1), tmp19226)

tmp19228 := GCons(__e, symshen_4print_1freshterm, tmp19227)

tmp19229 := GCons(__e, MakeInteger(1), tmp19228)

tmp19230 := GCons(__e, symshen_4print_1prolog_1vector, tmp19229)

tmp19231 := GCons(__e, MakeInteger(1), tmp19230)

tmp19232 := GCons(__e, symprofile, tmp19231)

tmp19233 := GCons(__e, MakeInteger(1), tmp19232)

tmp19234 := GCons(__e, symprint, tmp19233)

tmp19235 := GCons(__e, MakeInteger(1), tmp19234)

tmp19236 := GCons(__e, sympreclude_1all_1but, tmp19235)

tmp19237 := GCons(__e, MakeInteger(2), tmp19236)

tmp19238 := GCons(__e, sympos, tmp19237)

tmp19239 := GCons(__e, MakeInteger(0), tmp19238)

tmp19240 := GCons(__e, symporters, tmp19239)

tmp19241 := GCons(__e, MakeInteger(0), tmp19240)

tmp19242 := GCons(__e, symport, tmp19241)

tmp19243 := GCons(__e, MakeInteger(1), tmp19242)

tmp19244 := GCons(__e, sympackage_2, tmp19243)

tmp19245 := GCons(__e, MakeInteger(3), tmp19244)

tmp19246 := GCons(__e, sympackage, tmp19245)

tmp19247 := GCons(__e, MakeInteger(0), tmp19246)

tmp19248 := GCons(__e, symos, tmp19247)

tmp19249 := GCons(__e, MakeInteger(2), tmp19248)

tmp19250 := GCons(__e, symor, tmp19249)

tmp19251 := GCons(__e, MakeInteger(0), tmp19250)

tmp19252 := GCons(__e, symoptimise_2, tmp19251)

tmp19253 := GCons(__e, MakeInteger(1), tmp19252)

tmp19254 := GCons(__e, symoptimise, tmp19253)

tmp19255 := GCons(__e, MakeInteger(2), tmp19254)

tmp19256 := GCons(__e, symopen, tmp19255)

tmp19257 := GCons(__e, MakeInteger(1), tmp19256)

tmp19258 := GCons(__e, symoccurs_1check, tmp19257)

tmp19259 := GCons(__e, MakeInteger(0), tmp19258)

tmp19260 := GCons(__e, symoccurs_2, tmp19259)

tmp19261 := GCons(__e, MakeInteger(2), tmp19260)

tmp19262 := GCons(__e, symoccurrences, tmp19261)

tmp19263 := GCons(__e, MakeInteger(1), tmp19262)

tmp19264 := GCons(__e, symoccurs_1check, tmp19263)

tmp19265 := GCons(__e, MakeInteger(1), tmp19264)

tmp19266 := GCons(__e, symnumber_2, tmp19265)

tmp19267 := GCons(__e, MakeInteger(1), tmp19266)

tmp19268 := GCons(__e, symn_1_6string, tmp19267)

tmp19269 := GCons(__e, MakeInteger(2), tmp19268)

tmp19270 := GCons(__e, symnth, tmp19269)

tmp19271 := GCons(__e, MakeInteger(1), tmp19270)

tmp19272 := GCons(__e, symnot, tmp19271)

tmp19273 := GCons(__e, MakeInteger(1), tmp19272)

tmp19274 := GCons(__e, symnl, tmp19273)

tmp19275 := GCons(__e, MakeInteger(1), tmp19274)

tmp19276 := GCons(__e, symmaxinferences, tmp19275)

tmp19277 := GCons(__e, MakeInteger(2), tmp19276)

tmp19278 := GCons(__e, symmapcan, tmp19277)

tmp19279 := GCons(__e, MakeInteger(2), tmp19278)

tmp19280 := GCons(__e, symmap, tmp19279)

tmp19281 := GCons(__e, MakeInteger(1), tmp19280)

tmp19282 := GCons(__e, symmacroexpand, tmp19281)

tmp19283 := GCons(__e, MakeInteger(1), tmp19282)

tmp19284 := GCons(__e, symvector, tmp19283)

tmp19285 := GCons(__e, MakeInteger(2), tmp19284)

tmp19286 := GCons(__e, sym_5_a, tmp19285)

tmp19287 := GCons(__e, MakeInteger(2), tmp19286)

tmp19288 := GCons(__e, sym_5, tmp19287)

tmp19289 := GCons(__e, MakeInteger(1), tmp19288)

tmp19290 := GCons(__e, symload, tmp19289)

tmp19291 := GCons(__e, MakeInteger(1), tmp19290)

tmp19292 := GCons(__e, symlist, tmp19291)

tmp19293 := GCons(__e, MakeInteger(1), tmp19292)

tmp19294 := GCons(__e, symlineread, tmp19293)

tmp19295 := GCons(__e, MakeInteger(1), tmp19294)

tmp19296 := GCons(__e, symlimit, tmp19295)

tmp19297 := GCons(__e, MakeInteger(1), tmp19296)

tmp19298 := GCons(__e, symlength, tmp19297)

tmp19299 := GCons(__e, MakeInteger(0), tmp19298)

tmp19300 := GCons(__e, symlanguage, tmp19299)

tmp19301 := GCons(__e, MakeInteger(6), tmp19300)

tmp19302 := GCons(__e, symis_b, tmp19301)

tmp19303 := GCons(__e, MakeInteger(6), tmp19302)

tmp19304 := GCons(__e, symis, tmp19303)

tmp19305 := GCons(__e, MakeInteger(0), tmp19304)

tmp19306 := GCons(__e, symit, tmp19305)

tmp19307 := GCons(__e, MakeInteger(1), tmp19306)

tmp19308 := GCons(__e, syminternal, tmp19307)

tmp19309 := GCons(__e, MakeInteger(2), tmp19308)

tmp19310 := GCons(__e, symintersection, tmp19309)

tmp19311 := GCons(__e, MakeInteger(1), tmp19310)

tmp19312 := GCons(__e, syminclude_1all_1but, tmp19311)

tmp19313 := GCons(__e, MakeInteger(0), tmp19312)

tmp19314 := GCons(__e, symimplementation, tmp19313)

tmp19315 := GCons(__e, MakeInteger(2), tmp19314)

tmp19316 := GCons(__e, syminput_7, tmp19315)

tmp19317 := GCons(__e, MakeInteger(1), tmp19316)

tmp19318 := GCons(__e, syminput, tmp19317)

tmp19319 := GCons(__e, MakeInteger(0), tmp19318)

tmp19320 := GCons(__e, syminferences, tmp19319)

tmp19321 := GCons(__e, MakeInteger(1), tmp19320)

tmp19322 := GCons(__e, symintern, tmp19321)

tmp19323 := GCons(__e, MakeInteger(1), tmp19322)

tmp19324 := GCons(__e, syminternal, tmp19323)

tmp19325 := GCons(__e, MakeInteger(1), tmp19324)

tmp19326 := GCons(__e, syminteger_2, tmp19325)

tmp19327 := GCons(__e, MakeInteger(1), tmp19326)

tmp19328 := GCons(__e, symin_1package, tmp19327)

tmp19329 := GCons(__e, MakeInteger(0), tmp19328)

tmp19330 := GCons(__e, symincluded, tmp19329)

tmp19331 := GCons(__e, MakeInteger(1), tmp19330)

tmp19332 := GCons(__e, syminclude, tmp19331)

tmp19333 := GCons(__e, MakeInteger(3), tmp19332)

tmp19334 := GCons(__e, symif, tmp19333)

tmp19335 := GCons(__e, MakeInteger(1), tmp19334)

tmp19336 := GCons(__e, symhush, tmp19335)

tmp19337 := GCons(__e, MakeInteger(0), tmp19336)

tmp19338 := GCons(__e, symhush_2, tmp19337)

tmp19339 := GCons(__e, MakeInteger(1), tmp19338)

tmp19340 := GCons(__e, symhead, tmp19339)

tmp19341 := GCons(__e, MakeInteger(1), tmp19340)

tmp19342 := GCons(__e, symhdstr, tmp19341)

tmp19343 := GCons(__e, MakeInteger(1), tmp19342)

tmp19344 := GCons(__e, symhdv, tmp19343)

tmp19345 := GCons(__e, MakeInteger(1), tmp19344)

tmp19346 := GCons(__e, symhd, tmp19345)

tmp19347 := GCons(__e, MakeInteger(2), tmp19346)

tmp19348 := GCons(__e, symhash, tmp19347)

tmp19349 := GCons(__e, MakeInteger(2), tmp19348)

tmp19350 := GCons(__e, sym_a, tmp19349)

tmp19351 := GCons(__e, MakeInteger(2), tmp19350)

tmp19352 := GCons(__e, sym_6_a, tmp19351)

tmp19353 := GCons(__e, MakeInteger(2), tmp19352)

tmp19354 := GCons(__e, sym_6, tmp19353)

tmp19355 := GCons(__e, MakeInteger(2), tmp19354)

tmp19356 := GCons(__e, sym_5_1vector, tmp19355)

tmp19357 := GCons(__e, MakeInteger(2), tmp19356)

tmp19358 := GCons(__e, sym_5_1address, tmp19357)

tmp19359 := GCons(__e, MakeInteger(3), tmp19358)

tmp19360 := GCons(__e, symaddress_1_6, tmp19359)

tmp19361 := GCons(__e, MakeInteger(1), tmp19360)

tmp19362 := GCons(__e, symget_1time, tmp19361)

tmp19363 := GCons(__e, MakeInteger(3), tmp19362)

tmp19364 := GCons(__e, symget, tmp19363)

tmp19365 := GCons(__e, MakeInteger(1), tmp19364)

tmp19366 := GCons(__e, symgensym, tmp19365)

tmp19367 := GCons(__e, MakeInteger(1), tmp19366)

tmp19368 := GCons(__e, symfunction, tmp19367)

tmp19369 := GCons(__e, MakeInteger(1), tmp19368)

tmp19370 := GCons(__e, symfn, tmp19369)

tmp19371 := GCons(__e, MakeInteger(1), tmp19370)

tmp19372 := GCons(__e, symfst, tmp19371)

tmp19373 := GCons(__e, MakeInteger(0), tmp19372)

tmp19374 := GCons(__e, symfresh, tmp19373)

tmp19375 := GCons(__e, MakeInteger(1), tmp19374)

tmp19376 := GCons(__e, symfreeze, tmp19375)

tmp19377 := GCons(__e, MakeInteger(5), tmp19376)

tmp19378 := GCons(__e, symfork, tmp19377)

tmp19379 := GCons(__e, MakeInteger(1), tmp19378)

tmp19380 := GCons(__e, symforeign, tmp19379)

tmp19381 := GCons(__e, MakeInteger(7), tmp19380)

tmp19382 := GCons(__e, symfindall, tmp19381)

tmp19383 := GCons(__e, MakeInteger(2), tmp19382)

tmp19384 := GCons(__e, symfix, tmp19383)

tmp19385 := GCons(__e, MakeInteger(0), tmp19384)

tmp19386 := GCons(__e, symfail, tmp19385)

tmp19387 := GCons(__e, MakeInteger(2), tmp19386)

tmp19388 := GCons(__e, symfail_1if, tmp19387)

tmp19389 := GCons(__e, MakeInteger(0), tmp19388)

tmp19390 := GCons(__e, symfactorise_2, tmp19389)

tmp19391 := GCons(__e, MakeInteger(1), tmp19390)

tmp19392 := GCons(__e, symfactorise, tmp19391)

tmp19393 := GCons(__e, MakeInteger(1), tmp19392)

tmp19394 := GCons(__e, symexternal, tmp19393)

tmp19395 := GCons(__e, MakeInteger(1), tmp19394)

tmp19396 := GCons(__e, symexplode, tmp19395)

tmp19397 := GCons(__e, MakeInteger(1), tmp19396)

tmp19398 := GCons(__e, symeval_1kl, tmp19397)

tmp19399 := GCons(__e, MakeInteger(1), tmp19398)

tmp19400 := GCons(__e, symeval, tmp19399)

tmp19401 := GCons(__e, MakeInteger(1), tmp19400)

tmp19402 := GCons(__e, symerror_1to_1string, tmp19401)

tmp19403 := GCons(__e, MakeInteger(1), tmp19402)

tmp19404 := GCons(__e, symexternal, tmp19403)

tmp19405 := GCons(__e, MakeInteger(1), tmp19404)

tmp19406 := GCons(__e, symenable_1type_1theory, tmp19405)

tmp19407 := GCons(__e, MakeInteger(1), tmp19406)

tmp19408 := GCons(__e, symempty_2, tmp19407)

tmp19409 := GCons(__e, MakeInteger(2), tmp19408)

tmp19410 := GCons(__e, symelement_2, tmp19409)

tmp19411 := GCons(__e, MakeInteger(2), tmp19410)

tmp19412 := GCons(__e, symdo, tmp19411)

tmp19413 := GCons(__e, MakeInteger(2), tmp19412)

tmp19414 := GCons(__e, symdifference, tmp19413)

tmp19415 := GCons(__e, MakeInteger(1), tmp19414)

tmp19416 := GCons(__e, symdestroy, tmp19415)

tmp19417 := GCons(__e, MakeInteger(2), tmp19416)

tmp19418 := GCons(__e, symdeclare, tmp19417)

tmp19419 := GCons(__e, MakeInteger(0), tmp19418)

tmp19420 := GCons(__e, symdatatypes, tmp19419)

tmp19421 := GCons(__e, MakeInteger(1), tmp19420)

tmp19422 := GCons(__e, symclose, tmp19421)

tmp19423 := GCons(__e, MakeInteger(2), tmp19422)

tmp19424 := GCons(__e, symcn, tmp19423)

tmp19425 := GCons(__e, MakeInteger(1), tmp19424)

tmp19426 := GCons(__e, symcons_2, tmp19425)

tmp19427 := GCons(__e, MakeInteger(2), tmp19426)

tmp19428 := GCons(__e, symcons, tmp19427)

tmp19429 := GCons(__e, MakeInteger(2), tmp19428)

tmp19430 := GCons(__e, symconcat, tmp19429)

tmp19431 := GCons(__e, MakeInteger(2), tmp19430)

tmp19432 := GCons(__e, symcompile, tmp19431)

tmp19433 := GCons(__e, MakeInteger(1), tmp19432)

tmp19434 := GCons(__e, symchar_1stream_2, tmp19433)

tmp19435 := GCons(__e, MakeInteger(1), tmp19434)

tmp19436 := GCons(__e, symcd, tmp19435)

tmp19437 := GCons(__e, MakeInteger(5), tmp19436)

tmp19438 := GCons(__e, symcall, tmp19437)

tmp19439 := GCons(__e, MakeInteger(6), tmp19438)

tmp19440 := GCons(__e, symbind, tmp19439)

tmp19441 := GCons(__e, MakeInteger(1), tmp19440)

tmp19442 := GCons(__e, symbound_2, tmp19441)

tmp19443 := GCons(__e, MakeInteger(1), tmp19442)

tmp19444 := GCons(__e, symbootstrap, tmp19443)

tmp19445 := GCons(__e, MakeInteger(1), tmp19444)

tmp19446 := GCons(__e, symboolean_2, tmp19445)

tmp19447 := GCons(__e, MakeInteger(1), tmp19446)

tmp19448 := GCons(__e, symatom_2, tmp19447)

tmp19449 := GCons(__e, MakeInteger(2), tmp19448)

tmp19450 := GCons(__e, symassoc, tmp19449)

tmp19451 := GCons(__e, MakeInteger(1), tmp19450)

tmp19452 := GCons(__e, symarity, tmp19451)

tmp19453 := GCons(__e, MakeInteger(2), tmp19452)

tmp19454 := GCons(__e, symappend, tmp19453)

tmp19455 := GCons(__e, MakeInteger(2), tmp19454)

tmp19456 := GCons(__e, symand, tmp19455)

tmp19457 := GCons(__e, MakeInteger(2), tmp19456)

tmp19458 := GCons(__e, symadjoin, tmp19457)

tmp19459 := GCons(__e, MakeInteger(3), tmp19458)

tmp19460 := GCons(__e, symaddress_1_6, tmp19459)

tmp19461 := GCons(__e, MakeInteger(1), tmp19460)

tmp19462 := GCons(__e, symabsvector, tmp19461)

tmp19463 := GCons(__e, MakeInteger(1), tmp19462)

tmp19464 := GCons(__e, symabsvector_2, tmp19463)

tmp19465 := GCons(__e, MakeInteger(1), tmp19464)

tmp19466 := GCons(__e, symabsolute, tmp19465)

tmp19467 := GCons(__e, MakeInteger(0), tmp19466)

tmp19468 := GCons(__e, symabort, tmp19467)

tmp19469 := Call1(__e, SymFn(symshen_4initialise_1arity_1table), tmp19468)


_ = tmp19469

tmp19470 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":"))

tmp19471 := PrimGuard1(__e, symintern, PrimIntern, MakeString(";"))

tmp19472 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":="))

tmp19473 := PrimGuard1(__e, symintern, PrimIntern, MakeString(","))

tmp19474 := Call1(__e, SymFn(symvector), MakeInteger(0))


tmp19475 := PrimGuard1(__e, symintern, PrimIntern, MakeString("bar!"))

tmp19476 := GCons(__e, symabort, Nil)

tmp19477 := GCons(__e, symabsolute, tmp19476)

tmp19478 := GCons(__e, symabsvector, tmp19477)

tmp19479 := GCons(__e, symabsvector_2, tmp19478)

tmp19480 := GCons(__e, symaddress_1_6, tmp19479)

tmp19481 := GCons(__e, sym_5_1address, tmp19480)

tmp19482 := GCons(__e, symadjoin, tmp19481)

tmp19483 := GCons(__e, symand, tmp19482)

tmp19484 := GCons(__e, symappend, tmp19483)

tmp19485 := GCons(__e, symarity, tmp19484)

tmp19486 := GCons(__e, symassoc, tmp19485)

tmp19487 := GCons(__e, symassertz, tmp19486)

tmp19488 := GCons(__e, symasserta, tmp19487)

tmp19489 := GCons(__e, symatom_2, tmp19488)

tmp19490 := GCons(__e, symstep_2, tmp19489)

tmp19491 := GCons(__e, symspy_2, tmp19490)

tmp19492 := GCons(__e, tmp19475, tmp19491)

tmp19493 := GCons(__e, symbootstrap, tmp19492)

tmp19494 := GCons(__e, symboolean, tmp19493)

tmp19495 := GCons(__e, symboolean_2, tmp19494)

tmp19496 := GCons(__e, symbound_2, tmp19495)

tmp19497 := GCons(__e, symbind, tmp19496)

tmp19498 := GCons(__e, symclose, tmp19497)

tmp19499 := GCons(__e, symchar_1stream_2, tmp19498)

tmp19500 := GCons(__e, symcall, tmp19499)

tmp19501 := GCons(__e, symcases, tmp19500)

tmp19502 := GCons(__e, symcd, tmp19501)

tmp19503 := GCons(__e, symcompile, tmp19502)

tmp19504 := GCons(__e, symconcat, tmp19503)

tmp19505 := GCons(__e, symcond, tmp19504)

tmp19506 := GCons(__e, symcons, tmp19505)

tmp19507 := GCons(__e, symcons_2, tmp19506)

tmp19508 := GCons(__e, symcn, tmp19507)

tmp19509 := GCons(__e, symctxt, tmp19508)

tmp19510 := GCons(__e, symdatatypes, tmp19509)

tmp19511 := GCons(__e, symdatatype, tmp19510)

tmp19512 := GCons(__e, symdeclare, tmp19511)

tmp19513 := GCons(__e, symdefprolog, tmp19512)

tmp19514 := GCons(__e, symdefcc, tmp19513)

tmp19515 := GCons(__e, symdefmacro, tmp19514)

tmp19516 := GCons(__e, symdefine, tmp19515)

tmp19517 := GCons(__e, symdefun, tmp19516)

tmp19518 := GCons(__e, symdestroy, tmp19517)

tmp19519 := GCons(__e, symdifference, tmp19518)

tmp19520 := GCons(__e, symdo, tmp19519)

tmp19521 := GCons(__e, symelement_2, tmp19520)

tmp19522 := GCons(__e, symempty_2, tmp19521)

tmp19523 := GCons(__e, symerror, tmp19522)

tmp19524 := GCons(__e, symerror_1to_1string, tmp19523)

tmp19525 := GCons(__e, symeval, tmp19524)

tmp19526 := GCons(__e, symeval_1kl, tmp19525)

tmp19527 := GCons(__e, symexception, tmp19526)

tmp19528 := GCons(__e, symexternal, tmp19527)

tmp19529 := GCons(__e, symexplode, tmp19528)

tmp19530 := GCons(__e, symenable_1type_1theory, tmp19529)

tmp19531 := GCons(__e, False, tmp19530)

tmp19532 := GCons(__e, symfindall, tmp19531)

tmp19533 := GCons(__e, symfactorise_2, tmp19532)

tmp19534 := GCons(__e, symfactorise, tmp19533)

tmp19535 := GCons(__e, symfail_1if, tmp19534)

tmp19536 := GCons(__e, symfail, tmp19535)

tmp19537 := GCons(__e, symfile, tmp19536)

tmp19538 := GCons(__e, symfix, tmp19537)

tmp19539 := GCons(__e, symforeign, tmp19538)

tmp19540 := GCons(__e, symfork, tmp19539)

tmp19541 := GCons(__e, symfresh, tmp19540)

tmp19542 := GCons(__e, symfreeze, tmp19541)

tmp19543 := GCons(__e, symfst, tmp19542)

tmp19544 := GCons(__e, symfunction, tmp19543)

tmp19545 := GCons(__e, symfn, tmp19544)

tmp19546 := GCons(__e, symgensym, tmp19545)

tmp19547 := GCons(__e, symget_1time, tmp19546)

tmp19548 := GCons(__e, symget, tmp19547)

tmp19549 := GCons(__e, symhash, tmp19548)

tmp19550 := GCons(__e, symhdstr, tmp19549)

tmp19551 := GCons(__e, symhdv, tmp19550)

tmp19552 := GCons(__e, symhd, tmp19551)

tmp19553 := GCons(__e, symhead, tmp19552)

tmp19554 := GCons(__e, symhush, tmp19553)

tmp19555 := GCons(__e, symhush_2, tmp19554)

tmp19556 := GCons(__e, symif, tmp19555)

tmp19557 := GCons(__e, symimplementation, tmp19556)

tmp19558 := GCons(__e, syminternal, tmp19557)

tmp19559 := GCons(__e, symin_1package, tmp19558)

tmp19560 := GCons(__e, symin, tmp19559)

tmp19561 := GCons(__e, symis_b, tmp19560)

tmp19562 := GCons(__e, symis, tmp19561)

tmp19563 := GCons(__e, symit, tmp19562)

tmp19564 := GCons(__e, syminclude_1all_1but, tmp19563)

tmp19565 := GCons(__e, syminclude, tmp19564)

tmp19566 := GCons(__e, symincluded, tmp19565)

tmp19567 := GCons(__e, syminline, tmp19566)

tmp19568 := GCons(__e, syminput_7, tmp19567)

tmp19569 := GCons(__e, syminput, tmp19568)

tmp19570 := GCons(__e, syminteger_2, tmp19569)

tmp19571 := GCons(__e, symintern, tmp19570)

tmp19572 := GCons(__e, syminferences, tmp19571)

tmp19573 := GCons(__e, symintersection, tmp19572)

tmp19574 := GCons(__e, symis, tmp19573)

tmp19575 := GCons(__e, symlanguage, tmp19574)

tmp19576 := GCons(__e, symlambda, tmp19575)

tmp19577 := GCons(__e, symlazy, tmp19576)

tmp19578 := GCons(__e, symlet, tmp19577)

tmp19579 := GCons(__e, symlength, tmp19578)

tmp19580 := GCons(__e, symlimit, tmp19579)

tmp19581 := GCons(__e, symlineread, tmp19580)

tmp19582 := GCons(__e, symlist, tmp19581)

tmp19583 := GCons(__e, symloaded, tmp19582)

tmp19584 := GCons(__e, symload, tmp19583)

tmp19585 := GCons(__e, symmake_1string, tmp19584)

tmp19586 := GCons(__e, symmap, tmp19585)

tmp19587 := GCons(__e, symmapcan, tmp19586)

tmp19588 := GCons(__e, symmaxinferences, tmp19587)

tmp19589 := GCons(__e, symmacroexpand, tmp19588)

tmp19590 := GCons(__e, symmode, tmp19589)

tmp19591 := GCons(__e, symnl, tmp19590)

tmp19592 := GCons(__e, symnot, tmp19591)

tmp19593 := GCons(__e, symnth, tmp19592)

tmp19594 := GCons(__e, symnull, tmp19593)

tmp19595 := GCons(__e, symnumber, tmp19594)

tmp19596 := GCons(__e, symnumber_2, tmp19595)

tmp19597 := GCons(__e, symn_1_6string, tmp19596)

tmp19598 := GCons(__e, symoccurs_2, tmp19597)

tmp19599 := GCons(__e, symoccurs_1check, tmp19598)

tmp19600 := GCons(__e, symoccurrences, tmp19599)

tmp19601 := GCons(__e, symopen, tmp19600)

tmp19602 := GCons(__e, symoptimise_2, tmp19601)

tmp19603 := GCons(__e, symoptimise, tmp19602)

tmp19604 := GCons(__e, symor, tmp19603)

tmp19605 := GCons(__e, symos, tmp19604)

tmp19606 := GCons(__e, symout, tmp19605)

tmp19607 := GCons(__e, symoutput, tmp19606)

tmp19608 := GCons(__e, sympackage, tmp19607)

tmp19609 := GCons(__e, symport, tmp19608)

tmp19610 := GCons(__e, symporters, tmp19609)

tmp19611 := GCons(__e, sympos, tmp19610)

tmp19612 := GCons(__e, sympr, tmp19611)

tmp19613 := GCons(__e, symprint, tmp19612)

tmp19614 := GCons(__e, symprolog_1memory, tmp19613)

tmp19615 := GCons(__e, symprofile, tmp19614)

tmp19616 := GCons(__e, symprofile_1results, tmp19615)

tmp19617 := GCons(__e, symprotect, tmp19616)

tmp19618 := GCons(__e, symprolog_2, tmp19617)

tmp19619 := GCons(__e, symps, tmp19618)

tmp19620 := GCons(__e, sympreclude_1all_1but, tmp19619)

tmp19621 := GCons(__e, sympreclude, tmp19620)

tmp19622 := GCons(__e, symput, tmp19621)

tmp19623 := GCons(__e, sympackage_2, tmp19622)

tmp19624 := GCons(__e, symread_1unit_1string, tmp19623)

tmp19625 := GCons(__e, symread_1from_1string_1unprocessed, tmp19624)

tmp19626 := GCons(__e, symread_1from_1string, tmp19625)

tmp19627 := GCons(__e, symread_1byte, tmp19626)

tmp19628 := GCons(__e, symread_1file_1as_1string, tmp19627)

tmp19629 := GCons(__e, symread_1file_1as_1bytelist, tmp19628)

tmp19630 := GCons(__e, symread_1file, tmp19629)

tmp19631 := GCons(__e, symreceive, tmp19630)

tmp19632 := GCons(__e, symread, tmp19631)

tmp19633 := GCons(__e, symrelease, tmp19632)

tmp19634 := GCons(__e, symremove, tmp19633)

tmp19635 := GCons(__e, symretract, tmp19634)

tmp19636 := GCons(__e, symreverse, tmp19635)

tmp19637 := GCons(__e, symrun, tmp19636)

tmp19638 := GCons(__e, symstr, tmp19637)

tmp19639 := GCons(__e, symsave, tmp19638)

tmp19640 := GCons(__e, symset, tmp19639)

tmp19641 := GCons(__e, symsimple_1error, tmp19640)

tmp19642 := GCons(__e, symsnd, tmp19641)

tmp19643 := GCons(__e, symspecialise, tmp19642)

tmp19644 := GCons(__e, symspy, tmp19643)

tmp19645 := GCons(__e, symsqts, tmp19644)

tmp19646 := GCons(__e, symstep, tmp19645)

tmp19647 := GCons(__e, symstoutput, tmp19646)

tmp19648 := GCons(__e, symsterror, tmp19647)

tmp19649 := GCons(__e, symstinput, tmp19648)

tmp19650 := GCons(__e, symstring, tmp19649)

tmp19651 := GCons(__e, symstream, tmp19650)

tmp19652 := GCons(__e, symstring_1_6n, tmp19651)

tmp19653 := GCons(__e, symstring_2, tmp19652)

tmp19654 := GCons(__e, symsubst, tmp19653)

tmp19655 := GCons(__e, symsum, tmp19654)

tmp19656 := GCons(__e, symstring_1_6symbol, tmp19655)

tmp19657 := GCons(__e, symsymbol_2, tmp19656)

tmp19658 := GCons(__e, symsymbol, tmp19657)

tmp19659 := GCons(__e, symsystem_1S_2, tmp19658)

tmp19660 := GCons(__e, symsynonyms, tmp19659)

tmp19661 := GCons(__e, symsystemf, tmp19660)

tmp19662 := GCons(__e, symtail, tmp19661)

tmp19663 := GCons(__e, symtlv, tmp19662)

tmp19664 := GCons(__e, symtlstr, tmp19663)

tmp19665 := GCons(__e, symtl, tmp19664)

tmp19666 := GCons(__e, symtc, tmp19665)

tmp19667 := GCons(__e, symtc_2, tmp19666)

tmp19668 := GCons(__e, symthaw, tmp19667)

tmp19669 := GCons(__e, symtime, tmp19668)

tmp19670 := GCons(__e, symtrack, tmp19669)

tmp19671 := GCons(__e, symtracked, tmp19670)

tmp19672 := GCons(__e, symtrap_1error, tmp19671)

tmp19673 := GCons(__e, True, tmp19672)

tmp19674 := GCons(__e, symtuple_2, tmp19673)

tmp19675 := GCons(__e, symtype, tmp19674)

tmp19676 := GCons(__e, symreturn, tmp19675)

tmp19677 := GCons(__e, symunabsolute, tmp19676)

tmp19678 := GCons(__e, symundefmacro, tmp19677)

tmp19679 := GCons(__e, symunprofile, tmp19678)

tmp19680 := GCons(__e, symunput, tmp19679)

tmp19681 := GCons(__e, symunion, tmp19680)

tmp19682 := GCons(__e, symunix, tmp19681)

tmp19683 := GCons(__e, symunit, tmp19682)

tmp19684 := GCons(__e, symuntrack, tmp19683)

tmp19685 := GCons(__e, symunspecialise, tmp19684)

tmp19686 := GCons(__e, symupdate_1lambda_1table, tmp19685)

tmp19687 := GCons(__e, symu_b, tmp19686)

tmp19688 := GCons(__e, symuserdefs, tmp19687)

tmp19689 := GCons(__e, symvector_2, tmp19688)

tmp19690 := GCons(__e, symvector, tmp19689)

tmp19691 := GCons(__e, sym_5_1vector, tmp19690)

tmp19692 := GCons(__e, symvector_1_6, tmp19691)

tmp19693 := GCons(__e, symvalue, tmp19692)

tmp19694 := GCons(__e, symvar_2, tmp19693)

tmp19695 := GCons(__e, symvariable_2, tmp19694)

tmp19696 := GCons(__e, symverified, tmp19695)

tmp19697 := GCons(__e, symversion, tmp19696)

tmp19698 := GCons(__e, symwarn, tmp19697)

tmp19699 := GCons(__e, symwhen, tmp19698)

tmp19700 := GCons(__e, symwhere, tmp19699)

tmp19701 := GCons(__e, symwrite_1byte, tmp19700)

tmp19702 := GCons(__e, symwrite_1to_1file, tmp19701)

tmp19703 := GCons(__e, symwrite_1string, tmp19702)

tmp19704 := GCons(__e, symy_1or_1n_2, tmp19703)

tmp19705 := GCons(__e, tmp19474, tmp19704)

tmp19706 := GCons(__e, sym_6_6, tmp19705)

tmp19707 := GCons(__e, sym_5, tmp19706)

tmp19708 := GCons(__e, sym_5_a, tmp19707)

tmp19709 := GCons(__e, sym_7, tmp19708)

tmp19710 := GCons(__e, sym_d, tmp19709)

tmp19711 := GCons(__e, sym_c, tmp19710)

tmp19712 := GCons(__e, sym_1, tmp19711)

tmp19713 := GCons(__e, sym_3, tmp19712)

tmp19714 := GCons(__e, sym_5end_6, tmp19713)

tmp19715 := GCons(__e, sym_5_b_6, tmp19714)

tmp19716 := GCons(__e, sym_c_4, tmp19715)

tmp19717 := GCons(__e, sym_a_a_6, tmp19716)

tmp19718 := GCons(__e, sym_6, tmp19717)

tmp19719 := GCons(__e, sym_6_a, tmp19718)

tmp19720 := GCons(__e, sym_a, tmp19719)

tmp19721 := GCons(__e, sym_a_a, tmp19720)

tmp19722 := GCons(__e, sym_5e_6, tmp19721)

tmp19723 := GCons(__e, sym_1_6, tmp19722)

tmp19724 := GCons(__e, sym_5_1, tmp19723)

tmp19725 := GCons(__e, sym_dhush_d, tmp19724)

tmp19726 := GCons(__e, sym_dporters_d, tmp19725)

tmp19727 := GCons(__e, sym_dport_d, tmp19726)

tmp19728 := GCons(__e, sym_8s, tmp19727)

tmp19729 := GCons(__e, sym_8p, tmp19728)

tmp19730 := GCons(__e, sym_8v, tmp19729)

tmp19731 := GCons(__e, sym_dproperty_1vector_d, tmp19730)

tmp19732 := GCons(__e, sym_drelease_d, tmp19731)

tmp19733 := GCons(__e, sym_dos_d, tmp19732)

tmp19734 := GCons(__e, sym_dmacros_d, tmp19733)

tmp19735 := GCons(__e, sym_dmaximum_1print_1sequence_1size_d, tmp19734)

tmp19736 := GCons(__e, sym_dversion_d, tmp19735)

tmp19737 := GCons(__e, sym_dhome_1directory_d, tmp19736)

tmp19738 := GCons(__e, sym_dstoutput_d, tmp19737)

tmp19739 := GCons(__e, sym_dsterror_d, tmp19738)

tmp19740 := GCons(__e, sym_dstinput_d, tmp19739)

tmp19741 := GCons(__e, sym_dimplementation_d, tmp19740)

tmp19742 := GCons(__e, sym_dlanguage_d, tmp19741)

tmp19743 := GCons(__e, sym__, tmp19742)

tmp19744 := GCons(__e, tmp19473, tmp19743)

tmp19745 := GCons(__e, tmp19472, tmp19744)

tmp19746 := GCons(__e, tmp19471, tmp19745)

tmp19747 := GCons(__e, tmp19470, tmp19746)

tmp19748 := GCons(__e, sym_e_e, tmp19747)

tmp19749 := GCons(__e, sym_5_1_1, tmp19748)

tmp19750 := GCons(__e, sym_1_1_6, tmp19749)

tmp19751 := GCons(__e, sym_i, tmp19750)

tmp19752 := GCons(__e, sym_j, tmp19751)

tmp19753 := GCons(__e, sym_b, tmp19752)

tmp19754 := GValue(__e, sym_dproperty_1vector_d)

tmp19755 := Call4(__e, SymFn(symput), symshen, symshen_4external_1symbols, tmp19753, tmp19754)


_ = tmp19755

tmp19756 := PrimGuard1(__e, symabsvector, PrimAbsvector, MakeInteger(0))

return PrimGuard2(__e, symset, PrimSet, symshen_4_dempty_1absvector_d, tmp19756)


panic("unreachable")
})

tmp19757 := Call2(__e, ns2_1set, symshen_4initialise_1environment, tmp18993)


_ = tmp19757

tmp19758 := MakeNative0(func(__e *ControlFlow) Obj {
tmp19759 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, Nil)

_ = tmp19759

tmp19760 := MakeNative1(func(__e *ControlFlow, V6073 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6069 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6070 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6071 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6072 Obj) Obj {
tmp19761 := Call1(__e, SymFn(symshen_4newpv), B6069)


let__18906 := tmp19761
_ = let__18906

tmp19762 := Call1(__e, SymFn(symshen_4newpv), B6069)


let__18907 := tmp19762
_ = let__18907

tmp19763 := GCons(__e, let__18907, Nil)

tmp19764 := GCons(__e, symlist, tmp19763)

tmp19765 := GCons(__e, tmp19764, Nil)

tmp19766 := GCons(__e, sym_1_1_6, tmp19765)

tmp19767 := GCons(__e, let__18906, tmp19766)

tmp19768 := GCons(__e, symsymbol, Nil)

tmp19769 := GCons(__e, sym_1_1_6, tmp19768)

tmp19770 := GCons(__e, tmp19767, tmp19769)

tmp19771 := GCons(__e, tmp19770, Nil)

tmp19772 := GCons(__e, sym_1_1_6, tmp19771)

tmp19773 := GCons(__e, symsymbol, tmp19772)

tmp19774 := Call6(__e, SymFn(symis_b), V6073, tmp19773, B6069, L6070, Key6071, C6072)


tmp19775 := Call2(__e, SymFn(symshen_4gc), B6069, tmp19774)


return TailCall2(__e, SymFn(symshen_4gc), B6069, tmp19775)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19776 := GValue(__e, symshen_4_dsigf_d)

tmp19777 := Call3(__e, SymFn(symshen_4assoc_1_6), symshen_4register_1source_1form, tmp19760, tmp19776)


tmp19778 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19777)

_ = tmp19778

tmp19779 := MakeNative1(func(__e *ControlFlow, V6078 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6074 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6075 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6076 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6077 Obj) Obj {
tmp19780 := GCons(__e, symsymbol, Nil)

tmp19781 := GCons(__e, sym_1_1_6, tmp19780)

tmp19782 := GCons(__e, symsymbol, tmp19781)

return TailCall6(__e, SymFn(symis_b), V6078, tmp19782, B6074, L6075, Key6076, C6077)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19783 := GValue(__e, symshen_4_dsigf_d)

tmp19784 := Call3(__e, SymFn(symshen_4assoc_1_6), symshen_4unregister_1source_1form, tmp19779, tmp19783)


tmp19785 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19784)

_ = tmp19785

tmp19786 := MakeNative1(func(__e *ControlFlow, V6083 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6079 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6080 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6081 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6082 Obj) Obj {
tmp19787 := Call1(__e, SymFn(symshen_4newpv), B6079)


let__18908 := tmp19787
_ = let__18908

tmp19788 := GCons(__e, let__18908, Nil)

tmp19789 := GCons(__e, sym_1_1_6, tmp19788)

tmp19790 := Call6(__e, SymFn(symis_b), V6083, tmp19789, B6079, L6080, Key6081, C6082)


return TailCall2(__e, SymFn(symshen_4gc), B6079, tmp19790)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19791 := GValue(__e, symshen_4_dsigf_d)

tmp19792 := Call3(__e, SymFn(symshen_4assoc_1_6), symabort, tmp19786, tmp19791)


tmp19793 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19792)

_ = tmp19793

tmp19794 := MakeNative1(func(__e *ControlFlow, V6088 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6084 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6085 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6086 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6087 Obj) Obj {
tmp19795 := GCons(__e, symstring, Nil)

tmp19796 := GCons(__e, symlist, tmp19795)

tmp19797 := GCons(__e, tmp19796, Nil)

tmp19798 := GCons(__e, sym_1_1_6, tmp19797)

tmp19799 := GCons(__e, symstring, tmp19798)

return TailCall6(__e, SymFn(symis_b), V6088, tmp19799, B6084, L6085, Key6086, C6087)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19800 := GValue(__e, symshen_4_dsigf_d)

tmp19801 := Call3(__e, SymFn(symshen_4assoc_1_6), symabsolute, tmp19794, tmp19800)


tmp19802 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19801)

_ = tmp19802

tmp19803 := MakeNative1(func(__e *ControlFlow, V6093 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6089 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6090 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6091 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6092 Obj) Obj {
tmp19804 := Call1(__e, SymFn(symshen_4newpv), B6089)


let__18909 := tmp19804
_ = let__18909

tmp19805 := GCons(__e, symboolean, Nil)

tmp19806 := GCons(__e, sym_1_1_6, tmp19805)

tmp19807 := GCons(__e, let__18909, tmp19806)

tmp19808 := Call6(__e, SymFn(symis_b), V6093, tmp19807, B6089, L6090, Key6091, C6092)


return TailCall2(__e, SymFn(symshen_4gc), B6089, tmp19808)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19809 := GValue(__e, symshen_4_dsigf_d)

tmp19810 := Call3(__e, SymFn(symshen_4assoc_1_6), symabsvector_2, tmp19803, tmp19809)


tmp19811 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19810)

_ = tmp19811

tmp19812 := MakeNative1(func(__e *ControlFlow, V6099 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6095 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6096 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6097 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6098 Obj) Obj {
tmp19813 := Call1(__e, SymFn(symshen_4newpv), B6095)


let__18910 := tmp19813
_ = let__18910

tmp19814 := GCons(__e, let__18910, Nil)

tmp19815 := GCons(__e, symlist, tmp19814)

tmp19816 := GCons(__e, let__18910, Nil)

tmp19817 := GCons(__e, symlist, tmp19816)

tmp19818 := GCons(__e, tmp19817, Nil)

tmp19819 := GCons(__e, sym_1_1_6, tmp19818)

tmp19820 := GCons(__e, tmp19815, tmp19819)

tmp19821 := GCons(__e, tmp19820, Nil)

tmp19822 := GCons(__e, sym_1_1_6, tmp19821)

tmp19823 := GCons(__e, let__18910, tmp19822)

tmp19824 := Call6(__e, SymFn(symis_b), V6099, tmp19823, B6095, L6096, Key6097, C6098)


return TailCall2(__e, SymFn(symshen_4gc), B6095, tmp19824)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19825 := GValue(__e, symshen_4_dsigf_d)

tmp19826 := Call3(__e, SymFn(symshen_4assoc_1_6), symadjoin, tmp19812, tmp19825)


tmp19827 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19826)

_ = tmp19827

tmp19828 := MakeNative1(func(__e *ControlFlow, V6104 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6100 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6101 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6102 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6103 Obj) Obj {
tmp19829 := GCons(__e, symboolean, Nil)

tmp19830 := GCons(__e, sym_1_1_6, tmp19829)

tmp19831 := GCons(__e, symboolean, tmp19830)

tmp19832 := GCons(__e, tmp19831, Nil)

tmp19833 := GCons(__e, sym_1_1_6, tmp19832)

tmp19834 := GCons(__e, symboolean, tmp19833)

return TailCall6(__e, SymFn(symis_b), V6104, tmp19834, B6100, L6101, Key6102, C6103)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19835 := GValue(__e, symshen_4_dsigf_d)

tmp19836 := Call3(__e, SymFn(symshen_4assoc_1_6), symand, tmp19828, tmp19835)


tmp19837 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19836)

_ = tmp19837

tmp19838 := MakeNative1(func(__e *ControlFlow, V6109 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6105 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6106 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6107 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6108 Obj) Obj {
tmp19839 := Call1(__e, SymFn(symshen_4newpv), B6105)


let__18911 := tmp19839
_ = let__18911

tmp19840 := GCons(__e, symstring, Nil)

tmp19841 := GCons(__e, sym_1_1_6, tmp19840)

tmp19842 := GCons(__e, symsymbol, tmp19841)

tmp19843 := GCons(__e, tmp19842, Nil)

tmp19844 := GCons(__e, sym_1_1_6, tmp19843)

tmp19845 := GCons(__e, symstring, tmp19844)

tmp19846 := GCons(__e, tmp19845, Nil)

tmp19847 := GCons(__e, sym_1_1_6, tmp19846)

tmp19848 := GCons(__e, let__18911, tmp19847)

tmp19849 := Call6(__e, SymFn(symis_b), V6109, tmp19848, B6105, L6106, Key6107, C6108)


return TailCall2(__e, SymFn(symshen_4gc), B6105, tmp19849)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19850 := GValue(__e, symshen_4_dsigf_d)

tmp19851 := Call3(__e, SymFn(symshen_4assoc_1_6), symshen_4app, tmp19838, tmp19850)


tmp19852 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19851)

_ = tmp19852

tmp19853 := MakeNative1(func(__e *ControlFlow, V6116 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6112 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6113 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6114 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6115 Obj) Obj {
tmp19854 := Call1(__e, SymFn(symshen_4newpv), B6112)


let__18912 := tmp19854
_ = let__18912

tmp19855 := GCons(__e, let__18912, Nil)

tmp19856 := GCons(__e, symlist, tmp19855)

tmp19857 := GCons(__e, let__18912, Nil)

tmp19858 := GCons(__e, symlist, tmp19857)

tmp19859 := GCons(__e, let__18912, Nil)

tmp19860 := GCons(__e, symlist, tmp19859)

tmp19861 := GCons(__e, tmp19860, Nil)

tmp19862 := GCons(__e, sym_1_1_6, tmp19861)

tmp19863 := GCons(__e, tmp19858, tmp19862)

tmp19864 := GCons(__e, tmp19863, Nil)

tmp19865 := GCons(__e, sym_1_1_6, tmp19864)

tmp19866 := GCons(__e, tmp19856, tmp19865)

tmp19867 := Call6(__e, SymFn(symis_b), V6116, tmp19866, B6112, L6113, Key6114, C6115)


return TailCall2(__e, SymFn(symshen_4gc), B6112, tmp19867)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19868 := GValue(__e, symshen_4_dsigf_d)

tmp19869 := Call3(__e, SymFn(symshen_4assoc_1_6), symappend, tmp19853, tmp19868)


tmp19870 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19869)

_ = tmp19870

tmp19871 := MakeNative1(func(__e *ControlFlow, V6121 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6117 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6118 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6119 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6120 Obj) Obj {
tmp19872 := Call1(__e, SymFn(symshen_4newpv), B6117)


let__18913 := tmp19872
_ = let__18913

tmp19873 := GCons(__e, symnumber, Nil)

tmp19874 := GCons(__e, sym_1_1_6, tmp19873)

tmp19875 := GCons(__e, let__18913, tmp19874)

tmp19876 := Call6(__e, SymFn(symis_b), V6121, tmp19875, B6117, L6118, Key6119, C6120)


return TailCall2(__e, SymFn(symshen_4gc), B6117, tmp19876)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19877 := GValue(__e, symshen_4_dsigf_d)

tmp19878 := Call3(__e, SymFn(symshen_4assoc_1_6), symarity, tmp19871, tmp19877)


tmp19879 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19878)

_ = tmp19879

tmp19880 := MakeNative1(func(__e *ControlFlow, V6127 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6123 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6124 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6125 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6126 Obj) Obj {
tmp19881 := Call1(__e, SymFn(symshen_4newpv), B6123)


let__18914 := tmp19881
_ = let__18914

tmp19882 := GCons(__e, let__18914, Nil)

tmp19883 := GCons(__e, symlist, tmp19882)

tmp19884 := GCons(__e, tmp19883, Nil)

tmp19885 := GCons(__e, symlist, tmp19884)

tmp19886 := GCons(__e, let__18914, Nil)

tmp19887 := GCons(__e, symlist, tmp19886)

tmp19888 := GCons(__e, tmp19887, Nil)

tmp19889 := GCons(__e, sym_1_1_6, tmp19888)

tmp19890 := GCons(__e, tmp19885, tmp19889)

tmp19891 := GCons(__e, tmp19890, Nil)

tmp19892 := GCons(__e, sym_1_1_6, tmp19891)

tmp19893 := GCons(__e, let__18914, tmp19892)

tmp19894 := Call6(__e, SymFn(symis_b), V6127, tmp19893, B6123, L6124, Key6125, C6126)


return TailCall2(__e, SymFn(symshen_4gc), B6123, tmp19894)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19895 := GValue(__e, symshen_4_dsigf_d)

tmp19896 := Call3(__e, SymFn(symshen_4assoc_1_6), symassoc, tmp19880, tmp19895)


tmp19897 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19896)

_ = tmp19897

tmp19898 := MakeNative1(func(__e *ControlFlow, V6132 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6128 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6129 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6130 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6131 Obj) Obj {
tmp19899 := Call1(__e, SymFn(symshen_4newpv), B6128)


let__18915 := tmp19899
_ = let__18915

tmp19900 := GCons(__e, symboolean, Nil)

tmp19901 := GCons(__e, sym_1_1_6, tmp19900)

tmp19902 := GCons(__e, let__18915, tmp19901)

tmp19903 := Call6(__e, SymFn(symis_b), V6132, tmp19902, B6128, L6129, Key6130, C6131)


return TailCall2(__e, SymFn(symshen_4gc), B6128, tmp19903)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19904 := GValue(__e, symshen_4_dsigf_d)

tmp19905 := Call3(__e, SymFn(symshen_4assoc_1_6), symatom_2, tmp19898, tmp19904)


tmp19906 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19905)

_ = tmp19906

tmp19907 := MakeNative1(func(__e *ControlFlow, V6137 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6133 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6134 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6135 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6136 Obj) Obj {
tmp19908 := Call1(__e, SymFn(symshen_4newpv), B6133)


let__18916 := tmp19908
_ = let__18916

tmp19909 := GCons(__e, symboolean, Nil)

tmp19910 := GCons(__e, sym_1_1_6, tmp19909)

tmp19911 := GCons(__e, let__18916, tmp19910)

tmp19912 := Call6(__e, SymFn(symis_b), V6137, tmp19911, B6133, L6134, Key6135, C6136)


return TailCall2(__e, SymFn(symshen_4gc), B6133, tmp19912)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19913 := GValue(__e, symshen_4_dsigf_d)

tmp19914 := Call3(__e, SymFn(symshen_4assoc_1_6), symboolean_2, tmp19907, tmp19913)


tmp19915 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19914)

_ = tmp19915

tmp19916 := MakeNative1(func(__e *ControlFlow, V6142 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6138 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6139 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6140 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6141 Obj) Obj {
tmp19917 := GCons(__e, symstring, Nil)

tmp19918 := GCons(__e, sym_1_1_6, tmp19917)

tmp19919 := GCons(__e, symstring, tmp19918)

return TailCall6(__e, SymFn(symis_b), V6142, tmp19919, B6138, L6139, Key6140, C6141)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19920 := GValue(__e, symshen_4_dsigf_d)

tmp19921 := Call3(__e, SymFn(symshen_4assoc_1_6), symbootstrap, tmp19916, tmp19920)


tmp19922 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19921)

_ = tmp19922

tmp19923 := MakeNative1(func(__e *ControlFlow, V6147 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6143 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6144 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6145 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6146 Obj) Obj {
tmp19924 := GCons(__e, symboolean, Nil)

tmp19925 := GCons(__e, sym_1_1_6, tmp19924)

tmp19926 := GCons(__e, symsymbol, tmp19925)

return TailCall6(__e, SymFn(symis_b), V6147, tmp19926, B6143, L6144, Key6145, C6146)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19927 := GValue(__e, symshen_4_dsigf_d)

tmp19928 := Call3(__e, SymFn(symshen_4assoc_1_6), symbound_2, tmp19923, tmp19927)


tmp19929 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19928)

_ = tmp19929

tmp19930 := MakeNative1(func(__e *ControlFlow, V6152 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6148 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6149 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6150 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6151 Obj) Obj {
tmp19931 := Call1(__e, SymFn(symshen_4newpv), B6148)


let__18917 := tmp19931
_ = let__18917

tmp19932 := GCons(__e, let__18917, Nil)

tmp19933 := GCons(__e, symlist, tmp19932)

tmp19934 := GCons(__e, symboolean, Nil)

tmp19935 := GCons(__e, sym_1_1_6, tmp19934)

tmp19936 := GCons(__e, tmp19933, tmp19935)

tmp19937 := Call6(__e, SymFn(symis_b), V6152, tmp19936, B6148, L6149, Key6150, C6151)


return TailCall2(__e, SymFn(symshen_4gc), B6148, tmp19937)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19938 := GValue(__e, symshen_4_dsigf_d)

tmp19939 := Call3(__e, SymFn(symshen_4assoc_1_6), symshen_4ccons_2, tmp19930, tmp19938)


tmp19940 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19939)

_ = tmp19940

tmp19941 := MakeNative1(func(__e *ControlFlow, V6157 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6153 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6154 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6155 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6156 Obj) Obj {
tmp19942 := GCons(__e, symstring, Nil)

tmp19943 := GCons(__e, sym_1_1_6, tmp19942)

tmp19944 := GCons(__e, symstring, tmp19943)

return TailCall6(__e, SymFn(symis_b), V6157, tmp19944, B6153, L6154, Key6155, C6156)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19945 := GValue(__e, symshen_4_dsigf_d)

tmp19946 := Call3(__e, SymFn(symshen_4assoc_1_6), symcd, tmp19941, tmp19945)


tmp19947 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19946)

_ = tmp19947

tmp19948 := MakeNative1(func(__e *ControlFlow, V6163 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6159 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6160 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6161 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6162 Obj) Obj {
tmp19949 := Call1(__e, SymFn(symshen_4newpv), B6159)


let__18918 := tmp19949
_ = let__18918

tmp19950 := Call1(__e, SymFn(symshen_4newpv), B6159)


let__18919 := tmp19950
_ = let__18919

tmp19951 := GCons(__e, let__18918, Nil)

tmp19952 := GCons(__e, symstream, tmp19951)

tmp19953 := GCons(__e, let__18919, Nil)

tmp19954 := GCons(__e, symlist, tmp19953)

tmp19955 := GCons(__e, tmp19954, Nil)

tmp19956 := GCons(__e, sym_1_1_6, tmp19955)

tmp19957 := GCons(__e, tmp19952, tmp19956)

tmp19958 := Call6(__e, SymFn(symis_b), V6163, tmp19957, B6159, L6160, Key6161, C6162)


tmp19959 := Call2(__e, SymFn(symshen_4gc), B6159, tmp19958)


return TailCall2(__e, SymFn(symshen_4gc), B6159, tmp19959)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19960 := GValue(__e, symshen_4_dsigf_d)

tmp19961 := Call3(__e, SymFn(symshen_4assoc_1_6), symclose, tmp19948, tmp19960)


tmp19962 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19961)

_ = tmp19962

tmp19963 := MakeNative1(func(__e *ControlFlow, V6168 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6164 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6165 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6166 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6167 Obj) Obj {
tmp19964 := GCons(__e, symstring, Nil)

tmp19965 := GCons(__e, sym_1_1_6, tmp19964)

tmp19966 := GCons(__e, symstring, tmp19965)

tmp19967 := GCons(__e, tmp19966, Nil)

tmp19968 := GCons(__e, sym_1_1_6, tmp19967)

tmp19969 := GCons(__e, symstring, tmp19968)

return TailCall6(__e, SymFn(symis_b), V6168, tmp19969, B6164, L6165, Key6166, C6167)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19970 := GValue(__e, symshen_4_dsigf_d)

tmp19971 := Call3(__e, SymFn(symshen_4assoc_1_6), symcn, tmp19963, tmp19970)


tmp19972 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19971)

_ = tmp19972

tmp19973 := MakeNative1(func(__e *ControlFlow, V6173 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6169 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6170 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6171 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6172 Obj) Obj {
tmp19974 := Call1(__e, SymFn(symshen_4newpv), B6169)


let__18920 := tmp19974
_ = let__18920

tmp19975 := GCons(__e, let__18920, Nil)

tmp19976 := GCons(__e, symstream, tmp19975)

tmp19977 := GCons(__e, symboolean, Nil)

tmp19978 := GCons(__e, sym_1_1_6, tmp19977)

tmp19979 := GCons(__e, tmp19976, tmp19978)

tmp19980 := Call6(__e, SymFn(symis_b), V6173, tmp19979, B6169, L6170, Key6171, C6172)


return TailCall2(__e, SymFn(symshen_4gc), B6169, tmp19980)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19981 := GValue(__e, symshen_4_dsigf_d)

tmp19982 := Call3(__e, SymFn(symshen_4assoc_1_6), symchar_1stream_2, tmp19973, tmp19981)


tmp19983 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19982)

_ = tmp19983

tmp19984 := MakeNative1(func(__e *ControlFlow, V6180 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6176 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6177 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6178 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6179 Obj) Obj {
tmp19985 := Call1(__e, SymFn(symshen_4newpv), B6176)


let__18921 := tmp19985
_ = let__18921

tmp19986 := Call1(__e, SymFn(symshen_4newpv), B6176)


let__18922 := tmp19986
_ = let__18922

tmp19987 := GCons(__e, let__18921, Nil)

tmp19988 := GCons(__e, symlist, tmp19987)

tmp19989 := GCons(__e, let__18921, Nil)

tmp19990 := GCons(__e, symlist, tmp19989)

tmp19991 := GCons(__e, let__18922, Nil)

tmp19992 := GCons(__e, tmp19990, tmp19991)

tmp19993 := GCons(__e, symstr, tmp19992)

tmp19994 := GCons(__e, tmp19993, Nil)

tmp19995 := GCons(__e, sym_1_1_6, tmp19994)

tmp19996 := GCons(__e, tmp19988, tmp19995)

tmp19997 := GCons(__e, let__18921, Nil)

tmp19998 := GCons(__e, symlist, tmp19997)

tmp19999 := GCons(__e, let__18922, Nil)

tmp20000 := GCons(__e, sym_1_1_6, tmp19999)

tmp20001 := GCons(__e, tmp19998, tmp20000)

tmp20002 := GCons(__e, tmp20001, Nil)

tmp20003 := GCons(__e, sym_1_1_6, tmp20002)

tmp20004 := GCons(__e, tmp19996, tmp20003)

tmp20005 := Call6(__e, SymFn(symis_b), V6180, tmp20004, B6176, L6177, Key6178, C6179)


tmp20006 := Call2(__e, SymFn(symshen_4gc), B6176, tmp20005)


return TailCall2(__e, SymFn(symshen_4gc), B6176, tmp20006)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20007 := GValue(__e, symshen_4_dsigf_d)

tmp20008 := Call3(__e, SymFn(symshen_4assoc_1_6), symcompile, tmp19984, tmp20007)


tmp20009 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20008)

_ = tmp20009

tmp20010 := MakeNative1(func(__e *ControlFlow, V6185 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6181 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6182 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6183 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6184 Obj) Obj {
tmp20011 := Call1(__e, SymFn(symshen_4newpv), B6181)


let__18923 := tmp20011
_ = let__18923

tmp20012 := GCons(__e, symboolean, Nil)

tmp20013 := GCons(__e, sym_1_1_6, tmp20012)

tmp20014 := GCons(__e, let__18923, tmp20013)

tmp20015 := Call6(__e, SymFn(symis_b), V6185, tmp20014, B6181, L6182, Key6183, C6184)


return TailCall2(__e, SymFn(symshen_4gc), B6181, tmp20015)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20016 := GValue(__e, symshen_4_dsigf_d)

tmp20017 := Call3(__e, SymFn(symshen_4assoc_1_6), symcons_2, tmp20010, tmp20016)


tmp20018 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20017)

_ = tmp20018

tmp20019 := MakeNative1(func(__e *ControlFlow, V6190 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6186 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6187 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6188 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6189 Obj) Obj {
tmp20020 := GCons(__e, symsymbol, Nil)

tmp20021 := GCons(__e, symlist, tmp20020)

tmp20022 := GCons(__e, tmp20021, Nil)

tmp20023 := GCons(__e, sym_1_1_6, tmp20022)

return TailCall6(__e, SymFn(symis_b), V6190, tmp20023, B6186, L6187, Key6188, C6189)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20024 := GValue(__e, symshen_4_dsigf_d)

tmp20025 := Call3(__e, SymFn(symshen_4assoc_1_6), symdatatypes, tmp20019, tmp20024)


tmp20026 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20025)

_ = tmp20026

tmp20027 := MakeNative1(func(__e *ControlFlow, V6195 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6191 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6192 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6193 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6194 Obj) Obj {
tmp20028 := GCons(__e, symsymbol, Nil)

tmp20029 := GCons(__e, sym_1_1_6, tmp20028)

tmp20030 := GCons(__e, symsymbol, tmp20029)

return TailCall6(__e, SymFn(symis_b), V6195, tmp20030, B6191, L6192, Key6193, C6194)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20031 := GValue(__e, symshen_4_dsigf_d)

tmp20032 := Call3(__e, SymFn(symshen_4assoc_1_6), symdestroy, tmp20027, tmp20031)


tmp20033 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20032)

_ = tmp20033

tmp20034 := MakeNative1(func(__e *ControlFlow, V6202 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6198 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6199 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6200 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6201 Obj) Obj {
tmp20035 := Call1(__e, SymFn(symshen_4newpv), B6198)


let__18924 := tmp20035
_ = let__18924

tmp20036 := GCons(__e, let__18924, Nil)

tmp20037 := GCons(__e, symlist, tmp20036)

tmp20038 := GCons(__e, let__18924, Nil)

tmp20039 := GCons(__e, symlist, tmp20038)

tmp20040 := GCons(__e, let__18924, Nil)

tmp20041 := GCons(__e, symlist, tmp20040)

tmp20042 := GCons(__e, tmp20041, Nil)

tmp20043 := GCons(__e, sym_1_1_6, tmp20042)

tmp20044 := GCons(__e, tmp20039, tmp20043)

tmp20045 := GCons(__e, tmp20044, Nil)

tmp20046 := GCons(__e, sym_1_1_6, tmp20045)

tmp20047 := GCons(__e, tmp20037, tmp20046)

tmp20048 := Call6(__e, SymFn(symis_b), V6202, tmp20047, B6198, L6199, Key6200, C6201)


return TailCall2(__e, SymFn(symshen_4gc), B6198, tmp20048)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20049 := GValue(__e, symshen_4_dsigf_d)

tmp20050 := Call3(__e, SymFn(symshen_4assoc_1_6), symdifference, tmp20034, tmp20049)


tmp20051 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20050)

_ = tmp20051

tmp20052 := MakeNative1(func(__e *ControlFlow, V6207 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6203 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6204 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6205 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6206 Obj) Obj {
tmp20053 := Call1(__e, SymFn(symshen_4newpv), B6203)


let__18925 := tmp20053
_ = let__18925

tmp20054 := Call1(__e, SymFn(symshen_4newpv), B6203)


let__18926 := tmp20054
_ = let__18926

tmp20055 := GCons(__e, let__18926, Nil)

tmp20056 := GCons(__e, sym_1_1_6, tmp20055)

tmp20057 := GCons(__e, let__18926, tmp20056)

tmp20058 := GCons(__e, tmp20057, Nil)

tmp20059 := GCons(__e, sym_1_1_6, tmp20058)

tmp20060 := GCons(__e, let__18925, tmp20059)

tmp20061 := Call6(__e, SymFn(symis_b), V6207, tmp20060, B6203, L6204, Key6205, C6206)


tmp20062 := Call2(__e, SymFn(symshen_4gc), B6203, tmp20061)


return TailCall2(__e, SymFn(symshen_4gc), B6203, tmp20062)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20063 := GValue(__e, symshen_4_dsigf_d)

tmp20064 := Call3(__e, SymFn(symshen_4assoc_1_6), symdo, tmp20052, tmp20063)


tmp20065 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20064)

_ = tmp20065

tmp20066 := MakeNative1(func(__e *ControlFlow, V6214 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6210 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6211 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6212 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6213 Obj) Obj {
tmp20067 := Call1(__e, SymFn(symshen_4newpv), B6210)


let__18927 := tmp20067
_ = let__18927

tmp20068 := Call1(__e, SymFn(symshen_4newpv), B6210)


let__18928 := tmp20068
_ = let__18928

tmp20069 := GCons(__e, let__18927, Nil)

tmp20070 := GCons(__e, symlist, tmp20069)

tmp20071 := GCons(__e, let__18927, Nil)

tmp20072 := GCons(__e, symlist, tmp20071)

tmp20073 := GCons(__e, let__18928, Nil)

tmp20074 := GCons(__e, symlist, tmp20073)

tmp20075 := GCons(__e, tmp20074, Nil)

tmp20076 := GCons(__e, tmp20072, tmp20075)

tmp20077 := GCons(__e, symstr, tmp20076)

tmp20078 := GCons(__e, tmp20077, Nil)

tmp20079 := GCons(__e, sym_1_1_6, tmp20078)

tmp20080 := GCons(__e, tmp20070, tmp20079)

tmp20081 := Call6(__e, SymFn(symis_b), V6214, tmp20080, B6210, L6211, Key6212, C6213)


tmp20082 := Call2(__e, SymFn(symshen_4gc), B6210, tmp20081)


return TailCall2(__e, SymFn(symshen_4gc), B6210, tmp20082)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20083 := GValue(__e, symshen_4_dsigf_d)

tmp20084 := Call3(__e, SymFn(symshen_4assoc_1_6), sym_5e_6, tmp20066, tmp20083)


tmp20085 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20084)

_ = tmp20085

tmp20086 := MakeNative1(func(__e *ControlFlow, V6221 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6217 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6218 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6219 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6220 Obj) Obj {
tmp20087 := Call1(__e, SymFn(symshen_4newpv), B6217)


let__18929 := tmp20087
_ = let__18929

tmp20088 := Call1(__e, SymFn(symshen_4newpv), B6217)


let__18930 := tmp20088
_ = let__18930

tmp20089 := GCons(__e, let__18930, Nil)

tmp20090 := GCons(__e, symlist, tmp20089)

tmp20091 := GCons(__e, let__18929, Nil)

tmp20092 := GCons(__e, symlist, tmp20091)

tmp20093 := GCons(__e, let__18930, Nil)

tmp20094 := GCons(__e, symlist, tmp20093)

tmp20095 := GCons(__e, tmp20094, Nil)

tmp20096 := GCons(__e, tmp20092, tmp20095)

tmp20097 := GCons(__e, symstr, tmp20096)

tmp20098 := GCons(__e, tmp20097, Nil)

tmp20099 := GCons(__e, sym_1_1_6, tmp20098)

tmp20100 := GCons(__e, tmp20090, tmp20099)

tmp20101 := Call6(__e, SymFn(symis_b), V6221, tmp20100, B6217, L6218, Key6219, C6220)


tmp20102 := Call2(__e, SymFn(symshen_4gc), B6217, tmp20101)


return TailCall2(__e, SymFn(symshen_4gc), B6217, tmp20102)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20103 := GValue(__e, symshen_4_dsigf_d)

tmp20104 := Call3(__e, SymFn(symshen_4assoc_1_6), sym_5_b_6, tmp20086, tmp20103)


tmp20105 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20104)

_ = tmp20105

tmp20106 := MakeNative1(func(__e *ControlFlow, V6228 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6224 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6225 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6226 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6227 Obj) Obj {
tmp20107 := Call1(__e, SymFn(symshen_4newpv), B6224)


let__18931 := tmp20107
_ = let__18931

tmp20108 := Call1(__e, SymFn(symshen_4newpv), B6224)


let__18932 := tmp20108
_ = let__18932

tmp20109 := GCons(__e, let__18931, Nil)

tmp20110 := GCons(__e, symlist, tmp20109)

tmp20111 := GCons(__e, let__18931, Nil)

tmp20112 := GCons(__e, symlist, tmp20111)

tmp20113 := GCons(__e, let__18932, Nil)

tmp20114 := GCons(__e, symlist, tmp20113)

tmp20115 := GCons(__e, tmp20114, Nil)

tmp20116 := GCons(__e, tmp20112, tmp20115)

tmp20117 := GCons(__e, symstr, tmp20116)

tmp20118 := GCons(__e, tmp20117, Nil)

tmp20119 := GCons(__e, sym_1_1_6, tmp20118)

tmp20120 := GCons(__e, tmp20110, tmp20119)

tmp20121 := Call6(__e, SymFn(symis_b), V6228, tmp20120, B6224, L6225, Key6226, C6227)


tmp20122 := Call2(__e, SymFn(symshen_4gc), B6224, tmp20121)


return TailCall2(__e, SymFn(symshen_4gc), B6224, tmp20122)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20123 := GValue(__e, symshen_4_dsigf_d)

tmp20124 := Call3(__e, SymFn(symshen_4assoc_1_6), sym_5end_6, tmp20106, tmp20123)


tmp20125 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20124)

_ = tmp20125

tmp20126 := MakeNative1(func(__e *ControlFlow, V6233 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6229 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6230 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6231 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6232 Obj) Obj {
tmp20127 := Call1(__e, SymFn(symshen_4newpv), B6229)


let__18933 := tmp20127
_ = let__18933

tmp20128 := Call1(__e, SymFn(symshen_4newpv), B6229)


let__18934 := tmp20128
_ = let__18934

tmp20129 := GCons(__e, let__18933, Nil)

tmp20130 := GCons(__e, symlist, tmp20129)

tmp20131 := GCons(__e, let__18934, Nil)

tmp20132 := GCons(__e, tmp20130, tmp20131)

tmp20133 := GCons(__e, symstr, tmp20132)

tmp20134 := GCons(__e, symboolean, Nil)

tmp20135 := GCons(__e, sym_1_1_6, tmp20134)

tmp20136 := GCons(__e, tmp20133, tmp20135)

tmp20137 := Call6(__e, SymFn(symis_b), V6233, tmp20136, B6229, L6230, Key6231, C6232)


tmp20138 := Call2(__e, SymFn(symshen_4gc), B6229, tmp20137)


return TailCall2(__e, SymFn(symshen_4gc), B6229, tmp20138)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20139 := GValue(__e, symshen_4_dsigf_d)

tmp20140 := Call3(__e, SymFn(symshen_4assoc_1_6), symshen_4parse_1failure_2, tmp20126, tmp20139)


tmp20141 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20140)

_ = tmp20141

tmp20142 := MakeNative1(func(__e *ControlFlow, V6238 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6234 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6235 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6236 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6237 Obj) Obj {
tmp20143 := Call1(__e, SymFn(symshen_4newpv), B6234)


let__18935 := tmp20143
_ = let__18935

tmp20144 := Call1(__e, SymFn(symshen_4newpv), B6234)


let__18936 := tmp20144
_ = let__18936

tmp20145 := GCons(__e, let__18935, Nil)

tmp20146 := GCons(__e, symlist, tmp20145)

tmp20147 := GCons(__e, let__18936, Nil)

tmp20148 := GCons(__e, tmp20146, tmp20147)

tmp20149 := GCons(__e, symstr, tmp20148)

tmp20150 := GCons(__e, tmp20149, Nil)

tmp20151 := GCons(__e, sym_1_1_6, tmp20150)

tmp20152 := Call6(__e, SymFn(symis_b), V6238, tmp20151, B6234, L6235, Key6236, C6237)


tmp20153 := Call2(__e, SymFn(symshen_4gc), B6234, tmp20152)


return TailCall2(__e, SymFn(symshen_4gc), B6234, tmp20153)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20154 := GValue(__e, symshen_4_dsigf_d)

tmp20155 := Call3(__e, SymFn(symshen_4assoc_1_6), symshen_4parse_1failure, tmp20142, tmp20154)


tmp20156 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20155)

_ = tmp20156

tmp20157 := MakeNative1(func(__e *ControlFlow, V6243 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6239 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6240 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6241 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6242 Obj) Obj {
tmp20158 := Call1(__e, SymFn(symshen_4newpv), B6239)


let__18937 := tmp20158
_ = let__18937

tmp20159 := Call1(__e, SymFn(symshen_4newpv), B6239)


let__18938 := tmp20159
_ = let__18938

tmp20160 := GCons(__e, let__18937, Nil)

tmp20161 := GCons(__e, symlist, tmp20160)

tmp20162 := GCons(__e, let__18938, Nil)

tmp20163 := GCons(__e, tmp20161, tmp20162)

tmp20164 := GCons(__e, symstr, tmp20163)

tmp20165 := GCons(__e, let__18938, Nil)

tmp20166 := GCons(__e, sym_1_1_6, tmp20165)

tmp20167 := GCons(__e, tmp20164, tmp20166)

tmp20168 := Call6(__e, SymFn(symis_b), V6243, tmp20167, B6239, L6240, Key6241, C6242)


tmp20169 := Call2(__e, SymFn(symshen_4gc), B6239, tmp20168)


return TailCall2(__e, SymFn(symshen_4gc), B6239, tmp20169)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20170 := GValue(__e, symshen_4_dsigf_d)

tmp20171 := Call3(__e, SymFn(symshen_4assoc_1_6), symshen_4_5_1out, tmp20157, tmp20170)


tmp20172 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20171)

_ = tmp20172

tmp20173 := MakeNative1(func(__e *ControlFlow, V6249 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6245 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6246 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6247 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6248 Obj) Obj {
tmp20174 := Call1(__e, SymFn(symshen_4newpv), B6245)


let__18939 := tmp20174
_ = let__18939

tmp20175 := Call1(__e, SymFn(symshen_4newpv), B6245)


let__18940 := tmp20175
_ = let__18940

tmp20176 := GCons(__e, let__18940, Nil)

tmp20177 := GCons(__e, symlist, tmp20176)

tmp20178 := GCons(__e, let__18939, Nil)

tmp20179 := GCons(__e, tmp20177, tmp20178)

tmp20180 := GCons(__e, symstr, tmp20179)

tmp20181 := GCons(__e, let__18940, Nil)

tmp20182 := GCons(__e, symlist, tmp20181)

tmp20183 := GCons(__e, tmp20182, Nil)

tmp20184 := GCons(__e, sym_1_1_6, tmp20183)

tmp20185 := GCons(__e, tmp20180, tmp20184)

tmp20186 := Call6(__e, SymFn(symis_b), V6249, tmp20185, B6245, L6246, Key6247, C6248)


tmp20187 := Call2(__e, SymFn(symshen_4gc), B6245, tmp20186)


return TailCall2(__e, SymFn(symshen_4gc), B6245, tmp20187)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20188 := GValue(__e, symshen_4_dsigf_d)

tmp20189 := Call3(__e, SymFn(symshen_4assoc_1_6), symshen_4in_1_6, tmp20173, tmp20188)


tmp20190 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20189)

_ = tmp20190

tmp20191 := MakeNative1(func(__e *ControlFlow, V6255 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6251 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6252 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6253 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6254 Obj) Obj {
tmp20192 := Call1(__e, SymFn(symshen_4newpv), B6251)


let__18941 := tmp20192
_ = let__18941

tmp20193 := Call1(__e, SymFn(symshen_4newpv), B6251)


let__18942 := tmp20193
_ = let__18942

tmp20194 := GCons(__e, let__18941, Nil)

tmp20195 := GCons(__e, symlist, tmp20194)

tmp20196 := GCons(__e, let__18941, Nil)

tmp20197 := GCons(__e, symlist, tmp20196)

tmp20198 := GCons(__e, let__18942, Nil)

tmp20199 := GCons(__e, tmp20197, tmp20198)

tmp20200 := GCons(__e, symstr, tmp20199)

tmp20201 := GCons(__e, tmp20200, Nil)

tmp20202 := GCons(__e, sym_1_1_6, tmp20201)

tmp20203 := GCons(__e, let__18942, tmp20202)

tmp20204 := GCons(__e, tmp20203, Nil)

tmp20205 := GCons(__e, sym_1_1_6, tmp20204)

tmp20206 := GCons(__e, tmp20195, tmp20205)

tmp20207 := Call6(__e, SymFn(symis_b), V6255, tmp20206, B6251, L6252, Key6253, C6254)


tmp20208 := Call2(__e, SymFn(symshen_4gc), B6251, tmp20207)


return TailCall2(__e, SymFn(symshen_4gc), B6251, tmp20208)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20209 := GValue(__e, symshen_4_dsigf_d)

tmp20210 := Call3(__e, SymFn(symshen_4assoc_1_6), symshen_4comb, tmp20191, tmp20209)


tmp20211 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20210)

_ = tmp20211

tmp20212 := MakeNative1(func(__e *ControlFlow, V6260 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6256 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6257 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6258 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6259 Obj) Obj {
tmp20213 := Call1(__e, SymFn(symshen_4newpv), B6256)


let__18943 := tmp20213
_ = let__18943

tmp20214 := GCons(__e, let__18943, Nil)

tmp20215 := GCons(__e, symlist, tmp20214)

tmp20216 := GCons(__e, symboolean, Nil)

tmp20217 := GCons(__e, sym_1_1_6, tmp20216)

tmp20218 := GCons(__e, tmp20215, tmp20217)

tmp20219 := GCons(__e, tmp20218, Nil)

tmp20220 := GCons(__e, sym_1_1_6, tmp20219)

tmp20221 := GCons(__e, let__18943, tmp20220)

tmp20222 := Call6(__e, SymFn(symis_b), V6260, tmp20221, B6256, L6257, Key6258, C6259)


return TailCall2(__e, SymFn(symshen_4gc), B6256, tmp20222)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20223 := GValue(__e, symshen_4_dsigf_d)

tmp20224 := Call3(__e, SymFn(symshen_4assoc_1_6), symelement_2, tmp20212, tmp20223)


tmp20225 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20224)

_ = tmp20225

tmp20226 := MakeNative1(func(__e *ControlFlow, V6265 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6261 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6262 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6263 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6264 Obj) Obj {
tmp20227 := Call1(__e, SymFn(symshen_4newpv), B6261)


let__18944 := tmp20227
_ = let__18944

tmp20228 := GCons(__e, symboolean, Nil)

tmp20229 := GCons(__e, sym_1_1_6, tmp20228)

tmp20230 := GCons(__e, let__18944, tmp20229)

tmp20231 := Call6(__e, SymFn(symis_b), V6265, tmp20230, B6261, L6262, Key6263, C6264)


return TailCall2(__e, SymFn(symshen_4gc), B6261, tmp20231)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20232 := GValue(__e, symshen_4_dsigf_d)

tmp20233 := Call3(__e, SymFn(symshen_4assoc_1_6), symempty_2, tmp20226, tmp20232)


tmp20234 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20233)

_ = tmp20234

tmp20235 := MakeNative1(func(__e *ControlFlow, V6270 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6266 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6267 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6268 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6269 Obj) Obj {
tmp20236 := GCons(__e, symboolean, Nil)

tmp20237 := GCons(__e, sym_1_1_6, tmp20236)

tmp20238 := GCons(__e, symsymbol, tmp20237)

return TailCall6(__e, SymFn(symis_b), V6270, tmp20238, B6266, L6267, Key6268, C6269)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20239 := GValue(__e, symshen_4_dsigf_d)

tmp20240 := Call3(__e, SymFn(symshen_4assoc_1_6), symenable_1type_1theory, tmp20235, tmp20239)


tmp20241 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20240)

_ = tmp20241

tmp20242 := MakeNative1(func(__e *ControlFlow, V6275 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6271 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6272 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6273 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6274 Obj) Obj {
tmp20243 := GCons(__e, symsymbol, Nil)

tmp20244 := GCons(__e, symlist, tmp20243)

tmp20245 := GCons(__e, tmp20244, Nil)

tmp20246 := GCons(__e, sym_1_1_6, tmp20245)

tmp20247 := GCons(__e, symsymbol, tmp20246)

return TailCall6(__e, SymFn(symis_b), V6275, tmp20247, B6271, L6272, Key6273, C6274)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20248 := GValue(__e, symshen_4_dsigf_d)

tmp20249 := Call3(__e, SymFn(symshen_4assoc_1_6), symexternal, tmp20242, tmp20248)


tmp20250 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20249)

_ = tmp20250

tmp20251 := MakeNative1(func(__e *ControlFlow, V6280 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6276 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6277 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6278 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6279 Obj) Obj {
tmp20252 := GCons(__e, symstring, Nil)

tmp20253 := GCons(__e, sym_1_1_6, tmp20252)

tmp20254 := GCons(__e, symexception, tmp20253)

return TailCall6(__e, SymFn(symis_b), V6280, tmp20254, B6276, L6277, Key6278, C6279)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20255 := GValue(__e, symshen_4_dsigf_d)

tmp20256 := Call3(__e, SymFn(symshen_4assoc_1_6), symerror_1to_1string, tmp20251, tmp20255)


tmp20257 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20256)

_ = tmp20257

tmp20258 := MakeNative1(func(__e *ControlFlow, V6285 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6281 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6282 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6283 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6284 Obj) Obj {
tmp20259 := Call1(__e, SymFn(symshen_4newpv), B6281)


let__18945 := tmp20259
_ = let__18945

tmp20260 := GCons(__e, symstring, Nil)

tmp20261 := GCons(__e, symlist, tmp20260)

tmp20262 := GCons(__e, tmp20261, Nil)

tmp20263 := GCons(__e, sym_1_1_6, tmp20262)

tmp20264 := GCons(__e, let__18945, tmp20263)

tmp20265 := Call6(__e, SymFn(symis_b), V6285, tmp20264, B6281, L6282, Key6283, C6284)


return TailCall2(__e, SymFn(symshen_4gc), B6281, tmp20265)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20266 := GValue(__e, symshen_4_dsigf_d)

tmp20267 := Call3(__e, SymFn(symshen_4assoc_1_6), symexplode, tmp20258, tmp20266)


tmp20268 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20267)

_ = tmp20268

tmp20269 := MakeNative1(func(__e *ControlFlow, V6290 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6286 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6287 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6288 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6289 Obj) Obj {
tmp20270 := GCons(__e, symsymbol, Nil)

tmp20271 := GCons(__e, sym_1_1_6, tmp20270)

tmp20272 := GCons(__e, symsymbol, tmp20271)

return TailCall6(__e, SymFn(symis_b), V6290, tmp20272, B6286, L6287, Key6288, C6289)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20273 := GValue(__e, symshen_4_dsigf_d)

tmp20274 := Call3(__e, SymFn(symshen_4assoc_1_6), symfactorise, tmp20269, tmp20273)


tmp20275 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20274)

_ = tmp20275

tmp20276 := MakeNative1(func(__e *ControlFlow, V6295 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6291 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6292 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6293 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6294 Obj) Obj {
tmp20277 := GCons(__e, symboolean, Nil)

tmp20278 := GCons(__e, sym_1_1_6, tmp20277)

return TailCall6(__e, SymFn(symis_b), V6295, tmp20278, B6291, L6292, Key6293, C6294)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20279 := GValue(__e, symshen_4_dsigf_d)

tmp20280 := Call3(__e, SymFn(symshen_4assoc_1_6), symfactorise_2, tmp20276, tmp20279)


tmp20281 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20280)

_ = tmp20281

tmp20282 := MakeNative1(func(__e *ControlFlow, V6300 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6296 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6297 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6298 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6299 Obj) Obj {
tmp20283 := GCons(__e, symsymbol, Nil)

tmp20284 := GCons(__e, sym_1_1_6, tmp20283)

return TailCall6(__e, SymFn(symis_b), V6300, tmp20284, B6296, L6297, Key6298, C6299)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20285 := GValue(__e, symshen_4_dsigf_d)

tmp20286 := Call3(__e, SymFn(symshen_4assoc_1_6), symfail, tmp20282, tmp20285)


tmp20287 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20286)

_ = tmp20287

tmp20288 := MakeNative1(func(__e *ControlFlow, V6306 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6302 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6303 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6304 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6305 Obj) Obj {
tmp20289 := Call1(__e, SymFn(symshen_4newpv), B6302)


let__18946 := tmp20289
_ = let__18946

tmp20290 := GCons(__e, let__18946, Nil)

tmp20291 := GCons(__e, sym_1_1_6, tmp20290)

tmp20292 := GCons(__e, let__18946, tmp20291)

tmp20293 := GCons(__e, let__18946, Nil)

tmp20294 := GCons(__e, sym_1_1_6, tmp20293)

tmp20295 := GCons(__e, let__18946, tmp20294)

tmp20296 := GCons(__e, tmp20295, Nil)

tmp20297 := GCons(__e, sym_1_1_6, tmp20296)

tmp20298 := GCons(__e, tmp20292, tmp20297)

tmp20299 := Call6(__e, SymFn(symis_b), V6306, tmp20298, B6302, L6303, Key6304, C6305)


return TailCall2(__e, SymFn(symshen_4gc), B6302, tmp20299)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20300 := GValue(__e, symshen_4_dsigf_d)

tmp20301 := Call3(__e, SymFn(symshen_4assoc_1_6), symfix, tmp20288, tmp20300)


tmp20302 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20301)

_ = tmp20302

tmp20303 := MakeNative1(func(__e *ControlFlow, V6311 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6307 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6308 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6309 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6310 Obj) Obj {
tmp20304 := Call1(__e, SymFn(symshen_4newpv), B6307)


let__18947 := tmp20304
_ = let__18947

tmp20305 := GCons(__e, let__18947, Nil)

tmp20306 := GCons(__e, symlazy, tmp20305)

tmp20307 := GCons(__e, tmp20306, Nil)

tmp20308 := GCons(__e, sym_1_1_6, tmp20307)

tmp20309 := GCons(__e, let__18947, tmp20308)

tmp20310 := Call6(__e, SymFn(symis_b), V6311, tmp20309, B6307, L6308, Key6309, C6310)


return TailCall2(__e, SymFn(symshen_4gc), B6307, tmp20310)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20311 := GValue(__e, symshen_4_dsigf_d)

tmp20312 := Call3(__e, SymFn(symshen_4assoc_1_6), symfreeze, tmp20303, tmp20311)


tmp20313 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20312)

_ = tmp20313

tmp20314 := MakeNative1(func(__e *ControlFlow, V6316 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6312 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6313 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6314 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6315 Obj) Obj {
tmp20315 := Call1(__e, SymFn(symshen_4newpv), B6312)


let__18948 := tmp20315
_ = let__18948

tmp20316 := Call1(__e, SymFn(symshen_4newpv), B6312)


let__18949 := tmp20316
_ = let__18949

tmp20317 := GCons(__e, let__18948, Nil)

tmp20318 := GCons(__e, sym_d, tmp20317)

tmp20319 := GCons(__e, let__18949, tmp20318)

tmp20320 := GCons(__e, let__18949, Nil)

tmp20321 := GCons(__e, sym_1_1_6, tmp20320)

tmp20322 := GCons(__e, tmp20319, tmp20321)

tmp20323 := Call6(__e, SymFn(symis_b), V6316, tmp20322, B6312, L6313, Key6314, C6315)


tmp20324 := Call2(__e, SymFn(symshen_4gc), B6312, tmp20323)


return TailCall2(__e, SymFn(symshen_4gc), B6312, tmp20324)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20325 := GValue(__e, symshen_4_dsigf_d)

tmp20326 := Call3(__e, SymFn(symshen_4assoc_1_6), symfst, tmp20314, tmp20325)


tmp20327 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20326)

_ = tmp20327

tmp20328 := MakeNative1(func(__e *ControlFlow, V6321 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6317 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6318 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6319 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6320 Obj) Obj {
tmp20329 := GCons(__e, symsymbol, Nil)

tmp20330 := GCons(__e, sym_1_1_6, tmp20329)

tmp20331 := GCons(__e, symsymbol, tmp20330)

return TailCall6(__e, SymFn(symis_b), V6321, tmp20331, B6317, L6318, Key6319, C6320)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20332 := GValue(__e, symshen_4_dsigf_d)

tmp20333 := Call3(__e, SymFn(symshen_4assoc_1_6), symgensym, tmp20328, tmp20332)


tmp20334 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20333)

_ = tmp20334

tmp20335 := MakeNative1(func(__e *ControlFlow, V6327 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6323 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6324 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6325 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6326 Obj) Obj {
tmp20336 := Call1(__e, SymFn(symshen_4newpv), B6323)


let__18950 := tmp20336
_ = let__18950

tmp20337 := GCons(__e, let__18950, Nil)

tmp20338 := GCons(__e, symlist, tmp20337)

tmp20339 := GCons(__e, symboolean, Nil)

tmp20340 := GCons(__e, sym_1_1_6, tmp20339)

tmp20341 := GCons(__e, let__18950, tmp20340)

tmp20342 := GCons(__e, tmp20341, Nil)

tmp20343 := GCons(__e, sym_1_1_6, tmp20342)

tmp20344 := GCons(__e, tmp20338, tmp20343)

tmp20345 := Call6(__e, SymFn(symis_b), V6327, tmp20344, B6323, L6324, Key6325, C6326)


return TailCall2(__e, SymFn(symshen_4gc), B6323, tmp20345)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20346 := GValue(__e, symshen_4_dsigf_d)

tmp20347 := Call3(__e, SymFn(symshen_4assoc_1_6), symshen_4hds_a_2, tmp20335, tmp20346)


tmp20348 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20347)

_ = tmp20348

tmp20349 := MakeNative1(func(__e *ControlFlow, V6332 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6328 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6329 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6330 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6331 Obj) Obj {
tmp20350 := GCons(__e, symboolean, Nil)

tmp20351 := GCons(__e, sym_1_1_6, tmp20350)

tmp20352 := GCons(__e, symsymbol, tmp20351)

return TailCall6(__e, SymFn(symis_b), V6332, tmp20352, B6328, L6329, Key6330, C6331)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20353 := GValue(__e, symshen_4_dsigf_d)

tmp20354 := Call3(__e, SymFn(symshen_4assoc_1_6), symhush, tmp20349, tmp20353)


tmp20355 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20354)

_ = tmp20355

tmp20356 := MakeNative1(func(__e *ControlFlow, V6337 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6333 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6334 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6335 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6336 Obj) Obj {
tmp20357 := GCons(__e, symboolean, Nil)

tmp20358 := GCons(__e, sym_1_1_6, tmp20357)

return TailCall6(__e, SymFn(symis_b), V6337, tmp20358, B6333, L6334, Key6335, C6336)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20359 := GValue(__e, symshen_4_dsigf_d)

tmp20360 := Call3(__e, SymFn(symshen_4assoc_1_6), symhush_2, tmp20356, tmp20359)


tmp20361 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20360)

_ = tmp20361

tmp20362 := MakeNative1(func(__e *ControlFlow, V6343 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6339 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6340 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6341 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6342 Obj) Obj {
tmp20363 := Call1(__e, SymFn(symshen_4newpv), B6339)


let__18951 := tmp20363
_ = let__18951

tmp20364 := GCons(__e, let__18951, Nil)

tmp20365 := GCons(__e, symvector, tmp20364)

tmp20366 := GCons(__e, let__18951, Nil)

tmp20367 := GCons(__e, sym_1_1_6, tmp20366)

tmp20368 := GCons(__e, symnumber, tmp20367)

tmp20369 := GCons(__e, tmp20368, Nil)

tmp20370 := GCons(__e, sym_1_1_6, tmp20369)

tmp20371 := GCons(__e, tmp20365, tmp20370)

tmp20372 := Call6(__e, SymFn(symis_b), V6343, tmp20371, B6339, L6340, Key6341, C6342)


return TailCall2(__e, SymFn(symshen_4gc), B6339, tmp20372)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20373 := GValue(__e, symshen_4_dsigf_d)

tmp20374 := Call3(__e, SymFn(symshen_4assoc_1_6), sym_5_1vector, tmp20362, tmp20373)


tmp20375 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20374)

_ = tmp20375

tmp20376 := MakeNative1(func(__e *ControlFlow, V6349 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6345 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6346 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6347 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6348 Obj) Obj {
tmp20377 := Call1(__e, SymFn(symshen_4newpv), B6345)


let__18952 := tmp20377
_ = let__18952

tmp20378 := GCons(__e, let__18952, Nil)

tmp20379 := GCons(__e, symvector, tmp20378)

tmp20380 := GCons(__e, let__18952, Nil)

tmp20381 := GCons(__e, symvector, tmp20380)

tmp20382 := GCons(__e, tmp20381, Nil)

tmp20383 := GCons(__e, sym_1_1_6, tmp20382)

tmp20384 := GCons(__e, let__18952, tmp20383)

tmp20385 := GCons(__e, tmp20384, Nil)

tmp20386 := GCons(__e, sym_1_1_6, tmp20385)

tmp20387 := GCons(__e, symnumber, tmp20386)

tmp20388 := GCons(__e, tmp20387, Nil)

tmp20389 := GCons(__e, sym_1_1_6, tmp20388)

tmp20390 := GCons(__e, tmp20379, tmp20389)

tmp20391 := Call6(__e, SymFn(symis_b), V6349, tmp20390, B6345, L6346, Key6347, C6348)


return TailCall2(__e, SymFn(symshen_4gc), B6345, tmp20391)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20392 := GValue(__e, symshen_4_dsigf_d)

tmp20393 := Call3(__e, SymFn(symshen_4assoc_1_6), symvector_1_6, tmp20376, tmp20392)


tmp20394 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20393)

_ = tmp20394

tmp20395 := MakeNative1(func(__e *ControlFlow, V6354 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6350 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6351 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6352 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6353 Obj) Obj {
tmp20396 := Call1(__e, SymFn(symshen_4newpv), B6350)


let__18953 := tmp20396
_ = let__18953

tmp20397 := GCons(__e, let__18953, Nil)

tmp20398 := GCons(__e, symvector, tmp20397)

tmp20399 := GCons(__e, tmp20398, Nil)

tmp20400 := GCons(__e, sym_1_1_6, tmp20399)

tmp20401 := GCons(__e, symnumber, tmp20400)

tmp20402 := Call6(__e, SymFn(symis_b), V6354, tmp20401, B6350, L6351, Key6352, C6353)


return TailCall2(__e, SymFn(symshen_4gc), B6350, tmp20402)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20403 := GValue(__e, symshen_4_dsigf_d)

tmp20404 := Call3(__e, SymFn(symshen_4assoc_1_6), symvector, tmp20395, tmp20403)


tmp20405 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20404)

_ = tmp20405

tmp20406 := MakeNative1(func(__e *ControlFlow, V6359 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6355 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6356 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6357 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6358 Obj) Obj {
tmp20407 := GCons(__e, symnumber, Nil)

tmp20408 := GCons(__e, sym_1_1_6, tmp20407)

tmp20409 := GCons(__e, symsymbol, tmp20408)

return TailCall6(__e, SymFn(symis_b), V6359, tmp20409, B6355, L6356, Key6357, C6358)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20410 := GValue(__e, symshen_4_dsigf_d)

tmp20411 := Call3(__e, SymFn(symshen_4assoc_1_6), symget_1time, tmp20406, tmp20410)


tmp20412 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20411)

_ = tmp20412

tmp20413 := MakeNative1(func(__e *ControlFlow, V6364 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6360 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6361 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6362 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6363 Obj) Obj {
tmp20414 := Call1(__e, SymFn(symshen_4newpv), B6360)


let__18954 := tmp20414
_ = let__18954

tmp20415 := GCons(__e, symnumber, Nil)

tmp20416 := GCons(__e, sym_1_1_6, tmp20415)

tmp20417 := GCons(__e, symnumber, tmp20416)

tmp20418 := GCons(__e, tmp20417, Nil)

tmp20419 := GCons(__e, sym_1_1_6, tmp20418)

tmp20420 := GCons(__e, let__18954, tmp20419)

tmp20421 := Call6(__e, SymFn(symis_b), V6364, tmp20420, B6360, L6361, Key6362, C6363)


return TailCall2(__e, SymFn(symshen_4gc), B6360, tmp20421)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20422 := GValue(__e, symshen_4_dsigf_d)

tmp20423 := Call3(__e, SymFn(symshen_4assoc_1_6), symhash, tmp20413, tmp20422)


tmp20424 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20423)

_ = tmp20424

tmp20425 := MakeNative1(func(__e *ControlFlow, V6369 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6365 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6366 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6367 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6368 Obj) Obj {
tmp20426 := Call1(__e, SymFn(symshen_4newpv), B6365)


let__18955 := tmp20426
_ = let__18955

tmp20427 := GCons(__e, let__18955, Nil)

tmp20428 := GCons(__e, symlist, tmp20427)

tmp20429 := GCons(__e, let__18955, Nil)

tmp20430 := GCons(__e, sym_1_1_6, tmp20429)

tmp20431 := GCons(__e, tmp20428, tmp20430)

tmp20432 := Call6(__e, SymFn(symis_b), V6369, tmp20431, B6365, L6366, Key6367, C6368)


return TailCall2(__e, SymFn(symshen_4gc), B6365, tmp20432)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20433 := GValue(__e, symshen_4_dsigf_d)

tmp20434 := Call3(__e, SymFn(symshen_4assoc_1_6), symhead, tmp20425, tmp20433)


tmp20435 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20434)

_ = tmp20435

tmp20436 := MakeNative1(func(__e *ControlFlow, V6374 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6370 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6371 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6372 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6373 Obj) Obj {
tmp20437 := Call1(__e, SymFn(symshen_4newpv), B6370)


let__18956 := tmp20437
_ = let__18956

tmp20438 := GCons(__e, let__18956, Nil)

tmp20439 := GCons(__e, symvector, tmp20438)

tmp20440 := GCons(__e, let__18956, Nil)

tmp20441 := GCons(__e, sym_1_1_6, tmp20440)

tmp20442 := GCons(__e, tmp20439, tmp20441)

tmp20443 := Call6(__e, SymFn(symis_b), V6374, tmp20442, B6370, L6371, Key6372, C6373)


return TailCall2(__e, SymFn(symshen_4gc), B6370, tmp20443)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20444 := GValue(__e, symshen_4_dsigf_d)

tmp20445 := Call3(__e, SymFn(symshen_4assoc_1_6), symhdv, tmp20436, tmp20444)


tmp20446 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20445)

_ = tmp20446

tmp20447 := MakeNative1(func(__e *ControlFlow, V6379 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6375 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6376 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6377 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6378 Obj) Obj {
tmp20448 := GCons(__e, symstring, Nil)

tmp20449 := GCons(__e, sym_1_1_6, tmp20448)

tmp20450 := GCons(__e, symstring, tmp20449)

return TailCall6(__e, SymFn(symis_b), V6379, tmp20450, B6375, L6376, Key6377, C6378)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20451 := GValue(__e, symshen_4_dsigf_d)

tmp20452 := Call3(__e, SymFn(symshen_4assoc_1_6), symhdstr, tmp20447, tmp20451)


tmp20453 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20452)

_ = tmp20453

tmp20454 := MakeNative1(func(__e *ControlFlow, V6384 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6380 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6381 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6382 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6383 Obj) Obj {
tmp20455 := Call1(__e, SymFn(symshen_4newpv), B6380)


let__18957 := tmp20455
_ = let__18957

tmp20456 := GCons(__e, let__18957, Nil)

tmp20457 := GCons(__e, sym_1_1_6, tmp20456)

tmp20458 := GCons(__e, let__18957, tmp20457)

tmp20459 := GCons(__e, tmp20458, Nil)

tmp20460 := GCons(__e, sym_1_1_6, tmp20459)

tmp20461 := GCons(__e, let__18957, tmp20460)

tmp20462 := GCons(__e, tmp20461, Nil)

tmp20463 := GCons(__e, sym_1_1_6, tmp20462)

tmp20464 := GCons(__e, symboolean, tmp20463)

tmp20465 := Call6(__e, SymFn(symis_b), V6384, tmp20464, B6380, L6381, Key6382, C6383)


return TailCall2(__e, SymFn(symshen_4gc), B6380, tmp20465)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20466 := GValue(__e, symshen_4_dsigf_d)

tmp20467 := Call3(__e, SymFn(symshen_4assoc_1_6), symif, tmp20454, tmp20466)


tmp20468 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20467)

_ = tmp20468

tmp20469 := MakeNative1(func(__e *ControlFlow, V6389 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6385 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6386 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6387 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6388 Obj) Obj {
tmp20470 := GCons(__e, symsymbol, Nil)

tmp20471 := GCons(__e, sym_1_1_6, tmp20470)

tmp20472 := GCons(__e, symsymbol, tmp20471)

return TailCall6(__e, SymFn(symis_b), V6389, tmp20472, B6385, L6386, Key6387, C6388)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20473 := GValue(__e, symshen_4_dsigf_d)

tmp20474 := Call3(__e, SymFn(symshen_4assoc_1_6), symin_1package, tmp20469, tmp20473)


tmp20475 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20474)

_ = tmp20475

tmp20476 := MakeNative1(func(__e *ControlFlow, V6394 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6390 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6391 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6392 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6393 Obj) Obj {
tmp20477 := GCons(__e, symstring, Nil)

tmp20478 := GCons(__e, sym_1_1_6, tmp20477)

return TailCall6(__e, SymFn(symis_b), V6394, tmp20478, B6390, L6391, Key6392, C6393)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20479 := GValue(__e, symshen_4_dsigf_d)

tmp20480 := Call3(__e, SymFn(symshen_4assoc_1_6), symit, tmp20476, tmp20479)


tmp20481 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20480)

_ = tmp20481

tmp20482 := MakeNative1(func(__e *ControlFlow, V6399 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6395 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6396 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6397 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6398 Obj) Obj {
tmp20483 := GCons(__e, symstring, Nil)

tmp20484 := GCons(__e, sym_1_1_6, tmp20483)

return TailCall6(__e, SymFn(symis_b), V6399, tmp20484, B6395, L6396, Key6397, C6398)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20485 := GValue(__e, symshen_4_dsigf_d)

tmp20486 := Call3(__e, SymFn(symshen_4assoc_1_6), symimplementation, tmp20482, tmp20485)


tmp20487 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20486)

_ = tmp20487

tmp20488 := MakeNative1(func(__e *ControlFlow, V6405 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6401 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6402 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6403 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6404 Obj) Obj {
tmp20489 := GCons(__e, symsymbol, Nil)

tmp20490 := GCons(__e, symlist, tmp20489)

tmp20491 := GCons(__e, symsymbol, Nil)

tmp20492 := GCons(__e, symlist, tmp20491)

tmp20493 := GCons(__e, tmp20492, Nil)

tmp20494 := GCons(__e, sym_1_1_6, tmp20493)

tmp20495 := GCons(__e, tmp20490, tmp20494)

return TailCall6(__e, SymFn(symis_b), V6405, tmp20495, B6401, L6402, Key6403, C6404)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20496 := GValue(__e, symshen_4_dsigf_d)

tmp20497 := Call3(__e, SymFn(symshen_4assoc_1_6), syminclude, tmp20488, tmp20496)


tmp20498 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20497)

_ = tmp20498

tmp20499 := MakeNative1(func(__e *ControlFlow, V6411 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6407 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6408 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6409 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6410 Obj) Obj {
tmp20500 := GCons(__e, symsymbol, Nil)

tmp20501 := GCons(__e, symlist, tmp20500)

tmp20502 := GCons(__e, symsymbol, Nil)

tmp20503 := GCons(__e, symlist, tmp20502)

tmp20504 := GCons(__e, tmp20503, Nil)

tmp20505 := GCons(__e, sym_1_1_6, tmp20504)

tmp20506 := GCons(__e, tmp20501, tmp20505)

return TailCall6(__e, SymFn(symis_b), V6411, tmp20506, B6407, L6408, Key6409, C6410)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20507 := GValue(__e, symshen_4_dsigf_d)

tmp20508 := Call3(__e, SymFn(symshen_4assoc_1_6), syminclude_1all_1but, tmp20499, tmp20507)


tmp20509 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20508)

_ = tmp20509

tmp20510 := MakeNative1(func(__e *ControlFlow, V6416 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6412 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6413 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6414 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6415 Obj) Obj {
tmp20511 := GCons(__e, symsymbol, Nil)

tmp20512 := GCons(__e, symlist, tmp20511)

tmp20513 := GCons(__e, tmp20512, Nil)

tmp20514 := GCons(__e, sym_1_1_6, tmp20513)

return TailCall6(__e, SymFn(symis_b), V6416, tmp20514, B6412, L6413, Key6414, C6415)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20515 := GValue(__e, symshen_4_dsigf_d)

tmp20516 := Call3(__e, SymFn(symshen_4assoc_1_6), symincluded, tmp20510, tmp20515)


tmp20517 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20516)

_ = tmp20517

tmp20518 := MakeNative1(func(__e *ControlFlow, V6421 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6417 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6418 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6419 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6420 Obj) Obj {
tmp20519 := GCons(__e, symnumber, Nil)

tmp20520 := GCons(__e, sym_1_1_6, tmp20519)

return TailCall6(__e, SymFn(symis_b), V6421, tmp20520, B6417, L6418, Key6419, C6420)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20521 := GValue(__e, symshen_4_dsigf_d)

tmp20522 := Call3(__e, SymFn(symshen_4assoc_1_6), syminferences, tmp20518, tmp20521)


tmp20523 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20522)

_ = tmp20523

tmp20524 := MakeNative1(func(__e *ControlFlow, V6426 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6422 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6423 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6424 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6425 Obj) Obj {
tmp20525 := Call1(__e, SymFn(symshen_4newpv), B6422)


let__18958 := tmp20525
_ = let__18958

tmp20526 := GCons(__e, symstring, Nil)

tmp20527 := GCons(__e, sym_1_1_6, tmp20526)

tmp20528 := GCons(__e, symstring, tmp20527)

tmp20529 := GCons(__e, tmp20528, Nil)

tmp20530 := GCons(__e, sym_1_1_6, tmp20529)

tmp20531 := GCons(__e, let__18958, tmp20530)

tmp20532 := Call6(__e, SymFn(symis_b), V6426, tmp20531, B6422, L6423, Key6424, C6425)


return TailCall2(__e, SymFn(symshen_4gc), B6422, tmp20532)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20533 := GValue(__e, symshen_4_dsigf_d)

tmp20534 := Call3(__e, SymFn(symshen_4assoc_1_6), symshen_4insert, tmp20524, tmp20533)


tmp20535 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20534)

_ = tmp20535

tmp20536 := MakeNative1(func(__e *ControlFlow, V6431 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6427 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6428 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6429 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6430 Obj) Obj {
tmp20537 := Call1(__e, SymFn(symshen_4newpv), B6427)


let__18959 := tmp20537
_ = let__18959

tmp20538 := GCons(__e, symboolean, Nil)

tmp20539 := GCons(__e, sym_1_1_6, tmp20538)

tmp20540 := GCons(__e, let__18959, tmp20539)

tmp20541 := Call6(__e, SymFn(symis_b), V6431, tmp20540, B6427, L6428, Key6429, C6430)


return TailCall2(__e, SymFn(symshen_4gc), B6427, tmp20541)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20542 := GValue(__e, symshen_4_dsigf_d)

tmp20543 := Call3(__e, SymFn(symshen_4assoc_1_6), syminteger_2, tmp20536, tmp20542)


tmp20544 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20543)

_ = tmp20544

tmp20545 := MakeNative1(func(__e *ControlFlow, V6436 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6432 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6433 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6434 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6435 Obj) Obj {
tmp20546 := GCons(__e, symsymbol, Nil)

tmp20547 := GCons(__e, symlist, tmp20546)

tmp20548 := GCons(__e, tmp20547, Nil)

tmp20549 := GCons(__e, sym_1_1_6, tmp20548)

tmp20550 := GCons(__e, symsymbol, tmp20549)

return TailCall6(__e, SymFn(symis_b), V6436, tmp20550, B6432, L6433, Key6434, C6435)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20551 := GValue(__e, symshen_4_dsigf_d)

tmp20552 := Call3(__e, SymFn(symshen_4assoc_1_6), syminternal, tmp20545, tmp20551)


tmp20553 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20552)

_ = tmp20553

tmp20554 := MakeNative1(func(__e *ControlFlow, V6443 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6439 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6440 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6441 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6442 Obj) Obj {
tmp20555 := Call1(__e, SymFn(symshen_4newpv), B6439)


let__18960 := tmp20555
_ = let__18960

tmp20556 := GCons(__e, let__18960, Nil)

tmp20557 := GCons(__e, symlist, tmp20556)

tmp20558 := GCons(__e, let__18960, Nil)

tmp20559 := GCons(__e, symlist, tmp20558)

tmp20560 := GCons(__e, let__18960, Nil)

tmp20561 := GCons(__e, symlist, tmp20560)

tmp20562 := GCons(__e, tmp20561, Nil)

tmp20563 := GCons(__e, sym_1_1_6, tmp20562)

tmp20564 := GCons(__e, tmp20559, tmp20563)

tmp20565 := GCons(__e, tmp20564, Nil)

tmp20566 := GCons(__e, sym_1_1_6, tmp20565)

tmp20567 := GCons(__e, tmp20557, tmp20566)

tmp20568 := Call6(__e, SymFn(symis_b), V6443, tmp20567, B6439, L6440, Key6441, C6442)


return TailCall2(__e, SymFn(symshen_4gc), B6439, tmp20568)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20569 := GValue(__e, symshen_4_dsigf_d)

tmp20570 := Call3(__e, SymFn(symshen_4assoc_1_6), symintersection, tmp20554, tmp20569)


tmp20571 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20570)

_ = tmp20571

tmp20572 := MakeNative1(func(__e *ControlFlow, V6448 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6444 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6445 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6446 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6447 Obj) Obj {
tmp20573 := GCons(__e, symstring, Nil)

tmp20574 := GCons(__e, sym_1_1_6, tmp20573)

return TailCall6(__e, SymFn(symis_b), V6448, tmp20574, B6444, L6445, Key6446, C6447)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20575 := GValue(__e, symshen_4_dsigf_d)

tmp20576 := Call3(__e, SymFn(symshen_4assoc_1_6), symlanguage, tmp20572, tmp20575)


tmp20577 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20576)

_ = tmp20577

tmp20578 := MakeNative1(func(__e *ControlFlow, V6453 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6449 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6450 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6451 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6452 Obj) Obj {
tmp20579 := Call1(__e, SymFn(symshen_4newpv), B6449)


let__18961 := tmp20579
_ = let__18961

tmp20580 := GCons(__e, let__18961, Nil)

tmp20581 := GCons(__e, symlist, tmp20580)

tmp20582 := GCons(__e, symnumber, Nil)

tmp20583 := GCons(__e, sym_1_1_6, tmp20582)

tmp20584 := GCons(__e, tmp20581, tmp20583)

tmp20585 := Call6(__e, SymFn(symis_b), V6453, tmp20584, B6449, L6450, Key6451, C6452)


return TailCall2(__e, SymFn(symshen_4gc), B6449, tmp20585)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20586 := GValue(__e, symshen_4_dsigf_d)

tmp20587 := Call3(__e, SymFn(symshen_4assoc_1_6), symlength, tmp20578, tmp20586)


tmp20588 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20587)

_ = tmp20588

tmp20589 := MakeNative1(func(__e *ControlFlow, V6458 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6454 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6455 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6456 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6457 Obj) Obj {
tmp20590 := Call1(__e, SymFn(symshen_4newpv), B6454)


let__18962 := tmp20590
_ = let__18962

tmp20591 := GCons(__e, let__18962, Nil)

tmp20592 := GCons(__e, symvector, tmp20591)

tmp20593 := GCons(__e, symnumber, Nil)

tmp20594 := GCons(__e, sym_1_1_6, tmp20593)

tmp20595 := GCons(__e, tmp20592, tmp20594)

tmp20596 := Call6(__e, SymFn(symis_b), V6458, tmp20595, B6454, L6455, Key6456, C6457)


return TailCall2(__e, SymFn(symshen_4gc), B6454, tmp20596)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20597 := GValue(__e, symshen_4_dsigf_d)

tmp20598 := Call3(__e, SymFn(symshen_4assoc_1_6), symlimit, tmp20589, tmp20597)


tmp20599 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20598)

_ = tmp20599

tmp20600 := MakeNative1(func(__e *ControlFlow, V6464 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6460 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6461 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6462 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6463 Obj) Obj {
tmp20601 := GCons(__e, symin, Nil)

tmp20602 := GCons(__e, symstream, tmp20601)

tmp20603 := GCons(__e, symunit, Nil)

tmp20604 := GCons(__e, symlist, tmp20603)

tmp20605 := GCons(__e, tmp20604, Nil)

tmp20606 := GCons(__e, sym_1_1_6, tmp20605)

tmp20607 := GCons(__e, tmp20602, tmp20606)

return TailCall6(__e, SymFn(symis_b), V6464, tmp20607, B6460, L6461, Key6462, C6463)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20608 := GValue(__e, symshen_4_dsigf_d)

tmp20609 := Call3(__e, SymFn(symshen_4assoc_1_6), symlineread, tmp20600, tmp20608)


tmp20610 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20609)

_ = tmp20610

tmp20611 := MakeNative1(func(__e *ControlFlow, V6469 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6465 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6466 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6467 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6468 Obj) Obj {
tmp20612 := GCons(__e, symsymbol, Nil)

tmp20613 := GCons(__e, sym_1_1_6, tmp20612)

tmp20614 := GCons(__e, symstring, tmp20613)

return TailCall6(__e, SymFn(symis_b), V6469, tmp20614, B6465, L6466, Key6467, C6468)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20615 := GValue(__e, symshen_4_dsigf_d)

tmp20616 := Call3(__e, SymFn(symshen_4assoc_1_6), symload, tmp20611, tmp20615)


tmp20617 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20616)

_ = tmp20617

tmp20618 := MakeNative1(func(__e *ControlFlow, V6476 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6472 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6473 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6474 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6475 Obj) Obj {
tmp20619 := Call1(__e, SymFn(symshen_4newpv), B6472)


let__18963 := tmp20619
_ = let__18963

tmp20620 := Call1(__e, SymFn(symshen_4newpv), B6472)


let__18964 := tmp20620
_ = let__18964

tmp20621 := GCons(__e, let__18964, Nil)

tmp20622 := GCons(__e, sym_1_1_6, tmp20621)

tmp20623 := GCons(__e, let__18963, tmp20622)

tmp20624 := GCons(__e, let__18963, Nil)

tmp20625 := GCons(__e, symlist, tmp20624)

tmp20626 := GCons(__e, let__18964, Nil)

tmp20627 := GCons(__e, symlist, tmp20626)

tmp20628 := GCons(__e, tmp20627, Nil)

tmp20629 := GCons(__e, sym_1_1_6, tmp20628)

tmp20630 := GCons(__e, tmp20625, tmp20629)

tmp20631 := GCons(__e, tmp20630, Nil)

tmp20632 := GCons(__e, sym_1_1_6, tmp20631)

tmp20633 := GCons(__e, tmp20623, tmp20632)

tmp20634 := Call6(__e, SymFn(symis_b), V6476, tmp20633, B6472, L6473, Key6474, C6475)


tmp20635 := Call2(__e, SymFn(symshen_4gc), B6472, tmp20634)


return TailCall2(__e, SymFn(symshen_4gc), B6472, tmp20635)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20636 := GValue(__e, symshen_4_dsigf_d)

tmp20637 := Call3(__e, SymFn(symshen_4assoc_1_6), symmap, tmp20618, tmp20636)


tmp20638 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20637)

_ = tmp20638

tmp20639 := MakeNative1(func(__e *ControlFlow, V6483 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6479 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6480 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6481 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6482 Obj) Obj {
tmp20640 := Call1(__e, SymFn(symshen_4newpv), B6479)


let__18965 := tmp20640
_ = let__18965

tmp20641 := Call1(__e, SymFn(symshen_4newpv), B6479)


let__18966 := tmp20641
_ = let__18966

tmp20642 := GCons(__e, let__18966, Nil)

tmp20643 := GCons(__e, symlist, tmp20642)

tmp20644 := GCons(__e, tmp20643, Nil)

tmp20645 := GCons(__e, sym_1_1_6, tmp20644)

tmp20646 := GCons(__e, let__18965, tmp20645)

tmp20647 := GCons(__e, let__18965, Nil)

tmp20648 := GCons(__e, symlist, tmp20647)

tmp20649 := GCons(__e, let__18966, Nil)

tmp20650 := GCons(__e, symlist, tmp20649)

tmp20651 := GCons(__e, tmp20650, Nil)

tmp20652 := GCons(__e, sym_1_1_6, tmp20651)

tmp20653 := GCons(__e, tmp20648, tmp20652)

tmp20654 := GCons(__e, tmp20653, Nil)

tmp20655 := GCons(__e, sym_1_1_6, tmp20654)

tmp20656 := GCons(__e, tmp20646, tmp20655)

tmp20657 := Call6(__e, SymFn(symis_b), V6483, tmp20656, B6479, L6480, Key6481, C6482)


tmp20658 := Call2(__e, SymFn(symshen_4gc), B6479, tmp20657)


return TailCall2(__e, SymFn(symshen_4gc), B6479, tmp20658)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20659 := GValue(__e, symshen_4_dsigf_d)

tmp20660 := Call3(__e, SymFn(symshen_4assoc_1_6), symmapcan, tmp20639, tmp20659)


tmp20661 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20660)

_ = tmp20661

tmp20662 := MakeNative1(func(__e *ControlFlow, V6488 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6484 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6485 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6486 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6487 Obj) Obj {
tmp20663 := GCons(__e, symnumber, Nil)

tmp20664 := GCons(__e, sym_1_1_6, tmp20663)

tmp20665 := GCons(__e, symnumber, tmp20664)

return TailCall6(__e, SymFn(symis_b), V6488, tmp20665, B6484, L6485, Key6486, C6487)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20666 := GValue(__e, symshen_4_dsigf_d)

tmp20667 := Call3(__e, SymFn(symshen_4assoc_1_6), symmaxinferences, tmp20662, tmp20666)


tmp20668 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20667)

_ = tmp20668

tmp20669 := MakeNative1(func(__e *ControlFlow, V6493 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6489 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6490 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6491 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6492 Obj) Obj {
tmp20670 := GCons(__e, symstring, Nil)

tmp20671 := GCons(__e, sym_1_1_6, tmp20670)

tmp20672 := GCons(__e, symnumber, tmp20671)

return TailCall6(__e, SymFn(symis_b), V6493, tmp20672, B6489, L6490, Key6491, C6492)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20673 := GValue(__e, symshen_4_dsigf_d)

tmp20674 := Call3(__e, SymFn(symshen_4assoc_1_6), symn_1_6string, tmp20669, tmp20673)


tmp20675 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20674)

_ = tmp20675

tmp20676 := MakeNative1(func(__e *ControlFlow, V6498 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6494 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6495 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6496 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6497 Obj) Obj {
tmp20677 := GCons(__e, symnumber, Nil)

tmp20678 := GCons(__e, sym_1_1_6, tmp20677)

tmp20679 := GCons(__e, symnumber, tmp20678)

return TailCall6(__e, SymFn(symis_b), V6498, tmp20679, B6494, L6495, Key6496, C6497)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20680 := GValue(__e, symshen_4_dsigf_d)

tmp20681 := Call3(__e, SymFn(symshen_4assoc_1_6), symnl, tmp20676, tmp20680)


tmp20682 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20681)

_ = tmp20682

tmp20683 := MakeNative1(func(__e *ControlFlow, V6503 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6499 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6500 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6501 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6502 Obj) Obj {
tmp20684 := GCons(__e, symboolean, Nil)

tmp20685 := GCons(__e, sym_1_1_6, tmp20684)

tmp20686 := GCons(__e, symboolean, tmp20685)

return TailCall6(__e, SymFn(symis_b), V6503, tmp20686, B6499, L6500, Key6501, C6502)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20687 := GValue(__e, symshen_4_dsigf_d)

tmp20688 := Call3(__e, SymFn(symshen_4assoc_1_6), symnot, tmp20683, tmp20687)


tmp20689 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20688)

_ = tmp20689

tmp20690 := MakeNative1(func(__e *ControlFlow, V6508 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6504 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6505 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6506 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6507 Obj) Obj {
tmp20691 := Call1(__e, SymFn(symshen_4newpv), B6504)


let__18967 := tmp20691
_ = let__18967

tmp20692 := GCons(__e, let__18967, Nil)

tmp20693 := GCons(__e, symlist, tmp20692)

tmp20694 := GCons(__e, let__18967, Nil)

tmp20695 := GCons(__e, sym_1_1_6, tmp20694)

tmp20696 := GCons(__e, tmp20693, tmp20695)

tmp20697 := GCons(__e, tmp20696, Nil)

tmp20698 := GCons(__e, sym_1_1_6, tmp20697)

tmp20699 := GCons(__e, symnumber, tmp20698)

tmp20700 := Call6(__e, SymFn(symis_b), V6508, tmp20699, B6504, L6505, Key6506, C6507)


return TailCall2(__e, SymFn(symshen_4gc), B6504, tmp20700)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20701 := GValue(__e, symshen_4_dsigf_d)

tmp20702 := Call3(__e, SymFn(symshen_4assoc_1_6), symnth, tmp20690, tmp20701)


tmp20703 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20702)

_ = tmp20703

tmp20704 := MakeNative1(func(__e *ControlFlow, V6513 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6509 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6510 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6511 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6512 Obj) Obj {
tmp20705 := Call1(__e, SymFn(symshen_4newpv), B6509)


let__18968 := tmp20705
_ = let__18968

tmp20706 := GCons(__e, symboolean, Nil)

tmp20707 := GCons(__e, sym_1_1_6, tmp20706)

tmp20708 := GCons(__e, let__18968, tmp20707)

tmp20709 := Call6(__e, SymFn(symis_b), V6513, tmp20708, B6509, L6510, Key6511, C6512)


return TailCall2(__e, SymFn(symshen_4gc), B6509, tmp20709)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20710 := GValue(__e, symshen_4_dsigf_d)

tmp20711 := Call3(__e, SymFn(symshen_4assoc_1_6), symnumber_2, tmp20704, tmp20710)


tmp20712 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20711)

_ = tmp20712

tmp20713 := MakeNative1(func(__e *ControlFlow, V6518 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6514 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6515 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6516 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6517 Obj) Obj {
tmp20714 := Call1(__e, SymFn(symshen_4newpv), B6514)


let__18969 := tmp20714
_ = let__18969

tmp20715 := Call1(__e, SymFn(symshen_4newpv), B6514)


let__18970 := tmp20715
_ = let__18970

tmp20716 := GCons(__e, symnumber, Nil)

tmp20717 := GCons(__e, sym_1_1_6, tmp20716)

tmp20718 := GCons(__e, let__18970, tmp20717)

tmp20719 := GCons(__e, tmp20718, Nil)

tmp20720 := GCons(__e, sym_1_1_6, tmp20719)

tmp20721 := GCons(__e, let__18969, tmp20720)

tmp20722 := Call6(__e, SymFn(symis_b), V6518, tmp20721, B6514, L6515, Key6516, C6517)


tmp20723 := Call2(__e, SymFn(symshen_4gc), B6514, tmp20722)


return TailCall2(__e, SymFn(symshen_4gc), B6514, tmp20723)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20724 := GValue(__e, symshen_4_dsigf_d)

tmp20725 := Call3(__e, SymFn(symshen_4assoc_1_6), symoccurrences, tmp20713, tmp20724)


tmp20726 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20725)

_ = tmp20726

tmp20727 := MakeNative1(func(__e *ControlFlow, V6523 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6519 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6520 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6521 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6522 Obj) Obj {
tmp20728 := GCons(__e, symboolean, Nil)

tmp20729 := GCons(__e, sym_1_1_6, tmp20728)

tmp20730 := GCons(__e, symsymbol, tmp20729)

return TailCall6(__e, SymFn(symis_b), V6523, tmp20730, B6519, L6520, Key6521, C6522)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20731 := GValue(__e, symshen_4_dsigf_d)

tmp20732 := Call3(__e, SymFn(symshen_4assoc_1_6), symoccurs_1check, tmp20727, tmp20731)


tmp20733 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20732)

_ = tmp20733

tmp20734 := MakeNative1(func(__e *ControlFlow, V6528 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6524 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6525 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6526 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6527 Obj) Obj {
tmp20735 := GCons(__e, symboolean, Nil)

tmp20736 := GCons(__e, sym_1_1_6, tmp20735)

return TailCall6(__e, SymFn(symis_b), V6528, tmp20736, B6524, L6525, Key6526, C6527)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20737 := GValue(__e, symshen_4_dsigf_d)

tmp20738 := Call3(__e, SymFn(symshen_4assoc_1_6), symoccurs_2, tmp20734, tmp20737)


tmp20739 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20738)

_ = tmp20739

tmp20740 := MakeNative1(func(__e *ControlFlow, V6533 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6529 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6530 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6531 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6532 Obj) Obj {
tmp20741 := GCons(__e, symboolean, Nil)

tmp20742 := GCons(__e, sym_1_1_6, tmp20741)

tmp20743 := GCons(__e, symsymbol, tmp20742)

return TailCall6(__e, SymFn(symis_b), V6533, tmp20743, B6529, L6530, Key6531, C6532)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20744 := GValue(__e, symshen_4_dsigf_d)

tmp20745 := Call3(__e, SymFn(symshen_4assoc_1_6), symoptimise, tmp20740, tmp20744)


tmp20746 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20745)

_ = tmp20746

tmp20747 := MakeNative1(func(__e *ControlFlow, V6538 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6534 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6535 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6536 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6537 Obj) Obj {
tmp20748 := GCons(__e, symboolean, Nil)

tmp20749 := GCons(__e, sym_1_1_6, tmp20748)

return TailCall6(__e, SymFn(symis_b), V6538, tmp20749, B6534, L6535, Key6536, C6537)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20750 := GValue(__e, symshen_4_dsigf_d)

tmp20751 := Call3(__e, SymFn(symshen_4assoc_1_6), symoptimise_2, tmp20747, tmp20750)


tmp20752 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20751)

_ = tmp20752

tmp20753 := MakeNative1(func(__e *ControlFlow, V6543 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6539 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6540 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6541 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6542 Obj) Obj {
tmp20754 := GCons(__e, symboolean, Nil)

tmp20755 := GCons(__e, sym_1_1_6, tmp20754)

tmp20756 := GCons(__e, symboolean, tmp20755)

tmp20757 := GCons(__e, tmp20756, Nil)

tmp20758 := GCons(__e, sym_1_1_6, tmp20757)

tmp20759 := GCons(__e, symboolean, tmp20758)

return TailCall6(__e, SymFn(symis_b), V6543, tmp20759, B6539, L6540, Key6541, C6542)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20760 := GValue(__e, symshen_4_dsigf_d)

tmp20761 := Call3(__e, SymFn(symshen_4assoc_1_6), symor, tmp20753, tmp20760)


tmp20762 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20761)

_ = tmp20762

tmp20763 := MakeNative1(func(__e *ControlFlow, V6548 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6544 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6545 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6546 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6547 Obj) Obj {
tmp20764 := GCons(__e, symstring, Nil)

tmp20765 := GCons(__e, sym_1_1_6, tmp20764)

return TailCall6(__e, SymFn(symis_b), V6548, tmp20765, B6544, L6545, Key6546, C6547)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20766 := GValue(__e, symshen_4_dsigf_d)

tmp20767 := Call3(__e, SymFn(symshen_4assoc_1_6), symos, tmp20763, tmp20766)


tmp20768 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20767)

_ = tmp20768

tmp20769 := MakeNative1(func(__e *ControlFlow, V6553 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6549 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6550 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6551 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6552 Obj) Obj {
tmp20770 := GCons(__e, symboolean, Nil)

tmp20771 := GCons(__e, sym_1_1_6, tmp20770)

tmp20772 := GCons(__e, symsymbol, tmp20771)

return TailCall6(__e, SymFn(symis_b), V6553, tmp20772, B6549, L6550, Key6551, C6552)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20773 := GValue(__e, symshen_4_dsigf_d)

tmp20774 := Call3(__e, SymFn(symshen_4assoc_1_6), sympackage_2, tmp20769, tmp20773)


tmp20775 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20774)

_ = tmp20775

tmp20776 := MakeNative1(func(__e *ControlFlow, V6558 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6554 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6555 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6556 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6557 Obj) Obj {
tmp20777 := GCons(__e, symstring, Nil)

tmp20778 := GCons(__e, sym_1_1_6, tmp20777)

return TailCall6(__e, SymFn(symis_b), V6558, tmp20778, B6554, L6555, Key6556, C6557)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20779 := GValue(__e, symshen_4_dsigf_d)

tmp20780 := Call3(__e, SymFn(symshen_4assoc_1_6), symport, tmp20776, tmp20779)


tmp20781 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20780)

_ = tmp20781

tmp20782 := MakeNative1(func(__e *ControlFlow, V6563 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6559 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6560 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6561 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6562 Obj) Obj {
tmp20783 := GCons(__e, symstring, Nil)

tmp20784 := GCons(__e, sym_1_1_6, tmp20783)

return TailCall6(__e, SymFn(symis_b), V6563, tmp20784, B6559, L6560, Key6561, C6562)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20785 := GValue(__e, symshen_4_dsigf_d)

tmp20786 := Call3(__e, SymFn(symshen_4assoc_1_6), symporters, tmp20782, tmp20785)


tmp20787 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20786)

_ = tmp20787

tmp20788 := MakeNative1(func(__e *ControlFlow, V6568 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6564 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6565 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6566 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6567 Obj) Obj {
tmp20789 := GCons(__e, symstring, Nil)

tmp20790 := GCons(__e, sym_1_1_6, tmp20789)

tmp20791 := GCons(__e, symnumber, tmp20790)

tmp20792 := GCons(__e, tmp20791, Nil)

tmp20793 := GCons(__e, sym_1_1_6, tmp20792)

tmp20794 := GCons(__e, symstring, tmp20793)

return TailCall6(__e, SymFn(symis_b), V6568, tmp20794, B6564, L6565, Key6566, C6567)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20795 := GValue(__e, symshen_4_dsigf_d)

tmp20796 := Call3(__e, SymFn(symshen_4assoc_1_6), sympos, tmp20788, tmp20795)


tmp20797 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20796)

_ = tmp20797

tmp20798 := MakeNative1(func(__e *ControlFlow, V6573 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6569 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6570 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6571 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6572 Obj) Obj {
tmp20799 := GCons(__e, symout, Nil)

tmp20800 := GCons(__e, symstream, tmp20799)

tmp20801 := GCons(__e, symstring, Nil)

tmp20802 := GCons(__e, sym_1_1_6, tmp20801)

tmp20803 := GCons(__e, tmp20800, tmp20802)

tmp20804 := GCons(__e, tmp20803, Nil)

tmp20805 := GCons(__e, sym_1_1_6, tmp20804)

tmp20806 := GCons(__e, symstring, tmp20805)

return TailCall6(__e, SymFn(symis_b), V6573, tmp20806, B6569, L6570, Key6571, C6572)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20807 := GValue(__e, symshen_4_dsigf_d)

tmp20808 := Call3(__e, SymFn(symshen_4assoc_1_6), sympr, tmp20798, tmp20807)


tmp20809 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20808)

_ = tmp20809

tmp20810 := MakeNative1(func(__e *ControlFlow, V6578 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6574 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6575 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6576 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6577 Obj) Obj {
tmp20811 := Call1(__e, SymFn(symshen_4newpv), B6574)


let__18971 := tmp20811
_ = let__18971

tmp20812 := GCons(__e, let__18971, Nil)

tmp20813 := GCons(__e, sym_1_1_6, tmp20812)

tmp20814 := GCons(__e, let__18971, tmp20813)

tmp20815 := Call6(__e, SymFn(symis_b), V6578, tmp20814, B6574, L6575, Key6576, C6577)


return TailCall2(__e, SymFn(symshen_4gc), B6574, tmp20815)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20816 := GValue(__e, symshen_4_dsigf_d)

tmp20817 := Call3(__e, SymFn(symshen_4assoc_1_6), symprint, tmp20810, tmp20816)


tmp20818 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20817)

_ = tmp20818

tmp20819 := MakeNative1(func(__e *ControlFlow, V6583 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6579 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6580 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6581 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6582 Obj) Obj {
tmp20820 := GCons(__e, symsymbol, Nil)

tmp20821 := GCons(__e, sym_1_1_6, tmp20820)

tmp20822 := GCons(__e, symsymbol, tmp20821)

return TailCall6(__e, SymFn(symis_b), V6583, tmp20822, B6579, L6580, Key6581, C6582)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20823 := GValue(__e, symshen_4_dsigf_d)

tmp20824 := Call3(__e, SymFn(symshen_4assoc_1_6), symprofile, tmp20819, tmp20823)


tmp20825 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20824)

_ = tmp20825

tmp20826 := MakeNative1(func(__e *ControlFlow, V6589 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6585 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6586 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6587 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6588 Obj) Obj {
tmp20827 := GCons(__e, symsymbol, Nil)

tmp20828 := GCons(__e, symlist, tmp20827)

tmp20829 := GCons(__e, symsymbol, Nil)

tmp20830 := GCons(__e, symlist, tmp20829)

tmp20831 := GCons(__e, tmp20830, Nil)

tmp20832 := GCons(__e, sym_1_1_6, tmp20831)

tmp20833 := GCons(__e, tmp20828, tmp20832)

return TailCall6(__e, SymFn(symis_b), V6589, tmp20833, B6585, L6586, Key6587, C6588)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20834 := GValue(__e, symshen_4_dsigf_d)

tmp20835 := Call3(__e, SymFn(symshen_4assoc_1_6), sympreclude, tmp20826, tmp20834)


tmp20836 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20835)

_ = tmp20836

tmp20837 := MakeNative1(func(__e *ControlFlow, V6594 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6590 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6591 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6592 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6593 Obj) Obj {
tmp20838 := GCons(__e, symstring, Nil)

tmp20839 := GCons(__e, sym_1_1_6, tmp20838)

tmp20840 := GCons(__e, symstring, tmp20839)

return TailCall6(__e, SymFn(symis_b), V6594, tmp20840, B6590, L6591, Key6592, C6593)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20841 := GValue(__e, symshen_4_dsigf_d)

tmp20842 := Call3(__e, SymFn(symshen_4assoc_1_6), symshen_4proc_1nl, tmp20837, tmp20841)


tmp20843 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20842)

_ = tmp20843

tmp20844 := MakeNative1(func(__e *ControlFlow, V6599 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6595 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6596 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6597 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6598 Obj) Obj {
tmp20845 := GCons(__e, symnumber, Nil)

tmp20846 := GCons(__e, sym_d, tmp20845)

tmp20847 := GCons(__e, symsymbol, tmp20846)

tmp20848 := GCons(__e, tmp20847, Nil)

tmp20849 := GCons(__e, sym_1_1_6, tmp20848)

tmp20850 := GCons(__e, symsymbol, tmp20849)

return TailCall6(__e, SymFn(symis_b), V6599, tmp20850, B6595, L6596, Key6597, C6598)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20851 := GValue(__e, symshen_4_dsigf_d)

tmp20852 := Call3(__e, SymFn(symshen_4assoc_1_6), symprofile_1results, tmp20844, tmp20851)


tmp20853 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20852)

_ = tmp20853

tmp20854 := MakeNative1(func(__e *ControlFlow, V6604 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6600 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6601 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6602 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6603 Obj) Obj {
tmp20855 := Call1(__e, SymFn(symshen_4newpv), B6600)


let__18972 := tmp20855
_ = let__18972

tmp20856 := GCons(__e, let__18972, Nil)

tmp20857 := GCons(__e, sym_1_1_6, tmp20856)

tmp20858 := GCons(__e, let__18972, tmp20857)

tmp20859 := Call6(__e, SymFn(symis_b), V6604, tmp20858, B6600, L6601, Key6602, C6603)


return TailCall2(__e, SymFn(symshen_4gc), B6600, tmp20859)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20860 := GValue(__e, symshen_4_dsigf_d)

tmp20861 := Call3(__e, SymFn(symshen_4assoc_1_6), symprotect, tmp20854, tmp20860)


tmp20862 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20861)

_ = tmp20862

tmp20863 := MakeNative1(func(__e *ControlFlow, V6610 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6606 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6607 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6608 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6609 Obj) Obj {
tmp20864 := GCons(__e, symsymbol, Nil)

tmp20865 := GCons(__e, symlist, tmp20864)

tmp20866 := GCons(__e, symsymbol, Nil)

tmp20867 := GCons(__e, symlist, tmp20866)

tmp20868 := GCons(__e, tmp20867, Nil)

tmp20869 := GCons(__e, sym_1_1_6, tmp20868)

tmp20870 := GCons(__e, tmp20865, tmp20869)

return TailCall6(__e, SymFn(symis_b), V6610, tmp20870, B6606, L6607, Key6608, C6609)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20871 := GValue(__e, symshen_4_dsigf_d)

tmp20872 := Call3(__e, SymFn(symshen_4assoc_1_6), sympreclude_1all_1but, tmp20863, tmp20871)


tmp20873 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20872)

_ = tmp20873

tmp20874 := MakeNative1(func(__e *ControlFlow, V6615 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6611 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6612 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6613 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6614 Obj) Obj {
tmp20875 := GCons(__e, symout, Nil)

tmp20876 := GCons(__e, symstream, tmp20875)

tmp20877 := GCons(__e, symstring, Nil)

tmp20878 := GCons(__e, sym_1_1_6, tmp20877)

tmp20879 := GCons(__e, tmp20876, tmp20878)

tmp20880 := GCons(__e, tmp20879, Nil)

tmp20881 := GCons(__e, sym_1_1_6, tmp20880)

tmp20882 := GCons(__e, symstring, tmp20881)

return TailCall6(__e, SymFn(symis_b), V6615, tmp20882, B6611, L6612, Key6613, C6614)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20883 := GValue(__e, symshen_4_dsigf_d)

tmp20884 := Call3(__e, SymFn(symshen_4assoc_1_6), symshen_4prhush, tmp20874, tmp20883)


tmp20885 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20884)

_ = tmp20885

tmp20886 := MakeNative1(func(__e *ControlFlow, V6620 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6616 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6617 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6618 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6619 Obj) Obj {
tmp20887 := GCons(__e, symnumber, Nil)

tmp20888 := GCons(__e, sym_1_1_6, tmp20887)

tmp20889 := GCons(__e, symnumber, tmp20888)

return TailCall6(__e, SymFn(symis_b), V6620, tmp20889, B6616, L6617, Key6618, C6619)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20890 := GValue(__e, symshen_4_dsigf_d)

tmp20891 := Call3(__e, SymFn(symshen_4assoc_1_6), symprolog_1memory, tmp20886, tmp20890)


tmp20892 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20891)

_ = tmp20892

tmp20893 := MakeNative1(func(__e *ControlFlow, V6625 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6621 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6622 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6623 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6624 Obj) Obj {
tmp20894 := GCons(__e, symunit, Nil)

tmp20895 := GCons(__e, symlist, tmp20894)

tmp20896 := GCons(__e, tmp20895, Nil)

tmp20897 := GCons(__e, sym_1_1_6, tmp20896)

tmp20898 := GCons(__e, symsymbol, tmp20897)

return TailCall6(__e, SymFn(symis_b), V6625, tmp20898, B6621, L6622, Key6623, C6624)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20899 := GValue(__e, symshen_4_dsigf_d)

tmp20900 := Call3(__e, SymFn(symshen_4assoc_1_6), symps, tmp20893, tmp20899)


tmp20901 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20900)

_ = tmp20901

tmp20902 := MakeNative1(func(__e *ControlFlow, V6630 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6626 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6627 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6628 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6629 Obj) Obj {
tmp20903 := GCons(__e, symin, Nil)

tmp20904 := GCons(__e, symstream, tmp20903)

tmp20905 := GCons(__e, symunit, Nil)

tmp20906 := GCons(__e, sym_1_1_6, tmp20905)

tmp20907 := GCons(__e, tmp20904, tmp20906)

return TailCall6(__e, SymFn(symis_b), V6630, tmp20907, B6626, L6627, Key6628, C6629)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20908 := GValue(__e, symshen_4_dsigf_d)

tmp20909 := Call3(__e, SymFn(symshen_4assoc_1_6), symread, tmp20902, tmp20908)


tmp20910 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20909)

_ = tmp20910

tmp20911 := MakeNative1(func(__e *ControlFlow, V6635 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6631 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6632 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6633 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6634 Obj) Obj {
tmp20912 := GCons(__e, symin, Nil)

tmp20913 := GCons(__e, symstream, tmp20912)

tmp20914 := GCons(__e, symstring, Nil)

tmp20915 := GCons(__e, sym_1_1_6, tmp20914)

tmp20916 := GCons(__e, tmp20913, tmp20915)

return TailCall6(__e, SymFn(symis_b), V6635, tmp20916, B6631, L6632, Key6633, C6634)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20917 := GValue(__e, symshen_4_dsigf_d)

tmp20918 := Call3(__e, SymFn(symshen_4assoc_1_6), symread_1unit_1string, tmp20911, tmp20917)


tmp20919 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20918)

_ = tmp20919

tmp20920 := MakeNative1(func(__e *ControlFlow, V6640 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6636 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6637 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6638 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6639 Obj) Obj {
tmp20921 := GCons(__e, symin, Nil)

tmp20922 := GCons(__e, symstream, tmp20921)

tmp20923 := GCons(__e, symnumber, Nil)

tmp20924 := GCons(__e, sym_1_1_6, tmp20923)

tmp20925 := GCons(__e, tmp20922, tmp20924)

return TailCall6(__e, SymFn(symis_b), V6640, tmp20925, B6636, L6637, Key6638, C6639)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20926 := GValue(__e, symshen_4_dsigf_d)

tmp20927 := Call3(__e, SymFn(symshen_4assoc_1_6), symread_1byte, tmp20920, tmp20926)


tmp20928 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20927)

_ = tmp20928

tmp20929 := MakeNative1(func(__e *ControlFlow, V6645 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6641 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6642 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6643 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6644 Obj) Obj {
tmp20930 := GCons(__e, symnumber, Nil)

tmp20931 := GCons(__e, symlist, tmp20930)

tmp20932 := GCons(__e, tmp20931, Nil)

tmp20933 := GCons(__e, sym_1_1_6, tmp20932)

tmp20934 := GCons(__e, symstring, tmp20933)

return TailCall6(__e, SymFn(symis_b), V6645, tmp20934, B6641, L6642, Key6643, C6644)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20935 := GValue(__e, symshen_4_dsigf_d)

tmp20936 := Call3(__e, SymFn(symshen_4assoc_1_6), symread_1file_1as_1bytelist, tmp20929, tmp20935)


tmp20937 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20936)

_ = tmp20937

tmp20938 := MakeNative1(func(__e *ControlFlow, V6650 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6646 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6647 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6648 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6649 Obj) Obj {
tmp20939 := GCons(__e, symstring, Nil)

tmp20940 := GCons(__e, sym_1_1_6, tmp20939)

tmp20941 := GCons(__e, symstring, tmp20940)

return TailCall6(__e, SymFn(symis_b), V6650, tmp20941, B6646, L6647, Key6648, C6649)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20942 := GValue(__e, symshen_4_dsigf_d)

tmp20943 := Call3(__e, SymFn(symshen_4assoc_1_6), symread_1file_1as_1string, tmp20938, tmp20942)


tmp20944 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20943)

_ = tmp20944

tmp20945 := MakeNative1(func(__e *ControlFlow, V6655 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6651 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6652 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6653 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6654 Obj) Obj {
tmp20946 := GCons(__e, symunit, Nil)

tmp20947 := GCons(__e, symlist, tmp20946)

tmp20948 := GCons(__e, tmp20947, Nil)

tmp20949 := GCons(__e, sym_1_1_6, tmp20948)

tmp20950 := GCons(__e, symstring, tmp20949)

return TailCall6(__e, SymFn(symis_b), V6655, tmp20950, B6651, L6652, Key6653, C6654)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20951 := GValue(__e, symshen_4_dsigf_d)

tmp20952 := Call3(__e, SymFn(symshen_4assoc_1_6), symread_1file, tmp20945, tmp20951)


tmp20953 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20952)

_ = tmp20953

tmp20954 := MakeNative1(func(__e *ControlFlow, V6660 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6656 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6657 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6658 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6659 Obj) Obj {
tmp20955 := GCons(__e, symunit, Nil)

tmp20956 := GCons(__e, symlist, tmp20955)

tmp20957 := GCons(__e, tmp20956, Nil)

tmp20958 := GCons(__e, sym_1_1_6, tmp20957)

tmp20959 := GCons(__e, symstring, tmp20958)

return TailCall6(__e, SymFn(symis_b), V6660, tmp20959, B6656, L6657, Key6658, C6659)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20960 := GValue(__e, symshen_4_dsigf_d)

tmp20961 := Call3(__e, SymFn(symshen_4assoc_1_6), symread_1from_1string, tmp20954, tmp20960)


tmp20962 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20961)

_ = tmp20962

tmp20963 := MakeNative1(func(__e *ControlFlow, V6665 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6661 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6662 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6663 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6664 Obj) Obj {
tmp20964 := GCons(__e, symunit, Nil)

tmp20965 := GCons(__e, symlist, tmp20964)

tmp20966 := GCons(__e, tmp20965, Nil)

tmp20967 := GCons(__e, sym_1_1_6, tmp20966)

tmp20968 := GCons(__e, symstring, tmp20967)

return TailCall6(__e, SymFn(symis_b), V6665, tmp20968, B6661, L6662, Key6663, C6664)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20969 := GValue(__e, symshen_4_dsigf_d)

tmp20970 := Call3(__e, SymFn(symshen_4assoc_1_6), symread_1from_1string_1unprocessed, tmp20963, tmp20969)


tmp20971 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20970)

_ = tmp20971

tmp20972 := MakeNative1(func(__e *ControlFlow, V6670 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6666 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6667 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6668 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6669 Obj) Obj {
tmp20973 := GCons(__e, symstring, Nil)

tmp20974 := GCons(__e, sym_1_1_6, tmp20973)

return TailCall6(__e, SymFn(symis_b), V6670, tmp20974, B6666, L6667, Key6668, C6669)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20975 := GValue(__e, symshen_4_dsigf_d)

tmp20976 := Call3(__e, SymFn(symshen_4assoc_1_6), symrelease, tmp20972, tmp20975)


tmp20977 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20976)

_ = tmp20977

tmp20978 := MakeNative1(func(__e *ControlFlow, V6676 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6672 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6673 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6674 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6675 Obj) Obj {
tmp20979 := Call1(__e, SymFn(symshen_4newpv), B6672)


let__18973 := tmp20979
_ = let__18973

tmp20980 := GCons(__e, let__18973, Nil)

tmp20981 := GCons(__e, symlist, tmp20980)

tmp20982 := GCons(__e, let__18973, Nil)

tmp20983 := GCons(__e, symlist, tmp20982)

tmp20984 := GCons(__e, tmp20983, Nil)

tmp20985 := GCons(__e, sym_1_1_6, tmp20984)

tmp20986 := GCons(__e, tmp20981, tmp20985)

tmp20987 := GCons(__e, tmp20986, Nil)

tmp20988 := GCons(__e, sym_1_1_6, tmp20987)

tmp20989 := GCons(__e, let__18973, tmp20988)

tmp20990 := Call6(__e, SymFn(symis_b), V6676, tmp20989, B6672, L6673, Key6674, C6675)


return TailCall2(__e, SymFn(symshen_4gc), B6672, tmp20990)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20991 := GValue(__e, symshen_4_dsigf_d)

tmp20992 := Call3(__e, SymFn(symshen_4assoc_1_6), symremove, tmp20978, tmp20991)


tmp20993 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20992)

_ = tmp20993

tmp20994 := MakeNative1(func(__e *ControlFlow, V6682 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6678 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6679 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6680 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6681 Obj) Obj {
tmp20995 := Call1(__e, SymFn(symshen_4newpv), B6678)


let__18974 := tmp20995
_ = let__18974

tmp20996 := GCons(__e, let__18974, Nil)

tmp20997 := GCons(__e, symlist, tmp20996)

tmp20998 := GCons(__e, let__18974, Nil)

tmp20999 := GCons(__e, symlist, tmp20998)

tmp21000 := GCons(__e, tmp20999, Nil)

tmp21001 := GCons(__e, sym_1_1_6, tmp21000)

tmp21002 := GCons(__e, tmp20997, tmp21001)

tmp21003 := Call6(__e, SymFn(symis_b), V6682, tmp21002, B6678, L6679, Key6680, C6681)


return TailCall2(__e, SymFn(symshen_4gc), B6678, tmp21003)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp21004 := GValue(__e, symshen_4_dsigf_d)

tmp21005 := Call3(__e, SymFn(symshen_4assoc_1_6), symreverse, tmp20994, tmp21004)


tmp21006 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp21005)

_ = tmp21006

tmp21007 := MakeNative1(func(__e *ControlFlow, V6687 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6683 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6684 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6685 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6686 Obj) Obj {
tmp21008 := Call1(__e, SymFn(symshen_4newpv), B6683)


let__18975 := tmp21008
_ = let__18975

tmp21009 := GCons(__e, let__18975, Nil)

tmp21010 := GCons(__e, sym_1_1_6, tmp21009)

tmp21011 := GCons(__e, symstring, tmp21010)

tmp21012 := Call6(__e, SymFn(symis_b), V6687, tmp21011, B6683, L6684, Key6685, C6686)


return TailCall2(__e, SymFn(symshen_4gc), B6683, tmp21012)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp21013 := GValue(__e, symshen_4_dsigf_d)

tmp21014 := Call3(__e, SymFn(symshen_4assoc_1_6), symsimple_1error, tmp21007, tmp21013)


tmp21015 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp21014)

_ = tmp21015

tmp21016 := MakeNative1(func(__e *ControlFlow, V6692 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6688 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6689 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6690 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6691 Obj) Obj {
tmp21017 := Call1(__e, SymFn(symshen_4newpv), B6688)


let__18976 := tmp21017
_ = let__18976

tmp21018 := Call1(__e, SymFn(symshen_4newpv), B6688)


let__18977 := tmp21018
_ = let__18977

tmp21019 := GCons(__e, let__18977, Nil)

tmp21020 := GCons(__e, sym_d, tmp21019)

tmp21021 := GCons(__e, let__18976, tmp21020)

tmp21022 := GCons(__e, let__18977, Nil)

tmp21023 := GCons(__e, sym_1_1_6, tmp21022)

tmp21024 := GCons(__e, tmp21021, tmp21023)

tmp21025 := Call6(__e, SymFn(symis_b), V6692, tmp21024, B6688, L6689, Key6690, C6691)


tmp21026 := Call2(__e, SymFn(symshen_4gc), B6688, tmp21025)


return TailCall2(__e, SymFn(symshen_4gc), B6688, tmp21026)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp21027 := GValue(__e, symshen_4_dsigf_d)

tmp21028 := Call3(__e, SymFn(symshen_4assoc_1_6), symsnd, tmp21016, tmp21027)


tmp21029 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp21028)

_ = tmp21029

tmp21030 := MakeNative1(func(__e *ControlFlow, V6697 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6693 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6694 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6695 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6696 Obj) Obj {
tmp21031 := GCons(__e, symsymbol, Nil)

tmp21032 := GCons(__e, sym_1_1_6, tmp21031)

tmp21033 := GCons(__e, symnumber, tmp21032)

tmp21034 := GCons(__e, tmp21033, Nil)

tmp21035 := GCons(__e, sym_1_1_6, tmp21034)

tmp21036 := GCons(__e, symsymbol, tmp21035)

return TailCall6(__e, SymFn(symis_b), V6697, tmp21036, B6693, L6694, Key6695, C6696)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp21037 := GValue(__e, symshen_4_dsigf_d)

tmp21038 := Call3(__e, SymFn(symshen_4assoc_1_6), symspecialise, tmp21030, tmp21037)


tmp21039 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp21038)

_ = tmp21039

tmp21040 := MakeNative1(func(__e *ControlFlow, V6702 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6698 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6699 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6700 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6701 Obj) Obj {
tmp21041 := GCons(__e, symboolean, Nil)

tmp21042 := GCons(__e, sym_1_1_6, tmp21041)

tmp21043 := GCons(__e, symsymbol, tmp21042)

return TailCall6(__e, SymFn(symis_b), V6702, tmp21043, B6698, L6699, Key6700, C6701)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp21044 := GValue(__e, symshen_4_dsigf_d)

tmp21045 := Call3(__e, SymFn(symshen_4assoc_1_6), symspy, tmp21040, tmp21044)


tmp21046 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp21045)

_ = tmp21046

tmp21047 := MakeNative1(func(__e *ControlFlow, V6707 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6703 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6704 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6705 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6706 Obj) Obj {
tmp21048 := GCons(__e, symboolean, Nil)

tmp21049 := GCons(__e, sym_1_1_6, tmp21048)

return TailCall6(__e, SymFn(symis_b), V6707, tmp21049, B6703, L6704, Key6705, C6706)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp21050 := GValue(__e, symshen_4_dsigf_d)

tmp21051 := Call3(__e, SymFn(symshen_4assoc_1_6), symspy_2, tmp21047, tmp21050)


tmp21052 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp21051)

_ = tmp21052

tmp21053 := MakeNative1(func(__e *ControlFlow, V6712 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6708 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6709 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6710 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6711 Obj) Obj {
tmp21054 := GCons(__e, symboolean, Nil)

tmp21055 := GCons(__e, sym_1_1_6, tmp21054)

tmp21056 := GCons(__e, symsymbol, tmp21055)

return TailCall6(__e, SymFn(symis_b), V6712, tmp21056, B6708, L6709, Key6710, C6711)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp21057 := GValue(__e, symshen_4_dsigf_d)

tmp21058 := Call3(__e, SymFn(symshen_4assoc_1_6), symstep, tmp21053, tmp21057)


tmp21059 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp21058)

_ = tmp21059

tmp21060 := MakeNative1(func(__e *ControlFlow, V6717 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6713 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6714 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6715 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6716 Obj) Obj {
tmp21061 := GCons(__e, symboolean, Nil)

tmp21062 := GCons(__e, sym_1_1_6, tmp21061)

return TailCall6(__e, SymFn(symis_b), V6717, tmp21062, B6713, L6714, Key6715, C6716)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp21063 := GValue(__e, symshen_4_dsigf_d)

tmp21064 := Call3(__e, SymFn(symshen_4assoc_1_6), symstep_2, tmp21060, tmp21063)


tmp21065 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp21064)

_ = tmp21065

tmp21066 := MakeNative1(func(__e *ControlFlow, V6722 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6718 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6719 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6720 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6721 Obj) Obj {
tmp21067 := GCons(__e, symin, Nil)

tmp21068 := GCons(__e, symstream, tmp21067)

tmp21069 := GCons(__e, tmp21068, Nil)

tmp21070 := GCons(__e, sym_1_1_6, tmp21069)

return TailCall6(__e, SymFn(symis_b), V6722, tmp21070, B6718, L6719, Key6720, C6721)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp21071 := GValue(__e, symshen_4_dsigf_d)

tmp21072 := Call3(__e, SymFn(symshen_4assoc_1_6), symstinput, tmp21066, tmp21071)


tmp21073 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp21072)

_ = tmp21073

tmp21074 := MakeNative1(func(__e *ControlFlow, V6727 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6723 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6724 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6725 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6726 Obj) Obj {
tmp21075 := GCons(__e, symout, Nil)

tmp21076 := GCons(__e, symstream, tmp21075)

tmp21077 := GCons(__e, tmp21076, Nil)

tmp21078 := GCons(__e, sym_1_1_6, tmp21077)

return TailCall6(__e, SymFn(symis_b), V6727, tmp21078, B6723, L6724, Key6725, C6726)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp21079 := GValue(__e, symshen_4_dsigf_d)

tmp21080 := Call3(__e, SymFn(symshen_4assoc_1_6), symsterror, tmp21074, tmp21079)


tmp21081 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp21080)

_ = tmp21081

tmp21082 := MakeNative1(func(__e *ControlFlow, V6732 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6728 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6729 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6730 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6731 Obj) Obj {
tmp21083 := GCons(__e, symout, Nil)

tmp21084 := GCons(__e, symstream, tmp21083)

tmp21085 := GCons(__e, tmp21084, Nil)

tmp21086 := GCons(__e, sym_1_1_6, tmp21085)

return TailCall6(__e, SymFn(symis_b), V6732, tmp21086, B6728, L6729, Key6730, C6731)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp21087 := GValue(__e, symshen_4_dsigf_d)

tmp21088 := Call3(__e, SymFn(symshen_4assoc_1_6), symstoutput, tmp21082, tmp21087)


tmp21089 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp21088)

_ = tmp21089

tmp21090 := MakeNative1(func(__e *ControlFlow, V6737 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6733 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6734 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6735 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6736 Obj) Obj {
tmp21091 := Call1(__e, SymFn(symshen_4newpv), B6733)


let__18978 := tmp21091
_ = let__18978

tmp21092 := GCons(__e, symboolean, Nil)

tmp21093 := GCons(__e, sym_1_1_6, tmp21092)

tmp21094 := GCons(__e, let__18978, tmp21093)

tmp21095 := Call6(__e, SymFn(symis_b), V6737, tmp21094, B6733, L6734, Key6735, C6736)


return TailCall2(__e, SymFn(symshen_4gc), B6733, tmp21095)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp21096 := GValue(__e, symshen_4_dsigf_d)

tmp21097 := Call3(__e, SymFn(symshen_4assoc_1_6), symstring_2, tmp21090, tmp21096)


tmp21098 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp21097)

_ = tmp21098

tmp21099 := MakeNative1(func(__e *ControlFlow, V6742 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6738 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6739 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6740 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6741 Obj) Obj {
tmp21100 := Call1(__e, SymFn(symshen_4newpv), B6738)


let__18979 := tmp21100
_ = let__18979

tmp21101 := GCons(__e, symstring, Nil)

tmp21102 := GCons(__e, sym_1_1_6, tmp21101)

tmp21103 := GCons(__e, let__18979, tmp21102)

tmp21104 := Call6(__e, SymFn(symis_b), V6742, tmp21103, B6738, L6739, Key6740, C6741)


return TailCall2(__e, SymFn(symshen_4gc), B6738, tmp21104)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp21105 := GValue(__e, symshen_4_dsigf_d)

tmp21106 := Call3(__e, SymFn(symshen_4assoc_1_6), symstr, tmp21099, tmp21105)


tmp21107 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp21106)

_ = tmp21107

tmp21108 := MakeNative1(func(__e *ControlFlow, V6747 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6743 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6744 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6745 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6746 Obj) Obj {
tmp21109 := GCons(__e, symnumber, Nil)

tmp21110 := GCons(__e, sym_1_1_6, tmp21109)

tmp21111 := GCons(__e, symstring, tmp21110)

return TailCall6(__e, SymFn(symis_b), V6747, tmp21111, B6743, L6744, Key6745, C6746)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp21112 := GValue(__e, symshen_4_dsigf_d)

tmp21113 := Call3(__e, SymFn(symshen_4assoc_1_6), symstring_1_6n, tmp21108, tmp21112)


tmp21114 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp21113)

_ = tmp21114

tmp21115 := MakeNative1(func(__e *ControlFlow, V6752 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6748 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6749 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6750 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6751 Obj) Obj {
tmp21116 := GCons(__e, symsymbol, Nil)

tmp21117 := GCons(__e, sym_1_1_6, tmp21116)

tmp21118 := GCons(__e, symstring, tmp21117)

return TailCall6(__e, SymFn(symis_b), V6752, tmp21118, B6748, L6749, Key6750, C6751)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp21119 := GValue(__e, symshen_4_dsigf_d)

tmp21120 := Call3(__e, SymFn(symshen_4assoc_1_6), symstring_1_6symbol, tmp21115, tmp21119)


tmp21121 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp21120)

_ = tmp21121

tmp21122 := MakeNative1(func(__e *ControlFlow, V6757 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6753 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6754 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6755 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6756 Obj) Obj {
tmp21123 := GCons(__e, symnumber, Nil)

tmp21124 := GCons(__e, symlist, tmp21123)

tmp21125 := GCons(__e, symnumber, Nil)

tmp21126 := GCons(__e, sym_1_1_6, tmp21125)

tmp21127 := GCons(__e, tmp21124, tmp21126)

return TailCall6(__e, SymFn(symis_b), V6757, tmp21127, B6753, L6754, Key6755, C6756)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp21128 := GValue(__e, symshen_4_dsigf_d)

tmp21129 := Call3(__e, SymFn(symshen_4assoc_1_6), symsum, tmp21122, tmp21128)


tmp21130 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp21129)

_ = tmp21130

tmp21131 := MakeNative1(func(__e *ControlFlow, V6762 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6758 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6759 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6760 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6761 Obj) Obj {
tmp21132 := Call1(__e, SymFn(symshen_4newpv), B6758)


let__18980 := tmp21132
_ = let__18980

tmp21133 := GCons(__e, symboolean, Nil)

tmp21134 := GCons(__e, sym_1_1_6, tmp21133)

tmp21135 := GCons(__e, let__18980, tmp21134)

tmp21136 := Call6(__e, SymFn(symis_b), V6762, tmp21135, B6758, L6759, Key6760, C6761)


return TailCall2(__e, SymFn(symshen_4gc), B6758, tmp21136)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp21137 := GValue(__e, symshen_4_dsigf_d)

tmp21138 := Call3(__e, SymFn(symshen_4assoc_1_6), symsymbol_2, tmp21131, tmp21137)


tmp21139 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp21138)

_ = tmp21139

tmp21140 := MakeNative1(func(__e *ControlFlow, V6767 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6763 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6764 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6765 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6766 Obj) Obj {
tmp21141 := GCons(__e, symsymbol, Nil)

tmp21142 := GCons(__e, sym_1_1_6, tmp21141)

tmp21143 := GCons(__e, symsymbol, tmp21142)

return TailCall6(__e, SymFn(symis_b), V6767, tmp21143, B6763, L6764, Key6765, C6766)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp21144 := GValue(__e, symshen_4_dsigf_d)

tmp21145 := Call3(__e, SymFn(symshen_4assoc_1_6), symsystemf, tmp21140, tmp21144)


tmp21146 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp21145)

_ = tmp21146

tmp21147 := MakeNative1(func(__e *ControlFlow, V6772 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6768 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6769 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6770 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6771 Obj) Obj {
tmp21148 := GCons(__e, symboolean, Nil)

tmp21149 := GCons(__e, sym_1_1_6, tmp21148)

return TailCall6(__e, SymFn(symis_b), V6772, tmp21149, B6768, L6769, Key6770, C6771)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp21150 := GValue(__e, symshen_4_dsigf_d)

tmp21151 := Call3(__e, SymFn(symshen_4assoc_1_6), symsystem_1S_2, tmp21147, tmp21150)


tmp21152 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp21151)

_ = tmp21152

tmp21153 := MakeNative1(func(__e *ControlFlow, V6778 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6774 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6775 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6776 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6777 Obj) Obj {
tmp21154 := Call1(__e, SymFn(symshen_4newpv), B6774)


let__18981 := tmp21154
_ = let__18981

tmp21155 := GCons(__e, let__18981, Nil)

tmp21156 := GCons(__e, symlist, tmp21155)

tmp21157 := GCons(__e, let__18981, Nil)

tmp21158 := GCons(__e, symlist, tmp21157)

tmp21159 := GCons(__e, tmp21158, Nil)

tmp21160 := GCons(__e, sym_1_1_6, tmp21159)

tmp21161 := GCons(__e, tmp21156, tmp21160)

tmp21162 := Call6(__e, SymFn(symis_b), V6778, tmp21161, B6774, L6775, Key6776, C6777)


return TailCall2(__e, SymFn(symshen_4gc), B6774, tmp21162)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp21163 := GValue(__e, symshen_4_dsigf_d)

tmp21164 := Call3(__e, SymFn(symshen_4assoc_1_6), symtail, tmp21153, tmp21163)


tmp21165 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp21164)

_ = tmp21165

tmp21166 := MakeNative1(func(__e *ControlFlow, V6783 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6779 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6780 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6781 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6782 Obj) Obj {
tmp21167 := GCons(__e, symstring, Nil)

tmp21168 := GCons(__e, sym_1_1_6, tmp21167)

tmp21169 := GCons(__e, symstring, tmp21168)

return TailCall6(__e, SymFn(symis_b), V6783, tmp21169, B6779, L6780, Key6781, C6782)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp21170 := GValue(__e, symshen_4_dsigf_d)

tmp21171 := Call3(__e, SymFn(symshen_4assoc_1_6), symtlstr, tmp21166, tmp21170)


tmp21172 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp21171)

_ = tmp21172

tmp21173 := MakeNative1(func(__e *ControlFlow, V6789 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6785 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6786 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6787 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6788 Obj) Obj {
tmp21174 := Call1(__e, SymFn(symshen_4newpv), B6785)


let__18982 := tmp21174
_ = let__18982

tmp21175 := GCons(__e, let__18982, Nil)

tmp21176 := GCons(__e, symvector, tmp21175)

tmp21177 := GCons(__e, let__18982, Nil)

tmp21178 := GCons(__e, symvector, tmp21177)

tmp21179 := GCons(__e, tmp21178, Nil)

tmp21180 := GCons(__e, sym_1_1_6, tmp21179)

tmp21181 := GCons(__e, tmp21176, tmp21180)

tmp21182 := Call6(__e, SymFn(symis_b), V6789, tmp21181, B6785, L6786, Key6787, C6788)


return TailCall2(__e, SymFn(symshen_4gc), B6785, tmp21182)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp21183 := GValue(__e, symshen_4_dsigf_d)

tmp21184 := Call3(__e, SymFn(symshen_4assoc_1_6), symtlv, tmp21173, tmp21183)


tmp21185 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp21184)

_ = tmp21185

tmp21186 := MakeNative1(func(__e *ControlFlow, V6794 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6790 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6791 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6792 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6793 Obj) Obj {
tmp21187 := GCons(__e, symboolean, Nil)

tmp21188 := GCons(__e, sym_1_1_6, tmp21187)

tmp21189 := GCons(__e, symsymbol, tmp21188)

return TailCall6(__e, SymFn(symis_b), V6794, tmp21189, B6790, L6791, Key6792, C6793)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp21190 := GValue(__e, symshen_4_dsigf_d)

tmp21191 := Call3(__e, SymFn(symshen_4assoc_1_6), symtc, tmp21186, tmp21190)


tmp21192 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp21191)

_ = tmp21192

tmp21193 := MakeNative1(func(__e *ControlFlow, V6799 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6795 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6796 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6797 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6798 Obj) Obj {
tmp21194 := GCons(__e, symboolean, Nil)

tmp21195 := GCons(__e, sym_1_1_6, tmp21194)

return TailCall6(__e, SymFn(symis_b), V6799, tmp21195, B6795, L6796, Key6797, C6798)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp21196 := GValue(__e, symshen_4_dsigf_d)

tmp21197 := Call3(__e, SymFn(symshen_4assoc_1_6), symtc_2, tmp21193, tmp21196)


tmp21198 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp21197)

_ = tmp21198

tmp21199 := MakeNative1(func(__e *ControlFlow, V6804 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6800 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6801 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6802 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6803 Obj) Obj {
tmp21200 := Call1(__e, SymFn(symshen_4newpv), B6800)


let__18983 := tmp21200
_ = let__18983

tmp21201 := GCons(__e, let__18983, Nil)

tmp21202 := GCons(__e, symlazy, tmp21201)

tmp21203 := GCons(__e, let__18983, Nil)

tmp21204 := GCons(__e, sym_1_1_6, tmp21203)

tmp21205 := GCons(__e, tmp21202, tmp21204)

tmp21206 := Call6(__e, SymFn(symis_b), V6804, tmp21205, B6800, L6801, Key6802, C6803)


return TailCall2(__e, SymFn(symshen_4gc), B6800, tmp21206)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp21207 := GValue(__e, symshen_4_dsigf_d)

tmp21208 := Call3(__e, SymFn(symshen_4assoc_1_6), symthaw, tmp21199, tmp21207)


tmp21209 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp21208)

_ = tmp21209

tmp21210 := MakeNative1(func(__e *ControlFlow, V6809 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6805 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6806 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6807 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6808 Obj) Obj {
tmp21211 := GCons(__e, symsymbol, Nil)

tmp21212 := GCons(__e, sym_1_1_6, tmp21211)

tmp21213 := GCons(__e, symsymbol, tmp21212)

return TailCall6(__e, SymFn(symis_b), V6809, tmp21213, B6805, L6806, Key6807, C6808)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp21214 := GValue(__e, symshen_4_dsigf_d)

tmp21215 := Call3(__e, SymFn(symshen_4assoc_1_6), symtrack, tmp21210, tmp21214)


tmp21216 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp21215)

_ = tmp21216

tmp21217 := MakeNative1(func(__e *ControlFlow, V6814 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6810 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6811 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6812 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6813 Obj) Obj {
tmp21218 := GCons(__e, symsymbol, Nil)

tmp21219 := GCons(__e, symlist, tmp21218)

tmp21220 := GCons(__e, tmp21219, Nil)

tmp21221 := GCons(__e, sym_1_1_6, tmp21220)

return TailCall6(__e, SymFn(symis_b), V6814, tmp21221, B6810, L6811, Key6812, C6813)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp21222 := GValue(__e, symshen_4_dsigf_d)

tmp21223 := Call3(__e, SymFn(symshen_4assoc_1_6), symtracked, tmp21217, tmp21222)


tmp21224 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp21223)

_ = tmp21224

tmp21225 := MakeNative1(func(__e *ControlFlow, V6819 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6815 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6816 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6817 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6818 Obj) Obj {
tmp21226 := Call1(__e, SymFn(symshen_4newpv), B6815)


let__18984 := tmp21226
_ = let__18984

tmp21227 := GCons(__e, let__18984, Nil)

tmp21228 := GCons(__e, sym_1_1_6, tmp21227)

tmp21229 := GCons(__e, symexception, tmp21228)

tmp21230 := GCons(__e, let__18984, Nil)

tmp21231 := GCons(__e, sym_1_1_6, tmp21230)

tmp21232 := GCons(__e, tmp21229, tmp21231)

tmp21233 := GCons(__e, tmp21232, Nil)

tmp21234 := GCons(__e, sym_1_1_6, tmp21233)

tmp21235 := GCons(__e, let__18984, tmp21234)

tmp21236 := Call6(__e, SymFn(symis_b), V6819, tmp21235, B6815, L6816, Key6817, C6818)


return TailCall2(__e, SymFn(symshen_4gc), B6815, tmp21236)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp21237 := GValue(__e, symshen_4_dsigf_d)

tmp21238 := Call3(__e, SymFn(symshen_4assoc_1_6), symtrap_1error, tmp21225, tmp21237)


tmp21239 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp21238)

_ = tmp21239

tmp21240 := MakeNative1(func(__e *ControlFlow, V6824 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6820 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6821 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6822 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6823 Obj) Obj {
tmp21241 := Call1(__e, SymFn(symshen_4newpv), B6820)


let__18985 := tmp21241
_ = let__18985

tmp21242 := GCons(__e, symboolean, Nil)

tmp21243 := GCons(__e, sym_1_1_6, tmp21242)

tmp21244 := GCons(__e, let__18985, tmp21243)

tmp21245 := Call6(__e, SymFn(symis_b), V6824, tmp21244, B6820, L6821, Key6822, C6823)


return TailCall2(__e, SymFn(symshen_4gc), B6820, tmp21245)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp21246 := GValue(__e, symshen_4_dsigf_d)

tmp21247 := Call3(__e, SymFn(symshen_4assoc_1_6), symtuple_2, tmp21240, tmp21246)


tmp21248 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp21247)

_ = tmp21248

tmp21249 := MakeNative1(func(__e *ControlFlow, V6829 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6825 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6826 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6827 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6828 Obj) Obj {
tmp21250 := GCons(__e, symstring, Nil)

tmp21251 := GCons(__e, symlist, tmp21250)

tmp21252 := GCons(__e, tmp21251, Nil)

tmp21253 := GCons(__e, sym_1_1_6, tmp21252)

tmp21254 := GCons(__e, symstring, tmp21253)

return TailCall6(__e, SymFn(symis_b), V6829, tmp21254, B6825, L6826, Key6827, C6828)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp21255 := GValue(__e, symshen_4_dsigf_d)

tmp21256 := Call3(__e, SymFn(symshen_4assoc_1_6), symunabsolute, tmp21249, tmp21255)


tmp21257 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp21256)

_ = tmp21257

tmp21258 := MakeNative1(func(__e *ControlFlow, V6834 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6830 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6831 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6832 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6833 Obj) Obj {
tmp21259 := GCons(__e, symsymbol, Nil)

tmp21260 := GCons(__e, sym_1_1_6, tmp21259)

tmp21261 := GCons(__e, symsymbol, tmp21260)

return TailCall6(__e, SymFn(symis_b), V6834, tmp21261, B6830, L6831, Key6832, C6833)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp21262 := GValue(__e, symshen_4_dsigf_d)

tmp21263 := Call3(__e, SymFn(symshen_4assoc_1_6), symundefmacro, tmp21258, tmp21262)


tmp21264 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp21263)

_ = tmp21264

tmp21265 := MakeNative1(func(__e *ControlFlow, V6841 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6837 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6838 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6839 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6840 Obj) Obj {
tmp21266 := Call1(__e, SymFn(symshen_4newpv), B6837)


let__18986 := tmp21266
_ = let__18986

tmp21267 := GCons(__e, let__18986, Nil)

tmp21268 := GCons(__e, symlist, tmp21267)

tmp21269 := GCons(__e, let__18986, Nil)

tmp21270 := GCons(__e, symlist, tmp21269)

tmp21271 := GCons(__e, let__18986, Nil)

tmp21272 := GCons(__e, symlist, tmp21271)

tmp21273 := GCons(__e, tmp21272, Nil)

tmp21274 := GCons(__e, sym_1_1_6, tmp21273)

tmp21275 := GCons(__e, tmp21270, tmp21274)

tmp21276 := GCons(__e, tmp21275, Nil)

tmp21277 := GCons(__e, sym_1_1_6, tmp21276)

tmp21278 := GCons(__e, tmp21268, tmp21277)

tmp21279 := Call6(__e, SymFn(symis_b), V6841, tmp21278, B6837, L6838, Key6839, C6840)


return TailCall2(__e, SymFn(symshen_4gc), B6837, tmp21279)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp21280 := GValue(__e, symshen_4_dsigf_d)

tmp21281 := Call3(__e, SymFn(symshen_4assoc_1_6), symunion, tmp21265, tmp21280)


tmp21282 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp21281)

_ = tmp21282

tmp21283 := MakeNative1(func(__e *ControlFlow, V6846 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6842 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6843 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6844 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6845 Obj) Obj {
tmp21284 := GCons(__e, symsymbol, Nil)

tmp21285 := GCons(__e, sym_1_1_6, tmp21284)

tmp21286 := GCons(__e, symsymbol, tmp21285)

return TailCall6(__e, SymFn(symis_b), V6846, tmp21286, B6842, L6843, Key6844, C6845)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp21287 := GValue(__e, symshen_4_dsigf_d)

tmp21288 := Call3(__e, SymFn(symshen_4assoc_1_6), symunprofile, tmp21283, tmp21287)


tmp21289 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp21288)

_ = tmp21289

tmp21290 := MakeNative1(func(__e *ControlFlow, V6851 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6847 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6848 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6849 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6850 Obj) Obj {
tmp21291 := GCons(__e, symsymbol, Nil)

tmp21292 := GCons(__e, sym_1_1_6, tmp21291)

tmp21293 := GCons(__e, symsymbol, tmp21292)

return TailCall6(__e, SymFn(symis_b), V6851, tmp21293, B6847, L6848, Key6849, C6850)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp21294 := GValue(__e, symshen_4_dsigf_d)

tmp21295 := Call3(__e, SymFn(symshen_4assoc_1_6), symuntrack, tmp21290, tmp21294)


tmp21296 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp21295)

_ = tmp21296

tmp21297 := MakeNative1(func(__e *ControlFlow, V6856 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6852 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6853 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6854 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6855 Obj) Obj {
tmp21298 := GCons(__e, symsymbol, Nil)

tmp21299 := GCons(__e, symlist, tmp21298)

tmp21300 := GCons(__e, tmp21299, Nil)

tmp21301 := GCons(__e, sym_1_1_6, tmp21300)

return TailCall6(__e, SymFn(symis_b), V6856, tmp21301, B6852, L6853, Key6854, C6855)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp21302 := GValue(__e, symshen_4_dsigf_d)

tmp21303 := Call3(__e, SymFn(symshen_4assoc_1_6), symuserdefs, tmp21297, tmp21302)


tmp21304 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp21303)

_ = tmp21304

tmp21305 := MakeNative1(func(__e *ControlFlow, V6861 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6857 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6858 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6859 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6860 Obj) Obj {
tmp21306 := Call1(__e, SymFn(symshen_4newpv), B6857)


let__18987 := tmp21306
_ = let__18987

tmp21307 := GCons(__e, symboolean, Nil)

tmp21308 := GCons(__e, sym_1_1_6, tmp21307)

tmp21309 := GCons(__e, let__18987, tmp21308)

tmp21310 := Call6(__e, SymFn(symis_b), V6861, tmp21309, B6857, L6858, Key6859, C6860)


return TailCall2(__e, SymFn(symshen_4gc), B6857, tmp21310)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp21311 := GValue(__e, symshen_4_dsigf_d)

tmp21312 := Call3(__e, SymFn(symshen_4assoc_1_6), symvariable_2, tmp21305, tmp21311)


tmp21313 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp21312)

_ = tmp21313

tmp21314 := MakeNative1(func(__e *ControlFlow, V6866 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6862 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6863 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6864 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6865 Obj) Obj {
tmp21315 := Call1(__e, SymFn(symshen_4newpv), B6862)


let__18988 := tmp21315
_ = let__18988

tmp21316 := GCons(__e, symboolean, Nil)

tmp21317 := GCons(__e, sym_1_1_6, tmp21316)

tmp21318 := GCons(__e, let__18988, tmp21317)

tmp21319 := Call6(__e, SymFn(symis_b), V6866, tmp21318, B6862, L6863, Key6864, C6865)


return TailCall2(__e, SymFn(symshen_4gc), B6862, tmp21319)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp21320 := GValue(__e, symshen_4_dsigf_d)

tmp21321 := Call3(__e, SymFn(symshen_4assoc_1_6), symvector_2, tmp21314, tmp21320)


tmp21322 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp21321)

_ = tmp21322

tmp21323 := MakeNative1(func(__e *ControlFlow, V6871 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6867 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6868 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6869 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6870 Obj) Obj {
tmp21324 := GCons(__e, symstring, Nil)

tmp21325 := GCons(__e, sym_1_1_6, tmp21324)

return TailCall6(__e, SymFn(symis_b), V6871, tmp21325, B6867, L6868, Key6869, C6870)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp21326 := GValue(__e, symshen_4_dsigf_d)

tmp21327 := Call3(__e, SymFn(symshen_4assoc_1_6), symversion, tmp21323, tmp21326)


tmp21328 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp21327)

_ = tmp21328

tmp21329 := MakeNative1(func(__e *ControlFlow, V6876 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6872 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6873 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6874 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6875 Obj) Obj {
tmp21330 := Call1(__e, SymFn(symshen_4newpv), B6872)


let__18989 := tmp21330
_ = let__18989

tmp21331 := GCons(__e, let__18989, Nil)

tmp21332 := GCons(__e, sym_1_1_6, tmp21331)

tmp21333 := GCons(__e, let__18989, tmp21332)

tmp21334 := GCons(__e, tmp21333, Nil)

tmp21335 := GCons(__e, sym_1_1_6, tmp21334)

tmp21336 := GCons(__e, symstring, tmp21335)

tmp21337 := Call6(__e, SymFn(symis_b), V6876, tmp21336, B6872, L6873, Key6874, C6875)


return TailCall2(__e, SymFn(symshen_4gc), B6872, tmp21337)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp21338 := GValue(__e, symshen_4_dsigf_d)

tmp21339 := Call3(__e, SymFn(symshen_4assoc_1_6), symwrite_1to_1file, tmp21329, tmp21338)


tmp21340 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp21339)

_ = tmp21340

tmp21341 := MakeNative1(func(__e *ControlFlow, V6881 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6877 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6878 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6879 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6880 Obj) Obj {
tmp21342 := GCons(__e, symout, Nil)

tmp21343 := GCons(__e, symstream, tmp21342)

tmp21344 := GCons(__e, symnumber, Nil)

tmp21345 := GCons(__e, sym_1_1_6, tmp21344)

tmp21346 := GCons(__e, tmp21343, tmp21345)

tmp21347 := GCons(__e, tmp21346, Nil)

tmp21348 := GCons(__e, sym_1_1_6, tmp21347)

tmp21349 := GCons(__e, symnumber, tmp21348)

return TailCall6(__e, SymFn(symis_b), V6881, tmp21349, B6877, L6878, Key6879, C6880)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp21350 := GValue(__e, symshen_4_dsigf_d)

tmp21351 := Call3(__e, SymFn(symshen_4assoc_1_6), symwrite_1byte, tmp21341, tmp21350)


tmp21352 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp21351)

_ = tmp21352

tmp21353 := MakeNative1(func(__e *ControlFlow, V6886 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6882 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6883 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6884 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6885 Obj) Obj {
tmp21354 := GCons(__e, symout, Nil)

tmp21355 := GCons(__e, symstream, tmp21354)

tmp21356 := GCons(__e, symstring, Nil)

tmp21357 := GCons(__e, sym_1_1_6, tmp21356)

tmp21358 := GCons(__e, tmp21355, tmp21357)

tmp21359 := GCons(__e, tmp21358, Nil)

tmp21360 := GCons(__e, sym_1_1_6, tmp21359)

tmp21361 := GCons(__e, symstring, tmp21360)

return TailCall6(__e, SymFn(symis_b), V6886, tmp21361, B6882, L6883, Key6884, C6885)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp21362 := GValue(__e, symshen_4_dsigf_d)

tmp21363 := Call3(__e, SymFn(symshen_4assoc_1_6), symwrite_1string, tmp21353, tmp21362)


tmp21364 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp21363)

_ = tmp21364

tmp21365 := MakeNative1(func(__e *ControlFlow, V6891 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6887 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6888 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6889 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6890 Obj) Obj {
tmp21366 := GCons(__e, symboolean, Nil)

tmp21367 := GCons(__e, sym_1_1_6, tmp21366)

tmp21368 := GCons(__e, symstring, tmp21367)

return TailCall6(__e, SymFn(symis_b), V6891, tmp21368, B6887, L6888, Key6889, C6890)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp21369 := GValue(__e, symshen_4_dsigf_d)

tmp21370 := Call3(__e, SymFn(symshen_4assoc_1_6), symy_1or_1n_2, tmp21365, tmp21369)


tmp21371 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp21370)

_ = tmp21371

tmp21372 := MakeNative1(func(__e *ControlFlow, V6896 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6892 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6893 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6894 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6895 Obj) Obj {
tmp21373 := GCons(__e, symboolean, Nil)

tmp21374 := GCons(__e, sym_1_1_6, tmp21373)

tmp21375 := GCons(__e, symnumber, tmp21374)

tmp21376 := GCons(__e, tmp21375, Nil)

tmp21377 := GCons(__e, sym_1_1_6, tmp21376)

tmp21378 := GCons(__e, symnumber, tmp21377)

return TailCall6(__e, SymFn(symis_b), V6896, tmp21378, B6892, L6893, Key6894, C6895)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp21379 := GValue(__e, symshen_4_dsigf_d)

tmp21380 := Call3(__e, SymFn(symshen_4assoc_1_6), sym_6, tmp21372, tmp21379)


tmp21381 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp21380)

_ = tmp21381

tmp21382 := MakeNative1(func(__e *ControlFlow, V6901 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6897 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6898 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6899 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6900 Obj) Obj {
tmp21383 := GCons(__e, symboolean, Nil)

tmp21384 := GCons(__e, sym_1_1_6, tmp21383)

tmp21385 := GCons(__e, symnumber, tmp21384)

tmp21386 := GCons(__e, tmp21385, Nil)

tmp21387 := GCons(__e, sym_1_1_6, tmp21386)

tmp21388 := GCons(__e, symnumber, tmp21387)

return TailCall6(__e, SymFn(symis_b), V6901, tmp21388, B6897, L6898, Key6899, C6900)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp21389 := GValue(__e, symshen_4_dsigf_d)

tmp21390 := Call3(__e, SymFn(symshen_4assoc_1_6), sym_5, tmp21382, tmp21389)


tmp21391 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp21390)

_ = tmp21391

tmp21392 := MakeNative1(func(__e *ControlFlow, V6906 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6902 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6903 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6904 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6905 Obj) Obj {
tmp21393 := GCons(__e, symboolean, Nil)

tmp21394 := GCons(__e, sym_1_1_6, tmp21393)

tmp21395 := GCons(__e, symnumber, tmp21394)

tmp21396 := GCons(__e, tmp21395, Nil)

tmp21397 := GCons(__e, sym_1_1_6, tmp21396)

tmp21398 := GCons(__e, symnumber, tmp21397)

return TailCall6(__e, SymFn(symis_b), V6906, tmp21398, B6902, L6903, Key6904, C6905)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp21399 := GValue(__e, symshen_4_dsigf_d)

tmp21400 := Call3(__e, SymFn(symshen_4assoc_1_6), sym_6_a, tmp21392, tmp21399)


tmp21401 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp21400)

_ = tmp21401

tmp21402 := MakeNative1(func(__e *ControlFlow, V6911 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6907 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6908 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6909 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6910 Obj) Obj {
tmp21403 := GCons(__e, symboolean, Nil)

tmp21404 := GCons(__e, sym_1_1_6, tmp21403)

tmp21405 := GCons(__e, symnumber, tmp21404)

tmp21406 := GCons(__e, tmp21405, Nil)

tmp21407 := GCons(__e, sym_1_1_6, tmp21406)

tmp21408 := GCons(__e, symnumber, tmp21407)

return TailCall6(__e, SymFn(symis_b), V6911, tmp21408, B6907, L6908, Key6909, C6910)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp21409 := GValue(__e, symshen_4_dsigf_d)

tmp21410 := Call3(__e, SymFn(symshen_4assoc_1_6), sym_5_a, tmp21402, tmp21409)


tmp21411 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp21410)

_ = tmp21411

tmp21412 := MakeNative1(func(__e *ControlFlow, V6916 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6912 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6913 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6914 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6915 Obj) Obj {
tmp21413 := Call1(__e, SymFn(symshen_4newpv), B6912)


let__18990 := tmp21413
_ = let__18990

tmp21414 := GCons(__e, symboolean, Nil)

tmp21415 := GCons(__e, sym_1_1_6, tmp21414)

tmp21416 := GCons(__e, let__18990, tmp21415)

tmp21417 := GCons(__e, tmp21416, Nil)

tmp21418 := GCons(__e, sym_1_1_6, tmp21417)

tmp21419 := GCons(__e, let__18990, tmp21418)

tmp21420 := Call6(__e, SymFn(symis_b), V6916, tmp21419, B6912, L6913, Key6914, C6915)


return TailCall2(__e, SymFn(symshen_4gc), B6912, tmp21420)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp21421 := GValue(__e, symshen_4_dsigf_d)

tmp21422 := Call3(__e, SymFn(symshen_4assoc_1_6), sym_a, tmp21412, tmp21421)


tmp21423 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp21422)

_ = tmp21423

tmp21424 := MakeNative1(func(__e *ControlFlow, V6921 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6917 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6918 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6919 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6920 Obj) Obj {
tmp21425 := GCons(__e, symnumber, Nil)

tmp21426 := GCons(__e, sym_1_1_6, tmp21425)

tmp21427 := GCons(__e, symnumber, tmp21426)

tmp21428 := GCons(__e, tmp21427, Nil)

tmp21429 := GCons(__e, sym_1_1_6, tmp21428)

tmp21430 := GCons(__e, symnumber, tmp21429)

return TailCall6(__e, SymFn(symis_b), V6921, tmp21430, B6917, L6918, Key6919, C6920)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp21431 := GValue(__e, symshen_4_dsigf_d)

tmp21432 := Call3(__e, SymFn(symshen_4assoc_1_6), sym_7, tmp21424, tmp21431)


tmp21433 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp21432)

_ = tmp21433

tmp21434 := MakeNative1(func(__e *ControlFlow, V6926 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6922 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6923 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6924 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6925 Obj) Obj {
tmp21435 := GCons(__e, symnumber, Nil)

tmp21436 := GCons(__e, sym_1_1_6, tmp21435)

tmp21437 := GCons(__e, symnumber, tmp21436)

tmp21438 := GCons(__e, tmp21437, Nil)

tmp21439 := GCons(__e, sym_1_1_6, tmp21438)

tmp21440 := GCons(__e, symnumber, tmp21439)

return TailCall6(__e, SymFn(symis_b), V6926, tmp21440, B6922, L6923, Key6924, C6925)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp21441 := GValue(__e, symshen_4_dsigf_d)

tmp21442 := Call3(__e, SymFn(symshen_4assoc_1_6), sym_c, tmp21434, tmp21441)


tmp21443 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp21442)

_ = tmp21443

tmp21444 := MakeNative1(func(__e *ControlFlow, V6931 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6927 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6928 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6929 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6930 Obj) Obj {
tmp21445 := GCons(__e, symnumber, Nil)

tmp21446 := GCons(__e, sym_1_1_6, tmp21445)

tmp21447 := GCons(__e, symnumber, tmp21446)

tmp21448 := GCons(__e, tmp21447, Nil)

tmp21449 := GCons(__e, sym_1_1_6, tmp21448)

tmp21450 := GCons(__e, symnumber, tmp21449)

return TailCall6(__e, SymFn(symis_b), V6931, tmp21450, B6927, L6928, Key6929, C6930)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp21451 := GValue(__e, symshen_4_dsigf_d)

tmp21452 := Call3(__e, SymFn(symshen_4assoc_1_6), sym_1, tmp21444, tmp21451)


tmp21453 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp21452)

_ = tmp21453

tmp21454 := MakeNative1(func(__e *ControlFlow, V6936 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6932 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6933 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6934 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6935 Obj) Obj {
tmp21455 := GCons(__e, symnumber, Nil)

tmp21456 := GCons(__e, sym_1_1_6, tmp21455)

tmp21457 := GCons(__e, symnumber, tmp21456)

tmp21458 := GCons(__e, tmp21457, Nil)

tmp21459 := GCons(__e, sym_1_1_6, tmp21458)

tmp21460 := GCons(__e, symnumber, tmp21459)

return TailCall6(__e, SymFn(symis_b), V6936, tmp21460, B6932, L6933, Key6934, C6935)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp21461 := GValue(__e, symshen_4_dsigf_d)

tmp21462 := Call3(__e, SymFn(symshen_4assoc_1_6), sym_d, tmp21454, tmp21461)


tmp21463 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp21462)

_ = tmp21463

tmp21464 := MakeNative1(func(__e *ControlFlow, V6941 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6937 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6938 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6939 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6940 Obj) Obj {
tmp21465 := Call1(__e, SymFn(symshen_4newpv), B6937)


let__18991 := tmp21465
_ = let__18991

tmp21466 := Call1(__e, SymFn(symshen_4newpv), B6937)


let__18992 := tmp21466
_ = let__18992

tmp21467 := GCons(__e, symboolean, Nil)

tmp21468 := GCons(__e, sym_1_1_6, tmp21467)

tmp21469 := GCons(__e, let__18992, tmp21468)

tmp21470 := GCons(__e, tmp21469, Nil)

tmp21471 := GCons(__e, sym_1_1_6, tmp21470)

tmp21472 := GCons(__e, let__18991, tmp21471)

tmp21473 := Call6(__e, SymFn(symis_b), V6941, tmp21472, B6937, L6938, Key6939, C6940)


tmp21474 := Call2(__e, SymFn(symshen_4gc), B6937, tmp21473)


return TailCall2(__e, SymFn(symshen_4gc), B6937, tmp21474)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp21475 := GValue(__e, symshen_4_dsigf_d)

tmp21476 := Call3(__e, SymFn(symshen_4assoc_1_6), sym_a_a, tmp21464, tmp21475)


return PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp21476)


panic("unreachable")
})

tmp21477 := Call2(__e, ns2_1set, symshen_4initialise_1signedfuncs, tmp19758)


_ = tmp21477

tmp21478 := MakeNative0(func(__e *ControlFlow) Obj {
tmp21479 := MakeNative1(func(__e *ControlFlow, Y954 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4tuple), Y954)
panic("unreachable")
})

tmp21480 := GCons(__e, symshen_4tuple, tmp21479)

tmp21481 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21480)


_ = tmp21481

tmp21482 := MakeNative1(func(__e *ControlFlow, Y955 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4pvar), Y955)
panic("unreachable")
})

tmp21483 := GCons(__e, symshen_4pvar, tmp21482)

tmp21484 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21483)


_ = tmp21484

tmp21485 := MakeNative1(func(__e *ControlFlow, Y956 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4dictionary), Y956)
panic("unreachable")
})

tmp21486 := GCons(__e, symshen_4dictionary, tmp21485)

tmp21487 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21486)


_ = tmp21487

tmp21488 := MakeNative1(func(__e *ControlFlow, Y957 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4print_1prolog_1vector), Y957)
panic("unreachable")
})

tmp21489 := GCons(__e, symshen_4print_1prolog_1vector, tmp21488)

tmp21490 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21489)


_ = tmp21490

tmp21491 := MakeNative1(func(__e *ControlFlow, Y958 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4print_1freshterm), Y958)
panic("unreachable")
})

tmp21492 := GCons(__e, symshen_4print_1freshterm, tmp21491)

tmp21493 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21492)


_ = tmp21493

tmp21494 := MakeNative1(func(__e *ControlFlow, Y959 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4printF), Y959)
panic("unreachable")
})

tmp21495 := GCons(__e, symshen_4printF, tmp21494)

tmp21496 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21495)


_ = tmp21496

tmp21497 := MakeNative1(func(__e *ControlFlow, Y960 Obj) Obj {
return TailCall1(__e, SymFn(symabsolute), Y960)
panic("unreachable")
})

tmp21498 := GCons(__e, symabsolute, tmp21497)

tmp21499 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21498)


_ = tmp21499

tmp21500 := MakeNative1(func(__e *ControlFlow, Y961 Obj) Obj {
return GVectorP(__e, Y961)
panic("unreachable")
})

tmp21501 := GCons(__e, symabsvector_2, tmp21500)

tmp21502 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21501)


_ = tmp21502

tmp21503 := MakeNative1(func(__e *ControlFlow, Y962 Obj) Obj {
return PrimGuard1(__e, symabsvector, PrimAbsvector, Y962)
panic("unreachable")
})

tmp21504 := GCons(__e, symabsvector, tmp21503)

tmp21505 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21504)


_ = tmp21505

tmp21506 := MakeNative1(func(__e *ControlFlow, Y963 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y964 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y965 Obj) Obj {
return GVecSet(__e, Y963, Y964, Y965)
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp21507 := GCons(__e, symaddress_1_6, tmp21506)

tmp21508 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21507)


_ = tmp21508

tmp21509 := MakeNative1(func(__e *ControlFlow, Y966 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y967 Obj) Obj {
return TailCall2(__e, SymFn(symadjoin), Y966, Y967)
panic("unreachable")
})
panic("unreachable")
})

tmp21510 := GCons(__e, symadjoin, tmp21509)

tmp21511 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21510)


_ = tmp21511

tmp21512 := MakeNative1(func(__e *ControlFlow, Y968 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y969 Obj) Obj {
if True == Y968 {
if True == Y969 {
return True
} else {
return False
}
} else {
return False
}
panic("unreachable")
})
panic("unreachable")
})

tmp21515 := GCons(__e, symand, tmp21512)

tmp21516 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21515)


_ = tmp21516

tmp21517 := MakeNative1(func(__e *ControlFlow, Y970 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y971 Obj) Obj {
return TailCall2(__e, SymFn(symappend), Y970, Y971)
panic("unreachable")
})
panic("unreachable")
})

tmp21518 := GCons(__e, symappend, tmp21517)

tmp21519 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21518)


_ = tmp21519

tmp21520 := MakeNative1(func(__e *ControlFlow, Y972 Obj) Obj {
return TailCall1(__e, SymFn(symarity), Y972)
panic("unreachable")
})

tmp21521 := GCons(__e, symarity, tmp21520)

tmp21522 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21521)


_ = tmp21522

tmp21523 := MakeNative1(func(__e *ControlFlow, Y973 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y974 Obj) Obj {
return TailCall2(__e, SymFn(symassoc), Y973, Y974)
panic("unreachable")
})
panic("unreachable")
})

tmp21524 := GCons(__e, symassoc, tmp21523)

tmp21525 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21524)


_ = tmp21525

tmp21526 := MakeNative1(func(__e *ControlFlow, Y975 Obj) Obj {
return TailCall1(__e, SymFn(symatom_2), Y975)
panic("unreachable")
})

tmp21527 := GCons(__e, symatom_2, tmp21526)

tmp21528 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21527)


_ = tmp21528

tmp21529 := MakeNative1(func(__e *ControlFlow, Y976 Obj) Obj {
return TailCall1(__e, SymFn(symboolean_2), Y976)
panic("unreachable")
})

tmp21530 := GCons(__e, symboolean_2, tmp21529)

tmp21531 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21530)


_ = tmp21531

tmp21532 := MakeNative1(func(__e *ControlFlow, Y977 Obj) Obj {
return TailCall1(__e, SymFn(symbootstrap), Y977)
panic("unreachable")
})

tmp21533 := GCons(__e, symbootstrap, tmp21532)

tmp21534 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21533)


_ = tmp21534

tmp21535 := MakeNative1(func(__e *ControlFlow, Y978 Obj) Obj {
return TailCall1(__e, SymFn(symbound_2), Y978)
panic("unreachable")
})

tmp21536 := GCons(__e, symbound_2, tmp21535)

tmp21537 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21536)


_ = tmp21537

tmp21538 := MakeNative1(func(__e *ControlFlow, Y979 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y980 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y981 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y982 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y983 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y984 Obj) Obj {
return TailCall6(__e, SymFn(symbind), Y979, Y980, Y981, Y982, Y983, Y984)
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp21539 := GCons(__e, symbind, tmp21538)

tmp21540 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21539)


_ = tmp21540

tmp21541 := MakeNative1(func(__e *ControlFlow, Y985 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y986 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y987 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y988 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y989 Obj) Obj {
return TailCall5(__e, SymFn(symcall), Y985, Y986, Y987, Y988, Y989)
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp21542 := GCons(__e, symcall, tmp21541)

tmp21543 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21542)


_ = tmp21543

tmp21544 := MakeNative1(func(__e *ControlFlow, Y990 Obj) Obj {
return TailCall1(__e, SymFn(symcd), Y990)
panic("unreachable")
})

tmp21545 := GCons(__e, symcd, tmp21544)

tmp21546 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21545)


_ = tmp21546

tmp21547 := MakeNative1(func(__e *ControlFlow, Y991 Obj) Obj {
return TailCall1(__e, SymFn(symchar_1stream_2), Y991)
panic("unreachable")
})

tmp21548 := GCons(__e, symchar_1stream_2, tmp21547)

tmp21549 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21548)


_ = tmp21549

tmp21550 := MakeNative1(func(__e *ControlFlow, Y992 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y993 Obj) Obj {
return TailCall2(__e, SymFn(symcompile), Y992, Y993)
panic("unreachable")
})
panic("unreachable")
})

tmp21551 := GCons(__e, symcompile, tmp21550)

tmp21552 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21551)


_ = tmp21552

tmp21553 := MakeNative1(func(__e *ControlFlow, Y994 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y995 Obj) Obj {
return TailCall2(__e, SymFn(symconcat), Y994, Y995)
panic("unreachable")
})
panic("unreachable")
})

tmp21554 := GCons(__e, symconcat, tmp21553)

tmp21555 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21554)


_ = tmp21555

tmp21556 := MakeNative1(func(__e *ControlFlow, Y996 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y997 Obj) Obj {
return GCons(__e, Y996, Y997)
panic("unreachable")
})
panic("unreachable")
})

tmp21557 := GCons(__e, symcons, tmp21556)

tmp21558 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21557)


_ = tmp21558

tmp21559 := MakeNative1(func(__e *ControlFlow, Y998 Obj) Obj {
return GConsP(__e, Y998)
panic("unreachable")
})

tmp21560 := GCons(__e, symcons_2, tmp21559)

tmp21561 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21560)


_ = tmp21561

tmp21562 := MakeNative1(func(__e *ControlFlow, Y999 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1000 Obj) Obj {
return PrimGuard2(__e, symcn, PrimStringConcat, Y999, Y1000)
panic("unreachable")
})
panic("unreachable")
})

tmp21563 := GCons(__e, symcn, tmp21562)

tmp21564 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21563)


_ = tmp21564

tmp21565 := MakeNative1(func(__e *ControlFlow, Y1001 Obj) Obj {
return PrimGuard1(__e, symclose, PrimCloseStream, Y1001)
panic("unreachable")
})

tmp21566 := GCons(__e, symclose, tmp21565)

tmp21567 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21566)


_ = tmp21567

tmp21568 := MakeNative1(func(__e *ControlFlow, Y1002 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1003 Obj) Obj {
return TailCall2(__e, SymFn(symdeclare), Y1002, Y1003)
panic("unreachable")
})
panic("unreachable")
})

tmp21569 := GCons(__e, symdeclare, tmp21568)

tmp21570 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21569)


_ = tmp21570

tmp21571 := MakeNative1(func(__e *ControlFlow, Y1004 Obj) Obj {
return TailCall1(__e, SymFn(symdestroy), Y1004)
panic("unreachable")
})

tmp21572 := GCons(__e, symdestroy, tmp21571)

tmp21573 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21572)


_ = tmp21573

tmp21574 := MakeNative1(func(__e *ControlFlow, Y1005 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1006 Obj) Obj {
return TailCall2(__e, SymFn(symdifference), Y1005, Y1006)
panic("unreachable")
})
panic("unreachable")
})

tmp21575 := GCons(__e, symdifference, tmp21574)

tmp21576 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21575)


_ = tmp21576

tmp21577 := MakeNative1(func(__e *ControlFlow, Y1007 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1008 Obj) Obj {
_ = Y1007

return Y1008


panic("unreachable")
})
panic("unreachable")
})

tmp21578 := GCons(__e, symdo, tmp21577)

tmp21579 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21578)


_ = tmp21579

tmp21580 := MakeNative1(func(__e *ControlFlow, Y1009 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1010 Obj) Obj {
return TailCall2(__e, SymFn(symelement_2), Y1009, Y1010)
panic("unreachable")
})
panic("unreachable")
})

tmp21581 := GCons(__e, symelement_2, tmp21580)

tmp21582 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21581)


_ = tmp21582

tmp21583 := MakeNative1(func(__e *ControlFlow, Y1011 Obj) Obj {
return TailCall1(__e, SymFn(symempty_2), Y1011)
panic("unreachable")
})

tmp21584 := GCons(__e, symempty_2, tmp21583)

tmp21585 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21584)


_ = tmp21585

tmp21586 := MakeNative1(func(__e *ControlFlow, Y1012 Obj) Obj {
return TailCall1(__e, SymFn(symenable_1type_1theory), Y1012)
panic("unreachable")
})

tmp21587 := GCons(__e, symenable_1type_1theory, tmp21586)

tmp21588 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21587)


_ = tmp21588

tmp21589 := MakeNative1(func(__e *ControlFlow, Y1013 Obj) Obj {
return TailCall1(__e, SymFn(symexternal), Y1013)
panic("unreachable")
})

tmp21590 := GCons(__e, symexternal, tmp21589)

tmp21591 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21590)


_ = tmp21591

tmp21592 := MakeNative1(func(__e *ControlFlow, Y1014 Obj) Obj {
return PrimGuard1(__e, symerror_1to_1string, PrimErrorToString, Y1014)
panic("unreachable")
})

tmp21593 := GCons(__e, symerror_1to_1string, tmp21592)

tmp21594 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21593)


_ = tmp21594

tmp21595 := MakeNative1(func(__e *ControlFlow, Y1015 Obj) Obj {
return TailCall1(__e, SymFn(symeval), Y1015)
panic("unreachable")
})

tmp21596 := GCons(__e, symeval, tmp21595)

tmp21597 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21596)


_ = tmp21597

tmp21598 := MakeNative1(func(__e *ControlFlow, Y1016 Obj) Obj {
return TailCall1(__e, SymFn(symeval_1kl), Y1016)
panic("unreachable")
})

tmp21599 := GCons(__e, symeval_1kl, tmp21598)

tmp21600 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21599)


_ = tmp21600

tmp21601 := MakeNative1(func(__e *ControlFlow, Y1017 Obj) Obj {
return TailCall1(__e, SymFn(symexplode), Y1017)
panic("unreachable")
})

tmp21602 := GCons(__e, symexplode, tmp21601)

tmp21603 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21602)


_ = tmp21603

tmp21604 := MakeNative1(func(__e *ControlFlow, Y1018 Obj) Obj {
return TailCall1(__e, SymFn(symexternal), Y1018)
panic("unreachable")
})

tmp21605 := GCons(__e, symexternal, tmp21604)

tmp21606 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21605)


_ = tmp21606

tmp21607 := MakeNative1(func(__e *ControlFlow, Y1019 Obj) Obj {
return TailCall1(__e, SymFn(symfactorise), Y1019)
panic("unreachable")
})

tmp21608 := GCons(__e, symfactorise, tmp21607)

tmp21609 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21608)


_ = tmp21609

tmp21610 := MakeNative1(func(__e *ControlFlow, Y1020 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1021 Obj) Obj {
return TailCall2(__e, SymFn(symfail_1if), Y1020, Y1021)
panic("unreachable")
})
panic("unreachable")
})

tmp21611 := GCons(__e, symfail_1if, tmp21610)

tmp21612 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21611)


_ = tmp21612

tmp21613 := MakeNative1(func(__e *ControlFlow, Y1022 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1023 Obj) Obj {
return TailCall2(__e, SymFn(symfix), Y1022, Y1023)
panic("unreachable")
})
panic("unreachable")
})

tmp21614 := GCons(__e, symfix, tmp21613)

tmp21615 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21614)


_ = tmp21615

tmp21616 := MakeNative1(func(__e *ControlFlow, Y1024 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1025 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1026 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1027 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1028 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1029 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1030 Obj) Obj {
__e.TailApply(SymFn(symfindall), Y1024, Y1025, Y1026, Y1027, Y1028, Y1029, Y1030)
return nil
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp21617 := GCons(__e, symfindall, tmp21616)

tmp21618 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21617)


_ = tmp21618

tmp21619 := MakeNative1(func(__e *ControlFlow, Y1031 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1032 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1033 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1034 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1035 Obj) Obj {
return TailCall5(__e, SymFn(symfork), Y1031, Y1032, Y1033, Y1034, Y1035)
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp21620 := GCons(__e, symfork, tmp21619)

tmp21621 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21620)


_ = tmp21621

tmp21622 := MakeNative1(func(__e *ControlFlow, Y1036 Obj) Obj {
return MakeNative0(func(__e *ControlFlow) Obj {
return Y1036
panic("unreachable")
})
panic("unreachable")
})

tmp21623 := GCons(__e, symfreeze, tmp21622)

tmp21624 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21623)


_ = tmp21624

tmp21625 := MakeNative1(func(__e *ControlFlow, Y1037 Obj) Obj {
return TailCall1(__e, SymFn(symfst), Y1037)
panic("unreachable")
})

tmp21626 := GCons(__e, symfst, tmp21625)

tmp21627 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21626)


_ = tmp21627

tmp21628 := MakeNative1(func(__e *ControlFlow, Y1038 Obj) Obj {
return TailCall1(__e, SymFn(symfn), Y1038)
panic("unreachable")
})

tmp21629 := GCons(__e, symfn, tmp21628)

tmp21630 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21629)


_ = tmp21630

tmp21631 := MakeNative1(func(__e *ControlFlow, Y1039 Obj) Obj {
return TailCall1(__e, SymFn(symfunction), Y1039)
panic("unreachable")
})

tmp21632 := GCons(__e, symfunction, tmp21631)

tmp21633 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21632)


_ = tmp21633

tmp21634 := MakeNative1(func(__e *ControlFlow, Y1040 Obj) Obj {
return TailCall1(__e, SymFn(symgensym), Y1040)
panic("unreachable")
})

tmp21635 := GCons(__e, symgensym, tmp21634)

tmp21636 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21635)


_ = tmp21636

tmp21637 := MakeNative1(func(__e *ControlFlow, Y1041 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1042 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1043 Obj) Obj {
return TailCall3(__e, SymFn(symget), Y1041, Y1042, Y1043)
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp21638 := GCons(__e, symget, tmp21637)

tmp21639 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21638)


_ = tmp21639

tmp21640 := MakeNative1(func(__e *ControlFlow, Y1044 Obj) Obj {
return PrimGuard1(__e, symget_1time, PrimGetTime, Y1044)
panic("unreachable")
})

tmp21641 := GCons(__e, symget_1time, tmp21640)

tmp21642 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21641)


_ = tmp21642

tmp21643 := MakeNative1(func(__e *ControlFlow, Y1045 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1046 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1047 Obj) Obj {
return GVecSet(__e, Y1045, Y1046, Y1047)
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp21644 := GCons(__e, symaddress_1_6, tmp21643)

tmp21645 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21644)


_ = tmp21645

tmp21646 := MakeNative1(func(__e *ControlFlow, Y1048 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1049 Obj) Obj {
return GVecGet(__e, Y1048, Y1049)
panic("unreachable")
})
panic("unreachable")
})

tmp21647 := GCons(__e, sym_5_1address, tmp21646)

tmp21648 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21647)


_ = tmp21648

tmp21649 := MakeNative1(func(__e *ControlFlow, Y1050 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1051 Obj) Obj {
return TailCall2(__e, SymFn(sym_5_1vector), Y1050, Y1051)
panic("unreachable")
})
panic("unreachable")
})

tmp21650 := GCons(__e, sym_5_1vector, tmp21649)

tmp21651 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21650)


_ = tmp21651

tmp21652 := MakeNative1(func(__e *ControlFlow, Y1052 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1053 Obj) Obj {
return GGt(__e, Y1052, Y1053)
panic("unreachable")
})
panic("unreachable")
})

tmp21653 := GCons(__e, sym_6, tmp21652)

tmp21654 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21653)


_ = tmp21654

tmp21655 := MakeNative1(func(__e *ControlFlow, Y1054 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1055 Obj) Obj {
return GGe(__e, Y1054, Y1055)
panic("unreachable")
})
panic("unreachable")
})

tmp21656 := GCons(__e, sym_6_a, tmp21655)

tmp21657 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21656)


_ = tmp21657

tmp21658 := MakeNative1(func(__e *ControlFlow, Y1056 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1057 Obj) Obj {
return GEq(__e, Y1056, Y1057)
panic("unreachable")
})
panic("unreachable")
})

tmp21659 := GCons(__e, sym_a, tmp21658)

tmp21660 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21659)


_ = tmp21660

tmp21661 := MakeNative1(func(__e *ControlFlow, Y1058 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1059 Obj) Obj {
return TailCall2(__e, SymFn(symhash), Y1058, Y1059)
panic("unreachable")
})
panic("unreachable")
})

tmp21662 := GCons(__e, symhash, tmp21661)

tmp21663 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21662)


_ = tmp21663

tmp21664 := MakeNative1(func(__e *ControlFlow, Y1060 Obj) Obj {
return GHd(__e, Y1060)
panic("unreachable")
})

tmp21665 := GCons(__e, symhd, tmp21664)

tmp21666 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21665)


_ = tmp21666

tmp21667 := MakeNative1(func(__e *ControlFlow, Y1061 Obj) Obj {
return TailCall1(__e, SymFn(symhdv), Y1061)
panic("unreachable")
})

tmp21668 := GCons(__e, symhdv, tmp21667)

tmp21669 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21668)


_ = tmp21669

tmp21670 := MakeNative1(func(__e *ControlFlow, Y1062 Obj) Obj {
return TailCall1(__e, SymFn(symhdstr), Y1062)
panic("unreachable")
})

tmp21671 := GCons(__e, symhdstr, tmp21670)

tmp21672 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21671)


_ = tmp21672

tmp21673 := MakeNative1(func(__e *ControlFlow, Y1063 Obj) Obj {
return TailCall1(__e, SymFn(symhead), Y1063)
panic("unreachable")
})

tmp21674 := GCons(__e, symhead, tmp21673)

tmp21675 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21674)


_ = tmp21675

tmp21676 := MakeNative1(func(__e *ControlFlow, Y1064 Obj) Obj {
return TailCall1(__e, SymFn(symhush), Y1064)
panic("unreachable")
})

tmp21677 := GCons(__e, symhush, tmp21676)

tmp21678 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21677)


_ = tmp21678

tmp21679 := MakeNative1(func(__e *ControlFlow, Y1065 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1066 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1067 Obj) Obj {
if True == Y1065 {
return Y1066
} else {
return Y1067
}
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp21681 := GCons(__e, symif, tmp21679)

tmp21682 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21681)


_ = tmp21682

tmp21683 := MakeNative1(func(__e *ControlFlow, Y1068 Obj) Obj {
return TailCall1(__e, SymFn(syminclude), Y1068)
panic("unreachable")
})

tmp21684 := GCons(__e, syminclude, tmp21683)

tmp21685 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21684)


_ = tmp21685

tmp21686 := MakeNative1(func(__e *ControlFlow, Y1069 Obj) Obj {
return TailCall1(__e, SymFn(symin_1package), Y1069)
panic("unreachable")
})

tmp21687 := GCons(__e, symin_1package, tmp21686)

tmp21688 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21687)


_ = tmp21688

tmp21689 := MakeNative1(func(__e *ControlFlow, Y1070 Obj) Obj {
return PrimGuard1(__e, syminteger_2, PrimIsInteger, Y1070)
panic("unreachable")
})

tmp21690 := GCons(__e, syminteger_2, tmp21689)

tmp21691 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21690)


_ = tmp21691

tmp21692 := MakeNative1(func(__e *ControlFlow, Y1071 Obj) Obj {
return TailCall1(__e, SymFn(syminternal), Y1071)
panic("unreachable")
})

tmp21693 := GCons(__e, syminternal, tmp21692)

tmp21694 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21693)


_ = tmp21694

tmp21695 := MakeNative1(func(__e *ControlFlow, Y1072 Obj) Obj {
return PrimGuard1(__e, symintern, PrimIntern, Y1072)
panic("unreachable")
})

tmp21696 := GCons(__e, symintern, tmp21695)

tmp21697 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21696)


_ = tmp21697

tmp21698 := MakeNative1(func(__e *ControlFlow, Y1073 Obj) Obj {
return TailCall1(__e, SymFn(syminput), Y1073)
panic("unreachable")
})

tmp21699 := GCons(__e, syminput, tmp21698)

tmp21700 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21699)


_ = tmp21700

tmp21701 := MakeNative1(func(__e *ControlFlow, Y1074 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1075 Obj) Obj {
return TailCall2(__e, SymFn(symshen_4input_1h_7), Y1074, Y1075)
panic("unreachable")
})
panic("unreachable")
})

tmp21702 := GCons(__e, syminput_7, tmp21701)

tmp21703 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21702)


_ = tmp21703

tmp21704 := MakeNative1(func(__e *ControlFlow, Y1076 Obj) Obj {
return TailCall1(__e, SymFn(syminclude_1all_1but), Y1076)
panic("unreachable")
})

tmp21705 := GCons(__e, syminclude_1all_1but, tmp21704)

tmp21706 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21705)


_ = tmp21706

tmp21707 := MakeNative1(func(__e *ControlFlow, Y1077 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1078 Obj) Obj {
return TailCall2(__e, SymFn(symintersection), Y1077, Y1078)
panic("unreachable")
})
panic("unreachable")
})

tmp21708 := GCons(__e, symintersection, tmp21707)

tmp21709 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21708)


_ = tmp21709

tmp21710 := MakeNative1(func(__e *ControlFlow, Y1079 Obj) Obj {
return TailCall1(__e, SymFn(syminternal), Y1079)
panic("unreachable")
})

tmp21711 := GCons(__e, syminternal, tmp21710)

tmp21712 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21711)


_ = tmp21712

tmp21713 := MakeNative1(func(__e *ControlFlow, Y1080 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1081 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1082 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1083 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1084 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1085 Obj) Obj {
return TailCall6(__e, SymFn(symis), Y1080, Y1081, Y1082, Y1083, Y1084, Y1085)
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp21714 := GCons(__e, symis, tmp21713)

tmp21715 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21714)


_ = tmp21715

tmp21716 := MakeNative1(func(__e *ControlFlow, Y1086 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1087 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1088 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1089 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1090 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1091 Obj) Obj {
return TailCall6(__e, SymFn(symis_b), Y1086, Y1087, Y1088, Y1089, Y1090, Y1091)
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp21717 := GCons(__e, symis_b, tmp21716)

tmp21718 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21717)


_ = tmp21718

tmp21719 := MakeNative1(func(__e *ControlFlow, Y1092 Obj) Obj {
return TailCall1(__e, SymFn(symlength), Y1092)
panic("unreachable")
})

tmp21720 := GCons(__e, symlength, tmp21719)

tmp21721 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21720)


_ = tmp21721

tmp21722 := MakeNative1(func(__e *ControlFlow, Y1093 Obj) Obj {
return TailCall1(__e, SymFn(symlimit), Y1093)
panic("unreachable")
})

tmp21723 := GCons(__e, symlimit, tmp21722)

tmp21724 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21723)


_ = tmp21724

tmp21725 := MakeNative1(func(__e *ControlFlow, Y1094 Obj) Obj {
return TailCall1(__e, SymFn(symlineread), Y1094)
panic("unreachable")
})

tmp21726 := GCons(__e, symlineread, tmp21725)

tmp21727 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21726)


_ = tmp21727

tmp21728 := MakeNative1(func(__e *ControlFlow, Y1095 Obj) Obj {
return TailCall1(__e, SymFn(symload), Y1095)
panic("unreachable")
})

tmp21729 := GCons(__e, symload, tmp21728)

tmp21730 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21729)


_ = tmp21730

tmp21731 := MakeNative1(func(__e *ControlFlow, Y1096 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1097 Obj) Obj {
return GLt(__e, Y1096, Y1097)
panic("unreachable")
})
panic("unreachable")
})

tmp21732 := GCons(__e, sym_5, tmp21731)

tmp21733 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21732)


_ = tmp21733

tmp21734 := MakeNative1(func(__e *ControlFlow, Y1098 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1099 Obj) Obj {
return GLe(__e, Y1098, Y1099)
panic("unreachable")
})
panic("unreachable")
})

tmp21735 := GCons(__e, sym_5_a, tmp21734)

tmp21736 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21735)


_ = tmp21736

tmp21737 := MakeNative1(func(__e *ControlFlow, Y1100 Obj) Obj {
return TailCall1(__e, SymFn(symvector), Y1100)
panic("unreachable")
})

tmp21738 := GCons(__e, symvector, tmp21737)

tmp21739 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21738)


_ = tmp21739

tmp21740 := MakeNative1(func(__e *ControlFlow, Y1101 Obj) Obj {
return TailCall1(__e, SymFn(symmacroexpand), Y1101)
panic("unreachable")
})

tmp21741 := GCons(__e, symmacroexpand, tmp21740)

tmp21742 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21741)


_ = tmp21742

tmp21743 := MakeNative1(func(__e *ControlFlow, Y1102 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1103 Obj) Obj {
return TailCall2(__e, SymFn(symmap), Y1102, Y1103)
panic("unreachable")
})
panic("unreachable")
})

tmp21744 := GCons(__e, symmap, tmp21743)

tmp21745 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21744)


_ = tmp21745

tmp21746 := MakeNative1(func(__e *ControlFlow, Y1104 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1105 Obj) Obj {
return TailCall2(__e, SymFn(symmapcan), Y1104, Y1105)
panic("unreachable")
})
panic("unreachable")
})

tmp21747 := GCons(__e, symmapcan, tmp21746)

tmp21748 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21747)


_ = tmp21748

tmp21749 := MakeNative1(func(__e *ControlFlow, Y1106 Obj) Obj {
return TailCall1(__e, SymFn(symmaxinferences), Y1106)
panic("unreachable")
})

tmp21750 := GCons(__e, symmaxinferences, tmp21749)

tmp21751 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21750)


_ = tmp21751

tmp21752 := MakeNative1(func(__e *ControlFlow, Y1107 Obj) Obj {
return TailCall1(__e, SymFn(symnl), Y1107)
panic("unreachable")
})

tmp21753 := GCons(__e, symnl, tmp21752)

tmp21754 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21753)


_ = tmp21754

tmp21755 := MakeNative1(func(__e *ControlFlow, Y1108 Obj) Obj {
return GNot(__e, Y1108)
panic("unreachable")
})

tmp21756 := GCons(__e, symnot, tmp21755)

tmp21757 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21756)


_ = tmp21757

tmp21758 := MakeNative1(func(__e *ControlFlow, Y1109 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1110 Obj) Obj {
return TailCall2(__e, SymFn(symnth), Y1109, Y1110)
panic("unreachable")
})
panic("unreachable")
})

tmp21759 := GCons(__e, symnth, tmp21758)

tmp21760 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21759)


_ = tmp21760

tmp21761 := MakeNative1(func(__e *ControlFlow, Y1111 Obj) Obj {
return PrimGuard1(__e, symn_1_6string, PrimNumberToString, Y1111)
panic("unreachable")
})

tmp21762 := GCons(__e, symn_1_6string, tmp21761)

tmp21763 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21762)


_ = tmp21763

tmp21764 := MakeNative1(func(__e *ControlFlow, Y1112 Obj) Obj {
return PrimGuard1(__e, symnumber_2, PrimIsNumber, Y1112)
panic("unreachable")
})

tmp21765 := GCons(__e, symnumber_2, tmp21764)

tmp21766 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21765)


_ = tmp21766

tmp21767 := MakeNative1(func(__e *ControlFlow, Y1113 Obj) Obj {
return TailCall1(__e, SymFn(symoccurs_1check), Y1113)
panic("unreachable")
})

tmp21768 := GCons(__e, symoccurs_1check, tmp21767)

tmp21769 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21768)


_ = tmp21769

tmp21770 := MakeNative1(func(__e *ControlFlow, Y1114 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1115 Obj) Obj {
return TailCall2(__e, SymFn(symoccurrences), Y1114, Y1115)
panic("unreachable")
})
panic("unreachable")
})

tmp21771 := GCons(__e, symoccurrences, tmp21770)

tmp21772 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21771)


_ = tmp21772

tmp21773 := MakeNative1(func(__e *ControlFlow, Y1116 Obj) Obj {
return TailCall1(__e, SymFn(symoccurs_1check), Y1116)
panic("unreachable")
})

tmp21774 := GCons(__e, symoccurs_1check, tmp21773)

tmp21775 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21774)


_ = tmp21775

tmp21776 := MakeNative1(func(__e *ControlFlow, Y1117 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1118 Obj) Obj {
return PrimGuard2(__e, symopen, PrimOpenStream, Y1117, Y1118)
panic("unreachable")
})
panic("unreachable")
})

tmp21777 := GCons(__e, symopen, tmp21776)

tmp21778 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21777)


_ = tmp21778

tmp21779 := MakeNative1(func(__e *ControlFlow, Y1119 Obj) Obj {
return TailCall1(__e, SymFn(symoptimise), Y1119)
panic("unreachable")
})

tmp21780 := GCons(__e, symoptimise, tmp21779)

tmp21781 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21780)


_ = tmp21781

tmp21782 := MakeNative1(func(__e *ControlFlow, Y1120 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1121 Obj) Obj {
if True == Y1120 {
return True
} else {
if True == Y1121 {
return True
} else {
return False
}
}
panic("unreachable")
})
panic("unreachable")
})

tmp21785 := GCons(__e, symor, tmp21782)

tmp21786 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21785)


_ = tmp21786

tmp21787 := MakeNative1(func(__e *ControlFlow, Y1122 Obj) Obj {
return TailCall1(__e, SymFn(sympackage_2), Y1122)
panic("unreachable")
})

tmp21788 := GCons(__e, sympackage_2, tmp21787)

tmp21789 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21788)


_ = tmp21789

tmp21790 := MakeNative1(func(__e *ControlFlow, Y1123 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1124 Obj) Obj {
return PrimGuard2(__e, sympos, PrimPos, Y1123, Y1124)
panic("unreachable")
})
panic("unreachable")
})

tmp21791 := GCons(__e, sympos, tmp21790)

tmp21792 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21791)


_ = tmp21792

tmp21793 := MakeNative1(func(__e *ControlFlow, Y1125 Obj) Obj {
return TailCall1(__e, SymFn(sympreclude_1all_1but), Y1125)
panic("unreachable")
})

tmp21794 := GCons(__e, sympreclude_1all_1but, tmp21793)

tmp21795 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21794)


_ = tmp21795

tmp21796 := MakeNative1(func(__e *ControlFlow, Y1126 Obj) Obj {
return TailCall1(__e, SymFn(symprint), Y1126)
panic("unreachable")
})

tmp21797 := GCons(__e, symprint, tmp21796)

tmp21798 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21797)


_ = tmp21798

tmp21799 := MakeNative1(func(__e *ControlFlow, Y1127 Obj) Obj {
return TailCall1(__e, SymFn(symprofile), Y1127)
panic("unreachable")
})

tmp21800 := GCons(__e, symprofile, tmp21799)

tmp21801 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21800)


_ = tmp21801

tmp21802 := MakeNative1(func(__e *ControlFlow, Y1128 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4print_1prolog_1vector), Y1128)
panic("unreachable")
})

tmp21803 := GCons(__e, symshen_4print_1prolog_1vector, tmp21802)

tmp21804 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21803)


_ = tmp21804

tmp21805 := MakeNative1(func(__e *ControlFlow, Y1129 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4print_1freshterm), Y1129)
panic("unreachable")
})

tmp21806 := GCons(__e, symshen_4print_1freshterm, tmp21805)

tmp21807 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21806)


_ = tmp21807

tmp21808 := MakeNative1(func(__e *ControlFlow, Y1130 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4printF), Y1130)
panic("unreachable")
})

tmp21809 := GCons(__e, symshen_4printF, tmp21808)

tmp21810 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21809)


_ = tmp21810

tmp21811 := MakeNative1(func(__e *ControlFlow, Y1131 Obj) Obj {
return TailCall1(__e, SymFn(symprolog_1memory), Y1131)
panic("unreachable")
})

tmp21812 := GCons(__e, symprolog_1memory, tmp21811)

tmp21813 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21812)


_ = tmp21813

tmp21814 := MakeNative1(func(__e *ControlFlow, Y1132 Obj) Obj {
return TailCall1(__e, SymFn(symprofile_1results), Y1132)
panic("unreachable")
})

tmp21815 := GCons(__e, symprofile_1results, tmp21814)

tmp21816 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21815)


_ = tmp21816

tmp21817 := MakeNative1(func(__e *ControlFlow, Y1133 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1134 Obj) Obj {
return TailCall2(__e, SymFn(sympr), Y1133, Y1134)
panic("unreachable")
})
panic("unreachable")
})

tmp21818 := GCons(__e, sympr, tmp21817)

tmp21819 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21818)


_ = tmp21819

tmp21820 := MakeNative1(func(__e *ControlFlow, Y1135 Obj) Obj {
return TailCall1(__e, SymFn(symps), Y1135)
panic("unreachable")
})

tmp21821 := GCons(__e, symps, tmp21820)

tmp21822 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21821)


_ = tmp21822

tmp21823 := MakeNative1(func(__e *ControlFlow, Y1136 Obj) Obj {
return TailCall1(__e, SymFn(sympreclude), Y1136)
panic("unreachable")
})

tmp21824 := GCons(__e, sympreclude, tmp21823)

tmp21825 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21824)


_ = tmp21825

tmp21826 := MakeNative1(func(__e *ControlFlow, Y1137 Obj) Obj {
return TailCall1(__e, SymFn(sympreclude_1all_1but), Y1137)
panic("unreachable")
})

tmp21827 := GCons(__e, sympreclude_1all_1but, tmp21826)

tmp21828 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21827)


_ = tmp21828

tmp21829 := MakeNative1(func(__e *ControlFlow, Y1138 Obj) Obj {
return TailCall1(__e, SymFn(symprotect), Y1138)
panic("unreachable")
})

tmp21830 := GCons(__e, symprotect, tmp21829)

tmp21831 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21830)


_ = tmp21831

tmp21832 := MakeNative1(func(__e *ControlFlow, Y1139 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1140 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1141 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1142 Obj) Obj {
return TailCall4(__e, SymFn(symput), Y1139, Y1140, Y1141, Y1142)
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp21833 := GCons(__e, symput, tmp21832)

tmp21834 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21833)


_ = tmp21834

tmp21835 := MakeNative1(func(__e *ControlFlow, Y1143 Obj) Obj {
return PrimGuard1(__e, symread_1file_1as_1string, PrimReadFileAsString, Y1143)
panic("unreachable")
})

tmp21836 := GCons(__e, symread_1file_1as_1string, tmp21835)

tmp21837 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21836)


_ = tmp21837

tmp21838 := MakeNative1(func(__e *ControlFlow, Y1144 Obj) Obj {
return PrimGuard1(__e, symread_1file_1as_1bytelist, PrimReadFileAsByteList, Y1144)
panic("unreachable")
})

tmp21839 := GCons(__e, symread_1file_1as_1bytelist, tmp21838)

tmp21840 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21839)


_ = tmp21840

tmp21841 := MakeNative1(func(__e *ControlFlow, Y1145 Obj) Obj {
return TailCall1(__e, SymFn(symread_1file), Y1145)
panic("unreachable")
})

tmp21842 := GCons(__e, symread_1file, tmp21841)

tmp21843 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21842)


_ = tmp21843

tmp21844 := MakeNative1(func(__e *ControlFlow, Y1146 Obj) Obj {
return TailCall1(__e, SymFn(symread), Y1146)
panic("unreachable")
})

tmp21845 := GCons(__e, symread, tmp21844)

tmp21846 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21845)


_ = tmp21846

tmp21847 := MakeNative1(func(__e *ControlFlow, Y1147 Obj) Obj {
return PrimGuard1(__e, symread_1byte, PrimReadByte, Y1147)
panic("unreachable")
})

tmp21848 := GCons(__e, symread_1byte, tmp21847)

tmp21849 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21848)


_ = tmp21849

tmp21850 := MakeNative1(func(__e *ControlFlow, Y1148 Obj) Obj {
return TailCall1(__e, SymFn(symread_1from_1string), Y1148)
panic("unreachable")
})

tmp21851 := GCons(__e, symread_1from_1string, tmp21850)

tmp21852 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21851)


_ = tmp21852

tmp21853 := MakeNative1(func(__e *ControlFlow, Y1149 Obj) Obj {
return TailCall1(__e, SymFn(symread_1from_1string_1unprocessed), Y1149)
panic("unreachable")
})

tmp21854 := GCons(__e, symread_1from_1string_1unprocessed, tmp21853)

tmp21855 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21854)


_ = tmp21855

tmp21856 := MakeNative1(func(__e *ControlFlow, Y1150 Obj) Obj {
return TailCall1(__e, SymFn(symread_1unit_1string), Y1150)
panic("unreachable")
})

tmp21857 := GCons(__e, symread_1unit_1string, tmp21856)

tmp21858 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21857)


_ = tmp21858

tmp21859 := MakeNative1(func(__e *ControlFlow, Y1151 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1152 Obj) Obj {
return TailCall2(__e, SymFn(symshen_4register_1source_1form), Y1151, Y1152)
panic("unreachable")
})
panic("unreachable")
})

tmp21860 := GCons(__e, symshen_4register_1source_1form, tmp21859)

tmp21861 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21860)


_ = tmp21861

tmp21862 := MakeNative1(func(__e *ControlFlow, Y1153 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4unregister_1source_1form), Y1153)
panic("unreachable")
})

tmp21863 := GCons(__e, symshen_4unregister_1source_1form, tmp21862)

tmp21864 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21863)


_ = tmp21864

tmp21865 := MakeNative1(func(__e *ControlFlow, Y1154 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1155 Obj) Obj {
return TailCall2(__e, SymFn(symremove), Y1154, Y1155)
panic("unreachable")
})
panic("unreachable")
})

tmp21866 := GCons(__e, symremove, tmp21865)

tmp21867 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21866)


_ = tmp21867

tmp21868 := MakeNative1(func(__e *ControlFlow, Y1156 Obj) Obj {
return TailCall1(__e, SymFn(symreverse), Y1156)
panic("unreachable")
})

tmp21869 := GCons(__e, symreverse, tmp21868)

tmp21870 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21869)


_ = tmp21870

tmp21871 := MakeNative1(func(__e *ControlFlow, Y1157 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1158 Obj) Obj {
return PrimGuard2(__e, symset, PrimSet, Y1157, Y1158)
panic("unreachable")
})
panic("unreachable")
})

tmp21872 := GCons(__e, symset, tmp21871)

tmp21873 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21872)


_ = tmp21873

tmp21874 := MakeNative1(func(__e *ControlFlow, Y1159 Obj) Obj {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, Y1159)
panic("unreachable")
})

tmp21875 := GCons(__e, symsimple_1error, tmp21874)

tmp21876 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21875)


_ = tmp21876

tmp21877 := MakeNative1(func(__e *ControlFlow, Y1160 Obj) Obj {
return TailCall1(__e, SymFn(symsnd), Y1160)
panic("unreachable")
})

tmp21878 := GCons(__e, symsnd, tmp21877)

tmp21879 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21878)


_ = tmp21879

tmp21880 := MakeNative1(func(__e *ControlFlow, Y1161 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1162 Obj) Obj {
return TailCall2(__e, SymFn(symspecialise), Y1161, Y1162)
panic("unreachable")
})
panic("unreachable")
})

tmp21881 := GCons(__e, symspecialise, tmp21880)

tmp21882 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21881)


_ = tmp21882

tmp21883 := MakeNative1(func(__e *ControlFlow, Y1163 Obj) Obj {
return TailCall1(__e, SymFn(symspy), Y1163)
panic("unreachable")
})

tmp21884 := GCons(__e, symspy, tmp21883)

tmp21885 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21884)


_ = tmp21885

tmp21886 := MakeNative1(func(__e *ControlFlow, Y1164 Obj) Obj {
return TailCall1(__e, SymFn(symstep), Y1164)
panic("unreachable")
})

tmp21887 := GCons(__e, symstep, tmp21886)

tmp21888 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21887)


_ = tmp21888

tmp21889 := MakeNative1(func(__e *ControlFlow, Y1165 Obj) Obj {
return PrimGuard1(__e, symstr, PrimStr, Y1165)
panic("unreachable")
})

tmp21890 := GCons(__e, symstr, tmp21889)

tmp21891 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21890)


_ = tmp21891

tmp21892 := MakeNative1(func(__e *ControlFlow, Y1166 Obj) Obj {
return PrimGuard1(__e, symstring_1_6n, PrimStringToNumber, Y1166)
panic("unreachable")
})

tmp21893 := GCons(__e, symstring_1_6n, tmp21892)

tmp21894 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21893)


_ = tmp21894

tmp21895 := MakeNative1(func(__e *ControlFlow, Y1167 Obj) Obj {
return TailCall1(__e, SymFn(symstring_1_6symbol), Y1167)
panic("unreachable")
})

tmp21896 := GCons(__e, symstring_1_6symbol, tmp21895)

tmp21897 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21896)


_ = tmp21897

tmp21898 := MakeNative1(func(__e *ControlFlow, Y1168 Obj) Obj {
return PrimGuard1(__e, symstring_2, PrimIsString, Y1168)
panic("unreachable")
})

tmp21899 := GCons(__e, symstring_2, tmp21898)

tmp21900 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21899)


_ = tmp21900

tmp21901 := MakeNative1(func(__e *ControlFlow, Y1169 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1170 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1171 Obj) Obj {
return TailCall3(__e, SymFn(symsubst), Y1169, Y1170, Y1171)
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp21902 := GCons(__e, symsubst, tmp21901)

tmp21903 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21902)


_ = tmp21903

tmp21904 := MakeNative1(func(__e *ControlFlow, Y1172 Obj) Obj {
return TailCall1(__e, SymFn(symsum), Y1172)
panic("unreachable")
})

tmp21905 := GCons(__e, symsum, tmp21904)

tmp21906 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21905)


_ = tmp21906

tmp21907 := MakeNative1(func(__e *ControlFlow, Y1173 Obj) Obj {
return PrimGuard1(__e, symsymbol_2, PrimIsSymbol, Y1173)
panic("unreachable")
})

tmp21908 := GCons(__e, symsymbol_2, tmp21907)

tmp21909 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21908)


_ = tmp21909

tmp21910 := MakeNative1(func(__e *ControlFlow, Y1174 Obj) Obj {
return TailCall1(__e, SymFn(symsystemf), Y1174)
panic("unreachable")
})

tmp21911 := GCons(__e, symsystemf, tmp21910)

tmp21912 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21911)


_ = tmp21912

tmp21913 := MakeNative1(func(__e *ControlFlow, Y1175 Obj) Obj {
return TailCall1(__e, SymFn(symtail), Y1175)
panic("unreachable")
})

tmp21914 := GCons(__e, symtail, tmp21913)

tmp21915 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21914)


_ = tmp21915

tmp21916 := MakeNative1(func(__e *ControlFlow, Y1176 Obj) Obj {
return GTl(__e, Y1176)
panic("unreachable")
})

tmp21917 := GCons(__e, symtl, tmp21916)

tmp21918 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21917)


_ = tmp21918

tmp21919 := MakeNative1(func(__e *ControlFlow, Y1177 Obj) Obj {
return TailCall1(__e, SymFn(symtc), Y1177)
panic("unreachable")
})

tmp21920 := GCons(__e, symtc, tmp21919)

tmp21921 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21920)


_ = tmp21921

tmp21922 := MakeNative1(func(__e *ControlFlow, Y1178 Obj) Obj {
return TailCall1(__e, SymFn(symthaw), Y1178)
panic("unreachable")
})

tmp21923 := GCons(__e, symthaw, tmp21922)

tmp21924 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21923)


_ = tmp21924

tmp21925 := MakeNative1(func(__e *ControlFlow, Y1179 Obj) Obj {
return PrimGuard1(__e, symtlstr, PrimTailString, Y1179)
panic("unreachable")
})

tmp21926 := GCons(__e, symtlstr, tmp21925)

tmp21927 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21926)


_ = tmp21927

tmp21928 := MakeNative1(func(__e *ControlFlow, Y1180 Obj) Obj {
return TailCall1(__e, SymFn(symtrack), Y1180)
panic("unreachable")
})

tmp21929 := GCons(__e, symtrack, tmp21928)

tmp21930 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21929)


_ = tmp21930

tmp21931 := MakeNative1(func(__e *ControlFlow, Y1181 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1182 Obj) Obj {
tmp21932 := MakeNative0(func(__e *ControlFlow) Obj {
return Y1181
panic("unreachable")
})

return TailCall2(__e, try_1catch, tmp21932, Y1182)


panic("unreachable")
})
panic("unreachable")
})

tmp21933 := GCons(__e, symtrap_1error, tmp21931)

tmp21934 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21933)


_ = tmp21934

tmp21935 := MakeNative1(func(__e *ControlFlow, Y1183 Obj) Obj {
return TailCall1(__e, SymFn(symtuple_2), Y1183)
panic("unreachable")
})

tmp21936 := GCons(__e, symtuple_2, tmp21935)

tmp21937 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21936)


_ = tmp21937

tmp21938 := MakeNative1(func(__e *ControlFlow, Y1184 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1185 Obj) Obj {
return Y1185
panic("unreachable")
})
panic("unreachable")
})

tmp21939 := GCons(__e, symtype, tmp21938)

tmp21940 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21939)


_ = tmp21940

tmp21941 := MakeNative1(func(__e *ControlFlow, Y1186 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1187 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1188 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1189 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1190 Obj) Obj {
return TailCall5(__e, SymFn(symreturn), Y1186, Y1187, Y1188, Y1189, Y1190)
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp21942 := GCons(__e, symreturn, tmp21941)

tmp21943 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21942)


_ = tmp21943

tmp21944 := MakeNative1(func(__e *ControlFlow, Y1191 Obj) Obj {
return TailCall1(__e, SymFn(symunabsolute), Y1191)
panic("unreachable")
})

tmp21945 := GCons(__e, symunabsolute, tmp21944)

tmp21946 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21945)


_ = tmp21946

tmp21947 := MakeNative1(func(__e *ControlFlow, Y1192 Obj) Obj {
return TailCall1(__e, SymFn(symundefmacro), Y1192)
panic("unreachable")
})

tmp21948 := GCons(__e, symundefmacro, tmp21947)

tmp21949 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21948)


_ = tmp21949

tmp21950 := MakeNative1(func(__e *ControlFlow, Y1193 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1194 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1195 Obj) Obj {
return TailCall3(__e, SymFn(symunput), Y1193, Y1194, Y1195)
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp21951 := GCons(__e, symunput, tmp21950)

tmp21952 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21951)


_ = tmp21952

tmp21953 := MakeNative1(func(__e *ControlFlow, Y1196 Obj) Obj {
return TailCall1(__e, SymFn(symunprofile), Y1196)
panic("unreachable")
})

tmp21954 := GCons(__e, symunprofile, tmp21953)

tmp21955 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21954)


_ = tmp21955

tmp21956 := MakeNative1(func(__e *ControlFlow, Y1197 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1198 Obj) Obj {
return TailCall2(__e, SymFn(symunion), Y1197, Y1198)
panic("unreachable")
})
panic("unreachable")
})

tmp21957 := GCons(__e, symunion, tmp21956)

tmp21958 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21957)


_ = tmp21958

tmp21959 := MakeNative1(func(__e *ControlFlow, Y1199 Obj) Obj {
return TailCall1(__e, SymFn(symuntrack), Y1199)
panic("unreachable")
})

tmp21960 := GCons(__e, symuntrack, tmp21959)

tmp21961 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21960)


_ = tmp21961

tmp21962 := MakeNative1(func(__e *ControlFlow, Y1200 Obj) Obj {
return TailCall1(__e, SymFn(symundefmacro), Y1200)
panic("unreachable")
})

tmp21963 := GCons(__e, symundefmacro, tmp21962)

tmp21964 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21963)


_ = tmp21964

tmp21965 := MakeNative1(func(__e *ControlFlow, Y1201 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1202 Obj) Obj {
return TailCall2(__e, SymFn(symupdate_1lambda_1table), Y1201, Y1202)
panic("unreachable")
})
panic("unreachable")
})

tmp21966 := GCons(__e, symupdate_1lambda_1table, tmp21965)

tmp21967 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21966)


_ = tmp21967

tmp21968 := MakeNative1(func(__e *ControlFlow, Y1203 Obj) Obj {
return TailCall1(__e, SymFn(symvector), Y1203)
panic("unreachable")
})

tmp21969 := GCons(__e, symvector, tmp21968)

tmp21970 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21969)


_ = tmp21970

tmp21971 := MakeNative1(func(__e *ControlFlow, Y1204 Obj) Obj {
return TailCall1(__e, SymFn(symvector_2), Y1204)
panic("unreachable")
})

tmp21972 := GCons(__e, symvector_2, tmp21971)

tmp21973 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21972)


_ = tmp21973

tmp21974 := MakeNative1(func(__e *ControlFlow, Y1205 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1206 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1207 Obj) Obj {
return TailCall3(__e, SymFn(symvector_1_6), Y1205, Y1206, Y1207)
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp21975 := GCons(__e, symvector_1_6, tmp21974)

tmp21976 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21975)


_ = tmp21976

tmp21977 := MakeNative1(func(__e *ControlFlow, Y1208 Obj) Obj {
return GValue(__e, Y1208)
panic("unreachable")
})

tmp21978 := GCons(__e, symvalue, tmp21977)

tmp21979 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21978)


_ = tmp21979

tmp21980 := MakeNative1(func(__e *ControlFlow, Y1209 Obj) Obj {
return PrimGuard1(__e, symvariable_2, PrimIsVariable, Y1209)
panic("unreachable")
})

tmp21981 := GCons(__e, symvariable_2, tmp21980)

tmp21982 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21981)


_ = tmp21982

tmp21983 := MakeNative1(func(__e *ControlFlow, Y1210 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1211 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1212 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1213 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1214 Obj) Obj {
return TailCall5(__e, SymFn(symvar_2), Y1210, Y1211, Y1212, Y1213, Y1214)
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp21984 := GCons(__e, symvar_2, tmp21983)

tmp21985 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21984)


_ = tmp21985

tmp21986 := MakeNative1(func(__e *ControlFlow, Y1215 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1216 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1217 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1218 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1219 Obj) Obj {
return TailCall5(__e, SymFn(symwhen), Y1215, Y1216, Y1217, Y1218, Y1219)
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp21987 := GCons(__e, symwhen, tmp21986)

tmp21988 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21987)


_ = tmp21988

tmp21989 := MakeNative1(func(__e *ControlFlow, Y1220 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1221 Obj) Obj {
return PrimGuard2(__e, symwrite_1byte, PrimWriteByte, Y1220, Y1221)
panic("unreachable")
})
panic("unreachable")
})

tmp21990 := GCons(__e, symwrite_1byte, tmp21989)

tmp21991 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21990)


_ = tmp21991

tmp21992 := MakeNative1(func(__e *ControlFlow, Y1222 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1223 Obj) Obj {
return TailCall2(__e, SymFn(symwrite_1string), Y1222, Y1223)
panic("unreachable")
})
panic("unreachable")
})

tmp21993 := GCons(__e, symwrite_1string, tmp21992)

tmp21994 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21993)


_ = tmp21994

tmp21995 := MakeNative1(func(__e *ControlFlow, Y1224 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1225 Obj) Obj {
return TailCall2(__e, SymFn(symwrite_1to_1file), Y1224, Y1225)
panic("unreachable")
})
panic("unreachable")
})

tmp21996 := GCons(__e, symwrite_1to_1file, tmp21995)

tmp21997 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21996)


_ = tmp21997

tmp21998 := MakeNative1(func(__e *ControlFlow, Y1226 Obj) Obj {
return TailCall1(__e, SymFn(symy_1or_1n_2), Y1226)
panic("unreachable")
})

tmp21999 := GCons(__e, symy_1or_1n_2, tmp21998)

tmp22000 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21999)


_ = tmp22000

tmp22001 := MakeNative1(func(__e *ControlFlow, Y1227 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1228 Obj) Obj {
return GAdd(__e, Y1227, Y1228)
panic("unreachable")
})
panic("unreachable")
})

tmp22002 := GCons(__e, sym_7, tmp22001)

tmp22003 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp22002)


_ = tmp22003

tmp22004 := MakeNative1(func(__e *ControlFlow, Y1229 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1230 Obj) Obj {
return GMul(__e, Y1229, Y1230)
panic("unreachable")
})
panic("unreachable")
})

tmp22005 := GCons(__e, sym_d, tmp22004)

tmp22006 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp22005)


_ = tmp22006

tmp22007 := MakeNative1(func(__e *ControlFlow, Y1231 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1232 Obj) Obj {
return GDiv(__e, Y1231, Y1232)
panic("unreachable")
})
panic("unreachable")
})

tmp22008 := GCons(__e, sym_c, tmp22007)

tmp22009 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp22008)


_ = tmp22009

tmp22010 := MakeNative1(func(__e *ControlFlow, Y1233 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1234 Obj) Obj {
return GSub(__e, Y1233, Y1234)
panic("unreachable")
})
panic("unreachable")
})

tmp22011 := GCons(__e, sym_1, tmp22010)

tmp22012 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp22011)


_ = tmp22012

tmp22013 := MakeNative1(func(__e *ControlFlow, Y1235 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1236 Obj) Obj {
return TailCall2(__e, SymFn(sym_a_a), Y1235, Y1236)
panic("unreachable")
})
panic("unreachable")
})

tmp22014 := GCons(__e, sym_a_a, tmp22013)

tmp22015 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp22014)


_ = tmp22015

tmp22016 := MakeNative1(func(__e *ControlFlow, Y1237 Obj) Obj {
return TailCall1(__e, SymFn(sym_5e_6), Y1237)
panic("unreachable")
})

tmp22017 := GCons(__e, sym_5e_6, tmp22016)

tmp22018 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp22017)


_ = tmp22018

tmp22019 := MakeNative1(func(__e *ControlFlow, Y1238 Obj) Obj {
return TailCall1(__e, SymFn(sym_5end_6), Y1238)
panic("unreachable")
})

tmp22020 := GCons(__e, sym_5end_6, tmp22019)

tmp22021 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp22020)


_ = tmp22021

tmp22022 := MakeNative1(func(__e *ControlFlow, Y1239 Obj) Obj {
return TailCall1(__e, SymFn(sym_5_b_6), Y1239)
panic("unreachable")
})

tmp22023 := GCons(__e, sym_5_b_6, tmp22022)

tmp22024 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp22023)


_ = tmp22024

tmp22025 := MakeNative1(func(__e *ControlFlow, Y1240 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1241 Obj) Obj {
return TailCall2(__e, SymFn(sym_8p), Y1240, Y1241)
panic("unreachable")
})
panic("unreachable")
})

tmp22026 := GCons(__e, sym_8p, tmp22025)

tmp22027 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp22026)


_ = tmp22027

tmp22028 := MakeNative1(func(__e *ControlFlow, Y1242 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1243 Obj) Obj {
return TailCall2(__e, SymFn(sym_8v), Y1242, Y1243)
panic("unreachable")
})
panic("unreachable")
})

tmp22029 := GCons(__e, sym_8v, tmp22028)

tmp22030 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp22029)


_ = tmp22030

tmp22031 := MakeNative1(func(__e *ControlFlow, Y1244 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1245 Obj) Obj {
return TailCall2(__e, SymFn(sym_8s), Y1244, Y1245)
panic("unreachable")
})
panic("unreachable")
})

tmp22032 := GCons(__e, sym_8s, tmp22031)

return TailCall1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp22032)


panic("unreachable")
})

tmp22033 := Call2(__e, ns2_1set, symshen_4initialise_1lambda_1forms, tmp21478)


_ = tmp22033

tmp22034 := MakeNative0(func(__e *ControlFlow) Obj {
tmp22035 := Call0(__e, SymFn(symshen_4initialise_1environment))


_ = tmp22035

tmp22036 := Call0(__e, SymFn(symshen_4initialise_1lambda_1forms))


_ = tmp22036

return TailCall0(__e, SymFn(symshen_4initialise_1signedfuncs))


panic("unreachable")
})

return TailCall2(__e, ns2_1set, symshen_4initialise, tmp22034)




panic("unreachable")
})

