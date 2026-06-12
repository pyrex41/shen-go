package main

import . "github.com/tiancaiamao/shen-go/kl"

var InitMain = MakeNative(func(__e *ControlFlow) {
tmp17077 := MakeNative(func(__e *ControlFlow) {
tmp17078 := PrimSet(symshen_4_dhistory_d, Nil)

_ = tmp17078

tmp17079 := PrimSet(symshen_4_dtc_d, False)

_ = tmp17079

tmp17080 := Call(__e, PrimFunc(symshen_4dict), MakeNumber(20000))


tmp17081 := PrimSet(sym_dproperty_1vector_d, tmp17080)

_ = tmp17081

tmp17082 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
__e.TailApply(PrimFunc(symshen_4macros), X)
return
}, 1)

tmp17083 := PrimCons(symshen_4macros, tmp17082)

tmp17084 := PrimCons(tmp17083, Nil)

tmp17085 := PrimSet(sym_dmacros_d, tmp17084)

_ = tmp17085

tmp17086 := PrimSet(symshen_4_dgensym_d, MakeNumber(0))

_ = tmp17086

tmp17087 := PrimSet(symshen_4_dtracking_d, Nil)

_ = tmp17087

tmp17088 := PrimSet(symshen_4_dprofiled_d, Nil)

_ = tmp17088

tmp17089 := PrimCons(symtype, Nil)

tmp17090 := PrimCons(syminput_7, tmp17089)

tmp17091 := PrimCons(symopen, tmp17090)

tmp17092 := PrimCons(symset, tmp17091)

tmp17093 := PrimCons(symwhere, tmp17092)

tmp17094 := PrimCons(symlet, tmp17093)

tmp17095 := PrimCons(symlambda, tmp17094)

tmp17096 := PrimCons(symcons, tmp17095)

tmp17097 := PrimCons(sym_8v, tmp17096)

tmp17098 := PrimCons(sym_8s, tmp17097)

tmp17099 := PrimCons(sym_8p, tmp17098)

tmp17100 := PrimSet(symshen_4_dspecial_d, tmp17099)

_ = tmp17100

tmp17101 := PrimSet(symshen_4_dextraspecial_d, Nil)

_ = tmp17101

tmp17102 := PrimSet(symshen_4_dspy_d, False)

_ = tmp17102

tmp17103 := PrimSet(symshen_4_ddatatypes_d, Nil)

_ = tmp17103

tmp17104 := PrimSet(symshen_4_dalldatatypes_d, Nil)

_ = tmp17104

tmp17105 := PrimSet(symshen_4_dshen_1type_1theory_1enabled_2_d, True)

_ = tmp17105

tmp17106 := PrimSet(symshen_4_dpackage_d, symnull)

_ = tmp17106

tmp17107 := PrimSet(symshen_4_dsynonyms_d, Nil)

_ = tmp17107

tmp17108 := PrimSet(symshen_4_dsystem_d, Nil)

_ = tmp17108

tmp17109 := PrimSet(symshen_4_doccurs_d, True)

_ = tmp17109

tmp17110 := PrimSet(symshen_4_dfactorise_2_d, False)

_ = tmp17110

tmp17111 := PrimSet(symshen_4_dmaxinferences_d, MakeNumber(1000000))

_ = tmp17111

tmp17112 := PrimSet(sym_dmaximum_1print_1sequence_1size_d, MakeNumber(20))

_ = tmp17112

tmp17113 := PrimSet(symshen_4_dcall_d, MakeNumber(0))

_ = tmp17113

tmp17114 := PrimSet(symshen_4_dinfs_d, MakeNumber(0))

_ = tmp17114

tmp17115 := PrimSet(sym_dhush_d, False)

_ = tmp17115

tmp17116 := PrimSet(symshen_4_doptimise_d, False)

_ = tmp17116

tmp17117 := PrimSet(sym_dversion_d, MakeString("41.2"))

_ = tmp17117

tmp17118 := PrimSet(symshen_4_dnames_d, Nil)

_ = tmp17118

tmp17119 := PrimSet(symshen_4_dstep_d, False)

_ = tmp17119

tmp17120 := PrimSet(symshen_4_dit_d, MakeString(""))

_ = tmp17120

tmp17121 := PrimSet(symshen_4_dresidue_d, Nil)

_ = tmp17121

tmp17122 := PrimSet(sym_dabsolute_d, Nil)

_ = tmp17122

tmp17123 := PrimSet(symshen_4_dprolog_1memory_d, MakeNumber(1000))

_ = tmp17123

tmp17124 := PrimSet(symshen_4_dloading_2_d, False)

_ = tmp17124

tmp17125 := PrimSet(symshen_4_duserdefs_d, Nil)

_ = tmp17125

tmp17126 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
__e.Return(X)
return
}, 1)

tmp17127 := PrimSet(symshen_4_ddemodulation_1function_d, tmp17126)

_ = tmp17127

tmp17128 := PrimSet(symshen_4_dcustom_1pattern_1compiler_d, False)

_ = tmp17128

tmp17129 := PrimSet(symshen_4_dcustom_1pattern_1reducer_d, False)

_ = tmp17129

tmp17132 := Call(__e, PrimFunc(symbound_2), sym_dhome_1directory_d)


tmp17133 := PrimNot(tmp17132)

var ifres17130 Obj

if True == tmp17133 {
tmp17131 := PrimSet(sym_dhome_1directory_d, MakeString(""))

ifres17130 = tmp17131


} else {
ifres17130 = symshen_4skip


}

_ = ifres17130

tmp17137 := Call(__e, PrimFunc(symbound_2), sym_dsterror_d)


tmp17138 := PrimNot(tmp17137)

var ifres17134 Obj

if True == tmp17138 {
tmp17135 := PrimValue(sym_dstoutput_d)

tmp17136 := PrimSet(sym_dsterror_d, tmp17135)

ifres17134 = tmp17136


} else {
ifres17134 = symshen_4skip


}

_ = ifres17134

tmp17139 := Call(__e, PrimFunc(symprolog_1memory), MakeNumber(10000))


_ = tmp17139

tmp17140 := PrimSet(symshen_4_dloading_2_d, False)

_ = tmp17140

tmp17141 := PrimCons(MakeNumber(2), Nil)

tmp17142 := PrimCons(sym_8s, tmp17141)

tmp17143 := PrimCons(MakeNumber(2), tmp17142)

tmp17144 := PrimCons(sym_8v, tmp17143)

tmp17145 := PrimCons(MakeNumber(2), tmp17144)

tmp17146 := PrimCons(sym_8p, tmp17145)

tmp17147 := PrimCons(MakeNumber(1), tmp17146)

tmp17148 := PrimCons(sym_5_b_6, tmp17147)

tmp17149 := PrimCons(MakeNumber(1), tmp17148)

tmp17150 := PrimCons(sym_5end_6, tmp17149)

tmp17151 := PrimCons(MakeNumber(1), tmp17150)

tmp17152 := PrimCons(sym_5e_6, tmp17151)

tmp17153 := PrimCons(MakeNumber(2), tmp17152)

tmp17154 := PrimCons(sym_a_a, tmp17153)

tmp17155 := PrimCons(MakeNumber(2), tmp17154)

tmp17156 := PrimCons(sym_1, tmp17155)

tmp17157 := PrimCons(MakeNumber(2), tmp17156)

tmp17158 := PrimCons(sym_c, tmp17157)

tmp17159 := PrimCons(MakeNumber(2), tmp17158)

tmp17160 := PrimCons(sym_d, tmp17159)

tmp17161 := PrimCons(MakeNumber(2), tmp17160)

tmp17162 := PrimCons(sym_7, tmp17161)

tmp17163 := PrimCons(MakeNumber(1), tmp17162)

tmp17164 := PrimCons(symy_1or_1n_2, tmp17163)

tmp17165 := PrimCons(MakeNumber(2), tmp17164)

tmp17166 := PrimCons(symwrite_1to_1file, tmp17165)

tmp17167 := PrimCons(MakeNumber(2), tmp17166)

tmp17168 := PrimCons(symwrite_1byte, tmp17167)

tmp17169 := PrimCons(MakeNumber(5), tmp17168)

tmp17170 := PrimCons(symwhen, tmp17169)

tmp17171 := PrimCons(MakeNumber(0), tmp17170)

tmp17172 := PrimCons(symversion, tmp17171)

tmp17173 := PrimCons(MakeNumber(5), tmp17172)

tmp17174 := PrimCons(symvar_2, tmp17173)

tmp17175 := PrimCons(MakeNumber(1), tmp17174)

tmp17176 := PrimCons(symvariable_2, tmp17175)

tmp17177 := PrimCons(MakeNumber(1), tmp17176)

tmp17178 := PrimCons(symvalue, tmp17177)

tmp17179 := PrimCons(MakeNumber(3), tmp17178)

tmp17180 := PrimCons(symvector_1_6, tmp17179)

tmp17181 := PrimCons(MakeNumber(1), tmp17180)

tmp17182 := PrimCons(symvector_2, tmp17181)

tmp17183 := PrimCons(MakeNumber(1), tmp17182)

tmp17184 := PrimCons(symvector, tmp17183)

tmp17185 := PrimCons(MakeNumber(0), tmp17184)

tmp17186 := PrimCons(symuserdefs, tmp17185)

tmp17187 := PrimCons(MakeNumber(2), tmp17186)

tmp17188 := PrimCons(symupdate_1lambda_1table, tmp17187)

tmp17189 := PrimCons(MakeNumber(1), tmp17188)

tmp17190 := PrimCons(symundefmacro, tmp17189)

tmp17191 := PrimCons(MakeNumber(1), tmp17190)

tmp17192 := PrimCons(symuntrack, tmp17191)

tmp17193 := PrimCons(MakeNumber(2), tmp17192)

tmp17194 := PrimCons(symunion, tmp17193)

tmp17195 := PrimCons(MakeNumber(1), tmp17194)

tmp17196 := PrimCons(symunprofile, tmp17195)

tmp17197 := PrimCons(MakeNumber(3), tmp17196)

tmp17198 := PrimCons(symunput, tmp17197)

tmp17199 := PrimCons(MakeNumber(1), tmp17198)

tmp17200 := PrimCons(symundefmacro, tmp17199)

tmp17201 := PrimCons(MakeNumber(1), tmp17200)

tmp17202 := PrimCons(symunabsolute, tmp17201)

tmp17203 := PrimCons(MakeNumber(5), tmp17202)

tmp17204 := PrimCons(symreturn, tmp17203)

tmp17205 := PrimCons(MakeNumber(2), tmp17204)

tmp17206 := PrimCons(symtype, tmp17205)

tmp17207 := PrimCons(MakeNumber(1), tmp17206)

tmp17208 := PrimCons(symtuple_2, tmp17207)

tmp17209 := PrimCons(MakeNumber(2), tmp17208)

tmp17210 := PrimCons(symtrap_1error, tmp17209)

tmp17211 := PrimCons(MakeNumber(0), tmp17210)

tmp17212 := PrimCons(symtracked, tmp17211)

tmp17213 := PrimCons(MakeNumber(1), tmp17212)

tmp17214 := PrimCons(symtrack, tmp17213)

tmp17215 := PrimCons(MakeNumber(1), tmp17214)

tmp17216 := PrimCons(symtlstr, tmp17215)

tmp17217 := PrimCons(MakeNumber(1), tmp17216)

tmp17218 := PrimCons(symthaw, tmp17217)

tmp17219 := PrimCons(MakeNumber(0), tmp17218)

tmp17220 := PrimCons(symtc_2, tmp17219)

tmp17221 := PrimCons(MakeNumber(1), tmp17220)

tmp17222 := PrimCons(symtc, tmp17221)

tmp17223 := PrimCons(MakeNumber(1), tmp17222)

tmp17224 := PrimCons(symtl, tmp17223)

tmp17225 := PrimCons(MakeNumber(1), tmp17224)

tmp17226 := PrimCons(symtail, tmp17225)

tmp17227 := PrimCons(MakeNumber(1), tmp17226)

tmp17228 := PrimCons(symsystemf, tmp17227)

tmp17229 := PrimCons(MakeNumber(1), tmp17228)

tmp17230 := PrimCons(symsymbol_2, tmp17229)

tmp17231 := PrimCons(MakeNumber(1), tmp17230)

tmp17232 := PrimCons(symsum, tmp17231)

tmp17233 := PrimCons(MakeNumber(3), tmp17232)

tmp17234 := PrimCons(symsubst, tmp17233)

tmp17235 := PrimCons(MakeNumber(1), tmp17234)

tmp17236 := PrimCons(symstring_2, tmp17235)

tmp17237 := PrimCons(MakeNumber(1), tmp17236)

tmp17238 := PrimCons(symstring_1_6symbol, tmp17237)

tmp17239 := PrimCons(MakeNumber(1), tmp17238)

tmp17240 := PrimCons(symstring_1_6n, tmp17239)

tmp17241 := PrimCons(MakeNumber(1), tmp17240)

tmp17242 := PrimCons(symstr, tmp17241)

tmp17243 := PrimCons(MakeNumber(0), tmp17242)

tmp17244 := PrimCons(symstoutput, tmp17243)

tmp17245 := PrimCons(MakeNumber(0), tmp17244)

tmp17246 := PrimCons(symstinput, tmp17245)

tmp17247 := PrimCons(MakeNumber(0), tmp17246)

tmp17248 := PrimCons(symshen_4step_2, tmp17247)

tmp17249 := PrimCons(MakeNumber(1), tmp17248)

tmp17250 := PrimCons(symstep, tmp17249)

tmp17251 := PrimCons(MakeNumber(0), tmp17250)

tmp17252 := PrimCons(symshen_4spy_2, tmp17251)

tmp17253 := PrimCons(MakeNumber(1), tmp17252)

tmp17254 := PrimCons(symspy, tmp17253)

tmp17255 := PrimCons(MakeNumber(2), tmp17254)

tmp17256 := PrimCons(symspecialise, tmp17255)

tmp17257 := PrimCons(MakeNumber(1), tmp17256)

tmp17258 := PrimCons(symsnd, tmp17257)

tmp17259 := PrimCons(MakeNumber(1), tmp17258)

tmp17260 := PrimCons(symsimple_1error, tmp17259)

tmp17261 := PrimCons(MakeNumber(2), tmp17260)

tmp17262 := PrimCons(symset, tmp17261)

tmp17263 := PrimCons(MakeNumber(1), tmp17262)

tmp17264 := PrimCons(symreverse, tmp17263)

tmp17265 := PrimCons(MakeNumber(2), tmp17264)

tmp17266 := PrimCons(symremove, tmp17265)

tmp17267 := PrimCons(MakeNumber(0), tmp17266)

tmp17268 := PrimCons(symrelease, tmp17267)

tmp17269 := PrimCons(MakeNumber(1), tmp17268)

tmp17270 := PrimCons(symreceive, tmp17269)

tmp17271 := PrimCons(MakeNumber(1), tmp17270)

tmp17272 := PrimCons(symshen_4read_1unit_1string, tmp17271)

tmp17273 := PrimCons(MakeNumber(1), tmp17272)

tmp17274 := PrimCons(symread_1from_1string_1unprocessed, tmp17273)

tmp17275 := PrimCons(MakeNumber(1), tmp17274)

tmp17276 := PrimCons(symread_1from_1string, tmp17275)

tmp17277 := PrimCons(MakeNumber(1), tmp17276)

tmp17278 := PrimCons(symread_1byte, tmp17277)

tmp17279 := PrimCons(MakeNumber(1), tmp17278)

tmp17280 := PrimCons(symread, tmp17279)

tmp17281 := PrimCons(MakeNumber(1), tmp17280)

tmp17282 := PrimCons(symread_1file, tmp17281)

tmp17283 := PrimCons(MakeNumber(1), tmp17282)

tmp17284 := PrimCons(symread_1file_1as_1bytelist, tmp17283)

tmp17285 := PrimCons(MakeNumber(1), tmp17284)

tmp17286 := PrimCons(symread_1file_1as_1string, tmp17285)

tmp17287 := PrimCons(MakeNumber(4), tmp17286)

tmp17288 := PrimCons(symput, tmp17287)

tmp17289 := PrimCons(MakeNumber(1), tmp17288)

tmp17290 := PrimCons(symprotect, tmp17289)

tmp17291 := PrimCons(MakeNumber(1), tmp17290)

tmp17292 := PrimCons(sympreclude_1all_1but, tmp17291)

tmp17293 := PrimCons(MakeNumber(1), tmp17292)

tmp17294 := PrimCons(sympreclude, tmp17293)

tmp17295 := PrimCons(MakeNumber(1), tmp17294)

tmp17296 := PrimCons(symps, tmp17295)

tmp17297 := PrimCons(MakeNumber(2), tmp17296)

tmp17298 := PrimCons(sympr, tmp17297)

tmp17299 := PrimCons(MakeNumber(1), tmp17298)

tmp17300 := PrimCons(symprofile_1results, tmp17299)

tmp17301 := PrimCons(MakeNumber(1), tmp17300)

tmp17302 := PrimCons(symprolog_1memory, tmp17301)

tmp17303 := PrimCons(MakeNumber(1), tmp17302)

tmp17304 := PrimCons(symshen_4printF, tmp17303)

tmp17305 := PrimCons(MakeNumber(1), tmp17304)

tmp17306 := PrimCons(symshen_4print_1freshterm, tmp17305)

tmp17307 := PrimCons(MakeNumber(1), tmp17306)

tmp17308 := PrimCons(symshen_4print_1prolog_1vector, tmp17307)

tmp17309 := PrimCons(MakeNumber(1), tmp17308)

tmp17310 := PrimCons(symprofile, tmp17309)

tmp17311 := PrimCons(MakeNumber(1), tmp17310)

tmp17312 := PrimCons(symprint, tmp17311)

tmp17313 := PrimCons(MakeNumber(1), tmp17312)

tmp17314 := PrimCons(sympreclude_1all_1but, tmp17313)

tmp17315 := PrimCons(MakeNumber(2), tmp17314)

tmp17316 := PrimCons(sympos, tmp17315)

tmp17317 := PrimCons(MakeNumber(0), tmp17316)

tmp17318 := PrimCons(symporters, tmp17317)

tmp17319 := PrimCons(MakeNumber(0), tmp17318)

tmp17320 := PrimCons(symport, tmp17319)

tmp17321 := PrimCons(MakeNumber(1), tmp17320)

tmp17322 := PrimCons(sympackage_2, tmp17321)

tmp17323 := PrimCons(MakeNumber(3), tmp17322)

tmp17324 := PrimCons(sympackage, tmp17323)

tmp17325 := PrimCons(MakeNumber(0), tmp17324)

tmp17326 := PrimCons(symos, tmp17325)

tmp17327 := PrimCons(MakeNumber(2), tmp17326)

tmp17328 := PrimCons(symor, tmp17327)

tmp17329 := PrimCons(MakeNumber(0), tmp17328)

tmp17330 := PrimCons(symoptimise_2, tmp17329)

tmp17331 := PrimCons(MakeNumber(1), tmp17330)

tmp17332 := PrimCons(symoptimise, tmp17331)

tmp17333 := PrimCons(MakeNumber(2), tmp17332)

tmp17334 := PrimCons(symopen, tmp17333)

tmp17335 := PrimCons(MakeNumber(1), tmp17334)

tmp17336 := PrimCons(symoccurs_1check, tmp17335)

tmp17337 := PrimCons(MakeNumber(0), tmp17336)

tmp17338 := PrimCons(symoccurs_2, tmp17337)

tmp17339 := PrimCons(MakeNumber(2), tmp17338)

tmp17340 := PrimCons(symoccurrences, tmp17339)

tmp17341 := PrimCons(MakeNumber(1), tmp17340)

tmp17342 := PrimCons(symoccurs_1check, tmp17341)

tmp17343 := PrimCons(MakeNumber(1), tmp17342)

tmp17344 := PrimCons(symnumber_2, tmp17343)

tmp17345 := PrimCons(MakeNumber(1), tmp17344)

tmp17346 := PrimCons(symn_1_6string, tmp17345)

tmp17347 := PrimCons(MakeNumber(2), tmp17346)

tmp17348 := PrimCons(symnth, tmp17347)

tmp17349 := PrimCons(MakeNumber(1), tmp17348)

tmp17350 := PrimCons(symnot, tmp17349)

tmp17351 := PrimCons(MakeNumber(1), tmp17350)

tmp17352 := PrimCons(symnl, tmp17351)

tmp17353 := PrimCons(MakeNumber(1), tmp17352)

tmp17354 := PrimCons(symmaxinferences, tmp17353)

tmp17355 := PrimCons(MakeNumber(2), tmp17354)

tmp17356 := PrimCons(symmapcan, tmp17355)

tmp17357 := PrimCons(MakeNumber(2), tmp17356)

tmp17358 := PrimCons(symmap, tmp17357)

tmp17359 := PrimCons(MakeNumber(1), tmp17358)

tmp17360 := PrimCons(symmacroexpand, tmp17359)

tmp17361 := PrimCons(MakeNumber(1), tmp17360)

tmp17362 := PrimCons(symvector, tmp17361)

tmp17363 := PrimCons(MakeNumber(2), tmp17362)

tmp17364 := PrimCons(sym_5_a, tmp17363)

tmp17365 := PrimCons(MakeNumber(2), tmp17364)

tmp17366 := PrimCons(sym_5, tmp17365)

tmp17367 := PrimCons(MakeNumber(1), tmp17366)

tmp17368 := PrimCons(symload, tmp17367)

tmp17369 := PrimCons(MakeNumber(1), tmp17368)

tmp17370 := PrimCons(symlist, tmp17369)

tmp17371 := PrimCons(MakeNumber(1), tmp17370)

tmp17372 := PrimCons(symlineread, tmp17371)

tmp17373 := PrimCons(MakeNumber(1), tmp17372)

tmp17374 := PrimCons(symlimit, tmp17373)

tmp17375 := PrimCons(MakeNumber(1), tmp17374)

tmp17376 := PrimCons(symlength, tmp17375)

tmp17377 := PrimCons(MakeNumber(0), tmp17376)

tmp17378 := PrimCons(symlanguage, tmp17377)

tmp17379 := PrimCons(MakeNumber(6), tmp17378)

tmp17380 := PrimCons(symis_b, tmp17379)

tmp17381 := PrimCons(MakeNumber(6), tmp17380)

tmp17382 := PrimCons(symis, tmp17381)

tmp17383 := PrimCons(MakeNumber(0), tmp17382)

tmp17384 := PrimCons(symit, tmp17383)

tmp17385 := PrimCons(MakeNumber(1), tmp17384)

tmp17386 := PrimCons(syminternal, tmp17385)

tmp17387 := PrimCons(MakeNumber(2), tmp17386)

tmp17388 := PrimCons(symintersection, tmp17387)

tmp17389 := PrimCons(MakeNumber(1), tmp17388)

tmp17390 := PrimCons(syminclude_1all_1but, tmp17389)

tmp17391 := PrimCons(MakeNumber(0), tmp17390)

tmp17392 := PrimCons(symimplementation, tmp17391)

tmp17393 := PrimCons(MakeNumber(2), tmp17392)

tmp17394 := PrimCons(syminput_7, tmp17393)

tmp17395 := PrimCons(MakeNumber(1), tmp17394)

tmp17396 := PrimCons(syminput, tmp17395)

tmp17397 := PrimCons(MakeNumber(0), tmp17396)

tmp17398 := PrimCons(syminferences, tmp17397)

tmp17399 := PrimCons(MakeNumber(1), tmp17398)

tmp17400 := PrimCons(symintern, tmp17399)

tmp17401 := PrimCons(MakeNumber(1), tmp17400)

tmp17402 := PrimCons(syminternal, tmp17401)

tmp17403 := PrimCons(MakeNumber(1), tmp17402)

tmp17404 := PrimCons(syminteger_2, tmp17403)

tmp17405 := PrimCons(MakeNumber(1), tmp17404)

tmp17406 := PrimCons(symin_1package, tmp17405)

tmp17407 := PrimCons(MakeNumber(0), tmp17406)

tmp17408 := PrimCons(symshen_4included, tmp17407)

tmp17409 := PrimCons(MakeNumber(1), tmp17408)

tmp17410 := PrimCons(syminclude, tmp17409)

tmp17411 := PrimCons(MakeNumber(3), tmp17410)

tmp17412 := PrimCons(symif, tmp17411)

tmp17413 := PrimCons(MakeNumber(1), tmp17412)

tmp17414 := PrimCons(symhush, tmp17413)

tmp17415 := PrimCons(MakeNumber(0), tmp17414)

tmp17416 := PrimCons(symhush_2, tmp17415)

tmp17417 := PrimCons(MakeNumber(1), tmp17416)

tmp17418 := PrimCons(symhead, tmp17417)

tmp17419 := PrimCons(MakeNumber(1), tmp17418)

tmp17420 := PrimCons(symhdstr, tmp17419)

tmp17421 := PrimCons(MakeNumber(1), tmp17420)

tmp17422 := PrimCons(symhdv, tmp17421)

tmp17423 := PrimCons(MakeNumber(1), tmp17422)

tmp17424 := PrimCons(symhd, tmp17423)

tmp17425 := PrimCons(MakeNumber(2), tmp17424)

tmp17426 := PrimCons(symhash, tmp17425)

tmp17427 := PrimCons(MakeNumber(2), tmp17426)

tmp17428 := PrimCons(sym_a, tmp17427)

tmp17429 := PrimCons(MakeNumber(2), tmp17428)

tmp17430 := PrimCons(sym_6_a, tmp17429)

tmp17431 := PrimCons(MakeNumber(2), tmp17430)

tmp17432 := PrimCons(sym_6, tmp17431)

tmp17433 := PrimCons(MakeNumber(2), tmp17432)

tmp17434 := PrimCons(sym_5_1vector, tmp17433)

tmp17435 := PrimCons(MakeNumber(2), tmp17434)

tmp17436 := PrimCons(sym_5_1address, tmp17435)

tmp17437 := PrimCons(MakeNumber(3), tmp17436)

tmp17438 := PrimCons(symaddress_1_6, tmp17437)

tmp17439 := PrimCons(MakeNumber(1), tmp17438)

tmp17440 := PrimCons(symget_1time, tmp17439)

tmp17441 := PrimCons(MakeNumber(3), tmp17440)

tmp17442 := PrimCons(symget, tmp17441)

tmp17443 := PrimCons(MakeNumber(1), tmp17442)

tmp17444 := PrimCons(symgensym, tmp17443)

tmp17445 := PrimCons(MakeNumber(1), tmp17444)

tmp17446 := PrimCons(symfunction, tmp17445)

tmp17447 := PrimCons(MakeNumber(1), tmp17446)

tmp17448 := PrimCons(symfn, tmp17447)

tmp17449 := PrimCons(MakeNumber(1), tmp17448)

tmp17450 := PrimCons(symfst, tmp17449)

tmp17451 := PrimCons(MakeNumber(0), tmp17450)

tmp17452 := PrimCons(symfresh, tmp17451)

tmp17453 := PrimCons(MakeNumber(1), tmp17452)

tmp17454 := PrimCons(symfreeze, tmp17453)

tmp17455 := PrimCons(MakeNumber(5), tmp17454)

tmp17456 := PrimCons(symfork, tmp17455)

tmp17457 := PrimCons(MakeNumber(1), tmp17456)

tmp17458 := PrimCons(symforeign, tmp17457)

tmp17459 := PrimCons(MakeNumber(7), tmp17458)

tmp17460 := PrimCons(symfindall, tmp17459)

tmp17461 := PrimCons(MakeNumber(2), tmp17460)

tmp17462 := PrimCons(symfix, tmp17461)

tmp17463 := PrimCons(MakeNumber(0), tmp17462)

tmp17464 := PrimCons(symfail, tmp17463)

tmp17465 := PrimCons(MakeNumber(2), tmp17464)

tmp17466 := PrimCons(symfail_1if, tmp17465)

tmp17467 := PrimCons(MakeNumber(0), tmp17466)

tmp17468 := PrimCons(symfactorise_2, tmp17467)

tmp17469 := PrimCons(MakeNumber(1), tmp17468)

tmp17470 := PrimCons(symfactorise, tmp17469)

tmp17471 := PrimCons(MakeNumber(1), tmp17470)

tmp17472 := PrimCons(symexternal, tmp17471)

tmp17473 := PrimCons(MakeNumber(1), tmp17472)

tmp17474 := PrimCons(symexplode, tmp17473)

tmp17475 := PrimCons(MakeNumber(1), tmp17474)

tmp17476 := PrimCons(symeval_1kl, tmp17475)

tmp17477 := PrimCons(MakeNumber(1), tmp17476)

tmp17478 := PrimCons(symeval, tmp17477)

tmp17479 := PrimCons(MakeNumber(1), tmp17478)

tmp17480 := PrimCons(symerror_1to_1string, tmp17479)

tmp17481 := PrimCons(MakeNumber(1), tmp17480)

tmp17482 := PrimCons(symexternal, tmp17481)

tmp17483 := PrimCons(MakeNumber(1), tmp17482)

tmp17484 := PrimCons(symenable_1type_1theory, tmp17483)

tmp17485 := PrimCons(MakeNumber(1), tmp17484)

tmp17486 := PrimCons(symempty_2, tmp17485)

tmp17487 := PrimCons(MakeNumber(2), tmp17486)

tmp17488 := PrimCons(symelement_2, tmp17487)

tmp17489 := PrimCons(MakeNumber(2), tmp17488)

tmp17490 := PrimCons(symdo, tmp17489)

tmp17491 := PrimCons(MakeNumber(2), tmp17490)

tmp17492 := PrimCons(symdifference, tmp17491)

tmp17493 := PrimCons(MakeNumber(1), tmp17492)

tmp17494 := PrimCons(symdestroy, tmp17493)

tmp17495 := PrimCons(MakeNumber(2), tmp17494)

tmp17496 := PrimCons(symdeclare, tmp17495)

tmp17497 := PrimCons(MakeNumber(0), tmp17496)

tmp17498 := PrimCons(symdatatypes, tmp17497)

tmp17499 := PrimCons(MakeNumber(1), tmp17498)

tmp17500 := PrimCons(symclose, tmp17499)

tmp17501 := PrimCons(MakeNumber(2), tmp17500)

tmp17502 := PrimCons(symcn, tmp17501)

tmp17503 := PrimCons(MakeNumber(1), tmp17502)

tmp17504 := PrimCons(symcons_2, tmp17503)

tmp17505 := PrimCons(MakeNumber(2), tmp17504)

tmp17506 := PrimCons(symcons, tmp17505)

tmp17507 := PrimCons(MakeNumber(2), tmp17506)

tmp17508 := PrimCons(symconcat, tmp17507)

tmp17509 := PrimCons(MakeNumber(2), tmp17508)

tmp17510 := PrimCons(symcompile, tmp17509)

tmp17511 := PrimCons(MakeNumber(1), tmp17510)

tmp17512 := PrimCons(symcd, tmp17511)

tmp17513 := PrimCons(MakeNumber(5), tmp17512)

tmp17514 := PrimCons(symcall, tmp17513)

tmp17515 := PrimCons(MakeNumber(6), tmp17514)

tmp17516 := PrimCons(symbind, tmp17515)

tmp17517 := PrimCons(MakeNumber(1), tmp17516)

tmp17518 := PrimCons(symbound_2, tmp17517)

tmp17519 := PrimCons(MakeNumber(1), tmp17518)

tmp17520 := PrimCons(symbootstrap, tmp17519)

tmp17521 := PrimCons(MakeNumber(1), tmp17520)

tmp17522 := PrimCons(symboolean_2, tmp17521)

tmp17523 := PrimCons(MakeNumber(1), tmp17522)

tmp17524 := PrimCons(symatom_2, tmp17523)

tmp17525 := PrimCons(MakeNumber(2), tmp17524)

tmp17526 := PrimCons(symassoc, tmp17525)

tmp17527 := PrimCons(MakeNumber(1), tmp17526)

tmp17528 := PrimCons(symarity, tmp17527)

tmp17529 := PrimCons(MakeNumber(2), tmp17528)

tmp17530 := PrimCons(symappend, tmp17529)

tmp17531 := PrimCons(MakeNumber(2), tmp17530)

tmp17532 := PrimCons(symand, tmp17531)

tmp17533 := PrimCons(MakeNumber(2), tmp17532)

tmp17534 := PrimCons(symadjoin, tmp17533)

tmp17535 := PrimCons(MakeNumber(3), tmp17534)

tmp17536 := PrimCons(symaddress_1_6, tmp17535)

tmp17537 := PrimCons(MakeNumber(1), tmp17536)

tmp17538 := PrimCons(symabsvector, tmp17537)

tmp17539 := PrimCons(MakeNumber(1), tmp17538)

tmp17540 := PrimCons(symabsvector_2, tmp17539)

tmp17541 := PrimCons(MakeNumber(1), tmp17540)

tmp17542 := PrimCons(symabsolute, tmp17541)

tmp17543 := PrimCons(MakeNumber(0), tmp17542)

tmp17544 := PrimCons(symabort, tmp17543)

tmp17545 := Call(__e, PrimFunc(symshen_4initialise_1arity_1table), tmp17544)


_ = tmp17545

tmp17546 := PrimIntern(MakeString(":"))

tmp17547 := PrimIntern(MakeString(";"))

tmp17548 := PrimIntern(MakeString(":="))

tmp17549 := PrimIntern(MakeString(","))

tmp17550 := Call(__e, PrimFunc(symvector), MakeNumber(0))


tmp17551 := PrimIntern(MakeString("bar!"))

tmp17552 := PrimCons(symabort, Nil)

tmp17553 := PrimCons(symabsolute, tmp17552)

tmp17554 := PrimCons(symabsvector, tmp17553)

tmp17555 := PrimCons(symabsvector_2, tmp17554)

tmp17556 := PrimCons(symaddress_1_6, tmp17555)

tmp17557 := PrimCons(sym_5_1address, tmp17556)

tmp17558 := PrimCons(symadjoin, tmp17557)

tmp17559 := PrimCons(symand, tmp17558)

tmp17560 := PrimCons(symappend, tmp17559)

tmp17561 := PrimCons(symarity, tmp17560)

tmp17562 := PrimCons(symassoc, tmp17561)

tmp17563 := PrimCons(symassertz, tmp17562)

tmp17564 := PrimCons(symasserta, tmp17563)

tmp17565 := PrimCons(symatom_2, tmp17564)

tmp17566 := PrimCons(tmp17551, tmp17565)

tmp17567 := PrimCons(symbootstrap, tmp17566)

tmp17568 := PrimCons(symboolean, tmp17567)

tmp17569 := PrimCons(symboolean_2, tmp17568)

tmp17570 := PrimCons(symbound_2, tmp17569)

tmp17571 := PrimCons(symbind, tmp17570)

tmp17572 := PrimCons(symclose, tmp17571)

tmp17573 := PrimCons(symcall, tmp17572)

tmp17574 := PrimCons(symcases, tmp17573)

tmp17575 := PrimCons(symcd, tmp17574)

tmp17576 := PrimCons(symcompile, tmp17575)

tmp17577 := PrimCons(symconcat, tmp17576)

tmp17578 := PrimCons(symcond, tmp17577)

tmp17579 := PrimCons(symcons, tmp17578)

tmp17580 := PrimCons(symcons_2, tmp17579)

tmp17581 := PrimCons(symcn, tmp17580)

tmp17582 := PrimCons(symshen_4ctxt, tmp17581)

tmp17583 := PrimCons(symdatatypes, tmp17582)

tmp17584 := PrimCons(symdatatype, tmp17583)

tmp17585 := PrimCons(symdeclare, tmp17584)

tmp17586 := PrimCons(symdefprolog, tmp17585)

tmp17587 := PrimCons(symdefcc, tmp17586)

tmp17588 := PrimCons(symdefmacro, tmp17587)

tmp17589 := PrimCons(symdefine, tmp17588)

tmp17590 := PrimCons(symdefun, tmp17589)

tmp17591 := PrimCons(symdestroy, tmp17590)

tmp17592 := PrimCons(symdifference, tmp17591)

tmp17593 := PrimCons(symdo, tmp17592)

tmp17594 := PrimCons(symelement_2, tmp17593)

tmp17595 := PrimCons(symempty_2, tmp17594)

tmp17596 := PrimCons(symerror, tmp17595)

tmp17597 := PrimCons(symerror_1to_1string, tmp17596)

tmp17598 := PrimCons(symeval, tmp17597)

tmp17599 := PrimCons(symeval_1kl, tmp17598)

tmp17600 := PrimCons(symexception, tmp17599)

tmp17601 := PrimCons(symexternal, tmp17600)

tmp17602 := PrimCons(symexplode, tmp17601)

tmp17603 := PrimCons(symenable_1type_1theory, tmp17602)

tmp17604 := PrimCons(False, tmp17603)

tmp17605 := PrimCons(symfindall, tmp17604)

tmp17606 := PrimCons(symfactorise, tmp17605)

tmp17607 := PrimCons(symfail_1if, tmp17606)

tmp17608 := PrimCons(symfail, tmp17607)

tmp17609 := PrimCons(symfile, tmp17608)

tmp17610 := PrimCons(symfix, tmp17609)

tmp17611 := PrimCons(symforeign, tmp17610)

tmp17612 := PrimCons(symfork, tmp17611)

tmp17613 := PrimCons(symfresh, tmp17612)

tmp17614 := PrimCons(symfreeze, tmp17613)

tmp17615 := PrimCons(symfst, tmp17614)

tmp17616 := PrimCons(symfunction, tmp17615)

tmp17617 := PrimCons(symfn, tmp17616)

tmp17618 := PrimCons(symgensym, tmp17617)

tmp17619 := PrimCons(symget_1time, tmp17618)

tmp17620 := PrimCons(symget, tmp17619)

tmp17621 := PrimCons(symhash, tmp17620)

tmp17622 := PrimCons(symhdstr, tmp17621)

tmp17623 := PrimCons(symhdv, tmp17622)

tmp17624 := PrimCons(symhd, tmp17623)

tmp17625 := PrimCons(symhead, tmp17624)

tmp17626 := PrimCons(symif, tmp17625)

tmp17627 := PrimCons(symimplementation, tmp17626)

tmp17628 := PrimCons(syminternal, tmp17627)

tmp17629 := PrimCons(symin_1package, tmp17628)

tmp17630 := PrimCons(symin, tmp17629)

tmp17631 := PrimCons(symis_b, tmp17630)

tmp17632 := PrimCons(symis, tmp17631)

tmp17633 := PrimCons(symit, tmp17632)

tmp17634 := PrimCons(syminclude_1all_1but, tmp17633)

tmp17635 := PrimCons(syminclude, tmp17634)

tmp17636 := PrimCons(syminline, tmp17635)

tmp17637 := PrimCons(syminput_7, tmp17636)

tmp17638 := PrimCons(syminput, tmp17637)

tmp17639 := PrimCons(syminteger_2, tmp17638)

tmp17640 := PrimCons(symintern, tmp17639)

tmp17641 := PrimCons(syminferences, tmp17640)

tmp17642 := PrimCons(symintersection, tmp17641)

tmp17643 := PrimCons(symis, tmp17642)

tmp17644 := PrimCons(symlanguage, tmp17643)

tmp17645 := PrimCons(symlambda, tmp17644)

tmp17646 := PrimCons(symlazy, tmp17645)

tmp17647 := PrimCons(symlet, tmp17646)

tmp17648 := PrimCons(symlength, tmp17647)

tmp17649 := PrimCons(symlimit, tmp17648)

tmp17650 := PrimCons(symlineread, tmp17649)

tmp17651 := PrimCons(symlist, tmp17650)

tmp17652 := PrimCons(symloaded, tmp17651)

tmp17653 := PrimCons(symload, tmp17652)

tmp17654 := PrimCons(symmake_1string, tmp17653)

tmp17655 := PrimCons(symmap, tmp17654)

tmp17656 := PrimCons(symmapcan, tmp17655)

tmp17657 := PrimCons(symmaxinferences, tmp17656)

tmp17658 := PrimCons(symmacroexpand, tmp17657)

tmp17659 := PrimCons(symmode, tmp17658)

tmp17660 := PrimCons(symnl, tmp17659)

tmp17661 := PrimCons(symnot, tmp17660)

tmp17662 := PrimCons(symnth, tmp17661)

tmp17663 := PrimCons(symnull, tmp17662)

tmp17664 := PrimCons(symnumber, tmp17663)

tmp17665 := PrimCons(symnumber_2, tmp17664)

tmp17666 := PrimCons(symn_1_6string, tmp17665)

tmp17667 := PrimCons(symoccurs_1check, tmp17666)

tmp17668 := PrimCons(symoccurrences, tmp17667)

tmp17669 := PrimCons(symopen, tmp17668)

tmp17670 := PrimCons(symoptimise, tmp17669)

tmp17671 := PrimCons(symor, tmp17670)

tmp17672 := PrimCons(symos, tmp17671)

tmp17673 := PrimCons(symout, tmp17672)

tmp17674 := PrimCons(symoutput, tmp17673)

tmp17675 := PrimCons(sympackage, tmp17674)

tmp17676 := PrimCons(symport, tmp17675)

tmp17677 := PrimCons(symporters, tmp17676)

tmp17678 := PrimCons(sympos, tmp17677)

tmp17679 := PrimCons(sympr, tmp17678)

tmp17680 := PrimCons(symprint, tmp17679)

tmp17681 := PrimCons(symprolog_1memory, tmp17680)

tmp17682 := PrimCons(symprofile, tmp17681)

tmp17683 := PrimCons(symprofile_1results, tmp17682)

tmp17684 := PrimCons(symprotect, tmp17683)

tmp17685 := PrimCons(symprolog_2, tmp17684)

tmp17686 := PrimCons(symps, tmp17685)

tmp17687 := PrimCons(sympreclude_1all_1but, tmp17686)

tmp17688 := PrimCons(sympreclude, tmp17687)

tmp17689 := PrimCons(symput, tmp17688)

tmp17690 := PrimCons(sympackage_2, tmp17689)

tmp17691 := PrimCons(symread_1from_1string_1unprocessed, tmp17690)

tmp17692 := PrimCons(symread_1from_1string, tmp17691)

tmp17693 := PrimCons(symread_1byte, tmp17692)

tmp17694 := PrimCons(symread_1file_1as_1string, tmp17693)

tmp17695 := PrimCons(symread_1file_1as_1bytelist, tmp17694)

tmp17696 := PrimCons(symread_1file, tmp17695)

tmp17697 := PrimCons(symreceive, tmp17696)

tmp17698 := PrimCons(symread, tmp17697)

tmp17699 := PrimCons(symrelease, tmp17698)

tmp17700 := PrimCons(symremove, tmp17699)

tmp17701 := PrimCons(symretract, tmp17700)

tmp17702 := PrimCons(symreverse, tmp17701)

tmp17703 := PrimCons(symrun, tmp17702)

tmp17704 := PrimCons(symstr, tmp17703)

tmp17705 := PrimCons(symsave, tmp17704)

tmp17706 := PrimCons(symset, tmp17705)

tmp17707 := PrimCons(symsimple_1error, tmp17706)

tmp17708 := PrimCons(symsnd, tmp17707)

tmp17709 := PrimCons(symspecialise, tmp17708)

tmp17710 := PrimCons(symspy, tmp17709)

tmp17711 := PrimCons(symstep, tmp17710)

tmp17712 := PrimCons(symstoutput, tmp17711)

tmp17713 := PrimCons(symsterror, tmp17712)

tmp17714 := PrimCons(symstinput, tmp17713)

tmp17715 := PrimCons(symstring, tmp17714)

tmp17716 := PrimCons(symstream, tmp17715)

tmp17717 := PrimCons(symstring_1_6n, tmp17716)

tmp17718 := PrimCons(symstring_2, tmp17717)

tmp17719 := PrimCons(symsubst, tmp17718)

tmp17720 := PrimCons(symsum, tmp17719)

tmp17721 := PrimCons(symstring_1_6symbol, tmp17720)

tmp17722 := PrimCons(symsymbol_2, tmp17721)

tmp17723 := PrimCons(symsymbol, tmp17722)

tmp17724 := PrimCons(symsynonyms, tmp17723)

tmp17725 := PrimCons(symsystemf, tmp17724)

tmp17726 := PrimCons(symtail, tmp17725)

tmp17727 := PrimCons(symtlv, tmp17726)

tmp17728 := PrimCons(symtlstr, tmp17727)

tmp17729 := PrimCons(symtl, tmp17728)

tmp17730 := PrimCons(symtc, tmp17729)

tmp17731 := PrimCons(symtc_2, tmp17730)

tmp17732 := PrimCons(symthaw, tmp17731)

tmp17733 := PrimCons(symtime, tmp17732)

tmp17734 := PrimCons(symtrack, tmp17733)

tmp17735 := PrimCons(symtrap_1error, tmp17734)

tmp17736 := PrimCons(True, tmp17735)

tmp17737 := PrimCons(symtuple_2, tmp17736)

tmp17738 := PrimCons(symtype, tmp17737)

tmp17739 := PrimCons(symreturn, tmp17738)

tmp17740 := PrimCons(symunabsolute, tmp17739)

tmp17741 := PrimCons(symundefmacro, tmp17740)

tmp17742 := PrimCons(symunprofile, tmp17741)

tmp17743 := PrimCons(symunput, tmp17742)

tmp17744 := PrimCons(symunion, tmp17743)

tmp17745 := PrimCons(symshen_4unix, tmp17744)

tmp17746 := PrimCons(symunit, tmp17745)

tmp17747 := PrimCons(symuntrack, tmp17746)

tmp17748 := PrimCons(symunspecialise, tmp17747)

tmp17749 := PrimCons(symupdate_1lambda_1table, tmp17748)

tmp17750 := PrimCons(symu_b, tmp17749)

tmp17751 := PrimCons(symvector_2, tmp17750)

tmp17752 := PrimCons(symvector, tmp17751)

tmp17753 := PrimCons(sym_5_1vector, tmp17752)

tmp17754 := PrimCons(symvector_1_6, tmp17753)

tmp17755 := PrimCons(symvalue, tmp17754)

tmp17756 := PrimCons(symvar_2, tmp17755)

tmp17757 := PrimCons(symvariable_2, tmp17756)

tmp17758 := PrimCons(symverified, tmp17757)

tmp17759 := PrimCons(symversion, tmp17758)

tmp17760 := PrimCons(symwarn, tmp17759)

tmp17761 := PrimCons(symwhen, tmp17760)

tmp17762 := PrimCons(symwhere, tmp17761)

tmp17763 := PrimCons(symwrite_1byte, tmp17762)

tmp17764 := PrimCons(symwrite_1to_1file, tmp17763)

tmp17765 := PrimCons(symy_1or_1n_2, tmp17764)

tmp17766 := PrimCons(tmp17550, tmp17765)

tmp17767 := PrimCons(sym_6_6, tmp17766)

tmp17768 := PrimCons(sym_5, tmp17767)

tmp17769 := PrimCons(sym_5_a, tmp17768)

tmp17770 := PrimCons(sym_7, tmp17769)

tmp17771 := PrimCons(sym_d, tmp17770)

tmp17772 := PrimCons(sym_c, tmp17771)

tmp17773 := PrimCons(sym_1, tmp17772)

tmp17774 := PrimCons(sym_3, tmp17773)

tmp17775 := PrimCons(sym_5end_6, tmp17774)

tmp17776 := PrimCons(sym_5_b_6, tmp17775)

tmp17777 := PrimCons(sym_c_4, tmp17776)

tmp17778 := PrimCons(sym_a_a_6, tmp17777)

tmp17779 := PrimCons(sym_6, tmp17778)

tmp17780 := PrimCons(sym_6_a, tmp17779)

tmp17781 := PrimCons(sym_a, tmp17780)

tmp17782 := PrimCons(sym_a_a, tmp17781)

tmp17783 := PrimCons(sym_5e_6, tmp17782)

tmp17784 := PrimCons(sym_1_6, tmp17783)

tmp17785 := PrimCons(sym_5_1, tmp17784)

tmp17786 := PrimCons(sym_dhush_d, tmp17785)

tmp17787 := PrimCons(sym_dporters_d, tmp17786)

tmp17788 := PrimCons(sym_dport_d, tmp17787)

tmp17789 := PrimCons(sym_8s, tmp17788)

tmp17790 := PrimCons(sym_8p, tmp17789)

tmp17791 := PrimCons(sym_8v, tmp17790)

tmp17792 := PrimCons(sym_dproperty_1vector_d, tmp17791)

tmp17793 := PrimCons(sym_drelease_d, tmp17792)

tmp17794 := PrimCons(sym_dos_d, tmp17793)

tmp17795 := PrimCons(sym_dmacros_d, tmp17794)

tmp17796 := PrimCons(sym_dmaximum_1print_1sequence_1size_d, tmp17795)

tmp17797 := PrimCons(sym_dversion_d, tmp17796)

tmp17798 := PrimCons(sym_dhome_1directory_d, tmp17797)

tmp17799 := PrimCons(sym_dstoutput_d, tmp17798)

tmp17800 := PrimCons(sym_dsterror_d, tmp17799)

tmp17801 := PrimCons(sym_dstinput_d, tmp17800)

tmp17802 := PrimCons(sym_dimplementation_d, tmp17801)

tmp17803 := PrimCons(sym_dlanguage_d, tmp17802)

tmp17804 := PrimCons(sym__, tmp17803)

tmp17805 := PrimCons(tmp17549, tmp17804)

tmp17806 := PrimCons(tmp17548, tmp17805)

tmp17807 := PrimCons(tmp17547, tmp17806)

tmp17808 := PrimCons(tmp17546, tmp17807)

tmp17809 := PrimCons(sym_e_e, tmp17808)

tmp17810 := PrimCons(sym_5_1_1, tmp17809)

tmp17811 := PrimCons(sym_1_1_6, tmp17810)

tmp17812 := PrimCons(sym_i, tmp17811)

tmp17813 := PrimCons(sym_j, tmp17812)

tmp17814 := PrimCons(sym_b, tmp17813)

tmp17815 := PrimValue(sym_dproperty_1vector_d)

tmp17816 := Call(__e, PrimFunc(symput), symshen, symshen_4external_1symbols, tmp17814, tmp17815)


_ = tmp17816

tmp17817 := PrimAbsvector(MakeNumber(0))

__e.Return(PrimSet(symshen_4_dempty_1absvector_d, tmp17817))
return


}, 0)

tmp17818 := Call(__e, ns2_1set, symshen_4initialise_1environment, tmp17077)


_ = tmp17818

tmp17819 := MakeNative(func(__e *ControlFlow) {
tmp17820 := PrimSet(symshen_4_dsigf_d, Nil)

_ = tmp17820

tmp17821 := MakeNative(func(__e *ControlFlow) {
V5951 := __e.Get(1)
_ = V5951
__e.Return(MakeNative(func(__e *ControlFlow) {
B5947 := __e.Get(1)
_ = B5947
__e.Return(MakeNative(func(__e *ControlFlow) {
L5948 := __e.Get(1)
_ = L5948
__e.Return(MakeNative(func(__e *ControlFlow) {
Key5949 := __e.Get(1)
_ = Key5949
__e.Return(MakeNative(func(__e *ControlFlow) {
C5950 := __e.Get(1)
_ = C5950
tmp17822 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp17823 := PrimCons(A, Nil)

tmp17824 := PrimCons(sym_1_1_6, tmp17823)

tmp17825 := Call(__e, PrimFunc(symis_b), V5951, tmp17824, B5947, L5948, Key5949, C5950)


__e.TailApply(PrimFunc(symshen_4gc), B5947, tmp17825)
return


}, 1)

tmp17826 := Call(__e, PrimFunc(symshen_4newpv), B5947)


__e.TailApply(tmp17822, tmp17826)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp17827 := PrimValue(symshen_4_dsigf_d)

tmp17828 := Call(__e, PrimFunc(symshen_4assoc_1_6), symabort, tmp17821, tmp17827)


tmp17829 := PrimSet(symshen_4_dsigf_d, tmp17828)

_ = tmp17829

tmp17830 := MakeNative(func(__e *ControlFlow) {
V5956 := __e.Get(1)
_ = V5956
__e.Return(MakeNative(func(__e *ControlFlow) {
B5952 := __e.Get(1)
_ = B5952
__e.Return(MakeNative(func(__e *ControlFlow) {
L5953 := __e.Get(1)
_ = L5953
__e.Return(MakeNative(func(__e *ControlFlow) {
Key5954 := __e.Get(1)
_ = Key5954
__e.Return(MakeNative(func(__e *ControlFlow) {
C5955 := __e.Get(1)
_ = C5955
tmp17831 := PrimCons(symstring, Nil)

tmp17832 := PrimCons(symlist, tmp17831)

tmp17833 := PrimCons(tmp17832, Nil)

tmp17834 := PrimCons(sym_1_1_6, tmp17833)

tmp17835 := PrimCons(symstring, tmp17834)

__e.TailApply(PrimFunc(symis_b), V5956, tmp17835, B5952, L5953, Key5954, C5955)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp17836 := PrimValue(symshen_4_dsigf_d)

tmp17837 := Call(__e, PrimFunc(symshen_4assoc_1_6), symabsolute, tmp17830, tmp17836)


tmp17838 := PrimSet(symshen_4_dsigf_d, tmp17837)

_ = tmp17838

tmp17839 := MakeNative(func(__e *ControlFlow) {
V5961 := __e.Get(1)
_ = V5961
__e.Return(MakeNative(func(__e *ControlFlow) {
B5957 := __e.Get(1)
_ = B5957
__e.Return(MakeNative(func(__e *ControlFlow) {
L5958 := __e.Get(1)
_ = L5958
__e.Return(MakeNative(func(__e *ControlFlow) {
Key5959 := __e.Get(1)
_ = Key5959
__e.Return(MakeNative(func(__e *ControlFlow) {
C5960 := __e.Get(1)
_ = C5960
tmp17840 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp17841 := PrimCons(symboolean, Nil)

tmp17842 := PrimCons(sym_1_1_6, tmp17841)

tmp17843 := PrimCons(A, tmp17842)

tmp17844 := Call(__e, PrimFunc(symis_b), V5961, tmp17843, B5957, L5958, Key5959, C5960)


__e.TailApply(PrimFunc(symshen_4gc), B5957, tmp17844)
return


}, 1)

tmp17845 := Call(__e, PrimFunc(symshen_4newpv), B5957)


__e.TailApply(tmp17840, tmp17845)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp17846 := PrimValue(symshen_4_dsigf_d)

tmp17847 := Call(__e, PrimFunc(symshen_4assoc_1_6), symabsvector_2, tmp17839, tmp17846)


tmp17848 := PrimSet(symshen_4_dsigf_d, tmp17847)

_ = tmp17848

tmp17849 := MakeNative(func(__e *ControlFlow) {
V5966 := __e.Get(1)
_ = V5966
__e.Return(MakeNative(func(__e *ControlFlow) {
B5962 := __e.Get(1)
_ = B5962
__e.Return(MakeNative(func(__e *ControlFlow) {
L5963 := __e.Get(1)
_ = L5963
__e.Return(MakeNative(func(__e *ControlFlow) {
Key5964 := __e.Get(1)
_ = Key5964
__e.Return(MakeNative(func(__e *ControlFlow) {
C5965 := __e.Get(1)
_ = C5965
tmp17850 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp17851 := PrimCons(A, Nil)

tmp17852 := PrimCons(symlist, tmp17851)

tmp17853 := PrimCons(A, Nil)

tmp17854 := PrimCons(symlist, tmp17853)

tmp17855 := PrimCons(tmp17854, Nil)

tmp17856 := PrimCons(sym_1_1_6, tmp17855)

tmp17857 := PrimCons(tmp17852, tmp17856)

tmp17858 := PrimCons(tmp17857, Nil)

tmp17859 := PrimCons(sym_1_1_6, tmp17858)

tmp17860 := PrimCons(A, tmp17859)

tmp17861 := Call(__e, PrimFunc(symis_b), V5966, tmp17860, B5962, L5963, Key5964, C5965)


__e.TailApply(PrimFunc(symshen_4gc), B5962, tmp17861)
return


}, 1)

tmp17862 := Call(__e, PrimFunc(symshen_4newpv), B5962)


__e.TailApply(tmp17850, tmp17862)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp17863 := PrimValue(symshen_4_dsigf_d)

tmp17864 := Call(__e, PrimFunc(symshen_4assoc_1_6), symadjoin, tmp17849, tmp17863)


tmp17865 := PrimSet(symshen_4_dsigf_d, tmp17864)

_ = tmp17865

tmp17866 := MakeNative(func(__e *ControlFlow) {
V5971 := __e.Get(1)
_ = V5971
__e.Return(MakeNative(func(__e *ControlFlow) {
B5967 := __e.Get(1)
_ = B5967
__e.Return(MakeNative(func(__e *ControlFlow) {
L5968 := __e.Get(1)
_ = L5968
__e.Return(MakeNative(func(__e *ControlFlow) {
Key5969 := __e.Get(1)
_ = Key5969
__e.Return(MakeNative(func(__e *ControlFlow) {
C5970 := __e.Get(1)
_ = C5970
tmp17867 := PrimCons(symboolean, Nil)

tmp17868 := PrimCons(sym_1_1_6, tmp17867)

tmp17869 := PrimCons(symboolean, tmp17868)

tmp17870 := PrimCons(tmp17869, Nil)

tmp17871 := PrimCons(sym_1_1_6, tmp17870)

tmp17872 := PrimCons(symboolean, tmp17871)

__e.TailApply(PrimFunc(symis_b), V5971, tmp17872, B5967, L5968, Key5969, C5970)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp17873 := PrimValue(symshen_4_dsigf_d)

tmp17874 := Call(__e, PrimFunc(symshen_4assoc_1_6), symand, tmp17866, tmp17873)


tmp17875 := PrimSet(symshen_4_dsigf_d, tmp17874)

_ = tmp17875

tmp17876 := MakeNative(func(__e *ControlFlow) {
V5976 := __e.Get(1)
_ = V5976
__e.Return(MakeNative(func(__e *ControlFlow) {
B5972 := __e.Get(1)
_ = B5972
__e.Return(MakeNative(func(__e *ControlFlow) {
L5973 := __e.Get(1)
_ = L5973
__e.Return(MakeNative(func(__e *ControlFlow) {
Key5974 := __e.Get(1)
_ = Key5974
__e.Return(MakeNative(func(__e *ControlFlow) {
C5975 := __e.Get(1)
_ = C5975
tmp17877 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp17878 := PrimCons(symstring, Nil)

tmp17879 := PrimCons(sym_1_1_6, tmp17878)

tmp17880 := PrimCons(symsymbol, tmp17879)

tmp17881 := PrimCons(tmp17880, Nil)

tmp17882 := PrimCons(sym_1_1_6, tmp17881)

tmp17883 := PrimCons(symstring, tmp17882)

tmp17884 := PrimCons(tmp17883, Nil)

tmp17885 := PrimCons(sym_1_1_6, tmp17884)

tmp17886 := PrimCons(A, tmp17885)

tmp17887 := Call(__e, PrimFunc(symis_b), V5976, tmp17886, B5972, L5973, Key5974, C5975)


__e.TailApply(PrimFunc(symshen_4gc), B5972, tmp17887)
return


}, 1)

tmp17888 := Call(__e, PrimFunc(symshen_4newpv), B5972)


__e.TailApply(tmp17877, tmp17888)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp17889 := PrimValue(symshen_4_dsigf_d)

tmp17890 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4app, tmp17876, tmp17889)


tmp17891 := PrimSet(symshen_4_dsigf_d, tmp17890)

_ = tmp17891

tmp17892 := MakeNative(func(__e *ControlFlow) {
V5981 := __e.Get(1)
_ = V5981
__e.Return(MakeNative(func(__e *ControlFlow) {
B5977 := __e.Get(1)
_ = B5977
__e.Return(MakeNative(func(__e *ControlFlow) {
L5978 := __e.Get(1)
_ = L5978
__e.Return(MakeNative(func(__e *ControlFlow) {
Key5979 := __e.Get(1)
_ = Key5979
__e.Return(MakeNative(func(__e *ControlFlow) {
C5980 := __e.Get(1)
_ = C5980
tmp17893 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp17894 := PrimCons(A, Nil)

tmp17895 := PrimCons(symlist, tmp17894)

tmp17896 := PrimCons(A, Nil)

tmp17897 := PrimCons(symlist, tmp17896)

tmp17898 := PrimCons(A, Nil)

tmp17899 := PrimCons(symlist, tmp17898)

tmp17900 := PrimCons(tmp17899, Nil)

tmp17901 := PrimCons(sym_1_1_6, tmp17900)

tmp17902 := PrimCons(tmp17897, tmp17901)

tmp17903 := PrimCons(tmp17902, Nil)

tmp17904 := PrimCons(sym_1_1_6, tmp17903)

tmp17905 := PrimCons(tmp17895, tmp17904)

tmp17906 := Call(__e, PrimFunc(symis_b), V5981, tmp17905, B5977, L5978, Key5979, C5980)


__e.TailApply(PrimFunc(symshen_4gc), B5977, tmp17906)
return


}, 1)

tmp17907 := Call(__e, PrimFunc(symshen_4newpv), B5977)


__e.TailApply(tmp17893, tmp17907)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp17908 := PrimValue(symshen_4_dsigf_d)

tmp17909 := Call(__e, PrimFunc(symshen_4assoc_1_6), symappend, tmp17892, tmp17908)


tmp17910 := PrimSet(symshen_4_dsigf_d, tmp17909)

_ = tmp17910

tmp17911 := MakeNative(func(__e *ControlFlow) {
V5986 := __e.Get(1)
_ = V5986
__e.Return(MakeNative(func(__e *ControlFlow) {
B5982 := __e.Get(1)
_ = B5982
__e.Return(MakeNative(func(__e *ControlFlow) {
L5983 := __e.Get(1)
_ = L5983
__e.Return(MakeNative(func(__e *ControlFlow) {
Key5984 := __e.Get(1)
_ = Key5984
__e.Return(MakeNative(func(__e *ControlFlow) {
C5985 := __e.Get(1)
_ = C5985
tmp17912 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp17913 := PrimCons(symnumber, Nil)

tmp17914 := PrimCons(sym_1_1_6, tmp17913)

tmp17915 := PrimCons(A, tmp17914)

tmp17916 := Call(__e, PrimFunc(symis_b), V5986, tmp17915, B5982, L5983, Key5984, C5985)


__e.TailApply(PrimFunc(symshen_4gc), B5982, tmp17916)
return


}, 1)

tmp17917 := Call(__e, PrimFunc(symshen_4newpv), B5982)


__e.TailApply(tmp17912, tmp17917)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp17918 := PrimValue(symshen_4_dsigf_d)

tmp17919 := Call(__e, PrimFunc(symshen_4assoc_1_6), symarity, tmp17911, tmp17918)


tmp17920 := PrimSet(symshen_4_dsigf_d, tmp17919)

_ = tmp17920

tmp17921 := MakeNative(func(__e *ControlFlow) {
V5991 := __e.Get(1)
_ = V5991
__e.Return(MakeNative(func(__e *ControlFlow) {
B5987 := __e.Get(1)
_ = B5987
__e.Return(MakeNative(func(__e *ControlFlow) {
L5988 := __e.Get(1)
_ = L5988
__e.Return(MakeNative(func(__e *ControlFlow) {
Key5989 := __e.Get(1)
_ = Key5989
__e.Return(MakeNative(func(__e *ControlFlow) {
C5990 := __e.Get(1)
_ = C5990
tmp17922 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp17923 := PrimCons(A, Nil)

tmp17924 := PrimCons(symlist, tmp17923)

tmp17925 := PrimCons(tmp17924, Nil)

tmp17926 := PrimCons(symlist, tmp17925)

tmp17927 := PrimCons(A, Nil)

tmp17928 := PrimCons(symlist, tmp17927)

tmp17929 := PrimCons(tmp17928, Nil)

tmp17930 := PrimCons(sym_1_1_6, tmp17929)

tmp17931 := PrimCons(tmp17926, tmp17930)

tmp17932 := PrimCons(tmp17931, Nil)

tmp17933 := PrimCons(sym_1_1_6, tmp17932)

tmp17934 := PrimCons(A, tmp17933)

tmp17935 := Call(__e, PrimFunc(symis_b), V5991, tmp17934, B5987, L5988, Key5989, C5990)


__e.TailApply(PrimFunc(symshen_4gc), B5987, tmp17935)
return


}, 1)

tmp17936 := Call(__e, PrimFunc(symshen_4newpv), B5987)


__e.TailApply(tmp17922, tmp17936)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp17937 := PrimValue(symshen_4_dsigf_d)

tmp17938 := Call(__e, PrimFunc(symshen_4assoc_1_6), symassoc, tmp17921, tmp17937)


tmp17939 := PrimSet(symshen_4_dsigf_d, tmp17938)

_ = tmp17939

tmp17940 := MakeNative(func(__e *ControlFlow) {
V5996 := __e.Get(1)
_ = V5996
__e.Return(MakeNative(func(__e *ControlFlow) {
B5992 := __e.Get(1)
_ = B5992
__e.Return(MakeNative(func(__e *ControlFlow) {
L5993 := __e.Get(1)
_ = L5993
__e.Return(MakeNative(func(__e *ControlFlow) {
Key5994 := __e.Get(1)
_ = Key5994
__e.Return(MakeNative(func(__e *ControlFlow) {
C5995 := __e.Get(1)
_ = C5995
tmp17941 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp17942 := PrimCons(symboolean, Nil)

tmp17943 := PrimCons(sym_1_1_6, tmp17942)

tmp17944 := PrimCons(A, tmp17943)

tmp17945 := Call(__e, PrimFunc(symis_b), V5996, tmp17944, B5992, L5993, Key5994, C5995)


__e.TailApply(PrimFunc(symshen_4gc), B5992, tmp17945)
return


}, 1)

tmp17946 := Call(__e, PrimFunc(symshen_4newpv), B5992)


__e.TailApply(tmp17941, tmp17946)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp17947 := PrimValue(symshen_4_dsigf_d)

tmp17948 := Call(__e, PrimFunc(symshen_4assoc_1_6), symatom_2, tmp17940, tmp17947)


tmp17949 := PrimSet(symshen_4_dsigf_d, tmp17948)

_ = tmp17949

tmp17950 := MakeNative(func(__e *ControlFlow) {
V6001 := __e.Get(1)
_ = V6001
__e.Return(MakeNative(func(__e *ControlFlow) {
B5997 := __e.Get(1)
_ = B5997
__e.Return(MakeNative(func(__e *ControlFlow) {
L5998 := __e.Get(1)
_ = L5998
__e.Return(MakeNative(func(__e *ControlFlow) {
Key5999 := __e.Get(1)
_ = Key5999
__e.Return(MakeNative(func(__e *ControlFlow) {
C6000 := __e.Get(1)
_ = C6000
tmp17951 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp17952 := PrimCons(symboolean, Nil)

tmp17953 := PrimCons(sym_1_1_6, tmp17952)

tmp17954 := PrimCons(A, tmp17953)

tmp17955 := Call(__e, PrimFunc(symis_b), V6001, tmp17954, B5997, L5998, Key5999, C6000)


__e.TailApply(PrimFunc(symshen_4gc), B5997, tmp17955)
return


}, 1)

tmp17956 := Call(__e, PrimFunc(symshen_4newpv), B5997)


__e.TailApply(tmp17951, tmp17956)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp17957 := PrimValue(symshen_4_dsigf_d)

tmp17958 := Call(__e, PrimFunc(symshen_4assoc_1_6), symboolean_2, tmp17950, tmp17957)


tmp17959 := PrimSet(symshen_4_dsigf_d, tmp17958)

_ = tmp17959

tmp17960 := MakeNative(func(__e *ControlFlow) {
V6006 := __e.Get(1)
_ = V6006
__e.Return(MakeNative(func(__e *ControlFlow) {
B6002 := __e.Get(1)
_ = B6002
__e.Return(MakeNative(func(__e *ControlFlow) {
L6003 := __e.Get(1)
_ = L6003
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6004 := __e.Get(1)
_ = Key6004
__e.Return(MakeNative(func(__e *ControlFlow) {
C6005 := __e.Get(1)
_ = C6005
tmp17961 := PrimCons(symstring, Nil)

tmp17962 := PrimCons(sym_1_1_6, tmp17961)

tmp17963 := PrimCons(symstring, tmp17962)

__e.TailApply(PrimFunc(symis_b), V6006, tmp17963, B6002, L6003, Key6004, C6005)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp17964 := PrimValue(symshen_4_dsigf_d)

tmp17965 := Call(__e, PrimFunc(symshen_4assoc_1_6), symbootstrap, tmp17960, tmp17964)


tmp17966 := PrimSet(symshen_4_dsigf_d, tmp17965)

_ = tmp17966

tmp17967 := MakeNative(func(__e *ControlFlow) {
V6011 := __e.Get(1)
_ = V6011
__e.Return(MakeNative(func(__e *ControlFlow) {
B6007 := __e.Get(1)
_ = B6007
__e.Return(MakeNative(func(__e *ControlFlow) {
L6008 := __e.Get(1)
_ = L6008
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6009 := __e.Get(1)
_ = Key6009
__e.Return(MakeNative(func(__e *ControlFlow) {
C6010 := __e.Get(1)
_ = C6010
tmp17968 := PrimCons(symboolean, Nil)

tmp17969 := PrimCons(sym_1_1_6, tmp17968)

tmp17970 := PrimCons(symsymbol, tmp17969)

__e.TailApply(PrimFunc(symis_b), V6011, tmp17970, B6007, L6008, Key6009, C6010)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp17971 := PrimValue(symshen_4_dsigf_d)

tmp17972 := Call(__e, PrimFunc(symshen_4assoc_1_6), symbound_2, tmp17967, tmp17971)


tmp17973 := PrimSet(symshen_4_dsigf_d, tmp17972)

_ = tmp17973

tmp17974 := MakeNative(func(__e *ControlFlow) {
V6016 := __e.Get(1)
_ = V6016
__e.Return(MakeNative(func(__e *ControlFlow) {
B6012 := __e.Get(1)
_ = B6012
__e.Return(MakeNative(func(__e *ControlFlow) {
L6013 := __e.Get(1)
_ = L6013
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6014 := __e.Get(1)
_ = Key6014
__e.Return(MakeNative(func(__e *ControlFlow) {
C6015 := __e.Get(1)
_ = C6015
tmp17975 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp17976 := PrimCons(A, Nil)

tmp17977 := PrimCons(symlist, tmp17976)

tmp17978 := PrimCons(symboolean, Nil)

tmp17979 := PrimCons(sym_1_1_6, tmp17978)

tmp17980 := PrimCons(tmp17977, tmp17979)

tmp17981 := Call(__e, PrimFunc(symis_b), V6016, tmp17980, B6012, L6013, Key6014, C6015)


__e.TailApply(PrimFunc(symshen_4gc), B6012, tmp17981)
return


}, 1)

tmp17982 := Call(__e, PrimFunc(symshen_4newpv), B6012)


__e.TailApply(tmp17975, tmp17982)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp17983 := PrimValue(symshen_4_dsigf_d)

tmp17984 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4ccons_2, tmp17974, tmp17983)


tmp17985 := PrimSet(symshen_4_dsigf_d, tmp17984)

_ = tmp17985

tmp17986 := MakeNative(func(__e *ControlFlow) {
V6021 := __e.Get(1)
_ = V6021
__e.Return(MakeNative(func(__e *ControlFlow) {
B6017 := __e.Get(1)
_ = B6017
__e.Return(MakeNative(func(__e *ControlFlow) {
L6018 := __e.Get(1)
_ = L6018
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6019 := __e.Get(1)
_ = Key6019
__e.Return(MakeNative(func(__e *ControlFlow) {
C6020 := __e.Get(1)
_ = C6020
tmp17987 := PrimCons(symstring, Nil)

tmp17988 := PrimCons(sym_1_1_6, tmp17987)

tmp17989 := PrimCons(symstring, tmp17988)

__e.TailApply(PrimFunc(symis_b), V6021, tmp17989, B6017, L6018, Key6019, C6020)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp17990 := PrimValue(symshen_4_dsigf_d)

tmp17991 := Call(__e, PrimFunc(symshen_4assoc_1_6), symcd, tmp17986, tmp17990)


tmp17992 := PrimSet(symshen_4_dsigf_d, tmp17991)

_ = tmp17992

tmp17993 := MakeNative(func(__e *ControlFlow) {
V6026 := __e.Get(1)
_ = V6026
__e.Return(MakeNative(func(__e *ControlFlow) {
B6022 := __e.Get(1)
_ = B6022
__e.Return(MakeNative(func(__e *ControlFlow) {
L6023 := __e.Get(1)
_ = L6023
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6024 := __e.Get(1)
_ = Key6024
__e.Return(MakeNative(func(__e *ControlFlow) {
C6025 := __e.Get(1)
_ = C6025
tmp17994 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp17995 := MakeNative(func(__e *ControlFlow) {
B := __e.Get(1)
_ = B
tmp17996 := PrimCons(A, Nil)

tmp17997 := PrimCons(symstream, tmp17996)

tmp17998 := PrimCons(B, Nil)

tmp17999 := PrimCons(symlist, tmp17998)

tmp18000 := PrimCons(tmp17999, Nil)

tmp18001 := PrimCons(sym_1_1_6, tmp18000)

tmp18002 := PrimCons(tmp17997, tmp18001)

tmp18003 := Call(__e, PrimFunc(symis_b), V6026, tmp18002, B6022, L6023, Key6024, C6025)


__e.TailApply(PrimFunc(symshen_4gc), B6022, tmp18003)
return


}, 1)

tmp18004 := Call(__e, PrimFunc(symshen_4newpv), B6022)


tmp18005 := Call(__e, tmp17995, tmp18004)


__e.TailApply(PrimFunc(symshen_4gc), B6022, tmp18005)
return


}, 1)

tmp18006 := Call(__e, PrimFunc(symshen_4newpv), B6022)


__e.TailApply(tmp17994, tmp18006)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18007 := PrimValue(symshen_4_dsigf_d)

tmp18008 := Call(__e, PrimFunc(symshen_4assoc_1_6), symclose, tmp17993, tmp18007)


tmp18009 := PrimSet(symshen_4_dsigf_d, tmp18008)

_ = tmp18009

tmp18010 := MakeNative(func(__e *ControlFlow) {
V6031 := __e.Get(1)
_ = V6031
__e.Return(MakeNative(func(__e *ControlFlow) {
B6027 := __e.Get(1)
_ = B6027
__e.Return(MakeNative(func(__e *ControlFlow) {
L6028 := __e.Get(1)
_ = L6028
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6029 := __e.Get(1)
_ = Key6029
__e.Return(MakeNative(func(__e *ControlFlow) {
C6030 := __e.Get(1)
_ = C6030
tmp18011 := PrimCons(symstring, Nil)

tmp18012 := PrimCons(sym_1_1_6, tmp18011)

tmp18013 := PrimCons(symstring, tmp18012)

tmp18014 := PrimCons(tmp18013, Nil)

tmp18015 := PrimCons(sym_1_1_6, tmp18014)

tmp18016 := PrimCons(symstring, tmp18015)

__e.TailApply(PrimFunc(symis_b), V6031, tmp18016, B6027, L6028, Key6029, C6030)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18017 := PrimValue(symshen_4_dsigf_d)

tmp18018 := Call(__e, PrimFunc(symshen_4assoc_1_6), symcn, tmp18010, tmp18017)


tmp18019 := PrimSet(symshen_4_dsigf_d, tmp18018)

_ = tmp18019

tmp18020 := MakeNative(func(__e *ControlFlow) {
V6036 := __e.Get(1)
_ = V6036
__e.Return(MakeNative(func(__e *ControlFlow) {
B6032 := __e.Get(1)
_ = B6032
__e.Return(MakeNative(func(__e *ControlFlow) {
L6033 := __e.Get(1)
_ = L6033
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6034 := __e.Get(1)
_ = Key6034
__e.Return(MakeNative(func(__e *ControlFlow) {
C6035 := __e.Get(1)
_ = C6035
tmp18021 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp18022 := MakeNative(func(__e *ControlFlow) {
B := __e.Get(1)
_ = B
tmp18023 := PrimCons(A, Nil)

tmp18024 := PrimCons(symlist, tmp18023)

tmp18025 := PrimCons(A, Nil)

tmp18026 := PrimCons(symlist, tmp18025)

tmp18027 := PrimCons(B, Nil)

tmp18028 := PrimCons(tmp18026, tmp18027)

tmp18029 := PrimCons(symstr, tmp18028)

tmp18030 := PrimCons(tmp18029, Nil)

tmp18031 := PrimCons(sym_1_1_6, tmp18030)

tmp18032 := PrimCons(tmp18024, tmp18031)

tmp18033 := PrimCons(A, Nil)

tmp18034 := PrimCons(symlist, tmp18033)

tmp18035 := PrimCons(B, Nil)

tmp18036 := PrimCons(sym_1_1_6, tmp18035)

tmp18037 := PrimCons(tmp18034, tmp18036)

tmp18038 := PrimCons(tmp18037, Nil)

tmp18039 := PrimCons(sym_1_1_6, tmp18038)

tmp18040 := PrimCons(tmp18032, tmp18039)

tmp18041 := Call(__e, PrimFunc(symis_b), V6036, tmp18040, B6032, L6033, Key6034, C6035)


__e.TailApply(PrimFunc(symshen_4gc), B6032, tmp18041)
return


}, 1)

tmp18042 := Call(__e, PrimFunc(symshen_4newpv), B6032)


tmp18043 := Call(__e, tmp18022, tmp18042)


__e.TailApply(PrimFunc(symshen_4gc), B6032, tmp18043)
return


}, 1)

tmp18044 := Call(__e, PrimFunc(symshen_4newpv), B6032)


__e.TailApply(tmp18021, tmp18044)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18045 := PrimValue(symshen_4_dsigf_d)

tmp18046 := Call(__e, PrimFunc(symshen_4assoc_1_6), symcompile, tmp18020, tmp18045)


tmp18047 := PrimSet(symshen_4_dsigf_d, tmp18046)

_ = tmp18047

tmp18048 := MakeNative(func(__e *ControlFlow) {
V6041 := __e.Get(1)
_ = V6041
__e.Return(MakeNative(func(__e *ControlFlow) {
B6037 := __e.Get(1)
_ = B6037
__e.Return(MakeNative(func(__e *ControlFlow) {
L6038 := __e.Get(1)
_ = L6038
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6039 := __e.Get(1)
_ = Key6039
__e.Return(MakeNative(func(__e *ControlFlow) {
C6040 := __e.Get(1)
_ = C6040
tmp18049 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp18050 := PrimCons(symboolean, Nil)

tmp18051 := PrimCons(sym_1_1_6, tmp18050)

tmp18052 := PrimCons(A, tmp18051)

tmp18053 := Call(__e, PrimFunc(symis_b), V6041, tmp18052, B6037, L6038, Key6039, C6040)


__e.TailApply(PrimFunc(symshen_4gc), B6037, tmp18053)
return


}, 1)

tmp18054 := Call(__e, PrimFunc(symshen_4newpv), B6037)


__e.TailApply(tmp18049, tmp18054)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18055 := PrimValue(symshen_4_dsigf_d)

tmp18056 := Call(__e, PrimFunc(symshen_4assoc_1_6), symcons_2, tmp18048, tmp18055)


tmp18057 := PrimSet(symshen_4_dsigf_d, tmp18056)

_ = tmp18057

tmp18058 := MakeNative(func(__e *ControlFlow) {
V6046 := __e.Get(1)
_ = V6046
__e.Return(MakeNative(func(__e *ControlFlow) {
B6042 := __e.Get(1)
_ = B6042
__e.Return(MakeNative(func(__e *ControlFlow) {
L6043 := __e.Get(1)
_ = L6043
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6044 := __e.Get(1)
_ = Key6044
__e.Return(MakeNative(func(__e *ControlFlow) {
C6045 := __e.Get(1)
_ = C6045
tmp18059 := PrimCons(symsymbol, Nil)

tmp18060 := PrimCons(symlist, tmp18059)

tmp18061 := PrimCons(tmp18060, Nil)

tmp18062 := PrimCons(sym_1_1_6, tmp18061)

__e.TailApply(PrimFunc(symis_b), V6046, tmp18062, B6042, L6043, Key6044, C6045)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18063 := PrimValue(symshen_4_dsigf_d)

tmp18064 := Call(__e, PrimFunc(symshen_4assoc_1_6), symdatatypes, tmp18058, tmp18063)


tmp18065 := PrimSet(symshen_4_dsigf_d, tmp18064)

_ = tmp18065

tmp18066 := MakeNative(func(__e *ControlFlow) {
V6051 := __e.Get(1)
_ = V6051
__e.Return(MakeNative(func(__e *ControlFlow) {
B6047 := __e.Get(1)
_ = B6047
__e.Return(MakeNative(func(__e *ControlFlow) {
L6048 := __e.Get(1)
_ = L6048
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6049 := __e.Get(1)
_ = Key6049
__e.Return(MakeNative(func(__e *ControlFlow) {
C6050 := __e.Get(1)
_ = C6050
tmp18067 := PrimCons(symsymbol, Nil)

tmp18068 := PrimCons(sym_1_1_6, tmp18067)

tmp18069 := PrimCons(symsymbol, tmp18068)

__e.TailApply(PrimFunc(symis_b), V6051, tmp18069, B6047, L6048, Key6049, C6050)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18070 := PrimValue(symshen_4_dsigf_d)

tmp18071 := Call(__e, PrimFunc(symshen_4assoc_1_6), symdestroy, tmp18066, tmp18070)


tmp18072 := PrimSet(symshen_4_dsigf_d, tmp18071)

_ = tmp18072

tmp18073 := MakeNative(func(__e *ControlFlow) {
V6056 := __e.Get(1)
_ = V6056
__e.Return(MakeNative(func(__e *ControlFlow) {
B6052 := __e.Get(1)
_ = B6052
__e.Return(MakeNative(func(__e *ControlFlow) {
L6053 := __e.Get(1)
_ = L6053
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6054 := __e.Get(1)
_ = Key6054
__e.Return(MakeNative(func(__e *ControlFlow) {
C6055 := __e.Get(1)
_ = C6055
tmp18074 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp18075 := PrimCons(A, Nil)

tmp18076 := PrimCons(symlist, tmp18075)

tmp18077 := PrimCons(A, Nil)

tmp18078 := PrimCons(symlist, tmp18077)

tmp18079 := PrimCons(A, Nil)

tmp18080 := PrimCons(symlist, tmp18079)

tmp18081 := PrimCons(tmp18080, Nil)

tmp18082 := PrimCons(sym_1_1_6, tmp18081)

tmp18083 := PrimCons(tmp18078, tmp18082)

tmp18084 := PrimCons(tmp18083, Nil)

tmp18085 := PrimCons(sym_1_1_6, tmp18084)

tmp18086 := PrimCons(tmp18076, tmp18085)

tmp18087 := Call(__e, PrimFunc(symis_b), V6056, tmp18086, B6052, L6053, Key6054, C6055)


__e.TailApply(PrimFunc(symshen_4gc), B6052, tmp18087)
return


}, 1)

tmp18088 := Call(__e, PrimFunc(symshen_4newpv), B6052)


__e.TailApply(tmp18074, tmp18088)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18089 := PrimValue(symshen_4_dsigf_d)

tmp18090 := Call(__e, PrimFunc(symshen_4assoc_1_6), symdifference, tmp18073, tmp18089)


tmp18091 := PrimSet(symshen_4_dsigf_d, tmp18090)

_ = tmp18091

tmp18092 := MakeNative(func(__e *ControlFlow) {
V6061 := __e.Get(1)
_ = V6061
__e.Return(MakeNative(func(__e *ControlFlow) {
B6057 := __e.Get(1)
_ = B6057
__e.Return(MakeNative(func(__e *ControlFlow) {
L6058 := __e.Get(1)
_ = L6058
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6059 := __e.Get(1)
_ = Key6059
__e.Return(MakeNative(func(__e *ControlFlow) {
C6060 := __e.Get(1)
_ = C6060
tmp18093 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp18094 := MakeNative(func(__e *ControlFlow) {
B := __e.Get(1)
_ = B
tmp18095 := PrimCons(B, Nil)

tmp18096 := PrimCons(sym_1_1_6, tmp18095)

tmp18097 := PrimCons(B, tmp18096)

tmp18098 := PrimCons(tmp18097, Nil)

tmp18099 := PrimCons(sym_1_1_6, tmp18098)

tmp18100 := PrimCons(A, tmp18099)

tmp18101 := Call(__e, PrimFunc(symis_b), V6061, tmp18100, B6057, L6058, Key6059, C6060)


__e.TailApply(PrimFunc(symshen_4gc), B6057, tmp18101)
return


}, 1)

tmp18102 := Call(__e, PrimFunc(symshen_4newpv), B6057)


tmp18103 := Call(__e, tmp18094, tmp18102)


__e.TailApply(PrimFunc(symshen_4gc), B6057, tmp18103)
return


}, 1)

tmp18104 := Call(__e, PrimFunc(symshen_4newpv), B6057)


__e.TailApply(tmp18093, tmp18104)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18105 := PrimValue(symshen_4_dsigf_d)

tmp18106 := Call(__e, PrimFunc(symshen_4assoc_1_6), symdo, tmp18092, tmp18105)


tmp18107 := PrimSet(symshen_4_dsigf_d, tmp18106)

_ = tmp18107

tmp18108 := MakeNative(func(__e *ControlFlow) {
V6066 := __e.Get(1)
_ = V6066
__e.Return(MakeNative(func(__e *ControlFlow) {
B6062 := __e.Get(1)
_ = B6062
__e.Return(MakeNative(func(__e *ControlFlow) {
L6063 := __e.Get(1)
_ = L6063
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6064 := __e.Get(1)
_ = Key6064
__e.Return(MakeNative(func(__e *ControlFlow) {
C6065 := __e.Get(1)
_ = C6065
tmp18109 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp18110 := MakeNative(func(__e *ControlFlow) {
B := __e.Get(1)
_ = B
tmp18111 := PrimCons(A, Nil)

tmp18112 := PrimCons(symlist, tmp18111)

tmp18113 := PrimCons(A, Nil)

tmp18114 := PrimCons(symlist, tmp18113)

tmp18115 := PrimCons(B, Nil)

tmp18116 := PrimCons(symlist, tmp18115)

tmp18117 := PrimCons(tmp18116, Nil)

tmp18118 := PrimCons(tmp18114, tmp18117)

tmp18119 := PrimCons(symstr, tmp18118)

tmp18120 := PrimCons(tmp18119, Nil)

tmp18121 := PrimCons(sym_1_1_6, tmp18120)

tmp18122 := PrimCons(tmp18112, tmp18121)

tmp18123 := Call(__e, PrimFunc(symis_b), V6066, tmp18122, B6062, L6063, Key6064, C6065)


__e.TailApply(PrimFunc(symshen_4gc), B6062, tmp18123)
return


}, 1)

tmp18124 := Call(__e, PrimFunc(symshen_4newpv), B6062)


tmp18125 := Call(__e, tmp18110, tmp18124)


__e.TailApply(PrimFunc(symshen_4gc), B6062, tmp18125)
return


}, 1)

tmp18126 := Call(__e, PrimFunc(symshen_4newpv), B6062)


__e.TailApply(tmp18109, tmp18126)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18127 := PrimValue(symshen_4_dsigf_d)

tmp18128 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_5e_6, tmp18108, tmp18127)


tmp18129 := PrimSet(symshen_4_dsigf_d, tmp18128)

_ = tmp18129

tmp18130 := MakeNative(func(__e *ControlFlow) {
V6071 := __e.Get(1)
_ = V6071
__e.Return(MakeNative(func(__e *ControlFlow) {
B6067 := __e.Get(1)
_ = B6067
__e.Return(MakeNative(func(__e *ControlFlow) {
L6068 := __e.Get(1)
_ = L6068
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6069 := __e.Get(1)
_ = Key6069
__e.Return(MakeNative(func(__e *ControlFlow) {
C6070 := __e.Get(1)
_ = C6070
tmp18131 := MakeNative(func(__e *ControlFlow) {
B := __e.Get(1)
_ = B
tmp18132 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp18133 := PrimCons(A, Nil)

tmp18134 := PrimCons(symlist, tmp18133)

tmp18135 := PrimCons(B, Nil)

tmp18136 := PrimCons(symlist, tmp18135)

tmp18137 := PrimCons(A, Nil)

tmp18138 := PrimCons(symlist, tmp18137)

tmp18139 := PrimCons(tmp18138, Nil)

tmp18140 := PrimCons(tmp18136, tmp18139)

tmp18141 := PrimCons(symstr, tmp18140)

tmp18142 := PrimCons(tmp18141, Nil)

tmp18143 := PrimCons(sym_1_1_6, tmp18142)

tmp18144 := PrimCons(tmp18134, tmp18143)

tmp18145 := Call(__e, PrimFunc(symis_b), V6071, tmp18144, B6067, L6068, Key6069, C6070)


__e.TailApply(PrimFunc(symshen_4gc), B6067, tmp18145)
return


}, 1)

tmp18146 := Call(__e, PrimFunc(symshen_4newpv), B6067)


tmp18147 := Call(__e, tmp18132, tmp18146)


__e.TailApply(PrimFunc(symshen_4gc), B6067, tmp18147)
return


}, 1)

tmp18148 := Call(__e, PrimFunc(symshen_4newpv), B6067)


__e.TailApply(tmp18131, tmp18148)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18149 := PrimValue(symshen_4_dsigf_d)

tmp18150 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_5_b_6, tmp18130, tmp18149)


tmp18151 := PrimSet(symshen_4_dsigf_d, tmp18150)

_ = tmp18151

tmp18152 := MakeNative(func(__e *ControlFlow) {
V6076 := __e.Get(1)
_ = V6076
__e.Return(MakeNative(func(__e *ControlFlow) {
B6072 := __e.Get(1)
_ = B6072
__e.Return(MakeNative(func(__e *ControlFlow) {
L6073 := __e.Get(1)
_ = L6073
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6074 := __e.Get(1)
_ = Key6074
__e.Return(MakeNative(func(__e *ControlFlow) {
C6075 := __e.Get(1)
_ = C6075
tmp18153 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp18154 := MakeNative(func(__e *ControlFlow) {
B := __e.Get(1)
_ = B
tmp18155 := PrimCons(A, Nil)

tmp18156 := PrimCons(symlist, tmp18155)

tmp18157 := PrimCons(A, Nil)

tmp18158 := PrimCons(symlist, tmp18157)

tmp18159 := PrimCons(B, Nil)

tmp18160 := PrimCons(symlist, tmp18159)

tmp18161 := PrimCons(tmp18160, Nil)

tmp18162 := PrimCons(tmp18158, tmp18161)

tmp18163 := PrimCons(symstr, tmp18162)

tmp18164 := PrimCons(tmp18163, Nil)

tmp18165 := PrimCons(sym_1_1_6, tmp18164)

tmp18166 := PrimCons(tmp18156, tmp18165)

tmp18167 := Call(__e, PrimFunc(symis_b), V6076, tmp18166, B6072, L6073, Key6074, C6075)


__e.TailApply(PrimFunc(symshen_4gc), B6072, tmp18167)
return


}, 1)

tmp18168 := Call(__e, PrimFunc(symshen_4newpv), B6072)


tmp18169 := Call(__e, tmp18154, tmp18168)


__e.TailApply(PrimFunc(symshen_4gc), B6072, tmp18169)
return


}, 1)

tmp18170 := Call(__e, PrimFunc(symshen_4newpv), B6072)


__e.TailApply(tmp18153, tmp18170)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18171 := PrimValue(symshen_4_dsigf_d)

tmp18172 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_5end_6, tmp18152, tmp18171)


tmp18173 := PrimSet(symshen_4_dsigf_d, tmp18172)

_ = tmp18173

tmp18174 := MakeNative(func(__e *ControlFlow) {
V6081 := __e.Get(1)
_ = V6081
__e.Return(MakeNative(func(__e *ControlFlow) {
B6077 := __e.Get(1)
_ = B6077
__e.Return(MakeNative(func(__e *ControlFlow) {
L6078 := __e.Get(1)
_ = L6078
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6079 := __e.Get(1)
_ = Key6079
__e.Return(MakeNative(func(__e *ControlFlow) {
C6080 := __e.Get(1)
_ = C6080
tmp18175 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp18176 := MakeNative(func(__e *ControlFlow) {
B := __e.Get(1)
_ = B
tmp18177 := PrimCons(A, Nil)

tmp18178 := PrimCons(symlist, tmp18177)

tmp18179 := PrimCons(B, Nil)

tmp18180 := PrimCons(tmp18178, tmp18179)

tmp18181 := PrimCons(symstr, tmp18180)

tmp18182 := PrimCons(symboolean, Nil)

tmp18183 := PrimCons(sym_1_1_6, tmp18182)

tmp18184 := PrimCons(tmp18181, tmp18183)

tmp18185 := Call(__e, PrimFunc(symis_b), V6081, tmp18184, B6077, L6078, Key6079, C6080)


__e.TailApply(PrimFunc(symshen_4gc), B6077, tmp18185)
return


}, 1)

tmp18186 := Call(__e, PrimFunc(symshen_4newpv), B6077)


tmp18187 := Call(__e, tmp18176, tmp18186)


__e.TailApply(PrimFunc(symshen_4gc), B6077, tmp18187)
return


}, 1)

tmp18188 := Call(__e, PrimFunc(symshen_4newpv), B6077)


__e.TailApply(tmp18175, tmp18188)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18189 := PrimValue(symshen_4_dsigf_d)

tmp18190 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4parse_1failure_2, tmp18174, tmp18189)


tmp18191 := PrimSet(symshen_4_dsigf_d, tmp18190)

_ = tmp18191

tmp18192 := MakeNative(func(__e *ControlFlow) {
V6086 := __e.Get(1)
_ = V6086
__e.Return(MakeNative(func(__e *ControlFlow) {
B6082 := __e.Get(1)
_ = B6082
__e.Return(MakeNative(func(__e *ControlFlow) {
L6083 := __e.Get(1)
_ = L6083
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6084 := __e.Get(1)
_ = Key6084
__e.Return(MakeNative(func(__e *ControlFlow) {
C6085 := __e.Get(1)
_ = C6085
tmp18193 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp18194 := MakeNative(func(__e *ControlFlow) {
B := __e.Get(1)
_ = B
tmp18195 := PrimCons(A, Nil)

tmp18196 := PrimCons(symlist, tmp18195)

tmp18197 := PrimCons(B, Nil)

tmp18198 := PrimCons(tmp18196, tmp18197)

tmp18199 := PrimCons(symstr, tmp18198)

tmp18200 := PrimCons(tmp18199, Nil)

tmp18201 := PrimCons(sym_1_1_6, tmp18200)

tmp18202 := Call(__e, PrimFunc(symis_b), V6086, tmp18201, B6082, L6083, Key6084, C6085)


__e.TailApply(PrimFunc(symshen_4gc), B6082, tmp18202)
return


}, 1)

tmp18203 := Call(__e, PrimFunc(symshen_4newpv), B6082)


tmp18204 := Call(__e, tmp18194, tmp18203)


__e.TailApply(PrimFunc(symshen_4gc), B6082, tmp18204)
return


}, 1)

tmp18205 := Call(__e, PrimFunc(symshen_4newpv), B6082)


__e.TailApply(tmp18193, tmp18205)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18206 := PrimValue(symshen_4_dsigf_d)

tmp18207 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4parse_1failure, tmp18192, tmp18206)


tmp18208 := PrimSet(symshen_4_dsigf_d, tmp18207)

_ = tmp18208

tmp18209 := MakeNative(func(__e *ControlFlow) {
V6091 := __e.Get(1)
_ = V6091
__e.Return(MakeNative(func(__e *ControlFlow) {
B6087 := __e.Get(1)
_ = B6087
__e.Return(MakeNative(func(__e *ControlFlow) {
L6088 := __e.Get(1)
_ = L6088
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6089 := __e.Get(1)
_ = Key6089
__e.Return(MakeNative(func(__e *ControlFlow) {
C6090 := __e.Get(1)
_ = C6090
tmp18210 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp18211 := MakeNative(func(__e *ControlFlow) {
B := __e.Get(1)
_ = B
tmp18212 := PrimCons(A, Nil)

tmp18213 := PrimCons(symlist, tmp18212)

tmp18214 := PrimCons(B, Nil)

tmp18215 := PrimCons(tmp18213, tmp18214)

tmp18216 := PrimCons(symstr, tmp18215)

tmp18217 := PrimCons(B, Nil)

tmp18218 := PrimCons(sym_1_1_6, tmp18217)

tmp18219 := PrimCons(tmp18216, tmp18218)

tmp18220 := Call(__e, PrimFunc(symis_b), V6091, tmp18219, B6087, L6088, Key6089, C6090)


__e.TailApply(PrimFunc(symshen_4gc), B6087, tmp18220)
return


}, 1)

tmp18221 := Call(__e, PrimFunc(symshen_4newpv), B6087)


tmp18222 := Call(__e, tmp18211, tmp18221)


__e.TailApply(PrimFunc(symshen_4gc), B6087, tmp18222)
return


}, 1)

tmp18223 := Call(__e, PrimFunc(symshen_4newpv), B6087)


__e.TailApply(tmp18210, tmp18223)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18224 := PrimValue(symshen_4_dsigf_d)

tmp18225 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4_5_1out, tmp18209, tmp18224)


tmp18226 := PrimSet(symshen_4_dsigf_d, tmp18225)

_ = tmp18226

tmp18227 := MakeNative(func(__e *ControlFlow) {
V6096 := __e.Get(1)
_ = V6096
__e.Return(MakeNative(func(__e *ControlFlow) {
B6092 := __e.Get(1)
_ = B6092
__e.Return(MakeNative(func(__e *ControlFlow) {
L6093 := __e.Get(1)
_ = L6093
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6094 := __e.Get(1)
_ = Key6094
__e.Return(MakeNative(func(__e *ControlFlow) {
C6095 := __e.Get(1)
_ = C6095
tmp18228 := MakeNative(func(__e *ControlFlow) {
B := __e.Get(1)
_ = B
tmp18229 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp18230 := PrimCons(A, Nil)

tmp18231 := PrimCons(symlist, tmp18230)

tmp18232 := PrimCons(B, Nil)

tmp18233 := PrimCons(tmp18231, tmp18232)

tmp18234 := PrimCons(symstr, tmp18233)

tmp18235 := PrimCons(A, Nil)

tmp18236 := PrimCons(symlist, tmp18235)

tmp18237 := PrimCons(tmp18236, Nil)

tmp18238 := PrimCons(sym_1_1_6, tmp18237)

tmp18239 := PrimCons(tmp18234, tmp18238)

tmp18240 := Call(__e, PrimFunc(symis_b), V6096, tmp18239, B6092, L6093, Key6094, C6095)


__e.TailApply(PrimFunc(symshen_4gc), B6092, tmp18240)
return


}, 1)

tmp18241 := Call(__e, PrimFunc(symshen_4newpv), B6092)


tmp18242 := Call(__e, tmp18229, tmp18241)


__e.TailApply(PrimFunc(symshen_4gc), B6092, tmp18242)
return


}, 1)

tmp18243 := Call(__e, PrimFunc(symshen_4newpv), B6092)


__e.TailApply(tmp18228, tmp18243)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18244 := PrimValue(symshen_4_dsigf_d)

tmp18245 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4in_1_6, tmp18227, tmp18244)


tmp18246 := PrimSet(symshen_4_dsigf_d, tmp18245)

_ = tmp18246

tmp18247 := MakeNative(func(__e *ControlFlow) {
V6101 := __e.Get(1)
_ = V6101
__e.Return(MakeNative(func(__e *ControlFlow) {
B6097 := __e.Get(1)
_ = B6097
__e.Return(MakeNative(func(__e *ControlFlow) {
L6098 := __e.Get(1)
_ = L6098
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6099 := __e.Get(1)
_ = Key6099
__e.Return(MakeNative(func(__e *ControlFlow) {
C6100 := __e.Get(1)
_ = C6100
tmp18248 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp18249 := MakeNative(func(__e *ControlFlow) {
B := __e.Get(1)
_ = B
tmp18250 := PrimCons(A, Nil)

tmp18251 := PrimCons(symlist, tmp18250)

tmp18252 := PrimCons(A, Nil)

tmp18253 := PrimCons(symlist, tmp18252)

tmp18254 := PrimCons(B, Nil)

tmp18255 := PrimCons(tmp18253, tmp18254)

tmp18256 := PrimCons(symstr, tmp18255)

tmp18257 := PrimCons(tmp18256, Nil)

tmp18258 := PrimCons(sym_1_1_6, tmp18257)

tmp18259 := PrimCons(B, tmp18258)

tmp18260 := PrimCons(tmp18259, Nil)

tmp18261 := PrimCons(sym_1_1_6, tmp18260)

tmp18262 := PrimCons(tmp18251, tmp18261)

tmp18263 := Call(__e, PrimFunc(symis_b), V6101, tmp18262, B6097, L6098, Key6099, C6100)


__e.TailApply(PrimFunc(symshen_4gc), B6097, tmp18263)
return


}, 1)

tmp18264 := Call(__e, PrimFunc(symshen_4newpv), B6097)


tmp18265 := Call(__e, tmp18249, tmp18264)


__e.TailApply(PrimFunc(symshen_4gc), B6097, tmp18265)
return


}, 1)

tmp18266 := Call(__e, PrimFunc(symshen_4newpv), B6097)


__e.TailApply(tmp18248, tmp18266)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18267 := PrimValue(symshen_4_dsigf_d)

tmp18268 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4comb, tmp18247, tmp18267)


tmp18269 := PrimSet(symshen_4_dsigf_d, tmp18268)

_ = tmp18269

tmp18270 := MakeNative(func(__e *ControlFlow) {
V6106 := __e.Get(1)
_ = V6106
__e.Return(MakeNative(func(__e *ControlFlow) {
B6102 := __e.Get(1)
_ = B6102
__e.Return(MakeNative(func(__e *ControlFlow) {
L6103 := __e.Get(1)
_ = L6103
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6104 := __e.Get(1)
_ = Key6104
__e.Return(MakeNative(func(__e *ControlFlow) {
C6105 := __e.Get(1)
_ = C6105
tmp18271 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp18272 := PrimCons(A, Nil)

tmp18273 := PrimCons(symlist, tmp18272)

tmp18274 := PrimCons(symboolean, Nil)

tmp18275 := PrimCons(sym_1_1_6, tmp18274)

tmp18276 := PrimCons(tmp18273, tmp18275)

tmp18277 := PrimCons(tmp18276, Nil)

tmp18278 := PrimCons(sym_1_1_6, tmp18277)

tmp18279 := PrimCons(A, tmp18278)

tmp18280 := Call(__e, PrimFunc(symis_b), V6106, tmp18279, B6102, L6103, Key6104, C6105)


__e.TailApply(PrimFunc(symshen_4gc), B6102, tmp18280)
return


}, 1)

tmp18281 := Call(__e, PrimFunc(symshen_4newpv), B6102)


__e.TailApply(tmp18271, tmp18281)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18282 := PrimValue(symshen_4_dsigf_d)

tmp18283 := Call(__e, PrimFunc(symshen_4assoc_1_6), symelement_2, tmp18270, tmp18282)


tmp18284 := PrimSet(symshen_4_dsigf_d, tmp18283)

_ = tmp18284

tmp18285 := MakeNative(func(__e *ControlFlow) {
V6111 := __e.Get(1)
_ = V6111
__e.Return(MakeNative(func(__e *ControlFlow) {
B6107 := __e.Get(1)
_ = B6107
__e.Return(MakeNative(func(__e *ControlFlow) {
L6108 := __e.Get(1)
_ = L6108
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6109 := __e.Get(1)
_ = Key6109
__e.Return(MakeNative(func(__e *ControlFlow) {
C6110 := __e.Get(1)
_ = C6110
tmp18286 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp18287 := PrimCons(symboolean, Nil)

tmp18288 := PrimCons(sym_1_1_6, tmp18287)

tmp18289 := PrimCons(A, tmp18288)

tmp18290 := Call(__e, PrimFunc(symis_b), V6111, tmp18289, B6107, L6108, Key6109, C6110)


__e.TailApply(PrimFunc(symshen_4gc), B6107, tmp18290)
return


}, 1)

tmp18291 := Call(__e, PrimFunc(symshen_4newpv), B6107)


__e.TailApply(tmp18286, tmp18291)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18292 := PrimValue(symshen_4_dsigf_d)

tmp18293 := Call(__e, PrimFunc(symshen_4assoc_1_6), symempty_2, tmp18285, tmp18292)


tmp18294 := PrimSet(symshen_4_dsigf_d, tmp18293)

_ = tmp18294

tmp18295 := MakeNative(func(__e *ControlFlow) {
V6116 := __e.Get(1)
_ = V6116
__e.Return(MakeNative(func(__e *ControlFlow) {
B6112 := __e.Get(1)
_ = B6112
__e.Return(MakeNative(func(__e *ControlFlow) {
L6113 := __e.Get(1)
_ = L6113
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6114 := __e.Get(1)
_ = Key6114
__e.Return(MakeNative(func(__e *ControlFlow) {
C6115 := __e.Get(1)
_ = C6115
tmp18296 := PrimCons(symboolean, Nil)

tmp18297 := PrimCons(sym_1_1_6, tmp18296)

tmp18298 := PrimCons(symsymbol, tmp18297)

__e.TailApply(PrimFunc(symis_b), V6116, tmp18298, B6112, L6113, Key6114, C6115)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18299 := PrimValue(symshen_4_dsigf_d)

tmp18300 := Call(__e, PrimFunc(symshen_4assoc_1_6), symenable_1type_1theory, tmp18295, tmp18299)


tmp18301 := PrimSet(symshen_4_dsigf_d, tmp18300)

_ = tmp18301

tmp18302 := MakeNative(func(__e *ControlFlow) {
V6121 := __e.Get(1)
_ = V6121
__e.Return(MakeNative(func(__e *ControlFlow) {
B6117 := __e.Get(1)
_ = B6117
__e.Return(MakeNative(func(__e *ControlFlow) {
L6118 := __e.Get(1)
_ = L6118
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6119 := __e.Get(1)
_ = Key6119
__e.Return(MakeNative(func(__e *ControlFlow) {
C6120 := __e.Get(1)
_ = C6120
tmp18303 := PrimCons(symsymbol, Nil)

tmp18304 := PrimCons(symlist, tmp18303)

tmp18305 := PrimCons(tmp18304, Nil)

tmp18306 := PrimCons(sym_1_1_6, tmp18305)

tmp18307 := PrimCons(symsymbol, tmp18306)

__e.TailApply(PrimFunc(symis_b), V6121, tmp18307, B6117, L6118, Key6119, C6120)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18308 := PrimValue(symshen_4_dsigf_d)

tmp18309 := Call(__e, PrimFunc(symshen_4assoc_1_6), symexternal, tmp18302, tmp18308)


tmp18310 := PrimSet(symshen_4_dsigf_d, tmp18309)

_ = tmp18310

tmp18311 := MakeNative(func(__e *ControlFlow) {
V6126 := __e.Get(1)
_ = V6126
__e.Return(MakeNative(func(__e *ControlFlow) {
B6122 := __e.Get(1)
_ = B6122
__e.Return(MakeNative(func(__e *ControlFlow) {
L6123 := __e.Get(1)
_ = L6123
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6124 := __e.Get(1)
_ = Key6124
__e.Return(MakeNative(func(__e *ControlFlow) {
C6125 := __e.Get(1)
_ = C6125
tmp18312 := PrimCons(symstring, Nil)

tmp18313 := PrimCons(sym_1_1_6, tmp18312)

tmp18314 := PrimCons(symexception, tmp18313)

__e.TailApply(PrimFunc(symis_b), V6126, tmp18314, B6122, L6123, Key6124, C6125)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18315 := PrimValue(symshen_4_dsigf_d)

tmp18316 := Call(__e, PrimFunc(symshen_4assoc_1_6), symerror_1to_1string, tmp18311, tmp18315)


tmp18317 := PrimSet(symshen_4_dsigf_d, tmp18316)

_ = tmp18317

tmp18318 := MakeNative(func(__e *ControlFlow) {
V6131 := __e.Get(1)
_ = V6131
__e.Return(MakeNative(func(__e *ControlFlow) {
B6127 := __e.Get(1)
_ = B6127
__e.Return(MakeNative(func(__e *ControlFlow) {
L6128 := __e.Get(1)
_ = L6128
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6129 := __e.Get(1)
_ = Key6129
__e.Return(MakeNative(func(__e *ControlFlow) {
C6130 := __e.Get(1)
_ = C6130
tmp18319 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp18320 := PrimCons(symstring, Nil)

tmp18321 := PrimCons(symlist, tmp18320)

tmp18322 := PrimCons(tmp18321, Nil)

tmp18323 := PrimCons(sym_1_1_6, tmp18322)

tmp18324 := PrimCons(A, tmp18323)

tmp18325 := Call(__e, PrimFunc(symis_b), V6131, tmp18324, B6127, L6128, Key6129, C6130)


__e.TailApply(PrimFunc(symshen_4gc), B6127, tmp18325)
return


}, 1)

tmp18326 := Call(__e, PrimFunc(symshen_4newpv), B6127)


__e.TailApply(tmp18319, tmp18326)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18327 := PrimValue(symshen_4_dsigf_d)

tmp18328 := Call(__e, PrimFunc(symshen_4assoc_1_6), symexplode, tmp18318, tmp18327)


tmp18329 := PrimSet(symshen_4_dsigf_d, tmp18328)

_ = tmp18329

tmp18330 := MakeNative(func(__e *ControlFlow) {
V6136 := __e.Get(1)
_ = V6136
__e.Return(MakeNative(func(__e *ControlFlow) {
B6132 := __e.Get(1)
_ = B6132
__e.Return(MakeNative(func(__e *ControlFlow) {
L6133 := __e.Get(1)
_ = L6133
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6134 := __e.Get(1)
_ = Key6134
__e.Return(MakeNative(func(__e *ControlFlow) {
C6135 := __e.Get(1)
_ = C6135
tmp18331 := PrimCons(symsymbol, Nil)

tmp18332 := PrimCons(sym_1_1_6, tmp18331)

tmp18333 := PrimCons(symsymbol, tmp18332)

__e.TailApply(PrimFunc(symis_b), V6136, tmp18333, B6132, L6133, Key6134, C6135)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18334 := PrimValue(symshen_4_dsigf_d)

tmp18335 := Call(__e, PrimFunc(symshen_4assoc_1_6), symfactorise, tmp18330, tmp18334)


tmp18336 := PrimSet(symshen_4_dsigf_d, tmp18335)

_ = tmp18336

tmp18337 := MakeNative(func(__e *ControlFlow) {
V6141 := __e.Get(1)
_ = V6141
__e.Return(MakeNative(func(__e *ControlFlow) {
B6137 := __e.Get(1)
_ = B6137
__e.Return(MakeNative(func(__e *ControlFlow) {
L6138 := __e.Get(1)
_ = L6138
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6139 := __e.Get(1)
_ = Key6139
__e.Return(MakeNative(func(__e *ControlFlow) {
C6140 := __e.Get(1)
_ = C6140
tmp18338 := PrimCons(symboolean, Nil)

tmp18339 := PrimCons(sym_1_1_6, tmp18338)

__e.TailApply(PrimFunc(symis_b), V6141, tmp18339, B6137, L6138, Key6139, C6140)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18340 := PrimValue(symshen_4_dsigf_d)

tmp18341 := Call(__e, PrimFunc(symshen_4assoc_1_6), symfactorise_2, tmp18337, tmp18340)


tmp18342 := PrimSet(symshen_4_dsigf_d, tmp18341)

_ = tmp18342

tmp18343 := MakeNative(func(__e *ControlFlow) {
V6146 := __e.Get(1)
_ = V6146
__e.Return(MakeNative(func(__e *ControlFlow) {
B6142 := __e.Get(1)
_ = B6142
__e.Return(MakeNative(func(__e *ControlFlow) {
L6143 := __e.Get(1)
_ = L6143
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6144 := __e.Get(1)
_ = Key6144
__e.Return(MakeNative(func(__e *ControlFlow) {
C6145 := __e.Get(1)
_ = C6145
tmp18344 := PrimCons(symsymbol, Nil)

tmp18345 := PrimCons(sym_1_1_6, tmp18344)

__e.TailApply(PrimFunc(symis_b), V6146, tmp18345, B6142, L6143, Key6144, C6145)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18346 := PrimValue(symshen_4_dsigf_d)

tmp18347 := Call(__e, PrimFunc(symshen_4assoc_1_6), symfail, tmp18343, tmp18346)


tmp18348 := PrimSet(symshen_4_dsigf_d, tmp18347)

_ = tmp18348

tmp18349 := MakeNative(func(__e *ControlFlow) {
V6151 := __e.Get(1)
_ = V6151
__e.Return(MakeNative(func(__e *ControlFlow) {
B6147 := __e.Get(1)
_ = B6147
__e.Return(MakeNative(func(__e *ControlFlow) {
L6148 := __e.Get(1)
_ = L6148
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6149 := __e.Get(1)
_ = Key6149
__e.Return(MakeNative(func(__e *ControlFlow) {
C6150 := __e.Get(1)
_ = C6150
tmp18350 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp18351 := PrimCons(A, Nil)

tmp18352 := PrimCons(sym_1_1_6, tmp18351)

tmp18353 := PrimCons(A, tmp18352)

tmp18354 := PrimCons(A, Nil)

tmp18355 := PrimCons(sym_1_1_6, tmp18354)

tmp18356 := PrimCons(A, tmp18355)

tmp18357 := PrimCons(tmp18356, Nil)

tmp18358 := PrimCons(sym_1_1_6, tmp18357)

tmp18359 := PrimCons(tmp18353, tmp18358)

tmp18360 := Call(__e, PrimFunc(symis_b), V6151, tmp18359, B6147, L6148, Key6149, C6150)


__e.TailApply(PrimFunc(symshen_4gc), B6147, tmp18360)
return


}, 1)

tmp18361 := Call(__e, PrimFunc(symshen_4newpv), B6147)


__e.TailApply(tmp18350, tmp18361)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18362 := PrimValue(symshen_4_dsigf_d)

tmp18363 := Call(__e, PrimFunc(symshen_4assoc_1_6), symfix, tmp18349, tmp18362)


tmp18364 := PrimSet(symshen_4_dsigf_d, tmp18363)

_ = tmp18364

tmp18365 := MakeNative(func(__e *ControlFlow) {
V6156 := __e.Get(1)
_ = V6156
__e.Return(MakeNative(func(__e *ControlFlow) {
B6152 := __e.Get(1)
_ = B6152
__e.Return(MakeNative(func(__e *ControlFlow) {
L6153 := __e.Get(1)
_ = L6153
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6154 := __e.Get(1)
_ = Key6154
__e.Return(MakeNative(func(__e *ControlFlow) {
C6155 := __e.Get(1)
_ = C6155
tmp18366 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp18367 := PrimCons(A, Nil)

tmp18368 := PrimCons(symlazy, tmp18367)

tmp18369 := PrimCons(tmp18368, Nil)

tmp18370 := PrimCons(sym_1_1_6, tmp18369)

tmp18371 := PrimCons(A, tmp18370)

tmp18372 := Call(__e, PrimFunc(symis_b), V6156, tmp18371, B6152, L6153, Key6154, C6155)


__e.TailApply(PrimFunc(symshen_4gc), B6152, tmp18372)
return


}, 1)

tmp18373 := Call(__e, PrimFunc(symshen_4newpv), B6152)


__e.TailApply(tmp18366, tmp18373)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18374 := PrimValue(symshen_4_dsigf_d)

tmp18375 := Call(__e, PrimFunc(symshen_4assoc_1_6), symfreeze, tmp18365, tmp18374)


tmp18376 := PrimSet(symshen_4_dsigf_d, tmp18375)

_ = tmp18376

tmp18377 := MakeNative(func(__e *ControlFlow) {
V6161 := __e.Get(1)
_ = V6161
__e.Return(MakeNative(func(__e *ControlFlow) {
B6157 := __e.Get(1)
_ = B6157
__e.Return(MakeNative(func(__e *ControlFlow) {
L6158 := __e.Get(1)
_ = L6158
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6159 := __e.Get(1)
_ = Key6159
__e.Return(MakeNative(func(__e *ControlFlow) {
C6160 := __e.Get(1)
_ = C6160
tmp18378 := MakeNative(func(__e *ControlFlow) {
B := __e.Get(1)
_ = B
tmp18379 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp18380 := PrimCons(B, Nil)

tmp18381 := PrimCons(sym_d, tmp18380)

tmp18382 := PrimCons(A, tmp18381)

tmp18383 := PrimCons(A, Nil)

tmp18384 := PrimCons(sym_1_1_6, tmp18383)

tmp18385 := PrimCons(tmp18382, tmp18384)

tmp18386 := Call(__e, PrimFunc(symis_b), V6161, tmp18385, B6157, L6158, Key6159, C6160)


__e.TailApply(PrimFunc(symshen_4gc), B6157, tmp18386)
return


}, 1)

tmp18387 := Call(__e, PrimFunc(symshen_4newpv), B6157)


tmp18388 := Call(__e, tmp18379, tmp18387)


__e.TailApply(PrimFunc(symshen_4gc), B6157, tmp18388)
return


}, 1)

tmp18389 := Call(__e, PrimFunc(symshen_4newpv), B6157)


__e.TailApply(tmp18378, tmp18389)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18390 := PrimValue(symshen_4_dsigf_d)

tmp18391 := Call(__e, PrimFunc(symshen_4assoc_1_6), symfst, tmp18377, tmp18390)


tmp18392 := PrimSet(symshen_4_dsigf_d, tmp18391)

_ = tmp18392

tmp18393 := MakeNative(func(__e *ControlFlow) {
V6166 := __e.Get(1)
_ = V6166
__e.Return(MakeNative(func(__e *ControlFlow) {
B6162 := __e.Get(1)
_ = B6162
__e.Return(MakeNative(func(__e *ControlFlow) {
L6163 := __e.Get(1)
_ = L6163
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6164 := __e.Get(1)
_ = Key6164
__e.Return(MakeNative(func(__e *ControlFlow) {
C6165 := __e.Get(1)
_ = C6165
tmp18394 := PrimCons(symsymbol, Nil)

tmp18395 := PrimCons(sym_1_1_6, tmp18394)

tmp18396 := PrimCons(symsymbol, tmp18395)

__e.TailApply(PrimFunc(symis_b), V6166, tmp18396, B6162, L6163, Key6164, C6165)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18397 := PrimValue(symshen_4_dsigf_d)

tmp18398 := Call(__e, PrimFunc(symshen_4assoc_1_6), symgensym, tmp18393, tmp18397)


tmp18399 := PrimSet(symshen_4_dsigf_d, tmp18398)

_ = tmp18399

tmp18400 := MakeNative(func(__e *ControlFlow) {
V6171 := __e.Get(1)
_ = V6171
__e.Return(MakeNative(func(__e *ControlFlow) {
B6167 := __e.Get(1)
_ = B6167
__e.Return(MakeNative(func(__e *ControlFlow) {
L6168 := __e.Get(1)
_ = L6168
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6169 := __e.Get(1)
_ = Key6169
__e.Return(MakeNative(func(__e *ControlFlow) {
C6170 := __e.Get(1)
_ = C6170
tmp18401 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp18402 := PrimCons(A, Nil)

tmp18403 := PrimCons(symlist, tmp18402)

tmp18404 := PrimCons(symboolean, Nil)

tmp18405 := PrimCons(sym_1_1_6, tmp18404)

tmp18406 := PrimCons(A, tmp18405)

tmp18407 := PrimCons(tmp18406, Nil)

tmp18408 := PrimCons(sym_1_1_6, tmp18407)

tmp18409 := PrimCons(tmp18403, tmp18408)

tmp18410 := Call(__e, PrimFunc(symis_b), V6171, tmp18409, B6167, L6168, Key6169, C6170)


__e.TailApply(PrimFunc(symshen_4gc), B6167, tmp18410)
return


}, 1)

tmp18411 := Call(__e, PrimFunc(symshen_4newpv), B6167)


__e.TailApply(tmp18401, tmp18411)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18412 := PrimValue(symshen_4_dsigf_d)

tmp18413 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4hds_a_2, tmp18400, tmp18412)


tmp18414 := PrimSet(symshen_4_dsigf_d, tmp18413)

_ = tmp18414

tmp18415 := MakeNative(func(__e *ControlFlow) {
V6176 := __e.Get(1)
_ = V6176
__e.Return(MakeNative(func(__e *ControlFlow) {
B6172 := __e.Get(1)
_ = B6172
__e.Return(MakeNative(func(__e *ControlFlow) {
L6173 := __e.Get(1)
_ = L6173
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6174 := __e.Get(1)
_ = Key6174
__e.Return(MakeNative(func(__e *ControlFlow) {
C6175 := __e.Get(1)
_ = C6175
tmp18416 := PrimCons(symboolean, Nil)

tmp18417 := PrimCons(sym_1_1_6, tmp18416)

tmp18418 := PrimCons(symsymbol, tmp18417)

__e.TailApply(PrimFunc(symis_b), V6176, tmp18418, B6172, L6173, Key6174, C6175)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18419 := PrimValue(symshen_4_dsigf_d)

tmp18420 := Call(__e, PrimFunc(symshen_4assoc_1_6), symhush, tmp18415, tmp18419)


tmp18421 := PrimSet(symshen_4_dsigf_d, tmp18420)

_ = tmp18421

tmp18422 := MakeNative(func(__e *ControlFlow) {
V6181 := __e.Get(1)
_ = V6181
__e.Return(MakeNative(func(__e *ControlFlow) {
B6177 := __e.Get(1)
_ = B6177
__e.Return(MakeNative(func(__e *ControlFlow) {
L6178 := __e.Get(1)
_ = L6178
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6179 := __e.Get(1)
_ = Key6179
__e.Return(MakeNative(func(__e *ControlFlow) {
C6180 := __e.Get(1)
_ = C6180
tmp18423 := PrimCons(symboolean, Nil)

tmp18424 := PrimCons(sym_1_1_6, tmp18423)

__e.TailApply(PrimFunc(symis_b), V6181, tmp18424, B6177, L6178, Key6179, C6180)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18425 := PrimValue(symshen_4_dsigf_d)

tmp18426 := Call(__e, PrimFunc(symshen_4assoc_1_6), symhush_2, tmp18422, tmp18425)


tmp18427 := PrimSet(symshen_4_dsigf_d, tmp18426)

_ = tmp18427

tmp18428 := MakeNative(func(__e *ControlFlow) {
V6186 := __e.Get(1)
_ = V6186
__e.Return(MakeNative(func(__e *ControlFlow) {
B6182 := __e.Get(1)
_ = B6182
__e.Return(MakeNative(func(__e *ControlFlow) {
L6183 := __e.Get(1)
_ = L6183
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6184 := __e.Get(1)
_ = Key6184
__e.Return(MakeNative(func(__e *ControlFlow) {
C6185 := __e.Get(1)
_ = C6185
tmp18429 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp18430 := PrimCons(A, Nil)

tmp18431 := PrimCons(symvector, tmp18430)

tmp18432 := PrimCons(A, Nil)

tmp18433 := PrimCons(sym_1_1_6, tmp18432)

tmp18434 := PrimCons(symnumber, tmp18433)

tmp18435 := PrimCons(tmp18434, Nil)

tmp18436 := PrimCons(sym_1_1_6, tmp18435)

tmp18437 := PrimCons(tmp18431, tmp18436)

tmp18438 := Call(__e, PrimFunc(symis_b), V6186, tmp18437, B6182, L6183, Key6184, C6185)


__e.TailApply(PrimFunc(symshen_4gc), B6182, tmp18438)
return


}, 1)

tmp18439 := Call(__e, PrimFunc(symshen_4newpv), B6182)


__e.TailApply(tmp18429, tmp18439)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18440 := PrimValue(symshen_4_dsigf_d)

tmp18441 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_5_1vector, tmp18428, tmp18440)


tmp18442 := PrimSet(symshen_4_dsigf_d, tmp18441)

_ = tmp18442

tmp18443 := MakeNative(func(__e *ControlFlow) {
V6191 := __e.Get(1)
_ = V6191
__e.Return(MakeNative(func(__e *ControlFlow) {
B6187 := __e.Get(1)
_ = B6187
__e.Return(MakeNative(func(__e *ControlFlow) {
L6188 := __e.Get(1)
_ = L6188
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6189 := __e.Get(1)
_ = Key6189
__e.Return(MakeNative(func(__e *ControlFlow) {
C6190 := __e.Get(1)
_ = C6190
tmp18444 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp18445 := PrimCons(A, Nil)

tmp18446 := PrimCons(symvector, tmp18445)

tmp18447 := PrimCons(A, Nil)

tmp18448 := PrimCons(symvector, tmp18447)

tmp18449 := PrimCons(tmp18448, Nil)

tmp18450 := PrimCons(sym_1_1_6, tmp18449)

tmp18451 := PrimCons(A, tmp18450)

tmp18452 := PrimCons(tmp18451, Nil)

tmp18453 := PrimCons(sym_1_1_6, tmp18452)

tmp18454 := PrimCons(symnumber, tmp18453)

tmp18455 := PrimCons(tmp18454, Nil)

tmp18456 := PrimCons(sym_1_1_6, tmp18455)

tmp18457 := PrimCons(tmp18446, tmp18456)

tmp18458 := Call(__e, PrimFunc(symis_b), V6191, tmp18457, B6187, L6188, Key6189, C6190)


__e.TailApply(PrimFunc(symshen_4gc), B6187, tmp18458)
return


}, 1)

tmp18459 := Call(__e, PrimFunc(symshen_4newpv), B6187)


__e.TailApply(tmp18444, tmp18459)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18460 := PrimValue(symshen_4_dsigf_d)

tmp18461 := Call(__e, PrimFunc(symshen_4assoc_1_6), symvector_1_6, tmp18443, tmp18460)


tmp18462 := PrimSet(symshen_4_dsigf_d, tmp18461)

_ = tmp18462

tmp18463 := MakeNative(func(__e *ControlFlow) {
V6196 := __e.Get(1)
_ = V6196
__e.Return(MakeNative(func(__e *ControlFlow) {
B6192 := __e.Get(1)
_ = B6192
__e.Return(MakeNative(func(__e *ControlFlow) {
L6193 := __e.Get(1)
_ = L6193
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6194 := __e.Get(1)
_ = Key6194
__e.Return(MakeNative(func(__e *ControlFlow) {
C6195 := __e.Get(1)
_ = C6195
tmp18464 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp18465 := PrimCons(A, Nil)

tmp18466 := PrimCons(symvector, tmp18465)

tmp18467 := PrimCons(tmp18466, Nil)

tmp18468 := PrimCons(sym_1_1_6, tmp18467)

tmp18469 := PrimCons(symnumber, tmp18468)

tmp18470 := Call(__e, PrimFunc(symis_b), V6196, tmp18469, B6192, L6193, Key6194, C6195)


__e.TailApply(PrimFunc(symshen_4gc), B6192, tmp18470)
return


}, 1)

tmp18471 := Call(__e, PrimFunc(symshen_4newpv), B6192)


__e.TailApply(tmp18464, tmp18471)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18472 := PrimValue(symshen_4_dsigf_d)

tmp18473 := Call(__e, PrimFunc(symshen_4assoc_1_6), symvector, tmp18463, tmp18472)


tmp18474 := PrimSet(symshen_4_dsigf_d, tmp18473)

_ = tmp18474

tmp18475 := MakeNative(func(__e *ControlFlow) {
V6201 := __e.Get(1)
_ = V6201
__e.Return(MakeNative(func(__e *ControlFlow) {
B6197 := __e.Get(1)
_ = B6197
__e.Return(MakeNative(func(__e *ControlFlow) {
L6198 := __e.Get(1)
_ = L6198
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6199 := __e.Get(1)
_ = Key6199
__e.Return(MakeNative(func(__e *ControlFlow) {
C6200 := __e.Get(1)
_ = C6200
tmp18476 := PrimCons(symnumber, Nil)

tmp18477 := PrimCons(sym_1_1_6, tmp18476)

tmp18478 := PrimCons(symsymbol, tmp18477)

__e.TailApply(PrimFunc(symis_b), V6201, tmp18478, B6197, L6198, Key6199, C6200)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18479 := PrimValue(symshen_4_dsigf_d)

tmp18480 := Call(__e, PrimFunc(symshen_4assoc_1_6), symget_1time, tmp18475, tmp18479)


tmp18481 := PrimSet(symshen_4_dsigf_d, tmp18480)

_ = tmp18481

tmp18482 := MakeNative(func(__e *ControlFlow) {
V6206 := __e.Get(1)
_ = V6206
__e.Return(MakeNative(func(__e *ControlFlow) {
B6202 := __e.Get(1)
_ = B6202
__e.Return(MakeNative(func(__e *ControlFlow) {
L6203 := __e.Get(1)
_ = L6203
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6204 := __e.Get(1)
_ = Key6204
__e.Return(MakeNative(func(__e *ControlFlow) {
C6205 := __e.Get(1)
_ = C6205
tmp18483 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp18484 := PrimCons(symnumber, Nil)

tmp18485 := PrimCons(sym_1_1_6, tmp18484)

tmp18486 := PrimCons(symnumber, tmp18485)

tmp18487 := PrimCons(tmp18486, Nil)

tmp18488 := PrimCons(sym_1_1_6, tmp18487)

tmp18489 := PrimCons(A, tmp18488)

tmp18490 := Call(__e, PrimFunc(symis_b), V6206, tmp18489, B6202, L6203, Key6204, C6205)


__e.TailApply(PrimFunc(symshen_4gc), B6202, tmp18490)
return


}, 1)

tmp18491 := Call(__e, PrimFunc(symshen_4newpv), B6202)


__e.TailApply(tmp18483, tmp18491)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18492 := PrimValue(symshen_4_dsigf_d)

tmp18493 := Call(__e, PrimFunc(symshen_4assoc_1_6), symhash, tmp18482, tmp18492)


tmp18494 := PrimSet(symshen_4_dsigf_d, tmp18493)

_ = tmp18494

tmp18495 := MakeNative(func(__e *ControlFlow) {
V6211 := __e.Get(1)
_ = V6211
__e.Return(MakeNative(func(__e *ControlFlow) {
B6207 := __e.Get(1)
_ = B6207
__e.Return(MakeNative(func(__e *ControlFlow) {
L6208 := __e.Get(1)
_ = L6208
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6209 := __e.Get(1)
_ = Key6209
__e.Return(MakeNative(func(__e *ControlFlow) {
C6210 := __e.Get(1)
_ = C6210
tmp18496 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp18497 := PrimCons(A, Nil)

tmp18498 := PrimCons(symlist, tmp18497)

tmp18499 := PrimCons(A, Nil)

tmp18500 := PrimCons(sym_1_1_6, tmp18499)

tmp18501 := PrimCons(tmp18498, tmp18500)

tmp18502 := Call(__e, PrimFunc(symis_b), V6211, tmp18501, B6207, L6208, Key6209, C6210)


__e.TailApply(PrimFunc(symshen_4gc), B6207, tmp18502)
return


}, 1)

tmp18503 := Call(__e, PrimFunc(symshen_4newpv), B6207)


__e.TailApply(tmp18496, tmp18503)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18504 := PrimValue(symshen_4_dsigf_d)

tmp18505 := Call(__e, PrimFunc(symshen_4assoc_1_6), symhead, tmp18495, tmp18504)


tmp18506 := PrimSet(symshen_4_dsigf_d, tmp18505)

_ = tmp18506

tmp18507 := MakeNative(func(__e *ControlFlow) {
V6216 := __e.Get(1)
_ = V6216
__e.Return(MakeNative(func(__e *ControlFlow) {
B6212 := __e.Get(1)
_ = B6212
__e.Return(MakeNative(func(__e *ControlFlow) {
L6213 := __e.Get(1)
_ = L6213
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6214 := __e.Get(1)
_ = Key6214
__e.Return(MakeNative(func(__e *ControlFlow) {
C6215 := __e.Get(1)
_ = C6215
tmp18508 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp18509 := PrimCons(A, Nil)

tmp18510 := PrimCons(symvector, tmp18509)

tmp18511 := PrimCons(A, Nil)

tmp18512 := PrimCons(sym_1_1_6, tmp18511)

tmp18513 := PrimCons(tmp18510, tmp18512)

tmp18514 := Call(__e, PrimFunc(symis_b), V6216, tmp18513, B6212, L6213, Key6214, C6215)


__e.TailApply(PrimFunc(symshen_4gc), B6212, tmp18514)
return


}, 1)

tmp18515 := Call(__e, PrimFunc(symshen_4newpv), B6212)


__e.TailApply(tmp18508, tmp18515)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18516 := PrimValue(symshen_4_dsigf_d)

tmp18517 := Call(__e, PrimFunc(symshen_4assoc_1_6), symhdv, tmp18507, tmp18516)


tmp18518 := PrimSet(symshen_4_dsigf_d, tmp18517)

_ = tmp18518

tmp18519 := MakeNative(func(__e *ControlFlow) {
V6221 := __e.Get(1)
_ = V6221
__e.Return(MakeNative(func(__e *ControlFlow) {
B6217 := __e.Get(1)
_ = B6217
__e.Return(MakeNative(func(__e *ControlFlow) {
L6218 := __e.Get(1)
_ = L6218
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6219 := __e.Get(1)
_ = Key6219
__e.Return(MakeNative(func(__e *ControlFlow) {
C6220 := __e.Get(1)
_ = C6220
tmp18520 := PrimCons(symstring, Nil)

tmp18521 := PrimCons(sym_1_1_6, tmp18520)

tmp18522 := PrimCons(symstring, tmp18521)

__e.TailApply(PrimFunc(symis_b), V6221, tmp18522, B6217, L6218, Key6219, C6220)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18523 := PrimValue(symshen_4_dsigf_d)

tmp18524 := Call(__e, PrimFunc(symshen_4assoc_1_6), symhdstr, tmp18519, tmp18523)


tmp18525 := PrimSet(symshen_4_dsigf_d, tmp18524)

_ = tmp18525

tmp18526 := MakeNative(func(__e *ControlFlow) {
V6226 := __e.Get(1)
_ = V6226
__e.Return(MakeNative(func(__e *ControlFlow) {
B6222 := __e.Get(1)
_ = B6222
__e.Return(MakeNative(func(__e *ControlFlow) {
L6223 := __e.Get(1)
_ = L6223
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6224 := __e.Get(1)
_ = Key6224
__e.Return(MakeNative(func(__e *ControlFlow) {
C6225 := __e.Get(1)
_ = C6225
tmp18527 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp18528 := PrimCons(A, Nil)

tmp18529 := PrimCons(sym_1_1_6, tmp18528)

tmp18530 := PrimCons(A, tmp18529)

tmp18531 := PrimCons(tmp18530, Nil)

tmp18532 := PrimCons(sym_1_1_6, tmp18531)

tmp18533 := PrimCons(A, tmp18532)

tmp18534 := PrimCons(tmp18533, Nil)

tmp18535 := PrimCons(sym_1_1_6, tmp18534)

tmp18536 := PrimCons(symboolean, tmp18535)

tmp18537 := Call(__e, PrimFunc(symis_b), V6226, tmp18536, B6222, L6223, Key6224, C6225)


__e.TailApply(PrimFunc(symshen_4gc), B6222, tmp18537)
return


}, 1)

tmp18538 := Call(__e, PrimFunc(symshen_4newpv), B6222)


__e.TailApply(tmp18527, tmp18538)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18539 := PrimValue(symshen_4_dsigf_d)

tmp18540 := Call(__e, PrimFunc(symshen_4assoc_1_6), symif, tmp18526, tmp18539)


tmp18541 := PrimSet(symshen_4_dsigf_d, tmp18540)

_ = tmp18541

tmp18542 := MakeNative(func(__e *ControlFlow) {
V6231 := __e.Get(1)
_ = V6231
__e.Return(MakeNative(func(__e *ControlFlow) {
B6227 := __e.Get(1)
_ = B6227
__e.Return(MakeNative(func(__e *ControlFlow) {
L6228 := __e.Get(1)
_ = L6228
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6229 := __e.Get(1)
_ = Key6229
__e.Return(MakeNative(func(__e *ControlFlow) {
C6230 := __e.Get(1)
_ = C6230
tmp18543 := PrimCons(symsymbol, Nil)

tmp18544 := PrimCons(sym_1_1_6, tmp18543)

tmp18545 := PrimCons(symsymbol, tmp18544)

__e.TailApply(PrimFunc(symis_b), V6231, tmp18545, B6227, L6228, Key6229, C6230)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18546 := PrimValue(symshen_4_dsigf_d)

tmp18547 := Call(__e, PrimFunc(symshen_4assoc_1_6), symin_1package, tmp18542, tmp18546)


tmp18548 := PrimSet(symshen_4_dsigf_d, tmp18547)

_ = tmp18548

tmp18549 := MakeNative(func(__e *ControlFlow) {
V6236 := __e.Get(1)
_ = V6236
__e.Return(MakeNative(func(__e *ControlFlow) {
B6232 := __e.Get(1)
_ = B6232
__e.Return(MakeNative(func(__e *ControlFlow) {
L6233 := __e.Get(1)
_ = L6233
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6234 := __e.Get(1)
_ = Key6234
__e.Return(MakeNative(func(__e *ControlFlow) {
C6235 := __e.Get(1)
_ = C6235
tmp18550 := PrimCons(symstring, Nil)

tmp18551 := PrimCons(sym_1_1_6, tmp18550)

__e.TailApply(PrimFunc(symis_b), V6236, tmp18551, B6232, L6233, Key6234, C6235)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18552 := PrimValue(symshen_4_dsigf_d)

tmp18553 := Call(__e, PrimFunc(symshen_4assoc_1_6), symit, tmp18549, tmp18552)


tmp18554 := PrimSet(symshen_4_dsigf_d, tmp18553)

_ = tmp18554

tmp18555 := MakeNative(func(__e *ControlFlow) {
V6241 := __e.Get(1)
_ = V6241
__e.Return(MakeNative(func(__e *ControlFlow) {
B6237 := __e.Get(1)
_ = B6237
__e.Return(MakeNative(func(__e *ControlFlow) {
L6238 := __e.Get(1)
_ = L6238
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6239 := __e.Get(1)
_ = Key6239
__e.Return(MakeNative(func(__e *ControlFlow) {
C6240 := __e.Get(1)
_ = C6240
tmp18556 := PrimCons(symstring, Nil)

tmp18557 := PrimCons(sym_1_1_6, tmp18556)

__e.TailApply(PrimFunc(symis_b), V6241, tmp18557, B6237, L6238, Key6239, C6240)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18558 := PrimValue(symshen_4_dsigf_d)

tmp18559 := Call(__e, PrimFunc(symshen_4assoc_1_6), symimplementation, tmp18555, tmp18558)


tmp18560 := PrimSet(symshen_4_dsigf_d, tmp18559)

_ = tmp18560

tmp18561 := MakeNative(func(__e *ControlFlow) {
V6246 := __e.Get(1)
_ = V6246
__e.Return(MakeNative(func(__e *ControlFlow) {
B6242 := __e.Get(1)
_ = B6242
__e.Return(MakeNative(func(__e *ControlFlow) {
L6243 := __e.Get(1)
_ = L6243
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6244 := __e.Get(1)
_ = Key6244
__e.Return(MakeNative(func(__e *ControlFlow) {
C6245 := __e.Get(1)
_ = C6245
tmp18562 := PrimCons(symsymbol, Nil)

tmp18563 := PrimCons(symlist, tmp18562)

tmp18564 := PrimCons(symsymbol, Nil)

tmp18565 := PrimCons(symlist, tmp18564)

tmp18566 := PrimCons(tmp18565, Nil)

tmp18567 := PrimCons(sym_1_1_6, tmp18566)

tmp18568 := PrimCons(tmp18563, tmp18567)

__e.TailApply(PrimFunc(symis_b), V6246, tmp18568, B6242, L6243, Key6244, C6245)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18569 := PrimValue(symshen_4_dsigf_d)

tmp18570 := Call(__e, PrimFunc(symshen_4assoc_1_6), syminclude, tmp18561, tmp18569)


tmp18571 := PrimSet(symshen_4_dsigf_d, tmp18570)

_ = tmp18571

tmp18572 := MakeNative(func(__e *ControlFlow) {
V6251 := __e.Get(1)
_ = V6251
__e.Return(MakeNative(func(__e *ControlFlow) {
B6247 := __e.Get(1)
_ = B6247
__e.Return(MakeNative(func(__e *ControlFlow) {
L6248 := __e.Get(1)
_ = L6248
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6249 := __e.Get(1)
_ = Key6249
__e.Return(MakeNative(func(__e *ControlFlow) {
C6250 := __e.Get(1)
_ = C6250
tmp18573 := PrimCons(symsymbol, Nil)

tmp18574 := PrimCons(symlist, tmp18573)

tmp18575 := PrimCons(symsymbol, Nil)

tmp18576 := PrimCons(symlist, tmp18575)

tmp18577 := PrimCons(tmp18576, Nil)

tmp18578 := PrimCons(sym_1_1_6, tmp18577)

tmp18579 := PrimCons(tmp18574, tmp18578)

__e.TailApply(PrimFunc(symis_b), V6251, tmp18579, B6247, L6248, Key6249, C6250)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18580 := PrimValue(symshen_4_dsigf_d)

tmp18581 := Call(__e, PrimFunc(symshen_4assoc_1_6), syminclude_1all_1but, tmp18572, tmp18580)


tmp18582 := PrimSet(symshen_4_dsigf_d, tmp18581)

_ = tmp18582

tmp18583 := MakeNative(func(__e *ControlFlow) {
V6256 := __e.Get(1)
_ = V6256
__e.Return(MakeNative(func(__e *ControlFlow) {
B6252 := __e.Get(1)
_ = B6252
__e.Return(MakeNative(func(__e *ControlFlow) {
L6253 := __e.Get(1)
_ = L6253
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6254 := __e.Get(1)
_ = Key6254
__e.Return(MakeNative(func(__e *ControlFlow) {
C6255 := __e.Get(1)
_ = C6255
tmp18584 := PrimCons(symsymbol, Nil)

tmp18585 := PrimCons(symlist, tmp18584)

tmp18586 := PrimCons(tmp18585, Nil)

tmp18587 := PrimCons(sym_1_1_6, tmp18586)

__e.TailApply(PrimFunc(symis_b), V6256, tmp18587, B6252, L6253, Key6254, C6255)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18588 := PrimValue(symshen_4_dsigf_d)

tmp18589 := Call(__e, PrimFunc(symshen_4assoc_1_6), symincluded, tmp18583, tmp18588)


tmp18590 := PrimSet(symshen_4_dsigf_d, tmp18589)

_ = tmp18590

tmp18591 := MakeNative(func(__e *ControlFlow) {
V6261 := __e.Get(1)
_ = V6261
__e.Return(MakeNative(func(__e *ControlFlow) {
B6257 := __e.Get(1)
_ = B6257
__e.Return(MakeNative(func(__e *ControlFlow) {
L6258 := __e.Get(1)
_ = L6258
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6259 := __e.Get(1)
_ = Key6259
__e.Return(MakeNative(func(__e *ControlFlow) {
C6260 := __e.Get(1)
_ = C6260
tmp18592 := PrimCons(symnumber, Nil)

tmp18593 := PrimCons(sym_1_1_6, tmp18592)

__e.TailApply(PrimFunc(symis_b), V6261, tmp18593, B6257, L6258, Key6259, C6260)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18594 := PrimValue(symshen_4_dsigf_d)

tmp18595 := Call(__e, PrimFunc(symshen_4assoc_1_6), syminferences, tmp18591, tmp18594)


tmp18596 := PrimSet(symshen_4_dsigf_d, tmp18595)

_ = tmp18596

tmp18597 := MakeNative(func(__e *ControlFlow) {
V6266 := __e.Get(1)
_ = V6266
__e.Return(MakeNative(func(__e *ControlFlow) {
B6262 := __e.Get(1)
_ = B6262
__e.Return(MakeNative(func(__e *ControlFlow) {
L6263 := __e.Get(1)
_ = L6263
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6264 := __e.Get(1)
_ = Key6264
__e.Return(MakeNative(func(__e *ControlFlow) {
C6265 := __e.Get(1)
_ = C6265
tmp18598 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp18599 := PrimCons(symstring, Nil)

tmp18600 := PrimCons(sym_1_1_6, tmp18599)

tmp18601 := PrimCons(symstring, tmp18600)

tmp18602 := PrimCons(tmp18601, Nil)

tmp18603 := PrimCons(sym_1_1_6, tmp18602)

tmp18604 := PrimCons(A, tmp18603)

tmp18605 := Call(__e, PrimFunc(symis_b), V6266, tmp18604, B6262, L6263, Key6264, C6265)


__e.TailApply(PrimFunc(symshen_4gc), B6262, tmp18605)
return


}, 1)

tmp18606 := Call(__e, PrimFunc(symshen_4newpv), B6262)


__e.TailApply(tmp18598, tmp18606)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18607 := PrimValue(symshen_4_dsigf_d)

tmp18608 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4insert, tmp18597, tmp18607)


tmp18609 := PrimSet(symshen_4_dsigf_d, tmp18608)

_ = tmp18609

tmp18610 := MakeNative(func(__e *ControlFlow) {
V6271 := __e.Get(1)
_ = V6271
__e.Return(MakeNative(func(__e *ControlFlow) {
B6267 := __e.Get(1)
_ = B6267
__e.Return(MakeNative(func(__e *ControlFlow) {
L6268 := __e.Get(1)
_ = L6268
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6269 := __e.Get(1)
_ = Key6269
__e.Return(MakeNative(func(__e *ControlFlow) {
C6270 := __e.Get(1)
_ = C6270
tmp18611 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp18612 := PrimCons(symboolean, Nil)

tmp18613 := PrimCons(sym_1_1_6, tmp18612)

tmp18614 := PrimCons(A, tmp18613)

tmp18615 := Call(__e, PrimFunc(symis_b), V6271, tmp18614, B6267, L6268, Key6269, C6270)


__e.TailApply(PrimFunc(symshen_4gc), B6267, tmp18615)
return


}, 1)

tmp18616 := Call(__e, PrimFunc(symshen_4newpv), B6267)


__e.TailApply(tmp18611, tmp18616)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18617 := PrimValue(symshen_4_dsigf_d)

tmp18618 := Call(__e, PrimFunc(symshen_4assoc_1_6), syminteger_2, tmp18610, tmp18617)


tmp18619 := PrimSet(symshen_4_dsigf_d, tmp18618)

_ = tmp18619

tmp18620 := MakeNative(func(__e *ControlFlow) {
V6276 := __e.Get(1)
_ = V6276
__e.Return(MakeNative(func(__e *ControlFlow) {
B6272 := __e.Get(1)
_ = B6272
__e.Return(MakeNative(func(__e *ControlFlow) {
L6273 := __e.Get(1)
_ = L6273
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6274 := __e.Get(1)
_ = Key6274
__e.Return(MakeNative(func(__e *ControlFlow) {
C6275 := __e.Get(1)
_ = C6275
tmp18621 := PrimCons(symsymbol, Nil)

tmp18622 := PrimCons(symlist, tmp18621)

tmp18623 := PrimCons(tmp18622, Nil)

tmp18624 := PrimCons(sym_1_1_6, tmp18623)

tmp18625 := PrimCons(symsymbol, tmp18624)

__e.TailApply(PrimFunc(symis_b), V6276, tmp18625, B6272, L6273, Key6274, C6275)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18626 := PrimValue(symshen_4_dsigf_d)

tmp18627 := Call(__e, PrimFunc(symshen_4assoc_1_6), syminternal, tmp18620, tmp18626)


tmp18628 := PrimSet(symshen_4_dsigf_d, tmp18627)

_ = tmp18628

tmp18629 := MakeNative(func(__e *ControlFlow) {
V6281 := __e.Get(1)
_ = V6281
__e.Return(MakeNative(func(__e *ControlFlow) {
B6277 := __e.Get(1)
_ = B6277
__e.Return(MakeNative(func(__e *ControlFlow) {
L6278 := __e.Get(1)
_ = L6278
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6279 := __e.Get(1)
_ = Key6279
__e.Return(MakeNative(func(__e *ControlFlow) {
C6280 := __e.Get(1)
_ = C6280
tmp18630 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp18631 := PrimCons(A, Nil)

tmp18632 := PrimCons(symlist, tmp18631)

tmp18633 := PrimCons(A, Nil)

tmp18634 := PrimCons(symlist, tmp18633)

tmp18635 := PrimCons(A, Nil)

tmp18636 := PrimCons(symlist, tmp18635)

tmp18637 := PrimCons(tmp18636, Nil)

tmp18638 := PrimCons(sym_1_1_6, tmp18637)

tmp18639 := PrimCons(tmp18634, tmp18638)

tmp18640 := PrimCons(tmp18639, Nil)

tmp18641 := PrimCons(sym_1_1_6, tmp18640)

tmp18642 := PrimCons(tmp18632, tmp18641)

tmp18643 := Call(__e, PrimFunc(symis_b), V6281, tmp18642, B6277, L6278, Key6279, C6280)


__e.TailApply(PrimFunc(symshen_4gc), B6277, tmp18643)
return


}, 1)

tmp18644 := Call(__e, PrimFunc(symshen_4newpv), B6277)


__e.TailApply(tmp18630, tmp18644)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18645 := PrimValue(symshen_4_dsigf_d)

tmp18646 := Call(__e, PrimFunc(symshen_4assoc_1_6), symintersection, tmp18629, tmp18645)


tmp18647 := PrimSet(symshen_4_dsigf_d, tmp18646)

_ = tmp18647

tmp18648 := MakeNative(func(__e *ControlFlow) {
V6286 := __e.Get(1)
_ = V6286
__e.Return(MakeNative(func(__e *ControlFlow) {
B6282 := __e.Get(1)
_ = B6282
__e.Return(MakeNative(func(__e *ControlFlow) {
L6283 := __e.Get(1)
_ = L6283
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6284 := __e.Get(1)
_ = Key6284
__e.Return(MakeNative(func(__e *ControlFlow) {
C6285 := __e.Get(1)
_ = C6285
tmp18649 := PrimCons(symstring, Nil)

tmp18650 := PrimCons(sym_1_1_6, tmp18649)

__e.TailApply(PrimFunc(symis_b), V6286, tmp18650, B6282, L6283, Key6284, C6285)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18651 := PrimValue(symshen_4_dsigf_d)

tmp18652 := Call(__e, PrimFunc(symshen_4assoc_1_6), symlanguage, tmp18648, tmp18651)


tmp18653 := PrimSet(symshen_4_dsigf_d, tmp18652)

_ = tmp18653

tmp18654 := MakeNative(func(__e *ControlFlow) {
V6291 := __e.Get(1)
_ = V6291
__e.Return(MakeNative(func(__e *ControlFlow) {
B6287 := __e.Get(1)
_ = B6287
__e.Return(MakeNative(func(__e *ControlFlow) {
L6288 := __e.Get(1)
_ = L6288
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6289 := __e.Get(1)
_ = Key6289
__e.Return(MakeNative(func(__e *ControlFlow) {
C6290 := __e.Get(1)
_ = C6290
tmp18655 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp18656 := PrimCons(A, Nil)

tmp18657 := PrimCons(symlist, tmp18656)

tmp18658 := PrimCons(symnumber, Nil)

tmp18659 := PrimCons(sym_1_1_6, tmp18658)

tmp18660 := PrimCons(tmp18657, tmp18659)

tmp18661 := Call(__e, PrimFunc(symis_b), V6291, tmp18660, B6287, L6288, Key6289, C6290)


__e.TailApply(PrimFunc(symshen_4gc), B6287, tmp18661)
return


}, 1)

tmp18662 := Call(__e, PrimFunc(symshen_4newpv), B6287)


__e.TailApply(tmp18655, tmp18662)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18663 := PrimValue(symshen_4_dsigf_d)

tmp18664 := Call(__e, PrimFunc(symshen_4assoc_1_6), symlength, tmp18654, tmp18663)


tmp18665 := PrimSet(symshen_4_dsigf_d, tmp18664)

_ = tmp18665

tmp18666 := MakeNative(func(__e *ControlFlow) {
V6296 := __e.Get(1)
_ = V6296
__e.Return(MakeNative(func(__e *ControlFlow) {
B6292 := __e.Get(1)
_ = B6292
__e.Return(MakeNative(func(__e *ControlFlow) {
L6293 := __e.Get(1)
_ = L6293
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6294 := __e.Get(1)
_ = Key6294
__e.Return(MakeNative(func(__e *ControlFlow) {
C6295 := __e.Get(1)
_ = C6295
tmp18667 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp18668 := PrimCons(A, Nil)

tmp18669 := PrimCons(symvector, tmp18668)

tmp18670 := PrimCons(symnumber, Nil)

tmp18671 := PrimCons(sym_1_1_6, tmp18670)

tmp18672 := PrimCons(tmp18669, tmp18671)

tmp18673 := Call(__e, PrimFunc(symis_b), V6296, tmp18672, B6292, L6293, Key6294, C6295)


__e.TailApply(PrimFunc(symshen_4gc), B6292, tmp18673)
return


}, 1)

tmp18674 := Call(__e, PrimFunc(symshen_4newpv), B6292)


__e.TailApply(tmp18667, tmp18674)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18675 := PrimValue(symshen_4_dsigf_d)

tmp18676 := Call(__e, PrimFunc(symshen_4assoc_1_6), symlimit, tmp18666, tmp18675)


tmp18677 := PrimSet(symshen_4_dsigf_d, tmp18676)

_ = tmp18677

tmp18678 := MakeNative(func(__e *ControlFlow) {
V6301 := __e.Get(1)
_ = V6301
__e.Return(MakeNative(func(__e *ControlFlow) {
B6297 := __e.Get(1)
_ = B6297
__e.Return(MakeNative(func(__e *ControlFlow) {
L6298 := __e.Get(1)
_ = L6298
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6299 := __e.Get(1)
_ = Key6299
__e.Return(MakeNative(func(__e *ControlFlow) {
C6300 := __e.Get(1)
_ = C6300
tmp18679 := PrimCons(symin, Nil)

tmp18680 := PrimCons(symstream, tmp18679)

tmp18681 := PrimCons(symunit, Nil)

tmp18682 := PrimCons(symlist, tmp18681)

tmp18683 := PrimCons(tmp18682, Nil)

tmp18684 := PrimCons(sym_1_1_6, tmp18683)

tmp18685 := PrimCons(tmp18680, tmp18684)

__e.TailApply(PrimFunc(symis_b), V6301, tmp18685, B6297, L6298, Key6299, C6300)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18686 := PrimValue(symshen_4_dsigf_d)

tmp18687 := Call(__e, PrimFunc(symshen_4assoc_1_6), symlineread, tmp18678, tmp18686)


tmp18688 := PrimSet(symshen_4_dsigf_d, tmp18687)

_ = tmp18688

tmp18689 := MakeNative(func(__e *ControlFlow) {
V6306 := __e.Get(1)
_ = V6306
__e.Return(MakeNative(func(__e *ControlFlow) {
B6302 := __e.Get(1)
_ = B6302
__e.Return(MakeNative(func(__e *ControlFlow) {
L6303 := __e.Get(1)
_ = L6303
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6304 := __e.Get(1)
_ = Key6304
__e.Return(MakeNative(func(__e *ControlFlow) {
C6305 := __e.Get(1)
_ = C6305
tmp18690 := PrimCons(symsymbol, Nil)

tmp18691 := PrimCons(sym_1_1_6, tmp18690)

tmp18692 := PrimCons(symstring, tmp18691)

__e.TailApply(PrimFunc(symis_b), V6306, tmp18692, B6302, L6303, Key6304, C6305)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18693 := PrimValue(symshen_4_dsigf_d)

tmp18694 := Call(__e, PrimFunc(symshen_4assoc_1_6), symload, tmp18689, tmp18693)


tmp18695 := PrimSet(symshen_4_dsigf_d, tmp18694)

_ = tmp18695

tmp18696 := MakeNative(func(__e *ControlFlow) {
V6311 := __e.Get(1)
_ = V6311
__e.Return(MakeNative(func(__e *ControlFlow) {
B6307 := __e.Get(1)
_ = B6307
__e.Return(MakeNative(func(__e *ControlFlow) {
L6308 := __e.Get(1)
_ = L6308
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6309 := __e.Get(1)
_ = Key6309
__e.Return(MakeNative(func(__e *ControlFlow) {
C6310 := __e.Get(1)
_ = C6310
tmp18697 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp18698 := MakeNative(func(__e *ControlFlow) {
B := __e.Get(1)
_ = B
tmp18699 := PrimCons(B, Nil)

tmp18700 := PrimCons(sym_1_1_6, tmp18699)

tmp18701 := PrimCons(A, tmp18700)

tmp18702 := PrimCons(A, Nil)

tmp18703 := PrimCons(symlist, tmp18702)

tmp18704 := PrimCons(B, Nil)

tmp18705 := PrimCons(symlist, tmp18704)

tmp18706 := PrimCons(tmp18705, Nil)

tmp18707 := PrimCons(sym_1_1_6, tmp18706)

tmp18708 := PrimCons(tmp18703, tmp18707)

tmp18709 := PrimCons(tmp18708, Nil)

tmp18710 := PrimCons(sym_1_1_6, tmp18709)

tmp18711 := PrimCons(tmp18701, tmp18710)

tmp18712 := Call(__e, PrimFunc(symis_b), V6311, tmp18711, B6307, L6308, Key6309, C6310)


__e.TailApply(PrimFunc(symshen_4gc), B6307, tmp18712)
return


}, 1)

tmp18713 := Call(__e, PrimFunc(symshen_4newpv), B6307)


tmp18714 := Call(__e, tmp18698, tmp18713)


__e.TailApply(PrimFunc(symshen_4gc), B6307, tmp18714)
return


}, 1)

tmp18715 := Call(__e, PrimFunc(symshen_4newpv), B6307)


__e.TailApply(tmp18697, tmp18715)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18716 := PrimValue(symshen_4_dsigf_d)

tmp18717 := Call(__e, PrimFunc(symshen_4assoc_1_6), symmap, tmp18696, tmp18716)


tmp18718 := PrimSet(symshen_4_dsigf_d, tmp18717)

_ = tmp18718

tmp18719 := MakeNative(func(__e *ControlFlow) {
V6316 := __e.Get(1)
_ = V6316
__e.Return(MakeNative(func(__e *ControlFlow) {
B6312 := __e.Get(1)
_ = B6312
__e.Return(MakeNative(func(__e *ControlFlow) {
L6313 := __e.Get(1)
_ = L6313
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6314 := __e.Get(1)
_ = Key6314
__e.Return(MakeNative(func(__e *ControlFlow) {
C6315 := __e.Get(1)
_ = C6315
tmp18720 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp18721 := MakeNative(func(__e *ControlFlow) {
B := __e.Get(1)
_ = B
tmp18722 := PrimCons(B, Nil)

tmp18723 := PrimCons(symlist, tmp18722)

tmp18724 := PrimCons(tmp18723, Nil)

tmp18725 := PrimCons(sym_1_1_6, tmp18724)

tmp18726 := PrimCons(A, tmp18725)

tmp18727 := PrimCons(A, Nil)

tmp18728 := PrimCons(symlist, tmp18727)

tmp18729 := PrimCons(B, Nil)

tmp18730 := PrimCons(symlist, tmp18729)

tmp18731 := PrimCons(tmp18730, Nil)

tmp18732 := PrimCons(sym_1_1_6, tmp18731)

tmp18733 := PrimCons(tmp18728, tmp18732)

tmp18734 := PrimCons(tmp18733, Nil)

tmp18735 := PrimCons(sym_1_1_6, tmp18734)

tmp18736 := PrimCons(tmp18726, tmp18735)

tmp18737 := Call(__e, PrimFunc(symis_b), V6316, tmp18736, B6312, L6313, Key6314, C6315)


__e.TailApply(PrimFunc(symshen_4gc), B6312, tmp18737)
return


}, 1)

tmp18738 := Call(__e, PrimFunc(symshen_4newpv), B6312)


tmp18739 := Call(__e, tmp18721, tmp18738)


__e.TailApply(PrimFunc(symshen_4gc), B6312, tmp18739)
return


}, 1)

tmp18740 := Call(__e, PrimFunc(symshen_4newpv), B6312)


__e.TailApply(tmp18720, tmp18740)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18741 := PrimValue(symshen_4_dsigf_d)

tmp18742 := Call(__e, PrimFunc(symshen_4assoc_1_6), symmapcan, tmp18719, tmp18741)


tmp18743 := PrimSet(symshen_4_dsigf_d, tmp18742)

_ = tmp18743

tmp18744 := MakeNative(func(__e *ControlFlow) {
V6321 := __e.Get(1)
_ = V6321
__e.Return(MakeNative(func(__e *ControlFlow) {
B6317 := __e.Get(1)
_ = B6317
__e.Return(MakeNative(func(__e *ControlFlow) {
L6318 := __e.Get(1)
_ = L6318
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6319 := __e.Get(1)
_ = Key6319
__e.Return(MakeNative(func(__e *ControlFlow) {
C6320 := __e.Get(1)
_ = C6320
tmp18745 := PrimCons(symnumber, Nil)

tmp18746 := PrimCons(sym_1_1_6, tmp18745)

tmp18747 := PrimCons(symnumber, tmp18746)

__e.TailApply(PrimFunc(symis_b), V6321, tmp18747, B6317, L6318, Key6319, C6320)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18748 := PrimValue(symshen_4_dsigf_d)

tmp18749 := Call(__e, PrimFunc(symshen_4assoc_1_6), symmaxinferences, tmp18744, tmp18748)


tmp18750 := PrimSet(symshen_4_dsigf_d, tmp18749)

_ = tmp18750

tmp18751 := MakeNative(func(__e *ControlFlow) {
V6326 := __e.Get(1)
_ = V6326
__e.Return(MakeNative(func(__e *ControlFlow) {
B6322 := __e.Get(1)
_ = B6322
__e.Return(MakeNative(func(__e *ControlFlow) {
L6323 := __e.Get(1)
_ = L6323
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6324 := __e.Get(1)
_ = Key6324
__e.Return(MakeNative(func(__e *ControlFlow) {
C6325 := __e.Get(1)
_ = C6325
tmp18752 := PrimCons(symstring, Nil)

tmp18753 := PrimCons(sym_1_1_6, tmp18752)

tmp18754 := PrimCons(symnumber, tmp18753)

__e.TailApply(PrimFunc(symis_b), V6326, tmp18754, B6322, L6323, Key6324, C6325)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18755 := PrimValue(symshen_4_dsigf_d)

tmp18756 := Call(__e, PrimFunc(symshen_4assoc_1_6), symn_1_6string, tmp18751, tmp18755)


tmp18757 := PrimSet(symshen_4_dsigf_d, tmp18756)

_ = tmp18757

tmp18758 := MakeNative(func(__e *ControlFlow) {
V6331 := __e.Get(1)
_ = V6331
__e.Return(MakeNative(func(__e *ControlFlow) {
B6327 := __e.Get(1)
_ = B6327
__e.Return(MakeNative(func(__e *ControlFlow) {
L6328 := __e.Get(1)
_ = L6328
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6329 := __e.Get(1)
_ = Key6329
__e.Return(MakeNative(func(__e *ControlFlow) {
C6330 := __e.Get(1)
_ = C6330
tmp18759 := PrimCons(symnumber, Nil)

tmp18760 := PrimCons(sym_1_1_6, tmp18759)

tmp18761 := PrimCons(symnumber, tmp18760)

__e.TailApply(PrimFunc(symis_b), V6331, tmp18761, B6327, L6328, Key6329, C6330)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18762 := PrimValue(symshen_4_dsigf_d)

tmp18763 := Call(__e, PrimFunc(symshen_4assoc_1_6), symnl, tmp18758, tmp18762)


tmp18764 := PrimSet(symshen_4_dsigf_d, tmp18763)

_ = tmp18764

tmp18765 := MakeNative(func(__e *ControlFlow) {
V6336 := __e.Get(1)
_ = V6336
__e.Return(MakeNative(func(__e *ControlFlow) {
B6332 := __e.Get(1)
_ = B6332
__e.Return(MakeNative(func(__e *ControlFlow) {
L6333 := __e.Get(1)
_ = L6333
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6334 := __e.Get(1)
_ = Key6334
__e.Return(MakeNative(func(__e *ControlFlow) {
C6335 := __e.Get(1)
_ = C6335
tmp18766 := PrimCons(symboolean, Nil)

tmp18767 := PrimCons(sym_1_1_6, tmp18766)

tmp18768 := PrimCons(symboolean, tmp18767)

__e.TailApply(PrimFunc(symis_b), V6336, tmp18768, B6332, L6333, Key6334, C6335)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18769 := PrimValue(symshen_4_dsigf_d)

tmp18770 := Call(__e, PrimFunc(symshen_4assoc_1_6), symnot, tmp18765, tmp18769)


tmp18771 := PrimSet(symshen_4_dsigf_d, tmp18770)

_ = tmp18771

tmp18772 := MakeNative(func(__e *ControlFlow) {
V6341 := __e.Get(1)
_ = V6341
__e.Return(MakeNative(func(__e *ControlFlow) {
B6337 := __e.Get(1)
_ = B6337
__e.Return(MakeNative(func(__e *ControlFlow) {
L6338 := __e.Get(1)
_ = L6338
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6339 := __e.Get(1)
_ = Key6339
__e.Return(MakeNative(func(__e *ControlFlow) {
C6340 := __e.Get(1)
_ = C6340
tmp18773 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp18774 := PrimCons(A, Nil)

tmp18775 := PrimCons(symlist, tmp18774)

tmp18776 := PrimCons(A, Nil)

tmp18777 := PrimCons(sym_1_1_6, tmp18776)

tmp18778 := PrimCons(tmp18775, tmp18777)

tmp18779 := PrimCons(tmp18778, Nil)

tmp18780 := PrimCons(sym_1_1_6, tmp18779)

tmp18781 := PrimCons(symnumber, tmp18780)

tmp18782 := Call(__e, PrimFunc(symis_b), V6341, tmp18781, B6337, L6338, Key6339, C6340)


__e.TailApply(PrimFunc(symshen_4gc), B6337, tmp18782)
return


}, 1)

tmp18783 := Call(__e, PrimFunc(symshen_4newpv), B6337)


__e.TailApply(tmp18773, tmp18783)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18784 := PrimValue(symshen_4_dsigf_d)

tmp18785 := Call(__e, PrimFunc(symshen_4assoc_1_6), symnth, tmp18772, tmp18784)


tmp18786 := PrimSet(symshen_4_dsigf_d, tmp18785)

_ = tmp18786

tmp18787 := MakeNative(func(__e *ControlFlow) {
V6346 := __e.Get(1)
_ = V6346
__e.Return(MakeNative(func(__e *ControlFlow) {
B6342 := __e.Get(1)
_ = B6342
__e.Return(MakeNative(func(__e *ControlFlow) {
L6343 := __e.Get(1)
_ = L6343
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6344 := __e.Get(1)
_ = Key6344
__e.Return(MakeNative(func(__e *ControlFlow) {
C6345 := __e.Get(1)
_ = C6345
tmp18788 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp18789 := PrimCons(symboolean, Nil)

tmp18790 := PrimCons(sym_1_1_6, tmp18789)

tmp18791 := PrimCons(A, tmp18790)

tmp18792 := Call(__e, PrimFunc(symis_b), V6346, tmp18791, B6342, L6343, Key6344, C6345)


__e.TailApply(PrimFunc(symshen_4gc), B6342, tmp18792)
return


}, 1)

tmp18793 := Call(__e, PrimFunc(symshen_4newpv), B6342)


__e.TailApply(tmp18788, tmp18793)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18794 := PrimValue(symshen_4_dsigf_d)

tmp18795 := Call(__e, PrimFunc(symshen_4assoc_1_6), symnumber_2, tmp18787, tmp18794)


tmp18796 := PrimSet(symshen_4_dsigf_d, tmp18795)

_ = tmp18796

tmp18797 := MakeNative(func(__e *ControlFlow) {
V6351 := __e.Get(1)
_ = V6351
__e.Return(MakeNative(func(__e *ControlFlow) {
B6347 := __e.Get(1)
_ = B6347
__e.Return(MakeNative(func(__e *ControlFlow) {
L6348 := __e.Get(1)
_ = L6348
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6349 := __e.Get(1)
_ = Key6349
__e.Return(MakeNative(func(__e *ControlFlow) {
C6350 := __e.Get(1)
_ = C6350
tmp18798 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp18799 := MakeNative(func(__e *ControlFlow) {
B := __e.Get(1)
_ = B
tmp18800 := PrimCons(symnumber, Nil)

tmp18801 := PrimCons(sym_1_1_6, tmp18800)

tmp18802 := PrimCons(B, tmp18801)

tmp18803 := PrimCons(tmp18802, Nil)

tmp18804 := PrimCons(sym_1_1_6, tmp18803)

tmp18805 := PrimCons(A, tmp18804)

tmp18806 := Call(__e, PrimFunc(symis_b), V6351, tmp18805, B6347, L6348, Key6349, C6350)


__e.TailApply(PrimFunc(symshen_4gc), B6347, tmp18806)
return


}, 1)

tmp18807 := Call(__e, PrimFunc(symshen_4newpv), B6347)


tmp18808 := Call(__e, tmp18799, tmp18807)


__e.TailApply(PrimFunc(symshen_4gc), B6347, tmp18808)
return


}, 1)

tmp18809 := Call(__e, PrimFunc(symshen_4newpv), B6347)


__e.TailApply(tmp18798, tmp18809)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18810 := PrimValue(symshen_4_dsigf_d)

tmp18811 := Call(__e, PrimFunc(symshen_4assoc_1_6), symoccurrences, tmp18797, tmp18810)


tmp18812 := PrimSet(symshen_4_dsigf_d, tmp18811)

_ = tmp18812

tmp18813 := MakeNative(func(__e *ControlFlow) {
V6356 := __e.Get(1)
_ = V6356
__e.Return(MakeNative(func(__e *ControlFlow) {
B6352 := __e.Get(1)
_ = B6352
__e.Return(MakeNative(func(__e *ControlFlow) {
L6353 := __e.Get(1)
_ = L6353
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6354 := __e.Get(1)
_ = Key6354
__e.Return(MakeNative(func(__e *ControlFlow) {
C6355 := __e.Get(1)
_ = C6355
tmp18814 := PrimCons(symboolean, Nil)

tmp18815 := PrimCons(sym_1_1_6, tmp18814)

tmp18816 := PrimCons(symsymbol, tmp18815)

__e.TailApply(PrimFunc(symis_b), V6356, tmp18816, B6352, L6353, Key6354, C6355)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18817 := PrimValue(symshen_4_dsigf_d)

tmp18818 := Call(__e, PrimFunc(symshen_4assoc_1_6), symoccurs_1check, tmp18813, tmp18817)


tmp18819 := PrimSet(symshen_4_dsigf_d, tmp18818)

_ = tmp18819

tmp18820 := MakeNative(func(__e *ControlFlow) {
V6361 := __e.Get(1)
_ = V6361
__e.Return(MakeNative(func(__e *ControlFlow) {
B6357 := __e.Get(1)
_ = B6357
__e.Return(MakeNative(func(__e *ControlFlow) {
L6358 := __e.Get(1)
_ = L6358
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6359 := __e.Get(1)
_ = Key6359
__e.Return(MakeNative(func(__e *ControlFlow) {
C6360 := __e.Get(1)
_ = C6360
tmp18821 := PrimCons(symboolean, Nil)

tmp18822 := PrimCons(sym_1_1_6, tmp18821)

__e.TailApply(PrimFunc(symis_b), V6361, tmp18822, B6357, L6358, Key6359, C6360)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18823 := PrimValue(symshen_4_dsigf_d)

tmp18824 := Call(__e, PrimFunc(symshen_4assoc_1_6), symoccurs_2, tmp18820, tmp18823)


tmp18825 := PrimSet(symshen_4_dsigf_d, tmp18824)

_ = tmp18825

tmp18826 := MakeNative(func(__e *ControlFlow) {
V6366 := __e.Get(1)
_ = V6366
__e.Return(MakeNative(func(__e *ControlFlow) {
B6362 := __e.Get(1)
_ = B6362
__e.Return(MakeNative(func(__e *ControlFlow) {
L6363 := __e.Get(1)
_ = L6363
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6364 := __e.Get(1)
_ = Key6364
__e.Return(MakeNative(func(__e *ControlFlow) {
C6365 := __e.Get(1)
_ = C6365
tmp18827 := PrimCons(symboolean, Nil)

tmp18828 := PrimCons(sym_1_1_6, tmp18827)

tmp18829 := PrimCons(symsymbol, tmp18828)

__e.TailApply(PrimFunc(symis_b), V6366, tmp18829, B6362, L6363, Key6364, C6365)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18830 := PrimValue(symshen_4_dsigf_d)

tmp18831 := Call(__e, PrimFunc(symshen_4assoc_1_6), symoptimise, tmp18826, tmp18830)


tmp18832 := PrimSet(symshen_4_dsigf_d, tmp18831)

_ = tmp18832

tmp18833 := MakeNative(func(__e *ControlFlow) {
V6371 := __e.Get(1)
_ = V6371
__e.Return(MakeNative(func(__e *ControlFlow) {
B6367 := __e.Get(1)
_ = B6367
__e.Return(MakeNative(func(__e *ControlFlow) {
L6368 := __e.Get(1)
_ = L6368
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6369 := __e.Get(1)
_ = Key6369
__e.Return(MakeNative(func(__e *ControlFlow) {
C6370 := __e.Get(1)
_ = C6370
tmp18834 := PrimCons(symboolean, Nil)

tmp18835 := PrimCons(sym_1_1_6, tmp18834)

__e.TailApply(PrimFunc(symis_b), V6371, tmp18835, B6367, L6368, Key6369, C6370)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18836 := PrimValue(symshen_4_dsigf_d)

tmp18837 := Call(__e, PrimFunc(symshen_4assoc_1_6), symoptimise_2, tmp18833, tmp18836)


tmp18838 := PrimSet(symshen_4_dsigf_d, tmp18837)

_ = tmp18838

tmp18839 := MakeNative(func(__e *ControlFlow) {
V6376 := __e.Get(1)
_ = V6376
__e.Return(MakeNative(func(__e *ControlFlow) {
B6372 := __e.Get(1)
_ = B6372
__e.Return(MakeNative(func(__e *ControlFlow) {
L6373 := __e.Get(1)
_ = L6373
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6374 := __e.Get(1)
_ = Key6374
__e.Return(MakeNative(func(__e *ControlFlow) {
C6375 := __e.Get(1)
_ = C6375
tmp18840 := PrimCons(symboolean, Nil)

tmp18841 := PrimCons(sym_1_1_6, tmp18840)

tmp18842 := PrimCons(symboolean, tmp18841)

tmp18843 := PrimCons(tmp18842, Nil)

tmp18844 := PrimCons(sym_1_1_6, tmp18843)

tmp18845 := PrimCons(symboolean, tmp18844)

__e.TailApply(PrimFunc(symis_b), V6376, tmp18845, B6372, L6373, Key6374, C6375)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18846 := PrimValue(symshen_4_dsigf_d)

tmp18847 := Call(__e, PrimFunc(symshen_4assoc_1_6), symor, tmp18839, tmp18846)


tmp18848 := PrimSet(symshen_4_dsigf_d, tmp18847)

_ = tmp18848

tmp18849 := MakeNative(func(__e *ControlFlow) {
V6381 := __e.Get(1)
_ = V6381
__e.Return(MakeNative(func(__e *ControlFlow) {
B6377 := __e.Get(1)
_ = B6377
__e.Return(MakeNative(func(__e *ControlFlow) {
L6378 := __e.Get(1)
_ = L6378
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6379 := __e.Get(1)
_ = Key6379
__e.Return(MakeNative(func(__e *ControlFlow) {
C6380 := __e.Get(1)
_ = C6380
tmp18850 := PrimCons(symstring, Nil)

tmp18851 := PrimCons(sym_1_1_6, tmp18850)

__e.TailApply(PrimFunc(symis_b), V6381, tmp18851, B6377, L6378, Key6379, C6380)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18852 := PrimValue(symshen_4_dsigf_d)

tmp18853 := Call(__e, PrimFunc(symshen_4assoc_1_6), symos, tmp18849, tmp18852)


tmp18854 := PrimSet(symshen_4_dsigf_d, tmp18853)

_ = tmp18854

tmp18855 := MakeNative(func(__e *ControlFlow) {
V6386 := __e.Get(1)
_ = V6386
__e.Return(MakeNative(func(__e *ControlFlow) {
B6382 := __e.Get(1)
_ = B6382
__e.Return(MakeNative(func(__e *ControlFlow) {
L6383 := __e.Get(1)
_ = L6383
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6384 := __e.Get(1)
_ = Key6384
__e.Return(MakeNative(func(__e *ControlFlow) {
C6385 := __e.Get(1)
_ = C6385
tmp18856 := PrimCons(symboolean, Nil)

tmp18857 := PrimCons(sym_1_1_6, tmp18856)

tmp18858 := PrimCons(symsymbol, tmp18857)

__e.TailApply(PrimFunc(symis_b), V6386, tmp18858, B6382, L6383, Key6384, C6385)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18859 := PrimValue(symshen_4_dsigf_d)

tmp18860 := Call(__e, PrimFunc(symshen_4assoc_1_6), sympackage_2, tmp18855, tmp18859)


tmp18861 := PrimSet(symshen_4_dsigf_d, tmp18860)

_ = tmp18861

tmp18862 := MakeNative(func(__e *ControlFlow) {
V6391 := __e.Get(1)
_ = V6391
__e.Return(MakeNative(func(__e *ControlFlow) {
B6387 := __e.Get(1)
_ = B6387
__e.Return(MakeNative(func(__e *ControlFlow) {
L6388 := __e.Get(1)
_ = L6388
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6389 := __e.Get(1)
_ = Key6389
__e.Return(MakeNative(func(__e *ControlFlow) {
C6390 := __e.Get(1)
_ = C6390
tmp18863 := PrimCons(symstring, Nil)

tmp18864 := PrimCons(sym_1_1_6, tmp18863)

__e.TailApply(PrimFunc(symis_b), V6391, tmp18864, B6387, L6388, Key6389, C6390)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18865 := PrimValue(symshen_4_dsigf_d)

tmp18866 := Call(__e, PrimFunc(symshen_4assoc_1_6), symport, tmp18862, tmp18865)


tmp18867 := PrimSet(symshen_4_dsigf_d, tmp18866)

_ = tmp18867

tmp18868 := MakeNative(func(__e *ControlFlow) {
V6396 := __e.Get(1)
_ = V6396
__e.Return(MakeNative(func(__e *ControlFlow) {
B6392 := __e.Get(1)
_ = B6392
__e.Return(MakeNative(func(__e *ControlFlow) {
L6393 := __e.Get(1)
_ = L6393
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6394 := __e.Get(1)
_ = Key6394
__e.Return(MakeNative(func(__e *ControlFlow) {
C6395 := __e.Get(1)
_ = C6395
tmp18869 := PrimCons(symstring, Nil)

tmp18870 := PrimCons(sym_1_1_6, tmp18869)

__e.TailApply(PrimFunc(symis_b), V6396, tmp18870, B6392, L6393, Key6394, C6395)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18871 := PrimValue(symshen_4_dsigf_d)

tmp18872 := Call(__e, PrimFunc(symshen_4assoc_1_6), symporters, tmp18868, tmp18871)


tmp18873 := PrimSet(symshen_4_dsigf_d, tmp18872)

_ = tmp18873

tmp18874 := MakeNative(func(__e *ControlFlow) {
V6401 := __e.Get(1)
_ = V6401
__e.Return(MakeNative(func(__e *ControlFlow) {
B6397 := __e.Get(1)
_ = B6397
__e.Return(MakeNative(func(__e *ControlFlow) {
L6398 := __e.Get(1)
_ = L6398
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6399 := __e.Get(1)
_ = Key6399
__e.Return(MakeNative(func(__e *ControlFlow) {
C6400 := __e.Get(1)
_ = C6400
tmp18875 := PrimCons(symstring, Nil)

tmp18876 := PrimCons(sym_1_1_6, tmp18875)

tmp18877 := PrimCons(symnumber, tmp18876)

tmp18878 := PrimCons(tmp18877, Nil)

tmp18879 := PrimCons(sym_1_1_6, tmp18878)

tmp18880 := PrimCons(symstring, tmp18879)

__e.TailApply(PrimFunc(symis_b), V6401, tmp18880, B6397, L6398, Key6399, C6400)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18881 := PrimValue(symshen_4_dsigf_d)

tmp18882 := Call(__e, PrimFunc(symshen_4assoc_1_6), sympos, tmp18874, tmp18881)


tmp18883 := PrimSet(symshen_4_dsigf_d, tmp18882)

_ = tmp18883

tmp18884 := MakeNative(func(__e *ControlFlow) {
V6406 := __e.Get(1)
_ = V6406
__e.Return(MakeNative(func(__e *ControlFlow) {
B6402 := __e.Get(1)
_ = B6402
__e.Return(MakeNative(func(__e *ControlFlow) {
L6403 := __e.Get(1)
_ = L6403
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6404 := __e.Get(1)
_ = Key6404
__e.Return(MakeNative(func(__e *ControlFlow) {
C6405 := __e.Get(1)
_ = C6405
tmp18885 := PrimCons(symout, Nil)

tmp18886 := PrimCons(symstream, tmp18885)

tmp18887 := PrimCons(symstring, Nil)

tmp18888 := PrimCons(sym_1_1_6, tmp18887)

tmp18889 := PrimCons(tmp18886, tmp18888)

tmp18890 := PrimCons(tmp18889, Nil)

tmp18891 := PrimCons(sym_1_1_6, tmp18890)

tmp18892 := PrimCons(symstring, tmp18891)

__e.TailApply(PrimFunc(symis_b), V6406, tmp18892, B6402, L6403, Key6404, C6405)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18893 := PrimValue(symshen_4_dsigf_d)

tmp18894 := Call(__e, PrimFunc(symshen_4assoc_1_6), sympr, tmp18884, tmp18893)


tmp18895 := PrimSet(symshen_4_dsigf_d, tmp18894)

_ = tmp18895

tmp18896 := MakeNative(func(__e *ControlFlow) {
V6411 := __e.Get(1)
_ = V6411
__e.Return(MakeNative(func(__e *ControlFlow) {
B6407 := __e.Get(1)
_ = B6407
__e.Return(MakeNative(func(__e *ControlFlow) {
L6408 := __e.Get(1)
_ = L6408
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6409 := __e.Get(1)
_ = Key6409
__e.Return(MakeNative(func(__e *ControlFlow) {
C6410 := __e.Get(1)
_ = C6410
tmp18897 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp18898 := PrimCons(A, Nil)

tmp18899 := PrimCons(sym_1_1_6, tmp18898)

tmp18900 := PrimCons(A, tmp18899)

tmp18901 := Call(__e, PrimFunc(symis_b), V6411, tmp18900, B6407, L6408, Key6409, C6410)


__e.TailApply(PrimFunc(symshen_4gc), B6407, tmp18901)
return


}, 1)

tmp18902 := Call(__e, PrimFunc(symshen_4newpv), B6407)


__e.TailApply(tmp18897, tmp18902)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18903 := PrimValue(symshen_4_dsigf_d)

tmp18904 := Call(__e, PrimFunc(symshen_4assoc_1_6), symprint, tmp18896, tmp18903)


tmp18905 := PrimSet(symshen_4_dsigf_d, tmp18904)

_ = tmp18905

tmp18906 := MakeNative(func(__e *ControlFlow) {
V6416 := __e.Get(1)
_ = V6416
__e.Return(MakeNative(func(__e *ControlFlow) {
B6412 := __e.Get(1)
_ = B6412
__e.Return(MakeNative(func(__e *ControlFlow) {
L6413 := __e.Get(1)
_ = L6413
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6414 := __e.Get(1)
_ = Key6414
__e.Return(MakeNative(func(__e *ControlFlow) {
C6415 := __e.Get(1)
_ = C6415
tmp18907 := PrimCons(symsymbol, Nil)

tmp18908 := PrimCons(sym_1_1_6, tmp18907)

tmp18909 := PrimCons(symsymbol, tmp18908)

__e.TailApply(PrimFunc(symis_b), V6416, tmp18909, B6412, L6413, Key6414, C6415)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18910 := PrimValue(symshen_4_dsigf_d)

tmp18911 := Call(__e, PrimFunc(symshen_4assoc_1_6), symprofile, tmp18906, tmp18910)


tmp18912 := PrimSet(symshen_4_dsigf_d, tmp18911)

_ = tmp18912

tmp18913 := MakeNative(func(__e *ControlFlow) {
V6421 := __e.Get(1)
_ = V6421
__e.Return(MakeNative(func(__e *ControlFlow) {
B6417 := __e.Get(1)
_ = B6417
__e.Return(MakeNative(func(__e *ControlFlow) {
L6418 := __e.Get(1)
_ = L6418
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6419 := __e.Get(1)
_ = Key6419
__e.Return(MakeNative(func(__e *ControlFlow) {
C6420 := __e.Get(1)
_ = C6420
tmp18914 := PrimCons(symsymbol, Nil)

tmp18915 := PrimCons(symlist, tmp18914)

tmp18916 := PrimCons(symsymbol, Nil)

tmp18917 := PrimCons(symlist, tmp18916)

tmp18918 := PrimCons(tmp18917, Nil)

tmp18919 := PrimCons(sym_1_1_6, tmp18918)

tmp18920 := PrimCons(tmp18915, tmp18919)

__e.TailApply(PrimFunc(symis_b), V6421, tmp18920, B6417, L6418, Key6419, C6420)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18921 := PrimValue(symshen_4_dsigf_d)

tmp18922 := Call(__e, PrimFunc(symshen_4assoc_1_6), sympreclude, tmp18913, tmp18921)


tmp18923 := PrimSet(symshen_4_dsigf_d, tmp18922)

_ = tmp18923

tmp18924 := MakeNative(func(__e *ControlFlow) {
V6426 := __e.Get(1)
_ = V6426
__e.Return(MakeNative(func(__e *ControlFlow) {
B6422 := __e.Get(1)
_ = B6422
__e.Return(MakeNative(func(__e *ControlFlow) {
L6423 := __e.Get(1)
_ = L6423
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6424 := __e.Get(1)
_ = Key6424
__e.Return(MakeNative(func(__e *ControlFlow) {
C6425 := __e.Get(1)
_ = C6425
tmp18925 := PrimCons(symstring, Nil)

tmp18926 := PrimCons(sym_1_1_6, tmp18925)

tmp18927 := PrimCons(symstring, tmp18926)

__e.TailApply(PrimFunc(symis_b), V6426, tmp18927, B6422, L6423, Key6424, C6425)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18928 := PrimValue(symshen_4_dsigf_d)

tmp18929 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4proc_1nl, tmp18924, tmp18928)


tmp18930 := PrimSet(symshen_4_dsigf_d, tmp18929)

_ = tmp18930

tmp18931 := MakeNative(func(__e *ControlFlow) {
V6431 := __e.Get(1)
_ = V6431
__e.Return(MakeNative(func(__e *ControlFlow) {
B6427 := __e.Get(1)
_ = B6427
__e.Return(MakeNative(func(__e *ControlFlow) {
L6428 := __e.Get(1)
_ = L6428
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6429 := __e.Get(1)
_ = Key6429
__e.Return(MakeNative(func(__e *ControlFlow) {
C6430 := __e.Get(1)
_ = C6430
tmp18932 := PrimCons(symnumber, Nil)

tmp18933 := PrimCons(sym_d, tmp18932)

tmp18934 := PrimCons(symsymbol, tmp18933)

tmp18935 := PrimCons(tmp18934, Nil)

tmp18936 := PrimCons(sym_1_1_6, tmp18935)

tmp18937 := PrimCons(symsymbol, tmp18936)

__e.TailApply(PrimFunc(symis_b), V6431, tmp18937, B6427, L6428, Key6429, C6430)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18938 := PrimValue(symshen_4_dsigf_d)

tmp18939 := Call(__e, PrimFunc(symshen_4assoc_1_6), symprofile_1results, tmp18931, tmp18938)


tmp18940 := PrimSet(symshen_4_dsigf_d, tmp18939)

_ = tmp18940

tmp18941 := MakeNative(func(__e *ControlFlow) {
V6436 := __e.Get(1)
_ = V6436
__e.Return(MakeNative(func(__e *ControlFlow) {
B6432 := __e.Get(1)
_ = B6432
__e.Return(MakeNative(func(__e *ControlFlow) {
L6433 := __e.Get(1)
_ = L6433
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6434 := __e.Get(1)
_ = Key6434
__e.Return(MakeNative(func(__e *ControlFlow) {
C6435 := __e.Get(1)
_ = C6435
tmp18942 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp18943 := PrimCons(A, Nil)

tmp18944 := PrimCons(sym_1_1_6, tmp18943)

tmp18945 := PrimCons(A, tmp18944)

tmp18946 := Call(__e, PrimFunc(symis_b), V6436, tmp18945, B6432, L6433, Key6434, C6435)


__e.TailApply(PrimFunc(symshen_4gc), B6432, tmp18946)
return


}, 1)

tmp18947 := Call(__e, PrimFunc(symshen_4newpv), B6432)


__e.TailApply(tmp18942, tmp18947)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18948 := PrimValue(symshen_4_dsigf_d)

tmp18949 := Call(__e, PrimFunc(symshen_4assoc_1_6), symprotect, tmp18941, tmp18948)


tmp18950 := PrimSet(symshen_4_dsigf_d, tmp18949)

_ = tmp18950

tmp18951 := MakeNative(func(__e *ControlFlow) {
V6441 := __e.Get(1)
_ = V6441
__e.Return(MakeNative(func(__e *ControlFlow) {
B6437 := __e.Get(1)
_ = B6437
__e.Return(MakeNative(func(__e *ControlFlow) {
L6438 := __e.Get(1)
_ = L6438
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6439 := __e.Get(1)
_ = Key6439
__e.Return(MakeNative(func(__e *ControlFlow) {
C6440 := __e.Get(1)
_ = C6440
tmp18952 := PrimCons(symsymbol, Nil)

tmp18953 := PrimCons(symlist, tmp18952)

tmp18954 := PrimCons(symsymbol, Nil)

tmp18955 := PrimCons(symlist, tmp18954)

tmp18956 := PrimCons(tmp18955, Nil)

tmp18957 := PrimCons(sym_1_1_6, tmp18956)

tmp18958 := PrimCons(tmp18953, tmp18957)

__e.TailApply(PrimFunc(symis_b), V6441, tmp18958, B6437, L6438, Key6439, C6440)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18959 := PrimValue(symshen_4_dsigf_d)

tmp18960 := Call(__e, PrimFunc(symshen_4assoc_1_6), sympreclude_1all_1but, tmp18951, tmp18959)


tmp18961 := PrimSet(symshen_4_dsigf_d, tmp18960)

_ = tmp18961

tmp18962 := MakeNative(func(__e *ControlFlow) {
V6446 := __e.Get(1)
_ = V6446
__e.Return(MakeNative(func(__e *ControlFlow) {
B6442 := __e.Get(1)
_ = B6442
__e.Return(MakeNative(func(__e *ControlFlow) {
L6443 := __e.Get(1)
_ = L6443
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6444 := __e.Get(1)
_ = Key6444
__e.Return(MakeNative(func(__e *ControlFlow) {
C6445 := __e.Get(1)
_ = C6445
tmp18963 := PrimCons(symout, Nil)

tmp18964 := PrimCons(symstream, tmp18963)

tmp18965 := PrimCons(symstring, Nil)

tmp18966 := PrimCons(sym_1_1_6, tmp18965)

tmp18967 := PrimCons(tmp18964, tmp18966)

tmp18968 := PrimCons(tmp18967, Nil)

tmp18969 := PrimCons(sym_1_1_6, tmp18968)

tmp18970 := PrimCons(symstring, tmp18969)

__e.TailApply(PrimFunc(symis_b), V6446, tmp18970, B6442, L6443, Key6444, C6445)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18971 := PrimValue(symshen_4_dsigf_d)

tmp18972 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4prhush, tmp18962, tmp18971)


tmp18973 := PrimSet(symshen_4_dsigf_d, tmp18972)

_ = tmp18973

tmp18974 := MakeNative(func(__e *ControlFlow) {
V6451 := __e.Get(1)
_ = V6451
__e.Return(MakeNative(func(__e *ControlFlow) {
B6447 := __e.Get(1)
_ = B6447
__e.Return(MakeNative(func(__e *ControlFlow) {
L6448 := __e.Get(1)
_ = L6448
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6449 := __e.Get(1)
_ = Key6449
__e.Return(MakeNative(func(__e *ControlFlow) {
C6450 := __e.Get(1)
_ = C6450
tmp18975 := PrimCons(symnumber, Nil)

tmp18976 := PrimCons(sym_1_1_6, tmp18975)

tmp18977 := PrimCons(symnumber, tmp18976)

__e.TailApply(PrimFunc(symis_b), V6451, tmp18977, B6447, L6448, Key6449, C6450)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18978 := PrimValue(symshen_4_dsigf_d)

tmp18979 := Call(__e, PrimFunc(symshen_4assoc_1_6), symprolog_1memory, tmp18974, tmp18978)


tmp18980 := PrimSet(symshen_4_dsigf_d, tmp18979)

_ = tmp18980

tmp18981 := MakeNative(func(__e *ControlFlow) {
V6456 := __e.Get(1)
_ = V6456
__e.Return(MakeNative(func(__e *ControlFlow) {
B6452 := __e.Get(1)
_ = B6452
__e.Return(MakeNative(func(__e *ControlFlow) {
L6453 := __e.Get(1)
_ = L6453
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6454 := __e.Get(1)
_ = Key6454
__e.Return(MakeNative(func(__e *ControlFlow) {
C6455 := __e.Get(1)
_ = C6455
tmp18982 := PrimCons(symunit, Nil)

tmp18983 := PrimCons(symlist, tmp18982)

tmp18984 := PrimCons(tmp18983, Nil)

tmp18985 := PrimCons(sym_1_1_6, tmp18984)

tmp18986 := PrimCons(symsymbol, tmp18985)

__e.TailApply(PrimFunc(symis_b), V6456, tmp18986, B6452, L6453, Key6454, C6455)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18987 := PrimValue(symshen_4_dsigf_d)

tmp18988 := Call(__e, PrimFunc(symshen_4assoc_1_6), symps, tmp18981, tmp18987)


tmp18989 := PrimSet(symshen_4_dsigf_d, tmp18988)

_ = tmp18989

tmp18990 := MakeNative(func(__e *ControlFlow) {
V6461 := __e.Get(1)
_ = V6461
__e.Return(MakeNative(func(__e *ControlFlow) {
B6457 := __e.Get(1)
_ = B6457
__e.Return(MakeNative(func(__e *ControlFlow) {
L6458 := __e.Get(1)
_ = L6458
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6459 := __e.Get(1)
_ = Key6459
__e.Return(MakeNative(func(__e *ControlFlow) {
C6460 := __e.Get(1)
_ = C6460
tmp18991 := PrimCons(symin, Nil)

tmp18992 := PrimCons(symstream, tmp18991)

tmp18993 := PrimCons(symunit, Nil)

tmp18994 := PrimCons(sym_1_1_6, tmp18993)

tmp18995 := PrimCons(tmp18992, tmp18994)

__e.TailApply(PrimFunc(symis_b), V6461, tmp18995, B6457, L6458, Key6459, C6460)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18996 := PrimValue(symshen_4_dsigf_d)

tmp18997 := Call(__e, PrimFunc(symshen_4assoc_1_6), symread, tmp18990, tmp18996)


tmp18998 := PrimSet(symshen_4_dsigf_d, tmp18997)

_ = tmp18998

tmp18999 := MakeNative(func(__e *ControlFlow) {
V6466 := __e.Get(1)
_ = V6466
__e.Return(MakeNative(func(__e *ControlFlow) {
B6462 := __e.Get(1)
_ = B6462
__e.Return(MakeNative(func(__e *ControlFlow) {
L6463 := __e.Get(1)
_ = L6463
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6464 := __e.Get(1)
_ = Key6464
__e.Return(MakeNative(func(__e *ControlFlow) {
C6465 := __e.Get(1)
_ = C6465
tmp19000 := PrimCons(symin, Nil)

tmp19001 := PrimCons(symstream, tmp19000)

tmp19002 := PrimCons(symnumber, Nil)

tmp19003 := PrimCons(sym_1_1_6, tmp19002)

tmp19004 := PrimCons(tmp19001, tmp19003)

__e.TailApply(PrimFunc(symis_b), V6466, tmp19004, B6462, L6463, Key6464, C6465)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19005 := PrimValue(symshen_4_dsigf_d)

tmp19006 := Call(__e, PrimFunc(symshen_4assoc_1_6), symread_1byte, tmp18999, tmp19005)


tmp19007 := PrimSet(symshen_4_dsigf_d, tmp19006)

_ = tmp19007

tmp19008 := MakeNative(func(__e *ControlFlow) {
V6471 := __e.Get(1)
_ = V6471
__e.Return(MakeNative(func(__e *ControlFlow) {
B6467 := __e.Get(1)
_ = B6467
__e.Return(MakeNative(func(__e *ControlFlow) {
L6468 := __e.Get(1)
_ = L6468
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6469 := __e.Get(1)
_ = Key6469
__e.Return(MakeNative(func(__e *ControlFlow) {
C6470 := __e.Get(1)
_ = C6470
tmp19009 := PrimCons(symnumber, Nil)

tmp19010 := PrimCons(symlist, tmp19009)

tmp19011 := PrimCons(tmp19010, Nil)

tmp19012 := PrimCons(sym_1_1_6, tmp19011)

tmp19013 := PrimCons(symstring, tmp19012)

__e.TailApply(PrimFunc(symis_b), V6471, tmp19013, B6467, L6468, Key6469, C6470)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19014 := PrimValue(symshen_4_dsigf_d)

tmp19015 := Call(__e, PrimFunc(symshen_4assoc_1_6), symread_1file_1as_1bytelist, tmp19008, tmp19014)


tmp19016 := PrimSet(symshen_4_dsigf_d, tmp19015)

_ = tmp19016

tmp19017 := MakeNative(func(__e *ControlFlow) {
V6476 := __e.Get(1)
_ = V6476
__e.Return(MakeNative(func(__e *ControlFlow) {
B6472 := __e.Get(1)
_ = B6472
__e.Return(MakeNative(func(__e *ControlFlow) {
L6473 := __e.Get(1)
_ = L6473
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6474 := __e.Get(1)
_ = Key6474
__e.Return(MakeNative(func(__e *ControlFlow) {
C6475 := __e.Get(1)
_ = C6475
tmp19018 := PrimCons(symstring, Nil)

tmp19019 := PrimCons(sym_1_1_6, tmp19018)

tmp19020 := PrimCons(symstring, tmp19019)

__e.TailApply(PrimFunc(symis_b), V6476, tmp19020, B6472, L6473, Key6474, C6475)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19021 := PrimValue(symshen_4_dsigf_d)

tmp19022 := Call(__e, PrimFunc(symshen_4assoc_1_6), symread_1file_1as_1string, tmp19017, tmp19021)


tmp19023 := PrimSet(symshen_4_dsigf_d, tmp19022)

_ = tmp19023

tmp19024 := MakeNative(func(__e *ControlFlow) {
V6481 := __e.Get(1)
_ = V6481
__e.Return(MakeNative(func(__e *ControlFlow) {
B6477 := __e.Get(1)
_ = B6477
__e.Return(MakeNative(func(__e *ControlFlow) {
L6478 := __e.Get(1)
_ = L6478
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6479 := __e.Get(1)
_ = Key6479
__e.Return(MakeNative(func(__e *ControlFlow) {
C6480 := __e.Get(1)
_ = C6480
tmp19025 := PrimCons(symunit, Nil)

tmp19026 := PrimCons(symlist, tmp19025)

tmp19027 := PrimCons(tmp19026, Nil)

tmp19028 := PrimCons(sym_1_1_6, tmp19027)

tmp19029 := PrimCons(symstring, tmp19028)

__e.TailApply(PrimFunc(symis_b), V6481, tmp19029, B6477, L6478, Key6479, C6480)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19030 := PrimValue(symshen_4_dsigf_d)

tmp19031 := Call(__e, PrimFunc(symshen_4assoc_1_6), symread_1file, tmp19024, tmp19030)


tmp19032 := PrimSet(symshen_4_dsigf_d, tmp19031)

_ = tmp19032

tmp19033 := MakeNative(func(__e *ControlFlow) {
V6486 := __e.Get(1)
_ = V6486
__e.Return(MakeNative(func(__e *ControlFlow) {
B6482 := __e.Get(1)
_ = B6482
__e.Return(MakeNative(func(__e *ControlFlow) {
L6483 := __e.Get(1)
_ = L6483
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6484 := __e.Get(1)
_ = Key6484
__e.Return(MakeNative(func(__e *ControlFlow) {
C6485 := __e.Get(1)
_ = C6485
tmp19034 := PrimCons(symunit, Nil)

tmp19035 := PrimCons(symlist, tmp19034)

tmp19036 := PrimCons(tmp19035, Nil)

tmp19037 := PrimCons(sym_1_1_6, tmp19036)

tmp19038 := PrimCons(symstring, tmp19037)

__e.TailApply(PrimFunc(symis_b), V6486, tmp19038, B6482, L6483, Key6484, C6485)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19039 := PrimValue(symshen_4_dsigf_d)

tmp19040 := Call(__e, PrimFunc(symshen_4assoc_1_6), symread_1from_1string, tmp19033, tmp19039)


tmp19041 := PrimSet(symshen_4_dsigf_d, tmp19040)

_ = tmp19041

tmp19042 := MakeNative(func(__e *ControlFlow) {
V6491 := __e.Get(1)
_ = V6491
__e.Return(MakeNative(func(__e *ControlFlow) {
B6487 := __e.Get(1)
_ = B6487
__e.Return(MakeNative(func(__e *ControlFlow) {
L6488 := __e.Get(1)
_ = L6488
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6489 := __e.Get(1)
_ = Key6489
__e.Return(MakeNative(func(__e *ControlFlow) {
C6490 := __e.Get(1)
_ = C6490
tmp19043 := PrimCons(symunit, Nil)

tmp19044 := PrimCons(symlist, tmp19043)

tmp19045 := PrimCons(tmp19044, Nil)

tmp19046 := PrimCons(sym_1_1_6, tmp19045)

tmp19047 := PrimCons(symstring, tmp19046)

__e.TailApply(PrimFunc(symis_b), V6491, tmp19047, B6487, L6488, Key6489, C6490)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19048 := PrimValue(symshen_4_dsigf_d)

tmp19049 := Call(__e, PrimFunc(symshen_4assoc_1_6), symread_1from_1string_1unprocessed, tmp19042, tmp19048)


tmp19050 := PrimSet(symshen_4_dsigf_d, tmp19049)

_ = tmp19050

tmp19051 := MakeNative(func(__e *ControlFlow) {
V6496 := __e.Get(1)
_ = V6496
__e.Return(MakeNative(func(__e *ControlFlow) {
B6492 := __e.Get(1)
_ = B6492
__e.Return(MakeNative(func(__e *ControlFlow) {
L6493 := __e.Get(1)
_ = L6493
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6494 := __e.Get(1)
_ = Key6494
__e.Return(MakeNative(func(__e *ControlFlow) {
C6495 := __e.Get(1)
_ = C6495
tmp19052 := PrimCons(symstring, Nil)

tmp19053 := PrimCons(sym_1_1_6, tmp19052)

__e.TailApply(PrimFunc(symis_b), V6496, tmp19053, B6492, L6493, Key6494, C6495)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19054 := PrimValue(symshen_4_dsigf_d)

tmp19055 := Call(__e, PrimFunc(symshen_4assoc_1_6), symrelease, tmp19051, tmp19054)


tmp19056 := PrimSet(symshen_4_dsigf_d, tmp19055)

_ = tmp19056

tmp19057 := MakeNative(func(__e *ControlFlow) {
V6501 := __e.Get(1)
_ = V6501
__e.Return(MakeNative(func(__e *ControlFlow) {
B6497 := __e.Get(1)
_ = B6497
__e.Return(MakeNative(func(__e *ControlFlow) {
L6498 := __e.Get(1)
_ = L6498
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6499 := __e.Get(1)
_ = Key6499
__e.Return(MakeNative(func(__e *ControlFlow) {
C6500 := __e.Get(1)
_ = C6500
tmp19058 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp19059 := PrimCons(A, Nil)

tmp19060 := PrimCons(symlist, tmp19059)

tmp19061 := PrimCons(A, Nil)

tmp19062 := PrimCons(symlist, tmp19061)

tmp19063 := PrimCons(tmp19062, Nil)

tmp19064 := PrimCons(sym_1_1_6, tmp19063)

tmp19065 := PrimCons(tmp19060, tmp19064)

tmp19066 := PrimCons(tmp19065, Nil)

tmp19067 := PrimCons(sym_1_1_6, tmp19066)

tmp19068 := PrimCons(A, tmp19067)

tmp19069 := Call(__e, PrimFunc(symis_b), V6501, tmp19068, B6497, L6498, Key6499, C6500)


__e.TailApply(PrimFunc(symshen_4gc), B6497, tmp19069)
return


}, 1)

tmp19070 := Call(__e, PrimFunc(symshen_4newpv), B6497)


__e.TailApply(tmp19058, tmp19070)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19071 := PrimValue(symshen_4_dsigf_d)

tmp19072 := Call(__e, PrimFunc(symshen_4assoc_1_6), symremove, tmp19057, tmp19071)


tmp19073 := PrimSet(symshen_4_dsigf_d, tmp19072)

_ = tmp19073

tmp19074 := MakeNative(func(__e *ControlFlow) {
V6506 := __e.Get(1)
_ = V6506
__e.Return(MakeNative(func(__e *ControlFlow) {
B6502 := __e.Get(1)
_ = B6502
__e.Return(MakeNative(func(__e *ControlFlow) {
L6503 := __e.Get(1)
_ = L6503
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6504 := __e.Get(1)
_ = Key6504
__e.Return(MakeNative(func(__e *ControlFlow) {
C6505 := __e.Get(1)
_ = C6505
tmp19075 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp19076 := PrimCons(A, Nil)

tmp19077 := PrimCons(symlist, tmp19076)

tmp19078 := PrimCons(A, Nil)

tmp19079 := PrimCons(symlist, tmp19078)

tmp19080 := PrimCons(tmp19079, Nil)

tmp19081 := PrimCons(sym_1_1_6, tmp19080)

tmp19082 := PrimCons(tmp19077, tmp19081)

tmp19083 := Call(__e, PrimFunc(symis_b), V6506, tmp19082, B6502, L6503, Key6504, C6505)


__e.TailApply(PrimFunc(symshen_4gc), B6502, tmp19083)
return


}, 1)

tmp19084 := Call(__e, PrimFunc(symshen_4newpv), B6502)


__e.TailApply(tmp19075, tmp19084)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19085 := PrimValue(symshen_4_dsigf_d)

tmp19086 := Call(__e, PrimFunc(symshen_4assoc_1_6), symreverse, tmp19074, tmp19085)


tmp19087 := PrimSet(symshen_4_dsigf_d, tmp19086)

_ = tmp19087

tmp19088 := MakeNative(func(__e *ControlFlow) {
V6511 := __e.Get(1)
_ = V6511
__e.Return(MakeNative(func(__e *ControlFlow) {
B6507 := __e.Get(1)
_ = B6507
__e.Return(MakeNative(func(__e *ControlFlow) {
L6508 := __e.Get(1)
_ = L6508
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6509 := __e.Get(1)
_ = Key6509
__e.Return(MakeNative(func(__e *ControlFlow) {
C6510 := __e.Get(1)
_ = C6510
tmp19089 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp19090 := PrimCons(A, Nil)

tmp19091 := PrimCons(sym_1_1_6, tmp19090)

tmp19092 := PrimCons(symstring, tmp19091)

tmp19093 := Call(__e, PrimFunc(symis_b), V6511, tmp19092, B6507, L6508, Key6509, C6510)


__e.TailApply(PrimFunc(symshen_4gc), B6507, tmp19093)
return


}, 1)

tmp19094 := Call(__e, PrimFunc(symshen_4newpv), B6507)


__e.TailApply(tmp19089, tmp19094)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19095 := PrimValue(symshen_4_dsigf_d)

tmp19096 := Call(__e, PrimFunc(symshen_4assoc_1_6), symsimple_1error, tmp19088, tmp19095)


tmp19097 := PrimSet(symshen_4_dsigf_d, tmp19096)

_ = tmp19097

tmp19098 := MakeNative(func(__e *ControlFlow) {
V6516 := __e.Get(1)
_ = V6516
__e.Return(MakeNative(func(__e *ControlFlow) {
B6512 := __e.Get(1)
_ = B6512
__e.Return(MakeNative(func(__e *ControlFlow) {
L6513 := __e.Get(1)
_ = L6513
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6514 := __e.Get(1)
_ = Key6514
__e.Return(MakeNative(func(__e *ControlFlow) {
C6515 := __e.Get(1)
_ = C6515
tmp19099 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp19100 := MakeNative(func(__e *ControlFlow) {
B := __e.Get(1)
_ = B
tmp19101 := PrimCons(B, Nil)

tmp19102 := PrimCons(sym_d, tmp19101)

tmp19103 := PrimCons(A, tmp19102)

tmp19104 := PrimCons(B, Nil)

tmp19105 := PrimCons(sym_1_1_6, tmp19104)

tmp19106 := PrimCons(tmp19103, tmp19105)

tmp19107 := Call(__e, PrimFunc(symis_b), V6516, tmp19106, B6512, L6513, Key6514, C6515)


__e.TailApply(PrimFunc(symshen_4gc), B6512, tmp19107)
return


}, 1)

tmp19108 := Call(__e, PrimFunc(symshen_4newpv), B6512)


tmp19109 := Call(__e, tmp19100, tmp19108)


__e.TailApply(PrimFunc(symshen_4gc), B6512, tmp19109)
return


}, 1)

tmp19110 := Call(__e, PrimFunc(symshen_4newpv), B6512)


__e.TailApply(tmp19099, tmp19110)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19111 := PrimValue(symshen_4_dsigf_d)

tmp19112 := Call(__e, PrimFunc(symshen_4assoc_1_6), symsnd, tmp19098, tmp19111)


tmp19113 := PrimSet(symshen_4_dsigf_d, tmp19112)

_ = tmp19113

tmp19114 := MakeNative(func(__e *ControlFlow) {
V6521 := __e.Get(1)
_ = V6521
__e.Return(MakeNative(func(__e *ControlFlow) {
B6517 := __e.Get(1)
_ = B6517
__e.Return(MakeNative(func(__e *ControlFlow) {
L6518 := __e.Get(1)
_ = L6518
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6519 := __e.Get(1)
_ = Key6519
__e.Return(MakeNative(func(__e *ControlFlow) {
C6520 := __e.Get(1)
_ = C6520
tmp19115 := PrimCons(symsymbol, Nil)

tmp19116 := PrimCons(sym_1_1_6, tmp19115)

tmp19117 := PrimCons(symnumber, tmp19116)

tmp19118 := PrimCons(tmp19117, Nil)

tmp19119 := PrimCons(sym_1_1_6, tmp19118)

tmp19120 := PrimCons(symsymbol, tmp19119)

__e.TailApply(PrimFunc(symis_b), V6521, tmp19120, B6517, L6518, Key6519, C6520)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19121 := PrimValue(symshen_4_dsigf_d)

tmp19122 := Call(__e, PrimFunc(symshen_4assoc_1_6), symspecialise, tmp19114, tmp19121)


tmp19123 := PrimSet(symshen_4_dsigf_d, tmp19122)

_ = tmp19123

tmp19124 := MakeNative(func(__e *ControlFlow) {
V6526 := __e.Get(1)
_ = V6526
__e.Return(MakeNative(func(__e *ControlFlow) {
B6522 := __e.Get(1)
_ = B6522
__e.Return(MakeNative(func(__e *ControlFlow) {
L6523 := __e.Get(1)
_ = L6523
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6524 := __e.Get(1)
_ = Key6524
__e.Return(MakeNative(func(__e *ControlFlow) {
C6525 := __e.Get(1)
_ = C6525
tmp19125 := PrimCons(symboolean, Nil)

tmp19126 := PrimCons(sym_1_1_6, tmp19125)

tmp19127 := PrimCons(symsymbol, tmp19126)

__e.TailApply(PrimFunc(symis_b), V6526, tmp19127, B6522, L6523, Key6524, C6525)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19128 := PrimValue(symshen_4_dsigf_d)

tmp19129 := Call(__e, PrimFunc(symshen_4assoc_1_6), symspy, tmp19124, tmp19128)


tmp19130 := PrimSet(symshen_4_dsigf_d, tmp19129)

_ = tmp19130

tmp19131 := MakeNative(func(__e *ControlFlow) {
V6531 := __e.Get(1)
_ = V6531
__e.Return(MakeNative(func(__e *ControlFlow) {
B6527 := __e.Get(1)
_ = B6527
__e.Return(MakeNative(func(__e *ControlFlow) {
L6528 := __e.Get(1)
_ = L6528
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6529 := __e.Get(1)
_ = Key6529
__e.Return(MakeNative(func(__e *ControlFlow) {
C6530 := __e.Get(1)
_ = C6530
tmp19132 := PrimCons(symboolean, Nil)

tmp19133 := PrimCons(sym_1_1_6, tmp19132)

__e.TailApply(PrimFunc(symis_b), V6531, tmp19133, B6527, L6528, Key6529, C6530)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19134 := PrimValue(symshen_4_dsigf_d)

tmp19135 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4spy_2, tmp19131, tmp19134)


tmp19136 := PrimSet(symshen_4_dsigf_d, tmp19135)

_ = tmp19136

tmp19137 := MakeNative(func(__e *ControlFlow) {
V6536 := __e.Get(1)
_ = V6536
__e.Return(MakeNative(func(__e *ControlFlow) {
B6532 := __e.Get(1)
_ = B6532
__e.Return(MakeNative(func(__e *ControlFlow) {
L6533 := __e.Get(1)
_ = L6533
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6534 := __e.Get(1)
_ = Key6534
__e.Return(MakeNative(func(__e *ControlFlow) {
C6535 := __e.Get(1)
_ = C6535
tmp19138 := PrimCons(symboolean, Nil)

tmp19139 := PrimCons(sym_1_1_6, tmp19138)

tmp19140 := PrimCons(symsymbol, tmp19139)

__e.TailApply(PrimFunc(symis_b), V6536, tmp19140, B6532, L6533, Key6534, C6535)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19141 := PrimValue(symshen_4_dsigf_d)

tmp19142 := Call(__e, PrimFunc(symshen_4assoc_1_6), symstep, tmp19137, tmp19141)


tmp19143 := PrimSet(symshen_4_dsigf_d, tmp19142)

_ = tmp19143

tmp19144 := MakeNative(func(__e *ControlFlow) {
V6541 := __e.Get(1)
_ = V6541
__e.Return(MakeNative(func(__e *ControlFlow) {
B6537 := __e.Get(1)
_ = B6537
__e.Return(MakeNative(func(__e *ControlFlow) {
L6538 := __e.Get(1)
_ = L6538
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6539 := __e.Get(1)
_ = Key6539
__e.Return(MakeNative(func(__e *ControlFlow) {
C6540 := __e.Get(1)
_ = C6540
tmp19145 := PrimCons(symboolean, Nil)

tmp19146 := PrimCons(sym_1_1_6, tmp19145)

__e.TailApply(PrimFunc(symis_b), V6541, tmp19146, B6537, L6538, Key6539, C6540)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19147 := PrimValue(symshen_4_dsigf_d)

tmp19148 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4step_2, tmp19144, tmp19147)


tmp19149 := PrimSet(symshen_4_dsigf_d, tmp19148)

_ = tmp19149

tmp19150 := MakeNative(func(__e *ControlFlow) {
V6546 := __e.Get(1)
_ = V6546
__e.Return(MakeNative(func(__e *ControlFlow) {
B6542 := __e.Get(1)
_ = B6542
__e.Return(MakeNative(func(__e *ControlFlow) {
L6543 := __e.Get(1)
_ = L6543
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6544 := __e.Get(1)
_ = Key6544
__e.Return(MakeNative(func(__e *ControlFlow) {
C6545 := __e.Get(1)
_ = C6545
tmp19151 := PrimCons(symin, Nil)

tmp19152 := PrimCons(symstream, tmp19151)

tmp19153 := PrimCons(tmp19152, Nil)

tmp19154 := PrimCons(sym_1_1_6, tmp19153)

__e.TailApply(PrimFunc(symis_b), V6546, tmp19154, B6542, L6543, Key6544, C6545)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19155 := PrimValue(symshen_4_dsigf_d)

tmp19156 := Call(__e, PrimFunc(symshen_4assoc_1_6), symstinput, tmp19150, tmp19155)


tmp19157 := PrimSet(symshen_4_dsigf_d, tmp19156)

_ = tmp19157

tmp19158 := MakeNative(func(__e *ControlFlow) {
V6551 := __e.Get(1)
_ = V6551
__e.Return(MakeNative(func(__e *ControlFlow) {
B6547 := __e.Get(1)
_ = B6547
__e.Return(MakeNative(func(__e *ControlFlow) {
L6548 := __e.Get(1)
_ = L6548
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6549 := __e.Get(1)
_ = Key6549
__e.Return(MakeNative(func(__e *ControlFlow) {
C6550 := __e.Get(1)
_ = C6550
tmp19159 := PrimCons(symout, Nil)

tmp19160 := PrimCons(symstream, tmp19159)

tmp19161 := PrimCons(tmp19160, Nil)

tmp19162 := PrimCons(sym_1_1_6, tmp19161)

__e.TailApply(PrimFunc(symis_b), V6551, tmp19162, B6547, L6548, Key6549, C6550)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19163 := PrimValue(symshen_4_dsigf_d)

tmp19164 := Call(__e, PrimFunc(symshen_4assoc_1_6), symsterror, tmp19158, tmp19163)


tmp19165 := PrimSet(symshen_4_dsigf_d, tmp19164)

_ = tmp19165

tmp19166 := MakeNative(func(__e *ControlFlow) {
V6556 := __e.Get(1)
_ = V6556
__e.Return(MakeNative(func(__e *ControlFlow) {
B6552 := __e.Get(1)
_ = B6552
__e.Return(MakeNative(func(__e *ControlFlow) {
L6553 := __e.Get(1)
_ = L6553
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6554 := __e.Get(1)
_ = Key6554
__e.Return(MakeNative(func(__e *ControlFlow) {
C6555 := __e.Get(1)
_ = C6555
tmp19167 := PrimCons(symout, Nil)

tmp19168 := PrimCons(symstream, tmp19167)

tmp19169 := PrimCons(tmp19168, Nil)

tmp19170 := PrimCons(sym_1_1_6, tmp19169)

__e.TailApply(PrimFunc(symis_b), V6556, tmp19170, B6552, L6553, Key6554, C6555)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19171 := PrimValue(symshen_4_dsigf_d)

tmp19172 := Call(__e, PrimFunc(symshen_4assoc_1_6), symstoutput, tmp19166, tmp19171)


tmp19173 := PrimSet(symshen_4_dsigf_d, tmp19172)

_ = tmp19173

tmp19174 := MakeNative(func(__e *ControlFlow) {
V6561 := __e.Get(1)
_ = V6561
__e.Return(MakeNative(func(__e *ControlFlow) {
B6557 := __e.Get(1)
_ = B6557
__e.Return(MakeNative(func(__e *ControlFlow) {
L6558 := __e.Get(1)
_ = L6558
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6559 := __e.Get(1)
_ = Key6559
__e.Return(MakeNative(func(__e *ControlFlow) {
C6560 := __e.Get(1)
_ = C6560
tmp19175 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp19176 := PrimCons(symboolean, Nil)

tmp19177 := PrimCons(sym_1_1_6, tmp19176)

tmp19178 := PrimCons(A, tmp19177)

tmp19179 := Call(__e, PrimFunc(symis_b), V6561, tmp19178, B6557, L6558, Key6559, C6560)


__e.TailApply(PrimFunc(symshen_4gc), B6557, tmp19179)
return


}, 1)

tmp19180 := Call(__e, PrimFunc(symshen_4newpv), B6557)


__e.TailApply(tmp19175, tmp19180)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19181 := PrimValue(symshen_4_dsigf_d)

tmp19182 := Call(__e, PrimFunc(symshen_4assoc_1_6), symstring_2, tmp19174, tmp19181)


tmp19183 := PrimSet(symshen_4_dsigf_d, tmp19182)

_ = tmp19183

tmp19184 := MakeNative(func(__e *ControlFlow) {
V6566 := __e.Get(1)
_ = V6566
__e.Return(MakeNative(func(__e *ControlFlow) {
B6562 := __e.Get(1)
_ = B6562
__e.Return(MakeNative(func(__e *ControlFlow) {
L6563 := __e.Get(1)
_ = L6563
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6564 := __e.Get(1)
_ = Key6564
__e.Return(MakeNative(func(__e *ControlFlow) {
C6565 := __e.Get(1)
_ = C6565
tmp19185 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp19186 := PrimCons(symstring, Nil)

tmp19187 := PrimCons(sym_1_1_6, tmp19186)

tmp19188 := PrimCons(A, tmp19187)

tmp19189 := Call(__e, PrimFunc(symis_b), V6566, tmp19188, B6562, L6563, Key6564, C6565)


__e.TailApply(PrimFunc(symshen_4gc), B6562, tmp19189)
return


}, 1)

tmp19190 := Call(__e, PrimFunc(symshen_4newpv), B6562)


__e.TailApply(tmp19185, tmp19190)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19191 := PrimValue(symshen_4_dsigf_d)

tmp19192 := Call(__e, PrimFunc(symshen_4assoc_1_6), symstr, tmp19184, tmp19191)


tmp19193 := PrimSet(symshen_4_dsigf_d, tmp19192)

_ = tmp19193

tmp19194 := MakeNative(func(__e *ControlFlow) {
V6571 := __e.Get(1)
_ = V6571
__e.Return(MakeNative(func(__e *ControlFlow) {
B6567 := __e.Get(1)
_ = B6567
__e.Return(MakeNative(func(__e *ControlFlow) {
L6568 := __e.Get(1)
_ = L6568
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6569 := __e.Get(1)
_ = Key6569
__e.Return(MakeNative(func(__e *ControlFlow) {
C6570 := __e.Get(1)
_ = C6570
tmp19195 := PrimCons(symnumber, Nil)

tmp19196 := PrimCons(sym_1_1_6, tmp19195)

tmp19197 := PrimCons(symstring, tmp19196)

__e.TailApply(PrimFunc(symis_b), V6571, tmp19197, B6567, L6568, Key6569, C6570)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19198 := PrimValue(symshen_4_dsigf_d)

tmp19199 := Call(__e, PrimFunc(symshen_4assoc_1_6), symstring_1_6n, tmp19194, tmp19198)


tmp19200 := PrimSet(symshen_4_dsigf_d, tmp19199)

_ = tmp19200

tmp19201 := MakeNative(func(__e *ControlFlow) {
V6576 := __e.Get(1)
_ = V6576
__e.Return(MakeNative(func(__e *ControlFlow) {
B6572 := __e.Get(1)
_ = B6572
__e.Return(MakeNative(func(__e *ControlFlow) {
L6573 := __e.Get(1)
_ = L6573
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6574 := __e.Get(1)
_ = Key6574
__e.Return(MakeNative(func(__e *ControlFlow) {
C6575 := __e.Get(1)
_ = C6575
tmp19202 := PrimCons(symsymbol, Nil)

tmp19203 := PrimCons(sym_1_1_6, tmp19202)

tmp19204 := PrimCons(symstring, tmp19203)

__e.TailApply(PrimFunc(symis_b), V6576, tmp19204, B6572, L6573, Key6574, C6575)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19205 := PrimValue(symshen_4_dsigf_d)

tmp19206 := Call(__e, PrimFunc(symshen_4assoc_1_6), symstring_1_6symbol, tmp19201, tmp19205)


tmp19207 := PrimSet(symshen_4_dsigf_d, tmp19206)

_ = tmp19207

tmp19208 := MakeNative(func(__e *ControlFlow) {
V6581 := __e.Get(1)
_ = V6581
__e.Return(MakeNative(func(__e *ControlFlow) {
B6577 := __e.Get(1)
_ = B6577
__e.Return(MakeNative(func(__e *ControlFlow) {
L6578 := __e.Get(1)
_ = L6578
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6579 := __e.Get(1)
_ = Key6579
__e.Return(MakeNative(func(__e *ControlFlow) {
C6580 := __e.Get(1)
_ = C6580
tmp19209 := PrimCons(symnumber, Nil)

tmp19210 := PrimCons(symlist, tmp19209)

tmp19211 := PrimCons(symnumber, Nil)

tmp19212 := PrimCons(sym_1_1_6, tmp19211)

tmp19213 := PrimCons(tmp19210, tmp19212)

__e.TailApply(PrimFunc(symis_b), V6581, tmp19213, B6577, L6578, Key6579, C6580)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19214 := PrimValue(symshen_4_dsigf_d)

tmp19215 := Call(__e, PrimFunc(symshen_4assoc_1_6), symsum, tmp19208, tmp19214)


tmp19216 := PrimSet(symshen_4_dsigf_d, tmp19215)

_ = tmp19216

tmp19217 := MakeNative(func(__e *ControlFlow) {
V6586 := __e.Get(1)
_ = V6586
__e.Return(MakeNative(func(__e *ControlFlow) {
B6582 := __e.Get(1)
_ = B6582
__e.Return(MakeNative(func(__e *ControlFlow) {
L6583 := __e.Get(1)
_ = L6583
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6584 := __e.Get(1)
_ = Key6584
__e.Return(MakeNative(func(__e *ControlFlow) {
C6585 := __e.Get(1)
_ = C6585
tmp19218 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp19219 := PrimCons(symboolean, Nil)

tmp19220 := PrimCons(sym_1_1_6, tmp19219)

tmp19221 := PrimCons(A, tmp19220)

tmp19222 := Call(__e, PrimFunc(symis_b), V6586, tmp19221, B6582, L6583, Key6584, C6585)


__e.TailApply(PrimFunc(symshen_4gc), B6582, tmp19222)
return


}, 1)

tmp19223 := Call(__e, PrimFunc(symshen_4newpv), B6582)


__e.TailApply(tmp19218, tmp19223)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19224 := PrimValue(symshen_4_dsigf_d)

tmp19225 := Call(__e, PrimFunc(symshen_4assoc_1_6), symsymbol_2, tmp19217, tmp19224)


tmp19226 := PrimSet(symshen_4_dsigf_d, tmp19225)

_ = tmp19226

tmp19227 := MakeNative(func(__e *ControlFlow) {
V6591 := __e.Get(1)
_ = V6591
__e.Return(MakeNative(func(__e *ControlFlow) {
B6587 := __e.Get(1)
_ = B6587
__e.Return(MakeNative(func(__e *ControlFlow) {
L6588 := __e.Get(1)
_ = L6588
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6589 := __e.Get(1)
_ = Key6589
__e.Return(MakeNative(func(__e *ControlFlow) {
C6590 := __e.Get(1)
_ = C6590
tmp19228 := PrimCons(symsymbol, Nil)

tmp19229 := PrimCons(sym_1_1_6, tmp19228)

tmp19230 := PrimCons(symsymbol, tmp19229)

__e.TailApply(PrimFunc(symis_b), V6591, tmp19230, B6587, L6588, Key6589, C6590)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19231 := PrimValue(symshen_4_dsigf_d)

tmp19232 := Call(__e, PrimFunc(symshen_4assoc_1_6), symsystemf, tmp19227, tmp19231)


tmp19233 := PrimSet(symshen_4_dsigf_d, tmp19232)

_ = tmp19233

tmp19234 := MakeNative(func(__e *ControlFlow) {
V6596 := __e.Get(1)
_ = V6596
__e.Return(MakeNative(func(__e *ControlFlow) {
B6592 := __e.Get(1)
_ = B6592
__e.Return(MakeNative(func(__e *ControlFlow) {
L6593 := __e.Get(1)
_ = L6593
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6594 := __e.Get(1)
_ = Key6594
__e.Return(MakeNative(func(__e *ControlFlow) {
C6595 := __e.Get(1)
_ = C6595
tmp19235 := PrimCons(symboolean, Nil)

tmp19236 := PrimCons(sym_1_1_6, tmp19235)

__e.TailApply(PrimFunc(symis_b), V6596, tmp19236, B6592, L6593, Key6594, C6595)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19237 := PrimValue(symshen_4_dsigf_d)

tmp19238 := Call(__e, PrimFunc(symshen_4assoc_1_6), symsystem_1S_2, tmp19234, tmp19237)


tmp19239 := PrimSet(symshen_4_dsigf_d, tmp19238)

_ = tmp19239

tmp19240 := MakeNative(func(__e *ControlFlow) {
V6601 := __e.Get(1)
_ = V6601
__e.Return(MakeNative(func(__e *ControlFlow) {
B6597 := __e.Get(1)
_ = B6597
__e.Return(MakeNative(func(__e *ControlFlow) {
L6598 := __e.Get(1)
_ = L6598
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6599 := __e.Get(1)
_ = Key6599
__e.Return(MakeNative(func(__e *ControlFlow) {
C6600 := __e.Get(1)
_ = C6600
tmp19241 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp19242 := PrimCons(A, Nil)

tmp19243 := PrimCons(symlist, tmp19242)

tmp19244 := PrimCons(A, Nil)

tmp19245 := PrimCons(symlist, tmp19244)

tmp19246 := PrimCons(tmp19245, Nil)

tmp19247 := PrimCons(sym_1_1_6, tmp19246)

tmp19248 := PrimCons(tmp19243, tmp19247)

tmp19249 := Call(__e, PrimFunc(symis_b), V6601, tmp19248, B6597, L6598, Key6599, C6600)


__e.TailApply(PrimFunc(symshen_4gc), B6597, tmp19249)
return


}, 1)

tmp19250 := Call(__e, PrimFunc(symshen_4newpv), B6597)


__e.TailApply(tmp19241, tmp19250)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19251 := PrimValue(symshen_4_dsigf_d)

tmp19252 := Call(__e, PrimFunc(symshen_4assoc_1_6), symtail, tmp19240, tmp19251)


tmp19253 := PrimSet(symshen_4_dsigf_d, tmp19252)

_ = tmp19253

tmp19254 := MakeNative(func(__e *ControlFlow) {
V6606 := __e.Get(1)
_ = V6606
__e.Return(MakeNative(func(__e *ControlFlow) {
B6602 := __e.Get(1)
_ = B6602
__e.Return(MakeNative(func(__e *ControlFlow) {
L6603 := __e.Get(1)
_ = L6603
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6604 := __e.Get(1)
_ = Key6604
__e.Return(MakeNative(func(__e *ControlFlow) {
C6605 := __e.Get(1)
_ = C6605
tmp19255 := PrimCons(symstring, Nil)

tmp19256 := PrimCons(sym_1_1_6, tmp19255)

tmp19257 := PrimCons(symstring, tmp19256)

__e.TailApply(PrimFunc(symis_b), V6606, tmp19257, B6602, L6603, Key6604, C6605)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19258 := PrimValue(symshen_4_dsigf_d)

tmp19259 := Call(__e, PrimFunc(symshen_4assoc_1_6), symtlstr, tmp19254, tmp19258)


tmp19260 := PrimSet(symshen_4_dsigf_d, tmp19259)

_ = tmp19260

tmp19261 := MakeNative(func(__e *ControlFlow) {
V6611 := __e.Get(1)
_ = V6611
__e.Return(MakeNative(func(__e *ControlFlow) {
B6607 := __e.Get(1)
_ = B6607
__e.Return(MakeNative(func(__e *ControlFlow) {
L6608 := __e.Get(1)
_ = L6608
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6609 := __e.Get(1)
_ = Key6609
__e.Return(MakeNative(func(__e *ControlFlow) {
C6610 := __e.Get(1)
_ = C6610
tmp19262 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp19263 := PrimCons(A, Nil)

tmp19264 := PrimCons(symvector, tmp19263)

tmp19265 := PrimCons(A, Nil)

tmp19266 := PrimCons(symvector, tmp19265)

tmp19267 := PrimCons(tmp19266, Nil)

tmp19268 := PrimCons(sym_1_1_6, tmp19267)

tmp19269 := PrimCons(tmp19264, tmp19268)

tmp19270 := Call(__e, PrimFunc(symis_b), V6611, tmp19269, B6607, L6608, Key6609, C6610)


__e.TailApply(PrimFunc(symshen_4gc), B6607, tmp19270)
return


}, 1)

tmp19271 := Call(__e, PrimFunc(symshen_4newpv), B6607)


__e.TailApply(tmp19262, tmp19271)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19272 := PrimValue(symshen_4_dsigf_d)

tmp19273 := Call(__e, PrimFunc(symshen_4assoc_1_6), symtlv, tmp19261, tmp19272)


tmp19274 := PrimSet(symshen_4_dsigf_d, tmp19273)

_ = tmp19274

tmp19275 := MakeNative(func(__e *ControlFlow) {
V6616 := __e.Get(1)
_ = V6616
__e.Return(MakeNative(func(__e *ControlFlow) {
B6612 := __e.Get(1)
_ = B6612
__e.Return(MakeNative(func(__e *ControlFlow) {
L6613 := __e.Get(1)
_ = L6613
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6614 := __e.Get(1)
_ = Key6614
__e.Return(MakeNative(func(__e *ControlFlow) {
C6615 := __e.Get(1)
_ = C6615
tmp19276 := PrimCons(symboolean, Nil)

tmp19277 := PrimCons(sym_1_1_6, tmp19276)

tmp19278 := PrimCons(symsymbol, tmp19277)

__e.TailApply(PrimFunc(symis_b), V6616, tmp19278, B6612, L6613, Key6614, C6615)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19279 := PrimValue(symshen_4_dsigf_d)

tmp19280 := Call(__e, PrimFunc(symshen_4assoc_1_6), symtc, tmp19275, tmp19279)


tmp19281 := PrimSet(symshen_4_dsigf_d, tmp19280)

_ = tmp19281

tmp19282 := MakeNative(func(__e *ControlFlow) {
V6621 := __e.Get(1)
_ = V6621
__e.Return(MakeNative(func(__e *ControlFlow) {
B6617 := __e.Get(1)
_ = B6617
__e.Return(MakeNative(func(__e *ControlFlow) {
L6618 := __e.Get(1)
_ = L6618
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6619 := __e.Get(1)
_ = Key6619
__e.Return(MakeNative(func(__e *ControlFlow) {
C6620 := __e.Get(1)
_ = C6620
tmp19283 := PrimCons(symboolean, Nil)

tmp19284 := PrimCons(sym_1_1_6, tmp19283)

__e.TailApply(PrimFunc(symis_b), V6621, tmp19284, B6617, L6618, Key6619, C6620)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19285 := PrimValue(symshen_4_dsigf_d)

tmp19286 := Call(__e, PrimFunc(symshen_4assoc_1_6), symtc_2, tmp19282, tmp19285)


tmp19287 := PrimSet(symshen_4_dsigf_d, tmp19286)

_ = tmp19287

tmp19288 := MakeNative(func(__e *ControlFlow) {
V6626 := __e.Get(1)
_ = V6626
__e.Return(MakeNative(func(__e *ControlFlow) {
B6622 := __e.Get(1)
_ = B6622
__e.Return(MakeNative(func(__e *ControlFlow) {
L6623 := __e.Get(1)
_ = L6623
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6624 := __e.Get(1)
_ = Key6624
__e.Return(MakeNative(func(__e *ControlFlow) {
C6625 := __e.Get(1)
_ = C6625
tmp19289 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp19290 := PrimCons(A, Nil)

tmp19291 := PrimCons(symlazy, tmp19290)

tmp19292 := PrimCons(A, Nil)

tmp19293 := PrimCons(sym_1_1_6, tmp19292)

tmp19294 := PrimCons(tmp19291, tmp19293)

tmp19295 := Call(__e, PrimFunc(symis_b), V6626, tmp19294, B6622, L6623, Key6624, C6625)


__e.TailApply(PrimFunc(symshen_4gc), B6622, tmp19295)
return


}, 1)

tmp19296 := Call(__e, PrimFunc(symshen_4newpv), B6622)


__e.TailApply(tmp19289, tmp19296)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19297 := PrimValue(symshen_4_dsigf_d)

tmp19298 := Call(__e, PrimFunc(symshen_4assoc_1_6), symthaw, tmp19288, tmp19297)


tmp19299 := PrimSet(symshen_4_dsigf_d, tmp19298)

_ = tmp19299

tmp19300 := MakeNative(func(__e *ControlFlow) {
V6631 := __e.Get(1)
_ = V6631
__e.Return(MakeNative(func(__e *ControlFlow) {
B6627 := __e.Get(1)
_ = B6627
__e.Return(MakeNative(func(__e *ControlFlow) {
L6628 := __e.Get(1)
_ = L6628
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6629 := __e.Get(1)
_ = Key6629
__e.Return(MakeNative(func(__e *ControlFlow) {
C6630 := __e.Get(1)
_ = C6630
tmp19301 := PrimCons(symsymbol, Nil)

tmp19302 := PrimCons(sym_1_1_6, tmp19301)

tmp19303 := PrimCons(symsymbol, tmp19302)

__e.TailApply(PrimFunc(symis_b), V6631, tmp19303, B6627, L6628, Key6629, C6630)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19304 := PrimValue(symshen_4_dsigf_d)

tmp19305 := Call(__e, PrimFunc(symshen_4assoc_1_6), symtrack, tmp19300, tmp19304)


tmp19306 := PrimSet(symshen_4_dsigf_d, tmp19305)

_ = tmp19306

tmp19307 := MakeNative(func(__e *ControlFlow) {
V6636 := __e.Get(1)
_ = V6636
__e.Return(MakeNative(func(__e *ControlFlow) {
B6632 := __e.Get(1)
_ = B6632
__e.Return(MakeNative(func(__e *ControlFlow) {
L6633 := __e.Get(1)
_ = L6633
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6634 := __e.Get(1)
_ = Key6634
__e.Return(MakeNative(func(__e *ControlFlow) {
C6635 := __e.Get(1)
_ = C6635
tmp19308 := PrimCons(symsymbol, Nil)

tmp19309 := PrimCons(symlist, tmp19308)

tmp19310 := PrimCons(tmp19309, Nil)

tmp19311 := PrimCons(sym_1_1_6, tmp19310)

__e.TailApply(PrimFunc(symis_b), V6636, tmp19311, B6632, L6633, Key6634, C6635)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19312 := PrimValue(symshen_4_dsigf_d)

tmp19313 := Call(__e, PrimFunc(symshen_4assoc_1_6), symtracked, tmp19307, tmp19312)


tmp19314 := PrimSet(symshen_4_dsigf_d, tmp19313)

_ = tmp19314

tmp19315 := MakeNative(func(__e *ControlFlow) {
V6641 := __e.Get(1)
_ = V6641
__e.Return(MakeNative(func(__e *ControlFlow) {
B6637 := __e.Get(1)
_ = B6637
__e.Return(MakeNative(func(__e *ControlFlow) {
L6638 := __e.Get(1)
_ = L6638
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6639 := __e.Get(1)
_ = Key6639
__e.Return(MakeNative(func(__e *ControlFlow) {
C6640 := __e.Get(1)
_ = C6640
tmp19316 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp19317 := PrimCons(A, Nil)

tmp19318 := PrimCons(sym_1_1_6, tmp19317)

tmp19319 := PrimCons(symexception, tmp19318)

tmp19320 := PrimCons(A, Nil)

tmp19321 := PrimCons(sym_1_1_6, tmp19320)

tmp19322 := PrimCons(tmp19319, tmp19321)

tmp19323 := PrimCons(tmp19322, Nil)

tmp19324 := PrimCons(sym_1_1_6, tmp19323)

tmp19325 := PrimCons(A, tmp19324)

tmp19326 := Call(__e, PrimFunc(symis_b), V6641, tmp19325, B6637, L6638, Key6639, C6640)


__e.TailApply(PrimFunc(symshen_4gc), B6637, tmp19326)
return


}, 1)

tmp19327 := Call(__e, PrimFunc(symshen_4newpv), B6637)


__e.TailApply(tmp19316, tmp19327)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19328 := PrimValue(symshen_4_dsigf_d)

tmp19329 := Call(__e, PrimFunc(symshen_4assoc_1_6), symtrap_1error, tmp19315, tmp19328)


tmp19330 := PrimSet(symshen_4_dsigf_d, tmp19329)

_ = tmp19330

tmp19331 := MakeNative(func(__e *ControlFlow) {
V6646 := __e.Get(1)
_ = V6646
__e.Return(MakeNative(func(__e *ControlFlow) {
B6642 := __e.Get(1)
_ = B6642
__e.Return(MakeNative(func(__e *ControlFlow) {
L6643 := __e.Get(1)
_ = L6643
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6644 := __e.Get(1)
_ = Key6644
__e.Return(MakeNative(func(__e *ControlFlow) {
C6645 := __e.Get(1)
_ = C6645
tmp19332 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp19333 := PrimCons(symboolean, Nil)

tmp19334 := PrimCons(sym_1_1_6, tmp19333)

tmp19335 := PrimCons(A, tmp19334)

tmp19336 := Call(__e, PrimFunc(symis_b), V6646, tmp19335, B6642, L6643, Key6644, C6645)


__e.TailApply(PrimFunc(symshen_4gc), B6642, tmp19336)
return


}, 1)

tmp19337 := Call(__e, PrimFunc(symshen_4newpv), B6642)


__e.TailApply(tmp19332, tmp19337)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19338 := PrimValue(symshen_4_dsigf_d)

tmp19339 := Call(__e, PrimFunc(symshen_4assoc_1_6), symtuple_2, tmp19331, tmp19338)


tmp19340 := PrimSet(symshen_4_dsigf_d, tmp19339)

_ = tmp19340

tmp19341 := MakeNative(func(__e *ControlFlow) {
V6651 := __e.Get(1)
_ = V6651
__e.Return(MakeNative(func(__e *ControlFlow) {
B6647 := __e.Get(1)
_ = B6647
__e.Return(MakeNative(func(__e *ControlFlow) {
L6648 := __e.Get(1)
_ = L6648
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6649 := __e.Get(1)
_ = Key6649
__e.Return(MakeNative(func(__e *ControlFlow) {
C6650 := __e.Get(1)
_ = C6650
tmp19342 := PrimCons(symstring, Nil)

tmp19343 := PrimCons(symlist, tmp19342)

tmp19344 := PrimCons(tmp19343, Nil)

tmp19345 := PrimCons(sym_1_1_6, tmp19344)

tmp19346 := PrimCons(symstring, tmp19345)

__e.TailApply(PrimFunc(symis_b), V6651, tmp19346, B6647, L6648, Key6649, C6650)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19347 := PrimValue(symshen_4_dsigf_d)

tmp19348 := Call(__e, PrimFunc(symshen_4assoc_1_6), symunabsolute, tmp19341, tmp19347)


tmp19349 := PrimSet(symshen_4_dsigf_d, tmp19348)

_ = tmp19349

tmp19350 := MakeNative(func(__e *ControlFlow) {
V6656 := __e.Get(1)
_ = V6656
__e.Return(MakeNative(func(__e *ControlFlow) {
B6652 := __e.Get(1)
_ = B6652
__e.Return(MakeNative(func(__e *ControlFlow) {
L6653 := __e.Get(1)
_ = L6653
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6654 := __e.Get(1)
_ = Key6654
__e.Return(MakeNative(func(__e *ControlFlow) {
C6655 := __e.Get(1)
_ = C6655
tmp19351 := PrimCons(symsymbol, Nil)

tmp19352 := PrimCons(sym_1_1_6, tmp19351)

tmp19353 := PrimCons(symsymbol, tmp19352)

__e.TailApply(PrimFunc(symis_b), V6656, tmp19353, B6652, L6653, Key6654, C6655)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19354 := PrimValue(symshen_4_dsigf_d)

tmp19355 := Call(__e, PrimFunc(symshen_4assoc_1_6), symundefmacro, tmp19350, tmp19354)


tmp19356 := PrimSet(symshen_4_dsigf_d, tmp19355)

_ = tmp19356

tmp19357 := MakeNative(func(__e *ControlFlow) {
V6661 := __e.Get(1)
_ = V6661
__e.Return(MakeNative(func(__e *ControlFlow) {
B6657 := __e.Get(1)
_ = B6657
__e.Return(MakeNative(func(__e *ControlFlow) {
L6658 := __e.Get(1)
_ = L6658
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6659 := __e.Get(1)
_ = Key6659
__e.Return(MakeNative(func(__e *ControlFlow) {
C6660 := __e.Get(1)
_ = C6660
tmp19358 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp19359 := PrimCons(A, Nil)

tmp19360 := PrimCons(symlist, tmp19359)

tmp19361 := PrimCons(A, Nil)

tmp19362 := PrimCons(symlist, tmp19361)

tmp19363 := PrimCons(A, Nil)

tmp19364 := PrimCons(symlist, tmp19363)

tmp19365 := PrimCons(tmp19364, Nil)

tmp19366 := PrimCons(sym_1_1_6, tmp19365)

tmp19367 := PrimCons(tmp19362, tmp19366)

tmp19368 := PrimCons(tmp19367, Nil)

tmp19369 := PrimCons(sym_1_1_6, tmp19368)

tmp19370 := PrimCons(tmp19360, tmp19369)

tmp19371 := Call(__e, PrimFunc(symis_b), V6661, tmp19370, B6657, L6658, Key6659, C6660)


__e.TailApply(PrimFunc(symshen_4gc), B6657, tmp19371)
return


}, 1)

tmp19372 := Call(__e, PrimFunc(symshen_4newpv), B6657)


__e.TailApply(tmp19358, tmp19372)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19373 := PrimValue(symshen_4_dsigf_d)

tmp19374 := Call(__e, PrimFunc(symshen_4assoc_1_6), symunion, tmp19357, tmp19373)


tmp19375 := PrimSet(symshen_4_dsigf_d, tmp19374)

_ = tmp19375

tmp19376 := MakeNative(func(__e *ControlFlow) {
V6666 := __e.Get(1)
_ = V6666
__e.Return(MakeNative(func(__e *ControlFlow) {
B6662 := __e.Get(1)
_ = B6662
__e.Return(MakeNative(func(__e *ControlFlow) {
L6663 := __e.Get(1)
_ = L6663
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6664 := __e.Get(1)
_ = Key6664
__e.Return(MakeNative(func(__e *ControlFlow) {
C6665 := __e.Get(1)
_ = C6665
tmp19377 := PrimCons(symsymbol, Nil)

tmp19378 := PrimCons(sym_1_1_6, tmp19377)

tmp19379 := PrimCons(symsymbol, tmp19378)

__e.TailApply(PrimFunc(symis_b), V6666, tmp19379, B6662, L6663, Key6664, C6665)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19380 := PrimValue(symshen_4_dsigf_d)

tmp19381 := Call(__e, PrimFunc(symshen_4assoc_1_6), symunprofile, tmp19376, tmp19380)


tmp19382 := PrimSet(symshen_4_dsigf_d, tmp19381)

_ = tmp19382

tmp19383 := MakeNative(func(__e *ControlFlow) {
V6671 := __e.Get(1)
_ = V6671
__e.Return(MakeNative(func(__e *ControlFlow) {
B6667 := __e.Get(1)
_ = B6667
__e.Return(MakeNative(func(__e *ControlFlow) {
L6668 := __e.Get(1)
_ = L6668
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6669 := __e.Get(1)
_ = Key6669
__e.Return(MakeNative(func(__e *ControlFlow) {
C6670 := __e.Get(1)
_ = C6670
tmp19384 := PrimCons(symsymbol, Nil)

tmp19385 := PrimCons(sym_1_1_6, tmp19384)

tmp19386 := PrimCons(symsymbol, tmp19385)

__e.TailApply(PrimFunc(symis_b), V6671, tmp19386, B6667, L6668, Key6669, C6670)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19387 := PrimValue(symshen_4_dsigf_d)

tmp19388 := Call(__e, PrimFunc(symshen_4assoc_1_6), symuntrack, tmp19383, tmp19387)


tmp19389 := PrimSet(symshen_4_dsigf_d, tmp19388)

_ = tmp19389

tmp19390 := MakeNative(func(__e *ControlFlow) {
V6676 := __e.Get(1)
_ = V6676
__e.Return(MakeNative(func(__e *ControlFlow) {
B6672 := __e.Get(1)
_ = B6672
__e.Return(MakeNative(func(__e *ControlFlow) {
L6673 := __e.Get(1)
_ = L6673
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6674 := __e.Get(1)
_ = Key6674
__e.Return(MakeNative(func(__e *ControlFlow) {
C6675 := __e.Get(1)
_ = C6675
tmp19391 := PrimCons(symsymbol, Nil)

tmp19392 := PrimCons(symlist, tmp19391)

tmp19393 := PrimCons(tmp19392, Nil)

tmp19394 := PrimCons(sym_1_1_6, tmp19393)

__e.TailApply(PrimFunc(symis_b), V6676, tmp19394, B6672, L6673, Key6674, C6675)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19395 := PrimValue(symshen_4_dsigf_d)

tmp19396 := Call(__e, PrimFunc(symshen_4assoc_1_6), symuserdefs, tmp19390, tmp19395)


tmp19397 := PrimSet(symshen_4_dsigf_d, tmp19396)

_ = tmp19397

tmp19398 := MakeNative(func(__e *ControlFlow) {
V6681 := __e.Get(1)
_ = V6681
__e.Return(MakeNative(func(__e *ControlFlow) {
B6677 := __e.Get(1)
_ = B6677
__e.Return(MakeNative(func(__e *ControlFlow) {
L6678 := __e.Get(1)
_ = L6678
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6679 := __e.Get(1)
_ = Key6679
__e.Return(MakeNative(func(__e *ControlFlow) {
C6680 := __e.Get(1)
_ = C6680
tmp19399 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp19400 := PrimCons(symboolean, Nil)

tmp19401 := PrimCons(sym_1_1_6, tmp19400)

tmp19402 := PrimCons(A, tmp19401)

tmp19403 := Call(__e, PrimFunc(symis_b), V6681, tmp19402, B6677, L6678, Key6679, C6680)


__e.TailApply(PrimFunc(symshen_4gc), B6677, tmp19403)
return


}, 1)

tmp19404 := Call(__e, PrimFunc(symshen_4newpv), B6677)


__e.TailApply(tmp19399, tmp19404)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19405 := PrimValue(symshen_4_dsigf_d)

tmp19406 := Call(__e, PrimFunc(symshen_4assoc_1_6), symvariable_2, tmp19398, tmp19405)


tmp19407 := PrimSet(symshen_4_dsigf_d, tmp19406)

_ = tmp19407

tmp19408 := MakeNative(func(__e *ControlFlow) {
V6686 := __e.Get(1)
_ = V6686
__e.Return(MakeNative(func(__e *ControlFlow) {
B6682 := __e.Get(1)
_ = B6682
__e.Return(MakeNative(func(__e *ControlFlow) {
L6683 := __e.Get(1)
_ = L6683
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6684 := __e.Get(1)
_ = Key6684
__e.Return(MakeNative(func(__e *ControlFlow) {
C6685 := __e.Get(1)
_ = C6685
tmp19409 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp19410 := PrimCons(symboolean, Nil)

tmp19411 := PrimCons(sym_1_1_6, tmp19410)

tmp19412 := PrimCons(A, tmp19411)

tmp19413 := Call(__e, PrimFunc(symis_b), V6686, tmp19412, B6682, L6683, Key6684, C6685)


__e.TailApply(PrimFunc(symshen_4gc), B6682, tmp19413)
return


}, 1)

tmp19414 := Call(__e, PrimFunc(symshen_4newpv), B6682)


__e.TailApply(tmp19409, tmp19414)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19415 := PrimValue(symshen_4_dsigf_d)

tmp19416 := Call(__e, PrimFunc(symshen_4assoc_1_6), symvector_2, tmp19408, tmp19415)


tmp19417 := PrimSet(symshen_4_dsigf_d, tmp19416)

_ = tmp19417

tmp19418 := MakeNative(func(__e *ControlFlow) {
V6691 := __e.Get(1)
_ = V6691
__e.Return(MakeNative(func(__e *ControlFlow) {
B6687 := __e.Get(1)
_ = B6687
__e.Return(MakeNative(func(__e *ControlFlow) {
L6688 := __e.Get(1)
_ = L6688
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6689 := __e.Get(1)
_ = Key6689
__e.Return(MakeNative(func(__e *ControlFlow) {
C6690 := __e.Get(1)
_ = C6690
tmp19419 := PrimCons(symstring, Nil)

tmp19420 := PrimCons(sym_1_1_6, tmp19419)

__e.TailApply(PrimFunc(symis_b), V6691, tmp19420, B6687, L6688, Key6689, C6690)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19421 := PrimValue(symshen_4_dsigf_d)

tmp19422 := Call(__e, PrimFunc(symshen_4assoc_1_6), symversion, tmp19418, tmp19421)


tmp19423 := PrimSet(symshen_4_dsigf_d, tmp19422)

_ = tmp19423

tmp19424 := MakeNative(func(__e *ControlFlow) {
V6696 := __e.Get(1)
_ = V6696
__e.Return(MakeNative(func(__e *ControlFlow) {
B6692 := __e.Get(1)
_ = B6692
__e.Return(MakeNative(func(__e *ControlFlow) {
L6693 := __e.Get(1)
_ = L6693
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6694 := __e.Get(1)
_ = Key6694
__e.Return(MakeNative(func(__e *ControlFlow) {
C6695 := __e.Get(1)
_ = C6695
tmp19425 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp19426 := PrimCons(A, Nil)

tmp19427 := PrimCons(sym_1_1_6, tmp19426)

tmp19428 := PrimCons(A, tmp19427)

tmp19429 := PrimCons(tmp19428, Nil)

tmp19430 := PrimCons(sym_1_1_6, tmp19429)

tmp19431 := PrimCons(symstring, tmp19430)

tmp19432 := Call(__e, PrimFunc(symis_b), V6696, tmp19431, B6692, L6693, Key6694, C6695)


__e.TailApply(PrimFunc(symshen_4gc), B6692, tmp19432)
return


}, 1)

tmp19433 := Call(__e, PrimFunc(symshen_4newpv), B6692)


__e.TailApply(tmp19425, tmp19433)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19434 := PrimValue(symshen_4_dsigf_d)

tmp19435 := Call(__e, PrimFunc(symshen_4assoc_1_6), symwrite_1to_1file, tmp19424, tmp19434)


tmp19436 := PrimSet(symshen_4_dsigf_d, tmp19435)

_ = tmp19436

tmp19437 := MakeNative(func(__e *ControlFlow) {
V6701 := __e.Get(1)
_ = V6701
__e.Return(MakeNative(func(__e *ControlFlow) {
B6697 := __e.Get(1)
_ = B6697
__e.Return(MakeNative(func(__e *ControlFlow) {
L6698 := __e.Get(1)
_ = L6698
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6699 := __e.Get(1)
_ = Key6699
__e.Return(MakeNative(func(__e *ControlFlow) {
C6700 := __e.Get(1)
_ = C6700
tmp19438 := PrimCons(symout, Nil)

tmp19439 := PrimCons(symstream, tmp19438)

tmp19440 := PrimCons(symnumber, Nil)

tmp19441 := PrimCons(sym_1_1_6, tmp19440)

tmp19442 := PrimCons(tmp19439, tmp19441)

tmp19443 := PrimCons(tmp19442, Nil)

tmp19444 := PrimCons(sym_1_1_6, tmp19443)

tmp19445 := PrimCons(symnumber, tmp19444)

__e.TailApply(PrimFunc(symis_b), V6701, tmp19445, B6697, L6698, Key6699, C6700)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19446 := PrimValue(symshen_4_dsigf_d)

tmp19447 := Call(__e, PrimFunc(symshen_4assoc_1_6), symwrite_1byte, tmp19437, tmp19446)


tmp19448 := PrimSet(symshen_4_dsigf_d, tmp19447)

_ = tmp19448

tmp19449 := MakeNative(func(__e *ControlFlow) {
V6706 := __e.Get(1)
_ = V6706
__e.Return(MakeNative(func(__e *ControlFlow) {
B6702 := __e.Get(1)
_ = B6702
__e.Return(MakeNative(func(__e *ControlFlow) {
L6703 := __e.Get(1)
_ = L6703
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6704 := __e.Get(1)
_ = Key6704
__e.Return(MakeNative(func(__e *ControlFlow) {
C6705 := __e.Get(1)
_ = C6705
tmp19450 := PrimCons(symboolean, Nil)

tmp19451 := PrimCons(sym_1_1_6, tmp19450)

tmp19452 := PrimCons(symstring, tmp19451)

__e.TailApply(PrimFunc(symis_b), V6706, tmp19452, B6702, L6703, Key6704, C6705)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19453 := PrimValue(symshen_4_dsigf_d)

tmp19454 := Call(__e, PrimFunc(symshen_4assoc_1_6), symy_1or_1n_2, tmp19449, tmp19453)


tmp19455 := PrimSet(symshen_4_dsigf_d, tmp19454)

_ = tmp19455

tmp19456 := MakeNative(func(__e *ControlFlow) {
V6711 := __e.Get(1)
_ = V6711
__e.Return(MakeNative(func(__e *ControlFlow) {
B6707 := __e.Get(1)
_ = B6707
__e.Return(MakeNative(func(__e *ControlFlow) {
L6708 := __e.Get(1)
_ = L6708
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6709 := __e.Get(1)
_ = Key6709
__e.Return(MakeNative(func(__e *ControlFlow) {
C6710 := __e.Get(1)
_ = C6710
tmp19457 := PrimCons(symboolean, Nil)

tmp19458 := PrimCons(sym_1_1_6, tmp19457)

tmp19459 := PrimCons(symnumber, tmp19458)

tmp19460 := PrimCons(tmp19459, Nil)

tmp19461 := PrimCons(sym_1_1_6, tmp19460)

tmp19462 := PrimCons(symnumber, tmp19461)

__e.TailApply(PrimFunc(symis_b), V6711, tmp19462, B6707, L6708, Key6709, C6710)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19463 := PrimValue(symshen_4_dsigf_d)

tmp19464 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_6, tmp19456, tmp19463)


tmp19465 := PrimSet(symshen_4_dsigf_d, tmp19464)

_ = tmp19465

tmp19466 := MakeNative(func(__e *ControlFlow) {
V6716 := __e.Get(1)
_ = V6716
__e.Return(MakeNative(func(__e *ControlFlow) {
B6712 := __e.Get(1)
_ = B6712
__e.Return(MakeNative(func(__e *ControlFlow) {
L6713 := __e.Get(1)
_ = L6713
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6714 := __e.Get(1)
_ = Key6714
__e.Return(MakeNative(func(__e *ControlFlow) {
C6715 := __e.Get(1)
_ = C6715
tmp19467 := PrimCons(symboolean, Nil)

tmp19468 := PrimCons(sym_1_1_6, tmp19467)

tmp19469 := PrimCons(symnumber, tmp19468)

tmp19470 := PrimCons(tmp19469, Nil)

tmp19471 := PrimCons(sym_1_1_6, tmp19470)

tmp19472 := PrimCons(symnumber, tmp19471)

__e.TailApply(PrimFunc(symis_b), V6716, tmp19472, B6712, L6713, Key6714, C6715)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19473 := PrimValue(symshen_4_dsigf_d)

tmp19474 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_5, tmp19466, tmp19473)


tmp19475 := PrimSet(symshen_4_dsigf_d, tmp19474)

_ = tmp19475

tmp19476 := MakeNative(func(__e *ControlFlow) {
V6721 := __e.Get(1)
_ = V6721
__e.Return(MakeNative(func(__e *ControlFlow) {
B6717 := __e.Get(1)
_ = B6717
__e.Return(MakeNative(func(__e *ControlFlow) {
L6718 := __e.Get(1)
_ = L6718
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6719 := __e.Get(1)
_ = Key6719
__e.Return(MakeNative(func(__e *ControlFlow) {
C6720 := __e.Get(1)
_ = C6720
tmp19477 := PrimCons(symboolean, Nil)

tmp19478 := PrimCons(sym_1_1_6, tmp19477)

tmp19479 := PrimCons(symnumber, tmp19478)

tmp19480 := PrimCons(tmp19479, Nil)

tmp19481 := PrimCons(sym_1_1_6, tmp19480)

tmp19482 := PrimCons(symnumber, tmp19481)

__e.TailApply(PrimFunc(symis_b), V6721, tmp19482, B6717, L6718, Key6719, C6720)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19483 := PrimValue(symshen_4_dsigf_d)

tmp19484 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_6_a, tmp19476, tmp19483)


tmp19485 := PrimSet(symshen_4_dsigf_d, tmp19484)

_ = tmp19485

tmp19486 := MakeNative(func(__e *ControlFlow) {
V6726 := __e.Get(1)
_ = V6726
__e.Return(MakeNative(func(__e *ControlFlow) {
B6722 := __e.Get(1)
_ = B6722
__e.Return(MakeNative(func(__e *ControlFlow) {
L6723 := __e.Get(1)
_ = L6723
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6724 := __e.Get(1)
_ = Key6724
__e.Return(MakeNative(func(__e *ControlFlow) {
C6725 := __e.Get(1)
_ = C6725
tmp19487 := PrimCons(symboolean, Nil)

tmp19488 := PrimCons(sym_1_1_6, tmp19487)

tmp19489 := PrimCons(symnumber, tmp19488)

tmp19490 := PrimCons(tmp19489, Nil)

tmp19491 := PrimCons(sym_1_1_6, tmp19490)

tmp19492 := PrimCons(symnumber, tmp19491)

__e.TailApply(PrimFunc(symis_b), V6726, tmp19492, B6722, L6723, Key6724, C6725)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19493 := PrimValue(symshen_4_dsigf_d)

tmp19494 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_5_a, tmp19486, tmp19493)


tmp19495 := PrimSet(symshen_4_dsigf_d, tmp19494)

_ = tmp19495

tmp19496 := MakeNative(func(__e *ControlFlow) {
V6731 := __e.Get(1)
_ = V6731
__e.Return(MakeNative(func(__e *ControlFlow) {
B6727 := __e.Get(1)
_ = B6727
__e.Return(MakeNative(func(__e *ControlFlow) {
L6728 := __e.Get(1)
_ = L6728
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6729 := __e.Get(1)
_ = Key6729
__e.Return(MakeNative(func(__e *ControlFlow) {
C6730 := __e.Get(1)
_ = C6730
tmp19497 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp19498 := PrimCons(symboolean, Nil)

tmp19499 := PrimCons(sym_1_1_6, tmp19498)

tmp19500 := PrimCons(A, tmp19499)

tmp19501 := PrimCons(tmp19500, Nil)

tmp19502 := PrimCons(sym_1_1_6, tmp19501)

tmp19503 := PrimCons(A, tmp19502)

tmp19504 := Call(__e, PrimFunc(symis_b), V6731, tmp19503, B6727, L6728, Key6729, C6730)


__e.TailApply(PrimFunc(symshen_4gc), B6727, tmp19504)
return


}, 1)

tmp19505 := Call(__e, PrimFunc(symshen_4newpv), B6727)


__e.TailApply(tmp19497, tmp19505)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19506 := PrimValue(symshen_4_dsigf_d)

tmp19507 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_a, tmp19496, tmp19506)


tmp19508 := PrimSet(symshen_4_dsigf_d, tmp19507)

_ = tmp19508

tmp19509 := MakeNative(func(__e *ControlFlow) {
V6736 := __e.Get(1)
_ = V6736
__e.Return(MakeNative(func(__e *ControlFlow) {
B6732 := __e.Get(1)
_ = B6732
__e.Return(MakeNative(func(__e *ControlFlow) {
L6733 := __e.Get(1)
_ = L6733
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6734 := __e.Get(1)
_ = Key6734
__e.Return(MakeNative(func(__e *ControlFlow) {
C6735 := __e.Get(1)
_ = C6735
tmp19510 := PrimCons(symnumber, Nil)

tmp19511 := PrimCons(sym_1_1_6, tmp19510)

tmp19512 := PrimCons(symnumber, tmp19511)

tmp19513 := PrimCons(tmp19512, Nil)

tmp19514 := PrimCons(sym_1_1_6, tmp19513)

tmp19515 := PrimCons(symnumber, tmp19514)

__e.TailApply(PrimFunc(symis_b), V6736, tmp19515, B6732, L6733, Key6734, C6735)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19516 := PrimValue(symshen_4_dsigf_d)

tmp19517 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_7, tmp19509, tmp19516)


tmp19518 := PrimSet(symshen_4_dsigf_d, tmp19517)

_ = tmp19518

tmp19519 := MakeNative(func(__e *ControlFlow) {
V6741 := __e.Get(1)
_ = V6741
__e.Return(MakeNative(func(__e *ControlFlow) {
B6737 := __e.Get(1)
_ = B6737
__e.Return(MakeNative(func(__e *ControlFlow) {
L6738 := __e.Get(1)
_ = L6738
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6739 := __e.Get(1)
_ = Key6739
__e.Return(MakeNative(func(__e *ControlFlow) {
C6740 := __e.Get(1)
_ = C6740
tmp19520 := PrimCons(symnumber, Nil)

tmp19521 := PrimCons(sym_1_1_6, tmp19520)

tmp19522 := PrimCons(symnumber, tmp19521)

tmp19523 := PrimCons(tmp19522, Nil)

tmp19524 := PrimCons(sym_1_1_6, tmp19523)

tmp19525 := PrimCons(symnumber, tmp19524)

__e.TailApply(PrimFunc(symis_b), V6741, tmp19525, B6737, L6738, Key6739, C6740)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19526 := PrimValue(symshen_4_dsigf_d)

tmp19527 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_c, tmp19519, tmp19526)


tmp19528 := PrimSet(symshen_4_dsigf_d, tmp19527)

_ = tmp19528

tmp19529 := MakeNative(func(__e *ControlFlow) {
V6746 := __e.Get(1)
_ = V6746
__e.Return(MakeNative(func(__e *ControlFlow) {
B6742 := __e.Get(1)
_ = B6742
__e.Return(MakeNative(func(__e *ControlFlow) {
L6743 := __e.Get(1)
_ = L6743
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6744 := __e.Get(1)
_ = Key6744
__e.Return(MakeNative(func(__e *ControlFlow) {
C6745 := __e.Get(1)
_ = C6745
tmp19530 := PrimCons(symnumber, Nil)

tmp19531 := PrimCons(sym_1_1_6, tmp19530)

tmp19532 := PrimCons(symnumber, tmp19531)

tmp19533 := PrimCons(tmp19532, Nil)

tmp19534 := PrimCons(sym_1_1_6, tmp19533)

tmp19535 := PrimCons(symnumber, tmp19534)

__e.TailApply(PrimFunc(symis_b), V6746, tmp19535, B6742, L6743, Key6744, C6745)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19536 := PrimValue(symshen_4_dsigf_d)

tmp19537 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_1, tmp19529, tmp19536)


tmp19538 := PrimSet(symshen_4_dsigf_d, tmp19537)

_ = tmp19538

tmp19539 := MakeNative(func(__e *ControlFlow) {
V6751 := __e.Get(1)
_ = V6751
__e.Return(MakeNative(func(__e *ControlFlow) {
B6747 := __e.Get(1)
_ = B6747
__e.Return(MakeNative(func(__e *ControlFlow) {
L6748 := __e.Get(1)
_ = L6748
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6749 := __e.Get(1)
_ = Key6749
__e.Return(MakeNative(func(__e *ControlFlow) {
C6750 := __e.Get(1)
_ = C6750
tmp19540 := PrimCons(symnumber, Nil)

tmp19541 := PrimCons(sym_1_1_6, tmp19540)

tmp19542 := PrimCons(symnumber, tmp19541)

tmp19543 := PrimCons(tmp19542, Nil)

tmp19544 := PrimCons(sym_1_1_6, tmp19543)

tmp19545 := PrimCons(symnumber, tmp19544)

__e.TailApply(PrimFunc(symis_b), V6751, tmp19545, B6747, L6748, Key6749, C6750)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19546 := PrimValue(symshen_4_dsigf_d)

tmp19547 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_d, tmp19539, tmp19546)


tmp19548 := PrimSet(symshen_4_dsigf_d, tmp19547)

_ = tmp19548

tmp19549 := MakeNative(func(__e *ControlFlow) {
V6756 := __e.Get(1)
_ = V6756
__e.Return(MakeNative(func(__e *ControlFlow) {
B6752 := __e.Get(1)
_ = B6752
__e.Return(MakeNative(func(__e *ControlFlow) {
L6753 := __e.Get(1)
_ = L6753
__e.Return(MakeNative(func(__e *ControlFlow) {
Key6754 := __e.Get(1)
_ = Key6754
__e.Return(MakeNative(func(__e *ControlFlow) {
C6755 := __e.Get(1)
_ = C6755
tmp19550 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp19551 := MakeNative(func(__e *ControlFlow) {
B := __e.Get(1)
_ = B
tmp19552 := PrimCons(symboolean, Nil)

tmp19553 := PrimCons(sym_1_1_6, tmp19552)

tmp19554 := PrimCons(B, tmp19553)

tmp19555 := PrimCons(tmp19554, Nil)

tmp19556 := PrimCons(sym_1_1_6, tmp19555)

tmp19557 := PrimCons(A, tmp19556)

tmp19558 := Call(__e, PrimFunc(symis_b), V6756, tmp19557, B6752, L6753, Key6754, C6755)


__e.TailApply(PrimFunc(symshen_4gc), B6752, tmp19558)
return


}, 1)

tmp19559 := Call(__e, PrimFunc(symshen_4newpv), B6752)


tmp19560 := Call(__e, tmp19551, tmp19559)


__e.TailApply(PrimFunc(symshen_4gc), B6752, tmp19560)
return


}, 1)

tmp19561 := Call(__e, PrimFunc(symshen_4newpv), B6752)


__e.TailApply(tmp19550, tmp19561)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19562 := PrimValue(symshen_4_dsigf_d)

tmp19563 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_a_a, tmp19549, tmp19562)


__e.Return(PrimSet(symshen_4_dsigf_d, tmp19563))
return


}, 0)

tmp19564 := Call(__e, ns2_1set, symshen_4initialise_1signedfuncs, tmp17819)


_ = tmp19564

tmp19565 := MakeNative(func(__e *ControlFlow) {
tmp19566 := MakeNative(func(__e *ControlFlow) {
Y1220 := __e.Get(1)
_ = Y1220
__e.TailApply(PrimFunc(symshen_4tuple), Y1220)
return
}, 1)

tmp19567 := PrimCons(symshen_4tuple, tmp19566)

tmp19568 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19567)


_ = tmp19568

tmp19569 := MakeNative(func(__e *ControlFlow) {
Y1219 := __e.Get(1)
_ = Y1219
__e.TailApply(PrimFunc(symshen_4pvar), Y1219)
return
}, 1)

tmp19570 := PrimCons(symshen_4pvar, tmp19569)

tmp19571 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19570)


_ = tmp19571

tmp19572 := MakeNative(func(__e *ControlFlow) {
Y1218 := __e.Get(1)
_ = Y1218
__e.TailApply(PrimFunc(symshen_4dictionary), Y1218)
return
}, 1)

tmp19573 := PrimCons(symshen_4dictionary, tmp19572)

tmp19574 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19573)


_ = tmp19574

tmp19575 := MakeNative(func(__e *ControlFlow) {
Y1217 := __e.Get(1)
_ = Y1217
__e.TailApply(PrimFunc(symshen_4print_1prolog_1vector), Y1217)
return
}, 1)

tmp19576 := PrimCons(symshen_4print_1prolog_1vector, tmp19575)

tmp19577 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19576)


_ = tmp19577

tmp19578 := MakeNative(func(__e *ControlFlow) {
Y1216 := __e.Get(1)
_ = Y1216
__e.TailApply(PrimFunc(symshen_4print_1freshterm), Y1216)
return
}, 1)

tmp19579 := PrimCons(symshen_4print_1freshterm, tmp19578)

tmp19580 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19579)


_ = tmp19580

tmp19581 := MakeNative(func(__e *ControlFlow) {
Y1215 := __e.Get(1)
_ = Y1215
__e.TailApply(PrimFunc(symshen_4printF), Y1215)
return
}, 1)

tmp19582 := PrimCons(symshen_4printF, tmp19581)

tmp19583 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19582)


_ = tmp19583

tmp19584 := MakeNative(func(__e *ControlFlow) {
Y1214 := __e.Get(1)
_ = Y1214
__e.TailApply(PrimFunc(symabsolute), Y1214)
return
}, 1)

tmp19585 := PrimCons(symabsolute, tmp19584)

tmp19586 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19585)


_ = tmp19586

tmp19587 := MakeNative(func(__e *ControlFlow) {
Y1213 := __e.Get(1)
_ = Y1213
__e.Return(PrimIsVector(Y1213))
return
}, 1)

tmp19588 := PrimCons(symabsvector_2, tmp19587)

tmp19589 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19588)


_ = tmp19589

tmp19590 := MakeNative(func(__e *ControlFlow) {
Y1212 := __e.Get(1)
_ = Y1212
__e.Return(PrimAbsvector(Y1212))
return
}, 1)

tmp19591 := PrimCons(symabsvector, tmp19590)

tmp19592 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19591)


_ = tmp19592

tmp19593 := MakeNative(func(__e *ControlFlow) {
Y1209 := __e.Get(1)
_ = Y1209
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1210 := __e.Get(1)
_ = Y1210
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1211 := __e.Get(1)
_ = Y1211
__e.Return(PrimVectorSet(Y1209, Y1210, Y1211))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19594 := PrimCons(symaddress_1_6, tmp19593)

tmp19595 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19594)


_ = tmp19595

tmp19596 := MakeNative(func(__e *ControlFlow) {
Y1207 := __e.Get(1)
_ = Y1207
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1208 := __e.Get(1)
_ = Y1208
__e.TailApply(PrimFunc(symadjoin), Y1207, Y1208)
return
}, 1))
return
}, 1)

tmp19597 := PrimCons(symadjoin, tmp19596)

tmp19598 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19597)


_ = tmp19598

tmp19599 := MakeNative(func(__e *ControlFlow) {
Y1205 := __e.Get(1)
_ = Y1205
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1206 := __e.Get(1)
_ = Y1206
if True == Y1205 {
if True == Y1206 {
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
}, 1))
return
}, 1)

tmp19602 := PrimCons(symand, tmp19599)

tmp19603 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19602)


_ = tmp19603

tmp19604 := MakeNative(func(__e *ControlFlow) {
Y1203 := __e.Get(1)
_ = Y1203
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1204 := __e.Get(1)
_ = Y1204
__e.TailApply(PrimFunc(symappend), Y1203, Y1204)
return
}, 1))
return
}, 1)

tmp19605 := PrimCons(symappend, tmp19604)

tmp19606 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19605)


_ = tmp19606

tmp19607 := MakeNative(func(__e *ControlFlow) {
Y1202 := __e.Get(1)
_ = Y1202
__e.TailApply(PrimFunc(symarity), Y1202)
return
}, 1)

tmp19608 := PrimCons(symarity, tmp19607)

tmp19609 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19608)


_ = tmp19609

tmp19610 := MakeNative(func(__e *ControlFlow) {
Y1200 := __e.Get(1)
_ = Y1200
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1201 := __e.Get(1)
_ = Y1201
__e.TailApply(PrimFunc(symassoc), Y1200, Y1201)
return
}, 1))
return
}, 1)

tmp19611 := PrimCons(symassoc, tmp19610)

tmp19612 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19611)


_ = tmp19612

tmp19613 := MakeNative(func(__e *ControlFlow) {
Y1199 := __e.Get(1)
_ = Y1199
__e.TailApply(PrimFunc(symatom_2), Y1199)
return
}, 1)

tmp19614 := PrimCons(symatom_2, tmp19613)

tmp19615 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19614)


_ = tmp19615

tmp19616 := MakeNative(func(__e *ControlFlow) {
Y1198 := __e.Get(1)
_ = Y1198
__e.TailApply(PrimFunc(symboolean_2), Y1198)
return
}, 1)

tmp19617 := PrimCons(symboolean_2, tmp19616)

tmp19618 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19617)


_ = tmp19618

tmp19619 := MakeNative(func(__e *ControlFlow) {
Y1197 := __e.Get(1)
_ = Y1197
__e.TailApply(PrimFunc(symbootstrap), Y1197)
return
}, 1)

tmp19620 := PrimCons(symbootstrap, tmp19619)

tmp19621 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19620)


_ = tmp19621

tmp19622 := MakeNative(func(__e *ControlFlow) {
Y1196 := __e.Get(1)
_ = Y1196
__e.TailApply(PrimFunc(symbound_2), Y1196)
return
}, 1)

tmp19623 := PrimCons(symbound_2, tmp19622)

tmp19624 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19623)


_ = tmp19624

tmp19625 := MakeNative(func(__e *ControlFlow) {
Y1190 := __e.Get(1)
_ = Y1190
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1191 := __e.Get(1)
_ = Y1191
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1192 := __e.Get(1)
_ = Y1192
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1193 := __e.Get(1)
_ = Y1193
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1194 := __e.Get(1)
_ = Y1194
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1195 := __e.Get(1)
_ = Y1195
__e.TailApply(PrimFunc(symbind), Y1190, Y1191, Y1192, Y1193, Y1194, Y1195)
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19626 := PrimCons(symbind, tmp19625)

tmp19627 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19626)


_ = tmp19627

tmp19628 := MakeNative(func(__e *ControlFlow) {
Y1185 := __e.Get(1)
_ = Y1185
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1186 := __e.Get(1)
_ = Y1186
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1187 := __e.Get(1)
_ = Y1187
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1188 := __e.Get(1)
_ = Y1188
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1189 := __e.Get(1)
_ = Y1189
__e.TailApply(PrimFunc(symcall), Y1185, Y1186, Y1187, Y1188, Y1189)
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19629 := PrimCons(symcall, tmp19628)

tmp19630 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19629)


_ = tmp19630

tmp19631 := MakeNative(func(__e *ControlFlow) {
Y1184 := __e.Get(1)
_ = Y1184
__e.TailApply(PrimFunc(symcd), Y1184)
return
}, 1)

tmp19632 := PrimCons(symcd, tmp19631)

tmp19633 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19632)


_ = tmp19633

tmp19634 := MakeNative(func(__e *ControlFlow) {
Y1182 := __e.Get(1)
_ = Y1182
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1183 := __e.Get(1)
_ = Y1183
__e.TailApply(PrimFunc(symcompile), Y1182, Y1183)
return
}, 1))
return
}, 1)

tmp19635 := PrimCons(symcompile, tmp19634)

tmp19636 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19635)


_ = tmp19636

tmp19637 := MakeNative(func(__e *ControlFlow) {
Y1180 := __e.Get(1)
_ = Y1180
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1181 := __e.Get(1)
_ = Y1181
__e.TailApply(PrimFunc(symconcat), Y1180, Y1181)
return
}, 1))
return
}, 1)

tmp19638 := PrimCons(symconcat, tmp19637)

tmp19639 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19638)


_ = tmp19639

tmp19640 := MakeNative(func(__e *ControlFlow) {
Y1178 := __e.Get(1)
_ = Y1178
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1179 := __e.Get(1)
_ = Y1179
__e.Return(PrimCons(Y1178, Y1179))
return
}, 1))
return
}, 1)

tmp19641 := PrimCons(symcons, tmp19640)

tmp19642 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19641)


_ = tmp19642

tmp19643 := MakeNative(func(__e *ControlFlow) {
Y1177 := __e.Get(1)
_ = Y1177
__e.Return(PrimIsPair(Y1177))
return
}, 1)

tmp19644 := PrimCons(symcons_2, tmp19643)

tmp19645 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19644)


_ = tmp19645

tmp19646 := MakeNative(func(__e *ControlFlow) {
Y1175 := __e.Get(1)
_ = Y1175
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1176 := __e.Get(1)
_ = Y1176
__e.Return(PrimStringConcat(Y1175, Y1176))
return
}, 1))
return
}, 1)

tmp19647 := PrimCons(symcn, tmp19646)

tmp19648 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19647)


_ = tmp19648

tmp19649 := MakeNative(func(__e *ControlFlow) {
Y1174 := __e.Get(1)
_ = Y1174
__e.Return(PrimCloseStream(Y1174))
return
}, 1)

tmp19650 := PrimCons(symclose, tmp19649)

tmp19651 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19650)


_ = tmp19651

tmp19652 := MakeNative(func(__e *ControlFlow) {
Y1172 := __e.Get(1)
_ = Y1172
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1173 := __e.Get(1)
_ = Y1173
__e.TailApply(PrimFunc(symdeclare), Y1172, Y1173)
return
}, 1))
return
}, 1)

tmp19653 := PrimCons(symdeclare, tmp19652)

tmp19654 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19653)


_ = tmp19654

tmp19655 := MakeNative(func(__e *ControlFlow) {
Y1171 := __e.Get(1)
_ = Y1171
__e.TailApply(PrimFunc(symdestroy), Y1171)
return
}, 1)

tmp19656 := PrimCons(symdestroy, tmp19655)

tmp19657 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19656)


_ = tmp19657

tmp19658 := MakeNative(func(__e *ControlFlow) {
Y1169 := __e.Get(1)
_ = Y1169
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1170 := __e.Get(1)
_ = Y1170
__e.TailApply(PrimFunc(symdifference), Y1169, Y1170)
return
}, 1))
return
}, 1)

tmp19659 := PrimCons(symdifference, tmp19658)

tmp19660 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19659)


_ = tmp19660

tmp19661 := MakeNative(func(__e *ControlFlow) {
Y1167 := __e.Get(1)
_ = Y1167
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1168 := __e.Get(1)
_ = Y1168
_ = Y1167

__e.Return(Y1168)
return


}, 1))
return
}, 1)

tmp19662 := PrimCons(symdo, tmp19661)

tmp19663 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19662)


_ = tmp19663

tmp19664 := MakeNative(func(__e *ControlFlow) {
Y1165 := __e.Get(1)
_ = Y1165
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1166 := __e.Get(1)
_ = Y1166
__e.TailApply(PrimFunc(symelement_2), Y1165, Y1166)
return
}, 1))
return
}, 1)

tmp19665 := PrimCons(symelement_2, tmp19664)

tmp19666 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19665)


_ = tmp19666

tmp19667 := MakeNative(func(__e *ControlFlow) {
Y1164 := __e.Get(1)
_ = Y1164
__e.TailApply(PrimFunc(symempty_2), Y1164)
return
}, 1)

tmp19668 := PrimCons(symempty_2, tmp19667)

tmp19669 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19668)


_ = tmp19669

tmp19670 := MakeNative(func(__e *ControlFlow) {
Y1163 := __e.Get(1)
_ = Y1163
__e.TailApply(PrimFunc(symenable_1type_1theory), Y1163)
return
}, 1)

tmp19671 := PrimCons(symenable_1type_1theory, tmp19670)

tmp19672 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19671)


_ = tmp19672

tmp19673 := MakeNative(func(__e *ControlFlow) {
Y1162 := __e.Get(1)
_ = Y1162
__e.TailApply(PrimFunc(symexternal), Y1162)
return
}, 1)

tmp19674 := PrimCons(symexternal, tmp19673)

tmp19675 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19674)


_ = tmp19675

tmp19676 := MakeNative(func(__e *ControlFlow) {
Y1161 := __e.Get(1)
_ = Y1161
__e.Return(PrimErrorToString(Y1161))
return
}, 1)

tmp19677 := PrimCons(symerror_1to_1string, tmp19676)

tmp19678 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19677)


_ = tmp19678

tmp19679 := MakeNative(func(__e *ControlFlow) {
Y1160 := __e.Get(1)
_ = Y1160
__e.TailApply(PrimFunc(symeval), Y1160)
return
}, 1)

tmp19680 := PrimCons(symeval, tmp19679)

tmp19681 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19680)


_ = tmp19681

tmp19682 := MakeNative(func(__e *ControlFlow) {
Y1159 := __e.Get(1)
_ = Y1159
__e.TailApply(PrimFunc(symeval_1kl), Y1159)
return
}, 1)

tmp19683 := PrimCons(symeval_1kl, tmp19682)

tmp19684 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19683)


_ = tmp19684

tmp19685 := MakeNative(func(__e *ControlFlow) {
Y1158 := __e.Get(1)
_ = Y1158
__e.TailApply(PrimFunc(symexplode), Y1158)
return
}, 1)

tmp19686 := PrimCons(symexplode, tmp19685)

tmp19687 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19686)


_ = tmp19687

tmp19688 := MakeNative(func(__e *ControlFlow) {
Y1157 := __e.Get(1)
_ = Y1157
__e.TailApply(PrimFunc(symexternal), Y1157)
return
}, 1)

tmp19689 := PrimCons(symexternal, tmp19688)

tmp19690 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19689)


_ = tmp19690

tmp19691 := MakeNative(func(__e *ControlFlow) {
Y1156 := __e.Get(1)
_ = Y1156
__e.TailApply(PrimFunc(symfactorise), Y1156)
return
}, 1)

tmp19692 := PrimCons(symfactorise, tmp19691)

tmp19693 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19692)


_ = tmp19693

tmp19694 := MakeNative(func(__e *ControlFlow) {
Y1154 := __e.Get(1)
_ = Y1154
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1155 := __e.Get(1)
_ = Y1155
__e.TailApply(PrimFunc(symfail_1if), Y1154, Y1155)
return
}, 1))
return
}, 1)

tmp19695 := PrimCons(symfail_1if, tmp19694)

tmp19696 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19695)


_ = tmp19696

tmp19697 := MakeNative(func(__e *ControlFlow) {
Y1152 := __e.Get(1)
_ = Y1152
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1153 := __e.Get(1)
_ = Y1153
__e.TailApply(PrimFunc(symfix), Y1152, Y1153)
return
}, 1))
return
}, 1)

tmp19698 := PrimCons(symfix, tmp19697)

tmp19699 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19698)


_ = tmp19699

tmp19700 := MakeNative(func(__e *ControlFlow) {
Y1145 := __e.Get(1)
_ = Y1145
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1146 := __e.Get(1)
_ = Y1146
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1147 := __e.Get(1)
_ = Y1147
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1148 := __e.Get(1)
_ = Y1148
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1149 := __e.Get(1)
_ = Y1149
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1150 := __e.Get(1)
_ = Y1150
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1151 := __e.Get(1)
_ = Y1151
__e.TailApply(PrimFunc(symfindall), Y1145, Y1146, Y1147, Y1148, Y1149, Y1150, Y1151)
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19701 := PrimCons(symfindall, tmp19700)

tmp19702 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19701)


_ = tmp19702

tmp19703 := MakeNative(func(__e *ControlFlow) {
Y1140 := __e.Get(1)
_ = Y1140
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1141 := __e.Get(1)
_ = Y1141
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1142 := __e.Get(1)
_ = Y1142
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1143 := __e.Get(1)
_ = Y1143
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1144 := __e.Get(1)
_ = Y1144
__e.TailApply(PrimFunc(symfork), Y1140, Y1141, Y1142, Y1143, Y1144)
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19704 := PrimCons(symfork, tmp19703)

tmp19705 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19704)


_ = tmp19705

tmp19706 := MakeNative(func(__e *ControlFlow) {
Y1139 := __e.Get(1)
_ = Y1139
__e.Return(MakeNative(func(__e *ControlFlow) {
__e.Return(Y1139)
return
}, 0))
return
}, 1)

tmp19707 := PrimCons(symfreeze, tmp19706)

tmp19708 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19707)


_ = tmp19708

tmp19709 := MakeNative(func(__e *ControlFlow) {
Y1138 := __e.Get(1)
_ = Y1138
__e.TailApply(PrimFunc(symfst), Y1138)
return
}, 1)

tmp19710 := PrimCons(symfst, tmp19709)

tmp19711 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19710)


_ = tmp19711

tmp19712 := MakeNative(func(__e *ControlFlow) {
Y1137 := __e.Get(1)
_ = Y1137
__e.TailApply(PrimFunc(symfn), Y1137)
return
}, 1)

tmp19713 := PrimCons(symfn, tmp19712)

tmp19714 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19713)


_ = tmp19714

tmp19715 := MakeNative(func(__e *ControlFlow) {
Y1136 := __e.Get(1)
_ = Y1136
__e.TailApply(PrimFunc(symfunction), Y1136)
return
}, 1)

tmp19716 := PrimCons(symfunction, tmp19715)

tmp19717 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19716)


_ = tmp19717

tmp19718 := MakeNative(func(__e *ControlFlow) {
Y1135 := __e.Get(1)
_ = Y1135
__e.TailApply(PrimFunc(symgensym), Y1135)
return
}, 1)

tmp19719 := PrimCons(symgensym, tmp19718)

tmp19720 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19719)


_ = tmp19720

tmp19721 := MakeNative(func(__e *ControlFlow) {
Y1132 := __e.Get(1)
_ = Y1132
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1133 := __e.Get(1)
_ = Y1133
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1134 := __e.Get(1)
_ = Y1134
__e.TailApply(PrimFunc(symget), Y1132, Y1133, Y1134)
return
}, 1))
return
}, 1))
return
}, 1)

tmp19722 := PrimCons(symget, tmp19721)

tmp19723 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19722)


_ = tmp19723

tmp19724 := MakeNative(func(__e *ControlFlow) {
Y1131 := __e.Get(1)
_ = Y1131
__e.Return(PrimGetTime(Y1131))
return
}, 1)

tmp19725 := PrimCons(symget_1time, tmp19724)

tmp19726 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19725)


_ = tmp19726

tmp19727 := MakeNative(func(__e *ControlFlow) {
Y1128 := __e.Get(1)
_ = Y1128
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1129 := __e.Get(1)
_ = Y1129
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1130 := __e.Get(1)
_ = Y1130
__e.Return(PrimVectorSet(Y1128, Y1129, Y1130))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19728 := PrimCons(symaddress_1_6, tmp19727)

tmp19729 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19728)


_ = tmp19729

tmp19730 := MakeNative(func(__e *ControlFlow) {
Y1126 := __e.Get(1)
_ = Y1126
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1127 := __e.Get(1)
_ = Y1127
__e.Return(PrimVectorGet(Y1126, Y1127))
return
}, 1))
return
}, 1)

tmp19731 := PrimCons(sym_5_1address, tmp19730)

tmp19732 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19731)


_ = tmp19732

tmp19733 := MakeNative(func(__e *ControlFlow) {
Y1124 := __e.Get(1)
_ = Y1124
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1125 := __e.Get(1)
_ = Y1125
__e.TailApply(PrimFunc(sym_5_1vector), Y1124, Y1125)
return
}, 1))
return
}, 1)

tmp19734 := PrimCons(sym_5_1vector, tmp19733)

tmp19735 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19734)


_ = tmp19735

tmp19736 := MakeNative(func(__e *ControlFlow) {
Y1122 := __e.Get(1)
_ = Y1122
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1123 := __e.Get(1)
_ = Y1123
__e.Return(PrimGreatThan(Y1122, Y1123))
return
}, 1))
return
}, 1)

tmp19737 := PrimCons(sym_6, tmp19736)

tmp19738 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19737)


_ = tmp19738

tmp19739 := MakeNative(func(__e *ControlFlow) {
Y1120 := __e.Get(1)
_ = Y1120
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1121 := __e.Get(1)
_ = Y1121
__e.Return(PrimGreatEqual(Y1120, Y1121))
return
}, 1))
return
}, 1)

tmp19740 := PrimCons(sym_6_a, tmp19739)

tmp19741 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19740)


_ = tmp19741

tmp19742 := MakeNative(func(__e *ControlFlow) {
Y1118 := __e.Get(1)
_ = Y1118
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1119 := __e.Get(1)
_ = Y1119
__e.Return(PrimEqual(Y1118, Y1119))
return
}, 1))
return
}, 1)

tmp19743 := PrimCons(sym_a, tmp19742)

tmp19744 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19743)


_ = tmp19744

tmp19745 := MakeNative(func(__e *ControlFlow) {
Y1116 := __e.Get(1)
_ = Y1116
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1117 := __e.Get(1)
_ = Y1117
__e.TailApply(PrimFunc(symhash), Y1116, Y1117)
return
}, 1))
return
}, 1)

tmp19746 := PrimCons(symhash, tmp19745)

tmp19747 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19746)


_ = tmp19747

tmp19748 := MakeNative(func(__e *ControlFlow) {
Y1115 := __e.Get(1)
_ = Y1115
__e.Return(PrimHead(Y1115))
return
}, 1)

tmp19749 := PrimCons(symhd, tmp19748)

tmp19750 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19749)


_ = tmp19750

tmp19751 := MakeNative(func(__e *ControlFlow) {
Y1114 := __e.Get(1)
_ = Y1114
__e.TailApply(PrimFunc(symhdv), Y1114)
return
}, 1)

tmp19752 := PrimCons(symhdv, tmp19751)

tmp19753 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19752)


_ = tmp19753

tmp19754 := MakeNative(func(__e *ControlFlow) {
Y1113 := __e.Get(1)
_ = Y1113
__e.TailApply(PrimFunc(symhdstr), Y1113)
return
}, 1)

tmp19755 := PrimCons(symhdstr, tmp19754)

tmp19756 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19755)


_ = tmp19756

tmp19757 := MakeNative(func(__e *ControlFlow) {
Y1112 := __e.Get(1)
_ = Y1112
__e.TailApply(PrimFunc(symhead), Y1112)
return
}, 1)

tmp19758 := PrimCons(symhead, tmp19757)

tmp19759 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19758)


_ = tmp19759

tmp19760 := MakeNative(func(__e *ControlFlow) {
Y1111 := __e.Get(1)
_ = Y1111
__e.TailApply(PrimFunc(symhush), Y1111)
return
}, 1)

tmp19761 := PrimCons(symhush, tmp19760)

tmp19762 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19761)


_ = tmp19762

tmp19763 := MakeNative(func(__e *ControlFlow) {
Y1108 := __e.Get(1)
_ = Y1108
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1109 := __e.Get(1)
_ = Y1109
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1110 := __e.Get(1)
_ = Y1110
if True == Y1108 {
__e.Return(Y1109)
return
} else {
__e.Return(Y1110)
return
}
}, 1))
return
}, 1))
return
}, 1)

tmp19765 := PrimCons(symif, tmp19763)

tmp19766 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19765)


_ = tmp19766

tmp19767 := MakeNative(func(__e *ControlFlow) {
Y1107 := __e.Get(1)
_ = Y1107
__e.TailApply(PrimFunc(syminclude), Y1107)
return
}, 1)

tmp19768 := PrimCons(syminclude, tmp19767)

tmp19769 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19768)


_ = tmp19769

tmp19770 := MakeNative(func(__e *ControlFlow) {
Y1106 := __e.Get(1)
_ = Y1106
__e.TailApply(PrimFunc(symin_1package), Y1106)
return
}, 1)

tmp19771 := PrimCons(symin_1package, tmp19770)

tmp19772 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19771)


_ = tmp19772

tmp19773 := MakeNative(func(__e *ControlFlow) {
Y1105 := __e.Get(1)
_ = Y1105
__e.Return(PrimIsInteger(Y1105))
return
}, 1)

tmp19774 := PrimCons(syminteger_2, tmp19773)

tmp19775 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19774)


_ = tmp19775

tmp19776 := MakeNative(func(__e *ControlFlow) {
Y1104 := __e.Get(1)
_ = Y1104
__e.TailApply(PrimFunc(syminternal), Y1104)
return
}, 1)

tmp19777 := PrimCons(syminternal, tmp19776)

tmp19778 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19777)


_ = tmp19778

tmp19779 := MakeNative(func(__e *ControlFlow) {
Y1103 := __e.Get(1)
_ = Y1103
__e.Return(PrimIntern(Y1103))
return
}, 1)

tmp19780 := PrimCons(symintern, tmp19779)

tmp19781 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19780)


_ = tmp19781

tmp19782 := MakeNative(func(__e *ControlFlow) {
Y1102 := __e.Get(1)
_ = Y1102
__e.TailApply(PrimFunc(syminput), Y1102)
return
}, 1)

tmp19783 := PrimCons(syminput, tmp19782)

tmp19784 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19783)


_ = tmp19784

tmp19785 := MakeNative(func(__e *ControlFlow) {
Y1100 := __e.Get(1)
_ = Y1100
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1101 := __e.Get(1)
_ = Y1101
__e.TailApply(PrimFunc(syminput_7), Y1100, Y1101)
return
}, 1))
return
}, 1)

tmp19786 := PrimCons(syminput_7, tmp19785)

tmp19787 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19786)


_ = tmp19787

tmp19788 := MakeNative(func(__e *ControlFlow) {
Y1099 := __e.Get(1)
_ = Y1099
__e.TailApply(PrimFunc(syminclude_1all_1but), Y1099)
return
}, 1)

tmp19789 := PrimCons(syminclude_1all_1but, tmp19788)

tmp19790 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19789)


_ = tmp19790

tmp19791 := MakeNative(func(__e *ControlFlow) {
Y1097 := __e.Get(1)
_ = Y1097
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1098 := __e.Get(1)
_ = Y1098
__e.TailApply(PrimFunc(symintersection), Y1097, Y1098)
return
}, 1))
return
}, 1)

tmp19792 := PrimCons(symintersection, tmp19791)

tmp19793 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19792)


_ = tmp19793

tmp19794 := MakeNative(func(__e *ControlFlow) {
Y1096 := __e.Get(1)
_ = Y1096
__e.TailApply(PrimFunc(syminternal), Y1096)
return
}, 1)

tmp19795 := PrimCons(syminternal, tmp19794)

tmp19796 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19795)


_ = tmp19796

tmp19797 := MakeNative(func(__e *ControlFlow) {
Y1090 := __e.Get(1)
_ = Y1090
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1091 := __e.Get(1)
_ = Y1091
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1092 := __e.Get(1)
_ = Y1092
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1093 := __e.Get(1)
_ = Y1093
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1094 := __e.Get(1)
_ = Y1094
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1095 := __e.Get(1)
_ = Y1095
__e.TailApply(PrimFunc(symis), Y1090, Y1091, Y1092, Y1093, Y1094, Y1095)
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19798 := PrimCons(symis, tmp19797)

tmp19799 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19798)


_ = tmp19799

tmp19800 := MakeNative(func(__e *ControlFlow) {
Y1084 := __e.Get(1)
_ = Y1084
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1085 := __e.Get(1)
_ = Y1085
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1086 := __e.Get(1)
_ = Y1086
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1087 := __e.Get(1)
_ = Y1087
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1088 := __e.Get(1)
_ = Y1088
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1089 := __e.Get(1)
_ = Y1089
__e.TailApply(PrimFunc(symis_b), Y1084, Y1085, Y1086, Y1087, Y1088, Y1089)
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19801 := PrimCons(symis_b, tmp19800)

tmp19802 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19801)


_ = tmp19802

tmp19803 := MakeNative(func(__e *ControlFlow) {
Y1083 := __e.Get(1)
_ = Y1083
__e.TailApply(PrimFunc(symlength), Y1083)
return
}, 1)

tmp19804 := PrimCons(symlength, tmp19803)

tmp19805 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19804)


_ = tmp19805

tmp19806 := MakeNative(func(__e *ControlFlow) {
Y1082 := __e.Get(1)
_ = Y1082
__e.TailApply(PrimFunc(symlimit), Y1082)
return
}, 1)

tmp19807 := PrimCons(symlimit, tmp19806)

tmp19808 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19807)


_ = tmp19808

tmp19809 := MakeNative(func(__e *ControlFlow) {
Y1081 := __e.Get(1)
_ = Y1081
__e.TailApply(PrimFunc(symlineread), Y1081)
return
}, 1)

tmp19810 := PrimCons(symlineread, tmp19809)

tmp19811 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19810)


_ = tmp19811

tmp19812 := MakeNative(func(__e *ControlFlow) {
Y1080 := __e.Get(1)
_ = Y1080
__e.TailApply(PrimFunc(symload), Y1080)
return
}, 1)

tmp19813 := PrimCons(symload, tmp19812)

tmp19814 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19813)


_ = tmp19814

tmp19815 := MakeNative(func(__e *ControlFlow) {
Y1078 := __e.Get(1)
_ = Y1078
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1079 := __e.Get(1)
_ = Y1079
__e.Return(PrimLessThan(Y1078, Y1079))
return
}, 1))
return
}, 1)

tmp19816 := PrimCons(sym_5, tmp19815)

tmp19817 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19816)


_ = tmp19817

tmp19818 := MakeNative(func(__e *ControlFlow) {
Y1076 := __e.Get(1)
_ = Y1076
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1077 := __e.Get(1)
_ = Y1077
__e.Return(PrimLessEqual(Y1076, Y1077))
return
}, 1))
return
}, 1)

tmp19819 := PrimCons(sym_5_a, tmp19818)

tmp19820 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19819)


_ = tmp19820

tmp19821 := MakeNative(func(__e *ControlFlow) {
Y1075 := __e.Get(1)
_ = Y1075
__e.TailApply(PrimFunc(symvector), Y1075)
return
}, 1)

tmp19822 := PrimCons(symvector, tmp19821)

tmp19823 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19822)


_ = tmp19823

tmp19824 := MakeNative(func(__e *ControlFlow) {
Y1074 := __e.Get(1)
_ = Y1074
__e.TailApply(PrimFunc(symmacroexpand), Y1074)
return
}, 1)

tmp19825 := PrimCons(symmacroexpand, tmp19824)

tmp19826 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19825)


_ = tmp19826

tmp19827 := MakeNative(func(__e *ControlFlow) {
Y1072 := __e.Get(1)
_ = Y1072
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1073 := __e.Get(1)
_ = Y1073
__e.TailApply(PrimFunc(symmap), Y1072, Y1073)
return
}, 1))
return
}, 1)

tmp19828 := PrimCons(symmap, tmp19827)

tmp19829 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19828)


_ = tmp19829

tmp19830 := MakeNative(func(__e *ControlFlow) {
Y1070 := __e.Get(1)
_ = Y1070
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1071 := __e.Get(1)
_ = Y1071
__e.TailApply(PrimFunc(symmapcan), Y1070, Y1071)
return
}, 1))
return
}, 1)

tmp19831 := PrimCons(symmapcan, tmp19830)

tmp19832 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19831)


_ = tmp19832

tmp19833 := MakeNative(func(__e *ControlFlow) {
Y1069 := __e.Get(1)
_ = Y1069
__e.TailApply(PrimFunc(symmaxinferences), Y1069)
return
}, 1)

tmp19834 := PrimCons(symmaxinferences, tmp19833)

tmp19835 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19834)


_ = tmp19835

tmp19836 := MakeNative(func(__e *ControlFlow) {
Y1068 := __e.Get(1)
_ = Y1068
__e.TailApply(PrimFunc(symnl), Y1068)
return
}, 1)

tmp19837 := PrimCons(symnl, tmp19836)

tmp19838 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19837)


_ = tmp19838

tmp19839 := MakeNative(func(__e *ControlFlow) {
Y1067 := __e.Get(1)
_ = Y1067
__e.Return(PrimNot(Y1067))
return
}, 1)

tmp19840 := PrimCons(symnot, tmp19839)

tmp19841 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19840)


_ = tmp19841

tmp19842 := MakeNative(func(__e *ControlFlow) {
Y1065 := __e.Get(1)
_ = Y1065
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1066 := __e.Get(1)
_ = Y1066
__e.TailApply(PrimFunc(symnth), Y1065, Y1066)
return
}, 1))
return
}, 1)

tmp19843 := PrimCons(symnth, tmp19842)

tmp19844 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19843)


_ = tmp19844

tmp19845 := MakeNative(func(__e *ControlFlow) {
Y1064 := __e.Get(1)
_ = Y1064
__e.Return(PrimNumberToString(Y1064))
return
}, 1)

tmp19846 := PrimCons(symn_1_6string, tmp19845)

tmp19847 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19846)


_ = tmp19847

tmp19848 := MakeNative(func(__e *ControlFlow) {
Y1063 := __e.Get(1)
_ = Y1063
__e.Return(PrimIsNumber(Y1063))
return
}, 1)

tmp19849 := PrimCons(symnumber_2, tmp19848)

tmp19850 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19849)


_ = tmp19850

tmp19851 := MakeNative(func(__e *ControlFlow) {
Y1062 := __e.Get(1)
_ = Y1062
__e.TailApply(PrimFunc(symoccurs_1check), Y1062)
return
}, 1)

tmp19852 := PrimCons(symoccurs_1check, tmp19851)

tmp19853 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19852)


_ = tmp19853

tmp19854 := MakeNative(func(__e *ControlFlow) {
Y1060 := __e.Get(1)
_ = Y1060
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1061 := __e.Get(1)
_ = Y1061
__e.TailApply(PrimFunc(symoccurrences), Y1060, Y1061)
return
}, 1))
return
}, 1)

tmp19855 := PrimCons(symoccurrences, tmp19854)

tmp19856 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19855)


_ = tmp19856

tmp19857 := MakeNative(func(__e *ControlFlow) {
Y1059 := __e.Get(1)
_ = Y1059
__e.TailApply(PrimFunc(symoccurs_1check), Y1059)
return
}, 1)

tmp19858 := PrimCons(symoccurs_1check, tmp19857)

tmp19859 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19858)


_ = tmp19859

tmp19860 := MakeNative(func(__e *ControlFlow) {
Y1057 := __e.Get(1)
_ = Y1057
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1058 := __e.Get(1)
_ = Y1058
__e.Return(PrimOpenStream(Y1057, Y1058))
return
}, 1))
return
}, 1)

tmp19861 := PrimCons(symopen, tmp19860)

tmp19862 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19861)


_ = tmp19862

tmp19863 := MakeNative(func(__e *ControlFlow) {
Y1056 := __e.Get(1)
_ = Y1056
__e.TailApply(PrimFunc(symoptimise), Y1056)
return
}, 1)

tmp19864 := PrimCons(symoptimise, tmp19863)

tmp19865 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19864)


_ = tmp19865

tmp19866 := MakeNative(func(__e *ControlFlow) {
Y1054 := __e.Get(1)
_ = Y1054
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1055 := __e.Get(1)
_ = Y1055
if True == Y1054 {
__e.Return(True)
return
} else {
if True == Y1055 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}
}
}, 1))
return
}, 1)

tmp19869 := PrimCons(symor, tmp19866)

tmp19870 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19869)


_ = tmp19870

tmp19871 := MakeNative(func(__e *ControlFlow) {
Y1053 := __e.Get(1)
_ = Y1053
__e.TailApply(PrimFunc(sympackage_2), Y1053)
return
}, 1)

tmp19872 := PrimCons(sympackage_2, tmp19871)

tmp19873 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19872)


_ = tmp19873

tmp19874 := MakeNative(func(__e *ControlFlow) {
Y1051 := __e.Get(1)
_ = Y1051
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1052 := __e.Get(1)
_ = Y1052
__e.Return(PrimPos(Y1051, Y1052))
return
}, 1))
return
}, 1)

tmp19875 := PrimCons(sympos, tmp19874)

tmp19876 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19875)


_ = tmp19876

tmp19877 := MakeNative(func(__e *ControlFlow) {
Y1050 := __e.Get(1)
_ = Y1050
__e.TailApply(PrimFunc(sympreclude_1all_1but), Y1050)
return
}, 1)

tmp19878 := PrimCons(sympreclude_1all_1but, tmp19877)

tmp19879 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19878)


_ = tmp19879

tmp19880 := MakeNative(func(__e *ControlFlow) {
Y1049 := __e.Get(1)
_ = Y1049
__e.TailApply(PrimFunc(symprint), Y1049)
return
}, 1)

tmp19881 := PrimCons(symprint, tmp19880)

tmp19882 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19881)


_ = tmp19882

tmp19883 := MakeNative(func(__e *ControlFlow) {
Y1048 := __e.Get(1)
_ = Y1048
__e.TailApply(PrimFunc(symprofile), Y1048)
return
}, 1)

tmp19884 := PrimCons(symprofile, tmp19883)

tmp19885 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19884)


_ = tmp19885

tmp19886 := MakeNative(func(__e *ControlFlow) {
Y1047 := __e.Get(1)
_ = Y1047
__e.TailApply(PrimFunc(symshen_4print_1prolog_1vector), Y1047)
return
}, 1)

tmp19887 := PrimCons(symshen_4print_1prolog_1vector, tmp19886)

tmp19888 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19887)


_ = tmp19888

tmp19889 := MakeNative(func(__e *ControlFlow) {
Y1046 := __e.Get(1)
_ = Y1046
__e.TailApply(PrimFunc(symshen_4print_1freshterm), Y1046)
return
}, 1)

tmp19890 := PrimCons(symshen_4print_1freshterm, tmp19889)

tmp19891 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19890)


_ = tmp19891

tmp19892 := MakeNative(func(__e *ControlFlow) {
Y1045 := __e.Get(1)
_ = Y1045
__e.TailApply(PrimFunc(symshen_4printF), Y1045)
return
}, 1)

tmp19893 := PrimCons(symshen_4printF, tmp19892)

tmp19894 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19893)


_ = tmp19894

tmp19895 := MakeNative(func(__e *ControlFlow) {
Y1044 := __e.Get(1)
_ = Y1044
__e.TailApply(PrimFunc(symprolog_1memory), Y1044)
return
}, 1)

tmp19896 := PrimCons(symprolog_1memory, tmp19895)

tmp19897 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19896)


_ = tmp19897

tmp19898 := MakeNative(func(__e *ControlFlow) {
Y1043 := __e.Get(1)
_ = Y1043
__e.TailApply(PrimFunc(symprofile_1results), Y1043)
return
}, 1)

tmp19899 := PrimCons(symprofile_1results, tmp19898)

tmp19900 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19899)


_ = tmp19900

tmp19901 := MakeNative(func(__e *ControlFlow) {
Y1041 := __e.Get(1)
_ = Y1041
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1042 := __e.Get(1)
_ = Y1042
__e.TailApply(PrimFunc(sympr), Y1041, Y1042)
return
}, 1))
return
}, 1)

tmp19902 := PrimCons(sympr, tmp19901)

tmp19903 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19902)


_ = tmp19903

tmp19904 := MakeNative(func(__e *ControlFlow) {
Y1040 := __e.Get(1)
_ = Y1040
__e.TailApply(PrimFunc(symps), Y1040)
return
}, 1)

tmp19905 := PrimCons(symps, tmp19904)

tmp19906 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19905)


_ = tmp19906

tmp19907 := MakeNative(func(__e *ControlFlow) {
Y1039 := __e.Get(1)
_ = Y1039
__e.TailApply(PrimFunc(sympreclude), Y1039)
return
}, 1)

tmp19908 := PrimCons(sympreclude, tmp19907)

tmp19909 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19908)


_ = tmp19909

tmp19910 := MakeNative(func(__e *ControlFlow) {
Y1038 := __e.Get(1)
_ = Y1038
__e.TailApply(PrimFunc(sympreclude_1all_1but), Y1038)
return
}, 1)

tmp19911 := PrimCons(sympreclude_1all_1but, tmp19910)

tmp19912 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19911)


_ = tmp19912

tmp19913 := MakeNative(func(__e *ControlFlow) {
Y1037 := __e.Get(1)
_ = Y1037
__e.TailApply(PrimFunc(symprotect), Y1037)
return
}, 1)

tmp19914 := PrimCons(symprotect, tmp19913)

tmp19915 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19914)


_ = tmp19915

tmp19916 := MakeNative(func(__e *ControlFlow) {
Y1033 := __e.Get(1)
_ = Y1033
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1034 := __e.Get(1)
_ = Y1034
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1035 := __e.Get(1)
_ = Y1035
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1036 := __e.Get(1)
_ = Y1036
__e.TailApply(PrimFunc(symput), Y1033, Y1034, Y1035, Y1036)
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp19917 := PrimCons(symput, tmp19916)

tmp19918 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19917)


_ = tmp19918

tmp19919 := MakeNative(func(__e *ControlFlow) {
Y1032 := __e.Get(1)
_ = Y1032
__e.Return(PrimReadFileAsString(Y1032))
return
}, 1)

tmp19920 := PrimCons(symread_1file_1as_1string, tmp19919)

tmp19921 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19920)


_ = tmp19921

tmp19922 := MakeNative(func(__e *ControlFlow) {
Y1031 := __e.Get(1)
_ = Y1031
__e.Return(PrimReadFileAsByteList(Y1031))
return
}, 1)

tmp19923 := PrimCons(symread_1file_1as_1bytelist, tmp19922)

tmp19924 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19923)


_ = tmp19924

tmp19925 := MakeNative(func(__e *ControlFlow) {
Y1030 := __e.Get(1)
_ = Y1030
__e.TailApply(PrimFunc(symread_1file), Y1030)
return
}, 1)

tmp19926 := PrimCons(symread_1file, tmp19925)

tmp19927 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19926)


_ = tmp19927

tmp19928 := MakeNative(func(__e *ControlFlow) {
Y1029 := __e.Get(1)
_ = Y1029
__e.TailApply(PrimFunc(symread), Y1029)
return
}, 1)

tmp19929 := PrimCons(symread, tmp19928)

tmp19930 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19929)


_ = tmp19930

tmp19931 := MakeNative(func(__e *ControlFlow) {
Y1028 := __e.Get(1)
_ = Y1028
__e.Return(PrimReadByte(Y1028))
return
}, 1)

tmp19932 := PrimCons(symread_1byte, tmp19931)

tmp19933 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19932)


_ = tmp19933

tmp19934 := MakeNative(func(__e *ControlFlow) {
Y1027 := __e.Get(1)
_ = Y1027
__e.TailApply(PrimFunc(symread_1from_1string), Y1027)
return
}, 1)

tmp19935 := PrimCons(symread_1from_1string, tmp19934)

tmp19936 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19935)


_ = tmp19936

tmp19937 := MakeNative(func(__e *ControlFlow) {
Y1026 := __e.Get(1)
_ = Y1026
__e.TailApply(PrimFunc(symread_1from_1string_1unprocessed), Y1026)
return
}, 1)

tmp19938 := PrimCons(symread_1from_1string_1unprocessed, tmp19937)

tmp19939 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19938)


_ = tmp19939

tmp19940 := MakeNative(func(__e *ControlFlow) {
Y1025 := __e.Get(1)
_ = Y1025
__e.TailApply(PrimFunc(symshen_4read_1unit_1string), Y1025)
return
}, 1)

tmp19941 := PrimCons(symshen_4read_1unit_1string, tmp19940)

tmp19942 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19941)


_ = tmp19942

tmp19943 := MakeNative(func(__e *ControlFlow) {
Y1023 := __e.Get(1)
_ = Y1023
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1024 := __e.Get(1)
_ = Y1024
__e.TailApply(PrimFunc(symremove), Y1023, Y1024)
return
}, 1))
return
}, 1)

tmp19944 := PrimCons(symremove, tmp19943)

tmp19945 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19944)


_ = tmp19945

tmp19946 := MakeNative(func(__e *ControlFlow) {
Y1022 := __e.Get(1)
_ = Y1022
__e.TailApply(PrimFunc(symreverse), Y1022)
return
}, 1)

tmp19947 := PrimCons(symreverse, tmp19946)

tmp19948 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19947)


_ = tmp19948

tmp19949 := MakeNative(func(__e *ControlFlow) {
Y1020 := __e.Get(1)
_ = Y1020
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1021 := __e.Get(1)
_ = Y1021
__e.Return(PrimSet(Y1020, Y1021))
return
}, 1))
return
}, 1)

tmp19950 := PrimCons(symset, tmp19949)

tmp19951 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19950)


_ = tmp19951

tmp19952 := MakeNative(func(__e *ControlFlow) {
Y1019 := __e.Get(1)
_ = Y1019
__e.Return(PrimSimpleError(Y1019))
return
}, 1)

tmp19953 := PrimCons(symsimple_1error, tmp19952)

tmp19954 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19953)


_ = tmp19954

tmp19955 := MakeNative(func(__e *ControlFlow) {
Y1018 := __e.Get(1)
_ = Y1018
__e.TailApply(PrimFunc(symsnd), Y1018)
return
}, 1)

tmp19956 := PrimCons(symsnd, tmp19955)

tmp19957 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19956)


_ = tmp19957

tmp19958 := MakeNative(func(__e *ControlFlow) {
Y1016 := __e.Get(1)
_ = Y1016
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1017 := __e.Get(1)
_ = Y1017
__e.TailApply(PrimFunc(symspecialise), Y1016, Y1017)
return
}, 1))
return
}, 1)

tmp19959 := PrimCons(symspecialise, tmp19958)

tmp19960 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19959)


_ = tmp19960

tmp19961 := MakeNative(func(__e *ControlFlow) {
Y1015 := __e.Get(1)
_ = Y1015
__e.TailApply(PrimFunc(symspy), Y1015)
return
}, 1)

tmp19962 := PrimCons(symspy, tmp19961)

tmp19963 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19962)


_ = tmp19963

tmp19964 := MakeNative(func(__e *ControlFlow) {
Y1014 := __e.Get(1)
_ = Y1014
__e.TailApply(PrimFunc(symstep), Y1014)
return
}, 1)

tmp19965 := PrimCons(symstep, tmp19964)

tmp19966 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19965)


_ = tmp19966

tmp19967 := MakeNative(func(__e *ControlFlow) {
Y1013 := __e.Get(1)
_ = Y1013
__e.Return(PrimStr(Y1013))
return
}, 1)

tmp19968 := PrimCons(symstr, tmp19967)

tmp19969 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19968)


_ = tmp19969

tmp19970 := MakeNative(func(__e *ControlFlow) {
Y1012 := __e.Get(1)
_ = Y1012
__e.Return(PrimStringToNumber(Y1012))
return
}, 1)

tmp19971 := PrimCons(symstring_1_6n, tmp19970)

tmp19972 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19971)


_ = tmp19972

tmp19973 := MakeNative(func(__e *ControlFlow) {
Y1011 := __e.Get(1)
_ = Y1011
__e.TailApply(PrimFunc(symstring_1_6symbol), Y1011)
return
}, 1)

tmp19974 := PrimCons(symstring_1_6symbol, tmp19973)

tmp19975 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19974)


_ = tmp19975

tmp19976 := MakeNative(func(__e *ControlFlow) {
Y1010 := __e.Get(1)
_ = Y1010
__e.Return(PrimIsString(Y1010))
return
}, 1)

tmp19977 := PrimCons(symstring_2, tmp19976)

tmp19978 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19977)


_ = tmp19978

tmp19979 := MakeNative(func(__e *ControlFlow) {
Y1007 := __e.Get(1)
_ = Y1007
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1008 := __e.Get(1)
_ = Y1008
__e.Return(MakeNative(func(__e *ControlFlow) {
Y1009 := __e.Get(1)
_ = Y1009
__e.TailApply(PrimFunc(symsubst), Y1007, Y1008, Y1009)
return
}, 1))
return
}, 1))
return
}, 1)

tmp19980 := PrimCons(symsubst, tmp19979)

tmp19981 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19980)


_ = tmp19981

tmp19982 := MakeNative(func(__e *ControlFlow) {
Y1006 := __e.Get(1)
_ = Y1006
__e.TailApply(PrimFunc(symsum), Y1006)
return
}, 1)

tmp19983 := PrimCons(symsum, tmp19982)

tmp19984 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19983)


_ = tmp19984

tmp19985 := MakeNative(func(__e *ControlFlow) {
Y1005 := __e.Get(1)
_ = Y1005
__e.Return(PrimIsSymbol(Y1005))
return
}, 1)

tmp19986 := PrimCons(symsymbol_2, tmp19985)

tmp19987 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19986)


_ = tmp19987

tmp19988 := MakeNative(func(__e *ControlFlow) {
Y1004 := __e.Get(1)
_ = Y1004
__e.TailApply(PrimFunc(symsystemf), Y1004)
return
}, 1)

tmp19989 := PrimCons(symsystemf, tmp19988)

tmp19990 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19989)


_ = tmp19990

tmp19991 := MakeNative(func(__e *ControlFlow) {
Y1003 := __e.Get(1)
_ = Y1003
__e.TailApply(PrimFunc(symtail), Y1003)
return
}, 1)

tmp19992 := PrimCons(symtail, tmp19991)

tmp19993 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19992)


_ = tmp19993

tmp19994 := MakeNative(func(__e *ControlFlow) {
Y1002 := __e.Get(1)
_ = Y1002
__e.Return(PrimTail(Y1002))
return
}, 1)

tmp19995 := PrimCons(symtl, tmp19994)

tmp19996 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19995)


_ = tmp19996

tmp19997 := MakeNative(func(__e *ControlFlow) {
Y1001 := __e.Get(1)
_ = Y1001
__e.TailApply(PrimFunc(symtc), Y1001)
return
}, 1)

tmp19998 := PrimCons(symtc, tmp19997)

tmp19999 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp19998)


_ = tmp19999

tmp20000 := MakeNative(func(__e *ControlFlow) {
Y1000 := __e.Get(1)
_ = Y1000
__e.TailApply(PrimFunc(symthaw), Y1000)
return
}, 1)

tmp20001 := PrimCons(symthaw, tmp20000)

tmp20002 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp20001)


_ = tmp20002

tmp20003 := MakeNative(func(__e *ControlFlow) {
Y999 := __e.Get(1)
_ = Y999
__e.Return(PrimTailString(Y999))
return
}, 1)

tmp20004 := PrimCons(symtlstr, tmp20003)

tmp20005 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp20004)


_ = tmp20005

tmp20006 := MakeNative(func(__e *ControlFlow) {
Y998 := __e.Get(1)
_ = Y998
__e.TailApply(PrimFunc(symtrack), Y998)
return
}, 1)

tmp20007 := PrimCons(symtrack, tmp20006)

tmp20008 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp20007)


_ = tmp20008

tmp20009 := MakeNative(func(__e *ControlFlow) {
Y996 := __e.Get(1)
_ = Y996
__e.Return(MakeNative(func(__e *ControlFlow) {
Y997 := __e.Get(1)
_ = Y997
tmp20010 := MakeNative(func(__e *ControlFlow) {
__e.Return(Y996)
return
}, 0)

__e.TailApply(try_1catch, tmp20010, Y997)
return


}, 1))
return
}, 1)

tmp20011 := PrimCons(symtrap_1error, tmp20009)

tmp20012 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp20011)


_ = tmp20012

tmp20013 := MakeNative(func(__e *ControlFlow) {
Y995 := __e.Get(1)
_ = Y995
__e.TailApply(PrimFunc(symtuple_2), Y995)
return
}, 1)

tmp20014 := PrimCons(symtuple_2, tmp20013)

tmp20015 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp20014)


_ = tmp20015

tmp20016 := MakeNative(func(__e *ControlFlow) {
Y993 := __e.Get(1)
_ = Y993
__e.Return(MakeNative(func(__e *ControlFlow) {
Y994 := __e.Get(1)
_ = Y994
__e.Return(Y993)
return
}, 1))
return
}, 1)

tmp20017 := PrimCons(symtype, tmp20016)

tmp20018 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp20017)


_ = tmp20018

tmp20019 := MakeNative(func(__e *ControlFlow) {
Y988 := __e.Get(1)
_ = Y988
__e.Return(MakeNative(func(__e *ControlFlow) {
Y989 := __e.Get(1)
_ = Y989
__e.Return(MakeNative(func(__e *ControlFlow) {
Y990 := __e.Get(1)
_ = Y990
__e.Return(MakeNative(func(__e *ControlFlow) {
Y991 := __e.Get(1)
_ = Y991
__e.Return(MakeNative(func(__e *ControlFlow) {
Y992 := __e.Get(1)
_ = Y992
__e.TailApply(PrimFunc(symreturn), Y988, Y989, Y990, Y991, Y992)
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20020 := PrimCons(symreturn, tmp20019)

tmp20021 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp20020)


_ = tmp20021

tmp20022 := MakeNative(func(__e *ControlFlow) {
Y987 := __e.Get(1)
_ = Y987
__e.TailApply(PrimFunc(symunabsolute), Y987)
return
}, 1)

tmp20023 := PrimCons(symunabsolute, tmp20022)

tmp20024 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp20023)


_ = tmp20024

tmp20025 := MakeNative(func(__e *ControlFlow) {
Y986 := __e.Get(1)
_ = Y986
__e.TailApply(PrimFunc(symundefmacro), Y986)
return
}, 1)

tmp20026 := PrimCons(symundefmacro, tmp20025)

tmp20027 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp20026)


_ = tmp20027

tmp20028 := MakeNative(func(__e *ControlFlow) {
Y983 := __e.Get(1)
_ = Y983
__e.Return(MakeNative(func(__e *ControlFlow) {
Y984 := __e.Get(1)
_ = Y984
__e.Return(MakeNative(func(__e *ControlFlow) {
Y985 := __e.Get(1)
_ = Y985
__e.TailApply(PrimFunc(symunput), Y983, Y984, Y985)
return
}, 1))
return
}, 1))
return
}, 1)

tmp20029 := PrimCons(symunput, tmp20028)

tmp20030 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp20029)


_ = tmp20030

tmp20031 := MakeNative(func(__e *ControlFlow) {
Y982 := __e.Get(1)
_ = Y982
__e.TailApply(PrimFunc(symunprofile), Y982)
return
}, 1)

tmp20032 := PrimCons(symunprofile, tmp20031)

tmp20033 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp20032)


_ = tmp20033

tmp20034 := MakeNative(func(__e *ControlFlow) {
Y980 := __e.Get(1)
_ = Y980
__e.Return(MakeNative(func(__e *ControlFlow) {
Y981 := __e.Get(1)
_ = Y981
__e.TailApply(PrimFunc(symunion), Y980, Y981)
return
}, 1))
return
}, 1)

tmp20035 := PrimCons(symunion, tmp20034)

tmp20036 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp20035)


_ = tmp20036

tmp20037 := MakeNative(func(__e *ControlFlow) {
Y979 := __e.Get(1)
_ = Y979
__e.TailApply(PrimFunc(symuntrack), Y979)
return
}, 1)

tmp20038 := PrimCons(symuntrack, tmp20037)

tmp20039 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp20038)


_ = tmp20039

tmp20040 := MakeNative(func(__e *ControlFlow) {
Y978 := __e.Get(1)
_ = Y978
__e.TailApply(PrimFunc(symundefmacro), Y978)
return
}, 1)

tmp20041 := PrimCons(symundefmacro, tmp20040)

tmp20042 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp20041)


_ = tmp20042

tmp20043 := MakeNative(func(__e *ControlFlow) {
Y976 := __e.Get(1)
_ = Y976
__e.Return(MakeNative(func(__e *ControlFlow) {
Y977 := __e.Get(1)
_ = Y977
__e.TailApply(PrimFunc(symupdate_1lambda_1table), Y976, Y977)
return
}, 1))
return
}, 1)

tmp20044 := PrimCons(symupdate_1lambda_1table, tmp20043)

tmp20045 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp20044)


_ = tmp20045

tmp20046 := MakeNative(func(__e *ControlFlow) {
Y975 := __e.Get(1)
_ = Y975
__e.TailApply(PrimFunc(symvector), Y975)
return
}, 1)

tmp20047 := PrimCons(symvector, tmp20046)

tmp20048 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp20047)


_ = tmp20048

tmp20049 := MakeNative(func(__e *ControlFlow) {
Y974 := __e.Get(1)
_ = Y974
__e.TailApply(PrimFunc(symvector_2), Y974)
return
}, 1)

tmp20050 := PrimCons(symvector_2, tmp20049)

tmp20051 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp20050)


_ = tmp20051

tmp20052 := MakeNative(func(__e *ControlFlow) {
Y971 := __e.Get(1)
_ = Y971
__e.Return(MakeNative(func(__e *ControlFlow) {
Y972 := __e.Get(1)
_ = Y972
__e.Return(MakeNative(func(__e *ControlFlow) {
Y973 := __e.Get(1)
_ = Y973
__e.TailApply(PrimFunc(symvector_1_6), Y971, Y972, Y973)
return
}, 1))
return
}, 1))
return
}, 1)

tmp20053 := PrimCons(symvector_1_6, tmp20052)

tmp20054 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp20053)


_ = tmp20054

tmp20055 := MakeNative(func(__e *ControlFlow) {
Y970 := __e.Get(1)
_ = Y970
__e.Return(PrimValue(Y970))
return
}, 1)

tmp20056 := PrimCons(symvalue, tmp20055)

tmp20057 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp20056)


_ = tmp20057

tmp20058 := MakeNative(func(__e *ControlFlow) {
Y969 := __e.Get(1)
_ = Y969
__e.Return(PrimIsVariable(Y969))
return
}, 1)

tmp20059 := PrimCons(symvariable_2, tmp20058)

tmp20060 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp20059)


_ = tmp20060

tmp20061 := MakeNative(func(__e *ControlFlow) {
Y964 := __e.Get(1)
_ = Y964
__e.Return(MakeNative(func(__e *ControlFlow) {
Y965 := __e.Get(1)
_ = Y965
__e.Return(MakeNative(func(__e *ControlFlow) {
Y966 := __e.Get(1)
_ = Y966
__e.Return(MakeNative(func(__e *ControlFlow) {
Y967 := __e.Get(1)
_ = Y967
__e.Return(MakeNative(func(__e *ControlFlow) {
Y968 := __e.Get(1)
_ = Y968
__e.TailApply(PrimFunc(symvar_2), Y964, Y965, Y966, Y967, Y968)
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20062 := PrimCons(symvar_2, tmp20061)

tmp20063 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp20062)


_ = tmp20063

tmp20064 := MakeNative(func(__e *ControlFlow) {
Y959 := __e.Get(1)
_ = Y959
__e.Return(MakeNative(func(__e *ControlFlow) {
Y960 := __e.Get(1)
_ = Y960
__e.Return(MakeNative(func(__e *ControlFlow) {
Y961 := __e.Get(1)
_ = Y961
__e.Return(MakeNative(func(__e *ControlFlow) {
Y962 := __e.Get(1)
_ = Y962
__e.Return(MakeNative(func(__e *ControlFlow) {
Y963 := __e.Get(1)
_ = Y963
__e.TailApply(PrimFunc(symwhen), Y959, Y960, Y961, Y962, Y963)
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp20065 := PrimCons(symwhen, tmp20064)

tmp20066 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp20065)


_ = tmp20066

tmp20067 := MakeNative(func(__e *ControlFlow) {
Y957 := __e.Get(1)
_ = Y957
__e.Return(MakeNative(func(__e *ControlFlow) {
Y958 := __e.Get(1)
_ = Y958
__e.Return(PrimWriteByte(Y957, Y958))
return
}, 1))
return
}, 1)

tmp20068 := PrimCons(symwrite_1byte, tmp20067)

tmp20069 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp20068)


_ = tmp20069

tmp20070 := MakeNative(func(__e *ControlFlow) {
Y955 := __e.Get(1)
_ = Y955
__e.Return(MakeNative(func(__e *ControlFlow) {
Y956 := __e.Get(1)
_ = Y956
__e.TailApply(PrimFunc(symwrite_1to_1file), Y955, Y956)
return
}, 1))
return
}, 1)

tmp20071 := PrimCons(symwrite_1to_1file, tmp20070)

tmp20072 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp20071)


_ = tmp20072

tmp20073 := MakeNative(func(__e *ControlFlow) {
Y954 := __e.Get(1)
_ = Y954
__e.TailApply(PrimFunc(symy_1or_1n_2), Y954)
return
}, 1)

tmp20074 := PrimCons(symy_1or_1n_2, tmp20073)

tmp20075 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp20074)


_ = tmp20075

tmp20076 := MakeNative(func(__e *ControlFlow) {
Y952 := __e.Get(1)
_ = Y952
__e.Return(MakeNative(func(__e *ControlFlow) {
Y953 := __e.Get(1)
_ = Y953
__e.Return(PrimNumberAdd(Y952, Y953))
return
}, 1))
return
}, 1)

tmp20077 := PrimCons(sym_7, tmp20076)

tmp20078 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp20077)


_ = tmp20078

tmp20079 := MakeNative(func(__e *ControlFlow) {
Y950 := __e.Get(1)
_ = Y950
__e.Return(MakeNative(func(__e *ControlFlow) {
Y951 := __e.Get(1)
_ = Y951
__e.Return(PrimNumberMultiply(Y950, Y951))
return
}, 1))
return
}, 1)

tmp20080 := PrimCons(sym_d, tmp20079)

tmp20081 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp20080)


_ = tmp20081

tmp20082 := MakeNative(func(__e *ControlFlow) {
Y948 := __e.Get(1)
_ = Y948
__e.Return(MakeNative(func(__e *ControlFlow) {
Y949 := __e.Get(1)
_ = Y949
__e.Return(PrimNumberDivide(Y948, Y949))
return
}, 1))
return
}, 1)

tmp20083 := PrimCons(sym_c, tmp20082)

tmp20084 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp20083)


_ = tmp20084

tmp20085 := MakeNative(func(__e *ControlFlow) {
Y946 := __e.Get(1)
_ = Y946
__e.Return(MakeNative(func(__e *ControlFlow) {
Y947 := __e.Get(1)
_ = Y947
__e.Return(PrimNumberSubtract(Y946, Y947))
return
}, 1))
return
}, 1)

tmp20086 := PrimCons(sym_1, tmp20085)

tmp20087 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp20086)


_ = tmp20087

tmp20088 := MakeNative(func(__e *ControlFlow) {
Y944 := __e.Get(1)
_ = Y944
__e.Return(MakeNative(func(__e *ControlFlow) {
Y945 := __e.Get(1)
_ = Y945
__e.TailApply(PrimFunc(sym_a_a), Y944, Y945)
return
}, 1))
return
}, 1)

tmp20089 := PrimCons(sym_a_a, tmp20088)

tmp20090 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp20089)


_ = tmp20090

tmp20091 := MakeNative(func(__e *ControlFlow) {
Y943 := __e.Get(1)
_ = Y943
__e.TailApply(PrimFunc(sym_5e_6), Y943)
return
}, 1)

tmp20092 := PrimCons(sym_5e_6, tmp20091)

tmp20093 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp20092)


_ = tmp20093

tmp20094 := MakeNative(func(__e *ControlFlow) {
Y942 := __e.Get(1)
_ = Y942
__e.TailApply(PrimFunc(sym_5end_6), Y942)
return
}, 1)

tmp20095 := PrimCons(sym_5end_6, tmp20094)

tmp20096 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp20095)


_ = tmp20096

tmp20097 := MakeNative(func(__e *ControlFlow) {
Y941 := __e.Get(1)
_ = Y941
__e.TailApply(PrimFunc(sym_5_b_6), Y941)
return
}, 1)

tmp20098 := PrimCons(sym_5_b_6, tmp20097)

tmp20099 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp20098)


_ = tmp20099

tmp20100 := MakeNative(func(__e *ControlFlow) {
Y939 := __e.Get(1)
_ = Y939
__e.Return(MakeNative(func(__e *ControlFlow) {
Y940 := __e.Get(1)
_ = Y940
__e.TailApply(PrimFunc(sym_8p), Y939, Y940)
return
}, 1))
return
}, 1)

tmp20101 := PrimCons(sym_8p, tmp20100)

tmp20102 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp20101)


_ = tmp20102

tmp20103 := MakeNative(func(__e *ControlFlow) {
Y937 := __e.Get(1)
_ = Y937
__e.Return(MakeNative(func(__e *ControlFlow) {
Y938 := __e.Get(1)
_ = Y938
__e.TailApply(PrimFunc(sym_8v), Y937, Y938)
return
}, 1))
return
}, 1)

tmp20104 := PrimCons(sym_8v, tmp20103)

tmp20105 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp20104)


_ = tmp20105

tmp20106 := MakeNative(func(__e *ControlFlow) {
Y935 := __e.Get(1)
_ = Y935
__e.Return(MakeNative(func(__e *ControlFlow) {
Y936 := __e.Get(1)
_ = Y936
__e.TailApply(PrimFunc(sym_8s), Y935, Y936)
return
}, 1))
return
}, 1)

tmp20107 := PrimCons(sym_8s, tmp20106)

__e.TailApply(PrimFunc(symshen_4set_1lambda_1form_1entry), tmp20107)
return


}, 0)

tmp20108 := Call(__e, ns2_1set, symshen_4initialise_1lambda_1forms, tmp19565)


_ = tmp20108

tmp20109 := MakeNative(func(__e *ControlFlow) {
tmp20110 := Call(__e, PrimFunc(symshen_4initialise_1environment))


_ = tmp20110

tmp20111 := Call(__e, PrimFunc(symshen_4initialise_1lambda_1forms))


_ = tmp20111

__e.TailApply(PrimFunc(symshen_4initialise_1signedfuncs))
return


}, 0)

__e.TailApply(ns2_1set, symshen_4initialise, tmp20109)
return




}, 0)

var symshen_4choicepoint_b = MakeSymbol("shen.choicepoint!")
var symshen_4unpackage = MakeSymbol("shen.unpackage")
var symctxt = MakeSymbol("ctxt")
var symshen_4unprotect = MakeSymbol("shen.unprotect")
var symshen_4member = MakeSymbol("shen.member")
var symshen_4l_1rules = MakeSymbol("shen.l-rules")
var symshen_4write_1string = MakeSymbol("shen.write-string")
var symtime = MakeSymbol("time")
var symshen_4ticket_1number = MakeSymbol("shen.ticket-number")
var symshen_4source = MakeSymbol("shen.source")
var symshen_4walk = MakeSymbol("shen.walk")
var symshen_4compile_1prolog = MakeSymbol("shen.compile-prolog")
var sym_c_4 = MakeSymbol("/.")
var symshen_4_ddemodulation_1function_d = MakeSymbol("shen.*demodulation-function*")
var symshen_4_dnames_d = MakeSymbol("shen.*names*")
var symshen_4track_1function = MakeSymbol("shen.track-function")
var symis_b = MakeSymbol("is!")
var symAssumptions = MakeSymbol("Assumptions")
var symspy = MakeSymbol("spy")
var symshen_4r = MakeSymbol("shen.r")
var symshen_4bind_b = MakeSymbol("shen.bind!")
var symshen_4_5sym_6 = MakeSymbol("shen.<sym>")
var symP = MakeSymbol("P")
var symsum = MakeSymbol("sum")
var symshen_4rectify_1test = MakeSymbol("shen.rectify-test")
var symatom_2 = MakeSymbol("atom?")
var symshen_4_5colon_1equal_6 = MakeSymbol("shen.<colon-equal>")
var symshen_4str_1_6str = MakeSymbol("shen.str->str")
var symshen_4linearise_1clause = MakeSymbol("shen.linearise-clause")
var symshen_4prolog_1fbody = MakeSymbol("shen.prolog-fbody")
var symshen_4freshen_1sig = MakeSymbol("shen.freshen-sig")
var symoccurs_2 = MakeSymbol("occurs?")
var symporters = MakeSymbol("porters")
var symshen_4parse_1failure_2 = MakeSymbol("shen.parse-failure?")
var symshen_4compound_1pattern = MakeSymbol("shen.compound-pattern")
var symshen_4_5rrb_6 = MakeSymbol("shen.<rrb>")
var symvar_2 = MakeSymbol("var?")
var symshen_4system_1S = MakeSymbol("shen.system-S")
var symshen_4ok = MakeSymbol("shen.ok")
var sym_dlanguage_d = MakeSymbol("*language*")
var symshen_4iter_1vector = MakeSymbol("shen.iter-vector")
var sym_5end_6 = MakeSymbol("<end>")
var symread_1file_1as_1string = MakeSymbol("read-file-as-string")
var symshen_4peek_1history = MakeSymbol("shen.peek-history")
var symshen_4output_1track = MakeSymbol("shen.output-track")
var symsuccess = MakeSymbol("success")
var symshen_4_dsystem_d = MakeSymbol("shen.*system*")
var symunion = MakeSymbol("union")
var symshen_4_5return_6 = MakeSymbol("shen.<return>")
var symshen_4overapplication_2 = MakeSymbol("shen.overapplication?")
var symshen_4process_1lambda = MakeSymbol("shen.process-lambda")
var sym_6 = MakeSymbol(">")
var symshen_4freshterm = MakeSymbol("shen.freshterm")
var symshen_4mkstr_1r = MakeSymbol("shen.mkstr-r")
var symdefun = MakeSymbol("defun")
var symstr = MakeSymbol("str")
var symshen_4top = MakeSymbol("shen.top")
var symreturn = MakeSymbol("return")
var symshen_4_5ass_6 = MakeSymbol("shen.<ass>")
var symshen_4search_1user_1datatypes = MakeSymbol("shen.search-user-datatypes")
var symsubst = MakeSymbol("subst")
var symshen_4_5rule_d_6 = MakeSymbol("shen.<rule*>")
var symshen_4_5syntax_6 = MakeSymbol("shen.<syntax>")
var symshen_4analyse_1symbol_2 = MakeSymbol("shen.analyse-symbol?")
var symhash = MakeSymbol("hash")
var symshen_4pvar = MakeSymbol("shen.pvar")
var symshen_4invoke = MakeSymbol("shen.invoke")
var symshen_4variancy = MakeSymbol("shen.variancy")
var symfail_1if = MakeSymbol("fail-if")
var symshen_4_5iscolon_6 = MakeSymbol("shen.<iscolon>")
var symshen_4coll_1formulae = MakeSymbol("shen.coll-formulae")
var symboolean = MakeSymbol("boolean")
var symshen_4x_4launcher_4script_1command = MakeSymbol("shen.x.launcher.script-command")
var symshen_4_dtracking_d = MakeSymbol("shen.*tracking*")
var sym_dimplementation_d = MakeSymbol("*implementation*")
var symshen_4shendef_1_6kldef_1h = MakeSymbol("shen.shendef->kldef-h")
var symdatatypes = MakeSymbol("datatypes")
var symshen_4objectcode = MakeSymbol("shen.objectcode")
var symshen_4dict_1rm = MakeSymbol("shen.dict-rm")
var symshen_4iter_1list = MakeSymbol("shen.iter-list")
var symshen_4compile_1pattern_1fragment = MakeSymbol("shen.compile-pattern-fragment")
var symshen_4kl_1body = MakeSymbol("shen.kl-body")
var symerror = MakeSymbol("error")
var symshen_4unwind_1types = MakeSymbol("shen.unwind-types")
var symshen_4_5prem_6 = MakeSymbol("shen.<prem>")
var symAssumption = MakeSymbol("Assumption")
var symerror_1to_1string = MakeSymbol("error-to-string")
var symtc_2 = MakeSymbol("tc?")
var symshen_4internal_1to_1shen_2 = MakeSymbol("shen.internal-to-shen?")
var symshen_4typecheck = MakeSymbol("shen.typecheck")
var symshen_4_5minus_6 = MakeSymbol("shen.<minus>")
var symshen_4fn_1call_2 = MakeSymbol("shen.fn-call?")
var sym_3 = MakeSymbol("$")
var sympreclude = MakeSymbol("preclude")
var symshen_4unpackage_emacroexpand = MakeSymbol("shen.unpackage&macroexpand")
var symshen_4byte_1_6digit = MakeSymbol("shen.byte->digit")
var symshen_4lazyderef = MakeSymbol("shen.lazyderef")
var symshen_4op1 = MakeSymbol("shen.op1")
var symshen_4char_1stoutput_2 = MakeSymbol("shen.char-stoutput?")
var symshen_4_5patterns_6 = MakeSymbol("shen.<patterns>")
var symshen_4x_4launcher_4done = MakeSymbol("shen.x.launcher.done")
var symV = MakeSymbol("V")
var symshen_4sng_1h_2 = MakeSymbol("shen.sng-h?")
var symshen_4_5datatype_6 = MakeSymbol("shen.<datatype>")
var symshen_4_5hterm1_6 = MakeSymbol("shen.<hterm1>")
var symtc = MakeSymbol("tc")
var symcompile = MakeSymbol("compile")
var symshen_4unwind = MakeSymbol("shen.unwind")
var sym_5_b_6 = MakeSymbol("<!>")
var symshen_4_5pattern2_6 = MakeSymbol("shen.<pattern2>")
var symlineread = MakeSymbol("lineread")
var symshen_4nextticket = MakeSymbol("shen.nextticket")
var symshen_4custom_1pattern_1compiler = MakeSymbol("shen.custom-pattern-compiler")
var symshen_4expt = MakeSymbol("shen.expt")
var symshen_4record_1external = MakeSymbol("shen.record-external")
var symshen_4custom_1pattern_1body = MakeSymbol("shen.custom-pattern-body")
var symnumber_2 = MakeSymbol("number?")
var symshen_4_5lowC_6 = MakeSymbol("shen.<lowC>")
var symshen_4by_1hypothesis = MakeSymbol("shen.by-hypothesis")
var symshen_4t_d_1rule = MakeSymbol("shen.t*-rule")
var symshen_4x_4launcher_4eval_1string = MakeSymbol("shen.x.launcher.eval-string")
var symshen_4application_2 = MakeSymbol("shen.application?")
var symnumber = MakeSymbol("number")
var symshen_4arg_1_6str = MakeSymbol("shen.arg->str")
var symshen_4free_1variable_1error_1message = MakeSymbol("shen.free-variable-error-message")
var symshen_4_5control_6 = MakeSymbol("shen.<control>")
var symshen_4_5lowE_6 = MakeSymbol("shen.<lowE>")
var sympackage = MakeSymbol("package")
var symshen_4check_1eval_1and_1print = MakeSymbol("shen.check-eval-and-print")
var symshen_4demod = MakeSymbol("shen.demod")
var symloaded = MakeSymbol("loaded")
var symshen_4_5semicolon_6 = MakeSymbol("shen.<semicolon>")
var symshen_4_5lsb_6 = MakeSymbol("shen.<lsb>")
var symshen_4partial_1application_d_2 = MakeSymbol("shen.partial-application*?")
var symprofile = MakeSymbol("profile")
var symshen_4_5c_1rule_6 = MakeSymbol("shen.<c-rule>")
var symshen_4dict_1bucket_1_6 = MakeSymbol("shen.dict-bucket->")
var symshen_4x_4launcher_4help_1text = MakeSymbol("shen.x.launcher.help-text")
var symshen_4sysfunc_2 = MakeSymbol("shen.sysfunc?")
var symshen_4_5comma_6 = MakeSymbol("shen.<comma>")
var symprotect = MakeSymbol("protect")
var symshen_4factor_1recognisors = MakeSymbol("shen.factor-recognisors")
var symshen_4foreign_2 = MakeSymbol("shen.foreign?")
var symshen_4non_1terminalcode = MakeSymbol("shen.non-terminalcode")
var sym_c = MakeSymbol("/")
var symshen_4_5iscomma_6 = MakeSymbol("shen.<iscomma>")
var symexternal = MakeSymbol("external")
var symshen_4extraspecial_2 = MakeSymbol("shen.extraspecial?")
var symread_1byte = MakeSymbol("read-byte")
var symMessage = MakeSymbol("Message")
var symshen_4write_1kl = MakeSymbol("shen.write-kl")
var symabsvector = MakeSymbol("absvector")
var symshen_4initialise = MakeSymbol("shen.initialise")
var symabort = MakeSymbol("abort")
var symshen_4nothing_1doing_2 = MakeSymbol("shen.nothing-doing?")
var symshen_4included = MakeSymbol("shen.included")
var symshen_4_5type_6 = MakeSymbol("shen.<type>")
var symshen_4spy_2 = MakeSymbol("shen.spy?")
var symsave = MakeSymbol("save")
var symremove = MakeSymbol("remove")
var symshen_4assumetypes = MakeSymbol("shen.assumetypes")
var symshen_4_8v_1help = MakeSymbol("shen.@v-help")
var symshen_4rcons__form = MakeSymbol("shen.rcons_form")
var symY = MakeSymbol("Y")
var symshen_4_5multiline_6 = MakeSymbol("shen.<multiline>")
var symshen_4_dcall_d = MakeSymbol("shen.*call*")
var symwarn = MakeSymbol("warn")
var symshen_4_dempty_1absvector_d = MakeSymbol("shen.*empty-absvector*")
var symshen_4parse_1failure = MakeSymbol("shen.parse-failure")
var symshen_4arity_1chk = MakeSymbol("shen.arity-chk")
var symshen_4_5pattern_6 = MakeSymbol("shen.<pattern>")
var symFreeze = MakeSymbol("Freeze")
var symshen_4lambda_1form = MakeSymbol("shen.lambda-form")
var symshen_4prompt = MakeSymbol("shen.prompt")
var symshen_4assert_d = MakeSymbol("shen.assert*")
var symcases = MakeSymbol("cases")
var symshen_4function_1calls = MakeSymbol("shen.function-calls")
var symshen_4f = MakeSymbol("shen.f")
var symshen_4conscode = MakeSymbol("shen.conscode")
var symcond = MakeSymbol("cond")
var symis = MakeSymbol("is")
var symshen_4tame = MakeSymbol("shen.tame")
var symshen_4rule_1_6body = MakeSymbol("shen.rule->body")
var sym_4_4_4 = MakeSymbol("...")
var symhush_2 = MakeSymbol("hush?")
var sym_dmacros_d = MakeSymbol("*macros*")
var symshen_4overbind = MakeSymbol("shen.overbind")
var symshen_4type_1theory_1enabled_2 = MakeSymbol("shen.type-theory-enabled?")
var symstring = MakeSymbol("string")
var symexception = MakeSymbol("exception")
var symshen_4_dmaxinferences_d = MakeSymbol("shen.*maxinferences*")
var symshen_4atom_1_6str = MakeSymbol("shen.atom->str")
var symshen_4_dcustom_1pattern_1compiler_d = MakeSymbol("shen.*custom-pattern-compiler*")
var symvariable_2 = MakeSymbol("variable?")
var symsterror = MakeSymbol("sterror")
var symshen_4_5pattern1_6 = MakeSymbol("shen.<pattern1>")
var symshen_4cons_1form = MakeSymbol("shen.cons-form")
var symshen_4_5backslash_6 = MakeSymbol("shen.<backslash>")
var symshen_4zero_1place_2 = MakeSymbol("shen.zero-place?")
var symshen_4read_1evaluate_1print = MakeSymbol("shen.read-evaluate-print")
var symshen_4_5non_1terminal_2_6 = MakeSymbol("shen.<non-terminal?>")
var sym_dversion_d = MakeSymbol("*version*")
var symshen_4constructor_1error = MakeSymbol("shen.constructor-error")
var symshen_4find_1free_1vars = MakeSymbol("shen.find-free-vars")
var symshen_4x_4launcher_4version_1string = MakeSymbol("shen.x.launcher.version-string")
var sym_5_1address = MakeSymbol("<-address")
var sym_e_e = MakeSymbol("&&")
var symshen_4string_1prefix_2 = MakeSymbol("shen.string-prefix?")
var symshen_4prolog_1parameters = MakeSymbol("shen.prolog-parameters")
var symunknown_1arguments = MakeSymbol("unknown-arguments")
var symshen_4_7string_2 = MakeSymbol("shen.+string?")
var symshen_4_5shortnatter_6 = MakeSymbol("shen.<shortnatter>")
var symshen_4typename_1h = MakeSymbol("shen.typename-h")
var symK = MakeSymbol("K")
var symshen_4klfile = MakeSymbol("shen.klfile")
var symshen_4non_1application_2 = MakeSymbol("shen.non-application?")
var symshen_4variable_1case = MakeSymbol("shen.variable-case")
var symshen_4_5semantics_6 = MakeSymbol("shen.<semantics>")
var symParse = MakeSymbol("Parse")
var symshen_4_5digit_6 = MakeSymbol("shen.<digit>")
var symshen_4macro_1_8c = MakeSymbol("shen.macro-@c")
var symshen_4signal_1def = MakeSymbol("shen.signal-def")
var symshen_4internal_1symbols = MakeSymbol("shen.internal-symbols")
var symshen_4_5define_6 = MakeSymbol("shen.<define>")
var symshen_4lchh = MakeSymbol("shen.lchh")
var sym_j = MakeSymbol("}")
var sympackage_2 = MakeSymbol("package?")
var symshen_4lambda_1function = MakeSymbol("shen.lambda-function")
var symshen_4build_1lambda_1table = MakeSymbol("shen.build-lambda-table")
var symshen_4eval_1and_1print = MakeSymbol("shen.eval-and-print")
var symshen_4rule_1_6clause = MakeSymbol("shen.rule->clause")
var symshen_4nvars = MakeSymbol("shen.nvars")
var symshen_4x_4launcher_4execute_1all = MakeSymbol("shen.x.launcher.execute-all")
var symhdstr = MakeSymbol("hdstr")
var sym_1 = MakeSymbol("-")
var symshen_4loop = MakeSymbol("shen.loop")
var symmake_1string = MakeSymbol("make-string")
var symfail = MakeSymbol("fail")
var symfactorise_2 = MakeSymbol("factorise?")
var symshen_4tlv_1help = MakeSymbol("shen.tlv-help")
var symdifference = MakeSymbol("difference")
var symshen_4factor = MakeSymbol("shen.factor")
var symshen_4_8ch = MakeSymbol("shen.@ch")
var symshen_4_5head_6 = MakeSymbol("shen.<head>")
var symshen_4_5packagename_6 = MakeSymbol("shen.<packagename>")
var symshen_4alphanums_2 = MakeSymbol("shen.alphanums?")
var symZ = MakeSymbol("Z")
var symshen_4x_4launcher_4default_1handle_1result = MakeSymbol("shen.x.launcher.default-handle-result")
var symtail = MakeSymbol("tail")
var symadjoin = MakeSymbol("adjoin")
var symshen_4print_1vector_2 = MakeSymbol("shen.print-vector?")
var symshen_4dictionary = MakeSymbol("shen.dictionary")
var symshen_4cond_1form = MakeSymbol("shen.cond-form")
var symshen_4write_1kl_1h = MakeSymbol("shen.write-kl-h")
var symshen_4not_1pvar = MakeSymbol("shen.not-pvar")
var symshen_4external_1symbols = MakeSymbol("shen.external-symbols")
var symshen_4linearise = MakeSymbol("shen.linearise")
var symshen_4_5sides_6 = MakeSymbol("shen.<sides>")
var symfn = MakeSymbol("fn")
var symshen_4specialise_1consume = MakeSymbol("shen.specialise-consume")
var symshen_4_5stop_6 = MakeSymbol("shen.<stop>")
var symshen_4intern_1in_1package = MakeSymbol("shen.intern-in-package")
var symshen_4newname = MakeSymbol("shen.newname")
var symshen_4rules_1_6prolog = MakeSymbol("shen.rules->prolog")
var symshen_4macro_1_8ch = MakeSymbol("shen.macro-@ch")
var symshen_4show = MakeSymbol("shen.show")
var symshen_4_5comment_6 = MakeSymbol("shen.<comment>")
var symabsolute = MakeSymbol("absolute")
var symshen_4_5s_1exprs_6 = MakeSymbol("shen.<s-exprs>")
var symshen_4consume_1clause = MakeSymbol("shen.consume-clause")
var symshen_4free_1variable_2 = MakeSymbol("shen.free-variable?")
var symshen_4_5lcurly_6 = MakeSymbol("shen.<lcurly>")
var symshen_4internal_2 = MakeSymbol("shen.internal?")
var symshen_4dbl_2 = MakeSymbol("shen.dbl?")
var symC = MakeSymbol("C")
var symstep = MakeSymbol("step")
var symnth = MakeSymbol("nth")
var symoptimise = MakeSymbol("optimise")
var symshen_4initialise__environment = MakeSymbol("shen.initialise_environment")
var symshen_4_5bterms_6 = MakeSymbol("shen.<bterms>")
var symshen_4goto = MakeSymbol("shen.goto")
var symshen_4_5times_6 = MakeSymbol("shen.<times>")
var symshen_4process_1_8s = MakeSymbol("shen.process-@s")
var symbind = MakeSymbol("bind")
var symprofile_1results = MakeSymbol("profile-results")
var symshen_4map_1h = MakeSymbol("shen.map-h")
var symshen_4analyse_1variable_2 = MakeSymbol("shen.analyse-variable?")
var symshen_4create_1skeleton = MakeSymbol("shen.create-skeleton")
var symshen_4record_1and_1evaluate = MakeSymbol("shen.record-and-evaluate")
var symshen_4process_1sexprs = MakeSymbol("shen.process-sexprs")
var symshen_4char_1stinput_2 = MakeSymbol("shen.char-stinput?")
var symshen_4lzy_a_b = MakeSymbol("shen.lzy=!")
var symdefprolog = MakeSymbol("defprolog")
var symtrack = MakeSymbol("track")
var symshen_4p_1hyps = MakeSymbol("shen.p-hyps")
var symdo = MakeSymbol("do")
var symreverse = MakeSymbol("reverse")
var symshen_4_dspecial_d = MakeSymbol("shen.*special*")
var symResult = MakeSymbol("Result")
var symFinish = MakeSymbol("Finish")
var symshen_4monomorphic_2 = MakeSymbol("shen.monomorphic?")
var symlaunch_1repl = MakeSymbol("launch-repl")
var symfreeze = MakeSymbol("freeze")
var symshen_4simple_1curry = MakeSymbol("shen.simple-curry")
var symshen_4vector_1dereference = MakeSymbol("shen.vector-dereference")
var symshen_4_ddatatypes_d = MakeSymbol("shen.*datatypes*")
var symshen_4show_1assumptions = MakeSymbol("shen.show-assumptions")
var symshen_4funexstring = MakeSymbol("shen.funexstring")
var symshen_4_5alphanums_6 = MakeSymbol("shen.<alphanums>")
var sym_dproperty_1vector_d = MakeSymbol("*property-vector*")
var symrelease = MakeSymbol("release")
var symshen_4constructor_2 = MakeSymbol("shen.constructor?")
var symshen_4findall_1h = MakeSymbol("shen.findall-h")
var symshen_4this_1symbol_1is_1unbound = MakeSymbol("shen.this-symbol-is-unbound")
var symwrite_1to_1file = MakeSymbol("write-to-file")
var symshen_4i_1failed_b = MakeSymbol("shen.i-failed!")
var symshen_4wildcard_2 = MakeSymbol("shen.wildcard?")
var symshen_4cons_1case_1plus = MakeSymbol("shen.cons-case-plus")
var sympreclude_1all_1but = MakeSymbol("preclude-all-but")
var sympos = MakeSymbol("pos")
var symshen_4app = MakeSymbol("shen.app")
var symhdv = MakeSymbol("hdv")
var symcons_2 = MakeSymbol("cons?")
var syminput = MakeSymbol("input")
var sym_5_1_1 = MakeSymbol("<--")
var symshen_4callrec = MakeSymbol("shen.callrec")
var symshen_4_5datatype_1rule_6 = MakeSymbol("shen.<datatype-rule>")
var symAction = MakeSymbol("Action")
var syminput_7 = MakeSymbol("input+")
var sym_1_6 = MakeSymbol("->")
var symread_1file = MakeSymbol("read-file")
var symshen_4_5body_6 = MakeSymbol("shen.<body>")
var symshen_4incinfs = MakeSymbol("shen.incinfs")
var symshen_4side_1conditions_1_6goals = MakeSymbol("shen.side-conditions->goals")
var symshen_4freshen_1type = MakeSymbol("shen.freshen-type")
var sym_dsterror_d = MakeSymbol("*sterror*")
var symshen_4lock = MakeSymbol("shen.lock")
var symlanguage = MakeSymbol("language")
var symshen_4insert_1info = MakeSymbol("shen.insert-info")
var symshen_4_5clause_6 = MakeSymbol("shen.<clause>")
var symshen_4t_d_1rule_1h = MakeSymbol("shen.t*-rule-h")
var symconcat = MakeSymbol("concat")
var symenable_1type_1theory = MakeSymbol("enable-type-theory")
var symshen_4bad_1pivot_2 = MakeSymbol("shen.bad-pivot?")
var symset = MakeSymbol("set")
var symshen_4extract_1vars = MakeSymbol("shen.extract-vars")
var symprolog_1memory = MakeSymbol("prolog-memory")
var symshen_4pac_1h = MakeSymbol("shen.pac-h")
var symoccurs_1check = MakeSymbol("occurs-check")
var symshen_4partial_1parse_1failure_2 = MakeSymbol("shen.partial-parse-failure?")
var symshen_4_5packagenames_6 = MakeSymbol("shen.<packagenames>")
var symbar_b = MakeSymbol("bar!")
var symshen_4_dfactorise_2_d = MakeSymbol("shen.*factorise?*")
var symshen_4abs = MakeSymbol("shen.abs")
var sym_7 = MakeSymbol("+")
var symforeign = MakeSymbol("foreign")
var symfresh = MakeSymbol("fresh")
var symshen_4_doptimise_d = MakeSymbol("shen.*optimise*")
var symversion = MakeSymbol("version")
var symshen_4add_1sexpr = MakeSymbol("shen.add-sexpr")
var symshen_4passive_1variables = MakeSymbol("shen.passive-variables")
var symshen_4prolog_1abstraction = MakeSymbol("shen.prolog-abstraction")
var symunit = MakeSymbol("unit")
var symshen_4compile_1to_1kl = MakeSymbol("shen.compile-to-kl")
var symif = MakeSymbol("if")
var symshen_4alpha_1convert = MakeSymbol("shen.alpha-convert")
var symGoTo = MakeSymbol("GoTo")
var symshen_4call_1prolog = MakeSymbol("shen.call-prolog")
var sym_dporters_d = MakeSymbol("*porters*")
var symshen_4_5e_1number_6 = MakeSymbol("shen.<e-number>")
var symlimit = MakeSymbol("limit")
var symshen_4linearise_1h = MakeSymbol("shen.linearise-h")
var symtl = MakeSymbol("tl")
var symshen_4lowercase_2 = MakeSymbol("shen.lowercase?")
var symshen_4macros = MakeSymbol("shen.macros")
var syminteger_2 = MakeSymbol("integer?")
var symshen_4prolog_1vector = MakeSymbol("shen.prolog-vector")
var symshen_4explode_1h = MakeSymbol("shen.explode-h")
var symmode = MakeSymbol("mode")
var symshen_4hascut_2 = MakeSymbol("shen.hascut?")
var symshen_4dict_1count = MakeSymbol("shen.dict-count")
var sym_5 = MakeSymbol("<")
var symshen_4internal_1to_1P_2 = MakeSymbol("shen.internal-to-P?")
var syminline = MakeSymbol("inline")
var symshen_4copyfromvector = MakeSymbol("shen.copyfromvector")
var symshen_4_5str_6 = MakeSymbol("shen.<str>")
var symshen_4process_1datatype = MakeSymbol("shen.process-datatype")
var symshen_4process_1let = MakeSymbol("shen.process-let")
var symshen_4list_1_6str = MakeSymbol("shen.list->str")
var symshen_4retract_1clause = MakeSymbol("shen.retract-clause")
var symshen_4t_d_1correct = MakeSymbol("shen.t*-correct")
var symshen_4remove_1indirection = MakeSymbol("shen.remove-indirection")
var symshen_4make_1uppercase = MakeSymbol("shen.make-uppercase")
var symHypotheses = MakeSymbol("Hypotheses")
var symuntrack = MakeSymbol("untrack")
var symshen_4yacc_1semantics = MakeSymbol("shen.yacc-semantics")
var symKey = MakeSymbol("Key")
var symexplode = MakeSymbol("explode")
var symos = MakeSymbol("os")
var symshen_4return_2 = MakeSymbol("shen.return?")
var symshen_4find_1arity = MakeSymbol("shen.find-arity")
var symshen_4curry_1type = MakeSymbol("shen.curry-type")
var symunprofile = MakeSymbol("unprofile")
var symshen_4op2 = MakeSymbol("shen.op2")
var symshen_4_5notdbq_6 = MakeSymbol("shen.<notdbq>")
var symshen_4string_1match = MakeSymbol("shen.string-match")
var symshen_4process_1assoc = MakeSymbol("shen.process-assoc")
var symshen_4newpv = MakeSymbol("shen.newpv")
var symshen_4fbound_2 = MakeSymbol("shen.fbound?")
var symshen_4_5numeral_6 = MakeSymbol("shen.<numeral>")
var symshen_4deref_1calls = MakeSymbol("shen.deref-calls")
var symshen_4_5sig_drules_6 = MakeSymbol("shen.<sig*rules>")
var symshen_4_5number_6 = MakeSymbol("shen.<number>")
var symX = MakeSymbol("X")
var symshen_4digit_2 = MakeSymbol("shen.digit?")
var symoccurrences = MakeSymbol("occurrences")
var symshen_4dynamic = MakeSymbol("shen.dynamic")
var symshen_4x_4launcher_4eval_1flag_1map = MakeSymbol("shen.x.launcher.eval-flag-map")
var symprint = MakeSymbol("print")
var symshen_4out_1of_1bounds = MakeSymbol("shen.out-of-bounds")
var symshen_4shendef_1_6kldef = MakeSymbol("shen.shendef->kldef")
var symshen_4_5rsb_6 = MakeSymbol("shen.<rsb>")
var symassertz = MakeSymbol("assertz")
var symshen_4extract_1free_1vars = MakeSymbol("shen.extract-free-vars")
var symshen_4_5non_1terminal_1name_6 = MakeSymbol("shen.<non-terminal-name>")
var symshen_4dynamic_1default = MakeSymbol("shen.dynamic-default")
var symshen_4initialise_1signedfuncs = MakeSymbol("shen.initialise-signedfuncs")
var symsystemf = MakeSymbol("systemf")
var symshen_4pvar_2 = MakeSymbol("shen.pvar?")
var symshen_4_5conc_6 = MakeSymbol("shen.<conc>")
var symshen_4sigf = MakeSymbol("shen.sigf")
var symhead = MakeSymbol("head")
var symshen_4proc_1nl = MakeSymbol("shen.proc-nl")
var symshen_4_5constructor_6 = MakeSymbol("shen.<constructor>")
var symshen_4compute_1integer_1h = MakeSymbol("shen.compute-integer-h")
var symaddress_1_6 = MakeSymbol("address->")
var symshen_4alpha_2 = MakeSymbol("shen.alpha?")
var symshen_4_dit_d = MakeSymbol("shen.*it*")
var symshen_4stpart = MakeSymbol("shen.stpart")
var symshen_4lzy_a = MakeSymbol("shen.lzy=")
var symshen_4terpri_1or_1read_1char = MakeSymbol("shen.terpri-or-read-char")
var symshen_4find_1types = MakeSymbol("shen.find-types")
var sym_a = MakeSymbol("=")
var symundefmacro = MakeSymbol("undefmacro")
var symshen_4cut = MakeSymbol("shen.cut")
var symshen_4remove_1h = MakeSymbol("shen.remove-h")
var symshen_4demodulate = MakeSymbol("shen.demodulate")
var symshen_4fix_1help = MakeSymbol("shen.fix-help")
var symsynonyms = MakeSymbol("synonyms")
var symshen_4special_2 = MakeSymbol("shen.special?")
var symn_1_6string = MakeSymbol("n->string")
var symshen_4trim_1it = MakeSymbol("shen.trim-it")
var symprolog_2 = MakeSymbol("prolog?")
var symshen_4_5formulae_6 = MakeSymbol("shen.<formulae>")
var symshen_4_5dbq_6 = MakeSymbol("shen.<dbq>")
var symappend = MakeSymbol("append")
var syminferences = MakeSymbol("inferences")
var symshen_4typename = MakeSymbol("shen.typename")
var symshen_4t_d_1integrity = MakeSymbol("shen.t*-integrity")
var symshen_4dict = MakeSymbol("shen.dict")
var symshen_4x_4launcher_4eval_1command_1h = MakeSymbol("shen.x.launcher.eval-command-h")
var sym_8v = MakeSymbol("@v")
var symshen_4string_1_6bytes = MakeSymbol("shen.string->bytes")
var symread_1from_1string_1unprocessed = MakeSymbol("read-from-string-unprocessed")
var symshen_4special_1case = MakeSymbol("shen.special-case")
var symassoc = MakeSymbol("assoc")
var symshen_4unpack_1foreign = MakeSymbol("shen.unpack-foreign")
var symshen_4type_1F = MakeSymbol("shen.type-F")
var symstream = MakeSymbol("stream")
var symshow_1help = MakeSymbol("show-help")
var symshen_4factorise_1code = MakeSymbol("shen.factorise-code")
var symshen_4primitive = MakeSymbol("shen.primitive")
var symtrap_1error = MakeSymbol("trap-error")
var symshen_4c_1rules_1_6shen = MakeSymbol("shen.c-rules->shen")
var symsnd = MakeSymbol("snd")
var symshen_4update_1assoc = MakeSymbol("shen.update-assoc")
var symshen_4_dprofiled_d = MakeSymbol("shen.*profiled*")
var symshen_4_5rcurly_6 = MakeSymbol("shen.<rcurly>")
var symreceive = MakeSymbol("receive")
var symcons = MakeSymbol("cons")
var symempty_2 = MakeSymbol("empty?")
var symshen_4read_1loop = MakeSymbol("shen.read-loop")
var symshen = MakeSymbol("shen")
var sym_e = MakeSymbol("&")
var symshen_4t = MakeSymbol("shen.t")
var symshen_4empty_1absvector_2 = MakeSymbol("shen.empty-absvector?")
var symshen_4compile_1head = MakeSymbol("shen.compile-head")
var symshen_4freeze_1literals = MakeSymbol("shen.freeze-literals")
var symshen_4rule_1_6head = MakeSymbol("shen.rule->head")
var symstoutput = MakeSymbol("stoutput")
var sym_dos_d = MakeSymbol("*os*")
var symshen_4process_1application = MakeSymbol("shen.process-application")
var symfunction = MakeSymbol("function")
var symu_b = MakeSymbol("u!")
var symshen_4_5formula_6 = MakeSymbol("shen.<formula>")
var symshen_4wildcardcode = MakeSymbol("shen.wildcardcode")
var symW = MakeSymbol("W")
var symshen_4monotype = MakeSymbol("shen.monotype")
var symshen_4_5_1out = MakeSymbol("shen.<-out")
var symstinput = MakeSymbol("stinput")
var symshen_4mkstr_1l = MakeSymbol("shen.mkstr-l")
var symmacroexpand = MakeSymbol("macroexpand")
var symeval_1kl = MakeSymbol("eval-kl")
var sym_dhush_d = MakeSymbol("*hush*")
var symshen_4x_4launcher_4launch_1shen = MakeSymbol("shen.x.launcher.launch-shen")
var symshen_4shen_1_6kl = MakeSymbol("shen.shen->kl")
var symshen_4string_1_6byte = MakeSymbol("shen.string->byte")
var symshen_4terms = MakeSymbol("shen.terms")
var symshen_4str_1_6bytes = MakeSymbol("shen.str->bytes")
var symshen_4tracked_2 = MakeSymbol("shen.tracked?")
var symshen_4maxinfexceeded_2 = MakeSymbol("shen.maxinfexceeded?")
var symshen_4line = MakeSymbol("shen.line")
var symshen_4yacc_1syntax = MakeSymbol("shen.yacc-syntax")
var symget = MakeSymbol("get")
var symintern = MakeSymbol("intern")
var symshen_4process_1applications = MakeSymbol("shen.process-applications")
var symcn = MakeSymbol("cn")
var symshen_4choicepoint_2 = MakeSymbol("shen.choicepoint?")
var symshen_4remember_1datatype = MakeSymbol("shen.remember-datatype")
var symunspecialise = MakeSymbol("unspecialise")
var symshen_4reader_1error = MakeSymbol("shen.reader-error")
var symshen_4specialise_1member = MakeSymbol("shen.specialise-member")
var symshen_4processed = MakeSymbol("shen.processed")
var symshen_4parameters = MakeSymbol("shen.parameters")
var symshen_4work_1through = MakeSymbol("shen.work-through")
var symshen_4_dsize_1prolog_1vector_d = MakeSymbol("shen.*size-prolog-vector*")
var symgensym = MakeSymbol("gensym")
var symshen_4_dpackage_d = MakeSymbol("shen.*package*")
var symshen_4tuple_1up = MakeSymbol("shen.tuple-up")
var symshen_4freshen = MakeSymbol("shen.freshen")
var symshen_4decons = MakeSymbol("shen.decons")
var syminternal = MakeSymbol("internal")
var symshen_4read_1file_1as_1bytelist_1help = MakeSymbol("shen.read-file-as-bytelist-help")
var symshen_4_5literal_6 = MakeSymbol("shen.<literal>")
var symshen_4magless = MakeSymbol("shen.magless")
var symrun = MakeSymbol("run")
var symshen_4lch = MakeSymbol("shen.lch")
var symshen_4profiled_2 = MakeSymbol("shen.profiled?")
var sym_dmaximum_1print_1sequence_1size_d = MakeSymbol("*maximum-print-sequence-size*")
var symshen_4f_1error = MakeSymbol("shen.f-error")
var symunput = MakeSymbol("unput")
var symshen_4_5strcontents_6 = MakeSymbol("shen.<strcontents>")
var symshen_4atom_1case_1minus = MakeSymbol("shen.atom-case-minus")
var syminclude = MakeSymbol("include")
var symshen_4put_1profile = MakeSymbol("shen.put-profile")
var symshen_4process_1time = MakeSymbol("shen.process-time")
var symshen_4_5hterm2_6 = MakeSymbol("shen.<hterm2>")
var symshen_4_5prems_6 = MakeSymbol("shen.<prems>")
var symshen_4_5syntax_1item_6 = MakeSymbol("shen.<syntax-item>")
var symsymbol = MakeSymbol("symbol")
var symshen_4shen_1_6kl_1h = MakeSymbol("shen.shen->kl-h")
var symshen_4printF = MakeSymbol("shen.printF")
var symshen_4load_1help = MakeSymbol("shen.load-help")
var symshen_4lr_1rule = MakeSymbol("shen.lr-rule")
var symshen_4t_d = MakeSymbol("shen.t*")
var symshen_4show_1p = MakeSymbol("shen.show-p")
var symshen_4a = MakeSymbol("shen.a")
var symshen_4_5lrb_6 = MakeSymbol("shen.<lrb>")
var symshen_4bytes_1_6string = MakeSymbol("shen.bytes->string")
var symshen_4_5rules_d_6 = MakeSymbol("shen.<rules*>")
var symshen_4correct = MakeSymbol("shen.correct")
var sym_dabsolute_d = MakeSymbol("*absolute*")
var symshen_4rectify_1type = MakeSymbol("shen.rectify-type")
var symshen_4deref_1forked_1literals = MakeSymbol("shen.deref-forked-literals")
var symshen_4deref = MakeSymbol("shen.deref")
var sym_5_1vector = MakeSymbol("<-vector")
var symdeclare = MakeSymbol("declare")
var symabsvector_2 = MakeSymbol("absvector?")
var symlength = MakeSymbol("length")
var symshen_4mkstr = MakeSymbol("shen.mkstr")
var symshen_4_5equal_6 = MakeSymbol("shen.<equal>")
var symshen_4_dspy_d = MakeSymbol("shen.*spy*")
var symopen = MakeSymbol("open")
var symshen_4_5fraction_6 = MakeSymbol("shen.<fraction>")
var symshen_4vector_1_6str = MakeSymbol("shen.vector->str")
var symshen_4read_1unit_1string = MakeSymbol("shen.read-unit-string")
var symshen_4_5single_6 = MakeSymbol("shen.<single>")
var symshen_4profile_1help = MakeSymbol("shen.profile-help")
var symshen_4dict_1values = MakeSymbol("shen.dict-values")
var symshen_4dict_1count_1_6 = MakeSymbol("shen.dict-count->")
var symy_1or_1n_2 = MakeSymbol("y-or-n?")
var symelement_2 = MakeSymbol("element?")
var symbound_2 = MakeSymbol("bound?")
var symshen_4undefined_1f_2 = MakeSymbol("shen.undefined-f?")
var symshen_4_dsynonyms_d = MakeSymbol("shen.*synonyms*")
var symshen_4_5side_6 = MakeSymbol("shen.<side>")
var symshen_4initialise_1lambda_1forms = MakeSymbol("shen.initialise-lambda-forms")
var symshen_4_5yacc_6 = MakeSymbol("shen.<yacc>")
var symshen_4choicepoint = MakeSymbol("shen.choicepoint")
var symshen_4compute_1E = MakeSymbol("shen.compute-E")
var symmap = MakeSymbol("map")
var symshen_4_dgensym_d = MakeSymbol("shen.*gensym*")
var symshen_4print_1prolog_1vector = MakeSymbol("shen.print-prolog-vector")
var symshen_4_5atom_6 = MakeSymbol("shen.<atom>")
var symshen_4_dhistory_d = MakeSymbol("shen.*history*")
var symA = MakeSymbol("A")
var symshen_4_5wildcard_6 = MakeSymbol("shen.<wildcard>")
var symshen_4insert_1tracking_1code = MakeSymbol("shen.insert-tracking-code")
var symshen_4toplevel_1forms = MakeSymbol("shen.toplevel-forms")
var symarity = MakeSymbol("arity")
var symshen_4beta = MakeSymbol("shen.beta")
var symshen_4process_1cond_1clauses = MakeSymbol("shen.process-cond-clauses")
var symshen_4evaluate_1lineread = MakeSymbol("shen.evaluate-lineread")
var symStart = MakeSymbol("Start")
var symshen_4show_1datatypes = MakeSymbol("shen.show-datatypes")
var symshen_4prterm = MakeSymbol("shen.prterm")
var symshen_4cons_1form_1respect_1modes = MakeSymbol("shen.cons-form-respect-modes")
var symwhen = MakeSymbol("when")
var symfst = MakeSymbol("fst")
var symshen_4_5dbl_6 = MakeSymbol("shen.<dbl>")
var symshen_4print_1freshterm = MakeSymbol("shen.print-freshterm")
var symoutput = MakeSymbol("output")
var symshen_4input_1track = MakeSymbol("shen.input-track")
var symshen_4dict_1capacity = MakeSymbol("shen.dict-capacity")
var symshen_4x_4launcher_4quiet_1load = MakeSymbol("shen.x.launcher.quiet-load")
var symshen_4dict_1_6 = MakeSymbol("shen.dict->")
var symshen_4mod = MakeSymbol("shen.mod")
var symshen_4ctxt = MakeSymbol("shen.ctxt")
var symmaxinferences = MakeSymbol("maxinferences")
var symshen_4_5longnatter_6 = MakeSymbol("shen.<longnatter>")
var symshen_4raise_1syntax_1error = MakeSymbol("shen.raise-syntax-error")
var sym_dhome_1directory_d = MakeSymbol("*home-directory*")
var symshen_4_5rule_6 = MakeSymbol("shen.<rule>")
var symshen_4type_1error = MakeSymbol("shen.type-error")
var symshen_4atom_1case_1plus = MakeSymbol("shen.atom-case-plus")
var symshen_4construct_1context = MakeSymbol("shen.construct-context")
var symshen_4_dalldatatypes_d = MakeSymbol("shen.*alldatatypes*")
var symshen_4locked_2 = MakeSymbol("shen.locked?")
var symshen_4toplevel_1display_1exception = MakeSymbol("shen.toplevel-display-exception")
var symshen_4key_1in_1sequent_1calculus_2 = MakeSymbol("shen.key-in-sequent-calculus?")
var symshen_4s = MakeSymbol("shen.s")
var symnull = MakeSymbol("null")
var symboolean_2 = MakeSymbol("boolean?")
var symput = MakeSymbol("put")
var symshen_4insert = MakeSymbol("shen.insert")
var symshen_4factor_1cn = MakeSymbol("shen.factor-cn")
var symtype = MakeSymbol("type")
var symNewV = MakeSymbol("NewV")
var symcd = MakeSymbol("cd")
var symthaw = MakeSymbol("thaw")
var symshen_4_5simple_1pattern_6 = MakeSymbol("shen.<simple-pattern>")
var symshen_4partial = MakeSymbol("shen.partial")
var symshen_4_5bterm_6 = MakeSymbol("shen.<bterm>")
var symshen_4myassume = MakeSymbol("shen.myassume")
var symshen_4_5signature_6 = MakeSymbol("shen.<signature>")
var symshen_4whitespace_2 = MakeSymbol("shen.whitespace?")
var symshen_4_5strc_6 = MakeSymbol("shen.<strc>")
var symshen_4record_1internal = MakeSymbol("shen.record-internal")
var symshen_4process_1read_1byte = MakeSymbol("shen.process-read-byte")
var sym_d = MakeSymbol("*")
var symor = MakeSymbol("or")
var symshen_4intern_1type = MakeSymbol("shen.intern-type")
var sym_a_a_6 = MakeSymbol("==>")
var symshen_4freshen_1rule = MakeSymbol("shen.freshen-rule")
var symshen_4bucket_1fold = MakeSymbol("shen.bucket-fold")
var sympr = MakeSymbol("pr")
var symshen_4insert_1h = MakeSymbol("shen.insert-h")
var symshen_4fits_2 = MakeSymbol("shen.fits?")
var symlambda = MakeSymbol("lambda")
var symshen_4package_1symbols = MakeSymbol("shen.package-symbols")
var symload = MakeSymbol("load")
var symshen_4eos = MakeSymbol("shen.eos")
var symshen_4compute_1integer = MakeSymbol("shen.compute-integer")
var symshen_4compute_1fraction_1h = MakeSymbol("shen.compute-fraction-h")
var symshen_4gc = MakeSymbol("shen.gc")
var symshen_4passive_1bind = MakeSymbol("shen.passive-bind")
var symshen_4sigxrules = MakeSymbol("shen.sigxrules")
var symshen_4packaged_2 = MakeSymbol("shen.packaged?")
var sym_6_6 = MakeSymbol(">>")
var symS = MakeSymbol("S")
var symlazy = MakeSymbol("lazy")
var symmapcan = MakeSymbol("mapcan")
var symhd = MakeSymbol("hd")
var symshen_4_5whitespace_6 = MakeSymbol("shen.<whitespace>")
var symretract = MakeSymbol("retract")
var symshen_4horn_1clause_1procedure = MakeSymbol("shen.horn-clause-procedure")
var symnl = MakeSymbol("nl")
var symshen_4_dloading_2_d = MakeSymbol("shen.*loading?*")
var symshen_4lookupsig = MakeSymbol("shen.lookupsig")
var sym_dargv_d = MakeSymbol("*argv*")
var symshen_4goto_1h = MakeSymbol("shen.goto-h")
var symclose = MakeSymbol("close")
var sym_a_a = MakeSymbol("==")
var symshen_4_dresidue_d = MakeSymbol("shen.*residue*")
var symshen_4dict_1keys = MakeSymbol("shen.dict-keys")
var symshen_4write_1chars = MakeSymbol("shen.write-chars")
var symshen_4custom_1pattern_1reducer = MakeSymbol("shen.custom-pattern-reducer")
var symand = MakeSymbol("and")
var symdefmacro = MakeSymbol("defmacro")
var sym_1_1_6 = MakeSymbol("-->")
var symfindall = MakeSymbol("findall")
var symshen_4_dstep_d = MakeSymbol("shen.*step*")
var symshen_4non_1terminal_2 = MakeSymbol("shen.non-terminal?")
var symshen_4set_1lambda_1form_1entry = MakeSymbol("shen.set-lambda-form-entry")
var symshen_4fn_1print = MakeSymbol("shen.fn-print")
var symlet = MakeSymbol("let")
var symshen_4_7vector_2 = MakeSymbol("shen.+vector?")
var symTime = MakeSymbol("Time")
var symshen_4unlocked_2 = MakeSymbol("shen.unlocked?")
var symshen_4dict_1update_1count = MakeSymbol("shen.dict-update-count")
var symshen_4custom_1pattern_2 = MakeSymbol("shen.custom-pattern?")
var symshen_4call_1dynamic = MakeSymbol("shen.call-dynamic")
var symtracked = MakeSymbol("tracked")
var symvector_1_6 = MakeSymbol("vector->")
var symshen_4macroexpand_1h = MakeSymbol("shen.macroexpand-h")
var symshen_4received = MakeSymbol("shen.received")
var symshen_4remove_1datatypes = MakeSymbol("shen.remove-datatypes")
var symunabsolute = MakeSymbol("unabsolute")
var symshen_4reverse_1help = MakeSymbol("shen.reverse-help")
var symshen_4_5singleline_6 = MakeSymbol("shen.<singleline>")
var symshen_4_5float_6 = MakeSymbol("shen.<float>")
var symshen_4freshterms = MakeSymbol("shen.freshterms")
var symshen_4for_1each = MakeSymbol("shen.for-each")
var symshen_4free_1var_1chk = MakeSymbol("shen.free-var-chk")
var symshen_4op_1test = MakeSymbol("shen.op-test")
var symshen_4yacc_1_6shen = MakeSymbol("shen.yacc->shen")
var symNewAssumptions = MakeSymbol("NewAssumptions")
var symshen_4dict_2 = MakeSymbol("shen.dict?")
var symshen_4prhush = MakeSymbol("shen.prhush")
var symshen_4record_1kl = MakeSymbol("shen.record-kl")
var symshen_4pretty_1type = MakeSymbol("shen.pretty-type")
var symB = MakeSymbol("B")
var symshen_4_doccurs_d = MakeSymbol("shen.*occurs*")
var symshen_4modh = MakeSymbol("shen.modh")
var symshen_4_dinfs_d = MakeSymbol("shen.*infs*")
var symread_1from_1string = MakeSymbol("read-from-string")
var symshen_4prolog_1vector_1size = MakeSymbol("shen.prolog-vector-size")
var symshen_4custom_1pattern_1triple_1stack = MakeSymbol("shen.custom-pattern-triple-stack")
var symshen_4_5sc_6 = MakeSymbol("shen.<sc>")
var symshen_4_5c_1rules_6 = MakeSymbol("shen.<c-rules>")
var symshen_4update_1history = MakeSymbol("shen.update-history")
var symshen_4pause_1for_1user = MakeSymbol("shen.pause-for-user")
var symshen_4colon_1equal_2 = MakeSymbol("shen.colon-equal?")
var sym_dport_d = MakeSymbol("*port*")
var symshen_4store_1arity = MakeSymbol("shen.store-arity")
var symshen_4_1null_1 = MakeSymbol("shen.-null-")
var symshen_4curry = MakeSymbol("shen.curry")
var symshen_4semicolon_2 = MakeSymbol("shen.semicolon?")
var symfork = MakeSymbol("fork")
var symshen_4process_1after_1type = MakeSymbol("shen.process-after-type")
var symshen_4lambda_1of_1defun = MakeSymbol("shen.lambda-of-defun")
var symshen_4deref_1terms = MakeSymbol("shen.deref-terms")
var symshen_4scan_1body = MakeSymbol("shen.scan-body")
var symSelect = MakeSymbol("Select")
var symshen_4my_1read_1byte = MakeSymbol("shen.my-read-byte")
var symshen_4member_1clause = MakeSymbol("shen.member-clause")
var symshen_4variants_2 = MakeSymbol("shen.variants?")
var symshen_4pui_1h = MakeSymbol("shen.pui-h")
var symshen_4autocomplete = MakeSymbol("shen.autocomplete")
var symshen_4skip = MakeSymbol("shen.skip")
var symshen_4demode = MakeSymbol("shen.demode")
var symshen_4c_1rule_1_6shen = MakeSymbol("shen.c-rule->shen")
var symshen_4prodbutzero = MakeSymbol("shen.prodbutzero")
var symtlstr = MakeSymbol("tlstr")
var symshen_4make_1prolog_1variable = MakeSymbol("shen.make-prolog-variable")
var symshen_4spaces = MakeSymbol("shen.spaces")
var symshen_4assoc_1set = MakeSymbol("shen.assoc-set")
var symshen_4_5_1dict = MakeSymbol("shen.<-dict")
var symshen_4compile_1synonyms = MakeSymbol("shen.compile-synonyms")
var symshen_4bindv = MakeSymbol("shen.bindv")
var symshen_4decrement_1ticket = MakeSymbol("shen.decrement-ticket")
var symshen_4freshterm_2 = MakeSymbol("shen.freshterm?")
var symport = MakeSymbol("port")
var symshen_4prolog_1keyword_2 = MakeSymbol("shen.prolog-keyword?")
var sym_8s = MakeSymbol("@s")
var sym__ = MakeSymbol("_")
var symshen_4try_1parse = MakeSymbol("shen.try-parse")
var symshen_4repl = MakeSymbol("shen.repl")
var symdatatype = MakeSymbol("datatype")
var symget_1time = MakeSymbol("get-time")
var symshen_4_5sng_6 = MakeSymbol("shen.<sng>")
var sym_i = MakeSymbol("{")
var symshen_4_5bar_6 = MakeSymbol("shen.<bar>")
var symshen_4misc_2 = MakeSymbol("shen.misc?")
var symshen_4_5hterm_6 = MakeSymbol("shen.<hterm>")
var symshen_4dbl_1h_2 = MakeSymbol("shen.dbl-h?")
var symshen_4unix = MakeSymbol("shen.unix")
var symshen_4assoc_1rm = MakeSymbol("shen.assoc-rm")
var symshen_4sng_2 = MakeSymbol("shen.sng?")
var symshen_4_8c = MakeSymbol("shen.@c")
var symshen_4dict_1fold = MakeSymbol("shen.dict-fold")
var sym_8p = MakeSymbol("@p")
var symnot = MakeSymbol("not")
var symimplementation = MakeSymbol("implementation")
var symshen_4_dprolog_1memory_d = MakeSymbol("shen.*prolog-memory*")
var symvector_2 = MakeSymbol("vector?")
var sym_5_1 = MakeSymbol("<-")
var symin = MakeSymbol("in")
var symshen_4profile_1func = MakeSymbol("shen.profile-func")
var symshen_4uppercase_2 = MakeSymbol("shen.uppercase?")
var symshen_4prolog_1track = MakeSymbol("shen.prolog-track")
var symshen_4list_2 = MakeSymbol("shen.list?")
var symbootstrap = MakeSymbol("bootstrap")
var symshen_4prolog_1arity_1check = MakeSymbol("shen.prolog-arity-check")
var symshen_4_5log10_6 = MakeSymbol("shen.<log10>")
var symstring_1_6symbol = MakeSymbol("string->symbol")
var symshen_4rep_1X = MakeSymbol("shen.rep-X")
var symread = MakeSymbol("read")
var symshen_4_dsigf_d = MakeSymbol("shen.*sigf*")
var symwhere = MakeSymbol("where")
var symshen_4loading_2 = MakeSymbol("shen.loading?")
var symasserta = MakeSymbol("asserta")
var syminclude_1all_1but = MakeSymbol("include-all-but")
var symHd = MakeSymbol("Hd")
var sym_5_a = MakeSymbol("<=")
var symvalue = MakeSymbol("value")
var symshen_4_5rules_6 = MakeSymbol("shen.<rules>")
var symshen_4assoc_1_6 = MakeSymbol("shen.assoc->")
var symshen_4_5plus_6 = MakeSymbol("shen.<plus>")
var symshen_4_5name_6 = MakeSymbol("shen.<name>")
var symshen_4_5colon_6 = MakeSymbol("shen.<colon>")
var symshen_4initialise_1arity_1table = MakeSymbol("shen.initialise-arity-table")
var symit = MakeSymbol("it")
var symshen_4_5alphanum_6 = MakeSymbol("shen.<alphanum>")
var symshen_4maxseq = MakeSymbol("shen.maxseq")
var symshen_4fillvector = MakeSymbol("shen.fillvector")
var symshen_4compile_1body = MakeSymbol("shen.compile-body")
var symuserdefs = MakeSymbol("userdefs")
var symsystem_1S_2 = MakeSymbol("system-S?")
var symshen_4syntax_1item_2 = MakeSymbol("shen.syntax-item?")
var symshen_4_7m = MakeSymbol("shen.+m")
var symin_1package = MakeSymbol("in-package")
var symshen_4custom_1pattern = MakeSymbol("shen.custom-pattern")
var symshen_4pivot_1on = MakeSymbol("shen.pivot-on")
var symshen_4rfas_1h = MakeSymbol("shen.rfas-h")
var symshen_4ccons_2 = MakeSymbol("shen.ccons?")
var symshen_4consume = MakeSymbol("shen.consume")
var symshen_4g = MakeSymbol("shen.g")
var symshen_4bottom = MakeSymbol("shen.bottom")
var symshen_4combine_1c_1code = MakeSymbol("shen.combine-c-code")
var symoptimise_2 = MakeSymbol("optimise?")
var symshen_4unassoc = MakeSymbol("shen.unassoc")
var symshen_4multiples = MakeSymbol("shen.multiples")
var symshen_4length_1h = MakeSymbol("shen.length-h")
var symshen_4op = MakeSymbol("shen.op")
var symshen_4reader_1error_1message = MakeSymbol("shen.reader-error-message")
var symshen_4unlock = MakeSymbol("shen.unlock")
var symTm = MakeSymbol("Tm")
var symverified = MakeSymbol("verified")
var symshen_4terminalcode = MakeSymbol("shen.terminalcode")
var symincluded = MakeSymbol("included")
var symps = MakeSymbol("ps")
var symout = MakeSymbol("out")
var symshen_4_dextraspecial_d = MakeSymbol("shen.*extraspecial*")
var sym_6_a = MakeSymbol(">=")
var symsimple_1error = MakeSymbol("simple-error")
var symshen_4remove_1bystanders = MakeSymbol("shen.remove-bystanders")
var symshen_4prtl = MakeSymbol("shen.prtl")
var symshen_4find_1arities = MakeSymbol("shen.find-arities")
var symshen_4update_1lambdatable = MakeSymbol("shen.update-lambdatable")
var symL = MakeSymbol("L")
var symRemainder = MakeSymbol("Remainder")
var symeval = MakeSymbol("eval")
var symshen_4_5whitespaces_6 = MakeSymbol("shen.<whitespaces>")
var symshen_4compute_1fraction = MakeSymbol("shen.compute-fraction")
var symshen_4mu_1h = MakeSymbol("shen.mu-h")
var symshen_4_5clauses_6 = MakeSymbol("shen.<clauses>")
var symshen_4t_d_1rules = MakeSymbol("shen.t*-rules")
var symtuple_2 = MakeSymbol("tuple?")
var symshen_4_5s_1exprs2_6 = MakeSymbol("shen.<s-exprs2>")
var symshen_4_5shortnatters_6 = MakeSymbol("shen.<shortnatters>")
var symshen_4recursive_1string_1match = MakeSymbol("shen.recursive-string-match")
var symshen_4insert_1prolog_1variables = MakeSymbol("shen.insert-prolog-variables")
var symshen_4system_1S_1h = MakeSymbol("shen.system-S-h")
var sym_dstoutput_d = MakeSymbol("*stoutput*")
var symshen_4_5s_1exprs1_6 = MakeSymbol("shen.<s-exprs1>")
var symshen_4process_1cases = MakeSymbol("shen.process-cases")
var symshen_4_5defprolog_6 = MakeSymbol("shen.<defprolog>")
var symshen_4continue = MakeSymbol("shen.continue")
var symshen_4_1m = MakeSymbol("shen.-m")
var symshen_4_dtc_d = MakeSymbol("shen.*tc*")
var symsymbol_2 = MakeSymbol("symbol?")
var symshen_4tuple = MakeSymbol("shen.tuple")
var symshen_4factor_1selectors_1h = MakeSymbol("shen.factor-selectors-h")
var symshen_4fn_1call = MakeSymbol("shen.fn-call")
var symTl = MakeSymbol("Tl")
var symshen_4use_1type_1info = MakeSymbol("shen.use-type-info")
var symshen_4record_1it = MakeSymbol("shen.record-it")
var symshen_4process_1synonyms = MakeSymbol("shen.process-synonyms")
var symshen_4synonyms_1h = MakeSymbol("shen.synonyms-h")
var symstring_2 = MakeSymbol("string?")
var symshen_4openlock = MakeSymbol("shen.openlock")
var symshen_4process_1yacc_1semantics = MakeSymbol("shen.process-yacc-semantics")
var symshen_4_5integer_6 = MakeSymbol("shen.<integer>")
var symshen_4_5expr_6 = MakeSymbol("shen.<expr>")
var symshen_4hds_a_2 = MakeSymbol("shen.hds=?")
var symshen_4execute_1store_1arity = MakeSymbol("shen.execute-store-arity")
var symshen_4_5returns_6 = MakeSymbol("shen.<returns>")
var symshen_4package_1user_1input = MakeSymbol("shen.package-user-input")
var symRecord = MakeSymbol("Record")
var symshen_4x_4launcher_4eval_1command = MakeSymbol("shen.x.launcher.eval-command")
var symshen_4initialise_1environment = MakeSymbol("shen.initialise-environment")
var symshen_4posint_2 = MakeSymbol("shen.posint?")
var symupdate_1lambda_1table = MakeSymbol("update-lambda-table")
var sym_b = MakeSymbol("!")
var symshen_4recursively_1print = MakeSymbol("shen.recursively-print")
var symshen_4get_1profile = MakeSymbol("shen.get-profile")
var symshen_4not_1dictionary = MakeSymbol("shen.not-dictionary")
var symdestroy = MakeSymbol("destroy")
var sym_5e_6 = MakeSymbol("<e>")
var symshen_4x_4launcher_4main = MakeSymbol("shen.x.launcher.main")
var symfile = MakeSymbol("file")
var symshen_4not_1tuple = MakeSymbol("shen.not-tuple")
var symshen_4vector_1parameter = MakeSymbol("shen.vector-parameter")
var symwrite_1byte = MakeSymbol("write-byte")
var symshen_4hashkey = MakeSymbol("shen.hashkey")
var symintersection = MakeSymbol("intersection")
var symfix = MakeSymbol("fix")
var symshen_4in_1_6 = MakeSymbol("shen.in->")
var symshen_4_5hash_6 = MakeSymbol("shen.<hash>")
var symshen_4typetable = MakeSymbol("shen.typetable")
var symshen_4premises_1_6goals = MakeSymbol("shen.premises->goals")
var symvector = MakeSymbol("vector")
var symshen_4shen_1call_2 = MakeSymbol("shen.shen-call?")
var symshen_4integer_1test_2 = MakeSymbol("shen.integer-test?")
var symshen_4_duserdefs_d = MakeSymbol("shen.*userdefs*")
var symdefine = MakeSymbol("define")
var symshen_4_dcustom_1pattern_1reducer_d = MakeSymbol("shen.*custom-pattern-reducer*")
var symshen_4use_1history = MakeSymbol("shen.use-history")
var symshen_4step_2 = MakeSymbol("shen.step?")
var sym_dstinput_d = MakeSymbol("*stinput*")
var symlist = MakeSymbol("list")
var symshen_4_5double_6 = MakeSymbol("shen.<double>")
var symshen_4check_1byte = MakeSymbol("shen.check-byte")
var symshen_4variablecode = MakeSymbol("shen.variablecode")
var symhush = MakeSymbol("hush")
var symtlv = MakeSymbol("tlv")
var symshen_4comb = MakeSymbol("shen.comb")
var symshen_4credits = MakeSymbol("shen.credits")
var symshen_4occurs_1check_2 = MakeSymbol("shen.occurs-check?")
var symshen_4insert_1l = MakeSymbol("shen.insert-l")
var symread_1file_1as_1bytelist = MakeSymbol("read-file-as-bytelist")
var symshen_4_5_1dict_1bucket = MakeSymbol("shen.<-dict-bucket")
var symshen_4record_1macro = MakeSymbol("shen.record-macro")
var symcall = MakeSymbol("call")
var symshen_4_5datatype_1rules_6 = MakeSymbol("shen.<datatype-rules>")
var symshen_4syntax_1error_1message = MakeSymbol("shen.syntax-error-message")
var symshen_4lambda_1entry = MakeSymbol("shen.lambda-entry")
var symfactorise = MakeSymbol("factorise")
var symshen_4lowercase_1symbol_2 = MakeSymbol("shen.lowercase-symbol?")
var symstring_1_6n = MakeSymbol("string->n")
var symshen_4predicate = MakeSymbol("shen.predicate")
var symshen_4dict_1fold_1h = MakeSymbol("shen.dict-fold-h")
var symshen_4triple_1stack = MakeSymbol("shen.triple-stack")
var symshen_4factor_1selectors = MakeSymbol("shen.factor-selectors")
var symshen_4cons_1case_1minus = MakeSymbol("shen.cons-case-minus")
var symspecialise = MakeSymbol("specialise")
var symshen_4compound_1pattern_1h = MakeSymbol("shen.compound-pattern-h")
var symshen_4_5alpha_6 = MakeSymbol("shen.<alpha>")
var symshen_4_5digits_6 = MakeSymbol("shen.<digits>")
var symshen_4_5yaccsig_6 = MakeSymbol("shen.<yaccsig>")
var symshen_4_5packagechar_6 = MakeSymbol("shen.<packagechar>")
var sym_drelease_d = MakeSymbol("*release*")
var symshen_4_dshen_1type_1theory_1enabled_2_d = MakeSymbol("shen.*shen-type-theory-enabled?*")
var symshen_4process_1def = MakeSymbol("shen.process-def")
var symdefcc = MakeSymbol("defcc")
