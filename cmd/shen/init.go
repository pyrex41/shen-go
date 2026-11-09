package main

import . "github.com/pyrex41/shen-go/kl"

var InitMain = MakeNative0(func(__e *ControlFlow) Obj {
tmp18096 := MakeNative0(func(__e *ControlFlow) Obj {
tmp18097 := PrimGuard2(__e, symset, PrimSet, symshen_4_dhistory_d, Nil)

_ = tmp18097

tmp18098 := PrimGuard2(__e, symset, PrimSet, symshen_4_dtc_d, False)

_ = tmp18098

tmp18099 := Call1(__e, SymFn(symshen_4dict), MakeInteger(20000))


tmp18100 := PrimGuard2(__e, symset, PrimSet, sym_dproperty_1vector_d, tmp18099)

_ = tmp18100

tmp18101 := MakeNative1(func(__e *ControlFlow, X Obj) Obj {
return TailCall1(__e, SymFn(symshen_4macros), X)
panic("unreachable")
})

tmp18102 := GCons(__e, symshen_4macros, tmp18101)

tmp18103 := GCons(__e, tmp18102, Nil)

tmp18104 := PrimGuard2(__e, symset, PrimSet, sym_dmacros_d, tmp18103)

_ = tmp18104

tmp18105 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsource_1form_1handlers_d, Nil)

_ = tmp18105

tmp18106 := PrimGuard2(__e, symset, PrimSet, symshen_4_dgensym_d, MakeInteger(0))

_ = tmp18106

tmp18107 := PrimGuard2(__e, symset, PrimSet, symshen_4_dtracking_d, Nil)

_ = tmp18107

tmp18108 := PrimGuard2(__e, symset, PrimSet, symshen_4_dprofiled_d, Nil)

_ = tmp18108

tmp18109 := GCons(__e, symtype, Nil)

tmp18110 := GCons(__e, symshen_4input_1h_7, tmp18109)

tmp18111 := GCons(__e, symopen, tmp18110)

tmp18112 := GCons(__e, symset, tmp18111)

tmp18113 := GCons(__e, symwhere, tmp18112)

tmp18114 := GCons(__e, symlet, tmp18113)

tmp18115 := GCons(__e, symlambda, tmp18114)

tmp18116 := GCons(__e, symcons, tmp18115)

tmp18117 := GCons(__e, sym_8v, tmp18116)

tmp18118 := GCons(__e, sym_8s, tmp18117)

tmp18119 := GCons(__e, sym_8p, tmp18118)

tmp18120 := PrimGuard2(__e, symset, PrimSet, symshen_4_dspecial_d, tmp18119)

_ = tmp18120

tmp18121 := PrimGuard2(__e, symset, PrimSet, symshen_4_dextraspecial_d, Nil)

_ = tmp18121

tmp18122 := PrimGuard2(__e, symset, PrimSet, symshen_4_dspy_d, False)

_ = tmp18122

tmp18123 := PrimGuard2(__e, symset, PrimSet, symshen_4_ddatatypes_d, Nil)

_ = tmp18123

tmp18124 := PrimGuard2(__e, symset, PrimSet, symshen_4_dalldatatypes_d, Nil)

_ = tmp18124

tmp18125 := PrimGuard2(__e, symset, PrimSet, symshen_4_dshen_1type_1theory_1enabled_2_d, True)

_ = tmp18125

tmp18126 := PrimGuard2(__e, symset, PrimSet, symshen_4_dpackage_d, symnull)

_ = tmp18126

tmp18127 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsynonyms_d, Nil)

_ = tmp18127

tmp18128 := PrimGuard2(__e, symset, PrimSet, symshen_4_doccurs_d, True)

_ = tmp18128

tmp18129 := PrimGuard2(__e, symset, PrimSet, symshen_4_dfactorise_2_d, False)

_ = tmp18129

tmp18130 := PrimGuard2(__e, symset, PrimSet, symshen_4_dmaxinferences_d, MakeInteger(1000000))

_ = tmp18130

tmp18131 := PrimGuard2(__e, symset, PrimSet, sym_dmaximum_1print_1sequence_1size_d, MakeInteger(20))

_ = tmp18131

tmp18132 := PrimGuard2(__e, symset, PrimSet, symshen_4_dcall_d, MakeInteger(0))

_ = tmp18132

tmp18133 := PrimGuard2(__e, symset, PrimSet, symshen_4_dinfs_d, MakeInteger(0))

_ = tmp18133

tmp18134 := PrimGuard2(__e, symset, PrimSet, sym_dhush_d, False)

_ = tmp18134

tmp18135 := PrimGuard2(__e, symset, PrimSet, symshen_4_doptimise_d, False)

_ = tmp18135

tmp18136 := PrimGuard2(__e, symset, PrimSet, sym_dversion_d, MakeString("43.4"))

_ = tmp18136

tmp18137 := PrimGuard2(__e, symset, PrimSet, symshen_4_dnames_d, Nil)

_ = tmp18137

tmp18138 := PrimGuard2(__e, symset, PrimSet, symshen_4_dstep_d, False)

_ = tmp18138

tmp18139 := PrimGuard2(__e, symset, PrimSet, symshen_4_dit_d, MakeString(""))

_ = tmp18139

tmp18140 := PrimGuard2(__e, symset, PrimSet, symshen_4_dresidue_d, Nil)

_ = tmp18140

tmp18141 := PrimGuard2(__e, symset, PrimSet, sym_dabsolute_d, Nil)

_ = tmp18141

tmp18142 := PrimGuard2(__e, symset, PrimSet, symshen_4_dprolog_1memory_d, MakeInteger(1000))

_ = tmp18142

tmp18143 := PrimGuard2(__e, symset, PrimSet, symshen_4_dloading_2_d, False)

_ = tmp18143

tmp18144 := PrimGuard2(__e, symset, PrimSet, symshen_4_duserdefs_d, Nil)

_ = tmp18144

tmp18145 := MakeNative1(func(__e *ControlFlow, X Obj) Obj {
return X
panic("unreachable")
})

tmp18146 := PrimGuard2(__e, symset, PrimSet, symshen_4_ddemodulation_1function_d, tmp18145)

_ = tmp18146

tmp18147 := PrimGuard2(__e, symset, PrimSet, symshen_4_dcustom_1pattern_1compiler_d, False)

_ = tmp18147

tmp18148 := PrimGuard2(__e, symset, PrimSet, symshen_4_dcustom_1pattern_1reducer_d, False)

_ = tmp18148

tmp18151 := Call1(__e, SymFn(symbound_2), sym_dhome_1directory_d)


tmp18152 := GNot(__e, tmp18151)

var ifres18149 Obj

if True == tmp18152 {
tmp18150 := PrimGuard2(__e, symset, PrimSet, sym_dhome_1directory_d, MakeString(""))

ifres18149 = tmp18150


} else {
ifres18149 = symshen_4skip


}

_ = ifres18149

tmp18156 := Call1(__e, SymFn(symbound_2), sym_dsterror_d)


tmp18157 := GNot(__e, tmp18156)

var ifres18153 Obj

if True == tmp18157 {
tmp18154 := GValue(__e, sym_dstoutput_d)

tmp18155 := PrimGuard2(__e, symset, PrimSet, sym_dsterror_d, tmp18154)

ifres18153 = tmp18155


} else {
ifres18153 = symshen_4skip


}

_ = ifres18153

tmp18158 := Call1(__e, SymFn(symprolog_1memory), MakeInteger(10000))


_ = tmp18158

tmp18159 := PrimGuard2(__e, symset, PrimSet, symshen_4_dloading_2_d, False)

_ = tmp18159

tmp18160 := GCons(__e, MakeInteger(2), Nil)

tmp18161 := GCons(__e, sym_8s, tmp18160)

tmp18162 := GCons(__e, MakeInteger(2), tmp18161)

tmp18163 := GCons(__e, sym_8v, tmp18162)

tmp18164 := GCons(__e, MakeInteger(2), tmp18163)

tmp18165 := GCons(__e, sym_8p, tmp18164)

tmp18166 := GCons(__e, MakeInteger(1), tmp18165)

tmp18167 := GCons(__e, sym_5_b_6, tmp18166)

tmp18168 := GCons(__e, MakeInteger(1), tmp18167)

tmp18169 := GCons(__e, sym_5end_6, tmp18168)

tmp18170 := GCons(__e, MakeInteger(1), tmp18169)

tmp18171 := GCons(__e, sym_5e_6, tmp18170)

tmp18172 := GCons(__e, MakeInteger(2), tmp18171)

tmp18173 := GCons(__e, sym_a_a, tmp18172)

tmp18174 := GCons(__e, MakeInteger(2), tmp18173)

tmp18175 := GCons(__e, sym_1, tmp18174)

tmp18176 := GCons(__e, MakeInteger(2), tmp18175)

tmp18177 := GCons(__e, sym_c, tmp18176)

tmp18178 := GCons(__e, MakeInteger(2), tmp18177)

tmp18179 := GCons(__e, sym_d, tmp18178)

tmp18180 := GCons(__e, MakeInteger(2), tmp18179)

tmp18181 := GCons(__e, sym_7, tmp18180)

tmp18182 := GCons(__e, MakeInteger(1), tmp18181)

tmp18183 := GCons(__e, symy_1or_1n_2, tmp18182)

tmp18184 := GCons(__e, MakeInteger(2), tmp18183)

tmp18185 := GCons(__e, symwrite_1to_1file, tmp18184)

tmp18186 := GCons(__e, MakeInteger(2), tmp18185)

tmp18187 := GCons(__e, symwrite_1string, tmp18186)

tmp18188 := GCons(__e, MakeInteger(2), tmp18187)

tmp18189 := GCons(__e, symwrite_1byte, tmp18188)

tmp18190 := GCons(__e, MakeInteger(5), tmp18189)

tmp18191 := GCons(__e, symwhen, tmp18190)

tmp18192 := GCons(__e, MakeInteger(0), tmp18191)

tmp18193 := GCons(__e, symversion, tmp18192)

tmp18194 := GCons(__e, MakeInteger(5), tmp18193)

tmp18195 := GCons(__e, symvar_2, tmp18194)

tmp18196 := GCons(__e, MakeInteger(1), tmp18195)

tmp18197 := GCons(__e, symvariable_2, tmp18196)

tmp18198 := GCons(__e, MakeInteger(1), tmp18197)

tmp18199 := GCons(__e, symvalue, tmp18198)

tmp18200 := GCons(__e, MakeInteger(3), tmp18199)

tmp18201 := GCons(__e, symvector_1_6, tmp18200)

tmp18202 := GCons(__e, MakeInteger(1), tmp18201)

tmp18203 := GCons(__e, symvector_2, tmp18202)

tmp18204 := GCons(__e, MakeInteger(1), tmp18203)

tmp18205 := GCons(__e, symvector, tmp18204)

tmp18206 := GCons(__e, MakeInteger(0), tmp18205)

tmp18207 := GCons(__e, symuserdefs, tmp18206)

tmp18208 := GCons(__e, MakeInteger(2), tmp18207)

tmp18209 := GCons(__e, symupdate_1lambda_1table, tmp18208)

tmp18210 := GCons(__e, MakeInteger(1), tmp18209)

tmp18211 := GCons(__e, symundefmacro, tmp18210)

tmp18212 := GCons(__e, MakeInteger(1), tmp18211)

tmp18213 := GCons(__e, symuntrack, tmp18212)

tmp18214 := GCons(__e, MakeInteger(2), tmp18213)

tmp18215 := GCons(__e, symunion, tmp18214)

tmp18216 := GCons(__e, MakeInteger(1), tmp18215)

tmp18217 := GCons(__e, symunprofile, tmp18216)

tmp18218 := GCons(__e, MakeInteger(3), tmp18217)

tmp18219 := GCons(__e, symunput, tmp18218)

tmp18220 := GCons(__e, MakeInteger(1), tmp18219)

tmp18221 := GCons(__e, symundefmacro, tmp18220)

tmp18222 := GCons(__e, MakeInteger(1), tmp18221)

tmp18223 := GCons(__e, symunabsolute, tmp18222)

tmp18224 := GCons(__e, MakeInteger(5), tmp18223)

tmp18225 := GCons(__e, symreturn, tmp18224)

tmp18226 := GCons(__e, MakeInteger(2), tmp18225)

tmp18227 := GCons(__e, symtype, tmp18226)

tmp18228 := GCons(__e, MakeInteger(1), tmp18227)

tmp18229 := GCons(__e, symtuple_2, tmp18228)

tmp18230 := GCons(__e, MakeInteger(2), tmp18229)

tmp18231 := GCons(__e, symtrap_1error, tmp18230)

tmp18232 := GCons(__e, MakeInteger(0), tmp18231)

tmp18233 := GCons(__e, symtracked, tmp18232)

tmp18234 := GCons(__e, MakeInteger(1), tmp18233)

tmp18235 := GCons(__e, symtrack, tmp18234)

tmp18236 := GCons(__e, MakeInteger(1), tmp18235)

tmp18237 := GCons(__e, symtlstr, tmp18236)

tmp18238 := GCons(__e, MakeInteger(1), tmp18237)

tmp18239 := GCons(__e, symthaw, tmp18238)

tmp18240 := GCons(__e, MakeInteger(0), tmp18239)

tmp18241 := GCons(__e, symtc_2, tmp18240)

tmp18242 := GCons(__e, MakeInteger(1), tmp18241)

tmp18243 := GCons(__e, symtc, tmp18242)

tmp18244 := GCons(__e, MakeInteger(1), tmp18243)

tmp18245 := GCons(__e, symtl, tmp18244)

tmp18246 := GCons(__e, MakeInteger(1), tmp18245)

tmp18247 := GCons(__e, symtail, tmp18246)

tmp18248 := GCons(__e, MakeInteger(1), tmp18247)

tmp18249 := GCons(__e, symsystemf, tmp18248)

tmp18250 := GCons(__e, MakeInteger(1), tmp18249)

tmp18251 := GCons(__e, symsymbol_2, tmp18250)

tmp18252 := GCons(__e, MakeInteger(1), tmp18251)

tmp18253 := GCons(__e, symsum, tmp18252)

tmp18254 := GCons(__e, MakeInteger(3), tmp18253)

tmp18255 := GCons(__e, symsubst, tmp18254)

tmp18256 := GCons(__e, MakeInteger(1), tmp18255)

tmp18257 := GCons(__e, symstring_2, tmp18256)

tmp18258 := GCons(__e, MakeInteger(1), tmp18257)

tmp18259 := GCons(__e, symstring_1_6symbol, tmp18258)

tmp18260 := GCons(__e, MakeInteger(1), tmp18259)

tmp18261 := GCons(__e, symstring_1_6n, tmp18260)

tmp18262 := GCons(__e, MakeInteger(1), tmp18261)

tmp18263 := GCons(__e, symstr, tmp18262)

tmp18264 := GCons(__e, MakeInteger(0), tmp18263)

tmp18265 := GCons(__e, symstoutput, tmp18264)

tmp18266 := GCons(__e, MakeInteger(0), tmp18265)

tmp18267 := GCons(__e, symstinput, tmp18266)

tmp18268 := GCons(__e, MakeInteger(0), tmp18267)

tmp18269 := GCons(__e, symstep_2, tmp18268)

tmp18270 := GCons(__e, MakeInteger(1), tmp18269)

tmp18271 := GCons(__e, symstep, tmp18270)

tmp18272 := GCons(__e, MakeInteger(0), tmp18271)

tmp18273 := GCons(__e, symspy_2, tmp18272)

tmp18274 := GCons(__e, MakeInteger(1), tmp18273)

tmp18275 := GCons(__e, symspy, tmp18274)

tmp18276 := GCons(__e, MakeInteger(2), tmp18275)

tmp18277 := GCons(__e, symspecialise, tmp18276)

tmp18278 := GCons(__e, MakeInteger(1), tmp18277)

tmp18279 := GCons(__e, symsnd, tmp18278)

tmp18280 := GCons(__e, MakeInteger(1), tmp18279)

tmp18281 := GCons(__e, symsimple_1error, tmp18280)

tmp18282 := GCons(__e, MakeInteger(2), tmp18281)

tmp18283 := GCons(__e, symset, tmp18282)

tmp18284 := GCons(__e, MakeInteger(1), tmp18283)

tmp18285 := GCons(__e, symreverse, tmp18284)

tmp18286 := GCons(__e, MakeInteger(2), tmp18285)

tmp18287 := GCons(__e, symremove, tmp18286)

tmp18288 := GCons(__e, MakeInteger(1), tmp18287)

tmp18289 := GCons(__e, symshen_4unregister_1source_1form, tmp18288)

tmp18290 := GCons(__e, MakeInteger(2), tmp18289)

tmp18291 := GCons(__e, symshen_4register_1source_1form, tmp18290)

tmp18292 := GCons(__e, MakeInteger(0), tmp18291)

tmp18293 := GCons(__e, symrelease, tmp18292)

tmp18294 := GCons(__e, MakeInteger(1), tmp18293)

tmp18295 := GCons(__e, symreceive, tmp18294)

tmp18296 := GCons(__e, MakeInteger(1), tmp18295)

tmp18297 := GCons(__e, symread_1unit_1string, tmp18296)

tmp18298 := GCons(__e, MakeInteger(1), tmp18297)

tmp18299 := GCons(__e, symread_1from_1string_1unprocessed, tmp18298)

tmp18300 := GCons(__e, MakeInteger(1), tmp18299)

tmp18301 := GCons(__e, symread_1from_1string, tmp18300)

tmp18302 := GCons(__e, MakeInteger(1), tmp18301)

tmp18303 := GCons(__e, symread_1byte, tmp18302)

tmp18304 := GCons(__e, MakeInteger(1), tmp18303)

tmp18305 := GCons(__e, symread, tmp18304)

tmp18306 := GCons(__e, MakeInteger(1), tmp18305)

tmp18307 := GCons(__e, symread_1file, tmp18306)

tmp18308 := GCons(__e, MakeInteger(1), tmp18307)

tmp18309 := GCons(__e, symread_1file_1as_1bytelist, tmp18308)

tmp18310 := GCons(__e, MakeInteger(1), tmp18309)

tmp18311 := GCons(__e, symread_1file_1as_1string, tmp18310)

tmp18312 := GCons(__e, MakeInteger(4), tmp18311)

tmp18313 := GCons(__e, symput, tmp18312)

tmp18314 := GCons(__e, MakeInteger(1), tmp18313)

tmp18315 := GCons(__e, symprotect, tmp18314)

tmp18316 := GCons(__e, MakeInteger(1), tmp18315)

tmp18317 := GCons(__e, sympreclude_1all_1but, tmp18316)

tmp18318 := GCons(__e, MakeInteger(1), tmp18317)

tmp18319 := GCons(__e, sympreclude, tmp18318)

tmp18320 := GCons(__e, MakeInteger(1), tmp18319)

tmp18321 := GCons(__e, symps, tmp18320)

tmp18322 := GCons(__e, MakeInteger(2), tmp18321)

tmp18323 := GCons(__e, sympr, tmp18322)

tmp18324 := GCons(__e, MakeInteger(1), tmp18323)

tmp18325 := GCons(__e, symprofile_1results, tmp18324)

tmp18326 := GCons(__e, MakeInteger(1), tmp18325)

tmp18327 := GCons(__e, symprolog_1memory, tmp18326)

tmp18328 := GCons(__e, MakeInteger(1), tmp18327)

tmp18329 := GCons(__e, symshen_4printF, tmp18328)

tmp18330 := GCons(__e, MakeInteger(1), tmp18329)

tmp18331 := GCons(__e, symshen_4print_1freshterm, tmp18330)

tmp18332 := GCons(__e, MakeInteger(1), tmp18331)

tmp18333 := GCons(__e, symshen_4print_1prolog_1vector, tmp18332)

tmp18334 := GCons(__e, MakeInteger(1), tmp18333)

tmp18335 := GCons(__e, symprofile, tmp18334)

tmp18336 := GCons(__e, MakeInteger(1), tmp18335)

tmp18337 := GCons(__e, symprint, tmp18336)

tmp18338 := GCons(__e, MakeInteger(1), tmp18337)

tmp18339 := GCons(__e, sympreclude_1all_1but, tmp18338)

tmp18340 := GCons(__e, MakeInteger(2), tmp18339)

tmp18341 := GCons(__e, sympos, tmp18340)

tmp18342 := GCons(__e, MakeInteger(0), tmp18341)

tmp18343 := GCons(__e, symporters, tmp18342)

tmp18344 := GCons(__e, MakeInteger(0), tmp18343)

tmp18345 := GCons(__e, symport, tmp18344)

tmp18346 := GCons(__e, MakeInteger(1), tmp18345)

tmp18347 := GCons(__e, sympackage_2, tmp18346)

tmp18348 := GCons(__e, MakeInteger(3), tmp18347)

tmp18349 := GCons(__e, sympackage, tmp18348)

tmp18350 := GCons(__e, MakeInteger(0), tmp18349)

tmp18351 := GCons(__e, symos, tmp18350)

tmp18352 := GCons(__e, MakeInteger(2), tmp18351)

tmp18353 := GCons(__e, symor, tmp18352)

tmp18354 := GCons(__e, MakeInteger(0), tmp18353)

tmp18355 := GCons(__e, symoptimise_2, tmp18354)

tmp18356 := GCons(__e, MakeInteger(1), tmp18355)

tmp18357 := GCons(__e, symoptimise, tmp18356)

tmp18358 := GCons(__e, MakeInteger(2), tmp18357)

tmp18359 := GCons(__e, symopen, tmp18358)

tmp18360 := GCons(__e, MakeInteger(1), tmp18359)

tmp18361 := GCons(__e, symoccurs_1check, tmp18360)

tmp18362 := GCons(__e, MakeInteger(0), tmp18361)

tmp18363 := GCons(__e, symoccurs_2, tmp18362)

tmp18364 := GCons(__e, MakeInteger(2), tmp18363)

tmp18365 := GCons(__e, symoccurrences, tmp18364)

tmp18366 := GCons(__e, MakeInteger(1), tmp18365)

tmp18367 := GCons(__e, symoccurs_1check, tmp18366)

tmp18368 := GCons(__e, MakeInteger(1), tmp18367)

tmp18369 := GCons(__e, symnumber_2, tmp18368)

tmp18370 := GCons(__e, MakeInteger(1), tmp18369)

tmp18371 := GCons(__e, symn_1_6string, tmp18370)

tmp18372 := GCons(__e, MakeInteger(2), tmp18371)

tmp18373 := GCons(__e, symnth, tmp18372)

tmp18374 := GCons(__e, MakeInteger(1), tmp18373)

tmp18375 := GCons(__e, symnot, tmp18374)

tmp18376 := GCons(__e, MakeInteger(1), tmp18375)

tmp18377 := GCons(__e, symnl, tmp18376)

tmp18378 := GCons(__e, MakeInteger(1), tmp18377)

tmp18379 := GCons(__e, symmaxinferences, tmp18378)

tmp18380 := GCons(__e, MakeInteger(2), tmp18379)

tmp18381 := GCons(__e, symmapcan, tmp18380)

tmp18382 := GCons(__e, MakeInteger(2), tmp18381)

tmp18383 := GCons(__e, symmap, tmp18382)

tmp18384 := GCons(__e, MakeInteger(1), tmp18383)

tmp18385 := GCons(__e, symmacroexpand, tmp18384)

tmp18386 := GCons(__e, MakeInteger(1), tmp18385)

tmp18387 := GCons(__e, symvector, tmp18386)

tmp18388 := GCons(__e, MakeInteger(2), tmp18387)

tmp18389 := GCons(__e, sym_5_a, tmp18388)

tmp18390 := GCons(__e, MakeInteger(2), tmp18389)

tmp18391 := GCons(__e, sym_5, tmp18390)

tmp18392 := GCons(__e, MakeInteger(1), tmp18391)

tmp18393 := GCons(__e, symload, tmp18392)

tmp18394 := GCons(__e, MakeInteger(1), tmp18393)

tmp18395 := GCons(__e, symlist, tmp18394)

tmp18396 := GCons(__e, MakeInteger(1), tmp18395)

tmp18397 := GCons(__e, symlineread, tmp18396)

tmp18398 := GCons(__e, MakeInteger(1), tmp18397)

tmp18399 := GCons(__e, symlimit, tmp18398)

tmp18400 := GCons(__e, MakeInteger(1), tmp18399)

tmp18401 := GCons(__e, symlength, tmp18400)

tmp18402 := GCons(__e, MakeInteger(0), tmp18401)

tmp18403 := GCons(__e, symlanguage, tmp18402)

tmp18404 := GCons(__e, MakeInteger(6), tmp18403)

tmp18405 := GCons(__e, symis_b, tmp18404)

tmp18406 := GCons(__e, MakeInteger(6), tmp18405)

tmp18407 := GCons(__e, symis, tmp18406)

tmp18408 := GCons(__e, MakeInteger(0), tmp18407)

tmp18409 := GCons(__e, symit, tmp18408)

tmp18410 := GCons(__e, MakeInteger(1), tmp18409)

tmp18411 := GCons(__e, syminternal, tmp18410)

tmp18412 := GCons(__e, MakeInteger(2), tmp18411)

tmp18413 := GCons(__e, symintersection, tmp18412)

tmp18414 := GCons(__e, MakeInteger(1), tmp18413)

tmp18415 := GCons(__e, syminclude_1all_1but, tmp18414)

tmp18416 := GCons(__e, MakeInteger(0), tmp18415)

tmp18417 := GCons(__e, symimplementation, tmp18416)

tmp18418 := GCons(__e, MakeInteger(2), tmp18417)

tmp18419 := GCons(__e, syminput_7, tmp18418)

tmp18420 := GCons(__e, MakeInteger(1), tmp18419)

tmp18421 := GCons(__e, syminput, tmp18420)

tmp18422 := GCons(__e, MakeInteger(0), tmp18421)

tmp18423 := GCons(__e, syminferences, tmp18422)

tmp18424 := GCons(__e, MakeInteger(1), tmp18423)

tmp18425 := GCons(__e, symintern, tmp18424)

tmp18426 := GCons(__e, MakeInteger(1), tmp18425)

tmp18427 := GCons(__e, syminternal, tmp18426)

tmp18428 := GCons(__e, MakeInteger(1), tmp18427)

tmp18429 := GCons(__e, syminteger_2, tmp18428)

tmp18430 := GCons(__e, MakeInteger(1), tmp18429)

tmp18431 := GCons(__e, symin_1package, tmp18430)

tmp18432 := GCons(__e, MakeInteger(0), tmp18431)

tmp18433 := GCons(__e, symincluded, tmp18432)

tmp18434 := GCons(__e, MakeInteger(1), tmp18433)

tmp18435 := GCons(__e, syminclude, tmp18434)

tmp18436 := GCons(__e, MakeInteger(3), tmp18435)

tmp18437 := GCons(__e, symif, tmp18436)

tmp18438 := GCons(__e, MakeInteger(1), tmp18437)

tmp18439 := GCons(__e, symhush, tmp18438)

tmp18440 := GCons(__e, MakeInteger(0), tmp18439)

tmp18441 := GCons(__e, symhush_2, tmp18440)

tmp18442 := GCons(__e, MakeInteger(1), tmp18441)

tmp18443 := GCons(__e, symhead, tmp18442)

tmp18444 := GCons(__e, MakeInteger(1), tmp18443)

tmp18445 := GCons(__e, symhdstr, tmp18444)

tmp18446 := GCons(__e, MakeInteger(1), tmp18445)

tmp18447 := GCons(__e, symhdv, tmp18446)

tmp18448 := GCons(__e, MakeInteger(1), tmp18447)

tmp18449 := GCons(__e, symhd, tmp18448)

tmp18450 := GCons(__e, MakeInteger(2), tmp18449)

tmp18451 := GCons(__e, symhash, tmp18450)

tmp18452 := GCons(__e, MakeInteger(2), tmp18451)

tmp18453 := GCons(__e, sym_a, tmp18452)

tmp18454 := GCons(__e, MakeInteger(2), tmp18453)

tmp18455 := GCons(__e, sym_6_a, tmp18454)

tmp18456 := GCons(__e, MakeInteger(2), tmp18455)

tmp18457 := GCons(__e, sym_6, tmp18456)

tmp18458 := GCons(__e, MakeInteger(2), tmp18457)

tmp18459 := GCons(__e, sym_5_1vector, tmp18458)

tmp18460 := GCons(__e, MakeInteger(2), tmp18459)

tmp18461 := GCons(__e, sym_5_1address, tmp18460)

tmp18462 := GCons(__e, MakeInteger(3), tmp18461)

tmp18463 := GCons(__e, symaddress_1_6, tmp18462)

tmp18464 := GCons(__e, MakeInteger(1), tmp18463)

tmp18465 := GCons(__e, symget_1time, tmp18464)

tmp18466 := GCons(__e, MakeInteger(3), tmp18465)

tmp18467 := GCons(__e, symget, tmp18466)

tmp18468 := GCons(__e, MakeInteger(1), tmp18467)

tmp18469 := GCons(__e, symgensym, tmp18468)

tmp18470 := GCons(__e, MakeInteger(1), tmp18469)

tmp18471 := GCons(__e, symfunction, tmp18470)

tmp18472 := GCons(__e, MakeInteger(1), tmp18471)

tmp18473 := GCons(__e, symfn, tmp18472)

tmp18474 := GCons(__e, MakeInteger(1), tmp18473)

tmp18475 := GCons(__e, symfst, tmp18474)

tmp18476 := GCons(__e, MakeInteger(0), tmp18475)

tmp18477 := GCons(__e, symfresh, tmp18476)

tmp18478 := GCons(__e, MakeInteger(1), tmp18477)

tmp18479 := GCons(__e, symfreeze, tmp18478)

tmp18480 := GCons(__e, MakeInteger(5), tmp18479)

tmp18481 := GCons(__e, symfork, tmp18480)

tmp18482 := GCons(__e, MakeInteger(1), tmp18481)

tmp18483 := GCons(__e, symforeign, tmp18482)

tmp18484 := GCons(__e, MakeInteger(7), tmp18483)

tmp18485 := GCons(__e, symfindall, tmp18484)

tmp18486 := GCons(__e, MakeInteger(2), tmp18485)

tmp18487 := GCons(__e, symfix, tmp18486)

tmp18488 := GCons(__e, MakeInteger(0), tmp18487)

tmp18489 := GCons(__e, symfail, tmp18488)

tmp18490 := GCons(__e, MakeInteger(2), tmp18489)

tmp18491 := GCons(__e, symfail_1if, tmp18490)

tmp18492 := GCons(__e, MakeInteger(0), tmp18491)

tmp18493 := GCons(__e, symfactorise_2, tmp18492)

tmp18494 := GCons(__e, MakeInteger(1), tmp18493)

tmp18495 := GCons(__e, symfactorise, tmp18494)

tmp18496 := GCons(__e, MakeInteger(1), tmp18495)

tmp18497 := GCons(__e, symexternal, tmp18496)

tmp18498 := GCons(__e, MakeInteger(1), tmp18497)

tmp18499 := GCons(__e, symexplode, tmp18498)

tmp18500 := GCons(__e, MakeInteger(1), tmp18499)

tmp18501 := GCons(__e, symeval_1kl, tmp18500)

tmp18502 := GCons(__e, MakeInteger(1), tmp18501)

tmp18503 := GCons(__e, symeval, tmp18502)

tmp18504 := GCons(__e, MakeInteger(1), tmp18503)

tmp18505 := GCons(__e, symerror_1to_1string, tmp18504)

tmp18506 := GCons(__e, MakeInteger(1), tmp18505)

tmp18507 := GCons(__e, symexternal, tmp18506)

tmp18508 := GCons(__e, MakeInteger(1), tmp18507)

tmp18509 := GCons(__e, symenable_1type_1theory, tmp18508)

tmp18510 := GCons(__e, MakeInteger(1), tmp18509)

tmp18511 := GCons(__e, symempty_2, tmp18510)

tmp18512 := GCons(__e, MakeInteger(2), tmp18511)

tmp18513 := GCons(__e, symelement_2, tmp18512)

tmp18514 := GCons(__e, MakeInteger(2), tmp18513)

tmp18515 := GCons(__e, symdo, tmp18514)

tmp18516 := GCons(__e, MakeInteger(2), tmp18515)

tmp18517 := GCons(__e, symdifference, tmp18516)

tmp18518 := GCons(__e, MakeInteger(1), tmp18517)

tmp18519 := GCons(__e, symdestroy, tmp18518)

tmp18520 := GCons(__e, MakeInteger(2), tmp18519)

tmp18521 := GCons(__e, symdeclare, tmp18520)

tmp18522 := GCons(__e, MakeInteger(0), tmp18521)

tmp18523 := GCons(__e, symdatatypes, tmp18522)

tmp18524 := GCons(__e, MakeInteger(1), tmp18523)

tmp18525 := GCons(__e, symclose, tmp18524)

tmp18526 := GCons(__e, MakeInteger(2), tmp18525)

tmp18527 := GCons(__e, symcn, tmp18526)

tmp18528 := GCons(__e, MakeInteger(1), tmp18527)

tmp18529 := GCons(__e, symcons_2, tmp18528)

tmp18530 := GCons(__e, MakeInteger(2), tmp18529)

tmp18531 := GCons(__e, symcons, tmp18530)

tmp18532 := GCons(__e, MakeInteger(2), tmp18531)

tmp18533 := GCons(__e, symconcat, tmp18532)

tmp18534 := GCons(__e, MakeInteger(2), tmp18533)

tmp18535 := GCons(__e, symcompile, tmp18534)

tmp18536 := GCons(__e, MakeInteger(1), tmp18535)

tmp18537 := GCons(__e, symchar_1stream_2, tmp18536)

tmp18538 := GCons(__e, MakeInteger(1), tmp18537)

tmp18539 := GCons(__e, symcd, tmp18538)

tmp18540 := GCons(__e, MakeInteger(5), tmp18539)

tmp18541 := GCons(__e, symcall, tmp18540)

tmp18542 := GCons(__e, MakeInteger(6), tmp18541)

tmp18543 := GCons(__e, symbind, tmp18542)

tmp18544 := GCons(__e, MakeInteger(1), tmp18543)

tmp18545 := GCons(__e, symbound_2, tmp18544)

tmp18546 := GCons(__e, MakeInteger(1), tmp18545)

tmp18547 := GCons(__e, symbootstrap, tmp18546)

tmp18548 := GCons(__e, MakeInteger(1), tmp18547)

tmp18549 := GCons(__e, symboolean_2, tmp18548)

tmp18550 := GCons(__e, MakeInteger(1), tmp18549)

tmp18551 := GCons(__e, symatom_2, tmp18550)

tmp18552 := GCons(__e, MakeInteger(2), tmp18551)

tmp18553 := GCons(__e, symassoc, tmp18552)

tmp18554 := GCons(__e, MakeInteger(1), tmp18553)

tmp18555 := GCons(__e, symarity, tmp18554)

tmp18556 := GCons(__e, MakeInteger(2), tmp18555)

tmp18557 := GCons(__e, symappend, tmp18556)

tmp18558 := GCons(__e, MakeInteger(2), tmp18557)

tmp18559 := GCons(__e, symand, tmp18558)

tmp18560 := GCons(__e, MakeInteger(2), tmp18559)

tmp18561 := GCons(__e, symadjoin, tmp18560)

tmp18562 := GCons(__e, MakeInteger(3), tmp18561)

tmp18563 := GCons(__e, symaddress_1_6, tmp18562)

tmp18564 := GCons(__e, MakeInteger(1), tmp18563)

tmp18565 := GCons(__e, symabsvector, tmp18564)

tmp18566 := GCons(__e, MakeInteger(1), tmp18565)

tmp18567 := GCons(__e, symabsvector_2, tmp18566)

tmp18568 := GCons(__e, MakeInteger(1), tmp18567)

tmp18569 := GCons(__e, symabsolute, tmp18568)

tmp18570 := GCons(__e, MakeInteger(0), tmp18569)

tmp18571 := GCons(__e, symabort, tmp18570)

tmp18572 := Call1(__e, SymFn(symshen_4initialise_1arity_1table), tmp18571)


_ = tmp18572

tmp18573 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":"))

