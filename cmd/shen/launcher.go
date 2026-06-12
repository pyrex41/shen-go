package main

import . "github.com/tiancaiamao/shen-go/kl"

var LauncherMain = MakeNative(func(__e *ControlFlow) {
_ = MakeString("Copyright (c) 2019 Bruno Deferrari.\nBSD 3-Clause License: http://opensource.org/licenses/BSD-3-Clause")

tmp20112 := MakeNative(func(__e *ControlFlow) {
V7102 := __e.Get(1)
_ = V7102
tmp20113 := MakeNative(func(__e *ControlFlow) {
W7103 := __e.Get(1)
_ = W7103
tmp20114 := MakeNative(func(__e *ControlFlow) {
Z7104 := __e.Get(1)
_ = Z7104
__e.TailApply(PrimFunc(symeval), Z7104)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp20114, W7103)
return


}, 1)

tmp20115 := Call(__e, PrimFunc(symread_1file), V7102)


__e.TailApply(tmp20113, tmp20115)
return


}, 1)

tmp20116 := Call(__e, ns2_1set, symshen_4x_4launcher_4quiet_1load, tmp20112)


_ = tmp20116

tmp20117 := MakeNative(func(__e *ControlFlow) {
tmp20118 := Call(__e, PrimFunc(symversion))


tmp20119 := Call(__e, PrimFunc(symlanguage))


tmp20120 := Call(__e, PrimFunc(symport))


tmp20121 := PrimCons(tmp20120, Nil)

tmp20122 := PrimCons(tmp20119, tmp20121)

tmp20123 := Call(__e, PrimFunc(symimplementation))


tmp20124 := Call(__e, PrimFunc(symrelease))


tmp20125 := PrimCons(tmp20124, Nil)

tmp20126 := PrimCons(tmp20123, tmp20125)

tmp20127 := PrimCons(tmp20126, Nil)

tmp20128 := PrimCons(symimplementation, tmp20127)

tmp20129 := PrimCons(tmp20122, tmp20128)

tmp20130 := PrimCons(symport, tmp20129)

tmp20131 := Call(__e, PrimFunc(symshen_4app), tmp20130, MakeString("\n"), symshen_4r)


tmp20132 := PrimStringConcat(MakeString(" "), tmp20131)

__e.TailApply(PrimFunc(symshen_4app), tmp20118, tmp20132, symshen_4a)
return


}, 0)

tmp20133 := Call(__e, ns2_1set, symshen_4x_4launcher_4version_1string, tmp20117)


_ = tmp20133

tmp20134 := MakeNative(func(__e *ControlFlow) {
V7105 := __e.Get(1)
_ = V7105
tmp20135 := Call(__e, PrimFunc(symshen_4app), V7105, MakeString(" [--version] [--help] <COMMAND> [<ARGS>]\n\ncommands:\n    repl\n        Launches the interactive REPL.\n        Default action if no command is supplied.\n\n    script <FILE> [<ARGS>]\n        Runs the script in FILE. *argv* is set to [FILE | ARGS].\n\n    eval <ARGS>\n        Evaluates expressions and files. ARGS are evaluated from\n        left to right and can be a combination of:\n            -e, --eval <EXPR>\n                Evaluates EXPR and prints result.\n            -l, --load <FILE>\n                Reads and evaluates FILE.\n            -q, --quiet\n                Silences interactive output.\n            -s, --set <KEY> <VALUE>\n                Evaluates KEY, VALUE and sets as global.\n            -r, --repl\n                Launches the interactive REPL after evaluating\n                all the previous expresions."), symshen_4a)


__e.Return(PrimStringConcat(MakeString("Usage: "), tmp20135))
return


}, 1)

tmp20136 := Call(__e, ns2_1set, symshen_4x_4launcher_4help_1text, tmp20134)


_ = tmp20136

tmp20137 := MakeNative(func(__e *ControlFlow) {
V7106 := __e.Get(1)
_ = V7106
tmp20144 := PrimEqual(Nil, V7106)

if True == tmp20144 {
__e.Return(PrimCons(symsuccess, Nil))
return
} else {
tmp20142 := PrimIsPair(V7106)

if True == tmp20142 {
tmp20138 := PrimHead(V7106)

tmp20139 := Call(__e, PrimFunc(symthaw), tmp20138)


_ = tmp20139

tmp20140 := PrimTail(V7106)

__e.TailApply(PrimFunc(symshen_4x_4launcher_4execute_1all), tmp20140)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4x_4launcher_4execute_1all)
return
}


}


}, 1)

tmp20145 := Call(__e, ns2_1set, symshen_4x_4launcher_4execute_1all, tmp20137)


_ = tmp20145

tmp20146 := MakeNative(func(__e *ControlFlow) {
V7107 := __e.Get(1)
_ = V7107
tmp20147 := Call(__e, PrimFunc(symread_1from_1string), V7107)


tmp20148 := Call(__e, PrimFunc(symhead), tmp20147)


__e.TailApply(PrimFunc(symeval), tmp20148)
return


}, 1)

tmp20149 := Call(__e, ns2_1set, symshen_4x_4launcher_4eval_1string, tmp20146)


_ = tmp20149

tmp20150 := MakeNative(func(__e *ControlFlow) {
V7110 := __e.Get(1)
_ = V7110
tmp20160 := PrimEqual(MakeString("-e"), V7110)

if True == tmp20160 {
__e.Return(MakeString("--eval"))
return
} else {
tmp20158 := PrimEqual(MakeString("-l"), V7110)

if True == tmp20158 {
__e.Return(MakeString("--load"))
return
} else {
tmp20156 := PrimEqual(MakeString("-q"), V7110)

if True == tmp20156 {
__e.Return(MakeString("--quiet"))
return
} else {
tmp20154 := PrimEqual(MakeString("-s"), V7110)

if True == tmp20154 {
__e.Return(MakeString("--set"))
return
} else {
tmp20152 := PrimEqual(MakeString("-r"), V7110)

if True == tmp20152 {
__e.Return(MakeString("--repl"))
return
} else {
__e.Return(False)
return
}


}


}


}


}


}, 1)

tmp20161 := Call(__e, ns2_1set, symshen_4x_4launcher_4eval_1flag_1map, tmp20150)


_ = tmp20161

tmp20162 := MakeNative(func(__e *ControlFlow) {
V7115 := __e.Get(1)
_ = V7115
V7116 := __e.Get(2)
_ = V7116
tmp20266 := PrimEqual(Nil, V7115)

if True == tmp20266 {
tmp20163 := Call(__e, PrimFunc(symreverse), V7116)


__e.TailApply(PrimFunc(symshen_4x_4launcher_4execute_1all), tmp20163)
return


} else {
tmp20264 := PrimIsPair(V7115)

var ifres20256 Obj

if True == tmp20264 {
tmp20262 := PrimHead(V7115)

tmp20263 := PrimEqual(MakeString("--eval"), tmp20262)

var ifres20258 Obj

if True == tmp20263 {
tmp20260 := PrimTail(V7115)

tmp20261 := PrimIsPair(tmp20260)

var ifres20259 Obj

if True == tmp20261 {
ifres20259 = True


} else {
ifres20259 = False


}

ifres20258 = ifres20259


} else {
ifres20258 = False


}

var ifres20257 Obj

if True == ifres20258 {
ifres20257 = True


} else {
ifres20257 = False


}

ifres20256 = ifres20257


} else {
ifres20256 = False


}

if True == ifres20256 {
tmp20164 := PrimTail(V7115)

tmp20165 := PrimTail(tmp20164)

tmp20166 := MakeNative(func(__e *ControlFlow) {
tmp20167 := PrimTail(V7115)

tmp20168 := PrimHead(tmp20167)

tmp20169 := Call(__e, PrimFunc(symshen_4x_4launcher_4eval_1string), tmp20168)


tmp20170 := Call(__e, PrimFunc(symshen_4app), tmp20169, MakeString("\n"), symshen_4a)


tmp20171 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp20170, tmp20171)
return


}, 0)

tmp20172 := PrimCons(tmp20166, V7116)

__e.TailApply(PrimFunc(symshen_4x_4launcher_4eval_1command_1h), tmp20165, tmp20172)
return


} else {
tmp20254 := PrimIsPair(V7115)

var ifres20246 Obj

if True == tmp20254 {
tmp20252 := PrimHead(V7115)

tmp20253 := PrimEqual(MakeString("--load"), tmp20252)

var ifres20248 Obj

if True == tmp20253 {
tmp20250 := PrimTail(V7115)

tmp20251 := PrimIsPair(tmp20250)

var ifres20249 Obj

if True == tmp20251 {
ifres20249 = True


} else {
ifres20249 = False


}

ifres20248 = ifres20249


} else {
ifres20248 = False


}

var ifres20247 Obj

if True == ifres20248 {
ifres20247 = True


} else {
ifres20247 = False


}

ifres20246 = ifres20247


} else {
ifres20246 = False


}

if True == ifres20246 {
tmp20173 := PrimTail(V7115)

tmp20174 := PrimTail(tmp20173)

tmp20175 := MakeNative(func(__e *ControlFlow) {
tmp20176 := PrimTail(V7115)

tmp20177 := PrimHead(tmp20176)

__e.TailApply(PrimFunc(symload), tmp20177)
return


}, 0)

tmp20178 := PrimCons(tmp20175, V7116)

__e.TailApply(PrimFunc(symshen_4x_4launcher_4eval_1command_1h), tmp20174, tmp20178)
return


} else {
tmp20244 := PrimIsPair(V7115)

var ifres20240 Obj

if True == tmp20244 {
tmp20242 := PrimHead(V7115)

tmp20243 := PrimEqual(MakeString("--quiet"), tmp20242)

var ifres20241 Obj

if True == tmp20243 {
ifres20241 = True


} else {
ifres20241 = False


}

ifres20240 = ifres20241


} else {
ifres20240 = False


}

if True == ifres20240 {
tmp20179 := PrimTail(V7115)

tmp20180 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimSet(sym_dhush_d, True))
return
}, 0)

tmp20181 := PrimCons(tmp20180, V7116)

__e.TailApply(PrimFunc(symshen_4x_4launcher_4eval_1command_1h), tmp20179, tmp20181)
return


} else {
tmp20238 := PrimIsPair(V7115)

var ifres20225 Obj

if True == tmp20238 {
tmp20236 := PrimHead(V7115)

tmp20237 := PrimEqual(MakeString("--set"), tmp20236)

var ifres20227 Obj

if True == tmp20237 {
tmp20234 := PrimTail(V7115)

tmp20235 := PrimIsPair(tmp20234)

var ifres20229 Obj

if True == tmp20235 {
tmp20231 := PrimTail(V7115)

tmp20232 := PrimTail(tmp20231)

tmp20233 := PrimIsPair(tmp20232)

var ifres20230 Obj

if True == tmp20233 {
ifres20230 = True


} else {
ifres20230 = False


}

ifres20229 = ifres20230


} else {
ifres20229 = False


}

var ifres20228 Obj

if True == ifres20229 {
ifres20228 = True


} else {
ifres20228 = False


}

ifres20227 = ifres20228


} else {
ifres20227 = False


}

var ifres20226 Obj

if True == ifres20227 {
ifres20226 = True


} else {
ifres20226 = False


}

ifres20225 = ifres20226


} else {
ifres20225 = False


}

if True == ifres20225 {
tmp20182 := PrimTail(V7115)

tmp20183 := PrimTail(tmp20182)

tmp20184 := PrimTail(tmp20183)

tmp20185 := MakeNative(func(__e *ControlFlow) {
tmp20186 := PrimTail(V7115)

tmp20187 := PrimHead(tmp20186)

tmp20188 := Call(__e, PrimFunc(symshen_4x_4launcher_4eval_1string), tmp20187)


tmp20189 := PrimTail(V7115)

tmp20190 := PrimTail(tmp20189)

tmp20191 := PrimHead(tmp20190)

tmp20192 := Call(__e, PrimFunc(symshen_4x_4launcher_4eval_1string), tmp20191)


__e.Return(PrimSet(tmp20188, tmp20192))
return


}, 0)

tmp20193 := PrimCons(tmp20185, V7116)

__e.TailApply(PrimFunc(symshen_4x_4launcher_4eval_1command_1h), tmp20184, tmp20193)
return


} else {
tmp20223 := PrimIsPair(V7115)

var ifres20219 Obj

if True == tmp20223 {
tmp20221 := PrimHead(V7115)

tmp20222 := PrimEqual(MakeString("--repl"), tmp20221)

var ifres20220 Obj

if True == tmp20222 {
ifres20220 = True


} else {
ifres20220 = False


}

ifres20219 = ifres20220


} else {
ifres20219 = False


}

if True == ifres20219 {
tmp20194 := Call(__e, PrimFunc(symshen_4x_4launcher_4eval_1command_1h), Nil, V7116)


_ = tmp20194

tmp20195 := PrimTail(V7115)

__e.Return(PrimCons(symlaunch_1repl, tmp20195))
return


} else {
tmp20196 := MakeNative(func(__e *ControlFlow) {
Freeze7119 := __e.Get(1)
_ = Freeze7119
tmp20210 := PrimIsPair(V7115)

if True == tmp20210 {
tmp20197 := MakeNative(func(__e *ControlFlow) {
Result7118 := __e.Get(1)
_ = Result7118
tmp20199 := Call(__e, PrimFunc(symfail))


tmp20200 := PrimEqual(Result7118, tmp20199)

if True == tmp20200 {
__e.TailApply(PrimFunc(symthaw), Freeze7119)
return
} else {
__e.Return(Result7118)
return
}


}, 1)

tmp20201 := MakeNative(func(__e *ControlFlow) {
W7117 := __e.Get(1)
_ = W7117
tmp20205 := PrimEqual(False, W7117)

if True == tmp20205 {
__e.TailApply(PrimFunc(symfail))
return
} else {
tmp20202 := PrimTail(V7115)

tmp20203 := PrimCons(W7117, tmp20202)

__e.TailApply(PrimFunc(symshen_4x_4launcher_4eval_1command_1h), tmp20203, V7116)
return


}


}, 1)

tmp20206 := PrimHead(V7115)

tmp20207 := Call(__e, PrimFunc(symshen_4x_4launcher_4eval_1flag_1map), tmp20206)


tmp20208 := Call(__e, tmp20201, tmp20207)


__e.TailApply(tmp20197, tmp20208)
return


} else {
__e.TailApply(PrimFunc(symthaw), Freeze7119)
return
}


}, 1)

tmp20211 := MakeNative(func(__e *ControlFlow) {
tmp20217 := PrimIsPair(V7115)

if True == tmp20217 {
tmp20212 := PrimHead(V7115)

tmp20213 := Call(__e, PrimFunc(symshen_4app), tmp20212, MakeString(""), symshen_4a)


tmp20214 := PrimStringConcat(MakeString("Invalid eval argument: "), tmp20213)

tmp20215 := PrimCons(tmp20214, Nil)

__e.Return(PrimCons(symerror, tmp20215))
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4x_4launcher_4eval_1command_1h)
return
}


}, 0)

__e.TailApply(tmp20196, tmp20211)
return


}


}


}


}


}


}


}, 2)

