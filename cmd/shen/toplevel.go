package main

import . "github.com/tiancaiamao/shen-go/kl"

var TopLevelMain = MakeNative(func(__e *ControlFlow) {
tmp8438 := MakeNative(func(__e *ControlFlow) {
tmp8439 := Call(__e, PrimFunc(symshen_4credits))


_ = tmp8439

__e.TailApply(PrimFunc(symshen_4loop))
return


}, 0)

tmp8440 := Call(__e, ns2_1set, symshen_4shen, tmp8438)


_ = tmp8440

tmp8441 := MakeNative(func(__e *ControlFlow) {
tmp8442 := Call(__e, PrimFunc(symshen_4initialise__environment))


_ = tmp8442

tmp8443 := Call(__e, PrimFunc(symshen_4prompt))


_ = tmp8443

tmp8444 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4read_1evaluate_1print))
return
}, 0)

tmp8445 := MakeNative(func(__e *ControlFlow) {
Z5248 := __e.Get(1)
_ = Z5248
tmp8446 := PrimErrorToString(Z5248)

tmp8447 := Call(__e, PrimFunc(symstoutput))


tmp8448 := Call(__e, PrimFunc(sympr), tmp8446, tmp8447)


_ = tmp8448

__e.TailApply(PrimFunc(symnl), MakeNumber(0))
return


}, 1)

tmp8449 := Call(__e, try_1catch, tmp8444, tmp8445)


_ = tmp8449

__e.TailApply(PrimFunc(symshen_4loop))
return


}, 0)

tmp8450 := Call(__e, ns2_1set, symshen_4loop, tmp8441)


_ = tmp8450

tmp8451 := MakeNative(func(__e *ControlFlow) {
tmp8452 := Call(__e, PrimFunc(symstoutput))


tmp8453 := Call(__e, PrimFunc(sympr), MakeString("\nShen, www.shenlanguage.org, copyright (C) 2010-2024, Mark Tarver\n"), tmp8452)


_ = tmp8453

tmp8454 := PrimValue(sym_dversion_d)

tmp8455 := PrimValue(sym_dlanguage_d)

tmp8456 := PrimValue(sym_dimplementation_d)

tmp8457 := PrimValue(sym_drelease_d)

tmp8458 := Call(__e, PrimFunc(symshen_4app), tmp8457, MakeString("\n"), symshen_4a)


tmp8459 := PrimStringConcat(MakeString(" "), tmp8458)

tmp8460 := Call(__e, PrimFunc(symshen_4app), tmp8456, tmp8459, symshen_4a)


tmp8461 := PrimStringConcat(MakeString(", platform: "), tmp8460)

tmp8462 := Call(__e, PrimFunc(symshen_4app), tmp8455, tmp8461, symshen_4a)


tmp8463 := PrimStringConcat(MakeString(", language: "), tmp8462)

tmp8464 := Call(__e, PrimFunc(symshen_4app), tmp8454, tmp8463, symshen_4a)


tmp8465 := PrimStringConcat(MakeString("version: S"), tmp8464)

tmp8466 := Call(__e, PrimFunc(symstoutput))


tmp8467 := Call(__e, PrimFunc(sympr), tmp8465, tmp8466)


_ = tmp8467

tmp8468 := PrimValue(sym_dport_d)

tmp8469 := PrimValue(sym_dporters_d)

tmp8470 := Call(__e, PrimFunc(symshen_4app), tmp8469, MakeString("\n\n"), symshen_4a)


tmp8471 := PrimStringConcat(MakeString(", ported by "), tmp8470)

tmp8472 := Call(__e, PrimFunc(symshen_4app), tmp8468, tmp8471, symshen_4a)


tmp8473 := PrimStringConcat(MakeString("port "), tmp8472)

tmp8474 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp8473, tmp8474)
return


}, 0)

tmp8475 := Call(__e, ns2_1set, symshen_4credits, tmp8451)


_ = tmp8475

tmp8476 := MakeNative(func(__e *ControlFlow) {
tmp8477 := PrimSet(symshen_4_dcall_d, MakeNumber(0))

_ = tmp8477

__e.Return(PrimSet(symshen_4_dinfs_d, MakeNumber(0)))
return


}, 0)

tmp8478 := Call(__e, ns2_1set, symshen_4initialise__environment, tmp8476)


_ = tmp8478