tmp18574 := PrimGuard1(__e, symintern, PrimIntern, MakeString(";"))

tmp18575 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":="))

tmp18576 := PrimGuard1(__e, symintern, PrimIntern, MakeString(","))

tmp18577 := Call1(__e, SymFn(symvector), MakeInteger(0))


tmp18578 := PrimGuard1(__e, symintern, PrimIntern, MakeString("bar!"))

tmp18579 := GCons(__e, symabort, Nil)

tmp18580 := GCons(__e, symabsolute, tmp18579)

tmp18581 := GCons(__e, symabsvector, tmp18580)

tmp18582 := GCons(__e, symabsvector_2, tmp18581)

tmp18583 := GCons(__e, symaddress_1_6, tmp18582)

tmp18584 := GCons(__e, sym_5_1address, tmp18583)

tmp18585 := GCons(__e, symadjoin, tmp18584)

tmp18586 := GCons(__e, symand, tmp18585)

tmp18587 := GCons(__e, symappend, tmp18586)

tmp18588 := GCons(__e, symarity, tmp18587)

tmp18589 := GCons(__e, symassoc, tmp18588)

tmp18590 := GCons(__e, symassertz, tmp18589)

tmp18591 := GCons(__e, symasserta, tmp18590)

tmp18592 := GCons(__e, symatom_2, tmp18591)

tmp18593 := GCons(__e, symstep_2, tmp18592)

tmp18594 := GCons(__e, symspy_2, tmp18593)

tmp18595 := GCons(__e, tmp18578, tmp18594)

tmp18596 := GCons(__e, symbootstrap, tmp18595)

tmp18597 := GCons(__e, symboolean, tmp18596)

tmp18598 := GCons(__e, symboolean_2, tmp18597)

tmp18599 := GCons(__e, symbound_2, tmp18598)

tmp18600 := GCons(__e, symbind, tmp18599)

tmp18601 := GCons(__e, symclose, tmp18600)

tmp18602 := GCons(__e, symchar_1stream_2, tmp18601)

tmp18603 := GCons(__e, symcall, tmp18602)

tmp18604 := GCons(__e, symcases, tmp18603)

tmp18605 := GCons(__e, symcd, tmp18604)

tmp18606 := GCons(__e, symcompile, tmp18605)

tmp18607 := GCons(__e, symconcat, tmp18606)

tmp18608 := GCons(__e, symcond, tmp18607)

tmp18609 := GCons(__e, symcons, tmp18608)

tmp18610 := GCons(__e, symcons_2, tmp18609)

tmp18611 := GCons(__e, symcn, tmp18610)

tmp18612 := GCons(__e, symctxt, tmp18611)

tmp18613 := GCons(__e, symdatatypes, tmp18612)

tmp18614 := GCons(__e, symdatatype, tmp18613)

tmp18615 := GCons(__e, symdeclare, tmp18614)

tmp18616 := GCons(__e, symdefprolog, tmp18615)

tmp18617 := GCons(__e, symdefcc, tmp18616)

tmp18618 := GCons(__e, symdefmacro, tmp18617)

tmp18619 := GCons(__e, symdefine, tmp18618)

tmp18620 := GCons(__e, symdefun, tmp18619)

tmp18621 := GCons(__e, symdestroy, tmp18620)

tmp18622 := GCons(__e, symdifference, tmp18621)

tmp18623 := GCons(__e, symdo, tmp18622)

tmp18624 := GCons(__e, symelement_2, tmp18623)

tmp18625 := GCons(__e, symempty_2, tmp18624)

tmp18626 := GCons(__e, symerror, tmp18625)

tmp18627 := GCons(__e, symerror_1to_1string, tmp18626)

tmp18628 := GCons(__e, symeval, tmp18627)

tmp18629 := GCons(__e, symeval_1kl, tmp18628)

tmp18630 := GCons(__e, symexception, tmp18629)

tmp18631 := GCons(__e, symexternal, tmp18630)

tmp18632 := GCons(__e, symexplode, tmp18631)

tmp18633 := GCons(__e, symenable_1type_1theory, tmp18632)

tmp18634 := GCons(__e, False, tmp18633)

tmp18635 := GCons(__e, symfindall, tmp18634)

tmp18636 := GCons(__e, symfactorise_2, tmp18635)

tmp18637 := GCons(__e, symfactorise, tmp18636)

tmp18638 := GCons(__e, symfail_1if, tmp18637)

tmp18639 := GCons(__e, symfail, tmp18638)

tmp18640 := GCons(__e, symfile, tmp18639)

tmp18641 := GCons(__e, symfix, tmp18640)

tmp18642 := GCons(__e, symforeign, tmp18641)

tmp18643 := GCons(__e, symfork, tmp18642)

tmp18644 := GCons(__e, symfresh, tmp18643)

tmp18645 := GCons(__e, symfreeze, tmp18644)

tmp18646 := GCons(__e, symfst, tmp18645)

tmp18647 := GCons(__e, symfunction, tmp18646)

tmp18648 := GCons(__e, symfn, tmp18647)

tmp18649 := GCons(__e, symgensym, tmp18648)

tmp18650 := GCons(__e, symget_1time, tmp18649)

tmp18651 := GCons(__e, symget, tmp18650)

tmp18652 := GCons(__e, symhash, tmp18651)

tmp18653 := GCons(__e, symhdstr, tmp18652)

tmp18654 := GCons(__e, symhdv, tmp18653)

tmp18655 := GCons(__e, symhd, tmp18654)

tmp18656 := GCons(__e, symhead, tmp18655)

tmp18657 := GCons(__e, symhush, tmp18656)

tmp18658 := GCons(__e, symhush_2, tmp18657)

tmp18659 := GCons(__e, symif, tmp18658)

tmp18660 := GCons(__e, symimplementation, tmp18659)

tmp18661 := GCons(__e, syminternal, tmp18660)

tmp18662 := GCons(__e, symin_1package, tmp18661)

tmp18663 := GCons(__e, symin, tmp18662)

tmp18664 := GCons(__e, symis_b, tmp18663)

tmp18665 := GCons(__e, symis, tmp18664)

tmp18666 := GCons(__e, symit, tmp18665)

tmp18667 := GCons(__e, syminclude_1all_1but, tmp18666)

tmp18668 := GCons(__e, syminclude, tmp18667)

tmp18669 := GCons(__e, symincluded, tmp18668)

tmp18670 := GCons(__e, syminline, tmp18669)

tmp18671 := GCons(__e, syminput_7, tmp18670)

tmp18672 := GCons(__e, syminput, tmp18671)

tmp18673 := GCons(__e, syminteger_2, tmp18672)

tmp18674 := GCons(__e, symintern, tmp18673)

tmp18675 := GCons(__e, syminferences, tmp18674)

tmp18676 := GCons(__e, symintersection, tmp18675)

tmp18677 := GCons(__e, symis, tmp18676)

tmp18678 := GCons(__e, symlanguage, tmp18677)

tmp18679 := GCons(__e, symlambda, tmp18678)

tmp18680 := GCons(__e, symlazy, tmp18679)

tmp18681 := GCons(__e, symlet, tmp18680)

tmp18682 := GCons(__e, symlength, tmp18681)

tmp18683 := GCons(__e, symlimit, tmp18682)

tmp18684 := GCons(__e, symlineread, tmp18683)

tmp18685 := GCons(__e, symlist, tmp18684)

tmp18686 := GCons(__e, symloaded, tmp18685)

tmp18687 := GCons(__e, symload, tmp18686)

tmp18688 := GCons(__e, symmake_1string, tmp18687)

tmp18689 := GCons(__e, symmap, tmp18688)

tmp18690 := GCons(__e, symmapcan, tmp18689)

tmp18691 := GCons(__e, symmaxinferences, tmp18690)

tmp18692 := GCons(__e, symmacroexpand, tmp18691)

tmp18693 := GCons(__e, symmode, tmp18692)

tmp18694 := GCons(__e, symnl, tmp18693)

tmp18695 := GCons(__e, symnot, tmp18694)

tmp18696 := GCons(__e, symnth, tmp18695)

tmp18697 := GCons(__e, symnull, tmp18696)

tmp18698 := GCons(__e, symnumber, tmp18697)

tmp18699 := GCons(__e, symnumber_2, tmp18698)

tmp18700 := GCons(__e, symn_1_6string, tmp18699)

tmp18701 := GCons(__e, symoccurs_2, tmp18700)

tmp18702 := GCons(__e, symoccurs_1check, tmp18701)

tmp18703 := GCons(__e, symoccurrences, tmp18702)

tmp18704 := GCons(__e, symopen, tmp18703)

tmp18705 := GCons(__e, symoptimise_2, tmp18704)

tmp18706 := GCons(__e, symoptimise, tmp18705)

tmp18707 := GCons(__e, symor, tmp18706)

tmp18708 := GCons(__e, symos, tmp18707)

tmp18709 := GCons(__e, symout, tmp18708)

tmp18710 := GCons(__e, symoutput, tmp18709)

tmp18711 := GCons(__e, sympackage, tmp18710)

tmp18712 := GCons(__e, symport, tmp18711)

tmp18713 := GCons(__e, symporters, tmp18712)

tmp18714 := GCons(__e, sympos, tmp18713)

tmp18715 := GCons(__e, sympr, tmp18714)

tmp18716 := GCons(__e, symprint, tmp18715)

tmp18717 := GCons(__e, symprolog_1memory, tmp18716)

tmp18718 := GCons(__e, symprofile, tmp18717)

tmp18719 := GCons(__e, symprofile_1results, tmp18718)

tmp18720 := GCons(__e, symprotect, tmp18719)

tmp18721 := GCons(__e, symprolog_2, tmp18720)

tmp18722 := GCons(__e, symps, tmp18721)

tmp18723 := GCons(__e, sympreclude_1all_1but, tmp18722)

tmp18724 := GCons(__e, sympreclude, tmp18723)

tmp18725 := GCons(__e, symput, tmp18724)

tmp18726 := GCons(__e, sympackage_2, tmp18725)

tmp18727 := GCons(__e, symread_1unit_1string, tmp18726)

tmp18728 := GCons(__e, symread_1from_1string_1unprocessed, tmp18727)

tmp18729 := GCons(__e, symread_1from_1string, tmp18728)

tmp18730 := GCons(__e, symread_1byte, tmp18729)

tmp18731 := GCons(__e, symread_1file_1as_1string, tmp18730)

tmp18732 := GCons(__e, symread_1file_1as_1bytelist, tmp18731)

tmp18733 := GCons(__e, symread_1file, tmp18732)

tmp18734 := GCons(__e, symreceive, tmp18733)

tmp18735 := GCons(__e, symread, tmp18734)

tmp18736 := GCons(__e, symrelease, tmp18735)

tmp18737 := GCons(__e, symremove, tmp18736)

tmp18738 := GCons(__e, symretract, tmp18737)

tmp18739 := GCons(__e, symreverse, tmp18738)

tmp18740 := GCons(__e, symrun, tmp18739)

tmp18741 := GCons(__e, symstr, tmp18740)

tmp18742 := GCons(__e, symsave, tmp18741)

tmp18743 := GCons(__e, symset, tmp18742)

tmp18744 := GCons(__e, symsimple_1error, tmp18743)

tmp18745 := GCons(__e, symsnd, tmp18744)

tmp18746 := GCons(__e, symspecialise, tmp18745)

tmp18747 := GCons(__e, symspy, tmp18746)

tmp18748 := GCons(__e, symsqts, tmp18747)

tmp18749 := GCons(__e, symstep, tmp18748)

tmp18750 := GCons(__e, symstoutput, tmp18749)

tmp18751 := GCons(__e, symsterror, tmp18750)

tmp18752 := GCons(__e, symstinput, tmp18751)

tmp18753 := GCons(__e, symstring, tmp18752)

tmp18754 := GCons(__e, symstream, tmp18753)

tmp18755 := GCons(__e, symstring_1_6n, tmp18754)

tmp18756 := GCons(__e, symstring_2, tmp18755)

tmp18757 := GCons(__e, symsubst, tmp18756)

tmp18758 := GCons(__e, symsum, tmp18757)

tmp18759 := GCons(__e, symstring_1_6symbol, tmp18758)

tmp18760 := GCons(__e, symsymbol_2, tmp18759)

tmp18761 := GCons(__e, symsymbol, tmp18760)

tmp18762 := GCons(__e, symsystem_1S_2, tmp18761)

tmp18763 := GCons(__e, symsynonyms, tmp18762)

tmp18764 := GCons(__e, symsystemf, tmp18763)

tmp18765 := GCons(__e, symtail, tmp18764)

tmp18766 := GCons(__e, symtlv, tmp18765)

tmp18767 := GCons(__e, symtlstr, tmp18766)

tmp18768 := GCons(__e, symtl, tmp18767)

tmp18769 := GCons(__e, symtc, tmp18768)

tmp18770 := GCons(__e, symtc_2, tmp18769)

tmp18771 := GCons(__e, symthaw, tmp18770)

tmp18772 := GCons(__e, symtime, tmp18771)

tmp18773 := GCons(__e, symtrack, tmp18772)

tmp18774 := GCons(__e, symtracked, tmp18773)

tmp18775 := GCons(__e, symtrap_1error, tmp18774)

tmp18776 := GCons(__e, True, tmp18775)

tmp18777 := GCons(__e, symtuple_2, tmp18776)

tmp18778 := GCons(__e, symtype, tmp18777)

tmp18779 := GCons(__e, symreturn, tmp18778)

tmp18780 := GCons(__e, symunabsolute, tmp18779)

tmp18781 := GCons(__e, symundefmacro, tmp18780)

tmp18782 := GCons(__e, symunprofile, tmp18781)

tmp18783 := GCons(__e, symunput, tmp18782)

tmp18784 := GCons(__e, symunion, tmp18783)

tmp18785 := GCons(__e, symunix, tmp18784)

tmp18786 := GCons(__e, symunit, tmp18785)

tmp18787 := GCons(__e, symuntrack, tmp18786)

tmp18788 := GCons(__e, symunspecialise, tmp18787)

tmp18789 := GCons(__e, symupdate_1lambda_1table, tmp18788)

tmp18790 := GCons(__e, symu_b, tmp18789)

tmp18791 := GCons(__e, symuserdefs, tmp18790)

tmp18792 := GCons(__e, symvector_2, tmp18791)

tmp18793 := GCons(__e, symvector, tmp18792)

tmp18794 := GCons(__e, sym_5_1vector, tmp18793)

tmp18795 := GCons(__e, symvector_1_6, tmp18794)

tmp18796 := GCons(__e, symvalue, tmp18795)

tmp18797 := GCons(__e, symvar_2, tmp18796)

tmp18798 := GCons(__e, symvariable_2, tmp18797)

tmp18799 := GCons(__e, symverified, tmp18798)

tmp18800 := GCons(__e, symversion, tmp18799)

tmp18801 := GCons(__e, symwarn, tmp18800)

tmp18802 := GCons(__e, symwhen, tmp18801)

tmp18803 := GCons(__e, symwhere, tmp18802)

tmp18804 := GCons(__e, symwrite_1byte, tmp18803)

tmp18805 := GCons(__e, symwrite_1to_1file, tmp18804)

tmp18806 := GCons(__e, symwrite_1string, tmp18805)

tmp18807 := GCons(__e, symy_1or_1n_2, tmp18806)

tmp18808 := GCons(__e, tmp18577, tmp18807)

tmp18809 := GCons(__e, sym_6_6, tmp18808)

tmp18810 := GCons(__e, sym_5, tmp18809)

tmp18811 := GCons(__e, sym_5_a, tmp18810)

tmp18812 := GCons(__e, sym_7, tmp18811)

tmp18813 := GCons(__e, sym_d, tmp18812)

tmp18814 := GCons(__e, sym_c, tmp18813)

tmp18815 := GCons(__e, sym_1, tmp18814)

tmp18816 := GCons(__e, sym_3, tmp18815)

tmp18817 := GCons(__e, sym_5end_6, tmp18816)

tmp18818 := GCons(__e, sym_5_b_6, tmp18817)

tmp18819 := GCons(__e, sym_c_4, tmp18818)

tmp18820 := GCons(__e, sym_a_a_6, tmp18819)

tmp18821 := GCons(__e, sym_6, tmp18820)

tmp18822 := GCons(__e, sym_6_a, tmp18821)

tmp18823 := GCons(__e, sym_a, tmp18822)

tmp18824 := GCons(__e, sym_a_a, tmp18823)

tmp18825 := GCons(__e, sym_5e_6, tmp18824)

tmp18826 := GCons(__e, sym_1_6, tmp18825)

tmp18827 := GCons(__e, sym_5_1, tmp18826)

tmp18828 := GCons(__e, sym_dhush_d, tmp18827)

tmp18829 := GCons(__e, sym_dporters_d, tmp18828)

tmp18830 := GCons(__e, sym_dport_d, tmp18829)

tmp18831 := GCons(__e, sym_8s, tmp18830)

tmp18832 := GCons(__e, sym_8p, tmp18831)

tmp18833 := GCons(__e, sym_8v, tmp18832)

tmp18834 := GCons(__e, sym_dproperty_1vector_d, tmp18833)

tmp18835 := GCons(__e, sym_drelease_d, tmp18834)

tmp18836 := GCons(__e, sym_dos_d, tmp18835)

tmp18837 := GCons(__e, sym_dmacros_d, tmp18836)

tmp18838 := GCons(__e, sym_dmaximum_1print_1sequence_1size_d, tmp18837)

tmp18839 := GCons(__e, sym_dversion_d, tmp18838)

tmp18840 := GCons(__e, sym_dhome_1directory_d, tmp18839)

tmp18841 := GCons(__e, sym_dstoutput_d, tmp18840)

tmp18842 := GCons(__e, sym_dsterror_d, tmp18841)

tmp18843 := GCons(__e, sym_dstinput_d, tmp18842)

tmp18844 := GCons(__e, sym_dimplementation_d, tmp18843)

tmp18845 := GCons(__e, sym_dlanguage_d, tmp18844)

tmp18846 := GCons(__e, sym__, tmp18845)

tmp18847 := GCons(__e, tmp18576, tmp18846)

tmp18848 := GCons(__e, tmp18575, tmp18847)

tmp18849 := GCons(__e, tmp18574, tmp18848)

tmp18850 := GCons(__e, tmp18573, tmp18849)

tmp18851 := GCons(__e, sym_e_e, tmp18850)

tmp18852 := GCons(__e, sym_5_1_1, tmp18851)

tmp18853 := GCons(__e, sym_1_1_6, tmp18852)

tmp18854 := GCons(__e, sym_i, tmp18853)

tmp18855 := GCons(__e, sym_j, tmp18854)

tmp18856 := GCons(__e, sym_b, tmp18855)

tmp18857 := GValue(__e, sym_dproperty_1vector_d)

tmp18858 := Call4(__e, SymFn(symput), symshen, symshen_4external_1symbols, tmp18856, tmp18857)


_ = tmp18858

tmp18859 := PrimGuard1(__e, symabsvector, PrimAbsvector, MakeInteger(0))

return PrimGuard2(__e, symset, PrimSet, symshen_4_dempty_1absvector_d, tmp18859)


panic("unreachable")
})

tmp18860 := Call2(__e, ns2_1set, symshen_4initialise_1environment, tmp18096)


_ = tmp18860

