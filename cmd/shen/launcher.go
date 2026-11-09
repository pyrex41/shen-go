package main

import . "github.com/pyrex41/shen-go/kl"

var LauncherMain = MakeNative0(func(__e *ControlFlow) Obj {
_ = MakeString("Copyright (c) 2019 Bruno Deferrari.\nBSD 3-Clause License: http://opensource.org/licenses/BSD-3-Clause")

tmp21142 := MakeNative1(func(__e *ControlFlow, V7374 Obj) Obj {
tmp21143 := Call1(__e, SymFn(symread_1file), V7374)


let__21138 := tmp21143
_ = let__21138

tmp21144 := MakeNative1(func(__e *ControlFlow, Z7376 Obj) Obj {
return TailCall1(__e, SymFn(symeval), Z7376)
panic("unreachable")
})

return TailCall2(__e, SymFn(symmap), tmp21144, let__21138)


panic("unreachable")
})

tmp21145 := Call2(__e, ns2_1set, symshen_4x_4launcher_4quiet_1load, tmp21142)


_ = tmp21145

tmp21146 := MakeNative0(func(__e *ControlFlow) Obj {
tmp21147 := Call0(__e, SymFn(symversion))


tmp21148 := Call0(__e, SymFn(symlanguage))


tmp21149 := Call0(__e, SymFn(symport))


tmp21150 := GCons(__e, tmp21149, Nil)

tmp21151 := GCons(__e, tmp21148, tmp21150)

tmp21152 := Call0(__e, SymFn(symimplementation))


tmp21153 := Call0(__e, SymFn(symrelease))


tmp21154 := GCons(__e, tmp21153, Nil)

tmp21155 := GCons(__e, tmp21152, tmp21154)

tmp21156 := GCons(__e, tmp21155, Nil)

tmp21157 := GCons(__e, symimplementation, tmp21156)

tmp21158 := GCons(__e, tmp21151, tmp21157)

tmp21159 := GCons(__e, symport, tmp21158)

tmp21160 := Call3(__e, SymFn(symshen_4app), tmp21159, MakeString("\n"), symshen_4r)


return TailCall3(__e, SymFn(symshen_4app), tmp21147, PrimGuard2(__e, symcn, PrimStringConcat, MakeString(" "), tmp21160), symshen_4a)


panic("unreachable")
})

tmp21162 := Call2(__e, ns2_1set, symshen_4x_4launcher_4version_1string, tmp21146)


_ = tmp21162

tmp21163 := MakeNative1(func(__e *ControlFlow, V7377 Obj) Obj {
tmp21164 := Call3(__e, SymFn(symshen_4app), V7377, MakeString(" [--version] [--help] <COMMAND> [<ARGS>]\n\ncommands:\n    repl\n        Launches the interactive REPL.\n        Default action if no command is supplied.\n\n    script <FILE> [<ARGS>]\n        Runs the script in FILE. *argv* is set to [FILE | ARGS].\n\n    eval <ARGS>\n        Evaluates expressions and files. ARGS are evaluated from\n        left to right and can be a combination of:\n            -e, --eval <EXPR>\n                Evaluates EXPR and prints result.\n            -l, --load <FILE>\n                Reads and evaluates FILE.\n            -q, --quiet\n                Silences interactive output.\n            -s, --set <KEY> <VALUE>\n                Evaluates KEY, VALUE and sets as global.\n            -r, --repl\n                Launches the interactive REPL after evaluating\n                all the previous expresions."), symshen_4a)


return PrimGuard2(__e, symcn, PrimStringConcat, MakeString("Usage: "), tmp21164)


panic("unreachable")
})

tmp21165 := Call2(__e, ns2_1set, symshen_4x_4launcher_4help_1text, tmp21163)


_ = tmp21165

tmp21166 := (func() Obj {
var __self Obj
__self = MakeNative1(func(__e *ControlFlow, __self1 Obj) Obj {
__selftop:
V7378 := __self1
_ = V7378
tmp21173 := GEq(__e, Nil, V7378)

if True == tmp21173 {
return GCons(__e, symsuccess, Nil)
} else {
tmp21171 := GConsP(__e, V7378)

if True == tmp21171 {
tmp21167 := GHd(__e, V7378)

tmp21168 := Call1(__e, SymFn(symthaw), tmp21167)


_ = tmp21168

tmp21169 := GTl(__e, V7378)

if SymFn(symshen_4x_4launcher_4execute_1all) == __self {
__self1 = tmp21169
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4x_4launcher_4execute_1all), tmp21169)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4x_4launcher_4execute_1all)
}


}


panic("unreachable")
})
return __self
})()

tmp21174 := Call2(__e, ns2_1set, symshen_4x_4launcher_4execute_1all, tmp21166)


_ = tmp21174

tmp21175 := MakeNative1(func(__e *ControlFlow, V7379 Obj) Obj {
tmp21176 := Call1(__e, SymFn(symread_1from_1string), V7379)


tmp21177 := Call1(__e, SymFn(symhead), tmp21176)


return TailCall1(__e, SymFn(symeval), tmp21177)


panic("unreachable")
})

tmp21178 := Call2(__e, ns2_1set, symshen_4x_4launcher_4eval_1string, tmp21175)


_ = tmp21178

tmp21179 := MakeNative1(func(__e *ControlFlow, V7382 Obj) Obj {
tmp21189 := GEq(__e, MakeString("-e"), V7382)

if True == tmp21189 {
return MakeString("--eval")
} else {
tmp21187 := GEq(__e, MakeString("-l"), V7382)

if True == tmp21187 {
return MakeString("--load")
} else {
tmp21185 := GEq(__e, MakeString("-q"), V7382)

if True == tmp21185 {
return MakeString("--quiet")
} else {
tmp21183 := GEq(__e, MakeString("-s"), V7382)

if True == tmp21183 {
return MakeString("--set")
} else {
tmp21181 := GEq(__e, MakeString("-r"), V7382)

if True == tmp21181 {
return MakeString("--repl")
} else {
return False
}


}


}


}


}


panic("unreachable")
})

tmp21190 := Call2(__e, ns2_1set, symshen_4x_4launcher_4eval_1flag_1map, tmp21179)


_ = tmp21190

tmp21191 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V7387 := __self1
_ = V7387
V7388 := __self2
_ = V7388
tmp21284 := GEq(__e, Nil, V7387)

if True == tmp21284 {
tmp21192 := Call1(__e, SymFn(symreverse), V7388)


return TailCall1(__e, SymFn(symshen_4x_4launcher_4execute_1all), tmp21192)


} else {
tmp21282 := GConsP(__e, V7387)

var ifres21276 Obj

if True == tmp21282 {
tmp21280 := GHd(__e, V7387)

tmp21281 := GEq(__e, MakeString("--eval"), tmp21280)

var ifres21277 Obj

if True == tmp21281 {
tmp21278 := GTl(__e, V7387)

tmp21279 := GConsP(__e, tmp21278)

ifres21277 = tmp21279


} else {
ifres21277 = False


}

ifres21276 = ifres21277


} else {
ifres21276 = False


}

if True == ifres21276 {
tmp21193 := GTl(__e, V7387)

tmp21194 := GTl(__e, tmp21193)

tmp21195 := MakeNative0(func(__e *ControlFlow) Obj {
tmp21196 := GTl(__e, V7387)

tmp21197 := GHd(__e, tmp21196)

tmp21198 := Call1(__e, SymFn(symshen_4x_4launcher_4eval_1string), tmp21197)


tmp21199 := Call3(__e, SymFn(symshen_4app), tmp21198, MakeString("\n"), symshen_4a)


tmp21200 := Call0(__e, SymFn(symstoutput))


return TailCall2(__e, SymFn(sympr), tmp21199, tmp21200)


panic("unreachable")
})

tmp21201 := GCons(__e, tmp21195, V7388)

if SymFn(symshen_4x_4launcher_4eval_1command_1h) == __self {
__self1, __self2 = tmp21194, tmp21201
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4x_4launcher_4eval_1command_1h), tmp21194, tmp21201)


} else {
tmp21274 := GConsP(__e, V7387)

var ifres21268 Obj

if True == tmp21274 {
tmp21272 := GHd(__e, V7387)

tmp21273 := GEq(__e, MakeString("--load"), tmp21272)

var ifres21269 Obj

if True == tmp21273 {
tmp21270 := GTl(__e, V7387)

tmp21271 := GConsP(__e, tmp21270)

ifres21269 = tmp21271


} else {
ifres21269 = False


}

ifres21268 = ifres21269


} else {
ifres21268 = False


}

if True == ifres21268 {
tmp21202 := GTl(__e, V7387)

tmp21203 := GTl(__e, tmp21202)

tmp21204 := MakeNative0(func(__e *ControlFlow) Obj {
tmp21205 := GTl(__e, V7387)

tmp21206 := GHd(__e, tmp21205)

return TailCall1(__e, SymFn(symload), tmp21206)


panic("unreachable")
})

tmp21207 := GCons(__e, tmp21204, V7388)

if SymFn(symshen_4x_4launcher_4eval_1command_1h) == __self {
__self1, __self2 = tmp21203, tmp21207
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4x_4launcher_4eval_1command_1h), tmp21203, tmp21207)


} else {
tmp21266 := GConsP(__e, V7387)

var ifres21263 Obj

if True == tmp21266 {
tmp21264 := GHd(__e, V7387)

tmp21265 := GEq(__e, MakeString("--quiet"), tmp21264)

ifres21263 = tmp21265


} else {
ifres21263 = False


}

if True == ifres21263 {
tmp21208 := GTl(__e, V7387)

tmp21209 := MakeNative0(func(__e *ControlFlow) Obj {
return PrimGuard2(__e, symset, PrimSet, sym_dhush_d, True)
panic("unreachable")
})

tmp21210 := GCons(__e, tmp21209, V7388)

if SymFn(symshen_4x_4launcher_4eval_1command_1h) == __self {
__self1, __self2 = tmp21208, tmp21210
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4x_4launcher_4eval_1command_1h), tmp21208, tmp21210)


} else {
tmp21261 := GConsP(__e, V7387)

var ifres21251 Obj

if True == tmp21261 {
tmp21259 := GHd(__e, V7387)

tmp21260 := GEq(__e, MakeString("--set"), tmp21259)

var ifres21252 Obj

if True == tmp21260 {
tmp21257 := GTl(__e, V7387)

tmp21258 := GConsP(__e, tmp21257)

var ifres21253 Obj

if True == tmp21258 {
tmp21254 := GTl(__e, V7387)

tmp21255 := GTl(__e, tmp21254)

tmp21256 := GConsP(__e, tmp21255)

ifres21253 = tmp21256


} else {
ifres21253 = False


}

ifres21252 = ifres21253


} else {
ifres21252 = False


}

ifres21251 = ifres21252


} else {
ifres21251 = False


}

if True == ifres21251 {
tmp21211 := GTl(__e, V7387)

tmp21212 := GTl(__e, tmp21211)

tmp21213 := GTl(__e, tmp21212)

tmp21214 := MakeNative0(func(__e *ControlFlow) Obj {
tmp21215 := GTl(__e, V7387)

tmp21216 := GHd(__e, tmp21215)

tmp21217 := Call1(__e, SymFn(symshen_4x_4launcher_4eval_1string), tmp21216)


tmp21218 := GTl(__e, V7387)

tmp21219 := GTl(__e, tmp21218)

tmp21220 := GHd(__e, tmp21219)

tmp21221 := Call1(__e, SymFn(symshen_4x_4launcher_4eval_1string), tmp21220)


return PrimGuard2(__e, symset, PrimSet, tmp21217, tmp21221)


panic("unreachable")
})

tmp21222 := GCons(__e, tmp21214, V7388)

if SymFn(symshen_4x_4launcher_4eval_1command_1h) == __self {
__self1, __self2 = tmp21213, tmp21222
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4x_4launcher_4eval_1command_1h), tmp21213, tmp21222)


} else {
tmp21249 := GConsP(__e, V7387)

var ifres21246 Obj

if True == tmp21249 {
tmp21247 := GHd(__e, V7387)

tmp21248 := GEq(__e, MakeString("--repl"), tmp21247)

ifres21246 = tmp21248


} else {
ifres21246 = False


}

if True == ifres21246 {
tmp21223 := Call2(__e, SymFn(symshen_4x_4launcher_4eval_1command_1h), Nil, V7388)


_ = tmp21223

tmp21224 := GTl(__e, V7387)

return GCons(__e, symlaunch_1repl, tmp21224)


} else {
tmp21225 := MakeNative0(func(__e *ControlFlow) Obj {
tmp21231 := GConsP(__e, V7387)

if True == tmp21231 {
tmp21226 := GHd(__e, V7387)

tmp21227 := Call3(__e, SymFn(symshen_4app), tmp21226, MakeString(""), symshen_4a)


tmp21229 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("Invalid eval argument: "), tmp21227), Nil)
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("Invalid eval argument: "), tmp21227)
__typedArg1 := Nil
return Call2(__e, SymFn(symcons), __typedArg0, __typedArg1)
})()