tmp8479 := MakeNative(func(__e *ControlFlow) {
tmp8491 := PrimValue(symshen_4_dtc_d)

if True == tmp8491 {
tmp8480 := PrimValue(symshen_4_dhistory_d)

tmp8481 := Call(__e, PrimFunc(symlength), tmp8480)


tmp8482 := Call(__e, PrimFunc(symshen_4app), tmp8481, MakeString("+) "), symshen_4a)


tmp8483 := PrimStringConcat(MakeString("\n("), tmp8482)

tmp8484 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp8483, tmp8484)
return


} else {
tmp8485 := PrimValue(symshen_4_dhistory_d)

tmp8486 := Call(__e, PrimFunc(symlength), tmp8485)


tmp8487 := Call(__e, PrimFunc(symshen_4app), tmp8486, MakeString("-) "), symshen_4a)


tmp8488 := PrimStringConcat(MakeString("\n("), tmp8487)

tmp8489 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp8488, tmp8489)
return


}


}, 0)

tmp8492 := Call(__e, ns2_1set, symshen_4prompt, tmp8479)


_ = tmp8492

tmp8493 := MakeNative(func(__e *ControlFlow) {
tmp8494 := MakeNative(func(__e *ControlFlow) {
W5249 := __e.Get(1)
_ = W5249
tmp8495 := MakeNative(func(__e *ControlFlow) {
W5250 := __e.Get(1)
_ = W5250
tmp8496 := MakeNative(func(__e *ControlFlow) {
W5251 := __e.Get(1)
_ = W5251
tmp8497 := PrimValue(symshen_4_dtc_d)

__e.TailApply(PrimFunc(symshen_4evaluate_1lineread), W5250, W5251, tmp8497)
return


}, 1)

tmp8498 := Call(__e, PrimFunc(symshen_4update_1history))


__e.TailApply(tmp8496, tmp8498)
return


}, 1)

tmp8499 := Call(__e, PrimFunc(symstinput))


tmp8500 := Call(__e, PrimFunc(symlineread), tmp8499)


tmp8501 := Call(__e, PrimFunc(symshen_4package_1user_1input), W5249, tmp8500)


__e.TailApply(tmp8495, tmp8501)
return


}, 1)

tmp8502 := PrimValue(symshen_4_dpackage_d)

__e.TailApply(tmp8494, tmp8502)
return


}, 0)

tmp8503 := Call(__e, ns2_1set, symshen_4read_1evaluate_1print, tmp8493)


_ = tmp8503