tmp18861 := MakeNative0(func(__e *ControlFlow) Obj {
tmp18862 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, Nil)

_ = tmp18862

tmp18863 := MakeNative1(func(__e *ControlFlow, V6094 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6090 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6091 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6092 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6093 Obj) Obj {
tmp18864 := Call1(__e, SymFn(symshen_4newpv), B6090)


let__18009 := tmp18864
_ = let__18009

tmp18865 := Call1(__e, SymFn(symshen_4newpv), B6090)


let__18010 := tmp18865
_ = let__18010

tmp18866 := GCons(__e, let__18010, Nil)

tmp18867 := GCons(__e, symlist, tmp18866)

tmp18868 := GCons(__e, tmp18867, Nil)

tmp18869 := GCons(__e, sym_1_1_6, tmp18868)

tmp18870 := GCons(__e, let__18009, tmp18869)

tmp18871 := GCons(__e, symsymbol, Nil)

tmp18872 := GCons(__e, sym_1_1_6, tmp18871)

tmp18873 := GCons(__e, tmp18870, tmp18872)

tmp18874 := GCons(__e, tmp18873, Nil)

tmp18875 := GCons(__e, sym_1_1_6, tmp18874)

tmp18876 := GCons(__e, symsymbol, tmp18875)

tmp18877 := Call6(__e, SymFn(symis_b), V6094, tmp18876, B6090, L6091, Key6092, C6093)


tmp18878 := Call2(__e, SymFn(symshen_4gc), B6090, tmp18877)


return TailCall2(__e, SymFn(symshen_4gc), B6090, tmp18878)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp18879 := GValue(__e, symshen_4_dsigf_d)

tmp18880 := Call3(__e, SymFn(symshen_4assoc_1_6), symshen_4register_1source_1form, tmp18863, tmp18879)


tmp18881 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp18880)

_ = tmp18881

tmp18882 := MakeNative1(func(__e *ControlFlow, V6099 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6095 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6096 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6097 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6098 Obj) Obj {
tmp18883 := GCons(__e, symsymbol, Nil)

tmp18884 := GCons(__e, sym_1_1_6, tmp18883)

tmp18885 := GCons(__e, symsymbol, tmp18884)

return TailCall6(__e, SymFn(symis_b), V6099, tmp18885, B6095, L6096, Key6097, C6098)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp18886 := GValue(__e, symshen_4_dsigf_d)

tmp18887 := Call3(__e, SymFn(symshen_4assoc_1_6), symshen_4unregister_1source_1form, tmp18882, tmp18886)


tmp18888 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp18887)

_ = tmp18888

tmp18889 := MakeNative1(func(__e *ControlFlow, V6104 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6100 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6101 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6102 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6103 Obj) Obj {
tmp18890 := Call1(__e, SymFn(symshen_4newpv), B6100)


let__18011 := tmp18890
_ = let__18011

tmp18891 := GCons(__e, let__18011, Nil)

tmp18892 := GCons(__e, sym_1_1_6, tmp18891)

tmp18893 := Call6(__e, SymFn(symis_b), V6104, tmp18892, B6100, L6101, Key6102, C6103)


return TailCall2(__e, SymFn(symshen_4gc), B6100, tmp18893)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp18894 := GValue(__e, symshen_4_dsigf_d)

tmp18895 := Call3(__e, SymFn(symshen_4assoc_1_6), symabort, tmp18889, tmp18894)


tmp18896 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp18895)

_ = tmp18896

tmp18897 := MakeNative1(func(__e *ControlFlow, V6109 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6105 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6106 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6107 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6108 Obj) Obj {
tmp18898 := GCons(__e, symstring, Nil)

tmp18899 := GCons(__e, symlist, tmp18898)

tmp18900 := GCons(__e, tmp18899, Nil)

tmp18901 := GCons(__e, sym_1_1_6, tmp18900)

tmp18902 := GCons(__e, symstring, tmp18901)

return TailCall6(__e, SymFn(symis_b), V6109, tmp18902, B6105, L6106, Key6107, C6108)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp18903 := GValue(__e, symshen_4_dsigf_d)

tmp18904 := Call3(__e, SymFn(symshen_4assoc_1_6), symabsolute, tmp18897, tmp18903)


tmp18905 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp18904)

_ = tmp18905

tmp18906 := MakeNative1(func(__e *ControlFlow, V6114 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6110 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6111 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6112 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6113 Obj) Obj {
tmp18907 := Call1(__e, SymFn(symshen_4newpv), B6110)


let__18012 := tmp18907
_ = let__18012

tmp18908 := GCons(__e, symboolean, Nil)

tmp18909 := GCons(__e, sym_1_1_6, tmp18908)

tmp18910 := GCons(__e, let__18012, tmp18909)

tmp18911 := Call6(__e, SymFn(symis_b), V6114, tmp18910, B6110, L6111, Key6112, C6113)


return TailCall2(__e, SymFn(symshen_4gc), B6110, tmp18911)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp18912 := GValue(__e, symshen_4_dsigf_d)

tmp18913 := Call3(__e, SymFn(symshen_4assoc_1_6), symabsvector_2, tmp18906, tmp18912)


tmp18914 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp18913)

_ = tmp18914

tmp18915 := MakeNative1(func(__e *ControlFlow, V6120 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6116 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6117 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6118 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6119 Obj) Obj {
tmp18916 := Call1(__e, SymFn(symshen_4newpv), B6116)


let__18013 := tmp18916
_ = let__18013

tmp18917 := GCons(__e, let__18013, Nil)

tmp18918 := GCons(__e, symlist, tmp18917)

tmp18919 := GCons(__e, let__18013, Nil)

tmp18920 := GCons(__e, symlist, tmp18919)

tmp18921 := GCons(__e, tmp18920, Nil)

tmp18922 := GCons(__e, sym_1_1_6, tmp18921)

tmp18923 := GCons(__e, tmp18918, tmp18922)

tmp18924 := GCons(__e, tmp18923, Nil)

tmp18925 := GCons(__e, sym_1_1_6, tmp18924)

tmp18926 := GCons(__e, let__18013, tmp18925)

tmp18927 := Call6(__e, SymFn(symis_b), V6120, tmp18926, B6116, L6117, Key6118, C6119)


return TailCall2(__e, SymFn(symshen_4gc), B6116, tmp18927)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp18928 := GValue(__e, symshen_4_dsigf_d)

tmp18929 := Call3(__e, SymFn(symshen_4assoc_1_6), symadjoin, tmp18915, tmp18928)


tmp18930 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp18929)

_ = tmp18930

tmp18931 := MakeNative1(func(__e *ControlFlow, V6125 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6121 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6122 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6123 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6124 Obj) Obj {
tmp18932 := GCons(__e, symboolean, Nil)

tmp18933 := GCons(__e, sym_1_1_6, tmp18932)

tmp18934 := GCons(__e, symboolean, tmp18933)

tmp18935 := GCons(__e, tmp18934, Nil)

tmp18936 := GCons(__e, sym_1_1_6, tmp18935)

tmp18937 := GCons(__e, symboolean, tmp18936)

return TailCall6(__e, SymFn(symis_b), V6125, tmp18937, B6121, L6122, Key6123, C6124)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp18938 := GValue(__e, symshen_4_dsigf_d)

tmp18939 := Call3(__e, SymFn(symshen_4assoc_1_6), symand, tmp18931, tmp18938)


tmp18940 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp18939)

_ = tmp18940

tmp18941 := MakeNative1(func(__e *ControlFlow, V6130 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6126 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6127 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6128 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6129 Obj) Obj {
tmp18942 := Call1(__e, SymFn(symshen_4newpv), B6126)


let__18014 := tmp18942
_ = let__18014

tmp18943 := GCons(__e, symstring, Nil)

tmp18944 := GCons(__e, sym_1_1_6, tmp18943)

tmp18945 := GCons(__e, symsymbol, tmp18944)

tmp18946 := GCons(__e, tmp18945, Nil)

tmp18947 := GCons(__e, sym_1_1_6, tmp18946)

tmp18948 := GCons(__e, symstring, tmp18947)

tmp18949 := GCons(__e, tmp18948, Nil)

tmp18950 := GCons(__e, sym_1_1_6, tmp18949)

tmp18951 := GCons(__e, let__18014, tmp18950)

tmp18952 := Call6(__e, SymFn(symis_b), V6130, tmp18951, B6126, L6127, Key6128, C6129)


return TailCall2(__e, SymFn(symshen_4gc), B6126, tmp18952)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp18953 := GValue(__e, symshen_4_dsigf_d)

tmp18954 := Call3(__e, SymFn(symshen_4assoc_1_6), symshen_4app, tmp18941, tmp18953)


tmp18955 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp18954)

_ = tmp18955

tmp18956 := MakeNative1(func(__e *ControlFlow, V6137 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6133 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6134 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6135 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6136 Obj) Obj {
tmp18957 := Call1(__e, SymFn(symshen_4newpv), B6133)


let__18015 := tmp18957
_ = let__18015

tmp18958 := GCons(__e, let__18015, Nil)

tmp18959 := GCons(__e, symlist, tmp18958)

tmp18960 := GCons(__e, let__18015, Nil)

tmp18961 := GCons(__e, symlist, tmp18960)

tmp18962 := GCons(__e, let__18015, Nil)

tmp18963 := GCons(__e, symlist, tmp18962)

tmp18964 := GCons(__e, tmp18963, Nil)

tmp18965 := GCons(__e, sym_1_1_6, tmp18964)

tmp18966 := GCons(__e, tmp18961, tmp18965)

tmp18967 := GCons(__e, tmp18966, Nil)

tmp18968 := GCons(__e, sym_1_1_6, tmp18967)

tmp18969 := GCons(__e, tmp18959, tmp18968)

tmp18970 := Call6(__e, SymFn(symis_b), V6137, tmp18969, B6133, L6134, Key6135, C6136)


return TailCall2(__e, SymFn(symshen_4gc), B6133, tmp18970)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp18971 := GValue(__e, symshen_4_dsigf_d)

tmp18972 := Call3(__e, SymFn(symshen_4assoc_1_6), symappend, tmp18956, tmp18971)


tmp18973 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp18972)

_ = tmp18973

tmp18974 := MakeNative1(func(__e *ControlFlow, V6142 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6138 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6139 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6140 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6141 Obj) Obj {
tmp18975 := Call1(__e, SymFn(symshen_4newpv), B6138)


let__18016 := tmp18975
_ = let__18016

tmp18976 := GCons(__e, symnumber, Nil)

tmp18977 := GCons(__e, sym_1_1_6, tmp18976)

tmp18978 := GCons(__e, let__18016, tmp18977)

tmp18979 := Call6(__e, SymFn(symis_b), V6142, tmp18978, B6138, L6139, Key6140, C6141)


return TailCall2(__e, SymFn(symshen_4gc), B6138, tmp18979)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp18980 := GValue(__e, symshen_4_dsigf_d)

tmp18981 := Call3(__e, SymFn(symshen_4assoc_1_6), symarity, tmp18974, tmp18980)


tmp18982 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp18981)

_ = tmp18982

tmp18983 := MakeNative1(func(__e *ControlFlow, V6148 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6144 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6145 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6146 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6147 Obj) Obj {
tmp18984 := Call1(__e, SymFn(symshen_4newpv), B6144)


let__18017 := tmp18984
_ = let__18017

tmp18985 := GCons(__e, let__18017, Nil)

tmp18986 := GCons(__e, symlist, tmp18985)

tmp18987 := GCons(__e, tmp18986, Nil)

tmp18988 := GCons(__e, symlist, tmp18987)

tmp18989 := GCons(__e, let__18017, Nil)

tmp18990 := GCons(__e, symlist, tmp18989)

tmp18991 := GCons(__e, tmp18990, Nil)

tmp18992 := GCons(__e, sym_1_1_6, tmp18991)

tmp18993 := GCons(__e, tmp18988, tmp18992)

tmp18994 := GCons(__e, tmp18993, Nil)

tmp18995 := GCons(__e, sym_1_1_6, tmp18994)

tmp18996 := GCons(__e, let__18017, tmp18995)

tmp18997 := Call6(__e, SymFn(symis_b), V6148, tmp18996, B6144, L6145, Key6146, C6147)


return TailCall2(__e, SymFn(symshen_4gc), B6144, tmp18997)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp18998 := GValue(__e, symshen_4_dsigf_d)

tmp18999 := Call3(__e, SymFn(symshen_4assoc_1_6), symassoc, tmp18983, tmp18998)


tmp19000 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp18999)

_ = tmp19000

tmp19001 := MakeNative1(func(__e *ControlFlow, V6153 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6149 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6150 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6151 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6152 Obj) Obj {
tmp19002 := Call1(__e, SymFn(symshen_4newpv), B6149)


let__18018 := tmp19002
_ = let__18018

tmp19003 := GCons(__e, symboolean, Nil)

tmp19004 := GCons(__e, sym_1_1_6, tmp19003)

tmp19005 := GCons(__e, let__18018, tmp19004)

tmp19006 := Call6(__e, SymFn(symis_b), V6153, tmp19005, B6149, L6150, Key6151, C6152)


return TailCall2(__e, SymFn(symshen_4gc), B6149, tmp19006)


panic("unreachable")
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

tmp19008 := Call3(__e, SymFn(symshen_4assoc_1_6), symatom_2, tmp19001, tmp19007)


tmp19009 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19008)

_ = tmp19009

tmp19010 := MakeNative1(func(__e *ControlFlow, V6158 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6154 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6155 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6156 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6157 Obj) Obj {
tmp19011 := Call1(__e, SymFn(symshen_4newpv), B6154)


let__18019 := tmp19011
_ = let__18019

tmp19012 := GCons(__e, symboolean, Nil)

tmp19013 := GCons(__e, sym_1_1_6, tmp19012)

tmp19014 := GCons(__e, let__18019, tmp19013)

tmp19015 := Call6(__e, SymFn(symis_b), V6158, tmp19014, B6154, L6155, Key6156, C6157)


return TailCall2(__e, SymFn(symshen_4gc), B6154, tmp19015)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19016 := GValue(__e, symshen_4_dsigf_d)

tmp19017 := Call3(__e, SymFn(symshen_4assoc_1_6), symboolean_2, tmp19010, tmp19016)


tmp19018 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19017)

_ = tmp19018

tmp19019 := MakeNative1(func(__e *ControlFlow, V6163 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6159 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6160 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6161 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6162 Obj) Obj {
tmp19020 := GCons(__e, symstring, Nil)

tmp19021 := GCons(__e, sym_1_1_6, tmp19020)

tmp19022 := GCons(__e, symstring, tmp19021)

return TailCall6(__e, SymFn(symis_b), V6163, tmp19022, B6159, L6160, Key6161, C6162)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19023 := GValue(__e, symshen_4_dsigf_d)

tmp19024 := Call3(__e, SymFn(symshen_4assoc_1_6), symbootstrap, tmp19019, tmp19023)


tmp19025 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19024)

_ = tmp19025

tmp19026 := MakeNative1(func(__e *ControlFlow, V6168 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6164 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6165 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6166 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6167 Obj) Obj {
tmp19027 := GCons(__e, symboolean, Nil)

tmp19028 := GCons(__e, sym_1_1_6, tmp19027)

tmp19029 := GCons(__e, symsymbol, tmp19028)

return TailCall6(__e, SymFn(symis_b), V6168, tmp19029, B6164, L6165, Key6166, C6167)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19030 := GValue(__e, symshen_4_dsigf_d)

tmp19031 := Call3(__e, SymFn(symshen_4assoc_1_6), symbound_2, tmp19026, tmp19030)


tmp19032 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19031)

_ = tmp19032

tmp19033 := MakeNative1(func(__e *ControlFlow, V6173 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6169 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6170 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6171 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6172 Obj) Obj {
tmp19034 := Call1(__e, SymFn(symshen_4newpv), B6169)


let__18020 := tmp19034
_ = let__18020

tmp19035 := GCons(__e, let__18020, Nil)

tmp19036 := GCons(__e, symlist, tmp19035)

tmp19037 := GCons(__e, symboolean, Nil)

tmp19038 := GCons(__e, sym_1_1_6, tmp19037)

tmp19039 := GCons(__e, tmp19036, tmp19038)

tmp19040 := Call6(__e, SymFn(symis_b), V6173, tmp19039, B6169, L6170, Key6171, C6172)


return TailCall2(__e, SymFn(symshen_4gc), B6169, tmp19040)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19041 := GValue(__e, symshen_4_dsigf_d)

tmp19042 := Call3(__e, SymFn(symshen_4assoc_1_6), symshen_4ccons_2, tmp19033, tmp19041)


tmp19043 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19042)

_ = tmp19043

tmp19044 := MakeNative1(func(__e *ControlFlow, V6178 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6174 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6175 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6176 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6177 Obj) Obj {
tmp19045 := GCons(__e, symstring, Nil)

tmp19046 := GCons(__e, sym_1_1_6, tmp19045)

tmp19047 := GCons(__e, symstring, tmp19046)

return TailCall6(__e, SymFn(symis_b), V6178, tmp19047, B6174, L6175, Key6176, C6177)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19048 := GValue(__e, symshen_4_dsigf_d)

tmp19049 := Call3(__e, SymFn(symshen_4assoc_1_6), symcd, tmp19044, tmp19048)


tmp19050 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19049)

_ = tmp19050

tmp19051 := MakeNative1(func(__e *ControlFlow, V6184 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6180 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6181 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6182 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6183 Obj) Obj {
tmp19052 := Call1(__e, SymFn(symshen_4newpv), B6180)


let__18021 := tmp19052
_ = let__18021

tmp19053 := Call1(__e, SymFn(symshen_4newpv), B6180)


let__18022 := tmp19053
_ = let__18022

tmp19054 := GCons(__e, let__18021, Nil)

tmp19055 := GCons(__e, symstream, tmp19054)

tmp19056 := GCons(__e, let__18022, Nil)

tmp19057 := GCons(__e, symlist, tmp19056)

tmp19058 := GCons(__e, tmp19057, Nil)

tmp19059 := GCons(__e, sym_1_1_6, tmp19058)

tmp19060 := GCons(__e, tmp19055, tmp19059)

tmp19061 := Call6(__e, SymFn(symis_b), V6184, tmp19060, B6180, L6181, Key6182, C6183)


tmp19062 := Call2(__e, SymFn(symshen_4gc), B6180, tmp19061)


return TailCall2(__e, SymFn(symshen_4gc), B6180, tmp19062)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19063 := GValue(__e, symshen_4_dsigf_d)

tmp19064 := Call3(__e, SymFn(symshen_4assoc_1_6), symclose, tmp19051, tmp19063)


tmp19065 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19064)

_ = tmp19065

tmp19066 := MakeNative1(func(__e *ControlFlow, V6189 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6185 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6186 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6187 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6188 Obj) Obj {
tmp19067 := GCons(__e, symstring, Nil)

tmp19068 := GCons(__e, sym_1_1_6, tmp19067)

tmp19069 := GCons(__e, symstring, tmp19068)

tmp19070 := GCons(__e, tmp19069, Nil)

tmp19071 := GCons(__e, sym_1_1_6, tmp19070)

tmp19072 := GCons(__e, symstring, tmp19071)

return TailCall6(__e, SymFn(symis_b), V6189, tmp19072, B6185, L6186, Key6187, C6188)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19073 := GValue(__e, symshen_4_dsigf_d)

tmp19074 := Call3(__e, SymFn(symshen_4assoc_1_6), symcn, tmp19066, tmp19073)


tmp19075 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19074)

_ = tmp19075

tmp19076 := MakeNative1(func(__e *ControlFlow, V6194 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6190 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6191 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6192 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6193 Obj) Obj {
tmp19077 := Call1(__e, SymFn(symshen_4newpv), B6190)


let__18023 := tmp19077
_ = let__18023

tmp19078 := GCons(__e, let__18023, Nil)

tmp19079 := GCons(__e, symstream, tmp19078)

tmp19080 := GCons(__e, symboolean, Nil)

tmp19081 := GCons(__e, sym_1_1_6, tmp19080)

tmp19082 := GCons(__e, tmp19079, tmp19081)

tmp19083 := Call6(__e, SymFn(symis_b), V6194, tmp19082, B6190, L6191, Key6192, C6193)


return TailCall2(__e, SymFn(symshen_4gc), B6190, tmp19083)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19084 := GValue(__e, symshen_4_dsigf_d)

tmp19085 := Call3(__e, SymFn(symshen_4assoc_1_6), symchar_1stream_2, tmp19076, tmp19084)


tmp19086 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19085)

_ = tmp19086

tmp19087 := MakeNative1(func(__e *ControlFlow, V6201 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6197 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6198 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6199 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6200 Obj) Obj {
tmp19088 := Call1(__e, SymFn(symshen_4newpv), B6197)


let__18024 := tmp19088
_ = let__18024

tmp19089 := Call1(__e, SymFn(symshen_4newpv), B6197)


let__18025 := tmp19089
_ = let__18025

tmp19090 := GCons(__e, let__18024, Nil)

tmp19091 := GCons(__e, symlist, tmp19090)

tmp19092 := GCons(__e, let__18024, Nil)

tmp19093 := GCons(__e, symlist, tmp19092)

tmp19094 := GCons(__e, let__18025, Nil)

tmp19095 := GCons(__e, tmp19093, tmp19094)

tmp19096 := GCons(__e, symstr, tmp19095)

tmp19097 := GCons(__e, tmp19096, Nil)

tmp19098 := GCons(__e, sym_1_1_6, tmp19097)

tmp19099 := GCons(__e, tmp19091, tmp19098)

tmp19100 := GCons(__e, let__18024, Nil)

tmp19101 := GCons(__e, symlist, tmp19100)

tmp19102 := GCons(__e, let__18025, Nil)

tmp19103 := GCons(__e, sym_1_1_6, tmp19102)

tmp19104 := GCons(__e, tmp19101, tmp19103)

tmp19105 := GCons(__e, tmp19104, Nil)

tmp19106 := GCons(__e, sym_1_1_6, tmp19105)

tmp19107 := GCons(__e, tmp19099, tmp19106)

tmp19108 := Call6(__e, SymFn(symis_b), V6201, tmp19107, B6197, L6198, Key6199, C6200)


tmp19109 := Call2(__e, SymFn(symshen_4gc), B6197, tmp19108)


return TailCall2(__e, SymFn(symshen_4gc), B6197, tmp19109)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19110 := GValue(__e, symshen_4_dsigf_d)

tmp19111 := Call3(__e, SymFn(symshen_4assoc_1_6), symcompile, tmp19087, tmp19110)


tmp19112 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19111)

_ = tmp19112

tmp19113 := MakeNative1(func(__e *ControlFlow, V6206 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6202 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6203 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6204 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6205 Obj) Obj {
tmp19114 := Call1(__e, SymFn(symshen_4newpv), B6202)


let__18026 := tmp19114
_ = let__18026

tmp19115 := GCons(__e, symboolean, Nil)

tmp19116 := GCons(__e, sym_1_1_6, tmp19115)

tmp19117 := GCons(__e, let__18026, tmp19116)

tmp19118 := Call6(__e, SymFn(symis_b), V6206, tmp19117, B6202, L6203, Key6204, C6205)


return TailCall2(__e, SymFn(symshen_4gc), B6202, tmp19118)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19119 := GValue(__e, symshen_4_dsigf_d)

tmp19120 := Call3(__e, SymFn(symshen_4assoc_1_6), symcons_2, tmp19113, tmp19119)


tmp19121 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19120)

_ = tmp19121

tmp19122 := MakeNative1(func(__e *ControlFlow, V6211 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6207 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6208 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6209 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6210 Obj) Obj {
tmp19123 := GCons(__e, symsymbol, Nil)

tmp19124 := GCons(__e, symlist, tmp19123)

tmp19125 := GCons(__e, tmp19124, Nil)

tmp19126 := GCons(__e, sym_1_1_6, tmp19125)

return TailCall6(__e, SymFn(symis_b), V6211, tmp19126, B6207, L6208, Key6209, C6210)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19127 := GValue(__e, symshen_4_dsigf_d)

tmp19128 := Call3(__e, SymFn(symshen_4assoc_1_6), symdatatypes, tmp19122, tmp19127)


tmp19129 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19128)

_ = tmp19129

tmp19130 := MakeNative1(func(__e *ControlFlow, V6216 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6212 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6213 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6214 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6215 Obj) Obj {
tmp19131 := GCons(__e, symsymbol, Nil)

tmp19132 := GCons(__e, sym_1_1_6, tmp19131)

tmp19133 := GCons(__e, symsymbol, tmp19132)

return TailCall6(__e, SymFn(symis_b), V6216, tmp19133, B6212, L6213, Key6214, C6215)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19134 := GValue(__e, symshen_4_dsigf_d)

tmp19135 := Call3(__e, SymFn(symshen_4assoc_1_6), symdestroy, tmp19130, tmp19134)


tmp19136 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19135)

_ = tmp19136

tmp19137 := MakeNative1(func(__e *ControlFlow, V6223 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6219 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6220 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6221 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6222 Obj) Obj {
tmp19138 := Call1(__e, SymFn(symshen_4newpv), B6219)


let__18027 := tmp19138
_ = let__18027

tmp19139 := GCons(__e, let__18027, Nil)

tmp19140 := GCons(__e, symlist, tmp19139)

tmp19141 := GCons(__e, let__18027, Nil)

tmp19142 := GCons(__e, symlist, tmp19141)

tmp19143 := GCons(__e, let__18027, Nil)

tmp19144 := GCons(__e, symlist, tmp19143)

tmp19145 := GCons(__e, tmp19144, Nil)

tmp19146 := GCons(__e, sym_1_1_6, tmp19145)

tmp19147 := GCons(__e, tmp19142, tmp19146)

tmp19148 := GCons(__e, tmp19147, Nil)

tmp19149 := GCons(__e, sym_1_1_6, tmp19148)

tmp19150 := GCons(__e, tmp19140, tmp19149)

tmp19151 := Call6(__e, SymFn(symis_b), V6223, tmp19150, B6219, L6220, Key6221, C6222)


return TailCall2(__e, SymFn(symshen_4gc), B6219, tmp19151)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19152 := GValue(__e, symshen_4_dsigf_d)

tmp19153 := Call3(__e, SymFn(symshen_4assoc_1_6), symdifference, tmp19137, tmp19152)


tmp19154 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19153)

_ = tmp19154

tmp19155 := MakeNative1(func(__e *ControlFlow, V6228 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6224 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6225 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6226 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6227 Obj) Obj {
tmp19156 := Call1(__e, SymFn(symshen_4newpv), B6224)


let__18028 := tmp19156
_ = let__18028

tmp19157 := Call1(__e, SymFn(symshen_4newpv), B6224)


let__18029 := tmp19157
_ = let__18029

tmp19158 := GCons(__e, let__18029, Nil)

tmp19159 := GCons(__e, sym_1_1_6, tmp19158)

tmp19160 := GCons(__e, let__18029, tmp19159)

tmp19161 := GCons(__e, tmp19160, Nil)

tmp19162 := GCons(__e, sym_1_1_6, tmp19161)

tmp19163 := GCons(__e, let__18028, tmp19162)

tmp19164 := Call6(__e, SymFn(symis_b), V6228, tmp19163, B6224, L6225, Key6226, C6227)


tmp19165 := Call2(__e, SymFn(symshen_4gc), B6224, tmp19164)


return TailCall2(__e, SymFn(symshen_4gc), B6224, tmp19165)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19166 := GValue(__e, symshen_4_dsigf_d)

tmp19167 := Call3(__e, SymFn(symshen_4assoc_1_6), symdo, tmp19155, tmp19166)


tmp19168 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19167)

_ = tmp19168

tmp19169 := MakeNative1(func(__e *ControlFlow, V6235 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6231 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6232 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6233 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6234 Obj) Obj {
tmp19170 := Call1(__e, SymFn(symshen_4newpv), B6231)


let__18030 := tmp19170
_ = let__18030

tmp19171 := Call1(__e, SymFn(symshen_4newpv), B6231)


let__18031 := tmp19171
_ = let__18031

tmp19172 := GCons(__e, let__18030, Nil)

tmp19173 := GCons(__e, symlist, tmp19172)

tmp19174 := GCons(__e, let__18030, Nil)

tmp19175 := GCons(__e, symlist, tmp19174)

tmp19176 := GCons(__e, let__18031, Nil)

tmp19177 := GCons(__e, symlist, tmp19176)

tmp19178 := GCons(__e, tmp19177, Nil)

tmp19179 := GCons(__e, tmp19175, tmp19178)

tmp19180 := GCons(__e, symstr, tmp19179)

tmp19181 := GCons(__e, tmp19180, Nil)

tmp19182 := GCons(__e, sym_1_1_6, tmp19181)

tmp19183 := GCons(__e, tmp19173, tmp19182)

tmp19184 := Call6(__e, SymFn(symis_b), V6235, tmp19183, B6231, L6232, Key6233, C6234)


tmp19185 := Call2(__e, SymFn(symshen_4gc), B6231, tmp19184)


return TailCall2(__e, SymFn(symshen_4gc), B6231, tmp19185)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19186 := GValue(__e, symshen_4_dsigf_d)

tmp19187 := Call3(__e, SymFn(symshen_4assoc_1_6), sym_5e_6, tmp19169, tmp19186)


tmp19188 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19187)

_ = tmp19188

tmp19189 := MakeNative1(func(__e *ControlFlow, V6242 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6238 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6239 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6240 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6241 Obj) Obj {
tmp19190 := Call1(__e, SymFn(symshen_4newpv), B6238)


let__18032 := tmp19190
_ = let__18032

tmp19191 := Call1(__e, SymFn(symshen_4newpv), B6238)


let__18033 := tmp19191
_ = let__18033

tmp19192 := GCons(__e, let__18033, Nil)

tmp19193 := GCons(__e, symlist, tmp19192)

tmp19194 := GCons(__e, let__18032, Nil)

tmp19195 := GCons(__e, symlist, tmp19194)

tmp19196 := GCons(__e, let__18033, Nil)

tmp19197 := GCons(__e, symlist, tmp19196)

tmp19198 := GCons(__e, tmp19197, Nil)

tmp19199 := GCons(__e, tmp19195, tmp19198)

tmp19200 := GCons(__e, symstr, tmp19199)

tmp19201 := GCons(__e, tmp19200, Nil)

tmp19202 := GCons(__e, sym_1_1_6, tmp19201)

tmp19203 := GCons(__e, tmp19193, tmp19202)

tmp19204 := Call6(__e, SymFn(symis_b), V6242, tmp19203, B6238, L6239, Key6240, C6241)


tmp19205 := Call2(__e, SymFn(symshen_4gc), B6238, tmp19204)


return TailCall2(__e, SymFn(symshen_4gc), B6238, tmp19205)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19206 := GValue(__e, symshen_4_dsigf_d)

tmp19207 := Call3(__e, SymFn(symshen_4assoc_1_6), sym_5_b_6, tmp19189, tmp19206)


tmp19208 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19207)

_ = tmp19208

tmp19209 := MakeNative1(func(__e *ControlFlow, V6249 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6245 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6246 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6247 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6248 Obj) Obj {
tmp19210 := Call1(__e, SymFn(symshen_4newpv), B6245)


let__18034 := tmp19210
_ = let__18034

tmp19211 := Call1(__e, SymFn(symshen_4newpv), B6245)


let__18035 := tmp19211
_ = let__18035

tmp19212 := GCons(__e, let__18034, Nil)

tmp19213 := GCons(__e, symlist, tmp19212)

tmp19214 := GCons(__e, let__18034, Nil)

tmp19215 := GCons(__e, symlist, tmp19214)

tmp19216 := GCons(__e, let__18035, Nil)

tmp19217 := GCons(__e, symlist, tmp19216)

tmp19218 := GCons(__e, tmp19217, Nil)

tmp19219 := GCons(__e, tmp19215, tmp19218)

tmp19220 := GCons(__e, symstr, tmp19219)

tmp19221 := GCons(__e, tmp19220, Nil)

tmp19222 := GCons(__e, sym_1_1_6, tmp19221)

tmp19223 := GCons(__e, tmp19213, tmp19222)

tmp19224 := Call6(__e, SymFn(symis_b), V6249, tmp19223, B6245, L6246, Key6247, C6248)


tmp19225 := Call2(__e, SymFn(symshen_4gc), B6245, tmp19224)


return TailCall2(__e, SymFn(symshen_4gc), B6245, tmp19225)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19226 := GValue(__e, symshen_4_dsigf_d)

tmp19227 := Call3(__e, SymFn(symshen_4assoc_1_6), sym_5end_6, tmp19209, tmp19226)


tmp19228 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19227)

_ = tmp19228

tmp19229 := MakeNative1(func(__e *ControlFlow, V6254 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6250 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6251 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6252 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6253 Obj) Obj {
tmp19230 := Call1(__e, SymFn(symshen_4newpv), B6250)


let__18036 := tmp19230
_ = let__18036

tmp19231 := Call1(__e, SymFn(symshen_4newpv), B6250)


let__18037 := tmp19231
_ = let__18037

tmp19232 := GCons(__e, let__18036, Nil)

tmp19233 := GCons(__e, symlist, tmp19232)

tmp19234 := GCons(__e, let__18037, Nil)

tmp19235 := GCons(__e, tmp19233, tmp19234)

tmp19236 := GCons(__e, symstr, tmp19235)

tmp19237 := GCons(__e, symboolean, Nil)

tmp19238 := GCons(__e, sym_1_1_6, tmp19237)

tmp19239 := GCons(__e, tmp19236, tmp19238)

tmp19240 := Call6(__e, SymFn(symis_b), V6254, tmp19239, B6250, L6251, Key6252, C6253)


tmp19241 := Call2(__e, SymFn(symshen_4gc), B6250, tmp19240)


return TailCall2(__e, SymFn(symshen_4gc), B6250, tmp19241)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19242 := GValue(__e, symshen_4_dsigf_d)

tmp19243 := Call3(__e, SymFn(symshen_4assoc_1_6), symshen_4parse_1failure_2, tmp19229, tmp19242)


tmp19244 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19243)

_ = tmp19244

tmp19245 := MakeNative1(func(__e *ControlFlow, V6259 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6255 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6256 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6257 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6258 Obj) Obj {
tmp19246 := Call1(__e, SymFn(symshen_4newpv), B6255)


let__18038 := tmp19246
_ = let__18038

tmp19247 := Call1(__e, SymFn(symshen_4newpv), B6255)


let__18039 := tmp19247
_ = let__18039

tmp19248 := GCons(__e, let__18038, Nil)

tmp19249 := GCons(__e, symlist, tmp19248)

tmp19250 := GCons(__e, let__18039, Nil)

tmp19251 := GCons(__e, tmp19249, tmp19250)

tmp19252 := GCons(__e, symstr, tmp19251)

tmp19253 := GCons(__e, tmp19252, Nil)

tmp19254 := GCons(__e, sym_1_1_6, tmp19253)

tmp19255 := Call6(__e, SymFn(symis_b), V6259, tmp19254, B6255, L6256, Key6257, C6258)


tmp19256 := Call2(__e, SymFn(symshen_4gc), B6255, tmp19255)


return TailCall2(__e, SymFn(symshen_4gc), B6255, tmp19256)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19257 := GValue(__e, symshen_4_dsigf_d)

tmp19258 := Call3(__e, SymFn(symshen_4assoc_1_6), symshen_4parse_1failure, tmp19245, tmp19257)


tmp19259 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19258)

_ = tmp19259

tmp19260 := MakeNative1(func(__e *ControlFlow, V6264 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6260 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6261 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6262 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6263 Obj) Obj {
tmp19261 := Call1(__e, SymFn(symshen_4newpv), B6260)


let__18040 := tmp19261
_ = let__18040

tmp19262 := Call1(__e, SymFn(symshen_4newpv), B6260)


let__18041 := tmp19262
_ = let__18041

tmp19263 := GCons(__e, let__18040, Nil)

tmp19264 := GCons(__e, symlist, tmp19263)

tmp19265 := GCons(__e, let__18041, Nil)

tmp19266 := GCons(__e, tmp19264, tmp19265)

tmp19267 := GCons(__e, symstr, tmp19266)

tmp19268 := GCons(__e, let__18041, Nil)

tmp19269 := GCons(__e, sym_1_1_6, tmp19268)

tmp19270 := GCons(__e, tmp19267, tmp19269)

tmp19271 := Call6(__e, SymFn(symis_b), V6264, tmp19270, B6260, L6261, Key6262, C6263)


tmp19272 := Call2(__e, SymFn(symshen_4gc), B6260, tmp19271)


return TailCall2(__e, SymFn(symshen_4gc), B6260, tmp19272)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19273 := GValue(__e, symshen_4_dsigf_d)

tmp19274 := Call3(__e, SymFn(symshen_4assoc_1_6), symshen_4_5_1out, tmp19260, tmp19273)


tmp19275 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19274)

_ = tmp19275

tmp19276 := MakeNative1(func(__e *ControlFlow, V6270 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6266 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6267 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6268 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6269 Obj) Obj {
tmp19277 := Call1(__e, SymFn(symshen_4newpv), B6266)


let__18042 := tmp19277
_ = let__18042

tmp19278 := Call1(__e, SymFn(symshen_4newpv), B6266)


let__18043 := tmp19278
_ = let__18043

tmp19279 := GCons(__e, let__18043, Nil)

tmp19280 := GCons(__e, symlist, tmp19279)

tmp19281 := GCons(__e, let__18042, Nil)

tmp19282 := GCons(__e, tmp19280, tmp19281)

tmp19283 := GCons(__e, symstr, tmp19282)

tmp19284 := GCons(__e, let__18043, Nil)

tmp19285 := GCons(__e, symlist, tmp19284)

tmp19286 := GCons(__e, tmp19285, Nil)

tmp19287 := GCons(__e, sym_1_1_6, tmp19286)

tmp19288 := GCons(__e, tmp19283, tmp19287)

tmp19289 := Call6(__e, SymFn(symis_b), V6270, tmp19288, B6266, L6267, Key6268, C6269)


tmp19290 := Call2(__e, SymFn(symshen_4gc), B6266, tmp19289)


return TailCall2(__e, SymFn(symshen_4gc), B6266, tmp19290)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19291 := GValue(__e, symshen_4_dsigf_d)

tmp19292 := Call3(__e, SymFn(symshen_4assoc_1_6), symshen_4in_1_6, tmp19276, tmp19291)


tmp19293 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19292)

_ = tmp19293

tmp19294 := MakeNative1(func(__e *ControlFlow, V6276 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6272 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6273 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6274 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6275 Obj) Obj {
tmp19295 := Call1(__e, SymFn(symshen_4newpv), B6272)


let__18044 := tmp19295
_ = let__18044

tmp19296 := Call1(__e, SymFn(symshen_4newpv), B6272)


let__18045 := tmp19296
_ = let__18045

tmp19297 := GCons(__e, let__18044, Nil)

tmp19298 := GCons(__e, symlist, tmp19297)

tmp19299 := GCons(__e, let__18044, Nil)

tmp19300 := GCons(__e, symlist, tmp19299)

tmp19301 := GCons(__e, let__18045, Nil)

tmp19302 := GCons(__e, tmp19300, tmp19301)

tmp19303 := GCons(__e, symstr, tmp19302)

tmp19304 := GCons(__e, tmp19303, Nil)

tmp19305 := GCons(__e, sym_1_1_6, tmp19304)

tmp19306 := GCons(__e, let__18045, tmp19305)

tmp19307 := GCons(__e, tmp19306, Nil)

tmp19308 := GCons(__e, sym_1_1_6, tmp19307)

tmp19309 := GCons(__e, tmp19298, tmp19308)

tmp19310 := Call6(__e, SymFn(symis_b), V6276, tmp19309, B6272, L6273, Key6274, C6275)


tmp19311 := Call2(__e, SymFn(symshen_4gc), B6272, tmp19310)


return TailCall2(__e, SymFn(symshen_4gc), B6272, tmp19311)


panic("unreachable")
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

tmp19313 := Call3(__e, SymFn(symshen_4assoc_1_6), symshen_4comb, tmp19294, tmp19312)


tmp19314 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19313)

_ = tmp19314

tmp19315 := MakeNative1(func(__e *ControlFlow, V6281 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6277 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6278 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6279 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6280 Obj) Obj {
tmp19316 := Call1(__e, SymFn(symshen_4newpv), B6277)


let__18046 := tmp19316
_ = let__18046

tmp19317 := GCons(__e, let__18046, Nil)

tmp19318 := GCons(__e, symlist, tmp19317)

tmp19319 := GCons(__e, symboolean, Nil)

tmp19320 := GCons(__e, sym_1_1_6, tmp19319)

tmp19321 := GCons(__e, tmp19318, tmp19320)

tmp19322 := GCons(__e, tmp19321, Nil)

tmp19323 := GCons(__e, sym_1_1_6, tmp19322)

tmp19324 := GCons(__e, let__18046, tmp19323)

tmp19325 := Call6(__e, SymFn(symis_b), V6281, tmp19324, B6277, L6278, Key6279, C6280)


return TailCall2(__e, SymFn(symshen_4gc), B6277, tmp19325)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19326 := GValue(__e, symshen_4_dsigf_d)

tmp19327 := Call3(__e, SymFn(symshen_4assoc_1_6), symelement_2, tmp19315, tmp19326)


tmp19328 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19327)

_ = tmp19328

tmp19329 := MakeNative1(func(__e *ControlFlow, V6286 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6282 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6283 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6284 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6285 Obj) Obj {
tmp19330 := Call1(__e, SymFn(symshen_4newpv), B6282)


let__18047 := tmp19330
_ = let__18047

tmp19331 := GCons(__e, symboolean, Nil)

tmp19332 := GCons(__e, sym_1_1_6, tmp19331)

tmp19333 := GCons(__e, let__18047, tmp19332)

tmp19334 := Call6(__e, SymFn(symis_b), V6286, tmp19333, B6282, L6283, Key6284, C6285)


return TailCall2(__e, SymFn(symshen_4gc), B6282, tmp19334)


panic("unreachable")
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

tmp19336 := Call3(__e, SymFn(symshen_4assoc_1_6), symempty_2, tmp19329, tmp19335)


tmp19337 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19336)

_ = tmp19337

tmp19338 := MakeNative1(func(__e *ControlFlow, V6291 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6287 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6288 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6289 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6290 Obj) Obj {
tmp19339 := GCons(__e, symboolean, Nil)

tmp19340 := GCons(__e, sym_1_1_6, tmp19339)

tmp19341 := GCons(__e, symsymbol, tmp19340)

return TailCall6(__e, SymFn(symis_b), V6291, tmp19341, B6287, L6288, Key6289, C6290)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19342 := GValue(__e, symshen_4_dsigf_d)

tmp19343 := Call3(__e, SymFn(symshen_4assoc_1_6), symenable_1type_1theory, tmp19338, tmp19342)


tmp19344 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19343)

_ = tmp19344

tmp19345 := MakeNative1(func(__e *ControlFlow, V6296 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6292 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6293 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6294 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6295 Obj) Obj {
tmp19346 := GCons(__e, symsymbol, Nil)

tmp19347 := GCons(__e, symlist, tmp19346)

tmp19348 := GCons(__e, tmp19347, Nil)

tmp19349 := GCons(__e, sym_1_1_6, tmp19348)

tmp19350 := GCons(__e, symsymbol, tmp19349)

return TailCall6(__e, SymFn(symis_b), V6296, tmp19350, B6292, L6293, Key6294, C6295)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19351 := GValue(__e, symshen_4_dsigf_d)

tmp19352 := Call3(__e, SymFn(symshen_4assoc_1_6), symexternal, tmp19345, tmp19351)


tmp19353 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19352)

_ = tmp19353

tmp19354 := MakeNative1(func(__e *ControlFlow, V6301 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6297 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6298 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6299 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6300 Obj) Obj {
tmp19355 := GCons(__e, symstring, Nil)

tmp19356 := GCons(__e, sym_1_1_6, tmp19355)

tmp19357 := GCons(__e, symexception, tmp19356)

return TailCall6(__e, SymFn(symis_b), V6301, tmp19357, B6297, L6298, Key6299, C6300)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19358 := GValue(__e, symshen_4_dsigf_d)

tmp19359 := Call3(__e, SymFn(symshen_4assoc_1_6), symerror_1to_1string, tmp19354, tmp19358)


tmp19360 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19359)

_ = tmp19360

tmp19361 := MakeNative1(func(__e *ControlFlow, V6306 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6302 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6303 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6304 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6305 Obj) Obj {
tmp19362 := Call1(__e, SymFn(symshen_4newpv), B6302)


let__18048 := tmp19362
_ = let__18048

tmp19363 := GCons(__e, symstring, Nil)

tmp19364 := GCons(__e, symlist, tmp19363)

tmp19365 := GCons(__e, tmp19364, Nil)

tmp19366 := GCons(__e, sym_1_1_6, tmp19365)

tmp19367 := GCons(__e, let__18048, tmp19366)

tmp19368 := Call6(__e, SymFn(symis_b), V6306, tmp19367, B6302, L6303, Key6304, C6305)


return TailCall2(__e, SymFn(symshen_4gc), B6302, tmp19368)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19369 := GValue(__e, symshen_4_dsigf_d)

tmp19370 := Call3(__e, SymFn(symshen_4assoc_1_6), symexplode, tmp19361, tmp19369)


tmp19371 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19370)

_ = tmp19371

tmp19372 := MakeNative1(func(__e *ControlFlow, V6311 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6307 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6308 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6309 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6310 Obj) Obj {
tmp19373 := GCons(__e, symsymbol, Nil)

tmp19374 := GCons(__e, sym_1_1_6, tmp19373)

tmp19375 := GCons(__e, symsymbol, tmp19374)

return TailCall6(__e, SymFn(symis_b), V6311, tmp19375, B6307, L6308, Key6309, C6310)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19376 := GValue(__e, symshen_4_dsigf_d)

tmp19377 := Call3(__e, SymFn(symshen_4assoc_1_6), symfactorise, tmp19372, tmp19376)


tmp19378 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19377)

_ = tmp19378

tmp19379 := MakeNative1(func(__e *ControlFlow, V6316 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6312 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6313 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6314 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6315 Obj) Obj {
tmp19380 := GCons(__e, symboolean, Nil)

tmp19381 := GCons(__e, sym_1_1_6, tmp19380)

return TailCall6(__e, SymFn(symis_b), V6316, tmp19381, B6312, L6313, Key6314, C6315)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19382 := GValue(__e, symshen_4_dsigf_d)

tmp19383 := Call3(__e, SymFn(symshen_4assoc_1_6), symfactorise_2, tmp19379, tmp19382)


tmp19384 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19383)

_ = tmp19384

tmp19385 := MakeNative1(func(__e *ControlFlow, V6321 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6317 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6318 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6319 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6320 Obj) Obj {
tmp19386 := GCons(__e, symsymbol, Nil)

tmp19387 := GCons(__e, sym_1_1_6, tmp19386)

return TailCall6(__e, SymFn(symis_b), V6321, tmp19387, B6317, L6318, Key6319, C6320)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19388 := GValue(__e, symshen_4_dsigf_d)

tmp19389 := Call3(__e, SymFn(symshen_4assoc_1_6), symfail, tmp19385, tmp19388)


tmp19390 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19389)

_ = tmp19390

tmp19391 := MakeNative1(func(__e *ControlFlow, V6327 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6323 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6324 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6325 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6326 Obj) Obj {
tmp19392 := Call1(__e, SymFn(symshen_4newpv), B6323)


let__18049 := tmp19392
_ = let__18049

tmp19393 := GCons(__e, let__18049, Nil)

tmp19394 := GCons(__e, sym_1_1_6, tmp19393)

tmp19395 := GCons(__e, let__18049, tmp19394)

tmp19396 := GCons(__e, let__18049, Nil)

tmp19397 := GCons(__e, sym_1_1_6, tmp19396)

tmp19398 := GCons(__e, let__18049, tmp19397)

tmp19399 := GCons(__e, tmp19398, Nil)

tmp19400 := GCons(__e, sym_1_1_6, tmp19399)

tmp19401 := GCons(__e, tmp19395, tmp19400)

tmp19402 := Call6(__e, SymFn(symis_b), V6327, tmp19401, B6323, L6324, Key6325, C6326)


return TailCall2(__e, SymFn(symshen_4gc), B6323, tmp19402)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19403 := GValue(__e, symshen_4_dsigf_d)

tmp19404 := Call3(__e, SymFn(symshen_4assoc_1_6), symfix, tmp19391, tmp19403)


tmp19405 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19404)

_ = tmp19405

tmp19406 := MakeNative1(func(__e *ControlFlow, V6332 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6328 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6329 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6330 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6331 Obj) Obj {
tmp19407 := Call1(__e, SymFn(symshen_4newpv), B6328)


let__18050 := tmp19407
_ = let__18050

tmp19408 := GCons(__e, let__18050, Nil)

tmp19409 := GCons(__e, symlazy, tmp19408)

tmp19410 := GCons(__e, tmp19409, Nil)

tmp19411 := GCons(__e, sym_1_1_6, tmp19410)

tmp19412 := GCons(__e, let__18050, tmp19411)

tmp19413 := Call6(__e, SymFn(symis_b), V6332, tmp19412, B6328, L6329, Key6330, C6331)


return TailCall2(__e, SymFn(symshen_4gc), B6328, tmp19413)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19414 := GValue(__e, symshen_4_dsigf_d)

tmp19415 := Call3(__e, SymFn(symshen_4assoc_1_6), symfreeze, tmp19406, tmp19414)


tmp19416 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19415)

_ = tmp19416

tmp19417 := MakeNative1(func(__e *ControlFlow, V6337 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6333 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6334 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6335 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6336 Obj) Obj {
tmp19418 := Call1(__e, SymFn(symshen_4newpv), B6333)


let__18051 := tmp19418
_ = let__18051

tmp19419 := Call1(__e, SymFn(symshen_4newpv), B6333)


let__18052 := tmp19419
_ = let__18052

tmp19420 := GCons(__e, let__18051, Nil)

tmp19421 := GCons(__e, sym_d, tmp19420)

tmp19422 := GCons(__e, let__18052, tmp19421)

tmp19423 := GCons(__e, let__18052, Nil)

tmp19424 := GCons(__e, sym_1_1_6, tmp19423)

tmp19425 := GCons(__e, tmp19422, tmp19424)

tmp19426 := Call6(__e, SymFn(symis_b), V6337, tmp19425, B6333, L6334, Key6335, C6336)


tmp19427 := Call2(__e, SymFn(symshen_4gc), B6333, tmp19426)


return TailCall2(__e, SymFn(symshen_4gc), B6333, tmp19427)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19428 := GValue(__e, symshen_4_dsigf_d)

tmp19429 := Call3(__e, SymFn(symshen_4assoc_1_6), symfst, tmp19417, tmp19428)


tmp19430 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19429)

_ = tmp19430

tmp19431 := MakeNative1(func(__e *ControlFlow, V6342 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6338 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6339 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6340 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6341 Obj) Obj {
tmp19432 := GCons(__e, symsymbol, Nil)

tmp19433 := GCons(__e, sym_1_1_6, tmp19432)

tmp19434 := GCons(__e, symsymbol, tmp19433)

return TailCall6(__e, SymFn(symis_b), V6342, tmp19434, B6338, L6339, Key6340, C6341)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19435 := GValue(__e, symshen_4_dsigf_d)

tmp19436 := Call3(__e, SymFn(symshen_4assoc_1_6), symgensym, tmp19431, tmp19435)


tmp19437 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19436)

_ = tmp19437

tmp19438 := MakeNative1(func(__e *ControlFlow, V6348 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6344 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6345 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6346 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6347 Obj) Obj {
tmp19439 := Call1(__e, SymFn(symshen_4newpv), B6344)


let__18053 := tmp19439
_ = let__18053

tmp19440 := GCons(__e, let__18053, Nil)

tmp19441 := GCons(__e, symlist, tmp19440)

tmp19442 := GCons(__e, symboolean, Nil)

tmp19443 := GCons(__e, sym_1_1_6, tmp19442)

tmp19444 := GCons(__e, let__18053, tmp19443)

tmp19445 := GCons(__e, tmp19444, Nil)

tmp19446 := GCons(__e, sym_1_1_6, tmp19445)

tmp19447 := GCons(__e, tmp19441, tmp19446)

tmp19448 := Call6(__e, SymFn(symis_b), V6348, tmp19447, B6344, L6345, Key6346, C6347)


return TailCall2(__e, SymFn(symshen_4gc), B6344, tmp19448)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19449 := GValue(__e, symshen_4_dsigf_d)

tmp19450 := Call3(__e, SymFn(symshen_4assoc_1_6), symshen_4hds_a_2, tmp19438, tmp19449)


tmp19451 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19450)

_ = tmp19451

tmp19452 := MakeNative1(func(__e *ControlFlow, V6353 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6349 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6350 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6351 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6352 Obj) Obj {
tmp19453 := GCons(__e, symboolean, Nil)

tmp19454 := GCons(__e, sym_1_1_6, tmp19453)

tmp19455 := GCons(__e, symsymbol, tmp19454)

return TailCall6(__e, SymFn(symis_b), V6353, tmp19455, B6349, L6350, Key6351, C6352)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19456 := GValue(__e, symshen_4_dsigf_d)

tmp19457 := Call3(__e, SymFn(symshen_4assoc_1_6), symhush, tmp19452, tmp19456)


tmp19458 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19457)

_ = tmp19458

tmp19459 := MakeNative1(func(__e *ControlFlow, V6358 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6354 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6355 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6356 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6357 Obj) Obj {
tmp19460 := GCons(__e, symboolean, Nil)

tmp19461 := GCons(__e, sym_1_1_6, tmp19460)

return TailCall6(__e, SymFn(symis_b), V6358, tmp19461, B6354, L6355, Key6356, C6357)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19462 := GValue(__e, symshen_4_dsigf_d)

tmp19463 := Call3(__e, SymFn(symshen_4assoc_1_6), symhush_2, tmp19459, tmp19462)


tmp19464 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19463)

_ = tmp19464

tmp19465 := MakeNative1(func(__e *ControlFlow, V6364 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6360 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6361 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6362 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6363 Obj) Obj {
tmp19466 := Call1(__e, SymFn(symshen_4newpv), B6360)


let__18054 := tmp19466
_ = let__18054

tmp19467 := GCons(__e, let__18054, Nil)

tmp19468 := GCons(__e, symvector, tmp19467)

tmp19469 := GCons(__e, let__18054, Nil)

tmp19470 := GCons(__e, sym_1_1_6, tmp19469)

tmp19471 := GCons(__e, symnumber, tmp19470)

tmp19472 := GCons(__e, tmp19471, Nil)

tmp19473 := GCons(__e, sym_1_1_6, tmp19472)

tmp19474 := GCons(__e, tmp19468, tmp19473)

tmp19475 := Call6(__e, SymFn(symis_b), V6364, tmp19474, B6360, L6361, Key6362, C6363)


return TailCall2(__e, SymFn(symshen_4gc), B6360, tmp19475)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19476 := GValue(__e, symshen_4_dsigf_d)

tmp19477 := Call3(__e, SymFn(symshen_4assoc_1_6), sym_5_1vector, tmp19465, tmp19476)


tmp19478 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19477)

_ = tmp19478

tmp19479 := MakeNative1(func(__e *ControlFlow, V6370 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6366 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6367 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6368 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6369 Obj) Obj {
tmp19480 := Call1(__e, SymFn(symshen_4newpv), B6366)


let__18055 := tmp19480
_ = let__18055

tmp19481 := GCons(__e, let__18055, Nil)

tmp19482 := GCons(__e, symvector, tmp19481)

tmp19483 := GCons(__e, let__18055, Nil)

tmp19484 := GCons(__e, symvector, tmp19483)

tmp19485 := GCons(__e, tmp19484, Nil)

tmp19486 := GCons(__e, sym_1_1_6, tmp19485)

tmp19487 := GCons(__e, let__18055, tmp19486)

tmp19488 := GCons(__e, tmp19487, Nil)

tmp19489 := GCons(__e, sym_1_1_6, tmp19488)

tmp19490 := GCons(__e, symnumber, tmp19489)

tmp19491 := GCons(__e, tmp19490, Nil)

tmp19492 := GCons(__e, sym_1_1_6, tmp19491)

tmp19493 := GCons(__e, tmp19482, tmp19492)

tmp19494 := Call6(__e, SymFn(symis_b), V6370, tmp19493, B6366, L6367, Key6368, C6369)


return TailCall2(__e, SymFn(symshen_4gc), B6366, tmp19494)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19495 := GValue(__e, symshen_4_dsigf_d)

tmp19496 := Call3(__e, SymFn(symshen_4assoc_1_6), symvector_1_6, tmp19479, tmp19495)


tmp19497 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19496)

_ = tmp19497

tmp19498 := MakeNative1(func(__e *ControlFlow, V6375 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6371 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6372 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6373 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6374 Obj) Obj {
tmp19499 := Call1(__e, SymFn(symshen_4newpv), B6371)


let__18056 := tmp19499
_ = let__18056

tmp19500 := GCons(__e, let__18056, Nil)

tmp19501 := GCons(__e, symvector, tmp19500)

tmp19502 := GCons(__e, tmp19501, Nil)

tmp19503 := GCons(__e, sym_1_1_6, tmp19502)

tmp19504 := GCons(__e, symnumber, tmp19503)

tmp19505 := Call6(__e, SymFn(symis_b), V6375, tmp19504, B6371, L6372, Key6373, C6374)


return TailCall2(__e, SymFn(symshen_4gc), B6371, tmp19505)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19506 := GValue(__e, symshen_4_dsigf_d)

tmp19507 := Call3(__e, SymFn(symshen_4assoc_1_6), symvector, tmp19498, tmp19506)


tmp19508 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19507)

_ = tmp19508

tmp19509 := MakeNative1(func(__e *ControlFlow, V6380 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6376 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6377 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6378 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6379 Obj) Obj {
tmp19510 := GCons(__e, symnumber, Nil)

tmp19511 := GCons(__e, sym_1_1_6, tmp19510)

tmp19512 := GCons(__e, symsymbol, tmp19511)

return TailCall6(__e, SymFn(symis_b), V6380, tmp19512, B6376, L6377, Key6378, C6379)


panic("unreachable")
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

tmp19514 := Call3(__e, SymFn(symshen_4assoc_1_6), symget_1time, tmp19509, tmp19513)


tmp19515 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19514)

_ = tmp19515

tmp19516 := MakeNative1(func(__e *ControlFlow, V6385 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6381 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6382 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6383 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6384 Obj) Obj {
tmp19517 := Call1(__e, SymFn(symshen_4newpv), B6381)


let__18057 := tmp19517
_ = let__18057

tmp19518 := GCons(__e, symnumber, Nil)

tmp19519 := GCons(__e, sym_1_1_6, tmp19518)

tmp19520 := GCons(__e, symnumber, tmp19519)

tmp19521 := GCons(__e, tmp19520, Nil)

tmp19522 := GCons(__e, sym_1_1_6, tmp19521)

tmp19523 := GCons(__e, let__18057, tmp19522)

tmp19524 := Call6(__e, SymFn(symis_b), V6385, tmp19523, B6381, L6382, Key6383, C6384)


return TailCall2(__e, SymFn(symshen_4gc), B6381, tmp19524)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19525 := GValue(__e, symshen_4_dsigf_d)

tmp19526 := Call3(__e, SymFn(symshen_4assoc_1_6), symhash, tmp19516, tmp19525)


tmp19527 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19526)

_ = tmp19527

tmp19528 := MakeNative1(func(__e *ControlFlow, V6390 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6386 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6387 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6388 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6389 Obj) Obj {
tmp19529 := Call1(__e, SymFn(symshen_4newpv), B6386)


let__18058 := tmp19529
_ = let__18058

tmp19530 := GCons(__e, let__18058, Nil)

tmp19531 := GCons(__e, symlist, tmp19530)

tmp19532 := GCons(__e, let__18058, Nil)

tmp19533 := GCons(__e, sym_1_1_6, tmp19532)

tmp19534 := GCons(__e, tmp19531, tmp19533)

tmp19535 := Call6(__e, SymFn(symis_b), V6390, tmp19534, B6386, L6387, Key6388, C6389)


return TailCall2(__e, SymFn(symshen_4gc), B6386, tmp19535)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19536 := GValue(__e, symshen_4_dsigf_d)

tmp19537 := Call3(__e, SymFn(symshen_4assoc_1_6), symhead, tmp19528, tmp19536)


tmp19538 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19537)

_ = tmp19538

tmp19539 := MakeNative1(func(__e *ControlFlow, V6395 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6391 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6392 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6393 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6394 Obj) Obj {
tmp19540 := Call1(__e, SymFn(symshen_4newpv), B6391)


let__18059 := tmp19540
_ = let__18059

tmp19541 := GCons(__e, let__18059, Nil)

tmp19542 := GCons(__e, symvector, tmp19541)

tmp19543 := GCons(__e, let__18059, Nil)

tmp19544 := GCons(__e, sym_1_1_6, tmp19543)

tmp19545 := GCons(__e, tmp19542, tmp19544)

tmp19546 := Call6(__e, SymFn(symis_b), V6395, tmp19545, B6391, L6392, Key6393, C6394)


return TailCall2(__e, SymFn(symshen_4gc), B6391, tmp19546)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19547 := GValue(__e, symshen_4_dsigf_d)

tmp19548 := Call3(__e, SymFn(symshen_4assoc_1_6), symhdv, tmp19539, tmp19547)


tmp19549 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19548)

_ = tmp19549

tmp19550 := MakeNative1(func(__e *ControlFlow, V6400 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6396 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6397 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6398 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6399 Obj) Obj {
tmp19551 := GCons(__e, symstring, Nil)

tmp19552 := GCons(__e, sym_1_1_6, tmp19551)

tmp19553 := GCons(__e, symstring, tmp19552)

return TailCall6(__e, SymFn(symis_b), V6400, tmp19553, B6396, L6397, Key6398, C6399)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19554 := GValue(__e, symshen_4_dsigf_d)

tmp19555 := Call3(__e, SymFn(symshen_4assoc_1_6), symhdstr, tmp19550, tmp19554)


tmp19556 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19555)

_ = tmp19556

tmp19557 := MakeNative1(func(__e *ControlFlow, V6405 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6401 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6402 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6403 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6404 Obj) Obj {
tmp19558 := Call1(__e, SymFn(symshen_4newpv), B6401)


let__18060 := tmp19558
_ = let__18060

tmp19559 := GCons(__e, let__18060, Nil)

tmp19560 := GCons(__e, sym_1_1_6, tmp19559)

tmp19561 := GCons(__e, let__18060, tmp19560)

tmp19562 := GCons(__e, tmp19561, Nil)

tmp19563 := GCons(__e, sym_1_1_6, tmp19562)

tmp19564 := GCons(__e, let__18060, tmp19563)

tmp19565 := GCons(__e, tmp19564, Nil)

tmp19566 := GCons(__e, sym_1_1_6, tmp19565)

tmp19567 := GCons(__e, symboolean, tmp19566)

tmp19568 := Call6(__e, SymFn(symis_b), V6405, tmp19567, B6401, L6402, Key6403, C6404)


return TailCall2(__e, SymFn(symshen_4gc), B6401, tmp19568)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19569 := GValue(__e, symshen_4_dsigf_d)

tmp19570 := Call3(__e, SymFn(symshen_4assoc_1_6), symif, tmp19557, tmp19569)


tmp19571 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19570)

_ = tmp19571

tmp19572 := MakeNative1(func(__e *ControlFlow, V6410 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6406 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6407 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6408 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6409 Obj) Obj {
tmp19573 := GCons(__e, symsymbol, Nil)

tmp19574 := GCons(__e, sym_1_1_6, tmp19573)

tmp19575 := GCons(__e, symsymbol, tmp19574)

return TailCall6(__e, SymFn(symis_b), V6410, tmp19575, B6406, L6407, Key6408, C6409)


panic("unreachable")
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

tmp19577 := Call3(__e, SymFn(symshen_4assoc_1_6), symin_1package, tmp19572, tmp19576)


tmp19578 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19577)

_ = tmp19578

tmp19579 := MakeNative1(func(__e *ControlFlow, V6415 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6411 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6412 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6413 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6414 Obj) Obj {
tmp19580 := GCons(__e, symstring, Nil)

tmp19581 := GCons(__e, sym_1_1_6, tmp19580)

return TailCall6(__e, SymFn(symis_b), V6415, tmp19581, B6411, L6412, Key6413, C6414)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19582 := GValue(__e, symshen_4_dsigf_d)

tmp19583 := Call3(__e, SymFn(symshen_4assoc_1_6), symit, tmp19579, tmp19582)


tmp19584 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19583)

_ = tmp19584

tmp19585 := MakeNative1(func(__e *ControlFlow, V6420 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6416 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6417 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6418 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6419 Obj) Obj {
tmp19586 := GCons(__e, symstring, Nil)

tmp19587 := GCons(__e, sym_1_1_6, tmp19586)

return TailCall6(__e, SymFn(symis_b), V6420, tmp19587, B6416, L6417, Key6418, C6419)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19588 := GValue(__e, symshen_4_dsigf_d)

tmp19589 := Call3(__e, SymFn(symshen_4assoc_1_6), symimplementation, tmp19585, tmp19588)


tmp19590 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19589)

_ = tmp19590

tmp19591 := MakeNative1(func(__e *ControlFlow, V6426 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6422 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6423 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6424 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6425 Obj) Obj {
tmp19592 := GCons(__e, symsymbol, Nil)

tmp19593 := GCons(__e, symlist, tmp19592)

tmp19594 := GCons(__e, symsymbol, Nil)

tmp19595 := GCons(__e, symlist, tmp19594)

tmp19596 := GCons(__e, tmp19595, Nil)

tmp19597 := GCons(__e, sym_1_1_6, tmp19596)

tmp19598 := GCons(__e, tmp19593, tmp19597)

return TailCall6(__e, SymFn(symis_b), V6426, tmp19598, B6422, L6423, Key6424, C6425)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19599 := GValue(__e, symshen_4_dsigf_d)

tmp19600 := Call3(__e, SymFn(symshen_4assoc_1_6), syminclude, tmp19591, tmp19599)


tmp19601 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19600)

_ = tmp19601

tmp19602 := MakeNative1(func(__e *ControlFlow, V6432 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6428 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6429 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6430 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6431 Obj) Obj {
tmp19603 := GCons(__e, symsymbol, Nil)

tmp19604 := GCons(__e, symlist, tmp19603)

tmp19605 := GCons(__e, symsymbol, Nil)

tmp19606 := GCons(__e, symlist, tmp19605)

tmp19607 := GCons(__e, tmp19606, Nil)

tmp19608 := GCons(__e, sym_1_1_6, tmp19607)

tmp19609 := GCons(__e, tmp19604, tmp19608)

return TailCall6(__e, SymFn(symis_b), V6432, tmp19609, B6428, L6429, Key6430, C6431)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19610 := GValue(__e, symshen_4_dsigf_d)

tmp19611 := Call3(__e, SymFn(symshen_4assoc_1_6), syminclude_1all_1but, tmp19602, tmp19610)


tmp19612 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19611)

_ = tmp19612

tmp19613 := MakeNative1(func(__e *ControlFlow, V6437 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6433 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6434 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6435 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6436 Obj) Obj {
tmp19614 := GCons(__e, symsymbol, Nil)

tmp19615 := GCons(__e, symlist, tmp19614)

tmp19616 := GCons(__e, tmp19615, Nil)

tmp19617 := GCons(__e, sym_1_1_6, tmp19616)

return TailCall6(__e, SymFn(symis_b), V6437, tmp19617, B6433, L6434, Key6435, C6436)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19618 := GValue(__e, symshen_4_dsigf_d)

tmp19619 := Call3(__e, SymFn(symshen_4assoc_1_6), symincluded, tmp19613, tmp19618)


tmp19620 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19619)

_ = tmp19620

tmp19621 := MakeNative1(func(__e *ControlFlow, V6442 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6438 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6439 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6440 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6441 Obj) Obj {
tmp19622 := GCons(__e, symnumber, Nil)

tmp19623 := GCons(__e, sym_1_1_6, tmp19622)

return TailCall6(__e, SymFn(symis_b), V6442, tmp19623, B6438, L6439, Key6440, C6441)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19624 := GValue(__e, symshen_4_dsigf_d)

tmp19625 := Call3(__e, SymFn(symshen_4assoc_1_6), syminferences, tmp19621, tmp19624)


tmp19626 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19625)

_ = tmp19626

tmp19627 := MakeNative1(func(__e *ControlFlow, V6447 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6443 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6444 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6445 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6446 Obj) Obj {
tmp19628 := Call1(__e, SymFn(symshen_4newpv), B6443)


let__18061 := tmp19628
_ = let__18061

tmp19629 := GCons(__e, symstring, Nil)

tmp19630 := GCons(__e, sym_1_1_6, tmp19629)

tmp19631 := GCons(__e, symstring, tmp19630)

tmp19632 := GCons(__e, tmp19631, Nil)

tmp19633 := GCons(__e, sym_1_1_6, tmp19632)

tmp19634 := GCons(__e, let__18061, tmp19633)

tmp19635 := Call6(__e, SymFn(symis_b), V6447, tmp19634, B6443, L6444, Key6445, C6446)


return TailCall2(__e, SymFn(symshen_4gc), B6443, tmp19635)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19636 := GValue(__e, symshen_4_dsigf_d)

tmp19637 := Call3(__e, SymFn(symshen_4assoc_1_6), symshen_4insert, tmp19627, tmp19636)


tmp19638 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19637)

_ = tmp19638

tmp19639 := MakeNative1(func(__e *ControlFlow, V6452 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6448 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6449 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6450 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6451 Obj) Obj {
tmp19640 := Call1(__e, SymFn(symshen_4newpv), B6448)


let__18062 := tmp19640
_ = let__18062

tmp19641 := GCons(__e, symboolean, Nil)

tmp19642 := GCons(__e, sym_1_1_6, tmp19641)

tmp19643 := GCons(__e, let__18062, tmp19642)

tmp19644 := Call6(__e, SymFn(symis_b), V6452, tmp19643, B6448, L6449, Key6450, C6451)


return TailCall2(__e, SymFn(symshen_4gc), B6448, tmp19644)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19645 := GValue(__e, symshen_4_dsigf_d)

tmp19646 := Call3(__e, SymFn(symshen_4assoc_1_6), syminteger_2, tmp19639, tmp19645)


tmp19647 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19646)

_ = tmp19647

tmp19648 := MakeNative1(func(__e *ControlFlow, V6457 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6453 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6454 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6455 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6456 Obj) Obj {
tmp19649 := GCons(__e, symsymbol, Nil)

tmp19650 := GCons(__e, symlist, tmp19649)

tmp19651 := GCons(__e, tmp19650, Nil)

tmp19652 := GCons(__e, sym_1_1_6, tmp19651)

tmp19653 := GCons(__e, symsymbol, tmp19652)

return TailCall6(__e, SymFn(symis_b), V6457, tmp19653, B6453, L6454, Key6455, C6456)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19654 := GValue(__e, symshen_4_dsigf_d)

tmp19655 := Call3(__e, SymFn(symshen_4assoc_1_6), syminternal, tmp19648, tmp19654)


tmp19656 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19655)

_ = tmp19656

tmp19657 := MakeNative1(func(__e *ControlFlow, V6464 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6460 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6461 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6462 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6463 Obj) Obj {
tmp19658 := Call1(__e, SymFn(symshen_4newpv), B6460)


let__18063 := tmp19658
_ = let__18063

tmp19659 := GCons(__e, let__18063, Nil)

tmp19660 := GCons(__e, symlist, tmp19659)

tmp19661 := GCons(__e, let__18063, Nil)

tmp19662 := GCons(__e, symlist, tmp19661)

tmp19663 := GCons(__e, let__18063, Nil)

tmp19664 := GCons(__e, symlist, tmp19663)

tmp19665 := GCons(__e, tmp19664, Nil)

tmp19666 := GCons(__e, sym_1_1_6, tmp19665)

tmp19667 := GCons(__e, tmp19662, tmp19666)

tmp19668 := GCons(__e, tmp19667, Nil)

tmp19669 := GCons(__e, sym_1_1_6, tmp19668)

tmp19670 := GCons(__e, tmp19660, tmp19669)

tmp19671 := Call6(__e, SymFn(symis_b), V6464, tmp19670, B6460, L6461, Key6462, C6463)


return TailCall2(__e, SymFn(symshen_4gc), B6460, tmp19671)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19672 := GValue(__e, symshen_4_dsigf_d)

tmp19673 := Call3(__e, SymFn(symshen_4assoc_1_6), symintersection, tmp19657, tmp19672)


tmp19674 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19673)

_ = tmp19674

tmp19675 := MakeNative1(func(__e *ControlFlow, V6469 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6465 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6466 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6467 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6468 Obj) Obj {
tmp19676 := GCons(__e, symstring, Nil)

tmp19677 := GCons(__e, sym_1_1_6, tmp19676)

return TailCall6(__e, SymFn(symis_b), V6469, tmp19677, B6465, L6466, Key6467, C6468)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19678 := GValue(__e, symshen_4_dsigf_d)

tmp19679 := Call3(__e, SymFn(symshen_4assoc_1_6), symlanguage, tmp19675, tmp19678)


tmp19680 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19679)

_ = tmp19680

tmp19681 := MakeNative1(func(__e *ControlFlow, V6474 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6470 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6471 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6472 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6473 Obj) Obj {
tmp19682 := Call1(__e, SymFn(symshen_4newpv), B6470)


let__18064 := tmp19682
_ = let__18064

tmp19683 := GCons(__e, let__18064, Nil)

tmp19684 := GCons(__e, symlist, tmp19683)

tmp19685 := GCons(__e, symnumber, Nil)

tmp19686 := GCons(__e, sym_1_1_6, tmp19685)

tmp19687 := GCons(__e, tmp19684, tmp19686)

tmp19688 := Call6(__e, SymFn(symis_b), V6474, tmp19687, B6470, L6471, Key6472, C6473)


return TailCall2(__e, SymFn(symshen_4gc), B6470, tmp19688)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19689 := GValue(__e, symshen_4_dsigf_d)

tmp19690 := Call3(__e, SymFn(symshen_4assoc_1_6), symlength, tmp19681, tmp19689)


tmp19691 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19690)

_ = tmp19691

tmp19692 := MakeNative1(func(__e *ControlFlow, V6479 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6475 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6476 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6477 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6478 Obj) Obj {
tmp19693 := Call1(__e, SymFn(symshen_4newpv), B6475)


let__18065 := tmp19693
_ = let__18065

tmp19694 := GCons(__e, let__18065, Nil)

tmp19695 := GCons(__e, symvector, tmp19694)

tmp19696 := GCons(__e, symnumber, Nil)

tmp19697 := GCons(__e, sym_1_1_6, tmp19696)

tmp19698 := GCons(__e, tmp19695, tmp19697)

tmp19699 := Call6(__e, SymFn(symis_b), V6479, tmp19698, B6475, L6476, Key6477, C6478)


return TailCall2(__e, SymFn(symshen_4gc), B6475, tmp19699)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19700 := GValue(__e, symshen_4_dsigf_d)

tmp19701 := Call3(__e, SymFn(symshen_4assoc_1_6), symlimit, tmp19692, tmp19700)


tmp19702 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19701)

_ = tmp19702

tmp19703 := MakeNative1(func(__e *ControlFlow, V6485 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6481 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6482 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6483 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6484 Obj) Obj {
tmp19704 := GCons(__e, symin, Nil)

tmp19705 := GCons(__e, symstream, tmp19704)

tmp19706 := GCons(__e, symunit, Nil)

tmp19707 := GCons(__e, symlist, tmp19706)

tmp19708 := GCons(__e, tmp19707, Nil)

tmp19709 := GCons(__e, sym_1_1_6, tmp19708)

tmp19710 := GCons(__e, tmp19705, tmp19709)

return TailCall6(__e, SymFn(symis_b), V6485, tmp19710, B6481, L6482, Key6483, C6484)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19711 := GValue(__e, symshen_4_dsigf_d)

tmp19712 := Call3(__e, SymFn(symshen_4assoc_1_6), symlineread, tmp19703, tmp19711)


tmp19713 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19712)

_ = tmp19713

tmp19714 := MakeNative1(func(__e *ControlFlow, V6490 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6486 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6487 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6488 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6489 Obj) Obj {
tmp19715 := GCons(__e, symsymbol, Nil)

tmp19716 := GCons(__e, sym_1_1_6, tmp19715)

tmp19717 := GCons(__e, symstring, tmp19716)

return TailCall6(__e, SymFn(symis_b), V6490, tmp19717, B6486, L6487, Key6488, C6489)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19718 := GValue(__e, symshen_4_dsigf_d)

tmp19719 := Call3(__e, SymFn(symshen_4assoc_1_6), symload, tmp19714, tmp19718)


tmp19720 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19719)

_ = tmp19720

tmp19721 := MakeNative1(func(__e *ControlFlow, V6497 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6493 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6494 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6495 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6496 Obj) Obj {
tmp19722 := Call1(__e, SymFn(symshen_4newpv), B6493)


let__18066 := tmp19722
_ = let__18066

tmp19723 := Call1(__e, SymFn(symshen_4newpv), B6493)


let__18067 := tmp19723
_ = let__18067

tmp19724 := GCons(__e, let__18067, Nil)

tmp19725 := GCons(__e, sym_1_1_6, tmp19724)

tmp19726 := GCons(__e, let__18066, tmp19725)

tmp19727 := GCons(__e, let__18066, Nil)

tmp19728 := GCons(__e, symlist, tmp19727)

tmp19729 := GCons(__e, let__18067, Nil)

tmp19730 := GCons(__e, symlist, tmp19729)

tmp19731 := GCons(__e, tmp19730, Nil)

tmp19732 := GCons(__e, sym_1_1_6, tmp19731)

tmp19733 := GCons(__e, tmp19728, tmp19732)

tmp19734 := GCons(__e, tmp19733, Nil)

tmp19735 := GCons(__e, sym_1_1_6, tmp19734)

tmp19736 := GCons(__e, tmp19726, tmp19735)

tmp19737 := Call6(__e, SymFn(symis_b), V6497, tmp19736, B6493, L6494, Key6495, C6496)


tmp19738 := Call2(__e, SymFn(symshen_4gc), B6493, tmp19737)


return TailCall2(__e, SymFn(symshen_4gc), B6493, tmp19738)


panic("unreachable")
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

tmp19740 := Call3(__e, SymFn(symshen_4assoc_1_6), symmap, tmp19721, tmp19739)


tmp19741 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19740)

_ = tmp19741

tmp19742 := MakeNative1(func(__e *ControlFlow, V6504 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6500 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6501 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6502 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6503 Obj) Obj {
tmp19743 := Call1(__e, SymFn(symshen_4newpv), B6500)


let__18068 := tmp19743
_ = let__18068

tmp19744 := Call1(__e, SymFn(symshen_4newpv), B6500)


let__18069 := tmp19744
_ = let__18069

tmp19745 := GCons(__e, let__18069, Nil)

tmp19746 := GCons(__e, symlist, tmp19745)

tmp19747 := GCons(__e, tmp19746, Nil)

tmp19748 := GCons(__e, sym_1_1_6, tmp19747)

tmp19749 := GCons(__e, let__18068, tmp19748)

tmp19750 := GCons(__e, let__18068, Nil)

tmp19751 := GCons(__e, symlist, tmp19750)

tmp19752 := GCons(__e, let__18069, Nil)

tmp19753 := GCons(__e, symlist, tmp19752)

tmp19754 := GCons(__e, tmp19753, Nil)

tmp19755 := GCons(__e, sym_1_1_6, tmp19754)

tmp19756 := GCons(__e, tmp19751, tmp19755)

tmp19757 := GCons(__e, tmp19756, Nil)

tmp19758 := GCons(__e, sym_1_1_6, tmp19757)

tmp19759 := GCons(__e, tmp19749, tmp19758)

tmp19760 := Call6(__e, SymFn(symis_b), V6504, tmp19759, B6500, L6501, Key6502, C6503)


tmp19761 := Call2(__e, SymFn(symshen_4gc), B6500, tmp19760)


return TailCall2(__e, SymFn(symshen_4gc), B6500, tmp19761)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19762 := GValue(__e, symshen_4_dsigf_d)

tmp19763 := Call3(__e, SymFn(symshen_4assoc_1_6), symmapcan, tmp19742, tmp19762)


tmp19764 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19763)

_ = tmp19764

tmp19765 := MakeNative1(func(__e *ControlFlow, V6509 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6505 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6506 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6507 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6508 Obj) Obj {
tmp19766 := GCons(__e, symnumber, Nil)

tmp19767 := GCons(__e, sym_1_1_6, tmp19766)

tmp19768 := GCons(__e, symnumber, tmp19767)

return TailCall6(__e, SymFn(symis_b), V6509, tmp19768, B6505, L6506, Key6507, C6508)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19769 := GValue(__e, symshen_4_dsigf_d)

tmp19770 := Call3(__e, SymFn(symshen_4assoc_1_6), symmaxinferences, tmp19765, tmp19769)


tmp19771 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19770)

_ = tmp19771

tmp19772 := MakeNative1(func(__e *ControlFlow, V6514 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6510 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6511 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6512 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6513 Obj) Obj {
tmp19773 := GCons(__e, symstring, Nil)

tmp19774 := GCons(__e, sym_1_1_6, tmp19773)

tmp19775 := GCons(__e, symnumber, tmp19774)

return TailCall6(__e, SymFn(symis_b), V6514, tmp19775, B6510, L6511, Key6512, C6513)


panic("unreachable")
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

tmp19777 := Call3(__e, SymFn(symshen_4assoc_1_6), symn_1_6string, tmp19772, tmp19776)


tmp19778 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19777)

_ = tmp19778

tmp19779 := MakeNative1(func(__e *ControlFlow, V6519 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6515 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6516 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6517 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6518 Obj) Obj {
tmp19780 := GCons(__e, symnumber, Nil)

tmp19781 := GCons(__e, sym_1_1_6, tmp19780)

tmp19782 := GCons(__e, symnumber, tmp19781)

return TailCall6(__e, SymFn(symis_b), V6519, tmp19782, B6515, L6516, Key6517, C6518)


panic("unreachable")
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

tmp19784 := Call3(__e, SymFn(symshen_4assoc_1_6), symnl, tmp19779, tmp19783)


tmp19785 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19784)

_ = tmp19785

tmp19786 := MakeNative1(func(__e *ControlFlow, V6524 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6520 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6521 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6522 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6523 Obj) Obj {
tmp19787 := GCons(__e, symboolean, Nil)

tmp19788 := GCons(__e, sym_1_1_6, tmp19787)

tmp19789 := GCons(__e, symboolean, tmp19788)

return TailCall6(__e, SymFn(symis_b), V6524, tmp19789, B6520, L6521, Key6522, C6523)


panic("unreachable")
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

tmp19791 := Call3(__e, SymFn(symshen_4assoc_1_6), symnot, tmp19786, tmp19790)


tmp19792 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19791)

_ = tmp19792

tmp19793 := MakeNative1(func(__e *ControlFlow, V6529 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6525 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6526 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6527 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6528 Obj) Obj {
tmp19794 := Call1(__e, SymFn(symshen_4newpv), B6525)


let__18070 := tmp19794
_ = let__18070

tmp19795 := GCons(__e, let__18070, Nil)

tmp19796 := GCons(__e, symlist, tmp19795)

tmp19797 := GCons(__e, let__18070, Nil)

tmp19798 := GCons(__e, sym_1_1_6, tmp19797)

tmp19799 := GCons(__e, tmp19796, tmp19798)

tmp19800 := GCons(__e, tmp19799, Nil)

tmp19801 := GCons(__e, sym_1_1_6, tmp19800)

tmp19802 := GCons(__e, symnumber, tmp19801)

tmp19803 := Call6(__e, SymFn(symis_b), V6529, tmp19802, B6525, L6526, Key6527, C6528)


return TailCall2(__e, SymFn(symshen_4gc), B6525, tmp19803)


panic("unreachable")
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

tmp19805 := Call3(__e, SymFn(symshen_4assoc_1_6), symnth, tmp19793, tmp19804)


tmp19806 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19805)

_ = tmp19806

tmp19807 := MakeNative1(func(__e *ControlFlow, V6534 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6530 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6531 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6532 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6533 Obj) Obj {
tmp19808 := Call1(__e, SymFn(symshen_4newpv), B6530)


let__18071 := tmp19808
_ = let__18071

tmp19809 := GCons(__e, symboolean, Nil)

tmp19810 := GCons(__e, sym_1_1_6, tmp19809)

tmp19811 := GCons(__e, let__18071, tmp19810)

tmp19812 := Call6(__e, SymFn(symis_b), V6534, tmp19811, B6530, L6531, Key6532, C6533)


return TailCall2(__e, SymFn(symshen_4gc), B6530, tmp19812)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19813 := GValue(__e, symshen_4_dsigf_d)

tmp19814 := Call3(__e, SymFn(symshen_4assoc_1_6), symnumber_2, tmp19807, tmp19813)


tmp19815 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19814)

_ = tmp19815

tmp19816 := MakeNative1(func(__e *ControlFlow, V6539 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6535 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6536 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6537 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6538 Obj) Obj {
tmp19817 := Call1(__e, SymFn(symshen_4newpv), B6535)


let__18072 := tmp19817
_ = let__18072

tmp19818 := Call1(__e, SymFn(symshen_4newpv), B6535)


let__18073 := tmp19818
_ = let__18073

tmp19819 := GCons(__e, symnumber, Nil)

tmp19820 := GCons(__e, sym_1_1_6, tmp19819)

tmp19821 := GCons(__e, let__18073, tmp19820)

tmp19822 := GCons(__e, tmp19821, Nil)

tmp19823 := GCons(__e, sym_1_1_6, tmp19822)

tmp19824 := GCons(__e, let__18072, tmp19823)

tmp19825 := Call6(__e, SymFn(symis_b), V6539, tmp19824, B6535, L6536, Key6537, C6538)


tmp19826 := Call2(__e, SymFn(symshen_4gc), B6535, tmp19825)


return TailCall2(__e, SymFn(symshen_4gc), B6535, tmp19826)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19827 := GValue(__e, symshen_4_dsigf_d)

tmp19828 := Call3(__e, SymFn(symshen_4assoc_1_6), symoccurrences, tmp19816, tmp19827)


tmp19829 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19828)

_ = tmp19829

tmp19830 := MakeNative1(func(__e *ControlFlow, V6544 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6540 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6541 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6542 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6543 Obj) Obj {
tmp19831 := GCons(__e, symboolean, Nil)

tmp19832 := GCons(__e, sym_1_1_6, tmp19831)

tmp19833 := GCons(__e, symsymbol, tmp19832)

return TailCall6(__e, SymFn(symis_b), V6544, tmp19833, B6540, L6541, Key6542, C6543)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19834 := GValue(__e, symshen_4_dsigf_d)

tmp19835 := Call3(__e, SymFn(symshen_4assoc_1_6), symoccurs_1check, tmp19830, tmp19834)


tmp19836 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19835)

_ = tmp19836

tmp19837 := MakeNative1(func(__e *ControlFlow, V6549 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6545 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6546 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6547 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6548 Obj) Obj {
tmp19838 := GCons(__e, symboolean, Nil)

tmp19839 := GCons(__e, sym_1_1_6, tmp19838)

return TailCall6(__e, SymFn(symis_b), V6549, tmp19839, B6545, L6546, Key6547, C6548)


panic("unreachable")
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

tmp19841 := Call3(__e, SymFn(symshen_4assoc_1_6), symoccurs_2, tmp19837, tmp19840)


tmp19842 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19841)

_ = tmp19842

tmp19843 := MakeNative1(func(__e *ControlFlow, V6554 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6550 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6551 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6552 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6553 Obj) Obj {
tmp19844 := GCons(__e, symboolean, Nil)

tmp19845 := GCons(__e, sym_1_1_6, tmp19844)

tmp19846 := GCons(__e, symsymbol, tmp19845)

return TailCall6(__e, SymFn(symis_b), V6554, tmp19846, B6550, L6551, Key6552, C6553)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19847 := GValue(__e, symshen_4_dsigf_d)

tmp19848 := Call3(__e, SymFn(symshen_4assoc_1_6), symoptimise, tmp19843, tmp19847)


tmp19849 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19848)

_ = tmp19849

tmp19850 := MakeNative1(func(__e *ControlFlow, V6559 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6555 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6556 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6557 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6558 Obj) Obj {
tmp19851 := GCons(__e, symboolean, Nil)

tmp19852 := GCons(__e, sym_1_1_6, tmp19851)

return TailCall6(__e, SymFn(symis_b), V6559, tmp19852, B6555, L6556, Key6557, C6558)


panic("unreachable")
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

tmp19854 := Call3(__e, SymFn(symshen_4assoc_1_6), symoptimise_2, tmp19850, tmp19853)


tmp19855 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19854)

_ = tmp19855

tmp19856 := MakeNative1(func(__e *ControlFlow, V6564 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6560 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6561 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6562 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6563 Obj) Obj {
tmp19857 := GCons(__e, symboolean, Nil)

tmp19858 := GCons(__e, sym_1_1_6, tmp19857)

tmp19859 := GCons(__e, symboolean, tmp19858)

tmp19860 := GCons(__e, tmp19859, Nil)

tmp19861 := GCons(__e, sym_1_1_6, tmp19860)

tmp19862 := GCons(__e, symboolean, tmp19861)

return TailCall6(__e, SymFn(symis_b), V6564, tmp19862, B6560, L6561, Key6562, C6563)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19863 := GValue(__e, symshen_4_dsigf_d)

tmp19864 := Call3(__e, SymFn(symshen_4assoc_1_6), symor, tmp19856, tmp19863)


tmp19865 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19864)

_ = tmp19865

tmp19866 := MakeNative1(func(__e *ControlFlow, V6569 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6565 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6566 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6567 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6568 Obj) Obj {
tmp19867 := GCons(__e, symstring, Nil)

tmp19868 := GCons(__e, sym_1_1_6, tmp19867)

return TailCall6(__e, SymFn(symis_b), V6569, tmp19868, B6565, L6566, Key6567, C6568)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19869 := GValue(__e, symshen_4_dsigf_d)

tmp19870 := Call3(__e, SymFn(symshen_4assoc_1_6), symos, tmp19866, tmp19869)


tmp19871 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19870)

_ = tmp19871

tmp19872 := MakeNative1(func(__e *ControlFlow, V6574 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6570 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6571 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6572 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6573 Obj) Obj {
tmp19873 := GCons(__e, symboolean, Nil)

tmp19874 := GCons(__e, sym_1_1_6, tmp19873)

tmp19875 := GCons(__e, symsymbol, tmp19874)

return TailCall6(__e, SymFn(symis_b), V6574, tmp19875, B6570, L6571, Key6572, C6573)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19876 := GValue(__e, symshen_4_dsigf_d)

tmp19877 := Call3(__e, SymFn(symshen_4assoc_1_6), sympackage_2, tmp19872, tmp19876)


tmp19878 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19877)

_ = tmp19878

tmp19879 := MakeNative1(func(__e *ControlFlow, V6579 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6575 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6576 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6577 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6578 Obj) Obj {
tmp19880 := GCons(__e, symstring, Nil)

tmp19881 := GCons(__e, sym_1_1_6, tmp19880)

return TailCall6(__e, SymFn(symis_b), V6579, tmp19881, B6575, L6576, Key6577, C6578)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19882 := GValue(__e, symshen_4_dsigf_d)

tmp19883 := Call3(__e, SymFn(symshen_4assoc_1_6), symport, tmp19879, tmp19882)


tmp19884 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19883)

_ = tmp19884

tmp19885 := MakeNative1(func(__e *ControlFlow, V6584 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6580 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6581 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6582 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6583 Obj) Obj {
tmp19886 := GCons(__e, symstring, Nil)

tmp19887 := GCons(__e, sym_1_1_6, tmp19886)

return TailCall6(__e, SymFn(symis_b), V6584, tmp19887, B6580, L6581, Key6582, C6583)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19888 := GValue(__e, symshen_4_dsigf_d)

tmp19889 := Call3(__e, SymFn(symshen_4assoc_1_6), symporters, tmp19885, tmp19888)


tmp19890 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19889)

_ = tmp19890

tmp19891 := MakeNative1(func(__e *ControlFlow, V6589 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6585 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6586 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6587 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6588 Obj) Obj {
tmp19892 := GCons(__e, symstring, Nil)

tmp19893 := GCons(__e, sym_1_1_6, tmp19892)

tmp19894 := GCons(__e, symnumber, tmp19893)

tmp19895 := GCons(__e, tmp19894, Nil)

tmp19896 := GCons(__e, sym_1_1_6, tmp19895)

tmp19897 := GCons(__e, symstring, tmp19896)

return TailCall6(__e, SymFn(symis_b), V6589, tmp19897, B6585, L6586, Key6587, C6588)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19898 := GValue(__e, symshen_4_dsigf_d)

tmp19899 := Call3(__e, SymFn(symshen_4assoc_1_6), sympos, tmp19891, tmp19898)


tmp19900 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19899)

_ = tmp19900

tmp19901 := MakeNative1(func(__e *ControlFlow, V6594 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6590 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6591 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6592 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6593 Obj) Obj {
tmp19902 := GCons(__e, symout, Nil)

tmp19903 := GCons(__e, symstream, tmp19902)

tmp19904 := GCons(__e, symstring, Nil)

tmp19905 := GCons(__e, sym_1_1_6, tmp19904)

tmp19906 := GCons(__e, tmp19903, tmp19905)

tmp19907 := GCons(__e, tmp19906, Nil)

tmp19908 := GCons(__e, sym_1_1_6, tmp19907)

tmp19909 := GCons(__e, symstring, tmp19908)

return TailCall6(__e, SymFn(symis_b), V6594, tmp19909, B6590, L6591, Key6592, C6593)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19910 := GValue(__e, symshen_4_dsigf_d)

tmp19911 := Call3(__e, SymFn(symshen_4assoc_1_6), sympr, tmp19901, tmp19910)


tmp19912 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19911)

_ = tmp19912

tmp19913 := MakeNative1(func(__e *ControlFlow, V6599 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6595 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6596 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6597 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6598 Obj) Obj {
tmp19914 := Call1(__e, SymFn(symshen_4newpv), B6595)


let__18074 := tmp19914
_ = let__18074

tmp19915 := GCons(__e, let__18074, Nil)

tmp19916 := GCons(__e, sym_1_1_6, tmp19915)

tmp19917 := GCons(__e, let__18074, tmp19916)

tmp19918 := Call6(__e, SymFn(symis_b), V6599, tmp19917, B6595, L6596, Key6597, C6598)


return TailCall2(__e, SymFn(symshen_4gc), B6595, tmp19918)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19919 := GValue(__e, symshen_4_dsigf_d)

tmp19920 := Call3(__e, SymFn(symshen_4assoc_1_6), symprint, tmp19913, tmp19919)


tmp19921 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19920)

_ = tmp19921

tmp19922 := MakeNative1(func(__e *ControlFlow, V6604 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6600 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6601 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6602 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6603 Obj) Obj {
tmp19923 := GCons(__e, symsymbol, Nil)

tmp19924 := GCons(__e, sym_1_1_6, tmp19923)

tmp19925 := GCons(__e, symsymbol, tmp19924)

return TailCall6(__e, SymFn(symis_b), V6604, tmp19925, B6600, L6601, Key6602, C6603)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19926 := GValue(__e, symshen_4_dsigf_d)

tmp19927 := Call3(__e, SymFn(symshen_4assoc_1_6), symprofile, tmp19922, tmp19926)


tmp19928 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19927)

_ = tmp19928

tmp19929 := MakeNative1(func(__e *ControlFlow, V6610 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6606 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6607 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6608 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6609 Obj) Obj {
tmp19930 := GCons(__e, symsymbol, Nil)

tmp19931 := GCons(__e, symlist, tmp19930)

tmp19932 := GCons(__e, symsymbol, Nil)

tmp19933 := GCons(__e, symlist, tmp19932)

tmp19934 := GCons(__e, tmp19933, Nil)

tmp19935 := GCons(__e, sym_1_1_6, tmp19934)

tmp19936 := GCons(__e, tmp19931, tmp19935)

return TailCall6(__e, SymFn(symis_b), V6610, tmp19936, B6606, L6607, Key6608, C6609)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19937 := GValue(__e, symshen_4_dsigf_d)

tmp19938 := Call3(__e, SymFn(symshen_4assoc_1_6), sympreclude, tmp19929, tmp19937)


tmp19939 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19938)

_ = tmp19939

tmp19940 := MakeNative1(func(__e *ControlFlow, V6615 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6611 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6612 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6613 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6614 Obj) Obj {
tmp19941 := GCons(__e, symstring, Nil)

tmp19942 := GCons(__e, sym_1_1_6, tmp19941)

tmp19943 := GCons(__e, symstring, tmp19942)

return TailCall6(__e, SymFn(symis_b), V6615, tmp19943, B6611, L6612, Key6613, C6614)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19944 := GValue(__e, symshen_4_dsigf_d)

tmp19945 := Call3(__e, SymFn(symshen_4assoc_1_6), symshen_4proc_1nl, tmp19940, tmp19944)


tmp19946 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19945)

_ = tmp19946

tmp19947 := MakeNative1(func(__e *ControlFlow, V6620 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6616 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6617 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6618 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6619 Obj) Obj {
tmp19948 := GCons(__e, symnumber, Nil)

tmp19949 := GCons(__e, sym_d, tmp19948)

tmp19950 := GCons(__e, symsymbol, tmp19949)

tmp19951 := GCons(__e, tmp19950, Nil)

tmp19952 := GCons(__e, sym_1_1_6, tmp19951)

tmp19953 := GCons(__e, symsymbol, tmp19952)

return TailCall6(__e, SymFn(symis_b), V6620, tmp19953, B6616, L6617, Key6618, C6619)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19954 := GValue(__e, symshen_4_dsigf_d)

tmp19955 := Call3(__e, SymFn(symshen_4assoc_1_6), symprofile_1results, tmp19947, tmp19954)


tmp19956 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19955)

_ = tmp19956

tmp19957 := MakeNative1(func(__e *ControlFlow, V6625 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6621 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6622 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6623 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6624 Obj) Obj {
tmp19958 := Call1(__e, SymFn(symshen_4newpv), B6621)


let__18075 := tmp19958
_ = let__18075

tmp19959 := GCons(__e, let__18075, Nil)

tmp19960 := GCons(__e, sym_1_1_6, tmp19959)

tmp19961 := GCons(__e, let__18075, tmp19960)

tmp19962 := Call6(__e, SymFn(symis_b), V6625, tmp19961, B6621, L6622, Key6623, C6624)


return TailCall2(__e, SymFn(symshen_4gc), B6621, tmp19962)


panic("unreachable")
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

tmp19964 := Call3(__e, SymFn(symshen_4assoc_1_6), symprotect, tmp19957, tmp19963)


tmp19965 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19964)

_ = tmp19965

tmp19966 := MakeNative1(func(__e *ControlFlow, V6631 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6627 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6628 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6629 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6630 Obj) Obj {
tmp19967 := GCons(__e, symsymbol, Nil)

tmp19968 := GCons(__e, symlist, tmp19967)

tmp19969 := GCons(__e, symsymbol, Nil)

tmp19970 := GCons(__e, symlist, tmp19969)

tmp19971 := GCons(__e, tmp19970, Nil)

tmp19972 := GCons(__e, sym_1_1_6, tmp19971)

tmp19973 := GCons(__e, tmp19968, tmp19972)

return TailCall6(__e, SymFn(symis_b), V6631, tmp19973, B6627, L6628, Key6629, C6630)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19974 := GValue(__e, symshen_4_dsigf_d)

tmp19975 := Call3(__e, SymFn(symshen_4assoc_1_6), sympreclude_1all_1but, tmp19966, tmp19974)


tmp19976 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19975)

_ = tmp19976

tmp19977 := MakeNative1(func(__e *ControlFlow, V6636 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6632 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6633 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6634 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6635 Obj) Obj {
tmp19978 := GCons(__e, symout, Nil)

tmp19979 := GCons(__e, symstream, tmp19978)

tmp19980 := GCons(__e, symstring, Nil)

tmp19981 := GCons(__e, sym_1_1_6, tmp19980)

tmp19982 := GCons(__e, tmp19979, tmp19981)

tmp19983 := GCons(__e, tmp19982, Nil)

tmp19984 := GCons(__e, sym_1_1_6, tmp19983)

tmp19985 := GCons(__e, symstring, tmp19984)

return TailCall6(__e, SymFn(symis_b), V6636, tmp19985, B6632, L6633, Key6634, C6635)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19986 := GValue(__e, symshen_4_dsigf_d)

tmp19987 := Call3(__e, SymFn(symshen_4assoc_1_6), symshen_4prhush, tmp19977, tmp19986)


tmp19988 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19987)

_ = tmp19988

tmp19989 := MakeNative1(func(__e *ControlFlow, V6641 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6637 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6638 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6639 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6640 Obj) Obj {
tmp19990 := GCons(__e, symnumber, Nil)

tmp19991 := GCons(__e, sym_1_1_6, tmp19990)

tmp19992 := GCons(__e, symnumber, tmp19991)

return TailCall6(__e, SymFn(symis_b), V6641, tmp19992, B6637, L6638, Key6639, C6640)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp19993 := GValue(__e, symshen_4_dsigf_d)

tmp19994 := Call3(__e, SymFn(symshen_4assoc_1_6), symprolog_1memory, tmp19989, tmp19993)


tmp19995 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp19994)

_ = tmp19995

tmp19996 := MakeNative1(func(__e *ControlFlow, V6646 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6642 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6643 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6644 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6645 Obj) Obj {
tmp19997 := GCons(__e, symunit, Nil)

tmp19998 := GCons(__e, symlist, tmp19997)

tmp19999 := GCons(__e, tmp19998, Nil)

tmp20000 := GCons(__e, sym_1_1_6, tmp19999)

tmp20001 := GCons(__e, symsymbol, tmp20000)

return TailCall6(__e, SymFn(symis_b), V6646, tmp20001, B6642, L6643, Key6644, C6645)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20002 := GValue(__e, symshen_4_dsigf_d)

tmp20003 := Call3(__e, SymFn(symshen_4assoc_1_6), symps, tmp19996, tmp20002)


tmp20004 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20003)

_ = tmp20004

tmp20005 := MakeNative1(func(__e *ControlFlow, V6651 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6647 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6648 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6649 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6650 Obj) Obj {
tmp20006 := GCons(__e, symin, Nil)

tmp20007 := GCons(__e, symstream, tmp20006)

tmp20008 := GCons(__e, symunit, Nil)

tmp20009 := GCons(__e, sym_1_1_6, tmp20008)

tmp20010 := GCons(__e, tmp20007, tmp20009)

return TailCall6(__e, SymFn(symis_b), V6651, tmp20010, B6647, L6648, Key6649, C6650)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20011 := GValue(__e, symshen_4_dsigf_d)

tmp20012 := Call3(__e, SymFn(symshen_4assoc_1_6), symread, tmp20005, tmp20011)


tmp20013 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20012)

_ = tmp20013

tmp20014 := MakeNative1(func(__e *ControlFlow, V6656 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6652 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6653 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6654 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6655 Obj) Obj {
tmp20015 := GCons(__e, symin, Nil)

tmp20016 := GCons(__e, symstream, tmp20015)

tmp20017 := GCons(__e, symstring, Nil)

tmp20018 := GCons(__e, sym_1_1_6, tmp20017)

tmp20019 := GCons(__e, tmp20016, tmp20018)

return TailCall6(__e, SymFn(symis_b), V6656, tmp20019, B6652, L6653, Key6654, C6655)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20020 := GValue(__e, symshen_4_dsigf_d)

tmp20021 := Call3(__e, SymFn(symshen_4assoc_1_6), symread_1unit_1string, tmp20014, tmp20020)


tmp20022 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20021)

_ = tmp20022

tmp20023 := MakeNative1(func(__e *ControlFlow, V6661 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6657 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6658 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6659 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6660 Obj) Obj {
tmp20024 := GCons(__e, symin, Nil)

tmp20025 := GCons(__e, symstream, tmp20024)

tmp20026 := GCons(__e, symnumber, Nil)

tmp20027 := GCons(__e, sym_1_1_6, tmp20026)

tmp20028 := GCons(__e, tmp20025, tmp20027)

return TailCall6(__e, SymFn(symis_b), V6661, tmp20028, B6657, L6658, Key6659, C6660)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20029 := GValue(__e, symshen_4_dsigf_d)

tmp20030 := Call3(__e, SymFn(symshen_4assoc_1_6), symread_1byte, tmp20023, tmp20029)


tmp20031 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20030)

_ = tmp20031

tmp20032 := MakeNative1(func(__e *ControlFlow, V6666 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6662 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6663 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6664 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6665 Obj) Obj {
tmp20033 := GCons(__e, symnumber, Nil)

tmp20034 := GCons(__e, symlist, tmp20033)

tmp20035 := GCons(__e, tmp20034, Nil)

tmp20036 := GCons(__e, sym_1_1_6, tmp20035)

tmp20037 := GCons(__e, symstring, tmp20036)

return TailCall6(__e, SymFn(symis_b), V6666, tmp20037, B6662, L6663, Key6664, C6665)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20038 := GValue(__e, symshen_4_dsigf_d)

tmp20039 := Call3(__e, SymFn(symshen_4assoc_1_6), symread_1file_1as_1bytelist, tmp20032, tmp20038)


tmp20040 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20039)

_ = tmp20040

tmp20041 := MakeNative1(func(__e *ControlFlow, V6671 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6667 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6668 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6669 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6670 Obj) Obj {
tmp20042 := GCons(__e, symstring, Nil)

tmp20043 := GCons(__e, sym_1_1_6, tmp20042)

tmp20044 := GCons(__e, symstring, tmp20043)

return TailCall6(__e, SymFn(symis_b), V6671, tmp20044, B6667, L6668, Key6669, C6670)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20045 := GValue(__e, symshen_4_dsigf_d)

tmp20046 := Call3(__e, SymFn(symshen_4assoc_1_6), symread_1file_1as_1string, tmp20041, tmp20045)


tmp20047 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20046)

_ = tmp20047

tmp20048 := MakeNative1(func(__e *ControlFlow, V6676 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6672 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6673 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6674 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6675 Obj) Obj {
tmp20049 := GCons(__e, symunit, Nil)

tmp20050 := GCons(__e, symlist, tmp20049)

tmp20051 := GCons(__e, tmp20050, Nil)

tmp20052 := GCons(__e, sym_1_1_6, tmp20051)

tmp20053 := GCons(__e, symstring, tmp20052)

return TailCall6(__e, SymFn(symis_b), V6676, tmp20053, B6672, L6673, Key6674, C6675)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20054 := GValue(__e, symshen_4_dsigf_d)

tmp20055 := Call3(__e, SymFn(symshen_4assoc_1_6), symread_1file, tmp20048, tmp20054)


tmp20056 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20055)

_ = tmp20056

tmp20057 := MakeNative1(func(__e *ControlFlow, V6681 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6677 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6678 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6679 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6680 Obj) Obj {
tmp20058 := GCons(__e, symunit, Nil)

tmp20059 := GCons(__e, symlist, tmp20058)

tmp20060 := GCons(__e, tmp20059, Nil)

tmp20061 := GCons(__e, sym_1_1_6, tmp20060)

tmp20062 := GCons(__e, symstring, tmp20061)

return TailCall6(__e, SymFn(symis_b), V6681, tmp20062, B6677, L6678, Key6679, C6680)


panic("unreachable")
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

tmp20064 := Call3(__e, SymFn(symshen_4assoc_1_6), symread_1from_1string, tmp20057, tmp20063)


tmp20065 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20064)

_ = tmp20065

tmp20066 := MakeNative1(func(__e *ControlFlow, V6686 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6682 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6683 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6684 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6685 Obj) Obj {
tmp20067 := GCons(__e, symunit, Nil)

tmp20068 := GCons(__e, symlist, tmp20067)

tmp20069 := GCons(__e, tmp20068, Nil)

tmp20070 := GCons(__e, sym_1_1_6, tmp20069)

tmp20071 := GCons(__e, symstring, tmp20070)

return TailCall6(__e, SymFn(symis_b), V6686, tmp20071, B6682, L6683, Key6684, C6685)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20072 := GValue(__e, symshen_4_dsigf_d)

tmp20073 := Call3(__e, SymFn(symshen_4assoc_1_6), symread_1from_1string_1unprocessed, tmp20066, tmp20072)


tmp20074 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20073)

_ = tmp20074

tmp20075 := MakeNative1(func(__e *ControlFlow, V6691 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6687 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6688 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6689 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6690 Obj) Obj {
tmp20076 := GCons(__e, symstring, Nil)

tmp20077 := GCons(__e, sym_1_1_6, tmp20076)

return TailCall6(__e, SymFn(symis_b), V6691, tmp20077, B6687, L6688, Key6689, C6690)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20078 := GValue(__e, symshen_4_dsigf_d)

tmp20079 := Call3(__e, SymFn(symshen_4assoc_1_6), symrelease, tmp20075, tmp20078)


tmp20080 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20079)

_ = tmp20080

tmp20081 := MakeNative1(func(__e *ControlFlow, V6697 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6693 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6694 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6695 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6696 Obj) Obj {
tmp20082 := Call1(__e, SymFn(symshen_4newpv), B6693)


let__18076 := tmp20082
_ = let__18076

tmp20083 := GCons(__e, let__18076, Nil)

tmp20084 := GCons(__e, symlist, tmp20083)

tmp20085 := GCons(__e, let__18076, Nil)

tmp20086 := GCons(__e, symlist, tmp20085)

tmp20087 := GCons(__e, tmp20086, Nil)

tmp20088 := GCons(__e, sym_1_1_6, tmp20087)

tmp20089 := GCons(__e, tmp20084, tmp20088)

tmp20090 := GCons(__e, tmp20089, Nil)

tmp20091 := GCons(__e, sym_1_1_6, tmp20090)

tmp20092 := GCons(__e, let__18076, tmp20091)

tmp20093 := Call6(__e, SymFn(symis_b), V6697, tmp20092, B6693, L6694, Key6695, C6696)


return TailCall2(__e, SymFn(symshen_4gc), B6693, tmp20093)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20094 := GValue(__e, symshen_4_dsigf_d)

tmp20095 := Call3(__e, SymFn(symshen_4assoc_1_6), symremove, tmp20081, tmp20094)


tmp20096 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20095)

_ = tmp20096

tmp20097 := MakeNative1(func(__e *ControlFlow, V6703 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6699 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6700 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6701 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6702 Obj) Obj {
tmp20098 := Call1(__e, SymFn(symshen_4newpv), B6699)


let__18077 := tmp20098
_ = let__18077

tmp20099 := GCons(__e, let__18077, Nil)

tmp20100 := GCons(__e, symlist, tmp20099)

tmp20101 := GCons(__e, let__18077, Nil)

tmp20102 := GCons(__e, symlist, tmp20101)

tmp20103 := GCons(__e, tmp20102, Nil)

tmp20104 := GCons(__e, sym_1_1_6, tmp20103)

tmp20105 := GCons(__e, tmp20100, tmp20104)

tmp20106 := Call6(__e, SymFn(symis_b), V6703, tmp20105, B6699, L6700, Key6701, C6702)


return TailCall2(__e, SymFn(symshen_4gc), B6699, tmp20106)


panic("unreachable")
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

tmp20108 := Call3(__e, SymFn(symshen_4assoc_1_6), symreverse, tmp20097, tmp20107)


tmp20109 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20108)

_ = tmp20109

tmp20110 := MakeNative1(func(__e *ControlFlow, V6708 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6704 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6705 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6706 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6707 Obj) Obj {
tmp20111 := Call1(__e, SymFn(symshen_4newpv), B6704)


let__18078 := tmp20111
_ = let__18078

tmp20112 := GCons(__e, let__18078, Nil)

tmp20113 := GCons(__e, sym_1_1_6, tmp20112)

tmp20114 := GCons(__e, symstring, tmp20113)

tmp20115 := Call6(__e, SymFn(symis_b), V6708, tmp20114, B6704, L6705, Key6706, C6707)


return TailCall2(__e, SymFn(symshen_4gc), B6704, tmp20115)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20116 := GValue(__e, symshen_4_dsigf_d)

tmp20117 := Call3(__e, SymFn(symshen_4assoc_1_6), symsimple_1error, tmp20110, tmp20116)


tmp20118 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20117)

_ = tmp20118

tmp20119 := MakeNative1(func(__e *ControlFlow, V6713 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6709 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6710 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6711 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6712 Obj) Obj {
tmp20120 := Call1(__e, SymFn(symshen_4newpv), B6709)


let__18079 := tmp20120
_ = let__18079

tmp20121 := Call1(__e, SymFn(symshen_4newpv), B6709)


let__18080 := tmp20121
_ = let__18080

tmp20122 := GCons(__e, let__18080, Nil)

tmp20123 := GCons(__e, sym_d, tmp20122)

tmp20124 := GCons(__e, let__18079, tmp20123)

tmp20125 := GCons(__e, let__18080, Nil)

tmp20126 := GCons(__e, sym_1_1_6, tmp20125)

tmp20127 := GCons(__e, tmp20124, tmp20126)

tmp20128 := Call6(__e, SymFn(symis_b), V6713, tmp20127, B6709, L6710, Key6711, C6712)


tmp20129 := Call2(__e, SymFn(symshen_4gc), B6709, tmp20128)


return TailCall2(__e, SymFn(symshen_4gc), B6709, tmp20129)


panic("unreachable")
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

tmp20131 := Call3(__e, SymFn(symshen_4assoc_1_6), symsnd, tmp20119, tmp20130)


tmp20132 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20131)

_ = tmp20132

tmp20133 := MakeNative1(func(__e *ControlFlow, V6718 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6714 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6715 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6716 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6717 Obj) Obj {
tmp20134 := GCons(__e, symsymbol, Nil)

tmp20135 := GCons(__e, sym_1_1_6, tmp20134)

tmp20136 := GCons(__e, symnumber, tmp20135)

tmp20137 := GCons(__e, tmp20136, Nil)

tmp20138 := GCons(__e, sym_1_1_6, tmp20137)

tmp20139 := GCons(__e, symsymbol, tmp20138)

return TailCall6(__e, SymFn(symis_b), V6718, tmp20139, B6714, L6715, Key6716, C6717)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20140 := GValue(__e, symshen_4_dsigf_d)

tmp20141 := Call3(__e, SymFn(symshen_4assoc_1_6), symspecialise, tmp20133, tmp20140)


tmp20142 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20141)

_ = tmp20142

tmp20143 := MakeNative1(func(__e *ControlFlow, V6723 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6719 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6720 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6721 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6722 Obj) Obj {
tmp20144 := GCons(__e, symboolean, Nil)

tmp20145 := GCons(__e, sym_1_1_6, tmp20144)

tmp20146 := GCons(__e, symsymbol, tmp20145)

return TailCall6(__e, SymFn(symis_b), V6723, tmp20146, B6719, L6720, Key6721, C6722)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20147 := GValue(__e, symshen_4_dsigf_d)

tmp20148 := Call3(__e, SymFn(symshen_4assoc_1_6), symspy, tmp20143, tmp20147)


tmp20149 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20148)

_ = tmp20149

tmp20150 := MakeNative1(func(__e *ControlFlow, V6728 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6724 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6725 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6726 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6727 Obj) Obj {
tmp20151 := GCons(__e, symboolean, Nil)

tmp20152 := GCons(__e, sym_1_1_6, tmp20151)

return TailCall6(__e, SymFn(symis_b), V6728, tmp20152, B6724, L6725, Key6726, C6727)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20153 := GValue(__e, symshen_4_dsigf_d)

tmp20154 := Call3(__e, SymFn(symshen_4assoc_1_6), symspy_2, tmp20150, tmp20153)


tmp20155 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20154)

_ = tmp20155

tmp20156 := MakeNative1(func(__e *ControlFlow, V6733 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6729 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6730 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6731 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6732 Obj) Obj {
tmp20157 := GCons(__e, symboolean, Nil)

tmp20158 := GCons(__e, sym_1_1_6, tmp20157)

tmp20159 := GCons(__e, symsymbol, tmp20158)

return TailCall6(__e, SymFn(symis_b), V6733, tmp20159, B6729, L6730, Key6731, C6732)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20160 := GValue(__e, symshen_4_dsigf_d)

tmp20161 := Call3(__e, SymFn(symshen_4assoc_1_6), symstep, tmp20156, tmp20160)


tmp20162 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20161)

_ = tmp20162

tmp20163 := MakeNative1(func(__e *ControlFlow, V6738 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6734 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6735 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6736 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6737 Obj) Obj {
tmp20164 := GCons(__e, symboolean, Nil)

tmp20165 := GCons(__e, sym_1_1_6, tmp20164)

return TailCall6(__e, SymFn(symis_b), V6738, tmp20165, B6734, L6735, Key6736, C6737)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20166 := GValue(__e, symshen_4_dsigf_d)

tmp20167 := Call3(__e, SymFn(symshen_4assoc_1_6), symstep_2, tmp20163, tmp20166)


tmp20168 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20167)

_ = tmp20168

tmp20169 := MakeNative1(func(__e *ControlFlow, V6743 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6739 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6740 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6741 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6742 Obj) Obj {
tmp20170 := GCons(__e, symin, Nil)

tmp20171 := GCons(__e, symstream, tmp20170)

tmp20172 := GCons(__e, tmp20171, Nil)

tmp20173 := GCons(__e, sym_1_1_6, tmp20172)

return TailCall6(__e, SymFn(symis_b), V6743, tmp20173, B6739, L6740, Key6741, C6742)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20174 := GValue(__e, symshen_4_dsigf_d)

tmp20175 := Call3(__e, SymFn(symshen_4assoc_1_6), symstinput, tmp20169, tmp20174)


tmp20176 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20175)

_ = tmp20176

tmp20177 := MakeNative1(func(__e *ControlFlow, V6748 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6744 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6745 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6746 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6747 Obj) Obj {
tmp20178 := GCons(__e, symout, Nil)

tmp20179 := GCons(__e, symstream, tmp20178)

tmp20180 := GCons(__e, tmp20179, Nil)

tmp20181 := GCons(__e, sym_1_1_6, tmp20180)

return TailCall6(__e, SymFn(symis_b), V6748, tmp20181, B6744, L6745, Key6746, C6747)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20182 := GValue(__e, symshen_4_dsigf_d)

tmp20183 := Call3(__e, SymFn(symshen_4assoc_1_6), symsterror, tmp20177, tmp20182)


tmp20184 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20183)

_ = tmp20184

tmp20185 := MakeNative1(func(__e *ControlFlow, V6753 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6749 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6750 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6751 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6752 Obj) Obj {
tmp20186 := GCons(__e, symout, Nil)

tmp20187 := GCons(__e, symstream, tmp20186)

tmp20188 := GCons(__e, tmp20187, Nil)

tmp20189 := GCons(__e, sym_1_1_6, tmp20188)

return TailCall6(__e, SymFn(symis_b), V6753, tmp20189, B6749, L6750, Key6751, C6752)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20190 := GValue(__e, symshen_4_dsigf_d)

tmp20191 := Call3(__e, SymFn(symshen_4assoc_1_6), symstoutput, tmp20185, tmp20190)


tmp20192 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20191)

_ = tmp20192

tmp20193 := MakeNative1(func(__e *ControlFlow, V6758 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6754 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6755 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6756 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6757 Obj) Obj {
tmp20194 := Call1(__e, SymFn(symshen_4newpv), B6754)


let__18081 := tmp20194
_ = let__18081

tmp20195 := GCons(__e, symboolean, Nil)

tmp20196 := GCons(__e, sym_1_1_6, tmp20195)

tmp20197 := GCons(__e, let__18081, tmp20196)

tmp20198 := Call6(__e, SymFn(symis_b), V6758, tmp20197, B6754, L6755, Key6756, C6757)


return TailCall2(__e, SymFn(symshen_4gc), B6754, tmp20198)


panic("unreachable")
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

tmp20200 := Call3(__e, SymFn(symshen_4assoc_1_6), symstring_2, tmp20193, tmp20199)


tmp20201 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20200)

_ = tmp20201

tmp20202 := MakeNative1(func(__e *ControlFlow, V6763 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6759 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6760 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6761 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6762 Obj) Obj {
tmp20203 := Call1(__e, SymFn(symshen_4newpv), B6759)


let__18082 := tmp20203
_ = let__18082

tmp20204 := GCons(__e, symstring, Nil)

tmp20205 := GCons(__e, sym_1_1_6, tmp20204)

tmp20206 := GCons(__e, let__18082, tmp20205)

tmp20207 := Call6(__e, SymFn(symis_b), V6763, tmp20206, B6759, L6760, Key6761, C6762)


return TailCall2(__e, SymFn(symshen_4gc), B6759, tmp20207)


panic("unreachable")
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

tmp20209 := Call3(__e, SymFn(symshen_4assoc_1_6), symstr, tmp20202, tmp20208)


tmp20210 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20209)

_ = tmp20210

tmp20211 := MakeNative1(func(__e *ControlFlow, V6768 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6764 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6765 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6766 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6767 Obj) Obj {
tmp20212 := GCons(__e, symnumber, Nil)

tmp20213 := GCons(__e, sym_1_1_6, tmp20212)

tmp20214 := GCons(__e, symstring, tmp20213)

return TailCall6(__e, SymFn(symis_b), V6768, tmp20214, B6764, L6765, Key6766, C6767)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20215 := GValue(__e, symshen_4_dsigf_d)

tmp20216 := Call3(__e, SymFn(symshen_4assoc_1_6), symstring_1_6n, tmp20211, tmp20215)


tmp20217 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20216)

_ = tmp20217

tmp20218 := MakeNative1(func(__e *ControlFlow, V6773 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6769 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6770 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6771 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6772 Obj) Obj {
tmp20219 := GCons(__e, symsymbol, Nil)

tmp20220 := GCons(__e, sym_1_1_6, tmp20219)

tmp20221 := GCons(__e, symstring, tmp20220)

return TailCall6(__e, SymFn(symis_b), V6773, tmp20221, B6769, L6770, Key6771, C6772)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20222 := GValue(__e, symshen_4_dsigf_d)

tmp20223 := Call3(__e, SymFn(symshen_4assoc_1_6), symstring_1_6symbol, tmp20218, tmp20222)


tmp20224 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20223)

_ = tmp20224

tmp20225 := MakeNative1(func(__e *ControlFlow, V6778 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6774 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6775 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6776 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6777 Obj) Obj {
tmp20226 := GCons(__e, symnumber, Nil)

tmp20227 := GCons(__e, symlist, tmp20226)

tmp20228 := GCons(__e, symnumber, Nil)

tmp20229 := GCons(__e, sym_1_1_6, tmp20228)

tmp20230 := GCons(__e, tmp20227, tmp20229)

return TailCall6(__e, SymFn(symis_b), V6778, tmp20230, B6774, L6775, Key6776, C6777)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20231 := GValue(__e, symshen_4_dsigf_d)

tmp20232 := Call3(__e, SymFn(symshen_4assoc_1_6), symsum, tmp20225, tmp20231)


tmp20233 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20232)

_ = tmp20233

tmp20234 := MakeNative1(func(__e *ControlFlow, V6783 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6779 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6780 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6781 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6782 Obj) Obj {
tmp20235 := Call1(__e, SymFn(symshen_4newpv), B6779)


let__18083 := tmp20235
_ = let__18083

tmp20236 := GCons(__e, symboolean, Nil)

tmp20237 := GCons(__e, sym_1_1_6, tmp20236)

tmp20238 := GCons(__e, let__18083, tmp20237)

tmp20239 := Call6(__e, SymFn(symis_b), V6783, tmp20238, B6779, L6780, Key6781, C6782)


return TailCall2(__e, SymFn(symshen_4gc), B6779, tmp20239)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20240 := GValue(__e, symshen_4_dsigf_d)

tmp20241 := Call3(__e, SymFn(symshen_4assoc_1_6), symsymbol_2, tmp20234, tmp20240)


tmp20242 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20241)

_ = tmp20242

tmp20243 := MakeNative1(func(__e *ControlFlow, V6788 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6784 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6785 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6786 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6787 Obj) Obj {
tmp20244 := GCons(__e, symsymbol, Nil)

tmp20245 := GCons(__e, sym_1_1_6, tmp20244)

tmp20246 := GCons(__e, symsymbol, tmp20245)

return TailCall6(__e, SymFn(symis_b), V6788, tmp20246, B6784, L6785, Key6786, C6787)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20247 := GValue(__e, symshen_4_dsigf_d)

tmp20248 := Call3(__e, SymFn(symshen_4assoc_1_6), symsystemf, tmp20243, tmp20247)


tmp20249 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20248)

_ = tmp20249

tmp20250 := MakeNative1(func(__e *ControlFlow, V6793 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6789 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6790 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6791 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6792 Obj) Obj {
tmp20251 := GCons(__e, symboolean, Nil)

tmp20252 := GCons(__e, sym_1_1_6, tmp20251)

return TailCall6(__e, SymFn(symis_b), V6793, tmp20252, B6789, L6790, Key6791, C6792)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20253 := GValue(__e, symshen_4_dsigf_d)

tmp20254 := Call3(__e, SymFn(symshen_4assoc_1_6), symsystem_1S_2, tmp20250, tmp20253)


tmp20255 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20254)

_ = tmp20255

tmp20256 := MakeNative1(func(__e *ControlFlow, V6799 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6795 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6796 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6797 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6798 Obj) Obj {
tmp20257 := Call1(__e, SymFn(symshen_4newpv), B6795)


let__18084 := tmp20257
_ = let__18084

tmp20258 := GCons(__e, let__18084, Nil)

tmp20259 := GCons(__e, symlist, tmp20258)

tmp20260 := GCons(__e, let__18084, Nil)

tmp20261 := GCons(__e, symlist, tmp20260)

tmp20262 := GCons(__e, tmp20261, Nil)

tmp20263 := GCons(__e, sym_1_1_6, tmp20262)

tmp20264 := GCons(__e, tmp20259, tmp20263)

tmp20265 := Call6(__e, SymFn(symis_b), V6799, tmp20264, B6795, L6796, Key6797, C6798)


return TailCall2(__e, SymFn(symshen_4gc), B6795, tmp20265)


panic("unreachable")
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

tmp20267 := Call3(__e, SymFn(symshen_4assoc_1_6), symtail, tmp20256, tmp20266)


tmp20268 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20267)

_ = tmp20268

tmp20269 := MakeNative1(func(__e *ControlFlow, V6804 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6800 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6801 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6802 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6803 Obj) Obj {
tmp20270 := GCons(__e, symstring, Nil)

tmp20271 := GCons(__e, sym_1_1_6, tmp20270)

tmp20272 := GCons(__e, symstring, tmp20271)

return TailCall6(__e, SymFn(symis_b), V6804, tmp20272, B6800, L6801, Key6802, C6803)


panic("unreachable")
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

tmp20274 := Call3(__e, SymFn(symshen_4assoc_1_6), symtlstr, tmp20269, tmp20273)


tmp20275 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20274)

_ = tmp20275

tmp20276 := MakeNative1(func(__e *ControlFlow, V6810 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6806 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6807 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6808 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6809 Obj) Obj {
tmp20277 := Call1(__e, SymFn(symshen_4newpv), B6806)


let__18085 := tmp20277
_ = let__18085

tmp20278 := GCons(__e, let__18085, Nil)

tmp20279 := GCons(__e, symvector, tmp20278)

tmp20280 := GCons(__e, let__18085, Nil)

tmp20281 := GCons(__e, symvector, tmp20280)

tmp20282 := GCons(__e, tmp20281, Nil)

tmp20283 := GCons(__e, sym_1_1_6, tmp20282)

tmp20284 := GCons(__e, tmp20279, tmp20283)

tmp20285 := Call6(__e, SymFn(symis_b), V6810, tmp20284, B6806, L6807, Key6808, C6809)


return TailCall2(__e, SymFn(symshen_4gc), B6806, tmp20285)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20286 := GValue(__e, symshen_4_dsigf_d)

tmp20287 := Call3(__e, SymFn(symshen_4assoc_1_6), symtlv, tmp20276, tmp20286)


tmp20288 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20287)

_ = tmp20288

tmp20289 := MakeNative1(func(__e *ControlFlow, V6815 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6811 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6812 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6813 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6814 Obj) Obj {
tmp20290 := GCons(__e, symboolean, Nil)

tmp20291 := GCons(__e, sym_1_1_6, tmp20290)

tmp20292 := GCons(__e, symsymbol, tmp20291)

return TailCall6(__e, SymFn(symis_b), V6815, tmp20292, B6811, L6812, Key6813, C6814)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20293 := GValue(__e, symshen_4_dsigf_d)

tmp20294 := Call3(__e, SymFn(symshen_4assoc_1_6), symtc, tmp20289, tmp20293)


tmp20295 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20294)

_ = tmp20295

tmp20296 := MakeNative1(func(__e *ControlFlow, V6820 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6816 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6817 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6818 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6819 Obj) Obj {
tmp20297 := GCons(__e, symboolean, Nil)

tmp20298 := GCons(__e, sym_1_1_6, tmp20297)

return TailCall6(__e, SymFn(symis_b), V6820, tmp20298, B6816, L6817, Key6818, C6819)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20299 := GValue(__e, symshen_4_dsigf_d)

tmp20300 := Call3(__e, SymFn(symshen_4assoc_1_6), symtc_2, tmp20296, tmp20299)


tmp20301 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20300)

_ = tmp20301

tmp20302 := MakeNative1(func(__e *ControlFlow, V6825 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6821 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6822 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6823 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6824 Obj) Obj {
tmp20303 := Call1(__e, SymFn(symshen_4newpv), B6821)


let__18086 := tmp20303
_ = let__18086

tmp20304 := GCons(__e, let__18086, Nil)

tmp20305 := GCons(__e, symlazy, tmp20304)

tmp20306 := GCons(__e, let__18086, Nil)

tmp20307 := GCons(__e, sym_1_1_6, tmp20306)

tmp20308 := GCons(__e, tmp20305, tmp20307)

tmp20309 := Call6(__e, SymFn(symis_b), V6825, tmp20308, B6821, L6822, Key6823, C6824)


return TailCall2(__e, SymFn(symshen_4gc), B6821, tmp20309)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20310 := GValue(__e, symshen_4_dsigf_d)

tmp20311 := Call3(__e, SymFn(symshen_4assoc_1_6), symthaw, tmp20302, tmp20310)


tmp20312 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20311)

_ = tmp20312

tmp20313 := MakeNative1(func(__e *ControlFlow, V6830 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6826 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6827 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6828 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6829 Obj) Obj {
tmp20314 := GCons(__e, symsymbol, Nil)

tmp20315 := GCons(__e, sym_1_1_6, tmp20314)

tmp20316 := GCons(__e, symsymbol, tmp20315)

return TailCall6(__e, SymFn(symis_b), V6830, tmp20316, B6826, L6827, Key6828, C6829)


panic("unreachable")
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

tmp20318 := Call3(__e, SymFn(symshen_4assoc_1_6), symtrack, tmp20313, tmp20317)


tmp20319 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20318)

_ = tmp20319

tmp20320 := MakeNative1(func(__e *ControlFlow, V6835 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6831 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6832 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6833 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6834 Obj) Obj {
tmp20321 := GCons(__e, symsymbol, Nil)

tmp20322 := GCons(__e, symlist, tmp20321)

tmp20323 := GCons(__e, tmp20322, Nil)

tmp20324 := GCons(__e, sym_1_1_6, tmp20323)

return TailCall6(__e, SymFn(symis_b), V6835, tmp20324, B6831, L6832, Key6833, C6834)


panic("unreachable")
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

tmp20326 := Call3(__e, SymFn(symshen_4assoc_1_6), symtracked, tmp20320, tmp20325)


tmp20327 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20326)

_ = tmp20327

tmp20328 := MakeNative1(func(__e *ControlFlow, V6840 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6836 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6837 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6838 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6839 Obj) Obj {
tmp20329 := Call1(__e, SymFn(symshen_4newpv), B6836)


let__18087 := tmp20329
_ = let__18087

tmp20330 := GCons(__e, let__18087, Nil)

tmp20331 := GCons(__e, sym_1_1_6, tmp20330)

tmp20332 := GCons(__e, symexception, tmp20331)

tmp20333 := GCons(__e, let__18087, Nil)

tmp20334 := GCons(__e, sym_1_1_6, tmp20333)

tmp20335 := GCons(__e, tmp20332, tmp20334)

tmp20336 := GCons(__e, tmp20335, Nil)

tmp20337 := GCons(__e, sym_1_1_6, tmp20336)

tmp20338 := GCons(__e, let__18087, tmp20337)

tmp20339 := Call6(__e, SymFn(symis_b), V6840, tmp20338, B6836, L6837, Key6838, C6839)


return TailCall2(__e, SymFn(symshen_4gc), B6836, tmp20339)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20340 := GValue(__e, symshen_4_dsigf_d)

tmp20341 := Call3(__e, SymFn(symshen_4assoc_1_6), symtrap_1error, tmp20328, tmp20340)


tmp20342 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20341)

_ = tmp20342

tmp20343 := MakeNative1(func(__e *ControlFlow, V6845 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6841 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6842 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6843 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6844 Obj) Obj {
tmp20344 := Call1(__e, SymFn(symshen_4newpv), B6841)


let__18088 := tmp20344
_ = let__18088

tmp20345 := GCons(__e, symboolean, Nil)

tmp20346 := GCons(__e, sym_1_1_6, tmp20345)

tmp20347 := GCons(__e, let__18088, tmp20346)

tmp20348 := Call6(__e, SymFn(symis_b), V6845, tmp20347, B6841, L6842, Key6843, C6844)


return TailCall2(__e, SymFn(symshen_4gc), B6841, tmp20348)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20349 := GValue(__e, symshen_4_dsigf_d)

tmp20350 := Call3(__e, SymFn(symshen_4assoc_1_6), symtuple_2, tmp20343, tmp20349)


tmp20351 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20350)

_ = tmp20351

tmp20352 := MakeNative1(func(__e *ControlFlow, V6850 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6846 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6847 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6848 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6849 Obj) Obj {
tmp20353 := GCons(__e, symstring, Nil)

tmp20354 := GCons(__e, symlist, tmp20353)

tmp20355 := GCons(__e, tmp20354, Nil)

tmp20356 := GCons(__e, sym_1_1_6, tmp20355)

tmp20357 := GCons(__e, symstring, tmp20356)

return TailCall6(__e, SymFn(symis_b), V6850, tmp20357, B6846, L6847, Key6848, C6849)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20358 := GValue(__e, symshen_4_dsigf_d)

tmp20359 := Call3(__e, SymFn(symshen_4assoc_1_6), symunabsolute, tmp20352, tmp20358)


tmp20360 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20359)

_ = tmp20360

tmp20361 := MakeNative1(func(__e *ControlFlow, V6855 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6851 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6852 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6853 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6854 Obj) Obj {
tmp20362 := GCons(__e, symsymbol, Nil)

tmp20363 := GCons(__e, sym_1_1_6, tmp20362)

tmp20364 := GCons(__e, symsymbol, tmp20363)

return TailCall6(__e, SymFn(symis_b), V6855, tmp20364, B6851, L6852, Key6853, C6854)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20365 := GValue(__e, symshen_4_dsigf_d)

tmp20366 := Call3(__e, SymFn(symshen_4assoc_1_6), symundefmacro, tmp20361, tmp20365)


tmp20367 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20366)

_ = tmp20367

tmp20368 := MakeNative1(func(__e *ControlFlow, V6862 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6858 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6859 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6860 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6861 Obj) Obj {
tmp20369 := Call1(__e, SymFn(symshen_4newpv), B6858)


let__18089 := tmp20369
_ = let__18089

tmp20370 := GCons(__e, let__18089, Nil)

tmp20371 := GCons(__e, symlist, tmp20370)

tmp20372 := GCons(__e, let__18089, Nil)

tmp20373 := GCons(__e, symlist, tmp20372)

tmp20374 := GCons(__e, let__18089, Nil)

tmp20375 := GCons(__e, symlist, tmp20374)

tmp20376 := GCons(__e, tmp20375, Nil)

tmp20377 := GCons(__e, sym_1_1_6, tmp20376)

tmp20378 := GCons(__e, tmp20373, tmp20377)

tmp20379 := GCons(__e, tmp20378, Nil)

tmp20380 := GCons(__e, sym_1_1_6, tmp20379)

tmp20381 := GCons(__e, tmp20371, tmp20380)

tmp20382 := Call6(__e, SymFn(symis_b), V6862, tmp20381, B6858, L6859, Key6860, C6861)


return TailCall2(__e, SymFn(symshen_4gc), B6858, tmp20382)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20383 := GValue(__e, symshen_4_dsigf_d)

tmp20384 := Call3(__e, SymFn(symshen_4assoc_1_6), symunion, tmp20368, tmp20383)


tmp20385 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20384)

_ = tmp20385

tmp20386 := MakeNative1(func(__e *ControlFlow, V6867 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6863 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6864 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6865 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6866 Obj) Obj {
tmp20387 := GCons(__e, symsymbol, Nil)

tmp20388 := GCons(__e, sym_1_1_6, tmp20387)

tmp20389 := GCons(__e, symsymbol, tmp20388)

return TailCall6(__e, SymFn(symis_b), V6867, tmp20389, B6863, L6864, Key6865, C6866)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20390 := GValue(__e, symshen_4_dsigf_d)

tmp20391 := Call3(__e, SymFn(symshen_4assoc_1_6), symunprofile, tmp20386, tmp20390)


tmp20392 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20391)

_ = tmp20392

tmp20393 := MakeNative1(func(__e *ControlFlow, V6872 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6868 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6869 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6870 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6871 Obj) Obj {
tmp20394 := GCons(__e, symsymbol, Nil)

tmp20395 := GCons(__e, sym_1_1_6, tmp20394)

tmp20396 := GCons(__e, symsymbol, tmp20395)

return TailCall6(__e, SymFn(symis_b), V6872, tmp20396, B6868, L6869, Key6870, C6871)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20397 := GValue(__e, symshen_4_dsigf_d)

tmp20398 := Call3(__e, SymFn(symshen_4assoc_1_6), symuntrack, tmp20393, tmp20397)


tmp20399 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20398)

_ = tmp20399

tmp20400 := MakeNative1(func(__e *ControlFlow, V6877 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6873 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6874 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6875 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6876 Obj) Obj {
tmp20401 := GCons(__e, symsymbol, Nil)

tmp20402 := GCons(__e, symlist, tmp20401)

tmp20403 := GCons(__e, tmp20402, Nil)

tmp20404 := GCons(__e, sym_1_1_6, tmp20403)

return TailCall6(__e, SymFn(symis_b), V6877, tmp20404, B6873, L6874, Key6875, C6876)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20405 := GValue(__e, symshen_4_dsigf_d)

tmp20406 := Call3(__e, SymFn(symshen_4assoc_1_6), symuserdefs, tmp20400, tmp20405)


tmp20407 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20406)

_ = tmp20407

tmp20408 := MakeNative1(func(__e *ControlFlow, V6882 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6878 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6879 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6880 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6881 Obj) Obj {
tmp20409 := Call1(__e, SymFn(symshen_4newpv), B6878)


let__18090 := tmp20409
_ = let__18090

tmp20410 := GCons(__e, symboolean, Nil)

tmp20411 := GCons(__e, sym_1_1_6, tmp20410)

tmp20412 := GCons(__e, let__18090, tmp20411)

tmp20413 := Call6(__e, SymFn(symis_b), V6882, tmp20412, B6878, L6879, Key6880, C6881)


return TailCall2(__e, SymFn(symshen_4gc), B6878, tmp20413)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20414 := GValue(__e, symshen_4_dsigf_d)

tmp20415 := Call3(__e, SymFn(symshen_4assoc_1_6), symvariable_2, tmp20408, tmp20414)


tmp20416 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20415)

_ = tmp20416

tmp20417 := MakeNative1(func(__e *ControlFlow, V6887 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6883 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6884 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6885 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6886 Obj) Obj {
tmp20418 := Call1(__e, SymFn(symshen_4newpv), B6883)


let__18091 := tmp20418
_ = let__18091

tmp20419 := GCons(__e, symboolean, Nil)

tmp20420 := GCons(__e, sym_1_1_6, tmp20419)

tmp20421 := GCons(__e, let__18091, tmp20420)

tmp20422 := Call6(__e, SymFn(symis_b), V6887, tmp20421, B6883, L6884, Key6885, C6886)


return TailCall2(__e, SymFn(symshen_4gc), B6883, tmp20422)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20423 := GValue(__e, symshen_4_dsigf_d)

tmp20424 := Call3(__e, SymFn(symshen_4assoc_1_6), symvector_2, tmp20417, tmp20423)


tmp20425 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20424)

_ = tmp20425

tmp20426 := MakeNative1(func(__e *ControlFlow, V6892 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6888 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6889 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6890 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6891 Obj) Obj {
tmp20427 := GCons(__e, symstring, Nil)

tmp20428 := GCons(__e, sym_1_1_6, tmp20427)

return TailCall6(__e, SymFn(symis_b), V6892, tmp20428, B6888, L6889, Key6890, C6891)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20429 := GValue(__e, symshen_4_dsigf_d)

tmp20430 := Call3(__e, SymFn(symshen_4assoc_1_6), symversion, tmp20426, tmp20429)


tmp20431 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20430)

_ = tmp20431

tmp20432 := MakeNative1(func(__e *ControlFlow, V6897 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6893 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6894 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6895 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6896 Obj) Obj {
tmp20433 := Call1(__e, SymFn(symshen_4newpv), B6893)


let__18092 := tmp20433
_ = let__18092

tmp20434 := GCons(__e, let__18092, Nil)

tmp20435 := GCons(__e, sym_1_1_6, tmp20434)

tmp20436 := GCons(__e, let__18092, tmp20435)

tmp20437 := GCons(__e, tmp20436, Nil)

tmp20438 := GCons(__e, sym_1_1_6, tmp20437)

tmp20439 := GCons(__e, symstring, tmp20438)

tmp20440 := Call6(__e, SymFn(symis_b), V6897, tmp20439, B6893, L6894, Key6895, C6896)


return TailCall2(__e, SymFn(symshen_4gc), B6893, tmp20440)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20441 := GValue(__e, symshen_4_dsigf_d)

tmp20442 := Call3(__e, SymFn(symshen_4assoc_1_6), symwrite_1to_1file, tmp20432, tmp20441)


tmp20443 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20442)

_ = tmp20443

tmp20444 := MakeNative1(func(__e *ControlFlow, V6902 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6898 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6899 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6900 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6901 Obj) Obj {
tmp20445 := GCons(__e, symout, Nil)

tmp20446 := GCons(__e, symstream, tmp20445)

tmp20447 := GCons(__e, symnumber, Nil)

tmp20448 := GCons(__e, sym_1_1_6, tmp20447)

tmp20449 := GCons(__e, tmp20446, tmp20448)

tmp20450 := GCons(__e, tmp20449, Nil)

tmp20451 := GCons(__e, sym_1_1_6, tmp20450)

tmp20452 := GCons(__e, symnumber, tmp20451)

return TailCall6(__e, SymFn(symis_b), V6902, tmp20452, B6898, L6899, Key6900, C6901)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20453 := GValue(__e, symshen_4_dsigf_d)

tmp20454 := Call3(__e, SymFn(symshen_4assoc_1_6), symwrite_1byte, tmp20444, tmp20453)


tmp20455 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20454)

_ = tmp20455

tmp20456 := MakeNative1(func(__e *ControlFlow, V6907 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6903 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6904 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6905 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6906 Obj) Obj {
tmp20457 := GCons(__e, symout, Nil)

tmp20458 := GCons(__e, symstream, tmp20457)

tmp20459 := GCons(__e, symstring, Nil)

tmp20460 := GCons(__e, sym_1_1_6, tmp20459)

tmp20461 := GCons(__e, tmp20458, tmp20460)

tmp20462 := GCons(__e, tmp20461, Nil)

tmp20463 := GCons(__e, sym_1_1_6, tmp20462)

tmp20464 := GCons(__e, symstring, tmp20463)

return TailCall6(__e, SymFn(symis_b), V6907, tmp20464, B6903, L6904, Key6905, C6906)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20465 := GValue(__e, symshen_4_dsigf_d)

tmp20466 := Call3(__e, SymFn(symshen_4assoc_1_6), symwrite_1string, tmp20456, tmp20465)


tmp20467 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20466)

_ = tmp20467

tmp20468 := MakeNative1(func(__e *ControlFlow, V6912 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6908 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6909 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6910 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6911 Obj) Obj {
tmp20469 := GCons(__e, symboolean, Nil)

tmp20470 := GCons(__e, sym_1_1_6, tmp20469)

tmp20471 := GCons(__e, symstring, tmp20470)

return TailCall6(__e, SymFn(symis_b), V6912, tmp20471, B6908, L6909, Key6910, C6911)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20472 := GValue(__e, symshen_4_dsigf_d)

tmp20473 := Call3(__e, SymFn(symshen_4assoc_1_6), symy_1or_1n_2, tmp20468, tmp20472)


tmp20474 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20473)

_ = tmp20474

tmp20475 := MakeNative1(func(__e *ControlFlow, V6917 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6913 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6914 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6915 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6916 Obj) Obj {
tmp20476 := GCons(__e, symboolean, Nil)

tmp20477 := GCons(__e, sym_1_1_6, tmp20476)

tmp20478 := GCons(__e, symnumber, tmp20477)

tmp20479 := GCons(__e, tmp20478, Nil)

tmp20480 := GCons(__e, sym_1_1_6, tmp20479)

tmp20481 := GCons(__e, symnumber, tmp20480)

return TailCall6(__e, SymFn(symis_b), V6917, tmp20481, B6913, L6914, Key6915, C6916)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20482 := GValue(__e, symshen_4_dsigf_d)

tmp20483 := Call3(__e, SymFn(symshen_4assoc_1_6), sym_6, tmp20475, tmp20482)


tmp20484 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20483)

_ = tmp20484

tmp20485 := MakeNative1(func(__e *ControlFlow, V6922 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6918 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6919 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6920 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6921 Obj) Obj {
tmp20486 := GCons(__e, symboolean, Nil)

tmp20487 := GCons(__e, sym_1_1_6, tmp20486)

tmp20488 := GCons(__e, symnumber, tmp20487)

tmp20489 := GCons(__e, tmp20488, Nil)

tmp20490 := GCons(__e, sym_1_1_6, tmp20489)

tmp20491 := GCons(__e, symnumber, tmp20490)

return TailCall6(__e, SymFn(symis_b), V6922, tmp20491, B6918, L6919, Key6920, C6921)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20492 := GValue(__e, symshen_4_dsigf_d)

tmp20493 := Call3(__e, SymFn(symshen_4assoc_1_6), sym_5, tmp20485, tmp20492)


tmp20494 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20493)

_ = tmp20494

tmp20495 := MakeNative1(func(__e *ControlFlow, V6927 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6923 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6924 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6925 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6926 Obj) Obj {
tmp20496 := GCons(__e, symboolean, Nil)

tmp20497 := GCons(__e, sym_1_1_6, tmp20496)

tmp20498 := GCons(__e, symnumber, tmp20497)

tmp20499 := GCons(__e, tmp20498, Nil)

tmp20500 := GCons(__e, sym_1_1_6, tmp20499)

tmp20501 := GCons(__e, symnumber, tmp20500)

return TailCall6(__e, SymFn(symis_b), V6927, tmp20501, B6923, L6924, Key6925, C6926)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20502 := GValue(__e, symshen_4_dsigf_d)

tmp20503 := Call3(__e, SymFn(symshen_4assoc_1_6), sym_6_a, tmp20495, tmp20502)


tmp20504 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20503)

_ = tmp20504

tmp20505 := MakeNative1(func(__e *ControlFlow, V6932 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6928 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6929 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6930 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6931 Obj) Obj {
tmp20506 := GCons(__e, symboolean, Nil)

tmp20507 := GCons(__e, sym_1_1_6, tmp20506)

tmp20508 := GCons(__e, symnumber, tmp20507)

tmp20509 := GCons(__e, tmp20508, Nil)

tmp20510 := GCons(__e, sym_1_1_6, tmp20509)

tmp20511 := GCons(__e, symnumber, tmp20510)

return TailCall6(__e, SymFn(symis_b), V6932, tmp20511, B6928, L6929, Key6930, C6931)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20512 := GValue(__e, symshen_4_dsigf_d)

tmp20513 := Call3(__e, SymFn(symshen_4assoc_1_6), sym_5_a, tmp20505, tmp20512)


tmp20514 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20513)

_ = tmp20514

tmp20515 := MakeNative1(func(__e *ControlFlow, V6937 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6933 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6934 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6935 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6936 Obj) Obj {
tmp20516 := Call1(__e, SymFn(symshen_4newpv), B6933)


let__18093 := tmp20516
_ = let__18093

tmp20517 := GCons(__e, symboolean, Nil)

tmp20518 := GCons(__e, sym_1_1_6, tmp20517)

tmp20519 := GCons(__e, let__18093, tmp20518)

tmp20520 := GCons(__e, tmp20519, Nil)

tmp20521 := GCons(__e, sym_1_1_6, tmp20520)

tmp20522 := GCons(__e, let__18093, tmp20521)

tmp20523 := Call6(__e, SymFn(symis_b), V6937, tmp20522, B6933, L6934, Key6935, C6936)


return TailCall2(__e, SymFn(symshen_4gc), B6933, tmp20523)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20524 := GValue(__e, symshen_4_dsigf_d)

tmp20525 := Call3(__e, SymFn(symshen_4assoc_1_6), sym_a, tmp20515, tmp20524)


tmp20526 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20525)

_ = tmp20526

tmp20527 := MakeNative1(func(__e *ControlFlow, V6942 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6938 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6939 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6940 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6941 Obj) Obj {
tmp20528 := GCons(__e, symnumber, Nil)

tmp20529 := GCons(__e, sym_1_1_6, tmp20528)

tmp20530 := GCons(__e, symnumber, tmp20529)

tmp20531 := GCons(__e, tmp20530, Nil)

tmp20532 := GCons(__e, sym_1_1_6, tmp20531)

tmp20533 := GCons(__e, symnumber, tmp20532)

return TailCall6(__e, SymFn(symis_b), V6942, tmp20533, B6938, L6939, Key6940, C6941)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20534 := GValue(__e, symshen_4_dsigf_d)

tmp20535 := Call3(__e, SymFn(symshen_4assoc_1_6), sym_7, tmp20527, tmp20534)


tmp20536 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20535)

_ = tmp20536

tmp20537 := MakeNative1(func(__e *ControlFlow, V6947 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6943 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6944 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6945 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6946 Obj) Obj {
tmp20538 := GCons(__e, symnumber, Nil)

tmp20539 := GCons(__e, sym_1_1_6, tmp20538)

tmp20540 := GCons(__e, symnumber, tmp20539)

tmp20541 := GCons(__e, tmp20540, Nil)

tmp20542 := GCons(__e, sym_1_1_6, tmp20541)

tmp20543 := GCons(__e, symnumber, tmp20542)

return TailCall6(__e, SymFn(symis_b), V6947, tmp20543, B6943, L6944, Key6945, C6946)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20544 := GValue(__e, symshen_4_dsigf_d)

tmp20545 := Call3(__e, SymFn(symshen_4assoc_1_6), sym_c, tmp20537, tmp20544)


tmp20546 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20545)

_ = tmp20546

tmp20547 := MakeNative1(func(__e *ControlFlow, V6952 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6948 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6949 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6950 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6951 Obj) Obj {
tmp20548 := GCons(__e, symnumber, Nil)

tmp20549 := GCons(__e, sym_1_1_6, tmp20548)

tmp20550 := GCons(__e, symnumber, tmp20549)

tmp20551 := GCons(__e, tmp20550, Nil)

tmp20552 := GCons(__e, sym_1_1_6, tmp20551)

tmp20553 := GCons(__e, symnumber, tmp20552)

return TailCall6(__e, SymFn(symis_b), V6952, tmp20553, B6948, L6949, Key6950, C6951)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20554 := GValue(__e, symshen_4_dsigf_d)

tmp20555 := Call3(__e, SymFn(symshen_4assoc_1_6), sym_1, tmp20547, tmp20554)


tmp20556 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20555)

_ = tmp20556

tmp20557 := MakeNative1(func(__e *ControlFlow, V6957 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6953 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6954 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6955 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6956 Obj) Obj {
tmp20558 := GCons(__e, symnumber, Nil)

tmp20559 := GCons(__e, sym_1_1_6, tmp20558)

tmp20560 := GCons(__e, symnumber, tmp20559)

tmp20561 := GCons(__e, tmp20560, Nil)

tmp20562 := GCons(__e, sym_1_1_6, tmp20561)

tmp20563 := GCons(__e, symnumber, tmp20562)

return TailCall6(__e, SymFn(symis_b), V6957, tmp20563, B6953, L6954, Key6955, C6956)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20564 := GValue(__e, symshen_4_dsigf_d)

tmp20565 := Call3(__e, SymFn(symshen_4assoc_1_6), sym_d, tmp20557, tmp20564)


tmp20566 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20565)

_ = tmp20566

tmp20567 := MakeNative1(func(__e *ControlFlow, V6962 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, B6958 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, L6959 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Key6960 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, C6961 Obj) Obj {
tmp20568 := Call1(__e, SymFn(symshen_4newpv), B6958)


let__18094 := tmp20568
_ = let__18094

tmp20569 := Call1(__e, SymFn(symshen_4newpv), B6958)


let__18095 := tmp20569
_ = let__18095

tmp20570 := GCons(__e, symboolean, Nil)

tmp20571 := GCons(__e, sym_1_1_6, tmp20570)

tmp20572 := GCons(__e, let__18095, tmp20571)

tmp20573 := GCons(__e, tmp20572, Nil)

tmp20574 := GCons(__e, sym_1_1_6, tmp20573)

tmp20575 := GCons(__e, let__18094, tmp20574)

tmp20576 := Call6(__e, SymFn(symis_b), V6962, tmp20575, B6958, L6959, Key6960, C6961)


tmp20577 := Call2(__e, SymFn(symshen_4gc), B6958, tmp20576)


return TailCall2(__e, SymFn(symshen_4gc), B6958, tmp20577)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20578 := GValue(__e, symshen_4_dsigf_d)

tmp20579 := Call3(__e, SymFn(symshen_4assoc_1_6), sym_a_a, tmp20567, tmp20578)


return PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp20579)


panic("unreachable")
})

tmp20580 := Call2(__e, ns2_1set, symshen_4initialise_1signedfuncs, tmp18861)


_ = tmp20580

tmp20581 := MakeNative0(func(__e *ControlFlow) Obj {
tmp20582 := MakeNative1(func(__e *ControlFlow, Y954 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4tuple), Y954)
panic("unreachable")
})

tmp20583 := GCons(__e, symshen_4tuple, tmp20582)

tmp20584 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20583)


_ = tmp20584

tmp20585 := MakeNative1(func(__e *ControlFlow, Y955 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4pvar), Y955)
panic("unreachable")
})

tmp20586 := GCons(__e, symshen_4pvar, tmp20585)

tmp20587 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20586)


_ = tmp20587

tmp20588 := MakeNative1(func(__e *ControlFlow, Y956 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4dictionary), Y956)
panic("unreachable")
})

tmp20589 := GCons(__e, symshen_4dictionary, tmp20588)

tmp20590 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20589)


_ = tmp20590

tmp20591 := MakeNative1(func(__e *ControlFlow, Y957 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4print_1prolog_1vector), Y957)
panic("unreachable")
})

tmp20592 := GCons(__e, symshen_4print_1prolog_1vector, tmp20591)

tmp20593 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20592)


_ = tmp20593

tmp20594 := MakeNative1(func(__e *ControlFlow, Y958 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4print_1freshterm), Y958)
panic("unreachable")
})

tmp20595 := GCons(__e, symshen_4print_1freshterm, tmp20594)

tmp20596 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20595)


_ = tmp20596

tmp20597 := MakeNative1(func(__e *ControlFlow, Y959 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4printF), Y959)
panic("unreachable")
})

tmp20598 := GCons(__e, symshen_4printF, tmp20597)

tmp20599 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20598)


_ = tmp20599

tmp20600 := MakeNative1(func(__e *ControlFlow, Y960 Obj) Obj {
return TailCall1(__e, SymFn(symabsolute), Y960)
panic("unreachable")
})

tmp20601 := GCons(__e, symabsolute, tmp20600)

tmp20602 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20601)


_ = tmp20602

tmp20603 := MakeNative1(func(__e *ControlFlow, Y961 Obj) Obj {
return GVectorP(__e, Y961)
panic("unreachable")
})

tmp20604 := GCons(__e, symabsvector_2, tmp20603)

tmp20605 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20604)


_ = tmp20605

tmp20606 := MakeNative1(func(__e *ControlFlow, Y962 Obj) Obj {
return PrimGuard1(__e, symabsvector, PrimAbsvector, Y962)
panic("unreachable")
})

tmp20607 := GCons(__e, symabsvector, tmp20606)

tmp20608 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20607)


_ = tmp20608

tmp20609 := MakeNative1(func(__e *ControlFlow, Y963 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y964 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y965 Obj) Obj {
return GVecSet(__e, Y963, Y964, Y965)
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20610 := GCons(__e, symaddress_1_6, tmp20609)

tmp20611 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20610)


_ = tmp20611

tmp20612 := MakeNative1(func(__e *ControlFlow, Y966 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y967 Obj) Obj {
return TailCall2(__e, SymFn(symadjoin), Y966, Y967)
panic("unreachable")
})
panic("unreachable")
})

tmp20613 := GCons(__e, symadjoin, tmp20612)

tmp20614 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20613)


_ = tmp20614

tmp20615 := MakeNative1(func(__e *ControlFlow, Y968 Obj) Obj {
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

tmp20617 := GCons(__e, symand, tmp20615)

tmp20618 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20617)


_ = tmp20618

tmp20619 := MakeNative1(func(__e *ControlFlow, Y970 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y971 Obj) Obj {
return TailCall2(__e, SymFn(symappend), Y970, Y971)
panic("unreachable")
})
panic("unreachable")
})

tmp20620 := GCons(__e, symappend, tmp20619)

tmp20621 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20620)


_ = tmp20621

tmp20622 := MakeNative1(func(__e *ControlFlow, Y972 Obj) Obj {
return TailCall1(__e, SymFn(symarity), Y972)
panic("unreachable")
})

tmp20623 := GCons(__e, symarity, tmp20622)

tmp20624 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20623)


_ = tmp20624

tmp20625 := MakeNative1(func(__e *ControlFlow, Y973 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y974 Obj) Obj {
return TailCall2(__e, SymFn(symassoc), Y973, Y974)
panic("unreachable")
})
panic("unreachable")
})

tmp20626 := GCons(__e, symassoc, tmp20625)

tmp20627 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20626)


_ = tmp20627

tmp20628 := MakeNative1(func(__e *ControlFlow, Y975 Obj) Obj {
return TailCall1(__e, SymFn(symatom_2), Y975)
panic("unreachable")
})

tmp20629 := GCons(__e, symatom_2, tmp20628)

tmp20630 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20629)


_ = tmp20630

tmp20631 := MakeNative1(func(__e *ControlFlow, Y976 Obj) Obj {
return TailCall1(__e, SymFn(symboolean_2), Y976)
panic("unreachable")
})

tmp20632 := GCons(__e, symboolean_2, tmp20631)

tmp20633 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20632)


_ = tmp20633

tmp20634 := MakeNative1(func(__e *ControlFlow, Y977 Obj) Obj {
return TailCall1(__e, SymFn(symbootstrap), Y977)
panic("unreachable")
})

tmp20635 := GCons(__e, symbootstrap, tmp20634)

tmp20636 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20635)


_ = tmp20636

tmp20637 := MakeNative1(func(__e *ControlFlow, Y978 Obj) Obj {
return TailCall1(__e, SymFn(symbound_2), Y978)
panic("unreachable")
})

tmp20638 := GCons(__e, symbound_2, tmp20637)

tmp20639 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20638)


_ = tmp20639

tmp20640 := MakeNative1(func(__e *ControlFlow, Y979 Obj) Obj {
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

tmp20641 := GCons(__e, symbind, tmp20640)

tmp20642 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20641)


_ = tmp20642

tmp20643 := MakeNative1(func(__e *ControlFlow, Y985 Obj) Obj {
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

tmp20644 := GCons(__e, symcall, tmp20643)

tmp20645 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20644)


_ = tmp20645

tmp20646 := MakeNative1(func(__e *ControlFlow, Y990 Obj) Obj {
return TailCall1(__e, SymFn(symcd), Y990)
panic("unreachable")
})

tmp20647 := GCons(__e, symcd, tmp20646)

tmp20648 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20647)


_ = tmp20648

tmp20649 := MakeNative1(func(__e *ControlFlow, Y991 Obj) Obj {
return TailCall1(__e, SymFn(symchar_1stream_2), Y991)
panic("unreachable")
})

tmp20650 := GCons(__e, symchar_1stream_2, tmp20649)

tmp20651 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20650)


_ = tmp20651

tmp20652 := MakeNative1(func(__e *ControlFlow, Y992 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y993 Obj) Obj {
return TailCall2(__e, SymFn(symcompile), Y992, Y993)
panic("unreachable")
})
panic("unreachable")
})

tmp20653 := GCons(__e, symcompile, tmp20652)

tmp20654 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20653)


_ = tmp20654

tmp20655 := MakeNative1(func(__e *ControlFlow, Y994 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y995 Obj) Obj {
return TailCall2(__e, SymFn(symconcat), Y994, Y995)
panic("unreachable")
})
panic("unreachable")
})

tmp20656 := GCons(__e, symconcat, tmp20655)

tmp20657 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20656)


_ = tmp20657

tmp20658 := MakeNative1(func(__e *ControlFlow, Y996 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y997 Obj) Obj {
return GCons(__e, Y996, Y997)
panic("unreachable")
})
panic("unreachable")
})

tmp20659 := GCons(__e, symcons, tmp20658)

tmp20660 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20659)


_ = tmp20660

tmp20661 := MakeNative1(func(__e *ControlFlow, Y998 Obj) Obj {
return GConsP(__e, Y998)
panic("unreachable")
})

tmp20662 := GCons(__e, symcons_2, tmp20661)

tmp20663 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20662)


_ = tmp20663

tmp20664 := MakeNative1(func(__e *ControlFlow, Y999 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1000 Obj) Obj {
return PrimGuard2(__e, symcn, PrimStringConcat, Y999, Y1000)
panic("unreachable")
})
panic("unreachable")
})

tmp20665 := GCons(__e, symcn, tmp20664)

tmp20666 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20665)


_ = tmp20666

tmp20667 := MakeNative1(func(__e *ControlFlow, Y1001 Obj) Obj {
return PrimGuard1(__e, symclose, PrimCloseStream, Y1001)
panic("unreachable")
})

tmp20668 := GCons(__e, symclose, tmp20667)

tmp20669 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20668)


_ = tmp20669

tmp20670 := MakeNative1(func(__e *ControlFlow, Y1002 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1003 Obj) Obj {
return TailCall2(__e, SymFn(symdeclare), Y1002, Y1003)
panic("unreachable")
})
panic("unreachable")
})

tmp20671 := GCons(__e, symdeclare, tmp20670)

tmp20672 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20671)


_ = tmp20672

tmp20673 := MakeNative1(func(__e *ControlFlow, Y1004 Obj) Obj {
return TailCall1(__e, SymFn(symdestroy), Y1004)
panic("unreachable")
})

tmp20674 := GCons(__e, symdestroy, tmp20673)

tmp20675 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20674)


_ = tmp20675

tmp20676 := MakeNative1(func(__e *ControlFlow, Y1005 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1006 Obj) Obj {
return TailCall2(__e, SymFn(symdifference), Y1005, Y1006)
panic("unreachable")
})
panic("unreachable")
})

tmp20677 := GCons(__e, symdifference, tmp20676)

tmp20678 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20677)


_ = tmp20678

tmp20679 := MakeNative1(func(__e *ControlFlow, Y1007 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1008 Obj) Obj {
_ = Y1007

return Y1008


panic("unreachable")
})
panic("unreachable")
})

tmp20680 := GCons(__e, symdo, tmp20679)

tmp20681 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20680)


_ = tmp20681

tmp20682 := MakeNative1(func(__e *ControlFlow, Y1009 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1010 Obj) Obj {
return TailCall2(__e, SymFn(symelement_2), Y1009, Y1010)
panic("unreachable")
})
panic("unreachable")
})

tmp20683 := GCons(__e, symelement_2, tmp20682)

tmp20684 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20683)


_ = tmp20684

tmp20685 := MakeNative1(func(__e *ControlFlow, Y1011 Obj) Obj {
return TailCall1(__e, SymFn(symempty_2), Y1011)
panic("unreachable")
})

tmp20686 := GCons(__e, symempty_2, tmp20685)

tmp20687 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20686)


_ = tmp20687

tmp20688 := MakeNative1(func(__e *ControlFlow, Y1012 Obj) Obj {
return TailCall1(__e, SymFn(symenable_1type_1theory), Y1012)
panic("unreachable")
})

tmp20689 := GCons(__e, symenable_1type_1theory, tmp20688)

tmp20690 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20689)


_ = tmp20690

tmp20691 := MakeNative1(func(__e *ControlFlow, Y1013 Obj) Obj {
return TailCall1(__e, SymFn(symexternal), Y1013)
panic("unreachable")
})

tmp20692 := GCons(__e, symexternal, tmp20691)

tmp20693 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20692)


_ = tmp20693

tmp20694 := MakeNative1(func(__e *ControlFlow, Y1014 Obj) Obj {
return PrimGuard1(__e, symerror_1to_1string, PrimErrorToString, Y1014)
panic("unreachable")
})

tmp20695 := GCons(__e, symerror_1to_1string, tmp20694)

tmp20696 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20695)


_ = tmp20696

tmp20697 := MakeNative1(func(__e *ControlFlow, Y1015 Obj) Obj {
return TailCall1(__e, SymFn(symeval), Y1015)
panic("unreachable")
})

tmp20698 := GCons(__e, symeval, tmp20697)

tmp20699 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20698)


_ = tmp20699

tmp20700 := MakeNative1(func(__e *ControlFlow, Y1016 Obj) Obj {
return TailCall1(__e, SymFn(symeval_1kl), Y1016)
panic("unreachable")
})

tmp20701 := GCons(__e, symeval_1kl, tmp20700)

tmp20702 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20701)


_ = tmp20702

tmp20703 := MakeNative1(func(__e *ControlFlow, Y1017 Obj) Obj {
return TailCall1(__e, SymFn(symexplode), Y1017)
panic("unreachable")
})

tmp20704 := GCons(__e, symexplode, tmp20703)

tmp20705 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20704)


_ = tmp20705

tmp20706 := MakeNative1(func(__e *ControlFlow, Y1018 Obj) Obj {
return TailCall1(__e, SymFn(symexternal), Y1018)
panic("unreachable")
})

tmp20707 := GCons(__e, symexternal, tmp20706)

tmp20708 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20707)


_ = tmp20708

tmp20709 := MakeNative1(func(__e *ControlFlow, Y1019 Obj) Obj {
return TailCall1(__e, SymFn(symfactorise), Y1019)
panic("unreachable")
})

tmp20710 := GCons(__e, symfactorise, tmp20709)

tmp20711 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20710)


_ = tmp20711

tmp20712 := MakeNative1(func(__e *ControlFlow, Y1020 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1021 Obj) Obj {
return TailCall2(__e, SymFn(symfail_1if), Y1020, Y1021)
panic("unreachable")
})
panic("unreachable")
})

tmp20713 := GCons(__e, symfail_1if, tmp20712)

tmp20714 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20713)


_ = tmp20714

tmp20715 := MakeNative1(func(__e *ControlFlow, Y1022 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1023 Obj) Obj {
return TailCall2(__e, SymFn(symfix), Y1022, Y1023)
panic("unreachable")
})
panic("unreachable")
})

tmp20716 := GCons(__e, symfix, tmp20715)

tmp20717 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20716)


_ = tmp20717

tmp20718 := MakeNative1(func(__e *ControlFlow, Y1024 Obj) Obj {
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

tmp20719 := GCons(__e, symfindall, tmp20718)

tmp20720 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20719)


_ = tmp20720

tmp20721 := MakeNative1(func(__e *ControlFlow, Y1031 Obj) Obj {
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

tmp20722 := GCons(__e, symfork, tmp20721)

tmp20723 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20722)


_ = tmp20723

tmp20724 := MakeNative1(func(__e *ControlFlow, Y1036 Obj) Obj {
return MakeNative0(func(__e *ControlFlow) Obj {
return Y1036
panic("unreachable")
})
panic("unreachable")
})

tmp20725 := GCons(__e, symfreeze, tmp20724)

tmp20726 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20725)


_ = tmp20726

tmp20727 := MakeNative1(func(__e *ControlFlow, Y1037 Obj) Obj {
return TailCall1(__e, SymFn(symfst), Y1037)
panic("unreachable")
})

tmp20728 := GCons(__e, symfst, tmp20727)

tmp20729 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20728)


_ = tmp20729

tmp20730 := MakeNative1(func(__e *ControlFlow, Y1038 Obj) Obj {
return TailCall1(__e, SymFn(symfn), Y1038)
panic("unreachable")
})

tmp20731 := GCons(__e, symfn, tmp20730)

tmp20732 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20731)


_ = tmp20732

tmp20733 := MakeNative1(func(__e *ControlFlow, Y1039 Obj) Obj {
return TailCall1(__e, SymFn(symfunction), Y1039)
panic("unreachable")
})

tmp20734 := GCons(__e, symfunction, tmp20733)

tmp20735 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20734)


_ = tmp20735

tmp20736 := MakeNative1(func(__e *ControlFlow, Y1040 Obj) Obj {
return TailCall1(__e, SymFn(symgensym), Y1040)
panic("unreachable")
})

tmp20737 := GCons(__e, symgensym, tmp20736)

tmp20738 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20737)


_ = tmp20738

tmp20739 := MakeNative1(func(__e *ControlFlow, Y1041 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1042 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1043 Obj) Obj {
return TailCall3(__e, SymFn(symget), Y1041, Y1042, Y1043)
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20740 := GCons(__e, symget, tmp20739)

tmp20741 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20740)


_ = tmp20741

tmp20742 := MakeNative1(func(__e *ControlFlow, Y1044 Obj) Obj {
return PrimGuard1(__e, symget_1time, PrimGetTime, Y1044)
panic("unreachable")
})

tmp20743 := GCons(__e, symget_1time, tmp20742)

tmp20744 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20743)


_ = tmp20744

tmp20745 := MakeNative1(func(__e *ControlFlow, Y1045 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1046 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1047 Obj) Obj {
return GVecSet(__e, Y1045, Y1046, Y1047)
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp20746 := GCons(__e, symaddress_1_6, tmp20745)

tmp20747 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20746)


_ = tmp20747

tmp20748 := MakeNative1(func(__e *ControlFlow, Y1048 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1049 Obj) Obj {
return GVecGet(__e, Y1048, Y1049)
panic("unreachable")
})
panic("unreachable")
})

tmp20749 := GCons(__e, sym_5_1address, tmp20748)

tmp20750 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20749)


_ = tmp20750

tmp20751 := MakeNative1(func(__e *ControlFlow, Y1050 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1051 Obj) Obj {
return TailCall2(__e, SymFn(sym_5_1vector), Y1050, Y1051)
panic("unreachable")
})
panic("unreachable")
})

tmp20752 := GCons(__e, sym_5_1vector, tmp20751)

tmp20753 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20752)


_ = tmp20753

tmp20754 := MakeNative1(func(__e *ControlFlow, Y1052 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1053 Obj) Obj {
return GGt(__e, Y1052, Y1053)
panic("unreachable")
})
panic("unreachable")
})

tmp20755 := GCons(__e, sym_6, tmp20754)

tmp20756 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20755)


_ = tmp20756

tmp20757 := MakeNative1(func(__e *ControlFlow, Y1054 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1055 Obj) Obj {
return GGe(__e, Y1054, Y1055)
panic("unreachable")
})
panic("unreachable")
})

tmp20758 := GCons(__e, sym_6_a, tmp20757)

tmp20759 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20758)


_ = tmp20759

tmp20760 := MakeNative1(func(__e *ControlFlow, Y1056 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1057 Obj) Obj {
return GEq(__e, Y1056, Y1057)
panic("unreachable")
})
panic("unreachable")
})

tmp20761 := GCons(__e, sym_a, tmp20760)

tmp20762 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20761)


_ = tmp20762

tmp20763 := MakeNative1(func(__e *ControlFlow, Y1058 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1059 Obj) Obj {
return TailCall2(__e, SymFn(symhash), Y1058, Y1059)
panic("unreachable")
})
panic("unreachable")
})

tmp20764 := GCons(__e, symhash, tmp20763)

tmp20765 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20764)


_ = tmp20765

tmp20766 := MakeNative1(func(__e *ControlFlow, Y1060 Obj) Obj {
return GHd(__e, Y1060)
panic("unreachable")
})

tmp20767 := GCons(__e, symhd, tmp20766)

tmp20768 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20767)


_ = tmp20768

tmp20769 := MakeNative1(func(__e *ControlFlow, Y1061 Obj) Obj {
return TailCall1(__e, SymFn(symhdv), Y1061)
panic("unreachable")
})

tmp20770 := GCons(__e, symhdv, tmp20769)

tmp20771 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20770)


_ = tmp20771

tmp20772 := MakeNative1(func(__e *ControlFlow, Y1062 Obj) Obj {
return TailCall1(__e, SymFn(symhdstr), Y1062)
panic("unreachable")
})

tmp20773 := GCons(__e, symhdstr, tmp20772)

tmp20774 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20773)


_ = tmp20774

tmp20775 := MakeNative1(func(__e *ControlFlow, Y1063 Obj) Obj {
return TailCall1(__e, SymFn(symhead), Y1063)
panic("unreachable")
})

tmp20776 := GCons(__e, symhead, tmp20775)

tmp20777 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20776)


_ = tmp20777

tmp20778 := MakeNative1(func(__e *ControlFlow, Y1064 Obj) Obj {
return TailCall1(__e, SymFn(symhush), Y1064)
panic("unreachable")
})

tmp20779 := GCons(__e, symhush, tmp20778)

tmp20780 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20779)


_ = tmp20780

tmp20781 := MakeNative1(func(__e *ControlFlow, Y1065 Obj) Obj {
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

tmp20783 := GCons(__e, symif, tmp20781)

tmp20784 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20783)


_ = tmp20784

tmp20785 := MakeNative1(func(__e *ControlFlow, Y1068 Obj) Obj {
return TailCall1(__e, SymFn(syminclude), Y1068)
panic("unreachable")
})

tmp20786 := GCons(__e, syminclude, tmp20785)

tmp20787 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20786)


_ = tmp20787

tmp20788 := MakeNative1(func(__e *ControlFlow, Y1069 Obj) Obj {
return TailCall1(__e, SymFn(symin_1package), Y1069)
panic("unreachable")
})

tmp20789 := GCons(__e, symin_1package, tmp20788)

tmp20790 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20789)


_ = tmp20790

tmp20791 := MakeNative1(func(__e *ControlFlow, Y1070 Obj) Obj {
return PrimGuard1(__e, syminteger_2, PrimIsInteger, Y1070)
panic("unreachable")
})

tmp20792 := GCons(__e, syminteger_2, tmp20791)

tmp20793 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20792)


_ = tmp20793

tmp20794 := MakeNative1(func(__e *ControlFlow, Y1071 Obj) Obj {
return TailCall1(__e, SymFn(syminternal), Y1071)
panic("unreachable")
})

tmp20795 := GCons(__e, syminternal, tmp20794)

tmp20796 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20795)


_ = tmp20796

tmp20797 := MakeNative1(func(__e *ControlFlow, Y1072 Obj) Obj {
return PrimGuard1(__e, symintern, PrimIntern, Y1072)
panic("unreachable")
})

tmp20798 := GCons(__e, symintern, tmp20797)

tmp20799 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20798)


_ = tmp20799

tmp20800 := MakeNative1(func(__e *ControlFlow, Y1073 Obj) Obj {
return TailCall1(__e, SymFn(syminput), Y1073)
panic("unreachable")
})

tmp20801 := GCons(__e, syminput, tmp20800)

tmp20802 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20801)


_ = tmp20802

tmp20803 := MakeNative1(func(__e *ControlFlow, Y1074 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1075 Obj) Obj {
return TailCall2(__e, SymFn(symshen_4input_1h_7), Y1074, Y1075)
panic("unreachable")
})
panic("unreachable")
})

tmp20804 := GCons(__e, syminput_7, tmp20803)

tmp20805 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20804)


_ = tmp20805

tmp20806 := MakeNative1(func(__e *ControlFlow, Y1076 Obj) Obj {
return TailCall1(__e, SymFn(syminclude_1all_1but), Y1076)
panic("unreachable")
})

tmp20807 := GCons(__e, syminclude_1all_1but, tmp20806)

tmp20808 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20807)


_ = tmp20808

tmp20809 := MakeNative1(func(__e *ControlFlow, Y1077 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1078 Obj) Obj {
return TailCall2(__e, SymFn(symintersection), Y1077, Y1078)
panic("unreachable")
})
panic("unreachable")
})

tmp20810 := GCons(__e, symintersection, tmp20809)

tmp20811 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20810)


_ = tmp20811

tmp20812 := MakeNative1(func(__e *ControlFlow, Y1079 Obj) Obj {
return TailCall1(__e, SymFn(syminternal), Y1079)
panic("unreachable")
})

tmp20813 := GCons(__e, syminternal, tmp20812)

tmp20814 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20813)


_ = tmp20814

tmp20815 := MakeNative1(func(__e *ControlFlow, Y1080 Obj) Obj {
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

tmp20816 := GCons(__e, symis, tmp20815)

tmp20817 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20816)


_ = tmp20817

tmp20818 := MakeNative1(func(__e *ControlFlow, Y1086 Obj) Obj {
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

tmp20819 := GCons(__e, symis_b, tmp20818)

tmp20820 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20819)


_ = tmp20820

tmp20821 := MakeNative1(func(__e *ControlFlow, Y1092 Obj) Obj {
return TailCall1(__e, SymFn(symlength), Y1092)
panic("unreachable")
})

tmp20822 := GCons(__e, symlength, tmp20821)

tmp20823 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20822)


_ = tmp20823

tmp20824 := MakeNative1(func(__e *ControlFlow, Y1093 Obj) Obj {
return TailCall1(__e, SymFn(symlimit), Y1093)
panic("unreachable")
})

tmp20825 := GCons(__e, symlimit, tmp20824)

tmp20826 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20825)


_ = tmp20826

tmp20827 := MakeNative1(func(__e *ControlFlow, Y1094 Obj) Obj {
return TailCall1(__e, SymFn(symlineread), Y1094)
panic("unreachable")
})

tmp20828 := GCons(__e, symlineread, tmp20827)

tmp20829 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20828)


_ = tmp20829

tmp20830 := MakeNative1(func(__e *ControlFlow, Y1095 Obj) Obj {
return TailCall1(__e, SymFn(symload), Y1095)
panic("unreachable")
})

tmp20831 := GCons(__e, symload, tmp20830)

tmp20832 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20831)


_ = tmp20832

tmp20833 := MakeNative1(func(__e *ControlFlow, Y1096 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1097 Obj) Obj {
return GLt(__e, Y1096, Y1097)
panic("unreachable")
})
panic("unreachable")
})

tmp20834 := GCons(__e, sym_5, tmp20833)

tmp20835 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20834)


_ = tmp20835

tmp20836 := MakeNative1(func(__e *ControlFlow, Y1098 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1099 Obj) Obj {
return GLe(__e, Y1098, Y1099)
panic("unreachable")
})
panic("unreachable")
})

tmp20837 := GCons(__e, sym_5_a, tmp20836)

tmp20838 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20837)


_ = tmp20838

tmp20839 := MakeNative1(func(__e *ControlFlow, Y1100 Obj) Obj {
return TailCall1(__e, SymFn(symvector), Y1100)
panic("unreachable")
})

tmp20840 := GCons(__e, symvector, tmp20839)

tmp20841 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20840)


_ = tmp20841

tmp20842 := MakeNative1(func(__e *ControlFlow, Y1101 Obj) Obj {
return TailCall1(__e, SymFn(symmacroexpand), Y1101)
panic("unreachable")
})

tmp20843 := GCons(__e, symmacroexpand, tmp20842)

tmp20844 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20843)


_ = tmp20844

tmp20845 := MakeNative1(func(__e *ControlFlow, Y1102 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1103 Obj) Obj {
return TailCall2(__e, SymFn(symmap), Y1102, Y1103)
panic("unreachable")
})
panic("unreachable")
})

tmp20846 := GCons(__e, symmap, tmp20845)

tmp20847 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20846)


_ = tmp20847

tmp20848 := MakeNative1(func(__e *ControlFlow, Y1104 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1105 Obj) Obj {
return TailCall2(__e, SymFn(symmapcan), Y1104, Y1105)
panic("unreachable")
})
panic("unreachable")
})

tmp20849 := GCons(__e, symmapcan, tmp20848)

tmp20850 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20849)


_ = tmp20850

tmp20851 := MakeNative1(func(__e *ControlFlow, Y1106 Obj) Obj {
return TailCall1(__e, SymFn(symmaxinferences), Y1106)
panic("unreachable")
})

tmp20852 := GCons(__e, symmaxinferences, tmp20851)

tmp20853 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20852)


_ = tmp20853

tmp20854 := MakeNative1(func(__e *ControlFlow, Y1107 Obj) Obj {
return TailCall1(__e, SymFn(symnl), Y1107)
panic("unreachable")
})

tmp20855 := GCons(__e, symnl, tmp20854)

tmp20856 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20855)


_ = tmp20856

tmp20857 := MakeNative1(func(__e *ControlFlow, Y1108 Obj) Obj {
return GNot(__e, Y1108)
panic("unreachable")
})

tmp20858 := GCons(__e, symnot, tmp20857)

tmp20859 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20858)


_ = tmp20859

tmp20860 := MakeNative1(func(__e *ControlFlow, Y1109 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1110 Obj) Obj {
return TailCall2(__e, SymFn(symnth), Y1109, Y1110)
panic("unreachable")
})
panic("unreachable")
})

tmp20861 := GCons(__e, symnth, tmp20860)

tmp20862 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20861)


_ = tmp20862

tmp20863 := MakeNative1(func(__e *ControlFlow, Y1111 Obj) Obj {
return PrimGuard1(__e, symn_1_6string, PrimNumberToString, Y1111)
panic("unreachable")
})

tmp20864 := GCons(__e, symn_1_6string, tmp20863)

tmp20865 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20864)


_ = tmp20865

tmp20866 := MakeNative1(func(__e *ControlFlow, Y1112 Obj) Obj {
return PrimGuard1(__e, symnumber_2, PrimIsNumber, Y1112)
panic("unreachable")
})

tmp20867 := GCons(__e, symnumber_2, tmp20866)

tmp20868 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20867)


_ = tmp20868

tmp20869 := MakeNative1(func(__e *ControlFlow, Y1113 Obj) Obj {
return TailCall1(__e, SymFn(symoccurs_1check), Y1113)
panic("unreachable")
})

tmp20870 := GCons(__e, symoccurs_1check, tmp20869)

tmp20871 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20870)


_ = tmp20871

tmp20872 := MakeNative1(func(__e *ControlFlow, Y1114 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1115 Obj) Obj {
return TailCall2(__e, SymFn(symoccurrences), Y1114, Y1115)
panic("unreachable")
})
panic("unreachable")
})

tmp20873 := GCons(__e, symoccurrences, tmp20872)

tmp20874 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20873)


_ = tmp20874

tmp20875 := MakeNative1(func(__e *ControlFlow, Y1116 Obj) Obj {
return TailCall1(__e, SymFn(symoccurs_1check), Y1116)
panic("unreachable")
})

tmp20876 := GCons(__e, symoccurs_1check, tmp20875)

tmp20877 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20876)


_ = tmp20877

tmp20878 := MakeNative1(func(__e *ControlFlow, Y1117 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1118 Obj) Obj {
return PrimGuard2(__e, symopen, PrimOpenStream, Y1117, Y1118)
panic("unreachable")
})
panic("unreachable")
})

tmp20879 := GCons(__e, symopen, tmp20878)

tmp20880 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20879)


_ = tmp20880

tmp20881 := MakeNative1(func(__e *ControlFlow, Y1119 Obj) Obj {
return TailCall1(__e, SymFn(symoptimise), Y1119)
panic("unreachable")
})

tmp20882 := GCons(__e, symoptimise, tmp20881)

tmp20883 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20882)


_ = tmp20883

tmp20884 := MakeNative1(func(__e *ControlFlow, Y1120 Obj) Obj {
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

tmp20886 := GCons(__e, symor, tmp20884)

tmp20887 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20886)


_ = tmp20887

tmp20888 := MakeNative1(func(__e *ControlFlow, Y1122 Obj) Obj {
return TailCall1(__e, SymFn(sympackage_2), Y1122)
panic("unreachable")
})

tmp20889 := GCons(__e, sympackage_2, tmp20888)

tmp20890 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20889)


_ = tmp20890

tmp20891 := MakeNative1(func(__e *ControlFlow, Y1123 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1124 Obj) Obj {
return PrimGuard2(__e, sympos, PrimPos, Y1123, Y1124)
panic("unreachable")
})
panic("unreachable")
})

tmp20892 := GCons(__e, sympos, tmp20891)

tmp20893 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20892)


_ = tmp20893

tmp20894 := MakeNative1(func(__e *ControlFlow, Y1125 Obj) Obj {
return TailCall1(__e, SymFn(sympreclude_1all_1but), Y1125)
panic("unreachable")
})

tmp20895 := GCons(__e, sympreclude_1all_1but, tmp20894)

tmp20896 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20895)


_ = tmp20896

tmp20897 := MakeNative1(func(__e *ControlFlow, Y1126 Obj) Obj {
return TailCall1(__e, SymFn(symprint), Y1126)
panic("unreachable")
})

tmp20898 := GCons(__e, symprint, tmp20897)

tmp20899 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20898)


_ = tmp20899

tmp20900 := MakeNative1(func(__e *ControlFlow, Y1127 Obj) Obj {
return TailCall1(__e, SymFn(symprofile), Y1127)
panic("unreachable")
})

tmp20901 := GCons(__e, symprofile, tmp20900)

tmp20902 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20901)


_ = tmp20902

tmp20903 := MakeNative1(func(__e *ControlFlow, Y1128 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4print_1prolog_1vector), Y1128)
panic("unreachable")
})

tmp20904 := GCons(__e, symshen_4print_1prolog_1vector, tmp20903)

tmp20905 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20904)


_ = tmp20905

tmp20906 := MakeNative1(func(__e *ControlFlow, Y1129 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4print_1freshterm), Y1129)
panic("unreachable")
})

tmp20907 := GCons(__e, symshen_4print_1freshterm, tmp20906)

tmp20908 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20907)


_ = tmp20908

tmp20909 := MakeNative1(func(__e *ControlFlow, Y1130 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4printF), Y1130)
panic("unreachable")
})

tmp20910 := GCons(__e, symshen_4printF, tmp20909)

tmp20911 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20910)


_ = tmp20911

tmp20912 := MakeNative1(func(__e *ControlFlow, Y1131 Obj) Obj {
return TailCall1(__e, SymFn(symprolog_1memory), Y1131)
panic("unreachable")
})

tmp20913 := GCons(__e, symprolog_1memory, tmp20912)

tmp20914 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20913)


_ = tmp20914

tmp20915 := MakeNative1(func(__e *ControlFlow, Y1132 Obj) Obj {
return TailCall1(__e, SymFn(symprofile_1results), Y1132)
panic("unreachable")
})

tmp20916 := GCons(__e, symprofile_1results, tmp20915)

tmp20917 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20916)


_ = tmp20917

tmp20918 := MakeNative1(func(__e *ControlFlow, Y1133 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1134 Obj) Obj {
return TailCall2(__e, SymFn(sympr), Y1133, Y1134)
panic("unreachable")
})
panic("unreachable")
})

tmp20919 := GCons(__e, sympr, tmp20918)

tmp20920 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20919)


_ = tmp20920

tmp20921 := MakeNative1(func(__e *ControlFlow, Y1135 Obj) Obj {
return TailCall1(__e, SymFn(symps), Y1135)
panic("unreachable")
})

tmp20922 := GCons(__e, symps, tmp20921)

tmp20923 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20922)


_ = tmp20923

tmp20924 := MakeNative1(func(__e *ControlFlow, Y1136 Obj) Obj {
return TailCall1(__e, SymFn(sympreclude), Y1136)
panic("unreachable")
})

tmp20925 := GCons(__e, sympreclude, tmp20924)

tmp20926 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20925)