tmp20267 := Call(__e, ns2_1set, symshen_4x_4launcher_4eval_1command_1h, tmp20162)


_ = tmp20267

tmp20268 := MakeNative(func(__e *ControlFlow) {
V7120 := __e.Get(1)
_ = V7120
__e.TailApply(PrimFunc(symshen_4x_4launcher_4eval_1command_1h), V7120, Nil)
return
}, 1)

tmp20269 := Call(__e, ns2_1set, symshen_4x_4launcher_4eval_1command, tmp20268)


_ = tmp20269

tmp20270 := MakeNative(func(__e *ControlFlow) {
V7121 := __e.Get(1)
_ = V7121
V7122 := __e.Get(2)
_ = V7122
tmp20271 := PrimCons(V7121, V7122)

tmp20272 := PrimSet(sym_dargv_d, tmp20271)

_ = tmp20272

tmp20273 := Call(__e, PrimFunc(symshen_4x_4launcher_4quiet_1load), V7121)


_ = tmp20273

__e.Return(PrimCons(symsuccess, Nil))
return


}, 2)

tmp20274 := Call(__e, ns2_1set, symshen_4x_4launcher_4script_1command, tmp20270)


_ = tmp20274

tmp20275 := MakeNative(func(__e *ControlFlow) {
V7123 := __e.Get(1)
_ = V7123
tmp20362 := PrimIsPair(V7123)

var ifres20358 Obj

if True == tmp20362 {
tmp20360 := PrimTail(V7123)

tmp20361 := PrimEqual(Nil, tmp20360)

var ifres20359 Obj

if True == tmp20361 {
ifres20359 = True


} else {
ifres20359 = False


}

ifres20358 = ifres20359


} else {
ifres20358 = False


}

if True == ifres20358 {
__e.Return(PrimCons(symlaunch_1repl, Nil))
return
} else {
tmp20356 := PrimIsPair(V7123)

var ifres20347 Obj

if True == tmp20356 {
tmp20354 := PrimTail(V7123)

tmp20355 := PrimIsPair(tmp20354)

var ifres20349 Obj

if True == tmp20355 {
tmp20351 := PrimTail(V7123)

tmp20352 := PrimHead(tmp20351)

tmp20353 := PrimEqual(MakeString("--help"), tmp20352)

var ifres20350 Obj

if True == tmp20353 {
ifres20350 = True


} else {
ifres20350 = False


}

ifres20349 = ifres20350


} else {
ifres20349 = False


}

var ifres20348 Obj

if True == ifres20349 {
ifres20348 = True


} else {
ifres20348 = False


}

ifres20347 = ifres20348


} else {
ifres20347 = False


}

if True == ifres20347 {
tmp20276 := PrimHead(V7123)

tmp20277 := Call(__e, PrimFunc(symshen_4x_4launcher_4help_1text), tmp20276)


tmp20278 := PrimCons(tmp20277, Nil)

__e.Return(PrimCons(symshow_1help, tmp20278))
return


} else {
tmp20345 := PrimIsPair(V7123)

var ifres20336 Obj

if True == tmp20345 {
tmp20343 := PrimTail(V7123)

tmp20344 := PrimIsPair(tmp20343)

var ifres20338 Obj

if True == tmp20344 {
tmp20340 := PrimTail(V7123)

tmp20341 := PrimHead(tmp20340)

tmp20342 := PrimEqual(MakeString("--version"), tmp20341)

var ifres20339 Obj

if True == tmp20342 {
ifres20339 = True


} else {
ifres20339 = False


}

ifres20338 = ifres20339


} else {
ifres20338 = False


}

var ifres20337 Obj

if True == ifres20338 {
ifres20337 = True


} else {
ifres20337 = False


}

ifres20336 = ifres20337


} else {
ifres20336 = False


}

if True == ifres20336 {
tmp20279 := Call(__e, PrimFunc(symshen_4x_4launcher_4version_1string))


tmp20280 := PrimCons(tmp20279, Nil)

__e.Return(PrimCons(symsuccess, tmp20280))
return


} else {
tmp20334 := PrimIsPair(V7123)

var ifres20325 Obj

if True == tmp20334 {
tmp20332 := PrimTail(V7123)

tmp20333 := PrimIsPair(tmp20332)

var ifres20327 Obj

if True == tmp20333 {
tmp20329 := PrimTail(V7123)

tmp20330 := PrimHead(tmp20329)

tmp20331 := PrimEqual(MakeString("repl"), tmp20330)

var ifres20328 Obj

if True == tmp20331 {
ifres20328 = True


} else {
ifres20328 = False


}

ifres20327 = ifres20328


} else {
ifres20327 = False


}

var ifres20326 Obj

if True == ifres20327 {
ifres20326 = True


} else {
ifres20326 = False


}

ifres20325 = ifres20326


} else {
ifres20325 = False


}

if True == ifres20325 {
tmp20281 := PrimTail(V7123)

tmp20282 := PrimTail(tmp20281)

__e.Return(PrimCons(symlaunch_1repl, tmp20282))
return


} else {
tmp20323 := PrimIsPair(V7123)

var ifres20309 Obj

if True == tmp20323 {
tmp20321 := PrimTail(V7123)

tmp20322 := PrimIsPair(tmp20321)

var ifres20311 Obj

if True == tmp20322 {
tmp20318 := PrimTail(V7123)

tmp20319 := PrimHead(tmp20318)

tmp20320 := PrimEqual(MakeString("script"), tmp20319)

var ifres20313 Obj

if True == tmp20320 {
tmp20315 := PrimTail(V7123)

tmp20316 := PrimTail(tmp20315)

tmp20317 := PrimIsPair(tmp20316)

var ifres20314 Obj

if True == tmp20317 {
ifres20314 = True


} else {
ifres20314 = False


}

ifres20313 = ifres20314


} else {
ifres20313 = False


}

var ifres20312 Obj

if True == ifres20313 {
ifres20312 = True


} else {
ifres20312 = False


}

ifres20311 = ifres20312


} else {
ifres20311 = False


}

var ifres20310 Obj

if True == ifres20311 {
ifres20310 = True


} else {
ifres20310 = False


}

ifres20309 = ifres20310


} else {
ifres20309 = False


}

if True == ifres20309 {
tmp20283 := PrimTail(V7123)

tmp20284 := PrimTail(tmp20283)

tmp20285 := PrimHead(tmp20284)

tmp20286 := PrimTail(V7123)

tmp20287 := PrimTail(tmp20286)

tmp20288 := PrimTail(tmp20287)

__e.TailApply(PrimFunc(symshen_4x_4launcher_4script_1command), tmp20285, tmp20288)
return


} else {
tmp20307 := PrimIsPair(V7123)

var ifres20298 Obj

if True == tmp20307 {
tmp20305 := PrimTail(V7123)

tmp20306 := PrimIsPair(tmp20305)

var ifres20300 Obj

if True == tmp20306 {
tmp20302 := PrimTail(V7123)

tmp20303 := PrimHead(tmp20302)

tmp20304 := PrimEqual(MakeString("eval"), tmp20303)

var ifres20301 Obj

if True == tmp20304 {
ifres20301 = True


} else {
ifres20301 = False


}

ifres20300 = ifres20301


} else {
ifres20300 = False


}

var ifres20299 Obj

if True == ifres20300 {
ifres20299 = True


} else {
ifres20299 = False


}

ifres20298 = ifres20299


} else {
ifres20298 = False


}

if True == ifres20298 {
tmp20289 := PrimTail(V7123)

tmp20290 := PrimTail(tmp20289)

__e.TailApply(PrimFunc(symshen_4x_4launcher_4eval_1command), tmp20290)
return


} else {
tmp20296 := PrimIsPair(V7123)

var ifres20292 Obj

if True == tmp20296 {
tmp20294 := PrimTail(V7123)

tmp20295 := PrimIsPair(tmp20294)

var ifres20293 Obj

if True == tmp20295 {
ifres20293 = True


} else {
ifres20293 = False


}

ifres20292 = ifres20293


} else {
ifres20292 = False


}

if True == ifres20292 {
__e.Return(PrimCons(symunknown_1arguments, V7123))
return
} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4x_4launcher_4launch_1shen)
return
}


}


}


}


}


}


}


}, 1)

