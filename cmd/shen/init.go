package main

import . "github.com/pyrex41/shen-go/kl"

var InitMain = MakeNative0(func(__e *ControlFlow) Obj {
tmp18073 := MakeNative0(func(__e *ControlFlow) Obj {
tmp18074 := PrimGuard2(__e, symset, PrimSet, symshen_4_dhistory_d, Nil)

_ = tmp18074

tmp18075 := PrimGuard2(__e, symset, PrimSet, symshen_4_dtc_d, False)

_ = tmp18075

tmp18076 := Call1(__e, SymFn(symshen_4dict), MakeInteger(20000))


tmp18077 := PrimGuard2(__e, symset, PrimSet, sym_dproperty_1vector_d, tmp18076)

_ = tmp18077

tmp18078 := MakeNative1(func(__e *ControlFlow, X Obj) Obj {
return TailCall1(__e, SymFn(symshen_4macros), X)
panic("unreachable")
})

tmp18079 := GCons(__e, symshen_4macros, tmp18078)

tmp18080 := GCons(__e, tmp18079, Nil)

tmp18081 := PrimGuard2(__e, symset, PrimSet, sym_dmacros_d, tmp18080)

_ = tmp18081

tmp18082 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsource_1form_1handlers_d, Nil)

_ = tmp18082

tmp18083 := PrimGuard2(__e, symset, PrimSet, symshen_4_dgensym_d, MakeInteger(0))

_ = tmp18083

tmp18084 := PrimGuard2(__e, symset, PrimSet, symshen_4_dtracking_d, Nil)

_ = tmp18084

tmp18085 := PrimGuard2(__e, symset, PrimSet, symshen_4_dprofiled_d, Nil)

_ = tmp18085

tmp18086 := GCons(__e, symtype, Nil)

tmp18087 := GCons(__e, symshen_4input_1h_7, tmp18086)

tmp18088 := GCons(__e, symopen, tmp18087)

tmp18089 := GCons(__e, symset, tmp18088)

tmp18090 := GCons(__e, symwhere, tmp18089)

tmp18091 := GCons(__e, symlet, tmp18090)

tmp18092 := GCons(__e, symlambda, tmp18091)

tmp18093 := GCons(__e, symcons, tmp18092)

tmp18094 := GCons(__e, sym_8v, tmp18093)

tmp18095 := GCons(__e, sym_8s, tmp18094)

tmp18096 := GCons(__e, sym_8p, tmp18095)

tmp18097 := PrimGuard2(__e, symset, PrimSet, symshen_4_dspecial_d, tmp18096)

_ = tmp18097

tmp18098 := PrimGuard2(__e, symset, PrimSet, symshen_4_dextraspecial_d, Nil)

_ = tmp18098

tmp18099 := PrimGuard2(__e, symset, PrimSet, symshen_4_dspy_d, False)

_ = tmp18099

tmp18100 := PrimGuard2(__e, symset, PrimSet, symshen_4_ddatatypes_d, Nil)

_ = tmp18100

tmp18101 := PrimGuard2(__e, symset, PrimSet, symshen_4_dalldatatypes_d, Nil)

_ = tmp18101

tmp18102 := PrimGuard2(__e, symset, PrimSet, symshen_4_dshen_1type_1theory_1enabled_2_d, True)

_ = tmp18102

tmp18103 := PrimGuard2(__e, symset, PrimSet, symshen_4_dpackage_d, symnull)

_ = tmp18103

tmp18104 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsynonyms_d, Nil)

_ = tmp18104

tmp18105 := PrimGuard2(__e, symset, PrimSet, symshen_4_doccurs_d, True)

_ = tmp18105

tmp18106 := PrimGuard2(__e, symset, PrimSet, symshen_4_dfactorise_2_d, False)

_ = tmp18106

tmp18107 := PrimGuard2(__e, symset, PrimSet, symshen_4_dmaxinferences_d, MakeInteger(1000000))

_ = tmp18107

tmp18108 := PrimGuard2(__e, symset, PrimSet, sym_dmaximum_1print_1sequence_1size_d, MakeInteger(20))

_ = tmp18108

tmp18109 := PrimGuard2(__e, symset, PrimSet, symshen_4_dcall_d, MakeInteger(0))

_ = tmp18109

tmp18110 := PrimGuard2(__e, symset, PrimSet, symshen_4_dinfs_d, MakeInteger(0))

_ = tmp18110

tmp18111 := PrimGuard2(__e, symset, PrimSet, sym_dhush_d, False)

_ = tmp18111

tmp18112 := PrimGuard2(__e, symset, PrimSet, symshen_4_doptimise_d, False)

_ = tmp18112

tmp18113 := PrimGuard2(__e, symset, PrimSet, sym_dversion_d, MakeString("43.2"))

_ = tmp18113

tmp18114 := PrimGuard2(__e, symset, PrimSet, symshen_4_dnames_d, Nil)

_ = tmp18114

tmp18115 := PrimGuard2(__e, symset, PrimSet, symshen_4_dstep_d, False)

_ = tmp18115

tmp18116 := PrimGuard2(__e, symset, PrimSet, symshen_4_dit_d, MakeString(""))

_ = tmp18116

tmp18117 := PrimGuard2(__e, symset, PrimSet, symshen_4_dresidue_d, Nil)

_ = tmp18117

tmp18118 := PrimGuard2(__e, symset, PrimSet, sym_dabsolute_d, Nil)

_ = tmp18118

tmp18119 := PrimGuard2(__e, symset, PrimSet, symshen_4_dprolog_1memory_d, MakeInteger(1000))

_ = tmp18119

tmp18120 := PrimGuard2(__e, symset, PrimSet, symshen_4_dloading_2_d, False)

_ = tmp18120

tmp18121 := PrimGuard2(__e, symset, PrimSet, symshen_4_duserdefs_d, Nil)

_ = tmp18121

tmp18122 := MakeNative1(func(__e *ControlFlow, X Obj) Obj {
return X
panic("unreachable")
})

tmp18123 := PrimGuard2(__e, symset, PrimSet, symshen_4_ddemodulation_1function_d, tmp18122)

_ = tmp18123

tmp18124 := PrimGuard2(__e, symset, PrimSet, symshen_4_dcustom_1pattern_1compiler_d, False)

_ = tmp18124

tmp18125 := PrimGuard2(__e, symset, PrimSet, symshen_4_dcustom_1pattern_1reducer_d, False)

_ = tmp18125

tmp18128 := Call1(__e, SymFn(symbound_2), sym_dhome_1directory_d)


tmp18129 := GNot(__e, tmp18128)

var ifres18126 Obj

if True == tmp18129 {
tmp18127 := PrimGuard2(__e, symset, PrimSet, sym_dhome_1directory_d, MakeString(""))

ifres18126 = tmp18127


} else {
ifres18126 = symshen_4skip


}

_ = ifres18126

tmp18133 := Call1(__e, SymFn(symbound_2), sym_dsterror_d)


tmp18134 := GNot(__e, tmp18133)

var ifres18130 Obj

if True == tmp18134 {
tmp18131 := GValue(__e, sym_dstoutput_d)

tmp18132 := PrimGuard2(__e, symset, PrimSet, sym_dsterror_d, tmp18131)

ifres18130 = tmp18132


} else {
ifres18130 = symshen_4skip


}

_ = ifres18130

tmp18135 := Call1(__e, SymFn(symprolog_1memory), MakeInteger(10000))


_ = tmp18135

tmp18136 := PrimGuard2(__e, symset, PrimSet, symshen_4_dloading_2_d, False)

_ = tmp18136

tmp18137 := GCons(__e, MakeInteger(2), Nil)

tmp18138 := GCons(__e, sym_8s, tmp18137)

tmp18139 := GCons(__e, MakeInteger(2), tmp18138)

tmp18140 := GCons(__e, sym_8v, tmp18139)

tmp18141 := GCons(__e, MakeInteger(2), tmp18140)

tmp18142 := GCons(__e, sym_8p, tmp18141)

tmp18143 := GCons(__e, MakeInteger(1), tmp18142)

tmp18144 := GCons(__e, sym_5_b_6, tmp18143)

tmp18145 := GCons(__e, MakeInteger(1), tmp18144)

tmp18146 := GCons(__e, sym_5end_6, tmp18145)

tmp18147 := GCons(__e, MakeInteger(1), tmp18146)

tmp18148 := GCons(__e, sym_5e_6, tmp18147)

tmp18149 := GCons(__e, MakeInteger(2), tmp18148)

tmp18150 := GCons(__e, sym_a_a, tmp18149)

tmp18151 := GCons(__e, MakeInteger(2), tmp18150)

tmp18152 := GCons(__e, sym_1, tmp18151)

tmp18153 := GCons(__e, MakeInteger(2), tmp18152)

tmp18154 := GCons(__e, sym_c, tmp18153)

tmp18155 := GCons(__e, MakeInteger(2), tmp18154)

tmp18156 := GCons(__e, sym_d, tmp18155)

tmp18157 := GCons(__e, MakeInteger(2), tmp18156)

tmp18158 := GCons(__e, sym_7, tmp18157)

tmp18159 := GCons(__e, MakeInteger(1), tmp18158)

tmp18160 := GCons(__e, symy_1or_1n_2, tmp18159)

tmp18161 := GCons(__e, MakeInteger(2), tmp18160)

tmp18162 := GCons(__e, symwrite_1to_1file, tmp18161)

tmp18163 := GCons(__e, MakeInteger(2), tmp18162)

tmp18164 := GCons(__e, symwrite_1string, tmp18163)

tmp18165 := GCons(__e, MakeInteger(2), tmp18164)

tmp18166 := GCons(__e, symwrite_1byte, tmp18165)

tmp18167 := GCons(__e, MakeInteger(5), tmp18166)

tmp18168 := GCons(__e, symwhen, tmp18167)

tmp18169 := GCons(__e, MakeInteger(0), tmp18168)

tmp18170 := GCons(__e, symversion, tmp18169)

tmp18171 := GCons(__e, MakeInteger(5), tmp18170)

tmp18172 := GCons(__e, symvar_2, tmp18171)

tmp18173 := GCons(__e, MakeInteger(1), tmp18172)

tmp18174 := GCons(__e, symvariable_2, tmp18173)

tmp18175 := GCons(__e, MakeInteger(1), tmp18174)

tmp18176 := GCons(__e, symvalue, tmp18175)

tmp18177 := GCons(__e, MakeInteger(3), tmp18176)

tmp18178 := GCons(__e, symvector_1_6, tmp18177)

tmp18179 := GCons(__e, MakeInteger(1), tmp18178)

tmp18180 := GCons(__e, symvector_2, tmp18179)

tmp18181 := GCons(__e, MakeInteger(1), tmp18180)

tmp18182 := GCons(__e, symvector, tmp18181)

tmp18183 := GCons(__e, MakeInteger(0), tmp18182)

tmp18184 := GCons(__e, symuserdefs, tmp18183)

tmp18185 := GCons(__e, MakeInteger(2), tmp18184)

tmp18186 := GCons(__e, symupdate_1lambda_1table, tmp18185)

tmp18187 := GCons(__e, MakeInteger(1), tmp18186)

tmp18188 := GCons(__e, symundefmacro, tmp18187)

tmp18189 := GCons(__e, MakeInteger(1), tmp18188)

tmp18190 := GCons(__e, symuntrack, tmp18189)

tmp18191 := GCons(__e, MakeInteger(2), tmp18190)

tmp18192 := GCons(__e, symunion, tmp18191)

tmp18193 := GCons(__e, MakeInteger(1), tmp18192)

tmp18194 := GCons(__e, symunprofile, tmp18193)

tmp18195 := GCons(__e, MakeInteger(3), tmp18194)

tmp18196 := GCons(__e, symunput, tmp18195)

tmp18197 := GCons(__e, MakeInteger(1), tmp18196)

tmp18198 := GCons(__e, symundefmacro, tmp18197)

tmp18199 := GCons(__e, MakeInteger(1), tmp18198)

tmp18200 := GCons(__e, symunabsolute, tmp18199)

tmp18201 := GCons(__e, MakeInteger(5), tmp18200)

tmp18202 := GCons(__e, symreturn, tmp18201)

tmp18203 := GCons(__e, MakeInteger(2), tmp18202)

tmp18204 := GCons(__e, symtype, tmp18203)

tmp18205 := GCons(__e, MakeInteger(1), tmp18204)

tmp18206 := GCons(__e, symtuple_2, tmp18205)

tmp18207 := GCons(__e, MakeInteger(2), tmp18206)

tmp18208 := GCons(__e, symtrap_1error, tmp18207)

tmp18209 := GCons(__e, MakeInteger(0), tmp18208)

tmp18210 := GCons(__e, symtracked, tmp18209)

tmp18211 := GCons(__e, MakeInteger(1), tmp18210)

tmp18212 := GCons(__e, symtrack, tmp18211)

tmp18213 := GCons(__e, MakeInteger(1), tmp18212)

tmp18214 := GCons(__e, symtlstr, tmp18213)

tmp18215 := GCons(__e, MakeInteger(1), tmp18214)

tmp18216 := GCons(__e, symthaw, tmp18215)

tmp18217 := GCons(__e, MakeInteger(0), tmp18216)

tmp18218 := GCons(__e, symtc_2, tmp18217)

tmp18219 := GCons(__e, MakeInteger(1), tmp18218)

tmp18220 := GCons(__e, symtc, tmp18219)

tmp18221 := GCons(__e, MakeInteger(1), tmp18220)

tmp18222 := GCons(__e, symtl, tmp18221)

tmp18223 := GCons(__e, MakeInteger(1), tmp18222)

tmp18224 := GCons(__e, symtail, tmp18223)

tmp18225 := GCons(__e, MakeInteger(1), tmp18224)

tmp18226 := GCons(__e, symsystemf, tmp18225)

tmp18227 := GCons(__e, MakeInteger(1), tmp18226)

tmp18228 := GCons(__e, symsymbol_2, tmp18227)

tmp18229 := GCons(__e, MakeInteger(1), tmp18228)

tmp18230 := GCons(__e, symsum, tmp18229)

tmp18231 := GCons(__e, MakeInteger(3), tmp18230)

tmp18232 := GCons(__e, symsubst, tmp18231)

tmp18233 := GCons(__e, MakeInteger(1), tmp18232)

tmp18234 := GCons(__e, symstring_2, tmp18233)

tmp18235 := GCons(__e, MakeInteger(1), tmp18234)

tmp18236 := GCons(__e, symstring_1_6symbol, tmp18235)

tmp18237 := GCons(__e, MakeInteger(1), tmp18236)

tmp18238 := GCons(__e, symstring_1_6n, tmp18237)

tmp18239 := GCons(__e, MakeInteger(1), tmp18238)

tmp18240 := GCons(__e, symstr, tmp18239)

tmp18241 := GCons(__e, MakeInteger(0), tmp18240)

tmp18242 := GCons(__e, symstoutput, tmp18241)

tmp18243 := GCons(__e, MakeInteger(0), tmp18242)

tmp18244 := GCons(__e, symstinput, tmp18243)

tmp18245 := GCons(__e, MakeInteger(0), tmp18244)

tmp18246 := GCons(__e, symstep_2, tmp18245)

tmp18247 := GCons(__e, MakeInteger(1), tmp18246)

tmp18248 := GCons(__e, symstep, tmp18247)

tmp18249 := GCons(__e, MakeInteger(0), tmp18248)

tmp18250 := GCons(__e, symspy_2, tmp18249)

tmp18251 := GCons(__e, MakeInteger(1), tmp18250)

tmp18252 := GCons(__e, symspy, tmp18251)

tmp18253 := GCons(__e, MakeInteger(2), tmp18252)

tmp18254 := GCons(__e, symspecialise, tmp18253)

tmp18255 := GCons(__e, MakeInteger(1), tmp18254)

tmp18256 := GCons(__e, symsnd, tmp18255)

tmp18257 := GCons(__e, MakeInteger(1), tmp18256)

tmp18258 := GCons(__e, symsimple_1error, tmp18257)

tmp18259 := GCons(__e, MakeInteger(2), tmp18258)

tmp18260 := GCons(__e, symset, tmp18259)

tmp18261 := GCons(__e, MakeInteger(1), tmp18260)

tmp18262 := GCons(__e, symreverse, tmp18261)

tmp18263 := GCons(__e, MakeInteger(2), tmp18262)

tmp18264 := GCons(__e, symremove, tmp18263)

tmp18265 := GCons(__e, MakeInteger(1), tmp18264)

tmp18266 := GCons(__e, symshen_4unregister_1source_1form, tmp18265)

tmp18267 := GCons(__e, MakeInteger(2), tmp18266)

tmp18268 := GCons(__e, symshen_4register_1source_1form, tmp18267)

tmp18269 := GCons(__e, MakeInteger(0), tmp18268)

tmp18270 := GCons(__e, symrelease, tmp18269)

tmp18271 := GCons(__e, MakeInteger(1), tmp18270)

tmp18272 := GCons(__e, symreceive, tmp18271)

tmp18273 := GCons(__e, MakeInteger(1), tmp18272)

tmp18274 := GCons(__e, symread_1unit_1string, tmp18273)

tmp18275 := GCons(__e, MakeInteger(1), tmp18274)

tmp18276 := GCons(__e, symread_1from_1string_1unprocessed, tmp18275)

tmp18277 := GCons(__e, MakeInteger(1), tmp18276)

tmp18278 := GCons(__e, symread_1from_1string, tmp18277)

tmp18279 := GCons(__e, MakeInteger(1), tmp18278)

tmp18280 := GCons(__e, symread_1byte, tmp18279)

tmp18281 := GCons(__e, MakeInteger(1), tmp18280)

tmp18282 := GCons(__e, symread, tmp18281)

tmp18283 := GCons(__e, MakeInteger(1), tmp18282)

tmp18284 := GCons(__e, symread_1file, tmp18283)

tmp18285 := GCons(__e, MakeInteger(1), tmp18284)

tmp18286 := GCons(__e, symread_1file_1as_1bytelist, tmp18285)

tmp18287 := GCons(__e, MakeInteger(1), tmp18286)

tmp18288 := GCons(__e, symread_1file_1as_1string, tmp18287)

tmp18289 := GCons(__e, MakeInteger(4), tmp18288)

tmp18290 := GCons(__e, symput, tmp18289)

tmp18291 := GCons(__e, MakeInteger(1), tmp18290)

tmp18292 := GCons(__e, symprotect, tmp18291)

tmp18293 := GCons(__e, MakeInteger(1), tmp18292)

tmp18294 := GCons(__e, sympreclude_1all_1but, tmp18293)

tmp18295 := GCons(__e, MakeInteger(1), tmp18294)

tmp18296 := GCons(__e, sympreclude, tmp18295)

tmp18297 := GCons(__e, MakeInteger(1), tmp18296)

tmp18298 := GCons(__e, symps, tmp18297)

tmp18299 := GCons(__e, MakeInteger(2), tmp18298)

tmp18300 := GCons(__e, sympr, tmp18299)

tmp18301 := GCons(__e, MakeInteger(1), tmp18300)

tmp18302 := GCons(__e, symprofile_1results, tmp18301)

tmp18303 := GCons(__e, MakeInteger(1), tmp18302)

tmp18304 := GCons(__e, symprolog_1memory, tmp18303)

tmp18305 := GCons(__e, MakeInteger(1), tmp18304)

tmp18306 := GCons(__e, symshen_4printF, tmp18305)

tmp18307 := GCons(__e, MakeInteger(1), tmp18306)

tmp18308 := GCons(__e, symshen_4print_1freshterm, tmp18307)

tmp18309 := GCons(__e, MakeInteger(1), tmp18308)

tmp18310 := GCons(__e, symshen_4print_1prolog_1vector, tmp18309)

tmp18311 := GCons(__e, MakeInteger(1), tmp18310)

tmp18312 := GCons(__e, symprofile, tmp18311)

tmp18313 := GCons(__e, MakeInteger(1), tmp18312)

tmp18314 := GCons(__e, symprint, tmp18313)

tmp18315 := GCons(__e, MakeInteger(1), tmp18314)

tmp18316 := GCons(__e, sympreclude_1all_1but, tmp18315)

tmp18317 := GCons(__e, MakeInteger(2), tmp18316)

tmp18318 := GCons(__e, sympos, tmp18317)

tmp18319 := GCons(__e, MakeInteger(0), tmp18318)

tmp18320 := GCons(__e, symporters, tmp18319)

tmp18321 := GCons(__e, MakeInteger(0), tmp18320)

tmp18322 := GCons(__e, symport, tmp18321)

tmp18323 := GCons(__e, MakeInteger(1), tmp18322)

tmp18324 := GCons(__e, sympackage_2, tmp18323)

tmp18325 := GCons(__e, MakeInteger(3), tmp18324)

tmp18326 := GCons(__e, sympackage, tmp18325)

tmp18327 := GCons(__e, MakeInteger(0), tmp18326)

tmp18328 := GCons(__e, symos, tmp18327)

tmp18329 := GCons(__e, MakeInteger(2), tmp18328)

tmp18330 := GCons(__e, symor, tmp18329)

tmp18331 := GCons(__e, MakeInteger(0), tmp18330)

tmp18332 := GCons(__e, symoptimise_2, tmp18331)

tmp18333 := GCons(__e, MakeInteger(1), tmp18332)

tmp18334 := GCons(__e, symoptimise, tmp18333)

tmp18335 := GCons(__e, MakeInteger(2), tmp18334)

tmp18336 := GCons(__e, symopen, tmp18335)

tmp18337 := GCons(__e, MakeInteger(1), tmp18336)

tmp18338 := GCons(__e, symoccurs_1check, tmp18337)

tmp18339 := GCons(__e, MakeInteger(0), tmp18338)

tmp18340 := GCons(__e, symoccurs_2, tmp18339)

tmp18341 := GCons(__e, MakeInteger(2), tmp18340)

tmp18342 := GCons(__e, symoccurrences, tmp18341)

tmp18343 := GCons(__e, MakeInteger(1), tmp18342)

tmp18344 := GCons(__e, symoccurs_1check, tmp18343)

tmp18345 := GCons(__e, MakeInteger(1), tmp18344)

tmp18346 := GCons(__e, symnumber_2, tmp18345)

tmp18347 := GCons(__e, MakeInteger(1), tmp18346)

tmp18348 := GCons(__e, symn_1_6string, tmp18347)

tmp18349 := GCons(__e, MakeInteger(2), tmp18348)

tmp18350 := GCons(__e, symnth, tmp18349)

tmp18351 := GCons(__e, MakeInteger(1), tmp18350)

tmp18352 := GCons(__e, symnot, tmp18351)

tmp18353 := GCons(__e, MakeInteger(1), tmp18352)

tmp18354 := GCons(__e, symnl, tmp18353)

tmp18355 := GCons(__e, MakeInteger(1), tmp18354)

tmp18356 := GCons(__e, symmaxinferences, tmp18355)

tmp18357 := GCons(__e, MakeInteger(2), tmp18356)

tmp18358 := GCons(__e, symmapcan, tmp18357)

tmp18359 := GCons(__e, MakeInteger(2), tmp18358)

tmp18360 := GCons(__e, symmap, tmp18359)

tmp18361 := GCons(__e, MakeInteger(1), tmp18360)

tmp18362 := GCons(__e, symmacroexpand, tmp18361)

tmp18363 := GCons(__e, MakeInteger(1), tmp18362)

tmp18364 := GCons(__e, symvector, tmp18363)

tmp18365 := GCons(__e, MakeInteger(2), tmp18364)

tmp18366 := GCons(__e, sym_5_a, tmp18365)

tmp18367 := GCons(__e, MakeInteger(2), tmp18366)

tmp18368 := GCons(__e, sym_5, tmp18367)

tmp18369 := GCons(__e, MakeInteger(1), tmp18368)

tmp18370 := GCons(__e, symload, tmp18369)

tmp18371 := GCons(__e, MakeInteger(1), tmp18370)

tmp18372 := GCons(__e, symlist, tmp18371)

tmp18373 := GCons(__e, MakeInteger(1), tmp18372)

tmp18374 := GCons(__e, symlineread, tmp18373)

tmp18375 := GCons(__e, MakeInteger(1), tmp18374)

tmp18376 := GCons(__e, symlimit, tmp18375)

tmp18377 := GCons(__e, MakeInteger(1), tmp18376)

tmp18378 := GCons(__e, symlength, tmp18377)

tmp18379 := GCons(__e, MakeInteger(0), tmp18378)

tmp18380 := GCons(__e, symlanguage, tmp18379)

tmp18381 := GCons(__e, MakeInteger(6), tmp18380)

tmp18382 := GCons(__e, symis_b, tmp18381)

tmp18383 := GCons(__e, MakeInteger(6), tmp18382)

tmp18384 := GCons(__e, symis, tmp18383)

tmp18385 := GCons(__e, MakeInteger(0), tmp18384)

tmp18386 := GCons(__e, symit, tmp18385)

tmp18387 := GCons(__e, MakeInteger(1), tmp18386)

tmp18388 := GCons(__e, syminternal, tmp18387)

tmp18389 := GCons(__e, MakeInteger(2), tmp18388)

tmp18390 := GCons(__e, symintersection, tmp18389)

tmp18391 := GCons(__e, MakeInteger(1), tmp18390)

tmp18392 := GCons(__e, syminclude_1all_1but, tmp18391)

tmp18393 := GCons(__e, MakeInteger(0), tmp18392)

tmp18394 := GCons(__e, symimplementation, tmp18393)

tmp18395 := GCons(__e, MakeInteger(2), tmp18394)

tmp18396 := GCons(__e, syminput_7, tmp18395)

tmp18397 := GCons(__e, MakeInteger(1), tmp18396)

tmp18398 := GCons(__e, syminput, tmp18397)

tmp18399 := GCons(__e, MakeInteger(0), tmp18398)

tmp18400 := GCons(__e, syminferences, tmp18399)

tmp18401 := GCons(__e, MakeInteger(1), tmp18400)

tmp18402 := GCons(__e, symintern, tmp18401)

tmp18403 := GCons(__e, MakeInteger(1), tmp18402)

tmp18404 := GCons(__e, syminternal, tmp18403)

tmp18405 := GCons(__e, MakeInteger(1), tmp18404)

tmp18406 := GCons(__e, syminteger_2, tmp18405)

tmp18407 := GCons(__e, MakeInteger(1), tmp18406)

tmp18408 := GCons(__e, symin_1package, tmp18407)

tmp18409 := GCons(__e, MakeInteger(0), tmp18408)

tmp18410 := GCons(__e, symincluded, tmp18409)

tmp18411 := GCons(__e, MakeInteger(1), tmp18410)

tmp18412 := GCons(__e, syminclude, tmp18411)

tmp18413 := GCons(__e, MakeInteger(3), tmp18412)

tmp18414 := GCons(__e, symif, tmp18413)

tmp18415 := GCons(__e, MakeInteger(1), tmp18414)

tmp18416 := GCons(__e, symhush, tmp18415)

tmp18417 := GCons(__e, MakeInteger(0), tmp18416)

tmp18418 := GCons(__e, symhush_2, tmp18417)

tmp18419 := GCons(__e, MakeInteger(1), tmp18418)

tmp18420 := GCons(__e, symhead, tmp18419)

tmp18421 := GCons(__e, MakeInteger(1), tmp18420)

tmp18422 := GCons(__e, symhdstr, tmp18421)

tmp18423 := GCons(__e, MakeInteger(1), tmp18422)

tmp18424 := GCons(__e, symhdv, tmp18423)

tmp18425 := GCons(__e, MakeInteger(1), tmp18424)

tmp18426 := GCons(__e, symhd, tmp18425)

tmp18427 := GCons(__e, MakeInteger(2), tmp18426)

tmp18428 := GCons(__e, symhash, tmp18427)

tmp18429 := GCons(__e, MakeInteger(2), tmp18428)

tmp18430 := GCons(__e, sym_a, tmp18429)

tmp18431 := GCons(__e, MakeInteger(2), tmp18430)

tmp18432 := GCons(__e, sym_6_a, tmp18431)

tmp18433 := GCons(__e, MakeInteger(2), tmp18432)

tmp18434 := GCons(__e, sym_6, tmp18433)

tmp18435 := GCons(__e, MakeInteger(2), tmp18434)

tmp18436 := GCons(__e, sym_5_1vector, tmp18435)

tmp18437 := GCons(__e, MakeInteger(2), tmp18436)

tmp18438 := GCons(__e, sym_5_1address, tmp18437)

tmp18439 := GCons(__e, MakeInteger(3), tmp18438)

tmp18440 := GCons(__e, symaddress_1_6, tmp18439)

tmp18441 := GCons(__e, MakeInteger(1), tmp18440)

tmp18442 := GCons(__e, symget_1time, tmp18441)

tmp18443 := GCons(__e, MakeInteger(3), tmp18442)

tmp18444 := GCons(__e, symget, tmp18443)

tmp18445 := GCons(__e, MakeInteger(1), tmp18444)

tmp18446 := GCons(__e, symgensym, tmp18445)

tmp18447 := GCons(__e, MakeInteger(1), tmp18446)

tmp18448 := GCons(__e, symfunction, tmp18447)

tmp18449 := GCons(__e, MakeInteger(1), tmp18448)

tmp18450 := GCons(__e, symfn, tmp18449)

tmp18451 := GCons(__e, MakeInteger(1), tmp18450)

tmp18452 := GCons(__e, symfst, tmp18451)

tmp18453 := GCons(__e, MakeInteger(0), tmp18452)

tmp18454 := GCons(__e, symfresh, tmp18453)

tmp18455 := GCons(__e, MakeInteger(1), tmp18454)

tmp18456 := GCons(__e, symfreeze, tmp18455)

tmp18457 := GCons(__e, MakeInteger(5), tmp18456)

tmp18458 := GCons(__e, symfork, tmp18457)

tmp18459 := GCons(__e, MakeInteger(1), tmp18458)

tmp18460 := GCons(__e, symforeign, tmp18459)

tmp18461 := GCons(__e, MakeInteger(7), tmp18460)

tmp18462 := GCons(__e, symfindall, tmp18461)

tmp18463 := GCons(__e, MakeInteger(2), tmp18462)

tmp18464 := GCons(__e, symfix, tmp18463)

tmp18465 := GCons(__e, MakeInteger(0), tmp18464)

tmp18466 := GCons(__e, symfail, tmp18465)

tmp18467 := GCons(__e, MakeInteger(2), tmp18466)

tmp18468 := GCons(__e, symfail_1if, tmp18467)

tmp18469 := GCons(__e, MakeInteger(0), tmp18468)

tmp18470 := GCons(__e, symfactorise_2, tmp18469)

tmp18471 := GCons(__e, MakeInteger(1), tmp18470)

tmp18472 := GCons(__e, symfactorise, tmp18471)

tmp18473 := GCons(__e, MakeInteger(1), tmp18472)

tmp18474 := GCons(__e, symexternal, tmp18473)

tmp18475 := GCons(__e, MakeInteger(1), tmp18474)

tmp18476 := GCons(__e, symexplode, tmp18475)

tmp18477 := GCons(__e, MakeInteger(1), tmp18476)

tmp18478 := GCons(__e, symeval_1kl, tmp18477)

tmp18479 := GCons(__e, MakeInteger(1), tmp18478)

tmp18480 := GCons(__e, symeval, tmp18479)

tmp18481 := GCons(__e, MakeInteger(1), tmp18480)

tmp18482 := GCons(__e, symerror_1to_1string, tmp18481)

tmp18483 := GCons(__e, MakeInteger(1), tmp18482)

tmp18484 := GCons(__e, symexternal, tmp18483)

tmp18485 := GCons(__e, MakeInteger(1), tmp18484)

tmp18486 := GCons(__e, symenable_1type_1theory, tmp18485)

tmp18487 := GCons(__e, MakeInteger(1), tmp18486)

tmp18488 := GCons(__e, symempty_2, tmp18487)

tmp18489 := GCons(__e, MakeInteger(2), tmp18488)

tmp18490 := GCons(__e, symelement_2, tmp18489)

tmp18491 := GCons(__e, MakeInteger(2), tmp18490)

tmp18492 := GCons(__e, symdo, tmp18491)

tmp18493 := GCons(__e, MakeInteger(2), tmp18492)

tmp18494 := GCons(__e, symdifference, tmp18493)

tmp18495 := GCons(__e, MakeInteger(1), tmp18494)

tmp18496 := GCons(__e, symdestroy, tmp18495)

tmp18497 := GCons(__e, MakeInteger(2), tmp18496)

tmp18498 := GCons(__e, symdeclare, tmp18497)

tmp18499 := GCons(__e, MakeInteger(0), tmp18498)

tmp18500 := GCons(__e, symdatatypes, tmp18499)

tmp18501 := GCons(__e, MakeInteger(1), tmp18500)

tmp18502 := GCons(__e, symclose, tmp18501)

tmp18503 := GCons(__e, MakeInteger(2), tmp18502)

tmp18504 := GCons(__e, symcn, tmp18503)

tmp18505 := GCons(__e, MakeInteger(1), tmp18504)

tmp18506 := GCons(__e, symcons_2, tmp18505)

tmp18507 := GCons(__e, MakeInteger(2), tmp18506)

tmp18508 := GCons(__e, symcons, tmp18507)

tmp18509 := GCons(__e, MakeInteger(2), tmp18508)

tmp18510 := GCons(__e, symconcat, tmp18509)

tmp18511 := GCons(__e, MakeInteger(2), tmp18510)

tmp18512 := GCons(__e, symcompile, tmp18511)

tmp18513 := GCons(__e, MakeInteger(1), tmp18512)

tmp18514 := GCons(__e, symchar_1stream_2, tmp18513)

tmp18515 := GCons(__e, MakeInteger(1), tmp18514)

tmp18516 := GCons(__e, symcd, tmp18515)

tmp18517 := GCons(__e, MakeInteger(5), tmp18516)

tmp18518 := GCons(__e, symcall, tmp18517)

tmp18519 := GCons(__e, MakeInteger(6), tmp18518)

tmp18520 := GCons(__e, symbind, tmp18519)

tmp18521 := GCons(__e, MakeInteger(1), tmp18520)

tmp18522 := GCons(__e, symbound_2, tmp18521)

tmp18523 := GCons(__e, MakeInteger(1), tmp18522)

tmp18524 := GCons(__e, symbootstrap, tmp18523)

tmp18525 := GCons(__e, MakeInteger(1), tmp18524)

tmp18526 := GCons(__e, symboolean_2, tmp18525)

tmp18527 := GCons(__e, MakeInteger(1), tmp18526)

tmp18528 := GCons(__e, symatom_2, tmp18527)

tmp18529 := GCons(__e, MakeInteger(2), tmp18528)

tmp18530 := GCons(__e, symassoc, tmp18529)

tmp18531 := GCons(__e, MakeInteger(1), tmp18530)

tmp18532 := GCons(__e, symarity, tmp18531)

tmp18533 := GCons(__e, MakeInteger(2), tmp18532)

tmp18534 := GCons(__e, symappend, tmp18533)

tmp18535 := GCons(__e, MakeInteger(2), tmp18534)

tmp18536 := GCons(__e, symand, tmp18535)

tmp18537 := GCons(__e, MakeInteger(2), tmp18536)

tmp18538 := GCons(__e, symadjoin, tmp18537)

tmp18539 := GCons(__e, MakeInteger(3), tmp18538)

tmp18540 := GCons(__e, symaddress_1_6, tmp18539)

tmp18541 := GCons(__e, MakeInteger(1), tmp18540)

tmp18542 := GCons(__e, symabsvector, tmp18541)

tmp18543 := GCons(__e, MakeInteger(1), tmp18542)

tmp18544 := GCons(__e, symabsvector_2, tmp18543)

tmp18545 := GCons(__e, MakeInteger(1), tmp18544)

tmp18546 := GCons(__e, symabsolute, tmp18545)

tmp18547 := GCons(__e, MakeInteger(0), tmp18546)

tmp18548 := GCons(__e, symabort, tmp18547)

tmp18549 := Call1(__e, SymFn(symshen_4initialise_1arity_1table), tmp18548)


_ = tmp18549

tmp18550 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":"))