tmp8504 := MakeNative(func(__e *ControlFlow) {
V5252 := __e.Get(1)
_ = V5252
V5253 := __e.Get(2)
_ = V5253
tmp8511 := PrimEqual(symnull, V5252)

if True == tmp8511 {
__e.Return(V5253)
return
} else {
tmp8505 := MakeNative(func(__e *ControlFlow) {
W5254 := __e.Get(1)
_ = W5254
tmp8506 := MakeNative(func(__e *ControlFlow) {
W5255 := __e.Get(1)
_ = W5255
tmp8507 := MakeNative(func(__e *ControlFlow) {
Z5256 := __e.Get(1)
_ = Z5256
__e.TailApply(PrimFunc(symshen_4pui_1h), W5254, W5255, Z5256)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp8507, V5253)
return


}, 1)

tmp8508 := Call(__e, PrimFunc(symexternal), V5252)


__e.TailApply(tmp8506, tmp8508)
return


}, 1)

tmp8509 := PrimStr(V5252)

__e.TailApply(tmp8505, tmp8509)
return


}


}, 2)

tmp8512 := Call(__e, ns2_1set, symshen_4package_1user_1input, tmp8504)


_ = tmp8512

tmp8513 := MakeNative(func(__e *ControlFlow) {
V5261 := __e.Get(1)
_ = V5261
V5262 := __e.Get(2)
_ = V5262
V5263 := __e.Get(3)
_ = V5263
tmp8554 := PrimIsPair(V5263)

var ifres8541 Obj

if True == tmp8554 {
tmp8552 := PrimHead(V5263)

tmp8553 := PrimEqual(symfn, tmp8552)

var ifres8543 Obj

if True == tmp8553 {
tmp8550 := PrimTail(V5263)

tmp8551 := PrimIsPair(tmp8550)

var ifres8545 Obj

if True == tmp8551 {
tmp8547 := PrimTail(V5263)

tmp8548 := PrimTail(tmp8547)

tmp8549 := PrimEqual(Nil, tmp8548)

var ifres8546 Obj

if True == tmp8549 {
ifres8546 = True


} else {
ifres8546 = False


}

ifres8545 = ifres8546


} else {
ifres8545 = False


}

var ifres8544 Obj

if True == ifres8545 {
ifres8544 = True


} else {
ifres8544 = False


}

ifres8543 = ifres8544


} else {
ifres8543 = False


}

var ifres8542 Obj

if True == ifres8543 {
ifres8542 = True


} else {
ifres8542 = False


}

ifres8541 = ifres8542


} else {
ifres8541 = False


}

if True == ifres8541 {
tmp8519 := PrimTail(V5263)

tmp8520 := PrimHead(tmp8519)

tmp8521 := Call(__e, PrimFunc(symshen_4internal_2), tmp8520, V5261, V5262)


if True == tmp8521 {
tmp8514 := PrimTail(V5263)

tmp8515 := PrimHead(tmp8514)

tmp8516 := Call(__e, PrimFunc(symshen_4intern_1in_1package), V5261, tmp8515)


tmp8517 := PrimCons(tmp8516, Nil)

__e.Return(PrimCons(symfn, tmp8517))
return


} else {
__e.Return(V5263)
return
}


} else {
tmp8539 := PrimIsPair(V5263)

if True == tmp8539 {
tmp8536 := PrimHead(V5263)

tmp8537 := Call(__e, PrimFunc(symshen_4internal_2), tmp8536, V5261, V5262)


if True == tmp8537 {
tmp8522 := PrimHead(V5263)

tmp8523 := Call(__e, PrimFunc(symshen_4intern_1in_1package), V5261, tmp8522)


tmp8524 := MakeNative(func(__e *ControlFlow) {
Z5264 := __e.Get(1)
_ = Z5264
__e.TailApply(PrimFunc(symshen_4pui_1h), V5261, V5262, Z5264)
return
}, 1)

tmp8525 := PrimTail(V5263)

tmp8526 := Call(__e, PrimFunc(symmap), tmp8524, tmp8525)


__e.Return(PrimCons(tmp8523, tmp8526))
return


} else {
tmp8533 := PrimHead(V5263)

tmp8534 := PrimIsPair(tmp8533)

if True == tmp8534 {
tmp8527 := MakeNative(func(__e *ControlFlow) {
Z5265 := __e.Get(1)
_ = Z5265
__e.TailApply(PrimFunc(symshen_4pui_1h), V5261, V5262, Z5265)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp8527, V5263)
return


} else {
tmp8528 := PrimHead(V5263)

tmp8529 := MakeNative(func(__e *ControlFlow) {
Z5266 := __e.Get(1)
_ = Z5266
__e.TailApply(PrimFunc(symshen_4pui_1h), V5261, V5262, Z5266)
return
}, 1)

tmp8530 := PrimTail(V5263)

tmp8531 := Call(__e, PrimFunc(symmap), tmp8529, tmp8530)


__e.Return(PrimCons(tmp8528, tmp8531))
return


}


}


} else {
__e.Return(V5263)
return
}


}


}, 3)

tmp8555 := Call(__e, ns2_1set, symshen_4pui_1h, tmp8513)


_ = tmp8555

tmp8556 := MakeNative(func(__e *ControlFlow) {
tmp8557 := Call(__e, PrimFunc(symit))


tmp8558 := Call(__e, PrimFunc(symshen_4trim_1it), tmp8557)


tmp8559 := PrimValue(symshen_4_dhistory_d)

tmp8560 := PrimCons(tmp8558, tmp8559)

__e.Return(PrimSet(symshen_4_dhistory_d, tmp8560))
return


}, 0)

tmp8561 := Call(__e, ns2_1set, symshen_4update_1history, tmp8556)


_ = tmp8561

tmp8562 := MakeNative(func(__e *ControlFlow) {
V5267 := __e.Get(1)
_ = V5267
tmp8570 := Call(__e, PrimFunc(symshen_4_7string_2), V5267)


var ifres8565 Obj

if True == tmp8570 {
tmp8567 := Call(__e, PrimFunc(symhdstr), V5267)


tmp8568 := PrimStringToNumber(tmp8567)

tmp8569 := Call(__e, PrimFunc(symshen_4whitespace_2), tmp8568)


var ifres8566 Obj

if True == tmp8569 {
ifres8566 = True


} else {
ifres8566 = False


}

ifres8565 = ifres8566


} else {
ifres8565 = False


}

if True == ifres8565 {
tmp8563 := PrimTailString(V5267)

__e.TailApply(PrimFunc(symshen_4trim_1it), tmp8563)
return


} else {
__e.Return(V5267)
return
}


}, 1)

tmp8571 := Call(__e, ns2_1set, symshen_4trim_1it, tmp8562)


_ = tmp8571

tmp8572 := MakeNative(func(__e *ControlFlow) {
V5286 := __e.Get(1)
_ = V5286
V5287 := __e.Get(2)
_ = V5287
V5288 := __e.Get(3)
_ = V5288
tmp8702 := PrimIsPair(V5286)

var ifres8671 Obj

if True == tmp8702 {
tmp8700 := PrimTail(V5286)

tmp8701 := PrimEqual(Nil, tmp8700)

var ifres8673 Obj

if True == tmp8701 {
tmp8699 := PrimIsPair(V5287)

var ifres8675 Obj

if True == tmp8699 {
tmp8697 := PrimHead(V5287)

tmp8698 := Call(__e, PrimFunc(symshen_4_7string_2), tmp8697)


var ifres8677 Obj

if True == tmp8698 {
tmp8694 := PrimHead(V5287)

tmp8695 := Call(__e, PrimFunc(symhdstr), tmp8694)


tmp8696 := PrimEqual(MakeString("!"), tmp8695)

var ifres8679 Obj

if True == tmp8696 {
tmp8691 := PrimHead(V5287)

tmp8692 := PrimTailString(tmp8691)

tmp8693 := Call(__e, PrimFunc(symshen_4_7string_2), tmp8692)


var ifres8681 Obj

if True == tmp8693 {
tmp8687 := PrimHead(V5287)

tmp8688 := PrimTailString(tmp8687)

tmp8689 := Call(__e, PrimFunc(symhdstr), tmp8688)


tmp8690 := PrimEqual(MakeString("!"), tmp8689)

var ifres8683 Obj

if True == tmp8690 {
tmp8685 := PrimTail(V5287)

tmp8686 := PrimIsPair(tmp8685)

var ifres8684 Obj

if True == tmp8686 {
ifres8684 = True


} else {
ifres8684 = False


}

ifres8683 = ifres8684


} else {
ifres8683 = False


}

var ifres8682 Obj

if True == ifres8683 {
ifres8682 = True


} else {
ifres8682 = False


}

ifres8681 = ifres8682


} else {
ifres8681 = False


}

var ifres8680 Obj

if True == ifres8681 {
ifres8680 = True


} else {
ifres8680 = False


}

ifres8679 = ifres8680


} else {
ifres8679 = False


}

var ifres8678 Obj

if True == ifres8679 {
ifres8678 = True


} else {
ifres8678 = False


}

ifres8677 = ifres8678


} else {
ifres8677 = False


}

var ifres8676 Obj

if True == ifres8677 {
ifres8676 = True


} else {
ifres8676 = False


}

ifres8675 = ifres8676


} else {
ifres8675 = False


}

var ifres8674 Obj

if True == ifres8675 {
ifres8674 = True


} else {
ifres8674 = False


}

ifres8673 = ifres8674


} else {
ifres8673 = False


}

var ifres8672 Obj

if True == ifres8673 {
ifres8672 = True


} else {
ifres8672 = False


}

ifres8671 = ifres8672


} else {
ifres8671 = False


}

if True == ifres8671 {
tmp8573 := MakeNative(func(__e *ControlFlow) {
W5289 := __e.Get(1)
_ = W5289
tmp8574 := MakeNative(func(__e *ControlFlow) {
W5290 := __e.Get(1)
_ = W5290
tmp8575 := MakeNative(func(__e *ControlFlow) {
W5291 := __e.Get(1)
_ = W5291
__e.TailApply(PrimFunc(symshen_4evaluate_1lineread), W5289, W5290, V5288)
return
}, 1)

tmp8576 := PrimTail(V5287)

tmp8577 := PrimHead(tmp8576)

tmp8578 := Call(__e, PrimFunc(symshen_4app), tmp8577, MakeString("\n"), symshen_4a)


tmp8579 := Call(__e, PrimFunc(symstoutput))


tmp8580 := Call(__e, PrimFunc(sympr), tmp8578, tmp8579)


__e.TailApply(tmp8575, tmp8580)
return


}, 1)

tmp8581 := PrimTail(V5287)

tmp8582 := PrimHead(tmp8581)

tmp8583 := PrimTail(V5287)

tmp8584 := PrimCons(tmp8582, tmp8583)

tmp8585 := PrimSet(symshen_4_dhistory_d, tmp8584)

__e.TailApply(tmp8574, tmp8585)
return


}, 1)

tmp8586 := PrimTail(V5287)

tmp8587 := PrimHead(tmp8586)

tmp8588 := Call(__e, PrimFunc(symread_1from_1string), tmp8587)


__e.TailApply(tmp8573, tmp8588)
return


} else {
tmp8669 := PrimIsPair(V5286)

var ifres8653 Obj

if True == tmp8669 {
tmp8667 := PrimTail(V5286)

tmp8668 := PrimEqual(Nil, tmp8667)

var ifres8655 Obj

if True == tmp8668 {
tmp8666 := PrimIsPair(V5287)

var ifres8657 Obj

if True == tmp8666 {
tmp8664 := PrimHead(V5287)

tmp8665 := Call(__e, PrimFunc(symshen_4_7string_2), tmp8664)


var ifres8659 Obj

if True == tmp8665 {
tmp8661 := PrimHead(V5287)

tmp8662 := Call(__e, PrimFunc(symhdstr), tmp8661)


tmp8663 := PrimEqual(MakeString("!"), tmp8662)

var ifres8660 Obj

if True == tmp8663 {
ifres8660 = True


} else {
ifres8660 = False


}

ifres8659 = ifres8660


} else {
ifres8659 = False


}

var ifres8658 Obj

if True == ifres8659 {
ifres8658 = True


} else {
ifres8658 = False


}

ifres8657 = ifres8658


} else {
ifres8657 = False


}

var ifres8656 Obj

if True == ifres8657 {
ifres8656 = True


} else {
ifres8656 = False


}

ifres8655 = ifres8656


} else {
ifres8655 = False


}

var ifres8654 Obj

if True == ifres8655 {
ifres8654 = True


} else {
ifres8654 = False


}

ifres8653 = ifres8654


} else {
ifres8653 = False


}

if True == ifres8653 {
tmp8589 := MakeNative(func(__e *ControlFlow) {
W5292 := __e.Get(1)
_ = W5292
tmp8590 := MakeNative(func(__e *ControlFlow) {
W5293 := __e.Get(1)
_ = W5293
tmp8591 := MakeNative(func(__e *ControlFlow) {
W5294 := __e.Get(1)
_ = W5294
tmp8592 := MakeNative(func(__e *ControlFlow) {
W5295 := __e.Get(1)
_ = W5295
tmp8593 := MakeNative(func(__e *ControlFlow) {
W5296 := __e.Get(1)
_ = W5296
__e.TailApply(PrimFunc(symshen_4evaluate_1lineread), W5295, W5296, V5288)
return
}, 1)

tmp8594 := PrimTail(V5287)

tmp8595 := PrimCons(W5293, tmp8594)

tmp8596 := PrimSet(symshen_4_dhistory_d, tmp8595)

__e.TailApply(tmp8593, tmp8596)
return


}, 1)

tmp8597 := Call(__e, PrimFunc(symread_1from_1string), W5293)


__e.TailApply(tmp8592, tmp8597)
return


}, 1)

tmp8598 := Call(__e, PrimFunc(symshen_4app), W5293, MakeString("\n"), symshen_4a)


tmp8599 := Call(__e, PrimFunc(symstoutput))


tmp8600 := Call(__e, PrimFunc(sympr), tmp8598, tmp8599)


__e.TailApply(tmp8591, tmp8600)
return


}, 1)

tmp8601 := PrimHead(V5287)

tmp8602 := PrimTailString(tmp8601)

tmp8603 := PrimTail(V5287)

tmp8604 := Call(__e, PrimFunc(symshen_4use_1history), W5292, tmp8602, tmp8603)


__e.TailApply(tmp8590, tmp8604)
return


}, 1)

tmp8610 := PrimHead(V5287)

tmp8611 := PrimTailString(tmp8610)

tmp8612 := PrimEqual(tmp8611, MakeString(""))

var ifres8605 Obj

if True == tmp8612 {
ifres8605 = Nil


} else {
tmp8606 := PrimHead(V5287)

tmp8607 := PrimTailString(tmp8606)

tmp8608 := Call(__e, PrimFunc(symread_1from_1string), tmp8607)


tmp8609 := PrimHead(tmp8608)

ifres8605 = tmp8609


}

__e.TailApply(tmp8589, ifres8605)
return


} else {
tmp8651 := PrimIsPair(V5286)

var ifres8635 Obj

if True == tmp8651 {
tmp8649 := PrimTail(V5286)

tmp8650 := PrimEqual(Nil, tmp8649)

var ifres8637 Obj

if True == tmp8650 {
tmp8648 := PrimIsPair(V5287)

var ifres8639 Obj

if True == tmp8648 {
tmp8646 := PrimHead(V5287)

tmp8647 := Call(__e, PrimFunc(symshen_4_7string_2), tmp8646)


var ifres8641 Obj

if True == tmp8647 {
tmp8643 := PrimHead(V5287)

tmp8644 := Call(__e, PrimFunc(symhdstr), tmp8643)


tmp8645 := PrimEqual(MakeString("%"), tmp8644)

var ifres8642 Obj

if True == tmp8645 {
ifres8642 = True


} else {
ifres8642 = False


}

ifres8641 = ifres8642


} else {
ifres8641 = False


}

var ifres8640 Obj

if True == ifres8641 {
ifres8640 = True


} else {
ifres8640 = False


}

ifres8639 = ifres8640


} else {
ifres8639 = False


}

var ifres8638 Obj

if True == ifres8639 {
ifres8638 = True


} else {
ifres8638 = False


}

ifres8637 = ifres8638


} else {
ifres8637 = False


}

var ifres8636 Obj

if True == ifres8637 {
ifres8636 = True


} else {
ifres8636 = False


}

ifres8635 = ifres8636


} else {
ifres8635 = False


}

if True == ifres8635 {
tmp8613 := MakeNative(func(__e *ControlFlow) {
W5297 := __e.Get(1)
_ = W5297
tmp8614 := MakeNative(func(__e *ControlFlow) {
W5298 := __e.Get(1)
_ = W5298
tmp8615 := MakeNative(func(__e *ControlFlow) {
W5299 := __e.Get(1)
_ = W5299
__e.TailApply(PrimFunc(symabort))
return
}, 1)

tmp8616 := PrimTail(V5287)

tmp8617 := PrimSet(symshen_4_dhistory_d, tmp8616)

__e.TailApply(tmp8615, tmp8617)
return


}, 1)

tmp8618 := PrimHead(V5287)

tmp8619 := PrimTailString(tmp8618)

tmp8620 := PrimTail(V5287)

tmp8621 := Call(__e, PrimFunc(symshen_4peek_1history), W5297, tmp8619, tmp8620)


__e.TailApply(tmp8614, tmp8621)
return


}, 1)

tmp8627 := PrimHead(V5287)

tmp8628 := PrimTailString(tmp8627)

tmp8629 := PrimEqual(tmp8628, MakeString(""))

var ifres8622 Obj

if True == tmp8629 {
ifres8622 = Nil


} else {
tmp8623 := PrimHead(V5287)

tmp8624 := PrimTailString(tmp8623)

tmp8625 := Call(__e, PrimFunc(symread_1from_1string), tmp8624)


tmp8626 := PrimHead(tmp8625)

ifres8622 = tmp8626


}

__e.TailApply(tmp8613, ifres8622)
return


} else {
tmp8633 := PrimEqual(True, V5288)

if True == tmp8633 {
__e.TailApply(PrimFunc(symshen_4check_1eval_1and_1print), V5286)
return
} else {
tmp8631 := PrimEqual(False, V5288)

if True == tmp8631 {
__e.TailApply(PrimFunc(symshen_4eval_1and_1print), V5286)
return
} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.evaluate-lineread")))
return
}


}


}


}


}


}, 3)

tmp8703 := Call(__e, ns2_1set, symshen_4evaluate_1lineread, tmp8572)


_ = tmp8703

tmp8704 := MakeNative(func(__e *ControlFlow) {
V5300 := __e.Get(1)
_ = V5300
V5301 := __e.Get(2)
_ = V5301
V5302 := __e.Get(3)
_ = V5302
tmp8710 := PrimIsInteger(V5300)

if True == tmp8710 {
tmp8705 := PrimNumberAdd(MakeNumber(1), V5300)

tmp8706 := Call(__e, PrimFunc(symreverse), V5302)


__e.TailApply(PrimFunc(symnth), tmp8705, tmp8706)
return


} else {
tmp8708 := PrimIsSymbol(V5300)

if True == tmp8708 {
__e.TailApply(PrimFunc(symshen_4string_1match), V5301, V5302)
return
} else {
__e.Return(PrimSimpleError(MakeString("! expects a number or a symbol\n")))
return
}


}


}, 3)

tmp8711 := Call(__e, ns2_1set, symshen_4use_1history, tmp8704)


_ = tmp8711

tmp8712 := MakeNative(func(__e *ControlFlow) {
V5303 := __e.Get(1)
_ = V5303
V5304 := __e.Get(2)
_ = V5304
V5305 := __e.Get(3)
_ = V5305
tmp8726 := PrimIsInteger(V5303)

if True == tmp8726 {
tmp8713 := PrimNumberAdd(MakeNumber(1), V5303)

tmp8714 := Call(__e, PrimFunc(symreverse), V5305)


tmp8715 := Call(__e, PrimFunc(symnth), tmp8713, tmp8714)


tmp8716 := Call(__e, PrimFunc(symshen_4app), tmp8715, MakeString(""), symshen_4a)


tmp8717 := PrimStringConcat(MakeString("\n"), tmp8716)

tmp8718 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp8717, tmp8718)
return


} else {
tmp8724 := PrimEqual(V5304, MakeString(""))

var ifres8721 Obj

if True == tmp8724 {
ifres8721 = True


} else {
tmp8723 := PrimIsSymbol(V5303)

var ifres8722 Obj

if True == tmp8723 {
ifres8722 = True


} else {
ifres8722 = False


}

ifres8721 = ifres8722


}

if True == ifres8721 {
tmp8719 := Call(__e, PrimFunc(symreverse), V5305)


__e.TailApply(PrimFunc(symshen_4recursive_1string_1match), MakeNumber(0), V5304, tmp8719)
return


} else {
__e.Return(PrimSimpleError(MakeString("% expects a number or a symbol\n")))
return
}


}


}, 3)