_ = tmp20926

tmp20927 := MakeNative1(func(__e *ControlFlow, Y1137 Obj) Obj {
return TailCall1(__e, SymFn(sympreclude_1all_1but), Y1137)
panic("unreachable")
})

tmp20928 := GCons(__e, sympreclude_1all_1but, tmp20927)

tmp20929 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20928)


_ = tmp20929

tmp20930 := MakeNative1(func(__e *ControlFlow, Y1138 Obj) Obj {
return TailCall1(__e, SymFn(symprotect), Y1138)
panic("unreachable")
})

tmp20931 := GCons(__e, symprotect, tmp20930)

tmp20932 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20931)


_ = tmp20932

tmp20933 := MakeNative1(func(__e *ControlFlow, Y1139 Obj) Obj {
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

tmp20934 := GCons(__e, symput, tmp20933)

tmp20935 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20934)


_ = tmp20935

tmp20936 := MakeNative1(func(__e *ControlFlow, Y1143 Obj) Obj {
return PrimGuard1(__e, symread_1file_1as_1string, PrimReadFileAsString, Y1143)
panic("unreachable")
})

tmp20937 := GCons(__e, symread_1file_1as_1string, tmp20936)

tmp20938 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20937)


_ = tmp20938

tmp20939 := MakeNative1(func(__e *ControlFlow, Y1144 Obj) Obj {
return PrimGuard1(__e, symread_1file_1as_1bytelist, PrimReadFileAsByteList, Y1144)
panic("unreachable")
})