tmp18551 := PrimGuard1(__e, symintern, PrimIntern, MakeString(";"))

tmp18552 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":="))

tmp18553 := PrimGuard1(__e, symintern, PrimIntern, MakeString(","))

tmp18554 := Call1(__e, SymFn(symvector), MakeInteger(0))


tmp18555 := PrimGuard1(__e, symintern, PrimIntern, MakeString("bar!"))

tmp18556 := GCons(__e, symabort, Nil)

tmp18557 := GCons(__e, symabsolute, tmp18556)

tmp18558 := GCons(__e, symabsvector, tmp18557)

tmp18559 := GCons(__e, symabsvector_2, tmp18558)

tmp18560 := GCons(__e, symaddress_1_6, tmp18559)

tmp18561 := GCons(__e, sym_5_1address, tmp18560)

tmp18562 := GCons(__e, symadjoin, tmp18561)

tmp18563 := GCons(__e, symand, tmp18562)

tmp18564 := GCons(__e, symappend, tmp18563)

tmp18565 := GCons(__e, symarity, tmp18564)

tmp18566 := GCons(__e, symassoc, tmp18565)

tmp18567 := GCons(__e, symassertz, tmp18566)

tmp18568 := GCons(__e, symasserta, tmp18567)

tmp18569 := GCons(__e, symatom_2, tmp18568)

tmp18570 := GCons(__e, symstep_2, tmp18569)

tmp18571 := GCons(__e, symspy_2, tmp18570)

tmp18572 := GCons(__e, tmp18555, tmp18571)

tmp18573 := GCons(__e, symbootstrap, tmp18572)

tmp18574 := GCons(__e, symboolean, tmp18573)

tmp18575 := GCons(__e, symboolean_2, tmp18574)

tmp18576 := GCons(__e, symbound_2, tmp18575)

tmp18577 := GCons(__e, symbind, tmp18576)

tmp18578 := GCons(__e, symclose, tmp18577)

tmp18579 := GCons(__e, symchar_1stream_2, tmp18578)

tmp18580 := GCons(__e, symcall, tmp18579)

tmp18581 := GCons(__e, symcases, tmp18580)

tmp18582 := GCons(__e, symcd, tmp18581)

tmp18583 := GCons(__e, symcompile, tmp18582)

tmp18584 := GCons(__e, symconcat, tmp18583)

tmp18585 := GCons(__e, symcond, tmp18584)

tmp18586 := GCons(__e, symcons, tmp18585)

tmp18587 := GCons(__e, symcons_2, tmp18586)

tmp18588 := GCons(__e, symcn, tmp18587)

tmp18589 := GCons(__e, symctxt, tmp18588)

tmp18590 := GCons(__e, symdatatypes, tmp18589)

tmp18591 := GCons(__e, symdatatype, tmp18590)

tmp18592 := GCons(__e, symdeclare, tmp18591)

tmp18593 := GCons(__e, symdefprolog, tmp18592)

tmp18594 := GCons(__e, symdefcc, tmp18593)

tmp18595 := GCons(__e, symdefmacro, tmp18594)

tmp18596 := GCons(__e, symdefine, tmp18595)

tmp18597 := GCons(__e, symdefun, tmp18596)

tmp18598 := GCons(__e, symdestroy, tmp18597)

tmp18599 := GCons(__e, symdifference, tmp18598)

tmp18600 := GCons(__e, symdo, tmp18599)

tmp18601 := GCons(__e, symelement_2, tmp18600)

tmp18602 := GCons(__e, symempty_2, tmp18601)

tmp18603 := GCons(__e, symerror, tmp18602)

tmp18604 := GCons(__e, symerror_1to_1string, tmp18603)

tmp18605 := GCons(__e, symeval, tmp18604)

tmp18606 := GCons(__e, symeval_1kl, tmp18605)

tmp18607 := GCons(__e, symexception, tmp18606)

tmp18608 := GCons(__e, symexternal, tmp18607)

tmp18609 := GCons(__e, symexplode, tmp18608)

tmp18610 := GCons(__e, symenable_1type_1theory, tmp18609)

tmp18611 := GCons(__e, False, tmp18610)

tmp18612 := GCons(__e, symfindall, tmp18611)

tmp18613 := GCons(__e, symfactorise_2, tmp18612)

tmp18614 := GCons(__e, symfactorise, tmp18613)

tmp18615 := GCons(__e, symfail_1if, tmp18614)

tmp18616 := GCons(__e, symfail, tmp18615)

tmp18617 := GCons(__e, symfile, tmp18616)

tmp18618 := GCons(__e, symfix, tmp18617)

tmp18619 := GCons(__e, symforeign, tmp18618)

tmp18620 := GCons(__e, symfork, tmp18619)

tmp18621 := GCons(__e, symfresh, tmp18620)

tmp18622 := GCons(__e, symfreeze, tmp18621)

tmp18623 := GCons(__e, symfst, tmp18622)

tmp18624 := GCons(__e, symfunction, tmp18623)

tmp18625 := GCons(__e, symfn, tmp18624)

tmp18626 := GCons(__e, symgensym, tmp18625)

tmp18627 := GCons(__e, symget_1time, tmp18626)

tmp18628 := GCons(__e, symget, tmp18627)

tmp18629 := GCons(__e, symhash, tmp18628)

tmp18630 := GCons(__e, symhdstr, tmp18629)

tmp18631 := GCons(__e, symhdv, tmp18630)

tmp18632 := GCons(__e, symhd, tmp18631)

tmp18633 := GCons(__e, symhead, tmp18632)

tmp18634 := GCons(__e, symhush, tmp18633)

tmp18635 := GCons(__e, symhush_2, tmp18634)

tmp18636 := GCons(__e, symif, tmp18635)

tmp18637 := GCons(__e, symimplementation, tmp18636)

tmp18638 := GCons(__e, syminternal, tmp18637)

tmp18639 := GCons(__e, symin_1package, tmp18638)

tmp18640 := GCons(__e, symin, tmp18639)

tmp18641 := GCons(__e, symis_b, tmp18640)

tmp18642 := GCons(__e, symis, tmp18641)

tmp18643 := GCons(__e, symit, tmp18642)

tmp18644 := GCons(__e, syminclude_1all_1but, tmp18643)

tmp18645 := GCons(__e, syminclude, tmp18644)

tmp18646 := GCons(__e, symincluded, tmp18645)

tmp18647 := GCons(__e, syminline, tmp18646)

tmp18648 := GCons(__e, syminput_7, tmp18647)

tmp18649 := GCons(__e, syminput, tmp18648)

tmp18650 := GCons(__e, syminteger_2, tmp18649)

tmp18651 := GCons(__e, symintern, tmp18650)

tmp18652 := GCons(__e, syminferences, tmp18651)

tmp18653 := GCons(__e, symintersection, tmp18652)

tmp18654 := GCons(__e, symis, tmp18653)

tmp18655 := GCons(__e, symlanguage, tmp18654)

tmp18656 := GCons(__e, symlambda, tmp18655)

tmp18657 := GCons(__e, symlazy, tmp18656)

tmp18658 := GCons(__e, symlet, tmp18657)

tmp18659 := GCons(__e, symlength, tmp18658)

tmp18660 := GCons(__e, symlimit, tmp18659)

tmp18661 := GCons(__e, symlineread, tmp18660)

tmp18662 := GCons(__e, symlist, tmp18661)

tmp18663 := GCons(__e, symloaded, tmp18662)

tmp18664 := GCons(__e, symload, tmp18663)

tmp18665 := GCons(__e, symmake_1string, tmp18664)

tmp18666 := GCons(__e, symmap, tmp18665)

tmp18667 := GCons(__e, symmapcan, tmp18666)

tmp18668 := GCons(__e, symmaxinferences, tmp18667)

tmp18669 := GCons(__e, symmacroexpand, tmp18668)

tmp18670 := GCons(__e, symmode, tmp18669)

tmp18671 := GCons(__e, symnl, tmp18670)

tmp18672 := GCons(__e, symnot, tmp18671)

tmp18673 := GCons(__e, symnth, tmp18672)

tmp18674 := GCons(__e, symnull, tmp18673)

tmp18675 := GCons(__e, symnumber, tmp18674)

tmp18676 := GCons(__e, symnumber_2, tmp18675)

tmp18677 := GCons(__e, symn_1_6string, tmp18676)

tmp18678 := GCons(__e, symoccurs_2, tmp18677)

tmp18679 := GCons(__e, symoccurs_1check, tmp18678)

tmp18680 := GCons(__e, symoccurrences, tmp18679)

tmp18681 := GCons(__e, symopen, tmp18680)

tmp18682 := GCons(__e, symoptimise_2, tmp18681)

tmp18683 := GCons(__e, symoptimise, tmp18682)

tmp18684 := GCons(__e, symor, tmp18683)

tmp18685 := GCons(__e, symos, tmp18684)

tmp18686 := GCons(__e, symout, tmp18685)

tmp18687 := GCons(__e, symoutput, tmp18686)

tmp18688 := GCons(__e, sympackage, tmp18687)

tmp18689 := GCons(__e, symport, tmp18688)

tmp18690 := GCons(__e, symporters, tmp18689)

tmp18691 := GCons(__e, sympos, tmp18690)

tmp18692 := GCons(__e, sympr, tmp18691)

tmp18693 := GCons(__e, symprint, tmp18692)

tmp18694 := GCons(__e, symprolog_1memory, tmp18693)

tmp18695 := GCons(__e, symprofile, tmp18694)

tmp18696 := GCons(__e, symprofile_1results, tmp18695)

tmp18697 := GCons(__e, symprotect, tmp18696)

tmp18698 := GCons(__e, symprolog_2, tmp18697)

tmp18699 := GCons(__e, symps, tmp18698)

tmp18700 := GCons(__e, sympreclude_1all_1but, tmp18699)

tmp18701 := GCons(__e, sympreclude, tmp18700)

tmp18702 := GCons(__e, symput, tmp18701)

tmp18703 := GCons(__e, sympackage_2, tmp18702)

tmp18704 := GCons(__e, symread_1unit_1string, tmp18703)

tmp18705 := GCons(__e, symread_1from_1string_1unprocessed, tmp18704)

tmp18706 := GCons(__e, symread_1from_1string, tmp18705)

tmp18707 := GCons(__e, symread_1byte, tmp18706)

tmp18708 := GCons(__e, symread_1file_1as_1string, tmp18707)

tmp18709 := GCons(__e, symread_1file_1as_1bytelist, tmp18708)

tmp18710 := GCons(__e, symread_1file, tmp18709)

tmp18711 := GCons(__e, symreceive, tmp18710)

tmp18712 := GCons(__e, symread, tmp18711)

tmp18713 := GCons(__e, symrelease, tmp18712)

tmp18714 := GCons(__e, symremove, tmp18713)

tmp18715 := GCons(__e, symretract, tmp18714)

tmp18716 := GCons(__e, symreverse, tmp18715)

tmp18717 := GCons(__e, symrun, tmp18716)

tmp18718 := GCons(__e, symstr, tmp18717)

tmp18719 := GCons(__e, symsave, tmp18718)

tmp18720 := GCons(__e, symset, tmp18719)

tmp18721 := GCons(__e, symsimple_1error, tmp18720)

tmp18722 := GCons(__e, symsnd, tmp18721)

tmp18723 := GCons(__e, symspecialise, tmp18722)

tmp18724 := GCons(__e, symspy, tmp18723)

tmp18725 := GCons(__e, symsqts, tmp18724)

tmp18726 := GCons(__e, symstep, tmp18725)

tmp18727 := GCons(__e, symstoutput, tmp18726)

tmp18728 := GCons(__e, symsterror, tmp18727)

tmp18729 := GCons(__e, symstinput, tmp18728)

tmp18730 := GCons(__e, symstring, tmp18729)

tmp18731 := GCons(__e, symstream, tmp18730)

tmp18732 := GCons(__e, symstring_1_6n, tmp18731)

tmp18733 := GCons(__e, symstring_2, tmp18732)

tmp18734 := GCons(__e, symsubst, tmp18733)

tmp18735 := GCons(__e, symsum, tmp18734)

tmp18736 := GCons(__e, symstring_1_6symbol, tmp18735)

tmp18737 := GCons(__e, symsymbol_2, tmp18736)

tmp18738 := GCons(__e, symsymbol, tmp18737)

tmp18739 := GCons(__e, symsystem_1S_2, tmp18738)

tmp18740 := GCons(__e, symsynonyms, tmp18739)

tmp18741 := GCons(__e, symsystemf, tmp18740)

tmp18742 := GCons(__e, symtail, tmp18741)

tmp18743 := GCons(__e, symtlv, tmp18742)

tmp18744 := GCons(__e, symtlstr, tmp18743)

tmp18745 := GCons(__e, symtl, tmp18744)

tmp18746 := GCons(__e, symtc, tmp18745)

tmp18747 := GCons(__e, symtc_2, tmp18746)

tmp18748 := GCons(__e, symthaw, tmp18747)

tmp18749 := GCons(__e, symtime, tmp18748)

tmp18750 := GCons(__e, symtrack, tmp18749)

tmp18751 := GCons(__e, symtracked, tmp18750)

tmp18752 := GCons(__e, symtrap_1error, tmp18751)

tmp18753 := GCons(__e, True, tmp18752)

tmp18754 := GCons(__e, symtuple_2, tmp18753)

tmp18755 := GCons(__e, symtype, tmp18754)

tmp18756 := GCons(__e, symreturn, tmp18755)

tmp18757 := GCons(__e, symunabsolute, tmp18756)

tmp18758 := GCons(__e, symundefmacro, tmp18757)

tmp18759 := GCons(__e, symunprofile, tmp18758)

tmp18760 := GCons(__e, symunput, tmp18759)

tmp18761 := GCons(__e, symunion, tmp18760)

tmp18762 := GCons(__e, symunix, tmp18761)

tmp18763 := GCons(__e, symunit, tmp18762)

tmp18764 := GCons(__e, symuntrack, tmp18763)

tmp18765 := GCons(__e, symunspecialise, tmp18764)

tmp18766 := GCons(__e, symupdate_1lambda_1table, tmp18765)

tmp18767 := GCons(__e, symu_b, tmp18766)

tmp18768 := GCons(__e, symuserdefs, tmp18767)

tmp18769 := GCons(__e, symvector_2, tmp18768)

tmp18770 := GCons(__e, symvector, tmp18769)

tmp18771 := GCons(__e, sym_5_1vector, tmp18770)

tmp18772 := GCons(__e, symvector_1_6, tmp18771)

tmp18773 := GCons(__e, symvalue, tmp18772)

tmp18774 := GCons(__e, symvar_2, tmp18773)

tmp18775 := GCons(__e, symvariable_2, tmp18774)

tmp18776 := GCons(__e, symverified, tmp18775)

tmp18777 := GCons(__e, symversion, tmp18776)

tmp18778 := GCons(__e, symwarn, tmp18777)

tmp18779 := GCons(__e, symwhen, tmp18778)

tmp18780 := GCons(__e, symwhere, tmp18779)

tmp18781 := GCons(__e, symwrite_1byte, tmp18780)

tmp18782 := GCons(__e, symwrite_1to_1file, tmp18781)

tmp18783 := GCons(__e, symwrite_1string, tmp18782)

tmp18784 := GCons(__e, symy_1or_1n_2, tmp18783)

tmp18785 := GCons(__e, tmp18554, tmp18784)

tmp18786 := GCons(__e, sym_6_6, tmp18785)

tmp18787 := GCons(__e, sym_5, tmp18786)

tmp18788 := GCons(__e, sym_5_a, tmp18787)

tmp18789 := GCons(__e, sym_7, tmp18788)

tmp18790 := GCons(__e, sym_d, tmp18789)

tmp18791 := GCons(__e, sym_c, tmp18790)

tmp18792 := GCons(__e, sym_1, tmp18791)

tmp18793 := GCons(__e, sym_3, tmp18792)

tmp18794 := GCons(__e, sym_5end_6, tmp18793)

tmp18795 := GCons(__e, sym_5_b_6, tmp18794)

tmp18796 := GCons(__e, sym_c_4, tmp18795)

tmp18797 := GCons(__e, sym_a_a_6, tmp18796)

tmp18798 := GCons(__e, sym_6, tmp18797)

tmp18799 := GCons(__e, sym_6_a, tmp18798)

tmp18800 := GCons(__e, sym_a, tmp18799)

tmp18801 := GCons(__e, sym_a_a, tmp18800)

tmp18802 := GCons(__e, sym_5e_6, tmp18801)

tmp18803 := GCons(__e, sym_1_6, tmp18802)

tmp18804 := GCons(__e, sym_5_1, tmp18803)

tmp18805 := GCons(__e, sym_dhush_d, tmp18804)

tmp18806 := GCons(__e, sym_dporters_d, tmp18805)

tmp18807 := GCons(__e, sym_dport_d, tmp18806)

tmp18808 := GCons(__e, sym_8s, tmp18807)

tmp18809 := GCons(__e, sym_8p, tmp18808)

tmp18810 := GCons(__e, sym_8v, tmp18809)

tmp18811 := GCons(__e, sym_dproperty_1vector_d, tmp18810)

tmp18812 := GCons(__e, sym_drelease_d, tmp18811)

tmp18813 := GCons(__e, sym_dos_d, tmp18812)

tmp18814 := GCons(__e, sym_dmacros_d, tmp18813)

tmp18815 := GCons(__e, sym_dmaximum_1print_1sequence_1size_d, tmp18814)

tmp18816 := GCons(__e, sym_dversion_d, tmp18815)

tmp18817 := GCons(__e, sym_dhome_1directory_d, tmp18816)

tmp18818 := GCons(__e, sym_dstoutput_d, tmp18817)

tmp18819 := GCons(__e, sym_dsterror_d, tmp18818)

tmp18820 := GCons(__e, sym_dstinput_d, tmp18819)

tmp18821 := GCons(__e, sym_dimplementation_d, tmp18820)

tmp18822 := GCons(__e, sym_dlanguage_d, tmp18821)

tmp18823 := GCons(__e, sym__, tmp18822)

tmp18824 := GCons(__e, tmp18553, tmp18823)

tmp18825 := GCons(__e, tmp18552, tmp18824)

tmp18826 := GCons(__e, tmp18551, tmp18825)

tmp18827 := GCons(__e, tmp18550, tmp18826)

tmp18828 := GCons(__e, sym_e_e, tmp18827)

tmp18829 := GCons(__e, sym_5_1_1, tmp18828)

tmp18830 := GCons(__e, sym_1_1_6, tmp18829)

tmp18831 := GCons(__e, sym_i, tmp18830)

tmp18832 := GCons(__e, sym_j, tmp18831)

tmp18833 := GCons(__e, sym_b, tmp18832)

tmp18834 := GValue(__e, sym_dproperty_1vector_d)

tmp18835 := Call4(__e, SymFn(symput), symshen, symshen_4external_1symbols, tmp18833, tmp18834)


_ = tmp18835

tmp18836 := PrimGuard1(__e, symabsvector, PrimAbsvector, MakeInteger(0))

return PrimGuard2(__e, symset, PrimSet, symshen_4_dempty_1absvector_d, tmp18836)


panic("unreachable")
})

tmp18837 := Call2(__e, ns2_1set, symshen_4initialise_1environment, tmp18073)


_ = tmp18837