tmp8727 := Call(__e, ns2_1set, symshen_4peek_1history, tmp8712)


_ = tmp8727

tmp8728 := MakeNative(func(__e *ControlFlow) {
V5315 := __e.Get(1)
_ = V5315
V5316 := __e.Get(2)
_ = V5316
tmp8739 := PrimEqual(Nil, V5316)

if True == tmp8739 {
__e.Return(PrimSimpleError(MakeString("\ninput not found")))
return
} else {
tmp8737 := PrimIsPair(V5316)

var ifres8733 Obj

if True == tmp8737 {
tmp8735 := PrimHead(V5316)

tmp8736 := Call(__e, PrimFunc(symshen_4string_1prefix_2), V5315, tmp8735)


var ifres8734 Obj

if True == tmp8736 {
ifres8734 = True


} else {
ifres8734 = False


}

ifres8733 = ifres8734


} else {
ifres8733 = False


}

if True == ifres8733 {
__e.Return(PrimHead(V5316))
return
} else {
tmp8731 := PrimIsPair(V5316)

if True == tmp8731 {
tmp8729 := PrimTail(V5316)

__e.TailApply(PrimFunc(symshen_4string_1match), V5315, tmp8729)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.string-match")))
return
}


}


}


}, 2)

tmp8740 := Call(__e, ns2_1set, symshen_4string_1match, tmp8728)