tmp20363 := Call(__e, ns2_1set, symshen_4x_4launcher_4launch_1shen, tmp20275)


_ = tmp20363

tmp20364 := MakeNative(func(__e *ControlFlow) {
V7126 := __e.Get(1)
_ = V7126
tmp20463 := PrimIsPair(V7126)

var ifres20455 Obj

if True == tmp20463 {
tmp20461 := PrimHead(V7126)

tmp20462 := PrimEqual(symsuccess, tmp20461)

var ifres20457 Obj

if True == tmp20462 {
tmp20459 := PrimTail(V7126)

tmp20460 := PrimEqual(Nil, tmp20459)

var ifres20458 Obj

if True == tmp20460 {
ifres20458 = True


} else {
ifres20458 = False


}

ifres20457 = ifres20458


} else {
ifres20457 = False


}

var ifres20456 Obj

if True == ifres20457 {
ifres20456 = True


} else {
ifres20456 = False


}

ifres20455 = ifres20456


} else {
ifres20455 = False


}

if True == ifres20455 {
__e.Return(symshen_4x_4launcher_4done)
return
} else {
tmp20453 := PrimIsPair(V7126)

var ifres20440 Obj

if True == tmp20453 {
tmp20451 := PrimHead(V7126)

tmp20452 := PrimEqual(symsuccess, tmp20451)

var ifres20442 Obj

if True == tmp20452 {
tmp20449 := PrimTail(V7126)

tmp20450 := PrimIsPair(tmp20449)

var ifres20444 Obj

if True == tmp20450 {
tmp20446 := PrimTail(V7126)

tmp20447 := PrimTail(tmp20446)

tmp20448 := PrimEqual(Nil, tmp20447)

var ifres20445 Obj

if True == tmp20448 {
ifres20445 = True


} else {
ifres20445 = False


}

ifres20444 = ifres20445


} else {
ifres20444 = False


}

var ifres20443 Obj

if True == ifres20444 {
ifres20443 = True


} else {
ifres20443 = False


}

ifres20442 = ifres20443


} else {
ifres20442 = False


}

var ifres20441 Obj

if True == ifres20442 {
ifres20441 = True


} else {
ifres20441 = False


}

ifres20440 = ifres20441


} else {
ifres20440 = False


}

if True == ifres20440 {
tmp20365 := PrimTail(V7126)

tmp20366 := PrimHead(tmp20365)

tmp20367 := Call(__e, PrimFunc(symshen_4app), tmp20366, MakeString("\n"), symshen_4a)


tmp20368 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp20367, tmp20368)
return


} else {
tmp20438 := PrimIsPair(V7126)

var ifres20425 Obj

if True == tmp20438 {
tmp20436 := PrimHead(V7126)

tmp20437 := PrimEqual(symerror, tmp20436)

var ifres20427 Obj

if True == tmp20437 {
tmp20434 := PrimTail(V7126)

tmp20435 := PrimIsPair(tmp20434)

var ifres20429 Obj

if True == tmp20435 {
tmp20431 := PrimTail(V7126)

tmp20432 := PrimTail(tmp20431)

tmp20433 := PrimEqual(Nil, tmp20432)

var ifres20430 Obj

if True == tmp20433 {
ifres20430 = True


} else {
ifres20430 = False


}

ifres20429 = ifres20430


} else {
ifres20429 = False


}

var ifres20428 Obj

if True == ifres20429 {
ifres20428 = True


} else {
ifres20428 = False


}

ifres20427 = ifres20428


} else {
ifres20427 = False


}

var ifres20426 Obj

if True == ifres20427 {
ifres20426 = True


} else {
ifres20426 = False


}

ifres20425 = ifres20426


} else {
ifres20425 = False


}

if True == ifres20425 {
tmp20369 := PrimTail(V7126)

tmp20370 := PrimHead(tmp20369)

tmp20371 := Call(__e, PrimFunc(symshen_4app), tmp20370, MakeString("\n"), symshen_4a)


tmp20372 := PrimStringConcat(MakeString("ERROR: "), tmp20371)

tmp20373 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp20372, tmp20373)
return


} else {
tmp20423 := PrimIsPair(V7126)

var ifres20419 Obj

if True == tmp20423 {
tmp20421 := PrimHead(V7126)

tmp20422 := PrimEqual(symlaunch_1repl, tmp20421)

var ifres20420 Obj

if True == tmp20422 {
ifres20420 = True


} else {
ifres20420 = False


}

ifres20419 = ifres20420


} else {
ifres20419 = False


}

if True == ifres20419 {
__e.TailApply(PrimFunc(symshen_4repl))
return
} else {
tmp20417 := PrimIsPair(V7126)

var ifres20404 Obj

if True == tmp20417 {
tmp20415 := PrimHead(V7126)

tmp20416 := PrimEqual(symshow_1help, tmp20415)

var ifres20406 Obj

if True == tmp20416 {
tmp20413 := PrimTail(V7126)

tmp20414 := PrimIsPair(tmp20413)

var ifres20408 Obj

if True == tmp20414 {
tmp20410 := PrimTail(V7126)

tmp20411 := PrimTail(tmp20410)

tmp20412 := PrimEqual(Nil, tmp20411)

var ifres20409 Obj

if True == tmp20412 {
ifres20409 = True


} else {
ifres20409 = False


}

ifres20408 = ifres20409


} else {
ifres20408 = False


}

var ifres20407 Obj

if True == ifres20408 {
ifres20407 = True


} else {
ifres20407 = False


}

ifres20406 = ifres20407


} else {
ifres20406 = False


}

var ifres20405 Obj

if True == ifres20406 {
ifres20405 = True


} else {
ifres20405 = False


}

ifres20404 = ifres20405


} else {
ifres20404 = False


}

if True == ifres20404 {
tmp20374 := PrimTail(V7126)

tmp20375 := PrimHead(tmp20374)

tmp20376 := Call(__e, PrimFunc(symshen_4app), tmp20375, MakeString("\n"), symshen_4a)


tmp20377 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp20376, tmp20377)
return


} else {
tmp20402 := PrimIsPair(V7126)

var ifres20389 Obj

if True == tmp20402 {
tmp20400 := PrimHead(V7126)

tmp20401 := PrimEqual(symunknown_1arguments, tmp20400)

var ifres20391 Obj

if True == tmp20401 {
tmp20398 := PrimTail(V7126)

tmp20399 := PrimIsPair(tmp20398)

var ifres20393 Obj

if True == tmp20399 {
tmp20395 := PrimTail(V7126)

tmp20396 := PrimTail(tmp20395)

tmp20397 := PrimIsPair(tmp20396)

var ifres20394 Obj

if True == tmp20397 {
ifres20394 = True


} else {
ifres20394 = False


}

ifres20393 = ifres20394


} else {
ifres20393 = False


}

var ifres20392 Obj

if True == ifres20393 {
ifres20392 = True


} else {
ifres20392 = False


}

ifres20391 = ifres20392


} else {
ifres20391 = False


}

var ifres20390 Obj

if True == ifres20391 {
ifres20390 = True


} else {
ifres20390 = False


}

ifres20389 = ifres20390


} else {
ifres20389 = False


}

if True == ifres20389 {
tmp20378 := PrimTail(V7126)

tmp20379 := PrimTail(tmp20378)

tmp20380 := PrimHead(tmp20379)

tmp20381 := PrimTail(V7126)

tmp20382 := PrimHead(tmp20381)

tmp20383 := Call(__e, PrimFunc(symshen_4app), tmp20382, MakeString(" --help' for more information.\n"), symshen_4a)


tmp20384 := PrimStringConcat(MakeString("\nTry `"), tmp20383)

tmp20385 := Call(__e, PrimFunc(symshen_4app), tmp20380, tmp20384, symshen_4a)


tmp20386 := PrimStringConcat(MakeString("ERROR: Invalid argument: "), tmp20385)

tmp20387 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp20386, tmp20387)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4x_4launcher_4default_1handle_1result)
return
}


}


}


}


}


}


}, 1)

tmp20464 := Call(__e, ns2_1set, symshen_4x_4launcher_4default_1handle_1result, tmp20364)


_ = tmp20464

tmp20465 := MakeNative(func(__e *ControlFlow) {
V7127 := __e.Get(1)
_ = V7127
tmp20466 := Call(__e, PrimFunc(symshen_4x_4launcher_4launch_1shen), V7127)


__e.TailApply(PrimFunc(symshen_4x_4launcher_4default_1handle_1result), tmp20466)
return


}, 1)

__e.TailApply(ns2_1set, symshen_4x_4launcher_4main, tmp20465)
return




}, 0)