tmp18838 := MakeNative0(func(__e *ControlFlow) Obj {
tmp18839 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, Nil)

_ = tmp18839

tmp18840 := MakeNative1(func(__e *ControlFlow, V6073 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6069 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6070 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6071 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6072 Obj) Obj {
tmp18841 := Call1(__e, SymFn(symshen_4newpv), B6069)


let__17986 := tmp18841
_ = let__17986

tmp18842 := Call1(__e, SymFn(symshen_4newpv), B6069)


let__17987 := tmp18842
_ = let__17987

tmp18843 := GCons(__e, let__17987, Nil)

tmp18844 := GCons(__e, symlist, tmp18843)

tmp18845 := GCons(__e, tmp18844, Nil)

tmp18846 := GCons(__e, sym_1_1_6, tmp18845)

tmp18847 := GCons(__e, let__17986, tmp18846)

tmp18848 := GCons(__e, symsymbol, Nil)

tmp18849 := GCons(__e, sym_1_1_6, tmp18848)

tmp18850 := GCons(__e, tmp18847, tmp18849)

tmp18851 := GCons(__e, tmp18850, Nil)

tmp18852 := GCons(__e, sym_1_1_6, tmp18851)

tmp18853 := GCons(__e, symsymbol, tmp18852)

tmp18854 := Call6(__e, SymFn(symis_b), V6073, tmp18853, B6069, L6070, Key6071, C6072)


tmp18855 := Call2(__e, SymFn(symshen_4gc), B6069, tmp18854)


return TailCall2(__e, SymFn(symshen_4gc), B6069, tmp18855)


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

tmp18856 := GValue(__e, symshen_4_dsigf_d)

tmp18857 := Call3(__e, SymFn(symshen_4assoc_1_6), symshen_4register_1source_1form, tmp18840, tmp18856)


tmp18858 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp18857)

_ = tmp18858

tmp18859 := MakeNative1(func(__e *ControlFlow, V6078 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6074 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6075 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6076 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6077 Obj) Obj {
tmp18860 := GCons(__e, symsymbol, Nil)

tmp18861 := GCons(__e, sym_1_1_6, tmp18860)

tmp18862 := GCons(__e, symsymbol, tmp18861)

return TailCall6(__e, SymFn(symis_b), V6078, tmp18862, B6074, L6075, Key6076, C6077)


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

tmp18863 := GValue(__e, symshen_4_dsigf_d)

tmp18864 := Call3(__e, SymFn(symshen_4assoc_1_6), symshen_4unregister_1source_1form, tmp18859, tmp18863)


tmp18865 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp18864)

_ = tmp18865

tmp18866 := MakeNative1(func(__e *ControlFlow, V6083 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6079 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6080 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6081 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6082 Obj) Obj {
tmp18867 := Call1(__e, SymFn(symshen_4newpv), B6079)


let__17988 := tmp18867
_ = let__17988

tmp18868 := GCons(__e, let__17988, Nil)

tmp18869 := GCons(__e, sym_1_1_6, tmp18868)

tmp18870 := Call6(__e, SymFn(symis_b), V6083, tmp18869, B6079, L6080, Key6081, C6082)


return TailCall2(__e, SymFn(symshen_4gc), B6079, tmp18870)


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

tmp18871 := GValue(__e, symshen_4_dsigf_d)

tmp18872 := Call3(__e, SymFn(symshen_4assoc_1_6), symabort, tmp18866, tmp18871)


tmp18873 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp18872)

_ = tmp18873

tmp18874 := MakeNative1(func(__e *ControlFlow, V6088 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6084 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6085 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6086 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6087 Obj) Obj {
tmp18875 := GCons(__e, symstring, Nil)

tmp18876 := GCons(__e, symlist, tmp18875)

tmp18877 := GCons(__e, tmp18876, Nil)

tmp18878 := GCons(__e, sym_1_1_6, tmp18877)

tmp18879 := GCons(__e, symstring, tmp18878)

return TailCall6(__e, SymFn(symis_b), V6088, tmp18879, B6084, L6085, Key6086, C6087)


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

tmp18880 := GValue(__e, symshen_4_dsigf_d)

tmp18881 := Call3(__e, SymFn(symshen_4assoc_1_6), symabsolute, tmp18874, tmp18880)


tmp18882 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp18881)

_ = tmp18882

tmp18883 := MakeNative1(func(__e *ControlFlow, V6093 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6089 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6090 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6091 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6092 Obj) Obj {
tmp18884 := Call1(__e, SymFn(symshen_4newpv), B6089)


let__17989 := tmp18884
_ = let__17989

tmp18885 := GCons(__e, symboolean, Nil)

tmp18886 := GCons(__e, sym_1_1_6, tmp18885)

tmp18887 := GCons(__e, let__17989, tmp18886)

tmp18888 := Call6(__e, SymFn(symis_b), V6093, tmp18887, B6089, L6090, Key6091, C6092)


return TailCall2(__e, SymFn(symshen_4gc), B6089, tmp18888)


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

tmp18889 := GValue(__e, symshen_4_dsigf_d)

tmp18890 := Call3(__e, SymFn(symshen_4assoc_1_6), symabsvector_2, tmp18883, tmp18889)


tmp18891 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp18890)

_ = tmp18891

tmp18892 := MakeNative1(func(__e *ControlFlow, V6099 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6095 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6096 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6097 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6098 Obj) Obj {
tmp18893 := Call1(__e, SymFn(symshen_4newpv), B6095)


let__17990 := tmp18893
_ = let__17990

tmp18894 := GCons(__e, let__17990, Nil)

tmp18895 := GCons(__e, symlist, tmp18894)

tmp18896 := GCons(__e, let__17990, Nil)

tmp18897 := GCons(__e, symlist, tmp18896)

tmp18898 := GCons(__e, tmp18897, Nil)

tmp18899 := GCons(__e, sym_1_1_6, tmp18898)

tmp18900 := GCons(__e, tmp18895, tmp18899)

tmp18901 := GCons(__e, tmp18900, Nil)

tmp18902 := GCons(__e, sym_1_1_6, tmp18901)

tmp18903 := GCons(__e, let__17990, tmp18902)

tmp18904 := Call6(__e, SymFn(symis_b), V6099, tmp18903, B6095, L6096, Key6097, C6098)


return TailCall2(__e, SymFn(symshen_4gc), B6095, tmp18904)


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

tmp18905 := GValue(__e, symshen_4_dsigf_d)

tmp18906 := Call3(__e, SymFn(symshen_4assoc_1_6), symadjoin, tmp18892, tmp18905)


tmp18907 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp18906)

_ = tmp18907

tmp18908 := MakeNative1(func(__e *ControlFlow, V6104 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6100 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6101 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6102 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6103 Obj) Obj {
tmp18909 := GCons(__e, symboolean, Nil)

tmp18910 := GCons(__e, sym_1_1_6, tmp18909)

tmp18911 := GCons(__e, symboolean, tmp18910)

tmp18912 := GCons(__e, tmp18911, Nil)

tmp18913 := GCons(__e, sym_1_1_6, tmp18912)

tmp18914 := GCons(__e, symboolean, tmp18913)

return TailCall6(__e, SymFn(symis_b), V6104, tmp18914, B6100, L6101, Key6102, C6103)


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

tmp18915 := GValue(__e, symshen_4_dsigf_d)

tmp18916 := Call3(__e, SymFn(symshen_4assoc_1_6), symand, tmp18908, tmp18915)


tmp18917 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp18916)

_ = tmp18917

tmp18918 := MakeNative1(func(__e *ControlFlow, V6109 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6105 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6106 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6107 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6108 Obj) Obj {
tmp18919 := Call1(__e, SymFn(symshen_4newpv), B6105)


let__17991 := tmp18919
_ = let__17991

tmp18920 := GCons(__e, symstring, Nil)

tmp18921 := GCons(__e, sym_1_1_6, tmp18920)

tmp18922 := GCons(__e, symsymbol, tmp18921)

tmp18923 := GCons(__e, tmp18922, Nil)

tmp18924 := GCons(__e, sym_1_1_6, tmp18923)

tmp18925 := GCons(__e, symstring, tmp18924)

tmp18926 := GCons(__e, tmp18925, Nil)

tmp18927 := GCons(__e, sym_1_1_6, tmp18926)

tmp18928 := GCons(__e, let__17991, tmp18927)

tmp18929 := Call6(__e, SymFn(symis_b), V6109, tmp18928, B6105, L6106, Key6107, C6108)


return TailCall2(__e, SymFn(symshen_4gc), B6105, tmp18929)


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

tmp18930 := GValue(__e, symshen_4_dsigf_d)

tmp18931 := Call3(__e, SymFn(symshen_4assoc_1_6), symshen_4app, tmp18918, tmp18930)


tmp18932 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp18931)

_ = tmp18932

tmp18933 := MakeNative1(func(__e *ControlFlow, V6116 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6112 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6113 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6114 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6115 Obj) Obj {
tmp18934 := Call1(__e, SymFn(symshen_4newpv), B6112)


let__17992 := tmp18934
_ = let__17992

tmp18935 := GCons(__e, let__17992, Nil)

tmp18936 := GCons(__e, symlist, tmp18935)

tmp18937 := GCons(__e, let__17992, Nil)

tmp18938 := GCons(__e, symlist, tmp18937)

tmp18939 := GCons(__e, let__17992, Nil)

tmp18940 := GCons(__e, symlist, tmp18939)

tmp18941 := GCons(__e, tmp18940, Nil)

tmp18942 := GCons(__e, sym_1_1_6, tmp18941)

tmp18943 := GCons(__e, tmp18938, tmp18942)

tmp18944 := GCons(__e, tmp18943, Nil)

tmp18945 := GCons(__e, sym_1_1_6, tmp18944)

tmp18946 := GCons(__e, tmp18936, tmp18945)

tmp18947 := Call6(__e, SymFn(symis_b), V6116, tmp18946, B6112, L6113, Key6114, C6115)


return TailCall2(__e, SymFn(symshen_4gc), B6112, tmp18947)


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

tmp18948 := GValue(__e, symshen_4_dsigf_d)

tmp18949 := Call3(__e, SymFn(symshen_4assoc_1_6), symappend, tmp18933, tmp18948)


tmp18950 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp18949)

_ = tmp18950

tmp18951 := MakeNative1(func(__e *ControlFlow, V6121 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6117 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6118 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6119 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6120 Obj) Obj {
tmp18952 := Call1(__e, SymFn(symshen_4newpv), B6117)


let__17993 := tmp18952
_ = let__17993

tmp18953 := GCons(__e, symnumber, Nil)

tmp18954 := GCons(__e, sym_1_1_6, tmp18953)

tmp18955 := GCons(__e, let__17993, tmp18954)

tmp18956 := Call6(__e, SymFn(symis_b), V6121, tmp18955, B6117, L6118, Key6119, C6120)


return TailCall2(__e, SymFn(symshen_4gc), B6117, tmp18956)


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

tmp18957 := GValue(__e, symshen_4_dsigf_d)

tmp18958 := Call3(__e, SymFn(symshen_4assoc_1_6), symarity, tmp18951, tmp18957)


tmp18959 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp18958)

_ = tmp18959

tmp18960 := MakeNative1(func(__e *ControlFlow, V6127 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6123 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6124 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6125 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6126 Obj) Obj {
tmp18961 := Call1(__e, SymFn(symshen_4newpv), B6123)


let__17994 := tmp18961
_ = let__17994

tmp18962 := GCons(__e, let__17994, Nil)

tmp18963 := GCons(__e, symlist, tmp18962)

tmp18964 := GCons(__e, tmp18963, Nil)

tmp18965 := GCons(__e, symlist, tmp18964)

tmp18966 := GCons(__e, let__17994, Nil)

tmp18967 := GCons(__e, symlist, tmp18966)

tmp18968 := GCons(__e, tmp18967, Nil)

tmp18969 := GCons(__e, sym_1_1_6, tmp18968)

tmp18970 := GCons(__e, tmp18965, tmp18969)

tmp18971 := GCons(__e, tmp18970, Nil)

tmp18972 := GCons(__e, sym_1_1_6, tmp18971)

tmp18973 := GCons(__e, let__17994, tmp18972)

tmp18974 := Call6(__e, SymFn(symis_b), V6127, tmp18973, B6123, L6124, Key6125, C6126)


return TailCall2(__e, SymFn(symshen_4gc), B6123, tmp18974)


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

tmp18975 := GValue(__e, symshen_4_dsigf_d)

tmp18976 := Call3(__e, SymFn(symshen_4assoc_1_6), symassoc, tmp18960, tmp18975)


tmp18977 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp18976)

_ = tmp18977

tmp18978 := MakeNative1(func(__e *ControlFlow, V6132 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6128 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6129 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6130 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6131 Obj) Obj {
tmp18979 := Call1(__e, SymFn(symshen_4newpv), B6128)


let__17995 := tmp18979
_ = let__17995

tmp18980 := GCons(__e, symboolean, Nil)

tmp18981 := GCons(__e, sym_1_1_6, tmp18980)

tmp18982 := GCons(__e, let__17995, tmp18981)

tmp18983 := Call6(__e, SymFn(symis_b), V6132, tmp18982, B6128, L6129, Key6130, C6131)


return TailCall2(__e, SymFn(symshen_4gc), B6128, tmp18983)


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

tmp18984 := GValue(__e, symshen_4_dsigf_d)

tmp18985 := Call3(__e, SymFn(symshen_4assoc_1_6), symatom_2, tmp18978, tmp18984)


tmp18986 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp18985)

_ = tmp18986

tmp18987 := MakeNative1(func(__e *ControlFlow, V6137 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6133 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6134 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6135 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6136 Obj) Obj {
tmp18988 := Call1(__e, SymFn(symshen_4newpv), B6133)


let__17996 := tmp18988
_ = let__17996

tmp18989 := GCons(__e, symboolean, Nil)

tmp18990 := GCons(__e, sym_1_1_6, tmp18989)

tmp18991 := GCons(__e, let__17996, tmp18990)

tmp18992 := Call6(__e, SymFn(symis_b), V6137, tmp18991, B6133, L6134, Key6135, C6136)


return TailCall2(__e, SymFn(symshen_4gc), B6133, tmp18992)


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

tmp18993 := GValue(__e, symshen_4_dsigf_d)

tmp18994 := Call3(__e, SymFn(symshen_4assoc_1_6), symboolean_2, tmp18987, tmp18993)


tmp18995 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp18994)

_ = tmp18995

tmp18996 := MakeNative1(func(__e *ControlFlow, V6142 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6138 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6139 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6140 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6141 Obj) Obj {
tmp18997 := GCons(__e, symstring, Nil)

tmp18998 := GCons(__e, sym_1_1_6, tmp18997)

tmp18999 := GCons(__e, symstring, tmp18998)

return TailCall6(__e, SymFn(symis_b), V6142, tmp18999, B6138, L6139, Key6140, C6141)


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

tmp19000 := GValue(__e, symshen_4_dsigf_d)

tmp19001 := Call3(__e, SymFn(symshen_4assoc_1_6), symbootstrap, tmp18996, tmp19000)


tmp19002 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19001)

_ = tmp19002

tmp19003 := MakeNative1(func(__e *ControlFlow, V6147 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6143 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6144 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6145 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6146 Obj) Obj {
tmp19004 := GCons(__e, symboolean, Nil)

tmp19005 := GCons(__e, sym_1_1_6, tmp19004)

tmp19006 := GCons(__e, symsymbol, tmp19005)

return TailCall6(__e, SymFn(symis_b), V6147, tmp19006, B6143, L6144, Key6145, C6146)


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

tmp19007 := GValue(__e, symshen_4_dsigf_d)

tmp19008 := Call3(__e, SymFn(symshen_4assoc_1_6), symbound_2, tmp19003, tmp19007)


tmp19009 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19008)

_ = tmp19009

tmp19010 := MakeNative1(func(__e *ControlFlow, V6152 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6148 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6149 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6150 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6151 Obj) Obj {
tmp19011 := Call1(__e, SymFn(symshen_4newpv), B6148)


let__17997 := tmp19011
_ = let__17997

tmp19012 := GCons(__e, let__17997, Nil)

tmp19013 := GCons(__e, symlist, tmp19012)

tmp19014 := GCons(__e, symboolean, Nil)

tmp19015 := GCons(__e, sym_1_1_6, tmp19014)

tmp19016 := GCons(__e, tmp19013, tmp19015)

tmp19017 := Call6(__e, SymFn(symis_b), V6152, tmp19016, B6148, L6149, Key6150, C6151)


return TailCall2(__e, SymFn(symshen_4gc), B6148, tmp19017)


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

tmp19018 := GValue(__e, symshen_4_dsigf_d)

tmp19019 := Call3(__e, SymFn(symshen_4assoc_1_6), symshen_4ccons_2, tmp19010, tmp19018)


tmp19020 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19019)

_ = tmp19020

tmp19021 := MakeNative1(func(__e *ControlFlow, V6157 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6153 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6154 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6155 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6156 Obj) Obj {
tmp19022 := GCons(__e, symstring, Nil)

tmp19023 := GCons(__e, sym_1_1_6, tmp19022)

tmp19024 := GCons(__e, symstring, tmp19023)

return TailCall6(__e, SymFn(symis_b), V6157, tmp19024, B6153, L6154, Key6155, C6156)


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

tmp19025 := GValue(__e, symshen_4_dsigf_d)

tmp19026 := Call3(__e, SymFn(symshen_4assoc_1_6), symcd, tmp19021, tmp19025)


tmp19027 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19026)

_ = tmp19027

tmp19028 := MakeNative1(func(__e *ControlFlow, V6163 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6159 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6160 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6161 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6162 Obj) Obj {
tmp19029 := Call1(__e, SymFn(symshen_4newpv), B6159)


let__17998 := tmp19029
_ = let__17998

tmp19030 := Call1(__e, SymFn(symshen_4newpv), B6159)


let__17999 := tmp19030
_ = let__17999

tmp19031 := GCons(__e, let__17998, Nil)

tmp19032 := GCons(__e, symstream, tmp19031)

tmp19033 := GCons(__e, let__17999, Nil)

tmp19034 := GCons(__e, symlist, tmp19033)

tmp19035 := GCons(__e, tmp19034, Nil)

tmp19036 := GCons(__e, sym_1_1_6, tmp19035)

tmp19037 := GCons(__e, tmp19032, tmp19036)

tmp19038 := Call6(__e, SymFn(symis_b), V6163, tmp19037, B6159, L6160, Key6161, C6162)


tmp19039 := Call2(__e, SymFn(symshen_4gc), B6159, tmp19038)


return TailCall2(__e, SymFn(symshen_4gc), B6159, tmp19039)


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

tmp19040 := GValue(__e, symshen_4_dsigf_d)

tmp19041 := Call3(__e, SymFn(symshen_4assoc_1_6), symclose, tmp19028, tmp19040)


tmp19042 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19041)

_ = tmp19042

tmp19043 := MakeNative1(func(__e *ControlFlow, V6168 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6164 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6165 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6166 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6167 Obj) Obj {
tmp19044 := GCons(__e, symstring, Nil)

tmp19045 := GCons(__e, sym_1_1_6, tmp19044)

tmp19046 := GCons(__e, symstring, tmp19045)

tmp19047 := GCons(__e, tmp19046, Nil)

tmp19048 := GCons(__e, sym_1_1_6, tmp19047)

tmp19049 := GCons(__e, symstring, tmp19048)

return TailCall6(__e, SymFn(symis_b), V6168, tmp19049, B6164, L6165, Key6166, C6167)


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

tmp19050 := GValue(__e, symshen_4_dsigf_d)

tmp19051 := Call3(__e, SymFn(symshen_4assoc_1_6), symcn, tmp19043, tmp19050)


tmp19052 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19051)

_ = tmp19052

tmp19053 := MakeNative1(func(__e *ControlFlow, V6173 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6169 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6170 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6171 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6172 Obj) Obj {
tmp19054 := Call1(__e, SymFn(symshen_4newpv), B6169)


let__18000 := tmp19054
_ = let__18000

tmp19055 := GCons(__e, let__18000, Nil)

tmp19056 := GCons(__e, symstream, tmp19055)

tmp19057 := GCons(__e, symboolean, Nil)

tmp19058 := GCons(__e, sym_1_1_6, tmp19057)

tmp19059 := GCons(__e, tmp19056, tmp19058)

tmp19060 := Call6(__e, SymFn(symis_b), V6173, tmp19059, B6169, L6170, Key6171, C6172)


return TailCall2(__e, SymFn(symshen_4gc), B6169, tmp19060)


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

tmp19061 := GValue(__e, symshen_4_dsigf_d)

tmp19062 := Call3(__e, SymFn(symshen_4assoc_1_6), symchar_1stream_2, tmp19053, tmp19061)


tmp19063 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19062)

_ = tmp19063

tmp19064 := MakeNative1(func(__e *ControlFlow, V6180 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6176 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6177 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6178 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6179 Obj) Obj {
tmp19065 := Call1(__e, SymFn(symshen_4newpv), B6176)


let__18001 := tmp19065
_ = let__18001

tmp19066 := Call1(__e, SymFn(symshen_4newpv), B6176)


let__18002 := tmp19066
_ = let__18002

tmp19067 := GCons(__e, let__18001, Nil)

tmp19068 := GCons(__e, symlist, tmp19067)

tmp19069 := GCons(__e, let__18001, Nil)

tmp19070 := GCons(__e, symlist, tmp19069)

tmp19071 := GCons(__e, let__18002, Nil)

tmp19072 := GCons(__e, tmp19070, tmp19071)

tmp19073 := GCons(__e, symstr, tmp19072)

tmp19074 := GCons(__e, tmp19073, Nil)

tmp19075 := GCons(__e, sym_1_1_6, tmp19074)

tmp19076 := GCons(__e, tmp19068, tmp19075)

tmp19077 := GCons(__e, let__18001, Nil)

tmp19078 := GCons(__e, symlist, tmp19077)

tmp19079 := GCons(__e, let__18002, Nil)

tmp19080 := GCons(__e, sym_1_1_6, tmp19079)

tmp19081 := GCons(__e, tmp19078, tmp19080)

tmp19082 := GCons(__e, tmp19081, Nil)

tmp19083 := GCons(__e, sym_1_1_6, tmp19082)

tmp19084 := GCons(__e, tmp19076, tmp19083)

tmp19085 := Call6(__e, SymFn(symis_b), V6180, tmp19084, B6176, L6177, Key6178, C6179)


tmp19086 := Call2(__e, SymFn(symshen_4gc), B6176, tmp19085)


return TailCall2(__e, SymFn(symshen_4gc), B6176, tmp19086)


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

tmp19087 := GValue(__e, symshen_4_dsigf_d)

tmp19088 := Call3(__e, SymFn(symshen_4assoc_1_6), symcompile, tmp19064, tmp19087)


tmp19089 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19088)

_ = tmp19089

tmp19090 := MakeNative1(func(__e *ControlFlow, V6185 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6181 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6182 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6183 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6184 Obj) Obj {
tmp19091 := Call1(__e, SymFn(symshen_4newpv), B6181)


let__18003 := tmp19091
_ = let__18003

tmp19092 := GCons(__e, symboolean, Nil)

tmp19093 := GCons(__e, sym_1_1_6, tmp19092)

tmp19094 := GCons(__e, let__18003, tmp19093)

tmp19095 := Call6(__e, SymFn(symis_b), V6185, tmp19094, B6181, L6182, Key6183, C6184)


return TailCall2(__e, SymFn(symshen_4gc), B6181, tmp19095)


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

tmp19096 := GValue(__e, symshen_4_dsigf_d)

tmp19097 := Call3(__e, SymFn(symshen_4assoc_1_6), symcons_2, tmp19090, tmp19096)


tmp19098 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19097)

_ = tmp19098

tmp19099 := MakeNative1(func(__e *ControlFlow, V6190 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6186 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6187 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6188 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6189 Obj) Obj {
tmp19100 := GCons(__e, symsymbol, Nil)

tmp19101 := GCons(__e, symlist, tmp19100)

tmp19102 := GCons(__e, tmp19101, Nil)

tmp19103 := GCons(__e, sym_1_1_6, tmp19102)

return TailCall6(__e, SymFn(symis_b), V6190, tmp19103, B6186, L6187, Key6188, C6189)


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

tmp19104 := GValue(__e, symshen_4_dsigf_d)

tmp19105 := Call3(__e, SymFn(symshen_4assoc_1_6), symdatatypes, tmp19099, tmp19104)


tmp19106 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19105)

_ = tmp19106

tmp19107 := MakeNative1(func(__e *ControlFlow, V6195 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6191 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6192 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6193 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6194 Obj) Obj {
tmp19108 := GCons(__e, symsymbol, Nil)

tmp19109 := GCons(__e, sym_1_1_6, tmp19108)

tmp19110 := GCons(__e, symsymbol, tmp19109)

return TailCall6(__e, SymFn(symis_b), V6195, tmp19110, B6191, L6192, Key6193, C6194)


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

tmp19111 := GValue(__e, symshen_4_dsigf_d)

tmp19112 := Call3(__e, SymFn(symshen_4assoc_1_6), symdestroy, tmp19107, tmp19111)


tmp19113 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19112)

_ = tmp19113

tmp19114 := MakeNative1(func(__e *ControlFlow, V6202 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6198 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6199 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6200 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6201 Obj) Obj {
tmp19115 := Call1(__e, SymFn(symshen_4newpv), B6198)


let__18004 := tmp19115
_ = let__18004

tmp19116 := GCons(__e, let__18004, Nil)

tmp19117 := GCons(__e, symlist, tmp19116)

tmp19118 := GCons(__e, let__18004, Nil)

tmp19119 := GCons(__e, symlist, tmp19118)

tmp19120 := GCons(__e, let__18004, Nil)

tmp19121 := GCons(__e, symlist, tmp19120)

tmp19122 := GCons(__e, tmp19121, Nil)

tmp19123 := GCons(__e, sym_1_1_6, tmp19122)

tmp19124 := GCons(__e, tmp19119, tmp19123)

tmp19125 := GCons(__e, tmp19124, Nil)

tmp19126 := GCons(__e, sym_1_1_6, tmp19125)

tmp19127 := GCons(__e, tmp19117, tmp19126)

tmp19128 := Call6(__e, SymFn(symis_b), V6202, tmp19127, B6198, L6199, Key6200, C6201)


return TailCall2(__e, SymFn(symshen_4gc), B6198, tmp19128)


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

tmp19129 := GValue(__e, symshen_4_dsigf_d)

tmp19130 := Call3(__e, SymFn(symshen_4assoc_1_6), symdifference, tmp19114, tmp19129)


tmp19131 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19130)

_ = tmp19131

tmp19132 := MakeNative1(func(__e *ControlFlow, V6207 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6203 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6204 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6205 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6206 Obj) Obj {
tmp19133 := Call1(__e, SymFn(symshen_4newpv), B6203)


let__18005 := tmp19133
_ = let__18005

tmp19134 := Call1(__e, SymFn(symshen_4newpv), B6203)


let__18006 := tmp19134
_ = let__18006

tmp19135 := GCons(__e, let__18006, Nil)

tmp19136 := GCons(__e, sym_1_1_6, tmp19135)

tmp19137 := GCons(__e, let__18006, tmp19136)

tmp19138 := GCons(__e, tmp19137, Nil)

tmp19139 := GCons(__e, sym_1_1_6, tmp19138)

tmp19140 := GCons(__e, let__18005, tmp19139)

tmp19141 := Call6(__e, SymFn(symis_b), V6207, tmp19140, B6203, L6204, Key6205, C6206)


tmp19142 := Call2(__e, SymFn(symshen_4gc), B6203, tmp19141)


return TailCall2(__e, SymFn(symshen_4gc), B6203, tmp19142)


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

tmp19143 := GValue(__e, symshen_4_dsigf_d)

tmp19144 := Call3(__e, SymFn(symshen_4assoc_1_6), symdo, tmp19132, tmp19143)


tmp19145 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19144)

_ = tmp19145

tmp19146 := MakeNative1(func(__e *ControlFlow, V6214 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6210 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6211 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6212 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6213 Obj) Obj {
tmp19147 := Call1(__e, SymFn(symshen_4newpv), B6210)


let__18007 := tmp19147
_ = let__18007

tmp19148 := Call1(__e, SymFn(symshen_4newpv), B6210)


let__18008 := tmp19148
_ = let__18008

tmp19149 := GCons(__e, let__18007, Nil)

tmp19150 := GCons(__e, symlist, tmp19149)

tmp19151 := GCons(__e, let__18007, Nil)

tmp19152 := GCons(__e, symlist, tmp19151)

tmp19153 := GCons(__e, let__18008, Nil)

tmp19154 := GCons(__e, symlist, tmp19153)

tmp19155 := GCons(__e, tmp19154, Nil)

tmp19156 := GCons(__e, tmp19152, tmp19155)

tmp19157 := GCons(__e, symstr, tmp19156)

tmp19158 := GCons(__e, tmp19157, Nil)

tmp19159 := GCons(__e, sym_1_1_6, tmp19158)

tmp19160 := GCons(__e, tmp19150, tmp19159)

tmp19161 := Call6(__e, SymFn(symis_b), V6214, tmp19160, B6210, L6211, Key6212, C6213)


tmp19162 := Call2(__e, SymFn(symshen_4gc), B6210, tmp19161)


return TailCall2(__e, SymFn(symshen_4gc), B6210, tmp19162)


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

tmp19163 := GValue(__e, symshen_4_dsigf_d)

tmp19164 := Call3(__e, SymFn(symshen_4assoc_1_6), sym_5e_6, tmp19146, tmp19163)


tmp19165 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19164)

_ = tmp19165

tmp19166 := MakeNative1(func(__e *ControlFlow, V6221 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6217 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6218 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6219 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6220 Obj) Obj {
tmp19167 := Call1(__e, SymFn(symshen_4newpv), B6217)


let__18009 := tmp19167
_ = let__18009

tmp19168 := Call1(__e, SymFn(symshen_4newpv), B6217)


let__18010 := tmp19168
_ = let__18010

tmp19169 := GCons(__e, let__18010, Nil)

tmp19170 := GCons(__e, symlist, tmp19169)

tmp19171 := GCons(__e, let__18009, Nil)

tmp19172 := GCons(__e, symlist, tmp19171)

tmp19173 := GCons(__e, let__18010, Nil)

tmp19174 := GCons(__e, symlist, tmp19173)

tmp19175 := GCons(__e, tmp19174, Nil)

tmp19176 := GCons(__e, tmp19172, tmp19175)

tmp19177 := GCons(__e, symstr, tmp19176)

tmp19178 := GCons(__e, tmp19177, Nil)

tmp19179 := GCons(__e, sym_1_1_6, tmp19178)

tmp19180 := GCons(__e, tmp19170, tmp19179)

tmp19181 := Call6(__e, SymFn(symis_b), V6221, tmp19180, B6217, L6218, Key6219, C6220)


tmp19182 := Call2(__e, SymFn(symshen_4gc), B6217, tmp19181)


return TailCall2(__e, SymFn(symshen_4gc), B6217, tmp19182)


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

tmp19183 := GValue(__e, symshen_4_dsigf_d)

tmp19184 := Call3(__e, SymFn(symshen_4assoc_1_6), sym_5_b_6, tmp19166, tmp19183)


tmp19185 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19184)

_ = tmp19185

tmp19186 := MakeNative1(func(__e *ControlFlow, V6228 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6224 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6225 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6226 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6227 Obj) Obj {
tmp19187 := Call1(__e, SymFn(symshen_4newpv), B6224)


let__18011 := tmp19187
_ = let__18011

tmp19188 := Call1(__e, SymFn(symshen_4newpv), B6224)


let__18012 := tmp19188
_ = let__18012

tmp19189 := GCons(__e, let__18011, Nil)

tmp19190 := GCons(__e, symlist, tmp19189)

tmp19191 := GCons(__e, let__18011, Nil)

tmp19192 := GCons(__e, symlist, tmp19191)

tmp19193 := GCons(__e, let__18012, Nil)

tmp19194 := GCons(__e, symlist, tmp19193)

tmp19195 := GCons(__e, tmp19194, Nil)

tmp19196 := GCons(__e, tmp19192, tmp19195)

tmp19197 := GCons(__e, symstr, tmp19196)

tmp19198 := GCons(__e, tmp19197, Nil)

tmp19199 := GCons(__e, sym_1_1_6, tmp19198)

tmp19200 := GCons(__e, tmp19190, tmp19199)

tmp19201 := Call6(__e, SymFn(symis_b), V6228, tmp19200, B6224, L6225, Key6226, C6227)


tmp19202 := Call2(__e, SymFn(symshen_4gc), B6224, tmp19201)


return TailCall2(__e, SymFn(symshen_4gc), B6224, tmp19202)


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

tmp19203 := GValue(__e, symshen_4_dsigf_d)

tmp19204 := Call3(__e, SymFn(symshen_4assoc_1_6), sym_5end_6, tmp19186, tmp19203)


tmp19205 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19204)

_ = tmp19205

tmp19206 := MakeNative1(func(__e *ControlFlow, V6233 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6229 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6230 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6231 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6232 Obj) Obj {
tmp19207 := Call1(__e, SymFn(symshen_4newpv), B6229)


let__18013 := tmp19207
_ = let__18013

tmp19208 := Call1(__e, SymFn(symshen_4newpv), B6229)


let__18014 := tmp19208
_ = let__18014

tmp19209 := GCons(__e, let__18013, Nil)

tmp19210 := GCons(__e, symlist, tmp19209)

tmp19211 := GCons(__e, let__18014, Nil)

tmp19212 := GCons(__e, tmp19210, tmp19211)

tmp19213 := GCons(__e, symstr, tmp19212)

tmp19214 := GCons(__e, symboolean, Nil)

tmp19215 := GCons(__e, sym_1_1_6, tmp19214)

tmp19216 := GCons(__e, tmp19213, tmp19215)

tmp19217 := Call6(__e, SymFn(symis_b), V6233, tmp19216, B6229, L6230, Key6231, C6232)


tmp19218 := Call2(__e, SymFn(symshen_4gc), B6229, tmp19217)


return TailCall2(__e, SymFn(symshen_4gc), B6229, tmp19218)


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

tmp19219 := GValue(__e, symshen_4_dsigf_d)

tmp19220 := Call3(__e, SymFn(symshen_4assoc_1_6), symshen_4parse_1failure_2, tmp19206, tmp19219)


tmp19221 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19220)

_ = tmp19221

tmp19222 := MakeNative1(func(__e *ControlFlow, V6238 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6234 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6235 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6236 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6237 Obj) Obj {
tmp19223 := Call1(__e, SymFn(symshen_4newpv), B6234)


let__18015 := tmp19223
_ = let__18015

tmp19224 := Call1(__e, SymFn(symshen_4newpv), B6234)


let__18016 := tmp19224
_ = let__18016

tmp19225 := GCons(__e, let__18015, Nil)

tmp19226 := GCons(__e, symlist, tmp19225)

tmp19227 := GCons(__e, let__18016, Nil)

tmp19228 := GCons(__e, tmp19226, tmp19227)

tmp19229 := GCons(__e, symstr, tmp19228)

tmp19230 := GCons(__e, tmp19229, Nil)

tmp19231 := GCons(__e, sym_1_1_6, tmp19230)

tmp19232 := Call6(__e, SymFn(symis_b), V6238, tmp19231, B6234, L6235, Key6236, C6237)


tmp19233 := Call2(__e, SymFn(symshen_4gc), B6234, tmp19232)


return TailCall2(__e, SymFn(symshen_4gc), B6234, tmp19233)


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

tmp19234 := GValue(__e, symshen_4_dsigf_d)

tmp19235 := Call3(__e, SymFn(symshen_4assoc_1_6), symshen_4parse_1failure, tmp19222, tmp19234)


tmp19236 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19235)

_ = tmp19236

tmp19237 := MakeNative1(func(__e *ControlFlow, V6243 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6239 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6240 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6241 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6242 Obj) Obj {
tmp19238 := Call1(__e, SymFn(symshen_4newpv), B6239)


let__18017 := tmp19238
_ = let__18017

tmp19239 := Call1(__e, SymFn(symshen_4newpv), B6239)


let__18018 := tmp19239
_ = let__18018

tmp19240 := GCons(__e, let__18017, Nil)

tmp19241 := GCons(__e, symlist, tmp19240)

tmp19242 := GCons(__e, let__18018, Nil)

tmp19243 := GCons(__e, tmp19241, tmp19242)

tmp19244 := GCons(__e, symstr, tmp19243)

tmp19245 := GCons(__e, let__18018, Nil)

tmp19246 := GCons(__e, sym_1_1_6, tmp19245)

tmp19247 := GCons(__e, tmp19244, tmp19246)

tmp19248 := Call6(__e, SymFn(symis_b), V6243, tmp19247, B6239, L6240, Key6241, C6242)


tmp19249 := Call2(__e, SymFn(symshen_4gc), B6239, tmp19248)


return TailCall2(__e, SymFn(symshen_4gc), B6239, tmp19249)


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

tmp19250 := GValue(__e, symshen_4_dsigf_d)

tmp19251 := Call3(__e, SymFn(symshen_4assoc_1_6), symshen_4_5_1out, tmp19237, tmp19250)


tmp19252 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19251)

_ = tmp19252

tmp19253 := MakeNative1(func(__e *ControlFlow, V6249 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6245 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6246 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6247 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6248 Obj) Obj {
tmp19254 := Call1(__e, SymFn(symshen_4newpv), B6245)


let__18019 := tmp19254
_ = let__18019

tmp19255 := Call1(__e, SymFn(symshen_4newpv), B6245)


let__18020 := tmp19255
_ = let__18020

tmp19256 := GCons(__e, let__18020, Nil)

tmp19257 := GCons(__e, symlist, tmp19256)

tmp19258 := GCons(__e, let__18019, Nil)

tmp19259 := GCons(__e, tmp19257, tmp19258)

tmp19260 := GCons(__e, symstr, tmp19259)

tmp19261 := GCons(__e, let__18020, Nil)

tmp19262 := GCons(__e, symlist, tmp19261)

tmp19263 := GCons(__e, tmp19262, Nil)

tmp19264 := GCons(__e, sym_1_1_6, tmp19263)

tmp19265 := GCons(__e, tmp19260, tmp19264)

tmp19266 := Call6(__e, SymFn(symis_b), V6249, tmp19265, B6245, L6246, Key6247, C6248)


tmp19267 := Call2(__e, SymFn(symshen_4gc), B6245, tmp19266)


return TailCall2(__e, SymFn(symshen_4gc), B6245, tmp19267)


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

tmp19268 := GValue(__e, symshen_4_dsigf_d)

tmp19269 := Call3(__e, SymFn(symshen_4assoc_1_6), symshen_4in_1_6, tmp19253, tmp19268)


tmp19270 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19269)

_ = tmp19270

tmp19271 := MakeNative1(func(__e *ControlFlow, V6255 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6251 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6252 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6253 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6254 Obj) Obj {
tmp19272 := Call1(__e, SymFn(symshen_4newpv), B6251)


let__18021 := tmp19272
_ = let__18021

tmp19273 := Call1(__e, SymFn(symshen_4newpv), B6251)


let__18022 := tmp19273
_ = let__18022

tmp19274 := GCons(__e, let__18021, Nil)

tmp19275 := GCons(__e, symlist, tmp19274)

tmp19276 := GCons(__e, let__18021, Nil)

tmp19277 := GCons(__e, symlist, tmp19276)

tmp19278 := GCons(__e, let__18022, Nil)

tmp19279 := GCons(__e, tmp19277, tmp19278)

tmp19280 := GCons(__e, symstr, tmp19279)

tmp19281 := GCons(__e, tmp19280, Nil)

tmp19282 := GCons(__e, sym_1_1_6, tmp19281)

tmp19283 := GCons(__e, let__18022, tmp19282)

tmp19284 := GCons(__e, tmp19283, Nil)

tmp19285 := GCons(__e, sym_1_1_6, tmp19284)

tmp19286 := GCons(__e, tmp19275, tmp19285)

tmp19287 := Call6(__e, SymFn(symis_b), V6255, tmp19286, B6251, L6252, Key6253, C6254)


tmp19288 := Call2(__e, SymFn(symshen_4gc), B6251, tmp19287)


return TailCall2(__e, SymFn(symshen_4gc), B6251, tmp19288)


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

tmp19289 := GValue(__e, symshen_4_dsigf_d)

tmp19290 := Call3(__e, SymFn(symshen_4assoc_1_6), symshen_4comb, tmp19271, tmp19289)


tmp19291 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19290)

_ = tmp19291

tmp19292 := MakeNative1(func(__e *ControlFlow, V6260 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6256 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6257 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6258 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6259 Obj) Obj {
tmp19293 := Call1(__e, SymFn(symshen_4newpv), B6256)


let__18023 := tmp19293
_ = let__18023

tmp19294 := GCons(__e, let__18023, Nil)

tmp19295 := GCons(__e, symlist, tmp19294)

tmp19296 := GCons(__e, symboolean, Nil)

tmp19297 := GCons(__e, sym_1_1_6, tmp19296)

tmp19298 := GCons(__e, tmp19295, tmp19297)

tmp19299 := GCons(__e, tmp19298, Nil)

tmp19300 := GCons(__e, sym_1_1_6, tmp19299)

tmp19301 := GCons(__e, let__18023, tmp19300)

tmp19302 := Call6(__e, SymFn(symis_b), V6260, tmp19301, B6256, L6257, Key6258, C6259)


return TailCall2(__e, SymFn(symshen_4gc), B6256, tmp19302)


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

tmp19303 := GValue(__e, symshen_4_dsigf_d)

tmp19304 := Call3(__e, SymFn(symshen_4assoc_1_6), symelement_2, tmp19292, tmp19303)


tmp19305 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19304)

_ = tmp19305

tmp19306 := MakeNative1(func(__e *ControlFlow, V6265 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6261 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6262 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6263 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6264 Obj) Obj {
tmp19307 := Call1(__e, SymFn(symshen_4newpv), B6261)


let__18024 := tmp19307
_ = let__18024

tmp19308 := GCons(__e, symboolean, Nil)

tmp19309 := GCons(__e, sym_1_1_6, tmp19308)

tmp19310 := GCons(__e, let__18024, tmp19309)

tmp19311 := Call6(__e, SymFn(symis_b), V6265, tmp19310, B6261, L6262, Key6263, C6264)


return TailCall2(__e, SymFn(symshen_4gc), B6261, tmp19311)


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

tmp19312 := GValue(__e, symshen_4_dsigf_d)

tmp19313 := Call3(__e, SymFn(symshen_4assoc_1_6), symempty_2, tmp19306, tmp19312)


tmp19314 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19313)

_ = tmp19314

tmp19315 := MakeNative1(func(__e *ControlFlow, V6270 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6266 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6267 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6268 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6269 Obj) Obj {
tmp19316 := GCons(__e, symboolean, Nil)

tmp19317 := GCons(__e, sym_1_1_6, tmp19316)

tmp19318 := GCons(__e, symsymbol, tmp19317)

return TailCall6(__e, SymFn(symis_b), V6270, tmp19318, B6266, L6267, Key6268, C6269)


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

tmp19319 := GValue(__e, symshen_4_dsigf_d)

tmp19320 := Call3(__e, SymFn(symshen_4assoc_1_6), symenable_1type_1theory, tmp19315, tmp19319)


tmp19321 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19320)

_ = tmp19321

tmp19322 := MakeNative1(func(__e *ControlFlow, V6275 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6271 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6272 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6273 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6274 Obj) Obj {
tmp19323 := GCons(__e, symsymbol, Nil)

tmp19324 := GCons(__e, symlist, tmp19323)

tmp19325 := GCons(__e, tmp19324, Nil)

tmp19326 := GCons(__e, sym_1_1_6, tmp19325)

tmp19327 := GCons(__e, symsymbol, tmp19326)

return TailCall6(__e, SymFn(symis_b), V6275, tmp19327, B6271, L6272, Key6273, C6274)


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

tmp19328 := GValue(__e, symshen_4_dsigf_d)

tmp19329 := Call3(__e, SymFn(symshen_4assoc_1_6), symexternal, tmp19322, tmp19328)


tmp19330 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19329)

_ = tmp19330

tmp19331 := MakeNative1(func(__e *ControlFlow, V6280 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6276 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6277 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6278 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6279 Obj) Obj {
tmp19332 := GCons(__e, symstring, Nil)

tmp19333 := GCons(__e, sym_1_1_6, tmp19332)

tmp19334 := GCons(__e, symexception, tmp19333)

return TailCall6(__e, SymFn(symis_b), V6280, tmp19334, B6276, L6277, Key6278, C6279)


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

tmp19335 := GValue(__e, symshen_4_dsigf_d)

tmp19336 := Call3(__e, SymFn(symshen_4assoc_1_6), symerror_1to_1string, tmp19331, tmp19335)


tmp19337 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19336)

_ = tmp19337

tmp19338 := MakeNative1(func(__e *ControlFlow, V6285 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6281 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6282 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6283 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6284 Obj) Obj {
tmp19339 := Call1(__e, SymFn(symshen_4newpv), B6281)


let__18025 := tmp19339
_ = let__18025

tmp19340 := GCons(__e, symstring, Nil)

tmp19341 := GCons(__e, symlist, tmp19340)

tmp19342 := GCons(__e, tmp19341, Nil)

tmp19343 := GCons(__e, sym_1_1_6, tmp19342)

tmp19344 := GCons(__e, let__18025, tmp19343)

tmp19345 := Call6(__e, SymFn(symis_b), V6285, tmp19344, B6281, L6282, Key6283, C6284)


return TailCall2(__e, SymFn(symshen_4gc), B6281, tmp19345)


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

tmp19346 := GValue(__e, symshen_4_dsigf_d)

tmp19347 := Call3(__e, SymFn(symshen_4assoc_1_6), symexplode, tmp19338, tmp19346)


tmp19348 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19347)

_ = tmp19348

tmp19349 := MakeNative1(func(__e *ControlFlow, V6290 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6286 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6287 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6288 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6289 Obj) Obj {
tmp19350 := GCons(__e, symsymbol, Nil)

tmp19351 := GCons(__e, sym_1_1_6, tmp19350)

tmp19352 := GCons(__e, symsymbol, tmp19351)

return TailCall6(__e, SymFn(symis_b), V6290, tmp19352, B6286, L6287, Key6288, C6289)


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

tmp19353 := GValue(__e, symshen_4_dsigf_d)

tmp19354 := Call3(__e, SymFn(symshen_4assoc_1_6), symfactorise, tmp19349, tmp19353)


tmp19355 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19354)

_ = tmp19355

tmp19356 := MakeNative1(func(__e *ControlFlow, V6295 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6291 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6292 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6293 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6294 Obj) Obj {
tmp19357 := GCons(__e, symboolean, Nil)

tmp19358 := GCons(__e, sym_1_1_6, tmp19357)

return TailCall6(__e, SymFn(symis_b), V6295, tmp19358, B6291, L6292, Key6293, C6294)


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

tmp19359 := GValue(__e, symshen_4_dsigf_d)

tmp19360 := Call3(__e, SymFn(symshen_4assoc_1_6), symfactorise_2, tmp19356, tmp19359)


tmp19361 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19360)

_ = tmp19361

tmp19362 := MakeNative1(func(__e *ControlFlow, V6300 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6296 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6297 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6298 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6299 Obj) Obj {
tmp19363 := GCons(__e, symsymbol, Nil)

tmp19364 := GCons(__e, sym_1_1_6, tmp19363)

return TailCall6(__e, SymFn(symis_b), V6300, tmp19364, B6296, L6297, Key6298, C6299)


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

tmp19365 := GValue(__e, symshen_4_dsigf_d)

tmp19366 := Call3(__e, SymFn(symshen_4assoc_1_6), symfail, tmp19362, tmp19365)


tmp19367 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19366)

_ = tmp19367

tmp19368 := MakeNative1(func(__e *ControlFlow, V6306 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6302 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6303 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6304 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6305 Obj) Obj {
tmp19369 := Call1(__e, SymFn(symshen_4newpv), B6302)


let__18026 := tmp19369
_ = let__18026

tmp19370 := GCons(__e, let__18026, Nil)

tmp19371 := GCons(__e, sym_1_1_6, tmp19370)

tmp19372 := GCons(__e, let__18026, tmp19371)

tmp19373 := GCons(__e, let__18026, Nil)

tmp19374 := GCons(__e, sym_1_1_6, tmp19373)

tmp19375 := GCons(__e, let__18026, tmp19374)

tmp19376 := GCons(__e, tmp19375, Nil)

tmp19377 := GCons(__e, sym_1_1_6, tmp19376)

tmp19378 := GCons(__e, tmp19372, tmp19377)

tmp19379 := Call6(__e, SymFn(symis_b), V6306, tmp19378, B6302, L6303, Key6304, C6305)


return TailCall2(__e, SymFn(symshen_4gc), B6302, tmp19379)


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

tmp19380 := GValue(__e, symshen_4_dsigf_d)

tmp19381 := Call3(__e, SymFn(symshen_4assoc_1_6), symfix, tmp19368, tmp19380)


tmp19382 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19381)

_ = tmp19382

tmp19383 := MakeNative1(func(__e *ControlFlow, V6311 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6307 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6308 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6309 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6310 Obj) Obj {
tmp19384 := Call1(__e, SymFn(symshen_4newpv), B6307)


let__18027 := tmp19384
_ = let__18027

tmp19385 := GCons(__e, let__18027, Nil)

tmp19386 := GCons(__e, symlazy, tmp19385)

tmp19387 := GCons(__e, tmp19386, Nil)

tmp19388 := GCons(__e, sym_1_1_6, tmp19387)

tmp19389 := GCons(__e, let__18027, tmp19388)

tmp19390 := Call6(__e, SymFn(symis_b), V6311, tmp19389, B6307, L6308, Key6309, C6310)


return TailCall2(__e, SymFn(symshen_4gc), B6307, tmp19390)


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

tmp19391 := GValue(__e, symshen_4_dsigf_d)

tmp19392 := Call3(__e, SymFn(symshen_4assoc_1_6), symfreeze, tmp19383, tmp19391)


tmp19393 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19392)

_ = tmp19393

tmp19394 := MakeNative1(func(__e *ControlFlow, V6316 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6312 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6313 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6314 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6315 Obj) Obj {
tmp19395 := Call1(__e, SymFn(symshen_4newpv), B6312)


let__18028 := tmp19395
_ = let__18028

tmp19396 := Call1(__e, SymFn(symshen_4newpv), B6312)


let__18029 := tmp19396
_ = let__18029

tmp19397 := GCons(__e, let__18028, Nil)

tmp19398 := GCons(__e, sym_d, tmp19397)

tmp19399 := GCons(__e, let__18029, tmp19398)

tmp19400 := GCons(__e, let__18029, Nil)

tmp19401 := GCons(__e, sym_1_1_6, tmp19400)

tmp19402 := GCons(__e, tmp19399, tmp19401)

tmp19403 := Call6(__e, SymFn(symis_b), V6316, tmp19402, B6312, L6313, Key6314, C6315)


tmp19404 := Call2(__e, SymFn(symshen_4gc), B6312, tmp19403)


return TailCall2(__e, SymFn(symshen_4gc), B6312, tmp19404)


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

tmp19405 := GValue(__e, symshen_4_dsigf_d)

tmp19406 := Call3(__e, SymFn(symshen_4assoc_1_6), symfst, tmp19394, tmp19405)


tmp19407 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19406)

_ = tmp19407

tmp19408 := MakeNative1(func(__e *ControlFlow, V6321 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6317 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6318 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6319 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6320 Obj) Obj {
tmp19409 := GCons(__e, symsymbol, Nil)

tmp19410 := GCons(__e, sym_1_1_6, tmp19409)

tmp19411 := GCons(__e, symsymbol, tmp19410)

return TailCall6(__e, SymFn(symis_b), V6321, tmp19411, B6317, L6318, Key6319, C6320)


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

tmp19412 := GValue(__e, symshen_4_dsigf_d)

tmp19413 := Call3(__e, SymFn(symshen_4assoc_1_6), symgensym, tmp19408, tmp19412)


tmp19414 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19413)

_ = tmp19414

tmp19415 := MakeNative1(func(__e *ControlFlow, V6327 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6323 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6324 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6325 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6326 Obj) Obj {
tmp19416 := Call1(__e, SymFn(symshen_4newpv), B6323)


let__18030 := tmp19416
_ = let__18030

tmp19417 := GCons(__e, let__18030, Nil)

tmp19418 := GCons(__e, symlist, tmp19417)

tmp19419 := GCons(__e, symboolean, Nil)

tmp19420 := GCons(__e, sym_1_1_6, tmp19419)

tmp19421 := GCons(__e, let__18030, tmp19420)

tmp19422 := GCons(__e, tmp19421, Nil)

tmp19423 := GCons(__e, sym_1_1_6, tmp19422)

tmp19424 := GCons(__e, tmp19418, tmp19423)

tmp19425 := Call6(__e, SymFn(symis_b), V6327, tmp19424, B6323, L6324, Key6325, C6326)


return TailCall2(__e, SymFn(symshen_4gc), B6323, tmp19425)


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

tmp19426 := GValue(__e, symshen_4_dsigf_d)

tmp19427 := Call3(__e, SymFn(symshen_4assoc_1_6), symshen_4hds_a_2, tmp19415, tmp19426)


tmp19428 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19427)

_ = tmp19428

tmp19429 := MakeNative1(func(__e *ControlFlow, V6332 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6328 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6329 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6330 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6331 Obj) Obj {
tmp19430 := GCons(__e, symboolean, Nil)

tmp19431 := GCons(__e, sym_1_1_6, tmp19430)

tmp19432 := GCons(__e, symsymbol, tmp19431)

return TailCall6(__e, SymFn(symis_b), V6332, tmp19432, B6328, L6329, Key6330, C6331)


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

tmp19433 := GValue(__e, symshen_4_dsigf_d)

tmp19434 := Call3(__e, SymFn(symshen_4assoc_1_6), symhush, tmp19429, tmp19433)


tmp19435 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19434)

_ = tmp19435

tmp19436 := MakeNative1(func(__e *ControlFlow, V6337 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6333 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6334 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6335 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6336 Obj) Obj {
tmp19437 := GCons(__e, symboolean, Nil)

tmp19438 := GCons(__e, sym_1_1_6, tmp19437)

return TailCall6(__e, SymFn(symis_b), V6337, tmp19438, B6333, L6334, Key6335, C6336)


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

tmp19439 := GValue(__e, symshen_4_dsigf_d)

tmp19440 := Call3(__e, SymFn(symshen_4assoc_1_6), symhush_2, tmp19436, tmp19439)


tmp19441 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19440)

_ = tmp19441

tmp19442 := MakeNative1(func(__e *ControlFlow, V6343 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6339 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6340 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6341 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6342 Obj) Obj {
tmp19443 := Call1(__e, SymFn(symshen_4newpv), B6339)


let__18031 := tmp19443
_ = let__18031

tmp19444 := GCons(__e, let__18031, Nil)

tmp19445 := GCons(__e, symvector, tmp19444)

tmp19446 := GCons(__e, let__18031, Nil)

tmp19447 := GCons(__e, sym_1_1_6, tmp19446)

tmp19448 := GCons(__e, symnumber, tmp19447)

tmp19449 := GCons(__e, tmp19448, Nil)

tmp19450 := GCons(__e, sym_1_1_6, tmp19449)

tmp19451 := GCons(__e, tmp19445, tmp19450)

tmp19452 := Call6(__e, SymFn(symis_b), V6343, tmp19451, B6339, L6340, Key6341, C6342)


return TailCall2(__e, SymFn(symshen_4gc), B6339, tmp19452)


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

tmp19453 := GValue(__e, symshen_4_dsigf_d)

tmp19454 := Call3(__e, SymFn(symshen_4assoc_1_6), sym_5_1vector, tmp19442, tmp19453)


tmp19455 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19454)

_ = tmp19455

tmp19456 := MakeNative1(func(__e *ControlFlow, V6349 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6345 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6346 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6347 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6348 Obj) Obj {
tmp19457 := Call1(__e, SymFn(symshen_4newpv), B6345)


let__18032 := tmp19457
_ = let__18032

tmp19458 := GCons(__e, let__18032, Nil)

tmp19459 := GCons(__e, symvector, tmp19458)

tmp19460 := GCons(__e, let__18032, Nil)

tmp19461 := GCons(__e, symvector, tmp19460)

tmp19462 := GCons(__e, tmp19461, Nil)

tmp19463 := GCons(__e, sym_1_1_6, tmp19462)

tmp19464 := GCons(__e, let__18032, tmp19463)

tmp19465 := GCons(__e, tmp19464, Nil)

tmp19466 := GCons(__e, sym_1_1_6, tmp19465)

tmp19467 := GCons(__e, symnumber, tmp19466)

tmp19468 := GCons(__e, tmp19467, Nil)

tmp19469 := GCons(__e, sym_1_1_6, tmp19468)

tmp19470 := GCons(__e, tmp19459, tmp19469)

tmp19471 := Call6(__e, SymFn(symis_b), V6349, tmp19470, B6345, L6346, Key6347, C6348)


return TailCall2(__e, SymFn(symshen_4gc), B6345, tmp19471)


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

tmp19472 := GValue(__e, symshen_4_dsigf_d)

tmp19473 := Call3(__e, SymFn(symshen_4assoc_1_6), symvector_1_6, tmp19456, tmp19472)


tmp19474 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19473)

_ = tmp19474

tmp19475 := MakeNative1(func(__e *ControlFlow, V6354 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6350 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6351 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6352 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6353 Obj) Obj {
tmp19476 := Call1(__e, SymFn(symshen_4newpv), B6350)


let__18033 := tmp19476
_ = let__18033

tmp19477 := GCons(__e, let__18033, Nil)

tmp19478 := GCons(__e, symvector, tmp19477)

tmp19479 := GCons(__e, tmp19478, Nil)

tmp19480 := GCons(__e, sym_1_1_6, tmp19479)

tmp19481 := GCons(__e, symnumber, tmp19480)

tmp19482 := Call6(__e, SymFn(symis_b), V6354, tmp19481, B6350, L6351, Key6352, C6353)


return TailCall2(__e, SymFn(symshen_4gc), B6350, tmp19482)


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

tmp19483 := GValue(__e, symshen_4_dsigf_d)

tmp19484 := Call3(__e, SymFn(symshen_4assoc_1_6), symvector, tmp19475, tmp19483)


tmp19485 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19484)

_ = tmp19485

tmp19486 := MakeNative1(func(__e *ControlFlow, V6359 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6355 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6356 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6357 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6358 Obj) Obj {
tmp19487 := GCons(__e, symnumber, Nil)

tmp19488 := GCons(__e, sym_1_1_6, tmp19487)

tmp19489 := GCons(__e, symsymbol, tmp19488)

return TailCall6(__e, SymFn(symis_b), V6359, tmp19489, B6355, L6356, Key6357, C6358)


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

tmp19490 := GValue(__e, symshen_4_dsigf_d)

tmp19491 := Call3(__e, SymFn(symshen_4assoc_1_6), symget_1time, tmp19486, tmp19490)


tmp19492 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19491)

_ = tmp19492

tmp19493 := MakeNative1(func(__e *ControlFlow, V6364 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6360 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6361 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6362 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6363 Obj) Obj {
tmp19494 := Call1(__e, SymFn(symshen_4newpv), B6360)


let__18034 := tmp19494
_ = let__18034

tmp19495 := GCons(__e, symnumber, Nil)

tmp19496 := GCons(__e, sym_1_1_6, tmp19495)

tmp19497 := GCons(__e, symnumber, tmp19496)

tmp19498 := GCons(__e, tmp19497, Nil)

tmp19499 := GCons(__e, sym_1_1_6, tmp19498)

tmp19500 := GCons(__e, let__18034, tmp19499)

tmp19501 := Call6(__e, SymFn(symis_b), V6364, tmp19500, B6360, L6361, Key6362, C6363)


return TailCall2(__e, SymFn(symshen_4gc), B6360, tmp19501)


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

tmp19502 := GValue(__e, symshen_4_dsigf_d)

tmp19503 := Call3(__e, SymFn(symshen_4assoc_1_6), symhash, tmp19493, tmp19502)


tmp19504 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19503)

_ = tmp19504

tmp19505 := MakeNative1(func(__e *ControlFlow, V6369 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6365 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6366 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6367 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6368 Obj) Obj {
tmp19506 := Call1(__e, SymFn(symshen_4newpv), B6365)


let__18035 := tmp19506
_ = let__18035

tmp19507 := GCons(__e, let__18035, Nil)

tmp19508 := GCons(__e, symlist, tmp19507)

tmp19509 := GCons(__e, let__18035, Nil)

tmp19510 := GCons(__e, sym_1_1_6, tmp19509)

tmp19511 := GCons(__e, tmp19508, tmp19510)

tmp19512 := Call6(__e, SymFn(symis_b), V6369, tmp19511, B6365, L6366, Key6367, C6368)


return TailCall2(__e, SymFn(symshen_4gc), B6365, tmp19512)


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

tmp19513 := GValue(__e, symshen_4_dsigf_d)

tmp19514 := Call3(__e, SymFn(symshen_4assoc_1_6), symhead, tmp19505, tmp19513)


tmp19515 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19514)

_ = tmp19515

tmp19516 := MakeNative1(func(__e *ControlFlow, V6374 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6370 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6371 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6372 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6373 Obj) Obj {
tmp19517 := Call1(__e, SymFn(symshen_4newpv), B6370)


let__18036 := tmp19517
_ = let__18036

tmp19518 := GCons(__e, let__18036, Nil)

tmp19519 := GCons(__e, symvector, tmp19518)

tmp19520 := GCons(__e, let__18036, Nil)

tmp19521 := GCons(__e, sym_1_1_6, tmp19520)

tmp19522 := GCons(__e, tmp19519, tmp19521)

tmp19523 := Call6(__e, SymFn(symis_b), V6374, tmp19522, B6370, L6371, Key6372, C6373)


return TailCall2(__e, SymFn(symshen_4gc), B6370, tmp19523)


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

tmp19524 := GValue(__e, symshen_4_dsigf_d)

tmp19525 := Call3(__e, SymFn(symshen_4assoc_1_6), symhdv, tmp19516, tmp19524)


tmp19526 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19525)

_ = tmp19526

tmp19527 := MakeNative1(func(__e *ControlFlow, V6379 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6375 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6376 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6377 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6378 Obj) Obj {
tmp19528 := GCons(__e, symstring, Nil)

tmp19529 := GCons(__e, sym_1_1_6, tmp19528)

tmp19530 := GCons(__e, symstring, tmp19529)

return TailCall6(__e, SymFn(symis_b), V6379, tmp19530, B6375, L6376, Key6377, C6378)


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

tmp19531 := GValue(__e, symshen_4_dsigf_d)

tmp19532 := Call3(__e, SymFn(symshen_4assoc_1_6), symhdstr, tmp19527, tmp19531)


tmp19533 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19532)

_ = tmp19533

tmp19534 := MakeNative1(func(__e *ControlFlow, V6384 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6380 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6381 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6382 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6383 Obj) Obj {
tmp19535 := Call1(__e, SymFn(symshen_4newpv), B6380)


let__18037 := tmp19535
_ = let__18037

tmp19536 := GCons(__e, let__18037, Nil)

tmp19537 := GCons(__e, sym_1_1_6, tmp19536)

tmp19538 := GCons(__e, let__18037, tmp19537)

tmp19539 := GCons(__e, tmp19538, Nil)

tmp19540 := GCons(__e, sym_1_1_6, tmp19539)

tmp19541 := GCons(__e, let__18037, tmp19540)

tmp19542 := GCons(__e, tmp19541, Nil)

tmp19543 := GCons(__e, sym_1_1_6, tmp19542)

tmp19544 := GCons(__e, symboolean, tmp19543)

tmp19545 := Call6(__e, SymFn(symis_b), V6384, tmp19544, B6380, L6381, Key6382, C6383)


return TailCall2(__e, SymFn(symshen_4gc), B6380, tmp19545)


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

tmp19546 := GValue(__e, symshen_4_dsigf_d)

tmp19547 := Call3(__e, SymFn(symshen_4assoc_1_6), symif, tmp19534, tmp19546)


tmp19548 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19547)

_ = tmp19548

tmp19549 := MakeNative1(func(__e *ControlFlow, V6389 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6385 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6386 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6387 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6388 Obj) Obj {
tmp19550 := GCons(__e, symsymbol, Nil)

tmp19551 := GCons(__e, sym_1_1_6, tmp19550)

tmp19552 := GCons(__e, symsymbol, tmp19551)

return TailCall6(__e, SymFn(symis_b), V6389, tmp19552, B6385, L6386, Key6387, C6388)


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

tmp19553 := GValue(__e, symshen_4_dsigf_d)

tmp19554 := Call3(__e, SymFn(symshen_4assoc_1_6), symin_1package, tmp19549, tmp19553)


tmp19555 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19554)

_ = tmp19555

tmp19556 := MakeNative1(func(__e *ControlFlow, V6394 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6390 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6391 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6392 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6393 Obj) Obj {
tmp19557 := GCons(__e, symstring, Nil)

tmp19558 := GCons(__e, sym_1_1_6, tmp19557)

return TailCall6(__e, SymFn(symis_b), V6394, tmp19558, B6390, L6391, Key6392, C6393)


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

tmp19559 := GValue(__e, symshen_4_dsigf_d)

tmp19560 := Call3(__e, SymFn(symshen_4assoc_1_6), symit, tmp19556, tmp19559)


tmp19561 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19560)

_ = tmp19561

tmp19562 := MakeNative1(func(__e *ControlFlow, V6399 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6395 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6396 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6397 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6398 Obj) Obj {
tmp19563 := GCons(__e, symstring, Nil)

tmp19564 := GCons(__e, sym_1_1_6, tmp19563)

return TailCall6(__e, SymFn(symis_b), V6399, tmp19564, B6395, L6396, Key6397, C6398)


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

tmp19565 := GValue(__e, symshen_4_dsigf_d)

tmp19566 := Call3(__e, SymFn(symshen_4assoc_1_6), symimplementation, tmp19562, tmp19565)


tmp19567 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19566)

_ = tmp19567

tmp19568 := MakeNative1(func(__e *ControlFlow, V6405 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6401 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6402 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6403 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6404 Obj) Obj {
tmp19569 := GCons(__e, symsymbol, Nil)

tmp19570 := GCons(__e, symlist, tmp19569)

tmp19571 := GCons(__e, symsymbol, Nil)

tmp19572 := GCons(__e, symlist, tmp19571)

tmp19573 := GCons(__e, tmp19572, Nil)

tmp19574 := GCons(__e, sym_1_1_6, tmp19573)

tmp19575 := GCons(__e, tmp19570, tmp19574)

return TailCall6(__e, SymFn(symis_b), V6405, tmp19575, B6401, L6402, Key6403, C6404)


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

tmp19576 := GValue(__e, symshen_4_dsigf_d)

tmp19577 := Call3(__e, SymFn(symshen_4assoc_1_6), syminclude, tmp19568, tmp19576)


tmp19578 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19577)

_ = tmp19578

tmp19579 := MakeNative1(func(__e *ControlFlow, V6411 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6407 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6408 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6409 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6410 Obj) Obj {
tmp19580 := GCons(__e, symsymbol, Nil)

tmp19581 := GCons(__e, symlist, tmp19580)

tmp19582 := GCons(__e, symsymbol, Nil)

tmp19583 := GCons(__e, symlist, tmp19582)

tmp19584 := GCons(__e, tmp19583, Nil)

tmp19585 := GCons(__e, sym_1_1_6, tmp19584)

tmp19586 := GCons(__e, tmp19581, tmp19585)

return TailCall6(__e, SymFn(symis_b), V6411, tmp19586, B6407, L6408, Key6409, C6410)


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

tmp19587 := GValue(__e, symshen_4_dsigf_d)

tmp19588 := Call3(__e, SymFn(symshen_4assoc_1_6), syminclude_1all_1but, tmp19579, tmp19587)


tmp19589 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19588)

_ = tmp19589

tmp19590 := MakeNative1(func(__e *ControlFlow, V6416 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6412 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6413 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6414 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6415 Obj) Obj {
tmp19591 := GCons(__e, symsymbol, Nil)

tmp19592 := GCons(__e, symlist, tmp19591)

tmp19593 := GCons(__e, tmp19592, Nil)

tmp19594 := GCons(__e, sym_1_1_6, tmp19593)

return TailCall6(__e, SymFn(symis_b), V6416, tmp19594, B6412, L6413, Key6414, C6415)


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

tmp19595 := GValue(__e, symshen_4_dsigf_d)

tmp19596 := Call3(__e, SymFn(symshen_4assoc_1_6), symincluded, tmp19590, tmp19595)


tmp19597 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19596)

_ = tmp19597

tmp19598 := MakeNative1(func(__e *ControlFlow, V6421 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6417 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6418 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6419 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6420 Obj) Obj {
tmp19599 := GCons(__e, symnumber, Nil)

tmp19600 := GCons(__e, sym_1_1_6, tmp19599)

return TailCall6(__e, SymFn(symis_b), V6421, tmp19600, B6417, L6418, Key6419, C6420)


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

tmp19601 := GValue(__e, symshen_4_dsigf_d)

tmp19602 := Call3(__e, SymFn(symshen_4assoc_1_6), syminferences, tmp19598, tmp19601)


tmp19603 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19602)

_ = tmp19603

tmp19604 := MakeNative1(func(__e *ControlFlow, V6426 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6422 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6423 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6424 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6425 Obj) Obj {
tmp19605 := Call1(__e, SymFn(symshen_4newpv), B6422)


let__18038 := tmp19605
_ = let__18038

tmp19606 := GCons(__e, symstring, Nil)

tmp19607 := GCons(__e, sym_1_1_6, tmp19606)

tmp19608 := GCons(__e, symstring, tmp19607)

tmp19609 := GCons(__e, tmp19608, Nil)

tmp19610 := GCons(__e, sym_1_1_6, tmp19609)

tmp19611 := GCons(__e, let__18038, tmp19610)

tmp19612 := Call6(__e, SymFn(symis_b), V6426, tmp19611, B6422, L6423, Key6424, C6425)


return TailCall2(__e, SymFn(symshen_4gc), B6422, tmp19612)


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

tmp19613 := GValue(__e, symshen_4_dsigf_d)

tmp19614 := Call3(__e, SymFn(symshen_4assoc_1_6), symshen_4insert, tmp19604, tmp19613)


tmp19615 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19614)

_ = tmp19615

tmp19616 := MakeNative1(func(__e *ControlFlow, V6431 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6427 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6428 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6429 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6430 Obj) Obj {
tmp19617 := Call1(__e, SymFn(symshen_4newpv), B6427)


let__18039 := tmp19617
_ = let__18039

tmp19618 := GCons(__e, symboolean, Nil)

tmp19619 := GCons(__e, sym_1_1_6, tmp19618)

tmp19620 := GCons(__e, let__18039, tmp19619)

tmp19621 := Call6(__e, SymFn(symis_b), V6431, tmp19620, B6427, L6428, Key6429, C6430)


return TailCall2(__e, SymFn(symshen_4gc), B6427, tmp19621)


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

tmp19622 := GValue(__e, symshen_4_dsigf_d)

tmp19623 := Call3(__e, SymFn(symshen_4assoc_1_6), syminteger_2, tmp19616, tmp19622)


tmp19624 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19623)

_ = tmp19624

tmp19625 := MakeNative1(func(__e *ControlFlow, V6436 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6432 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6433 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6434 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6435 Obj) Obj {
tmp19626 := GCons(__e, symsymbol, Nil)

tmp19627 := GCons(__e, symlist, tmp19626)

tmp19628 := GCons(__e, tmp19627, Nil)

tmp19629 := GCons(__e, sym_1_1_6, tmp19628)

tmp19630 := GCons(__e, symsymbol, tmp19629)

return TailCall6(__e, SymFn(symis_b), V6436, tmp19630, B6432, L6433, Key6434, C6435)


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

tmp19631 := GValue(__e, symshen_4_dsigf_d)

tmp19632 := Call3(__e, SymFn(symshen_4assoc_1_6), syminternal, tmp19625, tmp19631)


tmp19633 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19632)

_ = tmp19633

tmp19634 := MakeNative1(func(__e *ControlFlow, V6443 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6439 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6440 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6441 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6442 Obj) Obj {
tmp19635 := Call1(__e, SymFn(symshen_4newpv), B6439)


let__18040 := tmp19635
_ = let__18040

tmp19636 := GCons(__e, let__18040, Nil)

tmp19637 := GCons(__e, symlist, tmp19636)

tmp19638 := GCons(__e, let__18040, Nil)

tmp19639 := GCons(__e, symlist, tmp19638)

tmp19640 := GCons(__e, let__18040, Nil)

tmp19641 := GCons(__e, symlist, tmp19640)

tmp19642 := GCons(__e, tmp19641, Nil)

tmp19643 := GCons(__e, sym_1_1_6, tmp19642)

tmp19644 := GCons(__e, tmp19639, tmp19643)

tmp19645 := GCons(__e, tmp19644, Nil)

tmp19646 := GCons(__e, sym_1_1_6, tmp19645)

tmp19647 := GCons(__e, tmp19637, tmp19646)

tmp19648 := Call6(__e, SymFn(symis_b), V6443, tmp19647, B6439, L6440, Key6441, C6442)


return TailCall2(__e, SymFn(symshen_4gc), B6439, tmp19648)


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

tmp19649 := GValue(__e, symshen_4_dsigf_d)

tmp19650 := Call3(__e, SymFn(symshen_4assoc_1_6), symintersection, tmp19634, tmp19649)


tmp19651 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19650)

_ = tmp19651

tmp19652 := MakeNative1(func(__e *ControlFlow, V6448 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6444 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6445 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6446 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6447 Obj) Obj {
tmp19653 := GCons(__e, symstring, Nil)

tmp19654 := GCons(__e, sym_1_1_6, tmp19653)

return TailCall6(__e, SymFn(symis_b), V6448, tmp19654, B6444, L6445, Key6446, C6447)


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

tmp19655 := GValue(__e, symshen_4_dsigf_d)

tmp19656 := Call3(__e, SymFn(symshen_4assoc_1_6), symlanguage, tmp19652, tmp19655)


tmp19657 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19656)

_ = tmp19657

tmp19658 := MakeNative1(func(__e *ControlFlow, V6453 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6449 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6450 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6451 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6452 Obj) Obj {
tmp19659 := Call1(__e, SymFn(symshen_4newpv), B6449)


let__18041 := tmp19659
_ = let__18041

tmp19660 := GCons(__e, let__18041, Nil)

tmp19661 := GCons(__e, symlist, tmp19660)

tmp19662 := GCons(__e, symnumber, Nil)

tmp19663 := GCons(__e, sym_1_1_6, tmp19662)

tmp19664 := GCons(__e, tmp19661, tmp19663)

tmp19665 := Call6(__e, SymFn(symis_b), V6453, tmp19664, B6449, L6450, Key6451, C6452)


return TailCall2(__e, SymFn(symshen_4gc), B6449, tmp19665)


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

tmp19666 := GValue(__e, symshen_4_dsigf_d)

tmp19667 := Call3(__e, SymFn(symshen_4assoc_1_6), symlength, tmp19658, tmp19666)


tmp19668 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19667)

_ = tmp19668

tmp19669 := MakeNative1(func(__e *ControlFlow, V6458 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6454 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6455 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6456 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6457 Obj) Obj {
tmp19670 := Call1(__e, SymFn(symshen_4newpv), B6454)


let__18042 := tmp19670
_ = let__18042

tmp19671 := GCons(__e, let__18042, Nil)

tmp19672 := GCons(__e, symvector, tmp19671)

tmp19673 := GCons(__e, symnumber, Nil)

tmp19674 := GCons(__e, sym_1_1_6, tmp19673)

tmp19675 := GCons(__e, tmp19672, tmp19674)

tmp19676 := Call6(__e, SymFn(symis_b), V6458, tmp19675, B6454, L6455, Key6456, C6457)


return TailCall2(__e, SymFn(symshen_4gc), B6454, tmp19676)


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

tmp19677 := GValue(__e, symshen_4_dsigf_d)

tmp19678 := Call3(__e, SymFn(symshen_4assoc_1_6), symlimit, tmp19669, tmp19677)


tmp19679 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19678)

_ = tmp19679

tmp19680 := MakeNative1(func(__e *ControlFlow, V6464 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6460 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6461 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6462 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6463 Obj) Obj {
tmp19681 := GCons(__e, symin, Nil)

tmp19682 := GCons(__e, symstream, tmp19681)

tmp19683 := GCons(__e, symunit, Nil)

tmp19684 := GCons(__e, symlist, tmp19683)

tmp19685 := GCons(__e, tmp19684, Nil)

tmp19686 := GCons(__e, sym_1_1_6, tmp19685)

tmp19687 := GCons(__e, tmp19682, tmp19686)

return TailCall6(__e, SymFn(symis_b), V6464, tmp19687, B6460, L6461, Key6462, C6463)


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

tmp19688 := GValue(__e, symshen_4_dsigf_d)

tmp19689 := Call3(__e, SymFn(symshen_4assoc_1_6), symlineread, tmp19680, tmp19688)


tmp19690 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19689)

_ = tmp19690

tmp19691 := MakeNative1(func(__e *ControlFlow, V6469 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6465 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6466 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6467 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6468 Obj) Obj {
tmp19692 := GCons(__e, symsymbol, Nil)

tmp19693 := GCons(__e, sym_1_1_6, tmp19692)

tmp19694 := GCons(__e, symstring, tmp19693)

return TailCall6(__e, SymFn(symis_b), V6469, tmp19694, B6465, L6466, Key6467, C6468)


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

tmp19695 := GValue(__e, symshen_4_dsigf_d)

tmp19696 := Call3(__e, SymFn(symshen_4assoc_1_6), symload, tmp19691, tmp19695)


tmp19697 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19696)

_ = tmp19697

tmp19698 := MakeNative1(func(__e *ControlFlow, V6476 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6472 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6473 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6474 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6475 Obj) Obj {
tmp19699 := Call1(__e, SymFn(symshen_4newpv), B6472)


let__18043 := tmp19699
_ = let__18043

tmp19700 := Call1(__e, SymFn(symshen_4newpv), B6472)


let__18044 := tmp19700
_ = let__18044

tmp19701 := GCons(__e, let__18044, Nil)

tmp19702 := GCons(__e, sym_1_1_6, tmp19701)

tmp19703 := GCons(__e, let__18043, tmp19702)

tmp19704 := GCons(__e, let__18043, Nil)

tmp19705 := GCons(__e, symlist, tmp19704)

tmp19706 := GCons(__e, let__18044, Nil)

tmp19707 := GCons(__e, symlist, tmp19706)

tmp19708 := GCons(__e, tmp19707, Nil)

tmp19709 := GCons(__e, sym_1_1_6, tmp19708)

tmp19710 := GCons(__e, tmp19705, tmp19709)

tmp19711 := GCons(__e, tmp19710, Nil)

tmp19712 := GCons(__e, sym_1_1_6, tmp19711)

tmp19713 := GCons(__e, tmp19703, tmp19712)

tmp19714 := Call6(__e, SymFn(symis_b), V6476, tmp19713, B6472, L6473, Key6474, C6475)


tmp19715 := Call2(__e, SymFn(symshen_4gc), B6472, tmp19714)


return TailCall2(__e, SymFn(symshen_4gc), B6472, tmp19715)


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

tmp19716 := GValue(__e, symshen_4_dsigf_d)

tmp19717 := Call3(__e, SymFn(symshen_4assoc_1_6), symmap, tmp19698, tmp19716)


tmp19718 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19717)

_ = tmp19718

tmp19719 := MakeNative1(func(__e *ControlFlow, V6483 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6479 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6480 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6481 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6482 Obj) Obj {
tmp19720 := Call1(__e, SymFn(symshen_4newpv), B6479)


let__18045 := tmp19720
_ = let__18045

tmp19721 := Call1(__e, SymFn(symshen_4newpv), B6479)


let__18046 := tmp19721
_ = let__18046

tmp19722 := GCons(__e, let__18046, Nil)

tmp19723 := GCons(__e, symlist, tmp19722)

tmp19724 := GCons(__e, tmp19723, Nil)

tmp19725 := GCons(__e, sym_1_1_6, tmp19724)

tmp19726 := GCons(__e, let__18045, tmp19725)

tmp19727 := GCons(__e, let__18045, Nil)

tmp19728 := GCons(__e, symlist, tmp19727)

tmp19729 := GCons(__e, let__18046, Nil)

tmp19730 := GCons(__e, symlist, tmp19729)

tmp19731 := GCons(__e, tmp19730, Nil)

tmp19732 := GCons(__e, sym_1_1_6, tmp19731)

tmp19733 := GCons(__e, tmp19728, tmp19732)

tmp19734 := GCons(__e, tmp19733, Nil)

tmp19735 := GCons(__e, sym_1_1_6, tmp19734)

tmp19736 := GCons(__e, tmp19726, tmp19735)

tmp19737 := Call6(__e, SymFn(symis_b), V6483, tmp19736, B6479, L6480, Key6481, C6482)


tmp19738 := Call2(__e, SymFn(symshen_4gc), B6479, tmp19737)


return TailCall2(__e, SymFn(symshen_4gc), B6479, tmp19738)


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

tmp19739 := GValue(__e, symshen_4_dsigf_d)

tmp19740 := Call3(__e, SymFn(symshen_4assoc_1_6), symmapcan, tmp19719, tmp19739)


tmp19741 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19740)

_ = tmp19741

tmp19742 := MakeNative1(func(__e *ControlFlow, V6488 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6484 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6485 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6486 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6487 Obj) Obj {
tmp19743 := GCons(__e, symnumber, Nil)

tmp19744 := GCons(__e, sym_1_1_6, tmp19743)

tmp19745 := GCons(__e, symnumber, tmp19744)

return TailCall6(__e, SymFn(symis_b), V6488, tmp19745, B6484, L6485, Key6486, C6487)


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

tmp19746 := GValue(__e, symshen_4_dsigf_d)

tmp19747 := Call3(__e, SymFn(symshen_4assoc_1_6), symmaxinferences, tmp19742, tmp19746)


tmp19748 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19747)

_ = tmp19748

tmp19749 := MakeNative1(func(__e *ControlFlow, V6493 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6489 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6490 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6491 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6492 Obj) Obj {
tmp19750 := GCons(__e, symstring, Nil)

tmp19751 := GCons(__e, sym_1_1_6, tmp19750)

tmp19752 := GCons(__e, symnumber, tmp19751)

return TailCall6(__e, SymFn(symis_b), V6493, tmp19752, B6489, L6490, Key6491, C6492)


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

tmp19753 := GValue(__e, symshen_4_dsigf_d)

tmp19754 := Call3(__e, SymFn(symshen_4assoc_1_6), symn_1_6string, tmp19749, tmp19753)


tmp19755 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19754)

_ = tmp19755

tmp19756 := MakeNative1(func(__e *ControlFlow, V6498 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6494 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6495 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6496 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6497 Obj) Obj {
tmp19757 := GCons(__e, symnumber, Nil)

tmp19758 := GCons(__e, sym_1_1_6, tmp19757)

tmp19759 := GCons(__e, symnumber, tmp19758)

return TailCall6(__e, SymFn(symis_b), V6498, tmp19759, B6494, L6495, Key6496, C6497)


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

tmp19760 := GValue(__e, symshen_4_dsigf_d)

tmp19761 := Call3(__e, SymFn(symshen_4assoc_1_6), symnl, tmp19756, tmp19760)


tmp19762 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19761)

_ = tmp19762

tmp19763 := MakeNative1(func(__e *ControlFlow, V6503 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6499 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6500 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6501 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6502 Obj) Obj {
tmp19764 := GCons(__e, symboolean, Nil)

tmp19765 := GCons(__e, sym_1_1_6, tmp19764)

tmp19766 := GCons(__e, symboolean, tmp19765)

return TailCall6(__e, SymFn(symis_b), V6503, tmp19766, B6499, L6500, Key6501, C6502)


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

tmp19767 := GValue(__e, symshen_4_dsigf_d)

tmp19768 := Call3(__e, SymFn(symshen_4assoc_1_6), symnot, tmp19763, tmp19767)


tmp19769 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19768)

_ = tmp19769

tmp19770 := MakeNative1(func(__e *ControlFlow, V6508 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6504 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6505 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6506 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6507 Obj) Obj {
tmp19771 := Call1(__e, SymFn(symshen_4newpv), B6504)


let__18047 := tmp19771
_ = let__18047

tmp19772 := GCons(__e, let__18047, Nil)

tmp19773 := GCons(__e, symlist, tmp19772)

tmp19774 := GCons(__e, let__18047, Nil)

tmp19775 := GCons(__e, sym_1_1_6, tmp19774)

tmp19776 := GCons(__e, tmp19773, tmp19775)

tmp19777 := GCons(__e, tmp19776, Nil)

tmp19778 := GCons(__e, sym_1_1_6, tmp19777)

tmp19779 := GCons(__e, symnumber, tmp19778)

tmp19780 := Call6(__e, SymFn(symis_b), V6508, tmp19779, B6504, L6505, Key6506, C6507)


return TailCall2(__e, SymFn(symshen_4gc), B6504, tmp19780)


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

tmp19781 := GValue(__e, symshen_4_dsigf_d)

tmp19782 := Call3(__e, SymFn(symshen_4assoc_1_6), symnth, tmp19770, tmp19781)


tmp19783 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19782)

_ = tmp19783

tmp19784 := MakeNative1(func(__e *ControlFlow, V6513 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6509 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6510 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6511 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6512 Obj) Obj {
tmp19785 := Call1(__e, SymFn(symshen_4newpv), B6509)


let__18048 := tmp19785
_ = let__18048

tmp19786 := GCons(__e, symboolean, Nil)

tmp19787 := GCons(__e, sym_1_1_6, tmp19786)

tmp19788 := GCons(__e, let__18048, tmp19787)

tmp19789 := Call6(__e, SymFn(symis_b), V6513, tmp19788, B6509, L6510, Key6511, C6512)


return TailCall2(__e, SymFn(symshen_4gc), B6509, tmp19789)


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

tmp19790 := GValue(__e, symshen_4_dsigf_d)

tmp19791 := Call3(__e, SymFn(symshen_4assoc_1_6), symnumber_2, tmp19784, tmp19790)


tmp19792 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19791)

_ = tmp19792

tmp19793 := MakeNative1(func(__e *ControlFlow, V6518 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6514 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6515 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6516 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6517 Obj) Obj {
tmp19794 := Call1(__e, SymFn(symshen_4newpv), B6514)


let__18049 := tmp19794
_ = let__18049

tmp19795 := Call1(__e, SymFn(symshen_4newpv), B6514)


let__18050 := tmp19795
_ = let__18050

tmp19796 := GCons(__e, symnumber, Nil)

tmp19797 := GCons(__e, sym_1_1_6, tmp19796)

tmp19798 := GCons(__e, let__18050, tmp19797)

tmp19799 := GCons(__e, tmp19798, Nil)

tmp19800 := GCons(__e, sym_1_1_6, tmp19799)

tmp19801 := GCons(__e, let__18049, tmp19800)

tmp19802 := Call6(__e, SymFn(symis_b), V6518, tmp19801, B6514, L6515, Key6516, C6517)


tmp19803 := Call2(__e, SymFn(symshen_4gc), B6514, tmp19802)


return TailCall2(__e, SymFn(symshen_4gc), B6514, tmp19803)


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

tmp19804 := GValue(__e, symshen_4_dsigf_d)

tmp19805 := Call3(__e, SymFn(symshen_4assoc_1_6), symoccurrences, tmp19793, tmp19804)


tmp19806 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19805)

_ = tmp19806

tmp19807 := MakeNative1(func(__e *ControlFlow, V6523 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6519 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6520 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6521 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6522 Obj) Obj {
tmp19808 := GCons(__e, symboolean, Nil)

tmp19809 := GCons(__e, sym_1_1_6, tmp19808)

tmp19810 := GCons(__e, symsymbol, tmp19809)

return TailCall6(__e, SymFn(symis_b), V6523, tmp19810, B6519, L6520, Key6521, C6522)


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

tmp19811 := GValue(__e, symshen_4_dsigf_d)

tmp19812 := Call3(__e, SymFn(symshen_4assoc_1_6), symoccurs_1check, tmp19807, tmp19811)


tmp19813 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19812)

_ = tmp19813

tmp19814 := MakeNative1(func(__e *ControlFlow, V6528 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6524 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6525 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6526 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6527 Obj) Obj {
tmp19815 := GCons(__e, symboolean, Nil)

tmp19816 := GCons(__e, sym_1_1_6, tmp19815)

return TailCall6(__e, SymFn(symis_b), V6528, tmp19816, B6524, L6525, Key6526, C6527)


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

tmp19817 := GValue(__e, symshen_4_dsigf_d)

tmp19818 := Call3(__e, SymFn(symshen_4assoc_1_6), symoccurs_2, tmp19814, tmp19817)


tmp19819 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19818)

_ = tmp19819

tmp19820 := MakeNative1(func(__e *ControlFlow, V6533 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6529 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6530 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6531 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6532 Obj) Obj {
tmp19821 := GCons(__e, symboolean, Nil)

tmp19822 := GCons(__e, sym_1_1_6, tmp19821)

tmp19823 := GCons(__e, symsymbol, tmp19822)

return TailCall6(__e, SymFn(symis_b), V6533, tmp19823, B6529, L6530, Key6531, C6532)


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

tmp19824 := GValue(__e, symshen_4_dsigf_d)

tmp19825 := Call3(__e, SymFn(symshen_4assoc_1_6), symoptimise, tmp19820, tmp19824)


tmp19826 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19825)

_ = tmp19826

tmp19827 := MakeNative1(func(__e *ControlFlow, V6538 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6534 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6535 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6536 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6537 Obj) Obj {
tmp19828 := GCons(__e, symboolean, Nil)

tmp19829 := GCons(__e, sym_1_1_6, tmp19828)

return TailCall6(__e, SymFn(symis_b), V6538, tmp19829, B6534, L6535, Key6536, C6537)


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

tmp19830 := GValue(__e, symshen_4_dsigf_d)

tmp19831 := Call3(__e, SymFn(symshen_4assoc_1_6), symoptimise_2, tmp19827, tmp19830)


tmp19832 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19831)

_ = tmp19832

tmp19833 := MakeNative1(func(__e *ControlFlow, V6543 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6539 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6540 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6541 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6542 Obj) Obj {
tmp19834 := GCons(__e, symboolean, Nil)

tmp19835 := GCons(__e, sym_1_1_6, tmp19834)

tmp19836 := GCons(__e, symboolean, tmp19835)

tmp19837 := GCons(__e, tmp19836, Nil)

tmp19838 := GCons(__e, sym_1_1_6, tmp19837)

tmp19839 := GCons(__e, symboolean, tmp19838)

return TailCall6(__e, SymFn(symis_b), V6543, tmp19839, B6539, L6540, Key6541, C6542)


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

tmp19840 := GValue(__e, symshen_4_dsigf_d)

tmp19841 := Call3(__e, SymFn(symshen_4assoc_1_6), symor, tmp19833, tmp19840)


tmp19842 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19841)

_ = tmp19842

tmp19843 := MakeNative1(func(__e *ControlFlow, V6548 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6544 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6545 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6546 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6547 Obj) Obj {
tmp19844 := GCons(__e, symstring, Nil)

tmp19845 := GCons(__e, sym_1_1_6, tmp19844)

return TailCall6(__e, SymFn(symis_b), V6548, tmp19845, B6544, L6545, Key6546, C6547)


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

tmp19846 := GValue(__e, symshen_4_dsigf_d)

tmp19847 := Call3(__e, SymFn(symshen_4assoc_1_6), symos, tmp19843, tmp19846)


tmp19848 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19847)

_ = tmp19848

tmp19849 := MakeNative1(func(__e *ControlFlow, V6553 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6549 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6550 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6551 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6552 Obj) Obj {
tmp19850 := GCons(__e, symboolean, Nil)

tmp19851 := GCons(__e, sym_1_1_6, tmp19850)

tmp19852 := GCons(__e, symsymbol, tmp19851)

return TailCall6(__e, SymFn(symis_b), V6553, tmp19852, B6549, L6550, Key6551, C6552)


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

tmp19853 := GValue(__e, symshen_4_dsigf_d)

tmp19854 := Call3(__e, SymFn(symshen_4assoc_1_6), sympackage_2, tmp19849, tmp19853)


tmp19855 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19854)

_ = tmp19855

tmp19856 := MakeNative1(func(__e *ControlFlow, V6558 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6554 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6555 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6556 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6557 Obj) Obj {
tmp19857 := GCons(__e, symstring, Nil)

tmp19858 := GCons(__e, sym_1_1_6, tmp19857)

return TailCall6(__e, SymFn(symis_b), V6558, tmp19858, B6554, L6555, Key6556, C6557)


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

tmp19859 := GValue(__e, symshen_4_dsigf_d)

tmp19860 := Call3(__e, SymFn(symshen_4assoc_1_6), symport, tmp19856, tmp19859)


tmp19861 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19860)

_ = tmp19861

tmp19862 := MakeNative1(func(__e *ControlFlow, V6563 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6559 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6560 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6561 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6562 Obj) Obj {
tmp19863 := GCons(__e, symstring, Nil)

tmp19864 := GCons(__e, sym_1_1_6, tmp19863)

return TailCall6(__e, SymFn(symis_b), V6563, tmp19864, B6559, L6560, Key6561, C6562)


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

tmp19865 := GValue(__e, symshen_4_dsigf_d)

tmp19866 := Call3(__e, SymFn(symshen_4assoc_1_6), symporters, tmp19862, tmp19865)


tmp19867 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19866)

_ = tmp19867

tmp19868 := MakeNative1(func(__e *ControlFlow, V6568 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6564 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6565 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6566 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6567 Obj) Obj {
tmp19869 := GCons(__e, symstring, Nil)

tmp19870 := GCons(__e, sym_1_1_6, tmp19869)

tmp19871 := GCons(__e, symnumber, tmp19870)

tmp19872 := GCons(__e, tmp19871, Nil)

tmp19873 := GCons(__e, sym_1_1_6, tmp19872)

tmp19874 := GCons(__e, symstring, tmp19873)

return TailCall6(__e, SymFn(symis_b), V6568, tmp19874, B6564, L6565, Key6566, C6567)


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

tmp19875 := GValue(__e, symshen_4_dsigf_d)

tmp19876 := Call3(__e, SymFn(symshen_4assoc_1_6), sympos, tmp19868, tmp19875)


tmp19877 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19876)

_ = tmp19877

tmp19878 := MakeNative1(func(__e *ControlFlow, V6573 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6569 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6570 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6571 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6572 Obj) Obj {
tmp19879 := GCons(__e, symout, Nil)

tmp19880 := GCons(__e, symstream, tmp19879)

tmp19881 := GCons(__e, symstring, Nil)

tmp19882 := GCons(__e, sym_1_1_6, tmp19881)

tmp19883 := GCons(__e, tmp19880, tmp19882)

tmp19884 := GCons(__e, tmp19883, Nil)

tmp19885 := GCons(__e, sym_1_1_6, tmp19884)

tmp19886 := GCons(__e, symstring, tmp19885)

return TailCall6(__e, SymFn(symis_b), V6573, tmp19886, B6569, L6570, Key6571, C6572)


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

tmp19887 := GValue(__e, symshen_4_dsigf_d)

tmp19888 := Call3(__e, SymFn(symshen_4assoc_1_6), sympr, tmp19878, tmp19887)


tmp19889 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19888)

_ = tmp19889

tmp19890 := MakeNative1(func(__e *ControlFlow, V6578 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6574 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6575 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6576 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6577 Obj) Obj {
tmp19891 := Call1(__e, SymFn(symshen_4newpv), B6574)


let__18051 := tmp19891
_ = let__18051

tmp19892 := GCons(__e, let__18051, Nil)

tmp19893 := GCons(__e, sym_1_1_6, tmp19892)

tmp19894 := GCons(__e, let__18051, tmp19893)

tmp19895 := Call6(__e, SymFn(symis_b), V6578, tmp19894, B6574, L6575, Key6576, C6577)


return TailCall2(__e, SymFn(symshen_4gc), B6574, tmp19895)


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

tmp19896 := GValue(__e, symshen_4_dsigf_d)

tmp19897 := Call3(__e, SymFn(symshen_4assoc_1_6), symprint, tmp19890, tmp19896)


tmp19898 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19897)

_ = tmp19898

tmp19899 := MakeNative1(func(__e *ControlFlow, V6583 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6579 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6580 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6581 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6582 Obj) Obj {
tmp19900 := GCons(__e, symsymbol, Nil)

tmp19901 := GCons(__e, sym_1_1_6, tmp19900)

tmp19902 := GCons(__e, symsymbol, tmp19901)

return TailCall6(__e, SymFn(symis_b), V6583, tmp19902, B6579, L6580, Key6581, C6582)


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

tmp19903 := GValue(__e, symshen_4_dsigf_d)

tmp19904 := Call3(__e, SymFn(symshen_4assoc_1_6), symprofile, tmp19899, tmp19903)


tmp19905 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19904)

_ = tmp19905

tmp19906 := MakeNative1(func(__e *ControlFlow, V6589 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6585 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6586 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6587 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6588 Obj) Obj {
tmp19907 := GCons(__e, symsymbol, Nil)

tmp19908 := GCons(__e, symlist, tmp19907)

tmp19909 := GCons(__e, symsymbol, Nil)

tmp19910 := GCons(__e, symlist, tmp19909)

tmp19911 := GCons(__e, tmp19910, Nil)

tmp19912 := GCons(__e, sym_1_1_6, tmp19911)

tmp19913 := GCons(__e, tmp19908, tmp19912)

return TailCall6(__e, SymFn(symis_b), V6589, tmp19913, B6585, L6586, Key6587, C6588)


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

tmp19914 := GValue(__e, symshen_4_dsigf_d)

tmp19915 := Call3(__e, SymFn(symshen_4assoc_1_6), sympreclude, tmp19906, tmp19914)


tmp19916 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19915)

_ = tmp19916

tmp19917 := MakeNative1(func(__e *ControlFlow, V6594 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6590 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6591 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6592 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6593 Obj) Obj {
tmp19918 := GCons(__e, symstring, Nil)

tmp19919 := GCons(__e, sym_1_1_6, tmp19918)

tmp19920 := GCons(__e, symstring, tmp19919)

return TailCall6(__e, SymFn(symis_b), V6594, tmp19920, B6590, L6591, Key6592, C6593)


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

tmp19921 := GValue(__e, symshen_4_dsigf_d)

tmp19922 := Call3(__e, SymFn(symshen_4assoc_1_6), symshen_4proc_1nl, tmp19917, tmp19921)


tmp19923 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19922)

_ = tmp19923

tmp19924 := MakeNative1(func(__e *ControlFlow, V6599 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6595 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6596 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6597 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6598 Obj) Obj {
tmp19925 := GCons(__e, symnumber, Nil)

tmp19926 := GCons(__e, sym_d, tmp19925)

tmp19927 := GCons(__e, symsymbol, tmp19926)

tmp19928 := GCons(__e, tmp19927, Nil)

tmp19929 := GCons(__e, sym_1_1_6, tmp19928)

tmp19930 := GCons(__e, symsymbol, tmp19929)

return TailCall6(__e, SymFn(symis_b), V6599, tmp19930, B6595, L6596, Key6597, C6598)


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

tmp19931 := GValue(__e, symshen_4_dsigf_d)

tmp19932 := Call3(__e, SymFn(symshen_4assoc_1_6), symprofile_1results, tmp19924, tmp19931)


tmp19933 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19932)

_ = tmp19933

tmp19934 := MakeNative1(func(__e *ControlFlow, V6604 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6600 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6601 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6602 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6603 Obj) Obj {
tmp19935 := Call1(__e, SymFn(symshen_4newpv), B6600)


let__18052 := tmp19935
_ = let__18052

tmp19936 := GCons(__e, let__18052, Nil)

tmp19937 := GCons(__e, sym_1_1_6, tmp19936)

tmp19938 := GCons(__e, let__18052, tmp19937)

tmp19939 := Call6(__e, SymFn(symis_b), V6604, tmp19938, B6600, L6601, Key6602, C6603)


return TailCall2(__e, SymFn(symshen_4gc), B6600, tmp19939)


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

tmp19940 := GValue(__e, symshen_4_dsigf_d)

tmp19941 := Call3(__e, SymFn(symshen_4assoc_1_6), symprotect, tmp19934, tmp19940)


tmp19942 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19941)

_ = tmp19942

tmp19943 := MakeNative1(func(__e *ControlFlow, V6610 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6606 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6607 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6608 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6609 Obj) Obj {
tmp19944 := GCons(__e, symsymbol, Nil)

tmp19945 := GCons(__e, symlist, tmp19944)

tmp19946 := GCons(__e, symsymbol, Nil)

tmp19947 := GCons(__e, symlist, tmp19946)

tmp19948 := GCons(__e, tmp19947, Nil)

tmp19949 := GCons(__e, sym_1_1_6, tmp19948)

tmp19950 := GCons(__e, tmp19945, tmp19949)

return TailCall6(__e, SymFn(symis_b), V6610, tmp19950, B6606, L6607, Key6608, C6609)


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

tmp19951 := GValue(__e, symshen_4_dsigf_d)

tmp19952 := Call3(__e, SymFn(symshen_4assoc_1_6), sympreclude_1all_1but, tmp19943, tmp19951)


tmp19953 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19952)

_ = tmp19953

tmp19954 := MakeNative1(func(__e *ControlFlow, V6615 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6611 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6612 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6613 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6614 Obj) Obj {
tmp19955 := GCons(__e, symout, Nil)

tmp19956 := GCons(__e, symstream, tmp19955)

tmp19957 := GCons(__e, symstring, Nil)

tmp19958 := GCons(__e, sym_1_1_6, tmp19957)

tmp19959 := GCons(__e, tmp19956, tmp19958)

tmp19960 := GCons(__e, tmp19959, Nil)

tmp19961 := GCons(__e, sym_1_1_6, tmp19960)

tmp19962 := GCons(__e, symstring, tmp19961)

return TailCall6(__e, SymFn(symis_b), V6615, tmp19962, B6611, L6612, Key6613, C6614)


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

tmp19963 := GValue(__e, symshen_4_dsigf_d)

tmp19964 := Call3(__e, SymFn(symshen_4assoc_1_6), symshen_4prhush, tmp19954, tmp19963)


tmp19965 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19964)

_ = tmp19965

tmp19966 := MakeNative1(func(__e *ControlFlow, V6620 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6616 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6617 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6618 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6619 Obj) Obj {
tmp19967 := GCons(__e, symnumber, Nil)

tmp19968 := GCons(__e, sym_1_1_6, tmp19967)

tmp19969 := GCons(__e, symnumber, tmp19968)

return TailCall6(__e, SymFn(symis_b), V6620, tmp19969, B6616, L6617, Key6618, C6619)


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

tmp19971 := Call3(__e, SymFn(symshen_4assoc_1_6), symprolog_1memory, tmp19966, tmp19970)


tmp19972 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19971)

_ = tmp19972

tmp19973 := MakeNative1(func(__e *ControlFlow, V6625 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6621 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6622 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6623 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6624 Obj) Obj {
tmp19974 := GCons(__e, symunit, Nil)

tmp19975 := GCons(__e, symlist, tmp19974)

tmp19976 := GCons(__e, tmp19975, Nil)

tmp19977 := GCons(__e, sym_1_1_6, tmp19976)

tmp19978 := GCons(__e, symsymbol, tmp19977)

return TailCall6(__e, SymFn(symis_b), V6625, tmp19978, B6621, L6622, Key6623, C6624)


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

tmp19979 := GValue(__e, symshen_4_dsigf_d)

tmp19980 := Call3(__e, SymFn(symshen_4assoc_1_6), symps, tmp19973, tmp19979)


tmp19981 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19980)

_ = tmp19981

tmp19982 := MakeNative1(func(__e *ControlFlow, V6630 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6626 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6627 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6628 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6629 Obj) Obj {
tmp19983 := GCons(__e, symin, Nil)

tmp19984 := GCons(__e, symstream, tmp19983)

tmp19985 := GCons(__e, symunit, Nil)

tmp19986 := GCons(__e, sym_1_1_6, tmp19985)

tmp19987 := GCons(__e, tmp19984, tmp19986)

return TailCall6(__e, SymFn(symis_b), V6630, tmp19987, B6626, L6627, Key6628, C6629)


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

tmp19988 := GValue(__e, symshen_4_dsigf_d)

tmp19989 := Call3(__e, SymFn(symshen_4assoc_1_6), symread, tmp19982, tmp19988)


tmp19990 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19989)

_ = tmp19990

tmp19991 := MakeNative1(func(__e *ControlFlow, V6635 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6631 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6632 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6633 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6634 Obj) Obj {
tmp19992 := GCons(__e, symin, Nil)

tmp19993 := GCons(__e, symstream, tmp19992)

tmp19994 := GCons(__e, symstring, Nil)

tmp19995 := GCons(__e, sym_1_1_6, tmp19994)

tmp19996 := GCons(__e, tmp19993, tmp19995)

return TailCall6(__e, SymFn(symis_b), V6635, tmp19996, B6631, L6632, Key6633, C6634)


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

tmp19997 := GValue(__e, symshen_4_dsigf_d)

tmp19998 := Call3(__e, SymFn(symshen_4assoc_1_6), symread_1unit_1string, tmp19991, tmp19997)


tmp19999 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19998)

_ = tmp19999

tmp20000 := MakeNative1(func(__e *ControlFlow, V6640 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6636 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6637 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6638 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6639 Obj) Obj {
tmp20001 := GCons(__e, symin, Nil)

tmp20002 := GCons(__e, symstream, tmp20001)

tmp20003 := GCons(__e, symnumber, Nil)

tmp20004 := GCons(__e, sym_1_1_6, tmp20003)

tmp20005 := GCons(__e, tmp20002, tmp20004)

return TailCall6(__e, SymFn(symis_b), V6640, tmp20005, B6636, L6637, Key6638, C6639)


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

tmp20006 := GValue(__e, symshen_4_dsigf_d)

tmp20007 := Call3(__e, SymFn(symshen_4assoc_1_6), symread_1byte, tmp20000, tmp20006)


tmp20008 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20007)

_ = tmp20008

tmp20009 := MakeNative1(func(__e *ControlFlow, V6645 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6641 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6642 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6643 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6644 Obj) Obj {
tmp20010 := GCons(__e, symnumber, Nil)

tmp20011 := GCons(__e, symlist, tmp20010)

tmp20012 := GCons(__e, tmp20011, Nil)

tmp20013 := GCons(__e, sym_1_1_6, tmp20012)

tmp20014 := GCons(__e, symstring, tmp20013)

return TailCall6(__e, SymFn(symis_b), V6645, tmp20014, B6641, L6642, Key6643, C6644)


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

tmp20015 := GValue(__e, symshen_4_dsigf_d)

tmp20016 := Call3(__e, SymFn(symshen_4assoc_1_6), symread_1file_1as_1bytelist, tmp20009, tmp20015)


tmp20017 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20016)

_ = tmp20017

tmp20018 := MakeNative1(func(__e *ControlFlow, V6650 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6646 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6647 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6648 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6649 Obj) Obj {
tmp20019 := GCons(__e, symstring, Nil)

tmp20020 := GCons(__e, sym_1_1_6, tmp20019)

tmp20021 := GCons(__e, symstring, tmp20020)

return TailCall6(__e, SymFn(symis_b), V6650, tmp20021, B6646, L6647, Key6648, C6649)


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

tmp20022 := GValue(__e, symshen_4_dsigf_d)

tmp20023 := Call3(__e, SymFn(symshen_4assoc_1_6), symread_1file_1as_1string, tmp20018, tmp20022)


tmp20024 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20023)

_ = tmp20024

tmp20025 := MakeNative1(func(__e *ControlFlow, V6655 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6651 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6652 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6653 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6654 Obj) Obj {
tmp20026 := GCons(__e, symunit, Nil)

tmp20027 := GCons(__e, symlist, tmp20026)

tmp20028 := GCons(__e, tmp20027, Nil)

tmp20029 := GCons(__e, sym_1_1_6, tmp20028)

tmp20030 := GCons(__e, symstring, tmp20029)

return TailCall6(__e, SymFn(symis_b), V6655, tmp20030, B6651, L6652, Key6653, C6654)


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

tmp20032 := Call3(__e, SymFn(symshen_4assoc_1_6), symread_1file, tmp20025, tmp20031)


tmp20033 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20032)

_ = tmp20033

tmp20034 := MakeNative1(func(__e *ControlFlow, V6660 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6656 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6657 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6658 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6659 Obj) Obj {
tmp20035 := GCons(__e, symunit, Nil)

tmp20036 := GCons(__e, symlist, tmp20035)

tmp20037 := GCons(__e, tmp20036, Nil)

tmp20038 := GCons(__e, sym_1_1_6, tmp20037)

tmp20039 := GCons(__e, symstring, tmp20038)

return TailCall6(__e, SymFn(symis_b), V6660, tmp20039, B6656, L6657, Key6658, C6659)


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

tmp20040 := GValue(__e, symshen_4_dsigf_d)

tmp20041 := Call3(__e, SymFn(symshen_4assoc_1_6), symread_1from_1string, tmp20034, tmp20040)


tmp20042 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20041)

_ = tmp20042

tmp20043 := MakeNative1(func(__e *ControlFlow, V6665 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6661 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6662 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6663 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6664 Obj) Obj {
tmp20044 := GCons(__e, symunit, Nil)

tmp20045 := GCons(__e, symlist, tmp20044)

tmp20046 := GCons(__e, tmp20045, Nil)

tmp20047 := GCons(__e, sym_1_1_6, tmp20046)

tmp20048 := GCons(__e, symstring, tmp20047)

return TailCall6(__e, SymFn(symis_b), V6665, tmp20048, B6661, L6662, Key6663, C6664)


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

tmp20050 := Call3(__e, SymFn(symshen_4assoc_1_6), symread_1from_1string_1unprocessed, tmp20043, tmp20049)


tmp20051 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20050)

_ = tmp20051

tmp20052 := MakeNative1(func(__e *ControlFlow, V6670 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6666 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6667 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6668 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6669 Obj) Obj {
tmp20053 := GCons(__e, symstring, Nil)

tmp20054 := GCons(__e, sym_1_1_6, tmp20053)

return TailCall6(__e, SymFn(symis_b), V6670, tmp20054, B6666, L6667, Key6668, C6669)


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

tmp20055 := GValue(__e, symshen_4_dsigf_d)

tmp20056 := Call3(__e, SymFn(symshen_4assoc_1_6), symrelease, tmp20052, tmp20055)


tmp20057 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20056)

_ = tmp20057

tmp20058 := MakeNative1(func(__e *ControlFlow, V6676 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6672 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6673 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6674 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6675 Obj) Obj {
tmp20059 := Call1(__e, SymFn(symshen_4newpv), B6672)


let__18053 := tmp20059
_ = let__18053

tmp20060 := GCons(__e, let__18053, Nil)

tmp20061 := GCons(__e, symlist, tmp20060)

tmp20062 := GCons(__e, let__18053, Nil)

tmp20063 := GCons(__e, symlist, tmp20062)

tmp20064 := GCons(__e, tmp20063, Nil)

tmp20065 := GCons(__e, sym_1_1_6, tmp20064)

tmp20066 := GCons(__e, tmp20061, tmp20065)

tmp20067 := GCons(__e, tmp20066, Nil)

tmp20068 := GCons(__e, sym_1_1_6, tmp20067)

tmp20069 := GCons(__e, let__18053, tmp20068)

tmp20070 := Call6(__e, SymFn(symis_b), V6676, tmp20069, B6672, L6673, Key6674, C6675)


return TailCall2(__e, SymFn(symshen_4gc), B6672, tmp20070)


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

tmp20071 := GValue(__e, symshen_4_dsigf_d)

tmp20072 := Call3(__e, SymFn(symshen_4assoc_1_6), symremove, tmp20058, tmp20071)


tmp20073 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20072)

_ = tmp20073

tmp20074 := MakeNative1(func(__e *ControlFlow, V6682 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6678 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6679 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6680 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6681 Obj) Obj {
tmp20075 := Call1(__e, SymFn(symshen_4newpv), B6678)


let__18054 := tmp20075
_ = let__18054

tmp20076 := GCons(__e, let__18054, Nil)

tmp20077 := GCons(__e, symlist, tmp20076)

tmp20078 := GCons(__e, let__18054, Nil)

tmp20079 := GCons(__e, symlist, tmp20078)

tmp20080 := GCons(__e, tmp20079, Nil)

tmp20081 := GCons(__e, sym_1_1_6, tmp20080)

tmp20082 := GCons(__e, tmp20077, tmp20081)

tmp20083 := Call6(__e, SymFn(symis_b), V6682, tmp20082, B6678, L6679, Key6680, C6681)


return TailCall2(__e, SymFn(symshen_4gc), B6678, tmp20083)


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

tmp20084 := GValue(__e, symshen_4_dsigf_d)

tmp20085 := Call3(__e, SymFn(symshen_4assoc_1_6), symreverse, tmp20074, tmp20084)


tmp20086 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20085)

_ = tmp20086

tmp20087 := MakeNative1(func(__e *ControlFlow, V6687 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6683 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6684 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6685 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6686 Obj) Obj {
tmp20088 := Call1(__e, SymFn(symshen_4newpv), B6683)


let__18055 := tmp20088
_ = let__18055

tmp20089 := GCons(__e, let__18055, Nil)

tmp20090 := GCons(__e, sym_1_1_6, tmp20089)

tmp20091 := GCons(__e, symstring, tmp20090)

tmp20092 := Call6(__e, SymFn(symis_b), V6687, tmp20091, B6683, L6684, Key6685, C6686)


return TailCall2(__e, SymFn(symshen_4gc), B6683, tmp20092)


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

tmp20093 := GValue(__e, symshen_4_dsigf_d)

tmp20094 := Call3(__e, SymFn(symshen_4assoc_1_6), symsimple_1error, tmp20087, tmp20093)


tmp20095 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20094)

_ = tmp20095

tmp20096 := MakeNative1(func(__e *ControlFlow, V6692 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6688 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6689 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6690 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6691 Obj) Obj {
tmp20097 := Call1(__e, SymFn(symshen_4newpv), B6688)


let__18056 := tmp20097
_ = let__18056

tmp20098 := Call1(__e, SymFn(symshen_4newpv), B6688)


let__18057 := tmp20098
_ = let__18057

tmp20099 := GCons(__e, let__18057, Nil)

tmp20100 := GCons(__e, sym_d, tmp20099)

tmp20101 := GCons(__e, let__18056, tmp20100)

tmp20102 := GCons(__e, let__18057, Nil)

tmp20103 := GCons(__e, sym_1_1_6, tmp20102)

tmp20104 := GCons(__e, tmp20101, tmp20103)

tmp20105 := Call6(__e, SymFn(symis_b), V6692, tmp20104, B6688, L6689, Key6690, C6691)


tmp20106 := Call2(__e, SymFn(symshen_4gc), B6688, tmp20105)


return TailCall2(__e, SymFn(symshen_4gc), B6688, tmp20106)


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

tmp20107 := GValue(__e, symshen_4_dsigf_d)

tmp20108 := Call3(__e, SymFn(symshen_4assoc_1_6), symsnd, tmp20096, tmp20107)


tmp20109 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20108)

_ = tmp20109

tmp20110 := MakeNative1(func(__e *ControlFlow, V6697 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6693 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6694 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6695 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6696 Obj) Obj {
tmp20111 := GCons(__e, symsymbol, Nil)

tmp20112 := GCons(__e, sym_1_1_6, tmp20111)

tmp20113 := GCons(__e, symnumber, tmp20112)

tmp20114 := GCons(__e, tmp20113, Nil)

tmp20115 := GCons(__e, sym_1_1_6, tmp20114)

tmp20116 := GCons(__e, symsymbol, tmp20115)

return TailCall6(__e, SymFn(symis_b), V6697, tmp20116, B6693, L6694, Key6695, C6696)


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

tmp20117 := GValue(__e, symshen_4_dsigf_d)

tmp20118 := Call3(__e, SymFn(symshen_4assoc_1_6), symspecialise, tmp20110, tmp20117)


tmp20119 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20118)

_ = tmp20119

tmp20120 := MakeNative1(func(__e *ControlFlow, V6702 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6698 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6699 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6700 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6701 Obj) Obj {
tmp20121 := GCons(__e, symboolean, Nil)

tmp20122 := GCons(__e, sym_1_1_6, tmp20121)

tmp20123 := GCons(__e, symsymbol, tmp20122)

return TailCall6(__e, SymFn(symis_b), V6702, tmp20123, B6698, L6699, Key6700, C6701)


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

tmp20124 := GValue(__e, symshen_4_dsigf_d)

tmp20125 := Call3(__e, SymFn(symshen_4assoc_1_6), symspy, tmp20120, tmp20124)


tmp20126 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20125)

_ = tmp20126

tmp20127 := MakeNative1(func(__e *ControlFlow, V6707 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6703 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6704 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6705 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6706 Obj) Obj {
tmp20128 := GCons(__e, symboolean, Nil)

tmp20129 := GCons(__e, sym_1_1_6, tmp20128)

return TailCall6(__e, SymFn(symis_b), V6707, tmp20129, B6703, L6704, Key6705, C6706)


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

tmp20130 := GValue(__e, symshen_4_dsigf_d)

tmp20131 := Call3(__e, SymFn(symshen_4assoc_1_6), symspy_2, tmp20127, tmp20130)


tmp20132 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20131)

_ = tmp20132

tmp20133 := MakeNative1(func(__e *ControlFlow, V6712 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6708 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6709 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6710 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6711 Obj) Obj {
tmp20134 := GCons(__e, symboolean, Nil)

tmp20135 := GCons(__e, sym_1_1_6, tmp20134)

tmp20136 := GCons(__e, symsymbol, tmp20135)

return TailCall6(__e, SymFn(symis_b), V6712, tmp20136, B6708, L6709, Key6710, C6711)


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

tmp20137 := GValue(__e, symshen_4_dsigf_d)

tmp20138 := Call3(__e, SymFn(symshen_4assoc_1_6), symstep, tmp20133, tmp20137)


tmp20139 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20138)

_ = tmp20139

tmp20140 := MakeNative1(func(__e *ControlFlow, V6717 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6713 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6714 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6715 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6716 Obj) Obj {
tmp20141 := GCons(__e, symboolean, Nil)

tmp20142 := GCons(__e, sym_1_1_6, tmp20141)

return TailCall6(__e, SymFn(symis_b), V6717, tmp20142, B6713, L6714, Key6715, C6716)


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

tmp20143 := GValue(__e, symshen_4_dsigf_d)

tmp20144 := Call3(__e, SymFn(symshen_4assoc_1_6), symstep_2, tmp20140, tmp20143)


tmp20145 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20144)

_ = tmp20145

tmp20146 := MakeNative1(func(__e *ControlFlow, V6722 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6718 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6719 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6720 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6721 Obj) Obj {
tmp20147 := GCons(__e, symin, Nil)

tmp20148 := GCons(__e, symstream, tmp20147)

tmp20149 := GCons(__e, tmp20148, Nil)

tmp20150 := GCons(__e, sym_1_1_6, tmp20149)

return TailCall6(__e, SymFn(symis_b), V6722, tmp20150, B6718, L6719, Key6720, C6721)


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

tmp20151 := GValue(__e, symshen_4_dsigf_d)

tmp20152 := Call3(__e, SymFn(symshen_4assoc_1_6), symstinput, tmp20146, tmp20151)


tmp20153 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20152)

_ = tmp20153

tmp20154 := MakeNative1(func(__e *ControlFlow, V6727 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6723 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6724 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6725 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6726 Obj) Obj {
tmp20155 := GCons(__e, symout, Nil)

tmp20156 := GCons(__e, symstream, tmp20155)

tmp20157 := GCons(__e, tmp20156, Nil)

tmp20158 := GCons(__e, sym_1_1_6, tmp20157)

return TailCall6(__e, SymFn(symis_b), V6727, tmp20158, B6723, L6724, Key6725, C6726)


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

tmp20159 := GValue(__e, symshen_4_dsigf_d)

tmp20160 := Call3(__e, SymFn(symshen_4assoc_1_6), symsterror, tmp20154, tmp20159)


tmp20161 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20160)

_ = tmp20161

tmp20162 := MakeNative1(func(__e *ControlFlow, V6732 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6728 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6729 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6730 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6731 Obj) Obj {
tmp20163 := GCons(__e, symout, Nil)

tmp20164 := GCons(__e, symstream, tmp20163)

tmp20165 := GCons(__e, tmp20164, Nil)

tmp20166 := GCons(__e, sym_1_1_6, tmp20165)

return TailCall6(__e, SymFn(symis_b), V6732, tmp20166, B6728, L6729, Key6730, C6731)


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

tmp20167 := GValue(__e, symshen_4_dsigf_d)

tmp20168 := Call3(__e, SymFn(symshen_4assoc_1_6), symstoutput, tmp20162, tmp20167)


tmp20169 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20168)

_ = tmp20169

tmp20170 := MakeNative1(func(__e *ControlFlow, V6737 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6733 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6734 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6735 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6736 Obj) Obj {
tmp20171 := Call1(__e, SymFn(symshen_4newpv), B6733)


let__18058 := tmp20171
_ = let__18058

tmp20172 := GCons(__e, symboolean, Nil)

tmp20173 := GCons(__e, sym_1_1_6, tmp20172)

tmp20174 := GCons(__e, let__18058, tmp20173)

tmp20175 := Call6(__e, SymFn(symis_b), V6737, tmp20174, B6733, L6734, Key6735, C6736)


return TailCall2(__e, SymFn(symshen_4gc), B6733, tmp20175)


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

tmp20176 := GValue(__e, symshen_4_dsigf_d)

tmp20177 := Call3(__e, SymFn(symshen_4assoc_1_6), symstring_2, tmp20170, tmp20176)


tmp20178 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20177)

_ = tmp20178

tmp20179 := MakeNative1(func(__e *ControlFlow, V6742 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6738 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6739 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6740 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6741 Obj) Obj {
tmp20180 := Call1(__e, SymFn(symshen_4newpv), B6738)


let__18059 := tmp20180
_ = let__18059

tmp20181 := GCons(__e, symstring, Nil)

tmp20182 := GCons(__e, sym_1_1_6, tmp20181)

tmp20183 := GCons(__e, let__18059, tmp20182)

tmp20184 := Call6(__e, SymFn(symis_b), V6742, tmp20183, B6738, L6739, Key6740, C6741)


return TailCall2(__e, SymFn(symshen_4gc), B6738, tmp20184)


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

tmp20185 := GValue(__e, symshen_4_dsigf_d)

tmp20186 := Call3(__e, SymFn(symshen_4assoc_1_6), symstr, tmp20179, tmp20185)


tmp20187 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20186)

_ = tmp20187

tmp20188 := MakeNative1(func(__e *ControlFlow, V6747 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6743 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6744 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6745 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6746 Obj) Obj {
tmp20189 := GCons(__e, symnumber, Nil)

tmp20190 := GCons(__e, sym_1_1_6, tmp20189)

tmp20191 := GCons(__e, symstring, tmp20190)

return TailCall6(__e, SymFn(symis_b), V6747, tmp20191, B6743, L6744, Key6745, C6746)


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

tmp20192 := GValue(__e, symshen_4_dsigf_d)

tmp20193 := Call3(__e, SymFn(symshen_4assoc_1_6), symstring_1_6n, tmp20188, tmp20192)


tmp20194 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20193)

_ = tmp20194

tmp20195 := MakeNative1(func(__e *ControlFlow, V6752 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6748 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6749 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6750 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6751 Obj) Obj {
tmp20196 := GCons(__e, symsymbol, Nil)

tmp20197 := GCons(__e, sym_1_1_6, tmp20196)

tmp20198 := GCons(__e, symstring, tmp20197)

return TailCall6(__e, SymFn(symis_b), V6752, tmp20198, B6748, L6749, Key6750, C6751)


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

tmp20199 := GValue(__e, symshen_4_dsigf_d)

tmp20200 := Call3(__e, SymFn(symshen_4assoc_1_6), symstring_1_6symbol, tmp20195, tmp20199)


tmp20201 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20200)

_ = tmp20201

tmp20202 := MakeNative1(func(__e *ControlFlow, V6757 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6753 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6754 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6755 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6756 Obj) Obj {
tmp20203 := GCons(__e, symnumber, Nil)

tmp20204 := GCons(__e, symlist, tmp20203)

tmp20205 := GCons(__e, symnumber, Nil)

tmp20206 := GCons(__e, sym_1_1_6, tmp20205)

tmp20207 := GCons(__e, tmp20204, tmp20206)

return TailCall6(__e, SymFn(symis_b), V6757, tmp20207, B6753, L6754, Key6755, C6756)


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

tmp20208 := GValue(__e, symshen_4_dsigf_d)

tmp20209 := Call3(__e, SymFn(symshen_4assoc_1_6), symsum, tmp20202, tmp20208)


tmp20210 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20209)

_ = tmp20210

tmp20211 := MakeNative1(func(__e *ControlFlow, V6762 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6758 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6759 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6760 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6761 Obj) Obj {
tmp20212 := Call1(__e, SymFn(symshen_4newpv), B6758)


let__18060 := tmp20212
_ = let__18060

tmp20213 := GCons(__e, symboolean, Nil)

tmp20214 := GCons(__e, sym_1_1_6, tmp20213)

tmp20215 := GCons(__e, let__18060, tmp20214)

tmp20216 := Call6(__e, SymFn(symis_b), V6762, tmp20215, B6758, L6759, Key6760, C6761)


return TailCall2(__e, SymFn(symshen_4gc), B6758, tmp20216)


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

tmp20217 := GValue(__e, symshen_4_dsigf_d)

tmp20218 := Call3(__e, SymFn(symshen_4assoc_1_6), symsymbol_2, tmp20211, tmp20217)


tmp20219 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20218)

_ = tmp20219

tmp20220 := MakeNative1(func(__e *ControlFlow, V6767 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6763 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6764 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6765 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6766 Obj) Obj {
tmp20221 := GCons(__e, symsymbol, Nil)

tmp20222 := GCons(__e, sym_1_1_6, tmp20221)

tmp20223 := GCons(__e, symsymbol, tmp20222)

return TailCall6(__e, SymFn(symis_b), V6767, tmp20223, B6763, L6764, Key6765, C6766)


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

tmp20224 := GValue(__e, symshen_4_dsigf_d)

tmp20225 := Call3(__e, SymFn(symshen_4assoc_1_6), symsystemf, tmp20220, tmp20224)


tmp20226 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20225)

_ = tmp20226

tmp20227 := MakeNative1(func(__e *ControlFlow, V6772 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6768 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6769 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6770 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6771 Obj) Obj {
tmp20228 := GCons(__e, symboolean, Nil)

tmp20229 := GCons(__e, sym_1_1_6, tmp20228)

return TailCall6(__e, SymFn(symis_b), V6772, tmp20229, B6768, L6769, Key6770, C6771)


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

tmp20230 := GValue(__e, symshen_4_dsigf_d)

tmp20231 := Call3(__e, SymFn(symshen_4assoc_1_6), symsystem_1S_2, tmp20227, tmp20230)


tmp20232 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20231)

_ = tmp20232

tmp20233 := MakeNative1(func(__e *ControlFlow, V6778 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6774 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6775 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6776 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6777 Obj) Obj {
tmp20234 := Call1(__e, SymFn(symshen_4newpv), B6774)


let__18061 := tmp20234
_ = let__18061

tmp20235 := GCons(__e, let__18061, Nil)

tmp20236 := GCons(__e, symlist, tmp20235)

tmp20237 := GCons(__e, let__18061, Nil)

tmp20238 := GCons(__e, symlist, tmp20237)

tmp20239 := GCons(__e, tmp20238, Nil)

tmp20240 := GCons(__e, sym_1_1_6, tmp20239)

tmp20241 := GCons(__e, tmp20236, tmp20240)

tmp20242 := Call6(__e, SymFn(symis_b), V6778, tmp20241, B6774, L6775, Key6776, C6777)


return TailCall2(__e, SymFn(symshen_4gc), B6774, tmp20242)


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

tmp20243 := GValue(__e, symshen_4_dsigf_d)

tmp20244 := Call3(__e, SymFn(symshen_4assoc_1_6), symtail, tmp20233, tmp20243)


tmp20245 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20244)

_ = tmp20245

tmp20246 := MakeNative1(func(__e *ControlFlow, V6783 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6779 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6780 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6781 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6782 Obj) Obj {
tmp20247 := GCons(__e, symstring, Nil)

tmp20248 := GCons(__e, sym_1_1_6, tmp20247)

tmp20249 := GCons(__e, symstring, tmp20248)

return TailCall6(__e, SymFn(symis_b), V6783, tmp20249, B6779, L6780, Key6781, C6782)


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

tmp20250 := GValue(__e, symshen_4_dsigf_d)

tmp20251 := Call3(__e, SymFn(symshen_4assoc_1_6), symtlstr, tmp20246, tmp20250)


tmp20252 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20251)

_ = tmp20252

tmp20253 := MakeNative1(func(__e *ControlFlow, V6789 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6785 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6786 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6787 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6788 Obj) Obj {
tmp20254 := Call1(__e, SymFn(symshen_4newpv), B6785)


let__18062 := tmp20254
_ = let__18062

tmp20255 := GCons(__e, let__18062, Nil)

tmp20256 := GCons(__e, symvector, tmp20255)

tmp20257 := GCons(__e, let__18062, Nil)

tmp20258 := GCons(__e, symvector, tmp20257)

tmp20259 := GCons(__e, tmp20258, Nil)

tmp20260 := GCons(__e, sym_1_1_6, tmp20259)

tmp20261 := GCons(__e, tmp20256, tmp20260)

tmp20262 := Call6(__e, SymFn(symis_b), V6789, tmp20261, B6785, L6786, Key6787, C6788)


return TailCall2(__e, SymFn(symshen_4gc), B6785, tmp20262)


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

tmp20263 := GValue(__e, symshen_4_dsigf_d)

tmp20264 := Call3(__e, SymFn(symshen_4assoc_1_6), symtlv, tmp20253, tmp20263)


tmp20265 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20264)

_ = tmp20265

tmp20266 := MakeNative1(func(__e *ControlFlow, V6794 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6790 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6791 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6792 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6793 Obj) Obj {
tmp20267 := GCons(__e, symboolean, Nil)

tmp20268 := GCons(__e, sym_1_1_6, tmp20267)

tmp20269 := GCons(__e, symsymbol, tmp20268)

return TailCall6(__e, SymFn(symis_b), V6794, tmp20269, B6790, L6791, Key6792, C6793)


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

tmp20270 := GValue(__e, symshen_4_dsigf_d)

tmp20271 := Call3(__e, SymFn(symshen_4assoc_1_6), symtc, tmp20266, tmp20270)


tmp20272 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20271)

_ = tmp20272

tmp20273 := MakeNative1(func(__e *ControlFlow, V6799 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6795 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6796 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6797 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6798 Obj) Obj {
tmp20274 := GCons(__e, symboolean, Nil)

tmp20275 := GCons(__e, sym_1_1_6, tmp20274)

return TailCall6(__e, SymFn(symis_b), V6799, tmp20275, B6795, L6796, Key6797, C6798)


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

tmp20276 := GValue(__e, symshen_4_dsigf_d)

tmp20277 := Call3(__e, SymFn(symshen_4assoc_1_6), symtc_2, tmp20273, tmp20276)


tmp20278 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20277)

_ = tmp20278

tmp20279 := MakeNative1(func(__e *ControlFlow, V6804 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6800 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6801 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6802 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6803 Obj) Obj {
tmp20280 := Call1(__e, SymFn(symshen_4newpv), B6800)


let__18063 := tmp20280
_ = let__18063

tmp20281 := GCons(__e, let__18063, Nil)

tmp20282 := GCons(__e, symlazy, tmp20281)

tmp20283 := GCons(__e, let__18063, Nil)

tmp20284 := GCons(__e, sym_1_1_6, tmp20283)

tmp20285 := GCons(__e, tmp20282, tmp20284)

tmp20286 := Call6(__e, SymFn(symis_b), V6804, tmp20285, B6800, L6801, Key6802, C6803)


return TailCall2(__e, SymFn(symshen_4gc), B6800, tmp20286)


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

tmp20287 := GValue(__e, symshen_4_dsigf_d)

tmp20288 := Call3(__e, SymFn(symshen_4assoc_1_6), symthaw, tmp20279, tmp20287)


tmp20289 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20288)

_ = tmp20289

tmp20290 := MakeNative1(func(__e *ControlFlow, V6809 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6805 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6806 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6807 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6808 Obj) Obj {
tmp20291 := GCons(__e, symsymbol, Nil)

tmp20292 := GCons(__e, sym_1_1_6, tmp20291)

tmp20293 := GCons(__e, symsymbol, tmp20292)

return TailCall6(__e, SymFn(symis_b), V6809, tmp20293, B6805, L6806, Key6807, C6808)


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

tmp20294 := GValue(__e, symshen_4_dsigf_d)

tmp20295 := Call3(__e, SymFn(symshen_4assoc_1_6), symtrack, tmp20290, tmp20294)


tmp20296 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20295)

_ = tmp20296

tmp20297 := MakeNative1(func(__e *ControlFlow, V6814 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6810 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6811 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6812 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6813 Obj) Obj {
tmp20298 := GCons(__e, symsymbol, Nil)

tmp20299 := GCons(__e, symlist, tmp20298)

tmp20300 := GCons(__e, tmp20299, Nil)

tmp20301 := GCons(__e, sym_1_1_6, tmp20300)

return TailCall6(__e, SymFn(symis_b), V6814, tmp20301, B6810, L6811, Key6812, C6813)


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

tmp20302 := GValue(__e, symshen_4_dsigf_d)

tmp20303 := Call3(__e, SymFn(symshen_4assoc_1_6), symtracked, tmp20297, tmp20302)


tmp20304 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20303)

_ = tmp20304

tmp20305 := MakeNative1(func(__e *ControlFlow, V6819 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6815 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6816 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6817 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6818 Obj) Obj {
tmp20306 := Call1(__e, SymFn(symshen_4newpv), B6815)


let__18064 := tmp20306
_ = let__18064

tmp20307 := GCons(__e, let__18064, Nil)

tmp20308 := GCons(__e, sym_1_1_6, tmp20307)

tmp20309 := GCons(__e, symexception, tmp20308)

tmp20310 := GCons(__e, let__18064, Nil)

tmp20311 := GCons(__e, sym_1_1_6, tmp20310)

tmp20312 := GCons(__e, tmp20309, tmp20311)

tmp20313 := GCons(__e, tmp20312, Nil)

tmp20314 := GCons(__e, sym_1_1_6, tmp20313)

tmp20315 := GCons(__e, let__18064, tmp20314)

tmp20316 := Call6(__e, SymFn(symis_b), V6819, tmp20315, B6815, L6816, Key6817, C6818)


return TailCall2(__e, SymFn(symshen_4gc), B6815, tmp20316)


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

tmp20317 := GValue(__e, symshen_4_dsigf_d)

tmp20318 := Call3(__e, SymFn(symshen_4assoc_1_6), symtrap_1error, tmp20305, tmp20317)


tmp20319 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20318)

_ = tmp20319

tmp20320 := MakeNative1(func(__e *ControlFlow, V6824 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6820 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6821 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6822 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6823 Obj) Obj {
tmp20321 := Call1(__e, SymFn(symshen_4newpv), B6820)


let__18065 := tmp20321
_ = let__18065

tmp20322 := GCons(__e, symboolean, Nil)

tmp20323 := GCons(__e, sym_1_1_6, tmp20322)

tmp20324 := GCons(__e, let__18065, tmp20323)

tmp20325 := Call6(__e, SymFn(symis_b), V6824, tmp20324, B6820, L6821, Key6822, C6823)


return TailCall2(__e, SymFn(symshen_4gc), B6820, tmp20325)


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

tmp20326 := GValue(__e, symshen_4_dsigf_d)

tmp20327 := Call3(__e, SymFn(symshen_4assoc_1_6), symtuple_2, tmp20320, tmp20326)


tmp20328 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20327)

_ = tmp20328

tmp20329 := MakeNative1(func(__e *ControlFlow, V6829 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6825 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6826 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6827 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6828 Obj) Obj {
tmp20330 := GCons(__e, symstring, Nil)

tmp20331 := GCons(__e, symlist, tmp20330)

tmp20332 := GCons(__e, tmp20331, Nil)

tmp20333 := GCons(__e, sym_1_1_6, tmp20332)

tmp20334 := GCons(__e, symstring, tmp20333)

return TailCall6(__e, SymFn(symis_b), V6829, tmp20334, B6825, L6826, Key6827, C6828)


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

tmp20335 := GValue(__e, symshen_4_dsigf_d)

tmp20336 := Call3(__e, SymFn(symshen_4assoc_1_6), symunabsolute, tmp20329, tmp20335)


tmp20337 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20336)

_ = tmp20337

tmp20338 := MakeNative1(func(__e *ControlFlow, V6834 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6830 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6831 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6832 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6833 Obj) Obj {
tmp20339 := GCons(__e, symsymbol, Nil)

tmp20340 := GCons(__e, sym_1_1_6, tmp20339)

tmp20341 := GCons(__e, symsymbol, tmp20340)

return TailCall6(__e, SymFn(symis_b), V6834, tmp20341, B6830, L6831, Key6832, C6833)


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

tmp20342 := GValue(__e, symshen_4_dsigf_d)

tmp20343 := Call3(__e, SymFn(symshen_4assoc_1_6), symundefmacro, tmp20338, tmp20342)


tmp20344 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20343)

_ = tmp20344

tmp20345 := MakeNative1(func(__e *ControlFlow, V6841 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6837 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6838 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6839 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6840 Obj) Obj {
tmp20346 := Call1(__e, SymFn(symshen_4newpv), B6837)


let__18066 := tmp20346
_ = let__18066

tmp20347 := GCons(__e, let__18066, Nil)

tmp20348 := GCons(__e, symlist, tmp20347)

tmp20349 := GCons(__e, let__18066, Nil)

tmp20350 := GCons(__e, symlist, tmp20349)

tmp20351 := GCons(__e, let__18066, Nil)

tmp20352 := GCons(__e, symlist, tmp20351)

tmp20353 := GCons(__e, tmp20352, Nil)

tmp20354 := GCons(__e, sym_1_1_6, tmp20353)

tmp20355 := GCons(__e, tmp20350, tmp20354)

tmp20356 := GCons(__e, tmp20355, Nil)

tmp20357 := GCons(__e, sym_1_1_6, tmp20356)

tmp20358 := GCons(__e, tmp20348, tmp20357)

tmp20359 := Call6(__e, SymFn(symis_b), V6841, tmp20358, B6837, L6838, Key6839, C6840)


return TailCall2(__e, SymFn(symshen_4gc), B6837, tmp20359)


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

tmp20360 := GValue(__e, symshen_4_dsigf_d)

tmp20361 := Call3(__e, SymFn(symshen_4assoc_1_6), symunion, tmp20345, tmp20360)


tmp20362 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20361)

_ = tmp20362

tmp20363 := MakeNative1(func(__e *ControlFlow, V6846 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6842 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6843 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6844 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6845 Obj) Obj {
tmp20364 := GCons(__e, symsymbol, Nil)

tmp20365 := GCons(__e, sym_1_1_6, tmp20364)

tmp20366 := GCons(__e, symsymbol, tmp20365)

return TailCall6(__e, SymFn(symis_b), V6846, tmp20366, B6842, L6843, Key6844, C6845)


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

tmp20367 := GValue(__e, symshen_4_dsigf_d)

tmp20368 := Call3(__e, SymFn(symshen_4assoc_1_6), symunprofile, tmp20363, tmp20367)


tmp20369 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20368)

_ = tmp20369

tmp20370 := MakeNative1(func(__e *ControlFlow, V6851 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6847 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6848 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6849 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6850 Obj) Obj {
tmp20371 := GCons(__e, symsymbol, Nil)

tmp20372 := GCons(__e, sym_1_1_6, tmp20371)

tmp20373 := GCons(__e, symsymbol, tmp20372)

return TailCall6(__e, SymFn(symis_b), V6851, tmp20373, B6847, L6848, Key6849, C6850)


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

tmp20374 := GValue(__e, symshen_4_dsigf_d)

tmp20375 := Call3(__e, SymFn(symshen_4assoc_1_6), symuntrack, tmp20370, tmp20374)


tmp20376 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20375)

_ = tmp20376

tmp20377 := MakeNative1(func(__e *ControlFlow, V6856 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6852 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6853 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6854 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6855 Obj) Obj {
tmp20378 := GCons(__e, symsymbol, Nil)

tmp20379 := GCons(__e, symlist, tmp20378)

tmp20380 := GCons(__e, tmp20379, Nil)

tmp20381 := GCons(__e, sym_1_1_6, tmp20380)

return TailCall6(__e, SymFn(symis_b), V6856, tmp20381, B6852, L6853, Key6854, C6855)


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

tmp20382 := GValue(__e, symshen_4_dsigf_d)

tmp20383 := Call3(__e, SymFn(symshen_4assoc_1_6), symuserdefs, tmp20377, tmp20382)


tmp20384 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20383)

_ = tmp20384

tmp20385 := MakeNative1(func(__e *ControlFlow, V6861 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6857 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6858 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6859 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6860 Obj) Obj {
tmp20386 := Call1(__e, SymFn(symshen_4newpv), B6857)


let__18067 := tmp20386
_ = let__18067

tmp20387 := GCons(__e, symboolean, Nil)

tmp20388 := GCons(__e, sym_1_1_6, tmp20387)

tmp20389 := GCons(__e, let__18067, tmp20388)

tmp20390 := Call6(__e, SymFn(symis_b), V6861, tmp20389, B6857, L6858, Key6859, C6860)


return TailCall2(__e, SymFn(symshen_4gc), B6857, tmp20390)


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

tmp20391 := GValue(__e, symshen_4_dsigf_d)

tmp20392 := Call3(__e, SymFn(symshen_4assoc_1_6), symvariable_2, tmp20385, tmp20391)


tmp20393 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20392)

_ = tmp20393

tmp20394 := MakeNative1(func(__e *ControlFlow, V6866 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6862 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6863 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6864 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6865 Obj) Obj {
tmp20395 := Call1(__e, SymFn(symshen_4newpv), B6862)


let__18068 := tmp20395
_ = let__18068

tmp20396 := GCons(__e, symboolean, Nil)

tmp20397 := GCons(__e, sym_1_1_6, tmp20396)

tmp20398 := GCons(__e, let__18068, tmp20397)

tmp20399 := Call6(__e, SymFn(symis_b), V6866, tmp20398, B6862, L6863, Key6864, C6865)


return TailCall2(__e, SymFn(symshen_4gc), B6862, tmp20399)


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

tmp20400 := GValue(__e, symshen_4_dsigf_d)

tmp20401 := Call3(__e, SymFn(symshen_4assoc_1_6), symvector_2, tmp20394, tmp20400)


tmp20402 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20401)

_ = tmp20402

tmp20403 := MakeNative1(func(__e *ControlFlow, V6871 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6867 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6868 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6869 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6870 Obj) Obj {
tmp20404 := GCons(__e, symstring, Nil)

tmp20405 := GCons(__e, sym_1_1_6, tmp20404)

return TailCall6(__e, SymFn(symis_b), V6871, tmp20405, B6867, L6868, Key6869, C6870)


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

tmp20406 := GValue(__e, symshen_4_dsigf_d)

tmp20407 := Call3(__e, SymFn(symshen_4assoc_1_6), symversion, tmp20403, tmp20406)


tmp20408 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20407)

_ = tmp20408

tmp20409 := MakeNative1(func(__e *ControlFlow, V6876 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6872 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6873 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6874 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6875 Obj) Obj {
tmp20410 := Call1(__e, SymFn(symshen_4newpv), B6872)


let__18069 := tmp20410
_ = let__18069

tmp20411 := GCons(__e, let__18069, Nil)

tmp20412 := GCons(__e, sym_1_1_6, tmp20411)

tmp20413 := GCons(__e, let__18069, tmp20412)

tmp20414 := GCons(__e, tmp20413, Nil)

tmp20415 := GCons(__e, sym_1_1_6, tmp20414)

tmp20416 := GCons(__e, symstring, tmp20415)

tmp20417 := Call6(__e, SymFn(symis_b), V6876, tmp20416, B6872, L6873, Key6874, C6875)


return TailCall2(__e, SymFn(symshen_4gc), B6872, tmp20417)


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

tmp20418 := GValue(__e, symshen_4_dsigf_d)

tmp20419 := Call3(__e, SymFn(symshen_4assoc_1_6), symwrite_1to_1file, tmp20409, tmp20418)


tmp20420 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20419)

_ = tmp20420

tmp20421 := MakeNative1(func(__e *ControlFlow, V6881 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6877 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6878 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6879 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6880 Obj) Obj {
tmp20422 := GCons(__e, symout, Nil)

tmp20423 := GCons(__e, symstream, tmp20422)

tmp20424 := GCons(__e, symnumber, Nil)

tmp20425 := GCons(__e, sym_1_1_6, tmp20424)

tmp20426 := GCons(__e, tmp20423, tmp20425)

tmp20427 := GCons(__e, tmp20426, Nil)

tmp20428 := GCons(__e, sym_1_1_6, tmp20427)

tmp20429 := GCons(__e, symnumber, tmp20428)

return TailCall6(__e, SymFn(symis_b), V6881, tmp20429, B6877, L6878, Key6879, C6880)


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

tmp20430 := GValue(__e, symshen_4_dsigf_d)

tmp20431 := Call3(__e, SymFn(symshen_4assoc_1_6), symwrite_1byte, tmp20421, tmp20430)


tmp20432 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20431)

_ = tmp20432

tmp20433 := MakeNative1(func(__e *ControlFlow, V6886 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6882 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6883 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6884 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6885 Obj) Obj {
tmp20434 := GCons(__e, symout, Nil)

tmp20435 := GCons(__e, symstream, tmp20434)

tmp20436 := GCons(__e, symstring, Nil)

tmp20437 := GCons(__e, sym_1_1_6, tmp20436)

tmp20438 := GCons(__e, tmp20435, tmp20437)

tmp20439 := GCons(__e, tmp20438, Nil)

tmp20440 := GCons(__e, sym_1_1_6, tmp20439)

tmp20441 := GCons(__e, symstring, tmp20440)

return TailCall6(__e, SymFn(symis_b), V6886, tmp20441, B6882, L6883, Key6884, C6885)


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

tmp20442 := GValue(__e, symshen_4_dsigf_d)

tmp20443 := Call3(__e, SymFn(symshen_4assoc_1_6), symwrite_1string, tmp20433, tmp20442)


tmp20444 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20443)

_ = tmp20444

tmp20445 := MakeNative1(func(__e *ControlFlow, V6891 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6887 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6888 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6889 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6890 Obj) Obj {
tmp20446 := GCons(__e, symboolean, Nil)

tmp20447 := GCons(__e, sym_1_1_6, tmp20446)

tmp20448 := GCons(__e, symstring, tmp20447)

return TailCall6(__e, SymFn(symis_b), V6891, tmp20448, B6887, L6888, Key6889, C6890)


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

tmp20449 := GValue(__e, symshen_4_dsigf_d)

tmp20450 := Call3(__e, SymFn(symshen_4assoc_1_6), symy_1or_1n_2, tmp20445, tmp20449)


tmp20451 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20450)

_ = tmp20451

tmp20452 := MakeNative1(func(__e *ControlFlow, V6896 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6892 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6893 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6894 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6895 Obj) Obj {
tmp20453 := GCons(__e, symboolean, Nil)

tmp20454 := GCons(__e, sym_1_1_6, tmp20453)

tmp20455 := GCons(__e, symnumber, tmp20454)

tmp20456 := GCons(__e, tmp20455, Nil)

tmp20457 := GCons(__e, sym_1_1_6, tmp20456)

tmp20458 := GCons(__e, symnumber, tmp20457)

return TailCall6(__e, SymFn(symis_b), V6896, tmp20458, B6892, L6893, Key6894, C6895)


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

tmp20459 := GValue(__e, symshen_4_dsigf_d)

tmp20460 := Call3(__e, SymFn(symshen_4assoc_1_6), sym_6, tmp20452, tmp20459)


tmp20461 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20460)

_ = tmp20461

tmp20462 := MakeNative1(func(__e *ControlFlow, V6901 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6897 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6898 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6899 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6900 Obj) Obj {
tmp20463 := GCons(__e, symboolean, Nil)

tmp20464 := GCons(__e, sym_1_1_6, tmp20463)

tmp20465 := GCons(__e, symnumber, tmp20464)

tmp20466 := GCons(__e, tmp20465, Nil)

tmp20467 := GCons(__e, sym_1_1_6, tmp20466)

tmp20468 := GCons(__e, symnumber, tmp20467)

return TailCall6(__e, SymFn(symis_b), V6901, tmp20468, B6897, L6898, Key6899, C6900)


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

tmp20469 := GValue(__e, symshen_4_dsigf_d)

tmp20470 := Call3(__e, SymFn(symshen_4assoc_1_6), sym_5, tmp20462, tmp20469)


tmp20471 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20470)

_ = tmp20471

tmp20472 := MakeNative1(func(__e *ControlFlow, V6906 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6902 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6903 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6904 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6905 Obj) Obj {
tmp20473 := GCons(__e, symboolean, Nil)

tmp20474 := GCons(__e, sym_1_1_6, tmp20473)

tmp20475 := GCons(__e, symnumber, tmp20474)

tmp20476 := GCons(__e, tmp20475, Nil)

tmp20477 := GCons(__e, sym_1_1_6, tmp20476)

tmp20478 := GCons(__e, symnumber, tmp20477)

return TailCall6(__e, SymFn(symis_b), V6906, tmp20478, B6902, L6903, Key6904, C6905)


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

tmp20480 := Call3(__e, SymFn(symshen_4assoc_1_6), sym_6_a, tmp20472, tmp20479)


tmp20481 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20480)

_ = tmp20481

tmp20482 := MakeNative1(func(__e *ControlFlow, V6911 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6907 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6908 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6909 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6910 Obj) Obj {
tmp20483 := GCons(__e, symboolean, Nil)

tmp20484 := GCons(__e, sym_1_1_6, tmp20483)

tmp20485 := GCons(__e, symnumber, tmp20484)

tmp20486 := GCons(__e, tmp20485, Nil)

tmp20487 := GCons(__e, sym_1_1_6, tmp20486)

tmp20488 := GCons(__e, symnumber, tmp20487)

return TailCall6(__e, SymFn(symis_b), V6911, tmp20488, B6907, L6908, Key6909, C6910)


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

tmp20489 := GValue(__e, symshen_4_dsigf_d)

tmp20490 := Call3(__e, SymFn(symshen_4assoc_1_6), sym_5_a, tmp20482, tmp20489)


tmp20491 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20490)

_ = tmp20491

tmp20492 := MakeNative1(func(__e *ControlFlow, V6916 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6912 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6913 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6914 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6915 Obj) Obj {
tmp20493 := Call1(__e, SymFn(symshen_4newpv), B6912)


let__18070 := tmp20493
_ = let__18070

tmp20494 := GCons(__e, symboolean, Nil)

tmp20495 := GCons(__e, sym_1_1_6, tmp20494)

tmp20496 := GCons(__e, let__18070, tmp20495)

tmp20497 := GCons(__e, tmp20496, Nil)

tmp20498 := GCons(__e, sym_1_1_6, tmp20497)

tmp20499 := GCons(__e, let__18070, tmp20498)

tmp20500 := Call6(__e, SymFn(symis_b), V6916, tmp20499, B6912, L6913, Key6914, C6915)


return TailCall2(__e, SymFn(symshen_4gc), B6912, tmp20500)


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

tmp20501 := GValue(__e, symshen_4_dsigf_d)

tmp20502 := Call3(__e, SymFn(symshen_4assoc_1_6), sym_a, tmp20492, tmp20501)


tmp20503 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20502)

_ = tmp20503

tmp20504 := MakeNative1(func(__e *ControlFlow, V6921 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6917 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6918 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6919 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6920 Obj) Obj {
tmp20505 := GCons(__e, symnumber, Nil)

tmp20506 := GCons(__e, sym_1_1_6, tmp20505)

tmp20507 := GCons(__e, symnumber, tmp20506)

tmp20508 := GCons(__e, tmp20507, Nil)

tmp20509 := GCons(__e, sym_1_1_6, tmp20508)

tmp20510 := GCons(__e, symnumber, tmp20509)

return TailCall6(__e, SymFn(symis_b), V6921, tmp20510, B6917, L6918, Key6919, C6920)


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

tmp20511 := GValue(__e, symshen_4_dsigf_d)

tmp20512 := Call3(__e, SymFn(symshen_4assoc_1_6), sym_7, tmp20504, tmp20511)


tmp20513 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20512)

_ = tmp20513

tmp20514 := MakeNative1(func(__e *ControlFlow, V6926 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6922 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6923 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6924 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6925 Obj) Obj {
tmp20515 := GCons(__e, symnumber, Nil)

tmp20516 := GCons(__e, sym_1_1_6, tmp20515)

tmp20517 := GCons(__e, symnumber, tmp20516)

tmp20518 := GCons(__e, tmp20517, Nil)

tmp20519 := GCons(__e, sym_1_1_6, tmp20518)

tmp20520 := GCons(__e, symnumber, tmp20519)

return TailCall6(__e, SymFn(symis_b), V6926, tmp20520, B6922, L6923, Key6924, C6925)


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

tmp20522 := Call3(__e, SymFn(symshen_4assoc_1_6), sym_c, tmp20514, tmp20521)


tmp20523 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20522)

_ = tmp20523

tmp20524 := MakeNative1(func(__e *ControlFlow, V6931 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6927 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6928 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6929 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6930 Obj) Obj {
tmp20525 := GCons(__e, symnumber, Nil)

tmp20526 := GCons(__e, sym_1_1_6, tmp20525)

tmp20527 := GCons(__e, symnumber, tmp20526)

tmp20528 := GCons(__e, tmp20527, Nil)

tmp20529 := GCons(__e, sym_1_1_6, tmp20528)

tmp20530 := GCons(__e, symnumber, tmp20529)

return TailCall6(__e, SymFn(symis_b), V6931, tmp20530, B6927, L6928, Key6929, C6930)


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

tmp20531 := GValue(__e, symshen_4_dsigf_d)

tmp20532 := Call3(__e, SymFn(symshen_4assoc_1_6), sym_1, tmp20524, tmp20531)


tmp20533 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20532)

_ = tmp20533

tmp20534 := MakeNative1(func(__e *ControlFlow, V6936 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6932 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6933 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6934 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6935 Obj) Obj {
tmp20535 := GCons(__e, symnumber, Nil)

tmp20536 := GCons(__e, sym_1_1_6, tmp20535)

tmp20537 := GCons(__e, symnumber, tmp20536)

tmp20538 := GCons(__e, tmp20537, Nil)

tmp20539 := GCons(__e, sym_1_1_6, tmp20538)

tmp20540 := GCons(__e, symnumber, tmp20539)

return TailCall6(__e, SymFn(symis_b), V6936, tmp20540, B6932, L6933, Key6934, C6935)


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

tmp20541 := GValue(__e, symshen_4_dsigf_d)

tmp20542 := Call3(__e, SymFn(symshen_4assoc_1_6), sym_d, tmp20534, tmp20541)


tmp20543 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20542)

_ = tmp20543

tmp20544 := MakeNative1(func(__e *ControlFlow, V6941 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6937 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6938 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6939 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6940 Obj) Obj {
tmp20545 := Call1(__e, SymFn(symshen_4newpv), B6937)


let__18071 := tmp20545
_ = let__18071

tmp20546 := Call1(__e, SymFn(symshen_4newpv), B6937)


let__18072 := tmp20546
_ = let__18072

tmp20547 := GCons(__e, symboolean, Nil)

tmp20548 := GCons(__e, sym_1_1_6, tmp20547)

tmp20549 := GCons(__e, let__18072, tmp20548)

tmp20550 := GCons(__e, tmp20549, Nil)

tmp20551 := GCons(__e, sym_1_1_6, tmp20550)

tmp20552 := GCons(__e, let__18071, tmp20551)

tmp20553 := Call6(__e, SymFn(symis_b), V6941, tmp20552, B6937, L6938, Key6939, C6940)


tmp20554 := Call2(__e, SymFn(symshen_4gc), B6937, tmp20553)


return TailCall2(__e, SymFn(symshen_4gc), B6937, tmp20554)


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

tmp20555 := GValue(__e, symshen_4_dsigf_d)

tmp20556 := Call3(__e, SymFn(symshen_4assoc_1_6), sym_a_a, tmp20544, tmp20555)


return PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20556)


panic("unreachable")
})

tmp20557 := Call2(__e, ns2_1set, symshen_4initialise_1signedfuncs, tmp18838)


_ = tmp20557

tmp20558 := MakeNative0(func(__e *ControlFlow) Obj {
tmp20559 := MakeNative1(func(__e *ControlFlow, Y954 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4tuple), Y954)
panic("unreachable")
})

tmp20560 := GCons(__e, symshen_4tuple, tmp20559)

tmp20561 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20560)


_ = tmp20561

tmp20562 := MakeNative1(func(__e *ControlFlow, Y955 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4pvar), Y955)
panic("unreachable")
})

tmp20563 := GCons(__e, symshen_4pvar, tmp20562)

tmp20564 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20563)


_ = tmp20564

tmp20565 := MakeNative1(func(__e *ControlFlow, Y956 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4dictionary), Y956)
panic("unreachable")
})

tmp20566 := GCons(__e, symshen_4dictionary, tmp20565)

tmp20567 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20566)


_ = tmp20567

tmp20568 := MakeNative1(func(__e *ControlFlow, Y957 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4print_1prolog_1vector), Y957)
panic("unreachable")
})

tmp20569 := GCons(__e, symshen_4print_1prolog_1vector, tmp20568)

tmp20570 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20569)


_ = tmp20570

tmp20571 := MakeNative1(func(__e *ControlFlow, Y958 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4print_1freshterm), Y958)
panic("unreachable")
})

tmp20572 := GCons(__e, symshen_4print_1freshterm, tmp20571)

tmp20573 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20572)


_ = tmp20573

tmp20574 := MakeNative1(func(__e *ControlFlow, Y959 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4printF), Y959)
panic("unreachable")
})

tmp20575 := GCons(__e, symshen_4printF, tmp20574)

tmp20576 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20575)


_ = tmp20576

tmp20577 := MakeNative1(func(__e *ControlFlow, Y960 Obj) Obj {
return TailCall1(__e, SymFn(symabsolute), Y960)
panic("unreachable")
})

tmp20578 := GCons(__e, symabsolute, tmp20577)

tmp20579 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20578)


_ = tmp20579

tmp20580 := MakeNative1(func(__e *ControlFlow, Y961 Obj) Obj {
return GVectorP(__e, Y961)
panic("unreachable")
})

tmp20581 := GCons(__e, symabsvector_2, tmp20580)

tmp20582 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20581)


_ = tmp20582

tmp20583 := MakeNative1(func(__e *ControlFlow, Y962 Obj) Obj {
return PrimGuard1(__e, symabsvector, PrimAbsvector, Y962)
panic("unreachable")
})

tmp20584 := GCons(__e, symabsvector, tmp20583)

tmp20585 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20584)


_ = tmp20585

tmp20586 := MakeNative1(func(__e *ControlFlow, Y963 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y964 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y965 Obj) Obj {
return GVecSet(__e, Y963, Y964, Y965)
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20587 := GCons(__e, symaddress_1_6, tmp20586)

tmp20588 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20587)


_ = tmp20588

tmp20589 := MakeNative1(func(__e *ControlFlow, Y966 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y967 Obj) Obj {
return TailCall2(__e, SymFn(symadjoin), Y966, Y967)
panic("unreachable")
})
panic("unreachable")
})

tmp20590 := GCons(__e, symadjoin, tmp20589)

tmp20591 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20590)


_ = tmp20591

tmp20592 := MakeNative1(func(__e *ControlFlow, Y968 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y969 Obj) Obj {
if True == Y968 {
return Y969
} else {
return False
}
panic("unreachable")
})
panic("unreachable")
})

tmp20594 := GCons(__e, symand, tmp20592)

tmp20595 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20594)


_ = tmp20595

tmp20596 := MakeNative1(func(__e *ControlFlow, Y970 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y971 Obj) Obj {
return TailCall2(__e, SymFn(symappend), Y970, Y971)
panic("unreachable")
})
panic("unreachable")
})

tmp20597 := GCons(__e, symappend, tmp20596)

tmp20598 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20597)


_ = tmp20598

tmp20599 := MakeNative1(func(__e *ControlFlow, Y972 Obj) Obj {
return TailCall1(__e, SymFn(symarity), Y972)
panic("unreachable")
})

tmp20600 := GCons(__e, symarity, tmp20599)

tmp20601 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20600)


_ = tmp20601

tmp20602 := MakeNative1(func(__e *ControlFlow, Y973 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y974 Obj) Obj {
return TailCall2(__e, SymFn(symassoc), Y973, Y974)
panic("unreachable")
})
panic("unreachable")
})

tmp20603 := GCons(__e, symassoc, tmp20602)

tmp20604 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20603)


_ = tmp20604

tmp20605 := MakeNative1(func(__e *ControlFlow, Y975 Obj) Obj {
return TailCall1(__e, SymFn(symatom_2), Y975)
panic("unreachable")
})

tmp20606 := GCons(__e, symatom_2, tmp20605)

tmp20607 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20606)


_ = tmp20607

tmp20608 := MakeNative1(func(__e *ControlFlow, Y976 Obj) Obj {
return TailCall1(__e, SymFn(symboolean_2), Y976)
panic("unreachable")
})

tmp20609 := GCons(__e, symboolean_2, tmp20608)

tmp20610 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20609)


_ = tmp20610

tmp20611 := MakeNative1(func(__e *ControlFlow, Y977 Obj) Obj {
return TailCall1(__e, SymFn(symbootstrap), Y977)
panic("unreachable")
})

tmp20612 := GCons(__e, symbootstrap, tmp20611)

tmp20613 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20612)


_ = tmp20613

tmp20614 := MakeNative1(func(__e *ControlFlow, Y978 Obj) Obj {
return TailCall1(__e, SymFn(symbound_2), Y978)
panic("unreachable")
})

tmp20615 := GCons(__e, symbound_2, tmp20614)

tmp20616 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20615)


_ = tmp20616

tmp20617 := MakeNative1(func(__e *ControlFlow, Y979 Obj) Obj {
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

tmp20618 := GCons(__e, symbind, tmp20617)

tmp20619 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20618)


_ = tmp20619

tmp20620 := MakeNative1(func(__e *ControlFlow, Y985 Obj) Obj {
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

tmp20621 := GCons(__e, symcall, tmp20620)

tmp20622 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20621)


_ = tmp20622

tmp20623 := MakeNative1(func(__e *ControlFlow, Y990 Obj) Obj {
return TailCall1(__e, SymFn(symcd), Y990)
panic("unreachable")
})

tmp20624 := GCons(__e, symcd, tmp20623)

tmp20625 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20624)


_ = tmp20625

tmp20626 := MakeNative1(func(__e *ControlFlow, Y991 Obj) Obj {
return TailCall1(__e, SymFn(symchar_1stream_2), Y991)
panic("unreachable")
})

tmp20627 := GCons(__e, symchar_1stream_2, tmp20626)

tmp20628 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20627)


_ = tmp20628

tmp20629 := MakeNative1(func(__e *ControlFlow, Y992 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y993 Obj) Obj {
return TailCall2(__e, SymFn(symcompile), Y992, Y993)
panic("unreachable")
})
panic("unreachable")
})

tmp20630 := GCons(__e, symcompile, tmp20629)

tmp20631 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20630)


_ = tmp20631

tmp20632 := MakeNative1(func(__e *ControlFlow, Y994 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y995 Obj) Obj {
return TailCall2(__e, SymFn(symconcat), Y994, Y995)
panic("unreachable")
})
panic("unreachable")
})

tmp20633 := GCons(__e, symconcat, tmp20632)

tmp20634 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20633)


_ = tmp20634

tmp20635 := MakeNative1(func(__e *ControlFlow, Y996 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y997 Obj) Obj {
return GCons(__e, Y996, Y997)
panic("unreachable")
})
panic("unreachable")
})

tmp20636 := GCons(__e, symcons, tmp20635)

tmp20637 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20636)


_ = tmp20637

tmp20638 := MakeNative1(func(__e *ControlFlow, Y998 Obj) Obj {
return GConsP(__e, Y998)
panic("unreachable")
})

tmp20639 := GCons(__e, symcons_2, tmp20638)

tmp20640 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20639)


_ = tmp20640

tmp20641 := MakeNative1(func(__e *ControlFlow, Y999 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1000 Obj) Obj {
return PrimGuard2(__e, symcn, PrimStringConcat, Y999, Y1000)
panic("unreachable")
})
panic("unreachable")
})

tmp20642 := GCons(__e, symcn, tmp20641)

tmp20643 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20642)


_ = tmp20643

tmp20644 := MakeNative1(func(__e *ControlFlow, Y1001 Obj) Obj {
return PrimGuard1(__e, symclose, PrimCloseStream, Y1001)
panic("unreachable")
})

tmp20645 := GCons(__e, symclose, tmp20644)

tmp20646 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20645)


_ = tmp20646

tmp20647 := MakeNative1(func(__e *ControlFlow, Y1002 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1003 Obj) Obj {
return TailCall2(__e, SymFn(symdeclare), Y1002, Y1003)
panic("unreachable")
})
panic("unreachable")
})

tmp20648 := GCons(__e, symdeclare, tmp20647)

tmp20649 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20648)


_ = tmp20649

tmp20650 := MakeNative1(func(__e *ControlFlow, Y1004 Obj) Obj {
return TailCall1(__e, SymFn(symdestroy), Y1004)
panic("unreachable")
})

tmp20651 := GCons(__e, symdestroy, tmp20650)

tmp20652 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20651)


_ = tmp20652

tmp20653 := MakeNative1(func(__e *ControlFlow, Y1005 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1006 Obj) Obj {
return TailCall2(__e, SymFn(symdifference), Y1005, Y1006)
panic("unreachable")
})
panic("unreachable")
})

tmp20654 := GCons(__e, symdifference, tmp20653)

tmp20655 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20654)


_ = tmp20655

tmp20656 := MakeNative1(func(__e *ControlFlow, Y1007 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1008 Obj) Obj {
_ = Y1007

return Y1008


panic("unreachable")
})
panic("unreachable")
})

tmp20657 := GCons(__e, symdo, tmp20656)

tmp20658 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20657)


_ = tmp20658

tmp20659 := MakeNative1(func(__e *ControlFlow, Y1009 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1010 Obj) Obj {
return TailCall2(__e, SymFn(symelement_2), Y1009, Y1010)
panic("unreachable")
})
panic("unreachable")
})

tmp20660 := GCons(__e, symelement_2, tmp20659)

tmp20661 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20660)


_ = tmp20661

tmp20662 := MakeNative1(func(__e *ControlFlow, Y1011 Obj) Obj {
return TailCall1(__e, SymFn(symempty_2), Y1011)
panic("unreachable")
})

tmp20663 := GCons(__e, symempty_2, tmp20662)

tmp20664 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20663)


_ = tmp20664

tmp20665 := MakeNative1(func(__e *ControlFlow, Y1012 Obj) Obj {
return TailCall1(__e, SymFn(symenable_1type_1theory), Y1012)
panic("unreachable")
})

tmp20666 := GCons(__e, symenable_1type_1theory, tmp20665)

tmp20667 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20666)


_ = tmp20667

tmp20668 := MakeNative1(func(__e *ControlFlow, Y1013 Obj) Obj {
return TailCall1(__e, SymFn(symexternal), Y1013)
panic("unreachable")
})

tmp20669 := GCons(__e, symexternal, tmp20668)

tmp20670 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20669)


_ = tmp20670

tmp20671 := MakeNative1(func(__e *ControlFlow, Y1014 Obj) Obj {
return PrimGuard1(__e, symerror_1to_1string, PrimErrorToString, Y1014)
panic("unreachable")
})

tmp20672 := GCons(__e, symerror_1to_1string, tmp20671)

tmp20673 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20672)


_ = tmp20673

tmp20674 := MakeNative1(func(__e *ControlFlow, Y1015 Obj) Obj {
return TailCall1(__e, SymFn(symeval), Y1015)
panic("unreachable")
})

tmp20675 := GCons(__e, symeval, tmp20674)

tmp20676 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20675)


_ = tmp20676

tmp20677 := MakeNative1(func(__e *ControlFlow, Y1016 Obj) Obj {
return TailCall1(__e, SymFn(symeval_1kl), Y1016)
panic("unreachable")
})

tmp20678 := GCons(__e, symeval_1kl, tmp20677)

tmp20679 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20678)


_ = tmp20679

tmp20680 := MakeNative1(func(__e *ControlFlow, Y1017 Obj) Obj {
return TailCall1(__e, SymFn(symexplode), Y1017)
panic("unreachable")
})

tmp20681 := GCons(__e, symexplode, tmp20680)

tmp20682 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20681)


_ = tmp20682

tmp20683 := MakeNative1(func(__e *ControlFlow, Y1018 Obj) Obj {
return TailCall1(__e, SymFn(symexternal), Y1018)
panic("unreachable")
})

tmp20684 := GCons(__e, symexternal, tmp20683)

tmp20685 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20684)


_ = tmp20685

tmp20686 := MakeNative1(func(__e *ControlFlow, Y1019 Obj) Obj {
return TailCall1(__e, SymFn(symfactorise), Y1019)
panic("unreachable")
})

tmp20687 := GCons(__e, symfactorise, tmp20686)

tmp20688 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20687)


_ = tmp20688

tmp20689 := MakeNative1(func(__e *ControlFlow, Y1020 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1021 Obj) Obj {
return TailCall2(__e, SymFn(symfail_1if), Y1020, Y1021)
panic("unreachable")
})
panic("unreachable")
})

tmp20690 := GCons(__e, symfail_1if, tmp20689)

tmp20691 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20690)


_ = tmp20691

tmp20692 := MakeNative1(func(__e *ControlFlow, Y1022 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1023 Obj) Obj {
return TailCall2(__e, SymFn(symfix), Y1022, Y1023)
panic("unreachable")
})
panic("unreachable")
})

tmp20693 := GCons(__e, symfix, tmp20692)

tmp20694 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20693)


_ = tmp20694

tmp20695 := MakeNative1(func(__e *ControlFlow, Y1024 Obj) Obj {
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

tmp20696 := GCons(__e, symfindall, tmp20695)

tmp20697 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20696)


_ = tmp20697

tmp20698 := MakeNative1(func(__e *ControlFlow, Y1031 Obj) Obj {
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

tmp20699 := GCons(__e, symfork, tmp20698)

tmp20700 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20699)


_ = tmp20700

tmp20701 := MakeNative1(func(__e *ControlFlow, Y1036 Obj) Obj {
return MakeNative0(func(__e *ControlFlow) Obj {
return Y1036
panic("unreachable")
})
panic("unreachable")
})

tmp20702 := GCons(__e, symfreeze, tmp20701)

tmp20703 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20702)


_ = tmp20703

tmp20704 := MakeNative1(func(__e *ControlFlow, Y1037 Obj) Obj {
return TailCall1(__e, SymFn(symfst), Y1037)
panic("unreachable")
})

tmp20705 := GCons(__e, symfst, tmp20704)

tmp20706 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20705)


_ = tmp20706

tmp20707 := MakeNative1(func(__e *ControlFlow, Y1038 Obj) Obj {
return TailCall1(__e, SymFn(symfn), Y1038)
panic("unreachable")
})

tmp20708 := GCons(__e, symfn, tmp20707)

tmp20709 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20708)


_ = tmp20709

tmp20710 := MakeNative1(func(__e *ControlFlow, Y1039 Obj) Obj {
return TailCall1(__e, SymFn(symfunction), Y1039)
panic("unreachable")
})

tmp20711 := GCons(__e, symfunction, tmp20710)

tmp20712 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20711)


_ = tmp20712

tmp20713 := MakeNative1(func(__e *ControlFlow, Y1040 Obj) Obj {
return TailCall1(__e, SymFn(symgensym), Y1040)
panic("unreachable")
})

tmp20714 := GCons(__e, symgensym, tmp20713)

tmp20715 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20714)


_ = tmp20715

tmp20716 := MakeNative1(func(__e *ControlFlow, Y1041 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1042 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1043 Obj) Obj {
return TailCall3(__e, SymFn(symget), Y1041, Y1042, Y1043)
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20717 := GCons(__e, symget, tmp20716)

tmp20718 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20717)


_ = tmp20718

tmp20719 := MakeNative1(func(__e *ControlFlow, Y1044 Obj) Obj {
return PrimGuard1(__e, symget_1time, PrimGetTime, Y1044)
panic("unreachable")
})

tmp20720 := GCons(__e, symget_1time, tmp20719)

tmp20721 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20720)


_ = tmp20721

tmp20722 := MakeNative1(func(__e *ControlFlow, Y1045 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1046 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1047 Obj) Obj {
return GVecSet(__e, Y1045, Y1046, Y1047)
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20723 := GCons(__e, symaddress_1_6, tmp20722)

tmp20724 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20723)


_ = tmp20724

tmp20725 := MakeNative1(func(__e *ControlFlow, Y1048 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1049 Obj) Obj {
return GVecGet(__e, Y1048, Y1049)
panic("unreachable")
})
panic("unreachable")
})

tmp20726 := GCons(__e, sym_5_1address, tmp20725)

tmp20727 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20726)


_ = tmp20727

tmp20728 := MakeNative1(func(__e *ControlFlow, Y1050 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1051 Obj) Obj {
return TailCall2(__e, SymFn(sym_5_1vector), Y1050, Y1051)
panic("unreachable")
})
panic("unreachable")
})

tmp20729 := GCons(__e, sym_5_1vector, tmp20728)

tmp20730 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20729)


_ = tmp20730

tmp20731 := MakeNative1(func(__e *ControlFlow, Y1052 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1053 Obj) Obj {
return GGt(__e, Y1052, Y1053)
panic("unreachable")
})
panic("unreachable")
})

tmp20732 := GCons(__e, sym_6, tmp20731)

tmp20733 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20732)


_ = tmp20733

tmp20734 := MakeNative1(func(__e *ControlFlow, Y1054 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1055 Obj) Obj {
return GGe(__e, Y1054, Y1055)
panic("unreachable")
})
panic("unreachable")
})

tmp20735 := GCons(__e, sym_6_a, tmp20734)

tmp20736 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20735)


_ = tmp20736

tmp20737 := MakeNative1(func(__e *ControlFlow, Y1056 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1057 Obj) Obj {
return GEq(__e, Y1056, Y1057)
panic("unreachable")
})
panic("unreachable")
})

tmp20738 := GCons(__e, sym_a, tmp20737)

tmp20739 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20738)


_ = tmp20739

tmp20740 := MakeNative1(func(__e *ControlFlow, Y1058 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1059 Obj) Obj {
return TailCall2(__e, SymFn(symhash), Y1058, Y1059)
panic("unreachable")
})
panic("unreachable")
})

tmp20741 := GCons(__e, symhash, tmp20740)

tmp20742 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20741)


_ = tmp20742

tmp20743 := MakeNative1(func(__e *ControlFlow, Y1060 Obj) Obj {
return GHd(__e, Y1060)
panic("unreachable")
})

tmp20744 := GCons(__e, symhd, tmp20743)

tmp20745 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20744)


_ = tmp20745

tmp20746 := MakeNative1(func(__e *ControlFlow, Y1061 Obj) Obj {
return TailCall1(__e, SymFn(symhdv), Y1061)
panic("unreachable")
})

tmp20747 := GCons(__e, symhdv, tmp20746)

tmp20748 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20747)


_ = tmp20748

tmp20749 := MakeNative1(func(__e *ControlFlow, Y1062 Obj) Obj {
return TailCall1(__e, SymFn(symhdstr), Y1062)
panic("unreachable")
})

tmp20750 := GCons(__e, symhdstr, tmp20749)

tmp20751 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20750)


_ = tmp20751

tmp20752 := MakeNative1(func(__e *ControlFlow, Y1063 Obj) Obj {
return TailCall1(__e, SymFn(symhead), Y1063)
panic("unreachable")
})

tmp20753 := GCons(__e, symhead, tmp20752)

tmp20754 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20753)


_ = tmp20754

tmp20755 := MakeNative1(func(__e *ControlFlow, Y1064 Obj) Obj {
return TailCall1(__e, SymFn(symhush), Y1064)
panic("unreachable")
})

tmp20756 := GCons(__e, symhush, tmp20755)

tmp20757 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20756)


_ = tmp20757

tmp20758 := MakeNative1(func(__e *ControlFlow, Y1065 Obj) Obj {
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

tmp20760 := GCons(__e, symif, tmp20758)

tmp20761 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20760)


_ = tmp20761

tmp20762 := MakeNative1(func(__e *ControlFlow, Y1068 Obj) Obj {
return TailCall1(__e, SymFn(syminclude), Y1068)
panic("unreachable")
})

tmp20763 := GCons(__e, syminclude, tmp20762)

tmp20764 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20763)


_ = tmp20764

tmp20765 := MakeNative1(func(__e *ControlFlow, Y1069 Obj) Obj {
return TailCall1(__e, SymFn(symin_1package), Y1069)
panic("unreachable")
})

tmp20766 := GCons(__e, symin_1package, tmp20765)

tmp20767 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20766)


_ = tmp20767

tmp20768 := MakeNative1(func(__e *ControlFlow, Y1070 Obj) Obj {
return PrimGuard1(__e, syminteger_2, PrimIsInteger, Y1070)
panic("unreachable")
})

tmp20769 := GCons(__e, syminteger_2, tmp20768)

tmp20770 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20769)


_ = tmp20770

tmp20771 := MakeNative1(func(__e *ControlFlow, Y1071 Obj) Obj {
return TailCall1(__e, SymFn(syminternal), Y1071)
panic("unreachable")
})

tmp20772 := GCons(__e, syminternal, tmp20771)

tmp20773 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20772)


_ = tmp20773

tmp20774 := MakeNative1(func(__e *ControlFlow, Y1072 Obj) Obj {
return PrimGuard1(__e, symintern, PrimIntern, Y1072)
panic("unreachable")
})

tmp20775 := GCons(__e, symintern, tmp20774)

tmp20776 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20775)


_ = tmp20776

tmp20777 := MakeNative1(func(__e *ControlFlow, Y1073 Obj) Obj {
return TailCall1(__e, SymFn(syminput), Y1073)
panic("unreachable")
})

tmp20778 := GCons(__e, syminput, tmp20777)

tmp20779 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20778)


_ = tmp20779

tmp20780 := MakeNative1(func(__e *ControlFlow, Y1074 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1075 Obj) Obj {
return TailCall2(__e, SymFn(symshen_4input_1h_7), Y1074, Y1075)
panic("unreachable")
})
panic("unreachable")
})

tmp20781 := GCons(__e, syminput_7, tmp20780)

tmp20782 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20781)


_ = tmp20782

tmp20783 := MakeNative1(func(__e *ControlFlow, Y1076 Obj) Obj {
return TailCall1(__e, SymFn(syminclude_1all_1but), Y1076)
panic("unreachable")
})

tmp20784 := GCons(__e, syminclude_1all_1but, tmp20783)

tmp20785 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20784)


_ = tmp20785

tmp20786 := MakeNative1(func(__e *ControlFlow, Y1077 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1078 Obj) Obj {
return TailCall2(__e, SymFn(symintersection), Y1077, Y1078)
panic("unreachable")
})
panic("unreachable")
})

tmp20787 := GCons(__e, symintersection, tmp20786)

tmp20788 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20787)


_ = tmp20788

tmp20789 := MakeNative1(func(__e *ControlFlow, Y1079 Obj) Obj {
return TailCall1(__e, SymFn(syminternal), Y1079)
panic("unreachable")
})

tmp20790 := GCons(__e, syminternal, tmp20789)

tmp20791 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20790)


_ = tmp20791

tmp20792 := MakeNative1(func(__e *ControlFlow, Y1080 Obj) Obj {
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

tmp20793 := GCons(__e, symis, tmp20792)

tmp20794 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20793)


_ = tmp20794

tmp20795 := MakeNative1(func(__e *ControlFlow, Y1086 Obj) Obj {
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

tmp20796 := GCons(__e, symis_b, tmp20795)

tmp20797 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20796)


_ = tmp20797

tmp20798 := MakeNative1(func(__e *ControlFlow, Y1092 Obj) Obj {
return TailCall1(__e, SymFn(symlength), Y1092)
panic("unreachable")
})

tmp20799 := GCons(__e, symlength, tmp20798)

tmp20800 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20799)


_ = tmp20800

tmp20801 := MakeNative1(func(__e *ControlFlow, Y1093 Obj) Obj {
return TailCall1(__e, SymFn(symlimit), Y1093)
panic("unreachable")
})

tmp20802 := GCons(__e, symlimit, tmp20801)

tmp20803 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20802)


_ = tmp20803

tmp20804 := MakeNative1(func(__e *ControlFlow, Y1094 Obj) Obj {
return TailCall1(__e, SymFn(symlineread), Y1094)
panic("unreachable")
})

tmp20805 := GCons(__e, symlineread, tmp20804)

tmp20806 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20805)


_ = tmp20806

tmp20807 := MakeNative1(func(__e *ControlFlow, Y1095 Obj) Obj {
return TailCall1(__e, SymFn(symload), Y1095)
panic("unreachable")
})

tmp20808 := GCons(__e, symload, tmp20807)

tmp20809 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20808)


_ = tmp20809

tmp20810 := MakeNative1(func(__e *ControlFlow, Y1096 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1097 Obj) Obj {
return GLt(__e, Y1096, Y1097)
panic("unreachable")
})
panic("unreachable")
})

tmp20811 := GCons(__e, sym_5, tmp20810)

tmp20812 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20811)


_ = tmp20812

tmp20813 := MakeNative1(func(__e *ControlFlow, Y1098 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1099 Obj) Obj {
return GLe(__e, Y1098, Y1099)
panic("unreachable")
})
panic("unreachable")
})

tmp20814 := GCons(__e, sym_5_a, tmp20813)

tmp20815 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20814)


_ = tmp20815

tmp20816 := MakeNative1(func(__e *ControlFlow, Y1100 Obj) Obj {
return TailCall1(__e, SymFn(symvector), Y1100)
panic("unreachable")
})

tmp20817 := GCons(__e, symvector, tmp20816)

tmp20818 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20817)


_ = tmp20818

tmp20819 := MakeNative1(func(__e *ControlFlow, Y1101 Obj) Obj {
return TailCall1(__e, SymFn(symmacroexpand), Y1101)
panic("unreachable")
})

tmp20820 := GCons(__e, symmacroexpand, tmp20819)

tmp20821 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20820)


_ = tmp20821

tmp20822 := MakeNative1(func(__e *ControlFlow, Y1102 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1103 Obj) Obj {
return TailCall2(__e, SymFn(symmap), Y1102, Y1103)
panic("unreachable")
})
panic("unreachable")
})

tmp20823 := GCons(__e, symmap, tmp20822)

tmp20824 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20823)


_ = tmp20824

tmp20825 := MakeNative1(func(__e *ControlFlow, Y1104 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1105 Obj) Obj {
return TailCall2(__e, SymFn(symmapcan), Y1104, Y1105)
panic("unreachable")
})
panic("unreachable")
})

tmp20826 := GCons(__e, symmapcan, tmp20825)

tmp20827 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20826)


_ = tmp20827

tmp20828 := MakeNative1(func(__e *ControlFlow, Y1106 Obj) Obj {
return TailCall1(__e, SymFn(symmaxinferences), Y1106)
panic("unreachable")
})

tmp20829 := GCons(__e, symmaxinferences, tmp20828)

tmp20830 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20829)


_ = tmp20830

tmp20831 := MakeNative1(func(__e *ControlFlow, Y1107 Obj) Obj {
return TailCall1(__e, SymFn(symnl), Y1107)
panic("unreachable")
})

tmp20832 := GCons(__e, symnl, tmp20831)

tmp20833 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20832)


_ = tmp20833

tmp20834 := MakeNative1(func(__e *ControlFlow, Y1108 Obj) Obj {
return GNot(__e, Y1108)
panic("unreachable")
})

tmp20835 := GCons(__e, symnot, tmp20834)

tmp20836 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20835)


_ = tmp20836

tmp20837 := MakeNative1(func(__e *ControlFlow, Y1109 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1110 Obj) Obj {
return TailCall2(__e, SymFn(symnth), Y1109, Y1110)
panic("unreachable")
})
panic("unreachable")
})

tmp20838 := GCons(__e, symnth, tmp20837)

tmp20839 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20838)


_ = tmp20839

tmp20840 := MakeNative1(func(__e *ControlFlow, Y1111 Obj) Obj {
return PrimGuard1(__e, symn_1_6string, PrimNumberToString, Y1111)
panic("unreachable")
})

tmp20841 := GCons(__e, symn_1_6string, tmp20840)

tmp20842 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20841)


_ = tmp20842

tmp20843 := MakeNative1(func(__e *ControlFlow, Y1112 Obj) Obj {
return PrimGuard1(__e, symnumber_2, PrimIsNumber, Y1112)
panic("unreachable")
})

tmp20844 := GCons(__e, symnumber_2, tmp20843)

tmp20845 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20844)


_ = tmp20845

tmp20846 := MakeNative1(func(__e *ControlFlow, Y1113 Obj) Obj {
return TailCall1(__e, SymFn(symoccurs_1check), Y1113)
panic("unreachable")
})

tmp20847 := GCons(__e, symoccurs_1check, tmp20846)

tmp20848 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20847)


_ = tmp20848

tmp20849 := MakeNative1(func(__e *ControlFlow, Y1114 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1115 Obj) Obj {
return TailCall2(__e, SymFn(symoccurrences), Y1114, Y1115)
panic("unreachable")
})
panic("unreachable")
})

tmp20850 := GCons(__e, symoccurrences, tmp20849)

tmp20851 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20850)


_ = tmp20851

tmp20852 := MakeNative1(func(__e *ControlFlow, Y1116 Obj) Obj {
return TailCall1(__e, SymFn(symoccurs_1check), Y1116)
panic("unreachable")
})

tmp20853 := GCons(__e, symoccurs_1check, tmp20852)

tmp20854 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20853)


_ = tmp20854

tmp20855 := MakeNative1(func(__e *ControlFlow, Y1117 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1118 Obj) Obj {
return PrimGuard2(__e, symopen, PrimOpenStream, Y1117, Y1118)
panic("unreachable")
})
panic("unreachable")
})

tmp20856 := GCons(__e, symopen, tmp20855)

tmp20857 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20856)


_ = tmp20857

tmp20858 := MakeNative1(func(__e *ControlFlow, Y1119 Obj) Obj {
return TailCall1(__e, SymFn(symoptimise), Y1119)
panic("unreachable")
})

tmp20859 := GCons(__e, symoptimise, tmp20858)

tmp20860 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20859)


_ = tmp20860

tmp20861 := MakeNative1(func(__e *ControlFlow, Y1120 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1121 Obj) Obj {
if True == Y1120 {
return True
} else {
return Y1121
}
panic("unreachable")
})
panic("unreachable")
})

tmp20863 := GCons(__e, symor, tmp20861)

tmp20864 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20863)


_ = tmp20864

tmp20865 := MakeNative1(func(__e *ControlFlow, Y1122 Obj) Obj {
return TailCall1(__e, SymFn(sympackage_2), Y1122)
panic("unreachable")
})

tmp20866 := GCons(__e, sympackage_2, tmp20865)

tmp20867 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20866)


_ = tmp20867

tmp20868 := MakeNative1(func(__e *ControlFlow, Y1123 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1124 Obj) Obj {
return PrimGuard2(__e, sympos, PrimPos, Y1123, Y1124)
panic("unreachable")
})
panic("unreachable")
})

tmp20869 := GCons(__e, sympos, tmp20868)

tmp20870 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20869)


_ = tmp20870

tmp20871 := MakeNative1(func(__e *ControlFlow, Y1125 Obj) Obj {
return TailCall1(__e, SymFn(sympreclude_1all_1but), Y1125)
panic("unreachable")
})

tmp20872 := GCons(__e, sympreclude_1all_1but, tmp20871)

tmp20873 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20872)


_ = tmp20873

tmp20874 := MakeNative1(func(__e *ControlFlow, Y1126 Obj) Obj {
return TailCall1(__e, SymFn(symprint), Y1126)
panic("unreachable")
})

tmp20875 := GCons(__e, symprint, tmp20874)

tmp20876 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20875)


_ = tmp20876

tmp20877 := MakeNative1(func(__e *ControlFlow, Y1127 Obj) Obj {
return TailCall1(__e, SymFn(symprofile), Y1127)
panic("unreachable")
})

tmp20878 := GCons(__e, symprofile, tmp20877)

tmp20879 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20878)


_ = tmp20879

tmp20880 := MakeNative1(func(__e *ControlFlow, Y1128 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4print_1prolog_1vector), Y1128)
panic("unreachable")
})

tmp20881 := GCons(__e, symshen_4print_1prolog_1vector, tmp20880)

tmp20882 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20881)


_ = tmp20882

tmp20883 := MakeNative1(func(__e *ControlFlow, Y1129 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4print_1freshterm), Y1129)
panic("unreachable")
})

tmp20884 := GCons(__e, symshen_4print_1freshterm, tmp20883)

tmp20885 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20884)


_ = tmp20885

tmp20886 := MakeNative1(func(__e *ControlFlow, Y1130 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4printF), Y1130)
panic("unreachable")
})