_ = tmp8740

tmp8741 := MakeNative(func(__e *ControlFlow) {
V5324 := __e.Get(1)
_ = V5324
V5325 := __e.Get(2)
_ = V5325
tmp8778 := PrimEqual(MakeString(""), V5324)

if True == tmp8778 {
__e.Return(True)
return
} else {
tmp8776 := Call(__e, PrimFunc(symshen_4_7string_2), V5324)


var ifres8771 Obj

if True == tmp8776 {
tmp8773 := Call(__e, PrimFunc(symhdstr), V5324)


tmp8774 := PrimStringToNumber(tmp8773)

tmp8775 := Call(__e, PrimFunc(symshen_4whitespace_2), tmp8774)


var ifres8772 Obj

if True == tmp8775 {
ifres8772 = True


} else {
ifres8772 = False


}

ifres8771 = ifres8772


} else {
ifres8771 = False


}

if True == ifres8771 {
tmp8742 := PrimTailString(V5324)

__e.TailApply(PrimFunc(symshen_4string_1prefix_2), tmp8742, V5325)
return


} else {
tmp8769 := Call(__e, PrimFunc(symshen_4_7string_2), V5325)


var ifres8764 Obj

if True == tmp8769 {
tmp8766 := Call(__e, PrimFunc(symhdstr), V5325)


tmp8767 := PrimStringToNumber(tmp8766)

tmp8768 := Call(__e, PrimFunc(symshen_4whitespace_2), tmp8767)


var ifres8765 Obj

if True == tmp8768 {
ifres8765 = True


} else {
ifres8765 = False


}

ifres8764 = ifres8765


} else {
ifres8764 = False


}

if True == ifres8764 {
tmp8743 := PrimTailString(V5325)

__e.TailApply(PrimFunc(symshen_4string_1prefix_2), V5324, tmp8743)
return


} else {
tmp8762 := Call(__e, PrimFunc(symshen_4_7string_2), V5325)


var ifres8758 Obj

if True == tmp8762 {
tmp8760 := Call(__e, PrimFunc(symhdstr), V5325)


tmp8761 := PrimEqual(MakeString("("), tmp8760)

var ifres8759 Obj

if True == tmp8761 {
ifres8759 = True


} else {
ifres8759 = False


}

ifres8758 = ifres8759


} else {
ifres8758 = False


}

if True == ifres8758 {
tmp8744 := PrimTailString(V5325)

__e.TailApply(PrimFunc(symshen_4string_1prefix_2), V5324, tmp8744)
return


} else {
tmp8756 := Call(__e, PrimFunc(symshen_4_7string_2), V5324)


var ifres8748 Obj

if True == tmp8756 {
tmp8755 := Call(__e, PrimFunc(symshen_4_7string_2), V5325)


var ifres8750 Obj

if True == tmp8755 {
tmp8752 := Call(__e, PrimFunc(symhdstr), V5324)


tmp8753 := Call(__e, PrimFunc(symhdstr), V5325)


tmp8754 := PrimEqual(tmp8752, tmp8753)

var ifres8751 Obj

if True == tmp8754 {
ifres8751 = True


} else {
ifres8751 = False


}

ifres8750 = ifres8751


} else {
ifres8750 = False


}

var ifres8749 Obj

if True == ifres8750 {
ifres8749 = True


} else {
ifres8749 = False


}

ifres8748 = ifres8749


} else {
ifres8748 = False


}

if True == ifres8748 {
tmp8745 := PrimTailString(V5324)

tmp8746 := PrimTailString(V5325)

__e.TailApply(PrimFunc(symshen_4string_1prefix_2), tmp8745, tmp8746)
return


} else {
__e.Return(False)
return
}


}


}


}


}


}, 2)