return GCons(__e, symerror, tmp21229)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4x_4launcher_4eval_1command_1h)
}


panic("unreachable")
})

let__21139 := tmp21225
_ = let__21139

tmp21244 := GConsP(__e, V7387)

if True == tmp21244 {
tmp21232 := GHd(__e, V7387)

tmp21233 := Call1(__e, SymFn(symshen_4x_4launcher_4eval_1flag_1map), tmp21232)


let__21141 := tmp21233
_ = let__21141

tmp21239 := GEq(__e, False, let__21141)

var ifres21234 Obj

if True == tmp21239 {
tmp21235 := Call0(__e, SymFn(symfail))


ifres21234 = tmp21235


} else {
tmp21236 := GTl(__e, V7387)

tmp21237 := GCons(__e, let__21141, tmp21236)

tmp21238 := Call2(__e, SymFn(symshen_4x_4launcher_4eval_1command_1h), tmp21237, V7388)


ifres21234 = tmp21238


}

let__21140 := ifres21234
_ = let__21140

tmp21241 := Call0(__e, SymFn(symfail))


tmp21242 := GEq(__e, let__21140, tmp21241)

if True == tmp21242 {
return TailCall1(__e, SymFn(symthaw), let__21139)
} else {
return let__21140
}


} else {
return TailCall1(__e, SymFn(symthaw), let__21139)
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

tmp21285 := Call2(__e, ns2_1set, symshen_4x_4launcher_4eval_1command_1h, tmp21191)


_ = tmp21285

tmp21286 := MakeNative1(func(__e *ControlFlow, V7392 Obj) Obj {
return TailCall2(__e, SymFn(symshen_4x_4launcher_4eval_1command_1h), V7392, Nil)
panic("unreachable")
})

tmp21287 := Call2(__e, ns2_1set, symshen_4x_4launcher_4eval_1command, tmp21286)


_ = tmp21287

tmp21288 := MakeNative2(func(__e *ControlFlow, V7393, V7394 Obj) Obj {
tmp21289 := GCons(__e, V7393, V7394)

tmp21290 := PrimGuard2(__e, symset, PrimSet, sym_dargv_d, tmp21289)

_ = tmp21290

tmp21291 := Call1(__e, SymFn(symshen_4x_4launcher_4quiet_1load), V7393)


_ = tmp21291

return GCons(__e, symsuccess, Nil)


panic("unreachable")
})

tmp21292 := Call2(__e, ns2_1set, symshen_4x_4launcher_4script_1command, tmp21288)


_ = tmp21292

tmp21293 := MakeNative1(func(__e *ControlFlow, V7395 Obj) Obj {
tmp21367 := GConsP(__e, V7395)

var ifres21364 Obj

if True == tmp21367 {
tmp21365 := GTl(__e, V7395)

tmp21366 := GEq(__e, Nil, tmp21365)

ifres21364 = tmp21366


} else {
ifres21364 = False


}

if True == ifres21364 {
return GCons(__e, symlaunch_1repl, Nil)
} else {
tmp21362 := GConsP(__e, V7395)

var ifres21355 Obj

if True == tmp21362 {
tmp21360 := GTl(__e, V7395)

tmp21361 := GConsP(__e, tmp21360)

var ifres21356 Obj

if True == tmp21361 {
tmp21357 := GTl(__e, V7395)

tmp21358 := GHd(__e, tmp21357)

tmp21359 := GEq(__e, MakeString("--help"), tmp21358)

ifres21356 = tmp21359


} else {
ifres21356 = False


}

ifres21355 = ifres21356


} else {
ifres21355 = False


}

if True == ifres21355 {
tmp21294 := GHd(__e, V7395)

tmp21295 := Call1(__e, SymFn(symshen_4x_4launcher_4help_1text), tmp21294)


tmp21296 := GCons(__e, tmp21295, Nil)

return GCons(__e, symshow_1help, tmp21296)


} else {
tmp21353 := GConsP(__e, V7395)

var ifres21346 Obj

if True == tmp21353 {
tmp21351 := GTl(__e, V7395)

tmp21352 := GConsP(__e, tmp21351)

var ifres21347 Obj

if True == tmp21352 {
tmp21348 := GTl(__e, V7395)

tmp21349 := GHd(__e, tmp21348)

tmp21350 := GEq(__e, MakeString("--version"), tmp21349)

ifres21347 = tmp21350


} else {
ifres21347 = False


}

ifres21346 = ifres21347


} else {
ifres21346 = False


}

if True == ifres21346 {
tmp21297 := Call0(__e, SymFn(symshen_4x_4launcher_4version_1string))


tmp21298 := GCons(__e, tmp21297, Nil)

return GCons(__e, symsuccess, tmp21298)


} else {
tmp21344 := GConsP(__e, V7395)

var ifres21337 Obj

if True == tmp21344 {
tmp21342 := GTl(__e, V7395)

tmp21343 := GConsP(__e, tmp21342)

var ifres21338 Obj

if True == tmp21343 {
tmp21339 := GTl(__e, V7395)

tmp21340 := GHd(__e, tmp21339)

tmp21341 := GEq(__e, MakeString("repl"), tmp21340)

ifres21338 = tmp21341


} else {
ifres21338 = False


}

ifres21337 = ifres21338


} else {
ifres21337 = False


}

if True == ifres21337 {
tmp21299 := GTl(__e, V7395)

tmp21300 := GTl(__e, tmp21299)

return GCons(__e, symlaunch_1repl, tmp21300)


} else {
tmp21335 := GConsP(__e, V7395)

var ifres21324 Obj

if True == tmp21335 {
tmp21333 := GTl(__e, V7395)

tmp21334 := GConsP(__e, tmp21333)

var ifres21325 Obj

if True == tmp21334 {
tmp21330 := GTl(__e, V7395)

tmp21331 := GHd(__e, tmp21330)

tmp21332 := GEq(__e, MakeString("script"), tmp21331)

var ifres21326 Obj

if True == tmp21332 {
tmp21327 := GTl(__e, V7395)

tmp21328 := GTl(__e, tmp21327)

tmp21329 := GConsP(__e, tmp21328)

ifres21326 = tmp21329


} else {
ifres21326 = False


}

ifres21325 = ifres21326


} else {
ifres21325 = False


}

ifres21324 = ifres21325


} else {
ifres21324 = False


}

if True == ifres21324 {
tmp21301 := GTl(__e, V7395)

tmp21302 := GTl(__e, tmp21301)

tmp21303 := GHd(__e, tmp21302)

tmp21304 := GTl(__e, V7395)

tmp21305 := GTl(__e, tmp21304)

tmp21306 := GTl(__e, tmp21305)

return TailCall2(__e, SymFn(symshen_4x_4launcher_4script_1command), tmp21303, tmp21306)


} else {
tmp21322 := GConsP(__e, V7395)

var ifres21315 Obj

if True == tmp21322 {
tmp21320 := GTl(__e, V7395)

tmp21321 := GConsP(__e, tmp21320)

var ifres21316 Obj

if True == tmp21321 {
tmp21317 := GTl(__e, V7395)

tmp21318 := GHd(__e, tmp21317)

tmp21319 := GEq(__e, MakeString("eval"), tmp21318)

ifres21316 = tmp21319


} else {
ifres21316 = False


}

ifres21315 = ifres21316


} else {
ifres21315 = False


}

if True == ifres21315 {
tmp21307 := GTl(__e, V7395)

tmp21308 := GTl(__e, tmp21307)

return TailCall1(__e, SymFn(symshen_4x_4launcher_4eval_1command), tmp21308)


} else {
tmp21313 := GConsP(__e, V7395)

var ifres21310 Obj

if True == tmp21313 {
tmp21311 := GTl(__e, V7395)

tmp21312 := GConsP(__e, tmp21311)

ifres21310 = tmp21312


} else {
ifres21310 = False


}

if True == ifres21310 {
return GCons(__e, symunknown_1arguments, V7395)
} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4x_4launcher_4launch_1shen)
}


}


}


}


}


}


}


panic("unreachable")
})

tmp21368 := Call2(__e, ns2_1set, symshen_4x_4launcher_4launch_1shen, tmp21293)


_ = tmp21368