tmp20940 := GCons(__e, symread_1file_1as_1bytelist, tmp20939)

tmp20941 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20940)


_ = tmp20941

tmp20942 := MakeNative1(func(__e *ControlFlow, Y1145 Obj) Obj {
return TailCall1(__e, SymFn(symread_1file), Y1145)
panic("unreachable")
})

tmp20943 := GCons(__e, symread_1file, tmp20942)

tmp20944 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20943)


_ = tmp20944

tmp20945 := MakeNative1(func(__e *ControlFlow, Y1146 Obj) Obj {
return TailCall1(__e, SymFn(symread), Y1146)
panic("unreachable")
})

tmp20946 := GCons(__e, symread, tmp20945)

tmp20947 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20946)


_ = tmp20947

tmp20948 := MakeNative1(func(__e *ControlFlow, Y1147 Obj) Obj {
return PrimGuard1(__e, symread_1byte, PrimReadByte, Y1147)
panic("unreachable")
})

tmp20949 := GCons(__e, symread_1byte, tmp20948)

tmp20950 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20949)


_ = tmp20950

tmp20951 := MakeNative1(func(__e *ControlFlow, Y1148 Obj) Obj {
return TailCall1(__e, SymFn(symread_1from_1string), Y1148)
panic("unreachable")
})