tmp20887 := GCons(__e, symshen_4printF, tmp20886)

tmp20888 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20887)


_ = tmp20888

tmp20889 := MakeNative1(func(__e *ControlFlow, Y1131 Obj) Obj {
return TailCall1(__e, SymFn(symprolog_1memory), Y1131)
panic("unreachable")
})

tmp20890 := GCons(__e, symprolog_1memory, tmp20889)

tmp20891 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20890)


_ = tmp20891

tmp20892 := MakeNative1(func(__e *ControlFlow, Y1132 Obj) Obj {
return TailCall1(__e, SymFn(symprofile_1results), Y1132)
panic("unreachable")
})

tmp20893 := GCons(__e, symprofile_1results, tmp20892)

tmp20894 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20893)


_ = tmp20894

tmp20895 := MakeNative1(func(__e *ControlFlow, Y1133 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1134 Obj) Obj {
return TailCall2(__e, SymFn(sympr), Y1133, Y1134)
panic("unreachable")
})
panic("unreachable")
})

tmp20896 := GCons(__e, sympr, tmp20895)

tmp20897 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20896)


_ = tmp20897

tmp20898 := MakeNative1(func(__e *ControlFlow, Y1135 Obj) Obj {
return TailCall1(__e, SymFn(symps), Y1135)
panic("unreachable")
})