tmp21369 := MakeNative1(func(__e *ControlFlow, V7398 Obj) Obj {
tmp21453 := GConsP(__e, V7398)

var ifres21447 Obj

if True == tmp21453 {
tmp21451 := GHd(__e, V7398)

tmp21452 := GEq(__e, symsuccess, tmp21451)

var ifres21448 Obj

if True == tmp21452 {
tmp21449 := GTl(__e, V7398)

tmp21450 := GEq(__e, Nil, tmp21449)

ifres21448 = tmp21450


} else {
ifres21448 = False


}

ifres21447 = ifres21448


} else {
ifres21447 = False


}

if True == ifres21447 {
return symshen_4x_4launcher_4done
} else {
tmp21445 := GConsP(__e, V7398)

var ifres21435 Obj

if True == tmp21445 {
tmp21443 := GHd(__e, V7398)

tmp21444 := GEq(__e, symsuccess, tmp21443)

var ifres21436 Obj

if True == tmp21444 {
tmp21441 := GTl(__e, V7398)

tmp21442 := GConsP(__e, tmp21441)

var ifres21437 Obj

if True == tmp21442 {
tmp21438 := GTl(__e, V7398)

tmp21439 := GTl(__e, tmp21438)

tmp21440 := GEq(__e, Nil, tmp21439)

ifres21437 = tmp21440


} else {
ifres21437 = False


}

ifres21436 = ifres21437


} else {
ifres21436 = False


}

ifres21435 = ifres21436


} else {
ifres21435 = False


}

if True == ifres21435 {
tmp21370 := GTl(__e, V7398)

tmp21371 := GHd(__e, tmp21370)

tmp21372 := Call3(__e, SymFn(symshen_4app), tmp21371, MakeString("\n"), symshen_4a)


tmp21373 := Call0(__e, SymFn(symstoutput))


return TailCall2(__e, SymFn(sympr), tmp21372, tmp21373)


} else {
tmp21433 := GConsP(__e, V7398)

var ifres21423 Obj

if True == tmp21433 {
tmp21431 := GHd(__e, V7398)

tmp21432 := GEq(__e, symerror, tmp21431)

var ifres21424 Obj

if True == tmp21432 {
tmp21429 := GTl(__e, V7398)

tmp21430 := GConsP(__e, tmp21429)

var ifres21425 Obj

if True == tmp21430 {
tmp21426 := GTl(__e, V7398)

tmp21427 := GTl(__e, tmp21426)

tmp21428 := GEq(__e, Nil, tmp21427)

ifres21425 = tmp21428


} else {
ifres21425 = False


}

ifres21424 = ifres21425


} else {
ifres21424 = False


}

ifres21423 = ifres21424


} else {
ifres21423 = False


}

if True == ifres21423 {
tmp21374 := GTl(__e, V7398)

tmp21375 := GHd(__e, tmp21374)

tmp21376 := Call3(__e, SymFn(symshen_4app), tmp21375, MakeString("\n"), symshen_4a)


tmp21377 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("ERROR: "), tmp21376)

tmp21378 := Call0(__e, SymFn(symstoutput))


return TailCall2(__e, SymFn(sympr), tmp21377, tmp21378)


} else {
tmp21421 := GConsP(__e, V7398)

var ifres21418 Obj

if True == tmp21421 {
tmp21419 := GHd(__e, V7398)

tmp21420 := GEq(__e, symlaunch_1repl, tmp21419)

ifres21418 = tmp21420


} else {
ifres21418 = False


}

if True == ifres21418 {
return TailCall0(__e, SymFn(symshen_4repl))
} else {
tmp21416 := GConsP(__e, V7398)

var ifres21406 Obj

if True == tmp21416 {
tmp21414 := GHd(__e, V7398)

tmp21415 := GEq(__e, symshow_1help, tmp21414)

var ifres21407 Obj

if True == tmp21415 {
tmp21412 := GTl(__e, V7398)

tmp21413 := GConsP(__e, tmp21412)

var ifres21408 Obj

if True == tmp21413 {
tmp21409 := GTl(__e, V7398)

tmp21410 := GTl(__e, tmp21409)

tmp21411 := GEq(__e, Nil, tmp21410)

ifres21408 = tmp21411


} else {
ifres21408 = False


}

ifres21407 = ifres21408


} else {
ifres21407 = False


}

ifres21406 = ifres21407


} else {
ifres21406 = False


}

if True == ifres21406 {
tmp21379 := GTl(__e, V7398)

tmp21380 := GHd(__e, tmp21379)

tmp21381 := Call3(__e, SymFn(symshen_4app), tmp21380, MakeString("\n"), symshen_4a)


tmp21382 := Call0(__e, SymFn(symstoutput))


return TailCall2(__e, SymFn(sympr), tmp21381, tmp21382)


} else {
tmp21404 := GConsP(__e, V7398)

var ifres21394 Obj

if True == tmp21404 {
tmp21402 := GHd(__e, V7398)

tmp21403 := GEq(__e, symunknown_1arguments, tmp21402)

var ifres21395 Obj

if True == tmp21403 {
tmp21400 := GTl(__e, V7398)

tmp21401 := GConsP(__e, tmp21400)

var ifres21396 Obj

if True == tmp21401 {
tmp21397 := GTl(__e, V7398)

tmp21398 := GTl(__e, tmp21397)

tmp21399 := GConsP(__e, tmp21398)

ifres21396 = tmp21399


} else {
ifres21396 = False


}

ifres21395 = ifres21396


} else {
ifres21395 = False


}

ifres21394 = ifres21395


} else {
ifres21394 = False


}

if True == ifres21394 {
tmp21383 := GTl(__e, V7398)

tmp21384 := GTl(__e, tmp21383)

tmp21385 := GHd(__e, tmp21384)

tmp21386 := GTl(__e, V7398)

tmp21387 := GHd(__e, tmp21386)

tmp21388 := Call3(__e, SymFn(symshen_4app), tmp21387, MakeString(" --help' for more information.\n"), symshen_4a)


tmp21390 := Call3(__e, SymFn(symshen_4app), tmp21385, PrimGuard2(__e, symcn, PrimStringConcat, MakeString("\nTry `"), tmp21388), symshen_4a)


tmp21391 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("ERROR: Invalid argument: "), tmp21390)

tmp21392 := Call0(__e, SymFn(symstoutput))


return TailCall2(__e, SymFn(sympr), tmp21391, tmp21392)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4x_4launcher_4default_1handle_1result)
}


}


}


}


}


}


panic("unreachable")
})

tmp21454 := Call2(__e, ns2_1set, symshen_4x_4launcher_4default_1handle_1result, tmp21369)


_ = tmp21454

tmp21455 := MakeNative1(func(__e *ControlFlow, V7399 Obj) Obj {
tmp21456 := Call1(__e, SymFn(symshen_4x_4launcher_4launch_1shen), V7399)


return TailCall1(__e, SymFn(symshen_4x_4launcher_4default_1handle_1result), tmp21456)


panic("unreachable")
})

return TailCall2(__e, ns2_1set, symshen_4x_4launcher_4main, tmp21455)




panic("unreachable")
})