tmp20952 := GCons(__e, symread_1from_1string, tmp20951)

tmp20953 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20952)


_ = tmp20953

tmp20954 := MakeNative1(func(__e *ControlFlow, Y1149 Obj) Obj {
return TailCall1(__e, SymFn(symread_1from_1string_1unprocessed), Y1149)
panic("unreachable")
})

tmp20955 := GCons(__e, symread_1from_1string_1unprocessed, tmp20954)

tmp20956 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20955)


_ = tmp20956

tmp20957 := MakeNative1(func(__e *ControlFlow, Y1150 Obj) Obj {
return TailCall1(__e, SymFn(symread_1unit_1string), Y1150)
panic("unreachable")
})

tmp20958 := GCons(__e, symread_1unit_1string, tmp20957)

tmp20959 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20958)


_ = tmp20959

tmp20960 := MakeNative1(func(__e *ControlFlow, Y1151 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1152 Obj) Obj {
return TailCall2(__e, SymFn(symshen_4register_1source_1form), Y1151, Y1152)
panic("unreachable")
})
panic("unreachable")
})

tmp20961 := GCons(__e, symshen_4register_1source_1form, tmp20960)

tmp20962 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20961)


_ = tmp20962

tmp20963 := MakeNative1(func(__e *ControlFlow, Y1153 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4unregister_1source_1form), Y1153)
panic("unreachable")
})