tmp20899 := GCons(__e, symps, tmp20898)

tmp20900 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20899)


_ = tmp20900

tmp20901 := MakeNative1(func(__e *ControlFlow, Y1136 Obj) Obj {
return TailCall1(__e, SymFn(sympreclude), Y1136)
panic("unreachable")
})

tmp20902 := GCons(__e, sympreclude, tmp20901)

tmp20903 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20902)


_ = tmp20903

tmp20904 := MakeNative1(func(__e *ControlFlow, Y1137 Obj) Obj {
return TailCall1(__e, SymFn(sympreclude_1all_1but), Y1137)
panic("unreachable")
})

tmp20905 := GCons(__e, sympreclude_1all_1but, tmp20904)

tmp20906 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20905)


_ = tmp20906

tmp20907 := MakeNative1(func(__e *ControlFlow, Y1138 Obj) Obj {
return TailCall1(__e, SymFn(symprotect), Y1138)
panic("unreachable")
})

tmp20908 := GCons(__e, symprotect, tmp20907)

tmp20909 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20908)


_ = tmp20909

tmp20910 := MakeNative1(func(__e *ControlFlow, Y1139 Obj) Obj {
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

tmp20911 := GCons(__e, symput, tmp20910)

tmp20912 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20911)


_ = tmp20912

tmp20913 := MakeNative1(func(__e *ControlFlow, Y1143 Obj) Obj {
return PrimGuard1(__e, symread_1file_1as_1string, PrimReadFileAsString, Y1143)
panic("unreachable")
})

tmp20914 := GCons(__e, symread_1file_1as_1string, tmp20913)

tmp20915 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20914)