tmp8779 := Call(__e, ns2_1set, symshen_4string_1prefix_2, tmp8741)


_ = tmp8779

tmp8780 := MakeNative(func(__e *ControlFlow) {
V5336 := __e.Get(1)
_ = V5336
V5337 := __e.Get(2)
_ = V5337
V5338 := __e.Get(3)
_ = V5338
tmp8795 := PrimEqual(Nil, V5338)

if True == tmp8795 {
__e.Return(symshen_4skip)
return
} else {
tmp8793 := PrimIsPair(V5338)

if True == tmp8793 {
tmp8788 := PrimHead(V5338)

tmp8789 := Call(__e, PrimFunc(symshen_4string_1prefix_2), V5337, tmp8788)


var ifres8781 Obj

if True == tmp8789 {
tmp8782 := PrimHead(V5338)

tmp8783 := Call(__e, PrimFunc(symshen_4app), tmp8782, MakeString("\n"), symshen_4a)


tmp8784 := PrimStringConcat(MakeString(". "), tmp8783)

tmp8785 := Call(__e, PrimFunc(symshen_4app), V5336, tmp8784, symshen_4a)


tmp8786 := Call(__e, PrimFunc(symstoutput))


tmp8787 := Call(__e, PrimFunc(sympr), tmp8785, tmp8786)


ifres8781 = tmp8787


} else {
ifres8781 = symshen_4skip


}

_ = ifres8781

tmp8790 := PrimNumberAdd(V5336, MakeNumber(1))

tmp8791 := PrimTail(V5338)

__e.TailApply(PrimFunc(symshen_4recursive_1string_1match), tmp8790, V5337, tmp8791)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.recursive-string-match")))
return
}


}


}, 3)

__e.TailApply(ns2_1set, symshen_4recursive_1string_1match, tmp8780)
return




}, 0)