tmp20964 := GCons(__e, symshen_4unregister_1source_1form, tmp20963)

tmp20965 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20964)


_ = tmp20965

tmp20966 := MakeNative1(func(__e *ControlFlow, Y1154 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1155 Obj) Obj {
return TailCall2(__e, SymFn(symremove), Y1154, Y1155)
panic("unreachable")
})
panic("unreachable")
})

tmp20967 := GCons(__e, symremove, tmp20966)

tmp20968 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20967)


_ = tmp20968

tmp20969 := MakeNative1(func(__e *ControlFlow, Y1156 Obj) Obj {
return TailCall1(__e, SymFn(symreverse), Y1156)
panic("unreachable")
})

tmp20970 := GCons(__e, symreverse, tmp20969)

tmp20971 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20970)


_ = tmp20971

tmp20972 := MakeNative1(func(__e *ControlFlow, Y1157 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1158 Obj) Obj {
return PrimGuard2(__e, symset, PrimSet, Y1157, Y1158)
panic("unreachable")
})
panic("unreachable")
})

tmp20973 := GCons(__e, symset, tmp20972)

tmp20974 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20973)


_ = tmp20974

tmp20975 := MakeNative1(func(__e *ControlFlow, Y1159 Obj) Obj {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, Y1159)
panic("unreachable")
})