_ = tmp20915

tmp20916 := MakeNative1(func(__e *ControlFlow, Y1144 Obj) Obj {
return PrimGuard1(__e, symread_1file_1as_1bytelist, PrimReadFileAsByteList, Y1144)
panic("unreachable")
})

tmp20917 := GCons(__e, symread_1file_1as_1bytelist, tmp20916)

tmp20918 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20917)


_ = tmp20918

tmp20919 := MakeNative1(func(__e *ControlFlow, Y1145 Obj) Obj {
return TailCall1(__e, SymFn(symread_1file), Y1145)
panic("unreachable")
})

tmp20920 := GCons(__e, symread_1file, tmp20919)

tmp20921 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20920)


_ = tmp20921

tmp20922 := MakeNative1(func(__e *ControlFlow, Y1146 Obj) Obj {
return TailCall1(__e, SymFn(symread), Y1146)
panic("unreachable")
})

tmp20923 := GCons(__e, symread, tmp20922)

tmp20924 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20923)


_ = tmp20924

tmp20925 := MakeNative1(func(__e *ControlFlow, Y1147 Obj) Obj {
return PrimGuard1(__e, symread_1byte, PrimReadByte, Y1147)
panic("unreachable")
})

tmp20926 := GCons(__e, symread_1byte, tmp20925)