var sym_b = MakeSymbol("!")
var sym_3 = MakeSymbol("$")
var sym_e = MakeSymbol("&")
var sym_e_e = MakeSymbol("&&")
var sym_d = MakeSymbol("*")
var sym_dabsolute_d = MakeSymbol("*absolute*")
var sym_dargv_d = MakeSymbol("*argv*")
var sym_dhome_1directory_d = MakeSymbol("*home-directory*")
var sym_dhush_d = MakeSymbol("*hush*")
var sym_dimplementation_d = MakeSymbol("*implementation*")
var sym_dlanguage_d = MakeSymbol("*language*")
var sym_dmacros_d = MakeSymbol("*macros*")
var sym_dmaximum_1print_1sequence_1size_d = MakeSymbol("*maximum-print-sequence-size*")
var sym_dos_d = MakeSymbol("*os*")
var sym_dport_d = MakeSymbol("*port*")
var sym_dporters_d = MakeSymbol("*porters*")
var sym_dproperty_1vector_d = MakeSymbol("*property-vector*")
var sym_drelease_d = MakeSymbol("*release*")
var sym_dsterror_d = MakeSymbol("*sterror*")
var sym_dstinput_d = MakeSymbol("*stinput*")
var sym_dstoutput_d = MakeSymbol("*stoutput*")
var sym_dversion_d = MakeSymbol("*version*")
var sym_7 = MakeSymbol("+")
var sym_1 = MakeSymbol("-")
var sym_1_1_6 = MakeSymbol("-->")
var sym_1_6 = MakeSymbol("->")
var sym_c = MakeSymbol("/")
var sym_c_4 = MakeSymbol("/.")
var sym_5 = MakeSymbol("<")
var sym_5_b_6 = MakeSymbol("<!>")
var sym_5_1 = MakeSymbol("<-")
var sym_5_1_1 = MakeSymbol("<--")
var sym_5_1address = MakeSymbol("<-address")
var sym_5_1vector = MakeSymbol("<-vector")
var sym_5_a = MakeSymbol("<=")
var sym_5e_6 = MakeSymbol("<e>")
var sym_5end_6 = MakeSymbol("<end>")
var sym_a = MakeSymbol("=")
var sym_a_a = MakeSymbol("==")
var sym_a_a_6 = MakeSymbol("==>")
var sym_6 = MakeSymbol(">")
var sym_6_a = MakeSymbol(">=")
var sym_6_6 = MakeSymbol(">>")
var sym_8p = MakeSymbol("@p")
var sym_8s = MakeSymbol("@s")
var sym_8v = MakeSymbol("@v")
var symA = MakeSymbol("A")
var symAction = MakeSymbol("Action")
var symAssumption = MakeSymbol("Assumption")
var symAssumptions = MakeSymbol("Assumptions")
var symB = MakeSymbol("B")
var symC = MakeSymbol("C")
var symFinish = MakeSymbol("Finish")
var symFreeze = MakeSymbol("Freeze")
var symGoTo = MakeSymbol("GoTo")
var symHd = MakeSymbol("Hd")
var symHypotheses = MakeSymbol("Hypotheses")
var symK = MakeSymbol("K")
var symKey = MakeSymbol("Key")
var symL = MakeSymbol("L")
var symMessage = MakeSymbol("Message")
var symNewAssumptions = MakeSymbol("NewAssumptions")
var symNewV = MakeSymbol("NewV")
var symP = MakeSymbol("P")
var symParse = MakeSymbol("Parse")
var symRecord = MakeSymbol("Record")
var symRemainder = MakeSymbol("Remainder")
var symResult = MakeSymbol("Result")
var symS = MakeSymbol("S")
var symSelect = MakeSymbol("Select")
var symStart = MakeSymbol("Start")
var symTime = MakeSymbol("Time")
var symTl = MakeSymbol("Tl")
var symTm = MakeSymbol("Tm")
var symV = MakeSymbol("V")
var symW = MakeSymbol("W")
var symX = MakeSymbol("X")
var symY = MakeSymbol("Y")
var symZ = MakeSymbol("Z")
var sym__ = MakeSymbol("_")
var symabort = MakeSymbol("abort")
var symabsolute = MakeSymbol("absolute")
var symabsvector = MakeSymbol("absvector")
var symabsvector_2 = MakeSymbol("absvector?")
var symaddress_1_6 = MakeSymbol("address->")
var symadjoin = MakeSymbol("adjoin")
var symand = MakeSymbol("and")
var symappend = MakeSymbol("append")
var symarity = MakeSymbol("arity")
var symasserta = MakeSymbol("asserta")
var symassertz = MakeSymbol("assertz")
var symassoc = MakeSymbol("assoc")
var symatom_2 = MakeSymbol("atom?")
var symbar_b = MakeSymbol("bar!")
var symbind = MakeSymbol("bind")
var symboolean = MakeSymbol("boolean")
var symboolean_2 = MakeSymbol("boolean?")
var symbootstrap = MakeSymbol("bootstrap")
var symbound_2 = MakeSymbol("bound?")
var symcall = MakeSymbol("call")
var symcases = MakeSymbol("cases")
var symcd = MakeSymbol("cd")
var symchar_1stream_2 = MakeSymbol("char-stream?")
var symclose = MakeSymbol("close")
var symcn = MakeSymbol("cn")
var symcompile = MakeSymbol("compile")
var symconcat = MakeSymbol("concat")
var symcond = MakeSymbol("cond")
var symcons = MakeSymbol("cons")
var symcons_2 = MakeSymbol("cons?")
var symctxt = MakeSymbol("ctxt")
var symdatatype = MakeSymbol("datatype")
var symdatatypes = MakeSymbol("datatypes")
var symdeclare = MakeSymbol("declare")
var symdefcc = MakeSymbol("defcc")
var symdefine = MakeSymbol("define")
var symdefmacro = MakeSymbol("defmacro")
var symdefprolog = MakeSymbol("defprolog")
var symdefun = MakeSymbol("defun")
var symdestroy = MakeSymbol("destroy")
var symdifference = MakeSymbol("difference")
var symdo = MakeSymbol("do")
var symelement_2 = MakeSymbol("element?")
var symempty_2 = MakeSymbol("empty?")
var symenable_1type_1theory = MakeSymbol("enable-type-theory")
var symerror = MakeSymbol("error")
var symerror_1to_1string = MakeSymbol("error-to-string")
var symeval = MakeSymbol("eval")
var symeval_1kl = MakeSymbol("eval-kl")
var symexception = MakeSymbol("exception")
var symexplode = MakeSymbol("explode")
var symexternal = MakeSymbol("external")
var symfactorise = MakeSymbol("factorise")
var symfactorise_2 = MakeSymbol("factorise?")
var symfail = MakeSymbol("fail")
var symfail_1if = MakeSymbol("fail-if")
var symfile = MakeSymbol("file")
var symfindall = MakeSymbol("findall")
var symfix = MakeSymbol("fix")
var symfn = MakeSymbol("fn")
var symforeign = MakeSymbol("foreign")
var symfork = MakeSymbol("fork")
var symfreeze = MakeSymbol("freeze")
var symfresh = MakeSymbol("fresh")
var symfst = MakeSymbol("fst")
var symfunction = MakeSymbol("function")
var symg = MakeSymbol("g")
var symgensym = MakeSymbol("gensym")
var symget = MakeSymbol("get")
var symget_1time = MakeSymbol("get-time")
var symhash = MakeSymbol("hash")
var symhd = MakeSymbol("hd")
var symhdstr = MakeSymbol("hdstr")
var symhdv = MakeSymbol("hdv")
var symhead = MakeSymbol("head")
var symhush = MakeSymbol("hush")
var symhush_2 = MakeSymbol("hush?")
var symif = MakeSymbol("if")
var symimplementation = MakeSymbol("implementation")
var symin = MakeSymbol("in")
var symin_1package = MakeSymbol("in-package")
var syminclude = MakeSymbol("include")
var syminclude_1all_1but = MakeSymbol("include-all-but")
var symincluded = MakeSymbol("included")
var syminferences = MakeSymbol("inferences")
var syminline = MakeSymbol("inline")
var syminput = MakeSymbol("input")
var syminput_7 = MakeSymbol("input+")
var syminteger_2 = MakeSymbol("integer?")
var symintern = MakeSymbol("intern")
var syminternal = MakeSymbol("internal")
var symintersection = MakeSymbol("intersection")
var symis = MakeSymbol("is")
var symis_b = MakeSymbol("is!")
var symit = MakeSymbol("it")
var symlambda = MakeSymbol("lambda")
var symlanguage = MakeSymbol("language")
var symlaunch_1repl = MakeSymbol("launch-repl")
var symlazy = MakeSymbol("lazy")
var symlength = MakeSymbol("length")
var symlet = MakeSymbol("let")
var symlimit = MakeSymbol("limit")
var symlineread = MakeSymbol("lineread")
var symlist = MakeSymbol("list")
var symload = MakeSymbol("load")
var symloaded = MakeSymbol("loaded")
var symmacroexpand = MakeSymbol("macroexpand")
var symmake_1string = MakeSymbol("make-string")
var symmap = MakeSymbol("map")
var symmapcan = MakeSymbol("mapcan")
var symmaxinferences = MakeSymbol("maxinferences")
var symmode = MakeSymbol("mode")
var symn_1_6string = MakeSymbol("n->string")
var symnl = MakeSymbol("nl")
var symnot = MakeSymbol("not")
var symnth = MakeSymbol("nth")
var symnull = MakeSymbol("null")
var symnumber = MakeSymbol("number")
var symnumber_2 = MakeSymbol("number?")
var symoccurrences = MakeSymbol("occurrences")
var symoccurs_1check = MakeSymbol("occurs-check")
var symoccurs_2 = MakeSymbol("occurs?")
var symopen = MakeSymbol("open")
var symoptimise = MakeSymbol("optimise")
var symoptimise_2 = MakeSymbol("optimise?")
var symor = MakeSymbol("or")
var symos = MakeSymbol("os")
var symout = MakeSymbol("out")
var symoutput = MakeSymbol("output")
var sympackage = MakeSymbol("package")
var sympackage_2 = MakeSymbol("package?")
var symport = MakeSymbol("port")
var symporters = MakeSymbol("porters")
var sympos = MakeSymbol("pos")
var sympr = MakeSymbol("pr")
var sympreclude = MakeSymbol("preclude")
var sympreclude_1all_1but = MakeSymbol("preclude-all-but")
var symprint = MakeSymbol("print")
var symprofile = MakeSymbol("profile")
var symprofile_1results = MakeSymbol("profile-results")
var symprolog_1memory = MakeSymbol("prolog-memory")
var symprolog_2 = MakeSymbol("prolog?")
var symprotect = MakeSymbol("protect")
var symps = MakeSymbol("ps")
var symput = MakeSymbol("put")
var symread = MakeSymbol("read")
var symread_1byte = MakeSymbol("read-byte")
var symread_1file = MakeSymbol("read-file")
var symread_1file_1as_1bytelist = MakeSymbol("read-file-as-bytelist")
var symread_1file_1as_1string = MakeSymbol("read-file-as-string")
var symread_1from_1string = MakeSymbol("read-from-string")
var symread_1from_1string_1unprocessed = MakeSymbol("read-from-string-unprocessed")
var symread_1unit_1string = MakeSymbol("read-unit-string")
var symreceive = MakeSymbol("receive")
var symrelease = MakeSymbol("release")
var symremove = MakeSymbol("remove")
var symretract = MakeSymbol("retract")
var symreturn = MakeSymbol("return")
var symreverse = MakeSymbol("reverse")
var symrun = MakeSymbol("run")
var symsave = MakeSymbol("save")
var symset = MakeSymbol("set")
var symshen = MakeSymbol("shen")
var symshen_4_dalldatatypes_d = MakeSymbol("shen.*alldatatypes*")
var symshen_4_dcall_d = MakeSymbol("shen.*call*")
var symshen_4_dcustom_1pattern_1compiler_d = MakeSymbol("shen.*custom-pattern-compiler*")
var symshen_4_dcustom_1pattern_1reducer_d = MakeSymbol("shen.*custom-pattern-reducer*")
var symshen_4_ddatatypes_d = MakeSymbol("shen.*datatypes*")
var symshen_4_ddemodulation_1function_d = MakeSymbol("shen.*demodulation-function*")
var symshen_4_dempty_1absvector_d = MakeSymbol("shen.*empty-absvector*")
var symshen_4_dextraspecial_d = MakeSymbol("shen.*extraspecial*")
var symshen_4_dfactorise_2_d = MakeSymbol("shen.*factorise?*")
var symshen_4_dgensym_d = MakeSymbol("shen.*gensym*")
var symshen_4_dhistory_d = MakeSymbol("shen.*history*")
var symshen_4_dinfs_d = MakeSymbol("shen.*infs*")
var symshen_4_dit_d = MakeSymbol("shen.*it*")
var symshen_4_dloading_2_d = MakeSymbol("shen.*loading?*")
var symshen_4_dmaxinferences_d = MakeSymbol("shen.*maxinferences*")
var symshen_4_dnames_d = MakeSymbol("shen.*names*")
var symshen_4_doccurs_d = MakeSymbol("shen.*occurs*")
var symshen_4_doptimise_d = MakeSymbol("shen.*optimise*")
var symshen_4_dpackage_d = MakeSymbol("shen.*package*")
var symshen_4_dprofiled_d = MakeSymbol("shen.*profiled*")
var symshen_4_dprolog_1memory_d = MakeSymbol("shen.*prolog-memory*")
var symshen_4_dresidue_d = MakeSymbol("shen.*residue*")
var symshen_4_dshen_1type_1theory_1enabled_2_d = MakeSymbol("shen.*shen-type-theory-enabled?*")
var symshen_4_dsigf_d = MakeSymbol("shen.*sigf*")
var symshen_4_dsize_1prolog_1vector_d = MakeSymbol("shen.*size-prolog-vector*")
var symshen_4_dsource_1form_1handlers_d = MakeSymbol("shen.*source-form-handlers*")
var symshen_4_dspecial_d = MakeSymbol("shen.*special*")
var symshen_4_dspy_d = MakeSymbol("shen.*spy*")
var symshen_4_dstep_d = MakeSymbol("shen.*step*")
var symshen_4_dsynonyms_d = MakeSymbol("shen.*synonyms*")
var symshen_4_dtc_d = MakeSymbol("shen.*tc*")
var symshen_4_dtracking_d = MakeSymbol("shen.*tracking*")
var symshen_4_duserdefs_d = MakeSymbol("shen.*userdefs*")
var symshen_4_7m = MakeSymbol("shen.+m")
var symshen_4_7string_2 = MakeSymbol("shen.+string?")
var symshen_4_7vector_2 = MakeSymbol("shen.+vector?")
var symshen_4_1m = MakeSymbol("shen.-m")
var symshen_4_1null_1 = MakeSymbol("shen.-null-")
var symshen_4_5_1dict = MakeSymbol("shen.<-dict")
var symshen_4_5_1dict_1bucket = MakeSymbol("shen.<-dict-bucket")
var symshen_4_5_1out = MakeSymbol("shen.<-out")
var symshen_4_5alpha_6 = MakeSymbol("shen.<alpha>")
var symshen_4_5alphanum_6 = MakeSymbol("shen.<alphanum>")
var symshen_4_5alphanums_6 = MakeSymbol("shen.<alphanums>")
var symshen_4_5ass_6 = MakeSymbol("shen.<ass>")
var symshen_4_5atom_6 = MakeSymbol("shen.<atom>")
var symshen_4_5backslash_6 = MakeSymbol("shen.<backslash>")
var symshen_4_5bar_6 = MakeSymbol("shen.<bar>")
var symshen_4_5body_6 = MakeSymbol("shen.<body>")
var symshen_4_5bterm_6 = MakeSymbol("shen.<bterm>")
var symshen_4_5bterms_6 = MakeSymbol("shen.<bterms>")
var symshen_4_5c_1rule_6 = MakeSymbol("shen.<c-rule>")
var symshen_4_5c_1rules_6 = MakeSymbol("shen.<c-rules>")
var symshen_4_5clause_6 = MakeSymbol("shen.<clause>")
var symshen_4_5clauses_6 = MakeSymbol("shen.<clauses>")
var symshen_4_5colon_1equal_6 = MakeSymbol("shen.<colon-equal>")
var symshen_4_5colon_6 = MakeSymbol("shen.<colon>")
var symshen_4_5comma_6 = MakeSymbol("shen.<comma>")
var symshen_4_5comment_6 = MakeSymbol("shen.<comment>")
var symshen_4_5conc_6 = MakeSymbol("shen.<conc>")
var symshen_4_5constructor_6 = MakeSymbol("shen.<constructor>")
var symshen_4_5control_6 = MakeSymbol("shen.<control>")
var symshen_4_5datatype_1rule_6 = MakeSymbol("shen.<datatype-rule>")
var symshen_4_5datatype_1rules_6 = MakeSymbol("shen.<datatype-rules>")
var symshen_4_5datatype_6 = MakeSymbol("shen.<datatype>")
var symshen_4_5dbl_6 = MakeSymbol("shen.<dbl>")
var symshen_4_5dbq_6 = MakeSymbol("shen.<dbq>")
var symshen_4_5define_6 = MakeSymbol("shen.<define>")
var symshen_4_5defprolog_6 = MakeSymbol("shen.<defprolog>")
var symshen_4_5digit_6 = MakeSymbol("shen.<digit>")
var symshen_4_5digits_6 = MakeSymbol("shen.<digits>")
var symshen_4_5double_6 = MakeSymbol("shen.<double>")
var symshen_4_5e_1number_6 = MakeSymbol("shen.<e-number>")
var symshen_4_5equal_6 = MakeSymbol("shen.<equal>")
var symshen_4_5expr_6 = MakeSymbol("shen.<expr>")
var symshen_4_5float_6 = MakeSymbol("shen.<float>")
var symshen_4_5formula_6 = MakeSymbol("shen.<formula>")
var symshen_4_5formulae_6 = MakeSymbol("shen.<formulae>")
var symshen_4_5fraction_6 = MakeSymbol("shen.<fraction>")
var symshen_4_5hash_6 = MakeSymbol("shen.<hash>")
var symshen_4_5head_6 = MakeSymbol("shen.<head>")
var symshen_4_5hterm1_6 = MakeSymbol("shen.<hterm1>")
var symshen_4_5hterm2_6 = MakeSymbol("shen.<hterm2>")
var symshen_4_5hterm_6 = MakeSymbol("shen.<hterm>")
var symshen_4_5integer_6 = MakeSymbol("shen.<integer>")
var symshen_4_5iscolon_6 = MakeSymbol("shen.<iscolon>")
var symshen_4_5iscomma_6 = MakeSymbol("shen.<iscomma>")
var symshen_4_5lcurly_6 = MakeSymbol("shen.<lcurly>")
var symshen_4_5literal_6 = MakeSymbol("shen.<literal>")
var symshen_4_5log10_6 = MakeSymbol("shen.<log10>")
var symshen_4_5longnatter_6 = MakeSymbol("shen.<longnatter>")
var symshen_4_5lowC_6 = MakeSymbol("shen.<lowC>")
var symshen_4_5lowE_6 = MakeSymbol("shen.<lowE>")
var symshen_4_5lrb_6 = MakeSymbol("shen.<lrb>")
var symshen_4_5lsb_6 = MakeSymbol("shen.<lsb>")
var symshen_4_5minus_6 = MakeSymbol("shen.<minus>")
var symshen_4_5multiline_6 = MakeSymbol("shen.<multiline>")
var symshen_4_5name_6 = MakeSymbol("shen.<name>")
var symshen_4_5non_1terminal_1name_6 = MakeSymbol("shen.<non-terminal-name>")
var symshen_4_5non_1terminal_2_6 = MakeSymbol("shen.<non-terminal?>")
var symshen_4_5notdbq_6 = MakeSymbol("shen.<notdbq>")
var symshen_4_5number_6 = MakeSymbol("shen.<number>")
var symshen_4_5numeral_6 = MakeSymbol("shen.<numeral>")
var symshen_4_5packagechar_6 = MakeSymbol("shen.<packagechar>")
var symshen_4_5packagename_6 = MakeSymbol("shen.<packagename>")
var symshen_4_5packagenames_6 = MakeSymbol("shen.<packagenames>")
var symshen_4_5pattern1_6 = MakeSymbol("shen.<pattern1>")
var symshen_4_5pattern2_6 = MakeSymbol("shen.<pattern2>")
var symshen_4_5pattern_6 = MakeSymbol("shen.<pattern>")
var symshen_4_5patterns_6 = MakeSymbol("shen.<patterns>")
var symshen_4_5plus_6 = MakeSymbol("shen.<plus>")
var symshen_4_5prem_6 = MakeSymbol("shen.<prem>")
var symshen_4_5prems_6 = MakeSymbol("shen.<prems>")
var symshen_4_5rcurly_6 = MakeSymbol("shen.<rcurly>")
var symshen_4_5return_6 = MakeSymbol("shen.<return>")
var symshen_4_5returns_6 = MakeSymbol("shen.<returns>")
var symshen_4_5rrb_6 = MakeSymbol("shen.<rrb>")
var symshen_4_5rsb_6 = MakeSymbol("shen.<rsb>")
var symshen_4_5rule_d_6 = MakeSymbol("shen.<rule*>")
var symshen_4_5rule_6 = MakeSymbol("shen.<rule>")
var symshen_4_5rules_d_6 = MakeSymbol("shen.<rules*>")
var symshen_4_5rules_6 = MakeSymbol("shen.<rules>")
var symshen_4_5s_1exprs1_6 = MakeSymbol("shen.<s-exprs1>")
var symshen_4_5s_1exprs2_6 = MakeSymbol("shen.<s-exprs2>")
var symshen_4_5s_1exprs_6 = MakeSymbol("shen.<s-exprs>")
var symshen_4_5sc_6 = MakeSymbol("shen.<sc>")
var symshen_4_5semantics_6 = MakeSymbol("shen.<semantics>")
var symshen_4_5semicolon_6 = MakeSymbol("shen.<semicolon>")
var symshen_4_5shortnatter_6 = MakeSymbol("shen.<shortnatter>")
var symshen_4_5shortnatters_6 = MakeSymbol("shen.<shortnatters>")
var symshen_4_5side_6 = MakeSymbol("shen.<side>")
var symshen_4_5sides_6 = MakeSymbol("shen.<sides>")
var symshen_4_5sig_drules_6 = MakeSymbol("shen.<sig*rules>")
var symshen_4_5signature_6 = MakeSymbol("shen.<signature>")
var symshen_4_5simple_1pattern_6 = MakeSymbol("shen.<simple-pattern>")
var symshen_4_5single_6 = MakeSymbol("shen.<single>")
var symshen_4_5singleline_6 = MakeSymbol("shen.<singleline>")
var symshen_4_5sng_6 = MakeSymbol("shen.<sng>")
var symshen_4_5stop_6 = MakeSymbol("shen.<stop>")
var symshen_4_5str_6 = MakeSymbol("shen.<str>")
var symshen_4_5strc_6 = MakeSymbol("shen.<strc>")
var symshen_4_5strcontents_6 = MakeSymbol("shen.<strcontents>")
var symshen_4_5sym_6 = MakeSymbol("shen.<sym>")
var symshen_4_5syntax_1item_6 = MakeSymbol("shen.<syntax-item>")
var symshen_4_5syntax_6 = MakeSymbol("shen.<syntax>")
var symshen_4_5times_6 = MakeSymbol("shen.<times>")
var symshen_4_5type_6 = MakeSymbol("shen.<type>")
var symshen_4_5whitespace_6 = MakeSymbol("shen.<whitespace>")
var symshen_4_5whitespaces_6 = MakeSymbol("shen.<whitespaces>")
var symshen_4_5wildcard_6 = MakeSymbol("shen.<wildcard>")
var symshen_4_5yacc_6 = MakeSymbol("shen.<yacc>")
var symshen_4_5yaccsig_6 = MakeSymbol("shen.<yaccsig>")
var symshen_4_8c = MakeSymbol("shen.@c")
var symshen_4_8ch = MakeSymbol("shen.@ch")
var symshen_4_8v_1help = MakeSymbol("shen.@v-help")
var symshen_4a = MakeSymbol("shen.a")
var symshen_4abs = MakeSymbol("shen.abs")
var symshen_4add_1sexpr = MakeSymbol("shen.add-sexpr")
var symshen_4alpha_1convert = MakeSymbol("shen.alpha-convert")
var symshen_4alpha_2 = MakeSymbol("shen.alpha?")
var symshen_4alphanums_2 = MakeSymbol("shen.alphanums?")
var symshen_4analyse_1symbol_2 = MakeSymbol("shen.analyse-symbol?")
var symshen_4analyse_1variable_2 = MakeSymbol("shen.analyse-variable?")
var symshen_4app = MakeSymbol("shen.app")
var symshen_4application_2 = MakeSymbol("shen.application?")
var symshen_4arg_1_6str = MakeSymbol("shen.arg->str")
var symshen_4arity_1chk = MakeSymbol("shen.arity-chk")
var symshen_4assert_d = MakeSymbol("shen.assert*")
var symshen_4assoc_1_6 = MakeSymbol("shen.assoc->")
var symshen_4assoc_1rm = MakeSymbol("shen.assoc-rm")
var symshen_4assoc_1set = MakeSymbol("shen.assoc-set")
var symshen_4assumetypes = MakeSymbol("shen.assumetypes")
var symshen_4atom_1_6str = MakeSymbol("shen.atom->str")
var symshen_4atom_1case_1minus = MakeSymbol("shen.atom-case-minus")
var symshen_4atom_1case_1plus = MakeSymbol("shen.atom-case-plus")
var symshen_4autocomplete = MakeSymbol("shen.autocomplete")
var symshen_4bad_1pivot_2 = MakeSymbol("shen.bad-pivot?")
var symshen_4beta = MakeSymbol("shen.beta")
var symshen_4bind_b = MakeSymbol("shen.bind!")
var symshen_4bindv = MakeSymbol("shen.bindv")
var symshen_4bottom = MakeSymbol("shen.bottom")
var symshen_4bucket_1fold = MakeSymbol("shen.bucket-fold")
var symshen_4build_1lambda_1table = MakeSymbol("shen.build-lambda-table")
var symshen_4by_1hypothesis = MakeSymbol("shen.by-hypothesis")
var symshen_4by_1hypothesis_1h = MakeSymbol("shen.by-hypothesis-h")
var symshen_4byte_1_6digit = MakeSymbol("shen.byte->digit")
var symshen_4bytes_1_6string = MakeSymbol("shen.bytes->string")
var symshen_4c_1rule_1_6shen = MakeSymbol("shen.c-rule->shen")
var symshen_4c_1rules_1_6shen = MakeSymbol("shen.c-rules->shen")
var symshen_4call_1dynamic = MakeSymbol("shen.call-dynamic")
var symshen_4call_1prolog = MakeSymbol("shen.call-prolog")
var symshen_4callrec = MakeSymbol("shen.callrec")
var symshen_4ccons_2 = MakeSymbol("shen.ccons?")
var symshen_4check_1byte = MakeSymbol("shen.check-byte")
var symshen_4check_1eval_1and_1print = MakeSymbol("shen.check-eval-and-print")
var symshen_4choicepoint = MakeSymbol("shen.choicepoint")
var symshen_4choicepoint_b = MakeSymbol("shen.choicepoint!")
var symshen_4choicepoint_2 = MakeSymbol("shen.choicepoint?")
var symshen_4coll_1formulae = MakeSymbol("shen.coll-formulae")
var symshen_4colon_1equal_2 = MakeSymbol("shen.colon-equal?")
var symshen_4comb = MakeSymbol("shen.comb")
var symshen_4combine_1c_1code = MakeSymbol("shen.combine-c-code")
var symshen_4compile_1body = MakeSymbol("shen.compile-body")
var symshen_4compile_1head = MakeSymbol("shen.compile-head")
var symshen_4compile_1pattern_1fragment = MakeSymbol("shen.compile-pattern-fragment")
var symshen_4compile_1prolog = MakeSymbol("shen.compile-prolog")
var symshen_4compile_1synonyms = MakeSymbol("shen.compile-synonyms")
var symshen_4compile_1to_1kl = MakeSymbol("shen.compile-to-kl")
var symshen_4compound_1pattern = MakeSymbol("shen.compound-pattern")
var symshen_4compound_1pattern_1h = MakeSymbol("shen.compound-pattern-h")
var symshen_4compute_1E = MakeSymbol("shen.compute-E")
var symshen_4compute_1fraction = MakeSymbol("shen.compute-fraction")
var symshen_4compute_1fraction_1h = MakeSymbol("shen.compute-fraction-h")
var symshen_4compute_1integer = MakeSymbol("shen.compute-integer")
var symshen_4compute_1integer_1h = MakeSymbol("shen.compute-integer-h")
var symshen_4cond_1form = MakeSymbol("shen.cond-form")
var symshen_4cons_1case_1minus = MakeSymbol("shen.cons-case-minus")
var symshen_4cons_1case_1plus = MakeSymbol("shen.cons-case-plus")
var symshen_4cons_1form = MakeSymbol("shen.cons-form")
var symshen_4cons_1form_1respect_1modes = MakeSymbol("shen.cons-form-respect-modes")
var symshen_4conscode = MakeSymbol("shen.conscode")
var symshen_4construct_1context = MakeSymbol("shen.construct-context")
var symshen_4constructor_1error = MakeSymbol("shen.constructor-error")
var symshen_4constructor_2 = MakeSymbol("shen.constructor?")
var symshen_4consume = MakeSymbol("shen.consume")
var symshen_4consume_1clause = MakeSymbol("shen.consume-clause")
var symshen_4continue = MakeSymbol("shen.continue")
var symshen_4copyfromvector = MakeSymbol("shen.copyfromvector")
var symshen_4correct = MakeSymbol("shen.correct")
var symshen_4create_1skeleton = MakeSymbol("shen.create-skeleton")
var symshen_4credits = MakeSymbol("shen.credits")
var symshen_4curry = MakeSymbol("shen.curry")
var symshen_4curry_1type = MakeSymbol("shen.curry-type")
var symshen_4custom_1pattern = MakeSymbol("shen.custom-pattern")
var symshen_4custom_1pattern_1body = MakeSymbol("shen.custom-pattern-body")
var symshen_4custom_1pattern_1compiler = MakeSymbol("shen.custom-pattern-compiler")
var symshen_4custom_1pattern_1reducer = MakeSymbol("shen.custom-pattern-reducer")
var symshen_4custom_1pattern_1triple_1stack = MakeSymbol("shen.custom-pattern-triple-stack")
var symshen_4custom_1pattern_2 = MakeSymbol("shen.custom-pattern?")
var symshen_4cut = MakeSymbol("shen.cut")
var symshen_4dbl_1h_2 = MakeSymbol("shen.dbl-h?")
var symshen_4dbl_2 = MakeSymbol("shen.dbl?")
var symshen_4declare_1with_1signature = MakeSymbol("shen.declare-with-signature")
var symshen_4decons = MakeSymbol("shen.decons")
var symshen_4decrement_1ticket = MakeSymbol("shen.decrement-ticket")
var symshen_4demod = MakeSymbol("shen.demod")
var symshen_4demode = MakeSymbol("shen.demode")
var symshen_4demodulate = MakeSymbol("shen.demodulate")
var symshen_4deref = MakeSymbol("shen.deref")
var symshen_4deref_1calls = MakeSymbol("shen.deref-calls")
var symshen_4deref_1forked_1literals = MakeSymbol("shen.deref-forked-literals")
var symshen_4deref_1terms = MakeSymbol("shen.deref-terms")
var symshen_4dict = MakeSymbol("shen.dict")
var symshen_4dict_1_6 = MakeSymbol("shen.dict->")
var symshen_4dict_1bucket_1_6 = MakeSymbol("shen.dict-bucket->")
var symshen_4dict_1capacity = MakeSymbol("shen.dict-capacity")
var symshen_4dict_1count = MakeSymbol("shen.dict-count")
var symshen_4dict_1count_1_6 = MakeSymbol("shen.dict-count->")
var symshen_4dict_1fold = MakeSymbol("shen.dict-fold")
var symshen_4dict_1fold_1h = MakeSymbol("shen.dict-fold-h")
var symshen_4dict_1keys = MakeSymbol("shen.dict-keys")
var symshen_4dict_1rm = MakeSymbol("shen.dict-rm")
var symshen_4dict_1update_1count = MakeSymbol("shen.dict-update-count")
var symshen_4dict_1values = MakeSymbol("shen.dict-values")
var symshen_4dict_2 = MakeSymbol("shen.dict?")
var symshen_4dictionary = MakeSymbol("shen.dictionary")
var symshen_4digit_2 = MakeSymbol("shen.digit?")
var symshen_4dynamic = MakeSymbol("shen.dynamic")
var symshen_4dynamic_1default = MakeSymbol("shen.dynamic-default")
var symshen_4empty_1absvector_2 = MakeSymbol("shen.empty-absvector?")
var symshen_4eos = MakeSymbol("shen.eos")
var symshen_4eval_1and_1print = MakeSymbol("shen.eval-and-print")
var symshen_4evaluate_1lineread = MakeSymbol("shen.evaluate-lineread")
var symshen_4execute_1store_1arity = MakeSymbol("shen.execute-store-arity")
var symshen_4explode_1h = MakeSymbol("shen.explode-h")
var symshen_4expt = MakeSymbol("shen.expt")
var symshen_4external_1symbols = MakeSymbol("shen.external-symbols")
var symshen_4extract_1free_1vars = MakeSymbol("shen.extract-free-vars")
var symshen_4extract_1vars = MakeSymbol("shen.extract-vars")
var symshen_4extraspecial_2 = MakeSymbol("shen.extraspecial?")
var symshen_4f = MakeSymbol("shen.f")
var symshen_4f_1error = MakeSymbol("shen.f-error")
var symshen_4factor = MakeSymbol("shen.factor")
var symshen_4factor_1cn = MakeSymbol("shen.factor-cn")
var symshen_4factor_1recognisors = MakeSymbol("shen.factor-recognisors")
var symshen_4factor_1selectors = MakeSymbol("shen.factor-selectors")
var symshen_4factor_1selectors_1h = MakeSymbol("shen.factor-selectors-h")
var symshen_4factorise_1code = MakeSymbol("shen.factorise-code")
var symshen_4fail_b = MakeSymbol("shen.fail!")
var symshen_4fbound_2 = MakeSymbol("shen.fbound?")
var symshen_4fillvector = MakeSymbol("shen.fillvector")
var symshen_4find_1arities = MakeSymbol("shen.find-arities")
var symshen_4find_1arity = MakeSymbol("shen.find-arity")
var symshen_4find_1free_1vars = MakeSymbol("shen.find-free-vars")
var symshen_4find_1types = MakeSymbol("shen.find-types")
var symshen_4findall_1h = MakeSymbol("shen.findall-h")
var symshen_4fits_2 = MakeSymbol("shen.fits?")
var symshen_4fix_1help = MakeSymbol("shen.fix-help")
var symshen_4fn_1call = MakeSymbol("shen.fn-call")
var symshen_4fn_1call_2 = MakeSymbol("shen.fn-call?")
var symshen_4fn_1print = MakeSymbol("shen.fn-print")
var symshen_4for_1each = MakeSymbol("shen.for-each")
var symshen_4foreign_2 = MakeSymbol("shen.foreign?")
var symshen_4free_1var_1chk = MakeSymbol("shen.free-var-chk")
var symshen_4free_1variable_1error_1message = MakeSymbol("shen.free-variable-error-message")
var symshen_4free_1variable_2 = MakeSymbol("shen.free-variable?")
var symshen_4freeze_1literals = MakeSymbol("shen.freeze-literals")
var symshen_4freshen = MakeSymbol("shen.freshen")
var symshen_4freshen_1rule = MakeSymbol("shen.freshen-rule")
var symshen_4freshen_1sig = MakeSymbol("shen.freshen-sig")
var symshen_4freshen_1type = MakeSymbol("shen.freshen-type")
var symshen_4freshterm = MakeSymbol("shen.freshterm")
var symshen_4freshterm_2 = MakeSymbol("shen.freshterm?")
var symshen_4freshterms = MakeSymbol("shen.freshterms")
var symshen_4function_1calls = MakeSymbol("shen.function-calls")
var symshen_4funexstring = MakeSymbol("shen.funexstring")
var symshen_4gc = MakeSymbol("shen.gc")
var symshen_4get_1profile = MakeSymbol("shen.get-profile")
var symshen_4goto = MakeSymbol("shen.goto")
var symshen_4goto_1h = MakeSymbol("shen.goto-h")
var symshen_4hascut_2 = MakeSymbol("shen.hascut?")
var symshen_4hashkey = MakeSymbol("shen.hashkey")
var symshen_4hds_a_2 = MakeSymbol("shen.hds=?")
var symshen_4horn_1clause_1procedure = MakeSymbol("shen.horn-clause-procedure")
var symshen_4i_1failed_b = MakeSymbol("shen.i-failed!")
var symshen_4in_1_6 = MakeSymbol("shen.in->")
var symshen_4incinfs = MakeSymbol("shen.incinfs")
var symshen_4initialise = MakeSymbol("shen.initialise")
var symshen_4initialise_1arity_1table = MakeSymbol("shen.initialise-arity-table")
var symshen_4initialise_1environment = MakeSymbol("shen.initialise-environment")
var symshen_4initialise_1lambda_1forms = MakeSymbol("shen.initialise-lambda-forms")
var symshen_4initialise_1signedfuncs = MakeSymbol("shen.initialise-signedfuncs")
var symshen_4initialise__environment = MakeSymbol("shen.initialise_environment")
var symshen_4input_1h_7 = MakeSymbol("shen.input-h+")
var symshen_4input_1track = MakeSymbol("shen.input-track")
var symshen_4insert = MakeSymbol("shen.insert")
var symshen_4insert_1h = MakeSymbol("shen.insert-h")
var symshen_4insert_1info = MakeSymbol("shen.insert-info")
var symshen_4insert_1l = MakeSymbol("shen.insert-l")
var symshen_4insert_1prolog_1variables = MakeSymbol("shen.insert-prolog-variables")
var symshen_4insert_1tracking_1code = MakeSymbol("shen.insert-tracking-code")
var symshen_4integer_1test_2 = MakeSymbol("shen.integer-test?")
var symshen_4intern_1in_1package = MakeSymbol("shen.intern-in-package")
var symshen_4intern_1type = MakeSymbol("shen.intern-type")
var symshen_4internal_1symbols = MakeSymbol("shen.internal-symbols")
var symshen_4internal_1to_1P_2 = MakeSymbol("shen.internal-to-P?")
var symshen_4internal_1to_1shen_2 = MakeSymbol("shen.internal-to-shen?")
var symshen_4internal_2 = MakeSymbol("shen.internal?")
var symshen_4invoke = MakeSymbol("shen.invoke")
var symshen_4iter_1list = MakeSymbol("shen.iter-list")
var symshen_4iter_1vector = MakeSymbol("shen.iter-vector")
var symshen_4key_1in_1sequent_1calculus_2 = MakeSymbol("shen.key-in-sequent-calculus?")
var symshen_4kl_1body = MakeSymbol("shen.kl-body")
var symshen_4klfile = MakeSymbol("shen.klfile")
var symshen_4l_1rules = MakeSymbol("shen.l-rules")
var symshen_4lambda_1entry = MakeSymbol("shen.lambda-entry")
var symshen_4lambda_1form = MakeSymbol("shen.lambda-form")
var symshen_4lambda_1function = MakeSymbol("shen.lambda-function")
var symshen_4lambda_1of_1defun = MakeSymbol("shen.lambda-of-defun")
var symshen_4lazyderef = MakeSymbol("shen.lazyderef")
var symshen_4lch = MakeSymbol("shen.lch")
var symshen_4lchh = MakeSymbol("shen.lchh")
var symshen_4length_1h = MakeSymbol("shen.length-h")
var symshen_4line = MakeSymbol("shen.line")
var symshen_4linearise = MakeSymbol("shen.linearise")
var symshen_4linearise_1clause = MakeSymbol("shen.linearise-clause")
var symshen_4linearise_1h = MakeSymbol("shen.linearise-h")
var symshen_4list_1_6str = MakeSymbol("shen.list->str")
var symshen_4list_2 = MakeSymbol("shen.list?")
var symshen_4load_1help = MakeSymbol("shen.load-help")
var symshen_4loading_2 = MakeSymbol("shen.loading?")
var symshen_4lock = MakeSymbol("shen.lock")
var symshen_4locked_2 = MakeSymbol("shen.locked?")
var symshen_4lookupsig = MakeSymbol("shen.lookupsig")
var symshen_4loop = MakeSymbol("shen.loop")
var symshen_4lowercase_1symbol_2 = MakeSymbol("shen.lowercase-symbol?")
var symshen_4lowercase_2 = MakeSymbol("shen.lowercase?")
var symshen_4lr_1rule = MakeSymbol("shen.lr-rule")
var symshen_4lzy_a = MakeSymbol("shen.lzy=")
var symshen_4lzy_a_b = MakeSymbol("shen.lzy=!")
var symshen_4macro_1_8c = MakeSymbol("shen.macro-@c")
var symshen_4macro_1_8ch = MakeSymbol("shen.macro-@ch")
var symshen_4macroexpand_1h = MakeSymbol("shen.macroexpand-h")
var symshen_4macros = MakeSymbol("shen.macros")
var symshen_4magless = MakeSymbol("shen.magless")
var symshen_4make_1prolog_1variable = MakeSymbol("shen.make-prolog-variable")
var symshen_4make_1uppercase = MakeSymbol("shen.make-uppercase")
var symshen_4map_1h = MakeSymbol("shen.map-h")
var symshen_4maxinfexceeded_2 = MakeSymbol("shen.maxinfexceeded?")
var symshen_4maxseq = MakeSymbol("shen.maxseq")
var symshen_4member = MakeSymbol("shen.member")
var symshen_4member_1clause = MakeSymbol("shen.member-clause")
var symshen_4misc_2 = MakeSymbol("shen.misc?")
var symshen_4mkstr = MakeSymbol("shen.mkstr")
var symshen_4mkstr_1l = MakeSymbol("shen.mkstr-l")
var symshen_4mkstr_1r = MakeSymbol("shen.mkstr-r")
var symshen_4mod = MakeSymbol("shen.mod")
var symshen_4modh = MakeSymbol("shen.modh")
var symshen_4monomorphic_2 = MakeSymbol("shen.monomorphic?")
var symshen_4monotype = MakeSymbol("shen.monotype")
var symshen_4mu_1h = MakeSymbol("shen.mu-h")
var symshen_4multiples = MakeSymbol("shen.multiples")
var symshen_4my_1read_1byte = MakeSymbol("shen.my-read-byte")
var symshen_4myassume = MakeSymbol("shen.myassume")
var symshen_4newname = MakeSymbol("shen.newname")
var symshen_4newpv = MakeSymbol("shen.newpv")
var symshen_4nextticket = MakeSymbol("shen.nextticket")
var symshen_4non_1application_2 = MakeSymbol("shen.non-application?")
var symshen_4non_1terminal_2 = MakeSymbol("shen.non-terminal?")
var symshen_4non_1terminalcode = MakeSymbol("shen.non-terminalcode")
var symshen_4not_1dictionary = MakeSymbol("shen.not-dictionary")
var symshen_4not_1pvar = MakeSymbol("shen.not-pvar")
var symshen_4not_1tuple = MakeSymbol("shen.not-tuple")
var symshen_4nothing_1doing_2 = MakeSymbol("shen.nothing-doing?")
var symshen_4nvars = MakeSymbol("shen.nvars")
var symshen_4objectcode = MakeSymbol("shen.objectcode")
var symshen_4occurs_1check_2 = MakeSymbol("shen.occurs-check?")
var symshen_4ok = MakeSymbol("shen.ok")
var symshen_4op = MakeSymbol("shen.op")
var symshen_4op_1test = MakeSymbol("shen.op-test")
var symshen_4op1 = MakeSymbol("shen.op1")
var symshen_4op2 = MakeSymbol("shen.op2")
var symshen_4openlock = MakeSymbol("shen.openlock")
var symshen_4out_1of_1bounds = MakeSymbol("shen.out-of-bounds")
var symshen_4output_1track = MakeSymbol("shen.output-track")
var symshen_4overapplication_2 = MakeSymbol("shen.overapplication?")
var symshen_4overbind = MakeSymbol("shen.overbind")
var symshen_4p_1hyps = MakeSymbol("shen.p-hyps")
var symshen_4pac_1h = MakeSymbol("shen.pac-h")
var symshen_4package_1symbols = MakeSymbol("shen.package-symbols")
var symshen_4package_1user_1input = MakeSymbol("shen.package-user-input")
var symshen_4packaged_2 = MakeSymbol("shen.packaged?")
var symshen_4parameters = MakeSymbol("shen.parameters")
var symshen_4parse_1failure = MakeSymbol("shen.parse-failure")
var symshen_4parse_1failure_2 = MakeSymbol("shen.parse-failure?")
var symshen_4partial = MakeSymbol("shen.partial")
var symshen_4partial_1application_d_2 = MakeSymbol("shen.partial-application*?")
var symshen_4partial_1parse_1failure_2 = MakeSymbol("shen.partial-parse-failure?")
var symshen_4passive_1bind = MakeSymbol("shen.passive-bind")
var symshen_4passive_1variables = MakeSymbol("shen.passive-variables")
var symshen_4pause_1for_1user = MakeSymbol("shen.pause-for-user")
var symshen_4peek_1history = MakeSymbol("shen.peek-history")
var symshen_4pivot_1on = MakeSymbol("shen.pivot-on")
var symshen_4posint_2 = MakeSymbol("shen.posint?")
var symshen_4precompiled_1function_2 = MakeSymbol("shen.precompiled-function?")
var symshen_4predicate = MakeSymbol("shen.predicate")
var symshen_4premises_1_6goals = MakeSymbol("shen.premises->goals")
var symshen_4pretty_1type = MakeSymbol("shen.pretty-type")
var symshen_4prhush = MakeSymbol("shen.prhush")
var symshen_4primitive = MakeSymbol("shen.primitive")
var symshen_4print_1freshterm = MakeSymbol("shen.print-freshterm")
var symshen_4print_1prolog_1vector = MakeSymbol("shen.print-prolog-vector")
var symshen_4print_1vector_2 = MakeSymbol("shen.print-vector?")
var symshen_4printF = MakeSymbol("shen.printF")
var symshen_4proc_1nl = MakeSymbol("shen.proc-nl")
var symshen_4process_1_8s = MakeSymbol("shen.process-@s")
var symshen_4process_1after_1type = MakeSymbol("shen.process-after-type")
var symshen_4process_1application = MakeSymbol("shen.process-application")
var symshen_4process_1applications = MakeSymbol("shen.process-applications")
var symshen_4process_1assoc = MakeSymbol("shen.process-assoc")
var symshen_4process_1cases = MakeSymbol("shen.process-cases")
var symshen_4process_1cond_1clauses = MakeSymbol("shen.process-cond-clauses")
var symshen_4process_1datatype = MakeSymbol("shen.process-datatype")
var symshen_4process_1def = MakeSymbol("shen.process-def")
var symshen_4process_1input_7 = MakeSymbol("shen.process-input+")
var symshen_4process_1lambda = MakeSymbol("shen.process-lambda")
var symshen_4process_1let = MakeSymbol("shen.process-let")
var symshen_4process_1read_1byte = MakeSymbol("shen.process-read-byte")
var symshen_4process_1sexprs = MakeSymbol("shen.process-sexprs")
var symshen_4process_1synonyms = MakeSymbol("shen.process-synonyms")
var symshen_4process_1time = MakeSymbol("shen.process-time")
var symshen_4process_1yacc_1semantics = MakeSymbol("shen.process-yacc-semantics")
var symshen_4processed = MakeSymbol("shen.processed")
var symshen_4prodbutzero = MakeSymbol("shen.prodbutzero")
var symshen_4profile_1func = MakeSymbol("shen.profile-func")
var symshen_4profile_1help = MakeSymbol("shen.profile-help")
var symshen_4profiled_2 = MakeSymbol("shen.profiled?")
var symshen_4prolog_1abstraction = MakeSymbol("shen.prolog-abstraction")
var symshen_4prolog_1arity_1check = MakeSymbol("shen.prolog-arity-check")
var symshen_4prolog_1fbody = MakeSymbol("shen.prolog-fbody")
var symshen_4prolog_1keyword_2 = MakeSymbol("shen.prolog-keyword?")
var symshen_4prolog_1parameters = MakeSymbol("shen.prolog-parameters")
var symshen_4prolog_1track = MakeSymbol("shen.prolog-track")
var symshen_4prolog_1vector = MakeSymbol("shen.prolog-vector")
var symshen_4prolog_1vector_1size = MakeSymbol("shen.prolog-vector-size")
var symshen_4prompt = MakeSymbol("shen.prompt")
var symshen_4prterm = MakeSymbol("shen.prterm")
var symshen_4prtl = MakeSymbol("shen.prtl")
var symshen_4pui_1h = MakeSymbol("shen.pui-h")
var symshen_4put_1profile = MakeSymbol("shen.put-profile")
var symshen_4pvar = MakeSymbol("shen.pvar")
var symshen_4pvar_2 = MakeSymbol("shen.pvar?")
var symshen_4r = MakeSymbol("shen.r")
var symshen_4raise_1syntax_1error = MakeSymbol("shen.raise-syntax-error")
var symshen_4rcons__form = MakeSymbol("shen.rcons_form")
var symshen_4rdecons = MakeSymbol("shen.rdecons")
var symshen_4read_1evaluate_1print = MakeSymbol("shen.read-evaluate-print")
var symshen_4read_1file_1as_1bytelist_1help = MakeSymbol("shen.read-file-as-bytelist-help")
var symshen_4read_1loop = MakeSymbol("shen.read-loop")
var symshen_4reader_1error = MakeSymbol("shen.reader-error")
var symshen_4reader_1error_1message = MakeSymbol("shen.reader-error-message")
var symshen_4received = MakeSymbol("shen.received")
var symshen_4record_1and_1evaluate = MakeSymbol("shen.record-and-evaluate")
var symshen_4record_1external = MakeSymbol("shen.record-external")
var symshen_4record_1internal = MakeSymbol("shen.record-internal")
var symshen_4record_1it = MakeSymbol("shen.record-it")
var symshen_4record_1kl = MakeSymbol("shen.record-kl")
var symshen_4record_1macro = MakeSymbol("shen.record-macro")
var symshen_4rectify_1test = MakeSymbol("shen.rectify-test")
var symshen_4rectify_1type = MakeSymbol("shen.rectify-type")
var symshen_4recursive_1string_1match = MakeSymbol("shen.recursive-string-match")
var symshen_4recursively_1factor_1selectors = MakeSymbol("shen.recursively-factor-selectors")
var symshen_4recursively_1print = MakeSymbol("shen.recursively-print")
var symshen_4register_1source_1form = MakeSymbol("shen.register-source-form")
var symshen_4remember_1datatype = MakeSymbol("shen.remember-datatype")
var symshen_4remove_1bystanders = MakeSymbol("shen.remove-bystanders")
var symshen_4remove_1datatypes = MakeSymbol("shen.remove-datatypes")
var symshen_4remove_1h = MakeSymbol("shen.remove-h")
var symshen_4remove_1indirection = MakeSymbol("shen.remove-indirection")
var symshen_4rep_1X = MakeSymbol("shen.rep-X")
var symshen_4repl = MakeSymbol("shen.repl")
var symshen_4restore_1P = MakeSymbol("shen.restore-P")
var symshen_4restore_1local = MakeSymbol("shen.restore-local")
var symshen_4retract_1clause = MakeSymbol("shen.retract-clause")
var symshen_4return_2 = MakeSymbol("shen.return?")
var symshen_4reverse_1help = MakeSymbol("shen.reverse-help")
var symshen_4rfas_1h = MakeSymbol("shen.rfas-h")
var symshen_4rule_1_6body = MakeSymbol("shen.rule->body")
var symshen_4rule_1_6clause = MakeSymbol("shen.rule->clause")
var symshen_4rule_1_6head = MakeSymbol("shen.rule->head")
var symshen_4rules_1_6prolog = MakeSymbol("shen.rules->prolog")
var symshen_4s = MakeSymbol("shen.s")
var symshen_4scan_1body = MakeSymbol("shen.scan-body")
var symshen_4search_1user_1datatypes = MakeSymbol("shen.search-user-datatypes")
var symshen_4search_1user_1datatypes_1h = MakeSymbol("shen.search-user-datatypes-h")
var symshen_4semicolon_2 = MakeSymbol("shen.semicolon?")
var symshen_4set_1lambda_1form_1entry = MakeSymbol("shen.set-lambda-form-entry")
var symshen_4shen_1_6kl = MakeSymbol("shen.shen->kl")
var symshen_4shen_1_6kl_1h = MakeSymbol("shen.shen->kl-h")
var symshen_4shen_1call_2 = MakeSymbol("shen.shen-call?")
var symshen_4shendef_1_6kldef = MakeSymbol("shen.shendef->kldef")
var symshen_4shendef_1_6kldef_1h = MakeSymbol("shen.shendef->kldef-h")
var symshen_4show = MakeSymbol("shen.show")
var symshen_4show_1assumptions = MakeSymbol("shen.show-assumptions")
var symshen_4show_1datatypes = MakeSymbol("shen.show-datatypes")
var symshen_4show_1p = MakeSymbol("shen.show-p")
var symshen_4side_1conditions_1_6goals = MakeSymbol("shen.side-conditions->goals")
var symshen_4sigf = MakeSymbol("shen.sigf")
var symshen_4signal_1def = MakeSymbol("shen.signal-def")
var symshen_4sigxrules = MakeSymbol("shen.sigxrules")
var symshen_4simple_1curry = MakeSymbol("shen.simple-curry")
var symshen_4skip = MakeSymbol("shen.skip")
var symshen_4sng_1h_2 = MakeSymbol("shen.sng-h?")
var symshen_4sng_2 = MakeSymbol("shen.sng?")
var symshen_4source = MakeSymbol("shen.source")
var symshen_4source_1form_1expansion = MakeSymbol("shen.source-form-expansion")
var symshen_4source_1form_1handler = MakeSymbol("shen.source-form-handler")
var symshen_4source_1form_1list_2 = MakeSymbol("shen.source-form-list?")
var symshen_4source_1macro_1walk = MakeSymbol("shen.source-macro-walk")
var symshen_4spaces = MakeSymbol("shen.spaces")
var symshen_4special_1case = MakeSymbol("shen.special-case")
var symshen_4special_2 = MakeSymbol("shen.special?")
var symshen_4specialise_1consume = MakeSymbol("shen.specialise-consume")
var symshen_4specialise_1member = MakeSymbol("shen.specialise-member")
var symshen_4store_1arity = MakeSymbol("shen.store-arity")
var symshen_4stpart = MakeSymbol("shen.stpart")
var symshen_4str_1_6bytes = MakeSymbol("shen.str->bytes")
var symshen_4str_1_6str = MakeSymbol("shen.str->str")
var symshen_4string_1_6byte = MakeSymbol("shen.string->byte")
var symshen_4string_1_6bytes = MakeSymbol("shen.string->bytes")
var symshen_4string_1match = MakeSymbol("shen.string-match")
var symshen_4string_1prefix_2 = MakeSymbol("shen.string-prefix?")
var symshen_4synonyms_1h = MakeSymbol("shen.synonyms-h")
var symshen_4syntax_1error_1message = MakeSymbol("shen.syntax-error-message")
var symshen_4syntax_1item_2 = MakeSymbol("shen.syntax-item?")
var symshen_4sysfunc_2 = MakeSymbol("shen.sysfunc?")
var symshen_4system_1S = MakeSymbol("shen.system-S")
var symshen_4system_1S_1h = MakeSymbol("shen.system-S-h")
var symshen_4t = MakeSymbol("shen.t")
var symshen_4t_d = MakeSymbol("shen.t*")
var symshen_4t_d_1correct = MakeSymbol("shen.t*-correct")
var symshen_4t_d_1integrity = MakeSymbol("shen.t*-integrity")
var symshen_4t_d_1rule = MakeSymbol("shen.t*-rule")
var symshen_4t_d_1rule_1h = MakeSymbol("shen.t*-rule-h")
var symshen_4t_d_1rules = MakeSymbol("shen.t*-rules")
var symshen_4tame = MakeSymbol("shen.tame")
var symshen_4terminalcode = MakeSymbol("shen.terminalcode")
var symshen_4terms = MakeSymbol("shen.terms")
var symshen_4terpri_1or_1read_1char = MakeSymbol("shen.terpri-or-read-char")
var symshen_4this_1symbol_1is_1unbound = MakeSymbol("shen.this-symbol-is-unbound")
var symshen_4ticket_1number = MakeSymbol("shen.ticket-number")
var symshen_4tlv_1help = MakeSymbol("shen.tlv-help")
var symshen_4top = MakeSymbol("shen.top")
var symshen_4toplevel_1display_1exception = MakeSymbol("shen.toplevel-display-exception")
var symshen_4toplevel_1forms = MakeSymbol("shen.toplevel-forms")
var symshen_4track_1function = MakeSymbol("shen.track-function")
var symshen_4tracked_2 = MakeSymbol("shen.tracked?")
var symshen_4trim_1it = MakeSymbol("shen.trim-it")
var symshen_4triple_1stack = MakeSymbol("shen.triple-stack")
var symshen_4try_1parse = MakeSymbol("shen.try-parse")
var symshen_4tuple = MakeSymbol("shen.tuple")
var symshen_4tuple_1up = MakeSymbol("shen.tuple-up")
var symshen_4type_1F = MakeSymbol("shen.type-F")
var symshen_4type_1error = MakeSymbol("shen.type-error")
var symshen_4type_1theory_1enabled_2 = MakeSymbol("shen.type-theory-enabled?")
var symshen_4typecheck = MakeSymbol("shen.typecheck")
var symshen_4typename = MakeSymbol("shen.typename")
var symshen_4typename_1h = MakeSymbol("shen.typename-h")
var symshen_4typetable = MakeSymbol("shen.typetable")
var symshen_4unassoc = MakeSymbol("shen.unassoc")
var symshen_4undefined_1f_2 = MakeSymbol("shen.undefined-f?")
var symshen_4unlock = MakeSymbol("shen.unlock")
var symshen_4unlocked_2 = MakeSymbol("shen.unlocked?")
var symshen_4unmark_1custom_1pattern = MakeSymbol("shen.unmark-custom-pattern")
var symshen_4unmark_1rule_1custom_1patterns = MakeSymbol("shen.unmark-rule-custom-patterns")
var symshen_4unpack_1foreign = MakeSymbol("shen.unpack-foreign")
var symshen_4unpackage = MakeSymbol("shen.unpackage")
var symshen_4unpackage_emacroexpand = MakeSymbol("shen.unpackage&macroexpand")
var symshen_4unprotect = MakeSymbol("shen.unprotect")
var symshen_4unregister_1source_1form = MakeSymbol("shen.unregister-source-form")
var symshen_4unwind = MakeSymbol("shen.unwind")
var symshen_4unwind_1types = MakeSymbol("shen.unwind-types")
var symshen_4update_1assoc = MakeSymbol("shen.update-assoc")
var symshen_4update_1history = MakeSymbol("shen.update-history")
var symshen_4update_1lambda_1table_cprecompiled = MakeSymbol("shen.update-lambda-table/precompiled")
var symshen_4update_1lambdatable = MakeSymbol("shen.update-lambdatable")
var symshen_4uppercase_2 = MakeSymbol("shen.uppercase?")
var symshen_4use_1history = MakeSymbol("shen.use-history")
var symshen_4use_1type_1info = MakeSymbol("shen.use-type-info")
var symshen_4variable_1case = MakeSymbol("shen.variable-case")
var symshen_4variablecode = MakeSymbol("shen.variablecode")
var symshen_4variancy = MakeSymbol("shen.variancy")
var symshen_4variancy_1signature = MakeSymbol("shen.variancy-signature")
var symshen_4variants_2 = MakeSymbol("shen.variants?")
var symshen_4vector_1_6str = MakeSymbol("shen.vector->str")
var symshen_4vector_1dereference = MakeSymbol("shen.vector-dereference")
var symshen_4vector_1parameter = MakeSymbol("shen.vector-parameter")
var symshen_4walk = MakeSymbol("shen.walk")
var symshen_4whitespace_2 = MakeSymbol("shen.whitespace?")
var symshen_4wildcard_2 = MakeSymbol("shen.wildcard?")
var symshen_4wildcardcode = MakeSymbol("shen.wildcardcode")
var symshen_4work_1through = MakeSymbol("shen.work-through")
var symshen_4write_1chars = MakeSymbol("shen.write-chars")
var symshen_4write_1kl = MakeSymbol("shen.write-kl")
var symshen_4write_1kl_1h = MakeSymbol("shen.write-kl-h")
var symshen_4x_4launcher_4default_1handle_1result = MakeSymbol("shen.x.launcher.default-handle-result")
var symshen_4x_4launcher_4done = MakeSymbol("shen.x.launcher.done")
var symshen_4x_4launcher_4eval_1command = MakeSymbol("shen.x.launcher.eval-command")
var symshen_4x_4launcher_4eval_1command_1h = MakeSymbol("shen.x.launcher.eval-command-h")
var symshen_4x_4launcher_4eval_1flag_1map = MakeSymbol("shen.x.launcher.eval-flag-map")
var symshen_4x_4launcher_4eval_1string = MakeSymbol("shen.x.launcher.eval-string")
var symshen_4x_4launcher_4execute_1all = MakeSymbol("shen.x.launcher.execute-all")
var symshen_4x_4launcher_4help_1text = MakeSymbol("shen.x.launcher.help-text")
var symshen_4x_4launcher_4launch_1shen = MakeSymbol("shen.x.launcher.launch-shen")
var symshen_4x_4launcher_4main = MakeSymbol("shen.x.launcher.main")
var symshen_4x_4launcher_4quiet_1load = MakeSymbol("shen.x.launcher.quiet-load")
var symshen_4x_4launcher_4script_1command = MakeSymbol("shen.x.launcher.script-command")
var symshen_4x_4launcher_4version_1string = MakeSymbol("shen.x.launcher.version-string")
var symshen_4yacc_1_6shen = MakeSymbol("shen.yacc->shen")
var symshen_4yacc_1semantics = MakeSymbol("shen.yacc-semantics")
var symshen_4yacc_1syntax = MakeSymbol("shen.yacc-syntax")
var symshen_4zero_1place_2 = MakeSymbol("shen.zero-place?")
var symshow_1help = MakeSymbol("show-help")
var symsimple_1error = MakeSymbol("simple-error")
var symsnd = MakeSymbol("snd")
var symspecialise = MakeSymbol("specialise")
var symspy = MakeSymbol("spy")
var symspy_2 = MakeSymbol("spy?")
var symsqts = MakeSymbol("sqts")
var symstep = MakeSymbol("step")
var symstep_2 = MakeSymbol("step?")
var symsterror = MakeSymbol("sterror")
var symstinput = MakeSymbol("stinput")
var symstoutput = MakeSymbol("stoutput")
var symstr = MakeSymbol("str")
var symstream = MakeSymbol("stream")
var symstring = MakeSymbol("string")
var symstring_1_6n = MakeSymbol("string->n")
var symstring_1_6symbol = MakeSymbol("string->symbol")
var symstring_2 = MakeSymbol("string?")
var symsubst = MakeSymbol("subst")
var symsuccess = MakeSymbol("success")
var symsum = MakeSymbol("sum")
var symsymbol = MakeSymbol("symbol")
var symsymbol_2 = MakeSymbol("symbol?")
var symsynonyms = MakeSymbol("synonyms")
var symsystem_1S_2 = MakeSymbol("system-S?")
var symsystemf = MakeSymbol("systemf")
var symtail = MakeSymbol("tail")
var symtc = MakeSymbol("tc")
var symtc_2 = MakeSymbol("tc?")
var symthaw = MakeSymbol("thaw")
var symtime = MakeSymbol("time")
var symtl = MakeSymbol("tl")
var symtlstr = MakeSymbol("tlstr")
var symtlv = MakeSymbol("tlv")
var symtrack = MakeSymbol("track")
var symtracked = MakeSymbol("tracked")
var symtrap_1error = MakeSymbol("trap-error")
var symtuple_2 = MakeSymbol("tuple?")
var symtype = MakeSymbol("type")
var symu_b = MakeSymbol("u!")
var symunabsolute = MakeSymbol("unabsolute")
var symundefmacro = MakeSymbol("undefmacro")
var symunion = MakeSymbol("union")
var symunit = MakeSymbol("unit")
var symunix = MakeSymbol("unix")
var symunknown_1arguments = MakeSymbol("unknown-arguments")
var symunprofile = MakeSymbol("unprofile")
var symunput = MakeSymbol("unput")
var symunspecialise = MakeSymbol("unspecialise")
var symuntrack = MakeSymbol("untrack")
var symupdate_1lambda_1table = MakeSymbol("update-lambda-table")
var symuserdefs = MakeSymbol("userdefs")
var symvalue = MakeSymbol("value")
var symvar_2 = MakeSymbol("var?")
var symvariable_2 = MakeSymbol("variable?")
var symvector = MakeSymbol("vector")
var symvector_1_6 = MakeSymbol("vector->")
var symvector_2 = MakeSymbol("vector?")
var symverified = MakeSymbol("verified")
var symversion = MakeSymbol("version")
var symwarn = MakeSymbol("warn")
var symwhen = MakeSymbol("when")
var symwhere = MakeSymbol("where")
var symwrite_1byte = MakeSymbol("write-byte")
var symwrite_1string = MakeSymbol("write-string")
var symwrite_1to_1file = MakeSymbol("write-to-file")
var symy_1or_1n_2 = MakeSymbol("y-or-n?")
var sym_i = MakeSymbol("{")
var sym_j = MakeSymbol("}")