tmp20976 := GCons(__e, symsimple_1error, tmp20975)

tmp20977 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20976)


_ = tmp20977

tmp20978 := MakeNative1(func(__e *ControlFlow, Y1160 Obj) Obj {
return TailCall1(__e, SymFn(symsnd), Y1160)
panic("unreachable")
})

tmp20979 := GCons(__e, symsnd, tmp20978)

tmp20980 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20979)


_ = tmp20980

tmp20981 := MakeNative1(func(__e *ControlFlow, Y1161 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1162 Obj) Obj {
return TailCall2(__e, SymFn(symspecialise), Y1161, Y1162)
panic("unreachable")
})
panic("unreachable")
})

tmp20982 := GCons(__e, symspecialise, tmp20981)

tmp20983 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20982)


_ = tmp20983

tmp20984 := MakeNative1(func(__e *ControlFlow, Y1163 Obj) Obj {
return TailCall1(__e, SymFn(symspy), Y1163)
panic("unreachable")
})

tmp20985 := GCons(__e, symspy, tmp20984)

tmp20986 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20985)


_ = tmp20986

tmp20987 := MakeNative1(func(__e *ControlFlow, Y1164 Obj) Obj {
return TailCall1(__e, SymFn(symstep), Y1164)
panic("unreachable")
})

tmp20988 := GCons(__e, symstep, tmp20987)

tmp20989 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20988)


_ = tmp20989

tmp20990 := MakeNative1(func(__e *ControlFlow, Y1165 Obj) Obj {
return PrimGuard1(__e, symstr, PrimStr, Y1165)
panic("unreachable")
})

tmp20991 := GCons(__e, symstr, tmp20990)

tmp20992 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20991)


_ = tmp20992

tmp20993 := MakeNative1(func(__e *ControlFlow, Y1166 Obj) Obj {
return PrimGuard1(__e, symstring_1_6n, PrimStringToNumber, Y1166)
panic("unreachable")
})

tmp20994 := GCons(__e, symstring_1_6n, tmp20993)

tmp20995 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20994)


_ = tmp20995

tmp20996 := MakeNative1(func(__e *ControlFlow, Y1167 Obj) Obj {
return TailCall1(__e, SymFn(symstring_1_6symbol), Y1167)
panic("unreachable")
})

tmp20997 := GCons(__e, symstring_1_6symbol, tmp20996)

tmp20998 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp20997)


_ = tmp20998

tmp20999 := MakeNative1(func(__e *ControlFlow, Y1168 Obj) Obj {
return PrimGuard1(__e, symstring_2, PrimIsString, Y1168)
panic("unreachable")
})

tmp21000 := GCons(__e, symstring_2, tmp20999)

tmp21001 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21000)


_ = tmp21001

tmp21002 := MakeNative1(func(__e *ControlFlow, Y1169 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1170 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1171 Obj) Obj {
return TailCall3(__e, SymFn(symsubst), Y1169, Y1170, Y1171)
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp21003 := GCons(__e, symsubst, tmp21002)

tmp21004 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21003)


_ = tmp21004

tmp21005 := MakeNative1(func(__e *ControlFlow, Y1172 Obj) Obj {
return TailCall1(__e, SymFn(symsum), Y1172)
panic("unreachable")
})

tmp21006 := GCons(__e, symsum, tmp21005)

tmp21007 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21006)


_ = tmp21007

tmp21008 := MakeNative1(func(__e *ControlFlow, Y1173 Obj) Obj {
return PrimGuard1(__e, symsymbol_2, PrimIsSymbol, Y1173)
panic("unreachable")
})

tmp21009 := GCons(__e, symsymbol_2, tmp21008)

tmp21010 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21009)


_ = tmp21010

tmp21011 := MakeNative1(func(__e *ControlFlow, Y1174 Obj) Obj {
return TailCall1(__e, SymFn(symsystemf), Y1174)
panic("unreachable")
})

tmp21012 := GCons(__e, symsystemf, tmp21011)

tmp21013 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21012)


_ = tmp21013

tmp21014 := MakeNative1(func(__e *ControlFlow, Y1175 Obj) Obj {
return TailCall1(__e, SymFn(symtail), Y1175)
panic("unreachable")
})

tmp21015 := GCons(__e, symtail, tmp21014)

tmp21016 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21015)


_ = tmp21016

tmp21017 := MakeNative1(func(__e *ControlFlow, Y1176 Obj) Obj {
return GTl(__e, Y1176)
panic("unreachable")
})

tmp21018 := GCons(__e, symtl, tmp21017)

tmp21019 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21018)


_ = tmp21019

tmp21020 := MakeNative1(func(__e *ControlFlow, Y1177 Obj) Obj {
return TailCall1(__e, SymFn(symtc), Y1177)
panic("unreachable")
})

tmp21021 := GCons(__e, symtc, tmp21020)

tmp21022 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21021)


_ = tmp21022

tmp21023 := MakeNative1(func(__e *ControlFlow, Y1178 Obj) Obj {
return TailCall1(__e, SymFn(symthaw), Y1178)
panic("unreachable")
})

tmp21024 := GCons(__e, symthaw, tmp21023)

tmp21025 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21024)


_ = tmp21025

tmp21026 := MakeNative1(func(__e *ControlFlow, Y1179 Obj) Obj {
return PrimGuard1(__e, symtlstr, PrimTailString, Y1179)
panic("unreachable")
})

tmp21027 := GCons(__e, symtlstr, tmp21026)

tmp21028 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21027)


_ = tmp21028

tmp21029 := MakeNative1(func(__e *ControlFlow, Y1180 Obj) Obj {
return TailCall1(__e, SymFn(symtrack), Y1180)
panic("unreachable")
})

tmp21030 := GCons(__e, symtrack, tmp21029)

tmp21031 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21030)


_ = tmp21031

tmp21032 := MakeNative1(func(__e *ControlFlow, Y1181 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1182 Obj) Obj {
tmp21033 := MakeNative0(func(__e *ControlFlow) Obj {
return Y1181
panic("unreachable")
})

return TailCall2(__e, try_1catch, tmp21033, Y1182)


panic("unreachable")
})
panic("unreachable")
})

tmp21034 := GCons(__e, symtrap_1error, tmp21032)

tmp21035 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21034)


_ = tmp21035

tmp21036 := MakeNative1(func(__e *ControlFlow, Y1183 Obj) Obj {
return TailCall1(__e, SymFn(symtuple_2), Y1183)
panic("unreachable")
})

tmp21037 := GCons(__e, symtuple_2, tmp21036)

tmp21038 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21037)


_ = tmp21038

tmp21039 := MakeNative1(func(__e *ControlFlow, Y1184 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1185 Obj) Obj {
return Y1185
panic("unreachable")
})
panic("unreachable")
})

tmp21040 := GCons(__e, symtype, tmp21039)

tmp21041 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21040)


_ = tmp21041

tmp21042 := MakeNative1(func(__e *ControlFlow, Y1186 Obj) Obj {
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

tmp21043 := GCons(__e, symreturn, tmp21042)

tmp21044 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21043)


_ = tmp21044

tmp21045 := MakeNative1(func(__e *ControlFlow, Y1191 Obj) Obj {
return TailCall1(__e, SymFn(symunabsolute), Y1191)
panic("unreachable")
})

tmp21046 := GCons(__e, symunabsolute, tmp21045)

tmp21047 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21046)


_ = tmp21047

tmp21048 := MakeNative1(func(__e *ControlFlow, Y1192 Obj) Obj {
return TailCall1(__e, SymFn(symundefmacro), Y1192)
panic("unreachable")
})

tmp21049 := GCons(__e, symundefmacro, tmp21048)

tmp21050 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21049)


_ = tmp21050

tmp21051 := MakeNative1(func(__e *ControlFlow, Y1193 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1194 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1195 Obj) Obj {
return TailCall3(__e, SymFn(symunput), Y1193, Y1194, Y1195)
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp21052 := GCons(__e, symunput, tmp21051)

tmp21053 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21052)


_ = tmp21053

tmp21054 := MakeNative1(func(__e *ControlFlow, Y1196 Obj) Obj {
return TailCall1(__e, SymFn(symunprofile), Y1196)
panic("unreachable")
})

tmp21055 := GCons(__e, symunprofile, tmp21054)

tmp21056 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21055)


_ = tmp21056

tmp21057 := MakeNative1(func(__e *ControlFlow, Y1197 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1198 Obj) Obj {
return TailCall2(__e, SymFn(symunion), Y1197, Y1198)
panic("unreachable")
})
panic("unreachable")
})

tmp21058 := GCons(__e, symunion, tmp21057)

tmp21059 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21058)


_ = tmp21059

tmp21060 := MakeNative1(func(__e *ControlFlow, Y1199 Obj) Obj {
return TailCall1(__e, SymFn(symuntrack), Y1199)
panic("unreachable")
})

tmp21061 := GCons(__e, symuntrack, tmp21060)

tmp21062 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21061)


_ = tmp21062

tmp21063 := MakeNative1(func(__e *ControlFlow, Y1200 Obj) Obj {
return TailCall1(__e, SymFn(symundefmacro), Y1200)
panic("unreachable")
})

tmp21064 := GCons(__e, symundefmacro, tmp21063)

tmp21065 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21064)


_ = tmp21065

tmp21066 := MakeNative1(func(__e *ControlFlow, Y1201 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1202 Obj) Obj {
return TailCall2(__e, SymFn(symupdate_1lambda_1table), Y1201, Y1202)
panic("unreachable")
})
panic("unreachable")
})

tmp21067 := GCons(__e, symupdate_1lambda_1table, tmp21066)

tmp21068 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21067)


_ = tmp21068

tmp21069 := MakeNative1(func(__e *ControlFlow, Y1203 Obj) Obj {
return TailCall1(__e, SymFn(symvector), Y1203)
panic("unreachable")
})

tmp21070 := GCons(__e, symvector, tmp21069)

tmp21071 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21070)


_ = tmp21071

tmp21072 := MakeNative1(func(__e *ControlFlow, Y1204 Obj) Obj {
return TailCall1(__e, SymFn(symvector_2), Y1204)
panic("unreachable")
})

tmp21073 := GCons(__e, symvector_2, tmp21072)

tmp21074 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21073)


_ = tmp21074

tmp21075 := MakeNative1(func(__e *ControlFlow, Y1205 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1206 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1207 Obj) Obj {
return TailCall3(__e, SymFn(symvector_1_6), Y1205, Y1206, Y1207)
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp21076 := GCons(__e, symvector_1_6, tmp21075)

tmp21077 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21076)


_ = tmp21077

tmp21078 := MakeNative1(func(__e *ControlFlow, Y1208 Obj) Obj {
return GValue(__e, Y1208)
panic("unreachable")
})

tmp21079 := GCons(__e, symvalue, tmp21078)

tmp21080 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21079)


_ = tmp21080

tmp21081 := MakeNative1(func(__e *ControlFlow, Y1209 Obj) Obj {
return PrimGuard1(__e, symvariable_2, PrimIsVariable, Y1209)
panic("unreachable")
})

tmp21082 := GCons(__e, symvariable_2, tmp21081)

tmp21083 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21082)


_ = tmp21083

tmp21084 := MakeNative1(func(__e *ControlFlow, Y1210 Obj) Obj {
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

tmp21085 := GCons(__e, symvar_2, tmp21084)

tmp21086 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21085)


_ = tmp21086

tmp21087 := MakeNative1(func(__e *ControlFlow, Y1215 Obj) Obj {
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

tmp21088 := GCons(__e, symwhen, tmp21087)

tmp21089 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21088)


_ = tmp21089

tmp21090 := MakeNative1(func(__e *ControlFlow, Y1220 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1221 Obj) Obj {
return PrimGuard2(__e, symwrite_1byte, PrimWriteByte, Y1220, Y1221)
panic("unreachable")
})
panic("unreachable")
})

tmp21091 := GCons(__e, symwrite_1byte, tmp21090)

tmp21092 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21091)


_ = tmp21092

tmp21093 := MakeNative1(func(__e *ControlFlow, Y1222 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1223 Obj) Obj {
return TailCall2(__e, SymFn(symwrite_1string), Y1222, Y1223)
panic("unreachable")
})
panic("unreachable")
})

tmp21094 := GCons(__e, symwrite_1string, tmp21093)

tmp21095 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21094)


_ = tmp21095

tmp21096 := MakeNative1(func(__e *ControlFlow, Y1224 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1225 Obj) Obj {
return TailCall2(__e, SymFn(symwrite_1to_1file), Y1224, Y1225)
panic("unreachable")
})
panic("unreachable")
})

tmp21097 := GCons(__e, symwrite_1to_1file, tmp21096)

tmp21098 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21097)


_ = tmp21098

tmp21099 := MakeNative1(func(__e *ControlFlow, Y1226 Obj) Obj {
return TailCall1(__e, SymFn(symy_1or_1n_2), Y1226)
panic("unreachable")
})

tmp21100 := GCons(__e, symy_1or_1n_2, tmp21099)

tmp21101 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21100)


_ = tmp21101

tmp21102 := MakeNative1(func(__e *ControlFlow, Y1227 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1228 Obj) Obj {
return GAdd(__e, Y1227, Y1228)
panic("unreachable")
})
panic("unreachable")
})

tmp21103 := GCons(__e, sym_7, tmp21102)

tmp21104 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21103)


_ = tmp21104

tmp21105 := MakeNative1(func(__e *ControlFlow, Y1229 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1230 Obj) Obj {
return GMul(__e, Y1229, Y1230)
panic("unreachable")
})
panic("unreachable")
})

tmp21106 := GCons(__e, sym_d, tmp21105)

tmp21107 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21106)


_ = tmp21107

tmp21108 := MakeNative1(func(__e *ControlFlow, Y1231 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1232 Obj) Obj {
return GDiv(__e, Y1231, Y1232)
panic("unreachable")
})
panic("unreachable")
})

tmp21109 := GCons(__e, sym_c, tmp21108)

tmp21110 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21109)


_ = tmp21110

tmp21111 := MakeNative1(func(__e *ControlFlow, Y1233 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1234 Obj) Obj {
return GSub(__e, Y1233, Y1234)
panic("unreachable")
})
panic("unreachable")
})

tmp21112 := GCons(__e, sym_1, tmp21111)

tmp21113 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21112)


_ = tmp21113

tmp21114 := MakeNative1(func(__e *ControlFlow, Y1235 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1236 Obj) Obj {
return TailCall2(__e, SymFn(sym_a_a), Y1235, Y1236)
panic("unreachable")
})
panic("unreachable")
})

tmp21115 := GCons(__e, sym_a_a, tmp21114)

tmp21116 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21115)


_ = tmp21116

tmp21117 := MakeNative1(func(__e *ControlFlow, Y1237 Obj) Obj {
return TailCall1(__e, SymFn(sym_5e_6), Y1237)
panic("unreachable")
})

tmp21118 := GCons(__e, sym_5e_6, tmp21117)

tmp21119 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21118)


_ = tmp21119

tmp21120 := MakeNative1(func(__e *ControlFlow, Y1238 Obj) Obj {
return TailCall1(__e, SymFn(sym_5end_6), Y1238)
panic("unreachable")
})

tmp21121 := GCons(__e, sym_5end_6, tmp21120)

tmp21122 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21121)


_ = tmp21122

tmp21123 := MakeNative1(func(__e *ControlFlow, Y1239 Obj) Obj {
return TailCall1(__e, SymFn(sym_5_b_6), Y1239)
panic("unreachable")
})

tmp21124 := GCons(__e, sym_5_b_6, tmp21123)

tmp21125 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21124)


_ = tmp21125

tmp21126 := MakeNative1(func(__e *ControlFlow, Y1240 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1241 Obj) Obj {
return TailCall2(__e, SymFn(sym_8p), Y1240, Y1241)
panic("unreachable")
})
panic("unreachable")
})

tmp21127 := GCons(__e, sym_8p, tmp21126)

tmp21128 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21127)


_ = tmp21128

tmp21129 := MakeNative1(func(__e *ControlFlow, Y1242 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1243 Obj) Obj {
return TailCall2(__e, SymFn(sym_8v), Y1242, Y1243)
panic("unreachable")
})
panic("unreachable")
})

tmp21130 := GCons(__e, sym_8v, tmp21129)

tmp21131 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21130)


_ = tmp21131

tmp21132 := MakeNative1(func(__e *ControlFlow, Y1244 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Y1245 Obj) Obj {
return TailCall2(__e, SymFn(sym_8s), Y1244, Y1245)
panic("unreachable")
})
panic("unreachable")
})

tmp21133 := GCons(__e, sym_8s, tmp21132)

return TailCall1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp21133)


panic("unreachable")
})

tmp21134 := Call2(__e, ns2_1set, symshen_4initialise_1lambda_1forms, tmp20581)


_ = tmp21134

tmp21135 := MakeNative0(func(__e *ControlFlow) Obj {
tmp21136 := Call0(__e, SymFn(symshen_4initialise_1environment))


_ = tmp21136

tmp21137 := Call0(__e, SymFn(symshen_4initialise_1lambda_1forms))


_ = tmp21137

return TailCall0(__e, SymFn(symshen_4initialise_1signedfuncs))


panic("unreachable")
})

return TailCall2(__e, ns2_1set, symshen_4initialise, tmp21135)




panic("unreachable")
})