tmp20927 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20926)


_ = tmp20927

tmp20928 := MakeNative1(func(__e *ControlFlow, Y1148 Obj) Obj {
return TailCall1(__e, SymFn(symread_1from_1string), Y1148)
panic("unreachable")
})

tmp20929 := GCons(__e, symread_1from_1string, tmp20928)

tmp20930 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20929)


_ = tmp20930

tmp20931 := MakeNative1(func(__e *ControlFlow, Y1149 Obj) Obj {
return TailCall1(__e, SymFn(symread_1from_1string_1unprocessed), Y1149)
panic("unreachable")
})

tmp20932 := GCons(__e, symread_1from_1string_1unprocessed, tmp20931)

tmp20933 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20932)


_ = tmp20933

tmp20934 := MakeNative1(func(__e *ControlFlow, Y1150 Obj) Obj {
return TailCall1(__e, SymFn(symread_1unit_1string), Y1150)
panic("unreachable")
})

tmp20935 := GCons(__e, symread_1unit_1string, tmp20934)

tmp20936 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20935)


_ = tmp20936

tmp20937 := MakeNative1(func(__e *ControlFlow, Y1151 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1152 Obj) Obj {
return TailCall2(__e, SymFn(symshen_4register_1source_1form), Y1151, Y1152)
panic("unreachable")
})
panic("unreachable")
})

tmp20938 := GCons(__e, symshen_4register_1source_1form, tmp20937)

tmp20939 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20938)


_ = tmp20939

tmp20940 := MakeNative1(func(__e *ControlFlow, Y1153 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4unregister_1source_1form), Y1153)
panic("unreachable")
})

tmp20941 := GCons(__e, symshen_4unregister_1source_1form, tmp20940)

tmp20942 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20941)


_ = tmp20942

tmp20943 := MakeNative1(func(__e *ControlFlow, Y1154 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1155 Obj) Obj {
return TailCall2(__e, SymFn(symremove), Y1154, Y1155)
panic("unreachable")
})
panic("unreachable")
})

tmp20944 := GCons(__e, symremove, tmp20943)

tmp20945 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20944)


_ = tmp20945

tmp20946 := MakeNative1(func(__e *ControlFlow, Y1156 Obj) Obj {
return TailCall1(__e, SymFn(symreverse), Y1156)
panic("unreachable")
})

tmp20947 := GCons(__e, symreverse, tmp20946)

tmp20948 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20947)


_ = tmp20948

tmp20949 := MakeNative1(func(__e *ControlFlow, Y1157 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1158 Obj) Obj {
return PrimGuard2(__e, symset, PrimSet, Y1157, Y1158)
panic("unreachable")
})
panic("unreachable")
})

tmp20950 := GCons(__e, symset, tmp20949)

tmp20951 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20950)


_ = tmp20951

tmp20952 := MakeNative1(func(__e *ControlFlow, Y1159 Obj) Obj {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, Y1159)
panic("unreachable")
})

tmp20953 := GCons(__e, symsimple_1error, tmp20952)

tmp20954 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20953)


_ = tmp20954

tmp20955 := MakeNative1(func(__e *ControlFlow, Y1160 Obj) Obj {
return TailCall1(__e, SymFn(symsnd), Y1160)
panic("unreachable")
})

tmp20956 := GCons(__e, symsnd, tmp20955)

tmp20957 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20956)


_ = tmp20957

tmp20958 := MakeNative1(func(__e *ControlFlow, Y1161 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1162 Obj) Obj {
return TailCall2(__e, SymFn(symspecialise), Y1161, Y1162)
panic("unreachable")
})
panic("unreachable")
})

tmp20959 := GCons(__e, symspecialise, tmp20958)

tmp20960 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20959)


_ = tmp20960

tmp20961 := MakeNative1(func(__e *ControlFlow, Y1163 Obj) Obj {
return TailCall1(__e, SymFn(symspy), Y1163)
panic("unreachable")
})

tmp20962 := GCons(__e, symspy, tmp20961)

tmp20963 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20962)


_ = tmp20963

tmp20964 := MakeNative1(func(__e *ControlFlow, Y1164 Obj) Obj {
return TailCall1(__e, SymFn(symstep), Y1164)
panic("unreachable")
})

tmp20965 := GCons(__e, symstep, tmp20964)

tmp20966 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20965)


_ = tmp20966

tmp20967 := MakeNative1(func(__e *ControlFlow, Y1165 Obj) Obj {
return PrimGuard1(__e, symstr, PrimStr, Y1165)
panic("unreachable")
})

tmp20968 := GCons(__e, symstr, tmp20967)

tmp20969 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20968)


_ = tmp20969

tmp20970 := MakeNative1(func(__e *ControlFlow, Y1166 Obj) Obj {
return PrimGuard1(__e, symstring_1_6n, PrimStringToNumber, Y1166)
panic("unreachable")
})

tmp20971 := GCons(__e, symstring_1_6n, tmp20970)

tmp20972 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20971)


_ = tmp20972

tmp20973 := MakeNative1(func(__e *ControlFlow, Y1167 Obj) Obj {
return TailCall1(__e, SymFn(symstring_1_6symbol), Y1167)
panic("unreachable")
})

tmp20974 := GCons(__e, symstring_1_6symbol, tmp20973)

tmp20975 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20974)


_ = tmp20975

tmp20976 := MakeNative1(func(__e *ControlFlow, Y1168 Obj) Obj {
return PrimGuard1(__e, symstring_2, PrimIsString, Y1168)
panic("unreachable")
})

tmp20977 := GCons(__e, symstring_2, tmp20976)

tmp20978 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20977)


_ = tmp20978

tmp20979 := MakeNative1(func(__e *ControlFlow, Y1169 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1170 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1171 Obj) Obj {
return TailCall3(__e, SymFn(symsubst), Y1169, Y1170, Y1171)
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20980 := GCons(__e, symsubst, tmp20979)

tmp20981 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20980)


_ = tmp20981

tmp20982 := MakeNative1(func(__e *ControlFlow, Y1172 Obj) Obj {
return TailCall1(__e, SymFn(symsum), Y1172)
panic("unreachable")
})

tmp20983 := GCons(__e, symsum, tmp20982)

tmp20984 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20983)


_ = tmp20984

tmp20985 := MakeNative1(func(__e *ControlFlow, Y1173 Obj) Obj {
return PrimGuard1(__e, symsymbol_2, PrimIsSymbol, Y1173)
panic("unreachable")
})

tmp20986 := GCons(__e, symsymbol_2, tmp20985)

tmp20987 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20986)


_ = tmp20987

tmp20988 := MakeNative1(func(__e *ControlFlow, Y1174 Obj) Obj {
return TailCall1(__e, SymFn(symsystemf), Y1174)
panic("unreachable")
})

tmp20989 := GCons(__e, symsystemf, tmp20988)

tmp20990 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20989)


_ = tmp20990

tmp20991 := MakeNative1(func(__e *ControlFlow, Y1175 Obj) Obj {
return TailCall1(__e, SymFn(symtail), Y1175)
panic("unreachable")
})

tmp20992 := GCons(__e, symtail, tmp20991)

tmp20993 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20992)


_ = tmp20993

tmp20994 := MakeNative1(func(__e *ControlFlow, Y1176 Obj) Obj {
return GTl(__e, Y1176)
panic("unreachable")
})

tmp20995 := GCons(__e, symtl, tmp20994)

tmp20996 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20995)


_ = tmp20996

tmp20997 := MakeNative1(func(__e *ControlFlow, Y1177 Obj) Obj {
return TailCall1(__e, SymFn(symtc), Y1177)
panic("unreachable")
})

tmp20998 := GCons(__e, symtc, tmp20997)

tmp20999 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20998)


_ = tmp20999

tmp21000 := MakeNative1(func(__e *ControlFlow, Y1178 Obj) Obj {
return TailCall1(__e, SymFn(symthaw), Y1178)
panic("unreachable")
})

tmp21001 := GCons(__e, symthaw, tmp21000)

tmp21002 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21001)


_ = tmp21002

tmp21003 := MakeNative1(func(__e *ControlFlow, Y1179 Obj) Obj {
return PrimGuard1(__e, symtlstr, PrimTailString, Y1179)
panic("unreachable")
})

tmp21004 := GCons(__e, symtlstr, tmp21003)

tmp21005 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21004)


_ = tmp21005

tmp21006 := MakeNative1(func(__e *ControlFlow, Y1180 Obj) Obj {
return TailCall1(__e, SymFn(symtrack), Y1180)
panic("unreachable")
})

tmp21007 := GCons(__e, symtrack, tmp21006)

tmp21008 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21007)


_ = tmp21008

tmp21009 := MakeNative1(func(__e *ControlFlow, Y1181 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1182 Obj) Obj {
tmp21010 := MakeNative0(func(__e *ControlFlow) Obj {
return Y1181
panic("unreachable")
})

return TailCall2(__e, try_1catch, tmp21010, Y1182)


panic("unreachable")
})
panic("unreachable")
})

tmp21011 := GCons(__e, symtrap_1error, tmp21009)

tmp21012 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21011)


_ = tmp21012

tmp21013 := MakeNative1(func(__e *ControlFlow, Y1183 Obj) Obj {
return TailCall1(__e, SymFn(symtuple_2), Y1183)
panic("unreachable")
})

tmp21014 := GCons(__e, symtuple_2, tmp21013)

tmp21015 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21014)


_ = tmp21015

tmp21016 := MakeNative1(func(__e *ControlFlow, Y1184 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1185 Obj) Obj {
return Y1185
panic("unreachable")
})
panic("unreachable")
})

tmp21017 := GCons(__e, symtype, tmp21016)

tmp21018 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21017)


_ = tmp21018

tmp21019 := MakeNative1(func(__e *ControlFlow, Y1186 Obj) Obj {
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

tmp21020 := GCons(__e, symreturn, tmp21019)

tmp21021 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21020)


_ = tmp21021

tmp21022 := MakeNative1(func(__e *ControlFlow, Y1191 Obj) Obj {
return TailCall1(__e, SymFn(symunabsolute), Y1191)
panic("unreachable")
})

tmp21023 := GCons(__e, symunabsolute, tmp21022)

tmp21024 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21023)


_ = tmp21024

tmp21025 := MakeNative1(func(__e *ControlFlow, Y1192 Obj) Obj {
return TailCall1(__e, SymFn(symundefmacro), Y1192)
panic("unreachable")
})

tmp21026 := GCons(__e, symundefmacro, tmp21025)

tmp21027 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21026)


_ = tmp21027

tmp21028 := MakeNative1(func(__e *ControlFlow, Y1193 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1194 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1195 Obj) Obj {
return TailCall3(__e, SymFn(symunput), Y1193, Y1194, Y1195)
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp21029 := GCons(__e, symunput, tmp21028)

tmp21030 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21029)


_ = tmp21030

tmp21031 := MakeNative1(func(__e *ControlFlow, Y1196 Obj) Obj {
return TailCall1(__e, SymFn(symunprofile), Y1196)
panic("unreachable")
})

tmp21032 := GCons(__e, symunprofile, tmp21031)

tmp21033 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21032)


_ = tmp21033

tmp21034 := MakeNative1(func(__e *ControlFlow, Y1197 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1198 Obj) Obj {
return TailCall2(__e, SymFn(symunion), Y1197, Y1198)
panic("unreachable")
})
panic("unreachable")
})

tmp21035 := GCons(__e, symunion, tmp21034)

tmp21036 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21035)


_ = tmp21036

tmp21037 := MakeNative1(func(__e *ControlFlow, Y1199 Obj) Obj {
return TailCall1(__e, SymFn(symuntrack), Y1199)
panic("unreachable")
})

tmp21038 := GCons(__e, symuntrack, tmp21037)

tmp21039 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21038)


_ = tmp21039

tmp21040 := MakeNative1(func(__e *ControlFlow, Y1200 Obj) Obj {
return TailCall1(__e, SymFn(symundefmacro), Y1200)
panic("unreachable")
})

tmp21041 := GCons(__e, symundefmacro, tmp21040)

tmp21042 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21041)


_ = tmp21042

tmp21043 := MakeNative1(func(__e *ControlFlow, Y1201 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1202 Obj) Obj {
return TailCall2(__e, SymFn(symupdate_1lambda_1table), Y1201, Y1202)
panic("unreachable")
})
panic("unreachable")
})

tmp21044 := GCons(__e, symupdate_1lambda_1table, tmp21043)

tmp21045 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21044)


_ = tmp21045

tmp21046 := MakeNative1(func(__e *ControlFlow, Y1203 Obj) Obj {
return TailCall1(__e, SymFn(symvector), Y1203)
panic("unreachable")
})

tmp21047 := GCons(__e, symvector, tmp21046)

tmp21048 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21047)


_ = tmp21048

tmp21049 := MakeNative1(func(__e *ControlFlow, Y1204 Obj) Obj {
return TailCall1(__e, SymFn(symvector_2), Y1204)
panic("unreachable")
})

tmp21050 := GCons(__e, symvector_2, tmp21049)

tmp21051 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21050)


_ = tmp21051

tmp21052 := MakeNative1(func(__e *ControlFlow, Y1205 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1206 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1207 Obj) Obj {
return TailCall3(__e, SymFn(symvector_1_6), Y1205, Y1206, Y1207)
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp21053 := GCons(__e, symvector_1_6, tmp21052)

tmp21054 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21053)


_ = tmp21054

tmp21055 := MakeNative1(func(__e *ControlFlow, Y1208 Obj) Obj {
return GValue(__e, Y1208)
panic("unreachable")
})

tmp21056 := GCons(__e, symvalue, tmp21055)

tmp21057 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21056)


_ = tmp21057

tmp21058 := MakeNative1(func(__e *ControlFlow, Y1209 Obj) Obj {
return PrimGuard1(__e, symvariable_2, PrimIsVariable, Y1209)
panic("unreachable")
})

tmp21059 := GCons(__e, symvariable_2, tmp21058)

tmp21060 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21059)


_ = tmp21060

tmp21061 := MakeNative1(func(__e *ControlFlow, Y1210 Obj) Obj {
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

tmp21062 := GCons(__e, symvar_2, tmp21061)

tmp21063 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21062)


_ = tmp21063

tmp21064 := MakeNative1(func(__e *ControlFlow, Y1215 Obj) Obj {
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

tmp21065 := GCons(__e, symwhen, tmp21064)

tmp21066 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21065)


_ = tmp21066

tmp21067 := MakeNative1(func(__e *ControlFlow, Y1220 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1221 Obj) Obj {
return PrimGuard2(__e, symwrite_1byte, PrimWriteByte, Y1220, Y1221)
panic("unreachable")
})
panic("unreachable")
})

tmp21068 := GCons(__e, symwrite_1byte, tmp21067)

tmp21069 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21068)


_ = tmp21069

tmp21070 := MakeNative1(func(__e *ControlFlow, Y1222 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1223 Obj) Obj {
return TailCall2(__e, SymFn(symwrite_1string), Y1222, Y1223)
panic("unreachable")
})
panic("unreachable")
})

tmp21071 := GCons(__e, symwrite_1string, tmp21070)

tmp21072 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21071)


_ = tmp21072

tmp21073 := MakeNative1(func(__e *ControlFlow, Y1224 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1225 Obj) Obj {
return TailCall2(__e, SymFn(symwrite_1to_1file), Y1224, Y1225)
panic("unreachable")
})
panic("unreachable")
})

tmp21074 := GCons(__e, symwrite_1to_1file, tmp21073)

tmp21075 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21074)


_ = tmp21075

tmp21076 := MakeNative1(func(__e *ControlFlow, Y1226 Obj) Obj {
return TailCall1(__e, SymFn(symy_1or_1n_2), Y1226)
panic("unreachable")
})

tmp21077 := GCons(__e, symy_1or_1n_2, tmp21076)

tmp21078 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21077)


_ = tmp21078

tmp21079 := MakeNative1(func(__e *ControlFlow, Y1227 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1228 Obj) Obj {
return GAdd(__e, Y1227, Y1228)
panic("unreachable")
})
panic("unreachable")
})

tmp21080 := GCons(__e, sym_7, tmp21079)

tmp21081 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21080)


_ = tmp21081

tmp21082 := MakeNative1(func(__e *ControlFlow, Y1229 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1230 Obj) Obj {
return GMul(__e, Y1229, Y1230)
panic("unreachable")
})
panic("unreachable")
})

tmp21083 := GCons(__e, sym_d, tmp21082)

tmp21084 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21083)


_ = tmp21084

tmp21085 := MakeNative1(func(__e *ControlFlow, Y1231 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1232 Obj) Obj {
return GDiv(__e, Y1231, Y1232)
panic("unreachable")
})
panic("unreachable")
})

tmp21086 := GCons(__e, sym_c, tmp21085)

tmp21087 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21086)


_ = tmp21087

tmp21088 := MakeNative1(func(__e *ControlFlow, Y1233 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1234 Obj) Obj {
return GSub(__e, Y1233, Y1234)
panic("unreachable")
})
panic("unreachable")
})

tmp21089 := GCons(__e, sym_1, tmp21088)

tmp21090 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21089)


_ = tmp21090

tmp21091 := MakeNative1(func(__e *ControlFlow, Y1235 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1236 Obj) Obj {
return TailCall2(__e, SymFn(sym_a_a), Y1235, Y1236)
panic("unreachable")
})
panic("unreachable")
})

tmp21092 := GCons(__e, sym_a_a, tmp21091)

tmp21093 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21092)


_ = tmp21093

tmp21094 := MakeNative1(func(__e *ControlFlow, Y1237 Obj) Obj {
return TailCall1(__e, SymFn(sym_5e_6), Y1237)
panic("unreachable")
})

tmp21095 := GCons(__e, sym_5e_6, tmp21094)

tmp21096 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21095)


_ = tmp21096

tmp21097 := MakeNative1(func(__e *ControlFlow, Y1238 Obj) Obj {
return TailCall1(__e, SymFn(sym_5end_6), Y1238)
panic("unreachable")
})

tmp21098 := GCons(__e, sym_5end_6, tmp21097)

tmp21099 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21098)


_ = tmp21099

tmp21100 := MakeNative1(func(__e *ControlFlow, Y1239 Obj) Obj {
return TailCall1(__e, SymFn(sym_5_b_6), Y1239)
panic("unreachable")
})

tmp21101 := GCons(__e, sym_5_b_6, tmp21100)

tmp21102 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21101)


_ = tmp21102

tmp21103 := MakeNative1(func(__e *ControlFlow, Y1240 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1241 Obj) Obj {
return TailCall2(__e, SymFn(sym_8p), Y1240, Y1241)
panic("unreachable")
})
panic("unreachable")
})

tmp21104 := GCons(__e, sym_8p, tmp21103)

tmp21105 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21104)


_ = tmp21105

tmp21106 := MakeNative1(func(__e *ControlFlow, Y1242 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1243 Obj) Obj {
return TailCall2(__e, SymFn(sym_8v), Y1242, Y1243)
panic("unreachable")
})
panic("unreachable")
})

tmp21107 := GCons(__e, sym_8v, tmp21106)

tmp21108 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21107)


_ = tmp21108

tmp21109 := MakeNative1(func(__e *ControlFlow, Y1244 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1245 Obj) Obj {
return TailCall2(__e, SymFn(sym_8s), Y1244, Y1245)
panic("unreachable")
})
panic("unreachable")
})

tmp21110 := GCons(__e, sym_8s, tmp21109)

return TailCall1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21110)


panic("unreachable")
})

tmp21111 := Call2(__e, ns2_1set, symshen_4initialise_1lambda_1forms, tmp20558)


_ = tmp21111

tmp21112 := MakeNative0(func(__e *ControlFlow) Obj {
tmp21113 := Call0(__e, SymFn(symshen_4initialise_1environment))


_ = tmp21113

tmp21114 := Call0(__e, SymFn(symshen_4initialise_1lambda_1forms))


_ = tmp21114

return TailCall0(__e, SymFn(symshen_4initialise_1signedfuncs))


panic("unreachable")
})

return TailCall2(__e, ns2_1set, symshen_4initialise, tmp21112)




panic("unreachable")
})

